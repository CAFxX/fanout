package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int8; F1 int8; F2 int64 }
type GenS3 struct { F0 int64; F1 uint32; F2 bool }
type GenS4 struct { F0 int64; F1 int8; F2 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int32; F1 string }
type GenS7 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint64 }

func main() {
    var v1 int64 = 42
    var v2 uint = 942
    var v3 int64 = (v1 - 9223372036854775807)
    v4 := []int{42, 9223372036854775807, 32768, 32767}
    v5 := make(map[uint32]bool)
    var v6 float64 = 0.1
    v7 := v5[557]
    if (!false) {
        var v8 int8 = 0
        var v9 int64 = (v1 << ((v2) & 63))
        v6 = (0.0 + ((3.14159 - 0.1) * v6))
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        if false {
            switch fi10 {
            case -1:
                v11 := GenS1{F0: 1, F1: 32767}
                _ = v11
            case -382:
                v12 := GenS1{F0: v1, F1: 32767}
                _ = v12
            case 32768:
                for fi13 := 0; fi13 < 12; fi13++ {
                    var v14 uint64 = 0
                    if v7 {
                        fmt.Println(v2, v1)
                    }
                    _ = fi13
                    _ = v14
                }
            default:
                fmt.Println(v3)
            }
            switch fi10 {
            case 32767:
                v15 := GenS1{F0: 0, F1: 32767}
                _ = v15
            case 2147483647:
                switch 128 {
                case -164:
                    if v7 {
                        v6 += float64(0.1)
                        fmt.Println(v3)
                        fmt.Println(v3)
                    }
                case 128:
                    if true {
                        v16 := GenS1{F0: -9223372036854775808, F1: -685}
                        _ = v16
                    }
                case 9223372036854775807:
                    switch "" {
                    case "a":
                        _ = len(v4)
                    case "":
                        v5[171] = v7
                    case "hello":
                        fmt.Println(v7, v3, fi10)
                    default:
                        _ = len(v5)
                    }
                }
            default:
                for fi17 := 0; fi17 < 14; fi17++ {
                    fmt.Println(v2)
                    _ = fi17
                }
            }
        }
        v18 := GenS2{F0: 42, F1: 1, F2: (-(v1))}
        _ = fi10
        _ = v18
    }
    v2 += 2
    var v19 int16 = 32767
    v3 -= (v1 % (((v1 & v3)) | 1))
    switch 128 {
    case 4294967295:
        v6 = v6
    case -429:
        v6 = 1.5
    case 128:
        v7 = (1 == 128)
    default:
        for fi20 := 0; fi20 < 20; fi20++ {
            if v7 {
                if v7 {
                    var v21 string = ""
                    v22 := GenS2{F0: 0, F1: -128, F2: v3}
                    var v23 int8 = 127
                    _ = v21
                    _ = v22
                    _ = v23
                }
                v24 := GenS3{F0: (v1 / (((v3 | v3)) | 1)), F1: 2147483648, F2: (v2 == 18446744073709551615)}
                fmt.Println(v1, v6, v1)
                _ = v24
            } else {
                v7 = v7
                var v25 uint = 42
                _ = v25
            }
            _ = fi20
        }
    }
    v1 = (v3 >> ((0) & 63))
    for fi26 := 0; fi26 < 13; fi26++ {
        switch "\n\t" {
        case "hello":
            v27 := GenS4{F0: v1, F1: 0, F2: v19}
            _ = v27
        case "a":
            v1 += (v1 ^ v3)
        default:
            var v28 string = "foo bar"
            _ = v28
        }
        switch "" {
        case "foo bar":
            for fi29 := 0; fi29 < 3; fi29++ {
                v7 = ((^(fi26)) >= (fi29 + fi26))
                _ = fi29
            }
        case "":
            switch 32768 {
            case 714:
                v6 = v6
            case 128:
                var v30 uint = 2
                _ = v30
            }
        case "hello":
            var v31 int16 = v19
            _ = v31
        default:
            for fi32 := 0; fi32 < 5; fi32++ {
                v33 := GenS5{F0: ((((("foo bar" + "hello") + "foo bar") + "\n\t") + ((("hello" + "foo bar") + "\n\t") + "\n\t")) + "foo bar")}
                switch 127 {
                case 198:
                    v34 := GenS6{F0: 42, F1: "hello"}
                    _ = v34
                case -850:
                    var v35 int8 = 127
                    if (v7 || (!false)) {
                        v4 = append(v4, (fi26 + (^(fi26))))
                        var v36 int16 = v19
                        var v37 uint32 = 2147483648
                        _ = v36
                        _ = v37
                    }
                    _ = v35
                }
                _ = fi32
                _ = v33
            }
        }
        _ = fi26
    }
    var v38 float64 = v6
    v3 -= v3
    var v39 uint16 = 42
    if false {
        if v7 {
            if false {
                for fi40 := 0; fi40 < 4; fi40++ {
                    for fi41 := 0; fi41 < 3; fi41++ {
                        if ((v2 ^ (^(v2))) != (v2 / ((v2) | 1))) {
                            var v42 int16 = 1000
                            _ = v42
                        } else {
                            var v43 bool = v7
                            v44 := GenS5{F0: "foo bar"}
                            _ = v43
                            _ = v44
                        }
                        var v45 int64 = (v3 >> (((v2 & 42)) & 63))
                        var v46 uint16 = v39
                        _ = fi41
                        _ = v45
                        _ = v46
                    }
                    if v7 {
                        var v47 int = (fi40 / (((fi40 >> ((v2) & 63))) | 1))
                        v47 += (fi40 ^ (fi40 | (fi40 >> (((v2 | uint(v39))) & 63))))
                        _ = v47
                    } else {
                        var v48 float32 = -2.25
                        if len(v4) > 0 {
                            v4[0] = (^(fi40))
                        }
                        for fi49 := 0; fi49 < 16; fi49++ {
                            v6 = v6
                            var v50 int32 = -1
                            _ = fi49
                            _ = v50
                        }
                        _ = v48
                    }
                    _ = fi40
                }
            } else {
                _ = len(v5)
                var v51 uint16 = v39
                _ = v51
            }
        } else {
            if (false || true) {
                switch ("" + "") {
                case "\n\t":
                    switch "hello" {
                    case "a":
                        if ((v39 / ((v39) | 1)) == 65535) {
                            var v52 string = ""
                            delete(v5, 965)
                            _ = v52
                        } else {
                            fmt.Println(v19, v6, v2)
                            var v53 int32 = -2147483648
                            _ = len(v4)
                            _ = v53
                        }
                    default:
                        if true {
                            for fi54 := 0; fi54 < 11; fi54++ {
                                var v55 float64 = (3.14159 * (v6 / ((float64(1e10) * (v6 * (((0.0 / v6) + 1.5) * (v6 - ((1e10 * v38) / ((v38 / v6) - v38)))))) + 3.14159)))
                                for fi56 := 0; fi56 < 13; fi56++ {
                                    fmt.Println(fi54, fi56, fi54)
                                    _ = fi56
                                }
                                _ = fi54
                                _ = v55
                            }
                        }
                    }
                default:
                    if (1 > 208) {
                        v2 -= (v2 ^ v2)
                    }
                }
            }
        }
        if false {
            var v57 int16 = (^(v19))
            v58 := GenS7{F0: (v39 >> ((uint(v39)) & 63)), F1: 65535, F2: 1}
            var v59 bool = (!(932 != 2147483648))
            _ = v57
            _ = v58
            _ = v59
        }
    }
    v7 = (!(!false))
    v39 = (v39 % (((^(v39))) | 1))
    if ((v3 >> ((v2) & 63)) <= v1) {
        var v60 int8 = 36
        switch "hello" {
        case "\n\t":
            fmt.Println(v19, v60)
        case "hello":
            switch "hello" {
            case "":
                v2 -= (v2 - 0)
            case "\n\t":
                v39 = (v39 - v39)
            default:
                if (18446744073709551615 > v2) {
                    var v61 int16 = 927
                    var v62 uint16 = v39
                    _ = v61
                    _ = v62
                } else {
                    var v63 int = 128
                    if (((v1 >> (((v2 >> (((^(v2))) & 63))) & 63)) >= v3) && v7) {
                        fmt.Println(v6, v2, v39)
                    } else {
                        switch ("\n\t" + "\n\t") {
                        case "\n\t":
                            if (!v7) {
                                for fi64 := 0; fi64 < 11; fi64++ {
                                    for fi65 := 0; fi65 < 19; fi65++ {
                                        if v7 {
                                            var v66 uint16 = v39
                                            var v67 uint32 = 0
                                            v68 := GenS8{F0: (v60 >> ((194) & 63))}
                                            if (255 <= 2) {
                                                var v69 string = (("foo bar" + "") + "a")
                                                v60 = 127
                                                _ = v69
                                            }
                                            var v70 int16 = v19
                                            _ = v66
                                            _ = v67
                                            _ = v68
                                            _ = v70
                                        }
                                        v71 := v5[4294967295]
                                        v60 = v60
                                        _ = fi65
                                        _ = v71
                                    }
                                    _ = fi64
                                }
                            } else {
                                v72 := GenS1{F0: v1, F1: -32768}
                                fmt.Println(v2, v19)
                                v73 := GenS7{F0: v39, F1: 65535, F2: 69}
                                _ = v72
                                _ = v73
                            }
                        }
                        var v74 uint32 = 0
                        switch (v63 * 2147483647) {
                        case 0:
                            if v7 {
                                switch "\n\t" {
                                case "a":
                                    var v75 int16 = v19
                                    _ = v75
                                }
                                for fi76 := 0; fi76 < 8; fi76++ {
                                    for fi77 := 0; fi77 < 20; fi77++ {
                                        var v78 int64 = 0
                                        _ = fi77
                                        _ = v78
                                    }
                                    _ = fi76
                                }
                            }
                        case 66:
                            if (true || (v7 && false)) {
                                if ((v74 % ((v74) | 1)) < (^(v74))) {
                                    var v79 float64 = (((float64(0.1) - (float64(1e10) / 1.0)) - v38) / v38)
                                    v80 := GenS9{F0: 32}
                                    var v81 string = "\n\t"
                                    _ = v79
                                    _ = v80
                                    _ = v81
                                } else {
                                    var v82 uint = v2
                                    v2 += v82
                                    var v83 int32 = -2147483648
                                    _ = v82
                                    _ = v83
                                }
                            } else {
                                var v84 bool = ((!false) || (42 > v60))
                                var v85 int = (v63 >> (((v2 ^ (v2 + 361))) & 63))
                                v63 = (v85 + (v85 ^ v85))
                                _ = v84
                                _ = v85
                            }
                        }
                        _ = v74
                    }
                    _ = v63
                }
            }
        case "":
            v3 += (-(v3))
        default:
            v86 := GenS7{F0: (v39 - v39), F1: 1, F2: 0}
            _ = v86
        }
        _ = v60
    } else {
        v87 := GenS5{F0: "foo bar"}
        var v88 uint32 = 2
        v6 -= (float64(((1e10 * -2.25) / (0.1 + (float32(v38) - 3.14159)))) + float64(((-2.25 / -2.25) + float32(v6))))
        _ = v87
        _ = v88
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
