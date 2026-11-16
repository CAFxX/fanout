package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint16; F2 bool }
type GenS3 struct { F0 bool; F1 int; F2 uint32 }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 uint; F1 int8 }

func main() {
    var v1 int = -151
    var v2 float64 = 1e10
    var v3 uint16 = 255
    var v4 float32 = float32(0.0)
    v5 := []uint32{769, 2147483648}
    v6 := []uint32{}
    v2 = 3.14159
    v4 = -2.25
    var v7 int64 = -178
    v2 = v2
    var v8 int64 = -983
    v2 -= 0.1
    switch (v1 ^ -910) {
    case 32768:
        var v9 uint = 2
        _ = v9
    case 127:
        if (false || (true && true)) {
            var v10 uint16 = 255
            _ = v10
        } else {
            var v11 bool = true
            if ((v3 - (^(v3))) < 42) {
                fmt.Println(v7)
                switch (v1 >> ((2) & 63)) {
                case 567:
                    v12 := GenS1{F0: 62, F1: 524, F2: 497}
                    _ = v12
                }
            } else {
                v13 := GenS1{F0: -1, F1: -2147483648, F2: 0}
                _ = len(v5)
                _ = v13
            }
            fmt.Println(v8)
            _ = v11
        }
    default:
        if len(v6) > 0 {
            v6[0] = 214
        }
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v15 := GenS1{F0: -2147483648, F1: 490, F2: 852}
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        for fi17 := 0; fi17 < 6; fi17++ {
            if false {
                fmt.Println(v4)
            }
            switch "" {
            case "foo bar":
                if (true || (!false)) {
                    v2 = -2.25
                    switch "" {
                    case "a":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case "foo bar":
                        v4 -= 1e10
                    case "hello":
                        v18 := GenS1{F0: -1, F1: -2147483648, F2: 812}
                        _ = v18
                    default:
                        fmt.Println(fi17)
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                } else {
                    if true {
                        var v19 string = "foo bar"
                        _ = v19
                    } else {
                        if (("a" + ("foo bar" + "a")) == "\n\t") {
                            switch fi17 {
                            case 58:
                                v20 := GenS2{F0: 42, F1: 1, F2: (32767 != -681)}
                                _ = v20
                            default:
                                v1 -= 286
                            }
                        }
                    }
                    fmt.Println(v4, v3, fi16)
                    if (57 > 2) {
                        v4 -= (((v4 + 1e10) - 0.0) + v4)
                        var v21 string = ("\n\t" + "")
                        _ = v21
                    } else {
                        var v22 uint16 = v3
                        if (440 <= 876) {
                            var v23 uint8 = 255
                            _ = v23
                        }
                        var v24 bool = false
                        _ = v22
                        _ = v24
                    }
                    var v25 int32 = 1
                    var v26 float32 = (v4 - v4)
                    if false {
                        switch "a" {
                        case "":
                            switch (v1 ^ v1) {
                            case 32768:
                                var v27 uint16 = v3
                                _ = v27
                            case 890:
                                var v28 float32 = ((v26 + v4) - (v26 + v26))
                                _ = v28
                            case -9223372036854775808:
                                var v29 int64 = v7
                                _ = v29
                            default:
                                if (-32768 == 32767) {
                                    switch 1 {
                                    case 935:
                                        v30 := GenS3{F0: true, F1: (fi16 & (v1 - fi17)), F2: 2}
                                        _ = v30
                                    case 0:
                                        if ((v8 ^ (v8 % (((v8 * (v7 * (v8 * (v8 >> ((33) & 63)))))) | 1))) >= (v7 / ((38) | 1))) {
                                            var v31 bool = false
                                            _ = v31
                                        }
                                    case 128:
                                        switch "a" {
                                        case "a":
                                            v32 := GenS3{F0: false, F1: 128, F2: 291}
                                            _ = v32
                                        case "\n\t":
                                            v33 := GenS1{F0: v25, F1: v25, F2: 212}
                                            _ = v33
                                        default:
                                            v34 := GenS4{F0: 65535, F1: 42}
                                            _ = v34
                                        }
                                    }
                                } else {
                                    v8 += v7
                                }
                            }
                        case "\n\t":
                            var v35 int = -646
                            _ = v35
                        }
                    }
                    _ = v25
                    _ = v26
                }
            case "\n\t":
                var v36 uint8 = 127
                _ = v36
            case "":
                if len(v5) > 0 {
                    v5[0] = 763
                }
            }
            _ = fi17
        }
        _ = fi16
    }
    if true {
        if true {
            var v37 int32 = -2147483648
            if false {
                v37 = 665
                for fi38 := 0; fi38 < 6; fi38++ {
                    v2 = (v2 / v2)
                    _ = fi38
                }
                switch v1 {
                case 42:
                    for fi39 := 0; fi39 < 19; fi39++ {
                        fmt.Println(v2, v8, v7)
                        for fi40 := 0; fi40 < 11; fi40++ {
                            if len(v5) > 0 {
                                v5[0] = 197
                            }
                            if true {
                                var v41 uint32 = 2
                                var v42 uint = 42
                                var v43 int64 = v8
                                switch "foo bar" {
                                case "":
                                    v8 += v43
                                    if true {
                                        if len(v5) > 0 {
                                            v5[0] = v41
                                        }
                                        fmt.Println(v37)
                                        var v44 int8 = 127
                                        _ = v44
                                    } else {
                                        v7 = (v43 / ((-9223372036854775808) | 1))
                                        var v45 uint64 = 4294967295
                                        _ = v45
                                    }
                                case "foo bar":
                                    var v46 uint16 = (v3 / ((1) | 1))
                                    _ = v46
                                case "hello":
                                    if (!false) {
                                        if false {
                                            if false {
                                                if false {
                                                    var v47 string = (((("a" + "a") + ("a" + "foo bar")) + "\n\t") + "foo bar")
                                                    switch v47 {
                                                    case "hello":
                                                        var v48 uint32 = 4294967295
                                                        _ = v48
                                                    case "":
                                                        var v49 int = 9223372036854775807
                                                        _ = v49
                                                    }
                                                    var v50 int64 = (v43 << ((37) & 63))
                                                    _ = v47
                                                    _ = v50
                                                } else {
                                                    var v51 uint32 = v41
                                                    var v52 int32 = v37
                                                    fmt.Println(v41)
                                                    fmt.Println(v52)
                                                    _ = v51
                                                    _ = v52
                                                }
                                                if true {
                                                    v5 = append(v5, (v41 ^ (v41 << (((v42 % (((v42 + 18446744073709551615)) | 1))) & 63))))
                                                    v4 += -2.25
                                                } else {
                                                    v3 = v3
                                                    v53 := GenS3{F0: true, F1: (fi39 | -1), F2: (^(v41))}
                                                    v5 = append(v5, (v41 | (^(v41))))
                                                    _ = v53
                                                }
                                            }
                                            v6 = append(v6, 656)
                                        } else {
                                            fmt.Println(v37, v7, v37)
                                            switch v1 {
                                            case -498:
                                                if len(v6) > 0 {
                                                    v6[0] = v41
                                                }
                                            }
                                            switch v1 {
                                            case -78:
                                                v37 -= (v37 * (v37 & v37))
                                            default:
                                                var v54 int = (fi39 << ((v42) & 63))
                                                _ = v54
                                            }
                                        }
                                        v1 -= fi40
                                    } else {
                                        if (v2 != v2) {
                                            switch 700 {
                                            case -1:
                                                v55 := GenS2{F0: (v37 | v37), F1: 65535, F2: ((^(fi39)) > (v1 - 379))}
                                                var v56 uint = uint(v3)
                                                _ = v55
                                                _ = v56
                                            case 32767:
                                                v57 := GenS5{F0: v41, F1: ((float32(-2.25) / v4) == v4)}
                                                var v58 uint64 = 492
                                                var v59 int16 = 955
                                                _ = v57
                                                _ = v58
                                                _ = v59
                                            default:
                                                fmt.Println(v41, v1)
                                            }
                                        } else {
                                            var v60 uint16 = (v3 << ((255) & 63))
                                            v43 -= (v8 & -229)
                                            var v61 uint8 = 42
                                            _ = v60
                                            _ = v61
                                        }
                                    }
                                }
                                if false {
                                    var v62 int32 = 42
                                    _ = v62
                                } else {
                                    v63 := GenS1{F0: v37, F1: (v37 / (((v37 + (^(v37)))) | 1)), F2: 4294967295}
                                    v7 -= 9223372036854775807
                                    var v64 string = (("\n\t" + ("foo bar" + "\n\t")) + ("foo bar" + "foo bar"))
                                    _ = v63
                                    _ = v64
                                }
                                fmt.Println(fi39, v42)
                                v65 := GenS4{F0: v42, F1: 2}
                                _ = v41
                                _ = v42
                                _ = v43
                                _ = v65
                            } else {
                                if (127 == 1) {
                                    if len(v6) > 0 {
                                        v6[0] = 2147483647
                                    }
                                    v66 := GenS2{F0: -656, F1: 255, F2: false}
                                    _ = v66
                                } else {
                                    v3 = (v3 & (v3 >> ((0) & 63)))
                                    if true {
                                        if (true || false) {
                                            v7 += v7
                                            var v67 int16 = 599
                                            _ = v67
                                        } else {
                                            if (1.5 == float32(1.5)) {
                                                v68 := GenS4{F0: 828, F1: 42}
                                                var v69 float32 = v4
                                                _ = v68
                                                _ = v69
                                            }
                                            if false {
                                                v6 = append(v6, 2147483647)
                                                v70 := GenS6{F0: v3, F1: "hello"}
                                                _ = v70
                                            }
                                        }
                                    } else {
                                        var v71 uint64 = 910
                                        v72 := GenS2{F0: 42, F1: 655, F2: true}
                                        v73 := GenS5{F0: 4294967295, F1: true}
                                        _ = v71
                                        _ = v72
                                        _ = v73
                                    }
                                }
                            }
                            var v74 int = -9223372036854775808
                            _ = fi40
                            _ = v74
                        }
                        _ = fi39
                    }
                case -402:
                    for fi75 := 0; fi75 < 6; fi75++ {
                        v2 = (float64(v4) / v2)
                        _ = fi75
                    }
                default:
                    v76 := GenS3{F0: false, F1: (v1 | v1), F2: 4294967295}
                    _ = v76
                }
            }
            var v77 float64 = float64(0.0)
            _ = v37
            _ = v77
        } else {
            fmt.Println(v8, v4)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            for fi78 := 0; fi78 < 8; fi78++ {
                v3 += 1
                switch (fi78 ^ fi78) {
                case 9223372036854775807:
                    fmt.Println(v2)
                case -132:
                    var v79 uint64 = 18446744073709551615
                    _ = v79
                case 2147483647:
                    v2 += 0.0
                default:
                    var v80 uint = 255
                    _ = v80
                }
                var v81 int64 = (v8 - (v7 - v7))
                _ = fi78
                _ = v81
            }
        }
        if (true || true) {
            v7 += v8
            var v82 uint32 = 42
            v1 += v1
            _ = v82
        } else {
            v1 -= v1
            for fi83 := 0; fi83 < 17; fi83++ {
                _ = len(v6)
                _ = fi83
            }
        }
    } else {
        switch (v1 | (v1 | -968)) {
        case 42:
            for fi84 := 0; fi84 < 7; fi84++ {
                var v85 uint8 = 154
                _ = fi84
                _ = v85
            }
        default:
            var v86 string = "foo bar"
            _ = v86
        }
    }
    switch "a" {
    case "\n\t":
        var v87 int8 = -113
        _ = v87
    case "":
        if false {
            switch 445 {
            case -59:
                for fi88 := 0; fi88 < 5; fi88++ {
                    fmt.Println(v2)
                    _ = fi88
                }
            default:
                var v89 int32 = -2147483648
                _ = v89
            }
            for fi90 := 0; fi90 < 4; fi90++ {
                for fi91 := 0; fi91 < 8; fi91++ {
                    var v92 uint32 = 2147483647
                    _ = fi91
                    _ = v92
                }
                for fi93 := 0; fi93 < 5; fi93++ {
                    v4 = v4
                    _ = fi93
                }
                v94 := GenS7{F0: 18446744073709551615, F1: int8(v7)}
                _ = fi90
                _ = v94
            }
            var v95 float32 = (v4 - (0.0 * v4))
            _ = v95
        } else {
            for fi96 := 0; fi96 < 11; fi96++ {
                for fi97 := 0; fi97 < 18; fi97++ {
                    _ = len(v5)
                    _ = fi97
                }
                _ = fi96
            }
        }
    case "foo bar":
        if (!(9223372036854775807 > v8)) {
            for fi98 := 0; fi98 < 6; fi98++ {
                _ = len(v6)
                switch int(v7) {
                case 4294967295:
                    var v99 uint64 = 0
                    _ = v99
                }
                var v100 int16 = 138
                _ = fi98
                _ = v100
            }
        } else {
            var v101 string = "a"
            _ = v101
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
