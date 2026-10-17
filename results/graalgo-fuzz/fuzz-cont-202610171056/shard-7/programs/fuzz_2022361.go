package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint64; F1 string; F2 string }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS6 struct { F0 bool; F1 uint64; F2 int32 }
type GenS7 struct { F0 int64; F1 int32; F2 int16 }
type GenS8 struct { F0 string; F1 bool }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint64 }
func genFn0(p0 uint32) int8 {
    p0 += p0
    p0 += (p0 / ((1) | 1))
    v6 := []int{}
    var v7 uint64 = 130
    if false {
        var v8 uint64 = v7
        v9 := GenS1{F0: 127, F1: (v8 >> ((65535) & 63)), F2: p0}
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    return 42
}
func genFn1() string {
    var v10 uint64 = 408
    v11 := GenS1{F0: -20, F1: (v10 | v10), F2: 1}
    v10 = 913
    switch "hello" {
    case "hello":
        var v12 int = 4294967295
        _ = v12
    case "foo bar":
        if false {
            switch 873 {
            case -9223372036854775808:
                v10 += (^(v10))
            case 9223372036854775807:
                if (!true) {
                    v13 := []int{-271, 42}
                    v10 += 4294967295
                    v14 := GenS2{F0: v10, F1: "hello", F2: ((((("" + "a") + "\n\t") + "") + (("hello" + "\n\t") + "foo bar")) + ("" + ("\n\t" + ("" + ("\n\t" + ("foo bar" + ""))))))}
                    _ = v13
                    _ = v14
                }
            case -2147483648:
                var v15 int64 = 42
                _ = v15
            default:
                for fi16 := 0; fi16 < 2; fi16++ {
                    v17 := GenS3{F0: 1, F1: 438}
                    switch (fi16 * (fi16 / ((127) | 1))) {
                    case 4294967295:
                        var v18 int64 = -9223372036854775808
                        _ = v18
                    }
                    v19 := GenS1{F0: 43, F1: 18446744073709551615, F2: 1}
                    _ = fi16
                    _ = v17
                    _ = v19
                }
            }
        } else {
            fmt.Println(v10)
            var v20 bool = (!(!false))
            if (-1 == 9223372036854775807) {
                v21 := GenS2{F0: v10, F1: "\n\t", F2: ("a" + "foo bar")}
                var v22 float32 = 0.1
                v22 -= ((0.0 - 0.0) / (-2.25 + v22))
                _ = v21
                _ = v22
            } else {
                v23 := GenS4{F0: -85}
                _ = v23
            }
            _ = v20
        }
    default:
        v10 = v10
    }
    var v24 uint = 968
    _ = v10
    _ = v11
    _ = v24
    return "a"
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]string)
    var v3 uint = 898
    var v4 uint16 = 255
    var v5 int8 = 0
    var v25 string = "hello"
    fmt.Println(v25)
    switch -203 {
    case 32768:
        v26 := GenS3{F0: -9223372036854775808, F1: 0}
        _ = v26
    case -1:
        v25 = ("" + (((v25 + v25) + (v25 + v25)) + "a"))
    default:
        for fi27 := 0; fi27 < 15; fi27++ {
            fmt.Println(v4)
            for fi28 := 0; fi28 < 1; fi28++ {
                if ((0.1 == 0.0) && ((v3 ^ v3) < v3)) {
                    v29 := v2[("" + "foo bar")]
                    var v30 int64 = -464
                    _ = v29
                    _ = v30
                }
                var v31 uint32 = 0
                v32 := GenS5{F0: 127, F1: (v3 % ((1) | 1)), F2: (v4 << ((v3) & 63))}
                v33 := GenS6{F0: true, F1: 2, F2: 0}
                v34 := GenS3{F0: -9223372036854775808, F1: (v31 - (v31 >> (((v3 | uint(v4))) & 63)))}
                _ = fi28
                _ = v31
                _ = v32
                _ = v33
                _ = v34
            }
            _ = fi27
        }
    }
    for fi35 := 0; fi35 < 11; fi35++ {
        v36 := []int{fi35, fi35, (^(fi35))}
        var v37 string = "a"
        _ = fi35
        _ = v36
        _ = v37
    }
    v38 := GenS1{F0: (v5 >> ((v3) & 63)), F1: 42, F2: 2}
    switch -462 {
    case 2147483647:
        for fi39 := 0; fi39 < 4; fi39++ {
            v5 -= 1
            var v40 int16 = int16(v5)
            _ = fi39
            _ = v40
        }
    }
    if ((v3 <= 0) && (108 == 2)) {
        fmt.Println(v5, v5)
    }
    var v41 bool = (4294967295 <= 192)
    if v41 {
        var v42 float64 = (((1.5 + 1e10) - 0.0) * 0.1)
        _ = v42
    } else {
        var v43 uint = (v3 * 18446744073709551615)
        var v44 string = ("a" + v25)
        if false {
            if v41 {
                if (!true) {
                    var v45 string = ("\n\t" + ((("\n\t" + "") + ((v25 + "a") + (("hello" + ("a" + "\n\t")) + v25))) + (((v25 + v25) + "") + "foo bar")))
                    var v46 uint64 = 573
                    _ = v45
                    _ = v46
                } else {
                    for fi47 := 0; fi47 < 7; fi47++ {
                        var v48 int8 = (v5 & 1)
                        v49 := []int{fi47, fi47, (fi47 | 9223372036854775807)}
                        _ = fi47
                        _ = v48
                        _ = v49
                    }
                    v2[(v25 + v44)] = (("\n\t" + "foo bar") + (v25 + v25))
                    for fi50 := 0; fi50 < 4; fi50++ {
                        v5 = (v5 & 38)
                        switch 1 {
                        case -1:
                            fmt.Println(v3)
                        case -349:
                            v3 += (v3 << (((v43 >> ((v43) & 63))) & 63))
                        default:
                            fmt.Println(v41, v4)
                        }
                        _ = fi50
                    }
                }
                if v41 {
                    if true {
                        for fi51 := 0; fi51 < 3; fi51++ {
                            v52 := []int{fi51, (fi51 ^ fi51), (^(fi51)), 32767}
                            v41 = true
                            _ = fi51
                            _ = v52
                        }
                        fmt.Println(v3, v25)
                        switch (v25 + ((v44 + "hello") + "\n\t")) {
                        case "":
                            switch ((v44 + v44) + v25) {
                            case "foo bar":
                                v3 = (v3 ^ v43)
                            }
                        case "a":
                            fmt.Println(v3, v5, v25)
                        default:
                            v53 := GenS7{F0: 42, F1: -2147483648, F2: 0}
                            _ = v53
                        }
                    } else {
                        switch 2147483647 {
                        case 32768:
                            fmt.Println(v41)
                        case -1:
                            v4 = (v4 * (^(v4)))
                        }
                        v43 = v3
                    }
                    switch (((v44 + v44) + "a") + v25) {
                    case "hello":
                        var v54 int8 = (v5 % ((v5) | 1))
                        _ = v54
                    case "\n\t":
                        for fi55 := 0; fi55 < 9; fi55++ {
                            v56 := GenS3{F0: -893, F1: 65535}
                            var v57 int = fi55
                            v58 := GenS8{F0: v25, F1: v41}
                            _ = fi55
                            _ = v56
                            _ = v57
                            _ = v58
                        }
                    default:
                        var v59 float64 = 0.1
                        _ = v59
                    }
                } else {
                    if true {
                        if (!(!false)) {
                            fmt.Println(v44, v43)
                            v60 := GenS9{F0: 2147483647}
                            _ = v60
                        } else {
                            var v61 bool = v41
                            var v62 uint8 = 157
                            _ = v61
                            _ = v62
                        }
                        switch "\n\t" {
                        case "":
                            v63 := []int{4294967295}
                            _ = v63
                        case "foo bar":
                            for fi64 := 0; fi64 < 19; fi64++ {
                                _ = len(v2)
                                var v65 int8 = (v5 + genFn0(270))
                                var v66 string = v25
                                _ = fi64
                                _ = v65
                                _ = v66
                            }
                        case "\n\t":
                            var v67 uint = (v3 * 702)
                            _ = v67
                        default:
                            fmt.Println(v44, v43)
                        }
                    }
                    v3 = 0
                }
            } else {
                for fi68 := 0; fi68 < 7; fi68++ {
                    v69 := GenS5{F0: 128, F1: (v3 ^ (v3 & (^(v43)))), F2: v4}
                    v70 := GenS10{F0: 4294967295}
                    v43 -= v43
                    _ = fi68
                    _ = v69
                    _ = v70
                }
                for fi71 := 0; fi71 < 16; fi71++ {
                    v72 := []int{1, 0}
                    _ = fi71
                    _ = v72
                }
                v73 := []int{134, 32768}
                _ = v73
            }
            v74 := []int{137, -9223372036854775808, -525}
            var v75 int16 = 162
            _ = v74
            _ = v75
        }
        _ = v43
        _ = v44
    }
    fmt.Println(v5, v3, v41)
    var v76 int64 = 247
    for fi77 := 0; fi77 < 7; fi77++ {
        if ((v25 != ("a" + v25)) && v41) {
            fmt.Println(v25, v4)
            switch (fi77 + fi77) {
            case 127:
                fmt.Println(v4)
            }
            v41 = true
        } else {
            v41 = v41
            switch (fi77 * fi77) {
            case 32768:
                v78 := []int{(fi77 | fi77), (^(fi77)), (fi77 - fi77)}
                _ = v78
            case -2147483648:
                var v79 uint8 = 26
                _ = v79
            default:
                for fi80 := 0; fi80 < 4; fi80++ {
                    if (267 == (v3 >> (((^(v3))) & 63))) {
                        v81 := GenS6{F0: false, F1: 18446744073709551615, F2: -848}
                        v82 := []int{(fi80 | (fi80 & 189)), (-(fi80)), (fi80 | fi80), fi77}
                        var v83 string = v25
                        v41 = false
                        var v84 int8 = 42
                        _ = v81
                        _ = v82
                        _ = v83
                        _ = v84
                    }
                    v85 := []int{(fi80 % (((-(fi77))) | 1)), (fi77 & fi80), fi77, 2147483647}
                    switch v25 {
                    case "":
                        v25 = (("hello" + "") + ((v25 + v25) + v25))
                        v2[v25] = v25
                        var v86 uint8 = 127
                        _ = v86
                    default:
                        v87 := GenS6{F0: (!true), F1: 0, F2: 2147483647}
                        _ = v87
                    }
                    _ = fi80
                    _ = v85
                }
            }
            v88 := []int{(fi77 | int(v76)), 127, fi77}
            _ = v88
        }
        _ = fi77
    }
    for fi89 := 0; fi89 < 5; fi89++ {
        v41 = v41
        var v90 int = (fi89 * (fi89 & fi89))
        _ = fi89
        _ = v90
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v38
    _ = v4
    _ = v41
    _ = v5
    _ = v76
}
