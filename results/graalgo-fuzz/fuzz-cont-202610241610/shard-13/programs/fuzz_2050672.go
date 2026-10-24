package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 uint16; F2 string }
type GenS4 struct { F0 uint32; F1 string; F2 bool }
type GenS5 struct { F0 string; F1 uint32; F2 uint32 }

func main() {
    var v1 uint16 = 0
    var v2 string = ("foo bar" + ("\n\t" + "foo bar"))
    v3 := []uint32{4294967295, 2147483648, 23, 1}
    var v4 float32 = ((float32((1.5 * 3.14159)) * (1.5 - (0.0 + ((3.14159 * 0.0) * 1.5)))) + 1e10)
    var v5 string = ("" + v2)
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 uint16 = 2
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi9 := 0; fi9 < 11; fi9++ {
            fmt.Println(fi9)
            _ = fi9
        }
        _ = fi8
    }
    v2 = v2
    if true {
        v4 -= ((1e10 - v4) - (v4 + (float32((0.0 - ((3.14159 * -2.25) * (-2.25 * (((1.5 + 1e10) + (((1.5 - 3.14159) * (0.0 / ((float64(0.0) / (3.14159 + float64(v4))) / (1e10 + 0.1)))) * -2.25)) + (1e10 * 1e10)))))) * (-2.25 / v4))))
    }
    if true {
        if (529 >= -727) {
            fmt.Println(v1, v2)
        } else {
            v10 := GenS1{F0: -31, F1: 970, F2: 4294967295}
            v11 := GenS1{F0: -1, F1: 42, F2: 981}
            fmt.Println(v4, v5, v1)
            _ = v10
            _ = v11
        }
        switch (v5 + "a") {
        case "\n\t":
            for fi12 := 0; fi12 < 4; fi12++ {
                switch (fi12 & -1) {
                case 2147483647:
                    if (true && true) {
                        for fi13 := 0; fi13 < 3; fi13++ {
                            v2 = "foo bar"
                            _ = fi13
                        }
                        v14 := GenS1{F0: 42, F1: 846, F2: 1}
                        v5 = "\n\t"
                        _ = v14
                    } else {
                        switch ("a" + v5) {
                        case "":
                            if false {
                                if ((3.14159 != 0.0) || (!((-749 <= -9223372036854775808) || false))) {
                                    var v15 float32 = v4
                                    for fi16 := 0; fi16 < 17; fi16++ {
                                        var v17 uint64 = 206
                                        v17 += (^(v17))
                                        _ = fi16
                                        _ = v17
                                    }
                                    v3 = append(v3, 1)
                                    _ = v15
                                } else {
                                    v2 = ((v5 + "foo bar") + v5)
                                    v18 := GenS1{F0: 42, F1: 4294967295, F2: uint64(v1)}
                                    _ = v18
                                }
                            }
                        case "a":
                            for fi19 := 0; fi19 < 9; fi19++ {
                                var v20 uint32 = 140
                                _ = fi19
                                _ = v20
                            }
                        }
                        fmt.Println(v5)
                    }
                case -1:
                    v5 = "\n\t"
                case 32768:
                    v2 = v5
                default:
                    v21 := GenS1{F0: 2147483647, F1: 2147483648, F2: 18446744073709551615}
                    _ = v21
                }
                _ = fi12
            }
        }
        var v22 uint = 432
        _ = v22
    }
    if true {
        v2 = v5
    } else {
        for fi23 := 0; fi23 < 14; fi23++ {
            var v24 int32 = 0
            for fi25 := 0; fi25 < 1; fi25++ {
                v24 = (v24 << ((65535) & 63))
                _ = fi25
            }
            _ = fi23
            _ = v24
        }
        fmt.Println(v2)
        var v26 uint = 79
        _ = v26
    }
    var v27 int16 = 338
    if false {
        var v28 int64 = 469
        _ = v28
    } else {
        switch "\n\t" {
        case "foo bar":
            for fi29 := 0; fi29 < 6; fi29++ {
                switch (fi29 | 32768) {
                case -9223372036854775808:
                    v3 = append(v3, 42)
                case -1:
                    if (((((((3.14159 - (0.1 * 0.0)) / ((0.0 * ((float64(v4) - (0.1 * float64(v4))) + (float64((1e10 - 1e10)) + 0.1))) - (float64(v4) * ((float64(1.5) + (0.1 / 1e10)) * float64((float32(1.5) + float32((0.0 - (-2.25 * ((0.1 - ((-2.25 - (0.0 / -2.25)) - (((-2.25 - ((1e10 / 0.1) - 0.0)) - 1e10) / 3.14159))) * 1.5)))))))))) / float64(v4)) / 0.1) == (0.0 - 3.14159)) || ((-79 != 0) && true)) && (60 == -1)) {
                        v30 := GenS2{F0: 31}
                        _ = v30
                    } else {
                        for fi31 := 0; fi31 < 2; fi31++ {
                            var v32 int64 = 137
                            _ = fi31
                            _ = v32
                        }
                        for fi33 := 0; fi33 < 5; fi33++ {
                            _ = len(v3)
                            _ = fi33
                        }
                        if true {
                            v34 := GenS3{F0: -128, F1: (v1 ^ (v1 >> ((65535) & 63))), F2: v2}
                            for fi35 := 0; fi35 < 20; fi35++ {
                                v36 := GenS4{F0: 2, F1: v5, F2: (fi29 != 0)}
                                v37 := GenS2{F0: 4294967295}
                                v2 = v5
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                            v27 = 0
                            _ = v34
                        } else {
                            if len(v3) > 0 {
                                v3[0] = uint32(v1)
                            }
                        }
                    }
                default:
                    if len(v3) > 0 {
                        v3[0] = 757
                    }
                }
                _ = fi29
            }
        case "a":
            v2 = v5
        case "\n\t":
            fmt.Println(v4, v5)
        default:
            v38 := GenS5{F0: v2, F1: 588, F2: 551}
            _ = v38
        }
    }
    var v39 uint16 = v1
    if len(v3) > 0 {
        v3[0] = 623
    }
    for fi40 := 0; fi40 < 6; fi40++ {
        fmt.Println(fi40, fi40)
        v27 += (v27 >> ((65535) & 63))
        _ = fi40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
