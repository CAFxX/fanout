package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int }
type GenS2 struct { F0 int; F1 int64 }

func main() {
    var v1 float64 = (3.14159 / 3.14159)
    var v2 int64 = 42
    var v3 uint32 = 65535
    v4 := []int{-678, -2147483648, 893, 0}
    var v5 uint64 = 4294967295
    v6 := make(map[int]bool)
    v2 = v2
    v2 = v2
    v1 = float64(((-2.25 + 0.0) + (1e10 / ((1.5 - (((float32(1e10) * float32((0.1 + v1))) + (1e10 * (float32(1e10) / (((-2.25 / 1e10) - 0.0) + 0.1)))) / 1e10)) + 1e10))))
    switch -2147483648 {
    case 0:
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 int64 = -1
            var v9 string = "foo bar"
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    v10 := v6[42]
    if ((((0.0 + 1.5) + -2.25) / 1e10) > float32((1.5 * float64(((float32((-2.25 / 1.0)) / 1e10) / 1.0))))) {
        _ = len(v6)
    }
    v11 := GenS1{F0: (v2 | (v2 % ((42) | 1))), F1: (v2 == -9223372036854775808), F2: -1}
    for fi12 := 0; fi12 < 15; fi12++ {
        if (!(471 == -1)) {
            v5 -= v5
            var v13 int8 = -128
            _ = v13
        }
        if (!(!v10)) {
            if false {
                fmt.Println(v3)
            }
            v3 += (v3 - uint32(v5))
            delete(v6, (fi12 | fi12))
        } else {
            if v10 {
                var v14 int8 = -27
                _ = v14
            }
            var v15 uint = 83
            _ = v15
        }
        _ = fi12
    }
    v10 = true
    for fi16 := 0; fi16 < 6; fi16++ {
        switch ((("foo bar" + ("hello" + "")) + "a") + (("foo bar" + "a") + "hello")) {
        case "":
            v17 := GenS2{F0: fi16, F1: (-(v2))}
            _ = v17
        case "\n\t":
            fmt.Println(v2, v5, fi16)
        default:
            v5 = v5
        }
        var v18 float32 = (1.5 - 1.5)
        _ = fi16
        _ = v18
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        v10 = v10
        _ = fi19
    }
    v4 = append(v4, 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
