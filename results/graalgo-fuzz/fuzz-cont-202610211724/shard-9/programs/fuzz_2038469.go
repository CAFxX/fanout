package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 int8; F1 int32; F2 uint64 }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS8 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint32; F1 int32; F2 uint64 }

func main() {
    var v1 uint16 = 255
    var v2 int64 = -9223372036854775808
    v3 := []string{""}
    v4 := []uint32{2147483648, 2147483648, 2147483648}
    switch 234 {
    case -9223372036854775808:
        v5 := GenS1{F0: 127, F1: 2, F2: -111}
        _ = v5
    case -2147483648:
        var v6 int8 = -97
        _ = v6
    default:
        v2 = v2
    }
    if (v1 == 65535) {
        if ((255 < 42) && (0 <= 1)) {
            v1 = v1
        }
    } else {
        v7 := GenS1{F0: 128, F1: 834, F2: 9223372036854775807}
        _ = v7
    }
    switch ("" + "\n\t") {
    case "":
        var v8 int8 = 1
        _ = v8
    case "\n\t":
        fmt.Println(v1, v2)
    }
    if (!(3.14159 >= (0.1 / (((float32((-2.25 - (1.5 - 0.0))) * (1.5 * 1e10)) - 0.0) / ((((-2.25 * -2.25) * (float32(3.14159) + (float32(0.1) * 0.1))) / 1.0) - (3.14159 * 0.1)))))) {
        if (-178 == -9223372036854775808) {
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 uint16 = v1
                v2 = v2
                var v11 int = fi9
                _ = fi9
                _ = v10
                _ = v11
            }
            var v12 uint64 = 444
            _ = v12
        } else {
            if (((v1 ^ v1) >= 65535) && (1 < 877)) {
                var v13 int = 2
                for fi14 := 0; fi14 < 4; fi14++ {
                    fmt.Println(fi14, v2)
                    if true {
                        v15 := GenS2{F0: 1}
                        _ = v15
                    }
                    v13 += fi14
                    _ = fi14
                }
                fmt.Println(v13, v2, v2)
                _ = v13
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        fmt.Println(v2)
    }
    var v16 float32 = -2.25
    if ("\n\t" != (((("" + ("\n\t" + "hello")) + (("foo bar" + "") + "a")) + "\n\t") + ("" + ("foo bar" + "")))) {
        if false {
            v17 := GenS3{F0: 598, F1: 0}
            var v18 uint = 42
            v19 := GenS4{F0: -6, F1: -93, F2: 42}
            _ = v17
            _ = v18
            _ = v19
        } else {
            for fi20 := 0; fi20 < 17; fi20++ {
                switch "a" {
                case "\n\t":
                    var v21 float64 = -2.25
                    _ = v21
                default:
                    var v22 float64 = 1.5
                    _ = v22
                }
                _ = fi20
            }
            v1 += (^(v1))
        }
        v23 := GenS5{F0: (^(v1)), F1: 68}
        _ = v23
    } else {
        _ = len(v4)
        for fi24 := 0; fi24 < 15; fi24++ {
            fmt.Println(fi24)
            _ = fi24
        }
        v25 := GenS4{F0: 8, F1: 899, F2: 177}
        _ = v25
    }
    if false {
        v2 = 1
        switch 595 {
        case -741:
            switch 42 {
            case 949:
                for fi26 := 0; fi26 < 5; fi26++ {
                    var v27 uint32 = 280
                    fmt.Println(fi26)
                    if (v2 < -726) {
                        if ((v16 * 0.0) <= v16) {
                            var v28 int64 = (v2 & (v2 << ((271) & 63)))
                            v29 := GenS1{F0: 223, F1: 42, F2: (fi26 ^ -777)}
                            for fi30 := 0; fi30 < 16; fi30++ {
                                var v31 string = ("hello" + "hello")
                                _ = fi30
                                _ = v31
                            }
                            _ = v28
                            _ = v29
                        } else {
                            if (true && false) {
                                switch (fi26 * 128) {
                                case -249:
                                    if true {
                                        for fi32 := 0; fi32 < 8; fi32++ {
                                            v2 += (v2 / (((v2 | 1)) | 1))
                                            _ = fi32
                                        }
                                        fmt.Println(v2)
                                        var v33 int64 = v2
                                        _ = v33
                                    }
                                case -1:
                                    v2 = v2
                                }
                                for fi34 := 0; fi34 < 9; fi34++ {
                                    v35 := GenS6{F0: 0, F1: (v1 & (v1 / ((1) | 1)))}
                                    var v36 uint = 255
                                    var v37 int8 = -2
                                    _ = fi34
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                }
                                v38 := GenS1{F0: 193, F1: 42, F2: -1}
                                _ = v38
                            }
                            var v39 uint64 = 18446744073709551615
                            for fi40 := 0; fi40 < 6; fi40++ {
                                v3 = append(v3, "a")
                                _ = fi40
                            }
                            _ = v39
                        }
                    }
                    _ = fi26
                    _ = v27
                }
            case 0:
                var v41 uint16 = v1
                _ = v41
            case -9223372036854775808:
                v1 = (v1 << ((2) & 63))
            default:
                v2 -= (v2 % ((-1) | 1))
            }
        }
    } else {
        if false {
            if (false && (!(v1 >= v1))) {
                v1 = 255
                for fi42 := 0; fi42 < 8; fi42++ {
                    if (true && false) {
                        for fi43 := 0; fi43 < 3; fi43++ {
                            v44 := GenS6{F0: 208, F1: 203}
                            _ = fi43
                            _ = v44
                        }
                        var v45 int8 = -1
                        _ = v45
                    } else {
                        for fi46 := 0; fi46 < 9; fi46++ {
                            v47 := GenS2{F0: 130}
                            _ = fi46
                            _ = v47
                        }
                        var v48 string = "a"
                        if true {
                            v2 += (v2 * v2)
                            v48 = "hello"
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        } else {
                            var v49 bool = (!(((fi42 / ((4294967295) | 1)) <= (^(fi42))) && false))
                            _ = v49
                        }
                        _ = v48
                    }
                    _ = fi42
                }
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                fmt.Println(v2, v1)
            }
            var v50 int64 = -9223372036854775808
            _ = v50
        }
    }
    switch 9223372036854775807 {
    case 32768:
        v51 := GenS5{F0: (v1 | v1), F1: 4294967295}
        _ = v51
    }
    if (!(0.0 <= 0.1)) {
        v52 := GenS5{F0: v1, F1: 428}
        var v53 bool = (610 <= 394)
        v54 := GenS1{F0: 0, F1: 825, F2: -1}
        _ = v52
        _ = v53
        _ = v54
    }
    fmt.Println(v1, v16)
    for fi55 := 0; fi55 < 13; fi55++ {
        var v56 string = ("hello" + "hello")
        if false {
            var v57 int = (fi55 << ((581) & 63))
            _ = v57
        }
        _ = fi55
        _ = v56
    }
    v16 = v16
    if true {
        if true {
            v16 += (0.1 * 0.1)
            if true {
                switch "hello" {
                case "foo bar":
                    var v58 float32 = ((3.14159 * 0.0) * 0.0)
                    _ = v58
                case "\n\t":
                    v3 = append(v3, ("\n\t" + (((((((("" + "hello") + "") + "\n\t") + "") + (("foo bar" + (("" + "a") + ("\n\t" + ("foo bar" + (("hello" + "a") + "a"))))) + (("a" + "") + ("a" + (("hello" + "hello") + "hello"))))) + "\n\t") + ("hello" + ("\n\t" + "hello"))) + "")))
                default:
                    v59 := GenS1{F0: 43, F1: 42, F2: 4294967295}
                    _ = v59
                }
                fmt.Println(v16)
                if (((1e10 / ((0.0 / 1.0) * ((1.5 / 1.0) / float64((1e10 - v16))))) / (1.5 * -2.25)) < (((1.5 * 1e10) - float64(v16)) - ((((-2.25 + 1e10) + (float64(v16) * 1.5)) - (3.14159 / float64(((v16 * (1e10 + v16)) * 1.5)))) / -2.25))) {
                    v1 -= (^(v1))
                    v60 := GenS2{F0: 1}
                    _ = v60
                }
            }
            for fi61 := 0; fi61 < 6; fi61++ {
                if true {
                    fmt.Println(fi61)
                    for fi62 := 0; fi62 < 14; fi62++ {
                        if (!false) {
                            v63 := GenS7{F0: 4294967295, F1: 42, F2: (v1 + v1)}
                            var v64 int8 = 9
                            _ = v63
                            _ = v64
                        } else {
                            if ((v1 + 42) < v1) {
                                if ("foo bar" == "a") {
                                    v16 = float32(float64(float32(((0.0 - float64(v16)) + (1.5 * 1.5)))))
                                    var v65 uint32 = 217
                                    _ = v65
                                }
                                switch ("" + ("a" + "")) {
                                case "":
                                    _ = len(v4)
                                }
                                fmt.Println(v1)
                            }
                            v16 -= (0.1 / 1e10)
                            v2 = (v2 << ((965) & 63))
                        }
                        var v66 uint8 = 89
                        v66 = v66
                        if true {
                            var v67 int8 = 42
                            if false {
                                var v68 int = (fi62 / ((-9223372036854775808) | 1))
                                _ = v68
                            } else {
                                v66 = (^(v66))
                            }
                            _ = v67
                        } else {
                            v2 += (v2 + (v2 + v2))
                            v69 := GenS8{F0: v1, F1: 943, F2: 2147483647}
                            switch "hello" {
                            case "\n\t":
                                v70 := GenS4{F0: 20, F1: -1, F2: 4294967295}
                                _ = v70
                            case "hello":
                                switch "\n\t" {
                                case "\n\t":
                                    v66 -= (^(v66))
                                }
                            case "":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            default:
                                if len(v4) > 0 {
                                    v4[0] = 4294967295
                                }
                            }
                            _ = v69
                        }
                        if (v2 < v2) {
                            v1 = (v1 & v1)
                            fmt.Println(fi62, v66, v66)
                        }
                        _ = fi62
                        _ = v66
                    }
                    if false {
                        switch 128 {
                        case 523:
                            if false {
                                var v71 float64 = (-2.25 + ((0.1 / 0.1) / 3.14159))
                                for fi72 := 0; fi72 < 18; fi72++ {
                                    switch "a" {
                                    case "\n\t":
                                        var v73 int8 = 42
                                        fmt.Println(fi72, fi72)
                                        v74 := GenS1{F0: 42, F1: 83, F2: (fi61 | (fi72 >> ((120) & 63)))}
                                        _ = v73
                                        _ = v74
                                    case "":
                                        fmt.Println(fi61, fi72)
                                    case "a":
                                        if true {
                                            var v75 uint8 = 127
                                            fmt.Println(v75, fi61, v75)
                                            var v76 int64 = v2
                                            _ = v75
                                            _ = v76
                                        } else {
                                            if (!(fi72 == fi72)) {
                                                if (!false) {
                                                    v77 := GenS9{F0: (!false)}
                                                    if (!true) {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        v16 = 0.0
                                                    } else {
                                                        var v78 uint = 42
                                                        _ = v78
                                                    }
                                                    _ = v77
                                                }
                                                var v79 bool = true
                                                v80 := GenS10{F0: -32768}
                                                _ = v79
                                                _ = v80
                                            } else {
                                                var v81 float64 = v71
                                                _ = v81
                                            }
                                        }
                                    }
                                    fmt.Println(v2, v16, fi72)
                                    fmt.Println(v1, fi72, v71)
                                    _ = fi72
                                }
                                v82 := GenS7{F0: 4294967295, F1: 127, F2: 255}
                                _ = v71
                                _ = v82
                            }
                        case -9223372036854775808:
                            for fi83 := 0; fi83 < 13; fi83++ {
                                var v84 int16 = -32768
                                v85 := GenS11{F0: 524, F1: 42, F2: 42}
                                _ = fi83
                                _ = v84
                                _ = v85
                            }
                        default:
                            var v86 int32 = -1
                            _ = v86
                        }
                        v16 = 3.14159
                    }
                }
                fmt.Println(fi61)
                var v87 float64 = 0.0
                _ = fi61
                _ = v87
            }
        }
        if true {
            var v88 int64 = v2
            switch 128 {
            case 907:
                v89 := GenS1{F0: 103, F1: 65535, F2: 32768}
                _ = v89
            case 2147483647:
                v16 = float32(1e10)
            case 128:
                v1 = v1
            }
            _ = v88
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v90 int8 = -100
            _ = v90
        }
        fmt.Println(v2, v16)
    }
    v2 = v2
    fmt.Println(v16, v16)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
