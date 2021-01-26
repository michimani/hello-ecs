list-s3-buckets
---

A Load Balanced Web Service that returns list of S3 Bucket name in your AWS account.

## Usage

1. Clone repository

    ```bash
    $ git clone https://github.com/michimani/hello-ecs.git
    $ cd hello-ecs/list-s3-buckets
    ```

2. Create Copilot environment

    ```bash
    $ copilot env init --name test --profile default --default-config
    ```

3. Deploy

    ```bash
    $ copilot deploy --env test
    ```

4. Rquest

    Check the application domain with the `copilot svc show` command, and request. (with `curl` command)

    ```bash
    $ curl -Ssi http://<your-app-domain>/buckets
    HTTP/1.1 200 OK
    Date: Tue, 26 Jan 2021 14:31:10 GMT
    Content-Type: application/json; charset=UTF-8
    Transfer-Encoding: chunked
    Connection: keep-alive

    {"buckets":["bucket1",...,"bucket999"]}
    ```


## Run at locally

1. Build

    ```bash
    $ docker build -t list-s3-buckets .
    ```

2. Run

    ```bash
    $ docker run \
    --rm \
    -p 9000:1323 \
    -e AWS_DEFAULT_REGION="ap-northeast-1" \
    -e AWS_ACCESS_KEY_ID="<your-aws-access-key-id>" \
    -e AWS_SECRET_ACCESS_KEY="<your-aws-secret-acess-key>" \
    list-s3-buckets:latest
    ```

3. Request

    ```bash
    $ curl -Ssi localhost:9000/buckets
    HTTP/1.1 200 OK
    Content-Type: application/json; charset=UTF-8
    Date: Tue, 26 Jan 2021 14:39:09 GMT
    Transfer-Encoding: chunked

    {"buckets":["bucket1",...,"bucket999"]}
    ```