package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 int64; F1 uint8; F2 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint16; F1 uint16; F2 int }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 bool; F1 int64 }
type GenS12 struct { F0 uint; F1 uint; F2 int16 }
type GenS13 struct { F0 bool; F1 int; F2 string }
type GenS14 struct { F0 uint; F1 int; F2 uint64 }
func genFn0(p0 uint16, p1 string) bool {
    var v5 uint64 = 42
    v6 := []int{1}
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 int64 = -1
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 1; fi10++ {
            switch "foo bar" {
            case "a":
                fmt.Println(p1, fi10, v5)
            }
            if false {
                if (0 != 1) {
                    var v11 int64 = 42
                    v5 = (v5 % ((v5) | 1))
                    _ = v11
                } else {
                    p0 = p0
                    fmt.Println(p0, fi10, p1)
                    v6 = append(v6, (fi10 / ((fi7) | 1)))
                    var v12 bool = true
                    _ = v12
                }
                var v13 uint16 = 2
                if true {
                    var v14 int8 = 9
                    v14 -= 0
                    if (true || false) {
                        _ = len(v6)
                        switch ("foo bar" + p1) {
                        case "a":
                            p0 = v13
                            v15 := GenS1{F0: v13}
                            _ = v15
                        case "":
                            var v16 float64 = 1.5
                            _ = v16
                        default:
                            if ("foo bar" < (p1 + "\n\t")) {
                                fmt.Println(p0)
                                var v17 float64 = 1.5
                                switch (fi10 >> ((1) & 63)) {
                                case 776:
                                    v13 = p0
                                default:
                                    v5 = v5
                                }
                                v6 = append(v6, (fi7 & 128))
                                fmt.Println(p0, v14, p0)
                                _ = v17
                            }
                        }
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                    switch (-(fi7)) {
                    case -477:
                        v6 = append(v6, int(v5))
                        var v18 uint = 18446744073709551615
                        _ = v18
                    case -597:
                        var v19 int64 = -943
                        _ = v19
                    case 127:
                        var v20 int8 = -62
                        switch fi7 {
                        case -466:
                            switch fi7 {
                            case -811:
                                _ = len(v6)
                            case 4294967295:
                                var v21 bool = ((!((false || true) && (553 != -1))) && ((v20 < v14) && (!false)))
                                _ = v21
                            }
                        case -169:
                            p1 = p1
                        case -2147483648:
                            switch p1 {
                            case "hello":
                                v22 := GenS1{F0: 1}
                                _ = v22
                            case "\n\t":
                                var v23 int8 = v20
                                _ = v23
                            }
                        default:
                            switch fi10 {
                            case 1:
                                v24 := GenS2{F0: 0, F1: (p0 + uint16(v5))}
                                _ = v24
                            case 9223372036854775807:
                                fmt.Println(v14, p1, p0)
                            case 128:
                                switch ("" + (p1 + p1)) {
                                case "":
                                    var v25 int64 = 9223372036854775807
                                    _ = v25
                                case "foo bar":
                                    v14 = v14
                                case "\n\t":
                                    var v26 int64 = -1
                                    _ = v26
                                default:
                                    p0 += p0
                                }
                            default:
                                var v27 float32 = 3.14159
                                _ = v27
                            }
                        }
                        _ = v20
                    }
                    _ = v14
                }
                _ = v13
            } else {
                switch (p1 + "foo bar") {
                case "a":
                    p0 = p0
                case "hello":
                    if false {
                        v5 = (v5 ^ (v5 / (((v5 >> ((255) & 63))) | 1)))
                        p1 = p1
                        fmt.Println(p1)
                    }
                case "":
                    if ((!true) || false) {
                        var v28 uint64 = (v5 >> ((300) & 63))
                        _ = v28
                    } else {
                        var v29 int64 = -9223372036854775808
                        fmt.Println(v5, fi10, v29)
                        _ = v29
                    }
                }
            }
            v5 += 551
            _ = fi10
        }
        _ = fi7
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if (!(255 >= 4294967295)) {
        _ = len(v6)
        var v30 bool = false
        _ = v30
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (950 <= 2)
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    var v4 int = -824
    v31 := []int{}
    var v32 uint = 255
    var v33 int8 = -123
    if true {
        var v34 float32 = (-2.25 * float32(0.0))
        _ = v34
    } else {
        v32 = (^(v32))
        v35 := v2[v4]
        v36 := GenS2{F0: 380, F1: 32767}
        _ = v35
        _ = v36
    }
    var v37 string = "a"
    v38 := GenS2{F0: 42, F1: 0}
    v31 = append(v31, 60)
    for fi39 := 0; fi39 < 14; fi39++ {
        v40 := GenS1{F0: 255}
        _ = fi39
        _ = v40
    }
    var v41 bool = (42 <= -1)
    delete(v3, -721)
    fmt.Println(v37, v33, v4)
    v32 -= v32
    for fi42 := 0; fi42 < 1; fi42++ {
        if (v41 && false) {
            for fi43 := 0; fi43 < 2; fi43++ {
                if true {
                    var v44 float32 = 0.0
                    _ = v44
                }
                if (v41 && (!v41)) {
                    var v45 uint = v32
                    v32 = (v32 * v45)
                    if v41 {
                        v41 = v41
                        var v46 int8 = v33
                        _ = v46
                    }
                    _ = v45
                }
                _ = fi43
            }
            var v47 float64 = -2.25
            _ = v47
        }
        var v48 uint8 = 83
        _ = fi42
        _ = v48
    }
    switch v4 {
    case -1:
        switch 32767 {
        case 1:
            var v49 uint = 18446744073709551615
            _ = v49
        case 822:
            var v50 string = v37
            _ = v50
        }
    default:
        v4 = 128
    }
    var v51 float64 = 0.1
    var v52 uint = v32
    fmt.Println(v33, v32)
    switch v4 {
    case 1:
        if ((false || v41) || (!(!v41))) {
            for fi53 := 0; fi53 < 11; fi53++ {
                var v54 int16 = 0
                if false {
                    v55 := GenS1{F0: 467}
                    if false {
                        var v56 string = v37
                        for fi57 := 0; fi57 < 5; fi57++ {
                            fmt.Println(v33, v4)
                            if (!(274 <= 1)) {
                                v58 := GenS1{F0: 42}
                                _ = v58
                            } else {
                                var v59 uint64 = 42
                                if (false || v41) {
                                    if (42 == 0) {
                                        v33 = (v33 >> (((v52 + v32)) & 63))
                                        var v60 float64 = v51
                                        if false {
                                            switch (fi53 & fi57) {
                                            case 127:
                                                v61 := GenS1{F0: 255}
                                                _ = v61
                                            case 32767:
                                                v62 := GenS3{F0: v32, F1: 2}
                                                if len(v31) > 0 {
                                                    v31[0] = 42
                                                }
                                                var v63 bool = (-38 == (fi57 / (((v4 / (((fi53 - fi57)) | 1))) | 1)))
                                                _ = v62
                                                _ = v63
                                            }
                                            if len(v31) > 0 {
                                                v31[0] = -260
                                            }
                                        } else {
                                            v64 := GenS1{F0: uint16(v52)}
                                            v65 := GenS4{F0: 163, F1: 1, F2: v41}
                                            var v66 float64 = ((v60 / v60) / 3.14159)
                                            if v41 {
                                                v54 -= v54
                                                v67 := GenS5{F0: (v33 / (((v33 * (v33 * v33))) | 1))}
                                                _ = v67
                                            }
                                            if v41 {
                                                v51 = v60
                                                if v41 {
                                                    switch v37 {
                                                    case "\n\t":
                                                        var v68 int = int(v33)
                                                        _ = v68
                                                    }
                                                    _ = len(v31)
                                                    if (fi53 <= -482) {
                                                        var v69 int8 = (v33 << (((^(v52))) & 63))
                                                        fmt.Println(fi57, v51, v33)
                                                        _ = v69
                                                    }
                                                }
                                            } else {
                                                switch "" {
                                                case "a":
                                                    var v70 uint = uint(v59)
                                                    _ = v70
                                                case "foo bar":
                                                    if len(v31) > 0 {
                                                        v31[0] = v4
                                                    }
                                                case "\n\t":
                                                    v71 := GenS6{F0: v33}
                                                    _ = v71
                                                default:
                                                    var v72 uint64 = 18446744073709551615
                                                    _ = v72
                                                }
                                                if len(v31) >= 2 {
                                                    v31 = v31[1:len(v31)]
                                                }
                                                if (fi57 != (fi57 | v4)) {
                                                    if v41 {
                                                        if (330 == v59) {
                                                            switch v37 {
                                                            case "a":
                                                                fmt.Println(fi57)
                                                            case "":
                                                                var v73 uint = (v32 - 0)
                                                                _ = v73
                                                            }
                                                            v54 = (v54 << ((v52) & 63))
                                                        } else {
                                                            if true {
                                                                fmt.Println(v56, v59)
                                                                var v74 string = (v37 + ((v56 + ("" + (v37 + ("a" + ((("a" + ((("hello" + ((v56 + (v37 + "")) + "hello")) + v37) + "foo bar")) + "\n\t") + "\n\t"))))) + v56))
                                                                var v75 uint32 = 4294967295
                                                                _ = v74
                                                                _ = v75
                                                            } else {
                                                                switch fi57 {
                                                                case 661:
                                                                    var v76 uint16 = 193
                                                                    _ = v76
                                                                default:
                                                                    if false {
                                                                        v77 := GenS7{F0: v37}
                                                                        _ = v77
                                                                    }
                                                                }
                                                                switch v37 {
                                                                case "hello":
                                                                    if (fi57 > int(v32)) {
                                                                        var v78 uint16 = 42
                                                                        fmt.Println(v78, v54, v37)
                                                                        _ = v78
                                                                    } else {
                                                                        var v79 string = v37
                                                                        if v41 {
                                                                            v32 = (v52 + v52)
                                                                            if v41 {
                                                                                if false {
                                                                                    var v80 int32 = -1
                                                                                    if (("" + v37) > ("\n\t" + (v56 + (v56 + v79)))) {
                                                                                        if v41 {
                                                                                            if v41 {
                                                                                                if len(v31) > 0 {
                                                                                                    v31[0] = fi53
                                                                                                }
                                                                                                v81 := GenS8{F0: v52}
                                                                                                _ = v81
                                                                                            } else {
                                                                                                if (3.14159 < float32(((3.14159 - (1.5 + float64(1.5))) - (3.14159 + 1.5)))) {
                                                                                                    v82 := GenS9{F0: uint16(v59), F1: 1, F2: int(v52)}
                                                                                                    _ = v82
                                                                                                } else {
                                                                                                    if (v33 <= (v33 ^ v33)) {
                                                                                                        var v83 int = (^(fi57))
                                                                                                        var v84 uint = (v32 % ((v52) | 1))
                                                                                                        _ = v83
                                                                                                        _ = v84
                                                                                                    }
                                                                                                }
                                                                                                v85 := GenS4{F0: -1, F1: 42, F2: (!false)}
                                                                                                if len(v31) > 0 {
                                                                                                    v31[0] = fi57
                                                                                                }
                                                                                                _ = v85
                                                                                            }
                                                                                            v86 := GenS10{F0: 0}
                                                                                            fmt.Println(v37, fi57, v4)
                                                                                            fmt.Println(v60, v51, v37)
                                                                                            if (!true) {
                                                                                                fmt.Println(v56, v51, v4)
                                                                                                switch (v4 >> ((v52) & 63)) {
                                                                                                case -9223372036854775808:
                                                                                                    _ = len(v2)
                                                                                                case 933:
                                                                                                    switch v56 {
                                                                                                    case "":
                                                                                                        v54 = v54
                                                                                                    case "foo bar":
                                                                                                        var v87 uint = (v52 / ((v52) | 1))
                                                                                                        _ = v87
                                                                                                    case "\n\t":
                                                                                                        switch ("hello" + (v79 + "\n\t")) {
                                                                                                        case "foo bar":
                                                                                                            fmt.Println(v33, fi53)
                                                                                                        }
                                                                                                    }
                                                                                                    var v88 int16 = 622
                                                                                                    _ = v88
                                                                                                }
                                                                                            } else {
                                                                                                v89 := GenS9{F0: 42, F1: 957, F2: fi57}
                                                                                                if v41 {
                                                                                                    switch -464 {
                                                                                                    case -362:
                                                                                                        v41 = true
                                                                                                    }
                                                                                                    v33 = v33
                                                                                                    fmt.Println(v60, v66)
                                                                                                    v90 := v1[(v4 / ((fi57) | 1))]
                                                                                                    v91 := GenS7{F0: ((v79 + "\n\t") + "a")}
                                                                                                    if (!v41) {
                                                                                                        if (!(true && v41)) {
                                                                                                            v92 := GenS2{F0: (-(v80)), F1: 32767}
                                                                                                            _ = v92
                                                                                                        }
                                                                                                        v93 := GenS11{F0: v41, F1: -475}
                                                                                                        v32 += (v52 & 0)
                                                                                                        _ = v93
                                                                                                    }
                                                                                                    _ = v90
                                                                                                    _ = v91
                                                                                                } else {
                                                                                                    if v41 {
                                                                                                        v94 := GenS8{F0: (v32 ^ v32)}
                                                                                                        if v41 {
                                                                                                            fmt.Println(fi57, v60, v33)
                                                                                                            if v41 {
                                                                                                                v95 := GenS12{F0: (v32 | 255), F1: v32, F2: (-(v54))}
                                                                                                                fmt.Println(v4, v52, v56)
                                                                                                                _ = v95
                                                                                                            }
                                                                                                        }
                                                                                                        fmt.Println(fi53, v79)
                                                                                                        _ = v94
                                                                                                    }
                                                                                                    if (!v41) {
                                                                                                        var v96 int16 = v54
                                                                                                        _ = v96
                                                                                                    } else {
                                                                                                        var v97 string = (v56 + v79)
                                                                                                        v98 := GenS12{F0: (v52 & v52), F1: v52, F2: (v54 >> ((v52) & 63))}
                                                                                                        switch ("foo bar" + v97) {
                                                                                                        case "a":
                                                                                                            v80 = v80
                                                                                                        default:
                                                                                                            switch v4 {
                                                                                                            case -9223372036854775808:
                                                                                                                v99 := v3[(v4 << (((v52 & v32)) & 63))]
                                                                                                                _ = v99
                                                                                                            case 111:
                                                                                                                fmt.Println(v97)
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v97
                                                                                                        _ = v98
                                                                                                    }
                                                                                                    v56 = ""
                                                                                                    var v100 bool = v41
                                                                                                    _ = v100
                                                                                                }
                                                                                                switch 127 {
                                                                                                case 56:
                                                                                                    if len(v31) >= 2 {
                                                                                                        v31 = v31[1:len(v31)]
                                                                                                    }
                                                                                                default:
                                                                                                    var v101 int8 = (v33 % ((v33) | 1))
                                                                                                    _ = v101
                                                                                                }
                                                                                                _ = v89
                                                                                            }
                                                                                            _ = v86
                                                                                        } else {
                                                                                            v102 := GenS13{F0: false, F1: 9223372036854775807, F2: v56}
                                                                                            v54 -= (v54 % ((v54) | 1))
                                                                                            switch ("\n\t" + (v56 + ("\n\t" + "foo bar"))) {
                                                                                            case "a":
                                                                                                if false {
                                                                                                    v60 += (v66 + v60)
                                                                                                    v103 := GenS10{F0: 9223372036854775807}
                                                                                                    switch (fi53 | 0) {
                                                                                                    case 128:
                                                                                                        v104 := GenS14{F0: v52, F1: 127, F2: (v59 % ((v59) | 1))}
                                                                                                        _ = v104
                                                                                                    }
                                                                                                    _ = v103
                                                                                                } else {
                                                                                                    v33 += 42
                                                                                                    var v105 int = fi53
                                                                                                    v60 = ((3.14159 - (3.14159 / v51)) + v66)
                                                                                                    _ = v105
                                                                                                }
                                                                                            case "":
                                                                                                v106 := GenS3{F0: v32, F1: 42}
                                                                                                _ = v106
                                                                                            case "hello":
                                                                                                v51 -= v51
                                                                                            }
                                                                                            _ = v102
                                                                                        }
                                                                                        switch (-(v4)) {
                                                                                        case 1:
                                                                                            if len(v31) > 0 {
                                                                                                v31[0] = 255
                                                                                            }
                                                                                        case -630:
                                                                                            switch int(v59) {
                                                                                            case 128:
                                                                                                v107 := GenS6{F0: v33}
                                                                                                _ = v107
                                                                                            }
                                                                                        case 127:
                                                                                            fmt.Println(v52, v41, v59)
                                                                                        default:
                                                                                            if (uint8(v59) != 255) {
                                                                                                var v108 float32 = (1e10 - (0.0 / float32(v66)))
                                                                                                _ = v108
                                                                                            } else {
                                                                                                fmt.Println(fi57, v37, v56)
                                                                                            }
                                                                                            v109 := v1[(v4 * (fi53 - (fi57 & fi53)))]
                                                                                            v80 += v80
                                                                                            var v110 uint64 = 1
                                                                                            _ = v109
                                                                                            _ = v110
                                                                                        }
                                                                                    } else {
                                                                                        fmt.Println(v66)
                                                                                    }
                                                                                    _ = v80
                                                                                } else {
                                                                                    v4 += fi53
                                                                                    switch "\n\t" {
                                                                                    case "":
                                                                                        _ = len(v31)
                                                                                    default:
                                                                                        v4 = 4294967295
                                                                                    }
                                                                                    _ = len(v2)
                                                                                }
                                                                            }
                                                                        }
                                                                        if (128 <= 42) {
                                                                            v56 = (v79 + "\n\t")
                                                                        } else {
                                                                            if ((v59 == 2) && ((2 != 127) && (false || (!true)))) {
                                                                                if ((v59 >> ((v52) & 63)) <= v59) {
                                                                                    var v111 int8 = (v33 / ((0) | 1))
                                                                                    v32 = (v52 + 18446744073709551615)
                                                                                    _ = v111
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v56)
                                                                                v60 -= (1e10 + v60)
                                                                                var v112 string = v37
                                                                                _ = v112
                                                                            }
                                                                        }
                                                                        _ = v79
                                                                    }
                                                                case "\n\t":
                                                                    v54 += v54
                                                                case "a":
                                                                    v113 := v1[128]
                                                                    _ = v113
                                                                }
                                                                var v114 int16 = (v54 % (((v54 >> (((v52 + (^(v52)))) & 63))) | 1))
                                                                _ = v114
                                                            }
                                                            v52 -= (v32 + (v32 & uint(v59)))
                                                            v37 = (("a" + v56) + "hello")
                                                        }
                                                        if (!true) {
                                                            v56 = (v37 + (v56 + "hello"))
                                                            var v115 uint8 = 219
                                                            _ = v115
                                                        }
                                                        v1[(fi57 % ((0) | 1))] = 2147483647
                                                    } else {
                                                        v56 = v56
                                                    }
                                                    var v116 string = "foo bar"
                                                    v117 := GenS2{F0: -332, F1: 2}
                                                    _ = v116
                                                    _ = v117
                                                } else {
                                                    v66 -= (1e10 * v51)
                                                    var v118 bool = v41
                                                    fmt.Println(v118)
                                                    _ = v118
                                                }
                                                if len(v31) > 0 {
                                                    v31[0] = fi57
                                                }
                                            }
                                            _ = v64
                                            _ = v65
                                            _ = v66
                                        }
                                        _ = v60
                                    }
                                    v52 = 2
                                } else {
                                    var v119 int32 = -1
                                    _ = v119
                                }
                                _ = v59
                            }
                            _ = fi57
                        }
                        _ = v56
                    }
                    _ = v55
                } else {
                    for fi120 := 0; fi120 < 2; fi120++ {
                        v4 = (fi53 | int(v54))
                        _ = fi120
                    }
                    var v121 int = fi53
                    _ = v121
                }
                switch v37 {
                case "a":
                    var v122 int16 = (v54 / (((v54 | (v54 * v54))) | 1))
                    _ = v122
                case "\n\t":
                    var v123 int16 = (-(v54))
                    _ = v123
                }
                _ = fi53
                _ = v54
            }
            var v124 uint16 = 1
            _ = v124
        }
    case -9223372036854775808:
        v51 = (v51 - float64(-2.25))
    case 9223372036854775807:
        v52 = (v52 | v32)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v37
    _ = v38
    _ = v4
    _ = v41
    _ = v51
    _ = v52
}
