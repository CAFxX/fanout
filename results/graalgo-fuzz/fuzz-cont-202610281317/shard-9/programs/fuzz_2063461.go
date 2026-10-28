package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int64; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int64, p1 int32, p2 uint32) uint8 {
    v5 := []int{128, 32767, -427}
    var v6 float32 = float32(-2.25)
    var v7 uint = 0
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if true {
        var v8 string = "hello"
        _ = v8
    } else {
        if false {
            v9 := GenS1{F0: -128, F1: (p1 * (-(p1))), F2: p2}
            v10 := GenS2{F0: uint8(p2), F1: 1}
            _ = v10
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 255
}
func genFn1(p0 uint64, p1 int32) int64 {
    var v11 int16 = 1000
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 int64 = int64(p1)
        var v14 int64 = -1
        _ = fi12
        _ = v13
        _ = v14
    }
    switch 42 {
    case -307:
        switch "foo bar" {
        case "foo bar":
            v15 := []int{38}
            _ = v15
        case "":
            fmt.Println(v11)
        case "a":
            var v16 uint32 = 0
            _ = v16
        default:
            fmt.Println(v11, p1, v11)
        }
    default:
        var v17 int64 = int64(p0)
        _ = v17
    }
    fmt.Println(p0, p1)
    for fi18 := 0; fi18 < 10; fi18++ {
        v19 := GenS1{F0: -92, F1: (p1 << ((824) & 63)), F2: 4294967295}
        var v20 float64 = 0.0
        _ = fi18
        _ = v19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v11
    return -9223372036854775808
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]string)
    v4 := make(map[int]uint32)
    var v21 int32 = 0
    v22 := GenS2{F0: 255, F1: -128}
    v21 = 2147483647
    v23 := GenS1{F0: -1, F1: (v21 % (((v21 / ((v21) | 1))) | 1)), F2: 2}
    switch 968 {
    case 9223372036854775807:
        v24 := GenS3{F0: genFn1(333, 42), F1: 127, F2: 0}
        _ = v24
    case -1:
        var v25 float32 = -2.25
        _ = v25
    default:
        if true {
            v3[1] = ("\n\t" + "foo bar")
            v26 := GenS4{F0: 2}
            var v27 string = "a"
            _ = v26
            _ = v27
        } else {
            var v28 uint8 = 127
            v29 := []int{350, 865, 989}
            for fi30 := 0; fi30 < 2; fi30++ {
                v31 := GenS1{F0: -40, F1: int32(fi30), F2: 0}
                _ = fi30
                _ = v31
            }
            _ = v28
            _ = v29
        }
    }
    var v32 uint = 255
    for fi33 := 0; fi33 < 4; fi33++ {
        v1 = 65535
        _ = fi33
    }
    fmt.Println(v21, v1)
    v34 := GenS3{F0: 453, F1: 100, F2: 481}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v32
    _ = v34
    _ = v4
}
