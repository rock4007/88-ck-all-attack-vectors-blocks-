package main

import (
"log"

"github.com/[username]/88ck-immune-layer/stability-engine/internal/incident"
"github.com/[username]/88ck-immune-layer/stability-engine/internal/lyapunov"
"github.com/[username]/88ck-immune-layer/stability-engine/internal/orchestrator"
"github.com/[username]/88ck-immune-layer/stability-engine/internal/stability"
)

func main() {
l := lyapunov.NewConstraint(0.82)
s := stability.NewMonitor(l)
o := orchestrator.NewController()
i := incident.NewSink()

state := s.Snapshot(0.77)
plan := o.Plan(state)
i.Emit(plan)

log.Printf("stability-engine state=%s action=%s", state.Status, plan.Action)
}
