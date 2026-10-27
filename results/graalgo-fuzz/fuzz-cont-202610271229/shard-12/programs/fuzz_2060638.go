package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 int64; F1 int32; F2 uint32 }
type GenS4 struct { F0 int32; F1 uint8; F2 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS8 struct { F0 bool; F1 bool; F2 uint32 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint16; F1 uint }
type GenS12 struct { F0 uint8; F1 uint16; F2 int }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]bool)
    var v3 uint = 1
    v4 := []string{}
    var v5 uint = 493
    _ = len(v1)
    v5 = 255
    v6 := GenS1{F0: 2147483647}
    v7 := GenS2{F0: 9223372036854775807, F1: 1}
    v5 = v3
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint32 = 917
        fmt.Println(v5)
        v4 = append(v4, "")
        _ = fi8
        _ = v9
    }
    fmt.Println(v3, v3)
    v3 += (v3 | v5)
    if false {
        if ((!false) || (true && (42 > 4294967295))) {
            var v10 int32 = -863
            _ = len(v4)
            _ = v10
        } else {
            fmt.Println(v5, v3)
            if (1 <= 0) {
                if (0 == 1) {
                    switch "a" {
                    case "\n\t":
                        if (23 != 2) {
                            for fi11 := 0; fi11 < 20; fi11++ {
                                fmt.Println(v5, v5, v3)
                                var v12 int8 = 127
                                _ = fi11
                                _ = v12
                            }
                            if true {
                                var v13 uint16 = 823
                                _ = v13
                            }
                        } else {
                            if true {
                                if true {
                                    var v14 float64 = 1.5
                                    _ = v14
                                }
                                for fi15 := 0; fi15 < 3; fi15++ {
                                    v3 += v3
                                    _ = fi15
                                }
                            }
                            if true {
                                fmt.Println(v5, v3)
                                var v16 uint8 = 2
                                _ = v16
                            } else {
                                v3 += v3
                                var v17 uint16 = 0
                                _ = v17
                            }
                        }
                    case "a":
                        var v18 uint8 = 42
                        _ = v18
                    default:
                        v19 := GenS2{F0: -44, F1: -128}
                        _ = v19
                    }
                    v5 = (v3 & (v3 * (v5 + v5)))
                    v20 := GenS3{F0: 42, F1: 426, F2: 2}
                    _ = v20
                }
            }
        }
        v3 = (v3 | 417)
    } else {
        v5 = 760
        if false {
            for fi21 := 0; fi21 < 18; fi21++ {
                fmt.Println(v3, v5)
                v22 := GenS1{F0: 87}
                for fi23 := 0; fi23 < 8; fi23++ {
                    if (-2.25 <= (0.1 * 0.0)) {
                        fmt.Println(fi23)
                        if (false || true) {
                            _ = len(v4)
                        }
                        var v24 int = -299
                        _ = v24
                    } else {
                        v25 := GenS1{F0: 885}
                        _ = v25
                    }
                    if false {
                        fmt.Println(v5, v3, v5)
                        v26 := GenS4{F0: 582, F1: 0, F2: 364}
                        var v27 float32 = float32((1.5 / 1.5))
                        _ = v26
                        _ = v27
                    }
                    var v28 int64 = 0
                    _ = fi23
                    _ = v28
                }
                _ = fi21
                _ = v22
            }
        } else {
            if false {
                switch "hello" {
                case "foo bar":
                    if false {
                        if (v3 != 226) {
                            switch ("hello" + "\n\t") {
                            case "\n\t":
                                var v29 int64 = 42
                                _ = v29
                            case "hello":
                                if false {
                                    switch 82 {
                                    case -364:
                                        for fi30 := 0; fi30 < 10; fi30++ {
                                            var v31 uint8 = 1
                                            _ = fi30
                                            _ = v31
                                        }
                                    case 32768:
                                        var v32 uint = (v5 * v5)
                                        _ = v32
                                    }
                                    v3 = (v5 + (^(v3)))
                                }
                            case "a":
                                for fi33 := 0; fi33 < 16; fi33++ {
                                    for fi34 := 0; fi34 < 15; fi34++ {
                                        v35 := GenS3{F0: 0, F1: -127, F2: 2}
                                        fmt.Println(v3)
                                        _ = fi34
                                        _ = v35
                                    }
                                    switch "hello" {
                                    case "a":
                                        var v36 float32 = float32(1e10)
                                        _ = v36
                                    case "hello":
                                        v3 -= (v3 << ((v5) & 63))
                                    }
                                    var v37 uint8 = 2
                                    _ = fi33
                                    _ = v37
                                }
                            default:
                                v3 += v5
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    } else {
                        for fi38 := 0; fi38 < 8; fi38++ {
                            fmt.Println(fi38, v5)
                            for fi39 := 0; fi39 < 18; fi39++ {
                                if false {
                                    v40 := GenS5{F0: fi39}
                                    if false {
                                        v41 := GenS6{F0: 127}
                                        var v42 int32 = 42
                                        _ = v41
                                        _ = v42
                                    } else {
                                        fmt.Println(v5)
                                        v43 := GenS1{F0: 42}
                                        _ = v43
                                    }
                                    _ = v40
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = ""
                                    }
                                }
                                _ = fi39
                            }
                            _ = fi38
                        }
                    }
                default:
                    v5 = v3
                }
            } else {
                if false {
                    fmt.Println(v3)
                    fmt.Println(v3)
                    if true {
                        if ((0.0 - -2.25) != 0.1) {
                            switch "" {
                            case "hello":
                                var v44 int16 = -32768
                                _ = v44
                            case "\n\t":
                                var v45 int8 = 0
                                _ = v45
                            case "a":
                                fmt.Println(v5)
                            default:
                                for fi46 := 0; fi46 < 20; fi46++ {
                                    if true {
                                        fmt.Println(v3)
                                    }
                                    v3 = 1
                                    for fi47 := 0; fi47 < 4; fi47++ {
                                        v3 -= (v3 | v5)
                                        _ = fi47
                                    }
                                    _ = fi46
                                }
                            }
                            delete(v2, -607)
                        }
                        for fi48 := 0; fi48 < 10; fi48++ {
                            if false {
                                var v49 uint16 = 255
                                v3 += (v5 + 0)
                                _ = v49
                            }
                            var v50 int32 = 377
                            v3 += (v5 << ((65535) & 63))
                            _ = fi48
                            _ = v50
                        }
                        var v51 bool = (2 < 2)
                        _ = v51
                    } else {
                        if true {
                            fmt.Println(v5, v5)
                            v52 := GenS7{F0: 65535, F1: -2147483648, F2: 2}
                            _ = v52
                        } else {
                            var v53 uint = 0
                            v54 := GenS3{F0: -1, F1: -1, F2: 48}
                            switch "" {
                            case "\n\t":
                                var v55 uint = 533
                                _ = v55
                            default:
                                var v56 uint32 = 2147483647
                                _ = v56
                            }
                            _ = v53
                            _ = v54
                        }
                    }
                }
                if true {
                    v57 := GenS2{F0: -1, F1: -122}
                    _ = v57
                }
            }
        }
    }
    var v58 uint16 = 65535
    if ((!(488 <= -637)) && ((3.14159 * float64(((3.14159 * -2.25) - -2.25))) >= (0.1 * 1e10))) {
        v59 := GenS8{F0: false, F1: (false || false), F2: 65535}
        if ((1 >= 176) && ((!true) || false)) {
            if ("foo bar" > "foo bar") {
                v5 += v5
                fmt.Println(v5)
                if false {
                    v58 += 42
                }
            } else {
                v60 := GenS4{F0: 1, F1: 255, F2: v58}
                _ = v60
            }
        } else {
            fmt.Println(v5)
            switch "foo bar" {
            case "":
                switch (("" + (("hello" + ("" + ("" + "foo bar"))) + "")) + "\n\t") {
                case "foo bar":
                    if true {
                        if len(v4) > 0 {
                            v4[0] = ("hello" + ("\n\t" + (("" + ("foo bar" + ("foo bar" + "hello"))) + "a")))
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        for fi61 := 0; fi61 < 5; fi61++ {
                            v62 := GenS9{F0: 199}
                            v63 := GenS3{F0: 1, F1: 2147483647, F2: 2}
                            v64 := GenS10{F0: uint32(v58)}
                            _ = fi61
                            _ = v62
                            _ = v63
                            _ = v64
                        }
                    }
                default:
                    for fi65 := 0; fi65 < 14; fi65++ {
                        if false {
                            var v66 uint8 = 128
                            var v67 float64 = 3.14159
                            _ = v66
                            _ = v67
                        } else {
                            for fi68 := 0; fi68 < 4; fi68++ {
                                v69 := GenS2{F0: -455, F1: -128}
                                _ = fi68
                                _ = v69
                            }
                            var v70 int32 = 1
                            if ((3.14159 - 3.14159) >= -2.25) {
                                fmt.Println(v5, v70, fi65)
                                v71 := GenS11{F0: 0, F1: (^(v5))}
                                v2[fi65] = true
                                _ = v71
                            } else {
                                var v72 float32 = (1e10 + (((((((((3.14159 * ((0.0 - 0.1) + (-2.25 * -2.25))) - 0.0) + (((((0.1 - 0.1) / -2.25) - float32(1e10)) - -2.25) / 0.1)) * 1e10) + 1.5) - (((-2.25 / (1e10 * ((float32(1e10) / 1e10) - 0.1))) - (1.5 + (-2.25 * 3.14159))) / 1.5)) + 0.0) - float32(1e10)) * (1.5 - float32(0.0))))
                                _ = v72
                            }
                            _ = v70
                        }
                        _ = fi65
                    }
                }
            case "\n\t":
                if false {
                    for fi73 := 0; fi73 < 8; fi73++ {
                        v5 = (v5 / ((v5) | 1))
                        _ = fi73
                    }
                    if (2147483648 > 2) {
                        delete(v1, 704)
                        v5 += (v3 / ((v5) | 1))
                        if ("hello" < (("\n\t" + "\n\t") + ("hello" + "hello"))) {
                            for fi74 := 0; fi74 < 6; fi74++ {
                                v75 := GenS12{F0: 128, F1: (^(v58)), F2: -2147483648}
                                var v76 uint64 = 18446744073709551615
                                _ = fi74
                                _ = v75
                                _ = v76
                            }
                            if true {
                                fmt.Println(v58, v58, v5)
                                var v77 int32 = 616
                                for fi78 := 0; fi78 < 2; fi78++ {
                                    switch (fi78 | fi78) {
                                    case 32767:
                                        var v79 uint = (^(v5))
                                        _ = v79
                                    case -2147483648:
                                        v5 = v5
                                    case 2147483647:
                                        fmt.Println(v5, v77, v5)
                                    default:
                                        if (v58 > 722) {
                                            if (0.1 != (0.1 * (0.0 + (1.5 - (3.14159 - 0.0))))) {
                                                fmt.Println(v5, v5)
                                                v3 = (v3 % (((v5 ^ v3)) | 1))
                                                v5 = v3
                                            }
                                            v80 := GenS9{F0: 202}
                                            for fi81 := 0; fi81 < 13; fi81++ {
                                                if true {
                                                    if (128 != 2) {
                                                        delete(v2, fi81)
                                                    } else {
                                                        var v82 int64 = 42
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            delete(v1, 2)
                                                        case "":
                                                            fmt.Println(v77, fi81, fi78)
                                                        case "\n\t":
                                                            _ = len(v2)
                                                        default:
                                                            var v83 uint64 = 938
                                                            _ = v83
                                                        }
                                                        fmt.Println(v3, v77)
                                                        _ = v82
                                                    }
                                                }
                                                _ = fi81
                                            }
                                            _ = v80
                                        }
                                    }
                                    var v84 float64 = -2.25
                                    _ = fi78
                                    _ = v84
                                }
                                _ = v77
                            } else {
                                v58 -= (^(v58))
                            }
                        } else {
                            switch -2147483648 {
                            case 1:
                                switch "a" {
                                case "":
                                    switch "" {
                                    case "a":
                                        _ = len(v1)
                                    }
                                case "\n\t":
                                    var v85 int = 32767
                                    _ = v85
                                case "hello":
                                    fmt.Println(v5, v3)
                                }
                            }
                            var v86 float64 = -2.25
                            _ = v86
                        }
                    }
                    switch -2147483648 {
                    case 127:
                        delete(v1, 65535)
                    case 32768:
                        var v87 int32 = -1
                        _ = v87
                    case 32767:
                        fmt.Println(v3, v58)
                    }
                } else {
                    v88 := GenS6{F0: 255}
                    _ = v88
                }
            case "a":
                var v89 string = ("a" + (("\n\t" + "foo bar") + "hello"))
                _ = v89
            }
        }
        _ = v59
    }
    fmt.Println(v5)
    var v90 uint = v3
    var v91 bool = true
    var v92 int16 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
    _ = v7
    _ = v90
    _ = v91
    _ = v92
}
