<div align="center">

<img src="https://img.shields.io/badge/88%2FCK-Immune%20Layer-0d1117?style=for-the-badge&labelColor=0d1117&color=00d4ff" alt="88/CK Immune Layer"/>

# 88/CK Immune Layer

**Autonomous defense infrastructure for distributed systems at scale.**

[![Build](https://img.shields.io/github/actions/workflow/status/rock4007/88-ck-all-attack-vectors-blocks-/ci.yml?branch=main&style=flat-square&label=CI)](https://github.com/rock4007/88-ck-all-attack-vectors-blocks-/actions)
[![Security Scan](https://img.shields.io/github/actions/workflow/status/rock4007/88-ck-all-attack-vectors-blocks-/security-scan.yml?branch=main&style=flat-square&label=Security%20Scan&color=green)](https://github.com/rock4007/88-ck-all-attack-vectors-blocks-/actions)
[![Go Version](https://img.shields.io/badge/Go-1.22%2B-00ADD8?style=flat-square&logo=go)](https://go.dev)
[![Python](https://img.shields.io/badge/Python-3.11-3776AB?style=flat-square&logo=python)](https://python.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow?style=flat-square)](./LICENSE)
[![Coverage](https://img.shields.io/badge/Coverage-Enforced-brightgreen?style=flat-square)](./github/workflows/ci.yml)

<br/>

> *Stop attacks before they happen. Block instability before it spreads. Defend every layer — automatically.*

</div>

---

## What is 88/CK Immune Layer?

88/CK Immune Layer is a **production-grade, multi-pillar security and resilience platform** designed to protect distributed systems from the ground up. It combines:

- 🧠 **Adaptive policy control** that morphs to attack patterns in real time
- 🔐 **Zero-knowledge consensus** so nodes agree without revealing state
- 🕵️ **AI-powered anomaly detection** with graph and embedding analysis
- ⚖️ **Lyapunov-constrained orchestration** that mathematically prevents destabilizing changes
- 🛡️ **Ingress threat neutralization** — SQL injection, malware payloads, command injection — blocked and defused at the gate

---

## Architecture at a Glance

```
                        ┌─────────────────────────────────────────────┐
                        │              88/CK Immune Layer              │
                        └──────────────────┬──────────────────────────┘
                                           │
            ┌──────────────────────────────┼──────────────────────────────┐
            │                             │                              │
    ┌───────▼────────┐           ┌────────▼───────┐            ┌────────▼───────┐
    │  Pillar 1      │           │  Pillar 2      │            │  Pillar 3      │
    │  Morphic       │           │  Consensus     │            │  Entropy       │
    │  ─────────     │           │  ─────────     │            │  ─────────     │
    │  Gateway       │           │  ZK Proofs     │            │  Graph Scorer  │
    │  SQLi Blocker  │           │  Replay Guard  │            │  Embedding     │
    │  xDS Policy    │           │  PQ Attest.    │            │  SHAP Explain  │
    │  Gamma Control │           │  3-Tier Gate   │            │  Anomaly Score │
    └───────┬────────┘           └────────┬───────┘            └────────┬───────┘
            │                             │                              │
            └──────────────────┬──────────┘──────────────────────────────
                               │
                    ┌──────────▼──────────┐
                    │   Stability Engine  │
                    │   ───────────────── │
                    │   Lyapunov Control  │
                    │   Guardrail API     │
                    │   Incident Sink     │
                    └─────────────────────┘
```

---

## Four Pillars, One Defense

### 🔷 Pillar 1 — Morphic (Adaptive Gateway)
The ingress guardian. Every request passes through a **multi-layer security filter** before reaching any service.

| Capability | Detail |
|---|---|
| **SQLi Blocker** | 8 regex patterns covering UNION, DROP, comment injection, error-based exfiltration |
| **Malware Defuser** | 11 signature patterns; dangerous bytes stripped before logging |
| **xDS Policy Engine** | Dynamic control-plane configuration pushed to Envoy proxies |
| **Gamma Coupling** | Rate-limits cross-pillar influence to prevent destabilization cascades |
| **Prometheus Metrics** | Per-reason security block counter exposed at `/metrics` |

### 🔷 Pillar 2 — Consensus (Zero-Knowledge Agreement)
Nodes reach agreement **without revealing proposal contents** — even under adversarial observation.

| Capability | Detail |
|---|---|
| **ZK Proof-of-Possession** | Ed25519 + SHA-256 transcript binding (NIZK-style) |
| **Replay Guard** | Mutex-protected nonce map prevents replay attacks |
| **PQ Attestation** | Post-quantum-safe attestation abstraction |
| **3-Tier Admission** | Nonce → PQ attestation → ZK proof, cheapest checks first |

### 🔷 Pillar 3 — Entropy (AI Anomaly Detection)
Detects what rules miss — using **graph neural patterns and embedding distance** to catch novel threats.

| Capability | Detail |
|---|---|
| **Graph Scoring** | Weisfeiler-Leman style structural comparison |
| **Embedding Detector** | ONNX-backed vector distance anomaly scoring |
| **Explainability** | SHAP value attribution for every anomaly decision |
| **Baseline Tracking** | Rolling normality windows with drift alerting |

### 🔷 Stability Engine (Lyapunov Control Gate)
The final checkpoint. **No change reaches production** without passing a stability prediction.

| Capability | Detail |
|---|---|
| **Predictive Evaluation** | `S(t+1) = S(t) - (γ·0.4) - (d·0.3)` |
| **Guardrail API** | POST `/evaluate` — approve or block a rollout proposal |
| **Risk Classification** | Critical / High / Medium / Low bands |
| **Orchestrator Plans** | `hold`, `stage-rollout`, `freeze-change-and-monitor`, `isolate-and-recover` |

---

## Security Stack

```
Request ──► SQLi/Malware Filter ──► ZK Admission Gate ──► Stability Guardrail ──► Service
              │                          │                        │
           Block + 403             Block + reason           Block + 409
           Defuse payload          Log nonce replay         Log risk band
           Increment metric        Verify ZK proof          Return rollout plan
```

Every blocked request is:
1. **Defused** — control characters and dangerous symbols stripped before any logging
2. **Traced** — SHA-256 derived trace ID attached to every deny response
3. **Measured** — `morphic_security_blocks_total{reason="..."}` incremented in Prometheus
4. **Alerted** — Prometheus rules fire if block rate exceeds operational thresholds

---

## Quick Start

**Prerequisites:** Go 1.22+, Python 3.11+, Docker, Node.js 20+

```bash
# Clone and bootstrap
git clone https://github.com/rock4007/88-ck-all-attack-vectors-blocks-
cd 88-ck-all-attack-vectors-blocks-/88ck-immune-layer
./scripts/bootstrap.sh
```

**Run all services with Docker Compose:**
```bash
cd infra
docker compose up --build
```

**Run the Stability Engine standalone:**
```bash
cd stability-engine
go run ./cmd/engine
# Listening on :8090
```

**Evaluate a rollout proposal:**
```bash
curl -sS -X POST http://localhost:8090/evaluate \
  -H 'Content-Type: application/json' \
  -d '{
    "proposal_id": "deploy-v2.1.0",
    "current_stability": 0.91,
    "gamma_delta": 0.04,
    "disturbance": 0.03
  }' | jq
```

**Test the security filter:**
```bash
# This should return 403
curl -i "http://localhost:8080/tick?q=1'+OR+'1'='1"

# This should return a 200
curl -i "http://localhost:8080/healthz"
```

---

## Guardrail API Reference

### `POST /evaluate`

**Request:**
```json
{
  "proposal_id": "string",
  "current_stability": 0.0–1.0,
  "gamma_delta": 0.0–1.0,
  "disturbance": 0.0–1.0
}
```

**Response (approved):** `200 OK`
```json
{
  "decision": {
    "approved": true,
    "reason": "stability nominal",
    "risk": "low",
    "predicted_stability": 0.88
  },
  "orchestrator": { "plan": "stage-rollout" }
}
```

**Response (blocked):** `409 Conflict`
```json
{
  "decision": {
    "approved": false,
    "reason": "predicted stability below minimum floor",
    "risk": "critical",
    "predicted_stability": 0.71
  },
  "orchestrator": { "plan": "freeze-change-and-monitor" }
}
```

### `GET /healthz`
Returns `200 ok` when the engine is live.

---

## Observability

**Prometheus metrics exposed:**

| Metric | Type | Description |
|---|---|---|
| `morphic_security_blocks_total` | Counter | Blocked requests by reason label |
| `morphic_security_block_rate_5m` | Recording Rule | 5-minute block rate |
| `morphic_security_block_rate_critical_5m` | Recording Rule | Critical payload rate |
| `stability_index` | Gauge | Current Lyapunov stability score |

**Active alerts:**

| Alert | Severity | Condition |
|---|---|---|
| `MorphicSecurityBlocksSpike` | warning | Block rate > 0.5/s for 5 minutes |
| `MorphicCriticalPayloadBlocks` | critical | Critical payload rate > 0.1/s for 3 minutes |
| `StabilityDegradation` | critical | `stability_index` < 0.80 for 2 minutes |

---

## Repository Layout

```
88ck-immune-layer/
├── pillar1-morphic/          # Gateway, SQLi/malware filter, gamma coupling, xDS
│   ├── cmd/gateway/          # HTTP entrypoint (port 8080)
│   └── internal/
│       ├── securityfilter/   # Ingress threat detection + defuser
│       ├── metrics/          # Prometheus + OTel declarations
│       ├── gamma/            # Lyapunov coupling controller
│       ├── scheduler/        # ChaCha20-secured adaptive scheduler
│       └── xds/              # Envoy control-plane publisher
│
├── pillar2-consensus/        # ZK agreement, replay guard, PQ attestation
│   ├── cmd/consensus/        # Consensus node entrypoint
│   └── internal/
│       ├── zkp/              # Ed25519 + SHA-256 proof-of-possession
│       └── security/         # 3-tier admission gate
│
├── pillar3-entropy/          # Python anomaly detection
│   └── src/
│       ├── detector.py       # Main detector entrypoint
│       ├── baseline.py       # Normality baseline tracker
│       ├── graph_scorer.py   # WL-style graph anomaly scorer
│       └── explainability.py # SHAP attribution layer
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

## Security Baseline

- **Distroless containers** — `gcr.io/distroless/static-debian12:nonroot` for all Go services
- **No `math/rand`** — all entropy sourced from `crypto/rand`
- **Post-quantum posture** — PQ-safe attestation abstraction in consensus path
- **Zero-knowledge proofs** — proposal content never exposed during agreement
- **Log injection prevention** — all logged payloads pass through the defuser before writing
- **Replay protection** — nonce map blocks duplicate proposal submissions

---

## Development

```bash
# Run all Go tests (pillar2 + stability-engine)
cd pillar2-consensus && go test ./...
cd stability-engine && go test ./...

# Run security filter tests
cd pillar1-morphic && go test ./internal/securityfilter/...

# Validate a coupling coefficient
./scripts/validate-coefficients.sh 0.42

# Run adversarial harness
cd adversarial-harness && python runner.py --strict
```

---

<div align="center">

**Built for systems that cannot afford to fail.**

[![MIT License](https://img.shields.io/badge/License-MIT-yellow?style=flat-square)](./LICENSE)

</div>
