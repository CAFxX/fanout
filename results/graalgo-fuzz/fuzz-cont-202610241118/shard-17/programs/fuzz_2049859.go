package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 string; F1 uint16; F2 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 uint16; F2 int8 }
type GenS7 struct { F0 uint16; F1 int8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint }
func genFn0() int64 {
    v7 := GenS1{F0: 100, F1: "foo bar", F2: 966}
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            v10 := GenS1{F0: -85, F1: "", F2: 18446744073709551615}
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: -128, F1: "\n\t", F2: 563}
        switch "\n\t" {
        case "hello":
            v12 := []int{}
            _ = v12
        }
        _ = fi8
        _ = v11
    }
    _ = v7
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 255
    var v3 uint = 255
    v4 := make(map[string]bool)
    v5 := []uint32{416}
    var v6 uint32 = 42
    var v13 int = -1
    v2 -= 128
    if false {
        v14 := v4[(("a" + ("\n\t" + "\n\t")) + "hello")]
        var v15 int32 = -1
        if false {
            if false {
                if (v14 && (v14 || (v14 && v14))) {
                    v14 = v14
                    if v14 {
                        v13 += -9223372036854775808
                        var v16 string = "hello"
                        if len(v5) > 0 {
                            v5[0] = (v6 % (((v6 / ((42) | 1))) | 1))
                        }
                        _ = v16
                    }
                }
                if ("" <= (("" + "foo bar") + "\n\t")) {
                    var v17 float64 = 0.1
                    switch v13 {
                    case -1:
                        switch "foo bar" {
                        case "":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        case "a":
                            if ((((-292 > 1) && v14) && ((((3.14159 / 1.0) - 0.1) / 1.0) < float32(0.1))) || v14) {
                                switch v13 {
                                case 0:
                                    var v18 bool = v14
                                    _ = v18
                                default:
                                    fmt.Println(v6, v15, v15)
                                }
                            } else {
                                fmt.Println(v13)
                                if true {
                                    switch "a" {
                                    case "\n\t":
                                        fmt.Println(v6, v2, v15)
                                    case "foo bar":
                                        v19 := GenS1{F0: -1, F1: ("hello" + ("a" + "a")), F2: 528}
                                        _ = v19
                                    }
                                    v1[v6] = "hello"
                                } else {
                                    if ((v2 >= v2) || v14) {
                                        var v20 uint = (v3 % ((v3) | 1))
                                        v20 += v3
                                        if true {
                                            v21 := GenS1{F0: 31, F1: "hello", F2: 506}
                                            _ = v21
                                        } else {
                                            switch "hello" {
                                            case "\n\t":
                                                for fi22 := 0; fi22 < 18; fi22++ {
                                                    switch ("\n\t" + "a") {
                                                    case "":
                                                        v3 = 255
                                                    case "\n\t":
                                                        var v23 int8 = 100
                                                        _ = v23
                                                    case "a":
                                                        var v24 bool = v14
                                                        _ = v24
                                                    default:
                                                        for fi25 := 0; fi25 < 13; fi25++ {
                                                            if (1 >= v15) {
                                                                var v26 int64 = -9223372036854775808
                                                                _ = v26
                                                            }
                                                            delete(v4, "")
                                                            v1[(v6 - 995)] = "a"
                                                            _ = fi25
                                                        }
                                                    }
                                                    _ = fi22
                                                }
                                            }
                                            v27 := GenS2{F0: v14, F1: 1, F2: (^(v6))}
                                            if v14 {
                                                if (1.5 == 1e10) {
                                                    var v28 uint64 = 42
                                                    _ = v28
                                                }
                                                if (1.5 >= (-2.25 - float32(v17))) {
                                                    if v14 {
                                                        switch (v13 ^ v13) {
                                                        case 2147483647:
                                                            var v29 int8 = 40
                                                            _ = v29
                                                        default:
                                                            var v30 uint32 = v6
                                                            _ = v30
                                                        }
                                                        var v31 int8 = 42
                                                        _ = v31
                                                    } else {
                                                        fmt.Println(v6, v20, v13)
                                                        fmt.Println(v6, v6)
                                                    }
                                                    if v14 {
                                                        v3 -= (v3 ^ uint(v2))
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            if v14 {
                                                                v32 := GenS1{F0: 1, F1: "a", F2: 42}
                                                                if (32767 > -754) {
                                                                    switch (-(v13)) {
                                                                    case 42:
                                                                        var v33 uint8 = 255
                                                                        _ = v33
                                                                    case 2147483647:
                                                                        var v34 float32 = (float32(float64(((0.0 + float32(v17)) / 1.5))) - 1.5)
                                                                        _ = v34
                                                                    case -309:
                                                                        v5 = append(v5, v6)
                                                                    default:
                                                                        v13 = v13
                                                                    }
                                                                    v17 -= float64(3.14159)
                                                                } else {
                                                                    if true {
                                                                        fmt.Println(v13, v20, v14)
                                                                    } else {
                                                                        for fi35 := 0; fi35 < 15; fi35++ {
                                                                            v36 := GenS1{F0: 42, F1: "a", F2: 1}
                                                                            v20 = 2
                                                                            fmt.Println(v15, v13)
                                                                            _ = fi35
                                                                            _ = v36
                                                                        }
                                                                        for fi37 := 0; fi37 < 14; fi37++ {
                                                                            var v38 int8 = 30
                                                                            var v39 int32 = (v15 | (v15 ^ (v15 | (^(v15)))))
                                                                            _ = fi37
                                                                            _ = v38
                                                                            _ = v39
                                                                        }
                                                                    }
                                                                }
                                                                _ = v32
                                                            }
                                                        case "foo bar":
                                                            delete(v4, ("foo bar" + "a"))
                                                        case "a":
                                                            v6 += (v6 + (v6 & 0))
                                                        }
                                                    } else {
                                                        v20 = v20
                                                        v40 := v4["\n\t"]
                                                        _ = v40
                                                    }
                                                    for fi41 := 0; fi41 < 18; fi41++ {
                                                        if ("foo bar" < "a") {
                                                            if (-1 < 968) {
                                                                var v42 int8 = 109
                                                                v20 -= v3
                                                                _ = v42
                                                            } else {
                                                                var v43 int = (fi41 - (v13 - v13))
                                                                for fi44 := 0; fi44 < 13; fi44++ {
                                                                    var v45 uint = (v20 >> ((uint(v6)) & 63))
                                                                    _ = fi44
                                                                    _ = v45
                                                                }
                                                                _ = v43
                                                            }
                                                            for fi46 := 0; fi46 < 12; fi46++ {
                                                                v47 := GenS1{F0: -128, F1: (("\n\t" + "\n\t") + "hello"), F2: uint64(v2)}
                                                                switch (fi46 | v13) {
                                                                case 142:
                                                                    var v48 int16 = -302
                                                                    _ = v48
                                                                case 1:
                                                                    v49 := GenS3{F0: v6, F1: "a"}
                                                                    _ = v49
                                                                case 128:
                                                                    var v50 uint8 = (v2 >> ((42) & 63))
                                                                    _ = v50
                                                                }
                                                                if false {
                                                                    if true {
                                                                        v51 := GenS4{F0: ("" + "a"), F1: 2, F2: 18446744073709551615}
                                                                        v5 = append(v5, (v6 + 2))
                                                                        _ = v51
                                                                    }
                                                                    if (1 != 42) {
                                                                        v13 = 32768
                                                                    } else {
                                                                        if (float64(float32((v17 * v17))) >= v17) {
                                                                            fmt.Println(v3)
                                                                            v17 = (1e10 + v17)
                                                                        }
                                                                        v13 = fi46
                                                                    }
                                                                    var v52 string = "foo bar"
                                                                    _ = v52
                                                                } else {
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                }
                                                                _ = fi46
                                                                _ = v47
                                                            }
                                                        }
                                                        if v14 {
                                                            var v53 int32 = v15
                                                            switch ("" + ("foo bar" + ("" + "a"))) {
                                                            case "":
                                                                var v54 int = (v13 % ((1) | 1))
                                                                _ = v54
                                                            case "hello":
                                                                v55 := v1[(v6 & 0)]
                                                                _ = v55
                                                            case "foo bar":
                                                                v56 := GenS5{F0: 181}
                                                                _ = v56
                                                            }
                                                            v3 = (v20 + uint(v2))
                                                            _ = v53
                                                        }
                                                        _ = fi41
                                                    }
                                                } else {
                                                    var v57 uint = (^(v3))
                                                    for fi58 := 0; fi58 < 11; fi58++ {
                                                        if v14 {
                                                            v59 := GenS1{F0: 0, F1: "a", F2: 17}
                                                            switch fi58 {
                                                            case 258:
                                                                var v60 uint8 = (v2 & 127)
                                                                _ = v60
                                                            }
                                                            _ = v59
                                                        } else {
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            if (false && v14) {
                                                                v61 := GenS6{F0: v15, F1: 540, F2: 0}
                                                                _ = v61
                                                            }
                                                        }
                                                        _ = fi58
                                                    }
                                                    v62 := GenS7{F0: 2, F1: -128}
                                                    _ = v57
                                                    _ = v62
                                                }
                                                v63 := GenS6{F0: (v15 % ((v15) | 1)), F1: 855, F2: -1}
                                                _ = v63
                                            }
                                            _ = v27
                                        }
                                        _ = v20
                                    } else {
                                        v64 := GenS4{F0: "foo bar", F1: 32767, F2: 42}
                                        _ = v64
                                    }
                                    _ = len(v4)
                                }
                                var v65 uint32 = v6
                                _ = v65
                            }
                        }
                    case 551:
                        for fi66 := 0; fi66 < 17; fi66++ {
                            if (2 <= 832) {
                                if v14 {
                                    v67 := GenS8{F0: 1}
                                    for fi68 := 0; fi68 < 9; fi68++ {
                                        var v69 uint = (v3 / (((v3 & (v3 >> (((v3 >> ((v3) & 63))) & 63)))) | 1))
                                        _ = fi68
                                        _ = v69
                                    }
                                    _ = v67
                                } else {
                                    fmt.Println(v13, v17)
                                }
                            }
                            _ = fi66
                        }
                    }
                    v70 := GenS7{F0: uint16(v6), F1: int8(v13)}
                    _ = v17
                    _ = v70
                }
                v71 := GenS4{F0: ("foo bar" + "\n\t"), F1: 2, F2: 42}
                _ = v71
            }
            v2 = (^(v2))
            var v72 uint16 = 913
            _ = v72
        }
        _ = v14
        _ = v15
    }
    if (!false) {
        var v73 int16 = 42
        _ = v73
    }
    var v74 float64 = ((0.1 - 0.0) * 1e10)
    var v75 int64 = -9223372036854775808
    v13 = v13
    v74 -= v74
    fmt.Println(v2, v75)
    if false {
        var v76 uint16 = 2
        v6 = 538
        v2 += (v2 - (^(v2)))
        _ = v76
    } else {
        if false {
            fmt.Println(v74, v75, v3)
        } else {
            _ = len(v5)
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v74 = v74
        }
        v3 -= v3
    }
    var v77 int = v13
    v75 = 9223372036854775807
    fmt.Println(v74, v75)
    if (2 <= 1) {
        v2 = v2
        var v78 float32 = 0.1
        _ = v78
    }
    for fi79 := 0; fi79 < 20; fi79++ {
        _ = len(v4)
        _ = fi79
    }
    for fi80 := 0; fi80 < 6; fi80++ {
        if true {
            v81 := GenS9{F0: 255}
            _ = v81
        }
        _ = len(v5)
        v2 += (v2 << (((v3 ^ (v3 + (v3 + 704)))) & 63))
        _ = fi80
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v74
    _ = v75
    _ = v77
}
