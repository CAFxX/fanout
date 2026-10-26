package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int32 }

func main() {
    var v1 uint32 = 0
    var v2 int = 4294967295
    var v3 bool = (!(-324 < (v2 % ((v2) | 1))))
    var v4 int = v2
    v2 -= (v4 >> ((153) & 63))
    if (!v3) {
        v5 := []int{v4, v4, v2}
        _ = v5
    } else {
        v2 = v4
        v2 = 42
    }
    v6 := GenS1{F0: 1, F1: 2147483647, F2: -1}
    if false {
        v7 := GenS1{F0: 780, F1: int32(v2), F2: 42}
        var v8 float32 = ((3.14159 + 1e10) / (float32(3.14159) / (1.5 - -2.25)))
        v9 := []int{2147483647, v4, -797}
        _ = v7
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 10; fi10++ {
            v11 := GenS1{F0: 32767, F1: 2147483647, F2: -2147483648}
            v12 := GenS1{F0: 297, F1: 851, F2: 597}
            var v13 int16 = -319
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        var v14 bool = v3
        _ = v14
    }
    var v15 uint32 = 42
    v15 = v1
    v16 := []int{}
    v2 += v4
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
