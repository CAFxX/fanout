package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint; F1 uint32; F2 int16 }
type GenS8 struct { F0 int; F1 uint8 }
type GenS9 struct { F0 int16; F1 uint16 }

func main() {
    v1 := []int{}
    var v2 bool = (-9223372036854775808 <= 9223372036854775807)
    var v3 bool = v2
    var v4 int8 = 100
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 9223372036854775807 {
    case 127:
        var v5 int = 128
        _ = v5
    case -1:
        var v6 uint8 = 222
        _ = v6
    default:
        switch 42 {
        case 0:
            fmt.Println(v4, v4)
        case 32767:
            var v7 uint64 = 2
            _ = v7
        }
    }
    var v8 uint = 18446744073709551615
    switch ("" + "a") {
    case "hello":
        for fi9 := 0; fi9 < 8; fi9++ {
            if (!(v2 || v2)) {
                var v10 uint16 = 894
                switch ("a" + ("foo bar" + "a")) {
                case "hello":
                    v11 := GenS1{F0: 2147483647, F1: 130, F2: 894}
                    _ = v11
                }
                if v2 {
                    var v12 int = fi9
                    if (42 > 0) {
                        if v3 {
                            switch ("a" + (("hello" + "\n\t") + "foo bar")) {
                            case "":
                                v1 = append(v1, (v12 - (fi9 - v12)))
                            case "hello":
                                for fi13 := 0; fi13 < 8; fi13++ {
                                    if len(v1) > 0 {
                                        v1[0] = (fi9 >> ((v8) & 63))
                                    }
                                    _ = fi13
                                }
                            case "\n\t":
                                var v14 float32 = 0.1
                                _ = v14
                            default:
                                v15 := GenS2{F0: 0, F1: -32768}
                                _ = v15
                            }
                            v2 = true
                        }
                        var v16 uint16 = 255
                        _ = v16
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v12
                } else {
                    switch (fi9 | -1) {
                    case 110:
                        fmt.Println(v8, fi9)
                    case 313:
                        var v17 float64 = 0.0
                        _ = v17
                    }
                }
                _ = v10
            }
            if ((!((((0.0 / -2.25) + (3.14159 / (1e10 * (((-2.25 * 0.0) - (1.5 * (0.0 + 0.1))) + ((3.14159 + 0.1) + (1e10 / -2.25)))))) != 3.14159) && v3)) && false) {
                var v18 int8 = v4
                _ = v18
            } else {
                var v19 int64 = 1
                _ = v19
            }
            for fi20 := 0; fi20 < 19; fi20++ {
                v21 := GenS2{F0: uint16(v8), F1: 1}
                _ = fi20
                _ = v21
            }
            _ = fi9
        }
    case "foo bar":
        if false {
            for fi22 := 0; fi22 < 8; fi22++ {
                var v23 uint = 2
                _ = fi22
                _ = v23
            }
        }
    }
    if v2 {
        v24 := GenS3{F0: 0}
        switch "" {
        case "a":
            var v25 int16 = 1000
            _ = v25
        case "hello":
            switch ("" + "foo bar") {
            case "\n\t":
                if len(v1) > 0 {
                    v1[0] = 9223372036854775807
                }
            }
        case "":
            for fi26 := 0; fi26 < 3; fi26++ {
                v8 += (v8 - v8)
                v2 = v3
                fmt.Println(v8)
                _ = fi26
            }
        }
        v1 = append(v1, 4294967295)
        _ = v24
    }
    if v3 {
        v27 := GenS4{F0: 256}
        switch 384 {
        case 729:
            v28 := GenS3{F0: 160}
            _ = v28
        default:
            for fi29 := 0; fi29 < 5; fi29++ {
                fmt.Println(fi29, v8, v8)
                for fi30 := 0; fi30 < 17; fi30++ {
                    if true {
                        if v3 {
                            var v31 int = 9223372036854775807
                            switch "foo bar" {
                            case "\n\t":
                                v8 += 719
                            case "hello":
                                v8 = (v8 >> (((v8 >> (((v8 >> (((v8 % ((468) | 1))) & 63))) & 63))) & 63))
                            default:
                                if ((!v3) || true) {
                                    fmt.Println(v3, fi29, v2)
                                    switch "hello" {
                                    case "hello":
                                        var v32 int64 = -9223372036854775808
                                        _ = v32
                                    case "foo bar":
                                        fmt.Println(v31, fi29, fi29)
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = v31
                                    }
                                    if v3 {
                                        if true {
                                            fmt.Println(fi29, fi30)
                                        }
                                    }
                                } else {
                                    if false {
                                        var v33 float64 = 0.0
                                        var v34 int64 = 0
                                        fmt.Println(v3, v4)
                                        var v35 int32 = -2147483648
                                        fmt.Println(fi30)
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    } else {
                                        var v36 uint = 2
                                        v37 := GenS5{F0: 0, F1: 2147483647}
                                        _ = v36
                                        _ = v37
                                    }
                                    switch (v31 + fi29) {
                                    case -2147483648:
                                        fmt.Println(v4)
                                    case -9223372036854775808:
                                        var v38 bool = v2
                                        _ = v38
                                    }
                                }
                            }
                            if (-32768 != 32767) {
                                v4 = (v4 + v4)
                            } else {
                                var v39 uint16 = 42
                                _ = v39
                            }
                            v3 = (0 < 2)
                            _ = v31
                        }
                    } else {
                        var v40 uint8 = 155
                        v41 := GenS6{F0: 1000}
                        v2 = (v2 || ((!(v8 >= (v8 / (((v8 / ((v8) | 1))) | 1)))) && (("" + "\n\t") > "a")))
                        v42 := GenS5{F0: -1, F1: 486}
                        var v43 float32 = 3.14159
                        _ = v40
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                    switch ("" + "foo bar") {
                    case "hello":
                        if ((v8 & 4294967295) >= v8) {
                            switch fi30 {
                            case 128:
                                if (!((42 >= -1) || v3)) {
                                    var v44 int32 = -1
                                    var v45 int8 = 1
                                    _ = v44
                                    _ = v45
                                } else {
                                    v4 += (v4 * -1)
                                    if v2 {
                                        v3 = (335 <= 65535)
                                        v3 = v3
                                        fmt.Println(fi30, fi29, fi30)
                                    }
                                    fmt.Println(fi30)
                                }
                            case -136:
                                _ = len(v1)
                            }
                            _ = len(v1)
                        } else {
                            v8 = 65535
                            var v46 string = "foo bar"
                            _ = v46
                        }
                        switch (fi29 & (fi30 * 1)) {
                        case 32768:
                            if (1 > 128) {
                                v8 -= 83
                            } else {
                                v3 = v2
                                v47 := GenS7{F0: (^(v8)), F1: 1, F2: 32767}
                                switch ("\n\t" + (((((("hello" + "") + "") + "foo bar") + ("a" + "")) + "a") + "")) {
                                case "\n\t":
                                    v48 := GenS8{F0: (fi30 * 218), F1: 255}
                                    _ = v48
                                case "hello":
                                    switch fi29 {
                                    case 1:
                                        if v2 {
                                            var v49 int8 = (^(v4))
                                            _ = v49
                                        }
                                    }
                                default:
                                    v1 = append(v1, (^(fi30)))
                                }
                                var v50 uint64 = 4294967295
                                _ = v47
                                _ = v50
                            }
                        }
                    }
                    _ = fi30
                }
                _ = fi29
            }
        }
        v3 = ((-1 > -6) && false)
        _ = v27
    }
    switch 0 {
    case 782:
        v4 = v4
    case 127:
        v4 = v4
    default:
        v51 := GenS1{F0: -2147483648, F1: 0, F2: uint32(v8)}
        _ = v51
    }
    var v52 uint16 = 255
    if false {
        var v53 bool = (v2 && (1 >= 1))
        for fi54 := 0; fi54 < 7; fi54++ {
            v55 := GenS3{F0: 67}
            _ = fi54
            _ = v55
        }
        _ = v53
    } else {
        _ = len(v1)
        var v56 float32 = (0.1 * float32(-2.25))
        var v57 uint32 = 2
        _ = v56
        _ = v57
    }
    switch "\n\t" {
    case "hello":
        for fi58 := 0; fi58 < 20; fi58++ {
            v8 -= 4294967295
            var v59 int64 = 9223372036854775807
            var v60 uint32 = 445
            _ = fi58
            _ = v59
            _ = v60
        }
    case "foo bar":
        switch "\n\t" {
        case "":
            v61 := GenS9{F0: -32768, F1: (v52 - v52)}
            _ = v61
        }
    case "\n\t":
        for fi62 := 0; fi62 < 4; fi62++ {
            v63 := GenS9{F0: int16(v8), F1: 555}
            v4 += v4
            v4 = (v4 << ((2) & 63))
            _ = fi62
            _ = v63
        }
    default:
        v52 = v52
    }
    if true {
        v2 = true
        v3 = ((("" + "foo bar") + "") < ((("foo bar" + "a") + "a") + "a"))
        v3 = (v2 && true)
    }
    for fi64 := 0; fi64 < 3; fi64++ {
        var v65 float64 = 1e10
        switch "\n\t" {
        case "":
            if (4294967295 >= uint64(v8)) {
                fmt.Println(v8, v4)
                v3 = true
            }
        case "foo bar":
            fmt.Println(v3, v8, v3)
        default:
            fmt.Println(v4, v4)
        }
        _ = fi64
        _ = v65
    }
    v66 := GenS5{F0: 890, F1: 775}
    v52 = v52
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v66
    _ = v8
}
