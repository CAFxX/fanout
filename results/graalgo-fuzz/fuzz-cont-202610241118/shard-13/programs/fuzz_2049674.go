package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int64 }
type GenS2 struct { F0 uint; F1 int8; F2 int }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 int8; F2 uint8 }
type GenS6 struct { F0 int16; F1 int }
type GenS7 struct { F0 int16; F1 bool; F2 uint64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int64; F1 int; F2 uint8 }
type GenS10 struct { F0 uint16 }
func genFn0(p0 int8, p1 int8) uint {
    if (!true) {
        var v6 int16 = 0
        if (false || (!(1e10 != 1e10))) {
            if (true && (1.5 == 3.14159)) {
                var v7 uint16 = 255
                var v8 int64 = 9223372036854775807
                var v9 int64 = (v8 / ((v8) | 1))
                _ = v7
                _ = v8
                _ = v9
            } else {
                switch "" {
                case "a":
                    var v10 uint8 = 128
                    _ = v10
                case "foo bar":
                    if true {
                        for fi11 := 0; fi11 < 20; fi11++ {
                            v12 := GenS1{F0: 855, F1: -833, F2: -413}
                            _ = fi11
                            _ = v12
                        }
                        p0 -= p0
                        v6 -= (v6 - -1)
                    } else {
                        var v13 bool = false
                        _ = v13
                    }
                case "\n\t":
                    v14 := []int{175}
                    _ = v14
                default:
                    for fi15 := 0; fi15 < 14; fi15++ {
                        switch "" {
                        case "":
                            p0 = -1
                        case "hello":
                            var v16 uint16 = 482
                            _ = v16
                        default:
                            p0 -= p1
                        }
                        _ = fi15
                    }
                }
                p1 -= 42
                fmt.Println(p1, p0, p1)
            }
            v17 := []int{42, 221}
            p0 += p1
            _ = v17
        }
        if true {
            v18 := []int{-400, 1, -876, 603}
            _ = len(v18)
            for fi19 := 0; fi19 < 14; fi19++ {
                var v20 int16 = (v6 * v6)
                _ = fi19
                _ = v20
            }
            _ = v18
        } else {
            var v21 uint16 = 559
            _ = v21
        }
        _ = v6
    }
    p0 = (^(p0))
    v22 := GenS1{F0: 2, F1: -880, F2: -294}
    _ = p0
    _ = p1
    _ = v22
    return 429
}
func genFn1(p0 bool, p1 int32) uint8 {
    v23 := []int{}
    if (!(!p0)) {
        v24 := GenS1{F0: genFn0(-1, 31), F1: p1, F2: -1}
        p1 = p1
        for fi25 := 0; fi25 < 20; fi25++ {
            p1 = (p1 - p1)
            _ = fi25
        }
        _ = v24
    }
    var v26 int16 = 42
    v27 := GenS1{F0: genFn0(0, -1), F1: p1, F2: -1}
    _ = p0
    _ = p1
    _ = v23
    _ = v26
    _ = v27
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{}
    v3 := make(map[int]int)
    var v4 string = ("foo bar" + ("hello" + "hello"))
    v5 := make(map[string]int)
    for fi28 := 0; fi28 < 10; fi28++ {
        var v29 uint16 = 65535
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        v2 = append(v2, (((v4 + "a") + v4) + (v4 + "hello")))
        v31 := GenS1{F0: 42, F1: 2147483647, F2: 1}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi30
        _ = v31
    }
    fmt.Println(v4)
    if true {
        for fi32 := 0; fi32 < 13; fi32++ {
            var v33 bool = true
            var v34 uint = 42
            for fi35 := 0; fi35 < 5; fi35++ {
                var v36 string = v4
                var v37 float64 = 0.1
                _ = fi35
                _ = v36
                _ = v37
            }
            _ = fi32
            _ = v33
            _ = v34
        }
        for fi38 := 0; fi38 < 8; fi38++ {
            var v39 int64 = -1
            var v40 int64 = v39
            for fi41 := 0; fi41 < 1; fi41++ {
                v42 := v3[(^(fi38))]
                fmt.Println(v39, v42, v4)
                var v43 int64 = (v39 << ((42) & 63))
                _ = fi41
                _ = v42
                _ = v43
            }
            _ = fi38
            _ = v39
            _ = v40
        }
    }
    v44 := GenS1{F0: 255, F1: 2147483647, F2: -9223372036854775808}
    switch 32768 {
    case 0:
        var v45 float64 = 1.5
        _ = v45
    case 973:
        var v46 uint64 = 42
        _ = v46
    default:
        if false {
            var v47 int64 = 1
            _ = v47
        } else {
            var v48 uint = 2
            _ = v48
        }
    }
    fmt.Println(v4)
    switch v4 {
    case "":
        fmt.Println(v4)
    case "\n\t":
        fmt.Println(v4)
    case "foo bar":
        v49 := GenS1{F0: genFn0(42, 127), F1: 596, F2: 740}
        _ = v49
    }
    var v50 uint8 = genFn1((!((false && (!true)) || (((-2.25 / 1.0) + 1e10) > -2.25))), 791)
    v4 = v4
    var v51 int = 4294967295
    v52 := GenS2{F0: uint(v50), F1: 42, F2: -324}
    for fi53 := 0; fi53 < 20; fi53++ {
        v4 = ((((("a" + (v4 + v4)) + "") + v4) + "\n\t") + (v4 + v4))
        var v54 int16 = -245
        v4 = ""
        _ = fi53
        _ = v54
    }
    switch "\n\t" {
    case "\n\t":
        v51 += v51
    case "hello":
        switch 146 {
        case -9223372036854775808:
            fmt.Println(v50, v50, v4)
        case 4294967295:
            v51 -= v51
        case -1:
            for fi55 := 0; fi55 < 7; fi55++ {
                v4 = "hello"
                _ = fi55
            }
        default:
            var v56 uint8 = (v50 / ((1) | 1))
            _ = v56
        }
    case "foo bar":
        switch (v51 % ((v51) | 1)) {
        case -9223372036854775808:
            v57 := GenS3{F0: 1, F1: 18446744073709551615}
            _ = v57
        case -386:
            var v58 uint32 = 996
            _ = v58
        case -2147483648:
            var v59 uint64 = 582
            _ = v59
        default:
            if (!true) {
                for fi60 := 0; fi60 < 10; fi60++ {
                    for fi61 := 0; fi61 < 19; fi61++ {
                        if (127 <= 0) {
                            if (1000 >= 0) {
                                if len(v2) > 0 {
                                    v2[0] = v4
                                }
                            }
                            v51 = (fi61 * 9223372036854775807)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v50 -= 42
                            var v62 int = 0
                            _ = v62
                        } else {
                            var v63 int32 = 2147483647
                            switch v51 {
                            case 255:
                                delete(v5, ("" + ("foo bar" + ((v4 + "") + v4))))
                            case 128:
                                if (false && ((32767 == (v51 * (^(v51)))) && true)) {
                                    v50 = (v50 >> ((genFn0(-128, 0)) & 63))
                                    v63 = -415
                                    if ((v63 << ((18446744073709551615) & 63)) <= (v63 & (v63 + (-(v63))))) {
                                        switch (v51 + (fi60 + v51)) {
                                        case 128:
                                            var v64 string = (((("a" + (v4 + ("foo bar" + ("" + "hello")))) + (((v4 + (v4 + "a")) + ("foo bar" + "")) + (v4 + "hello"))) + v4) + v4)
                                            _ = v64
                                        case 42:
                                            v50 -= v50
                                        default:
                                            var v65 uint8 = (v50 ^ 1)
                                            _ = v65
                                        }
                                        var v66 uint64 = 72
                                        fmt.Println(fi61)
                                        _ = v66
                                    } else {
                                        var v67 uint64 = 1
                                        v4 = v4
                                        switch 32767 {
                                        case 4294967295:
                                            if (42 == 1) {
                                                var v68 string = v4
                                                v69 := GenS2{F0: 0, F1: 100, F2: (fi61 * (^(fi61)))}
                                                var v70 int16 = 536
                                                _ = v68
                                                _ = v69
                                                _ = v70
                                            }
                                        default:
                                            if false {
                                                v63 += 2147483647
                                                v2 = append(v2, "hello")
                                            } else {
                                                v71 := GenS4{F0: v51}
                                                var v72 int32 = (v63 << ((0) & 63))
                                                _ = v71
                                                _ = v72
                                            }
                                        }
                                        _ = v67
                                    }
                                }
                            }
                            fmt.Println(v4)
                            _ = v63
                        }
                        _ = fi61
                    }
                    _ = fi60
                }
                switch ((v4 + v4) + v4) {
                case "":
                    v50 += (v50 % (((^(v50))) | 1))
                case "a":
                    _ = len(v1)
                default:
                    var v73 int8 = 1
                    _ = v73
                }
            } else {
                switch (v51 * v51) {
                case 128:
                    v3[-9223372036854775808] = v51
                case 781:
                    fmt.Println(v51)
                }
                v51 -= 0
                var v74 float64 = 0.1
                _ = v74
            }
        }
    default:
        if true {
            var v75 float32 = 3.14159
            _ = v75
        } else {
            if false {
                for fi76 := 0; fi76 < 9; fi76++ {
                    for fi77 := 0; fi77 < 10; fi77++ {
                        v78 := GenS1{F0: 42, F1: 2147483647, F2: -785}
                        _ = fi77
                        _ = v78
                    }
                    _ = fi76
                }
            } else {
                switch -9223372036854775808 {
                case -693:
                    v79 := GenS3{F0: 2, F1: 110}
                    _ = v79
                case 921:
                    if ((0.0 != (1e10 / -2.25)) && ((genFn1(false, -2147483648) >= (v50 | (v50 / ((v50) | 1)))) || ((true && true) || (v51 > (v51 + 798))))) {
                        v80 := GenS5{F0: v4, F1: 1, F2: genFn1(true, 2147483647)}
                        _ = v80
                    } else {
                        if len(v2) > 0 {
                            v2[0] = "\n\t"
                        }
                    }
                }
                switch v51 {
                case 0:
                    fmt.Println(v50)
                case -1:
                    v4 = v4
                case 128:
                    var v81 int8 = -1
                    _ = v81
                default:
                    var v82 uint8 = genFn1(true, 2147483647)
                    _ = v82
                }
                fmt.Println(v4, v4, v4)
            }
            var v83 string = v4
            _ = v83
        }
    }
    if (v51 == v51) {
        if (v51 > 155) {
            var v84 uint16 = 65535
            if (510 <= -2147483648) {
                for fi85 := 0; fi85 < 14; fi85++ {
                    if (-891 >= 499) {
                        for fi86 := 0; fi86 < 2; fi86++ {
                            v87 := GenS3{F0: 965, F1: 947}
                            if (v84 < (v84 | v84)) {
                                var v88 bool = (1 <= -128)
                                fmt.Println(fi86)
                                var v89 int = fi86
                                _ = v88
                                _ = v89
                            }
                            var v90 float64 = -2.25
                            if ((4294967295 == 2) || true) {
                                _ = len(v2)
                                var v91 string = (("" + v4) + v4)
                                var v92 int8 = 127
                                v1[v91] = 2147483647
                                _ = v91
                                _ = v92
                            } else {
                                _ = len(v2)
                                var v93 uint8 = 195
                                v94 := GenS6{F0: 32767, F1: fi85}
                                _ = v93
                                _ = v94
                            }
                            _ = fi86
                            _ = v87
                            _ = v90
                        }
                        _ = len(v2)
                    }
                    v51 = 128
                    switch ("hello" + v4) {
                    case "foo bar":
                        v1[("foo bar" + v4)] = 4294967295
                    case "\n\t":
                        fmt.Println(v51, v51, v84)
                    default:
                        var v95 uint = uint(v84)
                        _ = v95
                    }
                    _ = fi85
                }
                fmt.Println(v51, v50, v51)
            } else {
                if ((0.1 - 1e10) > 0.1) {
                    if true {
                        v96 := GenS7{F0: -32768, F1: true, F2: 0}
                        v50 += genFn1((2147483647 > 659), 2147483647)
                        var v97 float32 = 0.0
                        _ = v96
                        _ = v97
                    } else {
                        v4 = ("foo bar" + v4)
                    }
                    if (float64((3.14159 + 3.14159)) == float64(float32(0.0))) {
                        if (-201 >= -125) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v50 -= genFn1((-88 != 0), 2147483647)
                            for fi98 := 0; fi98 < 14; fi98++ {
                                v4 = (v4 + v4)
                                switch "foo bar" {
                                case "a":
                                    v50 = (v50 >> ((1) & 63))
                                default:
                                    v4 = v4
                                }
                                if (v4 < (v4 + ("\n\t" + v4))) {
                                    v99 := GenS5{F0: v4, F1: 0, F2: 14}
                                    for fi100 := 0; fi100 < 7; fi100++ {
                                        v101 := GenS8{F0: 134}
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        switch v4 {
                                        case "":
                                            v51 = fi100
                                        }
                                        _ = fi100
                                        _ = v101
                                    }
                                    var v102 int8 = 12
                                    _ = v102
                                    _ = v99
                                } else {
                                    var v103 int64 = -777
                                    _ = v103
                                }
                                _ = fi98
                            }
                        }
                        v4 = (v4 + "hello")
                        switch v4 {
                        case "\n\t":
                            v84 = v84
                        default:
                            v104 := GenS9{F0: 731, F1: (-(v51)), F2: v50}
                            _ = v104
                        }
                    } else {
                        v105 := GenS10{F0: v84}
                        if (!false) {
                            if (1 < 1) {
                                v106 := GenS7{F0: 405, F1: (651 != 148), F2: 4294967295}
                                _ = v106
                            }
                            if (-84 > 614) {
                                var v107 int = (v51 * (^(v51)))
                                _ = v107
                            } else {
                                v4 = (v4 + v4)
                                v84 = 1
                            }
                        } else {
                            var v108 int16 = -1
                            _ = v108
                        }
                        if false {
                            fmt.Println(v4, v51, v84)
                            for fi109 := 0; fi109 < 17; fi109++ {
                                delete(v1, "foo bar")
                                fmt.Println(v50, v51, fi109)
                                _ = fi109
                            }
                        }
                        _ = v105
                    }
                    fmt.Println(v51, v84)
                }
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
                switch v4 {
                case "foo bar":
                    var v110 uint = 42
                    _ = v110
                case "\n\t":
                    if (18446744073709551615 == 962) {
                        delete(v5, v4)
                    }
                case "":
                    var v111 int32 = 2147483647
                    _ = v111
                }
            }
            _ = v84
        }
    }
    for fi112 := 0; fi112 < 2; fi112++ {
        for fi113 := 0; fi113 < 11; fi113++ {
            fmt.Println(fi113, v51, fi112)
            _ = fi113
        }
        if false {
            v114 := GenS4{F0: (v51 / ((fi112) | 1))}
            for fi115 := 0; fi115 < 14; fi115++ {
                v50 -= v50
                v4 = (v4 + v4)
                _ = fi115
            }
            _ = v114
        } else {
            var v116 int32 = 2147483647
            for fi117 := 0; fi117 < 14; fi117++ {
                v2 = append(v2, v4)
                var v118 int = 2147483647
                _ = fi117
                _ = v118
            }
            v119 := GenS2{F0: 1, F1: 100, F2: 223}
            _ = v116
            _ = v119
        }
        v1[(v4 + v4)] = 2147483648
        _ = fi112
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v50
    _ = v51
    _ = v52
}
