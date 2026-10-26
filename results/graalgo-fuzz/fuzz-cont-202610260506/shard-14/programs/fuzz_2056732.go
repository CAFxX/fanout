package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint32; F1 string; F2 uint64 }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS6 struct { F0 bool; F1 uint8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint8 }
func genFn0(p0 int8) int32 {
    var v7 int8 = p0
    v8 := []int{777, -136}
    for fi9 := 0; fi9 < 12; fi9++ {
        p0 += 1
        _ = fi9
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = v7
    _ = v8
    return 1
}
func genFn1() bool {
    for fi10 := 0; fi10 < 15; fi10++ {
        for fi11 := 0; fi11 < 12; fi11++ {
            if (!(!false)) {
                if false {
                    var v12 uint = 1
                    _ = v12
                } else {
                    var v13 uint = 202
                    var v14 float32 = 0.1
                    v14 = 0.0
                    var v15 uint32 = 2147483648
                    v14 = v14
                    v15 = 2147483647
                    _ = v13
                    _ = v14
                    _ = v15
                }
            }
            var v16 float64 = 0.1
            fmt.Println(fi10, v16)
            var v17 string = ""
            _ = fi11
            _ = v16
            _ = v17
        }
        _ = fi10
    }
    var v18 int64 = -9223372036854775808
    var v19 float64 = 1.5
    v19 = 1e10
    _ = v18
    _ = v19
    return true
}

func main() {
    var v1 int = 32767
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    var v4 uint = 42
    v5 := make(map[int]bool)
    v6 := make(map[int]string)
    v4 = 968
    v20 := GenS1{F0: (((("\n\t" + "") + "a") + "hello") + "hello"), F1: 9223372036854775807}
    if ((212 != 42) || (32767 > 310)) {
        if ((("\n\t" + "\n\t") + (((("\n\t" + "\n\t") + ("" + (("" + "a") + ""))) + "\n\t") + "a")) == "a") {
            v4 = 255
            v21 := []int{32768, (v1 / (((v1 - 127)) | 1))}
            for fi22 := 0; fi22 < 3; fi22++ {
                fmt.Println(v1, v4, v1)
                if false {
                    _ = len(v5)
                    var v23 int64 = 1
                    switch v1 {
                    case 814:
                        fmt.Println(v1, v23)
                    }
                    _ = v23
                } else {
                    v4 -= (^(v4))
                    delete(v3, 2)
                    if (367 <= 65535) {
                        var v24 float32 = 0.1
                        _ = v24
                    } else {
                        var v25 int32 = 2147483647
                        v26 := GenS1{F0: "", F1: -9223372036854775808}
                        _ = v25
                        _ = v26
                    }
                }
                _ = fi22
            }
            _ = v21
        }
    }
    var v27 int64 = 1
    fmt.Println(v1, v4, v1)
    fmt.Println(v27)
    fmt.Println(v4)
    v28 := []int{(v1 << ((65535) & 63)), v1}
    v29 := GenS1{F0: "", F1: v27}
    for fi30 := 0; fi30 < 12; fi30++ {
        v27 -= (v27 & v27)
        _ = fi30
    }
    fmt.Println(v4)
    v6[v1] = ("" + "")
    if true {
        if true {
            switch "\n\t" {
            case "foo bar":
                if (!("hello" > "foo bar")) {
                    for fi31 := 0; fi31 < 18; fi31++ {
                        for fi32 := 0; fi32 < 13; fi32++ {
                            v33 := v5[9223372036854775807]
                            _ = fi32
                            _ = v33
                        }
                        switch fi31 {
                        case 128:
                            switch "foo bar" {
                            case "a":
                                _ = len(v28)
                            case "hello":
                                switch ((("" + "\n\t") + "hello") + ("foo bar" + "\n\t")) {
                                case "a":
                                    for fi34 := 0; fi34 < 14; fi34++ {
                                        if (!false) {
                                            v27 = (v27 % ((-9223372036854775808) | 1))
                                            v1 = v1
                                            if (0 <= 127) {
                                                if ("foo bar" >= "a") {
                                                    var v35 string = ("" + "hello")
                                                    v36 := GenS1{F0: v35, F1: -386}
                                                    _ = v35
                                                    _ = v36
                                                } else {
                                                    fmt.Println(fi34)
                                                }
                                                v27 = (v27 * v27)
                                            }
                                            v37 := GenS1{F0: "hello", F1: -1}
                                            _ = v37
                                        }
                                        var v38 uint64 = 170
                                        var v39 int = (fi31 | -621)
                                        _ = fi34
                                        _ = v38
                                        _ = v39
                                    }
                                case "foo bar":
                                    for fi40 := 0; fi40 < 1; fi40++ {
                                        var v41 bool = false
                                        _ = fi40
                                        _ = v41
                                    }
                                }
                            case "\n\t":
                                var v42 uint16 = 32767
                                _ = v42
                            default:
                                if true {
                                    v4 -= 18446744073709551615
                                }
                            }
                        case 630:
                            for fi43 := 0; fi43 < 4; fi43++ {
                                switch 32767 {
                                case 32767:
                                    var v44 string = ""
                                    _ = v44
                                case 1:
                                    v45 := v3[557]
                                    _ = v45
                                case 4294967295:
                                    if (!(!((!false) && (!(-9223372036854775808 >= (v1 - (fi43 ^ 127))))))) {
                                        v4 = v4
                                    } else {
                                        var v46 uint = (v4 * (v4 * (v4 ^ v4)))
                                        v27 = 0
                                        _ = v46
                                    }
                                default:
                                    var v47 int8 = 127
                                    _ = v47
                                }
                                _ = fi43
                            }
                        case -17:
                            if (false || (!(3.14159 != 0.1))) {
                                v27 += 9223372036854775807
                            }
                        default:
                            v1 -= (fi31 * 9223372036854775807)
                        }
                        if true {
                            if (65535 < 664) {
                                v4 = v4
                            }
                        }
                        _ = fi31
                    }
                    v48 := GenS1{F0: "", F1: v27}
                    var v49 uint64 = uint64(v4)
                    _ = v48
                    _ = v49
                }
            case "a":
                v27 = v27
            case "\n\t":
                v50 := GenS2{F0: 4294967295, F1: "", F2: 0}
                _ = v50
            }
            var v51 float32 = ((1e10 + 3.14159) / 0.1)
            v28 = append(v28, 1)
            _ = v51
        } else {
            delete(v5, v1)
            var v52 float64 = (1e10 * 0.0)
            _ = v52
        }
    }
    switch "\n\t" {
    case "\n\t":
        if false {
            v27 += (v27 % ((v27) | 1))
            v4 = (v4 | v4)
            if (1 >= -795) {
                var v53 uint16 = 255
                for fi54 := 0; fi54 < 9; fi54++ {
                    var v55 float32 = 1.5
                    v56 := GenS3{F0: 4294967295, F1: 42, F2: (v53 >> (((v4 % (((v4 * 255)) | 1))) & 63))}
                    v55 = v55
                    _ = fi54
                    _ = v55
                    _ = v56
                }
                _ = v53
            }
        } else {
            if false {
                fmt.Println(v27, v1)
                switch ("a" + "hello") {
                case "a":
                    var v57 int8 = 42
                    _ = v57
                case "foo bar":
                    if (v4 > v4) {
                        fmt.Println(v1, v27, v4)
                        v58 := v2[1]
                        v59 := GenS3{F0: 4294967295, F1: 856, F2: 42}
                        _ = v58
                        _ = v59
                    } else {
                        v60 := GenS4{F0: 65535, F1: -1}
                        _ = v60
                    }
                case "":
                    fmt.Println(v4)
                default:
                    switch (v1 + (v1 >> ((v4) & 63))) {
                    case -876:
                        var v61 uint64 = 2
                        _ = v61
                    default:
                        v1 = -489
                    }
                }
            }
        }
    case "hello":
        if len(v28) > 0 {
            v28[0] = (v1 * v1)
        }
    case "a":
        for fi62 := 0; fi62 < 9; fi62++ {
            v4 -= v4
            v63 := GenS3{F0: 42, F1: 178, F2: 739}
            _ = fi62
            _ = v63
        }
    }
    for fi64 := 0; fi64 < 18; fi64++ {
        v65 := GenS5{F0: 1, F1: -32768, F2: 32767}
        v27 -= v27
        if (true && (32767 >= 1)) {
            switch v1 {
            case 4294967295:
                var v66 int64 = 9223372036854775807
                _ = v66
            case -9223372036854775808:
                for fi67 := 0; fi67 < 10; fi67++ {
                    switch 128 {
                    case 32768:
                        v68 := GenS6{F0: (65535 < 588), F1: 2}
                        _ = v68
                    case 4294967295:
                        v1 = (-(v1))
                    case 9223372036854775807:
                        v4 = 1
                    }
                    v1 -= (^(fi67))
                    _ = fi67
                }
            }
            v69 := GenS7{F0: 701}
            _ = v69
        } else {
            v1 += fi64
            v5[-665] = (2 != 2147483647)
            for fi70 := 0; fi70 < 17; fi70++ {
                fmt.Println(v27, fi64)
                if true {
                    if ((191 != 446) || (!(true && false))) {
                        delete(v2, 2147483647)
                        if len(v28) >= 2 {
                            v28 = v28[1:len(v28)]
                        }
                        v28 = append(v28, 127)
                        switch "hello" {
                        case "foo bar":
                            fmt.Println(fi64)
                        case "a":
                            v4 -= v4
                        case "\n\t":
                            v1 -= v1
                        }
                        if (42 <= 588) {
                            var v71 int = (-(v1))
                            _ = v71
                        } else {
                            v27 = (v27 | 42)
                            var v72 int16 = -1
                            _ = v72
                        }
                        if (!((v27 >> ((v4) & 63)) >= -9223372036854775808)) {
                            fmt.Println(v4, fi64, v1)
                            var v73 int8 = 127
                            v74 := GenS8{F0: 1}
                            v75 := GenS9{F0: 2}
                            fmt.Println(fi64)
                            _ = v73
                            _ = v74
                            _ = v75
                        }
                        var v76 uint32 = 65535
                        _ = v76
                    }
                    if (!true) {
                        var v77 int = (v1 % ((-1) | 1))
                        switch "\n\t" {
                        case "\n\t":
                            var v78 uint8 = 52
                            _ = v78
                        }
                        _ = v77
                    }
                    if true {
                        fmt.Println(fi64)
                        v4 = v4
                    }
                } else {
                    var v79 int32 = 149
                    _ = v79
                }
                fmt.Println(v27)
                _ = fi70
            }
        }
        _ = fi64
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
