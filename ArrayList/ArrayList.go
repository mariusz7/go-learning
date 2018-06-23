package myarraylist

import (
	"errors"
	"fmt"
)

type IArrayList interface {
	PushFront(element interface{}) error
	PushBack(element interface{}) error
	Insert(index int, element interface{}) error

	PopFront() (interface{}, error)
	PopBack() (interface{}, error)
	Get(index int) (interface{}, error)

	IndexOf(element interface{}) (int, error)
	Contains(element interface{}) bool

	Remove(index int) error

	Size() int
	Clear()
	IsEmpty() bool
}

//TODO replace with "type MyArrayList []interface{}"
type MyArrayList struct {
	alist []interface{}
}

func (al *MyArrayList) getSliceForNewElements(newElements int) (newSlice bool , s []interface{}, e error) {

	//TODO handle overflowing int in the below addition (use e)
	var suffCap int = len(al.alist) + newElements

	if suffCap > cap(al.alist) {

		var newCap int
		if cap(al.alist) <= 0 {
			newCap = 1
		} else {
			newCap = cap(al.alist)
		}

		//TODO handle overflowing int by shifting (use e)
		for ; newCap < suffCap; newCap <<= 1 {
		}

		newSlice = true
		s = make([]interface{}, 0, newCap)
		e = nil
		fmt.Printf("old cap: %v, new cap: %v\n", cap(al.alist), cap(s))

	} else {
		newSlice = false
		s = al.alist
		e = nil
	}
	return
}

func (al *MyArrayList) PushFront(element interface{}) error {
	if newSlice, s, e := al.getSliceForNewElements(1); e == nil {

		//None 'append()' call below should reallocate

		if newSlice {
			s = append(s, element) 
			s = append(s, al.alist...)
			al.alist = s
			fmt.Printf("len(al.alist): %v +PF\n", len(al.alist))
			return nil
		} else {
			var appendLast bool = false
			var last interface{}
			if len(al.alist) > 0 {
				last = al.alist[len(al.alist) - 1]
				appendLast = true	
			}

			copy(al.alist[1:], al.alist[:len(al.alist)-1])
			al.alist[0] = element
			if appendLast {
				al.alist = append(al.alist, last)
			}

			fmt.Printf("len(al.alist): %v 0PF\n", len(al.alist))
			return nil
		}
	} else {
		return errors.New("Not enough space")
	}
}

func (al *MyArrayList) PushBack(element interface{}) error {

	if newSlice, s, e := al.getSliceForNewElements(1); e == nil {
		if newSlice {
			s = append(s, al.alist...)
		}
		s = append(s, element) //Should not reallocate
		al.alist = s
		fmt.Printf("len(al.alist): %v PB\n", len(al.alist))
		return nil
	} else {
		return errors.New("Not enough space")
	}
}

func (al *MyArrayList) Insert(index int, element interface{}) error {

	if index < 0 || index > len(al.alist) {
		return errors.New("Out of bound input")
	}

	//None of below 'append()' calls should reallocate

	if newSlice, s, e := al.getSliceForNewElements(1); e == nil {

		if len(al.alist) <= 0 { //Insert into empty
			s = append(s, element)
			al.alist = s
			return nil
		}

		if newSlice { //We need to copy everything
			var before []interface{} = al.alist[:index]
			var after []interface{} = al.alist[index:]

			if len(before) > 0 {
				s = append(s, before...)
			}

			s = append(s, element)

			if len(after) > 0 {
				s = append(s, after...)
			}

			al.alist = s
			return nil
		} else {
			if index == len(al.alist) {
				al.alist = append(al.alist, element)
				return nil
			} else {
				after := al.alist[index:]
				var appendLast bool = false
				var last interface{}
				if len(after) > 0 {
					last = after[len(after) - 1]
					appendLast = true	
				}
				
				copy(al.alist[index + 1:], after)
				if appendLast {
					al.alist = append(al.alist, last)
				}
				al.alist[index] = element
				return nil
			}
		}
	} else {
		return errors.New("Not enough space")
	}
}

func (al *MyArrayList) PopFront() (interface{}, error) {
	if len(al.alist) <= 0 {
		return nil, errors.New("Container is empty")
	}

	return al.alist[0], nil
}


func (al *MyArrayList) PopBack() (interface{}, error) {
	if len(al.alist) <= 0 {
		return nil, errors.New("Container is empty")
	}

	return al.alist[len(al.alist) - 1], nil
}


func (al *MyArrayList) Get(index int) (interface{}, error) {
	if len(al.alist) <= 0 {
		return nil, errors.New("Container is empty")
	}

	if index < 0 || index >= len(al.alist) {
		return nil, errors.New("Out of bound input index")
	}

	return al.alist[index], nil
}

func (al *MyArrayList) IndexOf(element interface{}) (int, error) {
	for i, v := range al.alist {
		if v == element {
			return i, nil
		}
	}

	return -1, errors.New("Not found")
}

func (al *MyArrayList) Contains(element interface{}) bool {
	for _, v := range al.alist {
		if v == element {
			return true
		}
	}

	return false
}

func (al *MyArrayList) Remove(index int) error {

	if index < 0 || index >= len(al.alist) {
		return errors.New("Out of bound input index")
	}

	if len(al.alist) == 1 {
		al.alist = make([]interface{}, 0)
	} else {
		var before []interface{} = al.alist[:index]
		var after []interface{}

		if index < len(al.alist) - 1 {
			after = al.alist[index + 1:]
		} else {
			after = make([]interface{}, 0)
		}

		al.alist = append(before, after...)
	}

	return nil
}

func (al *MyArrayList) Size() int {
	return len(al.alist)
}

func (al *MyArrayList) Clear() {
	al.alist = make([]interface{}, 0, 1)
}

func (al *MyArrayList) IsEmpty() bool {
	return len(al.alist) <= 0
} 

