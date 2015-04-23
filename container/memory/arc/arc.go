package arc

import (
	"github.com/wayn3h0/go-caching"
	"github.com/wayn3h0/go-caching/container/memory"
)

// New returns a new instance of caching.Container: in-memory caching container using arc (adaptive/adjustable replacement cache) arithmetic.
func New(capacity int) caching.Container {
	return &container{
		capacity: capacity,
		t1:       newItems(),
		t2:       newItems(),
		b1:       newItems(),
		b2:       newItems(),
	}
}

// register the container.
func init() {
	memory.ARC.Register(New)
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

type container struct {
	capacity int
	p        int // target size of t1
	t1       *items
	t2       *items
	b1       *items
	b2       *items
}

func (this *container) replace() {
	if this.t1.Count() >= max(1, this.p) { // t1's size exceeds target (t1 is too big)
		this.b1.Set(this.t1.Discard()) // grab from t1 and put to b1
	} else {
		this.b2.Set(this.t2.Discard()) // grab from t2 and put to b2
	}
}

func (this *container) Get(key string) (interface{}, error) {
	if this.t1 == nil {
		return nil, nil
	}

	if this.t1.Contains(key) { // seen twice recently, put it to t2
		value := this.t1.Remove(key)
		this.t2.Set(key, value)
		return value, nil
	}

	if this.t2.Contains(key) {
		return this.t2.Get(key), nil
	}

	if this.b1.Contains(key) {
		this.p = min(this.capacity, this.p+max(this.b2.Count()/this.b1.Count(), 1)) // adapt the target size of t1
		this.replace()
		value := this.b1.Remove(key)
		this.t2.Set(key, value) // seen twice recently, put it to t2
		return value, nil
	}

	if this.b2.Contains(key) {
		this.p = max(0, this.p-max(this.b1.Count()/this.b2.Count(), 1)) // adapt the target size of t1
		this.replace()
		value := this.b2.Remove(key)
		this.t2.Set(key, value) // seen twice recently, put it to t2
		return value, nil
	}

	return nil, nil
}

func (this *container) Set(key string, value interface{}) error {
	if this.t1 == nil {
		this.t1 = newItems()
		this.t2 = newItems()
		this.b1 = newItems()
		this.b2 = newItems()
	}

	// remove the item if exists
	this.Remove(key)

	if value == nil {
		return this.Remove(key)
	}

	if this.t1.Count()+this.b1.Count() == this.capacity { // b1 + t1 is full
		if this.t1.Count() < this.capacity { // still room in t1
			this.b1.Discard()
			this.replace()
		} else {
			this.t1.Discard()
		}
	} else { //this.t1.Count()+this.b1.Count() < this.capacity {
		total := this.t1.Count() + this.t2.Count() + this.b1.Count() + this.b2.Count()
		if total >= this.capacity { // cache full
			if total == 2*this.capacity {
				this.b2.Discard()
			}

			this.replace()
		}
	}

	this.t1.Set(key, value) // seen once recently, put on t1

	return nil
}

func (this *container) Remove(key string) error {
	this.t1.Remove(key)
	this.t2.Remove(key)
	this.b1.Remove(key)
	this.b2.Remove(key)

	return nil
}

func (this *container) Clear() error {
	if this.t1 == nil {
		return nil
	}

	this.p = 0
	this.t1 = newItems()
	this.t2 = newItems()
	this.b1 = newItems()
	this.b2 = newItems()

	return nil
}
