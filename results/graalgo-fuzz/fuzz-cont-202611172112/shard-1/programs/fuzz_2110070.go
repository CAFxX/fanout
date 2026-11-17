package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS3 struct { F0 int8; F1 uint16; F2 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 int8; F1 uint; F2 uint16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint8; F1 uint32 }
type GenS9 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS10 struct { F0 uint64; F1 uint32 }
type GenS11 struct { F0 uint32; F1 uint16; F2 string }
type GenS12 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = 1
    var v3 int32 = -2147483648
    switch "foo bar" {
    case "":
        if true {
            v2 = v2
            var v4 int32 = 0
            _ = v4
        } else {
            if false {
                _ = len(v1)
                v5 := GenS1{F0: 42, F1: true}
                var v6 int32 = v3
                _ = v5
                _ = v6
            } else {
                for fi7 := 0; fi7 < 20; fi7++ {
                    v8 := []int{(fi7 - 32767), fi7, 553}
                    if false {
                        switch (fi7 / ((922) | 1)) {
                        case 127:
                            var v9 uint32 = 65535
                            _ = v9
                        default:
                            v3 += v3
                        }
                        v10 := GenS2{F0: 2, F1: -32768, F2: 65535}
                        switch "" {
                        case "":
                            var v11 int8 = 127
                            _ = v11
                        case "a":
                            switch "hello" {
                            case "":
                                fmt.Println(v2, v3, v3)
                            case "hello":
                                switch "a" {
                                case "":
                                    for fi12 := 0; fi12 < 20; fi12++ {
                                        switch "hello" {
                                        case "a":
                                            var v13 uint32 = 4294967295
                                            _ = v13
                                        case "":
                                            var v14 int32 = v3
                                            _ = v14
                                        default:
                                            _ = len(v8)
                                        }
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                        var v15 int32 = (v3 & v3)
                                        _ = fi12
                                        _ = v15
                                    }
                                default:
                                    if (0 <= 4294967295) {
                                        v16 := GenS3{F0: 0, F1: 2, F2: (v2 << ((199) & 63))}
                                        var v17 uint32 = 0
                                        _ = v16
                                        _ = v17
                                    } else {
                                        var v18 uint16 = 913
                                        var v19 uint16 = (v18 & v18)
                                        _ = v18
                                        _ = v19
                                    }
                                }
                            default:
                                for fi20 := 0; fi20 < 6; fi20++ {
                                    fmt.Println(v2, v2)
                                    if false {
                                        var v21 uint32 = 257
                                        _ = v21
                                    }
                                    _ = fi20
                                }
                            }
                        case "\n\t":
                            v3 = 1
                        default:
                            if true {
                                var v22 uint8 = 1
                                if true {
                                    _ = len(v8)
                                    v3 += (v3 << ((65535) & 63))
                                } else {
                                    switch 46 {
                                    case 859:
                                        if (!false) {
                                            if (65535 <= 2) {
                                                var v23 uint = 2
                                                _ = v23
                                            } else {
                                                for fi24 := 0; fi24 < 12; fi24++ {
                                                    v25 := GenS2{F0: v22, F1: -1, F2: 490}
                                                    v3 -= v3
                                                    var v26 uint8 = v22
                                                    _ = fi24
                                                    _ = v25
                                                    _ = v26
                                                }
                                                if true {
                                                    v2 = -9223372036854775808
                                                    switch "foo bar" {
                                                    case "hello":
                                                        v27 := GenS3{F0: 21, F1: 0, F2: v2}
                                                        _ = v27
                                                    case "\n\t":
                                                        if (1e10 < ((1.5 - -2.25) * 3.14159)) {
                                                            for fi28 := 0; fi28 < 15; fi28++ {
                                                                if true {
                                                                    if len(v8) > 0 {
                                                                        v8[0] = fi28
                                                                    }
                                                                    if len(v8) > 0 {
                                                                        v8[0] = fi7
                                                                    }
                                                                    v8 = append(v8, (fi28 % (((fi7 + (fi28 / ((-59) | 1)))) | 1)))
                                                                } else {
                                                                    if len(v8) >= 2 {
                                                                        v8 = v8[1:len(v8)]
                                                                    }
                                                                }
                                                                _ = fi28
                                                            }
                                                            var v29 string = "hello"
                                                            _ = v29
                                                        } else {
                                                            _ = len(v1)
                                                            if false {
                                                                v3 = 2147483647
                                                                var v30 int8 = 100
                                                                var v31 uint64 = 590
                                                                _ = v30
                                                                _ = v31
                                                            }
                                                            fmt.Println(v2, v2)
                                                        }
                                                    case "":
                                                        if len(v8) >= 2 {
                                                            v8 = v8[1:len(v8)]
                                                        }
                                                    default:
                                                        switch fi7 {
                                                        case -9223372036854775808:
                                                            v22 += 128
                                                        }
                                                    }
                                                } else {
                                                    if false {
                                                        if len(v8) > 0 {
                                                            v8[0] = 2147483647
                                                        }
                                                    }
                                                    v32 := GenS4{F0: (-(v3))}
                                                    v3 = 42
                                                    _ = v32
                                                }
                                                var v33 int = fi7
                                                _ = v33
                                            }
                                            v34 := GenS5{F0: "", F1: 594}
                                            _ = v34
                                        } else {
                                            fmt.Println(v2, v2, v2)
                                        }
                                    case 32767:
                                        v22 -= (v22 + 0)
                                    }
                                    v35 := GenS1{F0: 2147483648, F1: false}
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                    _ = v35
                                }
                                _ = v22
                            }
                        }
                        _ = v10
                    } else {
                        v3 += v3
                        switch (-(fi7)) {
                        case -421:
                            var v36 bool = true
                            _ = v36
                        case 4294967295:
                            if (0.0 != -2.25) {
                                var v37 int64 = v2
                                switch "hello" {
                                case "\n\t":
                                    var v38 float32 = ((1.5 - (0.0 + (1e10 - 0.0))) / (0.0 + 3.14159))
                                    _ = v38
                                case "":
                                    if true {
                                        var v39 int32 = (^(v3))
                                        var v40 int8 = 110
                                        _ = v39
                                        _ = v40
                                    } else {
                                        v41 := GenS3{F0: 127, F1: 32767, F2: -468}
                                        for fi42 := 0; fi42 < 15; fi42++ {
                                            v43 := GenS4{F0: v3}
                                            _ = fi42
                                            _ = v43
                                        }
                                        _ = v41
                                    }
                                case "hello":
                                    for fi44 := 0; fi44 < 13; fi44++ {
                                        fmt.Println(v3, fi7)
                                        _ = fi44
                                    }
                                }
                                _ = v37
                            }
                        }
                        switch 1 {
                        case -71:
                            v2 -= (v2 - v2)
                        case 9223372036854775807:
                            if false {
                                v45 := GenS6{F0: -128, F1: 334, F2: 42}
                                for fi46 := 0; fi46 < 19; fi46++ {
                                    if (true || false) {
                                        if ((3.14159 < ((0.0 + 3.14159) - (3.14159 * (-2.25 * 1e10)))) && true) {
                                            if true {
                                                v2 -= v2
                                                v2 = v2
                                            } else {
                                                v1[2] = (167 != 127)
                                                if true {
                                                    var v47 int16 = -883
                                                    fmt.Println(v47, fi7, fi46)
                                                    v3 = v3
                                                    _ = v47
                                                } else {
                                                    if len(v8) > 0 {
                                                        v8[0] = (fi46 + (fi46 - (fi7 | 4294967295)))
                                                    }
                                                    var v48 bool = true
                                                    v49 := GenS5{F0: "\n\t", F1: 447}
                                                    _ = v48
                                                    _ = v49
                                                }
                                            }
                                            v3 -= v3
                                        }
                                    }
                                    v2 += v2
                                    var v50 float32 = ((0.0 * (1.5 - 1.5)) * 0.0)
                                    _ = fi46
                                    _ = v50
                                }
                                v51 := GenS4{F0: v3}
                                _ = v45
                                _ = v51
                            } else {
                                switch (((((((("foo bar" + ("hello" + ((((("\n\t" + "a") + "") + "hello") + ("\n\t" + "a")) + "foo bar"))) + "foo bar") + ("hello" + "hello")) + ("hello" + "foo bar")) + "a") + ("hello" + "foo bar")) + ("foo bar" + (("hello" + (("hello" + "\n\t") + "\n\t")) + ("" + "")))) + (("" + "hello") + "foo bar")) {
                                case "foo bar":
                                    var v52 uint8 = 42
                                    _ = v52
                                case "":
                                    if (111 < -100) {
                                        v3 = (v3 + (v3 / (((v3 | (v3 + (v3 | (v3 << ((255) & 63)))))) | 1)))
                                        var v53 uint16 = 65535
                                        if false {
                                            v3 += (v3 ^ -1)
                                            v3 -= 42
                                            v53 -= 732
                                        }
                                        _ = v53
                                    }
                                case "a":
                                    _ = len(v8)
                                }
                                switch ("" + "foo bar") {
                                case "\n\t":
                                    var v54 int32 = (^(v3))
                                    _ = v54
                                case "foo bar":
                                    v2 += (-(v2))
                                }
                            }
                        case 1:
                            v2 += 0
                        default:
                            _ = len(v1)
                        }
                    }
                    var v55 int8 = 41
                    _ = fi7
                    _ = v55
                    _ = v8
                }
            }
            for fi56 := 0; fi56 < 15; fi56++ {
                v57 := GenS5{F0: (((("\n\t" + "a") + ("a" + "a")) + "") + "a"), F1: 4294967295}
                _ = fi56
                _ = v57
            }
            v3 = v3
        }
    case "a":
        var v58 string = "hello"
        _ = v58
    default:
        v59 := GenS7{F0: ""}
        _ = v59
    }
    var v60 uint16 = 65535
    v61 := []int{-2147483648}
    if (true || false) {
        _ = len(v61)
        if true {
            var v62 int = 32768
            _ = v62
        }
    } else {
        var v63 uint64 = 42
        switch -233 {
        case -879:
            v2 = (v2 ^ v2)
        case -1:
            var v64 uint8 = 127
            _ = v64
        case -9223372036854775808:
            var v65 uint16 = 238
            _ = v65
        default:
            v66 := GenS6{F0: -128, F1: 0, F2: 255}
            _ = v66
        }
        for fi67 := 0; fi67 < 8; fi67++ {
            for fi68 := 0; fi68 < 10; fi68++ {
                v2 -= (v2 + -1)
                switch ((("" + "a") + "hello") + ("foo bar" + "hello")) {
                case "a":
                    v63 -= (^(v63))
                case "hello":
                    v69 := v1[2147483648]
                    _ = v69
                case "foo bar":
                    v2 = v2
                }
                v70 := GenS8{F0: 205, F1: 4294967295}
                _ = fi68
                _ = v70
            }
            if (!(true && true)) {
                switch fi67 {
                case 128:
                    v71 := GenS9{F0: uint64(v60), F1: (v63 / (((v63 - v63)) | 1)), F2: 32767}
                    _ = v71
                case 2147483647:
                    v72 := GenS10{F0: 167, F1: 556}
                    _ = v72
                case 370:
                    for fi73 := 0; fi73 < 18; fi73++ {
                        v2 += 482
                        switch "foo bar" {
                        case "":
                            if true {
                                fmt.Println(v63)
                            }
                        case "hello":
                            v74 := GenS11{F0: 2147483647, F1: v60, F2: ("hello" + "\n\t")}
                            _ = v74
                        case "\n\t":
                            v75 := v1[uint32(v63)]
                            _ = v75
                        default:
                            var v76 uint64 = 18446744073709551615
                            _ = v76
                        }
                        _ = fi73
                    }
                default:
                    v3 = (v3 % ((-1) | 1))
                }
                v63 -= 4294967295
            }
            _ = fi67
        }
        _ = v63
    }
    for fi77 := 0; fi77 < 6; fi77++ {
        var v78 float64 = (0.0 * 0.1)
        v79 := GenS8{F0: 128, F1: 42}
        var v80 uint16 = (v60 & v60)
        _ = fi77
        _ = v78
        _ = v79
        _ = v80
    }
    v81 := GenS8{F0: 98, F1: 123}
    if ((-830 < 32767) || (!(true || true))) {
        if true {
            var v82 int8 = 1
            var v83 int8 = v82
            if true {
                if len(v61) >= 2 {
                    v61 = v61[1:len(v61)]
                }
                var v84 int8 = v83
                _ = v84
            } else {
                var v85 uint64 = 101
                _ = v85
            }
            _ = v82
            _ = v83
        } else {
            v3 = v3
            v60 += (v60 & (^(v60)))
        }
        v86 := GenS12{F0: (v60 / (((^(v60))) | 1))}
        _ = v86
    }
    switch "" {
    case "hello":
        fmt.Println(v2)
    case "a":
        v87 := GenS8{F0: 134, F1: uint32(v60)}
        _ = v87
    case "\n\t":
        var v88 uint64 = 909
        _ = v88
    default:
        for fi89 := 0; fi89 < 2; fi89++ {
            var v90 float32 = (1.5 - 0.1)
            _ = fi89
            _ = v90
        }
    }
    delete(v1, 193)
    v2 = (v2 + (v2 + v2))
    for fi91 := 0; fi91 < 5; fi91++ {
        v2 = v2
        var v92 float64 = (((3.14159 * float64(1e10)) - (((1e10 * 1.5) / 3.14159) * ((1.5 - 3.14159) * 1e10))) * (0.1 / (0.1 - float64(float32(0.0)))))
        v93 := GenS8{F0: 255, F1: 2}
        _ = fi91
        _ = v92
        _ = v93
    }
    for fi94 := 0; fi94 < 5; fi94++ {
        var v95 float64 = 1e10
        var v96 int64 = v2
        var v97 int = fi94
        _ = fi94
        _ = v95
        _ = v96
        _ = v97
    }
    v3 = 42
    var v98 int64 = v2
    for fi99 := 0; fi99 < 2; fi99++ {
        var v100 int8 = 1
        v98 -= (v98 + (v2 - (-(v98))))
        _ = fi99
        _ = v100
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v61:", len(v61), v61)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v60
    _ = v61
    _ = v81
    _ = v98
}
