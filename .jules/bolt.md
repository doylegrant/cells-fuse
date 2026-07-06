## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.
## 2024-06-09 - Readdir parallelization overhead
**Learning:** Parallelizing lightweight CPU-bound operations like string parsing and map updates in a FUSE directory listing (`Readdir`) with goroutines, WaitGroups, and Mutexes introduces severe synchronization overhead. This causes increased execution time and large memory allocations compared to a simple sequential loop.
**Action:** Always process lightweight data sequentially in hot paths. Avoid throwing goroutines at a problem unless the operations involve significant I/O or heavy computation.
