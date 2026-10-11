package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint16; F1 uint8; F2 string }
func genFn0(p0 uint8) int32 {
    p0 = (p0 * (p0 | p0))
    var v6 bool = true
    fmt.Println(p0, p0)
    v7 := GenS1{F0: -349, F1: true}
    _ = p0
    _ = v6
    _ = v7
    return -1
}
func genFn1() uint64 {
    if false {
        var v8 int64 = 537
        if true {
            v8 -= v8
            var v9 float32 = 1.5
            _ = v9
        }
        _ = v8
    }
    var v10 int64 = -1
    v11 := []int{}
    _ = v10
    _ = v11
    return 2
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -2147483648
    v3 := []int{127, 32768, -2147483648}
    var v4 int64 = -1
    v5 := make(map[uint32]string)
    v12 := GenS2{F0: 1, F1: 643}
    var v13 bool = false
    for fi14 := 0; fi14 < 17; fi14++ {
        fmt.Println(fi14)
        v4 = 366
        _ = fi14
    }
    v4 += v4
    v15 := GenS3{F0: 42, F1: 121, F2: "hello"}
    var v16 int64 = v4
    v17 := GenS3{F0: 358, F1: 156, F2: "hello"}
    v16 += (v16 / (((-(v4))) | 1))
    v2 = int32(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
