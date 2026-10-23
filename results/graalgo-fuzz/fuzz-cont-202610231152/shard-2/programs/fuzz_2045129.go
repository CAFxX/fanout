package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 string; F2 uint }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint8; F1 uint8; F2 int8 }

func main() {
    var v1 int8 = 0
    var v2 int = 4294967295
    v3 := make(map[string]int)
    v4 := []uint32{2147483647, 0, 65535}
    v5 := GenS1{F0: (v1 << ((18446744073709551615) & 63)), F1: 699}
    v1 -= v1
    if false {
        if true {
            if false {
                v3[""] = 128
                v1 += -128
                switch ((("a" + "hello") + "hello") + "foo bar") {
                case "a":
                    var v6 bool = (false || true)
                    _ = v6
                }
            } else {
                v2 += v2
                if true {
                    if (!(-1 == -559)) {
                        if true {
                            for fi7 := 0; fi7 < 7; fi7++ {
                                v8 := v3["hello"]
                                v8 = fi7
                                _ = fi7
                                _ = v8
                            }
                            v1 += -128
                        }
                        v2 = (v2 | (v2 + (v2 % ((v2) | 1))))
                    }
                    if (42 != 0) {
                        fmt.Println(v2)
                    } else {
                        fmt.Println(v2, v2)
                        var v9 uint64 = 451
                        for fi10 := 0; fi10 < 17; fi10++ {
                            for fi11 := 0; fi11 < 20; fi11++ {
                                v9 = (v9 ^ (^(v9)))
                                switch (-(v2)) {
                                case -2147483648:
                                    v12 := GenS1{F0: (v1 - (v1 | (v1 % ((v1) | 1)))), F1: 1}
                                    _ = v12
                                case 42:
                                    fmt.Println(v2, v1)
                                case 9223372036854775807:
                                    if (!(!(!(132 != 129)))) {
                                        if false {
                                            var v13 bool = true
                                            if len(v4) > 0 {
                                                v4[0] = 2147483647
                                            }
                                            if v13 {
                                                v14 := GenS2{F0: (v9 & (^(v9)))}
                                                fmt.Println(v13, fi10)
                                                _ = v14
                                            } else {
                                                switch 217 {
                                                case 42:
                                                    if ((0.0 + 3.14159) != 3.14159) {
                                                        switch ("hello" + ((("\n\t" + (("" + "a") + "foo bar")) + (("\n\t" + "\n\t") + "")) + "hello")) {
                                                        case "\n\t":
                                                            v9 = v9
                                                        case "a":
                                                            v1 += (^(v1))
                                                        case "hello":
                                                            var v15 uint = 18446744073709551615
                                                            _ = v15
                                                        default:
                                                            fmt.Println(v9, v13, fi11)
                                                        }
                                                        var v16 float32 = (1.5 + -2.25)
                                                        _ = v16
                                                    }
                                                case 32768:
                                                    var v17 float64 = ((1e10 / float64(1e10)) - 3.14159)
                                                    _ = v17
                                                case -473:
                                                    fmt.Println(v1, v1)
                                                }
                                                v18 := GenS2{F0: v9}
                                                _ = v18
                                            }
                                            _ = v13
                                        } else {
                                            v19 := GenS2{F0: 1}
                                            var v20 uint64 = (v9 / ((v9) | 1))
                                            fmt.Println(v2, v1)
                                            var v21 bool = false
                                            fmt.Println(v21, v2)
                                            _ = v19
                                            _ = v20
                                            _ = v21
                                        }
                                        v2 = (fi11 << ((2) & 63))
                                    }
                                default:
                                    switch "\n\t" {
                                    case "\n\t":
                                        switch 32767 {
                                        case 32767:
                                            if ((42 != 4294967295) || false) {
                                                switch fi11 {
                                                case -676:
                                                    if len(v4) > 0 {
                                                        v4[0] = 65535
                                                    }
                                                case -9223372036854775808:
                                                    if len(v4) > 0 {
                                                        v4[0] = 127
                                                    }
                                                }
                                                fmt.Println(v1, fi11, fi10)
                                            }
                                        default:
                                            if (1 != 0) {
                                                _ = len(v3)
                                                var v22 int64 = 1
                                                var v23 int32 = 0
                                                v3[((("a" + ("\n\t" + ("\n\t" + ""))) + "a") + "")] = v2
                                                if false {
                                                    fmt.Println(fi11, v1, fi11)
                                                }
                                                _ = v22
                                                _ = v23
                                            }
                                        }
                                    }
                                }
                                v24 := GenS2{F0: (v9 | v9)}
                                _ = fi11
                                _ = v24
                            }
                            var v25 float64 = ((float64(-2.25) + (-2.25 * (3.14159 - -2.25))) - 0.1)
                            var v26 int64 = int64(v1)
                            _ = fi10
                            _ = v25
                            _ = v26
                        }
                        _ = v9
                    }
                    fmt.Println(v2, v1)
                } else {
                    fmt.Println(v1, v1)
                    var v27 int = -96
                    for fi28 := 0; fi28 < 3; fi28++ {
                        var v29 float32 = 0.0
                        _ = fi28
                        _ = v29
                    }
                    _ = v27
                }
            }
            if false {
                if (2147483647 >= -2147483648) {
                    var v30 int = 128
                    _ = v30
                } else {
                    if (v1 == 42) {
                        v1 -= v1
                        if (!(v2 >= (v2 >> ((0) & 63)))) {
                            v2 += (^(v2))
                            v2 += 1
                        }
                        switch (^(v2)) {
                        case -312:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case 32767:
                            if false {
                                v3[""] = 2147483647
                                switch "hello" {
                                case "a":
                                    v31 := GenS3{F0: (v1 * v1)}
                                    _ = v31
                                }
                            } else {
                                var v32 uint8 = 250
                                switch (v2 + -1) {
                                case -4:
                                    for fi33 := 0; fi33 < 3; fi33++ {
                                        for fi34 := 0; fi34 < 11; fi34++ {
                                            var v35 int = 1
                                            _ = fi34
                                            _ = v35
                                        }
                                        var v36 uint64 = 845
                                        _ = fi33
                                        _ = v36
                                    }
                                case 32767:
                                    v37 := GenS4{F0: -9223372036854775808}
                                    _ = v37
                                case 4294967295:
                                    for fi38 := 0; fi38 < 20; fi38++ {
                                        var v39 uint64 = 0
                                        var v40 int = fi38
                                        v3["hello"] = v40
                                        _ = fi38
                                        _ = v39
                                        _ = v40
                                    }
                                default:
                                    v3[""] = v2
                                }
                                v32 = v32
                                _ = v32
                            }
                        case 4294967295:
                            v1 -= (v1 & v1)
                        }
                    } else {
                        v1 += v1
                    }
                    _ = len(v4)
                }
            } else {
                for fi41 := 0; fi41 < 13; fi41++ {
                    fmt.Println(v1)
                    v42 := GenS1{F0: v1, F1: 65535}
                    var v43 bool = true
                    _ = fi41
                    _ = v42
                    _ = v43
                }
            }
            fmt.Println(v1)
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi44 := 0; fi44 < 3; fi44++ {
                v45 := GenS4{F0: 0}
                v1 = -128
                var v46 float32 = 0.1
                _ = fi44
                _ = v45
                _ = v46
            }
        }
        for fi47 := 0; fi47 < 6; fi47++ {
            if (-482 < 1) {
                var v48 int64 = 9223372036854775807
                _ = v48
            }
            _ = fi47
        }
    } else {
        fmt.Println(v1)
        if (42 <= 881) {
            var v49 uint8 = 46
            var v50 int16 = 397
            _ = v49
            _ = v50
        }
    }
    switch (-(v2)) {
    case 42:
        v2 = v2
    default:
        if true {
            v51 := GenS3{F0: (v1 % (((v1 | (v1 | v1))) | 1))}
            for fi52 := 0; fi52 < 2; fi52++ {
                fmt.Println(fi52, v1)
                v53 := GenS1{F0: (v1 * v1), F1: 457}
                var v54 float64 = (3.14159 + -2.25)
                _ = fi52
                _ = v53
                _ = v54
            }
            _ = v51
        } else {
            var v55 bool = (4294967295 < 166)
            if (false && (994 == 42)) {
                switch 9223372036854775807 {
                case -2147483648:
                    for fi56 := 0; fi56 < 8; fi56++ {
                        v57 := GenS5{F0: (((float32(1e10) * ((-2.25 * (1e10 - ((-2.25 / 1.0) + (1e10 + ((-2.25 / (-2.25 / (float32(1e10) + (float32(((3.14159 * 3.14159) - 0.0)) + 1.5)))) + float32((1e10 - 1.5))))))) - (0.0 / 1e10))) > float32(((1.5 * (((1.5 / (1e10 / -2.25)) + 1.5) - float64(1e10))) * (((((0.1 + 3.14159) + (3.14159 + 3.14159)) / 0.1) * 3.14159) + 1.5)))) && true), F1: "hello", F2: 65535}
                        fmt.Println(v2, fi56)
                        _ = fi56
                        _ = v57
                    }
                case 32767:
                    if true {
                        v58 := GenS1{F0: (v1 << ((0) & 63)), F1: 65535}
                        var v59 float32 = (1.5 * 1e10)
                        _ = v58
                        _ = v59
                    }
                case -1:
                    if false {
                        v60 := GenS3{F0: v1}
                        for fi61 := 0; fi61 < 4; fi61++ {
                            for fi62 := 0; fi62 < 13; fi62++ {
                                v63 := GenS5{F0: (42 == int32(v2)), F1: "a", F2: 65535}
                                _ = len(v3)
                                v64 := GenS5{F0: true, F1: ("" + "\n\t"), F2: 255}
                                _ = fi62
                                _ = v63
                                _ = v64
                            }
                            v1 = 4
                            _ = fi61
                        }
                        _ = v60
                    } else {
                        v2 -= v2
                        if v55 {
                            var v65 int = (^(v2))
                            _ = v65
                        }
                    }
                default:
                    var v66 uint = 65535
                    _ = v66
                }
            }
            var v67 uint8 = 2
            _ = v55
            _ = v67
        }
    }
    var v68 int = v2
    if (true && true) {
        switch "\n\t" {
        case "hello":
            v68 = 128
        case "a":
            switch 0 {
            case 9223372036854775807:
                for fi69 := 0; fi69 < 4; fi69++ {
                    for fi70 := 0; fi70 < 18; fi70++ {
                        v68 = fi69
                        v2 = -317
                        _ = fi70
                    }
                    switch (fi69 / (((v68 | v68)) | 1)) {
                    case -571:
                        switch "foo bar" {
                        case "foo bar":
                            for fi71 := 0; fi71 < 9; fi71++ {
                                if ((2 > 42) || (true && false)) {
                                    v2 = 4294967295
                                }
                                v72 := GenS6{F0: "hello"}
                                _ = fi71
                                _ = v72
                            }
                        default:
                            v2 += v68
                        }
                    case -768:
                        var v73 string = ("foo bar" + ((("" + "\n\t") + "a") + ((("a" + "\n\t") + "foo bar") + "\n\t")))
                        _ = v73
                    case 128:
                        if (("foo bar" <= ("hello" + ((("foo bar" + ("hello" + "")) + "a") + ((("" + "") + "\n\t") + "hello")))) || false) {
                            v74 := GenS4{F0: 636}
                            var v75 int16 = 32767
                            v68 = v2
                            _ = v74
                            _ = v75
                        }
                    default:
                        switch 32767 {
                        case -374:
                            var v76 int = (fi69 << ((512) & 63))
                            _ = v76
                        case 711:
                            var v77 uint32 = 2147483648
                            _ = v77
                        default:
                            var v78 int32 = 1
                            _ = v78
                        }
                    }
                    v2 = (v2 % ((v2) | 1))
                    _ = fi69
                }
            default:
                var v79 uint = 2
                _ = v79
            }
        }
        if (32767 == 914) {
            for fi80 := 0; fi80 < 9; fi80++ {
                for fi81 := 0; fi81 < 9; fi81++ {
                    v1 = -80
                    _ = fi81
                }
                _ = fi80
            }
        }
    } else {
        fmt.Println(v2, v2)
        v1 = (^(v1))
        v82 := GenS3{F0: v1}
        _ = v82
    }
    for fi83 := 0; fi83 < 6; fi83++ {
        switch "" {
        case "foo bar":
            var v84 uint64 = 782
            _ = v84
        case "\n\t":
            v68 = v2
        case "":
            fmt.Println(v2)
        default:
            if (3.14159 >= -2.25) {
                var v85 float32 = 3.14159
                var v86 uint = 0
                _ = v85
                _ = v86
            }
        }
        var v87 int = (v68 ^ 9223372036854775807)
        _ = fi83
        _ = v87
    }
    var v88 uint = 915
    v88 = 943
    v89 := GenS2{F0: 4294967295}
    if (!false) {
        for fi90 := 0; fi90 < 10; fi90++ {
            for fi91 := 0; fi91 < 1; fi91++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                switch (("a" + ("foo bar" + "")) + "a") {
                case "":
                    var v92 uint8 = 128
                    _ = v92
                case "a":
                    var v93 int = (fi91 - (v68 ^ v68))
                    _ = v93
                default:
                    v94 := GenS6{F0: ""}
                    _ = v94
                }
                _ = fi91
            }
            _ = fi90
        }
    }
    for fi95 := 0; fi95 < 11; fi95++ {
        v2 = (-(v68))
        _ = fi95
    }
    v68 = (v2 + -204)
    if (0 <= 32767) {
        v96 := GenS7{F0: 234, F1: 195, F2: (-(v1))}
        _ = v96
    } else {
        v3["hello"] = 42
        for fi97 := 0; fi97 < 1; fi97++ {
            var v98 uint64 = 135
            v4 = append(v4, 769)
            _ = fi97
            _ = v98
        }
    }
    v2 -= (v2 + v68)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v68:", v68)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v68
    _ = v88
    _ = v89
}
