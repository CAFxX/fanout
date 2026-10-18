package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int; F1 int64; F2 uint }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 int16; F1 uint16; F2 int16 }
type GenS7 struct { F0 int16; F1 uint16 }
type GenS8 struct { F0 uint16; F1 uint64 }
type GenS9 struct { F0 int16; F1 int32 }
func genFn0(p0 uint8, p1 uint8, p2 uint8) uint32 {
    for fi4 := 0; fi4 < 20; fi4++ {
        if (!(false || true)) {
            if (!(!((fi4 << ((255) & 63)) > (fi4 & -688)))) {
                v5 := []int{}
                var v6 uint32 = 42
                p1 += (^(p0))
                _ = v5
                _ = v6
            }
            v7 := GenS1{F0: "foo bar"}
            _ = v7
        }
        p0 = p2
        _ = fi4
    }
    v8 := []int{}
    p0 = p0
    if (true && (18446744073709551615 == 373)) {
        for fi9 := 0; fi9 < 20; fi9++ {
            p0 = (p2 | 127)
            for fi10 := 0; fi10 < 3; fi10++ {
                p1 -= p0
                switch (fi10 - (fi10 % ((fi10) | 1))) {
                case 9223372036854775807:
                    p2 = p1
                case -2147483648:
                    switch (((("hello" + "foo bar") + "foo bar") + "a") + "") {
                    case "":
                        p1 = (p1 & p2)
                    case "hello":
                        fmt.Println(fi10, fi10, p0)
                    default:
                        if (4294967295 < 65535) {
                            var v11 uint64 = 18446744073709551615
                            _ = v11
                        } else {
                            if (false || false) {
                                v8 = append(v8, 32768)
                                switch -442 {
                                case 2147483647:
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                case 9223372036854775807:
                                    if true {
                                        if len(v8) > 0 {
                                            v8[0] = fi10
                                        }
                                        p2 += p2
                                    }
                                case 127:
                                    p1 += (p2 * 1)
                                    switch ("foo bar" + "a") {
                                    case "\n\t":
                                        if (-1 == 32767) {
                                            var v12 int64 = -1
                                            var v13 uint32 = 65535
                                            var v14 string = ""
                                            _ = v12
                                            _ = v13
                                            _ = v14
                                        } else {
                                            fmt.Println(fi10)
                                            var v15 int8 = 7
                                            if (448 != 2) {
                                                var v16 bool = (!false)
                                                if false {
                                                    v15 = 42
                                                    p2 = 2
                                                    var v17 uint16 = 13
                                                    _ = v17
                                                }
                                                _ = v16
                                            }
                                            _ = v15
                                        }
                                    case "":
                                        switch ("a" + "") {
                                        case "hello":
                                            if (((1.5 - ((-2.25 - -2.25) + (3.14159 / 1.5))) + 0.1) >= (float64(((1.5 / (-2.25 / (1.5 / (float32(-2.25) * float32(-2.25))))) + 1.5)) / -2.25)) {
                                                var v18 float32 = 0.1
                                                if true {
                                                    v19 := GenS1{F0: (("a" + "\n\t") + "hello")}
                                                    var v20 int64 = 483
                                                    _ = v19
                                                    _ = v20
                                                }
                                                _ = v18
                                            }
                                        default:
                                            _ = len(v8)
                                        }
                                    case "a":
                                        p2 += 0
                                    }
                                    var v21 uint64 = 1
                                    _ = v21
                                default:
                                    var v22 uint8 = (p0 ^ (p1 - (p1 & 255)))
                                    _ = v22
                                }
                            } else {
                                switch 4294967295 {
                                case -458:
                                    p2 = p2
                                default:
                                    if (("foo bar" < ((("hello" + ((("foo bar" + "hello") + (("hello" + (("hello" + "\n\t") + "hello")) + "")) + ("hello" + ""))) + "\n\t") + "hello")) || true) {
                                        p2 += (p1 >> ((505) & 63))
                                        var v23 uint16 = 255
                                        _ = v23
                                    }
                                }
                                var v24 int32 = 2147483647
                                _ = v24
                            }
                            fmt.Println(p0, fi10)
                            v25 := GenS2{F0: 660}
                            var v26 int8 = 84
                            if len(v8) > 0 {
                                v8[0] = 2147483647
                            }
                            _ = v25
                            _ = v26
                        }
                    }
                    fmt.Println(p1)
                    fmt.Println(fi9, p2, p0)
                case -111:
                    var v27 uint32 = 2
                    _ = v27
                default:
                    var v28 bool = (!true)
                    _ = v28
                }
                p2 = 127
                _ = fi10
            }
            _ = fi9
        }
        var v29 int64 = -9223372036854775808
        _ = v29
    } else {
        v30 := GenS3{F0: 81, F1: 967}
        _ = v30
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 2
}

func main() {
    v1 := []uint32{1, 940, 35}
    v2 := []uint32{374, 1, 0, 2}
    var v3 int8 = -1
    v31 := GenS1{F0: ("a" + "")}
    _ = len(v1)
    for fi32 := 0; fi32 < 20; fi32++ {
        var v33 uint8 = 242
        fmt.Println(v33)
        v34 := GenS4{F0: 32768, F1: 9223372036854775807, F2: 42}
        _ = fi32
        _ = v33
        _ = v34
    }
    v1 = append(v1, 42)
    for fi35 := 0; fi35 < 7; fi35++ {
        switch "" {
        case "a":
            if true {
                if false {
                    v3 = (v3 / (((v3 >> ((65535) & 63))) | 1))
                    var v36 uint64 = 887
                    v37 := GenS4{F0: fi35, F1: 1, F2: 2}
                    _ = v36
                    _ = v37
                }
                v3 -= v3
            } else {
                switch int(v3) {
                case 9223372036854775807:
                    if (v3 > v3) {
                        v3 -= -68
                        if true {
                            var v38 int16 = 42
                            _ = v38
                        }
                    } else {
                        for fi39 := 0; fi39 < 5; fi39++ {
                            fmt.Println(fi39, v3)
                            _ = fi39
                        }
                        var v40 int32 = 2147483647
                        v3 -= (v3 / (((v3 << ((2) & 63))) | 1))
                        _ = v40
                    }
                case 128:
                    if (-1 == 2147483647) {
                        var v41 uint8 = 128
                        _ = v41
                    }
                case -1:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                if (!true) {
                    switch ("foo bar" + ("" + "foo bar")) {
                    case "a":
                        if (0 >= 1000) {
                            var v42 int64 = 529
                            var v43 uint32 = 2147483648
                            _ = v42
                            _ = v43
                        } else {
                            v2 = append(v2, 161)
                        }
                    case "\n\t":
                        for fi44 := 0; fi44 < 13; fi44++ {
                            v3 = v3
                            _ = fi44
                        }
                    default:
                        fmt.Println(fi35, fi35)
                    }
                    switch (fi35 - (fi35 ^ 9223372036854775807)) {
                    case 1:
                        for fi45 := 0; fi45 < 20; fi45++ {
                            v3 -= (v3 >> ((0) & 63))
                            _ = fi45
                        }
                    case 128:
                        fmt.Println(fi35, fi35)
                    case 42:
                        var v46 bool = (26 > 255)
                        _ = v46
                    default:
                        var v47 int64 = 761
                        _ = v47
                    }
                    var v48 bool = false
                    _ = v48
                }
                _ = len(v2)
            }
        }
        _ = fi35
    }
    for fi49 := 0; fi49 < 13; fi49++ {
        v3 = v3
        _ = fi49
    }
    _ = len(v1)
    var v50 int8 = v3
    var v51 uint16 = 42
    var v52 int32 = 2147483647
    var v53 uint32 = 121
    v54 := GenS5{F0: 0, F1: ((true || (v53 < uint32(v51))) || (!(0 == 32768)))}
    if true {
        v3 += (v3 | int8(v53))
        fmt.Println(v53)
        if false {
            var v55 float32 = (1.5 * (3.14159 + 1.5))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v55
        } else {
            fmt.Println(v52, v53)
        }
    } else {
        v52 = (v52 >> ((255) & 63))
        if true {
            var v56 int32 = (v52 & (-(v52)))
            if (758 == 4294967295) {
                for fi57 := 0; fi57 < 17; fi57++ {
                    switch "a" {
                    case "":
                        var v58 uint64 = 0
                        _ = v58
                    case "\n\t":
                        if ((v53 << ((18446744073709551615) & 63)) < (v53 % (((^(v53))) | 1))) {
                            for fi59 := 0; fi59 < 14; fi59++ {
                                if len(v1) > 0 {
                                    v1[0] = v53
                                }
                                switch fi57 {
                                case 9223372036854775807:
                                    fmt.Println(v52, v3)
                                case 127:
                                    v51 -= (v51 - (v51 & (^(v51))))
                                case 2147483647:
                                    v50 = (v50 - -81)
                                default:
                                    v52 = v52
                                }
                                var v60 int16 = 1000
                                if len(v2) > 0 {
                                    v2[0] = genFn0(234, 127, 234)
                                }
                                _ = fi59
                                _ = v60
                            }
                            if (42 != 42) {
                                var v61 int = (^(fi57))
                                switch (v61 / ((fi57) | 1)) {
                                case 0:
                                    var v62 uint16 = (v51 / (((v51 ^ v51)) | 1))
                                    _ = v62
                                case -2147483648:
                                    var v63 int64 = -718
                                    _ = v63
                                case -565:
                                    fmt.Println(v53, v53, v61)
                                default:
                                    if false {
                                        switch ("a" + (("" + (("a" + ((("a" + "a") + (("" + "") + ("a" + ("" + "")))) + (("a" + "hello") + "\n\t"))) + "a")) + "hello")) {
                                        case "foo bar":
                                            for fi64 := 0; fi64 < 10; fi64++ {
                                                if false {
                                                    var v65 uint64 = 557
                                                    _ = v65
                                                } else {
                                                    var v66 uint = 309
                                                    var v67 string = ""
                                                    switch v67 {
                                                    case "\n\t":
                                                        fmt.Println(v67, v61)
                                                    case "foo bar":
                                                        fmt.Println(v53, fi64, v56)
                                                    case "a":
                                                        switch 32768 {
                                                        case 2147483647:
                                                            v68 := GenS6{F0: 42, F1: 205, F2: 1000}
                                                            _ = v68
                                                        case 32768:
                                                            v69 := GenS2{F0: 32767}
                                                            _ = v69
                                                        default:
                                                            v70 := GenS7{F0: 0, F1: (^(v51))}
                                                            _ = v70
                                                        }
                                                    }
                                                    _ = v66
                                                    _ = v67
                                                }
                                                _ = fi64
                                            }
                                        case "":
                                            switch "" {
                                            case "":
                                                _ = len(v1)
                                            case "foo bar":
                                                v3 -= (^(v50))
                                            default:
                                                var v71 uint16 = 65535
                                                _ = v71
                                            }
                                        case "\n\t":
                                            var v72 float32 = -2.25
                                            _ = v72
                                        default:
                                            switch (v61 * -798) {
                                            case 1:
                                                var v73 bool = (304 < -32768)
                                                _ = v73
                                            case -2147483648:
                                                if false {
                                                    var v74 uint8 = 0
                                                    _ = v74
                                                }
                                            case 890:
                                                switch fi57 {
                                                case -1:
                                                    fmt.Println(v56, v50)
                                                case 4294967295:
                                                    fmt.Println(v51, v50)
                                                }
                                            }
                                        }
                                        switch (v61 | (fi57 & v61)) {
                                        case 128:
                                            v1 = append(v1, (^(v53)))
                                        default:
                                            var v75 int16 = 1
                                            _ = v75
                                        }
                                        var v76 string = "foo bar"
                                        _ = v76
                                    }
                                }
                                _ = v61
                            }
                            for fi77 := 0; fi77 < 9; fi77++ {
                                if false {
                                    v1 = append(v1, (v53 / ((genFn0(0, 127, 0)) | 1)))
                                    if (42 < 72) {
                                        fmt.Println(v51, fi77, fi77)
                                        v3 -= v50
                                        if len(v2) > 0 {
                                            v2[0] = (v53 % (((v53 % (((v53 << ((0) & 63))) | 1))) | 1))
                                        }
                                    }
                                }
                                _ = fi77
                            }
                        } else {
                            if ((v51 >= (v51 ^ (^(v51)))) && ((v50 * (^(v50))) < 0)) {
                                if ((^(v53)) <= v53) {
                                    v78 := GenS8{F0: v51, F1: 1}
                                    v53 -= genFn0(2, 128, 2)
                                    _ = v78
                                }
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v3 = (v50 % ((25) | 1))
                                v3 += 42
                            }
                            v50 -= (v50 & 118)
                        }
                    default:
                        v79 := GenS9{F0: 1, F1: (-(v56))}
                        _ = v79
                    }
                    fmt.Println(v52)
                    var v80 uint = 0
                    _ = fi57
                    _ = v80
                }
                v2 = append(v2, v53)
            }
            if false {
                _ = len(v2)
                v3 += -4
            }
            _ = v56
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
}
