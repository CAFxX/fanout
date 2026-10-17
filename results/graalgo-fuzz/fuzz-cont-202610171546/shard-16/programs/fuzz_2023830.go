package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 uint; F1 int64; F2 bool }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS11 struct { F0 int16; F1 uint }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 string; F1 int8 }
type GenS14 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS15 struct { F0 int8; F1 int8 }
type GenS16 struct { F0 int; F1 uint16; F2 uint64 }
type GenS17 struct { F0 uint8; F1 int; F2 int8 }
type GenS18 struct { F0 string; F1 bool; F2 uint8 }
type GenS19 struct { F0 uint16 }
type GenS20 struct { F0 int64 }
type GenS21 struct { F0 uint64 }
type GenS22 struct { F0 uint; F1 uint64; F2 string }
func genFn0(p0 int16) uint32 {
    v4 := []int{}
    fmt.Println(p0)
    p0 = (p0 % ((p0) | 1))
    _ = p0
    _ = v4
    return 328
}
func genFn1(p0 uint) uint32 {
    if true {
        p0 -= (p0 & (p0 + 255))
        p0 -= (^(p0))
    }
    var v5 float32 = 3.14159
    _ = p0
    _ = v5
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 401
    v3 := []int{127, 1, 32767}
    switch 9223372036854775807 {
    case 32768:
        if true {
            v2 -= 42
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch "" {
            case "hello":
                fmt.Println(v2)
            case "foo bar":
                if (false || (v2 != 4294967295)) {
                    delete(v1, "foo bar")
                } else {
                    v6 := GenS1{F0: 95, F1: 418, F2: -9223372036854775808}
                    for fi7 := 0; fi7 < 1; fi7++ {
                        for fi8 := 0; fi8 < 2; fi8++ {
                            var v9 float32 = 3.14159
                            var v10 uint = 1
                            fmt.Println(fi7)
                            _ = fi8
                            _ = v10
                            _ = v9
                        }
                        if false {
                            if true {
                                var v11 uint16 = 191
                                for fi12 := 0; fi12 < 6; fi12++ {
                                    fmt.Println(fi12)
                                    if true {
                                        v13 := GenS2{F0: "hello", F1: true}
                                        v3 = append(v3, -2147483648)
                                        var v14 uint64 = v2
                                        _ = v13
                                        _ = v14
                                    }
                                    var v15 uint16 = (v11 & (v11 >> ((632) & 63)))
                                    v2 = 1
                                    fmt.Println(fi7, v15)
                                    _ = fi12
                                    _ = v15
                                }
                                _ = v11
                            }
                        }
                        for fi16 := 0; fi16 < 2; fi16++ {
                            v2 -= 2
                            _ = fi16
                        }
                        _ = fi7
                    }
                    fmt.Println(v2)
                    _ = v6
                }
            case "a":
                for fi17 := 0; fi17 < 8; fi17++ {
                    if (-2.25 >= (1e10 * (((float32((-2.25 / ((-2.25 / 1e10) * 1e10))) * 3.14159) / float32((-2.25 / 1.5))) / 1.0))) {
                        if (-1 != 1) {
                            _ = len(v3)
                            var v18 int32 = 2147483647
                            v19 := GenS3{F0: 42, F1: int(v2)}
                            _ = v18
                            _ = v19
                        }
                    } else {
                        if false {
                            if true {
                                for fi20 := 0; fi20 < 19; fi20++ {
                                    if (!(1 > -244)) {
                                        v21 := GenS1{F0: 128, F1: 46, F2: -1}
                                        switch "foo bar" {
                                        case "":
                                            switch (fi17 | -197) {
                                            case -9223372036854775808:
                                                var v22 bool = (3.14159 >= (float32(0.1) + 3.14159))
                                                _ = v22
                                            case 2147483647:
                                                v2 = 18446744073709551615
                                                v23 := v1["hello"]
                                                _ = v23
                                            case 127:
                                                if (2147483647 < 0) {
                                                    if (!(!(2147483647 > 135))) {
                                                        var v24 int32 = -777
                                                        v24 += v24
                                                        v25 := GenS2{F0: "foo bar", F1: true}
                                                        fmt.Println(v24, fi17, v2)
                                                        v24 = v24
                                                        if (false || false) {
                                                            _ = len(v1)
                                                            v26 := GenS4{F0: "hello", F1: 0}
                                                            switch ("hello" + "foo bar") {
                                                            case "foo bar":
                                                                if true {
                                                                    delete(v1, "foo bar")
                                                                    switch 32767 {
                                                                    case 9223372036854775807:
                                                                        v24 = (v24 + 0)
                                                                    case 4294967295:
                                                                        v27 := v1["hello"]
                                                                        _ = v27
                                                                    case 32767:
                                                                        if len(v3) > 0 {
                                                                            v3[0] = (fi20 ^ (^(fi17)))
                                                                        }
                                                                    default:
                                                                        v28 := v1[("" + "a")]
                                                                        _ = v28
                                                                    }
                                                                } else {
                                                                    v2 = v2
                                                                    _ = len(v3)
                                                                    v29 := GenS3{F0: 71, F1: (-(fi17))}
                                                                    var v30 int16 = 1
                                                                    _ = v29
                                                                    _ = v30
                                                                }
                                                            case "hello":
                                                                v31 := GenS5{F0: (v2 >= 0), F1: -1}
                                                                _ = v31
                                                            default:
                                                                var v32 float64 = 0.1
                                                                _ = v32
                                                            }
                                                            var v33 uint16 = 933
                                                            _ = v26
                                                            _ = v33
                                                        }
                                                        _ = v24
                                                        _ = v25
                                                    } else {
                                                        var v34 uint8 = 24
                                                        var v35 int16 = -1
                                                        var v36 int64 = 409
                                                        _ = v34
                                                        _ = v35
                                                        _ = v36
                                                    }
                                                    if true {
                                                        v2 -= v2
                                                        var v37 uint8 = 166
                                                        var v38 int32 = -1
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                } else {
                                                    var v39 int = (fi20 << ((1) & 63))
                                                    _ = v39
                                                }
                                            }
                                        case "foo bar":
                                            var v40 int64 = -9223372036854775808
                                            _ = v40
                                        case "\n\t":
                                            v2 += v2
                                        }
                                        _ = v21
                                    } else {
                                        var v41 float32 = 1.5
                                        _ = v41
                                    }
                                    v2 = (v2 * v2)
                                    _ = fi20
                                }
                                switch "\n\t" {
                                case "a":
                                    for fi42 := 0; fi42 < 10; fi42++ {
                                        if true {
                                            if (1 >= 1) {
                                                var v43 uint8 = 1
                                                v44 := GenS6{F0: 477, F1: 461, F2: (!((fi42 * 248) == (fi17 / (((fi17 / ((fi42) | 1))) | 1))))}
                                                v45 := GenS7{F0: 951}
                                                var v46 uint = 916
                                                _ = v43
                                                _ = v44
                                                _ = v45
                                                _ = v46
                                            } else {
                                                var v47 float32 = 0.0
                                                _ = v47
                                            }
                                            v2 -= 269
                                        }
                                        _ = fi42
                                    }
                                default:
                                    var v48 uint = 4294967295
                                    _ = v48
                                }
                            } else {
                                var v49 int8 = 77
                                for fi50 := 0; fi50 < 3; fi50++ {
                                    var v51 int32 = -1
                                    v51 = v51
                                    _ = fi50
                                    _ = v51
                                }
                                _ = v49
                            }
                        }
                        var v52 uint64 = v2
                        var v53 bool = false
                        _ = v52
                        _ = v53
                    }
                    _ = fi17
                }
            default:
                fmt.Println(v2)
            }
        }
    case -2147483648:
        fmt.Println(v2)
    default:
        v54 := GenS4{F0: "hello", F1: -9223372036854775808}
        _ = v54
    }
    if true {
        if true {
            switch ("foo bar" + "a") {
            case "":
                v2 += v2
            case "foo bar":
                v2 -= v2
            case "hello":
                v2 += v2
            }
            if ((431 < 127) && false) {
                v2 += v2
                var v55 int16 = 32767
                var v56 int = -1
                _ = v55
                _ = v56
            }
            v2 = 18446744073709551615
        } else {
            for fi57 := 0; fi57 < 1; fi57++ {
                switch "foo bar" {
                case "":
                    v2 = (v2 % ((v2) | 1))
                case "foo bar":
                    for fi58 := 0; fi58 < 14; fi58++ {
                        v2 -= (v2 - (v2 << ((42) & 63)))
                        v59 := GenS2{F0: "foo bar", F1: (false && false)}
                        _ = fi58
                        _ = v59
                    }
                }
                _ = fi57
            }
        }
    } else {
        switch ("a" + "foo bar") {
        case "a":
            var v60 int64 = 0
            _ = v60
        }
        v61 := GenS8{F0: v2}
        _ = v61
    }
    delete(v1, ("" + "a"))
    var v62 uint16 = 0
    v63 := GenS9{F0: ("hello" + ("a" + "foo bar"))}
    fmt.Println(v2)
    var v64 bool = false
    if len(v3) > 0 {
        v3[0] = 32768
    }
    v62 = 343
    for fi65 := 0; fi65 < 4; fi65++ {
        v66 := GenS8{F0: (^(v2))}
        _ = len(v1)
        var v67 string = "foo bar"
        _ = fi65
        _ = v66
        _ = v67
    }
    for fi68 := 0; fi68 < 15; fi68++ {
        for fi69 := 0; fi69 < 14; fi69++ {
            var v70 int64 = 0
            if false {
                if v64 {
                    fmt.Println(fi68, v64, v70)
                    v71 := GenS8{F0: uint64(v62)}
                    _ = v71
                } else {
                    var v72 int32 = 1
                    _ = v72
                }
            } else {
                if (-1 > 42) {
                    v70 += v70
                }
                fmt.Println(v70)
                v73 := GenS9{F0: ("" + "")}
                _ = v73
            }
            if (-40 <= 100) {
                var v74 float32 = (((3.14159 - 1e10) + -2.25) / 1.5)
                var v75 uint32 = genFn0(32767)
                var v76 bool = v64
                if (-9223372036854775808 >= v70) {
                    v3 = append(v3, (fi69 >> ((1) & 63)))
                    v77 := v1["foo bar"]
                    switch (fi68 ^ fi68) {
                    case -9223372036854775808:
                        if v64 {
                            if false {
                                if v64 {
                                    var v78 int32 = 2147483647
                                    _ = v78
                                } else {
                                    var v79 uint8 = 94
                                    var v80 float64 = 0.1
                                    if v76 {
                                        var v81 uint32 = v75
                                        v81 -= v81
                                        var v82 uint16 = v62
                                        _ = v81
                                        _ = v82
                                    } else {
                                        switch "\n\t" {
                                        case "a":
                                            switch fi69 {
                                            case 523:
                                                var v83 bool = (v70 < v70)
                                                _ = v83
                                            case 9223372036854775807:
                                                if (v75 <= (^(v75))) {
                                                    v74 = ((0.0 / v74) + v74)
                                                }
                                            case 951:
                                                if v76 {
                                                    var v84 uint32 = uint32(v79)
                                                    v76 = true
                                                    v85 := GenS1{F0: 0, F1: (v62 >> ((4294967295) & 63)), F2: (v70 << ((206) & 63))}
                                                    _ = v84
                                                    _ = v85
                                                }
                                            }
                                        case "hello":
                                            v62 = (v62 & (v62 | 0))
                                        case "foo bar":
                                            var v86 int = (fi69 - -1)
                                            _ = v86
                                        }
                                    }
                                    _ = v79
                                    _ = v80
                                }
                                if ((-2.25 + (float64((v74 - v74)) + 3.14159)) > (0.1 * 0.1)) {
                                    switch (fi68 << ((220) & 63)) {
                                    case 4294967295:
                                        delete(v1, ("hello" + "\n\t"))
                                    default:
                                        var v87 bool = (125 > 100)
                                        if false {
                                            var v88 int = fi68
                                            _ = v88
                                        } else {
                                            switch 86 {
                                            case 616:
                                                fmt.Println(v62, v70)
                                            }
                                            if v87 {
                                                v62 = (^(v62))
                                                if len(v3) > 0 {
                                                    v3[0] = fi68
                                                }
                                                var v89 uint16 = v62
                                                var v90 int64 = 201
                                                _ = v89
                                                _ = v90
                                            } else {
                                                if v77 {
                                                    v76 = ((!((42 != -32768) || false)) || v77)
                                                    var v91 uint64 = v2
                                                    v92 := GenS10{F0: (v75 >> ((83) & 63)), F1: 442, F2: (v75 >> ((uint(v75)) & 63))}
                                                    _ = v91
                                                    _ = v92
                                                }
                                                switch "a" {
                                                case "foo bar":
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        _ = len(v1)
                                                    default:
                                                        v62 = v62
                                                    }
                                                default:
                                                    fmt.Println(v87, fi69, v87)
                                                    v74 += ((v74 / ((v74 / v74) - 1.5)) * v74)
                                                    var v93 int = (^(fi69))
                                                    var v94 uint16 = v62
                                                    var v95 float64 = 3.14159
                                                    _ = v93
                                                    _ = v94
                                                    _ = v95
                                                }
                                                if v87 {
                                                    if v87 {
                                                        fmt.Println(v70, v76, v76)
                                                    }
                                                }
                                                if (!false) {
                                                    v96 := GenS11{F0: 31, F1: 0}
                                                    _ = len(v3)
                                                    _ = v96
                                                } else {
                                                    fmt.Println(fi68, v77, fi68)
                                                    v76 = v87
                                                    if len(v3) > 0 {
                                                        v3[0] = (^(fi68))
                                                    }
                                                }
                                                v97 := GenS8{F0: (v2 >> ((2) & 63))}
                                                switch fi69 {
                                                case 162:
                                                    if v87 {
                                                        switch "hello" {
                                                        case "a":
                                                            v98 := v1[("a" + ("\n\t" + ("\n\t" + (("a" + "a") + ("\n\t" + "hello")))))]
                                                            _ = v98
                                                        case "":
                                                            var v99 float64 = (1.5 * (0.1 * 0.0))
                                                            _ = v99
                                                        default:
                                                            fmt.Println(v74)
                                                        }
                                                    }
                                                default:
                                                    var v100 string = ("\n\t" + (("foo bar" + "foo bar") + "a"))
                                                    _ = v100
                                                }
                                                var v101 string = "hello"
                                                _ = v101
                                                _ = v97
                                            }
                                            v77 = (((1 == 736) || v77) || (!(!false)))
                                            if ((!("hello" != ("hello" + "foo bar"))) || v87) {
                                                v102 := GenS9{F0: "\n\t"}
                                                v62 = (v62 / ((429) | 1))
                                                var v103 uint32 = (v75 * v75)
                                                _ = v102
                                                _ = v103
                                            }
                                            fmt.Println(v2)
                                            switch "a" {
                                            case "\n\t":
                                                if v76 {
                                                    if ((^(fi68)) != fi68) {
                                                        v104 := v1[("\n\t" + ("\n\t" + "\n\t"))]
                                                        _ = v104
                                                    } else {
                                                        switch (fi69 >> ((4294967295) & 63)) {
                                                        case 42:
                                                            v75 = 2
                                                        default:
                                                            switch (fi69 ^ fi68) {
                                                            case 918:
                                                                if (false || ((^(v70)) == (^(v70)))) {
                                                                    if (v70 != v70) {
                                                                        if (float32(-2.25) == float32(float64(((1.5 * v74) + v74)))) {
                                                                            v105 := GenS12{F0: 28}
                                                                            _ = v105
                                                                        } else {
                                                                            switch "a" {
                                                                            case "a":
                                                                                var v106 int = fi68
                                                                                _ = v106
                                                                            case "foo bar":
                                                                                v107 := v1[""]
                                                                                _ = v107
                                                                            default:
                                                                                v2 = (^(v2))
                                                                            }
                                                                        }
                                                                        var v108 string = "foo bar"
                                                                        v109 := v1[v108]
                                                                        _ = v108
                                                                        _ = v109
                                                                    }
                                                                    fmt.Println(v75, v77)
                                                                    if true {
                                                                        if ("a" >= "a") {
                                                                            var v110 uint = 1
                                                                            _ = v110
                                                                        } else {
                                                                            fmt.Println(v77, v87, v87)
                                                                        }
                                                                        v70 += -919
                                                                        var v111 int8 = 42
                                                                        _ = v111
                                                                    } else {
                                                                        var v112 int = (fi68 | 34)
                                                                        v113 := v1[("" + "foo bar")]
                                                                        _ = v112
                                                                        _ = v113
                                                                    }
                                                                } else {
                                                                    delete(v1, (("" + (((("foo bar" + "foo bar") + ("foo bar" + "")) + "") + ("\n\t" + ("\n\t" + "\n\t")))) + "\n\t"))
                                                                    if (42 <= 235) {
                                                                        var v114 string = "hello"
                                                                        var v115 int32 = 2147483647
                                                                        v70 = 335
                                                                        _ = v114
                                                                        _ = v115
                                                                    } else {
                                                                        var v116 float32 = float32((3.14159 - -2.25))
                                                                        var v117 float64 = (1.5 - float64(3.14159))
                                                                        var v118 float64 = ((v117 - float64((v116 + -2.25))) + v117)
                                                                        delete(v1, ("foo bar" + (((("\n\t" + ("\n\t" + "foo bar")) + "foo bar") + "a") + ("hello" + "hello"))))
                                                                        var v119 uint64 = (v2 - v2)
                                                                        _ = v116
                                                                        _ = v117
                                                                        _ = v118
                                                                        _ = v119
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        if ((^(v62)) >= v62) {
                                                            fmt.Println(v74, v76, v75)
                                                            if (2 != (v75 ^ v75)) {
                                                                if v77 {
                                                                    switch (fi68 ^ (fi68 + (^(fi69)))) {
                                                                    case -9223372036854775808:
                                                                        var v120 uint16 = 151
                                                                        _ = v120
                                                                    case -2147483648:
                                                                        switch "foo bar" {
                                                                        case "\n\t":
                                                                            var v121 bool = (true || ((fi68 >> ((2) & 63)) == fi69))
                                                                            switch fi68 {
                                                                            case 314:
                                                                                var v122 int16 = 1000
                                                                                _ = v122
                                                                            case 0:
                                                                                v121 = (false || true)
                                                                            }
                                                                            v70 += v70
                                                                            _ = v121
                                                                        case "a":
                                                                            v2 += 4294967295
                                                                        case "hello":
                                                                            var v123 uint64 = (^(v2))
                                                                            _ = v123
                                                                        }
                                                                    case 477:
                                                                        fmt.Println(v87)
                                                                        var v124 bool = (fi69 <= fi68)
                                                                        var v125 int64 = (^(v70))
                                                                        var v126 bool = (v74 > v74)
                                                                        var v127 bool = (!(true || v124))
                                                                        _ = v124
                                                                        _ = v125
                                                                        _ = v126
                                                                        _ = v127
                                                                    }
                                                                }
                                                                v64 = (((v2 & 301) > v2) || (!v76))
                                                                var v128 uint = 65535
                                                                _ = v128
                                                            }
                                                            v76 = ((fi68 & (fi69 & (fi68 - (-(fi68))))) >= fi68)
                                                        } else {
                                                            fmt.Println(v87, fi68)
                                                            switch (fi68 & (fi69 | 0)) {
                                                            case 9223372036854775807:
                                                                v62 = v62
                                                            case -974:
                                                                v3 = append(v3, -9223372036854775808)
                                                            default:
                                                                v129 := GenS5{F0: v77, F1: fi69}
                                                                _ = v129
                                                            }
                                                        }
                                                        var v130 uint8 = uint8(v75)
                                                        _ = v130
                                                    }
                                                } else {
                                                    v75 += v75
                                                    var v131 uint = uint(v75)
                                                    if v87 {
                                                        if ((v131 * (v131 & (^(v131)))) >= 42) {
                                                            switch (("a" + "foo bar") + "foo bar") {
                                                            case "\n\t":
                                                                var v132 bool = true
                                                                _ = v132
                                                            }
                                                        }
                                                    } else {
                                                        v133 := GenS13{F0: ("a" + "foo bar"), F1: -1}
                                                        if (false && true) {
                                                            v76 = v64
                                                            v131 += v131
                                                            v134 := GenS14{F0: 0, F1: -128, F2: (^(v75))}
                                                            _ = v134
                                                        } else {
                                                            var v135 int64 = (v70 + (v70 / ((102) | 1)))
                                                            var v136 int8 = 114
                                                            var v137 int32 = -1
                                                            _ = v135
                                                            _ = v136
                                                            _ = v137
                                                        }
                                                        var v138 int32 = 42
                                                        _ = v133
                                                        _ = v138
                                                    }
                                                    fmt.Println(fi68, fi69, v87)
                                                    _ = v131
                                                }
                                            case "hello":
                                                v74 = 1.5
                                            }
                                        }
                                        _ = v87
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    switch ("hello" + ("" + ("" + "hello"))) {
                                    case "foo bar":
                                        var v139 int = (^(fi69))
                                        _ = v139
                                    case "\n\t":
                                        fmt.Println(fi68)
                                    case "":
                                        if v76 {
                                            if v77 {
                                                v140 := GenS1{F0: 42, F1: (^(v62)), F2: v70}
                                                v70 = v70
                                                v141 := GenS15{F0: 42, F1: 42}
                                                _ = v140
                                                _ = v141
                                            } else {
                                                fmt.Println(v74, v2, v64)
                                                var v142 uint32 = 1
                                                fmt.Println(fi69, v70)
                                                delete(v1, "foo bar")
                                                _ = v142
                                            }
                                            var v143 uint32 = 0
                                            var v144 int64 = (v70 - -9223372036854775808)
                                            _ = v143
                                            _ = v144
                                        }
                                    }
                                    v70 = (-(v70))
                                }
                            } else {
                                if v64 {
                                    v74 = (1.5 / (float32((((1e10 * ((1e10 * -2.25) / (1e10 * 1e10))) * 1e10) * 1e10)) - (((((((-2.25 / float32(-2.25)) / v74) * (-2.25 * (float32(-2.25) * 1e10))) + (v74 * float32(((float64(v74) - 3.14159) + ((3.14159 - (1.5 / (((1.5 + (0.1 / (0.1 / 0.1))) - 0.1) / ((3.14159 - (0.1 - 0.1)) - ((1.5 - 0.0) - 1e10))))) * float64((1e10 / v74))))))) - v74) / -2.25) / v74)))
                                    var v145 uint16 = v62
                                    var v146 uint = 4294967295
                                    _ = v145
                                    _ = v146
                                }
                            }
                        } else {
                            v75 += genFn1(739)
                            switch "\n\t" {
                            case "a":
                                fmt.Println(v77)
                            case "foo bar":
                                if true {
                                    if (v77 || v77) {
                                        if false {
                                            switch "a" {
                                            case "foo bar":
                                                var v147 float32 = v74
                                                _ = v147
                                            case "hello":
                                                if ((1000 <= -574) && (122 > 201)) {
                                                    fmt.Println(v70)
                                                    var v148 float32 = v74
                                                    _ = v148
                                                } else {
                                                    switch fi69 {
                                                    case 127:
                                                        var v149 int16 = 42
                                                        _ = v149
                                                    default:
                                                        if v76 {
                                                            fmt.Println(v75, fi68)
                                                        } else {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            var v150 uint32 = (v75 >> ((1) & 63))
                                                            v77 = v77
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            _ = v150
                                                        }
                                                    }
                                                    fmt.Println(v2, fi69)
                                                }
                                            }
                                        } else {
                                            v151 := GenS4{F0: ("" + "hello"), F1: (^(v70))}
                                            _ = v151
                                        }
                                    } else {
                                        var v152 uint64 = 2
                                        v153 := GenS14{F0: 128, F1: -128, F2: uint32(v62)}
                                        if (v152 > v2) {
                                            switch "\n\t" {
                                            case "a":
                                                if v76 {
                                                    v154 := GenS16{F0: fi69, F1: 42, F2: v152}
                                                    var v155 bool = (42 >= 42)
                                                    fmt.Println(fi68, v76)
                                                    _ = v154
                                                    _ = v155
                                                } else {
                                                    fmt.Println(v62, v152, fi68)
                                                    if (v76 && (false && false)) {
                                                        if v76 {
                                                            if true {
                                                                var v156 uint32 = v75
                                                                _ = v156
                                                            } else {
                                                                if v76 {
                                                                    if (v77 || v64) {
                                                                        v1["a"] = (int64(v62) > v70)
                                                                    } else {
                                                                        if ((v64 || v64) && ((3.14159 + (0.1 / float64((v74 * v74)))) <= ((((0.0 / (3.14159 + -2.25)) - 3.14159) * (float64(v74) * 0.0)) - -2.25))) {
                                                                            v77 = v77
                                                                            v2 = v152
                                                                        } else {
                                                                            _ = len(v3)
                                                                            v75 -= genFn1(255)
                                                                        }
                                                                    }
                                                                }
                                                                v74 = 1.5
                                                            }
                                                            fmt.Println(v62, v76, fi69)
                                                        } else {
                                                            v157 := GenS17{F0: 121, F1: (fi68 % (((fi68 | (fi68 - -2147483648))) | 1)), F2: 1}
                                                            switch "a" {
                                                            case "":
                                                                v152 = v2
                                                            case "hello":
                                                                v64 = (!(!v77))
                                                            case "a":
                                                                v2 -= 18446744073709551615
                                                            }
                                                            delete(v1, "a")
                                                            var v158 float32 = v74
                                                            _ = v157
                                                            _ = v158
                                                        }
                                                        switch "" {
                                                        case "hello":
                                                            if v64 {
                                                                var v159 float64 = (0.0 / 0.1)
                                                                var v160 int8 = 79
                                                                _ = v159
                                                                _ = v160
                                                            }
                                                        case "\n\t":
                                                            var v161 int64 = (v70 & v70)
                                                            _ = v161
                                                        }
                                                        v162 := GenS14{F0: 128, F1: 90, F2: 65535}
                                                        _ = v162
                                                    }
                                                    if v76 {
                                                        var v163 uint = 65535
                                                        switch -678 {
                                                        case 32768:
                                                            var v164 uint8 = 42
                                                            _ = v164
                                                        default:
                                                            switch fi69 {
                                                            case -252:
                                                                fmt.Println(v152)
                                                            case 127:
                                                                v64 = v64
                                                            case -1:
                                                                v3 = append(v3, fi69)
                                                            }
                                                        }
                                                        v62 -= (v62 >> ((1) & 63))
                                                        if ("" != ("a" + "")) {
                                                            v2 = 4294967295
                                                            v152 = v152
                                                            var v165 int = (-(fi69))
                                                            _ = v165
                                                        }
                                                        v1[(("hello" + "") + ("hello" + ("hello" + "hello")))] = false
                                                        _ = v163
                                                    } else {
                                                        if true {
                                                            delete(v1, "foo bar")
                                                            v77 = ((v70 & (v70 << ((18446744073709551615) & 63))) <= v70)
                                                            v74 = 1.5
                                                        }
                                                        _ = len(v3)
                                                    }
                                                }
                                            case "hello":
                                                if true {
                                                    if (1 <= 71) {
                                                        switch "a" {
                                                        case "\n\t":
                                                            fmt.Println(v75)
                                                        case "foo bar":
                                                            if true {
                                                                v166 := GenS16{F0: (fi68 | (-(fi69))), F1: v62, F2: v152}
                                                                _ = len(v3)
                                                                v76 = (65535 > 641)
                                                                _ = v166
                                                            } else {
                                                                v167 := GenS18{F0: "foo bar", F1: v76, F2: 2}
                                                                _ = v167
                                                            }
                                                        case "":
                                                            v168 := GenS12{F0: 164}
                                                            _ = v168
                                                        }
                                                        fmt.Println(fi69, v62)
                                                        v76 = v64
                                                    }
                                                    _ = len(v3)
                                                    var v169 uint16 = (^(v62))
                                                    v170 := GenS16{F0: fi69, F1: v169, F2: uint64(v75)}
                                                    _ = v169
                                                    _ = v170
                                                }
                                            default:
                                                v76 = v76
                                            }
                                            switch 1 {
                                            case -2147483648:
                                                _ = len(v3)
                                            case 4294967295:
                                                if ("\n\t" != "foo bar") {
                                                    if (1 > 255) {
                                                        if false {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        } else {
                                                            v70 = (-(v70))
                                                            fmt.Println(v76, v74)
                                                            if false {
                                                                var v171 int16 = -32768
                                                                var v172 int32 = -2147483648
                                                                var v173 uint64 = (v2 >> ((42) & 63))
                                                                if v64 {
                                                                    var v174 string = "\n\t"
                                                                    v175 := GenS3{F0: 127, F1: (fi68 - (^(fi69)))}
                                                                    v62 -= uint16(v173)
                                                                    v176 := v1["hello"]
                                                                    _ = v174
                                                                    _ = v175
                                                                    _ = v176
                                                                }
                                                                var v177 int8 = 42
                                                                _ = v171
                                                                _ = v172
                                                                _ = v173
                                                                _ = v177
                                                            } else {
                                                                var v178 int8 = 42
                                                                _ = v178
                                                            }
                                                        }
                                                        v179 := GenS11{F0: -888, F1: 1}
                                                        _ = v179
                                                    } else {
                                                        var v180 uint64 = (v152 % (((v152 / (((v152 - v2)) | 1))) | 1))
                                                        var v181 uint = 2
                                                        if v64 {
                                                            var v182 uint16 = (v62 << ((uint(v75)) & 63))
                                                            var v183 int8 = 0
                                                            switch "hello" {
                                                            case "":
                                                                v2 = 273
                                                            case "\n\t":
                                                                var v184 uint16 = v182
                                                                _ = v184
                                                            case "foo bar":
                                                                if v76 {
                                                                    if v76 {
                                                                        if v77 {
                                                                            v76 = v64
                                                                            v75 = v75
                                                                            var v185 uint16 = v62
                                                                            _ = v185
                                                                        } else {
                                                                            v183 = (v183 / ((127) | 1))
                                                                            var v186 float64 = 0.1
                                                                            _ = v186
                                                                        }
                                                                        var v187 uint16 = v182
                                                                        _ = v187
                                                                    } else {
                                                                        v75 = (^(v75))
                                                                    }
                                                                    switch 839 {
                                                                    case 1:
                                                                        v182 -= v182
                                                                    case 0:
                                                                        if v76 {
                                                                            var v188 uint16 = v182
                                                                            if (v64 && (-2.25 != 3.14159)) {
                                                                                var v189 string = "foo bar"
                                                                                var v190 int8 = (v183 << (((v181 << ((18446744073709551615) & 63))) & 63))
                                                                                _ = v189
                                                                                _ = v190
                                                                            }
                                                                            _ = len(v1)
                                                                            _ = v188
                                                                        }
                                                                    default:
                                                                        fmt.Println(v75)
                                                                    }
                                                                } else {
                                                                    v1["foo bar"] = v76
                                                                    v75 = v75
                                                                    var v191 int8 = v183
                                                                    _ = v191
                                                                }
                                                            default:
                                                                v74 -= v74
                                                                v181 = (v181 << ((v181) & 63))
                                                                v181 += 42
                                                            }
                                                            v1[((("foo bar" + "hello") + ("" + ((("" + "\n\t") + ("\n\t" + (("\n\t" + "\n\t") + "\n\t"))) + "\n\t"))) + "\n\t")] = false
                                                            v2 = v152
                                                            _ = v182
                                                            _ = v183
                                                        }
                                                        _ = v180
                                                        _ = v181
                                                    }
                                                    v192 := GenS17{F0: 2, F1: fi68, F2: 68}
                                                    delete(v1, "")
                                                    delete(v1, "")
                                                    _ = v192
                                                } else {
                                                    var v193 uint8 = 128
                                                    _ = v193
                                                }
                                            case -506:
                                                delete(v1, ("" + "\n\t"))
                                            }
                                            fmt.Println(v70, v2)
                                        }
                                        _ = v152
                                        _ = v153
                                    }
                                    var v194 int = int(v70)
                                    var v195 float32 = v74
                                    if (4294967295 != 2) {
                                        if v76 {
                                            if true {
                                                fmt.Println(v2, v75, v194)
                                                if len(v3) > 0 {
                                                    v3[0] = (fi68 | fi69)
                                                }
                                            }
                                            fmt.Println(v194)
                                            switch (fi68 | fi69) {
                                            case -9223372036854775808:
                                                var v196 uint8 = uint8(v2)
                                                _ = v196
                                            default:
                                                fmt.Println(fi69, v64, v76)
                                            }
                                        }
                                    } else {
                                        var v197 int64 = 0
                                        _ = v197
                                    }
                                    _ = v194
                                    _ = v195
                                } else {
                                    v198 := v1[(("" + "hello") + "a")]
                                    v64 = (!v76)
                                    _ = v198
                                }
                            default:
                                v199 := GenS19{F0: v62}
                                _ = v199
                            }
                            v62 -= (^(v62))
                            var v200 int = fi69
                            _ = v200
                        }
                    case 2147483647:
                        if v77 {
                            v201 := GenS20{F0: v70}
                            v2 = v2
                            switch "hello" {
                            case "hello":
                                if v77 {
                                    v70 -= 0
                                } else {
                                    v62 -= (^(v62))
                                    if (!false) {
                                        if v64 {
                                            delete(v1, ("hello" + "a"))
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v74 += 1e10
                                    }
                                }
                            default:
                                v202 := v1["foo bar"]
                                _ = v202
                            }
                            _ = v201
                        }
                    default:
                        var v203 int16 = 0
                        _ = v203
                    }
                    _ = v77
                } else {
                    if (v76 || (!false)) {
                        if len(v3) > 0 {
                            v3[0] = (fi69 ^ (fi69 - (fi68 >> ((2) & 63))))
                        }
                        var v204 uint64 = 543
                        _ = v204
                    } else {
                        var v205 float32 = 3.14159
                        _ = len(v3)
                        var v206 int8 = -1
                        _ = v205
                        _ = v206
                    }
                    var v207 uint8 = uint8(v62)
                    if (fi69 == (fi69 ^ int(v75))) {
                        if (v64 && ("\n\t" > "a")) {
                            v208 := GenS17{F0: v207, F1: (-(fi69)), F2: -20}
                            _ = v208
                        } else {
                            v75 = 599
                        }
                        v209 := GenS4{F0: "", F1: v70}
                        _ = v209
                    } else {
                        var v210 string = "hello"
                        _ = v210
                    }
                    v62 -= 510
                    v70 = (v70 + -9223372036854775808)
                    _ = v207
                }
                switch 127 {
                case -9223372036854775808:
                    var v211 int = 32767
                    _ = v211
                }
                switch fi69 {
                case 163:
                    if v76 {
                        v212 := GenS3{F0: 127, F1: fi69}
                        v62 -= 42
                        v2 -= v2
                        _ = v212
                    } else {
                        var v213 uint32 = v75
                        var v214 uint32 = (v213 << ((255) & 63))
                        _ = v213
                        _ = v214
                    }
                case 128:
                    var v215 bool = (((-(v70)) <= (v70 % ((v70) | 1))) && (v76 && (-51 == 42)))
                    _ = v215
                case 127:
                    var v216 uint32 = (^(v75))
                    _ = v216
                default:
                    switch ("" + "foo bar") {
                    case "":
                        var v217 int = 0
                        _ = v217
                    case "\n\t":
                        var v218 int8 = -1
                        _ = v218
                    }
                }
                _ = v74
                _ = v75
                _ = v76
            } else {
                _ = len(v3)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v219 := GenS8{F0: v2}
                _ = v219
            }
            _ = fi69
            _ = v70
        }
        for fi220 := 0; fi220 < 9; fi220++ {
            var v221 string = ("foo bar" + ("a" + "\n\t"))
            _ = len(v1)
            v62 -= 1
            _ = fi220
            _ = v221
        }
        if v64 {
            fmt.Println(v64, v62)
            v2 -= v2
        } else {
            var v222 float64 = (-2.25 + (1e10 + ((0.1 * 1.5) / 1.5)))
            if v64 {
                if false {
                    for fi223 := 0; fi223 < 7; fi223++ {
                        if v64 {
                            var v224 int = (fi68 & fi223)
                            _ = v224
                        }
                        if v64 {
                            switch fi223 {
                            case 32768:
                                v225 := GenS21{F0: (v2 % (((v2 | 2)) | 1))}
                                _ = v225
                            case 32767:
                                v2 += v2
                            case 9223372036854775807:
                                v226 := GenS5{F0: ((v2 | (v2 << ((890) & 63))) == (^(v2))), F1: 9223372036854775807}
                                _ = v226
                            default:
                                var v227 float32 = 0.0
                                _ = v227
                            }
                        }
                        _ = fi223
                    }
                }
            } else {
                v228 := GenS15{F0: -128, F1: 127}
                fmt.Println(v64)
                v229 := GenS8{F0: v2}
                _ = v228
                _ = v229
            }
            for fi230 := 0; fi230 < 3; fi230++ {
                v231 := GenS22{F0: 522, F1: v2, F2: "a"}
                var v232 int = 791
                switch 2147483647 {
                case 127:
                    fmt.Println(v232, v232, v222)
                case -68:
                    v1[""] = (-9223372036854775808 <= -798)
                    v64 = (0 != 127)
                case -574:
                    v232 += -1
                }
                _ = fi230
                _ = v231
                _ = v232
            }
            _ = v222
        }
        _ = fi68
    }
    _ = len(v1)
    var v233 int16 = 0
    v234 := GenS16{F0: -609, F1: 42, F2: (v2 * (v2 ^ (v2 >> ((1) & 63))))}
    delete(v1, "a")
    for fi235 := 0; fi235 < 1; fi235++ {
        var v236 float64 = 1e10
        _ = fi235
        _ = v236
    }
    for fi237 := 0; fi237 < 5; fi237++ {
        _ = len(v3)
        _ = fi237
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v233:", v233)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v62:", v62)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v233
    _ = v234
    _ = v3
    _ = v62
    _ = v63
    _ = v64
}
