package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS3 struct { F0 int; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint; F1 uint16; F2 int16 }
type GenS5 struct { F0 string; F1 int16; F2 int16 }
type GenS6 struct { F0 uint; F1 int32 }
type GenS7 struct { F0 uint16; F1 int16 }
type GenS8 struct { F0 uint64; F1 string; F2 uint16 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int8; F1 int16; F2 bool }
type GenS11 struct { F0 int; F1 uint32 }
type GenS12 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS13 struct { F0 uint16; F1 uint32; F2 int }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 string; F1 int }
type GenS16 struct { F0 uint32 }
type GenS17 struct { F0 string }
type GenS18 struct { F0 int64; F1 uint16 }
type GenS19 struct { F0 uint32 }
type GenS20 struct { F0 uint32; F1 string; F2 uint64 }
type GenS21 struct { F0 bool }
type GenS22 struct { F0 int32; F1 bool }
type GenS23 struct { F0 uint16; F1 int }
type GenS24 struct { F0 uint32 }
type GenS25 struct { F0 int16; F1 string; F2 string }
func genFn0(p0 int32, p1 uint8, p2 string) uint32 {
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 float64 = (0.1 * 0.1)
        v7 := []int{(fi5 + (fi5 * (-(fi5)))), (fi5 << ((42) & 63))}
        _ = fi5
        _ = v6
        _ = v7
    }
    p0 = (^(p0))
    var v8 uint32 = 42
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 uint = 4294967295
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 2, F1: 42, F2: 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v8
    return (v8 % ((v8) | 1))
}
func genFn1() uint8 {
    var v12 int64 = 0
    v12 = (v12 ^ -9223372036854775808)
    for fi13 := 0; fi13 < 9; fi13++ {
        fmt.Println(v12, fi13)
        _ = fi13
    }
    v14 := GenS2{F0: 2, F1: 2, F2: -1}
    v15 := []int{669, -793, 32767, -2147483648}
    _ = v12
    _ = v14
    _ = v15
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 1
    v3 := make(map[int]string)
    v4 := make(map[string]bool)
    switch (v2 >> ((772) & 63)) {
    case -1:
        switch "foo bar" {
        case "\n\t":
            delete(v4, "")
        case "foo bar":
            v16 := GenS2{F0: 42, F1: 454, F2: -2147483648}
            _ = v16
        }
    }
    v2 -= (v2 ^ v2)
    var v17 float64 = float64((((float32(((3.14159 - 1e10) / -2.25)) + float32(0.1)) + -2.25) * 1e10))
    switch (v2 - v2) {
    case 32768:
        var v18 float32 = -2.25
        _ = v18
    case 4294967295:
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 uint32 = 735
            switch ("hello" + (("" + (("hello" + "") + "foo bar")) + (("\n\t" + "foo bar") + ("" + "\n\t")))) {
            case "":
                _ = len(v4)
            case "\n\t":
                if (-2147483648 < 97) {
                    v2 += fi19
                    v21 := GenS1{F0: 320, F1: 20, F2: 127}
                    var v22 uint64 = 2
                    _ = v21
                    _ = v22
                }
            default:
                v17 = v17
            }
            _ = fi19
            _ = v20
        }
    case 1:
        v2 -= (v2 << ((705) & 63))
    default:
        var v23 int16 = -817
        _ = v23
    }
    switch "a" {
    case "a":
        var v24 int16 = 42
        _ = v24
    case "hello":
        v2 = v2
    default:
        for fi25 := 0; fi25 < 1; fi25++ {
            if (42 == 49) {
                if ((false || (false || false)) || ((((3.14159 / (1.5 * 1e10)) + (1e10 + 1.5)) * 0.1) > 0.0)) {
                    switch v2 {
                    case 156:
                        var v26 int32 = 42
                        _ = v26
                    case 4294967295:
                        if true {
                            fmt.Println(v17, fi25, v2)
                            v2 -= 32768
                            var v27 uint = 18446744073709551615
                            _ = v27
                        }
                    default:
                        v28 := GenS3{F0: 2147483647, F1: true, F2: 255}
                        _ = v28
                    }
                    switch "" {
                    case "hello":
                        for fi29 := 0; fi29 < 5; fi29++ {
                            _ = len(v3)
                            _ = fi29
                        }
                    case "\n\t":
                        fmt.Println(fi25, fi25, v17)
                    case "":
                        switch (fi25 & fi25) {
                        case 703:
                            for fi30 := 0; fi30 < 5; fi30++ {
                                if (-1 == -9223372036854775808) {
                                    v31 := v3[(fi30 << ((840) & 63))]
                                    v32 := GenS4{F0: 4294967295, F1: 0, F2: -32768}
                                    _ = len(v1)
                                    _ = v31
                                    _ = v32
                                }
                                var v33 int16 = 1000
                                fmt.Println(fi25, fi30, fi30)
                                _ = fi30
                                _ = v33
                            }
                        case -2147483648:
                            fmt.Println(v17)
                        }
                    default:
                        v34 := []int{(-(v2)), 42}
                        _ = v34
                    }
                } else {
                    v35 := v4[""]
                    if (!(-356 < -143)) {
                        switch (fi25 >> ((18446744073709551615) & 63)) {
                        case -239:
                            switch ("a" + "") {
                            case "\n\t":
                                var v36 uint16 = 1
                                _ = v36
                            case "hello":
                                v2 = fi25
                            }
                        case -959:
                            if (v35 && v35) {
                                v35 = true
                                for fi37 := 0; fi37 < 11; fi37++ {
                                    v38 := GenS5{F0: "foo bar", F1: 178, F2: -32768}
                                    v2 += v2
                                    v39 := GenS4{F0: 2, F1: 1, F2: -160}
                                    _ = fi37
                                    _ = v38
                                    _ = v39
                                }
                                v35 = true
                            }
                        case 32767:
                            v17 -= ((v17 * v17) + float64(-2.25))
                        }
                        for fi40 := 0; fi40 < 3; fi40++ {
                            v17 = v17
                            fmt.Println(fi40)
                            if ((!(v35 || (!v35))) || true) {
                                v41 := []int{(fi25 + (fi25 >> ((2) & 63)))}
                                _ = v41
                            } else {
                                var v42 uint64 = 690
                                v42 = 2
                                _ = v42
                            }
                            var v43 uint16 = 94
                            _ = fi40
                            _ = v43
                        }
                    } else {
                        v2 -= (v2 * v2)
                        if v35 {
                            if (!true) {
                                var v44 uint8 = genFn1()
                                if v35 {
                                    if false {
                                        v35 = v35
                                        var v45 uint16 = 759
                                        for fi46 := 0; fi46 < 6; fi46++ {
                                            v47 := []int{}
                                            delete(v4, "\n\t")
                                            _ = fi46
                                            _ = v47
                                        }
                                        _ = v45
                                    } else {
                                        v48 := []int{}
                                        for fi49 := 0; fi49 < 15; fi49++ {
                                            if len(v48) > 0 {
                                                v48[0] = fi25
                                            }
                                            v2 = (fi49 & fi49)
                                            _ = fi49
                                        }
                                        _ = v48
                                    }
                                    switch "\n\t" {
                                    case "a":
                                        if false {
                                            v35 = (genFn0(633, 73, "\n\t") >= 65535)
                                            var v50 float32 = (3.14159 / (1e10 / (1.5 * (0.0 + 0.1))))
                                            v44 += (v44 % ((genFn1()) | 1))
                                            _ = v50
                                        }
                                    default:
                                        var v51 uint16 = 65535
                                        _ = v51
                                    }
                                    var v52 int64 = int64(v2)
                                    _ = v52
                                }
                                fmt.Println(v44, v2, v2)
                                _ = v44
                            } else {
                                if (v35 || ((!v35) && v35)) {
                                    v35 = ((1000 == 0) || (18446744073709551615 > 1))
                                    if (v35 && (0 > 0)) {
                                        fmt.Println(v2, v17)
                                        fmt.Println(v2)
                                        v2 += fi25
                                    } else {
                                        fmt.Println(v35, v35, fi25)
                                        var v53 int16 = -32768
                                        _ = v53
                                    }
                                } else {
                                    if v35 {
                                        v17 = (v17 + v17)
                                        var v54 uint16 = 65535
                                        _ = v54
                                    }
                                }
                                switch ("\n\t" + "hello") {
                                case "a":
                                    for fi55 := 0; fi55 < 12; fi55++ {
                                        var v56 uint = 18446744073709551615
                                        _ = fi55
                                        _ = v56
                                    }
                                }
                                v57 := GenS1{F0: 65535, F1: 0, F2: 100}
                                _ = v57
                            }
                            v17 = float64((0.1 - 3.14159))
                            var v58 uint8 = 0
                            _ = v58
                        } else {
                            if true {
                                v17 += float64(1e10)
                            } else {
                                var v59 uint16 = 65535
                                _ = v59
                            }
                            if (0.1 <= -2.25) {
                                if v35 {
                                    v60 := GenS1{F0: 0, F1: 42, F2: 1}
                                    v61 := []int{v2}
                                    v62 := v3[v2]
                                    _ = v60
                                    _ = v61
                                    _ = v62
                                }
                                var v63 int64 = 1
                                _ = v63
                            } else {
                                v3[(v2 - 42)] = "a"
                                v3[(fi25 - (v2 << ((0) & 63)))] = (("hello" + "\n\t") + "hello")
                                var v64 float32 = 1.5
                                _ = v64
                            }
                        }
                    }
                    _ = v35
                }
                v2 = fi25
                var v65 bool = false
                _ = v65
            }
            if (true || (0 != genFn1())) {
                if false {
                    var v66 int16 = 42
                    _ = v66
                } else {
                    switch (fi25 + v2) {
                    case 416:
                        v2 += fi25
                    default:
                        if (!false) {
                            for fi67 := 0; fi67 < 8; fi67++ {
                                if false {
                                    var v68 int64 = 42
                                    if ((float64((float32((v17 - v17)) / 1.0)) * 0.1) != v17) {
                                        v69 := GenS6{F0: 65535, F1: 42}
                                        v70 := []int{9223372036854775807, fi67, fi25, (^(fi67))}
                                        _ = len(v70)
                                        _ = v69
                                        _ = v70
                                    } else {
                                        if true {
                                            fmt.Println(fi25, fi67)
                                            var v71 int64 = (v68 << ((42) & 63))
                                            _ = v71
                                        } else {
                                            fmt.Println(fi25)
                                            switch "a" {
                                            case "\n\t":
                                                if (((0.0 - -2.25) != -2.25) && true) {
                                                    var v72 uint32 = genFn0(42, 188, "")
                                                    v73 := v3[v2]
                                                    v2 += (fi25 + -474)
                                                    _ = v72
                                                    _ = v73
                                                }
                                            default:
                                                v74 := v3[(fi67 * 4294967295)]
                                                _ = v74
                                            }
                                            v75 := []int{(fi67 - (-(fi25))), (^(fi67)), -2147483648}
                                            v17 -= v17
                                            var v76 int16 = 1
                                            v17 += (v17 / (v17 / -2.25))
                                            if (65535 == 65535) {
                                                v2 += (fi67 ^ (v2 >> ((18446744073709551615) & 63)))
                                                fmt.Println(fi25, fi25, v2)
                                                if true {
                                                    switch fi25 {
                                                    case -1:
                                                        v76 = (v76 << ((347) & 63))
                                                        var v77 uint16 = 438
                                                        v75 = append(v75, 32767)
                                                        _ = v77
                                                    default:
                                                        v78 := GenS6{F0: 1, F1: -1}
                                                        _ = v78
                                                    }
                                                    v79 := GenS3{F0: -27, F1: true, F2: genFn1()}
                                                    var v80 int32 = 2147483647
                                                    _ = v79
                                                    _ = v80
                                                }
                                            } else {
                                                v81 := GenS7{F0: 255, F1: (v76 % ((v76) | 1))}
                                                fmt.Println(v2)
                                                _ = v81
                                            }
                                            v68 += (v68 << ((4294967295) & 63))
                                            v17 += v17
                                            _ = v75
                                            _ = v76
                                        }
                                    }
                                    _ = v68
                                } else {
                                    v2 = (fi25 * fi67)
                                    v17 += float64(0.1)
                                    v82 := GenS8{F0: 1, F1: "\n\t", F2: 2}
                                    _ = v82
                                }
                                _ = fi67
                            }
                            var v83 uint = 2
                            v1[-171] = genFn0(42, 127, "foo bar")
                            _ = v83
                        }
                    }
                }
                if true {
                    if ("a" <= "\n\t") {
                        v84 := GenS4{F0: 1, F1: 675, F2: 669}
                        var v85 float64 = (1.5 * 1.5)
                        _ = v84
                        _ = v85
                    } else {
                        for fi86 := 0; fi86 < 16; fi86++ {
                            if (0 >= -281) {
                                _ = len(v3)
                            } else {
                                var v87 int64 = -336
                                _ = v87
                            }
                            _ = fi86
                        }
                    }
                    v88 := []int{v2, v2, (v2 * (fi25 | (fi25 | v2))), (v2 / ((v2) | 1))}
                    if true {
                        fmt.Println(v17)
                        if true {
                            var v89 int32 = 1
                            switch -765 {
                            case -1:
                                for fi90 := 0; fi90 < 10; fi90++ {
                                    switch "hello" {
                                    case "":
                                        fmt.Println(fi90, fi25)
                                    default:
                                        switch fi25 {
                                        case 0:
                                            switch "hello" {
                                            case "foo bar":
                                                var v91 int = (v2 | -9223372036854775808)
                                                _ = v91
                                            case "a":
                                                switch "foo bar" {
                                                case "a":
                                                    switch v2 {
                                                    case 0:
                                                        switch ((("hello" + "hello") + ("" + (("\n\t" + "") + ((("foo bar" + "hello") + "a") + "a")))) + ("a" + "foo bar")) {
                                                        case "":
                                                            v92 := GenS6{F0: 65535, F1: 1}
                                                            _ = v92
                                                        case "foo bar":
                                                            v2 -= v2
                                                        default:
                                                            v1[(v2 - v2)] = 65535
                                                        }
                                                    case -2147483648:
                                                        v2 -= (fi90 / (((v2 * (fi25 | fi90))) | 1))
                                                    }
                                                case "\n\t":
                                                    _ = len(v88)
                                                }
                                            case "\n\t":
                                                v93 := GenS9{F0: 0}
                                                _ = v93
                                            }
                                        default:
                                            v17 = -2.25
                                        }
                                    }
                                    _ = fi90
                                }
                            case 594:
                                var v94 int = -2147483648
                                _ = v94
                            case 4294967295:
                                fmt.Println(fi25)
                            default:
                                v95 := GenS10{F0: 86, F1: 0, F2: true}
                                _ = v95
                            }
                            _ = v89
                        }
                        var v96 uint8 = 0
                        _ = v96
                    } else {
                        v97 := GenS5{F0: "", F1: 32767, F2: -32768}
                        var v98 bool = (!(53 != 171))
                        _ = v97
                        _ = v98
                    }
                    _ = v88
                }
                if (!false) {
                    if false {
                        v17 = float64(3.14159)
                        for fi99 := 0; fi99 < 16; fi99++ {
                            if ((65535 != 143) || (!(true && false))) {
                                if ((float32(1.5) != float32(v17)) && false) {
                                    switch "" {
                                    case "foo bar":
                                        v100 := GenS11{F0: 9223372036854775807, F1: 42}
                                        _ = v100
                                    case "a":
                                        v101 := GenS12{F0: 480, F1: 65535, F2: -499}
                                        _ = v101
                                    case "\n\t":
                                        v102 := GenS13{F0: 286, F1: 2147483647, F2: (^(v2))}
                                        _ = v102
                                    }
                                } else {
                                    if ((1 == 42) && (!(255 != 4294967295))) {
                                        v1[-819] = 672
                                        var v103 float32 = (((-2.25 * ((3.14159 - float32(v17)) + ((0.1 + 3.14159) * 1.5))) - float32((v17 * (v17 * -2.25)))) - 1.5)
                                        switch (fi99 | -2147483648) {
                                        case 9223372036854775807:
                                            var v104 float64 = float64(v103)
                                            _ = v104
                                        case -581:
                                            switch v2 {
                                            case -491:
                                                if false {
                                                    v105 := GenS14{F0: -1}
                                                    var v106 float32 = v103
                                                    _ = v105
                                                    _ = v106
                                                } else {
                                                    var v107 bool = true
                                                    var v108 float32 = 3.14159
                                                    var v109 int16 = 42
                                                    _ = v107
                                                    _ = v108
                                                    _ = v109
                                                }
                                            case 127:
                                                v110 := GenS15{F0: "hello", F1: -133}
                                                _ = v110
                                            default:
                                                v2 = (v2 * fi25)
                                            }
                                        default:
                                            v111 := v4[("" + (("a" + "hello") + (("a" + ("\n\t" + "a")) + (("\n\t" + "foo bar") + "a"))))]
                                            _ = v111
                                        }
                                        switch v2 {
                                        case -1:
                                            switch (-(fi25)) {
                                            case 564:
                                                v112 := GenS14{F0: 100}
                                                _ = v112
                                            case 1:
                                                var v113 float64 = (float64(v103) / 1.0)
                                                _ = v113
                                            default:
                                                v114 := GenS16{F0: 95}
                                                _ = v114
                                            }
                                        default:
                                            v2 += v2
                                        }
                                        if false {
                                            if (1000 == 749) {
                                                v103 -= 0.1
                                                v115 := GenS13{F0: 2, F1: 2147483648, F2: -1}
                                                v17 = (v17 + 0.1)
                                                _ = v115
                                            } else {
                                                if true {
                                                    if (1e10 < ((v17 + v17) + (v17 - v17))) {
                                                        v2 -= (v2 & (v2 - fi99))
                                                        var v116 bool = false
                                                        var v117 string = "a"
                                                        _ = v116
                                                        _ = v117
                                                    } else {
                                                        var v118 uint16 = 554
                                                        v119 := GenS17{F0: ("foo bar" + (("foo bar" + ("\n\t" + "\n\t")) + "a"))}
                                                        _ = v118
                                                        _ = v119
                                                    }
                                                    fmt.Println(v103, v2)
                                                    v2 = (-(fi25))
                                                }
                                                if false {
                                                    var v120 int16 = -32768
                                                    _ = v120
                                                } else {
                                                    v103 = 3.14159
                                                }
                                            }
                                            var v121 int8 = 100
                                            if (v121 == (v121 ^ (v121 ^ v121))) {
                                                var v122 string = "foo bar"
                                                if true {
                                                    v123 := GenS1{F0: 255, F1: (v121 >> ((65535) & 63)), F2: (v121 & v121)}
                                                    _ = v123
                                                }
                                                v124 := []int{1}
                                                _ = v122
                                                _ = v124
                                            }
                                            var v125 string = "foo bar"
                                            _ = v121
                                            _ = v125
                                        } else {
                                            switch "" {
                                            case "hello":
                                                if (!true) {
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v126 uint16 = 0
                                                        _ = v126
                                                    case "a":
                                                        v17 += v17
                                                    case "":
                                                        if (2 < 1) {
                                                            fmt.Println(fi99, fi99, v17)
                                                        } else {
                                                            v127 := GenS6{F0: 0, F1: -776}
                                                            var v128 int64 = 915
                                                            _ = v127
                                                            _ = v128
                                                        }
                                                        var v129 uint = 548
                                                        _ = v129
                                                    default:
                                                        var v130 uint64 = 2
                                                        _ = v130
                                                    }
                                                }
                                            }
                                        }
                                        _ = v103
                                    } else {
                                        v131 := []int{}
                                        _ = v131
                                    }
                                }
                                var v132 uint16 = 255
                                _ = v132
                            } else {
                                fmt.Println(fi25)
                                switch ("hello" + "") {
                                case "\n\t":
                                    if true {
                                        v2 = (fi99 ^ (v2 << ((1) & 63)))
                                        v2 = fi25
                                    } else {
                                        var v133 int64 = 42
                                        var v134 uint16 = 2
                                        switch (v2 + 4294967295) {
                                        case 42:
                                            if false {
                                                var v135 uint16 = 415
                                                _ = v135
                                            }
                                        default:
                                            v136 := GenS13{F0: (v134 + (v134 ^ 2)), F1: genFn0(863, 127, (("hello" + "") + "foo bar")), F2: (fi25 / ((1) | 1))}
                                            _ = v136
                                        }
                                        _ = v133
                                        _ = v134
                                    }
                                case "a":
                                    delete(v4, "foo bar")
                                case "hello":
                                    var v137 int8 = 0
                                    _ = v137
                                default:
                                    v2 -= fi25
                                }
                            }
                            if true {
                                switch -423 {
                                case 2147483647:
                                    v138 := []int{127, fi25}
                                    _ = v138
                                case 4294967295:
                                    if ((v2 <= 550) && false) {
                                        var v139 uint32 = 4294967295
                                        var v140 uint8 = 161
                                        if (255 != 2) {
                                            switch ("a" + "\n\t") {
                                            case "foo bar":
                                                var v141 uint64 = 4294967295
                                                _ = v141
                                            case "\n\t":
                                                v142 := GenS18{F0: 0, F1: 2}
                                                _ = v142
                                            default:
                                                v143 := []int{(fi99 >> ((255) & 63)), 1, (fi25 * -2147483648)}
                                                v144 := GenS15{F0: ("" + "a"), F1: (fi99 + v2)}
                                                v140 -= (v140 << ((658) & 63))
                                                _ = v143
                                                _ = v144
                                            }
                                            fmt.Println(v2)
                                            if (0 != 18446744073709551615) {
                                                v17 = (0.0 * ((3.14159 * 1.5) / float64((float32(v17) * 3.14159))))
                                                fmt.Println(fi99, v140, v17)
                                                v145 := v4["hello"]
                                                _ = v145
                                            } else {
                                                if (true || true) {
                                                    var v146 bool = true
                                                    switch "a" {
                                                    case "\n\t":
                                                        fmt.Println(v2)
                                                    case "":
                                                        v147 := GenS19{F0: (v139 * (v139 + v139))}
                                                        _ = v147
                                                    case "a":
                                                        var v148 float32 = 3.14159
                                                        var v149 string = ("" + "a")
                                                        _ = v148
                                                        _ = v149
                                                    }
                                                    delete(v1, (^(v2)))
                                                    _ = v146
                                                }
                                            }
                                        } else {
                                            if false {
                                                v150 := []int{fi25, v2}
                                                _ = v150
                                            }
                                            switch "\n\t" {
                                            case "foo bar":
                                                v4["a"] = ((!(true || (967 != 2))) && ((!false) || false))
                                            default:
                                                v2 += (v2 / (((v2 << ((255) & 63))) | 1))
                                            }
                                            v151 := GenS20{F0: 372, F1: "\n\t", F2: 42}
                                            _ = v151
                                        }
                                        if false {
                                            switch 1 {
                                            case -823:
                                                var v152 uint64 = 231
                                                _ = v152
                                            case 908:
                                                if ((((false || false) && (464 == -971)) && ((!false) && true)) && false) {
                                                    _ = len(v3)
                                                    var v153 bool = true
                                                    _ = v153
                                                }
                                            default:
                                                if (1 == 2147483647) {
                                                    v154 := v1[v2]
                                                    var v155 int16 = -714
                                                    var v156 uint = 86
                                                    _ = v154
                                                    _ = v155
                                                    _ = v156
                                                }
                                            }
                                            v157 := GenS11{F0: v2, F1: v139}
                                            _ = v157
                                        } else {
                                            delete(v1, 32768)
                                            var v158 uint8 = 186
                                            fmt.Println(v17)
                                            _ = v158
                                        }
                                        _ = v139
                                        _ = v140
                                    }
                                default:
                                    if true {
                                        v159 := GenS4{F0: 4294967295, F1: 720, F2: 523}
                                        _ = v159
                                    }
                                }
                                var v160 float32 = (-2.25 / 3.14159)
                                if (false || (false || false)) {
                                    if false {
                                        switch "\n\t" {
                                        case "hello":
                                            var v161 float32 = (v160 + 1e10)
                                            _ = len(v1)
                                            _ = v161
                                        default:
                                            var v162 string = "a"
                                            _ = v162
                                        }
                                        v163 := GenS21{F0: true}
                                        var v164 uint64 = 18446744073709551615
                                        _ = v163
                                        _ = v164
                                    } else {
                                        fmt.Println(v160, v17)
                                        v165 := GenS22{F0: -1, F1: false}
                                        _ = v165
                                    }
                                    var v166 float64 = (v17 * ((v17 * v17) * v17))
                                    v17 += 0.1
                                    _ = v166
                                } else {
                                    var v167 uint = 2
                                    switch fi99 {
                                    case 197:
                                        v168 := GenS1{F0: 1, F1: -87, F2: 0}
                                        _ = v168
                                    case 1:
                                        if (!false) {
                                            var v169 int32 = 0
                                            if ((-41 > -128) && ((0 > 32767) && false)) {
                                                switch v2 {
                                                case -191:
                                                    var v170 uint16 = 55
                                                    _ = v170
                                                case 1:
                                                    switch "foo bar" {
                                                    case "a":
                                                        var v171 bool = (0 > -351)
                                                        _ = v171
                                                    case "hello":
                                                        var v172 int64 = -1
                                                        _ = v172
                                                    case "foo bar":
                                                        v173 := GenS12{F0: 42, F1: 42, F2: v169}
                                                        v167 += (v167 ^ (^(v167)))
                                                        var v174 uint32 = 0
                                                        var v175 string = ""
                                                        v17 = -2.25
                                                        _ = len(v1)
                                                        v4[v175] = (-21 <= -80)
                                                        _ = v173
                                                        _ = v174
                                                        _ = v175
                                                    }
                                                case 128:
                                                    if (!true) {
                                                        v4[("foo bar" + "a")] = false
                                                        var v176 uint64 = 1
                                                        _ = v176
                                                    }
                                                }
                                                v177 := GenS22{F0: (v169 | 1), F1: (v160 == 3.14159)}
                                                v178 := GenS11{F0: fi25, F1: genFn0(0, 2, "\n\t")}
                                                _ = v177
                                                _ = v178
                                            } else {
                                                v179 := []int{}
                                                var v180 int = -467
                                                switch fi25 {
                                                case -247:
                                                    fmt.Println(v160, fi99, v17)
                                                case 32767:
                                                    if false {
                                                        var v181 uint16 = 308
                                                        var v182 uint16 = 32767
                                                        var v183 int16 = -1
                                                        _ = v181
                                                        _ = v182
                                                        _ = v183
                                                    } else {
                                                        var v184 uint32 = 2
                                                        v185 := GenS23{F0: 874, F1: fi99}
                                                        v160 = v160
                                                        _ = v184
                                                        _ = v185
                                                    }
                                                default:
                                                    switch "" {
                                                    case "a":
                                                        if len(v179) > 0 {
                                                            v179[0] = fi99
                                                        }
                                                    default:
                                                        switch (v180 >> ((v167) & 63)) {
                                                        case -1:
                                                            if false {
                                                                v17 = float64(1.5)
                                                            } else {
                                                                if ((370 > 1) && ((v17 + ((v17 * (1.5 - float64((v160 - v160)))) - 0.1)) != 1.5)) {
                                                                    var v186 int16 = -1
                                                                    _ = v186
                                                                }
                                                            }
                                                        case 9223372036854775807:
                                                            v187 := GenS11{F0: v2, F1: 684}
                                                            _ = v187
                                                        case 32768:
                                                            v188 := GenS24{F0: 0}
                                                            _ = v188
                                                        }
                                                    }
                                                }
                                                _ = v179
                                                _ = v180
                                            }
                                            v189 := GenS25{F0: 132, F1: "hello", F2: ("" + "hello")}
                                            _ = v169
                                            _ = v189
                                        }
                                    default:
                                        fmt.Println(v160, v167, v167)
                                    }
                                    _ = v167
                                }
                                fmt.Println(v17, v2, v17)
                                v190 := []int{fi25, fi99, 700}
                                var v191 uint64 = 42
                                _ = v160
                                _ = v190
                                _ = v191
                            }
                            _ = fi99
                        }
                        fmt.Println(v2, v2, v2)
                    }
                    var v192 string = "\n\t"
                    _ = v192
                }
            } else {
                delete(v3, (v2 & v2))
            }
            _ = fi25
        }
    }
    v17 -= v17
    var v193 int64 = 651
    v3[0] = ("foo bar" + "")
    v194 := []int{(v2 * 9223372036854775807), (v2 * v2), (v2 / (((v2 ^ 75)) | 1)), (v2 + v2)}
    fmt.Println(v17, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v193:", v193)
    fmt.Println("slicelen v194:", len(v194), v194)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v193
    _ = v194
    _ = v2
    _ = v3
    _ = v4
}
