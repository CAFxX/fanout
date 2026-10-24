package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint; F1 uint32; F2 uint }
type GenS3 struct { F0 uint64; F1 uint64; F2 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint32; F1 uint16 }
type GenS7 struct { F0 int32; F1 int32; F2 bool }
type GenS8 struct { F0 string; F1 uint16; F2 uint64 }
type GenS9 struct { F0 int; F1 uint8 }
type GenS10 struct { F0 uint; F1 int }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int64; F1 uint16; F2 string }
type GenS13 struct { F0 int16; F1 int8 }

func main() {
    var v1 uint64 = 930
    v2 := make(map[uint32]bool)
    var v3 bool = (127 > 1)
    var v4 uint8 = 48
    v5 := make(map[int]bool)
    var v6 int16 = 1
    var v7 uint16 = 42
    v3 = (0 > 1)
    for fi8 := 0; fi8 < 14; fi8++ {
        fmt.Println(v3, v7, fi8)
        switch "" {
        case "a":
            for fi9 := 0; fi9 < 20; fi9++ {
                if v3 {
                    if true {
                        var v10 float64 = (0.0 * 0.1)
                        _ = v10
                    }
                    fmt.Println(v6, fi9, v4)
                } else {
                    v11 := GenS1{F0: v3, F1: 65535}
                    var v12 int32 = 130
                    _ = v11
                    _ = v12
                }
                v13 := []int{-1, 42, fi9, (fi9 - fi8)}
                _ = fi9
                _ = v13
            }
        case "foo bar":
            var v14 uint32 = uint32(v1)
            _ = v14
        case "\n\t":
            v4 = (v4 ^ v4)
        default:
            v15 := GenS2{F0: 4294967295, F1: 42, F2: 185}
            _ = v15
        }
        _ = fi8
    }
    v16 := GenS1{F0: true, F1: 2147483648}
    fmt.Println(v6)
    v6 += v6
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 int = fi17
        var v19 int = -213
        switch (v18 - (fi17 % (((v18 % (((v18 & 9223372036854775807)) | 1))) | 1))) {
        case -1:
            var v20 int16 = v6
            _ = v20
        default:
            fmt.Println(v1, v1)
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    v21 := GenS3{F0: (v1 >> ((0) & 63)), F1: 832, F2: ((v1 << ((478) & 63)) <= v1)}
    var v22 float32 = (1e10 * -2.25)
    switch -1 {
    case 276:
        v23 := v5[794]
        _ = v23
    case 42:
        v24 := GenS2{F0: 65535, F1: 42, F2: 1}
        _ = v24
    case 32768:
        var v25 int64 = -521
        _ = v25
    }
    v1 = v1
    for fi26 := 0; fi26 < 14; fi26++ {
        if (!true) {
            var v27 uint = 2
            if false {
                for fi28 := 0; fi28 < 4; fi28++ {
                    var v29 int8 = 1
                    if v3 {
                        if v3 {
                            var v30 int64 = 1
                            var v31 int8 = 1
                            _ = v30
                            _ = v31
                        } else {
                            _ = len(v5)
                        }
                        if (v7 >= (v7 | uint16(v4))) {
                            switch "" {
                            case "":
                                fmt.Println(v1, v7)
                            case "hello":
                                v27 = v27
                            case "foo bar":
                                var v32 float32 = v22
                                _ = v32
                            }
                            v2[0] = (v29 > (v29 & (v29 & (-(v29)))))
                        }
                    }
                    v2[4294967295] = v3
                    var v33 string = ("foo bar" + "\n\t")
                    if v3 {
                        v34 := v2[1]
                        v34 = v34
                        v33 = ("\n\t" + (((v33 + "a") + "hello") + v33))
                        if (v34 && v34) {
                            switch fi26 {
                            case 127:
                                if v34 {
                                    v35 := []int{(fi28 & (fi28 / ((22) | 1)))}
                                    v29 += v29
                                    if false {
                                        switch (fi26 - fi28) {
                                        case 674:
                                            var v36 uint8 = 42
                                            _ = v36
                                        case 2147483647:
                                            var v37 int = 32767
                                            _ = v37
                                        case 42:
                                            fmt.Println(v6, v7, v22)
                                            v38 := GenS4{F0: v33}
                                            _ = v38
                                        }
                                        switch v33 {
                                        case "foo bar":
                                            var v39 uint32 = 98
                                            _ = v39
                                        case "":
                                            v34 = true
                                        case "\n\t":
                                            if (!v3) {
                                                fmt.Println(v29, fi28)
                                                v1 = (v1 << (((v27 % (((v27 & (v27 >> ((v27) & 63)))) | 1))) & 63))
                                                v6 = v6
                                            } else {
                                                fmt.Println(v29, v1, v4)
                                                if (!(v3 && false)) {
                                                    if (float32(1e10) <= v22) {
                                                        if len(v35) > 0 {
                                                            v35[0] = 127
                                                        }
                                                        v5[fi26] = (!v3)
                                                    } else {
                                                        if (int8(fi28) == v29) {
                                                            v33 = (v33 + v33)
                                                        } else {
                                                            var v40 bool = (v34 || v34)
                                                            switch (-(fi28)) {
                                                            case 677:
                                                                v35 = append(v35, fi26)
                                                                if len(v35) > 0 {
                                                                    v35[0] = (fi28 * fi26)
                                                                }
                                                                if v40 {
                                                                    switch (fi26 & (-(fi28))) {
                                                                    case -1:
                                                                        switch v33 {
                                                                        case "a":
                                                                            fmt.Println(v29, v34, v29)
                                                                        }
                                                                    case 244:
                                                                        v7 = uint16(v4)
                                                                    default:
                                                                        var v41 bool = false
                                                                        _ = v41
                                                                    }
                                                                    v4 += (v4 ^ v4)
                                                                    switch fi26 {
                                                                    case 32768:
                                                                        switch (fi26 / ((42) | 1)) {
                                                                        case 32767:
                                                                            var v42 int = (fi28 & fi28)
                                                                            _ = v42
                                                                        case 71:
                                                                            v7 = 505
                                                                        }
                                                                    case -198:
                                                                        switch fi28 {
                                                                        case 32767:
                                                                            v6 -= v6
                                                                        default:
                                                                            fmt.Println(v40, v1, v34)
                                                                        }
                                                                    case 829:
                                                                        if true {
                                                                            fmt.Println(v27, v1)
                                                                            fmt.Println(v4)
                                                                        } else {
                                                                            var v43 int = (fi26 << ((v27) & 63))
                                                                            v40 = (v4 > (v4 & 0))
                                                                            var v44 string = v33
                                                                            _ = v43
                                                                            _ = v44
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v45 uint = v27
                                                                    var v46 float64 = -2.25
                                                                    var v47 int = (fi26 % (((fi26 * -2147483648)) | 1))
                                                                    v40 = false
                                                                    v48 := GenS3{F0: v1, F1: (v1 / (((v1 << ((v27) & 63))) | 1)), F2: (-133 >= -1)}
                                                                    _ = v45
                                                                    _ = v46
                                                                    _ = v47
                                                                    _ = v48
                                                                }
                                                            }
                                                            switch fi28 {
                                                            case -566:
                                                                v49 := GenS5{F0: 0}
                                                                _ = v49
                                                            default:
                                                                v1 = (v1 - 42)
                                                            }
                                                            _ = v40
                                                        }
                                                        if true {
                                                            v50 := GenS3{F0: (v1 << (((v27 / ((v27) | 1))) & 63)), F1: v1, F2: v3}
                                                            var v51 float32 = (v22 * v22)
                                                            v52 := GenS3{F0: (v1 ^ v1), F1: 1, F2: (v1 > (v1 >> (((v27 / (((v27 * 4294967295)) | 1))) & 63)))}
                                                            fmt.Println(v34, v33)
                                                            switch ((v33 + "") + v33) {
                                                            case "":
                                                                delete(v5, fi26)
                                                            case "a":
                                                                v29 -= 50
                                                                v1 += v1
                                                                var v53 bool = v34
                                                                _ = v53
                                                            default:
                                                                var v54 uint16 = (v7 * 2)
                                                                _ = v54
                                                            }
                                                            v55 := GenS6{F0: 2, F1: (v7 * (v7 >> (((v27 + (v27 & v27))) & 63)))}
                                                            _ = v50
                                                            _ = v51
                                                            _ = v52
                                                            _ = v55
                                                        } else {
                                                            if v3 {
                                                                var v56 int8 = 127
                                                                v56 += -94
                                                                _ = v56
                                                            } else {
                                                                if true {
                                                                    if v34 {
                                                                        fmt.Println(fi26, v34)
                                                                        v57 := GenS4{F0: "foo bar"}
                                                                        if len(v35) >= 2 {
                                                                            v35 = v35[1:len(v35)]
                                                                        }
                                                                        _ = v57
                                                                    }
                                                                }
                                                                v4 = (v4 * 95)
                                                            }
                                                            if (v33 != ((v33 + (("foo bar" + v33) + (((v33 + v33) + (("a" + v33) + v33)) + v33))) + ((v33 + ("foo bar" + v33)) + v33))) {
                                                                if v34 {
                                                                    v58 := GenS3{F0: (^(v1)), F1: (v1 << ((v27) & 63)), F2: (v34 || v3)}
                                                                    fmt.Println(fi26, fi26, v29)
                                                                    _ = v58
                                                                }
                                                                v33 = (v33 + (v33 + v33))
                                                                var v59 uint64 = 408
                                                                _ = v59
                                                            } else {
                                                                if v34 {
                                                                    v34 = (v34 || (v33 < (("\n\t" + v33) + v33)))
                                                                    if (732 > 42) {
                                                                        if true {
                                                                            var v60 int = 128
                                                                            _ = v60
                                                                        }
                                                                    } else {
                                                                        v61 := GenS7{F0: -849, F1: 2147483647, F2: (v3 || false)}
                                                                        switch fi28 {
                                                                        case 424:
                                                                            var v62 int8 = 1
                                                                            _ = v62
                                                                        }
                                                                        v1 -= v1
                                                                        v34 = false
                                                                        _ = v61
                                                                    }
                                                                    _ = len(v5)
                                                                    if ((^(v4)) >= v4) {
                                                                        v7 -= (v7 - 32767)
                                                                    } else {
                                                                        var v63 float64 = 0.1
                                                                        var v64 uint32 = 65535
                                                                        _ = v63
                                                                        _ = v64
                                                                    }
                                                                    v65 := GenS8{F0: (v33 + ""), F1: v7, F2: (^(v1))}
                                                                    fmt.Println(v3, v29)
                                                                    var v66 float64 = 3.14159
                                                                    var v67 bool = true
                                                                    _ = v65
                                                                    _ = v66
                                                                    _ = v67
                                                                }
                                                                switch (fi28 * -305) {
                                                                case 32767:
                                                                    _ = len(v35)
                                                                default:
                                                                    var v68 int8 = 0
                                                                    _ = v68
                                                                }
                                                                var v69 int64 = 9223372036854775807
                                                                _ = v69
                                                            }
                                                            fmt.Println(v22, v6, v27)
                                                        }
                                                    }
                                                    v34 = v3
                                                }
                                                var v70 uint = v27
                                                _ = v70
                                            }
                                        default:
                                            var v71 uint = 65535
                                            _ = v71
                                        }
                                        switch fi28 {
                                        case 382:
                                            _ = len(v5)
                                        case 32767:
                                            if (612 < -1) {
                                                v72 := GenS5{F0: 42}
                                                if (((!v34) && v3) || false) {
                                                    var v73 uint8 = v4
                                                    v27 = v27
                                                    if (("hello" + v33) >= "foo bar") {
                                                        if v34 {
                                                            var v74 string = "\n\t"
                                                            _ = v74
                                                        }
                                                        var v75 int32 = -614
                                                        v22 += 0.1
                                                        v29 = 43
                                                        switch -9223372036854775808 {
                                                        case 0:
                                                            _ = len(v2)
                                                            switch v33 {
                                                            case "a":
                                                                if (-113 <= -49) {
                                                                    var v76 string = "a"
                                                                    v77 := GenS3{F0: 480, F1: v1, F2: ((v1 * (v1 - 511)) >= 2)}
                                                                    _ = v76
                                                                    _ = v77
                                                                } else {
                                                                    v78 := GenS2{F0: (v27 << ((uint(v7)) & 63)), F1: 961, F2: v27}
                                                                    _ = v78
                                                                }
                                                            default:
                                                                var v79 int16 = (v6 * v6)
                                                                _ = v79
                                                            }
                                                        case 9223372036854775807:
                                                            fmt.Println(v7, v6, v34)
                                                        default:
                                                            if true {
                                                                if v3 {
                                                                    var v80 int64 = 601
                                                                    fmt.Println(v33, v7)
                                                                    var v81 uint64 = (v1 / ((945) | 1))
                                                                    _ = v80
                                                                    _ = v81
                                                                } else {
                                                                    v82 := GenS9{F0: int(v75), F1: v4}
                                                                    fmt.Println(v6, v75)
                                                                    _ = v82
                                                                }
                                                                v83 := GenS7{F0: (v75 / ((v75) | 1)), F1: (v75 % (((^(v75))) | 1)), F2: v3}
                                                                v84 := GenS10{F0: (v27 % ((v27) | 1)), F1: 128}
                                                                _ = v83
                                                                _ = v84
                                                            }
                                                        }
                                                        _ = v75
                                                    } else {
                                                        var v85 int8 = (v29 - (v29 * v29))
                                                        if true {
                                                            if v3 {
                                                                var v86 uint64 = (v1 + v1)
                                                                if false {
                                                                    v22 = ((1.5 / ((v22 * v22) / ((v22 - 0.1) / 1e10))) + v22)
                                                                    var v87 string = (v33 + (v33 + v33))
                                                                    v6 = (v6 >> (((v27 + (v27 & 2))) & 63))
                                                                    _ = v87
                                                                }
                                                                var v88 int32 = 42
                                                                _ = v86
                                                                _ = v88
                                                            }
                                                            switch (fi28 % ((440) | 1)) {
                                                            case -282:
                                                                v89 := GenS11{F0: fi28}
                                                                _ = v89
                                                            case 127:
                                                                v6 += v6
                                                            case -1:
                                                                var v90 uint8 = v73
                                                                _ = v90
                                                            default:
                                                                switch v33 {
                                                                case "a":
                                                                    var v91 uint8 = (v73 / (((v4 | 0)) | 1))
                                                                    _ = v91
                                                                case "":
                                                                    if (true || false) {
                                                                        if (3.14159 < 0.0) {
                                                                            fmt.Println(v7)
                                                                            var v92 int16 = (v6 | (v6 * v6))
                                                                            var v93 uint16 = 1
                                                                            _ = v92
                                                                            _ = v93
                                                                        } else {
                                                                            v3 = v34
                                                                            v35 = append(v35, int(v6))
                                                                            v94 := GenS12{F0: 393, F1: 1, F2: ""}
                                                                            _ = v94
                                                                        }
                                                                        v95 := GenS9{F0: (fi28 / ((fi28) | 1)), F1: (v4 / (((v4 / ((v73) | 1))) | 1))}
                                                                        switch v33 {
                                                                        case "foo bar":
                                                                            var v96 uint64 = v1
                                                                            _ = v96
                                                                        case "":
                                                                            if (9223372036854775807 > 9223372036854775807) {
                                                                                v35 = append(v35, fi28)
                                                                                v7 -= 370
                                                                                if v34 {
                                                                                    var v97 float32 = (0.1 / (0.1 + v22))
                                                                                    v97 = ((v97 + v97) * v22)
                                                                                    _ = v97
                                                                                }
                                                                            }
                                                                        case "\n\t":
                                                                            v33 = v33
                                                                        }
                                                                        _ = v95
                                                                    } else {
                                                                        v7 += v7
                                                                        var v98 uint32 = 412
                                                                        _ = v98
                                                                    }
                                                                }
                                                            }
                                                            var v99 bool = ((v27 & 255) > v27)
                                                            v100 := GenS1{F0: v34, F1: 774}
                                                            var v101 int16 = (v6 + (^(v6)))
                                                            _ = v100
                                                            _ = v101
                                                            _ = v99
                                                        }
                                                        v102 := GenS7{F0: 618, F1: -879, F2: (v1 <= (v1 % (((v1 - 18446744073709551615)) | 1)))}
                                                        var v103 string = (((v33 + "") + (v33 + v33)) + v33)
                                                        _ = v102
                                                        _ = v103
                                                        _ = v85
                                                    }
                                                    _ = v73
                                                }
                                                if ((v29 % ((-1) | 1)) == (v29 << (((v27 >> (((v27 % (((v27 | v27)) | 1))) & 63))) & 63))) {
                                                    var v104 int64 = 9223372036854775807
                                                    _ = len(v2)
                                                    _ = v104
                                                } else {
                                                    if (v3 || v34) {
                                                        v33 = (v33 + "hello")
                                                        v34 = v3
                                                        switch "a" {
                                                        case "a":
                                                            if len(v35) >= 2 {
                                                                v35 = v35[1:len(v35)]
                                                            }
                                                        case "foo bar":
                                                            if len(v35) >= 2 {
                                                                v35 = v35[1:len(v35)]
                                                            }
                                                            _ = len(v35)
                                                            var v105 float64 = 0.0
                                                            _ = v105
                                                        default:
                                                            v33 = (v33 + (v33 + v33))
                                                        }
                                                        var v106 string = "foo bar"
                                                        v2[2147483648] = (0.0 == float64(0.1))
                                                        _ = v106
                                                    }
                                                }
                                                _ = v72
                                            }
                                        case -277:
                                            if (!false) {
                                                var v107 uint32 = 2
                                                if v34 {
                                                    if ((v34 && v3) && true) {
                                                        fmt.Println(fi28, fi26)
                                                    }
                                                    v108 := GenS1{F0: true, F1: (v107 / ((v107) | 1))}
                                                    var v109 uint8 = (v4 & 2)
                                                    _ = v108
                                                    _ = v109
                                                }
                                                v110 := GenS13{F0: v6, F1: int8(v1)}
                                                var v111 uint16 = v7
                                                if ((v6 + v6) >= (v6 - -43)) {
                                                    v35 = append(v35, (fi28 | (fi28 >> (((^(v27))) & 63))))
                                                }
                                                _ = v107
                                                _ = v110
                                                _ = v111
                                            }
                                        }
                                    } else {
                                        v112 := GenS11{F0: 4294967295}
                                        v113 := v2[1]
                                        _ = v112
                                        _ = v113
                                    }
                                    _ = v35
                                } else {
                                    var v114 int16 = int16(v27)
                                    _ = v114
                                }
                            case 661:
                                var v115 float64 = 3.14159
                                _ = v115
                            case -427:
                                v116 := GenS1{F0: v34, F1: 2147483647}
                                _ = v116
                            default:
                                v117 := GenS11{F0: 9223372036854775807}
                                _ = v117
                            }
                        } else {
                            if ((v27 ^ uint(v1)) <= (v27 / ((v27) | 1))) {
                                fmt.Println(v7, fi26, v1)
                                if (v27 <= 0) {
                                    var v118 uint64 = v1
                                    _ = v118
                                } else {
                                    v119 := GenS9{F0: (fi26 ^ 0), F1: v4}
                                    _ = v119
                                }
                            } else {
                                var v120 float64 = 3.14159
                                v6 = (v6 | (-(v6)))
                                fmt.Println(v120)
                                _ = v120
                            }
                            if ((^(v27)) <= (v27 & v27)) {
                                var v121 uint64 = (v1 + v1)
                                var v122 int32 = int32(fi28)
                                v1 += (v1 >> (((v27 - (v27 << (((v27 % (((^(v27))) | 1))) & 63)))) & 63))
                                _ = v121
                                _ = v122
                            }
                            v123 := GenS9{F0: int(v7), F1: 255}
                            _ = v123
                        }
                        if (v3 && v3) {
                            fmt.Println(v6)
                        }
                        _ = v34
                    } else {
                        v124 := GenS13{F0: (v6 % (((v6 % (((v6 | -73)) | 1))) | 1)), F1: 1}
                        if (v6 <= v6) {
                            v3 = (!(!v3))
                            v3 = (!v3)
                            var v125 int32 = -2147483648
                            _ = v125
                        }
                        _ = v124
                    }
                    v126 := GenS1{F0: true, F1: uint32(v1)}
                    _ = fi28
                    _ = v126
                    _ = v29
                    _ = v33
                }
                var v127 bool = v3
                _ = v127
            }
            _ = v27
        }
        _ = fi26
    }
    var v128 int64 = -9223372036854775808
    switch ((((("hello" + "a") + (((("foo bar" + "a") + "hello") + ("foo bar" + ("hello" + "\n\t"))) + "hello")) + "hello") + "") + ((("" + "") + "foo bar") + ("hello" + "\n\t"))) {
    case "foo bar":
        var v129 string = (("foo bar" + ("a" + ("" + ""))) + ("\n\t" + ("" + (("hello" + (("a" + (("hello" + "a") + "hello")) + (((("\n\t" + "") + ("\n\t" + ("a" + ""))) + ("a" + "\n\t")) + "\n\t"))) + "\n\t"))))
        _ = v129
    case "hello":
        var v130 int8 = 0
        _ = v130
    }
    var v131 float32 = v22
    if ((true || v3) || (!(!(!true)))) {
        v4 = (v4 << ((4294967295) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v128:", v128)
    fmt.Println("v131:", v131)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v128
    _ = v131
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
