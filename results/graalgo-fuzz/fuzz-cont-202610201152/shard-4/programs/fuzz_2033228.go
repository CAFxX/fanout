package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() uint16 {
    if false {
        var v4 uint8 = 21
        _ = v4
    } else {
        var v5 float32 = float32(-2.25)
        var v6 uint16 = 2
        _ = v5
        _ = v6
    }
    var v7 int64 = -1
    _ = v7
    return 865
}
func genFn1(p0 int64, p1 int16, p2 int64) uint64 {
    p1 = (-(p1))
    var v8 int = 9223372036854775807
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := []int{317, v8, (^(v8))}
        _ = fi9
        _ = v10
    }
    p0 += (p2 % (((^(p0))) | 1))
    v11 := []int{2147483647, v8, 2147483647, (v8 | v8)}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v8
    return 1
}

func main() {
    v1 := []int{}
    v2 := []int{-257, -67, 9223372036854775807, 2147483647}
    var v3 float64 = (1.5 - 1e10)
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 uint = 149
        _ = fi12
        _ = v13
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v14 := GenS1{F0: -234}
    if (!true) {
        var v15 bool = (((("" + "a") + "hello") != "foo bar") || (-65 != -128))
        if len(v1) > 0 {
            v1[0] = 127
        }
        v16 := GenS1{F0: -322}
        _ = v15
        _ = v16
    }
    var v17 float64 = v3
    for fi18 := 0; fi18 < 8; fi18++ {
        for fi19 := 0; fi19 < 17; fi19++ {
            var v20 uint64 = 42
            _ = fi19
            _ = v20
        }
        if (false || (1 >= 16)) {
            _ = len(v1)
        }
        fmt.Println(v17, v17, v3)
        _ = fi18
    }
    v17 -= v3
    var v21 uint = 4294967295
    v3 += v3
    var v22 uint16 = 169
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
