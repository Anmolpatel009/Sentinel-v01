// Package observation implements Sentinel's Observation Engine.
//
// The Observation Engine is responsible for collecting the current
// state of the host operating system and publishing an atomic
// ObservationSnapshot into the Shared State Repository (SSR).
package observation

/*
Collector Philosophy (Sentinel v1)

A collector is a small, focused component responsible for observing
exactly one subsystem of the operating system.

Examples:

    - CPU
    - Memory
    - Network
    - Processes
    - Filesystem

Each collector has exactly one responsibility:

    Observe its subsystem and populate only its corresponding section
    of the ObservationSnapshot.

Collectors NEVER:

    - Modify another collector's state
    - Write directly to the Shared State Repository
    - Perform analysis or decision making
    - Execute reconciliation actions

Instead, the Observation Manager coordinates all collectors during
an observation cycle.

                Runtime Tick
                      │
                      ▼
             Observation Manager
                      │
      ┌───────────────┼───────────────┐
      ▼               ▼               ▼
     CPU           Memory         Network
      ▼               ▼               ▼
      └───────────────┼───────────────┘
                      ▼
          ObservationSnapshot
                      ▼
                 Publisher
                      ▼
                     SSR

Implementation Note

Sentinel v1 intentionally does NOT define a Collector interface.

Reason:

At this stage, Sentinel contains only built-in collectors and does
not require runtime extensibility or plugins.

A Collector interface will be introduced in Sentinel v2 when:

    - Plugin collectors are supported.
    - Dynamic collector registration is required.
    - Multiple collector implementations exist.

Until then, the Observation Manager directly coordinates concrete
collector implementations.

This follows Sentinel's engineering principle:

    "Do not introduce abstractions until they solve a real problem."
*/