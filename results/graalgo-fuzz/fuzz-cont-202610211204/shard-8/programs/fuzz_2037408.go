package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 uint; F1 int; F2 uint16 }
type GenS5 struct { F0 int8; F1 uint64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 bool; F1 int64 }
type GenS9 struct { F0 int8; F1 int64 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint64; F1 int8 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 int8; F1 uint8 }
type GenS14 struct { F0 int; F1 uint32; F2 int64 }
type GenS15 struct { F0 int8; F1 int64; F2 uint64 }
type GenS16 struct { F0 bool; F1 int32 }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 int32 }
func genFn0(p0 string, p1 string) int64 {
    fmt.Println(p1)
    var v6 float32 = -2.25
    for fi7 := 0; fi7 < 7; fi7++ {
        if false {
            p0 = "\n\t"
            var v8 string = p0
            p0 = p1
            _ = v8
        }
        v6 -= v6
        v9 := GenS1{F0: 128}
        _ = fi7
        _ = v9
    }
    if true {
        p1 = p1
        if false {
            for fi10 := 0; fi10 < 3; fi10++ {
                var v11 uint = 0
                _ = fi10
                _ = v11
            }
            v12 := []int{9223372036854775807, -348, 32767}
            var v13 string = p0
            _ = v12
            _ = v13
        } else {
            var v14 uint8 = 255
            for fi15 := 0; fi15 < 1; fi15++ {
                if false {
                    if (!false) {
                        if false {
                            v16 := GenS2{F0: 1, F1: -2147483648}
                            var v17 uint = 65535
                            if (true && (fi15 < (fi15 - (^(fi15))))) {
                                if (((p1 + p1) + "hello") == "\n\t") {
                                    var v18 uint64 = 4294967295
                                    fmt.Println(v18, v14, v6)
                                    v19 := []int{fi15, fi15, fi15, fi15}
                                    _ = v18
                                    _ = v19
                                } else {
                                    v6 -= 3.14159
                                    v20 := []int{fi15, fi15, (fi15 * int(v14)), (fi15 * fi15)}
                                    _ = v20
                                }
                                var v21 uint32 = 600
                                _ = v21
                            }
                            _ = v16
                            _ = v17
                        } else {
                            v6 -= v6
                        }
                    }
                    p0 = ("\n\t" + "")
                    for fi22 := 0; fi22 < 9; fi22++ {
                        if false {
                            fmt.Println(p0, p1)
                            var v23 float64 = float64(v6)
                            v24 := []int{}
                            _ = v23
                            _ = v24
                        }
                        _ = fi22
                    }
                } else {
                    if ((!(1 >= -1)) || false) {
                        v25 := []int{(fi15 ^ (fi15 * (fi15 / ((1) | 1)))), fi15, (fi15 >> ((1) & 63))}
                        _ = v25
                    }
                    p1 = p0
                }
                var v26 bool = ((65535 >= 899) && (!((!(105 > 1)) && (!true))))
                _ = fi15
                _ = v26
            }
            _ = v14
        }
    } else {
        p1 = ("a" + p0)
        p1 = ((((p1 + (p1 + ((p1 + p1) + (p0 + p0)))) + p0) + ("foo bar" + p1)) + (("foo bar" + p1) + "a"))
    }
    _ = p0
    _ = p1
    _ = v6
    return -9223372036854775808
}

func main() {
    var v1 int = 42
    var v2 uint64 = 103
    var v3 bool = false
    var v4 uint16 = 354
    var v5 uint8 = 214
    for fi27 := 0; fi27 < 5; fi27++ {
        if (306 == -422) {
            fmt.Println(v4, v5, fi27)
            if (100 >= 42) {
                v2 += (v2 * (v2 << ((852) & 63)))
                var v28 uint16 = 748
                switch "hello" {
                case "":
                    for fi29 := 0; fi29 < 1; fi29++ {
                        v30 := GenS1{F0: (v5 | (v5 & 226))}
                        _ = fi29
                        _ = v30
                    }
                case "hello":
                    if v3 {
                        var v31 bool = v3
                        if (599 < 255) {
                            v32 := GenS3{F0: "\n\t", F1: -236}
                            _ = v32
                        }
                        _ = v31
                    } else {
                        v33 := []int{(fi27 * -824), fi27}
                        for fi34 := 0; fi34 < 16; fi34++ {
                            if false {
                                _ = len(v33)
                                fmt.Println(v1, v2, v3)
                                v35 := GenS4{F0: 0, F1: 2147483647, F2: (v28 >> ((991) & 63))}
                                _ = v35
                            } else {
                                var v36 uint8 = v5
                                var v37 uint32 = 554
                                _ = v36
                                _ = v37
                            }
                            var v38 float32 = -2.25
                            v39 := GenS4{F0: 589, F1: (v1 * 4294967295), F2: (v28 * v28)}
                            v1 = (fi34 - v1)
                            _ = fi34
                            _ = v38
                            _ = v39
                        }
                        _ = v33
                    }
                }
                _ = v28
            } else {
                fmt.Println(v5, fi27, v2)
                v40 := GenS5{F0: -94, F1: v2}
                var v41 uint64 = 18446744073709551615
                _ = v40
                _ = v41
            }
        }
        _ = fi27
    }
    var v42 int16 = 0
    if v3 {
        var v43 int = v1
        var v44 int16 = v42
        var v45 uint64 = (^(v2))
        _ = v43
        _ = v44
        _ = v45
    } else {
        v46 := []int{v1, v1}
        _ = v46
    }
    v3 = v3
    v47 := []int{v1, (v1 | v1)}
    if (!v3) {
        v47 = append(v47, 1)
        for fi48 := 0; fi48 < 19; fi48++ {
            if len(v47) > 0 {
                v47[0] = (fi48 / ((fi48) | 1))
            }
            if (!v3) {
                for fi49 := 0; fi49 < 20; fi49++ {
                    if false {
                        switch "" {
                        case "":
                            switch 632 {
                            case 851:
                                v42 += v42
                            case 0:
                                var v50 uint32 = 2147483647
                                _ = v50
                            }
                        case "foo bar":
                            v1 += 127
                        }
                        if len(v47) >= 2 {
                            v47 = v47[1:len(v47)]
                        }
                    } else {
                        var v51 int = (fi49 + (fi48 / (((fi49 & (fi48 * 2147483647))) | 1)))
                        v4 += v4
                        _ = v51
                    }
                    _ = fi49
                }
            } else {
                switch (fi48 / (((v1 ^ (v1 << ((42) & 63)))) | 1)) {
                case -1:
                    var v52 uint32 = 65535
                    _ = v52
                case 32768:
                    v4 += 0
                case -510:
                    if (v3 && (v4 == v4)) {
                        for fi53 := 0; fi53 < 20; fi53++ {
                            if (-419 < -32768) {
                                v4 = v4
                                v42 += (v42 | -378)
                                var v54 float32 = 1.5
                                _ = v54
                            }
                            v42 = (v42 * v42)
                            var v55 int = (^(fi48))
                            _ = fi53
                            _ = v55
                        }
                        for fi56 := 0; fi56 < 12; fi56++ {
                            if len(v47) >= 2 {
                                v47 = v47[1:len(v47)]
                            }
                            _ = fi56
                        }
                    } else {
                        if len(v47) >= 2 {
                            v47 = v47[1:len(v47)]
                        }
                        if len(v47) > 0 {
                            v47[0] = (fi48 / ((985) | 1))
                        }
                    }
                default:
                    for fi57 := 0; fi57 < 13; fi57++ {
                        v58 := GenS2{F0: 1, F1: -2147483648}
                        v3 = (!v3)
                        switch "\n\t" {
                        case "hello":
                            if false {
                                fmt.Println(v1)
                            }
                        case "\n\t":
                            v4 = 1
                        case "foo bar":
                            v59 := GenS6{F0: (^(v2))}
                            _ = v59
                        }
                        _ = fi57
                        _ = v58
                    }
                }
            }
            var v60 int32 = 933
            _ = fi48
            _ = v60
        }
        if false {
            var v61 float64 = (((0.0 - 0.0) * float64(0.0)) * 0.0)
            v62 := GenS3{F0: "a", F1: -2147483648}
            if len(v47) >= 2 {
                v47 = v47[1:len(v47)]
            }
            _ = v61
            _ = v62
        }
    } else {
        for fi63 := 0; fi63 < 19; fi63++ {
            if len(v47) >= 2 {
                v47 = v47[1:len(v47)]
            }
            _ = fi63
        }
        var v64 float32 = 3.14159
        v64 += 1.5
        _ = v64
    }
    var v65 int = v1
    switch "\n\t" {
    case "foo bar":
        v66 := GenS7{F0: v5}
        _ = v66
    case "a":
        if (v4 <= (v4 << ((42) & 63))) {
            var v67 string = "foo bar"
            v3 = (42 >= 1)
            _ = v67
        } else {
            _ = len(v47)
        }
    }
    var v68 uint64 = (v2 % ((v2) | 1))
    var v69 uint64 = (^(v68))
    if (1 == (^(v68))) {
        if v3 {
            v70 := GenS8{F0: ("\n\t" < "hello"), F1: -1}
            _ = v70
        }
        v71 := GenS1{F0: v5}
        _ = v71
    }
    v65 += (v65 << ((453) & 63))
    for fi72 := 0; fi72 < 5; fi72++ {
        v73 := GenS8{F0: (("foo bar" + "\n\t") >= ""), F1: -1}
        var v74 uint64 = v69
        var v75 uint32 = 65535
        _ = fi72
        _ = v73
        _ = v74
        _ = v75
    }
    var v76 int64 = -862
    switch ("a" + ("foo bar" + "foo bar")) {
    case "foo bar":
        if true {
            v1 = (v1 % ((525) | 1))
            for fi77 := 0; fi77 < 14; fi77++ {
                var v78 int = v1
                _ = fi77
                _ = v78
            }
        } else {
            v79 := GenS2{F0: v68, F1: 2147483647}
            _ = v79
        }
    case "":
        if ((v2 + (v68 ^ v68)) < v68) {
            fmt.Println(v42)
            fmt.Println(v69, v3)
            if true {
                var v80 uint = 2
                if (v3 || false) {
                    var v81 int8 = -128
                    switch ("a" + ("foo bar" + "a")) {
                    case "a":
                        if v3 {
                            for fi82 := 0; fi82 < 8; fi82++ {
                                if ((((-2.25 * (1.5 - (0.1 / 1.5))) * 3.14159) == float32(-2.25)) || (v68 > 2)) {
                                    if ((v76 - v76) != v76) {
                                        if (!(!false)) {
                                            fmt.Println(v5)
                                            var v83 string = ""
                                            _ = v83
                                        } else {
                                            if len(v47) > 0 {
                                                v47[0] = (v65 / ((-1) | 1))
                                            }
                                            if len(v47) >= 2 {
                                                v47 = v47[1:len(v47)]
                                            }
                                        }
                                        switch (v65 >> ((v80) & 63)) {
                                        case 207:
                                            _ = len(v47)
                                        default:
                                            var v84 float64 = (3.14159 / (float64(((1.5 / 1.0) * float32(3.14159))) + (1.5 + (0.0 * 1e10))))
                                            _ = v84
                                        }
                                        var v85 int64 = v76
                                        _ = v85
                                    } else {
                                        v86 := GenS4{F0: 2, F1: v65, F2: (v4 | v4)}
                                        v87 := GenS9{F0: v81, F1: 42}
                                        var v88 uint16 = (v4 + v4)
                                        _ = v86
                                        _ = v87
                                        _ = v88
                                    }
                                    v1 += 0
                                    if (v3 && ("foo bar" != "a")) {
                                        v47 = append(v47, (v65 << ((255) & 63)))
                                        var v89 int8 = v81
                                        if v3 {
                                            switch (v65 >> ((0) & 63)) {
                                            case 1:
                                                for fi90 := 0; fi90 < 11; fi90++ {
                                                    v1 -= (fi82 ^ -2147483648)
                                                    v65 -= int(v81)
                                                    _ = fi90
                                                }
                                            case 9223372036854775807:
                                                var v91 int64 = v76
                                                _ = v91
                                            case -711:
                                                for fi92 := 0; fi92 < 11; fi92++ {
                                                    fmt.Println(fi92, v1, v76)
                                                    v93 := GenS10{F0: 271}
                                                    var v94 bool = v3
                                                    _ = fi92
                                                    _ = v93
                                                    _ = v94
                                                }
                                            }
                                        } else {
                                            if len(v47) >= 2 {
                                                v47 = v47[1:len(v47)]
                                            }
                                            if len(v47) > 0 {
                                                v47[0] = 9223372036854775807
                                            }
                                            v65 = (v65 - (fi82 * 2147483647))
                                        }
                                        _ = v89
                                    }
                                }
                                for fi95 := 0; fi95 < 7; fi95++ {
                                    v69 = (v2 | (v2 / (((v69 >> (((^(v80))) & 63))) | 1)))
                                    v96 := GenS1{F0: (^(v5))}
                                    _ = fi95
                                    _ = v96
                                }
                                v47 = append(v47, 42)
                                _ = fi82
                            }
                            switch "foo bar" {
                            case "":
                                if (((v5 + uint8(v4)) <= 174) && ((3.14159 <= ((((1.5 / -2.25) - 1e10) / 1e10) - float64(1e10))) && (!(!(!(!false)))))) {
                                    var v97 int8 = (v81 % ((-20) | 1))
                                    _ = v97
                                } else {
                                    v42 = -1
                                }
                            case "\n\t":
                                var v98 bool = (3.14159 > 0.0)
                                _ = v98
                            case "hello":
                                if false {
                                    switch v65 {
                                    case 32767:
                                        switch "" {
                                        case "a":
                                            var v99 string = (((("hello" + "a") + ((("" + "hello") + "\n\t") + ("\n\t" + ("" + "foo bar")))) + "a") + ("" + "a"))
                                            _ = v99
                                        case "hello":
                                            for fi100 := 0; fi100 < 20; fi100++ {
                                                for fi101 := 0; fi101 < 20; fi101++ {
                                                    var v102 uint16 = (^(v4))
                                                    v103 := GenS11{F0: (v68 + v68), F1: -128}
                                                    var v104 bool = (!false)
                                                    _ = fi101
                                                    _ = v102
                                                    _ = v103
                                                    _ = v104
                                                }
                                                _ = fi100
                                            }
                                        case "":
                                            switch "hello" {
                                            case "\n\t":
                                                var v105 int = (v1 | v65)
                                                _ = v105
                                            default:
                                                var v106 uint = 703
                                                _ = v106
                                            }
                                        default:
                                            for fi107 := 0; fi107 < 9; fi107++ {
                                                var v108 int8 = v81
                                                v2 = (v69 % ((301) | 1))
                                                var v109 int = 127
                                                _ = fi107
                                                _ = v108
                                                _ = v109
                                            }
                                        }
                                    default:
                                        fmt.Println(v2, v4, v5)
                                    }
                                    v110 := GenS12{F0: -128}
                                    _ = v110
                                } else {
                                    fmt.Println(v68)
                                }
                            default:
                                if len(v47) > 0 {
                                    v47[0] = v1
                                }
                            }
                            var v111 int = (v65 - v65)
                            _ = v111
                        } else {
                            var v112 int32 = -1
                            _ = v112
                        }
                    case "foo bar":
                        var v113 int = v1
                        _ = v113
                    case "\n\t":
                        var v114 string = "a"
                        _ = v114
                    }
                    fmt.Println(v4, v81)
                    _ = v81
                }
                if len(v47) >= 2 {
                    v47 = v47[1:len(v47)]
                }
                _ = v80
            }
        }
    case "\n\t":
        v76 += v76
    }
    for fi115 := 0; fi115 < 9; fi115++ {
        for fi116 := 0; fi116 < 14; fi116++ {
            switch (^(v1)) {
            case 9223372036854775807:
                var v117 float64 = float64((3.14159 + (1.5 / -2.25)))
                _ = v117
            case -1:
                if (4294967295 < (v69 >> ((163) & 63))) {
                    switch (-(fi115)) {
                    case -1:
                        switch ("" + "") {
                        case "":
                            v118 := GenS12{F0: -103}
                            _ = v118
                        case "foo bar":
                            var v119 int16 = 32767
                            _ = v119
                        }
                    default:
                        var v120 uint32 = 816
                        _ = v120
                    }
                    if v3 {
                        v3 = v3
                        if ("foo bar" == "\n\t") {
                            if ((v68 >> ((4294967295) & 63)) >= uint64(v5)) {
                                if (v3 && v3) {
                                    fmt.Println(v4, v3)
                                    var v121 bool = false
                                    _ = v121
                                }
                            } else {
                                if len(v47) > 0 {
                                    v47[0] = 32768
                                }
                                var v122 int8 = 0
                                v65 = v65
                                _ = v122
                            }
                            var v123 int32 = -1
                            if len(v47) > 0 {
                                v47[0] = (-(fi116))
                            }
                            if (false || false) {
                                v124 := GenS13{F0: 100, F1: v5}
                                var v125 uint16 = 4
                                _ = v124
                                _ = v125
                            }
                            v76 -= (v76 ^ v76)
                            switch (^(v65)) {
                            case 2147483647:
                                switch v1 {
                                case 32767:
                                    if false {
                                        var v126 uint8 = (v5 & (v5 / ((v5) | 1)))
                                        _ = v126
                                    } else {
                                        v127 := GenS14{F0: fi115, F1: 4294967295, F2: (v76 & 9223372036854775807)}
                                        _ = v127
                                    }
                                case 32768:
                                    if (v3 || (false && ((v5 << ((0) & 63)) < (v5 - 42)))) {
                                        v128 := GenS15{F0: 1, F1: (-(v76)), F2: 2}
                                        _ = v128
                                    }
                                case 42:
                                    v76 += (v76 / (((v76 >> ((4294967295) & 63))) | 1))
                                }
                                v129 := GenS1{F0: (^(v5))}
                                v130 := GenS3{F0: "hello", F1: (v123 * v123)}
                                _ = v129
                                _ = v130
                            }
                            _ = v123
                        }
                    }
                }
            default:
                if (!false) {
                    fmt.Println(v69, fi115, v4)
                    v65 = v1
                    _ = len(v47)
                } else {
                    var v131 uint8 = v5
                    if (!v3) {
                        fmt.Println(fi115, v3, v5)
                        v5 += v5
                        fmt.Println(v68, v1)
                        v4 += (v4 << ((795) & 63))
                        fmt.Println(v4, fi115)
                    } else {
                        if true {
                            v69 = 395
                            switch ("hello" + "\n\t") {
                            case "":
                                fmt.Println(fi116, v69)
                            case "foo bar":
                                var v132 string = "a"
                                _ = v132
                            case "a":
                                v69 = 640
                            default:
                                v65 += fi115
                            }
                        } else {
                            v69 += v2
                            if v3 {
                                switch "foo bar" {
                                case "a":
                                    v76 = (v76 / ((v76) | 1))
                                case "\n\t":
                                    v133 := GenS6{F0: (v2 ^ (v69 ^ v68))}
                                    _ = v133
                                case "foo bar":
                                    var v134 float32 = 1.5
                                    _ = v134
                                }
                                v131 -= 1
                            }
                        }
                        if false {
                            var v135 uint32 = 65535
                            var v136 string = "a"
                            _ = v135
                            _ = v136
                        }
                        if (v3 && false) {
                            switch (v1 * fi115) {
                            case 32768:
                                if v3 {
                                    if (v3 || v3) {
                                        v137 := GenS16{F0: true, F1: 2147483647}
                                        _ = len(v47)
                                        v138 := GenS3{F0: ("\n\t" + "a"), F1: -2147483648}
                                        fmt.Println(fi115, v5)
                                        fmt.Println(v2, fi115, v2)
                                        switch "" {
                                        case "":
                                            v47 = append(v47, fi116)
                                        case "hello":
                                            v47 = append(v47, int(v42))
                                        case "a":
                                            var v139 uint32 = 42
                                            _ = v139
                                        default:
                                            var v140 float32 = 0.1
                                            _ = v140
                                        }
                                        _ = v137
                                        _ = v138
                                    }
                                } else {
                                    _ = len(v47)
                                    v2 = uint64(v4)
                                    var v141 int16 = (v42 + 32767)
                                    _ = v141
                                }
                                if (v3 && ((int64(v4) < (v76 - 306)) && v3)) {
                                    switch "hello" {
                                    case "foo bar":
                                        var v142 float64 = (3.14159 / 0.1)
                                        fmt.Println(v69)
                                        _ = v142
                                    case "hello":
                                        v143 := GenS17{F0: 18446744073709551615}
                                        _ = v143
                                    default:
                                        v47 = append(v47, (v1 % (((fi115 + (^(fi116)))) | 1)))
                                    }
                                } else {
                                    _ = len(v47)
                                    var v144 int16 = v42
                                    v5 -= v5
                                    _ = v144
                                }
                                switch ((("a" + ("a" + "a")) + "") + "foo bar") {
                                case "foo bar":
                                    var v145 uint64 = (v69 | 36)
                                    _ = v145
                                case "hello":
                                    v69 += v2
                                case "":
                                    if (((v76 + (v76 & genFn0(("\n\t" + "a"), "\n\t"))) != v76) && (false || (false && v3))) {
                                        if false {
                                            switch (-(fi116)) {
                                            case 4294967295:
                                                var v146 int8 = 127
                                                _ = v146
                                            default:
                                                v147 := GenS18{F0: -2147483648}
                                                _ = v147
                                            }
                                            var v148 uint8 = (v5 ^ 21)
                                            _ = v148
                                        }
                                    } else {
                                        var v149 string = "hello"
                                        _ = v149
                                    }
                                }
                                fmt.Println(fi116)
                            case 23:
                                if true {
                                    v69 += v2
                                }
                            case 0:
                                v150 := GenS1{F0: v131}
                                _ = v150
                            default:
                                v47 = append(v47, (^(fi116)))
                            }
                            var v151 uint64 = (^(v69))
                            if (!v3) {
                                fmt.Println(v5, v65)
                            }
                            if v3 {
                                v131 += (^(v5))
                                if len(v47) >= 2 {
                                    v47 = v47[1:len(v47)]
                                }
                                v152 := GenS4{F0: 1, F1: v1, F2: v4}
                                switch ((("a" + ((("a" + "hello") + "\n\t") + "hello")) + "foo bar") + ("a" + ("foo bar" + ("foo bar" + ("foo bar" + "hello"))))) {
                                case "\n\t":
                                    switch 42 {
                                    case 293:
                                        var v153 bool = true
                                        _ = v153
                                    case 127:
                                        v68 = (v2 >> ((183) & 63))
                                    default:
                                        switch "" {
                                        case "a":
                                            var v154 string = "a"
                                            _ = v154
                                        default:
                                            var v155 float32 = (float32(float64((0.1 / 1e10))) + ((0.1 - 1.5) / float32((1e10 + 1.5))))
                                            _ = v155
                                        }
                                    }
                                }
                                var v156 float64 = 0.1
                                _ = v152
                                _ = v156
                            }
                            var v157 bool = false
                            _ = v151
                            _ = v157
                        }
                    }
                    v1 = v65
                    _ = v131
                }
                if (int32(v69) != -1) {
                    v158 := GenS3{F0: "", F1: -2147483648}
                    switch "foo bar" {
                    case "foo bar":
                        v65 = (fi115 >> ((815) & 63))
                    case "a":
                        if len(v47) >= 2 {
                            v47 = v47[1:len(v47)]
                        }
                    case "\n\t":
                        switch v65 {
                        case 128:
                            v1 = (fi115 + fi116)
                        case -289:
                            v5 = v5
                        }
                    default:
                        var v159 int64 = v76
                        _ = v159
                    }
                    var v160 int16 = (-(v42))
                    _ = v158
                    _ = v160
                }
                v47 = append(v47, v65)
            }
            v47 = append(v47, fi115)
            _ = fi116
        }
        _ = fi115
    }
    v69 -= (^(v2))
    var v161 uint64 = (v69 / ((1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v161:", v161)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v5:", v5)
    fmt.Println("v65:", v65)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v161
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v47
    _ = v5
    _ = v65
    _ = v68
    _ = v69
    _ = v76
}
