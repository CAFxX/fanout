package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int64; F1 string; F2 uint }
type GenS5 struct { F0 int16; F1 int64; F2 bool }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int32; F1 string }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int64; F1 uint64 }
type GenS10 struct { F0 int8; F1 string }
type GenS11 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]string)
    var v3 float32 = (0.0 - 3.14159)
    if (65535 > 2) {
        delete(v1, 1)
        v4 := GenS1{F0: 0}
        _ = v4
    } else {
        v5 := GenS1{F0: 1}
        if false {
            fmt.Println(v3)
        }
        if false {
            for fi6 := 0; fi6 < 16; fi6++ {
                v7 := GenS1{F0: 620}
                _ = fi6
                _ = v7
            }
        }
        _ = v5
    }
    v8 := GenS1{F0: 715}
    if (904 >= 2) {
        var v9 uint = 4294967295
        _ = v9
    } else {
        v10 := GenS2{F0: 0, F1: 42}
        v3 -= 1.5
        _ = v10
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := GenS2{F0: 32767, F1: -731}
        var v13 float64 = 1.5
        if true {
            if false {
                if true {
                    v14 := []int{(fi11 / ((329) | 1)), (fi11 - (fi11 - (-(fi11)))), (fi11 ^ fi11)}
                    var v15 int = (-(fi11))
                    _ = v14
                    _ = v15
                }
            }
            v16 := GenS2{F0: 500, F1: 1000}
            _ = v16
        } else {
            var v17 uint = 1
            _ = v17
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    if ("a" >= "\n\t") {
        var v18 uint = 1
        _ = v18
    } else {
        switch 948 {
        case 32768:
            switch ((("a" + ("foo bar" + ("" + (("hello" + "\n\t") + (("" + "foo bar") + "a"))))) + "hello") + "foo bar") {
            case "hello":
                v3 = 3.14159
            case "":
                v3 = v3
            default:
                v3 = ((v3 / 0.1) - 1.5)
            }
        case 128:
            for fi19 := 0; fi19 < 9; fi19++ {
                switch fi19 {
                case 805:
                    delete(v2, "a")
                case 32768:
                    v20 := GenS1{F0: 65535}
                    _ = v20
                case 2147483647:
                    var v21 uint = 819
                    _ = v21
                default:
                    if true {
                        fmt.Println(v3, fi19)
                        v3 = v3
                    } else {
                        for fi22 := 0; fi22 < 13; fi22++ {
                            var v23 int64 = 1
                            _ = fi22
                            _ = v23
                        }
                        var v24 float32 = 1e10
                        _ = v24
                    }
                }
                v25 := GenS1{F0: 42}
                if ((-109 > 0) || (1 >= 152)) {
                    for fi26 := 0; fi26 < 2; fi26++ {
                        v3 = v3
                        _ = fi26
                    }
                } else {
                    v27 := GenS3{F0: 1, F1: 1}
                    var v28 int64 = 42
                    _ = v27
                    _ = v28
                }
                _ = fi19
                _ = v25
            }
        case 1:
            var v29 int = 32768
            _ = v29
        default:
            for fi30 := 0; fi30 < 7; fi30++ {
                v31 := GenS2{F0: 0, F1: 1000}
                v3 = v3
                _ = fi30
                _ = v31
            }
        }
        for fi32 := 0; fi32 < 11; fi32++ {
            if false {
                if true {
                    for fi33 := 0; fi33 < 13; fi33++ {
                        var v34 float32 = 0.1
                        var v35 int16 = -32768
                        var v36 int16 = (v35 & (v35 | v35))
                        _ = fi33
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                    if true {
                        var v37 string = "a"
                        _ = v37
                    }
                } else {
                    v38 := GenS4{F0: 42, F1: "foo bar", F2: 42}
                    for fi39 := 0; fi39 < 15; fi39++ {
                        switch "foo bar" {
                        case "a":
                            var v40 string = ("a" + "\n\t")
                            _ = v40
                        case "\n\t":
                            _ = len(v1)
                        default:
                            v3 -= ((0.1 / v3) / v3)
                        }
                        v3 = v3
                        _ = fi39
                    }
                    v41 := GenS5{F0: -254, F1: -304, F2: (false && (18446744073709551615 <= 18446744073709551615))}
                    _ = v38
                    _ = v41
                }
                for fi42 := 0; fi42 < 16; fi42++ {
                    var v43 uint = 2
                    _ = fi42
                    _ = v43
                }
            }
            _ = fi32
        }
        if (!(-876 <= -496)) {
            var v44 int = -1
            switch "foo bar" {
            case "hello":
                var v45 uint8 = 2
                _ = v45
            }
            for fi46 := 0; fi46 < 19; fi46++ {
                for fi47 := 0; fi47 < 6; fi47++ {
                    var v48 uint32 = 2147483647
                    v49 := GenS6{F0: 42}
                    _ = fi47
                    _ = v48
                    _ = v49
                }
                v50 := v2["hello"]
                v3 -= 0.0
                _ = fi46
                _ = v50
            }
            _ = v44
        } else {
            v3 -= 1.5
        }
    }
    var v51 uint64 = 777
    var v52 int16 = 0
    v3 = v3
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        v53 := v2[("" + "foo bar")]
        _ = v53
    case "a":
        v51 = (v51 & (^(v51)))
    case "":
        v3 -= (v3 * ((float32(3.14159) - (v3 / float32(0.1))) - 0.0))
    default:
        for fi54 := 0; fi54 < 9; fi54++ {
            fmt.Println(v52)
            _ = fi54
        }
    }
    if ((v51 + v51) < v51) {
        v55 := []int{-9223372036854775808}
        var v56 float32 = (v3 * -2.25)
        fmt.Println(v51, v52, v52)
        _ = v55
        _ = v56
    } else {
        switch ("\n\t" + ("a" + "")) {
        case "hello":
            v3 = v3
        case "foo bar":
            v57 := GenS7{F0: 1, F1: "\n\t"}
            _ = v57
        case "":
            if true {
                if ((228 != 255) && (!((v52 ^ 1) != 0))) {
                    if false {
                        var v58 string = "hello"
                        _ = v58
                    } else {
                        v51 = v51
                    }
                    switch "hello" {
                    case "a":
                        v59 := GenS8{F0: 42}
                        _ = v59
                    case "\n\t":
                        v52 -= (v52 / ((v52) | 1))
                    default:
                        v60 := GenS5{F0: (v52 << ((65535) & 63)), F1: 812, F2: (v51 > 0)}
                        _ = v60
                    }
                    var v61 uint16 = 42
                    _ = v61
                }
                for fi62 := 0; fi62 < 16; fi62++ {
                    var v63 uint32 = 4294967295
                    var v64 float64 = 3.14159
                    fmt.Println(v64, v52)
                    _ = fi62
                    _ = v63
                    _ = v64
                }
                var v65 string = ""
                _ = v65
            } else {
                for fi66 := 0; fi66 < 17; fi66++ {
                    v67 := []int{}
                    switch "\n\t" {
                    case "":
                        if false {
                            v52 += 307
                            var v68 float32 = -2.25
                            for fi69 := 0; fi69 < 7; fi69++ {
                                if (1 != 0) {
                                    switch "" {
                                    case "hello":
                                        fmt.Println(fi66, v68, v51)
                                    case "foo bar":
                                        v70 := GenS3{F0: 844, F1: 0}
                                        _ = v70
                                    case "":
                                        var v71 uint16 = 65535
                                        _ = v71
                                    default:
                                        v3 = v68
                                    }
                                    if len(v67) >= 2 {
                                        v67 = v67[1:len(v67)]
                                    }
                                    if len(v67) > 0 {
                                        v67[0] = fi66
                                    }
                                } else {
                                    switch "\n\t" {
                                    case "a":
                                        v3 -= ((0.1 + v68) + 1.5)
                                    case "\n\t":
                                        fmt.Println(v68, fi66)
                                    case "":
                                        fmt.Println(v3)
                                    }
                                }
                                var v72 uint64 = v51
                                var v73 int32 = 42
                                var v74 int16 = (-(v52))
                                _ = fi69
                                _ = v72
                                _ = v73
                                _ = v74
                            }
                            _ = v68
                        } else {
                            for fi75 := 0; fi75 < 7; fi75++ {
                                var v76 int = (fi75 + -997)
                                v77 := GenS9{F0: 702, F1: v51}
                                _ = fi75
                                _ = v76
                                _ = v77
                            }
                        }
                    case "foo bar":
                        v51 -= v51
                    case "hello":
                        for fi78 := 0; fi78 < 4; fi78++ {
                            fmt.Println(v3)
                            var v79 uint = 636
                            _ = fi78
                            _ = v79
                        }
                    }
                    switch "\n\t" {
                    case "hello":
                        v80 := GenS5{F0: (v52 >> ((26) & 63)), F1: 42, F2: true}
                        _ = v80
                    case "a":
                        for fi81 := 0; fi81 < 9; fi81++ {
                            v82 := GenS9{F0: -9223372036854775808, F1: v51}
                            v51 = (v51 - (v51 % (((^(v51))) | 1)))
                            v51 += v51
                            _ = fi81
                            _ = v82
                        }
                    default:
                        _ = len(v67)
                    }
                    _ = fi66
                    _ = v67
                }
            }
        default:
            if false {
                v52 = (v52 | v52)
                if true {
                    fmt.Println(v3, v52)
                    if false {
                        v52 = (v52 - v52)
                    } else {
                        var v83 uint8 = 2
                        for fi84 := 0; fi84 < 5; fi84++ {
                            v51 = v51
                            v85 := []int{fi84, fi84, fi84}
                            _ = fi84
                            _ = v85
                        }
                        _ = v83
                    }
                } else {
                    if (v52 < (v52 - v52)) {
                        _ = len(v1)
                    } else {
                        for fi86 := 0; fi86 < 12; fi86++ {
                            var v87 int8 = 103
                            for fi88 := 0; fi88 < 9; fi88++ {
                                if false {
                                    v89 := GenS10{F0: (v87 % ((127) | 1)), F1: ("hello" + (((("hello" + "hello") + (("" + "hello") + "foo bar")) + ("a" + "foo bar")) + ((("hello" + ((("a" + "foo bar") + "hello") + "a")) + "foo bar") + "foo bar")))}
                                    switch (fi86 + fi86) {
                                    case 2147483647:
                                        fmt.Println(v3)
                                    case 32767:
                                        var v90 uint16 = 0
                                        _ = v90
                                    }
                                    v87 -= v87
                                    fmt.Println(v3)
                                    _ = v89
                                } else {
                                    v91 := GenS11{F0: 167}
                                    v1[2147483647] = ((((("foo bar" + "foo bar") + "\n\t") + ("a" + "a")) + "\n\t") >= "a")
                                    _ = v91
                                }
                                v51 = v51
                                v92 := GenS8{F0: 2147483648}
                                fmt.Println(v87, v52, fi88)
                                var v93 int16 = v52
                                _ = fi88
                                _ = v92
                                _ = v93
                            }
                            switch "\n\t" {
                            case "a":
                                v94 := GenS8{F0: 18}
                                _ = v94
                            case "":
                                v95 := v1[4294967295]
                                _ = v95
                            case "hello":
                                var v96 int32 = 1
                                _ = v96
                            default:
                                v97 := GenS3{F0: 1, F1: (^(v51))}
                                _ = v97
                            }
                            _ = fi86
                            _ = v87
                        }
                        fmt.Println(v3, v3)
                    }
                    v3 = 1e10
                }
                var v98 int = 4294967295
                _ = v98
            } else {
                delete(v1, 65535)
                var v99 string = "foo bar"
                _ = v99
            }
        }
    }
    v51 = (^(v51))
    v100 := []int{32767, 0, 257}
    if true {
        fmt.Println(v52, v3, v51)
        fmt.Println(v3, v52)
        fmt.Println(v52, v52)
    } else {
        var v101 uint64 = v51
        var v102 int32 = 2147483647
        _ = v101
        _ = v102
    }
    v52 = v52
    var v103 int32 = -2147483648
    for fi104 := 0; fi104 < 19; fi104++ {
        switch (("a" + ((("foo bar" + "a") + "foo bar") + "a")) + ("" + ("\n\t" + ("\n\t" + ((("" + "") + "foo bar") + "a"))))) {
        case "a":
            v3 -= 0.1
        case "\n\t":
            _ = len(v100)
        case "hello":
            v103 += v103
        }
        _ = fi104
    }
    v103 += v103
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v100:", len(v100), v100)
    fmt.Println("v103:", v103)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v100
    _ = v103
    _ = v2
    _ = v3
    _ = v51
    _ = v52
    _ = v8
}
