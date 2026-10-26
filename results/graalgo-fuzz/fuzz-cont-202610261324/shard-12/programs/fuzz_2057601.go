package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int64; F2 int64 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS5 struct { F0 string; F1 uint32; F2 string }
type GenS6 struct { F0 uint16; F1 int32 }
type GenS7 struct { F0 uint8; F1 uint8 }
type GenS8 struct { F0 int64; F1 int }
type GenS9 struct { F0 int; F1 int16 }

func main() {
    v1 := []int{128}
    v2 := make(map[uint32]bool)
    var v3 int32 = 0
    var v4 int64 = -308
    var v5 string = "hello"
    v6 := make(map[string]bool)
    v6[v5] = true
    if true {
        for fi7 := 0; fi7 < 14; fi7++ {
            switch (fi7 ^ 128) {
            case 1:
                v8 := v2[469]
                _ = v8
            case 128:
                var v9 uint64 = 352
                _ = v9
            case 127:
                var v10 bool = false
                _ = v10
            default:
                var v11 int8 = 96
                _ = v11
            }
            var v12 uint = 0
            v12 += (v12 | (v12 % ((4294967295) | 1)))
            _ = fi7
            _ = v12
        }
        var v13 int = 0
        _ = v13
    } else {
        for fi14 := 0; fi14 < 13; fi14++ {
            v15 := GenS1{F0: v4}
            _ = fi14
            _ = v15
        }
        fmt.Println(v5, v4, v5)
        for fi16 := 0; fi16 < 6; fi16++ {
            v3 -= v3
            for fi17 := 0; fi17 < 10; fi17++ {
                var v18 uint32 = 65535
                var v19 bool = (!(true || true))
                var v20 bool = true
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            var v21 bool = (((137 >= 0) || (!true)) && ((true && true) && false))
            _ = fi16
            _ = v21
        }
    }
    v22 := GenS2{F0: -344, F1: (v4 + 288), F2: v4}
    for fi23 := 0; fi23 < 10; fi23++ {
        switch "foo bar" {
        case "hello":
            v24 := GenS1{F0: (v4 + v4)}
            _ = v24
        case "\n\t":
            if false {
                var v25 string = v5
                _ = v25
            }
        }
        if false {
            v4 = (v4 | v4)
            _ = len(v6)
            var v26 int32 = v3
            _ = v26
        }
        _ = fi23
    }
    v27 := GenS1{F0: (v4 / ((1) | 1))}
    for fi28 := 0; fi28 < 7; fi28++ {
        if (100 == 127) {
            if true {
                if true {
                    var v29 int16 = -32768
                    v30 := GenS3{F0: v5, F1: (v5 + v5)}
                    _ = v29
                    _ = v30
                }
                var v31 uint8 = 255
                if false {
                    v32 := GenS4{F0: 65535, F1: 117, F2: 2147483647}
                    if true {
                        v31 = 31
                        for fi33 := 0; fi33 < 10; fi33++ {
                            if false {
                                v31 = 2
                                fmt.Println(fi33)
                                switch v5 {
                                case "\n\t":
                                    if (1 < 0) {
                                        if true {
                                            if true {
                                                if (243 < (v31 & (v31 % ((0) | 1)))) {
                                                    var v34 uint8 = v31
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    var v35 int32 = -318
                                                    _ = v34
                                                    _ = v35
                                                }
                                                v5 = ((v5 + (("hello" + v5) + (v5 + v5))) + (v5 + "a"))
                                            }
                                        }
                                    }
                                case "":
                                    v6[v5] = true
                                case "hello":
                                    v36 := GenS5{F0: "", F1: 0, F2: v5}
                                    _ = v36
                                default:
                                    v31 = v31
                                }
                            }
                            var v37 int16 = 0
                            var v38 uint32 = 585
                            var v39 int = int(v4)
                            var v40 int32 = v3
                            _ = fi33
                            _ = v37
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                        v1 = append(v1, 114)
                    }
                    if (65535 != 18446744073709551615) {
                        for fi41 := 0; fi41 < 6; fi41++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = len(v1)
                            _ = fi41
                        }
                        fmt.Println(v31, v5, v31)
                        if (true && false) {
                            if len(v1) > 0 {
                                v1[0] = 32767
                            }
                            if false {
                                var v42 uint32 = 2147483647
                                v43 := GenS6{F0: 2, F1: (v3 / (((v3 >> ((65535) & 63))) | 1))}
                                _ = v42
                                _ = v43
                            } else {
                                v5 = "foo bar"
                                var v44 uint16 = 923
                                _ = v44
                            }
                            for fi45 := 0; fi45 < 10; fi45++ {
                                switch (fi45 + fi28) {
                                case 128:
                                    fmt.Println(v4, v3, v3)
                                }
                                delete(v6, v5)
                                _ = fi45
                            }
                        }
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(v3)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    _ = v32
                } else {
                    switch fi28 {
                    case -9223372036854775808:
                        var v46 int8 = 30
                        _ = v46
                    default:
                        if (false && ((42 < 933) || true)) {
                            fmt.Println(v5, v3, v5)
                            var v47 int8 = -128
                            switch fi28 {
                            case 503:
                                fmt.Println(fi28, v5)
                            case 32768:
                                fmt.Println(v4, v5)
                            case 127:
                                var v48 int64 = (v4 / (((v4 % ((v4) | 1))) | 1))
                                _ = v48
                            }
                            _ = v47
                        }
                    }
                    v5 = v5
                }
                _ = v31
            }
            for fi49 := 0; fi49 < 8; fi49++ {
                v50 := GenS7{F0: 128, F1: 24}
                switch "\n\t" {
                case "":
                    var v51 int = fi28
                    _ = v51
                case "foo bar":
                    v1 = append(v1, fi49)
                default:
                    v52 := v2[65535]
                    _ = v52
                }
                if (1000 < 980) {
                    v5 = v5
                    switch v5 {
                    case "foo bar":
                        _ = len(v6)
                        switch v5 {
                        case "\n\t":
                            _ = len(v2)
                        case "foo bar":
                            fmt.Println(fi49)
                        default:
                            if true {
                                var v53 int8 = 1
                                if (!(!false)) {
                                    var v54 string = "a"
                                    v53 = (v53 << ((397) & 63))
                                    v55 := GenS1{F0: (v4 + 733)}
                                    var v56 int32 = (v3 & v3)
                                    var v57 uint8 = 191
                                    _ = v54
                                    _ = v55
                                    _ = v56
                                    _ = v57
                                }
                                v2[2147483647] = (247 > 255)
                                _ = v53
                            } else {
                                var v58 uint = 65535
                                fmt.Println(fi28, v4)
                                var v59 int32 = v3
                                _ = v58
                                _ = v59
                            }
                        }
                        var v60 int32 = -2147483648
                        _ = v60
                    case "\n\t":
                        var v61 int = (fi28 << ((0) & 63))
                        _ = v61
                    case "hello":
                        if false {
                            v4 = (v4 % ((v4) | 1))
                            var v62 int = fi49
                            var v63 int32 = 212
                            v3 = -2147483648
                            _ = v62
                            _ = v63
                        } else {
                            _ = len(v1)
                            v4 = (v4 * (v4 % ((v4) | 1)))
                        }
                    }
                }
                fmt.Println(v5, fi28, v4)
                var v64 bool = ((fi28 & (^(fi28))) < (-(fi28)))
                v65 := GenS8{F0: (-(v4)), F1: (fi49 + -129)}
                _ = fi49
                _ = v50
                _ = v64
                _ = v65
            }
            v66 := GenS9{F0: (fi28 & fi28), F1: -1}
            _ = v66
        } else {
            v67 := GenS5{F0: v5, F1: 2, F2: ((((v5 + ("foo bar" + "hello")) + "") + "hello") + v5)}
            if ((2 == 32) || true) {
                v68 := GenS6{F0: 42, F1: -419}
                var v69 int16 = -32768
                v4 += (-(v4))
                _ = v68
                _ = v69
            } else {
                v5 = v5
            }
            if ((v3 % ((v3) | 1)) != (v3 << ((4294967295) & 63))) {
                v5 = v5
            } else {
                v5 = "foo bar"
            }
            _ = v67
        }
        var v70 float32 = 0.1
        _ = fi28
        _ = v70
    }
    var v71 float64 = -2.25
    v5 = v5
    for fi72 := 0; fi72 < 13; fi72++ {
        v3 += (v3 * (v3 << ((65535) & 63)))
        _ = fi72
    }
    v1 = append(v1, 32768)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v71
}
