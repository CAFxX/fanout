package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint }
func genFn0(p0 int, p1 uint32, p2 bool) int32 {
    switch "foo bar" {
    case "foo bar":
        v6 := GenS1{F0: (^(p0))}
        _ = v6
    case "hello":
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 int32 = 673
            _ = fi7
            _ = v8
        }
    }
    if (((354 != 2) && p2) || (true || p2)) {
        switch ("hello" + "a") {
        case "hello":
            var v9 uint64 = 2
            _ = v9
        case "":
            v10 := GenS1{F0: p0}
            _ = v10
        case "a":
            var v11 int16 = 0
            _ = v11
        default:
            v12 := GenS2{F0: 100, F1: 887}
            _ = v12
        }
    } else {
        v13 := GenS2{F0: 29, F1: 577}
        p0 = 128
        v14 := GenS1{F0: (p0 / ((983) | 1))}
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}

func main() {
    var v1 string = "foo bar"
    var v2 uint32 = 4294967295
    v3 := []int{-2147483648, 2147483647, 0}
    var v4 float32 = (1e10 + 0.1)
    v5 := make(map[int]uint32)
    v15 := GenS2{F0: -128, F1: 1}
    var v16 int16 = 672
    _ = len(v3)
    switch "hello" {
    case "hello":
        var v17 int8 = 127
        _ = v17
    default:
        fmt.Println(v4, v1)
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        v19 := GenS1{F0: (-(fi18))}
        switch "hello" {
        case "a":
            var v20 float32 = 0.1
            _ = v20
        case "hello":
            for fi21 := 0; fi21 < 10; fi21++ {
                switch fi18 {
                case 42:
                    v22 := v5[fi18]
                    _ = v22
                default:
                    var v23 uint64 = 2
                    _ = v23
                }
                var v24 float64 = 3.14159
                v2 += v2
                _ = fi21
                _ = v24
            }
        case "\n\t":
            var v25 int16 = v16
            _ = v25
        }
        _ = fi18
        _ = v19
    }
    fmt.Println(v4, v16, v16)
    v2 -= v2
    var v26 uint64 = 829
    v1 = "a"
    fmt.Println(v26, v4)
    v26 += 359
    var v27 string = v1
    var v28 uint8 = 42
    v27 = v27
    fmt.Println(v4, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
