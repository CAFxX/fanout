package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 uint32; F1 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint64; F1 int32 }
type GenS8 struct { F0 uint; F1 uint16; F2 int8 }
type GenS9 struct { F0 int; F1 string }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int8; F1 int; F2 bool }
type GenS12 struct { F0 int64; F1 string }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int64 }
func genFn0(p0 int64) bool {
    for fi5 := 0; fi5 < 17; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            p0 = (p0 + 1)
            _ = fi6
        }
        v7 := GenS1{F0: (421 <= 42)}
        var v8 uint32 = 2147483647
        _ = fi5
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        p0 = (p0 << ((542) & 63))
        _ = fi9
    }
    v10 := []int{573, 128, -1, 432}
    _ = p0
    _ = v10
    return false
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = -361
    var v3 uint = 18446744073709551615
    v4 := []int{2147483647, int(v3), 0}
    v2 -= v2
    fmt.Println(v3, v3)
    for fi11 := 0; fi11 < 20; fi11++ {
        switch fi11 {
        case 127:
            v3 -= (v3 >> (((v3 << ((v3) & 63))) & 63))
        }
        var v12 uint32 = 882
        v13 := GenS1{F0: true}
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(v2, v3)
    if ("foo bar" > "") {
        var v14 uint16 = 614
        _ = v14
    } else {
        v3 += (v3 + (v3 >> (((^(v3))) & 63)))
        switch "" {
        case "a":
            var v15 uint16 = 65535
            _ = v15
        default:
            for fi16 := 0; fi16 < 5; fi16++ {
                v3 = v3
                _ = fi16
            }
        }
        switch 128 {
        case 497:
            switch 42 {
            case 2147483647:
                v17 := GenS2{F0: (v3 / ((783) | 1)), F1: ("hello" + "a")}
                _ = v17
            case -1:
                if true {
                    switch ("" + "") {
                    case "foo bar":
                        if (!true) {
                            var v18 uint = v3
                            _ = v18
                        }
                    case "a":
                        switch "a" {
                        case "":
                            v19 := GenS1{F0: false}
                            _ = v19
                        default:
                            var v20 uint8 = 187
                            _ = v20
                        }
                    case "":
                        v2 -= v2
                    }
                }
            default:
                for fi21 := 0; fi21 < 8; fi21++ {
                    switch (fi21 | fi21) {
                    case -1:
                        var v22 uint16 = 65535
                        _ = v22
                    case 1:
                        for fi23 := 0; fi23 < 20; fi23++ {
                            fmt.Println(v2, fi21)
                            v3 -= 4294967295
                            var v24 uint8 = 1
                            _ = fi23
                            _ = v24
                        }
                    case 757:
                        var v25 uint16 = 65535
                        _ = v25
                    }
                    v3 += 255
                    var v26 uint8 = 86
                    _ = fi21
                    _ = v26
                }
            }
        case 819:
            v2 += (v2 | v2)
        default:
            v3 = v3
        }
    }
    if false {
        if false {
            v2 -= (v2 / (((v2 - (v2 | (v2 | v2)))) | 1))
        } else {
            fmt.Println(v3, v2)
            delete(v1, 65)
        }
        switch 898 {
        case 748:
            var v27 uint16 = 255
            _ = v27
        case 2147483647:
            v2 += v2
        default:
            if ((((v2 + (-(v2))) > v2) && ((2 > 4294967295) || false)) && (1 > -490)) {
                switch "hello" {
                case "foo bar":
                    if false {
                        v28 := GenS2{F0: v3, F1: "\n\t"}
                        _ = len(v1)
                        _ = v28
                    } else {
                        var v29 uint = 315
                        if true {
                            for fi30 := 0; fi30 < 1; fi30++ {
                                fmt.Println(fi30, v3)
                                _ = fi30
                            }
                            v29 = v3
                            v3 = (v29 + (v29 | (^(v29))))
                        } else {
                            for fi31 := 0; fi31 < 6; fi31++ {
                                if (false && (!(3.14159 > 3.14159))) {
                                    var v32 int32 = 0
                                    v33 := GenS3{F0: v3, F1: 42}
                                    _ = v32
                                    _ = v33
                                }
                                _ = fi31
                            }
                            if (128 >= 199) {
                                var v34 uint = v3
                                delete(v1, 4294967295)
                                v35 := GenS4{F0: 2, F1: 937}
                                _ = v34
                                _ = v35
                            }
                            fmt.Println(v3)
                        }
                        for fi36 := 0; fi36 < 19; fi36++ {
                            if true {
                                switch ("foo bar" + "a") {
                                case "\n\t":
                                    switch (("hello" + (("hello" + "") + ("foo bar" + "\n\t"))) + "foo bar") {
                                    case "\n\t":
                                        v29 += 0
                                    case "foo bar":
                                        var v37 uint8 = 42
                                        _ = v37
                                    case "hello":
                                        v38 := v1[937]
                                        _ = v38
                                    }
                                case "hello":
                                    v4 = append(v4, (fi36 | fi36))
                                case "foo bar":
                                    v39 := GenS1{F0: true}
                                    _ = v39
                                default:
                                    if true {
                                        for fi40 := 0; fi40 < 4; fi40++ {
                                            if true {
                                                v3 = v29
                                                v3 += (v29 - (^(v29)))
                                                var v41 int32 = 2147483647
                                                _ = v41
                                            }
                                            _ = fi40
                                        }
                                        v29 = (v29 << ((4294967295) & 63))
                                        fmt.Println(v3, v29, v3)
                                    }
                                }
                                v2 += v2
                                if true {
                                    v42 := GenS5{F0: 290, F1: 0}
                                    _ = v42
                                } else {
                                    var v43 float64 = 0.0
                                    v44 := GenS4{F0: 2147483647, F1: 4294967295}
                                    _ = v43
                                    _ = v44
                                }
                            }
                            var v45 int8 = -65
                            if false {
                                var v46 int16 = -444
                                _ = v46
                            } else {
                                if true {
                                    if ((v45 - v45) != v45) {
                                        switch "hello" {
                                        case "hello":
                                            for fi47 := 0; fi47 < 1; fi47++ {
                                                fmt.Println(fi47, v45)
                                                _ = fi47
                                            }
                                        case "a":
                                            if true {
                                                var v48 uint8 = 1
                                                for fi49 := 0; fi49 < 6; fi49++ {
                                                    fmt.Println(v45)
                                                    _ = fi49
                                                }
                                                _ = v48
                                            }
                                        }
                                        for fi50 := 0; fi50 < 20; fi50++ {
                                            v51 := GenS6{F0: -9223372036854775808}
                                            var v52 uint32 = 2147483648
                                            v52 -= 327
                                            _ = fi50
                                            _ = v51
                                            _ = v52
                                        }
                                    } else {
                                        v45 += v45
                                        switch (fi36 >> (((^(v3))) & 63)) {
                                        case -1:
                                            switch "a" {
                                            case "foo bar":
                                                if (100 >= v45) {
                                                    var v53 int8 = -128
                                                    _ = v53
                                                } else {
                                                    switch int(v45) {
                                                    case 32767:
                                                        var v54 bool = true
                                                        _ = v54
                                                    case -192:
                                                        v2 = v2
                                                    default:
                                                        var v55 uint32 = 65535
                                                        _ = v55
                                                    }
                                                    var v56 uint32 = 4294967295
                                                    for fi57 := 0; fi57 < 14; fi57++ {
                                                        var v58 uint8 = 42
                                                        _ = fi57
                                                        _ = v58
                                                    }
                                                    _ = v56
                                                }
                                            case "":
                                                var v59 bool = true
                                                _ = v59
                                            case "hello":
                                                if len(v4) > 0 {
                                                    v4[0] = -699
                                                }
                                            default:
                                                if (true && ((false || (!false)) || false)) {
                                                    switch (fi36 * 32768) {
                                                    case -1:
                                                        for fi60 := 0; fi60 < 2; fi60++ {
                                                            fmt.Println(fi60, v45, v2)
                                                            _ = fi60
                                                        }
                                                    case 4294967295:
                                                        var v61 float64 = (1.5 * ((1e10 + 0.1) + 3.14159))
                                                        _ = v61
                                                    case 1:
                                                        if false {
                                                            var v62 int = fi36
                                                            v2 += (v2 | v2)
                                                            _ = v62
                                                        } else {
                                                            var v63 uint64 = 0
                                                            _ = v63
                                                        }
                                                    }
                                                    var v64 uint = (v29 >> (((v3 | (v29 >> (((v29 >> ((v3) & 63))) & 63)))) & 63))
                                                    _ = v64
                                                } else {
                                                    var v65 uint64 = 650
                                                    _ = v65
                                                }
                                            }
                                        }
                                        v3 = v3
                                    }
                                    if false {
                                        var v66 float64 = 0.1
                                        v67 := v1[827]
                                        _ = v66
                                        _ = v67
                                    } else {
                                        _ = len(v1)
                                        for fi68 := 0; fi68 < 20; fi68++ {
                                            switch fi36 {
                                            case 9223372036854775807:
                                                var v69 int = (fi68 | (-(fi36)))
                                                v45 -= (v45 / ((127) | 1))
                                                var v70 int = (fi68 >> ((42) & 63))
                                                if ((^(v45)) >= int8(fi36)) {
                                                    var v71 int16 = 42
                                                    v45 -= 0
                                                    if false {
                                                        var v72 int = (-(v69))
                                                        if (!true) {
                                                            v4 = append(v4, -866)
                                                            var v73 uint64 = 1
                                                            var v74 int64 = (^(v2))
                                                            _ = v73
                                                            _ = v74
                                                        } else {
                                                            if false {
                                                                fmt.Println(fi68, fi36, v69)
                                                            } else {
                                                                if (((true && true) || (true || false)) || true) {
                                                                    v29 += v3
                                                                    if (1 != 2147483647) {
                                                                        v75 := GenS7{F0: 4294967295, F1: 42}
                                                                        _ = v75
                                                                    } else {
                                                                        v76 := GenS3{F0: (v29 - v3), F1: (^(v45))}
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            if (!false) {
                                                                                var v77 float32 = 0.0
                                                                                fmt.Println(v2, v45, v77)
                                                                                fmt.Println(v69, v70)
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = 9223372036854775807
                                                                                }
                                                                                _ = v77
                                                                            } else {
                                                                                if ("" == "") {
                                                                                    _ = len(v1)
                                                                                } else {
                                                                                    v78 := GenS2{F0: v29, F1: "\n\t"}
                                                                                    v69 = 272
                                                                                    var v79 uint64 = 855
                                                                                    v45 = v45
                                                                                    fmt.Println(fi68, fi36)
                                                                                    switch "hello" {
                                                                                    case "\n\t":
                                                                                        v80 := GenS3{F0: 1, F1: (^(v45))}
                                                                                        _ = v80
                                                                                    default:
                                                                                        v81 := GenS8{F0: 18446744073709551615, F1: 255, F2: v45}
                                                                                        _ = v81
                                                                                    }
                                                                                    v3 = 547
                                                                                    fmt.Println(v3, v79)
                                                                                    _ = v78
                                                                                    _ = v79
                                                                                }
                                                                            }
                                                                        case "\n\t":
                                                                            switch ("foo bar" + "foo bar") {
                                                                            case "":
                                                                                v69 += v72
                                                                                var v82 uint = 18446744073709551615
                                                                                if false {
                                                                                    v83 := GenS9{F0: v70, F1: ""}
                                                                                    _ = v83
                                                                                }
                                                                                _ = v82
                                                                            case "hello":
                                                                                var v84 float32 = 1.5
                                                                                _ = v84
                                                                            case "foo bar":
                                                                                v72 = 0
                                                                            }
                                                                        case "a":
                                                                            if (("foo bar" + "") != "") {
                                                                                var v85 int32 = 42
                                                                                if (("" + "a") != "hello") {
                                                                                    var v86 int = fi36
                                                                                    if len(v4) >= 2 {
                                                                                        v4 = v4[1:len(v4)]
                                                                                    }
                                                                                    _ = v86
                                                                                } else {
                                                                                    v72 += 32768
                                                                                }
                                                                                _ = v85
                                                                            }
                                                                        }
                                                                        var v87 int16 = v71
                                                                        _ = v76
                                                                        _ = v87
                                                                    }
                                                                    v71 = v71
                                                                    v69 = v69
                                                                } else {
                                                                    if false {
                                                                        v72 = 767
                                                                        v88 := v1[42]
                                                                        switch (fi68 + fi68) {
                                                                        case 381:
                                                                            var v89 uint8 = 54
                                                                            _ = v89
                                                                        case 1:
                                                                            _ = len(v1)
                                                                        case 716:
                                                                            if false {
                                                                                delete(v1, (v88 / (((v88 >> (((v29 - 0)) & 63))) | 1)))
                                                                                if ("a" != "hello") {
                                                                                    _ = len(v1)
                                                                                } else {
                                                                                    var v90 int16 = (v71 ^ (v71 >> ((4294967295) & 63)))
                                                                                    var v91 uint8 = 127
                                                                                    _ = v90
                                                                                    _ = v91
                                                                                }
                                                                                delete(v1, (v88 % ((v88) | 1)))
                                                                            } else {
                                                                                switch ("hello" + "a") {
                                                                                case "":
                                                                                    var v92 uint16 = 0
                                                                                    _ = v92
                                                                                case "\n\t":
                                                                                    if ((v29 + v3) >= v3) {
                                                                                        var v93 string = "a"
                                                                                        if ((!(false || (255 != 255))) && true) {
                                                                                            v72 += (fi36 & v70)
                                                                                        } else {
                                                                                            var v94 bool = (!(!true))
                                                                                            var v95 string = ("a" + v93)
                                                                                            _ = v94
                                                                                            _ = v95
                                                                                        }
                                                                                        var v96 uint16 = 0
                                                                                        _ = v93
                                                                                        _ = v96
                                                                                    }
                                                                                    v29 += v3
                                                                                    v70 += (v72 >> (((v29 - (^(v3)))) & 63))
                                                                                case "hello":
                                                                                    switch (fi68 | int(v2)) {
                                                                                    case 32768:
                                                                                        fmt.Println(v71, v71, fi36)
                                                                                    default:
                                                                                        v29 = 42
                                                                                    }
                                                                                }
                                                                                if true {
                                                                                    var v97 uint8 = 8
                                                                                    _ = v97
                                                                                }
                                                                                fmt.Println(v69)
                                                                            }
                                                                        }
                                                                        _ = v88
                                                                    }
                                                                    if (1.5 > ((float32((1.5 / (3.14159 - 0.0))) * float32(float64((0.0 - 0.0)))) + float32(3.14159))) {
                                                                        v3 = v29
                                                                    } else {
                                                                        switch (((("" + "") + "") + "\n\t") + "\n\t") {
                                                                        case "hello":
                                                                            v1[2147483647] = 65535
                                                                        case "\n\t":
                                                                            if (false || (!(127 != 1))) {
                                                                                var v98 uint64 = 4294967295
                                                                                v3 = v29
                                                                                var v99 uint = v3
                                                                                _ = v98
                                                                                _ = v99
                                                                            } else {
                                                                                var v100 int16 = 42
                                                                                v1[0] = 0
                                                                                var v101 int32 = -2147483648
                                                                                _ = v100
                                                                                _ = v101
                                                                            }
                                                                        case "foo bar":
                                                                            var v102 int = v72
                                                                            _ = v102
                                                                        }
                                                                        v3 = v29
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            var v103 int64 = (v2 ^ v2)
                                                                            _ = v103
                                                                        }
                                                                    }
                                                                    fmt.Println(v72)
                                                                }
                                                                v71 += v71
                                                                if (!true) {
                                                                    if (true && true) {
                                                                        v104 := GenS4{F0: 1, F1: 2}
                                                                        var v105 int = v72
                                                                        v106 := GenS3{F0: v29, F1: v45}
                                                                        _ = v104
                                                                        _ = v105
                                                                        _ = v106
                                                                    }
                                                                    if false {
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                    }
                                                                    var v107 int8 = 16
                                                                    _ = v107
                                                                }
                                                            }
                                                        }
                                                        _ = v72
                                                    } else {
                                                        v71 += v71
                                                        if (0 != 18446744073709551615) {
                                                            var v108 uint8 = 255
                                                            if (27 <= 0) {
                                                                v108 += v108
                                                                if (false && ((0 < 1) && false)) {
                                                                    v109 := GenS2{F0: 255, F1: ("\n\t" + (((("\n\t" + ("foo bar" + (((("a" + "") + "hello") + "") + "hello"))) + "\n\t") + "\n\t") + ("hello" + ("a" + "foo bar"))))}
                                                                    var v110 uint64 = 157
                                                                    v108 = v108
                                                                    _ = v109
                                                                    _ = v110
                                                                }
                                                            } else {
                                                                var v111 int8 = v45
                                                                if (true && (v45 >= (v45 + (v111 & (v45 | (v111 << ((1) & 63))))))) {
                                                                    v29 -= v29
                                                                    var v112 float32 = 1e10
                                                                    _ = v112
                                                                }
                                                                _ = v111
                                                            }
                                                            v69 -= fi36
                                                            v108 = v108
                                                            _ = v108
                                                        }
                                                        switch fi68 {
                                                        case -348:
                                                            v45 = v45
                                                        default:
                                                            var v113 bool = (true || false)
                                                            _ = v113
                                                        }
                                                    }
                                                    _ = v71
                                                } else {
                                                    switch 127 {
                                                    case 32767:
                                                        v70 -= (^(fi36))
                                                    case -678:
                                                        if true {
                                                            v29 = (^(v3))
                                                        } else {
                                                            v114 := GenS6{F0: (v2 * 600)}
                                                            var v115 int = fi68
                                                            v45 += (v45 / ((v45) | 1))
                                                            _ = v114
                                                            _ = v115
                                                        }
                                                    case 7:
                                                        v4 = append(v4, v70)
                                                    default:
                                                        v29 = 65535
                                                    }
                                                }
                                                v116 := GenS6{F0: 724}
                                                _ = v116
                                                _ = v69
                                                _ = v70
                                            case -2147483648:
                                                v117 := GenS10{F0: "foo bar"}
                                                v118 := v1[2147483648]
                                                _ = v117
                                                _ = v118
                                            case -9223372036854775808:
                                                var v119 int = (fi68 >> (((^(v3))) & 63))
                                                _ = v119
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = fi68
                                        }
                                        for fi120 := 0; fi120 < 1; fi120++ {
                                            var v121 uint64 = 884
                                            _ = fi120
                                            _ = v121
                                        }
                                    }
                                } else {
                                    var v122 int8 = (^(v45))
                                    var v123 int = (^(fi36))
                                    _ = v122
                                    _ = v123
                                }
                            }
                            _ = fi36
                            _ = v45
                        }
                        _ = v29
                    }
                case "\n\t":
                    fmt.Println(v3, v3)
                }
            }
        }
        fmt.Println(v3, v3)
    } else {
        var v124 float32 = 3.14159
        if (true || true) {
            if len(v4) > 0 {
                v4[0] = 845
            }
            if false {
                for fi125 := 0; fi125 < 7; fi125++ {
                    var v126 uint32 = 0
                    var v127 bool = false
                    v127 = v127
                    _ = fi125
                    _ = v126
                    _ = v127
                }
                var v128 uint = (v3 >> ((18446744073709551615) & 63))
                if false {
                    switch (("a" + "hello") + ("" + "hello")) {
                    case "\n\t":
                        for fi129 := 0; fi129 < 13; fi129++ {
                            v130 := GenS11{F0: -44, F1: (fi129 * (^(fi129))), F2: (true && (v128 > (v3 * (v128 << (((v3 - (v128 / ((v3) | 1)))) & 63)))))}
                            v124 += v124
                            v124 += v124
                            _ = fi129
                            _ = v130
                        }
                    }
                }
                _ = v128
            } else {
                if true {
                    var v131 bool = (1 > -832)
                    _ = v131
                } else {
                    v124 = v124
                    v132 := GenS12{F0: (-(v2)), F1: "\n\t"}
                    _ = v132
                }
            }
            if ((3.14159 <= -2.25) && (false || false)) {
                var v133 int16 = 32767
                var v134 string = (("a" + ("foo bar" + ("\n\t" + ("\n\t" + "foo bar")))) + "\n\t")
                switch -930 {
                case -2147483648:
                    var v135 float64 = float64(v124)
                    _ = v135
                default:
                    var v136 int32 = 2147483647
                    _ = v136
                }
                _ = v133
                _ = v134
            }
        }
        var v137 uint = v3
        _ = v124
        _ = v137
    }
    v3 = v3
    if true {
        for fi138 := 0; fi138 < 5; fi138++ {
            v2 -= 483
            v2 = (v2 * 9223372036854775807)
            if true {
                if false {
                    _ = len(v4)
                } else {
                    switch (((("a" + "foo bar") + "a") + ("\n\t" + "foo bar")) + ("\n\t" + "\n\t")) {
                    case "":
                        var v139 uint8 = 225
                        _ = v139
                    case "\n\t":
                        fmt.Println(v2)
                    default:
                        for fi140 := 0; fi140 < 10; fi140++ {
                            if (true && ((749 <= 743) || false)) {
                                switch "a" {
                                case "\n\t":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                default:
                                    v141 := v1[2]
                                    _ = v141
                                }
                            }
                            v2 += v2
                            _ = fi140
                        }
                    }
                }
                v142 := GenS13{F0: 1}
                fmt.Println(fi138, fi138)
                _ = v142
            } else {
                var v143 uint64 = 962
                _ = len(v4)
                fmt.Println(v2)
                _ = v143
            }
            _ = fi138
        }
        for fi144 := 0; fi144 < 17; fi144++ {
            if (((((255 != 127) && (v2 < (v2 / ((v2) | 1)))) || ((!(true && false)) || true)) && (1 >= 1)) && false) {
                if (255 != 199) {
                    if (42 >= 1000) {
                        v3 = (^(v3))
                        if len(v4) > 0 {
                            v4[0] = (fi144 | (fi144 + fi144))
                        }
                        var v145 uint16 = 32767
                        _ = v145
                    }
                    var v146 int16 = 1
                    _ = v146
                }
                v2 = v2
            }
            _ = fi144
        }
    } else {
        v147 := GenS14{F0: v2}
        _ = v147
    }
    var v148 int32 = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v148:", v148)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v148
    _ = v2
    _ = v3
    _ = v4
}
