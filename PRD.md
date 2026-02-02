# Product Requirements Document (PRD)

## Vibly

**Local Live Streaming Platform**

## Author

Shahadath Hossen Sajib

## Purpose & Vision

Build a **real-engineering-focused live streaming platform** similar to Kick/Twitch, designed primarily as a **portfolio project**. The goal is not commercial scale but to demonstrate **backend systems thinking**, **media pipeline understanding**, and **scalability concepts** using **Go**, **FFmpeg**, and a **monolithic but scalable architecture**.

The platform will support:

- Live streaming ingestion
- Multi-quality adaptive playback
- Recording streams locally
- Thumbnail generation
- Concurrent stream handling with worker isolation
- **User as Channel**: Each user is a unique channel
- **Persistent Channel Chat**: Real-time and historical chat for each channel

---

## Target Audience

- Recruiters & senior backend engineers
- DevOps-curious backend roles
- Systems/backend interviews

This project intentionally prioritizes **architecture clarity over UI polish**.

---

## Non-Goals

- Global CDN distribution
- Paid subscriptions or monetization
- Chat moderation (Phase 1)
- Ultra-low-latency (WebRTC)

---

## High-Level Architecture

```
Streamer (OBS)
    ↓ RTMP
Ingest Server (Go)
    ↓ Job Creation
Job Queue
    ↓
FFmpeg Workers (Isolated Processes)
    ↓
HLS Output (Multi-bitrate)
    ↓
Video Player
```

Monolith at the API level, **distributed at the worker level**.

---

## Tech Stack

### Backend

- **Go (raw net/http or minimal router)**
- FFmpeg (external binary)
- Redis (job queue, optional but recommended)
- SQLite / PostgreSQL (metadata)

### Media

- RTMP ingest
- HLS output
- Multi-bitrate adaptive streaming

### Infrastructure (Local-first)

- Docker & Docker Compose
- Linux-based OS process management

---

## Core Functional Requirements

### 1. Live Stream Ingestion

**Description**

- Accept live streams via RTMP (OBS compatible)
- Each stream is uniquely identified

**Requirements**

- Stream key validation
- Reject streams if system is overloaded

**Success Criteria**

- OBS can successfully start streaming
- Backend acknowledges stream start

---

### 2. Stream Recording

**Description**

- All live streams must be recorded locally

**Requirements**

- Save full stream as a file (e.g., MP4)
- Recording starts automatically with live stream

**Notes**

- Recording happens inside the FFmpeg worker

---

### 3. Multi-Quality Streaming (Adaptive Bitrate)

**Description**

- Each stream is transcoded into multiple qualities

**Target Profiles**

- 1080p
- 720p
- 480p

**Technical Approach**

- Single FFmpeg process per stream
- Multiple output variants
- HLS master playlist generation

**Playback**

- Player auto-switches quality based on bandwidth

---

### 4. Thumbnail Generation

**Description**

- Automatically generate stream thumbnail

**Requirements**

- Extract frame using FFmpeg
- Thumbnail available before playback page loads

---

### 5. Worker System (Critical Requirement)

**Philosophy**
Workers are **not magical**. They are **real OS processes** that must be controlled.

**Worker Definition**

- One worker = one FFmpeg process handling one stream

**Worker Pool Model**

- Fixed maximum workers (configurable)
- Workers are assigned jobs from queue
- When all workers are busy, new streams are queued or rejected

**Key Guarantees**

- One stream cannot crash the whole server
- CPU usage remains bounded

---

### 6. Job Queue

**Description**

- Decouple stream ingestion from FFmpeg execution

**Requirements**

- FIFO job handling
- Stream start → job creation
- Stream end → job cleanup

**Implementation Options**

- In-memory queue (Phase 1)
- Redis-backed queue (Phase 2)

---

### 7. Concurrency & Load Handling

**Constraints**

- Multiple concurrent streams supported
- Each stream isolated in its own worker

**Failure Handling**

- Worker crash does not affect others
- API remains responsive

---

### 9. Channel & Chat System

**Description**

- Each user has a dedicated "Channel"
- Persistent chat room for each channel

**Requirements**

- Users can send messages to a channel's chat
- Chat history is stored and retrievable
- Real-time updates via WebSockets

**Success Criteria**

- Chat persists even if the stream ends
- New viewers can see previous chat history

### 8. API Layer

**Responsibilities**

- Stream lifecycle management
- Worker assignment
- Metadata persistence

**Explicitly NOT Responsible For**

- Video processing
- Heavy computation

---

## Scalability Strategy (Conceptual)

### Phase 1 — Local Scaling

- Fixed worker pool
- Manual limits

### Phase 2 — Simulated Autoscaling

- Spawn FFmpeg processes dynamically
- Enforce max worker count

### Phase 3 — Real Autoscaling (Optional)

- Kubernetes HPA
- Worker pods

> The application never creates machines. Infrastructure does.

---

## Data Model (Simplified)

### User / Channel

- id
- name
- email
- password_hash
- stream_key
- created_at

### Stream

- id
- channel_id
- status (live / ended)
- started_at
- ended_at

### Chat Message

- id
- channel_id
- user_id
- content
- created_at

### Recording

- stream_id
- file_path
- duration

### Worker

- id
- status (idle / busy)
- current_stream_id

---

## Security Considerations

- Stream key validation
- Path traversal prevention
- FFmpeg execution sandboxing

---

## Observability

**Metrics**

- Active streams
- Busy workers
- Queue length

**Logs**

- Stream lifecycle events
- FFmpeg failures

---

## Risks & Mitigations

| Risk         | Mitigation        |
| ------------ | ----------------- |
| CPU overload | Worker limit      |
| FFmpeg crash | Process isolation |
| Disk fill    | Storage quotas    |
| Stream abuse | Rate limiting     |

---

## Success Definition (Portfolio)

This project is successful if:

- It handles multiple live streams safely
- Quality switching works during playback
- Architecture can be clearly explained in interviews
- Demonstrates real backend & systems thinking

---

## Why This Project Matters

This project proves:

- Understanding of media pipelines
- Experience with concurrency & isolation
- Ability to design scalable systems without overengineering

**This is real backend engineering — not a toy CRUD app.**
