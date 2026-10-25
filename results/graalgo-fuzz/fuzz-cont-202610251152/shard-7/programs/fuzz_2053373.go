package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint64 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 uint64; F1 bool; F2 uint64 }
type GenS4 struct { F0 string; F1 int64; F2 uint32 }
type GenS5 struct { F0 int32; F1 int64; F2 int64 }

func main() {
    var v1 float64 = float64(float32((1.5 * float64(1e10))))
    var v2 int16 = -565
    var v3 uint16 = 722
    fmt.Println(v2)
    v4 := []int{}
    v5 := GenS1{F0: "foo bar", F1: v2, F2: 50}
    var v6 float32 = 0.0
    v2 = (v2 >> ((2) & 63))
    v1 = v1
    fmt.Println(v1)
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: ("foo bar" + ("a" + "a")), F1: v2, F2: 490}
        if false {
            v2 = (v2 + v2)
        } else {
            fmt.Println(v1)
            switch fi7 {
            case 191:
                v4 = append(v4, 32767)
            case 32767:
                v2 = 42
            case -172:
                v4 = append(v4, -2147483648)
            default:
                var v9 uint = 42
                _ = v9
            }
            v6 = ((1.5 + (3.14159 / (v6 - (((((v6 + 0.1) * float32(float64(v6))) * v6) + (1.5 / ((v6 - ((1e10 + 0.1) / (float32(1e10) - (0.0 / 1.0)))) - v6))) / (v6 * 0.0))))) / ((v6 * 1e10) - v6))
        }
        v6 = ((1.5 - ((3.14159 + (v6 / 1.0)) + (((v6 + 0.1) / v6) * (1e10 / (v6 - (1.5 * 0.1)))))) + (v6 * ((v6 + (float32(v1) * v6)) * v6)))
        _ = fi7
        _ = v8
    }
    if len(v4) > 0 {
        v4[0] = -200
    }
    v10 := GenS1{F0: "hello", F1: -865, F2: 42}
    v11 := GenS2{F0: 127, F1: 2147483647}
    if false {
        if (5 <= 2147483648) {
            var v12 uint = 255
            _ = v12
        } else {
            v6 -= (-2.25 - 1.5)
            for fi13 := 0; fi13 < 6; fi13++ {
                if (!((fi13 > 128) && false)) {
                    var v14 int8 = -1
                    for fi15 := 0; fi15 < 17; fi15++ {
                        if false {
                            if true {
                                if ((-(v2)) >= v2) {
                                    v6 -= v6
                                    if len(v4) > 0 {
                                        v4[0] = (fi15 + 9223372036854775807)
                                    }
                                } else {
                                    var v16 uint64 = 104
                                    _ = v16
                                }
                                v17 := GenS3{F0: 4294967295, F1: (true && true), F2: 2}
                                _ = v17
                            } else {
                                v6 = ((1.5 - v6) - float32(float64(1e10)))
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                if (v14 > int8(v2)) {
                                    if (((v3 ^ (v3 + v3)) < (v3 / (((v3 - v3)) | 1))) && (!true)) {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v1 -= 3.14159
                                        _ = len(v4)
                                        var v18 int8 = (v14 + (v14 - v14))
                                        fmt.Println(v3, v3, v6)
                                        _ = v18
                                    } else {
                                        v19 := GenS1{F0: "foo bar", F1: (v2 % ((v2) | 1)), F2: 1}
                                        v20 := GenS1{F0: "hello", F1: v2, F2: 2}
                                        _ = v19
                                        _ = v20
                                    }
                                    if false {
                                        if (1e10 < ((v1 * (v1 * ((v1 / 1.5) / 1e10))) + 0.0)) {
                                            v21 := GenS1{F0: "", F1: v2, F2: 1}
                                            if len(v4) > 0 {
                                                v4[0] = (fi15 + (fi15 << ((342) & 63)))
                                            }
                                            _ = v21
                                        } else {
                                            var v22 int32 = 0
                                            _ = v22
                                        }
                                        var v23 string = "a"
                                        var v24 string = ((("a" + v23) + v23) + v23)
                                        _ = v23
                                        _ = v24
                                    }
                                    var v25 uint32 = 128
                                    _ = v25
                                }
                            }
                            var v26 uint8 = 42
                            v6 = (((0.1 - 0.0) + v6) + v6)
                            var v27 int32 = -378
                            v28 := GenS4{F0: "\n\t", F1: -9223372036854775808, F2: 1}
                            _ = v26
                            _ = v27
                            _ = v28
                        } else {
                            v29 := GenS2{F0: v14, F1: fi13}
                            _ = v29
                        }
                        _ = fi15
                    }
                    _ = v14
                }
                _ = fi13
            }
            var v30 string = "a"
            _ = v30
        }
    }
    switch 127 {
    case -95:
        v1 -= v1
    }
    fmt.Println(v6, v3)
    v31 := GenS5{F0: 355, F1: -1, F2: 9223372036854775807}
    v3 += 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
