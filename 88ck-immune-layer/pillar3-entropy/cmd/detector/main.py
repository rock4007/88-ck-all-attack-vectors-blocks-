from internal.baseline.baseline import BaselineEstimator
from internal.embedding.embed import EmbeddingModel
from internal.graph.wl_scorer import WeisfeilerLehmanScorer
from internal.explainability.shap_explainer import ShapExplainer


def main() -> None:
    baseline = BaselineEstimator(window=128)
    embedder = EmbeddingModel(dim=64)
    scorer = WeisfeilerLehmanScorer(depth=2)
    explainer = ShapExplainer()

    event = {"actor": "nhi:svc-gateway", "target": "db:ledger", "action": "query"}
    vector = embedder.encode(event)
    baseline_score = baseline.score(vector)
    graph_score = scorer.score(event)
    explanation = explainer.explain(event, baseline_score + graph_score)

    print(
        {
            "baseline": round(baseline_score, 6),
            "graph": round(graph_score, 6),
            "total": round(baseline_score + graph_score, 6),
            "explanation": explanation,
        }
    )


if __name__ == "__main__":
    main()
