package main

import (
"log"
"time"

"github.com/[username]/88ck-immune-layer/pillar2-consensus/internal/dilithium"
"github.com/[username]/88ck-immune-layer/pillar2-consensus/internal/hotstuff"
"github.com/[username]/88ck-immune-layer/pillar2-consensus/internal/nhi"
)

func main() {
engine := hotstuff.NewEngine(4)
identity := nhi.NewIdentityProvider()
signer := dilithium.NewSigner()

proposal := engine.Propose("checkpoint-epoch-001")
attestation := signer.Sign(proposal)
verified := identity.Verify(attestation)

log.Printf("proposal=%s verified=%t quorum=%d", proposal, verified, engine.Quorum())
time.Sleep(100 * time.Millisecond)
}
