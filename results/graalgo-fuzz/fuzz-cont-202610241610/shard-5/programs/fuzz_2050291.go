package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int32; F2 int }
func genFn0() uint8 {
    var v4 int16 = 621
    for fi5 := 0; fi5 < 19; fi5++ {
        v4 = (v4 * v4)
        v4 += v4
        _ = fi5
    }
    v6 := GenS1{F0: 42}
    for fi7 := 0; fi7 < 1; fi7++ {
        switch (fi7 | fi7) {
        case 1:
            v4 += 475
        case 127:
            var v8 string = ""
            _ = v8
        default:
            switch fi7 {
            case 4294967295:
                v4 = v4
            }
        }
        _ = fi7
    }
    _ = v4
    _ = v6
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 127
    v3 := []int{-98, -9223372036854775808, -9223372036854775808}
    v2 += v2
    v2 -= v2
    if false {
        var v9 int = 4294967295
        _ = v9
    } else {
        v2 = v2
        v10 := GenS2{F0: 42, F1: 1, F2: 2147483647}
        var v11 float64 = 0.0
        _ = v10
        _ = v11
    }
    var v12 bool = true
    v13 := GenS2{F0: 42, F1: 42, F2: -608}
    if len(v3) > 0 {
        v3[0] = 408
    }
    var v14 int = 127
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
