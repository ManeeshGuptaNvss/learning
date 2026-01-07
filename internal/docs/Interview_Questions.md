# Golang Technical Interview Preparation — Q&A Guide

This README summarizes our preparation conversation in a **question → answer → follow‑up** format. Use it for **quick revision before interviews**.

---

## 1. `make` vs `new`

### ❓ Question
What is the difference between `make` and `new` in Go?

### ✅ Answer
- `new(T)` allocates zeroed memory for type `T` and returns `*T`.
- `make(T, ...)` is only for `slice`, `map`, `chan`. It allocates and **initializes** internal data structures and returns a value of type `T`.

```go
p := new(int)        // *int
s := make([]int, 5)  // []int
m := make(map[string]int)
c := make(chan int)
```

### 🔎 Follow‑up: Stack vs Heap
Memory placement is decided by **escape analysis**, not by `make` or `new`.

```bash
go build -gcflags="-m"
```

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
- **Unbuffered** → send/receive block until both ready (synchronization)
- **Buffered** → blocks only when full/empty (queue behavior)

```go
make(chan int)     // unbuffered
make(chan int, 3)  // buffered
```

---

## 5. Closing Channels

### ❓ Question
What happens when you close a channel?

### ✅ Answer
- Send on closed → ❌ panic
- Receive from closed → zero value
- Close closed → ❌ panic

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
- `sync.Mutex`, `sync.RWMutex`
- channels
- `atomic`

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
- jobs channel
- workers
- results channel (optional)
- shutdown via close/context

Used to control concurrency and resource usage.

---

## 11. Mutex vs RWMutex

### ❓ Question
Difference between `sync.Mutex` and `sync.RWMutex`?

### ✅ Answer
- `Mutex` → one reader/writer
- `RWMutex` → many readers OR one writer

Use RWMutex only when reads are heavy.

---

## 12. `defer`

### ❓ Question
What is `defer`?

### ✅ Answer
Schedules a function to run when the surrounding function returns.

Rules:
- Executes in **LIFO order**
- Arguments evaluated immediately

Used for closing files, unlocking mutexes, panic recovery.

---

## 13. `context.Context`

### ❓ Question
What is context mainly used for?

### ✅ Answer
- Cancellation
- Timeouts/deadlines
- Request‑scoped data
- Stopping goroutines

---

# 🎯 High‑Impact Interview Lines

- "Stack vs heap is decided by escape analysis."
- "An interface is nil only when both type and value are nil."
- "Goroutines must cooperate to stop, usually via context."

---

# 🚀 How to Use This File

- Read once daily before interview
- Practice explaining each answer aloud
- Try coding small examples for channels and worker pools