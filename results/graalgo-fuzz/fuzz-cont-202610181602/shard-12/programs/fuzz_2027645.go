package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int64; F1 uint64; F2 int64 }
type GenS8 struct { F0 int16; F1 int64 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int8 }
func genFn0(p0 uint16, p1 int32) string {
    fmt.Println(p1, p0)
    p0 += 2
    var v6 uint = 4294967295
    for fi7 := 0; fi7 < 9; fi7++ {
        if (4294967295 > 117) {
            var v8 uint = v6
            v9 := []int{(^(fi7)), (fi7 + fi7)}
            for fi10 := 0; fi10 < 13; fi10++ {
                if true {
                    p1 = p1
                    var v11 uint16 = uint16(v6)
                    _ = v11
                }
                p0 += (p0 + 280)
                _ = fi10
            }
            _ = v8
            _ = v9
        }
        var v12 int = int(p0)
        var v13 int8 = -93
        _ = fi7
        _ = v12
        _ = v13
    }
    p0 += uint16(v6)
    _ = p0
    _ = p1
    _ = v6
    return "foo bar"
}
func genFn1(p0 uint8, p1 int32, p2 uint) int64 {
    for fi14 := 0; fi14 < 19; fi14++ {
        if false {
            if ((0.1 - float64((1e10 + (-2.25 - (0.0 / (float32((1e10 * (((((1.5 + 3.14159) - (0.1 / float64(0.1))) / ((1.5 * -2.25) + 1e10)) - (0.0 + 1.5)) + (1.5 / (0.0 + 1e10))))) - float32((1.5 - 3.14159)))))))) <= ((3.14159 * (-2.25 + 1.5)) * (0.1 + 0.0))) {
                v15 := GenS1{F0: 42}
                switch ("" + ("hello" + "\n\t")) {
                case "hello":
                    var v16 float32 = (float32(-2.25) / 1.5)
                    _ = v16
                }
                _ = v15
            }
            p0 -= (p0 + (p0 - p0))
        }
        var v17 float64 = 1e10
        _ = fi14
        _ = v17
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        for fi19 := 0; fi19 < 10; fi19++ {
            v20 := GenS2{F0: -128, F1: ("\n\t" + ("\n\t" + (((("hello" + "foo bar") + "") + "") + ("" + ("\n\t" + "")))))}
            var v21 int64 = -9223372036854775808
            _ = fi19
            _ = v20
            _ = v21
        }
        switch "foo bar" {
        case "":
            var v22 float32 = (3.14159 * ((((-2.25 / (1.5 * 1e10)) - -2.25) / 1.5) - -2.25))
            _ = v22
        case "hello":
            var v23 int = (fi18 >> ((p2) & 63))
            _ = v23
        default:
            switch "hello" {
            case "":
                if true {
                    if false {
                        var v24 uint16 = 0
                        v25 := GenS3{F0: (fi18 * fi18), F1: 2}
                        for fi26 := 0; fi26 < 7; fi26++ {
                            var v27 int = fi18
                            var v28 int64 = 42
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        _ = v24
                        _ = v25
                    }
                    switch ("foo bar" + "") {
                    case "a":
                        p1 = (p1 * (^(p1)))
                    default:
                        fmt.Println(p2, p0, p2)
                    }
                } else {
                    var v29 uint32 = 4294967295
                    v30 := GenS1{F0: (p2 + (^(p2)))}
                    var v31 uint64 = 4294967295
                    _ = v29
                    _ = v30
                    _ = v31
                }
            }
        }
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    v1 := []uint32{2147483648, 42, 2}
    var v2 uint16 = 50
    var v3 int8 = 100
    var v4 uint = 403
    var v5 float32 = 3.14159
    v4 = v4
    for fi32 := 0; fi32 < 10; fi32++ {
        if (1e10 == float64(v5)) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        var v33 float64 = -2.25
        v2 -= (^(v2))
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 17; fi34++ {
        switch fi34 {
        case 32768:
            v35 := GenS1{F0: (^(v4))}
            _ = v35
        default:
            if len(v1) > 0 {
                v1[0] = 1
            }
        }
        v1 = append(v1, 42)
        if (!false) {
            if true {
                if len(v1) > 0 {
                    v1[0] = 576
                }
                switch "hello" {
                case "":
                    switch ("foo bar" + ("\n\t" + "foo bar")) {
                    case "a":
                        v36 := GenS3{F0: fi34, F1: 176}
                        _ = v36
                    case "\n\t":
                        v37 := GenS4{F0: 38, F1: 2}
                        _ = v37
                    default:
                        if (!(false || (!true))) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    }
                default:
                    for fi38 := 0; fi38 < 7; fi38++ {
                        v39 := GenS2{F0: 1, F1: "hello"}
                        var v40 string = "\n\t"
                        var v41 int = -39
                        _ = fi38
                        _ = v39
                        _ = v40
                        _ = v41
                    }
                }
                v1 = append(v1, uint32(v4))
            }
        }
        _ = fi34
    }
    var v42 bool = (false || ((v2 + (^(v2))) >= uint16(v4)))
    for fi43 := 0; fi43 < 19; fi43++ {
        v44 := GenS3{F0: 1, F1: 4}
        v5 -= (0.1 * -2.25)
        switch (fi43 ^ 42) {
        case 9223372036854775807:
            if (!(!(!((false && v42) && (!(179 < 1)))))) {
                var v45 int8 = 127
                _ = v45
            } else {
                _ = len(v1)
                switch ("" + (("a" + "a") + (("foo bar" + ("foo bar" + ((("a" + (("" + "") + "hello")) + "\n\t") + "\n\t"))) + "a"))) {
                case "foo bar":
                    switch (("foo bar" + "") + "hello") {
                    case "":
                        if (!false) {
                            v4 -= (^(v4))
                        } else {
                            if len(v1) > 0 {
                                v1[0] = 0
                            }
                            if (1.5 >= (float64(v5) + 1.5)) {
                                var v46 float32 = v5
                                _ = v46
                            }
                        }
                    case "hello":
                        if v42 {
                            for fi47 := 0; fi47 < 11; fi47++ {
                                fmt.Println(v2, v3)
                                _ = fi47
                            }
                            fmt.Println(v3)
                        }
                    default:
                        if false {
                            switch (fi43 * fi43) {
                            case 9223372036854775807:
                                if len(v1) > 0 {
                                    v1[0] = 98
                                }
                            case -788:
                                if v42 {
                                    var v48 float64 = float64((float32(float64(v5)) / ((((((v5 - v5) * 1.5) - (((v5 / ((1e10 * (v5 * v5)) - float32(1e10))) / v5) * (1e10 / v5))) - (float32(-2.25) / float32(0.1))) + float32(0.0)) - v5)))
                                    _ = v48
                                } else {
                                    _ = len(v1)
                                }
                            case 32768:
                                v42 = v42
                            }
                            if (!true) {
                                if v42 {
                                    switch ("\n\t" + "") {
                                    case "":
                                        switch "a" {
                                        case "foo bar":
                                            if (("" == "\n\t") || v42) {
                                                v49 := GenS5{F0: fi43}
                                                v4 += 65535
                                                _ = v49
                                            }
                                        default:
                                            if v42 {
                                                for fi50 := 0; fi50 < 11; fi50++ {
                                                    fmt.Println(v4, v4, v5)
                                                    _ = fi50
                                                }
                                            } else {
                                                if v42 {
                                                    if (0 < 42) {
                                                        for fi51 := 0; fi51 < 2; fi51++ {
                                                            var v52 int8 = (^(v3))
                                                            var v53 int = (fi51 >> ((557) & 63))
                                                            if v42 {
                                                                if true {
                                                                    var v54 int8 = v3
                                                                    if v42 {
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            switch (fi43 & 0) {
                                                                            case -9223372036854775808:
                                                                                v55 := GenS4{F0: 127, F1: 42}
                                                                                _ = v55
                                                                            case 128:
                                                                                v1 = append(v1, 42)
                                                                            case 32768:
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                            default:
                                                                                var v56 uint64 = 18446744073709551615
                                                                                _ = v56
                                                                            }
                                                                        default:
                                                                            fmt.Println(v4, v5, fi43)
                                                                        }
                                                                    }
                                                                    var v57 uint16 = v2
                                                                    v1 = append(v1, 0)
                                                                    v58 := GenS1{F0: (v4 / (((v4 - (v4 + v4))) | 1))}
                                                                    _ = v54
                                                                    _ = v57
                                                                    _ = v58
                                                                } else {
                                                                    v3 = v52
                                                                    var v59 uint = (v4 / (((v4 ^ v4)) | 1))
                                                                    v2 = v2
                                                                    _ = v59
                                                                }
                                                            }
                                                            _ = fi51
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                    } else {
                                                        for fi60 := 0; fi60 < 12; fi60++ {
                                                            if len(v1) > 0 {
                                                                v1[0] = 2147483647
                                                            }
                                                            if v42 {
                                                                v61 := GenS6{F0: uint32(v4)}
                                                                if ((v42 && (1000 >= 32767)) && (-1 <= (v3 << ((v4) & 63)))) {
                                                                    v4 = (v4 << ((673) & 63))
                                                                    var v62 string = "\n\t"
                                                                    _ = v62
                                                                }
                                                                _ = v61
                                                            }
                                                            _ = fi60
                                                        }
                                                        if (594 < 1) {
                                                            _ = len(v1)
                                                            var v63 uint = (v4 - 18446744073709551615)
                                                            _ = v63
                                                        } else {
                                                            var v64 bool = (v42 && false)
                                                            var v65 uint8 = 42
                                                            var v66 int64 = 9223372036854775807
                                                            _ = v64
                                                            _ = v65
                                                            _ = v66
                                                        }
                                                        var v67 uint32 = 4294967295
                                                        _ = v67
                                                    }
                                                    var v68 int16 = 0
                                                    if ((((0.0 - 3.14159) - 1.5) + 0.1) <= -2.25) {
                                                        var v69 float32 = v5
                                                        for fi70 := 0; fi70 < 16; fi70++ {
                                                            v69 += 1.5
                                                            var v71 int64 = -481
                                                            switch -924 {
                                                            case 42:
                                                                if (-117 < 2147483647) {
                                                                    v68 += (v68 & (v68 >> ((1) & 63)))
                                                                    fmt.Println(v4, v68, fi43)
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                }
                                                            case 0:
                                                                v4 += v4
                                                            }
                                                            _ = fi70
                                                            _ = v71
                                                        }
                                                        _ = v69
                                                    }
                                                    _ = v68
                                                } else {
                                                    for fi72 := 0; fi72 < 6; fi72++ {
                                                        v73 := GenS7{F0: 0, F1: 933, F2: -9223372036854775808}
                                                        v2 = (v2 ^ (v2 / (((v2 & v2)) | 1)))
                                                        _ = fi72
                                                        _ = v73
                                                    }
                                                    v1 = append(v1, 956)
                                                }
                                                if ((!v42) || false) {
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        fmt.Println(v3, v2)
                                                    case "":
                                                        v74 := GenS8{F0: 42, F1: -311}
                                                        _ = v74
                                                    default:
                                                        if true {
                                                            var v75 int64 = 0
                                                            fmt.Println(fi43)
                                                            _ = v75
                                                        } else {
                                                            for fi76 := 0; fi76 < 1; fi76++ {
                                                                var v77 uint8 = 128
                                                                var v78 uint8 = uint8(v4)
                                                                _ = fi76
                                                                _ = v77
                                                                _ = v78
                                                            }
                                                            fmt.Println(fi43)
                                                        }
                                                    }
                                                } else {
                                                    if (46 == 255) {
                                                        var v79 uint16 = v2
                                                        switch "" {
                                                        case "foo bar":
                                                            switch (fi43 ^ fi43) {
                                                            case 0:
                                                                var v80 float32 = (0.1 / v5)
                                                                _ = v80
                                                            default:
                                                                v81 := GenS5{F0: (^(fi43))}
                                                                _ = v81
                                                            }
                                                        case "a":
                                                            v82 := GenS7{F0: -1, F1: 42, F2: 9223372036854775807}
                                                            _ = v82
                                                        default:
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        }
                                                        if ((!(!v42)) || true) {
                                                            v83 := GenS8{F0: 131, F1: -1}
                                                            if len(v1) > 0 {
                                                                v1[0] = 2147483647
                                                            }
                                                            _ = v83
                                                        } else {
                                                            var v84 int16 = 1000
                                                            _ = v84
                                                        }
                                                        _ = v79
                                                    }
                                                    v3 -= (v3 * -125)
                                                }
                                                if v42 {
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        var v85 string = ((("a" + "hello") + (("\n\t" + "a") + "hello")) + "hello")
                                                        _ = v85
                                                    case "\n\t":
                                                        v5 = v5
                                                    case "":
                                                        fmt.Println(v5)
                                                    }
                                                    if v42 {
                                                        switch -9223372036854775808 {
                                                        case 128:
                                                            if (359 > 765) {
                                                                v3 = (v3 + (-(v3)))
                                                                v4 -= (v4 | v4)
                                                                v3 += (^(v3))
                                                            } else {
                                                                v5 += v5
                                                            }
                                                        case -729:
                                                            var v86 uint16 = (^(v2))
                                                            _ = v86
                                                        case -2147483648:
                                                            v87 := GenS9{F0: -1}
                                                            _ = v87
                                                        default:
                                                            _ = len(v1)
                                                        }
                                                        v4 += (v4 ^ v4)
                                                    }
                                                } else {
                                                    v42 = (!true)
                                                }
                                            }
                                        }
                                    case "a":
                                        var v88 uint8 = 139
                                        _ = v88
                                    case "hello":
                                        var v89 uint16 = v2
                                        _ = v89
                                    }
                                } else {
                                    var v90 int32 = int32(v3)
                                    _ = v90
                                }
                                v5 -= v5
                            }
                        } else {
                            _ = len(v1)
                            fmt.Println(v42, v5)
                        }
                    }
                case "a":
                    v4 += 255
                }
                _ = len(v1)
            }
        case 42:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 610:
            fmt.Println(v2)
        default:
            var v91 int32 = 42
            _ = v91
        }
        _ = fi43
        _ = v44
    }
    v92 := GenS7{F0: 9223372036854775807, F1: 42, F2: -1}
    var v93 int16 = 1
    for fi94 := 0; fi94 < 7; fi94++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        fmt.Println(v4)
        _ = fi94
    }
    var v95 float32 = v5
    v96 := GenS8{F0: 726, F1: -1}
    for fi97 := 0; fi97 < 1; fi97++ {
        var v98 int32 = 2147483647
        fmt.Println(fi97, v42, v2)
        _ = fi97
        _ = v98
    }
    for fi99 := 0; fi99 < 4; fi99++ {
        fmt.Println(v3, v93)
        fmt.Println(v95, v4, v4)
        v100 := GenS10{F0: 6}
        _ = fi99
        _ = v100
    }
    var v101 int32 = -2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v93:", v93)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v92
    _ = v93
    _ = v95
    _ = v96
}
