# 88/CK All Attack Vectors Blocks

Defense-first distributed systems stack with adaptive filtering, zero-knowledge consensus, anomaly detection, and stability guardrails.

## Repository Overview

This repository contains the full 88/CK platform as a monorepo. The core implementation is in [88ck-immune-layer](88ck-immune-layer/README.md), with four cooperating runtime domains:

1. `pillar1-morphic`: ingress gateway and adaptive security filter
2. `pillar2-consensus`: replay-safe and zero-knowledge oriented admission and consensus security
3. `pillar3-entropy`: anomaly detection and explainability pipeline (Python)
4. `stability-engine`: Lyapunov-inspired rollout guardrail and orchestrator

## Project Structure

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

## Quick Start

### Prerequisites

- Go 1.22+
- Python 3.11+
- Node.js 20+
- Docker + Docker Compose

### Bootstrap

```bash
git clone https://github.com/rock4007/88-ck-all-attack-vectors-blocks-.git
cd 88-ck-all-attack-vectors-blocks-/88ck-immune-layer
./scripts/bootstrap.sh
```

### Run with Docker Compose

```bash
cd infra
docker compose up --build
```

## Development Commands

Run these from `88ck-immune-layer` unless noted otherwise.

```bash
# Go tests for Pillar 2
cd pillar2-consensus && go test ./...

# Go tests for stability engine
cd ../stability-engine && go test ./...

# Security filter tests
cd ../pillar1-morphic && go test ./internal/securityfilter/...

# Adversarial harness
cd ../adversarial-harness && python runner.py --strict
```

## Documentation

Technical docs live in `88ck-immune-layer/docs/`:

- `architecture.md`
- `api.md`
- `theory.md`
- `coupling-problem.md`

For full implementation details and service-level information, start with:

- [88ck-immune-layer/README.md](88ck-immune-layer/README.md)

## License

MIT. See [88ck-immune-layer/LICENSE](88ck-immune-layer/LICENSE).