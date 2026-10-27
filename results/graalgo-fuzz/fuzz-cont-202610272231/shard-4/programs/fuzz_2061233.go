package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint8; F1 int8; F2 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 bool; F1 int16; F2 uint }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int; F1 int8; F2 int16 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int64; F1 int32; F2 int32 }
type GenS13 struct { F0 int8; F1 int }
type GenS14 struct { F0 string; F1 int8; F2 int8 }
type GenS15 struct { F0 uint16; F1 string; F2 uint }
func genFn0() uint {
    var v4 float32 = (float32(-2.25) * 3.14159)
    if true {
        for fi5 := 0; fi5 < 11; fi5++ {
            v6 := []int{-69, fi5, fi5, fi5}
            _ = fi5
            _ = v6
        }
        if ((0 == 429) && (true && false)) {
            v7 := GenS1{F0: 2147483647}
            _ = v7
        } else {
            v8 := GenS2{F0: -9223372036854775808}
            var v9 int16 = -1
            v9 = 545
            _ = v8
            _ = v9
        }
        v4 = v4
    } else {
        v4 = 0.0
    }
    v10 := GenS1{F0: -9223372036854775808}
    for fi11 := 0; fi11 < 18; fi11++ {
        v12 := GenS1{F0: fi11}
        if ((fi11 % ((fi11) | 1)) == (fi11 * fi11)) {
            fmt.Println(fi11, v4)
            if false {
                v13 := GenS1{F0: 32767}
                fmt.Println(v4, v4)
                for fi14 := 0; fi14 < 13; fi14++ {
                    var v15 float32 = 3.14159
                    v4 -= v4
                    _ = fi14
                    _ = v15
                }
                _ = v13
            }
        }
        switch (fi11 * 0) {
        case -1:
            v16 := []int{}
            _ = v16
        }
        _ = fi11
        _ = v12
    }
    _ = v10
    _ = v4
    return 255
}

func main() {
    var v1 uint = 414
    var v2 uint16 = 939
    v3 := make(map[string]bool)
    if true {
        switch "\n\t" {
        case "\n\t":
            var v17 uint64 = 1
            _ = v17
        case "":
            if true {
                fmt.Println(v2)
                for fi18 := 0; fi18 < 1; fi18++ {
                    v19 := GenS3{F0: 127}
                    v20 := []int{(fi18 * fi18)}
                    var v21 bool = true
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                }
                if false {
                    v2 = v2
                }
            } else {
                if false {
                    if (("\n\t" + "\n\t") >= ("" + "")) {
                        if false {
                            v1 -= genFn0()
                            if (112 <= v2) {
                                switch ("hello" + (("a" + ("foo bar" + "a")) + "a")) {
                                case "hello":
                                    if false {
                                        switch 128 {
                                        case 0:
                                            v22 := v3[("hello" + "foo bar")]
                                            _ = v22
                                        case -1:
                                            if true {
                                                v23 := []int{1, 9223372036854775807}
                                                v24 := GenS2{F0: -2147483648}
                                                v2 += 157
                                                _ = v23
                                                _ = v24
                                            } else {
                                                switch "" {
                                                case "foo bar":
                                                    var v25 uint16 = v2
                                                    _ = v25
                                                case "a":
                                                    v2 += v2
                                                case "hello":
                                                    v2 = 42
                                                }
                                                v26 := GenS3{F0: 0}
                                                v2 = v2
                                                _ = v26
                                            }
                                        case 4294967295:
                                            var v27 int = 9223372036854775807
                                            _ = v27
                                        }
                                        switch 9223372036854775807 {
                                        case 889:
                                            fmt.Println(v2, v2)
                                        case 350:
                                            if false {
                                                if true {
                                                    v1 = genFn0()
                                                    v28 := GenS4{F0: 0, F1: 1}
                                                    _ = v28
                                                } else {
                                                    var v29 uint = (v1 | (v1 + (v1 >> ((v1) & 63))))
                                                    v30 := GenS5{F0: 127, F1: -108, F2: uint(v2)}
                                                    var v31 int64 = 9223372036854775807
                                                    _ = v29
                                                    _ = v30
                                                    _ = v31
                                                }
                                            }
                                        case 4294967295:
                                            v1 = v1
                                        }
                                        switch -2147483648 {
                                        case 462:
                                            v2 = (v2 + (v2 + v2))
                                        case -2147483648:
                                            if true {
                                                fmt.Println(v2)
                                                var v32 int16 = -1
                                                if true {
                                                    var v33 int32 = 42
                                                    var v34 int = 32768
                                                    for fi35 := 0; fi35 < 15; fi35++ {
                                                        switch (v34 << (((^(v1))) & 63)) {
                                                        case -2147483648:
                                                            switch "\n\t" {
                                                            case "a":
                                                                v34 = 127
                                                            case "\n\t":
                                                                if (false && true) {
                                                                    switch fi35 {
                                                                    case 4294967295:
                                                                        v1 = 65535
                                                                    default:
                                                                        var v36 int = (fi35 | (fi35 << ((v1) & 63)))
                                                                        _ = v36
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                var v37 int16 = v32
                                                                _ = v37
                                                            }
                                                        default:
                                                            if false {
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    if (false && (0 > 4294967295)) {
                                                                        v34 = 316
                                                                    }
                                                                case "":
                                                                    var v38 int8 = 1
                                                                    _ = v38
                                                                default:
                                                                    for fi39 := 0; fi39 < 3; fi39++ {
                                                                        switch "a" {
                                                                        case "\n\t":
                                                                            v40 := GenS2{F0: 32767}
                                                                            _ = v40
                                                                        case "foo bar":
                                                                            v32 -= (v32 - 1)
                                                                        case "a":
                                                                            var v41 uint8 = 127
                                                                            _ = v41
                                                                        default:
                                                                            fmt.Println(fi39)
                                                                        }
                                                                        _ = fi39
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        var v42 string = ("\n\t" + ("" + (("" + "hello") + "a")))
                                                        _ = fi35
                                                        _ = v42
                                                    }
                                                    _ = v33
                                                    _ = v34
                                                }
                                                _ = v32
                                            } else {
                                                for fi43 := 0; fi43 < 2; fi43++ {
                                                    fmt.Println(v2, v2)
                                                    if false {
                                                        for fi44 := 0; fi44 < 19; fi44++ {
                                                            var v45 uint8 = 207
                                                            v46 := GenS6{F0: 32}
                                                            fmt.Println(v1)
                                                            _ = fi44
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                        for fi47 := 0; fi47 < 12; fi47++ {
                                                            fmt.Println(v1, v1)
                                                            switch "hello" {
                                                            case "a":
                                                                v3[(("hello" + "a") + ((((("foo bar" + "a") + "\n\t") + "hello") + (("\n\t" + (("\n\t" + ("" + "")) + "hello")) + ("\n\t" + ("foo bar" + "foo bar")))) + ("\n\t" + "hello")))] = (uint16(v1) > 771)
                                                            case "hello":
                                                                var v48 int16 = 1000
                                                                _ = v48
                                                            case "foo bar":
                                                                v2 = uint16(v1)
                                                            }
                                                            v2 = (v2 * (v2 | 65535))
                                                            _ = fi47
                                                        }
                                                    }
                                                    _ = fi43
                                                }
                                                v1 += v1
                                            }
                                        }
                                    }
                                }
                                var v49 int = 42
                                for fi50 := 0; fi50 < 16; fi50++ {
                                    for fi51 := 0; fi51 < 17; fi51++ {
                                        var v52 uint32 = 65535
                                        var v53 int = 2147483647
                                        _ = fi51
                                        _ = v52
                                        _ = v53
                                    }
                                    _ = fi50
                                }
                                _ = v49
                            }
                        }
                        fmt.Println(v1, v2)
                        v54 := GenS6{F0: 2}
                        _ = v54
                    } else {
                        if true {
                            if false {
                                if (uint32(v1) < 2147483647) {
                                    if (!(-887 != -1)) {
                                        if true {
                                            v2 += (^(v2))
                                            switch 299 {
                                            case 32768:
                                                for fi55 := 0; fi55 < 12; fi55++ {
                                                    switch int(v1) {
                                                    case -656:
                                                        switch "" {
                                                        case "":
                                                            var v56 float32 = float32(float64((0.0 + 1e10)))
                                                            _ = v56
                                                        }
                                                    }
                                                    var v57 int64 = -9223372036854775808
                                                    _ = fi55
                                                    _ = v57
                                                }
                                            }
                                        }
                                        if true {
                                            fmt.Println(v1, v1)
                                        }
                                        for fi58 := 0; fi58 < 6; fi58++ {
                                            var v59 uint = 1
                                            _ = fi58
                                            _ = v59
                                        }
                                    }
                                    v1 += v1
                                } else {
                                    switch "foo bar" {
                                    case "a":
                                        v3["foo bar"] = false
                                    }
                                }
                            } else {
                                var v60 uint = 2
                                v60 = 1
                                var v61 uint32 = 580
                                _ = v60
                                _ = v61
                            }
                        } else {
                            if ((v2 ^ (v2 / (((v2 - v2)) | 1))) == (v2 | v2)) {
                                v1 = v1
                                for fi62 := 0; fi62 < 5; fi62++ {
                                    switch 328 {
                                    case -1:
                                        for fi63 := 0; fi63 < 6; fi63++ {
                                            fmt.Println(fi62)
                                            v2 = v2
                                            _ = fi63
                                        }
                                    case -2147483648:
                                        if (18446744073709551615 >= 436) {
                                            switch (fi62 | fi62) {
                                            case 127:
                                                var v64 int32 = 0
                                                _ = v64
                                            case 1:
                                                switch fi62 {
                                                case 2147483647:
                                                    v65 := GenS7{F0: false, F1: 585, F2: (v1 % (((v1 + (v1 * v1))) | 1))}
                                                    _ = v65
                                                case 32768:
                                                    for fi66 := 0; fi66 < 17; fi66++ {
                                                        v67 := []int{42, fi66}
                                                        v68 := GenS8{F0: (^(v2))}
                                                        _ = fi66
                                                        _ = v67
                                                        _ = v68
                                                    }
                                                default:
                                                    v2 -= (v2 | v2)
                                                }
                                            default:
                                                v69 := []int{32767}
                                                _ = v69
                                            }
                                        } else {
                                            switch 1 {
                                            case -522:
                                                v70 := GenS2{F0: 2147483647}
                                                _ = v70
                                            }
                                            v71 := GenS6{F0: 1}
                                            _ = v71
                                        }
                                    case 4294967295:
                                        v72 := GenS9{F0: 0}
                                        _ = v72
                                    }
                                    if (-32768 <= 1000) {
                                        for fi73 := 0; fi73 < 8; fi73++ {
                                            v1 -= (v1 / ((v1) | 1))
                                            var v74 uint64 = uint64(v2)
                                            _ = fi73
                                            _ = v74
                                        }
                                        for fi75 := 0; fi75 < 7; fi75++ {
                                            v3[(("a" + ("foo bar" + "hello")) + "a")] = true
                                            var v76 string = "hello"
                                            v2 = 294
                                            _ = fi75
                                            _ = v76
                                        }
                                    } else {
                                        v1 = v1
                                        var v77 float32 = (0.1 * 0.0)
                                        switch (fi62 % ((42) | 1)) {
                                        case -553:
                                            fmt.Println(v1, v2, v77)
                                        default:
                                            v77 = v77
                                        }
                                        _ = v77
                                    }
                                    var v78 uint8 = 42
                                    _ = fi62
                                    _ = v78
                                }
                                if false {
                                    var v79 bool = false
                                    v80 := []int{-2147483648, -1, 957, 32768}
                                    _ = v79
                                    _ = v80
                                }
                            }
                        }
                        fmt.Println(v1, v2)
                        for fi81 := 0; fi81 < 19; fi81++ {
                            v82 := v3[("hello" + "foo bar")]
                            var v83 int64 = -735
                            _ = fi81
                            _ = v82
                            _ = v83
                        }
                    }
                    for fi84 := 0; fi84 < 17; fi84++ {
                        for fi85 := 0; fi85 < 9; fi85++ {
                            var v86 uint8 = 127
                            var v87 int32 = 153
                            _ = fi85
                            _ = v86
                            _ = v87
                        }
                        for fi88 := 0; fi88 < 3; fi88++ {
                            v1 = v1
                            v1 = (v1 + v1)
                            if true {
                                v2 = (v2 >> ((v1) & 63))
                                fmt.Println(v2, fi88)
                                v89 := GenS10{F0: fi84, F1: -128, F2: 0}
                                fmt.Println(v2, fi88)
                                _ = v89
                            }
                            switch ("a" + (("\n\t" + "\n\t") + "\n\t")) {
                            case "foo bar":
                                var v90 int8 = 96
                                _ = v90
                            }
                            var v91 uint8 = 127
                            _ = fi88
                            _ = v91
                        }
                        v2 += (v2 & (v2 | (^(v2))))
                        _ = fi84
                    }
                } else {
                    fmt.Println(v2, v2)
                }
            }
        }
        v1 = v1
    } else {
        v92 := []int{32767}
        if true {
            v93 := GenS7{F0: ("\n\t" != ("\n\t" + "a")), F1: -1, F2: v1}
            _ = v93
        }
        _ = v92
    }
    fmt.Println(v1, v2)
    v94 := GenS11{F0: 42}
    for fi95 := 0; fi95 < 8; fi95++ {
        var v96 uint64 = 18446744073709551615
        switch (fi95 | fi95) {
        case 32768:
            v96 = v96
        case 2147483647:
            for fi97 := 0; fi97 < 2; fi97++ {
                if false {
                    v98 := GenS12{F0: -1, F1: 566, F2: -1}
                    _ = v98
                }
                v96 += 245
                _ = fi97
            }
        default:
            v1 -= (v1 << ((255) & 63))
        }
        _ = fi95
        _ = v96
    }
    v2 -= (v2 ^ 42)
    v1 -= v1
    if (18446744073709551615 <= 2) {
        var v99 int32 = 0
        switch 32767 {
        case -26:
            if (4294967295 <= 708) {
                if false {
                    v3["\n\t"] = (5 != -1)
                }
            }
        case 9223372036854775807:
            if true {
                for fi100 := 0; fi100 < 3; fi100++ {
                    var v101 bool = (!false)
                    _ = fi100
                    _ = v101
                }
                v102 := GenS13{F0: 42, F1: 9223372036854775807}
                var v103 float32 = 1.5
                _ = v102
                _ = v103
            }
        case -631:
            var v104 int16 = -1
            _ = v104
        }
        _ = v99
    } else {
        v2 = (v2 - (^(v2)))
        for fi105 := 0; fi105 < 12; fi105++ {
            var v106 uint32 = 629
            v2 -= (v2 - v2)
            _ = fi105
            _ = v106
        }
    }
    fmt.Println(v2)
    v2 -= 32767
    v1 += v1
    if (((float64(-2.25) - 0.0) / 1.0) > 1e10) {
        switch "\n\t" {
        case "\n\t":
            v1 = v1
        }
        var v107 string = ("a" + ((((("foo bar" + "") + "\n\t") + ("foo bar" + ("" + ("foo bar" + "hello")))) + ((("foo bar" + ("a" + "")) + "\n\t") + "")) + "foo bar"))
        var v108 int8 = 127
        _ = v107
        _ = v108
    }
    v109 := v3["a"]
    var v110 uint32 = 612
    switch 1 {
    case 1:
        var v111 int64 = 876
        _ = v111
    case 4294967295:
        var v112 float32 = 1.5
        _ = v112
    }
    v113 := GenS14{F0: "", F1: -1, F2: -1}
    v109 = v109
    v114 := GenS15{F0: (^(v2)), F1: "\n\t", F2: uint(v2)}
    fmt.Println("v1:", v1)
    fmt.Println("v109:", v109)
    fmt.Println("v110:", v110)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v109
    _ = v110
    _ = v113
    _ = v114
    _ = v2
    _ = v3
    _ = v94
}
