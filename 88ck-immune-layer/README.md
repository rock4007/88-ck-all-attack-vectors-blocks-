# 88/CK Immune Layer

[![S(t) live](https://img.shields.io/endpoint?url=https%3A%2F%2Fstatus.example.com%2F88ck%2Fstability.json)](https://status.example.com/88ck)
[![Build](https://img.shields.io/github/actions/workflow/status/rock4007/88-ck-all-attack-vectors-blocks-/ci.yml?branch=main)](./.github/workflows/ci.yml)
[![Coverage](https://img.shields.io/badge/coverage-enforced-brightgreen)](./.github/workflows/ci.yml)
[![arXiv](https://img.shields.io/badge/arXiv-2401.88088-b31b1b.svg)](https://arxiv.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

88/CK Immune Layer is a resilience platform for autonomous defense in modern distributed systems. It combines adaptive policy control, consensus integrity, anomaly detection, and Lyapunov-constrained orchestration into one operational stack.

## Why It Matters
- Stops unsafe rollouts before they destabilize production.
- Connects detection, consensus, and control loops into measurable stability outcomes.
- Ships with security-first defaults: distroless runtimes, PQ-safe consensus posture, and adversarial validation paths.

## System Pillars
Three pillars and one supervisory engine power the architecture:
- Pillar 1 Morphic: adaptive scheduling and xDS policy propagation.
- Pillar 2 Consensus: HotStuff-inspired agreement with PQ-safe attestation abstraction.
- Pillar 3 Entropy: embedding and graph-based anomaly detection with explainability.
- Stability Engine: Lyapunov-constrained control loop that blocks destabilizing rollout proposals.

## Core capabilities
- Autonomous stability guardrail with predictive proposal evaluation.
- Gamma coupling controls to constrain cross-pillar destabilization.
- Distroless runtime images for all service containers.
- Security scanning and adversarial regression workflows in GitHub Actions.
- Helm and Docker Compose deployment assets for fast environment bring-up.

## Repository layout
- pillar1-morphic: gateway runtime, scheduler, gamma coupling, xDS publisher.
- pillar2-consensus: consensus proposal path, PQ signature abstraction, NHI verification.
- pillar3-entropy: detector entrypoint, baseline scoring, WL-style graph scorer, explainability.
- stability-engine: evaluation API, Lyapunov floor checks, orchestration and incident sink.
- frontend: React + Vite operations UI.
- adversarial-harness: MITRE-style scenario runner (private code path).
- infra: compose files, Helm chart, Prometheus alerts and recording rules.
- docs: architecture, theory, coupling problem statement, API summary.

## For Platform Teams
- Deploy each pillar independently, then couple them progressively through guarded rollout policies.
- Use Stability Engine as the final control gate before production changes.
- Wire Prometheus alerts on S(t) and gamma coupling drift to enforce operational guardrails.

## Quick start
```bash
cd 88ck-immune-layer
./scripts/bootstrap.sh
```

Run core services locally:
```bash
cd infra
docker compose up --build
```

Run the stability engine directly:
```bash
cd stability-engine
go run ./cmd/engine
```

In another shell, evaluate a rollout proposal:
```bash
curl -sS -X POST http://localhost:8090/evaluate \
	-H 'Content-Type: application/json' \
	-d '{
		"proposal_id": "rollout-2026-03-15-a",
		"current_stability": 0.91,
		"gamma_delta": 0.04,
		"disturbance": 0.03
	}'
```

Run adversarial scenarios locally:
```bash
cd adversarial-harness
python runner.py --strict
```

## Stability Guardrail API
- GET /healthz
- POST /evaluate

Decision semantics:
- 200 OK: proposal approved and rollout may proceed in staged mode.
- 409 Conflict: proposal blocked due to stability-floor or gamma-step policy breach.

The response includes:
- current state snapshot.
- decision object with approved, reason, risk, predicted_stability, coupling_impact.
- orchestrator plan for hold, stage-rollout, freeze-change-and-monitor, or isolate-and-recover.

## Development workflow
Run Go tests for the stability engine:
```bash
cd stability-engine
go test ./...
```

Validate coupling coefficient format:
```bash
./scripts/validate-coefficients.sh 0.42
```

CI pipelines:
- .github/workflows/ci.yml
- .github/workflows/security-scan.yml
- .github/workflows/adversarial-test.yml
- .github/workflows/release.yml

## Security baseline
- Distroless runtime images for production containers.
- No math/rand usage in Go code; entropy comes from crypto/rand.
- PQ-safe posture preserved through consensus attestation paths.
- Private adversarial harness code is ignored by default to keep attack-playbook internals out of public commits.

## License
MIT. See LICENSE.
