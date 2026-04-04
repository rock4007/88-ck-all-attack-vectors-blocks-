# 88/CK All Attack Vectors Blocks

Security-focused distributed systems project with adaptive filtering, consensus hardening, anomaly detection, and rollout guardrails.

## Change Authority

This repository is owner-controlled. See [AUTHORITY_POLICY.md](AUTHORITY_POLICY.md) for the exact approval and branch-protection rules.

## Hiring Snapshot

This repository can be used as a cybersecurity engineering portfolio. It demonstrates:

- Secure backend engineering in Go and Python
- Practical AppSec controls (injection blocking, payload defusing, replay resistance)
- Distributed system security patterns (admission gates, consensus hardening)
- Cloud-native operations (Docker, Compose, Helm, Prometheus)
- Security observability and incident-oriented design

Relevant roles:

- Security Engineer
- Application Security Engineer
- Platform Security Engineer
- Detection Engineer / SOC Engineer
- Cloud Security Engineer

## Repository Overview

This repository contains the full 88/CK platform as a monorepo. The core implementation is in [88ck-immune-layer](88ck-immune-layer/README.md), with four cooperating runtime domains:

1. pillar1-morphic: ingress gateway and adaptive security filter
2. pillar2-consensus: replay-safe and zero-knowledge oriented admission and consensus security
3. pillar3-entropy: anomaly detection and explainability pipeline (Python)
4. stability-engine: Lyapunov-inspired rollout guardrail and orchestrator

## One-by-One Breakdown

### 1. What this project is

88/CK is a cybersecurity-focused distributed systems project. The goal is to block attacks early, protect consensus traffic, detect abnormal behavior, and stop risky changes before they impact production.

### 2. Pillar 1: Morphic (gateway security)

This is the first control point for incoming traffic. Requests are inspected for suspicious patterns such as SQL injection and dangerous payload signatures. Unsafe requests are blocked before they reach internal services.

### 3. Pillar 2: Consensus (secure admission and verification)

This layer hardens distributed coordination. It applies staged admission checks, including replay protection and proof verification, so duplicated or untrusted proposals are rejected.

### 4. Pillar 3: Entropy (anomaly detection)

This component handles behavior-based detection. It uses scoring and explainability-oriented modules to identify activity that deviates from expected system behavior.

### 5. Stability Engine (rollout safety)

This service evaluates rollout risk before a change proceeds. If predicted stability risk is too high, the system can recommend holding, freezing, or isolating rollout activity.

### 6. Adversarial harness

The harness runs adversarial scenarios to validate whether controls hold under pressure. It helps verify that the system works not only in normal operation but also against realistic attack paths.

### 7. Infra and operations

Docker, Compose, Helm, and Prometheus assets support deployment and observability. This makes the platform easier to run in development and easier to monitor in test environments.

### 8. Frontend

The frontend provides an operator-facing interface for platform workflows and visibility, complementing the service-side APIs.

### 9. Why this matters for hiring

The repository demonstrates practical security engineering work across secure coding, detection, distributed systems hardening, testing, and operational telemetry.

### 10. Role fit

This project aligns well with:

- Security Engineer
- Application Security Engineer
- Platform Security Engineer
- Detection Engineer / SOC Engineer
- Cloud Security Engineer

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

### Run with Docker Compose

```bash
cd infra
docker compose up --build
```

## Development Commands

Run these from 88ck-immune-layer unless noted otherwise.

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

## Repository Layout

```text
88ck-immune-layer/
├── pillar1-morphic/          # Gateway, SQLi/malware filter, gamma coupling, xDS
│   ├── cmd/gateway/          # HTTP entrypoint (port 8080)
│   └── internal/
│       ├── securityfilter/   # Ingress threat detection + defuser
│       ├── metrics/          # Prometheus + OTel declarations
│       ├── gamma/            # Lyapunov coupling controller
│       ├── scheduler/        # Bounded score scheduler
│       └── xds/              # In-process xDS publisher stub
│
├── pillar2-consensus/        # ZK agreement, replay guard, PQ attestation
│   ├── cmd/consensus/        # Consensus node entrypoint
│   └── internal/
│       ├── zkp/              # Ed25519 + SHA-256 proof-of-possession
│       └── security/         # 3-tier admission gate
│
├── pillar3-entropy/          # Python anomaly detection
│   ├── cmd/detector/         # Detector entrypoint
│   └── internal/
│       ├── baseline/         # Baseline tracker
│       ├── embedding/        # Embedding-based scoring
│       ├── explainability/   # SHAP layer
│       └── graph/            # WL-style graph scorer
│
├── stability-engine/         # Lyapunov guardrail + orchestrator
│   └── cmd/engine/           # HTTP API (port 8090)
│
├── frontend/                 # React + Vite + Tailwind ops UI
│
├── adversarial-harness/      # MITRE-style scenario runner
│
├── infra/
│   ├── docker-compose.yml
│   ├── helm/                 # Kubernetes Helm chart
│   └── prometheus/           # Alert + recording rules
│
├── docs/                     # Architecture, theory, API docs
└── scripts/                  # Bootstrap and validation utilities
```

---

## CI / CD Pipelines

| Workflow | Trigger | Purpose |
|---|---|---|
| `ci.yml` | Push / PR | Build, test, coverage enforce |
| `security-scan.yml` | Push / PR | Dependency and secrets scan |
| `adversarial-test.yml` | Schedule + PR | MITRE scenario regression |
| `release.yml` | Tag `v*` | Multi-arch image build + push |

---

## Current Scope Notes

- This is an engineering project and learning platform, not a turnkey commercial product.
- Some controls are simplified prototypes intended to show design approach and integration patterns.
- Use the docs and tests in each component to understand current behavior and limitations.

## Documentation

Technical docs live in `88ck-immune-layer/docs/`:

- architecture.md
- api.md
- theory.md
- coupling-problem.md

For full implementation details and service-level information, start with:

- [88ck-immune-layer/README.md](88ck-immune-layer/README.md)

## License

MIT. See [88ck-immune-layer/LICENSE](88ck-immune-layer/LICENSE).