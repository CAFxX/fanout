package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int32 }
type GenS2 struct { F0 int; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool; F1 int; F2 int8 }
type GenS8 struct { F0 string; F1 int64; F2 int8 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 bool; F1 uint16 }

func main() {
    v1 := []uint32{2147483648, 476, 2147483647}
    var v2 float32 = (((float32(float64(((float32(0.0) * 1.5) + -2.25))) + 0.1) * 3.14159) - -2.25)
    var v3 string = ""
    var v4 int64 = -1
    var v5 int64 = 42
    var v6 int8 = 1
    var v7 int16 = 574
    if ((1 < 0) && (!true)) {
        fmt.Println(v4, v3, v3)
        v8 := GenS1{F0: "\n\t", F1: "hello", F2: 2147483647}
        _ = v8
    } else {
        v3 = (v3 + "a")
        if false {
            v5 = -9223372036854775808
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v5, v6)
        } else {
            var v9 float64 = -2.25
            var v10 int32 = -495
            _ = v10
            _ = v9
        }
        v1 = append(v1, 65535)
    }
    v3 = v3
    _ = len(v1)
    v3 = v3
    if (true && (530 >= 834)) {
        if len(v1) > 0 {
            v1[0] = 760
        }
    } else {
        v11 := GenS2{F0: int(v6), F1: 1, F2: 42}
        var v12 int8 = (v6 ^ v6)
        if (v2 <= 0.1) {
            if ((!true) && false) {
                var v13 uint = 917
                var v14 int32 = int32(v13)
                _ = v13
                _ = v14
            } else {
                var v15 string = "\n\t"
                for fi16 := 0; fi16 < 20; fi16++ {
                    fmt.Println(v12, fi16, v6)
                    _ = fi16
                }
                v4 -= v5
                _ = v15
            }
            var v17 int = 1
            v18 := GenS3{F0: 18446744073709551615, F1: false}
            _ = v17
            _ = v18
        } else {
            v12 = v6
        }
        _ = v11
        _ = v12
    }
    if false {
        if (!(!true)) {
            if false {
                v6 = (v6 * int8(v4))
                v6 = (v6 * v6)
                var v19 int16 = -1
                _ = v19
            }
            for fi20 := 0; fi20 < 13; fi20++ {
                v21 := GenS1{F0: "hello", F1: v3, F2: 42}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi20
                _ = v21
            }
            fmt.Println(v3)
        }
        if true {
            if (false && ((v6 > 38) && false)) {
                var v22 int = -9223372036854775808
                v1 = append(v1, 4294967295)
                v23 := GenS4{F0: v6}
                _ = v22
                _ = v23
            }
            v7 = v7
            var v24 int64 = (v4 ^ (v5 ^ (v4 >> ((0) & 63))))
            _ = v24
        }
        for fi25 := 0; fi25 < 12; fi25++ {
            for fi26 := 0; fi26 < 20; fi26++ {
                if true {
                    switch "a" {
                    case "hello":
                        var v27 float64 = 3.14159
                        _ = v27
                    case "foo bar":
                        v5 = (v5 * (^(v4)))
                    case "":
                        v28 := GenS4{F0: (v6 << ((0) & 63))}
                        _ = v28
                    }
                } else {
                    fmt.Println(v6)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                fmt.Println(v3, v4)
                _ = fi26
            }
            _ = fi25
        }
    } else {
        for fi29 := 0; fi29 < 9; fi29++ {
            var v30 uint8 = 209
            if (true || (!false)) {
                if ("hello" == (((v3 + (v3 + "\n\t")) + v3) + "foo bar")) {
                    if len(v1) > 0 {
                        v1[0] = 45
                    }
                    var v31 uint = 1
                    var v32 bool = ((v6 / (((v6 + (v6 % ((42) | 1)))) | 1)) <= (v6 | v6))
                    _ = v31
                    _ = v32
                }
                var v33 uint32 = uint32(v30)
                switch "foo bar" {
                case "foo bar":
                    v34 := GenS5{F0: v30}
                    _ = v34
                case "a":
                    for fi35 := 0; fi35 < 18; fi35++ {
                        v3 = v3
                        _ = fi35
                    }
                case "hello":
                    switch "hello" {
                    case "\n\t":
                        var v36 int64 = (v5 - v4)
                        _ = v36
                    case "foo bar":
                        var v37 int32 = 42
                        _ = v37
                    default:
                        switch "" {
                        case "a":
                            if false {
                                if false {
                                    var v38 bool = true
                                    _ = v38
                                }
                                fmt.Println(v33)
                                v39 := GenS6{F0: -9223372036854775808}
                                _ = v39
                            }
                        case "hello":
                            if (v6 <= -92) {
                                v5 = v5
                            } else {
                                v4 += 42
                                if ((-(v6)) > (v6 * v6)) {
                                    var v40 string = v3
                                    v41 := GenS2{F0: fi29, F1: 42, F2: v30}
                                    _ = v40
                                    _ = v41
                                } else {
                                    fmt.Println(v6, v33, v4)
                                }
                                if true {
                                    switch (fi29 ^ 128) {
                                    case 9223372036854775807:
                                        var v42 bool = true
                                        _ = v42
                                    default:
                                        if true {
                                            if false {
                                                var v43 int16 = -1
                                                _ = v43
                                            } else {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                for fi44 := 0; fi44 < 4; fi44++ {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    fmt.Println(fi29, v30)
                                                    switch fi29 {
                                                    case 2147483647:
                                                        v45 := GenS7{F0: true, F1: (-(fi44)), F2: v6}
                                                        _ = v45
                                                    case 42:
                                                        if (false && (46 <= (v30 << ((220) & 63)))) {
                                                            if true {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                if (!(v2 < v2)) {
                                                                    v46 := GenS8{F0: (v3 + (v3 + "foo bar")), F1: (v5 & v5), F2: v6}
                                                                    v47 := GenS4{F0: 0}
                                                                    v6 += 43
                                                                    _ = v46
                                                                    _ = v47
                                                                } else {
                                                                    var v48 int32 = 2147483647
                                                                    _ = v48
                                                                }
                                                                switch ((v3 + ("hello" + v3)) + v3) {
                                                                case "a":
                                                                    switch (v3 + "hello") {
                                                                    case "hello":
                                                                        var v49 int8 = (v6 << ((18446744073709551615) & 63))
                                                                        _ = v49
                                                                    case "\n\t":
                                                                        var v50 int8 = (v6 << ((255) & 63))
                                                                        v51 := GenS8{F0: ((v3 + "hello") + ((((("foo bar" + "foo bar") + "") + v3) + "hello") + "a")), F1: v5, F2: -128}
                                                                        if (true && ((3.14159 * float64(3.14159)) >= 1e10)) {
                                                                            var v52 uint64 = 4294967295
                                                                            var v53 int64 = v5
                                                                            if true {
                                                                                fmt.Println(v50, v33, v50)
                                                                                v54 := GenS6{F0: (-(v5))}
                                                                                _ = v54
                                                                            }
                                                                            _ = v52
                                                                            _ = v53
                                                                        }
                                                                        _ = v50
                                                                        _ = v51
                                                                    default:
                                                                        switch (-(fi44)) {
                                                                        case -414:
                                                                            var v55 bool = (!(!false))
                                                                            _ = v55
                                                                        case -940:
                                                                            var v56 int32 = 316
                                                                            if (1 != -1) {
                                                                                if (false || (v3 != "\n\t")) {
                                                                                    if true {
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = uint32(v30)
                                                                                        }
                                                                                        v56 = (v56 + v56)
                                                                                        var v57 int16 = 0
                                                                                        _ = v57
                                                                                    } else {
                                                                                        var v58 string = v3
                                                                                        _ = v58
                                                                                    }
                                                                                    v30 -= v30
                                                                                }
                                                                                var v59 uint = 4294967295
                                                                                _ = v59
                                                                            } else {
                                                                                var v60 bool = (("hello" + "foo bar") != (v3 + (v3 + "a")))
                                                                                v33 = (v33 & v33)
                                                                                v3 = v3
                                                                                v5 = v5
                                                                                _ = v60
                                                                            }
                                                                            _ = v56
                                                                        default:
                                                                            var v61 float32 = ((float32(-2.25) * (float32(0.0) - v2)) - float32(3.14159))
                                                                            _ = v61
                                                                        }
                                                                    }
                                                                case "foo bar":
                                                                    var v62 uint = 255
                                                                    _ = v62
                                                                case "hello":
                                                                    v6 = v6
                                                                default:
                                                                    var v63 float64 = 1.5
                                                                    _ = v63
                                                                }
                                                            }
                                                        } else {
                                                            _ = len(v1)
                                                            v5 -= v5
                                                            v1 = append(v1, v33)
                                                        }
                                                        _ = len(v1)
                                                        v4 -= 9223372036854775807
                                                    case 128:
                                                        switch v3 {
                                                        case "\n\t":
                                                            var v64 int64 = 0
                                                            var v65 int8 = 1
                                                            _ = v64
                                                            _ = v65
                                                        case "hello":
                                                            if (((v2 - v2) * v2) < (float32(-2.25) + v2)) {
                                                                _ = len(v1)
                                                                v6 = v6
                                                            }
                                                            var v66 float32 = (((v2 * (v2 / v2)) + v2) / v2)
                                                            _ = v66
                                                        case "a":
                                                            v3 = v3
                                                        default:
                                                            v7 = v7
                                                        }
                                                    }
                                                    _ = fi44
                                                }
                                                fmt.Println(v5, v6)
                                            }
                                            v30 = v30
                                        }
                                    }
                                    switch "hello" {
                                    case "a":
                                        v2 = (float32((0.0 + (1e10 - -2.25))) * (0.0 / v2))
                                    case "hello":
                                        var v67 uint64 = 1
                                        _ = v67
                                    case "":
                                        v68 := GenS1{F0: "hello", F1: v3, F2: -1}
                                        _ = v68
                                    default:
                                        v33 = (^(v33))
                                    }
                                } else {
                                    var v69 int64 = (v5 & int64(v33))
                                    fmt.Println(v30, v4)
                                    if len(v1) > 0 {
                                        v1[0] = 0
                                    }
                                    _ = v69
                                }
                            }
                        }
                    }
                }
                _ = v33
            }
            v70 := GenS8{F0: "hello", F1: (v4 / (((v5 / (((v4 % (((v5 & (v4 ^ v4))) | 1))) | 1))) | 1)), F2: (v6 << ((65535) & 63))}
            _ = fi29
            _ = v30
            _ = v70
        }
        v3 = v3
    }
    for fi71 := 0; fi71 < 8; fi71++ {
        v72 := GenS9{F0: 2147483647}
        v73 := GenS10{F0: true, F1: 42}
        v3 = "hello"
        _ = fi71
        _ = v72
        _ = v73
    }
    _ = len(v1)
    for fi74 := 0; fi74 < 20; fi74++ {
        if (true && (!false)) {
            v7 -= (v7 << ((2) & 63))
            var v75 string = (v3 + v3)
            var v76 int32 = 1
            _ = v75
            _ = v76
        }
        v2 -= 0.1
        _ = fi74
    }
    v5 = (v4 / ((int64(v7)) | 1))
    if false {
        v3 = (v3 + (v3 + "\n\t"))
        var v77 float32 = (v2 - -2.25)
        _ = v77
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
