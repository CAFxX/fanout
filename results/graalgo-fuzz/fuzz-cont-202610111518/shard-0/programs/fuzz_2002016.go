package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 int64; F2 uint64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    v4 := make(map[string]string)
    var v5 int32 = 2147483647
    v6 := make(map[uint32]int)
    v7 := []int{}
    v7 = append(v7, 42)
    v5 += (-(v5))
    v8 := GenS1{F0: -487, F1: 2}
    v5 -= 0
    v9 := GenS1{F0: -136, F1: 2}
    if ((^(v5)) != 42) {
        switch 9223372036854775807 {
        case 1:
            var v10 string = "\n\t"
            _ = v10
        case 127:
            var v11 bool = (true && (v5 == v5))
            _ = v11
        case 128:
            v12 := GenS2{F0: 79}
            _ = v12
        default:
            fmt.Println(v5)
        }
    }
    v5 = (v5 + (-(v5)))
    switch -2147483648 {
    case -395:
        if len(v7) > 0 {
            v7[0] = 128
        }
    case 42:
        var v13 uint8 = 128
        _ = v13
    case -2147483648:
        for fi14 := 0; fi14 < 12; fi14++ {
            switch "a" {
            case "hello":
                var v15 int32 = (v5 ^ v5)
                _ = v15
            default:
                v16 := GenS2{F0: 127}
                _ = v16
            }
            v5 = v5
            switch "a" {
            case "foo bar":
                var v17 uint = 2
                _ = v17
            case "":
                v5 = (^(v5))
            case "a":
                switch "hello" {
                case "foo bar":
                    var v18 uint16 = 255
                    _ = v18
                default:
                    if false {
                        v7 = append(v7, fi14)
                        v19 := GenS3{F0: (("\n\t" + "foo bar") + ("foo bar" + "\n\t"))}
                        _ = v19
                    }
                }
            default:
                fmt.Println(fi14, v5)
            }
            _ = fi14
        }
    default:
        var v20 float32 = -2.25
        _ = v20
    }
    if false {
        v5 = (v5 & (v5 / (((^(v5))) | 1)))
    }
    var v21 uint64 = 0
    if (257 <= 255) {
        v21 += 0
    } else {
        for fi22 := 0; fi22 < 17; fi22++ {
            v21 -= v21
            _ = len(v4)
            v23 := GenS3{F0: ("" + "a")}
            _ = fi22
            _ = v23
        }
        fmt.Println(v21, v5)
        var v24 int = -9223372036854775808
        _ = v24
    }
    v25 := GenS1{F0: -32768, F1: 0}
    switch "" {
    case "":
        var v26 float32 = (0.1 * 0.0)
        _ = v26
    case "\n\t":
        v27 := GenS3{F0: "\n\t"}
        _ = v27
    case "hello":
        if (!false) {
            v5 = (v5 ^ 42)
        }
    default:
        v28 := GenS4{F0: v5, F1: 0, F2: v21}
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
