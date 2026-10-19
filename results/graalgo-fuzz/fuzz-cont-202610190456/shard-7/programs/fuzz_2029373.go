package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int; F1 uint32; F2 uint16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int64; F1 int64; F2 string }
type GenS9 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS10 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS11 struct { F0 bool; F1 uint32; F2 uint32 }
type GenS12 struct { F0 int64; F1 int8; F2 int32 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint16; F1 uint16 }
type GenS16 struct { F0 int16; F1 uint64 }
type GenS17 struct { F0 int64; F1 int; F2 int32 }
type GenS18 struct { F0 int }
type GenS19 struct { F0 uint; F1 uint64 }
type GenS20 struct { F0 bool; F1 int32 }
type GenS21 struct { F0 bool; F1 int16; F2 int32 }
type GenS22 struct { F0 int16; F1 int; F2 int32 }
type GenS23 struct { F0 string; F1 int }
type GenS24 struct { F0 int8 }
type GenS25 struct { F0 int64; F1 int32 }
type GenS26 struct { F0 uint16 }
type GenS27 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS28 struct { F0 int64; F1 int16 }
type GenS29 struct { F0 uint64; F1 int8; F2 int8 }
func genFn0() bool {
    for fi5 := 0; fi5 < 7; fi5++ {
        fmt.Println(fi5)
        _ = fi5
    }
    v6 := GenS1{F0: 255}
    var v7 uint32 = 2147483647
    fmt.Println(v7)
    for fi8 := 0; fi8 < 5; fi8++ {
        switch (fi8 % (((fi8 & fi8)) | 1)) {
        case 127:
            for fi9 := 0; fi9 < 14; fi9++ {
                switch "a" {
                case "a":
                    switch fi8 {
                    case 32768:
                        v10 := GenS1{F0: 669}
                        _ = v10
                    default:
                        var v11 int = fi8
                        _ = v11
                    }
                }
                var v12 float32 = float32(3.14159)
                if (9223372036854775807 >= -9223372036854775808) {
                    v13 := []int{127, (fi9 << ((255) & 63))}
                    _ = v13
                } else {
                    var v14 int32 = -2147483648
                    v7 = 2147483648
                    v15 := GenS1{F0: 255}
                    switch ("a" + ("a" + (("foo bar" + "a") + "hello"))) {
                    case "\n\t":
                        v16 := []int{fi8, fi9}
                        _ = v16
                    case "a":
                        if false {
                            v17 := GenS2{F0: 9223372036854775807, F1: 42, F2: -95}
                            _ = v17
                        }
                    default:
                        v18 := GenS1{F0: 42}
                        _ = v18
                    }
                    _ = v14
                    _ = v15
                }
                _ = fi9
                _ = v12
            }
        case -654:
            var v19 int8 = -17
            _ = v19
        default:
            v20 := GenS2{F0: 0, F1: 0, F2: 100}
            _ = v20
        }
        v7 = v7
        _ = fi8
    }
    _ = v6
    _ = v7
    return true
}

func main() {
    v1 := []string{""}
    v2 := []int{42, 747, 0}
    var v3 uint16 = 294
    v4 := []int{}
    var v21 float64 = 1e10
    if false {
        _ = len(v2)
        v22 := GenS1{F0: 363}
        for fi23 := 0; fi23 < 11; fi23++ {
            if (1e10 <= (float64(float32(v21)) + (((3.14159 - v21) + (3.14159 / 0.1)) / (v21 * v21)))) {
                for fi24 := 0; fi24 < 3; fi24++ {
                    fmt.Println(v21)
                    if true {
                        if (42 != 0) {
                            v3 = 573
                            if false {
                                var v25 int8 = -1
                                _ = v25
                            } else {
                                var v26 string = "foo bar"
                                _ = len(v4)
                                _ = v26
                            }
                        } else {
                            switch fi24 {
                            case 612:
                                v3 += v3
                            case 127:
                                fmt.Println(v3, fi23)
                            default:
                                var v27 int32 = 0
                                _ = v27
                            }
                            v3 = v3
                            v28 := GenS3{F0: 1}
                            _ = v28
                        }
                        switch "a" {
                        case "foo bar":
                            if false {
                                if true {
                                    v3 = (^(v3))
                                    if len(v1) > 0 {
                                        v1[0] = "a"
                                    }
                                    v29 := GenS1{F0: (v3 >> ((255) & 63))}
                                    _ = v29
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = "\n\t"
                                    }
                                    var v30 int16 = 1
                                    if (127 == 127) {
                                        switch "hello" {
                                        case "foo bar":
                                            fmt.Println(v3)
                                        case "\n\t":
                                            switch fi24 {
                                            case 591:
                                                var v31 int8 = -128
                                                _ = v31
                                            case -1:
                                                switch (fi23 * fi23) {
                                                case 128:
                                                    var v32 uint16 = (v3 << ((1) & 63))
                                                    _ = v32
                                                default:
                                                    if len(v1) > 0 {
                                                        v1[0] = "foo bar"
                                                    }
                                                }
                                            }
                                        }
                                        _ = len(v1)
                                        var v33 float64 = v21
                                        var v34 float32 = 3.14159
                                        var v35 uint8 = 255
                                        v4 = append(v4, -397)
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    } else {
                                        var v36 uint32 = 4294967295
                                        v30 += int16(v36)
                                        _ = v36
                                    }
                                    _ = v30
                                }
                                v21 = v21
                                if false {
                                    v37 := GenS4{F0: (!(false && true)), F1: (v3 ^ v3)}
                                    _ = v37
                                }
                            }
                        case "hello":
                            v4 = append(v4, -936)
                        case "a":
                            var v38 int16 = 32767
                            _ = v38
                        default:
                            var v39 int8 = -128
                            _ = v39
                        }
                        fmt.Println(fi23, fi24)
                    }
                    _ = fi24
                }
                var v40 string = "a"
                for fi41 := 0; fi41 < 13; fi41++ {
                    if len(v2) > 0 {
                        v2[0] = fi23
                    }
                    switch -2147483648 {
                    case 666:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case -235:
                        if (0 == 32767) {
                            var v42 int16 = 1000
                            fmt.Println(fi41)
                            if false {
                                switch v40 {
                                case "a":
                                    v43 := GenS3{F0: 2147483648}
                                    _ = v43
                                case "":
                                    var v44 int64 = 466
                                    _ = v44
                                }
                                if false {
                                    if (42 >= 0) {
                                        v3 = (^(v3))
                                        var v45 string = "hello"
                                        _ = v45
                                    }
                                    if ((false && false) || ((v21 + 0.0) <= (v21 + float64((0.0 / 1e10))))) {
                                        var v46 float64 = 3.14159
                                        _ = v46
                                    } else {
                                        v3 = (v3 << ((4294967295) & 63))
                                        v47 := GenS1{F0: v3}
                                        _ = v47
                                    }
                                    if (18446744073709551615 > 1) {
                                        v48 := GenS3{F0: 2147483648}
                                        if true {
                                            var v49 int64 = 42
                                            switch 2147483647 {
                                            case 32767:
                                                switch 32767 {
                                                case 4294967295:
                                                    fmt.Println(v40, v40, v42)
                                                case 1:
                                                    _ = len(v2)
                                                case -1:
                                                    v49 += (v49 * v49)
                                                default:
                                                    switch fi23 {
                                                    case 9223372036854775807:
                                                        switch ((((("foo bar" + "foo bar") + (v40 + "foo bar")) + v40) + ("hello" + ("a" + v40))) + v40) {
                                                        case "a":
                                                            fmt.Println(fi41, v3)
                                                        }
                                                    case 4294967295:
                                                        if false {
                                                            v42 += v42
                                                            v40 = v40
                                                            if true {
                                                                v40 = ("foo bar" + v40)
                                                                if (true && false) {
                                                                    fmt.Println(v3, v21, v49)
                                                                    if false {
                                                                        _ = len(v4)
                                                                        var v50 string = "\n\t"
                                                                        _ = v50
                                                                    }
                                                                    fmt.Println(fi41)
                                                                    v1 = append(v1, v40)
                                                                }
                                                            }
                                                        }
                                                    case 42:
                                                        var v51 int16 = 1000
                                                        v52 := GenS3{F0: 503}
                                                        _ = v51
                                                        _ = v52
                                                    default:
                                                        if false {
                                                            v40 = v40
                                                            if false {
                                                                var v53 int16 = v42
                                                                var v54 uint8 = 1
                                                                _ = v53
                                                                _ = v54
                                                            }
                                                            v3 = v3
                                                            var v55 float32 = (3.14159 + -2.25)
                                                            if (!false) {
                                                                if (0 == 685) {
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    fmt.Println(fi41, v42)
                                                                    switch fi41 {
                                                                    case 9223372036854775807:
                                                                        switch (fi41 ^ fi23) {
                                                                        case 32767:
                                                                            v56 := GenS5{F0: 127}
                                                                            switch "hello" {
                                                                            case "\n\t":
                                                                                if true {
                                                                                    switch ((((v40 + v40) + v40) + v40) + v40) {
                                                                                    case "foo bar":
                                                                                        if false {
                                                                                            var v57 uint32 = 2147483648
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = (fi41 & int(v49))
                                                                                            }
                                                                                            _ = v57
                                                                                        } else {
                                                                                            v1 = append(v1, v40)
                                                                                            var v58 int16 = (-(v42))
                                                                                            var v59 int16 = v58
                                                                                            _ = v58
                                                                                            _ = v59
                                                                                        }
                                                                                    case "\n\t":
                                                                                        v21 = (1e10 / v21)
                                                                                    case "hello":
                                                                                        if false {
                                                                                            var v60 int64 = (v49 >> ((18446744073709551615) & 63))
                                                                                            _ = v60
                                                                                        } else {
                                                                                            var v61 float64 = 1e10
                                                                                            switch fi41 {
                                                                                            case 1:
                                                                                                var v62 uint = 65535
                                                                                                _ = v62
                                                                                            default:
                                                                                                fmt.Println(v55, v21, v55)
                                                                                                v63 := GenS2{F0: int64(v42), F1: 181, F2: 100}
                                                                                                v64 := GenS6{F0: (fi41 | fi41), F1: 2147483648, F2: 255}
                                                                                                v40 = "\n\t"
                                                                                                if false {
                                                                                                    var v65 uint16 = (v3 ^ (v3 * (v3 & 42)))
                                                                                                    v66 := GenS7{F0: 1}
                                                                                                    v1 = append(v1, v40)
                                                                                                    _ = v65
                                                                                                    _ = v66
                                                                                                }
                                                                                                _ = v63
                                                                                                _ = v64
                                                                                            }
                                                                                            var v67 int32 = 873
                                                                                            _ = v61
                                                                                            _ = v67
                                                                                        }
                                                                                    }
                                                                                    var v68 int = fi23
                                                                                    v55 = -2.25
                                                                                    _ = v68
                                                                                } else {
                                                                                    var v69 uint64 = 62
                                                                                    v70 := GenS3{F0: 294}
                                                                                    _ = v69
                                                                                    _ = v70
                                                                                }
                                                                            case "":
                                                                                if true {
                                                                                    var v71 uint32 = 4294967295
                                                                                    _ = v71
                                                                                } else {
                                                                                    v40 = v40
                                                                                }
                                                                            default:
                                                                                var v72 uint16 = v3
                                                                                _ = v72
                                                                            }
                                                                            v21 += (((1.5 * v21) + v21) - v21)
                                                                            fmt.Println(fi23, v3)
                                                                            if (1 < 127) {
                                                                                fmt.Println(v3)
                                                                                v49 = 967
                                                                                var v73 uint16 = (v3 / (((v3 / ((0) | 1))) | 1))
                                                                                _ = v73
                                                                            } else {
                                                                                v74 := GenS7{F0: 18446744073709551615}
                                                                                v40 = "hello"
                                                                                var v75 int32 = 0
                                                                                var v76 uint16 = (v3 | (v3 | v3))
                                                                                if (128 >= 42) {
                                                                                    if false {
                                                                                        var v77 uint8 = 152
                                                                                        v78 := GenS7{F0: 1}
                                                                                        _ = v77
                                                                                        _ = v78
                                                                                    }
                                                                                    v42 = -1
                                                                                    v42 = (^(v42))
                                                                                    v75 -= (-(v75))
                                                                                } else {
                                                                                    v55 = v55
                                                                                    switch fi41 {
                                                                                    case 832:
                                                                                        v40 = v40
                                                                                    default:
                                                                                        var v79 float64 = (float64(v55) - v21)
                                                                                        _ = v79
                                                                                    }
                                                                                }
                                                                                var v80 int32 = (v75 + (v75 - (^(v75))))
                                                                                v81 := GenS5{F0: 128}
                                                                                v42 = (v42 << ((65535) & 63))
                                                                                _ = v74
                                                                                _ = v75
                                                                                _ = v76
                                                                                _ = v80
                                                                                _ = v81
                                                                            }
                                                                            _ = v56
                                                                        case 572:
                                                                            var v82 float64 = (v21 / 0.1)
                                                                            _ = v82
                                                                        default:
                                                                            if true {
                                                                                var v83 uint64 = 204
                                                                                _ = v83
                                                                            } else {
                                                                                v2 = append(v2, (fi23 * (-(fi23))))
                                                                                v21 = v21
                                                                                var v84 uint32 = 504
                                                                                _ = v84
                                                                            }
                                                                        }
                                                                    default:
                                                                        v40 = v40
                                                                    }
                                                                } else {
                                                                    if (1 <= 2147483648) {
                                                                        v42 += 32767
                                                                        var v85 int64 = -917
                                                                        v86 := GenS1{F0: v3}
                                                                        _ = v85
                                                                        _ = v86
                                                                    }
                                                                }
                                                            } else {
                                                                switch "" {
                                                                case "a":
                                                                    var v87 uint8 = 128
                                                                    _ = v87
                                                                case "hello":
                                                                    var v88 int16 = v42
                                                                    _ = v88
                                                                default:
                                                                    if len(v4) > 0 {
                                                                        v4[0] = 42
                                                                    }
                                                                }
                                                                if true {
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    var v89 string = (v40 + (v40 + (v40 + (v40 + "hello"))))
                                                                    _ = v89
                                                                } else {
                                                                    v1 = append(v1, ("foo bar" + (v40 + "foo bar")))
                                                                    fmt.Println(fi41, fi41)
                                                                }
                                                                if (v49 >= v49) {
                                                                    if true {
                                                                        v3 = v3
                                                                        fmt.Println(v3, fi23)
                                                                    } else {
                                                                        v90 := GenS4{F0: ((v42 / ((v42) | 1)) == (v42 ^ -1)), F1: (v3 ^ (v3 / ((42) | 1)))}
                                                                        v91 := GenS8{F0: v49, F1: v49, F2: ("\n\t" + v40)}
                                                                        _ = v90
                                                                        _ = v91
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        if (true || true) {
                                                                            var v92 uint8 = 197
                                                                            var v93 uint = 1
                                                                            _ = v92
                                                                            _ = v93
                                                                        } else {
                                                                            v42 -= int16(fi41)
                                                                            var v94 bool = false
                                                                            _ = v94
                                                                        }
                                                                        v1 = append(v1, "hello")
                                                                        v49 = (v49 & v49)
                                                                    } else {
                                                                        v3 = 586
                                                                    }
                                                                }
                                                            }
                                                            _ = v55
                                                        }
                                                    }
                                                }
                                            case 0:
                                                var v95 uint8 = 2
                                                _ = v95
                                            case 4294967295:
                                                var v96 int32 = 2147483647
                                                _ = v96
                                            }
                                            switch v40 {
                                            case "foo bar":
                                                switch v40 {
                                                case "":
                                                    v21 -= 0.1
                                                case "a":
                                                    v49 = (v49 % ((int64(v42)) | 1))
                                                    v97 := GenS9{F0: (v3 - (v3 - v3)), F1: 2147483647, F2: -128}
                                                    if true {
                                                        fmt.Println(v40, fi41, v3)
                                                        if true {
                                                            fmt.Println(v40, v49, v42)
                                                            fmt.Println(v21)
                                                        } else {
                                                            var v98 bool = (((!(!(false && (1e10 <= v21)))) || true) && (!((v49 ^ (v49 ^ (^(v49)))) <= 774)))
                                                            v99 := GenS10{F0: 133, F1: 762, F2: 0}
                                                            _ = v98
                                                            _ = v99
                                                        }
                                                        switch (-(fi23)) {
                                                        case 4294967295:
                                                            var v100 bool = true
                                                            _ = v100
                                                        }
                                                    } else {
                                                        var v101 uint = 42
                                                        _ = v101
                                                    }
                                                    _ = v97
                                                }
                                                v4 = append(v4, 127)
                                                var v102 uint = 839
                                                if false {
                                                    v103 := GenS11{F0: false, F1: 2, F2: 530}
                                                    v40 = ("hello" + "")
                                                    fmt.Println(v21)
                                                    _ = v103
                                                }
                                                _ = v102
                                            case "":
                                                v42 -= (v42 << ((181) & 63))
                                            case "hello":
                                                if false {
                                                    v104 := GenS12{F0: (^(v49)), F1: -128, F2: -597}
                                                    v42 = (v42 % ((v42) | 1))
                                                    _ = v104
                                                } else {
                                                    var v105 uint8 = 119
                                                    _ = v105
                                                }
                                            }
                                            _ = v49
                                        } else {
                                            fmt.Println(v3, fi23)
                                            switch v40 {
                                            case "foo bar":
                                                var v106 uint16 = v3
                                                _ = v106
                                            case "hello":
                                                if true {
                                                    if true {
                                                        var v107 float32 = -2.25
                                                        if false {
                                                            v42 -= -307
                                                            fmt.Println(fi23, fi23, v40)
                                                        } else {
                                                            if false {
                                                                v108 := GenS13{F0: -1}
                                                                switch (v40 + v40) {
                                                                case "":
                                                                    if (true || true) {
                                                                        fmt.Println(v40, fi23, v107)
                                                                        if true {
                                                                            if false {
                                                                                v21 += (v21 - 0.1)
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                            }
                                                                            var v109 int8 = 1
                                                                            var v110 float64 = float64(0.1)
                                                                            _ = v109
                                                                            _ = v110
                                                                        } else {
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                _ = v108
                                                            }
                                                        }
                                                        v21 -= v21
                                                        _ = v107
                                                    }
                                                }
                                            default:
                                                var v111 uint64 = 563
                                                _ = v111
                                            }
                                            v3 += v3
                                        }
                                        fmt.Println(fi23, fi41, v3)
                                        _ = v48
                                    } else {
                                        var v112 bool = (4294967295 != 537)
                                        if len(v4) > 0 {
                                            v4[0] = (^(fi23))
                                        }
                                        _ = v112
                                    }
                                } else {
                                    var v113 uint = 42
                                    _ = v113
                                }
                                if true {
                                    var v114 int16 = v42
                                    _ = v114
                                }
                                v115 := GenS14{F0: (v42 - v42)}
                                var v116 float64 = v21
                                _ = v115
                                _ = v116
                            } else {
                                if (!true) {
                                    var v117 uint = 4294967295
                                    if len(v2) > 0 {
                                        v2[0] = 528
                                    }
                                    _ = v117
                                } else {
                                    fmt.Println(fi41, fi23)
                                }
                                switch v40 {
                                case "a":
                                    if (4294967295 >= 298) {
                                        v4 = append(v4, fi23)
                                        if (!(-684 == -2147483648)) {
                                            v40 = ("\n\t" + "")
                                            v3 = (v3 + v3)
                                            switch (fi23 >> ((2) & 63)) {
                                            case 523:
                                                switch v40 {
                                                case "":
                                                    fmt.Println(v21, v21, fi23)
                                                case "a":
                                                    var v118 int32 = 2147483647
                                                    _ = v118
                                                }
                                            default:
                                                if ("a" == v40) {
                                                    v119 := GenS5{F0: 255}
                                                    _ = v119
                                                }
                                            }
                                        }
                                        v3 -= (v3 / (((^(v3))) | 1))
                                    } else {
                                        var v120 int = (fi23 / ((fi23) | 1))
                                        _ = v120
                                    }
                                case "\n\t":
                                    v121 := GenS15{F0: v3, F1: v3}
                                    _ = v121
                                case "hello":
                                    if false {
                                        if len(v4) > 0 {
                                            v4[0] = (-(fi41))
                                        }
                                    }
                                }
                                v21 = v21
                            }
                            _ = v42
                        }
                    default:
                        var v122 bool = true
                        var v123 uint8 = 90
                        _ = v122
                        _ = v123
                    }
                    _ = fi41
                }
                _ = v40
            } else {
                fmt.Println(v21, v21)
            }
            for fi124 := 0; fi124 < 18; fi124++ {
                v3 = (v3 / (((v3 - v3)) | 1))
                _ = fi124
            }
            if true {
                for fi125 := 0; fi125 < 18; fi125++ {
                    switch (fi23 << ((697) & 63)) {
                    case -878:
                        switch 127 {
                        case 4294967295:
                            v21 -= (0.1 + v21)
                        }
                        v126 := GenS15{F0: (v3 ^ v3), F1: (v3 * (v3 + 42))}
                        if (!(!(true || (0 != int64(fi23))))) {
                            switch ("hello" + "\n\t") {
                            case "a":
                                v127 := GenS10{F0: 525, F1: 1, F2: 2}
                                _ = v127
                            case "\n\t":
                                switch "a" {
                                case "a":
                                    v21 -= v21
                                }
                            case "":
                                v4 = append(v4, fi23)
                            default:
                                v3 = v3
                            }
                            if (127 < 127) {
                                switch "hello" {
                                case "hello":
                                    var v128 uint64 = 4294967295
                                    _ = v128
                                case "":
                                    var v129 float32 = (float32(v21) * 0.1)
                                    _ = v129
                                case "\n\t":
                                    var v130 uint32 = 42
                                    _ = v130
                                default:
                                    v3 = v3
                                }
                                switch (fi125 ^ (fi125 - fi125)) {
                                case -511:
                                    var v131 uint = 708
                                    if len(v2) > 0 {
                                        v2[0] = (fi125 ^ (fi125 % (((fi125 + (fi23 >> (((v131 | 4294967295)) & 63)))) | 1)))
                                    }
                                    _ = v131
                                default:
                                    switch fi23 {
                                    case 86:
                                        var v132 uint16 = v3
                                        _ = v132
                                    case -208:
                                        var v133 bool = ((!(true && false)) && ((21 < 74) || false))
                                        var v134 int64 = -465
                                        v3 = v3
                                        _ = v133
                                        _ = v134
                                    }
                                }
                                if len(v4) > 0 {
                                    v4[0] = (-(fi125))
                                }
                            } else {
                                v3 = 32767
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            if true {
                                if false {
                                    _ = len(v1)
                                    if true {
                                        var v135 uint = 786
                                        v21 -= v21
                                        v3 = (v3 * (v3 & v3))
                                        _ = v135
                                    } else {
                                        v136 := GenS11{F0: false, F1: 0, F2: 733}
                                        v21 += v21
                                        _ = v136
                                    }
                                    switch "" {
                                    case "a":
                                        v137 := GenS11{F0: true, F1: 0, F2: 291}
                                        _ = v137
                                    }
                                } else {
                                    var v138 string = ("" + "a")
                                    v3 += 2
                                    _ = v138
                                }
                                var v139 float32 = ((1.5 + float32(v21)) / -2.25)
                                _ = v139
                            }
                        }
                        _ = v126
                    case 9223372036854775807:
                        var v140 float32 = 3.14159
                        _ = v140
                    case -9223372036854775808:
                        v141 := GenS16{F0: 184, F1: 2}
                        _ = v141
                    default:
                        if true {
                            if false {
                                v21 -= 1e10
                            } else {
                                var v142 float64 = v21
                                var v143 int8 = 1
                                if true {
                                    var v144 string = "foo bar"
                                    _ = len(v2)
                                    switch fi23 {
                                    case 9223372036854775807:
                                        if (!(false || false)) {
                                            if true {
                                                if (-112 > 9223372036854775807) {
                                                    switch v144 {
                                                    case "hello":
                                                        fmt.Println(fi23)
                                                    case "":
                                                        fmt.Println(v21, v144)
                                                    case "\n\t":
                                                        v3 += (v3 - (^(v3)))
                                                    }
                                                }
                                                if true {
                                                    if true {
                                                        v145 := GenS2{F0: 0, F1: v3, F2: -52}
                                                        fmt.Println(v3, v3)
                                                        _ = v145
                                                    }
                                                }
                                            }
                                            v142 -= v142
                                            switch -414 {
                                            case 42:
                                                var v146 string = ("\n\t" + "a")
                                                _ = v146
                                            }
                                        }
                                        v147 := GenS10{F0: 0, F1: v3, F2: 2}
                                        v148 := GenS5{F0: 127}
                                        _ = v147
                                        _ = v148
                                    }
                                    _ = v144
                                }
                                _ = v142
                                _ = v143
                            }
                        }
                    }
                    _ = fi125
                }
                if ((192 >= 2147483647) && true) {
                    v149 := GenS17{F0: 42, F1: (fi23 / (((fi23 >> ((418) & 63))) | 1)), F2: -606}
                    if (-1 <= 1) {
                        if true {
                            v1 = append(v1, "hello")
                            v150 := GenS7{F0: 0}
                            _ = v150
                        } else {
                            v3 += v3
                            v3 = v3
                        }
                        switch (fi23 >> ((329) & 63)) {
                        case 9:
                            v3 += v3
                        case 70:
                            v151 := GenS18{F0: fi23}
                            _ = v151
                        default:
                            v152 := GenS17{F0: 9223372036854775807, F1: fi23, F2: -1}
                            _ = v152
                        }
                        if (-989 <= 31) {
                            if (!false) {
                                if false {
                                    if (fi23 >= (fi23 | fi23)) {
                                        v3 += (v3 * v3)
                                    } else {
                                        v153 := GenS4{F0: true, F1: v3}
                                        _ = v153
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = "a"
                                    }
                                } else {
                                    switch "" {
                                    case "\n\t":
                                        var v154 string = "a"
                                        _ = v154
                                    case "hello":
                                        var v155 uint64 = 4294967295
                                        _ = v155
                                    case "foo bar":
                                        v3 -= v3
                                    }
                                    switch (fi23 ^ fi23) {
                                    case -744:
                                        for fi156 := 0; fi156 < 20; fi156++ {
                                            fmt.Println(fi156)
                                            var v157 uint64 = 525
                                            _ = fi156
                                            _ = v157
                                        }
                                    case -499:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    default:
                                        fmt.Println(v21, v21)
                                    }
                                }
                            }
                            fmt.Println(fi23, v21, v3)
                        } else {
                            v158 := GenS3{F0: 2147483647}
                            if true {
                                for fi159 := 0; fi159 < 10; fi159++ {
                                    var v160 float32 = 1.5
                                    _ = fi159
                                    _ = v160
                                }
                                switch fi23 {
                                case 552:
                                    v21 -= v21
                                default:
                                    for fi161 := 0; fi161 < 18; fi161++ {
                                        v162 := GenS11{F0: true, F1: 2147483648, F2: 322}
                                        v3 = 244
                                        v21 += float64(0.1)
                                        _ = fi161
                                        _ = v162
                                    }
                                }
                            }
                            v21 = (((v21 - 1e10) * ((3.14159 - (((-2.25 / 0.1) + float64(-2.25)) * v21)) / (((v21 - 0.1) - ((((v21 + v21) / v21) - (1e10 * v21)) + v21)) / 3.14159))) / (1.5 - v21))
                            _ = v158
                        }
                    } else {
                        var v163 uint32 = 42
                        fmt.Println(v3, v21)
                        v163 = (v163 - (^(v163)))
                        _ = v163
                    }
                    _ = v149
                }
            } else {
                fmt.Println(v3)
                if len(v2) > 0 {
                    v2[0] = fi23
                }
                if len(v4) > 0 {
                    v4[0] = (^(fi23))
                }
            }
            _ = fi23
        }
        _ = v22
    } else {
        v3 -= (^(v3))
        v3 -= (^(v3))
        if true {
            v21 = -2.25
        }
    }
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    for fi164 := 0; fi164 < 12; fi164++ {
        v3 += (v3 + v3)
        fmt.Println(v21, v3)
        _ = fi164
    }
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    switch (("hello" + "foo bar") + "hello") {
    case "a":
        if ((740 <= 2) && false) {
            if (!(!((1 <= 0) || (!true)))) {
                fmt.Println(v21)
                fmt.Println(v21)
            } else {
                var v165 int64 = -9223372036854775808
                _ = v165
            }
            fmt.Println(v3)
        }
    }
    v21 = v21
    for fi166 := 0; fi166 < 9; fi166++ {
        var v167 int8 = 27
        v4 = append(v4, (^(fi166)))
        _ = fi166
        _ = v167
    }
    v3 = 533
    v21 -= v21
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v21, v21)
    switch "hello" {
    case "\n\t":
        for fi168 := 0; fi168 < 20; fi168++ {
            if false {
                if false {
                    if (1 == 807) {
                        var v169 string = "\n\t"
                        var v170 float32 = float32(v21)
                        _ = v169
                        _ = v170
                    }
                    v21 = v21
                } else {
                    v3 = v3
                }
                var v171 float32 = 1e10
                fmt.Println(v21, v171, v21)
                _ = v171
            } else {
                fmt.Println(v21, v3, v21)
                for fi172 := 0; fi172 < 15; fi172++ {
                    v173 := GenS19{F0: 4294967295, F1: 4294967295}
                    switch "" {
                    case "":
                        fmt.Println(fi172, v3, fi172)
                    case "hello":
                        switch (fi168 % (((-(fi168))) | 1)) {
                        case 42:
                            var v174 uint64 = 2
                            _ = v174
                        case -474:
                            if false {
                                v175 := GenS19{F0: 4294967295, F1: 42}
                                if ((!false) && (!false)) {
                                    v176 := GenS20{F0: false, F1: 873}
                                    var v177 int8 = -1
                                    _ = v176
                                    _ = v177
                                }
                                v4 = append(v4, (fi172 >> ((uint(v3)) & 63)))
                                _ = v175
                            } else {
                                fmt.Println(fi172, fi168)
                                fmt.Println(fi172)
                                v178 := GenS21{F0: false, F1: -253, F2: 2147483647}
                                var v179 uint32 = 952
                                _ = v178
                                _ = v179
                            }
                            var v180 int32 = -649
                            if ((-(fi172)) == fi172) {
                                var v181 float64 = (v21 + v21)
                                v21 = v21
                                _ = len(v2)
                                _ = v181
                            } else {
                                if false {
                                    switch fi168 {
                                    case -1:
                                        v182 := GenS22{F0: -32768, F1: (fi172 | fi168), F2: v180}
                                        v183 := GenS12{F0: int64(fi168), F1: 102, F2: v180}
                                        v184 := GenS23{F0: (("\n\t" + "\n\t") + ("\n\t" + ((("hello" + ("a" + "a")) + "a") + "hello"))), F1: fi168}
                                        _ = v182
                                        _ = v183
                                        _ = v184
                                    default:
                                        var v185 uint64 = 4294967295
                                        _ = v185
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = (fi172 * -2147483648)
                                    }
                                } else {
                                    var v186 float32 = ((0.1 * ((((((1e10 / -2.25) - float32(v21)) / (-2.25 - 0.0)) * 0.1) / 1e10) + 1e10)) / 0.1)
                                    _ = v186
                                }
                                fmt.Println(fi168)
                                v180 = v180
                            }
                            if (false || true) {
                                v187 := GenS15{F0: (v3 >> ((1) & 63)), F1: (v3 / (((v3 ^ 0)) | 1))}
                                v3 += 32767
                                var v188 float64 = ((v21 * v21) / v21)
                                _ = v187
                                _ = v188
                            } else {
                                var v189 int64 = -1
                                _ = v189
                            }
                            if (true && ((v180 % ((v180) | 1)) < v180)) {
                                v3 += v3
                                var v190 int = fi172
                                v21 += v21
                                _ = v190
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v180 += v180
                                switch fi172 {
                                case 1:
                                    var v191 int64 = 635
                                    if true {
                                        var v192 string = ""
                                        _ = v192
                                    }
                                    if ((!true) && false) {
                                        v3 = (v3 + (v3 << ((255) & 63)))
                                        v21 = 1e10
                                        v2 = append(v2, (fi168 ^ 32767))
                                    }
                                    switch fi168 {
                                    case -1:
                                        v193 := GenS5{F0: 128}
                                        _ = v193
                                    default:
                                        v1 = append(v1, "")
                                        v2 = append(v2, fi168)
                                        if true {
                                            var v194 bool = true
                                            _ = v194
                                        } else {
                                            var v195 uint16 = (v3 / ((1) | 1))
                                            v180 = 1
                                            if true {
                                                if (("foo bar" + "hello") > ("" + (("hello" + "\n\t") + "\n\t"))) {
                                                    v180 -= 454
                                                    switch (fi168 / ((fi168) | 1)) {
                                                    case 127:
                                                        v196 := GenS24{F0: -3}
                                                        _ = v196
                                                    case 32768:
                                                        _ = len(v1)
                                                    default:
                                                        if false {
                                                            switch "hello" {
                                                            case "foo bar":
                                                                v197 := GenS15{F0: (v195 / ((v195) | 1)), F1: 32767}
                                                                _ = v197
                                                            case "hello":
                                                                if len(v4) > 0 {
                                                                    v4[0] = fi172
                                                                }
                                                            case "\n\t":
                                                                if true {
                                                                    if ("hello" != "a") {
                                                                        var v198 uint8 = 246
                                                                        v4 = append(v4, (fi168 * (fi168 | fi172)))
                                                                        _ = v198
                                                                    } else {
                                                                        fmt.Println(v195, v21)
                                                                        switch fi168 {
                                                                        case 1:
                                                                            v3 += v195
                                                                        }
                                                                        var v199 int16 = 42
                                                                        v1 = append(v1, "foo bar")
                                                                        v200 := GenS16{F0: v199, F1: 528}
                                                                        v191 += (v191 << ((18446744073709551615) & 63))
                                                                        if true {
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                v199 += v199
                                                                            case "":
                                                                                v3 = v3
                                                                            case "foo bar":
                                                                                switch ("foo bar" + ("a" + (("foo bar" + "\n\t") + "\n\t"))) {
                                                                                case "":
                                                                                    v201 := GenS25{F0: -1, F1: 553}
                                                                                    var v202 float32 = (1.5 - -2.25)
                                                                                    _ = v201
                                                                                    _ = v202
                                                                                }
                                                                                var v203 uint16 = v195
                                                                                switch (^(fi168)) {
                                                                                case 128:
                                                                                    v4 = append(v4, fi172)
                                                                                case -137:
                                                                                    var v204 int32 = 0
                                                                                    _ = v204
                                                                                case -9223372036854775808:
                                                                                    var v205 int16 = (-(v199))
                                                                                    _ = v205
                                                                                }
                                                                                _ = v203
                                                                            }
                                                                            fmt.Println(v3)
                                                                            v180 -= v180
                                                                            v3 = (v195 >> ((2) & 63))
                                                                        }
                                                                        _ = v199
                                                                        _ = v200
                                                                    }
                                                                } else {
                                                                    v3 = v3
                                                                }
                                                            }
                                                        } else {
                                                            fmt.Println(fi172, v195)
                                                            v206 := GenS16{F0: 1, F1: 627}
                                                            var v207 float32 = (((3.14159 + 3.14159) * -2.25) / 0.1)
                                                            _ = v206
                                                            _ = v207
                                                        }
                                                    }
                                                    v180 -= (v180 ^ (v180 ^ (v180 << ((2) & 63))))
                                                }
                                                v191 += (v191 & v191)
                                            } else {
                                                var v208 int8 = -64
                                                _ = v208
                                            }
                                            _ = v195
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = -2147483648
                                        }
                                    }
                                    fmt.Println(v180, fi172)
                                    _ = v191
                                default:
                                    if true {
                                        v3 -= v3
                                        var v209 float32 = (0.0 - float32((float64(1e10) * v21)))
                                        _ = v209
                                    }
                                }
                            }
                            _ = v180
                        default:
                            v210 := GenS18{F0: 4294967295}
                            _ = v210
                        }
                    case "foo bar":
                        v4 = append(v4, 42)
                    default:
                        v3 = 0
                    }
                    v1 = append(v1, ("foo bar" + "hello"))
                    _ = fi172
                    _ = v173
                }
            }
            for fi211 := 0; fi211 < 7; fi211++ {
                var v212 uint8 = 42
                if false {
                    v3 = v3
                    v4 = append(v4, 32768)
                }
                if false {
                    if (!true) {
                        v213 := GenS20{F0: false, F1: 1}
                        var v214 int16 = -1
                        v21 = (3.14159 - (0.1 * v21))
                        _ = v213
                        _ = v214
                    } else {
                        v21 = 0.0
                        v215 := GenS3{F0: 65535}
                        if (128 == v212) {
                            v216 := GenS10{F0: 18446744073709551615, F1: v3, F2: 288}
                            if (false || true) {
                                var v217 uint16 = 32767
                                var v218 uint64 = 0
                                var v219 uint32 = 2147483648
                                v220 := GenS4{F0: (0 < 0), F1: uint16(v219)}
                                switch ("" + ("foo bar" + (("\n\t" + "foo bar") + ("a" + ("\n\t" + ""))))) {
                                case "hello":
                                    v21 -= ((1.5 / v21) * 1.5)
                                case "\n\t":
                                    var v221 int8 = 1
                                    _ = v221
                                default:
                                    var v222 uint64 = v218
                                    _ = v222
                                }
                                _ = len(v1)
                                _ = v217
                                _ = v218
                                _ = v219
                                _ = v220
                            }
                            _ = v216
                        }
                        _ = v215
                    }
                    switch "\n\t" {
                    case "a":
                        v21 = float64((1e10 - 0.0))
                    default:
                        v212 = 255
                    }
                    if (!(2 != (v212 << ((255) & 63)))) {
                        var v223 float64 = (((1.5 + (float64(1e10) / -2.25)) + v21) * 3.14159)
                        fmt.Println(v3)
                        var v224 uint16 = v3
                        _ = v223
                        _ = v224
                    } else {
                        if (true && (!(v212 == 128))) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if ((32767 <= 32767) && (!true)) {
                                if (false && (1.5 >= -2.25)) {
                                    var v225 int16 = 32767
                                    var v226 uint = 325
                                    _ = v225
                                    _ = v226
                                }
                                if true {
                                    var v227 uint16 = v3
                                    var v228 bool = false
                                    _ = len(v2)
                                    _ = v227
                                    _ = v228
                                } else {
                                    switch 32768 {
                                    case -1:
                                        v3 = v3
                                    case 2147483647:
                                        fmt.Println(v3)
                                    }
                                    if (!(true && (106 == v212))) {
                                        v21 = 3.14159
                                    }
                                }
                                var v229 float32 = (((1e10 * (1.5 * 3.14159)) / (-2.25 + -2.25)) - 0.0)
                                var v230 bool = false
                                var v231 float32 = -2.25
                                _ = v229
                                _ = v230
                                _ = v231
                            }
                            v21 -= (((0.0 + (0.1 - 3.14159)) * 3.14159) / v21)
                        } else {
                            if false {
                                v212 += (v212 | (v212 >> ((337) & 63)))
                                if true {
                                    v1 = append(v1, ("a" + (("a" + "\n\t") + (("hello" + "a") + "\n\t"))))
                                    var v232 int32 = -902
                                    var v233 float32 = 1.5
                                    fmt.Println(v3)
                                    v21 -= (float64(1e10) - 0.1)
                                    v232 = 541
                                    if len(v4) > 0 {
                                        v4[0] = fi168
                                    }
                                    _ = len(v2)
                                    switch fi168 {
                                    case 127:
                                        v3 = v3
                                    case -784:
                                        v234 := GenS7{F0: 4294967295}
                                        _ = v234
                                    case 32767:
                                        if true {
                                            v233 += (v233 - v233)
                                            var v235 uint8 = 127
                                            _ = v235
                                        } else {
                                            switch fi168 {
                                            case -9223372036854775808:
                                                var v236 uint32 = 1
                                                _ = v236
                                            case -864:
                                                v3 += v3
                                            }
                                            fmt.Println(fi168, v233)
                                            if false {
                                                if len(v4) > 0 {
                                                    v4[0] = (fi211 & (fi168 ^ fi211))
                                                }
                                                if len(v1) > 0 {
                                                    v1[0] = "a"
                                                }
                                            } else {
                                                if false {
                                                    if (3.14159 <= 1.5) {
                                                        if (2 != 2) {
                                                            if (true || (0 >= 4294967295)) {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            }
                                                            v4 = append(v4, (fi168 & (fi168 >> ((2) & 63))))
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            v212 += (^(v212))
                                                            v237 := GenS24{F0: 77}
                                                            _ = v237
                                                        }
                                                        v21 = 3.14159
                                                    }
                                                    switch fi168 {
                                                    case 727:
                                                        v232 = (v232 / ((v232) | 1))
                                                    }
                                                    fmt.Println(v3, v232)
                                                } else {
                                                    v238 := GenS26{F0: 0}
                                                    fmt.Println(v212)
                                                    _ = v238
                                                }
                                            }
                                            v239 := GenS11{F0: false, F1: 1, F2: 802}
                                            v212 = (v212 >> ((42) & 63))
                                            _ = v239
                                        }
                                    default:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    v232 += (v232 ^ v232)
                                    _ = v232
                                    _ = v233
                                }
                                var v240 bool = true
                                _ = v240
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = fi168
                                }
                                switch 32768 {
                                case 852:
                                    _ = len(v1)
                                }
                            }
                        }
                        fmt.Println(v212, v212)
                    }
                } else {
                    if (0 >= 399) {
                        v241 := GenS27{F0: 711, F1: 18446744073709551615, F2: 1}
                        v21 -= (1e10 - 0.1)
                        _ = v241
                    } else {
                        if (uint64(v212) > 4294967295) {
                            var v242 uint8 = 128
                            _ = v242
                        } else {
                            var v243 uint32 = 802
                            _ = v243
                        }
                        var v244 int32 = -1
                        _ = v244
                    }
                    fmt.Println(v3, v3, v212)
                }
                if false {
                    switch "foo bar" {
                    case "foo bar":
                        _ = len(v4)
                    default:
                        if true {
                            switch (fi211 >> ((281) & 63)) {
                            case 9223372036854775807:
                                if len(v1) > 0 {
                                    v1[0] = "foo bar"
                                }
                            case 4294967295:
                                if false {
                                    var v245 int = fi168
                                    v21 = v21
                                    _ = v245
                                }
                            }
                            if (true && false) {
                                if (9223372036854775807 == -9223372036854775808) {
                                    v1 = append(v1, "hello")
                                    var v246 int8 = -98
                                    var v247 bool = (!true)
                                    _ = v246
                                    _ = v247
                                }
                                var v248 uint = 723
                                if (!(4294967295 <= 19)) {
                                    v212 = (^(v212))
                                    v249 := GenS18{F0: (-(fi211))}
                                    fmt.Println(v248, fi168, v248)
                                    _ = v249
                                }
                                _ = v248
                            }
                        }
                    }
                    v250 := GenS1{F0: v3}
                    fmt.Println(v212)
                    var v251 uint32 = 2
                    var v252 float64 = ((float64(3.14159) + v21) / float64(1e10))
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v251 -= v251
                    _ = v250
                    _ = v251
                    _ = v252
                }
                _ = fi211
                _ = v212
            }
            _ = len(v1)
            _ = fi168
        }
    case "":
        var v253 int16 = -1
        _ = v253
    case "hello":
        v3 = (v3 + v3)
    default:
        for fi254 := 0; fi254 < 7; fi254++ {
            var v255 int64 = 9223372036854775807
            v4 = append(v4, fi254)
            _ = fi254
            _ = v255
        }
    }
    v21 += float64(0.0)
    v256 := GenS8{F0: -9223372036854775808, F1: -1, F2: "hello"}
    if true {
        for fi257 := 0; fi257 < 4; fi257++ {
            if false {
                if (v21 == 1e10) {
                    if true {
                        for fi258 := 0; fi258 < 14; fi258++ {
                            fmt.Println(fi257, fi257, fi257)
                            v21 = 0.0
                            var v259 uint = 4294967295
                            _ = fi258
                            _ = v259
                        }
                        if (1 >= 65535) {
                            v21 = v21
                            v260 := GenS28{F0: 1, F1: 900}
                            v261 := GenS8{F0: -9223372036854775808, F1: 1, F2: ""}
                            _ = v260
                            _ = v261
                        } else {
                            for fi262 := 0; fi262 < 13; fi262++ {
                                var v263 float32 = 1e10
                                v264 := GenS29{F0: 4294967295, F1: 9, F2: 0}
                                _ = fi262
                                _ = v263
                                _ = v264
                            }
                        }
                        if len(v2) > 0 {
                            v2[0] = (-(fi257))
                        }
                    } else {
                        _ = len(v2)
                        var v265 int64 = 42
                        _ = v265
                    }
                    fmt.Println(v21)
                    switch (fi257 % ((4294967295) | 1)) {
                    case 42:
                        var v266 float32 = 0.0
                        _ = v266
                    case 127:
                        v21 = v21
                    }
                } else {
                    if true {
                        var v267 uint64 = 2
                        v1 = append(v1, "foo bar")
                        for fi268 := 0; fi268 < 12; fi268++ {
                            v3 -= (v3 ^ v3)
                            var v269 uint8 = 162
                            _ = fi268
                            _ = v269
                        }
                        _ = v267
                    }
                }
            }
            v3 = (v3 / ((v3) | 1))
            _ = fi257
        }
    } else {
        v3 = v3
        v21 = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v256
    _ = v3
    _ = v4
}
