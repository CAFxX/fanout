package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 int32 }
func genFn0(p0 int8) uint8 {
    var v4 int16 = -1
    var v5 uint64 = 442
    _ = p0
    _ = v4
    _ = v5
    return 2
}
func genFn1() string {
    if true {
        var v6 int16 = -1
        v7 := []int{128}
        _ = v6
        _ = v7
    } else {
        switch 2147483647 {
        case 0:
            if (127 != 42) {
                var v8 int64 = 739
                v9 := []int{}
                if (!false) {
                    if (!true) {
                        var v10 int64 = 1
                        var v11 bool = (1e10 < 0.1)
                        _ = v10
                        _ = v11
                    }
                    v8 = v8
                    fmt.Println(v8)
                }
                _ = v8
                _ = v9
            }
        case 606:
            switch "" {
            case "\n\t":
                v12 := []int{128, 2147483647, 9223372036854775807}
                _ = v12
            case "a":
                var v13 int = 4294967295
                _ = v13
            }
        case 4294967295:
            v14 := GenS1{F0: -32768}
            _ = v14
        }
        var v15 uint16 = 0
        _ = v15
    }
    var v16 uint32 = 0
    _ = v16
    return "foo bar"
}

func main() {
    v1 := []string{""}
    v2 := []int{-2147483648, 903, -9223372036854775808, 32767}
    var v3 bool = (!true)
    v17 := GenS1{F0: 32767}
    if (618 < 32768) {
        var v18 int8 = -128
        v3 = (255 == 255)
        v3 = (v18 <= (v18 / ((-1) | 1)))
        _ = v18
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        switch fi19 {
        case 128:
            fmt.Println(v3, v3)
        }
        _ = fi19
    }
    switch "\n\t" {
    case "\n\t":
        if len(v1) > 0 {
            v1[0] = (("a" + ("" + "foo bar")) + (("foo bar" + "a") + "a"))
        }
    case "":
        if v3 {
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
            switch "\n\t" {
            case "":
                var v20 string = "foo bar"
                _ = v20
            case "\n\t":
                if (((0 == -128) && false) || (9223372036854775807 >= 9223372036854775807)) {
                    v3 = (2 > 520)
                } else {
                    var v21 uint32 = 2147483648
                    switch -2147483648 {
                    case -296:
                        v3 = ((128 <= 225) || false)
                    case 115:
                        var v22 int64 = -1
                        _ = v22
                    case -471:
                        switch 0 {
                        case -9223372036854775808:
                            var v23 float64 = 0.1
                            _ = v23
                        case -2147483648:
                            v21 += v21
                        }
                    default:
                        v24 := GenS2{F0: 893}
                        _ = v24
                    }
                    v25 := GenS2{F0: 590}
                    _ = v21
                    _ = v25
                }
            }
        }
    default:
        var v26 uint64 = 42
        _ = v26
    }
    for fi27 := 0; fi27 < 12; fi27++ {
        fmt.Println(fi27)
        _ = fi27
    }
    if v3 {
        _ = len(v1)
        v28 := GenS3{F0: 9223372036854775807}
        switch -9223372036854775808 {
        case -9223372036854775808:
            if true {
                v3 = ("foo bar" >= "\n\t")
                for fi29 := 0; fi29 < 18; fi29++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v30 float32 = (-2.25 + 1e10)
                    _ = fi29
                    _ = v30
                }
                if (259 < 470) {
                    v31 := GenS2{F0: 2}
                    var v32 int = 0
                    if true {
                        v3 = v3
                    } else {
                        v33 := GenS4{F0: 846, F1: "\n\t"}
                        _ = v33
                    }
                    _ = v31
                    _ = v32
                }
            } else {
                for fi34 := 0; fi34 < 6; fi34++ {
                    v1 = append(v1, "\n\t")
                    v3 = (1 > 2147483647)
                    _ = fi34
                }
                fmt.Println(v3)
                v35 := GenS3{F0: -9223372036854775808}
                _ = v35
            }
        case 127:
            for fi36 := 0; fi36 < 19; fi36++ {
                var v37 uint8 = 42
                _ = fi36
                _ = v37
            }
        case -2147483648:
            var v38 uint64 = 969
            _ = v38
        }
        _ = v28
    }
    if (!(false && (1 != -1))) {
        switch "a" {
        case "hello":
            if true {
                for fi39 := 0; fi39 < 20; fi39++ {
                    v3 = (2 <= 754)
                    var v40 int16 = -273
                    _ = fi39
                    _ = v40
                }
            }
        case "":
            if v3 {
                if v3 {
                    if false {
                        v41 := GenS2{F0: 255}
                        var v42 int = -1
                        var v43 float32 = 0.0
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                } else {
                    switch "hello" {
                    case "a":
                        for fi44 := 0; fi44 < 4; fi44++ {
                            v45 := GenS4{F0: 255, F1: (("hello" + "foo bar") + ((((("foo bar" + (("a" + "hello") + "a")) + "") + "hello") + "foo bar") + ""))}
                            for fi46 := 0; fi46 < 17; fi46++ {
                                v47 := GenS5{F0: 2147483647}
                                if len(v2) > 0 {
                                    v2[0] = (fi44 & fi46)
                                }
                                v3 = (18446744073709551615 > 2)
                                _ = fi46
                                _ = v47
                            }
                            v3 = ((((0.0 / -2.25) + -2.25) >= 3.14159) && (v3 || true))
                            _ = fi44
                            _ = v45
                        }
                    case "\n\t":
                        for fi48 := 0; fi48 < 17; fi48++ {
                            _ = len(v1)
                            v49 := GenS1{F0: 0}
                            _ = fi48
                            _ = v49
                        }
                    }
                    for fi50 := 0; fi50 < 10; fi50++ {
                        if true {
                            if false {
                                var v51 string = "hello"
                                var v52 uint8 = 2
                                _ = v51
                                _ = v52
                            }
                        }
                        var v53 float32 = ((1.5 / (1e10 * 1.5)) * 1e10)
                        _ = fi50
                        _ = v53
                    }
                    var v54 int32 = 2147483647
                    _ = v54
                }
                var v55 uint16 = 108
                var v56 float32 = (0.1 + 0.0)
                _ = v55
                _ = v56
            }
        case "\n\t":
            var v57 int64 = 42
            _ = v57
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        if ((0 != 42) && true) {
            var v58 uint8 = 103
            var v59 uint64 = uint64(v58)
            var v60 uint8 = (v58 >> ((2) & 63))
            _ = v58
            _ = v59
            _ = v60
        } else {
            switch "foo bar" {
            case "hello":
                if (32768 != -2147483648) {
                    _ = len(v2)
                } else {
                    if false {
                        v3 = (145 == genFn0(-128))
                        if len(v2) > 0 {
                            v2[0] = -364
                        }
                        var v61 float64 = ((0.1 - 3.14159) - float64((1.5 - 3.14159)))
                        _ = v61
                    }
                    v3 = (255 > 890)
                }
            case "":
                var v62 float32 = (0.0 / (1.5 * 0.1))
                _ = v62
            default:
                v3 = v3
            }
            var v63 uint8 = 216
            var v64 uint8 = (v63 % (((v63 * v63)) | 1))
            _ = v63
            _ = v64
        }
        v3 = v3
        if (((!v3) && (!v3)) || true) {
            v65 := GenS4{F0: 2, F1: "\n\t"}
            _ = v65
        }
    }
    switch "\n\t" {
    case "hello":
        var v66 uint16 = 376
        _ = v66
    }
    v67 := GenS4{F0: 65535, F1: "hello"}
    v68 := GenS5{F0: -237}
    var v69 int8 = 25
    v2 = append(v2, -606)
    v70 := GenS2{F0: 42}
    fmt.Println(v69, v69)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v67
    _ = v68
    _ = v69
    _ = v70
}
