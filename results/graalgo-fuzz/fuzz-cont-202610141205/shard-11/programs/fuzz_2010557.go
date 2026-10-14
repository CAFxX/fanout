package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint64; F1 string; F2 uint16 }
type GenS3 struct { F0 int16; F1 int64; F2 int }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 int64; F1 int8; F2 uint64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int32; F1 int16 }
type GenS8 struct { F0 bool; F1 int64; F2 uint64 }
type GenS9 struct { F0 bool; F1 uint64 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint64; F1 int32 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 uint16; F1 int16; F2 bool }
type GenS15 struct { F0 int8; F1 uint32 }
type GenS16 struct { F0 int32; F1 uint8 }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 int8; F1 uint; F2 int }
type GenS19 struct { F0 int; F1 int; F2 uint8 }
type GenS20 struct { F0 int16 }
func genFn0() uint32 {
    switch 316 {
    case 4294967295:
        var v4 int64 = 9223372036854775807
        _ = v4
    case -9223372036854775808:
        var v5 bool = true
        _ = v5
    case 1:
        if (-1 <= 990) {
            var v6 int = -9223372036854775808
            var v7 int64 = 9223372036854775807
            _ = v6
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := GenS1{F0: (fi8 << ((42) & 63)), F1: 42}
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 uint16 = 1
            v11 -= v11
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    return 366
}
func genFn1(p0 uint32) uint16 {
    if true {
        p0 = (p0 << ((255) & 63))
    }
    var v12 float64 = 3.14159
    _ = p0
    _ = v12
    return 255
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 0
    var v3 bool = false
    v1[42] = v3
    for fi13 := 0; fi13 < 6; fi13++ {
        switch (fi13 >> ((628) & 63)) {
        case -663:
            v3 = (296 >= 2)
        }
        var v14 uint8 = 255
        switch (("\n\t" + "") + ("hello" + "hello")) {
        case "a":
            v15 := GenS1{F0: -9223372036854775808, F1: v14}
            _ = v15
        default:
            v16 := []int{}
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        v3 = v3
        var v18 float64 = 3.14159
        switch "" {
        case "":
            if false {
                fmt.Println(v2, v18)
                for fi19 := 0; fi19 < 3; fi19++ {
                    v20 := GenS1{F0: fi17, F1: 127}
                    _ = fi19
                    _ = v20
                }
                var v21 int = 32768
                _ = v21
            } else {
                for fi22 := 0; fi22 < 14; fi22++ {
                    v3 = ((v2 << ((1) & 63)) <= v2)
                    if (0 > 4294967295) {
                        switch "a" {
                        case "a":
                            v18 -= (v18 * -2.25)
                        }
                    } else {
                        delete(v1, 9223372036854775807)
                        if v3 {
                            var v23 bool = (((float32((v18 + (0.1 / v18))) + 3.14159) + (float32(1e10) / 0.1)) >= (0.1 / (0.1 + (0.1 / (1.5 - 0.0)))))
                            if v23 {
                                v2 = v2
                                var v24 int = -495
                                fmt.Println(v24, fi22, fi17)
                                var v25 uint64 = (v2 % (((v2 & (v2 ^ v2))) | 1))
                                _ = v24
                                _ = v25
                            } else {
                                v26 := GenS1{F0: (fi17 ^ 127), F1: 106}
                                fmt.Println(v3, v2)
                                v27 := GenS1{F0: fi22, F1: 127}
                                _ = v26
                                _ = v27
                            }
                            _ = v23
                        }
                    }
                    v28 := GenS2{F0: (v2 | (v2 / (((^(v2))) | 1))), F1: ("\n\t" + "hello"), F2: genFn1(genFn0())}
                    fmt.Println(v3)
                    _ = fi22
                    _ = v28
                }
            }
        default:
            v29 := GenS2{F0: (v2 | v2), F1: "", F2: 42}
            _ = v29
        }
        _ = fi17
        _ = v18
    }
    if (!((!false) || v3)) {
        var v30 int = 0
        fmt.Println(v3)
        v31 := v1[(v30 & 9223372036854775807)]
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 3; fi32++ {
        fmt.Println(v2, v3)
        if v3 {
            switch "foo bar" {
            case "hello":
                v3 = (135 < 591)
            case "\n\t":
                if (!(!(!v3))) {
                    switch fi32 {
                    case 1:
                        var v33 bool = (1 == 1)
                        _ = v33
                    case 139:
                        var v34 bool = v3
                        _ = v34
                    case 32767:
                        v35 := GenS1{F0: fi32, F1: 128}
                        _ = v35
                    default:
                        v36 := GenS3{F0: 1, F1: -481, F2: (^(fi32))}
                        _ = v36
                    }
                    var v37 uint64 = v2
                    var v38 float64 = 1e10
                    _ = v37
                    _ = v38
                }
            }
            for fi39 := 0; fi39 < 6; fi39++ {
                if ((v2 * v2) < (v2 % ((v2) | 1))) {
                    if (32767 == 1) {
                        if (false && (246 <= 1)) {
                            switch 106 {
                            case 848:
                                switch (fi32 | (^(fi39))) {
                                case 2147483647:
                                    v40 := []int{fi32, 32768, (fi39 * (fi39 >> ((0) & 63))), fi39}
                                    _ = v40
                                }
                            case 480:
                                switch "" {
                                case "hello":
                                    v41 := v1[42]
                                    _ = v41
                                }
                            default:
                                var v42 int16 = -32768
                                _ = v42
                            }
                            switch (fi32 + fi32) {
                            case 9223372036854775807:
                                if (2 == 49) {
                                    v43 := GenS1{F0: 127, F1: 81}
                                    if true {
                                        v3 = true
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v44 int8 = -1
                                            _ = v44
                                        case "hello":
                                            if (2 > 2147483647) {
                                                var v45 string = "hello"
                                                if v3 {
                                                    if (494 == 42) {
                                                        v46 := []int{}
                                                        v47 := GenS3{F0: 42, F1: 417, F2: fi39}
                                                        v3 = (-104 < 223)
                                                        _ = v46
                                                        _ = v47
                                                    } else {
                                                        _ = len(v1)
                                                    }
                                                }
                                                v48 := GenS4{F0: -476, F1: 1}
                                                _ = v45
                                                _ = v48
                                            } else {
                                                v49 := GenS1{F0: fi39, F1: 123}
                                                if (!v3) {
                                                    fmt.Println(fi32, v3, fi32)
                                                    v2 -= (^(v2))
                                                    switch (fi32 + (fi32 ^ 127)) {
                                                    case -2147483648:
                                                        fmt.Println(v2, v3)
                                                    case -715:
                                                        if ("hello" != "a") {
                                                            var v50 float64 = -2.25
                                                            var v51 bool = ((v2 >> ((2) & 63)) != 1)
                                                            switch -126 {
                                                            case 2147483647:
                                                                var v52 int = (^(fi32))
                                                                switch (-(fi39)) {
                                                                case 208:
                                                                    var v53 string = ""
                                                                    _ = v53
                                                                case -443:
                                                                    var v54 string = "\n\t"
                                                                    _ = v54
                                                                case -163:
                                                                    v55 := GenS5{F0: 0, F1: 42, F2: 42}
                                                                    _ = v55
                                                                default:
                                                                    if ((-124 > 79) && (((float32(-2.25) * 1.5) + (float32(v50) / float32(((0.0 + v50) * (v50 - 1.5))))) <= float32(float64(-2.25)))) {
                                                                        var v56 uint32 = 0
                                                                        v3 = (-1 != 130)
                                                                        var v57 uint32 = v56
                                                                        _ = v56
                                                                        _ = v57
                                                                    }
                                                                    var v58 uint64 = v2
                                                                    _ = v58
                                                                }
                                                                _ = v52
                                                            case 721:
                                                                switch (fi39 + fi39) {
                                                                case -465:
                                                                    switch (("hello" + "a") + "foo bar") {
                                                                    case "hello":
                                                                        v3 = ("foo bar" > "")
                                                                    }
                                                                case -9223372036854775808:
                                                                    v59 := []int{(fi39 * 32767), fi39}
                                                                    v50 -= float64((3.14159 * ((1e10 * 1e10) / 1.0)))
                                                                    _ = v59
                                                                default:
                                                                    v2 = 670
                                                                }
                                                            default:
                                                                if true {
                                                                    switch "a" {
                                                                    case "hello":
                                                                        var v60 uint16 = 42
                                                                        _ = v60
                                                                    case "\n\t":
                                                                        if (0 < 2147483647) {
                                                                            var v61 uint = 65535
                                                                            var v62 string = "\n\t"
                                                                            _ = v61
                                                                            _ = v62
                                                                        }
                                                                    default:
                                                                        var v63 int = (fi32 << ((42) & 63))
                                                                        _ = v63
                                                                    }
                                                                } else {
                                                                    if ("" <= "hello") {
                                                                        if v51 {
                                                                            switch "foo bar" {
                                                                            case "":
                                                                                v64 := GenS2{F0: v2, F1: "", F2: 372}
                                                                                _ = v64
                                                                            case "a":
                                                                                if (-33 != 0) {
                                                                                    delete(v1, 128)
                                                                                    switch -88 {
                                                                                    case 712:
                                                                                        fmt.Println(v51)
                                                                                    }
                                                                                    if (586 != 987) {
                                                                                        var v65 bool = (!(-2147483648 == -179))
                                                                                        if (!v51) {
                                                                                            fmt.Println(v65, v50)
                                                                                            v65 = (42 > 883)
                                                                                            v66 := GenS6{F0: -32768}
                                                                                            switch 32768 {
                                                                                            case 9223372036854775807:
                                                                                                v1[128] = v65
                                                                                            default:
                                                                                                var v67 uint64 = 2
                                                                                                _ = v67
                                                                                            }
                                                                                            _ = v66
                                                                                        }
                                                                                        if (int16(fi39) > -1) {
                                                                                            v68 := []int{(fi32 / ((fi32) | 1))}
                                                                                            _ = v68
                                                                                        } else {
                                                                                            v69 := GenS7{F0: 42, F1: 32767}
                                                                                            switch (-(fi32)) {
                                                                                            case 42:
                                                                                                delete(v1, fi32)
                                                                                            }
                                                                                            _ = v69
                                                                                        }
                                                                                        _ = v65
                                                                                    } else {
                                                                                        var v70 int16 = -914
                                                                                        switch "foo bar" {
                                                                                        case "a":
                                                                                            var v71 int16 = v70
                                                                                            _ = v71
                                                                                        case "foo bar":
                                                                                            v70 += (^(v70))
                                                                                        }
                                                                                        v2 = 1
                                                                                        _ = v70
                                                                                    }
                                                                                }
                                                                            case "foo bar":
                                                                                if v3 {
                                                                                    var v72 int32 = 42
                                                                                    _ = v72
                                                                                } else {
                                                                                    v73 := GenS8{F0: v51, F1: 42, F2: v2}
                                                                                    _ = v73
                                                                                }
                                                                            default:
                                                                                var v74 float32 = 1e10
                                                                                _ = v74
                                                                            }
                                                                            switch "a" {
                                                                            case "foo bar":
                                                                                v51 = v51
                                                                            case "\n\t":
                                                                                switch ("hello" + (("foo bar" + ("a" + ("hello" + ""))) + ("a" + (("hello" + "a") + ("foo bar" + "\n\t"))))) {
                                                                                case "":
                                                                                    var v75 int16 = -1
                                                                                    v3 = true
                                                                                    _ = v75
                                                                                case "foo bar":
                                                                                    v3 = v3
                                                                                    var v76 int = 483
                                                                                    if (((!(!((v76 * (-(v76))) == -1))) || true) && ((!(9223372036854775807 > 42)) || (v51 && (!((v76 - (^(fi32))) > fi39))))) {
                                                                                        v77 := GenS8{F0: (-9223372036854775808 != -1), F1: 1, F2: (v2 & v2)}
                                                                                        _ = v77
                                                                                    } else {
                                                                                        v1[(v76 / ((2147483647) | 1))] = (!(-2147483648 != -190))
                                                                                        v50 = float64(0.0)
                                                                                    }
                                                                                    _ = v76
                                                                                case "a":
                                                                                    v3 = v3
                                                                                default:
                                                                                    switch -397 {
                                                                                    case 32768:
                                                                                        var v78 uint32 = 1
                                                                                        _ = v78
                                                                                    }
                                                                                }
                                                                            default:
                                                                                var v79 uint64 = v2
                                                                                _ = v79
                                                                            }
                                                                        }
                                                                        fmt.Println(v3, v3, v3)
                                                                        var v80 string = "a"
                                                                        _ = v80
                                                                    }
                                                                    var v81 uint = 323
                                                                    _ = v81
                                                                }
                                                            }
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                    case 580:
                                                        var v82 int64 = -9223372036854775808
                                                        _ = v82
                                                    default:
                                                        var v83 uint8 = 128
                                                        _ = v83
                                                    }
                                                    var v84 uint32 = 288
                                                    _ = v84
                                                }
                                                v3 = false
                                                _ = v49
                                            }
                                        case "":
                                            var v85 float32 = (float32(0.0) * 1.5)
                                            switch (^(fi32)) {
                                            case 4294967295:
                                                var v86 bool = v3
                                                _ = v86
                                            default:
                                                v85 += (((3.14159 + 1e10) + (float32(0.0) / 1e10)) + v85)
                                            }
                                            _ = v85
                                        default:
                                            v3 = v3
                                        }
                                        var v87 float64 = 1.5
                                        _ = v87
                                    } else {
                                        switch (("foo bar" + "a") + (("foo bar" + ("" + "\n\t")) + "foo bar")) {
                                        case "a":
                                            if v3 {
                                                if (false || (1 == fi32)) {
                                                    var v88 float64 = (1e10 + 3.14159)
                                                    switch (fi32 ^ 32767) {
                                                    case -65:
                                                        switch (-(fi32)) {
                                                        case 868:
                                                            fmt.Println(v88, fi39, v3)
                                                        case 0:
                                                            if (genFn0() != genFn0()) {
                                                                if (!v3) {
                                                                    v89 := v1[(-(fi39))]
                                                                    v3 = v3
                                                                    var v90 uint8 = 42
                                                                    _ = v89
                                                                    _ = v90
                                                                }
                                                            }
                                                        case 436:
                                                            var v91 uint64 = 2
                                                            _ = v91
                                                        default:
                                                            var v92 int64 = 260
                                                            _ = v92
                                                        }
                                                    }
                                                    _ = v88
                                                }
                                            }
                                        default:
                                            var v93 uint = 540
                                            _ = v93
                                        }
                                    }
                                    fmt.Println(fi39)
                                    _ = v43
                                }
                            case -455:
                                var v94 uint16 = 2
                                _ = v94
                            case -368:
                                if false {
                                    v3 = v3
                                } else {
                                    switch (-(fi39)) {
                                    case 625:
                                        var v95 uint64 = v2
                                        _ = v95
                                    case 909:
                                        var v96 string = ("hello" + "")
                                        _ = v96
                                    case -734:
                                        fmt.Println(fi32, fi32, fi39)
                                    }
                                    var v97 uint32 = genFn0()
                                    v3 = v3
                                    _ = v97
                                }
                            default:
                                v2 -= (v2 ^ v2)
                            }
                        } else {
                            var v98 bool = v3
                            var v99 uint8 = 128
                            _ = v98
                            _ = v99
                        }
                        v100 := []int{2147483647, (fi32 - fi32), (fi32 % ((32768) | 1)), fi32}
                        switch "foo bar" {
                        case "hello":
                            var v101 int16 = -506
                            _ = v101
                        case "\n\t":
                            v3 = (v3 || (!true))
                        case "a":
                            v102 := GenS9{F0: (((2 == 65535) || false) || false), F1: (v2 | (v2 ^ v2))}
                            _ = v102
                        default:
                            switch "a" {
                            case "a":
                                switch (fi32 + (fi39 >> ((65535) & 63))) {
                                case 2147483647:
                                    if (32767 > 255) {
                                        var v103 uint64 = v2
                                        _ = v103
                                    }
                                default:
                                    v100 = append(v100, fi39)
                                }
                            case "hello":
                                if (32767 <= 1) {
                                    var v104 float32 = float32(0.0)
                                    var v105 bool = v3
                                    _ = v104
                                    _ = v105
                                } else {
                                    if len(v100) > 0 {
                                        v100[0] = (-(fi39))
                                    }
                                    var v106 bool = (0.1 <= 3.14159)
                                    _ = v106
                                }
                            case "\n\t":
                                var v107 int64 = -9223372036854775808
                                _ = v107
                            }
                            delete(v1, fi39)
                            v2 -= 2
                        }
                        v100 = append(v100, -9223372036854775808)
                        fmt.Println(v2)
                        _ = v100
                    } else {
                        v3 = (v2 != (v2 * (v2 - v2)))
                        var v108 uint64 = v2
                        fmt.Println(fi39)
                        _ = v108
                    }
                    v109 := []int{fi39}
                    _ = v109
                }
                if false {
                    if (v3 || false) {
                        switch fi39 {
                        case -642:
                            switch "" {
                            case "hello":
                                v110 := []int{fi32, fi39, 32768}
                                v2 += (v2 / ((v2) | 1))
                                _ = v110
                            case "\n\t":
                                var v111 bool = (!((v2 < 501) || (1 == 774)))
                                _ = v111
                            }
                        case -603:
                            var v112 uint16 = 2
                            _ = v112
                        case 279:
                            var v113 uint8 = 2
                            _ = v113
                        default:
                            var v114 int16 = 32767
                            _ = v114
                        }
                        v115 := v1[fi39]
                        var v116 uint16 = 42
                        if v3 {
                            v117 := GenS2{F0: (^(v2)), F1: "\n\t", F2: (^(v116))}
                            v118 := v1[32768]
                            _ = v117
                            _ = v118
                        } else {
                            v119 := []int{fi32, fi39}
                            if (-9223372036854775808 != fi32) {
                                switch (("" + "\n\t") + "\n\t") {
                                case "foo bar":
                                    var v120 float32 = ((0.0 / ((((3.14159 / (0.0 - ((1.5 + ((1e10 - (1.5 + 1.5)) - float32(1.5))) * (1e10 / 1.0)))) / (-2.25 - 1.5)) - 0.1) + (0.0 + 0.0))) / (-2.25 + float32(0.1)))
                                    _ = v120
                                case "\n\t":
                                    var v121 uint64 = v2
                                    _ = v121
                                case "":
                                    v115 = true
                                }
                            } else {
                                switch (("a" + "") + "hello") {
                                case "a":
                                    if false {
                                        var v122 uint = 18446744073709551615
                                        v2 = v2
                                        _ = v122
                                    } else {
                                        var v123 uint32 = 2147483648
                                        v115 = (v115 && (!(1000 < -870)))
                                        if false {
                                            v124 := GenS10{F0: 117}
                                            _ = len(v1)
                                            _ = v124
                                        } else {
                                            var v125 bool = (v2 > v2)
                                            var v126 float32 = 0.1
                                            var v127 int = (fi32 << ((65535) & 63))
                                            var v128 uint16 = (v116 * (^(v116)))
                                            if len(v119) >= 2 {
                                                v119 = v119[1:len(v119)]
                                            }
                                            v129 := GenS9{F0: (v126 >= v126), F1: 4294967295}
                                            _ = v125
                                            _ = v126
                                            _ = v127
                                            _ = v128
                                            _ = v129
                                        }
                                        _ = v123
                                    }
                                case "foo bar":
                                    if (!(!(v3 || v115))) {
                                        if (!v3) {
                                            v115 = v115
                                            var v130 uint64 = v2
                                            switch fi39 {
                                            case -2147483648:
                                                v131 := GenS2{F0: (^(v2)), F1: "a", F2: 345}
                                                _ = v131
                                            case 2147483647:
                                                v3 = ((false || v3) && (-1 > -1))
                                            default:
                                                if (-383 != 0) {
                                                    if v115 {
                                                        _ = len(v1)
                                                        v132 := GenS11{F0: (!((0.0 - 1.5) >= (3.14159 - 0.0)))}
                                                        if (v2 <= (v130 / ((v2) | 1))) {
                                                            v119 = append(v119, fi32)
                                                            if ((0.1 * -2.25) == (0.0 / 1e10)) {
                                                                switch "foo bar" {
                                                                case "":
                                                                    if (!true) {
                                                                        fmt.Println(v116, v2, v3)
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            _ = len(v119)
                                                                        case "\n\t":
                                                                            switch "foo bar" {
                                                                            case "\n\t":
                                                                                v2 += (v2 % ((v2) | 1))
                                                                            case "hello":
                                                                                v119 = append(v119, (fi32 % (((^(fi32))) | 1)))
                                                                            case "foo bar":
                                                                                fmt.Println(v2, fi39, fi32)
                                                                            }
                                                                        case "hello":
                                                                            if (((v2 | v2) > v130) && v115) {
                                                                                v3 = v3
                                                                                v133 := GenS6{F0: -545}
                                                                                if v3 {
                                                                                    if true {
                                                                                        if true {
                                                                                            if (1 >= 0) {
                                                                                                if (-1 <= 636) {
                                                                                                    var v134 uint8 = 2
                                                                                                    if (!true) {
                                                                                                        v119 = append(v119, (fi39 * fi39))
                                                                                                        v134 += (v134 & 2)
                                                                                                    } else {
                                                                                                        if ((uint16(v130) < v116) || v115) {
                                                                                                            if ((^(v134)) >= v134) {
                                                                                                                var v135 string = "\n\t"
                                                                                                                var v136 int = fi39
                                                                                                                switch (("foo bar" + v135) + v135) {
                                                                                                                case "\n\t":
                                                                                                                    v137 := GenS12{F0: 18446744073709551615, F1: 42}
                                                                                                                    var v138 uint16 = v116
                                                                                                                    _ = v137
                                                                                                                    _ = v138
                                                                                                                case "hello":
                                                                                                                    v3 = true
                                                                                                                case "a":
                                                                                                                    _ = len(v1)
                                                                                                                default:
                                                                                                                    if (v115 && v3) {
                                                                                                                        v116 += (v116 + 334)
                                                                                                                    } else {
                                                                                                                        fmt.Println(v116, fi32, v130)
                                                                                                                        v139 := GenS6{F0: 1}
                                                                                                                        _ = v139
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v135
                                                                                                                _ = v136
                                                                                                            }
                                                                                                            var v140 int = fi32
                                                                                                            var v141 string = "foo bar"
                                                                                                            _ = v140
                                                                                                            _ = v141
                                                                                                        }
                                                                                                    }
                                                                                                    var v142 float32 = 0.1
                                                                                                    if v115 {
                                                                                                        var v143 uint8 = v134
                                                                                                        if len(v119) >= 2 {
                                                                                                            v119 = v119[1:len(v119)]
                                                                                                        }
                                                                                                        if len(v119) >= 2 {
                                                                                                            v119 = v119[1:len(v119)]
                                                                                                        }
                                                                                                        _ = v143
                                                                                                    }
                                                                                                    var v144 uint = 2
                                                                                                    _ = v134
                                                                                                    _ = v142
                                                                                                    _ = v144
                                                                                                } else {
                                                                                                    v115 = (true && ((^(v116)) <= v116))
                                                                                                    v3 = true
                                                                                                    v145 := GenS10{F0: 42}
                                                                                                    _ = v145
                                                                                                }
                                                                                                if ((255 != 65535) && true) {
                                                                                                    fmt.Println(v2, v2)
                                                                                                    v2 = (v130 - 1)
                                                                                                    if len(v119) >= 2 {
                                                                                                        v119 = v119[1:len(v119)]
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                v2 = (^(v130))
                                                                                                v119 = append(v119, fi32)
                                                                                                _ = len(v119)
                                                                                                v146 := GenS13{F0: 884}
                                                                                                _ = v146
                                                                                            }
                                                                                            v147 := GenS14{F0: 65535, F1: -1, F2: ((-1 == 100) && v3)}
                                                                                            _ = v147
                                                                                        }
                                                                                        v115 = ((3.14159 - 1.5) < 0.1)
                                                                                        var v148 int16 = 1000
                                                                                        _ = v148
                                                                                    } else {
                                                                                        if (true && v3) {
                                                                                            var v149 uint64 = v130
                                                                                            var v150 int8 = 42
                                                                                            if len(v119) > 0 {
                                                                                                v119[0] = int(v150)
                                                                                            }
                                                                                            _ = v149
                                                                                            _ = v150
                                                                                        } else {
                                                                                            var v151 bool = (uint(v2) <= 2)
                                                                                            if v115 {
                                                                                                switch (fi32 / ((fi32) | 1)) {
                                                                                                case 290:
                                                                                                    if v151 {
                                                                                                        if (65535 != (v116 ^ 915)) {
                                                                                                            v152 := v1[(fi39 | (fi32 * (fi39 % ((fi39) | 1))))]
                                                                                                            if (!(v151 || false)) {
                                                                                                                var v153 int = (fi39 & (fi39 + (fi39 / (((fi32 - 32768)) | 1))))
                                                                                                                v130 -= v2
                                                                                                                if true {
                                                                                                                    v154 := GenS8{F0: (39 > 2), F1: -1, F2: v2}
                                                                                                                    _ = v154
                                                                                                                } else {
                                                                                                                    v155 := GenS6{F0: -32768}
                                                                                                                    var v156 float32 = (0.0 * (0.1 / -2.25))
                                                                                                                    var v157 int32 = 1
                                                                                                                    if true {
                                                                                                                        if v151 {
                                                                                                                            var v158 uint32 = 618
                                                                                                                            _ = v158
                                                                                                                        }
                                                                                                                        v130 -= (v2 >> ((255) & 63))
                                                                                                                    } else {
                                                                                                                        if len(v119) >= 2 {
                                                                                                                            v119 = v119[1:len(v119)]
                                                                                                                        }
                                                                                                                    }
                                                                                                                    delete(v1, fi32)
                                                                                                                    _ = v155
                                                                                                                    _ = v156
                                                                                                                    _ = v157
                                                                                                                }
                                                                                                                _ = v153
                                                                                                            } else {
                                                                                                                if len(v119) > 0 {
                                                                                                                    v119[0] = fi39
                                                                                                                }
                                                                                                                v159 := GenS9{F0: (42 == -631), F1: v130}
                                                                                                                _ = v159
                                                                                                            }
                                                                                                            v2 += (^(v130))
                                                                                                            _ = v152
                                                                                                        } else {
                                                                                                            var v160 uint64 = v130
                                                                                                            fmt.Println(v3, v2)
                                                                                                            _ = v160
                                                                                                        }
                                                                                                        v3 = (2 == 42)
                                                                                                        if (0.0 <= (-2.25 / 1e10)) {
                                                                                                            fmt.Println(v2, fi32)
                                                                                                            v161 := GenS9{F0: (-74 != -224), F1: v130}
                                                                                                            fmt.Println(fi32, v130, fi39)
                                                                                                            v1[(fi39 / ((fi32) | 1))] = (4294967295 >= v130)
                                                                                                            _ = v161
                                                                                                        } else {
                                                                                                            if len(v119) >= 2 {
                                                                                                                v119 = v119[1:len(v119)]
                                                                                                            }
                                                                                                            v162 := GenS11{F0: v115}
                                                                                                            var v163 int64 = 9223372036854775807
                                                                                                            var v164 int = -213
                                                                                                            v165 := v1[(-(v164))]
                                                                                                            _ = v162
                                                                                                            _ = v163
                                                                                                            _ = v164
                                                                                                            _ = v165
                                                                                                        }
                                                                                                        var v166 string = "\n\t"
                                                                                                        v167 := v1[(fi32 << ((18446744073709551615) & 63))]
                                                                                                        _ = v166
                                                                                                        _ = v167
                                                                                                    }
                                                                                                default:
                                                                                                    var v168 int = (fi39 >> ((42) & 63))
                                                                                                    _ = v168
                                                                                                }
                                                                                                v130 -= v130
                                                                                            } else {
                                                                                                fmt.Println(v130, fi32, v116)
                                                                                                var v169 float64 = float64(((((1.5 / ((float32(1.5) + (1e10 / 3.14159)) + (3.14159 / 1.0))) + 0.0) * float32(1.5)) + 1e10))
                                                                                                v169 -= (v169 / (v169 - (-2.25 / v169)))
                                                                                                var v170 uint16 = (v116 - 0)
                                                                                                _ = v169
                                                                                                _ = v170
                                                                                            }
                                                                                            v2 = (v130 & 604)
                                                                                            var v171 uint = 1
                                                                                            if (-1 != 2147483647) {
                                                                                                switch fi39 {
                                                                                                case 0:
                                                                                                    if v151 {
                                                                                                        var v172 int8 = 1
                                                                                                        if len(v119) >= 2 {
                                                                                                            v119 = v119[1:len(v119)]
                                                                                                        }
                                                                                                        v173 := GenS3{F0: 0, F1: 192, F2: 2147483647}
                                                                                                        _ = v172
                                                                                                        _ = v173
                                                                                                    }
                                                                                                default:
                                                                                                    if (-2147483648 > 744) {
                                                                                                        v115 = (!v3)
                                                                                                    } else {
                                                                                                        fmt.Println(v115, v3)
                                                                                                        switch "\n\t" {
                                                                                                        case "hello":
                                                                                                            fmt.Println(v3, v151)
                                                                                                        case "":
                                                                                                            switch fi39 {
                                                                                                            case 9223372036854775807:
                                                                                                                v174 := GenS15{F0: -124, F1: 675}
                                                                                                                if len(v119) >= 2 {
                                                                                                                    v119 = v119[1:len(v119)]
                                                                                                                }
                                                                                                                switch fi32 {
                                                                                                                case 1:
                                                                                                                    var v175 string = ("\n\t" + ("" + ""))
                                                                                                                    _ = v175
                                                                                                                }
                                                                                                                _ = v174
                                                                                                            default:
                                                                                                                v171 = uint(v116)
                                                                                                            }
                                                                                                        case "\n\t":
                                                                                                            fmt.Println(v130, v130)
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                v119 = append(v119, fi39)
                                                                                            } else {
                                                                                                fmt.Println(v3)
                                                                                                var v176 int64 = -995
                                                                                                v177 := GenS16{F0: 42, F1: 128}
                                                                                                _ = v176
                                                                                                _ = v177
                                                                                            }
                                                                                            _ = v151
                                                                                            _ = v171
                                                                                        }
                                                                                        v116 = v116
                                                                                        switch (fi32 * (fi32 & (fi32 >> ((18446744073709551615) & 63)))) {
                                                                                        case 128:
                                                                                            fmt.Println(fi39, v130)
                                                                                        case -9223372036854775808:
                                                                                            v2 -= (v2 / ((42) | 1))
                                                                                        }
                                                                                    }
                                                                                    var v178 int64 = -372
                                                                                    var v179 uint64 = (v130 | 2)
                                                                                    _ = v178
                                                                                    _ = v179
                                                                                } else {
                                                                                    var v180 uint = 73
                                                                                    _ = v180
                                                                                }
                                                                                _ = v133
                                                                            }
                                                                        default:
                                                                            if v3 {
                                                                                var v181 int16 = 1000
                                                                                fmt.Println(v115, v2)
                                                                                fmt.Println(v2, v3)
                                                                                _ = v181
                                                                            } else {
                                                                                var v182 uint32 = 521
                                                                                if ((1e10 - (0.0 / ((1e10 + 3.14159) + ((float32(0.0) / 1.0) * 3.14159)))) > (((0.1 * -2.25) + (((1.5 * 1e10) * 1.5) / 1.0)) + 0.1)) {
                                                                                    v2 = uint64(v182)
                                                                                    v119 = append(v119, int(v130))
                                                                                }
                                                                                _ = v182
                                                                            }
                                                                        }
                                                                        switch -9223372036854775808 {
                                                                        case 128:
                                                                            v130 = (^(v130))
                                                                        case 32768:
                                                                            delete(v1, fi39)
                                                                        case -1:
                                                                            v130 = (v2 % (((v2 & (v2 & v130))) | 1))
                                                                        }
                                                                        v2 -= v130
                                                                        v130 = v2
                                                                    }
                                                                }
                                                                v116 += 42
                                                                fmt.Println(fi39, v2, v116)
                                                            } else {
                                                                var v183 int8 = 1
                                                                var v184 int32 = 2147483647
                                                                _ = v183
                                                                _ = v184
                                                            }
                                                            v1[fi39] = (!false)
                                                        }
                                                        v130 -= v2
                                                        _ = v132
                                                    } else {
                                                        if (!v115) {
                                                            if true {
                                                                var v185 int64 = 1
                                                                _ = v185
                                                            }
                                                        }
                                                        switch ((((("hello" + ("\n\t" + "")) + "\n\t") + ("foo bar" + "foo bar")) + "hello") + ("a" + "foo bar")) {
                                                        case "\n\t":
                                                            var v186 bool = true
                                                            _ = v186
                                                        case "foo bar":
                                                            var v187 uint16 = v116
                                                            var v188 bool = (4294967295 == 65535)
                                                            _ = v187
                                                            _ = v188
                                                        }
                                                    }
                                                } else {
                                                    if (v116 > (v116 + uint16(v130))) {
                                                        delete(v1, (fi39 & fi32))
                                                        switch "" {
                                                        case "hello":
                                                            v130 = v130
                                                        case "\n\t":
                                                            v3 = ((3.14159 / (3.14159 * 3.14159)) != (0.0 / (1e10 / (0.1 - -2.25))))
                                                        case "foo bar":
                                                            if (100 < -128) {
                                                                switch "hello" {
                                                                case "hello":
                                                                    var v189 uint16 = (v116 & 0)
                                                                    _ = v189
                                                                case "":
                                                                    if (9223372036854775807 < fi32) {
                                                                        if false {
                                                                            var v190 uint8 = 42
                                                                            if v115 {
                                                                                if v115 {
                                                                                    var v191 int64 = 0
                                                                                    if len(v119) > 0 {
                                                                                        v119[0] = fi32
                                                                                    }
                                                                                    _ = v191
                                                                                }
                                                                                var v192 int32 = 2147483647
                                                                                _ = v192
                                                                            }
                                                                            switch "hello" {
                                                                            case "hello":
                                                                                fmt.Println(v116, v116)
                                                                            case "foo bar":
                                                                                v3 = v3
                                                                            default:
                                                                                if (v115 && (v3 || false)) {
                                                                                    _ = len(v1)
                                                                                    if v115 {
                                                                                        v2 = (v2 ^ (^(v130)))
                                                                                    }
                                                                                } else {
                                                                                    if (9223372036854775807 == -95) {
                                                                                        v116 = v116
                                                                                    }
                                                                                    v119 = append(v119, (fi32 % ((fi39) | 1)))
                                                                                }
                                                                            }
                                                                            _ = v190
                                                                        }
                                                                        v3 = (0 != -1)
                                                                    } else {
                                                                        fmt.Println(fi32, fi32, v130)
                                                                    }
                                                                default:
                                                                    if v115 {
                                                                        v193 := GenS9{F0: v115, F1: 0}
                                                                        if (65535 < (v116 >> ((42) & 63))) {
                                                                            if len(v119) > 0 {
                                                                                v119[0] = fi39
                                                                            }
                                                                            v2 = v130
                                                                            if false {
                                                                                v115 = (true || true)
                                                                                var v194 int = -751
                                                                                v195 := GenS7{F0: -1, F1: 0}
                                                                                _ = v194
                                                                                _ = v195
                                                                            } else {
                                                                                v196 := GenS16{F0: 1, F1: 128}
                                                                                if (((-293 < 42) || true) || (303 < 1)) {
                                                                                    v116 += (v116 / ((genFn1(344)) | 1))
                                                                                    var v197 uint16 = genFn1(11)
                                                                                    v119 = append(v119, (fi32 * fi39))
                                                                                    _ = v197
                                                                                } else {
                                                                                    _ = len(v1)
                                                                                    var v198 uint16 = (v116 & (v116 >> ((255) & 63)))
                                                                                    _ = v198
                                                                                }
                                                                                _ = v196
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v2)
                                                                            var v199 float32 = (((float32(1e10) + 1e10) * float32((((3.14159 + 3.14159) + 0.0) * 1e10))) - 3.14159)
                                                                            v116 = (v116 + (v116 ^ 32767))
                                                                            if (-121 <= 100) {
                                                                                var v200 float64 = 0.0
                                                                                _ = v200
                                                                            }
                                                                            _ = v199
                                                                        }
                                                                        switch fi39 {
                                                                        case -277:
                                                                            var v201 int16 = 1000
                                                                            _ = v201
                                                                        case 2147483647:
                                                                            switch "" {
                                                                            case "hello":
                                                                                v202 := GenS17{F0: 0}
                                                                                _ = v202
                                                                            case "a":
                                                                                if v115 {
                                                                                    var v203 bool = (v116 == 42)
                                                                                    v204 := GenS18{F0: 42, F1: 1, F2: (fi39 << ((2) & 63))}
                                                                                    v130 = v130
                                                                                    _ = v203
                                                                                    _ = v204
                                                                                } else {
                                                                                    if false {
                                                                                        fmt.Println(v3, fi39)
                                                                                    }
                                                                                }
                                                                            default:
                                                                                v205 := GenS3{F0: 777, F1: 0, F2: fi32}
                                                                                _ = v205
                                                                            }
                                                                        case 360:
                                                                            var v206 string = "a"
                                                                            _ = v206
                                                                        }
                                                                        _ = v193
                                                                    } else {
                                                                        v115 = (-433 == 723)
                                                                        v130 -= v2
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        var v207 int16 = -32768
                                                        _ = v207
                                                    }
                                                }
                                            }
                                            var v208 uint8 = 0
                                            if ((v116 << ((1) & 63)) < 247) {
                                                fmt.Println(fi39)
                                                var v209 int8 = 76
                                                _ = v209
                                            }
                                            _ = v130
                                            _ = v208
                                        }
                                    } else {
                                        fmt.Println(v115, v115)
                                    }
                                default:
                                    switch (fi39 % ((fi39) | 1)) {
                                    case -543:
                                        v3 = (((1 >= -1) || (uint32(v116) != 235)) && (!((fi39 != int(v2)) && ((^(v2)) < 837))))
                                    case 42:
                                        var v210 int16 = int16(v2)
                                        _ = v210
                                    }
                                }
                                v211 := GenS1{F0: fi39, F1: 164}
                                var v212 int64 = -9223372036854775808
                                _ = v211
                                _ = v212
                            }
                            _ = v119
                        }
                        _ = v115
                        _ = v116
                    } else {
                        v3 = (!(0.0 != ((1.5 * (((float64(float32((3.14159 / 1e10))) * (((1e10 * 0.0) + -2.25) + -2.25)) / 1e10) / 1.0)) + (float64((0.0 + float32((1e10 * (1.5 + 3.14159))))) * (((1e10 + 1e10) / -2.25) / -2.25)))))
                        v213 := GenS17{F0: 0}
                        if (int32(v2) != 1) {
                            var v214 int = fi32
                            _ = v214
                        } else {
                            if (!(1 == 2147483648)) {
                                if v3 {
                                    var v215 int16 = -32768
                                    if (!v3) {
                                        v3 = false
                                    }
                                    var v216 int16 = (v215 & v215)
                                    _ = v215
                                    _ = v216
                                }
                                v217 := []int{}
                                fmt.Println(v2, fi39, fi32)
                                if len(v217) >= 2 {
                                    v217 = v217[1:len(v217)]
                                }
                                _ = v217
                            } else {
                                var v218 bool = ((1 == 886) && v3)
                                v219 := GenS19{F0: fi32, F1: fi32, F2: 42}
                                _ = v218
                                _ = v219
                            }
                        }
                        _ = v213
                    }
                } else {
                    v220 := v1[fi39]
                    v221 := GenS9{F0: (127 == 83), F1: 18446744073709551615}
                    var v222 int16 = 1
                    var v223 float32 = (1.5 - (0.1 - 3.14159))
                    _ = v220
                    _ = v221
                    _ = v222
                    _ = v223
                }
                delete(v1, fi39)
                v2 += (v2 / ((v2) | 1))
                v224 := []int{(fi32 ^ fi39), fi39, (fi32 & fi32)}
                v225 := GenS12{F0: (v2 << ((175) & 63)), F1: -720}
                var v226 string = "a"
                _ = fi39
                _ = v224
                _ = v225
                _ = v226
            }
            for fi227 := 0; fi227 < 6; fi227++ {
                fmt.Println(fi227, v2, fi32)
                var v228 int32 = 2147483647
                _ = fi227
                _ = v228
            }
        }
        _ = fi32
    }
    for fi229 := 0; fi229 < 16; fi229++ {
        switch (fi229 * (fi229 & (fi229 ^ (fi229 | fi229)))) {
        case -310:
            switch "a" {
            case "hello":
                v3 = v3
            case "":
                switch "foo bar" {
                case "a":
                    var v230 uint64 = 337
                    _ = v230
                case "\n\t":
                    var v231 uint32 = genFn0()
                    _ = v231
                }
            }
        default:
            v3 = v3
        }
        _ = fi229
    }
    for fi232 := 0; fi232 < 1; fi232++ {
        _ = len(v1)
        for fi233 := 0; fi233 < 9; fi233++ {
            fmt.Println(v2, v3, v3)
            v2 = v2
            v3 = false
            v234 := []int{}
            if ((float64(1.5) * 3.14159) != 0.1) {
                v2 += 2
                v235 := GenS8{F0: false, F1: -1, F2: 544}
                v2 += 76
                var v236 float64 = (-2.25 * 1e10)
                if len(v234) > 0 {
                    v234[0] = fi233
                }
                _ = v235
                _ = v236
            }
            v237 := GenS20{F0: 820}
            _ = fi233
            _ = v234
            _ = v237
        }
        _ = fi232
    }
    var v238 uint8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v238:", v238)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v238
    _ = v3
}
