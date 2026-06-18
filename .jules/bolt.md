## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.
## 2024-06-10 - FUSE path reflection overhead
**Learning:** `fmt.Sprintf` creates significant heap allocations and performance overhead (~2x slower) when used on hot code paths, such as constructing cache keys and byte ranges in `fuse.go` and `worker.go`.
**Action:** Replaced `fmt.Sprintf` with string concatenations and `strconv.FormatInt` to minimize reflection, reduce garbage collection pressure, and improve speed.
