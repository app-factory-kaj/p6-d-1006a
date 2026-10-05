# greeter — PRD

## Problem Statement

Teams building and testing services on this platform need a minimal, predictable HTTP service to exercise against — for smoke-testing pipelines, onboarding new engineers to platform conventions, and serving as a reference implementation. Today there is no small, canonical example service they can point tooling at, so each team improvises its own throwaway stub.

## Solution

Greeter is a small Go HTTP service exposing a single endpoint that returns a JSON greeting for a given name. It follows the platform's reference conventions (as established in app-factory-kaj/e2e-reference) so it can double as both a working utility and a template for how a minimal service on this platform is built.

## Actors

- **API Consumer** — any client (person or system) that calls the greeter endpoint to obtain a greeting; no sign-in or identity is involved.

## User Stories

1. As an API Consumer, I want to GET /hello?name=X, so that I receive a JSON greeting addressed to that name.
2. As an API Consumer, I want to GET /hello without a name, so that I still receive a sensible default greeting rather than an error.

## Product Decisions

- Sign-in / access control: the service is an open, unauthenticated API — no login or API key required, consistent with its role as a small reference/demo service. *assumed*
- Missing-name behavior: when `name` is omitted, the service responds 200 with a generic default greeting (e.g. "Hello, World!") rather than an error. *assumed*
- Implementation follows the conventions established in app-factory-kaj/e2e-reference, per the brief.

## Out of Scope

- Persistence of any kind — the service is stateless.
- Any endpoint beyond GET /hello.
- Authentication, authorization, or rate limiting.
- Localization or multi-language greetings.

## Open Questions

None at this time.

## Further Notes

None.