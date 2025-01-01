package worker

func NewPool(poolSize int) Pool {
	return Pool{
		size: poolSize,
	}
}
