package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint64; F1 bool; F2 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint32; F1 string; F2 int32 }
type GenS8 struct { F0 uint64; F1 uint32 }
type GenS9 struct { F0 bool; F1 string }

func main() {
    v1 := []string{("" + "foo bar"), ("" + "a")}
    v2 := make(map[string]string)
    var v3 float32 = (1.5 + 1e10)
    var v4 uint8 = 42
    var v5 int8 = -114
    var v6 uint16 = 32767
    v5 = (v5 - v5)
    var v7 int8 = (v5 ^ (v5 + -110))
    v8 := GenS1{F0: 4294967295, F1: v4}
    v4 += (^(v4))
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    var v9 float64 = float64(v3)
    var v10 int16 = -704
    switch -742 {
    case 4294967295:
        switch 42 {
        case 1:
            switch -340 {
            case 549:
                delete(v2, ("foo bar" + ""))
            case 4294967295:
                switch 127 {
                case 128:
                    if false {
                        switch "foo bar" {
                        case "hello":
                            for fi11 := 0; fi11 < 12; fi11++ {
                                for fi12 := 0; fi12 < 9; fi12++ {
                                    var v13 uint16 = v6
                                    _ = fi12
                                    _ = v13
                                }
                                _ = fi11
                            }
                        case "\n\t":
                            var v14 uint16 = (v6 + v6)
                            _ = v14
                        }
                        for fi15 := 0; fi15 < 14; fi15++ {
                            for fi16 := 0; fi16 < 14; fi16++ {
                                fmt.Println(v3, v3, v6)
                                fmt.Println(v3, v10)
                                _ = fi16
                            }
                            _ = fi15
                        }
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                    } else {
                        for fi17 := 0; fi17 < 7; fi17++ {
                            for fi18 := 0; fi18 < 8; fi18++ {
                                if ("hello" != "a") {
                                    if (false && true) {
                                        fmt.Println(v7, v9)
                                    }
                                    switch (fi17 ^ fi18) {
                                    case 624:
                                        v3 += 1.5
                                    case 0:
                                        if (!true) {
                                            if len(v1) > 0 {
                                                v1[0] = "\n\t"
                                            }
                                            if false {
                                                if len(v1) > 0 {
                                                    v1[0] = "\n\t"
                                                }
                                                v19 := GenS1{F0: 1, F1: v4}
                                                _ = v19
                                            }
                                            switch int(v6) {
                                            case 32767:
                                                v20 := GenS2{F0: 840, F1: (!(42 == -9223372036854775808)), F2: v10}
                                                _ = v20
                                            default:
                                                var v21 uint16 = v6
                                                _ = v21
                                            }
                                        } else {
                                            v9 += 0.1
                                            v5 -= v5
                                        }
                                    default:
                                        v9 = (0.0 * v9)
                                    }
                                }
                                if ((!true) && true) {
                                    fmt.Println(fi18, v9, fi17)
                                    v2["\n\t"] = "foo bar"
                                    var v22 string = ""
                                    _ = v22
                                }
                                v9 = (1e10 / (((0.1 / (v9 - (v9 / v9))) / ((v9 - v9) + (v9 / -2.25))) + v9))
                                _ = fi18
                            }
                            v10 += 42
                            _ = fi17
                        }
                        v6 += (v6 - v6)
                    }
                default:
                    var v23 uint = 255
                    _ = v23
                }
            case 9223372036854775807:
                fmt.Println(v7, v5)
            default:
                for fi24 := 0; fi24 < 12; fi24++ {
                    if (!(false || false)) {
                        if true {
                            for fi25 := 0; fi25 < 3; fi25++ {
                                v26 := GenS2{F0: 364, F1: (!(!((int16(fi25) <= (v10 + -1)) || (v5 < 0)))), F2: (v10 & (v10 - v10))}
                                _ = fi25
                                _ = v26
                            }
                        } else {
                            v3 += (float32(float64((v3 / v3))) + -2.25)
                            v27 := v2[("foo bar" + ("hello" + "a"))]
                            _ = v27
                        }
                        fmt.Println(v4)
                        _ = len(v1)
                    } else {
                        if true {
                            switch "hello" {
                            case "hello":
                                v5 -= v5
                            case "\n\t":
                                v4 = v4
                            default:
                                switch ("\n\t" + "") {
                                case "":
                                    fmt.Println(v5, v4, v5)
                                default:
                                    v7 = (^(v5))
                                }
                            }
                            v28 := v2[((("hello" + "foo bar") + ("" + (("\n\t" + "hello") + ("hello" + ("hello" + ("\n\t" + ("" + ""))))))) + "\n\t")]
                            v29 := GenS1{F0: 0, F1: v4}
                            _ = v28
                            _ = v29
                        }
                    }
                    if (!(fi24 == (fi24 / (((fi24 - fi24)) | 1)))) {
                        if true {
                            v4 = v4
                        } else {
                            switch "" {
                            case "\n\t":
                                v10 = v10
                            case "foo bar":
                                if true {
                                    switch (((("hello" + ("" + "a")) + "a") + ("" + ("foo bar" + ((("a" + "") + "") + (("a" + "\n\t") + ""))))) + "\n\t") {
                                    case "foo bar":
                                        _ = len(v1)
                                    default:
                                        if ((v6 % ((0) | 1)) > (v6 >> ((4294967295) & 63))) {
                                            if (!(true && false)) {
                                                if len(v1) > 0 {
                                                    v1[0] = (("a" + ("foo bar" + "\n\t")) + "")
                                                }
                                                if (true || (!true)) {
                                                    for fi30 := 0; fi30 < 5; fi30++ {
                                                        if (v6 != (v6 * (^(v6)))) {
                                                            var v31 float32 = 1.5
                                                            fmt.Println(fi30)
                                                            _ = len(v1)
                                                            _ = v31
                                                        } else {
                                                            if true {
                                                                switch 984 {
                                                                case 9223372036854775807:
                                                                    fmt.Println(v4)
                                                                case 487:
                                                                    v32 := GenS1{F0: 1, F1: (v4 * 42)}
                                                                    _ = v32
                                                                }
                                                                if false {
                                                                    var v33 float64 = v9
                                                                    var v34 uint8 = (v4 * v4)
                                                                    if ((v4 * (v34 - (v34 - v34))) <= (v34 % (((^(v4))) | 1))) {
                                                                        v34 = (v34 + v34)
                                                                        v4 = 41
                                                                    }
                                                                    _ = v33
                                                                    _ = v34
                                                                } else {
                                                                    if (!(127 >= 0)) {
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                    }
                                                                    if true {
                                                                        switch fi30 {
                                                                        case 127:
                                                                            var v35 int64 = -9223372036854775808
                                                                            _ = v35
                                                                        }
                                                                        v6 = 1
                                                                        v36 := v2[("a" + "")]
                                                                        _ = v36
                                                                    } else {
                                                                        var v37 uint = 692
                                                                        var v38 int32 = 42
                                                                        v39 := GenS3{F0: 42}
                                                                        _ = v37
                                                                        _ = v38
                                                                        _ = v39
                                                                    }
                                                                }
                                                                fmt.Println(v6, v3, v10)
                                                            }
                                                            fmt.Println(v5)
                                                            var v40 int = fi30
                                                            _ = v40
                                                        }
                                                        var v41 uint64 = 42
                                                        _ = fi30
                                                        _ = v41
                                                    }
                                                    if false {
                                                        v9 += v9
                                                    } else {
                                                        if len(v1) > 0 {
                                                            v1[0] = ""
                                                        }
                                                        if len(v1) > 0 {
                                                            v1[0] = "\n\t"
                                                        }
                                                        v10 += 48
                                                    }
                                                    v2[""] = ""
                                                } else {
                                                    if ((v5 + v5) < (v5 * v7)) {
                                                        var v42 string = (("foo bar" + "foo bar") + "a")
                                                        _ = v42
                                                    } else {
                                                        var v43 float32 = v3
                                                        v2[""] = "a"
                                                        _ = v43
                                                    }
                                                    for fi44 := 0; fi44 < 10; fi44++ {
                                                        v45 := GenS2{F0: 948, F1: false, F2: 42}
                                                        v6 = v6
                                                        v46 := GenS2{F0: 1, F1: true, F2: (v10 - 1)}
                                                        _ = fi44
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                    v7 = (v5 ^ v5)
                                                }
                                            }
                                            v7 = (-(v5))
                                            v47 := GenS1{F0: 18446744073709551615, F1: v4}
                                            _ = v47
                                        }
                                    }
                                    v48 := GenS3{F0: -9223372036854775808}
                                    _ = v48
                                } else {
                                    v49 := GenS2{F0: 609, F1: true, F2: v10}
                                    for fi50 := 0; fi50 < 8; fi50++ {
                                        v51 := GenS3{F0: 710}
                                        var v52 uint16 = v6
                                        v53 := GenS4{F0: (!true), F1: uint32(v4)}
                                        _ = fi50
                                        _ = v51
                                        _ = v52
                                        _ = v53
                                    }
                                    _ = v49
                                }
                            case "":
                                if (false && (240 != 42)) {
                                    for fi54 := 0; fi54 < 19; fi54++ {
                                        var v55 string = "\n\t"
                                        if ((0 > 950) || true) {
                                            v55 = v55
                                            fmt.Println(v7, v4, v55)
                                            v2[v55] = (v55 + ("foo bar" + ((("" + (v55 + v55)) + ("\n\t" + ("hello" + v55))) + v55)))
                                        } else {
                                            fmt.Println(fi54, fi24, v6)
                                            if (1.5 <= 0.0) {
                                                v56 := GenS4{F0: false, F1: 514}
                                                _ = v56
                                            } else {
                                                switch (fi24 | (fi54 >> ((0) & 63))) {
                                                case 9223372036854775807:
                                                    var v57 float32 = (v3 * v3)
                                                    _ = v57
                                                case 127:
                                                    if (18446744073709551615 < 1) {
                                                        v58 := GenS5{F0: 566}
                                                        _ = v58
                                                    }
                                                case -9223372036854775808:
                                                    var v59 int8 = v5
                                                    v1 = append(v1, (("" + ("foo bar" + "foo bar")) + v55))
                                                    _ = v59
                                                default:
                                                    fmt.Println(fi24, v7)
                                                }
                                                switch v55 {
                                                case "hello":
                                                    if (!(fi54 == (fi24 + (fi54 << ((339) & 63))))) {
                                                        if (!("foo bar" == (((v55 + ("hello" + "")) + "\n\t") + "\n\t"))) {
                                                            v9 += v9
                                                        } else {
                                                            if false {
                                                                v2[v55] = (v55 + v55)
                                                                v60 := GenS6{F0: 255}
                                                                _ = v60
                                                            } else {
                                                                if len(v1) > 0 {
                                                                    v1[0] = v55
                                                                }
                                                                v3 -= 1.5
                                                                switch 1 {
                                                                case -506:
                                                                    v61 := GenS3{F0: 0}
                                                                    _ = v61
                                                                case 9223372036854775807:
                                                                    var v62 uint16 = (v6 >> ((255) & 63))
                                                                    _ = v62
                                                                default:
                                                                    fmt.Println(fi24)
                                                                }
                                                            }
                                                        }
                                                        switch -157 {
                                                        case 1:
                                                            if true {
                                                                v63 := v2[(("hello" + v55) + (v55 + "a"))]
                                                                if (v63 >= "") {
                                                                    var v64 int16 = (v10 | v10)
                                                                    v1 = append(v1, v63)
                                                                    _ = v64
                                                                }
                                                                _ = v63
                                                            } else {
                                                                v65 := GenS7{F0: 399, F1: (v55 + v55), F2: 0}
                                                                v66 := GenS5{F0: v6}
                                                                if (4294967295 > 1) {
                                                                    if ((v10 >= (-(v10))) && (42 < 2)) {
                                                                        var v67 int32 = 2147483647
                                                                        if true {
                                                                            v68 := GenS5{F0: (v6 - (v6 + (^(v6))))}
                                                                            _ = v68
                                                                        }
                                                                        fmt.Println(fi54, fi54, v55)
                                                                        _ = v67
                                                                    }
                                                                    v69 := GenS5{F0: v6}
                                                                    if ((v7 >= (v5 / ((v5) | 1))) || (((v5 | -1) <= (v7 >> ((65535) & 63))) && false)) {
                                                                        var v70 int64 = -641
                                                                        fmt.Println(v7, v70)
                                                                        delete(v2, "a")
                                                                        if len(v1) > 0 {
                                                                            v1[0] = v55
                                                                        }
                                                                        _ = v70
                                                                    } else {
                                                                        if (!true) {
                                                                            var v71 int = (fi24 & -2147483648)
                                                                            v4 = (v4 % ((255) | 1))
                                                                            _ = v71
                                                                        }
                                                                    }
                                                                    _ = v69
                                                                }
                                                                var v72 uint = 255
                                                                _ = v65
                                                                _ = v66
                                                                _ = v72
                                                            }
                                                        default:
                                                            var v73 uint32 = 4294967295
                                                            _ = v73
                                                        }
                                                        v5 -= v5
                                                    }
                                                    var v74 string = "a"
                                                    switch fi54 {
                                                    case 128:
                                                        switch (^(fi54)) {
                                                        case 42:
                                                            if ((!(v6 < (v6 << ((549) & 63)))) || true) {
                                                                _ = len(v1)
                                                                v6 = v6
                                                                v7 = int8(v6)
                                                            } else {
                                                                if true {
                                                                    if false {
                                                                        switch v55 {
                                                                        case "":
                                                                            var v75 float32 = ((v3 - (3.14159 * -2.25)) / (float32(float64((v3 * (v3 * 1e10)))) * v3))
                                                                            _ = v75
                                                                        }
                                                                    } else {
                                                                        v4 += v4
                                                                        if (!(543 > (fi54 * fi54))) {
                                                                            switch v74 {
                                                                            case "":
                                                                                v5 += v7
                                                                            case "hello":
                                                                                fmt.Println(v3, fi24, v74)
                                                                            }
                                                                        }
                                                                    }
                                                                    v7 = (v7 << ((2) & 63))
                                                                } else {
                                                                    if true {
                                                                        var v76 int = 42
                                                                        _ = v76
                                                                    } else {
                                                                        switch (fi24 >> ((255) & 63)) {
                                                                        case 32767:
                                                                            if false {
                                                                                v77 := GenS8{F0: 42, F1: 1}
                                                                                _ = len(v2)
                                                                                v74 = v74
                                                                                _ = v77
                                                                            }
                                                                        default:
                                                                            fmt.Println(v3, v4, fi54)
                                                                        }
                                                                        v10 += v10
                                                                    }
                                                                }
                                                                fmt.Println(v74)
                                                                var v78 string = v55
                                                                _ = v78
                                                            }
                                                        case 1:
                                                            v79 := GenS1{F0: 53, F1: (v4 / (((v4 ^ v4)) | 1))}
                                                            _ = v79
                                                        case 127:
                                                            var v80 uint = 1
                                                            _ = v80
                                                        }
                                                    case 32768:
                                                        fmt.Println(v5, v3, v4)
                                                    }
                                                    v74 = (v74 + (("foo bar" + v74) + v74))
                                                    var v81 string = ""
                                                    _ = v74
                                                    _ = v81
                                                case "foo bar":
                                                    fmt.Println(v7, v4)
                                                }
                                            }
                                        }
                                        _ = fi54
                                        _ = v55
                                    }
                                } else {
                                    v3 -= v3
                                    v82 := GenS9{F0: false, F1: "\n\t"}
                                    var v83 float64 = (-2.25 * v9)
                                    _ = v82
                                    _ = v83
                                }
                            default:
                                var v84 float32 = 3.14159
                                _ = v84
                            }
                        }
                    } else {
                        v2["foo bar"] = "a"
                        fmt.Println(v3, v3)
                        var v85 uint64 = 423
                        _ = v85
                    }
                    _ = fi24
                }
            }
        }
    case 682:
        fmt.Println(v5)
    }
    for fi86 := 0; fi86 < 20; fi86++ {
        v2[("foo bar" + (("\n\t" + "") + "a"))] = "\n\t"
        _ = fi86
    }
    _ = len(v2)
    switch -586 {
    case -1:
        var v87 uint = 4294967295
        _ = v87
    default:
        v88 := GenS4{F0: false, F1: 2}
        _ = v88
    }
    var v89 float32 = 1e10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v89:", v89)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v89
    _ = v9
}
