// Package observation implements Sentinel's Observation Engine.
//
// The Observation Engine is responsible for collecting the current
// state of the host operating system and publishing one complete
// ObservationState into the Shared State Repository (SSR).
//
// The Observation Engine is the first stage of Sentinel's execution
// pipeline. Every subsequent engine (Analysis, Decision, Planner,
// Backend, Telemetry, and Reconciliation) consumes the state
// produced here.
package observation

/*
Observation Engine Philosophy (Sentinel v1)

The Observation Engine provides a consistent and atomic view of the
current operating system state.

The engine is composed of small, focused collectors. Each collector
is responsible for observing exactly one subsystem of the operating
system.

Examples:

    - CPU
    - Memory
    - Network
    - Processes
    - Filesystem

Each collector has exactly one responsibility:

    Observe its subsystem and populate only its corresponding
    section of state.ObservationState.

Collectors NEVER:

    - Modify another collector's state.
    - Publish directly to the Shared State Repository.
    - Perform analysis or decision making.
    - Execute reconciliation actions.
    - Communicate with other collectors.

Instead, the Observation Manager coordinates every collector during
a single observation cycle.

                    Runtime Tick
                          │
                          ▼
                Observation Manager
                          │
        ┌─────────────────┼─────────────────┐
        ▼                 ▼                 ▼
       CPU             Memory           Network
        ▼                 ▼                 ▼
     Process         Filesystem      (future...)
        └─────────────────┼─────────────────┘
                          ▼
             state.ObservationState
                          ▼
               Repository.PublishObservation()
                          ▼
                         SSR

Observation Invariant

During every observation cycle, exactly one
state.ObservationState is created.

All collectors populate different sections of the same
ObservationState instance.

The completed ObservationState is published exactly once
to the Shared State Repository.

Partial observations are never published.

Engineering Principles

    - One observation cycle produces one ObservationState.
    - One allocation per observation cycle.
    - No intermediate snapshots.
    - No unnecessary memory copies.
    - No duplicate ownership of state.
    - Every collector owns exactly one subsystem.
    - The Observation Manager owns orchestration.
*/
