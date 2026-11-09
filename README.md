# BAO

An Irish marketplace connecting businesses with verified audit and accounting firms.

## Project tracking

- Stories and technical work: GitHub Issues.
- Delivery stages: GitHub Milestones.
- Progress board: GitHub Projects when configured.
- Technical designs and editable diagrams: versioned in this repository and linked from each story.

Start with the [delivery plan](docs/planning/milestones.md). Create work using the story issue template. Use the [technical design template](docs/design/story-template.md) before implementation.

The GitHub board, remote issues and milestones are separate resources; adding these files does not create them.

## Engineering requirements

Go, PostgreSQL, domain driven design, CQRS, SOLID and reliable event delivery. Business policy values must be configurable and versioned where needed. Encryption is required in transit and at rest, with controlled decryption and managed keys.

Service boundaries, cloud provider, payment provider and financial consistency boundaries require explicit design decisions. CQRS does not imply event sourcing or independently deployed read databases.

See [AGENTS.md](AGENTS.md) for the detailed engineering and delivery requirements. No application implementation has been added yet.
