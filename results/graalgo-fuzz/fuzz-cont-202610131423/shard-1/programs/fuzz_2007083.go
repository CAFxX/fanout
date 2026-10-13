package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 int8; F1 uint32 }
type GenS6 struct { F0 int; F1 bool; F2 uint64 }
type GenS7 struct { F0 bool; F1 uint; F2 uint32 }

func main() {
    var v1 float64 = 1.5
    v2 := []string{"hello", ""}
    var v3 int16 = 32767
    var v4 uint8 = 255
    v3 = (v3 % (((v3 << ((4294967295) & 63))) | 1))
    var v5 int16 = v3
    for fi6 := 0; fi6 < 18; fi6++ {
        if false {
            v4 = v4
            if true {
                v1 += v1
                if ((v5 - 584) < v5) {
                    switch ("hello" + "a") {
                    case "foo bar":
                        for fi7 := 0; fi7 < 17; fi7++ {
                            if false {
                                v4 = (v4 * (v4 << ((2) & 63)))
                            } else {
                                switch ((("\n\t" + "") + "hello") + "foo bar") {
                                case "foo bar":
                                    var v8 int16 = (^(v3))
                                    _ = v8
                                default:
                                    fmt.Println(v5, fi7)
                                }
                                v9 := GenS1{F0: fi6}
                                var v10 float32 = 3.14159
                                var v11 float64 = v1
                                v2 = append(v2, "hello")
                                v10 -= v10
                                v12 := GenS2{F0: 852, F1: 18446744073709551615}
                                _ = v10
                                _ = v11
                                _ = v12
                                _ = v9
                            }
                            v5 = 42
                            v1 = float64((((-2.25 - -2.25) * 3.14159) * 1.5))
                            v5 += 186
                            var v13 int32 = -1
                            var v14 float32 = ((0.0 + ((3.14159 * -2.25) / 1.0)) - float32(v1))
                            _ = fi7
                            _ = v13
                            _ = v14
                        }
                    case "":
                        var v15 int32 = -2147483648
                        _ = v15
                    }
                    v16 := GenS2{F0: 42, F1: 4294967295}
                    v17 := GenS2{F0: 30, F1: uint(v4)}
                    _ = v16
                    _ = v17
                } else {
                    switch "a" {
                    case "hello":
                        for fi18 := 0; fi18 < 13; fi18++ {
                            v19 := GenS2{F0: 2, F1: 65535}
                            fmt.Println(v1, fi6, fi6)
                            _ = fi18
                            _ = v19
                        }
                    }
                    switch "" {
                    case "hello":
                        v5 -= v3
                    case "a":
                        v5 = 32767
                    case "":
                        var v20 float32 = 0.0
                        _ = v20
                    }
                    v21 := GenS2{F0: 735, F1: 255}
                    _ = v21
                }
                v2 = append(v2, "foo bar")
            }
            v22 := GenS3{F0: (((float64(((float32((v1 - v1)) * 3.14159) / (-2.25 + (0.0 / -2.25)))) / float64((((float32(-2.25) / -2.25) + 0.0) / 1.5))) * (v1 + float64(-2.25))) != float64((1.5 * 0.0))), F1: -80}
            _ = v22
        } else {
            if (true && true) {
                if (42 != 100) {
                    fmt.Println(v3, v1)
                } else {
                    switch (fi6 | 9223372036854775807) {
                    case -1:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case 19:
                        var v23 int8 = 42
                        _ = v23
                    case -9223372036854775808:
                        for fi24 := 0; fi24 < 11; fi24++ {
                            if true {
                                var v25 int32 = 42
                                var v26 uint64 = 42
                                var v27 bool = (-97 >= -1)
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                            _ = fi24
                        }
                    default:
                        if (("a" + (("" + ("\n\t" + ("" + (("a" + (("" + "") + "foo bar")) + "a")))) + "foo bar")) < ("\n\t" + "hello")) {
                            v2 = append(v2, "")
                            for fi28 := 0; fi28 < 15; fi28++ {
                                if true {
                                    if (false && (-128 != -128)) {
                                        v5 = (v3 % (((v5 >> ((42) & 63))) | 1))
                                        switch fi6 {
                                        case 4294967295:
                                            v29 := GenS4{F0: 32767, F1: 65535}
                                            _ = v29
                                        default:
                                            v5 = v3
                                        }
                                        var v30 uint64 = 18446744073709551615
                                        _ = v30
                                    } else {
                                        if (1 > 18446744073709551615) {
                                            fmt.Println(v4, v3)
                                            var v31 int8 = 0
                                            _ = v31
                                        }
                                        v1 -= (v1 + ((v1 + 3.14159) - v1))
                                        var v32 int = -9223372036854775808
                                        _ = v32
                                    }
                                    v33 := GenS3{F0: (!true), F1: -128}
                                    _ = v33
                                }
                                _ = fi28
                            }
                            fmt.Println(v1, v4, v3)
                        }
                    }
                }
                fmt.Println(v3, v3)
                fmt.Println(v3)
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v4 = v4
        }
        var v34 uint = 65535
        _ = fi6
        _ = v34
    }
    var v35 float64 = ((1e10 - v1) / v1)
    var v36 int64 = 635
    switch 128 {
    case -1:
        if true {
            v36 += v36
            var v37 int = 4294967295
            _ = v37
        }
    case 4294967295:
        fmt.Println(v5)
    }
    if false {
        var v38 float64 = 1.5
        _ = v38
    }
    v5 = v5
    for fi39 := 0; fi39 < 19; fi39++ {
        v35 = v35
        _ = fi39
    }
    if (v35 >= (v35 - v35)) {
        v40 := GenS5{F0: -128, F1: 1}
        _ = v40
    } else {
        v41 := GenS5{F0: 42, F1: 979}
        _ = v41
    }
    switch ((("foo bar" + "hello") + "\n\t") + (("foo bar" + ("foo bar" + "a")) + "\n\t")) {
    case "\n\t":
        if true {
            if false {
                for fi42 := 0; fi42 < 11; fi42++ {
                    for fi43 := 0; fi43 < 7; fi43++ {
                        if (1e10 <= 0.0) {
                            var v44 uint16 = 1
                            _ = v44
                        }
                        var v45 int64 = (^(v36))
                        var v46 int16 = 1
                        v5 += v3
                        if true {
                            v35 += (v35 + (((3.14159 * float64(1e10)) * v35) + 0.1))
                            if (((^(v4)) == (v4 >> ((4294967295) & 63))) && (((false || false) && true) && true)) {
                                v45 = (^(v36))
                                v1 -= v35
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        }
                        _ = fi43
                        _ = v45
                        _ = v46
                    }
                    switch (-(fi42)) {
                    case 0:
                        if (0 <= 255) {
                            v3 = 620
                            if (uint16(v4) != 42) {
                                var v47 float32 = (1e10 + -2.25)
                                var v48 int64 = (v36 + v36)
                                switch (fi42 % ((fi42) | 1)) {
                                case 0:
                                    for fi49 := 0; fi49 < 7; fi49++ {
                                        v4 += (v4 >> ((302) & 63))
                                        if len(v2) > 0 {
                                            v2[0] = ("\n\t" + (("\n\t" + ("" + "a")) + "foo bar"))
                                        }
                                        _ = fi49
                                    }
                                case -9223372036854775808:
                                    var v50 int = fi42
                                    _ = v50
                                case 4294967295:
                                    v47 += v47
                                default:
                                    var v51 int32 = -1
                                    _ = v51
                                }
                                _ = v47
                                _ = v48
                            } else {
                                switch "hello" {
                                case "foo bar":
                                    var v52 uint64 = 42
                                    _ = v52
                                case "hello":
                                    if true {
                                        var v53 uint64 = 267
                                        _ = v53
                                    } else {
                                        v36 += v36
                                        var v54 bool = ((!(true || false)) && ((-806 <= (fi42 >> ((700) & 63))) || (((!false) && true) || false)))
                                        _ = v54
                                    }
                                }
                            }
                            switch "\n\t" {
                            case "a":
                                var v55 uint64 = 4294967295
                                _ = v55
                            case "hello":
                                if false {
                                    v56 := GenS6{F0: (fi42 % (((fi42 / ((688) | 1))) | 1)), F1: (-1 != 127), F2: 4294967295}
                                    _ = v56
                                } else {
                                    var v57 float32 = (0.1 + float32(v1))
                                    if (42 == fi42) {
                                        var v58 float32 = v57
                                        _ = v58
                                    } else {
                                        for fi59 := 0; fi59 < 13; fi59++ {
                                            fmt.Println(v1)
                                            if true {
                                                switch "\n\t" {
                                                case "\n\t":
                                                    v36 = -925
                                                    v5 = v3
                                                    v5 = (-(v5))
                                                case "foo bar":
                                                    v3 -= v5
                                                case "a":
                                                    fmt.Println(v5, v35)
                                                    switch fi59 {
                                                    case -5:
                                                        var v60 int16 = 32767
                                                        _ = v60
                                                    case -1:
                                                        v5 -= (v3 + 0)
                                                    }
                                                    fmt.Println(fi59, v57)
                                                    v5 -= v5
                                                default:
                                                    var v61 uint64 = 4294967295
                                                    _ = v61
                                                }
                                            } else {
                                                v35 = ((v1 + v35) - -2.25)
                                                v35 -= (1.5 / 0.1)
                                                switch fi42 {
                                                case 9223372036854775807:
                                                    v62 := GenS7{F0: false, F1: 1, F2: 2147483648}
                                                    _ = v62
                                                case -811:
                                                    fmt.Println(fi42, v3, fi42)
                                                case 42:
                                                    if (false && false) {
                                                        var v63 string = ""
                                                        v64 := GenS7{F0: true, F1: 241, F2: 0}
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                default:
                                                    v65 := GenS7{F0: (!true), F1: 18446744073709551615, F2: 2}
                                                    _ = v65
                                                }
                                                var v66 bool = true
                                                _ = v66
                                            }
                                            _ = fi59
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = "foo bar"
                                        }
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v57
                                }
                            default:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        }
                    default:
                        switch "a" {
                        case "a":
                            for fi67 := 0; fi67 < 8; fi67++ {
                                v68 := GenS7{F0: (((("" != "a") && false) || true) || (-967 != 42)), F1: 65535, F2: 457}
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi67
                                _ = v68
                            }
                        case "foo bar":
                            for fi69 := 0; fi69 < 5; fi69++ {
                                var v70 float32 = 1.5
                                _ = fi69
                                _ = v70
                            }
                        case "":
                            var v71 float32 = -2.25
                            _ = v71
                        default:
                            v1 = 1e10
                        }
                    }
                    _ = fi42
                }
                v72 := GenS4{F0: 392, F1: 259}
                v73 := GenS4{F0: 2, F1: 2}
                _ = v72
                _ = v73
            } else {
                switch -649 {
                case 0:
                    v74 := GenS4{F0: 42, F1: 659}
                    _ = v74
                case 32768:
                    var v75 uint = 0
                    _ = v75
                }
                v2 = append(v2, ((("foo bar" + ("a" + "\n\t")) + "\n\t") + ("\n\t" + "a")))
                var v76 bool = (1 == 65535)
                _ = v76
            }
            if true {
                var v77 uint = 42
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
                _ = v77
            }
        } else {
            for fi78 := 0; fi78 < 15; fi78++ {
                var v79 uint16 = 42
                for fi80 := 0; fi80 < 2; fi80++ {
                    v1 = v1
                    if len(v2) > 0 {
                        v2[0] = ""
                    }
                    _ = fi80
                }
                _ = fi78
                _ = v79
            }
            fmt.Println(v5, v5, v36)
        }
    }
    v1 -= 0.0
    _ = len(v2)
    v1 = (((v1 * -2.25) * (float64(0.0) - (3.14159 - v35))) + v35)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
