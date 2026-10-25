package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int; F1 uint32; F2 int }
func genFn0(p0 uint64, p1 uint64) int8 {
    p0 -= (^(p0))
    var v4 int16 = -610
    var v5 string = ((("a" + "") + "") + ("foo bar" + "a"))
    switch 32767 {
    case -1:
        var v6 bool = true
        _ = v6
    default:
        v7 := GenS1{F0: p1, F1: -128}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 int8 = 127
    var v2 string = "foo bar"
    v3 := make(map[string]string)
    v3[v2] = ((v2 + (v2 + v2)) + v2)
    fmt.Println(v1)
    v8 := GenS2{F0: 128, F1: 97, F2: 127}
    v1 -= 0
    fmt.Println(v1, v2)
    var v9 int = 128
    if false {
        var v10 float32 = ((3.14159 / 0.1) * 3.14159)
        v11 := []int{(v9 ^ -9223372036854775808), v9}
        if len(v11) > 0 {
            v11[0] = v9
        }
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 9; fi12++ {
            var v13 int = (fi12 & (v9 * fi12))
            fmt.Println(v9, v13, v2)
            _ = fi12
            _ = v13
        }
        v3[v2] = ("foo bar" + v2)
        var v14 float32 = (1.5 + -2.25)
        _ = v14
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 uint32 = 2
        if false {
            var v17 bool = false
            _ = v17
        } else {
            var v18 int64 = -695
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
