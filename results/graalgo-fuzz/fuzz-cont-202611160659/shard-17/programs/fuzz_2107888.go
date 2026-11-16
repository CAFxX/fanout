package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int16 }
type GenS2 struct { F0 uint }
func genFn0(p0 string) string {
    if false {
        p0 = (p0 + (p0 + ("foo bar" + (p0 + p0))))
    }
    v6 := GenS1{F0: true, F1: -452, F2: -383}
    p0 = p0
    fmt.Println(p0)
    _ = p0
    _ = v6
    return p0
}

func main() {
    var v1 int32 = 1
    var v2 uint16 = 2
    var v3 bool = true
    var v4 int64 = 0
    v5 := []int{-1, -647, -9223372036854775808, -306}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v7 uint = 255
    if (1 != 65535) {
        var v8 float64 = (0.1 - 1e10)
        var v9 int16 = 1
        v10 := GenS2{F0: v7}
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        fmt.Println(fi11, v4, v1)
        switch -9223372036854775808 {
        case -2147483648:
            var v12 uint16 = v2
            _ = v12
        case 127:
            if v3 {
                var v13 uint64 = 42
                fmt.Println(v2, fi11, v2)
                _ = v13
            }
        }
        v1 = (v1 * v1)
        _ = fi11
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        v1 += (v1 + (-(v1)))
        _ = fi14
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        v2 -= v2
        var v16 int64 = v4
        var v17 uint16 = v2
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 float64 = ((-2.25 - 0.0) - -2.25)
    var v19 bool = v3
    v2 += (v2 >> ((2) & 63))
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    v20 := GenS1{F0: v19, F1: (v1 / ((-846) | 1)), F2: 42}
    v19 = (!((4294967295 > 42) || true))
    v4 -= (v4 ^ v4)
    var v21 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
