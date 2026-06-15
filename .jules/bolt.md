## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.
## 2024-06-10 - fmt.Sprintf overhead on FUSE hot paths
**Learning:** `fmt.Sprintf` uses reflection and allocates more memory compared to simple string concatenation and integer conversion via `strconv`. In hot paths like file reading and prefetching (e.g., generating cache keys and HTTP byte ranges), this introduces unnecessary CPU cycles and garbage collection overhead. Micro-benchmarks confirmed a 5x improvement for cache key generation and 2x for byte ranges.
**Action:** Replace `fmt.Sprintf` with string concatenation (`+`) and `strconv.FormatInt` on high-frequency paths.
