package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 int8; F1 uint }
func genFn0(p0 int32) uint8 {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 2; fi4++ {
        p0 -= (p0 ^ (p0 | p0))
        _ = fi4
    }
    _ = p0
    return 127
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 65535
    v3 := []uint32{0, 0, 4294967295}
    v2 += v2
    fmt.Println(v2)
    v2 = (v2 ^ (^(v2)))
    var v5 uint8 = 0
    v2 = (v2 % ((768) | 1))
    if len(v3) > 0 {
        v3[0] = 2
    }
    v6 := GenS1{F0: (v5 << ((1) & 63)), F1: (v2 - v2), F2: 255}
    if true {
        v5 = genFn0(-277)
        switch "a" {
        case "a":
            for fi7 := 0; fi7 < 19; fi7++ {
                switch (fi7 >> ((1) & 63)) {
                case 42:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case 128:
                    v8 := v1[42]
                    _ = v8
                case 32768:
                    v9 := GenS2{F0: genFn0(42), F1: 980, F2: 42}
                    _ = v9
                default:
                    _ = len(v3)
                }
                fmt.Println(fi7)
                _ = fi7
            }
        case "\n\t":
            for fi10 := 0; fi10 < 14; fi10++ {
                if false {
                    if false {
                        for fi11 := 0; fi11 < 3; fi11++ {
                            v12 := GenS2{F0: 0, F1: 2, F2: 2147483647}
                            v2 -= v2
                            _ = fi11
                            _ = v12
                        }
                        v13 := GenS2{F0: genFn0(1), F1: 0, F2: 360}
                        v1[2147483647] = 42
                        _ = v13
                    } else {
                        if (275 > 0) {
                            if (42 != -987) {
                                switch "hello" {
                                case "a":
                                    switch fi10 {
                                    case 32768:
                                        v3 = append(v3, 65535)
                                    }
                                case "":
                                    if false {
                                        switch -410 {
                                        case -2147483648:
                                            v5 += v5
                                        case 4294967295:
                                            switch "\n\t" {
                                            case "":
                                                if false {
                                                    fmt.Println(fi10, v5, v2)
                                                    v5 -= (v5 + uint8(v2))
                                                    v2 = (v2 << ((394) & 63))
                                                }
                                            case "hello":
                                                fmt.Println(v2)
                                            default:
                                                var v14 int8 = 1
                                                _ = v14
                                            }
                                        default:
                                            if true {
                                                for fi15 := 0; fi15 < 13; fi15++ {
                                                    if (true && (true || (false && (-9223372036854775808 < fi10)))) {
                                                        var v16 uint = 2
                                                        var v17 uint16 = v2
                                                        _ = v16
                                                        _ = v17
                                                    } else {
                                                        if false {
                                                            v2 = (v2 | 0)
                                                        }
                                                        var v18 uint64 = 150
                                                        switch (fi10 | -1) {
                                                        case 462:
                                                            if true {
                                                                if true {
                                                                    v18 = v18
                                                                }
                                                            }
                                                        case 0:
                                                            var v19 uint32 = uint32(v5)
                                                            _ = v19
                                                        case -458:
                                                            v2 = (v2 / ((v2) | 1))
                                                            var v20 string = ("hello" + "a")
                                                            v18 = (v18 & 4294967295)
                                                            _ = v20
                                                        default:
                                                            v21 := GenS3{F0: 4294967295}
                                                            _ = v21
                                                        }
                                                        _ = v18
                                                    }
                                                    _ = fi15
                                                }
                                                if false {
                                                    for fi22 := 0; fi22 < 19; fi22++ {
                                                        v5 -= (v5 + (^(v5)))
                                                        var v23 float64 = 1e10
                                                        _ = fi22
                                                        _ = v23
                                                    }
                                                } else {
                                                    switch "" {
                                                    case "a":
                                                        var v24 uint16 = (v2 + uint16(v5))
                                                        _ = v24
                                                    case "\n\t":
                                                        v5 -= (v5 & (v5 | 92))
                                                    default:
                                                        v25 := v1[42]
                                                        _ = v25
                                                    }
                                                }
                                                var v26 int16 = 32767
                                                _ = v26
                                            }
                                        }
                                        var v27 string = (("hello" + "foo bar") + "a")
                                        _ = v27
                                    } else {
                                        v28 := GenS3{F0: 18446744073709551615}
                                        switch (fi10 | (fi10 & 128)) {
                                        case 980:
                                            switch 1 {
                                            case 4294967295:
                                                v2 -= v2
                                            }
                                        }
                                        _ = v28
                                    }
                                case "\n\t":
                                    var v29 float32 = (((((1e10 - float32(((3.14159 - (-2.25 * (1e10 / 1.0))) * 0.0))) * (-2.25 + 1e10)) / (((0.1 * 0.1) + float32((((float64(0.0) * 3.14159) + (float64(0.1) + 3.14159)) * (0.1 + 3.14159)))) - -2.25)) - (-2.25 * float32(1e10))) / 1.0)
                                    _ = v29
                                }
                            } else {
                                v5 = v5
                                for fi30 := 0; fi30 < 14; fi30++ {
                                    switch "a" {
                                    case "\n\t":
                                        var v31 int8 = 100
                                        _ = v31
                                    case "hello":
                                        v2 -= (v2 / (((v2 / (((^(v2))) | 1))) | 1))
                                    default:
                                        fmt.Println(v5)
                                    }
                                    switch "hello" {
                                    case "foo bar":
                                        if (!(fi10 <= fi30)) {
                                            fmt.Println(fi10)
                                        }
                                    case "":
                                        delete(v1, 1)
                                    }
                                    _ = fi30
                                }
                                v2 -= 42
                            }
                            switch ((("\n\t" + "") + "hello") + "foo bar") {
                            case "a":
                                v32 := v1[2147483647]
                                _ = v32
                            case "\n\t":
                                for fi33 := 0; fi33 < 15; fi33++ {
                                    v34 := GenS1{F0: 0, F1: 0, F2: 42}
                                    fmt.Println(fi10, fi33, fi33)
                                    _ = fi33
                                    _ = v34
                                }
                            case "hello":
                                var v35 string = ("foo bar" + ((((("" + "hello") + "foo bar") + ("\n\t" + (("hello" + ("" + (((("\n\t" + "a") + "foo bar") + "a") + (("\n\t" + ("foo bar" + "hello")) + "foo bar")))) + "a"))) + "hello") + "foo bar"))
                                _ = v35
                            }
                            fmt.Println(v2)
                        } else {
                            if false {
                                _ = len(v1)
                            }
                        }
                        for fi36 := 0; fi36 < 7; fi36++ {
                            v2 += (v2 - (v2 * (v2 & v2)))
                            _ = fi36
                        }
                    }
                    if true {
                        var v37 int16 = 1000
                        _ = v37
                    } else {
                        v38 := v1[561]
                        v39 := GenS3{F0: 2}
                        _ = v38
                        _ = v39
                    }
                    if len(v3) > 0 {
                        v3[0] = 1
                    }
                }
                v5 = (v5 << ((18446744073709551615) & 63))
                _ = fi10
            }
        case "foo bar":
            switch ("\n\t" + "hello") {
            case "":
                for fi40 := 0; fi40 < 2; fi40++ {
                    if ((110 > 415) && (true && true)) {
                        if true {
                            for fi41 := 0; fi41 < 12; fi41++ {
                                v5 += 1
                                if (!(104 == 127)) {
                                    v5 = genFn0(0)
                                    if false {
                                        v2 = (v2 & v2)
                                    }
                                }
                                _ = fi41
                            }
                        } else {
                            v5 = (v5 & v5)
                        }
                        switch "\n\t" {
                        case "":
                            fmt.Println(v5, v2, fi40)
                        case "a":
                            var v42 uint8 = 149
                            _ = v42
                        default:
                            fmt.Println(v2)
                        }
                    }
                    v43 := GenS1{F0: 2, F1: (^(v2)), F2: 4294967295}
                    _ = fi40
                    _ = v43
                }
            case "foo bar":
                if true {
                    for fi44 := 0; fi44 < 14; fi44++ {
                        v45 := v1[42]
                        var v46 float64 = 0.1
                        var v47 float32 = ((3.14159 / 1e10) / 1.0)
                        _ = fi44
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                    for fi48 := 0; fi48 < 11; fi48++ {
                        var v49 uint32 = 395
                        for fi50 := 0; fi50 < 11; fi50++ {
                            _ = len(v3)
                            _ = fi50
                        }
                        _ = fi48
                        _ = v49
                    }
                }
            case "hello":
                v51 := GenS2{F0: genFn0(724), F1: 928, F2: uint32(v2)}
                _ = v51
            }
        }
        if (((0.1 + 1e10) > 1.5) || (v5 > (v5 + 184))) {
            v52 := GenS3{F0: 769}
            var v53 uint32 = 753
            if len(v3) > 0 {
                v3[0] = (v53 >> ((255) & 63))
            }
            _ = v52
            _ = v53
        }
    }
    switch (("\n\t" + "a") + "\n\t") {
    case "hello":
        v5 = (^(v5))
    default:
        v54 := GenS1{F0: (v5 ^ (v5 | 1)), F1: (^(v2)), F2: 255}
        _ = v54
    }
    v5 = (v5 & (^(v5)))
    fmt.Println(v2)
    var v55 int64 = 7
    v56 := GenS4{F0: -655, F1: 127}
    var v57 int16 = -32768
    for fi58 := 0; fi58 < 5; fi58++ {
        if false {
            fmt.Println(v2)
        } else {
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
        }
        v59 := GenS5{F0: 100, F1: 306}
        if true {
            v2 -= (v2 & v2)
            var v60 string = ("\n\t" + ("foo bar" + ""))
            _ = v60
        }
        _ = fi58
        _ = v59
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v55
    _ = v56
    _ = v57
    _ = v6
}
