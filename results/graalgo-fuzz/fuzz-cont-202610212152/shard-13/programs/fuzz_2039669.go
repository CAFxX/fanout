package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int8; F1 int64; F2 int16 }
type GenS5 struct { F0 int; F1 uint32; F2 int64 }
type GenS6 struct { F0 uint; F1 int8 }
type GenS7 struct { F0 uint32; F1 int32 }
type GenS8 struct { F0 uint8; F1 string }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint8; F1 uint }
type GenS11 struct { F0 int32; F1 int8 }
type GenS12 struct { F0 uint8; F1 int16; F2 uint }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 uint16; F1 int }
type GenS15 struct { F0 int32 }
func genFn0() int64 {
    if (0 == -840) {
        var v5 uint64 = 852
        _ = v5
    }
    var v6 string = "hello"
    if true {
        var v7 uint16 = 32767
        _ = v7
    } else {
        switch v6 {
        case "a":
            var v8 int32 = -1
            _ = v8
        case "\n\t":
            switch 2147483647 {
            case 32768:
                switch v6 {
                case "foo bar":
                    var v9 int32 = 1
                    _ = v9
                case "":
                    var v10 uint64 = 79
                    _ = v10
                case "\n\t":
                    v11 := GenS1{F0: 877, F1: -1}
                    _ = v11
                default:
                    v12 := []int{}
                    _ = v12
                }
            case 128:
                switch -719 {
                case 2147483647:
                    var v13 int32 = -897
                    _ = v13
                case 1:
                    v6 = ("\n\t" + v6)
                case -2147483648:
                    var v14 int = 1
                    _ = v14
                default:
                    var v15 uint = 0
                    _ = v15
                }
            }
        }
        for fi16 := 0; fi16 < 7; fi16++ {
            v17 := GenS2{F0: (!false), F1: 357}
            var v18 int8 = -1
            v6 = v6
            _ = fi16
            _ = v17
            _ = v18
        }
        if true {
            var v19 string = (v6 + ("foo bar" + "\n\t"))
            var v20 string = "foo bar"
            fmt.Println(v20)
            _ = v19
            _ = v20
        } else {
            v21 := []int{42}
            for fi22 := 0; fi22 < 20; fi22++ {
                if false {
                    for fi23 := 0; fi23 < 10; fi23++ {
                        var v24 float64 = float64(1e10)
                        v6 = v6
                        _ = fi23
                        _ = v24
                    }
                    v25 := GenS1{F0: 2, F1: 100}
                    _ = v25
                }
                var v26 int64 = -1
                _ = fi22
                _ = v26
            }
            _ = v21
        }
    }
    v27 := []int{9223372036854775807, 128, 42, 42}
    _ = v27
    _ = v6
    return -1
}

func main() {
    var v1 string = "a"
    var v2 uint8 = 252
    v3 := []int{42, 2147483647}
    v4 := []int{}
    switch v1 {
    case "a":
        v2 = v2
    case "hello":
        var v28 uint8 = 0
        _ = v28
    default:
        for fi29 := 0; fi29 < 16; fi29++ {
            if (!true) {
                for fi30 := 0; fi30 < 8; fi30++ {
                    v31 := GenS3{F0: 65535, F1: 186}
                    fmt.Println(v2, fi30, fi29)
                    _ = fi30
                    _ = v31
                }
                var v32 uint16 = 115
                v2 += v2
                _ = v32
            } else {
                v1 = (v1 + "foo bar")
                var v33 uint64 = 18446744073709551615
                _ = v33
            }
            _ = fi29
        }
    }
    _ = len(v4)
    var v34 int32 = -2147483648
    _ = len(v4)
    v35 := GenS1{F0: 0, F1: 1}
    var v36 int32 = v34
    switch 127 {
    case 42:
        v4 = append(v4, -1)
    case 30:
        var v37 float32 = (0.0 + ((1e10 - 3.14159) * 3.14159))
        _ = v37
    }
    v38 := GenS1{F0: 148, F1: -128}
    v39 := GenS4{F0: -120, F1: 9223372036854775807, F2: -1}
    if (1.5 == 3.14159) {
        var v40 float64 = (float64(-2.25) - ((-2.25 / 0.1) + ((((float64(1.5) * 3.14159) / (1.5 / (1.5 * (-2.25 / ((3.14159 + 0.0) + ((float64(0.1) + (1e10 + (3.14159 + (float64(0.1) * (3.14159 + float64((((1e10 * (-2.25 * 3.14159)) - (0.0 - -2.25)) * (0.0 * 1.5)))))))) + ((-2.25 + (0.0 * (((1e10 - -2.25) / 0.1) / 1.5))) + (1.5 / 1.0)))))))) / -2.25) - 0.1)))
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if false {
            for fi41 := 0; fi41 < 14; fi41++ {
                _ = len(v4)
                _ = fi41
            }
            if false {
                for fi42 := 0; fi42 < 1; fi42++ {
                    v43 := GenS5{F0: 42, F1: 229, F2: 1}
                    _ = fi42
                    _ = v43
                }
            }
        } else {
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
            var v44 int8 = -128
            _ = v44
        }
        _ = v40
    } else {
        v45 := GenS1{F0: 1, F1: -128}
        var v46 uint16 = 955
        if (true && false) {
            var v47 int32 = (v34 ^ v34)
            _ = v47
        } else {
            v48 := GenS6{F0: 1, F1: 73}
            _ = v48
        }
        _ = v45
        _ = v46
    }
    var v49 uint16 = 42
    for fi50 := 0; fi50 < 5; fi50++ {
        switch 32768 {
        case 4294967295:
            if false {
                v34 -= v34
                switch 1 {
                case 128:
                    _ = len(v3)
                case -925:
                    var v51 uint64 = 779
                    _ = v51
                }
                if (true && true) {
                    switch ("a" + (("foo bar" + v1) + ("hello" + v1))) {
                    case "hello":
                        if true {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            switch (v1 + "a") {
                            case "":
                                _ = len(v3)
                            case "hello":
                                v2 -= (v2 * (^(v2)))
                            }
                            var v52 int = -994
                            _ = v52
                        } else {
                            if len(v4) > 0 {
                                v4[0] = (fi50 | fi50)
                            }
                            if false {
                                fmt.Println(v2)
                            } else {
                                for fi53 := 0; fi53 < 20; fi53++ {
                                    v36 += v36
                                    var v54 int8 = -128
                                    switch fi53 {
                                    case 128:
                                        v49 += (v49 | (v49 << ((42) & 63)))
                                    case 1:
                                        fmt.Println(v36)
                                    default:
                                        var v55 bool = true
                                        _ = v55
                                    }
                                    v34 = v36
                                    v56 := GenS7{F0: 2147483647, F1: 2147483647}
                                    _ = fi53
                                    _ = v54
                                    _ = v56
                                }
                                _ = len(v3)
                            }
                        }
                    case "foo bar":
                        for fi57 := 0; fi57 < 15; fi57++ {
                            if (false && ((!false) && false)) {
                                fmt.Println(fi50, v34)
                            } else {
                                var v58 int32 = v36
                                switch "" {
                                case "foo bar":
                                    switch -143 {
                                    case 847:
                                        v59 := GenS1{F0: 464, F1: 1}
                                        _ = v59
                                    }
                                    fmt.Println(v2, fi57, fi57)
                                    if true {
                                        switch "hello" {
                                        case "foo bar":
                                            fmt.Println(v49, v2)
                                        case "a":
                                            if ((0.0 + (3.14159 + (0.1 + 3.14159))) != 1.5) {
                                                v49 -= v49
                                            } else {
                                                v34 -= (v34 | -2147483648)
                                                var v60 string = (("hello" + (v1 + (v1 + v1))) + v1)
                                                _ = v60
                                            }
                                        case "":
                                            if true {
                                                v61 := GenS4{F0: 105, F1: 42, F2: -32768}
                                                var v62 uint8 = v2
                                                _ = v61
                                                _ = v62
                                            }
                                        default:
                                            v3 = append(v3, fi57)
                                        }
                                        switch (^(fi57)) {
                                        case 665:
                                            if false {
                                                if ((1 != 42) || (0.0 > 0.0)) {
                                                    fmt.Println(v58, fi57)
                                                    v34 = (v58 & v34)
                                                    var v63 int64 = -1
                                                    var v64 uint8 = (v2 - (v2 << ((1) & 63)))
                                                    if true {
                                                        v65 := GenS8{F0: v2, F1: "hello"}
                                                        _ = v65
                                                    } else {
                                                        v3 = append(v3, (^(fi57)))
                                                        fmt.Println(v34, fi50)
                                                        v66 := GenS9{F0: 127}
                                                        _ = v66
                                                    }
                                                    v67 := GenS9{F0: 127}
                                                    _ = v63
                                                    _ = v64
                                                    _ = v67
                                                } else {
                                                    switch (-(fi50)) {
                                                    case 42:
                                                        v36 += (^(v34))
                                                    }
                                                }
                                                var v68 uint = 0
                                                var v69 int8 = 100
                                                _ = v68
                                                _ = v69
                                            } else {
                                                fmt.Println(v2, fi57, v34)
                                            }
                                            v36 = (v58 | v34)
                                            if (2 < 675) {
                                                v1 = ("\n\t" + (v1 + "hello"))
                                            }
                                        case 9223372036854775807:
                                            switch v1 {
                                            case "foo bar":
                                                fmt.Println(fi57, fi50)
                                            case "":
                                                v70 := GenS1{F0: 4294967295, F1: -51}
                                                _ = v70
                                            default:
                                                if false {
                                                    v34 = v58
                                                    if false {
                                                        v2 -= (v2 >> ((255) & 63))
                                                    }
                                                    v49 -= 125
                                                }
                                            }
                                        }
                                    } else {
                                        v58 -= v58
                                        v1 = (v1 + ("foo bar" + (v1 + v1)))
                                        fmt.Println(fi50, v58, v58)
                                    }
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v71 int16 = 1
                                        _ = v71
                                    case "hello":
                                        v3 = append(v3, fi50)
                                    default:
                                        v4 = append(v4, (^(fi57)))
                                    }
                                    fmt.Println(v36)
                                case "hello":
                                    v36 = v36
                                    var v72 bool = (0 > (fi50 % ((fi57) | 1)))
                                    _ = v72
                                case "a":
                                    v73 := GenS2{F0: (-8 == 0), F1: 86}
                                    _ = v73
                                default:
                                    switch fi57 {
                                    case 435:
                                        var v74 int64 = genFn0()
                                        _ = v74
                                    case 4294967295:
                                        switch fi57 {
                                        case 2147483647:
                                            v75 := GenS5{F0: 2147483647, F1: 2, F2: 42}
                                            _ = v75
                                        case -2147483648:
                                            _ = len(v3)
                                        case 9223372036854775807:
                                            _ = len(v4)
                                        default:
                                            v3 = append(v3, int(v49))
                                        }
                                    case 32767:
                                        var v76 float64 = 0.1
                                        _ = v76
                                    default:
                                        v77 := GenS10{F0: (v2 - 1), F1: 42}
                                        _ = v77
                                    }
                                }
                                _ = v58
                            }
                            _ = fi57
                        }
                    case "":
                        if (9223372036854775807 <= 0) {
                            v78 := GenS1{F0: 447, F1: -31}
                            _ = v78
                        }
                    }
                } else {
                    fmt.Println(v36)
                    for fi79 := 0; fi79 < 18; fi79++ {
                        v36 = v34
                        var v80 float32 = 3.14159
                        v36 = 0
                        _ = fi79
                        _ = v80
                    }
                }
            } else {
                var v81 int8 = 0
                for fi82 := 0; fi82 < 4; fi82++ {
                    v34 = (v34 >> ((165) & 63))
                    fmt.Println(v1, v1)
                    v81 += v81
                    v34 -= -2147483648
                    var v83 int16 = -122
                    _ = fi82
                    _ = v83
                }
                var v84 uint64 = 371
                _ = v81
                _ = v84
            }
        case 233:
            fmt.Println(v49)
        default:
            fmt.Println(fi50, v34, v2)
        }
        if false {
            v2 += (v2 % ((1) | 1))
        } else {
            v36 += v34
            var v85 string = v1
            for fi86 := 0; fi86 < 16; fi86++ {
                v87 := GenS2{F0: false, F1: 0}
                v2 = v2
                _ = fi86
                _ = v87
            }
            _ = v85
        }
        v88 := GenS10{F0: (v2 << ((246) & 63)), F1: 42}
        _ = fi50
        _ = v88
    }
    v49 -= (^(v49))
    for fi89 := 0; fi89 < 7; fi89++ {
        if ((((3.14159 - float32((0.1 * -2.25))) / (1e10 / 3.14159)) - 1e10) >= ((3.14159 / -2.25) - (1e10 * 1e10))) {
            var v90 string = "a"
            _ = v90
        }
        if (!false) {
            if ((v2 == 127) && (2 < (v2 | (v2 >> ((18446744073709551615) & 63))))) {
                for fi91 := 0; fi91 < 16; fi91++ {
                    v36 += (v36 << ((654) & 63))
                    _ = fi91
                }
                var v92 bool = true
                _ = v92
            } else {
                v1 = "\n\t"
                switch (fi89 ^ fi89) {
                case 998:
                    var v93 uint8 = (^(v2))
                    _ = v93
                }
            }
        } else {
            if (-478 >= (fi89 / ((fi89) | 1))) {
                switch v1 {
                case "foo bar":
                    var v94 int8 = 1
                    _ = v94
                case "hello":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                v49 += (v49 ^ (v49 << ((255) & 63)))
            } else {
                if (false && true) {
                    switch "\n\t" {
                    case "\n\t":
                        switch ((("" + v1) + (v1 + v1)) + v1) {
                        case "hello":
                            switch v1 {
                            case "a":
                                v1 = v1
                            case "hello":
                                for fi95 := 0; fi95 < 16; fi95++ {
                                    fmt.Println(v36, fi89, v1)
                                    v36 = v36
                                    var v96 float64 = (((((1e10 / (1.5 * -2.25)) / -2.25) - 0.0) * 3.14159) / 1.5)
                                    _ = fi95
                                    _ = v96
                                }
                            case "\n\t":
                                v97 := GenS11{F0: (v34 | v34), F1: -128}
                                _ = v97
                            default:
                                var v98 uint = 230
                                _ = v98
                            }
                        case "":
                            v99 := GenS12{F0: 127, F1: 996, F2: 65535}
                            _ = v99
                        default:
                            var v100 uint32 = 2147483647
                            _ = v100
                        }
                    case "a":
                        for fi101 := 0; fi101 < 7; fi101++ {
                            fmt.Println(v34, v49)
                            var v102 int8 = int8(v49)
                            _ = fi101
                            _ = v102
                        }
                    }
                    switch ("foo bar" + v1) {
                    case "a":
                        v103 := GenS13{F0: 61}
                        _ = v103
                    case "":
                        for fi104 := 0; fi104 < 12; fi104++ {
                            var v105 int64 = 1
                            v1 = "foo bar"
                            var v106 float32 = (1.5 + 1e10)
                            _ = fi104
                            _ = v105
                            _ = v106
                        }
                    }
                    if false {
                        v36 += v34
                        if true {
                            v107 := GenS5{F0: -9223372036854775808, F1: 8, F2: 0}
                            _ = v107
                        }
                        if (!true) {
                            switch v1 {
                            case "a":
                                var v108 uint = 65535
                                _ = v108
                            }
                        } else {
                            if ((v1 + v1) == "foo bar") {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v109 int16 = -32768
                                _ = v109
                            }
                            var v110 float32 = (3.14159 - -2.25)
                            switch "" {
                            case "hello":
                                var v111 uint64 = 1
                                _ = v111
                            case "":
                                if true {
                                    if false {
                                        if true {
                                            for fi112 := 0; fi112 < 18; fi112++ {
                                                switch 42 {
                                                case 9223372036854775807:
                                                    var v113 int8 = -110
                                                    _ = v113
                                                }
                                                v3 = append(v3, fi89)
                                                _ = len(v4)
                                                _ = fi112
                                            }
                                        }
                                        if (0.1 >= 0.1) {
                                            fmt.Println(v49)
                                        }
                                        v49 += v49
                                    } else {
                                        v2 += (v2 % ((v2) | 1))
                                        v114 := GenS14{F0: (v49 >> ((130) & 63)), F1: (fi89 & -9223372036854775808)}
                                        var v115 int16 = -1
                                        _ = v114
                                        _ = v115
                                    }
                                } else {
                                    var v116 string = v1
                                    _ = v116
                                }
                            default:
                                for fi117 := 0; fi117 < 11; fi117++ {
                                    v36 = -824
                                    if false {
                                        var v118 float32 = v110
                                        _ = v118
                                    }
                                    v119 := GenS11{F0: (v34 - v36), F1: -106}
                                    _ = fi117
                                    _ = v119
                                }
                            }
                            _ = v110
                        }
                    } else {
                        var v120 int16 = -1
                        if false {
                            for fi121 := 0; fi121 < 9; fi121++ {
                                v34 = int32(v120)
                                _ = fi121
                            }
                            if true {
                                _ = len(v4)
                            }
                            if len(v4) > 0 {
                                v4[0] = 9223372036854775807
                            }
                        }
                        _ = v120
                    }
                }
                v122 := GenS15{F0: (v36 % ((v34) | 1))}
                _ = v122
            }
        }
        _ = fi89
    }
    var v123 float64 = (((1.5 / 1.0) * 0.1) + ((1.5 * ((-2.25 - ((-2.25 * 0.1) + 0.0)) + (0.0 + 0.0))) * float64(float32(((0.0 - 0.0) + float64(-2.25))))))
    var v124 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v123:", v123)
    fmt.Println("v124:", v124)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v123
    _ = v124
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v38
    _ = v39
    _ = v4
    _ = v49
}
