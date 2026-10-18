package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint32 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32; F1 uint32; F2 uint32 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint64; F1 int }
func genFn0(p0 bool, p1 uint) int64 {
    p0 = ((127 <= 0) || (32767 > 140))
    p1 -= (p1 >> ((p1) & 63))
    var v7 int32 = 0
    _ = p0
    _ = p1
    _ = v7
    return 286
}

func main() {
    var v1 int16 = 0
    var v2 float64 = (0.1 / 1.0)
    var v3 uint8 = 128
    var v4 int64 = 0
    var v5 int8 = 127
    v6 := []string{}
    if ("a" <= "foo bar") {
        switch "hello" {
        case "foo bar":
            v8 := GenS1{F0: 536, F1: "hello", F2: 58}
            _ = v8
        case "hello":
            for fi9 := 0; fi9 < 3; fi9++ {
                if (false || true) {
                    if ((!(false && false)) && (!true)) {
                        v10 := GenS2{F0: 900, F1: 1}
                        switch fi9 {
                        case -1:
                            fmt.Println(v2)
                        default:
                            switch "a" {
                            case "foo bar":
                                v5 -= (v5 + (v5 - v5))
                            case "":
                                var v11 string = ("a" + "\n\t")
                                _ = v11
                            case "hello":
                                if false {
                                    v12 := GenS2{F0: uint64(v3), F1: 65535}
                                    switch "\n\t" {
                                    case "":
                                        fmt.Println(v1, v1, v3)
                                    }
                                    _ = v12
                                } else {
                                    var v13 int32 = 42
                                    _ = len(v6)
                                    v14 := GenS3{F0: ""}
                                    _ = v13
                                    _ = v14
                                }
                            }
                        }
                        _ = v10
                    }
                    if false {
                        var v15 int32 = -1
                        v6 = append(v6, "foo bar")
                        v16 := GenS1{F0: 2147483647, F1: "", F2: 4294967295}
                        _ = v15
                        _ = v16
                    } else {
                        if len(v6) > 0 {
                            v6[0] = "foo bar"
                        }
                        var v17 int8 = 100
                        _ = v17
                    }
                    if true {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                } else {
                    if len(v6) > 0 {
                        v6[0] = "a"
                    }
                    if (fi9 < fi9) {
                        for fi18 := 0; fi18 < 14; fi18++ {
                            v1 += v1
                            _ = fi18
                        }
                        switch ("foo bar" + "") {
                        case "foo bar":
                            var v19 int = fi9
                            _ = v19
                        case "a":
                            v3 += v3
                        case "hello":
                            fmt.Println(v4, fi9)
                        }
                    }
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(v1, v4)
                    case "hello":
                        if len(v6) > 0 {
                            v6[0] = "hello"
                        }
                    case "a":
                        v1 = v1
                    }
                }
                var v20 uint8 = v3
                var v21 int16 = (-(v1))
                _ = fi9
                _ = v20
                _ = v21
            }
        }
        v2 = float64(3.14159)
        if (("" + ("hello" + (((("a" + "\n\t") + "") + "a") + "hello"))) != "\n\t") {
            v6 = append(v6, "\n\t")
            var v22 uint32 = 0
            v2 += 1e10
            _ = v22
        }
    }
    if ((-788 > -2147483648) || (2 >= 42)) {
        fmt.Println(v2, v2, v5)
        v4 += (v4 - (v4 + -1))
    } else {
        if true {
            if len(v6) > 0 {
                v6[0] = "a"
            }
            if false {
                v23 := GenS4{F0: -2147483648, F1: (v3 | (v3 % ((v3) | 1)))}
                _ = v23
            } else {
                fmt.Println(v2, v3, v1)
            }
            if false {
                v24 := GenS2{F0: 4294967295, F1: 42}
                _ = v24
            } else {
                v2 -= v2
                var v25 string = "\n\t"
                switch 42 {
                case 127:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                default:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                _ = v25
            }
        }
        switch 235 {
        case 884:
            v26 := GenS5{F0: 65535}
            _ = v26
        }
        v2 -= (v2 * v2)
    }
    var v27 uint16 = 899
    var v28 uint32 = 42
    if false {
        if ((v28 << ((uint(v28)) & 63)) <= v28) {
            var v29 int16 = (^(v1))
            switch "\n\t" {
            case "hello":
                v30 := GenS6{F0: 2147483647, F1: 1, F2: v28}
                _ = v30
            case "\n\t":
                var v31 uint16 = 254
                _ = v31
            case "foo bar":
                var v32 int8 = v5
                _ = v32
            }
            _ = v29
        } else {
            v33 := GenS7{F0: 456}
            for fi34 := 0; fi34 < 2; fi34++ {
                if (v2 >= (v2 - (0.0 / v2))) {
                    var v35 uint16 = (v27 ^ (v27 + (v27 & uint16(v3))))
                    _ = v35
                }
                for fi36 := 0; fi36 < 19; fi36++ {
                    if (-2147483648 == -2147483648) {
                        v3 = uint8(v28)
                        var v37 string = "\n\t"
                        var v38 int = 397
                        _ = v37
                        _ = v38
                    } else {
                        switch 0 {
                        case 1:
                            switch ("a" + "") {
                            case "\n\t":
                                switch (fi36 * fi34) {
                                case -449:
                                    v39 := GenS3{F0: "foo bar"}
                                    _ = v39
                                case -2147483648:
                                    var v40 int32 = -674
                                    _ = v40
                                }
                            case "foo bar":
                                v4 += (v4 / ((-1) | 1))
                            }
                        case 42:
                            switch "\n\t" {
                            case "foo bar":
                                var v41 uint64 = 568
                                _ = v41
                            case "\n\t":
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            default:
                                var v42 uint8 = v3
                                _ = v42
                            }
                        }
                        v43 := GenS6{F0: -2147483648, F1: (v28 >> ((301) & 63)), F2: (^(v28))}
                        _ = v43
                    }
                    if (!(-2.25 == -2.25)) {
                        if ((v5 % ((v5) | 1)) < (v5 ^ v5)) {
                            var v44 uint8 = (v3 << ((65535) & 63))
                            v45 := GenS1{F0: -1, F1: "foo bar", F2: v28}
                            if (0 == 2) {
                                v6 = append(v6, "hello")
                            }
                            _ = v44
                            _ = v45
                        }
                        v5 -= -125
                        var v46 int16 = 1000
                        _ = v46
                    } else {
                        fmt.Println(v2)
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    if (((^(v27)) > (^(v27))) && ((v5 % ((v5) | 1)) < (v5 << ((232) & 63)))) {
                        if len(v6) > 0 {
                            v6[0] = "foo bar"
                        }
                        var v47 bool = (!(((!false) || false) || (false && true)))
                        var v48 uint = uint(v3)
                        _ = v47
                        _ = v48
                    }
                    _ = fi36
                }
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi34
            }
            for fi49 := 0; fi49 < 11; fi49++ {
                v50 := GenS8{F0: (v4 << ((65535) & 63))}
                v6 = append(v6, "")
                _ = fi49
                _ = v50
            }
            _ = v33
        }
        for fi51 := 0; fi51 < 6; fi51++ {
            if true {
                var v52 int8 = -95
                _ = v52
            } else {
                if (!(v1 > v1)) {
                    var v53 bool = (v2 < v2)
                    v4 = (v4 * v4)
                    _ = len(v6)
                    _ = v53
                } else {
                    for fi54 := 0; fi54 < 8; fi54++ {
                        if len(v6) > 0 {
                            v6[0] = "foo bar"
                        }
                        v3 += (v3 >> ((1) & 63))
                        fmt.Println(fi54)
                        switch ("foo bar" + "a") {
                        case "a":
                            var v55 uint32 = 42
                            _ = v55
                        case "hello":
                            switch (("a" + "a") + "a") {
                            case "a":
                                var v56 float64 = (1.5 * 3.14159)
                                _ = v56
                            case "hello":
                                var v57 float64 = ((v2 / 1.5) / -2.25)
                                _ = v57
                            case "":
                                var v58 uint16 = v27
                                _ = v58
                            }
                        }
                        var v59 string = "hello"
                        _ = fi54
                        _ = v59
                    }
                    var v60 float64 = v2
                    if true {
                        v60 = v2
                    }
                    _ = v60
                }
                v61 := GenS1{F0: 461, F1: ("" + (("\n\t" + "a") + ((((("" + ("\n\t" + "\n\t")) + "") + "foo bar") + "foo bar") + ("\n\t" + ("hello" + "hello"))))), F2: (^(v28))}
                _ = v61
            }
            v62 := GenS7{F0: v27}
            for fi63 := 0; fi63 < 8; fi63++ {
                var v64 int64 = -9223372036854775808
                var v65 uint32 = v28
                v6 = append(v6, "a")
                _ = fi63
                _ = v64
                _ = v65
            }
            _ = fi51
            _ = v62
        }
        if true {
            v2 = (1.5 * (0.0 / float64(-2.25)))
            v1 = v1
            v66 := GenS8{F0: v4}
            _ = v66
        } else {
            v3 = (^(v3))
            for fi67 := 0; fi67 < 9; fi67++ {
                v68 := GenS5{F0: (v27 << ((42) & 63))}
                var v69 uint = 217
                v70 := GenS2{F0: uint64(v3), F1: (v69 ^ v69)}
                _ = fi67
                _ = v68
                _ = v69
                _ = v70
            }
        }
    }
    switch "hello" {
    case "":
        _ = len(v6)
    case "\n\t":
        v71 := GenS9{F0: 18446744073709551615, F1: 0}
        _ = v71
    case "a":
        var v72 int16 = 0
        _ = v72
    }
    var v73 bool = true
    for fi74 := 0; fi74 < 1; fi74++ {
        if true {
            var v75 float64 = float64(3.14159)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = v75
        }
        fmt.Println(v2)
        _ = fi74
    }
    var v76 float64 = ((3.14159 - float64(3.14159)) + 3.14159)
    fmt.Println(v3, v5, v73)
    for fi77 := 0; fi77 < 4; fi77++ {
        v78 := GenS1{F0: 680, F1: "hello", F2: (v28 - v28)}
        var v79 uint = 18446744073709551615
        _ = fi77
        _ = v78
        _ = v79
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v73:", v73)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v76
}
