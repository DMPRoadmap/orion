# Contributing to this project

_The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED", "MAY", and "OPTIONAL" in this document are to be interpreted as described in [RFC 2119](https://datatracker.ietf.org/doc/html/rfc2119)._


Thank you for your interest in contributing to this project!
We welcome contributions of all types, whether it is adding new features, fixing bugs, improving performance, or enhancing documentation.

## Development Workflow

This project follows the **[TUBA flow](https://www.adavanzo.com/articles/2025/tuba-trunk-based-batch-release-flow)** (Trunk-Based Batch Release) Flow. All contributions stem from `main`, and changes are released in coordinated **batch releases**.

### Summary of TUBA flow

* You MUST work off the `main` branch at all times.
* You SHOULD use **short-lived branches** to isolate changes.
* You SHOULD submit **small, frequent PRs**.
* Releases are **batched** and **coordinated**, not continuous or feature-based.
* Integration and release coordination are handled by maintainers.

#### Best Practices

* You SHOULD AVOID long-lived branches or merging from anything other than main.
* You SHOULD keep changes small and reviewable.
* You SHOULD communicate in your PR if your work impacts others or requires coordination.
* You SHOULD NOT block on getting your feature shipped — trust the batch release cycle.

---


### Contribution Steps

* You MUST fork this repository (you MUST NOT clone the upstream directly).
* Before starting work, you MUST sync your fork with the latest main branch from upstream to ensure you have the most recent changes.
* You MUST create a short-lived feature or fix branch from your fork’s main branch.
  * Your branch SHOULD only last a few days and focus on a single purpose.
  * You MUST write clean, testable code.
  * You SHOULD include or update unit/integration tests where relevant.
  * You MUST ensure all tests pass locally before submitting.
  * You MUST keep your fork updated with the latest main branch.
* You MUST open a Pull Request from your feature/fix branch to the upstream main branch.
  * You SHOULD follow the Pull Request guidelines.
  * PRs SHOULD be **atomic**, **well-scoped,** and **not dependent on other branches**.

Your PR WILL be reviewed, and once approved, MAY be merged into main and/or queued for the next batch release by the maintainers.

### Batch Releases

Maintainers group approved changes and release them in **batches**, typically on a schedule or milestone basis. However, it MAY be possible for a merged PR to be released immediately. This helps:

* Ensure greater release stability
* Reduce coordination overhead
* Keep `main` always production-ready

If your change is urgent or time-sensitive, mention it in the PR for consideration.

## Submitting a Pull Request

* Push your feature branch to your fork
* Start from the latest `main`, create a short-lived feature or fix branch, and push it to your fork.
* Open a Pull Request (PR) against the upstream `main` branch
  * Do not open PRs directly against any nr/ or next_release/ branches. These are managed by maintainers.
* In the PR description, include:
  * Any related issue numbers, e.g., Related to #42. You MUST NOT use words like 'Close' or 'Fix', as release timing is managed separately.
  * A brief summary of the change.
  * Any caveats, limitations, or follow-ups others should be aware of.
  * If the change is time-sensitive, you MUST explain why. This helps maintainers prioritize batch inclusion.


## Reporting Issues

Before submitting a new issue:

* Search existing issues to avoid duplicates
* Include clear steps to reproduce if it's a bug
* Be descriptive and specific

# Thank You

We appreciate your effort and enthusiasm.
Your contributions help drive the project forward!
