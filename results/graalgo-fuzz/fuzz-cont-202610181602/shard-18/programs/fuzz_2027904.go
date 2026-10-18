package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 int32; F1 uint }
type GenS7 struct { F0 uint; F1 int64 }
type GenS8 struct { F0 uint8; F1 uint8 }
type GenS9 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS10 struct { F0 uint64; F1 uint64 }
type GenS11 struct { F0 uint64; F1 uint16 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int16; F1 bool; F2 uint64 }
type GenS14 struct { F0 uint32; F1 bool }

func main() {
    v1 := []int{42, 4294967295, -1}
    v2 := make(map[uint32]bool)
    var v3 uint = 1
    var v4 int16 = 179
    v5 := []string{"foo bar"}
    v6 := []int{-2147483648, 32767, 128}
    var v7 uint32 = 4294967295
    switch "a" {
    case "":
        if false {
            switch 473 {
            case -2147483648:
                var v8 uint16 = 0
                _ = v8
            default:
                for fi9 := 0; fi9 < 12; fi9++ {
                    for fi10 := 0; fi10 < 14; fi10++ {
                        if true {
                            if false {
                                var v11 int = fi10
                                _ = v11
                            } else {
                                v12 := GenS1{F0: 2}
                                v13 := GenS1{F0: 475}
                                v14 := GenS1{F0: 42}
                                _ = v12
                                _ = v13
                                _ = v14
                            }
                        } else {
                            var v15 int16 = v4
                            _ = v15
                        }
                        v4 += (^(v4))
                        _ = fi10
                    }
                    var v16 int64 = 42
                    _ = fi9
                    _ = v16
                }
            }
            v4 -= (v4 - v4)
        }
    default:
        var v17 uint = 65535
        _ = v17
    }
    var v18 uint32 = (v7 << ((v3) & 63))
    if false {
        if true {
            var v19 string = (("foo bar" + "\n\t") + "hello")
            _ = v19
        }
    } else {
        v7 = 1
        v20 := GenS2{F0: 42}
        _ = v20
    }
    v21 := GenS3{F0: -258}
    fmt.Println(v7, v7, v4)
    if (-878 == 42) {
        if false {
            for fi22 := 0; fi22 < 4; fi22++ {
                switch ("a" + "\n\t") {
                case "hello":
                    v23 := GenS2{F0: 1}
                    _ = v23
                case "":
                    var v24 int16 = -32768
                    _ = v24
                case "foo bar":
                    for fi25 := 0; fi25 < 17; fi25++ {
                        switch fi22 {
                        case 0:
                            var v26 int32 = 42
                            _ = v26
                        default:
                            if ((-810 != -1) || true) {
                                if (!true) {
                                    v3 += v3
                                    if ((1e10 - 0.0) <= 1e10) {
                                        v4 += v4
                                        v7 = v18
                                    }
                                    v4 = (v4 - v4)
                                } else {
                                    v7 = v18
                                    switch int(v18) {
                                    case 98:
                                        _ = len(v6)
                                    case 9223372036854775807:
                                        v2[(^(v7))] = (((v18 > 1) || true) && true)
                                    case -1:
                                        v4 -= (v4 & -792)
                                        var v27 float32 = 0.1
                                        delete(v2, (^(v7)))
                                        _ = v27
                                    }
                                    switch fi25 {
                                    case 0:
                                        if false {
                                            v28 := GenS2{F0: 128}
                                            v29 := GenS2{F0: uint8(v3)}
                                            switch (fi22 + (fi25 | (^(fi25)))) {
                                            case 214:
                                                var v30 uint8 = 127
                                                _ = v30
                                            case -263:
                                                if true {
                                                    if true {
                                                        var v31 float32 = 0.0
                                                        v32 := GenS4{F0: -128}
                                                        var v33 float64 = 1e10
                                                        _ = v31
                                                        _ = v32
                                                        _ = v33
                                                    }
                                                    var v34 float64 = -2.25
                                                    v4 = (v4 << ((v3) & 63))
                                                    _ = v34
                                                }
                                            default:
                                                var v35 uint64 = 2
                                                _ = v35
                                            }
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v36 int8 = 17
                                                _ = v36
                                            case "hello":
                                                v1 = append(v1, (^(fi25)))
                                            default:
                                                v4 -= 1000
                                            }
                                            _ = v28
                                            _ = v29
                                        } else {
                                            switch "\n\t" {
                                            case "\n\t":
                                                v37 := GenS2{F0: 255}
                                                _ = v37
                                            case "":
                                                v18 -= v7
                                                v38 := GenS4{F0: 62}
                                                if (uint(v7) >= (v3 | (v3 + 0))) {
                                                    v4 = 32767
                                                    v39 := GenS5{F0: v4, F1: 33}
                                                    var v40 uint8 = 0
                                                    _ = v39
                                                    _ = v40
                                                }
                                                _ = v38
                                            }
                                            var v41 int64 = 0
                                            v3 += v3
                                            _ = v41
                                        }
                                    case -538:
                                        v3 = v3
                                        fmt.Println(v18, fi25, fi22)
                                        var v42 int64 = 802
                                        v42 = v42
                                        fmt.Println(v7, v3)
                                        var v43 float64 = 0.1
                                        _ = v42
                                        _ = v43
                                    default:
                                        _ = len(v2)
                                    }
                                }
                                var v44 int16 = v4
                                fmt.Println(v44, v18, fi25)
                                _ = v44
                            } else {
                                if false {
                                    v45 := GenS2{F0: 42}
                                    _ = v45
                                } else {
                                    v46 := GenS4{F0: -119}
                                    v3 += v3
                                    if (!(true && false)) {
                                        if true {
                                            var v47 int32 = 476
                                            var v48 float64 = (3.14159 + (-2.25 - 1e10))
                                            if false {
                                                switch (fi25 - (^(fi25))) {
                                                case 42:
                                                    v49 := GenS6{F0: (v47 | 989), F1: (v3 * 4294967295)}
                                                    _ = v49
                                                case -1:
                                                    _ = len(v2)
                                                case 9223372036854775807:
                                                    var v50 uint64 = 2
                                                    _ = v50
                                                default:
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    var v51 float32 = 0.0
                                                    v52 := v2[v7]
                                                    _ = v51
                                                    _ = v52
                                                }
                                            }
                                            _ = v47
                                            _ = v48
                                        } else {
                                            fmt.Println(v3)
                                            fmt.Println(v7)
                                            var v53 uint16 = 32767
                                            _ = v53
                                        }
                                        v6 = append(v6, (fi25 | fi22))
                                        var v54 uint16 = 255
                                        _ = v54
                                    } else {
                                        v7 = (v7 - (v18 % ((v7) | 1)))
                                        var v55 float64 = ((((0.1 * (3.14159 + 1.5)) + (float64(-2.25) / float64(-2.25))) + 0.1) / 0.1)
                                        _ = v55
                                    }
                                    _ = v46
                                }
                            }
                        }
                        _ = len(v1)
                        v7 += (v7 >> (((v3 >> (((v3 | (v3 - (v3 | (v3 << ((128) & 63)))))) & 63))) & 63))
                        _ = len(v2)
                        _ = fi25
                    }
                default:
                    if true {
                        fmt.Println(v7, v7)
                        if false {
                            if true {
                                v3 = v3
                            } else {
                                if (127 > 201) {
                                    v3 = (^(v3))
                                    v7 = (^(v18))
                                }
                                fmt.Println(v18, v18, v3)
                            }
                            v3 += 0
                        } else {
                            if (-150 >= -219) {
                                v18 += (v7 * (^(v18)))
                            }
                            fmt.Println(v4, v3)
                        }
                        switch ("a" + "hello") {
                        case "\n\t":
                            v3 = v3
                        }
                    } else {
                        if false {
                            switch "" {
                            case "foo bar":
                                var v56 float64 = 3.14159
                                _ = v56
                            }
                            v57 := GenS3{F0: (fi22 ^ (fi22 | (fi22 * fi22)))}
                            _ = v57
                        }
                        fmt.Println(v4, v3)
                    }
                }
                var v58 float64 = -2.25
                switch "" {
                case "\n\t":
                    if ((1e10 * 3.14159) > 3.14159) {
                        fmt.Println(v18)
                        fmt.Println(v4, v4)
                        if ((true || true) && (false || true)) {
                            if false {
                                switch (fi22 * (fi22 / ((fi22) | 1))) {
                                case 4294967295:
                                    switch (fi22 - -9223372036854775808) {
                                    case -360:
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    case 127:
                                        var v59 uint16 = 255
                                        _ = v59
                                    }
                                case 42:
                                    if (0 == 994) {
                                        var v60 float32 = float32(((v58 + v58) / ((v58 + v58) / (v58 * 1.5))))
                                        _ = v60
                                    } else {
                                        if true {
                                            if true {
                                                v61 := GenS5{F0: 1000, F1: 42}
                                                v2[0] = (true || ((^(v18)) < 0))
                                                v4 = -32768
                                                _ = v61
                                            } else {
                                                var v62 uint8 = 205
                                                fmt.Println(v7)
                                                var v63 uint = (v3 - v3)
                                                _ = v62
                                                _ = v63
                                            }
                                            for fi64 := 0; fi64 < 11; fi64++ {
                                                if (true || false) {
                                                    switch ("hello" + ("foo bar" + ("" + "\n\t"))) {
                                                    case "hello":
                                                        v58 += v58
                                                    case "foo bar":
                                                        var v65 int16 = v4
                                                        _ = v65
                                                    case "a":
                                                        v18 = (v7 * v18)
                                                    }
                                                    v6 = append(v6, fi22)
                                                } else {
                                                    if true {
                                                        fmt.Println(v7, fi22)
                                                        var v66 int32 = -2147483648
                                                        _ = v66
                                                    } else {
                                                        switch fi64 {
                                                        case -2147483648:
                                                            v3 = v3
                                                        case 32767:
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                var v67 int64 = -879
                                                                _ = v67
                                                            default:
                                                                if true {
                                                                    v3 += v3
                                                                    _ = len(v2)
                                                                    v18 -= v18
                                                                }
                                                            }
                                                        }
                                                        switch fi64 {
                                                        case 684:
                                                            v4 += v4
                                                        case 2147483647:
                                                            switch (fi64 + fi22) {
                                                            case 0:
                                                                if false {
                                                                    v4 -= (v4 / (((v4 / (((v4 / (((-(v4))) | 1))) | 1))) | 1))
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        v3 -= v3
                                                                    case "":
                                                                        var v68 int = (^(fi22))
                                                                        v58 += (1.5 * (float64(((((0.0 - -2.25) * 3.14159) - (0.0 - 1e10)) + ((0.0 + -2.25) * -2.25))) + (3.14159 + -2.25)))
                                                                        var v69 float32 = 0.1
                                                                        v3 = 255
                                                                        var v70 string = "foo bar"
                                                                        _ = v68
                                                                        _ = v69
                                                                        _ = v70
                                                                    case "a":
                                                                        v18 -= 4294967295
                                                                    default:
                                                                        var v71 uint = 1
                                                                        _ = v71
                                                                    }
                                                                } else {
                                                                    var v72 uint = 42
                                                                    if (!(!true)) {
                                                                        v73 := GenS7{F0: v72, F1: -1}
                                                                        v3 = v72
                                                                        fmt.Println(v18)
                                                                        _ = v73
                                                                    } else {
                                                                        v72 += (v72 | v3)
                                                                        v7 = (^(v18))
                                                                        _ = len(v6)
                                                                        var v74 int16 = (v4 & 1000)
                                                                        _ = v74
                                                                    }
                                                                    if true {
                                                                        v75 := GenS6{F0: -816, F1: 350}
                                                                        var v76 int32 = -2147483648
                                                                        v77 := GenS7{F0: (v3 << ((122) & 63)), F1: 565}
                                                                        _ = v75
                                                                        _ = v76
                                                                        _ = v77
                                                                    }
                                                                    _ = v72
                                                                }
                                                            case -268:
                                                                fmt.Println(v4, fi22)
                                                            case -1:
                                                                fmt.Println(fi64, v18)
                                                            default:
                                                                v4 += (v4 % ((1) | 1))
                                                            }
                                                            switch (fi22 & (fi22 << ((v3) & 63))) {
                                                            case -232:
                                                                fmt.Println(v4, fi22, v7)
                                                            }
                                                        default:
                                                            if ((42 < 127) || (46 > 2)) {
                                                                fmt.Println(v18, v3, v3)
                                                                var v78 uint16 = 505
                                                                var v79 int8 = 1
                                                                _ = v78
                                                                _ = v79
                                                            }
                                                        }
                                                    }
                                                    switch fi64 {
                                                    case 128:
                                                        var v80 uint = (v3 * (v3 ^ 255))
                                                        _ = v80
                                                    case -148:
                                                        if (-2.25 < v58) {
                                                            fmt.Println(v3, v7)
                                                            fmt.Println(fi64, v7)
                                                        } else {
                                                            var v81 int = fi22
                                                            v82 := GenS6{F0: 1, F1: v3}
                                                            _ = v81
                                                            _ = v82
                                                        }
                                                    default:
                                                        var v83 int = 32768
                                                        _ = v83
                                                    }
                                                }
                                                v3 = (^(v3))
                                                switch fi22 {
                                                case 567:
                                                    v58 -= (v58 - (3.14159 - (v58 - -2.25)))
                                                case -1:
                                                    fmt.Println(fi22, v3)
                                                }
                                                _ = fi64
                                            }
                                            v84 := GenS8{F0: 229, F1: 127}
                                            _ = v84
                                        }
                                        v85 := GenS9{F0: 0, F1: 18446744073709551615, F2: 42}
                                        v86 := GenS10{F0: 1, F1: 469}
                                        _ = v85
                                        _ = v86
                                    }
                                case -2147483648:
                                    var v87 int8 = -19
                                    _ = v87
                                default:
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                                v88 := GenS2{F0: 139}
                                _ = v88
                            } else {
                                switch (fi22 + (fi22 % ((fi22) | 1))) {
                                case 1:
                                    switch (fi22 * fi22) {
                                    case 9223372036854775807:
                                        if true {
                                            for fi89 := 0; fi89 < 12; fi89++ {
                                                v58 += float64(float32(v58))
                                                if ((v58 - v58) > v58) {
                                                    var v90 uint16 = 502
                                                    v91 := GenS3{F0: fi89}
                                                    _ = v90
                                                    _ = v91
                                                } else {
                                                    var v92 uint = (v3 & 0)
                                                    _ = v92
                                                }
                                                _ = fi89
                                            }
                                            v2[(^(v7))] = true
                                            if true {
                                                var v93 bool = ((v7 ^ (v18 & v18)) <= (v7 << ((717) & 63)))
                                                fmt.Println(v93)
                                                v94 := GenS9{F0: 255, F1: 77, F2: 725}
                                                _ = v93
                                                _ = v94
                                            }
                                        } else {
                                            if (!true) {
                                                var v95 float32 = 0.0
                                                var v96 string = "a"
                                                for fi97 := 0; fi97 < 13; fi97++ {
                                                    switch (fi22 | 0) {
                                                    case 128:
                                                        v95 = v95
                                                    case 528:
                                                        if true {
                                                            var v98 string = (v96 + "")
                                                            var v99 string = ((v96 + ((v98 + v98) + v96)) + "")
                                                            _ = v98
                                                            _ = v99
                                                        }
                                                    }
                                                    if false {
                                                        if false {
                                                            v3 -= v3
                                                            v100 := GenS11{F0: 489, F1: 1}
                                                            v95 = -2.25
                                                            _ = v100
                                                        } else {
                                                            fmt.Println(fi22, v18)
                                                            fmt.Println(v7)
                                                        }
                                                        v101 := GenS10{F0: 18446744073709551615, F1: 0}
                                                        var v102 uint16 = 255
                                                        _ = v101
                                                        _ = v102
                                                    } else {
                                                        var v103 int64 = 126
                                                        var v104 int8 = -46
                                                        _ = v103
                                                        _ = v104
                                                    }
                                                    _ = fi97
                                                }
                                                _ = v95
                                                _ = v96
                                            }
                                        }
                                    case 589:
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    }
                                case 32767:
                                    switch ("\n\t" + "") {
                                    case "foo bar":
                                        var v105 uint8 = 27
                                        _ = v105
                                    }
                                case 454:
                                    v106 := GenS12{F0: "a"}
                                    _ = v106
                                }
                                for fi107 := 0; fi107 < 17; fi107++ {
                                    fmt.Println(fi107)
                                    _ = fi107
                                }
                            }
                        }
                    } else {
                        v108 := GenS9{F0: 1, F1: 42, F2: 1}
                        var v109 int8 = 1
                        _ = v108
                        _ = v109
                    }
                default:
                    for fi110 := 0; fi110 < 19; fi110++ {
                        v7 = v7
                        _ = fi110
                    }
                }
                _ = fi22
                _ = v58
            }
            var v111 uint16 = 0
            v112 := GenS13{F0: 0, F1: (v4 == (v4 | (v4 >> (((v3 / (((^(v3))) | 1))) & 63)))), F2: 0}
            _ = v111
            _ = v112
        } else {
            _ = len(v5)
        }
        fmt.Println(v4, v4)
    }
    var v113 int = 32768
    v114 := GenS14{F0: (v18 % (((^(v7))) | 1)), F1: false}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v113:", v113)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v113
    _ = v114
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
