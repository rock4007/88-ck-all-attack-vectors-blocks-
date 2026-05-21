from __future__ import annotations

from collections import deque
from dataclasses import dataclass
from typing import Any, Dict, List, Optional

THRESHOLD = 0.6
IMPACT_THRESHOLD = 0.4

@dataclass
class Alert:
    id: str
    source: str
    severity: float
    description: str
    metadata: Dict[str, Any]

@dataclass
class AttackEdge:
    source: str
    target: str
    technique_id: str
    weight: float
    impact: float

@dataclass
class AttackPath:
    nodes: List[str]
    edges: List[AttackEdge]

@dataclass
class PotentialChain:
    path: AttackPath
    probability: float
    impact: float
    recommended_action: str


def map_alert_to_technique(alert: Alert) -> str:
    # Explicit technique in metadata always wins over heuristic description matching.
    if "technique_id" in alert.metadata:
        return alert.metadata["technique_id"]
    mapping = {
        "sql injection": "T1190",
        "replay": "T1078",
        "privilege escalation": "T1068",
        "webshell": "T1505",
        "dos": "T1499",
        "suspicious query": "T1059",
    }
    for keyword, technique in mapping.items():
        if keyword in alert.description.lower():
            return technique
    return "T1059"


def find_attack_paths(graph: Dict[str, List[AttackEdge]], source: str, technique_id: str) -> List[AttackPath]:
    """Find all paths from source to critical assets.

    technique_id gates only the first hop (matching the alert vector); subsequent
    hops follow any edge because attackers pivot using different techniques.
    """
    paths: List[AttackPath] = []
    visited: set = set()
    q: deque = deque([AttackPath(nodes=[source], edges=[])])

    while q:
        current = q.popleft()
        last = current.nodes[-1]
        if last.startswith("critical:"):
            paths.append(current)
            continue

        if len(current.nodes) > 8:
            continue

        for edge in graph.get(last, []):
            # Only apply technique filter on the first hop from the alert source.
            if last == source and edge.technique_id != technique_id:
                continue
            if edge.target in current.nodes:
                continue

            next_path = AttackPath(nodes=current.nodes + [edge.target], edges=current.edges + [edge])
            signature = tuple(next_path.nodes)
            if signature in visited:
                continue
            visited.add(signature)
            q.append(next_path)

    return paths


def calculate_chain_probability(path: AttackPath) -> float:
    base = 0.5
    length_factor = 1.0 - 0.1 * len(path.edges)
    edge_weight = sum(edge.weight for edge in path.edges) / max(len(path.edges), 1)
    probability = base + 0.5 * edge_weight + 0.2 * length_factor
    return min(1.0, max(0.0, probability))


def calculate_chain_impact(path: AttackPath) -> float:
    node_impact = 0.0
    for node in path.nodes:
        if node.startswith("critical:"):
            node_impact = max(node_impact, 0.8)
        elif node.startswith("db:"):
            node_impact = max(node_impact, 0.6)
        elif node.startswith("svc:"):
            node_impact = max(node_impact, 0.5)

    edge_impact = max((edge.impact for edge in path.edges), default=0.0)
    impact = 0.4 * node_impact + 0.6 * edge_impact
    return min(1.0, impact)


def calculate_mitigation(path: AttackPath) -> str:
    if any(node.startswith("critical:") for node in path.nodes):
        return "isolate critical asset and enforce network segmentation"
    if any(edge.technique_id == "T1059" for edge in path.edges):
        return "harden command execution controls and limit process spawn privileges"
    return "apply targeted firewall rules, update ACLs, and increase monitoring on path nodes"


def detect_attack_chains(graph: Dict[str, List[AttackEdge]], current_alerts: List[Alert]) -> List[PotentialChain]:
    potential_chains: List[PotentialChain] = []

    for alert in current_alerts:
        technique_id = map_alert_to_technique(alert)
        paths = find_attack_paths(graph, alert.source, technique_id)

        for path in paths:
            probability = calculate_chain_probability(path)
            impact = calculate_chain_impact(path)

            if probability > THRESHOLD and impact > IMPACT_THRESHOLD:
                potential_chains.append(
                    PotentialChain(
                        path=path,
                        probability=probability,
                        impact=impact,
                        recommended_action=calculate_mitigation(path),
                    )
                )

    return potential_chains
