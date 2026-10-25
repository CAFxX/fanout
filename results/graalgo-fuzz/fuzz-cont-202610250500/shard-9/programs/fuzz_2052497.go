package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
func genFn0(p0 int8) int {
    if (513 >= 350) {
        var v6 string = ""
        if (!(2 != 32767)) {
            v7 := GenS1{F0: 128, F1: 0}
            _ = v7
        }
        v6 = v6
        _ = v6
    }
    p0 = (p0 >> ((491) & 63))
    fmt.Println(p0)
    _ = p0
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]int)
    var v2 string = (("a" + "") + ("" + "foo bar"))
    v3 := make(map[string]string)
    var v4 uint = 18446744073709551615
    var v5 int16 = 32767
    v4 = v4
    var v8 int16 = (v5 * v5)
    var v9 int16 = -295
    v10 := []int{}
    switch genFn0(0) {
    case 32767:
        v11 := GenS1{F0: 115, F1: 798}
        _ = v11
    case 486:
        fmt.Println(v8)
    case -509:
        v2 = (v2 + ((v2 + v2) + v2))
    }
    var v12 float64 = ((((1.5 * 1.5) * 1.5) + 0.1) / 3.14159)
    fmt.Println(v8, v12, v2)
    if len(v10) > 0 {
        v10[0] = 67
    }
    v2 = v2
    var v13 uint64 = 86
    v14 := v3[("hello" + "a")]
    switch 0 {
    case 0:
        v5 = v8
    default:
        v5 -= (v9 % ((v5) | 1))
    }
    var v15 uint8 = 172
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
