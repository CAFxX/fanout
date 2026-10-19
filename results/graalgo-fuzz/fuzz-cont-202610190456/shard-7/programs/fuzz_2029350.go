package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint32; F1 int16 }
type GenS7 struct { F0 string; F1 int }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int64; F1 int }
type GenS10 struct { F0 uint8; F1 uint16 }
type GenS11 struct { F0 int8; F1 bool; F2 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = 4294967295
    var v3 uint64 = 18446744073709551615
    v4 := make(map[string]string)
    v2 -= (-(v2))
    switch v2 {
    case 42:
        var v5 uint64 = 0
        _ = v5
    case 32768:
        delete(v1, 65535)
    case 0:
        delete(v1, 2147483647)
    default:
        v6 := []int{(-(v2)), (-(v2)), (v2 & 9223372036854775807), 2147483647}
        _ = v6
    }
    v7 := GenS1{F0: 65535, F1: v2}
    v8 := GenS2{F0: -1}
    if (!(930 == 2147483647)) {
        v9 := GenS3{F0: 919}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 14; fi10++ {
            v2 += fi10
            fmt.Println(v2)
            _ = fi10
        }
        v2 += v2
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        if false {
            var v12 uint8 = 42
            for fi13 := 0; fi13 < 11; fi13++ {
                if (!(1 == -1)) {
                    v2 = 1
                    var v14 uint8 = (^(v12))
                    _ = v14
                } else {
                    if false {
                        var v15 uint32 = 52
                        if ((v2 - (-(fi11))) != -1) {
                            if true {
                                v2 = v2
                            } else {
                                var v16 uint32 = (^(v15))
                                switch "" {
                                case "foo bar":
                                    var v17 int16 = 142
                                    _ = v17
                                case "hello":
                                    v18 := GenS1{F0: 42, F1: (fi11 | 859)}
                                    _ = v18
                                case "":
                                    var v19 bool = false
                                    v16 = v16
                                    if (-337 != -1) {
                                        var v20 float64 = 0.1
                                        v16 -= (v15 ^ v15)
                                        var v21 uint32 = v16
                                        _ = v20
                                        _ = v21
                                    } else {
                                        if v19 {
                                            v22 := []int{-60, fi11}
                                            if len(v22) > 0 {
                                                v22[0] = -549
                                            }
                                            if v19 {
                                                v12 += (v12 + 155)
                                                if v19 {
                                                    fmt.Println(v15, v15)
                                                }
                                            }
                                            switch "a" {
                                            case "\n\t":
                                                v23 := GenS4{F0: 32767}
                                                _ = v23
                                            default:
                                                v15 = (v15 & 2147483648)
                                            }
                                            if (42 == 0) {
                                                v24 := GenS5{F0: 9223372036854775807}
                                                fmt.Println(v15)
                                                _ = v24
                                            } else {
                                                switch ("" + "a") {
                                                case "a":
                                                    var v25 uint = 104
                                                    _ = v25
                                                case "\n\t":
                                                    switch (("\n\t" + "foo bar") + "foo bar") {
                                                    case "hello":
                                                        if (("hello" + "a") != "a") {
                                                            var v26 int64 = -1
                                                            v27 := GenS6{F0: v16, F1: -1}
                                                            fmt.Println(v19)
                                                            v22 = append(v22, fi11)
                                                            _ = v26
                                                            _ = v27
                                                        } else {
                                                            if ((0.1 * (0.0 - (-2.25 - ((((0.0 + 0.1) + 0.1) + (0.0 - 1.5)) + 0.0)))) > (3.14159 - 0.1)) {
                                                                if true {
                                                                    v12 = (v12 << ((826) & 63))
                                                                    if false {
                                                                        if len(v22) > 0 {
                                                                            v22[0] = -2147483648
                                                                        }
                                                                        if (42 < 9223372036854775807) {
                                                                            var v28 string = "a"
                                                                            _ = len(v4)
                                                                            _ = v28
                                                                        }
                                                                        _ = len(v22)
                                                                        v29 := GenS6{F0: v16, F1: 1}
                                                                        _ = v29
                                                                    } else {
                                                                        if (539 < (v3 ^ v3)) {
                                                                            v4["a"] = "foo bar"
                                                                            v30 := GenS7{F0: "a", F1: fi13}
                                                                            _ = v30
                                                                        } else {
                                                                            delete(v4, ("\n\t" + "foo bar"))
                                                                            var v31 uint16 = 255
                                                                            var v32 uint16 = (v31 * v31)
                                                                            _ = v31
                                                                            _ = v32
                                                                        }
                                                                        v33 := v1[v15]
                                                                        _ = v33
                                                                    }
                                                                    switch ((("\n\t" + "\n\t") + "hello") + ("foo bar" + "")) {
                                                                    case "a":
                                                                        var v34 uint16 = 255
                                                                        _ = v34
                                                                    }
                                                                    if true {
                                                                        fmt.Println(v3, v16)
                                                                        v19 = true
                                                                        switch -277 {
                                                                        case -777:
                                                                            v35 := GenS3{F0: 42}
                                                                            fmt.Println(v16)
                                                                            _ = v35
                                                                        case -2147483648:
                                                                            v16 = (v15 << ((679) & 63))
                                                                        case 32768:
                                                                            fmt.Println(fi13, fi11)
                                                                        }
                                                                    }
                                                                } else {
                                                                    if v19 {
                                                                        var v36 float32 = 0.0
                                                                        _ = len(v1)
                                                                        switch ("a" + ("hello" + ("\n\t" + ""))) {
                                                                        case "a":
                                                                            var v37 uint = 18446744073709551615
                                                                            _ = v37
                                                                        case "foo bar":
                                                                            switch -9223372036854775808 {
                                                                            case 2147483647:
                                                                                switch fi13 {
                                                                                case -1:
                                                                                    v22 = append(v22, (fi13 | fi13))
                                                                                default:
                                                                                    v12 -= (v12 - (^(v12)))
                                                                                }
                                                                            case 127:
                                                                                var v38 uint32 = (v16 & v16)
                                                                                _ = v38
                                                                            case 195:
                                                                                var v39 string = "a"
                                                                                _ = v39
                                                                            }
                                                                        default:
                                                                            var v40 int64 = 1
                                                                            _ = v40
                                                                        }
                                                                        _ = v36
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    case "foo bar":
                                                        v12 = v12
                                                    case "a":
                                                        var v41 int64 = -664
                                                        _ = v41
                                                    }
                                                case "foo bar":
                                                    v2 -= (fi13 << ((65535) & 63))
                                                }
                                                if v19 {
                                                    fmt.Println(v16, v19, v19)
                                                    v3 = (^(v3))
                                                    if v19 {
                                                        switch ("hello" + (("foo bar" + "\n\t") + "a")) {
                                                        case "":
                                                            v16 -= v16
                                                        case "a":
                                                            v42 := GenS6{F0: (v16 & uint32(v3)), F1: -253}
                                                            _ = v42
                                                        case "foo bar":
                                                            v4["\n\t"] = "foo bar"
                                                        }
                                                        _ = len(v4)
                                                    } else {
                                                        v43 := GenS3{F0: 1}
                                                        if (65535 > (^(v16))) {
                                                            var v44 uint64 = 0
                                                            _ = v44
                                                        } else {
                                                            v45 := GenS8{F0: (10 == -128)}
                                                            _ = v45
                                                        }
                                                        v12 -= v12
                                                        _ = v43
                                                    }
                                                }
                                            }
                                            v12 -= (^(v12))
                                            _ = v22
                                        } else {
                                            v3 = v3
                                            var v46 uint = 42
                                            _ = v46
                                        }
                                    }
                                    _ = v19
                                }
                                v47 := GenS9{F0: 42, F1: v2}
                                _ = v16
                                _ = v47
                            }
                            var v48 string = "hello"
                            if (456 != 1) {
                                switch ((v48 + "hello") + v48) {
                                case "\n\t":
                                    delete(v4, "a")
                                }
                                v48 = v48
                                var v49 uint32 = v15
                                _ = v49
                            } else {
                                var v50 int32 = 0
                                _ = v50
                            }
                            v51 := GenS4{F0: 0}
                            _ = v48
                            _ = v51
                        }
                        delete(v1, v15)
                        _ = v15
                    } else {
                        var v52 uint32 = 943
                        _ = v52
                    }
                    var v53 int = v2
                    _ = v53
                }
                v3 = 0
                v3 = 1
                var v54 uint32 = 520
                _ = fi13
                _ = v54
            }
            v3 = v3
            _ = v12
        }
        _ = fi11
    }
    v3 += v3
    fmt.Println(v2, v3)
    switch v2 {
    case -612:
        v3 = 1
    case 42:
        v55 := []int{v2, (v2 / ((-2147483648) | 1))}
        _ = v55
    case -9223372036854775808:
        for fi56 := 0; fi56 < 13; fi56++ {
            v57 := []int{0, v2, (-(fi56)), fi56}
            for fi58 := 0; fi58 < 10; fi58++ {
                v59 := GenS10{F0: 42, F1: 65535}
                _ = fi58
                _ = v59
            }
            _ = fi56
            _ = v57
        }
    default:
        switch (-(v2)) {
        case 32767:
            v60 := GenS4{F0: 2}
            _ = v60
        case 2147483647:
            var v61 int16 = -1
            _ = v61
        case 0:
            for fi62 := 0; fi62 < 8; fi62++ {
                if ((false && false) || (false && true)) {
                    var v63 int16 = 1000
                    _ = v63
                } else {
                    if false {
                        v3 = v3
                        switch "\n\t" {
                        case "foo bar":
                            var v64 uint64 = (v3 % ((v3) | 1))
                            _ = v64
                        case "hello":
                            for fi65 := 0; fi65 < 11; fi65++ {
                                v2 = -1
                                switch 1 {
                                case 9223372036854775807:
                                    v3 += v3
                                }
                                switch "" {
                                case "\n\t":
                                    v66 := GenS6{F0: 136, F1: -95}
                                    fmt.Println(v3, fi65)
                                    _ = v66
                                case "hello":
                                    v67 := []int{9223372036854775807, 4294967295}
                                    _ = v67
                                }
                                _ = fi65
                            }
                        }
                    } else {
                        for fi68 := 0; fi68 < 10; fi68++ {
                            var v69 int64 = -360
                            _ = fi68
                            _ = v69
                        }
                    }
                    if (42 != 127) {
                        var v70 uint = 42
                        _ = v70
                    }
                }
                _ = fi62
            }
        }
    }
    fmt.Println(v2, v3)
    v71 := GenS3{F0: 42}
    v3 += v3
    if true {
        if false {
            if ((65535 >= 2147483648) || false) {
                var v72 int32 = 650
                _ = v72
            } else {
                if true {
                    v3 = (v3 & 291)
                    v73 := GenS11{F0: 127, F1: (true && (1.5 < 0.0)), F2: 843}
                    _ = v73
                }
                var v74 uint16 = 65535
                if true {
                    var v75 float64 = (-2.25 + 0.0)
                    v76 := GenS6{F0: 630, F1: 0}
                    if (!true) {
                        v2 = v2
                    }
                    _ = v75
                    _ = v76
                } else {
                    var v77 int = v2
                    var v78 uint64 = (v3 + v3)
                    _ = v77
                    _ = v78
                }
                _ = v74
            }
            var v79 uint8 = 26
            if true {
                var v80 uint8 = (v79 | 0)
                _ = v80
            } else {
                var v81 bool = false
                v79 = v79
                switch (v2 + v2) {
                case 0:
                    v82 := []int{1, v2, 42, v2}
                    _ = v82
                default:
                    var v83 uint64 = (v3 | (v3 >> ((0) & 63)))
                    _ = v83
                }
                _ = v81
            }
            _ = v79
        }
    }
    fmt.Println(v2, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v71
    _ = v8
}
