## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.

## 2024-06-10 - Readdir Concurrency Overhead
**Learning:** Parallelizing lightweight CPU-bound operations (like parsing `TreeNode` to `fuse.Stat_t` in `Readdir` with `sync.Mutex`, `sync.WaitGroup`, and Goroutines) introduces massive synchronization and context-switching overhead that outweighs any benefits. This is a common codebase anti-pattern for rapid string processing or map assignments.
**Action:** Remove parallelization wrappers for fast map iterations and string handling to increase speed and lower memory usage.
