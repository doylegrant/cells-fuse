## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.
## 2024-06-10 - Avoid fmt.Sprintf in FUSE hot paths
**Learning:** Using `fmt.Sprintf` for constructing cache keys and byte ranges in the `Read` hot path causes unnecessary heap allocations and CPU overhead due to reflection. FUSE filesystem operations require extremely low latency.
**Action:** Always prefer direct string concatenation (`+`) with `strconv.FormatInt` or similar functions in performance-critical paths instead of reflection-based string manipulation.
