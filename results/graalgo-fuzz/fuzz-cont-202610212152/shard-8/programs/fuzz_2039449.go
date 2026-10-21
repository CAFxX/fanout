package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS4 struct { F0 int16; F1 uint32; F2 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS9 struct { F0 uint8; F1 uint32 }
type GenS10 struct { F0 uint64; F1 int64; F2 uint }
type GenS11 struct { F0 uint; F1 int16; F2 uint16 }

func main() {
    var v1 uint = 774
    v2 := make(map[string]int)
    v3 := []string{("foo bar" + "\n\t"), "hello", (("\n\t" + "") + "\n\t")}
    var v4 int = 792
    v5 := make(map[string]string)
    var v6 uint = v1
    var v7 int64 = 42
    for fi8 := 0; fi8 < 5; fi8++ {
        if true {
            if (0.0 == float32(3.14159)) {
                for fi9 := 0; fi9 < 7; fi9++ {
                    switch "a" {
                    case "a":
                        if true {
                            v10 := GenS1{F0: v1}
                            var v11 uint64 = 4294967295
                            _ = v10
                            _ = v11
                        }
                    case "foo bar":
                        if (false && false) {
                            v12 := GenS1{F0: (v6 / (((v6 / ((v1) | 1))) | 1))}
                            var v13 uint64 = 333
                            v5["hello"] = "\n\t"
                            var v14 uint64 = 34
                            var v15 int = 4294967295
                            v15 -= 32768
                            _ = v12
                            _ = v13
                            _ = v14
                            _ = v15
                        } else {
                            v1 -= (v1 << ((v1) & 63))
                        }
                    }
                    fmt.Println(v6, v6)
                    var v16 int32 = int32(v4)
                    var v17 int16 = 32767
                    switch "" {
                    case "hello":
                        v18 := GenS1{F0: v6}
                        _ = v18
                    case "":
                        v7 += (v7 / (((v7 + v7)) | 1))
                    case "foo bar":
                        if len(v3) > 0 {
                            v3[0] = ("hello" + ("a" + ""))
                        }
                    default:
                        switch (((((("hello" + "foo bar") + "foo bar") + "\n\t") + "") + "\n\t") + ("foo bar" + (("hello" + "a") + ""))) {
                        case "":
                            if true {
                                fmt.Println(v17, v17)
                                if false {
                                    if false {
                                        fmt.Println(v1, v4)
                                        v4 = (v4 << ((v6) & 63))
                                        v17 -= (^(v17))
                                    }
                                } else {
                                    v19 := v5["\n\t"]
                                    var v20 float64 = (1e10 / -2.25)
                                    _ = v19
                                    _ = v20
                                }
                            } else {
                                if true {
                                    _ = len(v5)
                                    v3 = append(v3, ("\n\t" + ""))
                                } else {
                                    v21 := GenS2{F0: v16, F1: -1}
                                    switch v4 {
                                    case 32767:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    case 4294967295:
                                        switch ("\n\t" + "foo bar") {
                                        case "\n\t":
                                            delete(v2, "foo bar")
                                        }
                                    case 0:
                                        _ = len(v3)
                                    }
                                    switch "hello" {
                                    case "":
                                        if true {
                                            fmt.Println(fi9)
                                        }
                                    case "foo bar":
                                        var v22 uint16 = 1
                                        _ = v22
                                    default:
                                        switch ("a" + ("\n\t" + (((("\n\t" + "\n\t") + "foo bar") + "hello") + (("" + (("hello" + "\n\t") + "foo bar")) + "")))) {
                                        case "foo bar":
                                            var v23 uint16 = 130
                                            _ = v23
                                        default:
                                            if false {
                                                v2["\n\t"] = 2147483647
                                                var v24 int8 = 100
                                                _ = v24
                                            } else {
                                                if (((!(!true)) || (false || false)) || (((v17 ^ v17) > (v17 << ((v6) & 63))) && false)) {
                                                    var v25 uint8 = 255
                                                    _ = v25
                                                }
                                            }
                                        }
                                    }
                                    _ = v21
                                }
                            }
                            switch 1 {
                            case 9223372036854775807:
                                switch (fi9 + -2147483648) {
                                case 284:
                                    switch 128 {
                                    case 133:
                                        var v26 int = 851
                                        _ = v26
                                    case 66:
                                        var v27 int32 = v16
                                        _ = v27
                                    default:
                                        if (!(((18446744073709551615 <= 18446744073709551615) || ((-(v16)) > v16)) || true)) {
                                            v28 := GenS3{F0: 42, F1: uint32(v6), F2: 84}
                                            _ = v28
                                        } else {
                                            fmt.Println(v4, v1)
                                            v4 = (v4 << ((42) & 63))
                                        }
                                    }
                                    v6 = v6
                                default:
                                    switch (fi9 % ((9223372036854775807) | 1)) {
                                    case 1:
                                        if ((fi8 / ((v4) | 1)) == (fi8 & fi9)) {
                                            fmt.Println(fi8)
                                            v4 = (-(fi9))
                                        }
                                    case 2147483647:
                                        var v29 uint16 = 64
                                        _ = v29
                                    case 4294967295:
                                        v30 := GenS1{F0: (v1 / (((v1 & v6)) | 1))}
                                        _ = v30
                                    }
                                }
                            default:
                                var v31 uint64 = 4294967295
                                _ = v31
                            }
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    }
                    _ = len(v3)
                    var v32 bool = true
                    v33 := GenS4{F0: 639, F1: 2147483648, F2: v32}
                    var v34 uint16 = 2
                    _ = fi9
                    _ = v16
                    _ = v17
                    _ = v32
                    _ = v33
                    _ = v34
                }
                if true {
                    switch "hello" {
                    case "foo bar":
                        var v35 uint64 = 454
                        _ = v35
                    case "a":
                        var v36 uint32 = 2
                        _ = v36
                    default:
                        _ = len(v2)
                    }
                } else {
                    v37 := GenS5{F0: 32767}
                    if false {
                        for fi38 := 0; fi38 < 7; fi38++ {
                            v7 += v7
                            fmt.Println(v1)
                            _ = fi38
                        }
                    }
                    _ = v37
                }
            } else {
                fmt.Println(fi8, fi8, fi8)
                v39 := GenS6{F0: 65535}
                _ = v39
            }
            v7 += (-(v7))
        } else {
            v7 = v7
            for fi40 := 0; fi40 < 17; fi40++ {
                if true {
                    if true {
                        var v41 uint16 = 551
                        if false {
                            switch "foo bar" {
                            case "a":
                                var v42 float32 = (1e10 * 0.0)
                                _ = v42
                            case "foo bar":
                                fmt.Println(v7)
                            }
                        } else {
                            fmt.Println(fi40)
                            v3 = append(v3, ("hello" + "hello"))
                        }
                        _ = v41
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v1 = 554
                    }
                    var v43 uint16 = 0
                    if true {
                        v2[(("" + ("a" + (("" + "a") + ""))) + "\n\t")] = -1
                        switch "foo bar" {
                        case "\n\t":
                            v7 -= (v7 + 412)
                        case "foo bar":
                            v43 = (v43 & v43)
                            var v44 uint64 = 18446744073709551615
                            if ((1.5 - 0.1) <= (0.1 + 0.1)) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v44 += (v44 >> (((v6 * (v6 << ((v1) & 63)))) & 63))
                                v6 = v6
                                var v45 int16 = 1
                                if false {
                                    v44 += (v44 + (^(v44)))
                                }
                                var v46 string = (("\n\t" + "") + "foo bar")
                                var v47 uint64 = 383
                                _ = v45
                                _ = v46
                                _ = v47
                            }
                            v2["\n\t"] = fi8
                            v7 += v7
                            _ = v44
                        case "":
                            var v48 int64 = 9223372036854775807
                            _ = v48
                        default:
                            v6 = (v1 ^ 18446744073709551615)
                        }
                        switch "hello" {
                        case "a":
                            var v49 uint16 = (v43 >> (((v6 ^ (v6 - (v6 + v1)))) & 63))
                            _ = v49
                        }
                    } else {
                        v6 = v1
                        v6 = (^(v6))
                    }
                    v50 := GenS7{F0: -264}
                    fmt.Println(v43)
                    v43 = 2
                    v7 -= int64(v1)
                    var v51 int16 = 0
                    _ = v43
                    _ = v50
                    _ = v51
                }
                _ = fi40
            }
            switch -9223372036854775808 {
            case 42:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                if false {
                    switch fi8 {
                    case -173:
                        v1 += v6
                    }
                }
            }
        }
        _ = fi8
    }
    v52 := v2[""]
    v4 = (v4 >> (((v6 / ((v1) | 1))) & 63))
    fmt.Println(v4, v1, v1)
    fmt.Println(v6, v1, v4)
    for fi53 := 0; fi53 < 13; fi53++ {
        var v54 uint64 = 422
        _ = fi53
        _ = v54
    }
    if (32767 <= 594) {
        for fi55 := 0; fi55 < 14; fi55++ {
            _ = len(v3)
            v56 := GenS8{F0: 1, F1: 4294967295, F2: 0}
            _ = fi55
            _ = v56
        }
    } else {
        v57 := GenS9{F0: 0, F1: 1}
        var v58 uint32 = 1
        _ = v57
        _ = v58
    }
    fmt.Println(v52, v4)
    if len(v3) > 0 {
        v3[0] = ("foo bar" + "foo bar")
    }
    _ = len(v3)
    if ((!false) && ((uint64(v1) > 4294967295) || false)) {
        var v59 int = int(v1)
        _ = v59
    } else {
        for fi60 := 0; fi60 < 13; fi60++ {
            if false {
                for fi61 := 0; fi61 < 14; fi61++ {
                    var v62 uint64 = 252
                    switch "hello" {
                    case "\n\t":
                        v63 := GenS6{F0: 828}
                        _ = v63
                    }
                    _ = fi61
                    _ = v62
                }
                v64 := GenS10{F0: 1, F1: v7, F2: v1}
                v52 -= 128
                _ = v64
            }
            _ = fi60
        }
        for fi65 := 0; fi65 < 19; fi65++ {
            for fi66 := 0; fi66 < 11; fi66++ {
                var v67 uint8 = 1
                if (!(true && (!true))) {
                    v68 := GenS11{F0: v6, F1: 0, F2: 0}
                    fmt.Println(v67)
                    _ = v68
                }
                _ = fi66
                _ = v67
            }
            _ = fi65
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
    _ = v7
}
