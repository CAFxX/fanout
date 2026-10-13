package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 0
    var v3 bool = true
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]string)
    var v6 int64 = 0
    fmt.Println(v6)
    v7 := GenS1{F0: v2, F1: 42}
    fmt.Println(v3)
    v2 = 32767
    v8 := GenS1{F0: 255, F1: -103}
    if (!false) {
        v9 := GenS1{F0: 0, F1: 126}
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 float32 = 3.14159
            v6 += (v6 + (v6 | (v6 / ((v6) | 1))))
            _ = fi10
            _ = v11
        }
        v2 = 10
        _ = v9
    }
    v6 += v6
    var v12 uint64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
