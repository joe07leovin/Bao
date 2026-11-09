# BAO

An Irish marketplace connecting businesses with verified audit and accounting firms.

## Project tracking

- Stories and technical work: [GitHub Issues](https://github.com/joe07leovin/Bao/issues).
- Delivery stages: [GitHub Milestones](https://github.com/joe07leovin/Bao/milestones).
- Progress board: GitHub Projects when configured.
- Technical designs and editable diagrams: versioned in this repository and linked from each story.

Start with the [delivery plan](docs/planning/milestones.md). Create work using the story issue template. Use the [technical design template](docs/design/story-template.md) before implementation.

Seven GitHub milestones and eight initial design issues have been created. The separate project board is not configured because the current Git credential does not include Projects permission.

## Engineering requirements

Follow the [architecture learning guide](docs/architecture-learning.md) to study the design and cloud infrastructure as each story is implemented.
The [certification roadmap](docs/certification-roadmap.md) starts from fundamentals and maps BAO lessons to AWS, Azure and TOGAF study.

Go, PostgreSQL, domain driven design, CQRS, SOLID and reliable event delivery. Business policy values must be configurable and versioned where needed. Encryption is required in transit and at rest, with controlled decryption and managed keys.

Service boundaries, cloud provider, payment provider and financial consistency boundaries require explicit design decisions. CQRS does not imply event sourcing or independently deployed read databases.

The initial Go service provides HTTPS liveness and PostgreSQL readiness checks, validated environment configuration and graceful shutdown. See [development instructions](docs/development.md), the [foundation design](docs/design/foundation.md) and [AGENTS.md](AGENTS.md). Business endpoints and production infrastructure remain pending.
