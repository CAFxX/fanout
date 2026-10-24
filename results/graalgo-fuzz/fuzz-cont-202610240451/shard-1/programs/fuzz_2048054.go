package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8; F1 int; F2 int }

func main() {
    var v1 uint8 = 137
    v2 := make(map[int]uint32)
    var v3 string = (("" + ((("\n\t" + ("foo bar" + ("\n\t" + ""))) + "") + "\n\t")) + "foo bar")
    var v4 uint64 = 512
    v4 = v4
    v5 := []int{-2147483648, -421, -1, 4294967295}
    delete(v2, 0)
    fmt.Println(v1, v4, v3)
    fmt.Println(v4)
    v4 += (v4 % ((42) | 1))
    var v6 int64 = 1
    var v7 uint64 = (^(v4))
    fmt.Println(v6, v7)
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 float32 = 1.5
        v3 = ("hello" + "\n\t")
        v10 := GenS1{F0: (v1 + v1), F1: 937}
        _ = fi8
        _ = v10
        _ = v9
    }
    v4 = (v4 & v4)
    if ((v6 & -779) < v6) {
        switch 2147483647 {
        case -393:
            v11 := GenS2{F0: ((v3 + (v3 + v3)) + "a")}
            _ = v11
        case -213:
            for fi12 := 0; fi12 < 17; fi12++ {
                _ = len(v5)
                _ = fi12
            }
        case 9223372036854775807:
            for fi13 := 0; fi13 < 11; fi13++ {
                switch (fi13 | 32767) {
                case 9223372036854775807:
                    if (2 < v1) {
                        if (-2147483648 > 0) {
                            _ = len(v2)
                        } else {
                            fmt.Println(v6)
                            switch fi13 {
                            case -2147483648:
                                var v14 uint8 = v1
                                _ = v14
                            default:
                                var v15 int16 = 1000
                                _ = v15
                            }
                            if true {
                                v4 = 0
                                if (2147483647 != 2147483647) {
                                    var v16 int16 = -434
                                    var v17 bool = false
                                    _ = v16
                                    _ = v17
                                } else {
                                    switch ("a" + ("a" + v3)) {
                                    case "a":
                                        var v18 int64 = v6
                                        _ = v18
                                    case "":
                                        if len(v5) > 0 {
                                            v5[0] = (^(fi13))
                                        }
                                    case "foo bar":
                                        for fi19 := 0; fi19 < 16; fi19++ {
                                            if true {
                                                if false {
                                                    var v20 int8 = -1
                                                    if ((255 == (v1 & 1)) || (v4 == v4)) {
                                                        v21 := GenS1{F0: 0, F1: -829}
                                                        var v22 int64 = -1
                                                        switch ((v3 + v3) + (v3 + v3)) {
                                                        case "foo bar":
                                                            var v23 int16 = 0
                                                            _ = v23
                                                        }
                                                        _ = v21
                                                        _ = v22
                                                    }
                                                    if true {
                                                        var v24 uint8 = uint8(v7)
                                                        fmt.Println(v7)
                                                        _ = v24
                                                    }
                                                    _ = v20
                                                } else {
                                                    if (-1 != -581) {
                                                        if false {
                                                            if (!(-1 >= 484)) {
                                                                var v25 float32 = (1.5 + ((1.5 / -2.25) + 0.1))
                                                                _ = v25
                                                            } else {
                                                                if (!(("hello" + "hello") >= "\n\t")) {
                                                                    v26 := GenS1{F0: uint8(v4), F1: int32(v6)}
                                                                    v7 = 2
                                                                    _ = v26
                                                                } else {
                                                                    v27 := GenS2{F0: v3}
                                                                    var v28 bool = (!false)
                                                                    _ = v27
                                                                    _ = v28
                                                                }
                                                            }
                                                            switch ("\n\t" + (v3 + ("a" + (v3 + ((v3 + "\n\t") + v3))))) {
                                                            case "foo bar":
                                                                v4 = v7
                                                            case "a":
                                                                v3 = (v3 + "hello")
                                                            case "":
                                                                if (!true) {
                                                                    v4 += v7
                                                                    switch (-(fi19)) {
                                                                    case 32767:
                                                                        var v29 int16 = -983
                                                                        _ = v29
                                                                    case 2147483647:
                                                                        var v30 float64 = (0.0 + 0.0)
                                                                        _ = v30
                                                                    default:
                                                                        switch fi13 {
                                                                        case 128:
                                                                            _ = len(v2)
                                                                        case -9223372036854775808:
                                                                            v4 += v7
                                                                        default:
                                                                            switch v3 {
                                                                            case "foo bar":
                                                                                v6 = (v6 ^ (v6 * v6))
                                                                            case "a":
                                                                                v6 += 1
                                                                            }
                                                                        }
                                                                        fmt.Println(v6, v4)
                                                                    }
                                                                    if len(v5) > 0 {
                                                                        v5[0] = 4294967295
                                                                    }
                                                                    if true {
                                                                        v3 = (v3 + v3)
                                                                        v31 := GenS3{F0: 4294967295, F1: 450}
                                                                        _ = v31
                                                                    }
                                                                    var v32 float64 = float64(1e10)
                                                                    v7 -= 4294967295
                                                                    _ = v32
                                                                } else {
                                                                    fmt.Println(fi13, v7)
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        switch v3 {
                                                                        case "":
                                                                            var v33 int64 = int64(fi13)
                                                                            _ = v33
                                                                        default:
                                                                            switch (fi19 % (((^(fi19))) | 1)) {
                                                                            case 382:
                                                                                var v34 uint = 955
                                                                                _ = v34
                                                                            default:
                                                                                var v35 bool = (65535 <= 4294967295)
                                                                                _ = v35
                                                                            }
                                                                        }
                                                                    case "hello":
                                                                        v7 = v4
                                                                    case "\n\t":
                                                                        var v36 string = ("a" + "\n\t")
                                                                        _ = v36
                                                                    }
                                                                    switch v3 {
                                                                    case "":
                                                                        var v37 int8 = 43
                                                                        _ = v37
                                                                    case "foo bar":
                                                                        fmt.Println(fi13)
                                                                    case "a":
                                                                        var v38 int = fi19
                                                                        v6 = v6
                                                                        fmt.Println(fi19)
                                                                        _ = v38
                                                                    default:
                                                                        var v39 float32 = 3.14159
                                                                        _ = v39
                                                                    }
                                                                }
                                                            default:
                                                                if ((v1 - 109) > (v1 << ((838) & 63))) {
                                                                    if true {
                                                                        if false {
                                                                            fmt.Println(v7, v1, v1)
                                                                        }
                                                                        v40 := GenS3{F0: 335, F1: 1}
                                                                        v41 := GenS1{F0: 186, F1: -1}
                                                                        _ = v40
                                                                        _ = v41
                                                                    } else {
                                                                        v3 = "hello"
                                                                    }
                                                                    switch "hello" {
                                                                    case "":
                                                                        if true {
                                                                            switch (fi13 | fi13) {
                                                                            case -7:
                                                                                v1 = (^(v1))
                                                                            case -836:
                                                                                var v42 float64 = 0.1
                                                                                _ = v42
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            var v43 uint16 = 65535
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            _ = v43
                                                        } else {
                                                            var v44 int32 = 42
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            fmt.Println(fi13, v4, v6)
                                                            _ = v44
                                                        }
                                                        v45 := GenS1{F0: v1, F1: -1}
                                                        var v46 uint = 255
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                    v3 = ""
                                                }
                                                v47 := GenS1{F0: 37, F1: 1}
                                                v48 := GenS4{F0: v3, F1: 4294967295}
                                                _ = v47
                                                _ = v48
                                            }
                                            v49 := v2[(-(fi19))]
                                            var v50 int32 = 0
                                            _ = fi19
                                            _ = v49
                                            _ = v50
                                        }
                                    default:
                                        v51 := GenS1{F0: (^(v1)), F1: 427}
                                        _ = v51
                                    }
                                    switch fi13 {
                                    case 2147483647:
                                        v4 += (v4 / (((^(v7))) | 1))
                                    }
                                }
                            }
                        }
                        fmt.Println(v3, v1, v3)
                        if len(v5) > 0 {
                            v5[0] = fi13
                        }
                    } else {
                        v52 := GenS5{F0: fi13}
                        v53 := GenS5{F0: -755}
                        var v54 int64 = v6
                        _ = v52
                        _ = v53
                        _ = v54
                    }
                case -2147483648:
                    if true {
                        switch "" {
                        case "\n\t":
                            v6 += 1
                        case "hello":
                            if true {
                                v4 = (v7 << ((4294967295) & 63))
                                v5 = append(v5, (^(fi13)))
                            } else {
                                var v55 int64 = (^(v6))
                                switch (fi13 - 32767) {
                                case -220:
                                    var v56 uint16 = 249
                                    _ = v56
                                }
                                _ = v55
                            }
                        default:
                            if (136 >= 42) {
                                fmt.Println(v1)
                            }
                        }
                        var v57 int16 = 0
                        var v58 float32 = ((1e10 + 0.0) + 3.14159)
                        _ = v57
                        _ = v58
                    }
                }
                var v59 int16 = 1
                _ = len(v5)
                _ = fi13
                _ = v59
            }
        }
    }
    if false {
        _ = len(v2)
    } else {
        v60 := GenS3{F0: 255, F1: 680}
        for fi61 := 0; fi61 < 11; fi61++ {
            for fi62 := 0; fi62 < 1; fi62++ {
                v63 := GenS6{F0: (v1 - 127), F1: fi61, F2: 127}
                var v64 uint16 = 608
                _ = fi62
                _ = v63
                _ = v64
            }
            fmt.Println(v1)
            v3 = ("foo bar" + ("\n\t" + v3))
            _ = fi61
        }
        fmt.Println(v3)
        _ = v60
    }
    var v65 int16 = -1
    v7 += v4
    for fi66 := 0; fi66 < 15; fi66++ {
        if len(v5) > 0 {
            v5[0] = (-(fi66))
        }
        _ = fi66
    }
    for fi67 := 0; fi67 < 6; fi67++ {
        fmt.Println(fi67)
        if false {
            var v68 bool = false
            _ = v68
        } else {
            switch v3 {
            case "\n\t":
                if false {
                    var v69 uint8 = v1
                    v70 := v2[(^(fi67))]
                    _ = v69
                    _ = v70
                } else {
                    v65 = -505
                    for fi71 := 0; fi71 < 16; fi71++ {
                        v72 := GenS3{F0: 4294967295, F1: 2}
                        _ = fi71
                        _ = v72
                    }
                }
            default:
                switch "a" {
                case "a":
                    fmt.Println(v7)
                default:
                    var v73 uint = 18446744073709551615
                    _ = v73
                }
            }
            _ = len(v5)
            v4 += v7
        }
        _ = fi67
    }
    var v74 uint8 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v7:", v7)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v7
    _ = v74
}
