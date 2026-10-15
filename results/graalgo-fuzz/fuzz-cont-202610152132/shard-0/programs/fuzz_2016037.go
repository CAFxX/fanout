package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int32; F2 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 uint; F1 int32; F2 uint64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int8; F1 bool }
type GenS8 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint; F1 int8 }
type GenS11 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS12 struct { F0 int8; F1 int32 }
type GenS13 struct { F0 int; F1 int32 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 int32; F1 uint16 }
func genFn0() int16 {
    v6 := []int{-9223372036854775808, 4294967295}
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 float32 = (0.1 + ((-2.25 - 0.1) - ((0.0 - -2.25) * -2.25)))
        fmt.Println(v8, v8)
        _ = fi7
        _ = v8
    }
    var v9 uint32 = 42
    v9 -= v9
    _ = v6
    _ = v9
    return 32767
}
func genFn1() uint16 {
    var v10 float64 = float64((-2.25 * 0.0))
    v11 := GenS1{F0: 442}
    var v12 float64 = -2.25
    _ = v10
    _ = v11
    _ = v12
    return 255
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]uint32)
    var v3 int16 = -1
    var v4 int64 = 0
    var v5 int8 = 1
    if false {
        v13 := []int{}
        for fi14 := 0; fi14 < 9; fi14++ {
            for fi15 := 0; fi15 < 1; fi15++ {
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                _ = fi15
            }
            for fi16 := 0; fi16 < 17; fi16++ {
                v17 := GenS1{F0: 2147483648}
                _ = fi16
                _ = v17
            }
            var v18 int8 = int8(fi14)
            _ = fi14
            _ = v18
        }
        var v19 uint32 = 792
        _ = v13
        _ = v19
    } else {
        if true {
            v20 := v2["hello"]
            for fi21 := 0; fi21 < 3; fi21++ {
                for fi22 := 0; fi22 < 1; fi22++ {
                    v23 := GenS1{F0: (v20 & v20)}
                    v24 := GenS1{F0: v20}
                    if false {
                        v25 := GenS1{F0: 424}
                        _ = v25
                    }
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                for fi26 := 0; fi26 < 2; fi26++ {
                    v20 = 1
                    _ = fi26
                }
                _ = fi21
            }
            if true {
                fmt.Println(v3, v20)
                fmt.Println(v4, v5, v20)
                if (1 > 42) {
                    if (!(!false)) {
                        switch 4294967295 {
                        case -585:
                            if (!(928 > 13)) {
                                if (!(v3 >= (v3 / ((genFn0()) | 1)))) {
                                    if (194 > 68) {
                                        for fi27 := 0; fi27 < 5; fi27++ {
                                            fmt.Println(fi27, v4)
                                            _ = fi27
                                        }
                                    } else {
                                        if true {
                                            v28 := GenS1{F0: (^(v20))}
                                            _ = v28
                                        }
                                        var v29 int = -9223372036854775808
                                        _ = v29
                                    }
                                    v5 = -86
                                    for fi30 := 0; fi30 < 10; fi30++ {
                                        fmt.Println(v20, v3)
                                        v1[0] = "\n\t"
                                        _ = fi30
                                    }
                                } else {
                                    for fi31 := 0; fi31 < 9; fi31++ {
                                        for fi32 := 0; fi32 < 19; fi32++ {
                                            fmt.Println(v4, v20, v5)
                                            switch "\n\t" {
                                            case "hello":
                                                var v33 string = ("" + (("\n\t" + "foo bar") + ""))
                                                _ = v33
                                            case "a":
                                                delete(v1, (^(v20)))
                                            case "foo bar":
                                                if (42 == 73) {
                                                    switch (fi31 >> ((480) & 63)) {
                                                    case 128:
                                                        fmt.Println(v20, v3, fi32)
                                                    case 0:
                                                        v3 -= v3
                                                    case 2147483647:
                                                        v34 := GenS2{F0: (!(true || ("" >= "a"))), F1: -1, F2: fi31}
                                                        _ = v34
                                                    default:
                                                        v20 -= v20
                                                    }
                                                    fmt.Println(v4, v20)
                                                    v5 -= v5
                                                }
                                            default:
                                                v1[(v20 - v20)] = ""
                                            }
                                            var v35 string = (("" + "hello") + "\n\t")
                                            _ = fi32
                                            _ = v35
                                        }
                                        if ((1.5 / (0.1 - float32(-2.25))) == 0.0) {
                                            if false {
                                                if (!(1e10 >= (0.0 / -2.25))) {
                                                    var v36 uint = 255
                                                    v37 := GenS3{F0: 255}
                                                    _ = v36
                                                    _ = v37
                                                }
                                                switch -9223372036854775808 {
                                                case 1:
                                                    v5 += -128
                                                }
                                                for fi38 := 0; fi38 < 7; fi38++ {
                                                    v5 -= v5
                                                    _ = fi38
                                                }
                                            }
                                            for fi39 := 0; fi39 < 9; fi39++ {
                                                if (true && false) {
                                                    v40 := GenS4{F0: (0.0 <= 1.5), F1: 18446744073709551615}
                                                    _ = v40
                                                } else {
                                                    var v41 uint16 = 2
                                                    if (2 < 227) {
                                                        var v42 int8 = (v5 ^ (v5 % ((100) | 1)))
                                                        _ = v42
                                                    }
                                                    _ = v41
                                                }
                                                var v43 uint = 44
                                                var v44 int8 = v5
                                                _ = fi39
                                                _ = v43
                                                _ = v44
                                            }
                                        }
                                        if (!false) {
                                            if (1 >= 42) {
                                                if (255 != 2) {
                                                    var v45 int32 = 0
                                                    var v46 string = "hello"
                                                    v45 -= (v45 + (v45 / (((v45 % (((v45 & (v45 >> ((2) & 63)))) | 1))) | 1)))
                                                    _ = v45
                                                    _ = v46
                                                } else {
                                                    _ = len(v1)
                                                    v47 := GenS2{F0: false, F1: 0, F2: -822}
                                                    v48 := GenS2{F0: true, F1: 578, F2: (fi31 & (fi31 % ((2147483647) | 1)))}
                                                    _ = v47
                                                    _ = v48
                                                }
                                                var v49 bool = (false && false)
                                                v50 := v2["hello"]
                                                _ = v49
                                                _ = v50
                                            } else {
                                                fmt.Println(v5, fi31)
                                                var v51 bool = (("" + ("a" + "foo bar")) == "")
                                                _ = v51
                                            }
                                            if true {
                                                if false {
                                                    var v52 uint32 = (v20 | (v20 + (v20 / (((^(v20))) | 1))))
                                                    _ = v52
                                                } else {
                                                    v5 += v5
                                                }
                                                v3 = (v3 % ((genFn0()) | 1))
                                                if false {
                                                    if (-2.25 == (1e10 + float32(0.0))) {
                                                        var v53 uint64 = 994
                                                        _ = v53
                                                    } else {
                                                        v20 = 2147483648
                                                        if true {
                                                            var v54 int = fi31
                                                            _ = v54
                                                        } else {
                                                            v55 := GenS1{F0: (v20 & v20)}
                                                            _ = v55
                                                        }
                                                    }
                                                } else {
                                                    var v56 uint = 255
                                                    for fi57 := 0; fi57 < 8; fi57++ {
                                                        if false {
                                                            v58 := GenS5{F0: (v56 + (v56 + v56)), F1: -1, F2: 42}
                                                            v59 := GenS4{F0: (v3 >= (v3 << ((824) & 63))), F1: 42}
                                                            switch 2147483647 {
                                                            case 2147483647:
                                                                var v60 uint8 = 2
                                                                _ = v60
                                                            }
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                        var v61 uint64 = 42
                                                        _ = fi57
                                                        _ = v61
                                                    }
                                                    _ = v56
                                                }
                                            }
                                        } else {
                                            v62 := []int{32768}
                                            v1[v20] = ""
                                            _ = v62
                                        }
                                        _ = fi31
                                    }
                                    fmt.Println(v5)
                                }
                                if true {
                                    for fi63 := 0; fi63 < 5; fi63++ {
                                        fmt.Println(v5, v5)
                                        var v64 int32 = 1
                                        _ = fi63
                                        _ = v64
                                    }
                                    delete(v2, "foo bar")
                                } else {
                                    if ((v5 + (-(v5))) < (^(v5))) {
                                        if ((!false) && (!false)) {
                                            v65 := GenS1{F0: 2147483648}
                                            if (false && true) {
                                                v66 := GenS1{F0: 42}
                                                _ = v66
                                            }
                                            _ = v65
                                        } else {
                                            if (0.1 == 0.1) {
                                                for fi67 := 0; fi67 < 15; fi67++ {
                                                    if (255 <= 127) {
                                                        for fi68 := 0; fi68 < 13; fi68++ {
                                                            var v69 int8 = v5
                                                            if true {
                                                                v69 = 1
                                                                v69 += v5
                                                            }
                                                            _ = fi68
                                                            _ = v69
                                                        }
                                                        var v70 uint16 = genFn1()
                                                        _ = v70
                                                    } else {
                                                        var v71 uint8 = 2
                                                        if (1e10 != (-2.25 - 1e10)) {
                                                            if (("\n\t" + "foo bar") == ("" + "a")) {
                                                                v1[4294967295] = "hello"
                                                            }
                                                            var v72 uint32 = 2147483647
                                                            var v73 int = 9223372036854775807
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                        if false {
                                                            _ = len(v2)
                                                        } else {
                                                            v71 += (^(v71))
                                                        }
                                                        _ = v71
                                                    }
                                                    var v74 string = "a"
                                                    if (int32(fi67) > 2147483647) {
                                                        var v75 bool = (4294967295 > 18446744073709551615)
                                                        _ = v75
                                                    } else {
                                                        var v76 uint8 = 128
                                                        _ = v76
                                                    }
                                                    _ = fi67
                                                    _ = v74
                                                }
                                                v20 += 42
                                            } else {
                                                v5 = int8(v3)
                                            }
                                            v77 := GenS1{F0: 65535}
                                            v78 := GenS3{F0: 127}
                                            _ = v77
                                            _ = v78
                                        }
                                    }
                                    v79 := []int{}
                                    _ = v79
                                }
                            }
                        default:
                            v1[(v20 - v20)] = ""
                        }
                        if ((1.5 != ((-2.25 * ((0.1 * 3.14159) * -2.25)) + 0.0)) && false) {
                            if (!(true || true)) {
                                for fi80 := 0; fi80 < 9; fi80++ {
                                    if false {
                                        v4 = v4
                                        for fi81 := 0; fi81 < 2; fi81++ {
                                            if false {
                                                if ((1.5 < 0.0) && false) {
                                                    if false {
                                                        if (!(!(1e10 <= 1.5))) {
                                                            var v82 int64 = v4
                                                            _ = v82
                                                        }
                                                    }
                                                }
                                                var v83 int8 = (v5 & (v5 + v5))
                                                _ = v83
                                            }
                                            _ = fi81
                                        }
                                    } else {
                                        var v84 int16 = (^(v3))
                                        v85 := GenS6{F0: 65535}
                                        _ = v84
                                        _ = v85
                                    }
                                    _ = fi80
                                }
                                v86 := []int{598, 128}
                                _ = v86
                            }
                            if (false && (1 == (-(v4)))) {
                                for fi87 := 0; fi87 < 11; fi87++ {
                                    if (-949 <= int32(v5)) {
                                        v88 := []int{fi87, 232, -296}
                                        if (!((42 == 100) || (("a" + ("a" + "a")) > "\n\t"))) {
                                            if true {
                                                v89 := GenS5{F0: 4294967295, F1: 42, F2: 2}
                                                var v90 int8 = 127
                                                _ = v89
                                                _ = v90
                                            }
                                            for fi91 := 0; fi91 < 8; fi91++ {
                                                v2["a"] = v20
                                                switch fi91 {
                                                case -2147483648:
                                                    switch fi87 {
                                                    case -234:
                                                        var v92 uint = 4294967295
                                                        _ = v92
                                                    case 0:
                                                        switch "\n\t" {
                                                        case "a":
                                                            v93 := GenS7{F0: v5, F1: (v5 < v5)}
                                                            _ = v93
                                                        case "foo bar":
                                                            var v94 string = ("foo bar" + ("foo bar" + ""))
                                                            _ = v94
                                                        case "hello":
                                                            v2[""] = 65535
                                                        }
                                                    default:
                                                        v2["foo bar"] = (v20 ^ (v20 * v20))
                                                    }
                                                case -340:
                                                    switch 409 {
                                                    case -115:
                                                        switch 1 {
                                                        case -1:
                                                            v95 := GenS2{F0: ((fi91 ^ fi91) == fi91), F1: 0, F2: fi87}
                                                            _ = v95
                                                        default:
                                                            fmt.Println(v3, v20, v4)
                                                        }
                                                    case 32768:
                                                        switch "hello" {
                                                        case "\n\t":
                                                            v4 = (-(v4))
                                                        case "a":
                                                            fmt.Println(fi87)
                                                        }
                                                    case -9223372036854775808:
                                                        v5 = v5
                                                    }
                                                default:
                                                    v20 += (v20 << ((18446744073709551615) & 63))
                                                }
                                                _ = fi91
                                            }
                                        } else {
                                            switch 127 {
                                            case 127:
                                                v4 = 9223372036854775807
                                            case -512:
                                                var v96 float64 = 0.1
                                                _ = v96
                                            }
                                            v97 := GenS5{F0: 418, F1: -1, F2: 4294967295}
                                            for fi98 := 0; fi98 < 6; fi98++ {
                                                var v99 int16 = (v3 >> ((1) & 63))
                                                v5 += 100
                                                var v100 string = "foo bar"
                                                _ = fi98
                                                _ = v100
                                                _ = v99
                                            }
                                            _ = v97
                                        }
                                        _ = v88
                                    } else {
                                        v101 := GenS1{F0: 139}
                                        _ = v101
                                    }
                                    delete(v1, 2)
                                    _ = fi87
                                }
                                v4 = v4
                                switch "a" {
                                case "foo bar":
                                    v102 := GenS7{F0: v5, F1: true}
                                    _ = v102
                                case "\n\t":
                                    var v103 uint64 = 507
                                    _ = v103
                                case "a":
                                    v20 = (v20 << ((uint(v20)) & 63))
                                default:
                                    fmt.Println(v4)
                                }
                            } else {
                                switch "foo bar" {
                                case "\n\t":
                                    for fi104 := 0; fi104 < 17; fi104++ {
                                        switch ("" + ("foo bar" + "a")) {
                                        case "\n\t":
                                            switch (fi104 / ((int(v4)) | 1)) {
                                            case -1:
                                                fmt.Println(fi104, v4)
                                            case 152:
                                                fmt.Println(fi104, fi104, fi104)
                                            }
                                        default:
                                            if false {
                                                if (!((v5 << ((0) & 63)) <= (v5 & -128))) {
                                                    for fi105 := 0; fi105 < 8; fi105++ {
                                                        if ("" < ("foo bar" + ("a" + ("foo bar" + "a")))) {
                                                            var v106 uint64 = 698
                                                            v4 = -9223372036854775808
                                                            _ = v106
                                                        }
                                                        v107 := GenS8{F0: v20, F1: 524, F2: (v20 & v20)}
                                                        _ = len(v2)
                                                        _ = fi105
                                                        _ = v107
                                                    }
                                                    v108 := GenS2{F0: false, F1: 42, F2: (fi104 + fi104)}
                                                    _ = v108
                                                } else {
                                                    fmt.Println(v20, v5)
                                                    var v109 float64 = 0.1
                                                    _ = v109
                                                }
                                                fmt.Println(fi104, v5)
                                                var v110 string = "\n\t"
                                                _ = v110
                                            }
                                        }
                                        _ = fi104
                                    }
                                }
                            }
                        } else {
                            var v111 int16 = genFn0()
                            _ = v111
                        }
                    } else {
                        var v112 bool = (3.14159 >= 3.14159)
                        v4 -= v4
                        _ = v112
                    }
                    v4 = (-(v4))
                }
            } else {
                _ = len(v2)
                if (!true) {
                    var v113 int = -519
                    var v114 uint16 = 159
                    v113 = v113
                    _ = v113
                    _ = v114
                } else {
                    v115 := GenS7{F0: v5, F1: (!(4294967295 > (v20 >> ((42) & 63))))}
                    var v116 int8 = (v5 % ((v5) | 1))
                    _ = v115
                    _ = v116
                }
                v117 := GenS9{F0: v5}
                _ = v117
            }
            _ = v20
        } else {
            v118 := []int{127, 0}
            v5 -= v5
            if true {
                for fi119 := 0; fi119 < 8; fi119++ {
                    if true {
                        v2[("hello" + (("hello" + "\n\t") + "foo bar"))] = 4294967295
                        var v120 uint = 255
                        _ = v120
                    } else {
                        if (true && ((18446744073709551615 == 2) && true)) {
                            v3 = 262
                        }
                        if false {
                            for fi121 := 0; fi121 < 3; fi121++ {
                                v4 += v4
                                _ = fi121
                            }
                            switch (fi119 << ((587) & 63)) {
                            case -1:
                                v118 = append(v118, 42)
                            case 128:
                                fmt.Println(v3)
                            }
                        } else {
                            for fi122 := 0; fi122 < 15; fi122++ {
                                if false {
                                    var v123 int32 = -1
                                    v5 = v5
                                    _ = v123
                                } else {
                                    v124 := GenS4{F0: true, F1: 1}
                                    _ = v124
                                }
                                v3 -= (v3 | 0)
                                _ = fi122
                            }
                        }
                    }
                    if (3.14159 == 0.1) {
                        v3 += v3
                        var v125 int64 = (v4 >> ((720) & 63))
                        _ = v125
                    } else {
                        var v126 uint16 = 65535
                        v127 := GenS3{F0: 95}
                        var v128 uint = 2
                        _ = v126
                        _ = v127
                        _ = v128
                    }
                    _ = fi119
                }
                var v129 uint64 = 42
                switch (("a" + "foo bar") + "") {
                case "":
                    v130 := GenS4{F0: false, F1: 1}
                    _ = v130
                }
                _ = v129
            }
            _ = v118
        }
        var v131 float32 = 3.14159
        _ = v131
    }
    var v132 float32 = 0.0
    v132 += (((v132 + (((-2.25 + float32((float64(v132) + 1.5))) * v132) / 1.0)) + 0.1) + float32((1.5 + 3.14159)))
    for fi133 := 0; fi133 < 20; fi133++ {
        v134 := []int{(fi133 >> ((4294967295) & 63))}
        var v135 int8 = v5
        _ = fi133
        _ = v134
        _ = v135
    }
    fmt.Println(v132, v4)
    if false {
        v3 += (v3 & (v3 & (v3 - (v3 + (-(v3))))))
    } else {
        switch -2147483648 {
        case 1:
            v5 = 0
        case 220:
            for fi136 := 0; fi136 < 7; fi136++ {
                v1[4294967295] = (("foo bar" + (("" + (("" + "\n\t") + "")) + (("a" + ("a" + "hello")) + "a"))) + "foo bar")
                for fi137 := 0; fi137 < 4; fi137++ {
                    var v138 uint16 = 311
                    var v139 int = (fi136 << ((255) & 63))
                    v2[("\n\t" + "hello")] = 2147483647
                    _ = fi137
                    _ = v138
                    _ = v139
                }
                switch fi136 {
                case -1:
                    v5 -= (v5 - (v5 + 42))
                default:
                    var v140 uint16 = 42
                    _ = v140
                }
                _ = fi136
            }
        }
    }
    var v141 bool = ("foo bar" >= "")
    var v142 int32 = 1
    var v143 uint8 = 2
    fmt.Println(v141, v141, v5)
    switch "a" {
    case "foo bar":
        if ((v141 && v141) && true) {
            v144 := []int{-182, -1}
            for fi145 := 0; fi145 < 17; fi145++ {
                v4 = (v4 / ((int64(v3)) | 1))
                var v146 int16 = (^(v3))
                _ = fi145
                _ = v146
            }
            for fi147 := 0; fi147 < 14; fi147++ {
                var v148 bool = v141
                _ = fi147
                _ = v148
            }
            _ = v144
        } else {
            v149 := GenS4{F0: v141, F1: 164}
            v5 += (-(v5))
            if v141 {
                fmt.Println(v142, v5, v143)
                for fi150 := 0; fi150 < 19; fi150++ {
                    v4 = (-(v4))
                    _ = fi150
                }
            } else {
                var v151 uint8 = (v143 % ((255) | 1))
                fmt.Println(v132)
                _ = v151
            }
            _ = v149
        }
    case "a":
        if false {
            for fi152 := 0; fi152 < 2; fi152++ {
                switch (fi152 & 4294967295) {
                case 32768:
                    if (4294967295 >= 18446744073709551615) {
                        switch ("\n\t" + ("foo bar" + "hello")) {
                        case "":
                            v153 := []int{(fi152 | fi152), fi152, fi152}
                            _ = v153
                        default:
                            switch (("a" + "") + "a") {
                            case "\n\t":
                                switch "a" {
                                case "hello":
                                    switch (fi152 ^ (fi152 ^ (fi152 << ((1) & 63)))) {
                                    case 32768:
                                        var v154 float64 = (1.5 * (3.14159 * 0.0))
                                        _ = v154
                                    }
                                default:
                                    for fi155 := 0; fi155 < 10; fi155++ {
                                        v156 := []int{fi152, 4294967295, (fi152 / (((fi152 >> ((18446744073709551615) & 63))) | 1))}
                                        switch 708 {
                                        case -784:
                                            v4 += (v4 & v4)
                                        default:
                                            if (281 < 849) {
                                                var v157 int = fi152
                                                if (int16(v5) != (v3 & (v3 + (v3 - (v3 / ((genFn0()) | 1)))))) {
                                                    v5 = (v5 / (((v5 - v5)) | 1))
                                                    switch v157 {
                                                    case 334:
                                                        if (1.5 <= 1.5) {
                                                            v158 := GenS10{F0: 0, F1: v5}
                                                            _ = v158
                                                        } else {
                                                            fmt.Println(v4, v5)
                                                        }
                                                    case 9223372036854775807:
                                                        var v159 int64 = (v4 % (((v4 << ((743) & 63))) | 1))
                                                        _ = v159
                                                    case -277:
                                                        if len(v156) >= 2 {
                                                            v156 = v156[1:len(v156)]
                                                        }
                                                    default:
                                                        switch "hello" {
                                                        case "":
                                                            if len(v156) >= 2 {
                                                                v156 = v156[1:len(v156)]
                                                            }
                                                        case "hello":
                                                            if len(v156) >= 2 {
                                                                v156 = v156[1:len(v156)]
                                                            }
                                                        case "\n\t":
                                                            v142 = (v142 * v142)
                                                        default:
                                                            switch "a" {
                                                            case "foo bar":
                                                                switch "a" {
                                                                case "foo bar":
                                                                    if true {
                                                                        var v160 int = v157
                                                                        v142 = (-(v142))
                                                                        _ = v160
                                                                    }
                                                                default:
                                                                    v161 := GenS11{F0: 2, F1: 369, F2: v143}
                                                                    _ = v161
                                                                }
                                                            default:
                                                                fmt.Println(fi152)
                                                            }
                                                            var v162 int16 = (v3 ^ v3)
                                                            v5 = v5
                                                            _ = v162
                                                        }
                                                    }
                                                    v2["\n\t"] = 2
                                                    v4 -= (v4 / (((v4 << ((4294967295) & 63))) | 1))
                                                    v163 := v1[42]
                                                    _ = v163
                                                } else {
                                                    v5 -= (-(v5))
                                                }
                                                _ = v157
                                            } else {
                                                v164 := GenS4{F0: (!(!v141)), F1: 428}
                                                var v165 int32 = (v142 & (v142 ^ 0))
                                                switch "hello" {
                                                case "":
                                                    if true {
                                                        if (((2147483647 < 2147483647) && (42 >= (-(v4)))) && ("foo bar" <= "a")) {
                                                            var v166 bool = (v4 >= (v4 + int64(v3)))
                                                            _ = v166
                                                        }
                                                        var v167 int64 = -9223372036854775808
                                                        fmt.Println(v143, v3)
                                                        _ = v167
                                                    } else {
                                                        switch fi152 {
                                                        case 128:
                                                            switch int(v165) {
                                                            case 619:
                                                                v3 = (v3 % ((genFn0()) | 1))
                                                                switch fi155 {
                                                                case 32768:
                                                                    switch "hello" {
                                                                    case "":
                                                                        if v141 {
                                                                            v142 = (v142 & -1)
                                                                        } else {
                                                                            fmt.Println(v165, v141)
                                                                            v4 = (v4 * int64(v5))
                                                                        }
                                                                    case "a":
                                                                        var v168 uint8 = (v143 - (v143 >> ((2) & 63)))
                                                                        _ = v168
                                                                    case "hello":
                                                                        var v169 int64 = (v4 - v4)
                                                                        _ = v169
                                                                    default:
                                                                        switch fi155 {
                                                                        case 32767:
                                                                            v170 := GenS12{F0: (-(v5)), F1: -2147483648}
                                                                            _ = v170
                                                                        }
                                                                    }
                                                                case 190:
                                                                    if (!(genFn1() < 1)) {
                                                                        var v171 float64 = (1.5 - 0.0)
                                                                        switch 4294967295 {
                                                                        case 1:
                                                                            v165 -= v142
                                                                        default:
                                                                            if v141 {
                                                                                var v172 uint32 = 4294967295
                                                                                v173 := GenS13{F0: fi155, F1: (v165 / ((v142) | 1))}
                                                                                if v141 {
                                                                                    var v174 int32 = 2147483647
                                                                                    v132 -= (float32(v171) - v132)
                                                                                    _ = v174
                                                                                }
                                                                                _ = v172
                                                                                _ = v173
                                                                            }
                                                                        }
                                                                        if (!v141) {
                                                                            var v175 int64 = -368
                                                                            _ = v175
                                                                        } else {
                                                                            v176 := GenS1{F0: 2147483647}
                                                                            v1[2147483648] = ("hello" + ("hello" + "hello"))
                                                                            _ = v176
                                                                        }
                                                                        _ = v171
                                                                    }
                                                                }
                                                                _ = len(v2)
                                                                var v177 bool = v141
                                                                _ = v177
                                                            case -383:
                                                                var v178 int16 = int16(v4)
                                                                _ = v178
                                                            case 128:
                                                                if v141 {
                                                                    v179 := GenS14{F0: v141}
                                                                    _ = v179
                                                                } else {
                                                                    if (!v141) {
                                                                        v180 := GenS3{F0: 127}
                                                                        v5 = (v5 ^ -1)
                                                                        var v181 int16 = v3
                                                                        v182 := GenS1{F0: 1}
                                                                        _ = v180
                                                                        _ = v181
                                                                        _ = v182
                                                                    } else {
                                                                        var v183 int16 = 1000
                                                                        _ = v183
                                                                    }
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        delete(v2, "foo bar")
                                                                    case "\n\t":
                                                                        v4 -= v4
                                                                    default:
                                                                        fmt.Println(v143, v3, v132)
                                                                    }
                                                                }
                                                                var v184 float64 = 1e10
                                                                v4 -= v4
                                                                _ = v184
                                                            }
                                                        case 324:
                                                            if v141 {
                                                                if (3.14159 > -2.25) {
                                                                    v143 = (v143 / ((100) | 1))
                                                                } else {
                                                                    _ = len(v1)
                                                                    if (-2.25 == v132) {
                                                                        v142 -= v142
                                                                    } else {
                                                                        if (1 <= 32767) {
                                                                            var v185 uint8 = 2
                                                                            fmt.Println(v142, v3, v141)
                                                                            v186 := GenS3{F0: v185}
                                                                            _ = v185
                                                                            _ = v186
                                                                        } else {
                                                                            v143 -= 128
                                                                        }
                                                                        var v187 int32 = v165
                                                                        _ = v187
                                                                    }
                                                                    v165 += v165
                                                                }
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    var v188 uint16 = 2
                                                                    _ = v188
                                                                case "":
                                                                    fmt.Println(v143, v165)
                                                                default:
                                                                    if (v4 >= (v4 | v4)) {
                                                                        v189 := GenS14{F0: v141}
                                                                        switch (("a" + "hello") + "foo bar") {
                                                                        case "a":
                                                                            var v190 uint64 = 612
                                                                            v132 = v132
                                                                            _ = v190
                                                                        case "foo bar":
                                                                            var v191 int8 = v5
                                                                            _ = v191
                                                                        case "hello":
                                                                            v143 = v143
                                                                        default:
                                                                            var v192 string = (("hello" + "hello") + ("hello" + ("foo bar" + "\n\t")))
                                                                            _ = v192
                                                                        }
                                                                        delete(v2, "")
                                                                        v5 -= v5
                                                                        _ = v189
                                                                    } else {
                                                                        v132 += v132
                                                                        switch (fi155 >> ((351) & 63)) {
                                                                        case 9223372036854775807:
                                                                            var v193 int8 = (-(v5))
                                                                            _ = v193
                                                                        case 2147483647:
                                                                            var v194 int32 = 42
                                                                            _ = v194
                                                                        case 128:
                                                                            v195 := GenS4{F0: (!(!false)), F1: 894}
                                                                            _ = v195
                                                                        }
                                                                        if (!(259 < uint64(v143))) {
                                                                            var v196 int8 = v5
                                                                            _ = v196
                                                                        } else {
                                                                            if ((fi155 - fi152) < (fi155 + fi155)) {
                                                                                if v141 {
                                                                                    if false {
                                                                                        var v197 uint64 = 4294967295
                                                                                        fmt.Println(v143, v5, v197)
                                                                                        _ = v197
                                                                                    }
                                                                                    var v198 string = "\n\t"
                                                                                    var v199 int16 = v3
                                                                                    fmt.Println(v4)
                                                                                    v141 = v141
                                                                                    _ = v198
                                                                                    _ = v199
                                                                                }
                                                                                var v200 uint16 = genFn1()
                                                                                _ = v200
                                                                            }
                                                                            fmt.Println(fi155, v3)
                                                                            if v141 {
                                                                                v142 = (v142 >> ((42) & 63))
                                                                            }
                                                                        }
                                                                        if (-616 < 32767) {
                                                                            v201 := GenS9{F0: (v5 / (((^(v5))) | 1))}
                                                                            _ = v201
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        default:
                                                            if (((!true) || ("a" < "hello")) && (505 <= 18446744073709551615)) {
                                                                v132 -= 3.14159
                                                                var v202 uint64 = 550
                                                                var v203 float32 = (float32(((float64((v132 - (v132 - 0.0))) / 0.1) + -2.25)) - 0.1)
                                                                _ = v202
                                                                _ = v203
                                                            }
                                                        }
                                                    }
                                                case "foo bar":
                                                    v142 -= (v165 % ((v142) | 1))
                                                }
                                                _ = v164
                                                _ = v165
                                            }
                                        }
                                        v132 = v132
                                        _ = fi155
                                        _ = v156
                                    }
                                }
                            }
                        }
                    } else {
                        for fi204 := 0; fi204 < 16; fi204++ {
                            v132 -= (((0.1 + 0.1) + v132) + (v132 / 1.5))
                            if v141 {
                                if true {
                                    switch (fi204 % ((fi152) | 1)) {
                                    case 32768:
                                        var v205 uint = 42
                                        _ = v205
                                    case 128:
                                        v132 = (v132 * v132)
                                    default:
                                        if v141 {
                                            fmt.Println(v3, fi152, v143)
                                        }
                                    }
                                    var v206 float32 = -2.25
                                    _ = v206
                                }
                                var v207 uint64 = 2
                                fmt.Println(v143)
                                _ = v207
                            } else {
                                v3 += -32768
                            }
                            _ = fi204
                        }
                    }
                case 2147483647:
                    if (v3 < v3) {
                        var v208 uint = 201
                        _ = v208
                    }
                }
                _ = fi152
            }
            for fi209 := 0; fi209 < 10; fi209++ {
                for fi210 := 0; fi210 < 11; fi210++ {
                    if true {
                        v211 := GenS2{F0: v141, F1: (-(v142)), F2: fi209}
                        var v212 uint8 = v143
                        var v213 uint = 813
                        _ = v211
                        _ = v212
                        _ = v213
                    }
                    _ = fi210
                }
                switch fi209 {
                case 0:
                    var v214 uint8 = (v143 * v143)
                    _ = v214
                default:
                    v215 := GenS6{F0: genFn1()}
                    _ = v215
                }
                _ = fi209
            }
            if (641 <= 267) {
                fmt.Println(v3, v5)
                v4 -= -9223372036854775808
            }
        } else {
            switch 128 {
            case 32767:
                if v141 {
                    v142 = -1
                    var v216 uint64 = 4294967295
                    _ = v216
                }
            case 32768:
                v132 += float32(3.14159)
            default:
                if v141 {
                    v217 := GenS9{F0: -128}
                    _ = v217
                } else {
                    if (32767 > int(v143)) {
                        v218 := []int{-769, -425}
                        var v219 float64 = 3.14159
                        _ = v218
                        _ = v219
                    }
                    v143 += (v143 + 2)
                    switch -629 {
                    case 0:
                        v220 := GenS13{F0: 698, F1: (^(v142))}
                        _ = v220
                    default:
                        var v221 uint = 174
                        _ = v221
                    }
                }
            }
            var v222 uint32 = 0
            _ = v222
        }
    }
    switch "hello" {
    case "":
        switch "\n\t" {
        case "hello":
            fmt.Println(v5)
        case "a":
            v223 := GenS1{F0: 65535}
            _ = v223
        case "foo bar":
            if (1 != 42) {
                var v224 string = "\n\t"
                switch 2147483647 {
                case -814:
                    fmt.Println(v3, v3, v3)
                case 32767:
                    switch "\n\t" {
                    case "hello":
                        v4 += -9223372036854775808
                    case "\n\t":
                        for fi225 := 0; fi225 < 4; fi225++ {
                            var v226 uint = 255
                            _ = fi225
                            _ = v226
                        }
                    case "foo bar":
                        v3 = 42
                    default:
                        v227 := GenS1{F0: 2147483648}
                        _ = v227
                    }
                default:
                    v228 := []int{0}
                    _ = v228
                }
                _ = v224
            } else {
                if v141 {
                    v132 = v132
                    if (v143 < 2) {
                        v3 = (v3 >> ((1) & 63))
                        if ((v3 & genFn0()) >= v3) {
                            fmt.Println(v132, v132, v5)
                        } else {
                            switch ("hello" + "foo bar") {
                            case "a":
                                v229 := []int{-540}
                                _ = v229
                            case "\n\t":
                                var v230 int = 0
                                _ = v230
                            case "hello":
                                for fi231 := 0; fi231 < 10; fi231++ {
                                    fmt.Println(fi231, v4, v142)
                                    if false {
                                        var v232 uint = 4294967295
                                        v233 := v2[("hello" + "")]
                                        v234 := GenS15{F0: v142, F1: 65535}
                                        _ = v232
                                        _ = v233
                                        _ = v234
                                    }
                                    v141 = v141
                                    _ = fi231
                                }
                            }
                            var v235 uint = 2
                            var v236 uint8 = 255
                            _ = v235
                            _ = v236
                        }
                        var v237 float32 = -2.25
                        _ = v237
                    }
                    var v238 uint32 = 42
                    _ = v238
                } else {
                    switch 32768 {
                    case 1:
                        for fi239 := 0; fi239 < 16; fi239++ {
                            v132 = 0.1
                            var v240 uint32 = 0
                            _ = fi239
                            _ = v240
                        }
                    case -434:
                        for fi241 := 0; fi241 < 19; fi241++ {
                            var v242 int16 = (-(v3))
                            fmt.Println(v143, fi241)
                            _ = fi241
                            _ = v242
                        }
                    }
                }
                delete(v2, "a")
                v3 = int16(v5)
            }
        }
    case "\n\t":
        fmt.Println(v141, v143, v3)
    }
    var v243 float64 = 1e10
    if (42 != 127) {
        v3 = v3
    }
    var v244 uint64 = 4294967295
    v142 += (v142 & 736)
    fmt.Println(v3, v142)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v132:", v132)
    fmt.Println("v141:", v141)
    fmt.Println("v142:", v142)
    fmt.Println("v143:", v143)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v243:", v243)
    fmt.Println("v244:", v244)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v132
    _ = v141
    _ = v142
    _ = v143
    _ = v2
    _ = v243
    _ = v244
    _ = v3
    _ = v4
    _ = v5
}
