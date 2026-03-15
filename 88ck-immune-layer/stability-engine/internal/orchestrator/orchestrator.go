package orchestrator

import "github.com/[username]/88ck-immune-layer/stability-engine/internal/stability"

type Plan struct {
Action string
}

type Controller struct{}

func NewController() *Controller {
return &Controller{}
}

func (c *Controller) Plan(state stability.State) Plan {
if state.Status == "stable" {
return Plan{Action: "hold"}
}
return Plan{Action: "isolate-and-recover"}
}
