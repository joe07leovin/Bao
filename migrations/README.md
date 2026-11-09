# PostgreSQL migrations

Versioned SQL pairs are applied explicitly. Do not run down migrations automatically at startup. The bootstrap migration is transactional and must be applied once to a new database.

With a trusted CA and verified-TLS DATABASE_URL configured in the operator environment:

```sh
psql "$DATABASE_URL" -v ON_ERROR_STOP=1 -f migrations/000001_platform.up.sql
```

Use a dedicated migration identity. Runtime credentials must not have schema administration permissions. A migration history and locking tool must be selected before additional application migrations are introduced. Storage encryption and backups must be supplied and verified by infrastructure.
