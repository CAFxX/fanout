package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
func genFn0() bool {
    var v4 int32 = 1
    v5 := []int{42}
    for fi6 := 0; fi6 < 4; fi6++ {
        if ((2 <= 0) && true) {
            fmt.Println(fi6, fi6)
            var v7 int64 = -9223372036854775808
            _ = v7
        }
        _ = fi6
    }
    v4 = (v4 * -1)
    _ = v4
    _ = v5
    return false
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"", ("foo bar" + ("" + "a"))}
    var v3 int8 = -128
    var v8 float32 = (1e10 * ((1.5 + 0.1) + 1e10))
    if (255 > 2) {
        fmt.Println(v8)
    } else {
        v8 += v8
        var v9 uint8 = 106
        _ = v9
    }
    var v10 float64 = (1e10 * (0.0 + float64((v8 - 0.1))))
    switch 495 {
    case -2147483648:
        var v11 string = "hello"
        _ = v11
    case -9223372036854775808:
        v1["hello"] = 65535
    case 460:
        v3 -= (^(v3))
    }
    if false {
        var v12 float32 = (3.14159 / (v8 * 3.14159))
        _ = v12
    } else {
        var v13 int32 = -691
        v10 = ((float64(float32(float64(3.14159))) / ((v10 + v10) + (0.0 * v10))) / 1.5)
        _ = v13
    }
    fmt.Println(v3)
    var v14 int16 = -32768
    v3 -= v3
    var v15 float32 = float32(((v10 + (v10 / v10)) - (v10 + 1.5)))
    v15 = 0.1
    v16 := GenS1{F0: 2147483647, F1: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v8
}
