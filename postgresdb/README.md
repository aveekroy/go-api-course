### postgresdb resources

```
1. All the yaml files defination for deploying it into the minikube cluster.
2. shell script to delete and deploy the files all in one go
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

UPDATE users SET created_at = 'timestamp' where id=1;

```
