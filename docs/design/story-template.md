# Story technical design

Status: Pending design. This template is not an approved implementation design.

Story link:

URS IDs:

Milestone:

## Outcome and acceptance criteria

Describe the user outcome, scope, exclusions and observable acceptance criteria.

## Domain model

Identify the owning bounded context, aggregate, invariants, states and transitions. Define ownership across boundaries.

## Commands queries and contracts

Specify inputs, outputs, validation, errors, authorisation and compatibility. Define read freshness and when authoritative reads are required.

## Persistence and transactions

Specify PostgreSQL schema, constraints, indexes, migrations, concurrency protection, transaction boundaries and recovery. Use exact financial representations and explicit currencies.

## Events and asynchronous effects

Specify event ownership, schema versions, outbox, delivery guarantees, deduplication, ordering assumptions, retries, dead-letter handling and replay. Prefer identifiers to sensitive payloads.

## Security and encryption

Describe tenant isolation, permissions, data classification, encryption in transit and at rest, application encryption where needed, managed key access, authorised decryption, retention and log redaction.

## Policies and configuration

Name each changeable policy, its units, validation, owner and effective version. Explain which values are snapshotted on transactions. Do not invent business defaults.

## Failures and recovery

Cover partial failures, duplicate requests, provider timeouts, ambiguous outcomes and reconciliation. Define operational alerts and recovery ownership.

## Diagrams

Add editable Mermaid diagrams that explain this story. Use sequence diagrams for interactions and failures, state diagrams for lifecycle changes, and entity relationship diagrams for persistence. Link shared system diagrams. Verify rendered output. Mark irrelevant diagram types with a reason.

## Verification and rollout

Tie tests to acceptance criteria and invariants. Describe compatibility, migration, deployment, recovery and verification evidence.

## Dependencies and decisions

Link blocking stories and architecture decisions. Resolve material questions before coding. Keep the design aligned with the implemented result.

## Architecture walkthrough

Explain the problem in plain language, how the chosen design works, the alternatives and why they were rejected. Walk through one successful request and one failure. Identify what the learner should understand before implementing this story.

## Cloud implications

Describe the relevant compute, network boundaries, workload identity, database, events, encryption keys, secrets, logs, backups and deployment. Include cost drivers and recovery behaviour. Mark irrelevant areas explicitly. Link any provider selection decision; do not assume AWS or Azure has been selected.
