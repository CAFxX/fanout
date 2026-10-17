package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint64; F2 int }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 int8; F1 uint16; F2 int64 }
type GenS7 struct { F0 int; F1 string }

func main() {
    var v1 uint16 = 0
    v2 := make(map[string]uint32)
    var v3 int8 = 1
    var v4 float64 = float64(((1.5 + 1.5) / -2.25))
    if (!(("a" + ("" + "hello")) > (("" + "\n\t") + ("a" + ("" + "\n\t"))))) {
        v5 := GenS1{F0: -2147483648}
        if (1.5 != 0.0) {
            var v6 uint64 = 773
            _ = v6
        } else {
            switch "foo bar" {
            case "hello":
                for fi7 := 0; fi7 < 7; fi7++ {
                    v4 = (float64(0.1) + 1e10)
                    v2["a"] = 2147483648
                    _ = fi7
                }
            case "":
                if (42 != 42) {
                    var v8 string = "hello"
                    if ((-2.25 / 0.1) > v4) {
                        v9 := []int{9223372036854775807, 431, 128}
                        _ = v9
                    } else {
                        v4 += -2.25
                        switch "a" {
                        case "a":
                            v8 = (v8 + "")
                        case "":
                            var v10 float32 = ((3.14159 - float32((1e10 / v4))) * 0.0)
                            _ = v10
                        case "foo bar":
                            v3 = 100
                        }
                        for fi11 := 0; fi11 < 16; fi11++ {
                            fmt.Println(v3, v4)
                            for fi12 := 0; fi12 < 2; fi12++ {
                                if true {
                                    switch "" {
                                    case "a":
                                        switch 127 {
                                        case 4294967295:
                                            fmt.Println(v8)
                                        case 2147483647:
                                            if (42 != -2147483648) {
                                                var v13 float64 = ((0.1 * v4) * v4)
                                                v14 := GenS1{F0: (fi11 ^ (fi11 + -1))}
                                                v1 = v1
                                                _ = v13
                                                _ = v14
                                            } else {
                                                _ = len(v2)
                                                var v15 float32 = 3.14159
                                                v15 += (-2.25 - ((v15 - (0.0 - (v15 * 3.14159))) / 1.0))
                                                fmt.Println(v8, v3, v15)
                                                _ = v15
                                            }
                                        }
                                    }
                                }
                                var v16 int = fi12
                                _ = fi12
                                _ = v16
                            }
                            if true {
                                v8 = "hello"
                            }
                            _ = fi11
                        }
                    }
                    var v17 float64 = (((1e10 - float64((0.1 - 3.14159))) + ((v4 - v4) - float64(float32((v4 + 0.1))))) + v4)
                    _ = v17
                    _ = v8
                } else {
                    var v18 float32 = 0.1
                    if (v18 >= 1e10) {
                        switch 9223372036854775807 {
                        case -9223372036854775808:
                            v19 := v2["hello"]
                            _ = v19
                        case 1:
                            var v20 string = ("hello" + ("hello" + "a"))
                            _ = v20
                        case 543:
                            v4 += float64((1e10 * v18))
                        }
                    } else {
                        v21 := GenS2{F0: 0}
                        if (42 < 0) {
                            if false {
                                if (2 <= 915) {
                                    _ = len(v2)
                                    v18 += (0.1 + (float32(1.5) - (1.5 * float32(0.1))))
                                }
                            } else {
                                switch "a" {
                                case "\n\t":
                                    v4 = (-2.25 / (((-2.25 / v4) * v4) / 1.5))
                                case "foo bar":
                                    v1 = (v1 % ((1) | 1))
                                default:
                                    fmt.Println(v3, v1)
                                }
                                switch ("\n\t" + ("" + "\n\t")) {
                                case "foo bar":
                                    fmt.Println(v3)
                                case "a":
                                    _ = len(v2)
                                default:
                                    if true {
                                        v4 = float64(float32(v4))
                                        v22 := GenS2{F0: 65535}
                                        v4 -= v4
                                        _ = v22
                                    }
                                }
                                var v23 int32 = 2147483647
                                _ = v23
                            }
                        } else {
                            var v24 string = "foo bar"
                            v18 -= 0.0
                            var v25 uint8 = 78
                            _ = v24
                            _ = v25
                        }
                        _ = v21
                    }
                    _ = v18
                }
            case "a":
                switch (("a" + "a") + "a") {
                case "":
                    v26 := GenS3{F0: 4294967295, F1: 2, F2: -1}
                    _ = v26
                default:
                    var v27 uint8 = 128
                    _ = v27
                }
            }
            var v28 int8 = (v3 | (v3 & (^(v3))))
            _ = v28
        }
        _ = v5
    } else {
        if false {
            v29 := []int{-9223372036854775808, 9223372036854775807}
            if (2 >= 2147483648) {
                for fi30 := 0; fi30 < 11; fi30++ {
                    for fi31 := 0; fi31 < 12; fi31++ {
                        if (((3.14159 + (0.1 + -2.25)) + (1e10 - 1.5)) != float32(v4)) {
                            var v32 string = ("a" + ("foo bar" + "a"))
                            var v33 uint64 = 0
                            _ = v32
                            _ = v33
                        }
                        var v34 uint64 = 711
                        if len(v29) > 0 {
                            v29[0] = (fi31 + (^(fi30)))
                        }
                        _ = fi31
                        _ = v34
                    }
                    _ = fi30
                }
                if true {
                    v1 = (v1 + (v1 & v1))
                    switch "hello" {
                    case "a":
                        switch 0 {
                        case -348:
                            v35 := GenS3{F0: 255, F1: 18446744073709551615, F2: 9223372036854775807}
                            _ = v35
                        case 763:
                            if (!(127 != 1)) {
                                switch -909 {
                                case -692:
                                    v4 -= 1.5
                                case -2147483648:
                                    var v36 uint = 933
                                    _ = v36
                                default:
                                    var v37 string = ""
                                    _ = v37
                                }
                                fmt.Println(v1, v1, v1)
                            } else {
                                v1 += (v1 ^ (^(v1)))
                                if false {
                                    var v38 int8 = v3
                                    if true {
                                        if false {
                                            v39 := GenS3{F0: 143, F1: 423, F2: 9223372036854775807}
                                            switch 32768 {
                                            case -924:
                                                var v40 uint32 = 600
                                                _ = v40
                                            case 42:
                                                var v41 bool = false
                                                _ = v41
                                            case -433:
                                                v38 = (v38 >> ((688) & 63))
                                            }
                                            fmt.Println(v3, v4)
                                            _ = v39
                                        }
                                        fmt.Println(v4, v3, v38)
                                    } else {
                                        v38 = (v3 * 127)
                                        if true {
                                            v1 -= (v1 - 0)
                                            var v42 float64 = v4
                                            v43 := v2["\n\t"]
                                            _ = v42
                                            _ = v43
                                        }
                                        var v44 uint16 = v1
                                        _ = v44
                                    }
                                    _ = v38
                                } else {
                                    var v45 uint32 = 810
                                    _ = v45
                                }
                            }
                        case 608:
                            var v46 int8 = 0
                            _ = v46
                        }
                    case "foo bar":
                        var v47 int8 = 0
                        _ = v47
                    default:
                        fmt.Println(v4)
                    }
                    v3 = v3
                } else {
                    v2[("foo bar" + ("\n\t" + "a"))] = 42
                    v3 = (v3 / ((int8(v1)) | 1))
                    v3 += (v3 - v3)
                }
            }
            _ = v29
        }
        if (4294967295 != 0) {
            if ((v1 + 32767) < v1) {
                for fi48 := 0; fi48 < 7; fi48++ {
                    if true {
                        v2["foo bar"] = 124
                        v49 := []int{9223372036854775807, 2147483647, fi48}
                        v1 = v1
                        _ = v49
                    } else {
                        var v50 int32 = 1
                        v50 += v50
                        _ = v50
                    }
                    if ((0 <= -32768) || ((v1 & (v1 >> ((255) & 63))) <= 84)) {
                        v3 -= v3
                    }
                    if true {
                        fmt.Println(fi48, v3)
                        for fi51 := 0; fi51 < 10; fi51++ {
                            v52 := []int{(fi51 * (fi48 * (fi51 / (((fi51 | 1)) | 1)))), fi51, fi51}
                            fmt.Println(v4)
                            _ = fi51
                            _ = v52
                        }
                        v53 := []int{}
                        _ = v53
                    }
                    _ = fi48
                }
                v54 := []int{}
                _ = v54
            }
            var v55 int16 = -7
            v3 += (v3 >> ((uint(v1)) & 63))
            _ = v55
        } else {
            v1 -= 2
            switch 982 {
            case -2147483648:
                v1 = v1
            case -1:
                switch ("foo bar" + "foo bar") {
                case "\n\t":
                    v56 := GenS3{F0: 255, F1: 42, F2: -2147483648}
                    _ = v56
                }
            default:
                for fi57 := 0; fi57 < 7; fi57++ {
                    var v58 int16 = 32767
                    _ = fi57
                    _ = v58
                }
            }
            if false {
                var v59 int64 = -1
                switch "a" {
                case "":
                    if ((("hello" + (("\n\t" + "hello") + "foo bar")) <= ("" + "\n\t")) && false) {
                        if true {
                            var v60 int16 = 0
                            _ = v60
                        }
                        fmt.Println(v59, v59, v4)
                    } else {
                        v61 := GenS1{F0: 742}
                        switch "a" {
                        case "hello":
                            v62 := GenS4{F0: (-(v59)), F1: 0}
                            _ = v62
                        case "":
                            for fi63 := 0; fi63 < 19; fi63++ {
                                fmt.Println(v3, v59, v59)
                                v59 -= (v59 | v59)
                                if true {
                                    var v64 float64 = v4
                                    var v65 string = ""
                                    v66 := GenS2{F0: 65535}
                                    _ = v64
                                    _ = v65
                                    _ = v66
                                }
                                _ = fi63
                            }
                        }
                        _ = v61
                    }
                case "a":
                    var v67 uint16 = (v1 | (v1 - v1))
                    _ = v67
                default:
                    switch "foo bar" {
                    case "":
                        var v68 string = "\n\t"
                        _ = v68
                    }
                }
                if false {
                    fmt.Println(v3)
                    for fi69 := 0; fi69 < 5; fi69++ {
                        if false {
                            for fi70 := 0; fi70 < 8; fi70++ {
                                fmt.Println(fi69, v59, fi69)
                                var v71 string = (("a" + "") + "hello")
                                var v72 int32 = -2147483648
                                v73 := GenS1{F0: (^(fi69))}
                                _ = fi70
                                _ = v71
                                _ = v72
                                _ = v73
                            }
                        }
                        fmt.Println(v3, v3, v4)
                        _ = fi69
                    }
                } else {
                    var v74 int = 42
                    var v75 int = v74
                    var v76 int64 = (v59 >> ((721) & 63))
                    _ = v74
                    _ = v75
                    _ = v76
                }
                _ = v59
            }
        }
        switch 0 {
        case -221:
            var v77 bool = (true || false)
            _ = v77
        case 974:
            switch 42 {
            case 236:
                var v78 int64 = -319
                _ = v78
            case 127:
                fmt.Println(v1, v3, v4)
            case 32767:
                var v79 int = 25
                _ = v79
            default:
                for fi80 := 0; fi80 < 13; fi80++ {
                    var v81 int16 = 42
                    var v82 int = 0
                    v83 := GenS2{F0: 781}
                    _ = fi80
                    _ = v81
                    _ = v82
                    _ = v83
                }
            }
        default:
            v4 -= -2.25
        }
    }
    switch "\n\t" {
    case "foo bar":
        v84 := v2[""]
        _ = v84
    case "\n\t":
        for fi85 := 0; fi85 < 7; fi85++ {
            if ((1e10 != ((v4 + (v4 * v4)) + 1e10)) && true) {
                v86 := GenS4{F0: -747, F1: 128}
                v4 = ((-2.25 + float64(1e10)) - v4)
                _ = v86
            }
            var v87 uint64 = 0
            if false {
                for fi88 := 0; fi88 < 20; fi88++ {
                    v89 := GenS1{F0: fi85}
                    _ = fi88
                    _ = v89
                }
                fmt.Println(v1, v3)
                v3 += -41
            }
            _ = fi85
            _ = v87
        }
    }
    fmt.Println(v3, v1, v1)
    for fi90 := 0; fi90 < 12; fi90++ {
        v1 -= (v1 / ((v1) | 1))
        var v91 float32 = 0.0
        var v92 uint8 = 0
        _ = fi90
        _ = v91
        _ = v92
    }
    fmt.Println(v3, v1, v3)
    switch ("\n\t" + ("foo bar" + ((("" + "") + ("hello" + "foo bar")) + "foo bar"))) {
    case "hello":
        for fi93 := 0; fi93 < 13; fi93++ {
            v94 := []int{}
            fmt.Println(v4, v3)
            for fi95 := 0; fi95 < 9; fi95++ {
                if (893 <= 2) {
                    v1 += v1
                    if (532 != 628) {
                        fmt.Println(fi93, v4, v4)
                        if false {
                            var v96 int8 = v3
                            _ = v96
                        }
                    }
                } else {
                    v4 += v4
                    switch (fi93 * fi93) {
                    case -2147483648:
                        v97 := GenS4{F0: -1, F1: 0}
                        _ = v97
                    }
                    fmt.Println(v3, v3, v4)
                }
                fmt.Println(v4, fi93, v3)
                _ = fi95
            }
            _ = fi93
            _ = v94
        }
    case "\n\t":
        for fi98 := 0; fi98 < 16; fi98++ {
            v3 = v3
            _ = fi98
        }
    case "":
        v1 = 255
    default:
        if false {
            for fi99 := 0; fi99 < 10; fi99++ {
                fmt.Println(v4, v3)
                var v100 int = (fi99 + (fi99 + fi99))
                _ = fi99
                _ = v100
            }
            var v101 float32 = 3.14159
            v2[("foo bar" + "foo bar")] = 2147483647
            _ = v101
        }
    }
    if false {
        v4 = ((v4 - (1.5 / v4)) / v4)
    } else {
        v3 -= v3
        v4 -= 1e10
        switch "" {
        case "":
            switch -9223372036854775808 {
            case 9223372036854775807:
                var v102 uint32 = 42
                _ = v102
            default:
                v103 := GenS5{F0: 1, F1: 42}
                _ = v103
            }
        default:
            switch (("a" + "a") + "a") {
            case "hello":
                if (127 <= 233) {
                    if false {
                        var v104 float64 = float64(3.14159)
                        if (-552 != 42) {
                            v1 = v1
                        }
                        _ = v104
                    }
                    if (false && (((42 < 564) && true) && (4294967295 <= 2))) {
                        var v105 uint32 = 719
                        v1 = v1
                        _ = v105
                    }
                }
            }
        }
    }
    v106 := GenS6{F0: (v3 | 42), F1: 32767, F2: 9223372036854775807}
    v107 := GenS2{F0: 104}
    v4 += float64(1e10)
    for fi108 := 0; fi108 < 16; fi108++ {
        var v109 float32 = (3.14159 * 0.0)
        var v110 uint = 82
        _ = fi108
        _ = v109
        _ = v110
    }
    if ("\n\t" < "foo bar") {
        if false {
            v1 += v1
        } else {
            v4 = v4
            fmt.Println(v4)
        }
        var v111 string = "hello"
        _ = v111
    } else {
        var v112 int32 = 0
        switch ("\n\t" + "foo bar") {
        case "\n\t":
            v112 = v112
        case "":
            var v113 uint = 2
            _ = v113
        }
        switch int(v1) {
        case 227:
            var v114 int8 = v3
            _ = v114
        default:
            v115 := GenS7{F0: -9223372036854775808, F1: ""}
            _ = v115
        }
        _ = v112
    }
    for fi116 := 0; fi116 < 10; fi116++ {
        var v117 int8 = (v3 + v3)
        switch fi116 {
        case 2147483647:
            var v118 int = (fi116 / ((0) | 1))
            _ = v118
        default:
            v119 := []int{}
            _ = v119
        }
        _ = fi116
        _ = v117
    }
    if false {
        v4 -= (v4 / (float64(1.5) * ((v4 + (v4 / v4)) - ((0.0 + (((1e10 - -2.25) * ((v4 / v4) / v4)) - ((((((1.5 - v4) - 1.5) - v4) * ((v4 * (v4 - v4)) / float64(3.14159))) * 0.1) / v4))) - (-2.25 * v4)))))
        v3 = 100
        switch "foo bar" {
        case "foo bar":
            v3 -= v3
        case "\n\t":
            if (4294967295 < 521) {
                for fi120 := 0; fi120 < 8; fi120++ {
                    if (false || (((fi120 ^ (fi120 / (((fi120 * fi120)) | 1))) != 160) || false)) {
                        for fi121 := 0; fi121 < 11; fi121++ {
                            var v122 int8 = v3
                            v123 := GenS5{F0: 4294967295, F1: 1}
                            v124 := []int{643, fi120, fi121}
                            _ = fi121
                            _ = v122
                            _ = v123
                            _ = v124
                        }
                    }
                    v4 = -2.25
                    v3 = (v3 / (((v3 & (v3 & (v3 ^ 1)))) | 1))
                    _ = fi120
                }
                delete(v2, "\n\t")
            }
        case "hello":
            v4 = (v4 * v4)
        }
    } else {
        v3 = (v3 - 65)
        var v125 string = ("\n\t" + "foo bar")
        _ = v125
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v106
    _ = v107
    _ = v2
    _ = v3
    _ = v4
}
