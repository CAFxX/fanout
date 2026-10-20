package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS4 struct { F0 int32; F1 int8; F2 uint16 }
type GenS5 struct { F0 int; F1 int16; F2 int }
type GenS6 struct { F0 uint8; F1 string; F2 int8 }
type GenS7 struct { F0 int64; F1 int64 }
type GenS8 struct { F0 int; F1 string; F2 int16 }
type GenS9 struct { F0 bool }

func main() {
    var v1 int16 = 177
    var v2 int32 = 42
    v3 := []int{9223372036854775807, -486, -2147483648}
    v3 = append(v3, 989)
    for fi4 := 0; fi4 < 20; fi4++ {
        v2 = (v2 - 1)
        v1 -= (v1 | (v1 - (v1 >> ((0) & 63))))
        _ = fi4
    }
    if (!(false || (255 < 563))) {
        if ((v1 <= 1000) || (4294967295 == 18446744073709551615)) {
            fmt.Println(v2, v1)
        }
    }
    var v5 int8 = 1
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    var v6 int64 = 0
    if (v2 < v2) {
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch -673 {
            case 0:
                fmt.Println(v6, v1, v5)
            case 32767:
                v3 = append(v3, 1)
            default:
                v1 -= (v1 * v1)
            }
            v6 -= (v6 ^ v6)
        } else {
            var v7 uint8 = 3
            for fi8 := 0; fi8 < 2; fi8++ {
                for fi9 := 0; fi9 < 19; fi9++ {
                    if true {
                        if true {
                            v7 += v7
                        } else {
                            switch (fi9 + 1) {
                            case 419:
                                switch "" {
                                case "hello":
                                    var v10 int = 9223372036854775807
                                    _ = v10
                                case "a":
                                    if (1.5 != 1.5) {
                                        switch (^(fi8)) {
                                        case 127:
                                            v1 += (v1 >> ((551) & 63))
                                        case -785:
                                            if false {
                                                fmt.Println(v1, v7, v2)
                                                v11 := GenS1{F0: 139, F1: 0}
                                                if true {
                                                    if (!false) {
                                                        v12 := GenS2{F0: 0, F1: 42}
                                                        _ = v12
                                                    } else {
                                                        if len(v3) > 0 {
                                                            v3[0] = fi9
                                                        }
                                                        switch fi8 {
                                                        case -1:
                                                            var v13 float64 = (float64(1.5) / 1e10)
                                                            _ = v13
                                                        }
                                                    }
                                                    v2 = v2
                                                    v7 += v7
                                                }
                                                _ = v11
                                            }
                                        default:
                                            v14 := GenS1{F0: uint(v7), F1: 395}
                                            _ = v14
                                        }
                                    } else {
                                        v15 := GenS2{F0: v2, F1: 0}
                                        v5 = (v5 >> ((65535) & 63))
                                        _ = v15
                                    }
                                default:
                                    if (!("hello" >= "hello")) {
                                        if (!false) {
                                            var v16 float32 = -2.25
                                            var v17 int8 = v5
                                            v6 -= (v6 % (((v6 - (v6 << ((65535) & 63)))) | 1))
                                            _ = v16
                                            _ = v17
                                        }
                                        switch ("\n\t" + "hello") {
                                        case "":
                                            if false {
                                                v18 := GenS2{F0: -1, F1: 42}
                                                _ = v18
                                            } else {
                                                v6 = 42
                                            }
                                        case "hello":
                                            _ = len(v3)
                                            if false {
                                                v19 := GenS1{F0: 255, F1: 2}
                                                if ((-(fi8)) < fi8) {
                                                    var v20 string = "foo bar"
                                                    var v21 uint = 873
                                                    v20 = "a"
                                                    _ = v20
                                                    _ = v21
                                                }
                                                _ = v19
                                            } else {
                                                var v22 bool = false
                                                switch "\n\t" {
                                                case "a":
                                                    var v23 uint64 = 42
                                                    _ = v23
                                                case "":
                                                    var v24 uint16 = uint16(v7)
                                                    _ = v24
                                                default:
                                                    fmt.Println(v1, v7)
                                                }
                                                _ = v22
                                            }
                                            var v25 uint16 = 32767
                                            _ = v25
                                        case "a":
                                            v2 = v2
                                        default:
                                            if true {
                                                v7 = v7
                                            } else {
                                                var v26 uint64 = 18446744073709551615
                                                v27 := GenS2{F0: (v2 ^ (v2 | v2)), F1: v26}
                                                var v28 uint32 = 2147483648
                                                _ = v26
                                                _ = v27
                                                _ = v28
                                            }
                                        }
                                    }
                                }
                            case 4294967295:
                                if len(v3) > 0 {
                                    v3[0] = fi9
                                }
                            default:
                                _ = len(v3)
                            }
                        }
                        fmt.Println(v7, v6)
                    }
                    _ = fi9
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                for fi29 := 0; fi29 < 9; fi29++ {
                    v30 := GenS3{F0: 42, F1: 400, F2: 893}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v2 = (v2 + (v2 | v2))
                    _ = fi29
                    _ = v30
                }
                _ = fi8
            }
            _ = v7
        }
        fmt.Println(v6, v2)
    }
    v31 := GenS4{F0: -571, F1: 127, F2: 310}
    for fi32 := 0; fi32 < 9; fi32++ {
        for fi33 := 0; fi33 < 20; fi33++ {
            if (!(1 != 2)) {
                v5 = (v5 * v5)
                switch 0 {
                case 32768:
                    switch (fi32 | fi33) {
                    case 1:
                        v6 = v6
                    default:
                        var v34 int = (fi33 * (fi33 & -2147483648))
                        _ = v34
                    }
                case 1:
                    v6 += (^(v6))
                }
            }
            if (1 != 4294967295) {
                if true {
                    switch 2147483647 {
                    case -815:
                        v2 = (v2 << ((0) & 63))
                    case 127:
                        v2 = 42
                    case -1:
                        if false {
                            var v35 float64 = (3.14159 - 0.1)
                            switch "a" {
                            case "foo bar":
                                v3 = append(v3, (fi33 >> ((4294967295) & 63)))
                            case "a":
                                v35 = float64(3.14159)
                            case "hello":
                                v36 := GenS2{F0: 0, F1: 361}
                                _ = v36
                            }
                            if (("\n\t" + "foo bar") == "a") {
                                var v37 uint16 = 42
                                var v38 int = fi33
                                _ = v37
                                _ = v38
                            } else {
                                v1 = -1
                                v6 += (v6 + 42)
                            }
                            _ = v35
                        } else {
                            var v39 uint = 177
                            var v40 uint32 = 2147483648
                            v5 = (v5 << (((^(v39))) & 63))
                            _ = v39
                            _ = v40
                        }
                    default:
                        switch (("hello" + (("hello" + ("a" + "")) + "foo bar")) + (("" + "a") + "hello")) {
                        case "foo bar":
                            switch (^(fi33)) {
                            case 1:
                                var v41 int = int(v2)
                                _ = v41
                            case -504:
                                if false {
                                    v1 -= 1
                                    switch (fi33 - (fi33 + fi33)) {
                                    case -9223372036854775808:
                                        if len(v3) > 0 {
                                            v3[0] = fi33
                                        }
                                    case -451:
                                        var v42 string = ""
                                        _ = v42
                                    case -1:
                                        switch ("hello" + ("\n\t" + "\n\t")) {
                                        case "hello":
                                            v2 += v2
                                        }
                                    }
                                    v43 := GenS5{F0: (^(fi32)), F1: v1, F2: (-(fi33))}
                                    _ = v43
                                }
                            case 42:
                                v1 = (v1 - (v1 >> ((255) & 63)))
                            default:
                                switch fi33 {
                                case 1:
                                    v2 = -1
                                    var v44 float64 = float64(0.1)
                                    v45 := GenS1{F0: 320, F1: 4294967295}
                                    _ = v44
                                    _ = v45
                                case 0:
                                    v46 := GenS6{F0: 83, F1: "foo bar", F2: v5}
                                    _ = v46
                                default:
                                    var v47 int16 = 42
                                    _ = v47
                                }
                                var v48 int8 = (v5 >> ((1) & 63))
                                if (v5 < v48) {
                                    var v49 uint32 = 42
                                    _ = v49
                                } else {
                                    v50 := GenS2{F0: v2, F1: 4294967295}
                                    if (128 != 255) {
                                        v1 = (v1 << ((2) & 63))
                                        var v51 uint64 = 2
                                        switch "" {
                                        case "":
                                            v48 -= (v48 >> ((168) & 63))
                                        case "a":
                                            fmt.Println(v51, v1)
                                        case "foo bar":
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        _ = v51
                                    }
                                    _ = v50
                                }
                                _ = v48
                            }
                        }
                    }
                    var v52 int = fi33
                    if ((3.14159 - float32(1.5)) != -2.25) {
                        var v53 uint = 2
                        fmt.Println(v5)
                        var v54 uint64 = 18446744073709551615
                        _ = v53
                        _ = v54
                    } else {
                        var v55 uint16 = 1
                        v55 -= v55
                        var v56 uint32 = 2
                        _ = v55
                        _ = v56
                    }
                    _ = v52
                } else {
                    switch (fi32 << ((65535) & 63)) {
                    case 128:
                        v1 = v1
                    }
                    if (!(false || (!(4294967295 <= 549)))) {
                        fmt.Println(fi33, fi32)
                        var v57 uint16 = 518
                        v58 := GenS7{F0: 954, F1: v6}
                        _ = v57
                        _ = v58
                    } else {
                        switch (fi33 + fi33) {
                        case 1:
                            v1 += 234
                        case -114:
                            switch ("hello" + "hello") {
                            case "":
                                v6 = (v6 * v6)
                            case "\n\t":
                                v3 = append(v3, (fi33 | -124))
                            case "foo bar":
                                var v59 string = "\n\t"
                                _ = v59
                            default:
                                if len(v3) > 0 {
                                    v3[0] = (^(fi32))
                                }
                            }
                        }
                        if len(v3) > 0 {
                            v3[0] = fi32
                        }
                        v60 := GenS8{F0: fi32, F1: "hello", F2: v1}
                        _ = v60
                    }
                }
            } else {
                v61 := GenS9{F0: (255 == 0)}
                _ = v61
            }
            _ = fi33
        }
        _ = fi32
    }
    var v62 int32 = v2
    v62 -= (v62 / ((v62) | 1))
    switch ("a" + "foo bar") {
    case "":
        var v63 uint16 = 2
        _ = v63
    }
    if ((false && (65535 == 453)) || (42 <= 193)) {
        for fi64 := 0; fi64 < 10; fi64++ {
            v1 -= (v1 * -1)
            for fi65 := 0; fi65 < 14; fi65++ {
                fmt.Println(v1, v2)
                _ = fi65
            }
            _ = fi64
        }
    }
    var v66 uint64 = 0
    if ((3.14159 + 1e10) < 0.0) {
        _ = len(v3)
    }
    var v67 uint8 = 227
    var v68 int16 = v1
    for fi69 := 0; fi69 < 3; fi69++ {
        if true {
            var v70 bool = true
            var v71 int = (fi69 ^ fi69)
            _ = v70
            _ = v71
        } else {
            for fi72 := 0; fi72 < 1; fi72++ {
                var v73 uint16 = 42
                fmt.Println(v73, v2)
                _ = fi72
                _ = v73
            }
            v1 -= (v1 ^ (v1 % ((v1) | 1)))
            var v74 int = 0
            _ = v74
        }
        var v75 float32 = 0.1
        _ = fi69
        _ = v75
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v5
    _ = v6
    _ = v62
    _ = v66
    _ = v67
    _ = v68
}
