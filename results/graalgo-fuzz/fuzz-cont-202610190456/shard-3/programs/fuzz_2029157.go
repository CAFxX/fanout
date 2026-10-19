package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32; F1 uint; F2 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool; F1 int }
type GenS8 struct { F0 string }
func genFn0() uint {
    var v7 float32 = float32(3.14159)
    v8 := GenS1{F0: 32767, F1: -373}
    v7 = 3.14159
    v7 = v7
    var v9 string = "a"
    _ = v7
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]uint32)
    v3 := []uint32{2, 1, 1, 941}
    var v4 int64 = 42
    var v5 uint8 = 255
    var v6 float64 = float64((0.0 + 3.14159))
    if ((1e10 * 0.1) < 0.0) {
        var v10 float64 = v6
        v4 += (v4 & 9223372036854775807)
        _ = v10
    }
    var v11 string = v1
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 string = "foo bar"
        _ = fi12
        _ = v13
    }
    var v14 string = (v11 + v1)
    if (false || true) {
        var v15 int16 = -1
        v4 -= v4
        _ = v15
    } else {
        v16 := v2[1]
        _ = v16
    }
    var v17 uint32 = 0
    delete(v2, 4294967295)
    v18 := GenS1{F0: 355, F1: -769}
    switch 9223372036854775807 {
    case 128:
        switch -844 {
        case 2147483647:
            v19 := GenS2{F0: v4}
            _ = v19
        }
    case 9223372036854775807:
        if false {
            v20 := GenS1{F0: 251, F1: 0}
            if true {
                v11 = ("\n\t" + "\n\t")
                _ = len(v3)
                var v21 int = 127
                _ = v21
            } else {
                v22 := GenS2{F0: v4}
                var v23 string = (v1 + v11)
                _ = v22
                _ = v23
            }
            _ = v20
        }
    case -1:
        v14 = "hello"
    }
    if false {
        delete(v2, v17)
        var v24 int16 = -1
        for fi25 := 0; fi25 < 12; fi25++ {
            var v26 float64 = v6
            switch fi25 {
            case -5:
                if false {
                    var v27 float32 = 1e10
                    _ = v27
                } else {
                    fmt.Println(v24)
                    var v28 int64 = v4
                    _ = v28
                }
            case 311:
                switch 4294967295 {
                case -818:
                    if false {
                        var v29 float64 = (((v26 * 1e10) + (v6 * v26)) * v26)
                        if (true && (4294967295 != 0)) {
                            if (-9223372036854775808 < 127) {
                                switch v1 {
                                case "\n\t":
                                    v14 = (v11 + v14)
                                case "a":
                                    v30 := GenS2{F0: (v4 | (v4 * (v4 << ((887) & 63))))}
                                    _ = v30
                                }
                                if false {
                                    for fi31 := 0; fi31 < 19; fi31++ {
                                        fmt.Println(v11, v24, v1)
                                        v4 = v4
                                        _ = fi31
                                    }
                                    v26 += 1.5
                                    v32 := GenS1{F0: v24, F1: 330}
                                    _ = v32
                                } else {
                                    for fi33 := 0; fi33 < 10; fi33++ {
                                        var v34 uint = 42
                                        var v35 int64 = (v4 / ((int64(v24)) | 1))
                                        v36 := GenS3{F0: -1}
                                        _ = fi33
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                    }
                                    for fi37 := 0; fi37 < 18; fi37++ {
                                        fmt.Println(v1)
                                        if false {
                                            fmt.Println(v11, v29)
                                        } else {
                                            v2[v17] = 798
                                        }
                                        switch (fi25 % ((int(v4)) | 1)) {
                                        case 843:
                                            if (!(((1.5 - -2.25) + (0.1 - (float32((float64(1.5) / float64((((float32(0.1) * 3.14159) * 0.1) - 3.14159)))) * (1.5 + 1e10)))) == (-2.25 / -2.25))) {
                                                if true {
                                                    var v38 float64 = 0.1
                                                    v4 += v4
                                                    fmt.Println(v26)
                                                    _ = v38
                                                }
                                                v26 = v6
                                            }
                                        }
                                        if (true || (1.5 > -2.25)) {
                                            v39 := GenS3{F0: -1}
                                            v40 := v2[v17]
                                            _ = v39
                                            _ = v40
                                        } else {
                                            v41 := GenS3{F0: 42}
                                            var v42 int64 = -9223372036854775808
                                            if false {
                                                fmt.Println(v42)
                                                v5 += uint8(v17)
                                            } else {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                v1 = "a"
                                                v5 = v5
                                                v11 = v14
                                                var v43 uint8 = 255
                                                _ = v43
                                            }
                                            v44 := v2[42]
                                            _ = v41
                                            _ = v42
                                            _ = v44
                                        }
                                        delete(v2, (^(v17)))
                                        _ = fi37
                                    }
                                }
                            }
                            if true {
                                var v45 float32 = 1.5
                                _ = v45
                            }
                        } else {
                            if (1 > 42) {
                                var v46 int = fi25
                                fmt.Println(v4)
                                _ = v46
                            }
                        }
                        switch v11 {
                        case "":
                            v6 = (v6 * 3.14159)
                        case "hello":
                            if len(v3) > 0 {
                                v3[0] = 4294967295
                            }
                        case "a":
                            v24 += (v24 * v24)
                        }
                        _ = v29
                    } else {
                        v47 := v2[v17]
                        v3 = append(v3, (v47 ^ (^(v17))))
                        v4 -= v4
                        _ = v47
                    }
                case -57:
                    v3 = append(v3, 0)
                default:
                    v3 = append(v3, 454)
                }
            }
            _ = fi25
            _ = v26
        }
        _ = v24
    }
    v11 = (v11 + ("foo bar" + v1))
    for fi48 := 0; fi48 < 4; fi48++ {
        if (v17 != v17) {
            if (((3.14159 / 0.1) < -2.25) || true) {
                if (-2147483648 < -2147483648) {
                    if false {
                        v14 = "foo bar"
                        var v49 int64 = 42
                        v11 = "foo bar"
                        _ = v49
                    }
                } else {
                    if (164 < 658) {
                        for fi50 := 0; fi50 < 7; fi50++ {
                            if len(v3) > 0 {
                                v3[0] = v17
                            }
                            if (!(fi48 <= (fi48 / (((fi50 | -123)) | 1)))) {
                                var v51 bool = false
                                _ = len(v2)
                                switch fi50 {
                                case -1:
                                    fmt.Println(v11, v11, v14)
                                case 32767:
                                    delete(v2, 0)
                                case 9223372036854775807:
                                    if ((v1 + v14) < "") {
                                        if (false || v51) {
                                            fmt.Println(v51, v1)
                                            if v51 {
                                                v2[v17] = (^(v17))
                                                v14 = ((v11 + v14) + (v1 + v11))
                                            } else {
                                                switch v1 {
                                                case "\n\t":
                                                    v17 = v17
                                                case "a":
                                                    var v52 uint16 = 2
                                                    _ = v52
                                                default:
                                                    var v53 float32 = 1e10
                                                    _ = v53
                                                }
                                            }
                                            v2[(v17 + (v17 ^ v17))] = (v17 | (v17 >> ((950) & 63)))
                                            if v51 {
                                                var v54 int64 = 1
                                                v2[(v17 * v17)] = 4294967295
                                                fmt.Println(v1, v11, v1)
                                                _ = v54
                                            } else {
                                                switch fi50 {
                                                case 485:
                                                    if false {
                                                        if v51 {
                                                            fmt.Println(v1)
                                                        } else {
                                                            fmt.Println(fi48, v4)
                                                            fmt.Println(v1, v4)
                                                            _ = len(v3)
                                                        }
                                                        if len(v3) > 0 {
                                                            v3[0] = 42
                                                        }
                                                        v4 = (v4 / ((v4) | 1))
                                                    }
                                                case 32767:
                                                    fmt.Println(v14, fi48, v6)
                                                case -2147483648:
                                                    v14 = v11
                                                default:
                                                    fmt.Println(v17)
                                                }
                                            }
                                            v3 = append(v3, v17)
                                        } else {
                                            v55 := GenS3{F0: -939}
                                            _ = v55
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        switch (^(fi50)) {
                                        case 9223372036854775807:
                                            var v56 int16 = 32767
                                            switch (v1 + v14) {
                                            case "a":
                                                v57 := GenS4{F0: ("hello" + v1)}
                                                _ = v57
                                            case "foo bar":
                                                if (594 >= (^(v17))) {
                                                    fmt.Println(v1)
                                                    switch (fi48 % ((fi48) | 1)) {
                                                    case 127:
                                                        v51 = ("" < ("hello" + ((v11 + "foo bar") + (v11 + ""))))
                                                    }
                                                    if ((true && v51) && (v56 >= v56)) {
                                                        var v58 string = ""
                                                        _ = v58
                                                    } else {
                                                        v56 += (^(v56))
                                                        v2[(^(v17))] = v17
                                                        var v59 uint8 = 1
                                                        _ = v59
                                                    }
                                                } else {
                                                    fmt.Println(v51)
                                                    fmt.Println(fi50, v1, v4)
                                                    v11 = v1
                                                }
                                                var v60 uint64 = 2
                                                _ = v60
                                            case "hello":
                                                v1 = ("\n\t" + (((v14 + "a") + "foo bar") + v14))
                                            }
                                            _ = v56
                                        }
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if false {
                                        v4 = v4
                                    }
                                default:
                                    v6 -= (v6 + (v6 - (v6 * ((((((float64(-2.25) / v6) * (v6 + 1.5)) - 3.14159) * 1e10) + 0.1) / float64((-2.25 * 3.14159))))))
                                    var v61 int64 = v4
                                    v14 = (((("a" + v14) + v1) + (v1 + "")) + (("foo bar" + v14) + (v1 + "\n\t")))
                                    _ = v61
                                }
                                _ = v51
                            }
                            _ = fi50
                        }
                        v6 = v6
                        v1 = ("a" + ("foo bar" + v1))
                    } else {
                        switch fi48 {
                        case 1:
                            v11 = v11
                        }
                    }
                }
                if false {
                    var v62 uint = genFn0()
                    fmt.Println(v5)
                    var v63 uint = v62
                    _ = v62
                    _ = v63
                } else {
                    v6 = 3.14159
                }
            } else {
                var v64 int8 = 127
                var v65 bool = false
                v66 := GenS2{F0: v4}
                _ = v64
                _ = v65
                _ = v66
            }
            for fi67 := 0; fi67 < 3; fi67++ {
                fmt.Println(v4, v17)
                if (-12 <= int8(v17)) {
                    v17 = 1
                    if false {
                        if false {
                            v68 := GenS5{F0: 286, F1: 779, F2: -1}
                            var v69 int16 = -32768
                            fmt.Println(v14, v11, fi67)
                            _ = v68
                            _ = v69
                        } else {
                            fmt.Println(v1)
                            if (float32(1.5) < 1e10) {
                                v70 := GenS4{F0: v14}
                                _ = v70
                            }
                        }
                        v71 := GenS1{F0: 1000, F1: -1}
                        _ = v71
                    } else {
                        v5 -= (v5 << ((702) & 63))
                        fmt.Println(v17, v17, v11)
                        switch "a" {
                        case "hello":
                            if true {
                                v72 := GenS6{F0: -949}
                                _ = v72
                            } else {
                                switch 2147483647 {
                                case 2147483647:
                                    v6 += ((((v6 + v6) + 0.1) * (((((-2.25 * v6) / 1.0) / 1.5) * 0.0) * -2.25)) * ((v6 / (1.5 - (v6 / (v6 - float64(0.1))))) / (-2.25 / v6)))
                                case -969:
                                    switch int(v17) {
                                    case 127:
                                        switch "" {
                                        case "hello":
                                            if (false || true) {
                                                v73 := GenS5{F0: int32(v4), F1: 477, F2: 0}
                                                v14 = v1
                                                if (!(true || (!(!(1 > 0))))) {
                                                    v5 -= v5
                                                    v14 = ((("foo bar" + v14) + v1) + v14)
                                                } else {
                                                    v74 := GenS4{F0: v1}
                                                    v14 = v14
                                                    _ = v74
                                                }
                                                _ = v73
                                            } else {
                                                v75 := GenS1{F0: 42, F1: 189}
                                                v2[(v17 * v17)] = 42
                                                _ = v75
                                            }
                                        default:
                                            v17 += v17
                                        }
                                    case 1:
                                        v11 = v11
                                    case -160:
                                        v4 = (v4 / ((-9223372036854775808) | 1))
                                    default:
                                        fmt.Println(v11, v5, v1)
                                    }
                                default:
                                    var v76 int32 = 1
                                    _ = v76
                                }
                                fmt.Println(v17)
                                v5 = (v5 * v5)
                                if true {
                                    switch (v1 + v11) {
                                    case "\n\t":
                                        switch (fi67 - (fi48 + (fi48 - fi67))) {
                                        case -1:
                                            var v77 float64 = v6
                                            _ = v77
                                        case 337:
                                            v6 = -2.25
                                        }
                                    default:
                                        fmt.Println(fi48, v4, v6)
                                    }
                                    var v78 uint16 = 65535
                                    _ = v78
                                }
                                var v79 uint32 = (v17 | 2147483648)
                                _ = v79
                            }
                            switch "hello" {
                            case "foo bar":
                                if false {
                                    v6 = v6
                                }
                            case "\n\t":
                                fmt.Println(v17, v6)
                            }
                            if true {
                                v4 += v4
                            }
                        }
                    }
                    switch (-(fi67)) {
                    case -204:
                        v5 = (v5 & 205)
                    case 2147483647:
                        if true {
                            if false {
                                var v80 uint8 = 0
                                v80 -= v5
                                if (true && true) {
                                    var v81 uint64 = 4294967295
                                    var v82 float64 = (((v6 / ((((1.5 + (v6 + v6)) / v6) - ((-2.25 / 1.0) / v6)) - v6)) / (v6 + (1.5 / ((v6 - 1.5) - (v6 + 0.0))))) + v6)
                                    _ = v81
                                    _ = v82
                                }
                                _ = v80
                            } else {
                                var v83 int8 = 100
                                _ = v83
                            }
                            v84 := v2[v17]
                            v85 := GenS3{F0: 2147483647}
                            v86 := GenS2{F0: v4}
                            _ = v84
                            _ = v85
                            _ = v86
                        }
                        fmt.Println(v1, fi48)
                    case -775:
                        if true {
                            if true {
                                v87 := GenS1{F0: int16(v5), F1: 1}
                                fmt.Println(fi48, v4, v1)
                                var v88 int8 = 42
                                _ = v87
                                _ = v88
                            }
                            if (v17 == v17) {
                                var v89 uint8 = v5
                                var v90 float64 = (v6 / ((3.14159 / v6) / (v6 - 0.0)))
                                _ = v89
                                _ = v90
                            }
                            v11 = (v1 + "a")
                            if len(v3) > 0 {
                                v3[0] = (v17 & 2147483648)
                            }
                            if (v4 > 9223372036854775807) {
                                fmt.Println(fi67, v1)
                                v91 := GenS6{F0: -9223372036854775808}
                                _ = v91
                            }
                            if false {
                                if (false && false) {
                                    var v92 float64 = v6
                                    if (!true) {
                                        switch "" {
                                        case "":
                                            v92 += (0.1 + float64(0.0))
                                        case "a":
                                            v1 = (v11 + "a")
                                        }
                                    }
                                    _ = v92
                                } else {
                                    if (!false) {
                                        v11 = (v14 + "")
                                    }
                                }
                            } else {
                                fmt.Println(v14, v14)
                                var v93 uint64 = 2
                                v94 := GenS2{F0: v4}
                                _ = v93
                                _ = v94
                            }
                        }
                    default:
                        var v95 float64 = v6
                        _ = v95
                    }
                }
                _ = fi67
            }
            _ = len(v2)
        } else {
            v1 = ""
        }
        v5 += v5
        v3 = append(v3, v17)
        _ = fi48
    }
    for fi96 := 0; fi96 < 16; fi96++ {
        switch v1 {
        case "":
            for fi97 := 0; fi97 < 2; fi97++ {
                if false {
                    if true {
                        v98 := v2[(v17 + v17)]
                        if false {
                            v11 = ("a" + (v1 + v14))
                            v11 = v1
                        } else {
                            v17 -= v98
                            if len(v3) > 0 {
                                v3[0] = 2
                            }
                            var v99 float64 = float64((1e10 * ((0.0 * 1.5) * 0.1)))
                            v14 = v14
                            _ = v99
                        }
                        _ = v98
                    }
                    v100 := GenS7{F0: (false && (uint(v17) != genFn0())), F1: (fi96 & (fi97 % ((fi96) | 1)))}
                    switch ("\n\t" + v14) {
                    case "":
                        v101 := GenS5{F0: -429, F1: 320, F2: -1}
                        _ = v101
                    case "foo bar":
                        v5 = (v5 & 0)
                    default:
                        if true {
                            var v102 uint32 = (v17 & v17)
                            v103 := v2[716]
                            _ = v102
                            _ = v103
                        } else {
                            if true {
                                v4 += 838
                            }
                        }
                    }
                    _ = v100
                }
                if true {
                    v1 = "hello"
                    v104 := GenS8{F0: (v1 + (v14 + ((v11 + (v11 + ((v11 + v1) + v11))) + "a")))}
                    _ = v104
                }
                _ = fi97
            }
        case "\n\t":
            if ((v17 >> ((0) & 63)) == (v17 / ((0) | 1))) {
                var v105 int32 = -1
                v1 = "foo bar"
                _ = v105
            }
        case "hello":
            switch fi96 {
            case 0:
                var v106 float64 = 3.14159
                _ = v106
            case 127:
                v107 := GenS2{F0: int64(v5)}
                _ = v107
            case 618:
                var v108 int32 = -1
                _ = v108
            }
        }
        var v109 string = ""
        v110 := GenS5{F0: -1, F1: 1, F2: 42}
        _ = fi96
        _ = v109
        _ = v110
    }
    if len(v3) > 0 {
        v3[0] = (v17 >> ((4294967295) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
