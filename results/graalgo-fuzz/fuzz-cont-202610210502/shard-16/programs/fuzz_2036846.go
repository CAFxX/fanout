package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 bool; F1 int64; F2 int32 }
type GenS6 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS7 struct { F0 int16; F1 int32; F2 int8 }
type GenS8 struct { F0 uint8; F1 uint8 }
type GenS9 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS10 struct { F0 int; F1 int64 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 bool; F1 string; F2 bool }
type GenS13 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS14 struct { F0 int }
func genFn0(p0 string, p1 int32) bool {
    if (0 > -963) {
        v5 := []int{-2147483648, 0, -9223372036854775808, -272}
        if true {
            if (295 == -185) {
                switch 20 {
                case 205:
                    if len(v5) > 0 {
                        v5[0] = -9223372036854775808
                    }
                default:
                    var v6 int32 = (p1 >> ((42) & 63))
                    _ = v6
                }
            }
            fmt.Println(p0, p1)
            p1 -= (-(p1))
        }
        var v7 int32 = p1
        _ = v5
        _ = v7
    }
    var v8 float64 = 0.1
    fmt.Println(p0, p1, p0)
    if (255 == 65535) {
        v9 := GenS1{F0: "hello"}
        switch "a" {
        case "hello":
            v10 := []int{-9223372036854775808, -1}
            _ = v10
        }
        p1 = -560
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v8
    return ((((true || (843 >= 42)) && false) && false) || true)
}
func genFn1(p0 int32, p1 uint8) int64 {
    switch 782 {
    case -1:
        switch ("hello" + "\n\t") {
        case "\n\t":
            v11 := GenS1{F0: ("\n\t" + (("" + "") + "\n\t"))}
            _ = v11
        case "foo bar":
            if false {
                v12 := GenS1{F0: ("" + "")}
                _ = v12
            }
        }
    case -810:
        var v13 int64 = 861
        _ = v13
    }
    if false {
        p1 += p1
        p0 -= -1
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 uint64 = 373
            for fi16 := 0; fi16 < 10; fi16++ {
                var v17 string = "\n\t"
                _ = fi16
                _ = v17
            }
            p0 += -1
            _ = fi14
            _ = v15
        }
    } else {
        var v18 int8 = 127
        _ = v18
    }
    p0 -= -614
    switch "a" {
    case "\n\t":
        v19 := GenS2{F0: 127}
        _ = v19
    default:
        v20 := GenS3{F0: 14}
        _ = v20
    }
    _ = p0
    _ = p1
    return 117
}

func main() {
    var v1 uint8 = 9
    var v2 int64 = -354
    var v3 int16 = -1
    v4 := make(map[uint32]string)
    switch "hello" {
    case "hello":
        if (!(!(!true))) {
            if false {
                v21 := GenS4{F0: 1, F1: 847}
                if false {
                    switch (("" + (("a" + "a") + "foo bar")) + ("foo bar" + "a")) {
                    case "\n\t":
                        if true {
                            for fi22 := 0; fi22 < 12; fi22++ {
                                if ((true && true) && (!(-2147483648 <= -2147483648))) {
                                    fmt.Println(v3, v2, v2)
                                }
                                _ = fi22
                            }
                            _ = len(v4)
                            switch ("" + "") {
                            case "hello":
                                v2 = 1
                            case "":
                                v23 := GenS1{F0: (("\n\t" + "a") + "hello")}
                                _ = v23
                            default:
                                v24 := []int{-218}
                                _ = v24
                            }
                        }
                    default:
                        for fi25 := 0; fi25 < 14; fi25++ {
                            _ = len(v4)
                            _ = fi25
                        }
                    }
                    var v26 int8 = -66
                    for fi27 := 0; fi27 < 3; fi27++ {
                        v28 := GenS1{F0: ""}
                        _ = fi27
                        _ = v28
                    }
                    _ = v26
                } else {
                    v3 += 0
                }
                _ = v21
            } else {
                for fi29 := 0; fi29 < 14; fi29++ {
                    var v30 uint8 = v1
                    var v31 int8 = -1
                    fmt.Println(v1, v31, v3)
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                v2 -= genFn1(int32(v2), 127)
                v32 := GenS5{F0: false, F1: (v2 ^ int64(v3)), F2: -1}
                _ = v32
            }
            delete(v4, 2)
        } else {
            if true {
                fmt.Println(v1, v2, v3)
                v2 = (v2 / ((-1) | 1))
                v3 -= v3
            } else {
                v2 = 642
                v3 += 1000
            }
        }
    case "":
        v3 = (v3 % ((42) | 1))
    default:
        var v33 int64 = v2
        _ = v33
    }
    v34 := []int{-144}
    for fi35 := 0; fi35 < 17; fi35++ {
        v1 = (^(v1))
        if len(v34) >= 2 {
            v34 = v34[1:len(v34)]
        }
        switch "" {
        case "\n\t":
            var v36 uint = 18446744073709551615
            _ = v36
        case "hello":
            fmt.Println(v2, v3, fi35)
        case "a":
            for fi37 := 0; fi37 < 10; fi37++ {
                v2 = v2
                v1 = (v1 >> ((18446744073709551615) & 63))
                if true {
                    if false {
                        if len(v34) >= 2 {
                            v34 = v34[1:len(v34)]
                        }
                    }
                    v1 = (v1 - (v1 - (v1 | (v1 | v1))))
                    switch 4294967295 {
                    case 362:
                        if true {
                            v38 := GenS2{F0: -1}
                            _ = v38
                        } else {
                            fmt.Println(v3)
                            if len(v34) >= 2 {
                                v34 = v34[1:len(v34)]
                            }
                        }
                    case 1:
                        if (1e10 == (-2.25 + float32(3.14159))) {
                            switch "foo bar" {
                            case "a":
                                var v39 uint32 = 2147483647
                                v40 := GenS6{F0: 4294967295, F1: 1, F2: 2147483647}
                                if len(v34) >= 2 {
                                    v34 = v34[1:len(v34)]
                                }
                                _ = v39
                                _ = v40
                            case "foo bar":
                                if (!(!true)) {
                                    var v41 int8 = -91
                                    _ = v41
                                }
                            case "\n\t":
                                fmt.Println(fi37)
                            default:
                                v2 -= v2
                            }
                        } else {
                            if (!false) {
                                v42 := GenS4{F0: 18446744073709551615, F1: 42}
                                _ = v42
                            }
                            if (!(!(4294967295 == 63))) {
                                v2 = -914
                                var v43 int64 = -9223372036854775808
                                v44 := GenS5{F0: true, F1: 9223372036854775807, F2: 42}
                                var v45 int8 = -51
                                fmt.Println(v45, v1)
                                if (!(2 >= 399)) {
                                    switch 1 {
                                    case 128:
                                        v3 = (-(v3))
                                    case 2147483647:
                                        switch "a" {
                                        case "\n\t":
                                            if (!((0.0 - 0.1) < 1.5)) {
                                                var v46 int = 118
                                                _ = v46
                                            } else {
                                                v47 := GenS3{F0: -925}
                                                _ = v47
                                            }
                                            fmt.Println(v43, v3)
                                        case "":
                                            v3 = v3
                                        }
                                    default:
                                        var v48 float32 = (3.14159 + ((-2.25 + (float32(1e10) + (1.5 * -2.25))) / 3.14159))
                                        _ = v48
                                    }
                                    if (76 > 1) {
                                        switch int(v3) {
                                        case 127:
                                            switch 32768 {
                                            case 32768:
                                                v1 = 127
                                            case 0:
                                                v49 := GenS1{F0: "a"}
                                                _ = v49
                                            }
                                        case 406:
                                            v1 = v1
                                        case 32767:
                                            v50 := GenS5{F0: false, F1: (v43 | (v2 << ((4294967295) & 63))), F2: 90}
                                            _ = v50
                                        default:
                                            if (false && (false && true)) {
                                                var v51 int = -22
                                                _ = v51
                                            } else {
                                                var v52 bool = (-1 == 0)
                                                v1 = (^(v1))
                                                var v53 int16 = (-(v3))
                                                _ = v52
                                                _ = v53
                                            }
                                        }
                                    }
                                }
                                v45 -= v45
                                _ = v43
                                _ = v44
                                _ = v45
                            } else {
                                if len(v34) >= 2 {
                                    v34 = v34[1:len(v34)]
                                }
                                v54 := GenS7{F0: (v3 / (((v3 | v3)) | 1)), F1: 1, F2: 1}
                                _ = v54
                            }
                        }
                        v3 = v3
                        var v55 int64 = -1
                        _ = v55
                    case 119:
                        if false {
                            v1 -= 2
                        } else {
                            if false {
                                switch 32768 {
                                case -1:
                                    delete(v4, 42)
                                case 0:
                                    if (943 == 943) {
                                        v56 := GenS1{F0: "hello"}
                                        if true {
                                            var v57 int32 = -516
                                            v3 -= 32767
                                            _ = v57
                                        } else {
                                            _ = len(v34)
                                            v34 = append(v34, 128)
                                            if ((!(!false)) || (!(((float64(float32((3.14159 + float64(-2.25)))) / 1.0) != ((0.1 + (1.5 / 1e10)) * 0.0)) && true))) {
                                                var v58 uint16 = 255
                                                switch "\n\t" {
                                                case "hello":
                                                    v58 -= 308
                                                case "a":
                                                    switch (fi37 << ((2) & 63)) {
                                                    case 128:
                                                        if ("hello" <= ((("hello" + "hello") + ("a" + "\n\t")) + "foo bar")) {
                                                            var v59 uint32 = 2
                                                            v60 := GenS8{F0: (v1 >> ((18446744073709551615) & 63)), F1: v1}
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                    case 1:
                                                        fmt.Println(v3)
                                                    case 9223372036854775807:
                                                        fmt.Println(v2)
                                                    }
                                                case "foo bar":
                                                    var v61 uint64 = 0
                                                    _ = v61
                                                }
                                                var v62 int = fi37
                                                v63 := GenS5{F0: true, F1: 9223372036854775807, F2: 1}
                                                v64 := v4[854]
                                                _ = v58
                                                _ = v62
                                                _ = v63
                                                _ = v64
                                            }
                                        }
                                        v65 := GenS9{F0: 32767, F1: v1, F2: v1}
                                        _ = v56
                                        _ = v65
                                    } else {
                                        v3 = (v3 * (v3 >> ((42) & 63)))
                                        if (!(v3 > (^(v3)))) {
                                            var v66 float32 = (0.1 + float32(3.14159))
                                            _ = v66
                                        } else {
                                            var v67 int16 = 0
                                            _ = v67
                                        }
                                        var v68 int64 = 0
                                        _ = v68
                                    }
                                default:
                                    v1 += (v1 ^ v1)
                                }
                            }
                        }
                    }
                    v69 := GenS10{F0: fi35, F1: v2}
                    v70 := GenS10{F0: 876, F1: 0}
                    _ = v69
                    _ = v70
                } else {
                    _ = len(v4)
                }
                _ = fi37
            }
        default:
            if false {
                v2 += genFn1(-2147483648, v1)
                v71 := GenS4{F0: uint(v1), F1: 2147483647}
                v72 := GenS11{F0: 249}
                _ = v71
                _ = v72
            } else {
                var v73 uint32 = 2147483648
                _ = v73
            }
        }
        _ = fi35
    }
    var v74 int8 = 0
    for fi75 := 0; fi75 < 14; fi75++ {
        fmt.Println(v2)
        switch ("" + "") {
        case "foo bar":
            var v76 uint32 = uint32(v1)
            _ = v76
        case "\n\t":
            v1 += v1
        case "a":
            fmt.Println(v74, v2)
        }
        fmt.Println(v1)
        _ = fi75
    }
    switch "hello" {
    case "\n\t":
        if true {
            v77 := GenS9{F0: 129, F1: (^(v1)), F2: 127}
            if true {
                fmt.Println(v74, v1, v74)
                v3 = 346
                fmt.Println(v74, v74, v2)
            } else {
                v78 := GenS2{F0: v74}
                v2 -= v2
                _ = v78
            }
            _ = v77
        } else {
            var v79 int8 = v74
            v79 = v74
            v1 += (v1 ^ 92)
            _ = v79
        }
    case "foo bar":
        switch 128 {
        case -9223372036854775808:
            for fi80 := 0; fi80 < 14; fi80++ {
                switch "foo bar" {
                case "foo bar":
                    if ((false && ((255 <= 65535) && false)) && (v1 == v1)) {
                        v81 := GenS7{F0: -32768, F1: 42, F2: (v74 << ((633) & 63))}
                        _ = v81
                    } else {
                        delete(v4, 42)
                        switch (fi80 % ((fi80) | 1)) {
                        case 855:
                            v74 = v74
                        }
                    }
                case "":
                    fmt.Println(fi80, fi80)
                case "hello":
                    fmt.Println(v3, v1)
                default:
                    v82 := GenS2{F0: 100}
                    _ = v82
                }
                _ = fi80
            }
        case -1:
            if (4294967295 > 255) {
                v83 := GenS9{F0: (-(v3)), F1: v1, F2: (v1 + v1)}
                v84 := GenS12{F0: (95 == -959), F1: "", F2: ((!(364 == 139)) || (255 > v1))}
                _ = v83
                _ = v84
            }
        default:
            v2 -= v2
        }
    }
    for fi85 := 0; fi85 < 10; fi85++ {
        _ = len(v34)
        _ = fi85
    }
    switch -402 {
    case -9223372036854775808:
        var v86 int32 = -446
        _ = v86
    case 672:
        for fi87 := 0; fi87 < 8; fi87++ {
            _ = len(v4)
            _ = fi87
        }
    case -2147483648:
        if (true || false) {
            if (!false) {
                var v88 uint16 = 874
                var v89 uint32 = 2147483648
                var v90 int8 = -128
                _ = v88
                _ = v89
                _ = v90
            }
            var v91 bool = ((v2 + genFn1(-2147483648, v1)) >= -205)
            _ = v91
        } else {
            var v92 int8 = 100
            var v93 uint8 = v1
            _ = v92
            _ = v93
        }
    }
    v3 -= (v3 / (((-(v3))) | 1))
    var v94 int32 = -1
    fmt.Println(v1)
    var v95 uint = 2
    if ((32768 != 32767) && (int16(v94) != (v3 % ((-132) | 1)))) {
        var v96 int16 = 32767
        switch 2147483647 {
        case -9223372036854775808:
            v94 += (v94 / ((v94) | 1))
        case -682:
            if false {
                for fi97 := 0; fi97 < 6; fi97++ {
                    v98 := GenS8{F0: (^(v1)), F1: v1}
                    for fi99 := 0; fi99 < 15; fi99++ {
                        if false {
                            v95 += v95
                            v94 = (v94 | -2147483648)
                        } else {
                            v100 := GenS13{F0: 100, F1: (v1 << ((42) & 63)), F2: 42}
                            v74 = (v74 << ((v95) & 63))
                            var v101 int32 = (v94 >> (((v95 << (((v95 & (v95 >> ((2) & 63)))) & 63))) & 63))
                            _ = v100
                            _ = v101
                        }
                        _ = fi99
                    }
                    fmt.Println(v1, v95)
                    _ = fi97
                    _ = v98
                }
            } else {
                var v102 uint16 = 2
                v4[2147483647] = "hello"
                _ = v102
            }
        }
        for fi103 := 0; fi103 < 15; fi103++ {
            v104 := GenS5{F0: ((v1 <= (^(v1))) || ((^(v1)) < v1)), F1: 1, F2: -1}
            _ = fi103
            _ = v104
        }
        _ = v96
    } else {
        var v105 float64 = 1.5
        _ = v105
    }
    v4[2147483647] = ((("\n\t" + "foo bar") + "") + ("" + "foo bar"))
    for fi106 := 0; fi106 < 13; fi106++ {
        if true {
            v4[2147483648] = ("" + "a")
        } else {
            v107 := GenS4{F0: 1, F1: -80}
            if (true || true) {
                v108 := GenS6{F0: (v95 * (v95 & (v95 ^ v95))), F1: 255, F2: 145}
                if false {
                    var v109 uint64 = 4294967295
                    v110 := GenS14{F0: (fi106 | fi106)}
                    fmt.Println(v3, v2, v2)
                    _ = v109
                    _ = v110
                } else {
                    v111 := GenS10{F0: fi106, F1: (-(v2))}
                    _ = v111
                }
                if true {
                    fmt.Println(v2)
                } else {
                    if len(v34) >= 2 {
                        v34 = v34[1:len(v34)]
                    }
                    var v112 int = fi106
                    _ = v112
                }
                _ = v108
            }
            _ = v107
        }
        v1 = v1
        var v113 int64 = (v2 / (((v2 * (v2 << (((v95 & v95)) & 63)))) | 1))
        _ = fi106
        _ = v113
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v74:", v74)
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v74
    _ = v94
    _ = v95
}
