package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int32; F1 bool; F2 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8; F1 uint }
type GenS6 struct { F0 int64; F1 int32 }
type GenS7 struct { F0 uint64; F1 uint }
type GenS8 struct { F0 int64; F1 uint8; F2 uint32 }

func main() {
    var v1 int64 = 0
    var v2 uint64 = 210
    var v3 int8 = 42
    var v4 string = (("\n\t" + "hello") + "\n\t")
    var v5 int64 = 319
    v6 := []string{(("hello" + v4) + (v4 + (("a" + v4) + "foo bar"))), ("foo bar" + v4)}
    var v7 int8 = v3
    if false {
        _ = len(v6)
        if (((96 <= 2) && (false || true)) && (!false)) {
            if (!true) {
                switch 9223372036854775807 {
                case 4294967295:
                    switch "hello" {
                    case "a":
                        switch 128 {
                        case 14:
                            switch 1 {
                            case -1:
                                fmt.Println(v1)
                            }
                        }
                    case "hello":
                        v4 = v4
                    case "":
                        if false {
                            var v8 int32 = int32(v3)
                            v3 = (v7 ^ v7)
                            var v9 float32 = 0.0
                            _ = v8
                            _ = v9
                        } else {
                            v4 = (("\n\t" + ((("foo bar" + "foo bar") + v4) + "")) + v4)
                            var v10 int8 = v7
                            _ = v10
                        }
                    }
                case 737:
                    for fi11 := 0; fi11 < 4; fi11++ {
                        var v12 string = v4
                        v5 += v5
                        for fi13 := 0; fi13 < 12; fi13++ {
                            if false {
                                if (!((!true) && true)) {
                                    if false {
                                        fmt.Println(v3, v3)
                                        var v14 int16 = 261
                                        if ((uint32(v2) < 2) && true) {
                                            if false {
                                                fmt.Println(fi13, v4)
                                                v5 = v5
                                                fmt.Println(v12)
                                            }
                                        } else {
                                            v1 = v1
                                            if (true || (v1 >= v5)) {
                                                var v15 string = "\n\t"
                                                switch v15 {
                                                case "\n\t":
                                                    switch (fi11 ^ (fi13 ^ (fi13 - fi13))) {
                                                    case 4294967295:
                                                        v16 := GenS1{F0: 42}
                                                        _ = v16
                                                    }
                                                    var v17 string = (("a" + v4) + v4)
                                                    _ = v17
                                                case "foo bar":
                                                    var v18 uint8 = 122
                                                    _ = v18
                                                default:
                                                    if false {
                                                        var v19 uint32 = 65535
                                                        if (!true) {
                                                            v19 += (^(v19))
                                                            if (!true) {
                                                                v2 += (v2 << ((2) & 63))
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                                var v20 uint16 = 1
                                                                var v21 int16 = (v14 & int16(v20))
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                                _ = v20
                                                                _ = v21
                                                            } else {
                                                                if true {
                                                                    v4 = v15
                                                                    var v22 float64 = float64(0.0)
                                                                    v12 = "a"
                                                                    _ = v22
                                                                }
                                                                v6 = append(v6, "hello")
                                                            }
                                                            fmt.Println(v2)
                                                        } else {
                                                            var v23 uint64 = v2
                                                            _ = v23
                                                        }
                                                        switch v4 {
                                                        case "":
                                                            if len(v6) > 0 {
                                                                v6[0] = v12
                                                            }
                                                        case "hello":
                                                            fmt.Println(v15)
                                                        case "\n\t":
                                                            _ = len(v6)
                                                        }
                                                        switch (fi13 + 32767) {
                                                        case -1:
                                                            switch v4 {
                                                            case "":
                                                                v24 := GenS1{F0: 65535}
                                                                if false {
                                                                    if false {
                                                                        if false {
                                                                            v3 = (v3 << ((2) & 63))
                                                                            v2 += (v2 << ((1) & 63))
                                                                            if false {
                                                                                v4 = "\n\t"
                                                                            }
                                                                        }
                                                                        fmt.Println(v3, v4)
                                                                    }
                                                                }
                                                                _ = v24
                                                            }
                                                        case 32767:
                                                            var v25 int32 = 2147483647
                                                            _ = v25
                                                        }
                                                        _ = v19
                                                    }
                                                }
                                                fmt.Println(v12, v14, v14)
                                                _ = v15
                                            }
                                        }
                                        _ = v14
                                    }
                                    fmt.Println(v12, fi13, fi11)
                                    if true {
                                        if false {
                                            _ = len(v6)
                                            v5 = v1
                                        }
                                        switch fi13 {
                                        case 32768:
                                            v7 = (v7 | v3)
                                        case -462:
                                            fmt.Println(v1, v2)
                                        }
                                        v3 = (v3 << ((331) & 63))
                                        v2 += v2
                                    } else {
                                        v3 = (^(v7))
                                        v6 = append(v6, v4)
                                    }
                                }
                            }
                            if true {
                                if false {
                                    if true {
                                        var v26 string = (v4 + "hello")
                                        var v27 string = v12
                                        _ = len(v6)
                                        _ = v26
                                        _ = v27
                                    } else {
                                        v6 = append(v6, ((("\n\t" + "a") + v4) + ("foo bar" + ("\n\t" + v4))))
                                        switch (fi13 % ((fi13) | 1)) {
                                        case 318:
                                            if false {
                                                fmt.Println(fi11, v5)
                                                fmt.Println(v4, fi13, fi11)
                                            } else {
                                                v1 = (v1 * int64(v2))
                                            }
                                        case 127:
                                            v3 = (v7 ^ v7)
                                        }
                                        v28 := GenS1{F0: 65535}
                                        _ = v28
                                    }
                                    if false {
                                        var v29 int32 = 598
                                        var v30 bool = (!(255 > 255))
                                        var v31 int8 = (v7 % ((127) | 1))
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                    } else {
                                        v12 = "\n\t"
                                        var v32 uint32 = 732
                                        _ = v32
                                    }
                                } else {
                                    var v33 string = v12
                                    var v34 uint8 = 1
                                    if true {
                                        if true {
                                            var v35 bool = (true || (1 < 2147483647))
                                            v12 = (v12 + v4)
                                            switch v12 {
                                            case "foo bar":
                                                v12 = "a"
                                            case "hello":
                                                if v35 {
                                                    if v35 {
                                                        var v36 uint8 = v34
                                                        if len(v6) > 0 {
                                                            v6[0] = v12
                                                        }
                                                        _ = v36
                                                    }
                                                    switch (v12 + v12) {
                                                    case "\n\t":
                                                        if v35 {
                                                            fmt.Println(v1, v1, fi11)
                                                        } else {
                                                            var v37 uint = 808
                                                            var v38 int64 = v5
                                                            _ = v37
                                                            _ = v38
                                                        }
                                                    case "hello":
                                                        switch (fi11 | -9223372036854775808) {
                                                        case 260:
                                                            var v39 int = fi11
                                                            _ = v39
                                                        case -1:
                                                            if false {
                                                                v6 = append(v6, v4)
                                                            } else {
                                                                v4 = (v4 + (v33 + "a"))
                                                            }
                                                        }
                                                    case "foo bar":
                                                        var v40 uint = 764
                                                        _ = v40
                                                    }
                                                    if len(v6) > 0 {
                                                        v6[0] = ("\n\t" + (("a" + "foo bar") + v12))
                                                    }
                                                    v2 += (^(v2))
                                                    if false {
                                                        switch (fi13 & 192) {
                                                        case 32768:
                                                            if v35 {
                                                                fmt.Println(v2, v5)
                                                                if (0 == -227) {
                                                                    switch (v12 + ("a" + v12)) {
                                                                    case "foo bar":
                                                                        switch fi11 {
                                                                        case 127:
                                                                            var v41 int = fi13
                                                                            _ = v41
                                                                        case -409:
                                                                            v3 -= 42
                                                                        }
                                                                    case "hello":
                                                                        switch "hello" {
                                                                        case "\n\t":
                                                                            v42 := GenS2{F0: (^(v2)), F1: (v5 | (v5 + 932))}
                                                                            _ = v42
                                                                        default:
                                                                            v43 := GenS1{F0: 438}
                                                                            _ = v43
                                                                        }
                                                                    case "a":
                                                                        v44 := GenS3{F0: 2147483647, F1: ((^(v5)) > (v1 ^ 347)), F2: 0}
                                                                        _ = v44
                                                                    default:
                                                                        switch v12 {
                                                                        case "hello":
                                                                            v45 := GenS1{F0: 2}
                                                                            _ = len(v6)
                                                                            _ = len(v6)
                                                                            _ = v45
                                                                        case "a":
                                                                            v7 += 0
                                                                        case "foo bar":
                                                                            fmt.Println(v35, v3, fi11)
                                                                        default:
                                                                            var v46 bool = ((727 >= 0) || ((v34 & (v34 | (^(v34)))) <= 148))
                                                                            _ = v46
                                                                        }
                                                                    }
                                                                }
                                                            } else {
                                                                if ((false || false) && v35) {
                                                                    fmt.Println(v33, v35)
                                                                }
                                                            }
                                                        default:
                                                            if (1 != 1000) {
                                                                fmt.Println(v4)
                                                                fmt.Println(v4, v35)
                                                            }
                                                        }
                                                    }
                                                }
                                            case "\n\t":
                                                var v47 int = (fi13 * (fi11 % (((^(fi11))) | 1)))
                                                _ = v47
                                            default:
                                                _ = len(v6)
                                            }
                                            _ = v35
                                        }
                                        v48 := GenS4{F0: v2}
                                        var v49 string = ((v12 + v33) + ("foo bar" + v4))
                                        _ = v48
                                        _ = v49
                                    }
                                    _ = v33
                                    _ = v34
                                }
                                v3 -= (^(v7))
                                switch -2147483648 {
                                case 4294967295:
                                    switch fi13 {
                                    case -410:
                                        v50 := GenS1{F0: 65535}
                                        _ = v50
                                    default:
                                        var v51 int16 = -1
                                        _ = v51
                                    }
                                case 32768:
                                    v12 = v12
                                }
                            } else {
                                v7 = v7
                                if false {
                                    if (((-298 < 0) || (!(!(true || false)))) && (0 == 42)) {
                                        if (true && (v3 < v7)) {
                                            v6 = append(v6, ((v4 + "") + v12))
                                        } else {
                                            var v52 bool = (!true)
                                            _ = v52
                                        }
                                        v53 := GenS4{F0: v2}
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        _ = v53
                                    }
                                } else {
                                    fmt.Println(fi13, v4)
                                    var v54 uint32 = 39
                                    var v55 int32 = 42
                                    _ = v54
                                    _ = v55
                                }
                                var v56 int32 = 982
                                _ = v56
                            }
                            _ = fi13
                        }
                        _ = fi11
                        _ = v12
                    }
                default:
                    var v57 uint8 = 2
                    _ = v57
                }
                if (1 <= 133) {
                    for fi58 := 0; fi58 < 10; fi58++ {
                        var v59 uint8 = 53
                        v60 := GenS4{F0: v2}
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = fi58
                        _ = v59
                        _ = v60
                    }
                    for fi61 := 0; fi61 < 12; fi61++ {
                        switch v4 {
                        case "a":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        case "\n\t":
                            if true {
                                v62 := GenS5{F0: (v7 * v3), F1: 255}
                                _ = v62
                            } else {
                                for fi63 := 0; fi63 < 11; fi63++ {
                                    var v64 float64 = -2.25
                                    _ = fi63
                                    _ = v64
                                }
                            }
                        case "":
                            var v65 int = 646
                            _ = v65
                        }
                        var v66 int16 = 1
                        v4 = v4
                        _ = fi61
                        _ = v66
                    }
                    v7 += v3
                }
            }
            switch 4294967295 {
            case 1:
                switch 479 {
                case 42:
                    fmt.Println(v2)
                case 608:
                    v2 = v2
                }
            case 127:
                if false {
                    if len(v6) > 0 {
                        v6[0] = v4
                    }
                    var v67 int8 = 98
                    _ = v67
                }
            default:
                v7 = (-(v3))
            }
        } else {
            v5 += v5
        }
    } else {
        fmt.Println(v7)
    }
    var v68 int = 756
    switch v4 {
    case "":
        var v69 float64 = 0.0
        _ = v69
    default:
        v1 -= v5
    }
    v70 := GenS6{F0: (v1 & (v1 / ((v5) | 1))), F1: -2147483648}
    v5 -= (v5 | 9223372036854775807)
    if (0.1 >= 3.14159) {
        if ((v5 << ((237) & 63)) > 614) {
            switch v68 {
            case 9223372036854775807:
                fmt.Println(v1, v4, v3)
            default:
                switch (v68 & 42) {
                case -96:
                    for fi71 := 0; fi71 < 7; fi71++ {
                        var v72 int64 = v5
                        fmt.Println(v5)
                        _ = fi71
                        _ = v72
                    }
                case 390:
                    var v73 int8 = (v3 / (((v3 | v7)) | 1))
                    _ = v73
                case 2147483647:
                    v3 -= v7
                }
            }
        } else {
            v74 := GenS4{F0: v2}
            if true {
                var v75 int32 = 2147483647
                fmt.Println(v1, v4)
                v76 := GenS7{F0: v2, F1: 658}
                _ = v75
                _ = v76
            } else {
                v3 = (v3 + v7)
            }
            _ = len(v6)
            _ = v74
        }
    } else {
        v3 -= 0
    }
    v2 += (v2 + v2)
    v77 := GenS8{F0: v1, F1: 2, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v68:", v68)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v7
    _ = v70
    _ = v77
}
