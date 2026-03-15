package incident

import (
"log"

"github.com/[username]/88ck-immune-layer/stability-engine/internal/orchestrator"
)

type Sink struct{}

func NewSink() *Sink {
return &Sink{}
}

func (s *Sink) Emit(plan orchestrator.Plan) {
log.Printf("incident action=%s", plan.Action)
}
