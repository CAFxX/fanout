package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint; F1 uint; F2 uint8 }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 uint8; F1 int64; F2 bool }
func genFn0(p0 int8) uint32 {
    switch -2147483648 {
    case 415:
        v5 := []int{}
        _ = v5
    case 9223372036854775807:
        p0 += (p0 * p0)
    case -803:
        for fi6 := 0; fi6 < 10; fi6++ {
            p0 -= p0
            v7 := []int{fi6}
            p0 = p0
            _ = fi6
            _ = v7
        }
    }
    switch "a" {
    case "a":
        var v8 uint16 = 65535
        _ = v8
    case "hello":
        p0 = (p0 % (((p0 * (p0 & (p0 | p0)))) | 1))
    default:
        switch 128 {
        case 32768:
            p0 += 1
        case 198:
            p0 += (p0 * p0)
        case 4294967295:
            v9 := GenS1{F0: false, F1: 1}
            _ = v9
        default:
            v10 := GenS1{F0: (((127 >= 9) && (155 != 2)) && true), F1: 958}
            _ = v10
        }
    }
    var v11 int8 = (-(p0))
    _ = p0
    _ = v11
    return 513
}

func main() {
    var v1 int64 = 1
    v2 := make(map[uint32]bool)
    v3 := []int{-9223372036854775808, 127}
    var v4 uint8 = 255
    if (-2.25 >= 1.5) {
        fmt.Println(v1, v4)
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 16; fi13++ {
            v14 := GenS2{F0: 98, F1: 4294967295, F2: 251}
            var v15 int64 = 42
            _ = fi13
            _ = v14
            _ = v15
        }
        v1 += v1
        var v16 int16 = 1
        _ = fi12
        _ = v16
    }
    v4 = v4
    v1 = (v1 % ((v1) | 1))
    if true {
        v2[316] = (100 != 2)
    }
    switch "\n\t" {
    case "foo bar":
        var v17 string = "a"
        _ = v17
    default:
        v1 += (v1 - 42)
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v4 = v4
        for fi19 := 0; fi19 < 13; fi19++ {
            v20 := GenS3{F0: 105, F1: 42}
            v21 := GenS3{F0: -119, F1: 91}
            if true {
                v22 := GenS4{F0: (v4 - v4), F1: v1, F2: true}
                var v23 int64 = v1
                v1 = (v1 + (v23 / ((v1) | 1)))
                var v24 uint8 = v4
                if false {
                    v25 := GenS1{F0: true, F1: -1}
                    var v26 int = fi19
                    var v27 int32 = 1
                    delete(v2, 670)
                    _ = v25
                    _ = v26
                    _ = v27
                }
                _ = v22
                _ = v23
                _ = v24
            }
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi18
    }
    var v28 float32 = (0.0 * 0.0)
    var v29 float64 = (((-2.25 + 3.14159) + -2.25) - -2.25)
    for fi30 := 0; fi30 < 17; fi30++ {
        v2[genFn0(0)] = (-9223372036854775808 != v1)
        _ = fi30
    }
    v28 += 3.14159
    var v31 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v4
}
