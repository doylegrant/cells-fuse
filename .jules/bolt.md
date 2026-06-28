## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.

## 2026-06-28 - Avoid fmt.Sprintf in FUSE hot paths
**Learning:** In highly trafficked paths like FUSE Read operations and prefetch loops, using `fmt.Sprintf` for constructing strings (e.g. cache keys or HTTP byte range headers) introduces unnecessary CPU and allocation overhead compared to simple string concatenation. Benchmark tests show `fmt.Sprintf` taking ~2x-3x longer than concatenating strings with `strconv.FormatInt`.
**Action:** Replace `fmt.Sprintf` with string concatenation combined with `strconv.FormatInt` in file read/prefetch hot paths.
