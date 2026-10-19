package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int8 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool; F1 uint64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint32; F1 int; F2 bool }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS12 struct { F0 int16; F1 int64 }
type GenS13 struct { F0 string; F1 bool; F2 bool }
func genFn0() int8 {
    var v5 int16 = 42
    v6 := []int{}
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 uint32 = 2147483648
        fmt.Println(v8, v5)
        var v9 int = (fi7 % ((fi7) | 1))
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint64 = 0
    _ = v10
    _ = v5
    _ = v6
    return 127
}
func genFn1(p0 int16, p1 int16, p2 string) int64 {
    var v11 int32 = -2147483648
    v12 := []int{9223372036854775807, 127, 13, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    var v3 uint16 = 45
    var v4 int8 = -1
    if false {
        if (true || (!(0.1 >= 1.5))) {
            for fi13 := 0; fi13 < 16; fi13++ {
                var v14 string = ""
                if false {
                    if (2147483647 >= 2147483647) {
                        v15 := []int{}
                        if (0 < 523) {
                            var v16 uint32 = 35
                            v17 := GenS1{F0: -1, F1: 1, F2: (-(v4))}
                            _ = v16
                            _ = v17
                        }
                        _ = v15
                    }
                    v3 = v3
                } else {
                    fmt.Println(v4, v4, v4)
                }
                _ = fi13
                _ = v14
            }
        } else {
            v18 := GenS2{F0: 18446744073709551615, F1: 1}
            _ = v18
        }
        switch ("" + "foo bar") {
        case "":
            v19 := GenS3{F0: 42}
            _ = v19
        default:
            if true {
                var v20 int32 = 1
                _ = v20
            } else {
                switch -137 {
                case 32767:
                    switch 128 {
                    case 2147483647:
                        var v21 int64 = 1
                        _ = v21
                    default:
                        if (false || (((3.14159 + -2.25) - float32(3.14159)) < 1e10)) {
                            for fi22 := 0; fi22 < 17; fi22++ {
                                switch "a" {
                                case "\n\t":
                                    v2[-1] = 904
                                }
                                for fi23 := 0; fi23 < 10; fi23++ {
                                    var v24 uint = 1
                                    v4 = v4
                                    v3 = (v3 - 42)
                                    if (((4294967295 == 1) && false) || ("" < "")) {
                                        var v25 float32 = 1.5
                                        _ = v25
                                    } else {
                                        v4 = (v4 - (v4 & -62))
                                    }
                                    _ = fi23
                                    _ = v24
                                }
                                _ = fi22
                            }
                            var v26 float64 = (-2.25 + -2.25)
                            v3 = 255
                            _ = v26
                        } else {
                            var v27 bool = true
                            _ = v27
                        }
                    }
                case -119:
                    switch 127 {
                    case 2147483647:
                        if (v3 <= v3) {
                            v3 = v3
                            var v28 string = "hello"
                            _ = v28
                        } else {
                            v29 := []int{}
                            _ = v29
                        }
                    }
                }
                v4 -= 127
                if true {
                    v4 -= v4
                } else {
                    for fi30 := 0; fi30 < 11; fi30++ {
                        var v31 float64 = 1.5
                        if false {
                            var v32 uint32 = 2
                            v32 += v32
                            _ = v32
                        }
                        _ = fi30
                        _ = v31
                    }
                    if (4294967295 == 72) {
                        v4 = 100
                    } else {
                        if true {
                            var v33 int64 = int64(v4)
                            for fi34 := 0; fi34 < 12; fi34++ {
                                switch "" {
                                case "":
                                    v35 := []int{}
                                    _ = v35
                                default:
                                    if false {
                                        v33 += genFn1(1, 1000, "foo bar")
                                        v4 -= v4
                                        v36 := GenS2{F0: 555, F1: -32768}
                                        _ = v36
                                    }
                                }
                                var v37 int64 = -816
                                fmt.Println(v33)
                                _ = fi34
                                _ = v37
                            }
                            _ = v33
                        }
                    }
                }
            }
        }
    } else {
        var v38 uint8 = 165
        var v39 int32 = 2147483647
        var v40 int8 = v4
        _ = v38
        _ = v39
        _ = v40
    }
    if true {
        var v41 int32 = 0
        var v42 int16 = -111
        v43 := GenS2{F0: 822, F1: (v42 / (((v42 % (((v42 << ((42) & 63))) | 1))) | 1))}
        _ = v41
        _ = v42
        _ = v43
    } else {
        fmt.Println(v3)
        if (true || (0.0 <= 0.0)) {
            v4 = (v4 | genFn0())
        } else {
            if ((42 != -66) || false) {
                for fi44 := 0; fi44 < 5; fi44++ {
                    fmt.Println(v3, fi44, v4)
                    _ = fi44
                }
                if (51 == 255) {
                    switch 128 {
                    case -730:
                        v3 = 42
                    }
                    v45 := GenS4{F0: 881}
                    var v46 int16 = -1
                    _ = v45
                    _ = v46
                }
                var v47 float64 = 0.1
                _ = v47
            } else {
                v3 += v3
                v3 = v3
                v48 := v2[4294967295]
                _ = v48
            }
            var v49 float64 = 0.1
            _ = v49
        }
    }
    v50 := GenS2{F0: 442, F1: 1000}
    if true {
        for fi51 := 0; fi51 < 19; fi51++ {
            var v52 uint16 = v3
            for fi53 := 0; fi53 < 14; fi53++ {
                if (true && true) {
                    switch "a" {
                    case "foo bar":
                        switch ("" + "\n\t") {
                        case "hello":
                            v54 := GenS2{F0: 0, F1: 1}
                            _ = v54
                        case "":
                            v55 := []int{fi53, int(v4)}
                            _ = v55
                        case "\n\t":
                            switch (fi53 ^ (fi51 << ((599) & 63))) {
                            case -177:
                                v52 -= v52
                            case -21:
                                var v56 int8 = genFn0()
                                _ = v56
                            }
                        }
                    case "hello":
                        switch fi51 {
                        case 1:
                            v57 := GenS5{F0: 4294967295}
                            _ = v57
                        default:
                            var v58 float64 = -2.25
                            _ = v58
                        }
                    case "\n\t":
                        var v59 int32 = 1
                        _ = v59
                    }
                    var v60 bool = false
                    if (3.14159 < 0.0) {
                        if true {
                            if v60 {
                                if v60 {
                                    var v61 uint32 = 0
                                    if v60 {
                                        var v62 int32 = -650
                                        fmt.Println(fi51, v60, v62)
                                        _ = v62
                                    }
                                    var v63 int64 = 1
                                    _ = v61
                                    _ = v63
                                } else {
                                    var v64 string = "hello"
                                    v52 = (v3 - 480)
                                    _ = v64
                                }
                                if false {
                                    var v65 uint32 = 669
                                    v66 := GenS6{F0: true, F1: 2}
                                    var v67 uint32 = 2147483648
                                    var v68 int64 = 0
                                    _ = v65
                                    _ = v66
                                    _ = v67
                                    _ = v68
                                }
                                var v69 uint8 = 0
                                _ = v69
                            } else {
                                if (true && (1 <= 888)) {
                                    _ = len(v1)
                                    v70 := GenS7{F0: 2}
                                    _ = v70
                                }
                            }
                        } else {
                            var v71 int64 = genFn1(1, 245, ("\n\t" + (("a" + "a") + "")))
                            _ = v71
                        }
                        v4 -= -1
                    }
                    switch (-(fi51)) {
                    case -147:
                        var v72 int32 = int32(v4)
                        _ = v72
                    case 0:
                        v73 := GenS2{F0: 223, F1: 0}
                        _ = v73
                    }
                    v74 := GenS4{F0: -1}
                    fmt.Println(fi51, v4, fi51)
                    _ = v60
                    _ = v74
                } else {
                    var v75 uint = 317
                    _ = v75
                }
                if (!(!true)) {
                    v3 -= v52
                }
                var v76 int = (^(fi51))
                switch "" {
                case "a":
                    switch "hello" {
                    case "\n\t":
                        var v77 string = ("\n\t" + "a")
                        _ = v77
                    default:
                        v3 = 65535
                    }
                }
                var v78 string = "foo bar"
                var v79 bool = true
                v3 -= (v52 << ((65535) & 63))
                _ = fi53
                _ = v76
                _ = v78
                _ = v79
            }
            _ = fi51
            _ = v52
        }
        for fi80 := 0; fi80 < 3; fi80++ {
            var v81 string = "foo bar"
            for fi82 := 0; fi82 < 4; fi82++ {
                if false {
                    v3 += v3
                    switch v81 {
                    case "foo bar":
                        var v83 float64 = (0.0 + (((((((float64(1e10) + (-2.25 / ((1e10 / 0.1) - 1.5))) * (float64(float32(float64(0.1))) * 0.0)) * (1e10 - float64(-2.25))) - ((float64(0.1) + 3.14159) + 1.5)) / ((-2.25 * (1.5 / 0.1)) - float64(3.14159))) - 1.5) + -2.25))
                        _ = v83
                    }
                }
                var v84 string = ("\n\t" + "")
                v3 = 32767
                _ = fi82
                _ = v84
            }
            for fi85 := 0; fi85 < 1; fi85++ {
                var v86 int32 = -2147483648
                var v87 uint64 = 4294967295
                _ = fi85
                _ = v86
                _ = v87
            }
            _ = fi80
            _ = v81
        }
        var v88 string = "hello"
        _ = v88
    } else {
        switch "hello" {
        case "a":
            v3 += 473
        default:
            var v89 float64 = float64((((-2.25 + 0.0) / 1.5) / 1e10))
            _ = v89
        }
    }
    v90 := []int{32768, -2147483648, 626, 32767}
    v91 := GenS8{F0: 904, F1: -720, F2: (-991 != -1)}
    if (18446744073709551615 > 18446744073709551615) {
        delete(v2, -509)
    } else {
        var v92 bool = false
        if (1e10 > (1e10 * (float32(0.1) - (0.1 - 1e10)))) {
            delete(v2, 32767)
            if v92 {
                var v93 float32 = ((1.5 / 3.14159) * 0.0)
                if true {
                    for fi94 := 0; fi94 < 14; fi94++ {
                        delete(v1, 473)
                        v95 := GenS9{F0: 18446744073709551615}
                        _ = fi94
                        _ = v95
                    }
                } else {
                    switch 1 {
                    case 128:
                        if false {
                            v4 += v4
                            var v96 float32 = v93
                            for fi97 := 0; fi97 < 8; fi97++ {
                                for fi98 := 0; fi98 < 14; fi98++ {
                                    var v99 uint64 = uint64(v3)
                                    v100 := GenS4{F0: -1}
                                    fmt.Println(v96, fi97, v96)
                                    v99 -= 42
                                    if len(v90) >= 2 {
                                        v90 = v90[1:len(v90)]
                                    }
                                    _ = fi98
                                    _ = v100
                                    _ = v99
                                }
                                v2[127] = 42
                                _ = fi97
                            }
                            _ = v96
                        } else {
                            for fi101 := 0; fi101 < 18; fi101++ {
                                var v102 uint32 = 65535
                                var v103 string = ("a" + "")
                                switch (fi101 >> ((0) & 63)) {
                                case -884:
                                    v93 = (v93 * ((v93 + v93) - (0.0 / 3.14159)))
                                case 4294967295:
                                    for fi104 := 0; fi104 < 20; fi104++ {
                                        if v92 {
                                            v105 := GenS6{F0: v92, F1: uint64(v102)}
                                            if v92 {
                                                if ((v4 * v4) <= (v4 / (((v4 >> ((65535) & 63))) | 1))) {
                                                    v106 := v1[(v102 / (((v102 ^ v102)) | 1))]
                                                    var v107 int8 = genFn0()
                                                    if false {
                                                        v106 = (false || false)
                                                    } else {
                                                        var v108 int32 = 2147483647
                                                        _ = v108
                                                    }
                                                    switch (fi101 >> ((2) & 63)) {
                                                    case 360:
                                                        v109 := v2[(^(fi101))]
                                                        _ = v109
                                                    case 128:
                                                        v110 := GenS10{F0: v106}
                                                        _ = v110
                                                    case -3:
                                                        var v111 uint64 = 0
                                                        _ = v111
                                                    default:
                                                        if len(v90) > 0 {
                                                            v90[0] = fi104
                                                        }
                                                    }
                                                    v112 := GenS6{F0: false, F1: 737}
                                                    fmt.Println(v93, v4, v4)
                                                    _ = v106
                                                    _ = v107
                                                    _ = v112
                                                }
                                                if (!v92) {
                                                    var v113 uint32 = uint32(v3)
                                                    _ = v113
                                                } else {
                                                    v93 = (v93 / float32(1.5))
                                                    fmt.Println(v3, v93)
                                                    v114 := GenS1{F0: -1, F1: -273, F2: v4}
                                                    _ = v114
                                                }
                                            } else {
                                                var v115 uint8 = 255
                                                _ = len(v2)
                                                v3 = (v3 / (((^(v3))) | 1))
                                                _ = v115
                                            }
                                            v3 = (^(v3))
                                            var v116 uint8 = 1
                                            _ = v105
                                            _ = v116
                                        } else {
                                            v117 := GenS11{F0: (v102 + v102), F1: 42, F2: 127}
                                            _ = v117
                                        }
                                        _ = fi104
                                    }
                                case 32768:
                                    v118 := v2[(fi101 / (((fi101 / (((fi101 & 127)) | 1))) | 1))]
                                    _ = v118
                                default:
                                    if true {
                                        var v119 uint64 = 42
                                        if ((("a" + v103) > ((v103 + "hello") + v103)) || (int16(v4) < -532)) {
                                            v120 := GenS5{F0: v119}
                                            if len(v90) > 0 {
                                                v90[0] = -372
                                            }
                                            _ = v120
                                        } else {
                                            v121 := v1[(v102 | v102)]
                                            v102 -= (v102 << ((255) & 63))
                                            var v122 int32 = 1
                                            _ = v121
                                            _ = v122
                                        }
                                        v102 -= v102
                                        _ = v119
                                    } else {
                                        var v123 uint = 129
                                        _ = len(v2)
                                        _ = v123
                                    }
                                }
                                _ = fi101
                                _ = v102
                                _ = v103
                            }
                        }
                    case 2147483647:
                        for fi124 := 0; fi124 < 3; fi124++ {
                            v3 = 32767
                            _ = fi124
                        }
                    }
                    fmt.Println(v92, v92, v92)
                    v4 = (-(v4))
                }
                for fi125 := 0; fi125 < 12; fi125++ {
                    fmt.Println(v93, v3)
                    _ = fi125
                }
                _ = v93
            } else {
                v4 += (v4 & (v4 - genFn0()))
                switch 4294967295 {
                case 128:
                    switch 0 {
                    case 1:
                        var v126 uint8 = 253
                        _ = v126
                    case 589:
                        for fi127 := 0; fi127 < 8; fi127++ {
                            v92 = false
                            for fi128 := 0; fi128 < 20; fi128++ {
                                v129 := GenS12{F0: 454, F1: 706}
                                v92 = (774 != 105)
                                v130 := GenS11{F0: 470, F1: 0, F2: 14}
                                _ = fi128
                                _ = v129
                                _ = v130
                            }
                            _ = fi127
                        }
                    }
                case 1:
                    switch 0 {
                    case 4294967295:
                        fmt.Println(v4, v4, v92)
                    case 2147483647:
                        if len(v90) >= 2 {
                            v90 = v90[1:len(v90)]
                        }
                    default:
                        v131 := GenS8{F0: 4294967295, F1: 0, F2: v92}
                        _ = v131
                    }
                }
            }
        }
        _ = v92
    }
    for fi132 := 0; fi132 < 8; fi132++ {
        var v133 int8 = v4
        v4 -= v133
        _ = fi132
        _ = v133
    }
    if len(v90) > 0 {
        v90[0] = 682
    }
    fmt.Println(v4, v4)
    if len(v90) >= 2 {
        v90 = v90[1:len(v90)]
    }
    if false {
        var v134 int64 = 9223372036854775807
        if false {
            v134 = v134
            fmt.Println(v4)
            if ((v134 >> ((4294967295) & 63)) < v134) {
                if (!(0 < -2147483648)) {
                    for fi135 := 0; fi135 < 15; fi135++ {
                        for fi136 := 0; fi136 < 9; fi136++ {
                            if false {
                                v134 = (-(v134))
                                var v137 string = "hello"
                                if (v134 >= (v134 % ((genFn1(897, -166, "\n\t")) | 1))) {
                                    fmt.Println(v137, fi136, fi135)
                                }
                                var v138 int = (fi135 ^ (fi136 - 127))
                                _ = v137
                                _ = v138
                            }
                            fmt.Println(fi136, v134)
                            var v139 bool = false
                            _ = fi136
                            _ = v139
                        }
                        _ = fi135
                    }
                } else {
                    v90 = append(v90, -1)
                    v4 += (-(v4))
                }
                for fi140 := 0; fi140 < 14; fi140++ {
                    v134 = -9223372036854775808
                    v4 += 0
                    var v141 int16 = 1000
                    _ = fi140
                    _ = v141
                }
                if len(v90) >= 2 {
                    v90 = v90[1:len(v90)]
                }
            }
        } else {
            if false {
                if true {
                    v134 -= (v134 >> ((18446744073709551615) & 63))
                    v134 -= v134
                    v142 := GenS1{F0: (v134 / ((genFn1(int16(v3), -32768, "foo bar")) | 1)), F1: -983, F2: v4}
                    _ = v142
                }
                if true {
                    var v143 int32 = 0
                    _ = v143
                } else {
                    fmt.Println(v4, v3, v3)
                    var v144 int = 0
                    v134 = (v134 / ((9223372036854775807) | 1))
                    _ = v144
                }
                v145 := GenS13{F0: "\n\t", F1: (v4 >= 42), F2: (0 > 1)}
                _ = v145
            } else {
                for fi146 := 0; fi146 < 11; fi146++ {
                    var v147 int64 = v134
                    v3 += v3
                    for fi148 := 0; fi148 < 9; fi148++ {
                        var v149 int16 = 425
                        _ = fi148
                        _ = v149
                    }
                    _ = fi146
                    _ = v147
                }
                _ = len(v1)
                v134 += v134
            }
            v3 -= (^(v3))
        }
        _ = v134
    }
    fmt.Println(v4, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v90:", len(v90), v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v90
    _ = v91
}
