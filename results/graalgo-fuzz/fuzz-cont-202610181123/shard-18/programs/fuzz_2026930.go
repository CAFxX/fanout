package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 int64; F2 uint64 }
type GenS4 struct { F0 int64; F1 int32; F2 int8 }
type GenS5 struct { F0 uint8; F1 int8; F2 bool }
type GenS6 struct { F0 string; F1 uint; F2 uint8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint32; F1 int64 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int16; F1 string; F2 int }
type GenS12 struct { F0 int; F1 uint32; F2 bool }
func genFn0(p0 uint32, p1 uint, p2 int8) uint16 {
    var v5 uint = p1
    var v6 int16 = 1000
    v7 := GenS1{F0: -41}
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: (^(p2))}
        switch ("" + "a") {
        case "a":
            if false {
                v10 := GenS2{F0: 2147483647}
                _ = v10
            }
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{"\n\t", "\n\t"}
    v3 := []uint32{}
    var v4 uint32 = 2
    switch ("\n\t" + "hello") {
    case "\n\t":
        switch 32768 {
        case 646:
            v4 -= v4
        case 4294967295:
            fmt.Println(v4)
        case 32767:
            _ = len(v3)
        }
    case "foo bar":
        var v11 int32 = 0
        _ = v11
    }
    var v12 int16 = -761
    switch "hello" {
    case "":
        fmt.Println(v12)
    }
    if ((-(v12)) <= v12) {
        fmt.Println(v4)
        if ((!(!(!(!(-1 <= 42))))) || (4294967295 >= 188)) {
            v13 := GenS3{F0: -1, F1: -9223372036854775808, F2: 1}
            v1["\n\t"] = -140
            for fi14 := 0; fi14 < 2; fi14++ {
                if (477 <= 182) {
                    fmt.Println(v4, v12)
                } else {
                    for fi15 := 0; fi15 < 13; fi15++ {
                        if len(v2) > 0 {
                            v2[0] = "\n\t"
                        }
                        delete(v1, "")
                        v16 := GenS4{F0: 1, F1: 42, F2: 1}
                        _ = fi15
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 17; fi17++ {
                        _ = len(v1)
                        v4 += v4
                        switch (("foo bar" + "hello") + "a") {
                        case "":
                            if true {
                                switch 417 {
                                case -580:
                                    var v18 int16 = 1000
                                    _ = v18
                                case -700:
                                    switch "" {
                                    case "a":
                                        switch fi17 {
                                        case -2147483648:
                                            var v19 int64 = 42
                                            fmt.Println(fi17)
                                            _ = v19
                                        case 42:
                                            var v20 uint32 = v4
                                            _ = v20
                                        case 688:
                                            switch "hello" {
                                            case "":
                                                if (!true) {
                                                    fmt.Println(v4)
                                                    v21 := GenS4{F0: 70, F1: -520, F2: 97}
                                                    _ = v21
                                                }
                                            }
                                            fmt.Println(v12)
                                        default:
                                            v12 += v12
                                        }
                                    default:
                                        v4 += (v4 | v4)
                                    }
                                case 0:
                                    if (113 > 185) {
                                        var v22 int8 = 100
                                        _ = v22
                                    } else {
                                        if false {
                                            switch 32767 {
                                            case 32767:
                                                var v23 string = (("foo bar" + "\n\t") + "foo bar")
                                                _ = v23
                                            default:
                                                var v24 int64 = 0
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = v24
                                            }
                                            if false {
                                                v25 := GenS5{F0: 255, F1: -1, F2: false}
                                                var v26 float32 = -2.25
                                                v2 = append(v2, "a")
                                                _ = v25
                                                _ = v26
                                            } else {
                                                v4 = 2147483648
                                            }
                                        } else {
                                            v27 := GenS3{F0: -118, F1: -9223372036854775808, F2: 104}
                                            _ = v27
                                        }
                                        _ = len(v3)
                                        var v28 uint32 = 4294967295
                                        var v29 int64 = 1
                                        v30 := GenS2{F0: -974}
                                        v28 = v4
                                        _ = v28
                                        _ = v29
                                        _ = v30
                                    }
                                default:
                                    fmt.Println(fi17, fi17, fi17)
                                }
                                if (false || false) {
                                    if false {
                                        v31 := GenS3{F0: 1, F1: -442, F2: 2}
                                        var v32 int32 = 1
                                        v4 = v4
                                        _ = v31
                                        _ = v32
                                    }
                                    v33 := GenS4{F0: -1, F1: -1, F2: -29}
                                    var v34 string = "\n\t"
                                    _ = v33
                                    _ = v34
                                }
                                switch (((((("" + ("a" + "a")) + "a") + ("hello" + ((("\n\t" + ("a" + "hello")) + "a") + "foo bar"))) + "hello") + "hello") + "a") {
                                case "foo bar":
                                    v35 := GenS6{F0: ("a" + ("a" + ("foo bar" + ""))), F1: 983, F2: 22}
                                    _ = v35
                                case "\n\t":
                                    _ = len(v1)
                                    v12 -= (-(v12))
                                case "":
                                    switch "\n\t" {
                                    case "\n\t":
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v36 uint64 = 4294967295
                                            switch (fi17 ^ fi17) {
                                            case 0:
                                                var v37 uint8 = 2
                                                _ = v37
                                            case 32767:
                                                v12 = 521
                                            case 623:
                                                v12 = v12
                                            default:
                                                v1[("hello" + "a")] = fi14
                                            }
                                            _ = v36
                                        case "hello":
                                            fmt.Println(fi17, fi17, fi17)
                                        case "foo bar":
                                            fmt.Println(fi14)
                                        default:
                                            fmt.Println(fi14)
                                        }
                                    case "":
                                        var v38 int64 = 220
                                        _ = v38
                                    case "a":
                                        if true {
                                            v39 := GenS7{F0: 128}
                                            v1["hello"] = fi17
                                            var v40 uint = 18446744073709551615
                                            if (567 >= (fi17 + (fi17 >> ((527) & 63)))) {
                                                _ = len(v1)
                                                switch "hello" {
                                                case "foo bar":
                                                    v4 += 2
                                                case "hello":
                                                    v1[(("" + "") + ("hello" + "hello"))] = fi14
                                                case "a":
                                                    v40 = 1
                                                }
                                                v40 = v40
                                            }
                                            fmt.Println(v40)
                                            _ = v39
                                            _ = v40
                                        } else {
                                            if false {
                                                if (0 > 1) {
                                                    if (!(((fi17 ^ (fi14 >> ((273) & 63))) == fi14) && false)) {
                                                        v1["\n\t"] = (-(fi14))
                                                        fmt.Println(v12)
                                                        v41 := GenS3{F0: 42, F1: 1, F2: 529}
                                                        _ = v41
                                                    }
                                                    switch fi14 {
                                                    case -9223372036854775808:
                                                        v42 := GenS7{F0: 3}
                                                        _ = v42
                                                    default:
                                                        if (!(fi14 >= fi14)) {
                                                            v43 := GenS8{F0: false}
                                                            v4 += v4
                                                            _ = v43
                                                        } else {
                                                            switch (fi14 / (((fi14 ^ fi17)) | 1)) {
                                                            case 716:
                                                                if (uint8(v4) >= 127) {
                                                                    var v44 int = fi17
                                                                    _ = v44
                                                                } else {
                                                                    var v45 string = "hello"
                                                                    v12 = v12
                                                                    _ = v45
                                                                }
                                                            case 32767:
                                                                var v46 int = fi14
                                                                _ = v46
                                                            default:
                                                                v47 := GenS9{F0: (v4 & (v4 * v4)), F1: -9223372036854775808}
                                                                _ = v47
                                                            }
                                                            v12 = v12
                                                            _ = len(v2)
                                                            var v48 int32 = 2147483647
                                                            _ = v48
                                                        }
                                                    }
                                                } else {
                                                    v49 := GenS2{F0: -1}
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        if (0 < 100) {
                                                            var v50 int = 128
                                                            v51 := GenS10{F0: 0}
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                    default:
                                                        fmt.Println(fi14)
                                                    }
                                                    _ = v49
                                                }
                                            }
                                            fmt.Println(fi17, v4, v4)
                                        }
                                    default:
                                        v12 = (v12 ^ 911)
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                default:
                                    fmt.Println(v12, v12, fi14)
                                }
                            }
                        }
                        fmt.Println(v4)
                        v52 := v1["a"]
                        if true {
                            switch fi14 {
                            case 603:
                                if true {
                                    v52 += fi17
                                    v4 = v4
                                    v12 = -32768
                                    v4 = v4
                                }
                            case 815:
                                v4 = (^(v4))
                            case -316:
                                v12 += (v12 + v12)
                            default:
                                switch (^(fi17)) {
                                case -1:
                                    if (!((float32((float64((1e10 / ((((0.0 - -2.25) / (3.14159 + (0.1 - (3.14159 * (0.1 + 0.1))))) - -2.25) / (-2.25 + 0.1)))) / ((1.5 - -2.25) + 3.14159))) / (1.5 + 0.0)) >= 1.5)) {
                                        var v53 uint32 = 660
                                        v2 = append(v2, "")
                                        _ = v53
                                    } else {
                                        v54 := GenS5{F0: 98, F1: 127, F2: true}
                                        v4 = (v4 & v4)
                                        v55 := GenS11{F0: v12, F1: "hello", F2: (-(fi14))}
                                        _ = v54
                                        _ = v55
                                    }
                                }
                            }
                        } else {
                            v12 = (-(v12))
                        }
                        v56 := v1["hello"]
                        if ((32767 != 786) && true) {
                            switch (("hello" + ((("a" + ("\n\t" + ("foo bar" + "\n\t"))) + "\n\t") + ((("hello" + "a") + "foo bar") + "\n\t"))) + "a") {
                            case "a":
                                var v57 int = -684
                                _ = v57
                            case "":
                                v1["foo bar"] = (fi14 | fi17)
                            case "hello":
                                v12 -= -501
                            default:
                                _ = len(v1)
                            }
                        } else {
                            var v58 uint = 65535
                            _ = v58
                        }
                        var v59 int16 = v12
                        v60 := GenS1{F0: 127}
                        _ = fi17
                        _ = v52
                        _ = v56
                        _ = v59
                        _ = v60
                    }
                    var v61 uint16 = 1
                    _ = v61
                }
                var v62 int16 = (v12 >> ((829) & 63))
                _ = fi14
                _ = v62
            }
            _ = v13
        }
        v12 = v12
    } else {
        var v63 string = "foo bar"
        var v64 string = (("hello" + v63) + "a")
        if (((1e10 * 0.0) / 1e10) <= (1.5 - 0.0)) {
            _ = len(v1)
        }
        _ = v63
        _ = v64
    }
    var v65 string = "hello"
    if true {
        var v66 uint = 1
        for fi67 := 0; fi67 < 20; fi67++ {
            v65 = "\n\t"
            _ = fi67
        }
        _ = v66
    }
    delete(v1, "foo bar")
    v65 = ("\n\t" + v65)
    if false {
        var v68 uint16 = genFn0(803, 4294967295, 127)
        _ = v68
    } else {
        v69 := GenS12{F0: 187, F1: (v4 / (((v4 >> ((414) & 63))) | 1)), F2: (!false)}
        var v70 string = v65
        _ = v69
        _ = v70
    }
    var v71 int32 = 42
    switch (v65 + (v65 + "foo bar")) {
    case "hello":
        for fi72 := 0; fi72 < 12; fi72++ {
            var v73 int32 = (v71 * (v71 ^ v71))
            _ = fi72
            _ = v73
        }
    case "a":
        for fi74 := 0; fi74 < 14; fi74++ {
            switch v65 {
            case "\n\t":
                for fi75 := 0; fi75 < 3; fi75++ {
                    var v76 int = (fi75 >> ((1) & 63))
                    v12 -= 183
                    switch "hello" {
                    case "foo bar":
                        v77 := GenS5{F0: 2, F1: 114, F2: (-128 == -107)}
                        _ = v77
                    }
                    v78 := GenS5{F0: 42, F1: 42, F2: (v71 < (v71 % ((v71) | 1)))}
                    _ = fi75
                    _ = v76
                    _ = v78
                }
            case "a":
                switch (fi74 * fi74) {
                case -9223372036854775808:
                    for fi79 := 0; fi79 < 20; fi79++ {
                        v65 = v65
                        _ = fi79
                    }
                case 32767:
                    v65 = v65
                default:
                    for fi80 := 0; fi80 < 9; fi80++ {
                        switch (^(fi80)) {
                        case 1:
                            delete(v1, (("hello" + ("a" + v65)) + v65))
                        default:
                            v65 = ((v65 + "foo bar") + (v65 + v65))
                        }
                        _ = fi80
                    }
                }
            case "hello":
                _ = len(v1)
            default:
                var v81 uint64 = 110
                _ = v81
            }
            _ = fi74
        }
    default:
        var v82 uint64 = 18446744073709551615
        _ = v82
    }
    _ = len(v3)
    v71 -= v71
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v65:", v65)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v65
    _ = v71
}
