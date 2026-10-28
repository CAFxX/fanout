package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8 }
func genFn0(p0 uint64, p1 uint16) string {
    v5 := GenS1{F0: 255, F1: int32(p0), F2: 2147483647}
    switch 42 {
    case 127:
        if (false || (0 <= 1)) {
            for fi6 := 0; fi6 < 17; fi6++ {
                if (!(100 >= 1)) {
                    var v7 float32 = (1e10 - 3.14159)
                    _ = v7
                }
                _ = fi6
            }
            for fi8 := 0; fi8 < 13; fi8++ {
                switch "a" {
                case "\n\t":
                    v9 := GenS1{F0: 128, F1: 815, F2: 437}
                    _ = v9
                case "":
                    if true {
                        v10 := []int{fi8, fi8, (-(fi8))}
                        _ = v10
                    }
                default:
                    p0 = p0
                }
                p0 = p0
                _ = fi8
            }
            p0 = (p0 & p0)
        } else {
            var v11 int = 128
            _ = v11
        }
    default:
        v12 := []int{314, 2147483647, 1}
        _ = v12
    }
    v13 := []int{-1, 9223372036854775807, 32767, 128}
    p1 = p1
    _ = p0
    _ = p1
    _ = v13
    _ = v5
    return "foo bar"
}
func genFn1(p0 bool) uint {
    v14 := []int{42, -2147483648, 0, -124}
    v15 := GenS1{F0: 42, F1: -867, F2: 626}
    _ = p0
    _ = v14
    _ = v15
    return 65535
}

func main() {
    var v1 int8 = 42
    v2 := []string{}
    var v3 float32 = 1e10
    var v4 float64 = 0.0
    _ = len(v2)
    for fi16 := 0; fi16 < 4; fi16++ {
        v17 := GenS2{F0: 0}
        _ = fi16
        _ = v17
    }
    switch "hello" {
    case "a":
        fmt.Println(v1, v3)
    case "hello":
        v18 := GenS3{F0: (false && true)}
        _ = v18
    default:
        v19 := GenS3{F0: (false && ((32767 > 65535) && true))}
        _ = v19
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 uint64 = 4294967295
        _ = fi20
        _ = v21
    }
    var v22 uint8 = 128
    fmt.Println(v3)
    var v23 int64 = 0
    for fi24 := 0; fi24 < 14; fi24++ {
        v3 = v3
        var v25 uint = 1
        var v26 uint16 = 65535
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 float32 = v3
    var v28 int = -1
    v22 = (v22 << ((4294967295) & 63))
    fmt.Println(v22, v4)
    if true {
        for fi29 := 0; fi29 < 7; fi29++ {
            switch (fi29 * 32768) {
            case 180:
                if (true || true) {
                    var v30 uint64 = 0
                    switch ("a" + "foo bar") {
                    case "foo bar":
                        var v31 uint = 42
                        _ = v31
                    default:
                        var v32 uint = genFn1(false)
                        _ = v32
                    }
                    fmt.Println(v3)
                    _ = v30
                }
            case 977:
                for fi33 := 0; fi33 < 16; fi33++ {
                    v1 = 35
                    v34 := GenS4{F0: v22}
                    v4 = v4
                    _ = fi33
                    _ = v34
                }
            case 128:
                var v35 uint16 = 782
                _ = v35
            }
            _ = fi29
        }
        v23 += (v23 >> ((genFn1(((!(v4 > 0.1)) || true))) & 63))
    } else {
        var v36 float64 = v4
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
