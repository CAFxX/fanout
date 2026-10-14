package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint8; F2 int64 }
type GenS3 struct { F0 int32; F1 string; F2 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS6 struct { F0 uint; F1 uint; F2 uint16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS11 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS12 struct { F0 int; F1 uint32; F2 uint8 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 int32; F1 uint16; F2 string }

func main() {
    var v1 float32 = 1e10
    var v2 uint8 = 78
    v3 := make(map[string]uint32)
    var v4 bool = (9223372036854775807 >= -816)
    var v5 float64 = float64((-2.25 - v1))
    v6 := GenS1{F0: (v2 >> ((0) & 63))}
    v7 := GenS2{F0: 0, F1: (v2 * (v2 / ((2) | 1))), F2: -627}
    switch ("foo bar" + ("foo bar" + "hello")) {
    case "a":
        v8 := GenS3{F0: 981, F1: "a", F2: 18446744073709551615}
        _ = v8
    }
    var v9 uint32 = 0
    v1 = v1
    var v10 uint8 = 0
    for fi11 := 0; fi11 < 17; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            switch (fi12 - (fi12 % (((fi11 - fi11)) | 1))) {
            case 4294967295:
                v10 += v2
            case 42:
                switch (fi11 | fi12) {
                case 2147483647:
                    v2 -= (v10 >> ((500) & 63))
                case 1:
                    if (((v2 - (v10 - (^(v2)))) > v10) || false) {
                        if (v4 || (1 <= v10)) {
                            v13 := GenS3{F0: 1, F1: "\n\t", F2: 710}
                            switch fi12 {
                            case -99:
                                if ((v2 < 79) || v4) {
                                    fmt.Println(v5, v9, v10)
                                    var v14 int = fi11
                                    fmt.Println(fi12, fi11)
                                    var v15 uint16 = 42
                                    v16 := []int{}
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                } else {
                                    var v17 int32 = -683
                                    v5 = float64(-2.25)
                                    _ = v17
                                }
                            }
                            switch ("foo bar" + ("" + "hello")) {
                            case "":
                                v18 := []int{(fi12 / (((^(fi11))) | 1)), 0}
                                _ = v18
                            }
                            _ = v13
                        }
                    }
                }
            case -9223372036854775808:
                delete(v3, "hello")
            default:
                var v19 int16 = 240
                _ = v19
            }
            delete(v3, "")
            _ = fi12
        }
        v1 = (v1 - -2.25)
        _ = fi11
    }
    if false {
        v10 -= v2
    } else {
        fmt.Println(v2)
        v20 := []int{-2147483648, 127}
        _ = v20
    }
    var v21 float64 = 1.5
    if (127 == 42) {
        delete(v3, "hello")
        v22 := GenS4{F0: true}
        _ = v22
    }
    if (255 <= 153) {
        var v23 int16 = -553
        v24 := GenS5{F0: v2, F1: 18446744073709551615, F2: 0}
        v23 = v23
        _ = v23
        _ = v24
    } else {
        if (!(!(4294967295 != 42))) {
            fmt.Println(v4, v9, v9)
            v25 := GenS6{F0: 65535, F1: 32, F2: 359}
            _ = v25
        }
        v26 := GenS7{F0: 1}
        for fi27 := 0; fi27 < 16; fi27++ {
            for fi28 := 0; fi28 < 12; fi28++ {
                var v29 uint64 = 42
                v29 += 2
                _ = fi28
                _ = v29
            }
            for fi30 := 0; fi30 < 15; fi30++ {
                if v4 {
                    var v31 int8 = -128
                    if v4 {
                        if ((v1 <= 3.14159) && v4) {
                            v21 = v21
                            if true {
                                _ = len(v3)
                                fmt.Println(v2)
                            } else {
                                _ = len(v3)
                                var v32 bool = (42 <= 795)
                                _ = v32
                            }
                            v33 := []int{}
                            _ = v33
                        } else {
                            v34 := GenS4{F0: ((3.14159 + (1.5 * v5)) > v5)}
                            v3["\n\t"] = v9
                            _ = v34
                        }
                        if false {
                            v35 := GenS5{F0: v2, F1: 0, F2: (v31 * (v31 | (v31 >> ((540) & 63))))}
                            var v36 uint8 = (v2 + v2)
                            _ = v35
                            _ = v36
                        } else {
                            fmt.Println(v9, v9)
                            fmt.Println(v5, v4)
                        }
                        if v4 {
                            if v4 {
                                v37 := GenS8{F0: "foo bar"}
                                v38 := GenS1{F0: v10}
                                var v39 uint = 412
                                _ = v37
                                _ = v38
                                _ = v39
                            }
                            if (uint(v2) >= 1) {
                                var v40 float64 = (1e10 + (v21 + ((v21 * 0.1) - ((0.1 - (v21 - ((float64(v1) * v5) + ((v5 * v5) * v5)))) - v21))))
                                if (v4 && ((v9 * 1) <= (v9 - v9))) {
                                    var v41 uint = 456
                                    if (1 >= 2) {
                                        if (290 != 0) {
                                            if (!(1e10 <= (v40 / v40))) {
                                                if v4 {
                                                    var v42 bool = (v4 || (("a" + "") <= ("" + "hello")))
                                                    var v43 string = "foo bar"
                                                    _ = v42
                                                    _ = v43
                                                }
                                                if (-100 == (-(v31))) {
                                                    switch fi27 {
                                                    case -9223372036854775808:
                                                        if v4 {
                                                            fmt.Println(fi30, v21)
                                                            fmt.Println(fi27, v40, fi30)
                                                            if (296 < -225) {
                                                                var v44 float32 = (((v1 / v1) - v1) + (((v1 - v1) + (v1 / (v1 * (float32((v5 - v40)) - (1e10 + v1))))) + 0.1))
                                                                _ = v44
                                                            }
                                                            fmt.Println(v40, v41)
                                                        } else {
                                                            if v4 {
                                                                var v45 string = "hello"
                                                                v1 -= v1
                                                                var v46 uint32 = v9
                                                                _ = v45
                                                                _ = v46
                                                            } else {
                                                                switch (-(fi30)) {
                                                                case -9223372036854775808:
                                                                    var v47 int32 = 281
                                                                    _ = v47
                                                                case 179:
                                                                    var v48 int64 = 1
                                                                    if (!v4) {
                                                                        v41 -= (v41 ^ (v41 ^ 42))
                                                                        if (!v4) {
                                                                            switch "hello" {
                                                                            case "\n\t":
                                                                                v41 += v41
                                                                                switch 127 {
                                                                                case 32768:
                                                                                    var v49 int64 = -9223372036854775808
                                                                                    _ = v49
                                                                                case 1:
                                                                                    fmt.Println(v41, v2, v41)
                                                                                case 4294967295:
                                                                                    var v50 uint64 = 0
                                                                                    _ = v50
                                                                                }
                                                                            case "":
                                                                                var v51 uint = (v41 & v41)
                                                                                _ = v51
                                                                            case "hello":
                                                                                v48 = v48
                                                                            }
                                                                            v41 = (v41 / (((v41 | (v41 << (((v41 << (((v41 | v41)) & 63))) & 63)))) | 1))
                                                                            switch (fi27 + (fi30 << (((v41 | (v41 + v41))) & 63))) {
                                                                            case 42:
                                                                                v1 = 3.14159
                                                                                fmt.Println(fi30, v40, v40)
                                                                            case 127:
                                                                                v48 = (v48 / ((v48) | 1))
                                                                            default:
                                                                                if (65535 != 59) {
                                                                                    _ = len(v3)
                                                                                    var v52 uint = v41
                                                                                    _ = v52
                                                                                } else {
                                                                                    if (v2 < 255) {
                                                                                        delete(v3, "")
                                                                                        v48 += 42
                                                                                    }
                                                                                    var v53 uint = v41
                                                                                    fmt.Println(v4, v9)
                                                                                    _ = v53
                                                                                }
                                                                            }
                                                                            if (-32768 < 42) {
                                                                                var v54 float32 = ((((float32((float64(v1) * 0.0)) / 0.1) + (((v1 * (v1 * ((((3.14159 + (1e10 + v1)) - (1.5 * v1)) - (((float32(((-2.25 - v21) * (v5 - 1e10))) / 3.14159) * (v1 / 1.0)) - ((v1 * (3.14159 - v1)) / 1e10))) - v1))) + v1) / v1)) + (0.0 - (v1 / v1))) * v1)
                                                                                var v55 int64 = int64(v9)
                                                                                if v4 {
                                                                                    v1 = v1
                                                                                }
                                                                                _ = v54
                                                                                _ = v55
                                                                            }
                                                                            v56 := GenS3{F0: -2147483648, F1: "\n\t", F2: 637}
                                                                            _ = v56
                                                                        }
                                                                        v57 := []int{int(v9), fi30, fi30, (^(fi27))}
                                                                        _ = v57
                                                                    } else {
                                                                        fmt.Println(v31, v10)
                                                                        if v4 {
                                                                            v58 := []int{32768, (fi27 / (((fi30 >> (((v41 & 255)) & 63))) | 1)), fi27, (fi27 >> (((v41 & 255)) & 63))}
                                                                            if v4 {
                                                                                v41 += 522
                                                                                var v59 int = (fi27 + fi27)
                                                                                _ = v59
                                                                            } else {
                                                                                v60 := GenS2{F0: -475, F1: uint8(v9), F2: (v48 | (v48 % (((^(v48))) | 1)))}
                                                                                if len(v58) >= 2 {
                                                                                    v58 = v58[1:len(v58)]
                                                                                }
                                                                                v61 := GenS2{F0: -1, F1: v10, F2: (-(v48))}
                                                                                _ = v60
                                                                                _ = v61
                                                                            }
                                                                            fmt.Println(fi30)
                                                                            _ = v58
                                                                        } else {
                                                                            v3["\n\t"] = 872
                                                                            v62 := []int{fi27, fi27, (fi30 & fi30), (fi30 & (fi27 + -561))}
                                                                            if true {
                                                                                if v4 {
                                                                                    var v63 int16 = 510
                                                                                    _ = v63
                                                                                } else {
                                                                                    if len(v62) >= 2 {
                                                                                        v62 = v62[1:len(v62)]
                                                                                    }
                                                                                    var v64 float64 = 0.0
                                                                                    _ = v64
                                                                                }
                                                                            }
                                                                            _ = v62
                                                                        }
                                                                        v40 = ((v5 + ((((1.5 - 0.1) + 0.1) - v40) * ((float64(1e10) + (v40 + v40)) * (1e10 * v21)))) + v5)
                                                                    }
                                                                    _ = v48
                                                                case -12:
                                                                    v9 = uint32(v41)
                                                                default:
                                                                    var v65 int32 = 0
                                                                    _ = v65
                                                                }
                                                                fmt.Println(v21, v9)
                                                                v40 -= (0.0 - v40)
                                                                switch (fi27 ^ fi30) {
                                                                case 42:
                                                                    v40 -= (1.5 - v40)
                                                                case 127:
                                                                    switch fi30 {
                                                                    case 32767:
                                                                        v66 := GenS4{F0: v4}
                                                                        _ = v66
                                                                    default:
                                                                        var v67 float64 = v5
                                                                        _ = v67
                                                                    }
                                                                case 32767:
                                                                    var v68 float64 = float64((v1 - v1))
                                                                    _ = v68
                                                                default:
                                                                    fmt.Println(v5)
                                                                }
                                                            }
                                                            var v69 uint64 = 18446744073709551615
                                                            _ = v69
                                                        }
                                                    case 0:
                                                        var v70 uint32 = (v9 | (v9 / (((^(v9))) | 1)))
                                                        _ = v70
                                                    case 1:
                                                        var v71 uint32 = v9
                                                        _ = v71
                                                    }
                                                    if v4 {
                                                        var v72 uint8 = (v10 | v2)
                                                        var v73 uint8 = (v10 & 44)
                                                        _ = v72
                                                        _ = v73
                                                    } else {
                                                        switch (-(fi30)) {
                                                        case 1:
                                                            v21 -= 3.14159
                                                        case -9223372036854775808:
                                                            if ((^(v31)) > v31) {
                                                                if v4 {
                                                                    fmt.Println(v10, v21)
                                                                } else {
                                                                    v74 := GenS7{F0: 0}
                                                                    _ = v74
                                                                }
                                                            }
                                                        default:
                                                            v1 += (0.1 * v1)
                                                        }
                                                        v75 := []int{(^(fi27))}
                                                        _ = v75
                                                    }
                                                } else {
                                                    if (!true) {
                                                        fmt.Println(v1, v10, v2)
                                                    }
                                                }
                                            }
                                        }
                                        switch "" {
                                        case "foo bar":
                                            var v76 float32 = 0.0
                                            _ = v76
                                        case "a":
                                            v41 += v41
                                            var v77 int64 = 9223372036854775807
                                            switch (^(fi27)) {
                                            case 128:
                                                v41 += 861
                                            }
                                            _ = v77
                                        }
                                        v78 := GenS9{F0: (v41 / (((v41 | v41)) | 1))}
                                        if (239 <= -32768) {
                                            v79 := GenS10{F0: 255, F1: 990, F2: 96}
                                            fmt.Println(v41, v40, fi27)
                                            _ = v79
                                        } else {
                                            switch "" {
                                            case "":
                                                v41 = v41
                                            case "\n\t":
                                                if (211 < 2) {
                                                    fmt.Println(v2)
                                                } else {
                                                    v40 -= 1.5
                                                    var v80 uint16 = 2
                                                    _ = v80
                                                }
                                            }
                                            v4 = (!true)
                                        }
                                        var v81 int8 = (v31 << ((v41) & 63))
                                        var v82 int32 = 832
                                        v9 = v9
                                        _ = v78
                                        _ = v81
                                        _ = v82
                                    } else {
                                        if (69 <= int32(v41)) {
                                            switch fi30 {
                                            case -511:
                                                v31 += 42
                                            case -632:
                                                v83 := GenS11{F0: 619, F1: 0, F2: 32767}
                                                _ = v83
                                            case -2147483648:
                                                var v84 bool = ((fi27 >> ((v41) & 63)) <= fi30)
                                                _ = v84
                                            default:
                                                v40 = (v40 + v21)
                                            }
                                            v85 := GenS10{F0: 0, F1: -2147483648, F2: 213}
                                            v10 = (v10 % ((v10) | 1))
                                            if ((v4 && (!(723 == 137))) && (true && (573 < 491))) {
                                                v10 += v10
                                                var v86 int32 = 218
                                                _ = v86
                                            } else {
                                                v87 := []int{(fi27 % ((fi27) | 1)), 128}
                                                _ = v87
                                            }
                                            v88 := GenS9{F0: (v41 % ((18446744073709551615) | 1))}
                                            fmt.Println(v40, v10, v5)
                                            _ = v85
                                            _ = v88
                                        } else {
                                            var v89 int32 = -1
                                            _ = v89
                                        }
                                    }
                                    v90 := GenS2{F0: 1, F1: (v10 - v10), F2: int64(fi27)}
                                    _ = v41
                                    _ = v90
                                }
                                _ = v40
                            } else {
                                v91 := v3[("a" + "")]
                                v92 := []int{fi27, (-(fi30)), (-(fi30)), (fi27 + fi27)}
                                var v93 uint16 = 430
                                fmt.Println(v4)
                                _ = v91
                                _ = v92
                                _ = v93
                            }
                            v31 = v31
                            v94 := []int{(fi30 ^ (fi27 << ((1) & 63))), 2147483647, fi30}
                            v31 = v31
                            _ = v94
                        } else {
                            v5 += float64(1e10)
                            v31 = v31
                        }
                        switch "" {
                        case "hello":
                            if (v21 > (v5 / v5)) {
                                v3["hello"] = v9
                                var v95 uint32 = (v9 >> ((643) & 63))
                                if false {
                                    switch int(v95) {
                                    case 1:
                                        v96 := GenS5{F0: 244, F1: 42, F2: v31}
                                        _ = v96
                                    case 4294967295:
                                        v31 = (v31 | v31)
                                    default:
                                        switch -9223372036854775808 {
                                        case 127:
                                            var v97 uint64 = 746
                                            _ = v97
                                        }
                                    }
                                } else {
                                    v98 := GenS8{F0: "foo bar"}
                                    v95 += 1
                                    if v4 {
                                        var v99 int8 = int8(v10)
                                        var v100 uint64 = 4294967295
                                        var v101 uint16 = 32767
                                        _ = v100
                                        _ = v101
                                        _ = v99
                                    }
                                    _ = v98
                                }
                                _ = v95
                            } else {
                                v102 := GenS12{F0: (fi27 * fi30), F1: (v9 << ((1) & 63)), F2: 42}
                                var v103 float32 = v1
                                _ = v102
                                _ = v103
                            }
                        case "":
                            v10 = 2
                        default:
                            fmt.Println(fi30, v9)
                        }
                    } else {
                        fmt.Println(v10)
                        switch 128 {
                        case 4294967295:
                            if v4 {
                                v104 := []int{fi27}
                                v105 := GenS6{F0: 1, F1: 255, F2: 2}
                                var v106 int32 = -1
                                _ = v104
                                _ = v105
                                _ = v106
                            }
                            if (4294967295 != 272) {
                                if true {
                                    var v107 int16 = -221
                                    var v108 int32 = -889
                                    _ = v107
                                    _ = v108
                                } else {
                                    var v109 float64 = v21
                                    _ = v109
                                }
                                v10 = 0
                                if (v9 != (v9 >> ((784) & 63))) {
                                    if v4 {
                                        var v110 string = ""
                                        v111 := []int{1, (fi27 * (fi30 - (fi27 >> ((42) & 63)))), 32768, (fi30 << ((42) & 63))}
                                        _ = v110
                                        _ = v111
                                    }
                                    fmt.Println(v2)
                                }
                            } else {
                                fmt.Println(v10)
                                if (false && true) {
                                    v4 = (42 != 67)
                                    v112 := []int{fi30, (^(fi27)), (-(fi30)), 128}
                                    _ = v112
                                }
                                if ((v9 + (v9 | v9)) < v9) {
                                    var v113 uint32 = 2147483648
                                    fmt.Println(v2, fi30, v5)
                                    _ = v113
                                }
                            }
                        case 127:
                            switch (fi30 ^ fi27) {
                            case 9223372036854775807:
                                if (!v4) {
                                    v5 = v21
                                    if true {
                                        var v114 int8 = v31
                                        var v115 uint = uint(v9)
                                        _ = v114
                                        _ = v115
                                    }
                                } else {
                                    var v116 string = ""
                                    v117 := []int{fi30, (fi27 ^ (-(fi30))), (fi30 & (fi30 >> ((770) & 63)))}
                                    var v118 int32 = 679
                                    _ = v116
                                    _ = v117
                                    _ = v118
                                }
                            case 492:
                                var v119 uint64 = 2
                                _ = v119
                            case -2147483648:
                                v120 := GenS10{F0: 42, F1: -150, F2: 2}
                                _ = v120
                            default:
                                var v121 uint32 = v9
                                _ = v121
                            }
                        case 0:
                            _ = len(v3)
                        }
                        v2 -= 128
                        var v122 int16 = 482
                        if v4 {
                            if ((((!false) && (v122 != (v122 ^ (-(v122))))) && (868 != 1)) && v4) {
                                fmt.Println(v122)
                            } else {
                                if v4 {
                                    switch fi27 {
                                    case -744:
                                        v123 := GenS13{F0: (v10 ^ v2)}
                                        _ = v123
                                    }
                                } else {
                                    v21 = v21
                                }
                                v124 := GenS3{F0: -1, F1: ("foo bar" + "\n\t"), F2: 4294967295}
                                _ = v124
                            }
                        } else {
                            if true {
                                var v125 uint8 = 0
                                var v126 string = ("" + ("a" + ""))
                                _ = v125
                                _ = v126
                            } else {
                                switch ("\n\t" + "foo bar") {
                                case "hello":
                                    v5 -= 1e10
                                case "a":
                                    var v127 float64 = (0.1 * v5)
                                    _ = v127
                                case "\n\t":
                                    v128 := GenS14{F0: 199, F1: 2, F2: ("foo bar" + "\n\t")}
                                    switch "hello" {
                                    case "hello":
                                        var v129 int16 = (-(v122))
                                        _ = v129
                                    default:
                                        v1 -= v1
                                    }
                                    var v130 uint32 = (^(v9))
                                    _ = v128
                                    _ = v130
                                }
                                if v4 {
                                    var v131 int32 = 778
                                    _ = v131
                                } else {
                                    v132 := []int{(fi27 % ((fi27) | 1)), 1}
                                    var v133 int32 = -92
                                    _ = v132
                                    _ = v133
                                }
                                v134 := GenS4{F0: (v122 >= 1000)}
                                _ = v134
                            }
                            if false {
                                v4 = (629 > 616)
                                v4 = (0.0 < -2.25)
                            }
                            var v135 int32 = 1
                            _ = v135
                        }
                        _ = v122
                    }
                    var v136 uint32 = 317
                    _ = v136
                    _ = v31
                } else {
                    var v137 uint = 1
                    _ = v137
                }
                _ = fi30
            }
            _ = fi27
        }
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
