package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }

func main() {
    var v1 uint8 = 128
    v2 := make(map[int]int)
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]uint32)
    switch 2147483647 {
    case 42:
        for fi5 := 0; fi5 < 16; fi5++ {
            if false {
                v1 += (v1 - v1)
            } else {
                var v6 int16 = -976
                var v7 int8 = 42
                var v8 int32 = 330
                _ = v6
                _ = v7
                _ = v8
            }
            var v9 float64 = 3.14159
            _ = fi5
            _ = v9
        }
    case 4294967295:
        v10 := []int{9223372036854775807, 4294967295, 4294967295}
        _ = v10
    }
    v1 = v1
    if true {
        v1 -= 0
        var v11 uint32 = 1
        _ = v11
    }
    fmt.Println(v1)
    v1 = (v1 & v1)
    v1 = 2
    v12 := GenS1{F0: -32768, F1: 0}
    var v13 uint16 = 0
    var v14 int32 = -636
    var v15 string = "\n\t"
    if false {
        v13 += (v13 / ((v13) | 1))
        fmt.Println(v1, v1, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
