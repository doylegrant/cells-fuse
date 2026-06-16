## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.

## 2024-05-18 - FUSE Readdir Parallelization Overhead
**Learning:** Parallelizing lightweight, CPU-bound operations (like parsing basic strings or updating a map) within `Readdir` loop utilizing goroutines, a waitgroup, mutexes, and channels causes immense synchronization overhead. This approach is an anti-pattern when network latency or heavy I/O is not involved.
**Action:** Replaced the parallelization logic with a sequential loop, eliminating the scheduling and synchronization overhead, which speeds up `Readdir` by ~18x. Avoid parallelizing code where thread-management cost vastly outweighs the actual unit of work.
