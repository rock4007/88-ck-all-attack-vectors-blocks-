package security

import (
	"sync"

	"github.com/88ck/pillar2-consensus/internal/nhi"
	"github.com/88ck/pillar2-consensus/internal/zkp"
)

type AdmissionRequest struct {
	Statement   string
	Attestation string
	Proof       zkp.Proof
}

type AdmissionDecision struct {
	Allowed bool
	Reason  string
}

type ReplayGuard struct {
	mu   sync.Mutex
	seen map[string]struct{}
}

func NewReplayGuard() *ReplayGuard {
	return &ReplayGuard{seen: make(map[string]struct{})}
}

func (r *ReplayGuard) TryMark(nonce string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.seen[nonce]; ok {
		return false
	}
	r.seen[nonce] = struct{}{}
	return true
}

type Layer struct {
	identity *nhi.IdentityProvider
	zk       *zkp.Verifier
	replay   *ReplayGuard
}

func NewLayer(identity *nhi.IdentityProvider, zk *zkp.Verifier, replay *ReplayGuard) *Layer {
	return &Layer{identity: identity, zk: zk, replay: replay}
}

// Evaluate applies admission checks in strict order:
// 1) nonce freshness, 2) PQ attestation shape, 3) proof validity.
// This ordering keeps cheap rejection paths fast under attack traffic.
func (l *Layer) Evaluate(req AdmissionRequest) AdmissionDecision {
	if req.Proof.Nonce == "" {
		return AdmissionDecision{Allowed: false, Reason: "nonce_required"}
	}
	if !l.replay.TryMark(req.Proof.Nonce) {
		return AdmissionDecision{Allowed: false, Reason: "replay_detected"}
	}
	if !l.identity.Verify(req.Attestation) {
		return AdmissionDecision{Allowed: false, Reason: "pq_attestation_failed"}
	}
	if !l.zk.Verify(req.Statement, req.Proof) {
		return AdmissionDecision{Allowed: false, Reason: "zk_proof_invalid"}
	}
	return AdmissionDecision{Allowed: true, Reason: "admitted"}
}
