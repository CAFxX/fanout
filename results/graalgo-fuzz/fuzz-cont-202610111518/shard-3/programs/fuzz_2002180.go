package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint16 }
func genFn0() int32 {
    if false {
        var v5 float32 = float32(1.5)
        var v6 int16 = 0
        var v7 int64 = 42
        _ = v5
        _ = v6
        _ = v7
    }
    var v8 int64 = 42
    var v9 uint8 = 246
    v8 = -294
    v10 := GenS1{F0: (^(v8)), F1: (!(1 <= v8)), F2: 65535}
    _ = v10
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 int16 = -1
    var v2 uint = 65535
    var v3 int8 = -1
    var v4 bool = (((((0.0 / float64(float32(((1e10 + 3.14159) + (-2.25 - ((((0.0 / 1.0) - 1e10) - (-2.25 - 1.5)) * 0.0)))))) - 0.0) - 1e10) > 0.1) || ("hello" < ("a" + "")))
    switch 32767 {
    case -105:
        var v11 uint64 = 2
        _ = v11
    }
    v12 := GenS1{F0: 0, F1: (0.1 > (0.0 - float32((-2.25 / (0.1 / 3.14159))))), F2: 65535}
    v3 = v3
    v13 := []int{127, 0, 32767, 32767}
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 uint8 = 42
        var v16 uint16 = 186
        v17 := GenS1{F0: 9223372036854775807, F1: v4, F2: (v16 & v16)}
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    _ = len(v13)
    v13 = append(v13, 9223372036854775807)
    if ((1.5 >= (0.0 + 3.14159)) || v4) {
        var v18 int8 = (v3 + v3)
        var v19 uint = (v2 | (v2 ^ v2))
        fmt.Println(v1, v3)
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        for fi21 := 0; fi21 < 8; fi21++ {
            _ = len(v13)
            _ = fi21
        }
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
