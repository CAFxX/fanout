package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint }
type GenS2 struct { F0 bool; F1 uint; F2 int8 }
func genFn0(p0 int32) string {
    for fi5 := 0; fi5 < 1; fi5++ {
        for fi6 := 0; fi6 < 2; fi6++ {
            p0 += 42
            _ = fi6
        }
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 uint64 = 4294967295
            _ = fi7
            _ = v8
        }
        p0 += p0
        _ = fi5
    }
    v9 := []int{452, -2147483648}
    for fi10 := 0; fi10 < 3; fi10++ {
        v9 = append(v9, 848)
        _ = fi10
    }
    var v11 uint8 = 42
    _ = p0
    _ = v11
    _ = v9
    return ""
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{0, 42}
    var v3 int16 = 42
    v4 := make(map[int]string)
    if len(v2) > 0 {
        v2[0] = 527
    }
    v12 := GenS1{F0: 2147483647, F1: -834, F2: 4294967295}
    v13 := GenS1{F0: 981, F1: -1, F2: 267}
    v14 := GenS2{F0: (0 <= -1), F1: 18446744073709551615, F2: -128}
    _ = len(v2)
    var v15 float32 = float32(0.0)
    v3 -= v3
    var v16 string = "\n\t"
    var v17 int16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
