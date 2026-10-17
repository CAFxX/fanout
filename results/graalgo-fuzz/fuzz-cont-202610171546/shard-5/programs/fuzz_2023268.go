package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 string; F1 uint16; F2 int64 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int64; F1 uint16 }
type GenS13 struct { F0 uint; F1 uint64 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 int64; F1 int }
type GenS16 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS17 struct { F0 int32; F1 uint64 }
type GenS18 struct { F0 int; F1 int16; F2 uint }
func genFn0(p0 uint16) bool {
    fmt.Println(p0)
    if (!true) {
        for fi6 := 0; fi6 < 20; fi6++ {
            if true {
                p0 += p0
                var v7 uint32 = 2
                _ = v7
            } else {
                for fi8 := 0; fi8 < 8; fi8++ {
                    var v9 int16 = 1
                    _ = fi8
                    _ = v9
                }
                v10 := GenS1{F0: 1, F1: 1, F2: 42}
                switch fi6 {
                case 2147483647:
                    if true {
                        v11 := []int{fi6}
                        _ = v11
                    }
                case 0:
                    for fi12 := 0; fi12 < 20; fi12++ {
                        p0 += p0
                        var v13 int16 = 1
                        _ = fi12
                        _ = v13
                    }
                default:
                    v14 := GenS1{F0: -2147483648, F1: 0, F2: 255}
                    _ = v14
                }
                _ = v10
            }
            _ = fi6
        }
        v15 := GenS1{F0: -1, F1: 1, F2: 252}
        fmt.Println(p0)
        _ = v15
    }
    v16 := GenS1{F0: -1, F1: 265, F2: 3}
    _ = p0
    _ = v16
    return (!false)
}

func main() {
    v1 := []string{("foo bar" + "\n\t")}
    v2 := make(map[string]bool)
    var v3 bool = (!false)
    var v4 uint8 = 2
    var v5 uint32 = 65535
    _ = len(v1)
    var v17 uint16 = 65535
    v18 := GenS2{F0: 904, F1: -634}
    if v3 {
        fmt.Println(v4, v3, v5)
    } else {
        if (((("" + (("\n\t" + ("" + "\n\t")) + ("hello" + ("\n\t" + "foo bar")))) + "\n\t") + ("foo bar" + "a")) <= "\n\t") {
            var v19 uint64 = 1
            if (!v3) {
                if v3 {
                    for fi20 := 0; fi20 < 3; fi20++ {
                        for fi21 := 0; fi21 < 2; fi21++ {
                            switch "foo bar" {
                            case "\n\t":
                                switch fi20 {
                                case 2147483647:
                                    var v22 uint32 = (v5 >> ((2) & 63))
                                    _ = v22
                                }
                                if v3 {
                                    if ((^(fi21)) <= 9223372036854775807) {
                                        v23 := GenS2{F0: 559, F1: int64(v17)}
                                        if (!(112 != 8)) {
                                            v17 -= v17
                                            if (0.1 > (-2.25 * 1e10)) {
                                                var v24 int8 = -128
                                                _ = v24
                                            }
                                            v3 = (42 <= 0)
                                        } else {
                                            if true {
                                                v25 := GenS3{F0: "a", F1: (v17 & v17), F2: -682}
                                                _ = v25
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v26 := GenS4{F0: v4, F1: -542}
                                            _ = v26
                                        }
                                        if (true || ((v4 >> ((4294967295) & 63)) > v4)) {
                                            v3 = (!((0 == -553) || (v17 == (v17 - 42))))
                                            fmt.Println(v4)
                                            v19 = (v19 & 1)
                                        }
                                        _ = v23
                                    }
                                    if true {
                                        v1 = append(v1, ("a" + ("" + ("foo bar" + "foo bar"))))
                                        v19 += 703
                                        if (-2147483648 > -2147483648) {
                                            fmt.Println(v3, v3, fi21)
                                            var v27 string = "foo bar"
                                            v28 := v2[(v27 + ((v27 + "") + v27))]
                                            _ = v27
                                            _ = v28
                                        } else {
                                            fmt.Println(v5, v5)
                                        }
                                    } else {
                                        if (1 <= 32767) {
                                            v4 = v4
                                            if true {
                                                fmt.Println(v17, v4)
                                                var v29 int32 = -705
                                                _ = v29
                                            } else {
                                                fmt.Println(v3, v3)
                                                v30 := GenS2{F0: 0, F1: -9223372036854775808}
                                                _ = v30
                                            }
                                            v31 := GenS5{F0: "\n\t"}
                                            _ = v31
                                        } else {
                                            var v32 bool = v3
                                            v33 := GenS1{F0: 745, F1: 1, F2: v4}
                                            _ = v32
                                            _ = v33
                                        }
                                        v5 = v5
                                    }
                                } else {
                                    v34 := GenS6{F0: 0}
                                    v17 += v17
                                    _ = v34
                                }
                                fmt.Println(v19, fi20)
                            case "foo bar":
                                if (!v3) {
                                    v35 := GenS7{F0: (v5 / (((v5 / ((1) | 1))) | 1))}
                                    _ = v35
                                } else {
                                    v17 = (v17 / ((v17) | 1))
                                }
                            case "":
                                v2["hello"] = (!true)
                            default:
                                if ((0.1 + (0.1 * 0.1)) == 1e10) {
                                    fmt.Println(v17, fi21)
                                    v36 := GenS8{F0: 65535, F1: 42, F2: 0}
                                    _ = v36
                                } else {
                                    if v3 {
                                        var v37 bool = v3
                                        v19 -= (v19 / ((v19) | 1))
                                        v17 = 2
                                        var v38 int16 = 1000
                                        fmt.Println(v17, v17, v19)
                                        _ = v37
                                        _ = v38
                                    }
                                }
                            }
                            v39 := GenS7{F0: (^(v5))}
                            switch (((("" + "a") + ("foo bar" + "\n\t")) + "") + "\n\t") {
                            case "":
                                v3 = (2 < (v17 << ((uint(v4)) & 63)))
                            case "hello":
                                var v40 float64 = -2.25
                                _ = v40
                            case "\n\t":
                                fmt.Println(v3, v4, v5)
                            }
                            _ = fi21
                            _ = v39
                        }
                        for fi41 := 0; fi41 < 9; fi41++ {
                            v4 -= 127
                            _ = fi41
                        }
                        _ = fi20
                    }
                    var v42 int16 = -32768
                    _ = v42
                } else {
                    v2["foo bar"] = (-823 > 42)
                }
            } else {
                fmt.Println(v5, v5)
                v5 += v5
            }
            _ = v19
        } else {
            var v43 int = 811
            var v44 bool = (v43 >= (v43 + v43))
            v44 = v44
            _ = v43
            _ = v44
        }
        for fi45 := 0; fi45 < 16; fi45++ {
            if ("\n\t" >= "a") {
                if v3 {
                    v5 = v5
                    delete(v2, ("a" + "hello"))
                } else {
                    v17 = v17
                    delete(v2, "")
                }
                fmt.Println(v5, v5)
                v2["foo bar"] = true
            } else {
                var v46 uint = 2
                _ = v46
            }
            delete(v2, "")
            _ = fi45
        }
    }
    v47 := GenS5{F0: (((("foo bar" + "foo bar") + "foo bar") + "a") + "foo bar")}
    switch "hello" {
    case "\n\t":
        if v3 {
            if (!((-22 == int8(v17)) || (372 < 306))) {
                v48 := GenS1{F0: 167, F1: 564, F2: v4}
                for fi49 := 0; fi49 < 11; fi49++ {
                    v2[("hello" + "foo bar")] = v3
                    var v50 uint64 = 838
                    _ = fi49
                    _ = v50
                }
                _ = v48
            } else {
                if (v3 || v3) {
                    switch 9223372036854775807 {
                    case 32767:
                        for fi51 := 0; fi51 < 8; fi51++ {
                            var v52 float64 = 3.14159
                            if false {
                                v53 := GenS9{F0: ((^(v17)) <= 2)}
                                _ = v53
                            }
                            _ = fi51
                            _ = v52
                        }
                    case -532:
                        for fi54 := 0; fi54 < 11; fi54++ {
                            var v55 int16 = 384
                            for fi56 := 0; fi56 < 13; fi56++ {
                                v55 -= 42
                                if (v5 != (^(v5))) {
                                    if ((fi56 + (fi56 >> ((42) & 63))) > (fi56 / ((fi56) | 1))) {
                                        fmt.Println(fi56, v55)
                                        switch 42 {
                                        case -330:
                                            if len(v1) > 0 {
                                                v1[0] = "a"
                                            }
                                        case -1:
                                            var v57 float64 = 1e10
                                            _ = v57
                                        }
                                    }
                                }
                                v58 := GenS8{F0: 2, F1: -654, F2: (^(v17))}
                                _ = fi56
                                _ = v58
                            }
                            _ = fi54
                            _ = v55
                        }
                    default:
                        var v59 int32 = 1
                        _ = v59
                    }
                    var v60 float64 = 1.5
                    for fi61 := 0; fi61 < 9; fi61++ {
                        v62 := GenS10{F0: 4294967295, F1: 0, F2: uint64(v5)}
                        _ = fi61
                        _ = v62
                    }
                    _ = v60
                } else {
                    for fi63 := 0; fi63 < 4; fi63++ {
                        switch "" {
                        case "foo bar":
                            v64 := GenS7{F0: (v5 - v5)}
                            _ = v64
                        case "a":
                            fmt.Println(fi63, v17)
                        }
                        if (!(!((3.14159 + 3.14159) > ((1e10 + 0.0) / 0.1)))) {
                            fmt.Println(v5)
                            fmt.Println(v4)
                        }
                        _ = fi63
                    }
                    for fi65 := 0; fi65 < 2; fi65++ {
                        v66 := GenS11{F0: 254}
                        _ = fi65
                        _ = v66
                    }
                    var v67 int64 = 9223372036854775807
                    _ = v67
                }
                v3 = false
            }
            v17 -= 965
            for fi68 := 0; fi68 < 2; fi68++ {
                for fi69 := 0; fi69 < 16; fi69++ {
                    switch fi68 {
                    case 0:
                        if ((fi68 + (fi68 >> ((90) & 63))) == (fi69 >> ((255) & 63))) {
                            v4 = (v4 | (v4 >> ((4294967295) & 63)))
                            switch 0 {
                            case 32768:
                                var v70 int8 = 11
                                _ = v70
                            case 617:
                                var v71 uint = 175
                                _ = v71
                            case -223:
                                v72 := GenS11{F0: uint(v5)}
                                _ = v72
                            default:
                                v73 := GenS7{F0: 4294967295}
                                _ = v73
                            }
                        } else {
                            var v74 uint = 403
                            if true {
                                switch 206 {
                                case 1:
                                    fmt.Println(v17, v4, fi69)
                                case 2147483647:
                                    var v75 uint8 = (v4 % ((v4) | 1))
                                    _ = v75
                                case -9223372036854775808:
                                    if ((v3 && true) && v3) {
                                        fmt.Println(v17)
                                        if v3 {
                                            v5 -= (v5 - 65535)
                                            v76 := GenS8{F0: 65535, F1: -428, F2: (v17 << ((v74) & 63))}
                                            _ = v76
                                        }
                                        switch ("" + "foo bar") {
                                        case "a":
                                            switch (fi68 + -2147483648) {
                                            case 4294967295:
                                                if (-1 >= -62) {
                                                    var v77 uint8 = v4
                                                    if (!v3) {
                                                        if len(v1) > 0 {
                                                            v1[0] = "hello"
                                                        }
                                                        v1 = append(v1, ((((("foo bar" + ("a" + "foo bar")) + "a") + "") + "foo bar") + "hello"))
                                                        if false {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            if true {
                                                                fmt.Println(v5, v77, v4)
                                                                v4 = 42
                                                                v78 := v2["hello"]
                                                                _ = v78
                                                            }
                                                            fmt.Println(v5, v17)
                                                        }
                                                        var v79 int = (fi68 & fi68)
                                                        _ = v79
                                                    }
                                                    _ = v77
                                                }
                                            default:
                                                var v80 int = (fi68 % (((fi68 >> ((432) & 63))) | 1))
                                                _ = v80
                                            }
                                        case "hello":
                                            v3 = (!false)
                                        case "\n\t":
                                            var v81 int32 = int32(v5)
                                            _ = v81
                                        default:
                                            if v3 {
                                                v1 = append(v1, "hello")
                                                var v82 int8 = -84
                                                if (v3 || (!("hello" > "hello"))) {
                                                    v5 = v5
                                                }
                                                _ = v82
                                            } else {
                                                fmt.Println(v74, v3, v3)
                                                var v83 int = (fi68 & 1)
                                                _ = v83
                                            }
                                        }
                                        v84 := GenS10{F0: 18446744073709551615, F1: 306, F2: 795}
                                        if ((float32(1e10) * (0.1 * float32((((1e10 - 3.14159) * -2.25) / 1.5)))) <= (3.14159 * 0.1)) {
                                            v3 = (((1.5 / 0.1) * float32((3.14159 / ((-2.25 + 0.1) + 3.14159)))) == 1.5)
                                            fmt.Println(v5, v17, v5)
                                            v85 := GenS12{F0: 9223372036854775807, F1: uint16(v4)}
                                            _ = v85
                                        }
                                        _ = v84
                                    } else {
                                        var v86 uint32 = (v5 + (v5 / (((v5 >> ((v74) & 63))) | 1)))
                                        v87 := GenS6{F0: 127}
                                        _ = v86
                                        _ = v87
                                    }
                                }
                                var v88 int8 = -1
                                _ = v88
                            } else {
                                if true {
                                    var v89 float32 = -2.25
                                    v89 = v89
                                    var v90 int32 = 603
                                    if (!false) {
                                        v2["\n\t"] = ((3.14159 - v89) <= 1.5)
                                        if v3 {
                                            fmt.Println(v17)
                                            switch (("a" + ("foo bar" + "foo bar")) + "\n\t") {
                                            case "\n\t":
                                                var v91 int8 = 100
                                                _ = v91
                                            case "foo bar":
                                                switch "hello" {
                                                case "hello":
                                                    v89 += v89
                                                }
                                            case "":
                                                if v3 {
                                                    v92 := GenS12{F0: 867, F1: uint16(v74)}
                                                    var v93 int32 = (v90 / ((int32(fi68)) | 1))
                                                    v94 := GenS13{F0: v74, F1: 2}
                                                    _ = v92
                                                    _ = v93
                                                    _ = v94
                                                }
                                            }
                                        }
                                        v2[""] = v3
                                    } else {
                                        var v95 uint8 = v4
                                        v96 := GenS5{F0: ((("\n\t" + ("a" + "")) + "\n\t") + "\n\t")}
                                        v97 := GenS1{F0: v90, F1: 0, F2: 128}
                                        v17 = 0
                                        v74 -= 508
                                        _ = v95
                                        _ = v96
                                        _ = v97
                                    }
                                    _ = v89
                                    _ = v90
                                }
                                v5 = v5
                                var v98 uint32 = (v5 + 576)
                                _ = v98
                            }
                            _ = v74
                        }
                    case 965:
                        if true {
                            fmt.Println(v4, v4, v4)
                            v99 := GenS14{F0: v5}
                            v17 -= (v17 + (v17 | v17))
                            var v100 string = (("\n\t" + "\n\t") + "hello")
                            fmt.Println(v5, v4, v100)
                            _ = v100
                            _ = v99
                        }
                    case 128:
                        v2[""] = false
                    }
                    _ = len(v2)
                    _ = fi69
                }
                _ = fi68
            }
        }
    case "hello":
        var v101 uint = uint(v4)
        _ = v101
    case "a":
        var v102 uint8 = 63
        _ = v102
    }
    v4 = (v4 - v4)
    var v103 float32 = (1e10 + 0.1)
    if v3 {
        var v104 int = 629
        switch v104 {
        case -751:
            var v105 uint32 = (^(v5))
            _ = v105
        case 9223372036854775807:
            v106 := GenS10{F0: 65535, F1: 1, F2: 2}
            _ = v106
        case 32767:
            v4 = 1
        }
        for fi107 := 0; fi107 < 2; fi107++ {
            v108 := GenS15{F0: -1, F1: v104}
            v4 -= v4
            _ = fi107
            _ = v108
        }
        _ = v104
    }
    if false {
        for fi109 := 0; fi109 < 17; fi109++ {
            var v110 int16 = 0
            switch ((("" + "\n\t") + "a") + (("hello" + ("foo bar" + ("hello" + ((("a" + "a") + ("hello" + "\n\t")) + "hello")))) + "a")) {
            case "a":
                var v111 uint32 = v5
                _ = v111
            case "\n\t":
                for fi112 := 0; fi112 < 18; fi112++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi112
                }
            default:
                var v113 float64 = ((0.1 * 0.1) + 1e10)
                _ = v113
            }
            var v114 int = (^(fi109))
            _ = fi109
            _ = v110
            _ = v114
        }
    } else {
        v115 := GenS16{F0: 2147483647, F1: 1, F2: (^(v4))}
        v3 = v3
        _ = v115
    }
    switch "hello" {
    case "a":
        for fi116 := 0; fi116 < 8; fi116++ {
            switch fi116 {
            case -2147483648:
                var v117 int64 = 1
                _ = v117
            case -9223372036854775808:
                for fi118 := 0; fi118 < 8; fi118++ {
                    var v119 int16 = -291
                    _ = fi118
                    _ = v119
                }
            case 721:
                v120 := GenS4{F0: 255, F1: 1000}
                _ = v120
            default:
                v121 := GenS17{F0: 42, F1: 490}
                _ = v121
            }
            v3 = false
            _ = fi116
        }
    case "foo bar":
        for fi122 := 0; fi122 < 2; fi122++ {
            v103 = v103
            for fi123 := 0; fi123 < 16; fi123++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi123
            }
            var v124 int = -649
            _ = fi122
            _ = v124
        }
    case "":
        var v125 uint = 255
        _ = v125
    }
    if len(v1) > 0 {
        v1[0] = (("\n\t" + "a") + "")
    }
    delete(v2, ("foo bar" + ((("a" + "") + "a") + "foo bar")))
    fmt.Println(v3, v3)
    switch "foo bar" {
    case "hello":
        delete(v2, ((("a" + "") + ("a" + ("\n\t" + "\n\t"))) + ("" + "hello")))
    case "\n\t":
        if ((0.0 / (float64(v103) * ((3.14159 + 0.1) - 0.0))) < -2.25) {
            v126 := GenS18{F0: 1, F1: -1, F2: 229}
            _ = v126
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v103:", v103)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v103
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
