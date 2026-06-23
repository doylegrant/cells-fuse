## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.
## 2024-06-09 - FUSE fmt.Sprintf overhead in hot paths
**Learning:** In FUSE `Read` operations, creating strings like cache keys (`"%s\x00%d"`) and byte ranges (`"bytes=%d-%d"`) using `fmt.Sprintf` incurs significant reflection and allocation overhead (e.g., 235 ns/op vs 77 ns/op in benchmarks).
**Action:** Replace `fmt.Sprintf` with string concatenation and `strconv.FormatInt` in high-frequency FUSE paths to reduce allocations and improve throughput.
