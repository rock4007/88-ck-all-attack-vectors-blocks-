package gamma

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"time"

	"github.com/redis/go-redis/v9"

	"github.com/88ck/pillar1-morphic/internal/metrics"
)

const (
	DefaultDeltaT = 200 * time.Millisecond
	TauEmb        = 38 * time.Millisecond
	SafetyMargin  = DefaultDeltaT - TauEmb
	RedisStream   = "88ck:gamma:baseline-update"
)

type GammaProtocol struct {
	deltaT      time.Duration
	tauEmb      time.Duration
	redisClient redis.Client
	metrics     *metrics.Metrics
}

func NewGammaProtocol(deltaT time.Duration, rdb redis.Client) *GammaProtocol {
	if err := (&GammaProtocol{
		deltaT:      deltaT,
		tauEmb:      TauEmb,
		redisClient: rdb,
		metrics:     metrics.NewMetrics(context.Background()),
	}).ValidateLeadTime(deltaT); err != nil {
		panic(err) // constructor validate
	}
	metrics.GammaLeadTimeMs.Set(float64(deltaT.Milliseconds()))
	metrics.GammaSafetyMarginMs.Set(float64(SafetyMargin.Milliseconds()))
	return &GammaProtocol{
		deltaT:      deltaT,
		tauEmb:      TauEmb,
		redisClient: rdb,
		metrics:     metrics.NewMetrics(context.Background()),
	}
}

func (g *GammaProtocol) ValidateLeadTime(deltaT time.Duration) error {
	if deltaT < g.tauEmb {
		metrics.GammaFaultsTotal.Inc()
		return errors.New("lead time violation: deltaT < tauEmb")
	}
	return nil
}

func (g *GammaProtocol) ScheduleBaselineUpdate(mutationID string, mutationTime time.Time) error {
_ = mutationTime.Add(-g.deltaT)
	if err := g.ValidateLeadTime(g.deltaT); err != nil {
		return err
	}

	payload := map[string]interface{}{
		"mutation_id":    mutationID,
		"scheduled_at":   mutationTime.Unix(),
		"delta_t_ms":     g.deltaT.Milliseconds(),
		"safety_margin":  SafetyMargin.Milliseconds(),
		"seed":           "", // hex from topology, stub ""
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal payload: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = g.redisClient.XAdd(ctx, &redis.XAddArgs{
		Stream: RedisStream,
		Values: map[string]interface{}{"event": string(data)},
	}).Result()
	return err
}

