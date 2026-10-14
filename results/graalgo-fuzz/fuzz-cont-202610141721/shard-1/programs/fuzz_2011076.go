package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int; F1 string; F2 int8 }
type GenS3 struct { F0 int; F1 string; F2 int64 }
type GenS4 struct { F0 string; F1 int16; F2 int16 }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int8; F1 uint8 }
type GenS8 struct { F0 int16; F1 uint16 }
type GenS9 struct { F0 int8; F1 string }
type GenS10 struct { F0 int8; F1 int; F2 uint16 }
type GenS11 struct { F0 int64; F1 int; F2 uint16 }
func genFn0(p0 int, p1 int64, p2 uint8) int64 {
    var v5 int32 = int32(p2)
    v6 := []int{(p0 & (p0 * (p0 % (((p0 >> ((2) & 63))) | 1))))}
    var v7 string = "hello"
    v6 = append(v6, (-(p0)))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return -895
}
func genFn1(p0 int32, p1 int16) int8 {
    var v8 uint16 = 2
    switch "\n\t" {
    case "a":
        if (446 != 9223372036854775807) {
            if true {
                var v9 uint = 65535
                for fi10 := 0; fi10 < 20; fi10++ {
                    for fi11 := 0; fi11 < 1; fi11++ {
                        if true {
                            v12 := []int{4294967295}
                            if false {
                                switch "foo bar" {
                                case "foo bar":
                                    _ = len(v12)
                                case "a":
                                    var v13 bool = true
                                    _ = v13
                                }
                                var v14 uint32 = 160
                                _ = v14
                            }
                            if true {
                                v9 = 371
                                v8 = (v8 / ((v8) | 1))
                            }
                            _ = v12
                        }
                        _ = fi11
                    }
                    v15 := GenS1{F0: p1, F1: 1}
                    switch fi10 {
                    case 42:
                        if true {
                            if true {
                                var v16 string = ("foo bar" + "hello")
                                v17 := []int{}
                                switch (fi10 % (((fi10 << (((v9 / (((v9 << ((v9) & 63))) | 1))) & 63))) | 1)) {
                                case 32768:
                                    var v18 int = -2147483648
                                    _ = v18
                                case 127:
                                    if (float64(0.1) <= ((((1e10 * (1e10 * 1.5)) * 0.1) - (1e10 * 3.14159)) - (0.1 - (1.5 - 1.5)))) {
                                        for fi19 := 0; fi19 < 10; fi19++ {
                                            switch ((((("a" + v16) + v16) + v16) + (v16 + "")) + (v16 + (v16 + v16))) {
                                            case "\n\t":
                                                v9 += (^(v9))
                                            case "hello":
                                                v16 = "foo bar"
                                                var v20 uint16 = v8
                                                p1 = int16(fi19)
                                                p0 -= p0
                                                _ = v20
                                            }
                                            _ = fi19
                                        }
                                        var v21 uint16 = (v8 & (^(v8)))
                                        if false {
                                            v9 = 728
                                            switch fi10 {
                                            case -845:
                                                if len(v17) >= 2 {
                                                    v17 = v17[1:len(v17)]
                                                }
                                            default:
                                                var v22 int64 = 1
                                                _ = v22
                                            }
                                            var v23 int32 = (p0 - (p0 | 607))
                                            _ = v23
                                        } else {
                                            for fi24 := 0; fi24 < 8; fi24++ {
                                                v8 = 65535
                                                var v25 bool = (-2.25 < 0.1)
                                                _ = fi24
                                                _ = v25
                                            }
                                        }
                                        _ = v21
                                    }
                                }
                                _ = v16
                                _ = v17
                            }
                            if (1000 < (p1 >> ((v9) & 63))) {
                                fmt.Println(v9, v9)
                            }
                            var v26 uint32 = 2
                            _ = v26
                        } else {
                            v8 = 225
                            p0 = -1
                        }
                    case 128:
                        fmt.Println(p0, p0, v8)
                    case 32767:
                        var v27 float64 = 1.5
                        _ = v27
                    default:
                        if (!true) {
                            for fi28 := 0; fi28 < 12; fi28++ {
                                v8 -= 2
                                _ = fi28
                            }
                            for fi29 := 0; fi29 < 10; fi29++ {
                                v30 := GenS2{F0: 128, F1: "", F2: 42}
                                switch "hello" {
                                case "\n\t":
                                    if false {
                                        if (4294967295 >= 4294967295) {
                                            if true {
                                                var v31 int = -687
                                                _ = v31
                                            }
                                            fmt.Println(fi10, v8, fi10)
                                            if true {
                                                v32 := GenS3{F0: 1, F1: "a", F2: int64(v8)}
                                                fmt.Println(v8, p1, v8)
                                                _ = v32
                                            }
                                        } else {
                                            var v33 int64 = -1
                                            _ = v33
                                        }
                                        var v34 float64 = 0.0
                                        _ = v34
                                    } else {
                                        v35 := []int{(-(fi29)), fi29, -9223372036854775808}
                                        v36 := GenS4{F0: "", F1: 1, F2: 1}
                                        _ = v35
                                        _ = v36
                                    }
                                default:
                                    v37 := GenS2{F0: int(p0), F1: ("foo bar" + "hello"), F2: 42}
                                    _ = v37
                                }
                                fmt.Println(fi10, v8)
                                _ = fi29
                                _ = v30
                            }
                            var v38 int8 = 127
                            _ = v38
                        } else {
                            v8 = (v8 % ((v8) | 1))
                        }
                    }
                    _ = fi10
                    _ = v15
                }
                fmt.Println(p1)
                _ = v9
            }
            v39 := GenS5{F0: (p1 | (p1 >> ((255) & 63))), F1: v8}
            v40 := []int{-2147483648, 199}
            _ = v39
            _ = v40
        } else {
            if true {
                var v41 int16 = p1
                var v42 int16 = v41
                _ = v41
                _ = v42
            } else {
                var v43 int32 = (-(p0))
                if ((p1 & 1000) != 1000) {
                    for fi44 := 0; fi44 < 15; fi44++ {
                        var v45 float64 = ((-2.25 - 3.14159) * (-2.25 + (0.1 / 0.1)))
                        v46 := GenS6{F0: p1}
                        _ = fi44
                        _ = v45
                        _ = v46
                    }
                } else {
                    for fi47 := 0; fi47 < 5; fi47++ {
                        var v48 uint16 = 32767
                        var v49 string = "a"
                        _ = fi47
                        _ = v48
                        _ = v49
                    }
                    if (v8 == v8) {
                        var v50 uint64 = 42
                        var v51 string = ("" + "\n\t")
                        _ = v50
                        _ = v51
                    }
                }
                _ = v43
            }
            p1 -= (p1 - int16(p0))
        }
    case "":
        fmt.Println(v8, v8)
    }
    switch 128 {
    case -540:
        v52 := []int{}
        _ = v52
    default:
        v8 += (v8 + (v8 | (^(v8))))
    }
    v8 = (v8 - 255)
    _ = p0
    _ = p1
    _ = v8
    return -1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    var v3 bool = false
    v4 := make(map[string]int)
    v3 = true
    fmt.Println(v3)
    v53 := GenS7{F0: 1, F1: 255}
    v54 := GenS8{F0: 1, F1: 32767}
    v55 := v1[65535]
    for fi56 := 0; fi56 < 9; fi56++ {
        v57 := []int{fi56}
        _ = fi56
        _ = v57
    }
    if ((1.5 * 0.1) >= 0.1) {
        for fi58 := 0; fi58 < 15; fi58++ {
            var v59 int16 = 42
            v60 := v1[324]
            v55 -= 2147483647
            _ = fi58
            _ = v59
            _ = v60
        }
        fmt.Println(v3, v3)
        var v61 uint32 = 1
        _ = v61
    }
    v3 = v3
    for fi62 := 0; fi62 < 3; fi62++ {
        v63 := GenS9{F0: genFn1(0, -343), F1: "hello"}
        v64 := []int{32768, (fi62 & -1)}
        _ = fi62
        _ = v63
        _ = v64
    }
    for fi65 := 0; fi65 < 2; fi65++ {
        v66 := []int{}
        var v67 string = "hello"
        v68 := GenS10{F0: -118, F1: 1, F2: 1}
        _ = fi65
        _ = v66
        _ = v67
        _ = v68
    }
    var v69 uint16 = 255
    var v70 string = "foo bar"
    for fi71 := 0; fi71 < 5; fi71++ {
        var v72 uint = 255
        v2[2147483648] = 32767
        _ = fi71
        _ = v72
    }
    switch "foo bar" {
    case "foo bar":
        delete(v2, 23)
    case "hello":
        for fi73 := 0; fi73 < 10; fi73++ {
            if v3 {
                delete(v2, 956)
                for fi74 := 0; fi74 < 16; fi74++ {
                    var v75 uint8 = 82
                    if true {
                        fmt.Println(v69, fi73, v70)
                        var v76 int8 = -1
                        _ = v76
                    } else {
                        var v77 float64 = 1.5
                        fmt.Println(v75, v55, v69)
                        if false {
                            v78 := GenS4{F0: v70, F1: 32767, F2: 728}
                            if v3 {
                                fmt.Println(v75, fi73, v3)
                            }
                            switch (v70 + ("hello" + v70)) {
                            case "\n\t":
                                var v79 uint32 = uint32(v75)
                                _ = v79
                            case "":
                                var v80 bool = v3
                                _ = v80
                            default:
                                fmt.Println(fi74)
                            }
                            _ = v78
                        }
                        _ = v77
                    }
                    _ = fi74
                    _ = v75
                }
                v55 = -1
            }
            _ = fi73
        }
    case "a":
        for fi81 := 0; fi81 < 16; fi81++ {
            if (0 > 51) {
                v3 = (5 < genFn1(621, -969))
                v69 += v69
                v3 = v3
            }
            v3 = (v55 > v55)
            _ = fi81
        }
    }
    _ = len(v4)
    fmt.Println(v3, v55)
    if ((1 == v69) || v3) {
        var v82 int32 = 2147483647
        _ = v82
    } else {
        switch 4294967295 {
        case 513:
            v83 := GenS11{F0: 9223372036854775807, F1: v55, F2: v69}
            _ = v83
        case 0:
            v84 := v1[65535]
            _ = v84
        case 325:
            v55 -= v55
        default:
            switch v70 {
            case "hello":
                for fi85 := 0; fi85 < 11; fi85++ {
                    fmt.Println(v3, v3)
                    _ = fi85
                }
            case "foo bar":
                v70 = (v70 + "foo bar")
            case "a":
                switch v55 {
                case 404:
                    v69 += v69
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v55:", v55)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v54
    _ = v55
    _ = v69
    _ = v70
}
