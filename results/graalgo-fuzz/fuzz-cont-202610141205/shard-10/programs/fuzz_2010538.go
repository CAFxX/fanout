package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 bool; F1 int8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 string; F1 string; F2 uint16 }
type GenS10 struct { F0 uint64; F1 int64 }
type GenS11 struct { F0 int16; F1 uint16 }
type GenS12 struct { F0 uint32; F1 int32 }
type GenS13 struct { F0 bool; F1 int16; F2 bool }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 uint16; F1 int16 }
type GenS16 struct { F0 int64; F1 int32 }
type GenS17 struct { F0 int16 }
type GenS18 struct { F0 uint32 }
type GenS19 struct { F0 uint; F1 string; F2 int64 }
type GenS20 struct { F0 uint; F1 int }
type GenS21 struct { F0 uint }
type GenS22 struct { F0 bool }
type GenS23 struct { F0 int16; F1 int }
type GenS24 struct { F0 int32; F1 int16 }
type GenS25 struct { F0 int8; F1 string; F2 int16 }
type GenS26 struct { F0 uint16; F1 int8; F2 int32 }
func genFn0(p0 int16, p1 uint64, p2 uint) uint32 {
    if (!true) {
        for fi6 := 0; fi6 < 15; fi6++ {
            v7 := GenS1{F0: (p0 * 42), F1: true, F2: 65535}
            var v8 uint = (p2 >> ((p2) & 63))
            _ = fi6
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := []int{(fi9 / ((fi9) | 1)), 889}
            for fi11 := 0; fi11 < 13; fi11++ {
                _ = len(v10)
                switch "hello" {
                case "\n\t":
                    var v12 int32 = 42
                    _ = v12
                case "a":
                    var v13 int = fi9
                    _ = v13
                default:
                    p1 = (^(p1))
                }
                _ = fi11
            }
            for fi14 := 0; fi14 < 2; fi14++ {
                var v15 string = ""
                if len(v10) > 0 {
                    v10[0] = fi9
                }
                _ = fi14
                _ = v15
            }
            _ = fi9
            _ = v10
        }
        var v16 uint32 = 65535
        _ = v16
    } else {
        p2 += (^(p2))
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        switch -93 {
        case 0:
            p2 = p2
        case 2147483647:
            if false {
                v18 := []int{fi17, 9223372036854775807, 595, fi17}
                for fi19 := 0; fi19 < 7; fi19++ {
                    p1 = (^(p1))
                    switch ("a" + "a") {
                    case "\n\t":
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                    case "":
                        v20 := GenS2{F0: 1}
                        _ = v20
                    case "a":
                        fmt.Println(fi17)
                    }
                    var v21 uint64 = p1
                    _ = fi19
                    _ = v21
                }
                _ = v18
            } else {
                var v22 float32 = -2.25
                v23 := GenS3{F0: -1}
                if (4294967295 >= 42) {
                    v24 := GenS3{F0: 0}
                    v25 := GenS2{F0: 127}
                    _ = v24
                    _ = v25
                }
                _ = v22
                _ = v23
            }
        case 1:
            if true {
                var v26 uint = 18446744073709551615
                p0 = p0
                _ = v26
            }
        default:
            for fi27 := 0; fi27 < 1; fi27++ {
                v28 := GenS2{F0: 1}
                p1 += 1
                _ = fi27
                _ = v28
            }
        }
        p2 -= p2
        _ = fi17
    }
    for fi29 := 0; fi29 < 12; fi29++ {
        if (false || ((-589 >= -1) && true)) {
            for fi30 := 0; fi30 < 11; fi30++ {
                switch ("foo bar" + "a") {
                case "a":
                    if true {
                        v31 := []int{fi29, fi29}
                        var v32 uint = 39
                        if (p2 != 125) {
                            if (!true) {
                                v33 := GenS1{F0: p0, F1: (412 > 2147483647), F2: 2147483648}
                                p0 += p0
                                _ = v33
                            }
                            var v34 uint8 = 1
                            _ = v34
                        }
                        _ = v31
                        _ = v32
                    } else {
                        p2 = p2
                        v35 := []int{(-(fi29))}
                        v36 := GenS3{F0: -1}
                        _ = v35
                        _ = v36
                    }
                default:
                    if true {
                        p2 = p2
                        switch "a" {
                        case "a":
                            switch "a" {
                            case "\n\t":
                                v37 := GenS1{F0: 945, F1: false, F2: 2}
                                _ = v37
                            case "":
                                p2 = 0
                            case "a":
                                p2 = (p2 / ((0) | 1))
                            }
                        case "":
                            fmt.Println(fi29)
                        case "foo bar":
                            v38 := []int{fi29, (fi29 | 752), 32767}
                            _ = v38
                        default:
                            p2 -= (p2 % ((p2) | 1))
                        }
                    } else {
                        v39 := GenS2{F0: -2}
                        v40 := []int{fi29, -2147483648}
                        var v41 string = "foo bar"
                        var v42 int16 = p0
                        if len(v40) >= 2 {
                            v40 = v40[1:len(v40)]
                        }
                        _ = v39
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                }
                v43 := []int{fi30}
                _ = fi30
                _ = v43
            }
            switch 0 {
            case 2147483647:
                var v44 uint = 0
                _ = v44
            case 399:
                switch "foo bar" {
                case "hello":
                    p1 -= p1
                case "\n\t":
                    p0 -= int16(p1)
                case "foo bar":
                    v45 := GenS3{F0: 100}
                    _ = v45
                default:
                    p1 = 250
                }
            }
            p2 -= p2
        }
        fmt.Println(p0, p1, p0)
        _ = fi29
    }
    var v46 uint = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v46
    return 15
}
func genFn1(p0 int, p1 int32) uint64 {
    fmt.Println(p1)
    var v47 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = v47
    return 2
}

func main() {
    v1 := []string{"\n\t"}
    v2 := []uint32{}
    v3 := []uint32{115, 993, 2, 65535}
    var v4 uint64 = 2
    var v5 uint = 65535
    var v48 uint8 = 42
    var v49 uint32 = 2147483647
    var v50 uint16 = 2
    for fi51 := 0; fi51 < 5; fi51++ {
        for fi52 := 0; fi52 < 6; fi52++ {
            v53 := GenS4{F0: (v4 + genFn1((^(fi51)), -1)), F1: uint8(v50), F2: 702}
            var v54 int = (fi52 + fi51)
            _ = fi52
            _ = v53
            _ = v54
        }
        var v55 uint = 244
        switch "\n\t" {
        case "a":
            var v56 uint16 = (v50 * (v50 - 1))
            _ = v56
        case "foo bar":
            if len(v3) > 0 {
                v3[0] = genFn0(-278, 1, (v55 >> ((v5) & 63)))
            }
        }
        _ = fi51
        _ = v55
    }
    if ((894 <= 1) && true) {
        v5 = (v5 - (v5 & 255))
        v4 = 4294967295
        v4 = uint64(v48)
    }
    if false {
        var v57 int64 = -1
        _ = v57
    } else {
        for fi58 := 0; fi58 < 19; fi58++ {
            for fi59 := 0; fi59 < 16; fi59++ {
                fmt.Println(fi59, v50, v50)
                switch fi58 {
                case -431:
                    var v60 int64 = 42
                    _ = v60
                case 1:
                    fmt.Println(v49, v49)
                case -2147483648:
                    var v61 float32 = (((0.0 + (0.1 / -2.25)) * float32(1.5)) * 1e10)
                    _ = v61
                }
                if ((^(v49)) >= v49) {
                    if (fi59 >= 32767) {
                        if false {
                            v62 := GenS1{F0: 42, F1: true, F2: (v49 ^ genFn0(42, (v4 >> ((v5) & 63)), v5))}
                            v5 = uint(v50)
                            switch fi58 {
                            case 1:
                                var v63 uint32 = genFn0(42, genFn1((^(fi58)), 359), (v5 - v5))
                                _ = v63
                            case -20:
                                var v64 int64 = 9223372036854775807
                                _ = v64
                            }
                            v5 += (v5 * v5)
                            v49 = genFn0(-1, genFn1(fi58, -13), 152)
                            _ = v62
                        }
                        fmt.Println(v5, fi58, fi58)
                        fmt.Println(v49, fi59, v49)
                    } else {
                        if true {
                            fmt.Println(fi58, v4)
                            if false {
                                var v65 uint64 = (v4 * (v4 / (((^(v4))) | 1)))
                                if true {
                                    switch int(v5) {
                                    case 702:
                                        if true {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            var v66 int64 = 9223372036854775807
                                            var v67 uint = (v5 >> ((v5) & 63))
                                            _ = v66
                                            _ = v67
                                        } else {
                                            if (true && true) {
                                                switch 128 {
                                                case 13:
                                                    fmt.Println(fi59, v5)
                                                case -9223372036854775808:
                                                    fmt.Println(fi59)
                                                case 225:
                                                    switch "hello" {
                                                    case "foo bar":
                                                        v5 -= (^(v5))
                                                    case "hello":
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            v50 += 42
                                                        case "hello":
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        case "a":
                                                            if ((false && false) && (!((0.0 <= 3.14159) && false))) {
                                                                v48 -= (v48 | (v48 >> ((v5) & 63)))
                                                            } else {
                                                                switch fi58 {
                                                                case 1:
                                                                    v68 := GenS5{F0: 1, F1: (v65 << ((4294967295) & 63))}
                                                                    _ = v68
                                                                case 127:
                                                                    if true {
                                                                        switch 1 {
                                                                        case 768:
                                                                            v69 := GenS4{F0: v65, F1: v48, F2: -1}
                                                                            v49 += (v49 - v49)
                                                                            _ = v69
                                                                        case 42:
                                                                            fmt.Println(fi59, fi58)
                                                                        default:
                                                                            switch (fi58 & 42) {
                                                                            case -775:
                                                                                if (42 <= int32(v5)) {
                                                                                    switch fi59 {
                                                                                    case 128:
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = ("a" + ("a" + "foo bar"))
                                                                                        }
                                                                                    case 127:
                                                                                        if false {
                                                                                            v5 = v5
                                                                                        } else {
                                                                                            v4 = 0
                                                                                            v5 += v5
                                                                                            var v70 float32 = 0.0
                                                                                            v70 += v70
                                                                                            _ = v70
                                                                                        }
                                                                                    case 32767:
                                                                                        fmt.Println(v50, v65, fi59)
                                                                                    }
                                                                                    var v71 uint64 = v4
                                                                                    _ = v71
                                                                                } else {
                                                                                    v5 = uint(v50)
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = (v49 << ((uint(v49)) & 63))
                                                                                    }
                                                                                }
                                                                            default:
                                                                                v5 -= 42
                                                                            }
                                                                        }
                                                                        if len(v3) > 0 {
                                                                            v3[0] = uint32(v50)
                                                                        }
                                                                        if true {
                                                                            v2 = append(v2, uint32(v50))
                                                                            v50 = uint16(v5)
                                                                        }
                                                                    } else {
                                                                        var v72 string = ("" + "foo bar")
                                                                        _ = v72
                                                                    }
                                                                    _ = len(v3)
                                                                case 42:
                                                                    var v73 uint8 = (v48 + 0)
                                                                    _ = v73
                                                                default:
                                                                    v65 = genFn1(fi58, 1)
                                                                }
                                                                switch (fi58 - int(v48)) {
                                                                case 9223372036854775807:
                                                                    var v74 string = "hello"
                                                                    _ = v74
                                                                default:
                                                                    v75 := GenS6{F0: false, F1: 1}
                                                                    _ = v75
                                                                }
                                                                if false {
                                                                    v76 := GenS3{F0: 111}
                                                                    switch fi59 {
                                                                    case 383:
                                                                        var v77 uint64 = v65
                                                                        _ = v77
                                                                    default:
                                                                        if false {
                                                                            if (true && false) {
                                                                                if true {
                                                                                    switch 0 {
                                                                                    case 144:
                                                                                        _ = len(v2)
                                                                                        v78 := GenS7{F0: -409}
                                                                                        v79 := GenS8{F0: ("" + ("foo bar" + "a"))}
                                                                                        _ = v78
                                                                                        _ = v79
                                                                                    case 32768:
                                                                                        v4 -= uint64(v49)
                                                                                    case -9223372036854775808:
                                                                                        var v80 int32 = -1
                                                                                        _ = v80
                                                                                    }
                                                                                    var v81 int16 = 1000
                                                                                    _ = v81
                                                                                } else {
                                                                                    v50 = 65535
                                                                                    var v82 bool = true
                                                                                    _ = v82
                                                                                }
                                                                            } else {
                                                                                v83 := GenS4{F0: v4, F1: (^(v48)), F2: 9223372036854775807}
                                                                                _ = len(v1)
                                                                                _ = v83
                                                                            }
                                                                            var v84 int32 = -2147483648
                                                                            _ = v84
                                                                        }
                                                                    }
                                                                    switch ("a" + ("hello" + "\n\t")) {
                                                                    case "a":
                                                                        v48 = (v48 << (((^(v5))) & 63))
                                                                    case "":
                                                                        var v85 int64 = 0
                                                                        _ = v85
                                                                    case "hello":
                                                                        var v86 uint8 = uint8(v5)
                                                                        _ = v86
                                                                    default:
                                                                        v65 = v65
                                                                        if ((0.0 - float64(1.5)) > (0.1 - ((3.14159 * ((float64((0.0 * -2.25)) - 0.1) - ((0.0 * -2.25) / 1.0))) + 0.0))) {
                                                                            fmt.Println(v50)
                                                                            var v87 uint8 = v48
                                                                            _ = v87
                                                                        } else {
                                                                            var v88 int8 = -128
                                                                            var v89 int16 = int16(v88)
                                                                            if (!((0.0 > float64((float32(1e10) / (3.14159 / -2.25)))) && false)) {
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = (v49 % (((v49 * uint32(v48))) | 1))
                                                                                }
                                                                                v4 += genFn1(fi58, 0)
                                                                                v49 = (v49 - genFn0(32767, (v65 >> ((1) & 63)), v5))
                                                                                var v90 uint64 = uint64(v5)
                                                                                _ = v90
                                                                            }
                                                                            _ = v88
                                                                            _ = v89
                                                                        }
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            var v91 uint = v5
                                                                            _ = v91
                                                                        default:
                                                                            var v92 string = "\n\t"
                                                                            _ = v92
                                                                        }
                                                                    }
                                                                    _ = v76
                                                                }
                                                            }
                                                        default:
                                                            if true {
                                                                if (!((v50 ^ v50) < (v50 - (^(v50))))) {
                                                                    v93 := GenS6{F0: false, F1: -1}
                                                                    v2 = append(v2, v49)
                                                                    v5 -= (v5 >> ((v5) & 63))
                                                                    _ = v93
                                                                } else {
                                                                    fmt.Println(fi59, v48, v50)
                                                                    if ((v5 >> (((v5 + uint(v49))) & 63)) >= v5) {
                                                                        switch "a" {
                                                                        case "":
                                                                            v94 := GenS9{F0: ("foo bar" + "hello"), F1: ("" + ""), F2: (v50 + (v50 << ((v5) & 63)))}
                                                                            _ = v94
                                                                        case "hello":
                                                                            v95 := GenS10{F0: v4, F1: int64(fi58)}
                                                                            _ = v95
                                                                        case "\n\t":
                                                                            if false {
                                                                                _ = len(v2)
                                                                                var v96 int64 = -1
                                                                                v96 -= v96
                                                                                fmt.Println(v4)
                                                                                _ = v96
                                                                            } else {
                                                                                v65 += v4
                                                                                v65 = (^(v4))
                                                                                if ((("foo bar" + (((("a" + "foo bar") + ("" + ("a" + "hello"))) + "hello") + "a")) + "foo bar") > "foo bar") {
                                                                                    var v97 uint8 = v48
                                                                                    _ = v97
                                                                                }
                                                                                if false {
                                                                                    v98 := GenS11{F0: 1000, F1: v50}
                                                                                    v50 = (v50 + uint16(v48))
                                                                                    _ = v98
                                                                                }
                                                                            }
                                                                        default:
                                                                            fmt.Println(fi58, v65, v50)
                                                                        }
                                                                        fmt.Println(fi59, v5)
                                                                    } else {
                                                                        v99 := GenS10{F0: (v4 - 631), F1: 0}
                                                                        var v100 float32 = 1e10
                                                                        v4 -= v65
                                                                        _ = v100
                                                                        _ = v99
                                                                    }
                                                                    var v101 uint8 = v48
                                                                    _ = v101
                                                                }
                                                                var v102 uint16 = (v50 << (((v5 << ((845) & 63))) & 63))
                                                                var v103 int = (fi59 / ((fi59) | 1))
                                                                v104 := GenS12{F0: 65535, F1: 2147483647}
                                                                _ = v102
                                                                _ = v103
                                                                _ = v104
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                            var v105 int64 = 0
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = v105
                                        }
                                    case -2147483648:
                                        if true {
                                            if (((0.1 / float64(0.1)) == (-2.25 + 1e10)) || (false || false)) {
                                                v65 -= v4
                                                v65 = v4
                                                var v106 int64 = 58
                                                _ = v106
                                            }
                                            v4 -= (v65 % (((^(v65))) | 1))
                                            _ = len(v1)
                                            v4 += 1
                                        } else {
                                            var v107 float32 = 3.14159
                                            fmt.Println(v5)
                                            _ = v107
                                        }
                                    default:
                                        v49 += 2147483647
                                        _ = len(v3)
                                    }
                                    var v108 string = ""
                                    v108 = "a"
                                    _ = v108
                                }
                                v109 := GenS3{F0: -128}
                                _ = v109
                                _ = v65
                            } else {
                                v110 := GenS11{F0: 588, F1: (v50 ^ v50)}
                                if false {
                                    var v111 int32 = 2147483647
                                    _ = v111
                                } else {
                                    v2 = append(v2, genFn0(-1, v4, uint(v50)))
                                    _ = len(v1)
                                    var v112 int8 = -128
                                    _ = v112
                                }
                                var v113 float64 = ((((0.0 + float64(1.5)) + (float64(1e10) - float64(1.5))) - (0.0 - -2.25)) / (1.5 - (0.0 + (3.14159 + (3.14159 - 1e10)))))
                                if false {
                                    v48 += (^(v48))
                                    fmt.Println(fi59, v50)
                                    if (!true) {
                                        switch "\n\t" {
                                        case "foo bar":
                                            v113 += v113
                                        case "hello":
                                            fmt.Println(v50)
                                            v114 := GenS8{F0: "\n\t"}
                                            switch "a" {
                                            case "\n\t":
                                                switch "" {
                                                case "\n\t":
                                                    v5 = uint(v48)
                                                }
                                            case "":
                                                if false {
                                                    v49 += 0
                                                    v4 = (v4 - uint64(v48))
                                                    v115 := GenS13{F0: true, F1: -1, F2: (-812 <= 42)}
                                                    if false {
                                                        var v116 uint32 = (v49 / ((42) | 1))
                                                        var v117 uint16 = v50
                                                        fmt.Println(v48, v5)
                                                        _ = v116
                                                        _ = v117
                                                    } else {
                                                        switch "\n\t" {
                                                        case "":
                                                            if len(v1) > 0 {
                                                                v1[0] = "a"
                                                            }
                                                        default:
                                                            switch (fi59 * (fi59 ^ (fi59 - int(v49)))) {
                                                            case -176:
                                                                v118 := GenS14{F0: genFn0(1, 697, v5)}
                                                                _ = v118
                                                            case 9223372036854775807:
                                                                _ = len(v2)
                                                            }
                                                        }
                                                        var v119 float64 = 0.1
                                                        if false {
                                                            v50 = v50
                                                            if (false || true) {
                                                                v120 := GenS15{F0: (v50 ^ v50), F1: 564}
                                                                var v121 bool = true
                                                                _ = v120
                                                                _ = v121
                                                            }
                                                            v122 := GenS16{F0: 1, F1: -838}
                                                            _ = v122
                                                        }
                                                        _ = v119
                                                    }
                                                    var v123 bool = (true && (0.0 != 3.14159))
                                                    _ = v115
                                                    _ = v123
                                                }
                                            case "hello":
                                                v4 -= v4
                                            }
                                            _ = v114
                                        default:
                                            fmt.Println(v48, v50)
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v124 int16 = 0
                                        var v125 uint = (v5 * 0)
                                        _ = v124
                                        _ = v125
                                    }
                                    fmt.Println(fi58)
                                    v126 := GenS17{F0: 32767}
                                    _ = v126
                                }
                                var v127 uint8 = 143
                                if false {
                                    var v128 string = "a"
                                    _ = v128
                                }
                                switch (-(fi58)) {
                                case 9223372036854775807:
                                    switch "\n\t" {
                                    case "foo bar":
                                        if (!false) {
                                            switch (fi59 | -2147483648) {
                                            case 32767:
                                                var v129 uint64 = genFn1((^(fi58)), 42)
                                                _ = v129
                                            case 954:
                                                var v130 int32 = 2147483647
                                                _ = v130
                                            default:
                                                var v131 uint32 = 2
                                                _ = v131
                                            }
                                            var v132 float64 = (1.5 * v113)
                                            switch "hello" {
                                            case "\n\t":
                                                if len(v3) > 0 {
                                                    v3[0] = genFn0(-1, (v4 ^ (^(v4))), 234)
                                                }
                                                v48 = v48
                                                if false {
                                                    v49 += (^(v49))
                                                    fmt.Println(v50, v132, v113)
                                                }
                                            default:
                                                fmt.Println(v4)
                                            }
                                            _ = v132
                                        } else {
                                            if (((0.1 + (1e10 * 0.1)) == (1.5 * 1.5)) || (25 < (v5 * v5))) {
                                                if (0.1 >= -2.25) {
                                                    v133 := GenS2{F0: 42}
                                                    _ = v133
                                                } else {
                                                    var v134 int8 = 9
                                                    var v135 bool = false
                                                    switch (fi58 | 0) {
                                                    case -1:
                                                        if len(v2) > 0 {
                                                            v2[0] = uint32(v4)
                                                        }
                                                        var v136 uint = 255
                                                        _ = v136
                                                    case 1:
                                                        v137 := GenS10{F0: uint64(v5), F1: -9223372036854775808}
                                                        _ = v137
                                                    case 127:
                                                        v138 := GenS18{F0: v49}
                                                        _ = v138
                                                    }
                                                    _ = v134
                                                    _ = v135
                                                }
                                                _ = len(v3)
                                            }
                                            var v139 int8 = 1
                                            _ = v139
                                        }
                                    case "\n\t":
                                        v127 += v127
                                    case "":
                                        var v140 uint8 = v127
                                        _ = v140
                                    }
                                case -2147483648:
                                    var v141 uint32 = 1
                                    _ = v141
                                default:
                                    v3 = append(v3, uint32(v4))
                                }
                                _ = v110
                                _ = v113
                                _ = v127
                            }
                        }
                        var v142 uint8 = (v48 * v48)
                        switch "foo bar" {
                        case "a":
                            v2 = append(v2, v49)
                        }
                        _ = v142
                    }
                } else {
                    v5 = (v5 - v5)
                    switch ("a" + ("foo bar" + "hello")) {
                    case "foo bar":
                        if ((v48 * v48) < uint8(v49)) {
                            var v143 int32 = 42
                            if true {
                                v49 = v49
                                _ = len(v1)
                            }
                            if len(v1) > 0 {
                                v1[0] = "foo bar"
                            }
                            _ = v143
                        }
                    }
                    switch "hello" {
                    case "foo bar":
                        v144 := GenS18{F0: (v49 * genFn0(1, v4, 4294967295))}
                        var v145 int16 = -32768
                        _ = v144
                        _ = v145
                    case "":
                        _ = len(v3)
                    }
                }
                _ = fi59
            }
            _ = fi58
        }
        fmt.Println(v4)
        var v146 int16 = -32768
        _ = v146
    }
    for fi147 := 0; fi147 < 3; fi147++ {
        fmt.Println(fi147, v50, v4)
        if false {
            var v148 uint16 = v50
            _ = v148
        } else {
            for fi149 := 0; fi149 < 17; fi149++ {
                if (575 > (v50 / ((281) | 1))) {
                    switch (((((("a" + "foo bar") + ("foo bar" + "hello")) + "\n\t") + "foo bar") + "a") + ("" + "a")) {
                    case "":
                        v5 -= v5
                    case "foo bar":
                        v150 := GenS19{F0: 1, F1: "a", F2: -1}
                        _ = v150
                    default:
                        if (v50 > (v50 * 0)) {
                            var v151 float64 = (-2.25 * 3.14159)
                            var v152 bool = (82 == (v48 / (((v48 << (((v5 % ((uint(v50)) | 1))) & 63))) | 1)))
                            _ = v151
                            _ = v152
                        }
                    }
                    var v153 int32 = 2147483647
                    var v154 uint16 = uint16(v48)
                    _ = len(v3)
                    v2 = append(v2, (v49 & v49))
                    if (1.5 == -2.25) {
                        fmt.Println(v5, v154, fi149)
                        v155 := GenS8{F0: ""}
                        switch (-(fi149)) {
                        case 4294967295:
                            v49 -= v49
                        case 127:
                            switch ("hello" + (("hello" + "") + "a")) {
                            case "\n\t":
                                if true {
                                    switch fi149 {
                                    case 32768:
                                        v156 := GenS20{F0: (v5 + v5), F1: (fi147 + fi147)}
                                        _ = v156
                                    case 220:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    case 655:
                                        v157 := GenS8{F0: ("foo bar" + "a")}
                                        _ = v157
                                    }
                                    switch fi147 {
                                    case 98:
                                        switch fi147 {
                                        case -2147483648:
                                            v4 -= genFn1(fi149, v153)
                                        case 1:
                                            v154 -= (v154 + 0)
                                        }
                                    case -150:
                                        v5 = 7
                                    default:
                                        fmt.Println(v48, fi149, v50)
                                    }
                                } else {
                                    if (65535 >= 2) {
                                        if true {
                                            fmt.Println(v50, v5, v153)
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        } else {
                                            v158 := GenS2{F0: 0}
                                            if true {
                                                var v159 uint64 = v4
                                                if false {
                                                    if (false || false) {
                                                        v160 := GenS21{F0: 42}
                                                        v161 := GenS22{F0: (v49 >= (^(v49)))}
                                                        if true {
                                                            var v162 int32 = v153
                                                            var v163 int8 = -1
                                                            _ = v162
                                                            _ = v163
                                                        } else {
                                                            v50 = 32767
                                                            var v164 int64 = -1
                                                            fmt.Println(v49)
                                                            var v165 uint32 = (v49 >> ((v5) & 63))
                                                            var v166 uint32 = v49
                                                            _ = v164
                                                            _ = v165
                                                            _ = v166
                                                        }
                                                        _ = v160
                                                        _ = v161
                                                    }
                                                    var v167 uint64 = uint64(v5)
                                                    _ = v167
                                                } else {
                                                    v50 = 433
                                                    switch ("hello" + "") {
                                                    case "":
                                                        var v168 int16 = 42
                                                        _ = v168
                                                    default:
                                                        var v169 float32 = 1.5
                                                        _ = v169
                                                    }
                                                    if (0 != -9223372036854775808) {
                                                        var v170 bool = true
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            var v171 uint16 = v50
                                                            _ = v171
                                                        default:
                                                            switch "hello" {
                                                            case "\n\t":
                                                                _ = len(v1)
                                                            case "a":
                                                                var v172 uint8 = (v48 * (v48 & (v48 >> (((v5 >> ((v5) & 63))) & 63))))
                                                                fmt.Println(v48, v50, v159)
                                                                _ = v172
                                                            case "":
                                                                var v173 bool = false
                                                                _ = v173
                                                            default:
                                                                v174 := GenS3{F0: 1}
                                                                _ = v174
                                                            }
                                                        }
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        _ = v170
                                                    }
                                                }
                                                v48 = uint8(v5)
                                                _ = v159
                                            } else {
                                                var v175 int64 = 25
                                                _ = len(v2)
                                                v50 = 1
                                                fmt.Println(v154, fi147)
                                                _ = v175
                                            }
                                            _ = v158
                                        }
                                        _ = len(v1)
                                        if (2 != uint(v49)) {
                                            var v176 uint32 = 65535
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = v176
                                        }
                                        var v177 uint16 = 32767
                                        _ = v177
                                    } else {
                                        if false {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = (v49 | 1)
                                            }
                                            var v178 int8 = -31
                                            _ = v178
                                        }
                                        var v179 uint8 = v48
                                        v153 -= v153
                                        _ = v179
                                    }
                                    v49 = uint32(v5)
                                    if true {
                                        if (2 <= (v50 + v50)) {
                                            v48 += v48
                                            fmt.Println(v154, v5)
                                        }
                                        if true {
                                            v48 = v48
                                        }
                                        fmt.Println(v48, v153, v153)
                                        v4 = (v4 / ((genFn1(fi147, 2147483647)) | 1))
                                        v153 -= v153
                                    } else {
                                        v180 := GenS16{F0: 0, F1: (v153 + 42)}
                                        _ = v180
                                    }
                                }
                            case "a":
                                v48 -= (v48 << (((^(v5))) & 63))
                            }
                        case 32768:
                            v181 := GenS23{F0: 0, F1: fi147}
                            _ = v181
                        }
                        _ = v155
                    } else {
                        var v182 string = "\n\t"
                        _ = v182
                    }
                    v183 := GenS15{F0: (v154 << (((v5 << (((v5 / ((255) | 1))) & 63))) & 63)), F1: 42}
                    var v184 int16 = 496
                    _ = v153
                    _ = v154
                    _ = v183
                    _ = v184
                }
                if false {
                    if (false || true) {
                        var v185 uint32 = 2147483648
                        v5 = (v5 | (v5 | (v5 >> ((475) & 63))))
                        var v186 int = (-(fi149))
                        _ = v185
                        _ = v186
                    } else {
                        if ((1 != -195) && true) {
                            v187 := GenS24{F0: 42, F1: 355}
                            v188 := GenS15{F0: v50, F1: 1}
                            _ = v187
                            _ = v188
                        }
                        var v189 float32 = (0.1 * (0.0 + -2.25))
                        if false {
                            if (!true) {
                                if true {
                                    if true {
                                        _ = len(v2)
                                        var v190 uint64 = genFn1(fi147, int32(v50))
                                        _ = v190
                                    } else {
                                        if (0 == (v48 / (((v48 | (^(v48)))) | 1))) {
                                            switch 127 {
                                            case 704:
                                                var v191 int = fi149
                                                _ = v191
                                            }
                                        }
                                    }
                                }
                                v192 := GenS25{F0: 1, F1: "hello", F2: 0}
                                var v193 string = ""
                                v194 := GenS6{F0: false, F1: 0}
                                v4 = 0
                                _ = v192
                                _ = v193
                                _ = v194
                            }
                            if false {
                                var v195 bool = true
                                if (!(3.14159 == v189)) {
                                    if (v195 && v195) {
                                        var v196 string = "a"
                                        if (42 < 42) {
                                            v197 := GenS2{F0: int8(v49)}
                                            var v198 float64 = -2.25
                                            if (1 != 2147483647) {
                                                if ("hello" != v196) {
                                                    var v199 uint32 = v49
                                                    var v200 float64 = float64(1e10)
                                                    if true {
                                                        v201 := GenS14{F0: (^(v49))}
                                                        _ = v201
                                                    }
                                                    _ = v199
                                                    _ = v200
                                                } else {
                                                    v202 := GenS21{F0: (v5 ^ (v5 | v5))}
                                                    var v203 uint = v5
                                                    v203 = (v203 & v5)
                                                    _ = v202
                                                    _ = v203
                                                }
                                            }
                                            var v204 uint = uint(v49)
                                            _ = v197
                                            _ = v198
                                            _ = v204
                                        }
                                        _ = v196
                                    } else {
                                        fmt.Println(v50, v49, fi147)
                                        _ = len(v3)
                                    }
                                }
                                _ = v195
                            }
                            if (42 != 9223372036854775807) {
                                v4 += (^(v4))
                                _ = len(v1)
                                v1 = append(v1, "hello")
                            }
                        } else {
                            var v205 float32 = ((0.0 - float32(1e10)) + (v189 / v189))
                            if (-1 > 0) {
                                var v206 int64 = -469
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v206
                            } else {
                                v4 = v4
                            }
                            var v207 string = (((("" + ("foo bar" + "")) + "a") + "foo bar") + "foo bar")
                            v208 := GenS1{F0: 1000, F1: (!((((true && true) && false) && false) || false)), F2: (v49 ^ v49)}
                            var v209 bool = false
                            _ = v205
                            _ = v207
                            _ = v208
                            _ = v209
                        }
                        _ = v189
                    }
                } else {
                    v49 = (v49 << (((v5 >> ((v5) & 63))) & 63))
                    v5 = (^(v5))
                    v49 = 2147483647
                }
                _ = fi149
            }
            var v210 float64 = 1e10
            var v211 float64 = (0.0 * float64(float32(((v210 / -2.25) / (v210 + 0.1)))))
            _ = v210
            _ = v211
        }
        _ = fi147
    }
    v212 := GenS26{F0: (v50 - (v50 - v50)), F1: 127, F2: 0}
    v48 = (v48 ^ v48)
    var v213 uint64 = (v4 / (((v4 / ((v4) | 1))) | 1))
    for fi214 := 0; fi214 < 10; fi214++ {
        v49 = 773
        _ = fi214
    }
    switch "a" {
    case "hello":
        for fi215 := 0; fi215 < 7; fi215++ {
            switch 32768 {
            case -1:
                var v216 int8 = 114
                _ = v216
            case 379:
                v4 += (v213 | uint64(v49))
            case 42:
                fmt.Println(v5, v213)
            }
            v213 = v4
            _ = fi215
        }
    case "\n\t":
        fmt.Println(v4, v5, v48)
    case "":
        v3 = append(v3, (^(v49)))
    }
    v3 = append(v3, 1)
    if (-26 < 42) {
        v4 -= v213
    } else {
        switch -467 {
        case 4294967295:
            v217 := GenS25{F0: 127, F1: (("foo bar" + ((("" + "foo bar") + ("hello" + "\n\t")) + "")) + ("\n\t" + "")), F2: -1}
            _ = v217
        case 931:
            if (v49 <= 2) {
                for fi218 := 0; fi218 < 6; fi218++ {
                    v219 := GenS14{F0: (v49 * genFn0(1000, v4, (v5 << (((v5 & 255)) & 63))))}
                    _ = fi218
                    _ = v219
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        case -1:
            fmt.Println(v50, v4)
        default:
            fmt.Println(v48)
        }
        var v220 float64 = 0.1
        _ = v220
    }
    switch "hello" {
    case "a":
        if (-765 > 2147483647) {
            var v221 uint16 = 217
            _ = v221
        }
    default:
        fmt.Println(v49)
    }
    v4 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v213:", v213)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v212
    _ = v213
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
}
