package queue

type Queue struct {
	Items []int
}

func (q *Queue) AddRight(items ...int) {
	for _, v := range items {
		q.Items = append(q.Items, v)
	}
}

func (q *Queue) AddLeft(items ...int) {
	l := items
	q.Items = append(l, q.Items...)
}

func (q *Queue) PopLeft() int {
	if len(q.Items) == 0 {
		return -1
	}
	item := q.Items[0]
	q.Items = q.Items[1:]
	return item
}

func (q *Queue) PopRight() int {
	if len(q.Items) == 0 {
		return -1
	}
	index := len(q.Items) - 1
	item := q.Items[index]
	q.Items = q.Items[:index]
	return item
}

func (q *Queue) Length() int {
	return len(q.Items)
}
