package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int32; F1 int8; F2 uint64 }
type GenS3 struct { F0 bool; F1 int; F2 uint }
type GenS4 struct { F0 int8; F1 string; F2 uint32 }
func genFn0(p0 uint, p1 int) uint {
    v6 := GenS1{F0: 1, F1: 42}
    if false {
        var v7 int32 = int32(p0)
        _ = v7
    }
    var v8 int = p1
    v9 := []int{v8, (p1 << ((p0) & 63)), 285, 0}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    _ = v9
    return (p0 * (p0 << ((p0) & 63)))
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint32 = 4294967295
    v3 := make(map[uint32]bool)
    v4 := []uint32{v2, v2, (v2 / ((42) | 1)), (v2 >> ((74) & 63))}
    v5 := make(map[int]string)
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 float64 = 1.5
        v12 := GenS1{F0: uint8(v2), F1: 127}
        v13 := GenS2{F0: 42, F1: 1, F2: (^(v1))}
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    _ = len(v4)
    if (0.0 > ((0.1 + 0.0) / 1e10)) {
        v1 += 0
    } else {
        for fi14 := 0; fi14 < 2; fi14++ {
            v15 := GenS2{F0: 56, F1: 42, F2: (v1 / (((v1 + v1)) | 1))}
            var v16 uint32 = 934
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 int16 = -697
        _ = v17
    }
    if false {
        v18 := GenS3{F0: ("foo bar" > "a"), F1: 9223372036854775807, F2: genFn0(119, 42)}
        _ = v18
    }
    var v19 uint8 = 108
    _ = len(v4)
    v2 = (^(v2))
    for fi20 := 0; fi20 < 10; fi20++ {
        switch ("" + "hello") {
        case "hello":
            v2 = v2
        case "\n\t":
            var v21 int = fi20
            _ = v21
        }
        for fi22 := 0; fi22 < 12; fi22++ {
            switch fi22 {
            case 439:
                v23 := GenS4{F0: int8(fi22), F1: "", F2: 707}
                _ = v23
            case 32767:
                var v24 uint16 = 492
                _ = v24
            case 32768:
                v1 -= v1
            default:
                switch "\n\t" {
                case "":
                    v25 := GenS4{F0: 76, F1: ("foo bar" + "foo bar"), F2: v2}
                    _ = v25
                default:
                    v19 = uint8(v1)
                }
            }
            v19 += uint8(v1)
            v19 += uint8(v2)
            _ = fi22
        }
        var v26 string = ("hello" + "hello")
        _ = fi20
        _ = v26
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch 42 {
    case 128:
        var v27 int8 = 81
        _ = v27
    case 9223372036854775807:
        if true {
            v28 := GenS3{F0: (1 < -12), F1: -9223372036854775808, F2: 2}
            fmt.Println(v2, v1)
            v29 := GenS2{F0: 2147483647, F1: 127, F2: 42}
            _ = v28
            _ = v29
        } else {
            v2 = v2
            v1 += v1
        }
    case -9223372036854775808:
        v4 = append(v4, 65535)
    }
    fmt.Println(v1, v2, v2)
    var v30 uint8 = 42
    v31 := GenS2{F0: 2147483647, F1: 73, F2: (v1 << ((1) & 63))}
    for fi32 := 0; fi32 < 14; fi32++ {
        v30 = (^(v30))
        v33 := GenS1{F0: 127, F1: 100}
        var v34 uint16 = 32767
        _ = fi32
        _ = v33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
