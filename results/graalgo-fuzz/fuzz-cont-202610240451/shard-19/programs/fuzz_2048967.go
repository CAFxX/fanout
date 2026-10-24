package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS7 struct { F0 uint; F1 bool }
type GenS8 struct { F0 uint16; F1 int8; F2 uint16 }
func genFn0(p0 uint64, p1 bool, p2 int32) int8 {
    for fi7 := 0; fi7 < 6; fi7++ {
        if (247 == 0) {
            p1 = true
        }
        _ = fi7
    }
    v8 := GenS1{F0: (p2 << ((460) & 63)), F1: 2147483648}
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := []int{}
        switch "hello" {
        case "foo bar":
            p2 += (p2 - (p2 / (((-(p2))) | 1)))
        default:
            for fi11 := 0; fi11 < 20; fi11++ {
                _ = len(v10)
                v12 := GenS2{F0: 541, F1: 2}
                _ = fi11
                _ = v12
            }
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 98
}
func genFn1(p0 uint, p1 uint, p2 int8) int16 {
    switch "hello" {
    case "":
        var v13 int16 = 1000
        _ = v13
    case "a":
        for fi14 := 0; fi14 < 5; fi14++ {
            if true {
                v15 := GenS3{F0: 1}
                p1 = (p1 & p0)
                v16 := []int{fi14, (-(fi14)), fi14, (fi14 / ((fi14) | 1))}
                _ = v15
                _ = v16
            } else {
                p0 = (p0 | p1)
            }
            p2 -= 127
            var v17 int32 = -835
            _ = fi14
            _ = v17
        }
    }
    var v18 uint = (p1 - p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return 77
}

func main() {
    v1 := []uint32{1, 2147483648, 2147483648}
    var v2 int = 2147483647
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]bool)
    var v5 uint8 = 101
    var v6 float32 = 3.14159
    v19 := GenS4{F0: (663 != uint16(v5))}
    var v20 int = (v2 + (v2 << ((135) & 63)))
    fmt.Println(v20, v6, v20)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        v6 -= (1e10 - 3.14159)
    } else {
        v2 = (^(v20))
    }
    if (false || (42 <= 128)) {
        v21 := GenS5{F0: 0}
        _ = v21
    }
    var v22 string = "\n\t"
    fmt.Println(v20, v5)
    v23 := GenS6{F0: 4294967295, F1: ((!(4294967295 > 1)) && (0 <= 1)), F2: 2147483647}
    v5 = v5
    if (0.1 != 0.0) {
        v20 += (v20 >> ((uint(v5)) & 63))
        _ = len(v1)
        v24 := GenS7{F0: 255, F1: false}
        _ = v24
    } else {
        var v25 uint32 = 2147483647
        fmt.Println(v2, v6)
        if true {
            fmt.Println(v5)
            var v26 uint64 = 2
            _ = v26
        } else {
            var v27 float32 = float32(-2.25)
            _ = v27
        }
        _ = v25
    }
    v6 = v6
    v28 := GenS8{F0: 292, F1: 127, F2: 769}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
