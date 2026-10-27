package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 uint32; F1 int; F2 int16 }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint16; F1 int16; F2 int32 }
type GenS9 struct { F0 int32; F1 uint; F2 uint32 }
type GenS10 struct { F0 int8; F1 uint64; F2 uint }
type GenS11 struct { F0 bool }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    v4 := make(map[int]uint32)
    var v5 bool = (0 < 845)
    v5 = (206 == 127)
    var v6 uint32 = 918
    if ((0.0 * float32(1e10)) < (3.14159 / 1.0)) {
        v7 := GenS1{F0: 2147483647, F1: false}
        v6 = v6
        v8 := GenS1{F0: 2147483647, F1: false}
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 20; fi9++ {
            switch (fi9 % (((fi9 / ((fi9) | 1))) | 1)) {
            case -561:
                switch 534 {
                case -773:
                    var v10 int = (fi9 % ((9223372036854775807) | 1))
                    _ = v10
                case 1:
                    v5 = (674 == 4294967295)
                default:
                    v5 = (682 > 9223372036854775807)
                }
            case 128:
                fmt.Println(fi9)
            case 271:
                fmt.Println(v6, fi9)
            default:
                for fi11 := 0; fi11 < 2; fi11++ {
                    v12 := GenS1{F0: (v6 << ((1) & 63)), F1: true}
                    var v13 int = -2147483648
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            }
            _ = fi9
        }
    }
    if v5 {
        v5 = (v5 && false)
        if (!true) {
            v14 := GenS2{F0: -1}
            var v15 bool = (-812 < -1)
            _ = v14
            _ = v15
        }
        v6 -= v6
    } else {
        v5 = false
        for fi16 := 0; fi16 < 12; fi16++ {
            v5 = ((2 >= 32767) && (18446744073709551615 > 780))
            v5 = (0 == 0)
            v6 += v6
            _ = fi16
        }
        if (754 > uint64(v6)) {
            v6 += 225
        }
    }
    if (100 < 100) {
        var v17 float64 = -2.25
        fmt.Println(v17, v6)
        _ = v17
    }
    if false {
        v18 := GenS2{F0: 42}
        v19 := v3[32767]
        _ = v18
        _ = v19
    } else {
        var v20 uint = uint(v6)
        v21 := GenS2{F0: 468}
        _ = v20
        _ = v21
    }
    v22 := []int{128}
    fmt.Println(v5)
    switch 1 {
    case 2147483647:
        var v23 int64 = 9223372036854775807
        _ = v23
    case 128:
        v24 := GenS3{F0: ("hello" + "\n\t"), F1: 42}
        _ = v24
    case 32767:
        for fi25 := 0; fi25 < 2; fi25++ {
            if ((0.1 + -2.25) == ((((3.14159 * 0.0) / 1e10) + 1.5) + ((-2.25 - (((((-2.25 * 3.14159) - 0.0) - float32(-2.25)) * -2.25) - -2.25)) / 1.5))) {
                v26 := GenS4{F0: (v6 ^ v6), F1: fi25, F2: 1}
                if (false && true) {
                    var v27 uint = 265
                    if len(v22) > 0 {
                        v22[0] = (fi25 / ((fi25) | 1))
                    }
                    fmt.Println(v6)
                    _ = v27
                } else {
                    var v28 int32 = -2147483648
                    switch "hello" {
                    case "":
                        switch 32768 {
                        case -9223372036854775808:
                            switch "a" {
                            case "hello":
                                var v29 int = 32768
                                _ = v29
                            case "":
                                var v30 string = "foo bar"
                                _ = v30
                            }
                        case 32768:
                            if (!false) {
                                fmt.Println(v5, v28, v6)
                                if (v28 > v28) {
                                    v31 := GenS3{F0: ("" + "hello"), F1: 422}
                                    _ = v31
                                }
                            } else {
                                v32 := GenS5{F0: (!v5), F1: 42}
                                _ = v32
                            }
                        case 681:
                            if len(v22) > 0 {
                                v22[0] = fi25
                            }
                        }
                    default:
                        var v33 int64 = -9223372036854775808
                        _ = v33
                    }
                    if v5 {
                        if v5 {
                            if false {
                                for fi34 := 0; fi34 < 10; fi34++ {
                                    fmt.Println(v28, fi25, fi25)
                                    v28 -= v28
                                    _ = fi34
                                }
                            }
                            v28 = (v28 + v28)
                        }
                        v35 := GenS6{F0: 65535, F1: "\n\t"}
                        _ = v35
                    } else {
                        if v5 {
                            v36 := GenS6{F0: 255, F1: "\n\t"}
                            var v37 uint64 = 18446744073709551615
                            _ = v36
                            _ = v37
                        } else {
                            for fi38 := 0; fi38 < 12; fi38++ {
                                fmt.Println(fi38, v5)
                                var v39 bool = v5
                                _ = fi38
                                _ = v39
                            }
                            delete(v4, (fi25 / ((fi25) | 1)))
                            for fi40 := 0; fi40 < 12; fi40++ {
                                v41 := v1[v6]
                                var v42 uint16 = 225
                                _ = fi40
                                _ = v41
                                _ = v42
                            }
                        }
                        v5 = (4 > -33)
                    }
                    _ = v28
                }
                _ = v26
            }
            delete(v4, 32767)
            _ = fi25
        }
    default:
        fmt.Println(v6, v6)
    }
    switch 876 {
    case 2147483647:
        for fi43 := 0; fi43 < 12; fi43++ {
            v5 = false
            v5 = (255 > 127)
            _ = fi43
        }
    case 0:
        if true {
            for fi44 := 0; fi44 < 10; fi44++ {
                if len(v22) > 0 {
                    v22[0] = fi44
                }
                if true {
                    if v5 {
                        switch "foo bar" {
                        case "\n\t":
                            if true {
                                for fi45 := 0; fi45 < 1; fi45++ {
                                    switch 205 {
                                    case 42:
                                        if (!v5) {
                                            delete(v1, v6)
                                            var v46 uint = 235
                                            v6 = (v6 + v6)
                                            fmt.Println(v6, v46)
                                            _ = v46
                                        }
                                    default:
                                        v6 = v6
                                    }
                                    v6 -= 65535
                                    var v47 bool = v5
                                    _ = fi45
                                    _ = v47
                                }
                                v6 -= (v6 << ((65535) & 63))
                                var v48 uint16 = 65535
                                _ = v48
                            } else {
                                fmt.Println(v5, v5, fi44)
                            }
                        case "a":
                            for fi49 := 0; fi49 < 20; fi49++ {
                                v50 := GenS5{F0: v5, F1: int16(fi49)}
                                fmt.Println(fi44, fi49)
                                switch fi49 {
                                case -2147483648:
                                    var v51 uint16 = 42
                                    _ = v51
                                default:
                                    switch (fi44 % ((4294967295) | 1)) {
                                    case -2147483648:
                                        fmt.Println(fi44, fi44)
                                        fmt.Println(v5, v6, v5)
                                        v5 = v5
                                        v6 += 0
                                        switch (^(fi49)) {
                                        case 660:
                                            if v5 {
                                                switch fi44 {
                                                case -2147483648:
                                                    v22 = append(v22, (fi44 ^ 630))
                                                case 145:
                                                    var v52 uint16 = 388
                                                    _ = v52
                                                case -381:
                                                    v5 = (18446744073709551615 > 4294967295)
                                                }
                                            }
                                        case 2147483647:
                                            v22 = append(v22, (fi44 | fi44))
                                        case -581:
                                            var v53 uint8 = 42
                                            _ = v53
                                        }
                                        _ = len(v22)
                                    case 345:
                                        var v54 float64 = (3.14159 + float64(-2.25))
                                        _ = v54
                                    default:
                                        var v55 uint16 = 0
                                        switch int(v6) {
                                        case 1:
                                            var v56 int64 = -9223372036854775808
                                            _ = v56
                                        case 805:
                                            var v57 uint64 = 0
                                            _ = v57
                                        case 544:
                                            var v58 int64 = 42
                                            _ = v58
                                        }
                                        _ = v55
                                    }
                                }
                                _ = fi49
                                _ = v50
                            }
                        case "":
                            v1[(v6 % (((^(v6))) | 1))] = v5
                        default:
                            v5 = (false || ((("" + "foo bar") + (("foo bar" + "hello") + "\n\t")) < ("foo bar" + "a")))
                        }
                    } else {
                        switch "hello" {
                        case "hello":
                            v5 = (v5 || v5)
                        case "foo bar":
                            fmt.Println(v6, fi44, v6)
                        case "\n\t":
                            for fi59 := 0; fi59 < 17; fi59++ {
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                                v5 = (32767 >= int16(fi44))
                                v6 -= (v6 << ((666) & 63))
                                _ = fi59
                            }
                        default:
                            switch "foo bar" {
                            case "":
                                if (!v5) {
                                    var v60 uint8 = 0
                                    fmt.Println(fi44, v60, fi44)
                                    _ = v60
                                }
                            }
                        }
                    }
                }
                v5 = v5
                _ = fi44
            }
            if (!(255 <= 127)) {
                v1[2147483648] = (true && (-1 == 2147483647))
                for fi61 := 0; fi61 < 7; fi61++ {
                    if (9223372036854775807 <= 568) {
                        if v5 {
                            v5 = (37 == -32768)
                            if v5 {
                                var v62 int64 = int64(v6)
                                v63 := GenS7{F0: 121}
                                _ = len(v22)
                                _ = v62
                                _ = v63
                            }
                            v5 = true
                        }
                        if len(v22) >= 2 {
                            v22 = v22[1:len(v22)]
                        }
                    }
                    for fi64 := 0; fi64 < 13; fi64++ {
                        v5 = v5
                        _ = fi64
                    }
                    for fi65 := 0; fi65 < 19; fi65++ {
                        v5 = true
                        var v66 float64 = (-2.25 + 1e10)
                        fmt.Println(v5, fi61, fi61)
                        _ = fi65
                        _ = v66
                    }
                    _ = fi61
                }
                if ((-195 < -825) && v5) {
                    if true {
                        if len(v22) > 0 {
                            v22[0] = -245
                        }
                        _ = len(v22)
                    } else {
                        var v67 int8 = 42
                        if len(v22) >= 2 {
                            v22 = v22[1:len(v22)]
                        }
                        _ = v67
                    }
                    v5 = (255 <= 1)
                    if (!(!(0.1 >= ((0.0 + 1.5) + -2.25)))) {
                        if ((42 == 1) || false) {
                            for fi68 := 0; fi68 < 4; fi68++ {
                                v6 += v6
                                switch fi68 {
                                case -9223372036854775808:
                                    if (v5 || (262 < 730)) {
                                        switch (fi68 | fi68) {
                                        case 32767:
                                            switch fi68 {
                                            case 4294967295:
                                                v69 := GenS8{F0: 581, F1: -608, F2: 42}
                                                _ = v69
                                            default:
                                                for fi70 := 0; fi70 < 20; fi70++ {
                                                    if (-805 != -9223372036854775808) {
                                                        v71 := GenS4{F0: (v6 - v6), F1: (fi70 + (fi68 + fi68)), F2: 393}
                                                        _ = v71
                                                    }
                                                    var v72 int8 = 97
                                                    _ = fi70
                                                    _ = v72
                                                }
                                            }
                                        default:
                                            v73 := GenS7{F0: -1}
                                            _ = v73
                                        }
                                    } else {
                                        var v74 uint = 930
                                        var v75 bool = (v5 && (!v5))
                                        _ = v74
                                        _ = v75
                                    }
                                case -2147483648:
                                    if v5 {
                                        var v76 uint8 = 255
                                        _ = v76
                                    }
                                default:
                                    v6 += v6
                                }
                                _ = fi68
                            }
                        }
                    }
                }
            }
        }
    case 58:
        if (v5 || ("" != ("hello" + "hello"))) {
            switch "a" {
            case "foo bar":
                v3[787] = (v6 ^ (v6 / ((0) | 1)))
            case "\n\t":
                v6 -= v6
            default:
                if ((42 <= 1) && (!true)) {
                    for fi77 := 0; fi77 < 9; fi77++ {
                        if (!v5) {
                            var v78 int64 = -561
                            if ((v6 | 42) >= (v6 % ((v6) | 1))) {
                                var v79 uint = 818
                                _ = v79
                            }
                            var v80 int64 = 1
                            _ = v78
                            _ = v80
                        }
                        for fi81 := 0; fi81 < 1; fi81++ {
                            var v82 uint8 = 42
                            _ = fi81
                            _ = v82
                        }
                        var v83 float32 = 0.0
                        _ = fi77
                        _ = v83
                    }
                } else {
                    v5 = (101 <= 1)
                    if len(v22) > 0 {
                        v22[0] = 127
                    }
                    v6 = v6
                }
            }
            v84 := GenS4{F0: v6, F1: -1, F2: -1}
            for fi85 := 0; fi85 < 2; fi85++ {
                for fi86 := 0; fi86 < 18; fi86++ {
                    var v87 int8 = -1
                    var v88 uint16 = 233
                    _ = fi86
                    _ = v87
                    _ = v88
                }
                if v5 {
                    fmt.Println(v6)
                } else {
                    var v89 uint64 = 293
                    _ = len(v2)
                    _ = v89
                }
                _ = fi85
            }
            _ = v84
        }
    }
    for fi90 := 0; fi90 < 1; fi90++ {
        v5 = (v5 || v5)
        var v91 float32 = -2.25
        _ = fi90
        _ = v91
    }
    if false {
        v6 = (v6 + v6)
        for fi92 := 0; fi92 < 19; fi92++ {
            var v93 uint = 2
            _ = fi92
            _ = v93
        }
        var v94 int16 = -32768
        _ = v94
    } else {
        if (-60 >= -32768) {
            switch 574 {
            case 9223372036854775807:
                if v5 {
                    for fi95 := 0; fi95 < 19; fi95++ {
                        for fi96 := 0; fi96 < 11; fi96++ {
                            v97 := GenS2{F0: 9223372036854775807}
                            v98 := GenS9{F0: int32(v6), F1: 255, F2: v6}
                            _ = fi96
                            _ = v97
                            _ = v98
                        }
                        v5 = (-32768 >= -32768)
                        switch ("hello" + ("foo bar" + "hello")) {
                        case "":
                            if v5 {
                                v99 := GenS10{F0: -1, F1: 900, F2: 261}
                                v100 := GenS11{F0: v5}
                                v6 -= 2147483647
                                _ = v100
                                _ = v99
                            }
                        case "a":
                            var v101 float64 = (3.14159 - (-2.25 - float64(float32(((1.5 + 0.1) - -2.25)))))
                            _ = v101
                        case "hello":
                            v5 = ("\n\t" != "")
                        }
                        _ = fi95
                    }
                }
            case 2147483647:
                for fi102 := 0; fi102 < 14; fi102++ {
                    v5 = false
                    var v103 uint = 2
                    _ = fi102
                    _ = v103
                }
            case 4294967295:
                v5 = false
            }
        } else {
            fmt.Println(v6)
            v5 = (!v5)
            v104 := GenS8{F0: 65535, F1: 107, F2: -105}
            _ = v104
        }
    }
    delete(v2, 1)
    fmt.Println(v6, v6)
    if len(v22) > 0 {
        v22[0] = 898
    }
    switch (("" + "") + "\n\t") {
    case "foo bar":
        v6 += 556
    case "":
        v2[32767] = v5
    case "a":
        var v105 float64 = -2.25
        _ = v105
    default:
        if true {
            if (!(!(!(488 != 65535)))) {
                v106 := GenS11{F0: true}
                delete(v4, -868)
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
                _ = v106
            }
            for fi107 := 0; fi107 < 3; fi107++ {
                fmt.Println(fi107, fi107)
                v5 = v5
                _ = fi107
            }
            var v108 int64 = -9223372036854775808
            _ = v108
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
