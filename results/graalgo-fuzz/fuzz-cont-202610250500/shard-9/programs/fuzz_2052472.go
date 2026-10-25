package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int32 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 uint }

func main() {
    var v1 int32 = 42
    v2 := []uint32{1, 2147483648}
    var v3 int64 = -9223372036854775808
    v3 = (v3 & (v3 << ((2) & 63)))
    for fi4 := 0; fi4 < 8; fi4++ {
        if (!false) {
            var v5 uint16 = 467
            _ = v5
        } else {
            switch "\n\t" {
            case "hello":
                if (0.1 < -2.25) {
                    if (("a" != "foo bar") && ((v1 - (^(v1))) > v1)) {
                        _ = len(v2)
                        for fi6 := 0; fi6 < 2; fi6++ {
                            var v7 uint64 = 2
                            v8 := GenS1{F0: 255, F1: 0, F2: -308}
                            v3 = -1
                            fmt.Println(v7)
                            v1 -= v1
                            if len(v2) > 0 {
                                v2[0] = 42
                            }
                            _ = fi6
                            _ = v7
                            _ = v8
                        }
                        v9 := GenS2{F0: 65535, F1: -9223372036854775808}
                        _ = v9
                    } else {
                        var v10 int16 = 42
                        for fi11 := 0; fi11 < 8; fi11++ {
                            var v12 bool = (-1 < int8(fi11))
                            v12 = (!(!v12))
                            _ = fi11
                            _ = v12
                        }
                        if (0.0 != ((((((0.0 - (1e10 / 1.0)) * 1.5) + 3.14159) / (3.14159 * ((3.14159 * 1.5) * 0.1))) / (float32(-2.25) / 1.5)) - 0.0)) {
                            switch int(v1) {
                            case 1:
                                v13 := GenS1{F0: 4294967295, F1: 255, F2: v1}
                                _ = v13
                            case -1:
                                v14 := GenS3{F0: 432, F1: fi4}
                                _ = v14
                            default:
                                var v15 int8 = -128
                                _ = v15
                            }
                        } else {
                            var v16 uint16 = 65535
                            var v17 bool = true
                            _ = v16
                            _ = v17
                        }
                        _ = v10
                    }
                    fmt.Println(fi4, fi4)
                } else {
                    for fi18 := 0; fi18 < 10; fi18++ {
                        fmt.Println(fi18, v1)
                        v1 = -2147483648
                        v19 := GenS2{F0: 732, F1: v3}
                        switch "foo bar" {
                        case "foo bar":
                            var v20 float32 = float32(0.1)
                            _ = v20
                        case "":
                            var v21 int = 4294967295
                            _ = v21
                        default:
                            var v22 float64 = 0.1
                            _ = v22
                        }
                        _ = len(v2)
                        _ = fi18
                        _ = v19
                    }
                }
            }
        }
        if (1 > 4294967295) {
            var v23 uint32 = 2147483647
            switch "" {
            case "":
                for fi24 := 0; fi24 < 14; fi24++ {
                    v23 = 4294967295
                    v25 := GenS1{F0: 4294967295, F1: 2, F2: 2147483647}
                    _ = fi24
                    _ = v25
                }
            case "hello":
                switch "\n\t" {
                case "":
                    var v26 uint16 = 347
                    _ = v26
                case "\n\t":
                    v27 := GenS4{F0: 153, F1: (v3 * v3)}
                    _ = v27
                default:
                    v28 := GenS4{F0: 0, F1: v3}
                    _ = v28
                }
            case "a":
                v1 = v1
            default:
                switch ("hello" + "") {
                case "\n\t":
                    var v29 float64 = -2.25
                    _ = v29
                }
            }
            for fi30 := 0; fi30 < 11; fi30++ {
                fmt.Println(v23)
                if (false || (!true)) {
                    v1 += v1
                    if len(v2) > 0 {
                        v2[0] = (v23 / (((v23 * (v23 & (v23 ^ (v23 - (v23 ^ (v23 / ((0) | 1)))))))) | 1))
                    }
                } else {
                    var v31 uint16 = 803
                    var v32 uint64 = 2
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v31
                    _ = v32
                }
                _ = fi30
            }
            _ = v23
        } else {
            fmt.Println(v1, v3, v1)
            var v33 uint64 = 2
            if (v1 != v1) {
                var v34 uint = 536
                v35 := GenS3{F0: v33, F1: (fi4 << ((4294967295) & 63))}
                _ = v34
                _ = v35
            }
            _ = v33
        }
        _ = fi4
    }
    v3 = (v3 ^ v3)
    v36 := GenS2{F0: 408, F1: (v3 % (((v3 & (v3 | v3))) | 1))}
    var v37 int16 = 32767
    switch "\n\t" {
    case "":
        switch "a" {
        case "a":
            _ = len(v2)
        case "hello":
            var v38 float64 = (3.14159 + 1e10)
            _ = v38
        case "":
            switch int(v3) {
            case 202:
                v39 := GenS4{F0: 304, F1: v3}
                _ = v39
            case 4294967295:
                v40 := GenS2{F0: 0, F1: v3}
                _ = v40
            default:
                v2 = append(v2, 825)
            }
        }
    default:
        var v41 bool = (63 < 2)
        _ = v41
    }
    v1 += (-(v1))
    for fi42 := 0; fi42 < 12; fi42++ {
        fmt.Println(v1, v3, v37)
        _ = fi42
    }
    v1 = (v1 / ((v1) | 1))
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v43 bool = (!(!(!true)))
        _ = v43
    } else {
        for fi44 := 0; fi44 < 19; fi44++ {
            for fi45 := 0; fi45 < 6; fi45++ {
                var v46 string = (("" + "foo bar") + "hello")
                switch fi45 {
                case 332:
                    v47 := GenS4{F0: 18446744073709551615, F1: (v3 << ((2) & 63))}
                    _ = v47
                case -103:
                    v46 = v46
                    if (1 != int16(v3)) {
                        if len(v2) > 0 {
                            v2[0] = 2
                        }
                    }
                }
                _ = fi45
                _ = v46
            }
            fmt.Println(v3)
            var v48 int32 = (v1 * (v1 >> ((18446744073709551615) & 63)))
            _ = fi44
            _ = v48
        }
        if len(v2) > 0 {
            v2[0] = 311
        }
        for fi49 := 0; fi49 < 9; fi49++ {
            if true {
                fmt.Println(v3)
                v37 -= -32768
                for fi50 := 0; fi50 < 14; fi50++ {
                    var v51 int8 = 0
                    switch fi50 {
                    case 648:
                        v37 = (^(v37))
                    case 32768:
                        if true {
                            var v52 float64 = 1e10
                            _ = v52
                        }
                    }
                    _ = fi50
                    _ = v51
                }
            }
            v53 := GenS5{F0: 0}
            var v54 bool = true
            _ = fi49
            _ = v53
            _ = v54
        }
    }
    v37 = (v37 << ((29) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
}
