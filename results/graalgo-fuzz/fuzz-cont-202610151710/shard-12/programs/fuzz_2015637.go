package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
func genFn0(p0 uint8, p1 int8, p2 int8) uint8 {
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: false, F1: 621}
        fmt.Println(p1, p2, p0)
        _ = fi7
        _ = v8
    }
    p2 -= (p1 ^ -1)
    for fi9 := 0; fi9 < 9; fi9++ {
        p1 = p1
        var v10 int64 = -1
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return (p0 >> ((65535) & 63))
}

func main() {
    var v1 int8 = 1
    var v2 int = -1
    var v3 int64 = -9223372036854775808
    v4 := []int{(v2 % ((v2) | 1)), v2}
    var v5 int64 = -492
    v6 := make(map[string]bool)
    v11 := GenS1{F0: true, F1: 639}
    fmt.Println(v2)
    v1 += v1
    var v12 float64 = ((0.0 + 1.5) / (((1e10 + -2.25) / 0.1) - (0.1 + ((3.14159 * ((1e10 + 3.14159) * 1e10)) / 1.0))))
    if (v12 > float64(float32((((v12 * (0.0 - v12)) * (v12 / 3.14159)) * v12)))) {
        var v13 int64 = v5
        switch v2 {
        case 42:
            fmt.Println(v5, v5)
        case 127:
            switch "" {
            case "foo bar":
                v6["hello"] = true
            case "a":
                v2 = (v2 / ((v2) | 1))
            default:
                fmt.Println(v13, v1, v3)
            }
        }
        _ = v13
    } else {
        v6[""] = false
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 uint = 2
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 int16 = -776
            _ = fi16
            _ = v17
        }
    }
    var v18 int64 = (-(v3))
    v3 = (-(v3))
    v19 := GenS1{F0: true, F1: 856}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
