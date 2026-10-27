package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int32; F1 int32; F2 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 uint16; F1 int64; F2 string }
type GenS6 struct { F0 uint32; F1 int16; F2 uint64 }
type GenS7 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int64; F1 int }
type GenS12 struct { F0 int8; F1 uint8 }
type GenS13 struct { F0 int32; F1 uint32 }
type GenS14 struct { F0 uint8; F1 uint8; F2 uint }
type GenS15 struct { F0 int64; F1 int64 }
type GenS16 struct { F0 uint; F1 bool }
type GenS17 struct { F0 int16 }
type GenS18 struct { F0 bool; F1 bool }
func genFn0(p0 uint32, p1 int64) int64 {
    v4 := GenS1{F0: (p1 % (((^(p1))) | 1)), F1: p1}
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 uint64 = 18446744073709551615
        v7 := GenS1{F0: (p1 >> ((65535) & 63)), F1: p1}
        fmt.Println(p0, p1)
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(p0, p0)
    switch -409 {
    case 128:
        var v8 float32 = (float32((3.14159 - (-2.25 * float64(0.1)))) - 3.14159)
        _ = v8
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v4
    return p1
}

func main() {
    var v1 uint = 1
    v2 := make(map[uint32]int)
    var v3 string = ((("" + ("hello" + ("hello" + "foo bar"))) + ("foo bar" + "a")) + "")
    if (905 > 18446744073709551615) {
        fmt.Println(v3, v3)
        v1 = 42
    }
    switch 151 {
    case 4294967295:
        v9 := GenS2{F0: 2147483647, F1: -2147483648, F2: -1}
        _ = v9
    case -1:
        v10 := GenS2{F0: 2147483647, F1: -194, F2: 2147483647}
        _ = v10
    }
    v1 = 792
    var v11 int16 = 32767
    var v12 uint = 42
    v12 = (^(v12))
    v13 := GenS1{F0: int64(v11), F1: -759}
    if (!true) {
        if true {
            var v14 int8 = 42
            for fi15 := 0; fi15 < 13; fi15++ {
                if true {
                    fmt.Println(v11)
                    fmt.Println(fi15)
                    var v16 float32 = (float32(1.5) + float32(0.1))
                    _ = v16
                }
                var v17 uint16 = 32767
                var v18 uint = v1
                _ = fi15
                _ = v17
                _ = v18
            }
            v12 -= v1
            _ = v14
        } else {
            if false {
                if ((true && (((v3 + v3) + ("foo bar" + (v3 + v3))) == ("hello" + v3))) || (false && (-2147483648 < 0))) {
                    for fi19 := 0; fi19 < 2; fi19++ {
                        if (true || false) {
                            switch (fi19 - fi19) {
                            case -2147483648:
                                fmt.Println(v3, v1)
                            case 2147483647:
                                var v20 float32 = 3.14159
                                _ = v20
                            }
                        } else {
                            v21 := GenS2{F0: -2147483648, F1: 0, F2: 0}
                            v22 := []int{fi19, 275, (-(fi19)), (fi19 >> ((v1) & 63))}
                            _ = v21
                            _ = v22
                        }
                        _ = fi19
                    }
                }
                if ((false && (-196 > -1)) && (2 == 2147483648)) {
                    if false {
                        var v23 int = -1
                        for fi24 := 0; fi24 < 20; fi24++ {
                            delete(v2, 379)
                            _ = fi24
                        }
                        var v25 uint64 = 42
                        _ = v23
                        _ = v25
                    } else {
                        if (!true) {
                            for fi26 := 0; fi26 < 7; fi26++ {
                                fmt.Println(v1, v12)
                                fmt.Println(v12, v3)
                                if true {
                                    v3 = ((v3 + "foo bar") + v3)
                                    v11 = -32768
                                    if (562 == 395) {
                                        var v27 uint16 = 825
                                        _ = v27
                                    } else {
                                        v28 := v2[429]
                                        for fi29 := 0; fi29 < 19; fi29++ {
                                            if ((false && (740 != 439)) || true) {
                                                switch (("a" + v3) + (("hello" + v3) + (v3 + v3))) {
                                                case "hello":
                                                    if (3.14159 >= (3.14159 * (float64(1.5) + 1e10))) {
                                                        switch fi29 {
                                                        case 32768:
                                                            _ = len(v2)
                                                        case 32767:
                                                            var v30 float32 = (1.5 * (-2.25 - float32((-2.25 - (((1.5 * 1.5) * 3.14159) - -2.25)))))
                                                            _ = v30
                                                        case -1:
                                                            fmt.Println(v1)
                                                        }
                                                        switch (((v3 + "hello") + v3) + (v3 + "a")) {
                                                        case "a":
                                                            var v31 float64 = 0.0
                                                            _ = v31
                                                        case "\n\t":
                                                            var v32 uint32 = 1
                                                            _ = v32
                                                        }
                                                    } else {
                                                        switch fi29 {
                                                        case -796:
                                                            v3 = "a"
                                                        }
                                                        var v33 uint16 = 65535
                                                        v3 = v3
                                                        _ = v33
                                                    }
                                                case "":
                                                    v1 -= v1
                                                default:
                                                    v12 = (v1 << (((v12 + v12)) & 63))
                                                }
                                            } else {
                                                fmt.Println(v3, v3, v12)
                                                switch (v3 + "") {
                                                case "a":
                                                    v34 := []int{}
                                                    _ = v34
                                                case "\n\t":
                                                    fmt.Println(v28, fi26, fi26)
                                                    delete(v2, 2)
                                                    v35 := GenS2{F0: -316, F1: 1, F2: 594}
                                                    v36 := []int{fi29}
                                                    if (255 > uint16(v1)) {
                                                        if true {
                                                            v11 += v11
                                                        } else {
                                                            fmt.Println(v28, fi29)
                                                            v12 -= (v1 >> ((59) & 63))
                                                            _ = len(v2)
                                                        }
                                                        v2[90] = v28
                                                    } else {
                                                        v28 = -1
                                                        v37 := GenS3{F0: 47}
                                                        _ = v37
                                                    }
                                                    _ = v35
                                                    _ = v36
                                                }
                                                v38 := GenS2{F0: -1, F1: -141, F2: 1}
                                                if (!(42 > 255)) {
                                                    delete(v2, 2)
                                                    var v39 uint64 = 76
                                                    if false {
                                                        if true {
                                                            var v40 int64 = genFn0(0, 0)
                                                            delete(v2, 0)
                                                            if ((1.5 == 1.5) || false) {
                                                                switch (((v3 + v3) + ("a" + v3)) + (v3 + "hello")) {
                                                                case "foo bar":
                                                                    var v41 uint8 = 154
                                                                    _ = v41
                                                                default:
                                                                    if ((!(true || true)) || ((v40 / ((42) | 1)) < -1)) {
                                                                        if (v40 < v40) {
                                                                            _ = len(v2)
                                                                            var v42 string = ((("" + v3) + ((v3 + v3) + ("\n\t" + v3))) + "foo bar")
                                                                            fmt.Println(fi26, v42)
                                                                            _ = v42
                                                                        }
                                                                        var v43 bool = (-1 <= 2147483647)
                                                                        _ = v43
                                                                    }
                                                                }
                                                                _ = len(v2)
                                                                v1 = (v12 >> ((422) & 63))
                                                                v44 := v2[2]
                                                                _ = v44
                                                            }
                                                            v45 := GenS4{F0: (v1 ^ v1), F1: 2147483647}
                                                            v46 := v2[2]
                                                            if true {
                                                                v47 := v2[65535]
                                                                fmt.Println(fi26, fi29)
                                                                var v48 int = (fi26 / (((fi29 % (((v28 ^ -258)) | 1))) | 1))
                                                                _ = v47
                                                                _ = v48
                                                            } else {
                                                                v49 := GenS3{F0: 2}
                                                                if true {
                                                                    v39 = (^(v39))
                                                                    if true {
                                                                        switch fi26 {
                                                                        case 4294967295:
                                                                            v28 = (fi29 >> ((v1) & 63))
                                                                        case 101:
                                                                            v40 -= -701
                                                                        case -9223372036854775808:
                                                                            fmt.Println(fi26)
                                                                        }
                                                                    }
                                                                }
                                                                _ = v49
                                                            }
                                                            var v50 int16 = 32767
                                                            _ = v40
                                                            _ = v45
                                                            _ = v46
                                                            _ = v50
                                                        }
                                                        v3 = ""
                                                    }
                                                    _ = v39
                                                } else {
                                                    v12 -= v1
                                                    var v51 int64 = -512
                                                    v12 = 255
                                                    var v52 int8 = 124
                                                    _ = v51
                                                    _ = v52
                                                }
                                                _ = v38
                                            }
                                            v1 += 234
                                            _ = fi29
                                        }
                                        fmt.Println(v1, v1, v11)
                                        _ = v28
                                    }
                                }
                                _ = fi26
                            }
                            v53 := GenS4{F0: v12, F1: 317}
                            var v54 uint8 = 2
                            _ = v53
                            _ = v54
                        } else {
                            for fi55 := 0; fi55 < 8; fi55++ {
                                for fi56 := 0; fi56 < 17; fi56++ {
                                    v3 = ("\n\t" + "hello")
                                    switch v3 {
                                    case "hello":
                                        if true {
                                            if (4294967295 < 18446744073709551615) {
                                                if (false && false) {
                                                    v1 = 0
                                                } else {
                                                    fmt.Println(v1, fi55)
                                                    var v57 uint = v12
                                                    delete(v2, 65535)
                                                    _ = v57
                                                }
                                                if true {
                                                    v12 = v12
                                                    var v58 bool = (true || (false || (2147483647 >= 2)))
                                                    _ = v58
                                                }
                                            } else {
                                                v59 := GenS5{F0: 919, F1: 1, F2: (("hello" + "a") + ((((((v3 + "") + "foo bar") + "a") + ((v3 + (v3 + "a")) + "hello")) + v3) + (v3 + ((("" + "foo bar") + v3) + v3))))}
                                                v3 = (((v3 + "hello") + ((v3 + v3) + v3)) + v3)
                                                var v60 uint32 = 739
                                                _ = v59
                                                _ = v60
                                            }
                                            switch v3 {
                                            case "a":
                                                v61 := GenS6{F0: 580, F1: (v11 << ((65535) & 63)), F2: 864}
                                                switch (v3 + v3) {
                                                case "":
                                                    var v62 uint = 1
                                                    _ = v62
                                                case "a":
                                                    v2[218] = (fi55 / ((0) | 1))
                                                case "hello":
                                                    if (42 >= -9223372036854775808) {
                                                        var v63 uint32 = 42
                                                        if ((1.5 + 0.1) >= 1e10) {
                                                            if (((v1 % (((v12 & (v12 << ((v12) & 63)))) | 1)) == (v1 + v12)) || (-2.25 == 1.5)) {
                                                                fmt.Println(v12)
                                                                var v64 int8 = int8(fi56)
                                                                var v65 int8 = (v64 + v64)
                                                                _ = v64
                                                                _ = v65
                                                            } else {
                                                                var v66 uint32 = 177
                                                                fmt.Println(v63, v1, v1)
                                                                v67 := GenS7{F0: 0, F1: 255, F2: 127}
                                                                switch (v3 + v3) {
                                                                case "a":
                                                                    fmt.Println(v12, v11)
                                                                default:
                                                                    v68 := []int{fi56}
                                                                    _ = v68
                                                                }
                                                                v69 := []int{}
                                                                _ = v66
                                                                _ = v67
                                                                _ = v69
                                                            }
                                                            fmt.Println(v3, fi55, v3)
                                                        } else {
                                                            var v70 int = fi56
                                                            v11 -= v11
                                                            v71 := GenS8{F0: v3}
                                                            _ = v70
                                                            _ = v71
                                                        }
                                                        _ = v63
                                                    } else {
                                                        fmt.Println(v11, fi56, fi55)
                                                        var v72 uint32 = uint32(v1)
                                                        var v73 int16 = 0
                                                        _ = v72
                                                        _ = v73
                                                    }
                                                default:
                                                    v74 := GenS4{F0: v12, F1: -1}
                                                    _ = v74
                                                }
                                                switch "\n\t" {
                                                case "":
                                                    v12 += (v12 << (((v12 - v12)) & 63))
                                                case "a":
                                                    v75 := GenS5{F0: 2, F1: -301, F2: ((v3 + v3) + v3)}
                                                    _ = v75
                                                case "hello":
                                                    v76 := []int{2147483647, fi55, 32768, (fi56 * (^(fi56)))}
                                                    _ = v76
                                                }
                                                _ = v61
                                            case "hello":
                                                var v77 int16 = v11
                                                _ = v77
                                            case "foo bar":
                                                if true {
                                                    v78 := []int{4294967295, (fi55 + (fi56 / ((fi56) | 1))), 32768}
                                                    var v79 int16 = v11
                                                    _ = v78
                                                    _ = v79
                                                } else {
                                                    if true {
                                                        var v80 int = (fi55 / (((fi55 - (fi55 % ((fi55) | 1)))) | 1))
                                                        if ("\n\t" < "foo bar") {
                                                            if (-9223372036854775808 > 27) {
                                                                if true {
                                                                    v11 = v11
                                                                } else {
                                                                    if (-9223372036854775808 < 772) {
                                                                        var v81 int32 = 42
                                                                        switch 127 {
                                                                        case -512:
                                                                            var v82 string = v3
                                                                            v11 = (-(v11))
                                                                            switch (v80 % ((9223372036854775807) | 1)) {
                                                                            case 170:
                                                                                if false {
                                                                                    v83 := GenS9{F0: 0}
                                                                                    v81 = (v81 + (v81 * int32(fi55)))
                                                                                    _ = v83
                                                                                }
                                                                            case 707:
                                                                                fmt.Println(v81, fi56, fi55)
                                                                            }
                                                                            _ = v82
                                                                        case -9223372036854775808:
                                                                            if true {
                                                                                delete(v2, 733)
                                                                            }
                                                                        case -2147483648:
                                                                            v84 := GenS5{F0: 1, F1: genFn0(2, 0), F2: v3}
                                                                            _ = v84
                                                                        }
                                                                        _ = v81
                                                                    }
                                                                    v80 = -120
                                                                }
                                                            } else {
                                                                v85 := []int{(fi56 << ((v1) & 63)), 454, v80, 4294967295}
                                                                v11 -= (v11 % ((0) | 1))
                                                                v80 = (^(fi55))
                                                                _ = v85
                                                            }
                                                            if ((42 > 484) || (true && true)) {
                                                                if (0 <= 127) {
                                                                    var v86 int16 = -32768
                                                                    _ = v86
                                                                } else {
                                                                    var v87 float64 = -2.25
                                                                    _ = v87
                                                                }
                                                                switch fi56 {
                                                                case -512:
                                                                    v12 = 42
                                                                case -9223372036854775808:
                                                                    if false {
                                                                        if (32768 <= v80) {
                                                                            v88 := GenS2{F0: -2147483648, F1: 947, F2: 670}
                                                                            _ = v88
                                                                        }
                                                                    }
                                                                case -2147483648:
                                                                    fmt.Println(v12)
                                                                }
                                                                v89 := GenS2{F0: 1, F1: 1, F2: 42}
                                                                v90 := []int{fi55, (v80 | (v80 % ((fi56) | 1))), fi56, v80}
                                                                _ = v89
                                                                _ = v90
                                                            } else {
                                                                delete(v2, 2)
                                                            }
                                                        } else {
                                                            v91 := []int{(-(fi56)), -514, (fi56 | fi56)}
                                                            switch v80 {
                                                            case 702:
                                                                var v92 float32 = (-2.25 / 1.0)
                                                                _ = v92
                                                            case -701:
                                                                fmt.Println(v80)
                                                            case 1:
                                                                var v93 int64 = 0
                                                                _ = v93
                                                            }
                                                            v80 = 32767
                                                            v2[0] = fi55
                                                            if (!false) {
                                                                v1 = (v12 & v12)
                                                                v94 := GenS3{F0: 53}
                                                                if len(v91) >= 2 {
                                                                    v91 = v91[1:len(v91)]
                                                                }
                                                                _ = v94
                                                            } else {
                                                                var v95 uint16 = 507
                                                                switch 4294967295 {
                                                                case 127:
                                                                    v96 := GenS8{F0: v3}
                                                                    _ = v96
                                                                case 631:
                                                                    delete(v2, 2)
                                                                }
                                                                if (1e10 <= 3.14159) {
                                                                    if false {
                                                                        switch (v80 + 4294967295) {
                                                                        case -712:
                                                                            v95 -= 255
                                                                        case -3:
                                                                            v11 = v11
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v80, fi56)
                                                                        if true {
                                                                            delete(v2, 4294967295)
                                                                            v3 = (v3 + v3)
                                                                            var v97 string = v3
                                                                            _ = v97
                                                                        }
                                                                    }
                                                                    if false {
                                                                        _ = len(v2)
                                                                        var v98 float64 = ((0.0 * 0.1) + ((1.5 / 1.0) * (0.1 + -2.25)))
                                                                        _ = v98
                                                                    } else {
                                                                        v3 = ""
                                                                        switch 42 {
                                                                        case 98:
                                                                            v99 := v2[2]
                                                                            _ = v99
                                                                        default:
                                                                            v1 += v1
                                                                            fmt.Println(fi55, v3)
                                                                            v1 += (v12 & 4294967295)
                                                                        }
                                                                        v100 := GenS2{F0: 1, F1: 42, F2: 2147483647}
                                                                        _ = v100
                                                                    }
                                                                    v101 := GenS5{F0: (^(v95)), F1: 42, F2: ("foo bar" + v3)}
                                                                    var v102 int32 = -749
                                                                    _ = v101
                                                                    _ = v102
                                                                } else {
                                                                    var v103 string = v3
                                                                    if false {
                                                                        if true {
                                                                            if true {
                                                                                if (true || (!(573 == 795))) {
                                                                                    v1 -= v12
                                                                                    var v104 int = 32768
                                                                                    if false {
                                                                                        var v105 int32 = 0
                                                                                        v80 = (^(fi56))
                                                                                        var v106 int16 = (v11 & -32768)
                                                                                        _ = v105
                                                                                        _ = v106
                                                                                    }
                                                                                    _ = v104
                                                                                } else {
                                                                                    v107 := GenS10{F0: (v11 / ((-283) | 1))}
                                                                                    if (true || (!true)) {
                                                                                        var v108 uint = 255
                                                                                        fmt.Println(fi55, v11)
                                                                                        _ = v108
                                                                                    } else {
                                                                                        if true {
                                                                                            var v109 int8 = 100
                                                                                            if true {
                                                                                                _ = len(v91)
                                                                                            }
                                                                                            var v110 float32 = 0.1
                                                                                            _ = v109
                                                                                            _ = v110
                                                                                        }
                                                                                        v12 -= 4294967295
                                                                                    }
                                                                                    _ = v107
                                                                                }
                                                                                v80 = fi56
                                                                            }
                                                                            switch v103 {
                                                                            case "a":
                                                                                if false {
                                                                                    var v111 int8 = 127
                                                                                    _ = v111
                                                                                }
                                                                            case "":
                                                                                var v112 int64 = genFn0(2147483647, 1)
                                                                                _ = v112
                                                                            default:
                                                                                switch fi56 {
                                                                                case 0:
                                                                                    delete(v2, 4294967295)
                                                                                case 960:
                                                                                    if false {
                                                                                        delete(v2, 2147483648)
                                                                                    }
                                                                                case 32768:
                                                                                    switch fi56 {
                                                                                    case -1:
                                                                                        v95 = (v95 % ((v95) | 1))
                                                                                        var v113 int = v80
                                                                                        _ = len(v91)
                                                                                        _ = v113
                                                                                    }
                                                                                }
                                                                            }
                                                                            switch (v3 + v103) {
                                                                            case "foo bar":
                                                                                delete(v2, 2147483648)
                                                                            case "":
                                                                                v114 := GenS11{F0: 42, F1: fi56}
                                                                                _ = v114
                                                                            default:
                                                                                var v115 float32 = float32((3.14159 - 3.14159))
                                                                                if (!(true || false)) {
                                                                                    v115 = ((-2.25 * v115) - float32(0.1))
                                                                                    var v116 uint32 = 841
                                                                                    v117 := GenS12{F0: 64, F1: 127}
                                                                                    if (v80 > -36) {
                                                                                        _ = len(v2)
                                                                                        v116 = uint32(v95)
                                                                                        v118 := GenS13{F0: -626, F1: (^(v116))}
                                                                                        _ = v118
                                                                                    } else {
                                                                                        var v119 float64 = (0.1 - ((3.14159 * (((1.5 + 1e10) / 0.1) / 3.14159)) - 3.14159))
                                                                                        _ = v119
                                                                                    }
                                                                                    _ = v116
                                                                                    _ = v117
                                                                                } else {
                                                                                    switch v103 {
                                                                                    case "hello":
                                                                                        switch v3 {
                                                                                        case "hello":
                                                                                            v120 := GenS1{F0: -1, F1: 42}
                                                                                            _ = v120
                                                                                        case "":
                                                                                            delete(v2, uint32(v95))
                                                                                        default:
                                                                                            v2[2147483647] = fi56
                                                                                        }
                                                                                    }
                                                                                    switch fi56 {
                                                                                    case 32768:
                                                                                        fmt.Println(v95)
                                                                                    case -2147483648:
                                                                                        fmt.Println(v12)
                                                                                    default:
                                                                                        var v121 uint = v1
                                                                                        _ = v121
                                                                                    }
                                                                                    if (-128 <= -29) {
                                                                                        if ((v95 | v95) != 255) {
                                                                                            if true {
                                                                                                switch v103 {
                                                                                                case "foo bar":
                                                                                                    if (127 >= 2) {
                                                                                                        var v122 float32 = v115
                                                                                                        v12 = (v12 & v12)
                                                                                                        _ = v122
                                                                                                    }
                                                                                                case "":
                                                                                                    var v123 bool = ((1 != 394) || false)
                                                                                                    _ = v123
                                                                                                }
                                                                                                _ = len(v2)
                                                                                                v11 += -1
                                                                                            }
                                                                                        }
                                                                                        v91 = append(v91, (fi55 * fi55))
                                                                                    }
                                                                                }
                                                                                v124 := GenS8{F0: (v3 + (("hello" + "") + ("foo bar" + v3)))}
                                                                                v11 -= (v11 * -32768)
                                                                                _ = v115
                                                                                _ = v124
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v103
                                                                }
                                                                var v125 uint32 = 1
                                                                fmt.Println(v125)
                                                                _ = v125
                                                                _ = v95
                                                            }
                                                            _ = v91
                                                        }
                                                        _ = v80
                                                    } else {
                                                        var v126 float32 = float32(0.1)
                                                        v127 := v2[2147483647]
                                                        _ = v126
                                                        _ = v127
                                                    }
                                                    v128 := GenS14{F0: 246, F1: 64, F2: v12}
                                                    v1 -= v12
                                                    var v129 int16 = v11
                                                    var v130 uint8 = 127
                                                    _ = v128
                                                    _ = v129
                                                    _ = v130
                                                }
                                            default:
                                                var v131 int64 = 0
                                                _ = v131
                                            }
                                            if false {
                                                v3 = ("hello" + ((v3 + v3) + v3))
                                                var v132 uint = (^(v1))
                                                if false {
                                                    if (int64(v1) < 0) {
                                                        v133 := []int{fi56, -411, -9223372036854775808}
                                                        var v134 uint8 = 0
                                                        v135 := GenS15{F0: 42, F1: genFn0(0, -9223372036854775808)}
                                                        _ = v133
                                                        _ = v134
                                                        _ = v135
                                                    } else {
                                                        if (((0.0 - 0.1) <= 1.5) && true) {
                                                            var v136 uint32 = 42
                                                            _ = len(v2)
                                                            v3 = v3
                                                            var v137 uint32 = v136
                                                            _ = v136
                                                            _ = v137
                                                        } else {
                                                            var v138 uint16 = 32767
                                                            var v139 uint64 = 4294967295
                                                            _ = v138
                                                            _ = v139
                                                        }
                                                        v3 = v3
                                                    }
                                                    v140 := GenS14{F0: 144, F1: 2, F2: (v132 / (((v12 % ((240) | 1))) | 1))}
                                                    _ = v140
                                                }
                                                _ = v132
                                            } else {
                                                v3 = ("hello" + ("a" + v3))
                                                var v141 int64 = 985
                                                _ = v141
                                            }
                                        }
                                    case "foo bar":
                                        v1 += (v12 + v1)
                                    }
                                    _ = fi56
                                }
                                _ = fi55
                            }
                        }
                        var v142 uint16 = 190
                        v12 -= (v1 + (^(v1)))
                        _ = v142
                    }
                } else {
                    v11 = (v11 | (v11 << (((v1 + (v1 - 0))) & 63)))
                    if ((v11 + v11) > (v11 & (-(v11)))) {
                        v12 = (^(v12))
                        v143 := GenS7{F0: 65535, F1: 202, F2: 127}
                        var v144 uint8 = 10
                        _ = v143
                        _ = v144
                    } else {
                        v12 = v12
                        var v145 bool = (0 != 127)
                        for fi146 := 0; fi146 < 10; fi146++ {
                            switch 32768 {
                            case 32767:
                                switch v3 {
                                case "a":
                                    _ = len(v2)
                                }
                            case 2147483647:
                                v147 := GenS16{F0: v1, F1: (!v145)}
                                _ = v147
                            case -2147483648:
                                for fi148 := 0; fi148 < 3; fi148++ {
                                    v145 = v145
                                    v2[1] = fi146
                                    _ = fi148
                                }
                            default:
                                var v149 int16 = (v11 * v11)
                                _ = v149
                            }
                            if v145 {
                                var v150 int8 = -1
                                _ = v150
                            } else {
                                var v151 uint32 = 827
                                switch v3 {
                                case "foo bar":
                                    for fi152 := 0; fi152 < 16; fi152++ {
                                        v153 := GenS17{F0: 0}
                                        _ = fi152
                                        _ = v153
                                    }
                                case "hello":
                                    var v154 float64 = (float64(0.1) - (0.1 - 3.14159))
                                    _ = v154
                                case "a":
                                    var v155 bool = (1 < -349)
                                    _ = v155
                                }
                                _ = v151
                            }
                            _ = fi146
                        }
                        _ = v145
                    }
                    for fi156 := 0; fi156 < 4; fi156++ {
                        fmt.Println(v1, v3)
                        var v157 uint8 = 2
                        _ = fi156
                        _ = v157
                    }
                }
                if (100 > -1) {
                    v12 = 4294967295
                    v1 -= (v1 + (v1 >> ((2) & 63)))
                    v1 = v12
                }
            } else {
                v12 -= (v1 >> ((v1) & 63))
                v158 := []int{-981, 1, -1, 32768}
                v159 := GenS18{F0: false, F1: (966 == 748)}
                _ = v158
                _ = v159
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
