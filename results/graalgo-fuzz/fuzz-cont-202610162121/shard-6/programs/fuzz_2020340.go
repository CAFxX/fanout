package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 bool; F1 int64; F2 uint32 }
type GenS4 struct { F0 int16; F1 uint64; F2 int64 }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 bool; F1 int }
type GenS7 struct { F0 string; F1 uint32 }
type GenS8 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint; F1 uint; F2 uint8 }

func main() {
    v1 := []string{(("hello" + "\n\t") + ("foo bar" + "a")), "a", ((("foo bar" + "") + "") + "hello"), "foo bar"}
    var v2 bool = (!(0 >= 972))
    var v3 int8 = -128
    v4 := make(map[uint32]string)
    if (238 <= 838) {
        switch "foo bar" {
        case "a":
            if v2 {
                fmt.Println(v3, v2)
            } else {
                var v5 uint = 255
                for fi6 := 0; fi6 < 20; fi6++ {
                    v7 := GenS1{F0: (-(fi6))}
                    _ = len(v1)
                    if (("\n\t" + "a") < ("" + "foo bar")) {
                        var v8 bool = (1e10 <= ((-2.25 - 1e10) / 1.5))
                        _ = v8
                    }
                    _ = fi6
                    _ = v7
                }
                if v2 {
                    if v2 {
                        v9 := GenS2{F0: -32768, F1: ""}
                        delete(v4, 2)
                        switch "" {
                        case "":
                            v10 := GenS1{F0: 128}
                            _ = v10
                        case "\n\t":
                            var v11 string = "foo bar"
                            _ = v11
                        case "hello":
                            if v2 {
                                var v12 float32 = 1.5
                                v3 = 127
                                var v13 int32 = 0
                                _ = v12
                                _ = v13
                            } else {
                                var v14 uint16 = 255
                                _ = v14
                            }
                        default:
                            for fi15 := 0; fi15 < 10; fi15++ {
                                var v16 int8 = v3
                                v17 := GenS1{F0: fi15}
                                _ = len(v1)
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                        }
                        _ = v9
                    }
                }
                _ = v5
            }
        }
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        for fi19 := 0; fi19 < 17; fi19++ {
            delete(v4, 2147483647)
            _ = len(v4)
            fmt.Println(v3, v3)
            _ = fi19
        }
        for fi20 := 0; fi20 < 20; fi20++ {
            if ((v3 ^ v3) <= v3) {
                v21 := GenS3{F0: v2, F1: -9223372036854775808, F2: 4294967295}
                _ = v21
            }
            _ = fi20
        }
        _ = fi18
    }
    var v22 int8 = (v3 | v3)
    fmt.Println(v3)
    v23 := GenS4{F0: 280, F1: 275, F2: -84}
    for fi24 := 0; fi24 < 7; fi24++ {
        _ = len(v4)
        var v25 float64 = 0.1
        switch (^(fi24)) {
        case -9223372036854775808:
            switch (fi24 & fi24) {
            case 999:
                for fi26 := 0; fi26 < 2; fi26++ {
                    if true {
                        v27 := GenS4{F0: 1, F1: 18446744073709551615, F2: 961}
                        v28 := GenS5{F0: 0, F1: 42}
                        _ = v27
                        _ = v28
                    } else {
                        _ = len(v1)
                    }
                    _ = fi26
                }
            case 127:
                var v29 uint16 = 42
                _ = v29
            default:
                fmt.Println(v3, v25, v3)
            }
        }
        _ = fi24
        _ = v25
    }
    switch 215 {
    case -1:
        if ((!(((float32(0.1) * ((1.5 * 0.0) / float32(0.1))) <= 1e10) && (-58 > -1))) || (2147483647 <= 128)) {
            for fi30 := 0; fi30 < 20; fi30++ {
                var v31 uint = 18446744073709551615
                if ((v31 >> ((332) & 63)) <= v31) {
                    fmt.Println(v31)
                    var v32 uint = (v31 / (((v31 << ((v31) & 63))) | 1))
                    if false {
                        if len(v1) > 0 {
                            v1[0] = ""
                        }
                        v31 -= (^(v32))
                        var v33 uint = v31
                        _ = v33
                    } else {
                        var v34 int32 = 1
                        for fi35 := 0; fi35 < 10; fi35++ {
                            var v36 int16 = 425
                            if v2 {
                                fmt.Println(v2, v22, v3)
                            } else {
                                v4[317] = (("" + "") + "foo bar")
                                if true {
                                    var v37 uint16 = 42
                                    v36 = (v36 ^ (v36 | 42))
                                    if (65535 > 0) {
                                        var v38 float64 = -2.25
                                        v3 = v22
                                        v37 = 1
                                        _ = v38
                                    }
                                    _ = v37
                                } else {
                                    var v39 int = fi30
                                    switch "a" {
                                    case "":
                                        v40 := GenS6{F0: v2, F1: fi35}
                                        _ = v40
                                    case "a":
                                        fmt.Println(v36, fi30)
                                        v3 = v22
                                    }
                                    _ = v39
                                }
                                if (v2 && v2) {
                                    switch ("\n\t" + "") {
                                    case "a":
                                        v32 = v32
                                    }
                                    v32 += (v31 | v32)
                                }
                            }
                            fmt.Println(fi35)
                            _ = fi35
                            _ = v36
                        }
                        _ = v34
                    }
                    _ = v32
                } else {
                    v41 := GenS7{F0: "a", F1: 180}
                    for fi42 := 0; fi42 < 8; fi42++ {
                        switch ((("\n\t" + "") + "a") + "foo bar") {
                        case "a":
                            v22 -= v3
                        case "":
                            v2 = v2
                        default:
                            var v43 int16 = 1
                            _ = v43
                        }
                        _ = fi42
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v41
                }
                if false {
                    v3 = 127
                    for fi44 := 0; fi44 < 20; fi44++ {
                        v3 += (v22 - v22)
                        v45 := GenS3{F0: (v2 && (!v2)), F1: -713, F2: 2147483647}
                        _ = fi44
                        _ = v45
                    }
                }
                _ = fi30
                _ = v31
            }
            var v46 bool = (!false)
            for fi47 := 0; fi47 < 20; fi47++ {
                v3 = (v3 + 0)
                if (v22 < -85) {
                    for fi48 := 0; fi48 < 1; fi48++ {
                        if (0.1 > 3.14159) {
                            if v2 {
                                var v49 float64 = 1.5
                                _ = v49
                            }
                        }
                        _ = fi48
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v50 uint32 = 4294967295
                    _ = v50
                }
                _ = len(v1)
                _ = fi47
            }
            _ = v46
        }
    }
    switch "\n\t" {
    case "foo bar":
        delete(v4, 65535)
    }
    var v51 int16 = 39
    v2 = false
    if (v3 != v22) {
        var v52 float64 = ((3.14159 + 3.14159) + 1.5)
        var v53 int = 0
        _ = v52
        _ = v53
    }
    for fi54 := 0; fi54 < 1; fi54++ {
        v55 := GenS5{F0: 42, F1: 0}
        _ = fi54
        _ = v55
    }
    for fi56 := 0; fi56 < 14; fi56++ {
        var v57 bool = ((9 <= fi56) && false)
        _ = fi56
        _ = v57
    }
    if (true || (true && (!(1 > 42)))) {
        v1 = append(v1, "a")
        fmt.Println(v51, v22)
        if v2 {
            var v58 uint = 65535
            v2 = v2
            v58 += (v58 & 65535)
            _ = v58
        } else {
            if (!v2) {
                v59 := GenS8{F0: 2, F1: 32767, F2: 86}
                v3 += (-(v3))
                if (397 == -2147483648) {
                    v51 = v51
                    v60 := GenS1{F0: 275}
                    if (!true) {
                        if false {
                            v2 = v2
                        }
                    }
                    _ = v60
                }
                _ = v59
            } else {
                v3 += v22
                v61 := GenS9{F0: 2}
                _ = v61
            }
            if (!(v2 || v2)) {
                for fi62 := 0; fi62 < 5; fi62++ {
                    v2 = (-2.25 <= 1e10)
                    if (true || ("" <= "foo bar")) {
                        var v63 int8 = v22
                        v63 -= v3
                        _ = v63
                    } else {
                        v3 = 0
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                        var v64 uint8 = 57
                        _ = v64
                    }
                    v65 := v4[1]
                    _ = fi62
                    _ = v65
                }
            } else {
                v66 := GenS10{F0: 65535, F1: 65535, F2: 187}
                fmt.Println(v2, v2, v22)
                _ = v66
            }
            v67 := GenS2{F0: -1, F1: ""}
            _ = v67
        }
    }
    v22 = (v3 - 111)
    var v68 int16 = (-(v51))
    for fi69 := 0; fi69 < 16; fi69++ {
        for fi70 := 0; fi70 < 5; fi70++ {
            if v2 {
                var v71 uint16 = 541
                _ = v71
            }
            _ = len(v1)
            v68 -= v51
            _ = fi70
        }
        var v72 bool = (!v2)
        _ = fi69
        _ = v72
    }
    v2 = ((0.0 / 1.0) >= 1e10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v51:", v51)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v51
    _ = v68
}
