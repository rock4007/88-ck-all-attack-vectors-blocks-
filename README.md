# 88/CK All Attack Vectors Blocks

Security-focused distributed systems project that combines ingress protection, consensus admission hardening, anomaly detection, and rollout guardrails.

## Change Authority

This repository is owner-controlled. See [AUTHORITY_POLICY.md](AUTHORITY_POLICY.md) for approval and branch-protection policy.

## What's New (April 2026)

- Production hardening completed for Docker Compose runtime:
  - Healthchecks added for morphic, stability-engine, and frontend containers.
  - Restart policies, resource constraints, and log rotation configured in compose.
- Supply-chain and build hygiene improvements:
  - `.dockerignore` added for all service directories.
  - Root `.gitignore` added for generated and local-only artifacts.
- Frontend dependency security refresh:
  - Vite upgraded to `6.4.1`.
  - `npm audit` now reports 0 vulnerabilities.
- Deployment readiness validated end-to-end:
  - `docker compose up --build` completes successfully.
  - Runtime services report healthy state where expected.

## What This Project Demonstrates

- Secure backend engineering in Go and Python
- AppSec controls for SQLi and unsafe payload handling
- Replay-resistant and staged admission checks in consensus paths
- Containerized operations with Compose and Helm
- Security observability and incident-oriented design

## System Overview

The platform is implemented in [88ck-immune-layer/README.md](88ck-immune-layer/README.md) and organized into four runtime domains:

1. `pillar1-morphic`: ingress gateway and adaptive security filtering
2. `pillar2-consensus`: replay-safe and proof-oriented admission
3. `pillar3-entropy`: anomaly detection and explainability pipeline
4. `stability-engine`: Lyapunov-inspired rollout safety and orchestration

## Quick Start

### Prerequisites

- Go 1.25+
- Python 3.11+
- Node.js 20+
- Docker + Docker Compose

### Bootstrap

```bash
git clone https://github.com/rock4007/88-ck-all-attack-vectors-blocks-.git
cd 88-ck-all-attack-vectors-blocks-/88ck-immune-layer
./scripts/bootstrap.sh
```

### Run Full Stack

```bash
cd infra
docker compose up --build
```

## Validation Commands

Run from `88ck-immune-layer` unless noted.

```bash
# Go tests (example modules)
cd pillar1-morphic && go test ./...
cd ../pillar2-consensus && go test ./...
cd ../stability-engine && go test ./...

# Python harness
cd ../adversarial-harness && python runner.py --strict

# Frontend build
cd ../frontend && npm run build
```

## Deployment Paths

- Local and pre-production: `infra/docker-compose.yml`
- Kubernetes: Helm chart in `infra/helm/88ck`
- CI/CD and release automation: root `.github/workflows/`

## Repository Structure

```text
88-ck-all-attack-vectors-blocks-/
├── README.md
└── 88ck-immune-layer/
    ├── go.work
    ├── adversarial-harness/
    ├── docs/
    ├── frontend/
    ├── infra/
    ├── pillar1-morphic/
    ├── pillar2-consensus/
    ├── pillar3-entropy/
    ├── scripts/
    └── stability-engine/
```

## Role Fit

This project aligns well with:

- Security Engineer
- Application Security Engineer
- Platform Security Engineer
- Detection Engineer / SOC Engineer
- Cloud Security Engineer

## Documentation

Technical docs are in `88ck-immune-layer/docs/`:

- `architecture.md`
- `api.md`
- `theory.md`
- `coupling-problem.md`

## License

MIT. See [88ck-immune-layer/LICENSE](88ck-immune-layer/LICENSE).
