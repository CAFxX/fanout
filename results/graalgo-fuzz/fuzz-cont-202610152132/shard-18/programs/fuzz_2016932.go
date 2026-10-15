package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64; F1 string; F2 bool }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int16; F1 uint16; F2 uint }
type GenS8 struct { F0 string; F1 int8 }
func genFn0() int32 {
    var v4 float64 = 3.14159
    if true {
        for fi5 := 0; fi5 < 17; fi5++ {
            var v6 string = "hello"
            var v7 bool = false
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 float64 = v4
                _ = fi8
                _ = v9
            }
            _ = fi5
            _ = v6
            _ = v7
        }
    } else {
        var v10 string = ("" + "hello")
        _ = v10
    }
    if false {
        v11 := []int{}
        _ = v11
    }
    _ = v4
    return 0
}
func genFn1() string {
    v12 := []int{-878}
    var v13 string = "hello"
    var v14 float64 = (1e10 - (1e10 * 3.14159))
    if (!true) {
        if (4294967295 < 1) {
            if true {
                v15 := GenS1{F0: true, F1: 42, F2: (v13 + (("foo bar" + (v13 + (v13 + (v13 + v13)))) + v13))}
                for fi16 := 0; fi16 < 19; fi16++ {
                    switch (fi16 + (fi16 * (fi16 << ((0) & 63)))) {
                    case 32768:
                        v17 := GenS1{F0: false, F1: 128, F2: (("\n\t" + v13) + v13)}
                        _ = v17
                    case 1:
                        switch 24 {
                        case -139:
                            fmt.Println(v14)
                        case 234:
                            v13 = (v13 + (((v13 + v13) + v13) + "foo bar"))
                        }
                    }
                    if (-340 > -1) {
                        v14 += float64((1e10 / -2.25))
                        fmt.Println(v13, fi16)
                        if false {
                            fmt.Println(v14)
                        }
                    } else {
                        for fi18 := 0; fi18 < 16; fi18++ {
                            var v19 string = ""
                            _ = fi18
                            _ = v19
                        }
                        if false {
                            v14 -= (0.0 * ((v14 + v14) - v14))
                            var v20 uint64 = 257
                            _ = v20
                        } else {
                            if false {
                                v21 := GenS2{F0: 0}
                                _ = v21
                            }
                            if (-2.25 != (1.5 - 1.5)) {
                                for fi22 := 0; fi22 < 15; fi22++ {
                                    fmt.Println(fi16)
                                    _ = fi22
                                }
                            }
                        }
                    }
                    if ((!(127 == 65535)) || true) {
                        var v23 int64 = -1
                        for fi24 := 0; fi24 < 6; fi24++ {
                            v25 := GenS2{F0: 320}
                            _ = fi24
                            _ = v25
                        }
                        _ = v23
                    }
                    _ = fi16
                }
                _ = v15
            }
            fmt.Println(v14, v13)
            if ((v14 < float64(0.0)) && false) {
                switch 42 {
                case 42:
                    for fi26 := 0; fi26 < 17; fi26++ {
                        if len(v12) > 0 {
                            v12[0] = (fi26 & (fi26 - 1))
                        }
                        _ = fi26
                    }
                default:
                    if (!true) {
                        v14 = float64((1e10 - 0.0))
                        var v27 uint32 = 2
                        _ = v27
                    }
                }
            }
        } else {
            var v28 uint32 = 4294967295
            v29 := GenS3{F0: genFn0(), F1: 4294967295}
            for fi30 := 0; fi30 < 9; fi30++ {
                switch ((v13 + ((((("" + v13) + (v13 + v13)) + v13) + "hello") + "hello")) + "a") {
                case "\n\t":
                    v12 = append(v12, 42)
                }
                _ = fi30
            }
            _ = v28
            _ = v29
        }
        v12 = append(v12, 32767)
    }
    _ = v12
    _ = v13
    _ = v14
    return (v13 + (v13 + ""))
}

func main() {
    var v1 float32 = float32((3.14159 / float64(1e10)))
    var v2 int64 = 19
    v3 := make(map[int]string)
    _ = len(v3)
    var v31 uint64 = 321
    for fi32 := 0; fi32 < 13; fi32++ {
        if false {
            var v33 uint64 = (v31 % (((v31 & (v31 >> ((65535) & 63)))) | 1))
            for fi34 := 0; fi34 < 4; fi34++ {
                if false {
                    var v35 int32 = 0
                    fmt.Println(fi32, fi32, fi34)
                    if (!(127 > (fi32 - (fi32 ^ (fi34 << ((2) & 63)))))) {
                        v36 := []int{113, fi34, fi32}
                        v37 := GenS4{F0: v2}
                        _ = v36
                        _ = v37
                    }
                    _ = v35
                }
                v33 += (v31 % ((v33) | 1))
                v38 := []int{}
                _ = fi34
                _ = v38
            }
            v1 = v1
            _ = v33
        } else {
            var v39 bool = true
            if true {
                switch (fi32 + fi32) {
                case 4294967295:
                    var v40 int32 = -548
                    _ = v40
                case -2147483648:
                    v41 := GenS1{F0: v39, F1: 42, F2: "a"}
                    _ = v41
                case 1:
                    switch "a" {
                    case "hello":
                        fmt.Println(fi32, v1, v31)
                    }
                }
                if v39 {
                    v42 := GenS2{F0: v31}
                    switch -499 {
                    case -9223372036854775808:
                        if ("hello" < ("hello" + "\n\t")) {
                            for fi43 := 0; fi43 < 19; fi43++ {
                                if v39 {
                                    fmt.Println(v2)
                                    var v44 bool = (v39 && v39)
                                    v44 = (true || v39)
                                    v45 := GenS1{F0: v44, F1: 1, F2: "hello"}
                                    _ = v44
                                    _ = v45
                                } else {
                                    v46 := GenS1{F0: ((((((0.0 / -2.25) / 1e10) * -2.25) - 1.5) + 1.5) > float64(v1)), F1: 128, F2: "hello"}
                                    _ = len(v3)
                                    fmt.Println(v1, fi32, fi43)
                                    _ = v46
                                }
                                _ = fi43
                            }
                            if (v1 == (float32(0.1) - v1)) {
                                v2 = -9223372036854775808
                            }
                            if v39 {
                                for fi47 := 0; fi47 < 6; fi47++ {
                                    switch (-(fi47)) {
                                    case 0:
                                        var v48 int16 = -32768
                                        _ = v48
                                    default:
                                        var v49 float32 = v1
                                        _ = v49
                                    }
                                    v50 := []int{}
                                    _ = fi47
                                    _ = v50
                                }
                                var v51 int8 = 127
                                if (v1 < (3.14159 * float32(0.1))) {
                                    if (v2 == 1) {
                                        v52 := []int{}
                                        if len(v52) > 0 {
                                            v52[0] = 42
                                        }
                                        for fi53 := 0; fi53 < 8; fi53++ {
                                            v54 := GenS2{F0: (v31 >> ((65535) & 63))}
                                            if (v39 || false) {
                                                var v55 bool = (35 == 250)
                                                if len(v52) > 0 {
                                                    v52[0] = fi53
                                                }
                                                var v56 uint8 = 2
                                                _ = v55
                                                _ = v56
                                            }
                                            _ = fi53
                                            _ = v54
                                        }
                                        _ = v52
                                    }
                                    v57 := GenS5{F0: v31, F1: ("a" + ""), F2: v39}
                                    _ = v57
                                }
                                _ = v51
                            } else {
                                var v58 uint64 = (v31 >> ((255) & 63))
                                _ = v58
                            }
                        } else {
                            if v39 {
                                switch "hello" {
                                case "hello":
                                    for fi59 := 0; fi59 < 6; fi59++ {
                                        delete(v3, int(v31))
                                        if ((-2.25 * 0.1) >= 0.1) {
                                            v31 = 224
                                            fmt.Println(fi59, v2, v2)
                                            switch "hello" {
                                            case "\n\t":
                                                var v60 int8 = -52
                                                fmt.Println(v1, v31, v31)
                                                _ = v60
                                            default:
                                                var v61 uint64 = (v31 / ((v31) | 1))
                                                _ = v61
                                            }
                                            v39 = (v39 && v39)
                                            if v39 {
                                                v31 -= 1
                                                v31 = (v31 ^ (v31 + v31))
                                            } else {
                                                v1 = ((v1 / v1) * 0.0)
                                                v39 = true
                                            }
                                            var v62 uint32 = 0
                                            v62 -= (v62 - 823)
                                            var v63 uint32 = v62
                                            _ = v62
                                            _ = v63
                                        }
                                        _ = fi59
                                    }
                                default:
                                    for fi64 := 0; fi64 < 20; fi64++ {
                                        var v65 uint16 = 130
                                        v66 := GenS6{F0: 0}
                                        _ = fi64
                                        _ = v65
                                        _ = v66
                                    }
                                }
                            }
                        }
                    case 32768:
                        switch fi32 {
                        case 4294967295:
                            v67 := GenS2{F0: (v31 & (v31 * 42))}
                            _ = v67
                        case -934:
                            var v68 int8 = 19
                            _ = v68
                        }
                    case -484:
                        for fi69 := 0; fi69 < 15; fi69++ {
                            fmt.Println(v2)
                            _ = fi69
                        }
                    }
                    _ = v42
                }
            }
            _ = v39
        }
        var v70 int16 = -499
        switch ((("hello" + "foo bar") + "\n\t") + ("foo bar" + ("" + ""))) {
        case "\n\t":
            fmt.Println(v70)
        case "":
            for fi71 := 0; fi71 < 6; fi71++ {
                var v72 int8 = 0
                var v73 uint = 4294967295
                _ = fi71
                _ = v72
                _ = v73
            }
        default:
            var v74 uint16 = 255
            _ = v74
        }
        _ = fi32
        _ = v70
    }
    if false {
        for fi75 := 0; fi75 < 10; fi75++ {
            v2 = (v2 & 9223372036854775807)
            var v76 bool = (!false)
            fmt.Println(fi75, v1, v31)
            _ = fi75
            _ = v76
        }
        v77 := []int{-9223372036854775808, 0}
        _ = v77
    } else {
        v78 := GenS6{F0: 1}
        _ = v78
    }
    switch 0 {
    case 258:
        var v79 int16 = -1
        _ = v79
    case -738:
        for fi80 := 0; fi80 < 6; fi80++ {
            if (int64(fi80) >= v2) {
                var v81 float32 = 1e10
                _ = v81
            }
            v82 := GenS7{F0: 0, F1: 277, F2: 65535}
            fmt.Println(v2)
            _ = fi80
            _ = v82
        }
    case -946:
        v83 := GenS8{F0: "\n\t", F1: -115}
        _ = v83
    default:
        var v84 string = "\n\t"
        _ = v84
    }
    v1 += v1
    fmt.Println(v31)
    v31 += 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
}
