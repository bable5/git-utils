# Git-Utilities

Utility to automate git workflows.

![Build Badge](https://github.com/bable5/git-utils/actions/workflows/main.yml/badge.svg)

## Commands

* `userstory` get/set/update the userstory config value
* `commit-msg` git `commit-msg` hook. Treats the commit message as a go template.
  - `{{ .US }}` template in the userstory, if configured.

## Scrips

* `git-userstory`. Forwards arguments to `git-utils userstory`. When it is on the $PATH git will be able to translate `git userstory` into `git-userstory`

## Hooks

* `commit-msg`. Forwards arguments to `git-utils commit-msg`. Place in `.git/hooks`.

Plugins to rewrite

1. ~~git userstory~~
2. git userstory-branch
3. ~~commit hook to replay {{US}} with config value~~
4. git templates and git init