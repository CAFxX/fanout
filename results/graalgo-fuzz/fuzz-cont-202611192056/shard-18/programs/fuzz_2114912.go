package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int64; F1 int8; F2 uint8 }
type GenS3 struct { F0 int8; F1 uint32; F2 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 int32 }
type GenS6 struct { F0 int; F1 int; F2 uint }
type GenS7 struct { F0 int8; F1 int; F2 int8 }

func main() {
    v1 := make(map[string]bool)
    v2 := []string{"a"}
    var v3 float64 = -2.25
    var v4 int64 = 1
    v5 := []uint32{2, 65535, 2147483647}
    if true {
        v6 := GenS1{F0: 32767, F1: 1}
        v7 := GenS1{F0: 583, F1: 242}
        v3 = 0.0
        _ = v6
        _ = v7
    } else {
        if ((((0.0 >= float64(-2.25)) && false) && true) || (!(v3 < 0.0))) {
            var v8 int8 = 127
            delete(v1, ("" + "foo bar"))
            v8 += (v8 + v8)
            _ = v8
        }
        switch "hello" {
        case "a":
            v9 := GenS1{F0: 0, F1: 121}
            _ = v9
        case "":
            v10 := GenS1{F0: 255, F1: 128}
            _ = v10
        case "\n\t":
            if len(v2) > 0 {
                v2[0] = "hello"
            }
        }
    }
    fmt.Println(v3, v3)
    var v11 float64 = v3
    v11 = (0.1 / ((-2.25 - ((((float64(1e10) + float64((float32(v11) + float32(0.0)))) * v3) / 3.14159) * v3)) + (v3 - v3)))
    var v12 uint = 69
    var v13 uint8 = 42
    fmt.Println(v11)
    for fi14 := 0; fi14 < 14; fi14++ {
        var v15 string = "a"
        v3 += float64(1.5)
        if false {
            var v16 int = -568
            _ = v16
        }
        _ = fi14
        _ = v15
    }
    if (0 != -711) {
        v3 -= v11
        var v17 int = 506
        for fi18 := 0; fi18 < 7; fi18++ {
            for fi19 := 0; fi19 < 13; fi19++ {
                v11 = (float64((float32(v3) - -2.25)) + -2.25)
                fmt.Println(v17, v4, v11)
                _ = fi19
            }
            switch "\n\t" {
            case "\n\t":
                v3 += v3
            case "foo bar":
                v1["foo bar"] = (!false)
            case "a":
                for fi20 := 0; fi20 < 9; fi20++ {
                    switch fi20 {
                    case -429:
                        fmt.Println(v17)
                    case 1:
                        delete(v1, "hello")
                        var v21 int8 = 1
                        fmt.Println(fi20)
                        _ = v21
                    case -263:
                        v22 := GenS1{F0: 42, F1: (v13 >> ((v12) & 63))}
                        _ = v22
                    }
                    var v23 uint16 = 580
                    _ = fi20
                    _ = v23
                }
            }
            var v24 bool = true
            _ = fi18
            _ = v24
        }
        _ = v17
    } else {
        v25 := v1["foo bar"]
        _ = v25
    }
    if false {
        v26 := GenS2{F0: 9223372036854775807, F1: 0, F2: 160}
        var v27 uint8 = (v13 / (((v13 & (v13 & v13))) | 1))
        _ = v26
        _ = v27
    }
    if false {
        v13 = v13
        v12 = (v12 + v12)
    } else {
        v13 -= 128
        if (false || true) {
            v12 = (^(v12))
            v1[""] = true
        } else {
            if true {
                if true {
                    if (false || (4294967295 == 4294967295)) {
                        v28 := GenS2{F0: -861, F1: int8(v13), F2: v13}
                        if false {
                            delete(v1, "")
                            v29 := GenS3{F0: 42, F1: 0, F2: -128}
                            v4 -= v4
                            _ = v29
                        }
                        v4 = -1
                        _ = v28
                    } else {
                        v30 := GenS4{F0: 9223372036854775807}
                        if (!false) {
                            var v31 int64 = (v4 % ((-82) | 1))
                            _ = v31
                        }
                        for fi32 := 0; fi32 < 17; fi32++ {
                            for fi33 := 0; fi33 < 1; fi33++ {
                                var v34 int64 = -1
                                v3 -= float64((0.0 / float32(v3)))
                                var v35 int = fi32
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                            fmt.Println(v12, v3, v13)
                            v36 := GenS1{F0: 381, F1: (v13 / ((v13) | 1))}
                            _ = fi32
                            _ = v36
                        }
                        _ = v30
                    }
                } else {
                    v3 = v11
                    for fi37 := 0; fi37 < 10; fi37++ {
                        if ((v13 << (((v12 | v12)) & 63)) > v13) {
                            v38 := GenS4{F0: 128}
                            for fi39 := 0; fi39 < 19; fi39++ {
                                _ = len(v2)
                                _ = fi39
                            }
                            v4 -= v4
                            _ = v38
                        } else {
                            var v40 int8 = 0
                            v3 = v11
                            v13 -= (v13 << ((v12) & 63))
                            _ = v40
                        }
                        for fi41 := 0; fi41 < 3; fi41++ {
                            var v42 uint8 = 128
                            var v43 uint = (v12 * (v12 * v12))
                            _ = fi41
                            _ = v42
                            _ = v43
                        }
                        for fi44 := 0; fi44 < 19; fi44++ {
                            if (v4 == int64(fi44)) {
                                v45 := v1["foo bar"]
                                if (v11 >= (v11 / 1.0)) {
                                    v45 = true
                                    v5 = append(v5, 70)
                                } else {
                                    v11 = (-2.25 / (float64(1.5) * float64((float32(v11) / ((0.0 * (3.14159 / 3.14159)) - (0.1 - ((0.0 + (((3.14159 / 1.0) / 1.5) - 1.5)) + (-2.25 / (0.1 / 3.14159)))))))))
                                    delete(v1, "a")
                                    fmt.Println(v45, v45, v12)
                                    if true {
                                        v3 = -2.25
                                        switch "foo bar" {
                                        case "":
                                            var v46 string = "hello"
                                            _ = v46
                                        case "\n\t":
                                            var v47 float64 = v3
                                            _ = v47
                                        default:
                                            var v48 uint8 = v13
                                            _ = v48
                                        }
                                        if ((!(42 != 42)) && v45) {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        var v49 bool = false
                                        _ = v49
                                    } else {
                                        v3 = v3
                                        if ((v13 * v13) != (v13 ^ (^(v13)))) {
                                            switch (fi44 % ((fi44) | 1)) {
                                            case -2147483648:
                                                v50 := GenS1{F0: 42, F1: (v13 | v13)}
                                                _ = v50
                                            case -9223372036854775808:
                                                v51 := GenS5{F0: v45, F1: 0}
                                                _ = v51
                                            case 0:
                                                var v52 int32 = 469
                                                _ = v52
                                            default:
                                                switch "\n\t" {
                                                case "hello":
                                                    if v45 {
                                                        if v45 {
                                                            v3 = v3
                                                        }
                                                        delete(v1, ("\n\t" + ("" + "foo bar")))
                                                        v5 = append(v5, 1)
                                                    } else {
                                                        fmt.Println(fi37, fi44, v13)
                                                    }
                                                case "\n\t":
                                                    v4 += v4
                                                default:
                                                    if ("a" <= "") {
                                                        v4 += (v4 + v4)
                                                    }
                                                    v12 = v12
                                                    var v53 uint8 = (v13 % (((v13 - v13)) | 1))
                                                    _ = v53
                                                }
                                            }
                                            v54 := GenS5{F0: true, F1: -1}
                                            if v45 {
                                                fmt.Println(v3)
                                                v55 := GenS6{F0: (fi37 % (((fi37 & 332)) | 1)), F1: fi44, F2: 350}
                                                v56 := GenS5{F0: ((v4 & (-(v4))) < v4), F1: 382}
                                                _ = v55
                                                _ = v56
                                            }
                                            delete(v1, "foo bar")
                                            v13 -= (v13 * v13)
                                            if ("a" == "") {
                                                switch "foo bar" {
                                                case "a":
                                                    v57 := GenS3{F0: -1, F1: 0, F2: int8(v4)}
                                                    _ = v57
                                                }
                                                if false {
                                                    var v58 float64 = v11
                                                    _ = len(v2)
                                                    _ = v58
                                                }
                                            }
                                            v45 = (-892 >= 1000)
                                            _ = v54
                                        } else {
                                            var v59 int16 = -1
                                            v60 := GenS2{F0: v4, F1: 127, F2: (v13 >> (((v12 - (v12 + (v12 ^ v12)))) & 63))}
                                            _ = v59
                                            _ = v60
                                        }
                                    }
                                }
                                v4 -= -9223372036854775808
                                _ = v45
                            } else {
                                var v61 uint = v12
                                if (32767 != 193) {
                                    v62 := GenS5{F0: false, F1: -1}
                                    fmt.Println(v61, v13, fi37)
                                    _ = v62
                                } else {
                                    var v63 uint32 = 0
                                    _ = v63
                                }
                                _ = v61
                            }
                            v4 += (v4 & (v4 ^ v4))
                            _ = fi44
                        }
                        _ = fi37
                    }
                    v64 := GenS3{F0: -1, F1: 65, F2: 127}
                    _ = v64
                }
            }
        }
        switch "foo bar" {
        case "foo bar":
            if (true && false) {
                switch 32767 {
                case -2147483648:
                    for fi65 := 0; fi65 < 2; fi65++ {
                        v1[("\n\t" + ((("hello" + "") + "a") + "\n\t"))] = false
                        v1["a"] = ((v12 + v12) > (v12 << ((v12) & 63)))
                        _ = fi65
                    }
                case 127:
                    if ((1 >= 42) && ((float32(v11) * (0.1 / float32(1.5))) == ((1e10 * -2.25) + (-2.25 * 3.14159)))) {
                        for fi66 := 0; fi66 < 1; fi66++ {
                            for fi67 := 0; fi67 < 1; fi67++ {
                                if true {
                                    var v68 uint32 = 2147483648
                                    v3 += (0.0 * (1.5 - 3.14159))
                                    v12 += (v12 - (^(v12)))
                                    _ = v68
                                }
                                if (false || false) {
                                    var v69 uint64 = 298
                                    var v70 int = 128
                                    switch (fi66 | v70) {
                                    case 2147483647:
                                        var v71 uint32 = 2
                                        _ = v71
                                    case -9223372036854775808:
                                        if (0 == 2) {
                                            v4 += v4
                                        } else {
                                            v72 := GenS7{F0: -49, F1: 0, F2: -17}
                                            if (v12 < (v12 | (v12 ^ 523))) {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                delete(v1, ("" + ""))
                                                switch (("foo bar" + ("foo bar" + "hello")) + ("\n\t" + ("" + ("foo bar" + "")))) {
                                                case "a":
                                                    v73 := GenS4{F0: 989}
                                                    _ = v73
                                                }
                                            } else {
                                                v12 = v12
                                                fmt.Println(v3)
                                            }
                                            if (v69 != 18446744073709551615) {
                                                var v74 int32 = 0
                                                _ = v74
                                            }
                                            _ = v72
                                        }
                                    }
                                    fmt.Println(fi67)
                                    delete(v1, ("" + "hello"))
                                    v13 = v13
                                    v75 := GenS5{F0: false, F1: 0}
                                    v70 -= (v70 ^ (fi67 | fi67))
                                    fmt.Println(v70, v11)
                                    switch (("\n\t" + "hello") + ("hello" + ("hello" + "foo bar"))) {
                                    case "foo bar":
                                        var v76 float64 = (1.5 + ((v11 - v3) + v3))
                                        _ = v76
                                    case "a":
                                        v11 += v3
                                    }
                                    _ = v69
                                    _ = v70
                                    _ = v75
                                }
                                _ = fi67
                            }
                            v3 = v11
                            if (((v3 < 1e10) || false) || (-128 != 127)) {
                                var v77 uint16 = 42
                                _ = v77
                            }
                            _ = fi66
                        }
                        if true {
                            v78 := GenS2{F0: (v4 * (v4 << (((^(v12))) & 63))), F1: 51, F2: 109}
                            if (false || false) {
                                var v79 bool = true
                                v79 = false
                                var v80 int8 = 127
                                _ = v79
                                _ = v80
                            }
                            v4 = (v4 ^ (v4 * 286))
                            _ = v78
                        }
                    } else {
                        v12 = (v12 % ((v12) | 1))
                    }
                }
            } else {
                v13 -= v13
                var v81 int = -1
                _ = v81
            }
        default:
            var v82 uint16 = 32767
            _ = v82
        }
    }
    if (0 != 346) {
        var v83 float32 = 3.14159
        if ((^(v12)) == v12) {
            var v84 int32 = 42
            for fi85 := 0; fi85 < 8; fi85++ {
                if (!true) {
                    switch 0 {
                    case -172:
                        v12 -= v12
                    case -1:
                        if (2 < 2) {
                            fmt.Println(v12, v3)
                        } else {
                            v13 = v13
                            var v86 int64 = v4
                            _ = v86
                        }
                    case -831:
                        if (v13 <= (v13 % ((v13) | 1))) {
                            v87 := v1["a"]
                            v3 = -2.25
                            switch ("\n\t" + "hello") {
                            case "":
                                var v88 int32 = v84
                                _ = v88
                            }
                            _ = v87
                        } else {
                            for fi89 := 0; fi89 < 3; fi89++ {
                                v11 = (float64(((v83 * (v83 + v83)) - (((v83 / v83) * 3.14159) / v83))) + v11)
                                _ = fi89
                            }
                        }
                    default:
                        fmt.Println(v3, v11)
                    }
                    v83 += v83
                }
                _ = fi85
            }
            if false {
                v90 := v1[("" + "\n\t")]
                _ = v90
            }
            _ = v84
        } else {
            v12 = v12
        }
        _ = v83
    } else {
        switch int(v4) {
        case 2147483647:
            v13 = (^(v13))
        case -9223372036854775808:
            if len(v2) > 0 {
                v2[0] = ""
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
