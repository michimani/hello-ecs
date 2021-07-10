hello-ecs
---

This repository is a collection of samples of applications running on ECS that provisioned by AWS Copilot.


## Usage

0. Install AWS Copilot

    If your local machine's OS is macOS, install it using `brew` command.
    
    ```bash
    brew install aws/tap/copilot-cli
    ```
    
    In other OS case, please read AWS Copilot document.  
    [(Optional) Verify the installation - AWS Copilot CLI](https://aws.github.io/copilot-cli/docs/getting-started/verify/)
    
    ```bash
    copilot version
    ```
    
    ```bash
    version: v1.8.3, built for darwin
    ```

1. Please execute it referring to the README of each application.

## Sample Apps

### [list-s3-buckets](https://github.com/michimani/hello-ecs/tree/master/list-s3-buckets)  

A Load Balanced Web Service that returns list of S3 Bucket name in your AWS account.


### [message-to-slack](https://github.com/michimani/hello-ecs/tree/master/message-to-slack)  

A scheduled task that posts a message to Slack.

