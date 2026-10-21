package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 string; F2 uint }
type GenS5 struct { F0 uint; F1 bool }

func main() {
    v1 := []uint32{469, 2147483647}
    var v2 float64 = ((0.1 + (1.5 - 3.14159)) * 3.14159)
    v3 := []string{}
    v4 := []int{}
    var v5 uint32 = 1
    var v6 uint16 = 1
    var v7 int64 = -9223372036854775808
    var v8 int16 = 32767
    v7 = (v7 * 0)
    var v9 string = ""
    v8 -= 1000
    var v10 int8 = 46
    v11 := GenS1{F0: v10}
    if true {
        v9 = v9
    }
    switch 42 {
    case 4294967295:
        for fi12 := 0; fi12 < 6; fi12++ {
            if (!(1.5 > 0.0)) {
                if false {
                    switch fi12 {
                    case 1:
                        for fi13 := 0; fi13 < 4; fi13++ {
                            if false {
                                v7 -= (v7 >> ((18446744073709551615) & 63))
                                switch (v9 + v9) {
                                case "foo bar":
                                    fmt.Println(fi13, fi13)
                                }
                            } else {
                                v14 := GenS2{F0: 65535, F1: v9}
                                v7 += (v7 | v7)
                                fmt.Println(fi12, fi13)
                                fmt.Println(v5, v10, v2)
                                _ = v14
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v15 float32 = float32(0.1)
                            var v16 int64 = (^(v7))
                            v17 := GenS3{F0: (v10 & 0)}
                            var v18 bool = (v2 <= ((v2 + 1e10) / float64((1.5 - (1e10 - float32((v2 * (((v2 + (float64((v15 + v15)) / (3.14159 + (v2 * ((((v2 / 3.14159) + ((0.0 * (v2 / ((v2 + (v2 - 3.14159)) * (((v2 / (v2 - v2)) + (v2 + v2)) - v2)))) + ((v2 * v2) * (float64((((float32(3.14159) / 0.1) / (1.5 + (v15 + v15))) - v15)) - (-2.25 / -2.25))))) - 3.14159) / v2))))) * float64(0.0)) + v2))))))))
                            fmt.Println(fi12, v16)
                            v5 += (^(v5))
                            if (1.5 == (float64(v15) / (((v2 + float64(v15)) / v2) / (v2 / (v2 - 1.5))))) {
                                var v19 uint32 = v5
                                if v18 {
                                    if (2 <= 521) {
                                        var v20 int16 = v8
                                        v21 := GenS3{F0: int8(v20)}
                                        _ = v20
                                        _ = v21
                                    } else {
                                        var v22 uint64 = 42
                                        if v18 {
                                            v2 = float64((0.0 + (1e10 - ((v15 * float32(1.5)) + float32((((1.5 - float64(-2.25)) * v2) * v2))))))
                                            var v23 string = v9
                                            _ = v23
                                        } else {
                                            fmt.Println(v15)
                                        }
                                        var v24 uint64 = v22
                                        _ = v22
                                        _ = v24
                                    }
                                    v15 -= 0.0
                                }
                                _ = v19
                            } else {
                                var v25 uint = 1
                                _ = v25
                            }
                            _ = fi13
                            _ = v15
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                    case 42:
                        if (!false) {
                            if ((^(fi12)) > (fi12 << ((1) & 63))) {
                                v26 := GenS4{F0: ((("a" + (v9 + v9)) + v9) + "a"), F1: ((v9 + v9) + ("foo bar" + v9)), F2: 18446744073709551615}
                                _ = v26
                            }
                            v9 = v9
                        } else {
                            if false {
                                var v27 float64 = v2
                                switch "foo bar" {
                                case "foo bar":
                                    var v28 int16 = int16(v10)
                                    _ = v28
                                case "a":
                                    for fi29 := 0; fi29 < 1; fi29++ {
                                        if false {
                                            v7 += (v7 | v7)
                                        }
                                        v30 := GenS5{F0: uint(v6), F1: false}
                                        v6 = (v6 - v6)
                                        _ = fi29
                                        _ = v30
                                    }
                                default:
                                    v1 = append(v1, v5)
                                }
                                _ = v27
                            }
                        }
                    case 9223372036854775807:
                        var v31 uint8 = 8
                        _ = v31
                    default:
                        fmt.Println(v8, v10)
                    }
                    var v32 int8 = (v10 ^ int8(v6))
                    _ = v32
                } else {
                    v33 := GenS3{F0: v10}
                    var v34 float32 = (-2.25 * 1.5)
                    _ = v33
                    _ = v34
                }
            }
            v7 = -772
            var v35 int16 = -701
            _ = fi12
            _ = v35
        }
    case -1:
        var v36 uint16 = (v6 << ((65535) & 63))
        _ = v36
    }
    var v37 bool = ((false || ((v6 << ((293) & 63)) >= 65535)) && (true && false))
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    if false {
        var v38 int16 = (^(v8))
        if v37 {
            var v39 uint = 1
            _ = v39
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v10 = v10
            for fi40 := 0; fi40 < 3; fi40++ {
                v41 := GenS3{F0: (v10 / (((v10 >> ((1) & 63))) | 1))}
                var v42 float64 = 3.14159
                _ = fi40
                _ = v41
                _ = v42
            }
        }
        _ = v38
    }
    v8 -= v8
    v7 = (v7 ^ v7)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
