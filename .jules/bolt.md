## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.
## 2026-06-10 - Avoid fmt.Sprintf in hot paths
**Learning:** `fmt.Sprintf` creates significant heap allocations and CPU overhead in hot paths like the S3 prefetching logic. Using reflection-based string manipulation is not well suited for highly concurrent and heavily used filesystem operations like `Read`.
**Action:** Replace `fmt.Sprintf` with string concatenation combined with `strconv.FormatInt` (or `strconv.AppendInt` / `byte` slices for even fewer allocations) when generating cache keys or byte ranges to reduce memory allocations and improve overall throughput.
