package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 bool }
func genFn0(p0 int16, p1 uint64, p2 int) int8 {
    p2 = (p2 / ((-768) | 1))
    var v4 int8 = -72
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 127
}

func main() {
    var v1 bool = true
    var v2 float64 = 0.0
    v3 := []int{-550, 32768}
    v2 -= v2
    switch "foo bar" {
    case "":
        var v5 int8 = genFn0(32767, 462, 9223372036854775807)
        _ = v5
    default:
        var v6 string = ("a" + "a")
        _ = v6
    }
    switch 42 {
    case 4294967295:
        v7 := GenS1{F0: "\n\t", F1: 435}
        _ = v7
    case 310:
        v1 = true
    }
    fmt.Println(v1, v2)
    if (-984 > -1) {
        v8 := GenS1{F0: (("hello" + "") + "a"), F1: 1}
        _ = v8
    } else {
        fmt.Println(v2, v2)
        var v9 uint32 = 0
        v10 := GenS2{F0: (!(42 >= 4294967295))}
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: v1}
    var v12 uint64 = 0
    v13 := GenS1{F0: "a", F1: (v12 / ((4294967295) | 1))}
    var v14 string = "\n\t"
    _ = len(v3)
    v15 := GenS2{F0: v1}
    if true {
        fmt.Println(v12, v1)
        var v16 uint32 = 82
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
