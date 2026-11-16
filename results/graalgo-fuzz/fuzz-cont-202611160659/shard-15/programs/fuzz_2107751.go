package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 string; F1 uint; F2 bool }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint; F1 int8 }
type GenS10 struct { F0 string; F1 uint32 }
type GenS11 struct { F0 uint16; F1 int }

func main() {
    v1 := make(map[int]bool)
    var v2 int = 32768
    v3 := make(map[uint32]string)
    v4 := []uint32{0, 2147483647, 235}
    if len(v4) > 0 {
        v4[0] = 1
    }
    v2 -= (v2 & 1)
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 int8 = 100
        for fi7 := 0; fi7 < 9; fi7++ {
            _ = len(v3)
            if ((831 > 43) && (65535 >= 42)) {
                _ = len(v4)
                if false {
                    if false {
                        v2 = (^(fi5))
                    } else {
                        var v8 uint8 = 1
                        var v9 int = (fi5 % ((485) | 1))
                        fmt.Println(v8, v9)
                        _ = v8
                        _ = v9
                    }
                    if true {
                        if (true || (false && true)) {
                            switch ("a" + "\n\t") {
                            case "\n\t":
                                v3[65535] = "a"
                            case "hello":
                                v2 += -9223372036854775808
                            default:
                                if false {
                                    var v10 uint32 = 0
                                    v2 += (fi5 - (^(fi5)))
                                    v10 = 45
                                    _ = v10
                                }
                            }
                            var v11 float64 = (1e10 / (0.1 + float64(1e10)))
                            v1[(v2 / (((fi5 + fi5)) | 1))] = false
                            _ = v11
                        } else {
                            switch "a" {
                            case "a":
                                _ = len(v1)
                            default:
                                v6 -= 34
                            }
                            if false {
                                if (v6 > (v6 / ((-128) | 1))) {
                                    var v12 uint64 = 1
                                    _ = v12
                                }
                                if (18446744073709551615 != 918) {
                                    var v13 uint = 4294967295
                                    _ = v13
                                }
                                switch (fi7 >> ((1) & 63)) {
                                case -350:
                                    switch fi7 {
                                    case -9223372036854775808:
                                        var v14 uint64 = 2
                                        _ = v14
                                    case 0:
                                        v2 = (v2 - (fi5 >> ((42) & 63)))
                                    default:
                                        var v15 int64 = 1
                                        _ = v15
                                    }
                                    if false {
                                        v16 := GenS1{F0: 1, F1: 127}
                                        fmt.Println(v6, v6, v6)
                                        _ = v16
                                    }
                                    v17 := GenS2{F0: true, F1: 42, F2: 61}
                                    _ = v17
                                }
                            }
                        }
                    }
                }
                v3[0] = ""
            } else {
                var v18 uint = 985
                _ = v18
            }
            fmt.Println(v2, v6, fi5)
            _ = fi7
        }
        if true {
            if false {
                var v19 int16 = -32768
                for fi20 := 0; fi20 < 9; fi20++ {
                    _ = len(v4)
                    var v21 bool = false
                    _ = fi20
                    _ = v21
                }
                switch ("foo bar" + ("" + "")) {
                case "a":
                    if len(v4) > 0 {
                        v4[0] = 4294967295
                    }
                case "":
                    v6 = (v6 | 0)
                case "\n\t":
                    if true {
                        for fi22 := 0; fi22 < 5; fi22++ {
                            v2 = (v2 % ((fi5) | 1))
                            _ = fi22
                        }
                        switch (v2 % ((32768) | 1)) {
                        case 32768:
                            switch v2 {
                            case 2147483647:
                                v6 = (v6 >> ((448) & 63))
                            case 128:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            default:
                                var v23 int8 = (^(v6))
                                _ = v23
                            }
                        }
                    } else {
                        if len(v4) > 0 {
                            v4[0] = 417
                        }
                        if false {
                            v6 = (v6 + v6)
                            switch (v2 << ((65535) & 63)) {
                            case 128:
                                _ = len(v4)
                            case 1:
                                delete(v1, (fi5 & (fi5 | v2)))
                            case 775:
                                fmt.Println(v6, fi5, v6)
                            default:
                                if ((-9223372036854775808 < int64(fi5)) && false) {
                                    if false {
                                        if true {
                                            switch (fi5 + v2) {
                                            case 1:
                                                if (112 > 896) {
                                                    v24 := v3[42]
                                                    v24 = (((v24 + v24) + v24) + "hello")
                                                    for fi25 := 0; fi25 < 18; fi25++ {
                                                        var v26 float32 = 0.1
                                                        _ = fi25
                                                        _ = v26
                                                    }
                                                    _ = v24
                                                }
                                            case 9223372036854775807:
                                                switch v2 {
                                                case 1:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                case -271:
                                                    switch (fi5 & (v2 + fi5)) {
                                                    case -1:
                                                        var v27 int64 = 1
                                                        _ = v27
                                                    case -458:
                                                        var v28 uint8 = 127
                                                        _ = v28
                                                    }
                                                case 818:
                                                    var v29 uint8 = 116
                                                    _ = v29
                                                }
                                            }
                                            if true {
                                                for fi30 := 0; fi30 < 14; fi30++ {
                                                    var v31 int8 = (v6 ^ v6)
                                                    switch (("" + "hello") + "\n\t") {
                                                    case "":
                                                        v6 += v31
                                                    case "a":
                                                        var v32 int16 = (-(v19))
                                                        _ = v32
                                                    default:
                                                        if len(v4) > 0 {
                                                            v4[0] = 507
                                                        }
                                                    }
                                                    _ = fi30
                                                    _ = v31
                                                }
                                                fmt.Println(v2)
                                            }
                                            v33 := GenS3{F0: 0, F1: (fi5 & 32768)}
                                            _ = v33
                                        } else {
                                            var v34 float32 = 3.14159
                                            v2 -= fi5
                                            var v35 int8 = (^(v6))
                                            _ = v34
                                            _ = v35
                                        }
                                        var v36 string = ("hello" + "\n\t")
                                        var v37 int32 = -1
                                        _ = v36
                                        _ = v37
                                    } else {
                                        var v38 bool = (!true)
                                        for fi39 := 0; fi39 < 6; fi39++ {
                                            var v40 float64 = 1e10
                                            _ = fi39
                                            _ = v40
                                        }
                                        _ = v38
                                    }
                                    v4 = append(v4, 65535)
                                    v41 := GenS1{F0: v19, F1: 128}
                                    _ = v41
                                }
                            }
                            _ = len(v4)
                        }
                    }
                }
                _ = v19
            }
        }
        _ = fi5
        _ = v6
    }
    v42 := GenS1{F0: 32767, F1: 140}
    v2 -= v2
    v43 := GenS3{F0: 42, F1: v2}
    if ("" != "\n\t") {
        for fi44 := 0; fi44 < 2; fi44++ {
            switch "hello" {
            case "hello":
                switch (v2 | -2147483648) {
                case 0:
                    v45 := GenS1{F0: 233, F1: 90}
                    _ = v45
                case 727:
                    for fi46 := 0; fi46 < 18; fi46++ {
                        v47 := GenS2{F0: (!true), F1: 868, F2: 1}
                        var v48 uint = 2
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                default:
                    for fi49 := 0; fi49 < 5; fi49++ {
                        if false {
                            v4 = append(v4, 382)
                            switch fi49 {
                            case 32767:
                                var v50 uint64 = 1
                                _ = v50
                            case 2147483647:
                                fmt.Println(fi44, fi49, fi44)
                            case 127:
                                v51 := GenS4{F0: 255, F1: ((("hello" + "") + "foo bar") + "")}
                                _ = v51
                            default:
                                v4 = append(v4, 2)
                            }
                        } else {
                            if false {
                                v2 = (fi49 | (v2 / (((fi44 * 122)) | 1)))
                                _ = len(v4)
                                v2 = v2
                                if len(v4) > 0 {
                                    v4[0] = 83
                                }
                                var v52 int16 = 42
                                fmt.Println(v2)
                                _ = v52
                            } else {
                                fmt.Println(v2, fi44, v2)
                                var v53 uint = 908
                                var v54 uint8 = 0
                                var v55 string = ("foo bar" + "hello")
                                switch (fi44 - v2) {
                                case 2147483647:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case -9223372036854775808:
                                    if true {
                                        fmt.Println(v2)
                                        var v56 uint32 = 2147483647
                                        switch v55 {
                                        case "":
                                            v57 := GenS5{F0: -32768, F1: 788}
                                            _ = v57
                                        case "hello":
                                            v58 := v1[fi49]
                                            _ = v58
                                        default:
                                            var v59 uint = v53
                                            v1[(v2 % (((fi49 - 0)) | 1))] = false
                                            _ = v59
                                        }
                                        _ = v56
                                    } else {
                                        fmt.Println(v53, v53)
                                        delete(v3, 2147483648)
                                        v53 += 255
                                    }
                                default:
                                    v60 := GenS6{F0: v55, F1: 610}
                                    _ = v60
                                }
                                _ = v53
                                _ = v54
                                _ = v55
                            }
                            if true {
                                v61 := GenS5{F0: 1000, F1: 2}
                                _ = len(v3)
                                fmt.Println(v2, fi44)
                                v2 = fi49
                                var v62 uint16 = 42
                                v63 := GenS3{F0: v62, F1: (-(fi44))}
                                v64 := GenS7{F0: "hello", F1: 771, F2: (true || false)}
                                if false {
                                    if (("a" == ("hello" + "\n\t")) || false) {
                                        v65 := GenS2{F0: false, F1: 2, F2: 584}
                                        fmt.Println(fi44, v62)
                                        _ = v65
                                    } else {
                                        delete(v1, 64)
                                    }
                                } else {
                                    switch (("foo bar" + "\n\t") + "") {
                                    case "hello":
                                        switch "hello" {
                                        case "a":
                                            if true {
                                                v66 := GenS8{F0: -415}
                                                _ = v66
                                            }
                                        }
                                    case "\n\t":
                                        var v67 float32 = 3.14159
                                        _ = v67
                                    case "foo bar":
                                        switch ("" + ("\n\t" + "")) {
                                        case "foo bar":
                                            if (!false) {
                                                v68 := GenS9{F0: 4294967295, F1: -1}
                                                v2 -= (^(v2))
                                                _ = v68
                                            }
                                            var v69 uint16 = 1
                                            _ = v69
                                        case "\n\t":
                                            fmt.Println(v2)
                                        case "":
                                            v2 -= 4294967295
                                        }
                                    }
                                }
                                switch "foo bar" {
                                case "a":
                                    if len(v4) > 0 {
                                        v4[0] = 859
                                    }
                                case "":
                                    var v70 float64 = 3.14159
                                    _ = v70
                                }
                                v62 = 2
                                fmt.Println(v62)
                                _ = v61
                                _ = v62
                                _ = v63
                                _ = v64
                            }
                        }
                        if true {
                            v2 += fi49
                            if false {
                                v2 = (fi49 + fi44)
                                if true {
                                    var v71 uint8 = 42
                                    _ = len(v4)
                                    fmt.Println(v71, fi49)
                                    v1[(v2 % ((v2) | 1))] = true
                                    _ = v71
                                }
                            }
                            switch "" {
                            case "foo bar":
                                var v72 uint32 = 2
                                _ = v72
                            case "hello":
                                if true {
                                    switch "\n\t" {
                                    case "":
                                        fmt.Println(fi44)
                                    case "foo bar":
                                        v73 := GenS10{F0: "\n\t", F1: 1}
                                        _ = v73
                                    case "hello":
                                        v74 := GenS4{F0: 2, F1: ""}
                                        _ = v74
                                    }
                                    var v75 int16 = 32767
                                    v75 = v75
                                    _ = v75
                                } else {
                                    v76 := GenS8{F0: 0}
                                    v2 = fi44
                                    _ = v76
                                }
                            default:
                                if (1 < 965) {
                                    v77 := GenS2{F0: true, F1: 18446744073709551615, F2: 0}
                                    v78 := GenS11{F0: 669, F1: fi44}
                                    _ = v77
                                    _ = v78
                                }
                            }
                        }
                        fmt.Println(fi49, v2, fi49)
                        _ = fi49
                    }
                }
            default:
                var v79 int32 = 0
                _ = v79
            }
            _ = fi44
        }
        v80 := GenS1{F0: 1, F1: 1}
        _ = v80
    }
    v81 := GenS5{F0: -1, F1: 4294967295}
    var v82 bool = ((42 < 2) || (-128 >= -128))
    v82 = (515 <= 479)
    v2 += (v2 - (v2 * v2))
    fmt.Println(v2, v82)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v81
    _ = v82
}
