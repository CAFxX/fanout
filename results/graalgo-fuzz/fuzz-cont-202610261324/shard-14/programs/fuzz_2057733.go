package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint64 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 uint8; F1 uint }
type GenS7 struct { F0 uint16; F1 uint }
type GenS8 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS9 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS10 struct { F0 bool; F1 int }

func main() {
    var v1 int16 = 1000
    v2 := []uint32{42, 54, 42}
    v3 := make(map[string]bool)
    v1 = (v1 & v1)
    v4 := GenS1{F0: 32768, F1: (0 == -1), F2: 982}
    if true {
        var v5 uint8 = 1
        _ = v5
    } else {
        v1 += 1
        v6 := GenS2{F0: "", F1: -1}
        switch -9223372036854775808 {
        case -2147483648:
            var v7 uint16 = 423
            _ = v7
        case -9223372036854775808:
            fmt.Println(v1)
        case 32768:
            var v8 uint8 = 127
            _ = v8
        }
        _ = v6
    }
    var v9 int16 = v1
    if true {
        var v10 uint8 = 154
        switch "" {
        case "\n\t":
            v11 := GenS2{F0: (("foo bar" + "\n\t") + ("a" + "\n\t")), F1: -1}
            _ = v11
        case "a":
            var v12 uint = 42
            _ = v12
        default:
            var v13 string = (("\n\t" + ("\n\t" + "foo bar")) + ("\n\t" + "a"))
            _ = v13
        }
        var v14 int = 9223372036854775807
        _ = v10
        _ = v14
    } else {
        switch ("hello" + ((("a" + "a") + "a") + ("\n\t" + ""))) {
        case "hello":
            var v15 float32 = float32(-2.25)
            _ = v15
        case "":
            var v16 uint16 = 415
            _ = v16
        case "foo bar":
            switch -9223372036854775808 {
            case 32767:
                v9 += v1
            case 0:
                for fi17 := 0; fi17 < 17; fi17++ {
                    for fi18 := 0; fi18 < 19; fi18++ {
                        v19 := GenS1{F0: (fi17 - 220), F1: true, F2: 0}
                        _ = fi18
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 17; fi20++ {
                        var v21 uint8 = 1
                        if true {
                            var v22 uint32 = 2147483648
                            _ = v22
                        }
                        v23 := GenS3{F0: 947}
                        _ = fi20
                        _ = v21
                        _ = v23
                    }
                    _ = fi17
                }
            default:
                v24 := GenS4{F0: 42}
                _ = v24
            }
        default:
            for fi25 := 0; fi25 < 4; fi25++ {
                var v26 int = (fi25 << ((42) & 63))
                v1 += v9
                if false {
                    switch v26 {
                    case 2147483647:
                        var v27 string = ("" + "\n\t")
                        _ = v27
                    case 4294967295:
                        for fi28 := 0; fi28 < 11; fi28++ {
                            if (-129 == v1) {
                                fmt.Println(v1, v1, fi28)
                                v9 = v1
                                delete(v3, ("\n\t" + ""))
                            } else {
                                var v29 bool = true
                                var v30 float64 = (-2.25 + (((3.14159 - (float64(0.0) - ((3.14159 * 1.5) + (((3.14159 - -2.25) + 0.0) / 1e10)))) * 3.14159) + (0.1 + -2.25)))
                                _ = v29
                                _ = v30
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi28
                        }
                    default:
                        v31 := GenS5{F0: 4294967295, F1: 32767}
                        _ = v31
                    }
                } else {
                    v26 += v26
                    for fi32 := 0; fi32 < 6; fi32++ {
                        v9 = (v9 & v9)
                        var v33 float32 = (3.14159 + 3.14159)
                        _ = fi32
                        _ = v33
                    }
                    v9 = v1
                }
                _ = fi25
                _ = v26
            }
        }
        var v34 int32 = 497
        _ = v34
    }
    for fi35 := 0; fi35 < 3; fi35++ {
        if (42 == 58) {
            var v36 int64 = 856
            for fi37 := 0; fi37 < 8; fi37++ {
                v9 -= 119
                if true {
                    v9 = (^(v1))
                    var v38 uint64 = 636
                    _ = v38
                } else {
                    var v39 int8 = -128
                    if (-968 == 1) {
                        fmt.Println(v9, v39, fi35)
                        if ("\n\t" < "foo bar") {
                            switch -282 {
                            case 585:
                                if true {
                                    fmt.Println(v39)
                                    fmt.Println(v39, fi35)
                                    fmt.Println(fi35)
                                } else {
                                    v39 += (^(v39))
                                }
                            case 679:
                                var v40 int = -1
                                _ = v40
                            }
                        } else {
                            if (!(true || (226 < 18446744073709551615))) {
                                if (false || false) {
                                    if false {
                                        v36 -= 575
                                    } else {
                                        v39 -= -1
                                        var v41 uint32 = 4294967295
                                        if true {
                                            v42 := GenS4{F0: 42}
                                            _ = v42
                                        }
                                        if (v39 != 42) {
                                            fmt.Println(v36)
                                            if true {
                                                var v43 float32 = 3.14159
                                                var v44 bool = (true && false)
                                                _ = v43
                                                _ = v44
                                            } else {
                                                var v45 string = ((("\n\t" + "") + "hello") + "")
                                                v45 = "foo bar"
                                                _ = v45
                                            }
                                            v46 := GenS2{F0: "\n\t", F1: (v36 - v36)}
                                            _ = v46
                                        }
                                        _ = v41
                                    }
                                    delete(v3, "\n\t")
                                    var v47 uint8 = 1
                                    _ = v47
                                } else {
                                    v48 := GenS1{F0: (fi37 & (fi35 * int(v39))), F1: (42 >= 0), F2: 1}
                                    fmt.Println(fi37, fi35, fi35)
                                    _ = v48
                                }
                            } else {
                                var v49 uint32 = 776
                                _ = v49
                            }
                            var v50 int8 = v39
                            v1 = v9
                            _ = v50
                        }
                        var v51 int64 = 1
                        _ = v51
                    }
                    _ = len(v2)
                    var v52 int64 = v36
                    v9 += (v1 + (v1 - 32767))
                    _ = v39
                    _ = v52
                }
                v36 = (v36 % (((^(v36))) | 1))
                _ = fi37
            }
            v1 -= 0
            _ = v36
        }
        var v53 float64 = float64((1e10 / 1.5))
        for fi54 := 0; fi54 < 8; fi54++ {
            if true {
                fmt.Println(fi35, fi54, fi35)
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if false {
                    fmt.Println(v9)
                    v9 += 32767
                    v9 -= v1
                } else {
                    v55 := GenS1{F0: -2147483648, F1: (!(42 == 252)), F2: 670}
                    switch fi54 {
                    case 4294967295:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    default:
                        v9 = v1
                        switch "foo bar" {
                        case "hello":
                            var v56 int32 = -1
                            _ = v56
                        }
                        switch fi54 {
                        case 9223372036854775807:
                            var v57 int = 775
                            _ = v57
                        case -9223372036854775808:
                            if ((!true) || (0 == 855)) {
                                if (-2.25 >= v53) {
                                    var v58 bool = true
                                    _ = v58
                                }
                                v59 := GenS1{F0: -9223372036854775808, F1: (v53 != ((0.0 * ((-2.25 / 1.0) - 1e10)) / v53)), F2: 906}
                                _ = v59
                            } else {
                                fmt.Println(v53)
                            }
                        default:
                            switch fi54 {
                            case 2147483647:
                                var v60 float64 = (v53 / v53)
                                _ = v60
                            case 380:
                                var v61 bool = true
                                _ = v61
                            }
                            v62 := v3[("\n\t" + "\n\t")]
                            fmt.Println(v62, v62)
                            fmt.Println(v62, v62, v9)
                            _ = v62
                        }
                    }
                    _ = v55
                }
                v1 -= (v9 << ((65535) & 63))
            }
            v9 += (v1 | 586)
            v9 -= v9
            _ = fi54
        }
        _ = fi35
        _ = v53
    }
    for fi63 := 0; fi63 < 15; fi63++ {
        for fi64 := 0; fi64 < 5; fi64++ {
            v65 := GenS4{F0: 2}
            _ = fi64
            _ = v65
        }
        switch (("\n\t" + ("foo bar" + ("hello" + ("foo bar" + "a")))) + "foo bar") {
        case "hello":
            for fi66 := 0; fi66 < 3; fi66++ {
                if true {
                    v2 = append(v2, 2147483647)
                    var v67 uint64 = 540
                    v1 -= (v9 | (v9 + v9))
                    _ = v67
                }
                var v68 uint = 782
                var v69 uint64 = 4294967295
                v69 = (v69 ^ v69)
                v9 = (v9 / (((v1 & 1)) | 1))
                _ = fi66
                _ = v68
                _ = v69
            }
        case "a":
            var v70 string = ("foo bar" + "foo bar")
            _ = v70
        case "\n\t":
            if false {
                if true {
                    var v71 uint32 = 0
                    fmt.Println(v1)
                    _ = v71
                } else {
                    switch fi63 {
                    case -1:
                        v72 := GenS2{F0: "\n\t", F1: 885}
                        _ = v72
                    case 42:
                        if false {
                            v2 = append(v2, 938)
                        }
                    case 32767:
                        v1 += (v9 ^ (v1 * -32768))
                    }
                    var v73 uint = 18446744073709551615
                    _ = v73
                }
                delete(v3, ("\n\t" + "a"))
                v9 -= (v1 & v1)
            } else {
                if len(v2) > 0 {
                    v2[0] = 1
                }
                v74 := GenS6{F0: 254, F1: 492}
                var v75 int = fi63
                _ = v74
                _ = v75
            }
        default:
            if false {
                v1 = (^(v1))
                var v76 uint16 = 104
                switch int(v76) {
                case -537:
                    for fi77 := 0; fi77 < 7; fi77++ {
                        v78 := GenS2{F0: ("foo bar" + "hello"), F1: -1}
                        var v79 uint = 541
                        var v80 uint64 = 0
                        _ = fi77
                        _ = v78
                        _ = v79
                        _ = v80
                    }
                case -52:
                    v3[("a" + (("\n\t" + "") + "\n\t"))] = ((!false) || false)
                case 9223372036854775807:
                    var v81 int32 = -1
                    _ = v81
                }
                _ = v76
            }
        }
        v82 := GenS7{F0: 1, F1: 0}
        _ = fi63
        _ = v82
    }
    var v83 uint16 = 0
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        if len(v2) > 0 {
            v2[0] = 557
        }
        if true {
            for fi84 := 0; fi84 < 5; fi84++ {
                v9 += (-(v1))
                _ = fi84
            }
            if (695 <= 0) {
                v85 := GenS2{F0: "hello", F1: 1}
                if false {
                    for fi86 := 0; fi86 < 19; fi86++ {
                        if true {
                            v87 := v3["hello"]
                            var v88 bool = (fi86 != (fi86 % ((fi86) | 1)))
                            _ = v87
                            _ = v88
                        } else {
                            switch ("a" + ((("foo bar" + ("foo bar" + ("" + ("a" + ("a" + "a"))))) + "\n\t") + ("foo bar" + "a"))) {
                            case "foo bar":
                                v89 := GenS6{F0: 62, F1: 255}
                                _ = v89
                            case "hello":
                                v90 := GenS8{F0: 4294967295, F1: 42, F2: (v9 ^ v1)}
                                _ = v90
                            }
                        }
                        v1 = (v9 >> ((42) & 63))
                        switch int(v9) {
                        case 32768:
                            if true {
                                var v91 int64 = -1
                                var v92 int = fi86
                                v93 := GenS2{F0: "\n\t", F1: 0}
                                _ = v91
                                _ = v92
                                _ = v93
                            } else {
                                var v94 uint16 = v83
                                _ = v94
                            }
                        case 1:
                            if len(v2) > 0 {
                                v2[0] = 2147483647
                            }
                        }
                        _ = fi86
                    }
                    if true {
                        fmt.Println(v9)
                        v9 -= (v1 / ((-1) | 1))
                    } else {
                        if false {
                            var v95 uint8 = 2
                            switch "foo bar" {
                            case "":
                                delete(v3, "a")
                            case "\n\t":
                                v95 = (^(v95))
                            case "foo bar":
                                v96 := GenS5{F0: 2, F1: (v83 ^ v83)}
                                _ = v96
                            }
                            v9 -= 407
                            _ = v95
                        }
                    }
                    if true {
                        var v97 uint = 18446744073709551615
                        fmt.Println(v9)
                        v9 = (v9 ^ v9)
                        _ = v97
                    }
                } else {
                    if true {
                        v9 -= -1
                        if (true && false) {
                            for fi98 := 0; fi98 < 10; fi98++ {
                                var v99 int32 = 1
                                v9 += (v9 >> ((2) & 63))
                                for fi100 := 0; fi100 < 5; fi100++ {
                                    v101 := GenS2{F0: "hello", F1: 88}
                                    _ = fi100
                                    _ = v101
                                }
                                _ = fi98
                                _ = v99
                            }
                            v102 := GenS9{F0: 42, F1: 75, F2: v83}
                            v9 += int16(v83)
                            _ = v102
                        } else {
                            v9 = v1
                        }
                        fmt.Println(v83)
                    } else {
                        var v103 bool = (-9223372036854775808 >= 42)
                        if v103 {
                            v104 := GenS9{F0: v83, F1: -128, F2: v83}
                            _ = v104
                        }
                        _ = v103
                    }
                    delete(v3, "hello")
                    fmt.Println(v83, v1)
                }
                v105 := GenS7{F0: v83, F1: 42}
                _ = v105
                _ = v85
            }
            for fi106 := 0; fi106 < 3; fi106++ {
                var v107 int64 = -1
                switch ("\n\t" + "") {
                case "a":
                    switch ("hello" + ((("\n\t" + "") + ("a" + "a")) + "hello")) {
                    case "\n\t":
                        v108 := GenS8{F0: 42, F1: 1, F2: v1}
                        _ = v108
                    case "foo bar":
                        var v109 uint32 = 65535
                        _ = v109
                    case "":
                        v110 := GenS7{F0: v83, F1: 0}
                        _ = v110
                    default:
                        v111 := v3["\n\t"]
                        _ = v111
                    }
                case "hello":
                    v112 := GenS2{F0: "\n\t", F1: (v107 / (((v107 >> ((uint(v83)) & 63))) | 1))}
                    _ = v112
                case "":
                    if (!true) {
                        if (42 >= 717) {
                            for fi113 := 0; fi113 < 4; fi113++ {
                                switch 42 {
                                case -1:
                                    v9 = (v9 + int16(v107))
                                case 2147483647:
                                    var v114 uint = 4294967295
                                    v107 += -9223372036854775808
                                    _ = v114
                                default:
                                    _ = len(v3)
                                }
                                v9 = v1
                                _ = fi113
                            }
                        }
                    } else {
                        v2 = append(v2, 265)
                    }
                }
                if true {
                    v9 += (v9 << ((42) & 63))
                    var v115 int64 = int64(fi106)
                    for fi116 := 0; fi116 < 20; fi116++ {
                        var v117 uint8 = 15
                        v83 -= (v83 / (((v83 / (((^(v83))) | 1))) | 1))
                        _ = fi116
                        _ = v117
                    }
                    _ = v115
                } else {
                    v107 = v107
                    fmt.Println(v83, v107)
                    v83 = 65535
                }
                _ = fi106
                _ = v107
            }
        } else {
            v118 := v3["hello"]
            _ = v118
        }
        fmt.Println(v9, v9, v9)
    } else {
        switch 42 {
        case 32768:
            switch "hello" {
            case "a":
                switch -9223372036854775808 {
                case 9223372036854775807:
                    var v119 int8 = 100
                    _ = v119
                case 128:
                    var v120 uint8 = 158
                    _ = v120
                default:
                    v2 = append(v2, 2147483647)
                }
            case "foo bar":
                fmt.Println(v9, v83)
            case "":
                fmt.Println(v9, v83, v1)
            default:
                for fi121 := 0; fi121 < 4; fi121++ {
                    v9 += v9
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi121
                }
            }
        case 127:
            v122 := GenS10{F0: (2 < 0), F1: 1}
            _ = v122
        case 128:
            v123 := GenS3{F0: 255}
            _ = v123
        default:
            fmt.Println(v83, v1)
        }
    }
    v2 = append(v2, 42)
    v124 := GenS2{F0: ("a" + (("hello" + ("" + ("a" + ("hello" + ("a" + ((("" + ("\n\t" + "hello")) + (("foo bar" + "hello") + "hello")) + (("" + "foo bar") + "hello"))))))) + "\n\t")), F1: 202}
    fmt.Println(v83, v9, v1)
    if (1 < 559) {
        for fi125 := 0; fi125 < 17; fi125++ {
            v1 = (v1 * (v1 - (-(v9))))
            if ((v83 + (v83 + (v83 << ((1) & 63)))) == 1) {
                var v126 uint = 770
                var v127 int = (fi125 / ((2147483647) | 1))
                var v128 bool = false
                _ = v126
                _ = v127
                _ = v128
            }
            v1 = (v9 % ((v1) | 1))
            _ = fi125
        }
        for fi129 := 0; fi129 < 2; fi129++ {
            switch "hello" {
            case "a":
                for fi130 := 0; fi130 < 19; fi130++ {
                    v131 := GenS3{F0: 1}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi130
                    _ = v131
                }
            case "":
                var v132 int = (^(fi129))
                _ = v132
            }
            if false {
                v133 := GenS9{F0: v83, F1: -84, F2: 1}
                _ = v133
            }
            v9 -= (v9 / ((v1) | 1))
            _ = fi129
        }
    }
    if false {
        v9 = 32767
        if false {
            var v134 uint32 = 2
            var v135 float64 = (1e10 * 0.1)
            fmt.Println(v83)
            _ = v134
            _ = v135
        }
        v136 := GenS6{F0: 42, F1: 2}
        _ = v136
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v83:", v83)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v124
    _ = v2
    _ = v3
    _ = v4
    _ = v83
    _ = v9
}
