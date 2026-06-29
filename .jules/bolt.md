## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.

## 2024-06-29 - FUSE hot path allocation overhead from fmt.Sprintf
**Learning:** In a FUSE filesystem, `fmt.Sprintf` is often used for generating cache keys and byte ranges in the very hot read/prefetch paths. This creates significant garbage and CPU overhead (reflection and heap allocations). Benchmarks show that replacing `fmt.Sprintf` with native string concatenation (`+`) and `strconv.FormatInt` roughly halves the execution time for string generation in these paths.
**Action:** Replace `fmt.Sprintf` with string concatenation (`+`) and `strconv` functions for high-frequency string constructions (like cache keys or byte ranges) inside FUSE `Read` paths and worker pools.
