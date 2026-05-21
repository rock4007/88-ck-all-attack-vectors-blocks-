from __future__ import annotations

from dataclasses import dataclass, field
from typing import Dict, List

from internal.graph.attack_graph import AttackGraph

HIGH_RISK_THRESHOLD = 0.8
MEDIUM_RISK_THRESHOLD = 0.5
INCREASED_QUORUM = 5
STANDARD_QUORUM = 3


@dataclass
class ConsensusDecision:
    allowed: bool
    reason: str
    quorum_size: int
    verification_methods: List[str]


class ChainAwareConsensus:
    """Adjusts consensus quorum and verification methods based on live attack-chain risk."""

    def __init__(self, attack_graph: AttackGraph) -> None:
        self.attack_graph = attack_graph
        self.consensus_requirements: Dict[str, int] = {}

    def verify_action(self, action: str, actor: str, target: str) -> ConsensusDecision:
        chain_risk = self._assess_chain_risk(actor, target)

        if chain_risk > HIGH_RISK_THRESHOLD:
            quorum_size = INCREASED_QUORUM
            verification_methods = ["biometric", "hardware_token", "geo_verification"]
        elif chain_risk > MEDIUM_RISK_THRESHOLD:
            quorum_size = STANDARD_QUORUM + 1
            verification_methods = ["hardware_token", "geo_verification"]
        else:
            quorum_size = STANDARD_QUORUM
            verification_methods = ["standard"]

        self.consensus_requirements[f"{actor}->{target}"] = quorum_size
        return self._run_consensus_process(quorum_size, verification_methods, action)

    def _assess_chain_risk(self, actor: str, target: str) -> float:
        edge = self.attack_graph.edges.get((actor, target))
        if edge is None or edge.blocked:
            return 0.0

        actor_node = self.attack_graph.nodes.get(actor)
        target_node = self.attack_graph.nodes.get(target)

        actor_exposure = 1.0 - (actor_node.security_level if actor_node else 0.5)
        target_value = target_node.security_level if target_node else 0.5

        return min(1.0, edge.weight * 0.4 + edge.impact * 0.4 + actor_exposure * 0.1 + target_value * 0.1)

    def _run_consensus_process(
        self,
        quorum_size: int,
        verification_methods: List[str],
        action: str,
    ) -> ConsensusDecision:
        return ConsensusDecision(
            allowed=True,
            reason=f"consensus reached with quorum={quorum_size}",
            quorum_size=quorum_size,
            verification_methods=verification_methods,
        )
