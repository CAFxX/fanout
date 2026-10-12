package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16; F1 bool; F2 uint16 }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 uint16; F1 int; F2 bool }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint8; F1 string }
type GenS10 struct { F0 int16; F1 uint8; F2 int }
type GenS11 struct { F0 uint64; F1 int8 }
func genFn0(p0 uint, p1 uint8, p2 uint8) int8 {
    if (514 == 2147483648) {
        var v4 float32 = 1e10
        var v5 uint64 = 18446744073709551615
        _ = v4
        _ = v5
    } else {
        var v6 uint16 = uint16(p2)
        v7 := GenS1{F0: false}
        v8 := GenS1{F0: (!false)}
        _ = v6
        _ = v7
        _ = v8
    }
    v9 := []int{1, 4294967295, 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 100
}
func genFn1(p0 int64) int64 {
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := GenS1{F0: true}
        var v12 uint16 = 2
        _ = fi10
        _ = v11
        _ = v12
    }
    p0 = (p0 - 865)
    v13 := GenS2{F0: 40, F1: 2147483648, F2: 65535}
    p0 += p0
    p0 += p0
    _ = p0
    _ = v13
    return -575
}

func main() {
    var v1 uint64 = 177
    var v2 string = "a"
    var v3 uint16 = 42
    if (!(!(42 >= 42))) {
        var v14 int8 = 0
        v15 := []int{32767, 32768, 2147483647}
        var v16 int16 = 32767
        _ = v14
        _ = v15
        _ = v16
    }
    v2 = ("hello" + v2)
    v17 := GenS3{F0: 0}
    v18 := []int{32767, 954}
    var v19 uint64 = v1
    switch "a" {
    case "a":
        var v20 int32 = 1
        _ = v20
    }
    if false {
        v19 += (v19 >> ((42) & 63))
        v18 = append(v18, 9223372036854775807)
        fmt.Println(v3, v19)
    } else {
        for fi21 := 0; fi21 < 11; fi21++ {
            fmt.Println(v3, fi21)
            for fi22 := 0; fi22 < 17; fi22++ {
                v19 = 18446744073709551615
                v19 = v1
                _ = fi22
            }
            v23 := GenS4{F0: 1000, F1: ((1e10 < (3.14159 + float64(((0.0 / float32(0.1)) - 0.1)))) && false), F2: (v3 >> ((2) & 63))}
            _ = fi21
            _ = v23
        }
        var v24 string = (v2 + "a")
        _ = v24
    }
    v18 = append(v18, 32767)
    if len(v18) > 0 {
        v18[0] = 1
    }
    if (0.1 >= 0.1) {
        for fi25 := 0; fi25 < 15; fi25++ {
            switch v2 {
            case "\n\t":
                switch v2 {
                case "":
                    for fi26 := 0; fi26 < 19; fi26++ {
                        v27 := GenS5{F0: 83, F1: 255}
                        _ = fi26
                        _ = v27
                    }
                case "\n\t":
                    v19 -= (v19 / ((1) | 1))
                case "a":
                    var v28 int64 = 42
                    _ = v28
                }
            }
            for fi29 := 0; fi29 < 20; fi29++ {
                var v30 uint8 = 0
                switch v2 {
                case "a":
                    switch v2 {
                    case "":
                        switch (v2 + v2) {
                        case "\n\t":
                            v31 := GenS1{F0: false}
                            _ = v31
                        case "foo bar":
                            switch fi29 {
                            case 816:
                                fmt.Println(v3, fi29)
                            case -2147483648:
                                v1 = v19
                            default:
                                _ = len(v18)
                            }
                        }
                    }
                    if true {
                        v3 = v3
                        fmt.Println(fi25)
                        fmt.Println(fi29)
                        if (v30 >= v30) {
                            v19 += v19
                            _ = len(v18)
                        }
                        if (false || true) {
                            fmt.Println(fi29, fi25)
                            v1 += (^(v19))
                            if (v1 <= uint64(v30)) {
                                if true {
                                    var v32 float32 = 1e10
                                    switch (v2 + ("" + ((v2 + v2) + "\n\t"))) {
                                    case "a":
                                        if false {
                                            if (genFn0(uint(v30), v30, v30) > 1) {
                                                var v33 string = "\n\t"
                                                _ = v33
                                            } else {
                                                v2 = ((v2 + v2) + v2)
                                                var v34 int = fi29
                                                _ = v34
                                            }
                                            if len(v18) > 0 {
                                                v18[0] = 0
                                            }
                                        }
                                        v35 := GenS6{F0: (v30 ^ (v30 | 128)), F1: (fi29 % ((fi25) | 1))}
                                        var v36 uint8 = v30
                                        _ = v35
                                        _ = v36
                                    }
                                    fmt.Println(v1)
                                    if len(v18) > 0 {
                                        v18[0] = -734
                                    }
                                    if (true && true) {
                                        v2 = "a"
                                    } else {
                                        if false {
                                            v2 = (v2 + (v2 + v2))
                                        }
                                    }
                                    switch ((v2 + ("hello" + "hello")) + v2) {
                                    case "\n\t":
                                        fmt.Println(v30, v3)
                                    case "hello":
                                        v30 = 255
                                    case "foo bar":
                                        if ((uint(v3) != 0) && false) {
                                            v37 := GenS7{F0: 699, F1: (fi25 & fi25), F2: false}
                                            if true {
                                                var v38 bool = true
                                                var v39 uint32 = 65535
                                                _ = v38
                                                _ = v39
                                            }
                                            _ = v37
                                        } else {
                                            v3 = (v3 / (((v3 / (((^(v3))) | 1))) | 1))
                                        }
                                    }
                                    _ = v32
                                }
                                switch 1 {
                                case 4294967295:
                                    switch 785 {
                                    case 127:
                                        switch "\n\t" {
                                        case "":
                                            fmt.Println(v19, v1)
                                            v2 = v2
                                            v30 = 127
                                        default:
                                            if (!true) {
                                                var v40 bool = false
                                                switch v2 {
                                                case "":
                                                    switch ("foo bar" + (v2 + "")) {
                                                    case "hello":
                                                        switch fi25 {
                                                        case 520:
                                                            var v41 int64 = 9223372036854775807
                                                            _ = v41
                                                        case 32768:
                                                            fmt.Println(v3, v2, v3)
                                                        case 690:
                                                            var v42 bool = true
                                                            _ = v42
                                                        default:
                                                            if (genFn1(42) > -293) {
                                                                v30 += v30
                                                                if (1 <= -179) {
                                                                    var v43 uint8 = uint8(v19)
                                                                    _ = v43
                                                                }
                                                            } else {
                                                                v44 := GenS4{F0: -72, F1: (((^(fi29)) >= fi29) || v40), F2: (v3 << ((0) & 63))}
                                                                var v45 uint8 = 1
                                                                _ = v44
                                                                _ = v45
                                                            }
                                                        }
                                                    case "":
                                                        if len(v18) >= 2 {
                                                            v18 = v18[1:len(v18)]
                                                        }
                                                    }
                                                default:
                                                    _ = len(v18)
                                                }
                                                _ = v40
                                            }
                                        }
                                    case -9223372036854775808:
                                        if true {
                                            v1 = (^(v1))
                                        } else {
                                            switch (v2 + v2) {
                                            case "hello":
                                                v2 = "a"
                                                if false {
                                                    var v46 uint8 = (v30 << ((255) & 63))
                                                    _ = v46
                                                }
                                            case "":
                                                var v47 uint16 = (v3 << ((18446744073709551615) & 63))
                                                _ = v47
                                            case "a":
                                                if false {
                                                    fmt.Println(v1)
                                                    v48 := GenS7{F0: v3, F1: (-(fi25)), F2: false}
                                                    fmt.Println(v2, v3, v19)
                                                    _ = v48
                                                }
                                            }
                                            switch 4294967295 {
                                            case 896:
                                                v2 = v2
                                            case -9223372036854775808:
                                                var v49 string = v2
                                                _ = v49
                                            default:
                                                if true {
                                                    var v50 uint = 2
                                                    var v51 uint8 = v30
                                                    var v52 float64 = -2.25
                                                    _ = v50
                                                    _ = v51
                                                    _ = v52
                                                }
                                            }
                                        }
                                    case 32768:
                                        if true {
                                            var v53 int = fi25
                                            _ = v53
                                        }
                                    }
                                default:
                                    fmt.Println(v30, fi25, v1)
                                }
                            }
                        }
                    } else {
                        if false {
                            if false {
                                var v54 string = ((((v2 + ("hello" + "a")) + (v2 + "a")) + "\n\t") + v2)
                                switch (fi29 % ((fi25) | 1)) {
                                case -1:
                                    v55 := GenS2{F0: (v3 >> ((892) & 63)), F1: 829, F2: 4294967295}
                                    _ = v55
                                case -534:
                                    var v56 bool = (false && ((v30 << ((2) & 63)) < (v30 | 10)))
                                    _ = v56
                                case 32767:
                                    if len(v18) > 0 {
                                        v18[0] = fi29
                                    }
                                }
                                v57 := GenS8{F0: v3}
                                _ = v54
                                _ = v57
                            }
                        }
                        if false {
                            v30 += uint8(v1)
                        }
                        if len(v18) > 0 {
                            v18[0] = fi29
                        }
                        var v58 string = (v2 + ((v2 + ((v2 + v2) + "a")) + ""))
                        _ = v58
                    }
                case "\n\t":
                    v18 = append(v18, 722)
                }
                v18 = append(v18, (fi25 - 149))
                v59 := GenS9{F0: 127, F1: (v2 + v2)}
                _ = fi29
                _ = v30
                _ = v59
            }
            v60 := GenS10{F0: 32767, F1: 0, F2: fi25}
            _ = fi25
            _ = v60
        }
        for fi61 := 0; fi61 < 14; fi61++ {
            v62 := GenS11{F0: (v19 * (^(v19))), F1: 42}
            var v63 uint32 = 386
            _ = fi61
            _ = v62
            _ = v63
        }
    } else {
        v64 := GenS8{F0: 682}
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
