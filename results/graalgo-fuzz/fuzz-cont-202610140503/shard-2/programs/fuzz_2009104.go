package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32; F1 uint32 }
func genFn0() uint16 {
    if true {
        var v5 int16 = -32768
        v6 := []int{4294967295, 42, 4294967295, 4294967295}
        v6 = append(v6, 127)
        _ = v5
        _ = v6
    }
    var v7 string = ""
    v8 := []int{-2147483648, -1, 2147483647, 128}
    switch -1 {
    case -71:
        v9 := GenS1{F0: 778}
        _ = v9
    case -9223372036854775808:
        var v10 uint64 = 0
        _ = v10
    case 42:
        fmt.Println(v7)
    default:
        if (42 >= 771) {
            for fi11 := 0; fi11 < 11; fi11++ {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                v12 := GenS2{F0: 128, F1: 381}
                _ = fi11
                _ = v12
            }
        }
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 int64 = -137
        v15 := GenS1{F0: 65535}
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint32, p1 uint) uint {
    v16 := GenS1{F0: p0}
    v17 := []int{-873, 32768, 42}
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return p1
}

func main() {
    v1 := []string{}
    var v2 int8 = 0
    var v3 int = 0
    v4 := make(map[string]int)
    v3 -= (v3 / ((v3) | 1))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        var v19 int64 = 0
        v20 := GenS3{F0: 0, F1: 223}
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println(v2, v2)
    for fi21 := 0; fi21 < 18; fi21++ {
        v3 = (v3 << ((18446744073709551615) & 63))
        for fi22 := 0; fi22 < 8; fi22++ {
            var v23 float64 = 3.14159
            var v24 int64 = 261
            fmt.Println(v23, fi22)
            var v25 float64 = (0.1 + v23)
            var v26 float64 = (v25 * 1.5)
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
            _ = v26
        }
        _ = fi21
    }
    v27 := GenS2{F0: 127, F1: -492}
    v1 = append(v1, ("a" + ("a" + ("foo bar" + "a"))))
    var v28 uint = 356
    var v29 uint16 = 2
    v28 = 255
    v30 := GenS4{F0: -1}
    v31 := GenS5{F0: 129, F1: 2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
