package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
func genFn0(p0 uint8, p1 bool) int8 {
    if p1 {
        for fi7 := 0; fi7 < 13; fi7++ {
            for fi8 := 0; fi8 < 17; fi8++ {
                p0 -= (p0 & p0)
                _ = fi8
            }
            _ = fi7
        }
    }
    var v9 string = ("\n\t" + "a")
    if (1.5 <= 1.5) {
        if p1 {
            p1 = p1
            v9 = v9
            v10 := []int{128}
            _ = v10
        }
        fmt.Println(v9, p0, p0)
    } else {
        var v11 bool = p1
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v9
    return 127
}
func genFn1(p0 uint, p1 int64, p2 int16) uint64 {
    v12 := GenS1{F0: p0, F1: true}
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 1
    v2 := make(map[int]int)
    var v3 uint8 = 127
    v4 := []uint32{2147483648, (v1 * v1), v1}
    var v5 bool = true
    v6 := make(map[string]uint32)
    v5 = v5
    fmt.Println(v1)
    v1 -= v1
    v5 = v5
    if ((!v5) && (-1 > 42)) {
        if ("\n\t" != "") {
            for fi13 := 0; fi13 < 20; fi13++ {
                var v14 uint32 = v1
                _ = fi13
                _ = v14
            }
            v15 := GenS1{F0: 18446744073709551615, F1: (float32((0.0 / float64((3.14159 + -2.25)))) < (-2.25 / (0.1 * 1e10)))}
            _ = v15
        }
    } else {
        v5 = false
        var v16 uint32 = 42
        if ((42 == -128) && v5) {
            var v17 uint16 = 2
            _ = v17
        }
        _ = v16
    }
    v1 -= (v1 + 2)
    v3 = 191
    var v18 bool = v5
    switch ("hello" + ((((("hello" + "") + ("a" + "hello")) + "") + "") + ("" + ""))) {
    case "":
        var v19 int = 65
        _ = v19
    default:
        var v20 int8 = genFn0(184, v18)
        _ = v20
    }
    v6[""] = v1
    fmt.Println(v1, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
