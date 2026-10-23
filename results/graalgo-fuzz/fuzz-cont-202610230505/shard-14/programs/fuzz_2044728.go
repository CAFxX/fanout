package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint32; F1 string; F2 uint64 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS8 struct { F0 uint8; F1 uint16 }
func genFn0(p0 string, p1 int16) int32 {
    if ((true && (2 <= 325)) || false) {
        v5 := []int{-552, 4294967295, 803}
        _ = v5
    } else {
        if (!(32768 <= 971)) {
            var v6 bool = (true || false)
            _ = v6
        }
    }
    fmt.Println(p1)
    v7 := []int{-1}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v7
    return 1
}

func main() {
    v1 := []string{"\n\t"}
    var v2 int32 = -289
    var v3 uint = 42
    var v4 int = 128
    fmt.Println(v2)
    if (32767 == 42) {
        for fi8 := 0; fi8 < 16; fi8++ {
            if len(v1) > 0 {
                v1[0] = (("hello" + "hello") + ("hello" + "hello"))
            }
            var v9 int32 = int32(v3)
            if (9223372036854775807 != -532) {
                var v10 string = "a"
                _ = v10
            }
            _ = fi8
            _ = v9
        }
    } else {
        var v11 int16 = -966
        var v12 string = (("\n\t" + "\n\t") + ("hello" + ("a" + "hello")))
        v4 = 0
        _ = v11
        _ = v12
    }
    var v13 uint64 = 0
    var v14 int = 0
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 int64 = 13
        switch (v4 ^ 32768) {
        case -574:
            if (false || (!(127 < uint8(v13)))) {
                fmt.Println(v4)
                for fi17 := 0; fi17 < 16; fi17++ {
                    v18 := GenS1{F0: (fi17 >> ((v3) & 63)), F1: ""}
                    v2 = v2
                    _ = fi17
                    _ = v18
                }
                v19 := GenS1{F0: fi15, F1: "\n\t"}
                _ = v19
            }
        default:
            if true {
                v14 = (fi15 << (((^(v3))) & 63))
            }
        }
        if true {
            v2 -= v2
            var v20 string = "foo bar"
            _ = v20
        }
        _ = fi15
        _ = v16
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        switch ("a" + "hello") {
        case "":
            if false {
                if (1 != 127) {
                    if true {
                        v22 := GenS1{F0: 307, F1: ""}
                        v4 = 742
                        _ = v22
                    } else {
                        v2 -= v2
                        switch v14 {
                        case 915:
                            v4 += v4
                        case -9223372036854775808:
                            v2 += genFn0("hello", 42)
                        case 4294967295:
                            var v23 int = 1
                            _ = v23
                        }
                    }
                    var v24 int8 = 0
                    _ = v24
                }
                switch "" {
                case "\n\t":
                    var v25 uint16 = 65535
                    _ = v25
                case "a":
                    switch "foo bar" {
                    case "hello":
                        var v26 bool = false
                        _ = v26
                    }
                default:
                    v13 -= (v13 >> ((v3) & 63))
                }
                v27 := GenS1{F0: (v4 | fi21), F1: ("a" + "")}
                _ = v27
            }
        case "\n\t":
            switch ("foo bar" + "hello") {
            case "hello":
                if false {
                    var v28 int8 = 52
                    v28 = 1
                    v29 := GenS2{F0: 419, F1: ("foo bar" + "hello"), F2: (v13 / ((v13) | 1))}
                    _ = v28
                    _ = v29
                }
            case "foo bar":
                for fi30 := 0; fi30 < 10; fi30++ {
                    var v31 int8 = -94
                    var v32 uint64 = v13
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            case "":
                if (1e10 < 3.14159) {
                    v2 = v2
                    var v33 uint64 = (^(v13))
                    if (v14 < v4) {
                        v33 = v33
                    }
                    _ = v33
                } else {
                    v34 := GenS3{F0: 0, F1: v13}
                    fmt.Println(v14, v4)
                    _ = v34
                }
            default:
                for fi35 := 0; fi35 < 20; fi35++ {
                    if (true || false) {
                        if (int8(v2) < 55) {
                            v36 := GenS2{F0: 42, F1: ((((("" + ("foo bar" + ("" + ("\n\t" + "a")))) + "a") + (((("\n\t" + "") + "a") + ("\n\t" + "hello")) + ((("\n\t" + "hello") + "hello") + "a"))) + ("hello" + "foo bar")) + "\n\t"), F2: v13}
                            switch fi21 {
                            case 127:
                                switch ("hello" + "hello") {
                                case "\n\t":
                                    v1 = append(v1, ("foo bar" + "a"))
                                case "hello":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                default:
                                    var v37 int = (v14 ^ v14)
                                    v4 = fi35
                                    var v38 string = "\n\t"
                                    _ = v37
                                    _ = v38
                                }
                                var v39 string = ""
                                if false {
                                    v40 := GenS2{F0: 623, F1: "\n\t", F2: 0}
                                    _ = v40
                                }
                                _ = v39
                            }
                            _ = v36
                        } else {
                            v3 += 65535
                            var v41 string = ("" + "a")
                            _ = len(v1)
                            v14 -= v4
                            if false {
                                v14 = (v14 | v14)
                                fmt.Println(fi21, v4)
                            } else {
                                if (false && (true || ((-32768 > -316) && true))) {
                                    v42 := GenS1{F0: 32767, F1: v41}
                                    _ = v42
                                }
                            }
                            fmt.Println(v13, v13)
                            var v43 uint8 = 42
                            switch v4 {
                            case 9223372036854775807:
                                var v44 int16 = 406
                                _ = v44
                            }
                            _ = len(v1)
                            _ = v41
                            _ = v43
                        }
                    }
                    _ = fi35
                }
            }
        case "a":
            v4 += v14
        default:
            v2 -= v2
        }
        var v45 uint8 = 180
        _ = fi21
        _ = v45
    }
    if true {
        switch (v14 / ((-9223372036854775808) | 1)) {
        case 42:
            v3 = (v3 | v3)
        case 9223372036854775807:
            switch v4 {
            case -710:
                v1 = append(v1, "foo bar")
            case -117:
                var v46 uint32 = 1
                _ = v46
            case 42:
                if (678 != (v13 + 2)) {
                    var v47 int32 = 42
                    v13 -= (v13 ^ (v13 & v13))
                    var v48 string = "\n\t"
                    _ = v47
                    _ = v48
                } else {
                    if false {
                        for fi49 := 0; fi49 < 19; fi49++ {
                            for fi50 := 0; fi50 < 3; fi50++ {
                                fmt.Println(v14, fi50, v14)
                                v4 -= (v14 & fi50)
                                var v51 int32 = genFn0((("" + "") + "hello"), -1)
                                _ = fi50
                                _ = v51
                            }
                            _ = fi49
                        }
                    } else {
                        _ = len(v1)
                    }
                    v2 = 1
                    var v52 string = "hello"
                    _ = v52
                }
            }
        }
        switch (v4 / ((375) | 1)) {
        case 42:
            if true {
                v3 += v3
                switch v14 {
                case -1:
                    var v53 string = ("foo bar" + "\n\t")
                    _ = v53
                case 611:
                    var v54 bool = false
                    _ = v54
                }
            } else {
                v3 -= 39
                if len(v1) > 0 {
                    v1[0] = ("\n\t" + "foo bar")
                }
                fmt.Println(v2)
            }
        default:
            switch 264 {
            case -9223372036854775808:
                v55 := GenS4{F0: 1}
                _ = v55
            case 32767:
                v1 = append(v1, "")
            case -204:
                var v56 uint = (v3 | v3)
                _ = v56
            }
        }
        if (4294967295 >= (^(v13))) {
            fmt.Println(v13)
            v57 := GenS2{F0: 434, F1: ((("hello" + "\n\t") + "hello") + ("a" + (("" + ("\n\t" + "foo bar")) + "\n\t"))), F2: v13}
            _ = v57
        }
    }
    switch "hello" {
    case "hello":
        v14 = 1
    }
    v2 -= v2
    for fi58 := 0; fi58 < 14; fi58++ {
        for fi59 := 0; fi59 < 11; fi59++ {
            switch ("\n\t" + "\n\t") {
            case "":
                if (25 < 42) {
                    var v60 uint16 = 2
                    _ = len(v1)
                    v61 := GenS4{F0: uint8(v13)}
                    switch (-(fi59)) {
                    case -1:
                        var v62 uint16 = (v60 % ((v60) | 1))
                        _ = v62
                    case 48:
                        if ((3.14159 / 0.1) != 0.1) {
                            v3 += v3
                            v3 = (v3 << ((uint(v60)) & 63))
                            if false {
                                switch -954 {
                                case -299:
                                    var v63 int32 = v2
                                    _ = v63
                                case -1:
                                    v3 = v3
                                case 32768:
                                    switch (fi59 - -474) {
                                    case -353:
                                        if false {
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v64 int64 = 0
                                                _ = v64
                                            default:
                                                fmt.Println(v2, v4)
                                            }
                                            fmt.Println(v13, fi59)
                                            if false {
                                                var v65 string = "hello"
                                                _ = v65
                                            } else {
                                                if false {
                                                    var v66 float32 = 0.1
                                                    v67 := GenS5{F0: v14, F1: (true && (v13 <= (v13 + v13)))}
                                                    v68 := GenS2{F0: 2147483647, F1: "foo bar", F2: v13}
                                                    _ = v66
                                                    _ = v67
                                                    _ = v68
                                                } else {
                                                    v69 := GenS2{F0: 2, F1: (("\n\t" + "a") + (("a" + "\n\t") + ("a" + ""))), F2: v13}
                                                    v70 := GenS1{F0: v4, F1: "hello"}
                                                    _ = v69
                                                    _ = v70
                                                }
                                                var v71 bool = false
                                                _ = v71
                                            }
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = "\n\t"
                                            }
                                            switch v4 {
                                            case -2147483648:
                                                switch (v4 - fi58) {
                                                case 127:
                                                    v13 -= (v13 * 1)
                                                }
                                            case 9223372036854775807:
                                                v72 := GenS2{F0: 2147483648, F1: (("foo bar" + "") + "foo bar"), F2: (v13 >> ((v3) & 63))}
                                                _ = v72
                                            }
                                        }
                                    default:
                                        if true {
                                            v13 += v13
                                            switch fi58 {
                                            case -342:
                                                fmt.Println(v13, v60, v13)
                                            case 32768:
                                                if len(v1) > 0 {
                                                    v1[0] = "a"
                                                }
                                            }
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v14 = 42
                                        }
                                    }
                                default:
                                    switch "foo bar" {
                                    case "\n\t":
                                        if len(v1) > 0 {
                                            v1[0] = "a"
                                        }
                                    case "a":
                                        fmt.Println(fi58, v3, v3)
                                    case "hello":
                                        fmt.Println(v60, v60, v13)
                                    }
                                    var v73 int64 = 812
                                    if len(v1) > 0 {
                                        v1[0] = "\n\t"
                                    }
                                    v1 = append(v1, "")
                                    _ = v73
                                }
                                var v74 int8 = 42
                                _ = v74
                            }
                            var v75 int32 = v2
                            v1 = append(v1, (("a" + "foo bar") + "foo bar"))
                            v76 := GenS6{F0: -9223372036854775808}
                            _ = v75
                            _ = v76
                        }
                    case 42:
                        v3 -= v3
                        var v77 uint8 = 127
                        _ = v77
                    }
                    _ = v60
                    _ = v61
                }
            case "foo bar":
                var v78 float64 = (-2.25 / 0.1)
                _ = v78
            case "hello":
                v79 := GenS7{F0: 1, F1: 100, F2: 65535}
                v80 := GenS7{F0: (v13 / ((1) | 1)), F1: 127, F2: 127}
                _ = v79
                _ = v80
            default:
                if ((((0.1 * ((1.5 * -2.25) * 0.0)) * 0.1) + 0.0) == 0.0) {
                    if (0 < 850) {
                        var v81 int32 = -2147483648
                        _ = v81
                    } else {
                        switch "a" {
                        case "\n\t":
                            if ((v13 + (v13 + v13)) < (v13 / ((v13) | 1))) {
                                v4 += (v14 << ((65535) & 63))
                                var v82 int16 = 32767
                                v3 = 255
                                v3 = 451
                                if (!true) {
                                    v1 = append(v1, "a")
                                    switch int(v3) {
                                    case 42:
                                        var v83 int16 = v82
                                        _ = v83
                                    case 9223372036854775807:
                                        if len(v1) > 0 {
                                            v1[0] = "a"
                                        }
                                    case 0:
                                        fmt.Println(v4, v13)
                                    }
                                } else {
                                    v1 = append(v1, "foo bar")
                                    v84 := GenS8{F0: 42, F1: 32767}
                                    _ = v84
                                }
                                _ = v82
                            }
                        case "":
                            _ = len(v1)
                        case "a":
                            v13 = v13
                        }
                        switch (fi58 << ((255) & 63)) {
                        case -545:
                            v2 += v2
                        case 1:
                            if ((false || (!(32767 != 1))) && (((v3 - v3) == uint(v13)) || (!false))) {
                                v4 = 32768
                            } else {
                                v85 := GenS2{F0: 2147483648, F1: "hello", F2: v13}
                                v2 -= -984
                                v86 := GenS1{F0: 215, F1: ""}
                                _ = v85
                                _ = v86
                            }
                        }
                    }
                    v13 -= v13
                    var v87 uint16 = 536
                    v88 := GenS7{F0: 1, F1: 0, F2: 65535}
                    _ = v87
                    _ = v88
                }
            }
            var v89 int16 = 123
            _ = fi59
            _ = v89
        }
        _ = fi58
    }
    var v90 string = ("a" + "a")
    v14 = (v14 >> (((v3 | 42)) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v90
}
