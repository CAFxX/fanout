package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint32 }
func genFn0() int16 {
    var v5 int = -988
    v5 = (-(v5))
    fmt.Println(v5)
    v6 := []int{}
    _ = v5
    _ = v6
    return -655
}
func genFn1(p0 string, p1 int64, p2 bool) uint8 {
    v7 := GenS1{F0: 42, F1: 194, F2: 100}
    p2 = (false && true)
    v8 := []int{560, -2147483648, 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 175
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[string]uint32)
    var v3 int32 = 42
    var v4 int32 = (v3 / ((-2147483648) | 1))
    fmt.Println(v4)
    v4 += -1
    v9 := []int{}
    var v10 string = ""
    switch -1 {
    case 4294967295:
        delete(v2, v10)
    }
    _ = len(v2)
    var v11 int64 = 1
    delete(v2, ("" + v10))
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := GenS2{F0: 2}
        var v14 float64 = 0.1
        _ = len(v9)
        _ = fi12
        _ = v13
        _ = v14
    }
    switch "\n\t" {
    case "":
        var v15 uint = 43
        _ = v15
    }
    switch "hello" {
    case "foo bar":
        var v16 int64 = v11
        _ = v16
    case "\n\t":
        var v17 int8 = -60
        _ = v17
    default:
        v4 = v3
    }
    _ = len(v9)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
