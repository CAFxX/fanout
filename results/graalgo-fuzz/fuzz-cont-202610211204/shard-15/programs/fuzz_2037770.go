package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint32; F1 string; F2 bool }
type GenS4 struct { F0 int8; F1 int64; F2 uint32 }
type GenS5 struct { F0 int8; F1 uint }
type GenS6 struct { F0 int8 }

func main() {
    var v1 int64 = -311
    v2 := []int{2147483647, -627, -218, 544}
    var v3 uint = 18446744073709551615
    var v4 int16 = -1
    v5 := []uint32{}
    var v6 uint = (^(v3))
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 uint8 = 132
        _ = fi7
        _ = v8
    }
    if (0.1 < 0.1) {
        fmt.Println(v1, v4)
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 float32 = -2.25
            var v11 uint8 = 2
            switch 32768 {
            case 128:
                _ = len(v2)
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        v3 -= (^(v3))
    } else {
        v12 := GenS1{F0: 0, F1: 0}
        _ = v12
    }
    var v13 uint = v6
    _ = len(v5)
    switch "hello" {
    case "hello":
        fmt.Println(v4, v1, v4)
    case "foo bar":
        if false {
            switch 4294967295 {
            case -997:
                switch (("" + ("foo bar" + "\n\t")) + "hello") {
                case "a":
                    v14 := GenS1{F0: 42, F1: 2147483647}
                    _ = v14
                default:
                    for fi15 := 0; fi15 < 6; fi15++ {
                        switch (fi15 * (fi15 & (fi15 + (fi15 | (fi15 / ((int(v1)) | 1)))))) {
                        case 4294967295:
                            fmt.Println(fi15)
                        case -2147483648:
                            for fi16 := 0; fi16 < 8; fi16++ {
                                switch fi16 {
                                case 598:
                                    if len(v2) > 0 {
                                        v2[0] = (fi15 * fi16)
                                    }
                                case 0:
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                case -466:
                                    fmt.Println(v13)
                                }
                                var v17 int64 = v1
                                if true {
                                    fmt.Println(v6)
                                } else {
                                    var v18 int64 = v17
                                    if (!false) {
                                        v13 = v3
                                        _ = len(v5)
                                        var v19 uint8 = 42
                                        _ = v19
                                    }
                                    v18 += 0
                                    _ = v18
                                }
                                _ = fi16
                                _ = v17
                            }
                        }
                        switch fi15 {
                        case -2147483648:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case 4294967295:
                            var v20 uint = 18446744073709551615
                            _ = v20
                        case 127:
                            v21 := GenS1{F0: v1, F1: -91}
                            _ = v21
                        default:
                            for fi22 := 0; fi22 < 9; fi22++ {
                                var v23 int16 = 32767
                                _ = fi22
                                _ = v23
                            }
                        }
                        v24 := GenS1{F0: -9223372036854775808, F1: 42}
                        _ = fi15
                        _ = v24
                    }
                }
            case 32767:
                fmt.Println(v4, v4, v13)
            case 128:
                for fi25 := 0; fi25 < 4; fi25++ {
                    switch "foo bar" {
                    case "hello":
                        for fi26 := 0; fi26 < 16; fi26++ {
                            if false {
                                var v27 uint16 = 255
                                if (!false) {
                                    if (!((-2147483648 < 42) && false)) {
                                        if false {
                                            v28 := GenS1{F0: -9223372036854775808, F1: -2147483648}
                                            if (false || (0 == 2)) {
                                                v4 -= v4
                                                var v29 uint16 = (v27 | (^(v27)))
                                                v13 -= 346
                                                _ = v29
                                            } else {
                                                _ = len(v5)
                                            }
                                            _ = v28
                                        }
                                        var v30 bool = (!true)
                                        v6 = (v6 >> ((v13) & 63))
                                        if (!(v30 && v30)) {
                                            fmt.Println(v4, v3, fi25)
                                            fmt.Println(v6)
                                            fmt.Println(fi25)
                                            var v31 float32 = float32(0.0)
                                            v27 = 42
                                            _ = v31
                                        }
                                        v27 -= 535
                                        _ = v30
                                    } else {
                                        fmt.Println(v13, v27, v4)
                                        v3 = (v13 << (((v13 >> ((255) & 63))) & 63))
                                        v4 += v4
                                        v32 := GenS2{F0: v27, F1: int8(v1)}
                                        _ = v32
                                    }
                                    if false {
                                        if true {
                                            v5 = append(v5, 0)
                                            fmt.Println(fi26, v27, v4)
                                            if true {
                                                var v33 uint8 = 127
                                                v34 := GenS3{F0: 2, F1: "hello", F2: true}
                                                v1 = int64(v4)
                                                _ = v33
                                                _ = v34
                                            } else {
                                                switch (("a" + ("foo bar" + "hello")) + ("\n\t" + "a")) {
                                                case "foo bar":
                                                    fmt.Println(v6)
                                                case "hello":
                                                    if len(v5) > 0 {
                                                        v5[0] = 4294967295
                                                    }
                                                default:
                                                    if (v4 >= v4) {
                                                        var v35 uint = (v3 * v3)
                                                        _ = v35
                                                    }
                                                }
                                                fmt.Println(v6, v13, v13)
                                                if ((v27 + (v27 << (((v3 | (v3 * 130))) & 63))) == v27) {
                                                    var v36 uint16 = v27
                                                    _ = v36
                                                } else {
                                                    _ = len(v2)
                                                }
                                            }
                                        }
                                        v1 += v1
                                    }
                                }
                                _ = v27
                            } else {
                                v37 := GenS4{F0: 0, F1: (v1 ^ v1), F2: 4294967295}
                                v38 := GenS5{F0: 127, F1: v3}
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                var v39 uint64 = 2
                                v40 := GenS1{F0: (^(v1)), F1: 230}
                                _ = v37
                                _ = v38
                                _ = v39
                                _ = v40
                            }
                            v2 = append(v2, fi26)
                            _ = fi26
                        }
                    case "":
                        if true {
                            v3 -= (^(v3))
                            v41 := GenS3{F0: 4294967295, F1: (("\n\t" + "foo bar") + ""), F2: false}
                            v13 = 1
                            _ = v41
                        } else {
                            switch "\n\t" {
                            case "\n\t":
                                v1 = (v1 >> ((255) & 63))
                            case "a":
                                v13 = (v6 & (v13 * 18446744073709551615))
                            case "foo bar":
                                fmt.Println(v1, v1, v13)
                            default:
                                var v42 float64 = -2.25
                                _ = v42
                            }
                        }
                    case "foo bar":
                        v1 = (v1 >> ((850) & 63))
                    default:
                        if (false || ((0.1 + 1e10) <= (0.0 - (float32(1.5) * 1e10)))) {
                            v4 += (v4 ^ -32768)
                        } else {
                            var v43 uint = 0
                            switch fi25 {
                            case 9223372036854775807:
                                v1 += (^(v1))
                            case -9223372036854775808:
                                if len(v5) > 0 {
                                    v5[0] = 65535
                                }
                            default:
                                v1 = v1
                            }
                            v44 := GenS4{F0: 13, F1: (v1 * (v1 << (((v13 | v3)) & 63))), F2: 4294967295}
                            _ = v43
                            _ = v44
                        }
                    }
                    v45 := GenS2{F0: 1, F1: -128}
                    v46 := GenS2{F0: 402, F1: -1}
                    _ = fi25
                    _ = v45
                    _ = v46
                }
            default:
                var v47 bool = (-82 >= int8(v6))
                _ = v47
            }
        }
    case "":
        v48 := GenS6{F0: -1}
        _ = v48
    }
    for fi49 := 0; fi49 < 11; fi49++ {
        v1 -= v1
        _ = fi49
    }
    switch 773 {
    case 107:
        fmt.Println(v1, v4)
    }
    v6 -= 972
    var v50 int = 586
    v51 := GenS4{F0: 0, F1: v1, F2: 2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
