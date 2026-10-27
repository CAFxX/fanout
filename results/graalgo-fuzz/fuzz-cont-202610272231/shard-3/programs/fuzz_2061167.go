package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 bool; F2 uint64 }
type GenS4 struct { F0 int8; F1 bool; F2 int16 }
type GenS5 struct { F0 int8; F1 int; F2 uint64 }
type GenS6 struct { F0 string; F1 int; F2 bool }
type GenS7 struct { F0 uint8 }
func genFn0() bool {
    var v5 int64 = 42
    for fi6 := 0; fi6 < 14; fi6++ {
        if (0 == 18446744073709551615) {
            v5 += (v5 ^ v5)
            if false {
                var v7 uint64 = 2
                v8 := []int{fi6, 128, (fi6 >> ((18446744073709551615) & 63)), (fi6 << ((1) & 63))}
                _ = v7
                _ = v8
            } else {
                v9 := GenS1{F0: 716, F1: -929, F2: -374}
                _ = v9
            }
            var v10 float64 = (-2.25 * -2.25)
            _ = v10
        }
        var v11 int32 = 782
        v12 := GenS2{F0: v5}
        _ = fi6
        _ = v11
        _ = v12
    }
    _ = v5
    return false
}
func genFn1(p0 int8, p1 int64, p2 uint64) uint {
    for fi13 := 0; fi13 < 1; fi13++ {
        v14 := GenS3{F0: 42, F1: ("hello" > "a"), F2: (^(p2))}
        _ = fi13
        _ = v14
    }
    switch ("" + ((("a" + "a") + "a") + "\n\t")) {
    case "":
        v15 := GenS4{F0: p0, F1: (!(9223372036854775807 > 9223372036854775807)), F2: -657}
        _ = v15
    case "\n\t":
        v16 := GenS3{F0: 832, F1: false, F2: (p2 << ((787) & 63))}
        _ = v16
    default:
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 int8 = p0
            _ = fi17
            _ = v18
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return uint(p2)
}

func main() {
    var v1 uint8 = 78
    v2 := make(map[string]bool)
    v3 := []string{}
    v4 := make(map[string]int)
    v19 := GenS5{F0: 96, F1: 0, F2: 18446744073709551615}
    for fi20 := 0; fi20 < 20; fi20++ {
        if (65535 == 42) {
            switch ("foo bar" + ("\n\t" + "hello")) {
            case "hello":
                if (42 < -1) {
                    v1 = (v1 - v1)
                    var v21 uint16 = 32767
                    _ = v21
                }
            default:
                v22 := GenS6{F0: "hello", F1: (fi20 & (fi20 >> ((42) & 63))), F2: false}
                _ = v22
            }
            if len(v3) > 0 {
                v3[0] = ("\n\t" + "\n\t")
            }
            var v23 uint8 = (v1 & (v1 * v1))
            _ = v23
        } else {
            var v24 uint16 = 2
            fmt.Println(v1, v1)
            v1 += v1
            _ = v24
        }
        _ = fi20
    }
    fmt.Println(v1)
    if ("" >= "hello") {
        if (false || false) {
            v1 = v1
            v25 := GenS7{F0: v1}
            _ = v25
        }
    } else {
        var v26 float32 = float32(1e10)
        var v27 int16 = -1
        var v28 int64 = 0
        _ = v26
        _ = v27
        _ = v28
    }
    fmt.Println(v1)
    var v29 int64 = 42
    var v30 bool = (3.14159 >= 1e10)
    var v31 uint8 = (v1 >> ((667) & 63))
    var v32 uint = 1
    v30 = (v29 <= v29)
    v31 -= (v1 << (((v32 ^ v32)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
