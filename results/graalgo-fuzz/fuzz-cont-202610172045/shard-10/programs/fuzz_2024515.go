package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int64 }
type GenS2 struct { F0 int16; F1 int8; F2 int16 }
func genFn0() uint16 {
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 int = fi6
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 float32 = 0.0
        var v10 float32 = (v9 * v9)
        v9 -= (v9 * (v10 + v10))
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := []int{-775}
    _ = v11
    return 0
}
func genFn1(p0 uint16, p1 int16, p2 uint) uint8 {
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 float32 = 0.0
        for fi14 := 0; fi14 < 5; fi14++ {
            switch ("hello" + "\n\t") {
            case "foo bar":
                fmt.Println(fi12, fi12, fi14)
                fmt.Println(p0, v13, fi14)
                var v15 string = (("a" + "\n\t") + "\n\t")
                _ = v15
            case "hello":
                v16 := GenS1{F0: -1, F1: 36, F2: 1}
                _ = v16
            case "":
                v17 := []int{-512, (^(fi14)), fi12}
                _ = v17
            }
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    v18 := []int{9223372036854775807, -378}
    p1 = p1
    v19 := GenS1{F0: -128, F1: 42, F2: -238}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return 2
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[int]string)
    var v3 uint64 = 1
    v4 := make(map[uint32]int)
    var v5 uint = 42
    switch "\n\t" {
    case "":
        for fi20 := 0; fi20 < 7; fi20++ {
            v1 = 42
            var v21 int = 32768
            fmt.Println(v3, fi20, fi20)
            _ = fi20
            _ = v21
        }
    case "hello":
        if (42 == -1) {
            var v22 float64 = -2.25
            v3 += v3
            _ = v22
        }
    case "a":
        if false {
            delete(v2, 32768)
            v3 += (v3 >> (((v5 / (((v5 + v5)) | 1))) & 63))
        }
    default:
        v23 := []int{}
        _ = v23
    }
    if true {
        v2[128] = "foo bar"
    }
    v24 := []int{32767}
    switch -674 {
    case -1:
        v5 = v5
    case 32767:
        var v25 int32 = -1
        _ = v25
    case 4294967295:
        v26 := GenS2{F0: -32768, F1: -1, F2: 1}
        _ = v26
    }
    var v27 uint8 = genFn1(255, int16(v1), (v5 & 432))
    _ = len(v24)
    var v28 uint = v5
    fmt.Println(v1, v28, v1)
    delete(v2, -1)
    var v29 int32 = 2147483647
    _ = len(v24)
    fmt.Println(v3, v27, v29)
    v24 = append(v24, 42)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
