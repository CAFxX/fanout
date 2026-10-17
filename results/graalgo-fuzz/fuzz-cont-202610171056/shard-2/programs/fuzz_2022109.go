package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 uint; F1 bool; F2 int16 }

func main() {
    v1 := []int{}
    v2 := []int{32768, 1, 127, 918}
    v3 := make(map[int]uint32)
    var v4 uint = 65535
    var v5 uint = 65535
    fmt.Println(v5, v5)
    fmt.Println(v5, v4)
    v6 := GenS1{F0: 345}
    v5 += (^(v5))
    if (false && (!true)) {
        if true {
            v7 := GenS1{F0: 2}
            var v8 int = 0
            v4 = (v4 << ((1) & 63))
            _ = v7
            _ = v8
        }
        fmt.Println(v5, v5)
    }
    var v9 uint = (v4 | (v5 - (^(v4))))
    for fi10 := 0; fi10 < 8; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            if true {
                switch (fi11 / ((fi10) | 1)) {
                case 32767:
                    if false {
                        if false {
                            v12 := GenS1{F0: 4294967295}
                            _ = v12
                        }
                        switch 1 {
                        case -637:
                            v5 += 42
                        }
                        switch ((("\n\t" + "\n\t") + "\n\t") + "") {
                        case "\n\t":
                            v5 -= (v4 + (v4 / (((v5 / (((v4 % ((v9) | 1))) | 1))) | 1)))
                        case "":
                            var v13 uint8 = 248
                            _ = v13
                        case "hello":
                            var v14 uint64 = 1
                            _ = v14
                        default:
                            v2 = append(v2, (-(fi11)))
                        }
                    }
                }
                delete(v3, 163)
            } else {
                fmt.Println(v5)
                var v15 int32 = 805
                _ = v15
            }
            var v16 uint32 = 0
            var v17 int8 = 0
            if ((v4 | (v4 ^ 849)) >= 5) {
                var v18 uint16 = 1
                if false {
                    v19 := GenS2{F0: (v17 & (v17 % (((v17 % (((v17 | v17)) | 1))) | 1))), F1: int16(v16), F2: 2}
                    v20 := GenS3{F0: 2, F1: (false && true)}
                    _ = v19
                    _ = v20
                }
                _ = v18
            }
            var v21 uint8 = 255
            var v22 uint16 = 32767
            _ = fi11
            _ = v16
            _ = v17
            _ = v21
            _ = v22
        }
        var v23 float64 = 1.5
        _ = fi10
        _ = v23
    }
    if (127 < 8) {
        var v24 int = 2147483647
        v4 = (v4 | (v5 % ((v5) | 1)))
        for fi25 := 0; fi25 < 12; fi25++ {
            if true {
                if true {
                    v5 = 2
                    v26 := GenS2{F0: -128, F1: -32768, F2: 255}
                    v4 += v9
                    _ = v26
                }
                _ = len(v2)
            } else {
                if len(v1) > 0 {
                    v1[0] = (v24 * v24)
                }
                if false {
                    fmt.Println(fi25, v4, v9)
                    if (false && false) {
                        for fi27 := 0; fi27 < 15; fi27++ {
                            v4 -= (^(v5))
                            _ = len(v1)
                            _ = fi27
                        }
                        var v28 uint = v4
                        switch "foo bar" {
                        case "":
                            switch v24 {
                            case 32767:
                                v4 += (v28 << ((18446744073709551615) & 63))
                            }
                        }
                        _ = v28
                    }
                    if true {
                        switch "a" {
                        case "":
                            v2 = append(v2, (fi25 ^ (v24 | (v24 | (v24 * 127)))))
                        case "hello":
                            var v29 int8 = 12
                            _ = v29
                        }
                    } else {
                        if (false && true) {
                            var v30 uint16 = 915
                            _ = v30
                        } else {
                            v31 := GenS1{F0: 65535}
                            var v32 int16 = 0
                            v5 -= (v4 * (v4 & v4))
                            _ = v31
                            _ = v32
                        }
                        v33 := GenS3{F0: (v9 | v5), F1: ((787 < 281) && (!false))}
                        _ = v33
                    }
                }
            }
            v34 := GenS2{F0: 100, F1: 604, F2: 42}
            v24 = fi25
            _ = fi25
            _ = v34
        }
        _ = v24
    } else {
        if (false && ((3.14159 * (0.1 / 3.14159)) != 0.1)) {
            var v35 int8 = 1
            v4 = 2
            _ = v35
        } else {
            _ = len(v1)
        }
    }
    if true {
        v9 = (v4 & (v4 << ((255) & 63)))
        v36 := GenS2{F0: -1, F1: -710, F2: 1}
        var v37 int = 2147483647
        _ = v36
        _ = v37
    } else {
        v38 := GenS3{F0: v9, F1: false}
        for fi39 := 0; fi39 < 8; fi39++ {
            var v40 string = ("foo bar" + "\n\t")
            for fi41 := 0; fi41 < 12; fi41++ {
                v42 := GenS2{F0: 0, F1: 32767, F2: 104}
                v40 = (("hello" + "") + v40)
                _ = fi41
                _ = v42
            }
            _ = fi39
            _ = v40
        }
        var v43 int64 = -356
        _ = v38
        _ = v43
    }
    if false {
        var v44 uint32 = 42
        _ = v44
    } else {
        fmt.Println(v9)
    }
    for fi45 := 0; fi45 < 18; fi45++ {
        delete(v3, -1)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi46 := 0; fi46 < 15; fi46++ {
            if (1 > -9223372036854775808) {
                fmt.Println(fi45)
                v47 := GenS2{F0: -1, F1: -71, F2: 1}
                switch fi46 {
                case 2147483647:
                    v9 = (v4 ^ (v4 + 65535))
                case 1:
                    switch ((("\n\t" + ((("\n\t" + "foo bar") + "foo bar") + "foo bar")) + "foo bar") + "a") {
                    case "hello":
                        v48 := GenS2{F0: -8, F1: 1000, F2: 188}
                        _ = v48
                    case "a":
                        switch ("hello" + ("\n\t" + ("a" + ""))) {
                        case "hello":
                            v9 -= v4
                        case "":
                            switch "foo bar" {
                            case "a":
                                v9 = (v5 * v9)
                            case "":
                                v9 -= v9
                            case "\n\t":
                                var v49 int = fi45
                                _ = v49
                            default:
                                v50 := GenS4{F0: 127, F1: -128}
                                _ = v50
                            }
                        default:
                            var v51 float64 = float64(1e10)
                            _ = v51
                        }
                    default:
                        v5 += (v5 + (^(v5)))
                        fmt.Println(fi45, v4, fi46)
                        var v52 string = (("\n\t" + "") + "a")
                        _ = v52
                    }
                }
                _ = v47
            } else {
                var v53 uint = v9
                v4 = 2
                switch (fi45 | 390) {
                case 2147483647:
                    switch "" {
                    case "\n\t":
                        var v54 uint16 = 1
                        _ = v54
                    case "a":
                        switch (fi45 / ((-197) | 1)) {
                        case 32768:
                            v5 = (v5 / ((v4) | 1))
                        case 2147483647:
                            fmt.Println(fi46, fi45)
                        }
                    case "foo bar":
                        v55 := GenS1{F0: 2}
                        _ = v55
                    default:
                        switch (("" + "") + (("a" + "") + "a")) {
                        case "a":
                            if true {
                                if (fi45 > fi46) {
                                    var v56 int64 = 9223372036854775807
                                    _ = v56
                                }
                            } else {
                                var v57 uint64 = 2
                                v9 = (v4 >> ((v53) & 63))
                                var v58 uint8 = 39
                                v4 += (v53 << (((v53 >> ((v4) & 63))) & 63))
                                _ = v57
                                _ = v58
                            }
                        }
                    }
                case -2147483648:
                    switch "foo bar" {
                    case "\n\t":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        if len(v1) > 0 {
                            v1[0] = -995
                        }
                    }
                case 203:
                    if (false || false) {
                        v9 -= v5
                        delete(v3, (fi46 - (^(fi45))))
                        _ = len(v3)
                        switch "foo bar" {
                        case "\n\t":
                            v4 = v4
                        case "hello":
                            switch "a" {
                            case "hello":
                                delete(v3, -657)
                            case "\n\t":
                                switch (fi45 ^ (fi45 % (((fi45 << ((v5) & 63))) | 1))) {
                                case 9223372036854775807:
                                    v4 += (v9 * v4)
                                }
                            default:
                                fmt.Println(v5, v4, fi45)
                            }
                        case "foo bar":
                            v59 := GenS2{F0: -1, F1: -1, F2: 1}
                            _ = v59
                        default:
                            v60 := GenS3{F0: (v53 ^ (v5 & 18446744073709551615)), F1: (540 >= 468)}
                            _ = v60
                        }
                        v1 = append(v1, (fi45 >> ((v9) & 63)))
                    } else {
                        v61 := GenS1{F0: 462}
                        _ = v61
                    }
                    delete(v3, fi45)
                    v53 = (v5 << ((v4) & 63))
                }
                _ = v53
            }
            var v62 uint32 = 42
            _ = fi46
            _ = v62
        }
        _ = fi45
    }
    v5 -= 554
    v63 := GenS5{F0: 4294967295, F1: -1}
    var v64 uint32 = 4294967295
    switch (("\n\t" + "foo bar") + ("a" + "hello")) {
    case "hello":
        for fi65 := 0; fi65 < 3; fi65++ {
            if (-32768 != 273) {
                var v66 int64 = 62
                _ = v66
            } else {
                for fi67 := 0; fi67 < 8; fi67++ {
                    if (65535 >= 0) {
                        switch fi65 {
                        case 1:
                            v68 := GenS6{F0: v9, F1: (((float64(0.1) * ((1e10 * 0.1) * 3.14159)) - float64(float32(float64(0.1)))) >= 0.1), F2: 1000}
                            _ = v68
                        }
                        v64 -= (v64 - v64)
                        if true {
                            var v69 uint64 = 2
                            _ = v69
                        }
                        v9 = v5
                        var v70 int64 = -1
                        _ = v70
                    } else {
                        var v71 uint16 = 65535
                        var v72 uint8 = 127
                        _ = v71
                        _ = v72
                    }
                    var v73 uint64 = uint64(v4)
                    _ = fi67
                    _ = v73
                }
                var v74 int16 = 32767
                fmt.Println(v4, fi65, v4)
                _ = v74
            }
            _ = fi65
        }
    default:
        for fi75 := 0; fi75 < 10; fi75++ {
            for fi76 := 0; fi76 < 20; fi76++ {
                v4 = v9
                var v77 float32 = 1e10
                var v78 int = fi76
                _ = fi76
                _ = v77
                _ = v78
            }
            fmt.Println(v9, v4)
            v64 += v64
            _ = fi75
        }
    }
    v79 := GenS6{F0: (v4 + v5), F1: false, F2: 984}
    v64 = 1
    var v80 int = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    fmt.Println("v80:", v80)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v64
    _ = v79
    _ = v80
    _ = v9
}
