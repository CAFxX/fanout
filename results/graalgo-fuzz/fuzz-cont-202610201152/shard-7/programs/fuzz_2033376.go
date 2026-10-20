package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS4 struct { F0 uint16; F1 int; F2 int }
type GenS5 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS6 struct { F0 uint32; F1 uint16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS9 struct { F0 int8; F1 int32 }
type GenS10 struct { F0 uint16; F1 int32; F2 uint }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int32; F1 int32; F2 uint8 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 uint; F1 int32; F2 int64 }
type GenS15 struct { F0 uint64; F1 uint16 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 uint16; F1 int32; F2 int }
type GenS18 struct { F0 uint32 }
func genFn0(p0 uint8) int {
    v5 := GenS1{F0: 669}
    for fi6 := 0; fi6 < 2; fi6++ {
        if (2 < 255) {
            v7 := GenS2{F0: fi6}
            var v8 bool = true
            var v9 int32 = -2147483648
            _ = v7
            _ = v8
            _ = v9
        }
        var v10 float32 = 1.5
        _ = fi6
        _ = v10
    }
    p0 = p0
    _ = p0
    _ = v5
    return -857
}
func genFn1(p0 int64) uint {
    v11 := GenS3{F0: 42, F1: 0, F2: -2147483648}
    v12 := []int{4294967295, 0}
    _ = p0
    _ = v11
    _ = v12
    return 0
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint64 = (v1 >> ((18446744073709551615) & 63))
    v3 := []string{"foo bar"}
    var v4 string = (("\n\t" + ("\n\t" + "foo bar")) + "\n\t")
    if (false && (-619 < -776)) {
        v4 = (v4 + "a")
        if (-28 == -73) {
            switch genFn0(2) {
            case 128:
                for fi13 := 0; fi13 < 2; fi13++ {
                    v14 := GenS4{F0: 844, F1: 32768, F2: 432}
                    for fi15 := 0; fi15 < 8; fi15++ {
                        v2 = (^(v2))
                        _ = fi15
                    }
                    var v16 int16 = -111
                    _ = fi13
                    _ = v14
                    _ = v16
                }
            case -2147483648:
                var v17 int = 42
                _ = v17
            case 1:
                if false {
                    for fi18 := 0; fi18 < 7; fi18++ {
                        switch (fi18 + (fi18 - fi18)) {
                        case -2147483648:
                            if false {
                                switch v4 {
                                case "a":
                                    v19 := GenS1{F0: 32767}
                                    _ = v19
                                default:
                                    v20 := GenS3{F0: 1, F1: 2, F2: -1}
                                    _ = v20
                                }
                                _ = len(v3)
                            }
                        case 4294967295:
                            v2 -= v2
                        case 32768:
                            fmt.Println(v1, v4)
                        }
                        _ = fi18
                    }
                    for fi21 := 0; fi21 < 5; fi21++ {
                        var v22 string = "\n\t"
                        switch "hello" {
                        case "\n\t":
                            if false {
                                if (!false) {
                                    var v23 uint32 = uint32(v2)
                                    var v24 float32 = 0.1
                                    _ = v23
                                    _ = v24
                                } else {
                                    if false {
                                        fmt.Println(v2, fi21)
                                        switch -2147483648 {
                                        case 32768:
                                            var v25 int16 = 431
                                            _ = v25
                                        case -74:
                                            v22 = ("" + ("\n\t" + v4))
                                        case 32767:
                                            v26 := GenS5{F0: 0, F1: -891, F2: v2}
                                            _ = v26
                                        default:
                                            var v27 int8 = 11
                                            _ = v27
                                        }
                                        v1 += v2
                                    }
                                    _ = len(v3)
                                }
                            }
                        default:
                            switch genFn0(2) {
                            case 57:
                                if (!((0.0 - (((0.0 / 0.1) * ((-2.25 - (1.5 - float32(1.5))) + (0.1 * ((-2.25 + -2.25) + 1.5)))) / 1.0)) == -2.25)) {
                                    if ((fi21 * 32767) != (fi21 & -9223372036854775808)) {
                                        if len(v3) > 0 {
                                            v3[0] = v22
                                        }
                                        v28 := GenS6{F0: 173, F1: 32767}
                                        var v29 uint16 = 42
                                        _ = v28
                                        _ = v29
                                    }
                                    for fi30 := 0; fi30 < 3; fi30++ {
                                        v3 = append(v3, v22)
                                        if true {
                                            switch ("\n\t" + "") {
                                            case "":
                                                var v31 uint = 100
                                                _ = v31
                                            case "hello":
                                                v2 = v1
                                                v3 = append(v3, v22)
                                                _ = len(v3)
                                            case "a":
                                                var v32 int = fi30
                                                _ = v32
                                            default:
                                                var v33 float64 = float64(1e10)
                                                _ = v33
                                            }
                                            fmt.Println(v2)
                                        } else {
                                            _ = len(v3)
                                            if false {
                                                v22 = "a"
                                                var v34 int16 = 32767
                                                if (!false) {
                                                    fmt.Println(v34)
                                                } else {
                                                    v35 := GenS7{F0: int(v2)}
                                                    v36 := GenS8{F0: 0, F1: 42, F2: (v2 - (v1 + (v2 & v1)))}
                                                    switch "" {
                                                    case "foo bar":
                                                        v1 -= v1
                                                    default:
                                                        if true {
                                                            if len(v3) > 0 {
                                                                v3[0] = v22
                                                            }
                                                            v4 = v4
                                                            if false {
                                                                v37 := GenS9{F0: 127, F1: -714}
                                                                v1 = v2
                                                                switch (fi21 & fi21) {
                                                                case 2147483647:
                                                                    var v38 int = -1
                                                                    _ = v38
                                                                }
                                                                _ = v37
                                                            } else {
                                                                if true {
                                                                    if (!(0.1 != -2.25)) {
                                                                        _ = len(v3)
                                                                        if true {
                                                                            v39 := GenS5{F0: 42, F1: v34, F2: (v2 | 2)}
                                                                            _ = v39
                                                                        } else {
                                                                            switch 0 {
                                                                            case 42:
                                                                                v40 := GenS10{F0: 2, F1: -2147483648, F2: 357}
                                                                                _ = v40
                                                                            case 796:
                                                                                var v41 uint = 1
                                                                                _ = v41
                                                                            case -507:
                                                                                v22 = ("foo bar" + "foo bar")
                                                                            }
                                                                        }
                                                                        fmt.Println(v2)
                                                                        v2 -= (v1 << ((1) & 63))
                                                                        var v42 int64 = 42
                                                                        fmt.Println(v42, fi21, v1)
                                                                        _ = v42
                                                                    } else {
                                                                        v3 = append(v3, v4)
                                                                        v3 = append(v3, ((("\n\t" + "") + "hello") + v4))
                                                                    }
                                                                    switch v4 {
                                                                    case "\n\t":
                                                                        v22 = ("foo bar" + v4)
                                                                    case "":
                                                                        if false {
                                                                            switch genFn0(2) {
                                                                            case 9223372036854775807:
                                                                                if (true || false) {
                                                                                    v43 := GenS8{F0: -2147483648, F1: 32767, F2: (v2 ^ (^(v2)))}
                                                                                    v2 += (v2 * (v2 - v1))
                                                                                    fmt.Println(v2, fi30)
                                                                                    if ((0 != uint32(v1)) || true) {
                                                                                        v34 = int16(fi21)
                                                                                        v34 -= (v34 | v34)
                                                                                    } else {
                                                                                        fmt.Println(v1)
                                                                                    }
                                                                                    _ = v43
                                                                                }
                                                                            case -9223372036854775808:
                                                                                var v44 int8 = 11
                                                                                _ = v44
                                                                            }
                                                                        }
                                                                    case "a":
                                                                        v45 := GenS11{F0: 65535}
                                                                        _ = v45
                                                                    }
                                                                } else {
                                                                    fmt.Println(v2, v4, fi21)
                                                                }
                                                                v2 = v2
                                                                if false {
                                                                    var v46 float64 = float64(0.0)
                                                                    _ = v46
                                                                } else {
                                                                    var v47 uint = 229
                                                                    if (fi30 < (fi21 + (fi30 % ((0) | 1)))) {
                                                                        if false {
                                                                            v1 += 1
                                                                            v34 -= v34
                                                                            v22 = v4
                                                                        }
                                                                        if (42 < uint32(v47)) {
                                                                            fmt.Println(v4, fi30, v34)
                                                                            var v48 int8 = -128
                                                                            switch (fi21 + int(v48)) {
                                                                            case 0:
                                                                                if false {
                                                                                    _ = len(v3)
                                                                                }
                                                                                v48 -= -1
                                                                                var v49 float64 = float64((0.0 - -2.25))
                                                                                _ = v49
                                                                            case 186:
                                                                                v48 -= 42
                                                                            default:
                                                                                var v50 int8 = v48
                                                                                _ = v50
                                                                            }
                                                                            _ = v48
                                                                        }
                                                                    }
                                                                    var v51 bool = true
                                                                    var v52 bool = v51
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        v3 = append(v3, (("a" + "hello") + "a"))
                                                                    case "\n\t":
                                                                        v53 := GenS12{F0: 42, F1: -1, F2: 254}
                                                                        _ = v53
                                                                    }
                                                                    if ("\n\t" == "a") {
                                                                        v54 := GenS7{F0: (-(fi30))}
                                                                        var v55 float32 = 1e10
                                                                        var v56 string = "\n\t"
                                                                        _ = v54
                                                                        _ = v55
                                                                        _ = v56
                                                                    } else {
                                                                        switch (((v22 + v4) + "a") + (((("" + "a") + "foo bar") + (v4 + (v22 + "a"))) + "hello")) {
                                                                        case "hello":
                                                                            v57 := GenS5{F0: 42, F1: (v34 * (-(v34))), F2: (v2 + (v2 >> ((192) & 63)))}
                                                                            _ = v57
                                                                        case "\n\t":
                                                                            var v58 uint8 = 83
                                                                            _ = v58
                                                                        case "a":
                                                                            v47 += 255
                                                                            switch (v22 + "hello") {
                                                                            case "hello":
                                                                                switch (v4 + v22) {
                                                                                case "hello":
                                                                                    switch fi21 {
                                                                                    case 18:
                                                                                        var v59 uint32 = 849
                                                                                        _ = v59
                                                                                    default:
                                                                                        v3 = append(v3, ((((v22 + ((v22 + v4) + (v4 + ((v22 + v4) + (v22 + v4))))) + (v22 + ("\n\t" + "hello"))) + v4) + "foo bar"))
                                                                                    }
                                                                                    var v60 string = v4
                                                                                    _ = v60
                                                                                }
                                                                            }
                                                                            var v61 uint64 = 4294967295
                                                                            _ = v61
                                                                        default:
                                                                            var v62 int = fi30
                                                                            _ = v62
                                                                        }
                                                                        if v52 {
                                                                            switch fi30 {
                                                                            case -221:
                                                                                switch (fi30 ^ (fi21 - (^(fi30)))) {
                                                                                case 316:
                                                                                    var v63 uint32 = 149
                                                                                    _ = v63
                                                                                case 9223372036854775807:
                                                                                    if v52 {
                                                                                        switch (fi30 - int(v34)) {
                                                                                        case 128:
                                                                                            var v64 uint64 = (^(v1))
                                                                                            _ = v64
                                                                                        case -2147483648:
                                                                                            var v65 int = fi21
                                                                                            _ = v65
                                                                                        case 0:
                                                                                            fmt.Println(v51, v22, fi30)
                                                                                        }
                                                                                    }
                                                                                case -1:
                                                                                    switch "\n\t" {
                                                                                    case "\n\t":
                                                                                        v66 := GenS10{F0: 2, F1: -1, F2: 4294967295}
                                                                                        _ = v66
                                                                                    case "":
                                                                                        var v67 int8 = 100
                                                                                        _ = v67
                                                                                    default:
                                                                                        v51 = v52
                                                                                    }
                                                                                }
                                                                            }
                                                                            v68 := GenS10{F0: 42, F1: 0, F2: 255}
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                v69 := GenS8{F0: 1, F1: 42, F2: v2}
                                                                                _ = v69
                                                                            case "":
                                                                                v70 := GenS13{F0: (fi21 << ((genFn1(9223372036854775807)) & 63))}
                                                                                _ = v70
                                                                            case "a":
                                                                                if true {
                                                                                    v71 := GenS14{F0: v47, F1: 242, F2: -1}
                                                                                    _ = v71
                                                                                } else {
                                                                                    v34 += -1
                                                                                }
                                                                            default:
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            }
                                                                            _ = v68
                                                                        }
                                                                        if true {
                                                                            var v72 uint8 = 40
                                                                            v47 += v47
                                                                            v72 = (v72 | (^(v72)))
                                                                            _ = v72
                                                                        } else {
                                                                            v51 = (2147483647 <= 65535)
                                                                            var v73 uint = (v47 & v47)
                                                                            v52 = (255 >= 100)
                                                                            var v74 string = (("\n\t" + ("hello" + v22)) + v4)
                                                                            if (0.0 > (1e10 + 3.14159)) {
                                                                                var v75 float32 = (3.14159 * (3.14159 * 3.14159))
                                                                                v51 = (v52 || (255 < (v47 % ((890) | 1))))
                                                                                _ = v75
                                                                            } else {
                                                                                var v76 int = 128
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                _ = v76
                                                                            }
                                                                            _ = v73
                                                                            _ = v74
                                                                        }
                                                                    }
                                                                    _ = v47
                                                                    _ = v51
                                                                    _ = v52
                                                                }
                                                            }
                                                            v77 := GenS8{F0: 331, F1: 65535, F2: (^(v2))}
                                                            _ = len(v3)
                                                            _ = v77
                                                        } else {
                                                            var v78 string = v22
                                                            _ = v78
                                                        }
                                                    }
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v79 := GenS11{F0: 361}
                                                        _ = v79
                                                    default:
                                                        switch (fi21 / ((0) | 1)) {
                                                        case 32767:
                                                            fmt.Println(v4, fi30, fi30)
                                                        default:
                                                            v1 = (v1 ^ v1)
                                                        }
                                                    }
                                                    _ = v35
                                                    _ = v36
                                                }
                                                _ = v34
                                            } else {
                                                v2 -= 4294967295
                                            }
                                            var v80 uint64 = 403
                                            _ = v80
                                        }
                                        _ = fi30
                                    }
                                }
                            case -9223372036854775808:
                                switch v22 {
                                case "":
                                    switch 9223372036854775807 {
                                    case -9223372036854775808:
                                        if (42 > -1) {
                                            for fi81 := 0; fi81 < 4; fi81++ {
                                                switch 127 {
                                                case 1:
                                                    var v82 string = v22
                                                    _ = v82
                                                case -1:
                                                    _ = len(v3)
                                                }
                                                if false {
                                                    var v83 uint8 = 83
                                                    v4 = "hello"
                                                    v84 := GenS7{F0: (fi21 << ((4294967295) & 63))}
                                                    v1 -= (v1 * (^(v2)))
                                                    v85 := GenS15{F0: v1, F1: 2}
                                                    _ = v83
                                                    _ = v84
                                                    _ = v85
                                                }
                                                _ = len(v3)
                                                _ = fi81
                                            }
                                            var v86 uint16 = 42
                                            _ = v86
                                        } else {
                                            v87 := GenS16{F0: (fi21 ^ genFn0(218))}
                                            var v88 uint64 = (v1 / ((811) | 1))
                                            v22 = "\n\t"
                                            _ = v87
                                            _ = v88
                                        }
                                    default:
                                        var v89 int = (fi21 | int(v1))
                                        _ = v89
                                    }
                                }
                            case 128:
                                var v90 uint = 4294967295
                                _ = v90
                            }
                        }
                        _ = fi21
                        _ = v22
                    }
                }
            }
        }
    } else {
        var v91 int = genFn0(1)
        v92 := GenS13{F0: 458}
        var v93 int8 = 42
        _ = v91
        _ = v92
        _ = v93
    }
    v4 = ((v4 + v4) + v4)
    v4 = v4
    for fi94 := 0; fi94 < 12; fi94++ {
        var v95 uint8 = 255
        v2 -= v2
        if len(v3) > 0 {
            v3[0] = (v4 + v4)
        }
        _ = fi94
        _ = v95
    }
    var v96 float32 = 0.1
    var v97 uint32 = 4
    var v98 int32 = -801
    v3 = append(v3, "\n\t")
    v2 = (^(v1))
    var v99 uint64 = 4294967295
    switch -1 {
    case 127:
        if true {
            _ = len(v3)
            v3 = append(v3, (v4 + v4))
        }
    }
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        for fi100 := 0; fi100 < 6; fi100++ {
            var v101 string = (v4 + "\n\t")
            if false {
                if true {
                    var v102 int64 = -9223372036854775808
                    switch v101 {
                    case "hello":
                        var v103 int8 = 100
                        _ = v103
                    case "":
                        _ = len(v3)
                    case "foo bar":
                        v3 = append(v3, (v101 + v101))
                    default:
                        for fi104 := 0; fi104 < 1; fi104++ {
                            if true {
                                if ((float64(float32(((3.14159 * (((0.1 / (float64((0.0 * ((v96 + 1.5) - 0.0))) / (float64(v96) / float64(v96)))) / (0.1 + ((1e10 * -2.25) + 1.5))) + 1e10)) * (-2.25 + 0.0)))) < 1e10) || false) {
                                    v99 = 4294967295
                                    switch v101 {
                                    case "a":
                                        v101 = "hello"
                                    default:
                                        if true {
                                            v105 := GenS4{F0: 854, F1: 222, F2: fi104}
                                            _ = v105
                                        }
                                    }
                                } else {
                                    fmt.Println(v102, v101, v101)
                                    var v106 int64 = v102
                                    v107 := GenS15{F0: 18446744073709551615, F1: 0}
                                    _ = v106
                                    _ = v107
                                }
                            } else {
                                var v108 int32 = (v98 + v98)
                                var v109 int16 = 1
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v108
                                _ = v109
                            }
                            v110 := GenS14{F0: 4294967295, F1: v98, F2: v102}
                            fmt.Println(v4, v96)
                            _ = fi104
                            _ = v110
                        }
                    }
                    _ = v102
                }
                if (32767 <= 32767) {
                    v101 = (("hello" + "") + (("\n\t" + ("foo bar" + v4)) + (("hello" + v4) + v101)))
                    for fi111 := 0; fi111 < 6; fi111++ {
                        var v112 string = v101
                        _ = fi111
                        _ = v112
                    }
                }
            } else {
                for fi113 := 0; fi113 < 10; fi113++ {
                    if (int(v97) == fi100) {
                        if (!false) {
                            var v114 uint64 = v99
                            v115 := GenS17{F0: 32767, F1: (v98 >> ((42) & 63)), F2: fi113}
                            _ = v114
                            _ = v115
                        }
                    } else {
                        var v116 int64 = 42
                        v98 += (v98 & (v98 >> ((255) & 63)))
                        v1 = v2
                        v4 = "\n\t"
                        v117 := GenS14{F0: 56, F1: v98, F2: v116}
                        _ = v116
                        _ = v117
                    }
                    _ = fi113
                }
                v1 = v99
            }
            for fi118 := 0; fi118 < 13; fi118++ {
                if false {
                    if len(v3) > 0 {
                        v3[0] = "a"
                    }
                    v98 += (v98 << ((255) & 63))
                    var v119 bool = ((!true) && (-954 < 1000))
                    if v119 {
                        var v120 uint16 = 710
                        switch "\n\t" {
                        case "foo bar":
                            v121 := GenS17{F0: v120, F1: int32(fi100), F2: fi118}
                            _ = v121
                        default:
                            var v122 bool = v119
                            _ = v122
                        }
                        v101 = ""
                        _ = v120
                    }
                    _ = v119
                } else {
                    if false {
                        if (!true) {
                            if len(v3) > 0 {
                                v3[0] = ((("hello" + (v4 + "hello")) + v4) + "foo bar")
                            }
                            _ = len(v3)
                            if len(v3) > 0 {
                                v3[0] = v4
                            }
                            v1 += v99
                        }
                    }
                }
                v123 := GenS18{F0: (v97 % ((v97) | 1))}
                var v124 int8 = -1
                _ = fi118
                _ = v123
                _ = v124
            }
            _ = fi100
            _ = v101
        }
        if false {
            var v125 int = 76
            _ = v125
        }
        _ = len(v3)
    }
    for fi126 := 0; fi126 < 2; fi126++ {
        if true {
            _ = len(v3)
            v99 -= 42
            v98 = -1
        }
        fmt.Println(v97)
        v127 := GenS18{F0: v97}
        _ = fi126
        _ = v127
    }
    v99 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v96
    _ = v97
    _ = v98
    _ = v99
}
