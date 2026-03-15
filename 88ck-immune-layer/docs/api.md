# API Surface

## Morphic Gateway
- `GET /healthz` -> liveness status.
- `GET /tick` -> produce one policy scheduling cycle.

## Stability Engine (planned)
- `POST /evaluate` -> evaluate candidate coupling update.
- `POST /incident` -> execute containment playbook.

## Entropy Detector (planned)
- `POST /score` -> return baseline + graph anomaly decomposition.
