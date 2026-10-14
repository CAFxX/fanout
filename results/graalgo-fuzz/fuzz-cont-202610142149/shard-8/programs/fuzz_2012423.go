package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint64 }
func genFn0(p0 int) uint8 {
    fmt.Println(p0)
    var v5 int8 = 100
    var v6 float32 = 3.14159
    v7 := []int{-657, (p0 * p0)}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 36
}

func main() {
    var v1 bool = (767 == 1)
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    v4 := make(map[int]string)
    var v8 bool = (2147483647 < 4294967295)
    v4[0] = ("foo bar" + "hello")
    var v9 uint64 = 129
    if v1 {
        v10 := GenS1{F0: genFn0(9223372036854775807), F1: genFn0(-510), F2: 980}
        var v11 bool = (false || (9223372036854775807 > 0))
        _ = v10
        _ = v11
    } else {
        switch "" {
        case "a":
            fmt.Println(v9, v1, v9)
        }
        v9 += v9
        v1 = v1
    }
    v9 = (v9 | (v9 % (((v9 << ((42) & 63))) | 1)))
    v9 += v9
    v9 -= (^(v9))
    fmt.Println(v1)
    v9 = (v9 << ((626) & 63))
    v1 = (!(!(!v8)))
    v8 = (9223372036854775807 != 0)
    if v8 {
        var v12 uint64 = (v9 * 18446744073709551615)
        if (2 <= 2) {
            if v8 {
                var v13 int64 = 932
                _ = v13
            } else {
                if v8 {
                    var v14 string = ("" + "\n\t")
                    v8 = true
                    fmt.Println(v1, v9)
                    _ = v14
                }
            }
            for fi15 := 0; fi15 < 17; fi15++ {
                v12 += (v12 * (v9 << ((42) & 63)))
                switch fi15 {
                case 128:
                    v8 = (-1 >= 2147483647)
                case 353:
                    switch fi15 {
                    case 1:
                        if (!true) {
                            if (42 < 0) {
                                v12 -= (v12 - v12)
                                var v16 bool = (v1 && (false || v8))
                                v8 = (((v9 ^ (v12 << ((839) & 63))) > 0) || (!v16))
                                _ = v16
                            } else {
                                for fi17 := 0; fi17 < 3; fi17++ {
                                    var v18 float64 = (1e10 - 3.14159)
                                    v12 += v12
                                    _ = fi17
                                    _ = v18
                                }
                                switch "a" {
                                case "":
                                    var v19 uint64 = (v12 * v9)
                                    _ = v19
                                default:
                                    var v20 uint64 = v9
                                    _ = v20
                                }
                                v8 = (!v1)
                            }
                            fmt.Println(v12, v9, v1)
                        }
                    case -779:
                        var v21 uint64 = 0
                        _ = v21
                    default:
                        switch "hello" {
                        case "":
                            v22 := []int{(fi15 - (fi15 + (fi15 + -909)))}
                            _ = v22
                        case "hello":
                            for fi23 := 0; fi23 < 9; fi23++ {
                                var v24 uint = 163
                                _ = fi23
                                _ = v24
                            }
                        }
                    }
                default:
                    _ = len(v2)
                }
                fmt.Println(v12, v9)
                _ = fi15
            }
            v12 = (v9 >> ((65535) & 63))
        } else {
            if (1e10 != 1e10) {
                var v25 int16 = int16(v12)
                var v26 uint = 0
                v25 = (v25 >> ((v26) & 63))
                _ = v25
                _ = v26
            } else {
                if (!false) {
                    v1 = (!v1)
                } else {
                    v27 := []int{42, 32767, 2147483647}
                    v12 = v12
                    _ = v27
                }
                v1 = true
            }
        }
        _ = v12
    } else {
        v28 := v4[720]
        _ = v28
    }
    fmt.Println(v8)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
