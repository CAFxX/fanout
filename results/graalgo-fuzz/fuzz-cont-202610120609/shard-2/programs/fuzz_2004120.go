package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 bool; F1 int16; F2 int32 }
type GenS3 struct { F0 uint8; F1 int16; F2 int64 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 uint16; F1 int64; F2 int }
func genFn0(p0 int16, p1 int, p2 uint64) string {
    var v7 int = (p1 / ((p1) | 1))
    p0 = 1
    if ((-558 > 2147483647) && ((p0 * 42) != p0)) {
        var v8 uint64 = (p2 | (p2 * 354))
        _ = v8
    } else {
        if false {
            for fi9 := 0; fi9 < 18; fi9++ {
                v10 := []int{}
                _ = fi9
                _ = v10
            }
        }
    }
    if true {
        var v11 int8 = 100
        _ = v11
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 int = (v7 % (((fi12 >> ((210) & 63))) | 1))
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "foo bar"
}
func genFn1(p0 uint64) int {
    p0 += p0
    var v14 int8 = 1
    _ = p0
    _ = v14
    return 1
}

func main() {
    var v1 int16 = -341
    var v2 uint = 352
    var v3 int8 = 100
    var v4 int16 = 32767
    var v5 int16 = v4
    v6 := []string{(("a" + ("hello" + "hello")) + (("" + "hello") + ((("a" + "foo bar") + "\n\t") + "a")))}
    for fi15 := 0; fi15 < 15; fi15++ {
        for fi16 := 0; fi16 < 8; fi16++ {
            v17 := GenS1{F0: 32767, F1: 42}
            _ = fi16
            _ = v17
        }
        var v18 uint8 = 128
        _ = fi15
        _ = v18
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if true {
        v6 = append(v6, ("foo bar" + "\n\t"))
    }
    if true {
        if (0.1 == (0.0 / -2.25)) {
            var v19 int16 = v5
            _ = v19
        } else {
            fmt.Println(v5, v4)
        }
        if (56 > -698) {
            v5 += (^(v5))
            var v20 uint = v2
            fmt.Println(v5)
            _ = v20
        } else {
            for fi21 := 0; fi21 < 9; fi21++ {
                switch fi21 {
                case 4294967295:
                    v4 = 32767
                case -2147483648:
                    for fi22 := 0; fi22 < 16; fi22++ {
                        var v23 uint8 = 219
                        v24 := GenS2{F0: true, F1: v1, F2: -2147483648}
                        switch "foo bar" {
                        case "hello":
                            v6 = append(v6, "hello")
                        case "":
                            v25 := GenS3{F0: 231, F1: (^(v1)), F2: 225}
                            _ = v25
                        case "foo bar":
                            v6 = append(v6, "\n\t")
                        default:
                            if true {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                _ = len(v6)
                                v23 += 42
                            } else {
                                var v26 float64 = 1e10
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                _ = v26
                            }
                        }
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                default:
                    v6 = append(v6, ("" + ((("a" + "hello") + "\n\t") + ("hello" + ("hello" + "foo bar")))))
                }
                if ((v3 < (v3 + -35)) && ((384 == (v2 ^ (^(v2)))) || true)) {
                    var v27 uint64 = 4294967295
                    if false {
                        var v28 int64 = -1
                        for fi29 := 0; fi29 < 15; fi29++ {
                            v4 += v1
                            v30 := GenS3{F0: 127, F1: 0, F2: (v28 | (v28 + v28))}
                            _ = fi29
                            _ = v30
                        }
                        switch "\n\t" {
                        case "foo bar":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        case "\n\t":
                            var v31 uint32 = 839
                            _ = v31
                        case "":
                            for fi32 := 0; fi32 < 10; fi32++ {
                                var v33 bool = false
                                v27 += 1
                                v34 := GenS2{F0: (2147483647 == 0), F1: v4, F2: 42}
                                switch int(v2) {
                                case 128:
                                    var v35 uint32 = 655
                                    _ = v35
                                }
                                _ = fi32
                                _ = v33
                                _ = v34
                            }
                        }
                        _ = v28
                    } else {
                        if true {
                            v4 = v1
                            v3 -= (v3 - (v3 % (((-(v3))) | 1)))
                            fmt.Println(v4, fi21, fi21)
                        }
                    }
                    _ = v27
                } else {
                    v4 = v1
                    v36 := GenS3{F0: uint8(v2), F1: -152, F2: -1}
                    _ = v36
                }
                for fi37 := 0; fi37 < 2; fi37++ {
                    if true {
                        v4 = (v5 | (v1 ^ v1))
                    } else {
                        var v38 uint32 = 2147483647
                        _ = v38
                    }
                    if true {
                        v5 = v1
                        v6 = append(v6, ("\n\t" + "\n\t"))
                        if true {
                            v5 = v4
                        } else {
                            fmt.Println(v5)
                            var v39 uint16 = 65535
                            _ = v39
                        }
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = fi37
                }
                _ = fi21
            }
            var v40 int16 = (v1 % ((-1) | 1))
            v2 = 205
            _ = v40
        }
    } else {
        for fi41 := 0; fi41 < 3; fi41++ {
            if (("foo bar" + "") != "hello") {
                for fi42 := 0; fi42 < 11; fi42++ {
                    v5 = v4
                    var v43 uint64 = 4294967295
                    _ = fi42
                    _ = v43
                }
                fmt.Println(v2, v3, v4)
                for fi44 := 0; fi44 < 14; fi44++ {
                    var v45 int16 = (v5 ^ v5)
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    var v46 uint32 = 2147483648
                    _ = fi44
                    _ = v45
                    _ = v46
                }
            }
            for fi47 := 0; fi47 < 18; fi47++ {
                if (144 <= 127) {
                    v2 = (v2 ^ v2)
                    var v48 uint8 = 2
                    _ = v48
                } else {
                    fmt.Println(fi47, v4)
                }
                _ = len(v6)
                v49 := GenS4{F0: v2, F1: 0}
                _ = fi47
                _ = v49
            }
            for fi50 := 0; fi50 < 12; fi50++ {
                var v51 int = (fi50 + 2147483647)
                if true {
                    _ = len(v6)
                    v5 -= (v1 << (((v2 >> (((^(v2))) & 63))) & 63))
                } else {
                    _ = len(v6)
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v6 = append(v6, ("" + "hello"))
                }
                var v52 bool = true
                _ = fi50
                _ = v51
                _ = v52
            }
            _ = fi41
        }
        var v53 uint = 18446744073709551615
        _ = v53
    }
    var v54 int64 = 915
    var v55 uint32 = 299
    for fi56 := 0; fi56 < 20; fi56++ {
        if true {
            switch "hello" {
            case "":
                if (false || ("\n\t" >= "foo bar")) {
                    var v57 uint32 = (^(v55))
                    _ = v57
                } else {
                    var v58 uint = 16
                    if len(v6) > 0 {
                        v6[0] = ("a" + "hello")
                    }
                    _ = v58
                }
            }
            var v59 int64 = 42
            _ = v59
        } else {
            _ = len(v6)
        }
        _ = fi56
    }
    for fi60 := 0; fi60 < 14; fi60++ {
        switch 128 {
        case 650:
            if false {
                fmt.Println(v5)
                for fi61 := 0; fi61 < 18; fi61++ {
                    switch ("a" + "\n\t") {
                    case "":
                        fmt.Println(v5, v4)
                    case "foo bar":
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    default:
                        var v62 int64 = v54
                        _ = v62
                    }
                    v55 -= 1
                    switch "\n\t" {
                    case "foo bar":
                        var v63 uint = v2
                        _ = v63
                    }
                    var v64 uint8 = 2
                    _ = fi61
                    _ = v64
                }
                switch "" {
                case "a":
                    fmt.Println(v3, v1, v55)
                case "foo bar":
                    fmt.Println(v1, v55)
                case "\n\t":
                    fmt.Println(v4, v1, v3)
                }
            }
        case 127:
            if true {
                switch ("a" + "hello") {
                case "":
                    v5 = v5
                case "foo bar":
                    var v65 int16 = v4
                    _ = v65
                default:
                    if true {
                        var v66 float32 = 3.14159
                        var v67 int = fi60
                        _ = v66
                        _ = v67
                    } else {
                        v68 := GenS5{F0: 0, F1: v54, F2: 776}
                        var v69 uint8 = 42
                        _ = v68
                        _ = v69
                    }
                }
                var v70 int16 = (v4 / (((v4 + 1000)) | 1))
                _ = v70
            }
        case -2147483648:
            var v71 uint32 = (^(v55))
            _ = v71
        }
        _ = fi60
    }
    switch "foo bar" {
    case "":
        var v72 float32 = 3.14159
        _ = v72
    case "hello":
        var v73 uint = v2
        _ = v73
    case "a":
        v74 := GenS5{F0: 255, F1: v54, F2: 469}
        _ = v74
    default:
        for fi75 := 0; fi75 < 3; fi75++ {
            if false {
                fmt.Println(fi75, v2, v2)
                if (false || false) {
                    if false {
                        v5 = v5
                        var v76 int = 0
                        _ = v76
                    }
                }
                var v77 bool = false
                _ = v77
            }
            var v78 int64 = v54
            _ = fi75
            _ = v78
        }
    }
    for fi79 := 0; fi79 < 3; fi79++ {
        var v80 string = ""
        if true {
            var v81 int16 = (v5 & (v1 | (v1 * (v1 >> ((v2) & 63)))))
            var v82 float32 = 3.14159
            _ = v81
            _ = v82
        }
        v4 = v1
        _ = fi79
        _ = v80
    }
    var v83 int32 = -1
    switch "\n\t" {
    case "\n\t":
        if true {
            for fi84 := 0; fi84 < 19; fi84++ {
                for fi85 := 0; fi85 < 8; fi85++ {
                    v4 -= (v1 ^ (v5 - (v1 ^ 42)))
                    v55 -= 0
                    var v86 uint = (v2 >> (((v2 & v2)) & 63))
                    _ = fi85
                    _ = v86
                }
                _ = fi84
            }
            var v87 string = "foo bar"
            var v88 float64 = float64(1.5)
            _ = v87
            _ = v88
        } else {
            for fi89 := 0; fi89 < 3; fi89++ {
                fmt.Println(v55, v2, v1)
                if (-1 < v3) {
                    v55 = (v55 << ((v2) & 63))
                    switch "hello" {
                    case "foo bar":
                        var v90 int16 = v4
                        _ = v90
                    }
                    var v91 int16 = 32767
                    _ = v91
                }
                _ = fi89
            }
        }
    case "a":
        var v92 int = 32767
        _ = v92
    case "":
        var v93 int8 = int8(v83)
        _ = v93
    default:
        v1 = 42
    }
    var v94 int = -2147483648
    var v95 bool = (!true)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v83:", v83)
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v6
    _ = v83
    _ = v94
    _ = v95
}
