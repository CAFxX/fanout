package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 bool; F1 int8; F2 uint64 }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 int8; F1 int32; F2 string }
type GenS7 struct { F0 uint8; F1 int8 }
type GenS8 struct { F0 uint8; F1 int }
type GenS9 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS10 struct { F0 uint64; F1 int32 }
func genFn0(p0 uint, p1 uint16, p2 bool) uint32 {
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 int = fi6
        fmt.Println(v7)
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: (p1 ^ (^(p1)))}
    fmt.Println(p1, p1, p2)
    v9 := GenS1{F0: (p1 / ((p1) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = 200
    var v3 int8 = -57
    var v4 uint = 1
    var v5 uint64 = 42
    v2 = v2
    v10 := []int{}
    var v11 float32 = 3.14159
    var v12 int32 = 168
    v5 = (v5 >> ((uint(v5)) & 63))
    v4 -= (v4 % ((v4) | 1))
    v13 := GenS1{F0: 471}
    v14 := v1[914]
    switch 1 {
    case 128:
        v15 := GenS2{F0: v2, F1: (^(v5)), F2: 2}
        _ = v15
    case -896:
        switch ("hello" + "hello") {
        case "foo bar":
            if false {
                v16 := v1[v14]
                _ = v16
            } else {
                if (!false) {
                    v2 -= v2
                }
            }
        case "hello":
            v2 -= 42
        }
    }
    var v17 float32 = v11
    _ = len(v10)
    v18 := GenS2{F0: (v2 ^ (v2 ^ (v2 - v2))), F1: 59, F2: 1}
    if ((false || false) || true) {
        var v19 int32 = 524
        _ = v19
    } else {
        if (v2 != v2) {
            for fi20 := 0; fi20 < 15; fi20++ {
                fmt.Println(v11, v2)
                _ = fi20
            }
            if true {
                if true {
                    switch -2147483648 {
                    case 192:
                        v14 += 65535
                    case -2147483648:
                        v3 -= 0
                    default:
                        var v21 float64 = (1.5 - -2.25)
                        _ = v21
                    }
                    delete(v1, (v14 << (((^(v4))) & 63)))
                    switch -393 {
                    case 32768:
                        for fi22 := 0; fi22 < 18; fi22++ {
                            fmt.Println(v3, v3)
                            _ = fi22
                        }
                    case -1:
                        v23 := GenS2{F0: v2, F1: (v5 / ((v5) | 1)), F2: 32767}
                        _ = v23
                    }
                }
                v4 += v4
            }
        } else {
            for fi24 := 0; fi24 < 11; fi24++ {
                if true {
                    _ = len(v1)
                    fmt.Println(v17)
                }
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                for fi25 := 0; fi25 < 1; fi25++ {
                    switch "a" {
                    case "":
                        var v26 float32 = ((((v17 + float32(0.1)) / 1.0) * 1e10) / v11)
                        _ = v26
                    case "foo bar":
                        v27 := GenS2{F0: -9223372036854775808, F1: v5, F2: 906}
                        v11 = 0.1
                        _ = v27
                    default:
                        if false {
                            var v28 float64 = 1e10
                            var v29 int32 = (v12 >> (((v4 ^ v4)) & 63))
                            if (((!false) && false) || ((1e10 + v17) < v11)) {
                                v14 = (^(v14))
                                v3 -= 126
                            } else {
                                var v30 int8 = (v3 >> ((255) & 63))
                                _ = len(v1)
                                if (1 < int32(v2)) {
                                    v31 := v1[(v14 + (^(v14)))]
                                    switch (-(fi25)) {
                                    case 49:
                                        v4 = (^(v4))
                                    case -953:
                                        switch 593 {
                                        case 9223372036854775807:
                                            v32 := v1[v14]
                                            _ = v32
                                        case 2147483647:
                                            if true {
                                                var v33 uint32 = (^(v14))
                                                _ = v33
                                            }
                                        case 42:
                                            var v34 uint = (v4 % (((v4 % ((130) | 1))) | 1))
                                            _ = v34
                                        }
                                    }
                                    _ = v31
                                }
                                _ = v30
                            }
                            _ = v28
                            _ = v29
                        }
                    }
                    delete(v1, (v14 / ((v14) | 1)))
                    _ = fi25
                }
                _ = fi24
            }
            for fi35 := 0; fi35 < 16; fi35++ {
                fmt.Println(fi35)
                for fi36 := 0; fi36 < 20; fi36++ {
                    v10 = append(v10, 127)
                    var v37 bool = true
                    _ = fi36
                    _ = v37
                }
                v5 = (v5 ^ (v5 >> (((v4 - (v4 * v4))) & 63)))
                _ = fi35
            }
            v1[genFn0((^(v4)), 636, (true && (!true)))] = 0
        }
    }
    for fi38 := 0; fi38 < 10; fi38++ {
        if (true || false) {
            for fi39 := 0; fi39 < 18; fi39++ {
                switch "foo bar" {
                case "\n\t":
                    v11 = (v17 - float32(0.1))
                case "foo bar":
                    delete(v1, (v14 ^ (v14 | v14)))
                default:
                    if ((genFn0(v4, 649, false) == (v14 >> ((1) & 63))) || true) {
                        fmt.Println(v11, v3)
                    } else {
                        v40 := GenS2{F0: (v2 & (v2 | (-(v2)))), F1: v5, F2: 834}
                        fmt.Println(v12, v11, v3)
                        if (!false) {
                            v10 = append(v10, (fi38 & (fi38 | 4294967295)))
                            _ = len(v10)
                            v5 += v5
                            v41 := GenS1{F0: 477}
                            _ = v41
                        } else {
                            v3 -= (v3 >> (((v4 % (((v4 - (v4 & v4))) | 1))) & 63))
                        }
                        if (!((v14 % ((uint32(v5)) | 1)) != v14)) {
                            _ = len(v1)
                            v14 -= v14
                            v17 += 1.5
                        } else {
                            if (false || false) {
                                fmt.Println(fi38, v14)
                                _ = len(v1)
                                var v42 float32 = ((v17 / v11) + v17)
                                fmt.Println(v5, v4, v5)
                                _ = v42
                            } else {
                                switch "\n\t" {
                                case "foo bar":
                                    v5 -= v5
                                case "\n\t":
                                    v43 := GenS3{F0: uint16(v5), F1: 1}
                                    _ = v43
                                case "a":
                                    v4 = 255
                                default:
                                    if false {
                                        if false {
                                            v4 = v4
                                            switch ("hello" + "a") {
                                            case "":
                                                if ((^(v4)) == v4) {
                                                    if false {
                                                        var v44 uint8 = 223
                                                        _ = v44
                                                    }
                                                    var v45 int32 = v12
                                                    var v46 int32 = 812
                                                    _ = v45
                                                    _ = v46
                                                } else {
                                                    v1[(^(v14))] = v14
                                                }
                                            default:
                                                if ((v12 ^ -1) != -2147483648) {
                                                    if true {
                                                        v11 = float32(-2.25)
                                                    } else {
                                                        v14 += v14
                                                        fmt.Println(v2, v4, v4)
                                                        v3 -= (v3 << ((v4) & 63))
                                                    }
                                                    var v47 uint8 = 128
                                                    _ = v47
                                                }
                                            }
                                        }
                                    } else {
                                        fmt.Println(v5, fi38)
                                        v48 := GenS3{F0: 0, F1: (v12 / ((v12) | 1))}
                                        _ = v48
                                    }
                                }
                            }
                            var v49 int8 = (v3 & (v3 / ((v3) | 1)))
                            v50 := GenS3{F0: 65535, F1: (v12 + 0)}
                            fmt.Println(v3, fi38, v17)
                            var v51 uint = v4
                            _ = v49
                            _ = v50
                            _ = v51
                        }
                        v10 = append(v10, fi39)
                        v52 := GenS3{F0: 740, F1: 0}
                        if false {
                            if (42 <= int16(v4)) {
                                if len(v10) > 0 {
                                    v10[0] = -1
                                }
                                switch "\n\t" {
                                case "\n\t":
                                    if (42 < 203) {
                                        v11 += (float32((-2.25 * (float64(float32(0.0)) + 0.1))) * v11)
                                        v11 -= float32((1e10 - 3.14159))
                                        v53 := GenS4{F0: (false && (!((!(4294967295 <= 42)) || false))), F1: (v3 << (((v4 >> ((v4) & 63))) & 63)), F2: 725}
                                        _ = v53
                                    } else {
                                        var v54 int32 = (v12 + (v12 % ((v12) | 1)))
                                        v4 = 251
                                        _ = v54
                                    }
                                case "a":
                                    v55 := GenS5{F0: 106, F1: (v3 - v3)}
                                    _ = v55
                                case "hello":
                                    var v56 float64 = (-2.25 - 0.0)
                                    _ = v56
                                }
                                if false {
                                    var v57 uint = 198
                                    var v58 float32 = (v11 + (v17 / v17))
                                    _ = v57
                                    _ = v58
                                }
                            }
                            if false {
                                var v59 uint32 = genFn0((^(v4)), 343, (false || true))
                                var v60 int64 = v2
                                switch 32768 {
                                case -1:
                                    fmt.Println(v4, v59)
                                case 4294967295:
                                    fmt.Println(v59)
                                }
                                _ = v59
                                _ = v60
                            } else {
                                fmt.Println(v5, v11, fi39)
                                v61 := GenS6{F0: (v3 << ((495) & 63)), F1: 1, F2: (("a" + "") + "a")}
                                _ = v61
                            }
                        } else {
                            var v62 uint64 = v5
                            v17 = v17
                            fmt.Println(v11, v3)
                            _ = v62
                        }
                        if false {
                            if true {
                                v12 = 2147483647
                                v14 = genFn0(42, 255, (!false))
                                v5 = v5
                            } else {
                                var v63 int8 = v3
                                v64 := GenS5{F0: 255, F1: -68}
                                if true {
                                    v65 := GenS1{F0: 509}
                                    switch "a" {
                                    case "\n\t":
                                        v14 = (v14 & v14)
                                    case "foo bar":
                                        var v66 uint32 = (v14 * (^(v14)))
                                        _ = v66
                                    default:
                                        switch (fi38 - fi38) {
                                        case 32768:
                                            if true {
                                                fmt.Println(v5, v14)
                                                delete(v1, (v14 ^ 267))
                                                v67 := GenS7{F0: 2, F1: v63}
                                                _ = v67
                                            } else {
                                                v10 = append(v10, -2147483648)
                                                v14 = genFn0((v4 - uint(v14)), 255, (v14 == (v14 + (v14 | genFn0(v4, 32767, (v11 == -2.25))))))
                                                if true {
                                                    v4 = v4
                                                    fmt.Println(v3, v3)
                                                }
                                            }
                                        case 0:
                                            v5 -= v5
                                        }
                                    }
                                    switch (^(fi38)) {
                                    case -244:
                                        v68 := GenS5{F0: 0, F1: -1}
                                        v69 := GenS6{F0: v63, F1: v12, F2: "a"}
                                        switch "a" {
                                        case "hello":
                                            v10 = append(v10, int(v12))
                                        }
                                        _ = v68
                                        _ = v69
                                    case -9223372036854775808:
                                        fmt.Println(v11)
                                    default:
                                        v70 := GenS8{F0: 128, F1: (fi39 % ((fi39) | 1))}
                                        _ = v70
                                    }
                                    fmt.Println(v14)
                                    _ = v65
                                }
                                _ = v63
                                _ = v64
                            }
                            if len(v10) > 0 {
                                v10[0] = 1
                            }
                            v4 -= (^(v4))
                            if true {
                                switch 398 {
                                case -2147483648:
                                    v11 = v17
                                case -1:
                                    v71 := GenS9{F0: v12, F1: 255, F2: 122}
                                    _ = v71
                                default:
                                    var v72 uint64 = v5
                                    _ = v72
                                }
                                var v73 uint64 = v5
                                _ = v73
                            } else {
                                var v74 int = (^(fi38))
                                v75 := GenS2{F0: v2, F1: v5, F2: 2}
                                _ = v74
                                _ = v75
                            }
                            v3 = -1
                            v12 = (v12 * 42)
                        }
                        v12 = v12
                        _ = v40
                        _ = v52
                    }
                }
                v76 := GenS10{F0: v5, F1: 0}
                v11 -= v11
                _ = fi39
                _ = v76
            }
        }
        v12 -= (v12 - v12)
        _ = fi38
    }
    fmt.Println(v12, v12)
    v11 = ((v17 * (v17 - v17)) - float32(1.5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
