package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 uint32; F1 uint32; F2 string }
type GenS7 struct { F0 int }
func genFn0(p0 int, p1 uint8) bool {
    p1 -= (p1 ^ (p1 % ((127) | 1)))
    v4 := []int{p0, p0}
    var v5 int = p0
    var v6 int64 = -9223372036854775808
    for fi7 := 0; fi7 < 2; fi7++ {
        v6 += -246
        var v8 uint = 255
        var v9 int = 32768
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return true
}
func genFn1() uint16 {
    if true {
        switch 42 {
        case -2147483648:
            v10 := []int{32767, 4294967295, 42, 32767}
            _ = v10
        case -78:
            if (!(248 > 145)) {
                switch ("" + "") {
                case "foo bar":
                    if (1 == 152) {
                        var v11 int = 2147483647
                        var v12 uint32 = 164
                        _ = v11
                        _ = v12
                    } else {
                        for fi13 := 0; fi13 < 3; fi13++ {
                            for fi14 := 0; fi14 < 14; fi14++ {
                                switch fi13 {
                                case 0:
                                    switch "a" {
                                    case "a":
                                        var v15 int = -1
                                        _ = v15
                                    case "\n\t":
                                        fmt.Println(fi14, fi14)
                                    case "foo bar":
                                        var v16 uint64 = 4294967295
                                        _ = v16
                                    default:
                                        if (-739 <= 1) {
                                            var v17 uint8 = 128
                                            _ = v17
                                        }
                                    }
                                case 127:
                                    if (!true) {
                                        fmt.Println(fi14)
                                    } else {
                                        switch (fi14 - fi13) {
                                        case -22:
                                            var v18 uint64 = 2
                                            _ = v18
                                        case 0:
                                            fmt.Println(fi14, fi14)
                                        default:
                                            switch (fi14 << ((65535) & 63)) {
                                            case 32767:
                                                var v19 int8 = -128
                                                _ = v19
                                            case 128:
                                                if (false || true) {
                                                    switch fi14 {
                                                    case 32767:
                                                        switch ("a" + ((("" + "\n\t") + "\n\t") + ("hello" + "\n\t"))) {
                                                        case "":
                                                            v20 := []int{fi13, fi14, 183}
                                                            _ = v20
                                                        }
                                                    case 32768:
                                                        fmt.Println(fi14, fi13)
                                                    case -218:
                                                        var v21 uint32 = 2147483647
                                                        _ = v21
                                                    default:
                                                        switch fi14 {
                                                        case 42:
                                                            v22 := []int{352, (fi13 % ((fi13) | 1)), fi13}
                                                            _ = v22
                                                        case 937:
                                                            if false {
                                                                v23 := GenS1{F0: 65535, F1: -398}
                                                                _ = v23
                                                            } else {
                                                                if true {
                                                                    var v24 int64 = int64(fi13)
                                                                    fmt.Println(fi13, fi14)
                                                                    var v25 string = ""
                                                                    _ = v24
                                                                    _ = v25
                                                                }
                                                            }
                                                        case -2147483648:
                                                            var v26 uint8 = 0
                                                            _ = v26
                                                        default:
                                                            if (!false) {
                                                                v27 := GenS1{F0: 659, F1: 765}
                                                                var v28 int32 = -272
                                                                v29 := GenS1{F0: 32767, F1: int16(v28)}
                                                                _ = v27
                                                                _ = v28
                                                                _ = v29
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    v30 := []int{0, fi13}
                                                    _ = v30
                                                }
                                            default:
                                                var v31 string = "\n\t"
                                                _ = v31
                                            }
                                            var v32 int16 = 617
                                            _ = v32
                                        }
                                        if (0 > 42) {
                                            var v33 uint8 = 186
                                            fmt.Println(v33, v33, fi13)
                                            _ = v33
                                        }
                                        var v34 bool = (-1 == (fi13 & (fi13 << ((18446744073709551615) & 63))))
                                        var v35 float64 = ((0.1 + ((0.0 - 0.1) - 0.0)) / float64(3.14159))
                                        _ = v34
                                        _ = v35
                                    }
                                    fmt.Println(fi13)
                                }
                                var v36 uint32 = 4294967295
                                v37 := GenS1{F0: 409, F1: 1000}
                                _ = fi14
                                _ = v36
                                _ = v37
                            }
                            var v38 bool = true
                            if ("" != "hello") {
                                v39 := GenS2{F0: v38}
                                v38 = (!(!(0 != 2147483647)))
                                var v40 uint = 648
                                _ = v39
                                _ = v40
                            } else {
                                fmt.Println(fi13, v38)
                            }
                            _ = fi13
                            _ = v38
                        }
                    }
                case "":
                    v41 := GenS3{F0: 103, F1: -992}
                    _ = v41
                case "\n\t":
                    for fi42 := 0; fi42 < 15; fi42++ {
                        if (((true || true) && true) || false) {
                            fmt.Println(fi42)
                            var v43 int = (fi42 * 32768)
                            _ = v43
                        }
                        v44 := GenS3{F0: 80, F1: (fi42 - fi42)}
                        _ = fi42
                        _ = v44
                    }
                }
                if true {
                    var v45 string = "hello"
                    for fi46 := 0; fi46 < 20; fi46++ {
                        v47 := []int{(^(fi46)), (fi46 + (^(fi46)))}
                        if true {
                            var v48 int8 = -45
                            v45 = v45
                            v45 = v45
                            _ = v48
                        }
                        _ = fi46
                        _ = v47
                    }
                    _ = v45
                }
            } else {
                for fi49 := 0; fi49 < 14; fi49++ {
                    var v50 int16 = -315
                    v51 := GenS4{F0: v50, F1: 61, F2: 332}
                    switch "hello" {
                    case "a":
                        v50 -= 213
                    case "hello":
                        v52 := []int{(fi49 << ((255) & 63))}
                        _ = v52
                    case "\n\t":
                        v50 = -347
                    }
                    _ = fi49
                    _ = v50
                    _ = v51
                }
                var v53 float32 = (-2.25 - (1.5 / 1.0))
                _ = v53
            }
        case 127:
            var v54 uint32 = 4294967295
            _ = v54
        }
        v55 := []int{9223372036854775807}
        switch 0 {
        case 127:
            switch -2147483648 {
            case 42:
                _ = len(v55)
            case -2147483648:
                for fi56 := 0; fi56 < 11; fi56++ {
                    _ = len(v55)
                    _ = fi56
                }
            case 2147483647:
                var v57 uint32 = 489
                _ = v57
            }
        case 29:
            for fi58 := 0; fi58 < 14; fi58++ {
                if true {
                    if len(v55) >= 2 {
                        v55 = v55[1:len(v55)]
                    }
                    var v59 float32 = float32(0.1)
                    fmt.Println(fi58, fi58)
                    _ = v59
                }
                var v60 uint32 = 42
                _ = fi58
                _ = v60
            }
        case -241:
            v61 := GenS5{F0: 0, F1: (1 <= 2)}
            _ = v61
        default:
            var v62 uint32 = 0
            _ = v62
        }
        _ = v55
    }
    var v63 uint32 = 65535
    switch "\n\t" {
    case "\n\t":
        var v64 uint32 = (v63 << ((0) & 63))
        _ = v64
    default:
        if true {
            for fi65 := 0; fi65 < 18; fi65++ {
                switch "\n\t" {
                case "hello":
                    v66 := GenS2{F0: (65535 < 0)}
                    _ = v66
                default:
                    v63 = 2147483647
                }
                for fi67 := 0; fi67 < 18; fi67++ {
                    v63 = v63
                    fmt.Println(fi67, fi67, v63)
                    _ = fi67
                }
                v68 := []int{}
                _ = fi65
                _ = v68
            }
            for fi69 := 0; fi69 < 5; fi69++ {
                var v70 uint = 65535
                var v71 float32 = (3.14159 * 0.0)
                _ = fi69
                _ = v70
                _ = v71
            }
        }
    }
    _ = v63
    return uint16(v63)
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{552, 796}
    v3 := []string{("foo bar" + "foo bar"), "\n\t", "hello"}
    var v72 int8 = 127
    v72 += (v72 - (v72 | v72))
    for fi73 := 0; fi73 < 12; fi73++ {
        v72 -= 127
        if false {
            var v74 uint64 = 42
            _ = v74
        }
        _ = fi73
    }
    var v75 int32 = -657
    if true {
        var v76 string = ""
        var v77 uint8 = 23
        switch v76 {
        case "foo bar":
            if (42 > 2) {
                if false {
                    var v78 string = (v76 + (v76 + "hello"))
                    v75 = v75
                    _ = v78
                }
                if (false || (1 != 32767)) {
                    if (v76 <= v76) {
                        if false {
                            for fi79 := 0; fi79 < 13; fi79++ {
                                v76 = "foo bar"
                                _ = fi79
                            }
                        }
                    }
                }
            }
        }
        _ = v76
        _ = v77
    } else {
        var v80 uint64 = 963
        v75 += (v75 - v75)
        _ = v80
    }
    v75 += v75
    if (-865 == 42) {
        if (2147483647 == -737) {
            fmt.Println(v75, v72)
        }
        v72 += v72
        switch "" {
        case "\n\t":
            switch (("foo bar" + "\n\t") + ("\n\t" + "")) {
            case "foo bar":
                v75 -= -1
            default:
                var v81 uint = 171
                _ = v81
            }
        case "hello":
            v75 += 0
        }
    } else {
        v82 := GenS5{F0: 730, F1: ((!false) && false)}
        var v83 float64 = 3.14159
        v84 := GenS4{F0: 471, F1: 42, F2: 0}
        _ = v82
        _ = v83
        _ = v84
    }
    fmt.Println(v75, v72)
    v85 := GenS6{F0: 42, F1: 1, F2: ("" + ("foo bar" + "\n\t"))}
    var v86 uint64 = 42
    v75 -= (v75 >> ((2) & 63))
    for fi87 := 0; fi87 < 19; fi87++ {
        if true {
            for fi88 := 0; fi88 < 15; fi88++ {
                var v89 int16 = 0
                var v90 int8 = v72
                if (false && true) {
                    if ((-(v89)) != (v89 >> ((65535) & 63))) {
                        v91 := GenS2{F0: false}
                        _ = v91
                    }
                    fmt.Println(fi87, fi88)
                } else {
                    fmt.Println(fi87)
                }
                _ = fi88
                _ = v89
                _ = v90
            }
            v75 = int32(v86)
            if ("" > "") {
                if true {
                    switch ("a" + "a") {
                    case "foo bar":
                        fmt.Println(v75, v75)
                    case "\n\t":
                        var v92 float32 = -2.25
                        _ = v92
                    case "":
                        if (false && false) {
                            v75 += -2147483648
                            if true {
                                fmt.Println(fi87, v72)
                            } else {
                                if (0 <= 114) {
                                    v75 -= (v75 >> ((255) & 63))
                                }
                                var v93 int8 = v72
                                _ = v93
                            }
                            if (1 <= 4294967295) {
                                _ = len(v2)
                                v1[65535] = (fi87 % ((fi87) | 1))
                                var v94 int = (fi87 & int(v72))
                                _ = v94
                            } else {
                                v95 := GenS3{F0: (v72 % ((0) | 1)), F1: fi87}
                                v86 -= v86
                                _ = v95
                            }
                        } else {
                            fmt.Println(fi87, v86, v86)
                            v72 -= (v72 >> ((667) & 63))
                            for fi96 := 0; fi96 < 20; fi96++ {
                                v72 -= v72
                                var v97 float32 = 3.14159
                                _ = len(v1)
                                _ = fi96
                                _ = v97
                            }
                        }
                    default:
                        v98 := GenS7{F0: fi87}
                        _ = v98
                    }
                    var v99 uint = 2
                    v72 = v72
                    _ = v99
                }
            } else {
                for fi100 := 0; fi100 < 1; fi100++ {
                    if false {
                        var v101 int16 = 1
                        v1[65535] = (fi87 / ((1) | 1))
                        _ = v101
                    }
                    switch 2147483647 {
                    case 32767:
                        if false {
                            var v102 uint32 = 2147483647
                            if ("" != ("a" + "\n\t")) {
                                v103 := GenS2{F0: (((0.0 + 1.5) != ((0.0 * float64((float32(1e10) + 1.5))) + 0.1)) || false)}
                                var v104 float32 = (1e10 + 0.0)
                                var v105 int8 = (v72 >> ((18446744073709551615) & 63))
                                if (v75 == (-(v75))) {
                                    v106 := GenS6{F0: v102, F1: 4294967295, F2: ("\n\t" + ((("foo bar" + "foo bar") + "hello") + "\n\t"))}
                                    _ = v106
                                }
                                var v107 bool = false
                                _ = v103
                                _ = v104
                                _ = v105
                                _ = v107
                            }
                            v3 = append(v3, "a")
                            v108 := GenS4{F0: -32768, F1: 65535, F2: (v86 / ((v86) | 1))}
                            v109 := GenS5{F0: (v86 + 0), F1: (!(-623 < 715))}
                            fmt.Println(v72, fi100)
                            _ = v102
                            _ = v108
                            _ = v109
                        } else {
                            v75 -= -2147483648
                            v110 := GenS4{F0: 1, F1: 1, F2: v86}
                            delete(v1, 0)
                            _ = v110
                        }
                    case 32768:
                        if false {
                            if ((!((127 <= 25) || (((v72 >> ((564) & 63)) > (-(v72))) && ((^(v86)) >= (v86 | (v86 ^ v86)))))) || (!true)) {
                                var v111 int8 = v72
                                _ = v111
                            }
                        }
                    default:
                        _ = len(v1)
                    }
                    _ = fi100
                }
                v86 -= v86
            }
        }
        var v112 bool = (0 > -1)
        _ = fi87
        _ = v112
    }
    if (false || ("" > "")) {
        v75 -= 0
        v86 = (v86 % ((v86) | 1))
    } else {
        v113 := GenS6{F0: 0, F1: 2147483647, F2: ("\n\t" + "")}
        _ = v113
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v72:", v72)
    fmt.Println("v75:", v75)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v72
    _ = v75
    _ = v85
    _ = v86
}
