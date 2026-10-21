package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 bool; F2 uint8 }

func main() {
    v1 := []string{("\n\t" + "\n\t"), "foo bar"}
    v2 := []int{0, 4294967295, 4294967295, -9223372036854775808}
    v3 := make(map[uint32]string)
    var v4 int = 687
    var v5 int8 = 1
    delete(v3, 447)
    fmt.Println(v5, v5)
    v2 = append(v2, v4)
    v6 := GenS1{F0: v4}
    var v7 uint32 = 42
    switch v4 {
    case -678:
        v8 := GenS1{F0: 127}
        _ = v8
    case -9223372036854775808:
        v9 := GenS1{F0: (v4 / (((v4 * (v4 >> ((1) & 63)))) | 1))}
        _ = v9
    default:
        fmt.Println(v4, v5)
    }
    var v10 bool = (2 != 128)
    for fi11 := 0; fi11 < 2; fi11++ {
        for fi12 := 0; fi12 < 9; fi12++ {
            fmt.Println(v4)
            v4 = (fi11 * fi12)
            v10 = true
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v13 := GenS1{F0: (fi12 * fi12)}
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 uint16 = 862
            v5 -= (v5 - 100)
            _ = len(v2)
            v10 = ((v5 ^ (v5 >> ((42) & 63))) != 100)
            _ = fi14
            _ = v15
        }
        _ = fi11
    }
    delete(v3, 828)
    v16 := GenS1{F0: (v4 * 127)}
    var v17 uint32 = v7
    var v18 uint16 = 65535
    switch 0 {
    case 32767:
        if (!((v10 && v10) && (v10 && (!(((1.5 - 1e10) * -2.25) <= float64(0.1)))))) {
            if v10 {
                v4 -= (-(v4))
            }
        } else {
            fmt.Println(v10, v17)
        }
    }
    if (v10 && (v7 != (v7 << ((18446744073709551615) & 63)))) {
        var v19 int64 = -670
        _ = v19
    }
    v20 := GenS2{F0: 0, F1: false, F2: 2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
