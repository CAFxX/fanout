package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 string }
type GenS2 struct { F0 uint8; F1 uint32; F2 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 int64; F2 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int; F1 uint; F2 uint16 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 string; F1 int32 }
type GenS10 struct { F0 string; F1 uint8; F2 int64 }
type GenS11 struct { F0 int8; F1 int; F2 uint16 }
type GenS12 struct { F0 bool; F1 uint; F2 int16 }
func genFn0() string {
    var v7 uint16 = 933
    v7 += v7
    var v8 uint = 0
    _ = v7
    _ = v8
    return "a"
}

func main() {
    var v1 int = 880
    var v2 int64 = -9223372036854775808
    var v3 int64 = -1
    var v4 string = ""
    var v5 int = v1
    v6 := []int{32768, v1}
    var v9 float64 = 1e10
    fmt.Println(v5, v5)
    if true {
        v10 := GenS1{F0: 1, F1: -2147483648, F2: v4}
        switch (v4 + "\n\t") {
        case "":
            for fi11 := 0; fi11 < 14; fi11++ {
                for fi12 := 0; fi12 < 18; fi12++ {
                    switch v4 {
                    case "hello":
                        if (true || false) {
                            _ = len(v6)
                            v13 := GenS1{F0: 278, F1: 719, F2: v4}
                            switch v4 {
                            case "a":
                                v3 += v3
                            case "":
                                var v14 int64 = v3
                                _ = v14
                            case "foo bar":
                                switch fi12 {
                                case -9223372036854775808:
                                    v2 = v2
                                case 2147483647:
                                    if true {
                                        fmt.Println(v1)
                                        if len(v6) > 0 {
                                            v6[0] = (v5 << ((0) & 63))
                                        }
                                    } else {
                                        switch 1 {
                                        case 9223372036854775807:
                                            switch v5 {
                                            case 1:
                                                if len(v6) > 0 {
                                                    v6[0] = (v5 + (fi11 << ((651) & 63)))
                                                }
                                            case 127:
                                                v3 += (v2 / (((v3 % ((42) | 1))) | 1))
                                            }
                                        case 42:
                                            v3 += (v2 * 42)
                                        case 0:
                                            fmt.Println(v9, v2)
                                        default:
                                            v2 -= (v3 ^ v2)
                                        }
                                        v3 = v2
                                        fmt.Println(v5, v2, v9)
                                    }
                                case 32767:
                                    v15 := GenS1{F0: 215, F1: -2147483648, F2: ""}
                                    _ = v15
                                }
                            }
                            _ = v13
                        }
                    case "foo bar":
                        switch v4 {
                        case "a":
                            var v16 uint32 = 913
                            _ = v16
                        case "":
                            var v17 string = "foo bar"
                            _ = v17
                        case "\n\t":
                            _ = len(v6)
                        }
                    case "\n\t":
                        switch "\n\t" {
                        case "a":
                            fmt.Println(v4, v3, v5)
                        case "":
                            if true {
                                if len(v6) > 0 {
                                    v6[0] = (fi12 * (fi12 << ((1) & 63)))
                                }
                                var v18 int = v1
                                if (2 != 4294967295) {
                                    if (0 != (fi11 << ((65535) & 63))) {
                                        switch v4 {
                                        case "a":
                                            var v19 int32 = 1
                                            _ = v19
                                        case "hello":
                                            switch v1 {
                                            case -2147483648:
                                                v2 = v3
                                                _ = len(v6)
                                                var v20 float32 = 0.0
                                                _ = v20
                                            default:
                                                v21 := GenS2{F0: 128, F1: 2147483647, F2: ((-230 == 42) && false)}
                                                _ = v21
                                            }
                                        case "\n\t":
                                            switch 127 {
                                            case -9223372036854775808:
                                                switch (fi12 % ((2147483647) | 1)) {
                                                case 127:
                                                    v4 = ((((("\n\t" + v4) + v4) + "") + v4) + v4)
                                                case 490:
                                                    var v22 uint64 = 42
                                                    _ = v22
                                                default:
                                                    var v23 string = ""
                                                    _ = v23
                                                }
                                            default:
                                                if (!false) {
                                                    switch 0 {
                                                    case 128:
                                                        if (1.5 < v9) {
                                                            if false {
                                                                if false {
                                                                    fmt.Println(v2, v9, v4)
                                                                }
                                                                if ((v2 & v2) <= v2) {
                                                                    v3 += (v3 + (v3 % ((-171) | 1)))
                                                                    if len(v6) >= 2 {
                                                                        v6 = v6[1:len(v6)]
                                                                    }
                                                                } else {
                                                                    var v24 int = (^(fi12))
                                                                    _ = v24
                                                                }
                                                            } else {
                                                                switch "hello" {
                                                                case "hello":
                                                                    fmt.Println(v18, fi12, v4)
                                                                case "a":
                                                                    _ = len(v6)
                                                                default:
                                                                    switch (fi12 | (fi12 / ((-821) | 1))) {
                                                                    case 0:
                                                                        v25 := GenS3{F0: true}
                                                                        _ = v25
                                                                    }
                                                                }
                                                                _ = len(v6)
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                        default:
                                            if (false && true) {
                                                var v26 int8 = -26
                                                _ = v26
                                            }
                                        }
                                        v27 := GenS2{F0: 128, F1: 456, F2: (42 != 97)}
                                        _ = v27
                                    } else {
                                        var v28 int = 127
                                        v4 = v4
                                        _ = v28
                                    }
                                    switch (v4 + "a") {
                                    case "":
                                        v29 := GenS2{F0: 128, F1: 2, F2: false}
                                        _ = v29
                                    case "hello":
                                        if true {
                                            v2 = v3
                                            switch v4 {
                                            case "hello":
                                                if (!(v2 > v3)) {
                                                    var v30 int32 = 42
                                                    var v31 int64 = (v2 + 0)
                                                    _ = v30
                                                    _ = v31
                                                } else {
                                                    fmt.Println(v1)
                                                    v32 := GenS4{F0: v3, F1: -907, F2: -1}
                                                    v2 = v3
                                                    _ = v32
                                                }
                                            }
                                        }
                                    }
                                } else {
                                    var v33 float32 = (float32(((1.5 * (v9 * ((v9 - -2.25) / (v9 / (v9 / -2.25))))) - float64(0.1))) / (1.5 + (1e10 - ((((-2.25 + 0.0) - 1e10) / 1e10) - -2.25))))
                                    fmt.Println(v1)
                                    v5 -= 32768
                                    _ = v33
                                }
                                var v34 int32 = 203
                                var v35 float64 = (1.5 / v9)
                                _ = v18
                                _ = v34
                                _ = v35
                            }
                        }
                    }
                    fmt.Println(v5, v1, v5)
                    var v36 uint = 2
                    _ = fi12
                    _ = v36
                }
                _ = fi11
            }
        case "\n\t":
            if true {
                var v37 uint16 = 958
                var v38 uint32 = 0
                _ = v37
                _ = v38
            } else {
                v3 = (v2 * (v2 >> ((240) & 63)))
                v39 := GenS3{F0: (687 <= v3)}
                _ = v39
            }
        default:
            for fi40 := 0; fi40 < 18; fi40++ {
                if len(v6) > 0 {
                    v6[0] = (fi40 % ((1) | 1))
                }
                v41 := GenS4{F0: (v2 % (((v2 << ((1) & 63))) | 1)), F1: (v3 | (v3 | -9223372036854775808)), F2: (v2 & 268)}
                _ = fi40
                _ = v41
            }
        }
        _ = v10
    } else {
        v1 = v1
    }
    v4 = v4
    if (!true) {
        for fi42 := 0; fi42 < 12; fi42++ {
            v3 += (v3 * (v2 % (((v3 % (((v3 ^ v2)) | 1))) | 1)))
            v3 += v2
            _ = fi42
        }
        var v43 int16 = int16(v1)
        for fi44 := 0; fi44 < 7; fi44++ {
            v45 := GenS5{F0: 2}
            _ = fi44
            _ = v45
        }
        _ = v43
    } else {
        if (255 != 0) {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        } else {
            v46 := GenS6{F0: true}
            v3 = (v2 / ((v3) | 1))
            _ = v46
        }
    }
    for fi47 := 0; fi47 < 19; fi47++ {
        v5 = v5
        for fi48 := 0; fi48 < 6; fi48++ {
            switch ((v4 + (v4 + ((v4 + v4) + v4))) + v4) {
            case "foo bar":
                var v49 float32 = 1e10
                _ = v49
            case "":
                v50 := GenS5{F0: 126}
                _ = v50
            default:
                _ = len(v6)
            }
            v1 -= (v1 << ((1) & 63))
            _ = fi48
        }
        _ = fi47
    }
    for fi51 := 0; fi51 < 19; fi51++ {
        var v52 uint32 = 332
        _ = fi51
        _ = v52
    }
    switch v4 {
    case "foo bar":
        switch v1 {
        case -1:
            var v53 int32 = 2147483647
            _ = v53
        case 42:
            switch "hello" {
            case "hello":
                var v54 int = (v1 - (v1 / ((v1) | 1)))
                _ = v54
            }
        case -2147483648:
            var v55 int64 = (v3 / ((v2) | 1))
            _ = v55
        default:
            var v56 uint16 = 42
            _ = v56
        }
    case "\n\t":
        var v57 float32 = 1e10
        _ = v57
    case "hello":
        var v58 uint32 = 2147483647
        _ = v58
    default:
        if ((963 < 18446744073709551615) && ((1.5 - 1.5) != 1e10)) {
            var v59 int8 = 1
            v59 += v59
            _ = v59
        }
    }
    fmt.Println(v1, v4, v5)
    for fi60 := 0; fi60 < 9; fi60++ {
        fmt.Println(v1, fi60)
        if false {
            v5 -= (v5 / ((v1) | 1))
            v5 += v1
        } else {
            switch -9223372036854775808 {
            case -582:
                switch "hello" {
                case "":
                    if false {
                        var v61 uint = 255
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        fmt.Println(v4)
                        _ = v61
                    }
                case "\n\t":
                    v62 := GenS7{F0: (v5 + 32768), F1: 636, F2: 0}
                    _ = v62
                default:
                    v5 = (v5 + (^(v5)))
                }
            case 9223372036854775807:
                if (!true) {
                    if true {
                        v63 := GenS8{F0: 100}
                        v1 -= (fi60 & fi60)
                        var v64 int16 = 0
                        _ = v63
                        _ = v64
                    }
                }
            case 116:
                v65 := GenS9{F0: (v4 + "foo bar"), F1: 2147483647}
                _ = v65
            }
            v2 = 9223372036854775807
            v1 = fi60
        }
        _ = fi60
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if true {
        v66 := GenS10{F0: (v4 + v4), F1: 124, F2: 0}
        v67 := GenS8{F0: -49}
        v68 := GenS3{F0: (true || (v1 != (v5 % ((381) | 1))))}
        _ = v66
        _ = v67
        _ = v68
    } else {
        if true {
            if false {
                fmt.Println(v3, v9, v2)
                var v69 float32 = -2.25
                v70 := GenS1{F0: 18446744073709551615, F1: -196, F2: v4}
                _ = v69
                _ = v70
            }
        } else {
            v71 := GenS11{F0: -1, F1: (-(v1)), F2: 2}
            _ = v71
        }
        fmt.Println(v4, v2)
        v6 = append(v6, 32767)
    }
    if (-2147483648 != 9223372036854775807) {
        if (!true) {
            v72 := GenS4{F0: (v3 * v2), F1: v2, F2: 0}
            _ = v72
        }
        v9 = (1e10 - v9)
    } else {
        var v73 string = v4
        for fi74 := 0; fi74 < 8; fi74++ {
            var v75 int8 = 127
            _ = fi74
            _ = v75
        }
        _ = v73
    }
    var v76 uint32 = 65535
    if false {
        _ = len(v6)
        var v77 int64 = v3
        fmt.Println(v2, v2, v9)
        _ = v77
    }
    var v78 int16 = 42
    if len(v6) > 0 {
        v6[0] = v1
    }
    if true {
        var v79 uint8 = 42
        for fi80 := 0; fi80 < 4; fi80++ {
            if ((1e10 / 3.14159) > 1e10) {
                var v81 int = (v1 * (fi80 >> ((4294967295) & 63)))
                for fi82 := 0; fi82 < 2; fi82++ {
                    v83 := GenS9{F0: v4, F1: 1}
                    var v84 int32 = 0
                    _ = fi82
                    _ = v83
                    _ = v84
                }
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v81
            }
            v85 := GenS6{F0: false}
            fmt.Println(v9, v5, v79)
            _ = fi80
            _ = v85
        }
        var v86 int32 = -48
        _ = v79
        _ = v86
    } else {
        for fi87 := 0; fi87 < 2; fi87++ {
            v3 -= 557
            if true {
                if true {
                    switch "hello" {
                    case "foo bar":
                        var v88 float32 = (1e10 + (1e10 - (1e10 - 0.1)))
                        _ = v88
                    case "hello":
                        var v89 int32 = 952
                        _ = v89
                    }
                    var v90 uint8 = uint8(v76)
                    v3 = v2
                    _ = v90
                }
                _ = len(v6)
            } else {
                v76 = v76
                switch "hello" {
                case "foo bar":
                    var v91 uint8 = uint8(v76)
                    _ = v91
                case "hello":
                    v92 := GenS5{F0: 110}
                    _ = v92
                case "":
                    v93 := GenS12{F0: true, F1: 530, F2: (^(v78))}
                    _ = v93
                default:
                    if len(v6) > 0 {
                        v6[0] = v1
                    }
                }
                for fi94 := 0; fi94 < 2; fi94++ {
                    v76 = v76
                    v6 = append(v6, 770)
                    _ = fi94
                }
            }
            _ = fi87
        }
        if (!false) {
            for fi95 := 0; fi95 < 15; fi95++ {
                if len(v6) > 0 {
                    v6[0] = (v1 + (fi95 / (((v1 * (fi95 ^ (v5 / (((v5 - fi95)) | 1))))) | 1)))
                }
                _ = fi95
            }
            v76 = (v76 & v76)
        } else {
            if len(v6) > 0 {
                v6[0] = (v1 << ((1) & 63))
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v76:", v76)
    fmt.Println("v78:", v78)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
    _ = v78
    _ = v9
}
