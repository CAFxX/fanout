package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8; F1 int32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS9 struct { F0 int32; F1 int64; F2 uint32 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int; F1 uint32 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 int; F1 uint16 }

func main() {
    var v1 uint16 = 712
    var v2 int16 = 42
    var v3 int32 = -582
    v4 := []int{61, -2147483648, 537, 2147483647}
    if (!((v1 % ((0) | 1)) > v1)) {
        for fi5 := 0; fi5 < 3; fi5++ {
            switch fi5 {
            case -9223372036854775808:
                _ = len(v4)
            case 0:
                fmt.Println(v2)
            case 127:
                for fi6 := 0; fi6 < 18; fi6++ {
                    var v7 uint16 = v1
                    _ = fi6
                    _ = v7
                }
            default:
                v8 := GenS1{F0: (v3 / ((0) | 1)), F1: int8(v1)}
                _ = v8
            }
            _ = fi5
        }
        v2 = (-(v2))
    } else {
        fmt.Println(v2, v3, v3)
        for fi9 := 0; fi9 < 3; fi9++ {
            v4 = append(v4, fi9)
            v3 -= (v3 % ((2147483647) | 1))
            _ = fi9
        }
        for fi10 := 0; fi10 < 3; fi10++ {
            v3 += (v3 / ((1) | 1))
            var v11 string = (("hello" + "hello") + "\n\t")
            _ = fi10
            _ = v11
        }
    }
    fmt.Println(v1)
    _ = len(v4)
    switch -9223372036854775808 {
    case 1:
        if true {
            for fi12 := 0; fi12 < 12; fi12++ {
                var v13 bool = true
                _ = fi12
                _ = v13
            }
            v2 = (v2 - v2)
        } else {
            for fi14 := 0; fi14 < 16; fi14++ {
                v2 += (v2 ^ 105)
                _ = fi14
            }
            var v15 int = -9223372036854775808
            var v16 float64 = (0.0 / 1e10)
            _ = v15
            _ = v16
        }
    case 127:
        for fi17 := 0; fi17 < 2; fi17++ {
            if (!true) {
                if true {
                    v2 -= v2
                }
                v18 := GenS2{F0: 9223372036854775807}
                v1 -= (v1 & 255)
                _ = v18
            }
            if false {
                if true {
                    var v19 int64 = 438
                    switch (-(fi17)) {
                    case 763:
                        fmt.Println(fi17, v1, v19)
                    case 0:
                        v20 := GenS3{F0: ("\n\t" + ((((("" + ("hello" + ((("a" + ("a" + ("" + "hello"))) + "a") + ""))) + "\n\t") + "hello") + "") + ((("" + "foo bar") + "") + "\n\t")))}
                        _ = v20
                    default:
                        switch fi17 {
                        case 804:
                            if false {
                                switch fi17 {
                                case 262:
                                    if len(v4) > 0 {
                                        v4[0] = (fi17 | fi17)
                                    }
                                case 17:
                                    for fi21 := 0; fi21 < 5; fi21++ {
                                        var v22 uint32 = 2147483648
                                        var v23 int16 = v2
                                        _ = fi21
                                        _ = v22
                                        _ = v23
                                    }
                                case 128:
                                    if true {
                                        var v24 bool = false
                                        _ = v24
                                    } else {
                                        v2 -= int16(fi17)
                                        switch "" {
                                        case "a":
                                            var v25 uint64 = 2
                                            _ = v25
                                        case "\n\t":
                                            v2 -= 1
                                        }
                                        for fi26 := 0; fi26 < 2; fi26++ {
                                            if (true && false) {
                                                v27 := GenS4{F0: 2147483648, F1: (^(v19))}
                                                if false {
                                                    if (("hello" < "a") || (true && (-9223372036854775808 <= fi26))) {
                                                        v28 := GenS5{F0: (v2 << ((686) & 63))}
                                                        _ = v28
                                                    } else {
                                                        fmt.Println(fi17, v1, fi26)
                                                        if (569 <= 65535) {
                                                            v29 := GenS6{F0: 127, F1: -2147483648}
                                                            _ = v29
                                                        }
                                                        v30 := GenS7{F0: (true || (1.5 == (1.5 * 3.14159)))}
                                                        _ = v30
                                                    }
                                                    if ((v19 / (((^(v19))) | 1)) <= -9223372036854775808) {
                                                        v31 := GenS6{F0: -9, F1: 0}
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        if true {
                                                            var v32 string = "\n\t"
                                                            var v33 uint16 = v1
                                                            _ = v32
                                                            _ = v33
                                                        } else {
                                                            var v34 int = (-(fi26))
                                                            _ = v34
                                                        }
                                                        var v35 int = fi26
                                                        switch "" {
                                                        case "hello":
                                                            var v36 int8 = 127
                                                            _ = v36
                                                        default:
                                                            v4 = append(v4, fi17)
                                                        }
                                                        _ = v31
                                                        _ = v35
                                                    }
                                                    fmt.Println(v3, v2)
                                                }
                                                if ((1 >= (fi17 >> ((0) & 63))) && true) {
                                                    switch "hello" {
                                                    case "hello":
                                                        var v37 uint32 = 77
                                                        _ = v37
                                                    case "":
                                                        if (2 > 127) {
                                                            if true {
                                                                var v38 string = "\n\t"
                                                                v4 = append(v4, (fi26 - (^(fi26))))
                                                                _ = v38
                                                            } else {
                                                                var v39 int64 = v19
                                                                _ = v39
                                                            }
                                                            var v40 bool = false
                                                            fmt.Println(v2, v40, v3)
                                                            v41 := GenS7{F0: true}
                                                            _ = v40
                                                            _ = v41
                                                        }
                                                    default:
                                                        if (!(!true)) {
                                                            fmt.Println(fi17, v19)
                                                        }
                                                    }
                                                    fmt.Println(v3, fi17, v3)
                                                    var v42 uint8 = 42
                                                    _ = v42
                                                } else {
                                                    if false {
                                                        switch fi26 {
                                                        case -9223372036854775808:
                                                            var v43 int64 = int64(fi17)
                                                            _ = v43
                                                        case 4294967295:
                                                            v44 := GenS6{F0: 42, F1: int32(v2)}
                                                            _ = v44
                                                        case -2147483648:
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                        switch (("foo bar" + ("" + ("a" + "foo bar"))) + ("a" + "hello")) {
                                                        case "\n\t":
                                                            v1 -= v1
                                                        default:
                                                            v19 -= (v19 % ((v19) | 1))
                                                        }
                                                        v3 += 1
                                                    }
                                                    v19 = v19
                                                }
                                                switch fi17 {
                                                case 9223372036854775807:
                                                    var v45 uint64 = 18446744073709551615
                                                    _ = v45
                                                case 32767:
                                                    fmt.Println(v1)
                                                case 1:
                                                    v46 := GenS8{F0: v2, F1: 65535, F2: 0}
                                                    _ = v46
                                                }
                                                v2 -= 32767
                                                v47 := GenS9{F0: 2147483647, F1: (v19 & 514), F2: 14}
                                                _ = v27
                                                _ = v47
                                            }
                                            var v48 int64 = v19
                                            var v49 uint16 = v1
                                            _ = fi26
                                            _ = v48
                                            _ = v49
                                        }
                                    }
                                default:
                                    v50 := GenS10{F0: (v3 - (v3 * 494))}
                                    _ = v50
                                }
                                if (!true) {
                                    var v51 float32 = float32((float64(float32(1e10)) * (((-2.25 * (0.1 / 3.14159)) + -2.25) / float64((0.1 + float32(float64(1e10)))))))
                                    for fi52 := 0; fi52 < 12; fi52++ {
                                        var v53 uint64 = 38
                                        _ = len(v4)
                                        var v54 float64 = -2.25
                                        var v55 uint32 = 187
                                        _ = fi52
                                        _ = v53
                                        _ = v54
                                        _ = v55
                                    }
                                    for fi56 := 0; fi56 < 4; fi56++ {
                                        fmt.Println(v1)
                                        fmt.Println(v19)
                                        if ((false || (v2 < (v2 / ((1000) | 1)))) || (true || (v51 >= (((float32(1.5) * (0.0 + 0.1)) - ((float32(3.14159) + 1e10) / v51)) - 3.14159)))) {
                                            v57 := GenS4{F0: 2147483648, F1: v19}
                                            v19 -= -1
                                            _ = v57
                                        } else {
                                            v58 := GenS11{F0: 32767, F1: 270}
                                            switch (fi56 | 2147483647) {
                                            case -459:
                                                fmt.Println(v1, v1, v2)
                                            }
                                            _ = len(v4)
                                            _ = v58
                                        }
                                        v19 += (v19 * -1)
                                        v51 -= 1.5
                                        if true {
                                            var v59 int32 = -65
                                            _ = v59
                                        } else {
                                            var v60 uint8 = 118
                                            v4 = append(v4, fi17)
                                            switch ("hello" + "\n\t") {
                                            case "hello":
                                                var v61 uint64 = 1
                                                _ = v61
                                            case "":
                                                v1 -= (v1 / ((v1) | 1))
                                            case "\n\t":
                                                _ = len(v4)
                                            }
                                            _ = v60
                                        }
                                        v1 = (v1 * (v1 % (((v1 & (^(v1)))) | 1)))
                                        var v62 uint16 = (v1 + v1)
                                        _ = fi56
                                        _ = v62
                                    }
                                    _ = v51
                                }
                                v19 += 0
                            }
                        case 42:
                            fmt.Println(v1, v3, v19)
                        default:
                            if (144 == 0) {
                                switch -2147483648 {
                                case 0:
                                    switch 0 {
                                    case 588:
                                        v3 += (v3 & (v3 / ((v3) | 1)))
                                    case 42:
                                        if (false && true) {
                                            switch fi17 {
                                            case 32767:
                                                v3 = (^(v3))
                                            case 42:
                                                if true {
                                                    switch ("" + ("a" + "")) {
                                                    case "\n\t":
                                                        fmt.Println(v2, v1)
                                                    case "a":
                                                        v63 := GenS3{F0: (("\n\t" + "foo bar") + "foo bar")}
                                                        _ = v63
                                                    case "foo bar":
                                                        for fi64 := 0; fi64 < 17; fi64++ {
                                                            _ = len(v4)
                                                            var v65 float64 = (1e10 * (-2.25 + (0.1 / 0.1)))
                                                            fmt.Println(v2, v19)
                                                            _ = fi64
                                                            _ = v65
                                                        }
                                                    default:
                                                        v19 = v19
                                                    }
                                                    if (42 < 18446744073709551615) {
                                                        v66 := GenS9{F0: -40, F1: (v19 + (v19 & (^(v19)))), F2: 65535}
                                                        _ = v66
                                                    } else {
                                                        if false {
                                                            var v67 float32 = 0.1
                                                            v67 = v67
                                                            _ = v67
                                                        } else {
                                                            var v68 int16 = (v2 | (v2 % ((v2) | 1)))
                                                            _ = v68
                                                        }
                                                        switch fi17 {
                                                        case 32767:
                                                            _ = len(v4)
                                                        }
                                                    }
                                                } else {
                                                    v1 = v1
                                                }
                                            }
                                            switch -9223372036854775808 {
                                            case 32768:
                                                _ = len(v4)
                                            case 9223372036854775807:
                                                if true {
                                                    v3 += (v3 & 786)
                                                    v69 := GenS10{F0: 95}
                                                    _ = v69
                                                } else {
                                                    fmt.Println(v1)
                                                    if true {
                                                        v2 -= v2
                                                    }
                                                    if false {
                                                        fmt.Println(v2)
                                                        v70 := GenS5{F0: (v2 - 42)}
                                                        v71 := GenS12{F0: (!false)}
                                                        _ = v70
                                                        _ = v71
                                                    }
                                                }
                                            case 32767:
                                                switch ("" + "hello") {
                                                case "":
                                                    if len(v4) > 0 {
                                                        v4[0] = fi17
                                                    }
                                                }
                                            default:
                                                if false {
                                                    var v72 int8 = 127
                                                    _ = v72
                                                } else {
                                                    v3 = v3
                                                }
                                            }
                                            v2 += -2
                                        }
                                    default:
                                        var v73 uint64 = 336
                                        _ = v73
                                    }
                                case -9223372036854775808:
                                    v74 := GenS12{F0: false}
                                    _ = v74
                                }
                                var v75 string = "foo bar"
                                if (669 != 2147483647) {
                                    v76 := GenS8{F0: (v2 * v2), F1: (^(v1)), F2: 197}
                                    switch (fi17 % (((fi17 / ((fi17) | 1))) | 1)) {
                                    case 128:
                                        switch fi17 {
                                        case 438:
                                            for fi77 := 0; fi77 < 6; fi77++ {
                                                _ = len(v4)
                                                _ = fi77
                                            }
                                        default:
                                            var v78 uint16 = 32767
                                            _ = v78
                                        }
                                    default:
                                        v19 -= v19
                                    }
                                    _ = v76
                                }
                                _ = v75
                            }
                        }
                    }
                    _ = v19
                } else {
                    var v79 uint = 0
                    _ = v79
                }
                if len(v4) > 0 {
                    v4[0] = (-(fi17))
                }
            }
            switch (("foo bar" + (("\n\t" + "\n\t") + "a")) + "foo bar") {
            case "\n\t":
                if len(v4) > 0 {
                    v4[0] = (fi17 << ((634) & 63))
                }
            case "hello":
                v80 := GenS13{F0: 2}
                _ = v80
            case "a":
                for fi81 := 0; fi81 < 17; fi81++ {
                    v1 += v1
                    if true {
                        v3 -= -824
                        v2 = 32767
                        v1 = v1
                    }
                    if true {
                        if (!true) {
                            v1 = v1
                            v2 += v2
                            v82 := GenS4{F0: 42, F1: -590}
                            _ = v82
                        }
                        var v83 int32 = -2147483648
                        _ = v83
                    }
                    _ = fi81
                }
            default:
                v1 -= v1
            }
            _ = fi17
        }
    }
    v2 -= -1
    for fi84 := 0; fi84 < 20; fi84++ {
        v2 = v2
        v4 = append(v4, fi84)
        _ = fi84
    }
    fmt.Println(v1, v3, v2)
    v85 := GenS5{F0: (v2 ^ v2)}
    v1 = v1
    v86 := GenS14{F0: -9223372036854775808, F1: 32767}
    var v87 int16 = 32767
    fmt.Println(v87)
    var v88 int64 = 388
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v85
    _ = v86
    _ = v87
    _ = v88
}
