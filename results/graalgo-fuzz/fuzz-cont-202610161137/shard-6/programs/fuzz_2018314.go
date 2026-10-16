package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS3 struct { F0 bool; F1 uint64; F2 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 string; F2 int }
type GenS6 struct { F0 int8; F1 bool; F2 int }
type GenS7 struct { F0 uint; F1 int64; F2 int32 }
type GenS8 struct { F0 int8; F1 uint8; F2 int8 }
type GenS9 struct { F0 int8; F1 int8; F2 int }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int32; F1 int; F2 int8 }
type GenS12 struct { F0 uint64; F1 uint16 }
type GenS13 struct { F0 uint32; F1 int16 }
type GenS14 struct { F0 uint8; F1 string }
type GenS15 struct { F0 uint; F1 int64 }
type GenS16 struct { F0 int8; F1 string }
func genFn0() uint8 {
    var v5 int = 2147483647
    var v6 bool = true
    if (v6 && (1.5 == 3.14159)) {
        v7 := GenS1{F0: 127}
        for fi8 := 0; fi8 < 20; fi8++ {
            switch fi8 {
            case -1:
                for fi9 := 0; fi9 < 7; fi9++ {
                    var v10 int32 = 349
                    var v11 uint16 = 32767
                    var v12 int64 = 9223372036854775807
                    v13 := GenS1{F0: int8(v11)}
                    var v14 bool = (!v6)
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                    _ = v13
                    _ = v14
                }
            case 0:
                for fi15 := 0; fi15 < 11; fi15++ {
                    if v6 {
                        if (1e10 >= 1e10) {
                            if (v5 < v5) {
                                fmt.Println(fi8, v5, v5)
                            } else {
                                v16 := GenS1{F0: 1}
                                _ = v16
                            }
                            fmt.Println(fi15, fi8)
                            v17 := []int{(fi8 / ((-1) | 1)), -811, 32768, (fi8 | (v5 + (-(fi8))))}
                            _ = v17
                        } else {
                            v18 := []int{fi15}
                            v5 -= fi15
                            v5 = 1
                            var v19 float32 = float32(1.5)
                            var v20 int16 = -32
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                    } else {
                        if (v6 || v6) {
                            v21 := []int{v5, v5, 553}
                            v6 = v6
                            _ = v21
                        }
                    }
                    var v22 bool = false
                    _ = fi15
                    _ = v22
                }
            case -9223372036854775808:
                var v23 int16 = 1
                _ = v23
            default:
                v24 := GenS1{F0: -6}
                _ = v24
            }
            _ = fi8
        }
        switch "hello" {
        case "":
            v25 := GenS2{F0: -903, F1: 255, F2: 0}
            _ = v25
        case "\n\t":
            v6 = true
        case "a":
            v26 := []int{224, (-(v5))}
            _ = v26
        default:
            var v27 uint = 4294967295
            _ = v27
        }
        _ = v7
    } else {
        v6 = (!v6)
        var v28 int64 = 1
        v29 := GenS3{F0: ((v28 & 0) <= (v28 * v28)), F1: 180, F2: 127}
        _ = v28
        _ = v29
    }
    v30 := []int{(-(v5))}
    v31 := GenS2{F0: 487, F1: 255, F2: 42}
    _ = v30
    _ = v31
    _ = v5
    _ = v6
    return 69
}

func main() {
    var v1 uint64 = 776
    var v2 int32 = -2147483648
    var v3 uint = 4294967295
    v4 := make(map[string]uint32)
    var v32 int16 = 0
    var v33 int64 = -9223372036854775808
    fmt.Println(v2, v3, v3)
    v34 := GenS3{F0: ((v3 ^ 18446744073709551615) >= v3), F1: 1, F2: 38}
    for fi35 := 0; fi35 < 8; fi35++ {
        if false {
            var v36 float32 = (3.14159 + 0.1)
            for fi37 := 0; fi37 < 20; fi37++ {
                fmt.Println(v1)
                v33 -= (-(v33))
                if true {
                    v38 := v4["\n\t"]
                    _ = v38
                } else {
                    v2 = (v2 & (v2 * (v2 << ((v3) & 63))))
                    var v39 bool = (-2147483648 >= (v2 ^ v2))
                    var v40 string = "hello"
                    _ = v39
                    _ = v40
                }
                _ = fi37
            }
            _ = v36
        } else {
            if false {
                fmt.Println(fi35, v33, v1)
            }
            v41 := GenS4{F0: 89}
            v42 := GenS5{F0: 255, F1: ("a" + (("foo bar" + "\n\t") + (("foo bar" + "\n\t") + "foo bar"))), F2: (fi35 ^ fi35)}
            _ = v41
            _ = v42
        }
        v43 := GenS4{F0: -128}
        v1 -= (v1 << ((4294967295) & 63))
        _ = fi35
        _ = v43
    }
    var v44 uint8 = genFn0()
    for fi45 := 0; fi45 < 6; fi45++ {
        var v46 int32 = v2
        if (("foo bar" + "") == "hello") {
            v47 := GenS6{F0: 100, F1: true, F2: 4294967295}
            v33 += int64(v1)
            v32 -= v32
            _ = v47
        } else {
            v48 := []int{32767}
            v49 := GenS7{F0: (v3 * (v3 % (((v3 - v3)) | 1))), F1: v33, F2: (v2 + v2)}
            fmt.Println(v33, v3)
            _ = v48
            _ = v49
        }
        _ = fi45
        _ = v46
    }
    v2 = (v2 << (((v3 | (v3 + v3))) & 63))
    if true {
        var v50 int8 = 42
        if (0 >= (^(v33))) {
            for fi51 := 0; fi51 < 17; fi51++ {
                switch (fi51 * (fi51 % ((4294967295) | 1))) {
                case 921:
                    if false {
                        fmt.Println(v44, v44, v32)
                    } else {
                        v52 := GenS8{F0: (^(v50)), F1: 0, F2: v50}
                        v53 := []int{fi51, 128}
                        if true {
                            fmt.Println(fi51, v44, v2)
                        }
                        _ = v52
                        _ = v53
                    }
                default:
                    switch "\n\t" {
                    case "\n\t":
                        var v54 uint64 = 1
                        _ = v54
                    case "hello":
                        switch ("foo bar" + "hello") {
                        case "foo bar":
                            v50 = -1
                        case "hello":
                            v55 := GenS3{F0: ((840 > 826) || true), F1: (v1 ^ v1), F2: v50}
                            _ = v55
                        }
                    case "a":
                        if true {
                            v56 := GenS9{F0: v50, F1: v50, F2: (fi51 >> ((v3) & 63))}
                            fmt.Println(v44, v50)
                            var v57 uint64 = (v1 >> ((v3) & 63))
                            _ = v56
                            _ = v57
                        } else {
                            v58 := GenS10{F0: 1}
                            _ = v58
                        }
                    default:
                        if ((1e10 * 1.5) > 1.5) {
                            _ = len(v4)
                            v59 := GenS3{F0: true, F1: 749, F2: (v50 | v50)}
                            _ = v59
                        }
                    }
                }
                _ = fi51
            }
            v32 = (v32 / ((int16(v50)) | 1))
            fmt.Println(v33)
        } else {
            v60 := GenS11{F0: (^(v2)), F1: 2147483647, F2: -128}
            var v61 uint16 = 1
            switch -65 {
            case 32767:
                _ = len(v4)
            case 2147483647:
                v62 := []int{32768, -1}
                _ = v62
            case 42:
                if ((v50 + -2) < (v50 << (((v3 >> ((4294967295) & 63))) & 63))) {
                    var v63 uint16 = uint16(v1)
                    v64 := []int{502, -791, int(v44)}
                    for fi65 := 0; fi65 < 4; fi65++ {
                        for fi66 := 0; fi66 < 9; fi66++ {
                            var v67 uint8 = (v44 - v44)
                            if false {
                                var v68 int32 = 1
                                v32 += (v32 + v32)
                                v33 -= -9223372036854775808
                                _ = v68
                            } else {
                                fmt.Println(fi65, v32, v1)
                                var v69 bool = true
                                _ = v69
                            }
                            _ = fi66
                            _ = v67
                        }
                        _ = fi65
                    }
                    _ = v63
                    _ = v64
                } else {
                    v32 += 42
                    if true {
                        switch 1 {
                        case -443:
                            fmt.Println(v1, v1)
                        default:
                            for fi70 := 0; fi70 < 14; fi70++ {
                                fmt.Println(v61)
                                _ = fi70
                            }
                        }
                    }
                    for fi71 := 0; fi71 < 7; fi71++ {
                        fmt.Println(v1)
                        if true {
                            v72 := []int{int(v32)}
                            for fi73 := 0; fi73 < 9; fi73++ {
                                if (((v2 % (((v2 / (((v2 >> ((v3) & 63))) | 1))) | 1)) >= int32(fi73)) || false) {
                                    _ = len(v4)
                                    v1 -= v1
                                    var v74 float64 = 0.0
                                    _ = v74
                                }
                                switch "\n\t" {
                                case "a":
                                    v3 = uint(v1)
                                }
                                _ = fi73
                            }
                            var v75 int64 = v33
                            _ = v72
                            _ = v75
                        } else {
                            for fi76 := 0; fi76 < 4; fi76++ {
                                var v77 uint8 = v44
                                if true {
                                    fmt.Println(v3, v33, v32)
                                    var v78 bool = (v2 >= v2)
                                    _ = v78
                                } else {
                                    var v79 uint32 = 2147483648
                                    _ = v79
                                }
                                if ((("" + "") + ("hello" + "a")) >= "foo bar") {
                                    var v80 int64 = (v33 << ((v3) & 63))
                                    if (!true) {
                                        if (v50 == v50) {
                                            v50 -= v50
                                            var v81 bool = false
                                            switch "hello" {
                                            case "":
                                                fmt.Println(fi76, v2)
                                            }
                                            var v82 bool = (fi76 < (fi76 ^ (-(fi71))))
                                            _ = v81
                                            _ = v82
                                        } else {
                                            var v83 uint32 = 0
                                            switch "foo bar" {
                                            case "hello":
                                                if false {
                                                    var v84 uint64 = uint64(v44)
                                                    v85 := GenS12{F0: (^(v84)), F1: 42}
                                                    _ = v84
                                                    _ = v85
                                                }
                                            case "\n\t":
                                                v83 = 2147483648
                                                var v86 string = "a"
                                                _ = v86
                                            }
                                            v77 += 0
                                            _ = v83
                                        }
                                        v32 = 1
                                        v44 -= v44
                                    }
                                    _ = v80
                                }
                                var v87 uint64 = 42
                                _ = fi76
                                _ = v77
                                _ = v87
                            }
                            fmt.Println(v33, v44)
                        }
                        _ = fi71
                    }
                }
            default:
                v88 := GenS13{F0: 2147483648, F1: v32}
                _ = v88
            }
            _ = v60
            _ = v61
        }
        _ = v50
    } else {
        for fi89 := 0; fi89 < 11; fi89++ {
            var v90 string = "\n\t"
            _ = fi89
            _ = v90
        }
        if ((^(v44)) < v44) {
            var v91 uint = (^(v3))
            _ = v91
        } else {
            switch 4294967295 {
            case 42:
                var v92 int32 = 1
                _ = v92
            default:
                if (v33 < (-(v33))) {
                    if true {
                        fmt.Println(v44, v3)
                    } else {
                        var v93 float32 = 3.14159
                        _ = v93
                    }
                    fmt.Println(v3, v2)
                }
            }
        }
    }
    if ((false && (v3 != 65535)) && true) {
        v33 = v33
        var v94 bool = false
        var v95 uint64 = v1
        _ = v94
        _ = v95
    }
    if (true || false) {
        var v96 string = "a"
        _ = v96
    } else {
        var v97 uint = v3
        v98 := GenS14{F0: genFn0(), F1: ""}
        if false {
            for fi99 := 0; fi99 < 11; fi99++ {
                switch "foo bar" {
                case "foo bar":
                    if false {
                        var v100 string = "hello"
                        for fi101 := 0; fi101 < 12; fi101++ {
                            v1 = (v1 + v1)
                            var v102 uint64 = (v1 / ((4294967295) | 1))
                            _ = fi101
                            _ = v102
                        }
                        _ = v100
                    }
                case "":
                    v97 = v97
                case "a":
                    switch "hello" {
                    case "a":
                        if true {
                            v97 += 255
                            switch 4294967295 {
                            case -23:
                                fmt.Println(v3, v32)
                            case 128:
                                var v103 int8 = 0
                                _ = v103
                            default:
                                switch "foo bar" {
                                case "hello":
                                    var v104 int32 = 42
                                    _ = v104
                                case "":
                                    var v105 float32 = -2.25
                                    _ = v105
                                case "foo bar":
                                    for fi106 := 0; fi106 < 18; fi106++ {
                                        fmt.Println(fi99, fi106)
                                        v107 := []int{fi106, (fi99 | fi99), -9223372036854775808, (-(fi106))}
                                        _ = fi106
                                        _ = v107
                                    }
                                default:
                                    v108 := v4["foo bar"]
                                    _ = v108
                                }
                            }
                            v3 = v97
                        }
                    case "":
                        v33 += 9223372036854775807
                    case "hello":
                        fmt.Println(v97, v97, v2)
                    }
                default:
                    if true {
                        v44 = genFn0()
                        v109 := GenS7{F0: (v3 * 0), F1: (v33 / ((v33) | 1)), F2: -1}
                        _ = v109
                    }
                }
                _ = fi99
            }
            v110 := GenS9{F0: -54, F1: -104, F2: 2147483647}
            delete(v4, "")
            _ = v110
        } else {
            for fi111 := 0; fi111 < 16; fi111++ {
                v3 = v97
                _ = fi111
            }
            if (2 <= 32767) {
                var v112 int8 = -1
                _ = v112
            } else {
                v2 = v2
            }
            v33 += -1
        }
        _ = v97
        _ = v98
    }
    switch 2147483647 {
    case 2147483647:
        v3 -= 2
    case 42:
        var v113 float32 = float32(0.0)
        _ = v113
    default:
        for fi114 := 0; fi114 < 11; fi114++ {
            v115 := GenS15{F0: 427, F1: (-(v33))}
            fmt.Println(v32, v2)
            _ = fi114
            _ = v115
        }
    }
    var v116 int16 = 32767
    v117 := []int{-738, 128, 1}
    if (v33 <= (v33 / ((v33) | 1))) {
        var v118 uint64 = (v1 - (^(v1)))
        _ = v118
    } else {
        fmt.Println(v2, v116, v116)
    }
    for fi119 := 0; fi119 < 14; fi119++ {
        for fi120 := 0; fi120 < 10; fi120++ {
            if true {
                switch "foo bar" {
                case "foo bar":
                    v33 = v33
                default:
                    var v121 int8 = 100
                    _ = v121
                }
                var v122 int = (fi120 / ((fi119) | 1))
                v32 = int16(fi120)
                _ = v122
            } else {
                v33 -= v33
                v3 -= v3
            }
            if (!(0 <= 0)) {
                switch fi120 {
                case 9223372036854775807:
                    var v123 int = (fi119 >> ((2) & 63))
                    _ = v123
                default:
                    v116 -= int16(fi120)
                }
                if true {
                    switch (fi119 / ((fi120) | 1)) {
                    case 42:
                        fmt.Println(fi119, v33)
                    case 0:
                        v1 = 2
                    }
                    v124 := GenS16{F0: -94, F1: ("" + ("" + ("foo bar" + ((((("" + ("foo bar" + "hello")) + "a") + "a") + ("hello" + (("" + ("foo bar" + "a")) + (("foo bar" + (("foo bar" + "a") + "foo bar")) + "\n\t")))) + (("" + (("\n\t" + "hello") + "a")) + ("hello" + ((("foo bar" + (("foo bar" + ("a" + "a")) + "\n\t")) + "foo bar") + ("a" + "hello"))))))))}
                    _ = len(v4)
                    v4["foo bar"] = 162
                    _ = v124
                } else {
                    var v125 int16 = (v116 ^ (v116 - v116))
                    v126 := GenS2{F0: -2147483648, F1: 128, F2: (^(v1))}
                    _ = v125
                    _ = v126
                }
            }
            fmt.Println(v3, fi120, v2)
            _ = fi120
        }
        _ = fi119
    }
    var v127 float64 = 1.5
    var v128 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v116:", v116)
    fmt.Println("slicelen v117:", len(v117), v117)
    fmt.Println("v127:", v127)
    fmt.Println("v128:", v128)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    _ = v1
    _ = v116
    _ = v117
    _ = v127
    _ = v128
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v44
}
