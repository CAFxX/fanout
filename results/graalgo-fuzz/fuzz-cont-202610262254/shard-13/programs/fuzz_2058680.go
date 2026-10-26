package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int8; F1 uint16; F2 int32 }
type GenS5 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS6 struct { F0 int8; F1 uint64 }
type GenS7 struct { F0 int8; F1 string; F2 int32 }
func genFn0(p0 bool, p1 uint, p2 uint32) int32 {
    var v4 int8 = 100
    p0 = (p0 || (((65535 > 480) || (!p0)) && (-1 > 1)))
    var v5 bool = (18446744073709551615 > 255)
    for fi6 := 0; fi6 < 2; fi6++ {
        v4 = -106
        _ = fi6
    }
    p1 = (p1 >> (((^(p1))) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 int16) uint32 {
    fmt.Println(p0)
    var v7 uint32 = 465
    v8 := GenS1{F0: (v7 ^ (^(v7)))}
    v9 := []int{42}
    var v10 int64 = 0
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return (v7 ^ (v7 / ((v7) | 1)))
}

func main() {
    var v1 int = 9223372036854775807
    var v2 uint16 = 32767
    var v3 int32 = -1
    v11 := []int{(v1 % (((v1 - (v1 + v1))) | 1))}
    if true {
        v3 = -2147483648
        v12 := GenS1{F0: 2147483648}
        _ = v12
    }
    if ("\n\t" == "") {
        switch (v1 ^ (v1 >> ((255) & 63))) {
        case 807:
            if (false || true) {
                var v13 uint32 = 394
                var v14 uint = uint(v2)
                _ = v13
                _ = v14
            } else {
                fmt.Println(v2, v1)
            }
        case 32768:
            var v15 float64 = 3.14159
            _ = v15
        }
    }
    if len(v11) > 0 {
        v11[0] = v1
    }
    if false {
        var v16 uint = 2
        _ = v16
    }
    switch "hello" {
    case "foo bar":
        v3 -= -401
    case "hello":
        v3 = (v3 * (v3 | v3))
    }
    v3 += (v3 | v3)
    var v17 bool = (true && (("a" + ("" + "a")) == ("hello" + "")))
    v17 = v17
    if ((v2 << ((542) & 63)) < v2) {
        for fi18 := 0; fi18 < 19; fi18++ {
            if (false || v17) {
                v3 = (-(v3))
            } else {
                switch "\n\t" {
                case "hello":
                    if ((!v17) || v17) {
                        v19 := GenS2{F0: genFn1(int16(v3))}
                        var v20 float64 = (float64(((1.5 + 1e10) - -2.25)) / 0.1)
                        v11 = append(v11, fi18)
                        _ = v19
                        _ = v20
                    }
                case "":
                    for fi21 := 0; fi21 < 20; fi21++ {
                        v17 = (((((("a" + "\n\t") + "") != "hello") && v17) || ((1.5 + 1.5) != float64(1e10))) || false)
                        _ = fi21
                    }
                }
            }
            switch "" {
            case "\n\t":
                v1 = (v1 * v1)
            case "a":
                v22 := GenS3{F0: "foo bar", F1: -53}
                _ = v22
            case "hello":
                v3 = (v3 & -2147483648)
            default:
                switch "foo bar" {
                case "hello":
                    for fi23 := 0; fi23 < 20; fi23++ {
                        fmt.Println(v2, fi18, fi18)
                        v1 -= -2147483648
                        _ = fi23
                    }
                default:
                    v24 := GenS4{F0: 78, F1: (v2 ^ 65535), F2: (v3 >> ((2) & 63))}
                    _ = v24
                }
            }
            if true {
                var v25 uint16 = v2
                var v26 bool = false
                _ = v25
                _ = v26
            } else {
                var v27 uint8 = 253
                _ = v27
            }
            _ = fi18
        }
        var v28 int64 = -438
        if v17 {
            if (!v17) {
                v28 = v28
                v3 += (v3 | (v3 >> ((1) & 63)))
            }
        } else {
            switch ("\n\t" + ("foo bar" + "")) {
            case "foo bar":
                for fi29 := 0; fi29 < 1; fi29++ {
                    v30 := GenS5{F0: v17, F1: 2, F2: v2}
                    var v31 int32 = v3
                    for fi32 := 0; fi32 < 10; fi32++ {
                        switch (("a" + ("" + "")) + "\n\t") {
                        case "hello":
                            v33 := GenS1{F0: 4294967295}
                            _ = v33
                        case "foo bar":
                            if len(v11) > 0 {
                                v11[0] = fi32
                            }
                        default:
                            var v34 uint32 = 4294967295
                            _ = v34
                        }
                        v31 += (v3 & (v3 << ((2) & 63)))
                        _ = fi32
                    }
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            default:
                for fi35 := 0; fi35 < 15; fi35++ {
                    switch "" {
                    case "foo bar":
                        for fi36 := 0; fi36 < 14; fi36++ {
                            v1 += fi36
                            v28 = -393
                            _ = fi36
                        }
                    default:
                        v17 = (v17 && (2 > 1))
                    }
                    v2 = v2
                    if ((-33 > -9223372036854775808) && ((^(v2)) >= v2)) {
                        v17 = ((((float32((0.1 - 0.0)) + (1.5 / float32(float64((0.1 / -2.25))))) != -2.25) || (!v17)) || v17)
                        var v37 int32 = (v3 << ((171) & 63))
                        _ = v37
                    } else {
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        switch "" {
                        case "\n\t":
                            v1 += fi35
                        case "foo bar":
                            v28 -= (v28 << ((40) & 63))
                        default:
                            switch "hello" {
                            case "foo bar":
                                for fi38 := 0; fi38 < 7; fi38++ {
                                    v39 := GenS6{F0: 42, F1: 4294967295}
                                    _ = fi38
                                    _ = v39
                                }
                            case "a":
                                switch ("\n\t" + "foo bar") {
                                case "foo bar":
                                    fmt.Println(v28, fi35, fi35)
                                }
                            }
                        }
                    }
                    _ = fi35
                }
            }
            v40 := GenS4{F0: -1, F1: 86, F2: -1}
            _ = v40
        }
        _ = v28
    } else {
        var v41 uint16 = 932
        _ = len(v11)
        _ = v41
    }
    v42 := GenS3{F0: (("foo bar" + ((("foo bar" + "") + "\n\t") + "a")) + ("a" + "")), F1: 0}
    switch "foo bar" {
    case "a":
        var v43 uint64 = 2
        _ = v43
    case "hello":
        switch (("a" + "foo bar") + "hello") {
        case "":
            v1 += 0
        case "\n\t":
            switch (v1 & 62) {
            case 340:
                v1 = v1
            case -950:
                if len(v11) > 0 {
                    v11[0] = v1
                }
            case -246:
                for fi44 := 0; fi44 < 6; fi44++ {
                    if true {
                        v45 := GenS4{F0: -128, F1: v2, F2: (v3 * (^(v3)))}
                        v2 = v2
                        var v46 uint = 255
                        _ = v45
                        _ = v46
                    } else {
                        if v17 {
                            v47 := GenS4{F0: 109, F1: (v2 - 65535), F2: 737}
                            _ = v47
                        } else {
                            _ = len(v11)
                        }
                        for fi48 := 0; fi48 < 20; fi48++ {
                            v3 = v3
                            v3 -= (v3 % ((genFn0(v17, 904, genFn1(1))) | 1))
                            v49 := GenS4{F0: 100, F1: 0, F2: int32(v2)}
                            _ = fi48
                            _ = v49
                        }
                    }
                    var v50 uint16 = v2
                    _ = fi44
                    _ = v50
                }
            default:
                var v51 bool = ((v17 || ((v1 + (v1 ^ v1)) == (v1 ^ 0))) && (-13 <= 56))
                _ = v51
            }
        case "hello":
            var v52 int64 = -9223372036854775808
            _ = v52
        default:
            for fi53 := 0; fi53 < 3; fi53++ {
                v54 := GenS4{F0: 1, F1: v2, F2: int32(fi53)}
                for fi55 := 0; fi55 < 9; fi55++ {
                    var v56 int16 = -409
                    _ = fi55
                    _ = v56
                }
                var v57 int16 = -32768
                _ = fi53
                _ = v54
                _ = v57
            }
        }
    }
    if len(v11) > 0 {
        v11[0] = 32767
    }
    var v58 int32 = 42
    for fi59 := 0; fi59 < 13; fi59++ {
        v60 := GenS7{F0: 30, F1: ("\n\t" + "foo bar"), F2: genFn0(((v17 && (!(v17 && v17))) && true), 2, 2147483648)}
        v1 += (v1 % ((fi59) | 1))
        v1 = fi59
        _ = fi59
        _ = v60
    }
    v3 = (v3 + -275)
    v11 = append(v11, -810)
    if ((v17 || (-32 > -9223372036854775808)) || v17) {
        for fi61 := 0; fi61 < 12; fi61++ {
            v62 := GenS2{F0: 240}
            _ = fi61
            _ = v62
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v42
    _ = v58
}
