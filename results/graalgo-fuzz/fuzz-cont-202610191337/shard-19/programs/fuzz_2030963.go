package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint; F1 int; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 int16; F2 int16 }
type GenS5 struct { F0 uint; F1 uint16; F2 int64 }

func main() {
    v1 := []uint32{2147483648, 2147483648, 1, 2}
    var v2 uint32 = 65535
    var v3 int8 = 42
    fmt.Println(v3, v2)
    var v4 string = ""
    if len(v1) > 0 {
        v1[0] = v2
    }
    v5 := GenS1{F0: v2, F1: 2, F2: -1}
    var v6 int8 = (v3 & v3)
    fmt.Println(v3)
    if (42 != 1) {
        v2 = (v2 % (((v2 | 420)) | 1))
        if false {
            _ = len(v1)
        } else {
            for fi7 := 0; fi7 < 19; fi7++ {
                for fi8 := 0; fi8 < 12; fi8++ {
                    if (!true) {
                        var v9 int = fi8
                        _ = v9
                    }
                    _ = fi8
                }
                _ = fi7
            }
        }
    } else {
        v6 -= 52
        var v10 uint8 = 245
        for fi11 := 0; fi11 < 2; fi11++ {
            for fi12 := 0; fi12 < 8; fi12++ {
                v13 := GenS1{F0: v2, F1: 2, F2: 295}
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
        _ = v10
    }
    switch 42 {
    case 9223372036854775807:
        v3 -= v3
    case -923:
        if false {
            switch "foo bar" {
            case "a":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "":
                v1 = append(v1, 4294967295)
            }
            if len(v1) > 0 {
                v1[0] = v2
            }
        } else {
            var v14 uint64 = 18446744073709551615
            var v15 int16 = 0
            _ = v14
            _ = v15
        }
    }
    v2 = 1
    switch -1 {
    case 32767:
        if (0 > 32768) {
            if false {
                if len(v1) > 0 {
                    v1[0] = (^(v2))
                }
                for fi16 := 0; fi16 < 9; fi16++ {
                    var v17 uint = 381
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    for fi18 := 0; fi18 < 14; fi18++ {
                        v19 := GenS1{F0: v2, F1: 1, F2: 1}
                        v1 = append(v1, v2)
                        if false {
                            if true {
                                if true {
                                    v1 = append(v1, (v2 & uint32(v17)))
                                    var v20 string = "\n\t"
                                    _ = v20
                                }
                                _ = len(v1)
                                var v21 int64 = -9223372036854775808
                                _ = v21
                            } else {
                                switch "" {
                                case "a":
                                    switch v4 {
                                    case "\n\t":
                                        var v22 float32 = -2.25
                                        _ = v22
                                    }
                                case "\n\t":
                                    if ((v3 << ((v17) & 63)) < (-(v3))) {
                                        v6 -= 0
                                        fmt.Println(v17)
                                        var v23 bool = ((!true) && true)
                                        _ = v23
                                    }
                                }
                                switch (fi18 & fi18) {
                                case -9223372036854775808:
                                    switch fi16 {
                                    case 32767:
                                        switch (fi18 >> ((2) & 63)) {
                                        case 32767:
                                            var v24 bool = (false && ((true || (false || (-1 < 184))) || true))
                                            if v24 {
                                                var v25 uint8 = 22
                                                switch (fi18 ^ 1) {
                                                case 32768:
                                                    switch "foo bar" {
                                                    case "hello":
                                                        v26 := GenS1{F0: v2, F1: 376, F2: 2147483647}
                                                        _ = v26
                                                    case "a":
                                                        var v27 int8 = v3
                                                        _ = v27
                                                    default:
                                                        var v28 float64 = -2.25
                                                        _ = v28
                                                    }
                                                case -165:
                                                    var v29 string = v4
                                                    _ = v29
                                                case 438:
                                                    v24 = v24
                                                }
                                                _ = v25
                                            }
                                            switch (((v4 + ((v4 + (("hello" + ((v4 + "") + "\n\t")) + v4)) + v4)) + v4) + (v4 + "a")) {
                                            case "\n\t":
                                                v30 := GenS2{F0: (v17 & v17), F1: fi16, F2: -428}
                                                _ = v30
                                            }
                                            _ = v24
                                        case 104:
                                            if (!false) {
                                                var v31 int = fi16
                                                switch (v4 + v4) {
                                                case "hello":
                                                    fmt.Println(v31, v2)
                                                default:
                                                    if true {
                                                        var v32 int64 = 1
                                                        if (3.14159 != (((1.5 * 1.5) - 1e10) / 1.0)) {
                                                            v33 := GenS2{F0: 328, F1: (v31 + (fi16 >> ((v17) & 63))), F2: int64(fi16)}
                                                            v6 = 1
                                                            v6 += (v6 | 100)
                                                            _ = v33
                                                        }
                                                        v2 = (^(v2))
                                                        _ = v32
                                                    } else {
                                                        fmt.Println(fi18, fi18)
                                                        var v34 int8 = v6
                                                        if false {
                                                            var v35 float32 = (-2.25 + 1e10)
                                                            v2 -= (v2 | v2)
                                                            _ = v35
                                                        }
                                                        _ = v34
                                                    }
                                                }
                                                v36 := GenS3{F0: -9}
                                                _ = v31
                                                _ = v36
                                            }
                                        }
                                    default:
                                        v3 -= 20
                                    }
                                case 437:
                                    v37 := GenS3{F0: v3}
                                    _ = v37
                                }
                            }
                            v4 = v4
                            v6 = (^(v6))
                        } else {
                            v3 = -84
                            v3 = (v3 ^ 1)
                            fmt.Println(v6, fi18)
                        }
                        v3 += v6
                        if len(v1) > 0 {
                            v1[0] = (v2 | v2)
                        }
                        var v38 string = (v4 + "")
                        _ = fi18
                        _ = v19
                        _ = v38
                    }
                    _ = fi16
                    _ = v17
                }
                for fi39 := 0; fi39 < 13; fi39++ {
                    v40 := GenS2{F0: 4294967295, F1: -942, F2: 42}
                    if false {
                        v2 -= 42
                    }
                    _ = fi39
                    _ = v40
                }
            } else {
                if true {
                    v6 -= -1
                }
                switch (((((("\n\t" + "\n\t") + (v4 + "")) + "a") + "foo bar") + "hello") + ("\n\t" + v4)) {
                case "hello":
                    v41 := GenS1{F0: (v2 % ((v2) | 1)), F1: 2, F2: -301}
                    _ = v41
                }
            }
            for fi42 := 0; fi42 < 9; fi42++ {
                v4 = "a"
                for fi43 := 0; fi43 < 7; fi43++ {
                    fmt.Println(v2, v3, v4)
                    _ = fi43
                }
                _ = len(v1)
                _ = fi42
            }
        }
    default:
        var v44 int64 = -9223372036854775808
        _ = v44
    }
    switch "\n\t" {
    case "foo bar":
        switch 927 {
        case 2147483647:
            if false {
                v45 := GenS4{F0: (v6 << ((2) & 63)), F1: -32768, F2: -1}
                _ = v45
            } else {
                var v46 uint16 = 0
                switch "a" {
                case "hello":
                    var v47 uint32 = 2147483647
                    _ = v47
                default:
                    var v48 int64 = 0
                    _ = v48
                }
                var v49 uint = 42
                _ = v46
                _ = v49
            }
        case 127:
            if (!true) {
                var v50 int16 = 0
                var v51 int64 = 740
                v52 := GenS5{F0: 1, F1: 2, F2: (v51 / ((v51) | 1))}
                _ = v50
                _ = v51
                _ = v52
            } else {
                if len(v1) > 0 {
                    v1[0] = 383
                }
                switch v4 {
                case "a":
                    v53 := GenS3{F0: v6}
                    _ = v53
                }
                for fi54 := 0; fi54 < 20; fi54++ {
                    v1 = append(v1, v2)
                    _ = fi54
                }
            }
        }
    case "a":
        var v55 bool = false
        _ = v55
    case "hello":
        var v56 bool = (1000 > -769)
        _ = v56
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
