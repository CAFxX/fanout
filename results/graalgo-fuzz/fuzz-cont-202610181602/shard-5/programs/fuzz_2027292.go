package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int16; F2 bool }
type GenS3 struct { F0 string; F1 uint; F2 int32 }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 bool; F1 uint; F2 bool }
type GenS6 struct { F0 string; F1 int16; F2 int }
type GenS7 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS8 struct { F0 int16; F1 uint64 }
type GenS9 struct { F0 int64; F1 int }
type GenS10 struct { F0 string; F1 string; F2 int16 }
type GenS11 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS12 struct { F0 uint8; F1 uint32 }
type GenS13 struct { F0 uint64; F1 int64 }
type GenS14 struct { F0 string; F1 uint; F2 uint }
type GenS15 struct { F0 int32; F1 uint; F2 uint8 }
type GenS16 struct { F0 uint; F1 int32 }
type GenS17 struct { F0 int64; F1 bool }
func genFn0() int32 {
    if (false || false) {
        var v5 int = 1
        v6 := GenS1{F0: 2}
        _ = v5
        _ = v6
    } else {
        if (true || false) {
            for fi7 := 0; fi7 < 17; fi7++ {
                v8 := []int{fi7}
                _ = fi7
                _ = v8
            }
            switch "hello" {
            case "hello":
                var v9 uint8 = 0
                _ = v9
            case "foo bar":
                var v10 uint16 = 65535
                _ = v10
            }
        } else {
            var v11 uint16 = 1
            if true {
                v11 = 65535
                var v12 uint32 = 0
                v12 += v12
                _ = v12
            } else {
                v13 := []int{4294967295}
                var v14 int = 32768
                switch "hello" {
                case "":
                    switch (v14 | (v14 | -123)) {
                    case 127:
                        v14 += (v14 & v14)
                    }
                case "hello":
                    for fi15 := 0; fi15 < 19; fi15++ {
                        for fi16 := 0; fi16 < 11; fi16++ {
                            if false {
                                _ = len(v13)
                            }
                            _ = fi16
                        }
                        fmt.Println(v14, v14)
                        _ = fi15
                    }
                default:
                    var v17 uint32 = 717
                    _ = v17
                }
                _ = v13
                _ = v14
            }
            if ((1 == -128) || (!false)) {
                var v18 uint = 2
                v19 := []int{2147483647}
                _ = v18
                _ = v19
            }
            _ = v11
        }
        var v20 int16 = 1000
        v21 := []int{32767, -137, 32767}
        _ = v20
        _ = v21
    }
    switch 32768 {
    case 872:
        switch 582 {
        case 32768:
            for fi22 := 0; fi22 < 6; fi22++ {
                v23 := []int{}
                if true {
                    var v24 float32 = 0.0
                    v25 := GenS2{F0: "hello", F1: -32768, F2: true}
                    switch (("\n\t" + (("a" + "foo bar") + ("a" + ""))) + "a") {
                    case "":
                        if len(v23) >= 2 {
                            v23 = v23[1:len(v23)]
                        }
                    default:
                        var v26 int8 = 42
                        _ = v26
                    }
                    _ = v24
                    _ = v25
                } else {
                    var v27 int64 = 42
                    if len(v23) >= 2 {
                        v23 = v23[1:len(v23)]
                    }
                    v27 = -772
                    _ = v27
                }
                _ = fi22
                _ = v23
            }
        case 9223372036854775807:
            if false {
                v28 := GenS3{F0: "a", F1: 869, F2: 2147483647}
                for fi29 := 0; fi29 < 17; fi29++ {
                    v30 := []int{(fi29 % ((fi29) | 1)), fi29, fi29}
                    var v31 int16 = 519
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                _ = v28
            }
        case 0:
            var v32 string = ""
            _ = v32
        default:
            var v33 int = 1
            _ = v33
        }
    case -1:
        var v34 int64 = 724
        _ = v34
    case 32768:
        v35 := GenS1{F0: 533}
        _ = v35
    default:
        var v36 int16 = 602
        _ = v36
    }
    return -2147483648
}
func genFn1() int {
    var v37 int32 = 572
    v37 -= v37
    v38 := []int{1, 1, 317, -223}
    _ = v37
    _ = v38
    return 32767
}

func main() {
    v1 := []int{-2147483648}
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    var v4 uint16 = 32767
    fmt.Println(v4)
    v4 = 32767
    v39 := GenS4{F0: 127, F1: 298}
    v40 := GenS5{F0: (!(false || ((128 != 505) && false))), F1: 42, F2: false}
    v4 = (v4 / ((v4) | 1))
    if true {
        var v41 float32 = 0.0
        if (false && false) {
            var v42 string = "foo bar"
            _ = v42
        }
        _ = v41
    } else {
        var v43 uint32 = 2147483647
        v43 = (v43 - (v43 * 8))
        _ = v43
    }
    v44 := GenS6{F0: "hello", F1: 42, F2: 127}
    var v45 int32 = 42
    if (-1 >= 127) {
        var v46 float32 = 0.1
        _ = v46
    }
    v4 = v4
    for fi47 := 0; fi47 < 16; fi47++ {
        v45 = v45
        _ = fi47
    }
    v48 := GenS7{F0: 535, F1: 366, F2: 679}
    v49 := GenS6{F0: "a", F1: 32767, F2: -706}
    if ((v4 ^ (v4 | v4)) != v4) {
        var v50 uint8 = 255
        _ = len(v1)
        fmt.Println(v50)
        _ = v50
    }
    var v51 string = ""
    for fi52 := 0; fi52 < 12; fi52++ {
        if ((v4 >= 255) && ((fi52 ^ fi52) > fi52)) {
            for fi53 := 0; fi53 < 10; fi53++ {
                switch 624 {
                case 210:
                    if (0 < -1) {
                        switch fi52 {
                        case 9223372036854775807:
                            v4 -= v4
                        default:
                            v51 = ""
                        }
                        v4 -= v4
                        fmt.Println(v51)
                    } else {
                        var v54 int8 = 127
                        _ = v54
                    }
                case 9223372036854775807:
                    if true {
                        v55 := GenS8{F0: -1, F1: 0}
                        fmt.Println(v51, v4)
                        _ = v55
                    } else {
                        var v56 uint = 213
                        _ = v56
                    }
                    var v57 uint32 = 42
                    _ = v57
                default:
                    var v58 uint = 2
                    _ = v58
                }
                _ = fi53
            }
            for fi59 := 0; fi59 < 14; fi59++ {
                v51 = ("hello" + ("" + (((v51 + "") + v51) + (("" + (v51 + ("a" + (v51 + "a")))) + v51))))
                v4 = (^(v4))
                switch -177 {
                case 128:
                    var v60 int64 = 426
                    _ = v60
                case -1:
                    if (((!(true || ((42 == -678) && true))) && (0.0 >= 3.14159)) || false) {
                        if false {
                            var v61 int32 = (v45 & v45)
                            v45 += (^(v45))
                            if len(v1) > 0 {
                                v1[0] = -901
                            }
                            _ = v61
                        }
                        var v62 uint = 159
                        if (true && (0.1 == 0.0)) {
                            v63 := GenS1{F0: 2}
                            if (true && false) {
                                fmt.Println(v4, v4)
                                var v64 uint8 = 179
                                var v65 int64 = 1
                                _ = v64
                                _ = v65
                            } else {
                                v66 := GenS9{F0: 0, F1: fi52}
                                if (32767 <= 982) {
                                    v1 = append(v1, fi52)
                                } else {
                                    v45 += 1
                                    if false {
                                        v4 = (v4 + (v4 * (v4 + v4)))
                                    }
                                    switch 32768 {
                                    case 32767:
                                        v67 := GenS10{F0: v51, F1: ("\n\t" + (("" + v51) + v51)), F2: 961}
                                        _ = v67
                                    default:
                                        if (!true) {
                                            switch (fi59 + int(v62)) {
                                            case 42:
                                                fmt.Println(v4, v62, fi52)
                                                if (fi59 >= (fi52 << (((v62 / ((v62) | 1))) & 63))) {
                                                    if false {
                                                        switch (((v51 + "a") + (((((("\n\t" + "") + "\n\t") + v51) + v51) + "") + ("\n\t" + "\n\t"))) + (v51 + "hello")) {
                                                        case "":
                                                            v4 += v4
                                                        case "a":
                                                            v68 := GenS8{F0: -505, F1: 0}
                                                            v69 := v2[v51]
                                                            var v70 uint8 = 140
                                                            _ = v68
                                                            _ = v69
                                                            _ = v70
                                                        case "hello":
                                                            var v71 float32 = (1e10 * 3.14159)
                                                            _ = v71
                                                        }
                                                    }
                                                } else {
                                                    v51 = ""
                                                    var v72 int8 = 0
                                                    var v73 float32 = 0.1
                                                    _ = v72
                                                    _ = v73
                                                }
                                                var v74 float32 = (((float32(((float64((float32(1.5) - float32(0.0))) + (1.5 * 1e10)) / (0.1 / ((0.0 * float64((float32(0.0) / 0.1))) + 1e10)))) - float32((3.14159 * (1.5 - -2.25)))) - 1.5) + float32((0.0 + ((0.0 + 1.5) / ((0.1 + (1.5 / ((float64(3.14159) + float64(0.1)) + (3.14159 * 3.14159)))) / (0.0 - -2.25))))))
                                                _ = v74
                                            }
                                        } else {
                                            fmt.Println(fi52, fi59, fi52)
                                            v75 := GenS4{F0: 185, F1: 0}
                                            switch (-(fi59)) {
                                            case 2147483647:
                                                v76 := v2[v51]
                                                _ = v76
                                            case -2147483648:
                                                var v77 float32 = (float32(1e10) + (1e10 * 1e10))
                                                _ = v77
                                            }
                                            _ = v75
                                        }
                                        v4 = (v4 & (^(v4)))
                                    }
                                }
                                fmt.Println(v51, fi59)
                                _ = v66
                            }
                            _ = v63
                        }
                        if true {
                            if (42 > 110) {
                                v78 := GenS11{F0: 9223372036854775807, F1: 593, F2: 0}
                                v4 = (v4 % ((uint16(v62)) | 1))
                                fmt.Println(v4, fi52, fi59)
                                _ = v78
                            } else {
                                if (!false) {
                                    v79 := GenS11{F0: 1, F1: 18446744073709551615, F2: 2}
                                    v51 = v51
                                    _ = v79
                                }
                            }
                            var v80 uint32 = 4294967295
                            v62 += (v62 << ((255) & 63))
                            v81 := GenS12{F0: 0, F1: 2}
                            _ = v80
                            _ = v81
                        }
                        _ = v62
                    }
                case 127:
                    switch 42 {
                    case 4294967295:
                        var v82 uint32 = 353
                        _ = v82
                    case 32767:
                        if false {
                            var v83 float32 = (-2.25 * -2.25)
                            _ = len(v3)
                            if (-1 < 42) {
                                switch v51 {
                                case "hello":
                                    if len(v1) > 0 {
                                        v1[0] = fi52
                                    }
                                case "":
                                    if true {
                                        fmt.Println(v4, v45, fi52)
                                    } else {
                                        if false {
                                            var v84 int = fi59
                                            switch "foo bar" {
                                            case "hello":
                                                var v85 uint64 = 4294967295
                                                _ = v85
                                            }
                                            var v86 int16 = 199
                                            _ = v84
                                            _ = v86
                                        } else {
                                            v51 = v51
                                            v1 = append(v1, genFn1())
                                            v83 = 0.0
                                            switch "" {
                                            case "\n\t":
                                                v87 := GenS9{F0: 42, F1: 9223372036854775807}
                                                _ = v87
                                            }
                                            var v88 uint64 = 4294967295
                                            _ = v88
                                        }
                                        switch v51 {
                                        case "hello":
                                            v45 = (v45 + v45)
                                        case "a":
                                            if false {
                                                v4 = (v4 ^ 32767)
                                                fmt.Println(v51, fi52)
                                            } else {
                                                v83 = v83
                                                var v89 float32 = v83
                                                v90 := GenS13{F0: 18446744073709551615, F1: -9223372036854775808}
                                                v1 = append(v1, fi52)
                                                var v91 int32 = 377
                                                _ = v89
                                                _ = v90
                                                _ = v91
                                            }
                                        default:
                                            v92 := GenS14{F0: v51, F1: 4294967295, F2: 615}
                                            _ = v92
                                        }
                                        var v93 string = v51
                                        _ = v93
                                    }
                                case "a":
                                    v51 = v51
                                default:
                                    v83 = v83
                                }
                                fmt.Println(v51, v83, fi59)
                                v45 += (v45 % ((v45) | 1))
                            } else {
                                fmt.Println(v51, v51, fi59)
                            }
                            _ = v83
                        } else {
                            delete(v2, "hello")
                            var v94 uint = 18446744073709551615
                            _ = v94
                        }
                    case 127:
                        var v95 int8 = 42
                        _ = v95
                    default:
                        v96 := GenS15{F0: -230, F1: 452, F2: 255}
                        _ = v96
                    }
                }
                _ = fi59
            }
        } else {
            switch v51 {
            case "\n\t":
                var v97 int64 = 495
                _ = v97
            case "":
                if len(v1) > 0 {
                    v1[0] = fi52
                }
            default:
                v98 := v2[v51]
                _ = v98
            }
            if false {
                if len(v1) > 0 {
                    v1[0] = fi52
                }
                switch (((v51 + (((v51 + ((v51 + v51) + (v51 + "foo bar"))) + (v51 + v51)) + "foo bar")) + v51) + v51) {
                case "\n\t":
                    fmt.Println(v51)
                case "":
                    switch "foo bar" {
                    case "hello":
                        v99 := GenS6{F0: (("\n\t" + "foo bar") + ("foo bar" + "a")), F1: -1, F2: fi52}
                        _ = v99
                    case "a":
                        v4 = v4
                    case "\n\t":
                        for fi100 := 0; fi100 < 3; fi100++ {
                            v2[("foo bar" + "foo bar")] = true
                            v51 = (v51 + "")
                            v45 -= v45
                            _ = fi100
                        }
                    }
                default:
                    switch "\n\t" {
                    case "\n\t":
                        _ = len(v1)
                    case "foo bar":
                        switch "hello" {
                        case "hello":
                            v45 = genFn0()
                        }
                    case "a":
                        v101 := GenS1{F0: 4294967295}
                        _ = v101
                    default:
                        for fi102 := 0; fi102 < 20; fi102++ {
                            if (((127 >= -128) || (v51 > v51)) || (1e10 != -2.25)) {
                                fmt.Println(v45, v45)
                                fmt.Println(fi52)
                            }
                            fmt.Println(fi102, fi102)
                            v103 := GenS7{F0: 65535, F1: 432, F2: 533}
                            v3[(^(fi52))] = (fi52 ^ fi102)
                            if false {
                                v104 := GenS16{F0: 0, F1: v45}
                                _ = v104
                            }
                            switch (-(fi102)) {
                            case 42:
                                v51 = v51
                            case -9223372036854775808:
                                if (921 < 0) {
                                    switch int(v45) {
                                    case -1:
                                        var v105 float64 = 1.5
                                        _ = v105
                                    case 539:
                                        var v106 int16 = -32768
                                        _ = v106
                                    default:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    switch v51 {
                                    case "\n\t":
                                        if (false && (v4 > (v4 - (^(v4))))) {
                                            if len(v1) > 0 {
                                                v1[0] = (fi52 * fi52)
                                            }
                                            var v107 bool = (!true)
                                            var v108 int16 = 42
                                            var v109 uint8 = 2
                                            _ = v107
                                            _ = v108
                                            _ = v109
                                        } else {
                                            switch genFn1() {
                                            case 128:
                                                v110 := GenS4{F0: 21, F1: 495}
                                                _ = v110
                                            case 322:
                                                if (32767 < (v4 * 2)) {
                                                    v4 = (v4 * 32767)
                                                }
                                            case 636:
                                                fmt.Println(fi102, v45)
                                            }
                                            v3[790] = 0
                                            v51 = "foo bar"
                                        }
                                    case "a":
                                        var v111 int = (fi52 | (fi52 * (fi52 << ((uint(v4)) & 63))))
                                        _ = v111
                                    }
                                    fmt.Println(fi102, v45, fi52)
                                }
                            default:
                                if false {
                                    fmt.Println(v51, v4, v4)
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    if ((!true) && true) {
                                        v112 := GenS3{F0: ("hello" + ((("" + v51) + v51) + v51)), F1: 4294967295, F2: -2147483648}
                                        v113 := GenS17{F0: 9223372036854775807, F1: true}
                                        if false {
                                            v51 = "a"
                                            v51 = ("" + v51)
                                            v114 := GenS3{F0: v51, F1: 255, F2: v45}
                                            _ = v114
                                        } else {
                                            var v115 uint8 = 59
                                            switch "\n\t" {
                                            case "\n\t":
                                                switch "" {
                                                case "":
                                                    fmt.Println(v45)
                                                default:
                                                    _ = len(v1)
                                                }
                                            }
                                            fmt.Println(v51)
                                            _ = v115
                                        }
                                        _ = v112
                                        _ = v113
                                    }
                                } else {
                                    fmt.Println(v45)
                                }
                            }
                            _ = fi102
                            _ = v103
                        }
                    }
                }
                var v116 uint8 = 48
                _ = v116
            } else {
                v117 := GenS6{F0: ("\n\t" + "hello"), F1: -1, F2: (-(fi52))}
                v118 := GenS9{F0: 616, F1: int(v45)}
                _ = v117
                _ = v118
            }
        }
        _ = fi52
    }
    var v119 int64 = -258
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v119:", v119)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v119
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v44
    _ = v45
    _ = v48
    _ = v49
    _ = v51
}
