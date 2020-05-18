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

Resource:
https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands

```
### Postgres Commands

```
psql -h localhost -U postgresadmin -d postgresdb

CREATE TABLE users(
	id serial PRIMARY KEY,
	username VARCHAR (50) UNIQUE NOT NULL,
	email VARCHAR (355) UNIQUE NOT NULL,
	created_at TIMESTAMP NOT NULL
);

INSERT INTO users (id, username, email, created_at) VALUES (1, 'user1', 'user1@mail.com', current_timestamp);
INSERT INTO users (id, username, email, created_at) VALUES (2, 'user2', 'user2@mail.com', current_timestamp);
INSERT INTO users (id, username, email, created_at) VALUES (3, 'user3', 'user3@mail.com', current_timestamp);

UPDATE users SET created_at = '2020-05-07 04:29:39.482142' where id=1;
UPDATE users SET created_at = '2020-05-07 04:30:11.647136' where id=2;
UPDATE users SET created_at = '2020-05-07 04:30:27.29615' where id=3;
```
### Docker Commands
```
docker build -t avk19/goapi:latest .
docker push avk19/goapi:latest
docker.io/avk19/goapi
```
