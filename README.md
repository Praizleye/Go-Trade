# Go-Trade

> A Go-based, non-custodial trading platform that starts as a personal trading system, grows into a trusted family product, and can later expand into a public copy-trading and subscription business.

## Current Status

Go-Trade is currently in the vision and foundation stage.

This repository does not yet contain a production-ready trading platform. This README describes the intended direction, system boundaries, architectural vocabulary, and rollout strategy so implementation can begin from a clear plan instead of a vague idea.

## Project Thesis

The goal of Go-Trade is to build trading infrastructure that is useful at every stage of growth:

1. Start small and build a reliable system for my own spot and futures trading.
2. Extend that system to a small, trusted family or close-circle environment.
3. Mature it into a public-facing platform with subscription access and managed copy-trading controls.
4. Explore, only later and very carefully, whether a trusted self-hosted or source-available distribution model makes sense.

The platform is intended to be non-custodial. Users connect their own exchange accounts and keep control of their funds. Go-Trade should focus on automation, execution, risk management, observability, and copy-trading coordination rather than taking custody of assets.

## Principles

- Non-custodial first: exchange accounts stay under the user's control.
- Safety before scale: risk controls matter more than feature count.
- Honest staging: personal use comes before family rollout, and family rollout comes before public access.
- Auditability matters: actions, signals, and copied trades should be explainable and traceable.
- Modular growth: the system should begin small but support later expansion into multi-user and subscription workflows.
- Deterministic execution: if AI or prompt-based tooling is used, it should support research and operator workflows, not replace auditable trade controls.

## Product Direction

### Phase 1: Solo Operator

The first version should solve a real personal problem:

- Connec one or more exchange accounts for spot and futures trading.
- Run a small number of strategies with clear, deterministic rules.
- Enforce basic personal risk limits around exposure, leverage, and stop conditions.
- Track positions, orders, fills, balances, and trade history.
- Provide operator-friendly logging, alerts, and manual override controls.

This phase is about proving the core engine, not building a marketplace.

### Phase 2: Trusted Family / Small Circle

Once the solo flow is stable, Go-Trade can expand into a private multi-user system for a small family or trusted group:

- Invite-only access for a small number of users.
- Separate account connections and risk settings per user.
- Shared strategies with follower-specific limits.
- Early copy-trading relationships between a strategy leader and trusted followers.
- Better permissioning, audit trails, notifications, and operational visibility.

This phase should still optimize for trust, not growth.

### Phase 3: Public Platform

After the core trading workflows, observability, and private copy-trading controls are proven, Go-Trade can expand into a public product:

- Public onboarding for approved users.
- Monthly subscription plans with feature and usage limits.
- Public copy-trading workflows with stronger safety controls and better transparency.
- Operational monitoring, support tooling, and more formal compliance posture.
- Clear product separation between personal use, trusted-circle use, and commercial hosted use.

This phase turns a private tool into a real software business.

### Phase 4: Future Exploration

One longer-term idea is to let trusted users or organizations run a copy of the platform on their own infrastructure under a separate commercial arrangement. That could eventually include controlled source access, update entitlements, and license enforcement mechanisms.

This is only a future exploration. It is not a committed near-term feature, and it should not shape the early architecture more than necessary.

## Why Go

Go is the right foundation for this project because it supports:

- Fast and efficient concurrent execution for trading workflows.
- A simple deployment model with portable binaries.
- Good ergonomics for services, workers, APIs, and CLI tooling.
- Predictable performance for long-running infrastructure.
- A clean standard-library-first ecosystem that works well for a staged product buildout.

## Technical Direction

The current technical direction is intentionally practical and minimal:

| Area | Planned Choice | Reason |
| --- | --- | --- |
| Core language | Go | Strong concurrency, portability, and service ergonomics |
| HTTP/API layer | Go standard library (`net/http`, `ServeMux`) | Keeps the platform simple, dependency-light, and aligned with modern Go routing |
| CLI/admin tooling | `cobra` | Good fit for operational commands and local tooling |
| Primary database | PostgreSQL | Durable relational storage for users, trades, subscriptions, and audit data |
| Cache / queue support | Redis | Fast caching, lightweight queues, rate limiting, and temporary coordination |
| Exchange connectivity | Exchange adapter layer | Keeps spot and futures integrations modular |
| Observability | Structured logs, metrics, traces | Necessary for trust, debugging, and operations |

This stack should remain small at the beginning. The goal is to avoid over-engineering while still leaving room for scale.

For HTTP routing and middleware, the default direction should be standard-library first. Modern Go provides enough routing capability through `net/http` and `ServeMux` to support a clean initial API without needing a heavier web framework. If the API surface becomes more complex later, a thin router such as `chi` can still be added without changing the broader architecture.

For a practical setup guide and official reading links, see [docs/go-net-http-setup.md](docs/go-net-http-setup.md).

## High-Level Architecture

Go-Trade should evolve around a few clear modules:

### Strategy Engine

The strategy engine is responsible for generating signals and rule-based actions from market data, account state, and configured strategy parameters.

It should favor deterministic and testable logic, even if AI-assisted research or prompt-driven operator tooling is introduced later.

### Execution Engine

The execution engine turns approved strategy actions into exchange orders, manages retries and error handling, and records the final execution outcome.

This layer should be conservative, observable, and tightly integrated with risk controls.

### Portfolio and Risk Manager

The portfolio and risk manager enforces exposure limits, leverage boundaries, stop rules, account constraints, and user-specific guardrails before trade execution is allowed.

This is a core trust layer and should exist from the earliest phase.

### Copy-Trading Coordinator

The copy-trading coordinator manages how leader strategies map to follower accounts, how follower-specific limits are applied, and how copied actions are approved, scaled, paused, or rejected.

This should be designed for safety first, not for viral growth.

### Subscription and Billing Module

The subscription module becomes important once Go-Trade expands into a public service. It should model plan entitlements, billing status, feature gates, and access limits without affecting the trading core.

### Audit, Logging, and Notification Layer

Every important action should produce a useful trail:

- signal created
- risk rule checked
- order submitted
- order filled or rejected
- copy-trade decision applied
- user notified

This layer is essential for trust, support, and operations.

## Planned Core Concepts

The README should establish stable vocabulary for implementation even before final code APIs exist.

### `ExchangeAdapter`

A unified interface for exchange connectivity, market data, balances, positions, orders, fills, and account metadata across spot and futures venues.

### `StrategyEngine`

A service or module that converts market inputs, account state, and strategy configuration into deterministic trading actions.

### `RiskProfile`

A per-user or per-account policy object that defines leverage limits, exposure caps, instrument restrictions, stop behavior, and copy-trading permissions.

### `CopyRelationship`

A model that maps a leader strategy or leader account to one or more follower accounts, including opt-in rules, scaling rules, instrument filters, and stop conditions.

### `SubscriptionPlan`

A business-layer concept that controls billing entitlements, feature availability, account limits, and public product tiers.

### `DeploymentMode`

A concept representing how the platform is being used:

- personal
- family/private
- hosted public
- future trusted self-hosted

These are conceptual contracts, not finalized code interfaces.

## Copy Trading Model and Safety Boundaries

Copy trading is one of the long-term goals of Go-Trade, but it must be handled carefully.

The intended model is:

- Leaders do not take custody of follower funds.
- Followers connect their own exchange accounts.
- Followers explicitly opt in to a leader or strategy.
- Each follower keeps independent risk controls and limits.
- The system may scale or reject copied trades based on follower constraints.
- Followers should be able to pause, disconnect, or override copied execution.

Important boundaries:

- Copied trades may not match the leader exactly because of slippage, account size, latency, market conditions, or exchange constraints.
- Copy trading should never imply guaranteed returns.
- Risk controls must be applied on the follower side before execution.
- Transparency should be favored over opaque automation.

## Subscription Vision

If Go-Trade reaches the public phase, the business model should be straightforward:

- monthly subscription plans
- feature-based access tiers
- limits based on users, accounts, strategies, or copy relationships
- private or premium capabilities for advanced users

The initial commercial posture should stay simple. A subscription-first model is easier to operate and explain than complex fee structures in the beginning.

## Security and Trust

Security is central to this project because trading systems handle sensitive credentials and high-risk actions.

Core expectations:

- Exchange API credentials should be encrypted at rest.
- Secrets should be isolated from application logic where possible.
- Logs should never expose raw credentials or unsafe account data.
- Sensitive actions should be auditable.
- Users should have clear visibility into what strategies are active and why a trade happened.
- Admin and operator actions should be permissioned and traceable.

Over time, Go-Trade should also support:

- secret rotation workflows
- strong authentication for operators and users
- alerting for abnormal behavior
- role-based access controls
- operational dashboards and incident visibility

## Deployment Direction

The primary direction is hybrid local plus cloud:

- Early personal use may run locally or in a small private environment.
- Family or trusted-circle use may run in a lightly hosted setup with stronger monitoring.
- Public rollout likely requires a more formal hosted control plane.
- A future self-hosted option can be explored only after the hosted architecture is stable.

This direction preserves flexibility without forcing a heavy platform design too early.

## AI and Prompt Engineering

AI-assisted workflows can be useful in this project, but they should be used carefully.

Good uses:

- research assistance
- operator summaries
- log interpretation
- idea generation for strategies
- internal tooling support

Bad uses:

- unaudited autonomous trade execution
- vague natural-language rules with no deterministic safeguards
- replacing hard risk controls with probabilistic outputs

If prompt engineering becomes part of the system, it should live beside deterministic trading logic, not instead of it.

## Planned Repository Shape

One reasonable future repository structure could look like this:

```text
cmd/
  api/
  worker/
  cli/
internal/
  exchange/
  strategy/
  execution/
  risk/
  portfolio/
  copytrade/
  billing/
  audit/
  notify/
pkg/
  models/
  config/
deploy/
docs/
```

This is only a likely shape, not a fixed contract.

## Near-Term Build Order

To keep the project grounded, the most sensible implementation order is:

1. Build one clean exchange adapter for a single target exchange.
2. Add account configuration, risk profiles, and deterministic strategy execution.
3. Record orders, fills, balances, and operator-visible audit logs.
4. Add alerts, manual controls, and simple operational tooling.
5. Expand into trusted multi-user support and early copy-trading relationships.
6. Add public subscription and billing only after the private model is proven.

## Legal and Risk Notice

Trading spot and futures markets is risky. Software can fail. Exchanges can fail. Market conditions can change faster than automated systems react.

Go-Trade should be treated as trading infrastructure, not financial advice. Nothing in this repository should be read as a promise of profitability or loss prevention.

Any public version of the platform should make the following clear:

- users remain responsible for their trading decisions
- copy trading does not guarantee identical performance
- past results do not guarantee future outcomes
- users should understand the risks of leverage, liquidation, and market volatility

## License and Future Commercial Path

This repository is currently licensed under the MIT License.

That means the code published here today is open under MIT terms. At the same time, the broader Go-Trade vision may eventually include hosted services, private operational environments, or separately packaged commercial offerings with different terms.

One speculative future idea is a trusted self-hosted distribution model where a customer can run the system on their own infrastructure under a commercial agreement, potentially with update-gated entitlements or other license controls.

That idea is still exploratory. It is not part of the current MIT commitment for this repository, and it should be treated as a future business and product discussion rather than an active feature promise.

## Closing Direction

Go-Trade is meant to be built in stages:

- useful for one operator first
- safe for a trusted circle next
- credible as a public business later

If the project stays disciplined about non-custodial design, risk control, auditability, and gradual rollout, it can grow from a personal Go trading system into a serious platform without losing trust along the way.
