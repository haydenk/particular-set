# Particular Set of Skills

![particular-set README banner](docs/assets/readme-banner.png)

A collection of provider-agnostic agent skill specs.

Each skill is a self-contained directory under `skills/` with a `SKILL.md`
manifest (YAML frontmatter + Markdown body) and any bundled resources it
needs. Skills are written to be readable by any agent runtime, but the repo
is also packaged as a Claude Code plugin so it can be installed directly.

## Layout

```text
skills/
  <skill-name>/
    SKILL.md            required
    references/         optional supporting docs the skill links to
    scripts/            optional helper scripts
    assets/             optional static files
```

## Skill manifest

`SKILL.md` must start with YAML frontmatter:

```yaml
---
name: <kebab-case-slug>          # matches the directory name
description: <one-line trigger>  # used by agents to decide when to invoke
allowed-tools: [optional, list]  # runtime-specific, ignored when absent
---
```

Full spec: [`docs/skill-spec.md`](docs/skill-spec.md).

## Installing as a Claude Code plugin

```text
/plugin marketplace add <git-url>
/plugin install particular-set@particular-set
```

The marketplace manifest lives at `.claude-plugin/marketplace.json`.

## Using skills outside Claude Code

Each `SKILL.md` is plain Markdown. Read it, paste the body into your agent's
system prompt, or have your runtime load it directly.

## Development

All checks run via a Dagger Go module — no host Node toolchain required.
mise pins the Dagger CLI version.

```sh
mise install              # installs dagger (and go for module edits)
mise run install-hooks    # link .git/hooks/pre-commit (idempotent)
mise run check            # validate + lint
mise run validate         # SKILL.md frontmatter only
mise run lint             # markdownlint in a container
```

Or call Dagger directly: `dagger call check --source=.`.

The pre-commit hook lives at `.githooks/pre-commit` and is installed as
a symlink into `.git/hooks/` by `mise run install-hooks`. CI runs the
same `check` function on every PR.

## Adding a skill

1. `mkdir skills/<slug>` and add `SKILL.md`.
2. Run `mise run validate` to confirm frontmatter is well-formed.
3. Open a PR. CI enforces the spec.

See [`skills/example-skill/`](skills/example-skill/) for a reference.
