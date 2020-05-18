### goapi resources

```
    1. build_and_deploy shell script - helps in automating most of your build and deploy activities.
    2. deployment - file used in the shell script to create the go api deployment file.
    3. Dockerfile - used for building the goapi image.
    4. goapi-service.yaml - Service defination for the goapi.
    5. goapi.go - main source file for the goapi.
```

### Docker Commands

```
docker build -t avk19/goapi:latest .
docker push avk19/goapi:latest
Image name in docker hub : docker.io/avk19/goapi
```