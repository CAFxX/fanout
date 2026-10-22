package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int16; F1 int64 }

func main() {
    var v1 int8 = 14
    var v2 uint64 = 2
    v3 := make(map[int]bool)
    v4 := make(map[int]int)
    v5 := make(map[int]string)
    v6 := []int{1, -36, 32767, -164}
    if len(v6) > 0 {
        v6[0] = -971
    }
    var v7 int32 = 64
    v1 = v1
    for fi8 := 0; fi8 < 7; fi8++ {
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 int = -937
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v1 = (v1 - -1)
    v11 := GenS1{F0: 127, F1: 4294967295}
    v1 = v1
    var v12 float32 = 1.5
    v1 = v1
    v3[128] = true
    var v13 int16 = 1000
    v6 = append(v6, 42)
    v7 += v7
    if false {
        v14 := GenS2{F0: v13, F1: -9223372036854775808}
        fmt.Println(v13, v12)
        v1 -= (^(v1))
        _ = v14
    } else {
        v15 := GenS1{F0: 128, F1: 742}
        _ = v15
    }
    var v16 int = -746
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
