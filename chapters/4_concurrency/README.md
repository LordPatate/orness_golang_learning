# Concurrency

## Goroutines

- A goroutine is a lightweight thread managed by the Go runtime: `go f(x, y, z)`
- The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens in the new goroutine.
- Goroutines run in the same address space, so access to shared memory must be synchronized.
- Synchronization can be accomplished using channels or the [sync](https://pkg.go.dev/sync) package.

## Channels

- Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`.
- By default, sends and receives block until the other side is ready.
- Channels can be *buffered*.
    - Sends to a buffered channel block only when the buffer is full.
    - Receives block when the buffer is empty.
- A sender can close a channel to indicate that no more values will be sent.
- Receivers can test whether a channel has been closed with the `v, ok := <-ch` idiom.
- The loop `for i := range c` receives values from the channel repeatedly until it is closed.
- The `select` statement lets a goroutine wait on multiple communication operations.
- A `select` blocks until one of its cases can run, then it executes that case.

## Exercise


