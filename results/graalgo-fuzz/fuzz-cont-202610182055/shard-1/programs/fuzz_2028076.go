package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint32; F1 int8; F2 uint16 }
func genFn0() int8 {
    for fi4 := 0; fi4 < 19; fi4++ {
        var v5 uint16 = 65535
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: -533}
    _ = v6
    return 100
}
func genFn1(p0 uint, p1 uint64, p2 uint8) uint64 {
    p2 -= p2
    if true {
        var v7 int32 = -2147483648
        _ = v7
    }
    v8 := GenS1{F0: 0}
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := GenS1{F0: 1000}
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 uint8 = 1
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 uint8 = (^(p2))
            v15 := GenS2{F0: 18446744073709551615, F1: p1}
            p1 = p1
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 1
}

func main() {
    var v1 float32 = 1e10
    v2 := []string{"hello"}
    v3 := []string{(("a" + "\n\t") + "foo bar"), ((("\n\t" + "a") + ("hello" + ("foo bar" + "\n\t"))) + "foo bar")}
    v1 -= float32(0.1)
    var v16 uint16 = 994
    switch ("a" + "hello") {
    case "a":
        if true {
            v16 += v16
            v16 = v16
        }
    case "hello":
        v17 := GenS3{F0: 81}
        _ = v17
    default:
        fmt.Println(v16, v1)
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        fmt.Println(fi18, v16, v1)
        _ = fi18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        switch "" {
        case "":
            if true {
                for fi20 := 0; fi20 < 2; fi20++ {
                    v21 := GenS1{F0: 32767}
                    fmt.Println(fi19)
                    switch ("a" + "a") {
                    case "foo bar":
                        fmt.Println(v1, v16)
                    case "hello":
                        v1 = 3.14159
                    default:
                        if (((float64((0.1 + (v1 / v1))) + (0.0 / 1.0)) * -2.25) <= (((float64(float32(-2.25)) * 0.0) - 3.14159) / (1e10 * (float64((v1 / (0.0 * v1))) * 3.14159)))) {
                            var v22 bool = true
                            var v23 bool = (0 == 42)
                            var v24 uint16 = (v16 ^ v16)
                            _ = v22
                            _ = v23
                            _ = v24
                        } else {
                            if true {
                                v25 := GenS2{F0: 597, F1: uint64(v16)}
                                v26 := GenS2{F0: 2, F1: 2}
                                v3 = append(v3, ((("a" + "a") + "\n\t") + ""))
                                _ = v25
                                _ = v26
                            } else {
                                _ = len(v2)
                            }
                        }
                    }
                    v27 := GenS2{F0: 0, F1: genFn1(0, 164, 255)}
                    _ = fi20
                    _ = v21
                    _ = v27
                }
                v28 := GenS4{F0: (^(v16))}
                for fi29 := 0; fi29 < 20; fi29++ {
                    switch "hello" {
                    case "":
                        switch "foo bar" {
                        case "foo bar":
                            var v30 int16 = -978
                            if true {
                                var v31 bool = (!true)
                                v31 = ((v31 || v31) && false)
                                _ = len(v3)
                                _ = v31
                            }
                            switch 42 {
                            case 127:
                                var v32 bool = true
                                _ = v32
                            case -1:
                                if true {
                                    v3 = append(v3, "")
                                } else {
                                    v16 += v16
                                }
                            case -395:
                                if len(v3) > 0 {
                                    v3[0] = "\n\t"
                                }
                            default:
                                switch "hello" {
                                case "\n\t":
                                    v16 = 42
                                default:
                                    switch (fi29 ^ 830) {
                                    case -1:
                                        v2 = append(v2, ("" + ""))
                                    case 42:
                                        var v33 int64 = 9223372036854775807
                                        _ = v33
                                    }
                                }
                            }
                            _ = v30
                        case "a":
                            fmt.Println(fi29)
                        case "\n\t":
                            v34 := GenS5{F0: 227, F1: 2147483648}
                            _ = v34
                        }
                    case "a":
                        _ = len(v2)
                    default:
                        var v35 int32 = -2147483648
                        _ = v35
                    }
                    _ = fi29
                }
                _ = v28
            }
        case "hello":
            for fi36 := 0; fi36 < 8; fi36++ {
                v16 = v16
                _ = fi36
            }
        }
        var v37 uint = 31
        fmt.Println(fi19)
        _ = fi19
        _ = v37
    }
    if false {
        switch "hello" {
        case "foo bar":
            switch 32767 {
            case 989:
                fmt.Println(v16, v16)
            case 0:
                for fi38 := 0; fi38 < 8; fi38++ {
                    var v39 float32 = v1
                    if (("hello" + "") != ("" + "\n\t")) {
                        v40 := GenS6{F0: (fi38 >> ((42) & 63))}
                        v16 += 2
                        v1 = ((v1 / v1) * (1e10 + 1.5))
                        _ = v40
                    }
                    var v41 int16 = -1
                    _ = fi38
                    _ = v39
                    _ = v41
                }
            case 556:
                fmt.Println(v16, v1)
            }
        case "hello":
            v16 = 2
        case "\n\t":
            var v42 string = ("a" + ("foo bar" + "a"))
            _ = v42
        }
        v43 := GenS1{F0: -139}
        _ = v43
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if false {
            switch -2147483648 {
            case 32768:
                var v44 int = 2147483647
                _ = v44
            }
            fmt.Println(v1, v16)
            var v45 int16 = 1000
            _ = v45
        } else {
            switch 0 {
            case 4294967295:
                var v46 uint16 = (v16 % ((1000) | 1))
                _ = v46
            case -1:
                v1 += v1
            }
            var v47 int32 = 1
            v16 = (v16 << ((118) & 63))
            _ = v47
        }
    }
    fmt.Println(v1, v16)
    fmt.Println(v1)
    v1 -= 1.5
    v48 := GenS2{F0: 405, F1: 24}
    switch "\n\t" {
    case "foo bar":
        for fi49 := 0; fi49 < 6; fi49++ {
            if true {
                v50 := GenS3{F0: 42}
                v51 := GenS7{F0: 42}
                switch "" {
                case "hello":
                    switch "foo bar" {
                    case "foo bar":
                        switch "\n\t" {
                        case "hello":
                            for fi52 := 0; fi52 < 14; fi52++ {
                                var v53 bool = (!(-280 != 128))
                                v54 := GenS8{F0: int64(fi49)}
                                v55 := GenS6{F0: 2147483647}
                                _ = fi52
                                _ = v53
                                _ = v54
                                _ = v55
                            }
                        default:
                            switch "hello" {
                            case "a":
                                switch fi49 {
                                case -2147483648:
                                    v16 = v16
                                }
                            case "\n\t":
                                fmt.Println(v1, v16)
                            }
                        }
                    case "hello":
                        var v56 int8 = -1
                        _ = v56
                    case "a":
                        var v57 int8 = -128
                        _ = v57
                    default:
                        v1 -= 1e10
                    }
                case "a":
                    v58 := GenS9{F0: false}
                    _ = v58
                default:
                    v1 -= float32(1e10)
                }
                _ = v50
                _ = v51
            } else {
                for fi59 := 0; fi59 < 7; fi59++ {
                    var v60 uint64 = 1
                    var v61 uint16 = (v16 * (^(v16)))
                    _ = fi59
                    _ = v60
                    _ = v61
                }
            }
            var v62 bool = (((v16 ^ v16) >= (v16 * 1)) && (!false))
            _ = fi49
            _ = v62
        }
    case "a":
        switch -368 {
        case 944:
            for fi63 := 0; fi63 < 9; fi63++ {
                fmt.Println(v16)
                v64 := GenS10{F0: (fi63 % ((fi63) | 1))}
                _ = fi63
                _ = v64
            }
        case 0:
            var v65 uint32 = 42
            _ = v65
        default:
            v16 = (^(v16))
        }
    case "":
        v1 += v1
    }
    if (("\n\t" + (("" + "hello") + ("foo bar" + ""))) < "\n\t") {
        for fi66 := 0; fi66 < 15; fi66++ {
            _ = len(v2)
            for fi67 := 0; fi67 < 4; fi67++ {
                switch (fi67 % (((fi67 << ((234) & 63))) | 1)) {
                case -2147483648:
                    v1 = (float32(float64(0.0)) + (v1 + v1))
                case 42:
                    if true {
                        _ = len(v2)
                    } else {
                        switch (-(fi66)) {
                        case 740:
                            switch "a" {
                            case "foo bar":
                                fmt.Println(v16, fi66)
                            case "a":
                                v68 := GenS11{F0: 2147483648, F1: 1, F2: 954}
                                _ = v68
                            default:
                                switch 32768 {
                                case -2147483648:
                                    v69 := GenS5{F0: 2, F1: 1}
                                    _ = v69
                                case 525:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case 685:
                                    v70 := GenS6{F0: (-(fi67))}
                                    _ = v70
                                }
                            }
                        case 359:
                            if len(v2) > 0 {
                                v2[0] = "foo bar"
                            }
                        default:
                            var v71 bool = false
                            _ = v71
                        }
                        var v72 float64 = float64(v1)
                        fmt.Println(fi67)
                        _ = v72
                    }
                case 2147483647:
                    v16 += v16
                default:
                    if false {
                        var v73 uint64 = 42
                        var v74 uint32 = 4294967295
                        fmt.Println(v16)
                        _ = v73
                        _ = v74
                    }
                }
                _ = fi67
            }
            for fi75 := 0; fi75 < 16; fi75++ {
                var v76 float64 = -2.25
                v76 = v76
                v16 = 2
                _ = fi75
                _ = v76
            }
            _ = fi66
        }
        v16 = v16
        var v77 bool = true
        _ = v77
    }
    v78 := GenS4{F0: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v48
    _ = v78
}
