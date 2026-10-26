package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 bool; F1 int64; F2 string }
type GenS3 struct { F0 int64; F1 uint16; F2 int32 }
type GenS4 struct { F0 bool; F1 uint16 }

func main() {
    var v1 float32 = (-2.25 - (-2.25 / ((1e10 / 1.5) - float32(1.5))))
    var v2 float32 = v1
    var v3 int = 32768
    v4 := GenS1{F0: 5, F1: (v3 + v3)}
    v3 += v3
    v5 := []int{v3}
    var v6 uint16 = 1
    var v7 float64 = 1.5
    if false {
        if (((true || (v6 == v6)) && false) && true) {
            _ = len(v5)
        }
        v3 = (v3 / (((v3 >> ((2) & 63))) | 1))
        v2 = (v2 * v1)
    }
    var v8 int64 = -9223372036854775808
    v2 += (((0.0 + v2) - 3.14159) - v2)
    v6 -= (v6 / (((v6 - (v6 >> ((42) & 63)))) | 1))
    v5 = append(v5, v3)
    fmt.Println(v1)
    switch "foo bar" {
    case "foo bar":
        for fi9 := 0; fi9 < 6; fi9++ {
            if (42 <= 1) {
                switch -856 {
                case -63:
                    v10 := GenS2{F0: false, F1: (v8 | (v8 - v8)), F2: ("\n\t" + "")}
                    _ = v10
                case 127:
                    v11 := GenS3{F0: (v8 / (((^(v8))) | 1)), F1: (v6 * (v6 + v6)), F2: -576}
                    _ = v11
                case 2147483647:
                    if (0 > 4294967295) {
                        fmt.Println(v2)
                    }
                }
                var v12 string = "\n\t"
                _ = v12
            }
            fmt.Println(v6)
            if len(v5) > 0 {
                v5[0] = v3
            }
            _ = fi9
        }
    default:
        for fi13 := 0; fi13 < 4; fi13++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if (2147483647 <= 0) {
                v8 += (v8 * v8)
                if true {
                    if (false && (65535 >= 2)) {
                        fmt.Println(v8)
                        fmt.Println(v2)
                        if (false || false) {
                            if true {
                                v14 := GenS4{F0: (!true), F1: 65535}
                                v6 = 299
                                v5 = append(v5, 42)
                                _ = v14
                            } else {
                                if false {
                                    var v15 uint64 = 18446744073709551615
                                    fmt.Println(v7, fi13)
                                    _ = v15
                                }
                                for fi16 := 0; fi16 < 20; fi16++ {
                                    fmt.Println(fi16, v8, v1)
                                    v7 = v7
                                    fmt.Println(fi16, fi13, v8)
                                    _ = fi16
                                }
                            }
                            v3 = 2147483647
                            v1 = float32(float64(v2))
                        }
                    }
                    v3 = 127
                    var v17 uint32 = 0
                    _ = v17
                } else {
                    v1 += v2
                    if len(v5) > 0 {
                        v5[0] = 2147483647
                    }
                }
                fmt.Println(v1)
            }
            var v18 int16 = 1000
            _ = fi13
            _ = v18
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
