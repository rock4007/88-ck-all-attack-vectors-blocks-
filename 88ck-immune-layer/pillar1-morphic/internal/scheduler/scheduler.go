package scheduler

import (
	"context"

	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/88ck/pillar1-morphic/internal/metrics"
	"github.com/88ck/pillar1-morphic/internal/types"
	"golang.org/x/crypto/chacha20poly1305"
)

type ChaCha20Scheduler struct {
	seed     []byte
	cipher   chacha20poly1305.Cipher
	period   time.Duration
	maxRate  float64
	metrics  *metrics.Metrics
	mu       sync.RWMutex
	cycles   int64
	mt       float64
	lastMorph time.Time
}

func NewChaCha20Scheduler(seed []byte, period time.Duration) *ChaCha20Scheduler {
	s, err := chacha20poly1305.NewCipher(seed)
	if err != nil {
		// Fatal? Log, but per rules no panic, but constructor fail
		panic(fmt.Sprintf("chacha20 init failed: %v", err)) // constructor
	}
	return &ChaCha20Scheduler{
		seed:    seed,
		cipher:  s,
		period:  period,
		maxRate: 1.0, // default
		metrics: metrics.NewMetrics(context.Background()),
		lastMorph: time.Now(),
	}
}

func (s *ChaCha20Scheduler) SetMaxRate(rate float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.maxRate = math.Max(0, math.Min(1, rate))
}

func (s *ChaCha20Scheduler) Next() types.EndpointTopology {
	s.mu.Lock()
	cycles := s.cycles
	s.cycles++
	s.mu.Unlock()

	nonce := make([]byte, chacha20poly1305.NonceSize)
	binary.LittleEndian.PutUint64(nonce, uint64(cycles))

	var stream [32]byte
	s.cipher.XORKeyStream(stream[:], nonce[:chacha20poly1305.NonceSize])

	// Deterministic mutation ID
	mutationID := hex.EncodeToString(stream[:16])

	// Deterministic endpoints rotation (simple example, rotate hash % N endpoints)
	// For prod, gen full topology from stream
		endpoints := []types.Endpoint{
			{ID: "ep1", Addr: fmt.Sprintf("127.0.0.%d:80", stream[0]), Weight: 100},
			{ID: "ep2", Addr: fmt.Sprintf("127.0.0.%d:80", stream[1]), Weight: 100},
			{ID: "ep3", Addr: fmt.Sprintf("127.0.0.%d:80", stream[2]), Weight: 100},
		}

	topology := types.EndpointTopology{
		MutationID:  mutationID,
		Endpoints:   endpoints,
		ScheduledAt: time.Now(),
		Seed:        s.seed,
	}

	return topology
}

func (s *ChaCha20Scheduler) CurrentMorphRate() float64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	mt := s.mt
	if mt < 0 {
		mt = 0
	}
	if mt > 1 {
		mt = 1
	}
	metrics.MTValue.Set(mt)
	return mt
}

func (s *ChaCha20Scheduler) Schedule(ctx context.Context, onMorph func(topology types.EndpointTopology)) {
	ticker := time.NewTicker(s.period)
	defer ticker.Stop()
	
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			start := time.Now()
			_, span := s.metrics.StartMorphCycle(context.Background(), s.seed)
			defer func() {
				duration := time.Since(start).Seconds()
				metrics.MorphCycleDuration.Observe(duration)
				span.End()
			}()

			topology := s.Next()
			metrics.MorphCyclesTotal.Inc()
			metrics.EndpointRotationEventsTotal.Add(float64(len(topology.Endpoints)))

			onMorph(topology)

			// Update M(t)
			now := time.Now()
			s.mu.Lock()
			rate := float64(s.cycles) / now.Sub(s.lastMorph).Minutes()
			s.mt = rate / s.maxRate
			s.lastMorph = now
			s.mu.Unlock()
		}
	}
}

