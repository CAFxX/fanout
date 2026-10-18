package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int8; F1 uint; F2 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool; F1 int16; F2 int16 }
type GenS6 struct { F0 int32; F1 int32 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int16; F1 uint64; F2 int32 }
type GenS10 struct { F0 bool; F1 string; F2 bool }
type GenS11 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS12 struct { F0 int; F1 uint8 }
type GenS13 struct { F0 int32; F1 uint32; F2 uint64 }

func main() {
    var v1 float64 = float64(float32((1e10 / 1.5)))
    var v2 uint = 979
    var v3 float64 = 1.5
    var v4 int16 = 0
    v3 += v1
    if true {
        v3 = ((3.14159 + v1) * 1.5)
        var v5 uint = v2
        fmt.Println(v5)
        _ = v5
    }
    for fi6 := 0; fi6 < 6; fi6++ {
        fmt.Println(fi6, v4)
        var v7 float32 = (float32(-2.25) * 0.1)
        _ = fi6
        _ = v7
    }
    var v8 float32 = float32(1.5)
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 uint8 = 192
        var v11 uint32 = 65535
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: 0, F1: -603, F2: 100}
    fmt.Println(v3, v1)
    v13 := []int{2147483647, 128, -1}
    if false {
        for fi14 := 0; fi14 < 15; fi14++ {
            if (!false) {
                var v15 int32 = 42
                fmt.Println(v4, v2)
                switch ("foo bar" + "a") {
                case "\n\t":
                    fmt.Println(fi14)
                case "":
                    switch ("" + ("a" + "a")) {
                    case "a":
                        switch (fi14 ^ int(v4)) {
                        case 539:
                            for fi16 := 0; fi16 < 12; fi16++ {
                                _ = len(v13)
                                if (!false) {
                                    var v17 uint64 = 771
                                    var v18 bool = ((^(v2)) >= (v2 - v2))
                                    v19 := GenS2{F0: 571, F1: 2}
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                }
                                var v20 float64 = 1.5
                                v21 := GenS1{F0: -498, F1: v4, F2: 127}
                                switch (-(fi16)) {
                                case 9223372036854775807:
                                    v13 = append(v13, (fi16 ^ (fi16 >> (((v2 / ((255) | 1))) & 63))))
                                default:
                                    if false {
                                        var v22 uint = (v2 | 2)
                                        switch "hello" {
                                        case "":
                                            v15 = v15
                                        case "\n\t":
                                            if (1 != 1) {
                                                if (true || (1 == 0)) {
                                                    if len(v13) > 0 {
                                                        v13[0] = (fi14 + fi14)
                                                    }
                                                    switch "hello" {
                                                    case "":
                                                        if len(v13) >= 2 {
                                                            v13 = v13[1:len(v13)]
                                                        }
                                                    case "a":
                                                        if true {
                                                            v3 += (3.14159 + (0.1 / v20))
                                                        }
                                                        v22 = 4294967295
                                                        v13 = append(v13, (^(fi14)))
                                                        v1 += (0.1 - v3)
                                                        fmt.Println(v1, v2)
                                                        if len(v13) >= 2 {
                                                            v13 = v13[1:len(v13)]
                                                        }
                                                    case "hello":
                                                        var v23 uint8 = 128
                                                        _ = v23
                                                    }
                                                    switch fi14 {
                                                    case 464:
                                                        if false {
                                                            fmt.Println(v15, v2)
                                                            v24 := GenS1{F0: (v15 + (v15 + 42)), F1: 0, F2: 0}
                                                            switch -1 {
                                                            case 2147483647:
                                                                v22 += (v22 % ((v22) | 1))
                                                            default:
                                                                var v25 float64 = v1
                                                                _ = v25
                                                            }
                                                            v2 -= (^(v2))
                                                            v1 -= float64(v8)
                                                            _ = v24
                                                        }
                                                    case 127:
                                                        v15 = (v15 & (v15 | (v15 ^ (v15 % ((v15) | 1)))))
                                                        if (2147483647 == 465) {
                                                            if false {
                                                                if len(v13) > 0 {
                                                                    v13[0] = fi14
                                                                }
                                                            } else {
                                                                var v26 string = "hello"
                                                                var v27 uint = v22
                                                                var v28 int8 = -128
                                                                _ = v26
                                                                _ = v27
                                                                _ = v28
                                                            }
                                                            fmt.Println(v8, v8, v4)
                                                        }
                                                        v1 += v1
                                                    }
                                                    v29 := GenS1{F0: 0, F1: (v4 % ((-32768) | 1)), F2: 100}
                                                    _ = v29
                                                } else {
                                                    fmt.Println(v2, v8, v2)
                                                    v30 := GenS2{F0: 843, F1: 255}
                                                    _ = v30
                                                }
                                            } else {
                                                if true {
                                                    v31 := GenS3{F0: 126, F1: v2, F2: (v2 ^ v2)}
                                                    _ = v31
                                                } else {
                                                    v22 = v2
                                                }
                                                v32 := GenS1{F0: v15, F1: 1, F2: 100}
                                                fmt.Println(fi14)
                                                var v33 uint32 = 998
                                                _ = len(v13)
                                                v4 -= (v4 + (-(v4)))
                                                _ = v32
                                                _ = v33
                                            }
                                        case "foo bar":
                                            v20 += v20
                                        }
                                        if false {
                                            var v34 float32 = float32(v3)
                                            _ = v34
                                        } else {
                                            if ((v15 | v15) >= (v15 * v15)) {
                                                if len(v13) > 0 {
                                                    v13[0] = (fi14 % ((-1) | 1))
                                                }
                                                switch 128 {
                                                case 128:
                                                    v4 = (-(v4))
                                                case 127:
                                                    var v35 int16 = 432
                                                    _ = v35
                                                default:
                                                    if false {
                                                        var v36 uint32 = 1
                                                        v8 -= (v8 * (0.1 - (v8 * v8)))
                                                        _ = v36
                                                    } else {
                                                        var v37 int = (-(fi14))
                                                        var v38 uint8 = 128
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            v39 := GenS3{F0: -22, F1: 554, F2: v22}
                                                            _ = v39
                                                        case "foo bar":
                                                            v1 = v20
                                                        default:
                                                            if len(v13) >= 2 {
                                                                v13 = v13[1:len(v13)]
                                                            }
                                                        }
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                    v3 += ((3.14159 + 0.0) / 3.14159)
                                                    switch "a" {
                                                    case "a":
                                                        v40 := GenS1{F0: 2147483647, F1: 1000, F2: 106}
                                                        _ = v40
                                                    case "":
                                                        v13 = append(v13, fi14)
                                                    case "hello":
                                                        switch (fi16 & (fi14 / ((fi16) | 1))) {
                                                        case 1:
                                                            if true {
                                                                fmt.Println(fi14, v4)
                                                            }
                                                            var v41 int = 0
                                                            _ = v41
                                                        case -1:
                                                            if (65535 == 831) {
                                                                switch (fi16 ^ (-(fi16))) {
                                                                case 128:
                                                                    v15 += (v15 | v15)
                                                                    v13 = append(v13, (^(fi16)))
                                                                    var v42 uint64 = 0
                                                                    _ = v42
                                                                case 39:
                                                                    v15 -= (v15 & (v15 + 42))
                                                                case 894:
                                                                    if (0 != 413) {
                                                                        switch ((("hello" + ("" + "\n\t")) + "\n\t") + (("" + (("a" + (((("" + "hello") + "a") + "\n\t") + ("" + ("" + "foo bar")))) + "a")) + "\n\t")) {
                                                                        case "foo bar":
                                                                            v1 += (0.1 * float64(v8))
                                                                        }
                                                                        v43 := GenS4{F0: 71}
                                                                        _ = v43
                                                                    } else {
                                                                        if false {
                                                                            if (2 > 1) {
                                                                                v44 := GenS1{F0: v15, F1: 0, F2: -128}
                                                                                _ = v44
                                                                            } else {
                                                                                v13 = append(v13, (-(fi16)))
                                                                                var v45 int64 = 1
                                                                                _ = v45
                                                                            }
                                                                            if len(v13) >= 2 {
                                                                                v13 = v13[1:len(v13)]
                                                                            }
                                                                            var v46 float32 = -2.25
                                                                            _ = v46
                                                                        } else {
                                                                            var v47 int8 = 100
                                                                            var v48 uint16 = 61
                                                                            v22 = v2
                                                                            _ = v47
                                                                            _ = v48
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    default:
                                                        v49 := GenS2{F0: 1, F1: 708}
                                                        _ = v49
                                                    }
                                                }
                                            } else {
                                                v13 = append(v13, (fi14 * (^(fi16))))
                                                var v50 uint = v2
                                                v8 = v8
                                                _ = v50
                                            }
                                        }
                                        var v51 bool = (true || true)
                                        if len(v13) >= 2 {
                                            v13 = v13[1:len(v13)]
                                        }
                                        _ = v22
                                        _ = v51
                                    }
                                }
                                v52 := GenS5{F0: ((float64(v8) + v20) < v20), F1: v4, F2: v4}
                                v53 := GenS5{F0: false, F1: (v4 >> ((v2) & 63)), F2: v4}
                                v8 = 1.5
                                _ = fi16
                                _ = v20
                                _ = v21
                                _ = v52
                                _ = v53
                            }
                        default:
                            for fi54 := 0; fi54 < 16; fi54++ {
                                var v55 uint = (v2 / (((v2 | v2)) | 1))
                                _ = fi54
                                _ = v55
                            }
                        }
                    case "hello":
                        var v56 uint16 = 65535
                        _ = v56
                    case "":
                        for fi57 := 0; fi57 < 12; fi57++ {
                            v2 = (v2 & 4294967295)
                            switch fi57 {
                            case 42:
                                v58 := GenS4{F0: 1}
                                _ = v58
                            case 0:
                                if false {
                                    var v59 int8 = 42
                                    var v60 uint16 = 0
                                    if ((v4 - -32768) >= 636) {
                                        var v61 float64 = -2.25
                                        if true {
                                            _ = len(v13)
                                            v2 -= uint(v60)
                                            if ((fi14 >> ((4294967295) & 63)) <= fi14) {
                                                v62 := GenS6{F0: v15, F1: v15}
                                                switch fi57 {
                                                case 42:
                                                    v1 += v61
                                                default:
                                                    v13 = append(v13, fi57)
                                                }
                                                v2 = (v2 | (v2 / ((v2) | 1)))
                                                _ = v62
                                            } else {
                                                v63 := GenS7{F0: v4}
                                                _ = v63
                                            }
                                        } else {
                                            if true {
                                                var v64 int64 = 5
                                                _ = v64
                                            }
                                            switch "hello" {
                                            case "foo bar":
                                                var v65 float64 = v61
                                                _ = v65
                                            case "hello":
                                                switch 4294967295 {
                                                case -2147483648:
                                                    v1 -= 0.1
                                                case 9223372036854775807:
                                                    v4 = v4
                                                default:
                                                    if (-2147483648 < (fi14 + (fi14 | (^(fi14))))) {
                                                        if len(v13) > 0 {
                                                            v13[0] = int(v60)
                                                        }
                                                    } else {
                                                        if true {
                                                            if true {
                                                                var v66 uint16 = v60
                                                                v1 += float64(v8)
                                                                _ = v66
                                                            }
                                                            var v67 bool = (v59 != 100)
                                                            _ = v67
                                                        }
                                                        var v68 int16 = (v4 % ((v4) | 1))
                                                        if (4294967295 == 42) {
                                                            if false {
                                                                v69 := GenS8{F0: (v68 ^ 518)}
                                                                v70 := GenS9{F0: (v68 * v68), F1: 749, F2: 42}
                                                                fmt.Println(v3, v1, v61)
                                                                var v71 int16 = v4
                                                                v59 = 1
                                                                if false {
                                                                    var v72 int = (fi14 % ((fi57) | 1))
                                                                    if false {
                                                                        var v73 uint32 = 706
                                                                        _ = v73
                                                                    }
                                                                    _ = v72
                                                                }
                                                                if true {
                                                                    v74 := GenS10{F0: true, F1: "foo bar", F2: false}
                                                                    _ = len(v13)
                                                                    if (!(v8 < 1.5)) {
                                                                        fmt.Println(fi14)
                                                                        v75 := GenS5{F0: true, F1: (v71 - 42), F2: 0}
                                                                        v76 := GenS11{F0: v2, F1: (v60 * (v60 - 42)), F2: (^(v60))}
                                                                        v60 = v60
                                                                        _ = v75
                                                                        _ = v76
                                                                    }
                                                                    _ = v74
                                                                } else {
                                                                    if len(v13) > 0 {
                                                                        v13[0] = (fi14 / ((-1) | 1))
                                                                    }
                                                                }
                                                                var v77 uint = 4294967295
                                                                _ = v69
                                                                _ = v70
                                                                _ = v71
                                                                _ = v77
                                                            }
                                                        } else {
                                                            v15 += (v15 >> (((v2 << (((v2 - v2)) & 63))) & 63))
                                                            var v78 int64 = -9223372036854775808
                                                            _ = v78
                                                        }
                                                        fmt.Println(v61, v61, v8)
                                                        _ = v68
                                                    }
                                                    var v79 uint = uint(v60)
                                                    v8 = ((v8 - 0.0) * (float32((v61 * 1e10)) - v8))
                                                    var v80 int64 = -9223372036854775808
                                                    if len(v13) > 0 {
                                                        v13[0] = (fi14 & 0)
                                                    }
                                                    var v81 int = fi57
                                                    _ = v79
                                                    _ = v80
                                                    _ = v81
                                                }
                                            default:
                                                switch fi57 {
                                                case 128:
                                                    _ = len(v13)
                                                case 32767:
                                                    if false {
                                                        if false {
                                                            var v82 string = "\n\t"
                                                            var v83 float32 = (float32(v61) + -2.25)
                                                            v13 = append(v13, (fi57 * 0))
                                                            _ = v82
                                                            _ = v83
                                                        } else {
                                                            if false {
                                                                v4 += -1
                                                                v60 = (v60 << ((316) & 63))
                                                                if true {
                                                                    fmt.Println(v15)
                                                                    var v84 int = 751
                                                                    v85 := GenS12{F0: (fi57 * v84), F1: 42}
                                                                    _ = v84
                                                                    _ = v85
                                                                } else {
                                                                    var v86 float32 = (-2.25 - (v8 * -2.25))
                                                                    var v87 uint64 = 2
                                                                    _ = v86
                                                                    _ = v87
                                                                }
                                                            }
                                                            if len(v13) > 0 {
                                                                v13[0] = int(v4)
                                                            }
                                                            if true {
                                                                v60 = (v60 << (((v2 ^ (v2 % ((v2) | 1)))) & 63))
                                                                if ((fi14 + (fi14 + (^(fi57)))) > 2147483647) {
                                                                    fmt.Println(v59)
                                                                    v88 := GenS8{F0: (v4 ^ v4)}
                                                                    v59 += v59
                                                                    _ = v88
                                                                } else {
                                                                    var v89 string = ("a" + "a")
                                                                    _ = v89
                                                                }
                                                                var v90 uint = 570
                                                                _ = v90
                                                            } else {
                                                                switch ("\n\t" + (("foo bar" + "\n\t") + "hello")) {
                                                                case "foo bar":
                                                                    v2 += v2
                                                                case "a":
                                                                    var v91 float64 = (float64(v8) + 1e10)
                                                                    _ = v91
                                                                case "\n\t":
                                                                    if (-128 >= v59) {
                                                                        var v92 int = (fi14 & fi14)
                                                                        var v93 uint8 = 0
                                                                        v94 := GenS1{F0: v15, F1: v4, F2: v59}
                                                                        switch (("a" + "hello") + "hello") {
                                                                        case "a":
                                                                            var v95 uint8 = (v93 / ((1) | 1))
                                                                            _ = v95
                                                                        }
                                                                        _ = v92
                                                                        _ = v93
                                                                        _ = v94
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        var v96 int = fi14
                                                        var v97 uint8 = 136
                                                        _ = v96
                                                        _ = v97
                                                    } else {
                                                        if ((!true) && false) {
                                                            v98 := GenS3{F0: (-(v59)), F1: v2, F2: (v2 << (((^(v2))) & 63))}
                                                            v61 = 0.0
                                                            _ = v98
                                                        }
                                                    }
                                                case 8:
                                                    v59 -= v59
                                                }
                                            }
                                        }
                                        _ = v61
                                    } else {
                                        var v99 int64 = -1
                                        v8 = ((v8 - 1.5) - v8)
                                        _ = v99
                                    }
                                    v4 = (v4 & 1)
                                    _ = v59
                                    _ = v60
                                } else {
                                    if len(v13) > 0 {
                                        v13[0] = (fi14 + (fi14 % ((-230) | 1)))
                                    }
                                    if (!(true && true)) {
                                        var v100 uint = 255
                                        _ = v100
                                    } else {
                                        v101 := GenS3{F0: -128, F1: (v2 | 2), F2: (v2 % ((v2) | 1))}
                                        _ = len(v13)
                                        var v102 uint = (^(v2))
                                        var v103 int64 = 578
                                        _ = v101
                                        _ = v102
                                        _ = v103
                                    }
                                    v8 = ((v8 + (3.14159 - -2.25)) * ((((v8 / v8) * v8) - (-2.25 * (v8 + ((v8 + float32((0.0 - float64(v8)))) * v8)))) + 0.0))
                                }
                            case -2147483648:
                                var v104 int8 = 42
                                _ = v104
                            }
                            v1 = float64((-2.25 + 1e10))
                            _ = fi57
                        }
                    default:
                        if (v2 < v2) {
                            switch (fi14 / (((fi14 >> ((v2) & 63))) | 1)) {
                            case 352:
                                if (1.5 < (v8 + -2.25)) {
                                    switch "hello" {
                                    case "a":
                                        fmt.Println(v8)
                                    case "hello":
                                        var v105 int16 = v4
                                        _ = v105
                                    case "\n\t":
                                        var v106 string = ""
                                        _ = v106
                                    }
                                }
                            case 296:
                                v3 += 3.14159
                            default:
                                fmt.Println(v3, v8, v15)
                            }
                        }
                    }
                }
                _ = v15
            } else {
                for fi107 := 0; fi107 < 7; fi107++ {
                    v108 := GenS13{F0: -341, F1: 65535, F2: 588}
                    fmt.Println(v1)
                    fmt.Println(v3)
                    v1 = float64(v8)
                    var v109 string = "hello"
                    _ = fi107
                    _ = v108
                    _ = v109
                }
            }
            v1 = v1
            v110 := GenS9{F0: (v4 - 125), F1: 727, F2: 569}
            _ = fi14
            _ = v110
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
