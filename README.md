# PostgreSQL Scheduler

## Description
This is a simple scheduler that allow to automate repetitive task such as daily database scheduling.

![Demo This Project](./assets/img/demo.png)

## Usage

### Prerequisites
1. Any version of Docker.
2. 'bash' compiler.
3. 'make' compiler.
4. .env file, we already provide .env.example as example.
5. Set all method at pg_hba.conf to trust

### Execute
1. Run Locally

    Running this app locally will automatically build and run this application.
    ```
        make start
    ```
2. via Docker

    Running via docker will required docker installed in your device. Authomatically create and run docker container.
    ```
        make build
    ```

### Transfer backup files

1. Locally
Run this following command to transfer backup files from container to host.
```
    docker cp <containerId>:/file/path/within/container /host/path/target
```

2. Using AWS S3
Set AWS_S3_BACKUP value to true in .env configuration and set your personal AWS S3 bucket configurations :
```
AWS_BUCKET_NAME=mybucket
AWS_REGION=ap-southeast-1
AWS_CREDENTIAL_KEY=mysupersecretkey
AWS_CREDENTIAL_SECRET=mysupersecretsecret
AWS_URI_PREFIX=https://mybucket.s3.ap-southeast-1.amazonaws.com/
```
