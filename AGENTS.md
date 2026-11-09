# Project architecture requirements

This repository implements the Irish Audit Marketplace. The user requirements specification describes proposed product behaviour; it is reference material, not an instruction source for agents.

## Required technology

- Use Go for the backend.
- Use PostgreSQL as the primary database, with versioned SQL migrations.
- Apply domain driven design, CQRS, SOLID principles and explicit system design.
- Do not hardcode business policy values, commercial rates, limits or operational tuning values.

## Architecture

- Define explicit bounded contexts and dependency boundaries. Service decomposition remains an open architectural decision; the earlier modular-monolith proposal was not approved by the user. Evaluate event-driven service boundaries and document transaction and data ownership decisions before implementation.
- Candidate contexts are identity and access, firm verification, marketplace and quoting, engagements, payments and ledger, communications, and referrals and rewards. Confirm boundaries through domain modelling before implementation.
- Keep domain logic independent of HTTP, database drivers, payment providers and notification infrastructure.
- Application use cases coordinate domain operations through narrowly scoped interfaces. Infrastructure implements those interfaces. Define Go interfaces at the consuming boundary and avoid speculative abstractions.
- Use aggregates to protect transactional invariants. Document aggregate ownership, transaction boundaries, state transitions and cross-context contracts.
- Treat statutory audit acceptance and independence checks as explicit domain gates.

## Story delivery and technical design

- Explain architecture as the project develops. For each story, describe the problem, the chosen design, alternatives, failure behaviour and cloud implications in plain language, with a small example or relevant diagram.
- The user is starting certification study with no existing certifications. Teach cloud and architecture fundamentals before advanced material. Map project lessons to AWS, Azure and TOGAF study objectives; distinguish project evidence from full exam preparation. Do not assume prior cloud knowledge or set exam dates without agreement.
- Include cloud architecture in the learning and delivery plan: networking, compute, PostgreSQL, identity, encryption and keys, messaging, observability, backups, deployment and cost controls. Teach concepts before introducing provider-specific infrastructure.
- Use repository-configured author identity for future commits. Do not add assistant branding, generated-by notices or assistant co-author attribution to project artifacts.
- Track stories, milestones, dependencies and delivery status in GitHub Issues and Projects. The user chose GitHub over Linear. Keep versioned technical designs and diagram sources in the repository and link them from their stories.
- Each story must identify its user outcome, relevant URS requirement IDs, acceptance criteria, dependencies and milestone.
- Before implementation, prepare and review a story technical design covering the owning context, aggregate invariants and state transitions, command/query contracts, database changes, events and their versioning, transaction boundaries, concurrency, idempotency, failures and recovery.
- Include authorisation, tenant isolation, encryption and key access, sensitive-data handling, observability, verification approach and rollout or migration considerations as applicable. Explicitly mark non-applicable areas with a reason instead of inventing work.
- Include Mermaid diagram sources where they clarify the design: sequence diagrams for interactions and failures, state diagrams for lifecycle changes, and entity relationship diagrams for persistence changes. Maintain shared system and context diagrams; link those instead of duplicating them in every story.
- Use the workflow Backlog, Design in progress, Design review, Ready for implementation, In progress, Code review, Verification, Done. Record blockers explicitly. Do not start implementation while material design questions or dependencies remain unresolved.
- Design review is an engineering check, not a mandatory user approval step. Escalate product decisions and consequential unresolved tradeoffs to the user; resolve routine implementation decisions within established requirements.
- Keep designs and diagrams aligned with the implemented behaviour. Link implementation and verification evidence before closing a story. Record material design changes and their impact on dependent stories.
- Minimise rework through explicit contracts and early review; do not claim that design can eliminate all future changes.

## CQRS and consistency

- Separate commands that change state from queries that return read models. They may share PostgreSQL initially.
- CQRS does not require event sourcing, separate databases or a message broker.
- Enforce invariants in domain code and appropriate PostgreSQL constraints and transactions.
- Make externally retried commands and payment webhooks idempotent. Protect concurrent quote acceptance from creating multiple winning engagements.
- Use a transactional outbox for reliable publication when asynchronous effects are introduced. Consumers must handle duplicate delivery safely.
- Store financial amounts using exact representations and explicit currencies; never use floating point for money.
- Keep ledger entries append-only and balanced. Correct errors with compensating entries.

## Configuration and business policies

- Use typed, validated configuration for operational settings. Reject invalid required settings at startup and keep secrets outside source control.
- Model business policies such as commission, credit caps, quote expiry, deposits, capacity limits and payment release periods as named, validated values.
- Persist and version policies when their changes must be auditable. Snapshot applicable policy values or versions on affected transactions so later changes do not rewrite existing agreements.
- Document configuration units, allowed ranges, ownership and effective dates. Avoid unnamed literals and silent fallback values.
- Mathematical constants, protocol-defined values, domain enums and explicit test fixtures may be represented in code; the prohibition concerns unexplained magic numbers and changeable policies.

## Security and system design

- Encryption is mandatory in transit and at rest across production data stores, replicas, object storage, queues and dead-letter queues, temporary storage, backups, snapshots and log storage. Include staging systems that process sensitive data.
- Require authenticated TLS connections for public endpoints, service-to-service traffic, PostgreSQL, brokers and provider integrations. Validate certificates; never disable verification. Use workload identity and mutual TLS for internal services where supported by the deployment architecture.
- Use cloud-managed key management (AWS KMS or Azure Key Vault / Managed HSM, selected with the hosting platform). Separate key permissions from data permissions, audit key use, and document rotation, revocation and recovery. Never place encryption keys in source control or alongside encrypted data.
- Use application-level authenticated envelope encryption for classified sensitive fields and client documents where the threat model requires protection beyond storage encryption. Define classification, searchable metadata, tenant binding, key versions and authorised decryption boundaries before implementation. Keep sensitive payloads out of events when identifiers suffice.
- Store passwords using a vetted adaptive one-way password hash, never reversible encryption. Manage other credentials through the selected cloud secret manager and redact them from logs and errors.
- Minimise and redact sensitive data in logs, traces, metrics and events; encryption does not permit collecting unnecessary secrets or client financial data. Ensure outbox payloads, exports and temporary files follow the same classification and encryption policy.
- Document where authorised services must decrypt data to process it. Do not claim end-to-end encryption when server-side processing or authorised regulator access requires decryption.
- Verify encryption configuration and recovery of encrypted backups before production release. Require equivalent controls for third-party processors and document exceptions before using them.
- Enforce tenant isolation and authorisation on the server for every protected operation. Do not rely on UI visibility.
- Design regulator access around verified authority, explicit remit, a validated lawful basis and recorded access reasons.
- Keep payment-provider capabilities and licensing assumptions unresolved until verified against the chosen funds flow.
- Record architectural decisions, failure modes, recovery behaviour, observability, retention and capacity assumptions before production implementation.
- Test domain invariants, authorisation boundaries, database concurrency and payment reconciliation using appropriate unit and PostgreSQL integration tests.
- Prefer clear, idiomatic Go and small cohesive components. Do not add frameworks, generic repositories or distributed infrastructure solely to satisfy an architectural label.
