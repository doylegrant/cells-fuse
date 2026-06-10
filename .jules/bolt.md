## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.
## 2024-06-10 - String Concatenation vs. fmt.Sprintf in Hot Paths
**Learning:** Using `fmt.Sprintf` in high-frequency read hot paths introduces significant CPU overhead due to reflection and dynamic typing parsing.
**Action:** Replace `fmt.Sprintf("%s\x00%d", strVar, intVar)` with native string concatenation and `strconv.FormatInt(intVar, 10)`. In benchmarks, this reduces allocation latency by ~50% (from ~260ns to ~125ns).
