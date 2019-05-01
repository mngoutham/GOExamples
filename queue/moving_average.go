package queue

type MovingAverage struct {
	Queue []int
	Size  int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{Size: size}
}

func (this *MovingAverage) Next(val int) float64 {
	if this.Size == 0 {
		return 0
	}
	if len(this.Queue) >= this.Size {
		this.Queue = this.Queue[1:len(this.Queue)]
	}
	this.Queue = append(this.Queue, val)
	sum := 0
	for _, v := range this.Queue {
		sum += v
	}
	return float64(sum) / float64(len(this.Queue))
}
