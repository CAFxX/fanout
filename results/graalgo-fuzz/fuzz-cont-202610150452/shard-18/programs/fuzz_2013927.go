package main

import (
    "fmt"
)

func genFn0(p0 int) uint32 {
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 int = (fi6 % ((2147483647) | 1))
        fmt.Println(fi6, fi6, p0)
        var v8 bool = false
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int16 = -1
    _ = p0
    _ = v9
    return 570
}

func main() {
    var v1 float64 = (float64(0.1) - 0.0)
    v2 := make(map[int]uint32)
    v3 := []int{}
    v4 := []int{-844, 233}
    v5 := make(map[uint32]string)
    fmt.Println(v1)
    var v10 int8 = 100
    v10 = 75
    v11 := v5[0]
    var v12 int32 = 890
    var v13 int = 32768
    v12 = (v12 % ((v12) | 1))
    var v14 uint32 = 2
    for fi15 := 0; fi15 < 3; fi15++ {
        v16 := v2[(fi15 ^ fi15)]
        _ = fi15
        _ = v16
    }
    if len(v3) > 0 {
        v3[0] = (v13 >> ((272) & 63))
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 float32 = 0.0
        var v19 int32 = (v12 % (((v12 - (v12 * (v12 ^ (v12 | (v12 % ((v12) | 1))))))) | 1))
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
