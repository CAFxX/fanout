package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 bool; F1 uint16 }
type GenS8 struct { F0 uint64; F1 int8; F2 bool }
type GenS9 struct { F0 uint; F1 uint; F2 uint8 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 bool; F1 bool; F2 int64 }
type GenS12 struct { F0 uint8; F1 uint32 }
type GenS13 struct { F0 int8; F1 int8; F2 string }
type GenS14 struct { F0 bool; F1 int8 }
type GenS15 struct { F0 int32; F1 bool }
type GenS16 struct { F0 int64; F1 int32; F2 int }
type GenS17 struct { F0 int16; F1 uint64 }
type GenS18 struct { F0 int8 }
type GenS19 struct { F0 int }
type GenS20 struct { F0 uint16; F1 int16; F2 uint }
type GenS21 struct { F0 uint16 }
type GenS22 struct { F0 uint32; F1 int16 }
type GenS23 struct { F0 string; F1 string }
type GenS24 struct { F0 uint; F1 uint16; F2 int8 }
type GenS25 struct { F0 int8 }
type GenS26 struct { F0 uint8 }
type GenS27 struct { F0 uint16; F1 int8 }
type GenS28 struct { F0 int8; F1 string }
func genFn0(p0 int, p1 uint32, p2 int16) uint64 {
    p1 += p1
    var v7 uint16 = 0
    v7 += v7
    var v8 int64 = 1
    v9 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 273
}
func genFn1(p0 uint64, p1 string, p2 int) int32 {
    var v10 uint8 = 128
    switch p2 {
    case 32767:
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 uint8 = (v10 * (v10 | v10))
            _ = fi11
            _ = v12
        }
    case 128:
        var v13 float64 = (-2.25 * (0.0 - 3.14159))
        _ = v13
    case 259:
        fmt.Println(p0)
    }
    v14 := []int{(p2 | p2), p2}
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 bool = false
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        _ = fi15
        _ = v16
    }
    var v17 int = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v14
    _ = v17
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = 3.14159
    v3 := []uint32{714, 42, 2147483648}
    var v4 string = "a"
    var v5 bool = (46 != 255)
    var v6 uint16 = 0
    var v18 int = -2147483648
    if v5 {
        v19 := GenS1{F0: v5, F1: 0, F2: 0}
        _ = v19
    } else {
        v20 := GenS1{F0: v5, F1: 42, F2: 602}
        switch (v18 >> ((65535) & 63)) {
        case -593:
            v21 := GenS2{F0: ("a" + v4), F1: 65535}
            _ = v21
        case 32767:
            v6 = v6
        case 0:
            if v5 {
                v4 = v4
            } else {
                v22 := GenS1{F0: true, F1: 1, F2: 65535}
                v23 := GenS3{F0: (v6 + v6), F1: -1}
                _ = v22
                _ = v23
            }
        default:
            delete(v1, 0)
        }
        if (v6 <= v6) {
            var v24 uint16 = 2
            for fi25 := 0; fi25 < 5; fi25++ {
                v24 = 2
                for fi26 := 0; fi26 < 18; fi26++ {
                    if (v18 > fi26) {
                        v18 -= fi25
                        if v5 {
                            var v27 uint8 = 1
                            _ = len(v3)
                            _ = v27
                        }
                        v28 := GenS3{F0: v24, F1: -1}
                        _ = v28
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        switch (v18 % (((fi26 & fi26)) | 1)) {
                        case 32767:
                            v4 = v4
                            if (!(!(!(((fi26 + fi25) >= (fi26 ^ fi25)) && v5)))) {
                                v24 = 42
                                var v29 uint64 = 18446744073709551615
                                _ = v29
                            } else {
                                fmt.Println(v24)
                                if len(v3) > 0 {
                                    v3[0] = 777
                                }
                            }
                        }
                        if v5 {
                            if true {
                                switch (fi25 << ((139) & 63)) {
                                case -296:
                                    var v30 uint32 = 138
                                    _ = v30
                                case 42:
                                    v6 = v24
                                }
                                if v5 {
                                    fmt.Println(fi25, v4)
                                    v31 := GenS3{F0: 293, F1: 42}
                                    _ = v31
                                }
                                _ = len(v3)
                            } else {
                                v32 := GenS2{F0: (v4 + "foo bar"), F1: 779}
                                _ = v32
                            }
                            var v33 int8 = 1
                            v6 += 32767
                            _ = v33
                        } else {
                            var v34 float64 = 0.1
                            switch 32767 {
                            case 0:
                                v4 = (("foo bar" + "\n\t") + "")
                            default:
                                var v35 int = (fi25 << ((65535) & 63))
                                _ = v35
                            }
                            _ = v34
                        }
                        var v36 int16 = -579
                        v37 := GenS4{F0: 2, F1: 1}
                        _ = v36
                        _ = v37
                    }
                    if v5 {
                        fmt.Println(fi26)
                        v38 := GenS3{F0: v24, F1: -1}
                        v24 -= 255
                        _ = v38
                    }
                    _ = fi26
                }
                _ = fi25
            }
            if (443 >= 529) {
                _ = len(v1)
                for fi39 := 0; fi39 < 20; fi39++ {
                    switch v4 {
                    case "":
                        var v40 bool = true
                        _ = v40
                    case "hello":
                        for fi41 := 0; fi41 < 11; fi41++ {
                            switch v4 {
                            case "foo bar":
                                switch v4 {
                                case "":
                                    if len(v3) > 0 {
                                        v3[0] = 521
                                    }
                                case "foo bar":
                                    v2 += ((v2 + (float64(-2.25) * 1e10)) - float64(1e10))
                                case "\n\t":
                                    v2 = 0.1
                                default:
                                    v42 := GenS1{F0: ("hello" == (v4 + "a")), F1: 0, F2: 4294967295}
                                    switch v18 {
                                    case -1:
                                        _ = len(v1)
                                    case 32768:
                                        if (!false) {
                                            var v43 int8 = 42
                                            v2 += float64((((-2.25 + 1e10) * 3.14159) * ((1.5 / (float32((v2 * (float64((1.5 + 3.14159)) * 1.5))) * -2.25)) / 0.1)))
                                            _ = v43
                                        } else {
                                            switch v4 {
                                            case "hello":
                                                v5 = false
                                            default:
                                                var v44 int = -792
                                                _ = v44
                                            }
                                            v45 := GenS1{F0: (v4 <= "foo bar"), F1: 784, F2: 0}
                                            var v46 int8 = 1
                                            switch -2147483648 {
                                            case -9223372036854775808:
                                                switch 32768 {
                                                case 802:
                                                    fmt.Println(fi39, fi41)
                                                case 223:
                                                    v47 := GenS4{F0: 117, F1: (^(v24))}
                                                    _ = v47
                                                case -183:
                                                    var v48 int16 = -875
                                                    _ = v48
                                                default:
                                                    v49 := GenS5{F0: 42, F1: v4}
                                                    _ = v49
                                                }
                                            }
                                            switch ("\n\t" + v4) {
                                            case "":
                                                v50 := v1[0]
                                                _ = v50
                                            case "hello":
                                                var v51 float32 = 0.0
                                                _ = v51
                                            default:
                                                switch (("\n\t" + ("hello" + (v4 + v4))) + v4) {
                                                case "foo bar":
                                                    var v52 float32 = -2.25
                                                    _ = v52
                                                }
                                            }
                                            _ = v45
                                            _ = v46
                                        }
                                    case 127:
                                        v2 += v2
                                    }
                                    fmt.Println(v6, v18)
                                    _ = v42
                                }
                            }
                            v4 = v4
                            if v5 {
                                var v53 float64 = float64(3.14159)
                                var v54 string = v4
                                v18 = (fi39 / (((fi41 + v18)) | 1))
                                var v55 int16 = 32767
                                if (1 == -2147483648) {
                                    v53 -= 1.5
                                    _ = len(v3)
                                    v55 = v55
                                } else {
                                    var v56 uint16 = 255
                                    v57 := v1[369]
                                    _ = v56
                                    _ = v57
                                }
                                v2 = v53
                                if len(v3) > 0 {
                                    v3[0] = 2
                                }
                                if v5 {
                                    switch v4 {
                                    case "":
                                        if (690 > 82) {
                                            if v5 {
                                                v5 = false
                                                var v58 uint32 = 65535
                                                v24 += (v24 ^ 255)
                                                _ = v58
                                            } else {
                                                var v59 uint16 = v24
                                                fmt.Println(v55)
                                                switch fi41 {
                                                case -827:
                                                    v4 = v4
                                                case -822:
                                                    v59 += (v6 | (^(v6)))
                                                case 983:
                                                    fmt.Println(fi41, fi41)
                                                }
                                                _ = v59
                                            }
                                        }
                                    case "foo bar":
                                        if v5 {
                                            v18 -= 2147483647
                                        } else {
                                            if true {
                                                delete(v1, 2147483648)
                                                fmt.Println(fi41, v54)
                                            } else {
                                                v60 := GenS6{F0: 255}
                                                _ = v60
                                            }
                                        }
                                    default:
                                        var v61 int = (fi41 & (v18 ^ (v18 * (-(fi39)))))
                                        _ = v61
                                    }
                                    v24 = v6
                                }
                                _ = v53
                                _ = v54
                                _ = v55
                            }
                            _ = fi41
                        }
                    default:
                        for fi62 := 0; fi62 < 20; fi62++ {
                            switch "" {
                            case "foo bar":
                                if v5 {
                                    v63 := GenS3{F0: (v6 * v6), F1: 42}
                                    var v64 uint64 = 1
                                    v1[869] = 2
                                    _ = v63
                                    _ = v64
                                } else {
                                    v65 := GenS1{F0: v5, F1: 4294967295, F2: 723}
                                    _ = v65
                                }
                            default:
                                v66 := GenS3{F0: (v24 | (v24 ^ 65535)), F1: -42}
                                _ = v66
                            }
                            _ = fi62
                        }
                    }
                    _ = fi39
                }
                v67 := v1[1]
                _ = v67
            }
            _ = v24
        }
        _ = v20
    }
    v6 += v6
    if (!v5) {
        v68 := GenS2{F0: ("" + ("hello" + ("foo bar" + ((v4 + "hello") + (("a" + v4) + (v4 + ("hello" + "a"))))))), F1: 823}
        v6 = v6
        _ = v68
    } else {
        v5 = (1.5 >= 1.5)
    }
    var v69 uint32 = 65535
    switch (("\n\t" + "hello") + ((v4 + ("" + "\n\t")) + ("a" + v4))) {
    case "\n\t":
        if ((v18 >= (v18 / ((-2147483648) | 1))) || (v4 <= v4)) {
            var v70 uint = 889
            _ = len(v1)
            _ = v70
        }
    case "hello":
        v2 += v2
    case "":
        var v71 uint = 42
        _ = v71
    default:
        fmt.Println(v4)
    }
    if true {
        v72 := GenS2{F0: v4, F1: (v69 & uint32(v6))}
        var v73 float32 = -2.25
        _ = v72
        _ = v73
    }
    switch "" {
    case "\n\t":
        if (!true) {
            for fi74 := 0; fi74 < 17; fi74++ {
                v75 := GenS7{F0: v5, F1: (^(v6))}
                v2 -= ((float64(1.5) * v2) / 1.0)
                v76 := GenS8{F0: 2, F1: int8(fi74), F2: (!(!v5))}
                _ = fi74
                _ = v75
                _ = v76
            }
        }
    case "hello":
        v2 = v2
    case "":
        for fi77 := 0; fi77 < 8; fi77++ {
            v6 += (v6 & v6)
            var v78 uint = 4294967295
            var v79 string = "a"
            _ = fi77
            _ = v78
            _ = v79
        }
    default:
        v69 = (v69 - 625)
    }
    v2 -= (v2 + (3.14159 + (v2 * v2)))
    if (int16(v18) < -32768) {
        v1[(^(v69))] = 0
        fmt.Println(v5)
        if (-2.25 > 1e10) {
            for fi80 := 0; fi80 < 18; fi80++ {
                v5 = (42 <= -1)
                _ = fi80
            }
            v6 += (v6 / (((v6 << ((1) & 63))) | 1))
        } else {
            var v81 uint = 938
            _ = v81
        }
    } else {
        var v82 int8 = 127
        v83 := GenS6{F0: 8}
        _ = v82
        _ = v83
    }
    if (1000 < -1) {
        switch v4 {
        case "a":
            switch "a" {
            case "hello":
                var v84 uint64 = 42
                _ = v84
            }
        case "\n\t":
            v6 -= v6
        }
        switch v4 {
        case "foo bar":
            for fi85 := 0; fi85 < 15; fi85++ {
                v2 -= 0.1
                v86 := GenS9{F0: 254, F1: 2, F2: 0}
                _ = fi85
                _ = v86
            }
        default:
            for fi87 := 0; fi87 < 3; fi87++ {
                switch "hello" {
                case "\n\t":
                    if ((v5 || true) || v5) {
                        for fi88 := 0; fi88 < 17; fi88++ {
                            v3 = append(v3, v69)
                            _ = fi88
                        }
                        for fi89 := 0; fi89 < 1; fi89++ {
                            fmt.Println(v18)
                            _ = fi89
                        }
                        if (2 >= 4294967295) {
                            for fi90 := 0; fi90 < 8; fi90++ {
                                switch 9223372036854775807 {
                                case -771:
                                    v4 = ((("a" + v4) + v4) + v4)
                                }
                                _ = fi90
                            }
                            switch v18 {
                            case 128:
                                var v91 int32 = -2147483648
                                _ = v91
                            case 9223372036854775807:
                                _ = len(v3)
                            case 127:
                                switch fi87 {
                                case -391:
                                    for fi92 := 0; fi92 < 4; fi92++ {
                                        switch (fi87 ^ v18) {
                                        case 32767:
                                            if false {
                                                v6 -= 2
                                            } else {
                                                v5 = (true || (!true))
                                            }
                                        default:
                                            if false {
                                                switch ("" + (((("" + (v4 + "\n\t")) + v4) + v4) + "hello")) {
                                                case "\n\t":
                                                    v93 := GenS5{F0: 508, F1: v4}
                                                    _ = v93
                                                case "":
                                                    var v94 int16 = 1000
                                                    _ = v94
                                                default:
                                                    var v95 float32 = (0.0 - (float32((((v2 + ((3.14159 - ((v2 - v2) / v2)) - 0.1)) + (v2 * float64((((float32(-2.25) - 1.5) * 1e10) + (0.1 / (-2.25 - (3.14159 + (1e10 * 0.1)))))))) / 1e10)) + 3.14159))
                                                    _ = v95
                                                }
                                                v2 -= 1.5
                                            } else {
                                                var v96 string = v4
                                                var v97 int = 301
                                                v98 := GenS10{F0: v96}
                                                var v99 int = v18
                                                _ = v96
                                                _ = v97
                                                _ = v98
                                                _ = v99
                                            }
                                        }
                                        v1[749] = v69
                                        v100 := GenS11{F0: (v5 || v5), F1: v5, F2: 0}
                                        _ = fi92
                                        _ = v100
                                    }
                                default:
                                    v69 += 156
                                }
                            }
                            v101 := v1[(v69 & 2147483648)]
                            _ = v101
                        }
                    } else {
                        v69 -= (v69 - (^(v69)))
                        v2 += (((((v2 - float64(3.14159)) + v2) * (1e10 * v2)) / ((v2 - v2) * v2)) - (-2.25 * 1e10))
                        fmt.Println(fi87, fi87)
                    }
                case "foo bar":
                    if false {
                        var v102 float64 = v2
                        var v103 uint32 = 811
                        v104 := GenS12{F0: 127, F1: 0}
                        _ = v102
                        _ = v103
                        _ = v104
                    } else {
                        var v105 int16 = 1000
                        for fi106 := 0; fi106 < 11; fi106++ {
                            v2 += ((1.5 + float64(0.1)) * 1.5)
                            v18 -= (fi87 >> ((959) & 63))
                            _ = fi106
                        }
                        _ = v105
                    }
                case "a":
                    if (v5 || (!v5)) {
                        for fi107 := 0; fi107 < 14; fi107++ {
                            switch fi107 {
                            case 1:
                                v3 = append(v3, v69)
                            }
                            var v108 float64 = (0.1 * (v2 / 1e10))
                            if (255 < 42) {
                                v109 := GenS3{F0: (v6 + 146), F1: 42}
                                _ = v109
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if v5 {
                                    switch -821 {
                                    case 809:
                                        v4 = v4
                                    case 32767:
                                        var v110 uint16 = (v6 - v6)
                                        _ = v110
                                    default:
                                        fmt.Println(fi107)
                                    }
                                    v18 -= v18
                                    switch "a" {
                                    case "hello":
                                        if (271 <= -1) {
                                            v108 -= -2.25
                                        }
                                    }
                                } else {
                                    var v111 int8 = -128
                                    var v112 bool = true
                                    _ = v111
                                    _ = v112
                                }
                                v113 := GenS13{F0: -113, F1: int8(fi107), F2: "\n\t"}
                                _ = v113
                            }
                            _ = fi107
                            _ = v108
                        }
                        v114 := GenS14{F0: v5, F1: 0}
                        _ = v114
                    } else {
                        switch 32768 {
                        case 9223372036854775807:
                            v18 += v18
                        case -1:
                            v18 -= (v18 & fi87)
                        case 32768:
                            fmt.Println(v18, fi87, v69)
                        default:
                            if (!true) {
                                switch -844 {
                                case 1:
                                    var v115 uint32 = 623
                                    _ = v115
                                case -9223372036854775808:
                                    v4 = ("foo bar" + "foo bar")
                                case 949:
                                    var v116 int64 = 1
                                    _ = v116
                                default:
                                    fmt.Println(v4, v6)
                                }
                            } else {
                                v4 = "a"
                            }
                        }
                        for fi117 := 0; fi117 < 3; fi117++ {
                            fmt.Println(fi87)
                            if v5 {
                                var v118 int32 = 2147483647
                                v6 = (^(v6))
                                fmt.Println(fi87, fi87)
                                fmt.Println(v69, v69)
                                _ = v118
                            } else {
                                var v119 float64 = v2
                                _ = v119
                            }
                            if v5 {
                                var v120 bool = v5
                                switch ("foo bar" + v4) {
                                case "hello":
                                    if v120 {
                                        switch "a" {
                                        case "":
                                            switch v18 {
                                            case 0:
                                                v121 := GenS6{F0: 0}
                                                _ = v121
                                            case 4294967295:
                                                var v122 int64 = 9223372036854775807
                                                _ = v122
                                            case 32768:
                                                var v123 float64 = -2.25
                                                _ = v123
                                            default:
                                                var v124 uint = 42
                                                _ = v124
                                            }
                                        case "hello":
                                            if true {
                                                v69 = 494
                                                if v5 {
                                                    v125 := GenS15{F0: 870, F1: false}
                                                    v18 -= (fi117 & -9223372036854775808)
                                                    _ = v125
                                                }
                                                var v126 float32 = (1e10 / 1.0)
                                                _ = v126
                                            }
                                        case "a":
                                            var v127 int = fi87
                                            _ = v127
                                        }
                                        if (1000 != 1000) {
                                            v128 := GenS8{F0: 342, F1: 76, F2: v120}
                                            _ = v128
                                        } else {
                                            v5 = ((v4 == ("a" + (v4 + (("\n\t" + "") + v4)))) && (!(4294967295 != 4294967295)))
                                        }
                                    }
                                case "foo bar":
                                    v120 = v5
                                    fmt.Println(v4, v6, v4)
                                case "":
                                    v129 := v1[(v69 | v69)]
                                    _ = v129
                                default:
                                    var v130 uint8 = 135
                                    _ = v130
                                }
                                v131 := GenS16{F0: 42, F1: 2147483647, F2: int(v6)}
                                _ = v120
                                _ = v131
                            }
                            _ = fi117
                        }
                    }
                }
                _ = fi87
            }
        }
        if (607 <= 459) {
            var v132 uint64 = 878
            var v133 int8 = 78
            _ = v132
            _ = v133
        } else {
            if len(v3) > 0 {
                v3[0] = 1
            }
        }
    }
    v134 := GenS4{F0: 127, F1: (v6 << ((42) & 63))}
    for fi135 := 0; fi135 < 18; fi135++ {
        var v136 string = v4
        for fi137 := 0; fi137 < 14; fi137++ {
            v3 = append(v3, v69)
            _ = fi137
        }
        v138 := GenS3{F0: v6, F1: 127}
        _ = fi135
        _ = v136
        _ = v138
    }
    v18 = (-(v18))
    if false {
        if ((((1.5 * 1.5) * 0.0) / (0.0 + 1e10)) != 0.1) {
            if v5 {
                switch v18 {
                case 32767:
                    _ = len(v1)
                case 128:
                    if v5 {
                        v18 = (v18 << ((42) & 63))
                    } else {
                        var v139 float64 = v2
                        var v140 uint64 = 467
                        v139 = (v2 - 0.0)
                        _ = v139
                        _ = v140
                    }
                case -9223372036854775808:
                    _ = len(v3)
                }
                v6 = v6
            }
            if (!(1 != -1)) {
                var v141 uint = 255
                switch (v18 - 653) {
                case -1:
                    switch (v18 * (v18 << (((^(v141))) & 63))) {
                    case -1:
                        var v142 float64 = 0.0
                        _ = v142
                    default:
                        if (v5 || ((^(v6)) != v6)) {
                            v18 = v18
                            v4 = v4
                            v143 := GenS7{F0: (2147483647 > -2147483648), F1: v6}
                            _ = v143
                        }
                    }
                case 128:
                    var v144 string = (((v4 + (v4 + ("\n\t" + v4))) + v4) + v4)
                    _ = v144
                case 4294967295:
                    if (!(0.0 == 1.5)) {
                        v145 := GenS17{F0: 32767, F1: 2}
                        var v146 uint32 = (v69 * v69)
                        _ = v145
                        _ = v146
                    } else {
                        v147 := GenS18{F0: -1}
                        v18 = (v18 ^ (v18 | 0))
                        _ = v147
                    }
                }
                v6 = (^(v6))
                _ = v141
            }
        } else {
            if (v6 == 42) {
                if v5 {
                    if (v5 || true) {
                        switch "\n\t" {
                        case "foo bar":
                            switch (v18 + (v18 % (((v18 ^ v18)) | 1))) {
                            case 2147483647:
                                var v148 int = v18
                                _ = v148
                            case 0:
                                v2 -= v2
                            default:
                                switch v18 {
                                case 127:
                                    fmt.Println(v2, v2)
                                }
                            }
                        case "hello":
                            for fi149 := 0; fi149 < 15; fi149++ {
                                v150 := GenS11{F0: ((v69 ^ 2147483647) > 65535), F1: v5, F2: 535}
                                v151 := GenS16{F0: 0, F1: -983, F2: (-(fi149))}
                                if (0 > 1) {
                                    if (618 >= fi149) {
                                        v152 := GenS19{F0: 2147483647}
                                        v5 = v5
                                        _ = v152
                                    }
                                    v18 = (fi149 ^ 4294967295)
                                    var v153 float64 = v2
                                    _ = v153
                                } else {
                                    for fi154 := 0; fi154 < 9; fi154++ {
                                        var v155 int64 = -9223372036854775808
                                        _ = fi154
                                        _ = v155
                                    }
                                    v69 -= (v69 * v69)
                                }
                                _ = fi149
                                _ = v150
                                _ = v151
                            }
                        default:
                            for fi156 := 0; fi156 < 20; fi156++ {
                                switch (v4 + v4) {
                                case "hello":
                                    var v157 uint64 = 18446744073709551615
                                    _ = v157
                                case "foo bar":
                                    if v5 {
                                        if v5 {
                                            v158 := GenS20{F0: 1, F1: -494, F2: 4294967295}
                                            _ = v158
                                        }
                                        for fi159 := 0; fi159 < 11; fi159++ {
                                            var v160 int8 = 127
                                            fmt.Println(fi156, fi159, fi159)
                                            switch (v4 + v4) {
                                            case "\n\t":
                                                if v5 {
                                                    var v161 uint32 = 2
                                                    var v162 int32 = -977
                                                    _ = v161
                                                    _ = v162
                                                } else {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                            case "hello":
                                                fmt.Println(fi159)
                                            case "":
                                                var v163 int64 = 1
                                                _ = v163
                                            }
                                            _ = fi159
                                            _ = v160
                                        }
                                    }
                                case "\n\t":
                                    var v164 int64 = -98
                                    _ = v164
                                }
                                _ = fi156
                            }
                        }
                        v165 := GenS9{F0: 935, F1: 42, F2: 255}
                        _ = v165
                    } else {
                        v18 += v18
                    }
                } else {
                    v2 = v2
                    v166 := GenS16{F0: 1, F1: 2147483647, F2: (^(v18))}
                    _ = v166
                }
                v6 -= v6
                v167 := GenS19{F0: (^(v18))}
                _ = v167
            }
            v168 := v1[65535]
            for fi169 := 0; fi169 < 15; fi169++ {
                switch v4 {
                case "a":
                    switch -728 {
                    case -1:
                        var v170 string = "hello"
                        _ = v170
                    }
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                switch (fi169 % ((fi169) | 1)) {
                case 240:
                    var v171 bool = v5
                    _ = v171
                default:
                    v168 += 0
                }
                _ = fi169
            }
            _ = v168
        }
        var v172 uint8 = 79
        for fi173 := 0; fi173 < 11; fi173++ {
            switch v18 {
            case 119:
                var v174 bool = ((v172 << ((4294967295) & 63)) > (v172 / (((^(v172))) | 1)))
                _ = v174
            case -9223372036854775808:
                if (!v5) {
                    v172 += v172
                }
            default:
                v3 = append(v3, v69)
            }
            _ = fi173
        }
        _ = v172
    }
    for fi175 := 0; fi175 < 14; fi175++ {
        if v5 {
            v176 := GenS20{F0: v6, F1: 1, F2: 463}
            _ = v176
        } else {
            v177 := GenS21{F0: (v6 & v6)}
            v178 := GenS22{F0: v69, F1: 0}
            var v179 uint64 = 1
            _ = v177
            _ = v178
            _ = v179
        }
        for fi180 := 0; fi180 < 12; fi180++ {
            if v5 {
                if v5 {
                    v181 := GenS23{F0: (v4 + ("hello" + "a")), F1: v4}
                    switch (fi175 * 1) {
                    case 2147483647:
                        if (((-2.25 + float32((((v2 + ((0.1 + v2) + v2)) + (3.14159 / float64(0.1))) / float64(-2.25)))) / 0.1) > 1e10) {
                            v182 := GenS24{F0: 0, F1: 42, F2: 127}
                            switch v4 {
                            case "foo bar":
                                delete(v1, v69)
                            case "\n\t":
                                var v183 int16 = 41
                                _ = v183
                            case "":
                                if true {
                                    v1[v69] = 858
                                    switch (v4 + (v4 + "foo bar")) {
                                    case "a":
                                        var v184 float64 = v2
                                        _ = v184
                                    }
                                    v185 := GenS12{F0: 0, F1: v69}
                                    _ = v185
                                }
                            default:
                                v186 := GenS25{F0: -33}
                                _ = v186
                            }
                            var v187 float32 = (3.14159 + 1.5)
                            _ = v182
                            _ = v187
                        }
                    case 0:
                        var v188 uint8 = 126
                        _ = v188
                    }
                    _ = v181
                }
                if true {
                    var v189 bool = v5
                    v69 -= v69
                    if v189 {
                        var v190 uint32 = (^(v69))
                        v5 = v5
                        switch (v4 + v4) {
                        case "a":
                            var v191 int64 = 9223372036854775807
                            _ = v191
                        case "\n\t":
                            if v5 {
                                v192 := GenS15{F0: 1, F1: false}
                                _ = v192
                            }
                        default:
                            v193 := GenS26{F0: 218}
                            _ = v193
                        }
                        _ = v190
                    } else {
                        switch "" {
                        case "foo bar":
                            fmt.Println(v18, v6, v4)
                        case "":
                            v4 = v4
                        case "hello":
                            fmt.Println(v18, fi180)
                        default:
                            if true {
                                v2 = ((1.5 * v2) - v2)
                                v194 := GenS27{F0: v6, F1: 42}
                                _ = v194
                            }
                        }
                        delete(v1, 101)
                        if ((0.0 + 0.0) != 1.5) {
                            var v195 int8 = 29
                            switch fi175 {
                            case 42:
                                switch (fi175 + v18) {
                                case 4294967295:
                                    if v189 {
                                        v196 := GenS2{F0: v4, F1: (v69 & v69)}
                                        _ = v196
                                    } else {
                                        v6 = (v6 ^ v6)
                                        v2 = v2
                                        fmt.Println(v18, fi175)
                                        var v197 uint32 = v69
                                        if len(v3) > 0 {
                                            v3[0] = 42
                                        }
                                        var v198 uint8 = 120
                                        _ = v197
                                        _ = v198
                                    }
                                case 0:
                                    v2 -= (0.1 / v2)
                                }
                            case 9223372036854775807:
                                _ = len(v1)
                            case 4294967295:
                                fmt.Println(v5, fi180)
                            default:
                                _ = len(v3)
                            }
                            v69 -= v69
                            _ = v195
                        }
                    }
                    fmt.Println(v4)
                    switch (fi180 + -1) {
                    case 32767:
                        var v199 uint16 = (v6 + (^(v6)))
                        _ = v199
                    case -548:
                        fmt.Println(v69, v189)
                    }
                    _ = v189
                }
            } else {
                v200 := GenS26{F0: 1}
                fmt.Println(fi180)
                switch v18 {
                case -88:
                    v6 = 0
                case 127:
                    v4 = ""
                default:
                    v3 = append(v3, 307)
                }
                _ = v200
            }
            var v201 int32 = 292
            _ = fi180
            _ = v201
        }
        var v202 float32 = (float32(float64(0.0)) - 1e10)
        _ = fi175
        _ = v202
    }
    v203 := GenS28{F0: -1, F1: "\n\t"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v134
    _ = v18
    _ = v2
    _ = v203
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
}
