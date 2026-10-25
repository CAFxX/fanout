package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int16; F1 bool; F2 uint64 }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 int; F1 uint16; F2 int32 }
func genFn0(p0 int, p1 uint16) int16 {
    v4 := GenS1{F0: p1, F1: 0, F2: -792}
    v5 := []int{p0, (p0 | p0), 216, 4294967295}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 bool, p1 int32, p2 string) uint {
    fmt.Println(p0, p0, p1)
    var v6 int16 = 1
    if p0 {
        if (p0 && (((2 == 2) && true) && (-2.25 < (0.0 - (((3.14159 - float32((3.14159 * (float64(float32(1e10)) + 1e10)))) * (3.14159 + ((-2.25 * 0.1) - (-2.25 - 1e10)))) * 0.1))))) {
            for fi7 := 0; fi7 < 3; fi7++ {
                for fi8 := 0; fi8 < 19; fi8++ {
                    switch fi7 {
                    case 143:
                        var v9 float32 = float32(-2.25)
                        _ = v9
                    case 127:
                        switch p2 {
                        case "\n\t":
                            fmt.Println(p1, p2)
                        case "foo bar":
                            v6 -= -1
                        }
                    }
                    _ = fi8
                }
                v10 := []int{(fi7 / (((fi7 - 32767)) | 1))}
                switch (fi7 + (-(fi7))) {
                case 32768:
                    p2 = ("a" + p2)
                case -2147483648:
                    var v11 uint8 = 127
                    _ = v11
                case -933:
                    for fi12 := 0; fi12 < 15; fi12++ {
                        p0 = p0
                        fmt.Println(p0, p2)
                        _ = fi12
                    }
                }
                _ = fi7
                _ = v10
            }
        } else {
            if ((-1 == -178) && false) {
                switch 0 {
                case 127:
                    var v13 int8 = -1
                    _ = v13
                default:
                    v14 := []int{-9223372036854775808}
                    _ = v14
                }
            }
            p2 = p2
            var v15 int = 1
            _ = v15
        }
    } else {
        var v16 int16 = v6
        v16 -= genFn0(32767, 0)
        _ = v16
    }
    if (!((-1 < -1) && (true || p0))) {
        v17 := GenS1{F0: 0, F1: 2, F2: p1}
        var v18 int32 = (p1 << ((684) & 63))
        v19 := GenS1{F0: 65535, F1: 694, F2: p1}
        _ = v17
        _ = v18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 596
}

func main() {
    var v1 uint = 234
    v2 := []string{"foo bar", ("foo bar" + "a"), "\n\t"}
    var v3 int32 = 814
    for fi20 := 0; fi20 < 10; fi20++ {
        switch "foo bar" {
        case "":
            var v21 int8 = 100
            _ = v21
        case "\n\t":
            v3 = (-(v3))
        case "a":
            var v22 float64 = ((3.14159 + (0.0 - 1e10)) - -2.25)
            _ = v22
        }
        _ = fi20
    }
    v2 = append(v2, ("\n\t" + "\n\t"))
    for fi23 := 0; fi23 < 8; fi23++ {
        var v24 int8 = 42
        _ = fi23
        _ = v24
    }
    switch ("\n\t" + "hello") {
    case "a":
        v25 := GenS1{F0: 71, F1: 827, F2: -2147483648}
        _ = v25
    }
    v1 -= v1
    for fi26 := 0; fi26 < 4; fi26++ {
        switch (((((("a" + "foo bar") + "") + "a") + "") + ("\n\t" + ("\n\t" + "\n\t"))) + "\n\t") {
        case "hello":
            switch (fi26 / ((fi26) | 1)) {
            case -9223372036854775808:
                v27 := GenS1{F0: 65535, F1: 0, F2: 417}
                _ = v27
            case 128:
                if (-417 >= 1) {
                    switch "\n\t" {
                    case "\n\t":
                        var v28 uint8 = 0
                        _ = v28
                    default:
                        var v29 float64 = -2.25
                        _ = v29
                    }
                    _ = len(v2)
                    v30 := GenS2{F0: 4294967295}
                    _ = v30
                }
            case 0:
                var v31 int = -100
                _ = v31
            default:
                var v32 int = (fi26 - -2147483648)
                _ = v32
            }
        case "":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "a":
            var v33 uint32 = 495
            _ = v33
        default:
            for fi34 := 0; fi34 < 11; fi34++ {
                fmt.Println(fi34)
                v3 = v3
                _ = fi34
            }
        }
        _ = fi26
    }
    switch 42 {
    case -9223372036854775808:
        if false {
            fmt.Println(v3, v3)
        } else {
            switch -1 {
            case 0:
                for fi35 := 0; fi35 < 1; fi35++ {
                    v3 = v3
                    var v36 int16 = -1
                    _ = fi35
                    _ = v36
                }
            case -2147483648:
                fmt.Println(v1, v3)
            case 210:
                v37 := GenS1{F0: 1, F1: 0, F2: (v3 | (v3 - (v3 & (v3 | 0))))}
                _ = v37
            }
        }
    case 32768:
        fmt.Println(v1)
    case -466:
        for fi38 := 0; fi38 < 15; fi38++ {
            var v39 uint64 = 40
            fmt.Println(v39, v1)
            v1 -= genFn1((567 >= 42), v3, "")
            _ = fi38
            _ = v39
        }
    }
    v40 := GenS3{F0: -1, F1: 42, F2: 2147483647}
    fmt.Println(v3, v1)
    if (-9223372036854775808 >= 589) {
        var v41 uint64 = 171
        if ((!(956 <= 296)) || false) {
            var v42 uint16 = 61
            for fi43 := 0; fi43 < 9; fi43++ {
                for fi44 := 0; fi44 < 9; fi44++ {
                    fmt.Println(v41, v3)
                    v45 := GenS4{F0: -161, F1: (!(128 < 224)), F2: (v41 << ((v1) & 63))}
                    var v46 uint16 = 42
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                _ = fi43
            }
            var v47 bool = false
            _ = v42
            _ = v47
        }
        _ = v41
    } else {
        fmt.Println(v1, v3)
        v1 = 255
    }
    var v48 uint8 = 196
    v1 += v1
    for fi49 := 0; fi49 < 12; fi49++ {
        if (0.1 != (0.1 + (3.14159 - ((1e10 * float32(1e10)) + float32(3.14159))))) {
            for fi50 := 0; fi50 < 18; fi50++ {
                var v51 int8 = -119
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
                _ = fi50
                _ = v51
            }
        } else {
            _ = len(v2)
        }
        for fi52 := 0; fi52 < 14; fi52++ {
            if false {
                switch (fi52 >> ((uint(v48)) & 63)) {
                case -952:
                    v3 -= 972
                case 144:
                    if (false && (0 != 710)) {
                        if (0 >= 563) {
                            switch -1 {
                            case 2147483647:
                                if len(v2) > 0 {
                                    v2[0] = ("" + "hello")
                                }
                            case 32767:
                                v2 = append(v2, "hello")
                            case -1:
                                switch (fi49 & fi49) {
                                case 128:
                                    v2 = append(v2, "")
                                case -1:
                                    var v53 int = fi49
                                    _ = v53
                                default:
                                    var v54 uint64 = 521
                                    _ = v54
                                }
                            default:
                                if (!true) {
                                    v3 -= -2147483648
                                    fmt.Println(fi52, v3, fi52)
                                } else {
                                    if ((("foo bar" + "") == ("a" + ("a" + ""))) || false) {
                                        var v55 uint64 = 42
                                        v55 = uint64(v48)
                                        _ = v55
                                    }
                                    v1 += v1
                                    var v56 string = ("\n\t" + "")
                                    _ = v56
                                }
                            }
                        }
                    } else {
                        switch fi49 {
                        case 42:
                            if true {
                                switch fi52 {
                                case 32767:
                                    var v57 float64 = ((-2.25 + (float64((1e10 * float32((1e10 - 1e10)))) - 3.14159)) - 0.1)
                                    _ = v57
                                }
                                v58 := GenS1{F0: 1, F1: 1, F2: (v3 ^ -2147483648)}
                                switch "hello" {
                                case "foo bar":
                                    v3 = (v3 & v3)
                                default:
                                    v1 += v1
                                }
                                _ = v58
                            } else {
                                if (-84 >= int64(v1)) {
                                    var v59 bool = false
                                    _ = v59
                                }
                                if true {
                                    v1 += (^(v1))
                                    if (genFn1(true, v3, "\n\t") >= genFn1((254 == 1), (-(v3)), "")) {
                                        switch "a" {
                                        case "foo bar":
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        case "":
                                            v60 := GenS4{F0: -1, F1: (0 != -32768), F2: 1}
                                            var v61 uint8 = 63
                                            switch "foo bar" {
                                            case "hello":
                                                var v62 uint16 = 65535
                                                _ = v62
                                            case "a":
                                                var v63 int = (fi49 & fi52)
                                                _ = v63
                                            default:
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                            var v64 int = (-(fi52))
                                            var v65 int16 = -32768
                                            v66 := GenS1{F0: 1, F1: 992, F2: -29}
                                            _ = v60
                                            _ = v61
                                            _ = v64
                                            _ = v65
                                            _ = v66
                                        case "hello":
                                            fmt.Println(v1, v48)
                                        default:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                    }
                                } else {
                                    _ = len(v2)
                                    _ = len(v2)
                                }
                                v3 += (v3 >> ((0) & 63))
                            }
                        default:
                            v2 = append(v2, "foo bar")
                        }
                        v48 -= v48
                        if (1.5 > (0.1 - 1e10)) {
                            fmt.Println(fi52, v48)
                            v67 := GenS3{F0: 127, F1: 1, F2: 1}
                            v68 := GenS4{F0: 42, F1: (!true), F2: 42}
                            _ = v67
                            _ = v68
                        } else {
                            if true {
                                v69 := GenS4{F0: 1000, F1: (!(false || (!(0 >= 42)))), F2: 0}
                                _ = v69
                            }
                        }
                    }
                case 32768:
                    var v70 float32 = 1e10
                    _ = v70
                default:
                    v48 -= uint8(v1)
                }
                var v71 int32 = 983
                if false {
                    switch fi49 {
                    case 127:
                        v1 += v1
                    case 42:
                        if ("hello" != "foo bar") {
                            v3 -= (v71 << (((v1 - v1)) & 63))
                            var v72 string = "foo bar"
                            _ = v72
                        }
                    }
                    v73 := GenS5{F0: (!true), F1: 18446744073709551615}
                    v48 += v48
                    _ = v73
                } else {
                    fmt.Println(v1, v71)
                    switch 9223372036854775807 {
                    case -582:
                        v2 = append(v2, "foo bar")
                        v74 := GenS5{F0: (false || (float32(float64((-2.25 - ((1e10 / 1.0) - (0.0 / 3.14159))))) >= (3.14159 - (0.0 + 1.5)))), F1: 2}
                        _ = v74
                    case -9223372036854775808:
                        v3 = v3
                    case -1:
                        var v75 uint = (v1 + v1)
                        _ = v75
                    default:
                        switch fi49 {
                        case 303:
                            fmt.Println(fi52, fi49, v3)
                        case 461:
                            fmt.Println(v71)
                        }
                    }
                }
                if false {
                    v2 = append(v2, ("foo bar" + "\n\t"))
                } else {
                    v1 = v1
                }
                var v76 string = "foo bar"
                _ = v71
                _ = v76
            } else {
                v2 = append(v2, ("foo bar" + "foo bar"))
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            switch (fi49 << ((802) & 63)) {
            case 2147483647:
                var v77 uint32 = 42
                _ = v77
            default:
                if true {
                    var v78 string = "\n\t"
                    v79 := GenS6{F0: (^(fi52)), F1: 790, F2: (v3 % (((^(v3))) | 1))}
                    v3 = 0
                    _ = v78
                    _ = v79
                }
            }
            _ = fi52
        }
        fmt.Println(v3, v48, v3)
        _ = fi49
    }
    _ = len(v2)
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v48
}
