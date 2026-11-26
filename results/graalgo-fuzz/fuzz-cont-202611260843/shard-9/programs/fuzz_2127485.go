package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int8; F2 int8 }
type GenS4 struct { F0 int64; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint32; F1 int; F2 uint }
type GenS6 struct { F0 int64; F1 uint64; F2 int16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint16; F1 int16 }
type GenS9 struct { F0 int32; F1 bool }
type GenS10 struct { F0 int8; F1 uint16; F2 int16 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS13 struct { F0 uint16; F1 int16; F2 int }
type GenS14 struct { F0 int8; F1 int16 }
type GenS15 struct { F0 uint; F1 int; F2 uint }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 uint8; F1 uint16 }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 int16 }
type GenS20 struct { F0 int8 }
type GenS21 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS22 struct { F0 int64 }
type GenS23 struct { F0 uint64 }
type GenS24 struct { F0 int8; F1 uint64; F2 int16 }
type GenS25 struct { F0 uint32 }
type GenS26 struct { F0 int32; F1 int; F2 int64 }
type GenS27 struct { F0 int; F1 string }
func genFn0() int8 {
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 uint32 = 1
        _ = fi5
        _ = v6
    }
    var v7 uint = 1
    _ = v7
    return 0
}
func genFn1() bool {
    var v8 float32 = 1e10
    fmt.Println(v8)
    v9 := []int{128}
    var v10 int32 = -1
    _ = v10
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 float64 = (0.0 * 1e10)
    v2 := make(map[string]int)
    var v3 float64 = (v1 * v1)
    v4 := make(map[uint32]bool)
    var v11 int32 = 549
    switch 128 {
    case 9223372036854775807:
        v3 = 0.1
    case 1:
        var v12 string = "foo bar"
        _ = v12
    case -9223372036854775808:
        if ((9223372036854775807 < 42) && (!(!((v11 >> ((621) & 63)) <= (v11 << ((283) & 63)))))) {
            var v13 uint64 = 42
            v14 := GenS1{F0: 42, F1: "a", F2: ("a" + "foo bar")}
            fmt.Println(v13, v11, v13)
            _ = v13
            _ = v14
        }
    default:
        switch -997 {
        case 42:
            var v15 int64 = 597
            _ = v15
        case 897:
            for fi16 := 0; fi16 < 6; fi16++ {
                v17 := GenS2{F0: 274}
                _ = fi16
                _ = v17
            }
        default:
            switch 116 {
            case 0:
                var v18 uint = 255
                _ = v18
            case 1:
                if true {
                    for fi19 := 0; fi19 < 16; fi19++ {
                        for fi20 := 0; fi20 < 19; fi20++ {
                            v21 := []int{fi19, fi20, (fi19 | fi20)}
                            v22 := GenS1{F0: 320, F1: ((("a" + (("" + "a") + "foo bar")) + ("hello" + "hello")) + "\n\t"), F2: "hello"}
                            _ = len(v21)
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                        v3 += v3
                        var v23 int16 = -32768
                        _ = fi19
                        _ = v23
                    }
                } else {
                    for fi24 := 0; fi24 < 18; fi24++ {
                        switch "" {
                        case "\n\t":
                            v25 := []int{42}
                            _ = v25
                        case "foo bar":
                            switch fi24 {
                            case 42:
                                if true {
                                    v26 := GenS3{F0: (fi24 * fi24), F1: genFn0(), F2: 1}
                                    v27 := []int{-302, (-(fi24)), (fi24 / ((fi24) | 1)), 0}
                                    switch ("hello" + (("a" + "") + "a")) {
                                    case "":
                                        v28 := GenS1{F0: 0, F1: ("foo bar" + ("hello" + "a")), F2: ("\n\t" + ("" + "hello"))}
                                        _ = v28
                                    case "\n\t":
                                        if false {
                                            _ = len(v4)
                                            var v29 string = ("hello" + ("foo bar" + ""))
                                            v1 -= v1
                                            _ = v29
                                        }
                                    }
                                    _ = v26
                                    _ = v27
                                } else {
                                    switch fi24 {
                                    case 127:
                                        switch "hello" {
                                        case "foo bar":
                                            var v30 uint64 = 286
                                            _ = v30
                                        case "":
                                            v31 := []int{(fi24 & fi24), -9223372036854775808}
                                            _ = v31
                                        case "\n\t":
                                            v32 := GenS4{F0: 0, F1: 1, F2: 127}
                                            _ = v32
                                        default:
                                            for fi33 := 0; fi33 < 9; fi33++ {
                                                if ((!(0.1 != 3.14159)) || true) {
                                                    if true {
                                                        v34 := GenS4{F0: 185, F1: 42, F2: 1}
                                                        v35 := GenS5{F0: 569, F1: (fi33 & fi33), F2: 1}
                                                        _ = v34
                                                        _ = v35
                                                    } else {
                                                        fmt.Println(v11)
                                                        var v36 float64 = 3.14159
                                                        var v37 int16 = -939
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                    if false {
                                                        var v38 float32 = (-2.25 * 0.0)
                                                        _ = v38
                                                    }
                                                } else {
                                                    v3 += 0.1
                                                    v1 -= v3
                                                }
                                                _ = fi33
                                            }
                                        }
                                    case 32767:
                                        v39 := []int{(fi24 + fi24), fi24, (^(fi24)), (fi24 & (fi24 / ((-2147483648) | 1)))}
                                        _ = v39
                                    case -2147483648:
                                        for fi40 := 0; fi40 < 3; fi40++ {
                                            var v41 uint = 42
                                            v3 = v3
                                            fmt.Println(v11)
                                            _ = fi40
                                            _ = v41
                                        }
                                    default:
                                        var v42 uint16 = 255
                                        _ = v42
                                    }
                                    for fi43 := 0; fi43 < 20; fi43++ {
                                        if ((!true) && false) {
                                            v44 := []int{fi24}
                                            switch ((("foo bar" + "foo bar") + "\n\t") + "") {
                                            case "a":
                                                fmt.Println(v1)
                                            case "hello":
                                                var v45 int = (fi24 & fi24)
                                                v46 := GenS4{F0: 42, F1: 2, F2: 1}
                                                fmt.Println(v11, v1, fi43)
                                                _ = v45
                                                _ = v46
                                            }
                                            _ = v44
                                        }
                                        var v47 uint32 = 65535
                                        switch "a" {
                                        case "\n\t":
                                            fmt.Println(v11, fi43, fi24)
                                        case "foo bar":
                                            v47 = v47
                                            switch "foo bar" {
                                            case "\n\t":
                                                v11 = 1
                                            case "hello":
                                                v11 += (v11 - 0)
                                            default:
                                                var v48 int = -955
                                                _ = v48
                                            }
                                            var v49 int16 = 678
                                            v1 += 0.0
                                            if false {
                                                switch "" {
                                                case "hello":
                                                    fmt.Println(v11, fi24)
                                                case "":
                                                    v47 -= v47
                                                default:
                                                    if false {
                                                        switch (fi43 + fi24) {
                                                        case 32768:
                                                            v3 = 0.0
                                                        case -9223372036854775808:
                                                            fmt.Println(v11, v49)
                                                        case -2147483648:
                                                            var v50 int32 = v11
                                                            _ = v50
                                                        default:
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                switch (fi43 >> ((1) & 63)) {
                                                                case 417:
                                                                    if (true && true) {
                                                                        v51 := GenS6{F0: int64(v11), F1: 817, F2: (v49 % ((-447) | 1))}
                                                                        v49 = v49
                                                                        var v52 uint = 0
                                                                        _ = v51
                                                                        _ = v52
                                                                    }
                                                                case 2147483647:
                                                                    var v53 uint16 = 42
                                                                    _ = v53
                                                                case 300:
                                                                    v1 -= (v3 * 0.0)
                                                                }
                                                            case "\n\t":
                                                                fmt.Println(v49, v1, v49)
                                                            case "hello":
                                                                var v54 float32 = ((1e10 - (0.0 / 0.1)) / 1.0)
                                                                _ = v54
                                                            }
                                                        }
                                                        v55 := []int{(fi43 << ((876) & 63)), (-(fi24)), 556, (fi43 % ((4294967295) | 1))}
                                                        fmt.Println(v11, v1)
                                                        _ = v55
                                                    } else {
                                                        switch fi43 {
                                                        case 9223372036854775807:
                                                            v56 := GenS7{F0: (v47 >> ((55) & 63))}
                                                            _ = v56
                                                        case -158:
                                                            if (!true) {
                                                                var v57 uint32 = v47
                                                                var v58 int64 = -9223372036854775808
                                                                v59 := v2[(("" + "hello") + "a")]
                                                                _ = v57
                                                                _ = v58
                                                                _ = v59
                                                            }
                                                        case -393:
                                                            v49 += v49
                                                        default:
                                                            switch (fi24 << ((1) & 63)) {
                                                            case 906:
                                                                v3 = v3
                                                                var v60 string = ((("\n\t" + "a") + ("foo bar" + ("" + "a"))) + "\n\t")
                                                                _ = v60
                                                            default:
                                                                if ((fi43 - (-(fi43))) < (fi43 + (fi43 | 9223372036854775807))) {
                                                                    v11 = v11
                                                                    v11 = v11
                                                                } else {
                                                                    v11 = -2147483648
                                                                    v3 = float64((float32(-2.25) - 1.5))
                                                                    var v61 uint8 = 78
                                                                    _ = v61
                                                                }
                                                            }
                                                        }
                                                        if true {
                                                            var v62 int64 = 42
                                                            v63 := GenS8{F0: 547, F1: (v49 - 1000)}
                                                            _ = v62
                                                            _ = v63
                                                        }
                                                    }
                                                    v3 -= v3
                                                    switch 1 {
                                                    case -9223372036854775808:
                                                        if false {
                                                            if ((v49 << ((1) & 63)) > -32768) {
                                                                v2["foo bar"] = 4294967295
                                                                v64 := []int{9223372036854775807}
                                                                fmt.Println(v49)
                                                                v65 := GenS9{F0: (-(v11)), F1: (2 == 2)}
                                                                _ = v64
                                                                _ = v65
                                                            }
                                                            _ = len(v4)
                                                            var v66 uint16 = 463
                                                            if false {
                                                                v67 := []int{}
                                                                _ = v67
                                                            }
                                                            v68 := GenS4{F0: 1, F1: 255, F2: 5}
                                                            if true {
                                                                switch (fi43 >> ((1) & 63)) {
                                                                case -10:
                                                                    var v69 float32 = 3.14159
                                                                    _ = v69
                                                                case 599:
                                                                    var v70 uint64 = 4294967295
                                                                    _ = v70
                                                                default:
                                                                    if false {
                                                                        var v71 float64 = v1
                                                                        v72 := []int{(fi24 + fi43), fi43}
                                                                        var v73 int = (fi24 & fi43)
                                                                        _ = v71
                                                                        _ = v72
                                                                        _ = v73
                                                                    } else {
                                                                        v74 := GenS7{F0: (v47 & 2)}
                                                                        v75 := GenS9{F0: v11, F1: false}
                                                                        v76 := GenS9{F0: 2147483647, F1: true}
                                                                        v47 = 11
                                                                        _ = v74
                                                                        _ = v75
                                                                        _ = v76
                                                                    }
                                                                }
                                                                if (911 == (v66 % (((^(v66))) | 1))) {
                                                                    var v77 uint8 = 191
                                                                    v77 += v77
                                                                    _ = v77
                                                                } else {
                                                                    switch 0 {
                                                                    case 42:
                                                                        if (0 != -1) {
                                                                            var v78 uint32 = (v47 + v47)
                                                                            v79 := GenS2{F0: v47}
                                                                            if (0.0 != (0.0 - (float32(0.1) / 1.5))) {
                                                                                switch (fi24 - 4294967295) {
                                                                                case 9223372036854775807:
                                                                                    switch 0 {
                                                                                    case 32768:
                                                                                        if true {
                                                                                            fmt.Println(v3, v11, v11)
                                                                                            if (fi43 < 128) {
                                                                                                switch "hello" {
                                                                                                case "\n\t":
                                                                                                    var v80 int32 = (v11 * v11)
                                                                                                    var v81 uint = 4294967295
                                                                                                    _ = v80
                                                                                                    _ = v81
                                                                                                case "foo bar":
                                                                                                    fmt.Println(v3, fi24, v11)
                                                                                                    v82 := []int{fi43}
                                                                                                    v47 += v47
                                                                                                    var v83 uint8 = 1
                                                                                                    _ = v82
                                                                                                    _ = v83
                                                                                                default:
                                                                                                    v66 -= (^(v66))
                                                                                                }
                                                                                                v66 += v66
                                                                                            } else {
                                                                                                if (int(v49) <= (fi43 >> ((0) & 63))) {
                                                                                                    v84 := GenS10{F0: 81, F1: v66, F2: (v49 - int16(v47))}
                                                                                                    if true {
                                                                                                        switch fi24 {
                                                                                                        case 32767:
                                                                                                            v85 := []int{(fi43 & (fi24 | (^(fi24)))), (fi43 % ((fi24) | 1)), 42}
                                                                                                            _ = v85
                                                                                                        }
                                                                                                        if true {
                                                                                                            v86 := GenS2{F0: 2147483647}
                                                                                                            _ = v86
                                                                                                        } else {
                                                                                                            switch fi24 {
                                                                                                            case 32768:
                                                                                                                var v87 float64 = 1e10
                                                                                                                _ = v87
                                                                                                            case 128:
                                                                                                                _ = len(v4)
                                                                                                            default:
                                                                                                                if (v78 >= v78) {
                                                                                                                    v88 := GenS2{F0: v78}
                                                                                                                    v11 -= v11
                                                                                                                    v3 = 0.1
                                                                                                                    var v89 uint8 = 255
                                                                                                                    v90 := v2["a"]
                                                                                                                    _ = v88
                                                                                                                    _ = v89
                                                                                                                    _ = v90
                                                                                                                }
                                                                                                            }
                                                                                                            v1 = -2.25
                                                                                                        }
                                                                                                        fmt.Println(v66, v78)
                                                                                                    } else {
                                                                                                        if true {
                                                                                                            v47 += (^(v78))
                                                                                                            v47 -= v47
                                                                                                        }
                                                                                                    }
                                                                                                    switch (-(fi24)) {
                                                                                                    case 1:
                                                                                                        var v91 int = (^(fi43))
                                                                                                        _ = v91
                                                                                                    case -217:
                                                                                                        v47 += (^(v47))
                                                                                                    }
                                                                                                    if (true || (((v66 - v66) < v66) && (v49 >= v49))) {
                                                                                                        switch "" {
                                                                                                        case "foo bar":
                                                                                                            if ((float32(v1) + (-2.25 - float32(3.14159))) != 1e10) {
                                                                                                                v92 := []int{}
                                                                                                                fmt.Println(fi43)
                                                                                                                v49 += (^(v49))
                                                                                                                var v93 int8 = -128
                                                                                                                if len(v92) > 0 {
                                                                                                                    v92[0] = fi43
                                                                                                                }
                                                                                                                _ = v92
                                                                                                                _ = v93
                                                                                                            }
                                                                                                            v94 := GenS11{F0: fi43}
                                                                                                            _ = v94
                                                                                                        case "\n\t":
                                                                                                            var v95 uint64 = 1
                                                                                                            _ = v95
                                                                                                        default:
                                                                                                            v66 = 42
                                                                                                        }
                                                                                                    } else {
                                                                                                        if (uint(v66) != 218) {
                                                                                                            if (26 > 42) {
                                                                                                                var v96 string = "foo bar"
                                                                                                                if (false && (false || false)) {
                                                                                                                    fmt.Println(v66, fi24, v3)
                                                                                                                    v97 := GenS2{F0: v47}
                                                                                                                    v98 := []int{(-(fi43)), 321, 2147483647, fi43}
                                                                                                                    _ = v97
                                                                                                                    _ = v98
                                                                                                                } else {
                                                                                                                    v99 := v2["\n\t"]
                                                                                                                    _ = v99
                                                                                                                }
                                                                                                                var v100 uint8 = 106
                                                                                                                _ = v100
                                                                                                                _ = v96
                                                                                                            }
                                                                                                            fmt.Println(v78, v49)
                                                                                                        }
                                                                                                        v66 = v66
                                                                                                        v11 = (v11 % ((v11) | 1))
                                                                                                        if ((0.1 + 1.5) != 3.14159) {
                                                                                                            v2[("foo bar" + "")] = fi24
                                                                                                            if false {
                                                                                                                v101 := GenS12{F0: 65535, F1: 2, F2: 0}
                                                                                                                _ = v101
                                                                                                            }
                                                                                                            var v102 bool = true
                                                                                                            _ = v102
                                                                                                        }
                                                                                                    }
                                                                                                    var v103 uint64 = 925
                                                                                                    var v104 uint = 0
                                                                                                    v104 = 293
                                                                                                    _ = v103
                                                                                                    _ = v104
                                                                                                    _ = v84
                                                                                                } else {
                                                                                                    v2["\n\t"] = fi43
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            v78 = v47
                                                                                            if true {
                                                                                                if (v11 <= v11) {
                                                                                                    var v105 int16 = v49
                                                                                                    var v106 int64 = 0
                                                                                                    v78 = (^(v47))
                                                                                                    var v107 int = 452
                                                                                                    var v108 int16 = (v49 % ((int16(v106)) | 1))
                                                                                                    fmt.Println(v49, fi43, fi43)
                                                                                                    v109 := GenS11{F0: (fi24 & -2147483648)}
                                                                                                    v110 := GenS13{F0: 65535, F1: v108, F2: int(v78)}
                                                                                                    var v111 int32 = int32(fi24)
                                                                                                    fmt.Println(v78)
                                                                                                    if false {
                                                                                                        var v112 uint = 0
                                                                                                        v113 := v4[(v78 & (v47 ^ uint32(v66)))]
                                                                                                        _ = v112
                                                                                                        _ = v113
                                                                                                    }
                                                                                                    _ = v105
                                                                                                    _ = v106
                                                                                                    _ = v107
                                                                                                    _ = v108
                                                                                                    _ = v109
                                                                                                    _ = v110
                                                                                                    _ = v111
                                                                                                } else {
                                                                                                    var v114 float64 = 3.14159
                                                                                                    v115 := []int{fi24}
                                                                                                    _ = v114
                                                                                                    _ = v115
                                                                                                }
                                                                                                var v116 int8 = 0
                                                                                                _ = v116
                                                                                            } else {
                                                                                                var v117 uint64 = 42
                                                                                                var v118 int = fi24
                                                                                                v49 += v49
                                                                                                _ = v117
                                                                                                _ = v118
                                                                                            }
                                                                                            v1 = (1.5 - ((v3 - v3) - (float64(1e10) * (v3 / float64((3.14159 - 1.5))))))
                                                                                        }
                                                                                    default:
                                                                                        fmt.Println(v47, fi43, v78)
                                                                                    }
                                                                                case 0:
                                                                                    v78 -= (^(v78))
                                                                                default:
                                                                                    if true {
                                                                                        if (!((3.14159 < 1e10) && false)) {
                                                                                            v66 = 0
                                                                                            v119 := v4[683]
                                                                                            _ = v119
                                                                                        }
                                                                                    }
                                                                                }
                                                                                fmt.Println(v78, fi24)
                                                                                if (!((true && (255 >= 997)) && false)) {
                                                                                    switch -909 {
                                                                                    case 32767:
                                                                                        v11 = (-(v11))
                                                                                    case 127:
                                                                                        var v120 uint = 2
                                                                                        _ = v120
                                                                                    default:
                                                                                        var v121 string = ((("\n\t" + "") + "") + "")
                                                                                        _ = v121
                                                                                    }
                                                                                    if ((((1.5 / -2.25) > 1e10) && true) && (-1 > -919)) {
                                                                                        v122 := GenS7{F0: v78}
                                                                                        v47 -= (v47 & v47)
                                                                                        _ = v122
                                                                                    } else {
                                                                                        v123 := []int{-80, (-(fi24))}
                                                                                        var v124 int32 = (v11 ^ (v11 + v11))
                                                                                        _ = v123
                                                                                        _ = v124
                                                                                    }
                                                                                    switch 32767 {
                                                                                    case -2147483648:
                                                                                        v125 := GenS13{F0: (v66 | 898), F1: 32767, F2: (fi43 - (fi43 >> ((65535) & 63)))}
                                                                                        _ = v125
                                                                                    default:
                                                                                        if (!(!(1 >= 127))) {
                                                                                            switch "hello" {
                                                                                            case "foo bar":
                                                                                                if true {
                                                                                                    v126 := GenS10{F0: -128, F1: 203, F2: v49}
                                                                                                    _ = v126
                                                                                                }
                                                                                            default:
                                                                                                v66 = 255
                                                                                            }
                                                                                            var v127 uint8 = 83
                                                                                            v128 := GenS9{F0: v11, F1: false}
                                                                                            _ = v127
                                                                                            _ = v128
                                                                                        } else {
                                                                                            v129 := GenS4{F0: 9223372036854775807, F1: 127, F2: 119}
                                                                                            _ = v129
                                                                                        }
                                                                                    }
                                                                                    var v130 bool = true
                                                                                    v131 := GenS8{F0: 255, F1: (v49 >> ((255) & 63))}
                                                                                    _ = v130
                                                                                    _ = v131
                                                                                } else {
                                                                                    fmt.Println(v1)
                                                                                }
                                                                            }
                                                                            _ = v78
                                                                            _ = v79
                                                                        } else {
                                                                            v132 := GenS9{F0: 42, F1: false}
                                                                            _ = v132
                                                                        }
                                                                        var v133 float64 = (v3 + v3)
                                                                        if (((^(v47)) <= v47) || ((v11 - (v11 / ((v11) | 1))) != 0)) {
                                                                            if true {
                                                                                v133 = (3.14159 + float64(3.14159))
                                                                                if true {
                                                                                    switch -1 {
                                                                                    case 127:
                                                                                        var v134 uint32 = (v47 + (v47 - (^(v47))))
                                                                                        _ = v134
                                                                                    default:
                                                                                        var v135 uint8 = 2
                                                                                        _ = v135
                                                                                    }
                                                                                } else {
                                                                                    var v136 uint = 115
                                                                                    v66 -= v66
                                                                                    if false {
                                                                                        var v137 int32 = 2147483647
                                                                                        v1 -= v3
                                                                                        _ = v137
                                                                                    } else {
                                                                                        var v138 uint8 = 1
                                                                                        _ = v138
                                                                                    }
                                                                                    _ = v136
                                                                                }
                                                                            } else {
                                                                                var v139 float64 = 1.5
                                                                                var v140 bool = true
                                                                                var v141 int = (fi43 - (^(fi43)))
                                                                                _ = v139
                                                                                _ = v140
                                                                                _ = v141
                                                                            }
                                                                            var v142 uint8 = 146
                                                                            _ = v142
                                                                        } else {
                                                                            v11 = v11
                                                                            v143 := GenS4{F0: 1, F1: 127, F2: -73}
                                                                            _ = v143
                                                                        }
                                                                        _ = v133
                                                                    case -9223372036854775808:
                                                                        v3 -= v3
                                                                    default:
                                                                        var v144 float32 = ((3.14159 * 0.1) - (float32((float64(((((0.0 + ((float32(v1) / ((3.14159 - 0.0) - 1e10)) / -2.25)) / -2.25) - -2.25) / -2.25)) / (1e10 / (-2.25 * (float64(-2.25) / (v3 + float64((((float32((v1 / 1e10)) - float32((1.5 / v3))) - 1e10) + 0.0)))))))) * 0.0))
                                                                        _ = v144
                                                                    }
                                                                    fmt.Println(v66, v11)
                                                                    v66 = v66
                                                                    v2[("foo bar" + "")] = 9223372036854775807
                                                                }
                                                            }
                                                            _ = v66
                                                            _ = v68
                                                        }
                                                        var v145 int8 = genFn0()
                                                        var v146 uint = 0
                                                        _ = v145
                                                        _ = v146
                                                    case 562:
                                                        v11 += (v11 & -1)
                                                    default:
                                                        switch ("a" + ("" + ("" + "hello"))) {
                                                        case "hello":
                                                            v147 := GenS1{F0: 42, F1: (("a" + "\n\t") + ""), F2: ""}
                                                            _ = v147
                                                        case "foo bar":
                                                            if true {
                                                                var v148 bool = true
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        _ = len(v2)
                                                                    case "foo bar":
                                                                        v149 := GenS9{F0: -2147483648, F1: v148}
                                                                        _ = v149
                                                                    default:
                                                                        v150 := GenS5{F0: v47, F1: -1, F2: 19}
                                                                        _ = v150
                                                                    }
                                                                case "hello":
                                                                    v151 := []int{fi43, (-(fi43))}
                                                                    _ = v151
                                                                case "\n\t":
                                                                    _ = len(v2)
                                                                }
                                                                v152 := []int{0, 32768, (fi24 / (((fi24 / (((^(fi24))) | 1))) | 1)), fi43}
                                                                _ = v148
                                                                _ = v152
                                                            } else {
                                                                switch 32768 {
                                                                case 32767:
                                                                    v153 := v4[(v47 >> ((255) & 63))]
                                                                    _ = v153
                                                                case 1:
                                                                    var v154 uint8 = 216
                                                                    _ = v154
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                                v1 = v1
                                            } else {
                                                var v155 int = (fi24 % ((int(v49)) | 1))
                                                var v156 uint = 255
                                                _ = v155
                                                _ = v156
                                            }
                                            _ = v49
                                        case "a":
                                            var v157 int64 = 1
                                            _ = v157
                                        }
                                        _ = fi43
                                        _ = v47
                                    }
                                    for fi158 := 0; fi158 < 17; fi158++ {
                                        var v159 uint64 = 687
                                        fmt.Println(v159, fi24)
                                        fmt.Println(fi24, v3, fi24)
                                        _ = fi158
                                        _ = v159
                                    }
                                }
                            case -1:
                                for fi160 := 0; fi160 < 11; fi160++ {
                                    if (!true) {
                                        var v161 float32 = (3.14159 / ((0.1 - 3.14159) - 0.0))
                                        if true {
                                            var v162 int64 = -744
                                            _ = v162
                                        } else {
                                            switch fi24 {
                                            case 0:
                                                v3 = (v3 * (float64(0.0) / 1e10))
                                            case 4294967295:
                                                fmt.Println(v3, fi24, v161)
                                                v163 := []int{(-(fi160)), (fi24 * 128), (fi160 << ((727) & 63))}
                                                v164 := GenS8{F0: 255, F1: 42}
                                                _ = v163
                                                _ = v164
                                            default:
                                                if (4294967295 >= 639) {
                                                    v165 := GenS1{F0: 549, F1: "\n\t", F2: "\n\t"}
                                                    v11 += (v11 & 1)
                                                    var v166 string = "foo bar"
                                                    fmt.Println(fi24, v11)
                                                    fmt.Println(fi160)
                                                    fmt.Println(fi24, v166, fi160)
                                                    v1 += (((0.1 + v3) - v3) + v3)
                                                    if false {
                                                        if (false || (-9223372036854775808 <= -911)) {
                                                            var v167 int = fi24
                                                            var v168 uint16 = 708
                                                            _ = v167
                                                            _ = v168
                                                        }
                                                        v169 := []int{fi160}
                                                        _ = v169
                                                    } else {
                                                        var v170 float64 = (v3 / (((-2.25 * (-2.25 + v1)) - 1.5) * (0.0 - 1e10)))
                                                        if false {
                                                            if false {
                                                                v1 = 3.14159
                                                                v171 := GenS2{F0: 897}
                                                                _ = v171
                                                            } else {
                                                                v172 := GenS14{F0: 70, F1: -1}
                                                                v173 := GenS14{F0: -57, F1: 236}
                                                                _ = v172
                                                                _ = v173
                                                            }
                                                        }
                                                        var v174 uint32 = 2147483648
                                                        _ = v170
                                                        _ = v174
                                                    }
                                                    v3 += -2.25
                                                    _ = v165
                                                    _ = v166
                                                } else {
                                                    var v175 bool = true
                                                    if v175 {
                                                        var v176 uint8 = 127
                                                        _ = v176
                                                    } else {
                                                        v11 = -27
                                                    }
                                                    v161 += v161
                                                    _ = v175
                                                }
                                            }
                                        }
                                        _ = v161
                                    } else {
                                        if ((v11 ^ v11) >= v11) {
                                            var v177 int = 32767
                                            var v178 string = "hello"
                                            fmt.Println(v178)
                                            if (!true) {
                                                var v179 uint64 = 2
                                                fmt.Println(v177, v178, fi24)
                                                if (uint16(v179) != 198) {
                                                    v11 = v11
                                                    if false {
                                                        v4[2147483647] = true
                                                        var v180 uint64 = 992
                                                        var v181 int = 0
                                                        v182 := GenS8{F0: 65535, F1: -1}
                                                        _ = v180
                                                        _ = v181
                                                        _ = v182
                                                    } else {
                                                        switch (fi24 & (v177 - (^(v177)))) {
                                                        case -11:
                                                            v183 := GenS15{F0: 255, F1: fi160, F2: 2}
                                                            _ = v183
                                                        default:
                                                            v3 = float64(-2.25)
                                                        }
                                                        v184 := GenS16{F0: -9223372036854775808}
                                                        _ = v184
                                                    }
                                                    if (1 > 1) {
                                                        v185 := GenS15{F0: 833, F1: (fi160 + (-(v177))), F2: 302}
                                                        var v186 int16 = -289
                                                        v187 := GenS11{F0: v177}
                                                        v178 = v178
                                                        _ = v185
                                                        _ = v186
                                                        _ = v187
                                                    } else {
                                                        if true {
                                                            if (!((false || true) && false)) {
                                                                v188 := GenS3{F0: (^(fi160)), F1: 75, F2: -128}
                                                                delete(v2, (v178 + (("foo bar" + ((v178 + (v178 + (v178 + v178))) + v178)) + "\n\t")))
                                                                var v189 float32 = (-2.25 + 0.1)
                                                                _ = v188
                                                                _ = v189
                                                            }
                                                            if (1 >= 1) {
                                                                v11 = 651
                                                            } else {
                                                                var v190 float64 = -2.25
                                                                fmt.Println(v177, v11, v1)
                                                                var v191 int8 = 0
                                                                _ = v190
                                                                _ = v191
                                                            }
                                                        }
                                                        v192 := []int{int(v179)}
                                                        v177 = fi160
                                                        _ = v192
                                                    }
                                                }
                                                var v193 uint64 = 4294967295
                                                if true {
                                                    v11 += v11
                                                }
                                                _ = v179
                                                _ = v193
                                            }
                                            v3 = v1
                                            v194 := []int{}
                                            _ = v177
                                            _ = v178
                                            _ = v194
                                        }
                                    }
                                    _ = fi160
                                }
                            default:
                                var v195 uint = 0
                                _ = v195
                            }
                        case "":
                            switch "hello" {
                            case "hello":
                                var v196 float64 = (v3 - 1.5)
                                _ = v196
                            }
                        }
                        _ = len(v4)
                        if false {
                            v197 := GenS17{F0: 63, F1: 255}
                            v11 -= v11
                            v3 = (((1e10 - 0.0) * 0.0) + (1e10 - (v1 + (3.14159 + (0.0 / 0.1)))))
                            _ = v197
                        }
                        _ = fi24
                    }
                    v198 := GenS9{F0: v11, F1: (float64(1.5) != v1)}
                    for fi199 := 0; fi199 < 12; fi199++ {
                        v200 := GenS18{F0: (664 != 42)}
                        _ = fi199
                        _ = v200
                    }
                    _ = v198
                }
            case 438:
                v201 := GenS16{F0: -9223372036854775808}
                _ = v201
            default:
                v3 = (v3 * v3)
            }
        }
    }
    fmt.Println(v11)
    if true {
        switch 0 {
        case 127:
            switch ("hello" + "hello") {
            case "foo bar":
                if false {
                    v202 := GenS19{F0: 1000}
                    var v203 uint8 = 255
                    _ = v202
                    _ = v203
                }
            case "":
                v204 := []int{}
                _ = v204
            }
        default:
            var v205 int = 1
            _ = v205
        }
        for fi206 := 0; fi206 < 15; fi206++ {
            if true {
                fmt.Println(fi206, v3)
                v207 := []int{(fi206 ^ -2147483648)}
                v207 = append(v207, 127)
                _ = v207
            }
            var v208 int32 = 42
            for fi209 := 0; fi209 < 19; fi209++ {
                if (1 == 255) {
                    v3 = v3
                    if (0 >= 2) {
                        v210 := GenS20{F0: 117}
                        var v211 float64 = ((v1 + v1) * 0.0)
                        _ = v210
                        _ = v211
                    } else {
                        if ("" <= "a") {
                            v11 += (-(v208))
                            if false {
                                var v212 int = -1
                                _ = len(v4)
                                if true {
                                    var v213 int = fi206
                                    if (false || true) {
                                        fmt.Println(v1)
                                        v4[2147483647] = (true && true)
                                    } else {
                                        v214 := []int{fi209}
                                        v208 = v208
                                        _ = v214
                                    }
                                    _ = v213
                                }
                                _ = v212
                            }
                            var v215 string = ""
                            v215 = "hello"
                            _ = v215
                        } else {
                            v216 := GenS3{F0: fi209, F1: 0, F2: 28}
                            switch "\n\t" {
                            case "hello":
                                var v217 uint64 = 140
                                _ = v217
                            case "foo bar":
                                v11 -= (^(v208))
                            }
                            var v218 float64 = v1
                            _ = v216
                            _ = v218
                        }
                        v208 -= v11
                    }
                    var v219 uint32 = 436
                    v220 := GenS21{F0: 1, F1: 0, F2: 32767}
                    _ = v219
                    _ = v220
                } else {
                    if true {
                        var v221 string = (((("a" + "\n\t") + "") + "\n\t") + "\n\t")
                        v1 = float64((float32(((float64(((1.5 / -2.25) * -2.25)) + ((0.1 / v3) - (float64(-2.25) / (1e10 + v3)))) + v1)) / 1.5))
                        var v222 uint32 = 1
                        _ = v221
                        _ = v222
                    }
                    switch (fi206 >> ((509) & 63)) {
                    case 42:
                        var v223 int16 = 42
                        _ = v223
                    }
                    switch (fi206 | (fi206 & -23)) {
                    case 127:
                        v224 := []int{42, fi206}
                        _ = v224
                    case 0:
                        var v225 uint32 = 565
                        _ = v225
                    case -2147483648:
                        v226 := GenS22{F0: 0}
                        _ = v226
                    default:
                        v227 := GenS10{F0: genFn0(), F1: 32767, F2: -32768}
                        _ = v227
                    }
                }
                _ = fi209
            }
            _ = fi206
            _ = v208
        }
    }
    for fi228 := 0; fi228 < 13; fi228++ {
        v3 += 0.1
        _ = fi228
    }
    if true {
        if (!true) {
            var v229 uint = 890
            _ = v229
        }
    }
    if (false || (-887 == 559)) {
        switch 32767 {
        case -689:
            for fi230 := 0; fi230 < 1; fi230++ {
                if false {
                    if (0.1 != (3.14159 / 1.0)) {
                        v231 := GenS3{F0: fi230, F1: genFn0(), F2: 1}
                        var v232 int16 = -32768
                        v11 -= v11
                        _ = v231
                        _ = v232
                    }
                    v233 := []int{(fi230 % (((fi230 + 9223372036854775807)) | 1))}
                    _ = v233
                }
                var v234 string = ("foo bar" + ("hello" + "foo bar"))
                _ = fi230
                _ = v234
            }
        case 173:
            if (true && false) {
                v235 := GenS23{F0: 703}
                for fi236 := 0; fi236 < 9; fi236++ {
                    v237 := GenS6{F0: 0, F1: 596, F2: 0}
                    _ = fi236
                    _ = v237
                }
                _ = v235
            }
        }
        for fi238 := 0; fi238 < 6; fi238++ {
            v1 -= 1e10
            for fi239 := 0; fi239 < 17; fi239++ {
                if false {
                    if (int64(v11) > -1) {
                        v1 += -2.25
                        switch fi239 {
                        case 1:
                            var v240 int32 = v11
                            _ = v240
                        case 224:
                            v3 = 0.0
                        default:
                            if (false && false) {
                                if (1e10 <= v3) {
                                    v241 := GenS24{F0: 100, F1: 180, F2: 1}
                                    v242 := []int{(fi239 ^ (fi239 >> ((42) & 63))), (fi239 & fi239)}
                                    switch (((("a" + "hello") + ((("foo bar" + ("hello" + "hello")) + "foo bar") + ("hello" + "a"))) + "hello") + "hello") {
                                    case "a":
                                        v11 -= 0
                                    default:
                                        v3 = v3
                                    }
                                    v243 := GenS19{F0: -32768}
                                    fmt.Println(fi238, fi238, v1)
                                    v1 -= ((v1 / 3.14159) * ((1.5 / ((v3 * v3) + (v1 + 1.5))) + 1e10))
                                    if len(v242) > 0 {
                                        v242[0] = fi238
                                    }
                                    _ = v241
                                    _ = v242
                                    _ = v243
                                } else {
                                    var v244 uint = 65535
                                    _ = v244
                                }
                                var v245 bool = (0.0 != (1.5 + 0.1))
                                var v246 string = ("a" + "foo bar")
                                if (!false) {
                                    v247 := GenS25{F0: 4294967295}
                                    v11 = v11
                                    _ = v247
                                }
                                _ = v245
                                _ = v246
                            }
                        }
                        v248 := GenS26{F0: v11, F1: (fi238 - -97), F2: 42}
                        _ = v248
                    } else {
                        var v249 int32 = (-(v11))
                        v250 := []int{(fi239 & (-(fi238))), (fi238 << ((371) & 63)), fi239}
                        fmt.Println(v3, v3)
                        _ = v249
                        _ = v250
                    }
                }
                v3 += 3.14159
                _ = fi239
            }
            switch (-(fi238)) {
            case -265:
                fmt.Println(v11, v3)
            case -438:
                v11 = v11
            case -9223372036854775808:
                if true {
                    v251 := GenS27{F0: (fi238 - fi238), F1: "foo bar"}
                    _ = v251
                }
            }
            _ = fi238
        }
        v2[("foo bar" + "hello")] = -2147483648
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
