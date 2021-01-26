message-to-slack
---

A scheduled task that posts a message to Slack.

## Usage


1. Clone repository

    ```bash
    $ git clone https://github.com/michimani/hello-ecs.git
    $ cd hello-ecs/message-to-slack
    ```

2. Create Copilot environment

    ```bash
    $ copilot job deploy --name message-to-slack-job --env test
    ```

3. Put Slack Webhook URL to SSM Parameter Store

    ```bash
    $ aws ssm put-parameter \
    --name /test/slack-webhook \
    --value "https://hooks.slack.com/services/XXXXXXX" \
    --type String \
    --tags Key=copilot-environment,Value=test Key=copilot-application,Value=message-to-slack
    ```

4. Update manifest.yaml

    ```diff
    - schedule: "0 15 31 12 *"
    + schedule: "*/5 * * * *"  # your own schedule
    ```

5. Deploy

    ```bash
    $ copilot job deploy --name message-to-slack-job --env test
    ```

6. Receive message from app

    After a certain amount of time, the scheduled task will be executed and Slack will receive a message.  
    ![sc__2021-01-23_131233](https://user-images.githubusercontent.com/9986092/105568718-eee24200-5d7e-11eb-9c19-1f7a6e8da2c1.png)
