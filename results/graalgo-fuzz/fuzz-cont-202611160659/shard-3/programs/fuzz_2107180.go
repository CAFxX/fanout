package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int32 }
type GenS2 struct { F0 int64; F1 bool; F2 int16 }
type GenS3 struct { F0 int64; F1 uint16; F2 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 uint32; F1 int16 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int16; F1 uint }
type GenS10 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS11 struct { F0 uint8; F1 int64; F2 uint64 }
func genFn0() int32 {
    for fi5 := 0; fi5 < 12; fi5++ {
        for fi6 := 0; fi6 < 2; fi6++ {
            v7 := GenS1{F0: 1, F1: 42, F2: 42}
            var v8 int8 = -1
            if (0 <= 33) {
                if false {
                    fmt.Println(fi6, fi6, fi5)
                }
                switch "" {
                case "hello":
                    v9 := []int{fi5}
                    _ = v9
                case "foo bar":
                    v8 -= -1
                case "\n\t":
                    var v10 uint64 = 1
                    v11 := []int{}
                    v11 = append(v11, fi5)
                    _ = v10
                    _ = v11
                default:
                    var v12 int16 = 631
                    _ = v12
                }
                var v13 uint32 = 1
                _ = v13
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        var v14 uint32 = 2
        _ = fi5
        _ = v14
    }
    v15 := GenS2{F0: 9223372036854775807, F1: true, F2: 1}
    var v16 string = "hello"
    v17 := []int{4294967295}
    _ = v15
    _ = v16
    _ = v17
    return 270
}
func genFn1() string {
    var v18 int16 = -539
    var v19 int32 = 1
    v18 -= v18
    v19 += genFn0()
    _ = v18
    _ = v19
    return ("foo bar" + "hello")
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[int]uint32)
    var v3 int16 = 42
    v4 := make(map[string]string)
    v3 = (v3 + (v3 / ((32767) | 1)))
    var v20 int16 = v3
    var v21 float64 = 0.1
    for fi22 := 0; fi22 < 10; fi22++ {
        switch ("" + "a") {
        case "hello":
            switch fi22 {
            case 32767:
                v23 := GenS3{F0: 179, F1: 32767, F2: (2147483647 == (fi22 ^ fi22))}
                _ = v23
            case -225:
                if (v21 > (v21 + 3.14159)) {
                    fmt.Println(v1)
                    v4["hello"] = "a"
                    switch "foo bar" {
                    case "hello":
                        if false {
                            if ((!false) || false) {
                                for fi24 := 0; fi24 < 10; fi24++ {
                                    v1 = float32(3.14159)
                                    _ = fi24
                                }
                                v20 = (v3 << ((619) & 63))
                            }
                            var v25 int32 = 330
                            _ = v25
                        }
                    case "foo bar":
                        v21 = float64(v1)
                    case "":
                        switch (fi22 << ((2) & 63)) {
                        case -9223372036854775808:
                            v26 := GenS4{F0: 92}
                            _ = v26
                        default:
                            switch "\n\t" {
                            case "foo bar":
                                for fi27 := 0; fi27 < 7; fi27++ {
                                    var v28 int32 = 1
                                    var v29 uint32 = 249
                                    _ = fi27
                                    _ = v28
                                    _ = v29
                                }
                            default:
                                if ((65535 > 0) || (158 < 1)) {
                                    _ = len(v4)
                                    v30 := []int{}
                                    _ = v30
                                }
                            }
                        }
                    default:
                        if true {
                            _ = len(v4)
                        } else {
                            var v31 bool = (((v20 << ((42) & 63)) >= (v3 / ((v3) | 1))) || true)
                            var v32 int64 = -1
                            v33 := v2[(^(fi22))]
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                    }
                } else {
                    v34 := GenS4{F0: 180}
                    _ = v34
                }
            case 9223372036854775807:
                if ((!false) && (!((v21 * -2.25) == 0.0))) {
                    switch "\n\t" {
                    case "foo bar":
                        v21 = ((v21 + v21) - 0.0)
                    default:
                        v35 := GenS5{F0: -1}
                        _ = v35
                    }
                } else {
                    switch (-(fi22)) {
                    case 127:
                        fmt.Println(v1, v1)
                    }
                    v36 := []int{(fi22 ^ fi22), fi22}
                    _ = v36
                }
            default:
                for fi37 := 0; fi37 < 9; fi37++ {
                    v38 := []int{int(v20), fi37}
                    var v39 int64 = 686
                    var v40 string = ("" + "hello")
                    if (2 == 414) {
                        _ = len(v4)
                        delete(v2, fi37)
                        if (v21 < (v21 / v21)) {
                            fmt.Println(v39, fi22)
                            if false {
                                if false {
                                    v40 = ("" + "\n\t")
                                }
                                var v41 uint32 = 42
                                _ = v41
                            } else {
                                if true {
                                    var v42 uint64 = 42
                                    if false {
                                        var v43 string = ("a" + "")
                                        var v44 bool = (42 > (v42 % ((42) | 1)))
                                        v44 = (((v43 + "foo bar") + v40) == v40)
                                        v45 := GenS3{F0: v39, F1: 0, F2: (64 == 98)}
                                        v42 += (v42 % (((v42 ^ (v42 / (((v42 ^ (v42 >> ((1) & 63)))) | 1)))) | 1))
                                        _ = v43
                                        _ = v44
                                        _ = v45
                                    } else {
                                        v21 = v21
                                    }
                                    _ = v42
                                } else {
                                    v40 = ("hello" + (v40 + v40))
                                    if false {
                                        _ = len(v38)
                                        v46 := GenS6{F0: v3, F1: v40}
                                        v4[v40] = ((((("foo bar" + ("a" + v40)) + (v40 + v40)) + v40) + "\n\t") + (("a" + v40) + "hello"))
                                        _ = v46
                                    } else {
                                        var v47 float32 = (v1 - v1)
                                        _ = v47
                                    }
                                    if (false && (true || false)) {
                                        switch "hello" {
                                        case "a":
                                            var v48 int16 = v3
                                            _ = v48
                                        case "foo bar":
                                            v2[fi22] = 4294967295
                                        default:
                                            v39 += (v39 * 9223372036854775807)
                                        }
                                        if (!true) {
                                            v40 = v40
                                            delete(v4, "hello")
                                            var v49 uint16 = 654
                                            _ = v49
                                        }
                                        var v50 int64 = (v39 << ((65535) & 63))
                                        _ = v50
                                    }
                                }
                                switch fi22 {
                                case 32768:
                                    switch v40 {
                                    case "a":
                                        v38 = append(v38, fi37)
                                    case "foo bar":
                                        if false {
                                            if false {
                                                fmt.Println(v21, v39)
                                                _ = len(v38)
                                                switch ("foo bar" + v40) {
                                                case "hello":
                                                    var v51 bool = (!(705 < 1))
                                                    _ = v51
                                                case "a":
                                                    v52 := GenS4{F0: 0}
                                                    _ = v52
                                                default:
                                                    var v53 uint = 42
                                                    if ((1 >= 191) || true) {
                                                        if ((!true) && (1.5 >= v1)) {
                                                            var v54 int64 = (^(v39))
                                                            var v55 int64 = 1
                                                            if (0 != 108) {
                                                                fmt.Println(v39, v53)
                                                                fmt.Println(v53, v40, v21)
                                                            }
                                                            _ = v54
                                                            _ = v55
                                                        } else {
                                                            v56 := GenS7{F0: 42, F1: v20}
                                                            _ = v56
                                                        }
                                                        v20 = v3
                                                    }
                                                    _ = v53
                                                }
                                                v20 -= (^(v20))
                                            }
                                        } else {
                                            delete(v2, (fi37 | fi37))
                                        }
                                    }
                                    v57 := v4[v40]
                                    _ = v57
                                }
                                v58 := GenS8{F0: 2147483648}
                                _ = len(v38)
                                _ = v58
                            }
                            v4[("foo bar" + "hello")] = "hello"
                        }
                    }
                    _ = fi37
                    _ = v38
                    _ = v39
                    _ = v40
                }
            }
        case "foo bar":
            for fi59 := 0; fi59 < 18; fi59++ {
                var v60 uint32 = 42
                v61 := []int{}
                var v62 uint32 = 78
                var v63 uint8 = 2
                fmt.Println(v1, v3, v62)
                _ = fi59
                _ = v60
                _ = v61
                _ = v62
                _ = v63
            }
        }
        _ = fi22
    }
    v64 := GenS3{F0: -1, F1: 2, F2: ((401 >= 2147483648) && (-9223372036854775808 >= 0))}
    if (!true) {
        for fi65 := 0; fi65 < 7; fi65++ {
            v3 -= (v20 % ((-631) | 1))
            v66 := []int{}
            _ = fi65
            _ = v66
        }
        v21 = (v21 * (v21 / 1.0))
    } else {
        v67 := GenS3{F0: 9223372036854775807, F1: 65535, F2: ((false && true) && (0 > 2))}
        for fi68 := 0; fi68 < 4; fi68++ {
            var v69 int32 = 1
            v20 = (v20 % ((v20) | 1))
            _ = fi68
            _ = v69
        }
        _ = v67
    }
    v1 -= (v1 / 1.5)
    v21 = (-2.25 * v21)
    switch -9223372036854775808 {
    case 128:
        var v70 int64 = 9223372036854775807
        _ = v70
    case -247:
        fmt.Println(v20, v21, v3)
    case 32767:
        if (0 <= 2) {
            v20 = v3
            switch -9223372036854775808 {
            case 4294967295:
                switch 9223372036854775807 {
                case 1:
                    fmt.Println(v3, v3)
                case -335:
                    switch 4294967295 {
                    case -961:
                        delete(v2, 899)
                    }
                default:
                    if true {
                        for fi71 := 0; fi71 < 8; fi71++ {
                            if false {
                                v72 := []int{32767, (^(fi71)), fi71}
                                _ = v72
                            } else {
                                switch 9223372036854775807 {
                                case 201:
                                    fmt.Println(v21)
                                case -272:
                                    var v73 int64 = 1
                                    _ = v73
                                case -9223372036854775808:
                                    switch fi71 {
                                    case 1:
                                        v21 += v21
                                    case 375:
                                        switch ("hello" + "hello") {
                                        case "a":
                                            v21 += 0.1
                                        case "\n\t":
                                            v74 := GenS9{F0: v3, F1: 2}
                                            _ = v74
                                        case "foo bar":
                                            var v75 int64 = -492
                                            _ = v75
                                        }
                                    case -992:
                                        var v76 float64 = ((v21 + ((v21 - float64((3.14159 - v1))) + v21)) / v21)
                                        _ = v76
                                    }
                                }
                                v21 -= 1.5
                            }
                            var v77 int8 = 100
                            _ = fi71
                            _ = v77
                        }
                        if true {
                            var v78 string = ""
                            _ = v78
                        } else {
                            var v79 uint32 = 50
                            v3 = (v3 << ((42) & 63))
                            v80 := GenS5{F0: 42}
                            _ = v79
                            _ = v80
                        }
                    }
                }
            }
        }
    default:
        switch "\n\t" {
        case "hello":
            v81 := GenS10{F0: -128, F1: 2147483648, F2: 893}
            _ = v81
        default:
            v3 = v20
        }
    }
    var v82 float32 = ((((v1 / v1) / v1) - v1) - v1)
    v21 = 1.5
    for fi83 := 0; fi83 < 16; fi83++ {
        if true {
            fmt.Println(v20)
        } else {
            v84 := GenS11{F0: 117, F1: -872, F2: 4294967295}
            _ = v84
        }
        _ = fi83
    }
    v3 -= (v20 ^ v20)
    var v85 uint64 = 620
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v82:", v82)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v64
    _ = v82
    _ = v85
}
