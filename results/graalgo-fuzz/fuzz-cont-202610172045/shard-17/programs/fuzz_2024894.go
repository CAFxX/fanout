package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint; F2 bool }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int64 }

func main() {
    var v1 uint = 1
    var v2 bool = (0 >= -2147483648)
    v3 := []string{("a" + ("foo bar" + ("hello" + "\n\t"))), "\n\t", ""}
    v4 := make(map[uint32]bool)
    var v5 uint32 = 600
    if v2 {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        if (-1 != 442) {
            for fi6 := 0; fi6 < 6; fi6++ {
                fmt.Println(fi6, v5, fi6)
                _ = len(v4)
                switch fi6 {
                case 4294967295:
                    _ = len(v4)
                default:
                    fmt.Println(v1)
                }
                _ = fi6
            }
            for fi7 := 0; fi7 < 5; fi7++ {
                v5 -= (v5 + 2)
                v8 := GenS1{F0: 680, F1: (v1 / (((^(v1))) | 1)), F2: 47}
                var v9 float32 = 3.14159
                _ = fi7
                _ = v8
                _ = v9
            }
        }
        if v2 {
            v5 += v5
            var v10 uint32 = (^(v5))
            _ = v10
        } else {
            v2 = (-1 != 1000)
            for fi11 := 0; fi11 < 8; fi11++ {
                v2 = v2
                if v2 {
                    var v12 uint64 = 18446744073709551615
                    for fi13 := 0; fi13 < 14; fi13++ {
                        var v14 int8 = 122
                        switch "hello" {
                        case "foo bar":
                            if (137 > uint8(v12)) {
                                switch "\n\t" {
                                case "\n\t":
                                    var v15 uint16 = 89
                                    _ = v15
                                default:
                                    if (-1 != -844) {
                                        fmt.Println(v5, fi13, v5)
                                    } else {
                                        var v16 int16 = 0
                                        _ = v16
                                    }
                                }
                                v3 = append(v3, "")
                            } else {
                                if v2 {
                                    v17 := GenS1{F0: 2, F1: v1, F2: 65535}
                                    fmt.Println(v5)
                                    _ = v17
                                } else {
                                    if false {
                                        v4[v5] = ("\n\t" == "a")
                                        var v18 int8 = v14
                                        _ = v18
                                    }
                                    v19 := GenS2{F0: 2, F1: v1, F2: v2}
                                    fmt.Println(fi13)
                                    _ = v19
                                }
                                v5 = v5
                                _ = len(v4)
                                v5 += 768
                            }
                        default:
                            var v20 int8 = v14
                            _ = v20
                        }
                        v21 := GenS2{F0: 0, F1: (v1 / ((v1) | 1)), F2: (!v2)}
                        var v22 uint32 = v5
                        _ = fi13
                        _ = v14
                        _ = v21
                        _ = v22
                    }
                    if (-379 != -1) {
                        for fi23 := 0; fi23 < 7; fi23++ {
                            v24 := GenS2{F0: 204, F1: 42, F2: (42 != 127)}
                            var v25 uint64 = v12
                            v26 := GenS2{F0: 0, F1: uint(v5), F2: v2}
                            _ = fi23
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                    }
                    _ = v12
                }
                _ = fi11
            }
            for fi27 := 0; fi27 < 4; fi27++ {
                switch 127 {
                case 128:
                    v2 = false
                case 1:
                    if v2 {
                        var v28 float64 = 0.0
                        _ = v28
                    } else {
                        for fi29 := 0; fi29 < 2; fi29++ {
                            var v30 int16 = 32767
                            _ = fi29
                            _ = v30
                        }
                        fmt.Println(v5, v2)
                    }
                default:
                    if false {
                        var v31 string = ("" + "\n\t")
                        v32 := GenS3{F0: v1, F1: 799}
                        _ = v31
                        _ = v32
                    }
                }
                _ = fi27
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if (true && v2) {
        v33 := GenS2{F0: uint8(v5), F1: v1, F2: v2}
        v2 = false
        _ = v33
    } else {
        var v34 int16 = 42
        v3 = append(v3, (("a" + "\n\t") + "foo bar"))
        for fi35 := 0; fi35 < 2; fi35++ {
            var v36 string = ("foo bar" + ("foo bar" + ""))
            _ = fi35
            _ = v36
        }
        _ = v34
    }
    switch "a" {
    case "a":
        if (!((0 < 2) || v2)) {
            var v37 string = ""
            if v2 {
                if len(v3) > 0 {
                    v3[0] = "a"
                }
            } else {
                for fi38 := 0; fi38 < 2; fi38++ {
                    var v39 int8 = 127
                    v39 += (v39 | (v39 * 42))
                    v37 = (v37 + v37)
                    _ = fi38
                    _ = v39
                }
            }
            if false {
                delete(v4, v5)
                for fi40 := 0; fi40 < 13; fi40++ {
                    v41 := GenS1{F0: (^(v1)), F1: (v1 >> ((v1) & 63)), F2: 1}
                    v1 -= v1
                    v5 -= v5
                    _ = fi40
                    _ = v41
                }
            } else {
                for fi42 := 0; fi42 < 3; fi42++ {
                    if true {
                        var v43 int32 = 128
                        var v44 float32 = (0.1 - 3.14159)
                        for fi45 := 0; fi45 < 7; fi45++ {
                            var v46 float64 = (-2.25 - 1.5)
                            var v47 uint = (v1 / ((v1) | 1))
                            _ = fi45
                            _ = v46
                            _ = v47
                        }
                        _ = v43
                        _ = v44
                    }
                    for fi48 := 0; fi48 < 10; fi48++ {
                        v49 := GenS3{F0: 1, F1: v1}
                        var v50 bool = ((v1 ^ v1) != (v1 / ((364) | 1)))
                        _ = fi48
                        _ = v49
                        _ = v50
                    }
                    _ = fi42
                }
                for fi51 := 0; fi51 < 18; fi51++ {
                    v37 = v37
                    if v2 {
                        v5 = v5
                        if (v2 && v2) {
                            v1 = (v1 % (((v1 ^ v1)) | 1))
                        } else {
                            for fi52 := 0; fi52 < 20; fi52++ {
                                var v53 uint = v1
                                if false {
                                    v5 -= v5
                                    fmt.Println(v1, v53, fi51)
                                    var v54 uint8 = 255
                                    _ = v54
                                }
                                v55 := GenS1{F0: (v53 >> (((v1 >> (((^(v53))) & 63))) & 63)), F1: 65535, F2: 1}
                                _ = fi52
                                _ = v53
                                _ = v55
                            }
                        }
                        for fi56 := 0; fi56 < 9; fi56++ {
                            v2 = v2
                            v57 := GenS3{F0: 0, F1: 0}
                            _ = fi56
                            _ = v57
                        }
                    }
                    _ = fi51
                }
                var v58 int8 = 0
                _ = v58
            }
            _ = v37
        }
    case "foo bar":
        for fi59 := 0; fi59 < 11; fi59++ {
            v4[(v5 | 416)] = true
            _ = fi59
        }
    case "hello":
        if (-85 != -534) {
            for fi60 := 0; fi60 < 10; fi60++ {
                var v61 uint = (v1 + 1)
                if false {
                    v2 = (!((v5 >= 2147483647) && (52 < v5)))
                }
                _ = fi60
                _ = v61
            }
        }
    default:
        var v62 uint = 42
        _ = v62
    }
    var v63 bool = (2147483648 > 2)
    var v64 string = "foo bar"
    v2 = (3.14159 == float32(float64(0.1)))
    var v65 uint8 = 1
    if (0.1 <= -2.25) {
        v66 := GenS2{F0: (^(v65)), F1: v1, F2: false}
        fmt.Println(v63, v5, v63)
        fmt.Println(v1, v5, v64)
        _ = v66
    } else {
        for fi67 := 0; fi67 < 16; fi67++ {
            v1 = v1
            _ = fi67
        }
        v5 -= 1
    }
    for fi68 := 0; fi68 < 6; fi68++ {
        v65 = v65
        switch "foo bar" {
        case "foo bar":
            v65 += 127
        case "\n\t":
            if false {
                v63 = v63
            } else {
                var v69 int16 = 0
                var v70 string = ("a" + (v64 + (((v64 + (v64 + "foo bar")) + v64) + ((((v64 + v64) + v64) + "a") + v64))))
                switch -964 {
                case 0:
                    v63 = ((42 >= 9223372036854775807) && v63)
                case -952:
                    switch fi68 {
                    case 372:
                        v71 := GenS1{F0: 51, F1: (^(v1)), F2: 65535}
                        _ = v71
                    case 138:
                        v1 = 323
                    case -9223372036854775808:
                        fmt.Println(v69, fi68, v63)
                    }
                case -2147483648:
                    if (-2.25 == (3.14159 / (1.5 - 0.0))) {
                        switch fi68 {
                        case 127:
                            fmt.Println(v69, v5, v69)
                        }
                    } else {
                        fmt.Println(v1, v69, v2)
                    }
                default:
                    if len(v3) > 0 {
                        v3[0] = v70
                    }
                }
                _ = v69
                _ = v70
            }
        case "a":
            var v72 int64 = 0
            _ = v72
        }
        v73 := GenS4{F0: 9223372036854775807}
        _ = fi68
        _ = v73
    }
    var v74 bool = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v63
    _ = v64
    _ = v65
    _ = v74
}
