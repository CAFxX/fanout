package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int32; F1 int64; F2 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32; F1 uint64; F2 int8 }
type GenS8 struct { F0 uint32; F1 uint }
type GenS9 struct { F0 int16; F1 string; F2 int64 }
type GenS10 struct { F0 int; F1 int64 }
type GenS11 struct { F0 int }
func genFn0() uint {
    var v6 float32 = 0.0
    var v7 uint16 = 255
    fmt.Println(v7, v6)
    _ = v6
    _ = v7
    return 624
}
func genFn1() int64 {
    var v8 bool = (!(255 == 926))
    v8 = ((-1 != 1000) || true)
    v9 := []int{-1, 0, -306}
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 16; fi11++ {
            if true {
                v12 := GenS1{F0: -362, F1: 0, F2: 2147483647}
                var v13 float32 = (0.1 + 1.5)
                _ = v12
                _ = v13
            }
            v9 = append(v9, (fi11 & (-(fi10))))
            _ = fi11
        }
        for fi14 := 0; fi14 < 1; fi14++ {
            v15 := GenS2{F0: 0, F1: 42, F2: 42}
            _ = fi14
            _ = v15
        }
        v9 = append(v9, fi10)
        _ = fi10
    }
    _ = v8
    _ = v9
    return 714
}

func main() {
    var v1 uint64 = 487
    var v2 string = ""
    var v3 string = (v2 + (v2 + ("\n\t" + v2)))
    var v4 int16 = -767
    var v5 bool = (!(!(!false)))
    switch 118 {
    case 42:
        if (!(18446744073709551615 < 2)) {
            var v16 uint64 = 744
            v17 := GenS3{F0: 0}
            switch int(v4) {
            case -653:
                for fi18 := 0; fi18 < 20; fi18++ {
                    var v19 uint = 1
                    switch (v3 + ((v3 + v3) + "")) {
                    case "hello":
                        var v20 uint8 = 0
                        _ = v20
                    default:
                        if v5 {
                            v21 := GenS4{F0: 966}
                            v16 = (^(v1))
                            var v22 string = ((((v3 + "\n\t") + "foo bar") + "hello") + ("foo bar" + v2))
                            _ = v21
                            _ = v22
                        }
                    }
                    _ = fi18
                    _ = v19
                }
            case 854:
                if (v5 || v5) {
                    fmt.Println(v16)
                    v23 := GenS1{F0: 516, F1: 75, F2: 4294967295}
                    if false {
                        if (-86 < int8(v4)) {
                            var v24 uint8 = 199
                            var v25 float32 = 1.5
                            var v26 float32 = v25
                            _ = v24
                            _ = v25
                            _ = v26
                        } else {
                            var v27 string = "foo bar"
                            _ = v27
                        }
                        v3 = (v3 + "a")
                    }
                    _ = v23
                }
            case 1:
                for fi28 := 0; fi28 < 20; fi28++ {
                    for fi29 := 0; fi29 < 20; fi29++ {
                        if v5 {
                            v30 := GenS5{F0: v4, F1: (fi28 & -2147483648)}
                            v31 := []int{1, 32768, fi29, fi29}
                            fmt.Println(fi28, fi28, v16)
                            _ = v30
                            _ = v31
                        } else {
                            if (true || (((v4 << ((0) & 63)) == 935) && (v5 && (128 <= 0)))) {
                                fmt.Println(fi29)
                                switch v3 {
                                case "foo bar":
                                    v32 := []int{(-(fi29)), (fi28 % (((^(fi29))) | 1)), (^(fi29))}
                                    _ = v32
                                default:
                                    v2 = (((("a" + "") + "\n\t") + v2) + v3)
                                }
                                var v33 int16 = -32768
                                _ = v33
                            }
                        }
                        var v34 int64 = 0
                        _ = fi29
                        _ = v34
                    }
                    _ = fi28
                }
            default:
                if (400 == 255) {
                    fmt.Println(v5, v3, v1)
                    if (69 > 683) {
                        fmt.Println(v16, v4, v5)
                        if (1 > -1) {
                            var v35 uint16 = 0
                            var v36 int64 = genFn1()
                            _ = v35
                            _ = v36
                        } else {
                            var v37 int8 = -114
                            if (9223372036854775807 <= 32768) {
                                for fi38 := 0; fi38 < 10; fi38++ {
                                    v3 = "foo bar"
                                    _ = fi38
                                }
                                switch v2 {
                                case "hello":
                                    v39 := GenS1{F0: -575, F1: 120, F2: 42}
                                    _ = v39
                                case "a":
                                    v40 := []int{167, 32768}
                                    _ = v40
                                case "":
                                    if (53 == 117) {
                                        for fi41 := 0; fi41 < 13; fi41++ {
                                            if true {
                                                var v42 uint32 = 4294967295
                                                for fi43 := 0; fi43 < 17; fi43++ {
                                                    switch "" {
                                                    case "foo bar":
                                                        fmt.Println(fi41, fi43, v16)
                                                    case "a":
                                                        var v44 uint = 4294967295
                                                        _ = v44
                                                    }
                                                    _ = fi43
                                                }
                                                _ = v42
                                            } else {
                                                var v45 int32 = -1
                                                v46 := []int{(fi41 + fi41), (-(fi41)), (-(fi41)), (fi41 / (((fi41 - 128)) | 1))}
                                                for fi47 := 0; fi47 < 6; fi47++ {
                                                    switch -782 {
                                                    case 128:
                                                        var v48 int8 = v37
                                                        _ = v48
                                                    case 898:
                                                        if len(v46) >= 2 {
                                                            v46 = v46[1:len(v46)]
                                                        }
                                                    case 197:
                                                        var v49 int16 = v4
                                                        _ = v49
                                                    }
                                                    _ = fi47
                                                }
                                                _ = v45
                                                _ = v46
                                            }
                                            _ = fi41
                                        }
                                        var v50 uint64 = 4294967295
                                        switch v2 {
                                        case "a":
                                            if ((3.14159 * 1e10) >= (((0.1 / 0.1) - (1e10 / 0.1)) - 1.5)) {
                                                if true {
                                                    v16 += (^(v50))
                                                } else {
                                                    v51 := []int{545, -9223372036854775808, 128}
                                                    v51 = append(v51, 32768)
                                                    if v5 {
                                                        v3 = "\n\t"
                                                        v5 = v5
                                                        var v52 uint16 = 293
                                                        _ = v52
                                                    }
                                                    _ = v51
                                                }
                                            } else {
                                                v2 = "\n\t"
                                            }
                                        }
                                        _ = v50
                                    }
                                }
                                var v53 uint8 = 128
                                _ = v53
                            } else {
                                var v54 uint16 = 605
                                var v55 int8 = (v37 | (v37 | (v37 - v37)))
                                _ = v54
                                _ = v55
                            }
                            _ = v37
                        }
                    } else {
                        var v56 string = v3
                        v5 = (0.1 != 3.14159)
                        _ = v56
                    }
                    v16 = v16
                }
            }
            _ = v16
            _ = v17
        }
    case 2147483647:
        v57 := []int{0, 32767}
        _ = v57
    default:
        fmt.Println(v2)
    }
    var v58 string = "a"
    if (255 >= 42) {
        var v59 uint32 = 870
        fmt.Println(v58)
        for fi60 := 0; fi60 < 3; fi60++ {
            fmt.Println(v58, v2)
            _ = fi60
        }
        _ = v59
    }
    switch v2 {
    case "":
        if false {
            var v61 uint32 = 65535
            if ((v1 ^ (^(v1))) <= 2) {
                switch ((v58 + ((v58 + "hello") + "foo bar")) + ("" + v2)) {
                case "foo bar":
                    v62 := []int{128}
                    _ = v62
                default:
                    v3 = (v58 + v58)
                }
                if (831 == v1) {
                    v4 -= -32768
                    var v63 uint32 = v61
                    if v5 {
                        var v64 string = ""
                        _ = v64
                    }
                    _ = v63
                } else {
                    var v65 uint8 = 107
                    fmt.Println(v4, v2)
                    fmt.Println(v58, v1, v3)
                    _ = v65
                }
            } else {
                fmt.Println(v5, v61, v2)
                fmt.Println(v2, v2)
                switch v2 {
                case "\n\t":
                    v3 = (v2 + (v58 + (v3 + v2)))
                }
            }
            if (-13 > -52) {
                var v66 uint32 = 922
                v67 := []int{int(v4), 128, 127}
                _ = v66
                _ = v67
            } else {
                switch 117 {
                case 9223372036854775807:
                    var v68 uint16 = 65535
                    _ = v68
                default:
                    v61 = (v61 / (((v61 % (((v61 + (v61 % ((v61) | 1)))) | 1))) | 1))
                }
                for fi69 := 0; fi69 < 8; fi69++ {
                    for fi70 := 0; fi70 < 8; fi70++ {
                        v71 := []int{(fi70 + (fi69 ^ fi69)), fi69, -614, (fi70 | fi69)}
                        switch fi69 {
                        case 920:
                            v72 := GenS6{F0: 42}
                            v58 = v3
                            if true {
                                v73 := GenS2{F0: -824, F1: 971, F2: (v4 % (((v4 % ((42) | 1))) | 1))}
                                _ = v73
                            }
                            var v74 uint16 = 255
                            if v5 {
                                var v75 float64 = 0.0
                                var v76 int64 = genFn1()
                                v71 = append(v71, fi69)
                                _ = v75
                                _ = v76
                            }
                            if ((false && true) || (-250 <= -1)) {
                                if ((fi69 / ((9223372036854775807) | 1)) < fi70) {
                                    if (v5 || v5) {
                                        v74 -= 630
                                        v77 := GenS7{F0: -1, F1: (v1 | (v1 + (v1 << ((42) & 63)))), F2: -128}
                                        switch fi69 {
                                        case 32:
                                            switch (fi69 << ((uint(v74)) & 63)) {
                                            case 4294967295:
                                                switch 32767 {
                                                case 9223372036854775807:
                                                    var v78 string = v58
                                                    _ = v78
                                                case -436:
                                                    var v79 int64 = int64(v4)
                                                    _ = v79
                                                case 32768:
                                                    if v5 {
                                                        if (299 != v61) {
                                                            if false {
                                                                var v80 int64 = -1
                                                                if v5 {
                                                                    var v81 int = (fi70 >> ((2) & 63))
                                                                    var v82 uint16 = 255
                                                                    v71 = append(v71, 32768)
                                                                    _ = v81
                                                                    _ = v82
                                                                } else {
                                                                    fmt.Println(v5)
                                                                }
                                                                if v5 {
                                                                    fmt.Println(v3, v4)
                                                                    if v5 {
                                                                        v5 = (uint8(v1) > uint8(v1))
                                                                    } else {
                                                                        var v83 bool = true
                                                                        if (42 > (v4 * (v4 - 42))) {
                                                                            var v84 int = fi70
                                                                            fmt.Println(v80)
                                                                            _ = v84
                                                                        } else {
                                                                            if false {
                                                                                v4 = (v4 >> ((0) & 63))
                                                                            }
                                                                        }
                                                                        _ = v83
                                                                    }
                                                                    v2 = "a"
                                                                }
                                                                _ = v80
                                                            } else {
                                                                v85 := GenS6{F0: 1}
                                                                if true {
                                                                    var v86 uint = 255
                                                                    var v87 int64 = 9223372036854775807
                                                                    var v88 uint32 = 0
                                                                    _ = v86
                                                                    _ = v87
                                                                    _ = v88
                                                                } else {
                                                                    var v89 float64 = 0.0
                                                                    var v90 bool = ((v74 + (^(v74))) > (v74 & (^(v74))))
                                                                    _ = v89
                                                                    _ = v90
                                                                }
                                                                _ = v85
                                                            }
                                                            v3 = "a"
                                                            var v91 int64 = 0
                                                            _ = v91
                                                        } else {
                                                            v71 = append(v71, (fi70 + 86))
                                                            v92 := GenS8{F0: 335, F1: 42}
                                                            _ = v92
                                                        }
                                                        v2 = v58
                                                    } else {
                                                        var v93 uint32 = (v61 * v61)
                                                        if len(v71) > 0 {
                                                            v71[0] = 959
                                                        }
                                                        _ = v93
                                                    }
                                                default:
                                                    v1 = v1
                                                }
                                            }
                                        case -1:
                                            switch (v2 + "\n\t") {
                                            case "hello":
                                                v61 -= (v61 | (^(v61)))
                                            case "foo bar":
                                                v4 -= v4
                                            case "":
                                                v3 = v58
                                            }
                                        }
                                        _ = v77
                                    }
                                    fmt.Println(fi69)
                                    switch "" {
                                    case "hello":
                                        var v94 bool = true
                                        _ = v94
                                    case "\n\t":
                                        if (!v5) {
                                            v61 -= 65535
                                        } else {
                                            var v95 string = v2
                                            _ = v95
                                        }
                                    case "foo bar":
                                        if (1e10 < 1.5) {
                                            var v96 uint = 4294967295
                                            v97 := GenS1{F0: 0, F1: 127, F2: v61}
                                            fmt.Println(v74)
                                            _ = v96
                                            _ = v97
                                        }
                                    default:
                                        v4 = (v4 << ((255) & 63))
                                    }
                                    _ = len(v71)
                                    v3 = "a"
                                } else {
                                    if (-2.25 == ((((0.0 / float32((1e10 - (1e10 / 3.14159)))) / 0.1) + 3.14159) / 0.1)) {
                                        fmt.Println(v74, v61, v74)
                                        v58 = v58
                                    }
                                    v98 := GenS9{F0: v4, F1: ((v2 + v3) + "hello"), F2: genFn1()}
                                    _ = v98
                                }
                            }
                            _ = v72
                            _ = v74
                        case 362:
                            v4 += (v4 & 21)
                        default:
                            fmt.Println(v4)
                        }
                        switch v58 {
                        case "\n\t":
                            v58 = v2
                        default:
                            if v5 {
                                switch "" {
                                case "":
                                    var v99 bool = v5
                                    _ = v99
                                case "hello":
                                    fmt.Println(v1, v3)
                                }
                                v100 := GenS5{F0: v4, F1: 32767}
                                _ = v100
                            }
                        }
                        _ = fi70
                        _ = v71
                    }
                    if (((true && (!(v5 && v5))) && true) || ((!(v5 && false)) || true)) {
                        if (!(v1 > 4294967295)) {
                            var v101 uint16 = 28
                            _ = v101
                        }
                        v58 = ("a" + (("" + (v2 + v3)) + (("\n\t" + (v58 + ((v3 + v2) + v2))) + ("hello" + v58))))
                    }
                    _ = fi69
                }
            }
            _ = v61
        }
    case "a":
        if v5 {
            v4 = 1000
        } else {
            v102 := []int{int(v1), 9223372036854775807}
            _ = v102
        }
    default:
        for fi103 := 0; fi103 < 3; fi103++ {
            v2 = (("a" + v2) + v2)
            var v104 bool = (1e10 < (-2.25 * ((0.1 + (float64(1e10) - 3.14159)) + float64((float32(float64(-2.25)) + (1e10 / -2.25))))))
            _ = fi103
            _ = v104
        }
    }
    v105 := []int{}
    for fi106 := 0; fi106 < 9; fi106++ {
        fmt.Println(v4)
        if (true || true) {
            v107 := GenS10{F0: (^(fi106)), F1: 0}
            v58 = v2
            v58 = v2
            _ = v107
        } else {
            v108 := GenS7{F0: 2147483647, F1: v1, F2: 42}
            _ = v108
        }
        _ = fi106
    }
    if v5 {
        var v109 string = "hello"
        _ = v109
    }
    if (!v5) {
        var v110 float32 = 0.1
        if (121 >= 0) {
            v111 := GenS5{F0: v4, F1: 1}
            _ = v111
        }
        _ = v110
    }
    var v112 string = v58
    switch "foo bar" {
    case "foo bar":
        for fi113 := 0; fi113 < 8; fi113++ {
            _ = len(v105)
            if (!v5) {
                fmt.Println(v112, v2)
                fmt.Println(v1, v5)
            }
            _ = fi113
        }
    case "hello":
        for fi114 := 0; fi114 < 14; fi114++ {
            switch (fi114 << ((1) & 63)) {
            case 80:
                v115 := GenS4{F0: 112}
                _ = v115
            default:
                if len(v105) >= 2 {
                    v105 = v105[1:len(v105)]
                }
            }
            v4 += v4
            v3 = v2
            _ = fi114
        }
    case "":
        v116 := GenS11{F0: -149}
        _ = v116
    }
    if v5 {
        for fi117 := 0; fi117 < 4; fi117++ {
            var v118 int32 = 473
            v119 := GenS10{F0: -835, F1: genFn1()}
            _ = fi117
            _ = v118
            _ = v119
        }
        for fi120 := 0; fi120 < 19; fi120++ {
            v2 = v58
            _ = fi120
        }
    }
    var v121 int32 = -1
    for fi122 := 0; fi122 < 18; fi122++ {
        if false {
            for fi123 := 0; fi123 < 3; fi123++ {
                if len(v105) > 0 {
                    v105[0] = (-(fi123))
                }
                var v124 int16 = (v4 | v4)
                v105 = append(v105, (-(fi122)))
                _ = fi123
                _ = v124
            }
            v1 -= (v1 << ((uint(v1)) & 63))
            switch (fi122 + 1) {
            case 449:
                if false {
                    v105 = append(v105, (fi122 + 2147483647))
                    v5 = (!v5)
                    if len(v105) > 0 {
                        v105[0] = (fi122 >> ((905) & 63))
                    }
                }
            default:
                var v125 int32 = (v121 >> ((65535) & 63))
                _ = v125
            }
        }
        if len(v105) > 0 {
            v105[0] = 1
        }
        v4 -= (-(v4))
        _ = fi122
    }
    fmt.Println(v4)
    var v126 uint64 = v1
    if v5 {
        v127 := GenS11{F0: -2147483648}
        var v128 string = v3
        _ = v127
        _ = v128
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v105:", len(v105), v105)
    fmt.Println("v112:", v112)
    fmt.Println("v121:", v121)
    fmt.Println("v126:", v126)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v105
    _ = v112
    _ = v121
    _ = v126
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
}
