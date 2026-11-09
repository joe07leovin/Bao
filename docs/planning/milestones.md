# BAO delivery plan

This is the initial milestone plan. Dates and estimates are uncommitted. Items below are planned work, not completed designs or existing GitHub milestones.

| Milestone | Exit condition | Initial design work |
| --- | --- | --- |
| Domain and architecture | Context ownership, invariants and consequential architecture decisions reviewed | Context map, service boundaries, shared system diagrams |
| Secure foundation | Go and PostgreSQL foundation, identity boundaries, encryption, CI and reliable event delivery verified | Tenant isolation, key management, CQRS, outbox and notification delivery |
| Firm and business onboarding | Registration, authority checks, eligibility and permissions verified | Firm and business registration and verification states |
| Marketplace and quoting | Guided requests, discovery, private quotes and single-winner acceptance verified | Quote lifecycle, scope comparison, concurrent acceptance |
| Engagement delivery | Professional checks, signing, secure documents and lifecycle verified | Engagement state machine, document access and communication |
| Payments and ledger | Validated provider flow, balanced postings and reconciliation verified | Funds flow, payment states, transaction boundaries, disputes and recovery |
| Pilot readiness | Remaining URS scope explicitly triaged and production checks verified | Security, accessibility, recovery, operations and complete journeys |

## Story lifecycle

Backlog → Design in progress → Design review → Ready for implementation → In progress → Code review → Verification → Done.

Record blockers explicitly. The intended stages must be configured in a board or tracked in issue descriptions. A template alone does not satisfy design review.

Each implementation story links its exact URS requirements, acceptance criteria, technical design, relevant diagrams, dependencies and verification evidence. Split broad design areas into implementable stories after domain modelling.

## Open decisions

- Independently deployed service boundaries and their data ownership.
- AWS or Azure and the managed event and notification services.
- Payment provider, funds flow, holding constraints and applicable obligations.
- Payments and ledger transaction boundaries and read consistency.
- Business policy values, including commission, deposits, expiry and credits.
- Regulator authority, remit, data access and retention.

The existing Linear project is historical setup. GitHub is the user's chosen tracking destination. No automatic synchronisation is configured.
