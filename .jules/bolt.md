## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.

## 2026-06-12 - Readdir synchronization overhead
**Learning:** Parallelizing simple CPU-bound operations in a loop (like parsing timestamps/sizes in `Readdir` after an API call returns) using goroutines, waitgroups, semaphores, and mutexes degrades performance by ~7x compared to sequential processing. The synchronization overhead drastically outweighs the parallelization benefits since the IO operation (API call) has already completed.
**Action:** Avoid goroutines for lightweight CPU-bound tasks in hot paths like `Readdir` unless significant blocking IO or computationally heavy work is involved. Sequential map insertions are far more efficient.