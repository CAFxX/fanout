package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int; F1 int; F2 uint8 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 string; F1 uint64; F2 int32 }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 uint16; F1 int64; F2 uint }
type GenS7 struct { F0 uint64; F1 int32; F2 uint64 }
func genFn0(p0 int64, p1 int64, p2 int32) int16 {
    fmt.Println(p0)
    var v7 int8 = 41
    fmt.Println(p1, p1, v7)
    for fi8 := 0; fi8 < 9; fi8++ {
        v7 -= (v7 % (((v7 ^ (^(v7)))) | 1))
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    var v1 float32 = (((-2.25 + float32((((((0.1 / -2.25) / 1e10) - 0.1) / 0.1) * 3.14159))) * -2.25) / -2.25)
    var v2 int = 0
    var v3 int64 = -1
    v4 := []uint32{128, 518, 4294967295}
    var v5 int8 = -99
    var v6 int32 = 0
    v4 = append(v4, 2)
    if (v2 > (v2 | v2)) {
        var v9 uint8 = 0
        var v10 float32 = (-2.25 * (v1 + v1))
        _ = v10
        _ = v9
    } else {
        v5 -= -1
        for fi11 := 0; fi11 < 8; fi11++ {
            v4 = append(v4, 2147483647)
            if len(v4) > 0 {
                v4[0] = 2147483648
            }
            _ = fi11
        }
        var v12 uint = 2
        _ = v12
    }
    v13 := GenS1{F0: 18446744073709551615, F1: 923}
    if true {
        var v14 uint32 = 467
        if true {
            var v15 int64 = (v3 ^ 307)
            var v16 uint8 = 151
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v15
            _ = v16
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v14
    }
    fmt.Println(v6)
    v17 := GenS1{F0: 42, F1: int16(v2)}
    v18 := GenS2{F0: (v2 << ((4294967295) & 63)), F1: v2, F2: 42}
    switch ("foo bar" + ("foo bar" + ("foo bar" + ("\n\t" + "a")))) {
    case "\n\t":
        var v19 uint64 = 18446744073709551615
        _ = v19
    case "a":
        switch "hello" {
        case "\n\t":
            if (false || true) {
                v5 = (v5 / ((v5) | 1))
                switch -1 {
                case 42:
                    v3 = v3
                case 32767:
                    var v20 int64 = (v3 | v3)
                    _ = v20
                case -163:
                    if (int64(v5) > (v3 >> ((18446744073709551615) & 63))) {
                        var v21 uint16 = 323
                        if false {
                            for fi22 := 0; fi22 < 9; fi22++ {
                                if ((!(255 <= 2)) || true) {
                                    v23 := GenS2{F0: (^(v2)), F1: 32768, F2: 2}
                                    if (((v6 * 0) < v6) && true) {
                                        v24 := GenS2{F0: v2, F1: (^(fi22)), F2: 0}
                                        var v25 uint16 = (v21 | (v21 / (((v21 / (((v21 * 255)) | 1))) | 1)))
                                        if (("" + "foo bar") >= "hello") {
                                            var v26 float32 = (((v1 / float32(0.1)) * (v1 * ((v1 + (v1 + float32((0.0 + 0.0)))) / v1))) - v1)
                                            fmt.Println(v2, v2, fi22)
                                            _ = v26
                                        } else {
                                            if (v6 <= (-(v6))) {
                                                if true {
                                                    v27 := GenS3{F0: 18446744073709551615, F1: fi22}
                                                    v28 := GenS3{F0: 1, F1: fi22}
                                                    var v29 int32 = (-(v6))
                                                    _ = v27
                                                    _ = v28
                                                    _ = v29
                                                }
                                            }
                                            v3 = (v3 % (((v3 ^ (v3 >> ((157) & 63)))) | 1))
                                        }
                                        _ = v24
                                        _ = v25
                                    }
                                    _ = v23
                                } else {
                                    var v30 uint8 = 128
                                    if ((v1 - 1.5) <= 0.0) {
                                        if false {
                                            v31 := GenS4{F0: (((("\n\t" + ("hello" + "")) + "a") + "a") + "foo bar"), F1: 0, F2: -2147483648}
                                            switch fi22 {
                                            case 701:
                                                var v32 int16 = int16(v2)
                                                _ = v32
                                            default:
                                                var v33 string = "hello"
                                                _ = v33
                                            }
                                            _ = v31
                                        } else {
                                            v30 += uint8(v21)
                                            var v34 bool = ((!(2 == 845)) && true)
                                            for fi35 := 0; fi35 < 11; fi35++ {
                                                v1 += v1
                                                v6 += v6
                                                _ = fi35
                                            }
                                            _ = v34
                                        }
                                        var v36 int32 = (^(v6))
                                        v37 := GenS5{F0: fi22, F1: (v36 % (((v6 * int32(fi22))) | 1))}
                                        _ = v36
                                        _ = v37
                                    }
                                    for fi38 := 0; fi38 < 2; fi38++ {
                                        var v39 int = fi22
                                        if len(v4) > 0 {
                                            v4[0] = 65535
                                        }
                                        _ = fi38
                                        _ = v39
                                    }
                                    _ = v30
                                }
                                v6 -= 42
                                _ = fi22
                            }
                            v40 := GenS6{F0: v21, F1: 1, F2: 18446744073709551615}
                            _ = v40
                        }
                        var v41 uint = 4294967295
                        _ = v21
                        _ = v41
                    }
                default:
                    v4 = append(v4, 4294967295)
                }
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        case "hello":
            v2 -= 42
        case "a":
            for fi42 := 0; fi42 < 9; fi42++ {
                v43 := GenS5{F0: v2, F1: 1}
                v44 := GenS7{F0: 2, F1: 2147483647, F2: 933}
                _ = fi42
                _ = v43
                _ = v44
            }
        }
    }
    if (18446744073709551615 != 27) {
        fmt.Println(v2)
    } else {
        v5 += -109
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
