# 88/CK Immune Layer

[![S(t) live](https://img.shields.io/endpoint?url=https%3A%2F%2Fstatus.example.com%2F88ck%2Fstability.json)](https://status.example.com/88ck)
[![Build](https://img.shields.io/github/actions/workflow/status/rock4007/88-ck-all-attack-vectors-blocks-/ci.yml?branch=main)](./.github/workflows/ci.yml)
[![Coverage](https://img.shields.io/badge/coverage-enforced-brightgreen)](./.github/workflows/ci.yml)
[![arXiv](https://img.shields.io/badge/arXiv-2401.88088-b31b1b.svg)](https://arxiv.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

88/CK Immune Layer is a multi-pillar defensive monorepo combining adaptive policy control, consensus integrity, graph/embedding anomaly detection, and Lyapunov-constrained stability orchestration.

## Repository layout
- `pillar1-morphic`: adaptive scheduler and xDS gateway.
- `pillar2-consensus`: HotStuff-inspired control plane with PQ-safe attestation abstraction.
- `pillar3-entropy`: anomaly detection, graph scoring, and explainability components.
- `stability-engine`: system-level stability and incident control loop.
- `frontend`: operational dashboard.
- `infra`: Helm, compose, and Prometheus artifacts.

## Quick start
```bash
cd 88ck-immune-layer
./scripts/bootstrap.sh
```

Run services locally:
```bash
cd infra
docker compose up --build
```

Run adversarial scenarios:
```bash
cd adversarial-harness
python runner.py --strict
```

## Security baseline
- Distroless runtime images for all production Dockerfiles.
- No `math/rand` usage in Go code; entropy comes from `crypto/rand`.
- PQ-safe posture preserved through consensus attestation paths.
