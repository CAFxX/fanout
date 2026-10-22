package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 uint; F2 int16 }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 int32; F1 uint16 }
type GenS8 struct { F0 int32; F1 string; F2 uint32 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int32; F1 bool }
type GenS11 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS12 struct { F0 string; F1 uint32 }
type GenS13 struct { F0 uint; F1 uint16 }
type GenS14 struct { F0 string; F1 bool }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS18 struct { F0 int8 }
type GenS19 struct { F0 uint; F1 string }
type GenS20 struct { F0 uint16; F1 uint16 }
type GenS21 struct { F0 bool; F1 uint32; F2 bool }
type GenS22 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS23 struct { F0 int32 }
type GenS24 struct { F0 uint; F1 string }
type GenS25 struct { F0 string; F1 bool; F2 int }
type GenS26 struct { F0 uint32; F1 uint; F2 int64 }
type GenS27 struct { F0 int }
type GenS28 struct { F0 uint32; F1 int8 }
type GenS29 struct { F0 int8; F1 int16; F2 int64 }
func genFn0(p0 uint32, p1 int8, p2 uint64) int16 {
    v6 := []int{9223372036854775807}
    var v7 float64 = float64(0.0)
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 11; fi9++ {
            if true {
                if len(v6) > 0 {
                    v6[0] = fi8
                }
                switch fi9 {
                case -2147483648:
                    var v10 int8 = (p1 ^ 127)
                    _ = v10
                case 2147483647:
                    p0 = 2
                case 0:
                    var v11 uint16 = uint16(p2)
                    _ = v11
                }
                p1 += -41
            } else {
                fmt.Println(p1, v7)
            }
            var v12 bool = (32767 < 0)
            v13 := GenS1{F0: p0, F1: (!v12)}
            _ = fi9
            _ = v12
            _ = v13
        }
        var v14 string = ""
        _ = fi8
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return int16(p0)
}
func genFn1(p0 int64, p1 int16) int8 {
    fmt.Println(p1, p0)
    if (genFn0(863, 100, 42) > (p1 | (p1 & p1))) {
        fmt.Println(p1)
        fmt.Println(p0, p1)
    } else {
        var v15 uint16 = 255
        var v16 int8 = 127
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 uint16 = 630
    var v2 int64 = -630
    v3 := []string{(("\n\t" + "hello") + "\n\t"), (("a" + ("\n\t" + ("" + ("" + (("" + "hello") + "hello"))))) + "a"), "foo bar"}
    v4 := make(map[string]bool)
    var v5 int16 = 1
    for fi17 := 0; fi17 < 19; fi17++ {
        for fi18 := 0; fi18 < 2; fi18++ {
            _ = len(v3)
            fmt.Println(fi18, v1)
            _ = fi18
        }
        _ = len(v3)
        for fi19 := 0; fi19 < 12; fi19++ {
            delete(v4, ("" + ""))
            var v20 int16 = 534
            _ = fi19
            _ = v20
        }
        _ = fi17
    }
    if false {
        for fi21 := 0; fi21 < 17; fi21++ {
            switch (fi21 * (fi21 ^ (-(fi21)))) {
            case 700:
                fmt.Println(fi21, v2, v5)
            }
            _ = fi21
        }
        delete(v4, "hello")
        v2 = 0
    }
    if (!false) {
        fmt.Println(v1)
        if (66 != 2) {
            for fi22 := 0; fi22 < 9; fi22++ {
                var v23 float32 = (3.14159 + float32(0.1))
                v5 = (v5 / ((genFn0(42, genFn1(v2, 0), 0)) | 1))
                _ = fi22
                _ = v23
            }
        } else {
            var v24 uint32 = 4294967295
            if ((v2 ^ v2) < v2) {
                delete(v4, "foo bar")
                v25 := GenS2{F0: -347, F1: (-(v5))}
                _ = v25
            } else {
                v5 = v5
            }
            _ = v24
        }
    } else {
        v26 := GenS3{F0: 0}
        _ = v26
    }
    v27 := GenS1{F0: 4294967295, F1: ((int8(v5) <= -128) || true)}
    fmt.Println(v2)
    for fi28 := 0; fi28 < 14; fi28++ {
        _ = len(v4)
        v29 := GenS1{F0: 4294967295, F1: (!false)}
        var v30 int16 = v5
        _ = fi28
        _ = v29
        _ = v30
    }
    if (-1 != 128) {
        for fi31 := 0; fi31 < 14; fi31++ {
            var v32 int32 = 42
            v1 += v1
            if true {
                if ((100 != -1) && false) {
                    v32 += v32
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                if (((255 != 128) || false) || ((1 != 26) || (42 == 2))) {
                    for fi33 := 0; fi33 < 15; fi33++ {
                        if true {
                            v1 = (^(v1))
                            v3 = append(v3, "a")
                        }
                        _ = fi33
                    }
                    var v34 uint32 = 1
                    _ = v34
                }
                v35 := v4[("hello" + "hello")]
                for fi36 := 0; fi36 < 13; fi36++ {
                    var v37 uint8 = 255
                    v38 := GenS4{F0: "\n\t"}
                    v39 := GenS1{F0: 632, F1: v35}
                    _ = fi36
                    _ = v37
                    _ = v38
                    _ = v39
                }
                _ = v35
            }
            _ = fi31
            _ = v32
        }
    }
    v1 = (^(v1))
    v1 += 65535
    if ((("foo bar" == (("hello" + "foo bar") + "\n\t")) && ((856 < 42) && false)) && (-1 != (v5 | 32767))) {
        fmt.Println(v1)
        if ((float32((0.1 + 3.14159)) * (0.0 + 3.14159)) > (1e10 + (0.0 / -2.25))) {
            for fi40 := 0; fi40 < 5; fi40++ {
                switch "hello" {
                case "":
                    v4["hello"] = false
                default:
                    fmt.Println(v5, v2, v2)
                }
                v2 = (v2 * 0)
                _ = fi40
            }
        } else {
            switch -408 {
            case -1:
                v5 = v5
            case -2147483648:
                if false {
                    if true {
                        switch -1 {
                        case -2147483648:
                            fmt.Println(v1)
                        case 2147483647:
                            v2 = (v2 << ((4294967295) & 63))
                        default:
                            var v41 uint = 953
                            _ = v41
                        }
                    } else {
                        var v42 int8 = 81
                        var v43 uint16 = v1
                        v44 := GenS1{F0: 2147483647, F1: true}
                        _ = v42
                        _ = v43
                        _ = v44
                    }
                    v45 := GenS5{F0: 100, F1: 18446744073709551615, F2: genFn0(2147483647, 69, 1)}
                    _ = v45
                }
            default:
                v1 -= (v1 & 181)
            }
        }
        if (!(true && (((false && false) && (((92 < 2) || (0.0 < 0.1)) && (0.1 >= 1e10))) || ((v2 ^ (-(v2))) == (-(v2)))))) {
            if false {
                if (v2 == (v2 - v2)) {
                    var v46 float64 = float64(3.14159)
                    if true {
                        for fi47 := 0; fi47 < 3; fi47++ {
                            switch (-(fi47)) {
                            case 32767:
                                if ((!true) || true) {
                                    var v48 uint = 18446744073709551615
                                    _ = v48
                                }
                            default:
                                v49 := GenS4{F0: "a"}
                                _ = v49
                            }
                            if (1 >= 255) {
                                if (!false) {
                                    v46 = v46
                                    v50 := GenS2{F0: -2147483648, F1: 1000}
                                    switch "foo bar" {
                                    case "a":
                                        if (766 < 42) {
                                            v51 := GenS4{F0: "foo bar"}
                                            for fi52 := 0; fi52 < 3; fi52++ {
                                                v53 := GenS6{F0: 951, F1: v5}
                                                v54 := v4[("foo bar" + "")]
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = fi52
                                                _ = v53
                                                _ = v54
                                            }
                                            v2 -= -1
                                            _ = v51
                                        }
                                    default:
                                        if len(v3) > 0 {
                                            v3[0] = ""
                                        }
                                    }
                                    _ = v50
                                } else {
                                    var v55 float32 = (-2.25 * (1e10 * (1e10 / 1.0)))
                                    v1 += v1
                                    if false {
                                        if true {
                                            var v56 uint32 = 1
                                            _ = v56
                                        }
                                        var v57 uint8 = 16
                                        _ = v57
                                    } else {
                                        if false {
                                            v5 -= (v5 + (v5 >> ((905) & 63)))
                                        } else {
                                            if (("a" + "\n\t") >= "hello") {
                                                v4["foo bar"] = (128 > 0)
                                            } else {
                                                if len(v3) > 0 {
                                                    v3[0] = "hello"
                                                }
                                                for fi58 := 0; fi58 < 17; fi58++ {
                                                    if (560 < 4294967295) {
                                                        v46 -= ((v46 / v46) * (0.0 - 0.1))
                                                        if ((759 != 639) || (4294967295 <= 42)) {
                                                            if true {
                                                                if true {
                                                                    if ((fi58 << ((1) & 63)) == -671) {
                                                                        switch (fi58 & (fi47 % (((fi58 * fi58)) | 1))) {
                                                                        case -800:
                                                                            if (!false) {
                                                                                v5 -= v5
                                                                            } else {
                                                                                var v59 float64 = (float64(-2.25) / v46)
                                                                                _ = v59
                                                                            }
                                                                        default:
                                                                            if true {
                                                                                fmt.Println(fi58)
                                                                                if (291 > 2) {
                                                                                    var v60 int32 = 42
                                                                                    _ = v60
                                                                                }
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = ("a" + "\n\t")
                                                                                }
                                                                            } else {
                                                                                v46 = v46
                                                                            }
                                                                        }
                                                                    }
                                                                    var v61 uint = 18446744073709551615
                                                                    var v62 int = fi58
                                                                    _ = v61
                                                                    _ = v62
                                                                } else {
                                                                    v55 += (-2.25 - (v55 + (v55 / (0.0 + ((0.1 - ((1e10 + (((v55 + v55) + (1.5 / -2.25)) * v55)) * v55)) * v55)))))
                                                                    v63 := GenS1{F0: 801, F1: true}
                                                                    v5 = v5
                                                                    _ = v63
                                                                }
                                                                fmt.Println(fi47, v55)
                                                                switch fi47 {
                                                                case 9223372036854775807:
                                                                    v64 := GenS1{F0: 4294967295, F1: (18446744073709551615 > 2)}
                                                                    _ = v64
                                                                default:
                                                                    var v65 bool = false
                                                                    _ = v65
                                                                }
                                                                var v66 int8 = 127
                                                                v67 := GenS6{F0: v2, F1: 974}
                                                                var v68 string = "foo bar"
                                                                _ = v66
                                                                _ = v67
                                                                _ = v68
                                                            } else {
                                                                if true {
                                                                    if ((127 <= -128) || true) {
                                                                        var v69 uint32 = 687
                                                                        var v70 uint32 = (v69 * v69)
                                                                        var v71 int16 = 32767
                                                                        _ = v69
                                                                        _ = v70
                                                                        _ = v71
                                                                    } else {
                                                                        v5 -= 1000
                                                                        v72 := GenS2{F0: -261, F1: (-(v5))}
                                                                        _ = v72
                                                                    }
                                                                    var v73 int16 = v5
                                                                    switch -9223372036854775808 {
                                                                    case 501:
                                                                        v74 := GenS4{F0: ("a" + (("foo bar" + "foo bar") + ""))}
                                                                        _ = v74
                                                                    case 2147483647:
                                                                        v3 = append(v3, "\n\t")
                                                                    case -9223372036854775808:
                                                                        if false {
                                                                            v4[(("" + "foo bar") + "foo bar")] = (true || (false && (fi47 < (^(fi47)))))
                                                                        } else {
                                                                            var v75 uint8 = 127
                                                                            switch "" {
                                                                            case "":
                                                                                fmt.Println(fi58, v46)
                                                                            case "a":
                                                                                v4["a"] = false
                                                                            }
                                                                            var v76 uint = 164
                                                                            _ = v75
                                                                            _ = v76
                                                                        }
                                                                        v77 := v4["\n\t"]
                                                                        _ = v77
                                                                    }
                                                                    fmt.Println(v55, v2)
                                                                    var v78 string = ""
                                                                    _ = v73
                                                                    _ = v78
                                                                } else {
                                                                    v55 -= -2.25
                                                                    fmt.Println(v1, fi58)
                                                                }
                                                                var v79 int32 = -170
                                                                var v80 int = (fi47 & (fi58 % ((0) | 1)))
                                                                v79 = (v79 / (((^(v79))) | 1))
                                                                _ = v79
                                                                _ = v80
                                                            }
                                                            var v81 int8 = genFn1((v2 & (v2 / ((1) | 1))), (v5 & -841))
                                                            var v82 uint64 = 887
                                                            _ = v81
                                                            _ = v82
                                                        }
                                                        v1 = (v1 / ((v1) | 1))
                                                    } else {
                                                        switch (fi47 / (((fi47 | 32768)) | 1)) {
                                                        case 32768:
                                                            v46 = 3.14159
                                                        }
                                                        switch "hello" {
                                                        case "":
                                                            v46 -= float64((v55 + 1.5))
                                                        case "\n\t":
                                                            v3 = append(v3, "")
                                                        case "a":
                                                            var v83 uint = 985
                                                            v84 := GenS7{F0: 0, F1: 32767}
                                                            v85 := GenS8{F0: 1, F1: "foo bar", F2: 4294967295}
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            _ = v83
                                                            _ = v84
                                                            _ = v85
                                                        default:
                                                            var v86 bool = true
                                                            v87 := GenS3{F0: 671}
                                                            v88 := GenS1{F0: 65535, F1: true}
                                                            _ = v86
                                                            _ = v87
                                                            _ = v88
                                                        }
                                                    }
                                                    switch (fi58 % (((fi58 / ((fi47) | 1))) | 1)) {
                                                    case 127:
                                                        v5 += 32767
                                                    case 32768:
                                                        if false {
                                                            var v89 int32 = 827
                                                            if (true || false) {
                                                                v5 += (v5 + int16(fi58))
                                                            }
                                                            v90 := GenS9{F0: 1}
                                                            var v91 uint32 = 42
                                                            if false {
                                                                v92 := GenS6{F0: (v2 & 0), F1: v5}
                                                                v2 = (v2 / (((v2 ^ (v2 ^ (v2 % (((v2 | 769)) | 1))))) | 1))
                                                                _ = v92
                                                            } else {
                                                                v89 = 2147483647
                                                                v4[""] = false
                                                            }
                                                            _ = v89
                                                            _ = v90
                                                            _ = v91
                                                        } else {
                                                            switch fi47 {
                                                            case 9223372036854775807:
                                                                v93 := GenS6{F0: v2, F1: v5}
                                                                _ = v93
                                                            }
                                                            switch "a" {
                                                            case "\n\t":
                                                                v94 := GenS10{F0: -656, F1: false}
                                                                _ = v94
                                                            case "a":
                                                                v1 += (^(v1))
                                                            default:
                                                                if (false && (164 > 1)) {
                                                                    v5 = (-(v5))
                                                                    fmt.Println(fi58, v5, v46)
                                                                    delete(v4, ("hello" + "\n\t"))
                                                                    switch (-(fi47)) {
                                                                    case -298:
                                                                        v3 = append(v3, "foo bar")
                                                                    default:
                                                                        var v95 int32 = 256
                                                                        _ = v95
                                                                    }
                                                                }
                                                            }
                                                            fmt.Println(fi58, v1)
                                                            var v96 uint32 = 745
                                                            v97 := GenS11{F0: (v96 + (v96 % (((v96 ^ (^(v96)))) | 1))), F1: 452, F2: 0}
                                                            v98 := GenS5{F0: genFn1((v2 - (v2 % ((-9223372036854775808) | 1))), v5), F1: 18446744073709551615, F2: -32768}
                                                            if false {
                                                                v4[""] = false
                                                                v1 = uint16(v96)
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    v99 := GenS10{F0: 0, F1: true}
                                                                    v100 := GenS6{F0: 939, F1: (v5 * genFn0(v96, int8(v1), 1))}
                                                                    switch ((((("" + ("foo bar" + "hello")) + "foo bar") + "\n\t") + "hello") + "hello") {
                                                                    case "foo bar":
                                                                        v101 := GenS12{F0: "\n\t", F1: (v96 | v96)}
                                                                        v102 := v4[(("" + "\n\t") + "hello")]
                                                                        _ = v101
                                                                        _ = v102
                                                                    default:
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            if (true && false) {
                                                                                v2 -= v2
                                                                                if ("" >= ("" + (("hello" + "\n\t") + "a"))) {
                                                                                    if (1 == 1) {
                                                                                        v1 = (^(v1))
                                                                                        fmt.Println(v96)
                                                                                        var v103 int8 = -71
                                                                                        _ = v103
                                                                                    }
                                                                                    switch "a" {
                                                                                    case "\n\t":
                                                                                        var v104 uint64 = 18446744073709551615
                                                                                        _ = v104
                                                                                    case "hello":
                                                                                        fmt.Println(v96, v2, v1)
                                                                                    case "a":
                                                                                        if false {
                                                                                            v105 := GenS13{F0: 1, F1: (^(v1))}
                                                                                            fmt.Println(v46, v55, fi47)
                                                                                            v96 = (v96 & (v96 + v96))
                                                                                            _ = v105
                                                                                        }
                                                                                    default:
                                                                                        v106 := GenS10{F0: 427, F1: (v55 == v55)}
                                                                                        _ = v106
                                                                                    }
                                                                                    v5 += (-(v5))
                                                                                    v107 := GenS14{F0: "foo bar", F1: (false && (((42 < 148) && (-1 >= 293)) || ((v5 % ((32767) | 1)) == 32767)))}
                                                                                    _ = v107
                                                                                } else {
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = "foo bar"
                                                                                    }
                                                                                }
                                                                            }
                                                                        case "foo bar":
                                                                            var v108 uint = 2
                                                                            _ = v108
                                                                        }
                                                                    }
                                                                    _ = v100
                                                                    _ = v99
                                                                default:
                                                                    v96 += v96
                                                                }
                                                                v109 := GenS14{F0: "\n\t", F1: (4294967295 >= 18446744073709551615)}
                                                                var v110 int8 = 127
                                                                v111 := GenS10{F0: -2147483648, F1: ((((v55 * v55) + v55) - (v55 + (v55 / v55))) > (0.0 + 0.1))}
                                                                v96 -= v96
                                                                v112 := GenS10{F0: 771, F1: true}
                                                                _ = v109
                                                                _ = v110
                                                                _ = v111
                                                                _ = v112
                                                            } else {
                                                                var v113 uint16 = (v1 / (((v1 >> ((65535) & 63))) | 1))
                                                                if (2 >= 42) {
                                                                    switch (fi47 - fi58) {
                                                                    case 4294967295:
                                                                        v114 := GenS15{F0: (v2 / (((v2 + v2)) | 1))}
                                                                        _ = v114
                                                                    default:
                                                                        if (!(fi58 <= fi58)) {
                                                                            v55 = v55
                                                                            v115 := GenS13{F0: uint(v96), F1: v113}
                                                                            _ = v115
                                                                        } else {
                                                                            fmt.Println(v55, v96)
                                                                            fmt.Println(v46, fi58)
                                                                        }
                                                                    }
                                                                    var v116 bool = true
                                                                    if (0.0 == (v55 + ((v55 + (1.5 * 0.1)) / v55))) {
                                                                        v117 := GenS10{F0: 1, F1: v116}
                                                                        switch ("hello" + "a") {
                                                                        case "foo bar":
                                                                            switch "hello" {
                                                                            case "a":
                                                                                var v118 float32 = (float32(v46) - -2.25)
                                                                                _ = len(v3)
                                                                                _ = v118
                                                                            case "":
                                                                                if (42 >= v96) {
                                                                                    v5 = int16(v96)
                                                                                    if v116 {
                                                                                        v119 := GenS16{F0: 42}
                                                                                        if v116 {
                                                                                            var v120 uint = 935
                                                                                            _ = v120
                                                                                        } else {
                                                                                            v46 = 1.5
                                                                                        }
                                                                                        var v121 int8 = 107
                                                                                        _ = v119
                                                                                        _ = v121
                                                                                    }
                                                                                    v3 = append(v3, "hello")
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                    if ((v113 >> ((0) & 63)) >= 107) {
                                                                                        var v122 int16 = (v5 | (^(v5)))
                                                                                        _ = v122
                                                                                    } else {
                                                                                        var v123 bool = (0 >= v113)
                                                                                        fmt.Println(v96, v96, v116)
                                                                                        if (!(v46 >= v46)) {
                                                                                            v124 := GenS17{F0: 4294967295, F1: 4294967295, F2: 924}
                                                                                            if true {
                                                                                                _ = len(v3)
                                                                                            } else {
                                                                                                if v116 {
                                                                                                    var v125 uint = 0
                                                                                                    _ = v125
                                                                                                } else {
                                                                                                    v123 = true
                                                                                                    v96 += v96
                                                                                                    var v126 int = fi47
                                                                                                    _ = v126
                                                                                                }
                                                                                                v55 = 0.0
                                                                                            }
                                                                                            _ = v124
                                                                                        }
                                                                                        _ = v123
                                                                                    }
                                                                                } else {
                                                                                    v2 -= v2
                                                                                }
                                                                            default:
                                                                                var v127 float32 = ((((0.1 + v55) - 3.14159) / float32(0.1)) * v55)
                                                                                _ = v127
                                                                            }
                                                                        default:
                                                                            v46 = ((0.1 - float64(v55)) + 0.0)
                                                                        }
                                                                        if v116 {
                                                                            var v128 float32 = v55
                                                                            _ = v128
                                                                        }
                                                                        _ = v117
                                                                    } else {
                                                                        if true {
                                                                            v129 := GenS18{F0: genFn1(271, v5)}
                                                                            _ = v129
                                                                        } else {
                                                                            switch "foo bar" {
                                                                            case "\n\t":
                                                                                fmt.Println(v2)
                                                                            case "foo bar":
                                                                                fmt.Println(v2, v2, v2)
                                                                            case "":
                                                                                v130 := GenS19{F0: 4294967295, F1: "a"}
                                                                                if (!(-1 != -1)) {
                                                                                    v131 := GenS20{F0: v113, F1: (v1 & (v1 + (v113 * (v1 & (v113 << ((18446744073709551615) & 63))))))}
                                                                                    _ = v131
                                                                                }
                                                                                _ = v130
                                                                            }
                                                                            var v132 uint = 1
                                                                            _ = v132
                                                                        }
                                                                    }
                                                                    _ = v116
                                                                } else {
                                                                    v133 := GenS9{F0: v113}
                                                                    v134 := GenS21{F0: true, F1: v96, F2: true}
                                                                    v2 -= v2
                                                                    _ = v133
                                                                    _ = v134
                                                                }
                                                                if false {
                                                                    if false {
                                                                        switch fi58 {
                                                                        case 42:
                                                                            switch "\n\t" {
                                                                            case "a":
                                                                                var v135 uint8 = 1
                                                                                _ = v135
                                                                            case "":
                                                                                fmt.Println(v46)
                                                                            case "foo bar":
                                                                                switch "\n\t" {
                                                                                case "foo bar":
                                                                                    fmt.Println(fi58, v1)
                                                                                case "\n\t":
                                                                                    v3 = append(v3, "a")
                                                                                default:
                                                                                    v1 = 2
                                                                                }
                                                                            default:
                                                                                if (586 >= 71) {
                                                                                    fmt.Println(fi47, v96, v1)
                                                                                } else {
                                                                                    var v136 string = ("foo bar" + ("\n\t" + ((("foo bar" + ((("\n\t" + "a") + "") + "foo bar")) + ("foo bar" + ("" + "foo bar"))) + ((("\n\t" + "foo bar") + "foo bar") + "hello"))))
                                                                                    v136 = (v136 + "a")
                                                                                    _ = v136
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    var v137 int = (fi58 % (((fi47 << ((18446744073709551615) & 63))) | 1))
                                                                    if false {
                                                                        fmt.Println(fi58, fi58)
                                                                        delete(v4, "foo bar")
                                                                        fmt.Println(v55)
                                                                        var v138 float64 = v46
                                                                        _ = len(v3)
                                                                        fmt.Println(v96)
                                                                        var v139 int32 = 681
                                                                        var v140 uint = 2
                                                                        v96 = 801
                                                                        v113 = v113
                                                                        var v141 int32 = v139
                                                                        _ = v138
                                                                        _ = v139
                                                                        _ = v140
                                                                        _ = v141
                                                                    }
                                                                    v142 := GenS22{F0: v96, F1: v96, F2: 482}
                                                                    fmt.Println(v55, v1)
                                                                    _ = v137
                                                                    _ = v142
                                                                } else {
                                                                    v143 := GenS23{F0: 2147483647}
                                                                    _ = v143
                                                                }
                                                                _ = v113
                                                            }
                                                            v144 := GenS9{F0: (v1 - 2)}
                                                            _ = v144
                                                            _ = v96
                                                            _ = v97
                                                            _ = v98
                                                        }
                                                    case 730:
                                                        var v145 uint32 = 1
                                                        _ = v145
                                                    }
                                                    v146 := GenS17{F0: 1, F1: 0, F2: 42}
                                                    _ = fi58
                                                    _ = v146
                                                }
                                                v147 := GenS6{F0: v2, F1: (v5 / ((genFn0(2147483648, 126, 18446744073709551615)) | 1))}
                                                _ = v147
                                            }
                                            fmt.Println(v46, v1)
                                        }
                                        if (42 < 4294967295) {
                                            delete(v4, "foo bar")
                                            switch fi47 {
                                            case 2147483647:
                                                for fi148 := 0; fi148 < 10; fi148++ {
                                                    if false {
                                                        v55 -= float32(0.1)
                                                    }
                                                    _ = fi148
                                                }
                                            case 128:
                                                for fi149 := 0; fi149 < 20; fi149++ {
                                                    if ("" == (("" + "hello") + "hello")) {
                                                        var v150 uint32 = 65535
                                                        switch ((("foo bar" + "") + (("foo bar" + (("" + ("a" + "\n\t")) + "foo bar")) + "")) + "foo bar") {
                                                        case "a":
                                                            v151 := GenS14{F0: "hello", F1: (89 >= 162)}
                                                            _ = v151
                                                        case "hello":
                                                            switch (fi149 - int(v2)) {
                                                            case -983:
                                                                fmt.Println(v150, v5)
                                                                fmt.Println(v150)
                                                            case 32767:
                                                                if true {
                                                                    var v152 uint64 = 1
                                                                    v153 := GenS19{F0: 18446744073709551615, F1: "hello"}
                                                                    delete(v4, ("hello" + ("a" + "\n\t")))
                                                                    _ = v152
                                                                    _ = v153
                                                                } else {
                                                                    v154 := GenS9{F0: 42}
                                                                    v46 = v46
                                                                    switch 0 {
                                                                    case -947:
                                                                        var v155 int32 = 203
                                                                        _ = v155
                                                                    default:
                                                                        v46 -= 0.0
                                                                    }
                                                                    _ = v154
                                                                }
                                                            case 1:
                                                                v156 := GenS9{F0: v1}
                                                                _ = v156
                                                            }
                                                        case "foo bar":
                                                            if (v55 != (1.5 - 3.14159)) {
                                                                var v157 float32 = 1e10
                                                                _ = len(v4)
                                                                fmt.Println(v46)
                                                                _ = v157
                                                            }
                                                        }
                                                        _ = v150
                                                    }
                                                    fmt.Println(v46)
                                                    fmt.Println(v55, fi149)
                                                    _ = fi149
                                                }
                                            default:
                                                fmt.Println(v46, v46, v1)
                                            }
                                        } else {
                                            switch ("" + "foo bar") {
                                            case "":
                                                var v158 int16 = v5
                                                _ = v158
                                            default:
                                                if true {
                                                    var v159 int8 = 39
                                                    v160 := v4["hello"]
                                                    _ = v159
                                                    _ = v160
                                                }
                                            }
                                        }
                                        v55 = (v55 * float32(v46))
                                    }
                                    _ = v55
                                }
                                var v161 float32 = 0.1
                                _ = v161
                            }
                            _ = fi47
                        }
                    } else {
                        v162 := GenS10{F0: 2147483647, F1: true}
                        v163 := GenS8{F0: 1, F1: "a", F2: 2147483648}
                        if (v46 > 0.0) {
                            fmt.Println(v46, v1, v1)
                            var v164 string = "\n\t"
                            _ = v164
                        } else {
                            for fi165 := 0; fi165 < 3; fi165++ {
                                v166 := GenS19{F0: 1, F1: ""}
                                for fi167 := 0; fi167 < 19; fi167++ {
                                    switch fi165 {
                                    case 865:
                                        var v168 int = fi167
                                        _ = v168
                                    default:
                                        var v169 float32 = 1e10
                                        _ = v169
                                    }
                                    switch 4294967295 {
                                    case 9223372036854775807:
                                        v46 = v46
                                    case 32767:
                                        v2 = (v2 & v2)
                                    case 656:
                                        if false {
                                            v46 += (v46 * v46)
                                        } else {
                                            var v170 string = (((("\n\t" + "foo bar") + "") + "a") + (("hello" + "hello") + ("\n\t" + "a")))
                                            v1 -= v1
                                            var v171 uint16 = v1
                                            v172 := GenS24{F0: 1, F1: "foo bar"}
                                            _ = v170
                                            _ = v171
                                            _ = v172
                                        }
                                    }
                                    v46 += v46
                                    v173 := GenS3{F0: 2147483647}
                                    v174 := GenS25{F0: ("" + "foo bar"), F1: ((v1 & v1) >= (v1 + (v1 ^ (v1 % ((2) | 1))))), F2: -858}
                                    if (9223372036854775807 != (v2 << ((4294967295) & 63))) {
                                        switch (-(fi167)) {
                                        case 9223372036854775807:
                                            var v175 bool = (32767 <= 1000)
                                            fmt.Println(v2, v175, v46)
                                            switch "a" {
                                            case "":
                                                fmt.Println(v5, fi165, v175)
                                            case "hello":
                                                var v176 uint64 = 1
                                                _ = v176
                                            }
                                            _ = v175
                                        case 42:
                                            var v177 int = -1
                                            _ = v177
                                        }
                                        v5 = (-(v5))
                                    }
                                    _ = fi167
                                    _ = v173
                                    _ = v174
                                }
                                delete(v4, "\n\t")
                                _ = fi165
                                _ = v166
                            }
                            switch "hello" {
                            case "":
                                var v178 int16 = genFn0(2, 127, 42)
                                _ = v178
                            }
                            for fi179 := 0; fi179 < 9; fi179++ {
                                var v180 string = "foo bar"
                                v1 -= (v1 + (v1 & v1))
                                var v181 uint64 = 1
                                _ = fi179
                                _ = v180
                                _ = v181
                            }
                        }
                        _ = v162
                        _ = v163
                    }
                    v182 := GenS21{F0: (!("" != ("\n\t" + "a"))), F1: 42, F2: (true && ("\n\t" >= "hello"))}
                    _ = v182
                    _ = v46
                } else {
                    _ = len(v4)
                    delete(v4, "hello")
                    v2 += v2
                }
            } else {
                var v183 int = 32767
                v1 += 971
                _ = v183
            }
            v184 := GenS26{F0: 2147483647, F1: 679, F2: (-(v2))}
            _ = v184
        }
    }
    switch "hello" {
    case "\n\t":
        if (-7 > 0) {
            v185 := v4[""]
            fmt.Println(v1, v5)
            _ = v185
        } else {
            v5 += v5
            var v186 int8 = -128
            v1 -= (v1 / ((2) | 1))
            _ = v186
        }
    case "hello":
        v187 := GenS7{F0: 678, F1: v1}
        _ = v187
    case "":
        fmt.Println(v1, v1)
    }
    for fi188 := 0; fi188 < 14; fi188++ {
        switch "\n\t" {
        case "":
            v1 += v1
        case "hello":
            switch "" {
            case "hello":
                var v189 int8 = -74
                _ = v189
            case "foo bar":
                if len(v3) > 0 {
                    v3[0] = (("a" + "\n\t") + "foo bar")
                }
            case "\n\t":
                for fi190 := 0; fi190 < 6; fi190++ {
                    fmt.Println(v1)
                    if false {
                        var v191 int = fi190
                        if (true && false) {
                            var v192 int64 = v2
                            v193 := GenS4{F0: "\n\t"}
                            _ = v192
                            _ = v193
                        } else {
                            if (!false) {
                                v1 -= v1
                                var v194 uint16 = (v1 * (^(v1)))
                                if ((v2 % ((v2) | 1)) > -9223372036854775808) {
                                    v195 := GenS26{F0: 1, F1: 178, F2: (^(v2))}
                                    v3 = append(v3, "foo bar")
                                    if (!true) {
                                        v196 := GenS27{F0: 128}
                                        if (true || false) {
                                            v191 = 9223372036854775807
                                            v197 := GenS28{F0: 1, F1: 86}
                                            var v198 bool = (false && (-2.25 < 1.5))
                                            _ = v197
                                            _ = v198
                                        }
                                        _ = v196
                                    }
                                    v199 := GenS13{F0: 298, F1: (^(v194))}
                                    _ = v195
                                    _ = v199
                                }
                                _ = v194
                            }
                            v200 := GenS2{F0: 2147483647, F1: (v5 | -32768)}
                            if len(v3) > 0 {
                                v3[0] = "foo bar"
                            }
                            _ = v200
                        }
                        switch 32767 {
                        case 0:
                            if true {
                                var v201 int64 = (v2 & 9223372036854775807)
                                if (true && true) {
                                    v5 -= v5
                                }
                                _ = v201
                            }
                        case 216:
                            switch "a" {
                            case "hello":
                                v202 := GenS16{F0: 2147483647}
                                _ = v202
                            default:
                                var v203 uint = 256
                                _ = v203
                            }
                        }
                        _ = v191
                    } else {
                        var v204 uint = 18446744073709551615
                        _ = v204
                    }
                    _ = fi190
                }
            default:
                v1 -= (v1 * v1)
            }
        case "\n\t":
            if (!(v2 < (v2 * (v2 * 42)))) {
                var v205 int32 = -2147483648
                v206 := GenS29{F0: int8(v205), F1: v5, F2: (-(v2))}
                _ = v205
                _ = v206
            }
        }
        var v207 string = "foo bar"
        v1 = v1
        _ = fi188
        _ = v207
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
