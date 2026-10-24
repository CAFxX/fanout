package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int; F2 int16 }
type GenS3 struct { F0 int; F1 int; F2 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint64; F1 uint }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint8; F1 string }
func genFn0() uint {
    var v6 float64 = 1.5
    v7 := GenS1{F0: -9223372036854775808, F1: ("" + "hello"), F2: 1}
    for fi8 := 0; fi8 < 18; fi8++ {
        v6 += (v6 / v6)
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 uint32 = 4294967295
            _ = fi9
            _ = v10
        }
        fmt.Println(v6, fi8)
        _ = fi8
    }
    switch "hello" {
    case "foo bar":
        v6 -= (v6 / v6)
    case "\n\t":
        v11 := GenS1{F0: -1, F1: ("foo bar" + "\n\t"), F2: 900}
        _ = v11
    }
    v12 := GenS2{F0: 42, F1: 42, F2: 32767}
    _ = v12
    _ = v6
    _ = v7
    return 890
}

func main() {
    var v1 uint = 119
    v2 := []string{("" + ((("foo bar" + "a") + ("hello" + "")) + ("foo bar" + "foo bar"))), ("a" + "hello"), "hello", "hello"}
    var v3 uint16 = 65535
    var v4 uint8 = 207
    var v5 uint64 = 302
    fmt.Println(v5, v4)
    var v13 uint64 = v5
    v14 := GenS3{F0: 4294967295, F1: 2147483647, F2: 42}
    fmt.Println(v13, v5, v4)
    v15 := GenS4{F0: 1}
    v4 = v4
    if true {
        if len(v2) > 0 {
            v2[0] = "a"
        }
        var v16 float32 = -2.25
        _ = v16
    } else {
        for fi17 := 0; fi17 < 12; fi17++ {
            v1 = 580
            fmt.Println(v5, fi17)
            v1 -= 255
            _ = fi17
        }
        var v18 int32 = 0
        switch "" {
        case "":
            v13 += v5
        case "hello":
            for fi19 := 0; fi19 < 3; fi19++ {
                v20 := GenS2{F0: (v3 + v3), F1: fi19, F2: 0}
                if (true && (v1 > (v1 % ((v1) | 1)))) {
                    for fi21 := 0; fi21 < 11; fi21++ {
                        switch ("a" + "hello") {
                        case "\n\t":
                            v13 -= (v13 * v13)
                        }
                        _ = fi21
                    }
                    switch "foo bar" {
                    case "":
                        switch (fi19 << (((v1 % ((v1) | 1))) & 63)) {
                        case -802:
                            var v22 uint32 = 609
                            _ = v22
                        case -876:
                            switch (fi19 >> ((0) & 63)) {
                            case 42:
                                for fi23 := 0; fi23 < 7; fi23++ {
                                    if false {
                                        var v24 int16 = 32767
                                        var v25 uint16 = (v3 % (((v3 & (v3 >> (((v1 / (((v1 >> (((^(v1))) & 63))) | 1))) & 63)))) | 1))
                                        _ = v24
                                        _ = v25
                                    }
                                    _ = fi23
                                }
                            case 591:
                                var v26 int32 = v18
                                _ = v26
                            default:
                                switch "a" {
                                case "\n\t":
                                    var v27 uint32 = 114
                                    _ = v27
                                }
                            }
                        case 9223372036854775807:
                            v3 = (v3 | 65535)
                        }
                    case "\n\t":
                        switch fi19 {
                        case 600:
                            v28 := GenS1{F0: (fi19 >> (((v1 - (v1 - (v1 | 42)))) & 63)), F1: "a", F2: v13}
                            _ = v28
                        default:
                            switch "" {
                            case "\n\t":
                                _ = len(v2)
                            default:
                                var v29 uint = v1
                                _ = v29
                            }
                        }
                    }
                    v30 := GenS1{F0: int(v3), F1: (("foo bar" + "hello") + ""), F2: (v5 % (((v5 + v5)) | 1))}
                    _ = v30
                }
                var v31 int32 = v18
                _ = fi19
                _ = v20
                _ = v31
            }
        case "foo bar":
            v13 += (^(v13))
        default:
            v32 := GenS4{F0: (v13 - v13)}
            _ = v32
        }
        _ = v18
    }
    switch (("foo bar" + "hello") + ("a" + "a")) {
    case "\n\t":
        fmt.Println(v1)
    default:
        v33 := GenS5{F0: 873}
        _ = v33
    }
    v34 := GenS6{F0: (^(v1))}
    for fi35 := 0; fi35 < 5; fi35++ {
        if (v5 != 1) {
            _ = len(v2)
        } else {
            if true {
                v36 := GenS6{F0: (v1 - v1)}
                _ = v36
            }
            v37 := GenS2{F0: 255, F1: fi35, F2: 0}
            v1 += v1
            _ = v37
        }
        v2 = append(v2, ("hello" + ("hello" + (("foo bar" + "\n\t") + (("a" + "a") + "\n\t")))))
        var v38 uint8 = (v4 % (((v4 >> (((^(v1))) & 63))) | 1))
        _ = fi35
        _ = v38
    }
    for fi39 := 0; fi39 < 7; fi39++ {
        var v40 uint8 = v4
        if true {
            switch "foo bar" {
            case "hello":
                var v41 int = -9223372036854775808
                _ = v41
            case "\n\t":
                for fi42 := 0; fi42 < 4; fi42++ {
                    _ = len(v2)
                    v13 = (v5 << ((genFn0()) & 63))
                    _ = fi42
                }
            }
            var v43 string = ""
            _ = v43
        } else {
            for fi44 := 0; fi44 < 18; fi44++ {
                v45 := GenS2{F0: v3, F1: 9223372036854775807, F2: 42}
                v5 += 18446744073709551615
                _ = fi44
                _ = v45
            }
            for fi46 := 0; fi46 < 4; fi46++ {
                var v47 int8 = -85
                v2 = append(v2, "\n\t")
                v3 += uint16(v40)
                _ = fi46
                _ = v47
            }
            switch ("" + "hello") {
            case "\n\t":
                var v48 string = (("a" + ("" + "hello")) + (("foo bar" + "foo bar") + "hello"))
                _ = v48
            case "":
                switch "a" {
                case "":
                    if (!false) {
                        var v49 uint64 = v5
                        _ = v49
                    }
                case "foo bar":
                    v50 := GenS7{F0: 127}
                    _ = v50
                }
            case "a":
                _ = len(v2)
            }
        }
        switch fi39 {
        case -2147483648:
            if false {
                fmt.Println(v40)
            }
        case 0:
            switch "\n\t" {
            case "foo bar":
                if false {
                    if false {
                        v51 := GenS6{F0: v1}
                        var v52 int64 = 42
                        _ = v51
                        _ = v52
                    } else {
                        for fi53 := 0; fi53 < 20; fi53++ {
                            var v54 float64 = 0.0
                            v13 = (v13 % (((v5 ^ (v13 ^ 485))) | 1))
                            _ = fi53
                            _ = v54
                        }
                        if ("a" != ("foo bar" + ("a" + ("\n\t" + (((("" + "foo bar") + ("foo bar" + "foo bar")) + "foo bar") + "a"))))) {
                            switch "a" {
                            case "":
                                for fi55 := 0; fi55 < 12; fi55++ {
                                    if false {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        var v56 int = (fi55 & (fi55 % ((2147483647) | 1)))
                                        var v57 uint8 = v4
                                        _ = v56
                                        _ = v57
                                    }
                                    _ = fi55
                                }
                            case "hello":
                                for fi58 := 0; fi58 < 17; fi58++ {
                                    if (v1 != v1) {
                                        if len(v2) > 0 {
                                            v2[0] = "a"
                                        }
                                        var v59 uint16 = v3
                                        _ = v59
                                    } else {
                                        if true {
                                            if false {
                                                switch fi39 {
                                                case 1:
                                                    var v60 int = 128
                                                    _ = v60
                                                case -198:
                                                    var v61 uint32 = 42
                                                    _ = v61
                                                default:
                                                    var v62 int64 = -77
                                                    _ = v62
                                                }
                                            }
                                            if (false && false) {
                                                var v63 int16 = 323
                                                var v64 uint8 = v40
                                                if false {
                                                    if (v13 <= (v13 + v13)) {
                                                        if len(v2) > 0 {
                                                            v2[0] = ("a" + "foo bar")
                                                        }
                                                        var v65 float64 = (float64(1e10) * (3.14159 - 1e10))
                                                        v5 = (v5 * v5)
                                                        _ = v65
                                                    }
                                                    v3 += (^(v3))
                                                }
                                                var v66 int16 = -513
                                                _ = v63
                                                _ = v64
                                                _ = v66
                                            } else {
                                                v67 := GenS3{F0: fi39, F1: -225, F2: 2147483647}
                                                var v68 uint8 = (v4 << (((v1 * genFn0())) & 63))
                                                _ = v67
                                                _ = v68
                                            }
                                        }
                                    }
                                    v4 += 0
                                    _ = fi58
                                }
                            }
                            if true {
                                for fi69 := 0; fi69 < 2; fi69++ {
                                    _ = len(v2)
                                    _ = fi69
                                }
                                switch "foo bar" {
                                case "\n\t":
                                    switch fi39 {
                                    case 42:
                                        var v70 float32 = float32((1e10 * ((((-2.25 - ((-2.25 / float64(0.1)) * 1.5)) * (-2.25 + 1e10)) * float64(3.14159)) - 1.5)))
                                        _ = v70
                                    case 1:
                                        if false {
                                            var v71 float64 = (1e10 * 1.5)
                                            v4 = (v40 | 118)
                                            _ = v71
                                        }
                                    default:
                                        for fi72 := 0; fi72 < 7; fi72++ {
                                            _ = len(v2)
                                            _ = fi72
                                        }
                                    }
                                }
                                v4 = 1
                            }
                            if false {
                                v73 := GenS8{F0: ("\n\t" + ("" + "hello"))}
                                v1 = (^(v1))
                                v3 = v3
                                _ = v73
                            }
                        }
                    }
                    v2 = append(v2, (("\n\t" + ("hello" + "hello")) + "\n\t"))
                    if (!(true || ((true && ((-918 != -226) || (-2.25 < 1e10))) && false))) {
                        for fi74 := 0; fi74 < 6; fi74++ {
                            _ = len(v2)
                            v40 -= (v40 >> (((v1 << ((v1) & 63))) & 63))
                            v75 := GenS6{F0: (^(v1))}
                            _ = fi74
                            _ = v75
                        }
                    } else {
                        v4 = (v40 - (^(v40)))
                        v76 := GenS9{F0: (v13 % (((v5 / (((^(v13))) | 1))) | 1)), F1: 0}
                        var v77 int8 = -123
                        _ = v76
                        _ = v77
                    }
                }
            case "a":
                v4 += 135
            }
        case 128:
            if (((float64(0.1) - float64(((float32(((-2.25 / (3.14159 + 0.0)) + float64(float32(float64(-2.25))))) + (((-2.25 + ((1e10 - float32(1e10)) * 3.14159)) + (float32(0.0) - 3.14159)) - 1.5)) + ((float32((-2.25 * (1.5 / 0.1))) * 0.1) * -2.25)))) / 1.0) <= -2.25) {
                v1 = genFn0()
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                v3 = (v3 ^ (v3 / ((v3) | 1)))
            }
        default:
            switch "hello" {
            case "foo bar":
                if true {
                    v78 := GenS8{F0: "a"}
                    _ = v78
                } else {
                    if false {
                        v1 -= 307
                    }
                }
            case "hello":
                v79 := GenS10{F0: true}
                _ = v79
            }
        }
        _ = fi39
        _ = v40
    }
    var v80 float32 = ((1e10 * 0.1) / 3.14159)
    var v81 uint = v1
    v13 += v13
    v82 := GenS11{F0: (^(v4)), F1: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v80
    _ = v81
    _ = v82
}
