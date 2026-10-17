package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int; F1 int32; F2 uint64 }
func genFn0(p0 uint64, p1 int8) uint32 {
    v5 := []int{-2147483648, -438}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = p0
    _ = p1
    _ = v5
    return 0
}
func genFn1(p0 int16, p1 uint8) int {
    switch -2147483648 {
    case 128:
        var v6 float32 = (1e10 - (0.1 / 3.14159))
        _ = v6
    case -9223372036854775808:
        var v7 uint = uint(p1)
        _ = v7
    default:
        p0 += (p0 % (((p0 & p0)) | 1))
    }
    p1 -= (p1 + (p1 + (p1 * 42)))
    _ = p0
    _ = p1
    return 9223372036854775807
}

func main() {
    var v1 uint32 = 604
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    v4 := []int{}
    v4 = append(v4, genFn1(-32768, 1))
    if false {
        v8 := GenS1{F0: 320, F1: 829, F2: 1}
        v2[172] = "hello"
        v1 -= genFn0(0, -95)
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        v1 = v1
        _ = fi9
    }
    v1 -= genFn0(0, 0)
    v10 := GenS2{F0: 32767, F1: 994, F2: 0}
    var v11 int16 = 42
    switch ("a" + ((("foo bar" + "") + ((("foo bar" + ("\n\t" + "")) + "foo bar") + "")) + ("a" + ("foo bar" + ("" + (("" + "a") + "foo bar")))))) {
    case "\n\t":
        if true {
            v1 -= v1
            for fi12 := 0; fi12 < 5; fi12++ {
                v11 += 120
                _ = fi12
            }
        } else {
            if (-1 <= -1) {
                for fi13 := 0; fi13 < 11; fi13++ {
                    v1 += v1
                    _ = fi13
                }
                var v14 bool = false
                _ = v14
            }
            var v15 bool = true
            _ = v15
        }
    case "hello":
        for fi16 := 0; fi16 < 6; fi16++ {
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 float32 = 0.0
                var v19 uint16 = 213
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = fi16
        }
    }
    v1 = (^(v1))
    v11 = -32768
    if true {
        delete(v3, 742)
        var v20 bool = true
        if false {
            if v20 {
                var v21 uint16 = 2
                _ = v21
            }
        }
        _ = v20
    }
    v1 = v1
    v22 := GenS1{F0: 132, F1: 642, F2: 9}
    for fi23 := 0; fi23 < 11; fi23++ {
        fmt.Println(fi23, v1)
        v11 -= (v11 % (((-(v11))) | 1))
        switch (^(fi23)) {
        case 9223372036854775807:
            switch "foo bar" {
            case "a":
                if (genFn0(472, -5) > v1) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    fmt.Println(v1, v11)
                    v1 += v1
                    v11 += 0
                }
            case "foo bar":
                v4 = append(v4, -2147483648)
            }
        case 2147483647:
            v11 = (v11 >> ((2) & 63))
        default:
            var v24 float64 = 1.5
            _ = v24
        }
        _ = fi23
    }
    if len(v4) > 0 {
        v4[0] = genFn1(v11, 255)
    }
    v1 -= (v1 ^ 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
