package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int64; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint64; F1 uint64; F2 int8 }
func genFn0(p0 uint8) uint {
    if true {
        p0 += (p0 % ((1) | 1))
        v4 := GenS1{F0: 42, F1: -558}
        for fi5 := 0; fi5 < 1; fi5++ {
            for fi6 := 0; fi6 < 2; fi6++ {
                p0 -= (p0 / ((p0) | 1))
                var v7 float64 = 1e10
                v8 := []int{(fi5 / ((fi5) | 1)), fi6, (fi6 ^ (fi5 / (((-(fi6))) | 1)))}
                _ = fi6
                _ = v7
                _ = v8
            }
            _ = fi5
        }
        _ = v4
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        v10 := []int{fi9, (fi9 % ((fi9) | 1)), fi9}
        _ = fi9
        _ = v10
    }
    if true {
        var v11 string = "foo bar"
        _ = v11
    } else {
        var v12 bool = (0 > 4294967295)
        _ = v12
    }
    _ = p0
    return 2
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint32 = 2
    v3 := make(map[int]string)
    if true {
        var v13 uint32 = (v2 * (v2 + v2))
        switch -1 {
        case 0:
            var v14 string = ""
            _ = v14
        }
        if (!true) {
            v15 := []int{4294967295}
            var v16 bool = false
            _ = v15
            _ = v16
        }
        _ = v13
    }
    v17 := GenS1{F0: -1, F1: 743}
    var v18 uint64 = 2
    for fi19 := 0; fi19 < 3; fi19++ {
        fmt.Println(fi19, v2, v1)
        _ = fi19
    }
    switch -935 {
    case 2147483647:
        if false {
            var v20 int16 = -1
            v21 := GenS2{F0: 461, F1: 255, F2: 427}
            _ = v20
            _ = v21
        }
    case 127:
        v1 -= (v18 << ((255) & 63))
    }
    switch -962 {
    case 1:
        v2 = v2
    case 127:
        var v22 float32 = 0.1
        _ = v22
    case -9223372036854775808:
        var v23 uint16 = 793
        _ = v23
    }
    v1 = 1
    for fi24 := 0; fi24 < 15; fi24++ {
        switch "" {
        case "hello":
            v25 := GenS1{F0: 0, F1: 0}
            _ = v25
        case "a":
            v18 = (v1 & (v1 | v18))
        }
        switch fi24 {
        case 447:
            v18 -= (v1 << ((489) & 63))
        case 9223372036854775807:
            var v26 int64 = 1
            _ = v26
        default:
            v2 -= (v2 << ((0) & 63))
        }
        _ = fi24
    }
    v1 += v1
    var v27 uint32 = 42
    fmt.Println(v2, v18, v1)
    v28 := GenS3{F0: (v18 ^ 0), F1: v1, F2: -73}
    var v29 float32 = float32(3.14159)
    var v30 string = "foo bar"
    v31 := []int{-840, -1}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
