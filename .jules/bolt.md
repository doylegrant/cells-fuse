## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.
## 2024-06-09 - Reflection-based string formatting in hot paths
**Learning:** `fmt.Sprintf` uses reflection to determine types at runtime and typically escapes variables to the heap, causing significant CPU overhead and memory allocations. When used in heavily executed hot paths like FUSE read operations or cache key generation loops, this becomes a measurable bottleneck.
**Action:** Replace `fmt.Sprintf` with simple string concatenation combined with functions like `strconv.FormatInt` or `strconv.AppendInt` in performance-critical sections to eliminate reflection overhead and minimize allocations.
