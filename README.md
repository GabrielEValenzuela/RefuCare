# 🩺 RefuCare

RefuCare is a cross-language, fault-tolerant microservices lab designed for teaching distributed systems through the lens of post-crisis healthcare management. It simulates a resilient medical logistics and diagnostics system using Go, Java, and Python.

```mermaid
flowchart TD
    subgraph Gateway
        Traefik[Traefik]
    end

    subgraph Services
        Vitals[Fiber]
        Records[Spring]
        Analyzer[FastAPI]
    end

    subgraph Databases
        Postgres[(PostgreSQL)]
        Mongo[(MongoDB)]
        Redis[(Redis)]
    end

    subgraph Infrastructure
        Eureka[Eureka]
        RabbitMQ[RabbitMQ]
        Prometheus[Prometheus]
        Grafana[Grafana]
    end

    Client[Client Application] --> Traefik
    Traefik --> Vitals
    Vitals --> Redis


    Vitals --> RabbitMQ
    Analyzer --> RabbitMQ
    Records --> Postgres
    Analyzer --> Mongo

    Vitals --> Eureka
    Records --> Eureka

    Vitals --> Prometheus
    Records --> Prometheus
    Traefik --> Prometheus
    RabbitMQ --> Prometheus
    Prometheus --> Grafana
```
