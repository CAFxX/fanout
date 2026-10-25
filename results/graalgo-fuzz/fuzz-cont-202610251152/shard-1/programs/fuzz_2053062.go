package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 bool; F2 uint }
type GenS5 struct { F0 int8; F1 bool; F2 uint8 }
type GenS6 struct { F0 uint8 }

func main() {
    var v1 uint32 = 65535
    v2 := make(map[int]bool)
    v3 := []uint32{(^(v1)), v1}
    v4 := make(map[uint32]bool)
    v5 := []string{"foo bar", (("\n\t" + "") + "foo bar")}
    var v6 int16 = 555
    if (!((!(true || true)) && false)) {
        v1 += v1
        _ = len(v2)
    } else {
        var v7 int = -971
        if (true || (true || ((((28 != -474) && false) || (2 <= 1)) || ((^(v1)) >= v1)))) {
            var v8 uint64 = 18446744073709551615
            _ = v8
        } else {
            if true {
                var v9 uint64 = 0
                _ = len(v4)
                _ = v9
            } else {
                var v10 int8 = 127
                _ = v10
            }
        }
        fmt.Println(v6, v1, v7)
        _ = v7
    }
    v11 := GenS1{F0: (v6 * 32767)}
    v1 = (v1 % ((v1) | 1))
    if false {
        switch 4294967295 {
        case 2147483647:
            var v12 int = 128
            _ = v12
        case 127:
            var v13 uint64 = 4294967295
            _ = v13
        default:
            v6 += (v6 * -402)
        }
        for fi14 := 0; fi14 < 8; fi14++ {
            if false {
                switch ("\n\t" + "foo bar") {
                case "\n\t":
                    for fi15 := 0; fi15 < 16; fi15++ {
                        var v16 uint64 = 881
                        if false {
                            if true {
                                v1 = v1
                            }
                        }
                        var v17 float64 = float64(-2.25)
                        v16 -= (v16 % ((v16) | 1))
                        fmt.Println(v16, v6, v17)
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                case "":
                    for fi18 := 0; fi18 < 5; fi18++ {
                        v4[65535] = (2147483647 <= 2147483647)
                        switch 212 {
                        case -73:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        case 0:
                            if true {
                                if (176 != 42) {
                                    var v19 float32 = (-2.25 - -2.25)
                                    _ = len(v4)
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = v19
                                } else {
                                    var v20 uint64 = 547
                                    v20 = v20
                                    v6 += v6
                                    _ = v20
                                }
                                v21 := GenS1{F0: v6}
                                _ = v21
                            } else {
                                v22 := GenS1{F0: (v6 >> ((0) & 63))}
                                v23 := GenS2{F0: 839, F1: 2}
                                _ = v22
                                _ = v23
                            }
                        case 42:
                            fmt.Println(fi18, fi18, fi18)
                        default:
                            var v24 string = "foo bar"
                            _ = v24
                        }
                        var v25 int8 = -1
                        _ = fi18
                        _ = v25
                    }
                }
                switch 9223372036854775807 {
                case -540:
                    switch "foo bar" {
                    case "foo bar":
                        if ((!(65535 > 2)) || true) {
                            for fi26 := 0; fi26 < 19; fi26++ {
                                var v27 uint16 = uint16(v1)
                                if ("a" >= "a") {
                                    var v28 int32 = -1
                                    v27 = v27
                                    switch fi14 {
                                    case 9223372036854775807:
                                        if false {
                                            fmt.Println(fi14, fi26, fi14)
                                        } else {
                                            var v29 float32 = 0.0
                                            v6 -= -495
                                            _ = v29
                                        }
                                    }
                                    _ = v28
                                } else {
                                    switch fi26 {
                                    case 552:
                                        if true {
                                            _ = len(v5)
                                            if (2 > uint8(v1)) {
                                                var v30 float32 = float32(0.0)
                                                v31 := GenS1{F0: (v6 | (v6 - v6))}
                                                var v32 int64 = -794
                                                v33 := GenS2{F0: (^(v27)), F1: 0}
                                                if false {
                                                    v34 := GenS1{F0: v6}
                                                    var v35 int16 = v6
                                                    fmt.Println(v35)
                                                    _ = v34
                                                    _ = v35
                                                }
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                                _ = v33
                                            } else {
                                                if (!(-2.25 < (1.5 * -2.25))) {
                                                    v1 += 829
                                                    v27 -= 2
                                                    v36 := GenS3{F0: 738}
                                                    _ = v36
                                                }
                                                v27 += (v27 + v27)
                                                v27 -= 361
                                            }
                                        }
                                    case 32767:
                                        v37 := GenS2{F0: (v27 ^ v27), F1: 253}
                                        _ = v37
                                    }
                                    if true {
                                        var v38 float32 = 1e10
                                        var v39 float64 = 1e10
                                        _ = v38
                                        _ = v39
                                    }
                                }
                                _ = fi26
                                _ = v27
                            }
                        }
                    case "":
                        v1 = (^(v1))
                    }
                case 4294967295:
                    for fi40 := 0; fi40 < 11; fi40++ {
                        switch 915 {
                        case 128:
                            if len(v5) > 0 {
                                v5[0] = "hello"
                            }
                        case 4294967295:
                            v41 := GenS4{F0: (true && true), F1: true, F2: 18446744073709551615}
                            _ = v41
                        default:
                            v6 = (^(v6))
                        }
                        _ = fi40
                    }
                case 969:
                    var v42 float32 = 1.5
                    _ = v42
                default:
                    v43 := GenS5{F0: 1, F1: (2 < 42), F2: 42}
                    _ = v43
                }
            }
            switch (("a" + "\n\t") + "\n\t") {
            case "foo bar":
                switch (("a" + ((("\n\t" + "foo bar") + ("foo bar" + "hello")) + "hello")) + ("foo bar" + ("\n\t" + "foo bar"))) {
                case "\n\t":
                    var v44 int16 = 718
                    _ = v44
                case "":
                    var v45 int16 = -709
                    _ = v45
                default:
                    var v46 string = ""
                    _ = v46
                }
            case "\n\t":
                v1 = v1
            case "":
                v6 += -1
            default:
                v1 += v1
            }
            _ = fi14
        }
        v6 += v6
    } else {
        for fi47 := 0; fi47 < 6; fi47++ {
            for fi48 := 0; fi48 < 13; fi48++ {
                if true {
                    v6 = v6
                    fmt.Println(v6, fi47, fi47)
                    if ((((1.5 / 1.0) - float64(3.14159)) - 1.5) >= 0.0) {
                        switch "a" {
                        case "hello":
                            v49 := GenS5{F0: -30, F1: ((65535 < 0) && true), F2: 255}
                            _ = v49
                        default:
                            switch "a" {
                            case "hello":
                                var v50 string = "foo bar"
                                var v51 int = (^(fi48))
                                _ = v50
                                _ = v51
                            case "foo bar":
                                v1 = v1
                                fmt.Println(v6, v6, fi47)
                            }
                        }
                        v6 += (v6 / ((v6) | 1))
                    }
                }
                var v52 bool = true
                var v53 int32 = -700
                _ = fi48
                _ = v52
                _ = v53
            }
            v6 = (v6 + (^(v6)))
            _ = fi47
        }
        var v54 int16 = -459
        v55 := GenS6{F0: 42}
        _ = v54
        _ = v55
    }
    delete(v4, v1)
    fmt.Println(v6, v1)
    switch -1 {
    case 0:
        switch -1 {
        case -561:
            v6 -= (v6 - (v6 << ((256) & 63)))
        case 2147483647:
            var v56 uint = 2
            _ = v56
        case 109:
            var v57 uint8 = 0
            _ = v57
        }
    case 9223372036854775807:
        v1 = v1
    case 804:
        fmt.Println(v6, v1)
    default:
        switch 128 {
        case -435:
            if ((!(!false)) || (false || (!false))) {
                switch 9223372036854775807 {
                case 0:
                    v6 = (v6 / (((v6 ^ 0)) | 1))
                default:
                    v58 := GenS6{F0: 71}
                    _ = v58
                }
                delete(v4, 440)
            }
        case 42:
            v1 -= (v1 << ((4294967295) & 63))
        default:
            switch "foo bar" {
            case "hello":
                for fi59 := 0; fi59 < 7; fi59++ {
                    fmt.Println(v6, v6, v1)
                    v3 = append(v3, v1)
                    var v60 uint32 = (v1 >> ((0) & 63))
                    _ = fi59
                    _ = v60
                }
            default:
                v61 := GenS2{F0: 2, F1: 2}
                _ = v61
            }
        }
    }
    var v62 bool = ((0.1 - 3.14159) != 3.14159)
    v63 := GenS5{F0: 1, F1: (307 < 529), F2: 195}
    if v62 {
        for fi64 := 0; fi64 < 2; fi64++ {
            var v65 uint = 702
            v66 := GenS4{F0: ((v1 / ((257) | 1)) == (v1 - uint32(v65))), F1: v62, F2: (v65 ^ (v65 % ((0) | 1)))}
            switch "\n\t" {
            case "hello":
                v67 := GenS1{F0: v6}
                _ = v67
            case "foo bar":
                var v68 int32 = 2147483647
                _ = v68
            default:
                _ = len(v3)
            }
            _ = fi64
            _ = v65
            _ = v66
        }
        for fi69 := 0; fi69 < 11; fi69++ {
            for fi70 := 0; fi70 < 20; fi70++ {
                v62 = ((true || (1e10 == (1.5 * 1e10))) || v62)
                _ = fi70
            }
            _ = fi69
        }
    }
    if true {
        v71 := GenS2{F0: 0, F1: 127}
        _ = v71
    }
    if ("foo bar" <= "a") {
        v2[2147483647] = (v62 || (v62 && v62))
        var v72 uint32 = 2
        _ = v72
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v63
}
