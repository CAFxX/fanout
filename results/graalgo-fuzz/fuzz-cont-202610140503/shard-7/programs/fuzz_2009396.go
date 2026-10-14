package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 int64; F1 int; F2 int8 }
type GenS6 struct { F0 bool; F1 int }
type GenS7 struct { F0 int32; F1 int8 }
type GenS8 struct { F0 bool; F1 string }

func main() {
    var v1 uint = 338
    var v2 uint16 = 255
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]string)
    var v5 uint64 = uint64(v1)
    var v6 uint32 = 888
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            switch (("hello" + "foo bar") + "a") {
            case "a":
                v9 := GenS1{F0: 1, F1: v2}
                _ = v9
            }
            switch fi8 {
            case 9223372036854775807:
                switch "a" {
                case "hello":
                    v10 := []int{}
                    _ = v10
                case "a":
                    var v11 int32 = -1
                    _ = v11
                case "foo bar":
                    fmt.Println(fi8)
                }
            default:
                delete(v4, ("foo bar" + ("" + "")))
            }
            _ = fi8
        }
        v12 := GenS1{F0: -32768, F1: uint16(v1)}
        var v13 int64 = -9223372036854775808
        _ = fi7
        _ = v12
        _ = v13
    }
    if (!false) {
        _ = len(v4)
        for fi14 := 0; fi14 < 12; fi14++ {
            switch ("" + ("a" + ("\n\t" + "a"))) {
            case "foo bar":
                fmt.Println(v1, v5, v5)
            case "":
                v6 = 365
            case "a":
                switch (fi14 & int(v6)) {
                case 0:
                    for fi15 := 0; fi15 < 11; fi15++ {
                        v16 := GenS2{F0: 42}
                        v17 := GenS1{F0: 42, F1: 65535}
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                }
            }
            var v18 uint16 = v2
            _ = fi14
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        if false {
            if false {
                for fi20 := 0; fi20 < 3; fi20++ {
                    var v21 uint64 = (v5 - (v5 >> (((v1 << (((^(v1))) & 63))) & 63)))
                    var v22 int8 = 100
                    var v23 int8 = (v22 - 110)
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
                v6 = 44
            } else {
                for fi24 := 0; fi24 < 16; fi24++ {
                    fmt.Println(v2, fi19, fi19)
                    v25 := []int{}
                    _ = fi24
                    _ = v25
                }
                var v26 int16 = 42
                _ = v26
            }
        }
        for fi27 := 0; fi27 < 7; fi27++ {
            var v28 int64 = 889
            var v29 uint32 = v6
            _ = fi27
            _ = v28
            _ = v29
        }
        var v30 int16 = -253
        _ = fi19
        _ = v30
    }
    var v31 uint8 = 1
    fmt.Println(v1)
    var v32 float32 = -2.25
    fmt.Println(v2, v1, v5)
    fmt.Println(v1, v6, v32)
    switch 32768 {
    case -9223372036854775808:
        v2 += v2
    case -475:
        var v33 int = int(v2)
        _ = v33
    default:
        for fi34 := 0; fi34 < 5; fi34++ {
            for fi35 := 0; fi35 < 4; fi35++ {
                v5 = uint64(v1)
                switch "\n\t" {
                case "\n\t":
                    if true {
                        var v36 uint = uint(v5)
                        switch fi35 {
                        case 2147483647:
                            v36 -= (v1 * (^(v1)))
                            var v37 int64 = 1
                            v32 = float32(((float64(1.5) / 0.1) - 1e10))
                            _ = v37
                        default:
                            var v38 int16 = -1
                            _ = v38
                        }
                        _ = v36
                    } else {
                        if (false || false) {
                            v39 := GenS1{F0: -1, F1: (v2 ^ v2)}
                            var v40 int = 0
                            _ = v39
                            _ = v40
                        } else {
                            v5 -= v5
                            fmt.Println(v32)
                        }
                        var v41 string = ""
                        switch (fi34 >> ((v1) & 63)) {
                        case -980:
                            v42 := GenS2{F0: uint(v2)}
                            _ = v42
                        case -2147483648:
                            v43 := GenS1{F0: 1000, F1: v2}
                            v31 += 42
                            fmt.Println(v41, v6, fi34)
                            _ = v43
                        }
                        v44 := GenS3{F0: 32767, F1: (^(v5))}
                        _ = v41
                        _ = v44
                    }
                default:
                    var v45 uint8 = (v31 % ((v31) | 1))
                    _ = v45
                }
                v4[(("hello" + (("\n\t" + ((("" + (("foo bar" + "\n\t") + ("\n\t" + ""))) + ("\n\t" + "hello")) + "hello")) + "")) + "\n\t")] = "hello"
                switch ("foo bar" + (((("\n\t" + "\n\t") + "") + "a") + "\n\t")) {
                case "hello":
                    v46 := GenS4{F0: true, F1: 193}
                    _ = v46
                case "a":
                    v2 -= 255
                default:
                    switch ("\n\t" + "\n\t") {
                    case "a":
                        switch fi35 {
                        case 2147483647:
                            if true {
                                v47 := GenS3{F0: 268, F1: uint64(v2)}
                                var v48 uint64 = 646
                                var v49 int32 = 964
                                if (-317 <= 335) {
                                    switch (fi35 >> ((v1) & 63)) {
                                    case -2147483648:
                                        var v50 int64 = -1
                                        _ = v50
                                    case -862:
                                        v2 = (v2 ^ (v2 % (((v2 + v2)) | 1)))
                                    case -9223372036854775808:
                                        switch (fi34 ^ fi35) {
                                        case 32768:
                                            v2 += v2
                                        case 2147483647:
                                            if ((fi35 * (fi35 << (((^(v1))) & 63))) != fi34) {
                                                var v51 float32 = float32(-2.25)
                                                if (true || true) {
                                                    switch ("hello" + "\n\t") {
                                                    case "a":
                                                        var v52 uint64 = v48
                                                        _ = v52
                                                    case "foo bar":
                                                        v2 = (v2 - v2)
                                                    }
                                                    v31 += 229
                                                    var v53 float64 = (((0.1 + (-2.25 - ((float64(v51) - 1.5) * 3.14159))) * -2.25) + (float64(((v51 / 0.1) / -2.25)) * 0.0))
                                                    _ = v53
                                                }
                                                fmt.Println(v2, v5)
                                                v48 -= (v48 % ((v48) | 1))
                                                v54 := []int{(fi34 | fi35)}
                                                _ = v51
                                                _ = v54
                                            } else {
                                                if (!(!false)) {
                                                    var v55 uint64 = (v48 % (((v5 >> ((65535) & 63))) | 1))
                                                    switch fi35 {
                                                    case 9223372036854775807:
                                                        if true {
                                                            fmt.Println(v32)
                                                            if false {
                                                                if (v2 >= 65535) {
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        v31 = v31
                                                                    default:
                                                                        v4[("foo bar" + "foo bar")] = "a"
                                                                    }
                                                                    v56 := v3[v6]
                                                                    fmt.Println(v2)
                                                                    _ = v56
                                                                } else {
                                                                    if false {
                                                                        v55 -= v5
                                                                        var v57 int64 = -9223372036854775808
                                                                        v58 := GenS5{F0: (v57 / ((v57) | 1)), F1: fi35, F2: 42}
                                                                        _ = v57
                                                                        _ = v58
                                                                    } else {
                                                                        fmt.Println(fi34, v1)
                                                                    }
                                                                    if (!(!true)) {
                                                                        v6 = (^(v6))
                                                                        v59 := GenS3{F0: -637, F1: v5}
                                                                        _ = v59
                                                                    } else {
                                                                        fmt.Println(v1, v31)
                                                                    }
                                                                }
                                                                var v60 int16 = -747
                                                                fmt.Println(v60, v49)
                                                                _ = v60
                                                            }
                                                            v5 += v48
                                                        }
                                                    default:
                                                        v61 := GenS5{F0: -1, F1: fi35, F2: -128}
                                                        _ = v61
                                                    }
                                                    var v62 int64 = 42
                                                    _ = v55
                                                    _ = v62
                                                } else {
                                                    var v63 int = (fi35 / (((fi35 % (((-(fi35))) | 1))) | 1))
                                                    v1 = v1
                                                    _ = v63
                                                }
                                            }
                                        default:
                                            switch fi34 {
                                            case -2147483648:
                                                v64 := v4[""]
                                                _ = v64
                                            }
                                        }
                                    }
                                } else {
                                    v48 += (v48 << ((0) & 63))
                                    if (float32((0.0 / -2.25)) <= v32) {
                                        v1 -= v1
                                    } else {
                                        if true {
                                            v65 := GenS2{F0: (v1 % ((uint(v5)) | 1))}
                                            _ = v65
                                        }
                                        v66 := v3[v6]
                                        fmt.Println(fi34, fi34)
                                        _ = v66
                                    }
                                    v4["foo bar"] = "foo bar"
                                }
                                var v67 uint64 = (v5 & 2)
                                _ = v47
                                _ = v48
                                _ = v49
                                _ = v67
                            } else {
                                var v68 uint64 = v5
                                _ = v68
                            }
                        }
                    default:
                        switch fi35 {
                        case -1:
                            switch "a" {
                            case "foo bar":
                                if ((v5 & 2) < (v5 % (((v5 ^ v5)) | 1))) {
                                    fmt.Println(fi35, v32)
                                }
                            }
                        case 614:
                            if false {
                                var v69 float64 = 0.1
                                switch "" {
                                case "foo bar":
                                    switch "foo bar" {
                                    case "hello":
                                        fmt.Println(v32, v6)
                                    default:
                                        v70 := v4[("hello" + ("\n\t" + ""))]
                                        _ = v70
                                    }
                                default:
                                    v5 += v5
                                }
                                v71 := GenS5{F0: -9223372036854775808, F1: fi34, F2: 1}
                                _ = v69
                                _ = v71
                            } else {
                                switch 512 {
                                case 2147483647:
                                    fmt.Println(v5, fi35, v32)
                                case 9223372036854775807:
                                    if (!false) {
                                        v72 := []int{-510}
                                        _ = v72
                                    }
                                    var v73 string = ""
                                    v74 := GenS6{F0: ((1 == (v31 >> ((18446744073709551615) & 63))) || (!true)), F1: 9223372036854775807}
                                    _ = v73
                                    _ = v74
                                case 593:
                                    if true {
                                        v2 = (^(v2))
                                    } else {
                                        var v75 int64 = 42
                                        switch fi34 {
                                        case 4294967295:
                                            fmt.Println(v5, v32, v75)
                                        default:
                                            var v76 uint32 = 42
                                            _ = v76
                                        }
                                        v32 = v32
                                        _ = v75
                                    }
                                }
                            }
                        case -9223372036854775808:
                            v77 := GenS7{F0: -2147483648, F1: -128}
                            _ = v77
                        default:
                            var v78 uint16 = v2
                            _ = v78
                        }
                    }
                }
                _ = fi35
            }
            v79 := GenS2{F0: v1}
            var v80 int64 = -1
            _ = fi34
            _ = v79
            _ = v80
        }
    }
    v1 = 356
    v81 := GenS4{F0: true, F1: v31}
    var v82 float32 = ((-2.25 * 1.5) / 1.0)
    delete(v3, (v6 ^ (^(v6))))
    var v83 int32 = 42
    v84 := GenS8{F0: true, F1: "\n\t"}
    v85 := GenS2{F0: 776}
    v1 -= (v1 & 18446744073709551615)
    v31 += (v31 ^ v31)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v81
    _ = v82
    _ = v83
    _ = v84
    _ = v85
}
