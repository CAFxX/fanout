package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 int64 }

func main() {
    var v1 int32 = -1
    var v2 uint8 = 55
    var v3 int8 = -128
    var v4 uint32 = 0
    v5 := make(map[int]int)
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := GenS1{F0: v3}
            v9 := []int{(fi6 | fi6), (fi7 * (-(fi6))), fi7}
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 string = ""
        v11 := GenS2{F0: 42}
        _ = fi6
        _ = v10
        _ = v11
    }
    fmt.Println(v1, v4)
    v5[int(v1)] = 42
    v12 := GenS3{F0: (v1 >> ((1) & 63)), F1: -1}
    if (!((1 < 42) || false)) {
        v3 = (v3 + v3)
    } else {
        var v13 int = -78
        var v14 int32 = (v1 / ((v1) | 1))
        _ = v13
        _ = v14
    }
    var v15 uint64 = 42
    v1 = (v1 ^ 42)
    v16 := GenS1{F0: v3}
    var v17 uint = 0
    v2 = (^(v2))
    for fi18 := 0; fi18 < 20; fi18++ {
        v1 = (v1 << ((860) & 63))
        var v19 int8 = -1
        _ = fi18
        _ = v19
    }
    var v20 uint32 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
