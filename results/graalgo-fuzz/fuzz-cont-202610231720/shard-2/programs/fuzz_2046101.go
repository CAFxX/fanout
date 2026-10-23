package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 uint }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 848
    var v3 int8 = -128
    v4 := []int{4294967295}
    v5 := make(map[uint32]int)
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint64 = (v2 + (v2 - 0))
        switch fi6 {
        case -2147483648:
            var v8 int8 = v3
            _ = v8
        }
        v5[65535] = (fi6 ^ 2)
        _ = fi6
        _ = v7
    }
    if false {
        var v9 int64 = 9223372036854775807
        v10 := GenS1{F0: (255 >= 165), F1: v9}
        var v11 int32 = 42
        _ = v10
        _ = v11
        _ = v9
    } else {
        v12 := v1[774]
        v13 := GenS2{F0: 263, F1: -797}
        var v14 uint64 = 9
        _ = v12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        v16 := GenS3{F0: 127, F1: 42}
        var v17 int32 = 0
        _ = fi15
        _ = v16
        _ = v17
    }
    switch 2147483647 {
    case 1:
        for fi18 := 0; fi18 < 19; fi18++ {
            delete(v5, 2147483647)
            _ = fi18
        }
    case 32768:
        fmt.Println(v2, v3)
    }
    v3 -= (v3 + 127)
    fmt.Println(v3)
    if true {
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 float64 = 1.5
            _ = fi19
            _ = v20
        }
        v3 = (v3 | v3)
    } else {
        var v21 uint64 = (v2 * (v2 - 972))
        fmt.Println(v2, v2)
        _ = v21
    }
    switch 0 {
    case 423:
        fmt.Println(v3, v2)
    case 9223372036854775807:
        fmt.Println(v2, v2)
    case 32767:
        var v22 string = ("" + ("\n\t" + ("\n\t" + ((("a" + "hello") + "foo bar") + "hello"))))
        _ = v22
    default:
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 uint = 689
            _ = len(v1)
            _ = fi23
            _ = v24
        }
    }
    _ = len(v5)
    v3 += (v3 % ((v3) | 1))
    v25 := v1[128]
    v26 := GenS4{F0: 330}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
