package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int32; F1 int16 }
type GenS9 struct { F0 int; F1 uint8 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int16; F1 uint }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int32; F1 uint64; F2 int }
type GenS14 struct { F0 int32; F1 uint8 }
type GenS15 struct { F0 int32; F1 int64 }
type GenS16 struct { F0 uint8; F1 int; F2 int64 }
type GenS17 struct { F0 uint32; F1 uint }
type GenS18 struct { F0 int8 }
func genFn0(p0 uint) int64 {
    for fi7 := 0; fi7 < 18; fi7++ {
        switch 0 {
        case 42:
            var v8 uint32 = uint32(p0)
            _ = v8
        case 127:
            v9 := GenS1{F0: 783, F1: 1000, F2: -128}
            _ = v9
        case 2147483647:
            for fi10 := 0; fi10 < 12; fi10++ {
                if true {
                    p0 = p0
                }
                _ = fi10
            }
        }
        _ = fi7
    }
    if true {
        v11 := GenS2{F0: 379}
        for fi12 := 0; fi12 < 7; fi12++ {
            switch (^(fi12)) {
            case 844:
                for fi13 := 0; fi13 < 4; fi13++ {
                    fmt.Println(fi12, fi13, fi12)
                    _ = fi13
                }
            default:
                var v14 uint32 = 42
                _ = v14
            }
            if (9223372036854775807 >= 127) {
                var v15 int8 = 127
                var v16 uint = (p0 & 18446744073709551615)
                _ = v15
                _ = v16
            } else {
                v17 := GenS3{F0: 126, F1: 775}
                fmt.Println(p0)
                v18 := GenS4{F0: -32768}
                _ = v17
                _ = v18
            }
            v19 := []int{(^(fi12))}
            _ = fi12
            _ = v19
        }
        var v20 uint8 = 127
        _ = v11
        _ = v20
    } else {
        var v21 int16 = -1
        v22 := []int{}
        _ = v21
        _ = v22
    }
    _ = p0
    return 0
}

func main() {
    var v1 int16 = -1
    var v2 float32 = 1e10
    var v3 uint8 = 127
    v4 := []uint32{}
    v5 := []int{int(v1)}
    var v6 int32 = 42
    if false {
        var v23 float32 = v2
        for fi24 := 0; fi24 < 17; fi24++ {
            for fi25 := 0; fi25 < 3; fi25++ {
                var v26 string = "\n\t"
                _ = fi25
                _ = v26
            }
            v5 = append(v5, (fi24 - (-(fi24))))
            if len(v5) > 0 {
                v5[0] = fi24
            }
            _ = fi24
        }
        _ = v23
    } else {
        var v27 int8 = 42
        var v28 uint = 558
        v28 -= 4294967295
        _ = v27
        _ = v28
    }
    switch -2147483648 {
    case -9223372036854775808:
        var v29 int32 = -610
        _ = v29
    case -23:
        for fi30 := 0; fi30 < 18; fi30++ {
            if true {
                v31 := GenS2{F0: 4294967295}
                var v32 int16 = -32768
                var v33 string = ("hello" + "")
                _ = v31
                _ = v32
                _ = v33
            } else {
                _ = len(v5)
                switch -692 {
                case -9223372036854775808:
                    v34 := GenS5{F0: -1}
                    _ = v34
                default:
                    if false {
                        var v35 uint64 = 18446744073709551615
                        _ = v35
                    }
                }
                if len(v5) > 0 {
                    v5[0] = fi30
                }
            }
            for fi36 := 0; fi36 < 4; fi36++ {
                v6 = v6
                _ = fi36
            }
            _ = fi30
        }
    case 4:
        if len(v4) > 0 {
            v4[0] = 2147483648
        }
    default:
        v2 -= (((v2 / (float32((0.0 * (0.1 * (-2.25 / ((1.5 + -2.25) / float64((v2 - v2))))))) * v2)) + v2) + 1.5)
    }
    for fi37 := 0; fi37 < 4; fi37++ {
        v1 = v1
        v38 := GenS6{F0: (v6 << ((2) & 63))}
        _ = fi37
        _ = v38
    }
    if false {
        for fi39 := 0; fi39 < 10; fi39++ {
            v40 := GenS2{F0: 1}
            if (928 != 0) {
                v2 -= 1e10
                fmt.Println(v6, v1)
            } else {
                var v41 uint = 410
                _ = v41
            }
            v42 := GenS7{F0: 0}
            _ = fi39
            _ = v40
            _ = v42
        }
        for fi43 := 0; fi43 < 10; fi43++ {
            v1 -= v1
            _ = fi43
        }
    } else {
        if true {
            if true {
                for fi44 := 0; fi44 < 16; fi44++ {
                    for fi45 := 0; fi45 < 13; fi45++ {
                        v6 = (^(v6))
                        var v46 int16 = -1
                        _ = fi45
                        _ = v46
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v3 = (v3 % (((v3 * v3)) | 1))
                    _ = fi44
                }
                switch 910 {
                case 32767:
                    if (genFn0(4294967295) != -1) {
                        fmt.Println(v3)
                        if (42 != 2147483648) {
                            v47 := GenS7{F0: 4294967295}
                            _ = v47
                        } else {
                            v48 := GenS8{F0: (v6 / (((v6 / ((v6) | 1))) | 1)), F1: -1}
                            _ = v48
                        }
                    }
                case 343:
                    v3 += v3
                case -9223372036854775808:
                    fmt.Println(v2, v1, v3)
                default:
                    switch int(v3) {
                    case 32767:
                        if true {
                            switch ("\n\t" + "\n\t") {
                            case "":
                                v1 = (-(v1))
                            }
                            if (!false) {
                                if len(v4) > 0 {
                                    v4[0] = 4294967295
                                }
                                for fi49 := 0; fi49 < 11; fi49++ {
                                    v3 -= v3
                                    var v50 uint = 0
                                    _ = fi49
                                    _ = v50
                                }
                            }
                            v51 := GenS2{F0: 2}
                            _ = v51
                        }
                    case -9223372036854775808:
                        v52 := GenS9{F0: 32767, F1: 181}
                        _ = v52
                    }
                }
                v53 := GenS9{F0: 128, F1: v3}
                _ = v53
            }
        } else {
            if len(v4) > 0 {
                v4[0] = 65535
            }
        }
        v54 := GenS10{F0: -1}
        if false {
            if false {
                var v55 int = int(v1)
                v55 += (v55 % ((32768) | 1))
                _ = v55
            } else {
                v56 := GenS11{F0: -32768, F1: 255}
                var v57 float64 = 1.5
                _ = v56
                _ = v57
            }
            switch ("" + ("\n\t" + ("foo bar" + (("foo bar" + (("hello" + "foo bar") + (("" + "a") + ""))) + "a")))) {
            case "a":
                var v58 uint64 = 220
                _ = v58
            case "hello":
                var v59 int16 = (v1 % ((32767) | 1))
                _ = v59
            }
            var v60 int = 1
            _ = v60
        }
        _ = v54
    }
    if (0.1 <= 1e10) {
        var v61 int32 = int32(v3)
        if len(v4) > 0 {
            v4[0] = 915
        }
        if (!true) {
            var v62 string = "foo bar"
            if false {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if (((42 == 32767) || ((v1 << ((0) & 63)) >= -636)) && ((1.5 / (((float32(-2.25) - 1.5) / -2.25) * (((((((float32((-2.25 / (((-2.25 * float64(3.14159)) - -2.25) - 1e10))) * v2) / 1.5) - 1.5) * 0.0) * v2) - 0.1) - (-2.25 + v2)))) > v2)) {
                    for fi63 := 0; fi63 < 11; fi63++ {
                        switch "foo bar" {
                        case "a":
                            switch fi63 {
                            case -1:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case -245:
                                if (82 < -128) {
                                    if false {
                                        v2 = v2
                                        _ = len(v5)
                                        v64 := GenS12{F0: v1}
                                        _ = v64
                                    } else {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        if true {
                                            var v65 uint = 1
                                            var v66 uint = 255
                                            var v67 uint8 = (v3 % ((218) | 1))
                                            _ = v65
                                            _ = v66
                                            _ = v67
                                        } else {
                                            if false {
                                                var v68 int = fi63
                                                v68 = -9223372036854775808
                                                v68 -= -581
                                                _ = v68
                                            } else {
                                                v4 = append(v4, 1)
                                                for fi69 := 0; fi69 < 5; fi69++ {
                                                    var v70 float64 = (0.1 * 1e10)
                                                    var v71 string = (("a" + (v62 + (v62 + (((v62 + "foo bar") + (v62 + v62)) + ("foo bar" + "a"))))) + "foo bar")
                                                    if false {
                                                        if true {
                                                            if ("a" < (v62 + (("" + v71) + "hello"))) {
                                                                var v72 int32 = 2147483647
                                                                v2 += float32(0.1)
                                                                var v73 float32 = 3.14159
                                                                v1 -= -1
                                                                _ = v72
                                                                _ = v73
                                                            }
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            v74 := GenS13{F0: (v61 | v6), F1: 831, F2: fi69}
                                                            _ = v74
                                                        } else {
                                                            v75 := GenS14{F0: v6, F1: v3}
                                                            var v76 int64 = 0
                                                            _ = v75
                                                            _ = v76
                                                        }
                                                        v77 := GenS4{F0: v1}
                                                        _ = v77
                                                    }
                                                    v78 := GenS15{F0: 42, F1: 0}
                                                    _ = fi69
                                                    _ = v70
                                                    _ = v71
                                                    _ = v78
                                                }
                                            }
                                        }
                                    }
                                    var v79 float32 = v2
                                    _ = v79
                                } else {
                                    fmt.Println(v62)
                                    for fi80 := 0; fi80 < 10; fi80++ {
                                        if true {
                                            if (((v2 / 0.1) + (v2 + v2)) <= (v2 * v2)) {
                                                v81 := GenS2{F0: 18446744073709551615}
                                                v62 = v62
                                                var v82 int16 = v1
                                                _ = v81
                                                _ = v82
                                            }
                                            v1 = 32767
                                            v83 := GenS16{F0: 136, F1: fi80, F2: 0}
                                            var v84 int8 = int8(fi80)
                                            v5 = append(v5, fi80)
                                            v4 = append(v4, 1)
                                            _ = v83
                                            _ = v84
                                        } else {
                                            if false {
                                                var v85 float32 = (v2 * v2)
                                                var v86 int16 = 753
                                                var v87 string = (((v62 + v62) + ("" + v62)) + "\n\t")
                                                _ = v85
                                                _ = v86
                                                _ = v87
                                            }
                                        }
                                        var v88 uint = 9
                                        _ = fi80
                                        _ = v88
                                    }
                                    for fi89 := 0; fi89 < 12; fi89++ {
                                        v1 -= (v1 - (v1 - v1))
                                        switch -967 {
                                        case -9223372036854775808:
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        case 127:
                                            var v90 bool = false
                                            _ = v90
                                        case -452:
                                            var v91 bool = true
                                            _ = v91
                                        default:
                                            fmt.Println(v6, fi63, v62)
                                        }
                                        var v92 float32 = v2
                                        var v93 int64 = 42
                                        _ = fi89
                                        _ = v92
                                        _ = v93
                                    }
                                }
                            }
                        case "\n\t":
                            if len(v4) > 0 {
                                v4[0] = 2147483647
                            }
                        case "foo bar":
                            var v94 int64 = int64(v61)
                            _ = v94
                        }
                        _ = fi63
                    }
                    v95 := GenS17{F0: 249, F1: 2}
                    for fi96 := 0; fi96 < 11; fi96++ {
                        var v97 uint64 = 0
                        v98 := GenS18{F0: 1}
                        _ = fi96
                        _ = v97
                        _ = v98
                    }
                    _ = v95
                } else {
                    for fi99 := 0; fi99 < 6; fi99++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v100 int64 = -9223372036854775808
                        fmt.Println(fi99, v6)
                        _ = fi99
                        _ = v100
                    }
                    v4 = append(v4, 0)
                    switch 32768 {
                    case 128:
                        var v101 int32 = (-(v61))
                        _ = v101
                    }
                }
            }
            _ = v62
        }
        _ = v61
    } else {
        var v102 int = -9223372036854775808
        var v103 float64 = 0.0
        _ = v102
        _ = v103
    }
    for fi104 := 0; fi104 < 7; fi104++ {
        fmt.Println(v2)
        _ = len(v5)
        _ = fi104
    }
    var v105 uint16 = 42
    for fi106 := 0; fi106 < 8; fi106++ {
        var v107 uint8 = (v3 * (^(v3)))
        v4 = append(v4, 46)
        _ = fi106
        _ = v107
    }
    for fi108 := 0; fi108 < 20; fi108++ {
        var v109 float32 = (((-2.25 - v2) + v2) + v2)
        if true {
            var v110 float32 = v2
            _ = v110
        }
        _ = fi108
        _ = v109
    }
    for fi111 := 0; fi111 < 11; fi111++ {
        if false {
            _ = len(v4)
            for fi112 := 0; fi112 < 10; fi112++ {
                if true {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = fi112
            }
        }
        var v113 float64 = 1.5
        var v114 int32 = (v6 & v6)
        _ = fi111
        _ = v113
        _ = v114
    }
    fmt.Println(v2)
    fmt.Println(v105)
    v1 = v1
    for fi115 := 0; fi115 < 12; fi115++ {
        if (true && (!false)) {
            if (-95 == 42) {
                var v116 int16 = v1
                v6 = v6
                var v117 int64 = 0
                _ = v116
                _ = v117
            }
            v3 = 2
            v6 -= v6
        }
        _ = fi115
    }
    var v118 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v105:", v105)
    fmt.Println("v118:", v118)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v105
    _ = v118
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
