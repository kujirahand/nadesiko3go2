package nadesiko4

type ObjectStack struct {
	sp   int
	List []Object
}

func NewObjectStack() *ObjectStack {
	o := &ObjectStack{sp: 0}
	o.List = make([]Object, 0)
	return o
}

func (o *ObjectStack) Push(v Object) {
	o.List = append(o.List, v)
}

func (o *ObjectStack) Pop() Object {
	if o.sp <= 0 {
		return nil
	}
	v := o.List[o.sp-1]
	o.List = o.List[0 : o.sp-1]
	o.sp--
	return v
}

func (o *ObjectStack) Peek() Object {
	if o.sp <= 0 {
		return nil
	}
	v := o.List[o.sp-1]
	return v
}

func (o *ObjectStack) Length() int {
	return o.sp
}
