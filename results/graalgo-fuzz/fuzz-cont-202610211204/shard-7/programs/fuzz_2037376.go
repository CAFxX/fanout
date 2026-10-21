package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 int16; F2 int }
func genFn0(p0 uint16, p1 uint64) int8 {
    v6 := []int{828, -499}
    if ((!(-2.25 <= (0.0 / float64(1e10)))) && (!(((float64(0.1) <= -2.25) || true) && false))) {
        var v7 bool = (-2.25 < 1.5)
        var v8 uint32 = 2147483647
        _ = v7
        _ = v8
    } else {
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    _ = v6
    return 0
}

func main() {
    var v1 bool = false
    var v2 uint32 = 161
    var v3 int = -9223372036854775808
    var v4 float64 = (float64((3.14159 - 0.1)) + ((1e10 / (float64(1.5) * 1.5)) + (1e10 * 0.0)))
    var v5 float32 = 3.14159
    v9 := GenS1{F0: 42}
    v10 := []int{(v3 >> ((18446744073709551615) & 63)), (v3 * v3), v3}
    for fi11 := 0; fi11 < 2; fi11++ {
        v3 += fi11
        _ = fi11
    }
    var v12 float64 = v4
    var v13 int64 = 672
    var v14 uint16 = 978
    if (0 >= 1) {
        v15 := GenS1{F0: 42}
        var v16 int8 = genFn0(255, 504)
        v10 = append(v10, v3)
        _ = v15
        _ = v16
    }
    switch "\n\t" {
    case "":
        if (v5 >= ((v5 / 3.14159) * float32((0.1 * (-2.25 * v4))))) {
            v13 = -9223372036854775808
            var v17 bool = (!true)
            switch "foo bar" {
            case "hello":
                var v18 int16 = -32768
                _ = v18
            case "foo bar":
                switch "\n\t" {
                case "\n\t":
                    if len(v10) > 0 {
                        v10[0] = (v3 - v3)
                    }
                }
            default:
                var v19 uint16 = (v14 / ((255) | 1))
                _ = v19
            }
            _ = v17
        } else {
            v20 := GenS1{F0: 1}
            _ = v20
        }
    case "\n\t":
        var v21 uint8 = 0
        _ = v21
    case "foo bar":
        v5 = 0.0
    default:
        var v22 int64 = v13
        _ = v22
    }
    if (float32(v12) <= (float32(v12) / -2.25)) {
        v12 -= -2.25
    } else {
        fmt.Println(v12, v4, v2)
    }
    var v23 int32 = -2147483648
    if len(v10) > 0 {
        v10[0] = -2147483648
    }
    fmt.Println(v23, v4, v13)
    v24 := GenS2{F0: (^(v3)), F1: 0, F2: 32768}
    var v25 uint = 255
    var v26 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
