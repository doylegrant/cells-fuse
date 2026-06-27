## 2024-06-09 - FUSE path conversion allocation overhead
**Learning:** In a FUSE filesystem, `toInternalPath` is called for almost every filesystem operation (via `beginOp` and explicitly). Using `strings.Split` and `strings.Join` for simple path modifications (handling workspace aliases and `.recycle_bin`) creates significant garbage and CPU overhead (9 allocations per call).
**Action:** Replace slice-based string manipulation with string slicing and `strings.ReplaceAll` for `.recycle_bin` to minimize allocations in the hot path.

## 2024-06-09 - Avoid fmt.Sprintf in FUSE Fs/Worker hot paths
**Learning:** `fmt.Sprintf` uses reflection which incurs a notable CPU and heap allocation penalty (up to 40% slower) compared to simple string concatenations like `string + string + strconv.FormatInt`. This becomes a significant overhead in very hot paths like FUSE read/prefetch operations which execute continuously.
**Action:** Replace `fmt.Sprintf` with `strconv.FormatInt` and string concatenation for cache keys and chunk byte range strings across FUSE read and background worker prefetching paths.
