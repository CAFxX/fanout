package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 uint16; F2 bool }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 uint16; F1 int }
type GenS7 struct { F0 int32; F1 uint; F2 uint16 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint8; F1 int16 }
func genFn0(p0 int8, p1 uint16) uint16 {
    p1 = 42
    if false {
        var v7 float64 = (0.1 * 1e10)
        fmt.Println(p1, p1)
        _ = v7
    } else {
        switch 283 {
        case -2147483648:
            for fi8 := 0; fi8 < 17; fi8++ {
                var v9 uint64 = 966
                v10 := GenS1{F0: ("foo bar" + "hello")}
                _ = fi8
                _ = v10
                _ = v9
            }
        case 9223372036854775807:
            if (!false) {
                v11 := GenS2{F0: 0, F1: 65535, F2: (!true)}
                _ = v11
            }
        }
        v12 := []int{2147483647, 127, 1}
        if true {
            fmt.Println(p1)
        }
        _ = v12
    }
    var v13 uint32 = 2147483648
    var v14 uint64 = 1
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return (p1 << ((446) & 63))
}
func genFn1() uint8 {
    v15 := GenS1{F0: ("foo bar" + ("hello" + ""))}
    for fi16 := 0; fi16 < 7; fi16++ {
        var v17 string = ""
        var v18 uint32 = 65535
        _ = fi16
        _ = v17
        _ = v18
    }
    _ = v15
    return 255
}

func main() {
    var v1 int8 = -1
    var v2 bool = (true || (!(2147483648 > 42)))
    v3 := []uint32{150, 2147483647, 738}
    v4 := make(map[int]string)
    var v5 float64 = 3.14159
    var v6 uint16 = 729
    v1 = -83
    if v2 {
        for fi19 := 0; fi19 < 8; fi19++ {
            _ = len(v3)
            var v20 uint = 2
            v21 := GenS3{F0: 42, F1: v1}
            _ = fi19
            _ = v20
            _ = v21
        }
        if ((("" >= ("\n\t" + ("\n\t" + "\n\t"))) || true) || v2) {
            v22 := GenS1{F0: "a"}
            v23 := GenS4{F0: (1000 >= 1000), F1: ""}
            _ = v22
            _ = v23
        } else {
            if true {
                v6 -= (v6 + (v6 / ((v6) | 1)))
                v2 = v2
                v2 = (!false)
            }
        }
    } else {
        if (4294967295 > 1) {
            v4[507] = ""
            var v24 int32 = 158
            if true {
                for fi25 := 0; fi25 < 2; fi25++ {
                    if v2 {
                        var v26 int8 = v1
                        _ = v26
                    } else {
                        for fi27 := 0; fi27 < 8; fi27++ {
                            switch int(v6) {
                            case 32768:
                                fmt.Println(v24)
                            case 476:
                                v5 += (((0.0 - float64(((3.14159 - ((-2.25 + 3.14159) * 1.5)) * (0.1 / (3.14159 * 1e10))))) + float64(0.1)) * (-2.25 / v5))
                            case -9223372036854775808:
                                v1 += (v1 * v1)
                            default:
                                if v2 {
                                    if v2 {
                                        var v28 int16 = 860
                                        _ = v28
                                    } else {
                                        v4[32767] = "a"
                                    }
                                }
                            }
                            fmt.Println(v6)
                            _ = fi27
                        }
                        v29 := GenS5{F0: "\n\t", F1: 18446744073709551615}
                        _ = v29
                    }
                    v30 := GenS5{F0: "\n\t", F1: 355}
                    _ = fi25
                    _ = v30
                }
                for fi31 := 0; fi31 < 18; fi31++ {
                    var v32 int32 = (v24 ^ -1)
                    var v33 int64 = 844
                    var v34 int16 = 0
                    _ = fi31
                    _ = v32
                    _ = v33
                    _ = v34
                }
            }
            _ = v24
        } else {
            v35 := GenS2{F0: 54, F1: (^(v6)), F2: false}
            _ = v35
        }
        switch 42 {
        case 2147483647:
            for fi36 := 0; fi36 < 14; fi36++ {
                fmt.Println(v1, fi36)
                v1 += -72
                _ = fi36
            }
        case 9223372036854775807:
            v37 := GenS2{F0: 0, F1: (v6 ^ (v6 ^ v6)), F2: ((110 == 0) && false)}
            _ = v37
        }
    }
    v38 := GenS6{F0: 2, F1: 32768}
    var v39 string = (("hello" + "hello") + ((("foo bar" + ("\n\t" + "hello")) + "foo bar") + "a"))
    for fi40 := 0; fi40 < 17; fi40++ {
        v6 = v6
        var v41 float32 = ((3.14159 / 1e10) - float32(-2.25))
        _ = len(v3)
        _ = fi40
        _ = v41
    }
    for fi42 := 0; fi42 < 5; fi42++ {
        var v43 int32 = -1
        for fi44 := 0; fi44 < 8; fi44++ {
            if v2 {
                v45 := GenS4{F0: false, F1: v39}
                switch v39 {
                case "a":
                    fmt.Println(v39)
                default:
                    v46 := GenS5{F0: v39, F1: 0}
                    fmt.Println(v39)
                    if v2 {
                        fmt.Println(fi44, fi42)
                        v2 = false
                        var v47 bool = (false && (true && (true || true)))
                        switch ("\n\t" + (v39 + ((v39 + ("foo bar" + v39)) + v39))) {
                        case "\n\t":
                            v48 := GenS4{F0: true, F1: v39}
                            _ = v48
                        }
                        var v49 uint32 = 1
                        if (32767 == (v6 >> ((uint(v6)) & 63))) {
                            var v50 int16 = -973
                            v5 = -2.25
                            v51 := GenS6{F0: genFn0(127, v6), F1: (fi42 ^ 32768)}
                            _ = v50
                            _ = v51
                        }
                        var v52 int8 = v1
                        _ = v47
                        _ = v49
                        _ = v52
                    }
                    _ = v46
                }
                switch (fi42 % (((fi42 + 2147483647)) | 1)) {
                case -9223372036854775808:
                    v53 := GenS7{F0: v43, F1: 255, F2: (v6 + v6)}
                    _ = v53
                case 32767:
                    switch (v39 + v39) {
                    case "":
                        var v54 float64 = 0.0
                        _ = v54
                    case "\n\t":
                        v39 = ""
                    default:
                        var v55 int32 = 520
                        _ = v55
                    }
                case 42:
                    fmt.Println(fi44, v5)
                }
                _ = v45
            } else {
                if (v2 || true) {
                    switch -9223372036854775808 {
                    case 772:
                        v56 := GenS8{F0: 9223372036854775807}
                        _ = v56
                    case -1:
                        fmt.Println(v6, v1)
                    case 2147483647:
                        switch v39 {
                        case "":
                            var v57 uint = 795
                            _ = v57
                        default:
                            v4[int(v6)] = "\n\t"
                        }
                    default:
                        v3 = append(v3, 4294967295)
                    }
                    delete(v4, fi42)
                    switch v39 {
                    case "\n\t":
                        var v58 uint8 = genFn1()
                        _ = v58
                    case "foo bar":
                        _ = len(v3)
                    case "hello":
                        v5 += ((0.0 - -2.25) - v5)
                    default:
                        fmt.Println(v39)
                    }
                }
                _ = len(v3)
                var v59 uint = 18446744073709551615
                _ = v59
            }
            var v60 float64 = v5
            fmt.Println(fi42)
            _ = fi44
            _ = v60
        }
        for fi61 := 0; fi61 < 7; fi61++ {
            fmt.Println(v1)
            _ = fi61
        }
        _ = fi42
        _ = v43
    }
    var v62 uint16 = v6
    if true {
        if (626 != 0) {
            for fi63 := 0; fi63 < 10; fi63++ {
                fmt.Println(v62, v62, v2)
                var v64 uint8 = 1
                if v2 {
                    var v65 uint = 719
                    v2 = true
                    v65 += v65
                    _ = v65
                } else {
                    v1 = -128
                    for fi66 := 0; fi66 < 17; fi66++ {
                        if (true || v2) {
                            var v67 uint64 = 18446744073709551615
                            var v68 float32 = 1.5
                            switch ((v39 + (("\n\t" + v39) + "hello")) + (v39 + v39)) {
                            case "foo bar":
                                var v69 float64 = v5
                                _ = v69
                            }
                            _ = v67
                            _ = v68
                        }
                        _ = fi66
                    }
                }
                _ = fi63
                _ = v64
            }
            var v70 int16 = -32768
            if v2 {
                fmt.Println(v39, v2, v2)
                v71 := GenS5{F0: v39, F1: 65535}
                if (!v2) {
                    var v72 bool = v2
                    v72 = (!v2)
                    switch v39 {
                    case "a":
                        v39 = v39
                    case "foo bar":
                        var v73 uint = 657
                        _ = v73
                    default:
                        for fi74 := 0; fi74 < 8; fi74++ {
                            var v75 string = v39
                            for fi76 := 0; fi76 < 7; fi76++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if v72 {
                                    var v77 string = "\n\t"
                                    fmt.Println(v62, v77)
                                    v4[(^(fi74))] = v77
                                    _ = v77
                                }
                                _ = fi76
                            }
                            fmt.Println(fi74, fi74)
                            _ = fi74
                            _ = v75
                        }
                    }
                    _ = v72
                } else {
                    if (v2 && (-2147483648 >= 0)) {
                        v5 = 1.5
                        var v78 int = -1
                        switch "a" {
                        case "hello":
                            fmt.Println(v5, v1)
                        case "a":
                            switch "foo bar" {
                            case "hello":
                                v6 -= genFn0(v1, (v6 + (^(v6))))
                            default:
                                var v79 uint32 = 714
                                _ = v79
                            }
                        case "\n\t":
                            for fi80 := 0; fi80 < 19; fi80++ {
                                if v2 {
                                    var v81 bool = v2
                                    _ = v81
                                }
                                _ = fi80
                            }
                        }
                        _ = v78
                    }
                    v82 := GenS9{F0: 2, F1: (-(v70))}
                    _ = v82
                }
                _ = v71
            }
            _ = v70
        }
    }
    v83 := GenS5{F0: (v39 + v39), F1: 4294967295}
    v84 := GenS8{F0: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v83
    _ = v84
}
