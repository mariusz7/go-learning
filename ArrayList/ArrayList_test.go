package myarraylist

import "testing"
/*
func TestMyArrayList(t *testing.T) {

	var ali IArrayList = &MyArrayList{}

	var inserted [16]int = [16]int{
		43, 22, 1, 77, -21, 33, 54, -84,
		19, -52, 5, 1, 66, -0, 99, -11}

	for i := 0; i < len(inserted); {
		//4 operations, so 4 loops expected
		//and for the above table of size 16 we should
		//be good with increasing ‘i’ within this loop
		ali.PushFront(inserted[i])
		ali.PushBack(inserted[i+1])
		ali.PushBack(inserted[i+2])
		ali.PushFront(inserted[i+3])
	}

	var expected [16]int = [16]int{
		-11, 66, 1, 19, -84, -21, 77, 43,
		22, 1, 33, 54, -52, 5, -0, 99}

	for i := 0; i < len(expected); i++ {
		if got, err := ali.Get(i); err == nil {
		if expected[i] != got.(int) {
				t.Errorf("expected[%v]: %v != got: %v\n",
					i, expected[i], got.(int))
			}
		} else {
			t.Errorf("Error while using Get() function: %v\n",
			err)
	}
	}

	if got, err := ali.PopFront(); err == nil {
		if got != -11 {
			t.Errorf("PopFront() failed; exp: -11, got: %v\n",
				got)
		}
	} else {
		t.Errorf("Error while using PopFront() function: %v\n",
			err)
}


	if got, err := ali.PopBack(); err == nil {
		if got != 99 {
			t.Errorf("PopBack() failed; exp: 99, got: %v\n",
				got)
		}
	} else {
		t.Errorf("Error while using PopBack() function: %v\n",
			err)
}


if err := ali.Insert(16, -1000); err != nil {
	t.Errorf("Error while using Insert(16) function: %v\n",
			err)
}

if err := ali.Insert(0, -2000); err != nil {
	t.Errorf("Error while using Insert(0) function: %v\n",
			err)
}

if err := ali.Insert(9, -3000); err != nil {
	t.Errorf("Error while using Insert(9) function: %v\n",
			err)
}

var exp []int = []int{
-2000, -11, 66, 1, 19, -84, -21, 77, 43,
-3000, 22, 1, 33, 54, -52, 5, -0, 99, -1000}

	for i := 0; i < len(exp); i++ {
		if got, err := ali.Get(i); err == nil {
			if exp[i] != got.(int) {
				t.Errorf("exp[%v]: %v != got: %v\n", i, exp[i], got.(int))
			}
		} else {
			t.Errorf("Insert failed with err: %v\n", err)
		}
	}
}

func TestIndexOf(t *testing.T) {

	var ali IArrayList = &MyArrayList{arrayList: []interface{}{-12, 1, 789}}

	if v, e := ali.IndexOf(1); e == nil {
		if v != 1 {
			t.Error("Failed to get index of 1")
		}
	} else {
		t.Error("Got error while getting index of 1")
	}


	if v, e := ali.IndexOf(-12); e == nil {
		if v != 0 {
			t.Error("Failed to get index of -12")
		}
	} else {
		t.Error("Got error while getting index of -12")
	}


	if v, e := ali.IndexOf(789); e == nil {
		if v != 2 {
			t.Error("Failed to get index of 789")
		}
	} else {
		t.Error("Got error while getting index of 789")
	}

	if _, e := ali.IndexOf(-10); e == nil {
		t.Error("Expected error for non existing value")
	}
}

func TestContains(t *testing.T) {

	var ali IArrayList = &MyArrayList{arrayList: []interface{}{-12, 1, 789}}

	if ali.Contains(-10) {
		t.Error("Unexpected value exists in the container")
	}

	if !ali.Contains(789) {
		t.Error("Expected value does not exist in the container")
	}
}
*/

func compare(expected []int, actual IArrayList, t *testing.T) {
	if len(expected) != actual.Size() {
		t.Errorf("Size differs: %v != %v\n", len(expected), actual.Size())
	}

	for i := 0; i < len(expected); i++ {
		if got, err := actual.Get(i); err == nil {
			if expected[i] != got.(int) {
				t.Errorf("exp[%v]: %v != got: %v\n",
					i, expected[i], got.(int))
			}
		} else {
			t.Error("Unexpected error returned from Get()")
		}
	}
}

func getTestArrayList() IArrayList {
	return &MyArrayList{arrayList: []interface{}{
		43, 22, 1, 77, -21, 33, 54, -84,
		19, -52, 5, 1, 66, -0, 99, -11}}
}

func TestRemove(t *testing.T) {

	var ali IArrayList = getTestArrayList()


	if e := ali.Remove(112); e == nil {
		t.Error("Expected error not returned")
	}

	if e := ali.Remove(12); e != nil {
		t.Error("Unexpected error returned")
	}

	compare([]int{
		43, 22, 1, 77, -21, 33, 54, -84,
		19, -52, 5, 1, -0, 99, -11,
	}, ali, t)
}

func TestClearIsEmpty(t *testing.T) {

var ali IArrayList = getTestArrayList()

	if ali.Size() != 16 {
		t.Error("Unexpected size")
	}

	if ali.IsEmpty() {
		t.Error("Container shall not be empty")
	}

	ali.Clear()

	if !ali.IsEmpty() {
		t.Error("Container shall be empty")
	}
}
