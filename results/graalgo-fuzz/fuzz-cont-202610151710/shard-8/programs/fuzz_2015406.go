package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint }
func genFn0(p0 uint8, p1 int, p2 int8) int16 {
    v6 := GenS1{F0: 57, F1: 0}
    var v7 float32 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1000
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{}
    v3 := make(map[string]int)
    v4 := []int{128, 42, -789}
    v5 := make(map[uint32]bool)
    if len(v2) > 0 {
        v2[0] = 65535
    }
    v8 := GenS1{F0: 255, F1: 42}
    v9 := GenS1{F0: 42, F1: 2}
    var v10 int64 = -1
    if true {
        var v11 uint64 = 648
        v3["foo bar"] = -908
        _ = v11
    }
    v12 := GenS1{F0: 65535, F1: 42}
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 int32 = -761
        _ = fi13
        _ = v14
    }
    var v15 bool = (v10 == (v10 * (v10 ^ -1)))
    v16 := GenS2{F0: 42, F1: 1, F2: 720}
    delete(v1, "")
    v3["\n\t"] = 128
    for fi17 := 0; fi17 < 17; fi17++ {
        v15 = (0 == 2)
        v10 = 0
        _ = fi17
    }
    v10 = 585
    v10 = v10
    var v18 uint32 = 1
    var v19 float64 = -2.25
    var v20 float64 = v19
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v12
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
