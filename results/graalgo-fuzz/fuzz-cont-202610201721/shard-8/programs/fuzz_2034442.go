package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint }
func genFn0(p0 int16, p1 uint8) uint16 {
    p1 = p1
    v5 := []int{-9223372036854775808}
    var v6 int64 = -9223372036854775808
    p0 -= p0
    if (!true) {
        switch 128 {
        case -2147483648:
            v7 := GenS1{F0: 136, F1: (18446744073709551615 == 85), F2: 1}
            _ = v7
        case 9223372036854775807:
            if (3.14159 != -2.25) {
                var v8 float64 = 1.5
                fmt.Println(p0, v8)
                p0 -= p0
                _ = v8
            }
        case -1:
            var v9 int64 = 9223372036854775807
            _ = v9
        }
        var v10 float32 = ((0.1 - 0.0) / (3.14159 + float32(-2.25)))
        fmt.Println(v6)
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 uint8 = 1
    var v2 float32 = -2.25
    v3 := make(map[string]bool)
    v4 := []int{0, 32768, -275}
    var v11 uint32 = 0
    v12 := GenS1{F0: (v1 * 115), F1: (true && (true || true)), F2: 1}
    v13 := GenS1{F0: (v1 * v1), F1: false, F2: 1}
    fmt.Println(v2)
    var v14 int64 = 0
    if false {
        v15 := GenS1{F0: (v1 % (((^(v1))) | 1)), F1: (18446744073709551615 > 4294967295), F2: 4294967295}
        _ = v15
    } else {
        if true {
            var v16 int32 = -157
            _ = v16
        }
        fmt.Println(v11, v1, v14)
    }
    v17 := GenS1{F0: (^(v1)), F1: true, F2: 255}
    var v18 uint32 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
