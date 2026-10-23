package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS4 struct { F0 bool; F1 int64; F2 string }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 uint64; F1 bool }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 bool; F1 bool; F2 uint }
type GenS9 struct { F0 uint16; F1 uint8 }
type GenS10 struct { F0 int; F1 int16; F2 uint8 }
type GenS11 struct { F0 uint16; F1 uint16; F2 bool }
type GenS12 struct { F0 int64; F1 uint; F2 uint8 }
type GenS13 struct { F0 uint32 }

func main() {
    var v1 int64 = 1
    v2 := []string{(("a" + ((("hello" + "") + "a") + "foo bar")) + "foo bar"), "a", "\n\t", "\n\t"}
    v3 := make(map[uint32]string)
    var v4 bool = true
    for fi5 := 0; fi5 < 5; fi5++ {
        if ((fi5 * (fi5 | fi5)) == fi5) {
            var v6 int32 = 0
            v3[2147483648] = "a"
            var v7 int16 = -1
            _ = v6
            _ = v7
        }
        var v8 uint = 2
        _ = fi5
        _ = v8
    }
    _ = len(v3)
    v1 = (-(v1))
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 int16 = 662
        switch (fi9 * fi9) {
        case 751:
            for fi11 := 0; fi11 < 1; fi11++ {
                v12 := GenS1{F0: 65535, F1: 124, F2: -1}
                v1 = v1
                _ = len(v3)
                _ = fi11
                _ = v12
            }
        case 4294967295:
            if v4 {
                if (v4 && v4) {
                    var v13 uint = 850
                    _ = v13
                } else {
                    fmt.Println(v4)
                    for fi14 := 0; fi14 < 3; fi14++ {
                        if (v4 && ((872 == 255) || (((0.0 - (1.5 / 1.0)) / float32((float64(1.5) / (float64(-2.25) - 0.1)))) == ((-2.25 + 0.0) * 1.5)))) {
                            _ = len(v3)
                        } else {
                            if ((v1 != (v1 / (((-(v1))) | 1))) && v4) {
                                if (v10 == v10) {
                                    var v15 uint64 = 511
                                    _ = v15
                                }
                                v10 += 1
                                v1 = (v1 + (v1 + v1))
                            }
                            if (-1 != -128) {
                                v10 += v10
                            } else {
                                _ = len(v3)
                                delete(v3, 0)
                                _ = len(v2)
                                switch "\n\t" {
                                case "hello":
                                    var v16 uint64 = 1
                                    _ = v16
                                default:
                                    var v17 bool = (v1 < 42)
                                    _ = v17
                                }
                                switch -2147483648 {
                                case 741:
                                    var v18 uint8 = 131
                                    _ = v18
                                case 9223372036854775807:
                                    switch ("a" + "") {
                                    case "a":
                                        v19 := GenS1{F0: 4294967295, F1: 0, F2: int32(v1)}
                                        _ = v19
                                    case "":
                                        var v20 float32 = (1e10 * -2.25)
                                        _ = v20
                                    }
                                case -2147483648:
                                    var v21 uint32 = 1
                                    _ = v21
                                default:
                                    v22 := GenS1{F0: 4294967295, F1: 234, F2: 375}
                                    _ = v22
                                }
                            }
                            var v23 int = (fi14 | 128)
                            if v4 {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                switch "foo bar" {
                                case "\n\t":
                                    switch ("foo bar" + "hello") {
                                    case "a":
                                        v1 += v1
                                    case "hello":
                                        v23 += (v23 >> ((42) & 63))
                                    case "foo bar":
                                        switch "hello" {
                                        case "a":
                                            v24 := GenS1{F0: 363, F1: 99, F2: -1}
                                            _ = v24
                                        case "hello":
                                            v1 += (v1 >> ((1) & 63))
                                            v4 = false
                                        }
                                    }
                                case "":
                                    _ = len(v2)
                                case "hello":
                                    var v25 string = ("hello" + "foo bar")
                                    _ = v25
                                default:
                                    if (0 == 1) {
                                        if v4 {
                                            var v26 int16 = (v10 + -32768)
                                            if true {
                                                if (0 > 0) {
                                                    if (715 >= 2147483648) {
                                                        if (255 > 255) {
                                                            if v4 {
                                                                if (56 <= 65535) {
                                                                    var v27 float64 = -2.25
                                                                    switch fi9 {
                                                                    case 127:
                                                                        var v28 int16 = v10
                                                                        var v29 int32 = 0
                                                                        _ = v28
                                                                        _ = v29
                                                                    case 802:
                                                                        var v30 int = 32768
                                                                        _ = v30
                                                                    }
                                                                    _ = v27
                                                                } else {
                                                                    v23 = (fi14 | (fi14 % (((fi9 % (((fi9 >> ((42) & 63))) | 1))) | 1)))
                                                                    v31 := v3[42]
                                                                    switch ("a" + (((((v31 + v31) + "foo bar") + v31) + v31) + v31)) {
                                                                    case "a":
                                                                        switch fi14 {
                                                                        case 475:
                                                                            fmt.Println(v26, v23, fi9)
                                                                        case 128:
                                                                            var v32 uint8 = 42
                                                                            _ = v32
                                                                        }
                                                                    default:
                                                                        if (0 < 2) {
                                                                            var v33 uint64 = 42
                                                                            var v34 int8 = -128
                                                                            var v35 int = v23
                                                                            v33 = 42
                                                                            _ = v33
                                                                            _ = v34
                                                                            _ = v35
                                                                        }
                                                                    }
                                                                    _ = v31
                                                                }
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        v36 := GenS1{F0: 1, F1: 42, F2: 649}
                                                                        _ = v36
                                                                    case "":
                                                                        v1 = 0
                                                                    default:
                                                                        var v37 uint32 = 1
                                                                        _ = v37
                                                                    }
                                                                default:
                                                                    v1 = -1
                                                                }
                                                                var v38 int16 = v26
                                                                _ = v38
                                                            }
                                                            var v39 float64 = 1e10
                                                            var v40 float32 = (-2.25 * (1.5 / 0.1))
                                                            var v41 float64 = (v39 / v39)
                                                            _ = v39
                                                            _ = v40
                                                            _ = v41
                                                        }
                                                        if (302 >= 2) {
                                                            v1 += 513
                                                            if true {
                                                                v4 = (v4 && (255 < 709))
                                                            } else {
                                                                var v42 int64 = 42
                                                                v3[0] = "\n\t"
                                                                if v4 {
                                                                    fmt.Println(fi14, v26)
                                                                }
                                                                _ = v42
                                                            }
                                                            fmt.Println(fi14)
                                                        }
                                                    }
                                                    if (v26 == (v26 * 42)) {
                                                        if v4 {
                                                            if ((true || v4) && (697 != 32767)) {
                                                                v1 += v1
                                                                fmt.Println(v10, v23)
                                                                fmt.Println(fi9, fi14)
                                                                v10 -= (-(v26))
                                                            } else {
                                                                if (136 < 212) {
                                                                    switch ("\n\t" + "") {
                                                                    case "a":
                                                                        v10 -= v26
                                                                    default:
                                                                        v1 = (-(v1))
                                                                    }
                                                                    var v43 uint64 = 910
                                                                    if v4 {
                                                                        v44 := GenS1{F0: 1, F1: 16, F2: -1}
                                                                        fmt.Println(fi9, v10)
                                                                        var v45 float64 = 0.1
                                                                        _ = v44
                                                                        _ = v45
                                                                    } else {
                                                                        var v46 uint8 = 5
                                                                        _ = v46
                                                                    }
                                                                    _ = v43
                                                                } else {
                                                                    var v47 uint8 = 42
                                                                    _ = v47
                                                                }
                                                            }
                                                        } else {
                                                            switch "a" {
                                                            case "":
                                                                v23 = fi9
                                                            }
                                                        }
                                                        fmt.Println(v1)
                                                        fmt.Println(fi9)
                                                    } else {
                                                        delete(v3, 65535)
                                                        v48 := GenS2{F0: 255}
                                                        v49 := GenS3{F0: 65535, F1: 1, F2: 128}
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                }
                                            } else {
                                                if ((555 == 0) || v4) {
                                                    fmt.Println(v23, v26, v4)
                                                    v26 -= v26
                                                }
                                                if v4 {
                                                    v50 := GenS3{F0: 799, F1: 42, F2: 128}
                                                    _ = v50
                                                }
                                            }
                                            var v51 int16 = v10
                                            var v52 float32 = (1e10 * 0.0)
                                            v53 := GenS2{F0: 2}
                                            _ = v26
                                            _ = v51
                                            _ = v52
                                            _ = v53
                                        }
                                        if ((!((fi9 << ((0) & 63)) > fi14)) || (v4 && (!(0 > 42)))) {
                                            var v54 int8 = int8(v23)
                                            v55 := GenS3{F0: 663, F1: int32(v1), F2: 230}
                                            var v56 float32 = 0.1
                                            _ = v54
                                            _ = v55
                                            _ = v56
                                        } else {
                                            v23 -= 0
                                            v4 = v4
                                            switch (fi14 / ((v23) | 1)) {
                                            case 32767:
                                                fmt.Println(fi9)
                                            case 48:
                                                v4 = (v4 && (1000 < (v10 % (((v10 + v10)) | 1))))
                                            }
                                        }
                                        fmt.Println(v4, v10)
                                        var v57 string = "\n\t"
                                        _ = v57
                                    }
                                }
                                fmt.Println(fi14, v10, fi9)
                            } else {
                                if false {
                                    fmt.Println(v10)
                                    v58 := GenS4{F0: false, F1: 42, F2: "\n\t"}
                                    v10 = v10
                                    _ = v58
                                }
                                if v4 {
                                    switch ((("foo bar" + "foo bar") + "a") + ("a" + "foo bar")) {
                                    case "hello":
                                        var v59 uint64 = 417
                                        _ = v59
                                    case "\n\t":
                                        v60 := GenS5{F0: 0, F1: 2}
                                        _ = v60
                                    case "foo bar":
                                        v61 := GenS3{F0: 42, F1: 532, F2: 252}
                                        _ = v61
                                    default:
                                        if true {
                                            v62 := GenS6{F0: 637, F1: (!v4)}
                                            switch ("foo bar" + "a") {
                                            case "\n\t":
                                                var v63 float64 = 0.0
                                                v1 = 1
                                                _ = v63
                                            default:
                                                v23 = 0
                                            }
                                            v64 := GenS4{F0: ((4294967295 > 4294967295) && (1 <= int8(v10))), F1: (-(v1)), F2: "\n\t"}
                                            _ = v62
                                            _ = v64
                                        } else {
                                            v2 = append(v2, "hello")
                                            _ = len(v2)
                                            var v65 uint = 80
                                            _ = v65
                                        }
                                    }
                                    v66 := GenS7{F0: 65535}
                                    if (((1 <= -42) && true) || false) {
                                        v4 = true
                                    }
                                    _ = v66
                                }
                                switch "a" {
                                case "hello":
                                    if v4 {
                                        fmt.Println(v1)
                                        if v4 {
                                            if (255 != 42) {
                                                v10 -= (v10 / ((1000) | 1))
                                                if (42 <= 2) {
                                                    var v67 uint = 571
                                                    switch "hello" {
                                                    case "":
                                                        var v68 float64 = 1.5
                                                        _ = v68
                                                    case "hello":
                                                        delete(v3, 65535)
                                                    case "foo bar":
                                                        if v4 {
                                                            fmt.Println(v1)
                                                            var v69 int8 = -128
                                                            _ = v69
                                                        }
                                                    default:
                                                        var v70 int64 = (v1 ^ (v1 / ((v1) | 1)))
                                                        _ = v70
                                                    }
                                                    v71 := GenS8{F0: false, F1: v4, F2: 2}
                                                    v72 := GenS9{F0: 2, F1: 110}
                                                    _ = v67
                                                    _ = v71
                                                    _ = v72
                                                } else {
                                                    v73 := GenS10{F0: fi14, F1: (v10 ^ -32768), F2: 9}
                                                    if (v1 >= v1) {
                                                        var v74 uint8 = 127
                                                        var v75 int16 = (v10 + -1)
                                                        _ = v74
                                                        _ = v75
                                                    } else {
                                                        fmt.Println(v1)
                                                        _ = len(v3)
                                                        v76 := GenS8{F0: (127 < 1), F1: (-1 > v1), F2: 42}
                                                        v77 := GenS8{F0: v4, F1: (!v4), F2: 0}
                                                        _ = v76
                                                        _ = v77
                                                    }
                                                    _ = v73
                                                }
                                                if (v4 && (-47 >= -53)) {
                                                    v2 = append(v2, "foo bar")
                                                    fmt.Println(v4, v10, fi14)
                                                } else {
                                                    v78 := v3[1]
                                                    _ = v78
                                                }
                                            }
                                        }
                                        var v79 uint8 = 197
                                        _ = v79
                                    } else {
                                        if v4 {
                                            var v80 uint8 = 127
                                            v81 := GenS10{F0: (fi14 + -1), F1: (v10 >> ((42) & 63)), F2: 133}
                                            _ = v80
                                            _ = v81
                                        }
                                    }
                                case "a":
                                    v10 = (-(v10))
                                }
                                if ((((((1.5 - 1e10) + (((3.14159 * (3.14159 / (float64(float32((0.0 - (1.5 * 3.14159)))) / -2.25))) + 1.5) - 1.5)) + 0.0) - -2.25) / 1.0) < 3.14159) {
                                    v82 := GenS4{F0: v4, F1: v1, F2: "hello"}
                                    var v83 uint16 = 42
                                    _ = v82
                                    _ = v83
                                }
                                var v84 float32 = 3.14159
                                _ = v84
                            }
                            _ = v23
                        }
                        _ = fi14
                    }
                    _ = len(v3)
                }
                var v85 uint64 = 0
                v85 += 1
                _ = v85
            }
        case 297:
            for fi86 := 0; fi86 < 11; fi86++ {
                var v87 bool = (!(v4 || (-2.25 < (float32(float64(1.5)) * float32(-2.25)))))
                var v88 float64 = 1.5
                _ = fi86
                _ = v87
                _ = v88
            }
        default:
            for fi89 := 0; fi89 < 4; fi89++ {
                v1 = v1
                _ = fi89
            }
        }
        if ((!v4) && (255 != 255)) {
            v3[0] = ("\n\t" + "a")
        }
        _ = fi9
        _ = v10
    }
    var v90 int = -1
    if true {
        switch (v90 << ((4294967295) & 63)) {
        case -9223372036854775808:
            switch ("\n\t" + (("" + "foo bar") + ("a" + "\n\t"))) {
            case "":
                var v91 uint8 = 255
                _ = v91
            case "\n\t":
                var v92 uint = 2
                _ = v92
            default:
                v4 = v4
            }
        case 891:
            v4 = v4
        case -1:
            var v93 uint64 = 0
            _ = v93
        }
        var v94 string = (("\n\t" + "hello") + "")
        v90 += (v90 & -9223372036854775808)
        _ = v94
    } else {
        var v95 int8 = 0
        v96 := GenS11{F0: 880, F1: 42, F2: false}
        _ = v95
        _ = v96
    }
    if (42 >= 20) {
        v97 := v3[352]
        var v98 int = v90
        _ = v97
        _ = v98
    }
    v99 := v3[196]
    v100 := GenS12{F0: (v1 + v1), F1: 4294967295, F2: 10}
    v3[786] = ((("\n\t" + "hello") + "hello") + ("hello" + "a"))
    switch ((v99 + ("" + "foo bar")) + v99) {
    case "foo bar":
        for fi101 := 0; fi101 < 3; fi101++ {
            var v102 int32 = 1
            var v103 uint32 = 0
            _ = fi101
            _ = v102
            _ = v103
        }
    case "a":
        var v104 uint8 = 42
        _ = v104
    default:
        for fi105 := 0; fi105 < 9; fi105++ {
            v2 = append(v2, v99)
            v106 := GenS13{F0: 638}
            fmt.Println(v90, fi105)
            _ = fi105
            _ = v106
        }
    }
    v90 -= -288
    v90 = v90
    var v107 string = v99
    var v108 int16 = 0
    for fi109 := 0; fi109 < 11; fi109++ {
        fmt.Println(v90)
        for fi110 := 0; fi110 < 6; fi110++ {
            switch v107 {
            case "\n\t":
                _ = len(v2)
            default:
                if ((fi110 - int(v1)) < (^(fi109))) {
                    fmt.Println(v108, v99)
                    v99 = (("a" + ("" + v107)) + v107)
                    v99 = v99
                } else {
                    fmt.Println(v108, v108)
                    v107 = (((v99 + "hello") + "foo bar") + (v107 + "\n\t"))
                }
            }
            switch v99 {
            case "":
                v111 := GenS1{F0: 392, F1: 42, F2: 42}
                _ = v111
            default:
                v1 = v1
            }
            _ = len(v2)
            var v112 uint64 = 4294967295
            _ = fi110
            _ = v112
        }
        _ = fi109
    }
    for fi113 := 0; fi113 < 2; fi113++ {
        var v114 uint = 962
        _ = fi113
        _ = v114
    }
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("v108:", v108)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v90:", v90)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v107
    _ = v108
    _ = v2
    _ = v3
    _ = v4
    _ = v90
    _ = v99
}
