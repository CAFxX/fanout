package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16 }

func main() {
    var v1 uint16 = 32767
    var v2 int8 = -128
    v3 := []uint32{62, 4294967295, 2147483648}
    if (!true) {
        v4 := GenS1{F0: -9223372036854775808, F1: 0}
        _ = v4
    } else {
        var v5 uint16 = (v1 | v1)
        var v6 uint32 = 430
        for fi7 := 0; fi7 < 14; fi7++ {
            v1 = (v1 * v5)
            _ = fi7
        }
        _ = v5
        _ = v6
    }
    fmt.Println(v2)
    if (4294967295 == 4294967295) {
        var v8 float32 = 0.1
        var v9 uint8 = 92
        var v10 int16 = 1000
        _ = v10
        _ = v8
        _ = v9
    } else {
        for fi11 := 0; fi11 < 4; fi11++ {
            v1 += (v1 + 761)
            var v12 uint = 839
            _ = fi11
            _ = v12
        }
    }
    var v13 int32 = 0
    v14 := GenS1{F0: -9223372036854775808, F1: 255}
    for fi15 := 0; fi15 < 13; fi15++ {
        v1 = v1
        if true {
            fmt.Println(v1, v2, v1)
            for fi16 := 0; fi16 < 4; fi16++ {
                switch "hello" {
                case "hello":
                    v3 = append(v3, 0)
                case "a":
                    if (442 <= 663) {
                        if len(v3) > 0 {
                            v3[0] = 596
                        }
                        if true {
                            var v17 uint = 4294967295
                            switch "hello" {
                            case "foo bar":
                                if (v17 > (^(v17))) {
                                    if (1.5 < -2.25) {
                                        var v18 float64 = (1.5 + 1.5)
                                        v18 += (1.5 * v18)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v18
                                    } else {
                                        switch -243 {
                                        case -1:
                                            var v19 uint16 = v1
                                            _ = v19
                                        case 1:
                                            v20 := GenS1{F0: (fi15 | (fi16 % (((fi15 & fi16)) | 1))), F1: v17}
                                            _ = v20
                                        case 127:
                                            var v21 float64 = 1e10
                                            _ = v21
                                        default:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        if (true || true) {
                                            var v22 int32 = (v13 | (v13 * (v13 >> ((v17) & 63))))
                                            switch ("hello" + ("" + "a")) {
                                            case "foo bar":
                                                var v23 uint64 = 1
                                                _ = v23
                                            case "":
                                                var v24 int64 = -9223372036854775808
                                                _ = v24
                                            default:
                                                var v25 uint = 42
                                                _ = v25
                                            }
                                            _ = v22
                                        } else {
                                            if (false && ((0.1 + float64(3.14159)) == (0.0 - (1e10 + 0.0)))) {
                                                v26 := GenS1{F0: 354, F1: 42}
                                                if (((1e10 * (0.1 * (3.14159 / 1e10))) * 0.0) != 0.0) {
                                                    var v27 int16 = 1
                                                    if false {
                                                        var v28 uint64 = 4294967295
                                                        _ = v28
                                                    } else {
                                                        v29 := GenS1{F0: fi16, F1: (v17 + v17)}
                                                        var v30 uint = v17
                                                        if (float64(3.14159) == 3.14159) {
                                                            v17 = v30
                                                            v17 = (^(v17))
                                                            if false {
                                                                if (1000 != 1) {
                                                                    if false {
                                                                        v31 := GenS2{F0: (v2 * (^(v2)))}
                                                                        _ = v31
                                                                    }
                                                                    fmt.Println(fi16)
                                                                }
                                                                var v32 int64 = 246
                                                                var v33 uint64 = 42
                                                                _ = v32
                                                                _ = v33
                                                            } else {
                                                                if (!true) {
                                                                    switch ("\n\t" + ("" + ("foo bar" + "\n\t"))) {
                                                                    case "a":
                                                                        if false {
                                                                            v17 -= v17
                                                                        }
                                                                    case "":
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        var v34 uint = 42
                                                                        _ = v34
                                                                    }
                                                                    var v35 int16 = v27
                                                                    if true {
                                                                        var v36 uint64 = 393
                                                                        _ = v36
                                                                    }
                                                                    _ = v35
                                                                }
                                                                if len(v3) > 0 {
                                                                    v3[0] = 2147483647
                                                                }
                                                                _ = len(v3)
                                                            }
                                                        }
                                                        v37 := GenS3{F0: 475}
                                                        _ = v29
                                                        _ = v30
                                                        _ = v37
                                                    }
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    v1 = 678
                                                    var v38 string = ("a" + "a")
                                                    _ = v27
                                                    _ = v38
                                                }
                                                _ = v26
                                            }
                                            v1 = v1
                                        }
                                        v1 -= uint16(v17)
                                    }
                                    v3 = append(v3, 250)
                                    var v39 string = "hello"
                                    _ = v39
                                } else {
                                    v13 -= (v13 >> (((^(v17))) & 63))
                                }
                            case "hello":
                                switch "foo bar" {
                                case "":
                                    v40 := GenS4{F0: -32768}
                                    _ = v40
                                case "a":
                                    if false {
                                        var v41 int = (^(fi16))
                                        v1 -= 0
                                        v1 = v1
                                        v1 += (v1 / ((v1) | 1))
                                        _ = v41
                                    } else {
                                        _ = len(v3)
                                    }
                                default:
                                    if (100 < -128) {
                                        if len(v3) > 0 {
                                            v3[0] = 1
                                        }
                                        v17 += 565
                                    } else {
                                        fmt.Println(v2, v2, fi16)
                                        var v42 int64 = 0
                                        if (!false) {
                                            v13 = v13
                                        }
                                        _ = v42
                                    }
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case "":
                                if (((fi15 & (fi15 | -9223372036854775808)) == (fi15 * (fi15 << ((880) & 63)))) || true) {
                                    fmt.Println(fi15, v17)
                                } else {
                                    v13 = 1
                                    v3 = append(v3, 42)
                                }
                            }
                            _ = v17
                        } else {
                            if (true || false) {
                                v1 = 32767
                            } else {
                                v13 = (v13 << ((1) & 63))
                            }
                            var v43 uint16 = v1
                            _ = v43
                        }
                    }
                case "\n\t":
                    fmt.Println(v1, fi15, fi16)
                default:
                    v1 = (v1 & v1)
                }
                if len(v3) > 0 {
                    v3[0] = 42
                }
                v1 = v1
                _ = fi16
            }
        } else {
            var v44 float64 = -2.25
            _ = v44
        }
        for fi45 := 0; fi45 < 8; fi45++ {
            if false {
                v1 += v1
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
            }
            _ = fi45
        }
        _ = fi15
    }
    v1 += 42
    var v46 float32 = 3.14159
    v2 += v2
    if false {
        switch "hello" {
        case "\n\t":
            v46 = 0.1
        case "":
            if true {
                var v47 int = 2147483647
                var v48 string = "\n\t"
                _ = v47
                _ = v48
            } else {
                if len(v3) > 0 {
                    v3[0] = 2147483648
                }
                v2 += v2
            }
        case "a":
            var v49 uint8 = 128
            _ = v49
        }
        v50 := GenS2{F0: (-(v2))}
        var v51 int32 = (v13 - -809)
        _ = v50
        _ = v51
    }
    var v52 uint = 866
    v13 = 2147483647
    switch "" {
    case "a":
        var v53 int8 = v2
        _ = v53
    case "foo bar":
        switch ("hello" + "a") {
        case "foo bar":
            var v54 float32 = v46
            _ = v54
        case "\n\t":
            var v55 float64 = 0.1
            _ = v55
        case "a":
            if (-6 >= -230) {
                var v56 int8 = -69
                _ = v56
            }
        default:
            for fi57 := 0; fi57 < 8; fi57++ {
                v52 = (v52 >> ((v52) & 63))
                v58 := GenS4{F0: 781}
                _ = fi57
                _ = v58
            }
        }
    default:
        v1 = (v1 << (((v52 - v52)) & 63))
    }
    var v59 int8 = 100
    v59 += v59
    switch "a" {
    case "":
        v1 += v1
    }
    _ = len(v3)
    var v60 uint = v52
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v46:", v46)
    fmt.Println("v52:", v52)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v46
    _ = v52
    _ = v59
    _ = v60
}
