package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int; F1 int; F2 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS5 struct { F0 uint8; F1 int32; F2 string }
type GenS6 struct { F0 string }
type GenS7 struct { F0 bool; F1 string; F2 uint32 }
type GenS8 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint16; F1 int64 }

func main() {
    var v1 float32 = (1.5 + -2.25)
    var v2 int16 = -32768
    v3 := []string{(((("foo bar" + "foo bar") + (("foo bar" + (("foo bar" + "\n\t") + "hello")) + "foo bar")) + "a") + ("\n\t" + "a")), (("\n\t" + "a") + ("foo bar" + "\n\t"))}
    v4 := make(map[int]int)
    var v5 int32 = 42
    var v6 string = "hello"
    fmt.Println(v2)
    switch (v6 + "a") {
    case "":
        v6 = (("hello" + (v6 + v6)) + "hello")
    }
    v7 := GenS1{F0: -9223372036854775808, F1: v6}
    v2 = 1
    switch v6 {
    case "\n\t":
        v4[782] = 32768
    case "a":
        if (1 == 1) {
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 int = fi8
                _ = fi8
                _ = v9
            }
            if true {
                _ = len(v3)
            } else {
                if ((-(v2)) > v2) {
                    if true {
                        for fi10 := 0; fi10 < 6; fi10++ {
                            switch 743 {
                            case 577:
                                for fi11 := 0; fi11 < 3; fi11++ {
                                    var v12 string = v6
                                    v1 = v1
                                    _ = fi11
                                    _ = v12
                                }
                            case 159:
                                for fi13 := 0; fi13 < 18; fi13++ {
                                    var v14 bool = false
                                    v14 = ((v2 + (^(v2))) > (-(v2)))
                                    _ = fi13
                                    _ = v14
                                }
                            default:
                                var v15 uint32 = 196
                                _ = v15
                            }
                            _ = fi10
                        }
                    } else {
                        v5 += (v5 >> ((65535) & 63))
                        for fi16 := 0; fi16 < 4; fi16++ {
                            for fi17 := 0; fi17 < 14; fi17++ {
                                var v18 int32 = (v5 >> ((992) & 63))
                                switch v6 {
                                case "hello":
                                    if (!true) {
                                        switch (fi16 & (fi17 ^ (fi16 << ((42) & 63)))) {
                                        case 935:
                                            v19 := GenS1{F0: 232, F1: v6}
                                            _ = v19
                                        case 45:
                                            v20 := GenS1{F0: 9223372036854775807, F1: v6}
                                            _ = v20
                                        case 32768:
                                            v21 := GenS1{F0: 42, F1: "hello"}
                                            _ = v21
                                        default:
                                            v22 := GenS1{F0: 42, F1: ((v6 + "hello") + "a")}
                                            _ = v22
                                        }
                                    }
                                default:
                                    v2 -= 0
                                }
                                fmt.Println(fi16, v2)
                                _ = fi17
                                _ = v18
                            }
                            var v23 uint32 = 2147483648
                            v5 -= (v5 >> ((535) & 63))
                            _ = fi16
                            _ = v23
                        }
                    }
                } else {
                    if (32767 >= -836) {
                        if false {
                            var v24 bool = (((float64((0.0 * (v1 * (0.0 + v1)))) / 1.0) > (3.14159 / 1.0)) && (v2 < (v2 / (((v2 - 1)) | 1))))
                            var v25 float64 = 1e10
                            var v26 uint16 = 255
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        switch "foo bar" {
                        case "":
                            for fi27 := 0; fi27 < 18; fi27++ {
                                for fi28 := 0; fi28 < 5; fi28++ {
                                    v29 := GenS1{F0: 1, F1: v6}
                                    _ = fi28
                                    _ = v29
                                }
                                _ = fi27
                            }
                        case "foo bar":
                            switch v6 {
                            case "a":
                                v5 += 42
                            default:
                                v2 += -245
                            }
                        default:
                            switch (("" + v6) + v6) {
                            case "hello":
                                if true {
                                    switch 0 {
                                    case 506:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    case 2147483647:
                                        fmt.Println(v2, v2)
                                    }
                                }
                            case "a":
                                for fi30 := 0; fi30 < 1; fi30++ {
                                    if false {
                                        switch "foo bar" {
                                        case "\n\t":
                                            if (v2 < v2) {
                                                var v31 int32 = v5
                                                _ = v31
                                            }
                                        default:
                                            fmt.Println(v5, v5)
                                        }
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = v6
                                    }
                                    var v32 uint16 = 813
                                    _ = fi30
                                    _ = v32
                                }
                            }
                        }
                    } else {
                        v1 = 3.14159
                        if (!(true && (false && (true || false)))) {
                            for fi33 := 0; fi33 < 11; fi33++ {
                                switch v6 {
                                case "foo bar":
                                    if false {
                                        var v34 int16 = v2
                                        for fi35 := 0; fi35 < 17; fi35++ {
                                            var v36 uint32 = 2147483648
                                            _ = fi35
                                            _ = v36
                                        }
                                        var v37 uint16 = 879
                                        _ = v34
                                        _ = v37
                                    } else {
                                        var v38 int16 = -734
                                        _ = v38
                                    }
                                case "hello":
                                    if true {
                                        v5 = v5
                                    }
                                }
                                for fi39 := 0; fi39 < 5; fi39++ {
                                    var v40 uint16 = 252
                                    _ = fi39
                                    _ = v40
                                }
                                v1 += (3.14159 / v1)
                                _ = fi33
                            }
                            switch 42 {
                            case 32768:
                                var v41 int16 = v2
                                _ = v41
                            case 0:
                                v2 -= v2
                            case -545:
                                switch 127 {
                                case 32767:
                                    for fi42 := 0; fi42 < 17; fi42++ {
                                        if false {
                                            var v43 uint = 1
                                            var v44 uint32 = 0
                                            var v45 uint16 = 418
                                            _ = v43
                                            _ = v44
                                            _ = v45
                                        }
                                        var v46 bool = (true || (!true))
                                        fmt.Println(v46, v1, v2)
                                        _ = fi42
                                        _ = v46
                                    }
                                case 137:
                                    v47 := GenS2{F0: -9223372036854775808, F1: 2147483647, F2: (v2 | (v2 / ((v2) | 1)))}
                                    _ = v47
                                }
                            default:
                                fmt.Println(v2)
                            }
                            if ((((v6 + (v6 + v6)) + ("a" + v6)) < v6) && (940 >= 2)) {
                                switch v6 {
                                case "":
                                    var v48 uint = 1
                                    _ = v48
                                default:
                                    v6 = (("foo bar" + v6) + ("hello" + "foo bar"))
                                }
                            } else {
                                var v49 uint8 = 42
                                _ = v49
                            }
                        } else {
                            _ = len(v3)
                        }
                    }
                    fmt.Println(v2)
                    v5 -= 1
                }
            }
            for fi50 := 0; fi50 < 2; fi50++ {
                if false {
                    if ((v5 & (-(v5))) <= v5) {
                        if true {
                            if false {
                                v51 := GenS3{F0: v2}
                                var v52 float64 = (0.1 * (0.0 - 3.14159))
                                var v53 bool = (0 <= 127)
                                _ = v51
                                _ = v52
                                _ = v53
                            }
                            v2 += (v2 >> ((255) & 63))
                            if (true || false) {
                                v1 += 3.14159
                            } else {
                                var v54 int = (fi50 / (((fi50 - -1)) | 1))
                                v55 := GenS4{F0: (v5 % ((v5) | 1)), F1: 127, F2: 2}
                                _ = v54
                                _ = v55
                            }
                        } else {
                            _ = len(v4)
                            v2 -= -640
                        }
                        var v56 uint = 18446744073709551615
                        _ = v56
                    }
                    switch fi50 {
                    case 127:
                        if false {
                            if true {
                                for fi57 := 0; fi57 < 11; fi57++ {
                                    if true {
                                        v5 = (v5 / ((v5) | 1))
                                        if true {
                                            fmt.Println(fi50)
                                            var v58 int = (^(fi57))
                                            switch "\n\t" {
                                            case "\n\t":
                                                var v59 float64 = float64(v1)
                                                _ = v59
                                            case "a":
                                                v4[(fi57 + v58)] = 32768
                                            case "foo bar":
                                                var v60 float64 = (3.14159 + 3.14159)
                                                _ = v60
                                            }
                                            _ = v58
                                        } else {
                                            var v61 uint32 = 980
                                            _ = v61
                                        }
                                        switch (("hello" + "a") + v6) {
                                        case "hello":
                                            v6 = v6
                                        case "\n\t":
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        v62 := v4[fi50]
                                        if (18446744073709551615 == 2) {
                                            fmt.Println(v6, v1, v2)
                                            fmt.Println(fi57, fi57)
                                        } else {
                                            switch ("" + "") {
                                            case "hello":
                                                var v63 bool = false
                                                var v64 float32 = v1
                                                _ = v63
                                                _ = v64
                                            }
                                            var v65 uint = 174
                                            _ = v65
                                        }
                                        if (!false) {
                                            switch fi50 {
                                            case 4294967295:
                                                v6 = "a"
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v66 := v4[9223372036854775807]
                                            v62 = (v62 << ((685) & 63))
                                            _ = v66
                                        } else {
                                            if true {
                                                fmt.Println(v5)
                                                _ = len(v3)
                                                delete(v4, (fi57 - (-(fi57))))
                                            } else {
                                                if false {
                                                    var v67 int16 = (^(v2))
                                                    fmt.Println(v2, v1, fi50)
                                                    _ = v67
                                                }
                                                if ((0.1 * 1.5) <= (3.14159 * (float64(0.1) / 1.5))) {
                                                    switch fi50 {
                                                    case 128:
                                                        v62 += (v62 ^ (fi50 & 243))
                                                    case -1:
                                                        switch (fi57 << ((2) & 63)) {
                                                        case 32767:
                                                            if true {
                                                                v5 -= 42
                                                            }
                                                        case 2147483647:
                                                            fmt.Println(v6, v2)
                                                        }
                                                    case 4294967295:
                                                        v4[fi50] = int(v2)
                                                    }
                                                    v2 = 0
                                                    if true {
                                                        v68 := GenS3{F0: v2}
                                                        v5 = (v5 % ((v5) | 1))
                                                        _ = v68
                                                    }
                                                }
                                                v69 := GenS5{F0: 255, F1: (v5 - v5), F2: v6}
                                                if (v6 > (v6 + ("hello" + (v6 + (v6 + (("hello" + v6) + ("foo bar" + ("" + ("a" + v6))))))))) {
                                                    v1 -= v1
                                                    if true {
                                                        v70 := GenS6{F0: ("\n\t" + v6)}
                                                        _ = v70
                                                    } else {
                                                        _ = len(v4)
                                                        if false {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            var v71 string = v6
                                                            _ = v71
                                                        }
                                                        if (false || true) {
                                                            v6 = v6
                                                        }
                                                    }
                                                }
                                                _ = v69
                                            }
                                        }
                                        var v72 uint8 = 255
                                        fmt.Println(v6, fi50, v2)
                                        v2 += (v2 | v2)
                                        _ = v62
                                        _ = v72
                                    }
                                    if (true && (128 < 128)) {
                                        v2 = (v2 % (((^(v2))) | 1))
                                        switch (fi57 >> ((1) & 63)) {
                                        case -2147483648:
                                            v73 := GenS1{F0: 1, F1: ""}
                                            _ = v73
                                        case -704:
                                            v4[(fi57 | (fi57 + fi57))] = (fi57 & fi57)
                                        case 830:
                                            v74 := GenS7{F0: true, F1: (("foo bar" + ("a" + ("hello" + (("hello" + (v6 + (((((("a" + "") + v6) + "foo bar") + v6) + v6) + v6))) + ((v6 + ((v6 + "foo bar") + v6)) + ("hello" + "a")))))) + v6), F2: 252}
                                            _ = v74
                                        default:
                                            fmt.Println(v1)
                                        }
                                        fmt.Println(v2, v5, fi50)
                                    }
                                    var v75 int16 = v2
                                    _ = fi57
                                    _ = v75
                                }
                                var v76 uint8 = 227
                                _ = v76
                            }
                            for fi77 := 0; fi77 < 17; fi77++ {
                                var v78 uint32 = 2
                                _ = fi77
                                _ = v78
                            }
                        } else {
                            for fi79 := 0; fi79 < 13; fi79++ {
                                delete(v4, fi79)
                                _ = fi79
                            }
                            fmt.Println(v5, v5, v1)
                        }
                    case -1:
                        for fi80 := 0; fi80 < 20; fi80++ {
                            v3 = append(v3, v6)
                            fmt.Println(v6, fi50)
                            var v81 uint = 65535
                            _ = fi80
                            _ = v81
                        }
                    default:
                        var v82 uint32 = 887
                        _ = v82
                    }
                    var v83 float32 = 1.5
                    _ = v83
                } else {
                    for fi84 := 0; fi84 < 20; fi84++ {
                        var v85 float32 = v1
                        _ = fi84
                        _ = v85
                    }
                    if len(v3) > 0 {
                        v3[0] = v6
                    }
                    switch "hello" {
                    case "":
                        v86 := v4[fi50]
                        _ = v86
                    case "foo bar":
                        _ = len(v3)
                    case "a":
                        v5 = (v5 >> ((641) & 63))
                    }
                }
                _ = fi50
            }
        }
    default:
        for fi87 := 0; fi87 < 4; fi87++ {
            v4[fi87] = (fi87 << ((310) & 63))
            var v88 int16 = 1
            _ = fi87
            _ = v88
        }
    }
    var v89 float32 = v1
    var v90 uint64 = 42
    switch v6 {
    case "":
        v6 = ("hello" + "")
    case "a":
        var v91 int64 = -1
        _ = v91
    case "\n\t":
        v92 := GenS8{F0: (v2 + (v2 * (v2 + v2))), F1: 2, F2: 32767}
        _ = v92
    }
    switch 0 {
    case 0:
        v2 += v2
    case -2147483648:
        v93 := GenS7{F0: false, F1: ("" + v6), F2: 142}
        _ = v93
    }
    v94 := GenS9{F0: ((!true) || false)}
    v95 := GenS10{F0: 992, F1: 232}
    var v96 uint = 255
    for fi97 := 0; fi97 < 10; fi97++ {
        for fi98 := 0; fi98 < 16; fi98++ {
            v2 = -806
            _ = fi98
        }
        _ = fi97
    }
    delete(v4, 32768)
    if false {
        for fi99 := 0; fi99 < 14; fi99++ {
            var v100 int = fi99
            _ = fi99
            _ = v100
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v89
    _ = v90
    _ = v94
    _ = v95
    _ = v96
}
