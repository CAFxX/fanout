package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS4 struct { F0 uint32; F1 uint16; F2 bool }
func genFn0(p0 int64) string {
    v5 := GenS1{F0: 4294967295}
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int32 = -1
        if (false || (!(-42 < 42))) {
            var v8 int32 = (v7 * v7)
            v7 += 0
            var v9 int32 = 42
            _ = v8
            _ = v9
        } else {
            var v10 float32 = 1e10
            _ = v10
        }
        var v11 int64 = -1
        _ = fi6
        _ = v11
        _ = v7
    }
    _ = p0
    _ = v5
    return "a"
}
func genFn1(p0 int64) uint16 {
    p0 = p0
    p0 = p0
    _ = p0
    return 32767
}

func main() {
    var v1 int64 = 969
    var v2 uint8 = 255
    var v3 int8 = 1
    var v4 int8 = (v3 << ((103) & 63))
    switch 42 {
    case 4294967295:
        v2 += v2
    case 32767:
        switch int(v4) {
        case -80:
            fmt.Println(v1, v3)
        case 0:
            v12 := GenS1{F0: 547}
            _ = v12
        default:
            v2 = 206
        }
    default:
        v13 := []int{9223372036854775807, 719}
        _ = v13
    }
    if false {
        v14 := GenS1{F0: 1}
        for fi15 := 0; fi15 < 12; fi15++ {
            v16 := []int{}
            var v17 int8 = -128
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 float32 = 1e10
        _ = v14
        _ = v18
    } else {
        v19 := GenS2{F0: true}
        var v20 int64 = -9223372036854775808
        if (!false) {
            v21 := GenS3{F0: uint64(v2), F1: 4294967295, F2: 904}
            _ = v21
        }
        _ = v19
        _ = v20
    }
    v22 := GenS4{F0: 314, F1: 32767, F2: false}
    v23 := []int{100, -625}
    var v24 string = "a"
    v2 += v2
    switch ((((v24 + "foo bar") + v24) + (v24 + v24)) + "") {
    case "hello":
        var v25 uint8 = v2
        _ = v25
    case "foo bar":
        fmt.Println(v24, v1)
    case "":
        v2 -= v2
    }
    v4 -= (v3 << ((0) & 63))
    var v26 uint8 = 2
    _ = len(v23)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
}
