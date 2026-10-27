package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
func genFn0() uint64 {
    if (false || (((1e10 * float32(0.1)) <= 0.1) || false)) {
        var v5 uint = 4294967295
        if (false && (!false)) {
            for fi6 := 0; fi6 < 8; fi6++ {
                v5 = 0
                v7 := []int{(fi6 / ((2147483647) | 1)), 862, fi6, fi6}
                _ = fi6
                _ = v7
            }
            fmt.Println(v5)
        }
        _ = v5
    }
    var v8 uint32 = 245
    for fi9 := 0; fi9 < 14; fi9++ {
        v8 = (^(v8))
        v10 := GenS1{F0: 2, F1: 13}
        _ = fi9
        _ = v10
    }
    _ = v8
    return 1
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]int)
    var v3 int = 191
    v4 := make(map[uint32]bool)
    var v11 uint = 2
    var v12 float64 = 1e10
    v11 = (v11 & v11)
    if true {
        var v13 uint16 = 0
        var v14 int32 = -635
        _ = v13
        _ = v14
    } else {
        v11 = v11
        v3 = v3
    }
    var v15 uint = v11
    if false {
        if false {
            var v16 int8 = 42
            _ = v16
        }
        var v17 float64 = v12
        var v18 int64 = -746
        _ = v17
        _ = v18
    } else {
        v11 -= v15
        v3 = (v3 / ((v3) | 1))
        for fi19 := 0; fi19 < 18; fi19++ {
            for fi20 := 0; fi20 < 4; fi20++ {
                if (!true) {
                    v11 = v11
                    v21 := GenS1{F0: 169, F1: 42}
                    _ = v21
                } else {
                    var v22 float64 = v12
                    _ = v22
                }
                _ = len(v4)
                v23 := []int{v3, -9223372036854775808, 32767, v3}
                _ = fi20
                _ = v23
            }
            fmt.Println(fi19, v12, v3)
            if true {
                v12 -= (((float64((3.14159 / 1.0)) * (float64((3.14159 + 0.1)) / float64(0.1))) / (1e10 / (v12 + ((float64((0.0 / 1e10)) * float64(3.14159)) * -2.25)))) / v12)
            }
            _ = fi19
        }
    }
    var v24 uint32 = 65535
    if true {
        fmt.Println(v12)
        fmt.Println(v15, v24)
    }
    fmt.Println(v24, v24, v3)
    var v25 bool = (false || true)
    switch 4294967295 {
    case -2147483648:
        fmt.Println(v12)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
