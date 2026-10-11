package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8; F1 int; F2 int64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int8; F1 int32; F2 int16 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int32; F1 uint; F2 uint8 }
type GenS11 struct { F0 bool; F1 uint; F2 uint8 }
type GenS12 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS13 struct { F0 uint32; F1 int32 }

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{1}
    var v3 uint8 = 127
    var v4 float64 = 0.0
    var v5 int8 = 1
    v6 := make(map[string]string)
    for fi7 := 0; fi7 < 13; fi7++ {
        if false {
            if true {
                switch -8 {
                case -1:
                    var v8 int = 42
                    _ = v8
                case 32767:
                    v5 = v5
                default:
                    v9 := GenS1{F0: 960, F1: 189}
                    _ = v9
                }
                var v10 uint8 = v3
                if ((!false) && ((v5 - (v5 + (v5 % ((v5) | 1)))) < (^(v5)))) {
                    fmt.Println(v4)
                    for fi11 := 0; fi11 < 12; fi11++ {
                        var v12 int32 = -1
                        fmt.Println(v5)
                        _ = fi11
                        _ = v12
                    }
                }
                _ = v10
            } else {
                v3 -= v3
            }
            if len(v2) > 0 {
                v2[0] = 1
            }
            if false {
                v4 -= float64(1e10)
                switch (fi7 & -9223372036854775808) {
                case -1:
                    if false {
                        var v13 int8 = v5
                        _ = v13
                    }
                default:
                    var v14 string = "foo bar"
                    _ = v14
                }
                var v15 float64 = v4
                _ = v15
            } else {
                fmt.Println(fi7, fi7, v3)
                v3 += v3
                if ((896 < 32767) || (float32(0.0) >= (3.14159 - 1.5))) {
                    v5 = int8(fi7)
                    if (true || true) {
                        if (0 >= 879) {
                            var v16 int32 = 42
                            v17 := GenS2{F0: 327}
                            switch "foo bar" {
                            case "\n\t":
                                for fi18 := 0; fi18 < 13; fi18++ {
                                    switch "a" {
                                    case "foo bar":
                                        v4 = ((v4 + (v4 - v4)) + v4)
                                    case "hello":
                                        if len(v2) > 0 {
                                            v2[0] = 2147483647
                                        }
                                    default:
                                        var v19 uint64 = 18446744073709551615
                                        _ = v19
                                    }
                                    _ = fi18
                                }
                            case "foo bar":
                                fmt.Println(v3, v5)
                            }
                            _ = v16
                            _ = v17
                        }
                    } else {
                        fmt.Println(v5, v4, v4)
                        v20 := GenS2{F0: 1}
                        switch (-(fi7)) {
                        case -9223372036854775808:
                            fmt.Println(v5, v3, v3)
                        case -196:
                            v21 := v1["a"]
                            _ = v21
                        case 2147483647:
                            var v22 uint64 = 1
                            _ = v22
                        }
                        _ = v20
                    }
                    if true {
                        _ = len(v1)
                        for fi23 := 0; fi23 < 4; fi23++ {
                            if false {
                                var v24 uint = 2
                                _ = v24
                            }
                            v25 := GenS1{F0: 42, F1: 255}
                            _ = fi23
                            _ = v25
                        }
                    } else {
                        if false {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v26 int8 = (v5 % ((-55) | 1))
                            fmt.Println(fi7, v4)
                            _ = v26
                        } else {
                            v27 := GenS1{F0: -32768, F1: 0}
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v5 += (v5 / ((0) | 1))
                            _ = v27
                        }
                        var v28 int = 32768
                        var v29 bool = (1.5 != (1e10 / (0.0 - 1e10)))
                        _ = v28
                        _ = v29
                    }
                } else {
                    v4 += v4
                    for fi30 := 0; fi30 < 10; fi30++ {
                        var v31 uint8 = 128
                        _ = fi30
                        _ = v31
                    }
                    if false {
                        for fi32 := 0; fi32 < 7; fi32++ {
                            if true {
                                v33 := v1[""]
                                v34 := v1[("" + "\n\t")]
                                _ = v33
                                _ = v34
                            } else {
                                fmt.Println(v4)
                                if true {
                                    v3 = v3
                                    var v35 uint = 18446744073709551615
                                    _ = v35
                                } else {
                                    switch (("a" + (((("" + ("a" + "a")) + "foo bar") + ("\n\t" + "")) + "a")) + "foo bar") {
                                    case "foo bar":
                                        var v36 uint16 = 255
                                        _ = v36
                                    case "":
                                        v3 += v3
                                    default:
                                        var v37 uint16 = 1
                                        _ = v37
                                    }
                                    v38 := GenS3{F0: 1, F1: 4294967295, F2: (^(v3))}
                                    _ = v38
                                }
                                fmt.Println(v3, fi32, v3)
                            }
                            v39 := GenS3{F0: 0, F1: 4294967295, F2: v3}
                            v5 = -1
                            _ = fi32
                            _ = v39
                        }
                        v40 := v1["hello"]
                        _ = v40
                    } else {
                        if (("\n\t" != "a") || true) {
                            if true {
                                if false {
                                    fmt.Println(v3, v3)
                                    fmt.Println(v4)
                                    v3 += v3
                                }
                                var v41 uint16 = 1
                                switch fi7 {
                                case -2147483648:
                                    v5 -= (-(v5))
                                case 1:
                                    v4 -= ((v4 - (0.0 - 1e10)) * v4)
                                case 32768:
                                    var v42 uint = 2
                                    _ = v42
                                }
                                _ = v41
                            }
                            for fi43 := 0; fi43 < 12; fi43++ {
                                var v44 uint32 = 2147483648
                                if true {
                                    if true {
                                        if ((!false) && (255 >= 65535)) {
                                            var v45 float64 = 0.1
                                            _ = len(v6)
                                            switch (fi7 / ((fi7) | 1)) {
                                            case 9223372036854775807:
                                                v46 := GenS1{F0: -1, F1: 42}
                                                _ = v46
                                            case -9223372036854775808:
                                                var v47 float64 = v45
                                                _ = v47
                                            }
                                            _ = v45
                                        }
                                    }
                                    fmt.Println(v44, v5, v44)
                                } else {
                                    v1[(("foo bar" + "hello") + "a")] = (((1e10 * (0.0 * -2.25)) - 0.1) > float32(v4))
                                }
                                if (!false) {
                                    switch 0 {
                                    case 445:
                                        v48 := GenS4{F0: 42, F1: 65535, F2: 42}
                                        _ = v48
                                    case 2147483647:
                                        v49 := GenS4{F0: 4294967295, F1: 65535, F2: uint64(v44)}
                                        _ = v49
                                    case 128:
                                        v44 = v44
                                    }
                                    v50 := GenS3{F0: 665, F1: v44, F2: 95}
                                    switch (fi43 & fi7) {
                                    case -9223372036854775808:
                                        if true {
                                            v1[("foo bar" + "foo bar")] = true
                                            switch "" {
                                            case "hello":
                                                switch (fi43 << ((43) & 63)) {
                                                case -2147483648:
                                                    fmt.Println(v5, fi43)
                                                case -882:
                                                    v51 := GenS5{F0: 42}
                                                    _ = v51
                                                default:
                                                    v4 = v4
                                                }
                                            case "":
                                                var v52 uint16 = 32767
                                                _ = v52
                                            }
                                            if (1 < 2) {
                                                var v53 uint = 18446744073709551615
                                                _ = v53
                                            } else {
                                                var v54 float32 = 0.0
                                                _ = v54
                                            }
                                            v3 = v3
                                            _ = len(v1)
                                            v5 = int8(fi43)
                                            var v55 float64 = ((v4 + v4) / (0.0 / (float64(-2.25) + ((0.1 - (v4 + v4)) * 3.14159))))
                                            _ = v55
                                        } else {
                                            var v56 bool = false
                                            _ = v56
                                        }
                                    }
                                    _ = v50
                                } else {
                                    if (!(42 <= (v44 - v44))) {
                                        var v57 string = ("foo bar" + "hello")
                                        v57 = "a"
                                        var v58 uint8 = v3
                                        _ = v57
                                        _ = v58
                                    }
                                    fmt.Println(v44)
                                }
                                var v59 int = -9223372036854775808
                                var v60 string = (("\n\t" + (("hello" + "foo bar") + "\n\t")) + "")
                                _ = fi43
                                _ = v44
                                _ = v59
                                _ = v60
                            }
                            fmt.Println(fi7, v4)
                        }
                        for fi61 := 0; fi61 < 8; fi61++ {
                            v2 = append(v2, 1)
                            _ = fi61
                        }
                        if (!((1 < 1) && (!false))) {
                            v62 := GenS6{F0: 42, F1: fi7, F2: 9223372036854775807}
                            for fi63 := 0; fi63 < 2; fi63++ {
                                var v64 uint64 = 1
                                var v65 uint64 = 116
                                if (255 < 307) {
                                    v3 -= v3
                                    fmt.Println(v4)
                                    v66 := GenS7{F0: "\n\t"}
                                    _ = v66
                                } else {
                                    var v67 int32 = -2147483648
                                    _ = len(v2)
                                    v65 = (v65 >> ((609) & 63))
                                    v68 := GenS3{F0: 867, F1: 42, F2: (^(v3))}
                                    _ = v67
                                    _ = v68
                                }
                                _ = fi63
                                _ = v64
                                _ = v65
                            }
                            _ = v62
                        }
                    }
                }
            }
        }
        if ((fi7 << ((65535) & 63)) >= 1) {
            for fi69 := 0; fi69 < 5; fi69++ {
                fmt.Println(fi7, v5, v3)
                if (!(362 == 266)) {
                    v70 := GenS7{F0: "hello"}
                    var v71 float64 = v4
                    _ = v70
                    _ = v71
                }
                var v72 uint16 = 667
                _ = fi69
                _ = v72
            }
        }
        _ = fi7
    }
    var v73 int8 = (v5 % (((^(v5))) | 1))
    for fi74 := 0; fi74 < 14; fi74++ {
        v6["foo bar"] = ("\n\t" + ("" + ("hello" + "hello")))
        v4 = (3.14159 * v4)
        _ = fi74
    }
    for fi75 := 0; fi75 < 11; fi75++ {
        fmt.Println(v3)
        if false {
            var v76 float64 = (float64(3.14159) / -2.25)
            _ = v76
        }
        _ = fi75
    }
    fmt.Println(v3, v4, v5)
    var v77 int8 = 100
    if false {
        delete(v6, "")
    }
    if false {
        v78 := GenS7{F0: ("foo bar" + "")}
        v4 = ((v4 + 3.14159) / (1.5 - ((1.5 + v4) / ((v4 / (v4 * v4)) - v4))))
        _ = v78
    }
    if false {
        if true {
            _ = len(v6)
            var v79 bool = (1 >= 65535)
            _ = v79
        } else {
            if (42 == 2) {
                for fi80 := 0; fi80 < 19; fi80++ {
                    if true {
                        var v81 string = "hello"
                        if true {
                            var v82 float64 = 3.14159
                            if (2147483648 == 54) {
                                for fi83 := 0; fi83 < 1; fi83++ {
                                    v5 -= v73
                                    v5 = v77
                                    _ = fi83
                                }
                                var v84 uint16 = 65535
                                _ = v84
                            }
                            fmt.Println(v3, v81)
                            _ = v82
                        } else {
                            var v85 uint = 4294967295
                            var v86 string = ((v81 + "") + "foo bar")
                            v87 := GenS8{F0: (v77 * (v73 & v77)), F1: 42, F2: -1}
                            _ = v85
                            _ = v86
                            _ = v87
                        }
                        switch v81 {
                        case "\n\t":
                            switch v81 {
                            case "":
                                v88 := v6[""]
                                _ = v88
                            default:
                                fmt.Println(v73)
                            }
                        case "a":
                            for fi89 := 0; fi89 < 20; fi89++ {
                                v73 -= (v73 % (((v5 + (v73 >> ((65535) & 63)))) | 1))
                                var v90 float32 = 1.5
                                switch fi80 {
                                case -2147483648:
                                    switch fi80 {
                                    case 32768:
                                        var v91 int = fi80
                                        _ = v91
                                    case 714:
                                        v5 = (v5 >> ((730) & 63))
                                    default:
                                        switch fi89 {
                                        case 32767:
                                            fmt.Println(v77, v73)
                                        case -709:
                                            v92 := GenS9{F0: v77}
                                            _ = v92
                                        case 128:
                                            v93 := GenS4{F0: 42, F1: 4294967295, F2: 18446744073709551615}
                                            _ = v93
                                        }
                                    }
                                case 32767:
                                    fmt.Println(fi89, v81, fi89)
                                case 32768:
                                    var v94 int = fi80
                                    _ = v94
                                }
                                _ = fi89
                                _ = v90
                            }
                        default:
                            v95 := GenS8{F0: 0, F1: -2147483648, F2: 247}
                            _ = v95
                        }
                        _ = v81
                    } else {
                        for fi96 := 0; fi96 < 2; fi96++ {
                            switch 32767 {
                            case -9223372036854775808:
                                v97 := GenS10{F0: 42, F1: 0, F2: v3}
                                _ = v97
                            case -678:
                                if false {
                                    if (false || false) {
                                        var v98 int8 = 0
                                        _ = v98
                                    } else {
                                        v99 := GenS4{F0: 749, F1: 65535, F2: 806}
                                        if true {
                                            var v100 bool = false
                                            v101 := GenS11{F0: ((fi80 <= fi80) || true), F1: 11, F2: v3}
                                            _ = v100
                                            _ = v101
                                        }
                                        _ = v99
                                    }
                                } else {
                                    var v102 int16 = 1000
                                    var v103 int32 = 441
                                    fmt.Println(fi80)
                                    _ = v102
                                    _ = v103
                                }
                            }
                            _ = fi96
                        }
                        if ((v3 << ((42) & 63)) != 31) {
                            if true {
                                v77 -= v73
                            } else {
                                if true {
                                    var v104 uint32 = 2
                                    v77 += 127
                                    _ = v104
                                }
                                for fi105 := 0; fi105 < 6; fi105++ {
                                    v1["hello"] = (v5 != (v73 >> ((27) & 63)))
                                    _ = fi105
                                }
                            }
                        }
                        v5 += (v77 % ((1) | 1))
                    }
                    v73 = -123
                    v3 = (v3 % ((v3) | 1))
                    _ = fi80
                }
                v106 := GenS12{F0: 776, F1: 255, F2: v3}
                if (((v4 * -2.25) / v4) <= (v4 - ((v4 - v4) - v4))) {
                    for fi107 := 0; fi107 < 7; fi107++ {
                        switch (fi107 << ((667) & 63)) {
                        case 645:
                            var v108 uint64 = 194
                            _ = v108
                        case 0:
                            for fi109 := 0; fi109 < 2; fi109++ {
                                var v110 int64 = 42
                                _ = fi109
                                _ = v110
                            }
                        case 330:
                            switch "" {
                            case "hello":
                                fmt.Println(v5)
                            case "":
                                for fi111 := 0; fi111 < 10; fi111++ {
                                    fmt.Println(v4, v4)
                                    switch "" {
                                    case "":
                                        var v112 string = (("foo bar" + "\n\t") + "\n\t")
                                        _ = v112
                                    default:
                                        var v113 uint32 = 0
                                        _ = v113
                                    }
                                    fmt.Println(v77, v5)
                                    _ = fi111
                                }
                            case "foo bar":
                                v114 := GenS13{F0: 57, F1: 42}
                                _ = v114
                            default:
                                fmt.Println(v5, fi107, v3)
                            }
                        }
                        switch 32768 {
                        case 9223372036854775807:
                            fmt.Println(v5, v4)
                        }
                        _ = fi107
                    }
                    var v115 uint = 2
                    _ = v115
                } else {
                    if false {
                        var v116 uint64 = 431
                        _ = v116
                    }
                }
                _ = v106
            }
            var v117 int8 = v73
            var v118 uint32 = 1
            _ = v117
            _ = v118
        }
    }
    v119 := GenS9{F0: (v73 / ((v77) | 1))}
    fmt.Println(v4, v4, v77)
    if false {
        if len(v2) > 0 {
            v2[0] = 620
        }
        switch 252 {
        case -1:
            var v120 uint = 0
            _ = v120
        case 9223372036854775807:
            switch -1 {
            case -178:
                fmt.Println(v4, v73)
            }
        }
    } else {
        v121 := GenS3{F0: 65535, F1: 65535, F2: v3}
        _ = v121
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v73:", v73)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v119
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v77
}
