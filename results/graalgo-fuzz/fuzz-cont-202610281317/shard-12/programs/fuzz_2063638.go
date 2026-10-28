package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 bool; F2 int64 }
type GenS3 struct { F0 int8; F1 uint32; F2 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS7 struct { F0 uint; F1 int32; F2 int16 }
type GenS8 struct { F0 int8; F1 int32 }
type GenS9 struct { F0 uint64; F1 int32; F2 int64 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int8; F1 int }
type GenS12 struct { F0 uint64; F1 uint16 }
type GenS13 struct { F0 int64; F1 int64 }
type GenS14 struct { F0 int; F1 uint32 }
type GenS15 struct { F0 uint16; F1 string; F2 int32 }
type GenS16 struct { F0 int16; F1 string; F2 int8 }
type GenS17 struct { F0 int32; F1 uint16; F2 bool }
type GenS18 struct { F0 int64; F1 int16; F2 int16 }
type GenS19 struct { F0 int32 }
type GenS20 struct { F0 int8; F1 bool }
func genFn0(p0 int64) uint {
    var v7 string = ""
    v8 := []int{-1, 801, -9223372036854775808}
    var v9 string = v7
    _ = len(v8)
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 65535
}

func main() {
    var v1 float32 = (1.5 - float32((1e10 + float64(1.5))))
    var v2 int64 = 1
    v3 := make(map[int]bool)
    v4 := []int{9223372036854775807, 2147483647, int(v2), -874}
    var v5 bool = true
    var v6 bool = (v5 || v5)
    v6 = true
    v10 := GenS1{F0: (("foo bar" + "\n\t") + "foo bar")}
    v11 := GenS1{F0: ("a" + "")}
    v12 := GenS2{F0: 170, F1: (!("foo bar" > "a")), F2: v2}
    if ("\n\t" != "") {
        if (v6 || ((1.5 + v1) >= (v1 * v1))) {
            v6 = v5
            v13 := GenS3{F0: 125, F1: 1, F2: 0}
            _ = v13
        }
        if (1.5 > 1e10) {
            v14 := GenS1{F0: "foo bar"}
            switch -2147483648 {
            case 60:
                v5 = (18446744073709551615 > 104)
            }
            v6 = v5
            _ = v14
        } else {
            switch ((("a" + "\n\t") + "hello") + "a") {
            case "hello":
                _ = len(v4)
            case "foo bar":
                v5 = (true || ((1e10 * float64((v1 - v1))) >= (3.14159 * (1e10 - 0.0))))
            case "a":
                if true {
                    switch "\n\t" {
                    case "":
                        delete(v3, 0)
                    case "\n\t":
                        v15 := GenS2{F0: 2147483647, F1: v6, F2: v2}
                        _ = v15
                    default:
                        var v16 int8 = 42
                        _ = v16
                    }
                    delete(v3, int(v2))
                    v17 := v3[1]
                    _ = v17
                }
            default:
                if v5 {
                    v18 := v3[-830]
                    _ = v18
                } else {
                    switch -9223372036854775808 {
                    case -57:
                        v4 = append(v4, 32768)
                    case 2147483647:
                        for fi19 := 0; fi19 < 1; fi19++ {
                            fmt.Println(fi19, v5, v6)
                            var v20 uint16 = 32767
                            _ = fi19
                            _ = v20
                        }
                    case 127:
                        v6 = true
                    default:
                        for fi21 := 0; fi21 < 7; fi21++ {
                            for fi22 := 0; fi22 < 14; fi22++ {
                                switch fi21 {
                                case -458:
                                    var v23 uint64 = 1
                                    _ = v23
                                case 32768:
                                    v5 = ((v5 || (!(!((3.14159 - -2.25) == v1)))) && v6)
                                case 2147483647:
                                    var v24 uint16 = 0
                                    _ = v24
                                default:
                                    if len(v4) > 0 {
                                        v4[0] = (-(fi21))
                                    }
                                    if v5 {
                                        switch -525 {
                                        case 650:
                                            v25 := GenS1{F0: "\n\t"}
                                            _ = len(v4)
                                            _ = v25
                                        case 42:
                                            v2 += (-(v2))
                                        case 474:
                                            if (0 > 42) {
                                                var v26 int = (fi21 % ((fi22) | 1))
                                                v6 = v6
                                                _ = v26
                                            } else {
                                                if v5 {
                                                    v27 := GenS1{F0: "a"}
                                                    switch "hello" {
                                                    case "hello":
                                                        fmt.Println(fi21, fi21)
                                                    case "a":
                                                        switch "hello" {
                                                        case "\n\t":
                                                            v1 += -2.25
                                                        default:
                                                            v28 := GenS2{F0: 681, F1: false, F2: v2}
                                                            _ = v28
                                                        }
                                                    case "":
                                                        v5 = false
                                                    }
                                                    _ = v27
                                                } else {
                                                    v5 = (1 < 65535)
                                                    fmt.Println(fi21, v1)
                                                    var v29 string = "foo bar"
                                                    if len(v4) > 0 {
                                                        v4[0] = -2147483648
                                                    }
                                                    _ = v29
                                                }
                                            }
                                        }
                                    } else {
                                        v30 := v3[fi22]
                                        _ = v30
                                    }
                                    v31 := GenS2{F0: 0, F1: ((fi21 << ((genFn0((v2 << ((65535) & 63)))) & 63)) != (fi22 * fi22)), F2: v2}
                                    _ = v31
                                }
                                if (32767 <= -1) {
                                    switch ("" + ("hello" + ("\n\t" + "foo bar"))) {
                                    case "foo bar":
                                        v32 := GenS4{F0: -139}
                                        _ = v32
                                    default:
                                        fmt.Println(fi21, v2, v5)
                                    }
                                } else {
                                    v3[fi21] = (!(v6 || (v2 == -9223372036854775808)))
                                    v33 := GenS1{F0: "a"}
                                    fmt.Println(v5, fi22)
                                    if true {
                                        var v34 float64 = 3.14159
                                        var v35 float64 = 0.0
                                        _ = v34
                                        _ = v35
                                    }
                                    _ = v33
                                }
                                var v36 float32 = (1.5 + (v1 * float32(-2.25)))
                                _ = fi22
                                _ = v36
                            }
                            for fi37 := 0; fi37 < 8; fi37++ {
                                var v38 int = -9223372036854775808
                                _ = fi37
                                _ = v38
                            }
                            _ = fi21
                        }
                    }
                    v39 := GenS5{F0: -3}
                    var v40 uint32 = 106
                    _ = v39
                    _ = v40
                }
            }
        }
        _ = len(v4)
    }
    var v41 int64 = v2
    fmt.Println(v1)
    v42 := GenS6{F0: 127, F1: 940, F2: 858}
    if v6 {
        for fi43 := 0; fi43 < 12; fi43++ {
            switch "\n\t" {
            case "hello":
                v1 -= 0.0
            case "a":
                if (false && (v1 > (v1 * (0.1 - 0.1)))) {
                    if (("hello" + "hello") <= "") {
                        var v44 uint = 2
                        _ = len(v4)
                        _ = v44
                    } else {
                        v5 = v6
                        v3[(fi43 % ((0) | 1))] = (-1 > 32767)
                    }
                    for fi45 := 0; fi45 < 20; fi45++ {
                        var v46 int32 = 2147483647
                        _ = fi45
                        _ = v46
                    }
                    var v47 uint8 = 0
                    _ = v47
                }
            case "":
                v48 := GenS4{F0: 426}
                _ = v48
            default:
                fmt.Println(v41)
            }
            if (42 == int8(fi43)) {
                if v6 {
                    if v6 {
                        v5 = v6
                    } else {
                        v49 := GenS7{F0: 4294967295, F1: -246, F2: 42}
                        switch "hello" {
                        case "":
                            if ((0 < 42) && (65535 >= 4294967295)) {
                                if (!(true || true)) {
                                    switch fi43 {
                                    case -421:
                                        if (42 > (fi43 - fi43)) {
                                            var v50 float32 = v1
                                            for fi51 := 0; fi51 < 12; fi51++ {
                                                fmt.Println(v5)
                                                _ = fi51
                                            }
                                            _ = v50
                                        }
                                    case 997:
                                        for fi52 := 0; fi52 < 3; fi52++ {
                                            var v53 uint32 = 4294967295
                                            switch "" {
                                            case "":
                                                var v54 uint = 1
                                                _ = v54
                                            case "foo bar":
                                                switch "a" {
                                                case "":
                                                    fmt.Println(v53, v5)
                                                case "foo bar":
                                                    var v55 string = "foo bar"
                                                    _ = v55
                                                }
                                            }
                                            _ = fi52
                                            _ = v53
                                        }
                                    }
                                    v56 := GenS8{F0: 0, F1: 42}
                                    _ = v56
                                } else {
                                    fmt.Println(v5)
                                }
                                var v57 uint64 = 4294967295
                                _ = v57
                            }
                        }
                        _ = v49
                    }
                    v6 = ("\n\t" >= "")
                    if false {
                        if (!(127 == 127)) {
                            var v58 uint16 = 32767
                            if (42 <= genFn0((v2 << ((1) & 63)))) {
                                _ = len(v3)
                            } else {
                                for fi59 := 0; fi59 < 14; fi59++ {
                                    if (1 == 2) {
                                        delete(v3, fi43)
                                    } else {
                                        v6 = (!(!(v58 < v58)))
                                    }
                                    if (1 >= 211) {
                                        if (2147483647 <= 634) {
                                            switch "" {
                                            case "a":
                                                var v60 int64 = v2
                                                _ = v60
                                            case "hello":
                                                v5 = (165 <= 1)
                                            }
                                            v61 := GenS3{F0: 100, F1: 526, F2: 4294967295}
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = v61
                                        }
                                        if (((!v6) || (v58 > 542)) || (!false)) {
                                            switch fi43 {
                                            case -47:
                                                var v62 int32 = -614
                                                _ = v62
                                            case -111:
                                                if v5 {
                                                    var v63 float64 = 1e10
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        v4 = append(v4, fi59)
                                                    }
                                                    var v64 string = ("foo bar" + "hello")
                                                    if true {
                                                        v5 = ((42 != 1000) || v5)
                                                        if true {
                                                            switch (fi43 | fi43) {
                                                            case -1:
                                                                var v65 uint8 = 2
                                                                _ = v65
                                                            case 1:
                                                                v66 := GenS3{F0: 100, F1: 238, F2: 861}
                                                                _ = v66
                                                            case 32768:
                                                                v6 = (1 <= 1)
                                                            default:
                                                                var v67 float64 = ((v63 + ((v63 - v63) / 1.5)) * v63)
                                                                _ = v67
                                                            }
                                                        } else {
                                                            fmt.Println(v58, v41)
                                                        }
                                                        if ((v5 && v5) && v6) {
                                                            var v68 uint16 = v58
                                                            v4 = append(v4, (fi43 / ((892) | 1)))
                                                            var v69 int32 = -1
                                                            _ = v68
                                                            _ = v69
                                                        } else {
                                                            v70 := GenS9{F0: 18446744073709551615, F1: -291, F2: v2}
                                                            if v6 {
                                                                v71 := GenS9{F0: 2, F1: -909, F2: 205}
                                                                v72 := GenS2{F0: -1, F1: true, F2: v2}
                                                                _ = v71
                                                                _ = v72
                                                            } else {
                                                                var v73 string = "\n\t"
                                                                v2 = -503
                                                                v41 = 224
                                                                _ = v73
                                                            }
                                                            _ = v70
                                                        }
                                                    }
                                                    if (v41 > v41) {
                                                        if v5 {
                                                            _ = len(v4)
                                                            fmt.Println(v64, fi43)
                                                            if ((v58 + (v58 ^ (v58 - v58))) < v58) {
                                                                if v6 {
                                                                    v41 += (v2 & v2)
                                                                    if (v58 != (^(v58))) {
                                                                        var v74 int = fi59
                                                                        var v75 uint64 = 720
                                                                        _ = v74
                                                                        _ = v75
                                                                    } else {
                                                                        if (583 == (^(fi59))) {
                                                                            v64 = (v64 + v64)
                                                                        } else {
                                                                            v76 := GenS4{F0: 2147483647}
                                                                            if v5 {
                                                                                if (!(false && v5)) {
                                                                                    v2 += -784
                                                                                }
                                                                            }
                                                                            v64 = ""
                                                                            _ = v76
                                                                        }
                                                                    }
                                                                    var v77 int64 = -1
                                                                    _ = v77
                                                                }
                                                                v78 := GenS9{F0: 2, F1: 0, F2: v41}
                                                                _ = v78
                                                            }
                                                            v3[fi59] = (!v6)
                                                        } else {
                                                            var v79 float64 = ((v63 + v63) * 3.14159)
                                                            if v5 {
                                                                var v80 int8 = -1
                                                                _ = v80
                                                            } else {
                                                                fmt.Println(v64, v1, v2)
                                                            }
                                                            _ = v79
                                                        }
                                                        var v81 int64 = v41
                                                        if v5 {
                                                            var v82 uint = 0
                                                            v83 := GenS8{F0: -116, F1: 2147483647}
                                                            if ((v82 / ((genFn0(v41)) | 1)) < (v82 ^ (v82 - v82))) {
                                                                v84 := GenS3{F0: 127, F1: 306, F2: (v82 % (((v82 << ((v82) & 63))) | 1))}
                                                                _ = v84
                                                            }
                                                            _ = v82
                                                            _ = v83
                                                        } else {
                                                            fmt.Println(v81, fi43, v81)
                                                        }
                                                        _ = v81
                                                    } else {
                                                        v6 = (v64 <= ("foo bar" + ((v64 + ("hello" + (v64 + v64))) + "foo bar")))
                                                        if true {
                                                            v63 -= v63
                                                            v58 += 32767
                                                        }
                                                    }
                                                    _ = v63
                                                    _ = v64
                                                } else {
                                                    v1 += v1
                                                    v58 -= 255
                                                    var v85 int = (fi43 << ((795) & 63))
                                                    v58 += v58
                                                    if v5 {
                                                        if (-223 < -1) {
                                                            var v86 uint16 = 644
                                                            if v5 {
                                                                fmt.Println(fi59)
                                                                v87 := GenS8{F0: -1, F1: 42}
                                                                _ = v87
                                                            }
                                                            if v6 {
                                                                if (v5 && v5) {
                                                                    v2 -= -1
                                                                    v85 = v85
                                                                } else {
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    if (255 < v86) {
                                                                        var v88 int32 = -571
                                                                        _ = v88
                                                                    } else {
                                                                        var v89 uint8 = 128
                                                                        _ = v89
                                                                    }
                                                                    var v90 uint = 0
                                                                    _ = v90
                                                                }
                                                                if false {
                                                                    v6 = ("" <= "")
                                                                    v91 := GenS10{F0: (v86 & 32767)}
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    _ = len(v4)
                                                                    _ = v91
                                                                }
                                                            }
                                                            _ = v86
                                                        } else {
                                                            var v92 int8 = 42
                                                            switch "" {
                                                            case "hello":
                                                                v3[fi59] = v6
                                                                var v93 int64 = v41
                                                                _ = v93
                                                            case "":
                                                                v41 += (^(v41))
                                                            }
                                                            _ = len(v4)
                                                            var v94 int = v85
                                                            _ = len(v4)
                                                            _ = v92
                                                            _ = v94
                                                        }
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    }
                                                    v41 = v41
                                                    var v95 uint = uint(v58)
                                                    if v6 {
                                                        switch "a" {
                                                        case "a":
                                                            fmt.Println(v2, v58, v2)
                                                        case "hello":
                                                            v2 = 1
                                                        }
                                                    }
                                                    _ = v85
                                                    _ = v95
                                                }
                                            case -859:
                                                var v96 string = "\n\t"
                                                _ = v96
                                            }
                                        }
                                        v58 += (v58 + (v58 + 255))
                                    } else {
                                        var v97 int32 = 2147483647
                                        v98 := GenS6{F0: 255, F1: -32768, F2: 232}
                                        delete(v3, -9223372036854775808)
                                        fmt.Println(v58, v41, v5)
                                        v41 = 42
                                        _ = v97
                                        _ = v98
                                    }
                                    _ = fi59
                                }
                                if v5 {
                                    _ = len(v4)
                                } else {
                                    var v99 bool = v6
                                    var v100 float64 = 1.5
                                    _ = v100
                                    _ = v99
                                }
                            }
                            _ = v58
                        } else {
                            v101 := GenS11{F0: 36, F1: (fi43 % (((fi43 + (fi43 % ((fi43) | 1)))) | 1))}
                            switch "a" {
                            case "\n\t":
                                fmt.Println(v6, v5, v1)
                            case "":
                                v6 = v5
                            }
                            _ = v101
                        }
                        delete(v3, -426)
                    } else {
                        switch "foo bar" {
                        case "a":
                            v5 = (18446744073709551615 == 733)
                        case "hello":
                            v102 := GenS12{F0: 265, F1: 547}
                            _ = v102
                        case "\n\t":
                            var v103 bool = v5
                            _ = v103
                        }
                        v104 := GenS13{F0: (-(v2)), F1: v2}
                        _ = v104
                    }
                } else {
                    if (0 > 0) {
                        v1 -= float32(((-2.25 / 0.1) + 1.5))
                        if (2 < 988) {
                            var v105 uint64 = 1
                            v6 = (1e10 <= v1)
                            var v106 int16 = 1000
                            _ = v105
                            _ = v106
                        }
                    } else {
                        _ = len(v3)
                    }
                }
                var v107 uint32 = 275
                v5 = v6
                _ = v107
            }
            v108 := GenS14{F0: 2147483647, F1: 967}
            _ = fi43
            _ = v108
        }
        if false {
            if (589 == 32767) {
                v4 = append(v4, 9223372036854775807)
            } else {
                v109 := v3[-770]
                var v110 uint64 = 0
                _ = v109
                _ = v110
            }
        } else {
            if v5 {
                v6 = v6
            } else {
                for fi111 := 0; fi111 < 7; fi111++ {
                    if v5 {
                        v112 := v3[fi111]
                        switch 32768 {
                        case 905:
                            v1 -= 0.1
                        case -1:
                            _ = len(v3)
                        }
                        v1 = (0.1 / float32(-2.25))
                        _ = v112
                    } else {
                        v113 := v3[fi111]
                        _ = v113
                    }
                    v114 := GenS6{F0: 2, F1: 0, F2: 2}
                    for fi115 := 0; fi115 < 8; fi115++ {
                        if (42 <= 1) {
                            v5 = (v41 > -71)
                            if (-468 == -670) {
                                delete(v3, fi111)
                                if (-2.25 < 0.0) {
                                    v5 = (0.1 <= -2.25)
                                    fmt.Println(v41, fi115)
                                    v41 += (v2 << ((391) & 63))
                                }
                                if v6 {
                                    v116 := GenS7{F0: 2, F1: 0, F2: 1000}
                                    _ = v116
                                }
                            } else {
                                v6 = false
                                delete(v3, fi111)
                                switch ("hello" + ("a" + ("a" + "hello"))) {
                                case "foo bar":
                                    if (91 >= 1) {
                                        v6 = v6
                                        fmt.Println(v6, v41)
                                        v5 = (v5 && v5)
                                        switch "" {
                                        case "":
                                            v4 = append(v4, fi111)
                                        default:
                                            var v117 uint = 42
                                            _ = v117
                                        }
                                        v6 = ((1 != -56) && v6)
                                        var v118 int = fi115
                                        if (v1 >= v1) {
                                            if ((((v1 - v1) * ((-2.25 * (v1 + (v1 - v1))) / v1)) - (((v1 + v1) - v1) * v1)) != v1) {
                                                if v6 {
                                                    v4 = append(v4, (fi115 / ((32768) | 1)))
                                                } else {
                                                    v118 += v118
                                                    v118 = (^(fi111))
                                                }
                                                v6 = (((18446744073709551615 < 315) || v5) && (57 < 127))
                                            } else {
                                                v119 := GenS15{F0: 280, F1: (((("hello" + "") + "\n\t") + "") + (("foo bar" + "a") + "hello")), F2: 42}
                                                var v120 string = "\n\t"
                                                v4 = append(v4, v118)
                                                if true {
                                                    var v121 int = v118
                                                    _ = v121
                                                } else {
                                                    var v122 int = (fi115 << ((65535) & 63))
                                                    fmt.Println(v6, v6)
                                                    _ = v122
                                                }
                                                delete(v3, 42)
                                                v1 = (v1 / v1)
                                                _ = v119
                                                _ = v120
                                            }
                                            fmt.Println(v5, v1)
                                        } else {
                                            v123 := v3[(v118 / (((fi115 + fi115)) | 1))]
                                            if v6 {
                                                v118 = v118
                                            } else {
                                                switch fi115 {
                                                case -1:
                                                    v124 := GenS14{F0: fi111, F1: 42}
                                                    _ = v124
                                                }
                                                if (float32((0.1 / 3.14159)) >= 0.1) {
                                                    switch (v118 << ((82) & 63)) {
                                                    case -999:
                                                        switch ("a" + "foo bar") {
                                                        case "\n\t":
                                                            switch ("\n\t" + ("hello" + "foo bar")) {
                                                            case "\n\t":
                                                                fmt.Println(v2, v5, v41)
                                                                if false {
                                                                    var v125 int8 = 68
                                                                    var v126 int64 = v2
                                                                    _ = v125
                                                                    _ = v126
                                                                } else {
                                                                    switch fi111 {
                                                                    case -1:
                                                                        if (42 >= -2147483648) {
                                                                            v127 := GenS7{F0: 255, F1: -1, F2: 1000}
                                                                            v6 = v123
                                                                            var v128 bool = (v118 >= fi111)
                                                                            _ = v127
                                                                            _ = v128
                                                                        } else {
                                                                            var v129 uint16 = 32767
                                                                            switch "" {
                                                                            case "hello":
                                                                                v130 := GenS13{F0: -437, F1: v41}
                                                                                _ = v130
                                                                            case "":
                                                                                fmt.Println(v123, v41, fi115)
                                                                            case "\n\t":
                                                                                var v131 bool = false
                                                                                _ = v131
                                                                            }
                                                                            _ = v129
                                                                        }
                                                                    case -9223372036854775808:
                                                                        v3[(fi115 & (v118 << ((42) & 63)))] = (!(2 < 19))
                                                                    case 9223372036854775807:
                                                                        var v132 uint = 4294967295
                                                                        _ = v132
                                                                    default:
                                                                        if (1 > 690) {
                                                                            v6 = (("\n\t" + "\n\t") >= "\n\t")
                                                                            fmt.Println(v2, v5)
                                                                        } else {
                                                                            var v133 uint32 = 1
                                                                            v134 := GenS16{F0: 560, F1: ("hello" + ("a" + "a")), F2: 73}
                                                                            _ = v133
                                                                            _ = v134
                                                                        }
                                                                    }
                                                                    var v135 uint16 = 32767
                                                                    _ = v135
                                                                }
                                                            case "hello":
                                                                var v136 uint16 = 218
                                                                _ = v136
                                                            default:
                                                                var v137 bool = (v2 != (v2 << ((42) & 63)))
                                                                _ = v137
                                                            }
                                                        case "a":
                                                            v123 = (v5 || v6)
                                                        case "foo bar":
                                                            switch ("\n\t" + "a") {
                                                            case "\n\t":
                                                                v138 := GenS17{F0: 2147483647, F1: 32767, F2: (!(982 <= 1000))}
                                                                _ = v138
                                                            case "":
                                                                var v139 uint32 = 206
                                                                _ = v139
                                                            }
                                                        }
                                                    default:
                                                        var v140 int32 = -2147483648
                                                        _ = v140
                                                    }
                                                    v2 = (^(v41))
                                                    v4 = append(v4, v118)
                                                }
                                                switch ("a" + "a") {
                                                case "":
                                                    if v5 {
                                                        var v141 uint32 = 2147483648
                                                        var v142 float32 = 1e10
                                                        _ = v141
                                                        _ = v142
                                                    }
                                                case "a":
                                                    var v143 float32 = v1
                                                    _ = v143
                                                }
                                            }
                                            _ = v123
                                        }
                                        if (v6 || ((2147483647 < 2147483647) && (!("hello" <= (((("hello" + "a") + "\n\t") + (("a" + "\n\t") + "\n\t")) + ("hello" + "a")))))) {
                                            if v6 {
                                                v144 := GenS18{F0: (v41 % (((v2 % ((9223372036854775807) | 1))) | 1)), F1: -1, F2: 631}
                                                if (((("foo bar" + ("a" + "foo bar")) + "a") < "hello") && (float64(v1) >= ((((0.1 / (3.14159 / ((-2.25 / 3.14159) / ((3.14159 / 1.0) / float64(float32(1e10)))))) - 0.1) * -2.25) + (((float64((v1 - 0.0)) + (0.1 - 3.14159)) * (((1e10 / 1.0) / 3.14159) / (-2.25 + 3.14159))) * 0.0)))) {
                                                    if true {
                                                        if false {
                                                            v145 := GenS2{F0: -1, F1: (!v6), F2: -9223372036854775808}
                                                            v6 = v6
                                                            fmt.Println(v118)
                                                            _ = v145
                                                        } else {
                                                            var v146 int16 = 42
                                                            var v147 int8 = -128
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            _ = v146
                                                            _ = v147
                                                        }
                                                        switch ("" + (("hello" + "") + "hello")) {
                                                        case "a":
                                                            _ = len(v4)
                                                        }
                                                        v2 = v2
                                                    }
                                                } else {
                                                    v3[(fi115 / ((39) | 1))] = (65535 < 255)
                                                    v148 := GenS13{F0: v41, F1: (v41 - v41)}
                                                    _ = v148
                                                }
                                                _ = v144
                                            } else {
                                                var v149 uint16 = 0
                                                _ = len(v3)
                                                _ = v149
                                            }
                                            var v150 uint16 = 929
                                            _ = v150
                                        }
                                        v151 := GenS3{F0: 1, F1: 1, F2: 669}
                                        switch "hello" {
                                        case "hello":
                                            v152 := GenS2{F0: 377, F1: (v1 != (v1 - ((1e10 - v1) - (3.14159 * v1)))), F2: -1}
                                            _ = v152
                                        default:
                                            if v5 {
                                                v153 := GenS7{F0: 314, F1: 1, F2: 42}
                                                _ = v153
                                            }
                                        }
                                        _ = v118
                                        _ = v151
                                    } else {
                                        if (!((v6 && (185 == 42)) || v6)) {
                                            var v154 int8 = -95
                                            fmt.Println(v6)
                                            _ = v154
                                        } else {
                                            var v155 int = (fi111 % ((127) | 1))
                                            v156 := GenS19{F0: -2147483648}
                                            v5 = ((!(!(4294967295 >= 65535))) && (fi115 == 4294967295))
                                            _ = v155
                                            _ = v156
                                        }
                                        var v157 int64 = 0
                                        v158 := GenS20{F0: 0, F1: (1000 == 32767)}
                                        _ = v157
                                        _ = v158
                                    }
                                default:
                                    switch ("hello" + ("a" + "a")) {
                                    case "\n\t":
                                        if (v6 || v6) {
                                            switch "hello" {
                                            case "a":
                                                v41 += (v41 + (v2 ^ v41))
                                            default:
                                                v1 = v1
                                            }
                                            switch 381 {
                                            case -2147483648:
                                                fmt.Println(fi111, v1)
                                            case -9223372036854775808:
                                                v159 := GenS18{F0: (v41 << ((4294967295) & 63)), F1: 42, F2: -1}
                                                _ = v159
                                            default:
                                                var v160 float32 = v1
                                                _ = v160
                                            }
                                        }
                                        var v161 uint = 181
                                        _ = v161
                                    case "":
                                        v162 := v3[983]
                                        var v163 float64 = float64((v1 - float32((3.14159 + 3.14159))))
                                        v163 -= (1e10 * v163)
                                        _ = v162
                                        _ = v163
                                    default:
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                }
                            }
                            v3[(fi111 + 1)] = (!v6)
                        }
                        if v6 {
                            var v164 int = (fi111 ^ fi115)
                            _ = v164
                        } else {
                            var v165 uint16 = 255
                            var v166 int16 = 436
                            _ = v165
                            _ = v166
                        }
                        _ = fi115
                    }
                    _ = fi111
                    _ = v114
                }
                v167 := v3[9223372036854775807]
                _ = v167
            }
            var v168 uint = 2
            if v5 {
                v5 = false
                _ = len(v3)
            }
            _ = v168
        }
        _ = len(v3)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
