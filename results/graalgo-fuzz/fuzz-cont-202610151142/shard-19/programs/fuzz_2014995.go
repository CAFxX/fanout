package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 int32; F1 string; F2 uint8 }
type GenS6 struct { F0 int8; F1 uint; F2 uint8 }
type GenS7 struct { F0 int64; F1 uint16; F2 int32 }
type GenS8 struct { F0 int64; F1 int32; F2 string }
type GenS9 struct { F0 int8; F1 string }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 bool; F1 int64 }
type GenS12 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS13 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS14 struct { F0 uint32; F1 uint64 }

func main() {
    var v1 int8 = 127
    var v2 uint8 = 2
    var v3 uint32 = 501
    v4 := []string{}
    v5 := []uint32{v3}
    if (65535 > 2) {
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
        v6 := GenS1{F0: v1, F1: v2, F2: 2147483647}
        var v7 uint16 = 42
        _ = v6
        _ = v7
    } else {
        var v8 int = 4294967295
        _ = v8
    }
    var v9 uint16 = 2
    if ((v1 | (v1 + v1)) == (v1 % ((v1) | 1))) {
        if (877 > 4294967295) {
            v2 += v2
        }
        v2 += v2
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        v9 = 255
        _ = fi10
    }
    var v11 float32 = (float32(0.0) * 1.5)
    var v12 uint32 = v3
    v1 += int8(v2)
    v13 := GenS2{F0: (^(v9)), F1: ((v2 / ((v2) | 1)) != uint8(v12))}
    fmt.Println(v12, v3)
    v3 += 4294967295
    v12 = v3
    for fi14 := 0; fi14 < 1; fi14++ {
        if true {
            var v15 uint16 = (^(v9))
            _ = v15
        } else {
            if false {
                var v16 float64 = 1.5
                v2 = v2
                v12 -= 1
                _ = v16
            } else {
                v3 = (v12 >> ((255) & 63))
                v9 -= v9
            }
        }
        switch "" {
        case "\n\t":
            for fi17 := 0; fi17 < 7; fi17++ {
                v18 := GenS1{F0: (v1 >> ((2) & 63)), F1: (v2 ^ v2), F2: 662}
                _ = fi17
                _ = v18
            }
        case "hello":
            v11 = 1.5
        case "":
            _ = len(v5)
        }
        _ = fi14
    }
    switch "foo bar" {
    case "\n\t":
        var v19 uint64 = 0
        _ = v19
    case "a":
        for fi20 := 0; fi20 < 5; fi20++ {
            v12 = 513
            switch 0 {
            case -1:
                for fi21 := 0; fi21 < 17; fi21++ {
                    v9 = (v9 / ((v9) | 1))
                    _ = fi21
                }
            default:
                var v22 uint32 = v12
                _ = v22
            }
            _ = fi20
        }
    default:
        for fi23 := 0; fi23 < 2; fi23++ {
            v12 = 158
            _ = fi23
        }
    }
    v24 := GenS1{F0: v1, F1: (v2 - (v2 - 42)), F2: -2147483648}
    var v25 int8 = (v1 + v1)
    for fi26 := 0; fi26 < 2; fi26++ {
        if false {
            fmt.Println(v3, v2, v11)
            switch ("foo bar" + "\n\t") {
            case "hello":
                v3 = 2147483648
            case "a":
                var v27 float32 = v11
                _ = v27
            default:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            for fi28 := 0; fi28 < 8; fi28++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v5 = append(v5, (v3 & v3))
                if ((^(v9)) >= v9) {
                    v12 = v3
                    v3 += (^(v3))
                    v29 := GenS2{F0: (v9 / ((65535) | 1)), F1: (false && (v11 < v11))}
                    if ((1.5 - v11) < (v11 / ((v11 - 0.0) - -2.25))) {
                        switch "\n\t" {
                        case "":
                            v11 += v11
                        case "foo bar":
                            switch (-(fi26)) {
                            case 95:
                                switch "foo bar" {
                                case "\n\t":
                                    v2 = v2
                                default:
                                    var v30 float64 = 0.0
                                    _ = v30
                                }
                            case 0:
                                if true {
                                    if ((float32((3.14159 - 1.5)) / 1.0) <= -2.25) {
                                        if (float64(float32(1.5)) < 3.14159) {
                                            switch (-(fi26)) {
                                            case 727:
                                                if false {
                                                    v31 := GenS3{F0: 2147483647}
                                                    v11 += float32(3.14159)
                                                    _ = v31
                                                } else {
                                                    var v32 int8 = (v25 & (v1 - int8(v9)))
                                                    _ = len(v4)
                                                    _ = v32
                                                }
                                            default:
                                                if ((v2 - 1) <= (v2 + uint8(v9))) {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    if len(v4) > 0 {
                                                        v4[0] = ""
                                                    }
                                                } else {
                                                    var v33 float64 = 1.5
                                                    v3 -= v3
                                                    if true {
                                                        v34 := GenS2{F0: 335, F1: (-2.25 <= (3.14159 / (0.1 - float32((3.14159 * 1e10)))))}
                                                        var v35 uint64 = 18446744073709551615
                                                        v12 -= 0
                                                        if true {
                                                            var v36 int64 = -9223372036854775808
                                                            var v37 uint = 18446744073709551615
                                                            fmt.Println(v35, v3, v2)
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                        v33 -= 1.5
                                                        v1 -= int8(v2)
                                                        var v38 bool = false
                                                        _ = v34
                                                        _ = v35
                                                        _ = v38
                                                    } else {
                                                        v33 = (v33 + (v33 * 1.5))
                                                        var v39 int16 = 1000
                                                        v9 = 2
                                                        v9 = (v9 | 65535)
                                                        var v40 uint8 = (v2 + 0)
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        v4 = append(v4, "hello")
                                                        _ = v39
                                                        _ = v40
                                                    }
                                                    var v41 float32 = (v11 + 0.1)
                                                    v3 += v12
                                                    _ = v33
                                                    _ = v41
                                                }
                                            }
                                            if false {
                                                var v42 bool = false
                                                _ = v42
                                            } else {
                                                if false {
                                                    v43 := GenS4{F0: 0, F1: fi28}
                                                    if len(v4) > 0 {
                                                        v4[0] = "hello"
                                                    }
                                                    if false {
                                                        v2 += v2
                                                        v11 = (v11 * (float32(3.14159) * ((0.0 * v11) - (v11 + (v11 - 0.0)))))
                                                        fmt.Println(v12, v3, v9)
                                                    } else {
                                                        if false {
                                                            if ((255 >= (v2 * (v2 / ((1) | 1)))) || true) {
                                                                v44 := GenS2{F0: 42, F1: true}
                                                                if len(v5) > 0 {
                                                                    v5[0] = (v12 % (((^(v12))) | 1))
                                                                }
                                                                _ = v44
                                                            }
                                                            switch ("a" + "") {
                                                            case "foo bar":
                                                                var v45 int16 = -32768
                                                                var v46 uint16 = v9
                                                                fmt.Println(v2)
                                                                v47 := GenS5{F0: 2147483647, F1: "", F2: v2}
                                                                _ = v45
                                                                _ = v46
                                                                _ = v47
                                                            case "\n\t":
                                                                fmt.Println(fi28, v12, v11)
                                                            case "a":
                                                                switch ("a" + ("hello" + "\n\t")) {
                                                                case "hello":
                                                                    if (v1 >= v25) {
                                                                        if true {
                                                                            if (!false) {
                                                                                if true {
                                                                                    if (2 < 2) {
                                                                                        v48 := GenS3{F0: 328}
                                                                                        v25 = -1
                                                                                        v49 := GenS1{F0: v25, F1: (v2 << ((4294967295) & 63)), F2: -237}
                                                                                        _ = v48
                                                                                        _ = v49
                                                                                    } else {
                                                                                        fmt.Println(v25, v1)
                                                                                    }
                                                                                    v25 -= (v25 ^ -53)
                                                                                } else {
                                                                                    v25 = (v1 | -128)
                                                                                    var v50 uint64 = 253
                                                                                    _ = v50
                                                                                }
                                                                                var v51 float32 = 0.1
                                                                                var v52 uint8 = 1
                                                                                _ = v51
                                                                                _ = v52
                                                                            }
                                                                            switch ("foo bar" + "foo bar") {
                                                                            case "a":
                                                                                var v53 int16 = -32768
                                                                                _ = v53
                                                                            case "\n\t":
                                                                                v2 -= v2
                                                                                v11 = v11
                                                                            case "foo bar":
                                                                                switch (^(fi26)) {
                                                                                case 4294967295:
                                                                                    var v54 uint16 = (v9 >> ((4294967295) & 63))
                                                                                    _ = v54
                                                                                }
                                                                            default:
                                                                                v5 = append(v5, 407)
                                                                            }
                                                                            switch ("foo bar" + "hello") {
                                                                            case "":
                                                                                switch (fi28 / (((fi26 - 4294967295)) | 1)) {
                                                                                case -465:
                                                                                    v55 := GenS6{F0: v1, F1: 919, F2: v2}
                                                                                    _ = v55
                                                                                case 786:
                                                                                    v2 = uint8(v3)
                                                                                case -569:
                                                                                    switch ("\n\t" + "\n\t") {
                                                                                    case "":
                                                                                        v11 += (v11 - v11)
                                                                                    case "a":
                                                                                        switch "" {
                                                                                        case "hello":
                                                                                            fmt.Println(v1, v2)
                                                                                        case "\n\t":
                                                                                            v56 := GenS3{F0: 1}
                                                                                            _ = v56
                                                                                        case "":
                                                                                            switch 128 {
                                                                                            case 32768:
                                                                                                fmt.Println(v3, fi28, fi26)
                                                                                            case 128:
                                                                                                var v57 float32 = (0.1 * (float32(1e10) * v11))
                                                                                                _ = v57
                                                                                            case 4294967295:
                                                                                                if (!(128 == v2)) {
                                                                                                    var v58 float64 = 1e10
                                                                                                    v59 := GenS2{F0: v9, F1: true}
                                                                                                    _ = v58
                                                                                                    _ = v59
                                                                                                } else {
                                                                                                    fmt.Println(v11, fi26)
                                                                                                    fmt.Println(v11, v3)
                                                                                                    v11 = v11
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            var v60 int8 = -128
                                                                                            _ = v60
                                                                                        }
                                                                                    case "hello":
                                                                                        var v61 int8 = (-(v25))
                                                                                        _ = v61
                                                                                    }
                                                                                }
                                                                            case "a":
                                                                                switch fi26 {
                                                                                case -1:
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = (v3 - v3)
                                                                                    }
                                                                                case 4294967295:
                                                                                    v12 = (v3 & 4294967295)
                                                                                case 9223372036854775807:
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = "a"
                                                                                    }
                                                                                    if true {
                                                                                        switch (fi26 & (fi28 + fi26)) {
                                                                                        case 1:
                                                                                            switch ("\n\t" + "hello") {
                                                                                            case "":
                                                                                                if len(v4) >= 2 {
                                                                                                    v4 = v4[1:len(v4)]
                                                                                                }
                                                                                            case "foo bar":
                                                                                                var v62 int16 = 32767
                                                                                                _ = v62
                                                                                            case "\n\t":
                                                                                                var v63 string = (("a" + (((("foo bar" + ("a" + "a")) + ("" + "")) + "foo bar") + ("" + "\n\t"))) + "hello")
                                                                                                _ = v63
                                                                                            }
                                                                                        case 4294967295:
                                                                                            switch fi26 {
                                                                                            case 127:
                                                                                                v64 := GenS7{F0: 42, F1: 507, F2: 638}
                                                                                                _ = v64
                                                                                            case 9223372036854775807:
                                                                                                if (!true) {
                                                                                                    _ = len(v4)
                                                                                                    switch "a" {
                                                                                                    case "":
                                                                                                        v3 = 2147483647
                                                                                                    }
                                                                                                    var v65 uint8 = 255
                                                                                                    fmt.Println(v12, fi28, v2)
                                                                                                    v12 = v12
                                                                                                    _ = v65
                                                                                                }
                                                                                            case -2147483648:
                                                                                                var v66 uint8 = 2
                                                                                                _ = v66
                                                                                            default:
                                                                                                var v67 int64 = -9223372036854775808
                                                                                                _ = v67
                                                                                            }
                                                                                        default:
                                                                                            if ((v12 + v12) == 65535) {
                                                                                                v25 += (v1 % (((v25 * v25)) | 1))
                                                                                                if false {
                                                                                                    v1 = (-(v1))
                                                                                                    var v68 uint32 = v3
                                                                                                    if len(v4) > 0 {
                                                                                                        v4[0] = ("hello" + "\n\t")
                                                                                                    }
                                                                                                    _ = v68
                                                                                                } else {
                                                                                                    v4 = append(v4, "")
                                                                                                    var v69 uint16 = 65535
                                                                                                    if len(v4) >= 2 {
                                                                                                        v4 = v4[1:len(v4)]
                                                                                                    }
                                                                                                    _ = v69
                                                                                                }
                                                                                                var v70 uint = 341
                                                                                                if (uint16(v70) <= (v9 ^ v9)) {
                                                                                                    v71 := GenS8{F0: 220, F1: -2147483648, F2: ("foo bar" + "\n\t")}
                                                                                                    _ = v71
                                                                                                } else {
                                                                                                    v70 = (v70 - (v70 + 2))
                                                                                                    v70 = 65535
                                                                                                    if false {
                                                                                                        v9 = (v9 ^ v9)
                                                                                                    }
                                                                                                }
                                                                                                _ = v70
                                                                                            } else {
                                                                                                if false {
                                                                                                    v25 = 70
                                                                                                    var v72 uint32 = (v3 >> ((746) & 63))
                                                                                                    _ = v72
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        v73 := GenS6{F0: 127, F1: 50, F2: 0}
                                                                                        _ = v73
                                                                                    }
                                                                                    var v74 int16 = -946
                                                                                    _ = v74
                                                                                }
                                                                            }
                                                                            v75 := GenS9{F0: v25, F1: ""}
                                                                            switch (-(fi28)) {
                                                                            case -211:
                                                                                v2 = (v2 | 200)
                                                                            case 1:
                                                                                v25 -= v1
                                                                            case -1:
                                                                                v2 += v2
                                                                                fmt.Println(v25, v25)
                                                                            default:
                                                                                fmt.Println(v3, v12, fi28)
                                                                            }
                                                                            _ = v75
                                                                        }
                                                                        if ((!(0.0 >= v11)) || false) {
                                                                            if (62 <= 4294967295) {
                                                                                fmt.Println(v3, v9)
                                                                            } else {
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                            }
                                                                        }
                                                                        var v76 uint16 = (^(v9))
                                                                        _ = v76
                                                                    } else {
                                                                        var v77 uint64 = 2
                                                                        if true {
                                                                            v5 = append(v5, (v12 & (v12 & (v3 / (((^(v3))) | 1)))))
                                                                            var v78 int8 = v1
                                                                            _ = v78
                                                                        } else {
                                                                            v9 = v9
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                        }
                                                                        _ = v77
                                                                    }
                                                                case "foo bar":
                                                                    v4 = append(v4, "hello")
                                                                case "a":
                                                                    var v79 int = fi28
                                                                    _ = v79
                                                                }
                                                            default:
                                                                v3 = 2147483647
                                                            }
                                                        }
                                                        var v80 uint64 = 1
                                                        _ = v80
                                                    }
                                                    if true {
                                                        v81 := GenS10{F0: 1000}
                                                        if false {
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                        } else {
                                                            if (((v1 << ((812) & 63)) >= 124) || ((v3 / ((v12) | 1)) > 2147483647)) {
                                                                v82 := GenS11{F0: ((fi28 + fi26) >= (^(fi28))), F1: 801}
                                                                v12 -= (v12 + (v3 | v12))
                                                                v2 = uint8(v12)
                                                                fmt.Println(v3, v11)
                                                                _ = v82
                                                            }
                                                            var v83 float64 = (-2.25 - 0.1)
                                                            v5 = append(v5, 2147483647)
                                                            var v84 string = ("a" + "foo bar")
                                                            v3 = (v12 - v12)
                                                            _ = v83
                                                            _ = v84
                                                        }
                                                        v12 += uint32(v2)
                                                        _ = v81
                                                    }
                                                    switch fi28 {
                                                    case 4294967295:
                                                        var v85 uint64 = 42
                                                        _ = v85
                                                    case -987:
                                                        var v86 int64 = 9223372036854775807
                                                        _ = v86
                                                    case 272:
                                                        var v87 uint = 42
                                                        _ = v87
                                                    }
                                                    var v88 int8 = v25
                                                    _ = v43
                                                    _ = v88
                                                }
                                            }
                                        } else {
                                            v9 = 2
                                        }
                                        var v89 int8 = v25
                                        v25 = v25
                                        _ = v89
                                    }
                                    var v90 float64 = (float64(v11) / 1.0)
                                    _ = v90
                                } else {
                                    var v91 uint = 125
                                    var v92 int8 = (v1 / ((127) | 1))
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    _ = v91
                                    _ = v92
                                }
                            case 179:
                                switch (("foo bar" + "a") + ("\n\t" + "a")) {
                                case "a":
                                    var v93 uint8 = uint8(v3)
                                    _ = v93
                                case "foo bar":
                                    var v94 float64 = 1.5
                                    _ = v94
                                case "hello":
                                    if ((v9 - v9) > v9) {
                                        if false {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            v9 = (v9 & (v9 * 32767))
                                        } else {
                                            v9 += 0
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                        v95 := GenS3{F0: -1}
                                        fmt.Println(v11, v1)
                                        _ = v95
                                    } else {
                                        v2 -= (v2 & (v2 * 63))
                                        v96 := GenS2{F0: (v9 | v9), F1: false}
                                        if len(v4) > 0 {
                                            v4[0] = ""
                                        }
                                        _ = len(v4)
                                        v97 := GenS6{F0: 100, F1: 0, F2: v2}
                                        _ = v96
                                        _ = v97
                                    }
                                default:
                                    switch (fi28 & -911) {
                                    case 403:
                                        switch ("foo bar" + "\n\t") {
                                        case "hello":
                                            v11 += (3.14159 - 0.1)
                                        case "\n\t":
                                            fmt.Println(fi28, v1, v1)
                                        default:
                                            var v98 int32 = 42
                                            _ = v98
                                        }
                                    case 520:
                                        if len(v5) > 0 {
                                            v5[0] = 687
                                        }
                                        v1 += 0
                                        v3 = 65535
                                        if len(v4) > 0 {
                                            v4[0] = ("\n\t" + ("a" + "hello"))
                                        }
                                    case 9223372036854775807:
                                        v99 := GenS12{F0: (v9 + 1), F1: (v1 >> ((56) & 63)), F2: v3}
                                        _ = v99
                                    default:
                                        switch ("" + ("hello" + "a")) {
                                        case "hello":
                                            _ = len(v4)
                                        case "foo bar":
                                            if true {
                                                v5 = append(v5, 824)
                                                if false {
                                                    v3 -= v3
                                                    v11 = 1.5
                                                    var v100 uint8 = v2
                                                    fmt.Println(v2)
                                                    _ = v100
                                                } else {
                                                    var v101 int = int(v12)
                                                    v102 := GenS1{F0: (v1 + (v25 - v1)), F1: (^(v2)), F2: -715}
                                                    v1 = v1
                                                    _ = v101
                                                    _ = v102
                                                }
                                            }
                                        case "":
                                            var v103 string = "\n\t"
                                            _ = v103
                                        }
                                    }
                                }
                            }
                        default:
                            if (true || ((!true) || (1 > 42))) {
                                if (236 > 508) {
                                    v104 := GenS9{F0: v1, F1: "\n\t"}
                                    var v105 uint32 = (^(v12))
                                    _ = v104
                                    _ = v105
                                }
                            } else {
                                v1 -= (v1 >> ((609) & 63))
                                if (685 >= (v12 + (v3 >> ((1) & 63)))) {
                                    fmt.Println(v9)
                                }
                                var v106 uint64 = 0
                                v107 := GenS10{F0: 1000}
                                _ = v106
                                _ = v107
                            }
                        }
                        v1 = (v1 >> ((255) & 63))
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    _ = v29
                } else {
                    switch (^(fi28)) {
                    case 128:
                        fmt.Println(v11, v2)
                    default:
                        var v108 bool = false
                        _ = v108
                    }
                }
                _ = fi28
            }
        } else {
            v25 = (v1 & (v1 % ((v1) | 1)))
            v109 := GenS13{F0: 230, F1: 1, F2: (v9 * v9)}
            if true {
                var v110 uint = 2
                _ = v110
            } else {
                var v111 uint16 = v9
                _ = v111
            }
            _ = v109
        }
        v112 := GenS7{F0: 42, F1: (v9 ^ v9), F2: -2147483648}
        _ = fi26
        _ = v112
    }
    v113 := GenS13{F0: v12, F1: 980, F2: (v9 / (((v9 + v9)) | 1))}
    for fi114 := 0; fi114 < 2; fi114++ {
        v115 := GenS4{F0: -109, F1: 683}
        if ((-899 == 42) || (true || ("" < "a"))) {
            for fi116 := 0; fi116 < 15; fi116++ {
                var v117 string = ""
                v4 = append(v4, v117)
                _ = fi116
                _ = v117
            }
            switch -2147483648 {
            case 592:
                if false {
                    if true {
                        fmt.Println(v3)
                        v11 -= -2.25
                    } else {
                        switch (fi114 >> ((65535) & 63)) {
                        case 553:
                            v4 = append(v4, ("foo bar" + ("a" + "a")))
                        case 127:
                            for fi118 := 0; fi118 < 11; fi118++ {
                                v11 = v11
                                if (!(!true)) {
                                    if ((v9 | v9) != 42) {
                                        if (!false) {
                                            _ = len(v4)
                                            var v119 uint = 704
                                            _ = v119
                                        }
                                    } else {
                                        v25 -= v1
                                        fmt.Println(v12, v25, v3)
                                        switch (fi118 & (fi118 - fi114)) {
                                        case 4294967295:
                                            switch ("hello" + (("\n\t" + (("hello" + "a") + (("hello" + "foo bar") + ("hello" + ("hello" + "a"))))) + "foo bar")) {
                                            case "":
                                                var v120 uint = 255
                                                _ = v120
                                            case "hello":
                                                var v121 int16 = int16(fi114)
                                                _ = v121
                                            default:
                                                v11 += ((1.5 * (float32(float64((v11 / v11))) * -2.25)) * (float32(1.5) * (1.5 * -2.25)))
                                            }
                                        case -2147483648:
                                            v25 = (-(v1))
                                        case -247:
                                            v122 := GenS6{F0: v1, F1: 42, F2: v2}
                                            _ = v122
                                        }
                                    }
                                }
                                _ = fi118
                            }
                        }
                        if true {
                            v3 = v12
                            switch (((("\n\t" + (("" + "a") + "a")) + "foo bar") + "") + "foo bar") {
                            case "a":
                                for fi123 := 0; fi123 < 3; fi123++ {
                                    v124 := GenS3{F0: 2147483647}
                                    var v125 float64 = 0.1
                                    v12 -= 2147483647
                                    _ = fi123
                                    _ = v124
                                    _ = v125
                                }
                            case "":
                                v126 := GenS1{F0: (v1 - 127), F1: (v2 * (v2 << ((4294967295) & 63))), F2: 0}
                                _ = v126
                            case "hello":
                                v25 = v25
                            default:
                                if len(v5) > 0 {
                                    v5[0] = v3
                                }
                            }
                        }
                        if (true && true) {
                            var v127 uint8 = uint8(v9)
                            _ = v127
                        } else {
                            v11 -= (v11 * v11)
                            v4 = append(v4, "hello")
                        }
                    }
                    var v128 int8 = 127
                    _ = v128
                }
            case 32767:
                switch (fi114 % ((int(v9)) | 1)) {
                case 128:
                    if (fi114 == (fi114 & fi114)) {
                        v3 += 752
                    } else {
                        v129 := GenS14{F0: v3, F1: 2}
                        _ = v129
                    }
                }
            default:
                for fi130 := 0; fi130 < 8; fi130++ {
                    v1 = -18
                    v131 := GenS12{F0: 32767, F1: (^(v25)), F2: v3}
                    var v132 uint64 = 2
                    _ = fi130
                    _ = v131
                    _ = v132
                }
            }
        }
        var v133 uint16 = 345
        _ = fi114
        _ = v115
        _ = v133
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v113
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
