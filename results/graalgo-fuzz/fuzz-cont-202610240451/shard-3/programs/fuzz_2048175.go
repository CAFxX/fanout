package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint8 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int8 }
func genFn0() uint8 {
    switch "a" {
    case "":
        var v4 uint = 586
        _ = v4
    case "foo bar":
        var v5 float32 = 0.0
        _ = v5
    case "hello":
        if false {
            var v6 int64 = -9223372036854775808
            _ = v6
        } else {
            var v7 int64 = 0
            _ = v7
        }
    }
    var v8 uint = 65535
    var v9 int64 = 9223372036854775807
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int64, p1 bool) int32 {
    if p1 {
        var v10 int64 = (-(p0))
        var v11 float32 = float32(3.14159)
        var v12 int8 = 127
        _ = v10
        _ = v11
        _ = v12
    }
    if p1 {
        for fi13 := 0; fi13 < 14; fi13++ {
            v14 := []int{4294967295, (fi13 ^ fi13), fi13}
            var v15 uint = 520
            fmt.Println(fi13)
            _ = fi13
            _ = v14
            _ = v15
        }
        v16 := []int{-916}
        v17 := GenS1{F0: 42, F1: 127, F2: 128}
        _ = v16
        _ = v17
    }
    fmt.Println(p0)
    p0 = (p0 * (^(p0)))
    v18 := GenS1{F0: 836, F1: 943, F2: 0}
    _ = p0
    _ = p1
    _ = v18
    return -801
}

func main() {
    var v1 uint32 = 381
    var v2 int32 = -1
    v3 := make(map[uint32]uint32)
    var v19 uint16 = 0
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 int = (fi20 / ((fi20) | 1))
        var v22 int32 = (v2 * 901)
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 uint = 2
    switch "foo bar" {
    case "\n\t":
        if (-713 != genFn1(-1, ((0.1 * ((-2.25 * (((0.0 / (0.0 - 1.5)) * 0.1) / 1e10)) + 1.5)) != (-2.25 + 1.5)))) {
            v19 += v19
            var v24 string = "hello"
            _ = v24
        } else {
            switch 127 {
            case 32768:
                if (65535 <= 626) {
                    if false {
                        _ = len(v3)
                    } else {
                        var v25 int16 = int16(v23)
                        _ = v25
                    }
                } else {
                    var v26 int16 = 0
                    fmt.Println(v19, v2)
                    _ = v26
                }
            default:
                v19 += 443
            }
        }
    case "foo bar":
        v23 = 192
    case "a":
        for fi27 := 0; fi27 < 4; fi27++ {
            switch (fi27 & -2147483648) {
            case 42:
                if ((v23 + (v23 - v23)) != uint(v1)) {
                    v1 += v1
                } else {
                    var v28 bool = true
                    switch ("a" + ("hello" + "hello")) {
                    case "":
                        v29 := []int{(fi27 & -379), (fi27 - fi27)}
                        _ = v29
                    }
                    var v30 bool = true
                    _ = v28
                    _ = v30
                }
            case 2147483647:
                v31 := GenS2{F0: 0, F1: (v2 >> ((243) & 63))}
                _ = v31
            case 127:
                var v32 uint32 = 2147483647
                _ = v32
            }
            _ = fi27
        }
    default:
        if ((v23 & 383) <= v23) {
            var v33 uint32 = v1
            var v34 int8 = 0
            _ = v33
            _ = v34
        }
    }
    if false {
        fmt.Println(v1)
        var v35 float32 = ((1e10 + 1e10) / 1.5)
        for fi36 := 0; fi36 < 9; fi36++ {
            switch (("" + "\n\t") + "") {
            case "":
                if (genFn0() < 42) {
                    if false {
                        v37 := GenS1{F0: 42, F1: 127, F2: 255}
                        v38 := GenS3{F0: 1}
                        v23 = 443
                        _ = v37
                        _ = v38
                    } else {
                        for fi39 := 0; fi39 < 16; fi39++ {
                            var v40 uint = 42
                            _ = fi39
                            _ = v40
                        }
                        var v41 uint64 = 644
                        v23 = (^(v23))
                        _ = v41
                    }
                    if true {
                        for fi42 := 0; fi42 < 13; fi42++ {
                            var v43 float32 = v35
                            _ = fi42
                            _ = v43
                        }
                        switch fi36 {
                        case -1:
                            for fi44 := 0; fi44 < 14; fi44++ {
                                v45 := GenS1{F0: 605, F1: -22, F2: 0}
                                var v46 uint64 = 1
                                _ = fi44
                                _ = v45
                                _ = v46
                            }
                        default:
                            v47 := v3[v1]
                            _ = v47
                        }
                        v48 := GenS1{F0: 826, F1: (-(fi36)), F2: 87}
                        _ = v48
                    }
                }
            }
            if (false || ((v23 & (v23 % ((4294967295) | 1))) < (v23 * (v23 % ((2) | 1))))) {
                v49 := []int{9223372036854775807, 9223372036854775807, -127, fi36}
                for fi50 := 0; fi50 < 9; fi50++ {
                    _ = len(v3)
                    fmt.Println(fi36)
                    _ = fi50
                }
                _ = v49
            }
            var v51 int64 = -9223372036854775808
            _ = fi36
            _ = v51
        }
        _ = v35
    } else {
        v23 += v23
        v52 := GenS3{F0: -9223372036854775808}
        _ = v52
    }
    v19 -= v19
    fmt.Println(v19, v2)
    var v53 int8 = -128
    switch 908 {
    case 2147483647:
        if (-1 <= (v53 % ((int8(v2)) | 1))) {
            for fi54 := 0; fi54 < 19; fi54++ {
                if false {
                    switch -939 {
                    case -515:
                        fmt.Println(fi54)
                    case 364:
                        for fi55 := 0; fi55 < 19; fi55++ {
                            v2 = 42
                            _ = fi55
                        }
                    default:
                        var v56 int8 = 29
                        _ = v56
                    }
                } else {
                    switch ("foo bar" + ("" + ("hello" + "hello"))) {
                    case "foo bar":
                        for fi57 := 0; fi57 < 15; fi57++ {
                            var v58 bool = ((v53 | -1) <= (v53 + int8(v1)))
                            var v59 string = ""
                            v60 := []int{(fi57 & (fi57 / ((fi54) | 1))), 976, -360, 709}
                            if (!(((((904 == 2) && v58) || (v58 && (v58 && ((fi54 == fi54) || (77 >= uint16(v23)))))) && v58) || (-101 == 9223372036854775807))) {
                                var v61 int8 = int8(v23)
                                v59 = (v59 + "hello")
                                fmt.Println(v58)
                                var v62 float32 = (3.14159 - 3.14159)
                                _ = v61
                                _ = v62
                            }
                            var v63 uint16 = v19
                            _ = fi57
                            _ = v58
                            _ = v59
                            _ = v60
                            _ = v63
                        }
                    case "":
                        switch fi54 {
                        case -224:
                            if false {
                                switch fi54 {
                                case 32767:
                                    fmt.Println(fi54, fi54)
                                case 128:
                                    v2 -= 42
                                }
                            } else {
                                fmt.Println(v1, v23)
                                for fi64 := 0; fi64 < 5; fi64++ {
                                    v19 -= v19
                                    _ = fi64
                                }
                                switch (((("" + ((("\n\t" + ("" + "a")) + "foo bar") + ((("\n\t" + ("a" + "hello")) + ("" + ("foo bar" + (("" + "") + "")))) + "a"))) + ("a" + ((("hello" + "\n\t") + "foo bar") + ("" + "a")))) + "foo bar") + "hello") {
                                case "hello":
                                    var v65 int32 = (v2 + genFn1(9223372036854775807, true))
                                    _ = v65
                                case "a":
                                    var v66 int32 = 42
                                    _ = v66
                                case "":
                                    v1 -= (^(v1))
                                default:
                                    switch ("a" + "foo bar") {
                                    case "a":
                                        v67 := []int{}
                                        _ = v67
                                    case "\n\t":
                                        v2 = v2
                                    }
                                }
                            }
                        }
                    case "\n\t":
                        switch (fi54 >> ((v23) & 63)) {
                        case 1:
                            for fi68 := 0; fi68 < 2; fi68++ {
                                fmt.Println(v19)
                                var v69 bool = (int32(v23) != (v2 - -445))
                                _ = fi68
                                _ = v69
                            }
                        case 2147483647:
                            fmt.Println(v53, v1, v1)
                        }
                    }
                    fmt.Println(fi54, v23)
                }
                _ = fi54
            }
            var v70 int8 = v53
            if false {
                v71 := GenS4{F0: 32767}
                _ = v71
            }
            _ = v70
        }
    default:
        for fi72 := 0; fi72 < 10; fi72++ {
            switch (-(fi72)) {
            case 9223372036854775807:
                v73 := GenS5{F0: (v19 ^ v19)}
                _ = v73
            case -2147483648:
                if ((-2.25 <= -2.25) || (!false)) {
                    v2 = v2
                }
            default:
                v53 += (v53 % ((v53) | 1))
            }
            v23 -= v23
            _ = fi72
        }
    }
    var v74 uint64 = 2
    var v75 uint8 = uint8(v23)
    if (!("foo bar" > ("" + "foo bar"))) {
        v76 := GenS6{F0: 128}
        _ = v76
    } else {
        var v77 bool = ((!true) || false)
        var v78 float32 = -2.25
        _ = v77
        _ = v78
    }
    v79 := GenS7{F0: v53}
    v74 -= (v74 | (^(v74)))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v53:", v53)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v53
    _ = v74
    _ = v75
    _ = v79
}
