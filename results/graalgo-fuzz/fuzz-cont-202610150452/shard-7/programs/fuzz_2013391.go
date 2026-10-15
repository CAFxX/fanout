package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 uint; F2 string }
type GenS6 struct { F0 bool; F1 uint16; F2 int }
type GenS7 struct { F0 int32; F1 string }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int64 }
func genFn0(p0 uint8, p1 bool) uint64 {
    v7 := GenS1{F0: 786, F1: 209}
    var v8 uint8 = (p0 / (((p0 - 2)) | 1))
    switch "a" {
    case "a":
        v9 := GenS2{F0: 559}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 uint, p1 int8) uint {
    switch ("a" + "a") {
    case "":
        if false {
            for fi10 := 0; fi10 < 16; fi10++ {
                var v11 int = (fi10 << (((p0 | 18446744073709551615)) & 63))
                if (!true) {
                    switch fi10 {
                    case -345:
                        var v12 string = ("" + "hello")
                        _ = v12
                    case 32768:
                        switch "hello" {
                        case "foo bar":
                            var v13 int64 = 832
                            _ = v13
                        case "":
                            if ((0.0 + (0.1 + -2.25)) != 3.14159) {
                                for fi14 := 0; fi14 < 5; fi14++ {
                                    fmt.Println(p1, v11)
                                    _ = fi14
                                }
                                var v15 bool = (-9223372036854775808 >= 1)
                                _ = v15
                            } else {
                                for fi16 := 0; fi16 < 16; fi16++ {
                                    p0 = p0
                                    fmt.Println(p0, fi10)
                                    v17 := []int{841, fi16}
                                    _ = fi16
                                    _ = v17
                                }
                            }
                        default:
                            for fi18 := 0; fi18 < 16; fi18++ {
                                p0 = (p0 / ((2) | 1))
                                var v19 float32 = float32(0.0)
                                if true {
                                    var v20 int16 = 1000
                                    v21 := []int{fi10, fi18, 0, fi18}
                                    _ = v20
                                    _ = v21
                                }
                                switch 42 {
                                case -171:
                                    switch "" {
                                    case "":
                                        fmt.Println(p1, p0)
                                    }
                                }
                                _ = fi18
                                _ = v19
                            }
                        }
                    case -2147483648:
                        for fi22 := 0; fi22 < 2; fi22++ {
                            v23 := []int{fi10, (fi10 ^ -2147483648), -623, (v11 * 32768)}
                            if false {
                                _ = len(v23)
                                switch "hello" {
                                case "a":
                                    v24 := GenS3{F0: 4294967295, F1: 1, F2: 829}
                                    _ = v24
                                default:
                                    if (!(!false)) {
                                        if len(v23) > 0 {
                                            v23[0] = (fi22 << ((p0) & 63))
                                        }
                                        switch (("hello" + "hello") + "foo bar") {
                                        case "hello":
                                            var v25 uint16 = 0
                                            _ = v25
                                        case "":
                                            p1 = p1
                                        case "\n\t":
                                            v26 := GenS4{F0: fi22}
                                            _ = v26
                                        default:
                                            switch ("hello" + "a") {
                                            case "a":
                                                switch ((("a" + "a") + "\n\t") + "") {
                                                case "\n\t":
                                                    if (!(666 != genFn0(124, true))) {
                                                        v27 := GenS4{F0: (fi22 % (((fi10 | (-(v11)))) | 1))}
                                                        var v28 string = ("hello" + (("a" + "hello") + "a"))
                                                        _ = v27
                                                        _ = v28
                                                    }
                                                case "foo bar":
                                                    if true {
                                                        v11 += 42
                                                    } else {
                                                        if len(v23) >= 2 {
                                                            v23 = v23[1:len(v23)]
                                                        }
                                                    }
                                                case "hello":
                                                    v11 += fi22
                                                }
                                            case "":
                                                switch "hello" {
                                                case "hello":
                                                    v29 := GenS5{F0: 1, F1: p0, F2: "hello"}
                                                    _ = v29
                                                case "":
                                                    p0 = p0
                                                }
                                            case "\n\t":
                                                var v30 uint32 = 4294967295
                                                _ = v30
                                            default:
                                                if len(v23) > 0 {
                                                    v23[0] = 2147483647
                                                }
                                            }
                                        }
                                    }
                                }
                                var v31 int64 = 9223372036854775807
                                _ = v31
                            } else {
                                if false {
                                    fmt.Println(p1)
                                } else {
                                    v32 := GenS5{F0: 42, F1: p0, F2: ("\n\t" + "\n\t")}
                                    if ((395 < 42) || false) {
                                        p0 = (p0 >> (((p0 & (p0 + 18446744073709551615))) & 63))
                                        v23 = append(v23, fi22)
                                    } else {
                                        v33 := GenS3{F0: 42, F1: 867, F2: -32768}
                                        if len(v23) >= 2 {
                                            v23 = v23[1:len(v23)]
                                        }
                                        _ = v33
                                    }
                                    p0 += p0
                                    switch "hello" {
                                    case "a":
                                        var v34 int64 = 813
                                        _ = v34
                                    case "foo bar":
                                        if len(v23) > 0 {
                                            v23[0] = fi10
                                        }
                                    default:
                                        p0 += p0
                                    }
                                    _ = v32
                                }
                                var v35 int16 = 0
                                var v36 int = 2147483647
                                _ = v35
                                _ = v36
                            }
                            _ = fi22
                            _ = v23
                        }
                    default:
                        p1 -= (-(p1))
                    }
                    v37 := GenS2{F0: -513}
                    var v38 int64 = 9223372036854775807
                    _ = v37
                    _ = v38
                }
                _ = fi10
                _ = v11
            }
            for fi39 := 0; fi39 < 3; fi39++ {
                switch ("a" + ("hello" + "\n\t")) {
                case "foo bar":
                    var v40 uint32 = 404
                    _ = v40
                case "":
                    p1 += (p1 - p1)
                default:
                    switch ("foo bar" + "hello") {
                    case "\n\t":
                        if true {
                            fmt.Println(fi39)
                            p0 = p0
                            v41 := []int{fi39, (fi39 % ((fi39) | 1)), -59, 128}
                            _ = v41
                        } else {
                            var v42 float64 = (((1.5 / -2.25) - (3.14159 / 3.14159)) + -2.25)
                            _ = v42
                        }
                    case "a":
                        for fi43 := 0; fi43 < 2; fi43++ {
                            var v44 uint64 = 978
                            _ = fi43
                            _ = v44
                        }
                    }
                }
                v45 := GenS5{F0: 42, F1: 465, F2: "hello"}
                v46 := []int{}
                _ = fi39
                _ = v45
                _ = v46
            }
        } else {
            p1 -= -79
            v47 := []int{232, 752}
            v48 := GenS6{F0: true, F1: 2, F2: 2147483647}
            _ = v47
            _ = v48
        }
    case "foo bar":
        var v49 uint = (p0 * (p0 >> ((793) & 63)))
        _ = v49
    default:
        var v50 int64 = -300
        _ = v50
    }
    v51 := GenS2{F0: 0}
    var v52 uint64 = 1
    for fi53 := 0; fi53 < 20; fi53++ {
        if true {
            var v54 string = "a"
            switch fi53 {
            case -1:
                var v55 uint32 = 2147483647
                _ = v55
            }
            _ = v54
        } else {
            if false {
                v56 := GenS4{F0: fi53}
                var v57 int64 = 525
                _ = v56
                _ = v57
            } else {
                for fi58 := 0; fi58 < 17; fi58++ {
                    p1 -= (p1 >> ((42) & 63))
                    _ = fi58
                }
                fmt.Println(p0, p1)
                if true {
                    v59 := []int{fi53}
                    for fi60 := 0; fi60 < 17; fi60++ {
                        var v61 int64 = 9223372036854775807
                        _ = fi60
                        _ = v61
                    }
                    _ = v59
                }
            }
        }
        _ = fi53
    }
    switch "hello" {
    case "hello":
        v62 := GenS4{F0: 42}
        _ = v62
    case "foo bar":
        p1 = (p1 * p1)
    case "":
        v63 := []int{-1, 981, 42, 199}
        _ = v63
    }
    _ = p0
    _ = p1
    _ = v51
    _ = v52
    return 0
}

func main() {
    var v1 int8 = 0
    v2 := []string{"a", "hello"}
    var v3 uint = 18446744073709551615
    var v4 uint16 = 849
    v5 := make(map[string]bool)
    v6 := make(map[uint32]int)
    var v64 int64 = 1
    v64 = (v64 - (v64 / ((v64) | 1)))
    fmt.Println(v3, v1)
    for fi65 := 0; fi65 < 16; fi65++ {
        v64 = int64(v4)
        for fi66 := 0; fi66 < 8; fi66++ {
            if false {
                switch "a" {
                case "":
                    fmt.Println(v64, v3, fi65)
                case "foo bar":
                    if (!false) {
                        v64 = (v64 >> (((^(v3))) & 63))
                        var v67 float64 = float64((3.14159 * 1e10))
                        if false {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v68 := GenS5{F0: 42, F1: v3, F2: ""}
                            _ = v68
                        } else {
                            fmt.Println(fi65, fi66)
                            v69 := v5[("a" + "")]
                            var v70 uint8 = 128
                            if (0 == 857) {
                                if v69 {
                                    var v71 uint64 = 0
                                    _ = v71
                                }
                                v70 += (v70 << (((v3 ^ 65535)) & 63))
                            }
                            _ = v69
                            _ = v70
                        }
                        _ = v67
                    }
                }
            } else {
                v64 = (v64 >> ((v3) & 63))
                v64 += v64
                switch fi65 {
                case 128:
                    var v72 float64 = 3.14159
                    _ = v72
                case 42:
                    var v73 int16 = 32767
                    _ = v73
                case -1:
                    var v74 uint8 = 127
                    _ = v74
                default:
                    switch ("hello" + ("" + "foo bar")) {
                    case "\n\t":
                        var v75 int32 = 1
                        v1 -= v1
                        _ = v75
                    case "":
                        v2 = append(v2, "\n\t")
                    default:
                        fmt.Println(v3, v4)
                    }
                }
                fmt.Println(v3, v3)
                var v76 int8 = v1
                var v77 uint8 = 0
                if (v4 > 255) {
                    if (true || true) {
                        v78 := GenS7{F0: -2147483648, F1: "\n\t"}
                        v3 = 1
                        var v79 int = -2147483648
                        _ = v78
                        _ = v79
                    } else {
                        if true {
                            fmt.Println(v3, v64, v1)
                            fmt.Println(v4, fi65)
                            var v80 bool = ("foo bar" > ("a" + "\n\t"))
                            v81 := v6[2147483647]
                            _ = v80
                            _ = v81
                        }
                        fmt.Println(fi65, v77)
                    }
                    fmt.Println(v76, fi66)
                    v82 := GenS6{F0: false, F1: 0, F2: fi65}
                    _ = v82
                } else {
                    var v83 int16 = 42
                    switch "" {
                    case "foo bar":
                        switch fi65 {
                        case 127:
                            _ = len(v2)
                        case -471:
                            v1 = 42
                        case -9223372036854775808:
                            _ = len(v6)
                        }
                    case "a":
                        if true {
                            v84 := v6[42]
                            var v85 uint16 = 65535
                            v6[2147483647] = 396
                            _ = v84
                            _ = v85
                        }
                    case "":
                        var v86 float32 = ((0.1 * 0.1) + 3.14159)
                        _ = v86
                    default:
                        _ = len(v5)
                    }
                    _ = v83
                }
                _ = v76
                _ = v77
            }
            var v87 float32 = 1e10
            _ = fi66
            _ = v87
        }
        _ = fi65
    }
    var v88 string = ("" + "")
    switch (v88 + "hello") {
    case "hello":
        for fi89 := 0; fi89 < 18; fi89++ {
            if ((v4 - v4) == v4) {
                var v90 uint64 = 42
                fmt.Println(v1)
                v91 := GenS8{F0: false}
                _ = v90
                _ = v91
            } else {
                if (28 < 95) {
                    var v92 int16 = 1
                    _ = v92
                } else {
                    if (v88 >= (v88 + v88)) {
                        fmt.Println(v4)
                        var v93 bool = (true && true)
                        _ = v93
                    }
                }
                switch 32767 {
                case 0:
                    if (true || false) {
                        for fi94 := 0; fi94 < 19; fi94++ {
                            if false {
                                if false {
                                    var v95 uint = v3
                                    fmt.Println(v64, v4)
                                    if false {
                                        v96 := GenS1{F0: (^(v4)), F1: 65535}
                                        _ = v96
                                    } else {
                                        fmt.Println(v88, v3, v95)
                                        v95 -= (v3 - v3)
                                    }
                                    _ = v95
                                } else {
                                    var v97 string = v88
                                    switch fi89 {
                                    case 32767:
                                        var v98 bool = (v4 == (v4 & v4))
                                        _ = v98
                                    }
                                    if (true && true) {
                                        v99 := GenS4{F0: fi94}
                                        var v100 uint32 = 1
                                        _ = v100
                                        _ = v99
                                    } else {
                                        var v101 bool = true
                                        fmt.Println(v3, v1, fi89)
                                        v102 := GenS2{F0: -276}
                                        _ = v101
                                        _ = v102
                                    }
                                    _ = v97
                                }
                            }
                            _ = fi94
                        }
                        fmt.Println(fi89, v1)
                    }
                case 2147483647:
                    if true {
                        var v103 float64 = (0.1 - -2.25)
                        for fi104 := 0; fi104 < 6; fi104++ {
                            v105 := v6[828]
                            _ = fi104
                            _ = v105
                        }
                        if (0.0 >= ((float32(float64((3.14159 * 1e10))) - 0.1) * 0.1)) {
                            switch v88 {
                            case "foo bar":
                                var v106 int16 = 32767
                                _ = v106
                            case "hello":
                                v1 += 127
                            case "a":
                                v3 += (v3 - (v3 / ((v3) | 1)))
                            default:
                                var v107 uint64 = 0
                                _ = v107
                            }
                            v6[2147483647] = (fi89 << ((281) & 63))
                            v2 = append(v2, v88)
                        } else {
                            _ = len(v6)
                            var v108 int64 = v64
                            v1 = v1
                            _ = v108
                        }
                        _ = v103
                    } else {
                        if (fi89 == fi89) {
                            var v109 float64 = float64(-2.25)
                            v1 = -1
                            _ = v109
                        }
                        var v110 float32 = (0.1 - 1.5)
                        _ = v110
                    }
                case 32768:
                    v4 -= uint16(v3)
                default:
                    v111 := GenS8{F0: false}
                    _ = v111
                }
                var v112 uint64 = 18446744073709551615
                _ = v112
            }
            for fi113 := 0; fi113 < 19; fi113++ {
                _ = len(v5)
                _ = fi113
            }
            _ = fi89
        }
    case "foo bar":
        v114 := GenS9{F0: (v64 >> ((v3) & 63))}
        _ = v114
    case "a":
        _ = len(v5)
    }
    var v115 uint8 = uint8(v3)
    var v116 uint16 = v4
    var v117 uint16 = (v116 / ((659) | 1))
    v116 -= v4
    if false {
        v3 += (^(v3))
    } else {
        switch (v88 + v88) {
        case "":
            var v118 int8 = (^(v1))
            _ = v118
        case "hello":
            v115 += uint8(v3)
        default:
            if true {
                v1 = (v1 % ((v1) | 1))
            } else {
                v117 += 65535
                var v119 float64 = -2.25
                _ = v119
            }
        }
        var v120 uint16 = v116
        _ = v120
    }
    fmt.Println("v1:", v1)
    fmt.Println("v115:", v115)
    fmt.Println("v116:", v116)
    fmt.Println("v117:", v117)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v64:", v64)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v115
    _ = v116
    _ = v117
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v88
}
