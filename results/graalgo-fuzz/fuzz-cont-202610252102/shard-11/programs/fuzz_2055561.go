package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int }
type GenS2 struct { F0 uint8; F1 int; F2 int8 }
type GenS3 struct { F0 int32; F1 int; F2 int16 }
type GenS4 struct { F0 int; F1 uint8; F2 int16 }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 int64; F1 int }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int16; F1 uint8 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 string; F1 string }

func main() {
    var v1 string = "\n\t"
    var v2 uint8 = 2
    v3 := make(map[int]uint32)
    v4 := []uint32{547, 2147483648, 4294967295, 672}
    if (4294967295 != 18446744073709551615) {
        v1 = "foo bar"
        var v5 string = ""
        v6 := GenS1{F0: 684, F1: (true || false), F2: 32768}
        _ = v5
        _ = v6
    } else {
        v1 = "foo bar"
        var v7 int32 = -367
        _ = v7
    }
    var v8 float32 = 1e10
    for fi9 := 0; fi9 < 1; fi9++ {
        v2 += 74
        v1 = ""
        _ = fi9
    }
    v1 = ""
    if len(v4) > 0 {
        v4[0] = 464
    }
    if false {
        var v10 uint64 = 18446744073709551615
        var v11 uint16 = 65535
        _ = v10
        _ = v11
    }
    switch -853 {
    case 42:
        v12 := GenS2{F0: v2, F1: -1, F2: -16}
        _ = v12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 int32 = -721
        var v15 int16 = -32768
        for fi16 := 0; fi16 < 1; fi16++ {
            v17 := GenS1{F0: 255, F1: false, F2: 0}
            var v18 int32 = (v14 & (v14 >> ((1) & 63)))
            if true {
                if false {
                    var v19 bool = true
                    if v19 {
                        if v19 {
                            var v20 int32 = v14
                            switch v1 {
                            case "":
                                v19 = (((v20 & v14) == (^(v14))) && (!((v8 + (((((0.1 + v8) + 1.5) * float32(-2.25)) + 0.1) + float32(float64(float32(3.14159))))) < (v8 / v8))))
                            case "\n\t":
                                delete(v3, (^(fi13)))
                            case "a":
                                var v21 uint64 = 2
                                _ = v21
                            }
                            _ = v20
                        }
                    }
                    _ = v19
                } else {
                    v18 += v18
                    if ((false && (false || (false || (!((true && false) || false))))) && (v18 >= v14)) {
                        var v22 int32 = 1
                        var v23 float64 = 0.0
                        var v24 uint8 = v2
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                }
                var v25 int32 = v14
                _ = v25
            } else {
                v26 := GenS3{F0: v14, F1: fi13, F2: -32768}
                var v27 float32 = v8
                switch (fi13 >> ((42) & 63)) {
                case -9223372036854775808:
                    switch v1 {
                    case "a":
                        v18 = (v18 << ((42) & 63))
                        v28 := GenS4{F0: -2147483648, F1: (v2 << ((18446744073709551615) & 63)), F2: (v15 ^ (^(v15)))}
                        v1 = (v1 + "foo bar")
                        _ = v28
                    case "":
                        var v29 uint64 = 42
                        _ = v29
                    case "\n\t":
                        var v30 float32 = 0.1
                        _ = v30
                    }
                case 128:
                    if false {
                        if ((!false) || false) {
                            switch (v1 + v1) {
                            case "":
                                _ = len(v3)
                            case "hello":
                                v8 += 0.1
                            case "\n\t":
                                var v31 uint16 = 409
                                _ = v31
                            }
                            if (!(("foo bar" + "hello") == ("\n\t" + (v1 + (v1 + "hello"))))) {
                                v27 = v8
                                fmt.Println(v1)
                                v4 = append(v4, 613)
                            } else {
                                var v32 int16 = (-(v15))
                                if (true && true) {
                                    var v33 int16 = 0
                                    v1 = v1
                                    _ = v33
                                } else {
                                    var v34 uint8 = (v2 % (((v2 - (v2 >> ((4294967295) & 63)))) | 1))
                                    switch fi16 {
                                    case 32767:
                                        v32 += v15
                                        var v35 int32 = (^(v14))
                                        var v36 uint64 = 18446744073709551615
                                        _ = v35
                                        _ = v36
                                    case -727:
                                        delete(v3, (fi13 % ((fi13) | 1)))
                                    }
                                    _ = v34
                                }
                                v37 := GenS5{F0: 65535, F1: 1}
                                _ = v32
                                _ = v37
                            }
                        } else {
                            var v38 int16 = v15
                            switch fi13 {
                            case 549:
                                switch (fi16 / ((fi13) | 1)) {
                                case 4294967295:
                                    var v39 int64 = 42
                                    _ = v39
                                default:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                            case -2147483648:
                                var v40 float64 = -2.25
                                _ = v40
                            case 42:
                                v2 = v2
                            }
                            if true {
                                fmt.Println(v14, fi16)
                            } else {
                                var v41 uint = 65535
                                v4 = append(v4, 65535)
                                _ = len(v4)
                                v38 -= 182
                                _ = v41
                            }
                            if (0 != 725) {
                                v8 += v27
                                var v42 int64 = -9223372036854775808
                                v8 += v27
                                _ = v42
                            }
                            _ = v38
                        }
                        if false {
                            if true {
                                switch v1 {
                                case "a":
                                    _ = len(v4)
                                }
                                v15 = (-(v15))
                            } else {
                                v14 = -1
                                var v43 uint = 0
                                v44 := v3[(fi13 | fi13)]
                                _ = v43
                                _ = v44
                            }
                            switch v1 {
                            case "\n\t":
                                fmt.Println(v8, v15)
                            case "foo bar":
                                var v45 string = ""
                                _ = v45
                            default:
                                if false {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    var v46 string = v1
                                    _ = v46
                                } else {
                                    v2 = v2
                                    var v47 uint64 = 2
                                    v48 := v3[(fi16 * (fi13 | -899))]
                                    _ = v47
                                    _ = v48
                                }
                            }
                        } else {
                            switch "foo bar" {
                            case "hello":
                                var v49 string = (v1 + (((v1 + (v1 + ((("hello" + v1) + v1) + ("" + (v1 + v1))))) + (v1 + v1)) + ""))
                                _ = v49
                            case "foo bar":
                                v50 := GenS6{F0: 848, F1: fi16}
                                _ = v50
                            case "a":
                                fmt.Println(v27, fi16)
                            }
                            if (42 == 625) {
                                if (!(!false)) {
                                    fmt.Println(fi16, v14, fi13)
                                    var v51 bool = (0 == 1)
                                    _ = v51
                                } else {
                                    v52 := GenS1{F0: 11, F1: (false || true), F2: (fi16 / ((fi16) | 1))}
                                    if (!true) {
                                        if false {
                                            switch v1 {
                                            case "\n\t":
                                                switch v1 {
                                                case "\n\t":
                                                    switch (((v1 + (v1 + "\n\t")) + v1) + (v1 + "foo bar")) {
                                                    case "hello":
                                                        _ = len(v3)
                                                        v27 = v27
                                                        v3[fi16] = 2147483647
                                                    }
                                                case "hello":
                                                    var v53 int32 = v14
                                                    _ = v53
                                                default:
                                                    if true {
                                                        if false {
                                                            if true {
                                                                switch fi16 {
                                                                case 1:
                                                                    if false {
                                                                        var v54 float64 = (float64((v27 / float32(0.1))) / ((1e10 - (-2.25 * 1.5)) + (3.14159 - 0.1)))
                                                                        if true {
                                                                            v15 = v15
                                                                            if (true || (989 != -1)) {
                                                                                v55 := v3[fi16]
                                                                                _ = v55
                                                                            }
                                                                            switch fi16 {
                                                                            case -2147483648:
                                                                                v3[4294967295] = 467
                                                                            case -117:
                                                                                v15 = v15
                                                                            default:
                                                                                v27 -= v8
                                                                            }
                                                                        }
                                                                        _ = v54
                                                                    } else {
                                                                        var v56 uint16 = 2
                                                                        v57 := GenS7{F0: v14}
                                                                        _ = v56
                                                                        _ = v57
                                                                    }
                                                                case 115:
                                                                    v58 := GenS7{F0: -247}
                                                                    _ = v58
                                                                case 128:
                                                                    v8 = (float32(-2.25) * 0.1)
                                                                default:
                                                                    if true {
                                                                        if false {
                                                                            if false {
                                                                                fmt.Println(v15, v27)
                                                                            }
                                                                            var v59 uint8 = 165
                                                                            var v60 float64 = ((((((1.5 * -2.25) / 1.0) + float64(float32(0.1))) - (0.1 / (float64((float32(((1.5 * 0.1) + -2.25)) * 0.0)) + 1.5))) + (1.5 / 1.0)) * -2.25)
                                                                            var v61 float64 = v60
                                                                            v62 := GenS1{F0: 0, F1: (int(v59) >= fi13), F2: (fi13 * (-(fi13)))}
                                                                            _ = v59
                                                                            _ = v60
                                                                            _ = v61
                                                                            _ = v62
                                                                        }
                                                                    } else {
                                                                        if true {
                                                                            v27 -= 1.5
                                                                            fmt.Println(v15)
                                                                            v8 = v8
                                                                        } else {
                                                                            _ = len(v4)
                                                                            if (100 != 100) {
                                                                                v63 := GenS6{F0: -28, F1: fi16}
                                                                                v18 = v14
                                                                                _ = v63
                                                                            }
                                                                            fmt.Println(v1, v14)
                                                                        }
                                                                        var v64 string = "hello"
                                                                        _ = v64
                                                                    }
                                                                }
                                                                var v65 int = int(v18)
                                                                v66 := GenS5{F0: 0, F1: 577}
                                                                _ = v65
                                                                _ = v66
                                                            } else {
                                                                v4 = append(v4, 935)
                                                            }
                                                            v67 := GenS8{F0: (v1 + ("a" + ""))}
                                                            switch v1 {
                                                            case "\n\t":
                                                                _ = len(v4)
                                                            case "":
                                                                if ((65535 == 18446744073709551615) || true) {
                                                                    v68 := GenS4{F0: (fi16 / ((fi13) | 1)), F1: v2, F2: (v15 * v15)}
                                                                    _ = v68
                                                                }
                                                            case "foo bar":
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                            default:
                                                                var v69 string = v1
                                                                _ = v69
                                                            }
                                                            v8 = v27
                                                            v27 = (-2.25 * (v27 / v27))
                                                            var v70 string = ""
                                                            _ = v67
                                                            _ = v70
                                                        } else {
                                                            v14 += (v18 << ((18446744073709551615) & 63))
                                                        }
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        var v71 int32 = (^(v18))
                                                        if len(v4) > 0 {
                                                            v4[0] = 4294967295
                                                        }
                                                        switch fi13 {
                                                        case 1:
                                                            var v72 uint = 18446744073709551615
                                                            _ = v72
                                                        case 2147483647:
                                                            fmt.Println(v15)
                                                        }
                                                        _ = v71
                                                    } else {
                                                        v73 := GenS5{F0: 2, F1: 262}
                                                        switch ("hello" + (("a" + v1) + "hello")) {
                                                        case "\n\t":
                                                            if (!(!true)) {
                                                                var v74 uint8 = v2
                                                                _ = v74
                                                            }
                                                        }
                                                        v75 := v3[(fi16 * fi13)]
                                                        _ = v73
                                                        _ = v75
                                                    }
                                                }
                                            }
                                            if true {
                                                if false {
                                                    v3[(fi16 % ((2147483647) | 1))] = 767
                                                    delete(v3, fi13)
                                                    var v76 int64 = -479
                                                    _ = v76
                                                }
                                                fmt.Println(v14, v14)
                                            }
                                        }
                                    } else {
                                        var v77 int = fi16
                                        v15 = (v15 - 32767)
                                        _ = v77
                                    }
                                    v1 = "hello"
                                    _ = v52
                                }
                            }
                            switch v1 {
                            case "a":
                                switch ((v1 + ("a" + v1)) + "\n\t") {
                                case "foo bar":
                                    var v78 bool = (!false)
                                    _ = v78
                                case "a":
                                    var v79 bool = false
                                    _ = v79
                                case "":
                                    var v80 int64 = 561
                                    _ = v80
                                default:
                                    var v81 uint64 = 2
                                    _ = v81
                                }
                            case "":
                                v82 := GenS7{F0: -961}
                                _ = v82
                            }
                        }
                    }
                case 9223372036854775807:
                    switch "" {
                    case "foo bar":
                        switch fi13 {
                        case 217:
                            v4 = append(v4, 2147483647)
                        case 831:
                            switch "" {
                            case "foo bar":
                                if false {
                                    var v83 bool = true
                                    _ = v83
                                }
                            case "\n\t":
                                v1 = ("\n\t" + v1)
                            case "hello":
                                v84 := v3[fi13]
                                _ = v84
                            }
                        case 619:
                            var v85 uint = 769
                            _ = v85
                        default:
                            v86 := GenS6{F0: 9223372036854775807, F1: 807}
                            switch fi16 {
                            case 303:
                                delete(v3, fi13)
                            default:
                                switch (^(fi13)) {
                                case -1:
                                    var v87 int16 = v15
                                    _ = v87
                                case 183:
                                    fmt.Println(fi16)
                                case 9223372036854775807:
                                    if (!((v18 + v18) > (^(v18)))) {
                                        switch fi13 {
                                        case 13:
                                            if ((-9223372036854775808 > (fi16 ^ (fi16 % ((-857) | 1)))) && ((0 >= 1) || (!true))) {
                                                if true {
                                                    v2 = (v2 >> ((932) & 63))
                                                    v27 = 1e10
                                                }
                                                var v88 uint8 = (v2 * (v2 / (((v2 << ((389) & 63))) | 1)))
                                                _ = v88
                                            } else {
                                                if (42 < 249) {
                                                    var v89 uint = 255
                                                    var v90 uint32 = 2
                                                    _ = v89
                                                    _ = v90
                                                } else {
                                                    v91 := GenS8{F0: (v1 + ((v1 + (("hello" + (v1 + (("" + v1) + (v1 + ((v1 + v1) + "a"))))) + (v1 + v1))) + (v1 + v1)))}
                                                    var v92 bool = true
                                                    _ = v91
                                                    _ = v92
                                                }
                                                v93 := GenS9{F0: 1000, F1: (v2 | v2)}
                                                var v94 string = v1
                                                _ = v93
                                                _ = v94
                                            }
                                        default:
                                            v8 = -2.25
                                            v95 := GenS6{F0: 42, F1: 127}
                                            if (!("hello" == "")) {
                                                v96 := GenS10{F0: (v14 + (v14 * (v14 >> ((648) & 63))))}
                                                var v97 string = v1
                                                _ = v96
                                                _ = v97
                                            }
                                            _ = v95
                                        }
                                        v1 = "foo bar"
                                        v1 = v1
                                    } else {
                                        if (!(65535 != 68)) {
                                            v18 += (v14 * v18)
                                            var v98 uint64 = 634
                                            _ = v98
                                        } else {
                                            v2 = 2
                                            var v99 int8 = -128
                                            _ = v99
                                        }
                                        v2 -= v2
                                    }
                                default:
                                    v100 := GenS11{F0: (!true)}
                                    _ = v100
                                }
                            }
                            _ = v86
                        }
                    }
                default:
                    if true {
                        var v101 float64 = float64(v27)
                        var v102 int64 = 42
                        var v103 float64 = v101
                        _ = v101
                        _ = v102
                        _ = v103
                    } else {
                        var v104 uint64 = 42
                        fmt.Println(fi16)
                        switch fi13 {
                        case 128:
                            switch 0 {
                            case 899:
                                v2 = (^(v2))
                            case 128:
                                v8 = v8
                                var v105 uint32 = 2147483648
                                v2 += v2
                                _ = v105
                            }
                        case -485:
                            v3[fi16] = 2147483647
                        case 2147483647:
                            v104 += (^(v104))
                        default:
                            var v106 int32 = int32(v104)
                            _ = v106
                        }
                        _ = v104
                    }
                }
                var v107 uint8 = v2
                v4 = append(v4, 2147483648)
                _ = v107
                _ = v26
                _ = v27
            }
            if false {
                if false {
                    if len(v4) > 0 {
                        v4[0] = 2147483648
                    }
                    v108 := GenS11{F0: (517 > 685)}
                    v109 := GenS12{F0: "foo bar", F1: ((v1 + v1) + v1)}
                    _ = v108
                    _ = v109
                }
                var v110 uint = 65535
                v110 -= (v110 << ((42) & 63))
                _ = v110
            } else {
                var v111 uint32 = 57
                switch v1 {
                case "foo bar":
                    var v112 int32 = (v14 + (-(v14)))
                    _ = v112
                }
                v113 := GenS6{F0: 1, F1: 4294967295}
                _ = v111
                _ = v113
            }
            var v114 float64 = -2.25
            _ = fi16
            _ = v114
            _ = v17
            _ = v18
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    v2 = 255
    var v115 float32 = v8
    var v116 uint32 = 2
    if false {
        _ = len(v4)
        for fi117 := 0; fi117 < 12; fi117++ {
            var v118 float32 = float32(float64(1.5))
            _ = fi117
            _ = v118
        }
        fmt.Println(v1, v1)
    } else {
        for fi119 := 0; fi119 < 1; fi119++ {
            v1 = (((v1 + ("hello" + v1)) + v1) + (("hello" + ("\n\t" + ((v1 + "\n\t") + (v1 + (((("a" + v1) + ((v1 + "foo bar") + v1)) + (((v1 + v1) + (v1 + ((("foo bar" + v1) + v1) + v1))) + v1)) + ""))))) + (v1 + (v1 + ("\n\t" + "a")))))
            v2 = 128
            v8 = 1e10
            _ = fi119
        }
    }
    var v120 int = 154
    if ((v120 + (v120 | (v120 & (v120 >> ((42) & 63))))) > (v120 ^ v120)) {
        var v121 int64 = -425
        for fi122 := 0; fi122 < 20; fi122++ {
            v4 = append(v4, v116)
            for fi123 := 0; fi123 < 19; fi123++ {
                var v124 int8 = 42
                if ((v124 & v124) < (-(v124))) {
                    v125 := GenS3{F0: -709, F1: v120, F2: -32768}
                    _ = v125
                } else {
                    v120 = (fi123 * 42)
                    var v126 string = (("" + v1) + (((("a" + "hello") + v1) + ("hello" + v1)) + (v1 + v1)))
                    if true {
                        v127 := GenS9{F0: 276, F1: v2}
                        v126 = (v1 + "a")
                        if ((!((!(!false)) && (true && false))) && false) {
                            _ = len(v3)
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        _ = v127
                    }
                    _ = v126
                }
                _ = fi123
                _ = v124
            }
            v115 = 1.5
            _ = fi122
        }
        _ = v121
    } else {
        var v128 int64 = -624
        if true {
            v129 := GenS7{F0: 42}
            if (727 <= 143) {
                var v130 uint16 = 255
                _ = v130
            }
            _ = v129
        }
        _ = v128
    }
    fmt.Println("v1:", v1)
    fmt.Println("v115:", v115)
    fmt.Println("v116:", v116)
    fmt.Println("v120:", v120)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v115
    _ = v116
    _ = v120
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
