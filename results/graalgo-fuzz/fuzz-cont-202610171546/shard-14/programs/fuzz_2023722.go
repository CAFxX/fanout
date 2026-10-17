package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int32; F1 int64; F2 int16 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 uint16; F1 bool }
type GenS8 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS9 struct { F0 uint; F1 uint }
type GenS10 struct { F0 uint16; F1 int8 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint32 }

func main() {
    v1 := []uint32{2147483648, 2147483648}
    v2 := make(map[string]uint32)
    var v3 bool = false
    v4 := make(map[string]int)
    v5 := make(map[int]string)
    var v6 uint64 = 933
    if v3 {
        v6 = (v6 ^ (v6 | 808))
        v7 := GenS1{F0: -1, F1: -128}
        _ = v7
    }
    delete(v4, "foo bar")
    switch -47 {
    case 32767:
        _ = len(v1)
    case -2147483648:
        var v8 int32 = 1
        _ = v8
    case 127:
        var v9 int8 = 101
        _ = v9
    }
    v10 := GenS1{F0: 128, F1: -128}
    switch 128 {
    case 172:
        if v3 {
            switch ("" + "") {
            case "\n\t":
                switch "hello" {
                case "\n\t":
                    switch "a" {
                    case "":
                        if v3 {
                            delete(v2, "\n\t")
                            switch "a" {
                            case "hello":
                                v3 = (!(!(v3 && (!v3))))
                            default:
                                fmt.Println(v6)
                            }
                            v3 = v3
                        } else {
                            switch 604 {
                            case -2147483648:
                                switch ((((("\n\t" + "hello") + "\n\t") + ("" + ("" + "a"))) + (("" + "hello") + ("a" + ""))) + ("foo bar" + "")) {
                                case "":
                                    for fi11 := 0; fi11 < 10; fi11++ {
                                        if len(v1) > 0 {
                                            v1[0] = 591
                                        }
                                        var v12 uint = 859
                                        _ = fi11
                                        _ = v12
                                    }
                                case "a":
                                    var v13 float32 = float32(-2.25)
                                    _ = v13
                                }
                            case -1:
                                v3 = (65535 > 425)
                            case -9223372036854775808:
                                v14 := GenS1{F0: 333, F1: 127}
                                _ = v14
                            default:
                                if len(v1) > 0 {
                                    v1[0] = 433
                                }
                            }
                            v3 = true
                            for fi15 := 0; fi15 < 17; fi15++ {
                                if ((0 <= 9223372036854775807) && (302 != 792)) {
                                    if v3 {
                                        if false {
                                            fmt.Println(v6, v6, v6)
                                            if v3 {
                                                v6 += (v6 - 997)
                                                for fi16 := 0; fi16 < 17; fi16++ {
                                                    v4["foo bar"] = (fi16 ^ (fi15 >> ((780) & 63)))
                                                    _ = fi16
                                                }
                                            }
                                            if v3 {
                                                fmt.Println(fi15, v6)
                                                var v17 float64 = (((0.1 + 1.5) - (0.0 + 1e10)) - 1.5)
                                                _ = v17
                                            }
                                        }
                                    }
                                    var v18 uint64 = (v6 * 780)
                                    _ = v18
                                }
                                _ = fi15
                            }
                        }
                    case "foo bar":
                        v6 += 42
                    }
                case "hello":
                    v3 = (128 < 9223372036854775807)
                case "a":
                    v4["\n\t"] = -985
                default:
                    var v19 int8 = 42
                    _ = v19
                }
            case "hello":
                v20 := GenS2{F0: 2147483647, F1: -1, F2: 254}
                _ = v20
            case "":
                if true {
                    switch 0 {
                    case 32768:
                        for fi21 := 0; fi21 < 8; fi21++ {
                            var v22 uint8 = 127
                            v6 = 0
                            _ = fi21
                            _ = v22
                        }
                    case 0:
                        var v23 int8 = 106
                        _ = v23
                    default:
                        switch ("\n\t" + "hello") {
                        case "\n\t":
                            if v3 {
                                v3 = true
                                v24 := GenS2{F0: 1, F1: 1, F2: -44}
                                _ = v24
                            } else {
                                if v3 {
                                    v6 = (v6 | v6)
                                    switch 0 {
                                    case 0:
                                        v6 = (^(v6))
                                    case -359:
                                        var v25 int32 = -1
                                        _ = v25
                                    case -331:
                                        _ = len(v2)
                                    }
                                    var v26 int16 = 61
                                    _ = v26
                                }
                                if v3 {
                                    var v27 int64 = 9223372036854775807
                                    _ = v27
                                } else {
                                    fmt.Println(v3)
                                }
                            }
                        case "a":
                            for fi28 := 0; fi28 < 4; fi28++ {
                                var v29 int16 = 42
                                if ((0.0 == 1.5) && v3) {
                                    delete(v4, "a")
                                } else {
                                    v3 = v3
                                    v1 = append(v1, 760)
                                    if (!true) {
                                        v30 := GenS2{F0: -594, F1: 9223372036854775807, F2: v29}
                                        fmt.Println(v29, fi28, fi28)
                                        v29 = v29
                                        _ = v30
                                    } else {
                                        if false {
                                            v3 = (v29 != -1)
                                        }
                                        v3 = (0 <= -2147483648)
                                    }
                                }
                                if len(v1) > 0 {
                                    v1[0] = 694
                                }
                                _ = fi28
                                _ = v29
                            }
                        }
                    }
                } else {
                    v31 := GenS1{F0: int(v6), F1: -128}
                    v5[-790] = "\n\t"
                    _ = v31
                }
            }
            var v32 float32 = (float32((1e10 + 1e10)) + ((1e10 + -2.25) + ((-2.25 * -2.25) + 3.14159)))
            switch "a" {
            case "a":
                if (2147483647 != 42) {
                    v33 := GenS3{F0: v6, F1: 2147483647}
                    var v34 int8 = 42
                    switch -1 {
                    case -552:
                        var v35 int8 = (v34 | v34)
                        _ = v35
                    case 795:
                        v6 = 18446744073709551615
                    }
                    _ = v33
                    _ = v34
                }
            case "":
                switch 246 {
                case -1:
                    for fi36 := 0; fi36 < 11; fi36++ {
                        for fi37 := 0; fi37 < 19; fi37++ {
                            if (v6 == 4294967295) {
                                v32 = (float32(0.0) - ((v32 - 0.1) / (v32 - v32)))
                                v2["foo bar"] = 1
                                v3 = (!(960 > -510))
                            }
                            v32 -= (float32(-2.25) + 1e10)
                            if v3 {
                                v38 := GenS4{F0: -128}
                                v39 := GenS4{F0: 127}
                                _ = v38
                                _ = v39
                            }
                            _ = fi37
                        }
                        v40 := GenS5{F0: "foo bar"}
                        _ = fi36
                        _ = v40
                    }
                case 9223372036854775807:
                    switch ("a" + "\n\t") {
                    case "a":
                        var v41 uint8 = 127
                        _ = v41
                    case "":
                        if ((591 == 4294967295) || (-667 <= 1000)) {
                            fmt.Println(v3, v3)
                            for fi42 := 0; fi42 < 19; fi42++ {
                                for fi43 := 0; fi43 < 20; fi43++ {
                                    v5[1] = "hello"
                                    var v44 uint32 = 1
                                    _ = fi43
                                    _ = v44
                                }
                                _ = fi42
                            }
                        }
                    case "foo bar":
                        v45 := GenS6{F0: v6, F1: 9223372036854775807}
                        _ = v45
                    }
                case 127:
                    v46 := GenS7{F0: uint16(v6), F1: (v3 && (v3 || false))}
                    _ = v46
                }
            }
            _ = v32
        }
    }
    var v47 int64 = 347
    switch 42 {
    case 4294967295:
        v47 = -9223372036854775808
    case 416:
        for fi48 := 0; fi48 < 9; fi48++ {
            v6 = v6
            if (427 != 1) {
                if true {
                    fmt.Println(v47, v3)
                    if (v3 && v3) {
                        for fi49 := 0; fi49 < 20; fi49++ {
                            var v50 int8 = -65
                            v51 := GenS5{F0: "a"}
                            _ = fi49
                            _ = v50
                            _ = v51
                        }
                    }
                    v3 = v3
                }
                switch (fi48 % (((fi48 >> ((255) & 63))) | 1)) {
                case 470:
                    v3 = (v47 != (^(v47)))
                default:
                    v3 = (false || (1 != 255))
                }
            } else {
                v52 := GenS3{F0: v6, F1: -2147483648}
                for fi53 := 0; fi53 < 5; fi53++ {
                    if (true && v3) {
                        v3 = v3
                        switch "a" {
                        case "":
                            v54 := GenS8{F0: 143, F1: int32(v47), F2: 4294967295}
                            _ = v54
                        case "hello":
                            v3 = false
                        default:
                            v55 := GenS3{F0: v6, F1: 614}
                            _ = v55
                        }
                        var v56 uint = 837
                        _ = v56
                    }
                    _ = fi53
                }
                v57 := GenS3{F0: 42, F1: -730}
                _ = v52
                _ = v57
            }
            _ = fi48
        }
    default:
        v58 := GenS9{F0: 255, F1: 255}
        _ = v58
    }
    if ((197 > 255) || true) {
        fmt.Println(v3, v3)
    } else {
        for fi59 := 0; fi59 < 1; fi59++ {
            for fi60 := 0; fi60 < 9; fi60++ {
                v3 = v3
                _ = fi60
            }
            for fi61 := 0; fi61 < 6; fi61++ {
                var v62 float32 = (1e10 * (0.1 + (1e10 * 1.5)))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi61
                _ = v62
            }
            _ = fi59
        }
        for fi63 := 0; fi63 < 5; fi63++ {
            switch fi63 {
            case -2147483648:
                v3 = (v6 >= (v6 - v6))
            case -166:
                for fi64 := 0; fi64 < 7; fi64++ {
                    _ = len(v1)
                    _ = fi64
                }
            }
            _ = fi63
        }
    }
    var v65 uint16 = 32767
    v66 := GenS10{F0: v65, F1: -45}
    var v67 bool = (255 > 128)
    v68 := GenS11{F0: 127}
    if (-2.25 <= ((1.5 + float64((1e10 / -2.25))) - -2.25)) {
        for fi69 := 0; fi69 < 12; fi69++ {
            v70 := GenS12{F0: uint32(v6)}
            var v71 float64 = 3.14159
            var v72 uint32 = 1
            _ = fi69
            _ = v70
            _ = v71
            _ = v72
        }
        delete(v4, "hello")
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v65
    _ = v66
    _ = v67
    _ = v68
}
