package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 int32; F2 int64 }
type GenS4 struct { F0 uint; F1 string; F2 bool }
type GenS5 struct { F0 bool; F1 uint; F2 int }
type GenS6 struct { F0 uint16; F1 int64; F2 bool }
type GenS7 struct { F0 int32; F1 uint8; F2 int8 }
type GenS8 struct { F0 bool; F1 int; F2 int64 }
func genFn0() bool {
    v5 := []int{127, 4294967295}
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 string = ("a" + "")
        _ = fi6
        _ = v7
    }
    var v8 bool = false
    _ = v5
    _ = v8
    return v8
}
func genFn1() int64 {
    v9 := GenS1{F0: 475, F1: -2147483648, F2: 255}
    if false {
        var v10 uint16 = 1
        _ = v10
    } else {
        var v11 uint = 2
        v11 += v11
        _ = v11
    }
    _ = v9
    return -9223372036854775808
}

func main() {
    v1 := []string{"\n\t", "", "", (("hello" + "\n\t") + ("a" + ""))}
    var v2 int16 = -355
    var v3 uint16 = 2
    var v4 float64 = (3.14159 - (0.1 * 0.0))
    var v12 int32 = -2147483648
    switch "a" {
    case "a":
        if (-2.25 < 0.1) {
            v13 := GenS1{F0: int64(v3), F1: -291, F2: 2}
            v1 = append(v1, "foo bar")
            _ = v13
        }
    default:
        for fi14 := 0; fi14 < 8; fi14++ {
            for fi15 := 0; fi15 < 17; fi15++ {
                var v16 int8 = 42
                switch fi15 {
                case 4294967295:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if (v3 <= v3) {
                        switch (fi14 >> ((410) & 63)) {
                        case 42:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case 4294967295:
                            if (750 >= (fi14 << ((626) & 63))) {
                                if (1 != 4294967295) {
                                    fmt.Println(v2)
                                    var v17 int = (fi15 << ((2) & 63))
                                    v16 -= (v16 | 100)
                                    _ = v17
                                }
                            }
                        }
                        v2 = -124
                    } else {
                        var v18 uint8 = 155
                        switch -1 {
                        case 127:
                            var v19 uint16 = v3
                            _ = v19
                        case 963:
                            var v20 string = ""
                            _ = v20
                        case 32767:
                            var v21 string = (("\n\t" + "") + "\n\t")
                            _ = v21
                        default:
                            if false {
                                var v22 int64 = -945
                                _ = v22
                            }
                        }
                        _ = v18
                    }
                case 127:
                    var v23 float64 = (3.14159 - 1e10)
                    _ = v23
                }
                v2 += v2
                switch (fi15 / ((42) | 1)) {
                case -731:
                    v16 -= (v16 - -1)
                }
                _ = fi15
                _ = v16
            }
            v2 = (v2 >> ((1) & 63))
            _ = fi14
        }
    }
    v3 -= (v3 | 326)
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        _ = len(v1)
    case "":
        for fi24 := 0; fi24 < 14; fi24++ {
            switch (fi24 % (((fi24 | int(v3))) | 1)) {
            case 0:
                fmt.Println(fi24)
            }
            v25 := GenS1{F0: 0, F1: v12, F2: 600}
            var v26 uint64 = 511
            _ = fi24
            _ = v25
            _ = v26
        }
    case "hello":
        v27 := GenS2{F0: 0}
        _ = v27
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        v2 -= (v2 | v2)
        if (!(0 > (fi28 + 4294967295))) {
            fmt.Println(v2, v2, fi28)
            v12 = (v12 ^ 0)
            v29 := GenS3{F0: -1, F1: v12, F2: -9223372036854775808}
            _ = v29
        } else {
            if false {
                fmt.Println(v3)
                switch "a" {
                case "foo bar":
                    fmt.Println(v3, v12)
                case "a":
                    v12 -= 42
                case "hello":
                    switch "\n\t" {
                    case "a":
                        var v30 bool = (false && false)
                        _ = v30
                    default:
                        switch fi28 {
                        case 9223372036854775807:
                            v31 := GenS2{F0: 9223372036854775807}
                            _ = v31
                        default:
                            v2 -= v2
                        }
                    }
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                v32 := GenS4{F0: 1, F1: "\n\t", F2: (!false)}
                _ = v32
            }
            if len(v1) > 0 {
                v1[0] = "a"
            }
            if (!(!true)) {
                var v33 bool = (42 > 127)
                for fi34 := 0; fi34 < 1; fi34++ {
                    fmt.Println(fi28)
                    switch (fi28 - (-(fi28))) {
                    case 2147483647:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case 32767:
                        fmt.Println(v3)
                    default:
                        if (("" + (("\n\t" + ("a" + "foo bar")) + (((("" + "hello") + (("\n\t" + ("hello" + "a")) + "")) + "") + ("a" + "a")))) >= ("" + "\n\t")) {
                            fmt.Println(v12, v33)
                            v35 := GenS5{F0: (v33 || (0 <= 42)), F1: 982, F2: (^(fi34))}
                            _ = v35
                        }
                        fmt.Println(fi34, v3)
                        var v36 uint64 = 613
                        _ = v36
                    }
                    switch (fi34 - (fi28 | fi34)) {
                    case -226:
                        switch (^(fi34)) {
                        case 565:
                            v3 = (v3 << ((4) & 63))
                        case -2147483648:
                            v37 := GenS6{F0: (v3 + (v3 + (v3 + 141))), F1: 42, F2: v33}
                            _ = v37
                        }
                    default:
                        if v33 {
                            switch "" {
                            case "\n\t":
                                if (0 >= -2147483648) {
                                    switch 32767 {
                                    case 32767:
                                        var v38 uint = 42
                                        _ = v38
                                    }
                                    if false {
                                        if v33 {
                                            v12 -= (v12 + v12)
                                            v3 -= (v3 + (v3 + 2))
                                        } else {
                                            v33 = (-40 == int64(v12))
                                        }
                                        v33 = (!v33)
                                        if (!(79 <= 56)) {
                                            fmt.Println(fi28, v2, v12)
                                            var v39 uint16 = (v3 ^ (v3 << ((1) & 63)))
                                            if (-2.25 > (0.0 + (0.1 / float32(3.14159)))) {
                                                v40 := GenS3{F0: -1, F1: -22, F2: -1}
                                                if len(v1) > 0 {
                                                    v1[0] = ("hello" + "a")
                                                }
                                                _ = v40
                                            }
                                            _ = v39
                                        }
                                        fmt.Println(v4, fi28)
                                        fmt.Println(fi28, v12)
                                    }
                                    if (4294967295 >= 65535) {
                                        v41 := GenS6{F0: (^(v3)), F1: -9223372036854775808, F2: false}
                                        if (3.14159 <= 0.0) {
                                            v12 = -720
                                            switch "foo bar" {
                                            case "foo bar":
                                                v42 := GenS1{F0: -1, F1: 2147483647, F2: 353}
                                                _ = v42
                                            case "hello":
                                                var v43 float64 = (3.14159 - float64(-2.25))
                                                _ = v43
                                            case "a":
                                                v44 := GenS5{F0: (v3 >= (^(v3))), F1: 34, F2: fi34}
                                                _ = v44
                                            default:
                                                if false {
                                                    var v45 float32 = ((3.14159 + 0.1) * ((1e10 + 3.14159) / 3.14159))
                                                    _ = v45
                                                }
                                            }
                                        } else {
                                            fmt.Println(v4, v3, v33)
                                            v1 = append(v1, ("a" + "hello"))
                                            var v46 uint8 = 127
                                            _ = v46
                                        }
                                        _ = v41
                                    }
                                }
                            }
                        }
                    }
                    _ = fi34
                }
                _ = v33
            }
        }
        _ = fi28
    }
    var v47 uint8 = 127
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v48 uint8 = v47
    for fi49 := 0; fi49 < 9; fi49++ {
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        fmt.Println(v12, v3, fi49)
        _ = fi49
    }
    if false {
        v47 = v48
    } else {
        if false {
            if (979 == -2147483648) {
                switch 501 {
                case 9223372036854775807:
                    var v50 bool = true
                    _ = v50
                case -516:
                    var v51 int16 = (v2 ^ (v2 + v2))
                    _ = v51
                default:
                    var v52 uint8 = v47
                    _ = v52
                }
            }
            var v53 uint = 65535
            _ = v53
        } else {
            v12 += (v12 * (v12 ^ -1))
            var v54 float64 = float64((((((0.0 + (3.14159 * float32(v4))) / (0.1 * -2.25)) - (0.0 - -2.25)) - (((0.1 / ((1.5 + 1e10) + -2.25)) - 0.1) / -2.25)) + 1.5))
            _ = v54
        }
    }
    if ((v48 << ((1) & 63)) != (v48 << ((65535) & 63))) {
        var v55 bool = false
        fmt.Println(v2)
        v2 += int16(v3)
        _ = v55
    }
    var v56 uint16 = v3
    if (9223372036854775807 == 989) {
        for fi57 := 0; fi57 < 17; fi57++ {
            if ((1e10 / 3.14159) >= (-2.25 + 1e10)) {
                var v58 int32 = v12
                _ = v58
            }
            _ = fi57
        }
        if (!(!(238 <= 600))) {
            switch "a" {
            case "\n\t":
                if true {
                    if false {
                        v59 := GenS7{F0: v12, F1: (v47 >> ((204) & 63)), F2: 1}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v59
                    } else {
                        fmt.Println(v12, v56, v2)
                    }
                    var v60 int32 = 42
                    var v61 int32 = -497
                    _ = v60
                    _ = v61
                }
            default:
                _ = len(v1)
            }
            var v62 float64 = v4
            if true {
                fmt.Println(v62)
                for fi63 := 0; fi63 < 5; fi63++ {
                    var v64 int = (fi63 >> ((255) & 63))
                    _ = fi63
                    _ = v64
                }
                var v65 float64 = -2.25
                _ = v65
            } else {
                if true {
                    for fi66 := 0; fi66 < 5; fi66++ {
                        v2 = (v2 / ((v2) | 1))
                        var v67 uint16 = v56
                        var v68 int64 = 42
                        _ = fi66
                        _ = v67
                        _ = v68
                    }
                }
            }
            _ = v62
        }
        v1 = append(v1, "foo bar")
    } else {
        v3 += (v3 | (v3 / ((12) | 1)))
        v48 -= v47
        for fi69 := 0; fi69 < 5; fi69++ {
            var v70 string = ""
            var v71 uint = 339
            v71 = v71
            _ = fi69
            _ = v70
            _ = v71
        }
    }
    v72 := GenS8{F0: (-409 <= 9223372036854775807), F1: 32768, F2: -1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v56
    _ = v72
}
