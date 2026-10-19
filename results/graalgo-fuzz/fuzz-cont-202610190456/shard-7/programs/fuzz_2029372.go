package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int32; F1 int32; F2 int32 }
type GenS7 struct { F0 bool; F1 uint; F2 uint8 }

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]string)
    switch -340 {
    case -9223372036854775808:
        v1 = append(v1, 2)
    case 2147483647:
        v4 := GenS1{F0: 4294967295, F1: 0}
        _ = v4
    default:
        v5 := GenS2{F0: 1}
        _ = v5
    }
    switch -1 {
    case -972:
        var v6 int32 = -2
        _ = v6
    }
    var v7 uint16 = 1
    delete(v3, 2)
    var v8 string = "\n\t"
    for fi9 := 0; fi9 < 2; fi9++ {
        if (839 != -1) {
            v8 = "\n\t"
            switch fi9 {
            case -1:
                for fi10 := 0; fi10 < 8; fi10++ {
                    v8 = (v8 + "")
                    _ = fi10
                }
            case 793:
                v11 := GenS2{F0: 128}
                _ = v11
            default:
                var v12 int = fi9
                _ = v12
            }
        }
        _ = fi9
    }
    var v13 float64 = 3.14159
    v3[0] = v8
    if true {
        v14 := GenS3{F0: "hello"}
        v13 -= (float64(float32((v13 - v13))) * (v13 / v13))
        _ = v14
    } else {
        var v15 uint8 = 211
        _ = v15
    }
    if false {
        var v16 string = "hello"
        _ = v16
    } else {
        switch (("a" + v8) + v8) {
        case "\n\t":
            if false {
                var v17 int = 760
                if false {
                    if (!true) {
                        if true {
                            v7 = (v7 | (v7 * (v7 >> ((0) & 63))))
                            var v18 uint16 = v7
                            v18 = v7
                            _ = v18
                        }
                        if false {
                            if (((1e10 * 1e10) <= (float32((((v13 * (((v13 - 0.0) + 0.1) / float64(-2.25))) + v13) / 1.5)) / float32(v13))) && (true || (!(127 > -1)))) {
                                for fi19 := 0; fi19 < 18; fi19++ {
                                    if (!(65535 > 1)) {
                                        v13 -= 1.5
                                    } else {
                                        for fi20 := 0; fi20 < 16; fi20++ {
                                            switch (v8 + "a") {
                                            case "a":
                                                if false {
                                                    v3[0] = ((((v8 + v8) + (v8 + v8)) + v8) + (("foo bar" + "hello") + v8))
                                                    v17 += fi20
                                                    var v21 uint64 = 4294967295
                                                    _ = v21
                                                }
                                            case "\n\t":
                                                v17 = 4294967295
                                            case "hello":
                                                fmt.Println(fi20)
                                                v22 := GenS1{F0: (fi19 >> ((4294967295) & 63)), F1: v17}
                                                v23 := GenS1{F0: 1, F1: (v17 / ((4294967295) | 1))}
                                                fmt.Println(v17)
                                                var v24 uint16 = v7
                                                v8 = v8
                                                _ = v22
                                                _ = v23
                                                _ = v24
                                            }
                                            _ = fi20
                                        }
                                        var v25 bool = (110 >= 1000)
                                        var v26 uint16 = (^(v7))
                                        _ = v25
                                        _ = v26
                                    }
                                    _ = fi19
                                }
                            } else {
                                if true {
                                    for fi27 := 0; fi27 < 16; fi27++ {
                                        if true {
                                            if false {
                                                var v28 int = (fi27 + (v17 % ((v17) | 1)))
                                                _ = v28
                                            }
                                        } else {
                                            _ = len(v3)
                                        }
                                        _ = len(v1)
                                        _ = fi27
                                    }
                                    for fi29 := 0; fi29 < 10; fi29++ {
                                        var v30 float64 = 0.1
                                        for fi31 := 0; fi31 < 12; fi31++ {
                                            delete(v3, 2147483647)
                                            var v32 string = "a"
                                            if false {
                                                var v33 int8 = 127
                                                v32 = v32
                                                switch (^(fi31)) {
                                                case 4294967295:
                                                    if (v13 < v13) {
                                                        var v34 bool = true
                                                        var v35 int32 = -33
                                                        _ = v34
                                                        _ = v35
                                                    }
                                                }
                                                _ = v33
                                            } else {
                                                v30 -= ((v30 / v30) * v13)
                                                v1 = append(v1, 65535)
                                            }
                                            v36 := GenS3{F0: v8}
                                            _ = fi31
                                            _ = v32
                                            _ = v36
                                        }
                                        _ = fi29
                                        _ = v30
                                    }
                                    v37 := GenS3{F0: "a"}
                                    _ = v37
                                }
                                v38 := GenS4{F0: 255, F1: (v17 % ((v17) | 1))}
                                _ = v38
                            }
                            for fi39 := 0; fi39 < 1; fi39++ {
                                var v40 float32 = 1e10
                                _ = fi39
                                _ = v40
                            }
                            switch ("a" + v8) {
                            case "a":
                                if (249 <= 18446744073709551615) {
                                    fmt.Println(v13)
                                    fmt.Println(v13, v17, v8)
                                } else {
                                    v41 := GenS3{F0: v8}
                                    if len(v1) > 0 {
                                        v1[0] = 54
                                    }
                                    _ = v41
                                }
                            case "foo bar":
                                v13 = v13
                            case "":
                                v3[65535] = (v8 + v8)
                            default:
                                var v42 uint16 = v7
                                _ = v42
                            }
                        } else {
                            var v43 uint = 4294967295
                            _ = v43
                        }
                    }
                    for fi44 := 0; fi44 < 6; fi44++ {
                        for fi45 := 0; fi45 < 9; fi45++ {
                            switch v8 {
                            case "\n\t":
                                switch ("a" + (v8 + v8)) {
                                case "\n\t":
                                    if false {
                                        var v46 int32 = 2147483647
                                        switch fi45 {
                                        case -9223372036854775808:
                                            var v47 int = fi45
                                            _ = v47
                                        }
                                        _ = v46
                                    } else {
                                        v48 := GenS5{F0: 2147483647}
                                        _ = v48
                                    }
                                case "foo bar":
                                    if len(v1) > 0 {
                                        v1[0] = 650
                                    }
                                }
                            default:
                                if len(v1) > 0 {
                                    v1[0] = 0
                                }
                            }
                            _ = fi45
                        }
                        var v49 int64 = 551
                        for fi50 := 0; fi50 < 7; fi50++ {
                            switch "" {
                            case "":
                                switch fi44 {
                                case 32768:
                                    var v51 int64 = v49
                                    _ = v51
                                case 540:
                                    v52 := GenS6{F0: -750, F1: 743, F2: 0}
                                    _ = v52
                                case 128:
                                    v53 := GenS2{F0: 103}
                                    _ = v53
                                default:
                                    fmt.Println(v13)
                                }
                            case "\n\t":
                                fmt.Println(v17)
                            default:
                                if ((true || true) || false) {
                                    var v54 int = 1
                                    v55 := GenS2{F0: 90}
                                    switch (-(fi50)) {
                                    case -580:
                                        var v56 uint32 = 2147483648
                                        _ = v56
                                    case 1:
                                        v54 = (v54 & v54)
                                    case -2147483648:
                                        var v57 float32 = float32(v13)
                                        _ = v57
                                    }
                                    _ = v54
                                    _ = v55
                                }
                            }
                            _ = fi50
                        }
                        _ = fi44
                        _ = v49
                    }
                    if true {
                        var v58 uint16 = v7
                        var v59 uint64 = 0
                        if false {
                            for fi60 := 0; fi60 < 10; fi60++ {
                                fmt.Println(v59)
                                _ = fi60
                            }
                        } else {
                            var v61 int8 = 42
                            fmt.Println(v58, v61)
                            _ = v61
                        }
                        _ = v58
                        _ = v59
                    }
                }
                v8 = v8
                _ = v17
            }
        case "a":
            v7 -= (v7 % (((^(v7))) | 1))
        case "hello":
            for fi62 := 0; fi62 < 9; fi62++ {
                for fi63 := 0; fi63 < 20; fi63++ {
                    switch ((v8 + v8) + (v8 + v8)) {
                    case "\n\t":
                        v64 := v2[2]
                        _ = v64
                    case "hello":
                        switch v8 {
                        case "hello":
                            v2[4294967295] = v8
                        }
                    default:
                        var v65 float64 = v13
                        if true {
                            fmt.Println(v7, fi63)
                            switch v8 {
                            case "hello":
                                v65 += -2.25
                                v66 := GenS5{F0: 1}
                                _ = v66
                            }
                        } else {
                            var v67 uint16 = (v7 << ((13) & 63))
                            _ = v67
                        }
                        _ = v65
                    }
                    v3[800] = ("" + v8)
                    v8 = ""
                    _ = fi63
                }
                var v68 int = 1
                _ = fi62
                _ = v68
            }
        default:
            v69 := GenS3{F0: v8}
            _ = v69
        }
        if false {
            var v70 int8 = 124
            var v71 uint8 = 0
            switch -594 {
            case 9223372036854775807:
                fmt.Println(v13, v70, v70)
            case -452:
                if (!true) {
                    v71 = (v71 >> ((18446744073709551615) & 63))
                } else {
                    switch ((v8 + v8) + "foo bar") {
                    case "":
                        v71 += (v71 + v71)
                    default:
                        v13 = (v13 * v13)
                    }
                }
            default:
                var v72 uint = 236
                _ = v72
            }
            _ = v70
            _ = v71
        } else {
            fmt.Println(v13, v8)
            v13 -= (3.14159 - (-2.25 + v13))
            v73 := GenS2{F0: 128}
            _ = v73
        }
        var v74 uint16 = (v7 & v7)
        _ = v74
    }
    v75 := v3[4294967295]
    v8 = v8
    fmt.Println(v7)
    if true {
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch (("" + ("a" + v75)) + "") {
            case "":
                switch "a" {
                case "hello":
                    for fi76 := 0; fi76 < 16; fi76++ {
                        if ((0 < 1) || true) {
                            fmt.Println(v13, v7, fi76)
                        }
                        var v77 int = fi76
                        _ = len(v2)
                        _ = fi76
                        _ = v77
                    }
                default:
                    if true {
                        v78 := GenS2{F0: 127}
                        _ = v78
                    } else {
                        if (false && false) {
                            var v79 uint32 = 0
                            if false {
                                fmt.Println(v75)
                            } else {
                                v79 = (v79 / (((v79 / ((2) | 1))) | 1))
                            }
                            _ = v79
                        } else {
                            v13 = v13
                        }
                        if false {
                            for fi80 := 0; fi80 < 17; fi80++ {
                                switch v8 {
                                case "hello":
                                    if (!false) {
                                        var v81 uint64 = 0
                                        if len(v1) > 0 {
                                            v1[0] = 340
                                        }
                                        if false {
                                            v7 = (v7 - (v7 * uint16(v81)))
                                        }
                                        _ = v81
                                    }
                                case "foo bar":
                                    v75 = "\n\t"
                                }
                                _ = fi80
                            }
                            var v82 uint = 0
                            var v83 bool = false
                            _ = v82
                            _ = v83
                        } else {
                            for fi84 := 0; fi84 < 8; fi84++ {
                                v7 = (v7 - v7)
                                fmt.Println(v13)
                                var v85 bool = (v7 < (v7 * 255))
                                _ = fi84
                                _ = v85
                            }
                            fmt.Println(v75, v7, v13)
                            v13 -= (1e10 * ((((v13 + v13) / v13) - v13) * v13))
                        }
                        for fi86 := 0; fi86 < 14; fi86++ {
                            switch (fi86 & (fi86 * (fi86 ^ fi86))) {
                            case 32768:
                                if (true || (true || ((!true) || (!(65535 >= 42))))) {
                                    fmt.Println(v8, v13, v8)
                                    v87 := v3[2]
                                    _ = v87
                                } else {
                                    fmt.Println(v75, v75)
                                }
                            }
                            if false {
                                var v88 uint32 = 2
                                v89 := GenS7{F0: (0 != 18446744073709551615), F1: 459, F2: 2}
                                v75 = ""
                                _ = v88
                                _ = v89
                            }
                            for fi90 := 0; fi90 < 10; fi90++ {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v8 = "foo bar"
                                fmt.Println(v7)
                                _ = fi90
                            }
                            _ = fi86
                        }
                    }
                }
            }
        } else {
            var v91 int16 = -1
            _ = v91
        }
        v7 += (^(v7))
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    for fi92 := 0; fi92 < 15; fi92++ {
        var v93 uint64 = 1
        fmt.Println(v13, v93)
        if false {
            v94 := GenS1{F0: fi92, F1: fi92}
            _ = v94
        }
        _ = fi92
        _ = v93
    }
    v75 = v75
    var v95 uint32 = 1
    for fi96 := 0; fi96 < 4; fi96++ {
        v13 = 1e10
        _ = fi96
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v75:", v75)
    fmt.Println("v8:", v8)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v7
    _ = v75
    _ = v8
    _ = v95
}
