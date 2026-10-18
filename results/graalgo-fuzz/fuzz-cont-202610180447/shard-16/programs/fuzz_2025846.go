package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]int)
    var v3 int8 = -68
    if false {
        delete(v2, "\n\t")
        delete(v2, "a")
    }
    v3 = (-(v3))
    v4 := GenS1{F0: 32768, F1: -476, F2: 1}
    var v5 uint64 = 499
    v5 = v5
    v6 := GenS1{F0: -2147483648, F1: int32(v5), F2: 0}
    delete(v2, "hello")
    if (4294967295 <= 65535) {
        v7 := GenS1{F0: 128, F1: 0, F2: 4294967295}
        v8 := GenS1{F0: 177, F1: 2147483647, F2: 2147483647}
        fmt.Println(v5, v3)
        _ = v7
        _ = v8
    } else {
        var v9 float64 = (3.14159 + 0.0)
        _ = v9
    }
    v5 -= v5
    if (0 >= 935) {
        fmt.Println(v3)
        var v10 int = 32768
        _ = v10
    } else {
        var v11 uint8 = 255
        var v12 bool = (2 != uint16(v5))
        for fi13 := 0; fi13 < 19; fi13++ {
            delete(v1, 912)
            v3 = v3
            _ = fi13
        }
        _ = v11
        _ = v12
    }
    var v14 float32 = 1e10
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 string = "\n\t"
        _ = fi15
        _ = v16
    }
    var v17 uint64 = v5
    var v18 int64 = 42
    v19 := GenS2{F0: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
