# particular-set

A repository of provider-agnostic agent skill specs. Also packaged as a
plugin for Claude Code and Codex so it can be installed in either
directly, but the specs themselves are not tied to any particular
runtime.

This file (`AGENTS.md`) is the canonical instruction document for any
coding agent working in this repo. Claude Code, Codex, Cursor, and
Copilot all read it natively; there is no `CLAUDE.md`.

## Scope

This repo holds general-purpose developer skills, agents, and related
specs: things any engineer on any codebase could use. Anything specific
to a company, client, team, or single project belongs in its own repo,
not here. If a skill only makes sense with knowledge of one
organisation's systems, conventions, or data, it is out of scope.

## Stack

- Markdown skill specs under `skills/<slug>/SKILL.md`
- Plugin wrappers: `.claude-plugin/marketplace.json` (Claude Code),
  `.codex-plugin/plugin.json` + `.agents/plugins/marketplace.json`
  (Codex). `.agents/skills` is a symlink to `skills/` so Codex picks
  up skills when run inside this repo.
- Dagger Go module under `.dagger/` runs every check in containers
- mise pins `dagger` (and `go`, for editing the module)
- GitHub Actions for CI — calls Dagger
- No host Node toolchain; markdownlint runs in a container

## Commands

Tasks are wired through mise so the same command works locally and in CI.

- `mise install` — install Dagger and Go to versions pinned in `mise.toml`
- `mise run check` — run validate + lint via Dagger
- `mise run validate` — frontmatter validation only
- `mise run manifests` — plugin manifest consistency only
- `mise run lint` — Markdown lint only
- `mise run install-hooks` — install `.git/hooks/pre-commit` (idempotent)

Direct equivalents if you prefer:

- `dagger call check --source=.`
- `dagger call validate --source=.`
- `dagger call manifests --source=.`
- `dagger call lint --source=.`

## Editing the Dagger module

Source lives at `.dagger/main.go`. After changing function signatures
run `dagger develop` to regenerate `dagger.gen.go`. `go.mod`/`go.sum`
updates happen with normal `go get` / `go mod tidy`.

## Conventions

- One skill per directory under `skills/`. Directory name must equal the
  `name:` field in frontmatter.
- Skill names are kebab-case.
- `description:` must be a single line — that line is what an agent reads
  to decide whether to invoke the skill. Lead with concrete triggers.
- Keep `SKILL.md` bodies short. Push depth into `references/` and link
  with relative paths so the agent can pull them in on demand.
- Specs should be provider-agnostic. Runtime-specific fields go in
  optional frontmatter keys (e.g. `allowed-tools` for Claude Code) and
  must not be required for the skill to make sense to other runtimes.

## Do not touch without approval

- `.claude-plugin/marketplace.json`, `.codex-plugin/plugin.json`,
  `.agents/plugins/marketplace.json` — changes here affect every consumer
  who has installed the plugin. Bump versions deliberately, and bump the
  Claude and Codex versions together; `mise run check` fails if they
  differ.
- `.github/workflows/` — CI config. Local edits often hide real issues.
- `dagger.json` — engine version is pinned to match the mise-installed
  CLI. Bump both together.

## Adding a skill

Use the example at `skills/example-skill/` as a template. Run
`mise run check` before opening a PR.
