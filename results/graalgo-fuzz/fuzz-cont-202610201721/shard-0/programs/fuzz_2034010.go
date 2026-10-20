package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int; F1 uint; F2 uint32 }
func genFn0() bool {
    var v5 int16 = 0
    v6 := GenS1{F0: 1, F1: 407}
    var v7 uint32 = 570
    _ = v5
    _ = v6
    _ = v7
    return false
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{75}
    v3 := make(map[string]uint32)
    v4 := make(map[int]int)
    var v8 bool = (2 > 0)
    v8 = v8
    v9 := GenS1{F0: 502, F1: 2}
    v10 := GenS2{F0: -91, F1: 18446744073709551615, F2: 488}
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    v8 = v8
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 int8 = 1
        _ = fi11
        _ = v12
    }
    var v13 int8 = 127
    v8 = true
    if ((129 <= 2) || (!v8)) {
        var v14 int64 = 279
        var v15 int64 = v14
        v16 := GenS2{F0: -2147483648, F1: 809, F2: 2}
        _ = v14
        _ = v15
        _ = v16
    } else {
        if (2147483647 >= 4294967295) {
            v13 -= (v13 + -128)
            v8 = (-2.25 <= (3.14159 * 1.5))
            for fi17 := 0; fi17 < 5; fi17++ {
                _ = len(v3)
                _ = fi17
            }
        }
        v13 += -1
        var v18 int = 0
        _ = v18
    }
    v19 := GenS2{F0: 4294967295, F1: 2, F2: 2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
