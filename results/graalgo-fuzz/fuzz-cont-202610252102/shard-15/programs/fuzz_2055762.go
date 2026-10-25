package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS5 struct { F0 uint32; F1 int }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint64; F1 uint }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int64; F1 int8; F2 int8 }
func genFn0(p0 uint64, p1 uint16, p2 int) string {
    switch "\n\t" {
    case "a":
        if true {
            var v4 float64 = 0.1
            fmt.Println(p2)
            _ = v4
        }
    case "\n\t":
        v5 := GenS1{F0: 3}
        _ = v5
    case "":
        var v6 uint = 1
        _ = v6
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 uint8 = 255
            v9 = (v9 + (v9 - 255))
            _ = fi8
            _ = v9
        }
        p1 = (p1 & 1)
        _ = fi7
    }
    fmt.Println(p1, p0, p0)
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return ""
}

func main() {
    var v1 uint = 65535
    var v2 int = 1
    v3 := make(map[string]int)
    if true {
        _ = len(v3)
        for fi10 := 0; fi10 < 14; fi10++ {
            fmt.Println(fi10, fi10, v1)
            _ = fi10
        }
    }
    v11 := v3[""]
    var v12 uint32 = 1
    if true {
        v13 := []int{v11, 1, 203}
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 int = (^(v2))
            _ = fi14
            _ = v15
        }
        if ((-2.25 + (((((3.14159 - (1.5 / (0.1 - 3.14159))) - 0.0) - 0.0) * 0.1) - (1e10 + 3.14159))) != ((1.5 - 0.1) * (0.1 / 3.14159))) {
            var v16 bool = (true || false)
            if v16 {
                var v17 int64 = 1
                v12 -= v12
                _ = v17
            } else {
                _ = len(v13)
                var v18 uint8 = 235
                var v19 int = v2
                _ = v18
                _ = v19
            }
            _ = v16
        }
        _ = v13
    }
    v20 := GenS1{F0: 89}
    switch "\n\t" {
    case "\n\t":
        var v21 float64 = -2.25
        _ = v21
    case "":
        v22 := GenS2{F0: (v1 / ((943) | 1)), F1: v1}
        _ = v22
    case "foo bar":
        if (!(!true)) {
            _ = len(v3)
        } else {
            if (v12 != (v12 | (^(v12)))) {
                v23 := GenS1{F0: 229}
                v11 = v11
                _ = v23
            } else {
                v24 := GenS3{F0: -1}
                var v25 uint32 = v12
                _ = v24
                _ = v25
            }
        }
    default:
        if false {
            for fi26 := 0; fi26 < 17; fi26++ {
                for fi27 := 0; fi27 < 12; fi27++ {
                    v28 := []int{32767}
                    _ = fi27
                    _ = v28
                }
                v2 = v11
                _ = fi26
            }
            fmt.Println(v1)
        } else {
            v12 -= v12
            v2 -= (v2 % (((v2 << (((v1 / (((v1 / ((255) | 1))) | 1))) & 63))) | 1))
        }
    }
    if (!(true && (0 > 91))) {
        for fi29 := 0; fi29 < 4; fi29++ {
            if ((!("a" < "hello")) && (v1 == 823)) {
                delete(v3, "\n\t")
                if (false || true) {
                    var v30 uint32 = (v12 & (v12 / ((65535) | 1)))
                    v31 := GenS4{F0: 248, F1: 255, F2: 0}
                    for fi32 := 0; fi32 < 16; fi32++ {
                        switch 9223372036854775807 {
                        case 42:
                            v30 = v30
                        }
                        switch "" {
                        case "a":
                            v3[""] = fi32
                        case "hello":
                            v30 = (v12 ^ (v12 >> ((v1) & 63)))
                        case "\n\t":
                            if true {
                                var v33 uint = 1
                                v1 = v33
                                _ = v33
                            }
                        }
                        var v34 int8 = 1
                        _ = fi32
                        _ = v34
                    }
                    _ = v30
                    _ = v31
                } else {
                    var v35 int64 = 9223372036854775807
                    fmt.Println(v1, fi29, v1)
                    for fi36 := 0; fi36 < 12; fi36++ {
                        v11 = (v11 / ((0) | 1))
                        v12 = (v12 + 42)
                        _ = fi36
                    }
                    _ = v35
                }
            } else {
                var v37 string = "\n\t"
                fmt.Println(v11, v2)
                if ((v37 + ("foo bar" + v37)) != v37) {
                    var v38 int64 = 9223372036854775807
                    v39 := v3[v37]
                    _ = v38
                    _ = v39
                } else {
                    for fi40 := 0; fi40 < 8; fi40++ {
                        v41 := GenS4{F0: 2, F1: 1, F2: 0}
                        v12 = 1
                        v12 = 4294967295
                        _ = fi40
                        _ = v41
                    }
                    if (-1 > 1) {
                        v12 = v12
                    }
                }
                _ = v37
            }
            switch (fi29 & (-(fi29))) {
            case 127:
                v11 = -998
            case 454:
                switch (("\n\t" + ((("foo bar" + (((("foo bar" + ("foo bar" + "")) + "\n\t") + "") + "\n\t")) + "") + "hello")) + "foo bar") {
                case "a":
                    v42 := GenS4{F0: 659, F1: 65535, F2: 42}
                    _ = v42
                case "hello":
                    v3[("hello" + ("foo bar" + ((("\n\t" + ("a" + "foo bar")) + "\n\t") + "foo bar")))] = (-(v11))
                }
            default:
                for fi43 := 0; fi43 < 11; fi43++ {
                    if true {
                        switch fi29 {
                        case 32767:
                            var v44 uint = v1
                            _ = v44
                        case 32768:
                            var v45 uint32 = v12
                            _ = v45
                        default:
                            if (-1 < 881) {
                                if true {
                                    if false {
                                        var v46 int64 = 0
                                        v47 := []int{-776}
                                        var v48 int64 = (v46 + 0)
                                        _ = v46
                                        _ = v47
                                        _ = v48
                                    } else {
                                        var v49 uint8 = 1
                                        if (2 != v49) {
                                            var v50 bool = false
                                            v12 += 4294967295
                                            v11 = (fi43 << ((437) & 63))
                                            v51 := []int{(fi43 & v2), v2, (fi43 >> (((v1 / ((v1) | 1))) & 63))}
                                            var v52 bool = v50
                                            _ = v50
                                            _ = v51
                                            _ = v52
                                        } else {
                                            if (false && (0.0 <= -2.25)) {
                                                _ = len(v3)
                                                if true {
                                                    v1 -= (v1 & (^(v1)))
                                                    var v53 int16 = -885
                                                    if (630 == -708) {
                                                        if false {
                                                            v3["foo bar"] = fi43
                                                        }
                                                    } else {
                                                        if ((v2 - (fi43 / (((v2 ^ int(v53))) | 1))) >= (v2 >> ((v1) & 63))) {
                                                            fmt.Println(fi43, v49, v11)
                                                            if (!false) {
                                                                v1 -= (v1 ^ 255)
                                                                v11 = v11
                                                            }
                                                            v49 -= v49
                                                            var v54 int8 = 44
                                                            v55 := []int{(fi29 - (v2 >> (((^(v1))) & 63))), v11}
                                                            _ = v54
                                                            _ = v55
                                                        } else {
                                                            v56 := []int{87, int(v1), (v11 - v11), (^(v11))}
                                                            v1 = (v1 / ((42) | 1))
                                                            if (true || false) {
                                                                v49 = (v49 | v49)
                                                                if (!true) {
                                                                    fmt.Println(v53)
                                                                }
                                                            }
                                                            var v57 string = "foo bar"
                                                            var v58 int16 = v53
                                                            switch v57 {
                                                            case "":
                                                                var v59 int16 = v58
                                                                _ = v59
                                                            case "hello":
                                                                var v60 int16 = 0
                                                                _ = v60
                                                            default:
                                                                switch (v57 + v57) {
                                                                case "\n\t":
                                                                    v3[""] = (fi29 ^ v2)
                                                                }
                                                            }
                                                            v61 := GenS1{F0: v49}
                                                            v58 += (v58 | v53)
                                                            var v62 int32 = 2147483647
                                                            _ = v56
                                                            _ = v57
                                                            _ = v58
                                                            _ = v61
                                                            _ = v62
                                                        }
                                                    }
                                                    var v63 int = v11
                                                    var v64 int64 = 42
                                                    _ = v53
                                                    _ = v63
                                                    _ = v64
                                                }
                                                var v65 float32 = 0.1
                                                v1 -= (v1 & 42)
                                                if false {
                                                    v66 := []int{fi29, (^(v11)), v2}
                                                    _ = v66
                                                } else {
                                                    v12 = v12
                                                }
                                                _ = v65
                                            }
                                            v67 := []int{}
                                            fmt.Println(v2, v1, fi29)
                                            _ = v67
                                        }
                                        _ = v49
                                    }
                                    fmt.Println(v12, fi29)
                                } else {
                                    if (v12 == 904) {
                                        switch "" {
                                        case "foo bar":
                                            var v68 float64 = -2.25
                                            _ = v68
                                        }
                                        switch "a" {
                                        case "":
                                            v69 := GenS2{F0: (v1 % ((v1) | 1)), F1: v1}
                                            switch "\n\t" {
                                            case "foo bar":
                                                v2 = int(v12)
                                            case "\n\t":
                                                v70 := []int{(v2 % ((v11) | 1)), -2147483648, v2}
                                                _ = v70
                                            case "":
                                                v71 := []int{}
                                                _ = v71
                                            }
                                            _ = v69
                                        case "hello":
                                            fmt.Println(v1, v1, v2)
                                        default:
                                            if ((!true) && true) {
                                                v11 = (v2 % (((fi43 & fi29)) | 1))
                                                v72 := GenS1{F0: 42}
                                                _ = v72
                                            }
                                        }
                                    }
                                    var v73 int = fi29
                                    _ = v73
                                }
                                var v74 int8 = 0
                                _ = v74
                            }
                        }
                        switch -915 {
                        case 1:
                            delete(v3, "foo bar")
                        case 32768:
                            _ = len(v3)
                        case 9223372036854775807:
                            switch "\n\t" {
                            case "a":
                                var v75 uint = (v1 + v1)
                                _ = v75
                            case "":
                                v11 = (v2 / ((fi29) | 1))
                            case "hello":
                                if (9223372036854775807 != -1) {
                                    var v76 bool = true
                                    if ((128 != uint8(v1)) && v76) {
                                        var v77 int8 = int8(v1)
                                        if ((v77 / ((v77) | 1)) != int8(v1)) {
                                            v12 = (v12 ^ (v12 % ((v12) | 1)))
                                            v77 += (v77 & v77)
                                        } else {
                                            v12 += v12
                                            fmt.Println(v1)
                                        }
                                        _ = v77
                                    }
                                    switch (fi43 - 9223372036854775807) {
                                    case 127:
                                        var v78 string = ((("a" + ("a" + ("hello" + "\n\t"))) + "a") + "foo bar")
                                        _ = v78
                                    }
                                    switch (fi29 * 42) {
                                    case 42:
                                        v12 += v12
                                        v79 := GenS5{F0: (^(v12)), F1: v11}
                                        v76 = v76
                                        _ = v79
                                    case -2147483648:
                                        v1 += v1
                                    default:
                                        var v80 uint32 = v12
                                        _ = v80
                                    }
                                    switch 9223372036854775807 {
                                    case -744:
                                        switch ("" + ("\n\t" + (("\n\t" + "") + "\n\t"))) {
                                        case "foo bar":
                                            v11 = (v2 / (((^(v2))) | 1))
                                        case "":
                                            fmt.Println(v12)
                                        case "\n\t":
                                            var v81 int = (v11 & 32767)
                                            _ = v81
                                        default:
                                            var v82 int64 = 265
                                            _ = v82
                                        }
                                    case 2147483647:
                                        switch (-(v11)) {
                                        case 9223372036854775807:
                                            fmt.Println(v76, v76, v11)
                                        default:
                                            if ("foo bar" >= "a") {
                                                var v83 float32 = float32(1e10)
                                                _ = v83
                                            }
                                        }
                                    case -478:
                                        v84 := GenS4{F0: 0, F1: 17, F2: 206}
                                        _ = v84
                                    default:
                                        fmt.Println(v1, v2)
                                    }
                                    _ = v76
                                }
                            default:
                                var v85 uint = 18446744073709551615
                                _ = v85
                            }
                        }
                    } else {
                        fmt.Println(v11, v11, fi29)
                    }
                    v86 := []int{fi29}
                    v11 -= v2
                    _ = fi43
                    _ = v86
                }
            }
            v87 := GenS1{F0: 234}
            _ = fi29
            _ = v87
        }
        for fi88 := 0; fi88 < 10; fi88++ {
            var v89 float64 = -2.25
            fmt.Println(v12)
            v12 += v12
            _ = fi88
            _ = v89
        }
    }
    var v90 float32 = float32((0.1 + -2.25))
    switch v11 {
    case -1:
        v2 = v11
    case -358:
        var v91 uint64 = 18446744073709551615
        _ = v91
    }
    v92 := GenS1{F0: 127}
    if (true && (true && (1 != -227))) {
        if false {
            if true {
                v93 := GenS1{F0: 42}
                v11 = v2
                v94 := GenS6{F0: 42, F1: 338}
                _ = v93
                _ = v94
            }
            for fi95 := 0; fi95 < 15; fi95++ {
                v96 := GenS6{F0: 3, F1: 42}
                _ = fi95
                _ = v96
            }
        }
        switch "" {
        case "hello":
            v97 := []int{v2, v2, v2}
            _ = v97
        case "":
            v90 += float32(3.14159)
        case "foo bar":
            fmt.Println(v1, v12)
        default:
            v90 += v90
        }
        fmt.Println(v11)
    } else {
        v98 := GenS7{F0: 117}
        for fi99 := 0; fi99 < 6; fi99++ {
            switch v2 {
            case 865:
                if true {
                    v100 := GenS7{F0: 127}
                    _ = v100
                } else {
                    var v101 int = (^(v11))
                    _ = v101
                }
            case 4294967295:
                v3["foo bar"] = v2
            case -1:
                if true {
                    switch v2 {
                    case 682:
                        var v102 int64 = 9223372036854775807
                        _ = v102
                    case 32768:
                        if (!((v12 << (((v1 - (v1 << (((v1 % (((^(v1))) | 1))) & 63)))) & 63)) == (v12 + (v12 - v12)))) {
                            v103 := GenS8{F0: 4294967295, F1: v1}
                            v104 := GenS5{F0: 2, F1: fi99}
                            for fi105 := 0; fi105 < 3; fi105++ {
                                v12 -= v12
                                _ = fi105
                            }
                            _ = v103
                            _ = v104
                        } else {
                            v1 += 0
                            for fi106 := 0; fi106 < 1; fi106++ {
                                var v107 uint = (v1 * (v1 | (v1 | 212)))
                                _ = fi106
                                _ = v107
                            }
                            var v108 uint16 = 65535
                            _ = v108
                        }
                    }
                    v109 := GenS9{F0: true}
                    switch (v2 + (fi99 >> ((v1) & 63))) {
                    case 127:
                        fmt.Println(v2, v90)
                    case -459:
                        for fi110 := 0; fi110 < 11; fi110++ {
                            var v111 bool = ((52 > 255) && false)
                            v112 := GenS10{F0: 263, F1: 1, F2: 127}
                            v90 -= float32(((-2.25 + 3.14159) * 0.1))
                            _ = fi110
                            _ = v111
                            _ = v112
                        }
                    }
                    _ = v109
                } else {
                    switch (v2 * (v11 & (v11 >> (((v1 & v1)) & 63)))) {
                    case 128:
                        v113 := GenS4{F0: 255, F1: 2, F2: 27}
                        _ = v113
                    case 13:
                        var v114 float64 = (0.1 / 1.0)
                        _ = v114
                    }
                    v1 = v1
                }
            default:
                var v115 bool = (!(true && false))
                _ = v115
            }
            _ = fi99
        }
        _ = v98
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v90:", v90)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v90
    _ = v92
}
