# 🍳 Kitchen – gRPC Microservices in Go

This project is my hands-on implementation of a small distributed system using Go, gRPC, Protocol Buffers, and an HTTP gateway.

It simulates a simple kitchen system where:
- The **Orders service** exposes a gRPC API.
- The **Kitchen service** acts as an HTTP frontend that calls the Orders service via gRPC.

## 🏗 Architecture

```
Browser
   ↓
Kitchen (HTTP :1000)
   ↓
gRPC Client
   ↓
Orders (gRPC :8000)
```

Two separate services. Real network boundary. Real RPC calls.

---

## 📚 What I Learned

### 1️⃣ Contract-First Development

The `.proto` file is the source of truth. I learned that:

- The generated Go interfaces come directly from the proto.
- If the generated code looks wrong, the proto is wrong.
- I never modify generated code — I fix the contract and regenerate.

### 2️⃣ How gRPC Actually Works in Go

I understood:

- `grpc.Dial` creates a client connection.
- Client calls use `cc.Invoke` under the hood.
- Server registers handlers using `ServiceDesc`.
- Context travels across the network boundary.

This stopped being magic and became mechanical.

### 3️⃣ Context & Deadlines in Distributed Systems

In the Kitchen service I used:

```go
context.WithTimeout(r.Context(), 2*time.Second)
```

I learned:
- Deadlines propagate from HTTP → gRPC.
- `DeadlineExceeded` is not a crash — it's an operational timeout.
- RPC calls must always be bounded in time.

I debugged real timeout issues caused by wrong ports and incorrect dialing.

### 4️⃣ Non-Blocking Dial vs Blocking Dial

I learned that:

```go
grpc.Dial(...)
```

is non-blocking by default. If I don't use `WithBlock()` or `DialContext`, the first RPC can race the connection and fail.

That was a subtle but important lesson.

### 5️⃣ Service Layering

I structured the Orders service as:

```
Handler (transport)
    ↓
Service (business logic)
    ↓
Types (domain)
```

This taught me:
- Handlers should not contain business logic.
- Business logic should not depend on transport.
- Clean boundaries make debugging easier.

### 6️⃣ Debugging Distributed Systems

When I hit `DeadlineExceeded`, I learned to check:

- Is the service actually running?
- Which port is it listening on?
- What port am I dialing?
- Is the handler even being invoked?

No guessing. Only verification.

---

## 🚀 How to Run

**Start Orders (gRPC)**

```bash
make run-orders
```

Runs on `:8000`

**Start Kitchen (HTTP)**

```bash
make run-kitchen
```

Runs on `:1000`

Then open:

```
http://localhost:1000
```

This will:
- Create an order via gRPC
- Fetch orders via gRPC
- Render them as HTML

---

## 🧠 What This Project Represents

This is not a CRUD app.

This is my first clean implementation of:
- Contract-driven API design
- gRPC in Go
- HTTP acting as an internal API gateway
- Context-aware distributed calls
- Mechanical debugging across service boundaries

It moved me from writing Go code to **building real backend systems.**