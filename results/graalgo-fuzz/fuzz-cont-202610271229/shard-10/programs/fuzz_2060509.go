package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }

func main() {
    var v1 uint8 = 26
    var v2 int64 = -9223372036854775808
    var v3 float64 = float64(float32(0.1))
    var v4 float32 = -2.25
    var v5 int64 = (^(v2))
    v6 := []uint32{4294967295}
    v1 = (v1 + (v1 << ((18446744073709551615) & 63)))
    v4 = 0.1
    v3 = (((float64(((v4 - v4) - v4)) - v3) / ((((1.5 - 1.5) + v3) * float64(v4)) - (float64(0.1) * (float64(3.14159) + 3.14159)))) * (float64((float32((0.1 / v3)) - 3.14159)) * 0.1))
    for fi7 := 0; fi7 < 6; fi7++ {
        switch (fi7 - -2147483648) {
        case 2147483647:
            var v8 float32 = (1.5 / v4)
            _ = v8
        default:
            fmt.Println(v1, v1)
        }
        fmt.Println(v3, v3)
        _ = fi7
    }
    fmt.Println(v5)
    var v9 float32 = (((float32(0.0) * (v4 + v4)) * (v4 * v4)) - 3.14159)
    var v10 uint32 = 42
    if (false || (1000 < -32768)) {
        v2 = (v5 | (v2 - v2))
    } else {
        if true {
            switch ("" + "") {
            case "foo bar":
                v3 = v3
            }
            v11 := GenS1{F0: v10, F1: -743}
            v2 = (-(v5))
            _ = v11
        }
        var v12 uint64 = 1
        var v13 uint8 = 255
        _ = v12
        _ = v13
    }
    fmt.Println(v9, v2)
    v10 = 1
    fmt.Println(v5, v3)
    if (45 > (v10 & 86)) {
        for fi14 := 0; fi14 < 13; fi14++ {
            v9 = v4
            switch (^(fi14)) {
            case 0:
                v15 := GenS1{F0: v10, F1: (-(fi14))}
                _ = v15
            }
            _ = fi14
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
