package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 bool; F1 int16; F2 int32 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int8; F1 int }
type GenS10 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 int8 = 100
    var v2 float64 = 0.1
    v3 := make(map[uint32]uint32)
    v4 := []string{"hello", ""}
    switch 127 {
    case 195:
        var v5 uint16 = 65535
        _ = v5
    case 128:
        var v6 string = "\n\t"
        _ = v6
    default:
        if (!((false && (598 >= 4294967295)) && false)) {
            for fi7 := 0; fi7 < 1; fi7++ {
                for fi8 := 0; fi8 < 8; fi8++ {
                    v9 := GenS1{F0: 215, F1: "\n\t"}
                    switch fi7 {
                    case -442:
                        var v10 uint32 = 4294967295
                        _ = v10
                    case -1:
                        fmt.Println(v1, fi7, v1)
                    }
                    _ = fi8
                    _ = v9
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if (-9223372036854775808 < 9223372036854775807) {
                    switch ("hello" + "\n\t") {
                    case "\n\t":
                        v1 += (v1 * 0)
                    }
                }
                _ = fi7
            }
            var v11 bool = true
            var v12 float32 = float32(v2)
            _ = v11
            _ = v12
        }
    }
    switch "a" {
    case "a":
        v4 = append(v4, ("foo bar" + "foo bar"))
    case "foo bar":
        fmt.Println(v2, v2)
    case "hello":
        for fi13 := 0; fi13 < 10; fi13++ {
            if (1 > 1) {
                if (!(!((!((!(true && true)) || false)) || ("" < "")))) {
                    for fi14 := 0; fi14 < 14; fi14++ {
                        if false {
                            var v15 uint8 = 1
                            _ = v15
                        }
                        if false {
                            if true {
                                v4 = append(v4, "a")
                                _ = len(v4)
                            } else {
                                v1 = (^(v1))
                            }
                            var v16 int = fi13
                            var v17 uint16 = 0
                            if (612 < uint32(v17)) {
                                if false {
                                    v18 := GenS2{F0: 472, F1: ("" + "a")}
                                    _ = v18
                                } else {
                                    var v19 int64 = 1
                                    _ = v19
                                }
                                v20 := GenS2{F0: -639, F1: (("hello" + (("\n\t" + (("foo bar" + "foo bar") + ("foo bar" + "foo bar"))) + ("hello" + ("a" + "hello")))) + "\n\t")}
                                _ = v20
                            } else {
                                v16 += fi14
                            }
                            v17 += v17
                            _ = v16
                            _ = v17
                        }
                        _ = fi14
                    }
                }
                v21 := GenS2{F0: -32768, F1: "hello"}
                v1 -= v1
                _ = v21
            }
            fmt.Println(fi13)
            _ = fi13
        }
    }
    v22 := GenS3{F0: (4294967295 < -2), F1: 1, F2: -1}
    v4 = append(v4, "\n\t")
    if (!true) {
        for fi23 := 0; fi23 < 5; fi23++ {
            var v24 int64 = 276
            _ = fi23
            _ = v24
        }
        switch "\n\t" {
        case "":
            switch 621 {
            case 0:
                v25 := GenS1{F0: 218, F1: ("" + "foo bar")}
                _ = v25
            default:
                v26 := GenS4{F0: 0, F1: 245}
                _ = v26
            }
        }
        v2 += ((1.5 + ((v2 / 1.0) + (0.1 * ((1e10 + v2) * v2)))) + (0.1 + ((float64(0.0) / 1.0) / ((0.0 / (3.14159 - ((v2 * 3.14159) / (1e10 - 0.1)))) - ((-2.25 + -2.25) * v2)))))
    }
    if true {
        if false {
            for fi27 := 0; fi27 < 4; fi27++ {
                v28 := GenS4{F0: 9223372036854775807, F1: 99}
                _ = fi27
                _ = v28
            }
            if true {
                var v29 uint64 = 439
                _ = v29
            } else {
                fmt.Println(v1, v1)
                var v30 int16 = -32768
                switch "\n\t" {
                case "":
                    v31 := GenS5{F0: ("hello" + "hello"), F1: ("a" + ("foo bar" + ""))}
                    _ = v31
                case "hello":
                    v32 := GenS4{F0: -9223372036854775808, F1: 103}
                    _ = v32
                default:
                    v30 += (^(v30))
                }
                _ = v30
            }
            if (!true) {
                for fi33 := 0; fi33 < 14; fi33++ {
                    v34 := GenS6{F0: 2147483647}
                    _ = fi33
                    _ = v34
                }
            } else {
                for fi35 := 0; fi35 < 5; fi35++ {
                    v36 := v3[2147483648]
                    _ = fi35
                    _ = v36
                }
                for fi37 := 0; fi37 < 18; fi37++ {
                    v4 = append(v4, "foo bar")
                    if (true || false) {
                        v1 -= (v1 - v1)
                        for fi38 := 0; fi38 < 1; fi38++ {
                            var v39 uint64 = 42
                            _ = fi38
                            _ = v39
                        }
                    } else {
                        switch (fi37 << ((549) & 63)) {
                        case 42:
                            switch "\n\t" {
                            case "foo bar":
                                fmt.Println(v1, v2, v1)
                            default:
                                var v40 int16 = 895
                                _ = v40
                            }
                        case 975:
                            for fi41 := 0; fi41 < 13; fi41++ {
                                switch "a" {
                                case "a":
                                    v2 -= v2
                                    fmt.Println(fi37, fi41, fi37)
                                    switch (^(fi37)) {
                                    case -271:
                                        if (!true) {
                                            _ = len(v4)
                                        }
                                    }
                                case "foo bar":
                                    v4 = append(v4, ("hello" + ""))
                                }
                                if false {
                                    var v42 bool = true
                                    v43 := GenS7{F0: (fi41 * fi37)}
                                    _ = v42
                                    _ = v43
                                }
                                _ = fi41
                            }
                        case 167:
                            v2 -= v2
                        }
                        for fi44 := 0; fi44 < 18; fi44++ {
                            v45 := GenS8{F0: -2147483648}
                            var v46 int16 = 42
                            v4 = append(v4, "hello")
                            _ = fi44
                            _ = v45
                            _ = v46
                        }
                    }
                    _ = fi37
                }
                v47 := GenS1{F0: 1, F1: "\n\t"}
                _ = v47
            }
        } else {
            v48 := GenS4{F0: -9223372036854775808, F1: 127}
            v49 := GenS2{F0: 442, F1: "hello"}
            _ = v48
            _ = v49
        }
        v50 := GenS9{F0: (^(v1)), F1: -1}
        _ = v50
    } else {
        for fi51 := 0; fi51 < 8; fi51++ {
            var v52 uint64 = 329
            v1 += v1
            switch -750 {
            case 127:
                _ = len(v3)
            default:
                var v53 bool = (0 <= 189)
                _ = v53
            }
            _ = fi51
            _ = v52
        }
        var v54 uint = 939
        v1 = 127
        _ = v54
    }
    var v55 uint = 4294967295
    if true {
        var v56 uint64 = 518
        _ = v56
    } else {
        if (-9223372036854775808 != -1) {
            v55 += (v55 ^ v55)
        } else {
            v57 := GenS10{F0: 127, F1: 925}
            for fi58 := 0; fi58 < 19; fi58++ {
                v59 := v3[531]
                fmt.Println(v2, v59)
                _ = fi58
                _ = v59
            }
            v60 := v3[2]
            _ = v57
            _ = v60
        }
        v1 = v1
        if (true && false) {
            switch 32767 {
            case 4294967295:
                var v61 uint8 = 255
                _ = v61
            }
        }
    }
    v1 += (v1 - v1)
    v2 += v2
    v2 -= 1e10
    if (2147483647 <= -1) {
        v1 -= 25
    } else {
        if (-1 == 609) {
            switch 128 {
            case 4294967295:
                v62 := GenS8{F0: -224}
                _ = v62
            }
        }
        switch 1 {
        case -260:
            fmt.Println(v2, v55)
        }
    }
    for fi63 := 0; fi63 < 17; fi63++ {
        fmt.Println(v1)
        if false {
            if (2147483647 < 167) {
                v2 -= (v2 - 3.14159)
                v2 -= -2.25
            }
            if (!false) {
                if true {
                    for fi64 := 0; fi64 < 1; fi64++ {
                        var v65 int8 = v1
                        switch 9223372036854775807 {
                        case 127:
                            var v66 float32 = (0.1 / (1e10 - 3.14159))
                            fmt.Println(v2, fi64, fi63)
                            var v67 uint8 = uint8(v55)
                            _ = v66
                            _ = v67
                        case 42:
                            var v68 int16 = 1000
                            _ = v68
                        }
                        v69 := GenS8{F0: 960}
                        var v70 string = "hello"
                        var v71 uint32 = 137
                        _ = fi64
                        _ = v65
                        _ = v69
                        _ = v70
                        _ = v71
                    }
                    v2 = v2
                    v55 = 859
                } else {
                    _ = len(v4)
                }
                var v72 int32 = 2147483647
                v1 = -31
                _ = v72
            } else {
                switch fi63 {
                case -9223372036854775808:
                    fmt.Println(v2)
                }
                if len(v4) > 0 {
                    v4[0] = "foo bar"
                }
            }
        } else {
            if false {
                fmt.Println(fi63, v2)
                v73 := v3[2147483648]
                _ = v73
            } else {
                v74 := v3[2]
                _ = v74
            }
            var v75 string = ("a" + "a")
            _ = v75
        }
        _ = fi63
    }
    var v76 uint32 = 562
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v55:", v55)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v55
    _ = v76
}
