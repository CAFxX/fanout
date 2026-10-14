package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int8 }
type GenS2 struct { F0 uint; F1 uint64; F2 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 int32; F2 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint32; F1 int32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int64; F1 uint; F2 int8 }
type GenS11 struct { F0 uint8; F1 uint16; F2 string }
type GenS12 struct { F0 bool; F1 uint8 }
type GenS13 struct { F0 int32; F1 uint8; F2 int }
type GenS14 struct { F0 int64; F1 string; F2 int8 }
func genFn0(p0 int16, p1 uint32, p2 int64) string {
    if false {
        v4 := []int{0, -190}
        var v5 uint16 = 2
        _ = v4
        _ = v5
    } else {
        var v6 int = 32767
        _ = v6
    }
    var v7 int64 = (p2 | (p2 + p2))
    fmt.Println(v7, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "a"
}

func main() {
    var v1 uint = 65535
    var v2 uint16 = 525
    var v3 float32 = 1e10
    for fi8 := 0; fi8 < 14; fi8++ {
        switch (fi8 ^ 455) {
        case 0:
            if false {
                switch "a" {
                case "hello":
                    if false {
                        var v9 uint8 = 71
                        v10 := GenS1{F0: 42, F1: true, F2: 42}
                        _ = v10
                        _ = v9
                    } else {
                        for fi11 := 0; fi11 < 10; fi11++ {
                            switch "foo bar" {
                            case "\n\t":
                                switch "a" {
                                case "\n\t":
                                    v2 = (^(v2))
                                case "a":
                                    var v12 int32 = 431
                                    _ = v12
                                default:
                                    v13 := GenS1{F0: 118, F1: false, F2: -9}
                                    _ = v13
                                }
                            }
                            if (!((!false) && (!((-(fi11)) > fi11)))) {
                                var v14 int8 = -128
                                fmt.Println(v14, v2, v3)
                                v15 := []int{}
                                _ = v14
                                _ = v15
                            } else {
                                var v16 uint16 = v2
                                _ = v16
                            }
                            _ = fi11
                        }
                    }
                case "foo bar":
                    v1 = v1
                }
                v2 = (v2 >> ((v1) & 63))
                for fi17 := 0; fi17 < 14; fi17++ {
                    var v18 string = ""
                    _ = fi17
                    _ = v18
                }
            }
        case 2147483647:
            v1 -= (v1 * 499)
        case 9223372036854775807:
            v19 := []int{}
            _ = v19
        }
        var v20 bool = (((("" + ("hello" + (("hello" + ((("\n\t" + "foo bar") + "") + "")) + "foo bar"))) + (("hello" + "") + "\n\t")) + ((("a" + "\n\t") + "") + "foo bar")) > "\n\t")
        v1 = (v1 ^ 2)
        _ = fi8
        _ = v20
    }
    switch "a" {
    case "hello":
        var v21 uint64 = 334
        _ = v21
    case "":
        v22 := GenS1{F0: 100, F1: (-2.25 == 1e10), F2: -128}
        _ = v22
    case "\n\t":
        for fi23 := 0; fi23 < 5; fi23++ {
            if false {
                v24 := GenS2{F0: (v1 % ((255) | 1)), F1: 42, F2: 1}
                if (216 != v2) {
                    v25 := GenS2{F0: v1, F1: 385, F2: -9223372036854775808}
                    fmt.Println(v2, v2)
                    for fi26 := 0; fi26 < 14; fi26++ {
                        fmt.Println(fi23, v1, v3)
                        v27 := []int{}
                        _ = fi26
                        _ = v27
                    }
                    _ = v25
                } else {
                    var v28 uint64 = 18446744073709551615
                    v2 += 65535
                    _ = v28
                }
                var v29 string = "\n\t"
                _ = v24
                _ = v29
            } else {
                if false {
                    v30 := []int{(fi23 / (((fi23 & (fi23 - (-(fi23))))) | 1)), (fi23 - 2147483647), (fi23 + fi23), -356}
                    _ = v30
                }
                v2 = (v2 ^ 255)
                for fi31 := 0; fi31 < 4; fi31++ {
                    v2 = (v2 >> (((v1 | v1)) & 63))
                    v2 = v2
                    _ = fi31
                }
            }
            v32 := []int{-1, (fi23 >> ((v1) & 63))}
            if true {
                fmt.Println(v2, v2)
                var v33 uint16 = 2
                switch ((((((((("\n\t" + "a") + "hello") + ("hello" + "a")) + "\n\t") + ("\n\t" + "")) + "") + "\n\t") + ("foo bar" + "")) + ("a" + "")) {
                case "hello":
                    for fi34 := 0; fi34 < 7; fi34++ {
                        v35 := GenS3{F0: 285}
                        _ = fi34
                        _ = v35
                    }
                case "foo bar":
                    fmt.Println(v3, v33)
                default:
                    _ = len(v32)
                }
                _ = v33
            }
            _ = fi23
            _ = v32
        }
    default:
        switch 127 {
        case 448:
            v3 = float32(1.5)
        case 9223372036854775807:
            v1 -= 307
        }
    }
    var v36 float32 = 0.1
    fmt.Println(v3, v2, v2)
    var v37 bool = (!false)
    fmt.Println(v3, v37)
    var v38 int64 = -469
    var v39 int8 = 5
    v40 := []int{-50, -2147483648, 2147483647, int(v39)}
    var v41 uint32 = uint32(v2)
    if false {
        for fi42 := 0; fi42 < 8; fi42++ {
            if false {
                var v43 uint8 = 255
                _ = len(v40)
                for fi44 := 0; fi44 < 14; fi44++ {
                    v40 = append(v40, 186)
                    v39 += v39
                    if v37 {
                        v45 := GenS2{F0: v1, F1: 4294967295, F2: v38}
                        var v46 uint16 = v2
                        v47 := GenS1{F0: v39, F1: true, F2: v39}
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                    if len(v40) > 0 {
                        v40[0] = fi44
                    }
                    _ = len(v40)
                    var v48 uint16 = (v2 | (^(v2)))
                    v49 := GenS1{F0: (^(v39)), F1: true, F2: 0}
                    _ = fi44
                    _ = v48
                    _ = v49
                }
                _ = v43
            } else {
                if (!false) {
                    var v50 string = "a"
                    var v51 uint64 = 168
                    var v52 uint64 = 27
                    _ = v50
                    _ = v51
                    _ = v52
                } else {
                    for fi53 := 0; fi53 < 13; fi53++ {
                        v54 := GenS4{F0: 0, F1: -80, F2: fi42}
                        _ = fi53
                        _ = v54
                    }
                }
                if (!v37) {
                    if ((v41 % (((v41 & (v41 << ((v1) & 63)))) | 1)) <= v41) {
                        v41 = (^(v41))
                        v55 := GenS5{F0: int16(v39)}
                        switch "\n\t" {
                        case "hello":
                            for fi56 := 0; fi56 < 10; fi56++ {
                                if v37 {
                                    if true {
                                        v57 := GenS5{F0: 42}
                                        var v58 uint = (v1 & (v1 >> ((v1) & 63)))
                                        _ = v57
                                        _ = v58
                                    }
                                    switch (fi56 & fi42) {
                                    case 32767:
                                        if v37 {
                                            v59 := GenS6{F0: (^(v41)), F1: 256}
                                            v36 = v3
                                            fmt.Println(v1)
                                            if true {
                                                v60 := GenS6{F0: v41, F1: 0}
                                                if v37 {
                                                    var v61 bool = ((v36 >= v3) || false)
                                                    var v62 uint64 = 42
                                                    _ = v61
                                                    _ = v62
                                                } else {
                                                    v1 = (^(v1))
                                                    var v63 float32 = v36
                                                    _ = v63
                                                }
                                                _ = v60
                                            } else {
                                                v64 := GenS7{F0: (v39 ^ (v39 << (((^(v1))) & 63)))}
                                                _ = v64
                                            }
                                            switch int(v41) {
                                            case 205:
                                                if ((v37 || (18446744073709551615 != 194)) && v37) {
                                                    fmt.Println(v1)
                                                    fmt.Println(v1, v41, v1)
                                                }
                                            case 843:
                                                switch "a" {
                                                case "":
                                                    switch "foo bar" {
                                                    case "hello":
                                                        if (!v37) {
                                                            if len(v40) > 0 {
                                                                v40[0] = fi42
                                                            }
                                                            v40 = append(v40, 9223372036854775807)
                                                            fmt.Println(fi42)
                                                        } else {
                                                            if v37 {
                                                                var v65 string = ("hello" + "a")
                                                                _ = v65
                                                            }
                                                            v66 := GenS4{F0: 0, F1: -185, F2: 4294967295}
                                                            _ = v66
                                                        }
                                                    case "a":
                                                        v67 := GenS7{F0: 42}
                                                        _ = v67
                                                    case "foo bar":
                                                        if len(v40) > 0 {
                                                            v40[0] = (fi42 & (fi42 ^ (^(fi42))))
                                                        }
                                                    }
                                                case "a":
                                                    v68 := GenS6{F0: v41, F1: -2147483648}
                                                    _ = v68
                                                }
                                            }
                                            _ = v59
                                        } else {
                                            var v69 uint64 = 417
                                            var v70 float32 = v36
                                            var v71 int32 = -2147483648
                                            _ = v69
                                            _ = v70
                                            _ = v71
                                        }
                                    case 374:
                                        if (v1 < (v1 ^ (^(v1)))) {
                                            switch (fi42 * (-(fi56))) {
                                            case 4294967295:
                                                if (v37 || true) {
                                                    if v37 {
                                                        switch "\n\t" {
                                                        case "a":
                                                            var v72 bool = (v1 == v1)
                                                            _ = v72
                                                        case "\n\t":
                                                            var v73 int64 = v38
                                                            _ = v73
                                                        }
                                                    }
                                                } else {
                                                    var v74 int32 = -594
                                                    v75 := GenS8{F0: 255}
                                                    _ = v74
                                                    _ = v75
                                                }
                                            case 9223372036854775807:
                                                var v76 uint = v1
                                                _ = v76
                                            case 1:
                                                v3 -= (0.0 * float32(0.1))
                                                v3 -= v3
                                            }
                                            if (false || false) {
                                                v38 -= (v38 + v38)
                                                v77 := GenS9{F0: 0}
                                                if (("a" + ("foo bar" + ("a" + ""))) < "") {
                                                    if v37 {
                                                        if v37 {
                                                            var v78 string = "a"
                                                            fmt.Println(fi42)
                                                            v41 = v41
                                                            _ = v78
                                                        } else {
                                                            switch "hello" {
                                                            case "hello":
                                                                var v79 int32 = 132
                                                                _ = v79
                                                            }
                                                        }
                                                        if (!(!v37)) {
                                                            if (-692 >= -2147483648) {
                                                                v80 := GenS2{F0: 1, F1: 42, F2: (-(v38))}
                                                                fmt.Println(v39, v36, v3)
                                                                _ = v80
                                                            } else {
                                                                v39 -= 127
                                                                var v81 int64 = v38
                                                                _ = v81
                                                            }
                                                            v39 -= (v39 / (((v39 % ((int8(fi42)) | 1))) | 1))
                                                        } else {
                                                            if len(v40) >= 2 {
                                                                v40 = v40[1:len(v40)]
                                                            }
                                                            var v82 uint16 = 1
                                                            _ = v82
                                                        }
                                                    } else {
                                                        v83 := GenS10{F0: 1, F1: 1, F2: 40}
                                                        switch fi42 {
                                                        case -1:
                                                            if ("foo bar" == ("a" + ("" + (("hello" + "") + ("hello" + "\n\t"))))) {
                                                                var v84 int64 = v38
                                                                var v85 float64 = -2.25
                                                                v84 = (v84 ^ -433)
                                                                v86 := GenS5{F0: 1}
                                                                _ = v84
                                                                _ = v85
                                                                _ = v86
                                                            }
                                                        case 127:
                                                            var v87 uint16 = (v2 ^ (v2 ^ v2))
                                                            _ = v87
                                                        case -9223372036854775808:
                                                            var v88 uint8 = 42
                                                            _ = v88
                                                        default:
                                                            fmt.Println(fi42, v37, v36)
                                                        }
                                                        fmt.Println(fi42, fi56)
                                                        fmt.Println(v1)
                                                        v36 = -2.25
                                                        switch fi56 {
                                                        case 574:
                                                            var v89 int8 = (v39 + 0)
                                                            _ = v89
                                                        case 32767:
                                                            var v90 uint64 = 982
                                                            _ = v90
                                                        case 2147483647:
                                                            fmt.Println(v39, fi56)
                                                        default:
                                                            v91 := GenS9{F0: 115}
                                                            _ = v91
                                                        }
                                                        _ = v83
                                                    }
                                                    var v92 uint64 = 1
                                                    _ = len(v40)
                                                    _ = v92
                                                }
                                                _ = v77
                                            }
                                            v1 -= (v1 & uint(v41))
                                            if len(v40) >= 2 {
                                                v40 = v40[1:len(v40)]
                                            }
                                        } else {
                                            var v93 uint16 = (^(v2))
                                            _ = v93
                                        }
                                    case -462:
                                        v37 = (("hello" + "foo bar") <= "")
                                    default:
                                        v94 := GenS11{F0: 224, F1: (^(v2)), F2: "\n\t"}
                                        var v95 string = "\n\t"
                                        var v96 uint8 = 53
                                        _ = v94
                                        _ = v95
                                        _ = v96
                                    }
                                } else {
                                    v1 = (v1 | v1)
                                    var v97 int64 = 42
                                    _ = v97
                                }
                                if (true || ((^(v39)) == v39)) {
                                    var v98 uint16 = 466
                                    v99 := GenS10{F0: (v38 * (v38 % ((-412) | 1))), F1: v1, F2: 104}
                                    _ = v98
                                    _ = v99
                                }
                                var v100 int16 = 0
                                var v101 uint8 = 2
                                fmt.Println(v37, v100)
                                v101 = 119
                                v102 := GenS1{F0: (v39 >> (((v1 / (((v1 & 42)) | 1))) & 63)), F1: v37, F2: int8(v101)}
                                _ = fi56
                                _ = v100
                                _ = v101
                                _ = v102
                            }
                        case "":
                            fmt.Println(fi42, v41, v36)
                        default:
                            v40 = append(v40, fi42)
                        }
                        _ = v55
                    } else {
                        var v103 uint = (^(v1))
                        fmt.Println(v36, v39)
                        _ = v103
                    }
                    for fi104 := 0; fi104 < 9; fi104++ {
                        v2 = uint16(v41)
                        fmt.Println(v39, v37)
                        fmt.Println(fi42, v36, v3)
                        _ = fi104
                    }
                    fmt.Println(v41, fi42, v37)
                } else {
                    var v105 string = ("foo bar" + "")
                    if false {
                        var v106 int32 = 1
                        fmt.Println(v37)
                        switch (-(fi42)) {
                        case -14:
                            v37 = true
                        }
                        _ = v106
                    } else {
                        v39 = (v39 ^ (v39 << ((v1) & 63)))
                        var v107 uint8 = 223
                        if (uint16(v1) == 818) {
                            if v37 {
                                if (false && ((v2 * (v2 % ((477) | 1))) <= (v2 ^ v2))) {
                                    switch fi42 {
                                    case 2147483647:
                                        for fi108 := 0; fi108 < 13; fi108++ {
                                            v107 = (v107 << (((v1 & (v1 - (^(v1))))) & 63))
                                            if (4294967295 > uint64(v41)) {
                                                var v109 int = fi42
                                                _ = v109
                                            }
                                            _ = fi108
                                        }
                                    case 32767:
                                        v110 := GenS9{F0: 1}
                                        _ = v110
                                    default:
                                        var v111 uint16 = 32767
                                        _ = v111
                                    }
                                    var v112 uint = 291
                                    v113 := GenS5{F0: -32768}
                                    _ = v112
                                    _ = v113
                                } else {
                                    for fi114 := 0; fi114 < 14; fi114++ {
                                        v115 := GenS12{F0: (!((true && false) || (446 > 896))), F1: (v107 * (v107 >> ((uint(v41)) & 63)))}
                                        v40 = append(v40, (fi42 & (fi42 ^ (fi114 * int(v1)))))
                                        v116 := GenS3{F0: -1}
                                        _ = fi114
                                        _ = v115
                                        _ = v116
                                    }
                                }
                                var v117 bool = v37
                                _ = v117
                            } else {
                                v38 -= 1
                                var v118 string = (v105 + ("" + (v105 + ((v105 + v105) + v105))))
                                _ = v118
                            }
                            fmt.Println(v38)
                            for fi119 := 0; fi119 < 20; fi119++ {
                                v39 -= -1
                                if (!(628 <= (^(v1)))) {
                                    v120 := GenS3{F0: int32(v2)}
                                    v121 := GenS13{F0: 0, F1: uint8(v2), F2: (fi119 % ((fi119) | 1))}
                                    _ = v120
                                    _ = v121
                                } else {
                                    if (v37 || v37) {
                                        v40 = append(v40, -259)
                                        v105 = ((v105 + (("\n\t" + ("a" + v105)) + ((v105 + v105) + v105))) + ((v105 + v105) + v105))
                                    }
                                }
                                _ = fi119
                            }
                        } else {
                            fmt.Println(v41, v39, v105)
                        }
                        _ = v107
                    }
                    fmt.Println(v1, v36)
                    _ = v105
                }
            }
            _ = fi42
        }
        if len(v40) >= 2 {
            v40 = v40[1:len(v40)]
        }
    } else {
        v122 := GenS7{F0: -33}
        _ = v122
    }
    v2 -= (^(v2))
    v123 := GenS14{F0: v38, F1: "foo bar", F2: v39}
    if v37 {
        v40 = append(v40, -736)
    } else {
        var v124 uint8 = 0
        if (!((^(v1)) < (v1 ^ 255))) {
            fmt.Println(v41, v38, v1)
            var v125 float64 = 0.0
            _ = v125
        } else {
            v1 -= (v1 - (v1 << (((v1 & 4294967295)) & 63)))
            if len(v40) >= 2 {
                v40 = v40[1:len(v40)]
            }
        }
        _ = v124
    }
    v3 -= -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v123
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v40
    _ = v41
}
