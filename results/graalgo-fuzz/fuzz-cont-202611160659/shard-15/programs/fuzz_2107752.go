package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int16 }
type GenS2 struct { F0 int32; F1 uint8; F2 uint }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 uint8; F1 string; F2 uint64 }
type GenS7 struct { F0 int32; F1 uint }
type GenS8 struct { F0 uint32; F1 uint64 }
type GenS9 struct { F0 uint; F1 uint32 }
type GenS10 struct { F0 uint16; F1 uint; F2 string }
type GenS11 struct { F0 int32; F1 int32 }
type GenS12 struct { F0 int64; F1 int16 }
type GenS13 struct { F0 uint8; F1 uint64; F2 int }

func main() {
    var v1 int32 = 1
    v2 := make(map[int]uint32)
    var v3 string = "a"
    var v4 int = -827
    var v5 uint64 = 2
    v2[v4] = uint32(v5)
    if (!(1 != 4294967295)) {
        fmt.Println(v5)
        var v6 int64 = 1
        _ = v6
    }
    var v7 uint32 = 2147483647
    v3 = v3
    delete(v2, -9223372036854775808)
    if (true && ("" == v3)) {
        v8 := []int{(v4 | 2147483647), 133, (^(v4))}
        _ = v8
    } else {
        v7 -= (v7 % ((2147483648) | 1))
        if (-9223372036854775808 != 9223372036854775807) {
            var v9 uint8 = 255
            for fi10 := 0; fi10 < 8; fi10++ {
                v11 := v2[-431]
                v12 := GenS1{F0: ((0 <= 0) && false), F1: 345, F2: 1000}
                _ = fi10
                _ = v11
                _ = v12
            }
            var v13 uint64 = v5
            _ = v13
            _ = v9
        }
        var v14 uint64 = (v5 % ((1) | 1))
        _ = v14
    }
    var v15 int16 = 351
    v7 = (^(v7))
    var v16 uint = 0
    v17 := []int{}
    v18 := GenS1{F0: true, F1: 42, F2: v15}
    _ = len(v17)
    v3 = (((((v3 + v3) + v3) + ("a" + v3)) + v3) + ("a" + ("a" + (v3 + v3))))
    if true {
        v3 = v3
        if false {
            v19 := GenS2{F0: v1, F1: 65, F2: v16}
            var v20 uint16 = uint16(v7)
            _ = v19
            _ = v20
        } else {
            v21 := GenS1{F0: (!((0 < (^(v7))) || (!(((!false) || (!(!false))) || false)))), F1: 65535, F2: (-(v15))}
            v22 := GenS1{F0: (32767 <= 65535), F1: 910, F2: (v15 - v15)}
            v2[(v4 & v4)] = v7
            _ = v21
            _ = v22
        }
    } else {
        v23 := GenS1{F0: true, F1: 0, F2: (-(v15))}
        _ = len(v2)
        _ = v23
    }
    switch 127 {
    case -697:
        switch 364 {
        case 2147483647:
            v3 = "\n\t"
        case 42:
            for fi24 := 0; fi24 < 16; fi24++ {
                var v25 int8 = -1
                _ = fi24
                _ = v25
            }
        case 32768:
            v4 = v4
        default:
            if (("a" + "\n\t") != v3) {
                var v26 uint8 = 255
                _ = len(v2)
                _ = v26
            }
        }
    case 943:
        v7 = 42
    }
    switch (v3 + ("" + ("\n\t" + (("\n\t" + "a") + ((v3 + v3) + v3))))) {
    case "a":
        for fi27 := 0; fi27 < 3; fi27++ {
            switch "a" {
            case "hello":
                for fi28 := 0; fi28 < 19; fi28++ {
                    v1 = v1
                    switch (v4 >> (((v16 + (v16 % (((v16 << ((2) & 63))) | 1)))) & 63)) {
                    case 0:
                        if true {
                            switch v3 {
                            case "foo bar":
                                v16 = v16
                            case "a":
                                var v29 uint64 = (v5 ^ v5)
                                _ = v29
                            case "hello":
                                v30 := GenS1{F0: false, F1: 255, F2: (^(v15))}
                                _ = v30
                            default:
                                var v31 string = ""
                                _ = v31
                            }
                        }
                    case 981:
                        var v32 int = fi27
                        v16 -= v16
                        switch int(v7) {
                        case 32768:
                            _ = len(v2)
                        case 0:
                            var v33 float64 = (((3.14159 + (-2.25 + 0.0)) / (0.1 + float64(3.14159))) / 3.14159)
                            _ = v33
                        }
                        _ = v32
                    default:
                        var v34 uint32 = (v7 & (v7 - uint32(v16)))
                        _ = v34
                    }
                    delete(v2, (fi28 << ((302) & 63)))
                    v35 := GenS2{F0: int32(v5), F1: 255, F2: (v16 - 255)}
                    v36 := GenS3{F0: 0, F1: 127}
                    _ = fi28
                    _ = v35
                    _ = v36
                }
            case "a":
                v37 := v2[fi27]
                _ = v37
            default:
                for fi38 := 0; fi38 < 19; fi38++ {
                    var v39 string = v3
                    var v40 int32 = (v1 - v1)
                    _ = fi38
                    _ = v39
                    _ = v40
                }
            }
            v15 -= v15
            v4 += fi27
            _ = fi27
        }
    case "foo bar":
        delete(v2, -2147483648)
    case "hello":
        if false {
            var v41 string = "foo bar"
            switch v4 {
            case 42:
                switch v41 {
                case "a":
                    switch 1 {
                    case -1:
                        fmt.Println(v16, v4, v16)
                    }
                case "hello":
                    var v42 uint16 = 32767
                    _ = v42
                case "foo bar":
                    v43 := GenS4{F0: "a", F1: ((!(((v4 | v4) < (^(v4))) || false)) || (v41 < (v41 + ("foo bar" + v41))))}
                    _ = v43
                default:
                    v4 = -915
                }
            case -9223372036854775808:
                if false {
                    if false {
                        var v44 uint64 = (v5 & 580)
                        _ = v44
                    } else {
                        switch (v4 * v4) {
                        case 620:
                            v2[v4] = v7
                        case 127:
                            v45 := GenS1{F0: true, F1: 2, F2: 0}
                            _ = v45
                        default:
                            var v46 uint = (^(v16))
                            _ = v46
                        }
                        if false {
                            for fi47 := 0; fi47 < 7; fi47++ {
                                switch "foo bar" {
                                case "\n\t":
                                    var v48 uint32 = v7
                                    _ = v48
                                case "foo bar":
                                    v49 := GenS3{F0: 595, F1: -62}
                                    _ = v49
                                }
                                if (!false) {
                                    v50 := GenS3{F0: -9223372036854775808, F1: 104}
                                    if len(v17) >= 2 {
                                        v17 = v17[1:len(v17)]
                                    }
                                    v51 := GenS4{F0: (((v3 + (v3 + v3)) + (v41 + v3)) + (v3 + v3)), F1: ((v4 << (((^(v16))) & 63)) > -2147483648)}
                                    _ = v50
                                    _ = v51
                                } else {
                                    v52 := GenS3{F0: -9223372036854775808, F1: -18}
                                    _ = v52
                                }
                                _ = fi47
                            }
                            var v53 bool = (v5 == v5)
                            switch int(v16) {
                            case 32767:
                                v16 -= (^(v16))
                            default:
                                switch (v4 / (((v4 << ((139) & 63))) | 1)) {
                                case 32768:
                                    var v54 int32 = v1
                                    _ = v54
                                case -9223372036854775808:
                                    v7 = (^(v7))
                                case 426:
                                    v55 := GenS5{F0: -309, F1: (v5 % ((4294967295) | 1))}
                                    _ = v55
                                default:
                                    fmt.Println(v3)
                                }
                            }
                            _ = v53
                        } else {
                            fmt.Println(v41, v1)
                            var v56 int8 = 100
                            _ = v56
                        }
                        switch (^(v4)) {
                        case 127:
                            v7 -= v7
                        case 2147483647:
                            var v57 float32 = (1.5 / ((0.1 * 0.1) - 1.5))
                            _ = v57
                        default:
                            var v58 uint64 = v5
                            _ = v58
                        }
                    }
                    fmt.Println(v1)
                    for fi59 := 0; fi59 < 15; fi59++ {
                        for fi60 := 0; fi60 < 9; fi60++ {
                            var v61 string = (((v41 + v3) + (((v3 + (v41 + (v3 + (v3 + (v41 + "\n\t"))))) + v41) + v41)) + v41)
                            _ = fi60
                            _ = v61
                        }
                        if (463 >= 65535) {
                            v62 := GenS4{F0: "a", F1: (2 <= 127)}
                            for fi63 := 0; fi63 < 4; fi63++ {
                                _ = len(v2)
                                v64 := GenS6{F0: 2, F1: (v41 + v3), F2: 18446744073709551615}
                                var v65 int32 = int32(v15)
                                _ = fi63
                                _ = v64
                                _ = v65
                            }
                            _ = v62
                        } else {
                            _ = len(v2)
                            for fi66 := 0; fi66 < 8; fi66++ {
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                                var v67 uint32 = (v7 / (((v7 * 2147483647)) | 1))
                                v68 := GenS4{F0: "", F1: ((!true) && true)}
                                _ = fi66
                                _ = v67
                                _ = v68
                            }
                        }
                        v2[(^(v4))] = (v7 / ((4294967295) | 1))
                        _ = fi59
                    }
                }
            }
            _ = v41
        } else {
            for fi69 := 0; fi69 < 19; fi69++ {
                if false {
                    v4 = (v4 | 1)
                    v70 := GenS4{F0: "foo bar", F1: false}
                    _ = v70
                } else {
                    for fi71 := 0; fi71 < 12; fi71++ {
                        fmt.Println(v3, v16)
                        _ = fi71
                    }
                    var v72 uint32 = (v7 * v7)
                    switch (fi69 | 694) {
                    case -354:
                        for fi73 := 0; fi73 < 2; fi73++ {
                            if ((v15 == (v15 >> ((uint(v7)) & 63))) && false) {
                                var v74 int = fi69
                                v15 -= (v15 - (v15 & v15))
                                switch v4 {
                                case 137:
                                    var v75 uint16 = 65535
                                    _ = v75
                                case 676:
                                    v76 := GenS4{F0: v3, F1: (0 > v16)}
                                    _ = v76
                                }
                                _ = v74
                            } else {
                                v77 := GenS5{F0: 42, F1: v5}
                                v15 -= v15
                                _ = v77
                            }
                            if (v1 < (v1 >> ((v16) & 63))) {
                                var v78 int = -9223372036854775808
                                v3 = v3
                                v3 = (v3 + "\n\t")
                                _ = v78
                            }
                            var v79 uint8 = 1
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                            v5 = v5
                            _ = fi73
                            _ = v79
                        }
                    }
                    _ = v72
                }
                switch v3 {
                case "":
                    v7 = 2147483647
                case "\n\t":
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                case "foo bar":
                    if false {
                        var v80 uint = v16
                        for fi81 := 0; fi81 < 9; fi81++ {
                            var v82 uint64 = uint64(v7)
                            var v83 bool = false
                            var v84 uint16 = 291
                            _ = fi81
                            _ = v82
                            _ = v83
                            _ = v84
                        }
                        fmt.Println(v3)
                        _ = v80
                    } else {
                        delete(v2, 875)
                    }
                default:
                    fmt.Println(v7, v3)
                }
                for fi85 := 0; fi85 < 17; fi85++ {
                    fmt.Println(fi69, v15, fi85)
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                    if false {
                        if (uint8(v16) < 0) {
                            switch v3 {
                            case "":
                                var v86 bool = true
                                if true {
                                    var v87 string = "hello"
                                    if ((v7 / ((4294967295) | 1)) != 65535) {
                                        v5 -= (v5 << ((uint(v7)) & 63))
                                    } else {
                                        var v88 uint8 = 0
                                        var v89 uint64 = (^(v5))
                                        v90 := GenS1{F0: (v86 && (((v87 != (v3 + "hello")) || true) || ((^(v88)) <= (v88 - (^(v88)))))), F1: 255, F2: -232}
                                        _ = v88
                                        _ = v89
                                        _ = v90
                                    }
                                    _ = v87
                                } else {
                                    _ = len(v17)
                                    if v86 {
                                        v91 := GenS6{F0: 238, F1: v3, F2: v5}
                                        switch (v3 + v3) {
                                        case "":
                                            v92 := GenS7{F0: -1, F1: v16}
                                            _ = v92
                                        case "a":
                                            switch v3 {
                                            case "foo bar":
                                                var v93 uint64 = (^(v5))
                                                _ = v93
                                            case "":
                                                fmt.Println(v7, fi69, v86)
                                            }
                                        }
                                        switch "a" {
                                        case "foo bar":
                                            v94 := v2[fi69]
                                            _ = v94
                                        case "\n\t":
                                            var v95 uint8 = 2
                                            _ = v95
                                        }
                                        v5 = v5
                                        _ = v91
                                    } else {
                                        var v96 int64 = 42
                                        _ = len(v17)
                                        _ = v96
                                    }
                                }
                                _ = v86
                            case "\n\t":
                                v97 := GenS8{F0: 65535, F1: (v5 + (v5 - 2))}
                                _ = v97
                            }
                            var v98 uint64 = v5
                            v99 := GenS4{F0: v3, F1: true}
                            _ = v98
                            _ = v99
                        }
                    } else {
                        fmt.Println(fi85, v1, v3)
                        switch (v3 + v3) {
                        case "hello":
                            if (9223372036854775807 > 9223372036854775807) {
                                switch "hello" {
                                case "a":
                                    if false {
                                        if false {
                                            fmt.Println(v3, v7, v15)
                                            fmt.Println(v5, fi85)
                                        } else {
                                            switch int(v15) {
                                            case 1:
                                                v100 := GenS8{F0: v7, F1: (v5 << ((65535) & 63))}
                                                _ = v100
                                            case 32768:
                                                v5 += 2
                                            default:
                                                var v101 int64 = 0
                                                _ = v101
                                            }
                                            v1 += v1
                                            fmt.Println(v5, v7)
                                        }
                                        var v102 int64 = -1
                                        switch int(v7) {
                                        case 32767:
                                            v103 := GenS8{F0: (v7 * (^(v7))), F1: 0}
                                            _ = v103
                                        default:
                                            if true {
                                                v104 := GenS6{F0: 2, F1: "hello", F2: (v5 | v5)}
                                                _ = v104
                                            }
                                        }
                                        _ = v102
                                    } else {
                                        var v105 uint32 = v7
                                        _ = len(v17)
                                        _ = v105
                                    }
                                case "foo bar":
                                    switch v4 {
                                    case 9223372036854775807:
                                        if len(v17) > 0 {
                                            v17[0] = 32768
                                        }
                                    case 32767:
                                        v106 := GenS2{F0: v1, F1: 1, F2: 42}
                                        _ = v106
                                    default:
                                        v3 = (v3 + (("" + (v3 + ((v3 + (v3 + "a")) + ((v3 + ("hello" + v3)) + (v3 + v3))))) + (v3 + v3)))
                                    }
                                    v107 := v2[v4]
                                    _ = v107
                                case "":
                                    v4 += fi69
                                }
                                v108 := GenS9{F0: 645, F1: 995}
                                switch v3 {
                                case "":
                                    switch ("" + "hello") {
                                    case "\n\t":
                                        v4 += (fi85 + fi85)
                                    case "":
                                        v109 := GenS10{F0: 32767, F1: 18446744073709551615, F2: v3}
                                        _ = v109
                                    default:
                                        switch (v3 + v3) {
                                        case "a":
                                            fmt.Println(v7, v4)
                                        default:
                                            switch v4 {
                                            case -214:
                                                fmt.Println(v3, v16)
                                            case -772:
                                                if (v1 < int32(v16)) {
                                                    var v110 uint32 = (v7 - 458)
                                                    _ = v110
                                                } else {
                                                    var v111 int = -266
                                                    if len(v17) > 0 {
                                                        v17[0] = (v4 % ((32767) | 1))
                                                    }
                                                    v16 = v16
                                                    v112 := GenS11{F0: 0, F1: -103}
                                                    _ = len(v2)
                                                    _ = v111
                                                    _ = v112
                                                }
                                                var v113 uint32 = (v7 >> ((v16) & 63))
                                                _ = v113
                                            case 128:
                                                if (-24 < int8(v5)) {
                                                    switch -9223372036854775808 {
                                                    case 2147483647:
                                                        delete(v2, (fi85 - (fi85 & (fi85 + fi85))))
                                                    case 964:
                                                        switch 32767 {
                                                        case 4294967295:
                                                            var v114 int8 = -18
                                                            _ = v114
                                                        case -9223372036854775808:
                                                            if false {
                                                                v4 += 4294967295
                                                            }
                                                        }
                                                        var v115 uint32 = 1
                                                        var v116 float32 = 1.5
                                                        _ = v115
                                                        _ = v116
                                                    default:
                                                        fmt.Println(v5, v4, fi85)
                                                    }
                                                    v4 += (fi85 % (((fi69 % ((fi69) | 1))) | 1))
                                                    v117 := v2[int(v5)]
                                                    _ = v117
                                                } else {
                                                    if false {
                                                        switch ("hello" + ("\n\t" + v3)) {
                                                        case "a":
                                                            v118 := GenS12{F0: 0, F1: (v15 | v15)}
                                                            _ = v118
                                                        default:
                                                            switch (-(v4)) {
                                                            case 2147483647:
                                                                switch v4 {
                                                                case -797:
                                                                    switch v3 {
                                                                    case "a":
                                                                        var v119 float32 = 1.5
                                                                        _ = v119
                                                                    case "\n\t":
                                                                        switch int(v15) {
                                                                        case 32767:
                                                                            fmt.Println(fi69, v1, v1)
                                                                        }
                                                                        _ = len(v17)
                                                                        switch v3 {
                                                                        case "hello":
                                                                            delete(v2, v4)
                                                                        }
                                                                    case "foo bar":
                                                                        v120 := GenS13{F0: 127, F1: (v5 & v5), F2: (fi85 ^ (v4 * (fi69 >> (((^(v16))) & 63))))}
                                                                        _ = v120
                                                                    default:
                                                                        var v121 uint32 = 894
                                                                        _ = v121
                                                                    }
                                                                    var v122 int = 2147483647
                                                                    var v123 uint8 = 255
                                                                    _ = v122
                                                                    _ = v123
                                                                case 4294967295:
                                                                    v124 := GenS9{F0: v16, F1: (v7 - (v7 / ((v7) | 1)))}
                                                                    _ = v124
                                                                }
                                                            case -478:
                                                                if (!true) {
                                                                    v125 := GenS10{F0: 2, F1: v16, F2: "a"}
                                                                    v126 := GenS2{F0: (-(v1)), F1: 77, F2: (v16 ^ v16)}
                                                                    v4 = (fi69 / (((v4 * fi69)) | 1))
                                                                    _ = v125
                                                                    _ = v126
                                                                }
                                                            case 820:
                                                                v7 = 619
                                                            }
                                                        }
                                                        fmt.Println(fi85, fi69, v5)
                                                        v15 += v15
                                                    }
                                                }
                                            }
                                        }
                                    }
                                case "foo bar":
                                    fmt.Println(fi69)
                                case "hello":
                                    if len(v17) >= 2 {
                                        v17 = v17[1:len(v17)]
                                    }
                                default:
                                    var v127 int64 = 1
                                    _ = v127
                                }
                                _ = v108
                            } else {
                                var v128 int8 = 100
                                _ = v128
                            }
                        }
                        if true {
                            if false {
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                                if (v7 == v7) {
                                    v5 = (v5 / ((uint64(v16)) | 1))
                                    var v129 string = v3
                                    _ = v129
                                }
                            } else {
                                v16 = (v16 * 255)
                                if true {
                                    var v130 int16 = (v15 - (v15 + (-(v15))))
                                    var v131 uint8 = 1
                                    _ = v130
                                    _ = v131
                                }
                                fmt.Println(v7, fi69)
                            }
                        }
                    }
                    var v132 int16 = v15
                    _ = fi85
                    _ = v132
                }
                _ = fi69
            }
        }
    }
    var v133 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v133:", v133)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v133
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
