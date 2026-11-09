-- Explicit rollback only. No CASCADE: refuse to remove a non-empty schema.
BEGIN;
DROP SCHEMA platform;
COMMIT;
