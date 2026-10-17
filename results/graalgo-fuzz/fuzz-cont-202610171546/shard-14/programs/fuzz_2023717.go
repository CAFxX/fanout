package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint16; F1 int16 }
func genFn0() uint16 {
    var v5 uint16 = 35
    v5 = (v5 % ((32767) | 1))
    var v6 float64 = float64((0.1 + 1e10))
    if true {
        if false {
            v7 := []int{-363}
            v6 = v6
            _ = v7
        }
    } else {
        for fi8 := 0; fi8 < 11; fi8++ {
            v5 = v5
            for fi9 := 0; fi9 < 1; fi9++ {
                v10 := GenS1{F0: 18446744073709551615, F1: 255}
                var v11 int8 = -16
                _ = fi9
                _ = v10
                _ = v11
            }
            var v12 uint32 = 2147483647
            _ = fi8
            _ = v12
        }
    }
    if (!(694 != 277)) {
        v6 += v6
    }
    _ = v5
    _ = v6
    return v5
}

func main() {
    var v1 uint64 = 2
    var v2 int16 = -32768
    var v3 int64 = -363
    v4 := make(map[int]int)
    var v13 int32 = 1
    for fi14 := 0; fi14 < 8; fi14++ {
        fmt.Println(v2)
        _ = len(v4)
        _ = fi14
    }
    if true {
        fmt.Println(v13, v3)
        fmt.Println(v3, v2, v1)
        v15 := GenS1{F0: (v1 * 18446744073709551615), F1: 140}
        _ = v15
    }
    if ((false && ((true && false) || (!true))) && (int(v3) > 0)) {
        if (0 < 1) {
            var v16 int = 42
            _ = v16
        }
        for fi17 := 0; fi17 < 1; fi17++ {
            for fi18 := 0; fi18 < 6; fi18++ {
                var v19 float32 = 1e10
                v13 += (v13 - v13)
                if true {
                    switch -9223372036854775808 {
                    case 42:
                        var v20 int64 = (v3 * (v3 + (v3 / ((int64(v1)) | 1))))
                        _ = v20
                    case 480:
                        if (!(((float32(1e10) / (v19 - v19)) * v19) == v19)) {
                            switch fi17 {
                            case 32767:
                                fmt.Println(v1)
                            case 128:
                                v21 := GenS1{F0: 18446744073709551615, F1: 127}
                                _ = v21
                            default:
                                var v22 uint16 = 820
                                _ = v22
                            }
                            var v23 int = (fi18 % ((138) | 1))
                            _ = len(v4)
                            _ = v23
                        } else {
                            if (false && true) {
                                v19 = v19
                                v1 -= v1
                            } else {
                                var v24 int32 = v13
                                _ = v24
                            }
                        }
                    default:
                        fmt.Println(fi17, v19)
                    }
                } else {
                    var v25 float32 = v19
                    if (!true) {
                        fmt.Println(v3)
                        switch (fi18 << ((4294967295) & 63)) {
                        case -393:
                            switch ("a" + (("" + "") + ("\n\t" + "foo bar"))) {
                            case "hello":
                                v2 = 1
                            case "a":
                                fmt.Println(v19, v1, v3)
                            case "\n\t":
                                var v26 int16 = (-(v2))
                                _ = v26
                            }
                        }
                        if false {
                            v27 := GenS2{F0: 2}
                            v28 := GenS3{F0: -11}
                            _ = v27
                            _ = v28
                        } else {
                            v3 += (v3 + -9223372036854775808)
                        }
                        var v29 uint8 = 2
                        switch "foo bar" {
                        case "":
                            if false {
                                v1 = v1
                            }
                        default:
                            switch "" {
                            case "a":
                                v4[(fi18 - 661)] = fi18
                            }
                        }
                        _ = v29
                    }
                    _ = v25
                }
                var v30 int16 = (v2 & (v2 << ((4294967295) & 63)))
                _ = len(v4)
                _ = fi18
                _ = v19
                _ = v30
            }
            for fi31 := 0; fi31 < 7; fi31++ {
                fmt.Println(fi31, v13, v3)
                v13 += v13
                var v32 int64 = v3
                _ = fi31
                _ = v32
            }
            _ = len(v4)
            _ = fi17
        }
    }
    v33 := []int{42, 127}
    fmt.Println(v2, v3)
    fmt.Println(v1, v13, v3)
    for fi34 := 0; fi34 < 3; fi34++ {
        v3 = 42
        for fi35 := 0; fi35 < 13; fi35++ {
            v1 = (^(v1))
            if (false && (false || false)) {
                fmt.Println(v2)
                _ = len(v4)
                var v36 int32 = v13
                if true {
                    fmt.Println(fi34, fi34)
                    fmt.Println(v36, v1, fi35)
                    v37 := GenS4{F0: (v36 ^ v36), F1: "\n\t"}
                    _ = v37
                } else {
                    var v38 uint32 = 0
                    _ = v38
                }
                switch "a" {
                case "a":
                    if false {
                        if true {
                            var v39 uint = 255
                            v2 = (v2 / ((v2) | 1))
                            _ = v39
                        }
                        switch fi35 {
                        case -1:
                            v40 := GenS3{F0: -116}
                            _ = v40
                        case 566:
                            v2 -= (v2 >> ((500) & 63))
                        default:
                            v33 = append(v33, (fi35 & fi35))
                        }
                        fmt.Println(fi35)
                        v41 := GenS3{F0: 42}
                        fmt.Println(v13, fi34)
                        var v42 uint16 = 648
                        v43 := GenS4{F0: v13, F1: "\n\t"}
                        _ = v41
                        _ = v42
                        _ = v43
                    } else {
                        fmt.Println(v13)
                    }
                case "foo bar":
                    if (((-2.25 - (0.0 - 1e10)) + ((-2.25 / 0.1) / (-2.25 + 3.14159))) == 3.14159) {
                        fmt.Println(v3)
                    } else {
                        v3 -= (v3 * (v3 >> ((65535) & 63)))
                        v2 += v2
                        v33 = append(v33, fi34)
                    }
                case "\n\t":
                    v44 := GenS3{F0: 79}
                    _ = v44
                default:
                    var v45 int32 = 42
                    _ = v45
                }
                v1 = v1
                if ((!false) || true) {
                    var v46 bool = false
                    if len(v33) > 0 {
                        v33[0] = (fi35 & fi34)
                    }
                    v47 := GenS2{F0: 251}
                    if true {
                        var v48 int64 = -707
                        _ = v48
                    } else {
                        v49 := GenS5{F0: 237, F1: 1}
                        v13 += (v13 ^ (-(v13)))
                        var v50 uint16 = genFn0()
                        v51 := GenS6{F0: ("a" + "a"), F1: -68}
                        var v52 uint = 65535
                        v53 := GenS7{F0: v52}
                        v54 := GenS8{F0: 0, F1: -1}
                        _ = v49
                        _ = v50
                        _ = v51
                        _ = v52
                        _ = v53
                        _ = v54
                    }
                    var v55 uint64 = v1
                    v4[fi34] = fi34
                    var v56 bool = (!(int8(v3) >= 42))
                    _ = v46
                    _ = v47
                    _ = v55
                    _ = v56
                } else {
                    var v57 uint = 638
                    _ = v57
                }
                _ = v36
            } else {
                var v58 uint64 = v1
                _ = v58
            }
            v4[fi34] = 1
            _ = fi35
        }
        _ = fi34
    }
    var v59 uint16 = 1
    if false {
        if true {
            switch 2147483647 {
            case 9223372036854775807:
                v13 += (-(v13))
            case -857:
                v1 -= 321
            case 32767:
                v59 += (v59 >> ((255) & 63))
            }
        }
    } else {
        for fi60 := 0; fi60 < 7; fi60++ {
            v59 = (v59 % ((32767) | 1))
            v13 += v13
            _ = fi60
        }
        if (42 == 52) {
            v61 := GenS6{F0: (("a" + "a") + ("hello" + "a")), F1: -1}
            var v62 uint8 = 2
            switch 767 {
            case -9223372036854775808:
                var v63 uint = 42
                _ = v63
            case 2147483647:
                var v64 int32 = v13
                _ = v64
            case 127:
                var v65 int64 = (v3 >> ((42) & 63))
                _ = v65
            default:
                for fi66 := 0; fi66 < 6; fi66++ {
                    v67 := v4[(^(fi66))]
                    v4[(v67 | (fi66 | fi66))] = (v67 / (((v67 + v67)) | 1))
                    _ = fi66
                    _ = v67
                }
            }
            _ = v61
            _ = v62
        } else {
            v68 := v4[-41]
            _ = v68
        }
    }
    var v69 uint = 337
    delete(v4, 9223372036854775807)
    switch -809 {
    case 2147483647:
        v3 += -1
    case 1:
        v69 += v69
    case 776:
        v1 += v1
    default:
        for fi70 := 0; fi70 < 8; fi70++ {
            v71 := GenS6{F0: ("\n\t" + "foo bar"), F1: 0}
            if (4294967295 != (v1 + 4294967295)) {
                v72 := v4[int(v59)]
                _ = v72
            }
            _ = fi70
            _ = v71
        }
    }
    for fi73 := 0; fi73 < 4; fi73++ {
        var v74 int16 = (v2 >> ((v69) & 63))
        _ = fi73
        _ = v74
    }
    for fi75 := 0; fi75 < 18; fi75++ {
        var v76 uint = (^(v69))
        _ = fi75
        _ = v76
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v59:", v59)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v59
    _ = v69
}
