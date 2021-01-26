hello-ecs
---

This repository is a collection of samples of applications running on ECS that provisioned by AWS Copilot.


## Usage

0. Install AWS Copilot

    If your local machine's OS is macOS, install it using `brew` command.
    
    ```bash
    $ brew install aws/tap/copilot-cli
    ```
    
    In other OS case, please read AWS Copilot document.  
    [(Optional) Verify the installation - AWS Copilot CLI](https://aws.github.io/copilot-cli/docs/getting-started/verify/)
    
    ```bash
    $ copilot version
    version: v1.1.0, built for darwin
    ```

1. Please execute it referring to the README of each application.

## Sample Apps

### [message-to-slack](https://github.com/michimani/hello-ecs/tree/master/message-to-slack)  

Deploy a scheduled task that posts a message to Slack.

