package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint; F1 int64; F2 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int16; F1 bool }
type GenS7 struct { F0 uint; F1 int; F2 uint16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint16 }

func main() {
    v1 := []string{}
    var v2 float64 = (3.14159 + 0.0)
    var v3 int64 = -9223372036854775808
    var v4 int64 = (^(v3))
    var v5 int32 = 2147483647
    var v6 int32 = 650
    v3 = (v4 ^ (v3 * -795))
    for fi7 := 0; fi7 < 7; fi7++ {
        if (1 > 1) {
            if (1 == 42) {
                var v8 string = "foo bar"
                _ = v8
            }
        }
        _ = fi7
    }
    switch ("foo bar" + (("" + "a") + ("hello" + "hello"))) {
    case "foo bar":
        v4 -= v4
    default:
        v3 -= v3
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        for fi10 := 0; fi10 < 17; fi10++ {
            v11 := GenS1{F0: (fi9 * 1), F1: "foo bar"}
            v2 += float64(((0.1 + -2.25) + 3.14159))
            v12 := GenS2{F0: -109, F1: (!true)}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    v13 := GenS2{F0: -1, F1: (((true || (4294967295 > -677)) || true) && (("" + "foo bar") == "a"))}
    v5 = (v6 + v5)
    v4 += v4
    switch 4294967295 {
    case 540:
        if (("" + "hello") > "\n\t") {
            var v14 uint = 201
            switch -9223372036854775808 {
            case 9223372036854775807:
                var v15 int16 = -774
                _ = v15
            case 0:
                v2 -= v2
            case -2147483648:
                v6 = -1
            default:
                switch "foo bar" {
                case "foo bar":
                    v16 := GenS3{F0: -200, F1: -114}
                    _ = v16
                }
            }
            v17 := GenS2{F0: 32767, F1: (v4 > (v3 << (((v14 ^ v14)) & 63)))}
            _ = v14
            _ = v17
        }
    case 992:
        if false {
            for fi18 := 0; fi18 < 19; fi18++ {
                var v19 bool = (!(false && true))
                if v19 {
                    v19 = true
                    for fi20 := 0; fi20 < 11; fi20++ {
                        switch fi18 {
                        case 32767:
                            var v21 string = "foo bar"
                            _ = v21
                        default:
                            if v19 {
                                fmt.Println(fi20, v3, v19)
                                switch (("" + "foo bar") + "hello") {
                                case "a":
                                    v22 := GenS1{F0: (fi20 & (fi20 | fi18)), F1: "a"}
                                    _ = v22
                                }
                                v4 += v3
                                var v23 uint32 = 4294967295
                                _ = v23
                            } else {
                                if v19 {
                                    v4 -= 42
                                    v19 = false
                                    if v19 {
                                        v1 = append(v1, "")
                                        fmt.Println(v19, v3, v3)
                                    } else {
                                        v6 += (v6 & (v6 ^ (v6 / ((42) | 1))))
                                    }
                                } else {
                                    fmt.Println(v4, fi18)
                                    switch (^(fi18)) {
                                    case 127:
                                        if v19 {
                                            fmt.Println(v6, fi20)
                                            v4 = 1
                                        } else {
                                            v2 += ((v2 + float64(3.14159)) + float64(0.0))
                                            v19 = (false || ((-1 != v3) || (!false)))
                                        }
                                    }
                                    switch (fi18 / ((fi18) | 1)) {
                                    case 128:
                                        v24 := GenS4{F0: 621, F1: v3, F2: true}
                                        var v25 bool = ("" >= ("a" + "a"))
                                        v26 := GenS5{F0: (v4 & (-(v4)))}
                                        if (fi20 == -2147483648) {
                                            if true {
                                                v27 := GenS6{F0: -571, F1: v25}
                                                v28 := GenS1{F0: fi20, F1: ("\n\t" + ("a" + "a"))}
                                                _ = v27
                                                _ = v28
                                            }
                                            var v29 int16 = -32768
                                            v29 = (v29 - 32767)
                                            _ = v29
                                        }
                                        if (3 <= 127) {
                                            if len(v1) > 0 {
                                                v1[0] = "\n\t"
                                            }
                                            v4 = -786
                                            if true {
                                                if true {
                                                    v4 = (v3 - (v4 / ((9223372036854775807) | 1)))
                                                }
                                                fmt.Println(v3, fi20)
                                                if len(v1) > 0 {
                                                    v1[0] = (((("hello" + "a") + "\n\t") + "\n\t") + ("a" + "hello"))
                                                }
                                            }
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            switch "\n\t" {
                                            case "hello":
                                                var v30 int32 = v5
                                                fmt.Println(v2, v30, v4)
                                                if (v19 && (4294967295 > 18446744073709551615)) {
                                                    if false {
                                                        v19 = v25
                                                    }
                                                }
                                                _ = v30
                                            default:
                                                switch (fi20 | fi20) {
                                                case 0:
                                                    var v31 float64 = v2
                                                    _ = v31
                                                case 2147483647:
                                                    if ((v5 >> ((18446744073709551615) & 63)) < v5) {
                                                        var v32 uint32 = 2147483648
                                                        if v19 {
                                                            var v33 int32 = 537
                                                            v34 := GenS7{F0: 42, F1: fi20, F2: 65535}
                                                            _ = v33
                                                            _ = v34
                                                        }
                                                        var v35 uint16 = 299
                                                        _ = v32
                                                        _ = v35
                                                    }
                                                default:
                                                    v36 := GenS8{F0: 0}
                                                    var v37 int32 = v5
                                                    _ = v36
                                                    _ = v37
                                                }
                                            }
                                        }
                                        _ = v24
                                        _ = v25
                                        _ = v26
                                    case 0:
                                        if v19 {
                                            var v38 float32 = 1e10
                                            var v39 uint = 65535
                                            v1 = append(v1, "foo bar")
                                            _ = v38
                                            _ = v39
                                        } else {
                                            var v40 bool = v19
                                            var v41 uint = 1
                                            _ = v40
                                            _ = v41
                                        }
                                    case 127:
                                        v42 := GenS9{F0: 65535}
                                        _ = v42
                                    default:
                                        if v19 {
                                            var v43 uint64 = 564
                                            _ = len(v1)
                                            switch fi18 {
                                            case 42:
                                                var v44 bool = (!false)
                                                _ = v44
                                            case 4294967295:
                                                if (false || false) {
                                                    var v45 string = "foo bar"
                                                    v46 := GenS2{F0: 32767, F1: v19}
                                                    var v47 bool = (-435 != -162)
                                                    _ = v45
                                                    _ = v46
                                                    _ = v47
                                                }
                                            default:
                                                if v19 {
                                                    var v48 uint32 = 2147483647
                                                    v2 = (0.1 * v2)
                                                    var v49 uint16 = 32767
                                                    fmt.Println(v43)
                                                    _ = v48
                                                    _ = v49
                                                }
                                            }
                                            fmt.Println(v5)
                                            _ = v43
                                        } else {
                                            _ = len(v1)
                                            var v50 float64 = (3.14159 / (((v2 / (v2 + v2)) / (float64(1e10) + v2)) * v2))
                                            if ((v4 & v4) != 1) {
                                                switch (((("\n\t" + "\n\t") + "hello") + "") + "\n\t") {
                                                case "foo bar":
                                                    v19 = false
                                                case "a":
                                                    fmt.Println(v3, v5)
                                                }
                                                switch ("" + "") {
                                                case "\n\t":
                                                    fmt.Println(v50)
                                                }
                                                var v51 uint32 = 4294967295
                                                _ = v51
                                            }
                                            _ = v50
                                        }
                                    }
                                }
                            }
                        }
                        _ = fi20
                    }
                } else {
                    v5 -= int32(v3)
                    var v52 bool = ("hello" >= "a")
                    _ = v52
                }
                _ = fi18
                _ = v19
            }
        }
    case -549:
        v2 -= v2
    default:
        for fi53 := 0; fi53 < 5; fi53++ {
            var v54 uint8 = 52
            _ = fi53
            _ = v54
        }
    }
    switch "a" {
    case "a":
        var v55 bool = (!(("" != ("hello" + ("" + "foo bar"))) || (!true)))
        _ = v55
    case "hello":
        var v56 int = 9223372036854775807
        _ = v56
    default:
        v6 = (v6 + (v6 % (((v5 ^ (-(v5)))) | 1)))
    }
    if true {
        var v57 float64 = ((v2 + 1e10) / 1.0)
        v58 := GenS9{F0: 42}
        v3 -= (v4 - v3)
        _ = v57
        _ = v58
    } else {
        if (18446744073709551615 > 4294967295) {
            switch "" {
            case "hello":
                if ((false || (!true)) && true) {
                    if (255 < 692) {
                        if (1 > 391) {
                            var v59 int32 = (v5 + 1)
                            _ = v59
                        } else {
                            fmt.Println(v4, v4, v3)
                        }
                        if len(v1) > 0 {
                            v1[0] = ("hello" + "")
                        }
                    } else {
                        if (254 <= 42) {
                            var v60 float32 = 0.1
                            _ = v60
                        }
                        v1 = append(v1, "hello")
                    }
                    var v61 float32 = -2.25
                    if (!(!(1 <= 100))) {
                        fmt.Println(v61, v4, v61)
                        if false {
                            fmt.Println(v6, v5, v6)
                        }
                        fmt.Println(v4, v6)
                    }
                    _ = v61
                }
            case "":
                for fi62 := 0; fi62 < 4; fi62++ {
                    var v63 float64 = 1e10
                    v2 = (-2.25 / v63)
                    v3 -= v4
                    _ = fi62
                    _ = v63
                }
            case "\n\t":
                var v64 uint16 = 238
                _ = v64
            default:
                var v65 int16 = -32768
                _ = v65
            }
        } else {
            for fi66 := 0; fi66 < 16; fi66++ {
                var v67 uint = 2
                var v68 float64 = 1.5
                for fi69 := 0; fi69 < 20; fi69++ {
                    if ((true && ((v6 - 614) < (-(v5)))) || (614 != 2)) {
                        if len(v1) > 0 {
                            v1[0] = ("foo bar" + "foo bar")
                        }
                    }
                    v4 = v4
                    _ = fi69
                }
                _ = fi66
                _ = v67
                _ = v68
            }
            for fi70 := 0; fi70 < 7; fi70++ {
                var v71 uint = 1
                v72 := GenS5{F0: (v4 % ((v4) | 1))}
                _ = fi70
                _ = v71
                _ = v72
            }
        }
    }
    v73 := GenS7{F0: 394, F1: 359, F2: 32767}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
}
