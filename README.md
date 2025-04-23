# 🩺 RefuCare

RefuCare is a cross-language, fault-tolerant microservices lab designed for teaching distributed systems through the lens of post-crisis healthcare management. It simulates a resilient medical logistics and diagnostics system using Go, Java, and Python.

```mermaid
flowchart TD
    subgraph Gateway
        Traefik
    end

    subgraph Services
        Vitals
        Records
        Analyzer
    end

    subgraph Databases
        Postgres[(PostgreSQL)]
        Mongo[(MongoDB)]
        Redis[(Redis Cache)]
    end

    subgraph Infrastructure
        Eureka
        RabbitMQ
        Prometheus
        Grafana
    end

    %% Client interacts with system
    Client --> Traefik

    %% Traefik routes to services
    Traefik --> Vitals
    Traefik --> Records
    Traefik --> Analyzer

    %% Vitals service interactions
    Vitals --> Redis
    Vitals --> RabbitMQ
    Vitals --> Eureka
    Vitals --> Prometheus

    %% Records service
    Records --> Postgres
    Records --> Eureka
    Records --> Prometheus

    %% Analyzer service
    Analyzer --> RabbitMQ
    Analyzer --> Mongo
    Analyzer --> Prometheus

    %% Infrastructure
    Traefik --> Prometheus
    RabbitMQ --> Prometheus
    Prometheus --> Grafana

```
