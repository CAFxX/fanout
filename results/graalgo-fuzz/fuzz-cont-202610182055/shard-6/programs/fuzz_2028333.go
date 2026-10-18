package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 int16; F2 uint32 }
func genFn0(p0 uint32, p1 uint16) int {
    var v7 int8 = -106
    v7 += (v7 % (((-(v7))) | 1))
    _ = p0
    _ = p1
    _ = v7
    return 127
}
func genFn1(p0 int64, p1 int8, p2 int8) bool {
    v8 := []int{-601, 32767}
    switch "a" {
    case "foo bar":
        var v9 int8 = (-(p2))
        _ = v9
    default:
        var v10 int16 = -526
        _ = v10
    }
    v11 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v8
    return false
}

func main() {
    var v1 float64 = -2.25
    var v2 uint8 = 127
    v3 := []int{32767, 32767}
    var v4 uint = 1
    var v5 int8 = 42
    var v6 uint8 = 1
    v12 := GenS1{F0: 2}
    var v13 string = "a"
    fmt.Println(v6, v2)
    v14 := GenS2{F0: -1, F1: -836, F2: 2}
    var v15 uint32 = 0
    for fi16 := 0; fi16 < 18; fi16++ {
        fmt.Println(v13, v5, v4)
        _ = fi16
    }
    v17 := GenS1{F0: 2}
    v18 := GenS2{F0: 2147483647, F1: 0, F2: 359}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
