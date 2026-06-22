## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.
## 2024-06-10 - FUSE read/prefetch string formatting overhead
**Learning:** In a FUSE filesystem, `fmt.Sprintf` creates significant heap allocation and CPU cycle overhead when used in hot paths like `Read` or prefetch workers. This is because `fmt.Sprintf` relies on reflection, which is much slower than direct string manipulations.
**Action:** Replace `fmt.Sprintf` with string concatenation (`+`) and `strconv` functions (like `strconv.FormatInt`) for variable conversions in high-frequency execution paths to reduce overhead and improve read/prefetch throughput.
