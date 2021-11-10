package main

import "fmt"

const NEG_INFINITY = -9999999

type MaxStack struct {
	size   int
	data   []int
	maxVal []int
}

func Constructor() MaxStack {
	s := MaxStack{}
	return s
}

func (this *MaxStack) Push(x int) {
	curr_max := this.PeekMax()

	if x >= curr_max {
		this.maxVal = append(this.maxVal, x)
	}

	this.data = append(this.data, x)
}

func (this *MaxStack) Pop() int {
	if len(this.data) <= 0 {
		return NEG_INFINITY
	}

	val := this.data[len(this.data)-1]
	if val == this.PeekMax() {
		this.maxVal = this.maxVal[:len(this.maxVal)-1]
	}
	this.data = this.data[:len(this.data)-1]
	return val
}

func (this *MaxStack) Top() int {
	if len(this.data) <= 0 {
		return NEG_INFINITY
	}
	fmt.Println(this.data)
	return this.data[len(this.data)-1]
}

func (this *MaxStack) PeekMax() int {
	if len(this.maxVal) == 0 {
		return NEG_INFINITY
	}
	return this.maxVal[len(this.maxVal)-1]
}

func (this *MaxStack) PopMax() int {
	max := this.PeekMax()
	buffer := Constructor()
	for this.Top() != max {
		buffer.data = append(buffer.data, this.Pop())
	}
	_ = this.Pop()

	for len(buffer.data) > 0 {
		this.Push(buffer.Pop())
	}
	return max
}

/**
 * Your MaxStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.PeekMax();
 * param_5 := obj.PopMax();
 */
