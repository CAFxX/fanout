package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 int; F1 bool; F2 int32 }
type GenS5 struct { F0 uint64; F1 int; F2 int16 }
type GenS6 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS7 struct { F0 int; F1 bool; F2 string }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 string; F1 uint64; F2 uint8 }
type GenS10 struct { F0 int16; F1 int32 }
type GenS11 struct { F0 int8; F1 uint8; F2 bool }
type GenS12 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS13 struct { F0 int32 }
func genFn0(p0 bool, p1 int8, p2 int16) uint {
    v4 := []int{0}
    p2 = p2
    if true {
        v5 := GenS1{F0: 2147483648, F1: "foo bar"}
        _ = v5
    } else {
        var v6 int32 = -2147483648
        v7 := GenS1{F0: 1, F1: ("foo bar" + ("\n\t" + "\n\t"))}
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 255
}
func genFn1(p0 bool, p1 uint64, p2 string) int32 {
    switch 42 {
    case 378:
        v8 := []int{42, -514, 351}
        _ = v8
    }
    fmt.Println(p2, p2, p2)
    var v9 uint16 = 1
    p1 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -1
}

func main() {
    v1 := []int{-2147483648}
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    var v10 int16 = -986
    v11 := GenS1{F0: 156, F1: ("\n\t" + "a")}
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 bool = true
        v13 = false
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 uint32 = 346
        for fi16 := 0; fi16 < 19; fi16++ {
            if (!false) {
                if false {
                    switch (("\n\t" + "") + ("foo bar" + ("hello" + "hello"))) {
                    case "hello":
                        v15 -= (v15 % (((v15 & (v15 >> ((65535) & 63)))) | 1))
                    case "a":
                        if false {
                            switch ("\n\t" + (("hello" + ("\n\t" + ("hello" + "foo bar"))) + "hello")) {
                            case "foo bar":
                                v2[(^(v15))] = fi14
                            }
                            var v17 bool = (false && (-2147483648 <= fi14))
                            _ = v17
                        }
                    case "foo bar":
                        if ((!true) && (genFn1(false, 2, "hello") < genFn1((!true), 882, ("foo bar" + "")))) {
                            fmt.Println(fi14)
                            var v18 string = (("a" + "hello") + "a")
                            fmt.Println(fi14, fi16)
                            v19 := v3[195]
                            _ = v18
                            _ = v19
                        }
                    default:
                        if false {
                            v10 = v10
                            v10 -= (v10 | (v10 ^ v10))
                        }
                    }
                }
                fmt.Println(v10, fi14)
                if false {
                    v20 := GenS1{F0: v15, F1: "hello"}
                    v21 := GenS2{F0: genFn1((1e10 == 3.14159), 0, "\n\t")}
                    v2[0] = 128
                    _ = v20
                    _ = v21
                }
            } else {
                var v22 uint = 0
                _ = v22
            }
            v10 = 1
            _ = fi16
        }
        var v23 int32 = 0
        _ = fi14
        _ = v15
        _ = v23
    }
    if (false && (2 != 770)) {
        delete(v3, 42)
        v24 := GenS2{F0: -2147483648}
        _ = v24
    }
    for fi25 := 0; fi25 < 19; fi25++ {
        switch (((("a" + "\n\t") + "\n\t") + "") + (("foo bar" + "") + "a")) {
        case "\n\t":
            v26 := GenS1{F0: 2147483648, F1: ""}
            _ = v26
        case "foo bar":
            _ = len(v1)
        case "":
            if true {
                if false {
                    v27 := GenS3{F0: 987, F1: -84}
                    v10 = v10
                    var v28 int8 = -1
                    _ = v27
                    _ = v28
                } else {
                    if false {
                        switch (fi25 >> ((255) & 63)) {
                        case 0:
                            if ((fi25 / ((fi25) | 1)) != -581) {
                                var v29 uint = 42
                                v30 := v2[2147483648]
                                _ = v29
                                _ = v30
                            }
                        default:
                            for fi31 := 0; fi31 < 16; fi31++ {
                                delete(v2, 1)
                                var v32 uint16 = 32767
                                v10 = 370
                                _ = fi31
                                _ = v32
                            }
                        }
                        v1 = append(v1, fi25)
                        v10 = (v10 ^ v10)
                    } else {
                        v33 := GenS4{F0: (^(fi25)), F1: ((-2.25 > 0.1) || true), F2: -244}
                        if (!(0 <= -27)) {
                            var v34 bool = ((float32((((float64(float32(((((0.1 + 1e10) + 0.1) + 3.14159) + (0.1 - 0.1)))) - 1e10) - float64((0.0 - (3.14159 - (1.5 / 1e10))))) * 0.1)) - (0.1 * 1.5)) == 0.1)
                            v34 = v34
                            v35 := GenS2{F0: 0}
                            _ = v34
                            _ = v35
                        } else {
                            var v36 uint16 = 32767
                            if (v36 > (v36 & v36)) {
                                fmt.Println(fi25)
                                if true {
                                    if ((false || false) && (true || (1 >= (fi25 << ((42) & 63))))) {
                                        v10 = (v10 % (((v10 % (((-(v10))) | 1))) | 1))
                                        var v37 int8 = 90
                                        _ = v37
                                    } else {
                                        if (42 != (fi25 | (fi25 | (fi25 << ((255) & 63))))) {
                                            for fi38 := 0; fi38 < 10; fi38++ {
                                                switch (("" + ("hello" + ("hello" + "foo bar"))) + "") {
                                                case "hello":
                                                    v10 += v10
                                                case "":
                                                    var v39 uint16 = (v36 >> ((genFn0(false, -95, v10)) & 63))
                                                    _ = v39
                                                case "a":
                                                    v40 := GenS5{F0: 42, F1: fi25, F2: (v10 | v10)}
                                                    _ = v40
                                                default:
                                                    var v41 bool = (65535 <= 255)
                                                    v36 -= (v36 % (((v36 + (v36 << ((255) & 63)))) | 1))
                                                    _ = v41
                                                }
                                                _ = fi38
                                            }
                                            v42 := v3[(fi25 ^ (^(fi25)))]
                                            _ = v42
                                        } else {
                                            v43 := v2[2147483647]
                                            fmt.Println(fi25)
                                            _ = v43
                                        }
                                        var v44 uint8 = 1
                                        if ((1.5 * 0.1) > 0.0) {
                                            for fi45 := 0; fi45 < 18; fi45++ {
                                                if false {
                                                    fmt.Println(v44, fi25)
                                                    v46 := GenS6{F0: (v36 + (^(v36))), F1: v10, F2: 0}
                                                    _ = v46
                                                } else {
                                                    var v47 int16 = 42
                                                    switch 32767 {
                                                    case 0:
                                                        v1 = append(v1, fi45)
                                                    case 42:
                                                        var v48 string = "a"
                                                        _ = v48
                                                    }
                                                    _ = v47
                                                }
                                                var v49 float32 = 1.5
                                                if true {
                                                    var v50 int8 = 42
                                                    fmt.Println(fi45)
                                                    fmt.Println(v10)
                                                    _ = v50
                                                } else {
                                                    if true {
                                                        var v51 int8 = 117
                                                        _ = v51
                                                    } else {
                                                        var v52 int64 = 0
                                                        _ = v52
                                                    }
                                                    v53 := GenS1{F0: 42, F1: "a"}
                                                    _ = v53
                                                }
                                                _ = fi45
                                                _ = v49
                                            }
                                            var v54 string = "a"
                                            switch 318 {
                                            case 9223372036854775807:
                                                for fi55 := 0; fi55 < 4; fi55++ {
                                                    delete(v3, fi55)
                                                    fmt.Println(fi55, v44, v36)
                                                    _ = fi55
                                                }
                                            case -861:
                                                v56 := GenS7{F0: 32768, F1: false, F2: ((v54 + "a") + v54)}
                                                _ = v56
                                            default:
                                                v57 := GenS8{F0: 9223372036854775807}
                                                _ = v57
                                            }
                                            _ = v54
                                        }
                                        _ = v44
                                    }
                                    if true {
                                        delete(v3, (fi25 % ((9223372036854775807) | 1)))
                                        v58 := GenS9{F0: "foo bar", F1: 963, F2: 1}
                                        v59 := v3[(fi25 & 4294967295)]
                                        _ = v58
                                        _ = v59
                                    }
                                }
                            }
                            switch (fi25 ^ -2147483648) {
                            case 2147483647:
                                var v60 int = int(v36)
                                _ = v60
                            case 1:
                                fmt.Println(v10, v10)
                            case 317:
                                var v61 uint32 = 2
                                _ = v61
                            }
                            _ = v36
                        }
                        _ = v33
                    }
                    if true {
                        v62 := GenS10{F0: v10, F1: 641}
                        for fi63 := 0; fi63 < 4; fi63++ {
                            if false {
                                switch ("a" + ("foo bar" + "a")) {
                                case "hello":
                                    v10 = -851
                                case "foo bar":
                                    var v64 bool = true
                                    _ = v64
                                default:
                                    v10 += 42
                                }
                                if (fi63 >= (^(fi25))) {
                                    var v65 int32 = genFn1(true, 576, "\n\t")
                                    switch fi25 {
                                    case -234:
                                        switch "foo bar" {
                                        case "a":
                                            var v66 uint = 0
                                            _ = v66
                                        case "\n\t":
                                            switch "a" {
                                            case "hello":
                                                if (((((fi63 <= (fi63 & fi25)) || (128 < 86)) && true) && (9223372036854775807 >= 1)) || ((v65 & (v65 - 42)) > genFn1((!((1 < 42) || false)), 18446744073709551615, (("" + "") + ("" + "a"))))) {
                                                    v67 := GenS1{F0: 360, F1: ("\n\t" + "\n\t")}
                                                    _ = v67
                                                } else {
                                                    fmt.Println(fi25, v65)
                                                }
                                            default:
                                                var v68 uint64 = 354
                                                _ = v68
                                            }
                                            v69 := GenS11{F0: -115, F1: 1, F2: true}
                                            _ = v69
                                        default:
                                            var v70 int16 = v10
                                            _ = v70
                                        }
                                    case 0:
                                        v71 := v2[42]
                                        _ = v71
                                    case 128:
                                        v72 := GenS11{F0: -49, F1: 127, F2: (v10 > (v10 % (((v10 >> ((65535) & 63))) | 1)))}
                                        _ = v72
                                    default:
                                        v10 += v10
                                    }
                                    v10 = (v10 | (-(v10)))
                                    _ = v65
                                } else {
                                    fmt.Println(v10, fi25, v10)
                                    if ((!false) || (-9223372036854775808 < -1)) {
                                        var v73 uint16 = 237
                                        v10 = (v10 << ((18446744073709551615) & 63))
                                        v10 = (v10 / ((int16(fi25)) | 1))
                                        _ = v73
                                    } else {
                                        var v74 uint32 = 811
                                        _ = v74
                                    }
                                    v75 := GenS12{F0: 42, F1: 67, F2: -9223372036854775808}
                                    if ((((0.1 * 0.0) / ((1.5 + 1e10) / ((0.1 / 0.1) / 1.0))) == 0.1) && true) {
                                        switch 676 {
                                        case 625:
                                            switch "hello" {
                                            case "foo bar":
                                                v76 := GenS4{F0: 9223372036854775807, F1: true, F2: -2147483648}
                                                _ = v76
                                            case "":
                                                v1 = append(v1, (fi63 ^ fi63))
                                            case "a":
                                                v77 := GenS9{F0: "a", F1: 0, F2: 2}
                                                _ = v77
                                            }
                                        case 0:
                                            var v78 uint8 = 0
                                            _ = v78
                                        }
                                    }
                                    var v79 int16 = v10
                                    switch "hello" {
                                    case "":
                                        var v80 uint8 = 128
                                        _ = v80
                                    case "\n\t":
                                        var v81 uint32 = 2
                                        _ = v81
                                    }
                                    _ = v75
                                    _ = v79
                                }
                                fmt.Println(fi63, fi63, fi25)
                                if false {
                                    if (!(("a" + ("foo bar" + "")) >= "")) {
                                        if (215 > 255) {
                                            if (-2.25 <= -2.25) {
                                                fmt.Println(v10)
                                            }
                                            if true {
                                                if true {
                                                    var v82 uint64 = 0
                                                    fmt.Println(v82, v10, fi63)
                                                    _ = v82
                                                } else {
                                                    v83 := v2[2147483648]
                                                    _ = v83
                                                }
                                                var v84 float64 = -2.25
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = v84
                                            } else {
                                                var v85 uint = 40
                                                v10 -= v10
                                                if true {
                                                    v85 = (^(v85))
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    v86 := GenS12{F0: 65535, F1: 199, F2: -9223372036854775808}
                                                    _ = v86
                                                }
                                                v87 := GenS13{F0: 1}
                                                _ = v85
                                                _ = v87
                                            }
                                            v10 += (v10 >> ((42) & 63))
                                        } else {
                                            var v88 uint64 = 0
                                            fmt.Println(fi25, fi25)
                                            _ = v88
                                        }
                                    } else {
                                        var v89 int64 = 1
                                        v3[(fi25 / ((42) | 1))] = (1 <= 868)
                                        _ = v89
                                    }
                                    v10 = -1
                                }
                            }
                            _ = fi63
                        }
                        for fi90 := 0; fi90 < 13; fi90++ {
                            switch 32768 {
                            case -2147483648:
                                fmt.Println(v10, v10, fi90)
                            case 0:
                                fmt.Println(v10, v10)
                            case 1:
                                switch "hello" {
                                case "\n\t":
                                    fmt.Println(v10)
                                default:
                                    v10 += (v10 | v10)
                                }
                            }
                            var v91 int8 = 127
                            _ = fi90
                            _ = v91
                        }
                        _ = v62
                    }
                    if (255 >= 340) {
                        fmt.Println(v10, v10)
                        var v92 float64 = ((0.0 / 0.1) * 0.0)
                        _ = v92
                    }
                }
            }
        default:
            v1 = append(v1, fi25)
        }
        _ = fi25
    }
    v93 := GenS3{F0: 0, F1: 1}
    fmt.Println(v10)
    fmt.Println(v10)
    v94 := GenS11{F0: 15, F1: 42, F2: (!(-1 < genFn1((333 > 195), 42, "a")))}
    fmt.Println(v10)
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (-964 < -368) {
            var v95 uint16 = 42
            _ = v95
        }
    } else {
        _ = len(v2)
    }
    for fi96 := 0; fi96 < 20; fi96++ {
        var v97 float64 = -2.25
        _ = fi96
        _ = v97
    }
    switch "hello" {
    case "\n\t":
        if (!false) {
            fmt.Println(v10)
        } else {
            v10 = v10
            var v98 uint16 = 2
            _ = v98
        }
    case "foo bar":
        switch "foo bar" {
        case "hello":
            var v99 uint8 = 127
            _ = v99
        case "a":
            v10 = v10
        default:
            switch 2147483647 {
            case -2147483648:
                v10 += (^(v10))
            case 1:
                switch "a" {
                case "":
                    for fi100 := 0; fi100 < 13; fi100++ {
                        v101 := GenS12{F0: 255, F1: 128, F2: 0}
                        _ = fi100
                        _ = v101
                    }
                }
            case -9223372036854775808:
                for fi102 := 0; fi102 < 14; fi102++ {
                    if (442 != 1) {
                        fmt.Println(fi102, fi102)
                    }
                    _ = fi102
                }
            default:
                v10 = v10
            }
        }
    case "hello":
        for fi103 := 0; fi103 < 20; fi103++ {
            for fi104 := 0; fi104 < 10; fi104++ {
                v10 = v10
                _ = fi104
            }
            _ = fi103
        }
    }
    var v105 uint = 943
    v105 = (v105 >> ((v105) & 63))
    if false {
        for fi106 := 0; fi106 < 1; fi106++ {
            if true {
                var v107 int = (fi106 % (((fi106 & fi106)) | 1))
                _ = v107
            } else {
                v10 = v10
            }
            for fi108 := 0; fi108 < 7; fi108++ {
                var v109 int32 = -97
                v105 += v105
                var v110 bool = true
                _ = fi108
                _ = v109
                _ = v110
            }
            var v111 uint16 = 1
            _ = fi106
            _ = v111
        }
        if (false && true) {
            v112 := GenS3{F0: 0, F1: 1}
            _ = v112
        } else {
            switch ("\n\t" + "a") {
            case "foo bar":
                var v113 uint8 = 128
                _ = v113
            case "":
                var v114 uint32 = 840
                _ = v114
            }
            v115 := v2[2147483647]
            _ = v115
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v105:", v105)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v105
    _ = v11
    _ = v2
    _ = v3
    _ = v93
    _ = v94
}
