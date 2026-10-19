package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 bool }
type GenS2 struct { F0 string; F1 int; F2 uint64 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 uint8; F1 bool; F2 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint32; F1 uint32; F2 uint32 }

func main() {
    v1 := []int{2147483647, 1, 4294967295, -173}
    v2 := []int{-1}
    var v3 uint64 = 0
    var v4 int16 = 291
    switch "a" {
    case "foo bar":
        for fi5 := 0; fi5 < 9; fi5++ {
            var v6 int8 = int8(v4)
            if (!true) {
                fmt.Println(v6)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v4 += (v4 * v4)
            } else {
                if (true || ("\n\t" < "")) {
                    v7 := GenS1{F0: 65535, F1: 65535, F2: true}
                    if (!true) {
                        if len(v1) > 0 {
                            v1[0] = fi5
                        }
                    }
                    var v8 int32 = 664
                    _ = v7
                    _ = v8
                }
                v6 -= -1
            }
            if len(v2) > 0 {
                v2[0] = -1
            }
            _ = fi5
            _ = v6
        }
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        v3 -= (v3 ^ 18446744073709551615)
        v3 = v3
        for fi10 := 0; fi10 < 17; fi10++ {
            if false {
                var v11 uint8 = 42
                var v12 int8 = 42
                v12 = -44
                var v13 int32 = -1
                v12 = (v12 | -128)
                _ = v11
                _ = v12
                _ = v13
            } else {
                var v14 int = fi10
                switch (fi10 & (fi10 ^ (v14 & -2147483648))) {
                case -9223372036854775808:
                    var v15 int = (fi9 / ((fi10) | 1))
                    _ = v15
                }
                v2 = append(v2, fi9)
                _ = v14
            }
            var v16 int = fi10
            _ = fi10
            _ = v16
        }
        _ = fi9
    }
    v17 := GenS1{F0: 18446744073709551615, F1: 429, F2: true}
    var v18 int32 = 42
    v3 = 1
    v4 -= (v4 & (v4 & (v4 * v4)))
    var v19 int16 = v4
    fmt.Println(v3, v19)
    v20 := GenS1{F0: 0, F1: 2147483648, F2: (!true)}
    switch "a" {
    case "hello":
        fmt.Println(v18, v19, v19)
    case "a":
        v18 += (v18 / ((-2147483648) | 1))
    case "":
        _ = len(v1)
    default:
        switch 9223372036854775807 {
        case 32768:
            for fi21 := 0; fi21 < 6; fi21++ {
                fmt.Println(v4, v19)
                v4 -= (v4 % ((v4) | 1))
                v4 -= 0
                _ = fi21
            }
        case 345:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 uint = 18446744073709551615
                _ = fi22
                _ = v23
            }
        }
    }
    v24 := GenS2{F0: "", F1: 9223372036854775807, F2: v3}
    v25 := GenS3{F0: 42, F1: ("foo bar" < "\n\t")}
    var v26 float32 = ((-2.25 - (-2.25 - ((((0.0 / 1.5) - -2.25) / -2.25) * (0.0 + ((0.0 / float32((-2.25 / 0.1))) / 1e10))))) / 1.0)
    if false {
        for fi27 := 0; fi27 < 15; fi27++ {
            var v28 uint8 = 255
            switch 42 {
            case 0:
                for fi29 := 0; fi29 < 13; fi29++ {
                    v30 := GenS4{F0: v28, F1: false, F2: 2147483647}
                    _ = fi29
                    _ = v30
                }
            }
            if false {
                for fi31 := 0; fi31 < 20; fi31++ {
                    if false {
                        var v32 uint8 = v28
                        if true {
                            v28 = (v28 ^ (v32 + 0))
                            var v33 int8 = 127
                            _ = v33
                        }
                        if true {
                            v34 := GenS5{F0: true}
                            var v35 int64 = 42
                            _ = v34
                            _ = v35
                        }
                        _ = v32
                    } else {
                        switch -354 {
                        case 127:
                            var v36 uint16 = 32767
                            _ = v36
                        case 2147483647:
                            var v37 uint64 = v3
                            _ = v37
                        }
                        fmt.Println(v4, v18, v4)
                        switch (fi31 | -9223372036854775808) {
                        case 32768:
                            switch "\n\t" {
                            case "":
                                switch "a" {
                                case "\n\t":
                                    fmt.Println(v4, fi27)
                                }
                            case "hello":
                                v38 := GenS6{F0: "foo bar"}
                                _ = v38
                            }
                        case 207:
                            if false {
                                if (v3 < (v3 | (v3 | 948))) {
                                    fmt.Println(v28, v28, fi27)
                                    var v39 bool = (!true)
                                    _ = v39
                                }
                                v18 -= int32(v3)
                            }
                        case 42:
                            var v40 float32 = ((((((v26 * v26) + 0.0) + float32((1e10 + 0.0))) - v26) - 1.5) + 1e10)
                            _ = v40
                        default:
                            v4 -= 0
                        }
                    }
                    switch fi31 {
                    case -1:
                        v28 = (v28 << ((2) & 63))
                    case 42:
                        v1 = append(v1, (fi27 ^ (-(fi31))))
                    case 32768:
                        var v41 uint16 = 2
                        _ = v41
                    }
                    fmt.Println(v19, v4, fi31)
                    _ = fi31
                }
                var v42 uint8 = v28
                if (true || (127 > (fi27 >> ((2) & 63)))) {
                    if true {
                        for fi43 := 0; fi43 < 18; fi43++ {
                            fmt.Println(v4)
                            if (!false) {
                                var v44 int8 = 73
                                var v45 int32 = v18
                                _ = v44
                                _ = v45
                            } else {
                                if true {
                                    if len(v2) > 0 {
                                        v2[0] = fi27
                                    }
                                    if (18446744073709551615 != v3) {
                                        v42 -= (v42 * 144)
                                        var v46 uint16 = 42
                                        if false {
                                            switch fi27 {
                                            case 0:
                                                var v47 string = (("" + ((("hello" + "") + "foo bar") + "\n\t")) + ("foo bar" + ""))
                                                _ = v47
                                            case 2147483647:
                                                if true {
                                                    v3 -= (v3 - (^(v3)))
                                                    var v48 int8 = 127
                                                    var v49 int8 = v48
                                                    _ = v48
                                                    _ = v49
                                                } else {
                                                    var v50 bool = true
                                                    v42 -= v42
                                                    _ = v50
                                                }
                                            default:
                                                switch fi27 {
                                                case 2147483647:
                                                    v51 := GenS1{F0: 988, F1: 2, F2: (!((!(true && (!false))) || false))}
                                                    _ = v51
                                                case 100:
                                                    var v52 uint64 = (v3 + v3)
                                                    _ = v52
                                                case 32767:
                                                    v4 += int16(v28)
                                                }
                                            }
                                            v53 := GenS7{F0: 0, F1: 1, F2: 311}
                                            fmt.Println(fi27, v42)
                                            _ = v53
                                        } else {
                                            if len(v1) > 0 {
                                                v1[0] = (-(fi43))
                                            }
                                        }
                                        _ = v46
                                    } else {
                                        if true {
                                            v19 -= v4
                                            switch "a" {
                                            case "hello":
                                                var v54 float32 = 0.0
                                                _ = v54
                                            case "a":
                                                fmt.Println(v42)
                                            }
                                            var v55 uint8 = 183
                                            if (!true) {
                                                fmt.Println(fi43, v42)
                                            } else {
                                                v55 += (v55 ^ v28)
                                            }
                                            var v56 uint = 4294967295
                                            _ = v55
                                            _ = v56
                                        }
                                        v57 := GenS6{F0: ("foo bar" + "hello")}
                                        _ = v57
                                    }
                                } else {
                                    v28 += (v42 - v28)
                                    var v58 int32 = v18
                                    v18 = (v18 + 2147483647)
                                    _ = v58
                                }
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v59 uint32 = 2147483647
                            _ = fi43
                            _ = v59
                        }
                    } else {
                        for fi60 := 0; fi60 < 3; fi60++ {
                            if false {
                                if len(v1) > 0 {
                                    v1[0] = fi60
                                }
                                v61 := GenS2{F0: ("foo bar" + "foo bar"), F1: (fi60 << ((223) & 63)), F2: v3}
                                v26 -= 3.14159
                                _ = v61
                            }
                            var v62 uint32 = 65535
                            if (-9223372036854775808 < 42) {
                                if true {
                                    if false {
                                        var v63 int8 = 42
                                        _ = v63
                                    }
                                    switch (("a" + "\n\t") + "foo bar") {
                                    case "":
                                        var v64 uint8 = v28
                                        _ = v64
                                    case "foo bar":
                                        v62 = (v62 + (v62 - uint32(v3)))
                                    case "\n\t":
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    default:
                                        v42 = (v42 - (v28 + (v28 + (v28 >> ((2) & 63)))))
                                    }
                                }
                            }
                            _ = fi60
                            _ = v62
                        }
                    }
                } else {
                    v26 += v26
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v65 := GenS7{F0: 65535, F1: 42, F2: 65535}
                    _ = v65
                }
                _ = v42
            } else {
                if (!(false || (0 >= 42))) {
                    switch "\n\t" {
                    case "":
                        v4 = 42
                    case "a":
                        var v66 int32 = v18
                        _ = v66
                    case "foo bar":
                        v4 -= 1000
                    default:
                        var v67 uint64 = (v3 % (((^(v3))) | 1))
                        _ = v67
                    }
                } else {
                    fmt.Println(v3, v3, v4)
                    for fi68 := 0; fi68 < 14; fi68++ {
                        if len(v2) > 0 {
                            v2[0] = fi68
                        }
                        _ = fi68
                    }
                    fmt.Println(v18)
                }
                _ = len(v2)
            }
            _ = fi27
            _ = v28
        }
    }
    if false {
        switch 292 {
        case 127:
            for fi69 := 0; fi69 < 7; fi69++ {
                v2 = append(v2, fi69)
                for fi70 := 0; fi70 < 10; fi70++ {
                    v71 := GenS1{F0: 898, F1: 151, F2: false}
                    if (-818 >= (v4 % ((v19) | 1))) {
                        var v72 uint8 = 127
                        v18 = (v18 + 1)
                        _ = v72
                    }
                    switch "hello" {
                    case "":
                        if true {
                            _ = len(v2)
                            var v73 uint8 = 42
                            _ = len(v1)
                            _ = v73
                        } else {
                            v26 -= float32(0.1)
                            v2 = append(v2, (fi70 - fi70))
                        }
                        fmt.Println(fi69, v18)
                    default:
                        fmt.Println(v4, fi70, v18)
                    }
                    _ = fi70
                    _ = v71
                }
                _ = fi69
            }
        default:
            v3 -= v3
        }
        if false {
            _ = len(v2)
            fmt.Println(v3, v26, v3)
        } else {
            fmt.Println(v3, v3, v19)
        }
    }
    switch "a" {
    case "\n\t":
        v19 = (v4 << ((42) & 63))
    case "foo bar":
        var v74 uint8 = 127
        _ = v74
    case "a":
        if len(v1) > 0 {
            v1[0] = 412
        }
    default:
        var v75 int8 = -128
        _ = v75
    }
    v2 = append(v2, -9223372036854775808)
    v1 = append(v1, 127)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
