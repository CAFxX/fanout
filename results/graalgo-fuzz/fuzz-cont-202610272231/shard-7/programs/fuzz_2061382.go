package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 int64; F1 string; F2 uint64 }
type GenS5 struct { F0 uint64; F1 int16; F2 uint32 }
func genFn0() int {
    v5 := []int{671, 2147483647}
    var v6 int32 = 42
    var v7 int64 = 0
    v6 = -1
    v8 := GenS1{F0: 18446744073709551615}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -1
}
func genFn1() bool {
    v9 := []int{genFn0(), 128, -9223372036854775808}
    _ = len(v9)
    var v10 int32 = 2147483647
    var v11 uint8 = 127
    v11 += (v11 % ((255) | 1))
    _ = v10
    _ = v11
    _ = v9
    return ((v10 << ((315) & 63)) < -471)
}

func main() {
    var v1 int16 = -32768
    var v2 float64 = 1e10
    var v3 uint = 449
    var v4 uint16 = 42
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := []int{(fi12 + (fi12 - genFn0()))}
        _ = fi12
        _ = v13
    }
    v14 := []int{9223372036854775807, 1}
    if true {
        if true {
            if true {
                if (1 <= 127) {
                    var v15 int16 = v1
                    v2 = (((float64((3.14159 / 1.5)) - (v2 - 1e10)) - 0.1) + v2)
                    for fi16 := 0; fi16 < 20; fi16++ {
                        var v17 uint64 = 4294967295
                        for fi18 := 0; fi18 < 15; fi18++ {
                            if false {
                                if true {
                                    var v19 uint16 = 998
                                    v19 = (v19 | (^(v4)))
                                    v19 -= 35
                                    _ = v19
                                }
                                switch ("hello" + ("a" + ("" + "a"))) {
                                case "a":
                                    var v20 int = 2147483647
                                    _ = v20
                                case "hello":
                                    v21 := GenS1{F0: 18446744073709551615}
                                    _ = v21
                                }
                            }
                            if (((("a" >= "hello") || (false || true)) && true) || ((v3 >> (((v3 / (((^(v3))) | 1))) & 63)) == (v3 / (((^(v3))) | 1)))) {
                                var v22 int = (^(fi16))
                                v3 = 4294967295
                                switch fi16 {
                                case -180:
                                    var v23 bool = (v17 <= 577)
                                    _ = v23
                                case 2147483647:
                                    fmt.Println(fi18)
                                case -424:
                                    if false {
                                        v1 += v1
                                        if (true || (!(2 == 1))) {
                                            var v24 float64 = v2
                                            var v25 int64 = 255
                                            v22 = (v22 * fi16)
                                            _ = v24
                                            _ = v25
                                        } else {
                                            switch "hello" {
                                            case "a":
                                                var v26 uint64 = (^(v17))
                                                _ = v26
                                            case "\n\t":
                                                var v27 int32 = -370
                                                _ = v27
                                            default:
                                                var v28 uint64 = v17
                                                _ = v28
                                            }
                                            var v29 float32 = (float32(v2) - (1e10 * 3.14159))
                                            v1 = (v1 << ((v3) & 63))
                                            _ = v29
                                        }
                                    }
                                default:
                                    if ((v15 - 32767) >= -1) {
                                        switch "" {
                                        case "":
                                            var v30 int64 = 1
                                            _ = v30
                                        case "a":
                                            var v31 int = fi18
                                            _ = v31
                                        }
                                        v32 := GenS2{F0: -922}
                                        _ = v32
                                    }
                                }
                                _ = v22
                            }
                            if true {
                                _ = len(v14)
                            } else {
                                var v33 uint64 = v17
                                _ = v33
                            }
                            _ = fi18
                        }
                        var v34 int16 = 0
                        _ = fi16
                        _ = v17
                        _ = v34
                    }
                    _ = v15
                } else {
                    var v35 bool = ((0.1 + 1e10) <= ((-2.25 + 1.5) / 1e10))
                    for fi36 := 0; fi36 < 17; fi36++ {
                        for fi37 := 0; fi37 < 12; fi37++ {
                            var v38 int8 = 127
                            switch -1 {
                            case -1:
                                if (-928 >= 2147483647) {
                                    v2 = -2.25
                                    switch fi36 {
                                    case 127:
                                        switch "\n\t" {
                                        case "hello":
                                            var v39 string = "foo bar"
                                            _ = v39
                                        default:
                                            v38 -= v38
                                            v38 = (v38 % (((-(v38))) | 1))
                                            switch (fi36 >> (((v3 ^ v3)) & 63)) {
                                            case 386:
                                                _ = len(v14)
                                            }
                                        }
                                    case 32767:
                                        v40 := GenS3{F0: (("\n\t" != ("hello" + ("a" + ("" + ("hello" + (("foo bar" + "foo bar") + "hello")))))) || false), F1: (fi37 + (fi37 << ((2) & 63)))}
                                        _ = v40
                                    case 128:
                                        var v41 uint16 = v4
                                        _ = v41
                                    }
                                    switch fi36 {
                                    case -261:
                                        switch genFn0() {
                                        case -1:
                                            fmt.Println(v2, fi36, v35)
                                        case -486:
                                            fmt.Println(v35)
                                        case 128:
                                            switch "\n\t" {
                                            case "":
                                                switch (fi37 & -525) {
                                                case -9223372036854775808:
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        var v42 uint16 = 138
                                                        _ = v42
                                                    case "foo bar":
                                                        if len(v14) > 0 {
                                                            v14[0] = (fi36 | (fi36 | (-(fi37))))
                                                        }
                                                    case "hello":
                                                        v43 := GenS3{F0: ((1.5 - 0.0) == float32((1e10 / 1.0))), F1: -213}
                                                        _ = v43
                                                    default:
                                                        v1 = v1
                                                    }
                                                default:
                                                    if v35 {
                                                        v38 -= (v38 % ((int8(v3)) | 1))
                                                    } else {
                                                        v44 := GenS4{F0: 9223372036854775807, F1: (("" + "a") + "\n\t"), F2: 122}
                                                        fmt.Println(fi37, fi37)
                                                        if v35 {
                                                            var v45 int16 = -579
                                                            _ = v45
                                                        }
                                                        _ = v44
                                                    }
                                                }
                                            default:
                                                fmt.Println(v1, fi37, v3)
                                            }
                                        default:
                                            v3 -= (v3 & (v3 / ((v3) | 1)))
                                        }
                                        fmt.Println(fi36, v3, v3)
                                    case -313:
                                        v4 -= (v4 >> ((v3) & 63))
                                    default:
                                        v14 = append(v14, genFn0())
                                    }
                                    v4 = (v4 & (v4 ^ (^(v4))))
                                } else {
                                    fmt.Println(v38, v35, v3)
                                    var v46 float64 = v2
                                    _ = v46
                                }
                            default:
                                v2 = (float64(1.5) / (v2 / (-2.25 - 0.1)))
                            }
                            _ = fi37
                            _ = v38
                        }
                        for fi47 := 0; fi47 < 14; fi47++ {
                            if v35 {
                                _ = len(v14)
                            } else {
                                switch ("foo bar" + "hello") {
                                case "foo bar":
                                    _ = len(v14)
                                }
                                v4 += v4
                                v48 := GenS4{F0: -769, F1: "", F2: 1}
                                _ = v48
                            }
                            _ = fi47
                        }
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        _ = fi36
                    }
                    _ = v35
                }
                var v49 uint16 = uint16(v3)
                _ = v49
            } else {
                if true {
                    switch "\n\t" {
                    case "foo bar":
                        for fi50 := 0; fi50 < 5; fi50++ {
                            var v51 uint16 = (v4 | v4)
                            v2 = (v2 / v2)
                            _ = fi50
                            _ = v51
                        }
                    case "":
                        fmt.Println(v1, v2, v2)
                    }
                }
                for fi52 := 0; fi52 < 14; fi52++ {
                    v3 = (v3 | v3)
                    for fi53 := 0; fi53 < 5; fi53++ {
                        v54 := GenS4{F0: -9223372036854775808, F1: "", F2: uint64(v4)}
                        if true {
                            var v55 uint16 = v4
                            v56 := GenS4{F0: 1, F1: (((("a" + "foo bar") + "a") + ("a" + "foo bar")) + ("\n\t" + "foo bar")), F2: 2}
                            var v57 uint = (v3 ^ (^(v3)))
                            _ = v55
                            _ = v56
                            _ = v57
                        } else {
                            var v58 int32 = 911
                            _ = v58
                        }
                        var v59 int8 = -128
                        _ = fi53
                        _ = v54
                        _ = v59
                    }
                    _ = fi52
                }
            }
            v1 -= 32767
            if (true && (!false)) {
                _ = len(v14)
                for fi60 := 0; fi60 < 9; fi60++ {
                    switch ("\n\t" + "foo bar") {
                    case "\n\t":
                        for fi61 := 0; fi61 < 12; fi61++ {
                            v2 -= 1e10
                            _ = fi61
                        }
                    case "hello":
                        v14 = append(v14, (fi60 + fi60))
                    case "foo bar":
                        var v62 bool = (false && true)
                        _ = v62
                    }
                    var v63 float32 = 0.1
                    _ = fi60
                    _ = v63
                }
                fmt.Println(v4, v4)
            }
        } else {
            if (128 != 31) {
                v64 := GenS2{F0: (v1 & (^(v1)))}
                var v65 float64 = (1.5 + float64(((1e10 + ((0.0 - -2.25) + 1e10)) * 0.0)))
                _ = v64
                _ = v65
            } else {
                v2 = (v2 + float64(3.14159))
                if (true && true) {
                    v3 = 18446744073709551615
                    fmt.Println(v3, v2, v4)
                }
                v66 := GenS5{F0: 1, F1: v1, F2: 859}
                _ = v66
            }
        }
        v2 += -2.25
        v67 := GenS4{F0: -9223372036854775808, F1: "\n\t", F2: 2}
        _ = v67
    } else {
        var v68 float32 = (1.5 / 0.1)
        if true {
            if true {
                switch "foo bar" {
                case "":
                    if false {
                        var v69 int16 = v1
                        switch 32768 {
                        case 821:
                            v70 := GenS2{F0: (v1 | v1)}
                            _ = v70
                        }
                        if true {
                            v2 = v2
                            var v71 uint8 = 49
                            var v72 string = ("foo bar" + "\n\t")
                            _ = v71
                            _ = v72
                        }
                        _ = v69
                    }
                default:
                    var v73 int = 4294967295
                    _ = v73
                }
                _ = len(v14)
                fmt.Println(v3, v4)
            }
        }
        var v74 int = 0
        _ = v68
        _ = v74
    }
    v75 := GenS2{F0: 1000}
    var v76 int16 = v1
    if (uint16(v3) != v4) {
        fmt.Println(v3)
    }
    v4 = (^(v4))
    var v77 float64 = v2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v75
    _ = v76
    _ = v77
}
