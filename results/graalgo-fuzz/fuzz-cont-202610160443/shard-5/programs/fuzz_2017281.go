package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int16; F1 int32; F2 string }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 int32; F1 int64; F2 int8 }
type GenS6 struct { F0 uint; F1 uint16 }
type GenS7 struct { F0 int; F1 uint64 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint16; F1 int16 }
type GenS11 struct { F0 uint8; F1 string }
type GenS12 struct { F0 int32; F1 uint; F2 int16 }
type GenS13 struct { F0 int32; F1 int16; F2 string }
type GenS14 struct { F0 int8; F1 uint64 }
type GenS15 struct { F0 uint16; F1 uint; F2 int }

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[string]string)
    var v3 bool = true
    var v4 int8 = 1
    var v5 uint8 = 180
    if (!v3) {
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 float64 = 0.0
            fmt.Println(v1, v4, fi6)
            if (false && false) {
                for fi8 := 0; fi8 < 4; fi8++ {
                    if true {
                        v3 = (!(v3 && v3))
                    }
                    v5 = v5
                    _ = fi8
                }
            } else {
                var v9 float64 = ((v7 + 1.5) + v7)
                v7 -= 1e10
                switch "foo bar" {
                case "":
                    var v10 float64 = (((0.0 + (float64(3.14159) / (((((((v9 * float64((-2.25 + -2.25))) * (v7 / ((0.0 / v7) - float64(0.0)))) * v7) / (v9 * v9)) * float64(0.0)) + (v9 / 1.5)) * v7))) * (v9 / float64(0.1))) - (1.5 * -2.25))
                    _ = v10
                default:
                    v5 = (^(v5))
                }
                _ = v9
            }
            _ = fi6
            _ = v7
        }
        if ((("a" + "a") + "foo bar") < "\n\t") {
            var v11 int64 = 42
            _ = v11
        } else {
            v5 = (v5 - v5)
            v5 = v5
        }
    } else {
        v1 -= (v1 << ((42) & 63))
        var v12 int = -9223372036854775808
        if true {
            var v13 int64 = 0
            var v14 uint = 754
            switch "a" {
            case "\n\t":
                if v3 {
                    var v15 float32 = (0.1 / float32((1.5 * (-2.25 * 1e10))))
                    _ = v15
                } else {
                    v16 := v2["foo bar"]
                    v14 = v14
                    for fi17 := 0; fi17 < 7; fi17++ {
                        v12 -= fi17
                        if v3 {
                            var v18 bool = v3
                            fmt.Println(v18, v14)
                            _ = v18
                        } else {
                            if (0 != v4) {
                                if true {
                                    if (9223372036854775807 == -9223372036854775808) {
                                        var v19 uint64 = 18446744073709551615
                                        var v20 int16 = 0
                                        _ = v19
                                        _ = v20
                                    } else {
                                        if v3 {
                                            for fi21 := 0; fi21 < 18; fi21++ {
                                                if (!(2 < 215)) {
                                                    v1 -= (v1 + 0)
                                                    v14 += v14
                                                    v1 += v1
                                                    v13 = (v13 / ((-9223372036854775808) | 1))
                                                }
                                                var v22 int16 = 1000
                                                _ = fi21
                                                _ = v22
                                            }
                                            for fi23 := 0; fi23 < 2; fi23++ {
                                                v3 = (v3 || v3)
                                                _ = fi23
                                            }
                                            v16 = ""
                                        } else {
                                            switch 0 {
                                            case -2147483648:
                                                for fi24 := 0; fi24 < 6; fi24++ {
                                                    if (!(32767 < 1000)) {
                                                        v5 -= (v5 % (((^(v5))) | 1))
                                                        if false {
                                                            var v25 int16 = 90
                                                            switch (fi24 % ((v12) | 1)) {
                                                            case 127:
                                                                switch (("a" + "hello") + ((v16 + v16) + v16)) {
                                                                case "\n\t":
                                                                    v26 := GenS1{F0: fi17, F1: v13}
                                                                    v13 = (v13 | 55)
                                                                    v1 = (v1 - (v1 | (^(v1))))
                                                                    if ((1.5 / (3.14159 * 3.14159)) > 0.1) {
                                                                        var v27 float64 = 1e10
                                                                        var v28 int = (v12 * 609)
                                                                        v28 = 127
                                                                        v3 = (!(-1 < -631))
                                                                        v16 = "\n\t"
                                                                        if (v25 == (v25 & (v25 % ((v25) | 1)))) {
                                                                            if v3 {
                                                                                var v29 int = v12
                                                                                var v30 int32 = v1
                                                                                _ = v29
                                                                                _ = v30
                                                                            }
                                                                            v12 = (v12 ^ (fi17 << (((v14 ^ v14)) & 63)))
                                                                            if v3 {
                                                                                if v3 {
                                                                                    var v31 int8 = v4
                                                                                    if (v3 && v3) {
                                                                                        var v32 uint64 = 2
                                                                                        if v3 {
                                                                                            fmt.Println(v3)
                                                                                        }
                                                                                        _ = v32
                                                                                    } else {
                                                                                        var v33 int64 = 0
                                                                                        if (!v3) {
                                                                                            var v34 uint64 = 0
                                                                                            _ = len(v2)
                                                                                            var v35 int16 = 1000
                                                                                            if v3 {
                                                                                                v5 = (^(v5))
                                                                                            } else {
                                                                                                v36 := []int{(fi17 - (fi17 >> ((v14) & 63))), (v28 >> (((^(v14))) & 63)), -2147483648}
                                                                                                _ = v36
                                                                                            }
                                                                                            fmt.Println(v13, v27, v35)
                                                                                            v4 -= -1
                                                                                            v37 := GenS2{F0: (^(v35)), F1: (v1 >> (((v14 - (v14 % (((^(v14))) | 1)))) & 63)), F2: (((v16 + (v16 + (("\n\t" + "hello") + v16))) + "\n\t") + v16)}
                                                                                            v2[v16] = (v16 + "\n\t")
                                                                                            fmt.Println(v27)
                                                                                            switch (fi24 % ((int(v1)) | 1)) {
                                                                                            case 1:
                                                                                                v38 := GenS3{F0: 1, F1: 200}
                                                                                                _ = v38
                                                                                            default:
                                                                                                if false {
                                                                                                    var v39 float64 = float64(1.5)
                                                                                                    switch ((("\n\t" + v16) + (((("" + (("a" + ("hello" + "hello")) + v16)) + v16) + v16) + v16)) + "hello") {
                                                                                                    case "\n\t":
                                                                                                        v27 -= 1e10
                                                                                                    default:
                                                                                                        var v40 uint32 = 2147483647
                                                                                                        _ = v40
                                                                                                    }
                                                                                                    _ = v39
                                                                                                }
                                                                                            }
                                                                                            if (-128 == int8(v35)) {
                                                                                                v5 += (v5 * (v5 << ((v14) & 63)))
                                                                                                delete(v2, (v16 + (v16 + v16)))
                                                                                                var v41 uint8 = (v5 << (((v14 * 18446744073709551615)) & 63))
                                                                                                _ = v41
                                                                                            }
                                                                                            v42 := GenS2{F0: v35, F1: -518, F2: (("" + "foo bar") + v16)}
                                                                                            switch ("\n\t" + v16) {
                                                                                            case "foo bar":
                                                                                                v43 := GenS2{F0: (v35 & 42), F1: (v1 >> ((v14) & 63)), F2: v16}
                                                                                                _ = v43
                                                                                            }
                                                                                            if v3 {
                                                                                                var v44 bool = false
                                                                                                v45 := []int{(fi24 & (-(v28))), fi24, -9223372036854775808, 275}
                                                                                                v46 := GenS1{F0: 4294967295, F1: -1}
                                                                                                _ = v44
                                                                                                _ = v45
                                                                                                _ = v46
                                                                                            } else {
                                                                                                if false {
                                                                                                    switch "" {
                                                                                                    case "":
                                                                                                        v47 := []int{}
                                                                                                        _ = v47
                                                                                                    case "foo bar":
                                                                                                        if v3 {
                                                                                                            v48 := GenS2{F0: v25, F1: int32(v14), F2: v16}
                                                                                                            delete(v2, "\n\t")
                                                                                                            _ = v48
                                                                                                        } else {
                                                                                                            var v49 uint = v14
                                                                                                            v13 += (-(v13))
                                                                                                            v50 := []int{int(v25)}
                                                                                                            _ = v49
                                                                                                            _ = v50
                                                                                                        }
                                                                                                    default:
                                                                                                        v2[v16] = ("\n\t" + ("hello" + (v16 + ("\n\t" + "a"))))
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            var v51 uint32 = 2147483647
                                                                                            v52 := []int{}
                                                                                            v53 := GenS1{F0: v12, F1: (v13 | (v33 / ((v13) | 1)))}
                                                                                            switch (fi24 << (((v14 | v14)) & 63)) {
                                                                                            case -702:
                                                                                                var v54 int64 = int64(v34)
                                                                                                _ = v54
                                                                                            case 1:
                                                                                                var v55 uint64 = (v34 ^ 42)
                                                                                                _ = v55
                                                                                            case 0:
                                                                                                var v56 float64 = float64(0.1)
                                                                                                _ = v56
                                                                                            }
                                                                                            _ = v34
                                                                                            _ = v35
                                                                                            _ = v37
                                                                                            _ = v42
                                                                                            _ = v51
                                                                                            _ = v52
                                                                                            _ = v53
                                                                                        } else {
                                                                                            v57 := []int{(^(v28))}
                                                                                            v58 := GenS1{F0: v28, F1: (v33 % ((v13) | 1))}
                                                                                            if false {
                                                                                                var v59 int32 = -1
                                                                                                if v3 {
                                                                                                    v60 := GenS4{F0: 32767, F1: v33}
                                                                                                    _ = v60
                                                                                                }
                                                                                                _ = v59
                                                                                            }
                                                                                            _ = v57
                                                                                            _ = v58
                                                                                        }
                                                                                        _ = v33
                                                                                    }
                                                                                    if v3 {
                                                                                        if ((v5 + 0) == (v5 >> (((v14 + v14)) & 63))) {
                                                                                            var v61 string = v16
                                                                                            var v62 int8 = -94
                                                                                            _ = v61
                                                                                            _ = v62
                                                                                        } else {
                                                                                            if v3 {
                                                                                                var v63 uint8 = v5
                                                                                                _ = v63
                                                                                            }
                                                                                            v25 -= -325
                                                                                        }
                                                                                        if (779 <= (v14 - 1)) {
                                                                                            switch v16 {
                                                                                            case "hello":
                                                                                                var v64 uint16 = 42
                                                                                                _ = v64
                                                                                            }
                                                                                            if v3 {
                                                                                                v65 := GenS4{F0: 32767, F1: (v13 | (-(v13)))}
                                                                                                v13 = int64(v25)
                                                                                                v66 := GenS5{F0: (v1 | v1), F1: v13, F2: v4}
                                                                                                v67 := []int{(v12 / ((4294967295) | 1)), (^(v12)), (^(v28)), -1}
                                                                                                var v68 uint64 = uint64(v5)
                                                                                                _ = v65
                                                                                                _ = v66
                                                                                                _ = v67
                                                                                                _ = v68
                                                                                            }
                                                                                        } else {
                                                                                            v2["hello"] = v16
                                                                                        }
                                                                                    }
                                                                                    _ = v31
                                                                                } else {
                                                                                    delete(v2, "foo bar")
                                                                                    fmt.Println(v25)
                                                                                }
                                                                                switch ("\n\t" + (v16 + "hello")) {
                                                                                case "\n\t":
                                                                                    v69 := GenS6{F0: v14, F1: 255}
                                                                                    _ = v69
                                                                                case "a":
                                                                                    var v70 uint64 = 468
                                                                                    _ = v70
                                                                                case "hello":
                                                                                    v71 := GenS1{F0: (v12 % ((int(v4)) | 1)), F1: (v13 >> ((42) & 63))}
                                                                                    _ = v71
                                                                                default:
                                                                                    v72 := GenS7{F0: int(v4), F1: 18446744073709551615}
                                                                                    _ = v72
                                                                                }
                                                                                var v73 string = (("" + (("hello" + v16) + "foo bar")) + ("foo bar" + "hello"))
                                                                                v25 -= (v25 / (((v25 | (v25 << ((935) & 63)))) | 1))
                                                                                _ = v73
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v12)
                                                                            v13 += (v13 << ((v14) & 63))
                                                                        }
                                                                        _ = v27
                                                                        _ = v28
                                                                    } else {
                                                                        v1 = v1
                                                                        fmt.Println(v3, v13, v13)
                                                                    }
                                                                    _ = v26
                                                                case "hello":
                                                                    var v74 string = v16
                                                                    _ = v74
                                                                }
                                                            case 2147483647:
                                                                var v75 uint = v14
                                                                _ = v75
                                                            default:
                                                                v76 := GenS1{F0: (v12 >> ((v14) & 63)), F1: -877}
                                                                _ = v76
                                                            }
                                                            _ = v25
                                                        } else {
                                                            if (32767 > 65535) {
                                                                switch (fi24 >> ((v14) & 63)) {
                                                                case -288:
                                                                    v77 := []int{}
                                                                    _ = v77
                                                                case -746:
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        var v78 bool = (4294967295 == (v12 >> ((v14) & 63)))
                                                                        _ = v78
                                                                    }
                                                                case 0:
                                                                    v16 = (((v16 + v16) + "") + v16)
                                                                }
                                                                v79 := GenS8{F0: 2}
                                                                _ = v79
                                                            } else {
                                                                if true {
                                                                    if v3 {
                                                                        v80 := GenS8{F0: v5}
                                                                        if false {
                                                                            var v81 int16 = 32767
                                                                            _ = v81
                                                                        }
                                                                        _ = v80
                                                                    }
                                                                    v13 = (v13 * (v13 >> (((v14 | v14)) & 63)))
                                                                }
                                                                var v82 float64 = 0.0
                                                                v14 += (^(v14))
                                                                _ = v82
                                                            }
                                                            v83 := []int{(fi17 & (v12 - (fi17 ^ -9223372036854775808))), (v12 << (((^(v14))) & 63)), 520}
                                                            if v3 {
                                                                var v84 int8 = (v4 << (((v14 * (v14 * v14))) & 63))
                                                                fmt.Println(v12)
                                                                _ = len(v2)
                                                                _ = v84
                                                            } else {
                                                                v85 := GenS9{F0: 228}
                                                                _ = v85
                                                            }
                                                            _ = v83
                                                        }
                                                        if ((^(v14)) > v14) {
                                                            fmt.Println(fi17, fi17, v1)
                                                            var v86 float32 = 3.14159
                                                            _ = v86
                                                        } else {
                                                            v1 = v1
                                                            var v87 uint16 = 254
                                                            var v88 uint64 = 2
                                                            var v89 int8 = (v4 - v4)
                                                            _ = v87
                                                            _ = v88
                                                            _ = v89
                                                        }
                                                    }
                                                    var v90 uint8 = uint8(v14)
                                                    _ = fi24
                                                    _ = v90
                                                }
                                            }
                                            for fi91 := 0; fi91 < 14; fi91++ {
                                                if (!(!((v14 | v14) < (v14 << ((255) & 63))))) {
                                                    var v92 bool = (!v3)
                                                    v13 += (v13 & (v13 ^ v13))
                                                    v93 := GenS10{F0: 376, F1: 0}
                                                    _ = v92
                                                    _ = v93
                                                } else {
                                                    _ = len(v2)
                                                    v94 := []int{(v12 & v12), v12}
                                                    _ = v94
                                                }
                                                var v95 uint32 = 4294967295
                                                var v96 uint = (v14 << ((v14) & 63))
                                                _ = fi91
                                                _ = v95
                                                _ = v96
                                            }
                                            v97 := GenS11{F0: v5, F1: v16}
                                            _ = v97
                                        }
                                    }
                                    if v3 {
                                        var v98 uint32 = 42
                                        _ = v98
                                    } else {
                                        fmt.Println(v16, v16, v16)
                                    }
                                    v13 -= int64(v14)
                                }
                                v12 = (fi17 | int(v14))
                                _ = len(v2)
                            }
                            v13 = (v13 / ((9223372036854775807) | 1))
                        }
                        _ = fi17
                    }
                    _ = v16
                }
            case "":
                if (65535 >= 786) {
                    v99 := GenS5{F0: -941, F1: v13, F2: -128}
                    switch v12 {
                    case -490:
                        v100 := GenS12{F0: v1, F1: 4294967295, F2: 0}
                        _ = v100
                    case 32768:
                        v14 = uint(v5)
                    default:
                        v12 += -1
                    }
                    _ = len(v2)
                    _ = v99
                } else {
                    var v101 float32 = 1e10
                    _ = v101
                }
            case "foo bar":
                v102 := []int{(-(v12)), (v12 ^ (^(v12))), v12, -297}
                _ = v102
            }
            _ = v13
            _ = v14
        }
        _ = v12
    }
    v103 := GenS4{F0: 32767, F1: 9223372036854775807}
    v3 = v3
    _ = len(v2)
    switch "foo bar" {
    case "hello":
        var v104 bool = v3
        _ = v104
    case "\n\t":
        switch "\n\t" {
        case "foo bar":
            if (("foo bar" + "") == ("" + "a")) {
                v105 := GenS13{F0: (v1 % ((-739) | 1)), F1: 748, F2: "foo bar"}
                v106 := []int{882, -863, 42, -2147483648}
                v5 += v5
                _ = v105
                _ = v106
            }
        case "\n\t":
            v3 = (!(((false && false) || true) || (v3 || false)))
        default:
            v4 += (v4 / (((v4 << ((117) & 63))) | 1))
        }
    default:
        switch 312 {
        case 127:
            if v3 {
                v5 = (v5 / ((v5) | 1))
            }
        }
    }
    for fi107 := 0; fi107 < 9; fi107++ {
        var v108 int = fi107
        if (!v3) {
            v109 := GenS14{F0: (v4 / ((1) | 1)), F1: 4294967295}
            if true {
                var v110 float64 = 0.1
                var v111 uint32 = 1
                if (32767 > 255) {
                    var v112 uint = 255
                    v112 -= v112
                    _ = v112
                } else {
                    if v3 {
                        if ((!v3) || false) {
                            if (985 == 42) {
                                _ = len(v2)
                                v113 := GenS15{F0: 255, F1: uint(v111), F2: fi107}
                                _ = v113
                            }
                            var v114 int8 = (v4 ^ 0)
                            _ = v114
                        }
                        if v3 {
                            v5 -= v5
                            v5 += v5
                        } else {
                            var v115 float32 = 0.1
                            delete(v2, "foo bar")
                            _ = v115
                        }
                    }
                    for fi116 := 0; fi116 < 5; fi116++ {
                        fmt.Println(v111, v5)
                        if v3 {
                            v3 = false
                        }
                        if (!false) {
                            fmt.Println(v1, v110, v108)
                            v117 := v2["\n\t"]
                            _ = v117
                        } else {
                            var v118 int64 = -9223372036854775808
                            _ = v118
                        }
                        _ = fi116
                    }
                }
                _ = v110
                _ = v111
            }
            v119 := []int{}
            _ = v109
            _ = v119
        } else {
            switch "hello" {
            case "":
                switch fi107 {
                case 127:
                    for fi120 := 0; fi120 < 9; fi120++ {
                        if v3 {
                            switch "\n\t" {
                            case "":
                                fmt.Println(v4, fi120)
                            case "a":
                                if (v3 && false) {
                                    switch 127 {
                                    case 0:
                                        v5 = (v5 | (^(v5)))
                                    case -845:
                                        v121 := []int{(v108 << ((4294967295) & 63)), fi120, (fi107 << ((464) & 63))}
                                        _ = v121
                                    }
                                    var v122 bool = v3
                                    v123 := []int{}
                                    _ = v122
                                    _ = v123
                                }
                            case "\n\t":
                                var v124 uint32 = 4294967295
                                _ = v124
                            default:
                                switch ("a" + (("\n\t" + (("\n\t" + ((("a" + ("foo bar" + "a")) + "\n\t") + ("hello" + "\n\t"))) + "a")) + "")) {
                                case "foo bar":
                                    v5 = (v5 * (v5 / ((v5) | 1)))
                                }
                            }
                            if v3 {
                                switch fi120 {
                                case -9223372036854775808:
                                    var v125 int32 = (v1 % ((0) | 1))
                                    _ = v125
                                default:
                                    v126 := []int{}
                                    _ = v126
                                }
                                var v127 float32 = (3.14159 * 1.5)
                                switch 42 {
                                case 4294967295:
                                    if v3 {
                                        v128 := []int{(fi107 / ((927) | 1)), fi120, (fi120 ^ fi120), (fi107 & fi120)}
                                        _ = v128
                                    } else {
                                        if v3 {
                                            if (v3 || false) {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    var v129 int = (-(v108))
                                                    _ = v129
                                                }
                                            }
                                            v130 := []int{}
                                            v3 = v3
                                            fmt.Println(v1, v5, v5)
                                            var v131 bool = v3
                                            _ = v130
                                            _ = v131
                                        } else {
                                            if (true || (!(v108 <= 4294967295))) {
                                                v108 = (fi120 + fi107)
                                            } else {
                                                fmt.Println(v1, v3)
                                            }
                                        }
                                        if ((1 == 42) && (!(488 != 346))) {
                                            fmt.Println(v127, v4, v5)
                                        } else {
                                            var v132 int32 = int32(v4)
                                            if v3 {
                                                fmt.Println(v108, v4, v127)
                                            }
                                            v3 = false
                                            _ = v132
                                        }
                                    }
                                case 127:
                                    var v133 bool = ((v3 && (!(127 <= (v108 ^ (v108 % ((42) | 1)))))) && ((v4 >> ((680) & 63)) > -128))
                                    _ = v133
                                case 42:
                                    v1 = v1
                                default:
                                    var v134 bool = false
                                    _ = v134
                                }
                                _ = v127
                            }
                            v135 := []int{fi107}
                            _ = v135
                        }
                        _ = fi120
                    }
                case 76:
                    _ = len(v2)
                case -292:
                    if (32767 > -1) {
                        if ((false && true) && (-770 < 498)) {
                            v136 := []int{fi107, fi107, v108}
                            _ = v136
                        } else {
                            for fi137 := 0; fi137 < 9; fi137++ {
                                fmt.Println(v4)
                                _ = fi137
                            }
                            v4 = (v4 / (((v4 | v4)) | 1))
                            v4 += int8(fi107)
                        }
                    }
                default:
                    v3 = true
                }
            case "foo bar":
                var v138 int = fi107
                _ = v138
            case "hello":
                for fi139 := 0; fi139 < 9; fi139++ {
                    var v140 int16 = -32768
                    _ = fi139
                    _ = v140
                }
            default:
                v141 := []int{fi107}
                _ = v141
            }
            var v142 float32 = ((0.0 / 1.5) - 0.1)
            v143 := GenS8{F0: 255}
            _ = v142
            _ = v143
        }
        _ = fi107
        _ = v108
    }
    for fi144 := 0; fi144 < 10; fi144++ {
        fmt.Println(v4, v3)
        var v145 int16 = -837
        v4 = (v4 & (v4 % ((v4) | 1)))
        _ = fi144
        _ = v145
    }
    if v3 {
        switch ("\n\t" + "\n\t") {
        case "":
            v1 = v1
        default:
            v146 := []int{1, 1, 127, 32768}
            _ = v146
        }
        if (!(true || (!v3))) {
            for fi147 := 0; fi147 < 11; fi147++ {
                var v148 string = "foo bar"
                _ = fi147
                _ = v148
            }
            fmt.Println(v4, v3)
            fmt.Println(v3)
        } else {
            switch 9223372036854775807 {
            case 127:
                v1 = v1
            case 1:
                var v149 uint64 = 0
                _ = v149
            case 9223372036854775807:
                v4 = v4
            default:
                var v150 bool = false
                _ = v150
            }
            if (-9223372036854775808 < -1) {
                v4 += 127
                v5 -= (^(v5))
                switch -984 {
                case -1:
                    var v151 uint = 2
                    _ = v151
                case -985:
                    var v152 uint64 = uint64(v5)
                    _ = v152
                }
            }
        }
        switch -2147483648 {
        case 32768:
            var v153 float32 = (((((((0.1 / 0.1) / 1.5) + 0.1) - (-2.25 + 0.0)) - 1e10) + (float32(-2.25) - 0.0)) * 3.14159)
            _ = v153
        }
    }
    for fi154 := 0; fi154 < 12; fi154++ {
        v3 = true
        for fi155 := 0; fi155 < 12; fi155++ {
            var v156 int16 = 1000
            _ = fi155
            _ = v156
        }
        v157 := GenS11{F0: (v5 ^ (v5 >> ((678) & 63))), F1: ("" + "\n\t")}
        _ = fi154
        _ = v157
    }
    for fi158 := 0; fi158 < 9; fi158++ {
        var v159 float32 = ((1.5 * (0.1 * 1.5)) - 0.1)
        _ = fi158
        _ = v159
    }
    v160 := []int{32768, 127, 1}
    v3 = (32767 <= -32768)
    var v161 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v160:", len(v160), v160)
    fmt.Println("v161:", v161)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v103
    _ = v160
    _ = v161
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
