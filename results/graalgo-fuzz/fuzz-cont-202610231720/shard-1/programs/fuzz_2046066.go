package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 uint64; F2 int }
type GenS3 struct { F0 string; F1 int32; F2 int }
type GenS4 struct { F0 uint8; F1 int64; F2 int8 }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int8; F1 uint32; F2 int32 }
type GenS8 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS9 struct { F0 uint16 }
func genFn0(p0 int, p1 int8, p2 uint) uint8 {
    p2 += (p2 << (((^(p2))) & 63))
    switch p0 {
    case 2147483647:
        if (!(2 >= 710)) {
            for fi6 := 0; fi6 < 2; fi6++ {
                if false {
                    v7 := []int{(fi6 / (((fi6 | (fi6 / ((-1) | 1)))) | 1))}
                    p2 = 65535
                    _ = v7
                }
                for fi8 := 0; fi8 < 8; fi8++ {
                    var v9 uint = (p2 << ((255) & 63))
                    var v10 uint16 = 65535
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                p1 = p1
                _ = fi6
            }
            fmt.Println(p2, p1)
        }
    }
    v11 := GenS1{F0: p0}
    v12 := GenS2{F0: (("foo bar" + "\n\t") + (("foo bar" + "hello") + "a")), F1: 42, F2: 2147483647}
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 229
}
func genFn1(p0 int16, p1 uint32, p2 int8) int64 {
    var v13 uint8 = 255
    var v14 int8 = p2
    var v15 float64 = 0.0
    v16 := GenS2{F0: ("" + "a"), F1: 2, F2: 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    return 42
}

func main() {
    var v1 int8 = 42
    var v2 int64 = 0
    var v3 bool = ((false || (1 == 1)) || true)
    var v4 uint8 = 131
    v5 := []uint32{2147483648, 2147483648}
    var v17 int16 = 0
    switch "foo bar" {
    case "\n\t":
        var v18 float32 = 3.14159
        _ = v18
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    _ = len(v5)
    var v19 string = ("hello" + (("a" + "a") + "\n\t"))
    v17 += (^(v17))
    if ((255 == 1) && true) {
        v20 := GenS1{F0: 2147483647}
        v2 = (v2 * v2)
        v2 += v2
        _ = v20
    }
    v21 := GenS3{F0: "\n\t", F1: 0, F2: -1}
    var v22 uint64 = 902
    for fi23 := 0; fi23 < 6; fi23++ {
        v3 = v3
        v22 = (v22 | v22)
        _ = fi23
    }
    if (v3 || (!true)) {
        var v24 uint = 18446744073709551615
        _ = v24
    } else {
        v25 := GenS3{F0: ("\n\t" + v19), F1: 2147483647, F2: 415}
        if (2147483648 < 2) {
            v3 = (v3 || true)
            var v26 int16 = (^(v17))
            _ = len(v5)
            _ = v26
        } else {
            v17 += v17
            v4 += (v4 / ((genFn0(32768, 96, 617)) | 1))
        }
        v4 = 225
        _ = v25
    }
    var v27 uint64 = v22
    if v3 {
        switch -891 {
        case -9223372036854775808:
            var v28 bool = v3
            _ = v28
        default:
            fmt.Println(v22, v22)
        }
        if v3 {
            var v29 uint16 = 65535
            v30 := GenS1{F0: 127}
            for fi31 := 0; fi31 < 6; fi31++ {
                var v32 int64 = -202
                if v3 {
                    for fi33 := 0; fi33 < 16; fi33++ {
                        var v34 uint16 = 32767
                        var v35 uint16 = 0
                        v3 = (v19 > ("\n\t" + v19))
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                    v17 = 1000
                    var v36 int16 = 42
                    _ = v36
                }
                switch (fi31 << ((659) & 63)) {
                case 127:
                    v4 = (v4 * (v4 | (v4 << ((2) & 63))))
                }
                _ = fi31
                _ = v32
            }
            _ = v29
            _ = v30
        } else {
            v37 := GenS4{F0: (v4 << ((255) & 63)), F1: (v2 & v2), F2: v1}
            _ = v37
        }
        v38 := GenS4{F0: (v4 & (v4 - v4)), F1: 0, F2: -128}
        _ = v38
    } else {
        if (2147483647 < -2) {
            v2 -= 1
            var v39 uint64 = v22
            _ = v39
        } else {
            if v3 {
                if len(v5) > 0 {
                    v5[0] = 65535
                }
                fmt.Println(v17, v22, v19)
            } else {
                v3 = v3
                if (v3 || v3) {
                    v5 = append(v5, 2147483647)
                    fmt.Println(v17)
                    if len(v5) > 0 {
                        v5[0] = 2
                    }
                }
                var v40 uint64 = (v27 >> ((2) & 63))
                _ = v40
            }
        }
    }
    if ((!v3) && (!(534 <= 65535))) {
        for fi41 := 0; fi41 < 20; fi41++ {
            var v42 uint32 = 2147483648
            for fi43 := 0; fi43 < 14; fi43++ {
                v44 := GenS5{F0: 2147483647, F1: false}
                v45 := GenS3{F0: "", F1: -2147483648, F2: (fi43 | fi41)}
                var v46 int16 = 1000
                v5 = append(v5, v42)
                v5 = append(v5, 627)
                v47 := GenS6{F0: (v42 >> ((65535) & 63))}
                _ = fi43
                _ = v44
                _ = v45
                _ = v46
                _ = v47
            }
            var v48 int8 = 127
            _ = fi41
            _ = v42
            _ = v48
        }
        var v49 uint64 = v27
        if (!v3) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            fmt.Println(v3, v17)
            var v50 int16 = int16(v4)
            _ = v50
        }
        _ = v49
    } else {
        v51 := GenS1{F0: int(v17)}
        _ = v51
    }
    v22 += 2
    if (4294967295 >= 2) {
        switch v19 {
        case "hello":
            v22 = v27
        case "foo bar":
            if ((1.5 * 1e10) <= (3.14159 / 0.1)) {
                if ((1.5 / 1e10) < (0.0 + (float64(1e10) / 0.1))) {
                    for fi52 := 0; fi52 < 8; fi52++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi52
                    }
                    fmt.Println(v19, v4, v27)
                }
                if (!v3) {
                    for fi53 := 0; fi53 < 14; fi53++ {
                        if len(v5) > 0 {
                            v5[0] = 962
                        }
                        _ = fi53
                    }
                } else {
                    fmt.Println(v19, v4, v27)
                    if (462 > 181) {
                        var v54 uint8 = v4
                        _ = v54
                    } else {
                        v5 = append(v5, 49)
                    }
                }
                v55 := GenS7{F0: 127, F1: 324, F2: 728}
                _ = v55
            }
        case "":
            v27 = v27
        default:
            if false {
                if v3 {
                    var v56 uint8 = (v4 / (((v4 * (v4 + v4))) | 1))
                    if v3 {
                        v57 := GenS6{F0: 1}
                        v58 := GenS3{F0: (v19 + v19), F1: 2147483647, F2: 1}
                        _ = v57
                        _ = v58
                    } else {
                        for fi59 := 0; fi59 < 18; fi59++ {
                            v60 := GenS6{F0: 42}
                            _ = fi59
                            _ = v60
                        }
                        v61 := GenS8{F0: v4, F1: (v27 / ((v27) | 1)), F2: v2}
                        _ = v61
                    }
                    v62 := GenS7{F0: 77, F1: 65535, F2: -635}
                    _ = v56
                    _ = v62
                }
                v22 = v22
            } else {
                switch 326 {
                case -602:
                    v4 = v4
                case 574:
                    if v3 {
                        if v3 {
                            for fi63 := 0; fi63 < 19; fi63++ {
                                v5 = append(v5, 42)
                                _ = fi63
                            }
                        }
                        var v64 uint16 = 255
                        switch 9223372036854775807 {
                        case -493:
                            var v65 int32 = -1
                            _ = v65
                        case -1:
                            var v66 uint8 = (^(v4))
                            _ = v66
                        case 2147483647:
                            for fi67 := 0; fi67 < 4; fi67++ {
                                switch fi67 {
                                case 0:
                                    fmt.Println(v64, v17)
                                }
                                v22 = 2
                                _ = fi67
                            }
                        default:
                            switch 1 {
                            case 4294967295:
                                v2 += int64(v27)
                            }
                        }
                        _ = v64
                    } else {
                        for fi68 := 0; fi68 < 2; fi68++ {
                            for fi69 := 0; fi69 < 7; fi69++ {
                                fmt.Println(v17, v19)
                                if v3 {
                                    var v70 bool = v3
                                    _ = v70
                                } else {
                                    v71 := GenS7{F0: v1, F1: 4294967295, F2: 2147483647}
                                    var v72 float64 = -2.25
                                    _ = v71
                                    _ = v72
                                }
                                _ = fi69
                            }
                            v1 = v1
                            _ = fi68
                        }
                        switch "\n\t" {
                        case "":
                            var v73 string = v19
                            _ = v73
                        }
                        fmt.Println(v17)
                    }
                case 1:
                    if ((("hello" < v19) || v3) || false) {
                        for fi74 := 0; fi74 < 5; fi74++ {
                            v19 = ""
                            switch 32768 {
                            case -9223372036854775808:
                                if true {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                } else {
                                    switch 2147483647 {
                                    case -2147483648:
                                        switch fi74 {
                                        case 330:
                                            v22 -= (^(v27))
                                        case -105:
                                            var v75 int16 = (v17 ^ (^(v17)))
                                            _ = v75
                                        }
                                    case 331:
                                        v2 = (v2 >> ((566) & 63))
                                    case -847:
                                        fmt.Println(v22, v1, v22)
                                    default:
                                        if true {
                                            var v76 int16 = (v17 % (((v17 ^ v17)) | 1))
                                            _ = v76
                                        } else {
                                            v77 := GenS9{F0: 343}
                                            for fi78 := 0; fi78 < 16; fi78++ {
                                                if v3 {
                                                    fmt.Println(fi74, v27, v27)
                                                    if v3 {
                                                        v79 := GenS1{F0: fi78}
                                                        var v80 uint32 = 2147483648
                                                        _ = v79
                                                        _ = v80
                                                    }
                                                }
                                                if len(v5) > 0 {
                                                    v5[0] = 153
                                                }
                                                v19 = v19
                                                _ = fi78
                                            }
                                            _ = v77
                                        }
                                    }
                                }
                            case 186:
                                if (-2.25 < ((float64(1.5) * -2.25) * 1.5)) {
                                    switch (v19 + ("a" + (v19 + v19))) {
                                    case "":
                                        for fi81 := 0; fi81 < 2; fi81++ {
                                            var v82 string = ""
                                            if len(v5) > 0 {
                                                v5[0] = 0
                                            }
                                            switch fi74 {
                                            case -116:
                                                var v83 int16 = v17
                                                _ = v83
                                            case -1:
                                                var v84 string = v82
                                                _ = v84
                                            default:
                                                var v85 float32 = float32(1.5)
                                                _ = v85
                                            }
                                            _ = fi81
                                            _ = v82
                                        }
                                    case "a":
                                        if v3 {
                                            v86 := GenS2{F0: (v19 + "a"), F1: 1, F2: int(v4)}
                                            _ = v86
                                        } else {
                                            fmt.Println(v3, v17)
                                            v22 -= 0
                                        }
                                    default:
                                        for fi87 := 0; fi87 < 2; fi87++ {
                                            var v88 uint8 = (v4 % (((v4 * 255)) | 1))
                                            v1 -= 1
                                            _ = fi87
                                            _ = v88
                                        }
                                    }
                                    switch fi74 {
                                    case 865:
                                        fmt.Println(v19, fi74)
                                    }
                                    v89 := GenS2{F0: ("foo bar" + "a"), F1: v27, F2: (fi74 + -1)}
                                    _ = v89
                                } else {
                                    switch (((v19 + v19) + ((v19 + "\n\t") + ("hello" + ("foo bar" + v19)))) + (v19 + "\n\t")) {
                                    case "a":
                                        var v90 int64 = -1
                                        _ = v90
                                    default:
                                        for fi91 := 0; fi91 < 11; fi91++ {
                                            var v92 float64 = ((-2.25 / 1e10) / 0.1)
                                            _ = fi91
                                            _ = v92
                                        }
                                    }
                                }
                            case 4294967295:
                                var v93 float64 = float64(1e10)
                                _ = v93
                            default:
                                var v94 int32 = 246
                                _ = v94
                            }
                            _ = fi74
                        }
                        if v3 {
                            var v95 uint64 = (v22 + 472)
                            _ = v95
                        } else {
                            v27 -= v27
                        }
                        v2 += (v2 ^ (v2 << ((275) & 63)))
                    }
                default:
                    for fi96 := 0; fi96 < 2; fi96++ {
                        fmt.Println(v19, v22, v2)
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v97 := GenS6{F0: 490}
                        _ = fi96
                        _ = v97
                    }
                }
            }
        }
        for fi98 := 0; fi98 < 11; fi98++ {
            fmt.Println(v2, v19, v4)
            _ = fi98
        }
        v17 += 42
    }
    for fi99 := 0; fi99 < 14; fi99++ {
        var v100 int32 = -1
        v5 = append(v5, 4294967295)
        if len(v5) > 0 {
            v5[0] = uint32(v27)
        }
        _ = fi99
        _ = v100
    }
    switch (v19 + (("" + "hello") + v19)) {
    case "hello":
        if (32768 != 128) {
            var v101 int32 = -2147483648
            if (42 <= v4) {
                var v102 uint8 = (^(v4))
                _ = v102
            }
            if v3 {
                v4 += (v4 >> ((4294967295) & 63))
            }
            _ = v101
        } else {
            _ = len(v5)
            if (18446744073709551615 == 4294967295) {
                if len(v5) > 0 {
                    v5[0] = 2147483648
                }
                switch 269 {
                case 127:
                    _ = len(v5)
                case 128:
                    v19 = v19
                default:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            } else {
                for fi103 := 0; fi103 < 5; fi103++ {
                    v22 += 1
                    _ = fi103
                }
            }
            var v104 bool = (((v4 > v4) || (!false)) || (1000 == v17))
            _ = v104
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
