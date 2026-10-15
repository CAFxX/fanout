package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 bool; F1 bool; F2 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64; F1 int16 }
type GenS8 struct { F0 uint32; F1 bool; F2 int }
type GenS9 struct { F0 uint8; F1 string; F2 uint16 }
type GenS10 struct { F0 int; F1 int; F2 string }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint16; F1 uint16 }
type GenS13 struct { F0 uint8; F1 int32 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 string; F1 uint; F2 uint }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS18 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS19 struct { F0 string }
type GenS20 struct { F0 uint8; F1 bool }
type GenS21 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS22 struct { F0 string }
type GenS23 struct { F0 int16; F1 int64; F2 uint8 }
type GenS24 struct { F0 int16; F1 string; F2 int16 }
func genFn0() uint8 {
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 string = "hello"
        _ = fi6
        _ = v7
    }
    var v8 bool = ((2 == 2147483647) || false)
    if v8 {
        v8 = v8
        var v9 int8 = -128
        _ = v9
    } else {
        v10 := GenS1{F0: 31, F1: 0, F2: 0}
        v8 = (v8 && (9223372036854775807 != -9223372036854775808))
        _ = v10
    }
    v11 := GenS1{F0: 42, F1: 1, F2: 0}
    _ = v11
    _ = v8
    return 244
}

func main() {
    var v1 uint = 34
    var v2 int32 = -1
    v3 := []string{}
    v4 := make(map[uint32]uint32)
    var v5 uint32 = 4294967295
    if (42 > genFn0()) {
        v12 := GenS2{F0: -2147483648, F1: "foo bar"}
        if (2 < 922) {
            v5 = v5
            if true {
                for fi13 := 0; fi13 < 4; fi13++ {
                    for fi14 := 0; fi14 < 13; fi14++ {
                        var v15 int16 = 42
                        _ = fi14
                        _ = v15
                    }
                    var v16 float64 = 1e10
                    if true {
                        v5 = 925
                    } else {
                        v16 = 3.14159
                    }
                    _ = fi13
                    _ = v16
                }
            } else {
                for fi17 := 0; fi17 < 7; fi17++ {
                    v5 -= v5
                    _ = fi17
                }
                var v18 int = 32767
                _ = v18
            }
            v19 := GenS3{F0: (!false), F1: (true || false), F2: 50}
            _ = v19
        }
        _ = v12
    }
    if true {
        v20 := GenS4{F0: -95}
        v21 := GenS4{F0: -110}
        if (((("" + "foo bar") + ("\n\t" + ("hello" + ((("" + "") + "foo bar") + ("a" + (("" + "foo bar") + "")))))) + "") <= "hello") {
            var v22 int32 = 2147483647
            v1 -= v1
            var v23 int = 42
            _ = v22
            _ = v23
        }
        _ = v20
        _ = v21
    } else {
        var v24 int16 = 42
        _ = v24
    }
    if true {
        if (!(!true)) {
            v25 := GenS5{F0: 990, F1: 164, F2: (v2 >> ((2) & 63))}
            if (false && ((true || (v1 <= (v1 * v1))) || (false || (true || ((!false) && (801 > 4294967295)))))) {
                var v26 int = 128
                var v27 int16 = 1000
                v27 = v27
                _ = v26
                _ = v27
            }
            if (v2 != (v2 / ((v2) | 1))) {
                switch -824 {
                case 127:
                    v1 = v1
                }
            } else {
                fmt.Println(v2)
                for fi28 := 0; fi28 < 4; fi28++ {
                    if (!false) {
                        if true {
                            v29 := GenS3{F0: false, F1: false, F2: (fi28 / ((fi28) | 1))}
                            _ = v29
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if true {
                                if (0.0 < (float32(0.1) / 1.0)) {
                                    for fi30 := 0; fi30 < 1; fi30++ {
                                        var v31 uint64 = 42
                                        _ = fi30
                                        _ = v31
                                    }
                                    for fi32 := 0; fi32 < 14; fi32++ {
                                        var v33 uint16 = 255
                                        if len(v3) > 0 {
                                            v3[0] = "hello"
                                        }
                                        _ = fi32
                                        _ = v33
                                    }
                                } else {
                                    var v34 uint16 = 1
                                    _ = v34
                                }
                                var v35 float64 = 0.0
                                v1 = 4294967295
                                _ = v35
                            }
                            v36 := v4[v5]
                            _ = v36
                        }
                        v37 := GenS6{F0: 1}
                        _ = v37
                    } else {
                        var v38 uint = 337
                        v39 := GenS3{F0: false, F1: false, F2: fi28}
                        _ = v38
                        _ = v39
                    }
                    _ = fi28
                }
            }
            _ = v25
        } else {
            var v40 uint = (v1 + (^(v1)))
            _ = v40
        }
        if true {
            var v41 int64 = 231
            var v42 uint64 = 2
            _ = v41
            _ = v42
        } else {
            v43 := GenS4{F0: 1}
            if false {
                switch -198 {
                case 4294967295:
                    if len(v3) > 0 {
                        v3[0] = "\n\t"
                    }
                case 32767:
                    var v44 float32 = 1e10
                    _ = v44
                case 127:
                    for fi45 := 0; fi45 < 5; fi45++ {
                        var v46 int8 = -1
                        switch (fi45 | -665) {
                        case 0:
                            for fi47 := 0; fi47 < 14; fi47++ {
                                if (((1e10 - (((-2.25 * 0.0) * 3.14159) * float32(0.0))) < (0.0 - 1.5)) || (!false)) {
                                    if false {
                                        if false {
                                            v46 = int8(v2)
                                            var v48 uint8 = genFn0()
                                            v49 := GenS7{F0: 42, F1: int16(fi45)}
                                            if (264 < 1) {
                                                if (true && true) {
                                                    fmt.Println(v1, fi47)
                                                    delete(v4, v5)
                                                    _ = len(v4)
                                                } else {
                                                    switch "" {
                                                    case "a":
                                                        if false {
                                                            if true {
                                                                switch int(v48) {
                                                                case 42:
                                                                    if (!false) {
                                                                        v46 -= (v46 - v46)
                                                                        delete(v4, (v5 % (((v5 / ((v5) | 1))) | 1)))
                                                                        if len(v3) > 0 {
                                                                            v3[0] = ("foo bar" + "")
                                                                        }
                                                                        var v50 float64 = 1e10
                                                                        _ = v50
                                                                    } else {
                                                                        fmt.Println(v48, v5)
                                                                    }
                                                                    v5 += v5
                                                                    if false {
                                                                        switch ("foo bar" + "") {
                                                                        case "\n\t":
                                                                            v48 += v48
                                                                        }
                                                                        var v51 float64 = 1e10
                                                                        _ = v51
                                                                    } else {
                                                                        var v52 uint = v1
                                                                        fmt.Println(v2)
                                                                        var v53 string = "a"
                                                                        var v54 string = "\n\t"
                                                                        _ = v52
                                                                        _ = v53
                                                                        _ = v54
                                                                    }
                                                                case -1:
                                                                    fmt.Println(v2, v48, v2)
                                                                default:
                                                                    if (false && true) {
                                                                        fmt.Println(fi47)
                                                                    }
                                                                }
                                                                var v55 string = "a"
                                                                var v56 int32 = v2
                                                                _ = v55
                                                                _ = v56
                                                            } else {
                                                                v2 = v2
                                                                var v57 uint64 = 2
                                                                _ = len(v4)
                                                                _ = v57
                                                            }
                                                        }
                                                    default:
                                                        var v58 string = (("\n\t" + "") + "hello")
                                                        _ = v58
                                                    }
                                                    v59 := GenS2{F0: v2, F1: ("" + ((("" + "hello") + "hello") + "hello"))}
                                                    _ = v59
                                                }
                                                v60 := GenS8{F0: 1, F1: false, F2: fi47}
                                                _ = v60
                                            } else {
                                                v46 += (v46 >> ((v1) & 63))
                                            }
                                            if false {
                                                if (32767 > 609) {
                                                    var v61 uint64 = 18446744073709551615
                                                    var v62 bool = (32767 > 255)
                                                    _ = v61
                                                    _ = v62
                                                } else {
                                                    if (2 == 2) {
                                                        fmt.Println(v48, v2)
                                                        v63 := GenS7{F0: 2, F1: 0}
                                                        _ = v63
                                                    } else {
                                                        if true {
                                                            v4[0] = v5
                                                            var v64 uint8 = 193
                                                            fmt.Println(v64, fi45, v48)
                                                            if true {
                                                                delete(v4, uint32(v48))
                                                                v2 = -466
                                                            } else {
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    v65 := GenS1{F0: 1, F1: 840, F2: -59}
                                                                    _ = v65
                                                                case "":
                                                                    var v66 string = ("hello" + "")
                                                                    var v67 bool = (true && true)
                                                                    _ = v66
                                                                    _ = v67
                                                                case "foo bar":
                                                                    v68 := GenS9{F0: v64, F1: "a", F2: 32767}
                                                                    _ = v68
                                                                }
                                                                var v69 int = (fi45 & (fi45 << (((v1 / ((uint(v64)) | 1))) & 63)))
                                                                _ = v69
                                                            }
                                                            if len(v3) > 0 {
                                                                v3[0] = "a"
                                                            }
                                                            if false {
                                                                if true {
                                                                    if true {
                                                                        switch "hello" {
                                                                        case "a":
                                                                            v4[v5] = 436
                                                                        case "foo bar":
                                                                            if true {
                                                                                v70 := GenS9{F0: 39, F1: "\n\t", F2: 42}
                                                                                v71 := GenS10{F0: (fi45 | fi47), F1: fi47, F2: "foo bar"}
                                                                                delete(v4, 1)
                                                                                _ = v70
                                                                                _ = v71
                                                                            } else {
                                                                                v46 += v46
                                                                                var v72 uint = 2
                                                                                var v73 uint8 = genFn0()
                                                                                _ = v72
                                                                                _ = v73
                                                                            }
                                                                        case "hello":
                                                                            if len(v3) > 0 {
                                                                                v3[0] = "foo bar"
                                                                            }
                                                                        default:
                                                                            if true {
                                                                                fmt.Println(v5, v2, v1)
                                                                                v64 -= genFn0()
                                                                                if true {
                                                                                    v4[(^(v5))] = (v5 * v5)
                                                                                    var v74 int16 = 0
                                                                                    v4[v5] = 2147483648
                                                                                    v74 = v74
                                                                                    delete(v4, v5)
                                                                                    _ = v74
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v2, v46, v1)
                                                                                switch (fi45 * 2147483647) {
                                                                                case -883:
                                                                                    var v75 int32 = 1
                                                                                    _ = v75
                                                                                case 293:
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                case 127:
                                                                                    fmt.Println(v5, fi45, v48)
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = len(v3)
                                                                        fmt.Println(fi45)
                                                                    }
                                                                    var v76 int32 = -1
                                                                    if ("" >= "") {
                                                                        var v77 float32 = ((float32(3.14159) / 1e10) + 1.5)
                                                                        _ = v77
                                                                    } else {
                                                                        v78 := GenS11{F0: v1}
                                                                        v48 = (v64 >> ((v1) & 63))
                                                                        _ = v78
                                                                    }
                                                                    v3 = append(v3, (("" + ("foo bar" + ((("\n\t" + ("a" + "")) + "foo bar") + ((("foo bar" + "foo bar") + ("foo bar" + ("hello" + "foo bar"))) + "")))) + ("a" + (("hello" + "") + "a"))))
                                                                    if false {
                                                                        switch (^(fi45)) {
                                                                        case 42:
                                                                            fmt.Println(fi45)
                                                                        }
                                                                    }
                                                                    _ = v76
                                                                }
                                                                fmt.Println(fi47, v48)
                                                                v48 = v64
                                                            } else {
                                                                if true {
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    var v79 uint = v1
                                                                    switch 32767 {
                                                                    case 127:
                                                                        var v80 uint64 = 18446744073709551615
                                                                        _ = v80
                                                                    default:
                                                                        fmt.Println(v79, v46)
                                                                    }
                                                                    v64 -= (v48 / ((v64) | 1))
                                                                    _ = v79
                                                                } else {
                                                                    v48 = 1
                                                                }
                                                                v81 := GenS2{F0: -269, F1: ""}
                                                                if true {
                                                                    fmt.Println(v5, v64, v2)
                                                                }
                                                                _ = v81
                                                            }
                                                            v82 := v4[2]
                                                            _ = v64
                                                            _ = v82
                                                        }
                                                        var v83 uint32 = 2
                                                        _ = len(v3)
                                                        _ = v83
                                                    }
                                                    v48 -= v48
                                                    v46 = (-(v46))
                                                }
                                            } else {
                                                switch fi45 {
                                                case -2147483648:
                                                    var v84 uint = 17
                                                    _ = v84
                                                case 231:
                                                    var v85 string = ("a" + ("foo bar" + "\n\t"))
                                                    fmt.Println(v46, fi45, v5)
                                                    _ = v85
                                                default:
                                                    fmt.Println(v1, v48, v46)
                                                }
                                                if false {
                                                    v86 := GenS12{F0: 42, F1: 0}
                                                    switch "\n\t" {
                                                    case "a":
                                                        v87 := GenS4{F0: (v46 + (v46 >> (((^(v1))) & 63)))}
                                                        _ = v87
                                                    case "foo bar":
                                                        v5 -= (v5 / (((v5 | (v5 | (^(v5))))) | 1))
                                                    case "":
                                                        switch (("" + "hello") + "\n\t") {
                                                        case "hello":
                                                            var v88 int64 = 9223372036854775807
                                                            _ = v88
                                                        case "a":
                                                            fmt.Println(v46)
                                                        case "\n\t":
                                                            v2 -= (v2 % (((v2 >> ((v1) & 63))) | 1))
                                                        default:
                                                            switch (fi45 / (((fi47 ^ fi47)) | 1)) {
                                                            case 32767:
                                                                v89 := GenS11{F0: 4294967295}
                                                                _ = v89
                                                            case 4294967295:
                                                                v2 += v2
                                                            }
                                                        }
                                                    default:
                                                        var v90 int8 = 24
                                                        _ = v90
                                                    }
                                                    _ = v86
                                                } else {
                                                    if false {
                                                        _ = len(v3)
                                                    } else {
                                                        switch (fi45 & fi45) {
                                                        case 42:
                                                            var v91 bool = false
                                                            _ = v91
                                                        case 0:
                                                            v92 := v4[v5]
                                                            if ((59 != v1) || true) {
                                                                v92 += (v5 - (v5 / ((v5) | 1)))
                                                                var v93 int64 = 42
                                                                fmt.Println(v93, fi45)
                                                                _ = v93
                                                            } else {
                                                                v94 := GenS13{F0: (v48 | 0), F1: -359}
                                                                _ = v94
                                                            }
                                                            v1 = (^(v1))
                                                            v95 := GenS4{F0: (v46 << ((18446744073709551615) & 63))}
                                                            _ = v92
                                                            _ = v95
                                                        }
                                                        fmt.Println(v5, v2)
                                                    }
                                                }
                                            }
                                            _ = v48
                                            _ = v49
                                        }
                                        _ = len(v3)
                                        v5 = v5
                                        var v96 string = (("hello" + "\n\t") + ("foo bar" + ""))
                                        if (42 > 18446744073709551615) {
                                            v96 = v96
                                            if (!(v5 < v5)) {
                                                if (v2 > (v2 & v2)) {
                                                    v4[(^(v5))] = v5
                                                    if len(v3) > 0 {
                                                        v3[0] = ((v96 + v96) + (v96 + v96))
                                                    }
                                                }
                                                var v97 uint64 = 874
                                                _ = v97
                                            } else {
                                                var v98 uint16 = 65535
                                                var v99 string = "a"
                                                if true {
                                                    v100 := GenS2{F0: v2, F1: (v96 + ("foo bar" + v99))}
                                                    v99 = v99
                                                    v3 = append(v3, v96)
                                                    v4[v5] = (v5 & v5)
                                                    v101 := GenS3{F0: false, F1: (false || false), F2: 717}
                                                    v1 = (v1 >> (((v1 - v1)) & 63))
                                                    _ = v100
                                                    _ = v101
                                                } else {
                                                    if true {
                                                        var v102 uint64 = 0
                                                        var v103 int64 = 1
                                                        if (-1 <= (v2 ^ v2)) {
                                                            v4[v5] = v5
                                                            switch "" {
                                                            case "foo bar":
                                                                var v104 bool = true
                                                                v102 = 291
                                                                if v104 {
                                                                    switch (-(fi45)) {
                                                                    case 127:
                                                                        switch v96 {
                                                                        case "\n\t":
                                                                            switch v99 {
                                                                            case "":
                                                                                v105 := GenS14{F0: (v103 * v103)}
                                                                                _ = v105
                                                                            default:
                                                                                v106 := GenS15{F0: ("\n\t" + ""), F1: v1, F2: (v1 ^ v1)}
                                                                                _ = v106
                                                                            }
                                                                        default:
                                                                            v1 = (v1 & (^(v1)))
                                                                        }
                                                                    default:
                                                                        v107 := GenS9{F0: 2, F1: "", F2: v98}
                                                                        _ = v107
                                                                    }
                                                                    var v108 int8 = -1
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    _ = v108
                                                                }
                                                                v1 -= (v1 ^ v1)
                                                                v109 := GenS12{F0: v98, F1: (v98 & 255)}
                                                                delete(v4, v5)
                                                                _ = v104
                                                                _ = v109
                                                            default:
                                                                var v110 float64 = 1.5
                                                                var v111 bool = false
                                                                if (-61 >= 1000) {
                                                                    var v112 uint = v1
                                                                    v96 = "hello"
                                                                    var v113 int32 = 1
                                                                    _ = v112
                                                                    _ = v113
                                                                } else {
                                                                    v46 = 100
                                                                }
                                                                _ = v110
                                                                _ = v111
                                                            }
                                                            fmt.Println(fi47)
                                                        }
                                                        _ = v102
                                                        _ = v103
                                                    } else {
                                                        var v114 bool = true
                                                        _ = len(v4)
                                                        if true {
                                                            v3 = append(v3, v99)
                                                            if v114 {
                                                                delete(v4, 2147483647)
                                                                var v115 float32 = (3.14159 - float32(0.0))
                                                                _ = v115
                                                            }
                                                        } else {
                                                            fmt.Println(v99, v2, v114)
                                                            switch (v99 + (("\n\t" + v96) + "foo bar")) {
                                                            case "a":
                                                                switch fi45 {
                                                                case 128:
                                                                    var v116 uint = (^(v1))
                                                                    _ = v116
                                                                case 4294967295:
                                                                    v98 += v98
                                                                default:
                                                                    v46 = (v46 * v46)
                                                                }
                                                            case "\n\t":
                                                                var v117 bool = v114
                                                                _ = v117
                                                            }
                                                            switch fi45 {
                                                            case -1:
                                                                _ = len(v3)
                                                            case 128:
                                                                var v118 uint8 = 255
                                                                _ = v118
                                                            case -2147483648:
                                                                v1 += v1
                                                                v96 = v99
                                                                fmt.Println(v114, v98, v99)
                                                            }
                                                            fmt.Println(v98, v99)
                                                            if len(v3) > 0 {
                                                                v3[0] = "foo bar"
                                                            }
                                                        }
                                                        v1 -= v1
                                                        v1 = 0
                                                        _ = v114
                                                    }
                                                }
                                                _ = v98
                                                _ = v99
                                            }
                                        } else {
                                            if true {
                                                var v119 uint8 = genFn0()
                                                v3 = append(v3, v96)
                                                _ = v119
                                            } else {
                                                if false {
                                                    switch fi47 {
                                                    case 990:
                                                        if true {
                                                            var v120 int16 = 637
                                                            switch v96 {
                                                            case "a":
                                                                switch -9223372036854775808 {
                                                                case 9223372036854775807:
                                                                    var v121 uint = (v1 << ((v1) & 63))
                                                                    _ = v121
                                                                default:
                                                                    fmt.Println(fi45, fi45)
                                                                }
                                                            case "\n\t":
                                                                v122 := GenS11{F0: v1}
                                                                _ = v122
                                                            case "":
                                                                v46 = v46
                                                                switch v96 {
                                                                case "a":
                                                                    var v123 uint32 = v5
                                                                    _ = v123
                                                                case "hello":
                                                                    fmt.Println(v1, fi47)
                                                                case "\n\t":
                                                                    v96 = v96
                                                                }
                                                                if ((((v46 | v46) <= v46) || ((v120 >> ((v1) & 63)) > v120)) && false) {
                                                                    v124 := GenS5{F0: 2, F1: 255, F2: -246}
                                                                    v125 := GenS16{F0: 285}
                                                                    var v126 uint32 = (v5 & v5)
                                                                    _ = v124
                                                                    _ = v125
                                                                    _ = v126
                                                                }
                                                            default:
                                                                if true {
                                                                    var v127 float32 = float32(float64(3.14159))
                                                                    var v128 uint8 = 114
                                                                    _ = v127
                                                                    _ = v128
                                                                }
                                                                var v129 uint = (v1 - v1)
                                                                _ = v129
                                                            }
                                                            v96 = v96
                                                            _ = v120
                                                        }
                                                    default:
                                                        if (v46 >= 1) {
                                                            delete(v4, uint32(v1))
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                v130 := GenS17{F0: 1, F1: -1, F2: 255}
                                                                _ = v130
                                                            default:
                                                                v46 = v46
                                                                var v131 float32 = 1.5
                                                                v96 = v96
                                                                _ = v131
                                                            }
                                                            if true {
                                                                v4[v5] = v5
                                                            }
                                                        }
                                                    }
                                                    v132 := GenS9{F0: 61, F1: "\n\t", F2: 497}
                                                    switch fi45 {
                                                    case 2147483647:
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                    v133 := GenS18{F0: v2, F1: 0, F2: 221}
                                                    switch v96 {
                                                    case "foo bar":
                                                        v134 := GenS4{F0: int8(v1)}
                                                        _ = v134
                                                    case "\n\t":
                                                        fmt.Println(fi45, v46, fi45)
                                                    }
                                                    _ = v132
                                                    _ = v133
                                                }
                                                var v135 uint = 1
                                                _ = v135
                                            }
                                            v2 = v2
                                            switch -2147483648 {
                                            case 2147483647:
                                                v4[0] = v5
                                            case -9223372036854775808:
                                                var v136 int32 = 1
                                                _ = v136
                                            case -2147483648:
                                                v4[(v5 << (((^(v1))) & 63))] = (v5 ^ (v5 << ((4294967295) & 63)))
                                            }
                                        }
                                        var v137 uint64 = 122
                                        if false {
                                            if false {
                                                v96 = "hello"
                                                if ((^(v5)) <= (v5 << (((v1 % (((v1 << ((uint(v137)) & 63))) | 1))) & 63))) {
                                                    var v138 uint = (v1 & (v1 + v1))
                                                    var v139 uint = v138
                                                    v140 := GenS5{F0: 526, F1: 42, F2: v2}
                                                    _ = v138
                                                    _ = v139
                                                    _ = v140
                                                }
                                                switch "\n\t" {
                                                case "a":
                                                    var v141 uint32 = 65535
                                                    _ = v141
                                                case "foo bar":
                                                    v5 += 4294967295
                                                case "hello":
                                                    fmt.Println(v96, v1, v137)
                                                default:
                                                    var v142 int = -377
                                                    _ = v142
                                                }
                                            }
                                            if (127 < genFn0()) {
                                                var v143 bool = false
                                                var v144 uint64 = 1
                                                fmt.Println(fi45)
                                                v145 := GenS15{F0: v96, F1: v1, F2: v1}
                                                _ = v143
                                                _ = v144
                                                _ = v145
                                            } else {
                                                v4[722] = v5
                                                switch v96 {
                                                case "a":
                                                    v1 = v1
                                                }
                                                fmt.Println(v2, fi45, v5)
                                            }
                                        } else {
                                            _ = len(v3)
                                            v146 := GenS3{F0: (!(!(!(false && ((true && false) || true))))), F1: true, F2: 42}
                                            v147 := v4[v5]
                                            _ = v146
                                            _ = v147
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v137
                                        _ = v96
                                    }
                                    if false {
                                        var v148 float64 = float64(3.14159)
                                        var v149 int64 = -9223372036854775808
                                        v150 := GenS19{F0: "a"}
                                        _ = v148
                                        _ = v149
                                        _ = v150
                                    }
                                    v151 := v4[v5]
                                    _ = v151
                                }
                                _ = fi47
                            }
                        case 1:
                            switch "\n\t" {
                            case "a":
                                var v152 uint64 = 2
                                _ = v152
                            case "":
                                fmt.Println(v46, v46, v5)
                            case "hello":
                                v5 = v5
                            default:
                                switch ((("a" + "") + "foo bar") + "a") {
                                case "hello":
                                    fmt.Println(fi45)
                                default:
                                    v153 := GenS13{F0: 42, F1: v2}
                                    _ = v153
                                }
                            }
                        case 9223372036854775807:
                            v2 += (v2 >> (((v1 * v1)) & 63))
                        default:
                            switch "a" {
                            case "\n\t":
                                v5 -= (v5 % ((v5) | 1))
                            }
                        }
                        _ = fi45
                        _ = v46
                    }
                }
                var v154 uint32 = v5
                for fi155 := 0; fi155 < 1; fi155++ {
                    v5 = 2
                    v3 = append(v3, "foo bar")
                    v156 := GenS12{F0: 18, F1: 860}
                    _ = fi155
                    _ = v156
                }
                _ = v154
            } else {
                v3 = append(v3, "\n\t")
                switch 915 {
                case -107:
                    if true {
                        for fi157 := 0; fi157 < 5; fi157++ {
                            var v158 int64 = 122
                            _ = fi157
                            _ = v158
                        }
                        for fi159 := 0; fi159 < 10; fi159++ {
                            v5 += (^(v5))
                            _ = fi159
                        }
                        if (!false) {
                            var v160 uint64 = 2
                            if ("hello" <= "foo bar") {
                                v161 := GenS20{F0: 255, F1: (!false)}
                                _ = v161
                            } else {
                                if true {
                                    fmt.Println(v2, v160, v1)
                                    for fi162 := 0; fi162 < 16; fi162++ {
                                        if true {
                                            for fi163 := 0; fi163 < 17; fi163++ {
                                                if true {
                                                    var v164 uint16 = 32767
                                                    _ = v164
                                                }
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                switch "hello" {
                                                case "\n\t":
                                                    v165 := GenS3{F0: (false && true), F1: (!true), F2: (fi162 | (fi163 << ((v1) & 63)))}
                                                    _ = v165
                                                }
                                                v5 -= 273
                                                _ = fi163
                                            }
                                        }
                                        if false {
                                            v5 = 359
                                        } else {
                                            v166 := GenS21{F0: genFn0(), F1: 375, F2: (v5 - 792)}
                                            switch fi162 {
                                            case 42:
                                                for fi167 := 0; fi167 < 13; fi167++ {
                                                    var v168 bool = true
                                                    v168 = v168
                                                    _ = fi167
                                                    _ = v168
                                                }
                                            case 0:
                                                var v169 uint = (v1 % ((188) | 1))
                                                _ = v169
                                            default:
                                                v160 += (v160 + 42)
                                            }
                                            _ = v166
                                        }
                                        _ = fi162
                                    }
                                    v170 := GenS20{F0: 17, F1: (!true)}
                                    _ = v170
                                }
                                var v171 float64 = float64(1e10)
                                switch "a" {
                                case "hello":
                                    if false {
                                        var v172 uint64 = 1
                                        _ = v172
                                    } else {
                                        fmt.Println(v1, v5, v160)
                                    }
                                default:
                                    for fi173 := 0; fi173 < 1; fi173++ {
                                        var v174 bool = false
                                        var v175 string = "hello"
                                        switch 157 {
                                        case -90:
                                            for fi176 := 0; fi176 < 7; fi176++ {
                                                if (562 < v2) {
                                                    if false {
                                                        fmt.Println(v175)
                                                        if (1.5 <= 0.1) {
                                                            v177 := GenS22{F0: (v175 + "")}
                                                            switch v175 {
                                                            case "":
                                                                v174 = (2 != (v1 | 65535))
                                                            default:
                                                                v178 := GenS23{F0: 0, F1: -1, F2: genFn0()}
                                                                _ = v178
                                                            }
                                                            _ = v177
                                                        } else {
                                                            v179 := GenS3{F0: v174, F1: v174, F2: fi173}
                                                            v174 = true
                                                            _ = v179
                                                        }
                                                    } else {
                                                        var v180 bool = v174
                                                        _ = v180
                                                    }
                                                }
                                                var v181 uint32 = v5
                                                switch v175 {
                                                case "":
                                                    var v182 int64 = 9223372036854775807
                                                    _ = v182
                                                }
                                                _ = fi176
                                                _ = v181
                                            }
                                        case -1:
                                            var v183 bool = v174
                                            _ = v183
                                        }
                                        _ = fi173
                                        _ = v174
                                        _ = v175
                                    }
                                }
                                _ = v171
                            }
                            switch -889 {
                            case 1:
                                switch 716 {
                                case 9223372036854775807:
                                    switch 32767 {
                                    case 480:
                                        if ((!(v5 <= 965)) && true) {
                                            for fi184 := 0; fi184 < 8; fi184++ {
                                                delete(v4, v5)
                                                var v185 uint = v1
                                                _ = fi184
                                                _ = v185
                                            }
                                            fmt.Println(v1, v1)
                                            v186 := v4[(^(v5))]
                                            _ = v186
                                        }
                                    case 9223372036854775807:
                                        var v187 int8 = 98
                                        _ = v187
                                    default:
                                        var v188 int32 = v2
                                        _ = v188
                                    }
                                case -9223372036854775808:
                                    var v189 int8 = 1
                                    _ = v189
                                case 4294967295:
                                    fmt.Println(v2, v2)
                                default:
                                    switch (((("a" + "hello") + (("foo bar" + "foo bar") + "a")) + ("foo bar" + "hello")) + ("hello" + "")) {
                                    case "a":
                                        v190 := GenS22{F0: "foo bar"}
                                        _ = v190
                                    default:
                                        switch 811 {
                                        case 32767:
                                            switch "foo bar" {
                                            case "a":
                                                v160 = (v160 - v160)
                                            default:
                                                v191 := GenS20{F0: 127, F1: false}
                                                _ = v191
                                            }
                                        }
                                    }
                                }
                            case -1:
                                var v192 float32 = 0.0
                                _ = v192
                            default:
                                switch 128 {
                                case -2147483648:
                                    for fi193 := 0; fi193 < 6; fi193++ {
                                        v160 += (^(v160))
                                        _ = fi193
                                    }
                                case -9223372036854775808:
                                    v194 := GenS9{F0: genFn0(), F1: ("a" + ("" + "foo bar")), F2: 42}
                                    _ = v194
                                case -154:
                                    _ = len(v4)
                                default:
                                    switch 1 {
                                    case 128:
                                        fmt.Println(v2, v1)
                                    case -897:
                                        for fi195 := 0; fi195 < 2; fi195++ {
                                            var v196 int8 = -128
                                            switch fi195 {
                                            case -673:
                                                var v197 uint32 = (v5 << ((v1) & 63))
                                                _ = v197
                                            }
                                            _ = fi195
                                            _ = v196
                                        }
                                    case 32768:
                                        v3 = append(v3, "\n\t")
                                    default:
                                        v3 = append(v3, "a")
                                    }
                                }
                            }
                            _ = v160
                        }
                    } else {
                        v3 = append(v3, ("hello" + "hello"))
                        _ = len(v3)
                        switch -2147483648 {
                        case 32768:
                            v198 := GenS20{F0: 127, F1: true}
                            _ = v198
                        case 2147483647:
                            for fi199 := 0; fi199 < 2; fi199++ {
                                v5 = (v5 - 879)
                                fmt.Println(v5, v5, fi199)
                                var v200 int16 = -32768
                                _ = fi199
                                _ = v200
                            }
                        default:
                            for fi201 := 0; fi201 < 8; fi201++ {
                                if true {
                                    v1 = v1
                                    switch (("foo bar" + ("foo bar" + ((((("\n\t" + "hello") + "foo bar") + "\n\t") + "hello") + "foo bar"))) + "\n\t") {
                                    case "\n\t":
                                        v202 := GenS7{F0: 732, F1: -1}
                                        _ = v202
                                    default:
                                        var v203 uint32 = (v5 & 1)
                                        _ = v203
                                    }
                                } else {
                                    var v204 uint16 = 255
                                    v2 -= (-(v2))
                                    _ = v204
                                }
                                for fi205 := 0; fi205 < 9; fi205++ {
                                    delete(v4, v5)
                                    v2 += (v2 << ((0) & 63))
                                    v206 := GenS15{F0: ("a" + ""), F1: 389, F2: 0}
                                    var v207 uint = (v1 - (v1 | 713))
                                    if false {
                                        fmt.Println(v2)
                                        v3 = append(v3, "")
                                        fmt.Println(v207, v5, v207)
                                        v2 = v2
                                        v208 := GenS6{F0: 255}
                                        _ = v208
                                    }
                                    _ = len(v4)
                                    _ = fi205
                                    _ = v206
                                    _ = v207
                                }
                                _ = fi201
                            }
                        }
                    }
                }
            }
            if false {
                var v209 float64 = (-2.25 + 0.1)
                if false {
                    v3 = append(v3, ("" + "\n\t"))
                    v209 -= (1.5 + v209)
                    for fi210 := 0; fi210 < 4; fi210++ {
                        v211 := v4[v5]
                        v3 = append(v3, "a")
                        v1 -= v1
                        _ = fi210
                        _ = v211
                    }
                } else {
                    var v212 float32 = 0.1
                    for fi213 := 0; fi213 < 3; fi213++ {
                        fmt.Println(v1)
                        _ = fi213
                    }
                    _ = v212
                }
                _ = v209
            }
            _ = v43
        }
    } else {
        fmt.Println(v5)
    }
    fmt.Println(v2)
    v214 := GenS8{F0: (^(v5)), F1: false, F2: 32767}
    fmt.Println(v1, v1)
    v215 := GenS11{F0: 837}
    v216 := GenS24{F0: -172, F1: (("hello" + "\n\t") + (("a" + "") + ("\n\t" + "\n\t"))), F2: -455}
    for fi217 := 0; fi217 < 18; fi217++ {
        v2 += 27
        _ = fi217
    }
    var v218 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v218:", v218)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v214
    _ = v215
    _ = v216
    _ = v218
    _ = v3
    _ = v4
    _ = v5
}
