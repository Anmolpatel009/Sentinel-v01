# Sentinel v1

> A lightweight, modular, autonomous Linux system runtime built in Go.

Sentinel is an experimental systems project that continuously observes a Linux host, analyzes its runtime state, makes decisions, and reconciles the system toward a desired state.

The project is heavily inspired by the reconciliation philosophy used in Kubernetes, but is designed for operating system level resource management instead of cluster management.

---

# Vision

Modern operating systems expose enormous amounts of runtime information but provide very little autonomous decision making.

Sentinel aims to become a lightweight runtime capable of:

- observing the host
- understanding system health
- detecting anomalies
- making optimization decisions
- executing corrective actions
- continuously reconciling the system

```
                Desired State
                      ▲
                      │
               Reconciliation
                      ▲
                      │
                 Backend Engine
                      ▲
                      │
                 Planner Engine
                      ▲
                      │
                Decision Engine
                      ▲
                      │
                 Analysis Engine
                      ▲
                      │
             Shared State Repository
                      ▲
                      │
              Observation Engine
                      ▲
                      │
                  Linux Kernel
```

---

# Sentinel Architecture

```
Runtime
│
├── Configuration
├── Logging
├── Shared State Repository (SSR)
│
├── Observation Engine
│      ├── CPU
│      ├── Memory
│      ├── Network
│      ├── Process
│      └── Filesystem
│
├── Analysis Engine
├── Decision Engine
├── Planner Engine
├── Backend Engine
├── Telemetry
│
└── Reconciliation Loop
```

---

# Current Status

## Completed

- Runtime Lifecycle
- Configuration System
- Logging
- Shared State Repository (SSR)
- Observation Manager
- CPU Collector
- Memory Collector
- Network Collector
- Process Collector
- Filesystem Collector

All components compile successfully.

```
go vet   ✅

go build ✅

go test  ✅
```

---

# Shared State Repository (SSR)

Sentinel uses a centralized Shared State Repository.

Instead of components communicating directly with one another, every engine exchanges information through the SSR.

```
Observation
      │
      ▼
  Publish Snapshot
      │
      ▼
+----------------------+
| Shared State         |
| Repository (SSR)     |
+----------------------+
      ▲
      │
Analysis
      ▲
      │
Decision
      ▲
      │
Planner
      ▲
      │
Backend
```

Advantages

- loose coupling
- thread-safe
- single source of truth
- one writer per engine
- zero duplicated runtime state

---

# Observation Engine

The Observation Engine owns all Linux data collection.

Built-in collectors

- CPU
- Memory
- Network
- Process
- Filesystem

Each observation cycle allocates exactly **one** ObservationState.

Every collector writes directly into its own section.

```
Manager
   │
   ▼
Allocate ObservationState
   │
   ├── CPU
   ├── Memory
   ├── Network
   ├── Process
   └── Filesystem
          │
          ▼
 PublishObservation()
          │
          ▼
         SSR
```

No intermediate copies are created.

---

# Design Principles

Sentinel follows several engineering principles.

## Single Responsibility

Each engine owns exactly one responsibility.

## One Writer

Every engine owns one section of the runtime state.

No engine modifies another engine's state.

## Shared State

All communication happens through SSR.

No cross-component dependencies.

## Lightweight

Sentinel avoids unnecessary abstractions.

Interfaces are introduced only when multiple implementations exist.

## Zero-Copy Observation

Collectors populate a single ObservationState.

The completed snapshot is published once.

---

# Repository Structure

```
cmd/
    sentinel/

internal/

    config/

    logging/

    runtime/

    state/

    observation/
        cpu/
        memory/
        network/
        process/
        filesystem/
```

---

# Roadmap

## Foundation

- [x] Runtime
- [x] Configuration
- [x] Logging
- [x] Shared State Repository
- [x] Observation Layer

## Intelligence

- [ ] Analysis Engine
- [ ] Decision Engine
- [ ] Planner Engine

## Execution

- [ ] Backend Engine

## Control

- [ ] Telemetry
- [ ] Reconciliation Loop

---

# Project Goals

Sentinel is being developed as a research project exploring:

- Linux internals
- Go systems programming
- Autonomous runtime management
- Reconciliation loops
- High-performance observability
- Modular operating-system architecture

---

# Current Version

```
Version: v1 (In Development)
Status : Foundation Complete
Language : Go
Platform : Linux
```

---

# License

This project is currently under active development.