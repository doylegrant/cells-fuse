## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.

## 2024-06-09 - Goroutine overhead for simple sequential operations
**Learning:** In `fuse.go`'s `Readdir` function, parallelizing the parsing of relatively small, in-memory string properties (e.g., `strconv.ParseUint`, `filepath.Base`) inside a goroutine for every node introduced significant performance regressions (~3.8x slower) due to the overhead of goroutine scheduling, channel operations for bounding, and `sync.Mutex` contention on the output map.
**Action:** Always measure the impact of concurrency for operations that are purely CPU-bound. If the operation is cheap and doesn't involve I/O, a sequential loop without locks is usually faster.
