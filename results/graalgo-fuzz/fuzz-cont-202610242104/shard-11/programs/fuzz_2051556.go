package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int32; F1 string }

func main() {
    v1 := []uint32{2}
    var v2 bool = true
    v3 := make(map[uint32]int)
    var v4 uint64 = 2
    fmt.Println(v2)
    var v5 uint64 = v4
    v5 = v5
    var v6 float64 = 1e10
    v4 += (v4 | (v5 ^ 18446744073709551615))
    if (v5 > 4294967295) {
        fmt.Println(v6, v2)
    } else {
        v6 -= v6
        var v7 bool = (100 > -128)
        _ = v7
    }
    v8 := GenS1{F0: 9223372036854775807, F1: 153}
    var v9 int32 = 2147483647
    v10 := GenS1{F0: 373, F1: 0}
    v2 = (18446744073709551615 <= 31)
    var v11 int64 = 900
    if (!(v2 && v2)) {
        v12 := GenS1{F0: v11, F1: 255}
        _ = v12
    } else {
        v6 = v6
        if (true || (!(!(4294967295 == 4294967295)))) {
            v5 = v5
        }
    }
    if ((true && (429 != 91)) || false) {
        delete(v3, 42)
        v13 := GenS2{F0: (v9 / ((v9) | 1)), F1: ("a" + "hello")}
        _ = v13
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        if len(v1) > 0 {
            v1[0] = 42
        }
        _ = fi14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
