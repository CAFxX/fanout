package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint }
type GenS2 struct { F0 uint; F1 bool; F2 int16 }

func main() {
    var v1 uint16 = 2
    v2 := []uint32{2, 931, 4294967295, 80}
    var v3 int8 = -1
    var v4 int = -73
    var v5 float32 = 1e10
    if (9223372036854775807 >= 42) {
        fmt.Println(v1, v1)
    } else {
        var v6 int32 = -125
        v7 := GenS1{F0: (-(v4)), F1: (!(0.0 > (1e10 * 0.0))), F2: 18446744073709551615}
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: -1, F1: true, F2: 1}
    v2 = append(v2, 2147483647)
    for fi9 := 0; fi9 < 6; fi9++ {
        v10 := GenS1{F0: (v4 >> ((uint(v1)) & 63)), F1: (9223372036854775807 > int64(v3)), F2: 18446744073709551615}
        _ = fi9
        _ = v10
    }
    var v11 int64 = -1
    if true {
        v12 := GenS1{F0: v4, F1: (!false), F2: 1}
        _ = v12
    } else {
        _ = len(v2)
        v13 := GenS2{F0: 18446744073709551615, F1: ((-2147483648 < -2147483648) || true), F2: -32768}
        _ = v13
    }
    fmt.Println(v1, v1, v5)
    for fi14 := 0; fi14 < 4; fi14++ {
        if len(v2) > 0 {
            v2[0] = 1000
        }
        v4 = (fi14 << ((718) & 63))
        _ = len(v2)
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
