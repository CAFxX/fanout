package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 int8 }
func genFn0(p0 uint16) int32 {
    v7 := []int{}
    v8 := GenS1{F0: 0}
    fmt.Println(p0)
    _ = p0
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint, p1 bool, p2 uint64) uint {
    var v9 string = ""
    p1 = (42 >= 546)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 107
}

func main() {
    var v1 bool = true
    v2 := make(map[int]uint32)
    var v3 float64 = float64(3.14159)
    v4 := make(map[int]string)
    var v5 int64 = 0
    v6 := []uint32{4294967295, 2, 65535}
    v1 = v1
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := GenS2{F0: 635, F1: -128}
        var v12 bool = (42 != -778)
        _ = fi10
        _ = v11
        _ = v12
    }
    if len(v6) > 0 {
        v6[0] = 4294967295
    }
    switch "\n\t" {
    case "a":
        v3 += ((v3 * -2.25) - v3)
    case "foo bar":
        fmt.Println(v3, v1)
    }
    _ = len(v6)
    v1 = v1
    v13 := v2[661]
    fmt.Println(v3, v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
