package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 int64; F1 bool; F2 uint32 }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS8 struct { F0 int; F1 string; F2 int16 }
type GenS9 struct { F0 string; F1 uint; F2 int32 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int32; F1 int; F2 int16 }
type GenS12 struct { F0 uint8; F1 uint8 }
type GenS13 struct { F0 int16; F1 int64; F2 uint8 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint16; F1 uint16 }
type GenS16 struct { F0 bool; F1 int8 }
type GenS17 struct { F0 string; F1 uint; F2 int32 }
func genFn0() int16 {
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 12; fi7++ {
            switch "hello" {
            case "":
                v8 := GenS1{F0: (0 < 654), F1: 993, F2: -271}
                _ = v8
            case "foo bar":
                v9 := GenS1{F0: true, F1: 715, F2: -9223372036854775808}
                _ = v9
            }
            _ = fi7
        }
        _ = fi6
    }
    var v10 int = 1
    _ = v10
    return -762
}

func main() {
    var v1 uint16 = 2
    var v2 int16 = -32768
    v3 := make(map[int]uint32)
    var v4 bool = (144 == 0)
    v5 := make(map[int]uint32)
    if v4 {
        v11 := GenS1{F0: ((0.1 * 3.14159) != 0.1), F1: 65535, F2: 42}
        for fi12 := 0; fi12 < 20; fi12++ {
            v2 = genFn0()
            switch "a" {
            case "":
                var v13 int8 = 0
                _ = v13
            case "\n\t":
                v14 := GenS2{F0: 283}
                _ = v14
            case "foo bar":
                for fi15 := 0; fi15 < 12; fi15++ {
                    v4 = (1e10 <= 1e10)
                    _ = fi15
                }
            }
            _ = fi12
        }
        _ = len(v5)
        _ = v11
    } else {
        var v16 int8 = 127
        _ = v16
    }
    v17 := GenS1{F0: v4, F1: 876, F2: 0}
    _ = len(v5)
    v4 = (v1 != (^(v1)))
    fmt.Println(v1)
    for fi18 := 0; fi18 < 2; fi18++ {
        switch ("foo bar" + "") {
        case "hello":
            for fi19 := 0; fi19 < 15; fi19++ {
                var v20 int8 = 100
                var v21 uint64 = 926
                v4 = (!v4)
                _ = fi19
                _ = v20
                _ = v21
            }
        case "":
            v1 = 42
        default:
            v22 := GenS2{F0: 2147483647}
            _ = v22
        }
        _ = fi18
    }
    if false {
        if (42 == 1) {
            v23 := v5[-495]
            for fi24 := 0; fi24 < 11; fi24++ {
                fmt.Println(v23, v4, v23)
                _ = fi24
            }
            _ = v23
        }
        fmt.Println(v2, v4)
    } else {
        v4 = ((true || v4) || ("" != ""))
        fmt.Println(v4, v2, v2)
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        switch ("" + "") {
        case "foo bar":
            v2 = v2
        case "hello":
            for fi26 := 0; fi26 < 19; fi26++ {
                if v4 {
                    fmt.Println(v4)
                    switch ("" + (("" + "") + "")) {
                    case "hello":
                        delete(v3, (-(fi26)))
                    case "\n\t":
                        v27 := GenS3{F0: 1, F1: -1}
                        _ = v27
                    case "a":
                        v2 += v2
                    default:
                        fmt.Println(v2)
                    }
                } else {
                    fmt.Println(v2, v4, fi26)
                    var v28 float64 = (float64(((1e10 * 1e10) * -2.25)) + (0.1 + (1.5 - 3.14159)))
                    _ = v28
                }
                if (243 < 1) {
                    var v29 uint8 = 1
                    _ = v29
                } else {
                    fmt.Println(fi25)
                    switch "a" {
                    case "":
                        var v30 int32 = -2147483648
                        _ = v30
                    case "foo bar":
                        v1 += v1
                    case "\n\t":
                        var v31 int = -2147483648
                        _ = v31
                    default:
                        fmt.Println(v1, fi25)
                    }
                    _ = len(v3)
                }
                v32 := GenS2{F0: 0}
                _ = fi26
                _ = v32
            }
        case "a":
            v33 := GenS2{F0: 2}
            _ = v33
        }
        _ = fi25
    }
    fmt.Println(v4, v2, v1)
    for fi34 := 0; fi34 < 3; fi34++ {
        v35 := []int{(fi34 | (^(fi34))), fi34, fi34, 127}
        v4 = ((1e10 - 0.0) == ((0.1 + (0.1 / 0.1)) + float32(float64(0.1))))
        _ = fi34
        _ = v35
    }
    v1 += 1
    switch ("foo bar" + "a") {
    case "a":
        v36 := GenS1{F0: false, F1: 65535, F2: 0}
        _ = v36
    }
    var v37 float64 = 1.5
    switch 32768 {
    case -1:
        switch "" {
        case "a":
            if ((-(v2)) != (v2 | genFn0())) {
                for fi38 := 0; fi38 < 7; fi38++ {
                    switch "a" {
                    case "foo bar":
                        for fi39 := 0; fi39 < 2; fi39++ {
                            delete(v5, (fi38 << ((18446744073709551615) & 63)))
                            v40 := GenS4{F0: false, F1: (v2 % ((v2) | 1))}
                            v41 := GenS5{F0: -167, F1: (42 < 128), F2: 600}
                            _ = fi39
                            _ = v40
                            _ = v41
                        }
                    case "":
                        v42 := GenS3{F0: -155, F1: 2147483647}
                        _ = v42
                    case "hello":
                        switch fi38 {
                        case 0:
                            if v4 {
                                for fi43 := 0; fi43 < 16; fi43++ {
                                    var v44 int64 = -1
                                    delete(v3, fi38)
                                    v45 := []int{fi38}
                                    v4 = (!v4)
                                    _ = fi43
                                    _ = v44
                                    _ = v45
                                }
                            } else {
                                _ = len(v3)
                            }
                        case -170:
                            v37 = v37
                        }
                    }
                    v46 := GenS6{F0: v4, F1: 2}
                    fmt.Println(v37, fi38)
                    _ = fi38
                    _ = v46
                }
                v47 := GenS7{F0: 273, F1: 1, F2: 926}
                _ = len(v3)
                _ = v47
            }
        case "foo bar":
            if true {
                v4 = (627 == 42)
                for fi48 := 0; fi48 < 15; fi48++ {
                    if (!(2147483648 < 0)) {
                        var v49 uint32 = 4294967295
                        _ = v49
                    }
                    _ = fi48
                }
            } else {
                if (-2147483648 < -2147483648) {
                    var v50 uint64 = 18446744073709551615
                    for fi51 := 0; fi51 < 4; fi51++ {
                        v52 := v3[fi51]
                        fmt.Println(v4, v52, v2)
                        _ = fi51
                        _ = v52
                    }
                    _ = v50
                } else {
                    var v53 bool = v4
                    v1 += 0
                    if v53 {
                        for fi54 := 0; fi54 < 18; fi54++ {
                            var v55 int16 = (^(v2))
                            var v56 int16 = 32767
                            v4 = v53
                            _ = fi54
                            _ = v55
                            _ = v56
                        }
                        for fi57 := 0; fi57 < 19; fi57++ {
                            if false {
                                v58 := v3[fi57]
                                switch fi57 {
                                case 127:
                                    for fi59 := 0; fi59 < 3; fi59++ {
                                        var v60 bool = true
                                        switch fi59 {
                                        case 32768:
                                            v61 := GenS1{F0: (0 < 4294967295), F1: 424, F2: 42}
                                            _ = v61
                                        case -152:
                                            if (!false) {
                                                switch (fi57 / ((fi59) | 1)) {
                                                case 42:
                                                    v2 = 1000
                                                case -9223372036854775808:
                                                    v62 := []int{-9223372036854775808, fi59, (fi57 / (((fi57 | fi57)) | 1)), fi59}
                                                    _ = v62
                                                case 546:
                                                    if ((v37 != 1e10) && (genFn0() >= (v2 | -351))) {
                                                        switch "" {
                                                        case "a":
                                                            fmt.Println(v58, v53)
                                                            var v63 uint8 = 2
                                                            _ = v63
                                                        case "hello":
                                                            v64 := GenS7{F0: 42, F1: 128, F2: v58}
                                                            if v53 {
                                                                v2 = v2
                                                            } else {
                                                                delete(v3, fi59)
                                                                v65 := GenS8{F0: -1, F1: "a", F2: (v2 + 1000)}
                                                                v66 := []int{(^(fi57)), fi59, (fi59 - fi59), fi57}
                                                                _ = v65
                                                                _ = v66
                                                            }
                                                            switch fi57 {
                                                            case -9223372036854775808:
                                                                v53 = (!v4)
                                                            case -1:
                                                                switch fi57 {
                                                                case 105:
                                                                    if false {
                                                                        v67 := GenS6{F0: (!false), F1: v1}
                                                                        v37 = (v37 + v37)
                                                                        v68 := GenS9{F0: ("" + "\n\t"), F1: 1, F2: 1}
                                                                        _ = v67
                                                                        _ = v68
                                                                    }
                                                                case -9223372036854775808:
                                                                    switch ("hello" + "\n\t") {
                                                                    case "foo bar":
                                                                        if (1 != int16(v58)) {
                                                                            var v69 int = (^(fi57))
                                                                            fmt.Println(v37)
                                                                            v70 := GenS10{F0: (v58 & 2147483648)}
                                                                            _ = v69
                                                                            _ = v70
                                                                        } else {
                                                                            if (1000 == v2) {
                                                                                var v71 uint32 = uint32(v1)
                                                                                _ = v71
                                                                            } else {
                                                                                if (162 >= v1) {
                                                                                    switch "\n\t" {
                                                                                    case "a":
                                                                                        switch "\n\t" {
                                                                                        case "foo bar":
                                                                                            fmt.Println(v37, v4, v2)
                                                                                        case "":
                                                                                            v1 = v1
                                                                                            fmt.Println(v60)
                                                                                            var v72 string = ("foo bar" + (((("a" + "foo bar") + "a") + "hello") + "\n\t"))
                                                                                            _ = v72
                                                                                        default:
                                                                                            fmt.Println(v4, fi59)
                                                                                        }
                                                                                    case "foo bar":
                                                                                        if (0.1 == 3.14159) {
                                                                                            switch "a" {
                                                                                            case "hello":
                                                                                                fmt.Println(fi59, v58, fi59)
                                                                                            case "\n\t":
                                                                                                v53 = (((42 == 185) || (!true)) || (!(("a" == "") && (v4 && (423 < 897)))))
                                                                                            }
                                                                                            if v53 {
                                                                                                v73 := []int{fi57}
                                                                                                var v74 string = "a"
                                                                                                var v75 uint16 = 32767
                                                                                                if (v4 || (255 == uint8(v58))) {
                                                                                                    v5[(fi59 / ((32767) | 1))] = (v58 * v58)
                                                                                                    if true {
                                                                                                        delete(v5, fi59)
                                                                                                        _ = len(v73)
                                                                                                        switch (v74 + "hello") {
                                                                                                        case "a":
                                                                                                            switch (fi57 | 0) {
                                                                                                            case 127:
                                                                                                                if v53 {
                                                                                                                    var v76 int = (fi57 + fi59)
                                                                                                                    fmt.Println(v53, v76)
                                                                                                                    _ = v76
                                                                                                                } else {
                                                                                                                    var v77 string = v74
                                                                                                                    _ = v77
                                                                                                                }
                                                                                                            case 32767:
                                                                                                                if (v58 >= v58) {
                                                                                                                    var v78 int64 = 9223372036854775807
                                                                                                                    switch ("" + (v74 + "a")) {
                                                                                                                    case "foo bar":
                                                                                                                        _ = len(v73)
                                                                                                                    default:
                                                                                                                        fmt.Println(v60, v53)
                                                                                                                    }
                                                                                                                    _ = v78
                                                                                                                }
                                                                                                            default:
                                                                                                                switch (fi59 * 128) {
                                                                                                                case 4294967295:
                                                                                                                    v2 = 648
                                                                                                                case 868:
                                                                                                                    var v79 int16 = genFn0()
                                                                                                                    _ = v79
                                                                                                                case 2147483647:
                                                                                                                    v58 = v58
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    v74 = ((v74 + v74) + v74)
                                                                                                } else {
                                                                                                    if true {
                                                                                                        if v53 {
                                                                                                            switch (fi57 / (((^(fi57))) | 1)) {
                                                                                                            case 361:
                                                                                                                v3[fi57] = v58
                                                                                                            default:
                                                                                                                v80 := GenS8{F0: 2147483647, F1: (v74 + v74), F2: genFn0()}
                                                                                                                _ = v80
                                                                                                            }
                                                                                                            var v81 float32 = 3.14159
                                                                                                            var v82 uint8 = 255
                                                                                                            _ = v81
                                                                                                            _ = v82
                                                                                                        } else {
                                                                                                            if v4 {
                                                                                                                _ = len(v73)
                                                                                                            }
                                                                                                        }
                                                                                                        var v83 int8 = 100
                                                                                                        _ = v83
                                                                                                    } else {
                                                                                                        fmt.Println(v60)
                                                                                                        v2 = 258
                                                                                                        if v60 {
                                                                                                            if v4 {
                                                                                                                if ((v75 <= v1) || (v53 || false)) {
                                                                                                                    fmt.Println(v75, fi57, fi57)
                                                                                                                } else {
                                                                                                                    switch 1 {
                                                                                                                    case 667:
                                                                                                                        v53 = ((fi57 * (fi59 ^ fi59)) < fi57)
                                                                                                                    case -977:
                                                                                                                        fmt.Println(v58, v53)
                                                                                                                    }
                                                                                                                    v37 = v37
                                                                                                                }
                                                                                                                _ = len(v3)
                                                                                                            }
                                                                                                            var v84 uint8 = 193
                                                                                                            _ = v84
                                                                                                        }
                                                                                                    }
                                                                                                    switch v74 {
                                                                                                    case "a":
                                                                                                        v2 -= genFn0()
                                                                                                    default:
                                                                                                        v2 -= v2
                                                                                                    }
                                                                                                }
                                                                                                _ = v73
                                                                                                _ = v74
                                                                                                _ = v75
                                                                                            } else {
                                                                                                switch fi57 {
                                                                                                case -1:
                                                                                                    v2 -= (v2 & genFn0())
                                                                                                case 128:
                                                                                                    v58 = (v58 << ((969) & 63))
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            v53 = (18446744073709551615 <= 18446744073709551615)
                                                                                            if (!v53) {
                                                                                                v85 := []int{-2147483648, (fi57 ^ 0), 84, -2147483648}
                                                                                                v37 -= (1.5 * 1e10)
                                                                                                _ = v85
                                                                                            } else {
                                                                                                var v86 float64 = v37
                                                                                                _ = v86
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        fmt.Println(v60)
                                                                                    }
                                                                                    var v87 float64 = float64(3.14159)
                                                                                    _ = v87
                                                                                } else {
                                                                                    var v88 uint8 = 128
                                                                                    fmt.Println(v58)
                                                                                    var v89 int16 = v2
                                                                                    _ = v88
                                                                                    _ = v89
                                                                                }
                                                                                var v90 int8 = 0
                                                                                if v53 {
                                                                                    v91 := GenS11{F0: 2147483647, F1: -299, F2: (v2 << ((4294967295) & 63))}
                                                                                    _ = v91
                                                                                }
                                                                                _ = v90
                                                                            }
                                                                            var v92 bool = (fi57 >= (fi59 << ((242) & 63)))
                                                                            fmt.Println(v58, fi57, v37)
                                                                            _ = v92
                                                                        }
                                                                    default:
                                                                        v93 := GenS4{F0: ((v2 | v2) < int16(v58)), F1: (-(v2))}
                                                                        _ = v93
                                                                    }
                                                                case 275:
                                                                    v94 := GenS4{F0: (42 == 0), F1: (^(v2))}
                                                                    _ = v94
                                                                }
                                                            default:
                                                                switch "hello" {
                                                                case "a":
                                                                    v2 = v2
                                                                case "\n\t":
                                                                    v60 = false
                                                                }
                                                            }
                                                            _ = v64
                                                        case "foo bar":
                                                            if false {
                                                                v2 += v2
                                                                var v95 uint16 = v1
                                                                v96 := GenS12{F0: 255, F1: 226}
                                                                _ = v95
                                                                _ = v96
                                                            }
                                                        default:
                                                            var v97 int8 = int8(fi59)
                                                            _ = v97
                                                        }
                                                        v98 := []int{fi59, (fi57 >> ((725) & 63)), (-(fi57)), (fi59 - (fi57 + fi57))}
                                                        _ = v98
                                                    }
                                                    var v99 float32 = 1e10
                                                    _ = v99
                                                default:
                                                    v100 := GenS9{F0: "", F1: 0, F2: 42}
                                                    _ = v100
                                                }
                                                if (v60 || ((fi57 - fi57) > (-(fi59)))) {
                                                    switch ("\n\t" + "a") {
                                                    case "":
                                                        delete(v5, (fi59 ^ 9223372036854775807))
                                                    case "hello":
                                                        if (!v53) {
                                                            v3[32768] = v58
                                                            var v101 float64 = ((v37 - (((v37 - v37) / float64(float32(v37))) - v37)) + v37)
                                                            _ = v101
                                                        }
                                                    }
                                                    v60 = (v4 && v60)
                                                    v102 := []int{fi59, (fi57 | fi59)}
                                                    delete(v5, (fi57 & fi57))
                                                    var v103 uint = 0
                                                    var v104 int64 = 1
                                                    _ = v102
                                                    _ = v103
                                                    _ = v104
                                                } else {
                                                    fmt.Println(fi59)
                                                    switch 1 {
                                                    case 42:
                                                        if (0 < 2) {
                                                            switch fi57 {
                                                            case -1:
                                                                if v60 {
                                                                    fmt.Println(v37, v2)
                                                                    if ((9223372036854775807 == 42) && (!(1000 >= (^(v2))))) {
                                                                        fmt.Println(v53)
                                                                        v105 := []int{(fi59 ^ -2147483648), fi57, 2147483647, (fi57 + (fi57 - fi59))}
                                                                        if len(v105) >= 2 {
                                                                            v105 = v105[1:len(v105)]
                                                                        }
                                                                        _ = v105
                                                                    }
                                                                } else {
                                                                    var v106 int32 = 1
                                                                    switch "" {
                                                                    case "":
                                                                        if (int16(v1) >= genFn0()) {
                                                                            if false {
                                                                                fmt.Println(v4)
                                                                            }
                                                                            v60 = false
                                                                            v37 = (v37 / 3.14159)
                                                                            if (false && ((v106 >> ((58) & 63)) == (v106 & (v106 << ((42) & 63))))) {
                                                                                v107 := GenS13{F0: v2, F1: 9223372036854775807, F2: 54}
                                                                                v108 := GenS14{F0: v1}
                                                                                v109 := []int{fi59, 656, fi59}
                                                                                v4 = (v4 || v60)
                                                                                _ = v107
                                                                                _ = v108
                                                                                _ = v109
                                                                            }
                                                                            if v53 {
                                                                                fmt.Println(fi59)
                                                                            }
                                                                        } else {
                                                                            switch "hello" {
                                                                            case "\n\t":
                                                                                var v110 uint8 = 0
                                                                                _ = v110
                                                                            case "hello":
                                                                                switch "hello" {
                                                                                case "foo bar":
                                                                                    fmt.Println(v106, v2)
                                                                                    v106 = (v106 ^ 42)
                                                                                case "":
                                                                                    v3[42] = (v58 & (v58 & v58))
                                                                                default:
                                                                                    var v111 int8 = -128
                                                                                    _ = v111
                                                                                }
                                                                            case "a":
                                                                                var v112 int64 = -934
                                                                                _ = v112
                                                                            default:
                                                                                if v53 {
                                                                                    v60 = (((!(!(v1 >= 32767))) && (!v4)) && (0.0 >= 1.5))
                                                                                } else {
                                                                                    v60 = (!v4)
                                                                                    if false {
                                                                                        var v113 float32 = -2.25
                                                                                        var v114 int32 = (^(v106))
                                                                                        _ = v113
                                                                                        _ = v114
                                                                                    }
                                                                                    v115 := GenS15{F0: v1, F1: v1}
                                                                                    _ = v115
                                                                                }
                                                                            }
                                                                            v116 := GenS16{F0: false, F1: 42}
                                                                            var v117 int16 = (v2 << ((255) & 63))
                                                                            _ = v116
                                                                            _ = v117
                                                                        }
                                                                    case "a":
                                                                        if (v4 || v60) {
                                                                            var v118 float64 = -2.25
                                                                            v119 := GenS13{F0: (^(v2)), F1: 9223372036854775807, F2: 0}
                                                                            _ = v118
                                                                            _ = v119
                                                                        }
                                                                    case "foo bar":
                                                                        v4 = (18446744073709551615 > 1)
                                                                    }
                                                                    v120 := GenS17{F0: "hello", F1: 1, F2: v106}
                                                                    _ = v106
                                                                    _ = v120
                                                                }
                                                                var v121 uint8 = 241
                                                                _ = v121
                                                            }
                                                            v122 := GenS2{F0: (v58 * 42)}
                                                            v60 = v60
                                                            _ = v122
                                                        }
                                                    }
                                                    v1 += 1
                                                }
                                            }
                                        default:
                                            v123 := GenS17{F0: "", F1: 0, F2: -2147483648}
                                            _ = v123
                                        }
                                        _ = fi59
                                        _ = v60
                                    }
                                }
                                _ = v58
                            }
                            v53 = (!v4)
                            _ = fi57
                        }
                        v1 = v1
                    }
                    _ = v53
                }
            }
        case "":
            fmt.Println(v4, v37, v2)
        default:
            var v124 uint8 = 1
            _ = v124
        }
    case 0:
        var v125 uint64 = 4294967295
        _ = v125
    default:
        switch (("\n\t" + ("foo bar" + "a")) + "a") {
        case "a":
            var v126 uint32 = 808
            _ = v126
        case "foo bar":
            var v127 uint32 = 1
            _ = v127
        case "\n\t":
            fmt.Println(v4)
        }
    }
    var v128 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v128:", v128)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v128
    _ = v17
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
