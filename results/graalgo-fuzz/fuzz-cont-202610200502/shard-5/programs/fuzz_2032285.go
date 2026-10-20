package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16; F1 uint; F2 int }
type GenS6 struct { F0 int; F1 bool; F2 int16 }
type GenS7 struct { F0 bool; F1 int }
type GenS8 struct { F0 string }

func main() {
    var v1 uint16 = 0
    v2 := make(map[string]int)
    var v3 uint64 = 0
    var v4 uint32 = 852
    var v5 uint32 = 2
    switch ("foo bar" + (("hello" + (("foo bar" + (("a" + "\n\t") + (("hello" + "foo bar") + ((("\n\t" + "") + "hello") + "hello")))) + (("a" + "") + "hello"))) + ("" + "hello"))) {
    case "hello":
        switch "hello" {
        case "\n\t":
            v6 := []int{int(v4), 32767, 467}
            _ = v6
        case "":
            var v7 float32 = -2.25
            _ = v7
        case "a":
            var v8 uint8 = 49
            _ = v8
        default:
            var v9 uint = 42
            _ = v9
        }
    default:
        v10 := []int{int(v3), -928, -694, 128}
        _ = v10
    }
    if (v1 > (^(v1))) {
        var v11 uint16 = (v1 & v1)
        _ = v11
    } else {
        var v12 float32 = (1.5 * 1.5)
        var v13 int = -281
        var v14 bool = false
        _ = v12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        fmt.Println(v1)
        _ = fi15
    }
    switch 127 {
    case 32768:
        var v16 uint32 = 65535
        _ = v16
    }
    v3 += v3
    for fi17 := 0; fi17 < 13; fi17++ {
        var v18 bool = true
        switch (fi17 / ((32768) | 1)) {
        case 482:
            fmt.Println(fi17, v18)
        case 9223372036854775807:
            for fi19 := 0; fi19 < 1; fi19++ {
                fmt.Println(v18, fi19)
                if v18 {
                    fmt.Println(v1, v3)
                    if v18 {
                        v20 := GenS1{F0: -113}
                        var v21 float32 = float32(float64((float32((1e10 + 0.1)) / 1.5)))
                        var v22 uint16 = (v1 << ((65535) & 63))
                        v2[""] = (^(fi17))
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    if ((!v18) || (-1 < 958)) {
                        var v23 uint64 = (v3 >> ((524) & 63))
                        fmt.Println(v1, fi17)
                        v24 := []int{(fi17 / ((1) | 1))}
                        _ = v23
                        _ = v24
                    }
                }
                _ = fi19
            }
        case -126:
            switch "hello" {
            case "\n\t":
                if v18 {
                    switch fi17 {
                    case 127:
                        var v25 uint = 107
                        _ = v25
                    case 0:
                        if (0 > 100) {
                            switch ("\n\t" + "") {
                            case "hello":
                                switch 507 {
                                case 0:
                                    var v26 string = "hello"
                                    _ = v26
                                case 525:
                                    var v27 int = -400
                                    _ = v27
                                case 4294967295:
                                    v5 -= 633
                                default:
                                    if ((3.14159 - -2.25) != 3.14159) {
                                        v28 := GenS1{F0: 42}
                                        v4 += 42
                                        switch (^(fi17)) {
                                        case -1:
                                            var v29 float32 = float32(3.14159)
                                            _ = v29
                                        default:
                                            v3 = (v3 << ((4294967295) & 63))
                                        }
                                        _ = v28
                                    }
                                }
                            }
                            var v30 uint64 = (v3 >> ((42) & 63))
                            for fi31 := 0; fi31 < 14; fi31++ {
                                var v32 uint = 1
                                _ = fi31
                                _ = v32
                            }
                            _ = v30
                        } else {
                            for fi33 := 0; fi33 < 14; fi33++ {
                                fmt.Println(fi17)
                                if v18 {
                                    var v34 uint = 18446744073709551615
                                    _ = v34
                                } else {
                                    switch fi17 {
                                    case 606:
                                        v3 = v3
                                    default:
                                        fmt.Println(v1, v5, v1)
                                    }
                                    v35 := GenS2{F0: 0, F1: (v18 || (!v18)), F2: v1}
                                    if ("\n\t" > "\n\t") {
                                        fmt.Println(v18, v3, v5)
                                    }
                                    _ = v35
                                }
                                _ = fi33
                            }
                            var v36 uint64 = (v3 | (v3 & (v3 & 897)))
                            _ = v36
                        }
                    default:
                        if false {
                            if (false && (0 >= 960)) {
                                v1 = v1
                            } else {
                                delete(v2, "foo bar")
                                var v37 uint8 = 128
                                var v38 int16 = -822
                                _ = v37
                                _ = v38
                            }
                            v39 := []int{-9223372036854775808, fi17, -632, 32767}
                            v2[(("hello" + "") + ("a" + (("hello" + "foo bar") + "a")))] = fi17
                            _ = v39
                        }
                    }
                } else {
                    v2["\n\t"] = fi17
                    var v40 uint8 = 127
                    _ = v40
                }
            default:
                for fi41 := 0; fi41 < 16; fi41++ {
                    switch fi17 {
                    case 1:
                        v42 := GenS3{F0: -147, F1: 42}
                        fmt.Println(fi41)
                        v43 := GenS1{F0: 1}
                        v1 += (v1 + (v1 + (v1 * v1)))
                        v4 += (^(v4))
                        v18 = (!v18)
                        _ = v42
                        _ = v43
                    }
                    _ = fi41
                }
            }
        }
        _ = fi17
        _ = v18
    }
    if true {
        switch 32768 {
        case 32768:
            v44 := GenS3{F0: 42, F1: (^(v3))}
            _ = v44
        case 906:
            for fi45 := 0; fi45 < 6; fi45++ {
                if (!true) {
                    var v46 int32 = -1
                    if (!false) {
                        for fi47 := 0; fi47 < 18; fi47++ {
                            v48 := GenS2{F0: 128, F1: true, F2: (^(v1))}
                            fmt.Println(v46, v3)
                            v49 := []int{}
                            _ = fi47
                            _ = v48
                            _ = v49
                        }
                        v50 := GenS3{F0: -535, F1: uint64(v5)}
                        _ = v50
                    } else {
                        if true {
                            v1 += (^(v1))
                            for fi51 := 0; fi51 < 1; fi51++ {
                                if ((false || false) && (false || true)) {
                                    v52 := []int{}
                                    fmt.Println(v4, v4)
                                    _ = v52
                                }
                                _ = fi51
                            }
                            v3 = (v3 * 4294967295)
                        }
                    }
                    _ = v46
                } else {
                    switch "foo bar" {
                    case "a":
                        v2["\n\t"] = fi45
                    case "":
                        if (!false) {
                            v53 := GenS4{F0: (!((^(v4)) > (^(v5))))}
                            for fi54 := 0; fi54 < 5; fi54++ {
                                var v55 int16 = 42
                                _ = fi54
                                _ = v55
                            }
                            _ = v53
                        }
                    case "hello":
                        switch fi45 {
                        case 247:
                            var v56 uint32 = v5
                            _ = v56
                        case 296:
                            v5 = v5
                        case 105:
                            for fi57 := 0; fi57 < 12; fi57++ {
                                switch "\n\t" {
                                case "a":
                                    fmt.Println(v1, fi45)
                                case "\n\t":
                                    _ = len(v2)
                                default:
                                    var v58 uint32 = (v5 >> ((545) & 63))
                                    _ = v58
                                }
                                _ = fi57
                            }
                        }
                    default:
                        v4 += (v5 / (((v5 + v5)) | 1))
                    }
                }
                v59 := GenS3{F0: 14, F1: (v3 - v3)}
                var v60 uint32 = (v4 - 0)
                _ = fi45
                _ = v59
                _ = v60
            }
        }
        v61 := GenS5{F0: int16(v5), F1: 65535, F2: 799}
        _ = v61
    }
    if false {
        if false {
            delete(v2, ("hello" + ((("foo bar" + "foo bar") + "a") + ("hello" + ("\n\t" + "hello")))))
            v5 = (v4 & (v5 % ((v5) | 1)))
            for fi62 := 0; fi62 < 13; fi62++ {
                if (!(!false)) {
                    v5 = 65535
                }
                if true {
                    var v63 bool = (("hello" + ("hello" + "a")) > "")
                    _ = v63
                }
                v64 := v2["hello"]
                _ = fi62
                _ = v64
            }
        }
        for fi65 := 0; fi65 < 7; fi65++ {
            v3 -= 1
            v4 += (^(v4))
            _ = fi65
        }
    }
    if (0 <= 967) {
        if true {
            var v66 int = -9223372036854775808
            _ = v66
        } else {
            v67 := GenS6{F0: -246, F1: false, F2: 557}
            _ = v67
        }
    }
    var v68 bool = ((v1 >> ((18446744073709551615) & 63)) <= 1)
    switch 2147483647 {
    case 1:
        for fi69 := 0; fi69 < 20; fi69++ {
            var v70 float64 = (float64(0.0) / 1.0)
            v68 = (1 <= 203)
            _ = fi69
            _ = v70
        }
    case 924:
        fmt.Println(v68)
    }
    switch (("\n\t" + (("a" + "hello") + "")) + "a") {
    case "hello":
        switch "\n\t" {
        case "hello":
            v71 := GenS4{F0: v68}
            _ = v71
        case "foo bar":
            v3 -= (v3 * 18446744073709551615)
        case "":
            switch "a" {
            case "hello":
                v68 = (!(v68 && true))
            case "a":
                var v72 float32 = 3.14159
                _ = v72
            default:
                var v73 bool = v68
                _ = v73
            }
        default:
            v74 := GenS7{F0: false, F1: 128}
            _ = v74
        }
    }
    for fi75 := 0; fi75 < 8; fi75++ {
        for fi76 := 0; fi76 < 7; fi76++ {
            if (65535 == 445) {
                if (!false) {
                    if ((fi76 ^ (fi76 * fi75)) > 99) {
                        switch (fi76 % ((fi75) | 1)) {
                        case 602:
                            v4 = v4
                            fmt.Println(v1, v3)
                        case 42:
                            if v68 {
                                v77 := []int{}
                                v78 := GenS2{F0: 141, F1: (0 != 1), F2: 187}
                                _ = len(v77)
                                if len(v77) > 0 {
                                    v77[0] = 127
                                }
                                v68 = true
                                _ = v77
                                _ = v78
                            }
                        }
                        if (v68 && (956 > -9223372036854775808)) {
                            fmt.Println(v3)
                        }
                        v5 = (v4 * v5)
                    }
                    v3 = uint64(v5)
                    var v79 float32 = (3.14159 / (0.1 / -2.25))
                    _ = v79
                } else {
                    var v80 int32 = -885
                    if (2147483647 == 2) {
                        var v81 float64 = float64(0.1)
                        _ = v81
                    } else {
                        if v68 {
                            v68 = (!false)
                            var v82 int = fi75
                            _ = v82
                        }
                    }
                    v4 -= (v4 + v5)
                    _ = v80
                }
                switch "\n\t" {
                case "hello":
                    _ = len(v2)
                }
                v68 = v68
                if v68 {
                    fmt.Println(v5, v3, fi76)
                    var v83 uint = 158
                    _ = v83
                } else {
                    fmt.Println(fi76, fi75, v3)
                }
            }
            v84 := GenS8{F0: ""}
            var v85 int32 = 1
            _ = fi76
            _ = v84
            _ = v85
        }
        v5 += v5
        _ = fi75
    }
    v4 += (v5 - (v4 + (v5 | v4)))
    if (float32(0.0) > 3.14159) {
        for fi86 := 0; fi86 < 16; fi86++ {
            var v87 string = (((("\n\t" + "foo bar") + "a") + "") + "hello")
            v5 = (v5 - v4)
            _ = fi86
            _ = v87
        }
        switch -992 {
        case -2147483648:
            v3 += v3
        case 42:
            fmt.Println(v4)
        }
        v1 = (v1 + v1)
    }
    fmt.Println(v1, v1, v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v68
}
