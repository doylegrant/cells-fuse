## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.

## 2024-10-27 - FUSE read path allocation overhead
**Learning:** In the FUSE filesystem, generating `cacheKey` and `byteRange` strings with `fmt.Sprintf` in hot loops like read-ahead fetching (`Read`) and prefetch queueing causes significant garbage collection pressure due to reflection allocations, scaling negatively with chunking operations. Benchmarks indicate concatenating strings with `strconv.FormatInt` is up to 3x faster and reduces allocations by half.
**Action:** Avoid `fmt.Sprintf` in hot paths involving iteration or FUSE operations. Use direct string concatenation combined with `strconv.FormatInt` (or similar `strconv` functions) for constructing identifiers and headers in high-throughput areas.
