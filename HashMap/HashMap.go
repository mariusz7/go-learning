package myhashtable

import (
	"errors"
    "hash/fnv"
	"strconv"
)

const arraySize = 1024

type MyHashMap struct {
	array [arraySize][]KeyValue
}

//For now this supports (int -> string) pairs. TODO generalize it for interface{} type
type KeyValue struct {
	key int
	value string
}

func (hm *MyHashMap) Insert(key int, value string) {
	var hash int = keyToHash(key)
	var index int = hashToIndex(hash)

	hm.array[index] = append(hm.array[index], KeyValue{key: key, value; value})
}

//Function from https://stackoverflow.com/questions/13582519/how-to-generate-hash-number-of-a-string-in-go
func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func (hm *MyHashMap) keyToHash(key int) uint32 {

	var keyString string = strconv.Itoa(key)
	var hash uint32 = hash(keyString)

	return hash
}

func (hm *MyHashMap) hashToIndex(hash int) int {
	return hash % len(hm.array)
}

func (hm *MyHashMap) GetValue(key int) (string, error) {

	var hash int = keyToHash(key)
	var index int = hashToIndex(hash)

	for _, pair := range hm.array {
		if key == pair.key {
			return pair.value, nil
		}
	}

	return "", errors.New("Key not found")
}
