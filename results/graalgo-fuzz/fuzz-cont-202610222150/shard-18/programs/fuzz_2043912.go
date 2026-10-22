package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int32 }
type GenS2 struct { F0 int32; F1 int32; F2 string }
type GenS3 struct { F0 bool; F1 int; F2 uint32 }
type GenS4 struct { F0 uint8; F1 uint32; F2 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint; F1 int8; F2 int32 }

func main() {
    var v1 uint16 = 255
    var v2 uint16 = 282
    v3 := make(map[uint32]uint32)
    v4 := []string{"a", ""}
    v5 := GenS1{F0: 18446744073709551615, F1: int8(v2), F2: 0}
    for fi6 := 0; fi6 < 14; fi6++ {
        v1 = v1
        _ = fi6
    }
    fmt.Println(v2, v2)
    switch ("" + "foo bar") {
    case "foo bar":
        for fi7 := 0; fi7 < 5; fi7++ {
            if (1.5 == float32(0.1)) {
                var v8 int64 = 0
                var v9 int64 = (v8 & (v8 >> ((491) & 63)))
                delete(v3, 4294967295)
                _ = v8
                _ = v9
            } else {
                switch fi7 {
                case -277:
                    v10 := GenS2{F0: 31, F1: 1, F2: "\n\t"}
                    _ = v10
                }
            }
            _ = fi7
        }
    case "\n\t":
        var v11 uint32 = 503
        _ = v11
    default:
        fmt.Println(v1, v1)
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        switch "a" {
        case "foo bar":
            switch (fi12 / ((fi12) | 1)) {
            case 1:
                fmt.Println(v2, v1, v2)
            default:
                for fi13 := 0; fi13 < 6; fi13++ {
                    if (0.1 <= 3.14159) {
                        if true {
                            var v14 int32 = 1
                            var v15 uint8 = 128
                            var v16 int64 = -188
                            _ = v14
                            _ = v15
                            _ = v16
                        } else {
                            v1 = v2
                            v1 -= (^(v1))
                            fmt.Println(v2, fi12, v1)
                            var v17 int = fi12
                            if (255 <= 4294967295) {
                                if false {
                                    switch "hello" {
                                    case "foo bar":
                                        fmt.Println(v2, fi12)
                                    case "hello":
                                        var v18 float32 = 0.1
                                        v3[4294967295] = 2
                                        _ = v18
                                    case "\n\t":
                                        v19 := GenS1{F0: 0, F1: 0, F2: 1}
                                        _ = v19
                                    default:
                                        v3[509] = 1
                                    }
                                } else {
                                    v20 := GenS2{F0: 42, F1: -2147483648, F2: "\n\t"}
                                    if true {
                                        if true {
                                            if false {
                                                if (!true) {
                                                    fmt.Println(v1, v1)
                                                }
                                                v1 -= (^(v1))
                                            } else {
                                                var v21 int64 = -1
                                                v1 = 1
                                                _ = v21
                                            }
                                        } else {
                                            _ = len(v4)
                                            v1 -= (^(v2))
                                            var v22 string = "\n\t"
                                            _ = v22
                                        }
                                    } else {
                                        if (!true) {
                                            var v23 int8 = 100
                                            v23 = int8(v1)
                                            _ = v23
                                        } else {
                                            if (-450 != -603) {
                                                fmt.Println(v2, fi12, v2)
                                                v24 := GenS3{F0: true, F1: fi12, F2: 483}
                                                _ = v24
                                            } else {
                                                v3[688] = 2147483647
                                                if false {
                                                    if (!(65535 < 2147483648)) {
                                                        v2 = v2
                                                        var v25 bool = (false && ((1 >= 42) && true))
                                                        fmt.Println(v25, fi12)
                                                        _ = v25
                                                    } else {
                                                        fmt.Println(fi12, v17)
                                                        v1 -= v2
                                                    }
                                                    var v26 uint8 = 90
                                                    v27 := GenS4{F0: v26, F1: 380, F2: fi13}
                                                    _ = v26
                                                    _ = v27
                                                } else {
                                                    switch fi12 {
                                                    case 4294967295:
                                                        switch (fi13 + (fi12 << ((42) & 63))) {
                                                        case 1:
                                                            switch v17 {
                                                            case 32768:
                                                                if (false && (true || (!true))) {
                                                                    switch (v17 << ((42) & 63)) {
                                                                    case 9223372036854775807:
                                                                        v28 := GenS5{F0: 0}
                                                                        _ = v28
                                                                    default:
                                                                        switch int(v2) {
                                                                        case 0:
                                                                            if (1e10 != 1e10) {
                                                                                if false {
                                                                                    v1 -= v1
                                                                                } else {
                                                                                    v2 = (^(v2))
                                                                                }
                                                                                v2 = v1
                                                                            }
                                                                        case 394:
                                                                            var v29 int16 = -1
                                                                            v17 = (fi12 ^ (v17 | v17))
                                                                            var v30 bool = (!(1 == 4294967295))
                                                                            v31 := GenS3{F0: v30, F1: fi13, F2: 843}
                                                                            _ = v29
                                                                            _ = v30
                                                                            _ = v31
                                                                        default:
                                                                            v1 -= v1
                                                                        }
                                                                    }
                                                                } else {
                                                                    switch -184 {
                                                                    case 0:
                                                                        v32 := GenS3{F0: (0 != -1), F1: v17, F2: 4294967295}
                                                                        _ = v32
                                                                    }
                                                                    v17 += fi13
                                                                    if false {
                                                                        _ = len(v4)
                                                                        var v33 int16 = 42
                                                                        _ = v33
                                                                    } else {
                                                                        v34 := GenS3{F0: (!true), F1: (v17 % (((^(v17))) | 1)), F2: 2147483647}
                                                                        fmt.Println(v1, v17)
                                                                        var v35 uint = 4294967295
                                                                        _ = v34
                                                                        _ = v35
                                                                    }
                                                                }
                                                            case 127:
                                                                var v36 uint64 = 2
                                                                _ = v36
                                                            case 0:
                                                                v17 = v17
                                                            default:
                                                                var v37 float32 = 1e10
                                                                if (!(566 >= 2)) {
                                                                    v1 += (v2 << ((4294967295) & 63))
                                                                    var v38 int = fi12
                                                                    _ = v38
                                                                }
                                                                var v39 uint = 2
                                                                _ = v37
                                                                _ = v39
                                                            }
                                                        default:
                                                            fmt.Println(v2, fi13)
                                                        }
                                                    case -186:
                                                        v1 = v1
                                                    default:
                                                        var v40 int32 = 0
                                                        _ = v40
                                                    }
                                                    var v41 uint16 = 436
                                                    _ = v41
                                                }
                                            }
                                            v2 = (v2 >> ((255) & 63))
                                            if (1 > 1) {
                                                fmt.Println(v1)
                                                var v42 int16 = -444
                                                _ = v42
                                            } else {
                                                v4 = append(v4, "\n\t")
                                                _ = len(v4)
                                            }
                                        }
                                        v2 += (v1 * (v1 ^ v2))
                                        v17 += (fi13 & (fi13 | v17))
                                    }
                                    _ = v20
                                }
                                switch ("hello" + "\n\t") {
                                case "a":
                                    v3[0] = 4294967295
                                case "":
                                    fmt.Println(fi13)
                                case "hello":
                                    if (!false) {
                                        var v43 uint32 = 2147483647
                                        _ = len(v3)
                                        v4 = append(v4, "\n\t")
                                        _ = v43
                                    } else {
                                        v17 += fi13
                                        v44 := GenS6{F0: ""}
                                        _ = v44
                                    }
                                }
                                var v45 string = ("" + "foo bar")
                                _ = v45
                            } else {
                                var v46 float32 = 0.0
                                var v47 int16 = 1000
                                _ = v46
                                _ = v47
                            }
                            _ = v17
                        }
                        var v48 int32 = 0
                        _ = v48
                    } else {
                        fmt.Println(fi12)
                    }
                    v3[494] = 0
                    if false {
                        v49 := GenS5{F0: 1000}
                        _ = v49
                    } else {
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                    }
                    _ = fi13
                }
            }
        case "":
            fmt.Println(v2, v1, v2)
        case "a":
            v2 = v2
        default:
            switch "" {
            case "foo bar":
                v2 = 32767
            case "\n\t":
                fmt.Println(v1, fi12, v1)
            default:
                var v50 int8 = 1
                _ = v50
            }
        }
        v51 := GenS6{F0: "foo bar"}
        _ = fi12
        _ = v51
    }
    fmt.Println(v1, v1)
    fmt.Println(v1, v2)
    fmt.Println(v2, v2)
    if false {
        fmt.Println(v1, v1)
        var v52 int8 = 127
        v53 := GenS4{F0: 127, F1: 946, F2: 0}
        _ = v52
        _ = v53
    } else {
        v54 := GenS7{F0: 42, F1: 42, F2: 2147483647}
        v2 = 0
        var v55 uint16 = v2
        _ = v54
        _ = v55
    }
    if (128 > 43) {
        if false {
            v56 := v3[4294967295]
            _ = v56
        } else {
            var v57 float32 = 1.5
            fmt.Println(v2, v57, v1)
            _ = v57
        }
        var v58 string = "foo bar"
        if ((170 == 570) && (!(-2.25 > 0.0))) {
            v59 := GenS6{F0: v58}
            var v60 uint32 = 828
            _ = v59
            _ = v60
        } else {
            v1 = v2
        }
        _ = v58
    } else {
        for fi61 := 0; fi61 < 12; fi61++ {
            for fi62 := 0; fi62 < 18; fi62++ {
                v4 = append(v4, "\n\t")
                switch fi62 {
                case 4294967295:
                    var v63 bool = true
                    _ = v63
                case 128:
                    if len(v4) > 0 {
                        v4[0] = "\n\t"
                    }
                default:
                    if len(v4) > 0 {
                        v4[0] = ("\n\t" + "hello")
                    }
                }
                _ = fi62
            }
            var v64 int8 = 42
            _ = fi61
            _ = v64
        }
    }
    v2 = (^(v1))
    v1 = v2
    for fi65 := 0; fi65 < 18; fi65++ {
        switch int(v2) {
        case 0:
            fmt.Println(fi65, v1, v1)
        case 32767:
            var v66 float64 = (0.0 / (0.0 + 1e10))
            _ = v66
        }
        for fi67 := 0; fi67 < 2; fi67++ {
            var v68 uint64 = 484
            fmt.Println(v68, v2, v2)
            switch "foo bar" {
            case "a":
                fmt.Println(v1)
            case "foo bar":
                if ((-2.25 * ((3.14159 * 1.5) - 0.0)) > 0.0) {
                    v69 := GenS3{F0: (("hello" + ("foo bar" + "a")) < "\n\t"), F1: 0, F2: 212}
                    switch "foo bar" {
                    case "":
                        var v70 uint32 = 42
                        _ = v70
                    case "foo bar":
                        v2 -= (v2 - uint16(v68))
                    case "\n\t":
                        fmt.Println(v1, fi65)
                    default:
                        v1 += v2
                    }
                    _ = v69
                }
            default:
                var v71 uint16 = (v2 % ((1) | 1))
                _ = v71
            }
            _ = fi67
            _ = v68
        }
        _ = fi65
    }
    var v72 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v72
}
