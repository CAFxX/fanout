package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 int16; F2 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 int32 }
type GenS6 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS7 struct { F0 string; F1 uint8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint64; F1 int; F2 int32 }
func genFn0(p0 string, p1 int, p2 string) int64 {
    fmt.Println(p0, p1)
    v7 := GenS1{F0: 32767, F1: 832}
    p2 = "foo bar"
    p2 = "a"
    switch p0 {
    case "":
        v8 := []int{(p1 % ((p1) | 1)), (p1 ^ -2147483648)}
        _ = v8
    case "a":
        p1 = 42
    case "hello":
        p0 = ""
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    var v1 uint64 = 149
    v2 := []uint32{2147483647, 65535, 14, 807}
    var v3 int64 = 0
    var v4 int = 2147483647
    v5 := make(map[string]uint32)
    var v6 bool = false
    switch (v4 >> ((42) & 63)) {
    case 1:
        if true {
            var v9 float64 = 0.0
            if (v6 && false) {
                if true {
                    v10 := GenS1{F0: 2, F1: 32767}
                    fmt.Println(v3)
                    _ = v10
                }
            }
            switch (v4 % (((v4 ^ (v4 ^ v4))) | 1)) {
            case -2147483648:
                if v6 {
                    var v11 int16 = 32767
                    var v12 float32 = (float32(-2.25) / float32(((0.1 - float64((3.14159 * 0.0))) - float64(3.14159))))
                    _ = v11
                    _ = v12
                }
            default:
                var v13 int32 = 42
                _ = v13
            }
            _ = v9
        }
    case -100:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        if v6 {
            var v14 int = v4
            v14 -= int(v3)
            _ = v14
        } else {
            if (53 > -27) {
                var v15 int32 = -1
                v3 += v3
                _ = v15
            }
        }
    }
    v16 := GenS1{F0: 42, F1: 855}
    var v17 int64 = v3
    for fi18 := 0; fi18 < 8; fi18++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v19 := GenS2{F0: 32767}
        for fi20 := 0; fi20 < 8; fi20++ {
            fmt.Println(fi18, v6)
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        fmt.Println(v17, v1)
        var v22 uint32 = uint32(v1)
        var v23 bool = v6
        _ = fi21
        _ = v22
        _ = v23
    }
    if (1 < 2147483647) {
        switch -1 {
        case 42:
            v6 = ((v1 * v1) <= v1)
        case 128:
            var v24 int = v4
            _ = v24
        }
        var v25 string = "\n\t"
        if ((^(v17)) <= v17) {
            _ = len(v2)
            fmt.Println(v1)
        } else {
            v3 = genFn0((v25 + v25), (v4 / ((2147483647) | 1)), v25)
            v6 = v6
        }
        _ = v25
    } else {
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        for fi26 := 0; fi26 < 1; fi26++ {
            v2 = append(v2, 312)
            if (!(128 == 42)) {
                if (384 < 172) {
                    switch fi26 {
                    case 1:
                        v5[((("" + ("\n\t" + (("hello" + "\n\t") + ""))) + "a") + "\n\t")] = 597
                    }
                    v4 += v4
                    switch ("hello" + "foo bar") {
                    case "\n\t":
                        var v27 int16 = -326
                        _ = v27
                    case "hello":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "a":
                        fmt.Println(v1, v1)
                    default:
                        v4 = (fi26 ^ (^(v4)))
                    }
                } else {
                    v28 := GenS3{F0: (fi26 >> ((119) & 63)), F1: 32767, F2: 42}
                    if ("\n\t" <= "") {
                        v29 := GenS1{F0: 420, F1: 951}
                        var v30 uint16 = 162
                        switch 42 {
                        case 9223372036854775807:
                            for fi31 := 0; fi31 < 9; fi31++ {
                                var v32 string = "foo bar"
                                _ = fi31
                                _ = v32
                            }
                        case 1:
                            if (((1.5 - 3.14159) + 1.5) == (3.14159 * ((float64(float32((-2.25 + (1e10 * 0.0)))) / -2.25) * (-2.25 - (0.1 * 0.1))))) {
                                v33 := GenS4{F0: "hello"}
                                if v6 {
                                    var v34 float32 = 0.1
                                    v34 = v34
                                    var v35 uint32 = 0
                                    _ = v34
                                    _ = v35
                                }
                                _ = v33
                            }
                        }
                        _ = v29
                        _ = v30
                    } else {
                        for fi36 := 0; fi36 < 14; fi36++ {
                            var v37 uint = 42
                            _ = fi36
                            _ = v37
                        }
                        v38 := GenS5{F0: v6, F1: 42}
                        _ = v38
                    }
                    _ = v28
                }
            }
            if v6 {
                switch (fi26 + fi26) {
                case 1:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case 2147483647:
                    var v39 int64 = (v3 | (^(v3)))
                    _ = v39
                case 0:
                    if (v1 > v1) {
                        if (1 >= 1) {
                            if (!v6) {
                                var v40 int32 = 2147483647
                                _ = v40
                            }
                            _ = len(v2)
                        }
                    }
                default:
                    v6 = false
                }
                v41 := GenS1{F0: 689, F1: 0}
                _ = v41
            } else {
                v4 = (-(fi26))
                if v6 {
                    for fi42 := 0; fi42 < 10; fi42++ {
                        v43 := GenS1{F0: 32767, F1: 0}
                        switch fi26 {
                        case -1:
                            fmt.Println(v1, fi42)
                        }
                        if (v6 || v6) {
                            v44 := GenS5{F0: (!v6), F1: 0}
                            _ = v44
                        } else {
                            fmt.Println(fi42)
                            fmt.Println(fi42)
                            switch (fi42 / (((fi42 + 2147483647)) | 1)) {
                            case 128:
                                if false {
                                    var v45 int8 = -1
                                    _ = v45
                                } else {
                                    v1 = v1
                                    switch v4 {
                                    case 60:
                                        v6 = (-1 <= -128)
                                        if v6 {
                                            v6 = false
                                            var v46 int32 = 1
                                            switch 1 {
                                            case 0:
                                                v1 += 2
                                            case -622:
                                                switch (fi26 / (((-(fi26))) | 1)) {
                                                case 2147483647:
                                                    var v47 uint64 = 42
                                                    _ = v47
                                                case 0:
                                                    var v48 string = "a"
                                                    _ = v48
                                                case 4294967295:
                                                    v3 = genFn0("hello", -195, ("a" + "a"))
                                                    delete(v5, "foo bar")
                                                    v4 = fi26
                                                    v49 := GenS6{F0: -21, F1: 42, F2: 42}
                                                    _ = v49
                                                }
                                            default:
                                                fmt.Println(v1, v46, v17)
                                            }
                                            _ = v46
                                        }
                                        v5[""] = 65535
                                        var v50 uint = 176
                                        _ = v50
                                    case 42:
                                        v5["hello"] = 0
                                    default:
                                        delete(v5, ("" + "\n\t"))
                                        v17 += v17
                                    }
                                }
                            }
                        }
                        fmt.Println(v4, v1)
                        var v51 uint = 386
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi42
                        _ = v43
                        _ = v51
                    }
                    var v52 string = ""
                    _ = v52
                } else {
                    _ = len(v5)
                }
                for fi53 := 0; fi53 < 13; fi53++ {
                    if (!false) {
                        v2 = append(v2, 42)
                        v1 -= 2
                        if v6 {
                            if (2147483648 >= 2147483647) {
                                v4 = fi53
                                v3 = (^(v17))
                                if (("a" + "a") > "\n\t") {
                                    var v54 string = "foo bar"
                                    v6 = v6
                                    if true {
                                        _ = len(v5)
                                        var v55 uint16 = 1
                                        var v56 string = ("hello" + "a")
                                        _ = v55
                                        _ = v56
                                    } else {
                                        if len(v2) > 0 {
                                            v2[0] = 65535
                                        }
                                        v4 = (v4 ^ 0)
                                        fmt.Println(v1, fi53, v54)
                                        _ = len(v5)
                                        v54 = ""
                                    }
                                    _ = v54
                                } else {
                                    v3 = 9223372036854775807
                                    var v57 uint = 27
                                    _ = v57
                                }
                            } else {
                                var v58 int8 = 57
                                if true {
                                    v17 = 0
                                    fmt.Println(v58, v1, v4)
                                }
                                if true {
                                    v2 = append(v2, 42)
                                }
                                fmt.Println(v1, v58)
                                _ = v58
                            }
                            var v59 int = fi26
                            v3 -= (v17 * (^(v3)))
                            _ = v59
                        } else {
                            switch (fi26 & v4) {
                            case -292:
                                var v60 string = ""
                                _ = v60
                            }
                            if (v1 == 42) {
                                switch "hello" {
                                case "hello":
                                    if (1.5 >= 0.0) {
                                        _ = len(v2)
                                        v61 := GenS7{F0: "\n\t", F1: 191}
                                        var v62 float64 = -2.25
                                        _ = v61
                                        _ = v62
                                    }
                                case "":
                                    switch ("foo bar" + "") {
                                    case "":
                                        v6 = false
                                    default:
                                        var v63 uint8 = 255
                                        _ = v63
                                    }
                                case "foo bar":
                                    v6 = v6
                                default:
                                    v64 := GenS8{F0: v3}
                                    _ = v64
                                }
                                fmt.Println(fi26, v6, v6)
                                switch ("" + (("" + ("\n\t" + ("" + ("" + "foo bar")))) + (((("hello" + "a") + ("a" + "foo bar")) + "a") + "foo bar"))) {
                                case "hello":
                                    var v65 int = 16
                                    _ = v65
                                case "a":
                                    v3 = v17
                                case "foo bar":
                                    var v66 int16 = 32767
                                    _ = v66
                                default:
                                    switch 1 {
                                    case 32768:
                                        v4 = v4
                                    case 42:
                                        switch (fi53 * (fi53 >> ((255) & 63))) {
                                        case 128:
                                            v67 := GenS6{F0: -44, F1: v1, F2: 846}
                                            _ = v67
                                        case 127:
                                            switch -9223372036854775808 {
                                            case 370:
                                                switch "foo bar" {
                                                case "foo bar":
                                                    switch "foo bar" {
                                                    case "":
                                                        var v68 string = "a"
                                                        _ = v68
                                                    }
                                                case "\n\t":
                                                    switch "\n\t" {
                                                    case "a":
                                                        var v69 string = ("foo bar" + ("hello" + ""))
                                                        _ = v69
                                                    }
                                                default:
                                                    var v70 int32 = 0
                                                    _ = v70
                                                }
                                            case -669:
                                                v6 = (127 == 100)
                                                var v71 int = (-(fi53))
                                                fmt.Println(v3)
                                                _ = v71
                                            }
                                        case -2147483648:
                                            if ((v1 >> ((255) & 63)) >= v1) {
                                                if false {
                                                    v17 -= v17
                                                    v1 = v1
                                                } else {
                                                    var v72 float32 = -2.25
                                                    _ = v72
                                                }
                                            } else {
                                                v2 = append(v2, 2147483647)
                                            }
                                        }
                                    case 127:
                                        v73 := GenS9{F0: (v1 | 1), F1: (fi53 % (((-(fi53))) | 1)), F2: -2147483648}
                                        _ = v73
                                    }
                                    v74 := v5[("a" + ("a" + (("hello" + ("a" + "hello")) + "hello")))]
                                    _ = v74
                                }
                                fmt.Println(v17)
                            }
                        }
                    }
                    var v75 uint16 = 419
                    v3 = int64(fi26)
                    _ = fi53
                    _ = v75
                }
            }
            _ = fi26
        }
    }
    v5[""] = 4294967295
    if (143 != 190) {
        var v76 float32 = (1.5 - (1.5 + 3.14159))
        _ = v76
    }
    for fi77 := 0; fi77 < 11; fi77++ {
        if (185 == 2) {
            var v78 bool = v6
            v5[""] = 408
            fmt.Println(fi77)
            _ = v78
        }
        if v6 {
            switch "" {
            case "hello":
                var v79 uint16 = 1
                _ = v79
            case "":
                v6 = ((3.14159 + 0.1) != 3.14159)
            case "foo bar":
                var v80 uint8 = 145
                _ = v80
            }
        }
        _ = fi77
    }
    var v81 uint16 = 0
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if v6 {
        for fi82 := 0; fi82 < 7; fi82++ {
            var v83 float64 = 1.5
            var v84 float64 = (v83 + (-2.25 * v83))
            _ = fi82
            _ = v83
            _ = v84
        }
        v2 = append(v2, 2147483647)
        v1 = 4294967295
    }
    for fi85 := 0; fi85 < 13; fi85++ {
        v17 = -9223372036854775808
        if true {
            for fi86 := 0; fi86 < 15; fi86++ {
                var v87 uint16 = v81
                var v88 int32 = -1
                _ = fi86
                _ = v87
                _ = v88
            }
            v89 := GenS4{F0: ("foo bar" + ("\n\t" + ""))}
            _ = v89
        }
        _ = fi85
    }
    switch (^(v4)) {
    case 387:
        v2 = append(v2, 206)
    case -1:
        var v90 int = (v4 % (((v4 % (((v4 / (((v4 | v4)) | 1))) | 1))) | 1))
        _ = v90
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v81
}
