package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int32 }

func main() {
    var v1 uint8 = 128
    var v2 int16 = 909
    v3 := make(map[uint32]bool)
    v4 := make(map[int]bool)
    v5 := make(map[int]int)
    var v6 uint16 = 32767
    v7 := GenS1{F0: false, F1: 42, F2: 2147483647}
    v8 := []int{}
    if false {
        v2 = (v2 | 352)
    } else {
        _ = len(v4)
        var v9 uint = 4294967295
        _ = v9
    }
    var v10 uint8 = 255
    for fi11 := 0; fi11 < 16; fi11++ {
        _ = len(v8)
        if len(v8) > 0 {
            v8[0] = 2147483647
        }
        v1 += v1
        _ = fi11
    }
    var v12 float32 = 3.14159
    fmt.Println(v12)
    v13 := GenS1{F0: false, F1: 0, F2: 0}
    fmt.Println(v1, v2, v10)
    _ = len(v5)
    v14 := GenS1{F0: true, F1: 4294967295, F2: 42}
    v1 = (v1 & (v1 << ((65535) & 63)))
    fmt.Println(v12, v1, v1)
    fmt.Println(v2, v6, v1)
    if false {
        var v15 int64 = -914
        var v16 uint16 = v6
        _ = v15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
