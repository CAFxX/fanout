package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 string; F2 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 bool; F2 int32 }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 uint16; F1 uint32 }
type GenS7 struct { F0 int; F1 int }

func main() {
    var v1 uint = 18446744073709551615
    var v2 int = 9223372036854775807
    var v3 float32 = (0.0 * ((1.5 / 1.0) * (-2.25 + -2.25)))
    v4 := []string{"hello", "\n\t"}
    var v5 float32 = v3
    if (255 >= 255) {
        switch v2 {
        case 128:
            for fi6 := 0; fi6 < 10; fi6++ {
                v3 = v3
                _ = fi6
            }
        default:
            switch v2 {
            case -451:
                fmt.Println(v2, v2)
            case 140:
                v4 = append(v4, ("" + "hello"))
            case 163:
                var v7 uint8 = 35
                _ = v7
            default:
                fmt.Println(v3, v5)
            }
        }
    } else {
        var v8 int64 = 42
        _ = v8
    }
    v1 = (^(v1))
    var v9 int64 = 1
    if (false && ((false || (-2.25 > (((1.5 / 3.14159) + 0.0) - float64(v3)))) || true)) {
        fmt.Println(v3, v5, v5)
        var v10 string = ("hello" + "hello")
        _ = v10
    }
    v11 := GenS1{F0: (v1 << (((v1 << ((v1) & 63))) & 63))}
    if false {
        var v12 int32 = 2147483647
        if len(v4) > 0 {
            v4[0] = "hello"
        }
        _ = v12
    }
    if false {
        if ((v2 - v2) <= (^(v2))) {
            v2 -= (v2 << (((v1 - (^(v1)))) & 63))
            switch v2 {
            case 32767:
                if ((v2 - (^(v2))) != v2) {
                    v1 = (v1 + (v1 & (v1 / ((101) | 1))))
                } else {
                    var v13 uint = v1
                    v1 += 18446744073709551615
                    v5 = (v3 / (0.1 * (((float32((0.0 / (0.1 * 3.14159))) - ((float32(0.1) * -2.25) * v3)) + (1e10 * 1.5)) * v5)))
                    _ = v13
                }
            case 42:
                for fi14 := 0; fi14 < 13; fi14++ {
                    switch ("a" + ("a" + ("hello" + ("foo bar" + "hello")))) {
                    case "a":
                        v1 = v1
                    case "hello":
                        v5 = -2.25
                    default:
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi14
                }
            }
            _ = len(v4)
        }
    } else {
        v15 := GenS2{F0: 874, F1: "foo bar", F2: -237}
        if (1 > 1) {
            var v16 int = (v2 >> ((1) & 63))
            _ = v16
        } else {
            v2 -= 0
        }
        fmt.Println(v2, v3, v2)
        _ = v15
    }
    switch (v2 % (((v2 >> (((v1 * v1)) & 63))) | 1)) {
    case 128:
        var v17 int8 = -108
        _ = v17
    case 32768:
        fmt.Println(v9, v5)
    case -9223372036854775808:
        if (((v1 ^ (v1 - 65535)) >= 42) && false) {
            v18 := GenS1{F0: (v1 & (v1 | (v1 ^ (v1 >> ((65535) & 63)))))}
            v19 := GenS3{F0: v9}
            v20 := GenS4{F0: 0, F1: ((57 > 42) && (!false)), F2: 1}
            _ = v18
            _ = v19
            _ = v20
        }
    default:
        if (((^(v2)) >= v2) || true) {
            v1 = v1
            v21 := GenS3{F0: v9}
            var v22 bool = (!false)
            _ = v21
            _ = v22
        } else {
            fmt.Println(v1)
        }
    }
    var v23 int8 = 111
    switch "\n\t" {
    case "a":
        if true {
            v5 -= float32(((0.0 / 1.5) * (float64((0.1 + 1e10)) + 0.1)))
            for fi24 := 0; fi24 < 19; fi24++ {
                var v25 bool = ((v1 | v1) < (^(v1)))
                _ = fi24
                _ = v25
            }
        } else {
            v1 = v1
            v26 := GenS1{F0: v1}
            switch v2 {
            case -589:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 42:
                switch "" {
                case "":
                    var v27 int8 = (v23 | -1)
                    _ = v27
                case "foo bar":
                    if true {
                        v9 -= v9
                        for fi28 := 0; fi28 < 7; fi28++ {
                            if (128 > 128) {
                                switch "hello" {
                                case "a":
                                    if false {
                                        var v29 bool = (false && ("foo bar" < "a"))
                                        _ = v29
                                    } else {
                                        switch (("hello" + "hello") + "hello") {
                                        case "\n\t":
                                            if (0.1 != 3.14159) {
                                                var v30 uint = v1
                                                _ = v30
                                            }
                                        case "":
                                            var v31 bool = (false && true)
                                            _ = v31
                                        case "hello":
                                            v4 = append(v4, "")
                                        default:
                                            if false {
                                                var v32 string = (("foo bar" + "") + "\n\t")
                                                _ = v32
                                            }
                                        }
                                        fmt.Println(v5, fi28)
                                        _ = len(v4)
                                    }
                                case "foo bar":
                                    var v33 bool = false
                                    _ = v33
                                case "hello":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                default:
                                    var v34 float32 = v3
                                    _ = v34
                                }
                                var v35 string = "foo bar"
                                _ = v35
                            } else {
                                v23 += -1
                            }
                            for fi36 := 0; fi36 < 17; fi36++ {
                                if true {
                                    if len(v4) > 0 {
                                        v4[0] = ""
                                    }
                                } else {
                                    _ = len(v4)
                                    if true {
                                        switch "hello" {
                                        case "foo bar":
                                            if ("" != (("hello" + "a") + "")) {
                                                v2 = 128
                                                var v37 uint16 = 65535
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v38 := GenS3{F0: -9223372036854775808}
                                                    _ = v38
                                                default:
                                                    v3 = v3
                                                }
                                                _ = v37
                                            }
                                        case "hello":
                                            v9 -= (v9 & 9223372036854775807)
                                        case "\n\t":
                                            v39 := GenS5{F0: -701, F1: false}
                                            _ = v39
                                        }
                                    } else {
                                        if (4294967295 <= 112) {
                                            v40 := GenS3{F0: (v9 | -857)}
                                            v1 -= (v1 << (((v1 * (v1 | (v1 / ((153) | 1))))) & 63))
                                            var v41 int16 = int16(v9)
                                            _ = v40
                                            _ = v41
                                        } else {
                                            v3 = float32(0.1)
                                            switch (("hello" + "hello") + ("hello" + "hello")) {
                                            case "foo bar":
                                                switch "hello" {
                                                case "":
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        var v42 int8 = (v23 - (v23 * (-(v23))))
                                                        _ = v42
                                                    case "":
                                                        switch fi36 {
                                                        case -1:
                                                            var v43 int32 = 2147483647
                                                            _ = v43
                                                        }
                                                    default:
                                                        v3 = ((v5 + (v3 * float32(-2.25))) * v3)
                                                    }
                                                    var v44 bool = (-1 >= 1000)
                                                    _ = v44
                                                case "hello":
                                                    v23 -= (v23 & v23)
                                                case "a":
                                                    if true {
                                                        switch ("foo bar" + "\n\t") {
                                                        case "\n\t":
                                                            var v45 string = "a"
                                                            _ = v45
                                                        case "":
                                                            v46 := GenS5{F0: (fi28 << (((v1 + 255)) & 63)), F1: true}
                                                            var v47 int64 = (v9 ^ int64(v1))
                                                            _ = v46
                                                            _ = v47
                                                        case "hello":
                                                            var v48 uint64 = 18446744073709551615
                                                            _ = v48
                                                        default:
                                                            var v49 string = "\n\t"
                                                            var v50 int = v2
                                                            _ = v49
                                                            _ = v50
                                                        }
                                                        switch v2 {
                                                        case 127:
                                                            v51 := GenS3{F0: (v9 & (v9 + v9))}
                                                            if (!((v1 | v1) <= (v1 >> ((4294967295) & 63)))) {
                                                                fmt.Println(v1)
                                                                var v52 int = 128
                                                                var v53 int16 = 320
                                                                _ = v52
                                                                _ = v53
                                                            }
                                                            _ = v51
                                                        case 32768:
                                                            fmt.Println(fi36)
                                                            var v54 uint32 = 0
                                                            if len(v4) > 0 {
                                                                v4[0] = "\n\t"
                                                            }
                                                            _ = v54
                                                        case 4294967295:
                                                            var v55 string = ""
                                                            _ = v55
                                                        default:
                                                            v4 = append(v4, (("" + "a") + "a"))
                                                        }
                                                        v23 = (v23 & v23)
                                                    } else {
                                                        v56 := GenS5{F0: fi36, F1: (32767 == 42)}
                                                        v57 := GenS6{F0: uint16(v1), F1: 544}
                                                        _ = v56
                                                        _ = v57
                                                    }
                                                default:
                                                    var v58 bool = (!(42 == 55))
                                                    _ = v58
                                                }
                                            case "hello":
                                                fmt.Println(v5, v9, v3)
                                            }
                                        }
                                        var v59 uint64 = 521
                                        switch ((("a" + "") + "\n\t") + ("foo bar" + ("\n\t" + ("" + "a")))) {
                                        case "":
                                            fmt.Println(fi36, fi28)
                                        case "a":
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        case "\n\t":
                                            v1 = (v1 / ((v1) | 1))
                                        default:
                                            fmt.Println(v23, v23, v9)
                                        }
                                        if (v1 <= v1) {
                                            if (false || true) {
                                                var v60 uint16 = 0
                                                v2 -= (v2 / ((0) | 1))
                                                if true {
                                                    var v61 int16 = 36
                                                    v2 += (v2 << ((v1) & 63))
                                                    v62 := GenS2{F0: -2147483648, F1: ("\n\t" + "foo bar"), F2: -323}
                                                    _ = len(v4)
                                                    _ = v61
                                                    _ = v62
                                                }
                                                _ = v60
                                            }
                                        }
                                        if (1 < 2147483647) {
                                            v23 = v23
                                            var v63 uint32 = 4294967295
                                            v9 -= (^(v9))
                                            _ = v63
                                        } else {
                                            switch v2 {
                                            case 128:
                                                _ = len(v4)
                                            default:
                                                v64 := GenS5{F0: -646, F1: (v59 != v59)}
                                                _ = v64
                                            }
                                        }
                                        _ = v59
                                    }
                                    switch (-(fi36)) {
                                    case 557:
                                        fmt.Println(v2)
                                    case -178:
                                        v65 := GenS5{F0: int(v1), F1: true}
                                        _ = v65
                                    case -309:
                                        _ = len(v4)
                                    }
                                }
                                fmt.Println(v5)
                                _ = fi36
                            }
                            v4 = append(v4, "foo bar")
                            _ = fi28
                        }
                        var v66 uint32 = 908
                        _ = v66
                    } else {
                        v9 += (-(v9))
                    }
                default:
                    _ = len(v4)
                }
            case 1:
                if false {
                    fmt.Println(v5, v9, v2)
                }
            }
            _ = v26
        }
    case "hello":
        v67 := GenS4{F0: (v2 >> ((v1) & 63)), F1: true, F2: 2147483647}
        _ = v67
    case "":
        if ((3.14159 != -2.25) && true) {
            v9 = v9
            switch "\n\t" {
            case "\n\t":
                switch "foo bar" {
                case "a":
                    if (!(v9 < 0)) {
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                        for fi68 := 0; fi68 < 6; fi68++ {
                            v2 = 128
                            v69 := GenS7{F0: v2, F1: -9223372036854775808}
                            _ = fi68
                            _ = v69
                        }
                    } else {
                        var v70 string = "a"
                        _ = v70
                    }
                case "":
                    v1 = (^(v1))
                default:
                    fmt.Println(v23, v2)
                }
            case "foo bar":
                v71 := GenS7{F0: (v2 >> ((v1) & 63)), F1: v2}
                _ = v71
            case "a":
                fmt.Println(v3, v3, v3)
            }
        } else {
            v72 := GenS1{F0: v1}
            if len(v4) > 0 {
                v4[0] = ((("\n\t" + "a") + "foo bar") + "hello")
            }
            _ = v72
        }
    default:
        var v73 int32 = 42
        _ = v73
    }
    v23 -= 52
    for fi74 := 0; fi74 < 17; fi74++ {
        var v75 uint8 = 168
        _ = fi74
        _ = v75
    }
    v3 += v5
    var v76 uint = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v76:", v76)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v76
    _ = v9
}
