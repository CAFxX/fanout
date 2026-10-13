package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint16; F1 uint32; F2 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32; F1 int16; F2 uint }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS9 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS10 struct { F0 int8; F1 uint16; F2 string }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 int64; F1 int64; F2 bool }
type GenS14 struct { F0 int; F1 uint16 }
type GenS15 struct { F0 int; F1 int16; F2 uint }
type GenS16 struct { F0 uint32 }
type GenS17 struct { F0 uint32; F1 string }
type GenS18 struct { F0 int16; F1 int32 }
func genFn0(p0 uint32) int64 {
    p0 -= p0
    if (!false) {
        v7 := GenS1{F0: 42, F1: (p0 << ((4294967295) & 63))}
        var v8 bool = true
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: -242, F1: 65535}
    v10 := []int{}
    _ = p0
    _ = v10
    _ = v9
    return 0
}

func main() {
    v1 := []uint32{0, 65535, 0}
    var v2 uint64 = 42
    var v3 uint64 = (v2 * 460)
    var v4 uint32 = 65535
    v5 := []int{}
    var v6 int16 = 1000
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v6, v2, v3)
    v4 -= 2
    if false {
        switch "" {
        case "\n\t":
            fmt.Println(v6, v3)
        default:
            if true {
                v3 += v3
                for fi11 := 0; fi11 < 3; fi11++ {
                    v4 = v4
                    v6 = (v6 / ((v6) | 1))
                    _ = fi11
                }
            }
        }
        for fi12 := 0; fi12 < 7; fi12++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v6 = (^(v6))
            v2 += v2
            _ = fi12
        }
        var v13 float64 = ((1.5 - 0.0) - 1e10)
        _ = v13
    }
    var v14 bool = false
    fmt.Println(v14, v3)
    v15 := GenS1{F0: 431, F1: (v4 & 1)}
    var v16 string = "hello"
    for fi17 := 0; fi17 < 9; fi17++ {
        var v18 string = ((((("a" + v16) + "foo bar") + "hello") + (("a" + v16) + (v16 + "hello"))) + v16)
        var v19 string = v18
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 float32 = 1.5
    v21 := GenS1{F0: v6, F1: v4}
    var v22 uint16 = 197
    v23 := GenS1{F0: 32767, F1: (v4 << ((65535) & 63))}
    if v14 {
        if false {
            fmt.Println(v2, v16, v2)
            fmt.Println(v20, v3, v3)
        } else {
            switch 90 {
            case 587:
                if (!v14) {
                    fmt.Println(v20, v3)
                    v16 = "foo bar"
                    fmt.Println(v22, v3)
                }
            case -1:
                if true {
                    v2 -= (v2 % (((v2 + 0)) | 1))
                }
            case 128:
                var v24 uint64 = v3
                _ = v24
            default:
                var v25 uint64 = (^(v3))
                _ = v25
            }
        }
        for fi26 := 0; fi26 < 5; fi26++ {
            fmt.Println(v6)
            switch (fi26 | (fi26 << ((639) & 63))) {
            case 2147483647:
                if v14 {
                    if false {
                        var v27 int64 = 0
                        v28 := GenS1{F0: -1, F1: 42}
                        _ = v27
                        _ = v28
                    } else {
                        for fi29 := 0; fi29 < 17; fi29++ {
                            v14 = (true && v14)
                            switch fi26 {
                            case -447:
                                var v30 uint32 = (v4 - (v4 * uint32(v22)))
                                _ = v30
                            case 128:
                                v4 += (^(v4))
                            case 1:
                                var v31 int16 = 0
                                _ = v31
                            }
                            if len(v5) > 0 {
                                v5[0] = (fi29 + (fi26 ^ -192))
                            }
                            switch fi26 {
                            case 971:
                                var v32 float64 = (1.5 / 1e10)
                                _ = v32
                            case -1:
                                switch v16 {
                                case "hello":
                                    v33 := GenS2{F0: (fi29 >> ((1) & 63)), F1: 753}
                                    _ = v33
                                default:
                                    var v34 bool = ((1.5 + 1.5) >= (((v20 + float32(((3.14159 * 1e10) - (3.14159 / (1e10 * 1.5))))) / ((float32(3.14159) * v20) - v20)) + v20))
                                    _ = v34
                                }
                            }
                            v35 := GenS2{F0: fi29, F1: v3}
                            _ = fi29
                            _ = v35
                        }
                    }
                    for fi36 := 0; fi36 < 4; fi36++ {
                        v16 = v16
                        if true {
                            if v14 {
                                v37 := GenS3{F0: 32767, F1: (v4 << ((18446744073709551615) & 63)), F2: ("hello" + (v16 + v16))}
                                switch fi36 {
                                case 9223372036854775807:
                                    fmt.Println(fi26)
                                case 2147483647:
                                    switch "\n\t" {
                                    case "foo bar":
                                        if v14 {
                                            v1 = append(v1, v4)
                                            fmt.Println(v3, v16)
                                            if (-2147483648 != -1) {
                                                switch "foo bar" {
                                                case "hello":
                                                    v38 := GenS4{F0: 72}
                                                    _ = v38
                                                default:
                                                    v39 := GenS5{F0: v4, F1: v6, F2: 640}
                                                    _ = v39
                                                }
                                                var v40 uint16 = v22
                                                if false {
                                                    if false {
                                                        fmt.Println(v22, v40, fi26)
                                                        v5 = append(v5, 32767)
                                                        switch fi26 {
                                                        case 762:
                                                            if ((!(-2147483648 > 42)) && v14) {
                                                                v41 := GenS5{F0: (v4 * uint32(v22)), F1: 0, F2: 255}
                                                                switch v16 {
                                                                case "\n\t":
                                                                    _ = len(v5)
                                                                }
                                                                fmt.Println(v40, v4, v4)
                                                                _ = v41
                                                            } else {
                                                                v1 = append(v1, (v4 & (v4 >> ((18446744073709551615) & 63))))
                                                                v42 := GenS2{F0: 32767, F1: (v2 | 341)}
                                                                _ = v42
                                                            }
                                                        default:
                                                            switch fi36 {
                                                            case -1:
                                                                var v43 int8 = -74
                                                                _ = v43
                                                            default:
                                                                v44 := GenS6{F0: (v2 % ((v2) | 1)), F1: 2}
                                                                _ = v44
                                                            }
                                                        }
                                                        var v45 bool = (42 == 42)
                                                        v46 := GenS7{F0: v3}
                                                        var v47 float32 = v20
                                                        if v14 {
                                                            if v14 {
                                                                var v48 int = 127
                                                                var v49 uint64 = (v2 - uint64(v4))
                                                                v22 -= v22
                                                                _ = v48
                                                                _ = v49
                                                            }
                                                        }
                                                        _ = v45
                                                        _ = v46
                                                        _ = v47
                                                    } else {
                                                        fmt.Println(v20, v16, v16)
                                                        if v14 {
                                                            if len(v5) > 0 {
                                                                v5[0] = 127
                                                            }
                                                            switch -991 {
                                                            case -2147483648:
                                                                fmt.Println(v22, v4)
                                                            case 32768:
                                                                v50 := GenS5{F0: 2147483648, F1: (v6 / ((v6) | 1)), F2: 42}
                                                                _ = v50
                                                            case -469:
                                                                if v14 {
                                                                    _ = len(v5)
                                                                    v40 = 1
                                                                }
                                                            }
                                                            switch fi36 {
                                                            case -2147483648:
                                                                _ = len(v1)
                                                            }
                                                            v5 = append(v5, fi26)
                                                            var v51 bool = v14
                                                            _ = v51
                                                        }
                                                        v52 := GenS8{F0: v2, F1: 255, F2: 1}
                                                        v5 = append(v5, 4294967295)
                                                        _ = v52
                                                    }
                                                    v2 = v2
                                                    var v53 uint32 = v4
                                                    _ = v53
                                                }
                                                switch fi36 {
                                                case 0:
                                                    if v14 {
                                                        var v54 uint8 = uint8(v3)
                                                        if false {
                                                            if (4294967295 <= 18446744073709551615) {
                                                                v6 = (v6 ^ int16(fi36))
                                                            }
                                                        }
                                                        _ = len(v1)
                                                        _ = v54
                                                    } else {
                                                        var v55 int32 = -1
                                                        var v56 int64 = genFn0((^(v4)))
                                                        v3 -= (v2 * 726)
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                case 9223372036854775807:
                                                    var v57 uint16 = 1
                                                    _ = v57
                                                default:
                                                    v58 := GenS9{F0: uint8(v3), F1: v6, F2: v2}
                                                    _ = v58
                                                }
                                                v40 += v22
                                                fmt.Println(fi36, fi36)
                                                v14 = v14
                                                v59 := GenS1{F0: v6, F1: (v4 % (((v4 / ((v4) | 1))) | 1))}
                                                _ = v40
                                                _ = v59
                                            } else {
                                                var v60 uint = 42
                                                fmt.Println(v60)
                                                var v61 bool = v14
                                                _ = v60
                                                _ = v61
                                            }
                                        } else {
                                            v62 := GenS10{F0: 1, F1: (v22 + v22), F2: ""}
                                            v22 = (^(v22))
                                            v63 := GenS5{F0: v4, F1: (v6 - v6), F2: 4294967295}
                                            _ = v62
                                            _ = v63
                                        }
                                    case "hello":
                                        var v64 float64 = (-2.25 / -2.25)
                                        _ = v64
                                    case "\n\t":
                                        if (!v14) {
                                            v65 := GenS11{F0: (65535 >= v4)}
                                            _ = v65
                                        } else {
                                            var v66 int = fi36
                                            _ = v66
                                        }
                                    default:
                                        if v14 {
                                            v6 = (-(v6))
                                            v67 := GenS5{F0: v4, F1: v6, F2: 18446744073709551615}
                                            _ = v67
                                        } else {
                                            var v68 uint32 = v4
                                            _ = v68
                                        }
                                    }
                                default:
                                    var v69 bool = v14
                                    _ = v69
                                }
                                v4 = (^(v4))
                                _ = v37
                            } else {
                                v3 += (v3 ^ (v3 & (^(v3))))
                            }
                        } else {
                            var v70 uint32 = uint32(v22)
                            _ = len(v1)
                            if v14 {
                                v16 = v16
                                v5 = append(v5, (fi36 - fi26))
                                if v14 {
                                    if len(v5) > 0 {
                                        v5[0] = 2147483647
                                    }
                                }
                            } else {
                                if (-128 > -1) {
                                    v16 = (v16 + ((v16 + ((v16 + ("a" + (v16 + v16))) + "hello")) + v16))
                                    var v71 int = int(v70)
                                    if false {
                                        _ = len(v1)
                                        v71 += fi26
                                    } else {
                                        var v72 uint = 1
                                        var v73 uint8 = 0
                                        fmt.Println(v6, v72, fi26)
                                        _ = v72
                                        _ = v73
                                    }
                                    _ = v71
                                }
                                if v14 {
                                    v5 = append(v5, fi36)
                                }
                            }
                            _ = v70
                        }
                        v4 += 2147483648
                        _ = fi36
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                } else {
                    v20 -= v20
                }
            case -869:
                if ((v2 >> ((255) & 63)) <= v3) {
                    v74 := GenS2{F0: (fi26 & 1), F1: v3}
                    if true {
                        switch v16 {
                        case "\n\t":
                            for fi75 := 0; fi75 < 14; fi75++ {
                                if (2 <= v3) {
                                    if (65535 <= 0) {
                                        v76 := GenS10{F0: -128, F1: (v22 + v22), F2: (v16 + v16)}
                                        _ = v76
                                    } else {
                                        v77 := GenS12{F0: v4}
                                        v22 -= 42
                                        _ = v77
                                    }
                                    v5 = append(v5, (fi75 & fi75))
                                    fmt.Println(fi75, v14)
                                } else {
                                    var v78 uint16 = (^(v22))
                                    _ = v78
                                }
                                fmt.Println(fi26)
                                v1 = append(v1, v4)
                                var v79 uint16 = (v22 | 0)
                                _ = fi75
                                _ = v79
                            }
                        case "foo bar":
                            var v80 uint = 487
                            _ = v80
                        }
                        v81 := GenS10{F0: 67, F1: 0, F2: "hello"}
                        v20 = -2.25
                        _ = v81
                    }
                    v82 := GenS2{F0: fi26, F1: v3}
                    _ = v74
                    _ = v82
                } else {
                    fmt.Println(fi26, v4)
                    if (18 > 297) {
                        var v83 int = fi26
                        if v14 {
                            v2 = (v2 ^ 166)
                        } else {
                            for fi84 := 0; fi84 < 3; fi84++ {
                                if len(v5) > 0 {
                                    v5[0] = (-(fi84))
                                }
                                if v14 {
                                    _ = len(v5)
                                    v14 = (uint8(v3) < 2)
                                    if false {
                                        v85 := GenS6{F0: (v3 | v2), F1: 2}
                                        switch v83 {
                                        case -93:
                                            var v86 float32 = v20
                                            _ = v86
                                        case 1:
                                            var v87 uint16 = (^(v22))
                                            if (v6 < (v6 & v6)) {
                                                switch fi26 {
                                                case 2147483647:
                                                    v88 := GenS2{F0: v83, F1: (v2 / ((v3) | 1))}
                                                    _ = v88
                                                default:
                                                    if v14 {
                                                        if len(v5) > 0 {
                                                            v5[0] = 4294967295
                                                        }
                                                        v2 = (v3 / (((v3 ^ (v3 * (v2 - 4294967295)))) | 1))
                                                        if len(v5) > 0 {
                                                            v5[0] = -2147483648
                                                        }
                                                    }
                                                }
                                            } else {
                                                v14 = (!("" > ""))
                                                if true {
                                                    v87 = (^(v22))
                                                    v89 := GenS13{F0: 1, F1: -9223372036854775808, F2: (!v14)}
                                                    if v14 {
                                                        fmt.Println(v22, v3)
                                                    } else {
                                                        v2 -= 608
                                                    }
                                                    _ = v89
                                                }
                                                if ((!true) && v14) {
                                                    v90 := GenS11{F0: v14}
                                                    _ = v90
                                                }
                                            }
                                            _ = v87
                                        default:
                                            if (fi26 >= (^(v83))) {
                                                switch (v16 + v16) {
                                                case "foo bar":
                                                    if (true && true) {
                                                        v91 := GenS12{F0: v4}
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        _ = v91
                                                    }
                                                case "":
                                                    if len(v1) > 0 {
                                                        v1[0] = 65535
                                                    }
                                                    var v92 int32 = 0
                                                    _ = v92
                                                case "a":
                                                    v16 = "a"
                                                default:
                                                    v93 := GenS14{F0: v83, F1: v22}
                                                    _ = v93
                                                }
                                                v22 = 501
                                            }
                                        }
                                        var v94 string = v16
                                        _ = v85
                                        _ = v94
                                    }
                                } else {
                                    v95 := GenS2{F0: 18, F1: (v2 / (((v2 / ((0) | 1))) | 1))}
                                    if (!false) {
                                        if v14 {
                                            var v96 float64 = (float64(v20) + float64(v20))
                                            v5 = append(v5, v83)
                                            var v97 float32 = v20
                                            _ = v96
                                            _ = v97
                                        } else {
                                            v98 := GenS15{F0: 4294967295, F1: (v6 / ((1) | 1)), F2: 18446744073709551615}
                                            v2 += (v2 % ((721) | 1))
                                            v99 := GenS13{F0: -358, F1: genFn0((v4 - (v4 / (((v4 >> ((4294967295) & 63))) | 1)))), F2: v14}
                                            _ = v98
                                            _ = v99
                                        }
                                        v3 -= v2
                                    } else {
                                        if v14 {
                                            fmt.Println(v2)
                                        }
                                        if v14 {
                                            var v100 int32 = 42
                                            _ = v100
                                        }
                                        v20 = float32(1.5)
                                        v14 = ("" <= ((v16 + v16) + ("\n\t" + v16)))
                                        fmt.Println(fi26, fi26, v14)
                                    }
                                    _ = v95
                                }
                                var v101 uint16 = 0
                                _ = fi84
                                _ = v101
                            }
                        }
                        fmt.Println(v3, v4)
                        _ = v83
                    } else {
                        v16 = v16
                        if (255 > 491) {
                            v102 := GenS13{F0: -586, F1: 9223372036854775807, F2: false}
                            _ = v102
                        } else {
                            var v103 int64 = -9223372036854775808
                            _ = v103
                        }
                    }
                    if (!v14) {
                        if v14 {
                            fmt.Println(v6)
                            if (869 >= 42) {
                                v104 := GenS16{F0: (v4 & v4)}
                                v105 := GenS17{F0: 2147483648, F1: v16}
                                var v106 string = "hello"
                                _ = v104
                                _ = v105
                                _ = v106
                            } else {
                                switch (("hello" + "foo bar") + (v16 + "foo bar")) {
                                case "foo bar":
                                    v22 = (v22 >> ((0) & 63))
                                case "hello":
                                    _ = len(v1)
                                case "":
                                    v22 = (v22 / (((v22 >> ((65535) & 63))) | 1))
                                }
                                fmt.Println(v3)
                                v2 += v3
                            }
                            var v107 int16 = 0
                            _ = v107
                        }
                    } else {
                        fmt.Println(v2)
                        for fi108 := 0; fi108 < 17; fi108++ {
                            var v109 uint64 = (^(v3))
                            if ((^(v6)) == v6) {
                                v110 := GenS18{F0: v6, F1: 665}
                                _ = v110
                            }
                            v6 -= 0
                            _ = fi108
                            _ = v109
                        }
                        fmt.Println(fi26)
                    }
                }
            default:
                if (false && ((^(v22)) != v22)) {
                    for fi111 := 0; fi111 < 13; fi111++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi111
                    }
                    if len(v1) > 0 {
                        v1[0] = 42
                    }
                    for fi112 := 0; fi112 < 20; fi112++ {
                        var v113 uint32 = (v4 - (v4 + v4))
                        _ = fi112
                        _ = v113
                    }
                } else {
                    var v114 bool = v14
                    fmt.Println(v20, v16)
                    _ = v114
                }
            }
            for fi115 := 0; fi115 < 13; fi115++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v4 -= (v4 << ((65535) & 63))
                _ = fi115
            }
            _ = fi26
        }
        v116 := GenS17{F0: uint32(v3), F1: v16}
        _ = v116
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
