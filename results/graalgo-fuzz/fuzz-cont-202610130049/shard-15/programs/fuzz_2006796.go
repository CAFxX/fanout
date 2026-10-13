package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 int64; F2 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 int64; F2 int64 }
type GenS5 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int8; F1 string; F2 int8 }
type GenS8 struct { F0 string; F1 uint8 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int16; F1 int16; F2 string }
type GenS11 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS12 struct { F0 uint64; F1 int64 }
type GenS13 struct { F0 int; F1 uint32 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 int16; F1 int }
func genFn0(p0 int, p1 uint8, p2 string) uint {
    v6 := GenS1{F0: p0}
    if true {
        if (42 > 4294967295) {
            var v7 uint32 = 2
            v8 := GenS1{F0: (p0 / ((984) | 1))}
            _ = v7
            _ = v8
        } else {
            var v9 uint32 = 0
            v10 := GenS2{F0: 2, F1: 0, F2: "a"}
            _ = v10
            _ = v9
        }
        v11 := GenS3{F0: 100}
        _ = v11
    }
    var v12 int32 = 42
    v13 := GenS3{F0: 22}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v6
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 65535
    v3 := make(map[uint32]string)
    var v4 uint8 = 199
    v5 := []int{}
    switch "" {
    case "":
        v1["foo bar"] = 453
    case "hello":
        for fi14 := 0; fi14 < 12; fi14++ {
            for fi15 := 0; fi15 < 18; fi15++ {
                if false {
                    v16 := GenS1{F0: (fi14 + (fi14 / ((32768) | 1)))}
                    fmt.Println(fi14, fi14)
                    var v17 uint32 = 791
                    _ = v16
                    _ = v17
                }
                if ("hello" <= "\n\t") {
                    v4 += v4
                    delete(v3, 2)
                    var v18 uint32 = uint32(v4)
                    _ = v18
                } else {
                    v4 = 113
                    v2 = 4294967295
                    if true {
                        v19 := GenS4{F0: (65535 <= 542), F1: 9223372036854775807, F2: 0}
                        v4 = (v4 ^ (v4 << (((v2 - v2)) & 63)))
                        _ = v19
                    } else {
                        v20 := GenS3{F0: 42}
                        switch "a" {
                        case "foo bar":
                            if true {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            switch (fi14 << ((genFn0((fi14 ^ fi14), (v4 + v4), ((("a" + "foo bar") + ("a" + "")) + "foo bar"))) & 63)) {
                            case -288:
                                if (!true) {
                                    if true {
                                        v21 := GenS5{F0: 1, F1: 2147483648, F2: 100}
                                        v4 = v4
                                        if false {
                                            v2 += 0
                                            v22 := GenS6{F0: 633}
                                            if false {
                                                v23 := GenS7{F0: 42, F1: "foo bar", F2: 42}
                                                v24 := GenS3{F0: -115}
                                                fmt.Println(fi15, fi15)
                                                _ = v23
                                                _ = v24
                                            } else {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v25 := GenS8{F0: "foo bar", F1: uint8(v2)}
                                                    _ = v25
                                                case "\n\t":
                                                    var v26 uint64 = 161
                                                    _ = v26
                                                case "hello":
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                }
                                                if len(v5) > 0 {
                                                    v5[0] = (fi14 * (^(fi15)))
                                                }
                                                var v27 int = 1
                                                _ = v27
                                            }
                                            _ = v22
                                        } else {
                                            var v28 int16 = -904
                                            v29 := v3[42]
                                            _ = v28
                                            _ = v29
                                        }
                                        _ = v21
                                    } else {
                                        v2 = (v2 / ((genFn0(fi14, v4, "")) | 1))
                                        var v30 int64 = 42
                                        v4 = v4
                                        _ = v30
                                    }
                                    switch 128 {
                                    case 869:
                                        v2 -= (^(v2))
                                    default:
                                        v5 = append(v5, (fi14 >> ((v2) & 63)))
                                    }
                                }
                            default:
                                if (612 == -32768) {
                                    var v31 uint = (v2 / (((^(v2))) | 1))
                                    _ = v31
                                } else {
                                    var v32 int32 = 42
                                    _ = v32
                                }
                            }
                            v33 := GenS9{F0: 0}
                            _ = v33
                        case "\n\t":
                            var v34 uint16 = 919
                            _ = v34
                        case "":
                            var v35 uint32 = 2
                            _ = v35
                        default:
                            v4 += uint8(v2)
                        }
                        var v36 float32 = 0.0
                        _ = v20
                        _ = v36
                    }
                    v4 += v4
                    v4 = v4
                    var v37 float32 = (0.1 / float32(1e10))
                    _ = v37
                }
                var v38 int = fi15
                _ = fi15
                _ = v38
            }
            switch (fi14 / ((32768) | 1)) {
            case 0:
                v3[435] = "foo bar"
            default:
                _ = len(v5)
            }
            _ = fi14
        }
    }
    switch ((("" + ("a" + "hello")) + ("\n\t" + "\n\t")) + (((("a" + "hello") + "") + "\n\t") + "hello")) {
    case "foo bar":
        if (v4 > 255) {
            v4 = 2
        } else {
            delete(v3, 0)
        }
    case "hello":
        for fi39 := 0; fi39 < 9; fi39++ {
            fmt.Println(fi39)
            if true {
                v1["foo bar"] = 341
            }
            _ = fi39
        }
    default:
        for fi40 := 0; fi40 < 18; fi40++ {
            var v41 float32 = ((float32((3.14159 * -2.25)) - (float32(float64(float32(1.5))) / 1.0)) / ((((((-2.25 - float32(0.1)) * 1.5) + 3.14159) * -2.25) - float32(-2.25)) / 1.0))
            _ = fi40
            _ = v41
        }
    }
    switch -1 {
    case -1:
        if ((((3.14159 - 0.0) * float64(0.0)) > 1.5) && true) {
            v42 := GenS2{F0: 65535, F1: -121, F2: ("a" + "hello")}
            switch -789 {
            case 2147483647:
                switch (("foo bar" + ("hello" + ("\n\t" + ("foo bar" + ("foo bar" + ""))))) + "a") {
                case "a":
                    var v43 uint16 = 527
                    _ = v43
                case "\n\t":
                    switch -372 {
                    case -736:
                        for fi44 := 0; fi44 < 19; fi44++ {
                            v4 -= uint8(v2)
                            switch ("" + ("a" + "\n\t")) {
                            case "\n\t":
                                _ = len(v5)
                            }
                            if false {
                                switch ("" + "hello") {
                                case "":
                                    if (65535 == 2147483648) {
                                        var v45 uint64 = 0
                                        var v46 int32 = 42
                                        fmt.Println(v45, fi44, v46)
                                        _ = v45
                                        _ = v46
                                    } else {
                                        fmt.Println(fi44, fi44, v2)
                                        switch "" {
                                        case "foo bar":
                                            switch fi44 {
                                            case 4294967295:
                                                v47 := GenS10{F0: -551, F1: -689, F2: "\n\t"}
                                                _ = v47
                                            case 496:
                                                v48 := GenS11{F0: (478 >= 2147483648), F1: v4, F2: uint32(v4)}
                                                _ = v48
                                            case 207:
                                                if (!true) {
                                                    var v49 uint32 = 2
                                                    v5 = append(v5, fi44)
                                                    _ = v49
                                                }
                                            }
                                        case "hello":
                                            var v50 float64 = 0.0
                                            _ = v50
                                        case "":
                                            for fi51 := 0; fi51 < 10; fi51++ {
                                                if false {
                                                    var v52 float64 = -2.25
                                                    var v53 int = (^(fi44))
                                                    _ = v52
                                                    _ = v53
                                                }
                                                _ = fi51
                                            }
                                        }
                                    }
                                }
                                var v54 int32 = 0
                                _ = v54
                            }
                            _ = fi44
                        }
                    case -84:
                        var v55 uint64 = 349
                        _ = v55
                    case -68:
                        var v56 uint64 = 901
                        _ = v56
                    default:
                        if true {
                            v1["hello"] = 161
                        } else {
                            switch 2147483647 {
                            case 32767:
                                var v57 int = 678
                                _ = v57
                            default:
                                _ = len(v5)
                            }
                            v4 -= v4
                            v2 += 1
                        }
                    }
                case "hello":
                    var v58 int64 = 459
                    _ = v58
                default:
                    var v59 uint64 = 4294967295
                    _ = v59
                }
            case 127:
                var v60 bool = true
                _ = v60
            case -440:
                v4 = (v4 - v4)
            default:
                var v61 string = "foo bar"
                _ = v61
            }
            _ = v42
        } else {
            var v62 int32 = 2147483647
            _ = v62
        }
    case 9223372036854775807:
        switch "foo bar" {
        case "foo bar":
            v63 := GenS1{F0: 32768}
            _ = v63
        case "\n\t":
            for fi64 := 0; fi64 < 7; fi64++ {
                for fi65 := 0; fi65 < 8; fi65++ {
                    v66 := GenS12{F0: 0, F1: 9223372036854775807}
                    v2 = (v2 + 965)
                    v67 := GenS13{F0: 42, F1: 1}
                    var v68 float32 = float32(0.1)
                    _ = fi65
                    _ = v66
                    _ = v67
                    _ = v68
                }
                _ = fi64
            }
        }
    default:
        v1["foo bar"] = 65535
    }
    if (!(false && true)) {
        v4 = (v4 & v4)
    }
    if ("" <= ("\n\t" + ("foo bar" + "hello"))) {
        v2 = (v2 | 4294967295)
    } else {
        fmt.Println(v4, v2)
        var v69 int64 = 308
        _ = v69
    }
    var v70 int32 = 978
    _ = len(v3)
    if (!false) {
        var v71 bool = false
        _ = v71
    }
    v70 -= (v70 % ((-1) | 1))
    if len(v5) > 0 {
        v5[0] = 410
    }
    v2 += genFn0(0, v4, ("hello" + "a"))
    var v72 uint32 = 2147483647
    if ("a" > "a") {
        for fi73 := 0; fi73 < 18; fi73++ {
            for fi74 := 0; fi74 < 17; fi74++ {
                var v75 int8 = -128
                v76 := GenS7{F0: (v75 * 0), F1: "a", F2: (v75 - v75)}
                switch "\n\t" {
                case "a":
                    if (!false) {
                        v77 := GenS11{F0: false, F1: v4, F2: 65535}
                        _ = v77
                    } else {
                        v75 -= (v75 ^ (v75 * v75))
                        fmt.Println(v4, v72, fi74)
                        if true {
                            if true {
                                v78 := GenS14{F0: v2}
                                _ = v78
                            }
                            var v79 uint8 = (^(v4))
                            fmt.Println(fi74)
                            _ = v79
                        } else {
                            var v80 uint32 = v72
                            v81 := GenS5{F0: 65535, F1: uint32(v4), F2: v75}
                            _ = v80
                            _ = v81
                        }
                    }
                case "foo bar":
                    switch (fi73 << ((65535) & 63)) {
                    case -9223372036854775808:
                        var v82 bool = (1 == 317)
                        _ = v82
                    default:
                        switch 127 {
                        case -218:
                            v75 += v75
                        }
                    }
                default:
                    if true {
                        v83 := GenS3{F0: int8(v72)}
                        var v84 float32 = 1e10
                        var v85 uint = (^(v2))
                        _ = v83
                        _ = v84
                        _ = v85
                    } else {
                        switch (("\n\t" + ("foo bar" + ("" + "hello"))) + ("hello" + ("foo bar" + ("foo bar" + ((("foo bar" + "hello") + "foo bar") + ("" + "foo bar")))))) {
                        case "foo bar":
                            v75 -= v75
                        case "\n\t":
                            v70 += v70
                        case "":
                            var v86 int = (fi73 | int(v75))
                            _ = v86
                        }
                        var v87 int64 = int64(v75)
                        if len(v5) > 0 {
                            v5[0] = (fi74 & 416)
                        }
                        _ = v87
                    }
                }
                var v88 uint16 = 65535
                _ = fi74
                _ = v75
                _ = v76
                _ = v88
            }
            _ = fi73
        }
        v4 = (v4 * (^(v4)))
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    var v89 int32 = (v70 / (((v70 % ((int32(v4)) | 1))) | 1))
    fmt.Println(v2)
    fmt.Println(v72, v72, v70)
    v90 := GenS15{F0: -1, F1: 127}
    if len(v5) > 0 {
        v5[0] = int(v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v70:", v70)
    fmt.Println("v72:", v72)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v70
    _ = v72
    _ = v89
    _ = v90
}
