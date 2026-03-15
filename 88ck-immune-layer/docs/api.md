# API Surface

## Morphic Gateway
- `GET /healthz` -> liveness status.
- `GET /tick` -> produce one policy scheduling cycle.

## Stability Engine
- `GET /healthz` -> liveness status.
- `POST /evaluate` -> evaluate candidate coupling update and return a guardrail decision.

Request body:

```json
{
	"proposal_id": "rollout-2026-03-15-a",
	"current_stability": 0.91,
	"gamma_delta": 0.04,
	"disturbance": 0.03
}
```

Response body:

```json
{
	"state": {
		"Status": "stable",
		"Value": 0.91
	},
	"decision": {
		"proposal_id": "rollout-2026-03-15-a",
		"approved": true,
		"action": "allow-rollout",
		"reason": "within_lyapunov_guardrails",
		"predicted_stability": 0.885,
		"risk": "low",
		"coupling_impact": 0.04
	},
	"plan": {
		"Action": "stage-rollout",
		"Reason": "within_lyapunov_guardrails"
	}
}
```

## Entropy Detector (planned)
- `POST /score` -> return baseline + graph anomaly decomposition.
