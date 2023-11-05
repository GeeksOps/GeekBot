# GeekBot

## Overview

The **GeekBot** is a Telegram bot built using Go Lang that integrates with the GeekOps community group to enhance user interaction and community engagement.

It is designed to authenticate new members, automate moderation tasks, reward active participants, and foster a friendly and constructive environment ...

## Features and Functionalities

1. New Member Verification
    - The bot will implement a check to determine if a new member is a bot or a human, using CAPTCHA-based challenges or interaction-based verification methods.

2. Account Maintenance
    - GeekBot will automatically remove profiles that have been marked with the status "Deleted."

3. Moderation and Gamification
    - The bot will serve as a moderator and will implement a gamified points system to incentivize positive behavior and community engagement.

4. Points Distribution
    - The bot will award points based on the following actions:
        - Single-word answer that resolves an issue/question: 50 points
        - Detailed answer with rationale for the solution that resolves an issue/question: 100 points
        - Detailed answer with external references that lead to resolution: 150 points
        - Contribution to discussion that aids in resolving an issue: 20 points
        - Inviting a colleague or friend with an interest in IT infrastructure: 50 points
        - Honest recommendation/feedback about an event or meeting in chat or via Google form: 200 points
        - Members will be able to view their points balance using the `/my_points` command.

5. Points Redemption
    - Points can be exchanged for community merch, with 100 point equivalent to 1 UAH. Future plans include the introduction of achievements, and members' suggestions for these are welcomed.

<!-- markdownlint-disable MD033 -->
<p style="display: flex; justify-content: space-around;">
  <img src=docs/images/gift-1.png width="25%" style="margin: 0 10%"/>
  <img src=docs/images/gift-2.png width="25%" style="margin: 0 10%" />
</p>
<!-- markdownlint-enable MD033 -->

1. Community Rule Enforcement
    - GeekBot will maintain a list of prohibited words, primarily related to politics and commonly found in spam messages. If these words are detected, the message will be deleted.

## Technical Specification for GeekBot Telegram Community Assistant

### Stack and Tooling

- **Kubernetes** (K3s + Raspberry Pi setup).
- **Helm**
- **ArgoCD** will automatically deploy the bot to Kubernetes when creating a new artifact with the tag.
- **Mozilla SOPS** for secrets
- **Go Lang** + Cobra + [telebot](https://github.com/tucnak/telebot)
- **GitHub Actions** CI/CD pipeline

## Study Materials

- [learn-go-with-tests](https://quii.gitbook.io/learn-go-with-tests)
- [Go in Action, Second Edition](https://www.manning.com/books/go-in-action-second-edition)
- [Shipping Go Develop, deliver, discuss, design, and go again](https://www.manning.com/books/shipping-go)
- [Build an Orchestrator in Go (From Scratch)](https://www.manning.com/books/build-an-orchestrator-in-go-from-scratch)
