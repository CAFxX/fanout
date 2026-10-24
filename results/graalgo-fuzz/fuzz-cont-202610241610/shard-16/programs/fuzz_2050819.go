package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 int8; F2 int8 }

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]uint32)
    var v3 int16 = 677
    var v4 string = "hello"
    v5 := GenS1{F0: v4}
    v6 := GenS2{F0: 308, F1: -1, F2: 42}
    v7 := GenS2{F0: 65535, F1: 0, F2: 100}
    fmt.Println(v1, v1)
    var v8 float64 = -2.25
    v9 := GenS2{F0: 18446744073709551615, F1: 16, F2: 86}
    v8 = (v8 - v8)
    var v10 float32 = (1.5 * float32(float64((float32((v8 / float64((1.5 + 1e10)))) + 0.0))))
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := GenS2{F0: 18446744073709551615, F1: 42, F2: -1}
        _ = fi11
        _ = v12
    }
    v2[851] = 605
    if ((-2147483648 < -1) || false) {
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 uint = 348
            _ = fi13
            _ = v14
        }
        var v15 int32 = 1
        _ = v15
    } else {
        var v16 float32 = 3.14159
        _ = v16
    }
    if true {
        var v17 int = -9223372036854775808
        var v18 int64 = 392
        var v19 uint8 = 101
        _ = v17
        _ = v18
        _ = v19
    }
    v20 := []int{}
    v1 = (v3 << ((2) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
