## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.

## 2024-06-10 - fmt.Sprintf allocation overhead in read hot paths
**Learning:** In the FUSE file read hot paths (`fuse.go` and `worker.go`), generating S3 byte ranges and cache keys via `fmt.Sprintf` creates unnecessary memory allocations and CPU cycles. Reflection-based string manipulations should be avoided inside loops executed for every file chunk. Benchmarks show that string concatenation using `+` and `strconv.FormatInt` is ~2x faster and produces fewer allocations.
**Action:** Replace `fmt.Sprintf` with string concatenation and `strconv` inside high-frequency operations, specifically when composing cache keys or HTTP header values.
