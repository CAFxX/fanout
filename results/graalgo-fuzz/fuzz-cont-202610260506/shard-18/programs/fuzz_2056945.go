package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 bool }
type GenS2 struct { F0 uint8; F1 int8; F2 int16 }
type GenS3 struct { F0 int8; F1 int64 }
func genFn0(p0 int16) uint16 {
    v4 := GenS1{F0: 65535, F1: 241, F2: (65535 > 565)}
    if true {
        v5 := GenS1{F0: 42, F1: 58, F2: false}
        var v6 int64 = 401
        v7 := []int{-2147483648, 32767}
        _ = v5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 42, F1: 149, F2: (!false)}
    _ = p0
    _ = v4
    _ = v8
    return 32767
}
func genFn1(p0 int8, p1 int32, p2 bool) uint8 {
    v9 := []int{}
    v10 := GenS1{F0: 840, F1: 128, F2: (!p2)}
    v11 := GenS2{F0: 240, F1: p0, F2: 1000}
    switch 446 {
    case -36:
        v12 := GenS3{F0: p0, F1: -938}
        _ = v12
    default:
        if (("\n\t" + "") >= "") {
            v13 := GenS2{F0: 42, F1: p0, F2: 1}
            p2 = (!true)
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return 42
}

func main() {
    v1 := []int{32768}
    var v2 float32 = 1.5
    v3 := make(map[string]string)
    for fi14 := 0; fi14 < 13; fi14++ {
        if false {
            for fi15 := 0; fi15 < 3; fi15++ {
                var v16 uint = 2
                _ = fi15
                _ = v16
            }
            v17 := GenS1{F0: 582, F1: 2, F2: (false && true)}
            _ = v17
        }
        _ = fi14
    }
    v2 = (0.1 / (float32(3.14159) - (0.1 / 0.1)))
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 uint = 255
        fmt.Println(fi18, v2)
        _ = fi18
        _ = v19
    }
    switch -614 {
    case 4294967295:
        fmt.Println(v2)
    }
    if (((0.0 * 0.0) * -2.25) == 0.1) {
        var v20 int16 = 0
        _ = v20
    } else {
        v2 = 1e10
    }
    var v21 uint64 = 18446744073709551615
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 string = "hello"
        _ = fi22
        _ = v23
    }
    fmt.Println(v2, v2)
    v2 -= 0.0
    switch 9223372036854775807 {
    case 1:
        for fi24 := 0; fi24 < 7; fi24++ {
            delete(v3, "")
            v21 += (v21 * 0)
            _ = fi24
        }
    case -9223372036854775808:
        var v25 int32 = 1
        _ = v25
    case 2147483647:
        v2 += v2
    default:
        _ = len(v3)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
}
