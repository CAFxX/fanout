package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS2 struct { F0 int8; F1 int64; F2 int64 }
type GenS3 struct { F0 uint64; F1 int; F2 string }
type GenS4 struct { F0 uint8; F1 bool; F2 int32 }
type GenS5 struct { F0 int64; F1 int64 }
func genFn0(p0 string, p1 uint8, p2 bool) uint8 {
    v5 := []int{}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p1
}

func main() {
    v1 := []string{(("hello" + "foo bar") + "\n\t"), ""}
    var v2 int64 = 42
    var v3 int8 = -128
    var v4 float64 = 0.0
    v3 = 0
    v6 := GenS1{F0: 32767, F1: 2147483648, F2: v2}
    if false {
        v7 := GenS2{F0: 42, F1: v2, F2: (^(v2))}
        fmt.Println(v3, v4)
        v1 = append(v1, "foo bar")
        _ = v7
    }
    var v8 int32 = 182
    if (89 == 0) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        switch "foo bar" {
        case "":
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
        case "hello":
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
        case "a":
            if ((v8 == v8) && ((-9223372036854775808 <= (v2 - v2)) && false)) {
                switch "hello" {
                case "a":
                    if false {
                        v4 = 0.1
                        v8 = 463
                    } else {
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                        if len(v1) > 0 {
                            v1[0] = "\n\t"
                        }
                    }
                case "":
                    v9 := GenS2{F0: (v3 >> ((255) & 63)), F1: (v2 % ((v2) | 1)), F2: (v2 ^ -1)}
                    _ = v9
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                if ((v3 * 42) == v3) {
                    fmt.Println(v8, v3, v3)
                    var v10 int32 = (v8 % (((v8 & int32(v3))) | 1))
                    _ = v10
                } else {
                    if true {
                        switch "\n\t" {
                        case "foo bar":
                            if ((v3 ^ (v3 | (v3 % (((v3 - -1)) | 1)))) <= v3) {
                                v11 := GenS3{F0: 18446744073709551615, F1: 9223372036854775807, F2: "foo bar"}
                                _ = v11
                            }
                        case "hello":
                            if true {
                                if false {
                                    var v12 int16 = -999
                                    for fi13 := 0; fi13 < 2; fi13++ {
                                        v4 += v4
                                        var v14 uint64 = 2
                                        fmt.Println(v12)
                                        _ = fi13
                                        _ = v14
                                    }
                                    _ = v12
                                }
                                var v15 int = -9223372036854775808
                                v16 := GenS4{F0: 255, F1: false, F2: (v8 / (((v8 | (v8 & v8))) | 1))}
                                _ = v15
                                _ = v16
                            } else {
                                var v17 int = 1
                                var v18 uint = 286
                                _ = v17
                                _ = v18
                            }
                        }
                    } else {
                        var v19 string = ("hello" + (("" + ("" + "a")) + "hello"))
                        v8 -= (v8 >> ((423) & 63))
                        _ = v19
                    }
                }
                v20 := GenS2{F0: (v3 + -117), F1: (v2 + 42), F2: -9223372036854775808}
                _ = v20
            }
        }
        v21 := GenS5{F0: v2, F1: (v2 % ((v2) | 1))}
        _ = v21
    } else {
        var v22 int16 = 538
        fmt.Println(v2, v8, v2)
        _ = v22
    }
    v3 += (v3 ^ v3)
    v1 = append(v1, ("a" + "a"))
    if ((v3 & v3) < v3) {
        v1 = append(v1, "\n\t")
        for fi23 := 0; fi23 < 15; fi23++ {
            v8 = (v8 >> ((18446744073709551615) & 63))
            _ = fi23
        }
        fmt.Println(v3, v8)
    }
    var v24 string = ("\n\t" + "a")
    var v25 uint16 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v6
    _ = v8
}
