package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint8 }
func genFn0(p0 int16, p1 int8) string {
    for fi7 := 0; fi7 < 15; fi7++ {
        p1 -= p1
        v8 := []int{(-(fi7)), 32767, (fi7 >> ((436) & 63))}
        _ = fi7
        _ = v8
    }
    var v9 uint32 = 2147483648
    v10 := []int{32768}
    var v11 uint = 1
    var v12 uint8 = 255
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return "a"
}

func main() {
    var v1 bool = ((1.5 <= ((3.14159 + (0.0 * ((0.1 / 1e10) / (3.14159 / ((0.0 + -2.25) * (1e10 * ((-2.25 / 1.0) + 0.1))))))) * (3.14159 / (0.1 - -2.25)))) && true)
    var v2 uint64 = 4294967295
    var v3 bool = false
    var v4 uint16 = 65535
    var v5 uint64 = uint64(v4)
    var v6 uint32 = 2
    v3 = (2 <= 42)
    if (!v3) {
        v13 := GenS1{F0: true, F1: 42, F2: 1}
        fmt.Println(v1, v1, v5)
        var v14 uint64 = v2
        _ = v13
        _ = v14
    }
    v1 = true
    switch -637 {
    case 9223372036854775807:
        v15 := GenS1{F0: (-128 == int8(v6)), F1: -9223372036854775808, F2: 209}
        _ = v15
    case 2147483647:
        if (!v3) {
            fmt.Println(v6, v6)
            if (float64(1.5) >= float64(0.0)) {
                var v16 int64 = 522
                _ = v16
            }
        }
    case -2147483648:
        v17 := []int{32768, 0, 0}
        _ = v17
    default:
        v18 := GenS1{F0: v1, F1: 1, F2: 213}
        _ = v18
    }
    var v19 int64 = -9223372036854775808
    v20 := []int{-2147483648, int(v6), 9223372036854775807, 202}
    v21 := GenS1{F0: (v4 >= 673), F1: v19, F2: 2}
    var v22 float64 = 0.0
    var v23 int16 = -32768
    fmt.Println(v1, v2, v22)
    var v24 bool = v1
    var v25 int8 = 107
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
