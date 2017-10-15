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
curl -u 'username:password' 'http://localhost:8080/admin/collections'
# list all collections
curl 'http://localhost:8080/admin/collections'
# create a new collection (ex. articles)
curl -X POST 'http://localhost:8080/admin/collections' -H 'Content-Type: application/json' --data '{"data":{"name":"articles","singular":"article","columns":{"title":{"type":"String"},"email":{"type":"String","validations":"required,email"},"description":{"type":"String"},"position":{"type":"Float"},"published":{"type":"Boolean"},"dt":{"type":"DateTime"}}}}'
# read a specific collection (ex. articles)
curl 'http://localhost:8080/admin/collections/articles'
# update a specific collection (ex. articles)
curl -X PUT 'http://localhost:8080/admin/collections/articles' -H 'Content-Type: application/json' --data '{"data":{"name":"articles","columns":{"subtitle":{"type":"String"},"email":{"type":"String","validations":"required,email"}}}}'
# destroy a specific collection (ex. articles)
curl -X DELETE 'http://localhost:8080/admin/collections/articles'
```

### Api examples

```sh
# list the content of a collection (ex. articles)
curl http://localhost:8080/api/articles/articles
# read an entry of the content of a collection (ex. articles)
curl http://localhost:8080/api/articles/articles/59468245cfba25329f3272db
# create an entry for the content of a collection (ex. articles)
curl -X POST http://localhost:8080/api/articles/articles -H 'Content-Type: application/json' --data '{"data":{"title":"A test"}}'
# update an entry for the content of a collection (ex. articles)
curl -X PUT http://localhost:8080/api/articles/articles/59468245cfba25329f3272db -H 'Content-Type: application/json' --data '{"data":{"title":"A test 2"}}'
# destroy an entry for the content of a collection (ex. articles)
curl -X DELETE http://localhost:8080/api/books/books/59462836cfba25329f3272d0
```
