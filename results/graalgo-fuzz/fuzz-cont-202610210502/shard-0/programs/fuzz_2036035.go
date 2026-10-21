package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 bool }
type GenS2 struct { F0 bool; F1 string; F2 uint64 }
type GenS3 struct { F0 bool; F1 int64; F2 int }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int8; F1 int8; F2 int32 }
type GenS9 struct { F0 int8; F1 uint64; F2 int32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]uint32)
    v3 := []uint32{}
    var v4 int = 127
    if len(v3) > 0 {
        v3[0] = 666
    }
    switch "a" {
    case "\n\t":
        var v5 int32 = 429
        _ = v5
    default:
        v4 -= (v4 << ((4294967295) & 63))
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 16; fi7++ {
            v8 := GenS1{F0: 1000, F1: 0, F2: (!false)}
            if false {
                var v9 int16 = -32768
                fmt.Println(v9, v9, fi7)
                _ = v9
            }
            v10 := GenS2{F0: (2 <= 18446744073709551615), F1: "", F2: 1}
            _ = fi7
            _ = v10
            _ = v8
        }
        _ = fi6
    }
    var v11 int64 = 1
    _ = len(v1)
    fmt.Println(v11, v4)
    for fi12 := 0; fi12 < 16; fi12++ {
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 string = "hello"
            _ = fi13
            _ = v14
        }
        var v15 int64 = 0
        _ = fi12
        _ = v15
    }
    delete(v2, (("a" + (("foo bar" + ("hello" + ("" + "hello"))) + "\n\t")) + ("foo bar" + ("foo bar" + "\n\t"))))
    for fi16 := 0; fi16 < 6; fi16++ {
        if (true || (3.14159 == 3.14159)) {
            var v17 uint = 697
            switch "hello" {
            case "":
                switch "foo bar" {
                case "hello":
                    if ((!true) || true) {
                        switch -2147483648 {
                        case 42:
                            v11 += 0
                        case 1:
                            if false {
                                var v18 float64 = float64(((0.1 - ((0.1 / 0.1) * 0.1)) / float32(0.1)))
                                v2[("foo bar" + "\n\t")] = 4294967295
                                fmt.Println(v4, v4, v17)
                                _ = v18
                            }
                        case -9223372036854775808:
                            for fi19 := 0; fi19 < 20; fi19++ {
                                if true {
                                    if true {
                                        var v20 bool = false
                                        if v20 {
                                            v11 -= v11
                                            if (482 <= 1000) {
                                                var v21 int8 = 100
                                                v17 = (v17 + v17)
                                                if false {
                                                    switch (("\n\t" + ("a" + ("a" + "a"))) + ("" + "\n\t")) {
                                                    case "foo bar":
                                                        fmt.Println(v20, v17, fi19)
                                                    default:
                                                        v22 := GenS3{F0: ("hello" <= "a"), F1: (v11 ^ (v11 | 1)), F2: -9223372036854775808}
                                                        _ = v22
                                                    }
                                                    switch "a" {
                                                    case "hello":
                                                        v4 = 2147483647
                                                    default:
                                                        v17 = v17
                                                    }
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    if len(v3) > 0 {
                                                        v3[0] = 0
                                                    }
                                                    v23 := GenS2{F0: v20, F1: "\n\t", F2: 1}
                                                    _ = v23
                                                } else {
                                                    v24 := GenS2{F0: (255 < 2), F1: "\n\t", F2: 0}
                                                    _ = v24
                                                }
                                                _ = v21
                                            }
                                            var v25 bool = v20
                                            _ = v25
                                        }
                                        var v26 uint32 = 42
                                        v27 := GenS2{F0: (v20 || v20), F1: "foo bar", F2: 1}
                                        _ = v20
                                        _ = v26
                                        _ = v27
                                    }
                                    if (!(false || false)) {
                                        var v28 int64 = (-(v11))
                                        if (false && false) {
                                            delete(v1, (fi19 >> ((876) & 63)))
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v17 = 42
                                            v29 := GenS4{F0: 1000, F1: 18446744073709551615}
                                            _ = v29
                                        }
                                        var v30 string = ("\n\t" + "\n\t")
                                        _ = v28
                                        _ = v30
                                    } else {
                                        v4 = v4
                                    }
                                } else {
                                    if false {
                                        switch fi16 {
                                        case -1:
                                            delete(v2, "\n\t")
                                        case 127:
                                            if false {
                                                switch v4 {
                                                case -53:
                                                    if (826 >= 65535) {
                                                        v17 = v17
                                                        var v31 float32 = -2.25
                                                        _ = v31
                                                    }
                                                case 454:
                                                    _ = len(v3)
                                                case 42:
                                                    var v32 uint = v17
                                                    _ = v32
                                                default:
                                                    switch ("foo bar" + "hello") {
                                                    case "hello":
                                                        var v33 int16 = -120
                                                        _ = v33
                                                    case "foo bar":
                                                        var v34 int8 = 42
                                                        _ = v34
                                                    case "a":
                                                        var v35 uint8 = 128
                                                        var v36 uint8 = v35
                                                        var v37 int16 = -79
                                                        _ = v35
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                }
                                                var v38 uint = (v17 ^ v17)
                                                v39 := GenS3{F0: false, F1: 997, F2: fi16}
                                                _ = v38
                                                _ = v39
                                            }
                                        }
                                    }
                                }
                                if (1 == 2147483648) {
                                    var v40 float64 = 0.1
                                    switch ("hello" + "") {
                                    case "hello":
                                        var v41 string = ("a" + ("\n\t" + "hello"))
                                        v4 = v4
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v41
                                    case "\n\t":
                                        v42 := GenS5{F0: 0, F1: v11}
                                        _ = v42
                                    default:
                                        v11 = v11
                                        fmt.Println(v11)
                                        fmt.Println(v17, fi16)
                                        fmt.Println(fi19, v4)
                                        switch "foo bar" {
                                        case "a":
                                            v40 = -2.25
                                        case "foo bar":
                                            if (true && (false || true)) {
                                                if true {
                                                    var v43 float64 = ((v40 * v40) - 1.5)
                                                    v43 += float64((1e10 + 0.1))
                                                    v44 := GenS6{F0: (v17 % (((v17 * v17)) | 1)), F1: -1}
                                                    if (18446744073709551615 != 18446744073709551615) {
                                                        if true {
                                                            v17 -= v17
                                                        }
                                                        v45 := GenS7{F0: fi19}
                                                        v4 -= fi16
                                                        _ = v45
                                                    } else {
                                                        v2["a"] = 1
                                                        _ = len(v1)
                                                        var v46 uint32 = 211
                                                        _ = v46
                                                    }
                                                    if ("foo bar" == "\n\t") {
                                                        var v47 bool = (102 <= 42)
                                                        _ = v47
                                                    } else {
                                                        v43 += (v40 - (0.0 * (v40 - v40)))
                                                    }
                                                    _ = v43
                                                    _ = v44
                                                } else {
                                                    switch fi16 {
                                                    case 4294967295:
                                                        fmt.Println(v4, v4, v4)
                                                    case 127:
                                                        if (true || false) {
                                                            var v48 uint32 = 651
                                                            _ = v48
                                                        } else {
                                                            v2["\n\t"] = 65535
                                                            switch (-(fi19)) {
                                                            case 32767:
                                                                v1[-977] = 257
                                                            case 127:
                                                                var v49 float32 = 1.5
                                                                _ = v49
                                                            case -2147483648:
                                                                fmt.Println(v4, v17, v11)
                                                            }
                                                        }
                                                    default:
                                                        v50 := GenS3{F0: (!true), F1: v11, F2: v4}
                                                        fmt.Println(v4, fi19, v11)
                                                        switch "foo bar" {
                                                        case "":
                                                            switch (("hello" + "hello") + ("\n\t" + "")) {
                                                            case "a":
                                                                var v51 uint64 = 805
                                                                _ = v51
                                                            default:
                                                                var v52 uint16 = 32767
                                                                _ = v52
                                                            }
                                                        case "foo bar":
                                                            if true {
                                                                v3 = append(v3, 2147483648)
                                                            } else {
                                                                v53 := GenS3{F0: ((v17 >> ((v17) & 63)) <= 221), F1: v11, F2: fi16}
                                                                switch fi19 {
                                                                case 4294967295:
                                                                    v54 := GenS7{F0: (v4 + (fi19 >> (((v17 & 0)) & 63)))}
                                                                    _ = v54
                                                                case -2147483648:
                                                                    if ((!false) || ((1.5 == 1e10) || (!(!(!true))))) {
                                                                        v4 = (^(fi19))
                                                                    } else {
                                                                        if ((((("hello" + "\n\t") + "foo bar") + ("foo bar" + "")) + "\n\t") < "foo bar") {
                                                                            var v55 string = "\n\t"
                                                                            _ = v55
                                                                        }
                                                                    }
                                                                    if (("hello" + "foo bar") >= "hello") {
                                                                        v17 = 2
                                                                    }
                                                                    if (9223372036854775807 != (^(v11))) {
                                                                        v17 += v17
                                                                        v3 = append(v3, 654)
                                                                    } else {
                                                                        v17 += v17
                                                                        fmt.Println(fi19, v4)
                                                                    }
                                                                case -168:
                                                                    v3 = append(v3, 0)
                                                                default:
                                                                    fmt.Println(v4, v11, v4)
                                                                }
                                                                if false {
                                                                    fmt.Println(v11, v40, v11)
                                                                } else {
                                                                    v11 -= 1
                                                                }
                                                                _ = v53
                                                            }
                                                        }
                                                        _ = v50
                                                    }
                                                    if true {
                                                        v56 := GenS4{F0: 199, F1: 42}
                                                        var v57 int32 = -2147483648
                                                        var v58 uint32 = 2147483647
                                                        _ = v56
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                    switch "foo bar" {
                                                    case "":
                                                        v59 := GenS7{F0: (fi19 * 127)}
                                                        _ = v59
                                                    case "foo bar":
                                                        if false {
                                                            v60 := GenS1{F0: -1, F1: v11, F2: false}
                                                            var v61 string = ("a" + "hello")
                                                            v11 = 907
                                                            _ = v60
                                                            _ = v61
                                                        }
                                                    case "a":
                                                        delete(v1, (fi16 + fi19))
                                                    default:
                                                        v62 := GenS4{F0: -619, F1: 4294967295}
                                                        var v63 int32 = 1
                                                        v64 := GenS2{F0: (!true), F1: "", F2: 18446744073709551615}
                                                        v65 := GenS4{F0: -32768, F1: 719}
                                                        fmt.Println(fi19)
                                                        if (1 >= 42) {
                                                            switch (fi19 % (((v4 + v4)) | 1)) {
                                                            case -293:
                                                                switch ((("foo bar" + "hello") + "foo bar") + ("" + ("hello" + "a"))) {
                                                                case "":
                                                                    var v66 float32 = 1.5
                                                                    _ = v66
                                                                case "foo bar":
                                                                    if (v4 < 2147483647) {
                                                                        v40 = ((v40 - float64(1e10)) * 3.14159)
                                                                    } else {
                                                                        v17 += 65535
                                                                        v67 := GenS6{F0: v17, F1: 1}
                                                                        var v68 float64 = float64(1e10)
                                                                        _ = v67
                                                                        _ = v68
                                                                    }
                                                                }
                                                            case 1:
                                                                v17 -= v17
                                                            case 127:
                                                                fmt.Println(v63, fi19, v17)
                                                            }
                                                            v69 := GenS5{F0: 32767, F1: 0}
                                                            _ = v69
                                                        }
                                                        v17 = (^(v17))
                                                        _ = v62
                                                        _ = v63
                                                        _ = v64
                                                        _ = v65
                                                    }
                                                    v70 := GenS8{F0: -1, F1: -128, F2: 42}
                                                    _ = v70
                                                }
                                            }
                                        case "":
                                            var v71 uint64 = 10
                                            _ = v71
                                        default:
                                            var v72 uint64 = 1
                                            _ = v72
                                        }
                                    }
                                    _ = v40
                                } else {
                                    var v73 int16 = 532
                                    var v74 int32 = 42
                                    _ = v73
                                    _ = v74
                                }
                                _ = fi19
                            }
                        default:
                            var v75 uint = (^(v17))
                            _ = v75
                        }
                    }
                default:
                    v76 := v1[9223372036854775807]
                    _ = v76
                }
            case "\n\t":
                v17 = 255
            default:
                for fi77 := 0; fi77 < 19; fi77++ {
                    switch fi16 {
                    case 128:
                        v78 := v1[v4]
                        _ = v78
                    case 32768:
                        var v79 int16 = 32767
                        _ = v79
                    case -2147483648:
                        fmt.Println(fi16)
                    default:
                        var v80 int = fi77
                        _ = v80
                    }
                    var v81 int = (fi16 / (((fi16 % ((fi16) | 1))) | 1))
                    _ = fi77
                    _ = v81
                }
            }
            _ = v17
        } else {
            for fi82 := 0; fi82 < 2; fi82++ {
                var v83 string = ""
                fmt.Println(v4, fi16, v83)
                switch v4 {
                case 32767:
                    fmt.Println(v11, v11)
                case 4294967295:
                    v83 = v83
                case -802:
                    if true {
                        var v84 int32 = -2147483648
                        if false {
                            if (339 <= 25) {
                                if false {
                                    var v85 bool = (true || ((!false) && false))
                                    v84 -= v84
                                    _ = v85
                                }
                                var v86 uint = 255
                                _ = v86
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = 506
                                }
                                v83 = (v83 + (v83 + v83))
                            }
                        }
                        switch (v4 & fi16) {
                        case 0:
                            if true {
                                var v87 int16 = int16(v84)
                                _ = v87
                            } else {
                                fmt.Println(fi82)
                                var v88 int16 = 1
                                v4 = 128
                                _ = v88
                            }
                        default:
                            _ = len(v3)
                        }
                        var v89 string = ("foo bar" + v83)
                        var v90 float64 = -2.25
                        _ = v84
                        _ = v89
                        _ = v90
                    }
                default:
                    v91 := GenS9{F0: 1, F1: 42, F2: 0}
                    _ = v91
                }
                _ = fi82
                _ = v83
            }
            v11 = -1
            for fi92 := 0; fi92 < 13; fi92++ {
                v11 -= v11
                v4 = v4
                if len(v3) > 0 {
                    v3[0] = 65535
                }
                var v93 uint64 = 2
                var v94 uint64 = (v93 >> ((1) & 63))
                _ = fi92
                _ = v93
                _ = v94
            }
        }
        _ = fi16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
