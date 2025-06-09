# Stock Price Fetcher

> This project was based on [Building a Microservice with Golang and Docker (JSON And gRPC)](https://www.youtube.com/watch?v=367qYRy39zw&list=PL0xRBLFXXsP5cru52B5GAQmIrTTAL8A66&index=1) and [Building a Microservice with Golang and Docker - gRPC Transport](https://www.youtube.com/watch?v=D0St2LH158Q&list=PL0xRBLFXXsP5cru52B5GAQmIrTTAL8A66&index=7)

Description
-

A Go-based microservice for fetching stock prices, designed to demonstrate idiomatic Go patterns in both HTTP and gRPC APIs:

- Exposes a JSON-based HTTP endpoint for on-demand stock price queries
- Supports real-time streaming of stock prices using gRPC
-️ Built as a lightweight, modular service suitable for distributed systems

This project explores Go’s concurrency model, networking packages, and clean service-oriented design with real-world data integration.

> [!NOTE]
> I will add a call to a third party system for fetching data instead of hard coding it a db in the RAM as in the videos above