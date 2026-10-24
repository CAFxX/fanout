package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64; F1 int }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint8; F1 uint8 }
type GenS12 struct { F0 int32; F1 uint64; F2 int16 }
type GenS13 struct { F0 uint64; F1 int }

func main() {
    v1 := []uint32{2147483648}
    v2 := make(map[uint32]string)
    var v3 string = ("" + "a")
    var v4 uint16 = 297
    v5 := GenS1{F0: 1, F1: 959}
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            switch fi7 {
            case 1:
                switch v3 {
                case "a":
                    v1 = append(v1, 0)
                }
                fmt.Println(v4, v3, v3)
            default:
                var v8 int32 = 0
                _ = v8
            }
            fmt.Println(fi7, v3, fi6)
            switch (fi6 ^ (fi7 / ((fi6) | 1))) {
            case 127:
                v9 := GenS2{F0: v3, F1: 18446744073709551615}
                switch (fi7 * fi6) {
                case -2147483648:
                    if (((v3 + v3) <= v3) && false) {
                        if false {
                            v10 := GenS1{F0: 1, F1: -90}
                            delete(v2, 344)
                            var v11 int8 = -98
                            _ = v10
                            _ = v11
                        }
                    }
                }
                _ = v9
            case -1:
                v1 = append(v1, 266)
            }
            _ = fi7
        }
        _ = fi6
    }
    v2[42] = (v3 + "a")
    v12 := GenS3{F0: 42}
    var v13 uint64 = 1
    v14 := GenS4{F0: 338}
    fmt.Println(v4, v3, v13)
    if (true && (((("foo bar" + (("hello" + v3) + v3)) != "\n\t") && (uint8(v13) < 255)) || false)) {
        for fi15 := 0; fi15 < 2; fi15++ {
            for fi16 := 0; fi16 < 18; fi16++ {
                switch int(v4) {
                case 32768:
                    v13 += 42
                case 1:
                    fmt.Println(fi16)
                case 4294967295:
                    _ = len(v1)
                default:
                    if (false || (v3 > v3)) {
                        if false {
                            v4 -= v4
                            v17 := GenS5{F0: 0}
                            switch v3 {
                            case "\n\t":
                                var v18 uint8 = 112
                                _ = v18
                            case "a":
                                v13 = (v13 | v13)
                            case "foo bar":
                                if false {
                                    _ = len(v2)
                                    var v19 uint16 = v4
                                    switch 64 {
                                    case 32768:
                                        v2[uint32(v13)] = v3
                                    case -9223372036854775808:
                                        switch v3 {
                                        case "hello":
                                            v20 := GenS4{F0: 0}
                                            _ = v20
                                        }
                                    }
                                    v13 = (v13 | (v13 - v13))
                                    if false {
                                        v13 += (v13 | v13)
                                        var v21 bool = (-2.25 < (3.14159 / 1e10))
                                        _ = v21
                                    }
                                    _ = v19
                                }
                            }
                            _ = v17
                        }
                    } else {
                        fmt.Println(v13, v13)
                        if ((0 < 348) || ((v4 > (v4 / (((v4 ^ (v4 % ((v4) | 1)))) | 1))) || true)) {
                            v13 -= 659
                            if len(v1) > 0 {
                                v1[0] = 2
                            }
                        } else {
                            var v22 uint = 4294967295
                            switch 0 {
                            case 32768:
                                fmt.Println(v22)
                            default:
                                fmt.Println(v22, fi15)
                            }
                            _ = v22
                        }
                        if (641 >= uint32(v4)) {
                            if true {
                                if (-9223372036854775808 == 0) {
                                    var v23 int32 = -436
                                    v1 = append(v1, 65535)
                                    var v24 float64 = 0.0
                                    _ = len(v2)
                                    v23 += v23
                                    _ = v23
                                    _ = v24
                                }
                                switch fi16 {
                                case -947:
                                    v4 = 709
                                case -1:
                                    if true {
                                        if ((!(!true)) || false) {
                                            v13 += (v13 & v13)
                                            fmt.Println(fi15, fi16, v3)
                                            v4 += v4
                                        } else {
                                            if false {
                                                fmt.Println(v4)
                                                var v25 int32 = 2147483647
                                                var v26 int32 = int32(v13)
                                                _ = v25
                                                _ = v26
                                            }
                                        }
                                        fmt.Println(v13, v3)
                                    } else {
                                        switch fi16 {
                                        case -281:
                                            var v27 int8 = 1
                                            v28 := GenS4{F0: 1}
                                            _ = v27
                                            _ = v28
                                        default:
                                            _ = len(v2)
                                            if true {
                                                fmt.Println(fi16, v4)
                                                switch (v3 + (v3 + "a")) {
                                                case "foo bar":
                                                    delete(v2, 2147483647)
                                                case "\n\t":
                                                    v29 := GenS4{F0: 4294967295}
                                                    _ = v29
                                                case "":
                                                    var v30 float32 = (3.14159 + 0.1)
                                                    _ = v30
                                                default:
                                                    delete(v2, 65535)
                                                }
                                                v31 := GenS5{F0: 2}
                                                _ = v31
                                            } else {
                                                v32 := GenS5{F0: 532}
                                                if false {
                                                    if false {
                                                        if true {
                                                            v3 = v3
                                                            fmt.Println(fi16)
                                                        } else {
                                                            var v33 int = (-(fi15))
                                                            var v34 bool = true
                                                            v34 = v34
                                                            _ = v33
                                                            _ = v34
                                                        }
                                                        v13 = v13
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    } else {
                                                        var v35 int8 = 100
                                                        v13 -= v13
                                                        switch v3 {
                                                        case "":
                                                            if (true || false) {
                                                                v13 -= v13
                                                            } else {
                                                                var v36 uint8 = 128
                                                                v37 := v2[803]
                                                                _ = v36
                                                                _ = v37
                                                            }
                                                        }
                                                        v35 += 110
                                                        fmt.Println(fi16)
                                                        if (0.0 == (-2.25 - -2.25)) {
                                                            v38 := GenS6{F0: 0}
                                                            if len(v1) > 0 {
                                                                v1[0] = 846
                                                            }
                                                            if false {
                                                                switch fi16 {
                                                                case 32767:
                                                                    if false {
                                                                        v39 := GenS2{F0: "a", F1: (v13 % (((v13 * v13)) | 1))}
                                                                        fmt.Println(v35)
                                                                        var v40 uint64 = v13
                                                                        _ = v39
                                                                        _ = v40
                                                                    }
                                                                case -2147483648:
                                                                    var v41 string = (v3 + v3)
                                                                    _ = v41
                                                                }
                                                            } else {
                                                                v35 -= 30
                                                                switch (((("foo bar" + (v3 + v3)) + ("foo bar" + ("foo bar" + (("a" + "") + ("foo bar" + ("" + v3)))))) + (v3 + ("" + (v3 + (v3 + (((v3 + "") + v3) + "")))))) + "\n\t") {
                                                                case "":
                                                                    v42 := GenS5{F0: 4294967295}
                                                                    _ = v42
                                                                }
                                                            }
                                                            _ = v38
                                                        } else {
                                                            var v43 bool = ((v3 + v3) >= (v3 + v3))
                                                            switch (fi16 | (fi15 % ((fi15) | 1))) {
                                                            case 74:
                                                                if (!v43) {
                                                                    if (v43 && (false || (!(-430 >= 42)))) {
                                                                        v44 := GenS7{F0: 0, F1: (fi16 - (fi15 & (fi15 * fi15)))}
                                                                        _ = v44
                                                                    } else {
                                                                        switch (fi16 * fi16) {
                                                                        case 32767:
                                                                            fmt.Println(fi16, v3)
                                                                        }
                                                                        var v45 int32 = 0
                                                                        v35 = v35
                                                                        _ = v45
                                                                    }
                                                                } else {
                                                                    fmt.Println(v43, fi16)
                                                                    var v46 uint32 = 136
                                                                    _ = v46
                                                                }
                                                            }
                                                            _ = v43
                                                        }
                                                        _ = v35
                                                    }
                                                    var v47 int32 = 269
                                                    v2[42] = ("a" + ("" + "foo bar"))
                                                    v3 = ""
                                                    _ = v47
                                                }
                                                switch ((v3 + ("\n\t" + v3)) + v3) {
                                                case "":
                                                    v3 = (v3 + (((((v3 + ((v3 + v3) + "foo bar")) + v3) + "") + "hello") + v3))
                                                case "hello":
                                                    v48 := GenS8{F0: (0.0 > (3.14159 / ((((-2.25 * 1e10) * (0.0 - (3.14159 * 0.1))) - 3.14159) - -2.25)))}
                                                    _ = v48
                                                case "foo bar":
                                                    v3 = ""
                                                }
                                                var v49 uint = 874
                                                _ = v32
                                                _ = v49
                                            }
                                            switch v3 {
                                            case "\n\t":
                                                switch "foo bar" {
                                                case "hello":
                                                    var v50 int32 = -2147483648
                                                    _ = v50
                                                case "a":
                                                    var v51 uint64 = v13
                                                    _ = v51
                                                case "\n\t":
                                                    v3 = ("a" + (v3 + ("a" + v3)))
                                                }
                                            case "hello":
                                                v52 := GenS9{F0: (0 >= 709)}
                                                _ = v52
                                            case "foo bar":
                                                v4 += 42
                                            }
                                            v53 := GenS8{F0: (false || (false || true))}
                                            _ = v53
                                        }
                                    }
                                    v54 := GenS10{F0: "\n\t"}
                                    v3 = (v3 + "\n\t")
                                    v4 = 524
                                    v13 = v13
                                    _ = v54
                                }
                                fmt.Println(fi16)
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            var v55 uint = 2
                            if false {
                                var v56 bool = ("a" >= (v3 + ("\n\t" + v3)))
                                var v57 uint8 = 179
                                v57 += 127
                                _ = v56
                                _ = v57
                            } else {
                                switch -1 {
                                case 0:
                                    var v58 uint = (^(v55))
                                    _ = v58
                                case 575:
                                    v3 = v3
                                case 32768:
                                    _ = len(v1)
                                }
                            }
                            _ = v55
                        } else {
                            v59 := GenS4{F0: 4294967295}
                            v60 := v2[65535]
                            _ = v59
                            _ = v60
                        }
                        var v61 int32 = -2147483648
                        _ = v61
                    }
                }
                v62 := v2[727]
                _ = fi16
                _ = v62
            }
            var v63 float32 = float32(0.1)
            v63 += v63
            _ = fi15
            _ = v63
        }
        var v64 uint64 = (v13 % ((4294967295) | 1))
        var v65 uint8 = 127
        _ = v64
        _ = v65
    } else {
        var v66 int = 4294967295
        var v67 uint32 = 792
        if (((((1.5 / float64(1e10)) * 1e10) / float64((((0.0 * ((0.1 * -2.25) * (1e10 - ((3.14159 * (0.0 * 3.14159)) + 0.1)))) + 1e10) - 0.1))) / 1e10) != 1.5) {
            if (-128 > 100) {
                var v68 uint32 = (v67 + v67)
                _ = v68
            }
            if (!true) {
                fmt.Println(v3, v66, v66)
                v69 := GenS11{F0: 128, F1: 1}
                v4 -= (v4 | v4)
                _ = v69
            }
            fmt.Println(v4, v66)
        } else {
            for fi70 := 0; fi70 < 19; fi70++ {
                if (v3 <= ("a" + v3)) {
                    var v71 uint16 = (v4 << ((255) & 63))
                    _ = v71
                }
                if false {
                    if false {
                        var v72 float32 = ((3.14159 * ((0.1 * 0.0) / (0.1 - 1.5))) / 3.14159)
                        _ = v72
                    } else {
                        var v73 int16 = 42
                        _ = v73
                    }
                }
                _ = fi70
            }
            fmt.Println(v67)
        }
        _ = v66
        _ = v67
    }
    v74 := GenS10{F0: v3}
    if true {
        v13 = (v13 * 192)
        fmt.Println(v4)
        for fi75 := 0; fi75 < 3; fi75++ {
            v3 = v3
            v76 := GenS12{F0: 42, F1: (v13 + 994), F2: -828}
            _ = fi75
            _ = v76
        }
    } else {
        if false {
            v77 := GenS6{F0: -1}
            v13 -= v13
            _ = v77
        }
    }
    var v78 string = (v3 + (v3 + v3))
    var v79 float64 = 0.0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v80 uint8 = 128
    if (true || (32767 > 4294967295)) {
        var v81 float32 = 0.0
        v79 += (1e10 / (((1e10 / 0.1) - (v79 * 0.0)) + v79))
        _ = v81
    } else {
        if (1 != uint(v4)) {
            for fi82 := 0; fi82 < 12; fi82++ {
                v3 = (v78 + v78)
                if true {
                    switch ((v3 + v78) + "a") {
                    case "foo bar":
                        var v83 int32 = 2147483647
                        _ = v83
                    case "":
                        switch (^(fi82)) {
                        case 461:
                            v84 := GenS13{F0: v13, F1: (fi82 >> ((1) & 63))}
                            _ = v84
                        case -2147483648:
                            switch (^(fi82)) {
                            case 127:
                                for fi85 := 0; fi85 < 16; fi85++ {
                                    v3 = ("a" + v3)
                                    v4 = (v4 | (v4 - v4))
                                    if true {
                                        switch "" {
                                        case "a":
                                            if false {
                                                fmt.Println(v78, fi82)
                                                v78 = "foo bar"
                                            } else {
                                                var v86 int = fi82
                                                _ = v86
                                            }
                                        case "foo bar":
                                            v80 -= (^(v80))
                                            v80 -= (^(v80))
                                            v87 := GenS4{F0: uint32(v80)}
                                            _ = v87
                                        default:
                                            fmt.Println(v80, v79)
                                        }
                                        v88 := GenS12{F0: 620, F1: v13, F2: 197}
                                        v13 = (^(v13))
                                        _ = v88
                                    } else {
                                        var v89 int32 = 1
                                        _ = len(v1)
                                        _ = v89
                                    }
                                    _ = fi85
                                }
                            case 42:
                                var v90 uint16 = v4
                                _ = v90
                            default:
                                v91 := GenS13{F0: v13, F1: 32768}
                                _ = v91
                            }
                        default:
                            delete(v2, 2)
                        }
                    }
                } else {
                    var v92 int16 = -265
                    _ = v92
                }
                fmt.Println(v3, v78)
                _ = fi82
            }
            fmt.Println(v13)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v74
    _ = v78
    _ = v79
    _ = v80
}
