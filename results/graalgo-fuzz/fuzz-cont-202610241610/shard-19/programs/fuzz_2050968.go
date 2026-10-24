package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int16 }
type GenS2 struct { F0 bool; F1 uint32; F2 int }
func genFn0(p0 bool) uint {
    v7 := []int{9223372036854775807, -975, -1, 32768}
    var v8 uint = 689
    _ = p0
    _ = v7
    _ = v8
    return (v8 & (v8 % ((59) | 1)))
}

func main() {
    var v1 int16 = 42
    v2 := []int{2147483647, 2147483647, 42}
    v3 := []string{"foo bar"}
    v4 := make(map[int]uint32)
    var v5 float64 = (3.14159 + float64(1e10))
    v6 := make(map[int]bool)
    v9 := GenS1{F0: -128, F1: 100, F2: 564}
    var v10 int8 = 0
    switch "" {
    case "foo bar":
        for fi11 := 0; fi11 < 13; fi11++ {
            for fi12 := 0; fi12 < 3; fi12++ {
                var v13 float64 = v5
                _ = fi12
                _ = v13
            }
            v14 := GenS1{F0: v10, F1: (-(v10)), F2: (^(v1))}
            v5 += (v5 + v5)
            _ = fi11
            _ = v14
        }
    case "a":
        v6[0] = (42 < 115)
    }
    fmt.Println(v5)
    v15 := GenS1{F0: (v10 | v10), F1: 100, F2: v1}
    switch 128 {
    case -1:
        switch (("a" + ("" + "hello")) + "foo bar") {
        case "a":
            v10 -= (-(v10))
        case "":
            var v16 bool = (false || true)
            _ = v16
        default:
            switch 32768 {
            case -691:
                v17 := GenS2{F0: false, F1: 65535, F2: 2147483647}
                _ = v17
            default:
                v18 := GenS2{F0: (!(0.0 != 1.5)), F1: 1, F2: -1}
                _ = v18
            }
        }
    default:
        var v19 int8 = (v10 >> ((2) & 63))
        _ = v19
    }
    var v20 int32 = 163
    v1 = (v1 + 42)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
