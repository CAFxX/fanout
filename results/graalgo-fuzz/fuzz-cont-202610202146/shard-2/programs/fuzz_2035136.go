package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int; F1 uint32; F2 uint }
type GenS3 struct { F0 int64 }
func genFn0(p0 bool, p1 uint32) int16 {
    v6 := GenS1{F0: "a", F1: p1, F2: 578}
    p1 = 65535
    var v7 bool = ((32768 >= 32767) || true)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -32768
}
func genFn1(p0 uint16, p1 uint64, p2 uint64) uint32 {
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := []int{fi8, (fi8 | -637), fi8}
        var v10 uint16 = (p0 | p0)
        v11 := GenS2{F0: (fi8 ^ -30), F1: 42, F2: 4294967295}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := []int{}
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return 404
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 0
    var v3 uint32 = 4294967295
    v4 := []int{}
    var v5 int16 = 1
    fmt.Println(v2, v5)
    var v14 float32 = 0.0
    v5 -= (v5 | 32767)
    switch "a" {
    case "a":
        _ = len(v1)
    }
    v5 = (v5 & 581)
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := GenS3{F0: -989}
        _ = fi15
        _ = v16
    }
    var v17 bool = false
    switch -639 {
    case -1:
        var v18 int32 = 2147483647
        _ = v18
    case -2147483648:
        v19 := GenS2{F0: 622, F1: (v3 + 2), F2: 1}
        _ = v19
    default:
        var v20 uint8 = 1
        _ = v20
    }
    var v21 uint64 = 604
    if false {
        var v22 int32 = 0
        fmt.Println(v14)
        var v23 float32 = v14
        _ = v22
        _ = v23
    }
    switch 32767 {
    case -9223372036854775808:
        v21 = (v21 << ((42) & 63))
    }
    fmt.Println(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
