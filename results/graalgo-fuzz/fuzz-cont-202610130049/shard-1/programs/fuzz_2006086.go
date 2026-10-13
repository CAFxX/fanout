package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint16 }
func genFn0(p0 int16, p1 uint64) uint8 {
    fmt.Println(p0, p0)
    var v4 int = -406
    p0 += p0
    v5 := []int{}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 string, p1 uint, p2 string) uint64 {
    var v6 uint8 = 132
    var v7 string = p0
    v7 = (p0 + p2)
    v8 := []int{-1, 128, -9223372036854775808, 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[int]int)
    v3 := []int{448, 32768}
    if false {
        if len(v3) > 0 {
            v3[0] = -775
        }
    } else {
        _ = len(v2)
        v1 -= v1
        var v9 float64 = -2.25
        _ = v9
    }
    v10 := GenS1{F0: 19}
    var v11 int8 = 0
    var v12 int = -2147483648
    for fi13 := 0; fi13 < 5; fi13++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if true {
            v12 += (^(v12))
            v11 -= v11
            v11 = (-(v11))
        } else {
            switch v12 {
            case -2147483648:
                v1 -= (v1 * genFn1("", 739, ("" + "foo bar")))
            case -259:
                for fi14 := 0; fi14 < 19; fi14++ {
                    var v15 uint8 = 2
                    v1 = (v1 & 0)
                    v15 = uint8(v1)
                    _ = fi14
                    _ = v15
                }
            }
            if (false || (v11 < v11)) {
                v16 := GenS2{F0: 2147483648, F1: 4294967295, F2: 255}
                _ = v16
            }
        }
        _ = fi13
    }
    v11 = v11
    for fi17 := 0; fi17 < 15; fi17++ {
        if true {
            fmt.Println(v1, fi17, v11)
            var v18 bool = (float64(0.0) < 0.1)
            _ = v18
        }
        if false {
            if (false && true) {
                switch "\n\t" {
                case "a":
                    var v19 int8 = (v11 | v11)
                    _ = v19
                case "hello":
                    fmt.Println(fi17)
                case "foo bar":
                    v11 -= v11
                default:
                    _ = len(v3)
                }
                var v20 int8 = v11
                _ = v20
            }
            if ("a" >= "") {
                for fi21 := 0; fi21 < 10; fi21++ {
                    v3 = append(v3, (fi17 / (((fi17 * fi17)) | 1)))
                    fmt.Println(v12)
                    v22 := GenS2{F0: 2147483647, F1: 2147483648, F2: 40}
                    _ = fi21
                    _ = v22
                }
                v12 = (v12 & (v12 - fi17))
                if ((int64(v11) <= -660) || false) {
                    v1 += (v1 - v1)
                    v23 := GenS1{F0: 2}
                    _ = v23
                }
            }
            if (3.14159 > 0.0) {
                var v24 uint8 = 189
                v24 = v24
                v1 -= (v1 % (((v1 + 18446744073709551615)) | 1))
                _ = v24
            }
        }
        _ = fi17
    }
    v1 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
}
