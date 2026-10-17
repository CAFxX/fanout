package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint }
type GenS2 struct { F0 bool; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 bool; F1 bool }

func main() {
    var v1 string = "hello"
    v2 := []uint32{169, 710, 4294967295, 65535}
    var v3 string = "a"
    var v4 uint32 = 2
    v5 := []uint32{v4}
    var v6 int64 = -9223372036854775808
    var v7 uint16 = 925
    var v8 int64 = 1
    v5 = append(v5, (v4 / (((v4 - (v4 >> ((18446744073709551615) & 63)))) | 1)))
    fmt.Println(v3, v4, v3)
    if (false && false) {
        _ = len(v5)
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := GenS1{F0: 0, F1: v3, F2: 42}
        switch "a" {
        case "":
            var v11 int8 = 100
            _ = v11
        case "foo bar":
            var v12 int32 = -537
            _ = v12
        case "hello":
            switch v1 {
            case "a":
                v13 := GenS2{F0: true, F1: 1, F2: -1}
                _ = v13
            case "foo bar":
                _ = len(v5)
            case "\n\t":
                _ = len(v5)
            }
        default:
            var v14 int = fi9
            _ = v14
        }
        v2 = append(v2, 1)
        _ = fi9
        _ = v10
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (v3 < (v3 + "a")) {
        v15 := GenS3{F0: 551, F1: v7, F2: 22}
        _ = len(v5)
        _ = v15
    }
    v16 := GenS4{F0: 1, F1: 160}
    switch ("" + "\n\t") {
    case "foo bar":
        for fi17 := 0; fi17 < 5; fi17++ {
            if false {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            } else {
                if (1 <= 384) {
                    for fi18 := 0; fi18 < 3; fi18++ {
                        switch ("\n\t" + v3) {
                        case "hello":
                            v7 -= (v7 >> ((2) & 63))
                        case "a":
                            v19 := GenS3{F0: 54, F1: (v7 & 475), F2: 127}
                            _ = v19
                        case "foo bar":
                            fmt.Println(fi17)
                        default:
                            switch v1 {
                            case "foo bar":
                                fmt.Println(v8, v7, v4)
                            case "hello":
                                if false {
                                    var v20 float32 = 0.0
                                    _ = v20
                                }
                            case "a":
                                var v21 string = v3
                                _ = v21
                            }
                        }
                        var v22 int = (fi18 - (fi17 & fi18))
                        if (127 > 38) {
                            if (!true) {
                                if ("\n\t" != "hello") {
                                    v23 := GenS5{F0: true, F1: true}
                                    v5 = append(v5, 2147483647)
                                    _ = v23
                                }
                            }
                            if (true && false) {
                                v6 = 457
                            } else {
                                if true {
                                    switch v3 {
                                    case "a":
                                        if false {
                                            switch ((v3 + v1) + (v1 + (v3 + (v1 + ("" + v3))))) {
                                            case "hello":
                                                if ((!(255 < (v7 & (v7 % ((v7) | 1))))) || (true || (float32(((0.1 - (3.14159 / 1.5)) * 0.0)) <= 3.14159))) {
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        v1 = "\n\t"
                                                    case "a":
                                                        v8 += v8
                                                    case "foo bar":
                                                        if len(v2) > 0 {
                                                            v2[0] = v4
                                                        }
                                                    default:
                                                        _ = len(v2)
                                                    }
                                                }
                                            case "foo bar":
                                                v22 = fi18
                                            case "\n\t":
                                                v24 := GenS1{F0: -594, F1: "hello", F2: 695}
                                                _ = v24
                                            }
                                            var v25 string = v3
                                            _ = v25
                                        } else {
                                            v8 -= int64(v4)
                                            v6 -= 42
                                        }
                                    case "hello":
                                        switch (v22 * (fi18 - (fi17 | (fi17 * (^(v22)))))) {
                                        case 0:
                                            var v26 int32 = int32(v22)
                                            _ = v26
                                        case -9223372036854775808:
                                            v4 -= (v4 ^ (v4 + v4))
                                            fmt.Println(v7)
                                        default:
                                            _ = len(v5)
                                        }
                                    case "":
                                        var v27 uint = 1
                                        _ = v27
                                    default:
                                        fmt.Println(v4, fi18)
                                    }
                                    var v28 uint32 = (v4 | 0)
                                    var v29 int8 = 0
                                    v7 += (^(v7))
                                    _ = v28
                                    _ = v29
                                }
                                _ = len(v5)
                                if (((((((v3 + v3) + "\n\t") + v3) + (v3 + "hello")) > v3) && (42 < 42)) && (!false)) {
                                    v1 = v1
                                    v30 := GenS2{F0: false, F1: 2, F2: 32767}
                                    switch -997 {
                                    case -9223372036854775808:
                                        var v31 bool = true
                                        if v31 {
                                            var v32 string = (v3 + v3)
                                            _ = v32
                                        } else {
                                            var v33 uint16 = v7
                                            v6 = (v6 >> ((0) & 63))
                                            _ = v33
                                        }
                                        _ = v31
                                    case 32768:
                                        var v34 int64 = v8
                                        _ = v34
                                    }
                                    _ = v30
                                }
                            }
                            v35 := GenS3{F0: 4294967295, F1: v7, F2: 127}
                            if true {
                                fmt.Println(v22, v3, v4)
                                switch (fi18 + (fi17 >> ((1) & 63))) {
                                case 668:
                                    fmt.Println(fi17)
                                case -9223372036854775808:
                                    var v36 float64 = float64((0.1 * 0.0))
                                    _ = v36
                                case 4294967295:
                                    fmt.Println(v1, v22, v8)
                                default:
                                    v8 = (v8 | 9223372036854775807)
                                }
                                var v37 uint32 = v4
                                _ = v37
                            } else {
                                _ = len(v2)
                                var v38 float64 = (0.1 - 0.0)
                                _ = v38
                            }
                            _ = v35
                        }
                        _ = fi18
                        _ = v22
                    }
                    var v39 float32 = ((0.1 * 0.0) / 0.1)
                    _ = v39
                } else {
                    v1 = "a"
                    for fi40 := 0; fi40 < 20; fi40++ {
                        if ((false || true) || (!(-1 < -2147483648))) {
                            if (435 > 65535) {
                                v41 := GenS4{F0: 42, F1: 255}
                                _ = v41
                            } else {
                                var v42 float32 = (-2.25 + 0.0)
                                v43 := GenS2{F0: ((1.5 * 1.5) >= 1.5), F1: 1, F2: 42}
                                if true {
                                    _ = len(v5)
                                    v4 += (v4 % (((v4 & v4)) | 1))
                                } else {
                                    fmt.Println(v7, v8)
                                    fmt.Println(v8, v6)
                                }
                                _ = v42
                                _ = v43
                            }
                            var v44 uint16 = v7
                            _ = v44
                        } else {
                            fmt.Println(v6, fi40, v6)
                            var v45 int32 = -2147483648
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v7 = v7
                            v45 = v45
                            var v46 uint = 1
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            switch (fi17 >> ((v46) & 63)) {
                            case -566:
                                if len(v2) > 0 {
                                    v2[0] = (v4 - 199)
                                }
                            }
                            if (!true) {
                                var v47 uint64 = 83
                                _ = v47
                            }
                            var v48 string = "foo bar"
                            _ = v45
                            _ = v46
                            _ = v48
                        }
                        fmt.Println(v7, v3)
                        _ = fi40
                    }
                }
                var v49 uint32 = (v4 * v4)
                _ = v49
            }
            _ = fi17
        }
    case "\n\t":
        v50 := GenS4{F0: 127, F1: 2}
        _ = v50
    }
    fmt.Println(v6)
    switch v3 {
    case "a":
        _ = len(v2)
    case "":
        var v51 uint16 = v7
        _ = v51
    default:
        switch v3 {
        case "\n\t":
            var v52 int64 = (v6 % (((v8 % ((v8) | 1))) | 1))
            _ = v52
        default:
            if false {
                var v53 uint64 = 4294967295
                var v54 int16 = int16(v4)
                _ = v53
                _ = v54
            }
        }
    }
    var v55 uint16 = (v7 & v7)
    if false {
        var v56 int32 = -2147483648
        if (false || false) {
            for fi57 := 0; fi57 < 20; fi57++ {
                v58 := GenS5{F0: (true && true), F1: false}
                v5 = append(v5, v4)
                _ = len(v2)
                _ = fi57
                _ = v58
            }
        }
        var v59 uint32 = 42
        _ = v56
        _ = v59
    } else {
        for fi60 := 0; fi60 < 6; fi60++ {
            _ = len(v5)
            var v61 uint32 = 0
            v62 := GenS4{F0: 152, F1: 27}
            _ = fi60
            _ = v61
            _ = v62
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
    _ = v7
    _ = v8
}
