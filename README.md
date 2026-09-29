# Particular Set of Skills

![particular-set README banner](docs/assets/readme-banner.png)

A collection of provider-agnostic agent skill specs.

Each skill is a self-contained directory under `skills/` with a `SKILL.md`
manifest (YAML frontmatter + Markdown body) and any bundled resources it
needs. Skills are written to be readable by any agent runtime, but the repo
is also packaged as a plugin for Claude Code and Codex so it can be
installed directly in either.

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

## Installing

The skill directories follow the [Agent Skills](https://agentskills.io)
format, so the same `skills/<slug>/` tree works in every runtime below.
Only the packaging wrapper differs.

### Claude Code

```text
/plugin marketplace add haydenk/particular-set
/plugin install particular-set@particular-set
```

Marketplace manifest: `.claude-plugin/marketplace.json`.

### Codex

Add the repo as a plugin marketplace, then install `particular-set`
from the plugin directory:

```sh
codex plugin marketplace add haydenk/particular-set
```

Plugin manifest: `.codex-plugin/plugin.json`. Marketplace catalogue:
`.agents/plugins/marketplace.json`.

To pull a single skill instead of the whole set, use the bundled
installer from inside Codex:

```text
$skill-installer install https://github.com/haydenk/particular-set/tree/master/skills/<slug>
```

Codex also discovers repo-local skills under `.agents/skills/`. That
directory is a symlink to `skills/` here, so any Codex session opened
inside a checkout of this repo sees every skill without installing.

### Other runtimes

Any agent that reads the Agent Skills format can consume a skill by
copying its directory into the runtime's skills folder, for example
`~/.claude/skills/` or `~/.agents/skills/`. Failing that, each
`SKILL.md` is plain Markdown: paste the body into your agent's system
prompt or have your runtime load it directly.

When bumping the plugin version, change it in both
`.claude-plugin/marketplace.json` and `.codex-plugin/plugin.json`.
`mise run check` fails if they disagree.

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
