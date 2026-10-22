package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int16; F2 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint; F1 int32; F2 uint }
func genFn0() uint32 {
    v7 := GenS1{F0: 4294967295, F1: 2}
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := []int{(fi8 ^ (-(fi8)))}
        _ = fi8
        _ = v9
    }
    if true {
        v10 := GenS1{F0: 529, F1: 140}
        if (!(float32(((0.0 / 1.5) / 1e10)) <= float32(0.1))) {
            var v11 int = -1
            _ = v11
        }
        var v12 int8 = -24
        _ = v10
        _ = v12
    } else {
        var v13 uint16 = 392
        var v14 int32 = 2147483647
        var v15 uint32 = 2147483648
        _ = v13
        _ = v14
        _ = v15
    }
    _ = v7
    return 1
}
func genFn1(p0 int, p1 uint8, p2 uint8) int64 {
    v16 := GenS1{F0: genFn0(), F1: p2}
    for fi17 := 0; fi17 < 4; fi17++ {
        fmt.Println(p2, p2)
        for fi18 := 0; fi18 < 11; fi18++ {
            fmt.Println(p1, p1, fi18)
            fmt.Println(p1, fi17, p0)
            _ = fi18
        }
        var v19 uint8 = p1
        _ = fi17
        _ = v19
    }
    if (!(!true)) {
        for fi20 := 0; fi20 < 18; fi20++ {
            v21 := []int{fi20}
            _ = fi20
            _ = v21
        }
        fmt.Println(p2)
    }
    v22 := []int{(p0 * p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v22
    return 0
}

func main() {
    var v1 uint16 = 32767
    var v2 uint32 = 65535
    var v3 uint = 65535
    var v4 uint8 = 1
    v5 := make(map[string]int)
    v6 := make(map[int]bool)
    if true {
        if ((0.1 < -2.25) && true) {
            var v23 uint32 = (v2 << ((v3) & 63))
            _ = v23
        }
        for fi24 := 0; fi24 < 6; fi24++ {
            fmt.Println(v3)
            var v25 int = (fi24 ^ (fi24 << ((2) & 63)))
            _ = fi24
            _ = v25
        }
    }
    switch -2147483648 {
    case 1:
        v1 += v1
    case -1:
        for fi26 := 0; fi26 < 11; fi26++ {
            if true {
                v4 = v4
                v2 -= v2
                switch ("hello" + "a") {
                case "\n\t":
                    for fi27 := 0; fi27 < 20; fi27++ {
                        v28 := []int{(fi26 * (fi26 & (fi27 % ((fi27) | 1)))), fi26}
                        v2 -= 4294967295
                        _ = fi27
                        _ = v28
                    }
                case "foo bar":
                    _ = len(v6)
                case "hello":
                    fmt.Println(v4, v2, v3)
                }
            } else {
                for fi29 := 0; fi29 < 18; fi29++ {
                    if (!(17 >= 0)) {
                        fmt.Println(v4, v3, v2)
                    }
                    v30 := GenS2{F0: -32768}
                    _ = fi29
                    _ = v30
                }
            }
            _ = fi26
        }
    case -365:
        if true {
            fmt.Println(v2, v3)
            v3 = 762
        } else {
            v4 -= v4
            if false {
                v1 += v1
                for fi31 := 0; fi31 < 18; fi31++ {
                    var v32 float32 = 0.0
                    if (v3 < (^(v3))) {
                        if (true && true) {
                            v2 = 881
                        }
                        switch (fi31 % (((fi31 / ((fi31) | 1))) | 1)) {
                        case 4294967295:
                            switch "\n\t" {
                            case "foo bar":
                                var v33 int = 919
                                _ = v33
                            case "":
                                if (!true) {
                                    fmt.Println(v4)
                                } else {
                                    if (!true) {
                                        fmt.Println(v4, v3, v1)
                                        for fi34 := 0; fi34 < 17; fi34++ {
                                            v3 -= (^(v3))
                                            _ = fi34
                                        }
                                    }
                                    if (uint8(v2) <= (v4 / ((255) | 1))) {
                                        for fi35 := 0; fi35 < 12; fi35++ {
                                            switch (fi35 | fi35) {
                                            case 568:
                                                if false {
                                                    v36 := GenS3{F0: "hello", F1: 42, F2: v3}
                                                    if (9223372036854775807 != 42) {
                                                        v32 = 0.1
                                                        switch fi35 {
                                                        case 82:
                                                            v37 := []int{}
                                                            _ = v37
                                                        case 0:
                                                            v38 := GenS3{F0: "hello", F1: 887, F2: v3}
                                                            _ = v38
                                                        }
                                                        var v39 int32 = -1
                                                        _ = v39
                                                    } else {
                                                        var v40 float64 = 3.14159
                                                        v41 := GenS4{F0: (v4 >> ((v3) & 63))}
                                                        v42 := GenS1{F0: (v2 + v2), F1: 128}
                                                        delete(v5, ("" + ""))
                                                        v43 := []int{(-(fi31))}
                                                        _ = v40
                                                        _ = v41
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                    var v44 int16 = 1
                                                    var v45 int16 = 42
                                                    var v46 uint32 = 2
                                                    fmt.Println(v3, v44)
                                                    v47 := GenS2{F0: v44}
                                                    _ = len(v6)
                                                    _ = v36
                                                    _ = v44
                                                    _ = v45
                                                    _ = v46
                                                    _ = v47
                                                }
                                            case 32768:
                                                switch "a" {
                                                case "":
                                                    v32 = 0.1
                                                    switch 0 {
                                                    case -1:
                                                        if (1 > 32767) {
                                                            v48 := []int{}
                                                            _ = v48
                                                        } else {
                                                            fmt.Println(fi35, v3, fi31)
                                                        }
                                                    case 0:
                                                        if true {
                                                            v49 := []int{(fi35 % ((fi35) | 1))}
                                                            _ = v49
                                                        }
                                                    }
                                                    v50 := GenS4{F0: (v4 % ((v4) | 1))}
                                                    if (true && (uint(v1) < (^(v3)))) {
                                                        v51 := []int{fi35, -1, (fi31 ^ (fi31 ^ (fi31 + fi35)))}
                                                        _ = v51
                                                    } else {
                                                        switch (^(fi31)) {
                                                        case 2147483647:
                                                            v52 := []int{(-(fi31))}
                                                            _ = v52
                                                        default:
                                                            v4 = v4
                                                        }
                                                        v53 := GenS2{F0: 1}
                                                        fmt.Println(fi31)
                                                        _ = v53
                                                    }
                                                    v54 := GenS3{F0: ("\n\t" + ("" + "a")), F1: 1000, F2: v3}
                                                    _ = v50
                                                    _ = v54
                                                case "a":
                                                    if true {
                                                        var v55 int = (-(fi35))
                                                        _ = v55
                                                    } else {
                                                        if true {
                                                            fmt.Println(v3, v1)
                                                            v56 := []int{fi31, 0, (fi31 | (fi31 ^ (fi35 & (fi31 + 128))))}
                                                            fmt.Println(v4, fi31)
                                                            _ = v56
                                                        } else {
                                                            v1 += 255
                                                            var v57 bool = (!true)
                                                            fmt.Println(fi31, fi31)
                                                            _ = v57
                                                        }
                                                    }
                                                }
                                            }
                                            var v58 uint8 = (v4 / ((uint8(v3)) | 1))
                                            _ = fi35
                                            _ = v58
                                        }
                                        var v59 int8 = 0
                                        if false {
                                            _ = len(v5)
                                            v60 := GenS4{F0: (v4 >> ((123) & 63))}
                                            var v61 int8 = v59
                                            _ = v60
                                            _ = v61
                                        } else {
                                            v62 := GenS5{F0: -363, F1: (v59 | -128)}
                                            if false {
                                                for fi63 := 0; fi63 < 20; fi63++ {
                                                    fmt.Println(v32, v2, v3)
                                                    v64 := v6[(fi31 - (fi63 + 42))]
                                                    _ = fi63
                                                    _ = v64
                                                }
                                                switch "hello" {
                                                case "":
                                                    var v65 float64 = -2.25
                                                    _ = v65
                                                case "foo bar":
                                                    for fi66 := 0; fi66 < 8; fi66++ {
                                                        v4 = v4
                                                        _ = fi66
                                                    }
                                                case "a":
                                                    switch -206 {
                                                    case 2147483647:
                                                        fmt.Println(v1)
                                                    }
                                                }
                                            }
                                            var v67 uint8 = v4
                                            _ = v62
                                            _ = v67
                                        }
                                        _ = v59
                                    } else {
                                        v68 := []int{}
                                        _ = v68
                                    }
                                }
                            case "\n\t":
                                var v69 int64 = 42
                                _ = v69
                            }
                        case 42:
                            var v70 uint32 = (v2 * genFn0())
                            _ = v70
                        case 2147483647:
                            if false {
                                v1 += v1
                                var v71 string = "\n\t"
                                _ = v71
                            } else {
                                v72 := GenS3{F0: "hello", F1: 32767, F2: v3}
                                _ = v72
                            }
                        }
                    }
                    var v73 uint32 = v2
                    _ = fi31
                    _ = v32
                    _ = v73
                }
                v4 = (v4 / ((62) | 1))
            }
        }
    default:
        switch "\n\t" {
        case "\n\t":
            v74 := GenS6{F0: -9223372036854775808}
            _ = v74
        case "hello":
            for fi75 := 0; fi75 < 6; fi75++ {
                for fi76 := 0; fi76 < 12; fi76++ {
                    v3 -= (v3 + v3)
                    _ = fi76
                }
                switch ((("a" + "") + "hello") + (((((("" + "\n\t") + ("" + "\n\t")) + ("" + "\n\t")) + "foo bar") + ("\n\t" + "\n\t")) + "a")) {
                case "":
                    v3 = uint(v1)
                case "foo bar":
                    var v77 uint16 = v1
                    _ = v77
                case "\n\t":
                    v2 = v2
                default:
                    var v78 string = "foo bar"
                    _ = v78
                }
                _ = fi75
            }
        case "foo bar":
            if false {
                fmt.Println(v3)
                switch "a" {
                case "hello":
                    var v79 float64 = (((3.14159 * (0.1 / (1.5 - (1e10 * (1e10 * (3.14159 - 0.1)))))) + float64(1e10)) / (-2.25 - 1e10))
                    _ = v79
                case "a":
                    v80 := []int{}
                    _ = v80
                case "":
                    for fi81 := 0; fi81 < 20; fi81++ {
                        for fi82 := 0; fi82 < 3; fi82++ {
                            v83 := []int{fi81}
                            v1 = (v1 ^ (v1 + (v1 * (^(v1)))))
                            if false {
                                _ = len(v6)
                                var v84 int = fi81
                                var v85 uint = (^(v3))
                                _ = v84
                                _ = v85
                            } else {
                                var v86 int = (fi82 - 554)
                                switch ("\n\t" + "\n\t") {
                                case "hello":
                                    v83 = append(v83, (fi82 / ((32767) | 1)))
                                    fmt.Println(fi82, v86, v2)
                                    v1 = (v1 - uint16(v4))
                                    if true {
                                        if len(v83) > 0 {
                                            v83[0] = fi82
                                        }
                                        fmt.Println(v2, v4)
                                    } else {
                                        if false {
                                            var v87 float64 = float64(0.0)
                                            _ = v87
                                        } else {
                                            if (9223372036854775807 >= 9223372036854775807) {
                                                v88 := v5[""]
                                                var v89 uint = 604
                                                _ = v88
                                                _ = v89
                                            } else {
                                                if (918 == 1) {
                                                    var v90 uint = 42
                                                    if false {
                                                        var v91 int32 = 2147483647
                                                        var v92 uint32 = uint32(v4)
                                                        var v93 uint8 = (v4 ^ (v4 | v4))
                                                        _ = v91
                                                        _ = v92
                                                        _ = v93
                                                    } else {
                                                        v4 = v4
                                                        v94 := v5[("" + "foo bar")]
                                                        if ("a" == "\n\t") {
                                                            switch ("hello" + "hello") {
                                                            case "":
                                                                var v95 uint8 = v4
                                                                _ = v95
                                                            default:
                                                                v96 := v6[42]
                                                                _ = v96
                                                            }
                                                            var v97 int64 = 1
                                                            var v98 uint16 = (v1 >> ((0) & 63))
                                                            _ = v97
                                                            _ = v98
                                                        }
                                                        _ = v94
                                                    }
                                                    _ = v90
                                                }
                                            }
                                        }
                                        v83 = append(v83, (v86 | 42))
                                    }
                                    var v99 int16 = int16(v2)
                                    _ = v99
                                case "":
                                    v4 -= v4
                                case "\n\t":
                                    fmt.Println(v2, v2, v86)
                                }
                                var v100 float64 = 1.5
                                _ = v100
                                _ = v86
                            }
                            _ = fi82
                            _ = v83
                        }
                        if false {
                            for fi101 := 0; fi101 < 1; fi101++ {
                                var v102 int16 = 0
                                _ = fi101
                                _ = v102
                            }
                        }
                        for fi103 := 0; fi103 < 8; fi103++ {
                            if (!(false && false)) {
                                v6[0] = ("a" <= "")
                                if (0.1 < (1e10 - 1e10)) {
                                    fmt.Println(v2)
                                    var v104 uint32 = 2147483648
                                    var v105 int8 = 42
                                    _ = v104
                                    _ = v105
                                }
                                if (!(1.5 < float32(float64(1e10)))) {
                                    var v106 int32 = int32(v4)
                                    fmt.Println(v3, v106, fi103)
                                    var v107 float64 = 1e10
                                    fmt.Println(fi81, fi81)
                                    _ = v106
                                    _ = v107
                                }
                            } else {
                                var v108 uint16 = v1
                                var v109 int8 = 0
                                _ = v108
                                _ = v109
                            }
                            if false {
                                var v110 float64 = 3.14159
                                if false {
                                    if false {
                                        v111 := GenS1{F0: (v2 & genFn0()), F1: v4}
                                        switch ("hello" + "\n\t") {
                                        case "":
                                            v112 := []int{}
                                            _ = v112
                                        case "foo bar":
                                            var v113 int8 = 13
                                            _ = v113
                                        }
                                        v1 += uint16(v2)
                                        _ = v111
                                    }
                                    fmt.Println(v2, v4)
                                } else {
                                    v114 := GenS6{F0: -1}
                                    if false {
                                        fmt.Println(fi103, v110)
                                    }
                                    _ = v114
                                }
                                v1 += 1
                                var v115 uint16 = (v1 >> (((v3 << ((v3) & 63))) & 63))
                                _ = v110
                                _ = v115
                            } else {
                                v4 += (^(v4))
                                switch ("" + "\n\t") {
                                case "\n\t":
                                    if ((42 == 42) && false) {
                                        switch int(v2) {
                                        case 32768:
                                            v3 = 2
                                        }
                                        fmt.Println(v1)
                                    }
                                default:
                                    fmt.Println(fi81, v1)
                                }
                                var v116 bool = ((float64(((((-2.25 - 3.14159) / 1.0) / 1e10) * 0.0)) + ((0.0 * 0.1) + (1e10 * 3.14159))) != (-2.25 - (0.1 / (0.1 - ((3.14159 + 3.14159) / 1.0)))))
                                var v117 uint16 = v1
                                v4 += 240
                                _ = v116
                                _ = v117
                            }
                            _ = fi103
                        }
                        _ = fi81
                    }
                default:
                    v4 = 2
                }
            } else {
                var v118 string = ("hello" + "")
                _ = v118
            }
        }
    }
    for fi119 := 0; fi119 < 3; fi119++ {
        var v120 uint32 = genFn0()
        fmt.Println(v3, v1, v1)
        var v121 uint16 = 354
        _ = fi119
        _ = v120
        _ = v121
    }
    delete(v5, ((("foo bar" + "a") + "foo bar") + ("a" + "")))
    v4 -= v4
    for fi122 := 0; fi122 < 8; fi122++ {
        v123 := GenS7{F0: v3, F1: 42, F2: (^(v3))}
        for fi124 := 0; fi124 < 11; fi124++ {
            v1 = (^(v1))
            delete(v5, "\n\t")
            fmt.Println(v1, v3, fi124)
            var v125 float32 = float32((0.1 * ((3.14159 / 0.1) - 0.1)))
            _ = fi124
            _ = v125
        }
        v3 = (v3 / (((v3 * (v3 >> ((v3) & 63)))) | 1))
        _ = fi122
        _ = v123
    }
    v126 := []int{128, -2147483648}
    for fi127 := 0; fi127 < 17; fi127++ {
        fmt.Println(v4, v2, v4)
        fmt.Println(v2)
        _ = fi127
    }
    var v128 int32 = 606
    switch "a" {
    case "hello":
        v4 = (v4 * 127)
    }
    if false {
        if true {
            var v129 int8 = 48
            v129 -= (v129 + (^(v129)))
            _ = v129
        } else {
            var v130 int64 = -779
            _ = v130
        }
        var v131 float64 = 0.0
        var v132 int = 243
        _ = v131
        _ = v132
    }
    if (!true) {
        v2 += 2147483647
        var v133 int32 = (v128 << ((v3) & 63))
        var v134 uint8 = (^(v4))
        _ = v133
        _ = v134
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v126:", len(v126), v126)
    fmt.Println("v128:", v128)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v126
    _ = v128
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
