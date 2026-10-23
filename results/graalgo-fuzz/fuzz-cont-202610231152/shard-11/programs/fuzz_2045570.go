package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint; F1 int; F2 int16 }
type GenS6 struct { F0 int64; F1 uint64; F2 int32 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int16; F1 uint8; F2 bool }
type GenS9 struct { F0 int64; F1 uint16 }
type GenS10 struct { F0 uint; F1 uint64 }
type GenS11 struct { F0 int8; F1 uint64; F2 string }
type GenS12 struct { F0 int8; F1 uint16 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 int; F1 int }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 int; F1 int8; F2 uint16 }
type GenS19 struct { F0 int32; F1 uint32 }
type GenS20 struct { F0 uint }
type GenS21 struct { F0 uint }
type GenS22 struct { F0 int64 }
type GenS23 struct { F0 int32 }
func genFn0(p0 uint16, p1 int64, p2 int64) uint32 {
    var v4 string = ("a" + "a")
    if (float32(1e10) == 3.14159) {
        if (p0 >= 1) {
            if false {
                v5 := []int{-1, 921, -9223372036854775808, 42}
                switch ((v4 + v4) + v4) {
                case "a":
                    if (9223372036854775807 == 430) {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        var v6 uint64 = 18446744073709551615
                        var v7 uint64 = (v6 << ((44) & 63))
                        var v8 int16 = 32767
                        _ = v6
                        _ = v7
                        _ = v8
                    }
                default:
                    if (385 >= 42) {
                        p0 = (p0 - 32767)
                        var v9 uint16 = p0
                        v10 := GenS1{F0: 32767, F1: 56, F2: p0}
                        _ = v10
                        _ = v9
                    } else {
                        v4 = "hello"
                        if false {
                            p2 = p1
                        } else {
                            var v11 uint = 18446744073709551615
                            if (980 != 2) {
                                var v12 string = ("foo bar" + (((v4 + (("a" + v4) + v4)) + ("" + "")) + "a"))
                                _ = len(v5)
                                _ = v12
                            } else {
                                p0 = 163
                                if true {
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v13 int32 = -2147483648
                                        _ = v13
                                    case "a":
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                    p0 = 201
                                    var v14 int16 = 1
                                    _ = v14
                                }
                                if len(v5) > 0 {
                                    v5[0] = 553
                                }
                            }
                            p2 += (p1 & int64(v11))
                            _ = v11
                        }
                        fmt.Println(v4)
                    }
                }
                if (0.0 <= 1e10) {
                    if true {
                        var v15 bool = (false && (!true))
                        v16 := GenS2{F0: 0}
                        _ = v15
                        _ = v16
                    }
                    if false {
                        var v17 uint8 = 153
                        if (2147483647 < 2147483647) {
                            v18 := GenS3{F0: 32767}
                            fmt.Println(v4, p1, v17)
                            v19 := GenS4{F0: "foo bar"}
                            _ = v18
                            _ = v19
                        }
                        v17 -= v17
                        _ = v17
                    }
                    v20 := GenS5{F0: 18446744073709551615, F1: 32768, F2: -7}
                    _ = v20
                } else {
                    fmt.Println(v4, p0, p2)
                    switch "hello" {
                    case "hello":
                        fmt.Println(p1, p2, v4)
                    case "":
                        fmt.Println(p0)
                    case "\n\t":
                        if false {
                            v21 := GenS6{F0: (^(p2)), F1: 4294967295, F2: 2147483647}
                            _ = v21
                        }
                    default:
                        v22 := GenS1{F0: (p0 ^ p0), F1: -103, F2: 495}
                        _ = v22
                    }
                    p0 = p0
                }
                _ = v5
            } else {
                v23 := GenS2{F0: 91}
                _ = v23
            }
            fmt.Println(p2, p0)
        }
        p2 -= p1
        v24 := GenS7{F0: 42}
        _ = v24
    } else {
        if (42 <= 42) {
            v25 := []int{1}
            var v26 uint = 42
            var v27 uint16 = 255
            _ = v25
            _ = v26
            _ = v27
        } else {
            v28 := GenS1{F0: p0, F1: -128, F2: p0}
            _ = v28
        }
    }
    switch (v4 + (v4 + "hello")) {
    case "hello":
        switch 42 {
        case 57:
            v29 := []int{-9223372036854775808}
            _ = v29
        case -858:
            p1 -= p2
        }
    case "foo bar":
        var v30 float32 = ((0.1 + (-2.25 + ((0.0 - (1.5 * (3.14159 - (1e10 * (float32(1e10) * float32(0.1)))))) - 0.1))) + float32(3.14159))
        _ = v30
    default:
        var v31 uint32 = 1
        _ = v31
    }
    p0 += p0
    var v32 string = v4
    _ = p0
    _ = p1
    _ = p2
    _ = v32
    _ = v4
    return 740
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 478
    var v3 string = "\n\t"
    switch 127 {
    case -314:
        switch "\n\t" {
        case "foo bar":
            var v33 bool = false
            _ = v33
        case "\n\t":
            for fi34 := 0; fi34 < 1; fi34++ {
                if false {
                    fmt.Println(fi34)
                }
                delete(v1, 268)
                v35 := GenS8{F0: -848, F1: 128, F2: (-9223372036854775808 > -483)}
                _ = fi34
                _ = v35
            }
        case "":
            if (-1 >= 0) {
                var v36 int64 = -1
                for fi37 := 0; fi37 < 15; fi37++ {
                    v38 := GenS9{F0: v36, F1: 32767}
                    for fi39 := 0; fi39 < 10; fi39++ {
                        v40 := GenS8{F0: 1000, F1: 0, F2: (!(0 < 0))}
                        v41 := GenS4{F0: v3}
                        if true {
                            if false {
                                if (true && true) {
                                    v42 := GenS10{F0: 461, F1: v2}
                                    _ = v42
                                } else {
                                    _ = len(v1)
                                }
                                var v43 int = fi39
                                if false {
                                    v44 := GenS1{F0: 65535, F1: int8(fi37), F2: 42}
                                    v45 := GenS6{F0: v36, F1: (v2 >> ((581) & 63)), F2: 0}
                                    _ = v44
                                    _ = v45
                                } else {
                                    fmt.Println(fi37)
                                    switch 1 {
                                    case -1:
                                        if false {
                                            switch v3 {
                                            case "hello":
                                                v2 += (v2 + v2)
                                            case "a":
                                                v46 := GenS9{F0: v36, F1: 2}
                                                _ = v46
                                            }
                                            if false {
                                                v3 = v3
                                                v3 = ("foo bar" + v3)
                                                switch fi37 {
                                                case 128:
                                                    v47 := GenS6{F0: (v36 | 42), F1: v2, F2: -2147483648}
                                                    _ = v47
                                                case 42:
                                                    if (true || (!true)) {
                                                        fmt.Println(v3, fi37)
                                                        var v48 int = fi37
                                                        _ = v48
                                                    }
                                                default:
                                                    var v49 uint8 = 255
                                                    _ = v49
                                                }
                                            } else {
                                                switch v3 {
                                                case "\n\t":
                                                    var v50 float32 = float32(1.5)
                                                    _ = v50
                                                case "hello":
                                                    delete(v1, v43)
                                                    var v51 int = 0
                                                    v52 := GenS11{F0: 42, F1: v2, F2: "a"}
                                                    switch "a" {
                                                    case "foo bar":
                                                        switch -693 {
                                                        case -2147483648:
                                                            fmt.Println(v51, v2)
                                                        case 9223372036854775807:
                                                            v53 := GenS9{F0: v36, F1: 42}
                                                            _ = v53
                                                        default:
                                                            v54 := GenS12{F0: -1, F1: 0}
                                                            _ = v54
                                                        }
                                                    case "\n\t":
                                                        var v55 uint64 = (v2 + v2)
                                                        _ = v55
                                                    case "a":
                                                        if (42 > -72) {
                                                            v51 = fi39
                                                        } else {
                                                            delete(v1, fi37)
                                                            v51 += v51
                                                            v56 := GenS13{F0: (v2 | (v2 - (v2 | (v2 ^ v2))))}
                                                            _ = v56
                                                        }
                                                    default:
                                                        fmt.Println(v36)
                                                        v57 := GenS14{F0: 18446744073709551615}
                                                        _ = v57
                                                    }
                                                    fmt.Println(v2, v43)
                                                    v58 := []int{fi39}
                                                    _ = v51
                                                    _ = v52
                                                    _ = v58
                                                case "a":
                                                    switch fi39 {
                                                    case 9223372036854775807:
                                                        switch fi37 {
                                                        case -705:
                                                            var v59 int32 = 400
                                                            _ = v59
                                                        }
                                                    case -9223372036854775808:
                                                        var v60 int = (fi39 & (fi39 * 127))
                                                        var v61 bool = (!true)
                                                        _ = v60
                                                        _ = v61
                                                    default:
                                                        var v62 int = fi39
                                                        _ = v62
                                                    }
                                                }
                                                var v63 uint16 = 32767
                                                switch ("hello" + "a") {
                                                case "":
                                                    if true {
                                                        if true {
                                                            if false {
                                                                var v64 uint16 = v63
                                                                v65 := []int{}
                                                                var v66 string = ((("a" + v3) + "") + v3)
                                                                _ = v64
                                                                _ = v65
                                                                _ = v66
                                                            }
                                                            _ = len(v1)
                                                            fmt.Println(v2)
                                                        } else {
                                                            var v67 uint8 = 42
                                                            _ = v67
                                                        }
                                                        var v68 int = fi37
                                                        _ = v68
                                                    }
                                                case "hello":
                                                    fmt.Println(fi37, fi37, v2)
                                                }
                                                _ = v63
                                            }
                                        }
                                        var v69 float32 = 0.1
                                        _ = v69
                                    default:
                                        if (130 > 1) {
                                            switch "a" {
                                            case "\n\t":
                                                v70 := GenS15{F0: 65535}
                                                _ = v70
                                            case "":
                                                if (!true) {
                                                    switch v43 {
                                                    case 0:
                                                        fmt.Println(fi39)
                                                    default:
                                                        v2 = 0
                                                    }
                                                }
                                            case "foo bar":
                                                if (false || false) {
                                                    if (2 != 0) {
                                                        v71 := GenS15{F0: 42}
                                                        fmt.Println(v36, v3, v3)
                                                        if (false && false) {
                                                            var v72 uint64 = v2
                                                            var v73 int = v43
                                                            v36 += 384
                                                            _ = v72
                                                            _ = v73
                                                        } else {
                                                            v74 := []int{fi37, fi39}
                                                            _ = v74
                                                        }
                                                        _ = v71
                                                    } else {
                                                        v75 := GenS16{F0: (fi37 ^ fi39), F1: 128}
                                                        _ = v75
                                                    }
                                                    fmt.Println(v43, fi39, fi39)
                                                }
                                            default:
                                                var v76 int16 = 1
                                                _ = v76
                                            }
                                            v77 := GenS5{F0: 255, F1: (v43 ^ 0), F2: 32767}
                                            v3 = v3
                                            var v78 bool = ((false && (float64((1.5 * (float32((0.1 + 3.14159)) + (((3.14159 * -2.25) / (3.14159 + (0.0 * (3.14159 * 3.14159)))) / 0.1)))) != 1.5)) && (1e10 < (3.14159 * 3.14159)))
                                            v78 = v78
                                            _ = v77
                                            _ = v78
                                        } else {
                                            switch (v43 ^ -793) {
                                            case 128:
                                                if (!true) {
                                                    var v79 uint = 146
                                                    var v80 float32 = 3.14159
                                                    _ = v79
                                                    _ = v80
                                                } else {
                                                    switch 225 {
                                                    case 9223372036854775807:
                                                        fmt.Println(v36, v3, v2)
                                                    }
                                                    v81 := []int{42}
                                                    _ = v81
                                                }
                                            }
                                            var v82 float64 = (1.5 - float64((0.1 + 0.1)))
                                            _ = v82
                                        }
                                    }
                                    if false {
                                        v3 = (v3 + ((("foo bar" + (("" + v3) + (v3 + "hello"))) + ((v3 + (v3 + "a")) + (("foo bar" + "foo bar") + v3))) + ("" + v3)))
                                        v3 = ((v3 + v3) + (v3 + v3))
                                        v83 := GenS5{F0: 1, F1: (v43 | (fi37 >> ((447) & 63))), F2: 593}
                                        _ = v83
                                    } else {
                                        switch fi39 {
                                        case 4294967295:
                                            if true {
                                                fmt.Println(v36, v3)
                                            }
                                            switch v43 {
                                            case 32767:
                                                switch "a" {
                                                case "foo bar":
                                                    v43 -= fi37
                                                default:
                                                    v3 = ((v3 + "hello") + "hello")
                                                }
                                            case 0:
                                                v2 -= v2
                                            case 158:
                                                if ((v2 + (^(v2))) >= v2) {
                                                    fmt.Println(v2, v43, v2)
                                                    var v84 uint32 = 2147483648
                                                    if true {
                                                        v3 = v3
                                                        var v85 uint32 = genFn0(uint16(v84), (v36 ^ v36), (v36 - (v36 - -9223372036854775808)))
                                                        _ = v85
                                                    } else {
                                                        v84 = 2147483648
                                                        v86 := []int{}
                                                        var v87 int = (fi39 + fi39)
                                                        v88 := GenS11{F0: 72, F1: v2, F2: v3}
                                                        _ = v86
                                                        _ = v87
                                                        _ = v88
                                                    }
                                                    _ = v84
                                                } else {
                                                    v36 = (-(v36))
                                                    if (13 < 28) {
                                                        v89 := []int{127, 32767, v43, (fi37 - fi37)}
                                                        v43 = v43
                                                        switch fi39 {
                                                        case -9223372036854775808:
                                                            if (0 <= -1) {
                                                                v43 = 32768
                                                                v2 = 4294967295
                                                                v2 += 981
                                                            } else {
                                                                var v90 int32 = 491
                                                                v90 = (^(v90))
                                                                _ = v90
                                                            }
                                                        case 1:
                                                            v91 := GenS5{F0: 65535, F1: fi37, F2: 422}
                                                            _ = v91
                                                        case 42:
                                                            switch "a" {
                                                            case "":
                                                                var v92 uint32 = 65535
                                                                _ = v92
                                                            case "hello":
                                                                v93 := GenS17{F0: false}
                                                                _ = v93
                                                            }
                                                        }
                                                        _ = v89
                                                    }
                                                }
                                            default:
                                                if true {
                                                    v94 := GenS17{F0: false}
                                                    var v95 uint16 = 65535
                                                    var v96 uint32 = uint32(v2)
                                                    _ = v94
                                                    _ = v95
                                                    _ = v96
                                                } else {
                                                    v97 := GenS17{F0: false}
                                                    var v98 uint8 = 255
                                                    _ = v97
                                                    _ = v98
                                                }
                                            }
                                            var v99 float32 = 3.14159
                                            var v100 uint = 255
                                            var v101 int8 = -11
                                            _ = v100
                                            _ = v101
                                            _ = v99
                                        case 511:
                                            v2 += 1
                                        }
                                        var v102 int8 = 127
                                        _ = v102
                                    }
                                    if (false && true) {
                                        if true {
                                            v103 := []int{9223372036854775807, fi37, fi39}
                                            _ = v103
                                        } else {
                                            v36 += (v36 % (((v36 + 705)) | 1))
                                            if ((fi39 <= (fi37 / ((v43) | 1))) || false) {
                                                var v104 bool = false
                                                _ = v104
                                            } else {
                                                switch ("hello" + v3) {
                                                case "":
                                                    var v105 string = v3
                                                    var v106 int16 = -32768
                                                    v107 := GenS2{F0: -35}
                                                    _ = v105
                                                    _ = v106
                                                    _ = v107
                                                }
                                                _ = len(v1)
                                                v36 = v36
                                                v2 = v2
                                                var v108 int32 = 1
                                                fmt.Println(v43)
                                                _ = v108
                                            }
                                            var v109 int64 = v36
                                            v43 = (v43 + (v43 + (v43 << ((444) & 63))))
                                            _ = v109
                                        }
                                        if false {
                                            if true {
                                                v3 = (v3 + v3)
                                            }
                                        } else {
                                            var v110 float64 = float64(float32(((((1.5 - -2.25) + (1e10 / 1e10)) - ((0.1 / -2.25) * (((0.0 + (float64(float32(-2.25)) * (((0.1 / 1.0) / (0.1 / 1e10)) + -2.25))) + 0.0) / 0.1))) - ((0.0 + 1e10) + float64((1e10 - ((0.1 + (1e10 + (3.14159 * 0.0))) + 3.14159)))))))
                                            v43 = fi37
                                            _ = v110
                                        }
                                        var v111 bool = (!(false || true))
                                        _ = v111
                                    }
                                    fmt.Println(fi37)
                                }
                                _ = v43
                            }
                            v2 = 18446744073709551615
                            var v112 int32 = 1
                            _ = v112
                        } else {
                            var v113 int32 = 42
                            _ = v113
                        }
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                    if ((fi37 * (^(fi37))) != 2147483647) {
                        for fi114 := 0; fi114 < 12; fi114++ {
                            fmt.Println(v36, v3, fi114)
                            _ = fi114
                        }
                    } else {
                        fmt.Println(v3, v2)
                    }
                    _ = fi37
                    _ = v38
                }
                v36 = -9223372036854775808
                _ = v36
            }
        default:
            fmt.Println(v2)
        }
    default:
        v115 := GenS18{F0: 1, F1: 42, F2: 248}
        _ = v115
    }
    v116 := []int{}
    for fi117 := 0; fi117 < 10; fi117++ {
        switch "hello" {
        case "\n\t":
            switch "hello" {
            case "":
                if false {
                    switch ("hello" + v3) {
                    case "a":
                        v3 = (v3 + "a")
                    default:
                        v2 = v2
                    }
                    var v118 uint = 2
                    v119 := GenS16{F0: 9223372036854775807, F1: fi117}
                    _ = v118
                    _ = v119
                } else {
                    var v120 uint64 = 0
                    var v121 bool = ((211 == 128) || false)
                    _ = v120
                    _ = v121
                }
            case "\n\t":
                v2 += (v2 & 1)
            default:
                v2 -= (v2 / ((v2) | 1))
            }
        case "":
            var v122 float64 = 1e10
            _ = v122
        case "foo bar":
            v123 := GenS1{F0: 1, F1: 42, F2: 9}
            _ = v123
        default:
            var v124 int8 = 42
            _ = v124
        }
        v3 = (v3 + (v3 + (v3 + v3)))
        _ = fi117
    }
    var v125 float32 = 1e10
    var v126 uint = 4294967295
    for fi127 := 0; fi127 < 12; fi127++ {
        v126 = (^(v126))
        if false {
            fmt.Println(v2)
            fmt.Println(v125, v2, fi127)
        }
        _ = fi127
    }
    switch (v3 + ("a" + (((v3 + "a") + "hello") + ("foo bar" + v3)))) {
    case "foo bar":
        var v128 uint32 = 1
        _ = v128
    case "a":
        v129 := GenS19{F0: 1, F1: 566}
        _ = v129
    }
    for fi130 := 0; fi130 < 5; fi130++ {
        v131 := GenS20{F0: v126}
        _ = fi130
        _ = v131
    }
    if false {
        var v132 uint8 = 42
        _ = v132
    }
    switch -9223372036854775808 {
    case 127:
        for fi133 := 0; fi133 < 11; fi133++ {
            switch (v3 + v3) {
            case "":
                if ((!true) || ((v3 + (v3 + "\n\t")) <= (v3 + "foo bar"))) {
                    for fi134 := 0; fi134 < 14; fi134++ {
                        var v135 uint = v126
                        _ = fi134
                        _ = v135
                    }
                    v136 := GenS21{F0: uint(v2)}
                    if len(v116) >= 2 {
                        v116 = v116[1:len(v116)]
                    }
                    _ = v136
                } else {
                    fmt.Println(fi133, fi133, v2)
                    var v137 int8 = -79
                    _ = v137
                }
            case "\n\t":
                if true {
                    delete(v1, (-(fi133)))
                }
            default:
                v138 := GenS22{F0: 479}
                _ = v138
            }
            fmt.Println(v126, v126, v3)
            v1[(fi133 << (((^(v126))) & 63))] = (((("" + v3) + "hello") + (("hello" + "a") + v3)) + ("" + (v3 + "a")))
            _ = fi133
        }
    case 1:
        v125 += v125
    }
    var v139 int32 = 0
    v140 := v1[1]
    switch "a" {
    case "foo bar":
        var v141 float32 = v125
        _ = v141
    case "hello":
        v142 := GenS22{F0: -11}
        _ = v142
    case "\n\t":
        var v143 string = v3
        _ = v143
    }
    v144 := GenS23{F0: v139}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v116:", len(v116), v116)
    fmt.Println("v125:", v125)
    fmt.Println("v126:", v126)
    fmt.Println("v139:", v139)
    fmt.Println("v140:", v140)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v116
    _ = v125
    _ = v126
    _ = v139
    _ = v140
    _ = v144
    _ = v2
    _ = v3
}
