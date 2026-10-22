package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int8; F2 uint32 }
type GenS3 struct { F0 int16; F1 int16; F2 uint }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int64; F1 string }
type GenS10 struct { F0 uint8; F1 string }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int32; F1 uint64; F2 uint16 }
func genFn0(p0 string, p1 bool, p2 uint) string {
    v6 := GenS1{F0: 100}
    if ((p1 || p1) || false) {
        v7 := []int{4294967295, 2147483647, -52}
        p1 = true
        _ = v7
    } else {
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 float64 = (0.0 - 1e10)
            switch (fi8 ^ fi8) {
            case 900:
                var v10 uint32 = 0
                _ = v10
            case -912:
                v11 := GenS2{F0: 0, F1: -60, F2: 42}
                _ = v11
            default:
                fmt.Println(v9, fi8)
            }
            _ = fi8
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p0
}
func genFn1(p0 uint8, p1 bool) int16 {
    switch -383 {
    case 4294967295:
        if false {
            for fi12 := 0; fi12 < 2; fi12++ {
                p1 = (p1 && false)
                fmt.Println(p1)
                _ = fi12
            }
            var v13 int64 = 72
            p0 = 131
            _ = v13
        }
    }
    var v14 int8 = -1
    _ = p0
    _ = p1
    _ = v14
    return -32768
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]bool)
    var v3 int64 = 42
    v4 := make(map[string]string)
    v5 := []int{9223372036854775807, 0}
    v3 -= (v3 * -20)
    var v15 string = "hello"
    v16 := GenS1{F0: -125}
    switch v15 {
    case "a":
        fmt.Println(v15, v3)
    case "hello":
        var v17 uint16 = 0
        _ = v17
    default:
        var v18 uint = 411
        _ = v18
    }
    if true {
        var v19 bool = false
        fmt.Println(v19)
        var v20 int64 = -46
        _ = v19
        _ = v20
    }
    var v21 int = -2147483648
    var v22 uint8 = 0
    if (!true) {
        for fi23 := 0; fi23 < 17; fi23++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi23
        }
        var v24 uint16 = 2
        _ = v24
    } else {
        var v25 float32 = -2.25
        _ = v25
    }
    var v26 int32 = -748
    fmt.Println(v21, v3, v15)
    for fi27 := 0; fi27 < 11; fi27++ {
        var v28 int8 = 1
        v3 += int64(v21)
        switch (v21 | fi27) {
        case 32767:
            if (!true) {
                v1[""] = false
                v29 := v4[("foo bar" + (v15 + v15))]
                _ = v29
            }
        default:
            for fi30 := 0; fi30 < 4; fi30++ {
                var v31 uint = 1
                v21 -= 199
                if false {
                    var v32 int16 = 0
                    var v33 int32 = (-(v26))
                    _ = v32
                    _ = v33
                }
                var v34 int16 = -32768
                var v35 uint32 = 123
                _ = fi30
                _ = v31
                _ = v34
                _ = v35
            }
        }
        _ = fi27
        _ = v28
    }
    if true {
        switch v21 {
        case -719:
            for fi36 := 0; fi36 < 1; fi36++ {
                if (1 >= -1) {
                    v21 -= v21
                    v22 = v22
                    switch fi36 {
                    case -332:
                        v22 += (v22 >> ((18446744073709551615) & 63))
                    case 32768:
                        var v37 int32 = (-(v26))
                        _ = v37
                    case 127:
                        fmt.Println(v21, v21, v22)
                    }
                } else {
                    switch "a" {
                    case "a":
                        if false {
                            fmt.Println(fi36)
                            var v38 uint64 = 0
                            for fi39 := 0; fi39 < 17; fi39++ {
                                switch -977 {
                                case 959:
                                    switch (v15 + ((("foo bar" + ("foo bar" + (v15 + v15))) + ("a" + (((v15 + v15) + ((v15 + "hello") + ("hello" + v15))) + "foo bar"))) + v15)) {
                                    case "foo bar":
                                        _ = len(v5)
                                    case "\n\t":
                                        var v40 int32 = (v26 & (v26 + -1))
                                        _ = v40
                                    default:
                                        var v41 uint16 = 824
                                        _ = v41
                                    }
                                default:
                                    v42 := GenS3{F0: 32767, F1: genFn1((v22 & (v22 % (((v22 / ((v22) | 1))) | 1))), true), F2: 18446744073709551615}
                                    switch "\n\t" {
                                    case "\n\t":
                                        fmt.Println(v21, fi36, v15)
                                    case "a":
                                        v38 += v38
                                    }
                                    if (false || true) {
                                        v21 += (fi36 & 1)
                                        v2[0] = (42 <= genFn1((v22 >> ((162) & 63)), ((42 != 1) && ((v26 * (v26 >> ((829) & 63))) == (v26 >> ((65535) & 63))))))
                                    } else {
                                        if true {
                                            fmt.Println(fi39)
                                            switch 627 {
                                            case -9223372036854775808:
                                                if (uint64(v22) > v38) {
                                                    switch ((((v15 + (("foo bar" + ((v15 + ("a" + v15)) + v15)) + "foo bar")) + (v15 + "hello")) + v15) + ("foo bar" + "a")) {
                                                    case "hello":
                                                        v43 := GenS2{F0: (v26 / (((v26 % ((v26) | 1))) | 1)), F1: 58, F2: 361}
                                                        _ = v43
                                                    default:
                                                        v22 = 128
                                                    }
                                                    _ = len(v4)
                                                }
                                            case 128:
                                                var v44 uint8 = v22
                                                _ = v44
                                            case 2147483647:
                                                if (!false) {
                                                    v5 = append(v5, fi36)
                                                    if (0.1 != 1e10) {
                                                        fmt.Println(v26, v22, v38)
                                                        v26 -= (v26 + -145)
                                                        fmt.Println(v3, v21)
                                                    } else {
                                                        v26 -= (^(v26))
                                                        switch fi39 {
                                                        case 919:
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                        case 32768:
                                                            var v45 int16 = 553
                                                            _ = v45
                                                        case 4294967295:
                                                            if true {
                                                                if len(v5) > 0 {
                                                                    v5[0] = fi39
                                                                }
                                                                v46 := GenS4{F0: v15, F1: 94}
                                                                v47 := GenS2{F0: int32(v21), F1: -104, F2: 1}
                                                                v48 := GenS5{F0: 111}
                                                                v15 = v15
                                                                _ = v46
                                                                _ = v47
                                                                _ = v48
                                                            } else {
                                                                v22 += v22
                                                                if (false || (false || true)) {
                                                                    v49 := v1[("\n\t" + "a")]
                                                                    var v50 string = "a"
                                                                    var v51 uint = 4294967295
                                                                    _ = v49
                                                                    _ = v50
                                                                    _ = v51
                                                                }
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    v15 = v15
                                                    _ = len(v5)
                                                }
                                            }
                                            var v52 uint16 = 799
                                            _ = v52
                                        }
                                        v26 = (v26 + (v26 ^ v26))
                                    }
                                    _ = v42
                                }
                                _ = fi39
                            }
                            _ = v38
                        }
                    default:
                        var v53 bool = true
                        _ = v53
                    }
                    fmt.Println(v26, v26, v22)
                    if true {
                        for fi54 := 0; fi54 < 3; fi54++ {
                            var v55 uint32 = 196
                            var v56 uint = 65535
                            if true {
                                switch fi36 {
                                case -2147483648:
                                    v26 = v26
                                case 42:
                                    if (!false) {
                                        if true {
                                            v15 = v15
                                        }
                                        v26 = 42
                                    }
                                case 1:
                                    fmt.Println(v21, v22)
                                default:
                                    v57 := GenS4{F0: v15, F1: v56}
                                    _ = v57
                                }
                                if ((!false) || (v3 != v3)) {
                                    if ((v55 * 4294967295) >= (^(v55))) {
                                        _ = len(v2)
                                    } else {
                                        if true {
                                            v55 = v55
                                        } else {
                                            v15 = (v15 + (("\n\t" + v15) + ""))
                                        }
                                        v22 = v22
                                        _ = len(v2)
                                    }
                                    delete(v2, 828)
                                    var v58 int16 = -409
                                    _ = v58
                                } else {
                                    var v59 int32 = 2147483647
                                    _ = v59
                                }
                            } else {
                                switch ("" + v15) {
                                case "":
                                    v26 += -204
                                }
                                if false {
                                    if true {
                                        _ = len(v5)
                                        if (!true) {
                                            var v60 int32 = v26
                                            _ = v60
                                        }
                                        v3 = v3
                                    } else {
                                        v15 = "\n\t"
                                        if false {
                                            var v61 int = -2147483648
                                            v21 = (^(fi54))
                                            _ = v61
                                        }
                                        if ((!(659 >= 42)) || (uint(v22) < v56)) {
                                            if true {
                                                var v62 int8 = -1
                                                _ = v62
                                            }
                                            var v63 float32 = float32(-2.25)
                                            var v64 int8 = 127
                                            _ = v63
                                            _ = v64
                                        } else {
                                            var v65 uint = v56
                                            var v66 int = (fi36 << ((uint(v55)) & 63))
                                            v5 = append(v5, fi54)
                                            _ = v65
                                            _ = v66
                                        }
                                    }
                                    if false {
                                        v67 := GenS6{F0: 462}
                                        switch v15 {
                                        case "\n\t":
                                            var v68 uint8 = v22
                                            _ = v68
                                        case "hello":
                                            v21 = v21
                                        default:
                                            if true {
                                                v21 = fi54
                                                var v69 uint = (^(v56))
                                                v70 := GenS7{F0: v21}
                                                _ = v69
                                                _ = v70
                                            }
                                        }
                                        var v71 uint = (^(v56))
                                        _ = v67
                                        _ = v71
                                    } else {
                                        if false {
                                            v1[v15] = false
                                            v22 = (v22 << (((v56 << (((v56 % (((v56 + 968)) | 1))) & 63))) & 63))
                                        }
                                        var v72 int32 = 2147483647
                                        var v73 int32 = v72
                                        _ = v72
                                        _ = v73
                                    }
                                } else {
                                    v74 := GenS8{F0: (v22 * uint8(v55))}
                                    fmt.Println(fi54)
                                    v75 := GenS9{F0: 9223372036854775807, F1: ("foo bar" + v15)}
                                    _ = v74
                                    _ = v75
                                }
                                v22 = 201
                                if true {
                                    if (3.14159 <= (float64(((1e10 - float32((((1.5 / 1e10) / -2.25) + float64(float32(0.1))))) + -2.25)) + 3.14159)) {
                                        v15 = v15
                                        v76 := GenS10{F0: (^(v22)), F1: ""}
                                        v77 := GenS2{F0: v26, F1: 127, F2: v55}
                                        _ = v76
                                        _ = v77
                                    } else {
                                        v26 = v26
                                        if true {
                                            v56 -= uint(v22)
                                        } else {
                                            fmt.Println(fi36)
                                        }
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                    var v78 float64 = (1.5 * 1.5)
                                    v78 -= 1e10
                                    v26 = 42
                                    _ = v78
                                } else {
                                    _ = len(v4)
                                    if true {
                                        _ = len(v5)
                                        v79 := GenS1{F0: 1}
                                        v80 := v4[v15]
                                        _ = v79
                                        _ = v80
                                    }
                                }
                                _ = len(v4)
                            }
                            _ = fi54
                            _ = v55
                            _ = v56
                        }
                        v1[(v15 + v15)] = ((true && (2 == 65535)) && true)
                    } else {
                        fmt.Println(fi36, fi36, v22)
                    }
                }
                _ = fi36
            }
        case 32768:
            for fi81 := 0; fi81 < 3; fi81++ {
                var v82 uint8 = v22
                v15 = v15
                _ = fi81
                _ = v82
            }
        case -134:
            v3 -= v3
        }
        if true {
            if (false || false) {
                switch v21 {
                case -9223372036854775808:
                    if false {
                        var v83 bool = true
                        if (569 == 859) {
                            var v84 uint8 = (v22 - 128)
                            for fi85 := 0; fi85 < 4; fi85++ {
                                if (1 <= 1) {
                                    fmt.Println(fi85, v21)
                                    for fi86 := 0; fi86 < 18; fi86++ {
                                        v87 := GenS7{F0: (fi85 + fi85)}
                                        if (!v83) {
                                            var v88 uint = 4294967295
                                            var v89 float32 = (1e10 * 1.5)
                                            _ = v88
                                            _ = v89
                                        } else {
                                            if (-9 <= -1) {
                                                var v90 string = "foo bar"
                                                var v91 uint64 = 0
                                                fmt.Println(v3, v91, v22)
                                                _ = v90
                                                _ = v91
                                            } else {
                                                var v92 uint8 = 0
                                                _ = v92
                                            }
                                            var v93 bool = v83
                                            switch v15 {
                                            case "\n\t":
                                                if ((68 >= 2) || ((4294967295 <= 42) && ((v3 & 0) <= (v3 << ((255) & 63))))) {
                                                    var v94 uint32 = 65535
                                                    _ = v94
                                                }
                                            default:
                                                if (!(-937 == (v26 & -759))) {
                                                    switch v15 {
                                                    case "a":
                                                        if len(v5) > 0 {
                                                            v5[0] = v21
                                                        }
                                                    case "foo bar":
                                                        if (1 < -66) {
                                                            var v95 int16 = 1000
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            _ = v95
                                                        }
                                                    }
                                                    var v96 int16 = 482
                                                    _ = v96
                                                } else {
                                                    fmt.Println(v15)
                                                }
                                            }
                                            _ = v93
                                        }
                                        var v97 uint32 = 2
                                        _ = fi86
                                        _ = v87
                                        _ = v97
                                    }
                                } else {
                                    v98 := GenS3{F0: -32768, F1: -291, F2: 255}
                                    _ = v98
                                }
                                _ = fi85
                            }
                            _ = v84
                        }
                        _ = v83
                    }
                }
                var v99 float32 = (1.5 - (0.1 - (float32(1.5) * ((1.5 - float32(1e10)) - 3.14159))))
                _ = v99
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            for fi100 := 0; fi100 < 15; fi100++ {
                for fi101 := 0; fi101 < 3; fi101++ {
                    var v102 int8 = 0
                    var v103 float64 = float64(((3.14159 - 1.5) - 1.5))
                    switch "a" {
                    case "\n\t":
                        if true {
                            v104 := GenS4{F0: ((v15 + (v15 + v15)) + ((("hello" + v15) + (v15 + (((("hello" + ((v15 + v15) + ("\n\t" + v15))) + "a") + v15) + (v15 + (v15 + ("hello" + "")))))) + "foo bar")), F1: 283}
                            _ = v104
                        } else {
                            v15 = ("\n\t" + ("" + "\n\t"))
                            if len(v5) > 0 {
                                v5[0] = (fi100 / ((v21) | 1))
                            }
                            var v105 uint = 2
                            v106 := GenS11{F0: v102}
                            _ = v105
                            _ = v106
                        }
                    default:
                        var v107 string = v15
                        _ = v107
                    }
                    v108 := v1[v15]
                    var v109 int8 = v102
                    var v110 uint64 = 133
                    _ = fi101
                    _ = v102
                    _ = v103
                    _ = v108
                    _ = v109
                    _ = v110
                }
                if true {
                    var v111 uint = 2
                    v112 := GenS9{F0: v3, F1: v15}
                    _ = v111
                    _ = v112
                }
                v2[776] = (false && false)
                _ = fi100
            }
        } else {
            var v113 uint32 = 2147483647
            switch v21 {
            case 683:
                var v114 float32 = 1.5
                _ = v114
            case 32767:
                v115 := GenS12{F0: (-(v26)), F1: 0, F2: 0}
                _ = v115
            }
            var v116 int16 = -1
            _ = v113
            _ = v116
        }
    }
    switch "\n\t" {
    case "":
        v26 = -729
    case "\n\t":
        v117 := GenS10{F0: (v22 | 2), F1: ""}
        _ = v117
    default:
        fmt.Println(v15, v3)
    }
    v5 = append(v5, v21)
    var v118 uint16 = 2
    for fi119 := 0; fi119 < 3; fi119++ {
        if ((!(!false)) || (1.5 < (3.14159 / -2.25))) {
            if false {
                var v120 int16 = genFn1((v22 ^ (v22 >> ((366) & 63))), ((v3 + (v3 - (-(v3)))) < (v3 ^ v3)))
                _ = v120
            } else {
                if false {
                    var v121 float64 = (1e10 / 0.1)
                    var v122 int32 = (v26 & v26)
                    _ = v121
                    _ = v122
                } else {
                    if ((((v22 < 255) && (!((42 < 20) && true))) && (false || false)) || (0.1 > 3.14159)) {
                        var v123 uint16 = 42
                        _ = v123
                    }
                    switch 9223372036854775807 {
                    case 1:
                        fmt.Println(v26, fi119, v22)
                    case 4294967295:
                        var v124 uint8 = v22
                        _ = v124
                    default:
                        _ = len(v1)
                    }
                }
            }
            for fi125 := 0; fi125 < 18; fi125++ {
                v126 := GenS5{F0: 1}
                var v127 int = fi125
                if (((-2.25 + (1.5 / 1e10)) - (0.1 * (1e10 - (((((-2.25 / ((0.0 - ((1e10 - ((-2.25 / 1.0) * (1.5 + float32(1e10)))) * -2.25)) + 0.0)) + (float32(1e10) * (float32((-2.25 + 0.1)) + 1e10))) + (0.0 - 0.1)) * 1e10) * (0.1 * -2.25))))) <= 3.14159) {
                    var v128 int = (fi125 / (((fi119 + int(v118))) | 1))
                    v4[v15] = (("foo bar" + "hello") + v15)
                    _ = v128
                }
                _ = fi125
                _ = v126
                _ = v127
            }
        }
        _ = fi119
    }
    var v129 bool = (!true)
    v130 := GenS11{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v118:", v118)
    fmt.Println("v129:", v129)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v118
    _ = v129
    _ = v130
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
