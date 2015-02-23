package goutils

func CopyMap(src, dst map[int]bool) {
	for k, v := range src {
		dst[k] = v
	}
}
