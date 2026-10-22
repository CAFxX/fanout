package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 bool }
type GenS2 struct { F0 int; F1 bool; F2 int32 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 int32; F1 uint32; F2 uint }
type GenS5 struct { F0 int64; F1 int8; F2 uint64 }
type GenS6 struct { F0 uint8; F1 int64 }
type GenS7 struct { F0 string; F1 int; F2 int16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint8; F1 int32; F2 int64 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int; F1 int }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 int8; F1 uint8 }
type GenS14 struct { F0 int32; F1 int8 }
type GenS15 struct { F0 uint; F1 int8; F2 uint }
type GenS16 struct { F0 uint32; F1 int; F2 int }
type GenS17 struct { F0 uint16; F1 uint64 }
func genFn0(p0 uint) int16 {
    if ((120 >= 2) && false) {
        if (false && (true || (false && false))) {
            switch (("" + ("a" + "\n\t")) + "a") {
            case "foo bar":
                if (!(363 < uint16(p0))) {
                    var v6 string = "a"
                    if (false && true) {
                        for fi7 := 0; fi7 < 7; fi7++ {
                            fmt.Println(fi7, v6)
                            _ = fi7
                        }
                        v6 = "\n\t"
                        v6 = v6
                    } else {
                        if true {
                            v6 = v6
                            p0 += p0
                            for fi8 := 0; fi8 < 15; fi8++ {
                                for fi9 := 0; fi9 < 4; fi9++ {
                                    var v10 uint64 = 42
                                    switch (fi8 | (fi8 << ((p0) & 63))) {
                                    case 4294967295:
                                        v11 := []int{}
                                        _ = v11
                                    case -891:
                                        var v12 float64 = -2.25
                                        _ = v12
                                    case -71:
                                        v13 := []int{(^(fi9)), fi8, (fi8 >> (((^(p0))) & 63))}
                                        _ = v13
                                    default:
                                        v6 = ("\n\t" + ((v6 + "a") + "foo bar"))
                                    }
                                    v14 := []int{(fi8 - (fi9 & fi8)), fi8, (fi9 ^ fi8), fi9}
                                    _ = fi9
                                    _ = v10
                                    _ = v14
                                }
                                _ = fi8
                            }
                        } else {
                            v6 = ("hello" + ((((("hello" + v6) + v6) + ("" + (v6 + v6))) + "") + v6))
                        }
                        fmt.Println(v6)
                        switch v6 {
                        case "hello":
                            if false {
                                fmt.Println(p0, v6)
                                v15 := []int{}
                                if true {
                                    v16 := GenS1{F0: 85, F1: (42 == 127), F2: true}
                                    v17 := GenS1{F0: 1, F1: (255 < uint16(p0)), F2: (false || true)}
                                    v6 = ("" + v6)
                                    _ = v16
                                    _ = v17
                                }
                                _ = v15
                            }
                        case "":
                            v18 := []int{127, 193, -9223372036854775808, 9223372036854775807}
                            _ = v18
                        default:
                            var v19 uint16 = 42
                            _ = v19
                        }
                    }
                    _ = v6
                } else {
                    var v20 float64 = ((1.5 * 0.1) - float64(1e10))
                    v20 -= (((v20 / 1e10) - (v20 / float64(1.5))) + v20)
                    _ = v20
                }
            case "hello":
                switch "hello" {
                case "\n\t":
                    p0 = 2
                case "":
                    if true {
                        fmt.Println(p0)
                        p0 = 42
                        v21 := []int{4294967295, 0, 127, -258}
                        _ = v21
                    } else {
                        switch "foo bar" {
                        case "":
                            for fi22 := 0; fi22 < 1; fi22++ {
                                v23 := GenS2{F0: (fi22 >> ((p0) & 63)), F1: (!true), F2: -402}
                                var v24 int16 = 361
                                v25 := GenS3{F0: -1, F1: 42}
                                _ = fi22
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                        case "foo bar":
                            var v26 uint8 = 42
                            _ = v26
                        default:
                            v27 := []int{}
                            _ = v27
                        }
                    }
                case "a":
                    var v28 bool = false
                    _ = v28
                default:
                    if (2 >= 255) {
                        var v29 int64 = -774
                        v30 := GenS4{F0: 179, F1: 488, F2: p0}
                        _ = v29
                        _ = v30
                    }
                }
            }
            v31 := []int{887}
            _ = v31
        } else {
            var v32 int = 4294967295
            var v33 int = -618
            fmt.Println(p0, v33)
            _ = v32
            _ = v33
        }
        v34 := GenS4{F0: -1, F1: 603, F2: p0}
        switch 987 {
        case 4294967295:
            v35 := []int{42, 4294967295}
            _ = v35
        }
        _ = v34
    }
    fmt.Println(p0)
    fmt.Println(p0)
    v36 := GenS2{F0: 888, F1: false, F2: 0}
    _ = p0
    _ = v36
    return -1
}
func genFn1() int {
    var v37 bool = false
    v37 = (((0 < 743) && (v37 || v37)) || v37)
    var v38 int16 = 1
    var v39 bool = true
    if v37 {
        v40 := GenS1{F0: 179, F1: (-666 < 752), F2: v37}
        fmt.Println(v39, v39)
        _ = v40
    }
    _ = v37
    _ = v38
    _ = v39
    return 2147483647
}

func main() {
    v1 := []string{("hello" + "\n\t"), ("" + "hello")}
    var v2 uint64 = 0
    var v3 int64 = 42
    var v4 int64 = (-(v3))
    v5 := make(map[string]bool)
    for fi41 := 0; fi41 < 4; fi41++ {
        v42 := v5[""]
        v43 := GenS4{F0: 2147483647, F1: 207, F2: 614}
        _ = fi41
        _ = v42
        _ = v43
    }
    for fi44 := 0; fi44 < 2; fi44++ {
        _ = len(v1)
        _ = fi44
    }
    for fi45 := 0; fi45 < 13; fi45++ {
        delete(v5, "\n\t")
        _ = fi45
    }
    if ((v3 < (v3 % ((int64(v2)) | 1))) || (!("hello" != "foo bar"))) {
        v46 := GenS5{F0: 0, F1: 0, F2: 235}
        fmt.Println(v2, v2, v2)
        _ = v46
    } else {
        v4 = (v3 >> ((2) & 63))
        for fi47 := 0; fi47 < 14; fi47++ {
            if false {
                v2 = 18446744073709551615
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v2 = 1
            } else {
                if false {
                    var v48 uint8 = 1
                    fmt.Println(v2, v48, v48)
                    _ = v48
                } else {
                    if false {
                        v4 = (v3 - v3)
                        if ("" <= "") {
                            switch "hello" {
                            case "a":
                                fmt.Println(fi47)
                            case "foo bar":
                                v4 += (v4 ^ v3)
                            default:
                                for fi49 := 0; fi49 < 4; fi49++ {
                                    fmt.Println(v4, fi47, fi49)
                                    _ = fi49
                                }
                            }
                            var v50 bool = false
                            for fi51 := 0; fi51 < 1; fi51++ {
                                switch ("foo bar" + "") {
                                case "\n\t":
                                    v50 = v50
                                case "hello":
                                    v1 = append(v1, "")
                                default:
                                    var v52 uint = 570
                                    _ = v52
                                }
                                v53 := GenS6{F0: 255, F1: 42}
                                fmt.Println(fi51, v50)
                                _ = fi51
                                _ = v53
                            }
                            _ = v50
                        } else {
                            if (!(0 != 42)) {
                                v5["a"] = (!true)
                                v4 = (v3 / ((v3) | 1))
                                switch int(v2) {
                                case 32767:
                                    if false {
                                        if true {
                                            for fi54 := 0; fi54 < 10; fi54++ {
                                                v55 := GenS7{F0: (("\n\t" + (("a" + "") + ("\n\t" + "\n\t"))) + (("\n\t" + "a") + "foo bar")), F1: genFn1(), F2: 0}
                                                _ = fi54
                                                _ = v55
                                            }
                                            fmt.Println(v3, fi47, v4)
                                            for fi56 := 0; fi56 < 14; fi56++ {
                                                v4 = v4
                                                v57 := GenS6{F0: 132, F1: (v3 * v4)}
                                                _ = fi56
                                                _ = v57
                                            }
                                        }
                                        var v58 uint32 = 0
                                        v59 := GenS8{F0: -1}
                                        _ = v58
                                        _ = v59
                                    }
                                case 9223372036854775807:
                                    v60 := GenS9{F0: 2, F1: 724, F2: (v3 ^ v4)}
                                    _ = v60
                                case -9223372036854775808:
                                    var v61 float32 = 1.5
                                    _ = v61
                                default:
                                    fmt.Println(fi47)
                                }
                            } else {
                                delete(v5, "hello")
                                switch (^(fi47)) {
                                case -180:
                                    if true {
                                        for fi62 := 0; fi62 < 8; fi62++ {
                                            var v63 uint32 = 0
                                            _ = fi62
                                            _ = v63
                                        }
                                        var v64 string = (("" + ("hello" + "hello")) + "foo bar")
                                        for fi65 := 0; fi65 < 2; fi65++ {
                                            if false {
                                                v66 := GenS7{F0: "hello", F1: (^(fi65)), F2: genFn0(4294967295)}
                                                fmt.Println(fi65, v64, v64)
                                                _ = v66
                                            }
                                            _ = fi65
                                        }
                                        _ = v64
                                    } else {
                                        v67 := GenS1{F0: 195, F1: false, F2: false}
                                        v68 := GenS10{F0: 97}
                                        _ = v67
                                        _ = v68
                                    }
                                default:
                                    v2 = v2
                                }
                                var v69 int16 = 949
                                _ = v69
                            }
                            for fi70 := 0; fi70 < 12; fi70++ {
                                var v71 int32 = 1
                                _ = fi70
                                _ = v71
                            }
                            v72 := GenS11{F0: fi47, F1: (-(fi47))}
                            _ = v72
                        }
                    } else {
                        switch fi47 {
                        case 127:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case 571:
                            fmt.Println(v2, v3, v2)
                        case 710:
                            var v73 float64 = float64(float32(1e10))
                            _ = v73
                        }
                    }
                    _ = len(v5)
                }
            }
            if ((false && false) && ((2147483647 == 320) || false)) {
                if false {
                    fmt.Println(v2, v3)
                } else {
                    v74 := GenS8{F0: 1}
                    _ = v74
                }
                v75 := GenS4{F0: 23, F1: 260, F2: 42}
                switch (fi47 - 127) {
                case 128:
                    v76 := v5["hello"]
                    _ = v76
                case 4294967295:
                    for fi77 := 0; fi77 < 13; fi77++ {
                        v78 := v5["a"]
                        if (true || v78) {
                            v3 -= 371
                            v79 := GenS12{F0: 65535}
                            v3 = (v3 << ((2) & 63))
                            var v80 int8 = -44
                            _ = v79
                            _ = v80
                        }
                        v5["a"] = v78
                        v78 = true
                        fmt.Println(v4, fi77, v4)
                        switch "foo bar" {
                        case "":
                            switch fi77 {
                            case 286:
                                var v81 bool = v78
                                _ = v81
                            case 2147483647:
                                v4 = 1
                            case 32767:
                                v3 = v3
                            default:
                                switch fi77 {
                                case 32768:
                                    v82 := GenS13{F0: -91, F1: 93}
                                    _ = v82
                                default:
                                    var v83 bool = v78
                                    _ = v83
                                }
                            }
                        case "\n\t":
                            var v84 uint16 = 255
                            _ = v84
                        case "hello":
                            if ((-207 < 1000) && (65535 != 837)) {
                                if len(v1) > 0 {
                                    v1[0] = "foo bar"
                                }
                                switch 9223372036854775807 {
                                case 1:
                                    v2 -= 1
                                case 0:
                                    v85 := GenS13{F0: 127, F1: 93}
                                    _ = v85
                                }
                                var v86 string = "a"
                                _ = v86
                            }
                        }
                        _ = fi77
                        _ = v78
                    }
                case -2147483648:
                    if true {
                        v4 -= (v4 & (^(v3)))
                    }
                default:
                    v87 := GenS9{F0: 176, F1: 42, F2: v3}
                    _ = v87
                }
                _ = v75
            }
            for fi88 := 0; fi88 < 4; fi88++ {
                if false {
                    if false {
                        var v89 uint16 = 917
                        if (!true) {
                            var v90 string = ("a" + ("\n\t" + "foo bar"))
                            v5[v90] = false
                            var v91 uint64 = v2
                            v92 := GenS14{F0: 0, F1: 119}
                            var v93 int8 = -51
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v90
                            _ = v91
                            _ = v92
                            _ = v93
                        }
                        _ = v89
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            v94 := GenS4{F0: 1, F1: 272, F2: 2}
                            var v95 uint32 = 207
                            switch "foo bar" {
                            case "foo bar":
                                switch genFn1() {
                                case -9223372036854775808:
                                    v96 := GenS2{F0: (fi88 % (((^(fi88))) | 1)), F1: false, F2: -2147483648}
                                    _ = v96
                                case 128:
                                    if (!false) {
                                        switch (fi88 % ((2147483647) | 1)) {
                                        case -1:
                                            v2 -= v2
                                        case 0:
                                            v3 += (v4 - (v3 | (v3 + v4)))
                                        case -489:
                                            var v97 int = (fi47 * (fi88 ^ (^(fi47))))
                                            _ = v97
                                        default:
                                            var v98 int = (fi88 & 826)
                                            _ = v98
                                        }
                                    } else {
                                        if len(v1) > 0 {
                                            v1[0] = "\n\t"
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = "a"
                                        }
                                    }
                                case 42:
                                    var v99 int = (fi47 + fi47)
                                    _ = v99
                                }
                            }
                            _ = v94
                            _ = v95
                        case "a":
                            if true {
                                if false {
                                    v100 := GenS15{F0: 1, F1: -1, F2: 1}
                                    v3 = v4
                                    fmt.Println(v2)
                                    _ = v100
                                } else {
                                    v4 -= v4
                                    v101 := GenS4{F0: 42, F1: 4294967295, F2: 4294967295}
                                    v102 := GenS16{F0: 1, F1: (fi88 + (fi88 % ((fi88) | 1))), F2: (^(fi47))}
                                    _ = v101
                                    _ = v102
                                }
                            }
                        case "\n\t":
                            if true {
                                if (-1 >= 0) {
                                    if len(v1) > 0 {
                                        v1[0] = "\n\t"
                                    }
                                } else {
                                    var v103 uint16 = 0
                                    _ = v103
                                }
                                v4 += (-(v4))
                            }
                        default:
                            if true {
                                if (v2 == v2) {
                                    var v104 string = ("foo bar" + ("\n\t" + (("" + ("foo bar" + "hello")) + ((("" + "") + "\n\t") + ("\n\t" + "foo bar")))))
                                    _ = v104
                                } else {
                                    if true {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        fmt.Println(v3)
                                    } else {
                                        var v105 bool = (v4 <= v3)
                                        v105 = (0 >= 2147483647)
                                        _ = v105
                                    }
                                    var v106 int8 = -128
                                    v107 := GenS15{F0: 761, F1: 127, F2: 2}
                                    _ = v106
                                    _ = v107
                                }
                                switch ((("hello" + ("a" + ("foo bar" + "foo bar"))) + "") + "") {
                                case "foo bar":
                                    var v108 uint = 4294967295
                                    _ = v108
                                case "a":
                                    v1 = append(v1, "")
                                case "":
                                    var v109 string = "hello"
                                    _ = v109
                                }
                                var v110 string = ""
                                _ = v110
                            }
                        }
                        v2 += v2
                        if (genFn0(42) >= -32768) {
                            switch (fi47 % (((fi47 + genFn1())) | 1)) {
                            case 673:
                                v3 = (v4 ^ (v3 % (((v3 - 0)) | 1)))
                            case 2147483647:
                                var v111 uint64 = (^(v2))
                                _ = v111
                            case -2147483648:
                                var v112 uint64 = 248
                                _ = v112
                            default:
                                v113 := GenS4{F0: -255, F1: 346, F2: 4294967295}
                                _ = v113
                            }
                        }
                    }
                    switch ((("\n\t" + ("hello" + "a")) + "\n\t") + "") {
                    case "hello":
                        if false {
                            v4 = (v4 | (v3 << ((18446744073709551615) & 63)))
                        } else {
                            v2 = v2
                            var v114 uint16 = 1
                            v2 -= v2
                            _ = v114
                        }
                        fmt.Println(v4, fi47)
                        _ = len(v5)
                    case "":
                        if (190 <= 1) {
                            var v115 uint64 = v2
                            var v116 float32 = 0.1
                            _ = v115
                            _ = v116
                        } else {
                            if false {
                                v3 = int64(v2)
                                if false {
                                    v117 := GenS12{F0: 836}
                                    fmt.Println(fi88)
                                    v2 = v2
                                    _ = v117
                                }
                                fmt.Println(v3, v4)
                            } else {
                                switch "hello" {
                                case "":
                                    v2 -= v2
                                case "a":
                                    v118 := GenS5{F0: (v3 - (-(v4))), F1: 79, F2: v2}
                                    _ = v118
                                case "\n\t":
                                    var v119 int = (fi88 >> ((0) & 63))
                                    _ = v119
                                default:
                                    var v120 float64 = -2.25
                                    _ = v120
                                }
                                if false {
                                    _ = len(v1)
                                    if true {
                                        if true {
                                            _ = len(v1)
                                            v3 = v3
                                            v121 := GenS8{F0: -544}
                                            _ = v121
                                        }
                                    }
                                }
                            }
                        }
                    case "foo bar":
                        v2 = 0
                    default:
                        var v122 bool = (!((787 <= 42) || true))
                        _ = v122
                    }
                    v4 = v4
                }
                v123 := GenS8{F0: 32767}
                _ = fi88
                _ = v123
            }
            _ = fi47
        }
    }
    v2 -= v2
    fmt.Println(v4, v2, v4)
    v3 = (v3 & -50)
    v124 := GenS6{F0: 169, F1: (v3 + v4)}
    var v125 uint = 65535
    if len(v1) > 0 {
        v1[0] = ("hello" + "\n\t")
    }
    if (599 == 64) {
        switch -2147483648 {
        case -754:
            v126 := GenS17{F0: 255, F1: (v2 & 18446744073709551615)}
            _ = v126
        default:
            v127 := v5[("a" + ("" + "foo bar"))]
            _ = v127
        }
        fmt.Println(v4, v4)
        for fi128 := 0; fi128 < 15; fi128++ {
            if ((0.0 - 1.5) > (0.1 * (1e10 - 0.0))) {
                _ = len(v5)
                for fi129 := 0; fi129 < 1; fi129++ {
                    var v130 int16 = 32767
                    _ = fi129
                    _ = v130
                }
            }
            fmt.Println(fi128, v3)
            _ = fi128
        }
    }
    for fi131 := 0; fi131 < 5; fi131++ {
        for fi132 := 0; fi132 < 20; fi132++ {
            if len(v1) > 0 {
                v1[0] = ("a" + "")
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi132
        }
        v125 = (v125 >> ((65535) & 63))
        _ = fi131
    }
    _ = len(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v125:", v125)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v124
    _ = v125
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
