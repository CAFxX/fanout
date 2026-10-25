package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int64; F1 int; F2 uint32 }
type GenS7 struct { F0 bool; F1 uint64 }

func main() {
    var v1 int16 = -25
    var v2 float64 = (((1e10 + (1e10 * 3.14159)) * 1.5) * (float64(1e10) + ((1.5 * 1e10) + (3.14159 * (0.1 + (1.5 / 3.14159))))))
    v3 := make(map[int]string)
    v4 := []uint32{600, 0, 2147483647, 4294967295}
    if (42 > 18446744073709551615) {
        var v5 int8 = -40
        var v6 string = "a"
        fmt.Println(v2, v6)
        _ = v5
        _ = v6
    }
    fmt.Println(v1)
    if true {
        v2 = 1.5
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 int16 = 726
            fmt.Println(v8, v1, v2)
            if (238 < 128) {
                switch (fi7 % (((fi7 >> ((200) & 63))) | 1)) {
                case 2147483647:
                    for fi9 := 0; fi9 < 19; fi9++ {
                        v1 += 42
                        fmt.Println(v8, v2, v2)
                        switch "hello" {
                        case "":
                            fmt.Println(v8)
                        case "foo bar":
                            v10 := GenS1{F0: "", F1: -863}
                            _ = v10
                        case "a":
                            v2 += v2
                        default:
                            v11 := GenS1{F0: (("foo bar" + "foo bar") + (("a" + ("\n\t" + ("hello" + "\n\t"))) + "hello")), F1: int64(v1)}
                            _ = v11
                        }
                        _ = fi9
                    }
                case 602:
                    _ = len(v4)
                case 9223372036854775807:
                    switch (-(fi7)) {
                    case 32767:
                        v4 = append(v4, 2147483647)
                    case -1:
                        for fi12 := 0; fi12 < 2; fi12++ {
                            v13 := GenS2{F0: 0, F1: 0}
                            if true {
                                v14 := GenS3{F0: 65535, F1: 0}
                                switch ("\n\t" + ("a" + "\n\t")) {
                                case "\n\t":
                                    switch "\n\t" {
                                    case "a":
                                        v2 = (v2 - (1.5 / 3.14159))
                                    default:
                                        var v15 uint64 = 0
                                        _ = v15
                                    }
                                    var v16 uint8 = 255
                                    v2 = (1.5 + v2)
                                    _ = v16
                                case "hello":
                                    var v17 uint32 = 2
                                    _ = v17
                                default:
                                    fmt.Println(v8)
                                }
                                v1 += 42
                                switch (("hello" + "\n\t") + "") {
                                case "hello":
                                    var v18 float32 = 0.1
                                    _ = v18
                                }
                                v8 -= (v1 / ((v8) | 1))
                                _ = v14
                            } else {
                                v1 += v8
                                v4 = append(v4, 2147483648)
                                if (-128 < -1) {
                                    v19 := GenS1{F0: "a", F1: 0}
                                    fmt.Println(fi12)
                                    _ = v19
                                } else {
                                    if false {
                                        var v20 int = 32768
                                        fmt.Println(fi7)
                                        _ = v20
                                    }
                                    var v21 float32 = (1.5 / float32(((v2 - v2) * 1e10)))
                                    _ = v21
                                }
                                if true {
                                    switch "hello" {
                                    case "hello":
                                        var v22 string = "foo bar"
                                        _ = v22
                                    }
                                    fmt.Println(fi12)
                                    fmt.Println(v1)
                                }
                                if false {
                                    if true {
                                        v8 = 1000
                                    }
                                    v8 -= (v1 >> ((42) & 63))
                                    if (18446744073709551615 == 4294967295) {
                                        if false {
                                            v1 = (-(v8))
                                            if false {
                                                _ = len(v4)
                                                fmt.Println(fi7, fi12)
                                            } else {
                                                v23 := GenS2{F0: 1, F1: 1}
                                                _ = len(v3)
                                                v24 := GenS4{F0: 255}
                                                var v25 int32 = 0
                                                _ = v23
                                                _ = v24
                                                _ = v25
                                            }
                                            v1 -= (v8 << ((255) & 63))
                                            var v26 uint64 = 2
                                            _ = v26
                                        } else {
                                            var v27 uint = 255
                                            v1 += -268
                                            _ = v27
                                        }
                                        v28 := GenS2{F0: 1, F1: 1}
                                        _ = v28
                                    } else {
                                        if ((694 <= -9223372036854775808) && (fi7 >= fi7)) {
                                            v1 -= -32768
                                            switch -397 {
                                            case 4294967295:
                                                var v29 uint8 = 250
                                                if ((true && true) || (42 >= 219)) {
                                                    if true {
                                                        var v30 bool = ((1 < 29) && false)
                                                        var v31 int8 = -128
                                                        _ = v30
                                                        _ = v31
                                                    }
                                                }
                                                v32 := GenS5{F0: false}
                                                if ((^(v29)) > 42) {
                                                    v1 = v8
                                                    if false {
                                                        fmt.Println(fi12, fi7)
                                                        fmt.Println(v1, fi7, fi7)
                                                    } else {
                                                        var v33 uint16 = 13
                                                        fmt.Println(v8, v33)
                                                        _ = v33
                                                    }
                                                    var v34 int32 = int32(fi7)
                                                    var v35 int64 = 173
                                                    var v36 int = (fi12 & fi12)
                                                    fmt.Println(v35, fi7, fi7)
                                                    _ = v34
                                                    _ = v35
                                                    _ = v36
                                                }
                                                delete(v3, 127)
                                                _ = v29
                                                _ = v32
                                            default:
                                                _ = len(v3)
                                            }
                                            v2 -= (v2 + (((0.0 - v2) / 1.5) * float64(1.5)))
                                        } else {
                                            var v37 float32 = (0.0 / -2.25)
                                            v38 := GenS2{F0: 100, F1: 65535}
                                            _ = v37
                                            _ = v38
                                        }
                                    }
                                }
                            }
                            switch (fi12 ^ fi12) {
                            case -1:
                                var v39 uint32 = 42
                                _ = v39
                            default:
                                var v40 int32 = 2147483647
                                _ = v40
                            }
                            _ = fi12
                            _ = v13
                        }
                    case 1:
                        fmt.Println(v2, v8)
                    }
                default:
                    if ((true || false) || true) {
                        var v41 int8 = -72
                        var v42 int8 = (v41 ^ v41)
                        _ = v41
                        _ = v42
                    } else {
                        delete(v3, 837)
                        for fi43 := 0; fi43 < 11; fi43++ {
                            v2 = v2
                            _ = fi43
                        }
                        for fi44 := 0; fi44 < 13; fi44++ {
                            switch fi44 {
                            case 0:
                                if (373 < 42) {
                                    var v45 uint64 = 1
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    _ = len(v4)
                                    _ = v45
                                } else {
                                    v46 := GenS5{F0: false}
                                    _ = v46
                                }
                            case 128:
                                var v47 string = ""
                                _ = v47
                            }
                            v8 -= v8
                            _ = fi44
                        }
                    }
                }
                var v48 uint = 751
                v49 := GenS4{F0: 127}
                _ = v48
                _ = v49
            } else {
                delete(v3, (fi7 + (fi7 - (fi7 ^ fi7))))
                var v50 int16 = v8
                _ = v50
            }
            _ = fi7
            _ = v8
        }
    }
    for fi51 := 0; fi51 < 10; fi51++ {
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
        switch 42 {
        case 32767:
            v2 += v2
        default:
            var v52 int8 = 1
            _ = v52
        }
        if (127 < 148) {
            switch "\n\t" {
            case "\n\t":
                v53 := GenS5{F0: false}
                _ = v53
            case "hello":
                _ = len(v4)
            default:
                var v54 float32 = 1.5
                _ = v54
            }
        } else {
            v1 -= (^(v1))
            if true {
                fmt.Println(fi51)
                switch "hello" {
                case "a":
                    v55 := GenS4{F0: 59}
                    _ = v55
                }
            }
        }
        _ = fi51
    }
    for fi56 := 0; fi56 < 9; fi56++ {
        var v57 int64 = 9223372036854775807
        _ = fi56
        _ = v57
    }
    v58 := v3[501]
    v58 = "a"
    var v59 uint64 = 18446744073709551615
    if false {
        var v60 int8 = -1
        var v61 uint64 = 4294967295
        if (true && (!false)) {
            for fi62 := 0; fi62 < 18; fi62++ {
                var v63 int8 = -19
                if (2 != v59) {
                    for fi64 := 0; fi64 < 4; fi64++ {
                        fmt.Println(v60, v63)
                        _ = fi64
                    }
                    var v65 int64 = -153
                    v66 := GenS5{F0: false}
                    _ = v65
                    _ = v66
                }
                _ = fi62
                _ = v63
            }
        } else {
            switch 128 {
            case 533:
                delete(v3, -755)
            case -34:
                for fi67 := 0; fi67 < 5; fi67++ {
                    var v68 uint = 65535
                    v4 = append(v4, 707)
                    _ = fi67
                    _ = v68
                }
            }
            var v69 string = ("hello" + v58)
            if true {
                v70 := GenS5{F0: false}
                var v71 int = -527
                if true {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    if (v71 <= v71) {
                        v59 += (^(v59))
                        for fi72 := 0; fi72 < 15; fi72++ {
                            v73 := GenS5{F0: true}
                            v74 := GenS2{F0: v60, F1: 42}
                            _ = fi72
                            _ = v73
                            _ = v74
                        }
                    }
                    if ((!(!true)) && ((v69 + (v58 + (v58 + "hello"))) == (((v58 + v69) + "hello") + v69))) {
                        v60 = v60
                        v61 = 18446744073709551615
                        var v75 uint16 = 0
                        _ = v75
                    }
                } else {
                    if (245 == 255) {
                        var v76 uint = uint(v59)
                        if (!(!(v71 == v71))) {
                            var v77 int16 = v1
                            _ = v77
                        } else {
                            v4 = append(v4, 65535)
                        }
                        _ = v76
                    } else {
                        _ = len(v4)
                        var v78 float32 = ((3.14159 + 3.14159) / ((1e10 / (float32((v2 * ((v2 + v2) * ((1e10 / ((1e10 / (float64((0.1 - float32((0.0 + 3.14159)))) * v2)) - v2)) - -2.25)))) + float32(v2))) + 0.1))
                        for fi79 := 0; fi79 < 18; fi79++ {
                            fmt.Println(v2, v78)
                            _ = fi79
                        }
                        _ = v78
                    }
                }
                _ = v70
                _ = v71
            }
            _ = v69
        }
        _ = v60
        _ = v61
    }
    switch (v58 + v58) {
    case "":
        switch v58 {
        case "a":
            delete(v3, -2147483648)
        case "\n\t":
            var v80 int = 501
            _ = v80
        case "":
            var v81 int16 = (v1 | v1)
            _ = v81
        default:
            if (v58 != v58) {
                v58 = ("hello" + v58)
                v1 = v1
                _ = len(v3)
            } else {
                v3[128] = ""
            }
        }
    default:
        if (true || (v1 > (-(v1)))) {
            v58 = "a"
            switch v58 {
            case "foo bar":
                fmt.Println(v2, v59, v2)
            case "a":
                fmt.Println(v2)
            default:
                if false {
                    for fi82 := 0; fi82 < 5; fi82++ {
                        if false {
                            var v83 uint8 = 42
                            fmt.Println(v58, v58, v1)
                            _ = v83
                        }
                        for fi84 := 0; fi84 < 17; fi84++ {
                            switch "\n\t" {
                            case "hello":
                                var v85 uint8 = 2
                                _ = v85
                            case "a":
                                v86 := GenS6{F0: -918, F1: 128, F2: 486}
                                _ = v86
                            }
                            switch ("" + (v58 + ((v58 + "\n\t") + v58))) {
                            case "hello":
                                switch (((v58 + (v58 + "a")) + v58) + "") {
                                case "foo bar":
                                    v59 = 0
                                case "\n\t":
                                    var v87 bool = (!true)
                                    _ = v87
                                }
                            case "foo bar":
                                _ = len(v4)
                            }
                            v88 := GenS1{F0: "\n\t", F1: 9223372036854775807}
                            _ = fi84
                            _ = v88
                        }
                        _ = fi82
                    }
                    for fi89 := 0; fi89 < 4; fi89++ {
                        fmt.Println(v1, v59)
                        _ = fi89
                    }
                } else {
                    v58 = v58
                }
            }
            fmt.Println(v58, v2)
        }
    }
    v90 := GenS7{F0: false, F1: (v59 >> ((0) & 63))}
    v58 = "foo bar"
    v58 = (v58 + (("foo bar" + "hello") + (((v58 + v58) + (v58 + (("hello" + v58) + "hello"))) + v58)))
    var v91 int32 = 1
    var v92 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v58
    _ = v59
    _ = v90
    _ = v91
    _ = v92
}
