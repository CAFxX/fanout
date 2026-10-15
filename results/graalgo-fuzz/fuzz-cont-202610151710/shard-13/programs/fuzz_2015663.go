package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 string; F1 bool; F2 int8 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64; F1 uint; F2 uint16 }
func genFn0() int8 {
    if (true || (65535 < 0)) {
        for fi6 := 0; fi6 < 10; fi6++ {
            fmt.Println(fi6)
            fmt.Println(fi6)
            _ = fi6
        }
        var v7 int = 127
        var v8 bool = (true || (!((0 >= 1) || (!(("" != "\n\t") || (false || true))))))
        _ = v7
        _ = v8
    }
    if true {
        for fi9 := 0; fi9 < 15; fi9++ {
            if (((850 != 194) || (141 != 1)) || ("" <= "hello")) {
                fmt.Println(fi9)
                v10 := []int{fi9, (fi9 - fi9)}
                if false {
                    var v11 string = ("" + "foo bar")
                    _ = v11
                } else {
                    v12 := GenS1{F0: 929}
                    for fi13 := 0; fi13 < 5; fi13++ {
                        if false {
                            var v14 uint = 847
                            var v15 int16 = 32767
                            var v16 uint16 = 42
                            var v17 int32 = 0
                            v17 += -2147483648
                            switch (("" + "hello") + ("" + "foo bar")) {
                            case "hello":
                                switch "foo bar" {
                                case "":
                                    fmt.Println(v17, fi13, v14)
                                case "\n\t":
                                    v18 := GenS2{F0: 94, F1: int64(v16)}
                                    _ = v18
                                case "foo bar":
                                    var v19 int32 = v17
                                    _ = v19
                                default:
                                    var v20 uint32 = 2147483648
                                    _ = v20
                                }
                            case "\n\t":
                                if len(v10) > 0 {
                                    v10[0] = (fi13 | (fi13 % ((128) | 1)))
                                }
                                switch "foo bar" {
                                case "foo bar":
                                    if true {
                                        switch "\n\t" {
                                        case "":
                                            var v21 uint16 = v16
                                            _ = v21
                                        }
                                        v10 = append(v10, fi13)
                                        if (9223372036854775807 == 1) {
                                            if ((2 > (v16 - (v16 * (v16 << (((v14 + (^(v14)))) & 63))))) || (!((^(v14)) != v14))) {
                                                if len(v10) > 0 {
                                                    v10[0] = (fi13 >> (((v14 / (((^(v14))) | 1))) & 63))
                                                }
                                                v22 := GenS1{F0: -1}
                                                _ = v22
                                            } else {
                                                var v23 int64 = -9223372036854775808
                                                var v24 uint32 = 288
                                                _ = v23
                                                _ = v24
                                            }
                                        } else {
                                            fmt.Println(v16)
                                            var v25 int = fi9
                                            if len(v10) > 0 {
                                                v10[0] = (fi9 / (((-(v25))) | 1))
                                            }
                                            _ = v25
                                        }
                                    }
                                }
                                if true {
                                    var v26 int32 = (v17 * (v17 | (v17 ^ v17)))
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                    var v27 uint8 = 183
                                    _ = v26
                                    _ = v27
                                } else {
                                    fmt.Println(v14)
                                    fmt.Println(fi9, v16)
                                }
                                var v28 uint = v14
                                _ = v28
                            case "a":
                                var v29 int32 = -2147483648
                                _ = v29
                            }
                            if ("a" != (("hello" + "") + "a")) {
                                v30 := GenS2{F0: int8(v14), F1: 0}
                                if true {
                                    v31 := GenS2{F0: 0, F1: 1}
                                    var v32 int = fi13
                                    _ = v31
                                    _ = v32
                                } else {
                                    _ = len(v10)
                                    switch ("a" + ("a" + "")) {
                                    case "hello":
                                        v33 := GenS3{F0: -128, F1: (^(v14))}
                                        switch ("hello" + ("\n\t" + ((("\n\t" + "a") + "foo bar") + ""))) {
                                        case "hello":
                                            if len(v10) > 0 {
                                                v10[0] = fi9
                                            }
                                        case "a":
                                            var v34 uint16 = 2
                                            _ = v34
                                        }
                                        _ = v33
                                    case "a":
                                        var v35 int = (fi9 * 127)
                                        _ = v35
                                    }
                                }
                                _ = v30
                            } else {
                                v36 := GenS1{F0: v17}
                                _ = v36
                            }
                            v17 = int32(v16)
                            v15 -= v15
                            _ = v14
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        _ = fi13
                    }
                    if (42 < -1) {
                        if (!(!true)) {
                            var v37 bool = false
                            if ((1.5 - 0.0) == 1e10) {
                                v10 = append(v10, fi9)
                            }
                            _ = v37
                        }
                    } else {
                        var v38 uint16 = 0
                        v38 = (v38 % (((^(v38))) | 1))
                        _ = v38
                    }
                    _ = v12
                }
                _ = v10
            } else {
                switch "a" {
                case "\n\t":
                    v39 := GenS3{F0: -1, F1: 2}
                    _ = v39
                default:
                    if (2147483647 != 4294967295) {
                        if false {
                            if (1 != 2147483647) {
                                v40 := GenS4{F0: "a", F1: true, F2: 100}
                                fmt.Println(fi9)
                                v41 := GenS3{F0: 127, F1: 18446744073709551615}
                                _ = v40
                                _ = v41
                            } else {
                                var v42 uint16 = 679
                                v42 += v42
                                var v43 string = ""
                                _ = v42
                                _ = v43
                            }
                            v44 := []int{fi9, (^(fi9)), (fi9 ^ fi9), (fi9 % ((0) | 1))}
                            switch (fi9 << ((18446744073709551615) & 63)) {
                            case 1:
                                fmt.Println(fi9)
                            case 128:
                                var v45 float64 = (0.0 - -2.25)
                                _ = v45
                            case 9223372036854775807:
                                if false {
                                    for fi46 := 0; fi46 < 8; fi46++ {
                                        fmt.Println(fi9, fi46)
                                        _ = fi46
                                    }
                                    var v47 uint16 = 1
                                    v47 += (v47 ^ 2)
                                    _ = v47
                                }
                            default:
                                v48 := GenS4{F0: ("" + "foo bar"), F1: (0 <= 245), F2: -1}
                                _ = v48
                            }
                            _ = v44
                        }
                        switch fi9 {
                        case 128:
                            var v49 int8 = -128
                            _ = v49
                        case 1:
                            v50 := []int{(fi9 - fi9), -1, fi9, 42}
                            _ = v50
                        case 42:
                            var v51 float64 = 1.5
                            _ = v51
                        }
                    }
                }
            }
            v52 := GenS1{F0: 237}
            switch "" {
            case "foo bar":
                v53 := []int{}
                _ = v53
            default:
                v54 := []int{fi9}
                _ = v54
            }
            _ = fi9
            _ = v52
        }
    }
    var v55 float32 = -2.25
    for fi56 := 0; fi56 < 3; fi56++ {
        if true {
            v55 -= 1.5
            v57 := []int{(fi56 + 2147483647)}
            v55 -= v55
            _ = v57
        }
        var v58 uint8 = 232
        _ = fi56
        _ = v58
    }
    fmt.Println(v55)
    _ = v55
    return -8
}

func main() {
    var v1 bool = true
    var v2 int64 = 0
    var v3 bool = (558 == 127)
    var v4 int = -9223372036854775808
    v5 := make(map[string]string)
    for fi59 := 0; fi59 < 2; fi59++ {
        switch fi59 {
        case 0:
            v60 := GenS4{F0: "", F1: v1, F2: 127}
            _ = v60
        case 9223372036854775807:
            v61 := []int{v4, (^(v4)), 128}
            _ = v61
        default:
            for fi62 := 0; fi62 < 15; fi62++ {
                var v63 uint64 = 413
                _ = fi62
                _ = v63
            }
        }
        for fi64 := 0; fi64 < 18; fi64++ {
            if v3 {
                v2 = v2
                fmt.Println(v1)
            }
            _ = fi64
        }
        for fi65 := 0; fi65 < 19; fi65++ {
            v66 := []int{fi59, fi59, (v4 >> ((168) & 63))}
            _ = fi65
            _ = v66
        }
        _ = fi59
    }
    for fi67 := 0; fi67 < 2; fi67++ {
        fmt.Println(v4, v1, v1)
        var v68 int8 = 42
        switch ("a" + ("\n\t" + "a")) {
        case "hello":
            var v69 int32 = -788
            _ = v69
        case "a":
            v70 := v5[""]
            _ = v70
        case "":
            if v3 {
                var v71 uint16 = 255
                var v72 uint16 = (^(v71))
                _ = v71
                _ = v72
            } else {
                v2 = (v2 & v2)
                fmt.Println(v4, v4)
                var v73 int16 = 1
                _ = v73
            }
        default:
            var v74 bool = true
            _ = v74
        }
        _ = fi67
        _ = v68
    }
    v1 = (v4 != (v4 * 285))
    var v75 uint64 = 42
    var v76 int64 = v2
    _ = len(v5)
    switch (("hello" + (("" + "") + "hello")) + "") {
    case "\n\t":
        v77 := v5["a"]
        _ = v77
    case "foo bar":
        var v78 uint16 = 32767
        _ = v78
    case "a":
        for fi79 := 0; fi79 < 7; fi79++ {
            for fi80 := 0; fi80 < 11; fi80++ {
                if v1 {
                    v76 -= (-(v2))
                    v81 := GenS5{F0: 0, F1: -128}
                    _ = v81
                }
                _ = fi80
            }
            v82 := GenS6{F0: 42}
            _ = fi79
            _ = v82
        }
    }
    switch v4 {
    case -48:
        var v83 uint64 = (v75 - v75)
        _ = v83
    case -269:
        v3 = (!(0.0 < 0.0))
    case 185:
        var v84 bool = (42 < 255)
        _ = v84
    default:
        var v85 bool = v1
        _ = v85
    }
    _ = len(v5)
    if false {
        fmt.Println(v76)
        v3 = false
    }
    var v86 uint32 = 349
    if v3 {
        v1 = true
        v87 := GenS3{F0: 127, F1: 1}
        for fi88 := 0; fi88 < 19; fi88++ {
            for fi89 := 0; fi89 < 13; fi89++ {
                var v90 int8 = 1
                v86 = (^(v86))
                _ = fi89
                _ = v90
            }
            for fi91 := 0; fi91 < 16; fi91++ {
                v92 := GenS7{F0: (v75 / ((v75) | 1)), F1: 252, F2: 42}
                _ = fi91
                _ = v92
            }
            v75 = 2
            _ = fi88
        }
        _ = v87
    }
    v93 := []int{}
    v94 := GenS7{F0: v75, F1: 255, F2: 0}
    var v95 string = "foo bar"
    var v96 uint16 = 2
    var v97 float32 = 1e10
    var v98 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v86:", v86)
    fmt.Println("slicelen v93:", len(v93), v93)
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v75
    _ = v76
    _ = v86
    _ = v93
    _ = v94
    _ = v95
    _ = v96
    _ = v97
    _ = v98
}
