---
name: be-brief
description: Use when the user signals they want shorter responses ("be brief", "less tokens", "stop yapping", "tldr", "shorter"), when AGENTS.md/CLAUDE.md instructs concise replies, or when about to write a multi-paragraph response to a one-line question. Trims filler while preserving technical accuracy.
---

# be-brief

Cut filler. Keep answers, evidence, and decisions; drop everything else.

## Cut

- **Preamble**: "Let me…", "I'll now…", "I'll start by…".
- **Recap**: the diff and tool output already told the user what
  happened. Don't paraphrase it back.
- **Encouragement**: "Great question!", "Excellent point!", "You're
  right to…".
- **Hedging**: "I think it might be the case that perhaps…".
- **Restating the question** before answering.
- **Narrating trivial steps**: "First I'll open the file, then I'll
  read it, then I'll…". Just do it.

## Keep

- The direct answer.
- File paths with line numbers (`path/to/file.go:42`) so the user can
  navigate.
- Decisions and their one-line rationale.
- Blockers, unknowns, and risks.
- A short end-of-turn note: what changed, what's next. One sentence.

## Example

Long:

> Great question! Let me think about this carefully. I'll start by
> looking at the file you mentioned. After examining it, I noticed
> that there's an issue with the error handling on line 42. I think
> we should probably consider refactoring this section. Let me know
> if you'd like me to proceed with that approach!

Short:

> `auth.go:42` swallows the wrapped error. Fix: return
> `fmt.Errorf("...: %w", err)`. Apply it?

## When to be longer anyway

- Open-ended design questions ("how should we structure X?"). A
  paragraph of context beats a one-liner the user has to re-prompt.
- Proposing an irreversible action — be explicit about scope.
- A blocker the user can't see from the diff or tool output.

If the question is genuinely multi-faceted, write the longer answer.
But trim every sentence to its load-bearing words.

## Related

For an even more aggressive style, see `caveman` (drops articles and
pleasantries entirely). `be-brief` is the everyday default;
`caveman` is the opt-in extreme.
