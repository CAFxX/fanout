package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int8 }
func genFn0(p0 string, p1 int8, p2 uint16) string {
    p1 -= (p1 % ((0) | 1))
    for fi4 := 0; fi4 < 16; fi4++ {
        v5 := GenS1{F0: 698}
        p0 = (((p0 + p0) + p0) + (("\n\t" + ("foo bar" + ((p0 + (("foo bar" + "foo bar") + p0)) + "hello"))) + "foo bar"))
        _ = fi4
        _ = v5
    }
    p0 = p0
    v6 := GenS2{F0: 1, F1: (p1 * p1)}
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "a"
}

func main() {
    v1 := []int{-1, 2147483647}
    var v2 int8 = 1
    var v3 float64 = -2.25
    for fi8 := 0; fi8 < 19; fi8++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi9 := 0; fi9 < 3; fi9++ {
            v3 += v3
            v1 = append(v1, fi9)
            _ = fi9
        }
        _ = fi8
    }
    v3 = v3
    _ = len(v1)
    v3 = 0.1
    var v10 float64 = v3
    for fi11 := 0; fi11 < 8; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            if len(v1) > 0 {
                v1[0] = (fi11 + fi12)
            }
            v13 := GenS2{F0: 42, F1: v2}
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    var v14 float32 = ((0.0 - 0.0) * 1.5)
    if false {
        var v15 float64 = 0.0
        _ = v15
    } else {
        var v16 uint8 = 127
        switch 306 {
        case -9223372036854775808:
            v16 += v16
        case -986:
            v14 -= (v14 + float32(v3))
        }
        _ = v16
    }
    var v17 string = ""
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v14
    _ = v17
    _ = v2
    _ = v3
}
