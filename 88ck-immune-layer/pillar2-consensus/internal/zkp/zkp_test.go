package zkp

import "testing"

func TestZKProofRoundTrip(t *testing.T) {
	prover, err := NewProver("svc-consensus")
	if err != nil {
		t.Fatalf("new prover: %v", err)
	}

	verifier := NewVerifier()
	verifier.RegisterIdentity("svc-consensus", prover.PublicKey())

	proof := prover.BuildProof("checkpoint-epoch-777", "nonce-1")
	if !verifier.Verify("checkpoint-epoch-777", proof) {
		t.Fatalf("expected proof verification to pass")
	}
}

func TestZKProofRejectsTamper(t *testing.T) {
	prover, err := NewProver("svc-consensus")
	if err != nil {
		t.Fatalf("new prover: %v", err)
	}

	verifier := NewVerifier()
	verifier.RegisterIdentity("svc-consensus", prover.PublicKey())

	proof := prover.BuildProof("checkpoint-epoch-777", "nonce-1")
	if verifier.Verify("checkpoint-epoch-778", proof) {
		t.Fatalf("expected verification failure for tampered statement")
	}
}
