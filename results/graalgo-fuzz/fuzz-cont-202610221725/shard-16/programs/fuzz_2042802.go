package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 bool; F1 int64; F2 int64 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool; F1 int8 }
type GenS7 struct { F0 uint8; F1 string }
type GenS8 struct { F0 uint64; F1 uint8; F2 bool }
type GenS9 struct { F0 uint; F1 int64; F2 bool }
type GenS10 struct { F0 int16; F1 uint64 }
type GenS11 struct { F0 uint32; F1 uint64 }
type GenS12 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS13 struct { F0 uint16; F1 int8 }
type GenS14 struct { F0 bool; F1 uint16 }
type GenS15 struct { F0 int; F1 int32; F2 uint16 }
type GenS16 struct { F0 uint16; F1 int32 }
type GenS17 struct { F0 uint64 }
type GenS18 struct { F0 int16 }
type GenS19 struct { F0 uint8; F1 int8 }
type GenS20 struct { F0 string }
func genFn0(p0 int, p1 int64) uint32 {
    if true {
        p0 = 0
    }
    var v4 int16 = 42
    switch "a" {
    case "foo bar":
        v4 = (^(v4))
    default:
        v5 := GenS1{F0: -287, F1: 2}
        _ = v5
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: 127, F1: 0}
        for fi8 := 0; fi8 < 16; fi8++ {
            v4 = -446
            v9 := []int{(p0 + -495), (^(fi8)), (^(fi6))}
            var v10 int16 = (v4 + (^(v4)))
            fmt.Println(v10)
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := GenS2{F0: v4, F1: 485}
        _ = fi6
        _ = v11
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    return 0
}
func genFn1(p0 int32) uint32 {
    var v12 float32 = 0.1
    for fi13 := 0; fi13 < 14; fi13++ {
        v12 = v12
        v14 := []int{(fi13 << ((18446744073709551615) & 63)), (fi13 << ((2) & 63)), (fi13 & fi13), (^(fi13))}
        _ = fi13
        _ = v14
    }
    v15 := GenS3{F0: (18446744073709551615 >= 0), F1: 812, F2: -100}
    p0 += p0
    fmt.Println(v12, v12)
    _ = p0
    _ = v12
    _ = v15
    return 2
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[string]int)
    v3 := []string{"foo bar"}
    switch 0 {
    case 127:
        for fi16 := 0; fi16 < 13; fi16++ {
            v1 = v1
            _ = fi16
        }
    case 128:
        if (18446744073709551615 > 0) {
            switch 32767 {
            case 42:
                var v17 int = 128
                _ = v17
            case -991:
                var v18 uint32 = 507
                _ = v18
            }
        }
    default:
        v1 -= 0.1
    }
    v2["hello"] = 436
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    switch 42 {
    case 313:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 32768:
        switch (((("a" + "") + "\n\t") + "") + "hello") {
        case "":
            v1 = -2.25
        case "hello":
            var v19 int32 = 0
            _ = v19
        default:
            for fi20 := 0; fi20 < 6; fi20++ {
                fmt.Println(fi20, v1)
                switch fi20 {
                case -243:
                    if (!(!true)) {
                        var v21 int64 = -1
                        _ = v21
                    } else {
                        v22 := GenS4{F0: -23, F1: 929}
                        if (false && false) {
                            v2[""] = fi20
                            if (!(!true)) {
                                var v23 bool = false
                                v23 = (!(!(0 >= 164)))
                                var v24 float32 = (-2.25 * v1)
                                _ = v23
                                _ = v24
                            }
                            var v25 uint = 4294967295
                            _ = v25
                        }
                        _ = v22
                    }
                }
                if (-232 > 1) {
                    var v26 float64 = 1e10
                    switch ("hello" + ("a" + ((("" + "hello") + "a") + "\n\t"))) {
                    case "hello":
                        v26 -= (v26 / v26)
                    case "":
                        var v27 uint8 = 127
                        _ = v27
                    case "\n\t":
                        switch "" {
                        case "hello":
                            var v28 string = (("foo bar" + "a") + "")
                            _ = v28
                        case "foo bar":
                            fmt.Println(v26)
                        case "\n\t":
                            if true {
                                var v29 int8 = 100
                                for fi30 := 0; fi30 < 2; fi30++ {
                                    if (32767 >= 591) {
                                        if len(v3) > 0 {
                                            v3[0] = ("\n\t" + "")
                                        }
                                        var v31 uint8 = 118
                                        _ = v31
                                    } else {
                                        v32 := GenS2{F0: 42, F1: 823}
                                        if true {
                                            v26 += 1.5
                                            switch ("foo bar" + ("a" + ("\n\t" + "foo bar"))) {
                                            case "":
                                                if (((1 <= 989) || false) && false) {
                                                    if (42 != -1) {
                                                        v1 = v1
                                                    }
                                                    var v33 int32 = -884
                                                    var v34 uint32 = 267
                                                    _ = v33
                                                    _ = v34
                                                } else {
                                                    v3 = append(v3, "foo bar")
                                                    v35 := GenS5{F0: 42}
                                                    v26 = (v26 + -2.25)
                                                    _ = v35
                                                }
                                            }
                                            var v36 uint16 = 1
                                            v37 := GenS1{F0: fi30, F1: 255}
                                            _ = v36
                                            _ = v37
                                        } else {
                                            if false {
                                                if true {
                                                    if (true && ((false && (!(!(!true)))) || ((float32(v26) - 0.1) <= (v1 + v1)))) {
                                                        var v38 int64 = -9223372036854775808
                                                        _ = v38
                                                    }
                                                    v29 -= (v29 >> ((255) & 63))
                                                    v39 := GenS4{F0: v29, F1: 2}
                                                    _ = v39
                                                } else {
                                                    fmt.Println(v29)
                                                    if (false && (166 < (-(fi30)))) {
                                                        v40 := GenS6{F0: true, F1: (v29 / ((v29) | 1))}
                                                        switch "foo bar" {
                                                        case "hello":
                                                            v41 := GenS5{F0: int64(fi20)}
                                                            _ = v41
                                                        }
                                                        _ = v40
                                                    }
                                                }
                                                var v42 uint8 = 220
                                                var v43 uint8 = (v42 * v42)
                                                _ = v42
                                                _ = v43
                                            } else {
                                                if ((149 <= 2147483647) || (196 > 255)) {
                                                    var v44 int8 = v29
                                                    v29 = v44
                                                    _ = v44
                                                } else {
                                                    var v45 int64 = -303
                                                    v1 += 1e10
                                                    _ = v45
                                                }
                                                v26 = ((-2.25 + ((v26 - 0.1) - (v26 / 1.0))) + float64(((v1 + (-2.25 * (-2.25 * 1e10))) / float32(1.5))))
                                                v46 := GenS7{F0: 255, F1: "hello"}
                                                _ = v46
                                            }
                                            var v47 float32 = (v1 / v1)
                                            switch fi30 {
                                            case 4294967295:
                                                v29 += (v29 ^ 42)
                                            }
                                            _ = v47
                                        }
                                        fmt.Println(fi30, v29, fi20)
                                        _ = v32
                                    }
                                    _ = fi30
                                }
                                if true {
                                    _ = len(v2)
                                    var v48 uint8 = 82
                                    if true {
                                        if (18446744073709551615 != 2) {
                                            v49 := GenS2{F0: 1000, F1: 1}
                                            _ = v49
                                        } else {
                                            for fi50 := 0; fi50 < 2; fi50++ {
                                                v26 = (((((float64(v1) * v26) * v26) + 0.0) / (v26 + 0.0)) - float64((v1 + -2.25)))
                                                _ = fi50
                                            }
                                        }
                                        var v51 uint32 = 272
                                        var v52 int32 = 2147483647
                                        _ = v51
                                        _ = v52
                                    } else {
                                        v53 := GenS1{F0: fi20, F1: 127}
                                        v26 = float64((v1 - v1))
                                        _ = v53
                                    }
                                    _ = v48
                                }
                                _ = v29
                            } else {
                                v54 := GenS6{F0: true, F1: 1}
                                for fi55 := 0; fi55 < 13; fi55++ {
                                    var v56 float64 = (-2.25 + v26)
                                    switch fi20 {
                                    case 0:
                                        switch "" {
                                        case "":
                                            var v57 string = ("" + "")
                                            _ = v57
                                        case "\n\t":
                                            v56 = -2.25
                                        default:
                                            if false {
                                                var v58 bool = (9223372036854775807 <= 181)
                                                fmt.Println(v56)
                                                _ = v58
                                            }
                                        }
                                    case 128:
                                        var v59 uint8 = 19
                                        _ = v59
                                    default:
                                        v26 += v56
                                    }
                                    _ = fi55
                                    _ = v56
                                }
                                _ = v54
                            }
                        default:
                            if (137 >= 4294967295) {
                                if true {
                                    var v60 int = fi20
                                    if (true && (9223372036854775807 != 343)) {
                                        v61 := GenS3{F0: false, F1: 447, F2: 42}
                                        v1 -= 3.14159
                                        for fi62 := 0; fi62 < 17; fi62++ {
                                            if len(v3) > 0 {
                                                v3[0] = ""
                                            }
                                            _ = fi62
                                        }
                                        _ = v61
                                    }
                                    for fi63 := 0; fi63 < 14; fi63++ {
                                        var v64 uint8 = 2
                                        _ = fi63
                                        _ = v64
                                    }
                                    _ = v60
                                }
                                for fi65 := 0; fi65 < 9; fi65++ {
                                    if (false || true) {
                                        v66 := GenS8{F0: 208, F1: 104, F2: (1 >= 0)}
                                        v1 += v1
                                        _ = v66
                                    } else {
                                        if true {
                                            var v67 bool = false
                                            if v67 {
                                                var v68 int = fi20
                                                if false {
                                                    var v69 uint = 255
                                                    _ = v69
                                                }
                                                v70 := GenS9{F0: 281, F1: -964, F2: v67}
                                                _ = v68
                                                _ = v70
                                            } else {
                                                v26 -= ((v26 - 1e10) - v26)
                                            }
                                            if v67 {
                                                if false {
                                                    var v71 uint32 = 4294967295
                                                    _ = v71
                                                }
                                                fmt.Println(fi20, fi20, v26)
                                            }
                                            _ = v67
                                        } else {
                                            if true {
                                                var v72 float64 = v26
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                switch fi20 {
                                                case 127:
                                                    var v73 bool = false
                                                    _ = v73
                                                }
                                                _ = v72
                                            }
                                            if (!(42 != 322)) {
                                                v26 = v26
                                                v2[("a" + "\n\t")] = -636
                                            }
                                            if true {
                                                v1 += float32(((v26 / v26) * v26))
                                                var v74 uint8 = 0
                                                v26 += float64((v1 / 1.0))
                                                _ = v74
                                            }
                                        }
                                        var v75 int32 = -1
                                        _ = v75
                                    }
                                    v1 -= v1
                                    _ = fi65
                                }
                            } else {
                                switch "" {
                                case "\n\t":
                                    switch "foo bar" {
                                    case "foo bar":
                                        if true {
                                            if (2147483647 != 29) {
                                                fmt.Println(v26, v26)
                                            } else {
                                                fmt.Println(v1)
                                                var v76 float32 = 1.5
                                                _ = v76
                                            }
                                            v3 = append(v3, ((("foo bar" + "a") + "\n\t") + ("a" + ((((("foo bar" + (("a" + "foo bar") + ("a" + "foo bar"))) + "") + "hello") + "hello") + ""))))
                                        } else {
                                            v77 := GenS10{F0: 1000, F1: 18446744073709551615}
                                            _ = len(v2)
                                            _ = v77
                                        }
                                    case "":
                                        var v78 uint8 = 42
                                        _ = v78
                                    }
                                case "hello":
                                    for fi79 := 0; fi79 < 19; fi79++ {
                                        v3 = append(v3, ("\n\t" + "\n\t"))
                                        if (-1 < -506) {
                                            switch (fi79 | (fi79 & fi20)) {
                                            case 492:
                                                fmt.Println(v1)
                                            default:
                                                fmt.Println(fi20)
                                            }
                                        } else {
                                            switch ("a" + ("a" + ("foo bar" + ("foo bar" + "\n\t")))) {
                                            case "":
                                                var v80 uint16 = 1
                                                _ = v80
                                            case "foo bar":
                                                switch (-(fi79)) {
                                                case -2147483648:
                                                    switch "a" {
                                                    case "\n\t":
                                                        v81 := GenS11{F0: 236, F1: 0}
                                                        _ = v81
                                                    case "hello":
                                                        v1 += v1
                                                    case "":
                                                        if true {
                                                            if false {
                                                                fmt.Println(fi20, fi20, v26)
                                                                v1 += v1
                                                                v82 := GenS12{F0: 127, F1: 65535, F2: 771}
                                                                var v83 int64 = -870
                                                                if false {
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        var v84 uint = 874
                                                                        _ = v84
                                                                    case "":
                                                                        if (fi20 == fi79) {
                                                                            if false {
                                                                                if false {
                                                                                    var v85 int8 = 1
                                                                                    if false {
                                                                                        var v86 int16 = 1000
                                                                                        v87 := GenS6{F0: (v83 < int64(v85)), F1: v85}
                                                                                        _ = v86
                                                                                        _ = v87
                                                                                    }
                                                                                    var v88 int64 = v83
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = "a"
                                                                                    }
                                                                                    v89 := GenS13{F0: 1, F1: 98}
                                                                                    _ = v85
                                                                                    _ = v88
                                                                                    _ = v89
                                                                                }
                                                                                switch "foo bar" {
                                                                                case "a":
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = (("\n\t" + ("\n\t" + (("" + "a") + ("a" + "foo bar")))) + "hello")
                                                                                    }
                                                                                default:
                                                                                    v3 = append(v3, ("hello" + ("\n\t" + ("hello" + (((("\n\t" + "a") + ("\n\t" + "foo bar")) + "\n\t") + ("a" + (("foo bar" + "\n\t") + "\n\t")))))))
                                                                                }
                                                                            } else {
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                var v90 int64 = -9223372036854775808
                                                                                _ = v90
                                                                            }
                                                                            v83 = v83
                                                                            var v91 int8 = -128
                                                                            _ = v91
                                                                        }
                                                                    case "\n\t":
                                                                        v1 += (1e10 + (-2.25 - (((v1 * v1) + v1) / v1)))
                                                                        fmt.Println(v1)
                                                                    }
                                                                    if false {
                                                                        v92 := GenS6{F0: (false || false), F1: -26}
                                                                        v83 = v83
                                                                        v83 = v83
                                                                        _ = v92
                                                                    } else {
                                                                        switch fi79 {
                                                                        case -9223372036854775808:
                                                                            v3 = append(v3, "foo bar")
                                                                        case 141:
                                                                            var v93 bool = false
                                                                            _ = v93
                                                                        }
                                                                    }
                                                                } else {
                                                                    if false {
                                                                        if (false || ((-32768 != -32768) && false)) {
                                                                            var v94 uint32 = 2
                                                                            _ = len(v3)
                                                                            _ = v94
                                                                        } else {
                                                                            var v95 string = ("foo bar" + "foo bar")
                                                                            var v96 int = fi20
                                                                            _ = v95
                                                                            _ = v96
                                                                        }
                                                                        var v97 float32 = v1
                                                                        _ = v97
                                                                    }
                                                                }
                                                                var v98 int8 = 100
                                                                var v99 uint8 = 249
                                                                _ = v82
                                                                _ = v83
                                                                _ = v98
                                                                _ = v99
                                                            }
                                                        }
                                                    default:
                                                        switch -538 {
                                                        case 2147483647:
                                                            fmt.Println(fi79)
                                                        default:
                                                            fmt.Println(v26)
                                                        }
                                                    }
                                                case 127:
                                                    if false {
                                                        fmt.Println(fi79, fi79)
                                                        fmt.Println(v26, v26)
                                                        fmt.Println(fi20, fi79, fi20)
                                                        fmt.Println(v1, v26, v1)
                                                    }
                                                case 191:
                                                    if false {
                                                        v1 = float32(3.14159)
                                                    } else {
                                                        var v100 uint32 = 2
                                                        v26 = ((v26 + v26) + v26)
                                                        var v101 int32 = 361
                                                        _ = v100
                                                        _ = v101
                                                    }
                                                }
                                            }
                                            fmt.Println(v26, fi20)
                                        }
                                        if (-1 != 127) {
                                            if ((("\n\t" + ((("hello" + (("hello" + "hello") + "")) + "hello") + ("a" + "hello"))) + ("foo bar" + "a")) < "") {
                                                if false {
                                                    v102 := GenS14{F0: false, F1: 255}
                                                    if true {
                                                        switch 42 {
                                                        case -9223372036854775808:
                                                            delete(v2, ("foo bar" + "foo bar"))
                                                        case 0:
                                                            fmt.Println(v26, fi20)
                                                        default:
                                                            fmt.Println(v26)
                                                        }
                                                        var v103 int16 = 178
                                                        _ = v103
                                                    }
                                                    var v104 uint = 42
                                                    _ = v102
                                                    _ = v104
                                                } else {
                                                    if (255 < 65535) {
                                                        var v105 uint8 = 231
                                                        switch fi20 {
                                                        case 32768:
                                                            var v106 uint64 = 0
                                                            _ = v106
                                                        case 0:
                                                            v26 -= ((v26 - v26) / (v26 * v26))
                                                            var v107 bool = (float64(1.5) < (v26 - v26))
                                                            _ = v107
                                                        case 1:
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            v1 = v1
                                                            v1 = v1
                                                        default:
                                                            v108 := GenS10{F0: 32767, F1: 2}
                                                            _ = v108
                                                        }
                                                        if (42 == -1) {
                                                            v3 = append(v3, (("\n\t" + (("a" + "") + "hello")) + ""))
                                                            if false {
                                                                v109 := GenS2{F0: 225, F1: 2}
                                                                _ = len(v2)
                                                                var v110 string = "hello"
                                                                _ = v109
                                                                _ = v110
                                                            }
                                                            var v111 int16 = 32767
                                                            delete(v2, "")
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            _ = v111
                                                        } else {
                                                            if false {
                                                                var v112 uint8 = v105
                                                                switch (fi20 & (-(fi79))) {
                                                                case 32768:
                                                                    var v113 uint32 = 886
                                                                    _ = v113
                                                                case -1:
                                                                    if ("foo bar" >= "a") {
                                                                        if (((v112 <= (v105 ^ v112)) && true) || false) {
                                                                            fmt.Println(v105, v1, fi20)
                                                                            delete(v2, "\n\t")
                                                                            if (!true) {
                                                                                v114 := GenS15{F0: (-(fi79)), F1: -561, F2: 0}
                                                                                var v115 int32 = 0
                                                                                _ = v114
                                                                                _ = v115
                                                                            }
                                                                        }
                                                                        fmt.Println(v112, v112, v105)
                                                                    }
                                                                }
                                                                if (!true) {
                                                                    fmt.Println(fi20, fi79)
                                                                    var v116 int16 = 32767
                                                                    var v117 int32 = 140
                                                                    _ = v116
                                                                    _ = v117
                                                                } else {
                                                                    var v118 uint = 18446744073709551615
                                                                    switch 978 {
                                                                    case 986:
                                                                        if false {
                                                                            fmt.Println(v105)
                                                                        } else {
                                                                            if false {
                                                                                v119 := GenS16{F0: 65535, F1: 1}
                                                                                _ = v119
                                                                            } else {
                                                                                if (1 <= 65) {
                                                                                    if true {
                                                                                        fmt.Println(v112, v112, v26)
                                                                                        if false {
                                                                                            if (520 < 0) {
                                                                                                var v120 bool = (("" + ("" + ("foo bar" + "a"))) == ("\n\t" + ""))
                                                                                                var v121 uint32 = genFn1(-912)
                                                                                                v122 := v2["foo bar"]
                                                                                                _ = v120
                                                                                                _ = v121
                                                                                                _ = v122
                                                                                            }
                                                                                        }
                                                                                        _ = len(v3)
                                                                                    }
                                                                                    v123 := GenS1{F0: (fi20 >> (((v118 % (((v118 + 812)) | 1))) & 63)), F1: 194}
                                                                                    _ = v123
                                                                                } else {
                                                                                    if (!(37 <= 2)) {
                                                                                        v124 := GenS13{F0: 42, F1: 100}
                                                                                        _ = v124
                                                                                    }
                                                                                    v125 := v2[""]
                                                                                    _ = v125
                                                                                }
                                                                                v126 := GenS16{F0: 137, F1: 0}
                                                                                switch (-(fi79)) {
                                                                                case -1:
                                                                                    v127 := GenS12{F0: v105, F1: v118, F2: 4294967295}
                                                                                    _ = v127
                                                                                case 881:
                                                                                    fmt.Println(v118, v26, v118)
                                                                                }
                                                                                fmt.Println(v118)
                                                                                _ = v126
                                                                            }
                                                                        }
                                                                    }
                                                                    switch -448 {
                                                                    case 4294967295:
                                                                        fmt.Println(v105)
                                                                    case -9223372036854775808:
                                                                        if len(v3) > 0 {
                                                                            v3[0] = ""
                                                                        }
                                                                    case -99:
                                                                        var v128 float32 = ((float32(1.5) / ((v1 + 1e10) * float32(((3.14159 * (1e10 - v26)) / v26)))) + float32(v26))
                                                                        _ = v128
                                                                    default:
                                                                        var v129 uint16 = 559
                                                                        _ = v129
                                                                    }
                                                                    _ = v118
                                                                }
                                                                _ = v112
                                                            }
                                                            fmt.Println(fi20, fi20, v105)
                                                        }
                                                        _ = v105
                                                    } else {
                                                        fmt.Println(fi20, v1)
                                                        switch (fi79 - (fi79 >> ((18446744073709551615) & 63))) {
                                                        case -697:
                                                            switch 1 {
                                                            case -248:
                                                                switch "\n\t" {
                                                                case "a":
                                                                    v1 += (((0.0 - v1) * v1) / ((v1 - v1) / v1))
                                                                case "\n\t":
                                                                    if ((-2.25 * 1.5) < float64(float32(float64(((v1 / (1.5 + v1)) * (v1 + float32((3.14159 / (0.1 + ((((v26 + float64((v1 * (v1 - 0.1)))) + (1e10 + 0.0)) + v26) - ((v26 + v26) / 1.5))))))))))) {
                                                                        switch (fi20 | (fi20 - (fi20 >> ((65535) & 63)))) {
                                                                        case -500:
                                                                            v1 = v1
                                                                        case 42:
                                                                            _ = len(v2)
                                                                        }
                                                                    } else {
                                                                        var v130 uint64 = 2
                                                                        if false {
                                                                            v1 += 1.5
                                                                        } else {
                                                                            delete(v2, "hello")
                                                                            if (26 < 1) {
                                                                                if true {
                                                                                    switch "foo bar" {
                                                                                    case "":
                                                                                        var v131 uint8 = 128
                                                                                        _ = v131
                                                                                    case "foo bar":
                                                                                        v132 := v2["a"]
                                                                                        _ = v132
                                                                                    case "hello":
                                                                                        v2["a"] = fi79
                                                                                    }
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                    v133 := GenS3{F0: (!(1 != 32767)), F1: 0, F2: 9223372036854775807}
                                                                                    _ = v133
                                                                                } else {
                                                                                    _ = len(v3)
                                                                                }
                                                                                if true {
                                                                                    if (-128 <= 1) {
                                                                                        var v134 uint8 = 128
                                                                                        _ = v134
                                                                                    }
                                                                                    switch ("foo bar" + "\n\t") {
                                                                                    case "\n\t":
                                                                                        fmt.Println(v1)
                                                                                    case "a":
                                                                                        v135 := GenS17{F0: v130}
                                                                                        v136 := v2["a"]
                                                                                        _ = v135
                                                                                        _ = v136
                                                                                    case "foo bar":
                                                                                        if ((uint(v130) > 1) || false) {
                                                                                            if (127 > 0) {
                                                                                                fmt.Println(v26, fi79, v26)
                                                                                                v1 = -2.25
                                                                                            } else {
                                                                                                v1 = (v1 / float32((-2.25 * ((v26 * v26) + -2.25))))
                                                                                                switch "a" {
                                                                                                case "a":
                                                                                                    v137 := GenS8{F0: (v130 ^ 2), F1: 42, F2: true}
                                                                                                    _ = v137
                                                                                                case "":
                                                                                                    v138 := GenS18{F0: 909}
                                                                                                    _ = v138
                                                                                                case "hello":
                                                                                                    fmt.Println(v26, v26, v1)
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        v2["a"] = -2147483648
                                                                                    }
                                                                                    v130 = (^(v130))
                                                                                }
                                                                            }
                                                                            var v139 int32 = -1
                                                                            _ = v139
                                                                        }
                                                                        _ = v130
                                                                    }
                                                                }
                                                            case 0:
                                                                v140 := GenS14{F0: false, F1: 65535}
                                                                _ = v140
                                                            default:
                                                                switch ("" + ("foo bar" + ("\n\t" + ("foo bar" + "")))) {
                                                                case "foo bar":
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                case "":
                                                                    v26 += ((1.5 - v26) + 3.14159)
                                                                case "a":
                                                                    if (18446744073709551615 >= 2) {
                                                                        var v141 bool = true
                                                                        if true {
                                                                            v2[("hello" + "foo bar")] = fi79
                                                                            v1 += v1
                                                                            var v142 uint16 = 42
                                                                            _ = v142
                                                                        } else {
                                                                            var v143 int32 = -2147483648
                                                                            _ = v143
                                                                        }
                                                                        v1 = 3.14159
                                                                        _ = v141
                                                                    } else {
                                                                        v26 -= (((1e10 - 0.0) - ((v26 * 3.14159) * v26)) - v26)
                                                                        _ = len(v2)
                                                                        var v144 bool = false
                                                                        _ = v144
                                                                    }
                                                                default:
                                                                    v145 := GenS7{F0: 168, F1: "a"}
                                                                    var v146 uint32 = 2147483648
                                                                    v147 := GenS14{F0: ((281 != 52) && (("" + "\n\t") > ((("\n\t" + "\n\t") + "a") + (("foo bar" + ("\n\t" + "\n\t")) + "foo bar")))), F1: uint16(v146)}
                                                                    _ = v145
                                                                    _ = v146
                                                                    _ = v147
                                                                }
                                                            }
                                                        case -1:
                                                            v148 := GenS19{F0: 127, F1: 1}
                                                            _ = v148
                                                        case 158:
                                                            var v149 int = 32767
                                                            _ = v149
                                                        default:
                                                            if false {
                                                                var v150 int = (fi20 | (^(fi20)))
                                                                fmt.Println(fi79, v150)
                                                                v1 += float32(1.5)
                                                                _ = v150
                                                            }
                                                        }
                                                    }
                                                    switch "a" {
                                                    case "foo bar":
                                                        var v151 uint16 = 0
                                                        _ = v151
                                                    case "":
                                                        if (0 <= 42) {
                                                            v26 += v26
                                                        } else {
                                                            var v152 int64 = -773
                                                            _ = v152
                                                        }
                                                    case "a":
                                                        v26 = ((v26 - v26) / v26)
                                                    default:
                                                        if false {
                                                            var v153 string = (("" + ("" + ((("foo bar" + "hello") + "\n\t") + (("foo bar" + "") + "foo bar")))) + "")
                                                            _ = v153
                                                        }
                                                    }
                                                }
                                                var v154 int16 = 32767
                                                _ = len(v2)
                                                _ = v154
                                            }
                                            var v155 string = "\n\t"
                                            _ = v155
                                        } else {
                                            fmt.Println(v26)
                                            fmt.Println(fi79, fi79)
                                        }
                                        _ = fi79
                                    }
                                case "a":
                                    _ = len(v3)
                                default:
                                    v1 += 3.14159
                                }
                                v2["foo bar"] = (fi20 ^ fi20)
                            }
                        }
                    }
                    _ = len(v2)
                    _ = v26
                } else {
                    for fi156 := 0; fi156 < 9; fi156++ {
                        v3 = append(v3, "\n\t")
                        if true {
                            v1 = ((1e10 - v1) / -2.25)
                            v3 = append(v3, "\n\t")
                        } else {
                            v157 := GenS20{F0: (((("a" + ("hello" + "a")) + "hello") + ("" + "\n\t")) + (("" + "") + "\n\t"))}
                            if (true || (426 <= 42)) {
                                v158 := GenS12{F0: 128, F1: 42, F2: 0}
                                var v159 uint16 = 2
                                _ = v158
                                _ = v159
                            }
                            _ = len(v3)
                            _ = v157
                        }
                        var v160 uint64 = 502
                        _ = fi156
                        _ = v160
                    }
                    var v161 bool = (("foo bar" + "\n\t") == "")
                    _ = v161
                }
                _ = fi20
            }
        }
    default:
        for fi162 := 0; fi162 < 13; fi162++ {
            v1 += (float32(1.5) / (3.14159 - v1))
            _ = fi162
        }
    }
    var v163 int64 = 9223372036854775807
    if (v163 == -539) {
        v1 -= 0.0
        v2[("a" + "foo bar")] = -1
        _ = len(v2)
    }
    var v164 string = ((("a" + "foo bar") + ("foo bar" + "\n\t")) + "\n\t")
    var v165 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v163:", v163)
    fmt.Println("v164:", v164)
    fmt.Println("v165:", v165)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v163
    _ = v164
    _ = v165
    _ = v2
    _ = v3
}
