package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 bool; F1 uint8; F2 int }
type GenS3 struct { F0 uint64; F1 bool }
func genFn0(p0 uint32, p1 uint) bool {
    v7 := GenS1{F0: p1, F1: "a"}
    if false {
        for fi8 := 0; fi8 < 1; fi8++ {
            for fi9 := 0; fi9 < 1; fi9++ {
                switch fi8 {
                case 2147483647:
                    fmt.Println(p1, p0, p1)
                default:
                    p1 = 18446744073709551615
                }
                _ = fi9
            }
            _ = fi8
        }
        p0 = p0
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 int16 = 0
            p1 -= (p1 + p1)
            fmt.Println(p1, p0, p1)
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return true
}
func genFn1(p0 string) uint32 {
    var v12 string = p0
    var v13 int32 = 42
    v14 := []int{-343, 47, 733}
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return 65535
}

func main() {
    var v1 int = 128
    v2 := make(map[uint32]int)
    var v3 int8 = -128
    var v4 uint = 4294967295
    var v5 uint64 = 2
    v6 := []uint32{1, 0, 62, 99}
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 int16 = -144
        v17 := GenS1{F0: v4, F1: ("\n\t" + ("\n\t" + "\n\t"))}
        var v18 uint64 = (v5 / ((217) | 1))
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        switch ("\n\t" + "foo bar") {
        case "hello":
            v5 += v5
        case "a":
            v3 = (v3 % ((v3) | 1))
        case "foo bar":
            for fi20 := 0; fi20 < 13; fi20++ {
                v21 := GenS2{F0: (!true), F1: 2, F2: (fi20 / ((fi20) | 1))}
                v3 = (v3 << (((^(v4))) & 63))
                fmt.Println(v1, v4)
                _ = fi20
                _ = v21
            }
        }
        fmt.Println(v4)
        _ = fi19
    }
    if (true || (409 == v4)) {
        var v22 int32 = -2147483648
        fmt.Println(v1, v4)
        _ = v22
    }
    v1 -= (v1 ^ -9223372036854775808)
    fmt.Println(v3, v1)
    switch v1 {
    case 0:
        v23 := GenS3{F0: (v5 / (((v5 / ((v5) | 1))) | 1)), F1: (-74 == v1)}
        _ = v23
    case 128:
        v24 := GenS2{F0: (65535 == 410), F1: 154, F2: v1}
        _ = v24
    }
    if (87 < -822) {
        v2[0] = v1
    }
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
