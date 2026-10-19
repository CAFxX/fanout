package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int16; F2 bool }
type GenS3 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16; F1 uint64 }
type GenS6 struct { F0 uint; F1 uint16; F2 bool }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS11 struct { F0 uint32; F1 int16 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint64; F1 int8; F2 int16 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 string }
type GenS16 struct { F0 uint32; F1 uint64; F2 int }
func genFn0() bool {
    v5 := []int{}
    if len(v5) > 0 {
        v5[0] = 128
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v6 := GenS1{F0: -1, F1: 663}
    var v7 float32 = 1e10
    _ = v5
    _ = v6
    _ = v7
    return (true || (1 > 255))
}
func genFn1(p0 string, p1 uint32, p2 uint64) int32 {
    p0 = "foo bar"
    for fi8 := 0; fi8 < 9; fi8++ {
        p1 += p1
        if (false && (-761 == (^(fi8)))) {
            v9 := []int{fi8}
            if len(v9) > 0 {
                v9[0] = fi8
            }
            _ = v9
        } else {
            var v10 uint8 = 255
            _ = v10
        }
        switch (("\n\t" + p0) + "") {
        case "a":
            v11 := GenS1{F0: -9223372036854775808, F1: 2}
            _ = v11
        default:
            if false {
                var v12 uint16 = 2
                if (false && true) {
                    v13 := []int{(fi8 + (fi8 << ((255) & 63)))}
                    p1 = 4294967295
                    v14 := GenS2{F0: 403, F1: -32768, F2: false}
                    _ = v13
                    _ = v14
                }
                _ = v12
            } else {
                v15 := []int{}
                p0 = p0
                v16 := GenS1{F0: 1, F1: 2}
                _ = v15
                _ = v16
            }
        }
        _ = fi8
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 uint = 650
        v19 := GenS3{F0: 4, F1: -1, F2: 2}
        _ = fi17
        _ = v18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = 275
    v3 := make(map[string]int)
    v4 := make(map[string]bool)
    delete(v3, ("foo bar" + "foo bar"))
    v20 := []int{9223372036854775807}
    v21 := GenS1{F0: 1, F1: 0}
    v2 += genFn1("", 1, 4294967295)
    v2 -= v2
    switch ("a" + "hello") {
    case "hello":
        fmt.Println(v2)
    case "":
        switch -423 {
        case 32768:
            fmt.Println(v2)
        case 365:
            v22 := GenS2{F0: 42, F1: 1, F2: (-986 >= -1)}
            _ = v22
        case -326:
            var v23 uint32 = 2147483648
            _ = v23
        }
    default:
        for fi24 := 0; fi24 < 14; fi24++ {
            switch ("a" + (("\n\t" + "hello") + "foo bar")) {
            case "a":
                switch fi24 {
                case 0:
                    v2 += v2
                }
            case "hello":
                v2 = v2
            case "":
                if (((!(2147483647 == 9223372036854775807)) || (true && (!true))) || true) {
                    var v25 int32 = v2
                    v2 -= v2
                    v25 -= v2
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 4; fi26++ {
                        fmt.Println(fi24, fi26, fi26)
                        fmt.Println(fi24)
                        _ = len(v1)
                        switch 298 {
                        case -810:
                            fmt.Println(fi24, v2)
                        default:
                            switch (^(fi24)) {
                            case 4294967295:
                                _ = len(v20)
                            case 9223372036854775807:
                                _ = len(v3)
                                var v27 int16 = 870
                                if len(v20) >= 2 {
                                    v20 = v20[1:len(v20)]
                                }
                                _ = v27
                            case -135:
                                if len(v20) >= 2 {
                                    v20 = v20[1:len(v20)]
                                }
                            }
                        }
                        if (!(float32(-2.25) >= 0.0)) {
                            v28 := GenS2{F0: 551, F1: -493, F2: (255 > 837)}
                            switch ("a" + "\n\t") {
                            case "foo bar":
                                var v29 int8 = -95
                                _ = v29
                            case "hello":
                                var v30 uint = 175
                                _ = v30
                            }
                            v2 += v2
                            _ = v28
                        }
                        _ = fi26
                    }
                }
            default:
                v2 -= 42
            }
            for fi31 := 0; fi31 < 17; fi31++ {
                if true {
                    if false {
                        if len(v20) >= 2 {
                            v20 = v20[1:len(v20)]
                        }
                    }
                }
                _ = fi31
            }
            var v32 int8 = 127
            _ = fi24
            _ = v32
        }
    }
    var v33 float32 = ((3.14159 + 1.5) / 0.1)
    var v34 uint8 = 2
    var v35 int32 = (v2 & (-(v2)))
    var v36 string = "\n\t"
    var v37 uint16 = 929
    for fi38 := 0; fi38 < 16; fi38++ {
        v35 = (v2 - v2)
        _ = fi38
    }
    if false {
        for fi39 := 0; fi39 < 3; fi39++ {
            for fi40 := 0; fi40 < 10; fi40++ {
                switch 128 {
                case -1:
                    v41 := v3["\n\t"]
                    _ = v41
                case 799:
                    v42 := GenS4{F0: (fi40 + (^(fi40)))}
                    var v43 int8 = 1
                    _ = v42
                    _ = v43
                default:
                    v44 := GenS4{F0: fi39}
                    _ = v44
                }
                v20 = append(v20, -2147483648)
                var v45 int8 = -128
                _ = fi40
                _ = v45
            }
            _ = fi39
        }
    } else {
        var v46 uint64 = 1
        _ = v46
    }
    if (v34 > (v34 ^ v34)) {
        var v47 int = 0
        if (float32(float64((v33 * 1e10))) <= (1.5 + v33)) {
            for fi48 := 0; fi48 < 11; fi48++ {
                if (v37 < (v37 % ((273) | 1))) {
                    if (!(!false)) {
                        v49 := v4["foo bar"]
                        for fi50 := 0; fi50 < 4; fi50++ {
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                            v33 += v33
                            _ = fi50
                        }
                        v51 := GenS4{F0: 4294967295}
                        _ = v49
                        _ = v51
                    } else {
                        fmt.Println(fi48, v35, v47)
                        v33 = v33
                        v33 = v33
                    }
                    for fi52 := 0; fi52 < 10; fi52++ {
                        v35 = (v35 + (^(v35)))
                        fmt.Println(v33)
                        v36 = "\n\t"
                        _ = fi52
                    }
                    switch "\n\t" {
                    case "":
                        for fi53 := 0; fi53 < 8; fi53++ {
                            var v54 float64 = (0.0 * 0.0)
                            _ = len(v20)
                            _ = fi53
                            _ = v54
                        }
                    default:
                        switch fi48 {
                        case 0:
                            delete(v1, 996)
                        default:
                            for fi55 := 0; fi55 < 18; fi55++ {
                                v3[v36] = (fi48 * 2147483647)
                                _ = fi55
                            }
                        }
                    }
                }
                _ = fi48
            }
        } else {
            for fi56 := 0; fi56 < 7; fi56++ {
                switch "hello" {
                case "hello":
                    switch ("foo bar" + v36) {
                    case "hello":
                        var v57 int8 = -128
                        _ = v57
                    case "":
                        var v58 float64 = 3.14159
                        _ = v58
                    case "\n\t":
                        for fi59 := 0; fi59 < 16; fi59++ {
                            var v60 uint32 = 517
                            var v61 float64 = 0.0
                            if (-1 != 42) {
                                if (v60 > (v60 - (v60 & uint32(v37)))) {
                                    if (!(("\n\t" > ((("a" + v36) + ("hello" + v36)) + v36)) || false)) {
                                        v34 = (v34 >> ((4294967295) & 63))
                                        fmt.Println(v47)
                                    }
                                } else {
                                    switch v36 {
                                    case "":
                                        fmt.Println(v35, v61, v37)
                                    case "\n\t":
                                        if (true || (!(42 < int16(fi59)))) {
                                            if false {
                                                var v62 float32 = v33
                                                _ = v62
                                            }
                                            fmt.Println(v2)
                                            var v63 uint = 715
                                            _ = v63
                                        } else {
                                            switch v36 {
                                            case "":
                                                if (!true) {
                                                    if (("hello" <= (v36 + v36)) || true) {
                                                        v36 = v36
                                                        if false {
                                                            v1[(v60 / (((v60 + 65535)) | 1))] = false
                                                            fmt.Println(v33, v2)
                                                            var v64 int32 = (v2 | v35)
                                                            _ = v64
                                                        } else {
                                                            if false {
                                                                if false {
                                                                    switch v36 {
                                                                    case "\n\t":
                                                                        var v65 int = v47
                                                                        _ = v65
                                                                    default:
                                                                        v66 := GenS5{F0: (v37 / ((65535) | 1)), F1: 4294967295}
                                                                        _ = v66
                                                                    }
                                                                }
                                                            } else {
                                                                v60 = 1
                                                            }
                                                            fmt.Println(v60, v37, v60)
                                                        }
                                                        v33 += (v33 * (v33 * v33))
                                                        var v67 int = (-(fi59))
                                                        switch v36 {
                                                        case "foo bar":
                                                            if (v36 > ("foo bar" + ((v36 + v36) + v36))) {
                                                                v67 += (fi59 % ((4294967295) | 1))
                                                                v61 -= -2.25
                                                            }
                                                        case "a":
                                                            if false {
                                                                v37 = v37
                                                            }
                                                        case "hello":
                                                            v37 += uint16(v60)
                                                        }
                                                        _ = v67
                                                    } else {
                                                        fmt.Println(v61, v37)
                                                        switch -793 {
                                                        case 128:
                                                            v61 = (v61 / v61)
                                                        }
                                                        if (false || (true || (584 >= v60))) {
                                                            if true {
                                                                v36 = v36
                                                                switch 2147483647 {
                                                                case -494:
                                                                    switch v36 {
                                                                    case "foo bar":
                                                                        v47 = (fi56 ^ 42)
                                                                    default:
                                                                        v68 := GenS5{F0: 1, F1: 2}
                                                                        _ = v68
                                                                    }
                                                                    switch int(v37) {
                                                                    case 38:
                                                                        switch v36 {
                                                                        case "a":
                                                                            var v69 int64 = 1
                                                                            _ = v69
                                                                        default:
                                                                            if true {
                                                                                var v70 uint16 = 872
                                                                                if (381 >= 42) {
                                                                                    v71 := GenS6{F0: 0, F1: (v37 + (v37 & v70)), F2: (355 > 9223372036854775807)}
                                                                                    _ = v71
                                                                                }
                                                                                v72 := v4[v36]
                                                                                _ = v70
                                                                                _ = v72
                                                                            }
                                                                        }
                                                                    case -140:
                                                                        var v73 int32 = (v2 | 2147483647)
                                                                        _ = v73
                                                                    default:
                                                                        switch "a" {
                                                                        case "\n\t":
                                                                            v4[v36] = (!false)
                                                                        case "hello":
                                                                            if true {
                                                                                v1[v60] = (0 == int64(v34))
                                                                            } else {
                                                                                v74 := v3[""]
                                                                                var v75 uint16 = v37
                                                                                fmt.Println(v36, v36, v75)
                                                                                _ = v74
                                                                                _ = v75
                                                                            }
                                                                        default:
                                                                            v33 = v33
                                                                        }
                                                                    }
                                                                case -283:
                                                                    v76 := GenS1{F0: 0, F1: v37}
                                                                    _ = v76
                                                                case 333:
                                                                    switch fi56 {
                                                                    case 89:
                                                                        v60 = v60
                                                                    case 128:
                                                                        var v77 uint16 = (v37 | (v37 & v37))
                                                                        _ = v77
                                                                    default:
                                                                        v37 += uint16(v60)
                                                                    }
                                                                }
                                                                switch v36 {
                                                                case "\n\t":
                                                                    v36 = "hello"
                                                                case "":
                                                                    if len(v20) > 0 {
                                                                        v20[0] = v47
                                                                    }
                                                                }
                                                            } else {
                                                                v2 -= int32(v37)
                                                                v47 = (^(fi59))
                                                            }
                                                            if false {
                                                                v36 = v36
                                                                v37 -= v37
                                                            }
                                                            var v78 int8 = 26
                                                            _ = v78
                                                        }
                                                    }
                                                    v1[2147483648] = false
                                                    if false {
                                                        v33 -= v33
                                                        var v79 uint16 = (v37 << ((2) & 63))
                                                        _ = v79
                                                    } else {
                                                        if false {
                                                            v20 = append(v20, fi56)
                                                            var v80 uint16 = 257
                                                            _ = v80
                                                        }
                                                        v37 -= 32767
                                                        if ((!true) && (!false)) {
                                                            v1[v60] = ((!false) && (false && false))
                                                            switch ("a" + ("a" + v36)) {
                                                            case "":
                                                                var v81 uint32 = 2
                                                                _ = v81
                                                            case "foo bar":
                                                                if false {
                                                                    var v82 int16 = 32767
                                                                    if false {
                                                                        v83 := GenS3{F0: 128, F1: 0, F2: uint8(v37)}
                                                                        v47 -= v47
                                                                        _ = v83
                                                                    } else {
                                                                        delete(v3, v36)
                                                                        v84 := GenS4{F0: (fi56 * 9223372036854775807)}
                                                                        switch -2147483648 {
                                                                        case -2147483648:
                                                                            v85 := GenS3{F0: v34, F1: -128, F2: (v34 + (v34 * v34))}
                                                                            _ = v85
                                                                        case 9223372036854775807:
                                                                            var v86 int8 = -62
                                                                            _ = v86
                                                                        case 0:
                                                                            if false {
                                                                                var v87 uint = 103
                                                                                var v88 int16 = (-(v82))
                                                                                v34 += 255
                                                                                _ = v87
                                                                                _ = v88
                                                                            } else {
                                                                                v89 := GenS5{F0: v37, F1: 2}
                                                                                v60 = 2147483648
                                                                                if false {
                                                                                    v33 = (v33 + float32(v61))
                                                                                } else {
                                                                                    var v90 float32 = float32(1.5)
                                                                                    _ = v90
                                                                                }
                                                                                _ = v89
                                                                            }
                                                                        }
                                                                        _ = v84
                                                                    }
                                                                    if (false && true) {
                                                                        v91 := GenS5{F0: 32767, F1: 291}
                                                                        if (true || true) {
                                                                            if false {
                                                                                var v92 uint8 = (v34 | (v34 / ((127) | 1)))
                                                                                _ = v92
                                                                            }
                                                                        } else {
                                                                            var v93 uint64 = 843
                                                                            _ = v93
                                                                        }
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            if ((607 > 1) && true) {
                                                                                switch v36 {
                                                                                case "":
                                                                                    var v94 uint8 = (^(v34))
                                                                                    _ = v94
                                                                                case "a":
                                                                                    v35 += v2
                                                                                case "hello":
                                                                                    _ = len(v4)
                                                                                }
                                                                                var v95 bool = (false && true)
                                                                                v96 := GenS7{F0: 0}
                                                                                _ = v95
                                                                                _ = v96
                                                                            } else {
                                                                                if true {
                                                                                    v33 -= v33
                                                                                    fmt.Println(fi59, v82, v37)
                                                                                    v97 := GenS8{F0: -32768}
                                                                                    _ = v97
                                                                                } else {
                                                                                    v47 -= 32767
                                                                                    var v98 int8 = 1
                                                                                    var v99 uint = 829
                                                                                    _ = v98
                                                                                    _ = v99
                                                                                }
                                                                                switch int(v35) {
                                                                                case -9223372036854775808:
                                                                                    var v100 uint32 = (v60 >> ((65535) & 63))
                                                                                    _ = v100
                                                                                case -1:
                                                                                    _ = len(v1)
                                                                                default:
                                                                                    v82 = (v82 - v82)
                                                                                    v82 = v82
                                                                                    v61 += v61
                                                                                }
                                                                            }
                                                                        default:
                                                                            var v101 uint8 = v34
                                                                            _ = v101
                                                                        }
                                                                        _ = v91
                                                                    } else {
                                                                        var v102 int64 = -108
                                                                        var v103 bool = true
                                                                        v20 = append(v20, fi59)
                                                                        v103 = ((416 <= 539) && true)
                                                                        _ = v102
                                                                        _ = v103
                                                                    }
                                                                    _ = v82
                                                                }
                                                            }
                                                            if false {
                                                                var v104 uint64 = 4294967295
                                                                if true {
                                                                    var v105 int16 = 1000
                                                                    var v106 float64 = float64(((v33 / (3.14159 - float32((v61 * float64(3.14159))))) - 0.0))
                                                                    _ = v105
                                                                    _ = v106
                                                                } else {
                                                                    switch v36 {
                                                                    case "foo bar":
                                                                        v107 := GenS9{F0: ((false && true) || (((v33 - v33) / ((float32(1.5) / -2.25) - v33)) != (1.5 / v33)))}
                                                                        _ = v107
                                                                    }
                                                                    v36 = v36
                                                                    fmt.Println(v2)
                                                                    v108 := GenS2{F0: 65535, F1: 0, F2: (v60 >= v60)}
                                                                    v109 := GenS10{F0: v60, F1: 42, F2: (v2 * 795)}
                                                                    _ = v108
                                                                    _ = v109
                                                                }
                                                                _ = v104
                                                            } else {
                                                                v110 := GenS11{F0: v60, F1: int16(v34)}
                                                                v111 := GenS6{F0: 911, F1: (^(v37)), F2: true}
                                                                v37 -= (v37 | (v37 / ((v37) | 1)))
                                                                _ = v110
                                                                _ = v111
                                                            }
                                                        } else {
                                                            _ = len(v20)
                                                            switch 0 {
                                                            case 0:
                                                                var v112 bool = (!((0.0 > 1.5) && false))
                                                                _ = v112
                                                            }
                                                            v47 += (fi59 & v47)
                                                        }
                                                        v35 -= v2
                                                    }
                                                }
                                            case "hello":
                                                var v113 int = (v47 & (v47 | fi59))
                                                _ = v113
                                            case "foo bar":
                                                var v114 string = (v36 + "hello")
                                                _ = v114
                                            default:
                                                v35 = genFn1(v36, v60, 198)
                                            }
                                            switch v36 {
                                            case "a":
                                                if true {
                                                    v36 = (("a" + v36) + ((v36 + "\n\t") + v36))
                                                } else {
                                                    var v115 uint = 364
                                                    var v116 int = 32767
                                                    if len(v20) > 0 {
                                                        v20[0] = (fi56 / ((v47) | 1))
                                                    }
                                                    var v117 uint16 = v37
                                                    if ((!true) && ((v61 < 1e10) || (((42 <= 0) || (v60 != v60)) || (v61 >= (float64((v33 * v33)) - ((float64(float32(v61)) - v61) / 1.0)))))) {
                                                        if len(v20) >= 2 {
                                                            v20 = v20[1:len(v20)]
                                                        }
                                                        v118 := GenS12{F0: 0}
                                                        v47 += fi59
                                                        _ = v118
                                                    } else {
                                                        if false {
                                                            var v119 int8 = 32
                                                            _ = v119
                                                        } else {
                                                            switch fi59 {
                                                            case -574:
                                                                v120 := v3["a"]
                                                                _ = v120
                                                            case 1:
                                                                v34 = v34
                                                            }
                                                            v121 := GenS13{F0: 32, F1: 42, F2: 1000}
                                                            fmt.Println(v47, v34, v117)
                                                            v122 := GenS14{F0: v115}
                                                            _ = v121
                                                            _ = v122
                                                        }
                                                        var v123 uint8 = 42
                                                        delete(v3, v36)
                                                        _ = v123
                                                    }
                                                    _ = v115
                                                    _ = v116
                                                    _ = v117
                                                }
                                            case "hello":
                                                v20 = append(v20, (fi56 & 1))
                                            case "foo bar":
                                                fmt.Println(v35, v37)
                                            }
                                        }
                                    default:
                                        var v124 bool = false
                                        _ = v124
                                    }
                                    var v125 uint = 0
                                    v61 = (float64(float32(1.5)) - v61)
                                    _ = v125
                                }
                            }
                            _ = fi59
                            _ = v60
                            _ = v61
                        }
                    }
                default:
                    for fi126 := 0; fi126 < 4; fi126++ {
                        if (("" + v36) < v36) {
                            if ((fi56 + (v47 << ((972) & 63))) != (v47 * 128)) {
                                var v127 uint8 = (v34 >> ((214) & 63))
                                var v128 float32 = ((0.1 * float32(((0.0 - 0.0) * (((0.1 / float64(float32(-2.25))) / 3.14159) / 1.0)))) * v33)
                                if ((1 >= -1) || (!(true && true))) {
                                    var v129 float64 = 0.1
                                    v37 = v37
                                    var v130 uint32 = 645
                                    _ = v129
                                    _ = v130
                                } else {
                                    v127 = 255
                                }
                                _ = v127
                                _ = v128
                            }
                        } else {
                            v20 = append(v20, fi56)
                            if ((v33 * v33) < (-2.25 * 1.5)) {
                                if len(v20) >= 2 {
                                    v20 = v20[1:len(v20)]
                                }
                            }
                        }
                        var v131 bool = ((true && true) || false)
                        if true {
                            fmt.Println(v36)
                            fmt.Println(v37, v34)
                            var v132 int32 = v2
                            v133 := GenS15{F0: ("\n\t" + "hello")}
                            _ = v132
                            _ = v133
                        } else {
                            v2 += (v2 << ((42) & 63))
                            if ((v47 == 9223372036854775807) || v131) {
                                if (v131 && ((-(fi56)) <= fi56)) {
                                    var v134 int8 = -128
                                    _ = v134
                                }
                                _ = len(v20)
                                v135 := GenS16{F0: 2147483648, F1: 0, F2: 528}
                                _ = v135
                            }
                            var v136 uint = 1
                            _ = v136
                        }
                        _ = fi126
                        _ = v131
                    }
                }
                delete(v3, v36)
                fmt.Println(v47, v35, v36)
                _ = fi56
            }
        }
        _ = v47
    }
    if false {
        v37 += v37
        switch v36 {
        case "hello":
            var v137 uint32 = 2147483648
            _ = v137
        default:
            var v138 uint8 = (v34 / (((v34 * 128)) | 1))
            _ = v138
        }
    } else {
        v3[v36] = 0
        fmt.Println(v35, v2, v36)
    }
    v34 = 56
    if (((128 >= -971) || (128 != -54)) && (((true && true) || (true || true)) && (!true))) {
        if true {
            if false {
                v37 += 255
                var v139 int16 = -1
                _ = v139
            } else {
                for fi140 := 0; fi140 < 12; fi140++ {
                    var v141 int64 = 129
                    v35 += v2
                    v35 = genFn1(v36, 1, 4294967295)
                    _ = fi140
                    _ = v141
                }
            }
            v142 := GenS1{F0: 42, F1: v37}
            _ = v142
        }
        v36 = ""
        fmt.Println(v36)
    } else {
        v36 = ("" + v36)
        switch ((v36 + v36) + ((v36 + ("hello" + "a")) + v36)) {
        case "a":
            fmt.Println(v34)
        default:
            v143 := GenS5{F0: (v37 % ((2) | 1)), F1: 2}
            _ = v143
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
