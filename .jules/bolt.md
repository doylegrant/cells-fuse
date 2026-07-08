## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.

## 2024-07-08 - FUSE hot path allocation overhead from fmt.Sprintf
**Learning:** In a FUSE filesystem, functions like `Read` and worker prefetching execute on extremely hot paths. Using reflection-based string manipulations like `fmt.Sprintf` for generating cache keys and byte ranges causes significant memory allocations and CPU overhead compared to string concatenation with `strconv.FormatInt`. Micro-benchmarks showed a ~35-45% reduction in execution time per operation.
**Action:** Avoid `fmt.Sprintf` in hot paths; prefer string concatenation (`+`) with `strconv.FormatInt` or byte buffer appending instead to minimize heap allocations and CPU cycles.
