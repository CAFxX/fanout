package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 uint; F2 int8 }
func genFn0(p0 int16, p1 uint16, p2 int) uint64 {
    for fi7 := 0; fi7 < 10; fi7++ {
        p0 = p0
        _ = fi7
    }
    v8 := []int{}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 2
}

func main() {
    v1 := []uint32{}
    var v2 uint32 = 42
    var v3 uint = 400
    var v4 float64 = (float64(1.5) / (1e10 * 1.5))
    var v5 string = "foo bar"
    v6 := make(map[string]int)
    fmt.Println(v4, v5, v4)
    v3 += 886
    if ((v2 % ((1) | 1)) == (v2 << ((v3) & 63))) {
        v9 := v6["foo bar"]
        _ = v9
    } else {
        v10 := GenS1{F0: false}
        var v11 string = (v5 + ("hello" + (v5 + v5)))
        v12 := GenS1{F0: false}
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 int = 2147483647
    switch (v13 / (((-(v13))) | 1)) {
    case 0:
        v3 += v3
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        v3 -= (^(v3))
        var v15 uint32 = v2
        var v16 int8 = 127
        _ = fi14
        _ = v15
        _ = v16
    }
    switch 32768 {
    case 9223372036854775807:
        v17 := GenS1{F0: false}
        _ = v17
    case 153:
        if ((2147483647 != 1) && true) {
            v18 := GenS2{F0: int64(v13), F1: v3, F2: 0}
            _ = v18
        }
    default:
        var v19 uint8 = 255
        _ = v19
    }
    var v20 int64 = 42
    var v21 uint = 0
    v5 = ""
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
