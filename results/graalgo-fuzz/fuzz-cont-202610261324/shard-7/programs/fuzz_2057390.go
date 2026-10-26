package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS5 struct { F0 int32; F1 int; F2 uint64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int; F1 int64 }
type GenS9 struct { F0 int64; F1 int64; F2 uint16 }
type GenS10 struct { F0 int32 }

func main() {
    var v1 int16 = -871
    var v2 bool = true
    var v3 uint32 = 2
    var v4 string = ((("foo bar" + "") + "a") + "foo bar")
    v5 := []string{}
    v6 := make(map[string]string)
    if (!((float64(1e10) <= float64(((3.14159 - 3.14159) + (float32(1.5) + (0.0 / 3.14159))))) || (114 > 127))) {
        _ = len(v5)
        if v2 {
            var v7 int32 = -1
            var v8 int8 = int8(v7)
            _ = v7
            _ = v8
        } else {
            switch 4294967295 {
            case -622:
                v1 = 42
            case 320:
                v4 = (v4 + ((v4 + "hello") + "\n\t"))
            case 2147483647:
                v3 = 65535
            default:
                v4 = "foo bar"
            }
            var v9 float32 = (1e10 / (1.5 - 0.1))
            var v10 bool = false
            _ = v10
            _ = v9
        }
    } else {
        switch 42 {
        case 4294967295:
            if (1 != 1) {
                v3 += (^(v3))
                v2 = (0 >= 745)
                if v2 {
                    var v11 uint32 = (v3 - 1)
                    if ((v3 & v3) <= v3) {
                        v2 = true
                    }
                    _ = v11
                }
            }
        case 0:
            v12 := GenS1{F0: 4294967295}
            _ = v12
        }
        if (!(!v2)) {
            v13 := GenS1{F0: -2147483648}
            v2 = v2
            v2 = v2
            _ = v13
        } else {
            switch 128 {
            case -174:
                fmt.Println(v2, v4)
            default:
                v2 = ((v1 % (((^(v1))) | 1)) <= v1)
            }
            v3 = v3
        }
        if true {
            v1 -= (v1 / (((v1 % (((v1 << ((772) & 63))) | 1))) | 1))
            for fi14 := 0; fi14 < 16; fi14++ {
                var v15 uint32 = 1
                _ = fi14
                _ = v15
            }
            if v2 {
                v4 = "foo bar"
                var v16 bool = ((403 > 863) || v2)
                _ = v16
            }
        } else {
            v17 := GenS2{F0: 128, F1: 2, F2: 18446744073709551615}
            v6[(v4 + "hello")] = v4
            v18 := GenS2{F0: 42, F1: 797, F2: 2}
            _ = v17
            _ = v18
        }
    }
    if true {
        for fi19 := 0; fi19 < 8; fi19++ {
            fmt.Println(v1)
            for fi20 := 0; fi20 < 18; fi20++ {
                _ = len(v6)
                fmt.Println(v1, v3, fi19)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = len(v6)
                fmt.Println(v1, fi20)
                v4 = "hello"
                v6["foo bar"] = v4
                _ = fi20
            }
            if (42 == 127) {
                v2 = false
            }
            _ = fi19
        }
        delete(v6, (v4 + "foo bar"))
    }
    switch 4294967295 {
    case 0:
        fmt.Println(v4, v2, v4)
    }
    if false {
        if v2 {
            fmt.Println(v4)
            if v2 {
                v2 = (2 == 4294967295)
            } else {
                v1 -= (v1 >> ((0) & 63))
            }
        }
        for fi21 := 0; fi21 < 4; fi21++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if (1 < 0) {
                for fi22 := 0; fi22 < 6; fi22++ {
                    switch (^(fi22)) {
                    case -166:
                        v23 := GenS1{F0: (-(fi22))}
                        switch fi21 {
                        case 32768:
                            if (-1 == 683) {
                                v24 := GenS3{F0: 127, F1: 1}
                                v2 = v2
                                var v25 bool = (-466 > 505)
                                v26 := GenS4{F0: v3, F1: 421, F2: 4294967295}
                                _ = v24
                                _ = v25
                                _ = v26
                            }
                        default:
                            if v2 {
                                v4 = v4
                                var v27 uint64 = 29
                                _ = v27
                            }
                        }
                        if (!(v2 || v2)) {
                            v28 := GenS1{F0: -653}
                            fmt.Println(v3)
                            v3 -= v3
                            _ = v28
                        }
                        _ = v23
                    case 127:
                        if (true || false) {
                            var v29 bool = v2
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            switch 1 {
                            case 9223372036854775807:
                                if false {
                                    var v30 uint16 = 65535
                                    var v31 int = fi22
                                    if v2 {
                                        delete(v6, ("foo bar" + v4))
                                        var v32 float64 = 1e10
                                        _ = v32
                                    }
                                    _ = v30
                                    _ = v31
                                }
                            default:
                                switch "a" {
                                case "a":
                                    if ((!false) && (0 <= 698)) {
                                        v33 := GenS5{F0: -2147483648, F1: (^(fi21)), F2: 375}
                                        v1 += (v1 * v1)
                                        _ = v33
                                    } else {
                                        var v34 int8 = 42
                                        if (v29 && (false || ((^(v1)) > v1))) {
                                            if (!(v2 || v29)) {
                                                v6[v4] = v4
                                                v3 = 90
                                                fmt.Println(fi21, v3)
                                                fmt.Println(v1)
                                            }
                                            if ((255 < 42) || (int8(v1) == -90)) {
                                                fmt.Println(fi22, v29, v4)
                                                fmt.Println(v4, v2)
                                            } else {
                                                var v35 string = v4
                                                _ = v35
                                            }
                                            var v36 string = (("\n\t" + (v4 + v4)) + "hello")
                                            _ = v36
                                        } else {
                                            v3 = (v3 % (((v3 % (((v3 >> ((282) & 63))) | 1))) | 1))
                                        }
                                        var v37 float64 = 1e10
                                        _ = v34
                                        _ = v37
                                    }
                                case "\n\t":
                                    v38 := GenS5{F0: 127, F1: int(v1), F2: 1}
                                    _ = v38
                                case "hello":
                                    var v39 int32 = -83
                                    _ = v39
                                default:
                                    v40 := v6[(("hello" + "a") + "a")]
                                    _ = v40
                                }
                            }
                            _ = v29
                        } else {
                            switch (v4 + v4) {
                            case "a":
                                v3 = 2147483648
                                v1 = v1
                            case "hello":
                                v41 := GenS4{F0: (v3 | (v3 / (((^(v3))) | 1))), F1: 569, F2: 972}
                                _ = v41
                            case "":
                                _ = len(v5)
                            }
                            var v42 bool = false
                            switch "hello" {
                            case "a":
                                if len(v5) > 0 {
                                    v5[0] = v4
                                }
                            }
                            _ = v42
                        }
                    default:
                        delete(v6, v4)
                    }
                    var v43 int8 = 127
                    fmt.Println(v1, v43, fi22)
                    _ = fi22
                    _ = v43
                }
            } else {
                for fi44 := 0; fi44 < 9; fi44++ {
                    var v45 int64 = 42
                    v1 -= -1
                    v2 = v2
                    v3 += v3
                    _ = fi44
                    _ = v45
                }
            }
            v3 = (^(v3))
            _ = fi21
        }
        var v46 uint = 18446744073709551615
        _ = v46
    }
    fmt.Println(v3, v1)
    var v47 int = 91
    v48 := GenS5{F0: 653, F1: 32768, F2: 246}
    for fi49 := 0; fi49 < 4; fi49++ {
        v5 = append(v5, v4)
        v47 -= fi49
        _ = fi49
    }
    var v50 int32 = int32(v3)
    for fi51 := 0; fi51 < 16; fi51++ {
        switch ((v4 + (v4 + "\n\t")) + "hello") {
        case "hello":
            for fi52 := 0; fi52 < 19; fi52++ {
                delete(v6, (v4 + ""))
                var v53 uint = 266
                switch fi52 {
                case -143:
                    if false {
                        fmt.Println(v50)
                        if (1.5 >= 0.0) {
                            v54 := v6["foo bar"]
                            v50 = (v50 % (((v50 + v50)) | 1))
                            switch 1 {
                            case 6:
                                if (!true) {
                                    v55 := GenS6{F0: v2}
                                    if (((v50 & v50) <= (v50 - v50)) && v2) {
                                        var v56 float64 = float64((0.0 * -2.25))
                                        v57 := GenS2{F0: 102, F1: 4294967295, F2: 42}
                                        _ = v56
                                        _ = v57
                                    } else {
                                        fmt.Println(v54)
                                        if (v3 <= v3) {
                                            v50 = v50
                                            var v58 uint16 = 42
                                            v1 = (v1 << ((v53) & 63))
                                            _ = v58
                                        } else {
                                            v59 := GenS7{F0: (v53 << ((v53) & 63))}
                                            var v60 uint16 = 42
                                            var v61 uint32 = (v3 | (v3 & v3))
                                            _ = v59
                                            _ = v60
                                            _ = v61
                                        }
                                        var v62 uint = (v53 << (((v53 >> ((v53) & 63))) & 63))
                                        var v63 int16 = v1
                                        v64 := GenS2{F0: 128, F1: 748, F2: 4294967295}
                                        v65 := GenS5{F0: (v50 % ((-1) | 1)), F1: fi52, F2: 0}
                                        _ = v62
                                        _ = v63
                                        _ = v64
                                        _ = v65
                                    }
                                    _ = v55
                                }
                            case -294:
                                var v66 int64 = -9223372036854775808
                                _ = v66
                            default:
                                var v67 uint8 = 128
                                _ = v67
                            }
                            var v68 int8 = 42
                            _ = v54
                            _ = v68
                        }
                    }
                case -622:
                    if true {
                        fmt.Println(v50, fi52)
                        v69 := GenS8{F0: 703, F1: 1}
                        var v70 uint16 = 406
                        _ = v69
                        _ = v70
                    }
                    if len(v5) > 0 {
                        v5[0] = v4
                    }
                    var v71 uint32 = 42
                    _ = v71
                case 1:
                    if v2 {
                        switch (v4 + "\n\t") {
                        case "hello":
                            if (65535 == 1) {
                                v72 := v6[v4]
                                switch int(v53) {
                                case 42:
                                    switch fi51 {
                                    case 42:
                                        v73 := GenS9{F0: -1, F1: 557, F2: 0}
                                        _ = v73
                                    case 611:
                                        if len(v5) > 0 {
                                            v5[0] = (v72 + ("a" + v72))
                                        }
                                    default:
                                        fmt.Println(v53)
                                    }
                                case 340:
                                    switch v4 {
                                    case "hello":
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    case "":
                                        v3 += v3
                                    }
                                case -9223372036854775808:
                                    fmt.Println(v2, v72)
                                }
                                _ = v72
                            } else {
                                v6[v4] = ((v4 + (v4 + "foo bar")) + v4)
                                v74 := GenS10{F0: 2147483647}
                                fmt.Println(v53)
                                _ = v74
                            }
                            var v75 int32 = (^(v50))
                            _ = v75
                        }
                        switch 0 {
                        case 779:
                            if (!v2) {
                                var v76 uint8 = 68
                                v3 -= 201
                                v53 = 18446744073709551615
                                _ = v76
                            }
                        case -2147483648:
                            var v77 int64 = int64(fi51)
                            var v78 int32 = (-(v50))
                            _ = v77
                            _ = v78
                        case -9223372036854775808:
                            v79 := GenS7{F0: v53}
                            _ = v79
                        }
                        if false {
                            var v80 int64 = 1
                            v6["hello"] = ""
                            _ = v80
                        } else {
                            v50 -= v50
                            fmt.Println(v53)
                        }
                    } else {
                        v2 = (v2 && (0.1 > 0.0))
                    }
                }
                _ = fi52
                _ = v53
            }
        case "\n\t":
            var v81 float64 = float64(0.1)
            _ = v81
        case "foo bar":
            v4 = (((((("foo bar" + (v4 + "")) + v4) + "foo bar") + v4) + v4) + "hello")
        }
        _ = fi51
    }
    v2 = ((v47 - (v47 ^ v47)) == (v47 / (((v47 | v47)) | 1)))
    v4 = v4
    _ = len(v6)
    delete(v6, (v4 + v4))
    switch ((("" + v4) + v4) + (("a" + v4) + (v4 + ("" + "a")))) {
    case "a":
        if ((v3 | (v3 >> ((18446744073709551615) & 63))) <= (v3 << ((18446744073709551615) & 63))) {
            var v82 int8 = 1
            var v83 string = v4
            _ = v82
            _ = v83
        } else {
            if (0.0 != ((1e10 / 1.0) + -2.25)) {
                if v2 {
                    for fi84 := 0; fi84 < 16; fi84++ {
                        var v85 uint8 = 2
                        _ = fi84
                        _ = v85
                    }
                } else {
                    for fi86 := 0; fi86 < 8; fi86++ {
                        if v2 {
                            for fi87 := 0; fi87 < 20; fi87++ {
                                if true {
                                    if (513 == 265) {
                                        var v88 uint16 = 255
                                        switch v4 {
                                        case "a":
                                            v47 = v47
                                        default:
                                            v88 -= 65535
                                        }
                                        _ = v88
                                    } else {
                                        fmt.Println(v50, fi86, v47)
                                        delete(v6, (v4 + v4))
                                    }
                                    v47 += (fi87 | 26)
                                    v47 = fi86
                                }
                                v89 := v6[v4]
                                v89 = v4
                                _ = fi87
                                _ = v89
                            }
                            v2 = true
                        }
                        _ = fi86
                    }
                    var v90 int64 = -9223372036854775808
                    if (int32(v1) > v50) {
                        v91 := GenS1{F0: v47}
                        var v92 int16 = 1000
                        _ = v91
                        _ = v92
                    }
                    _ = v90
                }
                v47 -= v47
                v93 := GenS3{F0: 89, F1: 127}
                _ = v93
            }
            v94 := v6[v4]
            v95 := GenS1{F0: (v47 - 128)}
            _ = v94
            _ = v95
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v50
    _ = v6
}
