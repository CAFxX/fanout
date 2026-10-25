package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int }
func genFn0(p0 uint8) int32 {
    p0 -= 1
    p0 += 2
    switch -1 {
    case 32767:
        p0 -= (p0 ^ p0)
    case 0:
        p0 -= 255
    case -271:
        v5 := GenS1{F0: "hello", F1: 929, F2: int(p0)}
        _ = v5
    default:
        p0 += 127
    }
    var v6 uint = 255
    _ = p0
    _ = v6
    return -1
}

func main() {
    var v1 string = "\n\t"
    var v2 int8 = -128
    var v3 float64 = 3.14159
    v4 := make(map[string]bool)
    var v7 uint = 255
    v2 -= v2
    v8 := []int{127, -1}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 int8 = (-(v2))
        _ = fi9
        _ = v10
    }
    var v11 uint = 1
    v7 -= (v11 ^ 2)
    v8 = append(v8, 1)
    v2 -= (v2 / (((v2 / ((v2) | 1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
