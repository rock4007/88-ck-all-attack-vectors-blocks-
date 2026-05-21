from __future__ import annotations

from dataclasses import dataclass
from typing import Any, Dict, List, Optional

from internal.graph.attack_chain import AttackEdge, AttackPath, PotentialChain
from internal.graph.attack_graph import AttackGraph

PREDICTION_THRESHOLD = 0.5


def isolate_component(target: str) -> None:
    # In an enterprise deployment, this would sever network access to the target.
    print(f"[DISRUPTION] isolating component: {target}")


def deploy_decoy(target: str) -> None:
    # Deploy a decoy or honeypot to divert attacker activity.
    print(f"[DISRUPTION] deploying decoy for: {target}")


def trigger_morphic_rotation(target: str) -> None:
    # Trigger morphological hardening routines around the target.
    print(f"[DISRUPTION] morphic rotation triggered for: {target}")


def require_consensus_verification(target: str) -> None:
    # Enforce additional consensus checks before the target can be used.
    print(f"[DISRUPTION] consensus verification required for: {target}")


def update_attack_graph(attack_graph: AttackGraph, source: str, target: str, blocked: bool = True) -> None:
    if blocked:
        attack_graph.block_edge(source, target)
    if target in attack_graph.nodes:
        attack_graph.update_exposure(target, {"blocked": blocked})


def normalize_edge(edge: Any) -> AttackEdge:
    if isinstance(edge, AttackEdge):
        return edge
    return AttackEdge(
        source=edge.get("source", ""),
        target=edge.get("target", ""),
        technique_id=edge.get("technique_id", ""),
        weight=edge.get("weight", 0.0),
        impact=edge.get("impact", 0.0),
    )


def disrupt_attack_chain(
    chain: Any,
    disruption_method: str,
    attack_graph: Optional[AttackGraph] = None,
) -> None:
    path = getattr(chain, "path", None)
    if path is None and isinstance(chain, dict):
        path = chain.get("path", [])

    edges = path if isinstance(path, list) else getattr(path, "edges", [])
    if not isinstance(edges, list) and hasattr(edges, "__iter__"):
        edges = list(edges)

    for raw_edge in edges:
        edge = normalize_edge(raw_edge)
        target = edge.target

        if disruption_method == "ISOLATION":
            isolate_component(target)
        elif disruption_method == "DECEPTION":
            deploy_decoy(target)
        elif disruption_method == "MORPHIC":
            trigger_morphic_rotation(target)
        elif disruption_method == "CONSENSUS":
            require_consensus_verification(target)
        else:
            print(f"[DISRUPTION] unknown method {disruption_method}, defaulting to isolation for {target}")
            isolate_component(target)

        if attack_graph is not None:
            update_attack_graph(attack_graph, edge.source, target, blocked=True)


# ---------------------------------------------------------------------------
# Graph rewriting
# ---------------------------------------------------------------------------

def rewrite_attack_graph(graph: AttackGraph, active_chains: List[PotentialChain]) -> None:
    """Block critical edges in the attack graph for each active chain."""
    for chain in active_chains:
        path = chain.path
        raw_edges: List[AttackEdge] = path.edges if isinstance(path, AttackPath) else []
        critical = [e for e in raw_edges if e.impact > 0.5 or e.weight > 0.7]

        for edge in critical:
            if edge.technique_id in ("T1078", "T1550", "T1119"):
                graph.set_mitigation(edge.technique_id, "rotate credentials and enforce MFA")
            elif edge.technique_id in ("T1190", "T1059"):
                graph.set_mitigation(
                    edge.technique_id,
                    "harden command execution and apply least-privilege",
                )
            elif edge.technique_id in ("T1055", "T1573"):
                graph.set_mitigation(
                    edge.technique_id,
                    "enforce process integrity and audit injection vectors",
                )
            graph.block_edge(edge.source, edge.target)


# ---------------------------------------------------------------------------
# Predictive analysis
# ---------------------------------------------------------------------------

@dataclass
class PredictedChain:
    path: List[str]
    probability: float
    timeframe: str
    recommended_mitigation: str


def predict_future_chains(
    graph: AttackGraph,
    current_state: Dict[str, Any],
    threat_intel: Any,
) -> List[PredictedChain]:
    """Predict probable future attack chains from high-risk components."""
    future_chains: List[PredictedChain] = []
    high_risk = _identify_high_risk_components(graph, current_state)
    techniques: List[str] = getattr(threat_intel, "relevant_techniques", [])

    for component in high_risk:
        for technique in techniques:
            simulated = _simulate_attack_path(graph, component, technique)
            if simulated["probability"] > PREDICTION_THRESHOLD:
                future_chains.append(
                    PredictedChain(
                        path=simulated["path"],
                        probability=simulated["probability"],
                        timeframe=simulated["timeframe"],
                        recommended_mitigation=_calculate_proactive_mitigation(simulated),
                    )
                )

    return future_chains


def _identify_high_risk_components(
    graph: AttackGraph,
    current_state: Dict[str, Any],
) -> List[str]:
    high_risk: List[str] = []
    for node_id, node in graph.nodes.items():
        state_info = current_state.get(node_id, {})
        if (
            node.current_state == "compromised"
            or state_info.get("alert_count", 0) > 2
            or node.security_level >= 0.8
        ):
            high_risk.append(node_id)
    return high_risk


def _simulate_attack_path(
    graph: AttackGraph,
    component: str,
    technique: str,
) -> Dict[str, Any]:
    paths = [
        path
        for path in graph.get_paths(component, "critical:")
        if _path_starts_with_technique(graph, path, technique)
    ]
    if not paths:
        return {"probability": 0.0, "path": [], "timeframe": "N/A"}

    best = min(paths, key=len)
    hops = max(len(best) - 1, 1)
    total_weight = 0.0
    for i in range(len(best) - 1):
        edge = graph.edges.get((best[i], best[i + 1]))
        total_weight += edge.weight if edge else 0.5

    probability = min(1.0, total_weight / hops)
    return {
        "probability": probability,
        "path": best,
        "timeframe": f"{hops * 5}-{hops * 15} minutes",
    }


def _path_starts_with_technique(
    graph: AttackGraph,
    path: List[str],
    technique: str,
) -> bool:
    if not technique or len(path) < 2:
        return True
    edge = graph.edges.get((path[0], path[1]))
    return edge is not None and edge.technique_id == technique


def _calculate_proactive_mitigation(simulated_path: Dict[str, Any]) -> str:
    path: List[str] = simulated_path.get("path", [])
    if any("critical:" in node for node in path):
        return "enforce network segmentation and activate deception layer before critical assets"
    return "increase monitoring, apply rate limiting, and pre-position forensic agents"
