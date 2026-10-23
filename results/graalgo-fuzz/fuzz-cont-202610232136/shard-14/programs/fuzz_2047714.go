package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 bool }
func genFn0(p0 int) int8 {
    p0 += p0
    p0 -= (p0 + (^(p0)))
    var v7 int64 = -1
    if (false && (((!false) && false) && true)) {
        v8 := []int{(-(p0)), (p0 + p0)}
        _ = v8
    }
    _ = p0
    _ = v7
    return -3
}
func genFn1(p0 uint64, p1 uint64, p2 bool) uint {
    switch "hello" {
    case "foo bar":
        var v9 uint = 4294967295
        _ = v9
    case "":
        switch 4294967295 {
        case 354:
            v10 := []int{32767, 127, 42, 127}
            _ = v10
        }
    case "hello":
        v11 := []int{}
        _ = v11
    default:
        fmt.Println(p0, p2, p0)
    }
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    var v1 uint16 = 868
    v2 := []string{"a"}
    var v3 int32 = -1
    v4 := make(map[string]bool)
    var v5 float64 = 3.14159
    var v6 int8 = 127
    var v12 uint64 = 740
    v13 := GenS1{F0: 4294967295, F1: -1, F2: true}
    fmt.Println(v3)
    fmt.Println(v6)
    v12 += 2
    v12 = v12
    v6 -= v6
    fmt.Println(v12)
    var v14 int32 = (v3 & (v3 >> ((genFn1(v12, v12, (!(0 >= 42)))) & 63)))
    v12 -= v12
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
