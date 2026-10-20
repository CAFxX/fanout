package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 string; F1 int; F2 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint32 }
func genFn0() int {
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 uint32 = 176
        fmt.Println(fi6)
        _ = fi6
        _ = v7
    }
    v8 := []int{128}
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS1{F0: "\n\t"}
        var v11 float64 = (0.0 / (((((-2.25 * (((3.14159 * 1.5) / 1.5) - 3.14159)) / 0.1) * -2.25) - 1.5) + (1.5 / 1.0)))
        switch (fi9 >> ((18446744073709551615) & 63)) {
        case 42:
            if len(v8) > 0 {
                v8[0] = fi9
            }
        case -117:
            var v12 uint = 18446744073709551615
            _ = v12
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    var v13 uint64 = 1
    var v14 float32 = 0.0
    _ = v13
    _ = v14
    _ = v8
    return 32767
}
func genFn1() bool {
    var v15 float64 = 1e10
    v16 := GenS1{F0: "\n\t"}
    for fi17 := 0; fi17 < 17; fi17++ {
        if (false || (2 == 42)) {
            var v18 uint = 42
            fmt.Println(v15, fi17)
            var v19 bool = false
            _ = v18
            _ = v19
        }
        _ = fi17
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        v15 = v15
        var v21 float32 = (0.1 / 1.5)
        for fi22 := 0; fi22 < 2; fi22++ {
            v23 := GenS2{F0: 1}
            v24 := []int{fi20, (fi22 / ((genFn0()) | 1)), (fi22 * (fi22 ^ (fi20 & (fi22 * (fi20 | fi22)))))}
            v21 = (v21 * v21)
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi20
        _ = v21
    }
    var v25 float32 = float32(v15)
    _ = v15
    _ = v16
    _ = v25
    return false
}

func main() {
    v1 := []uint32{42, 42, 65535, 4294967295}
    var v2 float64 = 0.1
    var v3 int = -213
    v4 := []string{(("foo bar" + ((("\n\t" + ("" + "")) + "hello") + "foo bar")) + ("foo bar" + ("\n\t" + "a")))}
    var v5 bool = ((1e10 == (0.1 / ((0.0 + (((1e10 * (((0.0 / (1.5 / (3.14159 - (0.0 + 1e10)))) * -2.25) + 3.14159)) - float32(0.0)) / 3.14159)) / 1e10))) || (1000 != 1000))
    var v26 uint32 = 1
    v4 = append(v4, ("\n\t" + (("" + "hello") + "\n\t")))
    v26 -= v26
    v5 = v5
    switch "hello" {
    case "\n\t":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "a":
        if v5 {
            if true {
                v27 := GenS2{F0: 71}
                _ = v27
            } else {
                v2 = (v2 + v2)
                for fi28 := 0; fi28 < 4; fi28++ {
                    var v29 float32 = float32(v2)
                    _ = fi28
                    _ = v29
                }
            }
        } else {
            v30 := GenS2{F0: 42}
            var v31 string = "a"
            _ = v30
            _ = v31
        }
    case "":
        if ((v3 | (v3 + v3)) == v3) {
            fmt.Println(v3)
            var v32 int = 1
            _ = v32
        }
    }
    fmt.Println(v26, v3)
    var v33 float32 = (3.14159 / 1e10)
    switch v3 {
    case 0:
        v33 = v33
    case -9223372036854775808:
        for fi34 := 0; fi34 < 12; fi34++ {
            var v35 int8 = -128
            if ((!v5) || v5) {
                for fi36 := 0; fi36 < 11; fi36++ {
                    v37 := GenS3{F0: 116, F1: 0}
                    _ = fi36
                    _ = v37
                }
                v38 := GenS4{F0: "foo bar", F1: fi34, F2: -1}
                if v5 {
                    fmt.Println(fi34)
                    var v39 uint8 = 2
                    _ = v39
                }
                _ = v38
            } else {
                v40 := GenS4{F0: ("a" + ((((("\n\t" + ("" + "")) + ("" + (("" + "\n\t") + ("a" + (("foo bar" + "hello") + ("\n\t" + ("" + ("" + "foo bar")))))))) + "foo bar") + (((("a" + "a") + "a") + "a") + ("foo bar" + (("hello" + "hello") + "\n\t")))) + "foo bar")), F1: genFn0(), F2: 1}
                _ = v40
            }
            _ = fi34
            _ = v35
        }
    default:
        switch ("hello" + "hello") {
        case "foo bar":
            switch (("a" + "\n\t") + "\n\t") {
            case "a":
                fmt.Println(v26, v33)
            case "hello":
                for fi41 := 0; fi41 < 16; fi41++ {
                    if false {
                        v42 := GenS3{F0: 96, F1: uint8(v26)}
                        _ = v42
                    }
                    var v43 float32 = v33
                    for fi44 := 0; fi44 < 9; fi44++ {
                        v26 = (v26 | 2)
                        _ = fi44
                    }
                    _ = fi41
                    _ = v43
                }
            default:
                var v45 uint16 = 659
                _ = v45
            }
        case "hello":
            if (false || (!v5)) {
                v3 = 32768
            }
        default:
            var v46 int64 = 1
            _ = v46
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v47 int = (v3 + v3)
    _ = len(v1)
    if false {
        var v48 float64 = (v2 / float64(float32(v2)))
        if (42 < 42) {
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v48)
            case "hello":
                switch ("\n\t" + (("" + "foo bar") + "hello")) {
                case "foo bar":
                    var v49 uint32 = (v26 % ((v26) | 1))
                    _ = v49
                }
            default:
                if false {
                    v3 = (v47 << ((255) & 63))
                    v2 += 0.1
                    fmt.Println(v47, v48, v3)
                } else {
                    for fi50 := 0; fi50 < 5; fi50++ {
                        v51 := GenS4{F0: ((((("foo bar" + "foo bar") + "foo bar") + "") + "hello") + "foo bar"), F1: (v47 >> ((271) & 63)), F2: 0}
                        if (!v5) {
                            var v52 int32 = -940
                            if ((fi50 * (fi50 % (((v47 ^ v3)) | 1))) == v3) {
                                for fi53 := 0; fi53 < 19; fi53++ {
                                    v52 = v52
                                    v3 = 9223372036854775807
                                    _ = fi53
                                }
                                v33 += 1.5
                                v54 := GenS2{F0: -128}
                                _ = v54
                            } else {
                                if (65535 > v26) {
                                    if true {
                                        var v55 int = (v3 | v47)
                                        fmt.Println(v52)
                                        var v56 int8 = 127
                                        _ = v55
                                        _ = v56
                                    }
                                    fmt.Println(v26)
                                    for fi57 := 0; fi57 < 6; fi57++ {
                                        switch genFn0() {
                                        case 32767:
                                            v48 = 3.14159
                                        }
                                        _ = fi57
                                    }
                                }
                                if (v33 != ((v33 / 0.1) * v33)) {
                                    switch (("foo bar" + ("a" + "hello")) + "foo bar") {
                                    case "\n\t":
                                        for fi58 := 0; fi58 < 18; fi58++ {
                                            v48 = v48
                                            if (!(-1 == 0)) {
                                                var v59 float32 = (float32(v48) / ((0.1 * v33) + v33))
                                                var v60 uint8 = 0
                                                _ = v59
                                                _ = v60
                                            }
                                            v2 = 0.1
                                            _ = fi58
                                        }
                                    case "":
                                        switch "a" {
                                        case "\n\t":
                                            if v5 {
                                                v5 = v5
                                            } else {
                                                for fi61 := 0; fi61 < 19; fi61++ {
                                                    fmt.Println(v52, fi61, v48)
                                                    if (v5 && (0 <= 1)) {
                                                        v62 := GenS5{F0: 639}
                                                        var v63 uint32 = 65535
                                                        _ = v62
                                                        _ = v63
                                                    } else {
                                                        fmt.Println(fi50, fi50)
                                                    }
                                                    _ = fi61
                                                }
                                                if v5 {
                                                    if ("hello" != ("foo bar" + "foo bar")) {
                                                        v47 -= (fi50 ^ fi50)
                                                        for fi64 := 0; fi64 < 10; fi64++ {
                                                            var v65 int64 = 9223372036854775807
                                                            _ = fi64
                                                            _ = v65
                                                        }
                                                        var v66 int16 = -32768
                                                        _ = v66
                                                    }
                                                    v2 -= v2
                                                    v67 := GenS5{F0: 9223372036854775807}
                                                    _ = v67
                                                } else {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                            }
                                        default:
                                            var v68 uint64 = 239
                                            _ = v68
                                        }
                                    default:
                                        fmt.Println(v33, fi50)
                                    }
                                    if v5 {
                                        var v69 int64 = 437
                                        _ = v69
                                    }
                                    if (!true) {
                                        if len(v4) > 0 {
                                            v4[0] = ("\n\t" + "a")
                                        }
                                    } else {
                                        var v70 uint64 = 42
                                        var v71 float64 = 1e10
                                        var v72 float64 = (0.1 - v2)
                                        _ = v70
                                        _ = v71
                                        _ = v72
                                    }
                                }
                            }
                            var v73 uint8 = 42
                            _ = v52
                            _ = v73
                        } else {
                            var v74 int8 = 1
                            switch "hello" {
                            case "a":
                                if len(v1) > 0 {
                                    v1[0] = (v26 & v26)
                                }
                            default:
                                var v75 uint32 = 1
                                _ = v75
                            }
                            v76 := GenS6{F0: 0, F1: (!false)}
                            _ = v74
                            _ = v76
                        }
                        v33 += float32(((-2.25 * 3.14159) - v2))
                        _ = fi50
                        _ = v51
                    }
                    switch v3 {
                    case -2147483648:
                        if (!((0 < 0) && (!v5))) {
                            var v77 string = "a"
                            _ = v77
                        }
                    case 0:
                        fmt.Println(v47, v26)
                    }
                    if len(v4) > 0 {
                        v4[0] = "foo bar"
                    }
                }
            }
            if true {
                v78 := GenS5{F0: 9223372036854775807}
                v79 := GenS4{F0: "foo bar", F1: (v47 + v47), F2: 42}
                _ = v78
                _ = v79
            }
        } else {
            var v80 string = "a"
            for fi81 := 0; fi81 < 9; fi81++ {
                v47 = 379
                v48 = float64(0.0)
                _ = fi81
            }
            v1 = append(v1, v26)
            _ = v80
        }
        _ = v48
    } else {
        for fi82 := 0; fi82 < 19; fi82++ {
            var v83 uint32 = 1
            v4 = append(v4, "\n\t")
            _ = fi82
            _ = v83
        }
    }
    if (!false) {
        if (1 < 255) {
            for fi84 := 0; fi84 < 18; fi84++ {
                v3 = genFn0()
                v2 = v2
                v85 := GenS6{F0: -1, F1: ((!v5) && (true && true))}
                _ = fi84
                _ = v85
            }
            if true {
                v2 += (v2 * (((((v2 + ((0.0 / 0.1) + (((v2 * ((0.1 - (v2 * (v2 + float64(v33)))) + v2)) - float64((v33 / (1e10 / -2.25)))) / float64((3.14159 * v33))))) - (v2 - v2)) + ((((1e10 / (v2 + (v2 - 3.14159))) / float64(1e10)) / 1.5) * (v2 - v2))) / (v2 + (1.5 + -2.25))) - v2))
                v86 := GenS7{F0: v5}
                _ = v86
            } else {
                _ = len(v1)
            }
        } else {
            if v5 {
                var v87 int32 = 2147483647
                _ = v87
            } else {
                if (float64((v33 + (v33 - v33))) == 0.1) {
                    fmt.Println(v47, v33)
                }
                var v88 bool = (!(v26 >= 65535))
                fmt.Println(v2, v47)
                _ = v88
            }
            for fi89 := 0; fi89 < 5; fi89++ {
                var v90 uint32 = 0
                switch "" {
                case "\n\t":
                    v5 = (231 != 0)
                case "hello":
                    for fi91 := 0; fi91 < 12; fi91++ {
                        if v5 {
                            var v92 int = (v3 | (v47 - 32767))
                            var v93 string = "foo bar"
                            _ = v92
                            _ = v93
                        }
                        _ = fi91
                    }
                case "foo bar":
                    if v5 {
                        v33 = v33
                        if ((1e10 * v33) != float32(v2)) {
                            var v94 uint64 = 1
                            v33 = (v33 * 1e10)
                            _ = v94
                        } else {
                            var v95 float32 = v33
                            var v96 int32 = -334
                            _ = v95
                            _ = v96
                        }
                        switch ("\n\t" + (("a" + "hello") + "a")) {
                        case "hello":
                            v1 = append(v1, 4294967295)
                        default:
                            v5 = (4294967295 > 0)
                        }
                    }
                }
                v2 -= v2
                _ = fi89
                _ = v90
            }
            for fi97 := 0; fi97 < 19; fi97++ {
                if (877 < 2) {
                    v26 = 2
                    v1 = append(v1, v26)
                    if false {
                        v33 += 0.1
                        if v5 {
                            for fi98 := 0; fi98 < 2; fi98++ {
                                v99 := GenS1{F0: "hello"}
                                v47 = (fi97 & v3)
                                _ = fi98
                                _ = v99
                            }
                            var v100 float32 = (v33 - ((v33 + float32((1e10 - (v2 * float64((((0.1 * v33) * ((0.1 * 0.0) + (((-2.25 + (1e10 * float32((v2 - float64((v33 + float32((v2 + v2)))))))) - v33) / 3.14159))) / v33)))))) + -2.25))
                            _ = v100
                        } else {
                            v101 := GenS8{F0: v26}
                            fmt.Println(v2)
                            fmt.Println(v2, v47)
                            _ = v101
                        }
                    }
                }
                _ = fi97
            }
        }
        var v102 int16 = -922
        var v103 uint8 = 2
        _ = v102
        _ = v103
    } else {
        var v104 uint = 65535
        var v105 bool = v5
        switch 0 {
        case -2147483648:
            v106 := GenS1{F0: "a"}
            _ = v106
        case 4294967295:
            v105 = v105
        default:
            var v107 uint32 = v26
            _ = v107
        }
        _ = v104
        _ = v105
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v33
    _ = v4
    _ = v47
    _ = v5
}
