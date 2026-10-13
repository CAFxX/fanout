package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32; F1 int64; F2 int8 }
type GenS7 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS8 struct { F0 uint64; F1 int16 }
type GenS9 struct { F0 int8; F1 string }
type GenS10 struct { F0 uint8; F1 int64 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int }

func main() {
    var v1 string = "\n\t"
    v2 := make(map[int]bool)
    v3 := []uint32{1, 681}
    v4 := []int{-1, 1, -1}
    v5 := []string{v1, "foo bar", v1, v1}
    var v6 uint16 = 0
    switch (v1 + "a") {
    case "a":
        switch 2147483647 {
        case 127:
            v7 := GenS1{F0: 4294967295, F1: (v1 <= "a"), F2: 0}
            _ = v7
        case 986:
            var v8 float32 = (float32(1e10) / -2.25)
            _ = v8
        default:
            switch v1 {
            case "hello":
                if (!(515 != -50)) {
                    if (239 < 128) {
                        fmt.Println(v1, v6)
                        v9 := v2[42]
                        v10 := GenS1{F0: 42, F1: (-341 == 1), F2: 9223372036854775807}
                        _ = v10
                        _ = v9
                    }
                    var v11 bool = false
                    _ = v11
                } else {
                    if (!false) {
                        var v12 uint64 = 42
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        var v13 string = (v1 + v1)
                        _ = v12
                        _ = v13
                    }
                    switch 1 {
                    case 4294967295:
                        var v14 uint8 = 46
                        _ = v14
                    case 128:
                        switch (v1 + v1) {
                        case "":
                            v6 -= 42
                        case "foo bar":
                            if (false && false) {
                                switch -9223372036854775808 {
                                case -1:
                                    v15 := GenS1{F0: 368, F1: false, F2: 42}
                                    _ = v15
                                case -228:
                                    v6 -= v6
                                case 127:
                                    var v16 uint32 = 1
                                    _ = v16
                                }
                                v1 = v1
                                fmt.Println(v1)
                            } else {
                                fmt.Println(v6, v6)
                                v1 = v1
                            }
                        }
                    }
                    for fi17 := 0; fi17 < 4; fi17++ {
                        v3 = append(v3, 4294967295)
                        _ = fi17
                    }
                }
            default:
                if false {
                    v6 = v6
                    v2[42] = (!true)
                }
            }
        }
    default:
        v18 := GenS2{F0: 943}
        _ = v18
    }
    v1 = ""
    v1 = v1
    var v19 int8 = 75
    var v20 int = -2147483648
    v21 := GenS1{F0: 746, F1: false, F2: 1}
    v22 := GenS2{F0: -470}
    v23 := GenS1{F0: 4294967295, F1: false, F2: -9223372036854775808}
    var v24 int32 = -1
    switch (v1 + (v1 + v1)) {
    case "hello":
        if false {
            var v25 int8 = -96
            fmt.Println(v1, v1)
            for fi26 := 0; fi26 < 13; fi26++ {
                var v27 float32 = -2.25
                _ = fi26
                _ = v27
            }
            _ = v25
        } else {
            var v28 float64 = 1.5
            if (2 < 42) {
                if true {
                    var v29 int8 = v19
                    switch 42 {
                    case 2147483647:
                        for fi30 := 0; fi30 < 12; fi30++ {
                            v19 -= (v29 + v19)
                            v3 = append(v3, 819)
                            _ = fi30
                        }
                    }
                    _ = v29
                } else {
                    for fi31 := 0; fi31 < 19; fi31++ {
                        for fi32 := 0; fi32 < 3; fi32++ {
                            v33 := GenS2{F0: -1}
                            _ = fi32
                            _ = v33
                        }
                        for fi34 := 0; fi34 < 17; fi34++ {
                            v20 -= (fi34 | fi31)
                            v2[(fi31 / (((v20 / (((v20 + v20)) | 1))) | 1))] = (!false)
                            _ = fi34
                        }
                        _ = fi31
                    }
                    _ = len(v3)
                }
                v20 += (v20 + v20)
                switch "a" {
                case "hello":
                    var v35 float64 = -2.25
                    _ = v35
                case "\n\t":
                    switch v20 {
                    case 42:
                        for fi36 := 0; fi36 < 11; fi36++ {
                            fmt.Println(v1)
                            var v37 float64 = (v28 + v28)
                            _ = fi36
                            _ = v37
                        }
                    default:
                        for fi38 := 0; fi38 < 8; fi38++ {
                            var v39 uint8 = 0
                            _ = fi38
                            _ = v39
                        }
                    }
                case "foo bar":
                    v6 -= v6
                }
            } else {
                v1 = v1
                var v40 float64 = 0.1
                v28 = (0.1 - float64(3.14159))
                _ = v40
            }
            for fi41 := 0; fi41 < 13; fi41++ {
                v42 := GenS3{F0: 372, F1: 1}
                if (!(1e10 < 1.5)) {
                    switch "" {
                    case "foo bar":
                        fmt.Println(v6, v28, fi41)
                    case "a":
                        var v43 uint = 68
                        _ = v43
                    }
                    switch ((v1 + "") + v1) {
                    case "hello":
                        v44 := GenS2{F0: 1000}
                        _ = v44
                    case "foo bar":
                        v19 -= v19
                    case "\n\t":
                        for fi45 := 0; fi45 < 15; fi45++ {
                            v1 = (v1 + "")
                            v46 := GenS4{F0: v6}
                            _ = fi45
                            _ = v46
                        }
                    default:
                        v28 -= float64(((0.0 / 1.0) / ((3.14159 + (0.0 / -2.25)) * 3.14159)))
                    }
                }
                _ = fi41
                _ = v42
            }
            _ = v28
        }
    case "foo bar":
        var v47 uint32 = 42
        _ = v47
    default:
        v48 := GenS5{F0: 2}
        _ = v48
    }
    if true {
        v49 := GenS2{F0: 42}
        switch "" {
        case "\n\t":
            for fi50 := 0; fi50 < 1; fi50++ {
                if true {
                    if (false && (2 < 0)) {
                        for fi51 := 0; fi51 < 19; fi51++ {
                            var v52 int16 = -969
                            var v53 uint = 0
                            if (65535 < uint32(v6)) {
                                var v54 uint32 = 1
                                if false {
                                    fmt.Println(v52)
                                    var v55 uint8 = 157
                                    _ = v55
                                }
                                if (28 != 42) {
                                    v56 := GenS1{F0: 4294967295, F1: (false && false), F2: -9223372036854775808}
                                    _ = v56
                                } else {
                                    if ((((3.14159 - (3.14159 / ((0.0 + 1.5) + (0.0 - -2.25)))) * (0.1 + ((-2.25 / ((1.5 + 1e10) * -2.25)) + 1e10))) * float64((float32(1.5) - 1e10))) <= ((float64(((0.0 / 1e10) + 1.5)) + 3.14159) + 3.14159)) {
                                        switch (v1 + v1) {
                                        case "hello":
                                            v6 -= v6
                                        case "foo bar":
                                            if false {
                                                v6 -= v6
                                                var v57 uint = v53
                                                _ = v57
                                            }
                                        case "":
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        default:
                                            fmt.Println(v6, v6)
                                        }
                                    } else {
                                        var v58 uint32 = (v54 ^ 765)
                                        if false {
                                            v20 -= (fi50 << (((v53 & v53)) & 63))
                                            var v59 int32 = -990
                                            var v60 int = fi50
                                            if false {
                                                if (((683 != 4294967295) && (0.0 == 1.5)) && true) {
                                                    if true {
                                                        var v61 int = int(v54)
                                                        _ = v61
                                                    }
                                                    var v62 uint32 = uint32(v53)
                                                    if (!false) {
                                                        var v63 uint8 = 1
                                                        var v64 uint32 = v54
                                                        var v65 int16 = (-(v52))
                                                        v66 := GenS6{F0: v24, F1: 0, F2: v19}
                                                        _ = v63
                                                        _ = v64
                                                        _ = v65
                                                        _ = v66
                                                    }
                                                    _ = v62
                                                } else {
                                                    var v67 int64 = 9223372036854775807
                                                    _ = v67
                                                }
                                                var v68 uint8 = 127
                                                _ = v68
                                            } else {
                                                var v69 float64 = (1.5 / (float64(1e10) - -2.25))
                                                _ = v69
                                            }
                                            switch (^(fi51)) {
                                            case -326:
                                                var v70 int8 = 100
                                                _ = v70
                                            case -2147483648:
                                                if false {
                                                    if (v24 >= (v24 % ((v59) | 1))) {
                                                        var v71 uint64 = 0
                                                        _ = v71
                                                    } else {
                                                        var v72 bool = (!(v58 > 762))
                                                        _ = v72
                                                    }
                                                    fmt.Println(v53)
                                                }
                                            default:
                                                v52 = 32767
                                            }
                                            _ = v59
                                            _ = v60
                                        } else {
                                            var v73 float64 = (1e10 - ((float64((1e10 * 1.5)) + float64(0.0)) + -2.25))
                                            _ = v73
                                        }
                                        v54 -= (v58 >> (((v53 & 18446744073709551615)) & 63))
                                        _ = v58
                                    }
                                    switch "" {
                                    case "\n\t":
                                        switch v1 {
                                        case "hello":
                                            fmt.Println(v20, v53)
                                        }
                                    case "a":
                                        v74 := GenS7{F0: v52, F1: 4294967295, F2: v54}
                                        _ = v74
                                    case "hello":
                                        v53 = (^(v53))
                                    default:
                                        fmt.Println(v52, v6)
                                    }
                                    fmt.Println(v1)
                                }
                                _ = v54
                            }
                            _ = fi51
                            _ = v52
                            _ = v53
                        }
                        var v75 bool = ((97 <= 1) || false)
                        if (v75 || ((((1.5 + (0.0 / float64(1.5))) + (1.5 * 0.1)) - -2.25) <= 1e10)) {
                            if (!(2 >= 42)) {
                                v75 = (!(!(!v75)))
                                var v76 int32 = v24
                                _ = v76
                            } else {
                                var v77 string = ("\n\t" + "\n\t")
                                if (-74 <= -637) {
                                    var v78 int = fi50
                                    _ = v78
                                }
                                _ = v77
                            }
                            fmt.Println(v19, v75, v75)
                        } else {
                            v79 := GenS6{F0: 12, F1: 1, F2: v19}
                            var v80 float64 = 3.14159
                            v20 = (v20 * 1)
                            _ = v79
                            _ = v80
                        }
                        _ = v75
                    } else {
                        var v81 float64 = (0.0 + 3.14159)
                        for fi82 := 0; fi82 < 10; fi82++ {
                            v83 := GenS8{F0: 4294967295, F1: 1}
                            _ = fi82
                            _ = v83
                        }
                        _ = v81
                    }
                    v84 := GenS7{F0: 42, F1: 2, F2: 42}
                    var v85 uint16 = v6
                    _ = v84
                    _ = v85
                }
                switch (fi50 / ((int(v24)) | 1)) {
                case -1:
                    v1 = ((v1 + "\n\t") + "\n\t")
                case 42:
                    var v86 int16 = 466
                    _ = v86
                case 9223372036854775807:
                    if (!false) {
                        for fi87 := 0; fi87 < 8; fi87++ {
                            v88 := GenS4{F0: v6}
                            v6 += (^(v6))
                            var v89 int64 = 0
                            if false {
                                _ = len(v5)
                                fmt.Println(v89, v89, fi50)
                                switch -2147483648 {
                                case 32768:
                                    v90 := v2[0]
                                    _ = v90
                                case -1:
                                    v91 := GenS2{F0: 42}
                                    _ = v91
                                case 0:
                                    switch (v1 + (v1 + "hello")) {
                                    case "\n\t":
                                        if ((-(v89)) == -1) {
                                            switch v1 {
                                            case "hello":
                                                v92 := GenS9{F0: v19, F1: ("a" + "foo bar")}
                                                _ = v92
                                            default:
                                                v24 = v24
                                            }
                                            fmt.Println(fi87)
                                        } else {
                                            var v93 uint64 = 725
                                            switch fi87 {
                                            case -818:
                                                if true {
                                                    fmt.Println(v89)
                                                } else {
                                                    var v94 uint32 = 4294967295
                                                    _ = v94
                                                }
                                            default:
                                                var v95 float64 = 0.1
                                                _ = v95
                                            }
                                            _ = v93
                                        }
                                    case "a":
                                        fmt.Println(v19)
                                    }
                                default:
                                    v5 = append(v5, v1)
                                }
                            } else {
                                var v96 uint16 = 255
                                switch -546 {
                                case 32767:
                                    v6 = v6
                                case -1:
                                    if (0 > 1) {
                                        switch v1 {
                                        case "foo bar":
                                            switch fi50 {
                                            case 0:
                                                v97 := GenS9{F0: v19, F1: v1}
                                                _ = v97
                                            case -980:
                                                var v98 int32 = v24
                                                _ = v98
                                            }
                                        }
                                        v99 := GenS1{F0: 0, F1: false, F2: (v89 / ((int64(v6)) | 1))}
                                        v100 := GenS3{F0: (^(v96)), F1: v6}
                                        _ = v100
                                        _ = v99
                                    } else {
                                        switch (fi87 & -811) {
                                        case -9223372036854775808:
                                            v20 -= v20
                                        case 1:
                                            fmt.Println(v24, v6)
                                        default:
                                            v101 := GenS4{F0: (v96 | 1)}
                                            _ = v101
                                        }
                                        v19 += (v19 + v19)
                                        if (false && ((!(!true)) || (!false))) {
                                            v89 += (v89 * (v89 * v89))
                                        }
                                        switch (-(fi87)) {
                                        case 32767:
                                            v24 = 42
                                        case -9223372036854775808:
                                            var v102 int16 = -727
                                            _ = v102
                                        case 9223372036854775807:
                                            v103 := GenS10{F0: 127, F1: 42}
                                            fmt.Println(v96)
                                            if false {
                                                if ((255 == 186) || false) {
                                                    v104 := v2[(-(fi50))]
                                                    fmt.Println(v6, v6)
                                                    _ = v104
                                                }
                                                switch v20 {
                                                case 0:
                                                    v96 += v96
                                                case 9223372036854775807:
                                                    v1 = v1
                                                case -1:
                                                    var v105 int16 = 42
                                                    _ = v105
                                                default:
                                                    if len(v5) > 0 {
                                                        v5[0] = ("hello" + ("a" + v1))
                                                    }
                                                }
                                            } else {
                                                v106 := GenS1{F0: 0, F1: false, F2: (v89 ^ -1)}
                                                _ = v106
                                            }
                                            var v107 string = "a"
                                            _ = v103
                                            _ = v107
                                        default:
                                            v108 := GenS1{F0: 42, F1: (!true), F2: int64(v24)}
                                            _ = v108
                                        }
                                    }
                                case 0:
                                    v109 := GenS11{F0: 4294967295}
                                    _ = v109
                                }
                                switch (v1 + v1) {
                                case "\n\t":
                                    v19 -= 42
                                case "foo bar":
                                    if false {
                                        var v110 float64 = 0.1
                                        _ = v110
                                    }
                                case "":
                                    var v111 uint32 = 280
                                    if true {
                                        v112 := GenS12{F0: 187}
                                        fmt.Println(v20, fi87)
                                        _ = v112
                                    }
                                    _ = v111
                                default:
                                    var v113 string = (v1 + (v1 + ((v1 + "") + (v1 + (v1 + (("\n\t" + v1) + v1))))))
                                    _ = v113
                                }
                                _ = v96
                            }
                            _ = len(v4)
                            _ = fi87
                            _ = v88
                            _ = v89
                        }
                        switch (fi50 << ((1) & 63)) {
                        case 32768:
                            var v114 int64 = -975
                            _ = v114
                        case -180:
                            _ = len(v3)
                        default:
                            var v115 int = (v20 * (fi50 >> ((762) & 63)))
                            _ = v115
                        }
                        switch fi50 {
                        case 127:
                            var v116 uint32 = 2147483647
                            _ = v116
                        default:
                            for fi117 := 0; fi117 < 13; fi117++ {
                                var v118 uint16 = (v6 / (((v6 | 793)) | 1))
                                v118 += v6
                                _ = fi117
                                _ = v118
                            }
                        }
                    } else {
                        v6 = v6
                        v119 := GenS9{F0: (v19 << ((2) & 63)), F1: ""}
                        _ = v119
                    }
                default:
                    switch (("" + v1) + "\n\t") {
                    case "\n\t":
                        var v120 string = ((v1 + v1) + v1)
                        _ = v120
                    case "hello":
                        switch v1 {
                        case "":
                            var v121 uint8 = 192
                            _ = v121
                        case "\n\t":
                            v122 := v2[(v20 ^ v20)]
                            _ = v122
                        case "foo bar":
                            fmt.Println(v19)
                        default:
                            fmt.Println(fi50, v19, v6)
                        }
                    case "foo bar":
                        var v123 int64 = 42
                        _ = v123
                    }
                }
                for fi124 := 0; fi124 < 6; fi124++ {
                    if (0 < 568) {
                        if true {
                            v1 = v1
                            var v125 int64 = 9223372036854775807
                            if true {
                                v20 = (-(v20))
                                var v126 uint32 = 270
                                _ = v126
                            }
                            _ = v125
                        }
                        var v127 float32 = 1.5
                        _ = v127
                    } else {
                        v128 := GenS7{F0: 42, F1: 2, F2: 2}
                        if len(v4) > 0 {
                            v4[0] = 4294967295
                        }
                        _ = v128
                    }
                    _ = fi124
                }
                _ = fi50
            }
        case "":
            fmt.Println(v1, v19, v20)
        default:
            for fi129 := 0; fi129 < 12; fi129++ {
                for fi130 := 0; fi130 < 17; fi130++ {
                    if true {
                        v19 += v19
                        fmt.Println(v6, v20)
                    }
                    switch v1 {
                    case "":
                        if ((29 < 1) || true) {
                            if false {
                                v1 = v1
                                var v131 uint32 = 0
                                if ((1e10 - -2.25) <= -2.25) {
                                    v132 := GenS11{F0: 55}
                                    v24 -= v24
                                    v133 := GenS8{F0: 2, F1: 32767}
                                    _ = v132
                                    _ = v133
                                }
                                _ = v131
                            }
                            switch v1 {
                            case "hello":
                                v20 = 4294967295
                            case "foo bar":
                                v24 = v24
                            case "\n\t":
                                delete(v2, 363)
                            }
                            var v134 int16 = 42
                            var v135 int32 = v24
                            _ = v134
                            _ = v135
                        } else {
                            switch v1 {
                            case "foo bar":
                                v1 = "a"
                            case "hello":
                                var v136 float32 = (1.5 / 1.0)
                                _ = v136
                            }
                            fmt.Println(v19, fi130, v19)
                        }
                    default:
                        v137 := GenS7{F0: -32768, F1: 1, F2: 42}
                        _ = v137
                    }
                    _ = fi130
                }
                v20 = fi129
                fmt.Println(v24)
                _ = fi129
            }
        }
        fmt.Println(v19)
        _ = v49
    }
    for fi138 := 0; fi138 < 16; fi138++ {
        v20 += 127
        var v139 string = "\n\t"
        v140 := GenS10{F0: 42, F1: -9223372036854775808}
        _ = fi138
        _ = v139
        _ = v140
    }
    switch 128 {
    case 445:
        if len(v4) > 0 {
            v4[0] = v20
        }
    case 1:
        var v141 float32 = (3.14159 * ((0.1 / 1.5) + (-2.25 + (((-2.25 + (3.14159 / 1.0)) + -2.25) * 0.0))))
        _ = v141
    case 2147483647:
        var v142 float64 = 3.14159
        _ = v142
    }
    var v143 uint32 = 255
    if true {
        v24 += (v24 - v24)
    }
    switch ("a" + "\n\t") {
    case "foo bar":
        for fi144 := 0; fi144 < 4; fi144++ {
            delete(v2, 0)
            _ = fi144
        }
    default:
        v145 := GenS13{F0: 239}
        _ = v145
    }
    fmt.Println("v1:", v1)
    fmt.Println("v143:", v143)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v143
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
