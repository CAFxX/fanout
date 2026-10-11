package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint64 }
func genFn0(p0 bool, p1 int16, p2 string) int {
    switch "" {
    case "":
        if (1.5 > 1e10) {
            var v6 int32 = -285
            p0 = (-324 >= 42)
            p1 += 917
            _ = v6
        } else {
            v7 := GenS1{F0: 2}
            _ = v7
        }
    case "hello":
        p1 = p1
    }
    p0 = false
    var v8 float64 = (0.1 / (-2.25 * 1e10))
    v9 := []int{836, -1}
    _ = len(v9)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return -1
}
func genFn1(p0 int32, p1 uint16) uint8 {
    var v10 uint8 = 1
    var v11 uint = 627
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return (^(v10))
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{984, 1, -395, 551}
    var v3 uint8 = 1
    var v4 uint16 = 874
    var v5 uint16 = (^(v4))
    if (0.1 > 3.14159) {
        for fi12 := 0; fi12 < 15; fi12++ {
            v4 = 65535
            v1[(-(fi12))] = 65535
            _ = fi12
        }
    } else {
        v3 -= (v3 + v3)
        fmt.Println(v4, v5)
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        v14 := GenS1{F0: 79}
        var v15 uint16 = v5
        _ = fi13
        _ = v14
        _ = v15
    }
    if (1 != 509) {
        var v16 int16 = -1
        for fi17 := 0; fi17 < 3; fi17++ {
            var v18 float32 = 0.1
            _ = fi17
            _ = v18
        }
        _ = v16
    } else {
        var v19 int16 = 32767
        switch -1 {
        case 0:
            var v20 int = 9223372036854775807
            _ = v20
        case 128:
            v21 := GenS2{F0: 652, F1: 42}
            _ = v21
        default:
            v3 -= v3
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v19
    }
    v22 := GenS2{F0: v4, F1: -128}
    switch "a" {
    case "hello":
        _ = len(v2)
    case "\n\t":
        v5 = v4
    default:
        v23 := GenS3{F0: v4, F1: -209, F2: 1}
        _ = v23
    }
    v4 += v5
    switch "a" {
    case "a":
        v24 := GenS3{F0: v5, F1: 42, F2: 2}
        _ = v24
    default:
        v5 = 2
    }
    if (false && false) {
        var v25 bool = true
        _ = v25
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v26 := GenS3{F0: 65535, F1: 2147483647, F2: 42}
    var v27 int = 32767
    var v28 uint32 = 2
    if len(v2) > 0 {
        v2[0] = v27
    }
    if true {
        var v29 int64 = -9223372036854775808
        _ = v29
    }
    v3 = 0
    var v30 uint = 18446744073709551615
    delete(v1, (v27 + (-(v27))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
