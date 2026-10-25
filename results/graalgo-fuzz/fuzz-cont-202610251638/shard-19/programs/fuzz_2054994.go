package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 uint64; F2 int }
type GenS4 struct { F0 uint8; F1 uint32; F2 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint; F1 uint; F2 int16 }
type GenS7 struct { F0 uint; F1 int64; F2 int8 }
type GenS8 struct { F0 uint8; F1 string; F2 int }
type GenS9 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS10 struct { F0 int32; F1 uint8 }
type GenS11 struct { F0 uint64; F1 uint16 }

func main() {
    var v1 uint16 = 32767
    var v2 float32 = (3.14159 - 1.5)
    v3 := []string{"foo bar", "foo bar", "", "\n\t"}
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]string)
    v6 := make(map[int]bool)
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: fi7}
        _ = fi7
        _ = v8
    }
    var v9 int32 = 862
    var v10 uint16 = v1
    if len(v3) > 0 {
        v3[0] = ("hello" + "a")
    }
    v11 := GenS2{F0: ((260 < 42) || false)}
    v10 -= (v10 << ((0) & 63))
    if true {
        for fi12 := 0; fi12 < 3; fi12++ {
            if (!(255 > 255)) {
                v13 := v5[42]
                _ = v13
            }
            v14 := GenS2{F0: (548 != 42)}
            _ = fi12
            _ = v14
        }
    } else {
        v15 := GenS1{F0: 416}
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 uint8 = 127
            v1 += (v1 / (((v1 - (v10 % ((348) | 1)))) | 1))
            if (!(531 < 42)) {
                var v18 uint32 = 65535
                for fi19 := 0; fi19 < 13; fi19++ {
                    if ((float32(((-2.25 / (float64(1.5) / 1.5)) / ((float64(v2) / (0.1 / 1e10)) - -2.25))) != v2) || true) {
                        var v20 uint = 2
                        _ = v20
                    } else {
                        var v21 int = 32767
                        if ((-2.25 * 1.5) != -2.25) {
                            if (true && false) {
                                var v22 int32 = v9
                                _ = v22
                            } else {
                                v2 -= (v2 + 0.1)
                                v23 := v6[fi19]
                                var v24 uint = 18446744073709551615
                                delete(v5, 2147483647)
                                if len(v3) > 0 {
                                    v3[0] = "foo bar"
                                }
                                v25 := v5[v18]
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                        } else {
                            if false {
                                var v26 float32 = float32(3.14159)
                                _ = v26
                            }
                            v18 -= (v18 % ((uint32(v10)) | 1))
                            v27 := GenS1{F0: v21}
                            _ = v27
                        }
                        if ((fi16 * (fi16 & (fi19 ^ v21))) > v21) {
                            switch (fi19 - (v21 / ((4294967295) | 1))) {
                            case 730:
                                switch (^(v21)) {
                                case 887:
                                    if true {
                                        var v28 uint64 = 2
                                        var v29 int8 = -1
                                        if ((((float64((1e10 - v2)) * 1e10) * (3.14159 * -2.25)) - 1e10) < 0.1) {
                                            if true {
                                                var v30 float64 = (1e10 * 3.14159)
                                                var v31 float64 = (((((v30 + 1.5) * ((3.14159 / 1.5) * v30)) * v30) * 0.0) / float64(v2))
                                                _ = v30
                                                _ = v31
                                            } else {
                                                delete(v5, 2147483648)
                                                if true {
                                                    switch "\n\t" {
                                                    case "":
                                                        switch (^(fi16)) {
                                                        case -1:
                                                            if false {
                                                                v10 += v1
                                                                if false {
                                                                    v29 = (v29 * -1)
                                                                    var v32 int8 = -128
                                                                    _ = v32
                                                                } else {
                                                                    delete(v5, (v18 * 152))
                                                                }
                                                                v33 := GenS2{F0: (false && true)}
                                                                v21 = (fi19 + fi19)
                                                                _ = v33
                                                            } else {
                                                                var v34 uint16 = (v1 >> ((65535) & 63))
                                                                _ = v34
                                                            }
                                                        }
                                                    }
                                                    var v35 bool = false
                                                    _ = v35
                                                } else {
                                                    switch (-(v21)) {
                                                    case 0:
                                                        if false {
                                                            if (float32(0.0) < v2) {
                                                                var v36 int8 = (^(v29))
                                                                if false {
                                                                    fmt.Println(v29, v17, v29)
                                                                } else {
                                                                    if (127 < 0) {
                                                                        fmt.Println(v17, v36)
                                                                    } else {
                                                                        v2 += 1.5
                                                                        var v37 int32 = 42
                                                                        _ = v37
                                                                    }
                                                                }
                                                                var v38 uint32 = v18
                                                                var v39 int = fi19
                                                                _ = v36
                                                                _ = v38
                                                                _ = v39
                                                            } else {
                                                                v40 := GenS3{F0: (v29 | 42), F1: 460, F2: int(v29)}
                                                                fmt.Println(v18)
                                                                v41 := v6[128]
                                                                _ = v40
                                                                _ = v41
                                                            }
                                                        }
                                                    default:
                                                        v42 := GenS4{F0: 127, F1: v18, F2: ("a" + "a")}
                                                        _ = v42
                                                    }
                                                    switch (fi16 ^ (fi16 << ((0) & 63))) {
                                                    case 127:
                                                        delete(v6, -515)
                                                    }
                                                    v43 := GenS5{F0: (fi16 % (((-(v21))) | 1))}
                                                    _ = v43
                                                }
                                                if true {
                                                    switch fi16 {
                                                    case -1:
                                                        switch (-(fi16)) {
                                                        case -1:
                                                            var v44 int = 127
                                                            v3 = append(v3, "")
                                                            if false {
                                                                v45 := v5[v18]
                                                                switch v44 {
                                                                case 42:
                                                                    v1 += v10
                                                                default:
                                                                    v2 = v2
                                                                }
                                                                var v46 uint8 = (^(v17))
                                                                _ = v45
                                                                _ = v46
                                                            }
                                                            _ = v44
                                                        case 42:
                                                            if false {
                                                                v47 := GenS1{F0: (v21 & (v21 * (v21 >> ((1) & 63))))}
                                                                switch (-(fi19)) {
                                                                case 853:
                                                                    v48 := v5[(v18 | 65535)]
                                                                    _ = v48
                                                                case -9223372036854775808:
                                                                    v29 = (v29 ^ v29)
                                                                default:
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        fmt.Println(v9)
                                                                    case "":
                                                                        v10 += 709
                                                                    default:
                                                                        fmt.Println(v18, v1, v9)
                                                                    }
                                                                }
                                                                var v49 uint64 = (^(v28))
                                                                if false {
                                                                    v50 := GenS3{F0: (^(v29)), F1: 18446744073709551615, F2: int(v9)}
                                                                    _ = v50
                                                                } else {
                                                                    fmt.Println(fi19)
                                                                    v3 = append(v3, "hello")
                                                                    var v51 float32 = v2
                                                                    _ = v51
                                                                }
                                                                _ = v47
                                                                _ = v49
                                                            } else {
                                                                v29 = 127
                                                                switch ("hello" + "foo bar") {
                                                                case "foo bar":
                                                                    v29 = -1
                                                                case "a":
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        v21 = 1
                                                                    default:
                                                                        v1 += (v1 - (v10 + 0))
                                                                    }
                                                                    if true {
                                                                        if (("" == ("" + "\n\t")) && true) {
                                                                            v52 := GenS4{F0: (v17 ^ uint8(v18)), F1: 2147483648, F2: ""}
                                                                            var v53 uint = 257
                                                                            var v54 bool = false
                                                                            _ = v52
                                                                            _ = v53
                                                                            _ = v54
                                                                        } else {
                                                                            var v55 uint32 = 561
                                                                            _ = v55
                                                                        }
                                                                    }
                                                                    v56 := GenS3{F0: (^(v29)), F1: 18446744073709551615, F2: 819}
                                                                    _ = v56
                                                                case "\n\t":
                                                                    _ = len(v3)
                                                                }
                                                            }
                                                        }
                                                    case 42:
                                                        if len(v3) > 0 {
                                                            v3[0] = "hello"
                                                        }
                                                    case 0:
                                                        fmt.Println(v10)
                                                    }
                                                } else {
                                                    if (1e10 < v2) {
                                                        fmt.Println(v17, v17)
                                                    }
                                                    switch (fi16 | -9223372036854775808) {
                                                    case 127:
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    default:
                                                        if (v18 != v18) {
                                                            switch fi19 {
                                                            case 42:
                                                                if true {
                                                                    v21 = (-(v21))
                                                                    if len(v3) > 0 {
                                                                        v3[0] = ""
                                                                    }
                                                                    if false {
                                                                        v57 := GenS6{F0: 0, F1: 65535, F2: 508}
                                                                        _ = v57
                                                                    } else {
                                                                        var v58 float32 = (v2 / -2.25)
                                                                        _ = len(v4)
                                                                        if false {
                                                                            if (((true || true) && (!false)) && true) {
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = ""
                                                                                }
                                                                                var v59 uint8 = (v17 - (v17 << ((uint(v18)) & 63)))
                                                                                switch int(v28) {
                                                                                case -1:
                                                                                    switch (("foo bar" + "a") + ("hello" + "hello")) {
                                                                                    case "hello":
                                                                                        if false {
                                                                                            var v60 int64 = 856
                                                                                            switch v21 {
                                                                                            case -35:
                                                                                                if ("" <= "foo bar") {
                                                                                                    v21 -= (-(fi16))
                                                                                                    _ = len(v6)
                                                                                                }
                                                                                                v10 -= (v1 & 446)
                                                                                                v61 := GenS7{F0: 65535, F1: int64(v59), F2: (^(v29))}
                                                                                                _ = v61
                                                                                            case 0:
                                                                                                if (true && true) {
                                                                                                    v62 := GenS1{F0: fi16}
                                                                                                    v63 := GenS8{F0: (v59 >> ((0) & 63)), F1: "hello", F2: v21}
                                                                                                    switch ("\n\t" + "\n\t") {
                                                                                                    case "hello":
                                                                                                        v21 += (v21 * (fi19 + (-(fi16))))
                                                                                                        switch (("a" + "") + ("" + ("foo bar" + ("a" + ("a" + ""))))) {
                                                                                                        case "\n\t":
                                                                                                            fmt.Println(v29)
                                                                                                        }
                                                                                                    case "\n\t":
                                                                                                        var v64 string = "hello"
                                                                                                        _ = v64
                                                                                                    case "":
                                                                                                        v58 -= (v2 + float32(((((0.0 / -2.25) * 3.14159) / 1.0) / float64((-2.25 + v58)))))
                                                                                                    default:
                                                                                                        fmt.Println(fi19, v60, v10)
                                                                                                    }
                                                                                                    _ = v62
                                                                                                    _ = v63
                                                                                                } else {
                                                                                                    if ("foo bar" == ("\n\t" + "")) {
                                                                                                        var v65 float64 = 0.1
                                                                                                        _ = v65
                                                                                                    }
                                                                                                    var v66 string = ((("" + "\n\t") + "hello") + (("hello" + "hello") + "hello"))
                                                                                                    fmt.Println(v17, v29, v21)
                                                                                                    v21 += -55
                                                                                                    _ = v66
                                                                                                }
                                                                                            case 9223372036854775807:
                                                                                                var v67 bool = true
                                                                                                _ = v67
                                                                                            default:
                                                                                                v58 = ((v2 * float32(3.14159)) * -2.25)
                                                                                            }
                                                                                            switch (fi19 | 127) {
                                                                                            case 2147483647:
                                                                                                var v68 uint = 1
                                                                                                _ = v68
                                                                                            }
                                                                                            _ = v60
                                                                                        }
                                                                                    default:
                                                                                        fmt.Println(v59, v29)
                                                                                        v69 := GenS8{F0: v17, F1: "", F2: (fi16 * v21)}
                                                                                        v70 := GenS4{F0: (v17 - 115), F1: (^(v18)), F2: (("\n\t" + "a") + ("a" + ""))}
                                                                                        _ = v69
                                                                                        _ = v70
                                                                                    }
                                                                                    v28 = (v28 >> ((979) & 63))
                                                                                    v59 += v17
                                                                                    var v71 int32 = v9
                                                                                    _ = v71
                                                                                case -846:
                                                                                    fmt.Println(v10)
                                                                                }
                                                                                _ = v59
                                                                            }
                                                                        } else {
                                                                            if len(v3) > 0 {
                                                                                v3[0] = "\n\t"
                                                                            }
                                                                            if ((v18 ^ (v18 - (^(v18)))) != v18) {
                                                                                v29 += v29
                                                                                if true {
                                                                                    v58 += -2.25
                                                                                    fmt.Println(v9, v17, v9)
                                                                                    if (4294967295 <= 42) {
                                                                                        switch "\n\t" {
                                                                                        case "":
                                                                                            switch ("\n\t" + "\n\t") {
                                                                                            case "\n\t":
                                                                                                if false {
                                                                                                    v9 = (v9 - 0)
                                                                                                    v28 = v28
                                                                                                } else {
                                                                                                    switch int(v10) {
                                                                                                    case 32767:
                                                                                                        if false {
                                                                                                            v72 := GenS8{F0: (v17 + uint8(v1)), F1: ((("\n\t" + "\n\t") + "\n\t") + "a"), F2: -1}
                                                                                                            if true {
                                                                                                                switch "hello" {
                                                                                                                case "foo bar":
                                                                                                                    var v73 uint64 = v28
                                                                                                                    _ = v73
                                                                                                                case "":
                                                                                                                    var v74 string = "hello"
                                                                                                                    _ = v74
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v72
                                                                                                        }
                                                                                                    case 32768:
                                                                                                        v1 += v10
                                                                                                        v17 += v17
                                                                                                    default:
                                                                                                        fmt.Println(v9)
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = ""
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = "foo bar"
                                                                                        }
                                                                                    }
                                                                                }
                                                                                if ((v28 > (v28 ^ v28)) && false) {
                                                                                    var v75 uint16 = v10
                                                                                    _ = len(v5)
                                                                                    _ = v75
                                                                                }
                                                                            }
                                                                            _ = len(v4)
                                                                        }
                                                                        _ = v58
                                                                    }
                                                                    switch 0 {
                                                                    case -903:
                                                                        v2 += v2
                                                                    case 32767:
                                                                        var v76 uint = 255
                                                                        _ = v76
                                                                    }
                                                                }
                                                            }
                                                            var v77 uint16 = (v1 << ((2) & 63))
                                                            var v78 float64 = -2.25
                                                            _ = v77
                                                            _ = v78
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                        switch 1 {
                                        case 42:
                                            var v79 int16 = -32768
                                            _ = v79
                                        }
                                        v80 := GenS4{F0: 1, F1: v18, F2: ((("foo bar" + "a") + "foo bar") + "foo bar")}
                                        _ = v28
                                        _ = v29
                                        _ = v80
                                    } else {
                                        if true {
                                            v1 = (^(v1))
                                            v81 := GenS1{F0: 2147483647}
                                            _ = v81
                                        } else {
                                            v1 = (v10 & uint16(v17))
                                        }
                                        if false {
                                            v9 = (v9 % ((v9) | 1))
                                            v21 = v21
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                    }
                                case 67:
                                    v2 += v2
                                }
                            case 32768:
                                var v82 uint64 = 2
                                _ = v82
                            }
                            if len(v3) > 0 {
                                v3[0] = "hello"
                            }
                        }
                        v2 = 3.14159
                        if len(v3) > 0 {
                            v3[0] = "\n\t"
                        }
                        if (!false) {
                            if len(v3) > 0 {
                                v3[0] = ("hello" + "a")
                            }
                        }
                        v1 = v1
                        v10 -= v1
                        v18 += v18
                        _ = v21
                    }
                    v83 := GenS9{F0: uint8(v1), F1: v1, F2: 42}
                    _ = fi19
                    _ = v83
                }
                v84 := GenS10{F0: -156, F1: 164}
                _ = v18
                _ = v84
            } else {
                var v85 int16 = 0
                var v86 uint64 = 2
                _ = v85
                _ = v86
            }
            _ = fi16
            _ = v17
        }
        v10 = v1
        _ = v15
    }
    switch "foo bar" {
    case "foo bar":
        v87 := GenS11{F0: 742, F1: v1}
        _ = v87
    default:
        v9 += 0
    }
    switch 42 {
    case 727:
        for fi88 := 0; fi88 < 8; fi88++ {
            for fi89 := 0; fi89 < 14; fi89++ {
                fmt.Println(v9, v2, v1)
                if true {
                    v10 -= v1
                    switch (("\n\t" + "hello") + "hello") {
                    case "foo bar":
                        switch fi89 {
                        case -9223372036854775808:
                            if ((128 < (fi89 ^ 605)) || true) {
                                switch 1 {
                                case 42:
                                    v10 -= (v10 >> ((1) & 63))
                                }
                                v2 = (3.14159 + 1e10)
                                v2 = v2
                            }
                        case 2147483647:
                            v90 := v4[2]
                            _ = v90
                        case 42:
                            if false {
                                var v91 bool = true
                                switch "foo bar" {
                                case "hello":
                                    fmt.Println(v2)
                                case "foo bar":
                                    if v91 {
                                        v9 += v9
                                    } else {
                                        var v92 int8 = 42
                                        fmt.Println(v9, fi89)
                                        _ = v92
                                    }
                                }
                                _ = v91
                            } else {
                                v2 = float32(1e10)
                                v2 = 0.1
                                fmt.Println(v1)
                                v93 := GenS6{F0: 881, F1: 255, F2: 0}
                                _ = v93
                            }
                        default:
                            fmt.Println(v10, v10)
                        }
                    }
                    v94 := GenS9{F0: 172, F1: v1, F2: (v9 / (((v9 - v9)) | 1))}
                    v10 += (v10 >> ((4294967295) & 63))
                    fmt.Println(fi89, fi88, v1)
                    _ = v94
                } else {
                    _ = len(v6)
                    fmt.Println(v10, fi88, v10)
                }
                v2 += ((3.14159 / ((v2 + (v2 * (0.0 * ((((float32(0.1) + v2) - (v2 + (3.14159 - v2))) * -2.25) - ((v2 * v2) / (float32(1e10) / v2)))))) / v2)) * (v2 / float32(1e10)))
                var v95 uint8 = 145
                v1 -= (v1 - 0)
                _ = fi89
                _ = v95
            }
            _ = fi88
        }
    default:
        if len(v3) > 0 {
            v3[0] = "hello"
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
