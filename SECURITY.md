# Security policy

## Supported versions

This repository tracks `master`. Security fixes land on `master` and are
picked up by anyone who has installed the plugin or vendored the skill
specs. There are no long-lived release branches.

## Reporting a vulnerability

Please do not file public GitHub issues for security reports. Instead,
use GitHub's private vulnerability reporting form:

<https://github.com/haydenk/particular-set/security/advisories/new>

Reports go straight to the maintainers and stay private until a fix is
published.

Include:

- A clear description of the issue and its impact.
- Steps to reproduce (a minimal example skill, payload, or command).
- The commit or tag you reproduced against.
- Any suggested remediation, if you have one.

You should expect:

- An acknowledgement within 3 business days.
- A triage update within 7 business days.
- Coordinated disclosure: we will agree a public disclosure date with
  the reporter before publishing a fix.

## Scope

In scope:

- Skill specs (`skills/**/SKILL.md` and supporting files) that contain
  prompt-injection payloads, credential-exfiltration patterns, or other
  content designed to subvert a consuming agent.
- The Dagger module under `.dagger/` and any CI workflows under
  `.github/workflows/` — supply-chain or code-execution issues.
- The Claude Code plugin manifest at `.claude-plugin/marketplace.json`.

Out of scope:

- Vulnerabilities in upstream tools we shell out to
  (`markdownlint-cli2`, Dagger itself, Go toolchain) — please report
  those to the upstream projects. We will track and update as fixes
  ship.
- Issues that require a malicious maintainer with commit access.

## Hardening guidance for consumers

If you are installing skills from this repo into an agent runtime:

- Review each `SKILL.md` body before granting the runtime broad tool
  access. Skills can include instructions that an agent will follow.
- Pin to a commit or tag rather than tracking `master` if you need a
  stable, reviewed surface.
- Treat `allowed-tools` as a hint, not a sandbox — your runtime is the
  authority on what a skill can actually call.
