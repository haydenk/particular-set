# Skill specification

A skill is a self-contained directory that teaches an agent how to do
one well-scoped thing. This repo ships skills as plain Markdown so any
agent runtime can consume them.

## Directory layout

```text
skills/<slug>/
  SKILL.md            required
  references/         optional — long-form docs the skill links to
  scripts/            optional — helper scripts the skill may invoke
  assets/             optional — static files (templates, fixtures)
```

`<slug>` is kebab-case and must equal the `name:` field in
`SKILL.md` frontmatter.

## SKILL.md format

```markdown
---
name: example-skill
description: One-line trigger describing when an agent should invoke this skill.
allowed-tools: [Read, Grep, Bash]   # optional, Claude Code only
---

# Example skill

Short body. State the goal, the steps, and any non-obvious rules.
Link to `references/*.md` for depth.
```

### Required fields

| field         | rule |
|---------------|------|
| `name`        | kebab-case, equals directory name |
| `description` | single line, ends with a period, leads with concrete triggers |

### Optional fields

| field           | meaning |
|-----------------|---------|
| `allowed-tools` | Claude Code: tools the skill is allowed to call |
| `metadata`      | free-form object for runtime-specific hints |

Unknown frontmatter keys are allowed but ignored by validation.

## Description style

The `description` is the only thing an agent sees before deciding to
load a skill. Treat it as the trigger sentence, not a tagline.

Good:

> Use when the user asks to scaffold a new Express route, add middleware,
> or wire up route-level error handling.

Bad:

> A skill for working with Express.

## Body style

- Lead with the goal in one sentence.
- Prefer numbered steps over prose for procedures.
- Push depth into `references/` files and link with relative paths.
- Keep the body short enough that an agent loads the whole thing
  without truncation — target under 150 lines.

## Validation

`mise run validate` (or `dagger call validate --source=.`) parses
every `SKILL.md`, checks the rules above, and fails the build on any
violation. Pre-commit runs the same check on staged files; CI runs
`mise run check` on every PR.
