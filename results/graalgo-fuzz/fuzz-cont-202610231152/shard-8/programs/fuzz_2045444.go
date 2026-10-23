package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
func genFn0(p0 int8, p1 int32) uint8 {
    v7 := []int{32767}
    var v8 string = ""
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 127
}
func genFn1(p0 uint8) int64 {
    v9 := GenS1{F0: 87, F1: -242}
    var v10 float64 = -2.25
    _ = p0
    _ = v10
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]int)
    v3 := make(map[string]int)
    v4 := make(map[string]uint32)
    var v5 uint8 = 108
    var v6 uint64 = 531
    _ = len(v4)
    for fi11 := 0; fi11 < 11; fi11++ {
        fmt.Println(v5, v5)
        v12 := v4["a"]
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 int8 = -65
        _ = fi13
        _ = v14
    }
    v6 = (v6 - 42)
    _ = len(v3)
    fmt.Println(v5, v5)
    var v15 bool = true
    v15 = (v15 && v15)
    fmt.Println(v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
