## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.
## 2024-06-09 - Reflection overhead in hot paths
**Learning:** `fmt.Sprintf` uses reflection and allocates memory dynamically, which becomes a significant bottleneck in frequently executed code paths like FUSE read operations and worker prefetch loops.
**Action:** Replace `fmt.Sprintf` with string concatenation (`+`) and specific formatting functions like `strconv.FormatInt` for measurable performance gains without sacrificing readability.
