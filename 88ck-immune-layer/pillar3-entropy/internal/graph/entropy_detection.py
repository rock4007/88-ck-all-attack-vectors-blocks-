from __future__ import annotations

import math
from collections import Counter
from dataclasses import dataclass, field
from typing import Any, Dict, List, Optional

from internal.graph.attack_graph import AttackGraph

CHAIN_ENTROPY_THRESHOLD = 0.6


@dataclass
class EntropyResult:
    entropy_level: float
    chain_detected: bool
    potential_chains: List[Dict[str, Any]] = field(default_factory=list)
    recommended_response: Optional[str] = None


class ChainAwareEntropyDetection:
    """Combines Shannon entropy over event actors with graph-topology chain-progression scoring."""

    def __init__(self, attack_graph: AttackGraph) -> None:
        self.attack_graph = attack_graph
        self.chain_models: Dict[str, Any] = {}

    def detect_entropy(self, system_events: List[Dict[str, Any]]) -> EntropyResult:
        standard_entropy = self._calculate_entropy(system_events)
        chain_entropy = self._analyze_chain_progression(system_events)
        # Weight graph-aware chain entropy more heavily; it carries stronger signal.
        combined_entropy = 0.4 * standard_entropy + 0.6 * chain_entropy

        if chain_entropy > CHAIN_ENTROPY_THRESHOLD:
            potential_chains = self._map_to_attack_chains(system_events)
            return EntropyResult(
                entropy_level=combined_entropy,
                chain_detected=True,
                potential_chains=potential_chains,
                recommended_response=self._calculate_chain_response(potential_chains),
            )

        return EntropyResult(
            entropy_level=combined_entropy,
            chain_detected=False,
        )

    def _calculate_entropy(self, system_events: List[Dict[str, Any]]) -> float:
        if not system_events:
            return 0.0
        counts = Counter(e.get("actor", "") for e in system_events)
        total = len(system_events)
        raw = -sum((c / total) * math.log2(c / total) for c in counts.values())
        # Normalise to [0, 1] against the theoretical maximum log2(|actors|).
        max_entropy = math.log2(max(len(counts), 2))
        return min(1.0, raw / max_entropy)

    def _analyze_chain_progression(self, system_events: List[Dict[str, Any]]) -> float:
        """Score how many events map to active (unblocked) edges in the attack graph."""
        if not system_events:
            return 0.0
        score = 0.0
        for event in system_events:
            actor = event.get("actor", "")
            target = event.get("target", "")
            edge = self.attack_graph.edges.get((actor, target))
            if edge and not edge.blocked:
                score += edge.weight * edge.impact
        return min(1.0, score / len(system_events))

    def _map_to_attack_chains(self, system_events: List[Dict[str, Any]]) -> List[Dict[str, Any]]:
        chains: List[Dict[str, Any]] = []
        for event in system_events:
            actor = event.get("actor", "")
            target = event.get("target", "")
            edge = self.attack_graph.edges.get((actor, target))
            if edge and not edge.blocked:
                chains.append(
                    {
                        "source": actor,
                        "target": target,
                        "technique_id": edge.technique_id,
                        "technique_name": edge.technique_name,
                        "risk": round(edge.weight * edge.impact, 4),
                    }
                )
        return chains

    def _calculate_chain_response(self, potential_chains: List[Dict[str, Any]]) -> str:
        if not potential_chains:
            return "monitor"
        max_risk = max((c.get("risk", 0.0) for c in potential_chains), default=0.0)
        if max_risk > 0.8:
            return "isolate"
        if max_risk > 0.5:
            return "deception"
        return "consensus_verification"
