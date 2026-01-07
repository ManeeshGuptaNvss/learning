# Golang Technical Interview Preparation — Q&A Guide

This README summarizes our preparation conversation in a **question → answer → follow‑up** format. Use it for **quick revision before interviews**.

---

## 1. `make` vs `new`

### ❓ Question

What is the difference between `make` and `new` in Go?

### ✅ Answer

* `new(T)` allocates zeroed memory for type `T` and returns `*T`.
* `make(T, ...)` is only for `slice`, `map`, `chan`. It allocates and **initializes** internal data structures and returns a value of type `T`.

```go
p := new(int)        // *int
s := make([]int, 5)  // []int
m := make(map[string]int)
c := make(chan int)
```

### 🔎 Follow‑up

Blocking a goroutine does **not** block the OS thread.

**Answer:**
When a goroutine blocks (on I/O, channel, mutex, etc.), the Go runtime **parks** it and schedules another runnable goroutine on the same OS thread. If the thread is blocked in a syscall, Go may create or reuse another OS thread so other goroutines can continue running.

Interview line:

> “Blocking a goroutine does not block the OS thread; the scheduler runs another goroutine instead.”

---

## 2. Interfaces in Go

### ❓ Question

What is an interface in Go?

### ✅ Answer

An interface defines a **set of method signatures**. Any type that implements those methods **implicitly** satisfies the interface.

```go
type Reader interface { Read([]byte) (int, error) }
```

Internally, an interface is stored as:

```
(type, value)
```

---

## 3. Nil Interface Trap

### ❓ Question

Difference between a nil interface and an interface holding a nil concrete value?

### ✅ Answer

An interface is **nil only when both type and value are nil**.

```go
var r io.Reader
fmt.Println(r == nil) // true

var f *os.File = nil
r = f
fmt.Println(r == nil) // false
```

### 🎯 Interview line

"An interface is nil only when both its dynamic type and value are nil."

---

## 4. Channels

### ❓ Question

Difference between buffered and unbuffered channels?

### ✅ Answer

* **Unbuffered** → send/receive block until both ready (synchronization)
* **Buffered** → blocks only when full/empty (queue behavior)

```go
make(chan int)     // unbuffered
make(chan int, 3)  // buffered
```

---

## 5. Closing Channels

### ❓ Question

What happens when you close a channel?

### ✅ Answer

* Send on closed → ❌ panic
* Receive from closed → zero value
* Close closed → ❌ panic

```go
v, ok := <-ch // ok = false when closed and empty
```

Only the **sender** should close a channel.

---

## 6. Goroutines

### ❓ Question

What is a goroutine? How is it different from an OS thread?

### ✅ Answer

A goroutine is a **lightweight function execution unit** managed by the Go runtime. Goroutines are multiplexed onto OS threads (M:N scheduler), have very small growing stacks, and are cheap to create.

### 🔎 Follow‑up

Blocking a goroutine does **not** block the OS thread. The scheduler parks it and runs another.

---

## 7. `select`

### ❓ Question

What is `select` used for?

### ✅ Answer

`select` lets a goroutine wait on **multiple channel operations**.

```go
select {
case v := <-ch:
case <-time.After(time.Second):
case <-ctx.Done():
}
```

Used for: multiplexing, timeouts, cancellation.

---

## 8. Graceful Goroutine Stop

### ❓ Question

How do you stop a goroutine gracefully?

### ✅ Answer

Goroutines cannot be force‑stopped. They must **listen for a signal**.

Best practice: `context.Context`

```go
select {
case <-ctx.Done():
    return
}
```

---

## 9. Race Conditions

### ❓ Question

What is a race condition and how do you prevent it?

### ✅ Answer

A race occurs when multiple goroutines access shared data concurrently and at least one is a write without synchronization.

Prevent using:

* `sync.Mutex`, `sync.RWMutex`
* channels
* `atomic`

Detect with:

```bash
go test -race
```

---

## 10. Worker Pool Pattern

### ❓ Question

What is a worker pool?

### ✅ Answer

A fixed number of goroutines (workers) pulling tasks from a jobs channel.

Components:

* jobs channel
* workers
* results channel (optional)
* shutdown via close/context

Used to control concurrency and resource usage.

---

## 11. Mutex vs RWMutex

### ❓ Question

Difference between `sync.Mutex` and `sync.RWMutex`?

### ✅ Answer

* `Mutex` → one reader/writer
* `RWMutex` → many readers OR one writer

Use RWMutex only when reads are heavy.

---

## 12. `defer`

### ❓ Question

What is `defer`?

### ✅ Answer

Schedules a function to run when the surrounding function returns.

Rules:

* Executes in **LIFO order**
* Arguments evaluated immediately

Used for closing files, unlocking mutexes, panic recovery.

---

## 13. `context.Context`

### ❓ Question

What is context mainly used for?

### ✅ Answer

* Cancellation
* Timeouts/deadlines
* Request‑scoped data
* Stopping goroutines

---

# 🎯 High‑Impact Interview Lines

* "Stack vs heap is decided by escape analysis."
* "An interface is nil only when both type and value are nil."
* "Goroutines must cooperate to stop, usually via context."

---

# 🚀 How to Use This File

* Read once daily before interview
* Practice explaining each answer aloud
* Try coding small examples for channels and worker pools

---

# ⚠️ Common Interview Traps & Trick Questions

1. **Nil interface trap**

   ```go
   var r io.Reader = (*os.File)(nil)
   fmt.Println(r == nil) // false
   ```

   Trap: Interface is non-nil because it has a dynamic type.

2. **Loop variable capture**

   ```go
   for i := 0; i < 3; i++ {
       go func() { fmt.Println(i) }()
   }
   ```

   Fix:

   ```go
   for i := 0; i < 3; i++ {
       i := i
       go func() { fmt.Println(i) }()
   }
   ```

3. **Goroutine leak**
   Goroutines blocked forever on channels or waiting without cancellation.
   Fix: always use `context`, done channels, or proper closing.

4. **Closing the wrong channel**
   Only the **sender** should close a channel. Receivers should never close.

5. **Send on closed channel** → panic

6. **Maps are not thread-safe**
   Must protect with mutex or use `sync.Map`.

7. **Assuming `make` means heap**
   Stack vs heap is decided by escape analysis.

---

# ✅ Last‑Day Revision Checklist

Before your interview, make sure you can confidently explain:

* [ ] `make` vs `new` + escape analysis
* [ ] Interface internals and nil interface trap
* [ ] Goroutines vs OS threads
* [ ] Buffered vs unbuffered channels
* [ ] Closing channels behavior
* [ ] `select` with timeout and cancellation
* [ ] Graceful goroutine shutdown using `context`
* [ ] Race conditions + prevention + race detector
* [ ] Mutex vs RWMutex
* [ ] Worker pool pattern
* [ ] `defer` rules (LIFO, argument evaluation)

Also revise:

* [ ] Writing clean Go code (error handling, naming)
* [ ] Basic DSA in Go (arrays, maps, trees)
* [ ] One concurrency coding problem

---

# 🎤 Mock Interview Q&A Section

These are extremely common questions interviewers ask after basics:

## Q1. How would you design a rate limiter in Go?

Expected keywords:

* token bucket / leaky bucket
* time.Ticker
* buffered channel
* goroutines

---

## Q2. How do you implement a thread-safe map?

Expected keywords:

* `sync.Mutex` + map
* `sync.RWMutex`
* `sync.Map` (when reads >> writes)

---

## Q3. How do you ensure a server shuts down gracefully?

Expected keywords:

* `context.WithCancel`
* `http.Server.Shutdown`
* wait groups
* stop accepting new work

---

## Q4. What happens if many goroutines write to the same channel?

Expected keywords:

* channels are thread-safe
* ordering is not guaranteed
* race only occurs if shared memory is modified unsafely

---

## Q5. How would you debug high memory usage in Go?

Expected keywords:

* `pprof`
* GC behavior
* escape analysis
* object reuse (`sync.Pool`)

---

## Q6. How would you design a concurrent job processing system?

Expected keywords:

* worker pool
* backpressure
* buffered channels
* graceful shutdown
* retries

---

# 🏁 Final Interview Advice

* Speak in **structured answers** (definition → how it works → example → trade‑offs)
* Always mention **context, cancellation, and graceful shutdown** in backend discussions
* If stuck, explain your **thought process** clearly

Good luck — you now have a solid Go interview preparation guide. 💪