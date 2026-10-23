package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int8; F1 string; F2 int8 }
type GenS8 struct { F0 int8; F1 int8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint32; F1 int; F2 int }
type GenS11 struct { F0 int8 }
func genFn0() bool {
    switch "a" {
    case "":
        for fi6 := 0; fi6 < 16; fi6++ {
            v7 := GenS1{F0: -9223372036854775808, F1: "", F2: 123}
            _ = fi6
            _ = v7
        }
    case "hello":
        v8 := GenS1{F0: 598, F1: "hello", F2: 2147483647}
        _ = v8
    case "a":
        switch ("hello" + "a") {
        case "foo bar":
            if true {
                for fi9 := 0; fi9 < 2; fi9++ {
                    v10 := []int{(fi9 + 9223372036854775807), -2147483648, (fi9 / ((4294967295) | 1))}
                    var v11 uint64 = 22
                    for fi12 := 0; fi12 < 4; fi12++ {
                        fmt.Println(fi12)
                        var v13 int32 = -51
                        v13 += (^(v13))
                        _ = fi12
                        _ = v13
                    }
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
        case "hello":
            v14 := []int{127, 1}
            _ = v14
        case "a":
            v15 := []int{}
            _ = v15
        default:
            var v16 uint8 = 0
            _ = v16
        }
    }
    var v17 float32 = 1.5
    var v18 int16 = -1
    switch "foo bar" {
    case "hello":
        var v19 uint64 = 455
        _ = v19
    default:
        var v20 uint = 4294967295
        _ = v20
    }
    v21 := []int{}
    _ = v17
    _ = v18
    _ = v21
    return ((1 > 42) || (true && false))
}

func main() {
    var v1 string = "hello"
    var v2 uint16 = 2
    v3 := make(map[uint32]int)
    var v4 uint8 = 128
    var v5 uint64 = 960
    v5 += (v5 / ((v5) | 1))
    for fi22 := 0; fi22 < 16; fi22++ {
        if (!true) {
            delete(v3, 4294967295)
            var v23 string = v1
            v24 := v3[921]
            _ = v23
            _ = v24
        }
        v2 += (^(v2))
        switch (fi22 - (fi22 >> ((2) & 63))) {
        case 9223372036854775807:
            if false {
                for fi25 := 0; fi25 < 12; fi25++ {
                    if (-81 == 42) {
                        var v26 int8 = int8(v4)
                        switch fi25 {
                        case 1:
                            if (42 > 0) {
                                var v27 uint16 = v2
                                var v28 int = 9223372036854775807
                                _ = v27
                                _ = v28
                            } else {
                                if false {
                                    v4 += (v4 / (((v4 & uint8(v2))) | 1))
                                    if (1.5 <= 0.0) {
                                        v29 := []int{127}
                                        if false {
                                            var v30 int8 = v26
                                            _ = v30
                                        }
                                        _ = v29
                                    } else {
                                        fmt.Println(v5, v26)
                                        v31 := GenS2{F0: false, F1: v2}
                                        _ = v31
                                    }
                                }
                                var v32 string = "foo bar"
                                _ = v32
                            }
                            if true {
                                var v33 int = (fi22 | fi22)
                                v26 += 19
                                _ = v33
                            } else {
                                v2 += v2
                            }
                            switch (fi25 * fi22) {
                            case -2147483648:
                                v34 := v3[65535]
                                _ = v34
                            case 32768:
                                var v35 int16 = 0
                                v5 += v5
                                var v36 uint8 = (v4 >> ((uint(v4)) & 63))
                                switch v1 {
                                case "":
                                    if false {
                                        var v37 int64 = -1
                                        _ = v37
                                    } else {
                                        v4 -= 2
                                        switch fi25 {
                                        case 128:
                                            if true {
                                                fmt.Println(v35, v26, v2)
                                                var v38 uint64 = 384
                                                v1 = v1
                                                _ = v38
                                            }
                                        case 1:
                                            v2 += v2
                                        case 9223372036854775807:
                                            v1 = (v1 + (v1 + v1))
                                        }
                                    }
                                case "a":
                                    switch (fi22 % (((fi25 * fi22)) | 1)) {
                                    case 506:
                                        var v39 uint32 = 868
                                        _ = v39
                                    case 2147483647:
                                        var v40 uint8 = 42
                                        _ = v40
                                    case -474:
                                        v41 := GenS3{F0: (9223372036854775807 <= 1), F1: 147}
                                        _ = v41
                                    default:
                                        switch fi22 {
                                        case 128:
                                            v42 := GenS2{F0: (false || false), F1: v2}
                                            _ = v42
                                        case 127:
                                            v2 -= (^(v2))
                                        case -2147483648:
                                            v43 := []int{fi25}
                                            _ = v43
                                        default:
                                            v44 := GenS2{F0: true, F1: (v2 + (v2 + 255))}
                                            _ = v44
                                        }
                                    }
                                case "\n\t":
                                    var v45 int = fi22
                                    _ = v45
                                }
                                switch ("foo bar" + v1) {
                                case "hello":
                                    v3[1] = (fi22 - fi22)
                                case "\n\t":
                                    var v46 int = int(v2)
                                    _ = v46
                                default:
                                    v4 = 108
                                }
                                var v47 int64 = -600
                                _ = v35
                                _ = v36
                                _ = v47
                            }
                        case 0:
                            fmt.Println(v1, v1)
                        case -224:
                            v48 := GenS4{F0: "", F1: (!(!(false && false)))}
                            _ = v48
                        default:
                            switch 127 {
                            case 42:
                                var v49 uint = 708
                                _ = v49
                            case -9223372036854775808:
                                v50 := GenS3{F0: false, F1: -1}
                                _ = v50
                            case 4294967295:
                                if false {
                                    switch 128 {
                                    case 127:
                                        v51 := GenS5{F0: 638}
                                        _ = v51
                                    }
                                } else {
                                    var v52 bool = false
                                    v26 = v26
                                    if v52 {
                                        switch (fi22 << ((4294967295) & 63)) {
                                        case 128:
                                            var v53 uint32 = 2147483647
                                            _ = v53
                                        default:
                                            v26 += (v26 ^ v26)
                                        }
                                    }
                                    _ = v52
                                }
                            }
                        }
                        v54 := GenS5{F0: 59}
                        switch v1 {
                        case "":
                            var v55 int32 = 1
                            _ = v55
                        case "\n\t":
                            if (uint64(v2) < v5) {
                                fmt.Println(fi25)
                                fmt.Println(v1, v4)
                                switch 32768 {
                                case -9223372036854775808:
                                    fmt.Println(v26, fi22, v2)
                                case -201:
                                    v56 := []int{(-(fi25)), int(v5)}
                                    v1 = v1
                                    _ = len(v56)
                                    _ = v56
                                }
                            } else {
                                v2 = 2
                                v57 := []int{fi22, fi25, 522, (fi25 + fi22)}
                                _ = v57
                            }
                        case "foo bar":
                            v58 := GenS6{F0: v5}
                            _ = v58
                        default:
                            if true {
                                v26 -= v26
                                var v59 int64 = 0
                                var v60 uint64 = (v5 / ((v5) | 1))
                                _ = v59
                                _ = v60
                            } else {
                                v61 := GenS1{F0: 1, F1: ("foo bar" + (v1 + v1)), F2: 919}
                                _ = v61
                            }
                        }
                        fmt.Println(v4, fi25)
                        _ = v26
                        _ = v54
                    }
                    _ = fi25
                }
            } else {
                if (!false) {
                    for fi62 := 0; fi62 < 4; fi62++ {
                        if (2 == 917) {
                            var v63 float32 = 0.0
                            var v64 int16 = -1
                            _ = v63
                            _ = v64
                        } else {
                            var v65 int = 236
                            if false {
                                v3[2] = fi62
                            }
                            fmt.Println(v65, fi62, v4)
                            _ = v65
                        }
                        _ = fi62
                    }
                    v2 = v2
                }
                v5 = (v5 - (v5 | (v5 % ((v5) | 1))))
            }
        case 4294967295:
            switch "a" {
            case "\n\t":
                for fi66 := 0; fi66 < 13; fi66++ {
                    switch fi66 {
                    case -405:
                        if (int8(v4) < 42) {
                            if false {
                                var v67 uint = 766
                                v2 = v2
                                _ = v67
                            }
                            v4 -= v4
                        } else {
                            var v68 int32 = 210
                            if ((!(false && false)) && (false || true)) {
                                v69 := v3[0]
                                _ = v69
                            }
                            switch (("foo bar" + ("\n\t" + v1)) + (v1 + v1)) {
                            case "":
                                switch (fi22 / ((fi22) | 1)) {
                                case 9223372036854775807:
                                    switch v1 {
                                    case "\n\t":
                                        if false {
                                            v1 = ((v1 + "hello") + "")
                                            v70 := GenS2{F0: true, F1: 380}
                                            if false {
                                                fmt.Println(v68, v2, fi22)
                                            }
                                            fmt.Println(v4)
                                            _ = v70
                                        } else {
                                            var v71 int16 = 42
                                            if false {
                                                switch ((("" + v1) + v1) + "hello") {
                                                case "hello":
                                                    _ = len(v3)
                                                case "":
                                                    var v72 uint = 671
                                                    _ = v72
                                                default:
                                                    var v73 float32 = (0.0 * ((float32(float64(0.0)) + 1.5) * float32(float64(-2.25))))
                                                    switch ("a" + v1) {
                                                    case "\n\t":
                                                        fmt.Println(v1)
                                                    case "foo bar":
                                                        switch "" {
                                                        case "foo bar":
                                                            v4 += 189
                                                        case "\n\t":
                                                            switch "\n\t" {
                                                            case "a":
                                                                switch fi22 {
                                                                case -9223372036854775808:
                                                                    var v74 string = (v1 + "")
                                                                    v3[65535] = 32767
                                                                    v2 += v2
                                                                    v1 = "a"
                                                                    v71 -= -32768
                                                                    _ = v74
                                                                case 127:
                                                                    fmt.Println(v73, v68, fi66)
                                                                default:
                                                                    var v75 bool = true
                                                                    _ = v75
                                                                }
                                                            default:
                                                                v1 = v1
                                                            }
                                                            v76 := v3[1]
                                                            _ = v76
                                                        case "":
                                                            v4 = v4
                                                        }
                                                    case "hello":
                                                        v77 := v3[2147483648]
                                                        _ = v77
                                                    default:
                                                        if (!false) {
                                                            v78 := []int{fi66, fi66, fi66}
                                                            v78 = append(v78, fi22)
                                                            var v79 int64 = int64(v2)
                                                            switch v1 {
                                                            case "a":
                                                                switch "foo bar" {
                                                                case "a":
                                                                    fmt.Println(v68, v4)
                                                                    v80 := GenS1{F0: int64(v71), F1: v1, F2: 65535}
                                                                    _ = v80
                                                                default:
                                                                    v79 -= -941
                                                                }
                                                            case "":
                                                                if len(v78) >= 2 {
                                                                    v78 = v78[1:len(v78)]
                                                                }
                                                            case "\n\t":
                                                                var v81 uint32 = 65535
                                                                _ = v81
                                                            default:
                                                                v82 := GenS7{F0: 0, F1: "", F2: -82}
                                                                _ = v82
                                                            }
                                                            v79 = 1
                                                            _ = v78
                                                            _ = v79
                                                        } else {
                                                            if (1e10 >= v73) {
                                                                v83 := GenS7{F0: -128, F1: v1, F2: -128}
                                                                _ = v83
                                                            }
                                                        }
                                                    }
                                                    _ = v73
                                                }
                                                switch (fi66 >> ((uint(v4)) & 63)) {
                                                case -9223372036854775808:
                                                    switch v1 {
                                                    case "":
                                                        v84 := GenS3{F0: true, F1: -829}
                                                        _ = v84
                                                    default:
                                                        v85 := GenS8{F0: 100, F1: 81}
                                                        _ = v85
                                                    }
                                                case 4294967295:
                                                    var v86 int32 = -2
                                                    _ = v86
                                                default:
                                                    v87 := GenS9{F0: 0}
                                                    _ = v87
                                                }
                                                var v88 uint16 = v2
                                                if ((!false) && false) {
                                                    v71 = int16(v68)
                                                    v2 = v88
                                                }
                                                _ = len(v3)
                                                _ = v88
                                            } else {
                                                v3[2] = (fi66 * fi22)
                                                switch v1 {
                                                case "\n\t":
                                                    var v89 uint32 = 4294967295
                                                    _ = v89
                                                case "":
                                                    v90 := []int{fi22}
                                                    _ = v90
                                                case "hello":
                                                    fmt.Println(v5)
                                                default:
                                                    var v91 int32 = (v68 ^ 1)
                                                    _ = v91
                                                }
                                            }
                                            v2 += v2
                                            _ = v71
                                        }
                                    default:
                                        _ = len(v3)
                                    }
                                default:
                                    v4 += v4
                                }
                            case "hello":
                                v2 = v2
                            case "a":
                                v92 := GenS10{F0: 2, F1: (fi22 | -642), F2: fi22}
                                _ = v92
                            default:
                                v93 := GenS11{F0: -128}
                                _ = v93
                            }
                            _ = v68
                        }
                    case 128:
                        var v94 uint = 0
                        _ = v94
                    case -2147483648:
                        if true {
                            switch (fi22 | fi22) {
                            case 32767:
                                v1 = "a"
                            }
                            v4 -= v4
                        } else {
                            v5 = v5
                        }
                    }
                    v95 := v3[892]
                    v96 := v3[600]
                    v97 := []int{0, (fi22 + v96)}
                    _ = fi66
                    _ = v95
                    _ = v96
                    _ = v97
                }
            case "a":
                switch v1 {
                case "\n\t":
                    var v98 float64 = float64((1.5 * -2.25))
                    _ = v98
                case "a":
                    if false {
                        var v99 int8 = 0
                        _ = v99
                    }
                case "":
                    var v100 uint16 = v2
                    _ = v100
                }
            default:
                delete(v3, 2)
            }
        case 185:
            v2 = v2
        }
        _ = fi22
    }
    v4 -= 1
    var v101 uint = 757
    v102 := GenS6{F0: 42}
    v1 = v1
    var v103 int64 = -1
    var v104 int8 = 1
    var v105 uint8 = 0
    v101 = v101
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v103:", v103)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v101
    _ = v102
    _ = v103
    _ = v104
    _ = v105
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
