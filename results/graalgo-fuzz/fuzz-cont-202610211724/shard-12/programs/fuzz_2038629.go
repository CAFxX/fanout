package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 bool }
type GenS2 struct { F0 uint8; F1 uint; F2 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8; F1 string }
type GenS7 struct { F0 int }
type GenS8 struct { F0 string; F1 uint16 }
type GenS9 struct { F0 string; F1 uint32 }
type GenS10 struct { F0 int16; F1 string; F2 int32 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int32 }

func main() {
    var v1 int32 = 0
    v2 := []int{1, 930}
    var v3 int = 128
    var v4 uint8 = 1
    var v5 string = "foo bar"
    if (("" + v5) < v5) {
        v1 = 2147483647
        for fi6 := 0; fi6 < 12; fi6++ {
            for fi7 := 0; fi7 < 2; fi7++ {
                v5 = "hello"
                v1 -= (v1 ^ v1)
                if len(v2) > 0 {
                    v2[0] = fi7
                }
                _ = fi7
            }
            if true {
                for fi8 := 0; fi8 < 7; fi8++ {
                    v3 -= (fi6 << ((18446744073709551615) & 63))
                    switch "a" {
                    case "":
                        var v9 bool = false
                        _ = v9
                    case "\n\t":
                        v2 = append(v2, (fi8 >> ((485) & 63)))
                        switch (v5 + "foo bar") {
                        case "\n\t":
                            var v10 uint = 4294967295
                            _ = v10
                        case "a":
                            if false {
                                var v11 int16 = -91
                                v12 := GenS1{F0: 4294967295, F1: -1, F2: (fi8 < v3)}
                                var v13 int32 = (v1 << ((42) & 63))
                                if (("hello" + "") != ("\n\t" + "a")) {
                                    if len(v2) > 0 {
                                        v2[0] = (fi8 << ((255) & 63))
                                    }
                                    v14 := GenS1{F0: 535, F1: 127, F2: (v11 < v11)}
                                    _ = v14
                                }
                                var v15 int8 = 42
                                _ = v11
                                _ = v12
                                _ = v13
                                _ = v15
                            } else {
                                var v16 bool = (65535 > uint16(v4))
                                _ = v16
                            }
                        case "hello":
                            var v17 bool = true
                            _ = v17
                        default:
                            fmt.Println(v4, v3, fi8)
                        }
                    case "foo bar":
                        switch v3 {
                        case 4294967295:
                            v18 := GenS2{F0: (v4 - (v4 % (((^(v4))) | 1))), F1: 2, F2: v3}
                            if len(v2) > 0 {
                                v2[0] = (fi8 | v3)
                            }
                            _ = v18
                        case -942:
                            _ = len(v2)
                        case 9223372036854775807:
                            var v19 uint32 = 721
                            _ = v19
                        default:
                            var v20 uint = 4294967295
                            _ = v20
                        }
                    }
                    _ = fi8
                }
                switch ("\n\t" + v5) {
                case "hello":
                    var v21 uint64 = 637
                    _ = v21
                case "\n\t":
                    v5 = v5
                default:
                    _ = len(v2)
                }
                if false {
                    v22 := GenS3{F0: (v4 << ((255) & 63))}
                    if len(v2) > 0 {
                        v2[0] = 128
                    }
                    v1 = (v1 ^ v1)
                    _ = v22
                }
            }
            for fi23 := 0; fi23 < 9; fi23++ {
                fmt.Println(v4)
                var v24 int16 = 0
                _ = fi23
                _ = v24
            }
            _ = fi6
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        var v25 int16 = -44
        _ = v25
    }
    if false {
        switch v3 {
        case 1:
            fmt.Println(v5, v5)
        case 32767:
            switch "" {
            case "":
                for fi26 := 0; fi26 < 18; fi26++ {
                    for fi27 := 0; fi27 < 15; fi27++ {
                        v3 -= (v3 * v3)
                        _ = fi27
                    }
                    fmt.Println(v3, v3)
                    _ = fi26
                }
            case "\n\t":
                v28 := GenS1{F0: 344, F1: -86, F2: true}
                _ = v28
            case "a":
                for fi29 := 0; fi29 < 7; fi29++ {
                    var v30 uint = 1
                    if true {
                        fmt.Println(v5, v5, v5)
                        if true {
                            v31 := GenS4{F0: 0, F1: v4, F2: 42}
                            v32 := GenS5{F0: 2}
                            _ = v31
                            _ = v32
                        } else {
                            switch fi29 {
                            case -69:
                                v33 := GenS1{F0: 42, F1: 100, F2: (42 == 1)}
                                _ = v33
                            case 42:
                                var v34 uint = v30
                                _ = v34
                            case -478:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                var v35 int = v3
                                _ = v35
                            }
                            if (true && ((true || (!false)) && (2 > 42))) {
                                if len(v2) > 0 {
                                    v2[0] = int(v1)
                                }
                            } else {
                                v2 = append(v2, int(v4))
                                v4 = v4
                                v36 := GenS6{F0: 0, F1: v5}
                                _ = v36
                            }
                        }
                    } else {
                        v37 := GenS7{F0: (v3 - v3)}
                        switch (fi29 - v3) {
                        case -2147483648:
                            v5 = "a"
                        case 0:
                            v30 = v30
                        case -1:
                            var v38 bool = (!true)
                            _ = v38
                        }
                        var v39 int64 = int64(fi29)
                        _ = v37
                        _ = v39
                    }
                    switch v3 {
                    case 93:
                        switch fi29 {
                        case -282:
                            v40 := GenS1{F0: 2147483648, F1: -128, F2: true}
                            _ = v40
                        case -638:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case -892:
                            v41 := GenS8{F0: v5, F1: 65535}
                            _ = v41
                        }
                    case 9223372036854775807:
                        if ((false && (!true)) || (-616 >= -1)) {
                            var v42 int = (-(fi29))
                            v43 := GenS9{F0: v5, F1: 65535}
                            _ = v42
                            _ = v43
                        }
                    default:
                        if (!((v4 << (((v30 * 18446744073709551615)) & 63)) >= 2)) {
                            if len(v2) > 0 {
                                v2[0] = -9223372036854775808
                            }
                            v44 := GenS2{F0: (v4 & (v4 ^ 128)), F1: 766, F2: (-(v3))}
                            if false {
                                var v45 string = ((("hello" + v5) + "") + v5)
                                _ = v45
                            } else {
                                var v46 uint = v30
                                if ((v1 * (^(v1))) > 293) {
                                    if (-9223372036854775808 == 1) {
                                        var v47 int8 = 42
                                        fmt.Println(v3)
                                        var v48 int32 = (v1 ^ (v1 ^ (v1 >> ((v30) & 63))))
                                        _ = v47
                                        _ = v48
                                    } else {
                                        v4 = v4
                                        var v49 int16 = 409
                                        _ = v49
                                    }
                                }
                                _ = v46
                            }
                            _ = v44
                        }
                    }
                    _ = fi29
                    _ = v30
                }
            }
        }
    } else {
        var v50 int64 = 651
        if (v4 < v4) {
            if (((1 != 2) && (969 == 748)) || true) {
                v5 = v5
                if ((v3 ^ v3) >= v3) {
                    var v51 uint = 42
                    _ = v51
                }
                for fi52 := 0; fi52 < 15; fi52++ {
                    switch fi52 {
                    case 9223372036854775807:
                        if false {
                            for fi53 := 0; fi53 < 1; fi53++ {
                                if false {
                                    var v54 int = (^(fi53))
                                    fmt.Println(v1, v3)
                                    _ = v54
                                } else {
                                    _ = len(v2)
                                }
                                var v55 int32 = v1
                                v56 := GenS4{F0: 2147483647, F1: v4, F2: 315}
                                v4 = (v4 * v4)
                                var v57 int16 = 1
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi53
                                _ = v55
                                _ = v56
                                _ = v57
                            }
                        } else {
                            var v58 bool = true
                            _ = v58
                        }
                    case 421:
                        v50 = v50
                    }
                    var v59 uint64 = 4294967295
                    v50 = v50
                    _ = fi52
                    _ = v59
                }
            } else {
                v60 := GenS10{F0: 0, F1: "\n\t", F2: -2147483648}
                switch "a" {
                case "foo bar":
                    var v61 int16 = 32767
                    _ = v61
                default:
                    v5 = v5
                }
                _ = v60
            }
            if len(v2) > 0 {
                v2[0] = (v3 | (v3 % (((v3 / ((v3) | 1))) | 1)))
            }
            fmt.Println(v3, v4, v4)
        }
        _ = v50
    }
    if len(v2) > 0 {
        v2[0] = v3
    }
    var v62 uint = 255
    for fi63 := 0; fi63 < 8; fi63++ {
        switch (v3 << ((v62) & 63)) {
        case 0:
            for fi64 := 0; fi64 < 16; fi64++ {
                var v65 string = v5
                var v66 uint32 = 933
                _ = fi64
                _ = v65
                _ = v66
            }
        case 397:
            if true {
                fmt.Println(v3, v62, v3)
            }
        }
        v1 -= int32(v4)
        _ = fi63
    }
    var v67 int8 = 1
    _ = len(v2)
    for fi68 := 0; fi68 < 6; fi68++ {
        fmt.Println(v3)
        var v69 int64 = 1
        _ = fi68
        _ = v69
    }
    for fi70 := 0; fi70 < 5; fi70++ {
        v62 = v62
        for fi71 := 0; fi71 < 1; fi71++ {
            v4 = v4
            switch v3 {
            case 1:
                v67 += v67
            }
            _ = fi71
        }
        v2 = append(v2, int(v1))
        _ = fi70
    }
    v4 -= (v4 ^ v4)
    v4 += v4
    for fi72 := 0; fi72 < 1; fi72++ {
        if false {
            var v73 bool = false
            _ = v73
        }
        switch v5 {
        case "\n\t":
            if (false || false) {
                v4 += (^(v4))
            }
        case "foo bar":
            fmt.Println(v67, v5)
        default:
            v67 = v67
        }
        for fi74 := 0; fi74 < 3; fi74++ {
            fmt.Println(v4, v62, v67)
            _ = fi74
        }
        _ = fi72
    }
    fmt.Println(v4, v3)
    v2 = append(v2, (v3 % (((v3 << (((v62 | v62)) & 63))) | 1)))
    if false {
        if true {
            if len(v2) > 0 {
                v2[0] = (v3 >> ((18446744073709551615) & 63))
            }
            if true {
                var v75 uint8 = (v4 & (v4 - (v4 * v4)))
                if true {
                    v76 := GenS11{F0: "a"}
                    v4 = (v4 >> (((v62 << (((v62 << (((v62 & (v62 + v62))) & 63))) & 63))) & 63))
                    _ = v76
                } else {
                    var v77 uint = 255
                    _ = v77
                }
                _ = v75
            }
        }
        if false {
            v78 := GenS12{F0: (v1 & 0)}
            if (0 <= 32767) {
                var v79 int64 = -311
                var v80 bool = true
                v79 += (v79 << ((v62) & 63))
                _ = v79
                _ = v80
            }
            _ = v78
        } else {
            v81 := GenS11{F0: v5}
            for fi82 := 0; fi82 < 8; fi82++ {
                switch fi82 {
                case -2147483648:
                    v62 = v62
                case 32767:
                    v83 := GenS8{F0: "foo bar", F1: 751}
                    _ = v83
                default:
                    var v84 uint64 = 18446744073709551615
                    _ = v84
                }
                for fi85 := 0; fi85 < 14; fi85++ {
                    if true {
                        fmt.Println(v1, fi82, v4)
                    }
                    if (false || (1.5 > (3.14159 + float32(1.5)))) {
                        if (!true) {
                            v1 += (v1 / ((v1) | 1))
                            switch ((v5 + "") + v5) {
                            case "hello":
                                v86 := GenS9{F0: ("" + (v5 + (v5 + v5))), F1: 625}
                                _ = v86
                            case "foo bar":
                                var v87 float32 = 0.1
                                _ = v87
                            case "":
                                v88 := GenS6{F0: -22, F1: v5}
                                _ = v88
                            default:
                                fmt.Println(v3, fi82, v1)
                            }
                            var v89 int = (fi82 * (-(fi85)))
                            _ = v89
                        }
                        v62 += (v62 / ((v62) | 1))
                    }
                    if true {
                        fmt.Println(fi85, v4, v67)
                    }
                    fmt.Println(fi82)
                    _ = fi85
                }
                _ = fi82
            }
            v90 := GenS12{F0: v1}
            _ = v81
            _ = v90
        }
        var v91 float64 = 1e10
        _ = v91
    } else {
        for fi92 := 0; fi92 < 6; fi92++ {
            for fi93 := 0; fi93 < 17; fi93++ {
                var v94 int16 = -307
                var v95 int64 = 0
                _ = fi93
                _ = v94
                _ = v95
            }
            _ = fi92
        }
    }
    switch 127 {
    case -342:
        v67 += int8(v3)
    case 2147483647:
        if (!(42 < 445)) {
            fmt.Println(v3, v62, v67)
            if (true || true) {
                for fi96 := 0; fi96 < 19; fi96++ {
                    fmt.Println(v5, v62)
                    _ = fi96
                }
                v97 := GenS8{F0: "", F1: 694}
                for fi98 := 0; fi98 < 4; fi98++ {
                    for fi99 := 0; fi99 < 5; fi99++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi99
                    }
                    v5 = (v5 + (v5 + (v5 + v5)))
                    _ = fi98
                }
                _ = v97
            }
            if false {
                var v100 bool = true
                _ = v100
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v62:", v62)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v62
    _ = v67
}
