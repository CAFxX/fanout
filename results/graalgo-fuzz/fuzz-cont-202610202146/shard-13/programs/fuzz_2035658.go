package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint64 }
func genFn0(p0 int8, p1 uint8) string {
    fmt.Println(p0)
    fmt.Println(p1, p0)
    v4 := []int{-621, -1}
    _ = p0
    _ = p1
    _ = v4
    return "foo bar"
}

func main() {
    var v1 uint8 = 78
    v2 := make(map[int]uint32)
    var v3 int16 = 519
    fmt.Println(v1, v3)
    fmt.Println(v1)
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := GenS1{F0: 2147483648, F1: (fi5 < fi5), F2: 0}
        var v7 int8 = 93
        _ = fi5
        _ = v6
        _ = v7
    }
    v1 = (v1 * 179)
    v2[4294967295] = 30
    v1 += (v1 & (v1 - v1))
    v1 = (^(v1))
    if (true && (-128 >= -128)) {
        var v8 int8 = -1
        _ = v8
    } else {
        var v9 int8 = 1
        _ = v9
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
