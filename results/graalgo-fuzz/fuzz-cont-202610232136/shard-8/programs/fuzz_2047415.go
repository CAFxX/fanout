package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 int64; F1 uint; F2 uint32 }
type GenS6 struct { F0 string; F1 int64; F2 uint8 }
type GenS7 struct { F0 uint; F1 string; F2 int16 }
type GenS8 struct { F0 uint; F1 int64 }
type GenS9 struct { F0 uint; F1 uint }
type GenS10 struct { F0 int64; F1 uint64; F2 string }
type GenS11 struct { F0 int; F1 int8; F2 int16 }
type GenS12 struct { F0 int64; F1 int16; F2 bool }
func genFn0(p0 int32, p1 bool, p2 int64) int8 {
    v6 := []int{60, -1, 4294967295, 32768}
    var v7 int8 = 42
    if p1 {
        _ = len(v6)
        var v8 uint16 = 32767
        var v9 string = ""
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (-(v7))
}

func main() {
    var v1 bool = (false && ((!true) || true))
    v2 := make(map[int]string)
    var v3 int8 = -1
    var v4 float32 = 3.14159
    v5 := make(map[uint32]uint32)
    v1 = (!false)
    var v10 string = "foo bar"
    v11 := []int{0, 32768}
    if v1 {
        _ = len(v5)
        _ = len(v2)
    }
    fmt.Println(v4, v10, v4)
    var v12 string = v10
    if (255 != 146) {
        fmt.Println(v10, v1, v3)
    } else {
        for fi13 := 0; fi13 < 1; fi13++ {
            if (142 >= 668) {
                if true {
                    v12 = "\n\t"
                    v1 = v1
                } else {
                    v3 = genFn0(-1, false, 9223372036854775807)
                }
                v10 = v10
            } else {
                v12 = v12
                fmt.Println(v10, v12, v4)
            }
            for fi14 := 0; fi14 < 11; fi14++ {
                if (0 >= 255) {
                    if true {
                        var v15 int = (fi13 % (((fi13 + (fi13 - fi13))) | 1))
                        if v1 {
                            if v1 {
                                var v16 bool = v1
                                var v17 float64 = -2.25
                                var v18 int16 = 1000
                                _ = v16
                                _ = v17
                                _ = v18
                            } else {
                                var v19 uint8 = 110
                                var v20 int = fi13
                                v21 := GenS1{F0: 2, F1: 286}
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                            if v1 {
                                v4 -= v4
                                if (v3 != -34) {
                                    v1 = (518 > 4294967295)
                                    if len(v11) > 0 {
                                        v11[0] = 334
                                    }
                                    switch v10 {
                                    case "a":
                                        fmt.Println(v12, fi13, v4)
                                    case "":
                                        fmt.Println(v10)
                                    }
                                    if v1 {
                                        if (v1 || true) {
                                            switch v10 {
                                            case "":
                                                v22 := GenS1{F0: 4294967295, F1: 705}
                                                _ = v22
                                            case "hello":
                                                var v23 uint8 = 1
                                                _ = v23
                                            }
                                            v3 = (v3 + (v3 * (-(v3))))
                                            var v24 int8 = v3
                                            _ = v24
                                        } else {
                                            v5[65535] = 0
                                            var v25 uint = 2
                                            _ = v25
                                        }
                                    }
                                } else {
                                    if (!(-472 <= 4294967295)) {
                                        v26 := GenS2{F0: 0, F1: 691}
                                        switch 32768 {
                                        case -2147483648:
                                            if len(v11) >= 2 {
                                                v11 = v11[1:len(v11)]
                                            }
                                        default:
                                            if len(v11) >= 2 {
                                                v11 = v11[1:len(v11)]
                                            }
                                        }
                                        var v27 int8 = (v3 ^ 0)
                                        _ = v26
                                        _ = v27
                                    } else {
                                        v5[1] = 2147483647
                                        var v28 uint8 = 0
                                        if v1 {
                                            switch (fi13 & fi13) {
                                            case 291:
                                                var v29 int = (v15 ^ fi13)
                                                _ = v29
                                            default:
                                                v4 += ((v4 - v4) / v4)
                                            }
                                            switch (fi14 * (fi13 & (fi14 ^ 9223372036854775807))) {
                                            case -732:
                                                if v1 {
                                                    v11 = append(v11, (-(v15)))
                                                    _ = len(v5)
                                                }
                                            default:
                                                v11 = append(v11, -888)
                                            }
                                            v30 := GenS3{F0: "foo bar", F1: 654}
                                            _ = v30
                                        } else {
                                            var v31 uint = 680
                                            fmt.Println(v3, fi13)
                                            switch v10 {
                                            case "hello":
                                                v32 := GenS4{F0: v1, F1: 129}
                                                _ = v32
                                            case "\n\t":
                                                switch 320 {
                                                case 0:
                                                    _ = len(v2)
                                                case 453:
                                                    var v33 float32 = v4
                                                    _ = v33
                                                default:
                                                    var v34 uint32 = 65535
                                                    if (v4 <= 3.14159) {
                                                        if ((v15 / ((fi13) | 1)) != fi14) {
                                                            if ((v31 << ((0) & 63)) != (v31 ^ (v31 % ((v31) | 1)))) {
                                                                v11 = append(v11, (v15 - fi13))
                                                                fmt.Println(v4, v15)
                                                            }
                                                            v35 := v2[4294967295]
                                                            switch (fi13 % ((fi14) | 1)) {
                                                            case 439:
                                                                if false {
                                                                    if len(v11) >= 2 {
                                                                        v11 = v11[1:len(v11)]
                                                                    }
                                                                } else {
                                                                    var v36 float64 = ((0.0 - (0.0 - float64(1e10))) / (0.1 * ((((3.14159 / 1.5) / 1.0) / (((-2.25 * 1.5) / (3.14159 * 3.14159)) * 1.5)) * -2.25)))
                                                                    v28 += v28
                                                                    if true {
                                                                        fmt.Println(fi14, v28)
                                                                        var v37 float64 = (float64(0.1) * ((3.14159 / v36) - ((v36 - (float64(v4) / (v36 / float64((float32(((v36 / (v36 / (v36 / v36))) - (v36 + 0.1))) + v4))))) - ((3.14159 * v36) * (0.0 / float64(v4))))))
                                                                        _ = v37
                                                                    } else {
                                                                        v4 -= v4
                                                                    }
                                                                    _ = v36
                                                                }
                                                            }
                                                            _ = v35
                                                        }
                                                        v28 += v28
                                                    }
                                                    switch (v12 + v12) {
                                                    case "hello":
                                                        fmt.Println(v1, fi13)
                                                    case "":
                                                        var v38 float32 = 0.1
                                                        _ = v38
                                                    }
                                                    v39 := GenS3{F0: v12, F1: (^(v34))}
                                                    _ = v34
                                                    _ = v39
                                                }
                                            case "foo bar":
                                                v28 = (v28 - (v28 & (v28 | (v28 ^ v28))))
                                                if ((2 > 2) || v1) {
                                                    v40 := GenS2{F0: 9223372036854775807, F1: 32767}
                                                    v31 -= 206
                                                    _ = v40
                                                } else {
                                                    var v41 uint = (^(v31))
                                                    _ = v41
                                                }
                                            }
                                            fmt.Println(v3)
                                            var v42 int8 = (v3 / ((genFn0(-2147483648, ((v31 < 2) || v1), 42)) | 1))
                                            _ = v31
                                            _ = v42
                                        }
                                        v43 := GenS5{F0: 9223372036854775807, F1: 525, F2: 4294967295}
                                        _ = v28
                                        _ = v43
                                    }
                                    switch ((((("foo bar" + (v10 + v10)) + v12) + ((("\n\t" + ("foo bar" + "\n\t")) + v12) + v12)) + (v12 + "hello")) + v10) {
                                    case "":
                                        switch (v10 + ("hello" + v12)) {
                                        case "hello":
                                            fmt.Println(v15)
                                        case "\n\t":
                                            v1 = (!(false || (1 != 39)))
                                        case "":
                                            switch v10 {
                                            case "hello":
                                                v10 = (v12 + v12)
                                            case "\n\t":
                                                if true {
                                                    if true {
                                                        v44 := GenS5{F0: 0, F1: 578, F2: 266}
                                                        var v45 int32 = int32(v3)
                                                        var v46 int8 = -95
                                                        var v47 bool = v1
                                                        var v48 uint16 = 492
                                                        switch (fi13 | 21) {
                                                        case 2147483647:
                                                            v49 := GenS6{F0: "", F1: 1, F2: 1}
                                                            _ = v49
                                                        }
                                                        v3 -= v3
                                                        if v1 {
                                                            if false {
                                                                if ((3.14159 - v4) >= (v4 + v4)) {
                                                                    v50 := GenS7{F0: 255, F1: "a", F2: 42}
                                                                    _ = v50
                                                                } else {
                                                                    if (("hello" + "foo bar") <= v10) {
                                                                        if v1 {
                                                                            v11 = append(v11, 2147483647)
                                                                            fmt.Println(v47)
                                                                        } else {
                                                                            _ = len(v2)
                                                                            var v51 string = ((v10 + v10) + v12)
                                                                            if (1 >= 86) {
                                                                                var v52 float64 = 3.14159
                                                                                fmt.Println(v10)
                                                                                _ = v52
                                                                            }
                                                                            _ = v51
                                                                        }
                                                                        switch (v10 + (v12 + v10)) {
                                                                        case "hello":
                                                                            v53 := GenS6{F0: "", F1: -9223372036854775808, F2: 1}
                                                                            _ = v53
                                                                        case "\n\t":
                                                                            if (!false) {
                                                                                var v54 int16 = 32767
                                                                                var v55 float64 = (((float64(0.0) - 0.1) - 1e10) / float64(0.1))
                                                                                v3 = int8(v54)
                                                                                _ = v54
                                                                                _ = v55
                                                                            } else {
                                                                                v56 := GenS8{F0: 255, F1: 543}
                                                                                if true {
                                                                                    v57 := v2[fi13]
                                                                                    _ = v57
                                                                                } else {
                                                                                    v15 += 32767
                                                                                    var v58 float64 = (0.0 + ((0.0 * 3.14159) / -2.25))
                                                                                    if len(v11) > 0 {
                                                                                        v11[0] = (fi14 | (fi14 * (fi13 & (v15 / (((fi13 << ((18446744073709551615) & 63))) | 1)))))
                                                                                    }
                                                                                    if v1 {
                                                                                        if v47 {
                                                                                            v59 := v5[133]
                                                                                            switch v10 {
                                                                                            case "":
                                                                                                if v47 {
                                                                                                    var v60 uint = 255
                                                                                                    _ = v60
                                                                                                } else {
                                                                                                    v1 = true
                                                                                                    var v61 int32 = v45
                                                                                                    if (false && v1) {
                                                                                                        var v62 uint64 = 4294967295
                                                                                                        _ = v62
                                                                                                    }
                                                                                                    _ = v61
                                                                                                }
                                                                                            }
                                                                                            v4 -= v4
                                                                                            _ = v59
                                                                                        }
                                                                                        var v63 uint16 = (v48 >> ((494) & 63))
                                                                                        fmt.Println(fi14, v12)
                                                                                        if true {
                                                                                            var v64 uint8 = 2
                                                                                            switch v10 {
                                                                                            case "a":
                                                                                                fmt.Println(v4, v10)
                                                                                            default:
                                                                                                if true {
                                                                                                    v12 = (v10 + v12)
                                                                                                    if (!(true && (v58 > (v58 + (v58 - v58))))) {
                                                                                                        var v65 uint8 = v64
                                                                                                        v1 = (1 < 2147483647)
                                                                                                        v66 := GenS7{F0: 2, F1: (v12 + "hello"), F2: -292}
                                                                                                        _ = v65
                                                                                                        _ = v66
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            _ = v64
                                                                                        }
                                                                                        v5[376] = 2
                                                                                        var v67 string = "foo bar"
                                                                                        fmt.Println(v1, fi13)
                                                                                        _ = v63
                                                                                        _ = v67
                                                                                    } else {
                                                                                        var v68 uint64 = 2
                                                                                        var v69 int64 = 0
                                                                                        _ = v68
                                                                                        _ = v69
                                                                                    }
                                                                                    _ = v58
                                                                                }
                                                                                _ = v56
                                                                            }
                                                                        default:
                                                                            fmt.Println(v4)
                                                                        }
                                                                        if len(v11) >= 2 {
                                                                            v11 = v11[1:len(v11)]
                                                                        }
                                                                    } else {
                                                                        if ((v4 * 0.0) <= v4) {
                                                                            fmt.Println(v45, v48, fi14)
                                                                        } else {
                                                                            var v70 uint64 = 18446744073709551615
                                                                            _ = v70
                                                                        }
                                                                        fmt.Println(v45)
                                                                        var v71 int32 = (v45 ^ 2147483647)
                                                                        _ = v71
                                                                    }
                                                                    var v72 uint16 = 2
                                                                    var v73 uint32 = 2147483647
                                                                    if v1 {
                                                                        switch fi14 {
                                                                        case 42:
                                                                            var v74 uint8 = 0
                                                                            _ = v74
                                                                        case -1:
                                                                            if (342 == 988) {
                                                                                v11 = append(v11, (fi13 & v15))
                                                                            }
                                                                        case -281:
                                                                            if (!(2 < 2)) {
                                                                                var v75 int16 = 858
                                                                                v76 := GenS8{F0: 42, F1: 9223372036854775807}
                                                                                if len(v11) >= 2 {
                                                                                    v11 = v11[1:len(v11)]
                                                                                }
                                                                                v77 := GenS7{F0: 4294967295, F1: "hello", F2: (v75 >> ((398) & 63))}
                                                                                _ = v75
                                                                                _ = v76
                                                                                _ = v77
                                                                            }
                                                                        default:
                                                                            v78 := GenS4{F0: ((v47 || true) || v1), F1: 42}
                                                                            _ = v78
                                                                        }
                                                                    } else {
                                                                        v79 := GenS9{F0: 93, F1: 255}
                                                                        _ = len(v11)
                                                                        _ = len(v5)
                                                                        _ = v79
                                                                    }
                                                                    _ = v72
                                                                    _ = v73
                                                                }
                                                            }
                                                            v80 := GenS5{F0: 0, F1: 4294967295, F2: 4294967295}
                                                            v15 -= v15
                                                            _ = v80
                                                        } else {
                                                            v2[127] = "foo bar"
                                                        }
                                                        _ = v44
                                                        _ = v45
                                                        _ = v46
                                                        _ = v47
                                                        _ = v48
                                                    } else {
                                                        v12 = ""
                                                        var v81 int32 = -1
                                                        var v82 int32 = 0
                                                        v83 := v5[455]
                                                        var v84 int64 = -909
                                                        var v85 uint8 = 128
                                                        _ = v81
                                                        _ = v82
                                                        _ = v83
                                                        _ = v84
                                                        _ = v85
                                                    }
                                                    v12 = v12
                                                    switch v15 {
                                                    case 0:
                                                        v15 = 2147483647
                                                    case -1:
                                                        fmt.Println(fi13, v15)
                                                    case 32768:
                                                        var v86 int32 = 460
                                                        _ = v86
                                                    default:
                                                        if true {
                                                            v10 = ("" + "\n\t")
                                                        }
                                                    }
                                                }
                                            case "foo bar":
                                                v15 += v15
                                            default:
                                                var v87 float32 = ((float32((float64(1.5) / 1.5)) / v4) * v4)
                                                _ = v87
                                            }
                                        }
                                    case "a":
                                        if v1 {
                                            var v88 float32 = v4
                                            v89 := GenS10{F0: 9223372036854775807, F1: 18446744073709551615, F2: v10}
                                            _ = v88
                                            _ = v89
                                        } else {
                                            fmt.Println(v15, v15)
                                            switch int(v3) {
                                            case 32767:
                                                switch "foo bar" {
                                                case "hello":
                                                    v90 := GenS7{F0: 0, F1: v10, F2: 1000}
                                                    _ = v90
                                                case "foo bar":
                                                    var v91 float32 = v4
                                                    _ = v91
                                                default:
                                                    v92 := GenS11{F0: (fi13 / ((127) | 1)), F1: v3, F2: 42}
                                                    _ = v92
                                                }
                                            case -2147483648:
                                                delete(v5, 41)
                                            case -9223372036854775808:
                                                if (!(-912 > 801)) {
                                                    if len(v11) > 0 {
                                                        v11[0] = (fi14 - (fi13 << ((1) & 63)))
                                                    }
                                                    v3 = -1
                                                }
                                            }
                                        }
                                    default:
                                        fmt.Println(v10, v12, v10)
                                    }
                                }
                            } else {
                                switch v10 {
                                case "":
                                    var v93 int32 = -2147483648
                                    _ = v93
                                case "a":
                                    v94 := GenS4{F0: (1 > 0), F1: 656}
                                    _ = v94
                                case "\n\t":
                                    v95 := GenS12{F0: 42, F1: -32768, F2: true}
                                    _ = v95
                                }
                                v12 = v12
                                switch fi13 {
                                case 127:
                                    var v96 uint16 = 1
                                    _ = v96
                                case 42:
                                    var v97 int16 = 1000
                                    _ = v97
                                case 9223372036854775807:
                                    if v1 {
                                        var v98 float32 = (v4 * float32(0.0))
                                        if (v1 || (v10 < ("a" + ("foo bar" + v12)))) {
                                            if false {
                                                if ("" >= v10) {
                                                    v15 = fi13
                                                }
                                            }
                                            switch v12 {
                                            case "foo bar":
                                                if len(v11) > 0 {
                                                    v11[0] = fi14
                                                }
                                            case "hello":
                                                if (v1 || true) {
                                                    var v99 uint32 = 1
                                                    v2[fi14] = ("\n\t" + (("\n\t" + v12) + "foo bar"))
                                                    if (208 == 35) {
                                                        var v100 uint32 = v99
                                                        fmt.Println(v98, v98)
                                                        _ = v100
                                                    } else {
                                                        fmt.Println(v98, v1, v10)
                                                    }
                                                    _ = v99
                                                }
                                            case "a":
                                                var v101 int = 1
                                                _ = v101
                                            default:
                                                v11 = append(v11, (-(v15)))
                                            }
                                            _ = len(v5)
                                            var v102 int32 = 2147483647
                                            _ = v102
                                        } else {
                                            v12 = v10
                                            var v103 int = fi14
                                            _ = v103
                                        }
                                        v104 := GenS6{F0: v10, F1: 0, F2: 61}
                                        _ = v104
                                        _ = v98
                                    } else {
                                        fmt.Println(v1, v15)
                                    }
                                default:
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                }
                                if true {
                                    v105 := GenS12{F0: 1, F1: 0, F2: ((v10 + v12) >= (v10 + (("" + "\n\t") + "")))}
                                    _ = len(v2)
                                    _ = v105
                                } else {
                                    fmt.Println(v15, v4)
                                    var v106 int16 = 1
                                    if v1 {
                                        var v107 bool = (127 > 2)
                                        _ = v107
                                    }
                                    _ = v106
                                }
                                var v108 uint8 = 2
                                _ = v108
                            }
                            var v109 bool = ((2 != 21) && v1)
                            _ = v109
                        }
                        v11 = append(v11, fi14)
                        _ = v15
                    }
                }
                v3 = (-(v3))
                v2[32768] = ((("hello" + v12) + v10) + ((v10 + "\n\t") + v10))
                _ = fi14
            }
            var v110 uint = 191
            _ = fi13
            _ = v110
        }
        var v111 int32 = -1
        var v112 uint64 = 1
        _ = v111
        _ = v112
    }
    v1 = true
    switch 32767 {
    case 32767:
        var v113 uint8 = 127
        _ = v113
    case 9223372036854775807:
        fmt.Println(v12, v1)
    case 32768:
        var v114 bool = v1
        _ = v114
    }
    fmt.Println(v3, v1, v3)
    v115 := v5[0]
    if (42 > 2147483647) {
        v4 += v4
        for fi116 := 0; fi116 < 1; fi116++ {
            if len(v11) > 0 {
                v11[0] = (fi116 >> ((0) & 63))
            }
            v117 := v2[fi116]
            _ = fi116
            _ = v117
        }
        if (-1 != 1000) {
            v12 = v12
            v10 = v10
        }
    }
    var v118 uint8 = 127
    var v119 int32 = -676
    var v120 uint = 0
    switch (v12 + v10) {
    case "\n\t":
        if false {
            v1 = v1
            var v121 bool = true
            if v1 {
                v122 := GenS3{F0: "a", F1: v115}
                v118 = v118
                var v123 string = "a"
                _ = v122
                _ = v123
            }
            _ = v121
        }
    }
    var v124 int8 = (v3 * 100)
    for fi125 := 0; fi125 < 6; fi125++ {
        v115 += (^(v115))
        var v126 uint32 = (^(v115))
        fmt.Println(v120, v126)
        _ = fi125
        _ = v126
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v115:", v115)
    fmt.Println("v118:", v118)
    fmt.Println("v119:", v119)
    fmt.Println("v12:", v12)
    fmt.Println("v120:", v120)
    fmt.Println("v124:", v124)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v115
    _ = v118
    _ = v119
    _ = v12
    _ = v120
    _ = v124
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
