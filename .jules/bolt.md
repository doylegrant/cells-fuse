## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.

## 2026-06-17 - Avoid fmt.Sprintf in FUSE hot paths
**Learning:** `fmt.Sprintf` is used heavily in the `CellsFuse.Read()` method for constructing cache keys and byte range strings for S3 requests. FUSE reads are very frequent, and the reflection used in `fmt.Sprintf` adds measurable overhead per iteration. A benchmark showed that avoiding `fmt.Sprintf` reduces time per operation from ~165ns to ~80ns (a ~2x improvement) and completely eliminates heap allocations caused by reflection.
**Action:** Replace `fmt.Sprintf` with string concatenation and `strconv.FormatInt` inside highly repeated FUSE operations like reading and writing chunks.
