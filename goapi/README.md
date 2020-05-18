### goapi resources

```
    1. build_and_deploy shell script - It can help you in building images, pushing the images to docker hub, 
                                       deploying the goapi to the minikube cluster, clearing up the images in your local MacOS.
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