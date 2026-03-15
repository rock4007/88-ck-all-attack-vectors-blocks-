# 88/CK Architecture

The immune layer is split into three pillars and one stability supervisor:

- Pillar 1 Morphic: control-plane policy adaptation and xDS emission.
- Pillar 2 Consensus: Byzantine-tolerant agreement and PQ-safe attestation.
- Pillar 3 Entropy: anomaly scoring across embeddings and graph structure.
- Stability Engine: Lyapunov-anchored guardrails and incident orchestration.

Data flows from detection and consensus into a shared control state. Γ coupling captures cross-pillar feedback and is constrained to preserve S(t) monotonic recovery after perturbations.
