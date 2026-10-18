package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int32; F1 bool; F2 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8; F1 int; F2 uint64 }
type GenS6 struct { F0 int16; F1 int32; F2 uint }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint16 }

func main() {
    var v1 uint = 65535
    v2 := make(map[string]bool)
    v3 := make(map[string]uint32)
    var v4 float32 = -2.25
    v5 := make(map[int]int)
    v6 := make(map[uint32]uint32)
    v7 := []int{128}
    if true {
        for fi8 := 0; fi8 < 19; fi8++ {
            for fi9 := 0; fi9 < 15; fi9++ {
                fmt.Println(fi9)
                _ = fi9
            }
            var v10 int8 = 0
            v11 := GenS1{F0: 0}
            _ = fi8
            _ = v10
            _ = v11
        }
        v4 -= v4
    }
    if true {
        v12 := GenS2{F0: (2147483647 != 9223372036854775807), F1: 128}
        for fi13 := 0; fi13 < 9; fi13++ {
            delete(v2, "")
            _ = fi13
        }
        _ = v12
    } else {
        if false {
            if true {
                if ((-909 <= 0) && true) {
                    var v14 int = -9223372036854775808
                    _ = v14
                }
                v4 -= v4
            } else {
                if (1 == 255) {
                    if true {
                        v1 -= v1
                    } else {
                        var v15 int32 = -1
                        var v16 bool = (1e10 >= (-2.25 + 1e10))
                        if (0 <= 1) {
                            if (v1 < 42) {
                                switch "a" {
                                case "foo bar":
                                    v16 = v16
                                }
                                _ = len(v5)
                            } else {
                                var v17 int16 = 1
                                if (!v16) {
                                    for fi18 := 0; fi18 < 13; fi18++ {
                                        fmt.Println(fi18, v1)
                                        var v19 uint16 = 269
                                        _ = fi18
                                        _ = v19
                                    }
                                }
                                v17 = 32767
                                _ = v17
                            }
                            if v16 {
                                v20 := GenS3{F0: 2147483647, F1: true, F2: 18446744073709551615}
                                var v21 uint32 = 65535
                                _ = v20
                                _ = v21
                            }
                        } else {
                            var v22 string = "\n\t"
                            if (!v16) {
                                for fi23 := 0; fi23 < 1; fi23++ {
                                    var v24 uint8 = 127
                                    _ = fi23
                                    _ = v24
                                }
                                v7 = append(v7, 1)
                            }
                            for fi25 := 0; fi25 < 8; fi25++ {
                                switch "foo bar" {
                                case "a":
                                    if (v1 != v1) {
                                        var v26 float32 = ((v4 - float32(((1e10 * ((1e10 + (0.0 * 1e10)) + float64(float32((float64(v4) * ((-2.25 * 1e10) / 3.14159)))))) / 1.0))) + float32((1.5 + 1.5)))
                                        var v27 int8 = 1
                                        _ = v26
                                        _ = v27
                                    } else {
                                        v28 := GenS3{F0: -713, F1: (((!(false || false)) || ((128 > fi25) || (42 > -387))) || ((v15 % ((v15) | 1)) == (v15 | v15))), F2: 1}
                                        var v29 float64 = -2.25
                                        if ((true || true) || (!v16)) {
                                            var v30 int32 = v15
                                            _ = v30
                                        } else {
                                            switch (v22 + v22) {
                                            case "hello":
                                                v31 := GenS1{F0: 0}
                                                _ = v31
                                            case "\n\t":
                                                for fi32 := 0; fi32 < 2; fi32++ {
                                                    v4 -= -2.25
                                                    v1 = 268
                                                    v16 = false
                                                    v33 := GenS1{F0: 42}
                                                    _ = len(v2)
                                                    v34 := GenS2{F0: (!false), F1: 25}
                                                    _ = fi32
                                                    _ = v33
                                                    _ = v34
                                                }
                                            default:
                                                v35 := v2[""]
                                                _ = v35
                                            }
                                            var v36 uint8 = 26
                                            _ = v36
                                        }
                                        _ = v28
                                        _ = v29
                                    }
                                }
                                v37 := GenS3{F0: 375, F1: (-119 < 127), F2: 79}
                                _ = fi25
                                _ = v37
                            }
                            _ = v22
                        }
                        _ = v15
                        _ = v16
                    }
                }
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                var v38 uint64 = 304
                _ = v38
            }
            if true {
                v39 := GenS1{F0: 1}
                v1 -= (v1 >> ((255) & 63))
                if ("hello" >= "a") {
                    v1 -= 1
                    var v40 string = "a"
                    if false {
                        switch 113 {
                        case -153:
                            for fi41 := 0; fi41 < 3; fi41++ {
                                var v42 uint64 = 42
                                _ = fi41
                                _ = v42
                            }
                        case -2147483648:
                            var v43 int16 = -1
                            _ = v43
                        case 82:
                            fmt.Println(v1, v4, v1)
                        }
                    } else {
                        v44 := GenS4{F0: 2147483648}
                        var v45 bool = (!(false || true))
                        v46 := GenS3{F0: -2147483648, F1: false, F2: 2}
                        _ = v44
                        _ = v45
                        _ = v46
                    }
                    _ = v40
                } else {
                    v47 := GenS5{F0: 127, F1: 843, F2: 2}
                    _ = len(v2)
                    _ = v47
                }
                _ = v39
            } else {
                v4 = float32(float64(float32(1.5)))
                for fi48 := 0; fi48 < 14; fi48++ {
                    for fi49 := 0; fi49 < 7; fi49++ {
                        var v50 int64 = -480
                        _ = fi49
                        _ = v50
                    }
                    var v51 uint8 = 2
                    fmt.Println(fi48, v51)
                    _ = fi48
                    _ = v51
                }
            }
            for fi52 := 0; fi52 < 20; fi52++ {
                v4 = (-2.25 / (v4 + 1e10))
                fmt.Println(v4, v4, v4)
                _ = fi52
            }
        }
        v4 = (v4 + 3.14159)
        v53 := GenS3{F0: -2147483648, F1: (0.0 >= v4), F2: 1}
        _ = v53
    }
    if (false && false) {
        v54 := GenS6{F0: 1, F1: 1, F2: v1}
        for fi55 := 0; fi55 < 8; fi55++ {
            var v56 bool = false
            for fi57 := 0; fi57 < 12; fi57++ {
                if (286 != v1) {
                    v1 += 65535
                    switch "\n\t" {
                    case "a":
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    default:
                        v4 -= v4
                    }
                    v58 := GenS4{F0: 498}
                    _ = v58
                }
                _ = fi57
            }
            _ = fi55
            _ = v56
        }
        _ = v54
    } else {
        var v59 uint32 = 787
        _ = v59
    }
    for fi60 := 0; fi60 < 5; fi60++ {
        switch (fi60 >> (((v1 * (v1 + 65535))) & 63)) {
        case -9223372036854775808:
            if false {
                if true {
                    v4 += v4
                    v1 -= (v1 << ((v1) & 63))
                } else {
                    switch "" {
                    case "hello":
                        if false {
                            var v61 int16 = 1000
                            _ = v61
                        }
                    }
                    if ("foo bar" != "foo bar") {
                        var v62 float32 = (float32(0.0) - (float32(0.1) + (v4 * ((1e10 * 1.5) - 0.1))))
                        if (106 != 18446744073709551615) {
                            if ((!(-1 >= -134)) || (!true)) {
                                v63 := v6[836]
                                _ = v63
                            }
                        }
                        _ = v62
                    }
                }
            } else {
                if ((true && false) || (!(true || false))) {
                    v1 -= v1
                }
                for fi64 := 0; fi64 < 15; fi64++ {
                    v1 = v1
                    v65 := GenS6{F0: 42, F1: 0, F2: 18446744073709551615}
                    var v66 float64 = (1.5 / float64(0.1))
                    _ = fi64
                    _ = v65
                    _ = v66
                }
                v67 := GenS4{F0: 636}
                _ = v67
            }
        case -989:
            for fi68 := 0; fi68 < 18; fi68++ {
                v4 = v4
                _ = fi68
            }
        default:
            fmt.Println(v1, v1, v4)
        }
        var v69 uint = (^(v1))
        _ = fi60
        _ = v69
    }
    for fi70 := 0; fi70 < 15; fi70++ {
        fmt.Println(fi70, fi70)
        _ = fi70
    }
    switch "hello" {
    case "hello":
        if (580 != 2) {
            fmt.Println(v1, v4)
            v1 = (v1 | (v1 - 18446744073709551615))
            switch "\n\t" {
            case "a":
                if (0.1 >= float32(float64(0.1))) {
                    var v71 uint16 = 264
                    switch "hello" {
                    case "hello":
                        v72 := GenS4{F0: 171}
                        _ = v72
                    case "":
                        fmt.Println(v1, v1)
                    case "a":
                        fmt.Println(v71)
                    }
                    _ = len(v2)
                    _ = v71
                }
            case "foo bar":
                v4 = (3.14159 + float32(float64((float32(0.1) + (((v4 + (0.1 - 0.1)) * float32((0.1 - -2.25))) + v4)))))
            case "\n\t":
                for fi73 := 0; fi73 < 9; fi73++ {
                    v74 := GenS1{F0: 42}
                    v75 := v5[(fi73 - -161)]
                    _ = fi73
                    _ = v74
                    _ = v75
                }
            }
        } else {
            fmt.Println(v4, v1)
            for fi76 := 0; fi76 < 18; fi76++ {
                switch fi76 {
                case 9223372036854775807:
                    var v77 float64 = 1.5
                    _ = v77
                case 128:
                    if false {
                        _ = len(v3)
                        fmt.Println(fi76)
                    } else {
                        if len(v7) > 0 {
                            v7[0] = 766
                        }
                    }
                }
                v78 := GenS4{F0: 4294967295}
                switch "hello" {
                case "a":
                    fmt.Println(v1, fi76)
                }
                _ = fi76
                _ = v78
            }
        }
    }
    for fi79 := 0; fi79 < 15; fi79++ {
        switch "a" {
        case "hello":
            for fi80 := 0; fi80 < 19; fi80++ {
                v81 := v5[-721]
                var v82 int = v81
                v83 := GenS7{F0: 856}
                _ = fi80
                _ = v81
                _ = v82
                _ = v83
            }
        default:
            var v84 int32 = 0
            _ = v84
        }
        var v85 uint64 = 193
        _ = fi79
        _ = v85
    }
    v86 := GenS8{F0: 65535}
    if (!true) {
        v1 = 18446744073709551615
        for fi87 := 0; fi87 < 20; fi87++ {
            v88 := GenS1{F0: 2}
            v4 += (float32(((3.14159 * 0.1) / float64(v4))) - v4)
            _ = fi87
            _ = v88
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v86
}
