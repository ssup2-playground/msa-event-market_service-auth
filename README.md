# msa-event-market_service-auth service-auth

service-auth is the service responsible for user management and authentication/authorization in the [msa-event-market](https://github.com/ssup2-playground/project_msa-event-market) Project. service-auth follows this [considerations](https://github.com/ssup2-playground/project_msa-event-market?tab=readme-ov-file#service-considerations)

* Architecture

<img src="/images/architecture.png" width="800"/>

* ER Diagram

<img src="/images/er-diagram.png" width="800"/>

* [Swagger](https://ssup2-playground.github.io/msa-event-market_service-auth/api/openapi/swagger.html)

## Authentication/Authorization

service-auth uses simple authentication based on **ID/Password**. A user can get the **access token** and **refresh token** based on  **JWT** required for authentication/authorization by entering ID/Password. Passwords are encrypted and stored using the **PBKDF2** algorithm.

In JWT token contains the following information.  

* User ID (UUID)
* Login ID
* Role (admin, user)

## Used main external packages and tools

service-auth uses following external packages and tools.

* **HTTP Server, Middleware** : [chi](https://github.com/go-chi/chi), [HTTP](https://pkg.go.dev/net/http), [oapi-codegen](https://github.com/deepmap/oapi-codegen)
* **GRPC Server, Intercepter** : [grpc](https://pkg.go.dev/google.golang.org/grpc), [protoc-gen-go](https://pkg.go.dev/github.com/golang/protobuf/protoc-gen-go)
* **MySQL** : [GORM](https://gorm.io/index.html)
* **Kafka** : [kafka-go](https://github.com/segmentio/kafka-go), [Debezium Outbox](https://debezium.io/documentation/reference/1.8/transformations/outbox-event-router.html)
* **Authorziation** : [Casbin](https://casbin.org/)
* **Logging, Tracking** : [zerolog](https://github.com/rs/zerolog), [Istio](https://istio.io/), [OpenTracing](https://opentracing.io/), [Jaeger](https://www.jaegertracing.io/)
* **Continuous Integration** : [Testify](https://github.com/stretchr/testify), [sqlmock](https://github.com/DATA-DOG/go-sqlmock), [Mockery](https://github.com/mockery/mockery), [Github Actions](https://github.com/features/actions)
* **Continuous Deployment** : [K8s](https://kubernetes.io/), [ArgoCD](https://argo-cd.readthedocs.io/en/stable/), [ArgoCD Image Updater](https://github.com/argoproj-labs/argocd-image-updater), [Kustomize](https://kustomize.io/)

## Development Environment

* Golang Version : 1.16

### Required CLI

* For HTTP Server, Swagger

```
// MacOS
$ go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.6.0
$ go install github.com/mikefarah/yq/v4
```

* For GRPC Server

```
// MacOS
$ brew install protobuf
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

* For Test, CI

```
// MacOS
$ brew install jq
$ brew install mockery
$ go install github.com/fullstorydev/grpcurl/cmd/grpcurl@v1.8.7
$ go install github.com/nektos/act@latest
```

## Reference

* UUID for DB
  * https://github.com/google/uuid/issues/20
* Istio GRPC 
  * https://stackoverflow.com/questions/62459006/how-to-route-multiple-grpc-services-based-on-path-in-istiokubernetes
* Casbin RBAC 
  * https://github.com/luk4z7/middleware-acl
* OpenTracing tracer
  * https://github.com/jaegertracing/jaeger-client-go/blob/master/zipkin/README.md
* OpenTracing middleware
  * https://github.com/go-chi/httptracer
* OpenTracing intercepter
  * https://github.com/opentracing-contrib/go-grpc
