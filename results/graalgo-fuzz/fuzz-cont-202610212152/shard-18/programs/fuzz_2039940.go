package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 int; F1 string; F2 int8 }
type GenS7 struct { F0 uint32 }

func main() {
    var v1 int8 = -1
    var v2 float64 = 0.0
    var v3 bool = (true && (((2 <= 2) && (-1 > -2147483648)) && true))
    v4 := make(map[uint32]bool)
    var v5 uint = 42
    if true {
        var v6 uint32 = 2
        _ = v6
    } else {
        fmt.Println(v5)
    }
    switch 2147483647 {
    case 42:
        fmt.Println(v5, v2)
    default:
        for fi7 := 0; fi7 < 4; fi7++ {
            switch -1 {
            case -898:
                var v8 int = 42
                _ = v8
            case 0:
                var v9 uint = (v5 * v5)
                _ = v9
            }
            if ((((1e10 / 1e10) + 0.0) + 0.1) != 0.1) {
                switch 9223372036854775807 {
                case 127:
                    var v10 string = ("foo bar" + "foo bar")
                    _ = v10
                case -2147483648:
                    var v11 uint = (v5 / (((^(v5))) | 1))
                    _ = v11
                case -1:
                    var v12 uint16 = 187
                    _ = v12
                }
            }
            v13 := []int{(fi7 + (^(fi7))), fi7, (fi7 >> (((v5 * (v5 << ((v5) & 63)))) & 63))}
            _ = fi7
            _ = v13
        }
    }
    var v14 bool = (!(1 <= 0))
    v14 = false
    for fi15 := 0; fi15 < 19; fi15++ {
        switch (fi15 % ((fi15) | 1)) {
        case -242:
            for fi16 := 0; fi16 < 13; fi16++ {
                v2 = float64(((0.0 - 0.0) / 1.5))
                switch "\n\t" {
                case "a":
                    switch "\n\t" {
                    case "foo bar":
                        v17 := GenS1{F0: 2}
                        _ = v17
                    case "":
                        var v18 int64 = -531
                        _ = v18
                    case "a":
                        fmt.Println(fi16, v5, v2)
                    default:
                        v5 += v5
                    }
                default:
                    var v19 uint64 = 18446744073709551615
                    _ = v19
                }
                fmt.Println(fi16)
                _ = fi16
            }
        case -9223372036854775808:
            v20 := GenS1{F0: 2}
            _ = v20
        case 710:
            for fi21 := 0; fi21 < 8; fi21++ {
                v2 = v2
                _ = fi21
            }
        }
        var v22 uint32 = 57
        _ = fi15
        _ = v22
    }
    var v23 int16 = -666
    v2 -= 3.14159
    v2 = float64((1e10 / 1.0))
    var v24 int32 = 2147483647
    switch -2147483648 {
    case -333:
        for fi25 := 0; fi25 < 3; fi25++ {
            v2 = float64(0.0)
            v14 = false
            if (v2 >= v2) {
                var v26 uint32 = 56
                _ = v26
            }
            _ = fi25
        }
    case -1:
        fmt.Println(v2)
    case 2147483647:
        v27 := []int{-714}
        _ = v27
    default:
        switch "hello" {
        case "a":
            var v28 int32 = 0
            _ = v28
        case "foo bar":
            for fi29 := 0; fi29 < 5; fi29++ {
                var v30 bool = true
                v31 := GenS2{F0: 383, F1: fi29}
                var v32 int64 = -9223372036854775808
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
        case "":
            var v33 uint32 = 42
            _ = v33
        }
    }
    if (v14 && v3) {
        switch "" {
        case "a":
            if true {
                var v34 int32 = (v24 * (v24 * v24))
                if true {
                    var v35 int64 = -1
                    v36 := GenS1{F0: 921}
                    for fi37 := 0; fi37 < 15; fi37++ {
                        v38 := GenS1{F0: 0}
                        fmt.Println(v2)
                        _ = fi37
                        _ = v38
                    }
                    _ = v35
                    _ = v36
                }
                delete(v4, 42)
                _ = v34
            }
        case "":
            var v39 uint16 = 769
            _ = v39
        case "foo bar":
            fmt.Println(v24)
        default:
            var v40 int32 = 331
            _ = v40
        }
        switch "a" {
        case "hello":
            v23 -= (v23 + v23)
        case "\n\t":
            v1 = -128
        case "a":
            fmt.Println(v14, v5)
        }
        var v41 uint8 = 42
        _ = v41
    }
    if v3 {
        if v14 {
            var v42 string = "a"
            if ((!(!true)) || (861 == (^(v5)))) {
                v43 := GenS2{F0: -1, F1: 414}
                _ = v43
            } else {
                switch 4294967295 {
                case -819:
                    switch (("" + (v42 + ((v42 + (v42 + (v42 + "foo bar"))) + "a"))) + ("a" + ("foo bar" + ("hello" + "foo bar")))) {
                    case "foo bar":
                        v44 := GenS1{F0: 1}
                        _ = v44
                    case "hello":
                        v3 = v14
                    case "a":
                        if true {
                            switch 32768 {
                            case 127:
                                if (2147483647 > 127) {
                                    var v45 int64 = 417
                                    _ = v45
                                }
                            case -976:
                                v5 = (v5 % ((2) | 1))
                            case 1:
                                var v46 uint64 = 42
                                _ = v46
                            }
                        } else {
                            var v47 float64 = (0.1 + 0.1)
                            fmt.Println(v2)
                            if (0 <= 2) {
                                var v48 uint = (v5 | 223)
                                _ = v48
                            }
                            _ = v47
                        }
                    default:
                        v42 = v42
                    }
                case 0:
                    fmt.Println(v3, v3, v3)
                default:
                    if v3 {
                        if v14 {
                            var v49 int64 = 190
                            _ = v49
                        } else {
                            fmt.Println(v14)
                            v24 = v24
                            for fi50 := 0; fi50 < 15; fi50++ {
                                v51 := []int{(fi50 ^ (fi50 - 4294967295)), (fi50 >> (((v5 * v5)) & 63))}
                                v23 = int16(fi50)
                                for fi52 := 0; fi52 < 9; fi52++ {
                                    var v53 float32 = 1.5
                                    _ = fi52
                                    _ = v53
                                }
                                _ = fi50
                                _ = v51
                            }
                        }
                    } else {
                        v42 = ""
                        for fi54 := 0; fi54 < 9; fi54++ {
                            switch (fi54 >> (((v5 * (v5 | v5))) & 63)) {
                            case 0:
                                for fi55 := 0; fi55 < 4; fi55++ {
                                    v1 = (v1 << ((v5) & 63))
                                    _ = fi55
                                }
                            }
                            for fi56 := 0; fi56 < 10; fi56++ {
                                if (v42 >= v42) {
                                    var v57 string = (v42 + (((v42 + v42) + v42) + ((v42 + "foo bar") + (v42 + v42))))
                                    var v58 int = fi56
                                    v59 := []int{fi56, 578}
                                    _ = len(v4)
                                    v60 := GenS3{F0: v1, F1: (v1 - 90)}
                                    if len(v59) > 0 {
                                        v59[0] = 4294967295
                                    }
                                    v61 := GenS2{F0: (^(fi54)), F1: -2147483648}
                                    v58 += (fi56 % ((0) | 1))
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                    _ = v61
                                }
                                v42 = "a"
                                v62 := []int{fi56, (fi56 + (-(fi56)))}
                                _ = fi56
                                _ = v62
                            }
                            for fi63 := 0; fi63 < 20; fi63++ {
                                fmt.Println(v3, v14, v2)
                                _ = fi63
                            }
                            _ = fi54
                        }
                    }
                }
                v23 += (v23 | (v23 | v23))
                v24 = v24
            }
            if true {
                v23 -= -32768
                if v3 {
                    v64 := []int{9223372036854775807, -701, -862}
                    if (v2 >= v2) {
                        _ = len(v64)
                    } else {
                        fmt.Println(v3)
                    }
                    var v65 int = 285
                    _ = v64
                    _ = v65
                } else {
                    if false {
                        fmt.Println(v42)
                    }
                    fmt.Println(v1, v3)
                }
                if ((2 >= 2) || (v14 && (-1 < v23))) {
                    var v66 float32 = 0.1
                    fmt.Println(v66, v3)
                    if (0 != 0) {
                        if (((v42 + v42) == "\n\t") && v3) {
                            v66 += 0.1
                            v42 = ("foo bar" + v42)
                            var v67 int64 = 548
                            _ = v67
                        } else {
                            delete(v4, 0)
                            for fi68 := 0; fi68 < 18; fi68++ {
                                v69 := []int{fi68, -6}
                                v70 := GenS3{F0: -33, F1: (-(v1))}
                                _ = fi68
                                _ = v69
                                _ = v70
                            }
                            v71 := GenS4{F0: v1}
                            _ = v71
                        }
                    }
                    _ = v66
                } else {
                    fmt.Println(v1, v2, v24)
                }
            }
            _ = v42
        } else {
            v72 := GenS3{F0: v1, F1: v1}
            v73 := GenS2{F0: 514, F1: 667}
            _ = v72
            _ = v73
        }
        for fi74 := 0; fi74 < 9; fi74++ {
            v23 = int16(fi74)
            switch "" {
            case "foo bar":
                switch "\n\t" {
                case "hello":
                    v75 := GenS5{F0: -1, F1: 128}
                    _ = v75
                case "":
                    v76 := GenS6{F0: (^(fi74)), F1: ("foo bar" + "hello"), F2: v1}
                    _ = v76
                }
            case "hello":
                for fi77 := 0; fi77 < 16; fi77++ {
                    v78 := v4[2147483647]
                    _ = fi77
                    _ = v78
                }
            case "":
                fmt.Println(v14, fi74)
            default:
                var v79 float32 = 3.14159
                _ = v79
            }
            var v80 int16 = 640
            _ = fi74
            _ = v80
        }
    } else {
        var v81 uint8 = 218
        _ = v81
    }
    v82 := GenS4{F0: -81}
    v83 := []int{-688, -9223372036854775808, -991, 1}
    switch -9223372036854775808 {
    case 4294967295:
        switch 127 {
        case 0:
            var v84 uint8 = 205
            _ = v84
        }
    case 32768:
        v85 := GenS7{F0: 2147483648}
        _ = v85
    }
    if (3.14159 != 0.0) {
        var v86 string = ""
        var v87 int8 = v1
        _ = v86
        _ = v87
    } else {
        _ = len(v4)
        v3 = true
        for fi88 := 0; fi88 < 10; fi88++ {
            var v89 int = (fi88 / ((fi88) | 1))
            _ = fi88
            _ = v89
        }
    }
    var v90 string = ("foo bar" + ("foo bar" + "a"))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v83:", len(v83), v83)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v82
    _ = v83
    _ = v90
}
