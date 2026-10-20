package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int32 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint; F1 string; F2 bool }
type GenS7 struct { F0 bool; F1 uint16 }
type GenS8 struct { F0 uint64; F1 uint }
type GenS9 struct { F0 uint32 }
func genFn0(p0 string, p1 uint64, p2 string) int16 {
    v4 := GenS1{F0: 0, F1: p0, F2: -852}
    v5 := GenS2{F0: 42, F1: 664}
    p2 = "\n\t"
    v6 := []int{-2147483648, 4294967295, 42, 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 47
}
func genFn1(p0 int16) bool {
    for fi7 := 0; fi7 < 20; fi7++ {
        switch fi7 {
        case 9223372036854775807:
            v8 := []int{0, 943}
            _ = v8
        case 0:
            v9 := GenS2{F0: 776, F1: 4294967295}
            _ = v9
        case 32767:
            var v10 float64 = 1.5
            _ = v10
        }
        p0 = -894
        _ = fi7
    }
    switch "a" {
    case "hello":
        for fi11 := 0; fi11 < 11; fi11++ {
            if true {
                v12 := GenS3{F0: 42, F1: 111, F2: 189}
                v13 := GenS2{F0: 406, F1: 65535}
                _ = v12
                _ = v13
            } else {
                p0 += (p0 | (^(p0)))
                p0 = p0
                p0 += genFn0(("\n\t" + "hello"), 2, "foo bar")
            }
            p0 += (p0 % ((p0) | 1))
            v14 := []int{(fi11 * (fi11 << ((2) & 63))), (-(fi11)), 128}
            _ = fi11
            _ = v14
        }
    }
    p0 = 46
    fmt.Println(p0)
    _ = p0
    return (32768 > 1)
}

func main() {
    var v1 uint8 = 128
    var v2 int = 1
    var v3 bool = false
    var v15 int64 = -1
    v15 = -107
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 int = v2
        if v3 {
            switch ((("hello" + "hello") + "foo bar") + "\n\t") {
            case "a":
                v18 := []int{v17, (v17 & (v2 / (((v17 & (v2 ^ 32767))) | 1))), v17}
                _ = v18
            case "\n\t":
                fmt.Println(v2, v3, v2)
            default:
                v2 -= fi16
            }
            v19 := GenS4{F0: 18446744073709551615}
            _ = v19
        } else {
            fmt.Println(fi16, v1, v2)
            var v20 uint64 = 18446744073709551615
            var v21 uint8 = (v1 >> ((18446744073709551615) & 63))
            _ = v20
            _ = v21
        }
        _ = fi16
        _ = v17
    }
    v15 = v15
    v22 := []int{128, v2}
    if v3 {
        v15 -= (v15 + (v15 - -1))
    } else {
        if (4294967295 == 155) {
            v15 = 9223372036854775807
            v23 := GenS5{F0: (^(v1))}
            _ = v23
        }
        switch -2147483648 {
        case 0:
            if (!v3) {
                if len(v22) > 0 {
                    v22[0] = v2
                }
            }
        case 333:
            var v24 string = (("" + (("a" + "foo bar") + ("foo bar" + ""))) + "foo bar")
            _ = v24
        case 9223372036854775807:
            for fi25 := 0; fi25 < 12; fi25++ {
                if v3 {
                    v26 := GenS3{F0: 0, F1: 2, F2: 4294967295}
                    for fi27 := 0; fi27 < 1; fi27++ {
                        v28 := GenS2{F0: v15, F1: 454}
                        _ = fi27
                        _ = v28
                    }
                    v29 := GenS6{F0: 4294967295, F1: ("\n\t" + "hello"), F2: v3}
                    _ = v26
                    _ = v29
                }
                if (v3 && ((-26 > 2147483647) || true)) {
                    var v30 bool = true
                    _ = v30
                } else {
                    switch ("hello" + "a") {
                    case "foo bar":
                        fmt.Println(v3, v3)
                    default:
                        var v31 uint64 = 471
                        _ = v31
                    }
                }
                if (197 != 65535) {
                    switch ("\n\t" + "foo bar") {
                    case "hello":
                        fmt.Println(v15)
                    case "\n\t":
                        for fi32 := 0; fi32 < 12; fi32++ {
                            switch v2 {
                            case -720:
                                var v33 int = 32767
                                _ = v33
                            case 128:
                                if (((float64(1.5) / 1e10) * -2.25) < 0.0) {
                                    var v34 int8 = 27
                                    _ = v34
                                } else {
                                    switch ((((("hello" + "\n\t") + "a") + "a") + ((("a" + "hello") + "foo bar") + ("" + ""))) + "hello") {
                                    case "\n\t":
                                        var v35 int16 = 1000
                                        _ = v35
                                    case "":
                                        fmt.Println(v15)
                                    case "hello":
                                        v15 = (v15 >> ((65535) & 63))
                                    }
                                    var v36 float32 = 1e10
                                    var v37 float64 = (-2.25 * float64(v36))
                                    v37 -= 0.0
                                    _ = v36
                                    _ = v37
                                }
                            case -2147483648:
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                            }
                            var v38 uint64 = 246
                            _ = fi32
                            _ = v38
                        }
                    case "a":
                        switch fi25 {
                        case 128:
                            switch "foo bar" {
                            case "hello":
                                switch (("hello" + "") + "") {
                                case "hello":
                                    switch "\n\t" {
                                    case "":
                                        fmt.Println(v3)
                                    case "hello":
                                        var v39 uint64 = 42
                                        _ = v39
                                    }
                                case "\n\t":
                                    _ = len(v22)
                                case "foo bar":
                                    v3 = false
                                }
                            default:
                                if v3 {
                                    if ((0 <= v1) || false) {
                                        if v3 {
                                            if false {
                                                if ((v3 || (!v3)) && ((v2 % (((fi25 * (v2 / (((fi25 % ((fi25) | 1))) | 1)))) | 1)) >= v2)) {
                                                    v22 = append(v22, 128)
                                                    _ = len(v22)
                                                    var v40 float64 = (1.5 * 0.0)
                                                    _ = v40
                                                }
                                                if false {
                                                    fmt.Println(v1, v3)
                                                    var v41 uint32 = 42
                                                    var v42 int64 = (v15 ^ v15)
                                                    _ = v41
                                                    _ = v42
                                                }
                                                for fi43 := 0; fi43 < 11; fi43++ {
                                                    if true {
                                                        if (-9223372036854775808 >= v15) {
                                                            var v44 bool = (82 != 2)
                                                            _ = v44
                                                        }
                                                    }
                                                    v45 := GenS4{F0: 42}
                                                    v46 := GenS6{F0: 18446744073709551615, F1: "a", F2: (!true)}
                                                    _ = fi43
                                                    _ = v45
                                                    _ = v46
                                                }
                                            }
                                            v22 = append(v22, (v2 | v2))
                                            if true {
                                                v2 += 962
                                            }
                                        }
                                    } else {
                                        if v3 {
                                            v15 = (v15 / (((v15 << ((4294967295) & 63))) | 1))
                                        } else {
                                            for fi47 := 0; fi47 < 8; fi47++ {
                                                switch "" {
                                                case "a":
                                                    switch "" {
                                                    case "foo bar":
                                                        fmt.Println(v2, fi25, v2)
                                                    case "a":
                                                        fmt.Println(v3)
                                                    }
                                                case "":
                                                    if v3 {
                                                        v2 += v2
                                                        var v48 int = 32767
                                                        _ = v48
                                                    } else {
                                                        v15 = (v15 & (v15 & v15))
                                                        var v49 int64 = 932
                                                        v15 -= 9223372036854775807
                                                        _ = v49
                                                    }
                                                }
                                                v3 = false
                                                _ = fi47
                                            }
                                            v15 = (v15 / ((v15) | 1))
                                        }
                                        var v50 float32 = 3.14159
                                        _ = v50
                                    }
                                    v15 -= (v15 & v15)
                                }
                            }
                        case -215:
                            v3 = v3
                        default:
                            if true {
                                v2 += v2
                                if v3 {
                                    if (v3 || v3) {
                                        for fi51 := 0; fi51 < 7; fi51++ {
                                            fmt.Println(v2, v1, v15)
                                            v3 = false
                                            if (673 <= 2) {
                                                var v52 uint32 = 42
                                                if false {
                                                    switch fi51 {
                                                    case -9223372036854775808:
                                                        v52 = 2
                                                    case 2147483647:
                                                        v15 = (v15 + -629)
                                                    default:
                                                        v53 := GenS7{F0: (!false), F1: 465}
                                                        _ = v53
                                                    }
                                                }
                                                _ = v52
                                            } else {
                                                v3 = v3
                                                v54 := GenS8{F0: 610, F1: 255}
                                                v3 = (0 <= 1000)
                                                _ = v54
                                            }
                                            fmt.Println(fi25, v2)
                                            _ = fi51
                                        }
                                        var v55 uint16 = 32767
                                        v22 = append(v22, v2)
                                        _ = v55
                                    }
                                    var v56 uint32 = 935
                                    v3 = v3
                                    _ = v56
                                } else {
                                    v2 = (fi25 % ((-2147483648) | 1))
                                    v22 = append(v22, v2)
                                }
                            } else {
                                if (2147483647 == -673) {
                                    for fi57 := 0; fi57 < 8; fi57++ {
                                        if v3 {
                                            var v58 int16 = 0
                                            if v3 {
                                                if len(v22) > 0 {
                                                    v22[0] = (v2 + fi25)
                                                }
                                                _ = len(v22)
                                            } else {
                                                v58 = genFn0("hello", 553, "a")
                                                if len(v22) > 0 {
                                                    v22[0] = fi25
                                                }
                                                v59 := GenS6{F0: 4294967295, F1: "a", F2: v3}
                                                _ = v59
                                            }
                                            _ = v58
                                        }
                                        _ = fi57
                                    }
                                }
                                var v60 string = ("\n\t" + "")
                                _ = v60
                            }
                        }
                    default:
                        v61 := GenS3{F0: 143, F1: 1, F2: 2}
                        _ = v61
                    }
                    var v62 uint32 = 455
                    _ = v62
                } else {
                    v1 = (v1 + v1)
                    for fi63 := 0; fi63 < 6; fi63++ {
                        v64 := GenS9{F0: 47}
                        _ = fi63
                        _ = v64
                    }
                }
                _ = fi25
            }
        default:
            switch "a" {
            case "hello":
                var v65 float64 = 0.0
                _ = v65
            default:
                var v66 int = (v2 * (v2 / ((v2) | 1)))
                _ = v66
            }
        }
    }
    var v67 uint8 = (v1 | 42)
    fmt.Println(v67)
    var v68 int16 = 94
    if false {
        v3 = v3
        for fi69 := 0; fi69 < 3; fi69++ {
            v2 += fi69
            v70 := GenS5{F0: (v67 * (^(v67)))}
            v68 -= 1000
            _ = fi69
            _ = v70
        }
    }
    v15 += 9223372036854775807
    for fi71 := 0; fi71 < 18; fi71++ {
        v2 += (v2 << ((42) & 63))
        switch "hello" {
        case "\n\t":
            v72 := GenS3{F0: 778, F1: 1, F2: 48}
            _ = v72
        default:
            _ = len(v22)
        }
        if false {
            var v73 int = v2
            _ = v73
        } else {
            var v74 uint16 = 221
            _ = v74
        }
        _ = fi71
    }
    v75 := GenS3{F0: 391, F1: 4294967295, F2: 2}
    v15 -= v15
    var v76 int = 32768
    v76 = v2
    for fi77 := 0; fi77 < 19; fi77++ {
        for fi78 := 0; fi78 < 3; fi78++ {
            v79 := GenS5{F0: (v1 + v67)}
            _ = fi78
            _ = v79
        }
        _ = fi77
    }
    var v80 int8 = -102
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v76:", v76)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v67
    _ = v68
    _ = v75
    _ = v76
    _ = v80
}
