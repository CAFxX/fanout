package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS5 struct { F0 int32; F1 uint16 }
type GenS6 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS7 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS11 struct { F0 int32; F1 int32; F2 uint8 }
func genFn0(p0 int8, p1 int32, p2 int16) uint8 {
    if false {
        if true {
            for fi7 := 0; fi7 < 20; fi7++ {
                for fi8 := 0; fi8 < 4; fi8++ {
                    var v9 uint32 = 2147483648
                    var v10 uint32 = 2147483648
                    v11 := GenS1{F0: 18446744073709551615, F1: 32767, F2: int64(fi7)}
                    _ = fi8
                    _ = v10
                    _ = v11
                    _ = v9
                }
                _ = fi7
            }
        } else {
            switch 203 {
            case 127:
                v12 := GenS1{F0: 42, F1: 255, F2: -9223372036854775808}
                _ = v12
            case -2147483648:
                fmt.Println(p1, p2, p1)
            default:
                v13 := GenS2{F0: 960}
                _ = v13
            }
            if (4294967295 >= 565) {
                if ((!false) || false) {
                    v14 := GenS2{F0: 4294967295}
                    for fi15 := 0; fi15 < 13; fi15++ {
                        var v16 string = ("foo bar" + (("foo bar" + "") + "a"))
                        var v17 int8 = (p0 ^ p0)
                        var v18 bool = (v16 == "")
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    _ = v14
                } else {
                    v19 := GenS2{F0: 1}
                    _ = v19
                }
                for fi20 := 0; fi20 < 11; fi20++ {
                    var v21 uint8 = 7
                    _ = fi20
                    _ = v21
                }
            }
        }
        fmt.Println(p2, p2, p1)
    } else {
        if (0 <= 42) {
            p0 = (p0 - (p0 & p0))
            p0 += p0
            if (p2 > (p2 << ((1) & 63))) {
                var v22 int64 = 0
                for fi23 := 0; fi23 < 12; fi23++ {
                    fmt.Println(p0, p2, v22)
                    var v24 bool = true
                    var v25 string = "hello"
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                switch "a" {
                case "\n\t":
                    p0 += (-(p0))
                case "a":
                    v26 := []int{422}
                    _ = v26
                default:
                    v27 := []int{0, 32767, 769}
                    _ = v27
                }
                _ = v22
            } else {
                v28 := []int{-541, 128, -9223372036854775808, 1}
                if (0 > 42) {
                    fmt.Println(p0, p2)
                    for fi29 := 0; fi29 < 11; fi29++ {
                        v30 := GenS1{F0: 893, F1: 0, F2: 1}
                        _ = fi29
                        _ = v30
                    }
                } else {
                    v31 := GenS3{F0: (((("hello" + (("foo bar" + "foo bar") + ("" + "hello"))) + ("hello" + "\n\t")) + ("a" + "a")) + "foo bar")}
                    switch "" {
                    case "hello":
                        _ = len(v28)
                    case "foo bar":
                        var v32 int8 = p0
                        _ = v32
                    }
                    var v33 string = "a"
                    _ = v31
                    _ = v33
                }
                if false {
                    if len(v28) >= 2 {
                        v28 = v28[1:len(v28)]
                    }
                }
                _ = v28
            }
        } else {
            fmt.Println(p1, p1, p1)
            if (!(162 != -2147483648)) {
                var v34 int = 1
                _ = v34
            }
            var v35 int16 = 42
            _ = v35
        }
        p1 = p1
        var v36 uint8 = 128
        _ = v36
    }
    v37 := []int{-2147483648, -247, 892, -81}
    v38 := GenS2{F0: 0}
    switch "\n\t" {
    case "":
        if false {
            if len(v37) >= 2 {
                v37 = v37[1:len(v37)]
            }
            var v39 int8 = (p0 << ((255) & 63))
            _ = v39
        }
    case "foo bar":
        if len(v37) > 0 {
            v37[0] = -715
        }
    case "\n\t":
        p0 = (p0 ^ -9)
    default:
        p1 += (p1 >> ((2) & 63))
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v37
    _ = v38
    return 127
}
func genFn1(p0 uint16, p1 int, p2 int64) int {
    for fi40 := 0; fi40 < 18; fi40++ {
        var v41 int16 = 485
        _ = fi40
        _ = v41
    }
    v42 := []int{(p1 % ((p1) | 1)), -2147483648, 4294967295}
    p1 += (p1 * p1)
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v42
    return 1
}

func main() {
    var v1 float64 = 1e10
    v2 := []string{("a" + ("foo bar" + "")), "\n\t", "a"}
    var v3 int64 = -9223372036854775808
    var v4 uint32 = 2147483648
    v5 := []string{}
    var v6 uint32 = v4
    v1 = ((v1 - v1) - 1.5)
    for fi43 := 0; fi43 < 6; fi43++ {
        v4 -= (v6 >> ((4294967295) & 63))
        var v44 int64 = (^(v3))
        _ = fi43
        _ = v44
    }
    for fi45 := 0; fi45 < 2; fi45++ {
        if true {
            v46 := GenS1{F0: 1, F1: 42, F2: v3}
            v4 -= 708
            _ = v46
        }
        fmt.Println(v1, v1)
        v47 := GenS2{F0: v4}
        _ = fi45
        _ = v47
    }
    var v48 int16 = 1
    var v49 int16 = (-(v48))
    for fi50 := 0; fi50 < 18; fi50++ {
        if true {
            v5 = append(v5, "foo bar")
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v49 += v49
        } else {
            var v51 int = fi50
            switch int(v3) {
            case 42:
                if (0 == 42) {
                    v3 = (v3 << ((uint(v4)) & 63))
                }
            default:
                var v52 uint8 = 1
                _ = v52
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v51
        }
        _ = fi50
    }
    _ = len(v2)
    for fi53 := 0; fi53 < 16; fi53++ {
        v54 := GenS4{F0: 42, F1: 32767, F2: 1}
        _ = fi53
        _ = v54
    }
    switch ("" + "") {
    case "":
        if false {
            fmt.Println(v48, v1)
            for fi55 := 0; fi55 < 20; fi55++ {
                fmt.Println(v48, fi55, v4)
                if true {
                    if true {
                        v56 := GenS4{F0: 42, F1: 0, F2: genFn0(-128, int32(v3), v48)}
                        for fi57 := 0; fi57 < 1; fi57++ {
                            var v58 uint8 = genFn0(int8(fi55), 1, (v49 >> ((4294967295) & 63)))
                            var v59 int32 = 1
                            v48 -= v48
                            switch "a" {
                            case "":
                                if (float32(float64(3.14159)) <= 1e10) {
                                    if ("hello" == "a") {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                } else {
                                    switch (("hello" + "foo bar") + ("a" + ((("\n\t" + "\n\t") + "hello") + "\n\t"))) {
                                    case "":
                                        v6 = v4
                                    case "a":
                                        v2 = append(v2, "hello")
                                    default:
                                        switch fi55 {
                                        case 1:
                                            if (100 > 124) {
                                                v60 := GenS4{F0: 0, F1: 0, F2: v58}
                                                _ = v60
                                            }
                                        default:
                                            if true {
                                                fmt.Println(fi55, v4)
                                                v61 := GenS4{F0: 719, F1: 65535, F2: v58}
                                                if ((v4 ^ (^(v4))) == (v4 / ((0) | 1))) {
                                                    if ("hello" == ("a" + "\n\t")) {
                                                        v62 := GenS5{F0: v59, F1: 42}
                                                        _ = v62
                                                    } else {
                                                        v6 = 2
                                                    }
                                                    v5 = append(v5, "\n\t")
                                                }
                                                if ((("" + "a") + "foo bar") < "") {
                                                    fmt.Println(v48, v1, v4)
                                                    v4 += (v6 % ((v6) | 1))
                                                } else {
                                                    fmt.Println(fi55)
                                                    v2 = append(v2, "")
                                                    v3 -= (v3 / (((v3 - (v3 ^ (v3 - v3)))) | 1))
                                                    if (255 <= 255) {
                                                        v59 = -2147483648
                                                        var v63 int8 = -104
                                                        v64 := GenS2{F0: v4}
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                    if true {
                                                        if (false || false) {
                                                            v3 -= v3
                                                            if false {
                                                                fmt.Println(v48)
                                                                fmt.Println(v59, v6, v49)
                                                                v65 := GenS5{F0: (v59 << ((2) & 63)), F1: 42}
                                                                v59 = (v59 % ((v59) | 1))
                                                                _ = v65
                                                            }
                                                            if true {
                                                                fmt.Println(fi55, v4)
                                                                v66 := GenS1{F0: 1, F1: 42, F2: (v3 | (v3 * (v3 + 0)))}
                                                                _ = v66
                                                            }
                                                        }
                                                        if true {
                                                            v67 := GenS2{F0: v6}
                                                            v68 := GenS6{F0: (^(v6)), F1: 1, F2: v59}
                                                            fmt.Println(v6)
                                                            _ = v67
                                                            _ = v68
                                                        } else {
                                                            var v69 uint = 42
                                                            v1 += (v1 - v1)
                                                            _ = v69
                                                        }
                                                    }
                                                }
                                                _ = v61
                                            } else {
                                                fmt.Println(fi55)
                                            }
                                        }
                                    }
                                }
                            case "foo bar":
                                if true {
                                    v59 = v59
                                    v70 := GenS2{F0: v4}
                                    v71 := GenS7{F0: 255, F1: 430, F2: v3}
                                    fmt.Println(v4, v49)
                                    var v72 int16 = int16(v59)
                                    v73 := GenS8{F0: v4}
                                    var v74 bool = (-128 != 90)
                                    var v75 uint64 = 2
                                    if (!(81 < 1)) {
                                        fmt.Println(v4)
                                    }
                                    _ = v70
                                    _ = v71
                                    _ = v72
                                    _ = v73
                                    _ = v74
                                    _ = v75
                                }
                            case "hello":
                                fmt.Println(v3)
                                var v76 uint16 = 255
                                _ = v76
                            default:
                                fmt.Println(v3)
                            }
                            _ = fi57
                            _ = v58
                            _ = v59
                        }
                        _ = v56
                    }
                    for fi77 := 0; fi77 < 4; fi77++ {
                        if false {
                            var v78 uint32 = v4
                            _ = v78
                        } else {
                            var v79 bool = (false && (!(616 != 2147483647)))
                            if (1 >= 0) {
                                v4 -= (^(v4))
                                var v80 uint = 946
                                fmt.Println(v48)
                                _ = v80
                            }
                            v4 = (v6 ^ (v6 / (((^(v6))) | 1)))
                            _ = v79
                        }
                        var v81 uint = 2
                        _ = fi77
                        _ = v81
                    }
                    var v82 int32 = -2147483648
                    _ = v82
                } else {
                    v4 = 0
                    v83 := GenS9{F0: -749}
                    switch fi55 {
                    case -1:
                        var v84 float32 = (1.5 + float32((0.0 + 1.5)))
                        _ = v84
                    case 2147483647:
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                    case -2147483648:
                        if false {
                            var v85 int32 = 2147483647
                            _ = v85
                        }
                    default:
                        var v86 bool = true
                        _ = v86
                    }
                    _ = v83
                }
                if false {
                    if (651 != 18446744073709551615) {
                        for fi87 := 0; fi87 < 8; fi87++ {
                            if (false || false) {
                                v88 := GenS8{F0: v4}
                                v6 -= (v6 - 65535)
                                _ = v88
                            } else {
                                v3 = v3
                                v1 = ((-2.25 + float64(-2.25)) - (0.0 * v1))
                                v89 := GenS1{F0: 457, F1: 0, F2: v3}
                                v48 += v49
                                var v90 uint = 42
                                _ = v89
                                _ = v90
                            }
                            v49 += v48
                            _ = fi87
                        }
                        v3 -= (v3 >> ((2) & 63))
                        if true {
                            for fi91 := 0; fi91 < 18; fi91++ {
                                v92 := GenS10{F0: 1, F1: (-(v3)), F2: v6}
                                var v93 uint16 = 120
                                switch 32768 {
                                case -845:
                                    fmt.Println(v49)
                                case 32768:
                                    if false {
                                        var v94 string = "\n\t"
                                        _ = v94
                                    }
                                case 850:
                                    switch (fi91 * fi91) {
                                    case 298:
                                        var v95 uint8 = 2
                                        _ = v95
                                    default:
                                        v96 := GenS5{F0: -764, F1: (v93 | 123)}
                                        _ = v96
                                    }
                                }
                                _ = fi91
                                _ = v92
                                _ = v93
                            }
                            if (fi55 >= fi55) {
                                var v97 float64 = v1
                                _ = v97
                            } else {
                                v98 := GenS11{F0: 1, F1: -1, F2: 127}
                                var v99 int = fi55
                                var v100 int32 = 529
                                _ = v100
                                _ = v98
                                _ = v99
                            }
                        } else {
                            for fi101 := 0; fi101 < 20; fi101++ {
                                v49 -= (v48 * v48)
                                _ = fi101
                            }
                        }
                    }
                }
                _ = fi55
            }
            for fi102 := 0; fi102 < 14; fi102++ {
                for fi103 := 0; fi103 < 5; fi103++ {
                    if len(v2) > 0 {
                        v2[0] = (("\n\t" + "hello") + "a")
                    }
                    if (v1 >= 3.14159) {
                        if (395 <= 807) {
                            var v104 uint = 0
                            _ = v104
                        }
                        v6 = (v4 % ((v4) | 1))
                        var v105 int = -637
                        var v106 int32 = 1
                        _ = v105
                        _ = v106
                    } else {
                        if false {
                            fmt.Println(v48, fi102, v1)
                            var v107 int16 = -118
                            switch genFn1(42, fi103, v3) {
                            case 250:
                                var v108 uint32 = v4
                                _ = v108
                            }
                            _ = v107
                        } else {
                            v109 := GenS2{F0: (v4 << ((2) & 63))}
                            _ = v109
                        }
                        var v110 int8 = 9
                        v48 -= (^(v49))
                        switch ("hello" + ("" + ("a" + ""))) {
                        case "":
                            if true {
                                if (false && (65535 != uint16(v6))) {
                                    if true {
                                        var v111 int8 = v110
                                        fmt.Println(v3, v111, v4)
                                        _ = v111
                                    } else {
                                        var v112 int32 = -354
                                        switch fi103 {
                                        case 38:
                                            var v113 int = (fi102 & fi103)
                                            _ = v113
                                        case 836:
                                            v114 := GenS6{F0: v4, F1: 255, F2: 2147483647}
                                            fmt.Println(v48, v6)
                                            _ = v114
                                        }
                                        _ = v112
                                    }
                                    v6 += 65535
                                    v115 := GenS11{F0: 2147483647, F1: 2147483647, F2: 128}
                                    _ = v115
                                } else {
                                    v3 -= (v3 % ((755) | 1))
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    if (false || (((false || (!(!(2147483647 >= v4)))) && false) && false)) {
                                        var v116 uint = 2
                                        fmt.Println(v110, v6)
                                        if len(v5) > 0 {
                                            v5[0] = ("foo bar" + "")
                                        }
                                        _ = v116
                                    } else {
                                        v4 = v6
                                        v117 := GenS4{F0: 18446744073709551615, F1: 32767, F2: genFn0((^(v110)), -1, (v48 & (v48 - -1)))}
                                        if true {
                                            v48 = (v49 + 1)
                                            var v118 int8 = (v110 + v110)
                                            if ((-(v3)) > (v3 / ((int64(v4)) | 1))) {
                                                v3 += (v3 + (v3 & (v3 - (v3 << ((2) & 63)))))
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                v3 = v3
                                            }
                                            _ = v118
                                        }
                                        _ = v117
                                    }
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                if true {
                                    fmt.Println(v1)
                                    v1 += (0.1 - (v1 * (v1 + v1)))
                                } else {
                                    if (1e10 == ((0.1 * 1.5) - 0.1)) {
                                        v119 := GenS7{F0: 255, F1: 2, F2: v3}
                                        _ = v119
                                    } else {
                                        var v120 float64 = 1e10
                                        switch fi103 {
                                        case -112:
                                            v121 := GenS8{F0: (v6 * v4)}
                                            _ = v121
                                        }
                                        _ = v120
                                    }
                                    v49 = v49
                                }
                            }
                        }
                        _ = v110
                    }
                    var v122 float32 = (1.5 - ((1.5 / (3.14159 - -2.25)) / 0.1))
                    _ = fi103
                    _ = v122
                }
                _ = len(v5)
                _ = fi102
            }
        } else {
            v123 := GenS2{F0: v4}
            fmt.Println(v49, v6, v49)
            fmt.Println(v4, v49, v49)
            _ = v123
        }
    default:
        v3 = v3
    }
    if false {
        var v124 float64 = v1
        if true {
            v3 = (v3 * v3)
        } else {
            v6 = (v4 & v4)
            var v125 uint32 = (v4 ^ (v4 >> ((255) & 63)))
            v126 := GenS11{F0: 42, F1: 0, F2: 2}
            _ = v125
            _ = v126
        }
        v2 = append(v2, "foo bar")
        _ = v124
    }
    for fi127 := 0; fi127 < 1; fi127++ {
        v48 = 331
        _ = len(v2)
        _ = fi127
    }
    v6 = 56
    v128 := GenS1{F0: 0, F1: 42, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v128
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
