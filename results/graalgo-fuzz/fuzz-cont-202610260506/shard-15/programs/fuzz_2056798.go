package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int8; F2 int64 }
type GenS4 struct { F0 uint8; F1 bool; F2 int64 }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 int64; F1 int }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS10 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS11 struct { F0 uint32; F1 int64; F2 int }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 bool; F1 int64; F2 uint }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint16; F1 uint32; F2 int64 }
func genFn0() int8 {
    var v6 bool = (730 > 65535)
    var v7 int = 2147483647
    v8 := GenS1{F0: 4294967295}
    _ = v6
    _ = v7
    _ = v8
    return 127
}
func genFn1() int64 {
    v9 := []int{2147483647, 1}
    var v10 string = "foo bar"
    _ = v10
    _ = v9
    return 42
}

func main() {
    var v1 float64 = (0.1 - 1.5)
    v2 := []int{128, -936}
    var v3 bool = false
    v4 := make(map[int]bool)
    var v5 uint16 = 0
    if (1 <= 6) {
        switch -9223372036854775808 {
        case 32767:
            v5 = (v5 % ((v5) | 1))
        case -9223372036854775808:
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 float32 = float32(v1)
                v4[(fi11 << ((uint(v5)) & 63))] = true
                var v13 int32 = -914
                _ = fi11
                _ = v12
                _ = v13
            }
        default:
            switch "foo bar" {
            case "hello":
                switch "hello" {
                case "a":
                    switch "\n\t" {
                    case "\n\t":
                        v14 := GenS2{F0: (v5 | (^(v5))), F1: 556, F2: 42}
                        _ = v14
                    case "foo bar":
                        for fi15 := 0; fi15 < 6; fi15++ {
                            _ = len(v2)
                            v5 = (v5 >> ((1) & 63))
                            v16 := GenS1{F0: 4294967295}
                            _ = fi15
                            _ = v16
                        }
                    case "":
                        v17 := GenS1{F0: 4294967295}
                        _ = v17
                    }
                }
            case "":
                if (!v3) {
                    v5 = (v5 % ((42) | 1))
                    v18 := GenS2{F0: v5, F1: 343, F2: 965}
                    _ = v18
                } else {
                    var v19 string = "foo bar"
                    _ = v19
                }
            default:
                v3 = (float32(v1) != (1.5 / -2.25))
            }
        }
        v3 = (18446744073709551615 == 0)
    }
    v5 -= (v5 / (((v5 - v5)) | 1))
    v20 := GenS1{F0: 78}
    for fi21 := 0; fi21 < 20; fi21++ {
        switch 919 {
        case 589:
            switch fi21 {
            case -9223372036854775808:
                var v22 int8 = 1
                _ = v22
            }
        case 984:
            v3 = ((0.0 - (1.5 * (1.5 / 1e10))) != (3.14159 + 1.5))
        case 226:
            if (1 <= 0) {
                v3 = (genFn0() >= genFn0())
                fmt.Println(v5, fi21, v5)
            } else {
                var v23 bool = v3
                v24 := GenS1{F0: 255}
                switch fi21 {
                case 4294967295:
                    var v25 uint16 = v5
                    _ = v25
                case 128:
                    var v26 uint = 0
                    _ = v26
                }
                _ = v23
                _ = v24
            }
        }
        fmt.Println(v3, v3)
        v27 := GenS1{F0: 1}
        _ = fi21
        _ = v27
    }
    v5 += v5
    if false {
        fmt.Println(v5, v1)
    } else {
        v4[-681] = (115 < 4294967295)
        if len(v2) > 0 {
            v2[0] = 42
        }
        if false {
            v28 := GenS2{F0: (^(v5)), F1: -1, F2: 4294967295}
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 4; fi29++ {
        v3 = v3
        _ = fi29
    }
    for fi30 := 0; fi30 < 12; fi30++ {
        for fi31 := 0; fi31 < 17; fi31++ {
            var v32 uint64 = 0
            v32 = (v32 & v32)
            if v3 {
                _ = len(v2)
                v4[fi31] = false
                if ((419 > -1) && (v3 && true)) {
                    var v33 int8 = -78
                    switch "\n\t" {
                    case "a":
                        _ = len(v4)
                    case "":
                        switch (fi30 - fi30) {
                        case 0:
                            v3 = v3
                        case 2147483647:
                            if (!true) {
                                v32 -= v32
                                v4[fi30] = (!v3)
                            }
                        }
                    case "hello":
                        delete(v4, fi31)
                    default:
                        if (true || v3) {
                            v34 := GenS1{F0: 65535}
                            fmt.Println(v5)
                            fmt.Println(fi31)
                            v33 -= 127
                            _ = v34
                        }
                    }
                    _ = v33
                } else {
                    v3 = (217 != 128)
                    v35 := GenS3{F0: 575, F1: 115, F2: 9223372036854775807}
                    _ = v35
                }
                v36 := GenS1{F0: 610}
                _ = v36
            }
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    if v3 {
        for fi37 := 0; fi37 < 8; fi37++ {
            v38 := GenS4{F0: 66, F1: ((v1 - v1) < (1e10 - float64(1.5))), F2: 0}
            _ = len(v4)
            var v39 bool = v3
            _ = fi37
            _ = v38
            _ = v39
        }
        if v3 {
            for fi40 := 0; fi40 < 5; fi40++ {
                v3 = false
                _ = fi40
            }
            var v41 int32 = 42
            _ = v41
        } else {
            switch ((("foo bar" + (("a" + ("foo bar" + "foo bar")) + ("\n\t" + (("" + ("foo bar" + "\n\t")) + "a")))) + ("hello" + "")) + (("a" + "a") + ("\n\t" + "\n\t"))) {
            case "":
                for fi42 := 0; fi42 < 10; fi42++ {
                    v43 := GenS4{F0: 128, F1: false, F2: -273}
                    v44 := GenS3{F0: 18446744073709551615, F1: -50, F2: 505}
                    v45 := GenS5{F0: -173, F1: 36}
                    _ = fi42
                    _ = v43
                    _ = v44
                    _ = v45
                }
            case "a":
                if true {
                    switch 2147483647 {
                    case -2147483648:
                        for fi46 := 0; fi46 < 18; fi46++ {
                            for fi47 := 0; fi47 < 10; fi47++ {
                                v1 += v1
                                _ = fi47
                            }
                            _ = fi46
                        }
                    case 32767:
                        if (-387 != 1) {
                            if (102 >= 42) {
                                v48 := GenS4{F0: 2, F1: (float32(float64(((((0.0 / 1e10) - 0.0) / ((0.1 + 0.1) + 3.14159)) / float32(float64((-2.25 * 0.1)))))) < float32((float64(1e10) + v1))), F2: -9223372036854775808}
                                if v3 {
                                    v2 = append(v2, 0)
                                }
                                _ = v48
                            } else {
                                if (((42 >= 42) && ((((1.5 + (1.5 * 1.5)) - 0.0) - 0.0) > (float32((((0.1 - (3.14159 * (float64((-2.25 + -2.25)) * ((0.1 + 1e10) * ((3.14159 + 0.0) + v1))))) * v1) + ((float64((0.1 + (((3.14159 + (0.0 + 1e10)) + 3.14159) / 0.1))) - -2.25) / -2.25))) / 3.14159))) && (!(!(!false)))) {
                                    delete(v4, 2147483647)
                                    v3 = v3
                                } else {
                                    if true {
                                        if (!false) {
                                            var v49 bool = false
                                            var v50 int = 42
                                            v51 := GenS6{F0: genFn1(), F1: 0}
                                            _ = v49
                                            _ = v50
                                            _ = v51
                                        } else {
                                            if v3 {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                switch (("foo bar" + "hello") + (("a" + ("foo bar" + "hello")) + "a")) {
                                                case "hello":
                                                    v1 = ((1e10 - -2.25) / 1.0)
                                                case "\n\t":
                                                    for fi52 := 0; fi52 < 2; fi52++ {
                                                        for fi53 := 0; fi53 < 16; fi53++ {
                                                            fmt.Println(fi52, fi52)
                                                            v3 = (2 < 4294967295)
                                                            if v3 {
                                                                var v54 uint8 = 21
                                                                _ = v54
                                                            }
                                                            _ = fi53
                                                        }
                                                        _ = fi52
                                                    }
                                                case "a":
                                                    var v55 int = 307
                                                    _ = v55
                                                }
                                            }
                                            for fi56 := 0; fi56 < 2; fi56++ {
                                                if false {
                                                    switch fi56 {
                                                    case -788:
                                                        for fi57 := 0; fi57 < 10; fi57++ {
                                                            if v3 {
                                                                _ = len(v2)
                                                                var v58 uint32 = 0
                                                                fmt.Println(v3)
                                                                if ((3.14159 / 1e10) > 1e10) {
                                                                    var v59 uint16 = (v5 << ((255) & 63))
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    var v60 string = "hello"
                                                                    v5 = 65535
                                                                    if false {
                                                                        var v61 int64 = -407
                                                                        switch v60 {
                                                                        case "\n\t":
                                                                            v5 += (v5 << ((4294967295) & 63))
                                                                        case "a":
                                                                            if len(v2) > 0 {
                                                                                v2[0] = 9223372036854775807
                                                                            }
                                                                        case "foo bar":
                                                                            var v62 int64 = -9223372036854775808
                                                                            _ = v62
                                                                        default:
                                                                            var v63 uint16 = v59
                                                                            _ = v63
                                                                        }
                                                                        _ = v61
                                                                    }
                                                                    _ = v59
                                                                    _ = v60
                                                                } else {
                                                                    var v64 uint = 255
                                                                    v1 = v1
                                                                    v1 -= (3.14159 / 1.0)
                                                                    _ = v64
                                                                }
                                                                var v65 bool = v3
                                                                _ = v58
                                                                _ = v65
                                                            } else {
                                                                v3 = v3
                                                                if (!v3) {
                                                                    var v66 uint = 2
                                                                    fmt.Println(fi56, fi57)
                                                                    var v67 int = -1
                                                                    _ = v66
                                                                    _ = v67
                                                                }
                                                                var v68 uint8 = 127
                                                                _ = v68
                                                            }
                                                            var v69 int64 = 1
                                                            _ = fi57
                                                            _ = v69
                                                        }
                                                    default:
                                                        fmt.Println(v1)
                                                    }
                                                } else {
                                                    if v3 {
                                                        var v70 int8 = 1
                                                        fmt.Println(fi56, v3)
                                                        _ = v70
                                                    }
                                                    for fi71 := 0; fi71 < 20; fi71++ {
                                                        if true {
                                                            var v72 int32 = 408
                                                            if (int32(fi71) == (v72 ^ 42)) {
                                                                var v73 int64 = 238
                                                                _ = v73
                                                            } else {
                                                                _ = len(v4)
                                                                var v74 int32 = -1
                                                                v75 := GenS6{F0: 1, F1: fi56}
                                                                _ = v74
                                                                _ = v75
                                                            }
                                                            var v76 uint16 = v5
                                                            v77 := GenS7{F0: 275}
                                                            _ = v72
                                                            _ = v76
                                                            _ = v77
                                                        } else {
                                                            v78 := GenS8{F0: "a"}
                                                            fmt.Println(fi56, v5, v5)
                                                            v4[fi56] = v3
                                                            v79 := GenS9{F0: -550, F1: 0, F2: (v5 * v5)}
                                                            fmt.Println(v5, fi56, v3)
                                                            _ = v78
                                                            _ = v79
                                                        }
                                                        _ = fi71
                                                    }
                                                    var v80 uint64 = 4294967295
                                                    _ = v80
                                                }
                                                switch fi56 {
                                                case 42:
                                                    for fi81 := 0; fi81 < 1; fi81++ {
                                                        switch "a" {
                                                        case "\n\t":
                                                            var v82 int = fi56
                                                            _ = v82
                                                        }
                                                        fmt.Println(v3, fi81, fi56)
                                                        var v83 int64 = -1
                                                        _ = fi81
                                                        _ = v83
                                                    }
                                                }
                                                _ = fi56
                                            }
                                        }
                                        if (18446744073709551615 != 1) {
                                            v3 = (25 != 127)
                                        }
                                        v84 := GenS10{F0: 42, F1: v5, F2: 0}
                                        _ = v84
                                    }
                                }
                                var v85 uint = 133
                                _ = v85
                            }
                            var v86 uint8 = 178
                            _ = v86
                        }
                    default:
                        v87 := v4[280]
                        _ = v87
                    }
                    if (0 >= 0) {
                        _ = len(v4)
                        var v88 int64 = -1
                        v89 := GenS11{F0: 65535, F1: (v88 | v88), F2: 42}
                        _ = v88
                        _ = v89
                    }
                    v1 = float64(float32(v1))
                } else {
                    for fi90 := 0; fi90 < 4; fi90++ {
                        for fi91 := 0; fi91 < 15; fi91++ {
                            fmt.Println(fi90, fi91, fi90)
                            fmt.Println(v1, v5)
                            v92 := GenS7{F0: 1}
                            _ = fi91
                            _ = v92
                        }
                        switch (^(fi90)) {
                        case 42:
                            var v93 int64 = 1
                            _ = v93
                        case 2147483647:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        if (1 < 0) {
                            switch "foo bar" {
                            case "foo bar":
                                switch "" {
                                case "":
                                    v5 = (^(v5))
                                case "\n\t":
                                    var v94 uint8 = 2
                                    _ = v94
                                }
                            case "":
                                switch ("foo bar" + ("\n\t" + "a")) {
                                case "":
                                    var v95 int8 = 42
                                    _ = v95
                                case "a":
                                    v1 += v1
                                }
                            default:
                                v3 = v3
                            }
                        } else {
                            if (!((!v3) || (!v3))) {
                                v96 := GenS7{F0: 1}
                                if (255 != 2) {
                                    var v97 uint8 = 0
                                    switch "a" {
                                    case "a":
                                        switch ("hello" + ("hello" + "hello")) {
                                        case "":
                                            switch (fi90 | (fi90 << ((4294967295) & 63))) {
                                            case 2147483647:
                                                fmt.Println(v5, v3)
                                            case -1:
                                                v98 := GenS12{F0: 1}
                                                _ = v98
                                            default:
                                                v99 := v4[fi90]
                                                _ = v99
                                            }
                                        case "\n\t":
                                            var v100 uint16 = v5
                                            _ = v100
                                        case "foo bar":
                                            switch "" {
                                            case "\n\t":
                                                v101 := GenS6{F0: -9223372036854775808, F1: fi90}
                                                _ = v101
                                            }
                                        default:
                                            var v102 string = ""
                                            _ = v102
                                        }
                                    case "foo bar":
                                        if v3 {
                                            for fi103 := 0; fi103 < 2; fi103++ {
                                                v104 := GenS13{F0: (1.5 == 0.1), F1: 1, F2: 1}
                                                var v105 float32 = (((((0.1 / 0.1) + (0.1 - 3.14159)) / 3.14159) / 1e10) * (-2.25 / float32(v1)))
                                                if len(v2) > 0 {
                                                    v2[0] = 0
                                                }
                                                _ = fi103
                                                _ = v104
                                                _ = v105
                                            }
                                        } else {
                                            var v106 int64 = 9223372036854775807
                                            var v107 string = ((("hello" + "a") + (("" + (("" + ("\n\t" + "")) + "foo bar")) + "")) + "")
                                            _ = v106
                                            _ = v107
                                        }
                                    case "\n\t":
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v108 uint32 = 4294967295
                                            _ = v108
                                        }
                                    default:
                                        if false {
                                            var v109 float32 = 1e10
                                            if (v1 != v1) {
                                                v4[(fi90 << ((667) & 63))] = v3
                                                var v110 uint = 2
                                                _ = v110
                                            } else {
                                                var v111 string = (("foo bar" + "hello") + "a")
                                                var v112 float32 = v109
                                                var v113 string = v111
                                                _ = v111
                                                _ = v112
                                                _ = v113
                                            }
                                            v3 = false
                                            _ = v109
                                        } else {
                                            var v114 uint8 = 42
                                            fmt.Println(v3)
                                            _ = v114
                                        }
                                    }
                                    v97 += 153
                                    _ = v97
                                }
                                _ = v96
                            } else {
                                v115 := GenS12{F0: 42}
                                v116 := GenS2{F0: v5, F1: genFn1(), F2: 42}
                                _ = v115
                                _ = v116
                            }
                            fmt.Println(v5, fi90, v1)
                        }
                        _ = fi90
                    }
                    v117 := GenS14{F0: v5}
                    _ = v117
                }
            }
        }
        fmt.Println(v3)
    } else {
        switch 1 {
        case 492:
            v1 -= v1
        case 606:
            for fi118 := 0; fi118 < 13; fi118++ {
                v1 = v1
                _ = fi118
            }
        default:
            for fi119 := 0; fi119 < 19; fi119++ {
                var v120 float64 = v1
                for fi121 := 0; fi121 < 10; fi121++ {
                    v122 := GenS15{F0: v5, F1: 42, F2: 9223372036854775807}
                    _ = fi121
                    _ = v122
                }
                _ = fi119
                _ = v120
            }
        }
    }
    fmt.Println(v1, v3)
    v3 = v3
    switch "foo bar" {
    case "":
        if ((32767 != -281) && true) {
            v3 = (!false)
            v2 = append(v2, 1)
            v3 = (!(0 != -1))
        } else {
            switch 9223372036854775807 {
            case 9223372036854775807:
                fmt.Println(v3)
            default:
                fmt.Println(v1, v5)
            }
        }
    }
    var v123 uint32 = 301
    fmt.Println("v1:", v1)
    fmt.Println("v123:", v123)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v123
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
