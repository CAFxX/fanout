package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS7 struct { F0 int64; F1 uint; F2 uint }
type GenS8 struct { F0 uint16 }
func genFn0() uint16 {
    var v6 uint = 0
    v7 := []int{}
    v6 = v6
    _ = v6
    _ = v7
    return 65535
}

func main() {
    var v1 int16 = 42
    var v2 uint16 = 710
    var v3 int32 = 0
    v4 := make(map[uint32]int)
    var v5 bool = ((-2.25 / 1.5) > 3.14159)
    var v8 int32 = v3
    if (("a" + "foo bar") < ("" + ("" + ""))) {
        switch ("" + ("a" + "")) {
        case "hello":
            v9 := GenS1{F0: (v8 + -1), F1: 1}
            _ = v9
        case "a":
            for fi10 := 0; fi10 < 4; fi10++ {
                v11 := []int{(fi10 / ((fi10) | 1)), (fi10 ^ fi10), fi10, fi10}
                switch (fi10 - (fi10 + fi10)) {
                case 386:
                    v12 := GenS2{F0: genFn0(), F1: 742, F2: 58}
                    _ = v12
                default:
                    v13 := GenS1{F0: -1, F1: 255}
                    _ = v13
                }
                _ = fi10
                _ = v11
            }
        case "foo bar":
            for fi14 := 0; fi14 < 4; fi14++ {
                switch "foo bar" {
                case "":
                    var v15 int64 = 459
                    _ = v15
                case "hello":
                    if (("" < "") || true) {
                        fmt.Println(v5)
                    }
                default:
                    delete(v4, 2147483647)
                }
                _ = fi14
            }
        }
        for fi16 := 0; fi16 < 15; fi16++ {
            for fi17 := 0; fi17 < 10; fi17++ {
                switch ("\n\t" + "\n\t") {
                case "\n\t":
                    var v18 float64 = 1.5
                    _ = v18
                }
                var v19 int8 = 1
                _ = fi17
                _ = v19
            }
            _ = fi16
        }
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        fmt.Println(v5, v3)
        v3 -= (^(v3))
        fmt.Println(v2, v1)
        _ = fi20
    }
    var v21 int8 = -15
    var v22 int32 = v8
    var v23 uint = 298
    var v24 int8 = -128
    v25 := v4[2]
    if (true || v5) {
        v26 := []int{(^(v25)), v25, v25}
        for fi27 := 0; fi27 < 9; fi27++ {
            fmt.Println(v22, v24, fi27)
            fmt.Println(v5, v25, v2)
            if (!((-2.25 + (1.5 + 1e10)) == 1e10)) {
                if (v5 && false) {
                    var v28 string = "foo bar"
                    var v29 int32 = v3
                    _ = v28
                    _ = v29
                }
            } else {
                v30 := v4[4294967295]
                var v31 bool = v5
                _ = v30
                _ = v31
            }
            _ = fi27
        }
        _ = v26
    } else {
        v32 := GenS3{F0: 0}
        var v33 int16 = (v1 ^ (v1 % ((v1) | 1)))
        _ = v32
        _ = v33
    }
    fmt.Println(v23)
    for fi34 := 0; fi34 < 14; fi34++ {
        switch int(v2) {
        case -9223372036854775808:
            v35 := GenS1{F0: v22, F1: 42}
            _ = v35
        case 1:
            var v36 int = (v25 + (v25 / (((fi34 >> ((v23) & 63))) | 1)))
            _ = v36
        case -855:
            for fi37 := 0; fi37 < 8; fi37++ {
                v21 = v21
                v38 := GenS1{F0: v3, F1: 128}
                fmt.Println(fi37)
                if false {
                    if v5 {
                        if v5 {
                            v8 = (v3 - v22)
                        } else {
                            if (1e10 > 1e10) {
                                delete(v4, 282)
                                v39 := []int{32768, v25, 1}
                                _ = v39
                            }
                            var v40 float32 = (1e10 * 0.1)
                            if false {
                                var v41 int64 = 746
                                _ = v41
                            } else {
                                switch "" {
                                case "hello":
                                    _ = len(v4)
                                    var v42 int64 = 9223372036854775807
                                    switch int(v42) {
                                    case 127:
                                        fmt.Println(v24, v1, v42)
                                    case 775:
                                        if v5 {
                                            var v43 int8 = 1
                                            var v44 int32 = (v3 | -2147483648)
                                            delete(v4, 42)
                                            _ = v43
                                            _ = v44
                                        }
                                    default:
                                        var v45 int = (^(v25))
                                        _ = v45
                                    }
                                    _ = v42
                                }
                                v21 = v21
                                v1 = (v1 + (^(v1)))
                            }
                            _ = v40
                        }
                        switch fi34 {
                        case -9223372036854775808:
                            v46 := GenS3{F0: 2}
                            _ = v46
                        case -2147483648:
                            v47 := []int{v25}
                            _ = v47
                        }
                    } else {
                        var v48 int32 = -1
                        v22 = (v8 << ((uint(v2)) & 63))
                        v49 := GenS1{F0: v3, F1: 42}
                        v50 := []int{}
                        switch fi34 {
                        case -1:
                            v8 = (v48 >> ((2) & 63))
                        default:
                            var v51 int8 = v21
                            _ = v51
                        }
                        _ = v48
                        _ = v49
                        _ = v50
                    }
                } else {
                    v52 := GenS4{F0: (v8 - v8), F1: 126}
                    v1 += (v1 << (((^(v23))) & 63))
                    if true {
                        v53 := GenS5{F0: 44, F1: v24}
                        switch 269 {
                        case 4294967295:
                            fmt.Println(v24, v22)
                        case 9223372036854775807:
                            v21 = (v24 >> (((^(v23))) & 63))
                        }
                        var v54 uint64 = 18446744073709551615
                        _ = v53
                        _ = v54
                    }
                    _ = v52
                }
                _ = fi37
                _ = v38
            }
        }
        switch ("a" + "") {
        case "hello":
            for fi55 := 0; fi55 < 3; fi55++ {
                var v56 int8 = (v24 - v21)
                _ = fi55
                _ = v56
            }
        case "":
            var v57 uint8 = 255
            _ = v57
        }
        v58 := []int{fi34, (fi34 * fi34)}
        _ = fi34
        _ = v58
    }
    v24 = (-(v24))
    v23 += (v23 - v23)
    if (1 > 0) {
        switch (^(v25)) {
        case 32768:
            switch "\n\t" {
            case "":
                delete(v4, uint32(v23))
            case "a":
                v59 := []int{}
                _ = v59
            default:
                v60 := GenS6{F0: (v5 || (!((-(v1)) < (^(v1))))), F1: 127, F2: 237}
                _ = v60
            }
        case 127:
            var v61 uint16 = 2
            _ = v61
        case 823:
            if false {
                switch (v25 + 42) {
                case -9223372036854775808:
                    v5 = false
                case 128:
                    v62 := []int{128, v25, -2147483648}
                    _ = v62
                case -483:
                    for fi63 := 0; fi63 < 6; fi63++ {
                        for fi64 := 0; fi64 < 18; fi64++ {
                            switch (fi64 << (((v23 + 941)) & 63)) {
                            case 127:
                                var v65 uint = (^(v23))
                                _ = v65
                            case 4294967295:
                                switch 4294967295 {
                                case -9223372036854775808:
                                    fmt.Println(v1, v2, fi64)
                                case 810:
                                    v66 := GenS7{F0: 0, F1: v23, F2: v23}
                                    _ = v66
                                default:
                                    v67 := GenS2{F0: 384, F1: 42, F2: 93}
                                    _ = v67
                                }
                            case 769:
                                v68 := GenS6{F0: ((v1 << (((v23 - (v23 ^ (^(v23))))) & 63)) != (v1 + v1)), F1: 211, F2: 1}
                                _ = v68
                            }
                            v69 := GenS1{F0: (v22 >> ((2) & 63)), F1: 174}
                            _ = fi64
                            _ = v69
                        }
                        for fi70 := 0; fi70 < 18; fi70++ {
                            if (!true) {
                                var v71 int32 = v8
                                _ = v71
                            }
                            v2 = v2
                            _ = fi70
                        }
                        _ = fi63
                    }
                }
            }
        default:
            for fi72 := 0; fi72 < 5; fi72++ {
                for fi73 := 0; fi73 < 13; fi73++ {
                    v74 := GenS7{F0: -9223372036854775808, F1: v23, F2: 683}
                    var v75 uint64 = 18446744073709551615
                    v76 := []int{fi73}
                    fmt.Println(fi72, v1, v75)
                    _ = fi73
                    _ = v74
                    _ = v75
                    _ = v76
                }
                _ = fi72
            }
        }
    } else {
        switch (v25 << (((^(v23))) & 63)) {
        case 0:
            var v77 int = -9223372036854775808
            _ = v77
        default:
            for fi78 := 0; fi78 < 3; fi78++ {
                v79 := []int{128, v25, v25}
                v80 := GenS2{F0: v2, F1: 42, F2: 186}
                _ = fi78
                _ = v79
                _ = v80
            }
        }
        for fi81 := 0; fi81 < 10; fi81++ {
            var v82 int8 = (v21 + (-(v24)))
            var v83 int16 = int16(v22)
            _ = fi81
            _ = v82
            _ = v83
        }
        if (255 >= 255) {
            v84 := GenS8{F0: 1}
            fmt.Println(v8)
            _ = v84
        } else {
            v85 := []int{v25}
            _ = v85
        }
    }
    v21 -= (v21 % ((v21) | 1))
    var v86 int16 = v1
    var v87 int32 = (v8 ^ 1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v86
    _ = v87
}
