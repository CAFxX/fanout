package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
func genFn0() int16 {
    var v5 float32 = float32((0.1 / -2.25))
    v6 := GenS1{F0: 18446744073709551615, F1: "foo bar"}
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 int8 = -128
    var v2 int = 313
    var v3 uint = 255
    v4 := make(map[string]uint32)
    var v7 int8 = (v1 % ((v1) | 1))
    v3 = 4294967295
    v1 = 127
    var v8 bool = false
    v9 := GenS1{F0: 4294967295, F1: ""}
    v7 = v1
    v7 = (v7 - v1)
    v10 := []int{}
    v3 -= 419
    if (18446744073709551615 <= 4294967295) {
        if len(v10) > 0 {
            v10[0] = (v2 / ((v2) | 1))
        }
    } else {
        v1 = (v1 & v1)
        v8 = ((v1 - (v1 << (((v3 >> ((v3) & 63))) & 63))) < (v1 & int8(v3)))
        var v11 uint = v3
        _ = v11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := GenS1{F0: 2, F1: "foo bar"}
        var v14 float32 = 1.5
        _ = fi12
        _ = v13
        _ = v14
    }
    if v8 {
        fmt.Println(v1, v7)
        var v15 string = "foo bar"
        _ = v15
    } else {
        fmt.Println(v8, v1, v7)
        v7 += 127
        fmt.Println(v1, v1, v8)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
