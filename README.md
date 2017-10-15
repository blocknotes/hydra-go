# Hydra Go

A multi-project API system.

Features:

- create REST API using specific endpoints (called collections)
- basic auth for admin routes
- built using Gin, a Mongo DB is required

To do:

- projects management

![hydra](hydra.png)

### Admin examples

```sh
# Basic auth - auth credentials are loaded from AuthCollection (see Conf)
curl -u 'username:password' 'http://localhost:8080/admin/projects'
# list all projects
curl 'http://127.0.0.1:8080/admin/projects'
# create a new project (ex. hydra1)
curl -X POST 'http://127.0.0.1:8080/admin/projects' -H 'Content-Type: application/json' --data '{"data":{"name":"hydra1","description":"Hydra 1","collections":[]}}'
# read a specific project (ex. hydra1)
curl 'http://127.0.0.1:8080/admin/projects/hydra1'
# update project - create collection (ex. hydra1 - articles)
curl -X PUT 'http://127.0.0.1:8080/admin/projects/hydra1' -H 'Content-Type: application/json' --data '{"data":{"collections":[{"name":"articles","singular":"article","columns":{"title":{"type":"String"},"email":{"type":"String","validations":"required,email"},"description":{"type":"String"},"position":{"type":"Float"},"published":{"type":"Boolean"},"dt":{"type":"DateTime"}}}]}}'
# destroy a specific project (ex. hydra1)
curl -X DELETE 'http://127.0.0.1:8080/admin/projects/hydra1'
```

### Api examples

```sh
# list the content of a collection (ex. articles)
curl 'http://127.0.0.1:8080/api/hydra1/articles'
# create an entry for the content of a collection (ex. articles)
curl -X POST 'http://127.0.0.1:8080/api/hydra1/articles' -H 'Content-Type: application/json' --data '{"data":{"title":"A test"}}'
# read an entry of the content of a collection (ex. articles)
curl 'http://127.0.0.1:8080/api/hydra1/articles/59e3b9b7d23d8028efd327c4'
# update an entry for the content of a collection (ex. articles)
curl -X PUT 'http://127.0.0.1:8080/api/hydra1/articles/59e3b9b7d23d8028efd327c4' -H 'Content-Type: application/json' --data '{"data":{"title":"A test 2"}}'
# destroy an entry for the content of a collection (ex. articles)
curl -X DELETE 'http://127.0.0.1:8080/api/hydra1/articles/59e3b9b7d23d8028efd327c4'
```
