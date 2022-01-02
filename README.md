# Git-Utilities

Utility to automate git workflows.

## Commands

* `userstory` get/set/update the userstory config value
* `commit-msg` git `commit-msg` hook. Treats the commit message as a go template.
  - `{{ .US }}` template in the userstory, if configured.

Plugins to rewrite

1. ~~git userstory~~
2. git userstory-branch
3. ~~commit hook to replay {{US}} with config value~~
4. git templates and git init