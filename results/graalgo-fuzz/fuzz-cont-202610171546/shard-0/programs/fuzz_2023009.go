package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint8; F2 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64; F1 uint8 }
type GenS8 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[uint32]string)
    v3 := []int{0, -9223372036854775808}
    var v4 uint8 = 1
    v5 := GenS1{F0: ""}
    v1 = (-(v1))
    var v6 uint64 = 18446744073709551615
    v4 = v4
    fmt.Println(v4, v1, v1)
    if (uint16(v6) < 790) {
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 int = (fi7 ^ (fi7 << ((18446744073709551615) & 63)))
            _ = fi7
            _ = v8
        }
        v1 += (v1 - 42)
    }
    switch "hello" {
    case "":
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := GenS1{F0: "foo bar"}
            switch ((("hello" + ("" + "")) + "\n\t") + "hello") {
            case "a":
                for fi11 := 0; fi11 < 14; fi11++ {
                    if len(v3) > 0 {
                        v3[0] = (fi9 | fi11)
                    }
                    _ = fi11
                }
            case "":
                for fi12 := 0; fi12 < 12; fi12++ {
                    switch ("hello" + "\n\t") {
                    case "foo bar":
                        v6 = v6
                    case "\n\t":
                        fmt.Println(v6, v1)
                    }
                    if true {
                        fmt.Println(v6, v6)
                        var v13 uint16 = 255
                        _ = v13
                    }
                    switch (fi9 / ((fi9) | 1)) {
                    case -1:
                        v4 = 75
                    }
                    fmt.Println(v6, fi9)
                    _ = fi12
                }
            }
            v6 = (v6 << ((0) & 63))
            _ = fi9
            _ = v10
        }
    case "a":
        fmt.Println(v6, v1, v1)
    }
    if (0 >= -952) {
        var v14 string = ("a" + "hello")
        var v15 int8 = 14
        _ = v14
        _ = v15
    } else {
        switch "\n\t" {
        case "hello":
            for fi16 := 0; fi16 < 11; fi16++ {
                v17 := GenS1{F0: (("hello" + "hello") + "a")}
                delete(v2, 42)
                _ = fi16
                _ = v17
            }
        }
        if true {
            for fi18 := 0; fi18 < 9; fi18++ {
                if len(v3) > 0 {
                    v3[0] = 9223372036854775807
                }
                _ = fi18
            }
            v1 += (v1 >> ((255) & 63))
        }
        var v19 uint64 = v6
        _ = v19
    }
    v1 -= (v1 >> ((65535) & 63))
    v20 := GenS2{F0: -558, F1: v4, F2: ""}
    v21 := v2[906]
    var v22 int = -2147483648
    switch ((v21 + v21) + "foo bar") {
    case "\n\t":
        switch (v22 >> ((255) & 63)) {
        case 32767:
            v6 += 18446744073709551615
        case 924:
            var v23 uint8 = v4
            _ = v23
        }
    case "":
        var v24 int32 = -2147483648
        _ = v24
    case "a":
        if true {
            v22 += (v22 * (v22 << ((519) & 63)))
            fmt.Println(v22, v22, v1)
        }
    default:
        var v25 int = v22
        _ = v25
    }
    var v26 uint = 632
    if (!false) {
        var v27 string = v21
        if true {
            if ((!(((!(!false)) || false) || true)) && false) {
                if true {
                    var v28 uint32 = 65535
                    if len(v3) > 0 {
                        v3[0] = (v22 & -352)
                    }
                    _ = v28
                }
                var v29 bool = ((!false) && (((0.1 - 0.0) - (1.5 - 3.14159)) > 1e10))
                _ = v29
            } else {
                delete(v2, uint32(v26))
                var v30 int32 = 42
                _ = v30
            }
        }
        _ = v27
    }
    v6 -= v6
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "a" {
    case "hello":
        switch v22 {
        case 2147483647:
            v6 += 4294967295
        case 42:
            v31 := GenS2{F0: -894, F1: 42, F2: ("" + v21)}
            _ = v31
        case -9223372036854775808:
            if false {
                if false {
                    for fi32 := 0; fi32 < 11; fi32++ {
                        v1 = (v1 >> ((699) & 63))
                        var v33 int = (fi32 + (^(fi32)))
                        switch "" {
                        case "\n\t":
                            var v34 int8 = -1
                            _ = v34
                        case "foo bar":
                            fmt.Println(v33, v1, v26)
                        default:
                            v4 += 127
                        }
                        _ = fi32
                        _ = v33
                    }
                    var v35 int32 = -1
                    if false {
                        v6 += v6
                        for fi36 := 0; fi36 < 3; fi36++ {
                            v4 = v4
                            v6 += v6
                            if (v26 < v26) {
                                if ((v21 + ((v21 + "\n\t") + (v21 + "a"))) == v21) {
                                    fmt.Println(v22)
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                } else {
                                    switch (v21 + ("foo bar" + v21)) {
                                    case "foo bar":
                                        if (!true) {
                                            var v37 uint16 = 42
                                            if (677 > (^(v26))) {
                                                if true {
                                                    switch (-(v22)) {
                                                    case 32768:
                                                        if true {
                                                            v22 = fi36
                                                            switch 369 {
                                                            case 4294967295:
                                                                fmt.Println(v6, fi36, v1)
                                                            case 1:
                                                                v4 += (^(v4))
                                                            }
                                                            var v38 int = (v22 ^ fi36)
                                                            _ = v38
                                                        }
                                                    case 0:
                                                        v3 = append(v3, 40)
                                                    case 9223372036854775807:
                                                        var v39 int8 = 0
                                                        _ = v39
                                                    default:
                                                        var v40 string = v21
                                                        _ = v40
                                                    }
                                                    if ((1.5 * 0.1) <= float32(1.5)) {
                                                        v26 += (v26 >> ((65535) & 63))
                                                        v22 -= fi36
                                                    } else {
                                                        if false {
                                                            v35 += (v35 >> ((18446744073709551615) & 63))
                                                        } else {
                                                            if true {
                                                                if (uint32(v4) == uint32(v37)) {
                                                                    for fi41 := 0; fi41 < 16; fi41++ {
                                                                        if false {
                                                                            var v42 uint64 = (^(v6))
                                                                            v43 := GenS3{F0: 439}
                                                                            _ = v42
                                                                            _ = v43
                                                                        } else {
                                                                            var v44 float32 = 1.5
                                                                            _ = v44
                                                                        }
                                                                        v2[952] = v21
                                                                        var v45 int = fi41
                                                                        _ = fi41
                                                                        _ = v45
                                                                    }
                                                                    for fi46 := 0; fi46 < 6; fi46++ {
                                                                        var v47 int64 = (^(v1))
                                                                        _ = fi46
                                                                        _ = v47
                                                                    }
                                                                    if true {
                                                                        if true {
                                                                            if (-1 > 1) {
                                                                                v48 := GenS1{F0: v21}
                                                                                _ = v48
                                                                            } else {
                                                                                v6 = (v6 % (((^(v6))) | 1))
                                                                                v37 -= 255
                                                                                v4 = (v4 - v4)
                                                                            }
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                        }
                                                                        v35 = v35
                                                                        v22 -= (v22 << (((v26 - (v26 / ((17) | 1)))) & 63))
                                                                    } else {
                                                                        v4 += (v4 / (((v4 ^ (v4 % ((uint8(v26)) | 1)))) | 1))
                                                                        for fi49 := 0; fi49 < 14; fi49++ {
                                                                            var v50 float32 = 1e10
                                                                            _ = fi49
                                                                            _ = v50
                                                                        }
                                                                    }
                                                                }
                                                                var v51 bool = ((2 != 42) || (!true))
                                                                _ = v51
                                                            }
                                                        }
                                                    }
                                                    var v52 float64 = 1.5
                                                    _ = v52
                                                } else {
                                                    v53 := GenS2{F0: v35, F1: 128, F2: v21}
                                                    _ = v53
                                                }
                                                if (4294967295 > uint32(v6)) {
                                                    fmt.Println(v26, v26, v26)
                                                    if false {
                                                        v54 := GenS4{F0: (v21 + ((v21 + v21) + (v21 + v21))), F1: 65535}
                                                        _ = v54
                                                    }
                                                    switch v22 {
                                                    case 0:
                                                        if true {
                                                            var v55 int32 = -2147483648
                                                            _ = v55
                                                        }
                                                    case -881:
                                                        v56 := GenS1{F0: "foo bar"}
                                                        _ = v56
                                                    case 516:
                                                        fmt.Println(v35, v6, v35)
                                                    }
                                                } else {
                                                    v21 = (("" + ("" + (("\n\t" + v21) + v21))) + v21)
                                                    var v57 uint32 = 2147483648
                                                    _ = v57
                                                }
                                            }
                                            _ = v37
                                        }
                                    default:
                                        for fi58 := 0; fi58 < 10; fi58++ {
                                            var v59 uint16 = 255
                                            v3 = append(v3, 2147483647)
                                            v60 := GenS3{F0: int16(v6)}
                                            _ = fi58
                                            _ = v59
                                            _ = v60
                                        }
                                    }
                                    for fi61 := 0; fi61 < 1; fi61++ {
                                        var v62 uint = 376
                                        switch (((v21 + (v21 + (((v21 + "\n\t") + v21) + ""))) + ((v21 + ((v21 + v21) + v21)) + "foo bar")) + ((v21 + (v21 + ((("a" + "a") + v21) + ("a" + "\n\t")))) + v21)) {
                                        case "hello":
                                            _ = len(v3)
                                        default:
                                            var v63 bool = (-871 >= v22)
                                            _ = v63
                                        }
                                        _ = fi61
                                        _ = v62
                                    }
                                    switch v21 {
                                    case "hello":
                                        switch (^(fi36)) {
                                        case 94:
                                            v64 := GenS3{F0: -480}
                                            _ = v64
                                        case 2147483647:
                                            v65 := GenS4{F0: v21, F1: 2147483647}
                                            _ = v65
                                        default:
                                            fmt.Println(v1, v21)
                                        }
                                    default:
                                        if false {
                                            for fi66 := 0; fi66 < 9; fi66++ {
                                                v1 = 304
                                                var v67 int16 = -625
                                                _ = fi66
                                                _ = v67
                                            }
                                            var v68 int64 = (v1 >> ((v26) & 63))
                                            _ = v68
                                        }
                                    }
                                }
                                if true {
                                    for fi69 := 0; fi69 < 10; fi69++ {
                                        v70 := GenS5{F0: fi36, F1: (^(v26))}
                                        v6 += 796
                                        _ = fi69
                                        _ = v70
                                    }
                                    v26 = v26
                                    for fi71 := 0; fi71 < 12; fi71++ {
                                        if len(v3) > 0 {
                                            v3[0] = fi71
                                        }
                                        var v72 float32 = 1.5
                                        _ = fi71
                                        _ = v72
                                    }
                                } else {
                                    fmt.Println(v26)
                                    v6 = v6
                                    v21 = (("foo bar" + v21) + ("foo bar" + "\n\t"))
                                }
                            } else {
                                v73 := GenS6{F0: (v4 / ((v4) | 1))}
                                v74 := GenS7{F0: v1, F1: (v4 << (((v26 % ((0) | 1))) & 63))}
                                _ = v73
                                _ = v74
                            }
                            _ = fi36
                        }
                        var v75 int = v22
                        _ = v75
                    }
                    _ = v35
                } else {
                    delete(v2, 2147483647)
                    var v76 bool = false
                    _ = v76
                }
                if (("foo bar" + v21) < ("" + (((v21 + v21) + "\n\t") + (("" + ("\n\t" + (v21 + "a"))) + (v21 + (v21 + (("a" + ("foo bar" + v21)) + (v21 + (((("" + ("hello" + v21)) + "") + ((v21 + "") + (v21 + "\n\t"))) + "a"))))))))) {
                    var v77 uint8 = (v4 % (((v4 * v4)) | 1))
                    _ = v77
                } else {
                    switch v22 {
                    case -9223372036854775808:
                        v78 := GenS8{F0: 32767, F1: 127}
                        _ = v78
                    default:
                        v21 = (v21 + v21)
                    }
                }
            }
        }
    default:
        v4 = (v4 | (v4 & (^(v4))))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
