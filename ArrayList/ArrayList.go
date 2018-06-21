package myarraylist

import (
	"errors"
)

type IArrayList interface {
	PushFront(element interface{})
	PushBack(element interface{})
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

type MyArrayList struct {
	arrayList []interface{}
}

func (ali *MyArrayList) PushFront(element interface{}) {
	ali.arrayList = append([]interface{}{element}, ali.arrayList)
}

func (ali *MyArrayList) PushBack(element interface{}) {
	ali.arrayList = append(ali.arrayList, element)
}










func (ali *MyArrayList) Insert(index int, element interface{}) error {
	if index > len(ali.arrayList) {
		return errors.New("Out of bound input")
	}

	ali.arrayList = append(
		ali.arrayList[:index],
		element,
		ali.arrayList[index + 1:])

	return nil
}

func (ali *MyArrayList) PopFront() (interface{}, error) {
	if len(ali.arrayList) <= 0 {
		return nil, errors.New(
"Attempt to read from an empty container")
	}

	return ali.arrayList[0], nil
}


func (ali *MyArrayList) PopBack() (interface{}, error) {
	if len(ali.arrayList) <= 0 {
		return nil, errors.New(
"Attempt to read from an empty container")
	}

	return ali.arrayList[len(ali.arrayList) - 1], nil
}


func (ali *MyArrayList) Get(index int) (interface{}, error) {
	if len(ali.arrayList) <= 0 {
		return nil, errors.New(
"Attempt to read from an empty container")
	}

	if index < 0 || index >= len(ali.arrayList) {
		return nil, errors.New(
"Out of bound input index")
	}

	return ali.arrayList[index], nil
}



func (ali *MyArrayList) IndexOf(element interface{}) (int, error) {
	for i, v := range ali.arrayList {
		if v == element {
			return i, nil
		}
	}

	return -1, errors.New("Not found")
}


func (ali *MyArrayList) Contains(element interface{}) bool {
	for _, v := range ali.arrayList {
		if v == element {
			return true
		}
	}

	return false
}

func (ali *MyArrayList) Remove(index int) error {
	if index < 0 || index >= len(ali.arrayList) {
		return errors.New("Out of bound input index")
	}

	if len(ali.arrayList) == 1 {
		ali.arrayList = make([]interface{}, 0)
	} else {
		ali.arrayList = ali.arrayList[:index]

		if index < len(ali.arrayList) - 1 {
			ali.arrayList = append(ali.arrayList, ali.arrayList[index + 1:])
		}
	}

	return nil
}


func (ali *MyArrayList) Size() int {
	return len(ali.arrayList)
}


func (ali *MyArrayList) Clear() {
	ali.arrayList = make([]interface{}, 0)
}


func (ali *MyArrayList) IsEmpty() bool {
	return len(ali.arrayList) <= 0
}
