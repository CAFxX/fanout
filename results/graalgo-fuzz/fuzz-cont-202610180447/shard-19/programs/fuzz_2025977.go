package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 bool }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint64; F1 uint; F2 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint16; F1 int8; F2 string }
type GenS11 struct { F0 uint16; F1 uint64 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int64; F1 uint }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint8; F1 bool }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 int64 }
type GenS19 struct { F0 int64 }
type GenS20 struct { F0 int32; F1 int; F2 int16 }
type GenS21 struct { F0 int; F1 uint; F2 int64 }
func genFn0(p0 int16) string {
    p0 = p0
    if false {
        v4 := GenS1{F0: ("foo bar" + ("foo bar" + "")), F1: 1, F2: (!true)}
        if (42 != -2147483648) {
            p0 = p0
            v5 := GenS2{F0: -9223372036854775808, F1: 365}
            v6 := []int{}
            _ = v5
            _ = v6
        }
        _ = v4
    } else {
        var v7 bool = (49 <= 0)
        for fi8 := 0; fi8 < 2; fi8++ {
            for fi9 := 0; fi9 < 1; fi9++ {
                v10 := []int{}
                v7 = (127 < 2)
                _ = fi9
                _ = v10
            }
            v7 = (false || (false && v7))
            v7 = (!(((float32(0.0) / -2.25) < 1e10) && (v7 && v7)))
            _ = fi8
        }
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 int8 = 127
            switch (fi11 & 516) {
            case 1:
                var v13 int = -9223372036854775808
                _ = v13
            case 729:
                switch (fi11 | (fi11 << ((42) & 63))) {
                case 9223372036854775807:
                    v12 += v12
                default:
                    v14 := []int{(fi11 % ((fi11) | 1)), (fi11 - 873)}
                    _ = v14
                }
            case 42:
                for fi15 := 0; fi15 < 9; fi15++ {
                    v16 := GenS1{F0: "\n\t", F1: 986, F2: (!(p0 <= (p0 * -1)))}
                    _ = fi15
                    _ = v16
                }
            default:
                v7 = false
            }
            _ = fi11
            _ = v12
        }
        _ = v7
    }
    switch 0 {
    case -861:
        for fi17 := 0; fi17 < 1; fi17++ {
            if false {
                for fi18 := 0; fi18 < 11; fi18++ {
                    v19 := []int{fi17, (fi17 ^ fi17)}
                    switch (fi17 / (((fi17 | (fi18 % ((fi17) | 1)))) | 1)) {
                    case 839:
                        _ = len(v19)
                    default:
                        p0 -= p0
                    }
                    _ = fi18
                    _ = v19
                }
                v20 := GenS3{F0: (p0 | (p0 | 32767)), F1: 918, F2: 0}
                _ = v20
            } else {
                p0 += (^(p0))
                for fi21 := 0; fi21 < 14; fi21++ {
                    var v22 bool = true
                    p0 = p0
                    v22 = true
                    v23 := []int{(fi17 % (((fi21 & (-(fi17)))) | 1)), (fi17 - fi17)}
                    var v24 uint64 = 18446744073709551615
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                }
            }
            for fi25 := 0; fi25 < 12; fi25++ {
                var v26 uint8 = 1
                _ = fi25
                _ = v26
            }
            _ = fi17
        }
    default:
        if (false || false) {
            var v27 float64 = ((0.1 + float64(1.5)) + 1e10)
            v28 := GenS2{F0: 932, F1: -1}
            _ = v27
            _ = v28
        } else {
            p0 += (-(p0))
            for fi29 := 0; fi29 < 9; fi29++ {
                for fi30 := 0; fi30 < 12; fi30++ {
                    v31 := []int{}
                    switch "" {
                    case "a":
                        p0 += p0
                        p0 += (p0 + p0)
                        if len(v31) >= 2 {
                            v31 = v31[1:len(v31)]
                        }
                        if true {
                            p0 = p0
                            var v32 uint8 = 106
                            v33 := GenS3{F0: (p0 & 210), F1: 0, F2: 42}
                            _ = v32
                            _ = v33
                        }
                        fmt.Println(fi30)
                    case "foo bar":
                        v34 := GenS2{F0: (fi30 + 2147483647), F1: (fi29 & fi29)}
                        _ = v34
                    case "":
                        p0 -= (^(p0))
                        p0 = -373
                        v31 = append(v31, fi30)
                        fmt.Println(fi29)
                        v35 := GenS2{F0: fi30, F1: (fi29 << ((1) & 63))}
                        _ = v35
                    default:
                        var v36 uint = 128
                        _ = v36
                    }
                    if (((-2.25 + 1.5) / ((0.0 - 1e10) * (-2.25 - ((-2.25 + (0.0 - -2.25)) - 0.1)))) < -2.25) {
                        switch "foo bar" {
                        case "\n\t":
                            if (true && true) {
                                p0 -= (p0 << ((1) & 63))
                                switch 1 {
                                case 127:
                                    p0 += p0
                                default:
                                    p0 = 622
                                }
                            } else {
                                p0 -= 32767
                                v37 := GenS4{F0: 2, F1: 468, F2: (false && true)}
                                if true {
                                    p0 = p0
                                } else {
                                    v31 = append(v31, (fi30 / ((-378) | 1)))
                                    if (!false) {
                                        var v38 uint8 = 255
                                        v39 := GenS2{F0: fi30, F1: fi30}
                                        if (!false) {
                                            if ((((!false) && false) && true) || false) {
                                                fmt.Println(p0, p0, v38)
                                            } else {
                                                var v40 int32 = 1
                                                v38 = 127
                                                _ = v40
                                            }
                                            var v41 int32 = 0
                                            _ = v41
                                        } else {
                                            if ((-(p0)) >= p0) {
                                                v42 := GenS5{F0: 0}
                                                if len(v31) > 0 {
                                                    v31[0] = fi29
                                                }
                                                var v43 uint8 = v38
                                                _ = v42
                                                _ = v43
                                            } else {
                                                switch ("" + ("" + "hello")) {
                                                case "foo bar":
                                                    v38 += 255
                                                case "a":
                                                    var v44 int = fi30
                                                    _ = v44
                                                default:
                                                    if len(v31) >= 2 {
                                                        v31 = v31[1:len(v31)]
                                                    }
                                                }
                                                switch fi30 {
                                                case -533:
                                                    fmt.Println(p0, v38, fi29)
                                                default:
                                                    var v45 uint8 = (v38 / ((98) | 1))
                                                    _ = v45
                                                }
                                                var v46 bool = (1e10 > 0.0)
                                                _ = v46
                                            }
                                            fmt.Println(fi30, fi30, fi29)
                                        }
                                        _ = v38
                                        _ = v39
                                    }
                                    if len(v31) >= 2 {
                                        v31 = v31[1:len(v31)]
                                    }
                                }
                                _ = v37
                            }
                            p0 -= p0
                            p0 -= (p0 & p0)
                        default:
                            fmt.Println(fi29, p0)
                        }
                        v47 := GenS3{F0: (p0 >> ((889) & 63)), F1: 297, F2: 200}
                        if len(v31) > 0 {
                            v31[0] = (fi30 | fi29)
                        }
                        _ = v47
                    }
                    var v48 int64 = 9223372036854775807
                    _ = fi30
                    _ = v31
                    _ = v48
                }
                p0 += 42
                _ = fi29
            }
        }
    }
    var v49 int16 = (^(p0))
    _ = p0
    _ = v49
    return ""
}
func genFn1() uint8 {
    v50 := GenS6{F0: (-1 < 177)}
    var v51 int16 = 1
    v51 = -757
    for fi52 := 0; fi52 < 10; fi52++ {
        if true {
            if (-9223372036854775808 != 0) {
                v51 = v51
                v53 := GenS5{F0: -847}
                _ = v53
            } else {
                for fi54 := 0; fi54 < 9; fi54++ {
                    switch fi54 {
                    case 9223372036854775807:
                        var v55 uint = 0
                        _ = v55
                    case 127:
                        var v56 uint32 = 42
                        _ = v56
                    case 894:
                        v57 := []int{fi54, (-(fi54)), (fi52 ^ (^(fi54))), (fi54 ^ (fi54 << ((65535) & 63)))}
                        _ = v57
                    default:
                        fmt.Println(v51)
                    }
                    fmt.Println(fi52)
                    _ = fi54
                }
                v51 += (v51 + (v51 << ((343) & 63)))
            }
            if (false || false) {
                v58 := GenS7{F0: true}
                v59 := GenS8{F0: 0}
                v60 := []int{}
                _ = v58
                _ = v59
                _ = v60
            }
            switch fi52 {
            case 42:
                fmt.Println(v51, v51)
            case 9223372036854775807:
                for fi61 := 0; fi61 < 9; fi61++ {
                    fmt.Println(fi61, fi52)
                    v62 := []int{fi61, 1, -2147483648}
                    _ = fi61
                    _ = v62
                }
            case 32767:
                for fi63 := 0; fi63 < 3; fi63++ {
                    var v64 uint8 = 0
                    var v65 float64 = float64(0.0)
                    _ = fi63
                    _ = v64
                    _ = v65
                }
            default:
                if false {
                    for fi66 := 0; fi66 < 18; fi66++ {
                        v67 := []int{(fi66 << ((552) & 63))}
                        var v68 float64 = (-2.25 * 1e10)
                        v51 = v51
                        _ = fi66
                        _ = v67
                        _ = v68
                    }
                    var v69 uint = 255
                    _ = v69
                }
            }
        } else {
            v70 := []int{(fi52 & fi52)}
            v71 := GenS9{F0: 1}
            for fi72 := 0; fi72 < 3; fi72++ {
                var v73 bool = (668 < -1)
                _ = fi72
                _ = v73
            }
            _ = v70
            _ = v71
        }
        var v74 uint16 = 42
        _ = fi52
        _ = v74
    }
    _ = v50
    _ = v51
    return 112
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[uint32]bool)
    var v3 uint = 0
    var v75 uint64 = (v1 + v1)
    v76 := []int{0}
    v1 -= 2
    for fi77 := 0; fi77 < 8; fi77++ {
        var v78 uint32 = uint32(v1)
        _ = fi77
        _ = v78
    }
    for fi79 := 0; fi79 < 11; fi79++ {
        switch "" {
        case "":
            switch fi79 {
            case 2147483647:
                var v80 uint8 = genFn1()
                _ = v80
            case 32768:
                _ = len(v2)
            case 0:
                v75 += 1
            default:
                if (!true) {
                    fmt.Println(v1, fi79)
                }
            }
        case "a":
            var v81 uint64 = 42
            _ = v81
        }
        v1 -= v1
        _ = fi79
    }
    var v82 uint16 = 32767
    if ((true || true) && true) {
        if false {
            var v83 string = ("hello" + "foo bar")
            _ = v83
        } else {
            var v84 string = "\n\t"
            _ = v84
        }
        if true {
            v1 += (^(v1))
            for fi85 := 0; fi85 < 7; fi85++ {
                for fi86 := 0; fi86 < 12; fi86++ {
                    var v87 uint16 = (v82 * (v82 & 2))
                    v88 := GenS4{F0: (v1 ^ (v75 % (((v75 - v1)) | 1))), F1: (^(v3)), F2: (!true)}
                    v89 := v2[737]
                    _ = fi86
                    _ = v87
                    _ = v88
                    _ = v89
                }
                v75 = (v75 / ((v75) | 1))
                _ = fi85
            }
        }
    } else {
        var v90 float32 = -2.25
        for fi91 := 0; fi91 < 10; fi91++ {
            if true {
                v75 = uint64(v82)
                var v92 int32 = -918
                _ = v92
            } else {
                var v93 float64 = (0.0 - 3.14159)
                _ = v93
            }
            v94 := GenS10{F0: (v82 % ((42) | 1)), F1: 1, F2: ""}
            fmt.Println(fi91, v82)
            _ = fi91
            _ = v94
        }
        v95 := GenS11{F0: v82, F1: 287}
        _ = v90
        _ = v95
    }
    fmt.Println(v3, v3)
    v76 = append(v76, 128)
    v82 = v82
    if false {
        var v96 int16 = -489
        _ = v96
    } else {
        v1 = (v1 ^ (v1 & (v1 >> ((v3) & 63))))
        if false {
            if (!true) {
                v75 += uint64(v82)
                v97 := GenS12{F0: 1000}
                if false {
                    if (true || true) {
                        var v98 bool = (127 == 100)
                        var v99 uint8 = genFn1()
                        switch -2147483648 {
                        case 4294967295:
                            v75 = (v75 + (v1 ^ (v75 >> (((v3 * 155)) & 63))))
                        }
                        _ = v98
                        _ = v99
                    }
                }
                _ = v97
            } else {
                delete(v2, 914)
            }
        }
    }
    switch 2147483647 {
    case -2147483648:
        if (false && (true && true)) {
            for fi100 := 0; fi100 < 14; fi100++ {
                v82 = (^(v82))
                _ = fi100
            }
            if (4294967295 >= 1) {
                var v101 int8 = 1
                var v102 int16 = 32767
                _ = v101
                _ = v102
            } else {
                if true {
                    if len(v76) > 0 {
                        v76[0] = -9223372036854775808
                    }
                    switch "hello" {
                    case "\n\t":
                        if false {
                            for fi103 := 0; fi103 < 16; fi103++ {
                                var v104 float64 = 0.1
                                fmt.Println(v104, v82, v3)
                                if (9223372036854775807 != 9223372036854775807) {
                                    fmt.Println(v1)
                                    for fi105 := 0; fi105 < 12; fi105++ {
                                        var v106 int32 = -2147483648
                                        if true {
                                            var v107 uint16 = v82
                                            v108 := GenS13{F0: -1, F1: (v3 % ((v3) | 1))}
                                            fmt.Println(v104, v3)
                                            _ = v107
                                            _ = v108
                                        }
                                        _ = fi105
                                        _ = v106
                                    }
                                } else {
                                    v1 = (v1 % ((uint64(v3)) | 1))
                                    fmt.Println(v1, v1, v104)
                                    v75 = 42
                                }
                                _ = fi103
                                _ = v104
                            }
                            v1 = v75
                            fmt.Println(v1, v3, v3)
                        }
                    default:
                        var v109 bool = false
                        _ = v109
                    }
                    v75 += v75
                } else {
                    switch "" {
                    case "hello":
                        for fi110 := 0; fi110 < 17; fi110++ {
                            if true {
                                v111 := GenS14{F0: 1000}
                                _ = v111
                            } else {
                                v112 := GenS15{F0: 0, F1: (false && (1.5 != 1e10))}
                                _ = v112
                            }
                            if (128 <= -655) {
                                var v113 int64 = 9223372036854775807
                                for fi114 := 0; fi114 < 17; fi114++ {
                                    var v115 uint64 = v1
                                    var v116 bool = (1 < -35)
                                    var v117 uint = (v3 + v3)
                                    if true {
                                        if (!v116) {
                                            v82 = uint16(v3)
                                            fmt.Println(v117)
                                        }
                                        var v118 float32 = (0.0 + 1e10)
                                        _ = v118
                                    }
                                    _ = fi114
                                    _ = v115
                                    _ = v116
                                    _ = v117
                                }
                                v113 += (v113 / ((0) | 1))
                                _ = v113
                            }
                            delete(v2, 4294967295)
                            _ = fi110
                        }
                    case "foo bar":
                        for fi119 := 0; fi119 < 2; fi119++ {
                            for fi120 := 0; fi120 < 20; fi120++ {
                                v121 := GenS8{F0: 180}
                                _ = fi120
                                _ = v121
                            }
                            v1 += (v1 << ((v3) & 63))
                            if len(v76) >= 2 {
                                v76 = v76[1:len(v76)]
                            }
                            _ = fi119
                        }
                    case "\n\t":
                        for fi122 := 0; fi122 < 7; fi122++ {
                            v123 := GenS16{F0: uint16(v3)}
                            _ = len(v2)
                            delete(v2, uint32(v1))
                            _ = fi122
                            _ = v123
                        }
                    default:
                        v75 += (^(v75))
                    }
                    for fi124 := 0; fi124 < 19; fi124++ {
                        v125 := GenS11{F0: 32, F1: v75}
                        for fi126 := 0; fi126 < 1; fi126++ {
                            if (!(!false)) {
                                if len(v76) >= 2 {
                                    v76 = v76[1:len(v76)]
                                }
                                switch "a" {
                                case "hello":
                                    if false {
                                        switch "hello" {
                                        case "\n\t":
                                            switch (^(fi126)) {
                                            case 32768:
                                                if false {
                                                    var v127 int32 = 0
                                                    v128 := GenS17{F0: (v82 / ((v82) | 1))}
                                                    if (genFn1() > 158) {
                                                        v82 = v82
                                                        v1 = (v75 >> ((v3) & 63))
                                                        if false {
                                                            var v129 uint = (v3 << (((v3 ^ v3)) & 63))
                                                            v3 -= v3
                                                            v75 = (v75 << ((v129) & 63))
                                                            _ = v129
                                                        } else {
                                                            v130 := GenS3{F0: -1, F1: (v1 & 42), F2: 107}
                                                            if true {
                                                                switch "a" {
                                                                case "foo bar":
                                                                    v75 = (v1 % ((v1) | 1))
                                                                case "a":
                                                                    switch fi124 {
                                                                    case 489:
                                                                        fmt.Println(v82)
                                                                    case -1:
                                                                        fmt.Println(v127, v75)
                                                                    case 128:
                                                                        if false {
                                                                            if false {
                                                                                if true {
                                                                                    fmt.Println(fi124)
                                                                                    v131 := GenS8{F0: 66}
                                                                                    v127 = int32(v82)
                                                                                    _ = v131
                                                                                } else {
                                                                                    switch (fi124 << ((18446744073709551615) & 63)) {
                                                                                    case 725:
                                                                                        if true {
                                                                                            v127 -= -1
                                                                                        }
                                                                                        v132 := GenS18{F0: 9223372036854775807}
                                                                                        fmt.Println(v75)
                                                                                        fmt.Println(v1, v1)
                                                                                        if (-2.25 > ((0.1 / -2.25) + ((1e10 - 1e10) + 3.14159))) {
                                                                                            if true {
                                                                                                v127 -= v127
                                                                                                var v133 uint = v3
                                                                                                fmt.Println(v133)
                                                                                                v1 += 2
                                                                                                v134 := GenS4{F0: 2, F1: (^(v133)), F2: ((!((!(4294967295 >= v3)) && (9223372036854775807 <= -1))) || false)}
                                                                                                _ = v133
                                                                                                _ = v134
                                                                                            } else {
                                                                                                switch (("a" + "a") + "") {
                                                                                                case "":
                                                                                                    v135 := GenS9{F0: 0}
                                                                                                    _ = v135
                                                                                                }
                                                                                            }
                                                                                            var v136 string = ""
                                                                                            _ = v136
                                                                                        } else {
                                                                                            var v137 uint16 = v82
                                                                                            v137 -= v82
                                                                                            _ = v137
                                                                                        }
                                                                                        v1 -= (^(v75))
                                                                                        fmt.Println(v82)
                                                                                        v127 += 0
                                                                                        var v138 int64 = 1
                                                                                        _ = v132
                                                                                        _ = v138
                                                                                    case -595:
                                                                                        v1 -= (v1 >> ((18446744073709551615) & 63))
                                                                                    }
                                                                                    v139 := GenS1{F0: "hello", F1: 4294967295, F2: (255 <= 99)}
                                                                                    _ = v139
                                                                                }
                                                                                var v140 float32 = 0.0
                                                                                _ = v140
                                                                            } else {
                                                                                fmt.Println(v1, fi126)
                                                                                var v141 uint = 65535
                                                                                _ = v141
                                                                            }
                                                                            v142 := GenS16{F0: 255}
                                                                            _ = v142
                                                                        }
                                                                    }
                                                                case "hello":
                                                                    if false {
                                                                        if true {
                                                                            if (false && ((false && false) || true)) {
                                                                                if true {
                                                                                    var v143 int16 = 32767
                                                                                    if true {
                                                                                        v82 = (v82 * uint16(v1))
                                                                                        v3 -= v3
                                                                                        var v144 int32 = v127
                                                                                        _ = v144
                                                                                    }
                                                                                    switch ("" + "foo bar") {
                                                                                    case "\n\t":
                                                                                        fmt.Println(v3)
                                                                                    }
                                                                                    _ = v143
                                                                                }
                                                                                v145 := GenS5{F0: (v127 * (v127 ^ v127))}
                                                                                fmt.Println(v1, fi124)
                                                                                _ = v145
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v127)
                                                                            v146 := GenS9{F0: 4294967295}
                                                                            v76 = append(v76, 42)
                                                                            _ = v146
                                                                        }
                                                                    } else {
                                                                        v147 := GenS7{F0: (100 > 93)}
                                                                        _ = v147
                                                                    }
                                                                default:
                                                                    v127 = 42
                                                                }
                                                            } else {
                                                                var v148 uint32 = 2
                                                                var v149 uint16 = 887
                                                                v75 -= (v1 >> (((v3 / ((v3) | 1))) & 63))
                                                                _ = v148
                                                                _ = v149
                                                            }
                                                            _ = v130
                                                        }
                                                        var v150 uint32 = 833
                                                        _ = v150
                                                    }
                                                    _ = v127
                                                    _ = v128
                                                } else {
                                                    if false {
                                                        v151 := GenS19{F0: 453}
                                                        _ = v151
                                                    } else {
                                                        v82 += (^(v82))
                                                        var v152 int16 = 32767
                                                        _ = v152
                                                    }
                                                }
                                            case -39:
                                                var v153 uint = 42
                                                _ = v153
                                            case -401:
                                                var v154 uint16 = 757
                                                _ = v154
                                            default:
                                                var v155 int8 = -128
                                                _ = v155
                                            }
                                        case "":
                                            var v156 float64 = 0.1
                                            _ = v156
                                        }
                                        _ = len(v76)
                                    }
                                case "a":
                                    v82 = (v82 ^ v82)
                                case "\n\t":
                                    v157 := GenS5{F0: 0}
                                    _ = v157
                                }
                            }
                            var v158 int64 = 57
                            _ = fi126
                            _ = v158
                        }
                        var v159 int64 = -1
                        _ = fi124
                        _ = v125
                        _ = v159
                    }
                }
                for fi160 := 0; fi160 < 7; fi160++ {
                    v1 = v75
                    _ = fi160
                }
            }
        }
    case 1:
        var v161 uint32 = 0
        _ = v161
    case 644:
        if (!(1 >= 2147483647)) {
            v1 = v1
            v3 -= 1
        } else {
            if ((v75 * v1) <= v1) {
                switch -2147483648 {
                case -1:
                    var v162 int16 = -879
                    _ = v162
                case -9223372036854775808:
                    var v163 uint = (v3 ^ (v3 * 1))
                    _ = v163
                default:
                    fmt.Println(v3)
                }
                fmt.Println(v3, v3)
                for fi164 := 0; fi164 < 8; fi164++ {
                    v165 := GenS20{F0: int32(v3), F1: (fi164 % ((fi164) | 1)), F2: 450}
                    _ = fi164
                    _ = v165
                }
            }
            for fi166 := 0; fi166 < 19; fi166++ {
                var v167 float32 = -2.25
                v75 += (v1 >> ((339) & 63))
                if ((1e10 + (0.0 / -2.25)) == 1.5) {
                    v2[0] = ((v82 & (v82 >> ((903) & 63))) == (v82 * (v82 | v82)))
                    v76 = append(v76, (fi166 >> ((v3) & 63)))
                } else {
                    switch "hello" {
                    case "a":
                        switch fi166 {
                        case 1:
                            v3 = v3
                        case 2147483647:
                            if false {
                                v168 := GenS10{F0: v82, F1: 127, F2: "hello"}
                                if (!true) {
                                    var v169 float32 = v167
                                    v170 := GenS3{F0: -1, F1: (v75 + v75), F2: 948}
                                    var v171 float32 = v169
                                    _ = v169
                                    _ = v170
                                    _ = v171
                                } else {
                                    v167 = (v167 * float32(-2.25))
                                }
                                _ = v168
                            } else {
                                v75 -= (v1 ^ v75)
                            }
                        case 128:
                            switch "a" {
                            case "foo bar":
                                var v172 int64 = -9223372036854775808
                                _ = v172
                            case "hello":
                                v82 = (v82 * v82)
                            default:
                                var v173 float32 = (v167 * 0.1)
                                _ = v173
                            }
                        default:
                            var v174 int8 = -45
                            _ = v174
                        }
                    case "hello":
                        var v175 uint64 = (^(v75))
                        _ = v175
                    case "foo bar":
                        var v176 float32 = float32(0.1)
                        _ = v176
                    default:
                        v82 = (v82 % (((v82 ^ 1)) | 1))
                    }
                    var v177 int64 = 9223372036854775807
                    switch -2147483648 {
                    case 32767:
                        for fi178 := 0; fi178 < 3; fi178++ {
                            if false {
                                v1 = (^(v75))
                            } else {
                                fmt.Println(v177, v3)
                            }
                            _ = fi178
                        }
                    case -1:
                        fmt.Println(v75, v75)
                    default:
                        fmt.Println(v82)
                    }
                    _ = v177
                }
                _ = fi166
                _ = v167
            }
            var v179 int64 = 352
            _ = v179
        }
    default:
        var v180 uint = (v3 ^ (v3 / ((v3) | 1)))
        _ = v180
    }
    var v181 uint8 = 0
    delete(v2, 42)
    var v182 int8 = 58
    fmt.Println(v3)
    v183 := GenS21{F0: 1, F1: uint(v82), F2: -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v181:", v181)
    fmt.Println("v182:", v182)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v75:", v75)
    fmt.Println("slicelen v76:", len(v76), v76)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v181
    _ = v182
    _ = v183
    _ = v2
    _ = v3
    _ = v75
    _ = v76
    _ = v82
}
