package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int16; F1 uint32; F2 int }
type GenS3 struct { F0 uint64; F1 int16; F2 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 int; F1 int16 }
type GenS7 struct { F0 bool }
func genFn0(p0 uint, p1 int16, p2 string) uint32 {
    v5 := []int{-713, 1, -566, 128}
    v6 := GenS1{F0: 2, F1: (p0 / (((p0 << ((p0) & 63))) | 1))}
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int16 = p1
        if (((((((1e10 * 3.14159) / (float64((-2.25 / (0.1 - 1.5))) + 1.5)) + 1e10) / 3.14159) * (3.14159 - -2.25)) / -2.25) != 0.1) {
            if (1 != 42) {
                if true {
                    p0 = 1
                }
                fmt.Println(p2)
                p1 = (p1 % ((-32768) | 1))
            } else {
                fmt.Println(p2)
                var v9 int8 = 42
                v5 = append(v5, 494)
                _ = v9
            }
            switch "\n\t" {
            case "":
                for fi10 := 0; fi10 < 3; fi10++ {
                    var v11 int16 = (p1 % (((p1 ^ (p1 & 876))) | 1))
                    _ = fi10
                    _ = v11
                }
            case "\n\t":
                var v12 uint16 = 0
                _ = v12
            case "foo bar":
                v5 = append(v5, fi7)
            }
        } else {
            fmt.Println(p1)
            for fi13 := 0; fi13 < 12; fi13++ {
                v14 := GenS2{F0: (-(v8)), F1: 2, F2: (fi7 ^ fi7)}
                var v15 uint16 = 141
                if (fi13 > (-(fi7))) {
                    p1 -= (p1 << ((660) & 63))
                    switch p2 {
                    case "":
                        v8 = 32767
                    case "\n\t":
                        p2 = ("" + p2)
                    default:
                        if ((((0 < 1) || ("hello" >= ("foo bar" + ""))) || (false && (!(-32768 == p1)))) && (!false)) {
                            switch "" {
                            case "a":
                                var v16 int = 32768
                                _ = v16
                            default:
                                v8 = 100
                            }
                            if (true && false) {
                                switch -9223372036854775808 {
                                case 42:
                                    v17 := GenS3{F0: 18446744073709551615, F1: 32767, F2: int64(p1)}
                                    _ = v17
                                case 1:
                                    var v18 uint16 = 65535
                                    _ = v18
                                default:
                                    p0 += (p0 * p0)
                                }
                                if false {
                                    p2 = p2
                                    var v19 int8 = 0
                                    var v20 string = "foo bar"
                                    _ = v19
                                    _ = v20
                                }
                            } else {
                                if (-1 > -347) {
                                    var v21 uint32 = 2147483647
                                    v15 = v15
                                    var v22 int16 = v8
                                    fmt.Println(fi13, p1, p0)
                                    v5 = append(v5, fi7)
                                    _ = len(v5)
                                    _ = v21
                                    _ = v22
                                } else {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = len(v5)
                                    if len(v5) > 0 {
                                        v5[0] = (fi7 >> ((30) & 63))
                                    }
                                    v23 := GenS3{F0: 212, F1: (p1 ^ v8), F2: -1}
                                    if true {
                                        fmt.Println(v8, p1)
                                        if (!true) {
                                            var v24 uint32 = 1
                                            _ = v24
                                        }
                                        p0 += 461
                                        if false {
                                            v5 = append(v5, -384)
                                            fmt.Println(p0, fi13)
                                        }
                                    }
                                    var v25 string = p2
                                    p2 = "hello"
                                    _ = v23
                                    _ = v25
                                }
                            }
                        } else {
                            var v26 uint64 = 76
                            var v27 float32 = float32(float64((((1e10 - 1e10) * 1.5) / float32(-2.25))))
                            _ = len(v5)
                            _ = v26
                            _ = v27
                        }
                    }
                    fmt.Println(fi7)
                } else {
                    fmt.Println(p1, p2)
                    v15 += v15
                    v8 = 584
                }
                v28 := GenS4{F0: 2}
                var v29 uint16 = 241
                _ = fi13
                _ = v14
                _ = v15
                _ = v28
                _ = v29
            }
            for fi30 := 0; fi30 < 3; fi30++ {
                var v31 int8 = 127
                _ = fi30
                _ = v31
            }
        }
        p2 = p2
        _ = fi7
        _ = v8
    }
    switch 1 {
    case 32767:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case 127:
        for fi32 := 0; fi32 < 16; fi32++ {
            var v33 uint = p0
            var v34 int32 = 0
            if false {
                p0 += (^(p0))
            } else {
                v35 := GenS2{F0: (p1 * p1), F1: 1, F2: 444}
                for fi36 := 0; fi36 < 12; fi36++ {
                    p0 -= (p0 % ((478) | 1))
                    _ = fi36
                }
                _ = v35
            }
            _ = fi32
            _ = v33
            _ = v34
        }
    case -50:
        fmt.Println(p0, p0, p2)
    }
    v37 := GenS4{F0: 654}
    _ = p0
    _ = p1
    _ = p2
    _ = v37
    _ = v5
    _ = v6
    return 816
}
func genFn1(p0 bool) uint64 {
    var v38 uint8 = 0
    v38 += v38
    switch 127 {
    case 32768:
        var v39 int16 = 42
        _ = v39
    case 870:
        p0 = p0
    default:
        if (0.0 <= (float64((1.5 / 3.14159)) + -2.25)) {
            for fi40 := 0; fi40 < 3; fi40++ {
                var v41 string = ""
                _ = fi40
                _ = v41
            }
            p0 = false
        }
    }
    for fi42 := 0; fi42 < 1; fi42++ {
        switch "a" {
        case "\n\t":
            for fi43 := 0; fi43 < 18; fi43++ {
                var v44 uint64 = 316
                _ = fi43
                _ = v44
            }
        default:
            fmt.Println(fi42, v38)
        }
        p0 = (42 < -64)
        _ = fi42
    }
    _ = p0
    _ = v38
    return 925
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[string]uint32)
    var v3 uint = 0
    v4 := make(map[uint32]int)
    v45 := GenS3{F0: 18446744073709551615, F1: -32768, F2: -9223372036854775808}
    if ((3.14159 * ((1.5 - float32(0.1)) + (((0.0 - (0.1 / (((float32(-2.25) * -2.25) + 1e10) + (3.14159 + float32((float64((float32((0.0 + float64(((0.0 + ((1e10 * 1e10) - 3.14159)) * 0.1)))) + float32(1e10))) * 0.1)))))) + 1e10) - 1e10))) >= -2.25) {
        fmt.Println(v1, v1)
        var v46 bool = (1 != 65535)
        _ = v46
    } else {
        var v47 int8 = 100
        v48 := GenS1{F0: 490, F1: 646}
        _ = v47
        _ = v48
    }
    if ((v1 > 828) && (816 > 42)) {
        v49 := GenS2{F0: v1, F1: 1, F2: 128}
        var v50 int64 = 1
        v51 := GenS5{F0: false, F1: (v3 / ((18446744073709551615) | 1))}
        _ = v49
        _ = v50
        _ = v51
    }
    v52 := GenS6{F0: 42, F1: (v1 - (v1 + v1))}
    var v53 string = "foo bar"
    v54 := []int{42, 9223372036854775807}
    v4[2147483648] = 128
    v55 := GenS7{F0: false}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v54:", len(v54), v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v52
    _ = v53
    _ = v54
    _ = v55
}
