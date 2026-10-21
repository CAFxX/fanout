package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 int16; F2 bool }
type GenS5 struct { F0 int8; F1 uint32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16; F1 int32 }
type GenS8 struct { F0 int32; F1 bool }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int64 }

func main() {
    var v1 uint8 = 227
    var v2 uint = 18446744073709551615
    var v3 bool = false
    v4 := []int{32768}
    for fi5 := 0; fi5 < 1; fi5++ {
        for fi6 := 0; fi6 < 14; fi6++ {
            fmt.Println(v3, v3, v3)
            _ = fi6
        }
        v2 += uint(v1)
        _ = fi5
    }
    v7 := GenS1{F0: 0}
    for fi8 := 0; fi8 < 7; fi8++ {
        v3 = (!v3)
        for fi9 := 0; fi9 < 20; fi9++ {
            if v3 {
                switch "" {
                case "foo bar":
                    v1 = v1
                case "":
                    v3 = (uint(v1) >= (^(v2)))
                case "hello":
                    var v10 int32 = 0
                    v11 := GenS2{F0: (!false), F1: 42, F2: uint64(v2)}
                    _ = v10
                    _ = v11
                }
                if (v3 && false) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                if v3 {
                    var v12 int16 = -256
                    _ = v12
                }
                var v13 uint = (v2 + (v2 / ((v2) | 1)))
                var v14 uint16 = 42
                v15 := GenS3{F0: 43}
                if v3 {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v16 := GenS4{F0: 1, F1: 1000, F2: (866 != fi9)}
                    _ = v16
                }
                if false {
                    fmt.Println(fi8)
                } else {
                    fmt.Println(v3, v14)
                    if ((v1 ^ (v1 / ((v1) | 1))) >= (v1 + v1)) {
                        var v17 uint8 = 225
                        if v3 {
                            switch ("\n\t" + "\n\t") {
                            case "\n\t":
                                if ((v3 && (!v3)) || (0 == 100)) {
                                    var v18 int16 = 32767
                                    if true {
                                        var v19 int32 = -649
                                        if len(v4) > 0 {
                                            v4[0] = -9223372036854775808
                                        }
                                        v4 = append(v4, fi8)
                                        if (127 >= 72) {
                                            v20 := GenS2{F0: (((v19 >> (((v2 ^ v13)) & 63)) > (^(v19))) && (v3 && false)), F1: 0, F2: 334}
                                            v1 = 255
                                            fmt.Println(fi9, v17)
                                            v18 -= -338
                                            _ = v20
                                        }
                                        if v3 {
                                            v4 = append(v4, (fi8 >> ((2) & 63)))
                                        } else {
                                            switch "hello" {
                                            case "\n\t":
                                                if v3 {
                                                    var v21 uint8 = v1
                                                    var v22 int64 = -1
                                                    v13 += (^(v13))
                                                    _ = v21
                                                    _ = v22
                                                }
                                            case "hello":
                                                v23 := GenS5{F0: int8(v18), F1: 65535}
                                                _ = v23
                                            case "":
                                                var v24 int8 = 1
                                                _ = v24
                                            default:
                                                if v3 {
                                                    v25 := GenS4{F0: v19, F1: v18, F2: (v19 > v19)}
                                                    _ = v25
                                                }
                                                switch (fi9 | (^(fi9))) {
                                                case 4294967295:
                                                    v13 = (v13 % ((2) | 1))
                                                case 32768:
                                                    var v26 bool = (0 > 42)
                                                    _ = v26
                                                case 871:
                                                    if (128 == (v17 - 54)) {
                                                        var v27 uint = v2
                                                        var v28 int8 = 0
                                                        v29 := GenS6{F0: v17}
                                                        if false {
                                                            v17 += (v17 % (((v17 & 1)) | 1))
                                                            if (v28 <= v28) {
                                                                v30 := GenS1{F0: v14}
                                                                v31 := GenS5{F0: 1, F1: 0}
                                                                var v32 uint16 = 42
                                                                _ = v30
                                                                _ = v31
                                                                _ = v32
                                                            }
                                                            if len(v4) > 0 {
                                                                v4[0] = 9223372036854775807
                                                            }
                                                        }
                                                        fmt.Println(v17, v17, v14)
                                                        var v33 uint64 = 2
                                                        fmt.Println(v14, v33)
                                                        _ = v27
                                                        _ = v28
                                                        _ = v29
                                                        _ = v33
                                                    }
                                                default:
                                                    var v34 float32 = 3.14159
                                                    _ = v34
                                                }
                                                var v35 uint64 = 1
                                                _ = v35
                                            }
                                            var v36 bool = (int8(v19) < 22)
                                            if len(v4) > 0 {
                                                v4[0] = 492
                                            }
                                            if (!true) {
                                                var v37 uint8 = 128
                                                _ = v37
                                            } else {
                                                _ = len(v4)
                                            }
                                            if v36 {
                                                v38 := GenS7{F0: v18, F1: (v19 ^ 1)}
                                                switch "hello" {
                                                case "hello":
                                                    var v39 uint64 = 2
                                                    var v40 int16 = 1
                                                    _ = v39
                                                    _ = v40
                                                case "":
                                                    fmt.Println(fi9)
                                                }
                                                _ = v38
                                            } else {
                                                var v41 uint16 = (v14 & v14)
                                                v13 = v13
                                                v41 -= uint16(v17)
                                                v42 := GenS3{F0: 127}
                                                switch (fi8 / (((fi9 >> ((uint(v1)) & 63))) | 1)) {
                                                case 32767:
                                                    v14 += (v14 - (v41 & (^(v14))))
                                                default:
                                                    if v36 {
                                                        var v43 int = (fi8 << ((v2) & 63))
                                                        fmt.Println(v3, v13, v41)
                                                        _ = v43
                                                    } else {
                                                        v41 -= v41
                                                        v18 = 42
                                                        v19 += (v19 + v19)
                                                    }
                                                }
                                                _ = v41
                                                _ = v42
                                            }
                                            _ = v36
                                        }
                                        var v44 float64 = (1e10 + 0.0)
                                        _ = v19
                                        _ = v44
                                    } else {
                                        switch -1 {
                                        case 377:
                                            fmt.Println(fi9, v13, v13)
                                        case -468:
                                            var v45 float32 = 0.1
                                            _ = v45
                                        case 42:
                                            _ = len(v4)
                                        default:
                                            var v46 float64 = 1e10
                                            _ = v46
                                        }
                                        v17 += v17
                                        var v47 bool = (false || (1 > (^(v2))))
                                        _ = v47
                                    }
                                    var v48 uint64 = 1
                                    _ = v18
                                    _ = v48
                                }
                            case "foo bar":
                                fmt.Println(v1, fi9)
                            default:
                                switch (("" + "hello") + "foo bar") {
                                case "foo bar":
                                    fmt.Println(v1, v3)
                                case "a":
                                    fmt.Println(v2)
                                    fmt.Println(fi9, fi8)
                                default:
                                    if false {
                                        var v49 string = (((("" + ("foo bar" + "")) + "\n\t") + ((("foo bar" + (("hello" + ("\n\t" + "hello")) + ("" + ((("\n\t" + ("a" + ("\n\t" + ""))) + "hello") + "foo bar")))) + (((("\n\t" + ("foo bar" + "\n\t")) + ("" + "")) + ("hello" + "\n\t")) + ("foo bar" + ""))) + "hello")) + "foo bar")
                                        _ = v49
                                    }
                                }
                            }
                            fmt.Println(v2)
                            fmt.Println(v13, v3)
                        } else {
                            if false {
                                switch "\n\t" {
                                case "a":
                                    var v50 int32 = 1
                                    if true {
                                        v17 = 127
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    _ = v50
                                case "":
                                    switch "foo bar" {
                                    case "":
                                        var v51 uint8 = 2
                                        _ = v51
                                    case "foo bar":
                                        v14 = uint16(v2)
                                    default:
                                        var v52 int64 = 180
                                        _ = v52
                                    }
                                case "foo bar":
                                    switch "a" {
                                    case "\n\t":
                                        v3 = (0.0 != float32(-2.25))
                                    case "":
                                        v3 = v3
                                    case "foo bar":
                                        fmt.Println(v1)
                                    default:
                                        var v53 int32 = 114
                                        _ = v53
                                    }
                                }
                                v54 := GenS7{F0: 32767, F1: int32(v1)}
                                _ = v54
                            } else {
                                v17 += uint8(v14)
                                if len(v4) > 0 {
                                    v4[0] = int(v17)
                                }
                                var v55 uint8 = (v1 * v1)
                                _ = v55
                            }
                            if (-917 <= -1) {
                                fmt.Println(v2, v17)
                                if (42 != 42) {
                                    v56 := GenS6{F0: (v17 & (v17 * v1))}
                                    v17 = v1
                                    var v57 int32 = 42
                                    _ = v56
                                    _ = v57
                                }
                            } else {
                                v58 := GenS8{F0: 2147483647, F1: false}
                                switch -24 {
                                case 368:
                                    if true {
                                        _ = len(v4)
                                        v14 = (^(v14))
                                        var v59 uint64 = 42
                                        _ = v59
                                    }
                                default:
                                    if (!v3) {
                                        switch "" {
                                        case "hello":
                                            switch ("" + "") {
                                            case "hello":
                                                v1 = 1
                                            case "":
                                                if v3 {
                                                    var v60 bool = v3
                                                    _ = v60
                                                }
                                            default:
                                                if v3 {
                                                    fmt.Println(v14)
                                                }
                                            }
                                        default:
                                            fmt.Println(v17, fi9)
                                        }
                                    } else {
                                        v61 := GenS9{F0: 9223372036854775807}
                                        _ = v61
                                    }
                                }
                                _ = v58
                            }
                            fmt.Println(v2)
                        }
                        var v62 int16 = -32768
                        _ = v17
                        _ = v62
                    }
                }
                _ = v13
                _ = v14
                _ = v15
            }
            v2 += (^(v2))
            _ = fi9
        }
        v2 -= (v2 | v2)
        _ = fi8
    }
    var v63 string = "\n\t"
    v3 = true
    var v64 int64 = 42
    v4 = append(v4, 1)
    for fi65 := 0; fi65 < 8; fi65++ {
        if len(v4) > 0 {
            v4[0] = fi65
        }
        v66 := GenS10{F0: ((v63 + "a") + ("foo bar" + v63))}
        if v3 {
            if (true && (!((v63 + ("foo bar" + "a")) <= v63))) {
                v64 -= v64
            } else {
                switch v63 {
                case "":
                    if len(v4) > 0 {
                        v4[0] = (-(fi65))
                    }
                case "\n\t":
                    for fi67 := 0; fi67 < 2; fi67++ {
                        v68 := GenS2{F0: ("foo bar" <= (((v63 + "foo bar") + v63) + "")), F1: -35, F2: 1}
                        _ = fi67
                        _ = v68
                    }
                default:
                    var v69 int16 = 1000
                    _ = v69
                }
            }
            v64 = v64
        }
        _ = fi65
        _ = v66
    }
    if v3 {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = len(v4)
        var v70 uint64 = 42
        _ = v70
    } else {
        v71 := GenS11{F0: v64}
        for fi72 := 0; fi72 < 3; fi72++ {
            v64 = v64
            var v73 int8 = 42
            v63 = v63
            _ = fi72
            _ = v73
        }
        _ = v71
    }
    var v74 int64 = v64
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v63
    _ = v64
    _ = v7
    _ = v74
}
