## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.

## 2024-06-10 - Readdir synchronization overhead
**Learning:** In a FUSE filesystem, `Readdir` is a critical hot path. Parallelizing the processing of API payload nodes using goroutines, a `sync.WaitGroup`, and a `sync.Mutex` introduces significant synchronization overhead for primarily lightweight CPU-bound tasks (like parsing strings and updating a map). A micro-benchmark showed a 3x speed degradation when using parallelization compared to sequential processing.
**Action:** Avoid parallelizing lightweight map building or string parsing logic in hot paths; execute them sequentially to eliminate context switching and lock contention overhead.
