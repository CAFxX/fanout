package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS5 struct { F0 int16; F1 int16; F2 int }
type GenS6 struct { F0 uint16; F1 uint; F2 int16 }
type GenS7 struct { F0 uint64; F1 int; F2 uint16 }
type GenS8 struct { F0 uint32; F1 uint16 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint; F1 bool }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 int8; F1 int32; F2 uint32 }
type GenS16 struct { F0 uint32 }
func genFn0(p0 int8) int {
    v6 := GenS1{F0: 0, F1: 9223372036854775807, F2: 42}
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 int32 = 1
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 18446744073709551615, F1: -9223372036854775808, F2: 4294967295}
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := GenS1{F0: 255, F1: 42, F2: 379}
        var v12 string = "a"
        for fi13 := 0; fi13 < 16; fi13++ {
            v12 = v12
            var v14 uint64 = 4294967295
            _ = fi13
            _ = v14
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v6
    _ = v9
    return 127
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 289
    v3 := make(map[string]uint32)
    var v4 int8 = 24
    v5 := make(map[string]string)
    v15 := GenS1{F0: 4294967295, F1: 42, F2: 65535}
    if (3.14159 == 0.0) {
        v16 := []int{128}
        v5[((((("hello" + "") + ("\n\t" + ("" + "a"))) + "\n\t") + ("\n\t" + "\n\t")) + ("" + ("hello" + "hello")))] = ""
        _ = v16
    }
    if ((!false) || (9223372036854775807 == 0)) {
        for fi17 := 0; fi17 < 19; fi17++ {
            v4 = v4
            v18 := []int{}
            for fi19 := 0; fi19 < 20; fi19++ {
                v2 += v2
                if false {
                    if (v4 >= v4) {
                        var v20 uint = 2
                        v21 := GenS1{F0: 1, F1: 1, F2: v20}
                        _ = v20
                        _ = v21
                    } else {
                        var v22 int8 = v4
                        fmt.Println(fi19)
                        var v23 uint64 = (v2 ^ 2)
                        v3["foo bar"] = 42
                        _ = v22
                        _ = v23
                    }
                    fmt.Println(fi17, fi17, fi17)
                    if (-2147483648 > -336) {
                        v24 := GenS2{F0: -9223372036854775808, F1: v4}
                        v4 = -108
                        fmt.Println(v2)
                        _ = v24
                    } else {
                        var v25 uint = 42
                        var v26 float32 = -2.25
                        var v27 int64 = 0
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    if false {
                        fmt.Println(fi17)
                        if true {
                            v4 = 100
                            var v28 uint8 = 2
                            var v29 uint16 = 853
                            _ = v28
                            _ = v29
                        }
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            v4 = v4
                        }
                        _ = len(v1)
                        var v30 int64 = 9223372036854775807
                        if (!false) {
                            var v31 int32 = 0
                            var v32 uint32 = 4294967295
                            _ = v31
                            _ = v32
                        } else {
                            switch (fi19 - 9223372036854775807) {
                            case -667:
                                if true {
                                    if true {
                                        if len(v18) >= 2 {
                                            v18 = v18[1:len(v18)]
                                        }
                                        if true {
                                            fmt.Println(fi17, fi19)
                                        } else {
                                            delete(v1, fi17)
                                            delete(v3, ("a" + "\n\t"))
                                            switch "hello" {
                                            case "\n\t":
                                                var v33 uint8 = 42
                                                _ = v33
                                            case "hello":
                                                var v34 int64 = (v30 / ((v30) | 1))
                                                _ = v34
                                            }
                                        }
                                        v35 := GenS3{F0: ((("\n\t" + ("hello" + (("hello" + "a") + "a"))) + "\n\t") + ("hello" + "\n\t"))}
                                        _ = v35
                                    } else {
                                        v5["hello"] = (("hello" + "foo bar") + "\n\t")
                                        fmt.Println(v2, v30, v4)
                                        if (18446744073709551615 <= 255) {
                                            fmt.Println(v30, v2)
                                        } else {
                                            _ = len(v3)
                                            v36 := GenS4{F0: v4, F1: 225, F2: 2147483648}
                                            _ = len(v18)
                                            _ = v36
                                        }
                                        if ((v4 + v4) > 32) {
                                            v30 -= v30
                                            var v37 int8 = v4
                                            _ = v37
                                        }
                                    }
                                    v1[2147483647] = false
                                }
                            default:
                                var v38 float64 = 3.14159
                                var v39 bool = true
                                if v39 {
                                    var v40 int8 = 100
                                    if len(v18) > 0 {
                                        v18[0] = fi17
                                    }
                                    v41 := GenS2{F0: v30, F1: -128}
                                    _ = v40
                                    _ = v41
                                }
                                _ = v38
                                _ = v39
                            }
                        }
                        var v42 uint = 18446744073709551615
                        fmt.Println(v30, v2, v4)
                        _ = v30
                        _ = v42
                    }
                } else {
                    if false {
                        v4 = v4
                    } else {
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                    }
                    delete(v5, "hello")
                }
                v4 = v4
                _ = fi19
            }
            _ = fi17
            _ = v18
        }
        var v43 int8 = v4
        _ = v43
    } else {
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            if (true && false) {
                v44 := GenS5{F0: 1000, F1: 42, F2: 574}
                var v45 int = -75
                _ = v44
                _ = v45
            }
        }
        switch "\n\t" {
        case "hello":
            v46 := []int{127, -2147483648}
            _ = v46
        case "foo bar":
            v47 := []int{32768}
            _ = v47
        case "":
            var v48 float64 = -2.25
            _ = v48
        }
    }
    v4 = (v4 | (v4 + (v4 | 42)))
    var v49 float32 = 1.5
    for fi50 := 0; fi50 < 1; fi50++ {
        if ((v2 ^ v2) < 1) {
            var v51 uint8 = 0
            v49 = v49
            fmt.Println(v4)
            _ = v51
        }
        if false {
            var v52 float32 = (v49 - v49)
            v53 := v1[-9223372036854775808]
            _ = v52
            _ = v53
        } else {
            v4 = (v4 ^ 8)
            for fi54 := 0; fi54 < 19; fi54++ {
                if false {
                    var v55 float64 = (1e10 - 0.0)
                    _ = v55
                }
                v56 := GenS4{F0: (-(v4)), F1: 127, F2: 65535}
                var v57 uint = 656
                var v58 int8 = 1
                _ = fi54
                _ = v56
                _ = v57
                _ = v58
            }
            v59 := GenS6{F0: 567, F1: 42, F2: 0}
            _ = v59
        }
        v2 -= v2
        _ = fi50
    }
    delete(v1, 0)
    var v60 int = -2147483648
    if (!(!false)) {
        switch (((("hello" + ("a" + "hello")) + "hello") + "\n\t") + "hello") {
        case "hello":
            if (!(384 == 1000)) {
                for fi61 := 0; fi61 < 20; fi61++ {
                    fmt.Println(fi61)
                    var v62 int64 = -9223372036854775808
                    _ = fi61
                    _ = v62
                }
                var v63 uint16 = 341
                _ = v63
            } else {
                if (18446744073709551615 >= 65535) {
                    var v64 bool = (false && (-54 <= 1))
                    _ = v64
                }
                v2 -= (^(v2))
            }
        }
        v3["a"] = 1
        fmt.Println(v60, v4)
    } else {
        if true {
            fmt.Println(v49, v49)
            switch "foo bar" {
            case "hello":
                if (667 > (^(v2))) {
                    _ = len(v1)
                    for fi65 := 0; fi65 < 13; fi65++ {
                        if true {
                            var v66 bool = false
                            _ = v66
                        } else {
                            for fi67 := 0; fi67 < 15; fi67++ {
                                if (!(!(!false))) {
                                    v68 := GenS7{F0: (v2 >> ((161) & 63)), F1: genFn0(127), F2: 532}
                                    v4 = (v4 | v4)
                                    v69 := v3[(("foo bar" + "foo bar") + ("hello" + "hello"))]
                                    var v70 int8 = v4
                                    _ = v68
                                    _ = v69
                                    _ = v70
                                } else {
                                    if true {
                                        if true {
                                            if (0 <= 32767) {
                                                fmt.Println(fi67)
                                            }
                                            v49 += (-2.25 + (float32((1e10 - float64((v49 * v49)))) - (1.5 - 3.14159)))
                                            var v71 uint = 4294967295
                                            _ = v71
                                        }
                                    }
                                }
                                switch (fi67 / ((fi67) | 1)) {
                                case 9223372036854775807:
                                    switch 0 {
                                    case 233:
                                        v72 := []int{0, -423, 4294967295, -370}
                                        if false {
                                            v49 = (v49 + (0.0 - 0.0))
                                            var v73 uint = 0
                                            if (!true) {
                                                if (1 < -831) {
                                                    v73 -= (^(v73))
                                                    v74 := GenS1{F0: v73, F1: 42, F2: (v73 << ((v73) & 63))}
                                                    _ = v74
                                                } else {
                                                    fmt.Println(v2, v2, v49)
                                                    switch 128 {
                                                    case 127:
                                                        switch "" {
                                                        case "\n\t":
                                                            _ = len(v72)
                                                        default:
                                                            v2 -= 42
                                                        }
                                                    default:
                                                        if true {
                                                            if ((^(v4)) != v4) {
                                                                if true {
                                                                    var v75 uint = v73
                                                                    _ = v75
                                                                }
                                                                v76 := GenS2{F0: 709, F1: (v4 >> ((65535) & 63))}
                                                                if (v73 != v73) {
                                                                    fmt.Println(v4)
                                                                    fmt.Println(v4)
                                                                } else {
                                                                    var v77 bool = (true && false)
                                                                    fmt.Println(v4, v4, v49)
                                                                    _ = v77
                                                                }
                                                                _ = v76
                                                            }
                                                            switch v60 {
                                                            case 42:
                                                                v60 = -341
                                                            }
                                                        }
                                                    }
                                                    v78 := GenS1{F0: (v73 + v73), F1: -1, F2: v73}
                                                    _ = v78
                                                }
                                            } else {
                                                if (!((667 == 735) && true)) {
                                                    switch -1 {
                                                    case -9223372036854775808:
                                                        var v79 int64 = 290
                                                        v4 += (v4 ^ v4)
                                                        fmt.Println(v73, v49)
                                                        _ = v79
                                                    }
                                                }
                                            }
                                            _ = v73
                                        } else {
                                            v60 -= (v60 + (v60 << ((1) & 63)))
                                            v80 := GenS8{F0: 65535, F1: 32767}
                                            _ = v80
                                        }
                                        switch "hello" {
                                        case "a":
                                            if ((v4 - 0) == (v4 % (((v4 - (v4 | v4))) | 1))) {
                                                var v81 uint8 = 235
                                                if ((!(32767 >= 1000)) || false) {
                                                    v82 := GenS3{F0: ("hello" + "foo bar")}
                                                    switch (("\n\t" + "a") + (("foo bar" + "\n\t") + ((("foo bar" + "a") + (((("\n\t" + "hello") + "hello") + "\n\t") + "foo bar")) + ("" + "a")))) {
                                                    case "":
                                                        v49 += v49
                                                    case "a":
                                                        var v83 bool = ((507 > 255) && ((v2 << ((905) & 63)) > (v2 * (v2 >> ((2) & 63)))))
                                                        _ = v83
                                                    }
                                                    v2 -= (v2 | v2)
                                                    _ = v82
                                                } else {
                                                    fmt.Println(v4)
                                                }
                                                v84 := GenS9{F0: 495}
                                                _ = v81
                                                _ = v84
                                            } else {
                                                v60 += (v60 ^ v60)
                                                if (42 <= 560) {
                                                    _ = len(v3)
                                                    switch 1 {
                                                    case 127:
                                                        switch fi67 {
                                                        case 32768:
                                                            if len(v72) >= 2 {
                                                                v72 = v72[1:len(v72)]
                                                            }
                                                        default:
                                                            var v85 string = "hello"
                                                            _ = v85
                                                        }
                                                    case 946:
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            var v86 uint8 = 255
                                                            _ = v86
                                                        default:
                                                            if true {
                                                                switch "a" {
                                                                case "hello":
                                                                    v4 += (^(v4))
                                                                }
                                                                switch ("foo bar" + "foo bar") {
                                                                case "foo bar":
                                                                    v49 = 1e10
                                                                case "a":
                                                                    if len(v72) >= 2 {
                                                                        v72 = v72[1:len(v72)]
                                                                    }
                                                                case "":
                                                                    switch (("foo bar" + "") + ("\n\t" + "a")) {
                                                                    case "":
                                                                        v60 = (fi67 << ((255) & 63))
                                                                    }
                                                                }
                                                                if true {
                                                                    if (2 != uint32(v2)) {
                                                                        fmt.Println(v2, v60, fi67)
                                                                        v87 := GenS9{F0: -367}
                                                                        switch -108 {
                                                                        case 742:
                                                                            v72 = append(v72, (fi65 >> ((1) & 63)))
                                                                        case 32768:
                                                                            delete(v1, (v60 ^ genFn0(v4)))
                                                                        case -888:
                                                                            fmt.Println(v60)
                                                                        }
                                                                        _ = v87
                                                                    }
                                                                    var v88 uint64 = 141
                                                                    var v89 uint16 = 42
                                                                    _ = v88
                                                                    _ = v89
                                                                }
                                                            }
                                                        }
                                                    default:
                                                        var v90 int64 = 0
                                                        _ = v90
                                                    }
                                                    if false {
                                                        v60 = (v60 | v60)
                                                    }
                                                    v91 := GenS8{F0: 0, F1: uint16(v2)}
                                                    _ = v91
                                                } else {
                                                    v3["a"] = 693
                                                    var v92 uint16 = 255
                                                    _ = v92
                                                }
                                                v93 := GenS4{F0: (v4 ^ v4), F1: 255, F2: 417}
                                                var v94 int = fi67
                                                var v95 uint16 = 1
                                                if len(v72) > 0 {
                                                    v72[0] = genFn0(127)
                                                }
                                                var v96 uint8 = 127
                                                _ = v93
                                                _ = v94
                                                _ = v95
                                                _ = v96
                                            }
                                        case "hello":
                                            v97 := v1[genFn0((v4 - 127))]
                                            _ = v97
                                        case "foo bar":
                                            if len(v72) >= 2 {
                                                v72 = v72[1:len(v72)]
                                            }
                                            var v98 uint8 = 0
                                            switch "a" {
                                            case "":
                                                v99 := GenS8{F0: 787, F1: 0}
                                                _ = v99
                                            default:
                                                _ = len(v72)
                                                v72 = append(v72, (fi67 ^ -2147483648))
                                            }
                                            v100 := GenS7{F0: (v2 % ((1) | 1)), F1: fi67, F2: 56}
                                            if (242 >= 242) {
                                                switch ("\n\t" + "foo bar") {
                                                case "foo bar":
                                                    v4 = (v4 * (v4 & (-(v4))))
                                                case "\n\t":
                                                    if (((v98 >= v98) && false) || true) {
                                                        var v101 uint16 = 255
                                                        if (2147483647 == 0) {
                                                            fmt.Println(v60, v49)
                                                            var v102 uint = 589
                                                            if (!(!(false && (-2147483648 > -9223372036854775808)))) {
                                                                v98 = v98
                                                                v103 := GenS7{F0: (v2 / ((v2) | 1)), F1: fi65, F2: (v101 * uint16(v98))}
                                                                _ = v103
                                                            } else {
                                                                v49 = 0.1
                                                                v2 = 42
                                                                var v104 int = (fi65 << ((728) & 63))
                                                                _ = v104
                                                            }
                                                            v2 -= 1
                                                            fmt.Println(v4, v102)
                                                            fmt.Println(v102)
                                                            _ = v102
                                                        }
                                                        if (394 != 255) {
                                                            v49 += (3.14159 / ((v49 * 0.0) / -2.25))
                                                            v101 -= 42
                                                        }
                                                        _ = v101
                                                    }
                                                default:
                                                    var v105 int8 = 100
                                                    _ = v105
                                                }
                                                if false {
                                                    _ = len(v72)
                                                    var v106 uint64 = (v2 / (((v2 / ((v2) | 1))) | 1))
                                                    _ = v106
                                                } else {
                                                    switch ("hello" + "hello") {
                                                    case "\n\t":
                                                        v60 += fi65
                                                    case "":
                                                        var v107 bool = (!false)
                                                        fmt.Println(v4, v107)
                                                        _ = v107
                                                    case "a":
                                                        v72 = append(v72, (fi65 * 32767))
                                                    default:
                                                        if (!true) {
                                                            if (2 >= 65535) {
                                                                v2 += v2
                                                                var v108 string = ("hello" + ("foo bar" + "a"))
                                                                fmt.Println(fi65, v108)
                                                                _ = v108
                                                            }
                                                        }
                                                    }
                                                }
                                            } else {
                                                v2 = (v2 >> ((255) & 63))
                                            }
                                            var v109 int32 = 584
                                            if true {
                                                var v110 int = int(v4)
                                                var v111 uint64 = (v2 & (v2 / ((4294967295) | 1)))
                                                v112 := GenS3{F0: "\n\t"}
                                                fmt.Println(fi65, fi65, v109)
                                                _ = v110
                                                _ = v111
                                                _ = v112
                                            } else {
                                                v49 = v49
                                                switch (("\n\t" + "a") + "foo bar") {
                                                case "hello":
                                                    var v113 uint16 = 285
                                                    _ = v113
                                                case "a":
                                                    var v114 int32 = 312
                                                    _ = v114
                                                }
                                                fmt.Println(v60, v49)
                                            }
                                            _ = v100
                                            _ = v109
                                            _ = v98
                                        default:
                                            var v115 int = (v60 >> ((1) & 63))
                                            _ = v115
                                        }
                                        _ = v72
                                    case 9223372036854775807:
                                        v2 += 0
                                        switch genFn0(v4) {
                                        case 32767:
                                            _ = len(v3)
                                        case -2147483648:
                                            v4 -= (-(v4))
                                        case -9223372036854775808:
                                            if (0 <= 65535) {
                                                fmt.Println(v2, v2, v49)
                                            }
                                        default:
                                            v49 -= float32(0.1)
                                        }
                                    }
                                case 32768:
                                    v4 += v4
                                case 2147483647:
                                    var v116 int16 = -1
                                    _ = v116
                                default:
                                    v117 := GenS8{F0: 354, F1: uint16(v2)}
                                    _ = v117
                                }
                                v118 := GenS10{F0: 202, F1: false}
                                _ = fi67
                                _ = v118
                            }
                            v49 = float32(1e10)
                        }
                        _ = fi65
                    }
                    for fi119 := 0; fi119 < 14; fi119++ {
                        switch "" {
                        case "foo bar":
                            _ = len(v3)
                        case "\n\t":
                            for fi120 := 0; fi120 < 7; fi120++ {
                                v49 -= v49
                                if (-2147483648 < 368) {
                                    fmt.Println(v60)
                                    var v121 int32 = 131
                                    _ = v121
                                }
                                var v122 uint16 = 2
                                fmt.Println(v122, fi120, v49)
                                v123 := GenS11{F0: v122}
                                var v124 int8 = v4
                                _ = fi120
                                _ = v122
                                _ = v123
                                _ = v124
                            }
                        }
                        var v125 string = (("a" + "") + "hello")
                        _ = fi119
                        _ = v125
                    }
                } else {
                    for fi126 := 0; fi126 < 14; fi126++ {
                        var v127 uint32 = 2147483647
                        for fi128 := 0; fi128 < 7; fi128++ {
                            v60 = fi126
                            if false {
                                v2 += v2
                            } else {
                                var v129 float32 = float32((0.1 * 3.14159))
                                var v130 uint16 = 2
                                if (false || (!false)) {
                                    var v131 int64 = -520
                                    if false {
                                        v132 := []int{fi128, (v60 ^ (v60 * genFn0((^(v4)))))}
                                        if (-1 >= (^(v4))) {
                                            delete(v5, (("\n\t" + "\n\t") + (("\n\t" + ("hello" + "\n\t")) + (("hello" + ("" + "")) + ("foo bar" + "")))))
                                        }
                                        _ = v132
                                    } else {
                                        var v133 int32 = 42
                                        var v134 string = ("a" + "hello")
                                        var v135 float32 = (((v129 / v49) / (((0.1 * v49) / float32(((0.0 - 3.14159) * (1.5 * 1e10)))) - float32(1.5))) + 3.14159)
                                        _ = v133
                                        _ = v134
                                        _ = v135
                                    }
                                    v4 = v4
                                    _ = v131
                                } else {
                                    v127 -= (v127 >> ((65535) & 63))
                                    v4 = (^(v4))
                                    switch "hello" {
                                    case "hello":
                                        v136 := GenS12{F0: (v2 ^ 0)}
                                        _ = v136
                                    }
                                }
                                _ = v129
                                _ = v130
                            }
                            delete(v3, "\n\t")
                            _ = fi128
                        }
                        fmt.Println(v49, v60)
                        _ = fi126
                        _ = v127
                    }
                }
            case "":
                for fi137 := 0; fi137 < 20; fi137++ {
                    if true {
                        for fi138 := 0; fi138 < 16; fi138++ {
                            v139 := []int{2147483647, v60}
                            _ = fi138
                            _ = v139
                        }
                        for fi140 := 0; fi140 < 5; fi140++ {
                            if false {
                                v3[("foo bar" + "\n\t")] = 0
                                var v141 int = fi137
                                _ = v141
                            }
                            v60 = genFn0((v4 - 14))
                            v2 += 2
                            _ = fi140
                        }
                        if false {
                            v49 += v49
                            if false {
                                var v142 float64 = ((float64(-2.25) - float64(3.14159)) + 3.14159)
                                if (255 < 127) {
                                    if false {
                                        v142 = v142
                                    }
                                    v143 := v3[""]
                                    var v144 float64 = v142
                                    _ = v143
                                    _ = v144
                                } else {
                                    if true {
                                        v145 := GenS9{F0: 42}
                                        v49 = ((0.0 * v49) / ((v49 * v49) - v49))
                                        _ = v145
                                    } else {
                                        v146 := []int{int(v2), (v60 & fi137)}
                                        v4 -= (^(v4))
                                        _ = v146
                                    }
                                    v60 -= (v60 % ((v60) | 1))
                                }
                                if true {
                                    if true {
                                        v49 = v49
                                    } else {
                                        var v147 string = "\n\t"
                                        v148 := GenS13{F0: 2147483647, F1: 52, F2: 1}
                                        if true {
                                            v149 := GenS14{F0: 1}
                                            _ = v149
                                        } else {
                                            switch (fi137 + -402) {
                                            case -20:
                                                var v150 bool = (1e10 <= v49)
                                                _ = v150
                                            case 42:
                                                if ((v2 + v2) < (^(v2))) {
                                                    var v151 int8 = (v4 / ((v4) | 1))
                                                    _ = v151
                                                }
                                            default:
                                                v147 = (v147 + "a")
                                            }
                                        }
                                        _ = v147
                                        _ = v148
                                    }
                                    fmt.Println(v4, fi137, v4)
                                    for fi152 := 0; fi152 < 11; fi152++ {
                                        v153 := []int{(^(fi152)), (v60 ^ 128), -884, genFn0((v4 + -128))}
                                        var v154 int = -9223372036854775808
                                        v5["a"] = (("hello" + ("hello" + "foo bar")) + ("" + ""))
                                        _ = fi152
                                        _ = v153
                                        _ = v154
                                    }
                                } else {
                                    delete(v1, (v60 << ((18446744073709551615) & 63)))
                                    v155 := GenS15{F0: v4, F1: 2147483647, F2: 957}
                                    _ = v155
                                }
                                _ = v142
                            }
                            v156 := []int{}
                            _ = v156
                        } else {
                            v157 := GenS9{F0: -1}
                            fmt.Println(v60, v49, fi137)
                            for fi158 := 0; fi158 < 18; fi158++ {
                                v159 := GenS4{F0: (-(v4)), F1: 127, F2: 388}
                                _ = fi158
                                _ = v159
                            }
                            _ = v157
                        }
                    } else {
                        for fi160 := 0; fi160 < 1; fi160++ {
                            v2 += v2
                            switch fi160 {
                            case -9223372036854775808:
                                v161 := []int{(v60 | fi137), v60, -761}
                                _ = v161
                            case -2147483648:
                                v162 := v1[877]
                                _ = v162
                            case 42:
                                v2 += 822
                            default:
                                var v163 int32 = -88
                                _ = v163
                            }
                            _ = fi160
                        }
                        switch "a" {
                        case "foo bar":
                            var v164 int = 2147483647
                            _ = v164
                        case "a":
                            var v165 uint32 = 348
                            _ = v165
                        default:
                            v166 := []int{1}
                            _ = v166
                        }
                        for fi167 := 0; fi167 < 5; fi167++ {
                            fmt.Println(v60)
                            _ = fi167
                        }
                    }
                    var v168 uint32 = 1
                    _ = fi137
                    _ = v168
                }
            }
            v169 := GenS16{F0: 2}
            _ = v169
        } else {
            v49 = float32((0.1 - float64(float32(1e10))))
            v2 -= v2
        }
        v170 := GenS9{F0: 42}
        _ = v170
    }
    fmt.Println(v60, v4, v49)
    var v171 bool = false
    v171 = true
    for fi172 := 0; fi172 < 6; fi172++ {
        v2 = v2
        _ = fi172
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v171:", v171)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v60:", v60)
    _ = v1
    _ = v15
    _ = v171
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v60
}
