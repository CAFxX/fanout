package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
func genFn0(p0 uint8) int8 {
    p0 -= 1
    var v7 int16 = 604
    v7 = v7
    _ = p0
    _ = v7
    return 42
}
func genFn1(p0 bool, p1 int32, p2 bool) uint {
    fmt.Println(p2, p0, p1)
    v8 := []int{127, 4294967295, 599}
    p1 += p1
    p2 = true
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 193
}

func main() {
    var v1 float32 = float32((0.0 / (1e10 - 1.5)))
    v2 := make(map[int]uint32)
    var v3 float64 = 0.0
    var v4 uint16 = 0
    var v5 int8 = 127
    var v6 int16 = -1
    v5 -= (v5 & (-(v5)))
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 uint32 = 1
        if true {
            v11 := GenS1{F0: v10, F1: 2}
            var v12 int8 = (v5 | (v5 | 82))
            _ = v11
            _ = v12
        } else {
            v13 := GenS1{F0: 65535, F1: 765}
            _ = v13
        }
        var v14 uint32 = v10
        _ = fi9
        _ = v10
        _ = v14
    }
    var v15 int8 = (-(v5))
    delete(v2, 42)
    v5 -= -128
    v5 -= 127
    v16 := []int{32768, -849, 127}
    v6 -= 1000
    var v17 int16 = (v6 & v6)
    var v18 uint64 = 320
    v3 = (v3 * 1e10)
    for fi19 := 0; fi19 < 19; fi19++ {
        v20 := GenS1{F0: 0, F1: (v18 / ((4294967295) | 1))}
        _ = fi19
        _ = v20
    }
    v5 -= v15
    v21 := GenS1{F0: 0, F1: (v18 % ((299) | 1))}
    v22 := GenS1{F0: 4294967295, F1: 42}
    if false {
        v18 += 0
        switch "foo bar" {
        case "hello":
            v18 = (v18 & v18)
        case "":
            fmt.Println(v3, v5, v18)
        }
        fmt.Println(v6, v15, v5)
    }
    var v23 uint8 = 127
    for fi24 := 0; fi24 < 10; fi24++ {
        v6 += v6
        for fi25 := 0; fi25 < 19; fi25++ {
            v3 += v3
            _ = fi25
        }
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
