package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 uint8; F1 int8; F2 int }
type GenS7 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS8 struct { F0 uint16; F1 int32; F2 int8 }
type GenS9 struct { F0 int32; F1 uint64; F2 string }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint64; F1 int16 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 uint32; F1 uint32; F2 uint16 }
func genFn0(p0 string, p1 bool) int64 {
    v5 := []int{32768, 672}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = p0
    _ = p1
    _ = v5
    return -1
}

func main() {
    var v1 uint32 = 42
    v2 := []int{32768}
    v3 := []int{42}
    var v4 uint = 0
    switch 32768 {
    case 127:
        v6 := GenS1{F0: uint(v1), F1: 1}
        _ = v6
    case 717:
        var v7 int32 = 1
        _ = v7
    case 32767:
        var v8 float64 = ((float64(float32(0.0)) + (1e10 / (float64(-2.25) + 3.14159))) - (3.14159 * (-2.25 / 0.1)))
        _ = v8
    }
    v1 = (v1 % ((v1) | 1))
    var v9 int32 = -586
    for fi10 := 0; fi10 < 5; fi10++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi10
    }
    if false {
        var v11 float32 = 0.1
        _ = len(v3)
        _ = v11
    } else {
        var v12 bool = false
        fmt.Println(v9)
        _ = v12
    }
    fmt.Println(v9, v4, v9)
    if len(v2) > 0 {
        v2[0] = 42
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        if false {
            if false {
                switch (fi13 + fi13) {
                case 71:
                    switch 128 {
                    case -86:
                        if (false || true) {
                            if true {
                                var v14 int = 127
                                var v15 string = ""
                                _ = v14
                                _ = v15
                            } else {
                                v16 := GenS1{F0: 18446744073709551615, F1: 62}
                                _ = v16
                            }
                            v2 = append(v2, 127)
                        }
                    case -1:
                        v17 := GenS1{F0: (v4 * v4), F1: 1}
                        _ = v17
                    case 127:
                        for fi18 := 0; fi18 < 1; fi18++ {
                            v1 = v1
                            var v19 uint64 = 727
                            _ = fi18
                            _ = v19
                        }
                    default:
                        if (-696 == 9223372036854775807) {
                            var v20 uint8 = 1
                            if false {
                                fmt.Println(v1)
                            } else {
                                fmt.Println(fi13)
                            }
                            switch (fi13 / ((fi13) | 1)) {
                            case 4294967295:
                                for fi21 := 0; fi21 < 12; fi21++ {
                                    var v22 int8 = 100
                                    fmt.Println(fi13, v9, fi21)
                                    _ = fi21
                                    _ = v22
                                }
                            case 835:
                                if true {
                                    fmt.Println(v4)
                                    for fi23 := 0; fi23 < 12; fi23++ {
                                        switch 763 {
                                        case -713:
                                            if true {
                                                v24 := GenS2{F0: -89, F1: 107}
                                                if (42 == 100) {
                                                    fmt.Println(v4, v20)
                                                } else {
                                                    var v25 float64 = 0.1
                                                    _ = v25
                                                }
                                                v9 = (-(v9))
                                                _ = v24
                                            } else {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                        case 115:
                                            switch "hello" {
                                            case "\n\t":
                                                var v26 uint = v4
                                                _ = v26
                                            }
                                        default:
                                            v1 = (v1 + (^(v1)))
                                        }
                                        if (-1 >= -121) {
                                            var v27 int8 = 100
                                            if true {
                                                switch fi23 {
                                                case 42:
                                                    v1 = v1
                                                default:
                                                    var v28 float64 = 1e10
                                                    _ = v28
                                                }
                                                if (((1.5 + 0.1) <= 1e10) && true) {
                                                    v29 := GenS1{F0: v4, F1: v27}
                                                    if false {
                                                        var v30 int32 = v9
                                                        switch 32767 {
                                                        case -9223372036854775808:
                                                            if true {
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                v1 = v1
                                                                if (v27 > v27) {
                                                                    v31 := GenS3{F0: v20, F1: 18446744073709551615}
                                                                    _ = v31
                                                                } else {
                                                                    var v32 bool = false
                                                                    v33 := GenS2{F0: int64(v30), F1: (v27 * (v27 / ((122) | 1)))}
                                                                    _ = v32
                                                                    _ = v33
                                                                }
                                                            } else {
                                                                v20 -= v20
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            }
                                                        case 42:
                                                            switch (fi23 << ((919) & 63)) {
                                                            case 9223372036854775807:
                                                                fmt.Println(fi23, v27)
                                                            case 42:
                                                                if (!((!false) || false)) {
                                                                    v3 = append(v3, fi13)
                                                                    v20 = 128
                                                                } else {
                                                                    if ((true && true) || (false || false)) {
                                                                        var v34 uint8 = v20
                                                                        _ = v34
                                                                    } else {
                                                                        var v35 uint16 = 0
                                                                        if (193 < 0) {
                                                                            if false {
                                                                                fmt.Println(v4)
                                                                            } else {
                                                                                var v36 float64 = float64((3.14159 * 0.1))
                                                                                v37 := GenS3{F0: 127, F1: 42}
                                                                                var v38 float32 = 1.5
                                                                                _ = v36
                                                                                _ = v37
                                                                                _ = v38
                                                                            }
                                                                            v30 -= 459
                                                                        } else {
                                                                            v27 = v27
                                                                            v4 -= 733
                                                                            v39 := GenS4{F0: 255, F1: -1}
                                                                            v1 -= (v1 / (((v1 >> (((v4 << (((v4 >> (((v4 | (^(v4)))) & 63))) & 63))) & 63))) | 1))
                                                                            switch fi13 {
                                                                            case 42:
                                                                                if (((!(false || true)) || true) && (42 <= int64(v35))) {
                                                                                    var v40 uint8 = 0
                                                                                    fmt.Println(v1)
                                                                                    _ = v40
                                                                                } else {
                                                                                    var v41 int8 = (v27 * v27)
                                                                                    v42 := GenS4{F0: v35, F1: 576}
                                                                                    if ((v30 - v30) < -55) {
                                                                                        switch "" {
                                                                                        case "a":
                                                                                            fmt.Println(fi13)
                                                                                        case "hello":
                                                                                            if (0.1 != 0.0) {
                                                                                                v43 := GenS5{F0: v9, F1: (!(true && false))}
                                                                                                _ = v43
                                                                                            }
                                                                                            v4 = v4
                                                                                            if (18446744073709551615 == 253) {
                                                                                                _ = len(v2)
                                                                                                v30 -= v9
                                                                                                var v44 int16 = -32768
                                                                                                _ = v44
                                                                                            } else {
                                                                                                var v45 float64 = 3.14159
                                                                                                v46 := GenS3{F0: v20, F1: uint64(v4)}
                                                                                                _ = v45
                                                                                                _ = v46
                                                                                            }
                                                                                        case "":
                                                                                            var v47 int16 = -1
                                                                                            _ = v47
                                                                                        }
                                                                                        switch ("hello" + ("\n\t" + "\n\t")) {
                                                                                        case "":
                                                                                            v48 := GenS6{F0: (v20 - (^(v20))), F1: (v27 ^ v41), F2: fi23}
                                                                                            var v49 bool = false
                                                                                            _ = v48
                                                                                            _ = v49
                                                                                        case "hello":
                                                                                            switch int(v4) {
                                                                                            case 0:
                                                                                                if true {
                                                                                                    if true {
                                                                                                        v20 = v20
                                                                                                        var v50 int = fi23
                                                                                                        fmt.Println(v35)
                                                                                                        _ = v50
                                                                                                    }
                                                                                                    v20 -= (v20 + v20)
                                                                                                }
                                                                                            case 909:
                                                                                                v27 -= v27
                                                                                            }
                                                                                        }
                                                                                        v9 -= (v30 + (v9 >> ((v4) & 63)))
                                                                                        if true {
                                                                                            var v51 uint16 = v35
                                                                                            _ = v51
                                                                                        } else {
                                                                                            v41 = v27
                                                                                        }
                                                                                        var v52 int16 = 32767
                                                                                        var v53 uint8 = v20
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                        switch "a" {
                                                                                        case "\n\t":
                                                                                            var v54 uint16 = (v35 << ((v4) & 63))
                                                                                            _ = v54
                                                                                        case "hello":
                                                                                            var v55 float64 = 1.5
                                                                                            _ = v55
                                                                                        default:
                                                                                            var v56 int16 = int16(v4)
                                                                                            _ = v56
                                                                                        }
                                                                                        switch "a" {
                                                                                        case "\n\t":
                                                                                            if ((v35 << ((v4) & 63)) > v35) {
                                                                                                var v57 int = fi13
                                                                                                var v58 int = fi13
                                                                                                _ = v57
                                                                                                _ = v58
                                                                                            }
                                                                                        }
                                                                                        _ = v52
                                                                                        _ = v53
                                                                                    }
                                                                                    v59 := GenS7{F0: (^(v35)), F1: (v1 / (((v1 | (v1 >> (((v4 << ((v4) & 63))) & 63)))) | 1)), F2: v20}
                                                                                    fmt.Println(v9, v35)
                                                                                    _ = v41
                                                                                    _ = v42
                                                                                    _ = v59
                                                                                }
                                                                            case 0:
                                                                                v60 := GenS7{F0: uint16(v1), F1: v1, F2: uint8(v4)}
                                                                                _ = v60
                                                                            case 104:
                                                                                var v61 int32 = (v30 >> (((^(v4))) & 63))
                                                                                _ = v61
                                                                            }
                                                                            _ = v39
                                                                        }
                                                                        _ = v35
                                                                    }
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    v20 -= (^(v20))
                                                                }
                                                            default:
                                                                v27 = -1
                                                                if ((v30 & v30) == v30) {
                                                                    v9 = (v30 << ((255) & 63))
                                                                    v62 := GenS7{F0: 42, F1: v1, F2: (^(v20))}
                                                                    var v63 int32 = -176
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 843
                                                                    }
                                                                    switch (fi23 | fi13) {
                                                                    case 32767:
                                                                        v64 := GenS7{F0: 32767, F1: v1, F2: (v20 & (v20 >> ((4294967295) & 63)))}
                                                                        _ = v64
                                                                    default:
                                                                        if true {
                                                                            switch "\n\t" {
                                                                            case "hello":
                                                                                var v65 float32 = 0.0
                                                                                _ = v65
                                                                            case "a":
                                                                                var v66 int16 = 627
                                                                                _ = v66
                                                                            case "\n\t":
                                                                                v1 += (v1 << ((v4) & 63))
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            default:
                                                                                fmt.Println(v9, v1, fi23)
                                                                            }
                                                                            v9 -= v9
                                                                            switch fi13 {
                                                                            case 447:
                                                                                if (858 >= 18446744073709551615) {
                                                                                    if false {
                                                                                        v67 := GenS8{F0: 628, F1: 42, F2: v27}
                                                                                        var v68 uint = (v4 | (v4 * v4))
                                                                                        switch (fi13 ^ (fi13 - -9223372036854775808)) {
                                                                                        case -374:
                                                                                            v69 := GenS9{F0: -2147483648, F1: 2, F2: ("a" + "")}
                                                                                            _ = v69
                                                                                        default:
                                                                                            fmt.Println(fi23)
                                                                                        }
                                                                                        v2 = append(v2, -2147483648)
                                                                                        v70 := GenS5{F0: v9, F1: true}
                                                                                        _ = v67
                                                                                        _ = v68
                                                                                        _ = v70
                                                                                    }
                                                                                }
                                                                            default:
                                                                                if (!true) {
                                                                                    v71 := GenS10{F0: (0 != 461)}
                                                                                    _ = v71
                                                                                } else {
                                                                                    var v72 int64 = int64(fi23)
                                                                                    v72 = v72
                                                                                    _ = v72
                                                                                }
                                                                            }
                                                                            v73 := GenS7{F0: uint16(v4), F1: v1, F2: (v20 / ((v20) | 1))}
                                                                            v74 := GenS9{F0: v30, F1: 901, F2: "a"}
                                                                            _ = v73
                                                                            _ = v74
                                                                        }
                                                                    }
                                                                    v27 = 42
                                                                    var v75 uint64 = 4294967295
                                                                    v76 := GenS2{F0: 1, F1: v27}
                                                                    _ = v62
                                                                    _ = v63
                                                                    _ = v75
                                                                    _ = v76
                                                                } else {
                                                                    v4 -= (v4 | (v4 & (v4 % ((v4) | 1))))
                                                                }
                                                                var v77 uint32 = (v1 % ((v1) | 1))
                                                                _ = v77
                                                            }
                                                        }
                                                        var v78 float64 = 0.1
                                                        if true {
                                                            switch ("a" + "foo bar") {
                                                            case "foo bar":
                                                                v79 := GenS10{F0: true}
                                                                _ = v79
                                                            case "hello":
                                                                var v80 int16 = 0
                                                                _ = v80
                                                            case "\n\t":
                                                                v30 -= v9
                                                            }
                                                        } else {
                                                            if false {
                                                                switch ("\n\t" + "foo bar") {
                                                                case "\n\t":
                                                                    var v81 int64 = 42
                                                                    _ = v81
                                                                case "foo bar":
                                                                    v82 := GenS11{F0: "foo bar"}
                                                                    _ = v82
                                                                }
                                                                v4 -= v4
                                                                v20 = (v20 * v20)
                                                            } else {
                                                                v83 := GenS5{F0: v9, F1: true}
                                                                v84 := GenS4{F0: 255, F1: 182}
                                                                _ = v83
                                                                _ = v84
                                                            }
                                                            v30 -= v30
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        _ = v30
                                                        _ = v78
                                                    }
                                                    _ = v29
                                                } else {
                                                    if (32768 < fi13) {
                                                        var v85 float32 = (0.0 + 0.0)
                                                        var v86 int64 = -34
                                                        switch ("hello" + "\n\t") {
                                                        case "hello":
                                                            switch "foo bar" {
                                                            case "a":
                                                                var v87 int64 = (-(v86))
                                                                _ = v87
                                                            case "\n\t":
                                                                switch "a" {
                                                                case "a":
                                                                    fmt.Println(v85, v85, v9)
                                                                case "\n\t":
                                                                    var v88 int32 = -1
                                                                    _ = v88
                                                                }
                                                            }
                                                        case "foo bar":
                                                            switch "a" {
                                                            case "foo bar":
                                                                var v89 int16 = 1000
                                                                _ = v89
                                                            }
                                                        default:
                                                            if true {
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    v90 := GenS8{F0: 0, F1: (^(v9)), F2: (-(v27))}
                                                                    _ = v90
                                                                default:
                                                                    v91 := GenS12{F0: 1, F1: -509}
                                                                    _ = v91
                                                                }
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                v92 := GenS13{F0: 577}
                                                                _ = v92
                                                            } else {
                                                                var v93 int16 = -570
                                                                _ = v93
                                                            }
                                                        }
                                                        _ = v85
                                                        _ = v86
                                                    }
                                                    if false {
                                                        var v94 uint = v4
                                                        _ = len(v3)
                                                        if false {
                                                            v4 = uint(v1)
                                                            v95 := GenS14{F0: "\n\t"}
                                                            _ = v95
                                                        }
                                                        _ = v94
                                                    }
                                                }
                                            }
                                            var v96 uint = (v4 / ((v4) | 1))
                                            v97 := GenS11{F0: "foo bar"}
                                            var v98 uint = (v96 >> (((v96 << (((v96 >> (((^(v96))) & 63))) & 63))) & 63))
                                            v9 += v9
                                            _ = v27
                                            _ = v96
                                            _ = v97
                                            _ = v98
                                        } else {
                                            var v99 uint8 = uint8(v4)
                                            _ = v99
                                        }
                                        _ = fi23
                                    }
                                    for fi100 := 0; fi100 < 2; fi100++ {
                                        var v101 int = 32768
                                        v20 = (v20 >> ((v4) & 63))
                                        _ = fi100
                                        _ = v101
                                    }
                                } else {
                                    v1 += 2147483648
                                    v4 -= 255
                                }
                            case -9223372036854775808:
                                v4 += v4
                            }
                            _ = v20
                        } else {
                            v102 := GenS15{F0: false}
                            var v103 float32 = 1e10
                            _ = v102
                            _ = v103
                        }
                    }
                case -987:
                    for fi104 := 0; fi104 < 15; fi104++ {
                        if false {
                            v4 = v4
                            v105 := GenS2{F0: 470, F1: 42}
                            var v106 bool = true
                            _ = v105
                            _ = v106
                        }
                        _ = fi104
                    }
                case 9223372036854775807:
                    v107 := GenS16{F0: 299, F1: v1, F2: 1}
                    _ = v107
                default:
                    for fi108 := 0; fi108 < 8; fi108++ {
                        switch (fi108 % ((fi108) | 1)) {
                        case 1:
                            var v109 int8 = 89
                            v9 -= (v9 - v9)
                            if ((v1 + v1) >= v1) {
                                var v110 int16 = -1
                                fmt.Println(v109)
                                _ = v110
                            }
                            _ = v109
                        default:
                            fmt.Println(v4)
                        }
                        _ = fi108
                    }
                }
            } else {
                v9 = (v9 / ((v9) | 1))
                fmt.Println(v9)
            }
            v4 += v4
        }
        v111 := GenS5{F0: v9, F1: (-1 != 611)}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi13
        _ = v111
    }
    v112 := GenS13{F0: -32768}
    fmt.Println(v1, v4, v1)
    v9 -= (v9 & -1)
    v9 = v9
    for fi113 := 0; fi113 < 5; fi113++ {
        if (!false) {
            v9 = -1
            v114 := GenS14{F0: "foo bar"}
            _ = v114
        } else {
            v4 = 42
        }
        if true {
            var v115 uint32 = v1
            switch fi113 {
            case 9223372036854775807:
                v115 -= 900
            case 106:
                var v116 uint16 = 255
                _ = v116
            default:
                if (!true) {
                    if (!true) {
                        switch "a" {
                        case "\n\t":
                            v117 := GenS9{F0: 1, F1: 18446744073709551615, F2: ""}
                            _ = v117
                        case "foo bar":
                            v115 = (v115 * 1)
                        default:
                            v118 := GenS2{F0: 0, F1: -128}
                            _ = v118
                        }
                    }
                } else {
                    if (true || ((1.5 * -2.25) > (float32(-2.25) - 3.14159))) {
                        v119 := GenS4{F0: 2, F1: 32767}
                        v120 := GenS14{F0: "hello"}
                        _ = v119
                        _ = v120
                    }
                }
            }
            var v121 int64 = 42
            _ = v115
            _ = v121
        }
        switch ("" + "") {
        case "a":
            v4 += 627
        }
        _ = fi113
    }
    var v122 int64 = 9223372036854775807
    v4 += v4
    v123 := GenS6{F0: 128, F1: -1, F2: 1}
    var v124 int = 2147483647
    v125 := GenS15{F0: false}
    fmt.Println("v1:", v1)
    fmt.Println("v122:", v122)
    fmt.Println("v124:", v124)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v112
    _ = v122
    _ = v123
    _ = v124
    _ = v125
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
