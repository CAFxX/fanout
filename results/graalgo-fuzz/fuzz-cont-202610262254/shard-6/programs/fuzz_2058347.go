package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int32 }
func genFn0() uint {
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := []int{(^(fi4)), (^(fi4)), fi4}
        _ = len(v5)
        var v6 int = (fi4 << ((4294967295) & 63))
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS1{F0: -1}
        var v9 int32 = 1
        v10 := []int{}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 uint = 1
    _ = v11
    return v11
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 151
    v3 := []int{32767, 128, -608}
    v12 := GenS1{F0: -1}
    v2 -= 132
    v2 = v2
    v2 = (v2 | 128)
    var v13 int16 = 235
    v14 := GenS1{F0: 1}
    var v15 int64 = 510
    v2 = v2
    v16 := GenS2{F0: (v13 >> ((42) & 63)), F1: 2147483647}
    fmt.Println(v13, v2, v15)
    v17 := GenS2{F0: -1, F1: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
