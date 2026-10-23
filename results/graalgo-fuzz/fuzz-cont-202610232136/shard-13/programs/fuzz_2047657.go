package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int }
func genFn0(p0 int8, p1 uint16, p2 int32) int64 {
    switch "foo bar" {
    case "\n\t":
        var v4 float64 = 0.1
        _ = v4
    }
    v5 := GenS1{F0: (-(p0)), F1: -9223372036854775808, F2: -946}
    if (599 <= 834) {
        v6 := []int{}
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 int16 = -699
            v6 = append(v6, int(p2))
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 2147483648
    var v3 float32 = 0.0
    if false {
        for fi9 := 0; fi9 < 14; fi9++ {
            v10 := GenS1{F0: 0, F1: 42, F2: (fi9 % ((fi9) | 1))}
            var v11 uint64 = 2
            _ = fi9
            _ = v10
            _ = v11
        }
        v2 = 352
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        if true {
            var v13 uint16 = 2
            _ = v13
        } else {
            for fi14 := 0; fi14 < 7; fi14++ {
                var v15 uint32 = v2
                v16 := []int{}
                v3 = (3.14159 + (1e10 + -2.25))
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        var v17 int32 = -333
        _ = fi12
        _ = v17
    }
    var v18 int32 = -850
    var v19 string = (("" + "a") + "foo bar")
    switch 127 {
    case 750:
        v3 += 0.0
    default:
        v20 := GenS1{F0: -128, F1: 730, F2: 42}
        _ = v20
    }
    v3 += 1.5
    v21 := []int{32768, -427, 2147483647}
    fmt.Println(v18, v19)
    v22 := v1[v19]
    var v23 int64 = 519
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
