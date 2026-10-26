package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 string }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int }
func genFn0(p0 uint, p1 int8, p2 int64) uint {
    fmt.Println(p2)
    p1 = (p1 & (p1 >> (((^(p0))) & 63)))
    v6 := GenS1{F0: 1, F1: 4294967295, F2: "hello"}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 610
}
func genFn1(p0 uint32) uint8 {
    v7 := GenS1{F0: p0, F1: 719, F2: ("hello" + "foo bar")}
    switch ("" + "a") {
    case "\n\t":
        v8 := []int{}
        _ = v8
    case "foo bar":
        if (0.1 != float64(1e10)) {
            var v9 uint32 = p0
            switch 0 {
            case 32767:
                v10 := []int{0, 42, 606, 32767}
                _ = v10
            case 42:
                p0 = v9
            default:
                if (uint16(v9) <= 958) {
                    v11 := GenS1{F0: (v9 / (((p0 << ((0) & 63))) | 1)), F1: 1, F2: "foo bar"}
                    _ = v11
                } else {
                    if (!(0 <= 715)) {
                        p0 = (v9 >> ((875) & 63))
                        p0 += (v9 % (((p0 | (v9 * v9))) | 1))
                        var v12 string = ("foo bar" + "\n\t")
                        _ = v12
                    }
                    var v13 int16 = 1000
                    var v14 string = "\n\t"
                    _ = v13
                    _ = v14
                }
            }
            _ = v9
        }
    case "hello":
        if true {
            v15 := []int{}
            p0 -= (^(p0))
            _ = v15
        }
    default:
        var v16 int64 = 109
        _ = v16
    }
    _ = p0
    _ = v7
    return 42
}

func main() {
    var v1 float32 = 0.0
    var v2 uint32 = 220
    var v3 int16 = 1
    v4 := make(map[int]uint32)
    var v5 bool = (!(true && true))
    v5 = ("\n\t" != "hello")
    v2 -= 2147483647
    fmt.Println(v5)
    if v5 {
        if (!v5) {
            var v17 bool = v5
            fmt.Println(v1, v2)
            _ = v17
        } else {
            for fi18 := 0; fi18 < 16; fi18++ {
                var v19 uint64 = 42
                fmt.Println(v2, v5, v5)
                _ = fi18
                _ = v19
            }
        }
        fmt.Println(v2, v3, v1)
    }
    _ = len(v4)
    if true {
        if (866 <= 494) {
            switch "foo bar" {
            case "":
                var v20 int16 = (v3 | (v3 + v3))
                _ = v20
            case "\n\t":
                v5 = true
            case "a":
                v21 := GenS2{F0: uint8(v2), F1: (("" + "a") + "foo bar")}
                _ = v21
            }
            v5 = (!v5)
        }
    }
    fmt.Println(v1, v5, v5)
    for fi22 := 0; fi22 < 1; fi22++ {
        v23 := GenS3{F0: fi22}
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
