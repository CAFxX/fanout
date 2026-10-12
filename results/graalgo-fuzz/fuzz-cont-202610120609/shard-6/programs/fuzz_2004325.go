package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 uint8; F1 int8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int32; F1 uint64; F2 int }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int16; F1 uint16 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint; F1 uint8; F2 int32 }
type GenS13 struct { F0 int8 }

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{}
    v3 := []uint32{2147483648, 3, 42, 615}
    var v4 uint = 65535
    v5 := make(map[int]string)
    fmt.Println(v4)
    var v6 float64 = (0.0 / (0.0 + 3.14159))
    var v7 uint16 = 717
    for fi8 := 0; fi8 < 19; fi8++ {
        v1[(fi8 + 891)] = false
        for fi9 := 0; fi9 < 17; fi9++ {
            var v10 string = "\n\t"
            _ = fi9
            _ = v10
        }
        if ((9223372036854775807 >= (fi8 & (fi8 / (((-(fi8))) | 1)))) || (true && ("a" <= ("a" + "\n\t")))) {
            for fi11 := 0; fi11 < 8; fi11++ {
                fmt.Println(v4, fi8)
                _ = fi11
            }
            v5[(fi8 >> ((135) & 63))] = (("\n\t" + "") + "foo bar")
            fmt.Println(v7)
        } else {
            for fi12 := 0; fi12 < 13; fi12++ {
                var v13 uint64 = 1
                v14 := GenS1{F0: 602}
                v15 := GenS1{F0: 316}
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
        }
        _ = fi8
    }
    var v16 int = 707
    fmt.Println(v7, v4, v6)
    if (42 == 617) {
        fmt.Println(v6, v6)
    }
    switch v16 {
    case 2147483647:
        v17 := GenS1{F0: 1000}
        _ = v17
    case -9223372036854775808:
        fmt.Println(v4, v7)
    case -962:
        if true {
            for fi18 := 0; fi18 < 8; fi18++ {
                v19 := GenS2{F0: -2147483648}
                _ = fi18
                _ = v19
            }
            if ("" == "\n\t") {
                for fi20 := 0; fi20 < 3; fi20++ {
                    if true {
                        var v21 uint32 = 42
                        _ = v21
                    } else {
                        v22 := GenS3{F0: (-1 != 42)}
                        v7 = (v7 / ((v7) | 1))
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 18; fi23++ {
                        var v24 uint16 = (v7 + v7)
                        var v25 uint16 = (v24 >> ((v4) & 63))
                        var v26 int32 = 0
                        switch "hello" {
                        case "":
                            delete(v5, (v16 << ((uint(v24)) & 63)))
                        }
                        _ = len(v1)
                        _ = fi23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    if len(v3) > 0 {
                        v3[0] = 2147483648
                    }
                    _ = fi20
                }
            }
            v6 -= 1.5
        }
    }
    switch ("hello" + "a") {
    case "a":
        if ((!((1 >= -747) || ((v16 ^ v16) != -97))) || (2 >= 0)) {
            v6 += (float64(1e10) / -2.25)
            var v27 string = ("foo bar" + "foo bar")
            _ = v27
        }
    case "\n\t":
        if (((((1e10 / -2.25) + float32(v6)) + float32((v6 + ((v6 / v6) / v6)))) / -2.25) == 3.14159) {
            for fi28 := 0; fi28 < 2; fi28++ {
                v7 += (v7 % ((42) | 1))
                if (0 <= 11) {
                    if true {
                        switch fi28 {
                        case 32768:
                            switch "hello" {
                            case "":
                                v29 := GenS4{F0: -128}
                                _ = v29
                            case "\n\t":
                                v2 = append(v2, uint32(v7))
                            case "hello":
                                v6 = 3.14159
                            }
                        case 127:
                            v5[(fi28 | (-(v16)))] = "hello"
                        }
                        var v30 int32 = 1
                        _ = v30
                    }
                    var v31 float64 = v6
                    _ = v31
                }
                v6 = (float64((((1.5 - float32(v6)) - (float32(-2.25) - float32(v6))) - 0.0)) / (v6 * (0.0 * 0.1)))
                _ = fi28
            }
            if true {
                if true {
                    switch ("hello" + "foo bar") {
                    case "hello":
                        switch "" {
                        case "\n\t":
                            fmt.Println(v4, v6, v6)
                        default:
                            v32 := GenS5{F0: v7, F1: 4294967295}
                            _ = v32
                        }
                    case "foo bar":
                        for fi33 := 0; fi33 < 1; fi33++ {
                            v34 := GenS2{F0: int32(v4)}
                            switch "" {
                            case "":
                                var v35 int16 = 42
                                _ = v35
                            case "foo bar":
                                var v36 float64 = (((v6 - v6) / (v6 + 1.5)) / float64(0.1))
                                _ = v36
                            case "a":
                                if true {
                                    var v37 uint = v4
                                    var v38 float64 = ((float64(1.5) - (1.5 / float64(-2.25))) + 0.1)
                                    var v39 int16 = -32768
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                } else {
                                    for fi40 := 0; fi40 < 4; fi40++ {
                                        v16 = v16
                                        if true {
                                            if (!(!true)) {
                                                var v41 float32 = (-2.25 + (float32((v6 * float64((0.0 - 0.0)))) + 1e10))
                                                _ = v41
                                            } else {
                                                var v42 uint64 = 4294967295
                                                switch (v16 / (((^(fi40))) | 1)) {
                                                case 128:
                                                    v43 := GenS1{F0: 32767}
                                                    _ = v43
                                                case -317:
                                                    var v44 uint8 = 255
                                                    _ = v44
                                                }
                                                _ = v42
                                            }
                                            var v45 uint16 = 65535
                                            v16 -= int(v4)
                                            _ = v45
                                        } else {
                                            v46 := GenS5{F0: 428, F1: 2}
                                            _ = v46
                                        }
                                        _ = fi40
                                    }
                                    if true {
                                        for fi47 := 0; fi47 < 10; fi47++ {
                                            switch v16 {
                                            case 127:
                                                delete(v1, v16)
                                            }
                                            var v48 uint = v4
                                            v1[fi33] = false
                                            _ = fi47
                                            _ = v48
                                        }
                                        switch fi33 {
                                        case 0:
                                            _ = len(v5)
                                        default:
                                            switch (fi33 % ((32767) | 1)) {
                                            case 128:
                                                v3 = append(v3, 561)
                                            case -9223372036854775808:
                                                if false {
                                                    v49 := GenS4{F0: 127}
                                                    v50 := GenS2{F0: -1}
                                                    _ = v49
                                                    _ = v50
                                                }
                                            case -1:
                                                v4 += 4294967295
                                            default:
                                                for fi51 := 0; fi51 < 2; fi51++ {
                                                    switch fi51 {
                                                    case -2147483648:
                                                        v52 := GenS6{F0: 1, F1: 97}
                                                        _ = v52
                                                    case 9223372036854775807:
                                                        switch ("" + (("a" + ("foo bar" + ((("foo bar" + "\n\t") + "\n\t") + "hello"))) + ("" + "a"))) {
                                                        case "":
                                                            v53 := GenS7{F0: 18446744073709551615}
                                                            _ = v53
                                                        case "hello":
                                                            if false {
                                                                v6 = v6
                                                                if (2 < 464) {
                                                                    delete(v1, (v16 % ((fi33) | 1)))
                                                                    var v54 bool = true
                                                                    v55 := v1[v16]
                                                                    _ = v54
                                                                    _ = v55
                                                                } else {
                                                                    v16 += (v16 | (fi33 / ((v16) | 1)))
                                                                    v56 := GenS6{F0: 42, F1: 0}
                                                                    _ = v56
                                                                }
                                                            } else {
                                                                switch ("" + "foo bar") {
                                                                case "\n\t":
                                                                    var v57 int64 = -835
                                                                    var v58 uint = 2
                                                                    _ = v57
                                                                    _ = v58
                                                                case "foo bar":
                                                                    var v59 float64 = (v6 * 3.14159)
                                                                    _ = v59
                                                                case "":
                                                                    switch fi51 {
                                                                    case 127:
                                                                        if false {
                                                                            fmt.Println(v7, fi51)
                                                                        } else {
                                                                            var v60 uint = (v4 * (v4 | v4))
                                                                            fmt.Println(v7, v6)
                                                                            _ = v60
                                                                        }
                                                                    case 208:
                                                                        var v61 uint16 = (v7 << ((v4) & 63))
                                                                        _ = v61
                                                                    case -255:
                                                                        var v62 uint8 = 0
                                                                        _ = v62
                                                                    default:
                                                                        v7 = 0
                                                                    }
                                                                }
                                                                fmt.Println(v6)
                                                                fmt.Println(fi33)
                                                            }
                                                        case "foo bar":
                                                            if true {
                                                                fmt.Println(v6)
                                                            }
                                                        default:
                                                            if false {
                                                                switch ("" + "") {
                                                                case "a":
                                                                    delete(v1, 0)
                                                                case "hello":
                                                                    switch -431 {
                                                                    case 804:
                                                                        v63 := v5[v16]
                                                                        _ = v63
                                                                    default:
                                                                        if len(v3) > 0 {
                                                                            v3[0] = 226
                                                                        }
                                                                    }
                                                                default:
                                                                    var v64 int8 = 0
                                                                    _ = v64
                                                                }
                                                            }
                                                        }
                                                    }
                                                    _ = fi51
                                                }
                                            }
                                        }
                                    }
                                }
                            default:
                                var v65 uint64 = 18446744073709551615
                                _ = v65
                            }
                            _ = fi33
                            _ = v34
                        }
                    default:
                        v16 += (v16 & 42)
                    }
                } else {
                    fmt.Println(v4)
                }
                v5[(v16 << ((uint(v7)) & 63))] = "a"
            }
        }
    case "foo bar":
        switch ("" + "foo bar") {
        case "\n\t":
            fmt.Println(v16, v6, v16)
        }
    default:
        if (true || (false || (float64(0.0) == v6))) {
            v66 := GenS7{F0: 0}
            var v67 int8 = 1
            for fi68 := 0; fi68 < 15; fi68++ {
                for fi69 := 0; fi69 < 7; fi69++ {
                    if (42 >= 2) {
                        _ = len(v3)
                        _ = len(v3)
                        if ((v6 - ((v6 / v6) - (v6 / v6))) < v6) {
                            var v70 int32 = 722
                            _ = v70
                        } else {
                            fmt.Println(v6, v67, v16)
                            v6 += (1e10 - v6)
                            v6 = (-2.25 + -2.25)
                        }
                    } else {
                        var v71 uint8 = 139
                        var v72 float32 = 1e10
                        v73 := GenS8{F0: -347, F1: 2, F2: -2147483648}
                        v74 := GenS9{F0: v16}
                        _ = v71
                        _ = v72
                        _ = v73
                        _ = v74
                    }
                    var v75 int32 = 0
                    var v76 int8 = v67
                    if false {
                        v67 = (v76 + v76)
                        v77 := GenS10{F0: 32767, F1: (v7 ^ 1)}
                        _ = v77
                    }
                    _ = fi69
                    _ = v75
                    _ = v76
                }
                v78 := GenS11{F0: v67}
                _ = fi68
                _ = v78
            }
            _ = v66
            _ = v67
        } else {
            var v79 float32 = -2.25
            v6 = v6
            _ = v79
        }
    }
    switch v16 {
    case -9223372036854775808:
        for fi80 := 0; fi80 < 16; fi80++ {
            if ((v16 ^ (fi80 % (((-(fi80))) | 1))) > fi80) {
                v81 := GenS5{F0: (v7 / ((v7) | 1)), F1: 797}
                var v82 bool = true
                v6 += v6
                _ = v81
                _ = v82
            } else {
                v3 = append(v3, 913)
                v83 := GenS12{F0: (v4 - 42), F1: 159, F2: -2147483648}
                switch (fi80 | (v16 % ((0) | 1))) {
                case 32768:
                    v6 = v6
                case 0:
                    if (!(!((0 > 2) && false))) {
                        _ = len(v3)
                        if false {
                            for fi84 := 0; fi84 < 17; fi84++ {
                                if (!(!(false || (!true)))) {
                                    var v85 bool = (!false)
                                    _ = v85
                                }
                                switch "" {
                                case "":
                                    fmt.Println(v6, v16, v6)
                                case "hello":
                                    v4 += 42
                                    v6 += v6
                                    fmt.Println(fi84)
                                }
                                _ = fi84
                            }
                            var v86 uint16 = (v7 >> ((v4) & 63))
                            for fi87 := 0; fi87 < 19; fi87++ {
                                delete(v5, fi80)
                                _ = fi87
                            }
                            _ = v86
                        }
                    }
                }
                _ = v83
            }
            for fi88 := 0; fi88 < 18; fi88++ {
                var v89 int8 = 127
                v89 = 109
                switch "\n\t" {
                case "a":
                    var v90 int16 = int16(v4)
                    var v91 uint64 = 42
                    v7 = v7
                    _ = v90
                    _ = v91
                case "":
                    v92 := GenS10{F0: 1, F1: 32767}
                    _ = v92
                }
                _ = fi88
                _ = v89
            }
            _ = fi80
        }
    }
    var v93 uint32 = 881
    v94 := GenS13{F0: -61}
    var v95 uint8 = 1
    v96 := GenS1{F0: 1000}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v93:", v93)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v93
    _ = v94
    _ = v95
    _ = v96
}
