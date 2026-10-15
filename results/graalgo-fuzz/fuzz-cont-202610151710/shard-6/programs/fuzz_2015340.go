package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 int16; F1 uint64; F2 uint }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 bool; F1 int16 }
type GenS8 struct { F0 int8; F1 int32; F2 uint8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint64; F1 uint32; F2 bool }
type GenS11 struct { F0 int32; F1 uint32 }

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[uint32]int)
    var v3 uint16 = 188
    v4 := []string{}
    var v5 int64 = 0
    var v6 int = int(v3)
    var v7 string = "\n\t"
    for fi8 := 0; fi8 < 11; fi8++ {
        if (-60 == 100) {
            _ = len(v2)
            var v9 int8 = 100
            if ((-1 == 877) || (false && (42 < 45))) {
                if (false || (!false)) {
                    v6 = (v6 % ((v6) | 1))
                    switch (^(v6)) {
                    case 42:
                        fmt.Println(v3, v5, v1)
                    case 32768:
                        v10 := GenS1{F0: 0, F1: (v5 % ((v5) | 1))}
                        _ = v10
                    case -1:
                        for fi11 := 0; fi11 < 3; fi11++ {
                            var v12 int = fi11
                            v13 := v2[(v1 << ((2) & 63))]
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                    }
                } else {
                    fmt.Println(v3)
                }
                fmt.Println(fi8, v3, v5)
                var v14 int16 = 261
                _ = v14
            } else {
                v7 = v7
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                delete(v2, 4294967295)
            }
            _ = v9
        }
        _ = fi8
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        if true {
            for fi16 := 0; fi16 < 14; fi16++ {
                _ = len(v2)
                switch fi16 {
                case 4294967295:
                    var v17 uint8 = uint8(v1)
                    _ = v17
                case 9223372036854775807:
                    v18 := GenS2{F0: (v6 * fi15)}
                    _ = v18
                case 571:
                    v19 := GenS2{F0: (v6 << ((368) & 63))}
                    _ = v19
                default:
                    if false {
                        var v20 uint64 = 1
                        v21 := GenS1{F0: (v20 ^ 18446744073709551615), F1: v5}
                        _ = v20
                        _ = v21
                    } else {
                        switch ((((v7 + v7) + "foo bar") + v7) + ("foo bar" + ("a" + "hello"))) {
                        case "\n\t":
                            var v22 float32 = float32(1e10)
                            _ = v22
                        default:
                            v5 += (^(v5))
                        }
                        var v23 int = 726
                        _ = v23
                    }
                }
                _ = fi16
            }
        } else {
            for fi24 := 0; fi24 < 16; fi24++ {
                switch (fi24 << ((42) & 63)) {
                case -1:
                    v25 := GenS2{F0: 4294967295}
                    _ = v25
                default:
                    v1 = uint32(v3)
                }
                var v26 int32 = -2147483648
                var v27 uint64 = 846
                _ = fi24
                _ = v26
                _ = v27
            }
            var v28 string = v7
            _ = v28
        }
        _ = fi15
    }
    switch v6 {
    case 2147483647:
        v29 := v2[2147483647]
        _ = v29
    }
    switch "hello" {
    case "\n\t":
        v30 := GenS3{F0: 127, F1: 100}
        _ = v30
    case "hello":
        if false {
            for fi31 := 0; fi31 < 10; fi31++ {
                fmt.Println(v3)
                var v32 int8 = -60
                _ = fi31
                _ = v32
            }
            if (false || (-128 < -128)) {
                fmt.Println(v7, v5, v7)
                v5 = 1
            } else {
                var v33 float64 = (1e10 - 0.0)
                var v34 int = v6
                var v35 float64 = -2.25
                _ = v33
                _ = v34
                _ = v35
            }
            if len(v4) > 0 {
                v4[0] = (v7 + v7)
            }
        } else {
            for fi36 := 0; fi36 < 14; fi36++ {
                delete(v2, (v1 << ((295) & 63)))
                if len(v4) > 0 {
                    v4[0] = "a"
                }
                _ = fi36
            }
            v37 := GenS4{F0: 820, F1: 1, F2: 2}
            for fi38 := 0; fi38 < 7; fi38++ {
                for fi39 := 0; fi39 < 15; fi39++ {
                    switch v7 {
                    case "a":
                        v4 = append(v4, v7)
                    case "\n\t":
                        switch v7 {
                        case "hello":
                            var v40 int32 = 0
                            fmt.Println(v5)
                            _ = v40
                        case "foo bar":
                            if false {
                                v41 := GenS5{F0: (v5 * v5)}
                                v6 += (fi38 & v6)
                                if true {
                                    fmt.Println(v7, fi39, v6)
                                } else {
                                    delete(v2, (v1 - v1))
                                    var v42 uint8 = 128
                                    v43 := GenS2{F0: (^(fi39))}
                                    _ = v42
                                    _ = v43
                                }
                                _ = v41
                            } else {
                                v7 = v7
                                v2[v1] = (v6 / ((-9223372036854775808) | 1))
                            }
                        case "":
                            if true {
                                v44 := GenS3{F0: -128, F1: 127}
                                switch int(v5) {
                                case 128:
                                    var v45 int64 = v5
                                    _ = v45
                                }
                                _ = v44
                            } else {
                                fmt.Println(v7)
                                switch -2147483648 {
                                case 42:
                                    v46 := GenS6{F0: v6}
                                    _ = v46
                                case 127:
                                    var v47 int64 = (v5 & (v5 << ((65535) & 63)))
                                    v48 := GenS3{F0: 44, F1: 127}
                                    _ = v47
                                    _ = v48
                                case -271:
                                    v1 = (^(v1))
                                default:
                                    if true {
                                        fmt.Println(v6, fi39, v7)
                                        var v49 uint = 42
                                        if (((524 < -1) || (float32(-2.25) > 0.0)) && ((32767 == 1000) && ((!(true || true)) || (v1 <= (v1 - (v1 * (v1 * v1))))))) {
                                            v50 := GenS6{F0: (v6 << (((v49 - 91)) & 63))}
                                            var v51 bool = (v7 >= "\n\t")
                                            _ = v50
                                            _ = v51
                                        } else {
                                            switch (("\n\t" + v7) + v7) {
                                            case "hello":
                                                if (int64(v3) < v5) {
                                                    v3 -= v3
                                                    var v52 uint64 = 69
                                                    _ = v52
                                                }
                                            case "foo bar":
                                                switch fi39 {
                                                case 32767:
                                                    fmt.Println(v49, fi38)
                                                }
                                            }
                                            if true {
                                                if len(v4) > 0 {
                                                    v4[0] = ""
                                                }
                                                if false {
                                                    fmt.Println(v5, v1, v3)
                                                    if len(v4) > 0 {
                                                        v4[0] = (v7 + "a")
                                                    }
                                                    v3 += (v3 >> ((18446744073709551615) & 63))
                                                } else {
                                                    var v53 int8 = 65
                                                    _ = v53
                                                }
                                                _ = len(v2)
                                                var v54 int = (v6 ^ (fi39 << ((18446744073709551615) & 63)))
                                                var v55 uint = (v49 * (v49 / ((v49) | 1)))
                                                _ = v54
                                                _ = v55
                                            }
                                            v56 := GenS7{F0: (int8(v1) <= -1), F1: 1000}
                                            _ = v56
                                        }
                                        _ = v49
                                    } else {
                                        _ = len(v2)
                                    }
                                }
                            }
                        default:
                            v6 = fi39
                        }
                    default:
                        var v57 uint16 = (^(v3))
                        v58 := GenS8{F0: 0, F1: -2147483648, F2: 63}
                        v4 = append(v4, "")
                        _ = v57
                        _ = v58
                    }
                    switch (v6 / ((fi39) | 1)) {
                    case -2147483648:
                        switch -347 {
                        case 86:
                            v6 -= -2147483648
                        case 9223372036854775807:
                            switch (^(fi38)) {
                            case 4294967295:
                                v5 += -1
                            case -1:
                                var v59 int = (fi39 % ((fi39) | 1))
                                _ = v59
                            case 32768:
                                var v60 float32 = 1.5
                                _ = v60
                            default:
                                switch "" {
                                case "hello":
                                    switch (v6 << ((2) & 63)) {
                                    case 537:
                                        if (42 <= 637) {
                                            var v61 bool = true
                                            var v62 float32 = ((((1.5 - (1.5 / ((1.5 * float32(((1.5 * (-2.25 * 1.5)) - (((float64((3.14159 - 0.0)) - 0.0) * 3.14159) + float64(((-2.25 - float32(1.5)) * 0.1)))))) - (-2.25 / (1.5 + 3.14159))))) - (1.5 - ((((0.0 + 1e10) + (1e10 / (3.14159 / -2.25))) / 1.5) / -2.25))) * 1e10) / 1e10)
                                            _ = len(v2)
                                            _ = v61
                                            _ = v62
                                        } else {
                                            var v63 uint = 65535
                                            delete(v2, 810)
                                            _ = v63
                                        }
                                    case -112:
                                        _ = len(v2)
                                    case 36:
                                        fmt.Println(v7, fi38, v3)
                                    default:
                                        v5 = -9223372036854775808
                                    }
                                case "foo bar":
                                    fmt.Println(v7, fi38, v5)
                                default:
                                    switch (^(v6)) {
                                    case 32768:
                                        fmt.Println(v1, v3, v1)
                                    case 0:
                                        if (false || (!true)) {
                                            if (((false || (fi39 <= 1)) || (0.1 <= 0.1)) && false) {
                                                var v64 int = fi38
                                                v65 := GenS2{F0: v64}
                                                var v66 int16 = 1000
                                                _ = v64
                                                _ = v65
                                                _ = v66
                                            } else {
                                                v7 = ""
                                            }
                                        }
                                    default:
                                        fmt.Println(v3)
                                    }
                                }
                            }
                        }
                    }
                    v6 += (-(v6))
                    _ = fi39
                }
                v67 := GenS9{F0: 1}
                _ = fi38
                _ = v67
            }
            _ = v37
        }
    default:
        var v68 int64 = v5
        _ = v68
    }
    v69 := GenS10{F0: 49, F1: v1, F2: true}
    var v70 int64 = v5
    if (v7 >= "") {
        if ((448 >= 220) || true) {
            v3 = (v3 * (v3 | (^(v3))))
            v70 += v70
            switch "foo bar" {
            case "a":
                v71 := GenS3{F0: 0, F1: int8(v6)}
                _ = v71
            case "hello":
                for fi72 := 0; fi72 < 1; fi72++ {
                    v3 = v3
                    _ = fi72
                }
            }
        } else {
            if ((v5 + v5) == v5) {
                var v73 bool = false
                _ = v73
            } else {
                if len(v4) > 0 {
                    v4[0] = v7
                }
                switch v6 {
                case 4294967295:
                    delete(v2, (^(v1)))
                default:
                    for fi74 := 0; fi74 < 20; fi74++ {
                        switch 32768 {
                        case 0:
                            v5 -= -1
                        case -121:
                            for fi75 := 0; fi75 < 15; fi75++ {
                                var v76 float64 = (3.14159 + 3.14159)
                                _ = len(v4)
                                fmt.Println(v1, v3, fi75)
                                _ = fi75
                                _ = v76
                            }
                        case 4294967295:
                            var v77 uint8 = 125
                            _ = v77
                        default:
                            switch v7 {
                            case "foo bar":
                                v78 := GenS3{F0: 127, F1: 42}
                                _ = v78
                            case "a":
                                v5 = v5
                            case "hello":
                                if false {
                                    _ = len(v2)
                                    if (51 > 42) {
                                        var v79 uint8 = 42
                                        _ = v79
                                    }
                                    for fi80 := 0; fi80 < 17; fi80++ {
                                        switch "a" {
                                        case "hello":
                                            fmt.Println(fi74, v5, v6)
                                        case "\n\t":
                                            var v81 string = "hello"
                                            _ = v81
                                        }
                                        if (true || (127 <= fi80)) {
                                            v6 = (v6 ^ v6)
                                            var v82 float64 = 0.1
                                            var v83 float64 = v82
                                            _ = v82
                                            _ = v83
                                        }
                                        v3 += (v3 ^ (^(v3)))
                                        _ = fi80
                                    }
                                } else {
                                    if true {
                                        if (!(-203 <= 0)) {
                                            for fi84 := 0; fi84 < 15; fi84++ {
                                                if (127 == 255) {
                                                    fmt.Println(v3, v6, v6)
                                                    var v85 uint = 255
                                                    _ = v85
                                                }
                                                v3 = (v3 % ((v3) | 1))
                                                delete(v2, v1)
                                                _ = fi84
                                            }
                                            v6 = fi74
                                        } else {
                                            delete(v2, 440)
                                        }
                                        fmt.Println(v7)
                                    } else {
                                        for fi86 := 0; fi86 < 8; fi86++ {
                                            fmt.Println(v6, fi74, fi86)
                                            v2[v1] = (fi74 | (fi74 ^ v6))
                                            _ = fi86
                                        }
                                    }
                                    v3 -= v3
                                }
                            }
                        }
                        v3 = 1
                        v5 = v5
                        _ = fi74
                    }
                }
            }
            v87 := GenS2{F0: v6}
            _ = v87
        }
    } else {
        v88 := GenS11{F0: 58, F1: (v1 - v1)}
        if ((true || false) || (true && (float64(float32((0.0 + -2.25))) >= 0.0))) {
            v7 = v7
        }
        fmt.Println(v6, v5, v6)
        _ = v88
    }
    var v89 uint64 = 2
    v1 = v1
    var v90 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v70:", v70)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v7
    _ = v70
    _ = v89
    _ = v90
}
