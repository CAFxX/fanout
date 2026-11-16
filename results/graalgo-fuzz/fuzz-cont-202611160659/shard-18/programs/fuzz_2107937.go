package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int64 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 uint64 }

func main() {
    var v1 uint64 = 0
    var v2 uint32 = 626
    v3 := make(map[int]uint32)
    v4 := make(map[int]int)
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]uint32)
    if true {
        if (127 > 0) {
            var v7 int16 = -600
            var v8 float32 = ((1.5 / -2.25) / (0.1 * 3.14159))
            switch 42 {
            case -2147483648:
                v9 := GenS1{F0: 116, F1: 1, F2: -9223372036854775808}
                _ = v9
            case 32767:
                v8 = v8
            case 42:
                var v10 bool = true
                _ = v10
            }
            _ = v7
            _ = v8
        } else {
            if true {
                for fi11 := 0; fi11 < 5; fi11++ {
                    var v12 int64 = -1
                    fmt.Println(fi11, v2, v1)
                    v12 -= (v12 + (v12 - v12))
                    _ = fi11
                    _ = v12
                }
            }
        }
        var v13 int64 = 1
        if (1.5 <= 0.0) {
            v13 -= 0
            v13 = 1
        } else {
            if (((42 <= 255) || (65535 == uint(v1))) && false) {
                v2 -= v2
                v5[v2] = 689
            }
            v14 := []int{}
            _ = v14
        }
        _ = v13
    }
    var v15 uint32 = v2
    switch 128 {
    case 128:
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 uint16 = 65535
            v2 = (v15 / ((v2) | 1))
            _ = fi16
            _ = v17
        }
    case 406:
        var v18 uint = 1
        _ = v18
    default:
        var v19 uint = 0
        _ = v19
    }
    v20 := GenS2{F0: ("a" + ("foo bar" + "hello")), F1: -1}
    delete(v6, v2)
    switch "\n\t" {
    case "\n\t":
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 int8 = 127
            for fi23 := 0; fi23 < 8; fi23++ {
                fmt.Println(v15, fi21, v2)
                v24 := []int{(fi21 * fi21), fi21, (fi23 % (((fi21 - (-(fi23)))) | 1))}
                fmt.Println(fi21)
                if (v1 != v1) {
                    if ("\n\t" >= "hello") {
                        if (!(955 != 32767)) {
                            v1 = (^(v1))
                            var v25 int8 = v22
                            _ = len(v6)
                            _ = v25
                        }
                    }
                } else {
                    var v26 int = -2147483648
                    var v27 int8 = v22
                    v2 = v15
                    switch "foo bar" {
                    case "":
                        if true {
                            var v28 uint = 255
                            if (!(true || true)) {
                                switch "\n\t" {
                                case "hello":
                                    fmt.Println(v15, v28, v15)
                                case "\n\t":
                                    if len(v24) > 0 {
                                        v24[0] = 27
                                    }
                                default:
                                    v15 = v15
                                }
                                var v29 uint8 = 42
                                _ = v29
                            } else {
                                if (242 > 1) {
                                    v3[955] = 853
                                    var v30 int8 = (v22 << (((^(v28))) & 63))
                                    _ = v30
                                } else {
                                    v26 += (fi21 + v26)
                                    v31 := v5[(v15 | (v2 - v15))]
                                    var v32 int64 = 0
                                    _ = v31
                                    _ = v32
                                }
                            }
                            if (0.1 < float64(1e10)) {
                                fmt.Println(v15)
                                var v33 int8 = (-(v27))
                                _ = v33
                            } else {
                                if true {
                                    var v34 string = ("foo bar" + ((("a" + ("hello" + "foo bar")) + "foo bar") + "a"))
                                    v26 += (fi21 + v26)
                                    _ = v34
                                }
                                v35 := GenS3{F0: ((1 != 82) || (!true)), F1: false}
                                _ = v35
                            }
                            switch "hello" {
                            case "foo bar":
                                delete(v4, (fi21 << ((1) & 63)))
                            case "hello":
                                if true {
                                    var v36 uint = (^(v28))
                                    v2 = (v15 << (((v36 | (v36 / ((v36) | 1)))) & 63))
                                    _ = v36
                                }
                            case "\n\t":
                                if true {
                                    switch ("hello" + "foo bar") {
                                    case "":
                                        if ((1 > -289) || true) {
                                            var v37 int8 = v27
                                            switch (fi21 % (((-(fi23))) | 1)) {
                                            case 32767:
                                                v38 := v3[fi23]
                                                _ = v38
                                            case -395:
                                                var v39 int8 = v22
                                                _ = v39
                                            }
                                            if true {
                                                var v40 int32 = 407
                                                var v41 float64 = 3.14159
                                                _ = v40
                                                _ = v41
                                            }
                                            _ = v37
                                        } else {
                                            var v42 uint16 = 0
                                            _ = v42
                                        }
                                    default:
                                        _ = len(v5)
                                    }
                                } else {
                                    var v43 uint32 = v15
                                    _ = v43
                                }
                            }
                            if ((3.14159 > 1e10) || false) {
                                v44 := GenS2{F0: "a", F1: -846}
                                _ = v44
                            } else {
                                v45 := GenS1{F0: v27, F1: 743, F2: 9223372036854775807}
                                _ = v45
                            }
                            var v46 int32 = 877
                            v22 = -1
                            _ = v28
                            _ = v46
                        } else {
                            v27 = (v27 | v22)
                            if false {
                                if false {
                                    var v47 float32 = (-2.25 - 3.14159)
                                    v26 += fi21
                                    _ = v47
                                } else {
                                    var v48 float32 = (1e10 / (0.1 + 0.0))
                                    _ = v48
                                }
                                v15 -= v15
                            } else {
                                if len(v24) > 0 {
                                    v24[0] = v26
                                }
                            }
                        }
                    case "a":
                        v1 = 461
                    case "hello":
                        if ((v2 * v2) >= 2147483648) {
                            var v49 bool = false
                            v27 = 42
                            v49 = (v49 || false)
                            if ("" > ("\n\t" + ("hello" + "foo bar"))) {
                                delete(v5, (^(v2)))
                            } else {
                                var v50 float64 = 0.1
                                var v51 int64 = int64(v27)
                                if ((v51 + (v51 / ((v51) | 1))) > v51) {
                                    fmt.Println(v26)
                                }
                                _ = v50
                                _ = v51
                            }
                            _ = v49
                        } else {
                            fmt.Println(v27)
                            switch fi21 {
                            case 772:
                                switch 4294967295 {
                                case -2147483648:
                                    var v52 int32 = 2147483647
                                    _ = v52
                                case 1:
                                    var v53 uint64 = (v1 + (v1 ^ v1))
                                    _ = v53
                                case 32767:
                                    if (((1.5 / float32(3.14159)) * 0.0) <= 0.1) {
                                        v54 := GenS2{F0: ("\n\t" + "foo bar"), F1: 42}
                                        fmt.Println(v1)
                                        var v55 int32 = -1
                                        switch "\n\t" {
                                        case "":
                                            switch v26 {
                                            case 0:
                                                v55 -= (v55 << ((4294967295) & 63))
                                            default:
                                                v56 := GenS3{F0: false, F1: ("a" < ("foo bar" + "\n\t"))}
                                                _ = v56
                                            }
                                        case "foo bar":
                                            v1 = v1
                                        case "\n\t":
                                            v57 := GenS4{F0: -162, F1: 18}
                                            _ = v57
                                        default:
                                            var v58 uint8 = 1
                                            _ = v58
                                        }
                                        _ = v54
                                        _ = v55
                                    } else {
                                        fmt.Println(v27)
                                        v59 := v5[(v15 | v15)]
                                        _ = v59
                                    }
                                default:
                                    v27 = (-(v27))
                                    fmt.Println(v27, v15, fi23)
                                    fmt.Println(v1, v15)
                                }
                            default:
                                delete(v6, (v2 << ((4294967295) & 63)))
                            }
                        }
                        switch "hello" {
                        case "hello":
                            switch fi23 {
                            case 2147483647:
                                var v60 float32 = 1e10
                                _ = v60
                            }
                        case "foo bar":
                            switch 9223372036854775807 {
                            case 2147483647:
                                if true {
                                    if true {
                                        if (3.14159 > 0.1) {
                                            var v61 uint32 = v2
                                            v26 = int(v27)
                                            if ((2 < 2) || true) {
                                                var v62 bool = (!(false || (!true)))
                                                if true {
                                                    v63 := v6[0]
                                                    var v64 int32 = 2147483647
                                                    _ = v63
                                                    _ = v64
                                                } else {
                                                    v65 := GenS2{F0: ("foo bar" + "a"), F1: 543}
                                                    _ = v65
                                                }
                                                v27 = -128
                                                var v66 uint8 = 0
                                                var v67 float64 = (float64((1.5 * 0.0)) + 1e10)
                                                _ = v62
                                                _ = v66
                                                _ = v67
                                            } else {
                                                v22 += v27
                                            }
                                            _ = v61
                                        } else {
                                            fmt.Println(v15)
                                            v24 = append(v24, (fi23 + (fi23 / ((438) | 1))))
                                        }
                                        if len(v24) > 0 {
                                            v24[0] = v26
                                        }
                                        v68 := GenS4{F0: v26, F1: 136}
                                        _ = v68
                                    }
                                    var v69 string = ""
                                    if (false && (((v69 + v69) < v69) || false)) {
                                        var v70 int32 = -676
                                        v69 = ""
                                        if (true && false) {
                                            _ = len(v3)
                                            v70 = (v70 & v70)
                                        }
                                        _ = v70
                                    }
                                    _ = v69
                                }
                                v22 = (v22 / (((^(v22))) | 1))
                            default:
                                v22 -= v27
                            }
                        }
                    }
                    var v71 int = (v26 << ((18446744073709551615) & 63))
                    _ = v26
                    _ = v27
                    _ = v71
                }
                _ = fi23
                _ = v24
            }
            _ = fi21
            _ = v22
        }
    }
    for fi72 := 0; fi72 < 6; fi72++ {
        v6[(v15 | (v2 - v15))] = (v15 * v2)
        _ = fi72
    }
    _ = len(v4)
    switch -1 {
    case -936:
        for fi73 := 0; fi73 < 18; fi73++ {
            switch "\n\t" {
            case "\n\t":
                if false {
                    v15 = (^(v2))
                    if true {
                        fmt.Println(fi73, v15)
                    }
                } else {
                    v1 -= (v1 - 18446744073709551615)
                    v15 += (^(v2))
                }
            }
            switch fi73 {
            case 0:
                if (!(false && (1 != 1))) {
                    delete(v6, v2)
                    v15 += v2
                    var v74 bool = (!false)
                    _ = v74
                }
            case 2147483647:
                var v75 uint8 = 9
                _ = v75
            case 32768:
                for fi76 := 0; fi76 < 15; fi76++ {
                    var v77 uint8 = 42
                    var v78 bool = true
                    _ = fi76
                    _ = v77
                    _ = v78
                }
            default:
                var v79 int64 = 9223372036854775807
                _ = v79
            }
            v15 -= v15
            _ = fi73
        }
    case -1:
        fmt.Println(v15, v15, v2)
    case 2147483647:
        v2 -= v2
    }
    var v80 int8 = 41
    v81 := []int{}
    v1 += v1
    v82 := GenS5{F0: v1}
    v2 += v15
    if len(v81) >= 2 {
        v81 = v81[1:len(v81)]
    }
    v1 = (v1 + v1)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v80:", v80)
    fmt.Println("slicelen v81:", len(v81), v81)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v80
    _ = v81
    _ = v82
}
