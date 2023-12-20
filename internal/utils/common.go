package utils

func Cut(i int, xs []int) (int, []int) {
	y := xs[i]
	ys := append(xs[:i], xs[i+1:]...)
	return y, ys
}
