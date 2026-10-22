package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int64) string {
    var v5 int8 = -1
    var v6 uint64 = 1
    fmt.Println(v5, p0, p0)
    _ = p0
    _ = v5
    _ = v6
    return "foo bar"
}

func main() {
    var v1 uint8 = 207
    v2 := []string{("a" + ""), "hello", ("" + "a")}
    v3 := make(map[uint32]int)
    var v4 uint64 = 0
    fmt.Println(v4, v1)
    var v7 int8 = 42
    switch -781 {
    case 482:
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 int32 = 42
            for fi10 := 0; fi10 < 4; fi10++ {
                v1 -= 1
                v11 := GenS1{F0: 755, F1: uint16(v1), F2: v4}
                var v12 uint8 = 42
                _ = fi10
                _ = v11
                _ = v12
            }
            v3[499] = (fi8 / ((fi8) | 1))
            _ = fi8
            _ = v9
        }
    case 32767:
        v13 := GenS1{F0: 484, F1: 32767, F2: (v4 >> ((18446744073709551615) & 63))}
        _ = v13
    default:
        v14 := GenS1{F0: 0, F1: 790, F2: (v4 % (((^(v4))) | 1))}
        _ = v14
    }
    v15 := GenS2{F0: "\n\t", F1: 1}
    switch "a" {
    case "hello":
        for fi16 := 0; fi16 < 19; fi16++ {
            v4 -= 2
            v17 := GenS3{F0: uint64(v1)}
            v4 = (v4 % ((v4) | 1))
            _ = fi16
            _ = v17
        }
    }
    v1 = v1
    fmt.Println(v1, v4, v7)
    var v18 uint32 = 240
    switch -283 {
    case 32768:
        var v19 int64 = -818
        _ = v19
    default:
        for fi20 := 0; fi20 < 9; fi20++ {
            var v21 uint = 331
            if (-1 <= -9223372036854775808) {
                var v22 string = ("foo bar" + "\n\t")
                if ((v1 << (((v21 * 42)) & 63)) != 255) {
                    fmt.Println(v18, v4)
                    fmt.Println(v21, v4, fi20)
                    var v23 float32 = -2.25
                    _ = v23
                } else {
                    var v24 uint64 = (v4 ^ (^(v4)))
                    v25 := GenS2{F0: "", F1: 42}
                    var v26 uint = 2
                    _ = v24
                    _ = v25
                    _ = v26
                }
                v4 = v4
                _ = v22
            } else {
                fmt.Println(v1, v1, fi20)
                var v27 float64 = 0.1
                var v28 string = "\n\t"
                _ = v27
                _ = v28
            }
            _ = fi20
            _ = v21
        }
    }
    if ((!(false || (127 == -786))) || true) {
        v1 = 42
        for fi29 := 0; fi29 < 13; fi29++ {
            v30 := v3[(v18 - v18)]
            _ = fi29
            _ = v30
        }
        var v31 float32 = -2.25
        _ = v31
    } else {
        v4 = 42
        v2 = append(v2, "hello")
    }
    var v32 int8 = (v7 ^ v7)
    v18 = 158
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v7
}
