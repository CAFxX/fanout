package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int16; F1 uint8; F2 int32 }
func genFn0(p0 int16) int64 {
    v7 := []int{}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println(p0)
    _ = p0
    _ = v7
    return -239
}
func genFn1(p0 uint) int32 {
    v8 := GenS1{F0: 4}
    var v9 int8 = 100
    if true {
        v10 := GenS1{F0: v9}
        var v11 float32 = 1e10
        v12 := []int{32767}
        _ = v10
        _ = v11
        _ = v12
    }
    if (true && true) {
        var v13 uint32 = 245
        _ = v13
    }
    v9 += (-(v9))
    _ = p0
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 float32 = 1.5
    var v2 int32 = 0
    var v3 uint32 = 65535
    v4 := make(map[uint32]uint32)
    var v5 uint = 211
    var v6 uint64 = 1
    if false {
        v5 = 65535
    }
    if (919 < 42) {
        v14 := GenS2{F0: (v2 & (v2 - (v2 >> ((v5) & 63)))), F1: v5}
        v2 -= 0
        for fi15 := 0; fi15 < 1; fi15++ {
            fmt.Println(v3)
            _ = fi15
        }
        _ = v14
    }
    v16 := GenS3{F0: -487, F1: 104, F2: genFn1(v5)}
    fmt.Println(v2, v5)
    v17 := []int{2147483647, -895, 4294967295}
    v5 += (v5 % ((v5) | 1))
    if false {
        fmt.Println(v1, v6)
        switch 32767 {
        case -459:
            var v18 uint = v5
            _ = v18
        }
        var v19 int = 747
        _ = v19
    }
    fmt.Println(v6, v5, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
