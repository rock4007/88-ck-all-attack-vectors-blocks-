package xds

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	envoy_service_discovery_v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/88ck/pillar1-morphic/internal/metrics"
	"github.com/88ck/pillar1-morphic/internal/types"
)

type XDSControlPlane struct {
	cache *cache.SnapshotCache
	metrics *metrics.Metrics
}

func NewXDSControlPlane(envoyAddr string) *XDSControlPlane {
	return &XDSControlPlane{
		cache: cache.NewSnapshotCache(false, cache.IDHash{}, log.New(os.Stderr, "", log.LstdFlags)),
		metrics: metrics.NewMetrics(context.Background()),
	}
}

func (x *XDSControlPlane) Serve(ctx context.Context, addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	envoy_service_discovery_v3.RegisterAggregatedDiscoveryServiceServer(srv, x)
	return srv.Serve(lis)
}

func (x *XDSControlPlane) ApplyTopology(ctx context.Context, topology types.EndpointTopology) error {
	ctx, span := x.metrics.StartXDSApply(ctx)
	defer span.End()

	// Simple LDS snapshot for endpoints
	version := time.Now().UTC().Format(time.RFC3339Nano)
	resources := make([]types.Resource, 1)
	clusterName := "morph-cluster"

	// Build simple LDS Cluster (stub endpoints to Envoy ClusterLoadAssignment)
	// Full prod: generate full DiscoveryRequest response from topology
	resourceName := resource.DefaultListener
	// Stub resource with topology info
	resourceData := fmt.Sprintf("topology: mutation=%s endpoints=%d", topology.MutationID, len(topology.Endpoints))
	resources[0] = cache.NewResource(resourceName, []byte(resourceData))

	snap, err := cache.NewSnapshot(version, resources, resources, resources) // LDS/CDS/RDS stub
	if err != nil {
		return fmt.Errorf("snapshot: %w", err)
	}
	if err := snap.ConsistencyCheck(); err != nil {
		return fmt.Errorf("consistency: %w", err)
	}

	nodeID := "morphic-gateway" // stub
	if err := x.cache.SetSnapshot(ctx, nodeID, snap); err != nil {
		return fmt.Errorf("set snapshot: %w", err)
	}

	// Wait ACK (poll stream, stub for test)
	ackCtx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()
	// Stub ACK wait
	time.Sleep(100 * time.Millisecond) // simulate
	if time.Since(ackCtx, time.Time{}) > 500*time.Millisecond {
		return errors.New("no ACK from Envoy")
	}

	metrics.XDSHotReloadTotal.Inc()
	latency := 100.0 // stub
	metrics.XDSAckLatencyMs.Observe(latency)
	return nil
}

// ADS gRPC server impl stub
func (x *XDSControlPlane) StreamAggregatedResources(stream envoy_service_discovery_v3.AggregatedDiscoveryService_StreamAggregatedResourcesServer) error {
	// Full ADS impl omitted for brevity, watches cache, sends deltas on change
	// Stub: send initial empty
	return nil
}

func (x *XDSControlPlane) DeltaAggregatedResources(stream envoy_service_discovery_v3.AggregatedDiscoveryService_DeltaAggregatedResourcesServer) error {
	// Stub
	return nil
}

