package query

import "context"

// Pinger is owned by the query that consumes it; pgxpool satisfies it directly.
type Pinger interface{ Ping(context.Context) error }
type Readiness struct{ Database Pinger }

func (q Readiness) Execute(ctx context.Context) error { return q.Database.Ping(ctx) }
