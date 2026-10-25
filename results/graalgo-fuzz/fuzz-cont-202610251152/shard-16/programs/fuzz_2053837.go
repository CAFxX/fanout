package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 int16; F2 int16 }
type GenS3 struct { F0 int8; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 int64; F1 uint64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int32; F1 int32 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint16; F1 int32; F2 int32 }
type GenS12 struct { F0 int64; F1 bool }
func genFn0(p0 bool, p1 int8) uint {
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 int64 = 0
        _ = fi6
        _ = v7
    }
    p1 -= p1
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 int) uint32 {
    v8 := []int{(p0 | p0), (-(p0))}
    var v9 uint16 = 529
    _ = p0
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := []string{("\n\t" + ""), "a"}
    var v2 int = -1
    v3 := []string{"foo bar", ""}
    v4 := make(map[uint32]uint32)
    var v5 bool = true
    delete(v4, 50)
    var v10 uint16 = 0
    fmt.Println(v5, v10, v2)
    fmt.Println(v2)
    v11 := GenS1{F0: v10}
    var v12 int8 = -81
    v3 = append(v3, "hello")
    v12 -= 100
    if v5 {
        fmt.Println(v5, v5)
        v13 := GenS1{F0: v10}
        _ = v13
    } else {
        var v14 int = 605
        if (-565 < 42) {
            for fi15 := 0; fi15 < 13; fi15++ {
                if v5 {
                    v10 = (v10 ^ 1)
                }
                switch (v2 * (v14 & (v14 - -2147483648))) {
                case -2147483648:
                    var v16 float64 = float64(0.1)
                    _ = v16
                }
                _ = fi15
            }
            if (127 <= 2) {
                v17 := GenS2{F0: 4294967295, F1: 274, F2: 1000}
                for fi18 := 0; fi18 < 12; fi18++ {
                    v14 = 42
                    _ = fi18
                }
                fmt.Println(v10)
                _ = v17
            }
        }
        _ = v14
    }
    if (!("foo bar" <= "foo bar")) {
        if true {
            for fi19 := 0; fi19 < 7; fi19++ {
                v10 += (^(v10))
                for fi20 := 0; fi20 < 5; fi20++ {
                    fmt.Println(v2, v5)
                    var v21 uint = 4294967295
                    var v22 string = (("foo bar" + "foo bar") + "foo bar")
                    var v23 uint8 = 255
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
                v10 -= (v10 | v10)
                _ = fi19
            }
            for fi24 := 0; fi24 < 19; fi24++ {
                fmt.Println(v2)
                var v25 uint64 = 2
                _ = fi24
                _ = v25
            }
            if false {
                var v26 int64 = 9223372036854775807
                _ = v26
            }
        } else {
            for fi27 := 0; fi27 < 1; fi27++ {
                for fi28 := 0; fi28 < 6; fi28++ {
                    if v5 {
                        if v5 {
                            var v29 float32 = -2.25
                            v4[110] = 42
                            var v30 string = "hello"
                            _ = v29
                            _ = v30
                        }
                    }
                    v12 = (v12 << ((0) & 63))
                    v12 -= 122
                    _ = fi28
                }
                if true {
                    if ((((v10 != v10) && v5) || false) || (2 > 4294967295)) {
                        var v31 uint64 = 932
                        _ = v31
                    } else {
                        if (1000 != 32767) {
                            fmt.Println(v10, v10, v10)
                        }
                    }
                    _ = len(v3)
                } else {
                    var v32 uint = 1
                    v10 = v10
                    if false {
                        switch v2 {
                        case 4294967295:
                            var v33 uint32 = 207
                            _ = v33
                        default:
                            v12 += (^(v12))
                        }
                        switch fi27 {
                        case 32767:
                            switch v2 {
                            case -2147483648:
                                v2 += (v2 << ((2) & 63))
                            case 9223372036854775807:
                                fmt.Println(v10, v10, v32)
                            case 0:
                                v34 := GenS3{F0: (v12 - 96), F1: 2147483647, F2: 1}
                                _ = v34
                            }
                        case -627:
                            for fi35 := 0; fi35 < 11; fi35++ {
                                var v36 uint64 = 42
                                fmt.Println(v36)
                                v4[0] = 42
                                if len(v3) > 0 {
                                    v3[0] = "\n\t"
                                }
                                v37 := GenS1{F0: (v10 ^ v10)}
                                if ("a" == "") {
                                    if ((42 == -335) || (v5 && v5)) {
                                        v10 = (v10 % (((v10 ^ (v10 ^ v10))) | 1))
                                    }
                                } else {
                                    v36 = (v36 / ((v36) | 1))
                                    v5 = ((!((1.5 >= 0.1) || (v12 != 116))) && v5)
                                }
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                        }
                        v38 := GenS3{F0: (v12 % (((-(v12))) | 1)), F1: 1, F2: 18446744073709551615}
                        _ = v38
                    } else {
                        var v39 float32 = -2.25
                        var v40 uint32 = 2
                        _ = v39
                        _ = v40
                    }
                    _ = v32
                }
                v2 += 4294967295
                _ = fi27
            }
        }
        for fi41 := 0; fi41 < 9; fi41++ {
            for fi42 := 0; fi42 < 7; fi42++ {
                switch ("hello" + (("a" + ("foo bar" + "")) + ("" + ""))) {
                case "foo bar":
                    switch "" {
                    case "a":
                        if true {
                            v2 = 4294967295
                            v43 := v4[42]
                            _ = v43
                        } else {
                            if (9223372036854775807 <= 353) {
                                v3 = append(v3, "")
                                if ((!false) && v5) {
                                    var v44 int32 = -1
                                    v3 = append(v3, "a")
                                    switch (fi42 % ((fi41) | 1)) {
                                    case -883:
                                        v44 += int32(v12)
                                    case 32768:
                                        switch "hello" {
                                        case "foo bar":
                                            var v45 float32 = 1e10
                                            _ = v45
                                        }
                                    }
                                    _ = v44
                                } else {
                                    var v46 int = fi41
                                    fmt.Println(fi42, v5)
                                    fmt.Println(v2, v10)
                                    if len(v1) > 0 {
                                        v1[0] = "\n\t"
                                    }
                                    switch "hello" {
                                    case "hello":
                                        var v47 uint16 = v10
                                        _ = v47
                                    case "a":
                                        v5 = true
                                    case "\n\t":
                                        var v48 uint16 = v10
                                        fmt.Println(fi42)
                                        v48 = v10
                                        _ = v48
                                    }
                                    _ = v46
                                }
                                var v49 uint8 = 2
                                _ = v49
                            }
                            v2 = 128
                            fmt.Println(v2, v10)
                        }
                    case "foo bar":
                        if len(v3) > 0 {
                            v3[0] = ("\n\t" + (("foo bar" + "hello") + ("foo bar" + "\n\t")))
                        }
                    }
                default:
                    if (42 <= 42) {
                        var v50 uint64 = 42
                        var v51 float32 = (1e10 - 3.14159)
                        if len(v3) > 0 {
                            v3[0] = ("\n\t" + (("foo bar" + "") + "foo bar"))
                        }
                        _ = v50
                        _ = v51
                    } else {
                        var v52 float32 = 3.14159
                        _ = v52
                    }
                }
                fmt.Println(v2)
                var v53 uint16 = (v10 >> ((1) & 63))
                _ = fi42
                _ = v53
            }
            _ = fi41
        }
        if false {
            v54 := GenS3{F0: v12, F1: 0, F2: 1}
            v5 = false
            if false {
                v55 := GenS3{F0: (v12 + v12), F1: 1, F2: 1}
                var v56 uint64 = 0
                _ = v55
                _ = v56
            } else {
                var v57 bool = v5
                var v58 int = 128
                _ = v57
                _ = v58
            }
            _ = v54
        }
    }
    switch 1 {
    case -593:
        var v59 uint8 = 47
        _ = v59
    case 42:
        v60 := GenS4{F0: 65535, F1: v12}
        _ = v60
    default:
        _ = len(v1)
    }
    v12 -= v12
    var v61 int16 = 526
    var v62 string = "\n\t"
    for fi63 := 0; fi63 < 3; fi63++ {
        var v64 uint8 = 207
        for fi65 := 0; fi65 < 7; fi65++ {
            v66 := GenS3{F0: (v12 - 100), F1: 255, F2: 803}
            v1 = append(v1, ("\n\t" + v62))
            _ = fi65
            _ = v66
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi63
        _ = v64
    }
    var v67 float64 = float64(-2.25)
    if (1 >= 2) {
        var v68 string = (v62 + ("hello" + ((v62 + (v62 + v62)) + v62)))
        _ = v68
    } else {
        if v5 {
            var v69 float32 = (float32(0.1) - 0.1)
            switch v2 {
            case 383:
                switch v62 {
                case "a":
                    if (v2 != int(v10)) {
                        if true {
                            var v70 int8 = v12
                            var v71 string = v62
                            v72 := GenS5{F0: 4294967295}
                            _ = v70
                            _ = v71
                            _ = v72
                        } else {
                            if ((v2 / (((v2 ^ 1)) | 1)) == v2) {
                                for fi73 := 0; fi73 < 12; fi73++ {
                                    for fi74 := 0; fi74 < 12; fi74++ {
                                        if v5 {
                                            var v75 uint16 = (v10 * 118)
                                            _ = v75
                                        } else {
                                            var v76 int64 = -13
                                            v77 := GenS6{F0: 65535, F1: v62}
                                            if (v69 != (v69 * v69)) {
                                                v4[2] = 287
                                                switch (fi73 ^ (fi73 << ((42) & 63))) {
                                                case 991:
                                                    var v78 uint = 42
                                                    if false {
                                                        _ = len(v4)
                                                    } else {
                                                        fmt.Println(v10)
                                                        v5 = (!false)
                                                    }
                                                    v69 = (v69 / (float32(v67) + v69))
                                                    _ = v78
                                                case -524:
                                                    v2 -= fi74
                                                case -2147483648:
                                                    var v79 uint32 = 4294967295
                                                    _ = v79
                                                }
                                            } else {
                                                switch "" {
                                                case "a":
                                                    if len(v3) > 0 {
                                                        v3[0] = (v62 + ("a" + ""))
                                                    }
                                                case "hello":
                                                    if (42 != (^(v12))) {
                                                        var v80 uint16 = (v10 | v10)
                                                        var v81 int = v2
                                                        delete(v4, 2147483648)
                                                        v4[genFn1((fi74 + int(v76)))] = 1
                                                        v1 = append(v1, (v62 + (("foo bar" + ("foo bar" + v62)) + v62)))
                                                        var v82 int8 = (-(v12))
                                                        var v83 bool = (42 != 4294967295)
                                                        _ = v80
                                                        _ = v81
                                                        _ = v82
                                                        _ = v83
                                                    }
                                                }
                                            }
                                            _ = v76
                                            _ = v77
                                        }
                                        _ = fi74
                                    }
                                    for fi84 := 0; fi84 < 18; fi84++ {
                                        v85 := GenS2{F0: 4294967295, F1: (v61 % ((v61) | 1)), F2: -631}
                                        if true {
                                            var v86 int = (v2 / (((fi84 * 32768)) | 1))
                                            if false {
                                                var v87 uint16 = v10
                                                if v5 {
                                                    var v88 uint = 394
                                                    if false {
                                                        var v89 int32 = 0
                                                        v88 = v88
                                                        switch (fi84 % (((v2 * 615)) | 1)) {
                                                        case 0:
                                                            fmt.Println(v69)
                                                        case 127:
                                                            var v90 float32 = v69
                                                            _ = v90
                                                        case 978:
                                                            if (!false) {
                                                                v10 -= (v87 + (v87 >> (((v88 % ((v88) | 1))) & 63)))
                                                                var v91 int32 = (v89 & v89)
                                                                if (!v5) {
                                                                    v62 = v62
                                                                }
                                                                v5 = (v5 && v5)
                                                                var v92 uint8 = 2
                                                                _ = v91
                                                                _ = v92
                                                            } else {
                                                                if (uint8(v87) >= 2) {
                                                                    v87 = uint16(v88)
                                                                    v69 = (((v69 + ((1e10 * (v69 * ((float32(v67) - ((3.14159 / v69) + v69)) / float32(v67)))) + v69)) * (((v69 + v69) / v69) / 1e10)) + (-2.25 + v69))
                                                                    v93 := GenS7{F0: 42, F1: 1}
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        v94 := GenS8{F0: v87}
                                                                        var v95 float32 = 1.5
                                                                        _ = v94
                                                                        _ = v95
                                                                    }
                                                                    _ = v93
                                                                } else {
                                                                    if (!false) {
                                                                        var v96 string = v62
                                                                        v69 -= (v69 * v69)
                                                                        _ = v96
                                                                    } else {
                                                                        v3 = append(v3, ((v62 + (("a" + ("hello" + v62)) + v62)) + v62))
                                                                        var v97 int32 = (^(v89))
                                                                        if len(v3) > 0 {
                                                                            v3[0] = ("foo bar" + v62)
                                                                        }
                                                                        switch v62 {
                                                                        case "hello":
                                                                            if v5 {
                                                                                switch fi84 {
                                                                                case -600:
                                                                                    v5 = (("a" + "\n\t") >= v62)
                                                                                case -532:
                                                                                    fmt.Println(v67, v61)
                                                                                case 128:
                                                                                    var v98 int64 = int64(v61)
                                                                                    fmt.Println(v61, v67, v89)
                                                                                    _ = len(v4)
                                                                                    var v99 int16 = (-(v61))
                                                                                    _ = v98
                                                                                    _ = v99
                                                                                default:
                                                                                    v100 := GenS7{F0: 0, F1: uint64(v87)}
                                                                                    _ = v100
                                                                                }
                                                                                v1 = append(v1, v62)
                                                                            }
                                                                        case "a":
                                                                            var v101 bool = (!false)
                                                                            _ = v101
                                                                        }
                                                                        _ = v97
                                                                    }
                                                                    v69 -= ((float32(v67) + v69) + v69)
                                                                    if false {
                                                                        if false {
                                                                            switch v62 {
                                                                            case "":
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            case "\n\t":
                                                                                if (-9223372036854775808 <= 42) {
                                                                                    var v102 float32 = v69
                                                                                    if ((!v5) || v5) {
                                                                                        v103 := GenS9{F0: (v89 << ((464) & 63)), F1: v89}
                                                                                        var v104 int16 = -52
                                                                                        _ = v103
                                                                                        _ = v104
                                                                                    }
                                                                                    _ = v102
                                                                                } else {
                                                                                    v10 += v10
                                                                                }
                                                                            case "a":
                                                                                fmt.Println(v5, v5, v69)
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                v86 = (fi84 | (^(v86)))
                                                            }
                                                        default:
                                                            switch v62 {
                                                            case "foo bar":
                                                                v105 := GenS4{F0: 2147483648, F1: 1}
                                                                _ = v105
                                                            case "hello":
                                                                var v106 float64 = 0.0
                                                                _ = v106
                                                            case "":
                                                                var v107 uint32 = 616
                                                                _ = v107
                                                            }
                                                        }
                                                        _ = v89
                                                    }
                                                    _ = v88
                                                } else {
                                                    var v108 uint64 = 1
                                                    v108 = (v108 % ((v108) | 1))
                                                    v109 := v4[516]
                                                    var v110 uint32 = (^(v109))
                                                    var v111 string = "a"
                                                    _ = v108
                                                    _ = v109
                                                    _ = v110
                                                    _ = v111
                                                }
                                                _ = v87
                                            }
                                            _ = v86
                                        }
                                        v112 := GenS10{F0: (v10 < (^(v10)))}
                                        _ = fi84
                                        _ = v112
                                        _ = v85
                                    }
                                    v3 = append(v3, (v62 + v62))
                                    _ = fi73
                                }
                            }
                        }
                        if (v12 <= v12) {
                            switch "\n\t" {
                            case "\n\t":
                                v10 = v10
                            default:
                                var v113 float64 = v67
                                _ = v113
                            }
                        } else {
                            var v114 int64 = -9223372036854775808
                            if ((!(!true)) && v5) {
                                if v5 {
                                    for fi115 := 0; fi115 < 13; fi115++ {
                                        var v116 bool = (v5 && true)
                                        if (!v116) {
                                            v10 += 255
                                            fmt.Println(v114, v5)
                                        }
                                        _ = fi115
                                        _ = v116
                                    }
                                } else {
                                    switch v2 {
                                    case 42:
                                        for fi117 := 0; fi117 < 11; fi117++ {
                                            if false {
                                                switch fi117 {
                                                case -2147483648:
                                                    v69 += 1.5
                                                case 961:
                                                    switch "hello" {
                                                    case "\n\t":
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    case "hello":
                                                        switch (v62 + (v62 + v62)) {
                                                        case "\n\t":
                                                            if ((646 != 558) || false) {
                                                                var v118 uint = 65535
                                                                switch -941 {
                                                                case -1:
                                                                    var v119 float64 = v67
                                                                    _ = v119
                                                                case 2147483647:
                                                                    v12 = 0
                                                                default:
                                                                    fmt.Println(v69, v62)
                                                                }
                                                                _ = v118
                                                            } else {
                                                                var v120 uint8 = 164
                                                                v121 := v4[65535]
                                                                v114 = 1
                                                                _ = v120
                                                                _ = v121
                                                            }
                                                        default:
                                                            v122 := GenS3{F0: -128, F1: -2147483648, F2: 903}
                                                            _ = v122
                                                        }
                                                    default:
                                                        v61 += (v61 >> ((42) & 63))
                                                    }
                                                case 128:
                                                    var v123 uint8 = 0
                                                    _ = v123
                                                default:
                                                    var v124 int8 = int8(v10)
                                                    _ = v124
                                                }
                                            } else {
                                                var v125 string = ((("a" + "hello") + (v62 + "")) + v62)
                                                delete(v4, 4294967295)
                                                _ = v125
                                            }
                                            var v126 uint32 = 839
                                            _ = fi117
                                            _ = v126
                                        }
                                    default:
                                        v127 := GenS11{F0: (v10 * 0), F1: 44, F2: 1}
                                        _ = v127
                                    }
                                }
                                v128 := GenS7{F0: 0, F1: 1}
                                _ = v128
                            }
                            if (v5 && (-1 != (^(v12)))) {
                                v4[248] = 11
                                for fi129 := 0; fi129 < 11; fi129++ {
                                    v62 = ("\n\t" + v62)
                                    v3 = append(v3, v62)
                                    _ = fi129
                                }
                                v130 := v4[42]
                                _ = v130
                            } else {
                                var v131 uint16 = 805
                                var v132 int8 = int8(v2)
                                _ = v131
                                _ = v132
                            }
                            _ = v114
                        }
                        v5 = (-1 != (v61 / (((v61 - 32767)) | 1)))
                    } else {
                        var v133 string = "hello"
                        fmt.Println(v5, v69, v2)
                        var v134 float32 = v69
                        _ = v133
                        _ = v134
                    }
                case "foo bar":
                    v61 = (v61 % ((1) | 1))
                default:
                    v135 := GenS12{F0: -758, F1: (v2 <= v2)}
                    _ = v135
                }
            case 4294967295:
                v69 += float32((float64((0.1 / -2.25)) * (-2.25 + (3.14159 + v67))))
            }
            _ = v69
        }
    }
    if true {
        for fi136 := 0; fi136 < 16; fi136++ {
            v61 -= v61
            _ = fi136
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v137 int64 = -868
        _ = v137
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v61
    _ = v62
    _ = v67
}
