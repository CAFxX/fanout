package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 uint8; F1 int32; F2 int32 }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint16; F1 uint64; F2 uint }
type GenS9 struct { F0 int32; F1 int; F2 uint8 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 bool; F1 uint64 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 int; F1 uint }
type GenS16 struct { F0 uint8 }
func genFn0(p0 bool, p1 int16) uint32 {
    for fi5 := 0; fi5 < 16; fi5++ {
        p1 = p1
        _ = fi5
    }
    p0 = (-2.25 >= (3.14159 * -2.25))
    v6 := GenS1{F0: 32767}
    v7 := []int{289, 0, 128, 0}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 42
}
func genFn1() uint {
    switch ((("foo bar" + ("foo bar" + ("foo bar" + ("a" + "\n\t")))) + ("\n\t" + "")) + "") {
    case "":
        switch (("foo bar" + "hello") + "") {
        case "a":
            var v8 uint32 = genFn0((591 != genFn0(true, 0)), 1)
            _ = v8
        case "\n\t":
            for fi9 := 0; fi9 < 11; fi9++ {
                for fi10 := 0; fi10 < 3; fi10++ {
                    if true {
                        if (1 <= 1) {
                            switch (fi9 + 0) {
                            case 540:
                                v11 := GenS1{F0: 0}
                                _ = v11
                            default:
                                switch "hello" {
                                case "a":
                                    if false {
                                        var v12 uint32 = 1
                                        var v13 float64 = 1e10
                                        var v14 bool = (!true)
                                        _ = v12
                                        _ = v13
                                        _ = v14
                                    }
                                case "foo bar":
                                    var v15 int64 = -128
                                    _ = v15
                                case "":
                                    fmt.Println(fi9, fi9)
                                default:
                                    if false {
                                        var v16 int32 = 823
                                        _ = v16
                                    } else {
                                        v17 := GenS1{F0: 943}
                                        var v18 uint8 = 42
                                        if ((fi10 | (-(fi9))) <= fi10) {
                                            v19 := GenS2{F0: v18}
                                            _ = v19
                                        } else {
                                            switch 127 {
                                            case 445:
                                                v18 -= (v18 + v18)
                                                v20 := GenS2{F0: v18}
                                                v18 = v18
                                                v21 := []int{fi9, (fi10 - (-(fi9))), (fi9 % (((fi9 - (fi10 | 1))) | 1)), fi9}
                                                var v22 float32 = 1.5
                                                _ = v20
                                                _ = v21
                                                _ = v22
                                            case 127:
                                                v23 := GenS2{F0: (v18 / ((128) | 1))}
                                                v24 := GenS3{F0: (65535 != 612), F1: 406}
                                                v25 := GenS4{F0: 2, F1: 533, F2: 0}
                                                _ = v23
                                                _ = v24
                                                _ = v25
                                            default:
                                                v26 := GenS4{F0: 255, F1: -704, F2: 2147483647}
                                                _ = v26
                                            }
                                            v18 = (^(v18))
                                            v18 = v18
                                        }
                                        _ = v17
                                        _ = v18
                                    }
                                }
                            }
                            var v27 int8 = 1
                            _ = v27
                        }
                        var v28 int8 = -128
                        _ = v28
                    } else {
                        if (!true) {
                            var v29 uint32 = 42
                            v29 += v29
                            _ = v29
                        } else {
                            v30 := GenS5{F0: 0, F1: 4294967295}
                            _ = v30
                        }
                    }
                    v31 := []int{fi9, (fi9 % ((fi10) | 1))}
                    var v32 string = "\n\t"
                    _ = fi10
                    _ = v31
                    _ = v32
                }
                for fi33 := 0; fi33 < 13; fi33++ {
                    fmt.Println(fi33, fi9)
                    _ = fi33
                }
                var v34 int = fi9
                _ = fi9
                _ = v34
            }
        case "":
            for fi35 := 0; fi35 < 3; fi35++ {
                v36 := []int{}
                _ = fi35
                _ = v36
            }
        }
    case "foo bar":
        if true {
            for fi37 := 0; fi37 < 11; fi37++ {
                var v38 int8 = 107
                switch (fi37 & (fi37 % ((-9223372036854775808) | 1))) {
                case 2147483647:
                    v38 -= -127
                case 4294967295:
                    var v39 int = (fi37 << ((18446744073709551615) & 63))
                    _ = v39
                }
                v40 := GenS6{F0: 127}
                _ = fi37
                _ = v38
                _ = v40
            }
            for fi41 := 0; fi41 < 20; fi41++ {
                var v42 int = (^(fi41))
                _ = fi41
                _ = v42
            }
            var v43 float64 = float64(1e10)
            _ = v43
        }
    default:
        switch "foo bar" {
        case "":
            var v44 uint16 = 0
            _ = v44
        case "hello":
            for fi45 := 0; fi45 < 16; fi45++ {
                var v46 int32 = 42
                v46 = v46
                v46 -= v46
                _ = fi45
                _ = v46
            }
        case "\n\t":
            v47 := GenS1{F0: 65535}
            _ = v47
        default:
            var v48 uint = 42
            _ = v48
        }
    }
    var v49 int64 = 42
    fmt.Println(v49)
    if true {
        for fi50 := 0; fi50 < 12; fi50++ {
            var v51 int = (-(fi50))
            _ = fi50
            _ = v51
        }
        if true {
            switch "foo bar" {
            case "\n\t":
                if (("foo bar" >= ("" + ("\n\t" + "hello"))) || false) {
                    for fi52 := 0; fi52 < 7; fi52++ {
                        var v53 string = "\n\t"
                        switch v53 {
                        case "foo bar":
                            v49 -= v49
                        }
                        _ = fi52
                        _ = v53
                    }
                } else {
                    if false {
                        if false {
                            var v54 float32 = 1.5
                            var v55 bool = true
                            v55 = v55
                            _ = v54
                            _ = v55
                        }
                        var v56 uint16 = 42
                        switch "a" {
                        case "":
                            if (683 > 455) {
                                v57 := GenS7{F0: 139}
                                var v58 uint32 = 0
                                _ = v57
                                _ = v58
                            } else {
                                v49 = v49
                                v59 := GenS8{F0: 32767, F1: 42, F2: 469}
                                _ = v59
                            }
                        case "hello":
                            for fi60 := 0; fi60 < 18; fi60++ {
                                if (!(!true)) {
                                    v61 := GenS6{F0: 167}
                                    _ = v61
                                }
                                for fi62 := 0; fi62 < 4; fi62++ {
                                    var v63 float64 = -2.25
                                    v64 := []int{128, (-(fi62))}
                                    switch fi62 {
                                    case -1:
                                        fmt.Println(v56)
                                    case 127:
                                        var v65 int8 = -128
                                        _ = v65
                                    case -727:
                                        fmt.Println(v63, v63)
                                    }
                                    _ = fi62
                                    _ = v63
                                    _ = v64
                                }
                                switch fi60 {
                                case -2147483648:
                                    if (1 >= 65535) {
                                        if (4294967295 != 0) {
                                            v66 := []int{fi60, (fi60 / ((fi60) | 1)), -428}
                                            v56 += 2
                                            _ = v66
                                        } else {
                                            if (0.0 != 1.5) {
                                                for fi67 := 0; fi67 < 18; fi67++ {
                                                    v68 := []int{(^(fi67)), fi60}
                                                    _ = fi67
                                                    _ = v68
                                                }
                                            } else {
                                                for fi69 := 0; fi69 < 1; fi69++ {
                                                    v49 = 904
                                                    var v70 uint32 = genFn0(false, 42)
                                                    _ = fi69
                                                    _ = v70
                                                }
                                                for fi71 := 0; fi71 < 1; fi71++ {
                                                    v72 := GenS4{F0: 127, F1: 1, F2: 42}
                                                    v73 := []int{(fi60 + -1), -9223372036854775808, fi60, fi60}
                                                    _ = fi71
                                                    _ = v72
                                                    _ = v73
                                                }
                                            }
                                            if false {
                                                var v74 uint16 = (v56 ^ v56)
                                                var v75 float64 = (((1e10 / 1.0) - 1e10) + 1.5)
                                                v76 := GenS7{F0: 42}
                                                _ = v74
                                                _ = v75
                                                _ = v76
                                            }
                                            var v77 int = -1
                                            _ = v77
                                        }
                                        var v78 int8 = -1
                                        fmt.Println(fi60, v49, fi60)
                                        _ = v78
                                    }
                                default:
                                    switch (fi60 % (((-(fi60))) | 1)) {
                                    case -406:
                                        if false {
                                            var v79 bool = false
                                            v80 := GenS9{F0: 1, F1: fi60, F2: 42}
                                            _ = v79
                                            _ = v80
                                        }
                                    }
                                }
                                _ = fi60
                            }
                        default:
                            var v81 int64 = v49
                            _ = v81
                        }
                        _ = v56
                    } else {
                        v49 -= v49
                        fmt.Println(v49)
                        v49 -= (v49 & -9223372036854775808)
                    }
                    v82 := GenS6{F0: 255}
                    _ = v82
                }
            case "":
                if true {
                    for fi83 := 0; fi83 < 20; fi83++ {
                        for fi84 := 0; fi84 < 11; fi84++ {
                            var v85 uint32 = 156
                            _ = fi84
                            _ = v85
                        }
                        fmt.Println(v49, v49)
                        for fi86 := 0; fi86 < 8; fi86++ {
                            v49 += v49
                            var v87 uint64 = 1
                            _ = fi86
                            _ = v87
                        }
                        _ = fi83
                    }
                }
            }
            switch 4294967295 {
            case -127:
                fmt.Println(v49)
            case -1:
                if true {
                    v88 := GenS10{F0: 739}
                    var v89 float64 = (-2.25 + -2.25)
                    v49 += v49
                    _ = v88
                    _ = v89
                } else {
                    var v90 uint64 = 42
                    if (-299 > 176) {
                        var v91 float64 = -2.25
                        _ = v91
                    } else {
                        v92 := GenS11{F0: false, F1: 1}
                        _ = v92
                    }
                    for fi93 := 0; fi93 < 13; fi93++ {
                        v94 := GenS2{F0: 248}
                        _ = fi93
                        _ = v94
                    }
                    _ = v90
                }
            case 2147483647:
                var v95 uint16 = 65535
                _ = v95
            default:
                fmt.Println(v49)
            }
            v49 += (v49 % (((v49 - (-(v49)))) | 1))
        }
        v96 := GenS12{F0: (true || (false || true))}
        _ = v96
    }
    for fi97 := 0; fi97 < 6; fi97++ {
        var v98 string = ("hello" + "foo bar")
        _ = fi97
        _ = v98
    }
    _ = v49
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 float32 = (((((-2.25 - 0.1) + 3.14159) * 1e10) + 1.5) + (((float32(0.1) * float32(0.0)) * 3.14159) + 3.14159))
    var v3 uint = 1
    v4 := []int{2147483647, -847}
    switch "\n\t" {
    case "a":
        v99 := GenS8{F0: 32767, F1: 18446744073709551615, F2: 1}
        _ = v99
    case "":
        if true {
            v3 = v3
        } else {
            var v100 string = "foo bar"
            var v101 int8 = 100
            _ = len(v4)
            _ = v100
            _ = v101
        }
    case "hello":
        if true {
            v2 -= 0.0
        } else {
            var v102 uint8 = 44
            switch ("\n\t" + ("foo bar" + "")) {
            case "":
                for fi103 := 0; fi103 < 14; fi103++ {
                    for fi104 := 0; fi104 < 20; fi104++ {
                        var v105 int8 = -122
                        _ = fi104
                        _ = v105
                    }
                    switch (fi103 ^ (fi103 ^ (-(fi103)))) {
                    case 127:
                        var v106 int8 = 113
                        _ = v106
                    default:
                        v107 := v1[fi103]
                        _ = v107
                    }
                    _ = fi103
                }
            }
            _ = v102
        }
    default:
        var v108 uint64 = 367
        _ = v108
    }
    if (((42 >= 1) && (2147483647 >= 128)) && (((3.14159 - (float64(-2.25) - 1e10)) - 0.1) <= float64(1e10))) {
        var v109 int16 = 657
        v2 = float32(-2.25)
        _ = v109
    }
    for fi110 := 0; fi110 < 2; fi110++ {
        v2 += (v2 - (v2 + (((((v2 - v2) / 1.0) + 0.0) * ((0.0 + v2) - ((((3.14159 + (v2 * 1.5)) + 3.14159) * 1e10) + ((((v2 * v2) + -2.25) + (v2 / float32(1e10))) + v2)))) / 1.0)))
        for fi111 := 0; fi111 < 12; fi111++ {
            if (227 <= 255) {
                if (0.0 >= ((1.5 + 3.14159) * 3.14159)) {
                    if false {
                        if true {
                            var v112 uint32 = 42
                            _ = v112
                        }
                        fmt.Println(v2, v3, fi111)
                    } else {
                        v113 := GenS8{F0: 0, F1: 0, F2: v3}
                        _ = v113
                    }
                    if (322 < 42) {
                        if true {
                            var v114 uint16 = 516
                            v115 := GenS8{F0: 82, F1: 141, F2: v3}
                            _ = v114
                            _ = v115
                        } else {
                            var v116 int8 = -70
                            v3 += (v3 + v3)
                            var v117 int = fi111
                            _ = v116
                            _ = v117
                        }
                        if (false && false) {
                            v2 = v2
                        } else {
                            switch "" {
                            case "foo bar":
                                var v118 uint64 = 409
                                _ = v118
                            case "":
                                _ = len(v1)
                            default:
                                v119 := GenS13{F0: 4294967295}
                                _ = v119
                            }
                        }
                        v120 := GenS10{F0: 42}
                        _ = v120
                    } else {
                        v3 -= (^(v3))
                    }
                    var v121 int = fi110
                    _ = v121
                } else {
                    v2 -= v2
                    fmt.Println(fi110)
                    v3 = v3
                }
                v122 := GenS5{F0: (v3 * (v3 >> ((226) & 63))), F1: 1}
                v1[fi111] = fi111
                _ = v122
            }
            var v123 string = ("\n\t" + "foo bar")
            fmt.Println(fi111, fi111, fi111)
            _ = fi111
            _ = v123
        }
        _ = fi110
    }
    if (552 < genFn0((-602 < 2147483647), 32767)) {
        v124 := GenS6{F0: 51}
        _ = v124
    } else {
        for fi125 := 0; fi125 < 20; fi125++ {
            v2 = 1e10
            switch "foo bar" {
            case "hello":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case "":
                var v126 uint16 = 255
                _ = v126
            default:
                fmt.Println(fi125, v2)
            }
            _ = fi125
        }
        if (((v2 / 0.1) + float32(-2.25)) < v2) {
            v3 -= 65535
        }
        v127 := GenS14{F0: 42}
        _ = v127
    }
    if true {
        for fi128 := 0; fi128 < 11; fi128++ {
            switch "" {
            case "":
                for fi129 := 0; fi129 < 1; fi129++ {
                    switch fi129 {
                    case -663:
                        v130 := v1[(fi129 / (((fi129 << ((v3) & 63))) | 1))]
                        _ = v130
                    case 42:
                        fmt.Println(fi128, v2)
                    case -707:
                        var v131 int8 = 36
                        _ = v131
                    }
                    var v132 uint64 = 42
                    fmt.Println(fi129)
                    _ = fi129
                    _ = v132
                }
            }
            v133 := GenS6{F0: 127}
            _ = fi128
            _ = v133
        }
    } else {
        for fi134 := 0; fi134 < 1; fi134++ {
            var v135 int = fi134
            v136 := GenS3{F0: true, F1: (v3 % ((v3) | 1))}
            _ = fi134
            _ = v135
            _ = v136
        }
        if (!true) {
            if (0.1 <= 0.1) {
                fmt.Println(v2)
                v137 := GenS4{F0: 35, F1: -17, F2: 2147483647}
                v138 := GenS15{F0: 4294967295, F1: genFn1()}
                _ = v137
                _ = v138
            } else {
                switch 127 {
                case 32767:
                    switch "" {
                    case "hello":
                        var v139 int32 = -68
                        _ = v139
                    case "a":
                        for fi140 := 0; fi140 < 8; fi140++ {
                            switch "a" {
                            case "":
                                for fi141 := 0; fi141 < 5; fi141++ {
                                    var v142 float32 = (v2 + v2)
                                    var v143 uint8 = 128
                                    _ = fi141
                                    _ = v142
                                    _ = v143
                                }
                            }
                            var v144 int32 = -2147483648
                            _ = fi140
                            _ = v144
                        }
                    }
                }
                for fi145 := 0; fi145 < 1; fi145++ {
                    if (-9223372036854775808 < -205) {
                        var v146 int8 = 127
                        fmt.Println(v2)
                        _ = v146
                    }
                    v2 -= (v2 * (v2 * float32(((3.14159 + 0.0) / 1.5))))
                    if ((1e10 + 0.1) == 1e10) {
                        var v147 int8 = 0
                        var v148 uint = (^(v3))
                        var v149 uint = (v148 << ((v3) & 63))
                        _ = v147
                        _ = v148
                        _ = v149
                    }
                    _ = fi145
                }
            }
            v150 := GenS7{F0: 198}
            _ = v150
        } else {
            v3 = v3
            fmt.Println(v3)
            v2 -= (1e10 + float32((1.5 * (0.0 - (3.14159 - (3.14159 - 1e10))))))
        }
    }
    var v151 int32 = 42
    v152 := v1[-9223372036854775808]
    delete(v1, (v152 / ((v152) | 1)))
    var v153 uint32 = 2
    if (!false) {
        for fi154 := 0; fi154 < 6; fi154++ {
            fmt.Println(v151, fi154)
            for fi155 := 0; fi155 < 11; fi155++ {
                v156 := GenS16{F0: 42}
                fmt.Println(v151, v153)
                var v157 uint64 = 1
                v157 += 999
                v152 -= -180
                v151 += v151
                _ = fi155
                _ = v156
                _ = v157
            }
            _ = fi154
        }
    }
    v2 -= float32(((1.5 * float64(((-2.25 / v2) - float32((-2.25 * (1e10 - 0.0)))))) * ((3.14159 * (1e10 + (float64(1e10) * 3.14159))) * (0.1 + 3.14159))))
    v151 += v151
    v152 -= (-(v152))
    v153 -= v153
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v151:", v151)
    fmt.Println("v152:", v152)
    fmt.Println("v153:", v153)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v151
    _ = v152
    _ = v153
    _ = v2
    _ = v3
    _ = v4
}
