package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 int64; F2 bool }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS6 struct { F0 uint8; F1 uint }
func genFn0(p0 bool, p1 int64) int8 {
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := GenS1{F0: (fi6 | fi6), F1: 765, F2: p1}
        var v8 uint8 = 127
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(p0)
    if false {
        p1 = (p1 / ((p1) | 1))
        if (0.1 <= 0.0) {
            p0 = (!(128 <= 32767))
            switch "foo bar" {
            case "a":
                var v9 uint16 = 0
                _ = v9
            case "foo bar":
                for fi10 := 0; fi10 < 12; fi10++ {
                    var v11 bool = (498 <= 229)
                    for fi12 := 0; fi12 < 16; fi12++ {
                        v13 := []int{fi12, 471}
                        _ = fi12
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 18; fi14++ {
                        if (!(!(float32(float64(1e10)) < (1e10 - (-2.25 - 0.1))))) {
                            v15 := []int{-9223372036854775808, 747, (fi14 >> ((4294967295) & 63)), -654}
                            _ = v15
                        }
                        p1 -= p1
                        var v16 uint32 = 4294967295
                        _ = fi14
                        _ = v16
                    }
                    _ = fi10
                    _ = v11
                }
            case "":
                v17 := GenS2{F0: 65535}
                _ = v17
            default:
                var v18 int8 = -1
                _ = v18
            }
            var v19 int16 = 1000
            _ = v19
        }
        var v20 uint64 = 18446744073709551615
        _ = v20
    } else {
        var v21 int64 = p1
        if (9223372036854775807 != 9223372036854775807) {
            switch 42 {
            case -9223372036854775808:
                for fi22 := 0; fi22 < 19; fi22++ {
                    var v23 int16 = 0
                    _ = fi22
                    _ = v23
                }
            default:
                fmt.Println(p0)
            }
            v24 := []int{128, 127, 128}
            v21 = (p1 / ((452) | 1))
            _ = v24
        }
        _ = v21
    }
    _ = p0
    _ = p1
    return 100
}
func genFn1(p0 int32, p1 bool, p2 uint8) uint32 {
    switch 467 {
    case 32767:
        v25 := []int{1}
        _ = v25
    case -2147483648:
        v26 := []int{128, -322, 1, 32767}
        _ = v26
    case 895:
        var v27 bool = (float64(0.1) == 0.0)
        _ = v27
    default:
        switch ("a" + "\n\t") {
        case "a":
            var v28 int16 = 32767
            _ = v28
        default:
            if p1 {
                p0 = p0
                var v29 string = "foo bar"
                p1 = p1
                _ = v29
            } else {
                switch 32767 {
                case 9223372036854775807:
                    for fi30 := 0; fi30 < 15; fi30++ {
                        var v31 int8 = 1
                        for fi32 := 0; fi32 < 15; fi32++ {
                            var v33 float64 = float64(float32(float64((0.1 + 1.5))))
                            switch (fi30 | (-(fi32))) {
                            case 409:
                                v34 := []int{(fi32 & fi30)}
                                _ = v34
                            case 32767:
                                p1 = (2 == 42)
                            case 1:
                                v35 := []int{(fi30 % (((fi32 ^ fi30)) | 1)), fi32, fi30}
                                _ = v35
                            default:
                                var v36 uint64 = 2
                                _ = v36
                            }
                            var v37 uint8 = 128
                            if p1 {
                                p2 = 255
                                if ((p2 * 0) > 2) {
                                    var v38 uint16 = 1
                                    if (false && (!false)) {
                                        var v39 uint16 = v38
                                        fmt.Println(v31)
                                        var v40 float32 = 1e10
                                        _ = v39
                                        _ = v40
                                    } else {
                                        v41 := []int{}
                                        _ = v41
                                    }
                                    v42 := []int{}
                                    _ = v38
                                    _ = v42
                                } else {
                                    fmt.Println(fi32)
                                    if ((-(v31)) >= v31) {
                                        switch ("a" + "") {
                                        case "a":
                                            v43 := []int{fi30, fi32}
                                            _ = v43
                                        case "foo bar":
                                            switch (fi32 & (fi30 >> ((65535) & 63))) {
                                            case 32767:
                                                if (p1 && false) {
                                                    p1 = (!(p1 && (v31 <= (v31 - 86))))
                                                    v44 := GenS3{F0: 1, F1: 9223372036854775807, F2: (944 > 1)}
                                                    _ = v44
                                                } else {
                                                    v45 := GenS4{F0: (fi32 + -176), F1: 918}
                                                    _ = v45
                                                }
                                            case 42:
                                                var v46 int32 = -2147483648
                                                _ = v46
                                            case 127:
                                                v47 := GenS1{F0: fi30, F1: 0, F2: 9223372036854775807}
                                                _ = v47
                                            default:
                                                var v48 bool = (255 >= 1)
                                                _ = v48
                                            }
                                        default:
                                            v49 := GenS5{F0: 434, F1: v37, F2: 2}
                                            _ = v49
                                        }
                                    }
                                }
                            }
                            _ = fi32
                            _ = v33
                            _ = v37
                        }
                        _ = fi30
                        _ = v31
                    }
                default:
                    v50 := []int{}
                    _ = v50
                }
                p0 -= (p0 | p0)
                if p1 {
                    if ("" <= (("foo bar" + ("\n\t" + "foo bar")) + (("foo bar" + "") + "\n\t"))) {
                        for fi51 := 0; fi51 < 9; fi51++ {
                            fmt.Println(fi51, p0, p1)
                            _ = fi51
                        }
                        p2 += p2
                    }
                    p0 -= (p0 % ((-1) | 1))
                }
            }
        }
    }
    switch -205 {
    case 0:
        if p1 {
            var v52 int64 = -152
            for fi53 := 0; fi53 < 6; fi53++ {
                var v54 int64 = (v52 ^ v52)
                var v55 float32 = (((1.5 / (((1e10 + (1.5 * 0.1)) * 1e10) * float32((3.14159 / -2.25)))) * (1e10 * 0.0)) * (3.14159 * -2.25))
                _ = fi53
                _ = v54
                _ = v55
            }
            _ = v52
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 int8 = -128
    var v2 int16 = 932
    var v3 int64 = 1
    v4 := make(map[int]int)
    v5 := []int{363, 4294967295, 1}
    var v56 string = ""
    fmt.Println(v3)
    for fi57 := 0; fi57 < 3; fi57++ {
        fmt.Println(v56, v1)
        v3 = v3
        var v58 int64 = -9223372036854775808
        _ = fi57
        _ = v58
    }
    for fi59 := 0; fi59 < 16; fi59++ {
        var v60 uint16 = 247
        for fi61 := 0; fi61 < 16; fi61++ {
            v3 = v3
            v4[(fi59 & fi59)] = 0
            fmt.Println(fi61, v56, fi61)
            v60 = (v60 << ((65535) & 63))
            v1 = (^(v1))
            _ = fi61
        }
        if true {
            v60 -= (v60 % (((v60 * v60)) | 1))
            if false {
                fmt.Println(v56, v60, v60)
                var v62 uint32 = 2147483647
                v63 := GenS3{F0: v3, F1: v3, F2: false}
                _ = v62
                _ = v63
            } else {
                var v64 uint16 = 2
                switch (v56 + ("hello" + ((v56 + (v56 + v56)) + ("a" + v56)))) {
                case "foo bar":
                    v5 = append(v5, (fi59 & 1))
                default:
                    for fi65 := 0; fi65 < 16; fi65++ {
                        switch "\n\t" {
                        case "\n\t":
                            var v66 uint8 = 167
                            _ = v66
                        case "hello":
                            v67 := GenS2{F0: 474}
                            _ = v67
                        default:
                            if (2 > 42) {
                                v1 = (-(v1))
                            } else {
                                fmt.Println(fi65, v1)
                                var v68 int32 = -747
                                _ = v68
                            }
                        }
                        var v69 string = (v56 + v56)
                        if false {
                            if true {
                                v70 := GenS6{F0: uint8(v64), F1: 42}
                                v71 := GenS2{F0: (^(v64))}
                                _ = v70
                                _ = v71
                            }
                        } else {
                            var v72 int8 = (v1 << ((42) & 63))
                            _ = v72
                        }
                        _ = fi65
                        _ = v69
                    }
                }
                v73 := GenS1{F0: fi59, F1: 940, F2: v3}
                _ = v64
                _ = v73
            }
            var v74 int = (fi59 + (fi59 * fi59))
            _ = v74
        } else {
            _ = len(v5)
            switch 0 {
            case 128:
                v75 := GenS1{F0: 489, F1: 495, F2: 325}
                _ = v75
            case -9223372036854775808:
                v2 -= v2
            case 4294967295:
                var v76 float32 = 1e10
                _ = v76
            default:
                switch (fi59 ^ (fi59 + fi59)) {
                case -327:
                    v56 = (v56 + "foo bar")
                case 376:
                    switch 1 {
                    case 2147483647:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case 32767:
                        for fi77 := 0; fi77 < 7; fi77++ {
                            var v78 float64 = 0.1
                            _ = fi77
                            _ = v78
                        }
                    }
                default:
                    for fi79 := 0; fi79 < 7; fi79++ {
                        fmt.Println(v2, v60)
                        v1 -= (v1 << ((42) & 63))
                        var v80 uint32 = genFn1(355, ((false && (86 > 2147483647)) && true), 204)
                        _ = fi79
                        _ = v80
                    }
                }
            }
        }
        _ = fi59
        _ = v60
    }
    switch -2147483648 {
    case 769:
        v3 += (v3 - v3)
    default:
        fmt.Println(v56)
    }
    fmt.Println(v3)
    var v81 uint8 = 42
    v56 = ((v56 + (v56 + (v56 + "a"))) + v56)
    delete(v4, -9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v81
}
