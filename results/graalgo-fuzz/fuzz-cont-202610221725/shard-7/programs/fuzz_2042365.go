package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[string]int)
    var v2 uint = 1
    v3 := []int{0}
    v4 := []string{((("a" + "a") + "hello") + ("\n\t" + "a")), "a", "\n\t", ("hello" + ("" + (("foo bar" + "\n\t") + ("foo bar" + "hello"))))}
    v5 := make(map[int]uint32)
    var v6 int64 = 1
    var v7 uint64 = 2
    switch 128 {
    case 4294967295:
        v5[-1] = 604
    case 127:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        fmt.Println(v6)
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 int32 = -1
        v2 = 81
        for fi10 := 0; fi10 < 19; fi10++ {
            fmt.Println(v7)
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    v3 = append(v3, -821)
    if (!(!true)) {
        v11 := GenS1{F0: 0, F1: 2147483647}
        if len(v3) > 0 {
            v3[0] = 42
        }
        var v12 bool = ((!true) && true)
        _ = v11
        _ = v12
    } else {
        var v13 bool = (v6 >= v6)
        v13 = v13
        _ = v13
    }
    var v14 int = -357
    var v15 float64 = (float64(0.1) - 1e10)
    v6 = (-(v6))
    v16 := GenS1{F0: 925, F1: 543}
    v17 := GenS1{F0: 941, F1: 444}
    var v18 int64 = v6
    v7 += v7
    v2 -= (v2 * uint(v7))
    fmt.Println(v2, v18)
    for fi19 := 0; fi19 < 16; fi19++ {
        _ = len(v1)
        if (true || (0 > 0)) {
            v18 = -75
        } else {
            for fi20 := 0; fi20 < 14; fi20++ {
                var v21 float32 = (1e10 + -2.25)
                _ = fi20
                _ = v21
            }
        }
        _ = fi19
    }
    v22 := GenS2{F0: "a"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
