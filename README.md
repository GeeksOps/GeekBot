# GeekBot

[![Go Report Card](https://goreportcard.com/badge/github.com/geekopsua/geekbot)](https://goreportcard.com/report/github.com/geekopsua/geekbot) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/a8a1691dee8a438780caab8cb578353c)](https://app.codacy.com/gh/GeekOpsUA/GeekBot/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)

## Overview

The **GeekBot** is a Telegram bot built using Go Lang that integrates with the GeekOps community group to enhance user interaction and community engagement.

It is designed to authenticate new members, automate moderation tasks, reward active participants, and foster a friendly and constructive environment...
Please read the [Concept](docs/plan/concept.md) for more information.

## Requirements

To run this project, you need to have the following tools installed:

- ![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/GeekOpsUA/GeekBot)

Also, please check that `kubectl` and `helm` are installed on your machine.
Additionally, we are using following tools:

- **Go Packages**
  - [Cobra](https://cobra.dev/#getting-started)
  - [Telebot](https://github.com/tucnak/telebot#getting-started)
- **Kubernetes**
  - [K3D](https://k3d.io/#installation)
  - [ArgoCD](https://argoproj.github.io/argo-cd/getting_started/#1-install-argo-cd)
  - [Mozilla SOPS](https://github.com/getsops/sops#sops-secrets-operations)
- **Helpful Tools**
  - [Taskfile](https://taskfile.dev/#/)
- **GitHub Actions**
  - [markdown-lint](https://github.com/marketplace/actions/markdown-lint)
  - [golangci-lint](https://github.com/marketplace/actions/run-golangci-lint)
  - [Memer Action](https://github.com/marketplace/actions/memer-action)

## Getting Started

### Working with Taskfile

To list all available tasks, run:

```bash
task --list-all
```

To get more information about a specific task, run:

```bash
task --summary <task-name>
```

Or, you can run:

```bash
task --summary
```

For default task.
