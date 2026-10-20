package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 string }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 bool; F1 int; F2 int8 }
type GenS4 struct { F0 int32; F1 int8; F2 uint8 }
type GenS5 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int8; F1 int32 }
func genFn0(p0 int32, p1 int8) uint16 {
    var v6 int64 = 9223372036854775807
    for fi7 := 0; fi7 < 13; fi7++ {
        v6 = (v6 ^ (v6 / (((v6 & (-(v6)))) | 1)))
        _ = fi7
    }
    var v8 int32 = (-(p0))
    fmt.Println(p1)
    v9 := GenS1{F0: false, F1: 128, F2: ("a" + ("" + "hello"))}
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := []string{"\n\t", (("foo bar" + ("\n\t" + "a")) + "\n\t")}
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]int)
    var v4 int32 = 484
    v5 := make(map[int]uint32)
    v4 -= (v4 % ((v4) | 1))
    v4 = 2147483647
    var v10 uint64 = 4294967295
    v11 := GenS2{F0: -846, F1: (("" + ("" + "\n\t")) + "a")}
    v12 := GenS3{F0: (!true), F1: -1, F2: 127}
    if false {
        if true {
            v10 = 610
            for fi13 := 0; fi13 < 11; fi13++ {
                delete(v5, fi13)
                v3[fi13] = fi13
                _ = fi13
            }
        }
    }
    var v14 float32 = 3.14159
    _ = len(v1)
    v14 = (-2.25 + v14)
    for fi15 := 0; fi15 < 5; fi15++ {
        v16 := v2[65535]
        v4 -= v4
        _ = fi15
        _ = v16
    }
    fmt.Println(v10, v10, v10)
    var v17 uint = 339
    v4 = v4
    if false {
        v14 += 0.0
        fmt.Println(v10, v10, v14)
        if (1 < 2147483648) {
            if false {
                if (!(false || (!true))) {
                    v18 := GenS3{F0: ((101 < 820) && (((3.14159 - 1e10) / float64(v14)) > 1e10)), F1: -2147483648, F2: 0}
                    var v19 string = "a"
                    _ = v18
                    _ = v19
                }
                var v20 uint32 = 4294967295
                v17 -= (v17 * (v17 - 147))
                _ = v20
            }
            if (("foo bar" + "a") == "a") {
                if true {
                    switch 9223372036854775807 {
                    case -9223372036854775808:
                        var v21 uint8 = 177
                        _ = v21
                    default:
                        for fi22 := 0; fi22 < 11; fi22++ {
                            if (1000 > 32767) {
                                fmt.Println(v14)
                            }
                            _ = fi22
                        }
                    }
                    v10 = (v10 >> (((v17 + (^(v17)))) & 63))
                    for fi23 := 0; fi23 < 14; fi23++ {
                        for fi24 := 0; fi24 < 11; fi24++ {
                            if (93 != 128) {
                                if (-1 != 1) {
                                    _ = len(v3)
                                    if ((true && (((false && (2147483647 <= 830)) || (2147483647 > 42)) || (v14 != 0.1))) && (!true)) {
                                        v17 = v17
                                    } else {
                                        if (fi23 > 1) {
                                            if true {
                                                v25 := v2[2147483648]
                                                _ = v25
                                            } else {
                                                v10 -= (^(v10))
                                                v4 -= v4
                                                v26 := GenS3{F0: false, F1: fi24, F2: 100}
                                                _ = v26
                                            }
                                        }
                                        switch "\n\t" {
                                        case "foo bar":
                                            if (0 > genFn0(v4, -128)) {
                                                if ((222 > 1) || false) {
                                                    var v27 uint64 = v10
                                                    _ = v27
                                                } else {
                                                    v28 := GenS4{F0: v4, F1: 68, F2: 0}
                                                    v29 := GenS1{F0: (v17 < v17), F1: (^(fi23)), F2: ("" + "a")}
                                                    v30 := v2[1]
                                                    _ = v28
                                                    _ = v29
                                                    _ = v30
                                                }
                                                v17 += v17
                                            } else {
                                                _ = len(v1)
                                                v31 := GenS5{F0: 1, F1: 4294967295, F2: 2147483647}
                                                if (!true) {
                                                    var v32 int8 = -1
                                                    fmt.Println(v17, fi24, v4)
                                                    var v33 string = "foo bar"
                                                    _ = v32
                                                    _ = v33
                                                }
                                                _ = v31
                                            }
                                        default:
                                            var v34 uint64 = (v10 * (^(v10)))
                                            _ = v34
                                        }
                                        var v35 uint = (v17 + v17)
                                        _ = v35
                                    }
                                }
                            }
                            var v36 uint32 = 4294967295
                            if (!false) {
                                var v37 string = "\n\t"
                                if false {
                                    v38 := GenS6{F0: 689}
                                    switch v37 {
                                    case "hello":
                                        v1 = append(v1, "hello")
                                    }
                                    if true {
                                        if (1.5 < (1e10 - (3.14159 + 1e10))) {
                                            fmt.Println(v37, v4)
                                            v39 := GenS4{F0: v4, F1: 100, F2: 2}
                                            switch "" {
                                            case "\n\t":
                                                var v40 uint32 = (v36 / ((1) | 1))
                                                v14 = 1e10
                                                fmt.Println(v4, fi24)
                                                _ = v40
                                            }
                                            v41 := GenS7{F0: -56, F1: (v4 ^ 524)}
                                            _ = v39
                                            _ = v41
                                        }
                                    } else {
                                        v17 = (v17 % ((uint(v10)) | 1))
                                        v14 -= (v14 + v14)
                                    }
                                    fmt.Println(fi23, v17, v4)
                                    _ = v38
                                }
                                _ = v37
                            } else {
                                v10 = (^(v10))
                                v4 = 2147483647
                            }
                            _ = fi24
                            _ = v36
                        }
                        fmt.Println(fi23)
                        _ = fi23
                    }
                }
                _ = len(v2)
                v42 := GenS1{F0: (-683 <= 0), F1: -620, F2: ("foo bar" + "hello")}
                _ = v42
            }
        } else {
            switch "foo bar" {
            case "foo bar":
                v14 = (v14 - v14)
            default:
                fmt.Println(v10, v10, v4)
            }
            fmt.Println(v14, v4, v14)
            v14 += 1e10
        }
    }
    v14 = (float32(float64((v14 * v14))) + (v14 - (((v14 - ((float32((float64(((0.0 * (v14 - ((1.5 * v14) + v14))) - (((v14 / -2.25) + v14) * 0.0))) + 1.5)) * v14) * 0.0)) - 3.14159) + (1e10 / ((((v14 / ((v14 * v14) / 1.5)) / v14) * (1e10 / v14)) * v14)))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
