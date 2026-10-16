package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int8; F1 int; F2 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 int64; F2 uint32 }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint32; F1 int64 }
type GenS9 struct { F0 int32; F1 uint; F2 int16 }
type GenS10 struct { F0 int32; F1 uint16; F2 bool }

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = true
    var v3 float32 = 1.5
    v4 := make(map[string]uint32)
    v5 := []int{}
    switch 299 {
    case -1:
        var v6 int8 = -1
        _ = v6
    case 323:
        for fi7 := 0; fi7 < 11; fi7++ {
            v8 := GenS1{F0: true, F1: -52}
            v2 = v2
            if (9 == 1000) {
                v2 = (65535 == 65535)
                if len(v5) > 0 {
                    v5[0] = fi7
                }
                v5 = append(v5, fi7)
            } else {
                for fi9 := 0; fi9 < 15; fi9++ {
                    v2 = true
                    _ = fi9
                }
                v3 -= v3
            }
            _ = fi7
            _ = v8
        }
    case 2147483647:
        if (!(90 != 234)) {
            v10 := GenS1{F0: v2, F1: 127}
            v11 := GenS2{F0: -877, F1: 754}
            switch 9223372036854775807 {
            case -430:
                v12 := GenS3{F0: -63, F1: 128, F2: 2}
                _ = v12
            case 32767:
                if (!true) {
                    switch -881 {
                    case 32767:
                        switch 9223372036854775807 {
                        case -212:
                            v13 := GenS1{F0: v2, F1: -26}
                            _ = v13
                        }
                    case 1:
                        var v14 int8 = 1
                        _ = v14
                    case 32768:
                        if (("\n\t" > ("" + "\n\t")) && false) {
                            v3 = 0.1
                            var v15 float32 = 0.1
                            _ = v15
                        } else {
                            if v2 {
                                if v2 {
                                    var v16 bool = v2
                                    _ = v16
                                }
                                v17 := GenS4{F0: 42}
                                v18 := GenS4{F0: 155}
                                _ = v17
                                _ = v18
                            }
                            switch 4294967295 {
                            case -2147483648:
                                fmt.Println(v2)
                            case 422:
                                v2 = (201 < 1)
                            default:
                                var v19 bool = false
                                _ = v19
                            }
                        }
                    default:
                        var v20 uint64 = 243
                        _ = v20
                    }
                    var v21 bool = (!(-2147483648 == 4294967295))
                    if v21 {
                        v2 = (v2 && (v3 >= (v3 * (0.1 * ((((1.5 + 1.5) * v3) * v3) - float32((float64((float32(0.1) + (float32(0.1) * 0.0))) - (1e10 / 0.1))))))))
                        for fi22 := 0; fi22 < 17; fi22++ {
                            if v2 {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                if (((fi22 * (^(fi22))) != fi22) && (!v21)) {
                                    var v23 uint32 = 4294967295
                                    _ = v23
                                }
                            } else {
                                var v24 uint16 = 0
                                var v25 int16 = 662
                                v21 = ((v24 * v24) != v24)
                                _ = v24
                                _ = v25
                            }
                            _ = fi22
                        }
                        v3 += (((v3 * v3) / 1.0) + (0.1 + float32((1.5 / 1.5))))
                    } else {
                        fmt.Println(v2, v3)
                        fmt.Println(v3, v2, v21)
                    }
                    _ = v21
                } else {
                    for fi26 := 0; fi26 < 17; fi26++ {
                        v3 -= -2.25
                        v2 = v2
                        _ = fi26
                    }
                    if (2147483647 < 0) {
                        v3 = (((v3 - (((v3 * float32((0.1 - ((1.5 * 1e10) - (0.0 / -2.25))))) * 1.5) + 0.0)) / v3) + ((v3 - float32(1.5)) + v3))
                        var v27 string = "foo bar"
                        _ = v27
                    } else {
                        v2 = (!(1 >= 42))
                        v5 = append(v5, 128)
                    }
                    var v28 int = -2147483648
                    _ = v28
                }
            }
            _ = v10
            _ = v11
        } else {
            switch 127 {
            case 32768:
                var v29 uint = 2
                _ = v29
            case 42:
                if (42 >= 18446744073709551615) {
                    if (42 >= -1) {
                        switch "" {
                        case "\n\t":
                            fmt.Println(v3, v3)
                        default:
                            v2 = (v2 && v2)
                        }
                        for fi30 := 0; fi30 < 4; fi30++ {
                            var v31 int32 = -299
                            _ = fi30
                            _ = v31
                        }
                        var v32 int = 176
                        _ = v32
                    }
                }
            case -188:
                v33 := GenS5{F0: 773, F1: 0, F2: 0}
                _ = v33
            default:
                var v34 int16 = 0
                _ = v34
            }
        }
    }
    fmt.Println(v3, v2)
    switch "a" {
    case "":
        v35 := GenS6{F0: (v2 && v2), F1: 2}
        _ = v35
    }
    if len(v5) > 0 {
        v5[0] = 0
    }
    if v2 {
        for fi36 := 0; fi36 < 10; fi36++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v37 := GenS7{F0: -880}
            _ = fi36
            _ = v37
        }
        switch 0 {
        case -147:
            var v38 int64 = -9223372036854775808
            _ = v38
        case 32768:
            if v2 {
                var v39 int16 = -32768
                for fi40 := 0; fi40 < 4; fi40++ {
                    var v41 float32 = 3.14159
                    _ = fi40
                    _ = v41
                }
                _ = v39
            } else {
                fmt.Println(v3, v2)
                v2 = (!false)
            }
        case 128:
            v3 = ((1.5 + ((0.1 - v3) + v3)) + -2.25)
        default:
            v42 := GenS3{F0: 100, F1: 833, F2: 18446744073709551615}
            _ = v42
        }
    } else {
        var v43 string = "hello"
        _ = v43
    }
    v44 := GenS3{F0: 127, F1: 127, F2: 255}
    if (!("hello" == ("foo bar" + "a"))) {
        fmt.Println(v2, v3)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = len(v5)
    } else {
        for fi45 := 0; fi45 < 19; fi45++ {
            if v2 {
                var v46 uint16 = 42
                if false {
                    switch (fi45 & 1) {
                    case -2147483648:
                        for fi47 := 0; fi47 < 11; fi47++ {
                            v48 := GenS6{F0: ((^(v46)) > 42), F1: 778}
                            switch (("" + ("\n\t" + "\n\t")) + "\n\t") {
                            case "":
                                var v49 bool = (v2 && ((2147483647 < 2147483647) || v2))
                                v50 := GenS1{F0: v2, F1: 122}
                                if len(v5) > 0 {
                                    v5[0] = fi47
                                }
                                _ = v49
                                _ = v50
                            case "hello":
                                fmt.Println(fi45, fi45, v3)
                            case "foo bar":
                                if (!true) {
                                    fmt.Println(v2, v3, v2)
                                } else {
                                    _ = len(v1)
                                }
                            }
                            var v51 int64 = -56
                            _ = fi47
                            _ = v48
                            _ = v51
                        }
                    default:
                        v5 = append(v5, (fi45 + fi45))
                    }
                    if (!(v2 && ((v46 <= v46) && (!(451 < -1))))) {
                        var v52 int64 = 1
                        _ = v52
                    } else {
                        delete(v4, "")
                        v46 += (v46 * (^(v46)))
                        if v2 {
                            for fi53 := 0; fi53 < 13; fi53++ {
                                fmt.Println(v46)
                                _ = fi53
                            }
                        }
                    }
                } else {
                    if (true || v2) {
                        var v54 float32 = ((float32(1.5) - 0.1) * float32(0.0))
                        v55 := v1[25]
                        v55 = (fi45 << ((4294967295) & 63))
                        _ = v54
                        _ = v55
                    } else {
                        fmt.Println(v46, v2)
                        v56 := GenS8{F0: 0, F1: -570}
                        if true {
                            for fi57 := 0; fi57 < 4; fi57++ {
                                if v2 {
                                    if (false && (v2 || true)) {
                                        var v58 int8 = 100
                                        switch "foo bar" {
                                        case "\n\t":
                                            v59 := v4["\n\t"]
                                            _ = v59
                                        case "":
                                            v60 := GenS1{F0: false, F1: v58}
                                            v2 = v2
                                            v3 = v3
                                            v61 := GenS3{F0: (v58 & v58), F1: 4294967295, F2: 1}
                                            _ = v60
                                            _ = v61
                                        case "hello":
                                            v62 := GenS4{F0: 792}
                                            _ = v62
                                        }
                                        _ = v58
                                    }
                                }
                                if true {
                                    if len(v5) > 0 {
                                        v5[0] = -508
                                    }
                                    var v63 uint8 = 255
                                    switch "\n\t" {
                                    case "a":
                                        var v64 int16 = 877
                                        _ = v64
                                    default:
                                        fmt.Println(v63, v3)
                                    }
                                    var v65 int = (fi45 / (((fi57 * fi45)) | 1))
                                    _ = v63
                                    _ = v65
                                } else {
                                    v2 = true
                                }
                                if v2 {
                                    if v2 {
                                        var v66 int32 = 42
                                        _ = v66
                                    } else {
                                        var v67 int8 = -123
                                        v68 := GenS9{F0: 42, F1: 65535, F2: 1}
                                        if (42 >= 18446744073709551615) {
                                            v46 = (v46 % ((579) | 1))
                                        }
                                        fmt.Println(v3, fi57)
                                        v67 -= v67
                                        _ = v67
                                        _ = v68
                                    }
                                    v3 -= (0.0 + (v3 / 1.0))
                                }
                                _ = fi57
                            }
                            if true {
                                v69 := GenS10{F0: int32(fi45), F1: v46, F2: (v46 != v46)}
                                for fi70 := 0; fi70 < 19; fi70++ {
                                    var v71 uint16 = v46
                                    v71 -= 526
                                    var v72 uint8 = 1
                                    fmt.Println(v2, v71)
                                    _ = fi70
                                    _ = v71
                                    _ = v72
                                }
                                var v73 float64 = 1.5
                                _ = v69
                                _ = v73
                            }
                            v46 -= (v46 & (v46 >> ((0) & 63)))
                        }
                        _ = v56
                    }
                    fmt.Println(v46, v3)
                    var v74 float64 = -2.25
                    _ = v74
                }
                if (!(-1 == 9223372036854775807)) {
                    var v75 int64 = -9223372036854775808
                    v75 -= (^(v75))
                    _ = v75
                }
                _ = v46
            }
            _ = fi45
        }
        fmt.Println(v3, v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
}
