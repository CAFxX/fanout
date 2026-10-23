package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS2 struct { F0 int; F1 uint16 }
func genFn0() uint32 {
    var v4 uint64 = 18446744073709551615
    var v5 uint16 = 600
    fmt.Println(v5, v5)
    _ = v4
    _ = v5
    return 2
}
func genFn1() bool {
    v6 := GenS1{F0: 127, F1: 42, F2: 42}
    var v7 int8 = -53
    _ = v6
    _ = v7
    return (34 == 450)
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 int16 = 1
    var v3 int = 994
    var v8 uint32 = 1
    for fi9 := 0; fi9 < 4; fi9++ {
        v10 := []int{(fi9 | v3)}
        if len(v10) > 0 {
            v10[0] = 127
        }
        v8 = genFn0()
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 int32 = int32(v3)
        if true {
            v13 := []int{}
            _ = v13
        }
        _ = fi11
        _ = v12
    }
    v2 += v2
    v14 := GenS2{F0: 128, F1: 925}
    var v15 uint64 = 0
    var v16 int8 = 0
    v2 = (v2 / ((42) | 1))
    var v17 float64 = 1.5
    v18 := []int{1}
    v8 = (^(v8))
    v19 := GenS1{F0: 8, F1: 127, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v8
}
