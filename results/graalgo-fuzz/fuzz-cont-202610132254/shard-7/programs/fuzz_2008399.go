package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint64; F2 int32 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 string }
type GenS9 struct { F0 string; F1 uint }
type GenS10 struct { F0 uint64 }

func main() {
    v1 := []string{"foo bar", "hello"}
    var v2 string = "a"
    var v3 float64 = 0.0
    var v4 int8 = 42
    v5 := GenS1{F0: ((v2 + v2) + ((("foo bar" + (((v2 + "a") + v2) + "hello")) + v2) + "hello")), F1: 2, F2: 727}
    if (!true) {
        if false {
            var v6 int64 = 828
            var v7 uint32 = 267
            var v8 float64 = float64(0.0)
            _ = v6
            _ = v7
            _ = v8
        }
        _ = len(v1)
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 uint64 = 18446744073709551615
            var v11 float64 = -2.25
            if false {
                switch (("" + v2) + ("foo bar" + "")) {
                case "a":
                    v10 = v10
                case "foo bar":
                    v1 = append(v1, v2)
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                if false {
                    var v12 string = (((v2 + (v2 + v2)) + v2) + "")
                    _ = v12
                } else {
                    var v13 uint64 = v10
                    for fi14 := 0; fi14 < 14; fi14++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if (!(!(348 > -344))) {
                            v11 += v11
                            fmt.Println(v3, v3, v13)
                        } else {
                            switch "hello" {
                            case "\n\t":
                                if false {
                                    var v15 uint32 = 1
                                    var v16 uint8 = 121
                                    _ = v15
                                    _ = v16
                                } else {
                                    var v17 int = 42
                                    _ = v17
                                }
                            case "foo bar":
                                v18 := GenS1{F0: "foo bar", F1: 18, F2: 520}
                                _ = v18
                            }
                        }
                        _ = fi14
                    }
                    _ = v13
                }
            } else {
                switch v2 {
                case "hello":
                    v19 := GenS2{F0: -833, F1: (^(v10)), F2: -34}
                    _ = v19
                case "\n\t":
                    fmt.Println(v2)
                default:
                    var v20 int16 = -492
                    _ = v20
                }
                if (!((98 >= 196) && true)) {
                    switch ("foo bar" + "foo bar") {
                    case "foo bar":
                        v21 := GenS1{F0: v2, F1: 128, F2: 2147483648}
                        _ = v21
                    default:
                        _ = len(v1)
                    }
                }
            }
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    var v22 uint32 = 42
    v4 = (v4 >> ((18446744073709551615) & 63))
    for fi23 := 0; fi23 < 19; fi23++ {
        v1 = append(v1, ("hello" + v2))
        _ = fi23
    }
    var v24 uint8 = 2
    if (false && true) {
        v2 = v2
        for fi25 := 0; fi25 < 5; fi25++ {
            fmt.Println(v24)
            var v26 int = (fi25 + (-(fi25)))
            _ = fi25
            _ = v26
        }
    }
    v3 = v3
    switch 9223372036854775807 {
    case 974:
        fmt.Println(v3, v24, v24)
    case 4294967295:
        if (702 <= 42) {
            switch ("foo bar" + "foo bar") {
            case "a":
                switch ("" + v2) {
                case "a":
                    if false {
                        var v27 uint64 = 4294967295
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 1; fi28++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi28
                        }
                    }
                case "\n\t":
                    for fi29 := 0; fi29 < 12; fi29++ {
                        if (((false && (!(!true))) && false) && (true && (true || true))) {
                            var v30 uint64 = 18446744073709551615
                            v22 = (v22 * (v22 | (v22 - v22)))
                            _ = v30
                        }
                        if true {
                            v3 = float64(1.5)
                            if (!false) {
                                v31 := GenS3{F0: 1, F1: fi29}
                                var v32 string = v2
                                if true {
                                    v33 := GenS4{F0: 0}
                                    _ = v33
                                } else {
                                    var v34 int32 = 42
                                    switch ("a" + v2) {
                                    case "foo bar":
                                        var v35 uint = 18446744073709551615
                                        _ = v35
                                    case "hello":
                                        v36 := GenS1{F0: (v32 + v32), F1: (v24 + uint8(v22)), F2: 308}
                                        _ = v36
                                    default:
                                        if (((!true) && ((v24 | v24) <= v24)) || (0 <= 1)) {
                                            var v37 uint = 211
                                            _ = v37
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v38 := GenS5{F0: 9223372036854775807, F1: (v4 - (-(v4)))}
                                            var v39 int32 = (v34 - (v34 >> ((1) & 63)))
                                            _ = v38
                                            _ = v39
                                        }
                                    }
                                    v1 = append(v1, v32)
                                    _ = v34
                                }
                                _ = v31
                                _ = v32
                            } else {
                                for fi40 := 0; fi40 < 9; fi40++ {
                                    v41 := GenS6{F0: 469, F1: v2}
                                    if false {
                                        v2 = v2
                                        var v42 int = fi29
                                        _ = v42
                                    } else {
                                        switch (fi29 & (^(fi29))) {
                                        case 127:
                                            v22 = (v22 ^ v22)
                                        case -518:
                                            if false {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v2 = "a"
                                                v1 = append(v1, "")
                                            } else {
                                                var v43 int16 = -597
                                                if ((v43 & v43) != 981) {
                                                    if len(v1) > 0 {
                                                        v1[0] = v2
                                                    }
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    var v44 int = (fi29 * int(v4))
                                                    _ = v44
                                                } else {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    v45 := GenS4{F0: -1}
                                                    _ = v45
                                                }
                                                _ = v43
                                            }
                                        case -252:
                                            _ = len(v1)
                                        default:
                                            switch v2 {
                                            case "\n\t":
                                                var v46 uint8 = 219
                                                _ = v46
                                            case "hello":
                                                var v47 int16 = 1
                                                _ = v47
                                            default:
                                                v3 -= v3
                                            }
                                        }
                                    }
                                    var v48 string = v2
                                    switch (fi29 >> ((0) & 63)) {
                                    case -1:
                                        var v49 uint8 = v24
                                        _ = v49
                                    case 9223372036854775807:
                                        var v50 int32 = 0
                                        _ = v50
                                    default:
                                        v51 := GenS3{F0: -32768, F1: (fi29 ^ 106)}
                                        _ = v51
                                    }
                                    _ = fi40
                                    _ = v41
                                    _ = v48
                                }
                                if (918 <= 42) {
                                    fmt.Println(v22, v24)
                                    switch "\n\t" {
                                    case "\n\t":
                                        fmt.Println(fi29, v3, v4)
                                    default:
                                        var v52 int16 = -970
                                        _ = v52
                                    }
                                    if true {
                                        v53 := GenS7{F0: v2, F1: (1 != 863)}
                                        _ = v53
                                    }
                                }
                            }
                        }
                        _ = fi29
                    }
                default:
                    v22 = 2147483647
                }
            case "hello":
                switch 461 {
                case 4294967295:
                    v54 := GenS1{F0: (v2 + ""), F1: uint8(v22), F2: v22}
                    _ = v54
                case -742:
                    for fi55 := 0; fi55 < 2; fi55++ {
                        switch fi55 {
                        case 934:
                            switch -1 {
                            case 531:
                                fmt.Println(fi55, v4)
                            case -1:
                                fmt.Println(v24, fi55, v24)
                            case -774:
                                for fi56 := 0; fi56 < 12; fi56++ {
                                    switch fi56 {
                                    case -2147483648:
                                        v24 = (v24 + v24)
                                    case 2147483647:
                                        switch ("foo bar" + "\n\t") {
                                        case "foo bar":
                                            fmt.Println(fi55, v2, v4)
                                        case "hello":
                                            var v57 int8 = (v4 * (v4 & v4))
                                            _ = v57
                                        case "":
                                            fmt.Println(v3, v24, v2)
                                        default:
                                            fmt.Println(v2)
                                        }
                                    case -949:
                                        var v58 uint64 = 500
                                        _ = v58
                                    default:
                                        switch (((v2 + v2) + v2) + "foo bar") {
                                        case "foo bar":
                                            if false {
                                                var v59 uint64 = 42
                                                _ = v59
                                            }
                                        case "a":
                                            if false {
                                                if (((v2 + (("" + v2) + (v2 + (v2 + v2)))) + (v2 + "hello")) < ("foo bar" + v2)) {
                                                    v22 += v22
                                                    var v60 uint32 = (v22 & (v22 + v22))
                                                    _ = v60
                                                }
                                                v22 += 487
                                                var v61 float64 = (v3 - v3)
                                                if (9223372036854775807 == 579) {
                                                    var v62 uint8 = (v24 & (v24 % (((v24 | (v24 >> ((1) & 63)))) | 1)))
                                                    v3 += 3.14159
                                                    _ = v62
                                                } else {
                                                    v2 = v2
                                                    fmt.Println(v2, fi55)
                                                    v4 = -23
                                                }
                                                _ = v61
                                            }
                                        case "hello":
                                            fmt.Println(v4)
                                        default:
                                            var v63 bool = (!(true && ((^(v22)) < v22)))
                                            _ = v63
                                        }
                                    }
                                    v24 = v24
                                    _ = fi56
                                }
                            default:
                                v64 := GenS8{F0: ("" + (("foo bar" + "hello") + ((((v2 + "foo bar") + "") + ((v2 + (v2 + v2)) + (((v2 + v2) + v2) + "a"))) + v2)))}
                                _ = v64
                            }
                        case -9223372036854775808:
                            for fi65 := 0; fi65 < 20; fi65++ {
                                var v66 bool = (-1 <= 0)
                                _ = fi65
                                _ = v66
                            }
                        case 128:
                            var v67 uint = 944
                            _ = v67
                        }
                        if len(v1) > 0 {
                            v1[0] = "hello"
                        }
                        _ = fi55
                    }
                }
            default:
                v68 := GenS3{F0: -1, F1: -645}
                _ = v68
            }
            for fi69 := 0; fi69 < 5; fi69++ {
                if true {
                    var v70 int = 32767
                    _ = v70
                }
                _ = fi69
            }
            if false {
                v3 = 1.5
                v71 := GenS9{F0: v2, F1: 40}
                _ = v71
            }
        }
    case -489:
        switch -9223372036854775808 {
        case 971:
            switch v2 {
            case "foo bar":
                switch 1 {
                case 42:
                    v22 = v22
                default:
                    v72 := GenS10{F0: 4294967295}
                    _ = v72
                }
            case "\n\t":
                v4 += v4
            case "hello":
                v4 = v4
            }
        }
    }
    v73 := GenS9{F0: (v2 + (((v2 + v2) + v2) + (v2 + v2))), F1: 18446744073709551615}
    v4 -= v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v73
}
