// Module particular-set exposes Dagger functions that check the skill
// specs in this repo. All checks run hermetically in containers, so the
// host only needs the Dagger CLI.
package main

import (
	"context"
	"fmt"
	"path"
	"regexp"
	"sort"
	"strings"

	"dagger/particular-set/internal/dagger"

	"gopkg.in/yaml.v3"
)

type ParticularSet struct{}

// Check runs every skill check and returns combined output.
func (m *ParticularSet) Check(ctx context.Context, source *dagger.Directory) (string, error) {
	v, err := m.Validate(ctx, source)
	if err != nil {
		return v, err
	}
	l, err := m.Lint(ctx, source)
	if err != nil {
		return v + "\n" + l, err
	}
	return v + "\n" + l, nil
}

// Validate parses every skills/<slug>/SKILL.md and verifies its frontmatter
// against the spec in docs/skill-spec.md.
func (m *ParticularSet) Validate(ctx context.Context, source *dagger.Directory) (string, error) {
	paths, err := source.Glob(ctx, "skills/*/SKILL.md")
	if err != nil {
		return "", fmt.Errorf("globbing SKILL.md files: %w", err)
	}
	sort.Strings(paths)

	if len(paths) == 0 {
		return "", fmt.Errorf("no skills found under skills/")
	}

	var out strings.Builder
	failed := 0

	for _, p := range paths {
		contents, err := source.File(p).Contents(ctx)
		if err != nil {
			return out.String(), fmt.Errorf("reading %s: %w", p, err)
		}

		errs := validateSkill(p, contents)
		if len(errs) == 0 {
			fmt.Fprintf(&out, "ok   %s\n", p)
			continue
		}
		failed++
		fmt.Fprintf(&out, "FAIL %s\n", p)
		for _, e := range errs {
			fmt.Fprintf(&out, "     - %s\n", e)
		}
	}

	fmt.Fprintf(&out, "\n%d skill(s) checked, %d failed\n", len(paths), failed)
	if failed > 0 {
		return out.String(), fmt.Errorf("%d skill(s) failed validation", failed)
	}
	return out.String(), nil
}

// Lint runs markdownlint-cli2 against all Markdown in the repo, inside a
// pinned container so no host Node toolchain is needed.
func (m *ParticularSet) Lint(ctx context.Context, source *dagger.Directory) (string, error) {
	return dag.Container().
		From("davidanson/markdownlint-cli2:v0.14.0").
		WithUser("root").
		WithMountedDirectory("/src", source).
		WithWorkdir("/src").
		WithExec([]string{
			"markdownlint-cli2",
			"skills/**/*.md",
			"docs/**/*.md",
			"README.md",
			"AGENTS.md",
			"SECURITY.md",
			"CODE_OF_CONDUCT.md",
		}).
		Stdout(ctx)
}

var slugRe = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)

type frontmatter struct {
	Name         string   `yaml:"name"`
	Description  string   `yaml:"description"`
	AllowedTools []string `yaml:"allowed-tools"`
}

func validateSkill(file, contents string) []string {
	var errs []string

	if !strings.HasPrefix(contents, "---\n") {
		return []string{"missing YAML frontmatter"}
	}
	end := strings.Index(contents[4:], "\n---")
	if end == -1 {
		return []string{"unterminated frontmatter"}
	}
	raw := contents[4 : 4+end]

	var fm frontmatter
	if err := yaml.Unmarshal([]byte(raw), &fm); err != nil {
		return []string{fmt.Sprintf("invalid YAML: %v", err)}
	}

	dirName := path.Base(path.Dir(file))

	switch {
	case fm.Name == "":
		errs = append(errs, "name: required")
	case !slugRe.MatchString(fm.Name):
		errs = append(errs, fmt.Sprintf("name: %q is not kebab-case", fm.Name))
	case fm.Name != dirName:
		errs = append(errs, fmt.Sprintf("name: %q does not match directory %q", fm.Name, dirName))
	}

	switch {
	case fm.Description == "":
		errs = append(errs, "description: required")
	case strings.Contains(fm.Description, "\n"):
		errs = append(errs, "description: must be a single line")
	case len(strings.TrimSpace(fm.Description)) < 20:
		errs = append(errs, "description: too short — lead with concrete triggers")
	}

	return errs
}
