package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 uint; F1 uint64; F2 int32 }
type GenS5 struct { F0 uint; F1 uint8 }
type GenS6 struct { F0 uint32; F1 string; F2 uint16 }
type GenS7 struct { F0 uint64; F1 int64 }
type GenS8 struct { F0 int8; F1 string; F2 int32 }
type GenS9 struct { F0 int }

func main() {
    v1 := []int{127, -2147483648, 1, -451}
    var v2 int8 = -31
    v3 := []string{"", "\n\t", ((("" + "\n\t") + "hello") + ""), ""}
    if false {
        v2 += 58
    } else {
        var v4 int = 315
        _ = v4
    }
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 19; fi6++ {
            fmt.Println(fi6, fi6, fi6)
            v2 += v2
            _ = fi6
        }
        fmt.Println(v2)
        v7 := GenS1{F0: 42, F1: "", F2: 32767}
        _ = fi5
        _ = v7
    }
    v3 = append(v3, "a")
    v2 += 127
    v8 := GenS1{F0: -2147483648, F1: (("hello" + (("a" + "\n\t") + "")) + ""), F2: 65535}
    var v9 uint16 = 255
    for fi10 := 0; fi10 < 18; fi10++ {
        for fi11 := 0; fi11 < 7; fi11++ {
            if true {
                v2 = -1
                if true {
                    v12 := GenS1{F0: 1, F1: "a", F2: 664}
                    var v13 int64 = 9223372036854775807
                    if false {
                        var v14 int8 = v2
                        if false {
                            v14 -= (v14 - (v14 >> ((566) & 63)))
                            var v15 uint16 = 42
                            _ = v15
                        }
                        switch (fi10 | fi11) {
                        case 903:
                            fmt.Println(fi10, v13, fi10)
                        case 2147483647:
                            v14 = 0
                        default:
                            if ((!false) || (57 != 2)) {
                                if false {
                                    v16 := GenS1{F0: 0, F1: "foo bar", F2: v9}
                                    v2 -= v2
                                    _ = v16
                                } else {
                                    v14 = (^(v14))
                                    var v17 uint32 = 38
                                    v13 = 42
                                    var v18 uint16 = v9
                                    if ((fi11 * 2147483647) < (fi10 % ((fi10) | 1))) {
                                        if true {
                                            var v19 float64 = 1e10
                                            v20 := GenS2{F0: 2, F1: v9}
                                            switch "a" {
                                            case "hello":
                                                var v21 uint8 = 0
                                                _ = v21
                                            }
                                            _ = v19
                                            _ = v20
                                        }
                                        if false {
                                            v3 = append(v3, ("" + ""))
                                            var v22 float32 = 3.14159
                                            _ = v22
                                        } else {
                                            if true {
                                                if false {
                                                    if false {
                                                        if len(v3) > 0 {
                                                            v3[0] = "a"
                                                        }
                                                    } else {
                                                        switch ("foo bar" + "\n\t") {
                                                        case "hello":
                                                            v14 = v14
                                                            if true {
                                                                fmt.Println(v13, v9, fi11)
                                                                if (1e10 >= 0.1) {
                                                                    if (true && true) {
                                                                        var v23 uint8 = 1
                                                                        v24 := GenS1{F0: 42, F1: "hello", F2: 870}
                                                                        if false {
                                                                            switch "foo bar" {
                                                                            case "a":
                                                                                var v25 uint64 = 0
                                                                                _ = v25
                                                                            case "":
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = (fi11 % ((int(v13)) | 1))
                                                                                }
                                                                            case "foo bar":
                                                                                v13 -= (^(v13))
                                                                                v2 -= v2
                                                                            }
                                                                            var v26 uint8 = 0
                                                                            v26 = 0
                                                                            _ = v26
                                                                        } else {
                                                                            if (true || ((!(-2.25 == ((1.5 + (0.1 / (((-2.25 * 0.1) + -2.25) - -2.25))) + 0.1))) || true)) {
                                                                                v13 -= v13
                                                                                var v27 int8 = v14
                                                                                if (true || false) {
                                                                                    var v28 int32 = int32(fi11)
                                                                                    _ = v28
                                                                                }
                                                                                _ = v27
                                                                            } else {
                                                                                fmt.Println(fi10)
                                                                                v18 += v9
                                                                                var v29 int8 = (v2 * (v2 | (v2 & (v2 - 1))))
                                                                                if (false && ((v13 + (v13 / (((v13 >> ((4294967295) & 63))) | 1))) != v13)) {
                                                                                    fmt.Println(v29, v13, v9)
                                                                                    var v30 int16 = 68
                                                                                    v14 += 0
                                                                                    if (-2.25 < (((float64((-2.25 / 1e10)) / (1.5 - -2.25)) + float64(-2.25)) + (3.14159 / 0.1))) {
                                                                                        switch "a" {
                                                                                        case "":
                                                                                            if (uint32(v23) == (^(v17))) {
                                                                                                v2 += int8(v23)
                                                                                                var v31 int16 = (v30 & -1)
                                                                                                var v32 float64 = -2.25
                                                                                                _ = v31
                                                                                                _ = v32
                                                                                            }
                                                                                            v18 = 42
                                                                                            v33 := GenS1{F0: -1, F1: ("" + ""), F2: v9}
                                                                                            _ = v33
                                                                                        case "hello":
                                                                                            v2 = (^(v29))
                                                                                        case "a":
                                                                                            var v34 bool = false
                                                                                            _ = v34
                                                                                        default:
                                                                                            if false {
                                                                                                var v35 string = "hello"
                                                                                                _ = v35
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    _ = v30
                                                                                } else {
                                                                                    var v36 int8 = (v14 + v14)
                                                                                    var v37 float64 = 0.1
                                                                                    var v38 int64 = v13
                                                                                    _ = v36
                                                                                    _ = v37
                                                                                    _ = v38
                                                                                }
                                                                                var v39 int16 = -32768
                                                                                v40 := GenS1{F0: 42, F1: "a", F2: v18}
                                                                                if (!true) {
                                                                                    switch fi10 {
                                                                                    case -29:
                                                                                        switch "hello" {
                                                                                        case "hello":
                                                                                            var v41 int64 = v13
                                                                                            _ = v41
                                                                                        default:
                                                                                            v29 -= -128
                                                                                        }
                                                                                    case -9223372036854775808:
                                                                                        var v42 uint64 = 38
                                                                                        _ = v42
                                                                                    default:
                                                                                        if (0.0 >= 1.5) {
                                                                                            if (492 >= (-(fi10))) {
                                                                                                if (false || true) {
                                                                                                    var v43 bool = (1 < (v39 % ((v39) | 1)))
                                                                                                    var v44 int32 = 1
                                                                                                    var v45 uint32 = v17
                                                                                                    v44 = (v44 * -2147483648)
                                                                                                    _ = v43
                                                                                                    _ = v44
                                                                                                    _ = v45
                                                                                                }
                                                                                            }
                                                                                            var v46 uint8 = uint8(v18)
                                                                                            v29 = v2
                                                                                            _ = v46
                                                                                        } else {
                                                                                            switch -628 {
                                                                                            case -1:
                                                                                                var v47 bool = (1.5 >= (1e10 / 0.1))
                                                                                                _ = v47
                                                                                            case -303:
                                                                                                if len(v3) >= 2 {
                                                                                                    v3 = v3[1:len(v3)]
                                                                                                }
                                                                                            case 128:
                                                                                                _ = len(v1)
                                                                                            }
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = (fi10 + (fi11 + (^(fi10))))
                                                                                            }
                                                                                            v2 = (^(v14))
                                                                                        }
                                                                                    }
                                                                                    if false {
                                                                                        var v48 uint16 = 32767
                                                                                        v14 += v29
                                                                                        _ = v48
                                                                                    } else {
                                                                                        var v49 uint64 = 0
                                                                                        v2 -= v14
                                                                                        _ = v49
                                                                                    }
                                                                                    v29 += int8(fi11)
                                                                                } else {
                                                                                    switch int(v14) {
                                                                                    case 42:
                                                                                        var v50 int16 = v39
                                                                                        _ = v50
                                                                                    case 2147483647:
                                                                                        v14 += int8(v23)
                                                                                        v14 -= (v29 ^ (v2 >> ((18446744073709551615) & 63)))
                                                                                        switch fi11 {
                                                                                        case -680:
                                                                                            if false {
                                                                                                switch (fi11 & fi10) {
                                                                                                case 2147483647:
                                                                                                    v51 := GenS3{F0: 2, F1: (v18 % (((v18 / (((v18 ^ 255)) | 1))) | 1))}
                                                                                                    _ = v51
                                                                                                }
                                                                                                if false {
                                                                                                    v52 := GenS3{F0: 1, F1: 2}
                                                                                                    v39 = 32767
                                                                                                    _ = v52
                                                                                                } else {
                                                                                                    v53 := GenS4{F0: 864, F1: 25, F2: -373}
                                                                                                    fmt.Println(v29)
                                                                                                    fmt.Println(v29, v17, fi11)
                                                                                                    _ = v53
                                                                                                }
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                            }
                                                                                        case -584:
                                                                                            if false {
                                                                                                fmt.Println(v17, fi10)
                                                                                            }
                                                                                        case -2147483648:
                                                                                            switch fi11 {
                                                                                            case -2147483648:
                                                                                                var v54 uint32 = v17
                                                                                                _ = v54
                                                                                            case 427:
                                                                                                v1 = append(v1, fi11)
                                                                                            default:
                                                                                                fmt.Println(fi10)
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(v13, v14, v29)
                                                                                    case 0:
                                                                                        var v55 int32 = 0
                                                                                        v17 = v17
                                                                                        if true {
                                                                                            v9 += (v9 << ((0) & 63))
                                                                                            var v56 string = (("\n\t" + ("" + "a")) + "a")
                                                                                            v14 += (^(v2))
                                                                                            _ = v56
                                                                                        } else {
                                                                                            var v57 uint64 = 979
                                                                                            fmt.Println(v17, v2)
                                                                                            _ = v57
                                                                                        }
                                                                                        _ = v55
                                                                                    default:
                                                                                        v2 = (-(v29))
                                                                                    }
                                                                                    _ = len(v1)
                                                                                    if true {
                                                                                        var v58 int64 = (v13 >> ((196) & 63))
                                                                                        _ = v58
                                                                                    }
                                                                                }
                                                                                _ = v29
                                                                                _ = v39
                                                                                _ = v40
                                                                            }
                                                                            var v59 uint16 = 2
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                            _ = v59
                                                                        }
                                                                        _ = v23
                                                                        _ = v24
                                                                    } else {
                                                                        v17 = v17
                                                                    }
                                                                    _ = len(v3)
                                                                    switch (fi10 ^ (fi10 % (((fi10 % (((fi10 % (((^(fi11))) | 1))) | 1))) | 1))) {
                                                                    case -6:
                                                                        v2 += (v2 - v14)
                                                                    default:
                                                                        var v60 uint16 = 272
                                                                        _ = v60
                                                                    }
                                                                }
                                                                fmt.Println(v9, v2)
                                                                var v61 uint64 = 609
                                                                _ = v61
                                                            }
                                                        default:
                                                            fmt.Println(v14)
                                                        }
                                                        v62 := GenS5{F0: 18446744073709551615, F1: 1}
                                                        _ = v62
                                                    }
                                                }
                                            } else {
                                                var v63 uint16 = (^(v18))
                                                _ = v63
                                            }
                                        }
                                    } else {
                                        fmt.Println(v17, fi11)
                                        v64 := GenS1{F0: 1, F1: ((("hello" + "foo bar") + (((("" + "\n\t") + "a") + ("" + ("foo bar" + ""))) + "hello")) + "hello"), F2: (v18 % (((v9 % (((v9 >> ((2) & 63))) | 1))) | 1))}
                                        v65 := GenS2{F0: 0, F1: v9}
                                        _ = v64
                                        _ = v65
                                    }
                                    var v66 int16 = 32767
                                    v18 = v18
                                    _ = v17
                                    _ = v18
                                    _ = v66
                                }
                                if (false || (!(!((false && true) || true)))) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    fmt.Println(v2, v9, fi10)
                                    var v67 string = ("hello" + "hello")
                                    v14 += -28
                                    _ = v67
                                }
                                var v68 string = ((("foo bar" + "") + "foo bar") + "a")
                                _ = len(v1)
                                _ = v68
                            } else {
                                v14 += v14
                                fmt.Println(fi10)
                            }
                        }
                        _ = v14
                    } else {
                        v69 := GenS6{F0: 65535, F1: "", F2: v9}
                        if len(v3) > 0 {
                            v3[0] = ("foo bar" + "")
                        }
                        var v70 int = fi10
                        _ = v69
                        _ = v70
                    }
                    fmt.Println(v13, fi11, v2)
                    if (((613 <= 517) || (0 == 4294967295)) && (-19 < 100)) {
                        if false {
                            v71 := GenS1{F0: -635, F1: "", F2: v9}
                            switch (fi10 | fi10) {
                            case 128:
                                if true {
                                    if len(v1) > 0 {
                                        v1[0] = (^(fi10))
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = fi10
                                    }
                                }
                            case -1:
                                switch (fi11 >> ((255) & 63)) {
                                case 9223372036854775807:
                                    fmt.Println(fi11)
                                case 1:
                                    switch "foo bar" {
                                    case "hello":
                                        _ = len(v1)
                                    case "":
                                        fmt.Println(v2)
                                        v72 := GenS7{F0: 4294967295, F1: 42}
                                        _ = v72
                                    default:
                                        switch fi11 {
                                        case 9223372036854775807:
                                            var v73 int8 = 42
                                            _ = v73
                                        }
                                    }
                                default:
                                    switch -1 {
                                    case 0:
                                        var v74 uint64 = 1
                                        _ = v74
                                    case 128:
                                        if (!false) {
                                            v9 = (v9 + 42)
                                        } else {
                                            fmt.Println(v13)
                                            v13 += (^(v13))
                                            if false {
                                                if (42 != 1) {
                                                    fmt.Println(fi10)
                                                    v9 += v9
                                                    _ = len(v1)
                                                }
                                                var v75 uint16 = (v9 & (v9 | v9))
                                                fmt.Println(v75, fi11, v2)
                                                _ = v75
                                            } else {
                                                var v76 uint64 = 4294967295
                                                var v77 float32 = 1e10
                                                fmt.Println(v2, v13, fi11)
                                                _ = v76
                                                _ = v77
                                            }
                                        }
                                    case 32767:
                                        fmt.Println(fi11, v9, v9)
                                    default:
                                        var v78 uint16 = 42
                                        _ = v78
                                    }
                                }
                            default:
                                fmt.Println(v9, v9)
                            }
                            var v79 float64 = ((0.1 - (1e10 / -2.25)) * 0.0)
                            _ = v71
                            _ = v79
                        } else {
                            switch "foo bar" {
                            case "hello":
                                if len(v1) > 0 {
                                    v1[0] = (^(fi10))
                                }
                            default:
                                _ = len(v1)
                            }
                            if (42 == 127) {
                                if ((65535 == 277) && ((false || true) && (true && true))) {
                                    var v80 uint = 18446744073709551615
                                    _ = v80
                                }
                            }
                            switch "foo bar" {
                            case "\n\t":
                                v9 = (^(v9))
                            case "hello":
                                var v81 uint64 = uint64(v9)
                                _ = v81
                            default:
                                if true {
                                    var v82 uint8 = 1
                                    _ = len(v1)
                                    _ = v82
                                }
                            }
                        }
                        if (!false) {
                            v2 = (v2 >> ((18446744073709551615) & 63))
                            v1 = append(v1, (fi10 >> ((65535) & 63)))
                            if true {
                                var v83 int16 = -608
                                _ = v83
                            }
                        }
                    }
                    switch "foo bar" {
                    case "\n\t":
                        var v84 float64 = (3.14159 / 1.0)
                        _ = v84
                    }
                    v85 := GenS8{F0: v2, F1: "foo bar", F2: -1}
                    _ = v12
                    _ = v13
                    _ = v85
                } else {
                    fmt.Println(v9, v9, v2)
                }
                v2 = (v2 % (((v2 ^ (v2 / ((1) | 1)))) | 1))
            }
            if false {
                var v86 bool = false
                var v87 int16 = 32767
                if (4294967295 <= uint(v9)) {
                    var v88 float32 = ((1.5 * -2.25) - (1e10 / (float32(1e10) + 1e10)))
                    if true {
                        var v89 float64 = 1e10
                        v86 = true
                        switch ("foo bar" + ((("hello" + "foo bar") + ("hello" + "")) + ("a" + ""))) {
                        case "foo bar":
                            var v90 float64 = 1.5
                            _ = v90
                        case "hello":
                            v91 := GenS7{F0: 1, F1: -702}
                            v87 = (^(v87))
                            var v92 string = ""
                            v93 := GenS9{F0: (fi10 ^ (fi11 % ((32768) | 1)))}
                            v88 = (float32(-2.25) + ((1e10 + (v88 - float32((1e10 / (v89 + v89))))) / v88))
                            _ = v91
                            _ = v92
                            _ = v93
                        case "":
                            var v94 uint8 = 100
                            _ = v94
                        }
                        _ = v89
                    }
                    _ = v88
                }
                _ = v86
                _ = v87
            } else {
                v2 = 100
                v2 -= -1
                var v95 int8 = v2
                _ = v95
            }
            var v96 string = "foo bar"
            _ = fi11
            _ = v96
        }
        fmt.Println(v9)
        _ = fi10
    }
    v97 := GenS6{F0: 352, F1: (("hello" + ("a" + "hello")) + "hello"), F2: 65535}
    if true {
        for fi98 := 0; fi98 < 4; fi98++ {
            _ = len(v3)
            _ = fi98
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (255 <= 40) {
            var v99 uint16 = (v9 - v9)
            v9 += v99
            var v100 int = -1
            _ = v100
            _ = v99
        }
    } else {
        for fi101 := 0; fi101 < 2; fi101++ {
            for fi102 := 0; fi102 < 14; fi102++ {
                if len(v3) > 0 {
                    v3[0] = ""
                }
                _ = fi102
            }
            _ = fi101
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
    _ = v9
    _ = v97
}
