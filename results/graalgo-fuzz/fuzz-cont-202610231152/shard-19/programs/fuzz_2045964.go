package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 uint32; F1 string; F2 int16 }
type GenS8 struct { F0 int32; F1 string; F2 int }
type GenS9 struct { F0 bool }

func main() {
    var v1 int = 127
    var v2 int32 = 42
    v3 := []uint32{2147483647, 128, 4294967295}
    var v4 float64 = -2.25
    var v5 int8 = 0
    var v6 float32 = (3.14159 / 3.14159)
    if (((v2 + v2) >= (-(v2))) || (float32(-2.25) <= v6)) {
        var v7 uint8 = 114
        var v8 float64 = (v4 - (v4 / v4))
        v9 := GenS1{F0: 802, F1: v7}
        _ = v7
        _ = v8
        _ = v9
    }
    fmt.Println(v5, v6, v1)
    v2 = 1
    for fi10 := 0; fi10 < 19; fi10++ {
        v5 -= (v5 % ((int8(v1)) | 1))
        for fi11 := 0; fi11 < 13; fi11++ {
            v4 += ((v4 + v4) * v4)
            _ = fi11
        }
        var v12 int32 = (v2 << ((401) & 63))
        _ = fi10
        _ = v12
    }
    v13 := GenS1{F0: 478, F1: 128}
    for fi14 := 0; fi14 < 7; fi14++ {
        if false {
            if (true && false) {
                switch ("" + "foo bar") {
                case "a":
                    fmt.Println(fi14)
                case "\n\t":
                    v15 := GenS2{F0: "\n\t"}
                    _ = v15
                case "foo bar":
                    var v16 int32 = v2
                    _ = v16
                }
                switch "\n\t" {
                case "foo bar":
                    var v17 int16 = 1
                    _ = v17
                case "\n\t":
                    v4 = 0.1
                case "hello":
                    switch fi14 {
                    case 138:
                        v1 = fi14
                    case 813:
                        v18 := GenS3{F0: "a"}
                        _ = v18
                    default:
                        for fi19 := 0; fi19 < 13; fi19++ {
                            switch fi19 {
                            case 262:
                                if false {
                                    var v20 uint16 = 145
                                    _ = v20
                                } else {
                                    if true {
                                        switch "\n\t" {
                                        case "hello":
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        case "\n\t":
                                            v5 += (v5 >> ((132) & 63))
                                        case "a":
                                            switch v1 {
                                            case 9223372036854775807:
                                                fmt.Println(v1)
                                            }
                                        default:
                                            fmt.Println(v6, fi19, v5)
                                        }
                                        fmt.Println(v1)
                                    } else {
                                        v1 = fi14
                                    }
                                    var v21 uint8 = 128
                                    v2 = 978
                                    _ = v21
                                }
                            default:
                                fmt.Println(fi14, v1)
                            }
                            fmt.Println(v6, v6)
                            _ = fi19
                        }
                    }
                default:
                    v3 = append(v3, 898)
                }
                for fi22 := 0; fi22 < 8; fi22++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v23 := GenS1{F0: 2, F1: 42}
                    _ = fi22
                    _ = v23
                }
            }
            var v24 uint16 = 1
            var v25 float64 = v4
            _ = v24
            _ = v25
        }
        fmt.Println(v5, v1, v1)
        var v26 uint32 = 2147483647
        _ = fi14
        _ = v26
    }
    if (65535 != 198) {
        v27 := GenS4{F0: (-(v2)), F1: 815}
        _ = v27
    } else {
        switch "hello" {
        case "":
            fmt.Println(v6, v4)
        case "foo bar":
            v28 := GenS2{F0: ("hello" + ("hello" + ("" + ("" + (("\n\t" + "foo bar") + "hello")))))}
            _ = v28
        default:
            v5 = v5
        }
        if false {
            if (!(true && true)) {
                var v29 int32 = (v2 >> ((4294967295) & 63))
                _ = v29
            }
            v5 = v5
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
        }
        v1 = v1
    }
    fmt.Println(v1, v6)
    v4 = (1.5 / 3.14159)
    v30 := GenS4{F0: 2147483647, F1: 2}
    if (!false) {
        v5 = 81
    } else {
        fmt.Println(v1)
    }
    v3 = append(v3, 4294967295)
    if false {
        v3 = append(v3, 65535)
        _ = len(v3)
    } else {
        if (true && true) {
            var v31 int = 1
            var v32 uint32 = 4294967295
            _ = v31
            _ = v32
        } else {
            v33 := GenS5{F0: (^(v2)), F1: true}
            if false {
                if false {
                    v2 += (v2 | 0)
                } else {
                    if ((v1 & 400) < (v1 - v1)) {
                        v4 += -2.25
                        v6 += -2.25
                        var v34 float64 = v4
                        _ = v34
                    }
                }
                v35 := GenS6{F0: (v2 << ((65535) & 63)), F1: 42}
                if (false && false) {
                    var v36 int32 = v2
                    var v37 uint64 = 836
                    _ = len(v3)
                    _ = v36
                    _ = v37
                } else {
                    var v38 float64 = ((0.1 - float64(v6)) * 3.14159)
                    _ = v38
                }
                _ = v35
            }
            for fi39 := 0; fi39 < 9; fi39++ {
                fmt.Println(v4, v6, v5)
                var v40 int64 = 539
                _ = fi39
                _ = v40
            }
            _ = v33
        }
        for fi41 := 0; fi41 < 3; fi41++ {
            for fi42 := 0; fi42 < 12; fi42++ {
                v4 -= (0.0 * v4)
                if true {
                    if true {
                        fmt.Println(fi42, v4)
                    } else {
                        _ = len(v3)
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    fmt.Println(v2)
                    switch (fi42 * (fi42 % ((9223372036854775807) | 1))) {
                    case -9223372036854775808:
                        switch ("\n\t" + "\n\t") {
                        case "\n\t":
                            v2 += v2
                        default:
                            if false {
                                var v43 float64 = 1e10
                                if false {
                                    v1 = (fi41 + (fi41 & 9223372036854775807))
                                    fmt.Println(v5, v1, v4)
                                    v4 -= (((float64(1e10) + float64(float32((3.14159 / v43)))) - (float64((v6 - (float32(v43) - 1.5))) + ((1e10 + -2.25) + float64(v6)))) / 1.0)
                                } else {
                                    fmt.Println(v43, v4)
                                    if false {
                                        if (-2147483648 > v2) {
                                            var v44 uint32 = 2147483647
                                            switch v1 {
                                            case -260:
                                                if true {
                                                    v3 = append(v3, 65535)
                                                    if ((2 <= 18446744073709551615) && ((1e10 <= ((v43 / (v4 * 1e10)) - v43)) || (int8(v1) <= v5))) {
                                                        v5 -= 96
                                                        switch "hello" {
                                                        case "\n\t":
                                                            var v45 uint8 = 42
                                                            _ = v45
                                                        default:
                                                            v46 := GenS4{F0: v2, F1: 0}
                                                            _ = v46
                                                        }
                                                    } else {
                                                        var v47 bool = ((v2 | int32(v1)) < (v2 ^ (^(v2))))
                                                        _ = v47
                                                    }
                                                }
                                            default:
                                                _ = len(v3)
                                            }
                                            _ = v44
                                        } else {
                                            fmt.Println(v1)
                                            var v48 uint32 = 215
                                            fmt.Println(v48, v2)
                                            _ = v48
                                        }
                                    } else {
                                        fmt.Println(v43, v5)
                                        if false {
                                            v2 += v2
                                            var v49 int8 = (v5 | 0)
                                            if ((v6 - ((float32((-2.25 / 1.5)) / (0.0 + (v6 * v6))) / v6)) == v6) {
                                                v50 := GenS2{F0: "\n\t"}
                                                v6 = 1e10
                                                switch (fi42 % (((fi41 - (fi41 % ((fi41) | 1)))) | 1)) {
                                                case 120:
                                                    var v51 uint64 = 0
                                                    _ = v51
                                                default:
                                                    fmt.Println(v4)
                                                }
                                                _ = v50
                                            } else {
                                                _ = len(v3)
                                                v52 := GenS4{F0: v2, F1: 474}
                                                _ = v52
                                            }
                                            _ = v49
                                        }
                                    }
                                    v53 := GenS4{F0: v2, F1: 30}
                                    _ = v53
                                }
                                _ = v43
                            }
                        }
                    case 127:
                        v6 = v6
                    default:
                        switch (("hello" + ("hello" + "a")) + "a") {
                        case "foo bar":
                            if false {
                                var v54 uint = 464
                                switch (v1 / ((v1) | 1)) {
                                case 32767:
                                    v3 = append(v3, 1)
                                case 2147483647:
                                    var v55 uint64 = 4294967295
                                    _ = v55
                                case 128:
                                    if true {
                                        var v56 uint8 = 127
                                        switch ("" + "") {
                                        case "a":
                                            var v57 int = v1
                                            v1 -= v1
                                            _ = v57
                                        case "":
                                            if false {
                                                switch (fi42 + v1) {
                                                case -9223372036854775808:
                                                    fmt.Println(v6)
                                                }
                                                fmt.Println(v6, v6, v1)
                                                var v58 uint32 = 939
                                                _ = v58
                                            }
                                        default:
                                            v3 = append(v3, 130)
                                        }
                                        _ = v56
                                    }
                                }
                                _ = v54
                            } else {
                                v5 = (v5 | (v5 / ((int8(v1)) | 1)))
                            }
                        case "hello":
                            var v59 bool = true
                            _ = v59
                        }
                    }
                    v3 = append(v3, 4294967295)
                }
                var v60 uint = 18446744073709551615
                _ = fi42
                _ = v60
            }
            if (!(2147483648 == 548)) {
                switch ("hello" + "\n\t") {
                case "foo bar":
                    for fi61 := 0; fi61 < 4; fi61++ {
                        v62 := GenS2{F0: ("foo bar" + "a")}
                        var v63 string = ""
                        if false {
                            switch (fi41 | v1) {
                            case 1:
                                var v64 uint64 = 18446744073709551615
                                if len(v3) > 0 {
                                    v3[0] = 2147483648
                                }
                                v2 = (v2 % (((v2 ^ (v2 % ((695) | 1)))) | 1))
                                if len(v3) > 0 {
                                    v3[0] = 4294967295
                                }
                                _ = v64
                            case -2147483648:
                                v63 = v63
                            default:
                                if len(v3) > 0 {
                                    v3[0] = 42
                                }
                            }
                            v5 += v5
                            fmt.Println(v2)
                        } else {
                            fmt.Println(v2, v4, v1)
                            v65 := GenS7{F0: 406, F1: v63, F2: 1}
                            switch ("" + v63) {
                            case "":
                                v5 = (v5 - (v5 % ((100) | 1)))
                            case "\n\t":
                                switch (fi41 - 671) {
                                case -135:
                                    v66 := GenS6{F0: -1, F1: v5}
                                    _ = v66
                                case 128:
                                    if len(v3) > 0 {
                                        v3[0] = 2147483648
                                    }
                                }
                            case "hello":
                                var v67 uint32 = 109
                                _ = v67
                            default:
                                v63 = v63
                            }
                            if (("a" + v63) < "foo bar") {
                                var v68 int8 = v5
                                v4 = v4
                                var v69 float32 = 0.0
                                _ = v68
                                _ = v69
                            }
                            if true {
                                v70 := GenS8{F0: v2, F1: v63, F2: 900}
                                _ = v70
                            }
                            var v71 int16 = -39
                            _ = v65
                            _ = v71
                        }
                        _ = len(v3)
                        _ = fi61
                        _ = v62
                        _ = v63
                    }
                case "":
                    v2 += 0
                case "\n\t":
                    v1 = fi41
                default:
                    switch (-(fi41)) {
                    case 4294967295:
                        var v72 uint8 = 255
                        _ = v72
                    case 9223372036854775807:
                        var v73 int64 = 0
                        _ = v73
                    default:
                        v74 := GenS9{F0: (false && true)}
                        _ = v74
                    }
                }
                fmt.Println(v5, v5)
                v6 += v6
            } else {
                var v75 uint16 = 32767
                _ = v75
            }
            _ = fi41
        }
    }
    v2 = (v2 % ((v2) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
