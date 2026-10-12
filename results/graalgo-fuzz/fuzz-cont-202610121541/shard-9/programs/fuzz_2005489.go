package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint32; F1 uint; F2 int }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 int16; F1 int32 }
type GenS7 struct { F0 int64; F1 uint }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"a", "foo bar"}
    v3 := make(map[int]uint32)
    var v4 int16 = -1
    var v5 uint8 = 255
    v6 := make(map[string]string)
    fmt.Println(v5, v4)
    fmt.Println(v4)
    if true {
        if false {
            if len(v2) > 0 {
                v2[0] = "hello"
            }
            for fi7 := 0; fi7 < 7; fi7++ {
                for fi8 := 0; fi8 < 9; fi8++ {
                    var v9 int64 = -9223372036854775808
                    fmt.Println(v5)
                    delete(v1, 2147483647)
                    if (0.0 < 0.1) {
                        switch fi7 {
                        case 1:
                            fmt.Println(v9)
                        case 32768:
                            if true {
                                var v10 string = "foo bar"
                                _ = v10
                            } else {
                                v4 -= 32767
                                fmt.Println(fi7, fi7, v4)
                            }
                        default:
                            if (0 == 0) {
                                if true {
                                    if (!(42 >= 186)) {
                                        if len(v2) > 0 {
                                            v2[0] = "a"
                                        }
                                        fmt.Println(fi8, v5, fi8)
                                        v11 := GenS1{F0: 1, F1: "hello"}
                                        _ = v11
                                    }
                                    var v12 string = "\n\t"
                                    _ = v12
                                }
                                switch ((("\n\t" + "foo bar") + "a") + ("a" + ("foo bar" + (((("hello" + "hello") + (((("" + "") + "hello") + "") + "\n\t")) + ("hello" + ("a" + ""))) + ("foo bar" + (("" + ("hello" + ((("hello" + (("\n\t" + "\n\t") + "a")) + "") + (("hello" + "\n\t") + ("" + "a"))))) + "a")))))) {
                                case "\n\t":
                                    fmt.Println(fi8)
                                case "":
                                    v13 := GenS2{F0: 2147483648, F1: 18446744073709551615, F2: fi7}
                                    _ = v13
                                case "foo bar":
                                    delete(v3, 32767)
                                    if (v4 <= (v4 - v4)) {
                                        v14 := GenS2{F0: 2147483648, F1: 18446744073709551615, F2: 0}
                                        if ((fi7 << ((18446744073709551615) & 63)) < fi8) {
                                            var v15 uint64 = 1
                                            if true {
                                                var v16 uint64 = uint64(v5)
                                                v17 := v3[fi7]
                                                _ = v16
                                                _ = v17
                                            } else {
                                                if true {
                                                    var v18 uint32 = 793
                                                    _ = v18
                                                }
                                            }
                                            switch -2147483648 {
                                            case 151:
                                                if (false || ((v9 % (((-(v9))) | 1)) > 1)) {
                                                    switch (("" + "hello") + "\n\t") {
                                                    case "a":
                                                        v9 = (v9 - v9)
                                                    case "foo bar":
                                                        var v19 uint64 = v15
                                                        _ = v19
                                                    }
                                                    if (590 > 18446744073709551615) {
                                                        v2 = append(v2, "a")
                                                        var v20 uint64 = (v15 >> ((565) & 63))
                                                        _ = v20
                                                    } else {
                                                        var v21 uint = 65535
                                                        _ = v21
                                                    }
                                                    var v22 string = "a"
                                                    var v23 int8 = 100
                                                    v24 := v3[(fi7 - fi8)]
                                                    _ = v22
                                                    _ = v23
                                                    _ = v24
                                                }
                                            case 32767:
                                                v2 = append(v2, "")
                                            case -230:
                                                var v25 uint8 = 145
                                                _ = v25
                                            }
                                            _ = v15
                                        }
                                        switch fi7 {
                                        case 744:
                                            var v26 uint16 = 0
                                            _ = v26
                                        case 127:
                                            _ = len(v6)
                                        }
                                        _ = len(v6)
                                        _ = v14
                                    }
                                    switch "foo bar" {
                                    case "a":
                                        v27 := GenS2{F0: 396, F1: 65535, F2: 9223372036854775807}
                                        _ = v27
                                    case "\n\t":
                                        v5 = (v5 % (((v5 + (v5 / (((v5 << ((1) & 63))) | 1)))) | 1))
                                        v28 := GenS2{F0: 4294967295, F1: 353, F2: int(v9)}
                                        var v29 uint8 = (v5 << ((2) & 63))
                                        if false {
                                            switch "\n\t" {
                                            case "":
                                                fmt.Println(v4, v29)
                                            case "foo bar":
                                                v9 = v9
                                                switch fi8 {
                                                case 320:
                                                    v30 := GenS1{F0: 335, F1: ""}
                                                    _ = v30
                                                case 1:
                                                    switch 42 {
                                                    case 42:
                                                        v4 = v4
                                                    default:
                                                        var v31 int = fi8
                                                        _ = v31
                                                    }
                                                default:
                                                    fmt.Println(v5, fi8)
                                                }
                                                v32 := GenS1{F0: 4294967295, F1: "\n\t"}
                                                _ = v32
                                            case "hello":
                                                switch 42 {
                                                case 4294967295:
                                                    fmt.Println(v29, v9, v9)
                                                case 9223372036854775807:
                                                    v33 := GenS1{F0: 65535, F1: "foo bar"}
                                                    _ = v33
                                                case 1:
                                                    if (true || (433 <= 617)) {
                                                        var v34 float32 = -2.25
                                                        fmt.Println(v34, v9, v5)
                                                        fmt.Println(v9, v29)
                                                        _ = v34
                                                    } else {
                                                        switch fi7 {
                                                        case 4294967295:
                                                            var v35 float64 = ((0.0 * 0.0) * 3.14159)
                                                            _ = v35
                                                        case 2147483647:
                                                            v4 = (v4 / ((1) | 1))
                                                        }
                                                        var v36 int16 = (v4 ^ v4)
                                                        _ = v36
                                                    }
                                                }
                                            }
                                            v29 = 232
                                            if (v4 != -1) {
                                                fmt.Println(fi8, v29, v9)
                                            }
                                            v5 -= 42
                                            fmt.Println(v4)
                                        }
                                        _ = v28
                                        _ = v29
                                    default:
                                        v37 := GenS1{F0: 2, F1: "hello"}
                                        _ = v37
                                    }
                                    if true {
                                        _ = len(v3)
                                        v9 -= (v9 + int64(fi8))
                                    } else {
                                        _ = len(v1)
                                        var v38 int8 = -112
                                        _ = v38
                                    }
                                }
                                var v39 float64 = 1e10
                                var v40 int = (fi8 - 1)
                                var v41 int8 = -128
                                _ = v39
                                _ = v40
                                _ = v41
                            }
                        }
                        var v42 bool = false
                        _ = v42
                    } else {
                        v9 += (v9 * 42)
                        v4 = (v4 % (((v4 + v4)) | 1))
                    }
                    _ = fi8
                    _ = v9
                }
                v43 := GenS3{F0: 65535, F1: 597}
                switch "foo bar" {
                case "hello":
                    v5 = (v5 / (((v5 * v5)) | 1))
                case "foo bar":
                    switch ("" + "hello") {
                    case "hello":
                        v4 -= -539
                    default:
                        v44 := GenS4{F0: true}
                        _ = v44
                    }
                default:
                    var v45 int16 = (v4 * v4)
                    _ = v45
                }
                _ = fi7
                _ = v43
            }
            fmt.Println(v5, v4)
        }
        v4 += -364
        var v46 float64 = (3.14159 / 0.1)
        _ = v46
    }
    var v47 int8 = 42
    fmt.Println(v47, v4)
    for fi48 := 0; fi48 < 16; fi48++ {
        v49 := GenS2{F0: 2147483648, F1: 65535, F2: fi48}
        switch "hello" {
        case "a":
            switch fi48 {
            case -199:
                v5 -= 2
            case -2147483648:
                if (!false) {
                    delete(v3, (-(fi48)))
                    v4 -= v4
                    fmt.Println(v4, v5)
                } else {
                    for fi50 := 0; fi50 < 9; fi50++ {
                        v51 := GenS3{F0: 224, F1: uint(v5)}
                        v52 := GenS5{F0: fi50, F1: 9223372036854775807}
                        var v53 uint16 = 699
                        _ = fi50
                        _ = v51
                        _ = v52
                        _ = v53
                    }
                    fmt.Println(v5, fi48)
                    v47 = (-(v47))
                }
            case 9223372036854775807:
                v5 = (v5 * 127)
            default:
                v5 -= (v5 ^ v5)
            }
        case "hello":
            switch "" {
            case "":
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
            case "\n\t":
                var v54 uint8 = v5
                _ = v54
            }
        case "":
            fmt.Println(fi48)
        }
        v47 = (-(v47))
        _ = fi48
        _ = v49
    }
    fmt.Println(v5, v5)
    v55 := GenS6{F0: int16(v47), F1: 1}
    v56 := GenS7{F0: -1, F1: 65535}
    var v57 float64 = 1e10
    var v58 uint16 = 1
    var v59 uint64 = 742
    v57 = (1.5 - v57)
    var v60 int16 = (v4 ^ 932)
    if true {
        switch -310 {
        case 269:
            _ = len(v3)
        default:
            var v61 uint16 = (v58 + 255)
            _ = v61
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v60
}
