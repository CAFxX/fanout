package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64 }
func genFn0(p0 int, p1 uint32) int8 {
    var v4 int64 = -1
    v5 := GenS1{F0: true, F1: 4294967295}
    for fi6 := 0; fi6 < 17; fi6++ {
        p1 -= 0
        _ = fi6
    }
    var v7 int16 = 0
    v8 := GenS1{F0: ((1e10 + 1.5) == (1e10 / ((0.1 - 1e10) - 1.5))), F1: 785}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    return 89
}

func main() {
    var v1 uint32 = 163
    var v2 float32 = 1.5
    var v3 float64 = -2.25
    v9 := []int{-9223372036854775808}
    v3 = (3.14159 - 0.1)
    v10 := GenS1{F0: false, F1: 18446744073709551615}
    fmt.Println(v2)
    fmt.Println(v1, v3, v3)
    fmt.Println(v1, v3, v1)
    v3 -= v3
    if true {
        v2 = 0.0
        v11 := GenS2{F0: true}
        var v12 float64 = (v3 - ((3.14159 - v3) / v3))
        _ = v11
        _ = v12
    }
    var v13 uint32 = (v1 ^ v1)
    _ = len(v9)
    v13 = (^(v13))
    for fi14 := 0; fi14 < 19; fi14++ {
        v13 = v1
        _ = fi14
    }
    v15 := GenS3{F0: -337}
    if len(v9) > 0 {
        v9[0] = -9223372036854775808
    }
    v2 = 3.14159
    var v16 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v9
}
