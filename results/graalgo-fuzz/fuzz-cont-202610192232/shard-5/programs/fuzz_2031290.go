package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int8; F1 string }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 bool; F1 uint64 }
type GenS11 struct { F0 int64; F1 int8 }
type GenS12 struct { F0 int; F1 uint32; F2 int64 }
type GenS13 struct { F0 uint8; F1 uint32; F2 uint }
func genFn0(p0 int, p1 uint32) uint32 {
    p1 = p1
    v6 := GenS1{F0: -878}
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := []int{(^(fi7))}
        p1 += (p1 >> ((255) & 63))
        _ = fi7
        _ = v8
    }
    var v9 int16 = 0
    for fi10 := 0; fi10 < 9; fi10++ {
        for fi11 := 0; fi11 < 19; fi11++ {
            fmt.Println(v9)
            var v12 uint16 = 65535
            _ = fi11
            _ = v12
        }
        var v13 int16 = (v9 - (v9 << ((42) & 63)))
        _ = fi10
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v9
    return p1
}
func genFn1(p0 uint8) int32 {
    v14 := GenS1{F0: 42}
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 13; fi16++ {
            v17 := []int{fi15}
            p0 = 1
            var v18 int64 = 1
            switch fi15 {
            case -9223372036854775808:
                if (4294967295 >= 322) {
                    v18 -= (v18 << ((255) & 63))
                    p0 += (p0 << ((1) & 63))
                    if true {
                        if (v18 >= v18) {
                            v17 = append(v17, int(p0))
                            _ = len(v17)
                        }
                    } else {
                        var v19 float64 = 3.14159
                        v20 := GenS1{F0: (^(v18))}
                        p0 -= 2
                        v17 = append(v17, (fi16 >> ((2) & 63)))
                        _ = v19
                        _ = v20
                    }
                } else {
                    _ = len(v17)
                }
            case 32767:
                if false {
                    v17 = append(v17, int(p0))
                    v18 = v18
                } else {
                    if len(v17) > 0 {
                        v17[0] = fi15
                    }
                }
            default:
                var v21 float32 = 1e10
                _ = v21
            }
            var v22 uint8 = p0
            _ = fi16
            _ = v17
            _ = v18
            _ = v22
        }
        for fi23 := 0; fi23 < 2; fi23++ {
            fmt.Println(fi23)
            var v24 string = ("\n\t" + "a")
            v24 = v24
            if false {
                if (!(2147483648 > 2147483648)) {
                    if (!true) {
                        v25 := GenS1{F0: 0}
                        fmt.Println(p0)
                        _ = v25
                    } else {
                        if false {
                            if (v24 > "a") {
                                p0 += (p0 % ((128) | 1))
                                var v26 float32 = -2.25
                                _ = v26
                            }
                            if false {
                                p0 -= p0
                                if false {
                                    v24 = (("\n\t" + (v24 + "hello")) + ((v24 + "hello") + v24))
                                    switch v24 {
                                    case "hello":
                                        v27 := []int{}
                                        _ = v27
                                    case "foo bar":
                                        var v28 float32 = (0.1 - -2.25)
                                        _ = v28
                                    default:
                                        switch fi23 {
                                        case 0:
                                            if (1 < -57) {
                                                v24 = "hello"
                                            } else {
                                                if (true && (!(!true))) {
                                                    v24 = v24
                                                    var v29 int16 = 1
                                                    var v30 uint32 = 441
                                                    _ = v29
                                                    _ = v30
                                                }
                                                v31 := GenS2{F0: (false && true)}
                                                var v32 int8 = 127
                                                _ = v31
                                                _ = v32
                                            }
                                        default:
                                            p0 += (^(p0))
                                        }
                                    }
                                } else {
                                    v33 := []int{fi15, 183, fi15, (fi23 << ((1) & 63))}
                                    if true {
                                        fmt.Println(v24, fi23)
                                    }
                                    _ = v33
                                }
                            } else {
                                switch -745 {
                                case 32768:
                                    fmt.Println(v24, p0, fi23)
                                case 674:
                                    v34 := GenS3{F0: 0, F1: 1}
                                    _ = v34
                                default:
                                    if (!true) {
                                        p0 = p0
                                        if false {
                                            v35 := []int{}
                                            _ = len(v35)
                                            _ = v35
                                        } else {
                                            fmt.Println(fi15, p0)
                                            v36 := GenS1{F0: -9223372036854775808}
                                            _ = v36
                                        }
                                        if true {
                                            v37 := GenS4{F0: 100}
                                            var v38 uint32 = uint32(p0)
                                            _ = v37
                                            _ = v38
                                        }
                                    } else {
                                        if true {
                                            var v39 int64 = 9223372036854775807
                                            v39 -= (^(v39))
                                            _ = v39
                                        } else {
                                            var v40 int = (fi23 ^ int(p0))
                                            _ = v40
                                        }
                                    }
                                }
                                switch 877 {
                                case 236:
                                    var v41 float32 = 1.5
                                    p0 += 255
                                    p0 = p0
                                    p0 += p0
                                    _ = v41
                                case 42:
                                    switch "a" {
                                    case "foo bar":
                                        if (255 >= 357) {
                                            var v42 float64 = float64(0.1)
                                            var v43 string = (v24 + "hello")
                                            v44 := []int{fi15}
                                            _ = v42
                                            _ = v43
                                            _ = v44
                                        } else {
                                            if (1000 >= 1000) {
                                                v24 = v24
                                                if true {
                                                    v45 := GenS5{F0: 2}
                                                    _ = v45
                                                } else {
                                                    fmt.Println(v24, v24)
                                                }
                                            }
                                        }
                                    case "":
                                        v24 = "hello"
                                    case "a":
                                        var v46 int64 = 0
                                        _ = v46
                                    default:
                                        if true {
                                            var v47 uint16 = 32767
                                            _ = v47
                                        } else {
                                            var v48 uint = 42
                                            v49 := []int{-2147483648, -542}
                                            _ = v48
                                            _ = v49
                                        }
                                    }
                                case 0:
                                    v50 := GenS2{F0: (v24 == v24)}
                                    v24 = "hello"
                                    v51 := GenS6{F0: 0}
                                    _ = v50
                                    _ = v51
                                default:
                                    if false {
                                        var v52 float32 = -2.25
                                        if (true && ((int8(p0) == 127) && true)) {
                                            p0 -= 255
                                        }
                                        v53 := GenS7{F0: 963}
                                        _ = v52
                                        _ = v53
                                    }
                                }
                                if (!(204 != 1)) {
                                    if (85 > 42) {
                                        switch v24 {
                                        case "foo bar":
                                            var v54 uint32 = 536
                                            _ = v54
                                        case "\n\t":
                                            fmt.Println(p0, fi15, p0)
                                        case "":
                                            var v55 int64 = 9223372036854775807
                                            _ = v55
                                        }
                                        var v56 float64 = (1.5 - ((-2.25 + ((1e10 / 0.1) + 1e10)) + 0.0))
                                        _ = v56
                                    }
                                    v57 := []int{(^(fi23)), (fi15 % (((fi23 << ((42) & 63))) | 1)), (^(fi23)), (^(fi15))}
                                    _ = v57
                                }
                            }
                            switch (fi23 << ((4294967295) & 63)) {
                            case -169:
                                var v58 bool = false
                                _ = v58
                            case -9223372036854775808:
                                if false {
                                    var v59 uint = 42
                                    if false {
                                        v59 -= v59
                                        var v60 uint16 = 0
                                        if true {
                                            if false {
                                                var v61 uint16 = (v60 + v60)
                                                fmt.Println(fi23, v24)
                                                var v62 bool = false
                                                v63 := []int{(fi23 & fi23), fi23, (fi23 << (((v59 % ((900) | 1))) & 63))}
                                                _ = v61
                                                _ = v62
                                                _ = v63
                                            }
                                        } else {
                                            fmt.Println(p0)
                                            v64 := []int{(fi15 | fi15), 83, 32767}
                                            switch v24 {
                                            case "foo bar":
                                                v64 = append(v64, fi15)
                                            case "hello":
                                                fmt.Println(v24)
                                            case "":
                                                var v65 int = fi15
                                                _ = v65
                                            }
                                            _ = v64
                                        }
                                        _ = v60
                                    } else {
                                        v24 = (v24 + "hello")
                                        var v66 float32 = (-2.25 - (((-2.25 * float32((((0.1 * -2.25) * 1e10) * ((float64((float32((3.14159 * float64(-2.25))) * 1e10)) * 3.14159) + ((((3.14159 * ((0.0 + (3.14159 * 0.1)) * 0.0)) - -2.25) + 0.0) * 0.0))))) - 0.1) * 0.1))
                                        _ = v66
                                    }
                                    fmt.Println(fi23, fi23)
                                    _ = v59
                                }
                            }
                        }
                    }
                    var v67 uint32 = 1
                    v68 := GenS8{F0: 106, F1: v24}
                    _ = v67
                    _ = v68
                }
            }
            var v69 float32 = -2.25
            fmt.Println(v69)
            var v70 int16 = 1
            fmt.Println(v24, p0)
            _ = fi23
            _ = v24
            _ = v69
            _ = v70
        }
        _ = fi15
    }
    for fi71 := 0; fi71 < 18; fi71++ {
        v72 := []int{265}
        p0 += p0
        _ = fi71
        _ = v72
    }
    switch 32767 {
    case 128:
        v73 := []int{-205, 42, 128}
        _ = v73
    }
    for fi74 := 0; fi74 < 8; fi74++ {
        v75 := GenS9{F0: (302 >= 33)}
        _ = fi74
        _ = v75
    }
    _ = p0
    _ = v14
    return 2147483647
}

func main() {
    var v1 string = "\n\t"
    var v2 float32 = (3.14159 + 1.5)
    v3 := []int{}
    var v4 uint32 = 2
    var v5 uint64 = 4294967295
    v76 := GenS3{F0: 754, F1: -2147483648}
    var v77 int = 621
    var v78 string = v1
    for fi79 := 0; fi79 < 3; fi79++ {
        switch v78 {
        case "a":
            v80 := GenS3{F0: 3, F1: 1}
            _ = v80
        case "foo bar":
            if true {
                switch fi79 {
                case -9223372036854775808:
                    if (true && true) {
                        v77 = fi79
                        fmt.Println(v78, v77)
                    } else {
                        switch ("" + v1) {
                        case "a":
                            v81 := GenS10{F0: true, F1: 0}
                            _ = v81
                        case "\n\t":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case "hello":
                            var v82 float32 = v2
                            _ = v82
                        default:
                            v78 = v78
                        }
                        if (168 < 32767) {
                            if ((-236 != -2147483648) || (127 == 100)) {
                                switch (fi79 + (fi79 | 42)) {
                                case -1:
                                    if false {
                                        for fi83 := 0; fi83 < 5; fi83++ {
                                            switch 929 {
                                            case -740:
                                                var v84 int8 = -1
                                                _ = v84
                                            case -381:
                                                _ = len(v3)
                                                fmt.Println(fi79, v4)
                                                var v85 int64 = -947
                                                v1 = v1
                                                v5 += 2
                                                _ = v85
                                            default:
                                                v3 = append(v3, v77)
                                            }
                                            var v86 int8 = 0
                                            _ = fi83
                                            _ = v86
                                        }
                                        v3 = append(v3, (-(fi79)))
                                        for fi87 := 0; fi87 < 6; fi87++ {
                                            fmt.Println(v77)
                                            _ = fi87
                                        }
                                    }
                                default:
                                    var v88 float32 = (v2 * 1e10)
                                    _ = v88
                                }
                                fmt.Println(fi79)
                            } else {
                                if true {
                                    for fi89 := 0; fi89 < 19; fi89++ {
                                        fmt.Println(v1, v2, v1)
                                        fmt.Println(fi89)
                                        var v90 int = (v77 >> ((979) & 63))
                                        _ = fi89
                                        _ = v90
                                    }
                                    switch 220 {
                                    case 127:
                                        v77 = (v77 & (fi79 + 32768))
                                    case -1:
                                        if true {
                                            switch ((v78 + "hello") + ("hello" + v1)) {
                                            case "a":
                                                var v91 bool = (false && (!false))
                                                _ = v91
                                            case "":
                                                v1 = (("a" + ((("\n\t" + v78) + v1) + "foo bar")) + (v1 + v1))
                                            case "hello":
                                                var v92 uint = 592
                                                _ = v92
                                            default:
                                                fmt.Println(v77, v5)
                                            }
                                            switch v77 {
                                            case 32767:
                                                if (genFn0((v77 % ((v77) | 1)), 934) != (v4 - 2)) {
                                                    var v93 uint = 255
                                                    var v94 int = fi79
                                                    _ = v93
                                                    _ = v94
                                                }
                                            default:
                                                var v95 uint64 = 688
                                                _ = v95
                                            }
                                        }
                                    case 9223372036854775807:
                                        if true {
                                            if ((425 < 0) || true) {
                                                switch "a" {
                                                case "foo bar":
                                                    fmt.Println(v78, fi79)
                                                case "\n\t":
                                                    v77 = -1
                                                default:
                                                    _ = len(v3)
                                                }
                                            }
                                        }
                                    default:
                                        fmt.Println(v78, v4, v5)
                                    }
                                    v3 = append(v3, v77)
                                } else {
                                    fmt.Println(v78, fi79, fi79)
                                    if true {
                                        if ((float32(3.14159) / 3.14159) >= v2) {
                                            v5 += (v5 ^ v5)
                                        } else {
                                            _ = len(v3)
                                            for fi96 := 0; fi96 < 4; fi96++ {
                                                var v97 int16 = -32768
                                                _ = fi96
                                                _ = v97
                                            }
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                }
                            }
                        }
                    }
                case 9223372036854775807:
                    v78 = ("" + (v1 + "\n\t"))
                case 42:
                    switch 0 {
                    case 42:
                        v4 = 65535
                    case -2147483648:
                        if true {
                            var v98 int64 = -1
                            _ = v98
                        }
                    case 2147483647:
                        v99 := GenS6{F0: 42}
                        _ = v99
                    }
                default:
                    v77 += (^(fi79))
                }
            } else {
                var v100 float32 = 1.5
                _ = v100
            }
        case "":
            var v101 bool = false
            _ = v101
        default:
            v102 := GenS11{F0: -9223372036854775808, F1: 42}
            _ = v102
        }
        _ = fi79
    }
    var v103 uint8 = 1
    for fi104 := 0; fi104 < 17; fi104++ {
        var v105 uint = 4294967295
        _ = fi104
        _ = v105
    }
    switch (-(v77)) {
    case 2147483647:
        var v106 int64 = 42
        _ = v106
    case 837:
        if (633 != -9223372036854775808) {
            fmt.Println(v103, v2)
        }
    case 1:
        var v107 float64 = 0.0
        _ = v107
    }
    v108 := GenS10{F0: false, F1: (^(v5))}
    switch 209 {
    case 42:
        fmt.Println(v77)
    }
    var v109 uint64 = v5
    switch v77 {
    case 32767:
        var v110 bool = (0 <= -996)
        _ = v110
    case 822:
        if (3.14159 != 0.1) {
            fmt.Println(v5, v2)
        } else {
            var v111 float64 = 1.5
            fmt.Println(v78)
            _ = v111
        }
    case -699:
        var v112 uint = 0
        _ = v112
    default:
        for fi113 := 0; fi113 < 4; fi113++ {
            fmt.Println(v77, v109)
            _ = len(v3)
            for fi114 := 0; fi114 < 5; fi114++ {
                v115 := GenS4{F0: 100}
                _ = fi114
                _ = v115
            }
            _ = fi113
        }
    }
    v116 := GenS12{F0: v77, F1: 0, F2: -1}
    var v117 uint = 203
    if ((0.1 + (-2.25 * -2.25)) != -2.25) {
        for fi118 := 0; fi118 < 8; fi118++ {
            v103 = v103
            v103 = 166
            _ = fi118
        }
        var v119 float32 = (float32(1.5) / (v2 * 1.5))
        _ = len(v3)
        _ = v119
    } else {
        v3 = append(v3, v77)
        v120 := GenS13{F0: (v103 + uint8(v109)), F1: v4, F2: (v117 / ((v117) | 1))}
        var v121 float64 = float64(float32((float64((0.0 * v2)) - 0.0)))
        _ = v120
        _ = v121
    }
    for fi122 := 0; fi122 < 2; fi122++ {
        if len(v3) > 0 {
            v3[0] = 9223372036854775807
        }
        _ = fi122
    }
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v109:", v109)
    fmt.Println("v117:", v117)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v103
    _ = v108
    _ = v109
    _ = v116
    _ = v117
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v76
    _ = v77
    _ = v78
}
