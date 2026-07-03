## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.

## 2025-02-12 - Readdir parallelization overhead
**Learning:** In `Readdir`, parallelizing the process of mapping `TreeNode` objects to FUSE `Stat_t` structures (which mainly involves string parsing and basic assignments) using goroutines, a `sync.WaitGroup`, a mutex, and a semaphore channel introduced significant overhead. Benchmarking showed sequential processing is ~5x faster because the synchronization cost far outweighs the lightweight CPU-bound work per node.
**Action:** Avoid parallelizing lightweight CPU-bound map updates and string parsing logic. Use sequential loops to minimize synchronization and allocation overhead.
