package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 int32; F2 uint }
type GenS5 struct { F0 int32; F1 uint8; F2 uint16 }
type GenS6 struct { F0 uint32; F1 uint; F2 int16 }
type GenS7 struct { F0 int32; F1 bool; F2 uint32 }

func main() {
    var v1 float64 = 1.5
    var v2 uint8 = 62
    var v3 int = 9223372036854775807
    for fi4 := 0; fi4 < 5; fi4++ {
        v3 -= (fi4 ^ 32768)
        v2 = 0
        v5 := GenS1{F0: 42, F1: 927, F2: -416}
        _ = fi4
        _ = v5
    }
    v6 := []int{(v3 * -2147483648), v3, (v3 << ((255) & 63))}
    var v7 float32 = float32(((((v1 * ((v1 + (v1 - -2.25)) + v1)) + v1) / (v1 - v1)) + ((v1 + v1) - 1.5)))
    var v8 float64 = v1
    v9 := GenS2{F0: 255, F1: 361}
    for fi10 := 0; fi10 < 7; fi10++ {
        v1 -= 0.0
        var v11 uint8 = v2
        _ = fi10
        _ = v11
    }
    var v12 string = "a"
    _ = len(v6)
    if (540 <= -134) {
        var v13 int32 = 0
        fmt.Println(v13)
        if ((1.5 * v7) <= 0.0) {
            var v14 bool = true
            if true {
                v7 -= (v7 * v7)
            } else {
                v15 := GenS2{F0: (v2 - v2), F1: 4294967295}
                if v14 {
                    v2 = (v2 ^ 2)
                } else {
                    var v16 uint32 = 42
                    v12 = ((v12 + "a") + "")
                    if len(v6) > 0 {
                        v6[0] = (v3 % (((v3 << ((18446744073709551615) & 63))) | 1))
                    }
                    _ = v16
                }
                for fi17 := 0; fi17 < 9; fi17++ {
                    fmt.Println(v1, v8, fi17)
                    v18 := GenS1{F0: (v13 % ((v13) | 1)), F1: 543, F2: 0}
                    _ = fi17
                    _ = v18
                }
                _ = v15
            }
            var v19 uint16 = 677
            _ = v14
            _ = v19
        }
        _ = v13
    } else {
        var v20 uint16 = 255
        v12 = ("a" + "\n\t")
        _ = v20
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        v7 = (3.14159 + v7)
        for fi22 := 0; fi22 < 16; fi22++ {
            switch (fi21 - 116) {
            case -324:
                if (2147483647 == -464) {
                    v6 = append(v6, 4294967295)
                }
            case -464:
                var v23 int64 = 42
                _ = v23
            }
            _ = fi22
        }
        _ = fi21
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        var v25 int16 = -1
        if (666 >= 268) {
            switch 446 {
            case 0:
                for fi26 := 0; fi26 < 11; fi26++ {
                    switch fi24 {
                    case 2147483647:
                        if true {
                            if false {
                                v27 := GenS1{F0: -810, F1: 953, F2: v25}
                                v7 = float32(v8)
                                var v28 uint = 0
                                switch "foo bar" {
                                case "a":
                                    v7 -= v7
                                    v29 := GenS2{F0: (v2 << ((v28) & 63)), F1: v28}
                                    _ = v29
                                }
                                fmt.Println(v1, v12, fi26)
                                switch (fi24 ^ 1) {
                                case -9223372036854775808:
                                    var v30 int32 = 2147483647
                                    _ = v30
                                }
                                _ = v27
                                _ = v28
                            }
                            if (v1 <= (0.0 - v8)) {
                                var v31 int = fi24
                                if len(v6) > 0 {
                                    v6[0] = 42
                                }
                                _ = v31
                            } else {
                                var v32 int64 = -1
                                switch v12 {
                                case "\n\t":
                                    v8 -= (3.14159 * 3.14159)
                                case "hello":
                                    v7 -= (v7 - (0.1 - -2.25))
                                    if false {
                                        var v33 int64 = (v32 / ((v32) | 1))
                                        if len(v6) > 0 {
                                            v6[0] = fi26
                                        }
                                        if true {
                                            var v34 uint16 = 794
                                            v12 = (v12 + v12)
                                            _ = v34
                                        }
                                        _ = v33
                                    }
                                    fmt.Println(v2)
                                case "":
                                    v25 += 0
                                }
                                _ = len(v6)
                                var v35 int8 = 76
                                _ = v32
                                _ = v35
                            }
                            v3 = 32767
                        } else {
                            var v36 uint8 = 238
                            v37 := GenS3{F0: 18446744073709551615}
                            if (!(648 == 2)) {
                                v7 = 3.14159
                            } else {
                                switch v12 {
                                case "":
                                    var v38 int8 = 0
                                    _ = v38
                                case "hello":
                                    v8 = v8
                                }
                                if ((18446744073709551615 >= 805) || true) {
                                    var v39 uint16 = 542
                                    var v40 uint32 = 1
                                    switch "\n\t" {
                                    case "":
                                        var v41 uint32 = (v40 + (v40 | 684))
                                        v41 = v41
                                        _ = v41
                                    }
                                    if true {
                                        v2 += v36
                                        if (!(42 < 4294967295)) {
                                            if (-1 <= 33) {
                                                if (((!true) || (32767 < (v39 >> ((255) & 63)))) || (true || false)) {
                                                    v40 += (v40 << ((42) & 63))
                                                    v6 = append(v6, (fi26 >> ((1) & 63)))
                                                    v42 := GenS2{F0: (^(v2)), F1: uint(v40)}
                                                    _ = v42
                                                }
                                                if (v39 >= v39) {
                                                    var v43 uint8 = v2
                                                    v6 = append(v6, -473)
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        v44 := GenS3{F0: 4294967295}
                                                        _ = v44
                                                    case "":
                                                        v45 := GenS2{F0: (v2 | 255), F1: 65535}
                                                        _ = v45
                                                    default:
                                                        v46 := GenS3{F0: 832}
                                                        _ = v46
                                                    }
                                                    var v47 bool = false
                                                    fmt.Println(v25, v43)
                                                    _ = v43
                                                    _ = v47
                                                }
                                            } else {
                                                v3 = fi24
                                            }
                                        } else {
                                            var v48 int = (fi26 + fi26)
                                            _ = v48
                                        }
                                    } else {
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        v25 = (v25 * (v25 >> ((596) & 63)))
                                    }
                                    _ = v39
                                    _ = v40
                                } else {
                                    v8 -= float64(-2.25)
                                    if true {
                                        var v49 uint8 = (v2 - 227)
                                        var v50 int64 = -9223372036854775808
                                        _ = v49
                                        _ = v50
                                    } else {
                                        v25 += v25
                                        if len(v6) > 0 {
                                            v6[0] = fi24
                                        }
                                        var v51 string = v12
                                        _ = v51
                                    }
                                    switch v12 {
                                    case "":
                                        v8 = (((((v8 / (1.5 - ((v1 - 1.5) * 0.1))) / float64(v7)) + 1.5) - (float64((v7 / -2.25)) - v8)) + (1e10 - (-2.25 + v8)))
                                    case "a":
                                        _ = len(v6)
                                    case "hello":
                                        v52 := GenS3{F0: uint(v36)}
                                        _ = v52
                                    }
                                }
                            }
                            _ = v36
                            _ = v37
                        }
                    case -9223372036854775808:
                        v53 := GenS4{F0: 529, F1: -1, F2: 0}
                        _ = v53
                    }
                    v7 = ((((v7 - 0.0) - (float32(1.5) + v7)) - (((0.0 - (((v7 * v7) * 1e10) + v7)) - 0.0) + (v7 + (v7 * (((float32(v1) * ((-2.25 + (1e10 - (v7 * v7))) * 0.1)) + v7) - (v7 / (v7 - v7))))))) * 0.1)
                    if (0 >= -9223372036854775808) {
                        v2 -= 1
                        v3 = fi24
                        v7 -= (1.5 - 0.1)
                    } else {
                        var v54 uint16 = 42
                        var v55 int16 = (v25 / ((v25) | 1))
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        if (!(0 <= 97)) {
                            v56 := GenS5{F0: 42, F1: 0, F2: 32767}
                            if false {
                                fmt.Println(v54, v54)
                                var v57 int32 = 2147483647
                                _ = v57
                            } else {
                                switch 32768 {
                                case 0:
                                    v6 = append(v6, (v3 | fi24))
                                case 9223372036854775807:
                                    v7 = v7
                                default:
                                    if (0 <= 0) {
                                        v54 = (^(v54))
                                        var v58 int16 = -166
                                        _ = v58
                                    }
                                }
                            }
                            _ = v56
                        } else {
                            switch (v3 & (v3 << ((627) & 63))) {
                            case -9223372036854775808:
                                switch (^(fi26)) {
                                case 32767:
                                    var v59 int8 = -1
                                    _ = v59
                                }
                            case -2147483648:
                                if (32767 >= (fi26 - fi26)) {
                                    var v60 uint8 = (^(v2))
                                    v61 := GenS6{F0: uint32(v54), F1: 42, F2: (v25 / ((v25) | 1))}
                                    _ = len(v6)
                                    _ = v60
                                    _ = v61
                                }
                            }
                            fmt.Println(v7, v25, v2)
                        }
                        v62 := GenS4{F0: 0, F1: 42, F2: 0}
                        _ = v54
                        _ = v55
                        _ = v62
                    }
                    _ = fi26
                }
            }
            v1 = ((1e10 - v1) + 0.0)
            if ((!true) || true) {
                var v63 bool = true
                _ = v63
            }
        }
        _ = fi24
        _ = v25
    }
    switch "a" {
    case "hello":
        v64 := GenS7{F0: 0, F1: false, F2: 2147483647}
        _ = v64
    case "a":
        for fi65 := 0; fi65 < 1; fi65++ {
            if false {
                for fi66 := 0; fi66 < 15; fi66++ {
                    var v67 int16 = 42
                    v68 := GenS4{F0: v67, F1: -2147483648, F2: 0}
                    v67 = (v67 * v67)
                    _ = fi66
                    _ = v67
                    _ = v68
                }
            }
            switch (v12 + "\n\t") {
            case "":
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case "foo bar":
                if len(v6) > 0 {
                    v6[0] = v3
                }
            }
            if (fi65 > (fi65 / (((v3 >> ((42) & 63))) | 1))) {
                v1 += ((float64(float32((1e10 * v1))) - v1) - (float64((float32(v8) + 0.0)) / ((float64(v7) / v1) * v8)))
                switch ("hello" + (v12 + "hello")) {
                case "a":
                    switch 821 {
                    case 9223372036854775807:
                        v12 = ("foo bar" + (v12 + (v12 + "")))
                    }
                case "":
                    var v69 uint32 = 65535
                    _ = v69
                }
                var v70 uint = 1
                _ = v70
            } else {
                v12 = (v12 + v12)
                if len(v6) > 0 {
                    v6[0] = (fi65 / ((1) | 1))
                }
                var v71 uint = 42
                _ = v71
            }
            _ = fi65
        }
    case "foo bar":
        if (-69 >= 42) {
            var v72 uint8 = 1
            fmt.Println(v3, v7)
            _ = v72
        } else {
            v8 = float64((((0.0 - v7) - v7) * (v7 - float32(v8))))
            v7 += (float32(v1) - (0.1 * v7))
        }
    default:
        switch v12 {
        case "":
            v8 = -2.25
        case "\n\t":
            var v73 int16 = -1
            _ = v73
        case "a":
            fmt.Println(v2)
        }
    }
    switch (v12 + ((("foo bar" + v12) + "foo bar") + v12)) {
    case "\n\t":
        var v74 int16 = 0
        _ = v74
    }
    v1 += (3.14159 - v8)
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
