# BAO architecture and certification learning plan

Starting point: no existing certifications. Dates, weekly study hours and exam bookings are not set. Begin with foundations and practical exercises, then progress to architectural decisions. This plan uses BAO as a case study; building one application does not cover a complete certification syllabus.

## Learning route

| Stage | Study focus | Certification relationship | BAO evidence |
| --- | --- | --- | --- |
| Foundations | Client and server, HTTP and TLS, DNS, IP addresses, compute, databases, identity, regions, availability, shared responsibility and cost | AWS Cloud Practitioner and Azure Fundamentals learning material; foundation exams are optional checkpoints | Explain a request travelling from a browser to the Go process and PostgreSQL |
| AWS architecture | Secure access, networking, compute, managed databases, messaging, resilience and cost tradeoffs | Initial architecture target AWS Solutions Architect Associate SAA-C03 | Reviewed AWS deployment design and small isolated labs |
| Azure administration | Identity, networking, compute, storage, monitoring and operational changes | Azure Administrator Associate AZ-104 | Implement and explain equivalent Azure labs |
| Azure architecture | Governance, data, business continuity and infrastructure design | Azure Solutions Architect Expert requires Azure Administrator Associate and AZ-305 | Compare Azure design decisions against the AWS design |
| Enterprise architecture foundations | Stakeholders, concerns, principles, scope, requirements and architecture development | TOGAF Enterprise Architecture Foundation | Architecture vision, stakeholder map and traceability |
| Enterprise architecture application | Apply architecture development, alternatives, roadmaps, governance and change management | TOGAF Enterprise Architecture Practitioner | Explain decisions and manage change across BAO milestones |

Start with the foundation material. Avoid studying three advanced syllabuses simultaneously. Use TOGAF concepts alongside the project from the beginning, with exam preparation as a distinct activity. Select one primary deployment cloud after introductory labs; use the other for equivalent exercises rather than requiring an active dual-cloud system.

## PostgreSQL comes first in the practical work

The repository contains a PostgreSQL driver and verified-TLS connection code. It does not yet have a running BAO database or a completed live database integration test.

The first practical database story must produce:

- An isolated development PostgreSQL instance with documented ownership and lifecycle.
- A trusted certificate and hostname-verified TLS; plaintext attempts must fail.
- Separate administrative, migration and runtime access with tested permissions.
- A documented and verified storage-encryption arrangement, including relevant backups and keys. An ordinary Docker volume alone is not evidence of encryption.
- A migration history and locking approach before application schema evolves.
- A passing live Go connection test and a migration/rollback exercise on disposable data.
- A backup and restore exercise with synthetic data, plus a clear account of what each layer protects.

Provisioning a paid cloud database requires a concrete design and cost review first. Local labs must also honour the encryption requirements. Do not silently relax TLS or storage controls to make a demonstration work.

## Lessons connected to delivery

| BAO topic | Cloud concepts to learn | Architecture concepts to learn |
| --- | --- | --- |
| PostgreSQL connection | Private connectivity, identity, secrets, TLS, managed database responsibilities | Interfaces, dependency inversion, configuration and failure boundaries |
| Schema and transactions | Managed database operations, storage and recovery | Invariants, constraints, atomicity and concurrency |
| Quote acceptance | Scaling, availability and database contention | Aggregates, command handling and consistency |
| Outbox and notifications | Queues, topics, retries and dead-letter handling | Event delivery, idempotency and eventual consistency |
| Documents | Object storage, access controls and encryption keys | Tenant isolation and data lifecycle |
| Payments and ledger | Audit logs, recovery and operational controls | Financial invariants, reconciliation and immutable records |
| Deployment | Networking, workload identity, observability and budgets | Quality attributes, tradeoffs and architecture governance |

## Per lesson routine

Explain the problem in plain language. Draw the smallest useful diagram. Make a small change together, inspect the result and deliberately test a failure. Finish with a teach-back question and link the evidence to the story. Explain decisions before introducing frameworks or service names.

Record learning progress separately from implementation progress: Not started, Studying, Lab completed, Can explain, Exam review. An automated implementation is not evidence that the learner understands it. Use official practice assessments and the full syllabus to identify gaps before scheduling an exam.

## Official study references

Reviewed 9 November 2026. Recheck exam requirements before booking.

- [AWS Cloud Practitioner](https://aws.amazon.com/certification/certified-cloud-practitioner/)
- [AWS Solutions Architect Associate exam guide](https://docs.aws.amazon.com/aws-certification/latest/solutions-architect-associate-03/solutions-architect-associate-03.html)
- [Azure Fundamentals](https://learn.microsoft.com/en-us/credentials/certifications/azure-fundamentals/)
- [Azure Administrator Associate](https://learn.microsoft.com/en-us/credentials/certifications/azure-administrator/)
- [Azure Solutions Architect Expert](https://learn.microsoft.com/en-us/credentials/certifications/azure-solutions-architect/)
- [TOGAF certification paths](https://help.opengroup.org/hc/en-us/articles/11147047948305-What-certifications-are-available-for-the-TOGAF-Standard-10th-Edition)
- [PostgreSQL TLS](https://www.postgresql.org/docs/current/ssl-tcp.html)
- [AWS RDS PostgreSQL TLS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Concepts.General.SSL.html)
- [Azure PostgreSQL TLS](https://learn.microsoft.com/en-us/azure/postgresql/security/security-tls)
