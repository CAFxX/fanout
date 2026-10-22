package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint64 }
type GenS2 struct { F0 bool; F1 int; F2 uint8 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint8, p1 int32, p2 int64) uint8 {
    v6 := []int{4294967295}
    if false {
        _ = len(v6)
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        if false {
            fmt.Println(p0)
        }
        _ = fi7
    }
    _ = len(v6)
    var v8 uint = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v8
    return (p0 >> ((v8) & 63))
}
func genFn1(p0 uint64, p1 int32, p2 uint8) uint32 {
    v9 := []int{-1, -429, -2147483648}
    p1 = -2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 2147483647
}

func main() {
    var v1 uint16 = 310
    var v2 int16 = -146
    v3 := make(map[string]uint32)
    v4 := []uint32{4294967295}
    v5 := make(map[int]bool)
    fmt.Println(v1)
    for fi10 := 0; fi10 < 16; fi10++ {
        v1 = (v1 + (^(v1)))
        v4 = append(v4, 0)
        _ = fi10
    }
    v4 = append(v4, 42)
    var v11 int = 800
    switch "a" {
    case "a":
        if (true || ((v2 == v2) && (!(false && ((-128 <= -128) && (!((0.1 + 1.5) != 0.0))))))) {
            v12 := GenS1{F0: 940, F1: 771, F2: 2}
            for fi13 := 0; fi13 < 1; fi13++ {
                var v14 uint = 65535
                _ = fi13
                _ = v14
            }
            var v15 string = "\n\t"
            _ = v12
            _ = v15
        } else {
            v2 = (^(v2))
            v4 = append(v4, 730)
        }
    default:
        v1 = 265
    }
    v11 = v11
    switch v11 {
    case -2147483648:
        v16 := GenS2{F0: ((true && false) && true), F1: v11, F2: 128}
        _ = v16
    case -391:
        _ = len(v3)
    case 2147483647:
        var v17 uint64 = 915
        _ = v17
    }
    v11 = (v11 ^ (v11 | (v11 % (((v11 | (v11 + 861))) | 1))))
    v18 := GenS3{F0: (v2 % ((v2) | 1))}
    switch v11 {
    case -9223372036854775808:
        v2 = (v2 * v2)
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if (817 > 26) {
        v11 = (v11 | v11)
        switch "foo bar" {
        case "hello":
            v19 := GenS2{F0: false, F1: v11, F2: 0}
            _ = v19
        }
    }
    fmt.Println(v11, v11)
    v2 += -1
    if (false && false) {
        v2 += (v2 ^ 0)
        v3[("\n\t" + "a")] = genFn1(355, 0, 221)
        if (1 == 1) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    } else {
        switch v11 {
        case -2147483648:
            switch "hello" {
            case "":
                v11 -= -984
            }
        case 32768:
            if true {
                v11 = v11
                if (false && false) {
                    v2 += v2
                }
            }
        case 127:
            v20 := GenS3{F0: (v2 >> ((961) & 63))}
            _ = v20
        default:
            v2 -= v2
        }
    }
    var v21 string = "a"
    fmt.Println(v11, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
