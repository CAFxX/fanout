package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int8; F2 uint8 }
type GenS3 struct { F0 string; F1 int8; F2 uint64 }
type GenS4 struct { F0 int; F1 int32; F2 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint32; F1 uint; F2 int }
type GenS9 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS10 struct { F0 int64; F1 int8; F2 int64 }

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{2147483647, 320}
    var v3 uint = 4294967295
    var v4 int8 = -1
    if ((!false) && (uint8(v3) < 42)) {
        fmt.Println(v4, v4)
        switch -2147483648 {
        case 829:
            v3 = 141
        case -627:
            _ = len(v2)
        default:
            fmt.Println(v3)
        }
        var v5 float64 = (0.1 / float64(float32(((0.1 / (3.14159 - ((float64((-2.25 / (3.14159 + 3.14159))) + 0.1) - ((((3.14159 * (-2.25 - 1.5)) - (1.5 / 0.1)) - 0.1) * ((3.14159 + (1.5 * 3.14159)) + 1e10))))) / -2.25))))
        _ = v5
    } else {
        v4 = v4
        for fi6 := 0; fi6 < 15; fi6++ {
            if false {
                var v7 uint32 = 2147483647
                _ = v7
            }
            for fi8 := 0; fi8 < 15; fi8++ {
                v4 = -1
                _ = fi8
            }
            _ = fi6
        }
        v2 = append(v2, 2147483648)
    }
    if (!(196 < 128)) {
        for fi9 := 0; fi9 < 12; fi9++ {
            switch fi9 {
            case 1:
                v4 = (v4 * v4)
            case 9223372036854775807:
                v10 := GenS1{F0: -2147483648}
                _ = v10
            default:
                if false {
                    v3 += v3
                }
            }
            v11 := GenS1{F0: fi9}
            for fi12 := 0; fi12 < 5; fi12++ {
                switch 0 {
                case 1:
                    switch "a" {
                    case "":
                        fmt.Println(v4, v3, v3)
                    case "\n\t":
                        v1[("a" + "foo bar")] = (fi12 ^ fi9)
                    case "a":
                        if true {
                            if false {
                                v13 := GenS1{F0: fi12}
                                _ = v13
                            } else {
                                v4 -= 0
                                fmt.Println(fi12)
                            }
                            var v14 float64 = 0.0
                            _ = len(v2)
                            _ = v14
                        }
                    default:
                        var v15 uint8 = 135
                        _ = v15
                    }
                default:
                    fmt.Println(fi12, v4)
                }
                v4 = -1
                _ = fi12
            }
            _ = fi9
            _ = v11
        }
    } else {
        for fi16 := 0; fi16 < 12; fi16++ {
            switch fi16 {
            case -683:
                var v17 uint64 = 54
                _ = v17
            case 748:
                var v18 int16 = int16(v4)
                _ = v18
            }
            if false {
                if true {
                    fmt.Println(fi16, fi16)
                    fmt.Println(fi16, fi16, v4)
                    var v19 int32 = -1
                    _ = v19
                }
                switch "hello" {
                case "foo bar":
                    for fi20 := 0; fi20 < 3; fi20++ {
                        v4 += v4
                        switch ("foo bar" + "\n\t") {
                        case "a":
                            if (false && false) {
                                var v21 string = "hello"
                                switch fi20 {
                                case 907:
                                    var v22 int = fi20
                                    _ = v22
                                }
                                fmt.Println(v4, fi20, fi20)
                                _ = v21
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        default:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        v4 -= (^(v4))
                        var v23 uint64 = 4294967295
                        v3 = v3
                        _ = fi20
                        _ = v23
                    }
                case "a":
                    v3 = 42
                default:
                    for fi24 := 0; fi24 < 17; fi24++ {
                        fmt.Println(fi16, v4, v3)
                        var v25 uint16 = 466
                        _ = fi24
                        _ = v25
                    }
                }
                var v26 uint64 = 184
                _ = v26
            } else {
                fmt.Println(v4, v4, v4)
                var v27 float32 = 0.1
                fmt.Println(v3, v4, fi16)
                _ = v27
            }
            if (1 != 255) {
                v1["foo bar"] = fi16
            }
            _ = fi16
        }
    }
    _ = len(v2)
    delete(v1, "foo bar")
    v4 = (v4 - v4)
    if false {
        v28 := GenS1{F0: 9223372036854775807}
        _ = v28
    } else {
        _ = len(v2)
    }
    switch 0 {
    case 32768:
        var v29 uint = (v3 >> ((v3) & 63))
        _ = v29
    }
    delete(v1, ("a" + "a"))
    v3 = 42
    switch "a" {
    case "\n\t":
        fmt.Println(v3, v4)
    case "foo bar":
        if (4294967295 > 1) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi30 := 0; fi30 < 6; fi30++ {
                var v31 bool = true
                _ = fi30
                _ = v31
            }
        }
    case "hello":
        var v32 uint = v3
        _ = v32
    }
    var v33 float32 = 1.5
    if (!(-323 >= 1)) {
        if false {
            fmt.Println(v4, v33)
            fmt.Println(v4, v3)
            var v34 uint16 = 638
            _ = v34
        } else {
            if false {
                var v35 uint16 = 229
                _ = v35
            } else {
                var v36 int8 = 1
                var v37 float32 = v33
                for fi38 := 0; fi38 < 9; fi38++ {
                    if false {
                        v33 = float32(0.0)
                        v39 := GenS2{F0: (v3 * 273), F1: (v36 + (v36 | (v36 * v36))), F2: 42}
                        _ = v39
                    }
                    v36 = 42
                    fmt.Println(v4, v37)
                    _ = fi38
                }
                _ = v36
                _ = v37
            }
            v4 += v4
            v40 := GenS1{F0: 2147483647}
            _ = v40
        }
        if true {
            if ((v3 < (v3 + 255)) && (true && (0 >= 2))) {
                var v41 int32 = -2147483648
                switch 128 {
                case 32768:
                    if true {
                        switch 1 {
                        case 1:
                            var v42 uint16 = 0
                            _ = v42
                        case -2147483648:
                            _ = len(v2)
                        case 127:
                            v43 := GenS3{F0: (("" + "") + ("foo bar" + "foo bar")), F1: v4, F2: 2}
                            _ = v43
                        default:
                            if len(v2) > 0 {
                                v2[0] = 513
                            }
                        }
                    }
                case -895:
                    var v44 uint32 = 11
                    _ = v44
                }
                _ = v41
            }
            v33 += (1e10 - (float32(1e10) * ((v33 + v33) / (0.1 / v33))))
            for fi45 := 0; fi45 < 6; fi45++ {
                if true {
                    var v46 int32 = -56
                    v3 = (v3 | v3)
                    for fi47 := 0; fi47 < 13; fi47++ {
                        fmt.Println(fi47)
                        _ = fi47
                    }
                    _ = v46
                } else {
                    if true {
                        v48 := GenS4{F0: fi45, F1: -1, F2: 128}
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if false {
                            v2 = append(v2, 802)
                            if (float64((v33 - float32((1e10 - 0.0)))) >= ((1.5 - 1.5) / 0.1)) {
                                _ = len(v2)
                                v49 := GenS1{F0: 9223372036854775807}
                                if (1 < -808) {
                                    v4 = 100
                                    v33 += v33
                                    var v50 uint32 = 2147483647
                                    _ = v50
                                }
                                _ = v49
                            } else {
                                v3 = v3
                                switch fi45 {
                                case 9223372036854775807:
                                    if true {
                                        v33 = (0.0 * (v33 * v33))
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    } else {
                                        v4 -= (v4 - v4)
                                        for fi51 := 0; fi51 < 13; fi51++ {
                                            v33 = float32((float64(v33) / 1.5))
                                            v33 = ((((v33 + 3.14159) * ((v33 * (v33 - 3.14159)) + 0.1)) * ((v33 - v33) + v33)) * 0.0)
                                            switch fi51 {
                                            case -539:
                                                switch "" {
                                                case "\n\t":
                                                    var v52 int32 = 0
                                                    _ = v52
                                                case "":
                                                    if (true && false) {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        v33 = ((3.14159 + v33) + (v33 + (0.1 + (v33 - v33))))
                                                        v53 := v1[""]
                                                        _ = v53
                                                    } else {
                                                        fmt.Println(fi45)
                                                    }
                                                }
                                            case 378:
                                                v4 -= v4
                                            }
                                            var v54 int = fi51
                                            _ = fi51
                                            _ = v54
                                        }
                                    }
                                case -428:
                                    v4 = (^(v4))
                                }
                            }
                        } else {
                            v33 += v33
                        }
                        _ = v48
                    } else {
                        if (9223372036854775807 > -1) {
                            fmt.Println(v3)
                        } else {
                            if (false && (v4 > (v4 ^ v4))) {
                                fmt.Println(v4, v4, v4)
                            } else {
                                v2 = append(v2, 2)
                            }
                            if (true && false) {
                                if ((^(v3)) >= (v3 & v3)) {
                                    v33 += v33
                                    v3 -= v3
                                    if false {
                                        _ = len(v2)
                                        v4 = 0
                                        delete(v1, "hello")
                                    }
                                } else {
                                    switch -2147483648 {
                                    case 32768:
                                        if true {
                                            v55 := GenS5{F0: 65535}
                                            var v56 float64 = (3.14159 / 1e10)
                                            _ = v55
                                            _ = v56
                                        } else {
                                            if false {
                                                for fi57 := 0; fi57 < 14; fi57++ {
                                                    v3 -= 1
                                                    _ = fi57
                                                }
                                            } else {
                                                fmt.Println(v3, v33, v4)
                                            }
                                            v58 := GenS6{F0: 200}
                                            var v59 int32 = -557
                                            _ = v58
                                            _ = v59
                                        }
                                    default:
                                        var v60 uint32 = 2147483648
                                        _ = v60
                                    }
                                    for fi61 := 0; fi61 < 14; fi61++ {
                                        if (((1.5 - (float32(0.1) - float32((float64(0.1) / 1.0)))) >= v33) || ((^(v3)) < 255)) {
                                            v33 += (0.0 + 0.1)
                                            var v62 int64 = 9223372036854775807
                                            _ = v62
                                        } else {
                                            var v63 uint64 = 18446744073709551615
                                            var v64 bool = (v4 >= (v4 + v4))
                                            _ = v63
                                            _ = v64
                                        }
                                        _ = fi61
                                    }
                                }
                            } else {
                                var v65 int8 = v4
                                _ = v65
                            }
                            v3 -= (v3 + v3)
                        }
                    }
                    v66 := GenS7{F0: "hello"}
                    _ = v66
                }
                _ = fi45
            }
        } else {
            switch "hello" {
            case "hello":
                switch ("hello" + "foo bar") {
                case "":
                    var v67 uint8 = 2
                    _ = v67
                case "foo bar":
                    if true {
                        if len(v2) > 0 {
                            v2[0] = 65535
                        }
                        if (4294967295 > v3) {
                            var v68 int = -1
                            _ = v68
                        } else {
                            if false {
                                var v69 uint32 = 378
                                _ = v69
                            } else {
                                v70 := v1["a"]
                                var v71 uint8 = 207
                                for fi72 := 0; fi72 < 4; fi72++ {
                                    if true {
                                        switch v70 {
                                        case 379:
                                            v4 -= -12
                                        case -1:
                                            var v73 uint8 = (v71 | 0)
                                            _ = v73
                                        case 1:
                                            v74 := GenS4{F0: (fi72 | v70), F1: -954, F2: 127}
                                            _ = v74
                                        default:
                                            var v75 float64 = (((0.1 * (-2.25 + (3.14159 - (0.0 * 0.1)))) * -2.25) * 3.14159)
                                            _ = v75
                                        }
                                        var v76 float64 = 1.5
                                        _ = v76
                                    } else {
                                        if len(v2) > 0 {
                                            v2[0] = 65535
                                        }
                                        var v77 uint64 = 18446744073709551615
                                        if (int32(v71) > 1) {
                                            v70 = (v70 | (fi72 + (fi72 ^ (v70 % ((v70) | 1)))))
                                            v3 = 428
                                        } else {
                                            v78 := GenS1{F0: v70}
                                            _ = v78
                                        }
                                        _ = v77
                                    }
                                    switch "" {
                                    case "\n\t":
                                        for fi79 := 0; fi79 < 19; fi79++ {
                                            var v80 string = (("foo bar" + "a") + ("" + "a"))
                                            _ = fi79
                                            _ = v80
                                        }
                                    case "":
                                        v81 := GenS8{F0: 2147483648, F1: (^(v3)), F2: (v70 << ((520) & 63))}
                                        _ = v81
                                    }
                                    for fi82 := 0; fi82 < 2; fi82++ {
                                        v4 = 42
                                        switch "a" {
                                        case "a":
                                            var v83 int32 = -2147483648
                                            _ = v83
                                        case "\n\t":
                                            v4 = (v4 & (v4 % (((v4 + (v4 + (v4 * (-(v4)))))) | 1)))
                                        default:
                                            _ = len(v1)
                                        }
                                        _ = fi82
                                    }
                                    _ = fi72
                                }
                                _ = v70
                                _ = v71
                            }
                            if len(v2) > 0 {
                                v2[0] = 532
                            }
                            fmt.Println(v33)
                        }
                        _ = len(v1)
                    } else {
                        v3 = 18446744073709551615
                        v4 -= (v4 & v4)
                    }
                case "hello":
                    v2 = append(v2, 2)
                default:
                    v33 += v33
                }
            default:
                var v84 int8 = v4
                _ = v84
            }
        }
    } else {
        fmt.Println(v33)
        switch "\n\t" {
        case "\n\t":
            v33 += v33
        case "":
            if false {
                var v85 bool = false
                _ = v85
            }
        }
    }
    if false {
        fmt.Println(v4, v33)
    } else {
        if false {
            for fi86 := 0; fi86 < 18; fi86++ {
                if true {
                    switch "\n\t" {
                    case "a":
                        var v87 int32 = -2147483648
                        _ = v87
                    default:
                        for fi88 := 0; fi88 < 5; fi88++ {
                            fmt.Println(fi88, v3, v33)
                            fmt.Println(v33)
                            var v89 int = fi88
                            _ = fi88
                            _ = v89
                        }
                    }
                    switch "" {
                    case "\n\t":
                        v33 = ((0.1 + -2.25) + v33)
                    case "":
                        var v90 uint = v3
                        _ = v90
                    case "a":
                        var v91 uint32 = 0
                        _ = v91
                    default:
                        v92 := GenS9{F0: 2147483648, F1: 18446744073709551615, F2: -109}
                        _ = v92
                    }
                    v93 := GenS10{F0: 1, F1: v4, F2: -1}
                    _ = v93
                } else {
                    var v94 int = (-(fi86))
                    v33 = v33
                    _ = v94
                }
                v95 := GenS7{F0: ""}
                if true {
                    v3 = v3
                    var v96 int = 9223372036854775807
                    v96 += -2147483648
                    _ = v96
                }
                _ = fi86
                _ = v95
            }
            var v97 int = 1
            _ = v97
        } else {
            if (true || true) {
                v2 = append(v2, 2147483648)
            }
            if false {
                v98 := GenS1{F0: 32768}
                var v99 int32 = 2147483647
                for fi100 := 0; fi100 < 16; fi100++ {
                    v4 -= v4
                    fmt.Println(v4, v99, v99)
                    _ = fi100
                }
                _ = v98
                _ = v99
            }
            var v101 uint8 = 127
            _ = v101
        }
    }
    v4 -= (v4 % ((v4) | 1))
    _ = len(v1)
    var v102 int = -2147483648
    v1[(("" + "foo bar") + "")] = v102
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v102:", v102)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v102
    _ = v2
    _ = v3
    _ = v33
    _ = v4
}
