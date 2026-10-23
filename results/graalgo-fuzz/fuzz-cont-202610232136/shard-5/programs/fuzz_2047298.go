package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 uint; F2 bool }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int8; F1 uint32 }
type GenS10 struct { F0 uint8; F1 int8; F2 int8 }
type GenS11 struct { F0 int; F1 int8; F2 bool }
func genFn0(p0 int, p1 uint32) bool {
    for fi7 := 0; fi7 < 9; fi7++ {
        p0 -= 42
        if true {
            for fi8 := 0; fi8 < 13; fi8++ {
                fmt.Println(fi8, p0, fi7)
                v9 := GenS1{F0: 911}
                v10 := GenS1{F0: 2147483647}
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        p0 = fi7
        _ = fi7
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return ((!(!(!(1 != 1)))) && (0 >= 32767))
}

func main() {
    var v1 int = 4294967295
    var v2 int = -583
    var v3 int8 = 1
    v4 := make(map[string]bool)
    v5 := []uint32{1}
    v6 := make(map[uint32]uint32)
    var v11 int64 = -1
    _ = len(v6)
    if true {
        var v12 float64 = float64((-2.25 * (1.5 + 3.14159)))
        var v13 uint8 = 128
        _ = len(v5)
        _ = v12
        _ = v13
    } else {
        v11 = -9223372036854775808
    }
    v14 := GenS1{F0: -2147483648}
    var v15 int8 = (v3 ^ (v3 - (v3 | (v3 >> ((13) & 63)))))
    if false {
        if false {
            switch "hello" {
            case "foo bar":
                var v16 uint = 865
                _ = v16
            case "\n\t":
                if (!false) {
                    v1 += -87
                    if true {
                        v17 := GenS2{F0: v2, F1: 178}
                        for fi18 := 0; fi18 < 16; fi18++ {
                            switch (-(fi18)) {
                            case 32768:
                                if (0 <= 42) {
                                    v19 := GenS1{F0: -1}
                                    _ = v19
                                }
                            case -2147483648:
                                if true {
                                    fmt.Println(fi18)
                                    v20 := GenS2{F0: fi18, F1: 2}
                                    v15 = int8(v2)
                                    _ = v20
                                }
                            }
                            _ = fi18
                        }
                        _ = v17
                    } else {
                        if true {
                            v3 = v3
                        } else {
                            v21 := GenS1{F0: 0}
                            _ = v21
                        }
                    }
                } else {
                    var v22 int64 = (v11 | v11)
                    v1 = 9223372036854775807
                    _ = v22
                }
            case "":
                v23 := GenS2{F0: -272, F1: 128}
                _ = v23
            }
            v24 := GenS1{F0: 295}
            _ = v24
        }
        for fi25 := 0; fi25 < 9; fi25++ {
            v26 := GenS2{F0: v2, F1: 255}
            var v27 uint16 = 284
            _ = fi25
            _ = v26
            _ = v27
        }
    } else {
        for fi28 := 0; fi28 < 15; fi28++ {
            v3 += v15
            _ = fi28
        }
    }
    if false {
        if (("" + "hello") > "hello") {
            if false {
                switch ("" + "\n\t") {
                case "":
                    for fi29 := 0; fi29 < 19; fi29++ {
                        switch (v1 - (fi29 - v2)) {
                        case -1:
                            for fi30 := 0; fi30 < 10; fi30++ {
                                v31 := GenS3{F0: "foo bar", F1: ("\n\t" + ("foo bar" + (("" + ("hello" + "\n\t")) + "foo bar")))}
                                var v32 string = ("a" + "")
                                fmt.Println(v3, v3)
                                var v33 float64 = 1.5
                                _ = fi30
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                        case 294:
                            switch fi29 {
                            case 128:
                                if (((((("\n\t" != (("" + "foo bar") + "a")) || (false && (637 != 0))) && true) || (42 < 18446744073709551615)) || false) || false) {
                                    delete(v4, "foo bar")
                                    for fi34 := 0; fi34 < 7; fi34++ {
                                        v2 = (fi29 - fi34)
                                        _ = fi34
                                    }
                                    for fi35 := 0; fi35 < 16; fi35++ {
                                        v11 = v11
                                        v15 = (v3 | v3)
                                        if false {
                                            var v36 bool = true
                                            _ = v36
                                        }
                                        _ = fi35
                                    }
                                }
                            case -9223372036854775808:
                                v37 := GenS2{F0: 42, F1: 255}
                                _ = v37
                            case 1:
                                for fi38 := 0; fi38 < 15; fi38++ {
                                    v2 = 32768
                                    _ = fi38
                                }
                            default:
                                v39 := GenS2{F0: fi29, F1: 195}
                                _ = v39
                            }
                        default:
                            delete(v6, 238)
                        }
                        v40 := GenS4{F0: 121}
                        _ = fi29
                        _ = v40
                    }
                case "\n\t":
                    fmt.Println(v2, v2, v2)
                default:
                    _ = len(v4)
                }
                for fi41 := 0; fi41 < 5; fi41++ {
                    var v42 float32 = 1e10
                    _ = fi41
                    _ = v42
                }
                _ = len(v5)
            }
            var v43 int32 = 1
            switch 2147483647 {
            case -2147483648:
                fmt.Println(v3, v11)
            case 988:
                v44 := GenS2{F0: v2, F1: 148}
                _ = v44
            default:
                fmt.Println(v15, v1)
            }
            _ = v43
        }
        var v45 int8 = (v15 % ((v15) | 1))
        v46 := GenS5{F0: 36, F1: 1, F2: false}
        _ = v45
        _ = v46
    } else {
        v1 = v2
        var v47 uint16 = 2
        _ = v47
    }
    var v48 int8 = 33
    if false {
        switch v2 {
        case 127:
            for fi49 := 0; fi49 < 3; fi49++ {
                var v50 uint32 = 839
                _ = fi49
                _ = v50
            }
        default:
            v6[65535] = 65535
        }
        v1 = v1
        if (!false) {
            fmt.Println(v3, v15, v15)
            for fi51 := 0; fi51 < 5; fi51++ {
                _ = len(v5)
                var v52 float32 = ((((1e10 * ((1e10 - 0.0) / ((1e10 + 0.0) / 1.5))) - 0.1) * -2.25) + 0.1)
                v1 -= 32768
                _ = fi51
                _ = v52
            }
            fmt.Println(v3, v2)
        } else {
            if false {
                var v53 uint = 42
                _ = v53
            } else {
                delete(v6, 65535)
                for fi54 := 0; fi54 < 17; fi54++ {
                    v48 = v48
                    var v55 float32 = -2.25
                    _ = fi54
                    _ = v55
                }
                for fi56 := 0; fi56 < 3; fi56++ {
                    switch "" {
                    case "\n\t":
                        for fi57 := 0; fi57 < 12; fi57++ {
                            if (!(-2.25 >= (float32(((0.1 - float64(1e10)) * ((-2.25 / 3.14159) / (1.5 - float64(0.0))))) + 1.5))) {
                                switch (("hello" + ("hello" + "hello")) + ("foo bar" + "hello")) {
                                case "hello":
                                    var v58 int32 = -1
                                    _ = v58
                                case "foo bar":
                                    v1 -= -838
                                case "\n\t":
                                    switch (v2 >> ((4294967295) & 63)) {
                                    case 619:
                                        var v59 uint16 = 255
                                        _ = v59
                                    case 14:
                                        delete(v6, 0)
                                    default:
                                        if len(v5) > 0 {
                                            v5[0] = 900
                                        }
                                    }
                                }
                                if false {
                                    fmt.Println(v48)
                                    if true {
                                        switch ("" + "") {
                                        case "":
                                            v60 := v4["hello"]
                                            _ = v60
                                        case "\n\t":
                                            var v61 int32 = -804
                                            _ = v61
                                        default:
                                            v15 = v48
                                        }
                                        var v62 int64 = (v11 + v11)
                                        if len(v5) > 0 {
                                            v5[0] = 42
                                        }
                                        v6[2147483648] = 4294967295
                                        _ = v62
                                    } else {
                                        var v63 uint16 = 32767
                                        switch (("\n\t" + (("a" + "\n\t") + ("a" + ""))) + "") {
                                        case "foo bar":
                                            var v64 string = ((("a" + "\n\t") + "hello") + ("a" + ("a" + (("hello" + "hello") + "foo bar"))))
                                            _ = v64
                                        }
                                        v6[2147483648] = uint32(v63)
                                        v65 := GenS6{F0: 99, F1: 0}
                                        _ = v63
                                        _ = v65
                                    }
                                    v66 := GenS6{F0: 4294967295, F1: 448}
                                    _ = v66
                                } else {
                                    v67 := GenS7{F0: 128}
                                    v5 = append(v5, 65535)
                                    _ = v67
                                }
                            }
                            v48 -= v15
                            switch "hello" {
                            case "\n\t":
                                fmt.Println(fi57)
                            }
                            var v68 uint = 42
                            _ = fi57
                            _ = v68
                        }
                    case "":
                        v15 = v15
                    case "hello":
                        v6[2147483648] = 2147483647
                    default:
                        v69 := GenS3{F0: "foo bar", F1: "foo bar"}
                        _ = v69
                    }
                    for fi70 := 0; fi70 < 19; fi70++ {
                        switch "a" {
                        case "":
                            if false {
                                v3 = v3
                                var v71 uint8 = 194
                                if ((^(v11)) < v11) {
                                    var v72 uint16 = 62
                                    var v73 string = "hello"
                                    fmt.Println(v2, v72, v2)
                                    _ = v72
                                    _ = v73
                                } else {
                                    v71 += (^(v71))
                                    switch "foo bar" {
                                    case "":
                                        var v74 int = (v2 - (v1 / (((v1 & v2)) | 1)))
                                        _ = v74
                                    case "a":
                                        switch ("a" + ("foo bar" + "")) {
                                        case "":
                                            v71 = v71
                                        case "hello":
                                            v75 := GenS3{F0: "hello", F1: (("foo bar" + "hello") + (((((("" + "hello") + "a") + ("" + "foo bar")) + "foo bar") + "hello") + "foo bar"))}
                                            _ = v75
                                        default:
                                            switch -1 {
                                            case 42:
                                                v2 = v2
                                            default:
                                                fmt.Println(v71)
                                            }
                                        }
                                    }
                                    if true {
                                        if false {
                                            v1 = (^(fi70))
                                            _ = len(v6)
                                        }
                                    } else {
                                        var v76 uint32 = uint32(v71)
                                        _ = v76
                                    }
                                }
                                _ = v71
                            } else {
                                if (!false) {
                                    v2 += -17
                                    v2 += (fi56 / (((-(v2))) | 1))
                                    if true {
                                        var v77 int64 = v11
                                        _ = v77
                                    } else {
                                        v3 = (v15 >> ((65535) & 63))
                                        v11 = 42
                                        _ = len(v5)
                                    }
                                } else {
                                    v15 += (-(v15))
                                    var v78 int = fi56
                                    _ = v78
                                }
                                fmt.Println(v15, v1)
                                switch (v2 / ((v2) | 1)) {
                                case 42:
                                    var v79 int32 = 2147483647
                                    _ = v79
                                case -579:
                                    v80 := GenS8{F0: 127}
                                    _ = v80
                                case 342:
                                    var v81 string = "\n\t"
                                    _ = v81
                                default:
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                            }
                        case "a":
                            v82 := GenS2{F0: fi56, F1: 48}
                            _ = v82
                        case "hello":
                            var v83 int16 = 1
                            _ = v83
                        default:
                            var v84 int64 = (v11 * (v11 / ((v11) | 1)))
                            _ = v84
                        }
                        if (0.0 >= float32(0.1)) {
                            if (-197 != 2147483647) {
                                switch "foo bar" {
                                case "\n\t":
                                    var v85 int8 = v3
                                    _ = v85
                                case "":
                                    var v86 float64 = 1.5
                                    _ = v86
                                case "a":
                                    delete(v6, 675)
                                default:
                                    v15 = 46
                                }
                                v87 := GenS9{F0: v3, F1: 0}
                                _ = v87
                            } else {
                                v1 = (^(fi56))
                                v15 = v15
                            }
                        } else {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            if ((("" + "foo bar") + "") >= "a") {
                                fmt.Println(v1, fi70, v15)
                            }
                        }
                        var v88 uint16 = 1
                        _ = fi70
                        _ = v88
                    }
                    v11 = (v11 % ((v11) | 1))
                    _ = fi56
                }
            }
            v2 = v2
        }
    } else {
        fmt.Println(v48, v3, v1)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v89 := GenS8{F0: 1}
    switch (v2 | (v2 * v1)) {
    case -2147483648:
        delete(v6, 285)
    case -1:
        var v90 float64 = -2.25
        _ = v90
    case 32767:
        switch ("a" + "foo bar") {
        case "a":
            var v91 float32 = ((1.5 + ((1e10 / float32(1e10)) * (((float32(-2.25) - 0.1) / -2.25) / 0.1))) * ((1.5 + -2.25) / 3.14159))
            _ = v91
        default:
            v92 := GenS10{F0: 126, F1: -128, F2: v48}
            _ = v92
        }
    default:
        v15 += (^(v15))
    }
    switch 32768 {
    case 4294967295:
        var v93 float64 = 1.5
        _ = v93
    case -827:
        v3 += -1
    case 127:
        for fi94 := 0; fi94 < 7; fi94++ {
            v15 += v15
            v95 := GenS11{F0: 1, F1: v48, F2: (true || (true && true))}
            _ = fi94
            _ = v95
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
    _ = v89
}
