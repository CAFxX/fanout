package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int }
type GenS2 struct { F0 int64 }
func genFn0(p0 int8, p1 int) uint8 {
    var v4 bool = (("a" + "hello") >= "\n\t")
    p1 -= p1
    _ = p0
    _ = p1
    _ = v4
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 874
    var v3 float64 = (float64(1.5) - 0.1)
    var v5 int32 = 292
    var v6 float32 = 0.0
    var v7 int = 0
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(v7)
        _ = fi8
    }
    v5 -= -2147483648
    _ = len(v1)
    v9 := GenS1{F0: 17, F1: false, F2: -265}
    v5 = 1
    var v10 bool = ((42 == 4294967295) || true)
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 int16 = 0
        v13 := GenS2{F0: 9223372036854775807}
        v2 = (v2 / ((v2) | 1))
        _ = fi11
        _ = v12
        _ = v13
    }
    v5 = v5
    if (((0.1 / v6) - v6) < v6) {
        v10 = v10
        if (42 > 752) {
            fmt.Println(v6, v6, v2)
        }
    }
    fmt.Println(v3, v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
