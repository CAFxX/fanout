package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }

func main() {
    v1 := []int{2147483647, 42, 0, -9223372036854775808}
    var v2 int = 127
    var v3 uint8 = 255
    var v4 float64 = float64(0.1)
    var v5 bool = true
    fmt.Println(v2, v4)
    fmt.Println(v2, v2)
    var v6 float64 = ((v4 - (v4 / (float64(-2.25) + (v4 - 1.5)))) + -2.25)
    var v7 uint64 = 4294967295
    if (0 != 2) {
        var v8 int64 = -285
        _ = v8
    } else {
        for fi9 := 0; fi9 < 5; fi9++ {
            v4 = 0.0
            _ = fi9
        }
        var v10 int64 = 18
        _ = v10
    }
    if false {
        var v11 float64 = 3.14159
        _ = v11
    } else {
        var v12 int16 = 1000
        if (((1e10 - (((((0.0 / ((3.14159 - 3.14159) - float32(v6))) * 1e10) * (0.0 * float32(1e10))) + 1e10) + float32(float64(0.1)))) / (-2.25 - 1e10)) >= 0.0) {
            v13 := GenS1{F0: -9223372036854775808, F1: 42}
            v12 += v12
            for fi14 := 0; fi14 < 14; fi14++ {
                fmt.Println(v12)
                if true {
                    v5 = v5
                    var v15 int16 = v12
                    v1 = append(v1, fi14)
                    _ = v15
                }
                if len(v1) > 0 {
                    v1[0] = fi14
                }
                _ = fi14
            }
            _ = v13
        } else {
            v16 := GenS1{F0: -784, F1: 42}
            var v17 float32 = 3.14159
            _ = v16
            _ = v17
        }
        fmt.Println(v3)
        _ = v12
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 int16 = 183
        var v20 int16 = -520
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 float32 = -2.25
    v6 = (((v4 + 3.14159) / float64(float32(v4))) - (1e10 / 1e10))
    _ = len(v1)
    var v22 uint32 = 874
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
