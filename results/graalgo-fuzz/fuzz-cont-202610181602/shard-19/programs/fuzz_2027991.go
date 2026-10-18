package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int64; F2 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 uint; F2 int16 }
type GenS6 struct { F0 int32; F1 int8; F2 uint8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint; F1 int32 }
type GenS11 struct { F0 uint16; F1 int32 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint64; F1 int64; F2 int32 }
type GenS14 struct { F0 int8; F1 int32 }
type GenS15 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 uint8 }
func genFn0() uint {
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 uint = 624
        v6 = 505
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 741}
    v8 := GenS1{F0: -600}
    var v9 uint = 4294967295
    _ = v7
    _ = v8
    _ = v9
    return (v9 * 0)
}
func genFn1(p0 uint32, p1 bool, p2 string) bool {
    var v10 int8 = -117
    fmt.Println(p1, p2, p0)
    v11 := []int{99, 1, -893}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return true
}

func main() {
    var v1 int8 = -108
    var v2 int16 = 42
    v3 := []int{}
    var v4 uint16 = 32767
    var v12 float32 = ((0.1 - 1.5) * 3.14159)
    v1 -= (v1 << ((42) & 63))
    var v13 float32 = (((1e10 + (v12 + 0.0)) * 1e10) * v12)
    v13 = v12
    var v14 float32 = -2.25
    v15 := GenS1{F0: -1}
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 uint = 2
        v3 = append(v3, fi16)
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 uint16 = (v4 + v4)
        _ = fi18
        _ = v19
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 = 113
    v13 += (v14 / ((float32((((3.14159 * 1e10) + (3.14159 * 1e10)) * 0.1)) - v14) + v13))
    if (v4 == v4) {
        for fi20 := 0; fi20 < 4; fi20++ {
            for fi21 := 0; fi21 < 19; fi21++ {
                v12 += v13
                _ = fi21
            }
            v22 := GenS2{F0: (fi20 >> ((4294967295) & 63)), F1: 1, F2: 225}
            var v23 int = fi20
            _ = fi20
            _ = v22
            _ = v23
        }
        v24 := GenS2{F0: 1, F1: 0, F2: (v4 + (v4 ^ 2))}
        var v25 int32 = -527
        _ = v24
        _ = v25
    } else {
        v4 = v4
        if false {
            v26 := GenS2{F0: -806, F1: 1, F2: v4}
            _ = v26
        } else {
            v13 = v13
            switch -923 {
            case 169:
                v27 := GenS3{F0: -678}
                _ = v27
            }
        }
        v28 := GenS3{F0: 9223372036854775807}
        _ = v28
    }
    if (false || ("a" < "\n\t")) {
        for fi29 := 0; fi29 < 12; fi29++ {
            v3 = append(v3, (fi29 / ((int(v1)) | 1)))
            var v30 uint8 = 128
            var v31 uint8 = v30
            _ = fi29
            _ = v30
            _ = v31
        }
        var v32 uint8 = 255
        v12 += v14
        _ = v32
    }
    if true {
        v4 += (v4 >> ((1) & 63))
        v33 := GenS4{F0: 42}
        _ = v33
    }
    switch 32768 {
    case 9223372036854775807:
        for fi34 := 0; fi34 < 1; fi34++ {
            switch fi34 {
            case -9223372036854775808:
                var v35 bool = ((false || false) && (!(!true)))
                _ = v35
            default:
                v12 = ((v13 - -2.25) + (v12 + ((v14 / float32(3.14159)) / (1.5 + v13))))
            }
            for fi36 := 0; fi36 < 8; fi36++ {
                if (!(!(!(0 < 42)))) {
                    v37 := GenS5{F0: -711, F1: uint(v4), F2: v2}
                    v1 = 100
                    _ = v37
                } else {
                    var v38 string = ("" + (((("" + "hello") + "hello") + ("a" + "a")) + ("a" + "hello")))
                    if (!false) {
                        var v39 float32 = v12
                        v13 = -2.25
                        fmt.Println(fi34)
                        _ = v39
                    } else {
                        if ((v38 + (v38 + (v38 + v38))) >= v38) {
                            var v40 int8 = (v1 * (v1 % ((v1) | 1)))
                            var v41 uint16 = (v4 ^ 42)
                            switch fi34 {
                            case 1:
                                v41 -= 65535
                            case 32768:
                                var v42 float32 = v13
                                _ = v42
                            default:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v43 uint = genFn0()
                                _ = v43
                            }
                            _ = v40
                            _ = v41
                        }
                        v44 := GenS6{F0: -766, F1: v1, F2: 0}
                        _ = v44
                    }
                    _ = v38
                }
                var v45 int8 = v1
                if true {
                    fmt.Println(v2)
                } else {
                    fmt.Println(v45, fi36, v13)
                    v46 := GenS7{F0: 65535}
                    _ = v46
                }
                _ = fi36
                _ = v45
            }
            _ = fi34
        }
    case 32768:
        var v47 string = (("hello" + "foo bar") + "\n\t")
        _ = v47
    case 32767:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        for fi48 := 0; fi48 < 10; fi48++ {
            var v49 uint = genFn0()
            _ = fi48
            _ = v49
        }
    }
    if (!false) {
        if false {
            fmt.Println(v12)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            if false {
                var v50 string = ((("hello" + "\n\t") + "foo bar") + ("\n\t" + (("foo bar" + "\n\t") + ((("\n\t" + "\n\t") + "") + "foo bar"))))
                v2 -= int16(v4)
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
                _ = v50
            } else {
                fmt.Println(v4, v1)
                v2 = 42
            }
            if (!(2 > 42)) {
                var v51 uint8 = 103
                _ = v51
            } else {
                v52 := GenS8{F0: ""}
                _ = v52
            }
            if false {
                var v53 string = ""
                var v54 uint8 = 215
                v54 -= 255
                _ = v53
                _ = v54
            } else {
                if true {
                    if len(v3) > 0 {
                        v3[0] = -1
                    }
                    var v55 int = 32767
                    _ = v55
                }
                v56 := GenS9{F0: (v2 << ((0) & 63))}
                _ = v56
            }
        }
    } else {
        if len(v3) > 0 {
            v3[0] = 0
        }
    }
    v57 := GenS10{F0: 2, F1: -1}
    if (2 >= 2) {
        switch 0 {
        case 48:
            if false {
                v58 := GenS11{F0: (v4 + (v4 + 2)), F1: 0}
                if true {
                    v3 = append(v3, 21)
                } else {
                    for fi59 := 0; fi59 < 14; fi59++ {
                        v12 += 1.5
                        _ = fi59
                    }
                }
                var v60 float32 = float32(3.14159)
                _ = v58
                _ = v60
            } else {
                for fi61 := 0; fi61 < 15; fi61++ {
                    fmt.Println(v1, v2, v12)
                    if false {
                        for fi62 := 0; fi62 < 16; fi62++ {
                            var v63 float64 = 3.14159
                            fmt.Println(v14)
                            _ = fi62
                            _ = v63
                        }
                        var v64 int64 = -367
                        _ = v64
                    }
                    var v65 string = "\n\t"
                    _ = fi61
                    _ = v65
                }
            }
        case -2147483648:
            v3 = append(v3, 128)
        default:
            if false {
                switch ("foo bar" + "") {
                case "a":
                    fmt.Println(v1)
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
                _ = len(v3)
            }
        }
        v1 = v1
    } else {
        if (0.0 == 1.5) {
            if (648 >= 42) {
                if true {
                    if (true || false) {
                        if (!((199 > 18446744073709551615) || false)) {
                            v13 = v14
                            for fi66 := 0; fi66 < 12; fi66++ {
                                _ = len(v3)
                                for fi67 := 0; fi67 < 4; fi67++ {
                                    var v68 string = ((("" + ((("hello" + "a") + "hello") + "hello")) + "") + "a")
                                    if ((0.0 - float64((1.5 / (((1e10 * (((3.14159 - (v12 - v14)) - v12) + float32((1.5 + 0.1)))) - (v14 - (v12 * v14))) - v12)))) < 3.14159) {
                                        switch "" {
                                        case "foo bar":
                                            v2 += v2
                                        case "hello":
                                            if false {
                                                v68 = ""
                                                v1 = -128
                                                if (42 > 627) {
                                                    var v69 int16 = v2
                                                    _ = v69
                                                } else {
                                                    switch ((v68 + ((v68 + (v68 + v68)) + v68)) + ((v68 + (("foo bar" + (v68 + "\n\t")) + (("" + v68) + v68))) + "")) {
                                                    case "\n\t":
                                                        fmt.Println(fi67, v13)
                                                    case "a":
                                                        if false {
                                                            if (211 <= 127) {
                                                                v14 = float32(((0.0 + -2.25) - float64(float32(0.1))))
                                                                v3 = append(v3, (fi67 * (fi66 * fi67)))
                                                                v1 = -1
                                                            } else {
                                                                var v70 int16 = (^(v2))
                                                                _ = v70
                                                            }
                                                            var v71 string = (v68 + v68)
                                                            fmt.Println(v2, v13)
                                                            _ = v71
                                                        } else {
                                                            v2 = (v2 - 32767)
                                                            v72 := GenS4{F0: 0}
                                                            _ = v72
                                                        }
                                                    case "hello":
                                                        var v73 int16 = 1
                                                        _ = v73
                                                    default:
                                                        v12 = (v13 * v12)
                                                    }
                                                    var v74 float32 = v14
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    v68 = ("" + (v68 + "foo bar"))
                                                    _ = v74
                                                }
                                            }
                                        case "a":
                                            v2 = v2
                                        default:
                                            fmt.Println(v1)
                                        }
                                    }
                                    v1 = -1
                                    v68 = ("a" + "\n\t")
                                    _ = fi67
                                    _ = v68
                                }
                                var v75 uint = 626
                                _ = fi66
                                _ = v75
                            }
                            v3 = append(v3, -777)
                        }
                        for fi76 := 0; fi76 < 19; fi76++ {
                            if false {
                                v13 += ((float32(3.14159) / v13) + 1.5)
                                v3 = append(v3, 127)
                            } else {
                                for fi77 := 0; fi77 < 4; fi77++ {
                                    v1 = (^(v1))
                                    _ = fi77
                                }
                            }
                            _ = fi76
                        }
                        var v78 uint32 = 65535
                        _ = v78
                    } else {
                        v79 := GenS12{F0: (-(v1))}
                        fmt.Println(v1, v2)
                        _ = v79
                    }
                } else {
                    v3 = append(v3, 120)
                    if false {
                        v80 := GenS13{F0: 909, F1: 73, F2: 953}
                        v14 = float32(float64(v13))
                        _ = v80
                    } else {
                        v13 = (v12 * v14)
                        switch 489 {
                        case -937:
                            if (58 == (v1 >> ((347) & 63))) {
                                v4 = (v4 | 42)
                            }
                        case -2147483648:
                            var v81 int8 = v1
                            _ = v81
                        }
                        v82 := GenS14{F0: (v1 ^ 94), F1: 2147483647}
                        _ = v82
                    }
                }
                var v83 int32 = -1
                for fi84 := 0; fi84 < 19; fi84++ {
                    if true {
                        if (-191 < 668) {
                            var v85 int64 = 517
                            _ = v85
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v86 uint64 = 0
                            _ = v86
                        }
                        fmt.Println(v83, v14)
                        var v87 uint = 2
                        _ = v87
                    } else {
                        v88 := GenS3{F0: -9223372036854775808}
                        _ = v88
                    }
                    _ = fi84
                }
                _ = v83
            } else {
                fmt.Println(v2, v2, v13)
                switch 560 {
                case 610:
                    switch 2147483647 {
                    case -1:
                        if (false && false) {
                            var v89 int64 = 9223372036854775807
                            switch (("\n\t" + ("\n\t" + "")) + "") {
                            case "":
                                var v90 int16 = v2
                                _ = v90
                            case "foo bar":
                                v91 := GenS15{F0: 1, F1: v89, F2: 185}
                                _ = v91
                            case "a":
                                for fi92 := 0; fi92 < 9; fi92++ {
                                    var v93 uint8 = 1
                                    v94 := GenS3{F0: 0}
                                    _ = fi92
                                    _ = v93
                                    _ = v94
                                }
                            default:
                                for fi95 := 0; fi95 < 2; fi95++ {
                                    var v96 int8 = int8(v2)
                                    v12 = float32(1e10)
                                    _ = fi95
                                    _ = v96
                                }
                            }
                            _ = v89
                        }
                    case 2147483647:
                        var v97 uint = 18446744073709551615
                        _ = v97
                    default:
                        fmt.Println(v2, v12)
                    }
                case 4294967295:
                    if false {
                        fmt.Println(v12)
                        v98 := GenS16{F0: (18446744073709551615 < 255)}
                        if false {
                            fmt.Println(v4, v13)
                            v3 = append(v3, 127)
                            var v99 uint8 = 2
                            _ = v99
                        } else {
                            for fi100 := 0; fi100 < 13; fi100++ {
                                var v101 int16 = v2
                                _ = fi100
                                _ = v101
                            }
                        }
                        _ = v98
                    } else {
                        for fi102 := 0; fi102 < 5; fi102++ {
                            if false {
                                var v103 float64 = -2.25
                                if true {
                                    var v104 int32 = 42
                                    _ = v104
                                } else {
                                    var v105 uint32 = 2147483648
                                    v4 = (v4 % ((v4) | 1))
                                    var v106 uint16 = 995
                                    _ = v105
                                    _ = v106
                                }
                                var v107 int32 = 707
                                _ = v103
                                _ = v107
                            } else {
                                v108 := GenS17{F0: 255}
                                _ = v108
                            }
                            fmt.Println(v4, v4)
                            _ = fi102
                        }
                        if (((!false) || true) || false) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            fmt.Println(v1)
                            if len(v3) > 0 {
                                v3[0] = 4294967295
                            }
                        } else {
                            for fi109 := 0; fi109 < 18; fi109++ {
                                v110 := GenS1{F0: 0}
                                var v111 uint = genFn0()
                                v1 += v1
                                _ = fi109
                                _ = v110
                                _ = v111
                            }
                            var v112 int64 = 1
                            var v113 int = 4294967295
                            _ = v112
                            _ = v113
                        }
                    }
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v57
}
