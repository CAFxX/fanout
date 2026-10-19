package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 string }

func main() {
    var v1 uint = 65535
    var v2 uint64 = 533
    var v3 uint = (v1 * 264)
    var v4 float64 = 3.14159
    v5 := make(map[uint32]uint32)
    v6 := make(map[uint32]uint32)
    if (255 >= 42) {
        if (!(!((0 >= 50) || (1 > 690)))) {
            var v7 int32 = 966
            fmt.Println(v1, v3, v4)
            v1 += v3
            _ = v7
        } else {
            v2 = v2
            switch "hello" {
            case "a":
                v8 := GenS1{F0: -2147483648, F1: 1}
                _ = v8
            case "":
                var v9 int16 = 0
                _ = v9
            default:
                v10 := []int{-1}
                _ = v10
            }
            var v11 uint16 = 42
            _ = v11
        }
    } else {
        v12 := GenS2{F0: 42}
        _ = v12
    }
    v13 := []int{843, 9223372036854775807, 118, -435}
    fmt.Println(v1)
    if (1.5 <= 1e10) {
        var v14 uint32 = 617
        _ = v14
    } else {
        v2 = 0
        v5[2] = 65535
        for fi15 := 0; fi15 < 19; fi15++ {
            fmt.Println(v4, fi15)
            if true {
                v2 -= v2
                for fi16 := 0; fi16 < 20; fi16++ {
                    v1 = v1
                    if false {
                        switch ((("" + "hello") + (("" + (("\n\t" + "") + ("" + "\n\t"))) + ("a" + (("hello" + "hello") + "foo bar")))) + "foo bar") {
                        case "hello":
                            var v17 bool = true
                            _ = v17
                        default:
                            v18 := GenS2{F0: 0}
                            if true {
                                if (!true) {
                                    v1 += (v1 - v1)
                                    v4 = (v4 / ((v4 + v4) * v4))
                                    var v19 int = (fi15 * 149)
                                    _ = v19
                                } else {
                                    v20 := GenS2{F0: -200}
                                    var v21 string = ((("" + "hello") + "\n\t") + ((("hello" + (("\n\t" + "hello") + "hello")) + "a") + ""))
                                    _ = v20
                                    _ = v21
                                }
                                if ((("foo bar" + (("\n\t" + "\n\t") + "a")) + "hello") > "a") {
                                    if (!(233 == 84)) {
                                        var v22 string = ("a" + (("\n\t" + "") + (("a" + "a") + "hello")))
                                        switch (fi16 << ((1) & 63)) {
                                        case -9223372036854775808:
                                            delete(v6, 2147483648)
                                        case -1:
                                            v13 = append(v13, -9223372036854775808)
                                        case 0:
                                            var v23 int = (^(fi16))
                                            _ = v23
                                        default:
                                            var v24 uint64 = v2
                                            _ = v24
                                        }
                                        v25 := GenS3{F0: 1}
                                        _ = v22
                                        _ = v25
                                    }
                                } else {
                                    switch ("\n\t" + "\n\t") {
                                    case "":
                                        switch (("\n\t" + "\n\t") + "\n\t") {
                                        case "":
                                            if len(v13) >= 2 {
                                                v13 = v13[1:len(v13)]
                                            }
                                        case "hello":
                                            var v26 int = 456
                                            _ = v26
                                        case "a":
                                            _ = len(v13)
                                        }
                                    case "\n\t":
                                        if (0 == 1) {
                                            v27 := GenS2{F0: -654}
                                            _ = v27
                                        }
                                    case "foo bar":
                                        var v28 bool = (300 <= 65535)
                                        _ = v28
                                    default:
                                        var v29 int64 = 0
                                        _ = v29
                                    }
                                }
                                fmt.Println(fi15, v4)
                            }
                            _ = v18
                        }
                    }
                    _ = fi16
                }
                fmt.Println(v3, v3)
            }
            fmt.Println(v2)
            _ = fi15
        }
    }
    if (true || (true && true)) {
        for fi30 := 0; fi30 < 1; fi30++ {
            v1 += v3
            if len(v13) > 0 {
                v13[0] = (fi30 & (fi30 | int(v3)))
            }
            switch fi30 {
            case 128:
                switch (("" + ("" + "\n\t")) + "") {
                case "hello":
                    fmt.Println(v1, v2)
                }
            default:
                v4 += float64(0.1)
            }
            _ = fi30
        }
    }
    fmt.Println(v4, v4, v4)
    for fi31 := 0; fi31 < 5; fi31++ {
        if true {
            v3 -= (v3 ^ (v1 ^ v3))
        } else {
            v32 := GenS1{F0: -108, F1: 381}
            fmt.Println(v1, v1, fi31)
            _ = v32
        }
        v13 = append(v13, fi31)
        for fi33 := 0; fi33 < 7; fi33++ {
            switch -1 {
            case -2147483648:
                fmt.Println(v3)
            default:
                var v34 int32 = 2147483647
                _ = v34
            }
            switch "a" {
            case "":
                v1 -= (v1 % (((v1 >> ((v1) & 63))) | 1))
            default:
                v35 := GenS4{F0: 1, F1: 1}
                _ = v35
            }
            v3 += (v3 | v1)
            _ = fi33
        }
        _ = fi31
    }
    fmt.Println(v4, v4)
    fmt.Println(v2, v2, v4)
    v1 += (v1 % ((811) | 1))
    if len(v13) > 0 {
        v13[0] = -1
    }
    if ("" < "\n\t") {
        v2 += v2
    } else {
        for fi36 := 0; fi36 < 1; fi36++ {
            var v37 int16 = 1000
            if len(v13) > 0 {
                v13[0] = fi36
            }
            delete(v5, 786)
            _ = fi36
            _ = v37
        }
        switch -376 {
        case 9223372036854775807:
            if false {
                v38 := GenS5{F0: ""}
                fmt.Println(v2, v1, v4)
                switch -408 {
                case 32767:
                    if true {
                        var v39 int16 = 1000
                        for fi40 := 0; fi40 < 14; fi40++ {
                            for fi41 := 0; fi41 < 9; fi41++ {
                                _ = len(v13)
                                v42 := GenS2{F0: 1}
                                _ = fi41
                                _ = v42
                            }
                            v4 += v4
                            _ = fi40
                        }
                        if false {
                            switch -1 {
                            case 2147483647:
                                v3 -= (^(v3))
                            case 9223372036854775807:
                                if ((false || (("" + "a") <= ("\n\t" + ""))) && false) {
                                    var v43 int32 = -1
                                    fmt.Println(v39, v2, v43)
                                    _ = v43
                                } else {
                                    if false {
                                        v44 := GenS5{F0: "hello"}
                                        if ((false && true) && ((v39 * v39) > v39)) {
                                            var v45 int8 = 0
                                            _ = v45
                                        } else {
                                            v46 := GenS1{F0: -818, F1: -1}
                                            _ = v46
                                        }
                                        v2 -= (v2 - 2)
                                        _ = v44
                                    }
                                    if (!false) {
                                        v2 = uint64(v3)
                                    } else {
                                        if (4294967295 <= v2) {
                                            var v47 int = 1
                                            _ = v47
                                        } else {
                                            _ = len(v5)
                                        }
                                        var v48 uint32 = 2147483647
                                        switch "foo bar" {
                                        case "":
                                            v3 -= (v1 * (v1 % ((v3) | 1)))
                                        case "\n\t":
                                            if (!false) {
                                                v1 = 377
                                                if (42 <= 2) {
                                                    v4 = (v4 + (v4 + 0.0))
                                                    v39 -= 0
                                                    var v49 bool = false
                                                    _ = v49
                                                } else {
                                                    var v50 float32 = (1.5 + 0.1)
                                                    fmt.Println(v2)
                                                    _ = v50
                                                }
                                            }
                                        default:
                                            if (0 >= 128) {
                                                fmt.Println(v1)
                                                v2 = 4294967295
                                                if false {
                                                    for fi51 := 0; fi51 < 10; fi51++ {
                                                        fmt.Println(fi51)
                                                        for fi52 := 0; fi52 < 18; fi52++ {
                                                            var v53 int64 = -9223372036854775808
                                                            if true {
                                                                var v54 int8 = 1
                                                                switch fi51 {
                                                                case 158:
                                                                    fmt.Println(fi51, v3)
                                                                case 0:
                                                                    var v55 int64 = (v53 ^ (v53 | (-(v53))))
                                                                    _ = v55
                                                                }
                                                                var v56 int64 = v53
                                                                _ = v54
                                                                _ = v56
                                                            } else {
                                                                switch fi52 {
                                                                case 32768:
                                                                    v57 := GenS4{F0: v48, F1: (v48 | (v48 & (v48 - (v48 | 2147483647))))}
                                                                    _ = v57
                                                                case -266:
                                                                    switch 782 {
                                                                    case 4294967295:
                                                                        if len(v13) > 0 {
                                                                            v13[0] = fi52
                                                                        }
                                                                    case 0:
                                                                        var v58 uint = v1
                                                                        var v59 uint8 = 127
                                                                        _ = v58
                                                                        _ = v59
                                                                    }
                                                                default:
                                                                    v4 -= v4
                                                                }
                                                            }
                                                            var v60 float32 = (1e10 / float32((v4 / v4)))
                                                            _ = fi52
                                                            _ = v53
                                                            _ = v60
                                                        }
                                                        _ = fi51
                                                    }
                                                    var v61 uint64 = v2
                                                    if (true || (42 < 584)) {
                                                        if (false && (49 <= 42)) {
                                                            if (((!true) || (-1 != 0)) || false) {
                                                                _ = len(v5)
                                                                for fi62 := 0; fi62 < 12; fi62++ {
                                                                    v63 := GenS1{F0: -496, F1: -388}
                                                                    v4 -= v4
                                                                    v64 := v5[(v48 | 76)]
                                                                    _ = fi62
                                                                    _ = v63
                                                                    _ = v64
                                                                }
                                                            } else {
                                                                v65 := GenS1{F0: -2147483648, F1: -74}
                                                                v66 := GenS4{F0: v48, F1: (v48 & 42)}
                                                                var v67 string = "a"
                                                                _ = v65
                                                                _ = v66
                                                                _ = v67
                                                            }
                                                        }
                                                    } else {
                                                        var v68 bool = true
                                                        var v69 float64 = (v4 - 0.1)
                                                        _ = v68
                                                        _ = v69
                                                    }
                                                    _ = v61
                                                } else {
                                                    for fi70 := 0; fi70 < 3; fi70++ {
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            if len(v13) > 0 {
                                                                v13[0] = (fi70 % (((fi70 & (fi70 - (fi70 % (((fi70 ^ fi70)) | 1))))) | 1))
                                                            }
                                                        case "a":
                                                            var v71 uint8 = 80
                                                            _ = v71
                                                        }
                                                        if len(v13) > 0 {
                                                            v13[0] = fi70
                                                        }
                                                        _ = fi70
                                                    }
                                                    v1 -= v1
                                                }
                                            }
                                        }
                                        _ = v48
                                    }
                                }
                            case 32768:
                                var v72 bool = false
                                _ = v72
                            default:
                                if false {
                                    fmt.Println(v2, v1)
                                    delete(v6, 2147483647)
                                } else {
                                    var v73 bool = true
                                    switch 127 {
                                    case 81:
                                        for fi74 := 0; fi74 < 5; fi74++ {
                                            v13 = append(v13, -9223372036854775808)
                                            _ = fi74
                                        }
                                    case 32767:
                                        var v75 int8 = -120
                                        _ = v75
                                    default:
                                        if (!(42 < (v2 / (((v2 / ((v2) | 1))) | 1)))) {
                                            var v76 uint8 = 0
                                            _ = v76
                                        } else {
                                            var v77 string = "foo bar"
                                            v1 = 112
                                            fmt.Println(v77, v77)
                                            _ = v77
                                        }
                                    }
                                    _ = v73
                                }
                            }
                            if (!(!(!true))) {
                                for fi78 := 0; fi78 < 19; fi78++ {
                                    switch fi78 {
                                    case -1:
                                        v13 = append(v13, fi78)
                                    case 42:
                                        v1 = (v1 ^ (^(v3)))
                                    default:
                                        for fi79 := 0; fi79 < 1; fi79++ {
                                            switch "hello" {
                                            case "\n\t":
                                                delete(v6, 2147483647)
                                            }
                                            _ = fi79
                                        }
                                    }
                                    _ = fi78
                                }
                                v80 := GenS3{F0: 379}
                                _ = v80
                            } else {
                                if false {
                                    v81 := GenS5{F0: ("foo bar" + "\n\t")}
                                    var v82 string = "\n\t"
                                    _ = v81
                                    _ = v82
                                }
                                var v83 float32 = (-2.25 + -2.25)
                                _ = v83
                            }
                            v3 = v1
                        }
                        _ = v39
                    }
                case 1:
                    fmt.Println(v3, v4, v3)
                case 861:
                    for fi84 := 0; fi84 < 2; fi84++ {
                        for fi85 := 0; fi85 < 15; fi85++ {
                            v2 = (v2 >> ((v1) & 63))
                            fmt.Println(fi84, fi85, v4)
                            _ = fi85
                        }
                        fmt.Println(v4, v3, v1)
                        delete(v5, 2147483648)
                        _ = fi84
                    }
                default:
                    var v86 int8 = -1
                    _ = v86
                }
                _ = v38
            } else {
                if (false && (uint32(v3) <= 445)) {
                    var v87 int64 = 9223372036854775807
                    _ = v87
                }
                v1 += v1
                v4 += (v4 - v4)
            }
        case 32768:
            switch -1 {
            case 2147483647:
                v2 = 2
            case -1:
                var v88 uint16 = 255
                _ = v88
            case 1:
                v89 := v5[4294967295]
                _ = v89
            }
        case 525:
            var v90 bool = true
            _ = v90
        }
        v2 -= v2
    }
    v6[575] = 42
    var v91 uint64 = v2
    if true {
        v2 += 328
    } else {
        var v92 uint = (v1 >> ((42) & 63))
        v91 = 0
        v3 = v1
        _ = v92
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v91:", v91)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v91
}
