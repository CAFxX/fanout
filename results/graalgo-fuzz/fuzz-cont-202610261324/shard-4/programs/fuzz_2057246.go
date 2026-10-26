package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint16; F1 int; F2 uint32 }

func main() {
    var v1 int16 = 0
    v2 := []int{-754, 32768, 128}
    v3 := make(map[uint32]int)
    for fi4 := 0; fi4 < 15; fi4++ {
        v1 -= v1
        v5 := GenS1{F0: 245, F1: 4294967295}
        _ = fi4
        _ = v5
    }
    switch "\n\t" {
    case "hello":
        var v6 float32 = ((1.5 - (float32(((0.1 + float64(float32(3.14159))) + (1e10 / ((3.14159 / 0.1) / 1.0)))) * 3.14159)) - ((-2.25 + 1.5) + float32((0.1 + (0.0 * 3.14159)))))
        _ = v6
    case "foo bar":
        var v7 bool = (!true)
        _ = v7
    case "":
        if (1e10 >= -2.25) {
            fmt.Println(v1)
            v8 := GenS1{F0: 2, F1: 42}
            var v9 int8 = -1
            _ = v8
            _ = v9
        }
    }
    _ = len(v3)
    v1 = (v1 & v1)
    if ((729 <= 65535) || false) {
        v10 := GenS1{F0: 433, F1: 4294967295}
        if false {
            if true {
                v1 = v1
                v11 := GenS2{F0: 319, F1: 32767, F2: 2147483648}
                _ = v11
            } else {
                var v12 int = -9223372036854775808
                v12 = (v12 + (v12 | v12))
                if false {
                    var v13 uint16 = 1
                    v12 += (v12 / ((v12) | 1))
                    _ = v13
                } else {
                    var v14 bool = false
                    _ = v14
                }
                _ = v12
            }
            if true {
                _ = len(v2)
            } else {
                if true {
                    switch "a" {
                    case "hello":
                        if false {
                            v1 += -1
                            for fi15 := 0; fi15 < 9; fi15++ {
                                if len(v2) > 0 {
                                    v2[0] = fi15
                                }
                                v16 := GenS1{F0: 1, F1: 18446744073709551615}
                                _ = fi15
                                _ = v16
                            }
                            var v17 int64 = -160
                            _ = v17
                        } else {
                            var v18 int32 = 811
                            _ = v18
                        }
                    default:
                        var v19 float64 = 3.14159
                        _ = v19
                    }
                    v20 := GenS1{F0: 849, F1: 1}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v20
                }
                v1 += v1
                fmt.Println(v1)
            }
        } else {
            var v21 int = 1
            _ = v21
        }
        _ = len(v3)
        _ = v10
    } else {
        for fi22 := 0; fi22 < 17; fi22++ {
            for fi23 := 0; fi23 < 12; fi23++ {
                fmt.Println(fi23)
                delete(v3, 4294967295)
                _ = fi23
            }
            _ = fi22
        }
    }
    delete(v3, 1)
    var v24 uint64 = 1
    fmt.Println(v24, v24)
    v24 = v24
    v3[1] = -435
    var v25 int16 = (-(v1))
    for fi26 := 0; fi26 < 9; fi26++ {
        v27 := GenS2{F0: 32767, F1: fi26, F2: 1}
        _ = fi26
        _ = v27
    }
    fmt.Println(v1, v25, v25)
    if ((float32(3.14159) < 0.0) || (64 == -128)) {
        v28 := GenS1{F0: 18446744073709551615, F1: 255}
        v25 -= (v25 % ((0) | 1))
        _ = v28
    }
    var v29 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
}
