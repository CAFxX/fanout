package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32 }

func main() {
    var v1 int8 = -1
    v2 := make(map[int]bool)
    v3 := make(map[string]string)
    v4 := []uint32{}
    var v5 string = "hello"
    v6 := make(map[int]string)
    v5 = "\n\t"
    _ = len(v4)
    var v7 uint64 = 2
    if len(v4) > 0 {
        v4[0] = 21
    }
    var v8 int = 9223372036854775807
    if false {
        var v9 bool = false
        var v10 int8 = v1
        v9 = ("foo bar" <= ("\n\t" + v5))
        _ = v10
        _ = v9
    }
    switch (v8 ^ v8) {
    case -933:
        v1 -= (v1 * (v1 << ((65535) & 63)))
    }
    var v11 int32 = 1
    v12 := GenS1{F0: 4294967295}
    var v13 string = ("a" + "")
    switch v8 {
    case -619:
        _ = len(v2)
    case 42:
        v14 := GenS2{F0: 2, F1: (380 > (v8 + 42))}
        _ = v14
    case 1:
        v11 = (v11 ^ 624)
    }
    var v15 int = v8
    v1 -= -1
    if true {
        var v16 int = v15
        _ = v16
    } else {
        v17 := GenS2{F0: 42, F1: ((v11 << ((479) & 63)) > (v11 | 42))}
        var v18 uint32 = 4294967295
        _ = v17
        _ = v18
    }
    switch v8 {
    case 839:
        var v19 float32 = -2.25
        _ = v19
    default:
        fmt.Println(v11, v7, v11)
    }
    v2[v15] = (!(false && (0.0 >= 1.5)))
    if true {
        if (1 >= v15) {
            for fi20 := 0; fi20 < 1; fi20++ {
                if ((float64(((0.0 + (0.0 * float32(((0.0 * (3.14159 - 0.0)) * ((3.14159 - 1e10) / -2.25))))) / 1.0)) + 0.1) != (0.1 / float64(((-2.25 * 1.5) + ((float32(-2.25) + 1e10) * 1.5))))) {
                    v11 += (v11 % ((v11) | 1))
                    var v21 uint64 = (v7 + (v7 & v7))
                    var v22 int8 = 39
                    _ = v21
                    _ = v22
                } else {
                    var v23 float64 = float64(1e10)
                    _ = v23
                }
                _ = fi20
            }
            for fi24 := 0; fi24 < 8; fi24++ {
                var v25 int64 = 9223372036854775807
                v26 := GenS2{F0: 2, F1: false}
                _ = fi24
                _ = v25
                _ = v26
            }
        }
        if (false && (9223372036854775807 <= v15)) {
            var v27 uint16 = 42
            if (!true) {
                switch "hello" {
                case "":
                    _ = len(v4)
                case "a":
                    fmt.Println(v11)
                default:
                    for fi28 := 0; fi28 < 4; fi28++ {
                        v15 = v15
                        _ = fi28
                    }
                }
            } else {
                fmt.Println(v1, v27, v27)
            }
            _ = v27
        } else {
            v29 := GenS3{F0: v11, F1: false}
            var v30 int32 = (-(v11))
            switch v15 {
            case -9223372036854775808:
                v4 = append(v4, 432)
            case -231:
                switch (v15 - int(v1)) {
                case -1:
                    if (172 <= 249) {
                        v30 = (^(v30))
                        if false {
                            fmt.Println(v13, v13)
                        } else {
                            var v31 int64 = 853
                            _ = v31
                        }
                    }
                case 128:
                    v4 = append(v4, 1)
                default:
                    _ = len(v2)
                }
            case 1:
                switch v5 {
                case "foo bar":
                    var v32 bool = (v15 < (v15 >> ((65535) & 63)))
                    _ = v32
                case "\n\t":
                    for fi33 := 0; fi33 < 15; fi33++ {
                        for fi34 := 0; fi34 < 4; fi34++ {
                            v35 := GenS2{F0: 2, F1: (false || true)}
                            if ((0.0 + -2.25) < 0.1) {
                                v30 += (v11 ^ (v11 | (v30 & v11)))
                                if ((!((!(42 < 255)) && false)) && true) {
                                    v11 -= (v11 - 609)
                                    var v36 int64 = 811
                                    switch 687 {
                                    case -2147483648:
                                        v37 := v3[(v5 + v5)]
                                        _ = v37
                                    default:
                                        v11 -= v11
                                    }
                                    var v38 uint64 = 42
                                    _ = v36
                                    _ = v38
                                }
                            }
                            v5 = "a"
                            _ = fi34
                            _ = v35
                        }
                        if (!(!(false && true))) {
                            v11 += (v11 >> ((18446744073709551615) & 63))
                        } else {
                            v39 := GenS4{F0: "a"}
                            if false {
                                var v40 uint = 65535
                                _ = v40
                            } else {
                                var v41 int8 = v1
                                switch (v8 & fi33) {
                                case 1:
                                    v11 = (v11 + v11)
                                case 4294967295:
                                    for fi42 := 0; fi42 < 6; fi42++ {
                                        switch ("hello" + ((v13 + "foo bar") + v5)) {
                                        case "a":
                                            var v43 bool = false
                                            _ = v43
                                        case "":
                                            var v44 string = v13
                                            _ = v44
                                        }
                                        if (!(((0.0 + -2.25) <= 3.14159) && (("a" + v13) > v13))) {
                                            switch ("\n\t" + v13) {
                                            case "hello":
                                                switch v8 {
                                                case -350:
                                                    v8 -= (-(fi33))
                                                default:
                                                    var v45 float64 = ((3.14159 + (1.5 / 0.1)) - 3.14159)
                                                    _ = v45
                                                }
                                            default:
                                                if false {
                                                    v46 := GenS3{F0: 2147483647, F1: false}
                                                    _ = v46
                                                } else {
                                                    var v47 int8 = 42
                                                    _ = v47
                                                }
                                            }
                                            if (!(0.0 == -2.25)) {
                                                v15 += v8
                                            } else {
                                                var v48 uint = 1
                                                v49 := GenS3{F0: (-(v11)), F1: (-2.25 > (3.14159 + (1.5 / (1e10 * -2.25))))}
                                                _ = v48
                                                _ = v49
                                            }
                                            v50 := GenS3{F0: v11, F1: (1.5 >= -2.25)}
                                            _ = v50
                                        }
                                        v7 += 42
                                        v51 := GenS5{F0: 255}
                                        _ = fi42
                                        _ = v51
                                    }
                                case 42:
                                    delete(v6, (fi33 | v15))
                                }
                                _ = v41
                            }
                            var v52 float32 = 3.14159
                            _ = v39
                            _ = v52
                        }
                        for fi53 := 0; fi53 < 4; fi53++ {
                            var v54 int32 = (v30 - v30)
                            _ = fi53
                            _ = v54
                        }
                        _ = fi33
                    }
                case "":
                    var v55 int32 = v11
                    _ = v55
                default:
                    for fi56 := 0; fi56 < 14; fi56++ {
                        var v57 uint8 = 127
                        var v58 int32 = v11
                        v6[fi56] = v5
                        _ = fi56
                        _ = v57
                        _ = v58
                    }
                }
            default:
                if (-9223372036854775808 == -561) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    if true {
                        switch "a" {
                        case "\n\t":
                            v59 := GenS4{F0: (v13 + "")}
                            _ = v59
                        case "foo bar":
                            v60 := v3["\n\t"]
                            _ = v60
                        default:
                            v4 = append(v4, uint32(v7))
                        }
                        for fi61 := 0; fi61 < 8; fi61++ {
                            if true {
                                for fi62 := 0; fi62 < 5; fi62++ {
                                    if (333 != 1) {
                                        delete(v2, v15)
                                        switch v15 {
                                        case 1:
                                            if false {
                                                fmt.Println(fi61, v8, v7)
                                            } else {
                                                fmt.Println(v7)
                                                var v63 int32 = 733
                                                _ = len(v4)
                                                _ = v63
                                            }
                                        case 2147483647:
                                            fmt.Println(v11, v11, v7)
                                        }
                                    }
                                    v11 -= v11
                                    _ = fi62
                                }
                                v13 = v5
                            } else {
                                switch 1 {
                                case -483:
                                    _ = len(v4)
                                case 0:
                                    if true {
                                        fmt.Println(v7, v1)
                                        if true {
                                            fmt.Println(fi61)
                                        }
                                    }
                                case 461:
                                    if len(v4) > 0 {
                                        v4[0] = 4294967295
                                    }
                                default:
                                    v64 := GenS5{F0: 127}
                                    _ = v64
                                }
                                if true {
                                    var v65 uint8 = 211
                                    _ = v65
                                } else {
                                    v66 := v3[v5]
                                    _ = v66
                                }
                            }
                            v7 = (v7 % (((v7 % ((v7) | 1))) | 1))
                            _ = fi61
                        }
                        if (true && false) {
                            v13 = v5
                            v5 = v5
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        } else {
                            var v67 bool = true
                            v68 := GenS6{F0: (v11 - (v11 ^ (v11 >> ((uint(v7)) & 63))))}
                            if (128 != 0) {
                                v4 = append(v4, 4294967295)
                                v67 = v67
                                if (v8 != (v8 + v8)) {
                                    v67 = v67
                                    v1 = (v1 & v1)
                                } else {
                                    if ((v1 >> ((2) & 63)) >= 127) {
                                        v5 = v5
                                        for fi69 := 0; fi69 < 1; fi69++ {
                                            v11 -= v11
                                            _ = len(v4)
                                            if v67 {
                                                fmt.Println(fi69)
                                                v11 -= v30
                                                fmt.Println(v30, v13)
                                            } else {
                                                if v67 {
                                                    v70 := GenS1{F0: 255}
                                                    v71 := GenS1{F0: 120}
                                                    _ = v70
                                                    _ = v71
                                                }
                                            }
                                            _ = fi69
                                        }
                                    } else {
                                        var v72 int64 = -779
                                        _ = v72
                                    }
                                    fmt.Println(v1)
                                    if false {
                                        for fi73 := 0; fi73 < 4; fi73++ {
                                            for fi74 := 0; fi74 < 16; fi74++ {
                                                var v75 uint64 = (v7 - v7)
                                                _ = fi74
                                                _ = v75
                                            }
                                            _ = fi73
                                        }
                                        var v76 int32 = v11
                                        _ = v76
                                    } else {
                                        v4 = append(v4, 2)
                                        v77 := v6[v15]
                                        _ = v77
                                    }
                                }
                            }
                            _ = v67
                            _ = v68
                        }
                    }
                    v78 := GenS2{F0: 2, F1: false}
                    _ = v78
                } else {
                    fmt.Println(v7, v1)
                }
            }
            _ = v29
            _ = v30
        }
        var v79 float32 = 1e10
        _ = v79
    }
    for fi80 := 0; fi80 < 9; fi80++ {
        v15 = (v8 >> ((65535) & 63))
        _ = fi80
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
