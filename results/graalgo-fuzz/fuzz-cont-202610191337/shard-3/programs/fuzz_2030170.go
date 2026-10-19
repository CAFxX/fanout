package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 bool; F1 uint }
type GenS8 struct { F0 uint; F1 uint32; F2 int32 }
type GenS9 struct { F0 uint64; F1 uint8 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint32; F1 uint64 }
type GenS13 struct { F0 uint64 }

func main() {
    var v1 string = "hello"
    v2 := []uint32{}
    var v3 int16 = 42
    var v4 float32 = 1e10
    var v5 uint8 = 255
    var v6 int = 4294967295
    switch v1 {
    case "hello":
        fmt.Println(v5, v4, v6)
    case "":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        switch ("hello" + ((v1 + v1) + v1)) {
        case "":
            v4 = 0.1
        case "a":
            fmt.Println(v6, v6)
        }
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        v1 = "\n\t"
        _ = fi7
    }
    if true {
        v6 = (-(v6))
    }
    if (-1 == 100) {
        if ((1e10 * 0.1) == 3.14159) {
            var v8 int64 = 632
            _ = v8
        } else {
            var v9 uint64 = 2
            var v10 int64 = 1
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 uint16 = 0
                v10 -= v10
                _ = fi11
                _ = v12
            }
            _ = v10
            _ = v9
        }
        v13 := GenS1{F0: 45, F1: v6}
        _ = len(v2)
        _ = v13
    }
    var v14 int8 = 35
    v15 := GenS2{F0: 873}
    switch v1 {
    case "\n\t":
        v16 := GenS3{F0: 255, F1: (v5 << ((4294967295) & 63)), F2: 252}
        _ = v16
    case "a":
        v3 += (-(v3))
    case "hello":
        for fi17 := 0; fi17 < 2; fi17++ {
            switch (v1 + v1) {
            case "a":
                v1 = v1
            default:
                if true {
                    v18 := GenS1{F0: 1, F1: fi17}
                    _ = v18
                } else {
                    var v19 int64 = -9223372036854775808
                    v20 := GenS4{F0: 0, F1: 1}
                    _ = v19
                    _ = v20
                }
            }
            switch (^(fi17)) {
            case -965:
                var v21 int64 = 9223372036854775807
                _ = v21
            case -665:
                if (637 < -9223372036854775808) {
                    v22 := GenS5{F0: 4294967295, F1: 255}
                    fmt.Println(v1, v4, v1)
                    _ = v22
                }
            default:
                for fi23 := 0; fi23 < 7; fi23++ {
                    v14 = (^(v14))
                    var v24 int = int(v3)
                    if (!true) {
                        if true {
                            var v25 bool = false
                            _ = v25
                        } else {
                            v26 := GenS3{F0: 255, F1: v5, F2: 4294967295}
                            var v27 int16 = (v3 / (((v3 ^ v3)) | 1))
                            _ = v26
                            _ = v27
                        }
                        _ = len(v2)
                    } else {
                        v28 := GenS2{F0: 0}
                        if len(v2) > 0 {
                            v2[0] = 1
                        }
                        var v29 uint16 = 65535
                        _ = v28
                        _ = v29
                    }
                    if false {
                        fmt.Println(v4)
                        var v30 float32 = (1e10 * -2.25)
                        _ = v30
                    }
                    _ = fi23
                    _ = v24
                }
            }
            _ = fi17
        }
    default:
        for fi31 := 0; fi31 < 4; fi31++ {
            var v32 int64 = 9223372036854775807
            v33 := GenS4{F0: 640, F1: 65535}
            if (0 > 0) {
                if (806 == 107) {
                    var v34 uint64 = 18446744073709551615
                    if (false && true) {
                        if (!((0 > -2147483648) && true)) {
                            switch v1 {
                            case "a":
                                fmt.Println(v5)
                            case "":
                                var v35 uint8 = 128
                                _ = v35
                            case "\n\t":
                                var v36 uint16 = 42
                                _ = v36
                            }
                        }
                        v4 -= (v4 + 0.1)
                        v32 = (v32 & v32)
                    }
                    var v37 float64 = 0.0
                    _ = v34
                    _ = v37
                } else {
                    fmt.Println(v4, fi31, v6)
                }
                if (false && (v3 >= (v3 * 1))) {
                    fmt.Println(v6, v6, v1)
                } else {
                    switch "hello" {
                    case "a":
                        v14 = (v14 / (((v14 % ((0) | 1))) | 1))
                    case "\n\t":
                        v38 := GenS3{F0: 156, F1: 42, F2: 2}
                        _ = v38
                    case "hello":
                        var v39 uint = 2
                        _ = v39
                    }
                    if false {
                        for fi40 := 0; fi40 < 13; fi40++ {
                            fmt.Println(v4, v14, fi31)
                            _ = fi40
                        }
                        v41 := GenS2{F0: 4294967295}
                        switch v6 {
                        case 286:
                            var v42 int32 = -827
                            _ = v42
                        }
                        _ = v41
                    } else {
                        var v43 int8 = (v14 / (((v14 % ((int8(v5)) | 1))) | 1))
                        for fi44 := 0; fi44 < 7; fi44++ {
                            if false {
                                switch (fi44 / ((727) | 1)) {
                                case 32767:
                                    v45 := GenS4{F0: 255, F1: 0}
                                    _ = v45
                                case -1:
                                    v2 = append(v2, 2147483647)
                                case 42:
                                    if true {
                                        switch ((v1 + v1) + "hello") {
                                        case "hello":
                                            if len(v2) > 0 {
                                                v2[0] = 42
                                            }
                                        case "":
                                            v46 := GenS6{F0: v14}
                                            _ = v46
                                        default:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                    } else {
                                        switch v1 {
                                        case "foo bar":
                                            v5 -= v5
                                            if true {
                                                _ = len(v2)
                                                if true {
                                                    var v47 uint = 4294967295
                                                    var v48 bool = false
                                                    var v49 int32 = 2147483647
                                                    _ = v47
                                                    _ = v48
                                                    _ = v49
                                                }
                                            } else {
                                                if (v14 >= (v14 / ((v14) | 1))) {
                                                    var v50 int16 = (v3 ^ int16(v32))
                                                    v3 = v3
                                                    v2 = append(v2, 2147483647)
                                                    _ = v50
                                                }
                                                if false {
                                                    if (true || false) {
                                                        v6 += -9223372036854775808
                                                        switch 195 {
                                                        case 32768:
                                                            switch (("\n\t" + "") + (("hello" + v1) + (v1 + "hello"))) {
                                                            case "foo bar":
                                                                switch 973 {
                                                                case -9223372036854775808:
                                                                    v14 = 1
                                                                case 32768:
                                                                    if true {
                                                                        fmt.Println(v1, v32)
                                                                    }
                                                                default:
                                                                    var v51 float32 = ((3.14159 * v4) - (v4 / (v4 / ((1.5 / ((v4 * v4) - 0.1)) - v4))))
                                                                    v52 := GenS6{F0: int8(v5)}
                                                                    v53 := GenS5{F0: 1, F1: 2}
                                                                    _ = v51
                                                                    _ = v52
                                                                    _ = v53
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        var v54 bool = (false || false)
                                                        v55 := GenS7{F0: v54, F1: 964}
                                                        v14 = (v14 >> ((2) & 63))
                                                        _ = v54
                                                        _ = v55
                                                    }
                                                    v1 = v1
                                                    fmt.Println(fi44, fi44, v6)
                                                } else {
                                                    v1 = v1
                                                }
                                            }
                                        case "hello":
                                            _ = len(v2)
                                        default:
                                            if false {
                                                if false {
                                                    v56 := GenS6{F0: (v43 ^ v43)}
                                                    var v57 int32 = 42
                                                    _ = v56
                                                    _ = v57
                                                } else {
                                                    if true {
                                                        var v58 uint8 = 232
                                                        v6 -= 0
                                                        v58 += v58
                                                        _ = v58
                                                    }
                                                    fmt.Println(v32, v6, v3)
                                                }
                                                _ = len(v2)
                                            } else {
                                                _ = len(v2)
                                                fmt.Println(v32, v32, v32)
                                                switch "a" {
                                                case "hello":
                                                    v43 += v43
                                                case "foo bar":
                                                    v59 := GenS4{F0: 2, F1: uint32(v5)}
                                                    _ = v59
                                                case "\n\t":
                                                    if len(v2) > 0 {
                                                        v2[0] = 188
                                                    }
                                                default:
                                                    v6 = fi31
                                                }
                                                if false {
                                                    v60 := GenS1{F0: 18446744073709551615, F1: -1}
                                                    _ = v60
                                                } else {
                                                    if ((((v4 - v4) - 1e10) * v4) < (v4 - v4)) {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        if (!true) {
                                                            var v61 uint32 = 1
                                                            var v62 int8 = (v14 << ((520) & 63))
                                                            var v63 int = v6
                                                            _ = v61
                                                            _ = v62
                                                            _ = v63
                                                        } else {
                                                            if false {
                                                                switch ("\n\t" + (v1 + "hello")) {
                                                                case "\n\t":
                                                                    v14 = (v43 % ((v14) | 1))
                                                                }
                                                            } else {
                                                                var v64 uint16 = 255
                                                                if false {
                                                                    var v65 bool = (((v6 + fi31) != -240) && true)
                                                                    v14 = (v43 | v43)
                                                                    _ = v65
                                                                }
                                                                _ = v64
                                                            }
                                                            var v66 uint16 = 0
                                                            _ = v66
                                                        }
                                                        v32 = v32
                                                        switch (fi31 + (fi44 / ((fi44) | 1))) {
                                                        case -91:
                                                            switch v1 {
                                                            case "a":
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    v43 = v14
                                                                }
                                                            default:
                                                                var v67 int32 = 165
                                                                _ = v67
                                                            }
                                                        }
                                                        v4 = (-2.25 + 0.0)
                                                        if len(v2) > 0 {
                                                            v2[0] = 233
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                                v68 := GenS8{F0: 30, F1: 2147483648, F2: 42}
                                _ = v68
                            } else {
                                var v69 bool = true
                                _ = v69
                            }
                            if true {
                                if true {
                                    var v70 int64 = (v32 - 42)
                                    var v71 uint64 = 0
                                    _ = len(v2)
                                    _ = v70
                                    _ = v71
                                }
                                if (!(("a" == v1) && true)) {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v43 = int8(v5)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                            }
                            switch ("a" + "foo bar") {
                            case "":
                                if (true && (803 <= 42)) {
                                    v14 -= (-(v43))
                                    v72 := GenS9{F0: 302, F1: (v5 / (((v5 / (((v5 - 1)) | 1))) | 1))}
                                    fmt.Println(v32, fi44, v32)
                                    _ = v72
                                }
                            default:
                                if true {
                                    v73 := GenS8{F0: 4294967295, F1: 2, F2: 345}
                                    v74 := GenS4{F0: 2, F1: 2147483647}
                                    var v75 int16 = (v3 >> ((65535) & 63))
                                    var v76 uint = 538
                                    switch v1 {
                                    case "hello":
                                        v77 := GenS1{F0: 207, F1: (-(fi44))}
                                        _ = v77
                                    case "":
                                        v78 := GenS8{F0: v76, F1: 479, F2: -2147483648}
                                        _ = v78
                                    case "foo bar":
                                        v79 := GenS10{F0: 728}
                                        _ = v79
                                    }
                                    if (!true) {
                                        v1 = ""
                                        var v80 uint8 = 0
                                        _ = v80
                                    } else {
                                        v1 = ""
                                        v75 = v3
                                        var v81 uint32 = 324
                                        if len(v2) > 0 {
                                            v2[0] = (v81 - 2)
                                        }
                                        fmt.Println(fi44, v1, v5)
                                        _ = v81
                                    }
                                    _ = v73
                                    _ = v74
                                    _ = v75
                                    _ = v76
                                }
                            }
                            _ = fi44
                        }
                        _ = v43
                    }
                }
                fmt.Println(v32)
            }
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    v5 += 0
    v82 := GenS11{F0: 1}
    switch 32768 {
    case -2147483648:
        v2 = append(v2, 2147483647)
    }
    for fi83 := 0; fi83 < 5; fi83++ {
        v1 = (v1 + ("" + (((v1 + ("" + v1)) + v1) + (v1 + "hello"))))
        _ = fi83
    }
    for fi84 := 0; fi84 < 7; fi84++ {
        v85 := GenS12{F0: 214, F1: 240}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi84
        _ = v85
    }
    if false {
        for fi86 := 0; fi86 < 7; fi86++ {
            v14 += (-(v14))
            _ = fi86
        }
    }
    var v87 int64 = int64(v5)
    for fi88 := 0; fi88 < 17; fi88++ {
        _ = len(v2)
        _ = len(v2)
        _ = fi88
    }
    v89 := GenS4{F0: 202, F1: 42}
    v90 := GenS13{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v82
    _ = v87
    _ = v89
    _ = v90
}
