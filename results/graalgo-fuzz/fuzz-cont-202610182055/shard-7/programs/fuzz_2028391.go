package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int8 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int8; F1 string; F2 uint64 }
type GenS5 struct { F0 uint8; F1 uint32; F2 string }
type GenS6 struct { F0 int32; F1 int16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint32; F1 bool }
type GenS9 struct { F0 int8; F1 int16 }
type GenS10 struct { F0 uint64; F1 uint8; F2 uint }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 int64; F1 uint32 }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 bool; F1 string }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 uint16; F1 uint32 }
type GenS19 struct { F0 string; F1 int; F2 string }
type GenS20 struct { F0 int32 }
type GenS21 struct { F0 int32; F1 uint64; F2 uint32 }
func genFn0(p0 int, p1 int16, p2 string) bool {
    var v6 int8 = 15
    v7 := GenS1{F0: int(p1), F1: -534, F2: (v6 - (v6 | 100))}
    for fi8 := 0; fi8 < 16; fi8++ {
        fmt.Println(p0, p1, p2)
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return false
}
func genFn1(p0 string, p1 int32, p2 uint8) uint8 {
    v9 := GenS1{F0: 127, F1: -32768, F2: -65}
    if (42 == 42) {
        v10 := GenS1{F0: 0, F1: -675, F2: 96}
        if true {
            var v11 float32 = 0.0
            v11 -= 1.5
            for fi12 := 0; fi12 < 4; fi12++ {
                v13 := GenS1{F0: (fi12 + fi12), F1: 1000, F2: 100}
                var v14 string = "\n\t"
                p0 = "foo bar"
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = v11
        }
        _ = v10
    }
    v15 := []int{4294967295, 32768, 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v9
    return (p2 << ((4294967295) & 63))
}

func main() {
    var v1 float32 = (0.1 * (0.0 + 1e10))
    var v2 uint64 = 2
    var v3 uint16 = 32767
    var v4 uint16 = 0
    v5 := make(map[string]string)
    switch 9223372036854775807 {
    case -2147483648:
        _ = len(v5)
    }
    if (42 <= 1) {
        if (false || (!(!false))) {
            for fi16 := 0; fi16 < 6; fi16++ {
                v17 := GenS2{F0: "a", F1: true}
                _ = fi16
                _ = v17
            }
            var v18 uint16 = v4
            _ = v18
        }
        v1 = v1
        var v19 int32 = 313
        _ = v19
    } else {
        for fi20 := 0; fi20 < 12; fi20++ {
            v5["hello"] = "hello"
            v21 := GenS3{F0: 1, F1: "foo bar"}
            for fi22 := 0; fi22 < 6; fi22++ {
                switch ("" + "hello") {
                case "":
                    var v23 uint = 2
                    _ = v23
                case "hello":
                    var v24 int = 0
                    if false {
                        var v25 uint32 = 1
                        fmt.Println(v2, fi22)
                        _ = v25
                    } else {
                        var v26 int8 = 1
                        _ = v26
                    }
                    _ = v24
                case "foo bar":
                    v27 := []int{fi20}
                    _ = v27
                default:
                    var v28 int32 = -766
                    switch fi20 {
                    case 418:
                        var v29 int16 = 808
                        _ = v29
                    }
                    _ = v28
                }
                delete(v5, (("a" + "") + ""))
                _ = fi22
            }
            _ = fi20
            _ = v21
        }
        if ((!(!false)) || true) {
            for fi30 := 0; fi30 < 5; fi30++ {
                if (1000 < 32767) {
                    if (!(107 <= 0)) {
                        fmt.Println(v4, v3, v4)
                        var v31 int16 = -490
                        v3 -= 42
                        _ = v31
                    } else {
                        v32 := GenS3{F0: 0, F1: "\n\t"}
                        var v33 uint32 = 2
                        _ = v32
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 16; fi34++ {
                        var v35 bool = true
                        v36 := []int{int(v4), (fi30 >> ((2) & 63)), fi34}
                        if (1 >= 65535) {
                            switch "" {
                            case "a":
                                if (!false) {
                                    if len(v36) > 0 {
                                        v36[0] = 42
                                    }
                                    if v35 {
                                        v2 = v2
                                    }
                                }
                            }
                            if v35 {
                                fmt.Println(v35, v3)
                            } else {
                                if v35 {
                                    var v37 string = ""
                                    v3 -= 130
                                    _ = v37
                                }
                                var v38 int = (fi34 / (((fi34 | fi30)) | 1))
                                _ = v38
                            }
                        }
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    v3 -= v3
                }
                for fi39 := 0; fi39 < 14; fi39++ {
                    v2 += v2
                    _ = fi39
                }
                _ = fi30
            }
        } else {
            switch 2147483647 {
            case 42:
                v40 := []int{4294967295}
                _ = v40
            default:
                switch "hello" {
                case "hello":
                    var v41 int64 = -9223372036854775808
                    _ = v41
                }
            }
            for fi42 := 0; fi42 < 13; fi42++ {
                if false {
                    v43 := GenS2{F0: "\n\t", F1: (1 >= 1)}
                    _ = v43
                }
                var v44 uint = 255
                switch "hello" {
                case "\n\t":
                    var v45 uint = (v44 / ((v44) | 1))
                    _ = v45
                case "foo bar":
                    v46 := []int{-52, (fi42 | (fi42 << (((v44 * 993)) & 63)))}
                    _ = v46
                case "hello":
                    var v47 int8 = -1
                    _ = v47
                }
                _ = fi42
                _ = v44
            }
            v48 := GenS1{F0: 0, F1: 548, F2: -38}
            _ = v48
        }
        for fi49 := 0; fi49 < 17; fi49++ {
            switch "\n\t" {
            case "a":
                v50 := GenS2{F0: ("hello" + ("a" + (("" + "foo bar") + "hello"))), F1: true}
                _ = v50
            default:
                v51 := GenS4{F0: 42, F1: "\n\t", F2: 910}
                _ = v51
            }
            _ = fi49
        }
    }
    for fi52 := 0; fi52 < 18; fi52++ {
        fmt.Println(v3, v4)
        if ((fi52 == fi52) && ((1e10 / ((v1 / ((v1 * v1) - 3.14159)) + v1)) == -2.25)) {
            fmt.Println(v4, fi52)
            v53 := GenS1{F0: (fi52 ^ (fi52 << ((0) & 63))), F1: 225, F2: 42}
            _ = v53
        } else {
            if false {
                v3 = (v4 / ((255) | 1))
                if false {
                    v4 = 32767
                    for fi54 := 0; fi54 < 20; fi54++ {
                        var v55 int32 = 0
                        _ = fi54
                        _ = v55
                    }
                }
            } else {
                var v56 float32 = 0.0
                var v57 uint32 = 1
                var v58 uint32 = (^(v57))
                _ = v56
                _ = v57
                _ = v58
            }
        }
        switch (-(fi52)) {
        case 4294967295:
            var v59 int64 = -813
            _ = v59
        case 9223372036854775807:
            fmt.Println(v4, v3, v3)
        }
        _ = fi52
    }
    v4 = (v4 & (v3 * v3))
    var v60 uint64 = v2
    if (!true) {
        var v61 uint8 = 0
        v1 = (v1 - float32((3.14159 + 0.1)))
        for fi62 := 0; fi62 < 15; fi62++ {
            switch "a" {
            case "\n\t":
                for fi63 := 0; fi63 < 17; fi63++ {
                    v64 := GenS5{F0: v61, F1: 1, F2: (("" + "a") + "\n\t")}
                    fmt.Println(v61)
                    var v65 uint16 = (v3 | (^(v3)))
                    _ = fi63
                    _ = v64
                    _ = v65
                }
            default:
                v61 = (v61 * (v61 & uint8(v3)))
            }
            delete(v5, "")
            switch ("\n\t" + "\n\t") {
            case "a":
                for fi66 := 0; fi66 < 19; fi66++ {
                    fmt.Println(v4)
                    v67 := GenS1{F0: (fi66 % ((127) | 1)), F1: 1, F2: 0}
                    fmt.Println(fi62, fi66)
                    _ = fi66
                    _ = v67
                }
            case "hello":
                fmt.Println(v2)
            }
            _ = fi62
        }
        _ = v61
    }
    for fi68 := 0; fi68 < 9; fi68++ {
        var v69 int64 = -1
        v4 = 1
        v70 := v5[("" + "\n\t")]
        _ = fi68
        _ = v69
        _ = v70
    }
    fmt.Println(v3, v3)
    if true {
        v71 := GenS5{F0: genFn1("\n\t", -257, 127), F1: 2, F2: "hello"}
        v4 += (v4 & (v4 << ((902) & 63)))
        _ = v71
    }
    delete(v5, "a")
    v1 = (3.14159 * ((((0.1 / v1) * float32((float64(v1) - 1e10))) + v1) - v1))
    if ((^(v3)) <= v3) {
        v2 -= (^(v2))
        for fi72 := 0; fi72 < 8; fi72++ {
            v1 = v1
            var v73 uint32 = 309
            v74 := GenS6{F0: -174, F1: 1}
            _ = fi72
            _ = v73
            _ = v74
        }
    }
    var v75 uint16 = (v3 - (v4 << ((0) & 63)))
    v75 = 1
    var v76 float64 = float64(v1)
    for fi77 := 0; fi77 < 7; fi77++ {
        if (true || false) {
            switch fi77 {
            case 2147483647:
                v3 -= (v4 * (v75 << ((1) & 63)))
            }
            for fi78 := 0; fi78 < 16; fi78++ {
                var v79 int16 = 514
                var v80 float64 = v76
                _ = fi78
                _ = v79
                _ = v80
            }
            if true {
                for fi81 := 0; fi81 < 17; fi81++ {
                    switch 1 {
                    case 57:
                        if (((fi77 / ((fi77) | 1)) == (fi77 << ((1) & 63))) && (41 != 42)) {
                            v82 := GenS6{F0: 42, F1: 42}
                            fmt.Println(v76, v3)
                            var v83 int64 = -1
                            if ((v60 - v2) != v60) {
                                v3 += 255
                                var v84 uint32 = 2
                                _ = v84
                            } else {
                                var v85 string = "a"
                                _ = v85
                            }
                            fmt.Println(v1, v4)
                            delete(v5, "\n\t")
                            _ = v82
                            _ = v83
                        }
                    case 2147483647:
                        if true {
                            v86 := GenS7{F0: (27 >= -29)}
                            switch "a" {
                            case "":
                                if false {
                                    v87 := []int{fi81, fi81}
                                    _ = v87
                                }
                                fmt.Println(v4, v1, v2)
                            case "foo bar":
                                if (("\n\t" + "\n\t") <= "hello") {
                                    fmt.Println(fi77, fi81)
                                    var v88 bool = false
                                    _ = v88
                                } else {
                                    var v89 uint16 = 65535
                                    if false {
                                        v1 = (1.5 / v1)
                                    }
                                    _ = v89
                                }
                            case "hello":
                                var v90 string = ("a" + (((("foo bar" + "\n\t") + "") + "a") + ""))
                                _ = v90
                            }
                            switch fi77 {
                            case 128:
                                v3 = 255
                                v1 = float32((1e10 * (float64((((((float32(v76) + 3.14159) * (v1 / ((v1 * ((v1 * v1) / v1)) * v1))) - v1) / v1) - (v1 * ((1.5 * 3.14159) - 0.1)))) / float64((v1 * v1)))))
                            case -1:
                                switch ("" + (("foo bar" + ("foo bar" + "\n\t")) + "")) {
                                case "\n\t":
                                    var v91 uint16 = (v75 / ((v75) | 1))
                                    _ = v91
                                }
                            case -607:
                                fmt.Println(fi81, v1, v3)
                            default:
                                v2 += v60
                            }
                            _ = v86
                        } else {
                            v92 := []int{42, 32768, -9223372036854775808}
                            v93 := GenS1{F0: -9223372036854775808, F1: 107, F2: int8(v3)}
                            fmt.Println(v1, v75)
                            _ = v92
                            _ = v93
                        }
                    case 32767:
                        v2 = 2
                    }
                    switch fi81 {
                    case 32768:
                        var v94 uint64 = v60
                        if false {
                            v95 := GenS8{F0: 2, F1: false}
                            if false {
                                var v96 int8 = 27
                                _ = v96
                            } else {
                                if true {
                                    v97 := GenS9{F0: -113, F1: int16(fi81)}
                                    v3 += v3
                                    if ("hello" != "") {
                                        var v98 bool = false
                                        v99 := GenS4{F0: 100, F1: ((((("" + "\n\t") + ("\n\t" + (((((("hello" + "a") + ((("a" + (("" + "a") + ("foo bar" + ("" + ("foo bar" + "\n\t"))))) + ("hello" + "foo bar")) + "")) + "hello") + ((("foo bar" + "hello") + "a") + "a")) + "") + ("\n\t" + ("hello" + "hello"))))) + "hello") + "foo bar") + (("hello" + ("hello" + ("\n\t" + "a"))) + ((("\n\t" + ((("\n\t" + "\n\t") + ("\n\t" + ("foo bar" + "a"))) + "\n\t")) + ("foo bar" + ("\n\t" + ((("" + "foo bar") + "a") + "foo bar")))) + ((("\n\t" + "") + "foo bar") + "foo bar")))), F2: 0}
                                        _ = v98
                                        _ = v99
                                    }
                                    var v100 uint32 = 0
                                    _ = v100
                                    _ = v97
                                } else {
                                    var v101 bool = (!true)
                                    v102 := GenS10{F0: (^(v60)), F1: 169, F2: uint(v4)}
                                    fmt.Println(fi77)
                                    if true {
                                        v75 = (^(v3))
                                    } else {
                                        v103 := v5["foo bar"]
                                        _ = v103
                                    }
                                    _ = v101
                                    _ = v102
                                }
                                if (!("hello" == "")) {
                                    v104 := GenS2{F0: "a", F1: (-1 > 32767)}
                                    switch "hello" {
                                    case "":
                                        if false {
                                            if (0.0 < (1e10 + (v76 / 0.1))) {
                                                var v105 uint = 738
                                                fmt.Println(fi81, v1)
                                                _ = v105
                                            }
                                        } else {
                                            v106 := GenS11{F0: (v75 % ((v75) | 1))}
                                            var v107 int16 = int16(fi81)
                                            switch "foo bar" {
                                            case "\n\t":
                                                v76 = float64((3.14159 * ((0.0 * (v1 - (-2.25 * float32(v76)))) + -2.25)))
                                            case "foo bar":
                                                v60 = v94
                                            }
                                            _ = v106
                                            _ = v107
                                        }
                                    case "a":
                                        v75 -= v75
                                    default:
                                        fmt.Println(v60, v4, v60)
                                    }
                                    if (42 <= -214) {
                                        var v108 int16 = 0
                                        _ = v108
                                    } else {
                                        v3 = 255
                                        v1 = v1
                                    }
                                    _ = v104
                                } else {
                                    switch (fi81 % ((fi81) | 1)) {
                                    case 9223372036854775807:
                                        v3 -= 2
                                    default:
                                        v60 -= v2
                                    }
                                    switch "" {
                                    case "a":
                                        v109 := GenS12{F0: (^(v60))}
                                        _ = v109
                                    default:
                                        var v110 int16 = 909
                                        if false {
                                            fmt.Println(v2, fi77)
                                            v4 = (v4 % (((v4 | 2)) | 1))
                                        } else {
                                            var v111 float32 = (v1 - v1)
                                            v112 := GenS13{F0: v110}
                                            fmt.Println(v75, v110, v94)
                                            _ = v111
                                            _ = v112
                                        }
                                        if true {
                                            v113 := GenS6{F0: 0, F1: (v110 + (v110 - 1000))}
                                            v114 := []int{}
                                            if (fi81 > fi77) {
                                                v76 += -2.25
                                            }
                                            if false {
                                                if (!false) {
                                                    if false {
                                                        fmt.Println(v1)
                                                        var v115 string = ("" + "hello")
                                                        _ = v115
                                                    }
                                                } else {
                                                    _ = len(v5)
                                                    if ((^(v4)) > (v75 + (v4 >> ((2) & 63)))) {
                                                        if len(v114) >= 2 {
                                                            v114 = v114[1:len(v114)]
                                                        }
                                                        if (true || false) {
                                                            var v116 uint32 = 2
                                                            fmt.Println(fi81, v75, v60)
                                                            var v117 int64 = 9223372036854775807
                                                            var v118 uint = 255
                                                            v119 := GenS4{F0: 0, F1: "hello", F2: (v60 * (^(v2)))}
                                                            if len(v114) >= 2 {
                                                                v114 = v114[1:len(v114)]
                                                            }
                                                            _ = v116
                                                            _ = v117
                                                            _ = v118
                                                            _ = v119
                                                        } else {
                                                            if true {
                                                                if true {
                                                                    if true {
                                                                        if (v60 != 37) {
                                                                            if ("\n\t" == "\n\t") {
                                                                                var v120 int = (fi77 | fi81)
                                                                                _ = v120
                                                                            }
                                                                        }
                                                                    } else {
                                                                        if true {
                                                                            if (128 <= 38) {
                                                                                var v121 float32 = 1e10
                                                                                _ = v121
                                                                            } else {
                                                                                v4 = 750
                                                                                if true {
                                                                                    var v122 int64 = 42
                                                                                    v123 := v5["foo bar"]
                                                                                    _ = v122
                                                                                    _ = v123
                                                                                } else {
                                                                                    v75 += 1
                                                                                    if true {
                                                                                        var v124 int16 = (v110 ^ (v110 - (v110 + (v110 / ((v110) | 1)))))
                                                                                        switch "hello" {
                                                                                        case "":
                                                                                            delete(v5, "")
                                                                                        }
                                                                                        var v125 uint16 = (v75 | uint16(v60))
                                                                                        var v126 bool = (false || true)
                                                                                        _ = v124
                                                                                        _ = v125
                                                                                        _ = v126
                                                                                    } else {
                                                                                        if false {
                                                                                            var v127 int = fi77
                                                                                            fmt.Println(v94)
                                                                                            var v128 float32 = 3.14159
                                                                                            _ = v127
                                                                                            _ = v128
                                                                                        } else {
                                                                                            switch (fi77 + fi77) {
                                                                                            case 341:
                                                                                                if len(v114) >= 2 {
                                                                                                    v114 = v114[1:len(v114)]
                                                                                                }
                                                                                            case 127:
                                                                                                if len(v114) > 0 {
                                                                                                    v114[0] = fi77
                                                                                                }
                                                                                            case 2147483647:
                                                                                                switch ("a" + "") {
                                                                                                case "foo bar":
                                                                                                    if (1 != int(v94)) {
                                                                                                        switch "" {
                                                                                                        case "foo bar":
                                                                                                            v129 := GenS10{F0: (v94 % (((v60 - 0)) | 1)), F1: uint8(v2), F2: 21}
                                                                                                            _ = v129
                                                                                                        case "hello":
                                                                                                            var v130 float32 = ((1e10 / 0.1) * 1.5)
                                                                                                            _ = v130
                                                                                                        }
                                                                                                        var v131 float32 = float32(((v76 * float64(v1)) + 0.0))
                                                                                                        _ = v131
                                                                                                    } else {
                                                                                                        switch "foo bar" {
                                                                                                        case "":
                                                                                                            if true {
                                                                                                                var v132 uint32 = 240
                                                                                                                _ = v132
                                                                                                            } else {
                                                                                                                fmt.Println(v110, v76, fi81)
                                                                                                                fmt.Println(v4, v75, v110)
                                                                                                            }
                                                                                                        default:
                                                                                                            if (true || true) {
                                                                                                                var v133 uint32 = 93
                                                                                                                _ = v133
                                                                                                            } else {
                                                                                                                if len(v114) >= 2 {
                                                                                                                    v114 = v114[1:len(v114)]
                                                                                                                }
                                                                                                                if true {
                                                                                                                    var v134 int32 = 1
                                                                                                                    switch "\n\t" {
                                                                                                                    case "hello":
                                                                                                                        v135 := GenS14{F0: 1, F1: 1}
                                                                                                                        _ = v135
                                                                                                                    case "a":
                                                                                                                        if false {
                                                                                                                            v136 := GenS15{F0: 367}
                                                                                                                            fmt.Println(v110)
                                                                                                                            var v137 uint32 = 923
                                                                                                                            _ = v136
                                                                                                                            _ = v137
                                                                                                                        } else {
                                                                                                                            _ = len(v114)
                                                                                                                        }
                                                                                                                    case "foo bar":
                                                                                                                        v114 = append(v114, (fi81 >> ((965) & 63)))
                                                                                                                    }
                                                                                                                    _ = v134
                                                                                                                }
                                                                                                                if (-1 < -108) {
                                                                                                                    _ = len(v5)
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        switch fi81 {
                                                                                                        case 449:
                                                                                                            v138 := GenS5{F0: genFn1("\n\t", -539, 127), F1: 65535, F2: "a"}
                                                                                                            _ = v138
                                                                                                        case 4294967295:
                                                                                                            v75 = v4
                                                                                                        case 32767:
                                                                                                            switch (("foo bar" + (("a" + ("" + "hello")) + ((("\n\t" + "") + "hello") + ("hello" + "foo bar")))) + "\n\t") {
                                                                                                            case "\n\t":
                                                                                                                v3 -= 785
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                case "a":
                                                                                                    var v139 bool = (!(false && ((903 != 553) && (!((-73 > -34) || (!(!true)))))))
                                                                                                    _ = v139
                                                                                                case "":
                                                                                                    if (-22 >= (v110 & 433)) {
                                                                                                        var v140 uint16 = (v4 | 3)
                                                                                                        _ = v140
                                                                                                    } else {
                                                                                                        if (!((9 <= -753) || true)) {
                                                                                                            v2 += (v60 << ((386) & 63))
                                                                                                        } else {
                                                                                                            var v141 bool = (true || false)
                                                                                                            _ = v141
                                                                                                        }
                                                                                                        if true {
                                                                                                            v3 = (v4 | (v4 + 868))
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    var v142 int16 = (-(v110))
                                                                                                    _ = v142
                                                                                                }
                                                                                            default:
                                                                                                var v143 int16 = v110
                                                                                                _ = v143
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    switch "a" {
                                                                                    case "foo bar":
                                                                                        var v144 int64 = 9223372036854775807
                                                                                        _ = v144
                                                                                    case "\n\t":
                                                                                        if false {
                                                                                            var v145 uint16 = 1
                                                                                            var v146 int8 = 100
                                                                                            _ = v145
                                                                                            _ = v146
                                                                                        } else {
                                                                                            v5["foo bar"] = "\n\t"
                                                                                        }
                                                                                    }
                                                                                }
                                                                                if false {
                                                                                    var v147 int64 = 662
                                                                                    v148 := GenS6{F0: 1, F1: int16(v3)}
                                                                                    v149 := GenS3{F0: -109, F1: "foo bar"}
                                                                                    v4 += (v4 % ((v75) | 1))
                                                                                    _ = v147
                                                                                    _ = v148
                                                                                    _ = v149
                                                                                } else {
                                                                                    var v150 float64 = ((v76 + v76) - float64(-2.25))
                                                                                    var v151 int16 = int16(v3)
                                                                                    v110 += 32767
                                                                                    _ = v150
                                                                                    _ = v151
                                                                                }
                                                                            }
                                                                        }
                                                                        v76 = float64((((v1 + 1.5) / float32(((v76 * v76) / (0.1 - v76)))) + 1e10))
                                                                    }
                                                                    _ = len(v114)
                                                                    var v152 int8 = 42
                                                                    _ = v152
                                                                }
                                                                if len(v114) > 0 {
                                                                    v114[0] = (-(fi81))
                                                                }
                                                                _ = len(v114)
                                                            }
                                                        }
                                                    } else {
                                                        fmt.Println(v4)
                                                        if false {
                                                            var v153 bool = (false || true)
                                                            _ = v153
                                                        } else {
                                                            v3 = v4
                                                            var v154 int8 = 1
                                                            if len(v114) > 0 {
                                                                v114[0] = fi77
                                                            }
                                                            var v155 int8 = v154
                                                            _ = v154
                                                            _ = v155
                                                        }
                                                        var v156 uint32 = 470
                                                        if true {
                                                            var v157 int32 = 42
                                                            v75 = v75
                                                            _ = v157
                                                        } else {
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                v158 := GenS13{F0: (v110 - v110)}
                                                                _ = v158
                                                            case "a":
                                                                v159 := GenS16{F0: false, F1: ("a" + ("a" + ""))}
                                                                _ = v159
                                                            }
                                                            fmt.Println(v3, v94, v94)
                                                        }
                                                        _ = v156
                                                    }
                                                }
                                            } else {
                                                if len(v114) > 0 {
                                                    v114[0] = (fi81 - (fi81 ^ 462))
                                                }
                                            }
                                            _ = v113
                                            _ = v114
                                        }
                                        _ = v110
                                    }
                                    v160 := GenS17{F0: (^(v75))}
                                    v5["hello"] = (("\n\t" + "\n\t") + "\n\t")
                                    _ = v160
                                }
                            }
                            v161 := GenS6{F0: 42, F1: 477}
                            _ = v161
                            _ = v95
                        } else {
                            if true {
                                var v162 int32 = 691
                                _ = v162
                            } else {
                                v60 = v2
                                if ((fi77 + fi81) > fi77) {
                                    fmt.Println(v4, v94)
                                } else {
                                    var v163 float32 = 0.1
                                    _ = v163
                                }
                            }
                        }
                        _ = v94
                    }
                    _ = fi81
                }
                _ = len(v5)
                var v164 uint8 = 128
                _ = v164
            } else {
                for fi165 := 0; fi165 < 10; fi165++ {
                    v166 := GenS18{F0: (v4 << ((2) & 63)), F1: 42}
                    var v167 int64 = 157
                    _ = fi165
                    _ = v166
                    _ = v167
                }
                for fi168 := 0; fi168 < 17; fi168++ {
                    if (2147483647 == 370) {
                        switch (("hello" + "\n\t") + "hello") {
                        case "\n\t":
                            var v169 uint32 = 1
                            _ = v169
                        case "":
                            var v170 int32 = 903
                            _ = v170
                        case "foo bar":
                            v4 += v75
                        }
                        v76 += ((((-2.25 * v76) + v76) * v76) + v76)
                    } else {
                        var v171 uint8 = 31
                        switch -1 {
                        case 1:
                            switch fi77 {
                            case 151:
                                switch "" {
                                case "":
                                    v172 := GenS19{F0: ("" + "foo bar"), F1: fi168, F2: ("\n\t" + ("" + "a"))}
                                    _ = v172
                                case "foo bar":
                                    if false {
                                        if (true && (-32768 > -306)) {
                                            switch fi168 {
                                            case -2147483648:
                                                v173 := GenS8{F0: 846, F1: false}
                                                _ = v173
                                            case -1:
                                                v174 := GenS8{F0: 42, F1: (42 == 937)}
                                                _ = v174
                                            default:
                                                if (32767 >= -1) {
                                                    var v175 bool = false
                                                    var v176 int32 = int32(v171)
                                                    _ = v175
                                                    _ = v176
                                                } else {
                                                    var v177 int16 = -32768
                                                    fmt.Println(v75, v75)
                                                    _ = v177
                                                }
                                            }
                                        }
                                    }
                                default:
                                    v75 = v3
                                    var v178 string = "\n\t"
                                    v178 = v178
                                    v76 = v76
                                    switch v178 {
                                    case "\n\t":
                                        var v179 int = 350
                                        _ = v179
                                    default:
                                        switch fi77 {
                                        case 127:
                                            if false {
                                                v180 := []int{(fi77 % ((fi77) | 1)), (fi168 / ((9223372036854775807) | 1)), fi168, fi168}
                                                _ = v180
                                            }
                                        case 32767:
                                            v181 := GenS6{F0: int32(v3), F1: int16(v171)}
                                            _ = v181
                                        case -9223372036854775808:
                                            if (v2 >= 750) {
                                                switch (v178 + v178) {
                                                case "foo bar":
                                                    if (v171 < (v171 - genFn1(((((("hello" + "\n\t") + (v178 + "\n\t")) + v178) + "hello") + v178), 0, v171))) {
                                                        var v182 int = 2147483647
                                                        _ = v182
                                                    }
                                                    v2 -= 2
                                                    var v183 uint16 = v3
                                                    _ = v183
                                                default:
                                                    v184 := []int{}
                                                    _ = v184
                                                }
                                            } else {
                                                var v185 uint64 = (v60 % (((v60 << ((0) & 63))) | 1))
                                                _ = v185
                                            }
                                        }
                                    }
                                    _ = v178
                                }
                            case 128:
                                if false {
                                    fmt.Println(v2, v2)
                                } else {
                                    var v186 int = (fi77 * (^(fi77)))
                                    var v187 int32 = -1
                                    _ = v186
                                    _ = v187
                                }
                            }
                        default:
                            v1 += 1.5
                        }
                        _ = v171
                    }
                    _ = fi168
                }
                for fi188 := 0; fi188 < 2; fi188++ {
                    v189 := []int{int(v2)}
                    _ = len(v189)
                    var v190 int16 = -1
                    _ = fi188
                    _ = v189
                    _ = v190
                }
            }
        } else {
            switch (fi77 ^ fi77) {
            case 32767:
                switch fi77 {
                case -1:
                    for fi191 := 0; fi191 < 12; fi191++ {
                        if true {
                            v76 -= 0.0
                            v192 := GenS12{F0: 4294967295}
                            _ = v192
                        }
                        if false {
                            var v193 float32 = v1
                            switch "hello" {
                            case "foo bar":
                                v194 := GenS20{F0: -1}
                                _ = v194
                            case "hello":
                                v195 := GenS10{F0: uint64(v3), F1: 205, F2: 1}
                                _ = v195
                            case "a":
                                v196 := []int{}
                                var v197 int = fi77
                                _ = v196
                                _ = v197
                            }
                            _ = v193
                        }
                        var v198 float64 = float64(v1)
                        v2 += v2
                        switch ("" + "a") {
                        case "hello":
                            v76 += (v198 * float64(v1))
                            switch "foo bar" {
                            case "foo bar":
                                _ = len(v5)
                            case "a":
                                v199 := GenS21{F0: 42, F1: v2, F2: uint32(v75)}
                                _ = v199
                            }
                        }
                        _ = fi191
                        _ = v198
                    }
                case 32767:
                    if false {
                        delete(v5, "\n\t")
                    } else {
                        v60 = 42
                        var v200 float32 = 0.0
                        var v201 uint = 876
                        _ = v200
                        _ = v201
                    }
                }
            case -130:
                v202 := []int{fi77}
                _ = v202
            case 9223372036854775807:
                fmt.Println(v1, v75, v4)
            }
            for fi203 := 0; fi203 < 6; fi203++ {
                v1 -= v1
                _ = fi203
            }
        }
        var v204 string = "\n\t"
        _ = fi77
        _ = v204
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v60:", v60)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v75
    _ = v76
}
