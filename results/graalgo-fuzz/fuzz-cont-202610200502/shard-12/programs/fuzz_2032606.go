package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 int; F1 uint; F2 uint }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int64; F1 int32; F2 int32 }
type GenS8 struct { F0 bool; F1 uint16; F2 int64 }
type GenS9 struct { F0 int8; F1 int8 }
type GenS10 struct { F0 int64; F1 string; F2 string }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS14 struct { F0 int8; F1 uint; F2 uint }
type GenS15 struct { F0 bool; F1 uint8 }
type GenS16 struct { F0 int8; F1 uint }

func main() {
    var v1 int32 = 285
    var v2 string = ("\n\t" + "hello")
    var v3 uint8 = 2
    v4 := make(map[uint32]string)
    var v5 uint16 = 162
    v6 := []uint32{0, 576, 1, 2147483647}
    var v7 uint32 = 814
    if false {
        v4[v7] = ""
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 uint64 = 683
            var v10 float64 = 0.0
            v7 -= v7
            _ = fi8
            _ = v10
            _ = v9
        }
    } else {
        var v11 int8 = 127
        _ = v11
    }
    if (!true) {
        var v12 int32 = 905
        v12 += v12
        if true {
            for fi13 := 0; fi13 < 20; fi13++ {
                if false {
                    for fi14 := 0; fi14 < 18; fi14++ {
                        _ = len(v4)
                        _ = fi14
                    }
                }
                v2 = ((("a" + v2) + (v2 + ("\n\t" + ((v2 + ((v2 + (v2 + v2)) + v2)) + v2)))) + (v2 + "hello"))
                v1 -= v1
                _ = fi13
            }
        }
        _ = v12
    }
    v3 = v3
    if (9223372036854775807 == 0) {
        v1 = (v1 - -1)
    } else {
        if len(v6) > 0 {
            v6[0] = v7
        }
        v15 := GenS1{F0: v2}
        if false {
            for fi16 := 0; fi16 < 1; fi16++ {
                for fi17 := 0; fi17 < 6; fi17++ {
                    fmt.Println(v5, v7, fi17)
                    var v18 int32 = int32(v7)
                    v19 := GenS2{F0: 1, F1: 0, F2: 4294967295}
                    var v20 int64 = -9223372036854775808
                    var v21 uint = 815
                    v3 = 255
                    _ = fi17
                    _ = v18
                    _ = v19
                    _ = v20
                    _ = v21
                }
                _ = fi16
            }
        }
        _ = v15
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        var v23 uint32 = 424
        v24 := GenS2{F0: v3, F1: 255, F2: 1}
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        v1 = int32(fi25)
        if ((((1.5 - -2.25) / 0.1) * 0.1) != 3.14159) {
            var v26 int8 = 127
            var v27 uint = 8
            _ = v26
            _ = v27
        } else {
            if len(v6) > 0 {
                v6[0] = v7
            }
        }
        for fi28 := 0; fi28 < 19; fi28++ {
            if len(v6) > 0 {
                v6[0] = v7
            }
            if true {
                switch v2 {
                case "\n\t":
                    fmt.Println(fi25, v7)
                case "":
                    v29 := GenS1{F0: "hello"}
                    _ = v29
                default:
                    v30 := GenS3{F0: 42}
                    _ = v30
                }
                v31 := GenS3{F0: v1}
                _ = v31
            } else {
                if true {
                    switch (("\n\t" + "hello") + ((v2 + (("a" + (v2 + v2)) + (v2 + "\n\t"))) + "hello")) {
                    case "foo bar":
                        switch ((v2 + v2) + v2) {
                        case "hello":
                            if (!(!false)) {
                                switch fi25 {
                                case -287:
                                    var v32 uint16 = 2
                                    _ = v32
                                case 9223372036854775807:
                                    if (fi28 >= fi28) {
                                        v33 := GenS1{F0: ((("hello" + "foo bar") + v2) + (v2 + (v2 + ("a" + v2))))}
                                        var v34 uint16 = 255
                                        var v35 float64 = 0.1
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    } else {
                                        v36 := GenS4{F0: 0, F1: (fi28 ^ fi28)}
                                        delete(v4, 42)
                                        _ = v36
                                    }
                                }
                            }
                        case "":
                            if false {
                                var v37 float32 = 1.5
                                v1 += -2147483648
                                var v38 bool = true
                                _ = v37
                                _ = v38
                            }
                            var v39 int16 = 544
                            fmt.Println(v5)
                            fmt.Println(v3, fi28, v5)
                            v40 := GenS4{F0: 704, F1: 446}
                            _ = v39
                            _ = v40
                        default:
                            var v41 int8 = 1
                            _ = v41
                        }
                    default:
                        switch v2 {
                        case "hello":
                            fmt.Println(v1, v3)
                        case "a":
                            v42 := GenS5{F0: -117, F1: 4294967295, F2: 292}
                            _ = v42
                        default:
                            v43 := GenS1{F0: ""}
                            _ = v43
                        }
                    }
                    switch v2 {
                    case "hello":
                        v1 = v1
                    case "\n\t":
                        if (false && true) {
                            v5 = (v5 & v5)
                            var v44 int8 = 1
                            if true {
                                switch fi28 {
                                case -2147483648:
                                    fmt.Println(v2, fi25)
                                case -1:
                                    if false {
                                        v1 += (^(v1))
                                    } else {
                                        v45 := GenS6{F0: v2}
                                        v3 = (v3 + 209)
                                        _ = v45
                                    }
                                }
                                var v46 int8 = v44
                                if (!true) {
                                    if ((v44 % (((v44 % (((-(v44))) | 1))) | 1)) > v46) {
                                        if (v46 != (v44 * 94)) {
                                            fmt.Println(fi25, v7)
                                            v4[(v7 % ((v7) | 1))] = v2
                                            var v47 bool = ((((v2 + (v2 + (v2 + (v2 + (v2 + "a"))))) + v2) + v2) != ("hello" + "foo bar"))
                                            _ = v47
                                        } else {
                                            v48 := GenS1{F0: v2}
                                            _ = v48
                                        }
                                    } else {
                                        delete(v4, v7)
                                    }
                                    var v49 uint = 65535
                                    _ = v49
                                } else {
                                    v50 := GenS4{F0: 175, F1: (fi28 ^ (fi25 % ((fi28) | 1)))}
                                    _ = v50
                                }
                                _ = v46
                            }
                            _ = v44
                        } else {
                            if true {
                                if false {
                                    if (true && true) {
                                        fmt.Println(v3, v2)
                                        v51 := GenS7{F0: 1, F1: (v1 * 0), F2: 0}
                                        var v52 uint16 = (^(v5))
                                        v7 = v7
                                        _ = v51
                                        _ = v52
                                    }
                                    v2 = ""
                                    v53 := GenS2{F0: v3, F1: 1, F2: 0}
                                    _ = v53
                                } else {
                                    var v54 bool = false
                                    if (!(18446744073709551615 >= 18446744073709551615)) {
                                        var v55 uint = 0
                                        if false {
                                            switch v2 {
                                            case "":
                                                fmt.Println(v55, v1, fi25)
                                            case "a":
                                                if (!(0.1 >= float64(1e10))) {
                                                    v56 := GenS8{F0: true, F1: 422, F2: 1}
                                                    if v54 {
                                                        v4[v7] = (v2 + v2)
                                                        fmt.Println(v7, v2)
                                                    }
                                                    var v57 uint64 = 203
                                                    _ = v56
                                                    _ = v57
                                                } else {
                                                    v58 := GenS9{F0: 127, F1: 116}
                                                    v54 = (281 >= 2)
                                                    fmt.Println(v1)
                                                    _ = v58
                                                }
                                                if v54 {
                                                    v59 := GenS2{F0: 1, F1: v55, F2: 980}
                                                    v60 := GenS1{F0: "\n\t"}
                                                    v55 -= 18446744073709551615
                                                    _ = v59
                                                    _ = v60
                                                }
                                                v61 := GenS8{F0: ((!v54) || (32767 < 0)), F1: (^(v5)), F2: 9223372036854775807}
                                                _ = v61
                                            case "foo bar":
                                                var v62 bool = ("a" != v2)
                                                _ = v62
                                            default:
                                                if (1 > (^(v5))) {
                                                    switch (fi25 % (((fi28 / (((fi28 - (fi28 | fi28))) | 1))) | 1)) {
                                                    case 4294967295:
                                                        if (v54 || v54) {
                                                            fmt.Println(v2, fi25)
                                                            var v63 int = (fi25 / ((4294967295) | 1))
                                                            _ = v63
                                                        } else {
                                                            if (v1 > v1) {
                                                                v5 = 32767
                                                                var v64 uint32 = v7
                                                                if v54 {
                                                                    var v65 bool = (!(128 > (v3 & (v3 | (v3 * v3)))))
                                                                    fmt.Println(v5)
                                                                    v66 := GenS10{F0: 0, F1: "a", F2: v2}
                                                                    v55 = (v55 >> (((v55 / ((v55) | 1))) & 63))
                                                                    switch -9223372036854775808 {
                                                                    case -347:
                                                                        var v67 bool = (v54 && (!(v54 || (42 != 42))))
                                                                        _ = v67
                                                                    case 128:
                                                                        switch 42 {
                                                                        case 4294967295:
                                                                            if (fi25 > (^(fi28))) {
                                                                                v3 = v3
                                                                                v4[(v64 & v7)] = ""
                                                                            } else {
                                                                                v7 -= v7
                                                                            }
                                                                        default:
                                                                            if (2 != 821) {
                                                                                v6 = append(v6, (v64 % (((v7 | v7)) | 1)))
                                                                                v68 := GenS2{F0: v3, F1: (v55 + (v55 - v55)), F2: uint64(v7)}
                                                                                _ = v68
                                                                            }
                                                                        }
                                                                    case 32767:
                                                                        if (675 <= -9223372036854775808) {
                                                                            var v69 int8 = 65
                                                                            _ = v69
                                                                        } else {
                                                                            var v70 uint32 = (v64 + v64)
                                                                            v55 += (^(v55))
                                                                            _ = v70
                                                                        }
                                                                    default:
                                                                        fmt.Println(v65, fi28, v64)
                                                                    }
                                                                    if (-32768 == 399) {
                                                                        _ = len(v4)
                                                                        if (v54 && (42 == 0)) {
                                                                            switch (v2 + ("hello" + v2)) {
                                                                            case "a":
                                                                                var v71 int16 = -468
                                                                                _ = v71
                                                                            case "hello":
                                                                                fmt.Println(v7, v64)
                                                                            default:
                                                                                var v72 uint8 = (v3 >> ((v55) & 63))
                                                                                _ = v72
                                                                            }
                                                                            v55 += 1
                                                                        }
                                                                    } else {
                                                                        v54 = (!(65535 != v55))
                                                                    }
                                                                    _ = v65
                                                                    _ = v66
                                                                } else {
                                                                    fmt.Println(v55, v5, v5)
                                                                    var v73 uint16 = v5
                                                                    _ = v73
                                                                }
                                                                _ = v64
                                                            } else {
                                                                switch fi28 {
                                                                case 816:
                                                                    if (-2147483648 != 1) {
                                                                        var v74 int16 = 284
                                                                        _ = v74
                                                                    }
                                                                case 127:
                                                                    var v75 bool = (0 <= 0)
                                                                    _ = v75
                                                                default:
                                                                    _ = len(v4)
                                                                }
                                                                var v76 uint32 = (v7 | (v7 * 838))
                                                                _ = v76
                                                            }
                                                            v77 := GenS9{F0: -128, F1: 96}
                                                            v3 -= (v3 - v3)
                                                            _ = v77
                                                        }
                                                        _ = len(v6)
                                                    default:
                                                        var v78 uint64 = 18446744073709551615
                                                        _ = v78
                                                    }
                                                }
                                            }
                                            v54 = (v5 != v5)
                                            switch v2 {
                                            case "hello":
                                                if v54 {
                                                    if (v2 < ("a" + ((v2 + v2) + v2))) {
                                                        if v54 {
                                                            switch (v2 + (v2 + v2)) {
                                                            case "":
                                                                var v79 int32 = v1
                                                                _ = v79
                                                            case "\n\t":
                                                                var v80 uint8 = (^(v3))
                                                                if (!(v7 != v7)) {
                                                                    switch v2 {
                                                                    case "a":
                                                                        _ = len(v4)
                                                                    }
                                                                    v81 := GenS3{F0: (^(v1))}
                                                                    var v82 uint8 = v80
                                                                    _ = v81
                                                                    _ = v82
                                                                }
                                                                v83 := v4[2]
                                                                _ = v80
                                                                _ = v83
                                                            }
                                                            switch ((v2 + v2) + v2) {
                                                            case "a":
                                                                if ((fi25 << (((^(v55))) & 63)) > (fi25 / (((fi25 ^ fi25)) | 1))) {
                                                                    var v84 uint16 = v5
                                                                    fmt.Println(v2, fi28, v1)
                                                                    _ = v84
                                                                }
                                                            case "hello":
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    v54 = (v1 >= (v1 ^ (v1 >> ((v55) & 63))))
                                                                }
                                                                if v54 {
                                                                    v85 := GenS11{F0: (fi28 + (fi28 & (fi25 >> ((v55) & 63))))}
                                                                    fmt.Println(v3, v2)
                                                                    v86 := v4[v7]
                                                                    _ = v85
                                                                    _ = v86
                                                                }
                                                                v1 += int32(fi25)
                                                            case "\n\t":
                                                                delete(v4, v7)
                                                            }
                                                        } else {
                                                            fmt.Println(v2, v3, v5)
                                                            switch v2 {
                                                            case "\n\t":
                                                                v87 := GenS12{F0: 0}
                                                                _ = v87
                                                            }
                                                        }
                                                        var v88 bool = (-128 != 0)
                                                        var v89 string = ("foo bar" + v2)
                                                        _ = v88
                                                        _ = v89
                                                    }
                                                    switch v2 {
                                                    case "a":
                                                        v55 = (v55 * 65535)
                                                    }
                                                    v7 = (v7 / ((231) | 1))
                                                } else {
                                                    switch fi28 {
                                                    case 0:
                                                        v55 -= v55
                                                    case -686:
                                                        _ = len(v6)
                                                        if true {
                                                            var v90 uint32 = v7
                                                            _ = v90
                                                        }
                                                    case 32767:
                                                        var v91 uint8 = v3
                                                        _ = v91
                                                    default:
                                                        v3 -= (v3 >> ((42) & 63))
                                                    }
                                                    var v92 float64 = 1e10
                                                    _ = len(v4)
                                                    v93 := GenS6{F0: (("\n\t" + v2) + v2)}
                                                    var v94 int64 = int64(v5)
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        v2 = ((v2 + "a") + v2)
                                                    case "":
                                                        v3 -= (^(v3))
                                                    default:
                                                        switch "foo bar" {
                                                        case "a":
                                                            var v95 int64 = -9223372036854775808
                                                            _ = v95
                                                        }
                                                    }
                                                    var v96 int = (fi25 | -452)
                                                    var v97 uint16 = v5
                                                    var v98 int = (v96 - -1)
                                                    _ = v92
                                                    _ = v93
                                                    _ = v94
                                                    _ = v96
                                                    _ = v97
                                                    _ = v98
                                                }
                                            case "\n\t":
                                                fmt.Println(v2)
                                            case "foo bar":
                                                if v54 {
                                                    v99 := GenS4{F0: 2, F1: fi25}
                                                    v100 := GenS13{F0: v5, F1: (v55 - 42), F2: (v3 - (v3 % ((v3) | 1)))}
                                                    var v101 uint32 = 1
                                                    _ = v100
                                                    _ = v101
                                                    _ = v99
                                                }
                                            }
                                        } else {
                                            if true {
                                                switch v2 {
                                                case "":
                                                    if v54 {
                                                        v102 := GenS14{F0: -1, F1: (v55 % ((v55) | 1)), F2: v55}
                                                        v5 += (v5 & 696)
                                                        _ = v102
                                                    }
                                                case "a":
                                                    v55 = v55
                                                    v103 := GenS5{F0: fi28, F1: (v55 + uint(v5)), F2: (^(v55))}
                                                    var v104 uint = uint(v3)
                                                    _ = v103
                                                    _ = v104
                                                default:
                                                    var v105 bool = (-1 > int32(v3))
                                                    _ = v105
                                                }
                                            }
                                            var v106 uint = (^(v55))
                                            fmt.Println(fi25)
                                            _ = v106
                                        }
                                        var v107 int32 = v1
                                        _ = v107
                                        _ = v55
                                    }
                                    var v108 int = (fi28 % ((fi28) | 1))
                                    _ = v108
                                    _ = v54
                                }
                                _ = len(v6)
                                v109 := GenS15{F0: false, F1: 0}
                                switch v2 {
                                case "hello":
                                    fmt.Println(v2, v1, v5)
                                case "":
                                    var v110 int64 = -9223372036854775808
                                    _ = v110
                                case "\n\t":
                                    var v111 uint32 = v7
                                    _ = v111
                                }
                                if false {
                                    v7 = (^(v7))
                                    v6 = append(v6, v7)
                                }
                                _ = v109
                            } else {
                                v112 := GenS8{F0: false, F1: v5, F2: -425}
                                v113 := v4[(v7 - v7)]
                                _ = v112
                                _ = v113
                            }
                        }
                    }
                    _ = len(v4)
                } else {
                    var v114 int = fi25
                    _ = v114
                }
            }
            var v115 bool = false
            _ = fi28
            _ = v115
        }
        _ = fi25
    }
    var v116 string = v2
    for fi117 := 0; fi117 < 1; fi117++ {
        v118 := GenS16{F0: 127, F1: 2}
        v2 = v2
        _ = fi117
        _ = v118
    }
    if false {
        if false {
            fmt.Println(v5, v2)
            var v119 uint = 1
            v120 := v4[v7]
            _ = v119
            _ = v120
        }
    } else {
        for fi121 := 0; fi121 < 20; fi121++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = len(v6)
            v116 = ("" + v116)
            _ = fi121
        }
        v2 = v116
    }
    if true {
        var v122 uint64 = 18446744073709551615
        var v123 uint = 42
        _ = v122
        _ = v123
    } else {
        switch -1 {
        case -1:
            var v124 int16 = -461
            _ = v124
        default:
            v2 = "\n\t"
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v116:", v116)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v116
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
