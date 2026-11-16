package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint32; F1 bool; F2 bool }

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 2147483647
    v3 := make(map[int]bool)
    v4 := []uint32{0, 521, 42}
    v5 := make(map[uint32]int)
    var v6 int32 = -411
    fmt.Println(v6, v6)
    fmt.Println(v6)
    if (0 != 255) {
        var v7 int64 = 1
        v8 := GenS1{F0: -1, F1: 255}
        _ = v7
        _ = v8
    }
    _ = len(v5)
    v9 := GenS2{F0: 829, F1: false, F2: (1 == 798)}
    if false {
        _ = len(v3)
        var v10 string = "\n\t"
        _ = v10
    } else {
        for fi11 := 0; fi11 < 3; fi11++ {
            for fi12 := 0; fi12 < 12; fi12++ {
                fmt.Println(fi12, fi11)
                fmt.Println(fi12)
                v6 = (-(v6))
                _ = fi12
            }
            var v13 bool = false
            v14 := GenS1{F0: v2, F1: 42}
            _ = fi11
            _ = v13
            _ = v14
        }
    }
    var v15 float32 = 3.14159
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 int = (fi16 % (((v2 << ((31) & 63))) | 1))
        v15 = 0.0
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        if ((!((true || (fi18 < (v2 - fi18))) && ((fi18 + fi18) != fi18))) && false) {
            var v19 uint64 = 0
            fmt.Println(v6, v19, v15)
            _ = v19
        }
        _ = fi18
    }
    v15 = (v15 * v15)
    var v20 int16 = 156
    v21 := v3[int(v20)]
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        _ = len(v3)
        v1[693] = (("" + "") + "\n\t")
        _ = fi22
    }
    fmt.Println(v21, v21, v20)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
