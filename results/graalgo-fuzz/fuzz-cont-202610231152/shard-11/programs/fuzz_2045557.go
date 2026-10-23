package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int }
func genFn0(p0 int16) uint32 {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: 100, F1: 255, F2: 0}
        _ = fi6
        _ = v7
    }
    _ = p0
    return 1
}

func main() {
    var v1 uint8 = 1
    v2 := []int{-702, -746, 127, 127}
    var v3 uint64 = 42
    var v4 uint8 = (v1 / ((128) | 1))
    v5 := make(map[string]bool)
    fmt.Println(v3, v4, v3)
    if false {
        v3 -= (v3 * (^(v3)))
    } else {
        if ((3.14159 / 1e10) <= (-2.25 / -2.25)) {
            v3 = (v3 + 1)
        }
        for fi8 := 0; fi8 < 10; fi8++ {
            fmt.Println(v3, v4, fi8)
            if (false || (1e10 == (-2.25 - 1.5))) {
                var v9 int16 = 1
                _ = len(v2)
                _ = v9
            }
            _ = fi8
        }
        var v10 int16 = 230
        _ = v10
    }
    var v11 bool = (0 > 9223372036854775807)
    var v12 uint = 2
    v11 = false
    v3 = v3
    switch -9223372036854775808 {
    case 32767:
        var v13 uint = v12
        _ = v13
    }
    var v14 int32 = 906
    v15 := GenS1{F0: 0, F1: (v12 << ((uint(v3)) & 63)), F2: 4294967295}
    for fi16 := 0; fi16 < 5; fi16++ {
        v4 = 2
        _ = fi16
    }
    switch 1 {
    case 32767:
        var v17 float32 = 1e10
        _ = v17
    case 127:
        fmt.Println(v3)
    case 4294967295:
        v18 := GenS1{F0: 50, F1: (v12 / ((v12) | 1)), F2: -131}
        _ = v18
    }
    var v19 int64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
