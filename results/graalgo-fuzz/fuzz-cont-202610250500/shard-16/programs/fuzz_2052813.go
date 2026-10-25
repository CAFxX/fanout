package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint64 }
func genFn0(p0 bool, p1 string, p2 uint) int32 {
    if true {
        v5 := []int{}
        if (-1 != -1) {
            var v6 float64 = 0.0
            v7 := GenS1{F0: (p2 | (p2 & (p2 | (p2 ^ (p2 ^ p2)))))}
            _ = v6
            _ = v7
        }
        _ = v5
    }
    fmt.Println(p2, p1)
    var v8 float64 = 0.1
    v9 := GenS1{F0: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 uint8 = 0
    var v2 bool = (!false)
    v3 := []int{706}
    var v4 uint16 = 32767
    v1 = v1
    var v10 int32 = -1
    fmt.Println(v10)
    v11 := GenS1{F0: 65535}
    v12 := GenS2{F0: 1, F1: 796}
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 uint16 = v4
        _ = fi13
        _ = v14
    }
    v4 = v4
    fmt.Println(v1, v10, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
