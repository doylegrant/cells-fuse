## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.
## 2024-07-01 - Synchronization overhead in FUSE readdir
**Learning:** Parallelizing lightweight CPU-bound operations (like simple string manipulations, map lookups, and struct creations) in the `Readdir` loop causes significant performance degradation due to synchronization overhead (`sync.WaitGroup`, `sync.Mutex`, goroutine scheduling, channel semantics) on a hot FUSE path.
**Action:** Always prefer sequential execution for lightweight loop iterations in hot paths, avoiding goroutines and synchronization primitives unless the workload per item is heavily CPU-bound or blocking (I/O).
