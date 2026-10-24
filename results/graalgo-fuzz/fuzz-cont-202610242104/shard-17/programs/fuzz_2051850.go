package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 int64; F2 uint }
type GenS4 struct { F0 string; F1 int16; F2 int64 }
type GenS5 struct { F0 uint64; F1 uint; F2 uint }
type GenS6 struct { F0 int8; F1 uint16; F2 int8 }
type GenS7 struct { F0 uint8; F1 int8 }
type GenS8 struct { F0 uint32; F1 int8; F2 uint }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int32; F1 int16; F2 string }
type GenS11 struct { F0 uint; F1 uint64 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint32; F1 int; F2 int }
type GenS14 struct { F0 uint16; F1 int8; F2 string }
type GenS15 struct { F0 int; F1 int16 }
type GenS16 struct { F0 int32; F1 bool; F2 uint16 }
func genFn0() bool {
    switch "a" {
    case "foo bar":
        v6 := []int{32768}
        _ = v6
    case "a":
        if (0.1 == 0.0) {
            for fi7 := 0; fi7 < 10; fi7++ {
                switch (^(fi7)) {
                case -133:
                    v8 := GenS1{F0: 766, F1: 224, F2: 0}
                    _ = v8
                case -952:
                    for fi9 := 0; fi9 < 12; fi9++ {
                        switch "a" {
                        case "":
                            var v10 int16 = -124
                            v10 = (^(v10))
                            v11 := GenS2{F0: 0}
                            _ = v10
                            _ = v11
                        case "hello":
                            if false {
                                var v12 float32 = float32(((3.14159 * 0.1) - float64(0.0)))
                                _ = v12
                            }
                        default:
                            v13 := GenS2{F0: 0}
                            _ = v13
                        }
                        _ = fi9
                    }
                case 0:
                    fmt.Println(fi7)
                default:
                    var v14 int32 = 2147483647
                    _ = v14
                }
                v15 := []int{(^(fi7))}
                _ = fi7
                _ = v15
            }
        }
    case "":
        var v16 uint16 = 747
        _ = v16
    }
    v17 := []int{32767, -757, -2147483648, 32768}
    var v18 int32 = 1
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    v19 := GenS2{F0: -128}
    _ = v17
    _ = v18
    _ = v19
    return ((!true) || ((1.5 + 3.14159) < 0.0))
}
func genFn1(p0 int8, p1 int, p2 int32) uint8 {
    p2 = (p2 & p2)
    v20 := []int{(^(p1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = -1
    var v3 uint16 = 2
    var v4 int8 = -1
    var v5 float64 = (float64(1.5) * -2.25)
    if (((1 >= 42) || ((1 != 42) && true)) && true) {
        v21 := []int{v2, v2}
        var v22 uint32 = 2147483648
        for fi23 := 0; fi23 < 5; fi23++ {
            switch (fi23 / ((9223372036854775807) | 1)) {
            case 0:
                var v24 int64 = 0
                _ = v24
            case -177:
                if false {
                    switch ("\n\t" + ("a" + "a")) {
                    case "a":
                        v25 := GenS3{F0: 850, F1: 968, F2: 720}
                        _ = v25
                    case "":
                        var v26 bool = ((-244 <= 42) && false)
                        _ = v26
                    case "foo bar":
                        delete(v1, (^(v22)))
                    default:
                        v27 := GenS2{F0: 78}
                        _ = v27
                    }
                    var v28 uint = 745
                    if false {
                        switch "foo bar" {
                        case "hello":
                            _ = len(v1)
                        case "a":
                            v2 -= v2
                        case "\n\t":
                            if false {
                                v4 = 100
                            }
                        }
                        switch ("foo bar" + "\n\t") {
                        case "":
                            var v29 uint32 = v22
                            _ = v29
                        default:
                            v30 := GenS2{F0: v4}
                            _ = v30
                        }
                    }
                    _ = v28
                } else {
                    v31 := GenS2{F0: 13}
                    if true {
                        for fi32 := 0; fi32 < 1; fi32++ {
                            v5 += 1.5
                            if (!("hello" != ("foo bar" + "hello"))) {
                                v5 -= (v5 / (v5 * float64(0.1)))
                                v33 := GenS4{F0: (("foo bar" + "") + ("foo bar" + "")), F1: 42, F2: 9223372036854775807}
                                var v34 uint8 = 34
                                _ = v33
                                _ = v34
                            }
                            v21 = append(v21, (fi23 % ((fi23) | 1)))
                            _ = fi32
                        }
                    } else {
                        var v35 float64 = v5
                        v36 := GenS5{F0: 42, F1: 65535, F2: 159}
                        if false {
                            var v37 int = fi23
                            v38 := GenS2{F0: (^(v4))}
                            _ = v37
                            _ = v38
                        } else {
                            v4 -= v4
                            v35 = 1.5
                        }
                        _ = v35
                        _ = v36
                    }
                    v3 = 36
                    _ = v31
                }
            case -940:
                v39 := GenS6{F0: (-(v4)), F1: v3, F2: (v4 + (v4 - (v4 - 127)))}
                _ = v39
            }
            if false {
                v5 = ((v5 - ((v5 - float64(3.14159)) + 0.0)) + -2.25)
                for fi40 := 0; fi40 < 17; fi40++ {
                    if len(v21) > 0 {
                        v21[0] = fi40
                    }
                    switch fi40 {
                    case 127:
                        switch "\n\t" {
                        case "\n\t":
                            if true {
                                if false {
                                    var v41 bool = (true || ((float64((3.14159 + 1.5)) < (v5 * 3.14159)) && (((-2.25 - ((float64(0.1) + v5) / 1.5)) + v5) != (v5 / 1.5))))
                                    _ = len(v21)
                                    var v42 uint32 = 263
                                    _ = v41
                                    _ = v42
                                }
                                v43 := GenS6{F0: v4, F1: (v3 ^ v3), F2: (v4 ^ (-(v4)))}
                                fmt.Println(v4)
                                _ = v43
                            } else {
                                v2 = 1
                                if len(v21) >= 2 {
                                    v21 = v21[1:len(v21)]
                                }
                                var v44 int = 0
                                if true {
                                    if true {
                                        if true {
                                            switch ("foo bar" + "") {
                                            case "a":
                                                v45 := GenS5{F0: 811, F1: 1, F2: 0}
                                                _ = v45
                                            case "hello":
                                                v21 = append(v21, 128)
                                            case "\n\t":
                                                if true {
                                                    v46 := GenS7{F0: 33, F1: v4}
                                                    _ = v46
                                                }
                                            }
                                            var v47 uint = 65535
                                            var v48 float32 = -2.25
                                            var v49 uint64 = 155
                                            var v50 string = "foo bar"
                                            if ((^(v44)) >= (fi40 << ((uint(v3)) & 63))) {
                                                if false {
                                                    var v51 float64 = float64((-2.25 + v48))
                                                    var v52 float64 = v5
                                                    v52 = float64((0.1 + v48))
                                                    _ = v51
                                                    _ = v52
                                                }
                                            }
                                            var v53 int = 1
                                            _ = v47
                                            _ = v48
                                            _ = v49
                                            _ = v50
                                            _ = v53
                                        } else {
                                            var v54 bool = (!("" > ("\n\t" + "hello")))
                                            _ = v54
                                        }
                                        if len(v21) > 0 {
                                            v21[0] = int(v3)
                                        }
                                    }
                                } else {
                                    v55 := GenS8{F0: (v22 * 42), F1: v4, F2: uint(v22)}
                                    _ = v55
                                }
                                var v56 int64 = 0
                                _ = v44
                                _ = v56
                            }
                        case "foo bar":
                            var v57 bool = false
                            _ = v57
                        }
                    case 0:
                        if len(v21) > 0 {
                            v21[0] = (fi23 | (fi40 & (v2 / (((v2 | (fi23 & 127))) | 1))))
                        }
                    case 42:
                        if len(v21) > 0 {
                            v21[0] = fi23
                        }
                    }
                    v3 += v3
                    _ = fi40
                }
                var v58 bool = true
                _ = v58
            } else {
                v4 = (v4 * v4)
            }
            if true {
                v59 := GenS1{F0: 65535, F1: 128, F2: 2147483647}
                for fi60 := 0; fi60 < 2; fi60++ {
                    var v61 uint32 = v22
                    if (!(!((^(fi60)) <= (fi60 >> ((uint(v3)) & 63))))) {
                        var v62 uint16 = (v3 << ((18446744073709551615) & 63))
                        fmt.Println(v62, v62)
                        v61 = v61
                        _ = v62
                    } else {
                        v63 := GenS8{F0: (v61 & v22), F1: int8(v3), F2: 255}
                        v5 = v5
                        fmt.Println(v2, v2)
                        _ = v63
                    }
                    fmt.Println(v22, v5, v61)
                    _ = fi60
                    _ = v61
                }
                _ = v59
            } else {
                v64 := v1[549]
                fmt.Println(v2, fi23)
                _ = v64
            }
            _ = fi23
        }
        _ = v21
        _ = v22
    } else {
        for fi65 := 0; fi65 < 17; fi65++ {
            fmt.Println(v4, v4, v4)
            v66 := GenS9{F0: v2}
            _ = fi65
            _ = v66
        }
        for fi67 := 0; fi67 < 2; fi67++ {
            var v68 int = (v2 & 724)
            switch (fi67 / ((fi67) | 1)) {
            case 2147483647:
                if (!true) {
                    for fi69 := 0; fi69 < 12; fi69++ {
                        v70 := GenS2{F0: (v4 << ((1) & 63))}
                        _ = fi69
                        _ = v70
                    }
                    var v71 int = v2
                    v72 := GenS7{F0: 42, F1: v4}
                    _ = v71
                    _ = v72
                } else {
                    var v73 float64 = (v5 * v5)
                    switch ("" + ("hello" + "hello")) {
                    case "\n\t":
                        switch -993 {
                        case -1:
                            var v74 int32 = -442
                            _ = v74
                        case 451:
                            fmt.Println(v68)
                        case 9223372036854775807:
                            v75 := GenS6{F0: v4, F1: v3, F2: v4}
                            _ = v75
                        default:
                            switch 32768 {
                            case 127:
                                for fi76 := 0; fi76 < 20; fi76++ {
                                    switch (fi76 | fi76) {
                                    case 32768:
                                        switch "foo bar" {
                                        case "hello":
                                            if ((v4 % ((v4) | 1)) < -128) {
                                                v77 := v1[707]
                                                _ = v77
                                            }
                                        case "":
                                            if true {
                                                switch ("\n\t" + "a") {
                                                case "hello":
                                                    v4 -= (v4 << ((18446744073709551615) & 63))
                                                case "a":
                                                    v78 := GenS9{F0: fi76}
                                                    _ = v78
                                                case "":
                                                    var v79 int32 = -1
                                                    _ = v79
                                                default:
                                                    fmt.Println(v73, v4, v4)
                                                }
                                                var v80 uint = 18446744073709551615
                                                v81 := GenS10{F0: 0, F1: 42, F2: "\n\t"}
                                                _ = v80
                                                _ = v81
                                            } else {
                                                v5 = 1e10
                                            }
                                            if (!((true && false) && (!true))) {
                                                var v82 int64 = -652
                                                fmt.Println(v2)
                                                switch "\n\t" {
                                                case "a":
                                                    v83 := []int{v68, v2, (v2 - fi76), (v68 ^ 32767)}
                                                    _ = v83
                                                }
                                                _ = v82
                                            } else {
                                                v68 = 1
                                                var v84 int64 = -1
                                                v84 += v84
                                                _ = v84
                                            }
                                        case "a":
                                            var v85 int64 = 762
                                            _ = v85
                                        default:
                                            fmt.Println(fi67, fi76)
                                        }
                                    }
                                    if ((427 >= 4294967295) || ((true && (false || (1 < 859))) && false)) {
                                        var v86 int16 = 339
                                        var v87 int16 = v86
                                        if true {
                                            v88 := GenS5{F0: 2, F1: 4294967295, F2: 0}
                                            v5 = -2.25
                                            v89 := GenS8{F0: 65535, F1: v4, F2: 255}
                                            _ = v88
                                            _ = v89
                                        } else {
                                            if false {
                                                v90 := []int{v68, (fi67 % ((fi67) | 1))}
                                                var v91 uint64 = 42
                                                v92 := GenS6{F0: v4, F1: v3, F2: v4}
                                                var v93 float64 = v73
                                                _ = v90
                                                _ = v91
                                                _ = v92
                                                _ = v93
                                            }
                                        }
                                        _ = v86
                                        _ = v87
                                    }
                                    var v94 int = -2147483648
                                    _ = fi76
                                    _ = v94
                                }
                            case 302:
                                switch "" {
                                case "\n\t":
                                    var v95 int8 = v4
                                    _ = v95
                                case "foo bar":
                                    for fi96 := 0; fi96 < 10; fi96++ {
                                        var v97 uint32 = 2147483648
                                        _ = len(v1)
                                        if true {
                                            fmt.Println(v68, v4, v5)
                                            if ((1 <= -32768) || false) {
                                                v2 -= v2
                                            }
                                            fmt.Println(fi96)
                                        }
                                        switch ("a" + ("hello" + "foo bar")) {
                                        case "foo bar":
                                            v5 = float64((1e10 * -2.25))
                                        case "hello":
                                            v3 -= (v3 * v3)
                                        }
                                        _ = fi96
                                        _ = v97
                                    }
                                case "":
                                    v5 += ((float64(((1e10 - 0.1) + (3.14159 - (((-2.25 / (1.5 - 0.0)) * float32(v73)) / 1.0)))) - 3.14159) * ((v73 - 3.14159) + float64(1e10)))
                                }
                            case 1:
                                v98 := []int{}
                                _ = v98
                            }
                        }
                    case "a":
                        if false {
                            var v99 uint8 = 220
                            v100 := []int{fi67, (fi67 & v68), (-(v68))}
                            v68 = 32767
                            _ = v100
                            _ = v99
                        } else {
                            if true {
                                v1[65535] = 2147483647
                                for fi101 := 0; fi101 < 8; fi101++ {
                                    v102 := GenS6{F0: -75, F1: v3, F2: (v4 | (v4 - v4))}
                                    delete(v1, 586)
                                    var v103 int16 = 42
                                    _ = fi101
                                    _ = v102
                                    _ = v103
                                }
                            }
                        }
                    }
                    _ = v73
                }
            case 474:
                v4 = (v4 >> ((57) & 63))
            default:
                fmt.Println(v4, v2, fi67)
            }
            _ = fi67
            _ = v68
        }
    }
    fmt.Println(v2, v4, v2)
    for fi104 := 0; fi104 < 5; fi104++ {
        v105 := GenS9{F0: -9223372036854775808}
        switch "a" {
        case "foo bar":
            for fi106 := 0; fi106 < 6; fi106++ {
                if true {
                    switch 9223372036854775807 {
                    case 32767:
                        var v107 uint8 = 128
                        _ = v107
                    default:
                        var v108 int = 9223372036854775807
                        _ = v108
                    }
                    var v109 string = ((("" + "foo bar") + ((("\n\t" + "") + "foo bar") + "hello")) + "")
                    var v110 int8 = v4
                    _ = v109
                    _ = v110
                } else {
                    switch "hello" {
                    case "hello":
                        if true {
                            v5 = float64((float32((3.14159 / v5)) / ((1e10 * 0.1) / 3.14159)))
                        }
                    case "foo bar":
                        v5 = (((v5 - (-2.25 + ((float64((1e10 * 1.5)) + 0.1) / (v5 + v5)))) - 0.1) + (1.5 / ((3.14159 - 1.5) * ((v5 / 1.5) + v5))))
                    case "":
                        if true {
                            v2 -= fi106
                            v111 := []int{fi106, (fi106 + v2), fi106}
                            _ = v111
                        } else {
                            v4 = (^(v4))
                            v5 -= (((((((float64(((1e10 - 1.5) / 0.1)) + 1e10) / 1.0) * (float64(((3.14159 / -2.25) / -2.25)) * v5)) + v5) * (v5 - v5)) - (v5 + v5)) * 3.14159)
                            v112 := GenS11{F0: 255, F1: 1}
                            _ = v112
                        }
                    default:
                        v5 = v5
                    }
                    v2 += fi106
                }
                v113 := GenS12{F0: 147}
                v114 := GenS13{F0: 2, F1: 4294967295, F2: fi106}
                _ = fi106
                _ = v113
                _ = v114
            }
        case "a":
            v115 := []int{(v2 / ((v2) | 1)), v2}
            _ = v115
        case "":
            var v116 bool = (255 >= 255)
            _ = v116
        }
        _ = fi104
        _ = v105
    }
    for fi117 := 0; fi117 < 15; fi117++ {
        v5 -= (v5 + 1e10)
        _ = fi117
    }
    fmt.Println(v2)
    var v118 int32 = -2147483648
    for fi119 := 0; fi119 < 7; fi119++ {
        if true {
            if true {
                var v120 float32 = (1.5 / 3.14159)
                v121 := GenS14{F0: v3, F1: v4, F2: "hello"}
                v122 := GenS15{F0: (fi119 / (((fi119 + 85)) | 1)), F1: 1}
                _ = v120
                _ = v121
                _ = v122
            } else {
                if (true && true) {
                    var v123 uint8 = 0
                    var v124 bool = true
                    var v125 int64 = 0
                    _ = v123
                    _ = v124
                    _ = v125
                } else {
                    v4 = (v4 >> ((249) & 63))
                    fmt.Println(v118)
                }
                switch ("" + "") {
                case "\n\t":
                    for fi126 := 0; fi126 < 17; fi126++ {
                        switch "foo bar" {
                        case "hello":
                            if (v5 < ((((1e10 - v5) - ((v5 / (v5 - ((-2.25 * v5) / ((v5 * float64(float32((v5 / 0.1)))) - 3.14159)))) - float64(1e10))) * 0.0) * (((1.5 / (v5 / (v5 + v5))) + v5) / v5))) {
                                switch ("" + (("hello" + ("a" + "a")) + ("hello" + "\n\t"))) {
                                case "foo bar":
                                    v127 := GenS16{F0: (v118 ^ v118), F1: true, F2: v3}
                                    _ = v127
                                case "\n\t":
                                    if false {
                                        var v128 int64 = 1
                                        var v129 int16 = -585
                                        var v130 uint16 = (v3 ^ (v3 << ((0) & 63)))
                                        _ = v128
                                        _ = v129
                                        _ = v130
                                    }
                                case "":
                                    fmt.Println(fi126)
                                }
                            } else {
                                var v131 uint16 = 1
                                _ = v131
                            }
                        case "a":
                            v3 = (v3 << ((750) & 63))
                        default:
                            v3 = v3
                        }
                        _ = fi126
                    }
                }
                v4 = v4
            }
        }
        v4 += (^(v4))
        _ = fi119
    }
    if false {
        var v132 uint32 = 1
        _ = v132
    } else {
        var v133 int = v2
        v5 -= ((1e10 - float64((1e10 / 1.0))) + (3.14159 + v5))
        v2 -= -2147483648
        _ = v133
    }
    if false {
        v118 = (v118 * (v118 >> ((4294967295) & 63)))
    }
    v1[925] = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v118:", v118)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v118
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
