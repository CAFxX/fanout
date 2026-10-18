package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint64 }
func genFn0() int64 {
    var v5 int8 = -43
    v6 := []int{-2147483648}
    v5 += (v5 | v5)
    if false {
        switch "foo bar" {
        case "hello":
            if false {
                v7 := GenS1{F0: true, F1: v5}
                _ = v7
            }
        default:
            switch 32768 {
            case 2147483647:
                var v8 bool = false
                _ = v8
            case -9223372036854775808:
                v5 += (^(v5))
            case 134:
                var v9 int64 = -9223372036854775808
                _ = v9
            default:
                var v10 int = 127
                _ = v10
            }
        }
        fmt.Println(v5)
        v5 = v5
    }
    switch -530 {
    case -2147483648:
        var v11 float64 = 0.1
        _ = v11
    case -849:
        v5 = v5
    default:
        var v12 float32 = 3.14159
        _ = v12
    }
    _ = v5
    _ = v6
    return -864
}
func genFn1(p0 int32) int8 {
    switch 4294967295 {
    case -478:
        v13 := []int{641, 32768}
        _ = v13
    }
    p0 = p0
    _ = p0
    return 0
}

func main() {
    var v1 int8 = -6
    v2 := make(map[string]bool)
    var v3 int = -333
    v4 := make(map[uint32]uint32)
    v1 = v1
    var v14 int8 = (v1 / (((^(v1))) | 1))
    v15 := []int{-400, 9223372036854775807, v3}
    var v16 int = v3
    var v17 uint8 = 128
    var v18 uint8 = v17
    var v19 uint = 18446744073709551615
    if false {
        for fi20 := 0; fi20 < 1; fi20++ {
            for fi21 := 0; fi21 < 19; fi21++ {
                var v22 bool = true
                _ = fi21
                _ = v22
            }
            v15 = append(v15, v3)
            _ = len(v4)
            _ = fi20
        }
        v18 = v17
        switch "" {
        case "a":
            fmt.Println(v16, v1, v14)
        case "\n\t":
            v23 := GenS2{F0: 33, F1: 4294967295}
            _ = v23
        case "hello":
            v24 := GenS2{F0: 334, F1: 622}
            _ = v24
        default:
            v1 -= 42
        }
    }
    v25 := GenS3{F0: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
