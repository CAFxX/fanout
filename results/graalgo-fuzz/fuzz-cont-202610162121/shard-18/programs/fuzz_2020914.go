package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }

func main() {
    var v1 int = 32768
    v2 := []uint32{2147483647, 2, 403, 399}
    var v3 bool = (-2.25 >= float64(0.1))
    v1 = 127
    v1 -= (v1 % ((v1) | 1))
    var v4 int64 = -1
    var v5 uint16 = 255
    v2 = append(v2, 180)
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := GenS1{F0: 2, F1: 151}
        var v8 uint8 = 127
        var v9 int32 = int32(fi6)
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    v5 += (v5 + (^(v5)))
    switch (v1 - (^(v1))) {
    case 4294967295:
        v1 -= v1
    case 32767:
        v1 = (v1 * v1)
    }
    _ = len(v2)
    v4 += 42
    v1 = v1
    var v10 uint8 = 0
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := GenS1{F0: 4294967295, F1: uint(v10)}
        if len(v2) > 0 {
            v2[0] = 0
        }
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        v1 = (fi13 - -9223372036854775808)
        if ((true && (3.14159 > (1e10 / (1.5 * float32(((0.0 + 0.0) + 1.5)))))) || (v1 != 1)) {
            fmt.Println(v1, v10, v10)
        }
        _ = fi13
    }
    v14 := GenS1{F0: 0, F1: uint(v5)}
    var v15 string = ""
    if (1.5 < 1e10) {
        var v16 uint16 = 255
        v4 = -790
        v10 = (v10 >> ((0) & 63))
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
