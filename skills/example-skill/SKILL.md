---
name: example-skill
description: Reference skill demonstrating the SKILL.md layout. Use when an author needs a template to copy when creating a new skill in this repo.
allowed-tools: [Read, Write, Edit]
---

# Example skill

This skill exists as a template. Copy the directory, rename it, and
edit the frontmatter and body to fit the new skill.

## When to use

The author description should describe a real trigger condition. For a
template skill the trigger is: "I want to add a new skill to this
repo and need a known-good starting point."

## Steps

1. `cp -r skills/example-skill skills/<new-slug>`.
2. Edit `SKILL.md` frontmatter — set `name` to `<new-slug>` and rewrite
   the `description` to lead with concrete triggers.
3. Replace the body with the actual instructions for the new skill.
4. Run `mise run validate` to confirm the frontmatter passes.
5. Open a PR.

## See also

- [`../../docs/skill-spec.md`](../../docs/skill-spec.md) — full spec
- [`references/style-notes.md`](references/style-notes.md) — body style
