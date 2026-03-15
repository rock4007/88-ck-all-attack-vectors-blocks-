package main

import (
"context"
"crypto/rand"
"encoding/hex"
"log"
"net/http"
"os"
"os/signal"
"syscall"
"time"

"github.com/88ck/pillar1-morphic/internal/gamma"
"github.com/88ck/pillar1-morphic/internal/scheduler"
"github.com/88ck/pillar1-morphic/internal/xds"
)

func main() {
ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
defer stop()

policyID := randomID(8)
s := scheduler.New(750 * time.Millisecond)
x := xds.NewPublisher("morphic-gateway")
g := gamma.NewCoupler(0.42)

mux := http.NewServeMux()
mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
w.WriteHeader(http.StatusOK)
_, _ = w.Write([]byte("ok"))
})
mux.HandleFunc("/tick", func(w http.ResponseWriter, _ *http.Request) {
decision := s.NextDecision(policyID)
coupled := g.Apply(decision.Score)
x.Publish(decision.Policy, coupled)
_, _ = w.Write([]byte("scheduled"))
})

srv := &http.Server{
Addr:              ":8080",
Handler:           mux,
ReadHeaderTimeout: 3 * time.Second,
}

go func() {
<-ctx.Done()
shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
_ = srv.Shutdown(shutdownCtx)
}()

log.Printf("morphic gateway listening on %s", srv.Addr)
if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
log.Fatalf("gateway failed: %v", err)
}
}

func randomID(n int) string {
b := make([]byte, n)
if _, err := rand.Read(b); err != nil {
return "fallback-policy"
}
return hex.EncodeToString(b)
}
