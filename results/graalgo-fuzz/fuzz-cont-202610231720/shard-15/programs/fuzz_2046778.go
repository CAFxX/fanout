package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int; F1 uint }
type GenS8 struct { F0 int16; F1 uint64 }
func genFn0(p0 int, p1 uint32) uint64 {
    var v6 int32 = 0
    switch (p0 / ((-2147483648) | 1)) {
    case 128:
        var v7 uint64 = 0
        _ = v7
    case 188:
        v6 = 1
    default:
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 int32 = v6
            switch "a" {
            case "hello":
                fmt.Println(fi8, fi8)
            }
            v10 := []int{4294967295}
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return 509
}
func genFn1(p0 string, p1 string) int {
    var v11 int64 = 42
    var v12 bool = false
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return -1
}

func main() {
    var v1 uint = 42
    v2 := []string{("\n\t" + "foo bar")}
    v3 := make(map[uint32]string)
    v4 := make(map[int]bool)
    var v5 int8 = 127
    switch 32767 {
    case 1:
        if true {
            v1 = v1
        }
    default:
        v13 := v4[-828]
        _ = v13
    }
    v14 := GenS1{F0: 0}
    if (!(32768 > -185)) {
        if true {
            var v15 bool = (!((0 <= 4294967295) || (1000 > 199)))
            var v16 float64 = 0.0
            fmt.Println(v1, v1, v5)
            _ = v15
            _ = v16
        }
        for fi17 := 0; fi17 < 7; fi17++ {
            v18 := GenS1{F0: 1}
            v1 = (v1 ^ v1)
            _ = fi17
            _ = v18
        }
        switch 4294967295 {
        case -920:
            delete(v3, 1)
        case 1:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 418:
            v19 := GenS2{F0: 390, F1: 128}
            _ = v19
        default:
            v5 += v5
        }
    } else {
        if ((v1 - v1) <= (v1 * (v1 * v1))) {
            fmt.Println(v1)
            v1 = (v1 << (((v1 % (((v1 ^ (v1 % (((^(v1))) | 1)))) | 1))) & 63))
        } else {
            var v20 uint = 65535
            v2 = append(v2, ("a" + ("" + ("a" + "\n\t"))))
            _ = v20
        }
        var v21 int = 605
        v5 = -128
        _ = v21
    }
    v22 := GenS1{F0: -32768}
    for fi23 := 0; fi23 < 1; fi23++ {
        switch ((("hello" + "hello") + (("" + ("a" + (("" + "") + "a"))) + "")) + "foo bar") {
        case "a":
            v1 -= v1
        case "foo bar":
            var v24 uint8 = 128
            _ = v24
        case "":
            switch "foo bar" {
            case "\n\t":
                v25 := GenS1{F0: 104}
                _ = v25
            }
        }
        switch "a" {
        case "foo bar":
            switch (fi23 ^ fi23) {
            case 2147483647:
                if false {
                    v1 = (v1 >> ((v1) & 63))
                    var v26 bool = (false && true)
                    _ = v26
                } else {
                    for fi27 := 0; fi27 < 17; fi27++ {
                        v1 = 577
                        var v28 uint32 = 948
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if false {
                            var v29 int = 616
                            if (2 != (v28 >> (((v1 << ((0) & 63))) & 63))) {
                                if false {
                                    switch "hello" {
                                    case "":
                                        v5 = -128
                                    }
                                    switch "\n\t" {
                                    case "hello":
                                        if (false && true) {
                                            switch v29 {
                                            case 16:
                                                var v30 uint64 = 0
                                                _ = v30
                                            default:
                                                var v31 uint64 = 18446744073709551615
                                                fmt.Println(v29, v5, fi23)
                                                _ = v31
                                            }
                                            fmt.Println(fi27)
                                            if len(v2) > 0 {
                                                v2[0] = "foo bar"
                                            }
                                        } else {
                                            v1 = v1
                                            v32 := GenS1{F0: 994}
                                            delete(v3, (v28 >> (((v1 / ((v1) | 1))) & 63)))
                                            _ = v32
                                        }
                                    default:
                                        fmt.Println(v29)
                                    }
                                } else {
                                    switch v29 {
                                    case 9223372036854775807:
                                        var v33 float32 = ((float32(3.14159) + 1e10) / 1.0)
                                        _ = v33
                                    case 127:
                                        var v34 float64 = (0.0 + (-2.25 + (0.0 - 0.1)))
                                        _ = v34
                                    case 884:
                                        fmt.Println(fi23)
                                        if false {
                                            var v35 uint16 = uint16(v1)
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            _ = v35
                                        } else {
                                            switch genFn1("a", ("\n\t" + "\n\t")) {
                                            case -9223372036854775808:
                                                var v36 int8 = (v5 % (((v5 | (v5 % (((v5 ^ v5)) | 1)))) | 1))
                                                _ = v36
                                            default:
                                                v37 := GenS3{F0: 42, F1: 255, F2: v5}
                                                var v38 int16 = 1000
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                v5 = (v5 & (v5 ^ 0))
                                                _ = v37
                                                _ = v38
                                            }
                                        }
                                        if ((!(true && (207 > 457))) || true) {
                                            switch genFn1("\n\t", "a") {
                                            case -9223372036854775808:
                                                var v39 uint = (^(v1))
                                                _ = v39
                                            case 2147483647:
                                                switch fi23 {
                                                case -2147483648:
                                                    _ = len(v3)
                                                case 231:
                                                    var v40 float64 = (3.14159 + (0.0 + 3.14159))
                                                    _ = v40
                                                }
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v29 = 32768
                                            if ((!false) && true) {
                                                v5 = int8(v28)
                                            } else {
                                                v41 := GenS4{F0: 15, F1: 255}
                                                _ = v41
                                            }
                                            v5 = v5
                                            if (1000 == 32767) {
                                                v5 = v5
                                            } else {
                                                fmt.Println(fi23, fi23)
                                                var v42 bool = true
                                                if (4294967295 < 425) {
                                                    _ = len(v3)
                                                    switch (v29 & fi27) {
                                                    case 0:
                                                        var v43 float64 = (0.0 + 0.0)
                                                        _ = v43
                                                    case 42:
                                                        v42 = true
                                                    case -2147483648:
                                                        v29 = v29
                                                    }
                                                    var v44 int64 = 5
                                                    _ = v44
                                                } else {
                                                    _ = len(v4)
                                                    v45 := GenS5{F0: -32768}
                                                    _ = v45
                                                }
                                                _ = v42
                                            }
                                            var v46 uint = v1
                                            _ = v46
                                        } else {
                                            v47 := GenS4{F0: 42, F1: 1}
                                            v3[(v28 >> ((1) & 63))] = ""
                                            _ = v47
                                        }
                                    }
                                    if true {
                                        v1 -= 354
                                        var v48 int8 = 42
                                        v49 := GenS1{F0: 122}
                                        _ = v48
                                        _ = v49
                                    }
                                    fmt.Println(v28, v28, v1)
                                }
                            } else {
                                var v50 int8 = -68
                                switch "" {
                                case "\n\t":
                                    switch (^(fi27)) {
                                    case 0:
                                        v29 = (fi27 % ((fi27) | 1))
                                        var v51 string = ""
                                        _ = v51
                                    }
                                case "a":
                                    fmt.Println(v28, v28, fi27)
                                    fmt.Println(fi27, fi27, v1)
                                default:
                                    v28 += (^(v28))
                                }
                                if ((0.1 + -2.25) == 1e10) {
                                    switch "hello" {
                                    case "\n\t":
                                        var v52 uint = (^(v1))
                                        _ = v52
                                    case "foo bar":
                                        v50 = int8(v1)
                                    default:
                                        v50 = (v5 ^ (-(v5)))
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                } else {
                                    var v53 int16 = int16(v5)
                                    var v54 uint = 238
                                    var v55 uint = 42
                                    _ = v53
                                    _ = v54
                                    _ = v55
                                }
                                _ = v50
                            }
                            _ = v29
                        }
                        _ = fi27
                        _ = v28
                    }
                }
            case -601:
                v56 := GenS4{F0: -1, F1: 128}
                _ = v56
            default:
                switch "hello" {
                case "hello":
                    var v57 float64 = 0.1
                    _ = v57
                case "\n\t":
                    v58 := GenS3{F0: 1, F1: 65535, F2: (v5 + (-(v5)))}
                    _ = v58
                }
            }
        case "":
            switch fi23 {
            case -9223372036854775808:
                v1 = v1
            case -173:
                switch fi23 {
                case 9223372036854775807:
                    for fi59 := 0; fi59 < 10; fi59++ {
                        v60 := GenS6{F0: 128}
                        _ = fi59
                        _ = v60
                    }
                default:
                    v2 = append(v2, "\n\t")
                }
            default:
                switch (fi23 | (fi23 ^ fi23)) {
                case 42:
                    if (127 == 2) {
                        for fi61 := 0; fi61 < 14; fi61++ {
                            if len(v2) > 0 {
                                v2[0] = ("foo bar" + "")
                            }
                            if (("hello" + "hello") < "") {
                                switch "\n\t" {
                                case "\n\t":
                                    var v62 int16 = 1
                                    _ = v62
                                case "":
                                    var v63 int32 = 42
                                    _ = v63
                                case "a":
                                    fmt.Println(v5)
                                }
                            }
                            _ = fi61
                        }
                        v64 := GenS5{F0: -578}
                        var v65 uint8 = 146
                        _ = v64
                        _ = v65
                    }
                case 127:
                    var v66 int32 = -115
                    _ = v66
                default:
                    for fi67 := 0; fi67 < 14; fi67++ {
                        v1 += (v1 + (^(v1)))
                        switch "\n\t" {
                        case "":
                            if true {
                                v68 := GenS5{F0: -1}
                                _ = v68
                            } else {
                                v69 := v4[fi23]
                                _ = v69
                            }
                        }
                        switch "\n\t" {
                        case "foo bar":
                            v70 := GenS7{F0: 9223372036854775807, F1: (v1 << (((v1 + v1)) & 63))}
                            v1 += (v1 + (v1 << (((v1 - v1)) & 63)))
                            fmt.Println(v1, v5)
                            _ = v70
                        case "\n\t":
                            switch "" {
                            case "\n\t":
                                var v71 int = fi67
                                _ = len(v3)
                                v5 = v5
                                _ = v71
                            default:
                                var v72 uint16 = 572
                                _ = v72
                            }
                        case "":
                            switch (fi67 * fi67) {
                            case -549:
                                v5 = int8(fi67)
                            }
                        default:
                            v5 = -1
                        }
                        _ = fi67
                    }
                }
            }
        case "a":
            v1 -= v1
        default:
            fmt.Println(v5)
        }
        for fi73 := 0; fi73 < 15; fi73++ {
            var v74 int8 = (v5 << (((v1 ^ (^(v1)))) & 63))
            _ = fi73
            _ = v74
        }
        _ = fi23
    }
    v75 := GenS4{F0: 633, F1: 0}
    for fi76 := 0; fi76 < 1; fi76++ {
        v77 := GenS8{F0: -32768, F1: 1}
        _ = fi76
        _ = v77
    }
    v5 += (v5 * v5)
    var v78 uint = 185
    switch 32768 {
    case 159:
        var v79 uint16 = 65535
        _ = v79
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v75
    _ = v78
}
