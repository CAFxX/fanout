package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint8; F2 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 uint16; F2 uint16 }
type GenS5 struct { F0 uint32; F1 string; F2 int16 }
func genFn0() int8 {
    v6 := []int{32768, -2147483648}
    var v7 uint8 = 228
    v7 = (v7 / (((^(v7))) | 1))
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint, p1 uint32, p2 bool) int32 {
    v8 := GenS1{F0: 581, F1: 106, F2: p1}
    p1 -= 42
    for fi9 := 0; fi9 < 6; fi9++ {
        p2 = true
        var v10 uint16 = 255
        v11 := GenS2{F0: 1, F1: 2, F2: ("hello" + "a")}
        _ = fi9
        _ = v10
        _ = v11
    }
    if p2 {
        for fi12 := 0; fi12 < 5; fi12++ {
            p2 = p2
            switch (fi12 + fi12) {
            case 4294967295:
                for fi13 := 0; fi13 < 13; fi13++ {
                    p2 = (2 != 18446744073709551615)
                    _ = fi13
                }
            case -115:
                fmt.Println(p2, fi12)
            case -379:
                v14 := []int{}
                _ = v14
            default:
                v15 := GenS3{F0: 0}
                _ = v15
            }
            if (0 >= 65535) {
                for fi16 := 0; fi16 < 15; fi16++ {
                    p2 = (-1 <= 42)
                    var v17 bool = true
                    _ = fi16
                    _ = v17
                }
                p1 += p1
            }
            _ = fi12
        }
        for fi18 := 0; fi18 < 9; fi18++ {
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 int8 = -88
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]bool)
    var v3 int16 = 452
    var v4 uint = 87
    v5 := make(map[uint32]int)
    switch ("a" + ("" + "")) {
    case "hello":
        v3 += v3
    default:
        v21 := []int{128}
        _ = v21
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        v3 += (v3 + v3)
        _ = fi22
    }
    var v23 int16 = (v3 ^ v3)
    v24 := GenS4{F0: 9223372036854775807, F1: 0, F2: 65535}
    v3 += 42
    var v25 float32 = float32(float64(float32(3.14159)))
    v4 -= v4
    v26 := GenS5{F0: 2, F1: "", F2: (v3 - (v23 ^ 1000))}
    if true {
        v23 = (^(v23))
        var v27 int8 = -33
        _ = v27
    } else {
        var v28 uint8 = 2
        for fi29 := 0; fi29 < 14; fi29++ {
            fmt.Println(v3, v4)
            fmt.Println(v4, v25)
            v30 := GenS1{F0: 18446744073709551615, F1: 49, F2: 4294967295}
            _ = fi29
            _ = v30
        }
        _ = v28
    }
    v23 -= (v3 >> (((v4 & 18446744073709551615)) & 63))
    var v31 bool = (v25 != float32(1e10))
    v32 := []int{206, 1}
    var v33 uint16 = 1
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
