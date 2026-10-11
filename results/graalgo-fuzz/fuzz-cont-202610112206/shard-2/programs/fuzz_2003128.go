package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int32; F2 string }
func genFn0(p0 int, p1 int32, p2 int16) string {
    for fi7 := 0; fi7 < 14; fi7++ {
        p2 = p2
        _ = fi7
    }
    p2 = (p2 * (^(p2)))
    _ = p0
    _ = p1
    _ = p2
    return "\n\t"
}

func main() {
    var v1 int8 = -70
    var v2 uint32 = 42
    var v3 string = ("foo bar" + "")
    var v4 bool = (12 <= 203)
    v5 := []int{1, 766}
    var v6 int = int(v2)
    if len(v5) > 0 {
        v5[0] = (v6 % ((128) | 1))
    }
    v2 = v2
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: ("a" + "a")}
        _ = fi8
        _ = v9
    }
    if len(v5) > 0 {
        v5[0] = (-(v6))
    }
    var v10 uint8 = 2
    switch 32768 {
    case 127:
        fmt.Println(v3, v4)
    case 2147483647:
        var v11 float32 = -2.25
        _ = v11
    }
    if v4 {
        for fi12 := 0; fi12 < 5; fi12++ {
            switch v6 {
            case -2147483648:
                var v13 float64 = (0.1 - (-2.25 + 0.0))
                _ = v13
            case -9223372036854775808:
                v10 = (v10 % (((^(v10))) | 1))
            }
            v14 := GenS2{F0: (v6 / ((fi12) | 1)), F1: -590, F2: v3}
            _ = fi12
            _ = v14
        }
    }
    v5 = append(v5, (v6 ^ (v6 >> ((334) & 63))))
    switch v3 {
    case "hello":
        var v15 int16 = -32768
        _ = v15
    case "\n\t":
        var v16 int64 = 237
        _ = v16
    case "a":
        var v17 uint32 = 2147483648
        _ = v17
    }
    v1 -= v1
    switch v3 {
    case "foo bar":
        v1 = (v1 / ((127) | 1))
    case "hello":
        fmt.Println(v2, v6)
    default:
        var v18 uint8 = v10
        _ = v18
    }
    var v19 int64 = -175
    v20 := GenS1{F0: (("foo bar" + v3) + (v3 + "hello"))}
    v4 = true
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
