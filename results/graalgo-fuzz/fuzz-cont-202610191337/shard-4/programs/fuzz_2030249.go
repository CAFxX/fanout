package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 int32; F1 int8; F2 string }
type GenS6 struct { F0 uint16; F1 uint; F2 uint }
type GenS7 struct { F0 uint32; F1 uint32 }
type GenS8 struct { F0 string; F1 int16 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint64; F1 uint8 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint64; F1 int64 }
type GenS14 struct { F0 int; F1 int64 }
type GenS15 struct { F0 int64 }
func genFn0(p0 uint32, p1 uint64) int16 {
    fmt.Println(p1)
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 uint32 = p0
        var v7 uint64 = 1
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 uint8 = 223
        p1 -= uint64(p0)
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        if (((fi10 >> ((169) & 63)) != -9223372036854775808) || (true && true)) {
            v11 := GenS1{F0: 266, F1: 0}
            v12 := []int{(fi10 >> ((2) & 63)), (fi10 / ((960) | 1)), (fi10 & fi10)}
            p0 += p0
            _ = v11
            _ = v12
        }
        v13 := []int{fi10}
        v14 := GenS1{F0: (p1 & 723), F1: 2147483647}
        _ = fi10
        _ = v13
        _ = v14
    }
    p1 -= 4294967295
    _ = p0
    _ = p1
    return 32767
}

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{}
    v3 := make(map[int]uint32)
    var v4 float64 = -2.25
    if (!true) {
        v4 = 1.5
        for fi15 := 0; fi15 < 20; fi15++ {
            for fi16 := 0; fi16 < 10; fi16++ {
                switch 128 {
                case 9223372036854775807:
                    v4 = float64(((-2.25 + 1e10) * 0.1))
                    var v17 int = (fi16 | (fi16 | (fi16 / ((fi16) | 1))))
                    if len(v2) > 0 {
                        v2[0] = 0
                    }
                    _ = v17
                case 32768:
                    if false {
                        v4 -= (v4 - 0.0)
                    } else {
                        if false {
                            switch fi15 {
                            case 1:
                                var v18 uint = 255
                                _ = v18
                            case 476:
                                fmt.Println(fi15)
                            default:
                                var v19 float32 = 0.1
                                _ = v19
                            }
                            if (42 < (fi15 ^ 715)) {
                                var v20 float32 = 1e10
                                v21 := GenS1{F0: 2, F1: 220}
                                switch ("a" + "foo bar") {
                                case "\n\t":
                                    v22 := GenS1{F0: 651, F1: 42}
                                    _ = v22
                                case "":
                                    var v23 uint8 = 1
                                    _ = v23
                                case "foo bar":
                                    var v24 uint = 2
                                    _ = v24
                                default:
                                    v20 = v20
                                }
                                _ = v20
                                _ = v21
                            } else {
                                var v25 uint16 = 32767
                                if ((108 == 42) || (!((65535 > 255) && (!(0.1 < (1.5 - 0.1)))))) {
                                    v26 := GenS1{F0: 1, F1: -349}
                                    v27 := GenS1{F0: 418, F1: 1}
                                    v28 := GenS2{F0: 870}
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                                if false {
                                    v29 := GenS2{F0: 488}
                                    _ = v29
                                }
                                var v30 int8 = 100
                                fmt.Println(v25)
                                _ = v25
                                _ = v30
                            }
                            switch ("" + "foo bar") {
                            case "\n\t":
                                fmt.Println(v4, v4, fi15)
                            case "hello":
                                fmt.Println(fi16, fi16, fi16)
                            default:
                                if false {
                                    fmt.Println(v4, fi15, fi15)
                                } else {
                                    if false {
                                        v4 += (v4 * v4)
                                        switch fi15 {
                                        case 825:
                                            v4 += (float64(3.14159) * 3.14159)
                                            v1[(fi16 | 847)] = ""
                                        case 1:
                                            v4 = (v4 + (float64((3.14159 - 1.5)) / v4))
                                            v31 := GenS1{F0: 18446744073709551615, F1: -568}
                                            fmt.Println(fi15, fi16)
                                            _ = v31
                                        default:
                                            _ = len(v1)
                                        }
                                        fmt.Println(v4, v4, fi16)
                                    } else {
                                        var v32 float64 = (float64(float32(v4)) / v4)
                                        _ = v32
                                    }
                                    if true {
                                        if false {
                                            var v33 float32 = 3.14159
                                            v34 := v3[9223372036854775807]
                                            _ = v33
                                            _ = v34
                                        } else {
                                            v4 = (v4 * (v4 * float64(-2.25)))
                                            var v35 float32 = (-2.25 + 1e10)
                                            _ = v35
                                        }
                                    } else {
                                        delete(v1, fi15)
                                    }
                                }
                            }
                            v36 := GenS3{F0: 204, F1: 1}
                            v37 := GenS4{F0: true, F1: -1}
                            _ = v36
                            _ = v37
                        }
                        if false {
                            v4 -= ((3.14159 * 1.5) * v4)
                            switch "\n\t" {
                            case "a":
                                switch "foo bar" {
                                case "\n\t":
                                    if false {
                                        _ = len(v2)
                                        v3[fi15] = 2
                                        var v38 uint64 = 42
                                        _ = v38
                                    }
                                case "":
                                    var v39 float32 = 0.1
                                    _ = v39
                                default:
                                    var v40 uint32 = 1
                                    _ = v40
                                }
                            default:
                                v4 += float64((1e10 + 0.0))
                                var v41 float32 = (float32(v4) * 1.5)
                                _ = v41
                            }
                            var v42 uint = 65535
                            _ = v42
                        } else {
                            v4 += ((v4 * (float64(((0.1 + ((1e10 * float32((v4 + 0.0))) * -2.25)) - 1.5)) - 3.14159)) / 1e10)
                            fmt.Println(v4, v4, v4)
                        }
                        fmt.Println(fi15, fi16)
                        if true {
                            delete(v3, 42)
                        } else {
                            if len(v2) > 0 {
                                v2[0] = 2147483648
                            }
                            var v43 float32 = 1e10
                            _ = v43
                        }
                    }
                case -1:
                    fmt.Println(fi15, v4)
                }
                _ = fi16
            }
            v44 := GenS5{F0: 2147483647, F1: 0, F2: "a"}
            _ = fi15
            _ = v44
        }
    }
    v45 := GenS6{F0: 65535, F1: 155, F2: 4294967295}
    v46 := GenS7{F0: 2147483647, F1: 4294967295}
    if true {
        _ = len(v2)
        v4 = ((v4 + v4) / -2.25)
    } else {
        if (32767 < -317) {
            v4 -= (float64((0.0 + (float32((((0.0 * 0.0) - v4) - 3.14159)) - (1.5 + 0.0)))) - float64(-2.25))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v4 = 1.5
        } else {
            for fi47 := 0; fi47 < 2; fi47++ {
                v4 -= v4
                if false {
                    switch "" {
                    case "hello":
                        var v48 int16 = 1
                        _ = v48
                    }
                    v4 = v4
                    v49 := GenS8{F0: "a", F1: -876}
                    _ = v49
                }
                var v50 bool = true
                _ = fi47
                _ = v50
            }
            v51 := v3[128]
            _ = v51
        }
        var v52 uint32 = 744
        _ = v52
    }
    var v53 int64 = 9223372036854775807
    for fi54 := 0; fi54 < 13; fi54++ {
        switch fi54 {
        case -9223372036854775808:
            v4 = ((v4 / v4) / v4)
        case -824:
            if true {
                var v55 int64 = v53
                _ = v55
            } else {
                if (!true) {
                    fmt.Println(v53, v4)
                    for fi56 := 0; fi56 < 7; fi56++ {
                        v2 = append(v2, 161)
                        _ = fi56
                    }
                    var v57 float64 = v4
                    _ = v57
                } else {
                    var v58 bool = (true && true)
                    switch "foo bar" {
                    case "\n\t":
                        var v59 int16 = -974
                        _ = v59
                    case "foo bar":
                        v60 := GenS9{F0: ("a" + ("a" + ("hello" + "")))}
                        _ = v60
                    case "a":
                        var v61 int64 = v53
                        _ = v61
                    }
                    _ = v58
                }
                v62 := GenS10{F0: 2, F1: 149}
                _ = v62
            }
        default:
            for fi63 := 0; fi63 < 17; fi63++ {
                v64 := GenS1{F0: 446, F1: -472}
                v2 = append(v2, 0)
                var v65 uint64 = 42
                _ = fi63
                _ = v64
                _ = v65
            }
        }
        _ = fi54
    }
    if false {
        var v66 uint = 4294967295
        switch "foo bar" {
        case "hello":
            if (true || true) {
                _ = len(v3)
                fmt.Println(v66, v4, v66)
                for fi67 := 0; fi67 < 18; fi67++ {
                    if (!true) {
                        _ = len(v1)
                        if ((255 == 128) && (0 > 1)) {
                            if (((true && ((0.1 + (0.0 - float32((float64((((1.5 / 3.14159) - ((0.1 - float32(float64(1.5))) * (-2.25 * -2.25))) - float32(1e10))) + (v4 / v4))))) >= float32(v4))) || (((127 >= 1) || (true && false)) && false)) && true) {
                                v1[fi67] = ("hello" + ("" + "hello"))
                                if true {
                                    var v68 int64 = (v53 << (((v66 << (((v66 % (((^(v66))) | 1))) & 63))) & 63))
                                    var v69 int = (fi67 ^ (^(fi67)))
                                    var v70 bool = (1000 > genFn0(65535, 0))
                                    _ = v68
                                    _ = v69
                                    _ = v70
                                }
                                switch (fi67 & (fi67 >> ((v66) & 63))) {
                                case 1:
                                    switch -700 {
                                    case 389:
                                        for fi71 := 0; fi71 < 20; fi71++ {
                                            var v72 uint16 = 2
                                            var v73 uint64 = 656
                                            if true {
                                                v74 := GenS5{F0: 2147483647, F1: -122, F2: ("\n\t" + ("foo bar" + "hello"))}
                                                v53 += v53
                                                _ = v74
                                            }
                                            _ = fi71
                                            _ = v72
                                            _ = v73
                                        }
                                    case 475:
                                        for fi75 := 0; fi75 < 5; fi75++ {
                                            switch ((("hello" + "hello") + "\n\t") + ("hello" + "foo bar")) {
                                            case "\n\t":
                                                switch (-(fi67)) {
                                                case 9223372036854775807:
                                                    var v76 string = "hello"
                                                    _ = v76
                                                case -9223372036854775808:
                                                    if (42 <= -47) {
                                                        v77 := GenS4{F0: true, F1: 42}
                                                        var v78 float64 = (3.14159 - (v4 / ((0.0 * v4) * float64(1.5))))
                                                        if ((float32(v78) / 1e10) <= float32((v78 - v78))) {
                                                            _ = len(v2)
                                                            v78 -= v4
                                                        } else {
                                                            switch "hello" {
                                                            case "foo bar":
                                                                v78 -= (v78 / (v4 / ((v4 / 1.5) - (v4 * v78))))
                                                            case "\n\t":
                                                                if (!false) {
                                                                    v79 := v3[fi67]
                                                                    _ = v79
                                                                }
                                                            case "":
                                                                var v80 bool = (-2.25 != 1.5)
                                                                _ = v80
                                                            default:
                                                                v78 = 0.1
                                                            }
                                                            var v81 float64 = 3.14159
                                                            v4 -= v4
                                                            _ = v81
                                                        }
                                                        _ = v77
                                                        _ = v78
                                                    } else {
                                                        switch fi75 {
                                                        case -9223372036854775808:
                                                            v82 := GenS1{F0: 29, F1: -2147483648}
                                                            v4 = v4
                                                            v2 = append(v2, 536)
                                                            _ = v82
                                                        case 32767:
                                                            _ = len(v2)
                                                            v53 = -431
                                                        }
                                                        fmt.Println(v53, v4)
                                                        var v83 int16 = int16(fi75)
                                                        _ = v83
                                                    }
                                                    var v84 int = (fi75 + 1)
                                                    _ = v84
                                                }
                                            case "":
                                                var v85 int16 = 1
                                                _ = v85
                                            default:
                                                if false {
                                                    switch "a" {
                                                    case "foo bar":
                                                        var v86 uint16 = 1
                                                        _ = v86
                                                    case "hello":
                                                        if (!false) {
                                                            if true {
                                                                var v87 uint64 = 18446744073709551615
                                                                switch ("\n\t" + ("foo bar" + "a")) {
                                                                case "":
                                                                    v88 := GenS11{F0: v66}
                                                                    _ = v88
                                                                }
                                                                v66 = 166
                                                                _ = v87
                                                            } else {
                                                                v89 := GenS12{F0: ("a" + ("" + "foo bar"))}
                                                                _ = v89
                                                            }
                                                        }
                                                    case "\n\t":
                                                        v3[32768] = 2147483648
                                                    default:
                                                        var v90 uint8 = 166
                                                        _ = v90
                                                    }
                                                    v91 := GenS12{F0: ("a" + ("" + "a"))}
                                                    v53 = 512
                                                    var v92 uint8 = 144
                                                    fmt.Println(v66, fi75)
                                                    _ = v91
                                                    _ = v92
                                                } else {
                                                    v93 := GenS4{F0: false, F1: genFn0(97, 42)}
                                                    switch -9223372036854775808 {
                                                    case -9223372036854775808:
                                                        var v94 int64 = 913
                                                        _ = v94
                                                    default:
                                                        if ((!false) && false) {
                                                            v95 := GenS11{F0: v66}
                                                            delete(v3, fi67)
                                                            _ = v95
                                                        }
                                                    }
                                                    _ = v93
                                                }
                                                if true {
                                                    v4 = ((v4 * (-2.25 + (float64(3.14159) - v4))) - 3.14159)
                                                    v96 := GenS9{F0: (("" + (("foo bar" + (("\n\t" + "foo bar") + "\n\t")) + "hello")) + ("foo bar" + "hello"))}
                                                    _ = v96
                                                }
                                                var v97 uint16 = 42
                                                _ = v97
                                            }
                                            _ = fi75
                                        }
                                    }
                                case -123:
                                    var v98 int32 = 263
                                    _ = v98
                                }
                            }
                            for fi99 := 0; fi99 < 6; fi99++ {
                                v100 := GenS2{F0: 2}
                                var v101 uint32 = 2
                                var v102 float32 = -2.25
                                _ = len(v1)
                                if (!(!false)) {
                                    v103 := GenS13{F0: 4294967295, F1: v53}
                                    _ = len(v2)
                                    switch (fi67 & int(v53)) {
                                    case -266:
                                        fmt.Println(v102, v4, v53)
                                    default:
                                        switch 42 {
                                        case 1:
                                            fmt.Println(v4)
                                        case 24:
                                            v66 -= v66
                                        default:
                                            v104 := GenS9{F0: ""}
                                            _ = v104
                                        }
                                    }
                                    _ = v103
                                } else {
                                    if (!false) {
                                        v102 += 0.0
                                    }
                                    v102 = ((v102 * float32((0.0 + (1.5 * (v4 * 3.14159))))) / v102)
                                    if true {
                                        v4 += v4
                                        var v105 uint16 = 65535
                                        if ((float64(v102) - 3.14159) >= v4) {
                                            v2 = append(v2, 530)
                                            if true {
                                                switch fi99 {
                                                case 0:
                                                    v66 += (v66 >> ((65535) & 63))
                                                case 1:
                                                    if true {
                                                        switch (fi67 & (-(fi67))) {
                                                        case -1:
                                                            var v106 float64 = v4
                                                            _ = v106
                                                        case 32767:
                                                            switch (("a" + ("hello" + "foo bar")) + "\n\t") {
                                                            case "":
                                                                switch "foo bar" {
                                                                case "a":
                                                                    v107 := GenS11{F0: v66}
                                                                    _ = v107
                                                                }
                                                            case "\n\t":
                                                                switch ("foo bar" + "\n\t") {
                                                                case "":
                                                                    v108 := GenS5{F0: 2147483647, F1: 1, F2: ""}
                                                                    _ = v108
                                                                case "foo bar":
                                                                    fmt.Println(v53)
                                                                case "hello":
                                                                    v101 += v101
                                                                }
                                                            case "foo bar":
                                                                fmt.Println(v4, v101, v66)
                                                            }
                                                        case 9223372036854775807:
                                                            _ = len(v3)
                                                        }
                                                        v109 := GenS14{F0: (^(fi99)), F1: 0}
                                                        if (v4 <= 3.14159) {
                                                            var v110 uint8 = 42
                                                            var v111 int8 = 22
                                                            _ = v110
                                                            _ = v111
                                                        } else {
                                                            var v112 uint8 = 149
                                                            v112 -= uint8(v66)
                                                            v102 += -2.25
                                                            _ = v112
                                                        }
                                                        _ = v109
                                                    } else {
                                                        v113 := v1[(fi99 & (fi99 & 817))]
                                                        v114 := GenS15{F0: -190}
                                                        v113 = "\n\t"
                                                        var v115 uint8 = 42
                                                        v116 := GenS5{F0: 42, F1: 94, F2: v113}
                                                        _ = v113
                                                        _ = v114
                                                        _ = v115
                                                        _ = v116
                                                    }
                                                default:
                                                    if (((127 == -20) || true) && true) {
                                                        v101 -= (v101 | v101)
                                                    } else {
                                                        var v117 uint8 = 111
                                                        _ = v117
                                                    }
                                                }
                                                switch ((("\n\t" + "hello") + "foo bar") + "foo bar") {
                                                case "\n\t":
                                                    v118 := GenS5{F0: -1, F1: 90, F2: "hello"}
                                                    _ = v118
                                                }
                                                v66 -= 0
                                                delete(v3, (-(fi67)))
                                                v119 := GenS4{F0: (true || true), F1: 1000}
                                                _ = v119
                                            } else {
                                                v102 -= -2.25
                                            }
                                        } else {
                                            var v120 float64 = 0.0
                                            _ = v120
                                        }
                                        if true {
                                            v4 -= ((v4 / v4) - 0.1)
                                            v105 -= 682
                                        } else {
                                            v66 = (v66 * 42)
                                            v4 -= v4
                                            switch fi99 {
                                            case 127:
                                                switch 60 {
                                                case 383:
                                                    v4 = (v4 - v4)
                                                case 831:
                                                    var v121 uint64 = 18446744073709551615
                                                    _ = v121
                                                case 757:
                                                    var v122 uint = (v66 ^ (v66 | (v66 * v66)))
                                                    _ = v122
                                                default:
                                                    var v123 int8 = 127
                                                    _ = v123
                                                }
                                            case 382:
                                                fmt.Println(v102, fi67)
                                            case 128:
                                                switch (^(fi99)) {
                                                case 940:
                                                    v124 := GenS2{F0: 1}
                                                    _ = v124
                                                case 127:
                                                    _ = len(v1)
                                                default:
                                                    v105 += v105
                                                }
                                            default:
                                                v4 += 0.0
                                            }
                                        }
                                        v3[32768] = v101
                                        _ = v105
                                    } else {
                                        var v125 float32 = float32(1e10)
                                        _ = v125
                                    }
                                }
                                _ = fi99
                                _ = v100
                                _ = v101
                                _ = v102
                            }
                        }
                        v53 = (-(v53))
                    }
                    for fi126 := 0; fi126 < 19; fi126++ {
                        if len(v2) > 0 {
                            v2[0] = 2147483648
                        }
                        var v127 int64 = v53
                        _ = fi126
                        _ = v127
                    }
                    _ = fi67
                }
            }
        default:
            var v128 int32 = 2147483647
            _ = v128
        }
        _ = v66
    }
    switch 9223372036854775807 {
    case 188:
        var v129 uint = 4294967295
        _ = v129
    }
    v130 := v3[9223372036854775807]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v130:", v130)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v130
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v53
}
