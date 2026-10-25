package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int64; F1 bool }
func genFn0(p0 string) string {
    fmt.Println(p0)
    p0 = "hello"
    v7 := GenS1{F0: (42 >= 1), F1: 607, F2: 42}
    var v8 uint8 = 0
    _ = p0
    _ = v7
    _ = v8
    return p0
}

func main() {
    var v1 uint = 65535
    v2 := make(map[uint32]int)
    v3 := make(map[string]uint32)
    var v4 float64 = ((0.1 + 3.14159) - (0.1 * float64(0.0)))
    var v5 int32 = 1
    var v6 int = 2147483647
    fmt.Println(v1, v6, v1)
    for fi9 := 0; fi9 < 3; fi9++ {
        fmt.Println(v4)
        v5 += v5
        if true {
            fmt.Println(v6, fi9)
        } else {
            if (2 > 0) {
                if true {
                    switch "hello" {
                    case "foo bar":
                        var v10 bool = (!true)
                        _ = v10
                    case "a":
                        fmt.Println(v5)
                    default:
                        fmt.Println(v1, v4)
                    }
                    if (1 < 695) {
                        var v11 uint64 = 42
                        v11 = v11
                        _ = v11
                    }
                    var v12 int32 = v5
                    _ = v12
                } else {
                    v13 := GenS2{F0: 97, F1: (!true)}
                    _ = v13
                }
            }
            for fi14 := 0; fi14 < 19; fi14++ {
                var v15 float32 = -2.25
                _ = fi14
                _ = v15
            }
        }
        _ = fi9
    }
    v16 := GenS1{F0: true, F1: 749, F2: 381}
    v17 := v3["hello"]
    v18 := GenS2{F0: -9223372036854775808, F1: true}
    v19 := []int{v6, (-(v6))}
    for fi20 := 0; fi20 < 13; fi20++ {
        var v21 bool = (v17 != v17)
        v3["hello"] = v17
        var v22 bool = true
        _ = fi20
        _ = v21
        _ = v22
    }
    v19 = append(v19, (v6 * 32767))
    var v23 float64 = v4
    for fi24 := 0; fi24 < 17; fi24++ {
        var v25 int16 = -527
        switch (fi24 + (fi24 & -1)) {
        case -771:
            v26 := GenS2{F0: 202, F1: false}
            _ = v26
        default:
            v1 += (v1 ^ 65535)
        }
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
