package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 int32 }
func genFn0(p0 int8) uint16 {
    v5 := GenS1{F0: 756}
    for fi6 := 0; fi6 < 7; fi6++ {
        p0 = (^(p0))
        v7 := GenS1{F0: 32767}
        _ = fi6
        _ = v7
    }
    switch "a" {
    case "foo bar":
        for fi8 := 0; fi8 < 11; fi8++ {
            if (true && false) {
                switch (fi8 + (fi8 & -9223372036854775808)) {
                case -300:
                    v9 := GenS1{F0: 65535}
                    _ = v9
                case 2147483647:
                    var v10 float32 = (((float32(1.5) * 0.0) - float32(1e10)) * 1e10)
                    _ = v10
                }
            } else {
                p0 = (^(p0))
                var v11 int64 = 9223372036854775807
                _ = v11
            }
            for fi12 := 0; fi12 < 5; fi12++ {
                var v13 string = "\n\t"
                fmt.Println(p0, fi12, v13)
                v14 := GenS1{F0: 871}
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi8
        }
    case "\n\t":
        v15 := []int{937, 17, -115}
        _ = v15
    case "a":
        p0 -= 127
    default:
        if false {
            switch ("foo bar" + ("hello" + (("\n\t" + "hello") + ("foo bar" + "hello")))) {
            case "hello":
                if (255 > 42) {
                    v16 := []int{128, 9223372036854775807}
                    _ = v16
                } else {
                    switch -561 {
                    case 4294967295:
                        if false {
                            var v17 float32 = 3.14159
                            _ = v17
                        }
                    case 534:
                        p0 = (-(p0))
                    default:
                        switch "foo bar" {
                        case "a":
                            v18 := []int{628, 1, 127, -1}
                            _ = v18
                        case "":
                            switch ((("hello" + (((("hello" + "a") + ("hello" + "\n\t")) + "\n\t") + ("foo bar" + "foo bar"))) + "") + "") {
                            case "\n\t":
                                v19 := GenS1{F0: 0}
                                _ = v19
                            }
                        }
                    }
                }
            case "\n\t":
                for fi20 := 0; fi20 < 13; fi20++ {
                    switch 32767 {
                    case 2147483647:
                        fmt.Println(p0)
                    case -1:
                        v21 := GenS2{F0: fi20, F1: true}
                        _ = v21
                    case -558:
                        var v22 float64 = ((1e10 / 3.14159) + 0.1)
                        _ = v22
                    default:
                        p0 = 42
                    }
                    var v23 int16 = 912
                    for fi24 := 0; fi24 < 20; fi24++ {
                        var v25 int = fi24
                        switch (^(fi24)) {
                        case -996:
                            v23 = (v23 / ((v23) | 1))
                        default:
                            v23 -= v23
                        }
                        v25 = fi24
                        _ = fi24
                        _ = v25
                    }
                    _ = fi20
                    _ = v23
                }
            }
            for fi26 := 0; fi26 < 20; fi26++ {
                v27 := []int{fi26, (fi26 >> ((609) & 63))}
                if len(v27) > 0 {
                    v27[0] = (fi26 / ((fi26) | 1))
                }
                _ = fi26
                _ = v27
            }
        } else {
            v28 := GenS3{F0: 127, F1: 4294967295}
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        switch fi29 {
        case -1:
            if false {
                var v30 int = (fi29 + (^(fi29)))
                _ = v30
            } else {
                fmt.Println(p0, p0)
            }
        case 4294967295:
            v31 := []int{(^(fi29)), (fi29 ^ fi29), fi29, (fi29 / (((fi29 ^ (fi29 / (((fi29 + (^(fi29)))) | 1)))) | 1))}
            _ = v31
        case 0:
            if false {
                fmt.Println(fi29)
                p0 += p0
                switch "hello" {
                case "\n\t":
                    var v32 uint = 1
                    _ = v32
                case "":
                    var v33 uint8 = 21
                    _ = v33
                }
            } else {
                for fi34 := 0; fi34 < 5; fi34++ {
                    p0 += (p0 * (p0 << ((255) & 63)))
                    switch fi34 {
                    case -281:
                        p0 += 1
                    }
                    p0 = (p0 - p0)
                    _ = fi34
                }
                p0 = p0
                var v35 string = ("\n\t" + ("" + (("a" + "\n\t") + ("a" + ("\n\t" + "")))))
                _ = v35
            }
        }
        switch (("" + ("a" + ((((("foo bar" + ((("a" + "foo bar") + (("hello" + "a") + "")) + "\n\t")) + "foo bar") + ("\n\t" + ("a" + ("hello" + "foo bar")))) + "a") + ""))) + "a") {
        case "hello":
            v36 := []int{(-(fi29)), 4294967295}
            _ = v36
        case "\n\t":
            v37 := GenS4{F0: -1}
            _ = v37
        default:
            var v38 uint64 = 423
            _ = v38
        }
        var v39 float32 = float32(1.5)
        _ = fi29
        _ = v39
    }
    var v40 string = "foo bar"
    _ = p0
    _ = v40
    _ = v5
    return 42
}
func genFn1() int8 {
    var v41 int16 = -513
    switch "foo bar" {
    case "foo bar":
        v41 -= (v41 - v41)
    default:
        v42 := GenS3{F0: 0, F1: 184}
        _ = v42
    }
    v43 := GenS1{F0: 733}
    for fi44 := 0; fi44 < 12; fi44++ {
        var v45 float64 = 0.0
        fmt.Println(v45, v45)
        switch "\n\t" {
        case "foo bar":
            v46 := GenS3{F0: 87, F1: 2147483647}
            _ = v46
        case "":
            var v47 uint32 = 2
            _ = v47
        }
        _ = fi44
        _ = v45
    }
    _ = v41
    _ = v43
    return 42
}

func main() {
    var v1 uint32 = 0
    var v2 int64 = 1
    var v3 uint64 = 1
    v4 := make(map[string]string)
    if false {
        for fi48 := 0; fi48 < 7; fi48++ {
            var v49 bool = true
            _ = fi48
            _ = v49
        }
    }
    var v50 string = "foo bar"
    var v51 uint32 = v1
    var v52 bool = ((v3 * (v3 / ((v3) | 1))) >= uint64(v1))
    var v53 uint16 = 32767
    v53 = genFn0(genFn1())
    _ = len(v4)
    v2 = (-(v2))
    v54 := GenS2{F0: 32768, F1: true}
    if ((-(v2)) >= (v2 & v2)) {
        if v52 {
            _ = len(v4)
            var v55 int16 = 32767
            var v56 int8 = -35
            _ = v55
            _ = v56
        }
        if true {
            var v57 bool = (-2.25 == 0.0)
            switch v50 {
            case "a":
                v58 := GenS2{F0: -167, F1: false}
                _ = v58
            case "":
                if (18446744073709551615 == 621) {
                    v57 = true
                    var v59 float64 = 0.0
                    _ = v59
                }
            case "\n\t":
                if v57 {
                    for fi60 := 0; fi60 < 19; fi60++ {
                        var v61 bool = (746 <= v51)
                        _ = fi60
                        _ = v61
                    }
                } else {
                    if (!v52) {
                        if false {
                            var v62 uint16 = genFn0(-128)
                            _ = v62
                        } else {
                            v3 -= v3
                            if v57 {
                                for fi63 := 0; fi63 < 1; fi63++ {
                                    switch fi63 {
                                    case 1:
                                        v1 = uint32(v3)
                                    case 867:
                                        if (42 == 255) {
                                            v53 += (^(v53))
                                        }
                                    case -2147483648:
                                        fmt.Println(v1)
                                    }
                                    var v64 int = fi63
                                    _ = fi63
                                    _ = v64
                                }
                                var v65 uint8 = 255
                                fmt.Println(v52, v50, v3)
                                _ = v65
                            }
                            fmt.Println(v3, v3)
                        }
                        for fi66 := 0; fi66 < 7; fi66++ {
                            fmt.Println(v52)
                            _ = fi66
                        }
                        v3 += v3
                    } else {
                        var v67 int16 = 0
                        _ = v67
                    }
                    v53 = v53
                }
            }
            _ = v57
        }
        v4[v50] = (v50 + v50)
    }
    for fi68 := 0; fi68 < 9; fi68++ {
        v51 = v51
        var v69 float32 = (3.14159 + (1e10 - ((1e10 * float32(((3.14159 - (1e10 - (1e10 * -2.25))) / -2.25))) + 0.1)))
        fmt.Println(v2, v69)
        _ = fi68
        _ = v69
    }
    var v70 uint = 2
    v52 = v52
    for fi71 := 0; fi71 < 19; fi71++ {
        var v72 float64 = (3.14159 * 0.1)
        _ = fi71
        _ = v72
    }
    var v73 uint64 = (v3 * (v3 | 1))
    switch 468 {
    case 105:
        v51 += v51
    case 42:
        var v74 float64 = float64(1.5)
        _ = v74
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v70:", v70)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v70
    _ = v73
}
