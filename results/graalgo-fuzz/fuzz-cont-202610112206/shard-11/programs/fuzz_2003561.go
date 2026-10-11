package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int64; F1 uint; F2 int }
type GenS4 struct { F0 uint; F1 int16; F2 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool; F1 int32; F2 uint64 }
type GenS7 struct { F0 uint64; F1 bool }
type GenS8 struct { F0 bool; F1 uint; F2 int32 }
type GenS9 struct { F0 int8; F1 int64; F2 int64 }

func main() {
    var v1 uint32 = 2147483648
    var v2 uint = 2
    var v3 uint16 = 110
    v4 := []string{("hello" + "\n\t"), ("" + "hello"), ("" + "\n\t")}
    fmt.Println(v3, v2)
    if len(v4) > 0 {
        v4[0] = ("" + "hello")
    }
    var v5 uint32 = (v1 >> (((^(v2))) & 63))
    if ((v1 - (v1 / ((65535) | 1))) > v5) {
        if ("foo bar" > ("foo bar" + "\n\t")) {
            switch -783 {
            case 0:
                var v6 uint64 = 4294967295
                _ = v6
            case 42:
                for fi7 := 0; fi7 < 14; fi7++ {
                    switch fi7 {
                    case -1:
                        var v8 int = (fi7 * 601)
                        _ = v8
                    case 127:
                        _ = len(v4)
                    }
                    for fi9 := 0; fi9 < 12; fi9++ {
                        v10 := GenS1{F0: (v5 % ((v1) | 1)), F1: (fi7 >> ((v2) & 63)), F2: fi9}
                        _ = fi9
                        _ = v10
                    }
                    v3 = v3
                    _ = fi7
                }
            default:
                v1 = (v1 / ((65535) | 1))
            }
            switch 9223372036854775807 {
            case 2147483647:
                _ = len(v4)
            }
            if false {
                var v11 int8 = 100
                fmt.Println(v2)
                _ = v11
            }
        } else {
            fmt.Println(v2)
            var v12 float32 = 1e10
            _ = v12
        }
        var v13 uint32 = v1
        v2 -= v2
        _ = v13
    }
    switch "\n\t" {
    case "hello":
        switch 128 {
        case 127:
            v1 = v1
        default:
            switch 121 {
            case 994:
                fmt.Println(v2, v5, v2)
            case -698:
                switch 1 {
                case 529:
                    var v14 string = ((("foo bar" + "a") + "hello") + "")
                    _ = v14
                case 9223372036854775807:
                    switch 802 {
                    case 128:
                        var v15 string = "a"
                        _ = v15
                    }
                }
            }
        }
    case "a":
        var v16 float32 = 3.14159
        _ = v16
    case "foo bar":
        switch "foo bar" {
        case "\n\t":
            for fi17 := 0; fi17 < 2; fi17++ {
                var v18 uint = 65535
                fmt.Println(v18)
                fmt.Println(fi17)
                _ = fi17
                _ = v18
            }
        case "a":
            if true {
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
                for fi19 := 0; fi19 < 4; fi19++ {
                    v20 := GenS2{F0: 255, F1: 1}
                    fmt.Println(fi19, v2)
                    _ = fi19
                    _ = v20
                }
                if false {
                    v21 := GenS2{F0: 173, F1: 2}
                    var v22 uint32 = v5
                    _ = v21
                    _ = v22
                } else {
                    v23 := GenS3{F0: -1, F1: v2, F2: 2147483647}
                    v5 = (v1 & 1)
                    _ = v23
                }
            } else {
                for fi24 := 0; fi24 < 4; fi24++ {
                    v2 = (v2 - (v2 % (((v2 % (((^(v2))) | 1))) | 1)))
                    if true {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    } else {
                        for fi25 := 0; fi25 < 9; fi25++ {
                            switch "hello" {
                            case "a":
                                v26 := GenS3{F0: 42, F1: (v2 / ((v2) | 1)), F2: 1}
                                var v27 uint = 65535
                                _ = v26
                                _ = v27
                            case "\n\t":
                                var v28 uint64 = 42
                                _ = v28
                            case "foo bar":
                                v1 -= (v1 >> ((11) & 63))
                            default:
                                v1 += (v1 >> ((v2) & 63))
                            }
                            switch "" {
                            case "a":
                                switch (-(fi24)) {
                                case 9223372036854775807:
                                    var v29 uint8 = 2
                                    _ = v29
                                default:
                                    switch (fi25 >> ((34) & 63)) {
                                    case -9223372036854775808:
                                        v30 := GenS3{F0: -1, F1: v2, F2: fi25}
                                        _ = v30
                                    case -1:
                                        if false {
                                            switch (fi25 | (fi24 | 748)) {
                                            case 128:
                                                v3 += 169
                                            case -564:
                                                v3 += v3
                                            }
                                            var v31 float32 = 0.0
                                            _ = v31
                                        } else {
                                            v4 = append(v4, ("a" + ("" + "")))
                                        }
                                    case -149:
                                        switch "hello" {
                                        case "foo bar":
                                            if (false && true) {
                                                v3 += v3
                                                fmt.Println(v2)
                                            }
                                        default:
                                            v1 -= 196
                                        }
                                        _ = len(v4)
                                        var v32 int8 = -1
                                        var v33 uint = v2
                                        _ = v32
                                        _ = v33
                                    default:
                                        v34 := GenS1{F0: (v5 + v1), F1: (fi25 >> (((v2 / (((^(v2))) | 1))) & 63)), F2: (fi25 + (fi25 / ((fi24) | 1)))}
                                        _ = v34
                                    }
                                }
                            case "\n\t":
                                fmt.Println(v5, v3, fi25)
                            }
                            switch "" {
                            case "":
                                v3 = v3
                            case "\n\t":
                                v35 := GenS3{F0: -52, F1: (v2 >> ((1) & 63)), F2: (fi24 + (fi25 - int(v3)))}
                                _ = v35
                            case "a":
                                v36 := GenS2{F0: 0, F1: 4294967295}
                                _ = v36
                            }
                            _ = fi25
                        }
                        _ = len(v4)
                        v4 = append(v4, "")
                    }
                    v3 = 32767
                    _ = fi24
                }
            }
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    for fi37 := 0; fi37 < 7; fi37++ {
        for fi38 := 0; fi38 < 6; fi38++ {
            if (!true) {
                if (-1 <= 1) {
                    if true {
                        if (127 > 0) {
                            if (((!true) || (0 > 18446744073709551615)) || false) {
                                var v39 uint16 = 1
                                _ = v39
                            }
                            v5 = v5
                            var v40 uint64 = 2
                            _ = v40
                        } else {
                            v5 = v1
                            _ = len(v4)
                            v5 -= (v1 + v5)
                        }
                        _ = len(v4)
                    }
                    v41 := GenS4{F0: (v2 >> ((2) & 63)), F1: 1, F2: "foo bar"}
                    _ = v41
                }
                v4 = append(v4, "")
            }
            _ = fi38
        }
        var v42 float64 = 1.5
        v1 -= 800
        _ = fi37
        _ = v42
    }
    if ((!false) || true) {
        if (152 < 162) {
            v5 += 65535
            for fi43 := 0; fi43 < 11; fi43++ {
                v1 += v1
                _ = fi43
            }
            if false {
                for fi44 := 0; fi44 < 4; fi44++ {
                    v1 = (^(v5))
                    _ = fi44
                }
            } else {
                v3 += v3
                if true {
                    if true {
                        fmt.Println(v5, v2, v1)
                        switch "hello" {
                        case "a":
                            fmt.Println(v5, v3, v3)
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    for fi45 := 0; fi45 < 15; fi45++ {
                        if len(v4) > 0 {
                            v4[0] = "hello"
                        }
                        fmt.Println(v1, fi45, fi45)
                        fmt.Println(v2)
                        _ = fi45
                    }
                    v46 := GenS5{F0: 2147483647}
                    _ = v46
                } else {
                    var v47 bool = false
                    v3 += 65535
                    v48 := GenS5{F0: 209}
                    _ = v47
                    _ = v48
                }
                if true {
                    if true {
                        var v49 int = 4294967295
                        for fi50 := 0; fi50 < 4; fi50++ {
                            v1 = v1
                            var v51 uint16 = (v3 - (v3 * (v3 * (v3 + 32767))))
                            _ = fi50
                            _ = v51
                        }
                        switch -92 {
                        case -1:
                            v49 = v49
                        case 32768:
                            switch (v49 + 0) {
                            case 163:
                                _ = len(v4)
                            case 9223372036854775807:
                                switch v49 {
                                case 2147483647:
                                    if ((!true) && (0 > 112)) {
                                        var v52 string = ("a" + "a")
                                        var v53 int32 = -2147483648
                                        _ = v52
                                        _ = v53
                                    } else {
                                        v54 := GenS6{F0: (!true), F1: -950, F2: 42}
                                        _ = v54
                                    }
                                }
                            }
                        }
                        _ = v49
                    }
                    for fi55 := 0; fi55 < 20; fi55++ {
                        fmt.Println(v5, v5)
                        var v56 int8 = -1
                        if ((-231 > 1) && ((v2 | v2) >= 2)) {
                            if (3.14159 > 3.14159) {
                                if (v2 < 192) {
                                    v56 = (^(v56))
                                    fmt.Println(v1, v56)
                                    var v57 string = ("\n\t" + ("a" + ""))
                                    _ = v57
                                } else {
                                    if true {
                                        var v58 int8 = v56
                                        v59 := GenS2{F0: 128, F1: 558}
                                        for fi60 := 0; fi60 < 15; fi60++ {
                                            v1 = v1
                                            _ = fi60
                                        }
                                        _ = v58
                                        _ = v59
                                    } else {
                                        var v61 int8 = v56
                                        _ = v61
                                    }
                                    switch ("foo bar" + "a") {
                                    case "a":
                                        for fi62 := 0; fi62 < 16; fi62++ {
                                            var v63 int32 = 1
                                            var v64 int8 = -100
                                            v4 = append(v4, "foo bar")
                                            v2 += (v2 + (v2 / ((4294967295) | 1)))
                                            _ = fi62
                                            _ = v63
                                            _ = v64
                                        }
                                    case "":
                                        if (false && (1 != 42)) {
                                            var v65 uint64 = 2
                                            v2 += v2
                                            for fi66 := 0; fi66 < 4; fi66++ {
                                                v3 = uint16(v65)
                                                _ = fi66
                                            }
                                            _ = v65
                                        } else {
                                            fmt.Println(v56, v1)
                                            var v67 bool = (-32768 <= -275)
                                            _ = v67
                                        }
                                    case "\n\t":
                                        v68 := GenS6{F0: false, F1: -1, F2: 42}
                                        _ = v68
                                    default:
                                        if true {
                                            if false {
                                                for fi69 := 0; fi69 < 4; fi69++ {
                                                    fmt.Println(v5, v1)
                                                    _ = fi69
                                                }
                                                if len(v4) > 0 {
                                                    v4[0] = ("" + "a")
                                                }
                                                fmt.Println(v56, v5, v3)
                                            }
                                        }
                                    }
                                }
                            }
                            var v70 uint16 = v3
                            if true {
                                v71 := GenS7{F0: 42, F1: false}
                                _ = v71
                            } else {
                                if (!(1e10 != 0.0)) {
                                    if len(v4) > 0 {
                                        v4[0] = "foo bar"
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    if false {
                                        if true {
                                            _ = len(v4)
                                        } else {
                                            if true {
                                                v72 := GenS4{F0: v2, F1: -711, F2: ""}
                                                switch "foo bar" {
                                                case "":
                                                    v3 = 32767
                                                }
                                                fmt.Println(v3, fi55, v3)
                                                _ = v72
                                            } else {
                                                _ = len(v4)
                                            }
                                        }
                                        v1 = 2147483647
                                        for fi73 := 0; fi73 < 19; fi73++ {
                                            v4 = append(v4, "\n\t")
                                            if true {
                                                if false {
                                                    if false {
                                                        if false {
                                                            v74 := GenS4{F0: 42, F1: 942, F2: ("hello" + "")}
                                                            var v75 float32 = -2.25
                                                            _ = v74
                                                            _ = v75
                                                        }
                                                        var v76 int32 = 42
                                                        _ = v76
                                                    } else {
                                                        v4 = append(v4, ("foo bar" + "a"))
                                                        if len(v4) > 0 {
                                                            v4[0] = "a"
                                                        }
                                                    }
                                                }
                                                fmt.Println(fi73)
                                            } else {
                                                _ = len(v4)
                                                v77 := GenS2{F0: 155, F1: 0}
                                                if len(v4) > 0 {
                                                    v4[0] = "\n\t"
                                                }
                                                if len(v4) > 0 {
                                                    v4[0] = ""
                                                }
                                                _ = v77
                                            }
                                            v5 = (v5 / (((v5 << ((v2) & 63))) | 1))
                                            _ = fi73
                                        }
                                    } else {
                                        var v78 int = (fi55 / (((fi55 * fi55)) | 1))
                                        v79 := GenS1{F0: uint32(v3), F1: int(v2), F2: int(v3)}
                                        v80 := GenS6{F0: (float64(1e10) <= 0.0), F1: 990, F2: 1}
                                        _ = v78
                                        _ = v79
                                        _ = v80
                                    }
                                }
                            }
                            _ = v70
                        } else {
                            if len(v4) > 0 {
                                v4[0] = ""
                            }
                            fmt.Println(v3, v3, v2)
                        }
                        _ = fi55
                        _ = v56
                    }
                }
            }
        } else {
            _ = len(v4)
        }
    } else {
        var v81 uint16 = (v3 % (((v3 * (v3 / (((^(v3))) | 1)))) | 1))
        fmt.Println(v2, v2)
        _ = v81
    }
    if true {
        v4 = append(v4, "a")
        v2 += v2
    } else {
        switch int(v2) {
        case 0:
            switch -548 {
            case -353:
                var v82 int8 = 127
                _ = v82
            case 629:
                if len(v4) > 0 {
                    v4[0] = ("a" + ((((("\n\t" + ((("\n\t" + "foo bar") + "hello") + (("" + "a") + "foo bar"))) + ("hello" + (("foo bar" + ("a" + "\n\t")) + ("" + "\n\t")))) + "hello") + "a") + "a"))
                }
            case 32767:
                if false {
                    fmt.Println(v1, v1, v5)
                    fmt.Println(v3)
                } else {
                    if len(v4) > 0 {
                        v4[0] = "foo bar"
                    }
                    v2 += 838
                    v5 += 65535
                }
            default:
                v5 = (^(v5))
            }
        case -1:
            v1 -= v1
        }
    }
    if len(v4) > 0 {
        v4[0] = ""
    }
    if (-9223372036854775808 <= 782) {
        if true {
            var v83 int16 = 969
            if false {
                for fi84 := 0; fi84 < 5; fi84++ {
                    if (-2.25 != 0.1) {
                        v3 = (v3 - v3)
                    } else {
                        var v85 int32 = -253
                        var v86 int32 = (v85 << (((v2 << ((v2) & 63))) & 63))
                        v87 := GenS8{F0: (!true), F1: (v2 ^ (v2 | (v2 >> ((v2) & 63)))), F2: (v85 / ((int32(v2)) | 1))}
                        _ = v85
                        _ = v86
                        _ = v87
                    }
                    if true {
                        v83 -= int16(v5)
                        var v88 uint64 = 1
                        for fi89 := 0; fi89 < 14; fi89++ {
                            v4 = append(v4, "")
                            v90 := GenS6{F0: (!(!false)), F1: 1, F2: (v88 * v88)}
                            v3 += v3
                            _ = fi89
                            _ = v90
                        }
                        _ = v88
                    }
                    _ = fi84
                }
                v3 += 1
                var v91 int32 = -2147483648
                _ = v91
            } else {
                v92 := GenS5{F0: v5}
                _ = len(v4)
                _ = v92
            }
            _ = v83
        }
        v93 := GenS9{F0: 42, F1: 620, F2: 9223372036854775807}
        _ = v93
    } else {
        var v94 bool = false
        var v95 bool = false
        var v96 int8 = -1
        _ = v94
        _ = v95
        _ = v96
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
