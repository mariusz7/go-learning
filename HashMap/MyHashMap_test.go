package myhashtable


import (
	"testing"
	"math/rand"
	"strconv"
)


func TestMyHashMap(t *testing.T) {

	const valCount = 1000
	var hm MyHashMap
	var testPairs []KeyValue

	for i := 0; i < valCount; i++ {
		key := rand.Intn(10000)
		value := "keyInStr:" + strconv.Itoa(key)

		//fmt.Printf("key: %v, value: %v\n", key, value)

		testPairs = append(testPairs, KeyValue{key: key, value: value})

		hm.Insert(key, value)
	}

	for _, pair := range testPairs {
		if actualValue, err := hm.GetValue(pair.key); err == nil {
			if pair.value != actualValue {
				t.Errorf("Key: %v, expected value: %v, actual value: %v\n",
					pair.key, pair.value, actualValue)
			}
		} else {
			t.Errorf("Unexpected error: %v for key: %v\n", err, pair.key)
		}
	}
}
