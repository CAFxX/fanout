package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64; F1 uint32 }
type GenS7 struct { F0 string; F1 int64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint8; F1 int16 }
type GenS10 struct { F0 int; F1 int; F2 uint16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 32768
    var v3 float64 = -2.25
    fmt.Println(v2, v3)
    if true {
        for fi4 := 0; fi4 < 13; fi4++ {
            switch (^(fi4)) {
            case 32767:
                fmt.Println(fi4, v2)
            case 0:
                v3 -= ((v3 + 0.1) * v3)
            case 127:
                var v5 uint8 = 1
                _ = v5
            }
            if false {
                if (true || true) {
                    if ((((!(0 <= -6)) && (1 <= -9223372036854775808)) || true) && (-9 > 127)) {
                        v3 += (float64(3.14159) + v3)
                        for fi6 := 0; fi6 < 19; fi6++ {
                            v3 = ((float64(((float32(1.5) + float32(1.5)) * 0.0)) - float64(0.0)) + v3)
                            _ = fi6
                        }
                    } else {
                        switch "a" {
                        case "foo bar":
                            var v7 uint8 = 143
                            _ = v7
                        case "":
                            switch "hello" {
                            case "foo bar":
                                v8 := GenS1{F0: 2, F1: ("" + "foo bar")}
                                _ = v8
                            default:
                                if true {
                                    switch "hello" {
                                    case "\n\t":
                                        for fi9 := 0; fi9 < 8; fi9++ {
                                            v2 += fi9
                                            v10 := []int{fi9}
                                            v3 += (0.0 * float64(-2.25))
                                            _ = fi9
                                            _ = v10
                                        }
                                    case "foo bar":
                                        var v11 int64 = 888
                                        _ = v11
                                    default:
                                        v12 := GenS2{F0: v2}
                                        _ = v12
                                    }
                                }
                            }
                        }
                    }
                } else {
                    v3 += v3
                    v2 += (fi4 << ((18446744073709551615) & 63))
                }
                for fi13 := 0; fi13 < 12; fi13++ {
                    switch (fi4 << ((0) & 63)) {
                    case 32767:
                        if (32767 >= -720) {
                            if (477 != 0) {
                                v14 := []int{fi4, v2, fi13}
                                v3 -= v3
                                _ = v14
                            } else {
                                var v15 int64 = 42
                                fmt.Println(v2)
                                var v16 float64 = 0.1
                                _ = v15
                                _ = v16
                            }
                            fmt.Println(fi4)
                        }
                    case -2147483648:
                        v17 := v1[(^(v2))]
                        _ = v17
                    case 2147483647:
                        v18 := []int{v2}
                        _ = v18
                    default:
                        v3 = ((v3 - float64(-2.25)) * 1.5)
                    }
                    v19 := GenS2{F0: -2147483648}
                    v2 = fi4
                    _ = fi13
                    _ = v19
                }
            } else {
                v20 := []int{278, 1, v2, fi4}
                switch "\n\t" {
                case "\n\t":
                    switch "a" {
                    case "\n\t":
                        v2 = v2
                    case "foo bar":
                        fmt.Println(v3, fi4)
                    case "a":
                        var v21 bool = (((v3 - 3.14159) / 3.14159) > v3)
                        _ = v21
                    default:
                        fmt.Println(v3, v3, v2)
                    }
                case "foo bar":
                    var v22 int16 = 42
                    _ = v22
                }
                if ("" == (("\n\t" + "hello") + ("hello" + (("foo bar" + ("" + ("" + "foo bar"))) + "\n\t")))) {
                    fmt.Println(v3, v3, v2)
                    v20 = append(v20, 0)
                } else {
                    fmt.Println(fi4, v3, fi4)
                    for fi23 := 0; fi23 < 18; fi23++ {
                        var v24 int8 = 104
                        var v25 int32 = 42
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    v3 -= ((-2.25 / 1e10) + v3)
                }
                _ = v20
            }
            _ = fi4
        }
    } else {
        if false {
            v26 := GenS3{F0: 14, F1: ""}
            switch "\n\t" {
            case "a":
                v27 := GenS4{F0: 27}
                _ = v27
            }
            _ = v26
        }
    }
    var v28 int = v2
    v3 = (float64((1.5 - 1.5)) - v3)
    if (true || (1 == 42)) {
        for fi29 := 0; fi29 < 12; fi29++ {
            var v30 float64 = v3
            _ = fi29
            _ = v30
        }
        v3 += (v3 / v3)
    } else {
        v31 := GenS5{F0: ("hello" <= ("\n\t" + ("\n\t" + "a")))}
        if true {
            v32 := GenS4{F0: 169}
            _ = v32
        } else {
            v1[v2] = 0
        }
        _ = v31
    }
    v1[239] = 1
    var v33 float32 = (float32((v3 + v3)) * 3.14159)
    v34 := GenS6{F0: 73, F1: 42}
    switch v2 {
    case 1:
        var v35 int = v28
        _ = v35
    default:
        var v36 uint = 0
        _ = v36
    }
    for fi37 := 0; fi37 < 3; fi37++ {
        delete(v1, (^(fi37)))
        _ = fi37
    }
    var v38 float64 = (float64(-2.25) * (v3 / v3))
    v2 -= v2
    v38 = v3
    fmt.Println(v33, v33, v2)
    if ("\n\t" > "hello") {
        if false {
            for fi39 := 0; fi39 < 11; fi39++ {
                var v40 int = 127
                delete(v1, -12)
                var v41 int = fi39
                _ = fi39
                _ = v40
                _ = v41
            }
            for fi42 := 0; fi42 < 3; fi42++ {
                v43 := GenS7{F0: ("hello" + "hello"), F1: 1}
                var v44 uint64 = 4294967295
                var v45 uint = 18446744073709551615
                _ = fi42
                _ = v43
                _ = v44
                _ = v45
            }
        }
        for fi46 := 0; fi46 < 11; fi46++ {
            if true {
                v47 := GenS2{F0: v28}
                _ = v47
            } else {
                for fi48 := 0; fi48 < 10; fi48++ {
                    if false {
                        if false {
                            if true {
                                v49 := GenS4{F0: 2}
                                fmt.Println(v2, v38)
                                _ = v49
                            }
                            v50 := v1[(fi48 / ((v28) | 1))]
                            _ = v50
                        } else {
                            fmt.Println(v38)
                            var v51 uint8 = 172
                            delete(v1, fi46)
                            _ = v51
                        }
                        var v52 uint8 = 239
                        v53 := GenS6{F0: 18446744073709551615, F1: 57}
                        _ = v52
                        _ = v53
                    } else {
                        if ((fi46 / (((v2 / (((fi48 + (v28 >> ((0) & 63)))) | 1))) | 1)) < (fi46 >> ((18446744073709551615) & 63))) {
                            _ = len(v1)
                        }
                    }
                    fmt.Println(fi46)
                    _ = fi48
                }
            }
            _ = fi46
        }
        if false {
            if ((v33 - v33) == v33) {
                v3 = v3
                if (true && true) {
                    for fi54 := 0; fi54 < 3; fi54++ {
                        v55 := GenS7{F0: "", F1: 0}
                        _ = fi54
                        _ = v55
                    }
                    v2 = v2
                }
            }
            if (!false) {
                if (v2 > 32768) {
                    fmt.Println(v28, v2, v2)
                    var v56 int = -9223372036854775808
                    _ = v56
                } else {
                    for fi57 := 0; fi57 < 9; fi57++ {
                        var v58 uint64 = 1
                        _ = fi57
                        _ = v58
                    }
                    fmt.Println(v38)
                    v1[(^(v28))] = 2
                }
            } else {
                var v59 int16 = -468
                fmt.Println(v38, v33, v33)
                v60 := GenS7{F0: (("foo bar" + "") + ""), F1: 1}
                _ = v59
                _ = v60
            }
            for fi61 := 0; fi61 < 9; fi61++ {
                for fi62 := 0; fi62 < 18; fi62++ {
                    var v63 int32 = -1
                    switch ("hello" + ((("" + "hello") + ("a" + ((("" + "foo bar") + "hello") + ("\n\t" + "")))) + "a")) {
                    case "hello":
                        v2 += (v2 >> ((65535) & 63))
                    case "":
                        v64 := GenS8{F0: 544}
                        _ = v64
                    }
                    fmt.Println(v63)
                    v28 -= (v28 | (v2 | (fi61 + (v2 << ((4294967295) & 63)))))
                    switch v2 {
                    case 879:
                        fmt.Println(fi62)
                    case -741:
                        v3 -= v3
                    default:
                        var v65 uint16 = 0
                        _ = v65
                    }
                    _ = fi62
                    _ = v63
                }
                for fi66 := 0; fi66 < 4; fi66++ {
                    var v67 uint = 42
                    v68 := GenS9{F0: 92, F1: 42}
                    v69 := GenS7{F0: "a", F1: 711}
                    _ = fi66
                    _ = v67
                    _ = v68
                    _ = v69
                }
                _ = fi61
            }
        } else {
            v38 -= v3
            for fi70 := 0; fi70 < 5; fi70++ {
                switch "a" {
                case "\n\t":
                    switch "" {
                    case "hello":
                        v38 = -2.25
                    case "foo bar":
                        v3 -= v38
                    case "a":
                        if false {
                            v71 := []int{42, v2}
                            v2 -= (-(v2))
                            _ = v71
                        }
                    default:
                        v72 := v1[2147483647]
                        _ = v72
                    }
                case "foo bar":
                    for fi73 := 0; fi73 < 5; fi73++ {
                        v74 := GenS2{F0: (v28 << ((0) & 63))}
                        var v75 uint16 = 255
                        _ = fi73
                        _ = v74
                        _ = v75
                    }
                case "hello":
                    v33 = (((v33 * float32((float64(v33) * v3))) + (v33 - (float32(v38) * ((v33 * (0.0 * 0.0)) / v33)))) * 1.5)
                }
                v3 = -2.25
                v76 := GenS5{F0: true}
                _ = fi70
                _ = v76
            }
            if ((!(v28 < 1)) || true) {
                v77 := GenS10{F0: 128, F1: (v2 / ((-2147483648) | 1)), F2: 32767}
                _ = v77
            }
        }
    } else {
        v1[(v28 % ((v28) | 1))] = 2147483648
        var v78 uint16 = 1
        fmt.Println(v3, v28)
        _ = v78
    }
    if false {
        var v79 int16 = 452
        v80 := []int{v28, (v28 & v28), (^(v28))}
        v38 += (v38 / v3)
        _ = v79
        _ = v80
    } else {
        delete(v1, v2)
        var v81 int8 = -128
        for fi82 := 0; fi82 < 14; fi82++ {
            var v83 uint32 = 42
            fmt.Println(v38, v2, v38)
            _ = fi82
            _ = v83
        }
        _ = v81
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v38
}
