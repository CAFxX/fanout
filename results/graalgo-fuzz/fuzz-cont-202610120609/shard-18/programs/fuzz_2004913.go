package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int16 }
type GenS2 struct { F0 uint; F1 string; F2 int8 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 uint8; F1 int64; F2 uint }
type GenS6 struct { F0 int8; F1 int }
type GenS7 struct { F0 int64; F1 uint16 }
func genFn0(p0 string) int8 {
    for fi4 := 0; fi4 < 16; fi4++ {
        var v5 int32 = -1
        if (-655 > 228) {
            fmt.Println(p0)
            for fi6 := 0; fi6 < 17; fi6++ {
                v7 := []int{(fi6 * 32768)}
                _ = len(v7)
                fmt.Println(v5, v5, fi6)
                _ = fi6
                _ = v7
            }
            p0 = "\n\t"
        }
        _ = fi4
        _ = v5
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        switch fi8 {
        case 4294967295:
            var v9 int8 = -1
            _ = v9
        case -43:
            var v10 uint16 = 255
            _ = v10
        case 1:
            var v11 uint32 = 445
            _ = v11
        }
        if true {
            fmt.Println(p0, p0)
        }
        fmt.Println(p0, p0)
        _ = fi8
    }
    switch (p0 + "\n\t") {
    case "hello":
        var v12 uint16 = 23
        _ = v12
    case "foo bar":
        var v13 int = 971
        _ = v13
    case "\n\t":
        var v14 float64 = 0.0
        _ = v14
    default:
        fmt.Println(p0)
    }
    if (42 > 42) {
        var v15 uint = 420
        _ = v15
    } else {
        var v16 string = ""
        for fi17 := 0; fi17 < 17; fi17++ {
            switch v16 {
            case "a":
                for fi18 := 0; fi18 < 1; fi18++ {
                    if true {
                        var v19 string = "hello"
                        if true {
                            v16 = "hello"
                        }
                        _ = v19
                    }
                    v20 := []int{fi17, 32767, -442}
                    var v21 uint16 = 42
                    v22 := GenS1{F0: 1, F1: -9223372036854775808, F2: 32767}
                    switch fi18 {
                    case -612:
                        var v23 int8 = -128
                        _ = v23
                    }
                    _ = len(v20)
                    _ = fi18
                    _ = v20
                    _ = v21
                    _ = v22
                }
            case "\n\t":
                p0 = p0
            default:
                v16 = (v16 + (p0 + v16))
            }
            if false {
                var v24 string = (p0 + "")
                if (2147483648 > 4294967295) {
                    v24 = "a"
                    p0 = (((("foo bar" + "a") + "foo bar") + (v24 + ((p0 + p0) + ""))) + (v24 + ((p0 + p0) + ("" + "foo bar"))))
                    v25 := []int{(fi17 & fi17), (fi17 << ((1) & 63)), 4294967295, 321}
                    _ = v25
                }
                _ = v24
            } else {
                if false {
                    switch (fi17 % (((-(fi17))) | 1)) {
                    case 127:
                        for fi26 := 0; fi26 < 3; fi26++ {
                            p0 = ("hello" + "hello")
                            _ = fi26
                        }
                    case -449:
                        if ((((1 > 914) && ((0 > 2) && (32767 != 32767))) || false) && (27 != 0)) {
                            if true {
                                var v27 float32 = float32(-2.25)
                                _ = v27
                            }
                            v16 = (p0 + p0)
                        } else {
                            var v28 int8 = 41
                            var v29 bool = (!true)
                            v30 := GenS2{F0: 278, F1: (p0 + "\n\t"), F2: -1}
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                    case -900:
                        v31 := GenS2{F0: 4294967295, F1: p0, F2: 31}
                        _ = v31
                    }
                } else {
                    if ((-2.25 * 0.0) != 0.0) {
                        switch "hello" {
                        case "hello":
                            var v32 int = 4294967295
                            _ = v32
                        case "a":
                            fmt.Println(v16, v16)
                        default:
                            v16 = "foo bar"
                        }
                        var v33 uint64 = 1
                        _ = v33
                    } else {
                        var v34 uint16 = 0
                        _ = v34
                    }
                    var v35 uint32 = 1
                    _ = v35
                }
                p0 = (p0 + p0)
                var v36 uint = 18446744073709551615
                _ = v36
            }
            if (130 <= 987) {
                for fi37 := 0; fi37 < 1; fi37++ {
                    fmt.Println(p0, fi17)
                    if (true || (100 > 127)) {
                        v16 = ("hello" + ("foo bar" + v16))
                    }
                    p0 = ("hello" + "")
                    _ = fi37
                }
                v38 := GenS2{F0: 923, F1: "\n\t", F2: 27}
                var v39 bool = true
                _ = v38
                _ = v39
            } else {
                var v40 uint = 539
                v41 := []int{(fi17 * fi17), 691}
                _ = v40
                _ = v41
            }
            _ = fi17
        }
        _ = v16
    }
    if true {
        switch p0 {
        case "hello":
            var v42 int16 = 0
            _ = v42
        case "foo bar":
            var v43 int64 = 14
            _ = v43
        case "\n\t":
            p0 = ("" + "")
        }
        v44 := GenS2{F0: 4294967295, F1: p0, F2: 127}
        _ = v44
    }
    _ = p0
    return 100
}

func main() {
    var v1 int32 = 0
    var v2 int8 = 100
    var v3 uint32 = 2147483648
    for fi45 := 0; fi45 < 14; fi45++ {
        v46 := GenS3{F0: 883, F1: 2}
        _ = fi45
        _ = v46
    }
    fmt.Println(v2)
    var v47 uint16 = 32767
    v3 = 1
    if (((2147483647 > v3) || (!((0.0 < (1e10 * (-2.25 / 1.0))) || false))) && (!((("a" + "\n\t") + (("a" + "\n\t") + (((((((("foo bar" + ("hello" + "hello")) + "") + "hello") + ("hello" + ("foo bar" + "foo bar"))) + ("foo bar" + "")) + "\n\t") + "foo bar") + "a"))) < "foo bar"))) {
        if false {
            v48 := GenS1{F0: 2, F1: 78, F2: 1}
            v3 += 564
            _ = v48
        } else {
            for fi49 := 0; fi49 < 7; fi49++ {
                fmt.Println(fi49, fi49)
                for fi50 := 0; fi50 < 16; fi50++ {
                    var v51 float64 = 1.5
                    if ((v2 - 100) <= genFn0((("a" + ("\n\t" + "")) + "hello"))) {
                        v52 := []int{-1}
                        var v53 uint64 = 1
                        _ = v52
                        _ = v53
                    } else {
                        fmt.Println(fi50, v3)
                        fmt.Println(fi49, v1, v47)
                        if (!(!true)) {
                            switch fi49 {
                            case 127:
                                fmt.Println(fi50)
                            }
                            fmt.Println(v3, fi50, v51)
                        } else {
                            var v54 int8 = int8(v3)
                            v55 := []int{-2147483648, fi49, fi50}
                            v56 := GenS4{F0: 255, F1: 2147483647}
                            _ = v54
                            _ = v55
                            _ = v56
                        }
                        var v57 int = fi49
                        var v58 bool = false
                        v59 := GenS4{F0: 4294967295, F1: -930}
                        v51 -= ((0.0 / 1.0) * float64(1e10))
                        _ = v57
                        _ = v58
                        _ = v59
                    }
                    fmt.Println(v2, v3)
                    v3 += v3
                    switch int(v3) {
                    case 32767:
                        v3 = v3
                    }
                    _ = fi50
                    _ = v51
                }
                _ = fi49
            }
            v60 := []int{}
            if ((-32768 >= -628) && false) {
                fmt.Println(v1)
            } else {
                var v61 int64 = 0
                v62 := GenS3{F0: v47, F1: v47}
                _ = v61
                _ = v62
            }
            _ = v60
        }
        v63 := []int{9223372036854775807, -691, 32768, 2147483647}
        _ = v63
    } else {
        switch 1 {
        case 930:
            v64 := GenS4{F0: uint(v47), F1: (v1 * (v1 + v1))}
            _ = v64
        default:
            v2 = 42
        }
        v65 := GenS1{F0: 746, F1: -163, F2: 1000}
        fmt.Println(v2, v2, v1)
        _ = v65
    }
    for fi66 := 0; fi66 < 2; fi66++ {
        v67 := GenS1{F0: (v1 << ((255) & 63)), F1: int64(v3), F2: 0}
        v1 -= 2147483647
        v2 = (v2 ^ (-(v2)))
        _ = fi66
        _ = v67
    }
    if (true && false) {
        v1 = v1
        for fi68 := 0; fi68 < 2; fi68++ {
            fmt.Println(v3, fi68)
            for fi69 := 0; fi69 < 18; fi69++ {
                if false {
                    var v70 int64 = 582
                    _ = v70
                } else {
                    v71 := []int{}
                    _ = v71
                }
                v72 := []int{}
                _ = fi69
                _ = v72
            }
            _ = fi68
        }
    } else {
        fmt.Println(v47, v1, v3)
        v73 := GenS3{F0: v47, F1: (v47 ^ v47)}
        if ((-9223372036854775808 == 9223372036854775807) || (((1e10 - (0.0 * 0.1)) * (0.0 - 1.5)) > (((0.0 - float64(((((1e10 - (1e10 / 1.0)) - (float32((1.5 * 1e10)) - 1e10)) + 3.14159) - 3.14159))) / 0.1) / (3.14159 / ((1.5 + -2.25) * (-2.25 + 0.0)))))) {
            v74 := []int{54, 364}
            for fi75 := 0; fi75 < 17; fi75++ {
                if true {
                    v74 = append(v74, int(v2))
                } else {
                    var v76 uint32 = 65535
                    _ = v76
                }
                var v77 string = "\n\t"
                _ = fi75
                _ = v77
            }
            _ = v74
        } else {
            switch "" {
            case "hello":
                var v78 int32 = v1
                _ = v78
            case "foo bar":
                if true {
                    for fi79 := 0; fi79 < 15; fi79++ {
                        v80 := GenS5{F0: 1, F1: -9223372036854775808, F2: 509}
                        _ = fi79
                        _ = v80
                    }
                    v47 += v47
                    var v81 uint = 42
                    _ = v81
                } else {
                    v2 -= 1
                    var v82 int64 = 0
                    _ = v82
                }
            default:
                v83 := GenS6{F0: (v2 * (v2 / ((v2) | 1))), F1: 20}
                _ = v83
            }
        }
        _ = v73
    }
    v47 -= (v47 << ((18446744073709551615) & 63))
    v84 := GenS7{F0: 9223372036854775807, F1: 42}
    if ((v1 << ((4294967295) & 63)) >= v1) {
        var v85 uint32 = (v3 ^ (v3 - (v3 - 1)))
        fmt.Println(v2, v47, v85)
        _ = v85
    }
    switch (("hello" + ("\n\t" + ("foo bar" + "hello"))) + "foo bar") {
    case "a":
        v86 := GenS6{F0: 0, F1: -990}
        _ = v86
    case "\n\t":
        switch 146 {
        case -249:
            var v87 int = int(v3)
            _ = v87
        case 127:
            var v88 int32 = 42
            _ = v88
        default:
            v89 := []int{-1, int(v1), 798, 2147483647}
            _ = v89
        }
    case "":
        if (true || ((v47 << ((4294967295) & 63)) > (v47 & v47))) {
            v90 := []int{int(v3), -1, 0}
            var v91 float32 = 1e10
            v92 := GenS5{F0: 186, F1: 1, F2: 1}
            _ = v90
            _ = v91
            _ = v92
        }
    default:
        switch "hello" {
        case "a":
            fmt.Println(v3, v1, v1)
        case "hello":
            v3 = (v3 / ((2147483648) | 1))
        case "":
            if (v1 < (v1 * (v1 << ((6) & 63)))) {
                v93 := GenS5{F0: 2, F1: 9223372036854775807, F2: 863}
                switch 32767 {
                case 1:
                    switch ("" + "") {
                    case "\n\t":
                        var v94 float64 = float64((0.1 - (1e10 / 1.0)))
                        _ = v94
                    case "foo bar":
                        v95 := []int{4294967295, -9223372036854775808, -385}
                        _ = v95
                    }
                case 4294967295:
                    fmt.Println(v47, v1, v1)
                case 2147483647:
                    v96 := []int{}
                    _ = v96
                }
                _ = v93
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v47
    _ = v84
}
