package store

import (
    "testing"
    "time"
)

var testData = map[string]string{
    "hello": "world" ,
    "hello1": "world1",
    "hello2": "world2",
    "hello3": "world3",
    "hello4": "world4",
}

func Set(t *testing.T, key, value string) {
    Cache.set(key, value)
}


func Get(t *testing.T, key string) {
    return Cache.get(key)
}


func TestAll(t *testing.T) {
    for {
        go func() {
            for key, value := range testData {
                go Set(t, key, value)       
            }    
        }()
        time.Sleep(1 * time.Second)
        for key, value := range testData {
            Get(t, key, value)       
        }
    }
}


