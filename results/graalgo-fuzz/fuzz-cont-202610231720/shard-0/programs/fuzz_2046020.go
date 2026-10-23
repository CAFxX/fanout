package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS7 struct { F0 int32; F1 uint64 }
type GenS8 struct { F0 int8; F1 int8; F2 uint }
type GenS9 struct { F0 int; F1 uint }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int8; F1 int }
type GenS12 struct { F0 string; F1 int64 }
type GenS13 struct { F0 uint64; F1 int; F2 bool }
type GenS14 struct { F0 uint }
func genFn0(p0 int64) uint16 {
    var v6 uint16 = 255
    var v7 float32 = 0.0
    if true {
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 int = fi8
            _ = fi8
            _ = v9
        }
        switch "" {
        case "\n\t":
            v6 = 504
        default:
            for fi10 := 0; fi10 < 19; fi10++ {
                v11 := []int{}
                if (127 > 128) {
                    if len(v11) > 0 {
                        v11[0] = (fi10 - 127)
                    }
                    for fi12 := 0; fi12 < 16; fi12++ {
                        v13 := GenS1{F0: (("a" + "") + ""), F1: false, F2: 0}
                        if (-128 != -128) {
                            v6 = 0
                        }
                        _ = fi12
                        _ = v13
                    }
                }
                _ = fi10
                _ = v11
            }
        }
        p0 = 1
    }
    var v14 int8 = 127
    _ = p0
    _ = v14
    _ = v6
    _ = v7
    return (v6 | (^(v6)))
}

func main() {
    var v1 uint = 1
    v2 := []string{"foo bar", ("a" + ("" + ""))}
    v3 := []uint32{1, 65535, 125}
    var v4 float64 = 1e10
    var v5 int8 = 127
    v15 := GenS2{F0: 4294967295, F1: 128}
    if false {
        if (2147483648 != 0) {
            switch "" {
            case "\n\t":
                v16 := GenS3{F0: 32768, F1: 255}
                _ = v16
            case "a":
                v5 += (v5 + (v5 - v5))
            }
            if (float64((0.0 + (3.14159 - (3.14159 * 0.0)))) <= 1e10) {
                v1 -= (v1 | (^(v1)))
                switch "" {
                case "hello":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "foo bar":
                    for fi17 := 0; fi17 < 8; fi17++ {
                        v18 := GenS4{F0: 0, F1: 1}
                        for fi19 := 0; fi19 < 9; fi19++ {
                            if len(v2) > 0 {
                                v2[0] = ""
                            }
                            switch (("\n\t" + ("foo bar" + ((("\n\t" + (("foo bar" + "\n\t") + "foo bar")) + "a") + ""))) + ("" + "a")) {
                            case "\n\t":
                                if true {
                                    var v20 int8 = (v5 % ((82) | 1))
                                    v5 += 102
                                    v5 = -1
                                    if len(v3) > 0 {
                                        v3[0] = 368
                                    }
                                    _ = v20
                                }
                            case "a":
                                v21 := GenS5{F0: 65535, F1: false}
                                v22 := GenS2{F0: 625, F1: 155}
                                v4 = 0.0
                                _ = v21
                                _ = v22
                            case "foo bar":
                                var v23 string = ("" + (("\n\t" + ("\n\t" + "")) + ""))
                                _ = v23
                            }
                            _ = fi19
                        }
                        _ = fi17
                        _ = v18
                    }
                }
            }
            if true {
                for fi24 := 0; fi24 < 19; fi24++ {
                    var v25 int = (fi24 % (((fi24 - fi24)) | 1))
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi24
                    _ = v25
                }
                var v26 uint = (v1 ^ 2)
                _ = v26
            }
        }
        if true {
            var v27 int64 = 0
            v1 = (v1 >> (((v1 | v1)) & 63))
            _ = v27
        }
        var v28 int16 = int16(v5)
        _ = v28
    }
    v4 -= float64(0.1)
    if (221 != 224) {
        if false {
            switch 9223372036854775807 {
            case 2147483647:
                if true {
                    v5 -= (v5 ^ v5)
                    if true {
                        for fi29 := 0; fi29 < 4; fi29++ {
                            v30 := GenS5{F0: 148, F1: false}
                            _ = fi29
                            _ = v30
                        }
                    } else {
                        var v31 int = -710
                        _ = v31
                    }
                } else {
                    for fi32 := 0; fi32 < 13; fi32++ {
                        fmt.Println(v5, v5)
                        _ = fi32
                    }
                    var v33 float32 = 1e10
                    if len(v3) > 0 {
                        v3[0] = 0
                    }
                    _ = v33
                }
            case 1:
                var v34 int32 = -2147483648
                _ = v34
            default:
                v35 := GenS6{F0: (^(v1)), F1: 0, F2: 9}
                _ = v35
            }
            for fi36 := 0; fi36 < 5; fi36++ {
                fmt.Println(v4)
                for fi37 := 0; fi37 < 2; fi37++ {
                    switch (("a" + "foo bar") + "\n\t") {
                    case "hello":
                        if true {
                            v38 := GenS5{F0: 1, F1: ("a" >= "\n\t")}
                            _ = v38
                        }
                    case "foo bar":
                        fmt.Println(fi36, v5, v1)
                    case "":
                        var v39 int16 = 42
                        _ = v39
                    }
                    _ = fi37
                }
                v40 := GenS7{F0: -1, F1: 42}
                _ = fi36
                _ = v40
            }
        }
        switch ("hello" + "foo bar") {
        case "a":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "":
            for fi41 := 0; fi41 < 17; fi41++ {
                if (false && false) {
                    if true {
                        v5 -= -22
                        var v42 float32 = (1e10 / 1e10)
                        var v43 int32 = 784
                        _ = v42
                        _ = v43
                    }
                    if false {
                        if (false || ("foo bar" >= "")) {
                            v44 := GenS8{F0: v5, F1: v5, F2: 670}
                            _ = v44
                        } else {
                            v5 = v5
                            v45 := GenS9{F0: -1, F1: v1}
                            switch ("hello" + "foo bar") {
                            case "hello":
                                var v46 uint8 = 127
                                _ = v46
                            default:
                                if true {
                                    var v47 int64 = 990
                                    _ = v47
                                } else {
                                    v5 = (v5 * v5)
                                    v48 := GenS8{F0: 100, F1: (v5 * (v5 - 86)), F2: v1}
                                    for fi49 := 0; fi49 < 3; fi49++ {
                                        v1 = v1
                                        _ = fi49
                                    }
                                    _ = v48
                                }
                            }
                            _ = v45
                        }
                    }
                }
                var v50 int64 = 1
                _ = fi41
                _ = v50
            }
        case "hello":
            var v51 uint = v1
            _ = v51
        default:
            v52 := GenS9{F0: -862, F1: v1}
            _ = v52
        }
    }
    if (!(2147483647 != 1)) {
        v1 = (v1 >> ((42) & 63))
        v53 := GenS2{F0: 164, F1: 2}
        var v54 string = ""
        _ = v53
        _ = v54
    }
    var v55 uint64 = 2
    if (("\n\t" + "hello") <= ("a" + "\n\t")) {
        for fi56 := 0; fi56 < 11; fi56++ {
            switch "\n\t" {
            case "hello":
                _ = len(v3)
            default:
                for fi57 := 0; fi57 < 8; fi57++ {
                    if ((float32(((float64((0.0 - float32(v4))) - float64(1.5)) + v4)) + (-2.25 - float32(v4))) == 1.5) {
                        if len(v3) > 0 {
                            v3[0] = 65535
                        }
                    } else {
                        if false {
                            var v58 int = 386
                            v59 := GenS10{F0: (!(209 == 661))}
                            v3 = append(v3, 65535)
                            switch ("hello" + "\n\t") {
                            case "":
                                if true {
                                    switch "" {
                                    case "":
                                        switch ("" + "\n\t") {
                                        case "":
                                            if len(v2) > 0 {
                                                v2[0] = "foo bar"
                                            }
                                        }
                                    default:
                                        v60 := GenS11{F0: int8(v55), F1: 32767}
                                        _ = v60
                                    }
                                    if false {
                                        fmt.Println(v58, fi56, v58)
                                        v58 = fi57
                                        var v61 int8 = 0
                                        _ = v61
                                    }
                                }
                            case "foo bar":
                                v55 += v55
                            }
                            _ = v58
                            _ = v59
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        var v62 string = ("" + "\n\t")
                        if (!false) {
                            v1 = v1
                            var v63 int = (-(fi56))
                            _ = v63
                        } else {
                            switch v62 {
                            case "hello":
                                v64 := GenS3{F0: fi56, F1: 2}
                                _ = v64
                            case "foo bar":
                                var v65 uint8 = 127
                                _ = v65
                            case "":
                                if false {
                                    switch -2147483648 {
                                    case 32768:
                                        if false {
                                            v1 = v1
                                        } else {
                                            v1 = 255
                                        }
                                    case 42:
                                        v3 = append(v3, 665)
                                    case 979:
                                        switch (^(fi57)) {
                                        case -445:
                                            v4 = v4
                                        case 32768:
                                            v66 := GenS1{F0: v62, F1: (0.0 == v4), F2: 2147483647}
                                            _ = v66
                                        }
                                    }
                                } else {
                                    var v67 uint32 = uint32(v1)
                                    v1 = 581
                                    var v68 int16 = 340
                                    fmt.Println(v55)
                                    _ = v67
                                    _ = v68
                                }
                            default:
                                var v69 int32 = 0
                                _ = v69
                            }
                            switch fi57 {
                            case -1:
                                if false {
                                    var v70 int16 = 152
                                    var v71 uint32 = 369
                                    v62 = v62
                                    _ = v70
                                    _ = v71
                                }
                            case 4294967295:
                                switch (fi57 ^ fi57) {
                                case 1:
                                    if len(v3) > 0 {
                                        v3[0] = 4294967295
                                    }
                                default:
                                    if false {
                                        if false {
                                            var v72 bool = (0.1 != 1.5)
                                            var v73 int8 = (^(v5))
                                            switch (fi56 << ((v1) & 63)) {
                                            case 987:
                                                if true {
                                                    var v74 string = (v62 + "hello")
                                                    if true {
                                                        if v72 {
                                                            v4 += (v4 - float64(3.14159))
                                                            v4 -= v4
                                                            var v75 int32 = -2147483648
                                                            _ = v75
                                                        } else {
                                                            v72 = (true && v72)
                                                        }
                                                        v76 := GenS2{F0: (v55 * v55), F1: 42}
                                                        var v77 int16 = 42
                                                        _ = v76
                                                        _ = v77
                                                    }
                                                    _ = v74
                                                } else {
                                                    switch v62 {
                                                    case "hello":
                                                        var v78 int32 = 910
                                                        _ = v78
                                                    case "foo bar":
                                                        if v72 {
                                                            _ = len(v2)
                                                        }
                                                    case "a":
                                                        if (!true) {
                                                            var v79 int32 = int32(fi57)
                                                            var v80 int16 = 42
                                                            _ = v79
                                                            _ = v80
                                                        }
                                                    }
                                                    fmt.Println(v5)
                                                    if v72 {
                                                        var v81 string = v62
                                                        _ = v81
                                                    }
                                                }
                                            case 0:
                                                v82 := GenS3{F0: (fi57 / (((-(fi56))) | 1)), F1: v1}
                                                _ = v82
                                            default:
                                                v83 := GenS12{F0: v62, F1: 848}
                                                _ = v83
                                            }
                                            _ = v72
                                            _ = v73
                                        }
                                        v84 := GenS13{F0: v55, F1: -919, F2: ((!false) || false)}
                                        v4 -= 1e10
                                        _ = v84
                                    }
                                }
                            }
                            var v85 uint8 = 127
                            _ = v85
                        }
                        v1 += v1
                        _ = v62
                    }
                    _ = fi57
                }
            }
            v86 := GenS14{F0: v1}
            _ = fi56
            _ = v86
        }
    }
    v55 += (v55 << (((v1 / (((v1 / ((v1) | 1))) | 1))) & 63))
    v55 += (v55 << ((100) & 63))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v87 float32 = -2.25
    _ = len(v3)
    var v88 string = ("" + "a")
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v87
    _ = v88
}
