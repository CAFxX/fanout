package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 string; F2 int64 }
type GenS3 struct { F0 int; F1 bool; F2 int8 }

func main() {
    v1 := []uint32{571, 2147483647, 592}
    v2 := make(map[string]bool)
    var v3 int8 = -109
    v4 := make(map[int]bool)
    var v5 int = 32767
    var v6 uint = 473
    v6 = 716
    v6 += v6
    v7 := v4[(v5 & v5)]
    switch "\n\t" {
    case "":
        if v7 {
            for fi8 := 0; fi8 < 6; fi8++ {
                v9 := v4[0]
                var v10 int = -566
                _ = fi8
                _ = v10
                _ = v9
            }
            var v11 uint = (v6 >> ((v6) & 63))
            _ = v11
        } else {
            var v12 uint32 = 4294967295
            _ = v12
        }
    default:
        var v13 float32 = -2.25
        _ = v13
    }
    if ((false || false) || (!(v7 && (!(((v5 + -2147483648) > 600) && ((3.14159 + (float32(-2.25) / -2.25)) < 0.0)))))) {
        if true {
            var v14 float64 = 0.1
            v14 += ((1e10 * 3.14159) + 3.14159)
            _ = v14
        } else {
            var v15 uint8 = 127
            fmt.Println(v5, v3)
            if ((!v7) && v7) {
                switch "hello" {
                case "":
                    v16 := v4[(v5 >> (((v6 / ((v6) | 1))) & 63))]
                    _ = v16
                case "\n\t":
                    for fi17 := 0; fi17 < 10; fi17++ {
                        v18 := GenS1{F0: 42}
                        _ = fi17
                        _ = v18
                    }
                default:
                    v6 = v6
                }
                v15 = v15
                for fi19 := 0; fi19 < 19; fi19++ {
                    var v20 int = (v5 | (^(fi19)))
                    for fi21 := 0; fi21 < 20; fi21++ {
                        v7 = true
                        _ = fi21
                    }
                    _ = fi19
                    _ = v20
                }
            } else {
                switch "foo bar" {
                case "\n\t":
                    v22 := GenS2{F0: 38, F1: "\n\t", F2: 1}
                    _ = v22
                case "hello":
                    if v7 {
                        v2[("a" + ("" + ""))] = (42 == 65535)
                    }
                default:
                    switch v5 {
                    case -727:
                        fmt.Println(v5)
                    case -253:
                        var v23 uint8 = v15
                        _ = v23
                    case -1:
                        if v7 {
                            for fi24 := 0; fi24 < 8; fi24++ {
                                var v25 uint8 = v15
                                v25 = 0
                                _ = fi24
                                _ = v25
                            }
                        } else {
                            if (true && (true && true)) {
                                v3 = v3
                                switch "foo bar" {
                                case "\n\t":
                                    v15 = 128
                                }
                                v7 = (-9223372036854775808 != -1)
                            }
                            fmt.Println(v3, v5, v6)
                            v26 := GenS1{F0: 255}
                            _ = v26
                        }
                    }
                }
            }
            _ = v15
        }
    } else {
        fmt.Println(v6, v6, v6)
        v6 += v6
    }
    _ = len(v1)
    v27 := GenS3{F0: (-(v5)), F1: (!(v7 && ((-1 >= v5) && true))), F2: v3}
    var v28 float64 = (-2.25 / 1.5)
    var v29 uint = (^(v6))
    v28 -= float64(3.14159)
    for fi30 := 0; fi30 < 9; fi30++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi30
    }
    fmt.Println(v7)
    var v31 uint16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
