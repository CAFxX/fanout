package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int, p1 uint32) int64 {
    p0 -= 565
    var v5 float32 = (0.1 + 3.14159)
    var v6 int8 = 0
    fmt.Println(v6)
    var v7 uint = 255
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint32, p1 int) int32 {
    fmt.Println(p0, p0)
    p1 += p1
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := GenS1{F0: 491, F1: 42}
        _ = fi8
        _ = v9
    }
    fmt.Println(p1, p0)
    switch (-(p1)) {
    case 32768:
        var v10 uint16 = 838
        _ = v10
    case 646:
        switch "hello" {
        case "\n\t":
            fmt.Println(p1)
        case "a":
            v11 := GenS2{F0: 100}
            _ = v11
        case "foo bar":
            var v12 int32 = 201
            _ = v12
        }
    default:
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 uint = 4294967295
            _ = fi13
            _ = v14
        }
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := []int{32767}
    v2 := make(map[int]bool)
    v3 := make(map[string]int)
    v4 := []string{"hello", "", "foo bar"}
    var v15 bool = (1000 < 1)
    for fi16 := 0; fi16 < 4; fi16++ {
        v15 = false
        _ = fi16
    }
    var v17 float64 = -2.25
    var v18 uint = 504
    var v19 int8 = -109
    v19 = v19
    if v15 {
        _ = len(v4)
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 int64 = 917
            _ = fi20
            _ = v21
        }
        v22 := GenS2{F0: -1}
        _ = v22
    }
    var v23 bool = (!true)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
