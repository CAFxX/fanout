package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 bool }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint8; F1 int64; F2 uint8 }
func genFn0(p0 uint32) int16 {
    v6 := GenS1{F0: 1, F1: 127, F2: (false && true)}
    v7 := []int{-2147483648, 0, 2147483647}
    if len(v7) > 0 {
        v7[0] = 32767
    }
    _ = p0
    _ = v6
    _ = v7
    return -943
}
func genFn1() int16 {
    var v8 uint8 = 2
    var v9 int8 = 42
    _ = v8
    _ = v9
    return 32767
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]uint32)
    var v4 uint16 = 282
    var v5 bool = (0.0 < 1e10)
    if (-1 > 1000) {
        v10 := v2[2147483648]
        var v11 uint32 = 4294967295
        var v12 bool = v10
        _ = v10
        _ = v11
        _ = v12
    } else {
        for fi13 := 0; fi13 < 11; fi13++ {
            if v5 {
                for fi14 := 0; fi14 < 4; fi14++ {
                    fmt.Println(fi14, fi14, fi13)
                    if (69 <= 0) {
                        switch "foo bar" {
                        case "hello":
                            v15 := GenS1{F0: v4, F1: (-(fi14)), F2: ("hello" >= "a")}
                            _ = v15
                        case "":
                            var v16 float64 = 0.0
                            _ = v16
                        }
                        v4 = v4
                    } else {
                        var v17 int = fi14
                        if (2 < 0) {
                            switch "foo bar" {
                            case "":
                                var v18 uint8 = 181
                                _ = v18
                            default:
                                switch v17 {
                                case 2147483647:
                                    switch (^(fi13)) {
                                    case 32768:
                                        v5 = (4294967295 != 65535)
                                    default:
                                        if ((("hello" <= "\n\t") || v5) && false) {
                                            var v19 int16 = -206
                                            _ = v19
                                        }
                                    }
                                case -1:
                                    v4 = v4
                                default:
                                    v4 -= (v4 << ((uint(v4)) & 63))
                                    fmt.Println(v5, v5, v4)
                                    var v20 int8 = 42
                                    v5 = (0 < 200)
                                    _ = v20
                                }
                            }
                        } else {
                            v4 = (v4 % ((1) | 1))
                        }
                        var v21 uint16 = (v4 << ((255) & 63))
                        _ = v17
                        _ = v21
                    }
                    v22 := GenS1{F0: (v4 >> ((1) & 63)), F1: (fi13 ^ fi13), F2: true}
                    _ = fi14
                    _ = v22
                }
            } else {
                v5 = v5
                var v23 float32 = (1.5 * 1.5)
                if (fi13 != fi13) {
                    v4 = v4
                    switch ("foo bar" + "hello") {
                    case "\n\t":
                        v4 = v4
                    default:
                        for fi24 := 0; fi24 < 13; fi24++ {
                            fmt.Println(v23, fi24)
                            _ = fi24
                        }
                    }
                    if (!(!v5)) {
                        if (586 <= 849) {
                            var v25 float64 = (0.0 * ((((-2.25 - 1.5) + 1.5) / (1e10 * 1.5)) + 1.5))
                            v23 -= v23
                            _ = v25
                        }
                        for fi26 := 0; fi26 < 16; fi26++ {
                            fmt.Println(fi26, fi13)
                            var v27 int64 = 620
                            fmt.Println(v23, v27)
                            _ = fi26
                            _ = v27
                        }
                    }
                }
                _ = v23
            }
            for fi28 := 0; fi28 < 17; fi28++ {
                v5 = true
                var v29 int32 = 42
                var v30 uint64 = 214
                v31 := []int{(fi28 & fi28), (fi13 - (fi28 % ((fi13) | 1)))}
                var v32 uint = 598
                switch "foo bar" {
                case "\n\t":
                    var v33 uint64 = uint64(v32)
                    _ = v33
                case "hello":
                    fmt.Println(v5, v29, fi13)
                    v29 += 1
                    var v34 int8 = 0
                    _ = v34
                case "a":
                    if ((v29 * (v29 ^ 2147483647)) <= v29) {
                        var v35 uint64 = (v30 / ((4294967295) | 1))
                        var v36 float64 = 0.1
                        _ = v35
                        _ = v36
                    } else {
                        if false {
                            v29 += (v29 / ((v29) | 1))
                        }
                        v29 += -2147483648
                        v37 := v3[("" + "foo bar")]
                        _ = v37
                    }
                }
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
                _ = v32
            }
            _ = fi13
        }
        if false {
            delete(v2, 4294967295)
            v5 = false
        } else {
            fmt.Println(v4)
            if (905 >= 2) {
                var v38 uint8 = 2
                v4 += (v4 | v4)
                v39 := []int{185, 42, 0, -2147483648}
                _ = v38
                _ = v39
            }
        }
    }
    if false {
        switch "foo bar" {
        case "\n\t":
            if v5 {
                switch ("" + "\n\t") {
                case "foo bar":
                    if true {
                        fmt.Println(v4, v4)
                    }
                }
            }
        case "a":
            if ((v4 >> ((1) & 63)) != (^(v4))) {
                var v40 string = ((((("" + (("foo bar" + "a") + "a")) + "") + "a") + "hello") + ((("a" + ("" + "hello")) + "") + ""))
                v5 = v5
                if false {
                    _ = len(v1)
                    switch (("\n\t" + "foo bar") + "") {
                    case "hello":
                        v2[65535] = ((!v5) && (!v5))
                    default:
                        var v41 uint8 = 0
                        _ = v41
                    }
                }
                _ = v40
            }
        case "foo bar":
            v5 = v5
        }
    } else {
        for fi42 := 0; fi42 < 10; fi42++ {
            var v43 int8 = 42
            delete(v2, 2)
            _ = fi42
            _ = v43
        }
        for fi44 := 0; fi44 < 2; fi44++ {
            var v45 float64 = 1.5
            _ = fi44
            _ = v45
        }
    }
    var v46 int64 = -9223372036854775808
    switch 2147483647 {
    case 2147483647:
        v47 := []int{1}
        _ = v47
    case 32768:
        var v48 int64 = v46
        _ = v48
    }
    for fi49 := 0; fi49 < 15; fi49++ {
        v1["foo bar"] = ("foo bar" + "\n\t")
        _ = fi49
    }
    v46 -= 1
    v4 = (v4 / ((998) | 1))
    v46 = (v46 + 878)
    if (104 > 255) {
        var v50 uint16 = (^(v4))
        for fi51 := 0; fi51 < 10; fi51++ {
            for fi52 := 0; fi52 < 4; fi52++ {
                v5 = (!false)
                var v53 string = "\n\t"
                var v54 uint8 = 1
                _ = fi52
                _ = v53
                _ = v54
            }
            switch ("\n\t" + ("foo bar" + "\n\t")) {
            case "hello":
                for fi55 := 0; fi55 < 2; fi55++ {
                    switch 9223372036854775807 {
                    case 127:
                        switch "\n\t" {
                        case "":
                            fmt.Println(fi55)
                        case "\n\t":
                            var v56 bool = ("" < "hello")
                            _ = v56
                        case "a":
                            if (97 == 42) {
                                var v57 uint = 65535
                                v46 += -609
                                _ = v57
                            }
                        }
                    case 9223372036854775807:
                        v58 := []int{(fi55 / (((-(fi55))) | 1)), fi51, fi55}
                        if (!v5) {
                            v58 = append(v58, 128)
                        } else {
                            v4 = v50
                            if (!(417 != 18446744073709551615)) {
                                switch (fi55 >> ((0) & 63)) {
                                case -1:
                                    if ((200 < 18446744073709551615) || true) {
                                        v59 := GenS2{F0: (float64(-2.25) <= 0.1), F1: 1}
                                        if (966 < -780) {
                                            if v5 {
                                                var v60 float64 = 1e10
                                                _ = v60
                                            } else {
                                                fmt.Println(fi55)
                                            }
                                        }
                                        v50 = v50
                                        _ = v59
                                    } else {
                                        var v61 bool = (!(2 >= 2))
                                        if len(v58) >= 2 {
                                            v58 = v58[1:len(v58)]
                                        }
                                        switch fi55 {
                                        case 128:
                                            var v62 bool = (65535 == 328)
                                            _ = v62
                                        case 802:
                                            var v63 bool = false
                                            _ = v63
                                        case 4294967295:
                                            var v64 bool = (false && v61)
                                            _ = v64
                                        }
                                        v65 := GenS3{F0: 255, F1: v46, F2: 255}
                                        _ = v61
                                        _ = v65
                                    }
                                default:
                                    v66 := GenS2{F0: (float64(float32(((1e10 * float64(0.0)) + float64((((1e10 * 0.1) + 1.5) - 0.0))))) == 1.5), F1: -864}
                                    _ = v66
                                }
                                _ = len(v2)
                            }
                            fmt.Println(fi55, v4, fi55)
                        }
                        if (!false) {
                            if v5 {
                                var v67 int8 = 0
                                _ = v67
                            }
                            _ = len(v2)
                        }
                        _ = v58
                    case 128:
                        v4 -= 32767
                    }
                    _ = fi55
                }
            case "\n\t":
                var v68 int16 = 42
                _ = v68
            case "":
                var v69 int8 = -1
                _ = v69
            default:
                v70 := v3[""]
                _ = v70
            }
            v71 := GenS3{F0: 1, F1: v46, F2: 160}
            _ = fi51
            _ = v71
        }
        _ = v50
    }
    for fi72 := 0; fi72 < 16; fi72++ {
        v4 += (^(v4))
        _ = fi72
    }
    v46 = v46
    if v5 {
        for fi73 := 0; fi73 < 17; fi73++ {
            v4 = (v4 / ((386) | 1))
            _ = fi73
        }
    }
    if (v5 || (v5 || v5)) {
        switch -2147483648 {
        case 4294967295:
            var v74 string = ("foo bar" + (("a" + "\n\t") + "foo bar"))
            _ = v74
        case 720:
            if v5 {
                var v75 uint = 255
                _ = v75
            } else {
                v76 := []int{0, 127, -1}
                if (0 != (v4 % ((v4) | 1))) {
                    if (1 != 843) {
                        v2[802] = ((v5 && (!v5)) || (4294967295 == 42))
                        var v77 int16 = -114
                        fmt.Println(v4, v4, v46)
                        _ = v77
                    }
                }
                var v78 uint16 = v4
                _ = v76
                _ = v78
            }
        default:
            for fi79 := 0; fi79 < 3; fi79++ {
                fmt.Println(fi79)
                switch "foo bar" {
                case "\n\t":
                    var v80 float64 = -2.25
                    _ = v80
                case "a":
                    var v81 uint64 = 467
                    _ = v81
                default:
                    var v82 int32 = -455
                    _ = v82
                }
                _ = fi79
            }
        }
    }
    for fi83 := 0; fi83 < 17; fi83++ {
        v84 := []int{fi83, 127}
        v85 := GenS1{F0: v4, F1: fi83, F2: v5}
        for fi86 := 0; fi86 < 10; fi86++ {
            if ((!("hello" == ("\n\t" + (("a" + ("foo bar" + "a")) + ("a" + ("" + (("foo bar" + ("foo bar" + "foo bar")) + ("\n\t" + ("\n\t" + "foo bar"))))))))) && (v5 || (-128 != 0))) {
                if false {
                    v87 := GenS3{F0: 0, F1: v46, F2: 2}
                    _ = v87
                }
                v4 = v4
            }
            v5 = v5
            var v88 int32 = 42
            fmt.Println(v88)
            _ = fi86
            _ = v88
        }
        _ = fi83
        _ = v84
        _ = v85
    }
    for fi89 := 0; fi89 < 5; fi89++ {
        if ((((0 <= 0) && v5) || (1 > 132)) && v5) {
            var v90 float64 = float64(1.5)
            switch fi89 {
            case 55:
                v4 -= 65535
            default:
                _ = len(v2)
            }
            _ = v90
        } else {
            for fi91 := 0; fi91 < 20; fi91++ {
                var v92 uint16 = 0
                _ = fi91
                _ = v92
            }
        }
        v93 := []int{(fi89 & (fi89 | fi89)), (fi89 << ((320) & 63)), fi89}
        fmt.Println(fi89)
        _ = fi89
        _ = v93
    }
    if (((v46 - v46) <= -9223372036854775808) || (v4 == 255)) {
        var v94 int16 = -32768
        _ = v94
    }
    fmt.Println(v46)
    var v95 uint32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v95
}
