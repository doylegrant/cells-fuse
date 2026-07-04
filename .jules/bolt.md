## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.

## 2025-02-23 - FUSE Readdir parallelization overhead
**Learning:** Parallelizing lightweight CPU-bound operations (like parsing strings or map updates in loops) using goroutines, waitgroups, and mutexes causes significant performance degradation compared to sequential execution due to synchronization overhead.
**Action:** Avoid parallelization for lightweight loop operations in hot paths; use sequential execution instead.
