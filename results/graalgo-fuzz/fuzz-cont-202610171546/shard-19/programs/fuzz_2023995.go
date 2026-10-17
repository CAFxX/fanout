package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint32; F1 int8; F2 int64 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint16; F1 int32; F2 int }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8; F1 bool }
type GenS7 struct { F0 uint32; F1 int64 }
type GenS8 struct { F0 int8; F1 int16; F2 int64 }
type GenS9 struct { F0 uint32; F1 int; F2 int16 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS12 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS13 struct { F0 uint; F1 bool }
type GenS14 struct { F0 int8; F1 bool; F2 uint64 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 uint16; F1 uint8 }
func genFn0(p0 uint, p1 uint) uint32 {
    p1 = p1
    for fi4 := 0; fi4 < 17; fi4++ {
        var v5 bool = (false || (0.1 >= float64(3.14159)))
        for fi6 := 0; fi6 < 13; fi6++ {
            if true {
                if ((true || v5) && (-768 > 840)) {
                    v7 := GenS1{F0: -128, F1: 9223372036854775807}
                    var v8 uint16 = 2
                    v9 := GenS1{F0: -1, F1: 9223372036854775807}
                    _ = v7
                    _ = v8
                    _ = v9
                }
            }
            _ = fi6
        }
        _ = fi4
        _ = v5
    }
    switch "hello" {
    case "hello":
        fmt.Println(p0, p1)
    case "foo bar":
        v10 := []int{0, 42}
        _ = v10
    case "\n\t":
        v11 := []int{32767, -133, 347, 9223372036854775807}
        _ = v11
    }
    _ = p0
    _ = p1
    return 2147483648
}
func genFn1(p0 int32) uint8 {
    p0 -= -521
    if false {
        switch 32768 {
        case -298:
            for fi12 := 0; fi12 < 18; fi12++ {
                for fi13 := 0; fi13 < 11; fi13++ {
                    switch ("hello" + "\n\t") {
                    case "foo bar":
                        p0 -= (p0 * 2147483647)
                    default:
                        if (((((-2.25 * 0.1) + -2.25) - 1.5) - 3.14159) > 3.14159) {
                            if (!(!true)) {
                                p0 += p0
                                fmt.Println(fi13, p0, fi12)
                            } else {
                                fmt.Println(fi12, fi12, fi13)
                                var v14 string = "\n\t"
                                _ = v14
                            }
                            if (128 > 28) {
                                p0 = p0
                            }
                        } else {
                            fmt.Println(fi13)
                            v15 := []int{-2147483648, (fi12 >> ((244) & 63))}
                            fmt.Println(p0, p0)
                            var v16 uint = 112
                            _ = v15
                            _ = v16
                        }
                    }
                    var v17 int8 = 42
                    v17 = (v17 & (v17 >> ((962) & 63)))
                    _ = fi13
                    _ = v17
                }
                if ((!false) && false) {
                    for fi18 := 0; fi18 < 18; fi18++ {
                        var v19 uint8 = 2
                        _ = fi18
                        _ = v19
                    }
                    var v20 float64 = 1e10
                    _ = v20
                } else {
                    switch "foo bar" {
                    case "hello":
                        if true {
                            v21 := GenS1{F0: 127, F1: 9223372036854775807}
                            if true {
                                p0 -= 42
                                v22 := GenS1{F0: 0, F1: -586}
                                _ = v22
                            }
                            _ = v21
                        }
                    case "a":
                        for fi23 := 0; fi23 < 7; fi23++ {
                            v24 := GenS2{F0: 0, F1: 1, F2: -1}
                            var v25 int16 = -914
                            switch (fi23 - fi23) {
                            case 128:
                                p0 = 2147483647
                            case -2147483648:
                                v26 := GenS3{F0: v25, F1: 431}
                                _ = v26
                            default:
                                v25 += v25
                            }
                            v27 := GenS1{F0: 0, F1: 9223372036854775807}
                            _ = fi23
                            _ = v24
                            _ = v25
                            _ = v27
                        }
                    default:
                        v28 := GenS2{F0: 2147483647, F1: -105, F2: -1}
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 18; fi29++ {
                        var v30 int64 = -752
                        fmt.Println(fi12, v30, v30)
                        if true {
                            v31 := []int{(fi29 * 32768), fi12}
                            _ = v31
                        }
                        var v32 int16 = 283
                        _ = fi29
                        _ = v30
                        _ = v32
                    }
                }
                _ = fi12
            }
        case 2147483647:
            v33 := GenS2{F0: genFn0(2, 971), F1: -101, F2: 9223372036854775807}
            _ = v33
        case -1:
            p0 = (p0 - p0)
        default:
            p0 += 0
        }
        var v34 uint64 = 4294967295
        _ = v34
    }
    _ = p0
    return 2
}

func main() {
    var v1 uint = 65535
    var v2 uint64 = 777
    v3 := make(map[uint32]string)
    for fi35 := 0; fi35 < 9; fi35++ {
        for fi36 := 0; fi36 < 1; fi36++ {
            v37 := []int{}
            if false {
                v2 = (v2 & (^(v2)))
                v2 -= (v2 >> ((0) & 63))
                switch ("a" + (("hello" + ("foo bar" + "foo bar")) + "")) {
                case "hello":
                    switch (^(fi36)) {
                    case 127:
                        var v38 float32 = -2.25
                        _ = v38
                    case 180:
                        v2 = v2
                    }
                case "a":
                    v2 = v2
                }
                if len(v37) >= 2 {
                    v37 = v37[1:len(v37)]
                }
            }
            v39 := GenS2{F0: 1, F1: 100, F2: 42}
            _ = fi36
            _ = v37
            _ = v39
        }
        if (false && (!false)) {
            v40 := GenS3{F0: 32767, F1: 255}
            switch "\n\t" {
            case "":
                v41 := GenS4{F0: 0, F1: -1, F2: fi35}
                _ = v41
            case "\n\t":
                fmt.Println(v1)
            case "foo bar":
                v42 := GenS5{F0: 0}
                _ = v42
            }
            switch fi35 {
            case 1:
                v43 := GenS1{F0: 21, F1: -371}
                _ = v43
            default:
                v1 -= v1
            }
            _ = v40
        }
        _ = fi35
    }
    fmt.Println(v2, v1)
    var v44 bool = true
    var v45 int16 = 42
    switch "foo bar" {
    case "hello":
        v46 := GenS6{F0: 100, F1: (v44 && (v1 < (^(v1))))}
        _ = v46
    case "a":
        if (!false) {
            var v47 bool = v44
            v48 := []int{4294967295, -1, -9223372036854775808}
            switch "\n\t" {
            case "\n\t":
                for fi49 := 0; fi49 < 20; fi49++ {
                    v50 := GenS1{F0: 42, F1: 335}
                    v45 += v45
                    _ = fi49
                    _ = v50
                }
            case "foo bar":
                switch "hello" {
                case "foo bar":
                    fmt.Println(v2, v45)
                default:
                    fmt.Println(v47, v2)
                }
            }
            _ = v47
            _ = v48
        } else {
            switch "" {
            case "\n\t":
                v51 := []int{4294967295, 1, 834, 91}
                _ = v51
            case "hello":
                v45 += -1
            }
            fmt.Println(v44)
            if (((42 == 0) && v44) && (!(("" + ("foo bar" + ("foo bar" + "\n\t"))) < "\n\t"))) {
                v1 += 4294967295
            }
        }
    default:
        for fi52 := 0; fi52 < 17; fi52++ {
            if true {
                v1 -= v1
                v53 := GenS5{F0: -53}
                _ = v53
            }
            v44 = v44
            _ = fi52
        }
    }
    v45 -= 1000
    v45 = v45
    v1 -= (v1 - (v1 / ((2) | 1)))
    if true {
        if v44 {
            v54 := GenS1{F0: -1, F1: 0}
            fmt.Println(v1, v44)
            _ = v54
        }
        var v55 bool = v44
        v56 := GenS5{F0: -128}
        _ = v55
        _ = v56
    } else {
        v57 := GenS7{F0: 4294967295, F1: -9223372036854775808}
        v58 := []int{-610, 9223372036854775807, 2147483647}
        v59 := GenS8{F0: -128, F1: (v45 & -1), F2: 42}
        _ = v57
        _ = v58
        _ = v59
    }
    var v60 int64 = 42
    var v61 int64 = (v60 + -1)
    fmt.Println(v60)
    for fi62 := 0; fi62 < 17; fi62++ {
        if (v44 && v44) {
            for fi63 := 0; fi63 < 18; fi63++ {
                if (42 <= 0) {
                    switch ("hello" + "foo bar") {
                    case "hello":
                        if v44 {
                            v44 = ((uint32(v2) >= 2147483648) || (v61 < v60))
                        } else {
                            switch ("foo bar" + "a") {
                            case "foo bar":
                                fmt.Println(fi62)
                            case "":
                                v61 = v60
                                var v64 int = (fi63 | (fi62 << ((v1) & 63)))
                                if v44 {
                                    v65 := GenS9{F0: 2, F1: 128, F2: (v45 << (((v1 + v1)) & 63))}
                                    _ = v65
                                }
                                _ = v64
                            case "\n\t":
                                v66 := GenS10{F0: (^(fi63))}
                                _ = v66
                            default:
                                v67 := GenS4{F0: 519, F1: -2147483648, F2: (fi63 ^ fi62)}
                                _ = v67
                            }
                            var v68 float32 = 1.5
                            var v69 int32 = -237
                            v70 := GenS4{F0: 161, F1: (v69 & v69), F2: 0}
                            _ = v68
                            _ = v69
                            _ = v70
                        }
                    case "a":
                        var v71 float32 = 3.14159
                        _ = v71
                    }
                } else {
                    var v72 float64 = 0.1
                    v73 := GenS11{F0: v2, F1: 420, F2: v2}
                    _ = v72
                    _ = v73
                }
                _ = fi63
            }
        } else {
            if ((^(v2)) > (v2 * 1)) {
                switch (fi62 | (fi62 | 2147483647)) {
                case 32767:
                    v74 := GenS12{F0: -567, F1: 0, F2: 18446744073709551615}
                    _ = v74
                case 905:
                    var v75 uint = (v1 * v1)
                    _ = v75
                case -9223372036854775808:
                    switch "a" {
                    case "hello":
                        var v76 uint = 220
                        _ = v76
                    case "\n\t":
                        v77 := GenS13{F0: v1, F1: (!true)}
                        _ = v77
                    default:
                        for fi78 := 0; fi78 < 12; fi78++ {
                            v79 := GenS11{F0: v2, F1: 1, F2: 456}
                            var v80 int32 = int32(v61)
                            v81 := GenS14{F0: 100, F1: v44, F2: 524}
                            _ = fi78
                            _ = v79
                            _ = v80
                            _ = v81
                        }
                    }
                }
                if false {
                    v82 := v3[4294967295]
                    var v83 int8 = 127
                    _ = v82
                    _ = v83
                }
                fmt.Println(v60, v60, v44)
            }
            switch (fi62 * 525) {
            case -1:
                var v84 float64 = 3.14159
                _ = v84
            }
            v85 := v3[629]
            _ = v85
        }
        fmt.Println(v2)
        v86 := GenS6{F0: -1, F1: true}
        _ = fi62
        _ = v86
    }
    var v87 float32 = 0.0
    for fi88 := 0; fi88 < 17; fi88++ {
        switch -473 {
        case 128:
            if ((1 == -2147483648) && v44) {
                if v44 {
                    v44 = (v44 && v44)
                }
                v89 := GenS2{F0: 2147483648, F1: -1, F2: 0}
                for fi90 := 0; fi90 < 4; fi90++ {
                    v91 := []int{(fi90 & fi90)}
                    _ = fi90
                    _ = v91
                }
                _ = v89
            } else {
                v92 := GenS10{F0: fi88}
                v45 = (^(v45))
                _ = v92
            }
        case 9223372036854775807:
            if v44 {
                var v93 uint8 = 240
                v94 := GenS12{F0: int32(v45), F1: (v93 | (v93 - 2)), F2: (v2 * 85)}
                for fi95 := 0; fi95 < 15; fi95++ {
                    var v96 uint8 = 139
                    v97 := GenS15{F0: 295}
                    v44 = (((0.0 * float64((0.1 * 1.5))) * float64(v87)) == 1.5)
                    _ = fi95
                    _ = v96
                    _ = v97
                }
                _ = v93
                _ = v94
            }
        }
        if true {
            var v98 int8 = 25
            v60 -= (^(v61))
            _ = v98
        } else {
            for fi99 := 0; fi99 < 5; fi99++ {
                switch 32768 {
                case 333:
                    if (!v44) {
                        v100 := GenS16{F0: 997, F1: 127}
                        v1 = v1
                        v87 = v87
                        var v101 string = "\n\t"
                        _ = v100
                        _ = v101
                    }
                default:
                    if ((fi99 << ((2) & 63)) == 1) {
                        v60 += v61
                        switch 127 {
                        case 128:
                            fmt.Println(v44)
                        case 651:
                            if (566 > genFn0(v1, (^(v1)))) {
                                fmt.Println(v2)
                                var v102 float32 = v87
                                _ = v102
                            }
                            var v103 uint8 = 127
                            if v44 {
                                switch "foo bar" {
                                case "":
                                    var v104 int = fi99
                                    _ = v104
                                case "hello":
                                    switch (fi88 / (((fi99 % (((fi88 * 42)) | 1))) | 1)) {
                                    case -41:
                                        v61 = 42
                                    default:
                                        switch fi99 {
                                        case -659:
                                            var v105 int8 = -128
                                            _ = v105
                                        }
                                    }
                                }
                                delete(v3, 4294967295)
                                var v106 float64 = (1e10 * 3.14159)
                                _ = v106
                            } else {
                                v44 = (!false)
                                var v107 int = -9223372036854775808
                                _ = v107
                            }
                            _ = v103
                        case 32767:
                            v108 := []int{}
                            _ = v108
                        default:
                            v109 := []int{fi88, 656, (fi99 + fi88)}
                            _ = v109
                        }
                    }
                }
                var v110 string = (("" + "") + "hello")
                var v111 int = 0
                v111 -= 180
                _ = fi99
                _ = v110
                _ = v111
            }
            v112 := []int{fi88, (fi88 >> ((v1) & 63)), (fi88 << ((841) & 63)), fi88}
            if (!true) {
                if len(v112) >= 2 {
                    v112 = v112[1:len(v112)]
                }
                v112 = append(v112, 526)
            } else {
                v2 -= (^(v2))
                var v113 float32 = ((1.5 / (v87 - v87)) / (v87 - v87))
                v45 -= -52
                _ = v113
            }
            _ = v112
        }
        _ = fi88
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v60
    _ = v61
    _ = v87
}
