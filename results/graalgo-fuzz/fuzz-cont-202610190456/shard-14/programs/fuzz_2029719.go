package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint32; F1 uint64; F2 int16 }

func main() {
    var v1 int8 = 42
    v2 := []uint32{42, 4294967295, 950}
    var v3 float32 = 0.1
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := GenS1{F0: (fi4 << ((4294967295) & 63)), F1: 2}
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 10; fi6++ {
        if (!false) {
            for fi7 := 0; fi7 < 17; fi7++ {
                switch fi7 {
                case 9223372036854775807:
                    v1 = (v1 ^ v1)
                case 2147483647:
                    v3 = (-2.25 + v3)
                default:
                    switch "foo bar" {
                    case "hello":
                        if true {
                            v3 = 0.0
                            v1 = 0
                            v1 -= 122
                        }
                    }
                }
                if true {
                    v1 = (^(v1))
                    v3 = v3
                    v1 = v1
                } else {
                    v8 := GenS2{F0: -260, F1: 42}
                    switch "\n\t" {
                    case "":
                        switch "" {
                        case "":
                            v9 := GenS1{F0: (fi6 + fi7), F1: 42}
                            _ = v9
                        case "hello":
                            v2 = append(v2, 42)
                        case "foo bar":
                            v1 -= 127
                        default:
                            var v10 bool = (388 >= 2)
                            _ = v10
                        }
                    case "\n\t":
                        v2 = append(v2, 0)
                        if len(v2) > 0 {
                            v2[0] = 0
                        }
                    case "a":
                        if ("a" == ((("a" + "a") + "a") + "\n\t")) {
                            v2 = append(v2, 42)
                        }
                    }
                    v3 += v3
                    v2 = append(v2, 65535)
                    _ = v8
                }
                var v11 float32 = ((((1e10 + 3.14159) / (float32(float64(float32((float64(float32(3.14159)) + (0.1 + (1e10 - 0.1)))))) - v3)) * 1e10) - v3)
                _ = fi7
                _ = v11
            }
            var v12 uint8 = 170
            _ = v12
        } else {
            v13 := GenS3{F0: 9223372036854775807}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v13
        }
        _ = fi6
    }
    if true {
        var v14 int64 = 775
        _ = v14
    } else {
        if (float32((0.1 - 3.14159)) >= v3) {
            if (!(32768 == 1)) {
                var v15 string = "a"
                var v16 int16 = -1
                for fi17 := 0; fi17 < 1; fi17++ {
                    switch fi17 {
                    case 127:
                        switch (fi17 % ((fi17) | 1)) {
                        case -769:
                            if true {
                                v18 := GenS4{F0: 18446744073709551615}
                                var v19 float64 = float64(0.1)
                                _ = v18
                                _ = v19
                            }
                        }
                    case -938:
                        for fi20 := 0; fi20 < 5; fi20++ {
                            switch fi20 {
                            case -2147483648:
                                var v21 uint = 0
                                _ = v21
                            case 128:
                                if (2 <= 42) {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    var v22 int64 = -9223372036854775808
                                    v15 = "a"
                                    _ = v22
                                }
                            }
                            v3 += -2.25
                            if (false && true) {
                                v1 = v1
                                var v23 int = -9223372036854775808
                                _ = v23
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            _ = fi20
                        }
                    default:
                        var v24 uint32 = 502
                        _ = v24
                    }
                    var v25 string = v15
                    _ = fi17
                    _ = v25
                }
                _ = v15
                _ = v16
            } else {
                if (!false) {
                    for fi26 := 0; fi26 < 10; fi26++ {
                        var v27 string = ("" + "\n\t")
                        if true {
                            v28 := GenS5{F0: (fi26 & fi26)}
                            var v29 int8 = 0
                            _ = v28
                            _ = v29
                        }
                        if false {
                            v1 = 42
                        } else {
                            if ((-2.25 == 3.14159) || (2 > 1)) {
                                switch (fi26 >> ((255) & 63)) {
                                case -758:
                                    fmt.Println(v3)
                                case 32767:
                                    for fi30 := 0; fi30 < 9; fi30++ {
                                        var v31 int16 = 32767
                                        v1 = (v1 - (v1 << ((18446744073709551615) & 63)))
                                        _ = fi30
                                        _ = v31
                                    }
                                default:
                                    switch 226 {
                                    case -1:
                                        v27 = ""
                                    default:
                                        var v32 int8 = 100
                                        _ = v32
                                    }
                                }
                                if (!((((((((v3 + ((v3 - v3) - ((v3 + v3) / v3))) * v3) * v3) / ((((-2.25 + 1e10) / v3) - v3) + (v3 * (v3 - (0.1 * float32((1.5 - 0.1))))))) + v3) > (v3 / float32(-2.25))) || true) || true)) {
                                    var v33 float64 = 1e10
                                    if (fi26 == -9223372036854775808) {
                                        var v34 int = 127
                                        v1 = 42
                                        _ = v34
                                    } else {
                                        var v35 uint32 = 2147483648
                                        if len(v2) > 0 {
                                            v2[0] = (v35 / ((0) | 1))
                                        }
                                        _ = v35
                                    }
                                    _ = v33
                                } else {
                                    switch fi26 {
                                    case 131:
                                        for fi36 := 0; fi36 < 14; fi36++ {
                                            var v37 int32 = 1
                                            fmt.Println(v37, v37)
                                            v38 := GenS5{F0: 4294967295}
                                            _ = fi36
                                            _ = v37
                                            _ = v38
                                        }
                                    case -2147483648:
                                        if true {
                                            for fi39 := 0; fi39 < 4; fi39++ {
                                                if true {
                                                    v1 = 127
                                                } else {
                                                    var v40 int = 4294967295
                                                    if (!(fi39 > (fi39 / ((32768) | 1)))) {
                                                        var v41 string = v27
                                                        v27 = (v41 + ("\n\t" + v41))
                                                        _ = v41
                                                    }
                                                    var v42 float64 = (0.1 / 1.0)
                                                    _ = v40
                                                    _ = v42
                                                }
                                                _ = fi39
                                            }
                                        }
                                    default:
                                        _ = len(v2)
                                    }
                                    fmt.Println(v3)
                                }
                                v43 := GenS6{F0: 2147483648, F1: 0, F2: 42}
                                _ = v43
                            } else {
                                var v44 int = (fi26 + fi26)
                                v27 = v27
                                _ = v44
                            }
                            v3 -= 0.1
                        }
                        _ = fi26
                        _ = v27
                    }
                    v1 = (v1 - (v1 | v1))
                    v1 -= 42
                } else {
                    v3 = (float32((3.14159 - 0.1)) + 1e10)
                    fmt.Println(v3)
                }
            }
            for fi45 := 0; fi45 < 11; fi45++ {
                v1 -= (v1 % (((v1 >> ((18446744073709551615) & 63))) | 1))
                for fi46 := 0; fi46 < 6; fi46++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v47 int8 = (v1 * v1)
                    v47 -= (-(v1))
                    _ = fi46
                    _ = v47
                }
                _ = fi45
            }
            fmt.Println(v1)
        } else {
            v1 = (v1 * v1)
            if true {
                var v48 float64 = 0.0
                var v49 uint32 = 50
                if false {
                    switch 0 {
                    case 2147483647:
                        var v50 int32 = -838
                        _ = v50
                    case -9223372036854775808:
                        if ((true && true) || (1000 > -1)) {
                            switch ("foo bar" + ("foo bar" + "\n\t")) {
                            case "":
                                var v51 int = -182
                                _ = v51
                            case "\n\t":
                                v1 += 13
                            case "a":
                                fmt.Println(v49)
                            default:
                                fmt.Println(v3, v3, v1)
                            }
                            if (v1 == (v1 ^ v1)) {
                                if ((159 == 1) && true) {
                                    for fi52 := 0; fi52 < 17; fi52++ {
                                        var v53 bool = true
                                        fmt.Println(v53, v49, v53)
                                        var v54 int16 = 1000
                                        _ = fi52
                                        _ = v53
                                        _ = v54
                                    }
                                } else {
                                    v49 -= v49
                                    if (true && true) {
                                        var v55 int32 = -1
                                        _ = v55
                                    }
                                }
                                switch "a" {
                                case "hello":
                                    if false {
                                        fmt.Println(v1, v1)
                                        var v56 int16 = 1000
                                        v3 = float32(v48)
                                        _ = v56
                                    }
                                }
                            } else {
                                var v57 float32 = v3
                                v49 = (^(v49))
                                _ = v57
                            }
                            v1 += v1
                        }
                    case -608:
                        fmt.Println(v3, v3, v48)
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    fmt.Println(v49, v49)
                    for fi58 := 0; fi58 < 13; fi58++ {
                        v1 = (v1 >> ((1) & 63))
                        _ = fi58
                    }
                }
                _ = v48
                _ = v49
            } else {
                v59 := GenS4{F0: 452}
                _ = v59
            }
        }
    }
    v3 += 1.5
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch 554 {
    case 846:
        switch 707 {
        case 42:
            var v60 uint8 = 0
            _ = v60
        default:
            if len(v2) > 0 {
                v2[0] = 0
            }
        }
    case 32767:
        if true {
            var v61 int16 = -32768
            fmt.Println(v61)
            _ = v61
        } else {
            for fi62 := 0; fi62 < 3; fi62++ {
                v1 += (v1 & (^(v1)))
                _ = fi62
            }
            var v63 int = -2147483648
            _ = v63
        }
    case 42:
        v64 := GenS1{F0: 42, F1: 4294967295}
        _ = v64
    }
    if (-1 <= -1) {
        var v65 uint32 = 4294967295
        _ = v65
    }
    fmt.Println(v3)
    switch 151 {
    case 32768:
        var v66 uint16 = 32767
        _ = v66
    case -2147483648:
        var v67 int32 = 0
        _ = v67
    }
    if true {
        v68 := GenS5{F0: 128}
        _ = v68
    }
    var v69 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v69
}
