package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 int8; F1 uint16 }
type GenS7 struct { F0 int32 }

func main() {
    var v1 float32 = ((float32(float64(float32((1e10 / (1.5 / 1.5))))) - -2.25) + 1.5)
    var v2 int = -2147483648
    var v3 string = "\n\t"
    v4 := []int{v2, -928, -9223372036854775808, v2}
    var v5 uint = 18446744073709551615
    var v6 uint8 = 255
    if false {
        if len(v4) > 0 {
            v4[0] = -905
        }
    }
    switch int(v5) {
    case -2147483648:
        var v7 float32 = v1
        _ = v7
    case 9223372036854775807:
        var v8 string = "foo bar"
        _ = v8
    default:
        switch ("foo bar" + v3) {
        case "a":
            for fi9 := 0; fi9 < 18; fi9++ {
                v10 := GenS1{F0: 981, F1: 0, F2: fi9}
                switch (v2 >> ((1) & 63)) {
                case -668:
                    if (!true) {
                        _ = len(v4)
                        var v11 float32 = ((0.0 + v1) * (1e10 - float32((((((0.0 * 0.0) * (((-2.25 + 0.0) - (0.1 + 1e10)) / 1.0)) - -2.25) / float64(0.1)) / 1.0))))
                        for fi12 := 0; fi12 < 10; fi12++ {
                            v5 = uint(v6)
                            v13 := GenS2{F0: v3, F1: "foo bar"}
                            _ = fi12
                            _ = v13
                        }
                        _ = v11
                    }
                default:
                    v1 = 0.0
                }
                if false {
                    _ = len(v4)
                    switch 1 {
                    case -2147483648:
                        switch "foo bar" {
                        case "\n\t":
                            if len(v4) > 0 {
                                v4[0] = (v2 - fi9)
                            }
                        case "foo bar":
                            if false {
                                v14 := GenS1{F0: 42, F1: 2, F2: v2}
                                _ = v14
                            }
                        default:
                            v15 := GenS1{F0: 0, F1: 62, F2: v2}
                            _ = v15
                        }
                    }
                    var v16 uint = 2
                    _ = v16
                } else {
                    v1 = v1
                }
                _ = fi9
                _ = v10
            }
        case "foo bar":
            v17 := GenS3{F0: v3}
            _ = v17
        }
    }
    var v18 uint64 = 998
    switch v3 {
    case "":
        v6 = v6
    case "a":
        if (207 <= -535) {
            var v19 uint = (v5 + v5)
            if false {
                var v20 int64 = 42
                if false {
                    switch "a" {
                    case "":
                        if (-1 >= 753) {
                            fmt.Println(v5, v20, v1)
                            if (2147483647 < -1) {
                                _ = len(v4)
                                if ((v5 * v5) <= (v5 ^ 242)) {
                                    var v21 int64 = (v20 % ((v20) | 1))
                                    switch v3 {
                                    case "foo bar":
                                        if (!(v20 != (^(v20)))) {
                                            var v22 bool = true
                                            var v23 string = (v3 + ("\n\t" + "\n\t"))
                                            _ = v22
                                            _ = v23
                                        } else {
                                            var v24 int = (^(v2))
                                            v20 -= (v21 ^ -9223372036854775808)
                                            var v25 float32 = v1
                                            _ = v24
                                            _ = v25
                                        }
                                    case "\n\t":
                                        fmt.Println(v6, v20)
                                    case "hello":
                                        if (!true) {
                                            for fi26 := 0; fi26 < 11; fi26++ {
                                                fmt.Println(v5, v3, v1)
                                                fmt.Println(v20, v1)
                                                fmt.Println(v2, v5, v1)
                                                _ = fi26
                                            }
                                            fmt.Println(v18)
                                        }
                                    default:
                                        fmt.Println(v3, v20, v2)
                                    }
                                    _ = v21
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    for fi27 := 0; fi27 < 13; fi27++ {
                                        v28 := GenS4{F0: 0, F1: 65535}
                                        _ = fi27
                                        _ = v28
                                    }
                                }
                            } else {
                                var v29 uint32 = 4294967295
                                switch (v2 << ((111) & 63)) {
                                case 9223372036854775807:
                                    v30 := GenS5{F0: v29, F1: 42}
                                    _ = v30
                                case 32768:
                                    if (!(v1 >= 3.14159)) {
                                        for fi31 := 0; fi31 < 15; fi31++ {
                                            var v32 int16 = 800
                                            _ = fi31
                                            _ = v32
                                        }
                                    }
                                case 4294967295:
                                    v3 = "foo bar"
                                }
                                _ = v29
                            }
                        }
                    default:
                        v6 = (v6 % ((42) | 1))
                    }
                    fmt.Println(v19)
                    v33 := GenS6{F0: 125, F1: 32767}
                    _ = v33
                } else {
                    fmt.Println(v19, v19)
                    v34 := GenS7{F0: 1}
                    _ = v34
                }
                _ = v20
            }
            fmt.Println(v19, v3, v19)
            _ = v19
        }
    case "\n\t":
        for fi35 := 0; fi35 < 3; fi35++ {
            v1 = v1
            for fi36 := 0; fi36 < 1; fi36++ {
                switch v3 {
                case "":
                    if (-19 != 0) {
                        var v37 uint = 0
                        if len(v4) > 0 {
                            v4[0] = -9223372036854775808
                        }
                        var v38 int16 = 0
                        _ = v37
                        _ = v38
                    } else {
                        var v39 uint8 = 127
                        fmt.Println(fi35, v6, v18)
                        _ = v39
                    }
                default:
                    if ((false || (3.14159 >= 1.5)) && (!true)) {
                        v1 += 0.1
                        fmt.Println(v3)
                    } else {
                        var v40 bool = (-1 > 42)
                        _ = v40
                    }
                }
                _ = fi36
            }
            _ = fi35
        }
    default:
        for fi41 := 0; fi41 < 6; fi41++ {
            if (-1 <= -50) {
                var v42 string = (v3 + v3)
                _ = v42
            }
            if false {
                var v43 uint = 73
                v3 = ("\n\t" + (v3 + v3))
                v1 = v1
                _ = v43
            } else {
                v18 -= (v18 - 4294967295)
            }
            _ = fi41
        }
    }
    for fi44 := 0; fi44 < 2; fi44++ {
        switch "\n\t" {
        case "a":
            v4 = append(v4, (fi44 << ((220) & 63)))
        case "":
            for fi45 := 0; fi45 < 15; fi45++ {
                var v46 bool = true
                var v47 uint32 = 740
                _ = fi45
                _ = v46
                _ = v47
            }
        default:
            fmt.Println(v1, v3)
        }
        _ = len(v4)
        _ = fi44
    }
    for fi48 := 0; fi48 < 20; fi48++ {
        var v49 string = "\n\t"
        for fi50 := 0; fi50 < 8; fi50++ {
            v51 := GenS1{F0: v18, F1: 255, F2: fi50}
            fmt.Println(v49, fi48, fi48)
            _ = fi50
            _ = v51
        }
        if true {
            var v52 uint32 = uint32(v6)
            var v53 int = (fi48 / (((v2 ^ v2)) | 1))
            _ = v52
            _ = v53
        }
        _ = fi48
        _ = v49
    }
    for fi54 := 0; fi54 < 7; fi54++ {
        if (true && (!(!true))) {
            var v55 int32 = 42
            switch 351 {
            case 2147483647:
                var v56 uint16 = 32767
                _ = v56
            }
            _ = v55
        } else {
            v1 = -2.25
            v57 := GenS1{F0: (v18 / ((v18) | 1)), F1: 1, F2: 0}
            _ = v57
        }
        v58 := GenS3{F0: (v3 + "foo bar")}
        switch "hello" {
        case "a":
            switch ("" + "") {
            case "\n\t":
                v59 := GenS2{F0: ("hello" + "\n\t"), F1: (v3 + ((v3 + "foo bar") + "\n\t"))}
                _ = v59
            case "":
                v6 = v6
            case "a":
                var v60 float64 = float64((v1 - ((((float32(0.1) / (v1 * v1)) + 0.0) * v1) / v1)))
                _ = v60
            default:
                switch (v3 + v3) {
                case "":
                    switch (fi54 % ((fi54) | 1)) {
                    case 128:
                        v5 += (v5 / (((^(v5))) | 1))
                    case 4294967295:
                        v61 := GenS3{F0: v3}
                        _ = v61
                    case 500:
                        if false {
                            if true {
                                v3 = "a"
                                v5 = v5
                                for fi62 := 0; fi62 < 2; fi62++ {
                                    var v63 string = "hello"
                                    _ = fi62
                                    _ = v63
                                }
                            }
                        } else {
                            if (false && true) {
                                for fi64 := 0; fi64 < 1; fi64++ {
                                    if ((fi54 % ((42) | 1)) == (^(v2))) {
                                        v6 += uint8(v18)
                                    }
                                    _ = fi64
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            } else {
                                v2 -= v2
                            }
                            var v65 int64 = -1
                            _ = v65
                        }
                    default:
                        v5 = (v5 << (((^(v5))) & 63))
                    }
                case "\n\t":
                    if (0 < -1) {
                        var v66 float32 = (v1 / v1)
                        var v67 string = ((v3 + "hello") + v3)
                        _ = v66
                        _ = v67
                    } else {
                        if true {
                            if (true && true) {
                                var v68 float32 = float32((((1e10 / 1.0) / -2.25) + -2.25))
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                for fi69 := 0; fi69 < 19; fi69++ {
                                    v4 = append(v4, (v2 - v2))
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v1 = -2.25
                                    _ = fi69
                                }
                                _ = v68
                            } else {
                                v5 -= (v5 + (v5 % ((v5) | 1)))
                            }
                            if false {
                                var v70 uint = v5
                                _ = v70
                            }
                        }
                        for fi71 := 0; fi71 < 14; fi71++ {
                            fmt.Println(v5, v3, v6)
                            _ = len(v4)
                            if len(v4) > 0 {
                                v4[0] = -9223372036854775808
                            }
                            _ = fi71
                        }
                        switch v3 {
                        case "foo bar":
                            _ = len(v4)
                        case "a":
                            v4 = append(v4, v2)
                        case "":
                            if false {
                                for fi72 := 0; fi72 < 5; fi72++ {
                                    v73 := GenS4{F0: 42, F1: 2}
                                    _ = fi72
                                    _ = v73
                                }
                                v5 -= (^(v5))
                                v4 = append(v4, 127)
                            } else {
                                switch v3 {
                                case "\n\t":
                                    fmt.Println(fi54, fi54)
                                case "":
                                    v2 = (-(v2))
                                case "hello":
                                    var v74 float64 = 1.5
                                    _ = v74
                                }
                            }
                        }
                    }
                }
            }
        case "foo bar":
            if false {
                if (v3 > v3) {
                    v5 = 65535
                    fmt.Println(v6, v1)
                }
                v18 -= 2
            } else {
                if (!true) {
                    fmt.Println(fi54, v1, v5)
                    fmt.Println(fi54, v6, v6)
                    for fi75 := 0; fi75 < 5; fi75++ {
                        v3 = v3
                        var v76 int64 = -833
                        if len(v4) > 0 {
                            v4[0] = v2
                        }
                        _ = fi75
                        _ = v76
                    }
                }
            }
        default:
            var v77 uint32 = 0
            _ = v77
        }
        _ = fi54
        _ = v58
    }
    fmt.Println(v1)
    var v78 int16 = 1
    var v79 int8 = -1
    if len(v4) > 0 {
        v4[0] = v2
    }
    switch "\n\t" {
    case "a":
        v80 := GenS2{F0: v3, F1: "hello"}
        _ = v80
    case "hello":
        v5 += v5
    default:
        for fi81 := 0; fi81 < 18; fi81++ {
            var v82 bool = true
            _ = fi81
            _ = v82
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v78
    _ = v79
}
