package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 bool; F2 uint }
type GenS3 struct { F0 uint; F1 uint8; F2 int16 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 bool; F1 int16 }
type GenS7 struct { F0 int; F1 int64; F2 uint64 }
type GenS8 struct { F0 uint; F1 uint }
type GenS9 struct { F0 bool; F1 bool }
type GenS10 struct { F0 bool; F1 int16 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint8; F1 int32; F2 int8 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint8; F1 string; F2 int8 }
type GenS15 struct { F0 uint16; F1 bool }
func genFn0() int {
    if true {
        if false {
            for fi4 := 0; fi4 < 18; fi4++ {
                for fi5 := 0; fi5 < 16; fi5++ {
                    if true {
                        v6 := []int{(-(fi5)), (^(fi5)), (fi5 % ((2147483647) | 1)), fi4}
                        if true {
                            _ = len(v6)
                        } else {
                            var v7 uint64 = 2
                            fmt.Println(fi5, v7, fi5)
                            _ = v7
                        }
                        var v8 uint16 = 65535
                        var v9 uint16 = 255
                        _ = v6
                        _ = v8
                        _ = v9
                    } else {
                        var v10 int32 = 0
                        _ = v10
                    }
                    var v11 bool = false
                    v11 = (!(v11 || v11))
                    _ = fi5
                    _ = v11
                }
                var v12 uint16 = 2
                var v13 uint = 0
                _ = fi4
                _ = v12
                _ = v13
            }
            v14 := []int{4294967295}
            _ = v14
        }
    }
    var v15 int = 127
    for fi16 := 0; fi16 < 4; fi16++ {
        if false {
            v15 = (v15 ^ (-(fi16)))
        } else {
            var v17 uint32 = 42
            _ = v17
        }
        _ = fi16
    }
    _ = v15
    return (-(v15))
}
func genFn1(p0 int16) int8 {
    if true {
        var v18 uint = 18446744073709551615
        _ = v18
    }
    var v19 uint = 862
    switch "\n\t" {
    case "":
        if true {
            p0 = (p0 & 782)
        }
    case "hello":
        v19 = v19
    default:
        for fi20 := 0; fi20 < 12; fi20++ {
            if (!("foo bar" == "hello")) {
                if (!false) {
                    if false {
                        for fi21 := 0; fi21 < 18; fi21++ {
                            v22 := []int{(fi20 * 127), 251}
                            v23 := GenS1{F0: 42}
                            v24 := GenS2{F0: -769, F1: false, F2: 42}
                            if true {
                                v22 = append(v22, (-(fi20)))
                                fmt.Println(v19)
                            }
                            _ = fi21
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                    }
                } else {
                    var v25 uint = v19
                    if false {
                        var v26 float64 = float64(1.5)
                        fmt.Println(fi20, fi20, p0)
                        if (1 == -675) {
                            fmt.Println(fi20, p0, v26)
                        }
                        _ = v26
                    } else {
                        for fi27 := 0; fi27 < 15; fi27++ {
                            v28 := []int{}
                            if true {
                                p0 -= -32768
                            }
                            switch ("\n\t" + "\n\t") {
                            case "hello":
                                if false {
                                    v29 := GenS3{F0: 424, F1: 42, F2: (p0 % ((p0) | 1))}
                                    _ = v29
                                }
                            case "":
                                _ = len(v28)
                            }
                            _ = fi27
                            _ = v28
                        }
                        var v30 int8 = 103
                        _ = v30
                    }
                    _ = v25
                }
            } else {
                var v31 uint64 = 0
                _ = v31
            }
            _ = fi20
        }
    }
    _ = p0
    _ = v19
    return -108
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 102
    v3 := make(map[string]int)
    v32 := GenS3{F0: 548, F1: 255, F2: 32767}
    fmt.Println(v2)
    v2 -= v2
    v33 := []int{127, genFn0()}
    if (!((0 < 2) || false)) {
        fmt.Println(v2)
        var v34 uint8 = 127
        _ = v34
    }
    v2 = 0
    fmt.Println(v2)
    switch 0 {
    case 127:
        if ("\n\t" <= "hello") {
            if true {
                var v35 float32 = 0.0
                var v36 bool = true
                v36 = (((v35 * v35) / v35) > v35)
                _ = v35
                _ = v36
            }
        } else {
            for fi37 := 0; fi37 < 13; fi37++ {
                if (genFn1(42) != (-(v2))) {
                    if false {
                        if false {
                            fmt.Println(v2)
                            if false {
                                v38 := GenS3{F0: 0, F1: 127, F2: -1}
                                for fi39 := 0; fi39 < 5; fi39++ {
                                    var v40 string = "hello"
                                    fmt.Println(fi39, v2)
                                    v41 := GenS1{F0: 0}
                                    switch v40 {
                                    case "\n\t":
                                        v42 := GenS3{F0: 317, F1: 42, F2: 0}
                                        _ = v42
                                    }
                                    var v43 uint8 = 255
                                    _ = fi39
                                    _ = v40
                                    _ = v41
                                    _ = v43
                                }
                                _ = v38
                            } else {
                                v44 := v1["\n\t"]
                                _ = v44
                            }
                            for fi45 := 0; fi45 < 19; fi45++ {
                                fmt.Println(v2, v2)
                                v2 = (v2 % ((37) | 1))
                                _ = fi45
                            }
                        } else {
                            for fi46 := 0; fi46 < 15; fi46++ {
                                fmt.Println(fi37)
                                _ = fi46
                            }
                            switch genFn0() {
                            case 127:
                                v2 += genFn1(641)
                            case 518:
                                v33 = append(v33, (fi37 * genFn0()))
                            default:
                                v2 = genFn1(42)
                            }
                        }
                    } else {
                        v47 := v1["foo bar"]
                        _ = v47
                    }
                    v48 := GenS4{F0: (("\n\t" + ("a" + "foo bar")) + ("" + "a")), F1: 801}
                    for fi49 := 0; fi49 < 2; fi49++ {
                        if (1 > fi49) {
                            if true {
                                if (1000 < 42) {
                                    switch (fi37 << ((640) & 63)) {
                                    case -1:
                                        fmt.Println(v2, fi37, fi37)
                                    case 0:
                                        var v50 float32 = 3.14159
                                        _ = v50
                                    case 32767:
                                        v2 += 100
                                    default:
                                        v3["foo bar"] = (fi37 ^ int(v2))
                                    }
                                } else {
                                    if (42 > 0) {
                                        var v51 uint16 = 32767
                                        _ = v51
                                    }
                                    v2 = genFn1(-21)
                                    var v52 uint32 = 2147483647
                                    if (!true) {
                                        v1["foo bar"] = -9223372036854775808
                                        delete(v1, "foo bar")
                                    }
                                    _ = v52
                                }
                                var v53 float32 = float32((1.5 + (3.14159 - 1.5)))
                                switch "a" {
                                case "a":
                                    v1["\n\t"] = 2147483647
                                case "":
                                    v2 -= (v2 * 42)
                                default:
                                    var v54 int = (fi37 - 32768)
                                    _ = v54
                                }
                                switch "" {
                                case "":
                                    if (((3.14159 / 0.1) * float64(v53)) != -2.25) {
                                        switch "" {
                                        case "hello":
                                            if true {
                                                if false {
                                                    v55 := v1["hello"]
                                                    switch "" {
                                                    case "":
                                                        if ((true && false) && (!(true || ((v53 > v53) || true)))) {
                                                            v33 = append(v33, (-(fi49)))
                                                        }
                                                    case "hello":
                                                        v55 += fi49
                                                    case "\n\t":
                                                        if false {
                                                            var v56 uint16 = 1
                                                            switch "" {
                                                            case "foo bar":
                                                                fmt.Println(v55, v55)
                                                            }
                                                            _ = v56
                                                        }
                                                    }
                                                    _ = v55
                                                }
                                                v2 = (^(v2))
                                            } else {
                                                switch "a" {
                                                case "hello":
                                                    switch 128 {
                                                    case 32767:
                                                        if ((982 == 0) && false) {
                                                            v57 := GenS5{F0: 42, F1: ("" + ("a" + "hello"))}
                                                            _ = v57
                                                        }
                                                    case -10:
                                                        v58 := GenS6{F0: (false && true), F1: 364}
                                                        _ = v58
                                                    }
                                                case "":
                                                    var v59 bool = (!(32767 == 693))
                                                    _ = v59
                                                }
                                            }
                                        case "a":
                                            v60 := GenS2{F0: 0, F1: true, F2: 0}
                                            _ = v60
                                        case "\n\t":
                                            if (!false) {
                                                if (633 > 0) {
                                                    var v61 uint32 = 65535
                                                    v33 = append(v33, 0)
                                                    if (255 != 32767) {
                                                        v2 = v2
                                                        v53 = (((((-2.25 / 1e10) + v53) / 1e10) - 3.14159) * 0.0)
                                                        if (!false) {
                                                            v62 := GenS7{F0: fi49, F1: int64(v2), F2: 2}
                                                            switch (fi49 - fi37) {
                                                            case 128:
                                                                v61 = 585
                                                            case 127:
                                                                v63 := GenS8{F0: 557, F1: 21}
                                                                _ = v63
                                                            }
                                                            if true {
                                                                v64 := GenS9{F0: (!true), F1: false}
                                                                v65 := GenS10{F0: false, F1: 348}
                                                                if len(v33) >= 2 {
                                                                    v33 = v33[1:len(v33)]
                                                                }
                                                                _ = v64
                                                                _ = v65
                                                            }
                                                            _ = v62
                                                        }
                                                    } else {
                                                        if true {
                                                            v33 = append(v33, 9223372036854775807)
                                                            var v66 uint = 65535
                                                            switch "" {
                                                            case "hello":
                                                                v3[""] = fi49
                                                            }
                                                            _ = v66
                                                        } else {
                                                            _ = len(v3)
                                                            switch (((("foo bar" + "foo bar") + "foo bar") + "hello") + "\n\t") {
                                                            case "hello":
                                                                v67 := GenS8{F0: 65535, F1: 345}
                                                                _ = v67
                                                            default:
                                                                switch 127 {
                                                                case -412:
                                                                    v68 := GenS11{F0: (v61 & v61)}
                                                                    _ = v68
                                                                case -620:
                                                                    v53 = v53
                                                                case 9223372036854775807:
                                                                    switch (fi49 ^ (fi49 ^ genFn0())) {
                                                                    case 830:
                                                                        var v69 uint = 148
                                                                        _ = v69
                                                                    case -862:
                                                                        v53 = (v53 / v53)
                                                                        v70 := GenS5{F0: 5, F1: "foo bar"}
                                                                        if len(v33) >= 2 {
                                                                            v33 = v33[1:len(v33)]
                                                                        }
                                                                        _ = v70
                                                                    case 993:
                                                                        switch 32767 {
                                                                        case 127:
                                                                            v71 := GenS12{F0: 154, F1: 42, F2: -1}
                                                                            if true {
                                                                                switch fi37 {
                                                                                case -9223372036854775808:
                                                                                    var v72 int8 = v2
                                                                                    _ = v72
                                                                                case 32767:
                                                                                    var v73 uint64 = 456
                                                                                    _ = v73
                                                                                case 964:
                                                                                    v74 := GenS2{F0: 1, F1: true, F2: 1}
                                                                                    _ = v74
                                                                                }
                                                                                v2 += -1
                                                                            } else {
                                                                                v53 += (v53 * 3.14159)
                                                                                v53 -= 0.0
                                                                                var v75 int8 = (v2 / ((v2) | 1))
                                                                                switch (^(fi49)) {
                                                                                case 4294967295:
                                                                                    fmt.Println(fi37, v75, fi37)
                                                                                default:
                                                                                    switch "a" {
                                                                                    case "hello":
                                                                                        switch "" {
                                                                                        case "foo bar":
                                                                                            fmt.Println(v53)
                                                                                        default:
                                                                                            var v76 uint64 = 2
                                                                                            _ = v76
                                                                                        }
                                                                                    case "foo bar":
                                                                                        v53 -= -2.25
                                                                                        var v77 int32 = 2147483647
                                                                                        _ = v77
                                                                                    case "":
                                                                                        var v78 string = "\n\t"
                                                                                        _ = v78
                                                                                    default:
                                                                                        v1[""] = (fi49 / ((fi49) | 1))
                                                                                    }
                                                                                }
                                                                                _ = v75
                                                                            }
                                                                            fmt.Println(v2, fi37, fi49)
                                                                            _ = v71
                                                                        case -167:
                                                                            var v79 uint8 = 2
                                                                            _ = v79
                                                                        case 32767:
                                                                            var v80 uint8 = 2
                                                                            _ = v80
                                                                        }
                                                                    default:
                                                                        v61 += (v61 & v61)
                                                                    }
                                                                    var v81 uint32 = v61
                                                                    if (828 < 255) {
                                                                        var v82 int32 = 1
                                                                        var v83 int8 = (v2 << ((619) & 63))
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            v61 = v61
                                                                        }
                                                                        _ = v82
                                                                        _ = v83
                                                                    }
                                                                    _ = v81
                                                                default:
                                                                    var v84 int16 = 42
                                                                    _ = v84
                                                                }
                                                            }
                                                        }
                                                        fmt.Println(v61, fi49, v61)
                                                        v2 -= v2
                                                    }
                                                    if false {
                                                        var v85 uint = 0
                                                        var v86 uint64 = 42
                                                        _ = v85
                                                        _ = v86
                                                    } else {
                                                        switch "\n\t" {
                                                        case "hello":
                                                            v87 := v3["foo bar"]
                                                            _ = v87
                                                        case "a":
                                                            v88 := GenS7{F0: 32767, F1: 0, F2: 181}
                                                            _ = v88
                                                        default:
                                                            if (42 != 18446744073709551615) {
                                                                switch ("a" + "hello") {
                                                                case "hello":
                                                                    if (54 == v61) {
                                                                        v89 := GenS13{F0: 42}
                                                                        _ = v89
                                                                    }
                                                                case "":
                                                                    var v90 int64 = 42
                                                                    _ = v90
                                                                case "a":
                                                                    v61 = (v61 % (((v61 / ((2147483647) | 1))) | 1))
                                                                default:
                                                                    v91 := GenS1{F0: 0}
                                                                    _ = v91
                                                                }
                                                            } else {
                                                                v92 := GenS5{F0: -128, F1: "a"}
                                                                if (!true) {
                                                                    var v93 bool = (1 == 42)
                                                                    _ = v93
                                                                }
                                                                _ = v92
                                                            }
                                                        }
                                                        v94 := GenS10{F0: ((0.0 / 0.1) == 3.14159), F1: -512}
                                                        _ = v94
                                                    }
                                                    _ = v61
                                                } else {
                                                    if (1 <= 271) {
                                                        v95 := GenS10{F0: true, F1: -1}
                                                        _ = v95
                                                    }
                                                    if (!false) {
                                                        v96 := v3[("" + (("" + "a") + "hello"))]
                                                        var v97 string = ""
                                                        v98 := GenS8{F0: 1, F1: 65535}
                                                        v99 := GenS4{F0: "\n\t", F1: 0}
                                                        v53 -= 1e10
                                                        var v100 float32 = 0.0
                                                        _ = v100
                                                        _ = v96
                                                        _ = v97
                                                        _ = v98
                                                        _ = v99
                                                    } else {
                                                        var v101 int16 = -158
                                                        v1[((("a" + (("" + "") + ("" + "hello"))) + "") + "foo bar")] = genFn0()
                                                        v102 := GenS7{F0: -14, F1: 9223372036854775807, F2: 2}
                                                        _ = v101
                                                        _ = v102
                                                    }
                                                }
                                                v103 := GenS14{F0: 2, F1: "a", F2: 0}
                                                v104 := GenS9{F0: true, F1: (!(65535 > 42))}
                                                v53 -= v53
                                                v105 := GenS15{F0: 1, F1: false}
                                                _ = v103
                                                _ = v104
                                                _ = v105
                                            }
                                        default:
                                            var v106 uint64 = 281
                                            _ = v106
                                        }
                                        v53 = 3.14159
                                        v2 -= (v2 ^ genFn1(32767))
                                    }
                                case "a":
                                    v53 -= (((((((-2.25 / 0.1) - (0.0 - 0.0)) / float32(1e10)) / (((v53 * 0.0) - -2.25) + v53)) / v53) * float32((1e10 * 0.1))) / (v53 + (-2.25 - (v53 - v53))))
                                case "\n\t":
                                    var v107 int16 = 42
                                    _ = v107
                                }
                                var v108 uint32 = 4294967295
                                var v109 uint16 = uint16(v108)
                                _ = v108
                                _ = v109
                                _ = v53
                            }
                        } else {
                            if (!true) {
                                var v110 float32 = (3.14159 / (float32(1e10) + 3.14159))
                                fmt.Println(v110, v2, fi49)
                                _ = v110
                            } else {
                                v2 -= v2
                                v2 = (v2 / ((v2) | 1))
                                if (!(true && (2 <= 701))) {
                                    if true {
                                        v2 += genFn1(-95)
                                    }
                                    var v111 string = ("" + ("hello" + "hello"))
                                    switch genFn0() {
                                    case 9223372036854775807:
                                        v111 = v111
                                    }
                                    var v112 uint8 = 128
                                    v111 = (((v111 + "hello") + "hello") + "hello")
                                    v1[(v111 + v111)] = fi49
                                    _ = v111
                                    _ = v112
                                } else {
                                    var v113 uint16 = 255
                                    v113 = (^(v113))
                                    var v114 float64 = 0.1
                                    _ = v113
                                    _ = v114
                                }
                            }
                        }
                        _ = fi49
                    }
                    _ = v48
                }
                v2 = (^(v2))
                _ = fi37
            }
            var v115 int16 = -32768
            _ = v115
        }
    }
    fmt.Println(v2)
    for fi116 := 0; fi116 < 16; fi116++ {
        if (-1 != v2) {
            var v117 int8 = 1
            var v118 uint64 = 518
            _ = v117
            _ = v118
        }
        for fi119 := 0; fi119 < 6; fi119++ {
            v2 = v2
            _ = fi119
        }
        _ = fi116
    }
    for fi120 := 0; fi120 < 1; fi120++ {
        v2 -= 120
        for fi121 := 0; fi121 < 10; fi121++ {
            v2 = v2
            _ = fi121
        }
        _ = fi120
    }
    v2 -= v2
    switch 1 {
    case 32767:
        v2 -= (v2 % (((-(v2))) | 1))
    }
    v122 := GenS6{F0: false, F1: -32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    _ = v1
    _ = v122
    _ = v2
    _ = v3
    _ = v32
    _ = v33
}
