## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.
## 2024-06-10 - fmt.Sprintf overhead in hot paths
**Learning:** `fmt.Sprintf` is a major performance bottleneck in hot paths like `Read` and `Prefetch` due to reflection and heap allocations. Benchmarks showed it taking ~150-230ns/op, while string concatenation combined with `strconv.FormatInt` reduces the time to ~33-125ns/op (about 2-4.5x faster).
**Action:** Always prefer string concatenation (`+`) with `strconv.FormatInt` over `fmt.Sprintf` for constructing strings in high-frequency FUSE operations.
