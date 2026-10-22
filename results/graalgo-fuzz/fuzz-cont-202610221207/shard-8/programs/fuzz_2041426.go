package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint }
type GenS2 struct { F0 uint32; F1 int; F2 uint32 }
type GenS3 struct { F0 uint; F1 uint8; F2 string }
type GenS4 struct { F0 string; F1 bool }
func genFn0(p0 string, p1 uint32, p2 bool) uint32 {
    if p2 {
        switch -286 {
        case -43:
            v7 := GenS1{F0: -9223372036854775808, F1: p1, F2: 895}
            _ = v7
        }
        v8 := GenS2{F0: p1, F1: 1, F2: p1}
        if p2 {
            p0 = p0
        } else {
            fmt.Println(p0)
            var v9 uint = 4294967295
            p0 = (("" + (p0 + p0)) + "foo bar")
            _ = v9
        }
        _ = v8
    }
    p0 = (p0 + (p0 + p0))
    v10 := GenS3{F0: 421, F1: 0, F2: ((p0 + "hello") + p0)}
    fmt.Println(p2, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return p1
}
func genFn1() uint8 {
    switch 1 {
    case 458:
        var v11 uint8 = 255
        _ = v11
    case 4294967295:
        var v12 int8 = -4
        _ = v12
    case 468:
        var v13 bool = (!false)
        _ = v13
    }
    var v14 uint64 = 1
    fmt.Println(v14)
    v14 = v14
    _ = v14
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    var v4 uint64 = 4294967295
    var v5 bool = true
    var v6 int64 = 1
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 uint16 = 1
        v6 = (v6 ^ -9223372036854775808)
        fmt.Println(v5)
        _ = fi15
        _ = v16
    }
    var v17 int16 = -968
    v18 := []int{-2147483648, 133, 128}
    v19 := GenS4{F0: "hello", F1: (1e10 == 1.5)}
    _ = len(v2)
    var v20 int64 = 42
    for fi21 := 0; fi21 < 15; fi21++ {
        v4 += v4
        v20 = (-(v20))
        _ = fi21
    }
    switch 42 {
    case -608:
        switch -476 {
        case 32767:
            var v22 int = -299
            _ = v22
        case 0:
            v20 = (-(v6))
        case -1:
            v17 = v17
        default:
            var v23 float32 = 0.1
            _ = v23
        }
    case 0:
        fmt.Println(v4, v4, v17)
    default:
        v6 = v20
    }
    var v24 int8 = -1
    switch -1 {
    case 42:
        switch "a" {
        case "foo bar":
            var v25 float64 = 3.14159
            _ = v25
        case "":
            var v26 uint16 = 255
            _ = v26
        case "\n\t":
            delete(v1, 65535)
        }
    default:
        if (v5 && false) {
            v20 += v20
            fmt.Println(v17)
            var v27 int16 = (v17 + (v17 & v17))
            _ = v27
        } else {
            var v28 float64 = 1e10
            var v29 int16 = (v17 ^ (v17 + v17))
            _ = v28
            _ = v29
        }
    }
    v17 = v17
    if v5 {
        fmt.Println(v24)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
