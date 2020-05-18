# go-api-course

Repository containing all the resources for go-api-course in Udemy


### goapi 
#### Contains the below resources for the goapi

```
		1. Dockerfile 
		2. build_and_deploy shell script
		3. YAML files for deploying the GO API inside the minikube cluster
		4. The go file for the API
```


### postgresdb 
#### Contains all the resources for the postgresdb

```
		1. YAML files for deploying the postgres db inside the minikube cluster
		2. deploy shell script
```

### Architecture Diagram

https://github.com/aveekroy/go-api-course/blob/master/Architecture.png

### Minikube commands

```
minikube start --driver=virtualbox
minikube status
minikube stop
minikube start
minikube delete
minikube service <service name> --url
minikube ip
minikube dashboard
```

### kubectl commands

Resource:
https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands

```
kubectl cluster-info
kubectl api-resources
kubectl get pods
kubectl get svc
kubectl describe pods - describes all the pods
kubectl describe pod <pod_name>
kubectl apply -f <yaml file>
kubectl delete -f <yaml file>
kubectl logs <pod_name>
kubectl proxy - after that open the localhost
kubectl explain <kind name>
```

