package zkp

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
)

// Proof is a non-interactive proof-of-possession transcript.
// The prover demonstrates knowledge of a private key without revealing it.
type Proof struct {
	IdentityID    string `json:"identity_id"`
	Nonce         string `json:"nonce"`
	StatementHash string `json:"statement_hash"`
	Signature     string `json:"signature"`
}

type Prover struct {
	identityID string
	privateKey ed25519.PrivateKey
	publicKey  ed25519.PublicKey
}

func NewProver(identityID string) (*Prover, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	return &Prover{identityID: identityID, privateKey: privateKey, publicKey: publicKey}, nil
}

func (p *Prover) PublicKey() ed25519.PublicKey {
	return p.publicKey
}

func (p *Prover) BuildProof(statement string, nonce string) Proof {
	// Bind signature to the statement hash plus nonce so proofs cannot be replayed
	// for a different proposal payload.
	statementHash := hashHex(statement)
	msg := transcriptMessage(statementHash, nonce)
	sig := ed25519.Sign(p.privateKey, msg)
	return Proof{
		IdentityID:    p.identityID,
		Nonce:         nonce,
		StatementHash: statementHash,
		Signature:     base64.StdEncoding.EncodeToString(sig),
	}
}

type Verifier struct {
	keys map[string]ed25519.PublicKey
}

func NewVerifier() *Verifier {
	return &Verifier{keys: make(map[string]ed25519.PublicKey)}
}

func (v *Verifier) RegisterIdentity(identityID string, key ed25519.PublicKey) {
	v.keys[identityID] = key
}

func (v *Verifier) Verify(statement string, proof Proof) bool {
	pub, ok := v.keys[proof.IdentityID]
	if !ok {
		return false
	}

	// Compare the claimed statement hash in constant time before signature check.
	expectedHash := hashHex(statement)
	if subtle.ConstantTimeCompare([]byte(proof.StatementHash), []byte(expectedHash)) != 1 {
		return false
	}

	sig, err := base64.StdEncoding.DecodeString(proof.Signature)
	if err != nil {
		return false
	}

	msg := transcriptMessage(proof.StatementHash, proof.Nonce)
	return ed25519.Verify(pub, msg, sig)
}

func hashHex(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}

func transcriptMessage(statementHash string, nonce string) []byte {
	// Keep transcript compact and deterministic for reproducible verification.
	sum := sha256.Sum256([]byte(statementHash + "|" + nonce))
	return sum[:]
}
