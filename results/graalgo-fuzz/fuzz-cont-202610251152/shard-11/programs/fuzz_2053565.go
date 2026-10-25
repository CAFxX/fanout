package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int32; F1 int32; F2 uint }
type GenS5 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS6 struct { F0 bool; F1 int32; F2 uint8 }
type GenS7 struct { F0 string; F1 uint16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint64; F1 int8 }
type GenS10 struct { F0 int; F1 uint16; F2 uint8 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint32; F1 uint32; F2 int }
type GenS13 struct { F0 uint32; F1 uint }
type GenS14 struct { F0 int; F1 int8; F2 uint16 }
func genFn0() bool {
    v7 := GenS1{F0: 1, F1: false}
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 string = ("" + "")
            if (730 == fi8) {
                v10 = (v10 + ("\n\t" + v10))
            }
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    var v11 int16 = 32767
    _ = v11
    _ = v7
    return (!true)
}
func genFn1() uint {
    if false {
        for fi12 := 0; fi12 < 12; fi12++ {
            for fi13 := 0; fi13 < 18; fi13++ {
                v14 := GenS1{F0: 97, F1: true}
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
        var v15 uint8 = 42
        _ = v15
    }
    var v16 int64 = -388
    v16 = (v16 % (((v16 % (((^(v16))) | 1))) | 1))
    for fi17 := 0; fi17 < 1; fi17++ {
        switch fi17 {
        case 42:
            if (!true) {
                v18 := GenS2{F0: (-(fi17)), F1: 1}
                _ = v18
            }
        case 38:
            v19 := []int{(fi17 / ((fi17) | 1)), (fi17 * fi17)}
            _ = v19
        case -245:
            v20 := GenS1{F0: 1, F1: (18446744073709551615 == 1)}
            _ = v20
        }
        var v21 int8 = -115
        v22 := GenS1{F0: 65535, F1: (((2 < 554) || true) || ((1e10 + 3.14159) >= (3.14159 + (0.0 + -2.25))))}
        _ = fi17
        _ = v21
        _ = v22
    }
    var v23 uint32 = 2147483647
    _ = v16
    _ = v23
    return 18446744073709551615
}

func main() {
    v1 := []int{-1, 9223372036854775807, 4294967295, -2147483648}
    var v2 int8 = -2
    v3 := []int{1, -1, 1}
    var v4 uint8 = 0
    var v5 float32 = 3.14159
    var v6 int16 = 32767
    var v24 int8 = (v2 << ((0) & 63))
    v25 := GenS1{F0: 853, F1: (!false)}
    if true {
        var v26 float64 = 3.14159
        _ = len(v3)
        for fi27 := 0; fi27 < 1; fi27++ {
            for fi28 := 0; fi28 < 12; fi28++ {
                var v29 int16 = 32767
                v29 = v6
                var v30 uint8 = (v4 / ((15) | 1))
                switch "foo bar" {
                case "":
                    var v31 int = (fi28 - (fi27 << ((4294967295) & 63)))
                    _ = v31
                case "a":
                    fmt.Println(v2)
                default:
                    v26 -= v26
                }
                v3 = append(v3, 32768)
                v26 = 0.1
                _ = fi28
                _ = v29
                _ = v30
            }
            var v32 int16 = (-(v6))
            _ = fi27
            _ = v32
        }
        _ = v26
    }
    switch "a" {
    case "\n\t":
        v24 = (v2 >> ((genFn1()) & 63))
    case "hello":
        v3 = append(v3, -2147483648)
    case "a":
        var v33 uint16 = 334
        _ = v33
    default:
        var v34 uint32 = 489
        _ = v34
    }
    if (true && (v6 >= (v6 + v6))) {
        var v35 uint32 = 0
        if false {
            if true {
                v36 := GenS2{F0: 0, F1: 766}
                v37 := GenS3{F0: 961, F1: -749}
                fmt.Println(v35, v4, v2)
                _ = v36
                _ = v37
            }
            var v38 int16 = (v6 - (-(v6)))
            v35 = (v35 - 65535)
            _ = v38
        } else {
            for fi39 := 0; fi39 < 12; fi39++ {
                fmt.Println(v2)
                _ = fi39
            }
            v40 := GenS4{F0: -1, F1: -2147483648, F2: 264}
            _ = v40
        }
        _ = v35
    } else {
        fmt.Println(v2, v24)
        if (708 == 32767) {
            for fi41 := 0; fi41 < 1; fi41++ {
                switch ("a" + ("\n\t" + "\n\t")) {
                case "\n\t":
                    v1 = append(v1, fi41)
                case "hello":
                    v42 := GenS5{F0: 18446744073709551615, F1: 1, F2: v6}
                    _ = v42
                case "a":
                    switch "foo bar" {
                    case "a":
                        if false {
                            if (2 > 213) {
                                v3 = append(v3, -979)
                                v24 -= (^(v24))
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = fi41
                                }
                                v6 = (^(v6))
                            }
                            v2 = (v24 * (-(v2)))
                            v5 -= v5
                        }
                    case "foo bar":
                        if ((2147483647 > -2147483648) || (!true)) {
                            v43 := GenS6{F0: (!false), F1: 0, F2: 0}
                            if (!true) {
                                var v44 int16 = v6
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if false {
                                    if len(v3) > 0 {
                                        v3[0] = (fi41 >> ((2) & 63))
                                    }
                                    switch "" {
                                    case "a":
                                        var v45 uint32 = 644
                                        _ = v45
                                    case "hello":
                                        var v46 int8 = v2
                                        _ = v46
                                    }
                                } else {
                                    for fi47 := 0; fi47 < 14; fi47++ {
                                        v48 := GenS7{F0: "", F1: 138}
                                        var v49 float64 = ((float64(0.0) - 1e10) - (-2.25 - ((float64(v5) + (1.5 - -2.25)) / float64(-2.25))))
                                        _ = fi47
                                        _ = v48
                                        _ = v49
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = (^(fi41))
                                    }
                                }
                                _ = v44
                            }
                            v50 := GenS8{F0: (v6 & v6)}
                            _ = v43
                            _ = v50
                        } else {
                            v4 = (^(v4))
                        }
                    default:
                        if (2 == genFn1()) {
                            v5 += -2.25
                            switch fi41 {
                            case 42:
                                if len(v1) > 0 {
                                    v1[0] = (fi41 ^ 973)
                                }
                            case -9223372036854775808:
                                if (v6 > -652) {
                                    if false {
                                        v51 := GenS3{F0: 0, F1: fi41}
                                        _ = v51
                                    } else {
                                        if (0 == -196) {
                                            if len(v1) > 0 {
                                                v1[0] = fi41
                                            }
                                            if (!("a" == ("a" + ""))) {
                                                var v52 uint32 = uint32(v4)
                                                if true {
                                                    if true {
                                                        for fi53 := 0; fi53 < 9; fi53++ {
                                                            var v54 float32 = (v5 - v5)
                                                            _ = fi53
                                                            _ = v54
                                                        }
                                                        var v55 uint64 = 171
                                                        _ = len(v3)
                                                        _ = v55
                                                    } else {
                                                        for fi56 := 0; fi56 < 7; fi56++ {
                                                            var v57 uint = 678
                                                            if true {
                                                                v58 := GenS8{F0: (v6 - v6)}
                                                                var v59 bool = true
                                                                _ = v58
                                                                _ = v59
                                                            } else {
                                                                _ = len(v1)
                                                                var v60 int32 = 235
                                                                var v61 bool = (1e10 > -2.25)
                                                                _ = v60
                                                                _ = v61
                                                            }
                                                            _ = fi56
                                                            _ = v57
                                                        }
                                                        v62 := GenS9{F0: 153, F1: (v2 / ((int8(v4)) | 1))}
                                                        _ = v62
                                                    }
                                                }
                                                for fi63 := 0; fi63 < 14; fi63++ {
                                                    v3 = append(v3, int(v24))
                                                    _ = fi63
                                                }
                                                _ = v52
                                            } else {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                var v64 uint8 = (v4 & v4)
                                                _ = v64
                                            }
                                            v5 = (((v5 / (v5 * 3.14159)) - v5) + -2.25)
                                        }
                                    }
                                } else {
                                    v65 := GenS10{F0: 0, F1: 53, F2: v4}
                                    if ((("" + "\n\t") + ("foo bar" + "")) > "\n\t") {
                                        v24 = v24
                                        if ((-(v6)) != v6) {
                                            if true {
                                                var v66 int32 = 2147483647
                                                if (float64(float32(3.14159)) > float64(-2.25)) {
                                                    v1 = append(v1, (fi41 >> ((2) & 63)))
                                                }
                                                _ = v66
                                            } else {
                                                switch fi41 {
                                                case -1:
                                                    v1 = append(v1, (fi41 << ((742) & 63)))
                                                case 32768:
                                                    for fi67 := 0; fi67 < 19; fi67++ {
                                                        var v68 int16 = v6
                                                        v69 := GenS9{F0: 1, F1: v2}
                                                        var v70 uint64 = 540
                                                        _ = fi67
                                                        _ = v68
                                                        _ = v69
                                                        _ = v70
                                                    }
                                                default:
                                                    var v71 int32 = -1
                                                    _ = v71
                                                }
                                                if false {
                                                    if len(v3) > 0 {
                                                        v3[0] = (fi41 - int(v6))
                                                    }
                                                } else {
                                                    var v72 uint8 = (v4 >> ((42) & 63))
                                                    for fi73 := 0; fi73 < 2; fi73++ {
                                                        fmt.Println(v6, v5, fi73)
                                                        v74 := GenS8{F0: (v6 << ((2) & 63))}
                                                        _ = fi73
                                                        _ = v74
                                                    }
                                                    _ = v72
                                                }
                                                for fi75 := 0; fi75 < 6; fi75++ {
                                                    v2 = (v24 / ((v2) | 1))
                                                    _ = fi75
                                                }
                                            }
                                            var v76 int = fi41
                                            _ = v76
                                        } else {
                                            for fi77 := 0; fi77 < 5; fi77++ {
                                                var v78 int8 = (-(v24))
                                                _ = fi77
                                                _ = v78
                                            }
                                            v6 = v6
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    } else {
                                        if true {
                                            switch -752 {
                                            case 240:
                                                if (255 > 379) {
                                                    var v79 string = "\n\t"
                                                    v5 += float32(float64(v5))
                                                    if (true && (-940 <= (fi41 * (fi41 >> ((65535) & 63))))) {
                                                        _ = len(v3)
                                                    }
                                                    _ = v79
                                                }
                                            case -9223372036854775808:
                                                var v80 int = (fi41 % ((1) | 1))
                                                _ = v80
                                            case 32768:
                                                for fi81 := 0; fi81 < 7; fi81++ {
                                                    var v82 int16 = (-(v6))
                                                    if len(v1) > 0 {
                                                        v1[0] = 128
                                                    }
                                                    _ = fi81
                                                    _ = v82
                                                }
                                            default:
                                                v6 += v6
                                            }
                                            switch 32767 {
                                            case 42:
                                                var v83 uint = 4294967295
                                                _ = v83
                                            default:
                                                fmt.Println(v5, v4)
                                            }
                                            v2 = (-(v24))
                                        }
                                    }
                                    _ = v65
                                }
                            case 128:
                                v84 := GenS10{F0: fi41, F1: 32767, F2: (v4 << ((255) & 63))}
                                _ = v84
                            default:
                                for fi85 := 0; fi85 < 20; fi85++ {
                                    if true {
                                        var v86 uint64 = 18446744073709551615
                                        if (!false) {
                                            v1 = append(v1, fi41)
                                        } else {
                                            v6 -= (v6 - (v6 - v6))
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            var v87 uint32 = 0
                                            _ = v87
                                        }
                                        _ = v86
                                    }
                                    _ = fi85
                                }
                            }
                        }
                    }
                default:
                    v6 = v6
                }
                _ = fi41
            }
            for fi88 := 0; fi88 < 14; fi88++ {
                if ("" != "") {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if true {
                        v89 := GenS4{F0: 0, F1: 0, F2: 18446744073709551615}
                        var v90 float64 = -2.25
                        _ = v89
                        _ = v90
                    }
                    if (!(false || true)) {
                        for fi91 := 0; fi91 < 8; fi91++ {
                            if (false || true) {
                                var v92 float64 = (((-2.25 * -2.25) - (3.14159 / ((-2.25 + 3.14159) + (0.0 / float64((v5 + 3.14159)))))) - (3.14159 / 3.14159))
                                v24 -= (v2 ^ int8(v4))
                                if (v92 >= v92) {
                                    if ((((0.0 / v5) == float32(0.1)) || (42 >= 1)) && false) {
                                        v93 := GenS6{F0: true, F1: -1, F2: v4}
                                        _ = v93
                                    } else {
                                        var v94 uint = 939
                                        _ = v94
                                    }
                                }
                                _ = v92
                            } else {
                                v95 := GenS1{F0: 255, F1: (v24 <= (v2 - 0))}
                                _ = v95
                            }
                            var v96 int8 = v24
                            _ = fi91
                            _ = v96
                        }
                        v97 := GenS6{F0: (!false), F1: -1, F2: 127}
                        _ = v97
                    } else {
                        var v98 float64 = -2.25
                        var v99 uint32 = 2
                        v2 = v24
                        _ = v98
                        _ = v99
                    }
                } else {
                    if true {
                        if (!false) {
                            fmt.Println(v2, v4)
                            for fi100 := 0; fi100 < 2; fi100++ {
                                v3 = append(v3, (^(fi100)))
                                _ = fi100
                            }
                        }
                        v6 = v6
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        for fi101 := 0; fi101 < 8; fi101++ {
                            if ((v24 * int8(v6)) != (v2 | (v24 & v24))) {
                                switch fi101 {
                                case 9223372036854775807:
                                    var v102 uint32 = 2147483647
                                    _ = v102
                                default:
                                    v103 := GenS11{F0: -1}
                                    _ = v103
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                fmt.Println(fi101, fi88, v6)
                            }
                            _ = fi101
                        }
                    }
                }
                fmt.Println(v2, v24, v4)
                v4 = (v4 - 163)
                _ = fi88
            }
            for fi104 := 0; fi104 < 1; fi104++ {
                if (32767 > v6) {
                    v105 := GenS12{F0: 1, F1: 4294967295, F2: fi104}
                    var v106 bool = (956 > 763)
                    _ = v105
                    _ = v106
                }
                v107 := GenS9{F0: 18446744073709551615, F1: (v24 * 111)}
                if false {
                    if (0 <= 56) {
                        fmt.Println(v2, v2, fi104)
                        _ = len(v1)
                        if ((true && (-1 == 0)) || (3.14159 == (float32(1e10) * 3.14159))) {
                            v108 := GenS13{F0: 42, F1: 748}
                            v109 := GenS6{F0: (42 > -32768), F1: 1, F2: v4}
                            _ = v108
                            _ = v109
                        }
                    }
                    for fi110 := 0; fi110 < 16; fi110++ {
                        if true {
                            v5 = 0.1
                        }
                        var v111 float32 = float32(((3.14159 * 0.1) * 1e10))
                        _ = fi110
                        _ = v111
                    }
                }
                _ = fi104
                _ = v107
            }
        }
    }
    v2 = v24
    for fi112 := 0; fi112 < 4; fi112++ {
        var v113 int64 = int64(v6)
        for fi114 := 0; fi114 < 8; fi114++ {
            if false {
                fmt.Println(v24)
            }
            if (v5 != (v5 + (v5 / v5))) {
                var v115 uint32 = 65535
                fmt.Println(v24, v5)
                var v116 int32 = 42
                _ = v115
                _ = v116
            }
            _ = fi114
        }
        _ = fi112
        _ = v113
    }
    if false {
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
        for fi117 := 0; fi117 < 18; fi117++ {
            var v118 uint64 = 291
            _ = fi117
            _ = v118
        }
        v119 := GenS14{F0: -2147483648, F1: 127, F2: 255}
        _ = v119
    } else {
        var v120 uint64 = 799
        if len(v1) > 0 {
            v1[0] = 127
        }
        if true {
            fmt.Println(v4, v5)
            var v121 uint64 = v120
            _ = v121
        }
        _ = v120
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
