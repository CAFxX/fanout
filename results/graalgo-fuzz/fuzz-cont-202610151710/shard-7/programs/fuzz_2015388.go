package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint32 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint16; F1 int16; F2 bool }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 bool; F1 bool }
func genFn0() uint64 {
    var v5 uint32 = 65535
    for fi6 := 0; fi6 < 19; fi6++ {
        if (3.14159 < 0.0) {
            v7 := []int{489, fi6, (fi6 - fi6), fi6}
            var v8 uint64 = 141
            v5 = v5
            _ = v7
            _ = v8
        } else {
            var v9 int8 = -82
            if (3.14159 <= 1.5) {
                var v10 uint64 = 0
                for fi11 := 0; fi11 < 16; fi11++ {
                    v9 -= v9
                    v12 := GenS1{F0: 1, F1: "foo bar", F2: (v5 * (v5 << ((65535) & 63)))}
                    var v13 uint8 = 0
                    switch "\n\t" {
                    case "hello":
                        var v14 uint = 255
                        _ = v14
                    case "":
                        fmt.Println(fi6, fi11)
                    }
                    v15 := GenS2{F0: v9, F1: -128}
                    switch int(v5) {
                    case 0:
                        v16 := []int{(^(fi11))}
                        _ = v16
                    case -2147483648:
                        switch ("a" + ("" + "a")) {
                        case "hello":
                            switch ("" + ("foo bar" + "hello")) {
                            case "a":
                                v17 := GenS1{F0: 2, F1: "\n\t", F2: 4294967295}
                                _ = v17
                            case "foo bar":
                                v9 = -107
                            }
                        case "a":
                            var v18 uint32 = v5
                            _ = v18
                        }
                    case 2147483647:
                        v9 = (v9 ^ v9)
                        var v19 uint16 = 1
                        v13 += 2
                        _ = v19
                    }
                    if (false && ((v5 < (^(v5))) || false)) {
                        fmt.Println(v13)
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v15
                }
                _ = v10
            }
            _ = v9
        }
        _ = fi6
    }
    v5 -= (v5 - v5)
    var v20 int = -1
    for fi21 := 0; fi21 < 7; fi21++ {
        v22 := GenS2{F0: 42, F1: 56}
        var v23 float64 = 1.5
        _ = fi21
        _ = v22
        _ = v23
    }
    _ = v20
    _ = v5
    return 994
}
func genFn1(p0 int8) uint8 {
    for fi24 := 0; fi24 < 12; fi24++ {
        for fi25 := 0; fi25 < 14; fi25++ {
            v26 := []int{1, 127, (fi24 & (fi24 / ((127) | 1)))}
            fmt.Println(p0, fi24)
            var v27 uint = 42
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = fi24
    }
    p0 = -43
    v28 := []int{2147483647, 182}
    var v29 string = ""
    p0 += 1
    _ = p0
    _ = v28
    _ = v29
    return 6
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 1
    v3 := make(map[uint32]string)
    var v4 uint64 = 2
    fmt.Println(v4, v4)
    v4 = (v4 / ((v4) | 1))
    if true {
        var v30 int32 = -2147483648
        var v31 uint8 = uint8(v4)
        if false {
            if (!(!((!true) && (-2.25 == ((1.5 / ((-2.25 + 1e10) - (3.14159 / 1.0))) + 1.5))))) {
                for fi32 := 0; fi32 < 20; fi32++ {
                    fmt.Println(v31, v30)
                    switch fi32 {
                    case 127:
                        var v33 uint16 = 0
                        _ = v33
                    case -2147483648:
                        v2 += (v2 & v2)
                    case 32767:
                        fmt.Println(v31, v31, v31)
                    }
                    v34 := []int{-2147483648, fi32, fi32, (fi32 | fi32)}
                    _ = fi32
                    _ = v34
                }
                fmt.Println(v4, v4)
            }
        }
        _ = v30
        _ = v31
    } else {
        v2 = 0
        v1["\n\t"] = false
    }
    var v35 int = 32768
    if ((1.5 + -2.25) <= (1e10 * (((3.14159 / float32(-2.25)) * 0.0) + 0.0))) {
        v36 := GenS3{F0: 0, F1: (v2 * v2), F2: (2 <= 32767)}
        fmt.Println(v2, v2, v4)
        for fi37 := 0; fi37 < 15; fi37++ {
            var v38 float64 = (float64((((0.1 + ((-2.25 - (((0.0 / 3.14159) - -2.25) - 1e10)) - ((-2.25 / 0.1) + 1e10))) - 3.14159) - 0.0)) - 0.0)
            v3[4294967295] = "foo bar"
            _ = fi37
            _ = v38
        }
        _ = v36
    }
    switch -2147483648 {
    case -176:
        if false {
            v35 += 88
            if false {
                for fi39 := 0; fi39 < 1; fi39++ {
                    v40 := []int{v35, v35}
                    if false {
                        var v41 string = "foo bar"
                        switch v41 {
                        case "\n\t":
                            if ((v35 % (((^(fi39))) | 1)) == (v35 & (-(v35)))) {
                                if len(v40) >= 2 {
                                    v40 = v40[1:len(v40)]
                                }
                            }
                        case "a":
                            v35 = (fi39 >> ((175) & 63))
                        default:
                            if ((-2.25 / (1.5 * 1.5)) != ((1e10 + 3.14159) + 3.14159)) {
                                var v42 uint = 2
                                v43 := GenS1{F0: v4, F1: ((v41 + v41) + ("\n\t" + v41)), F2: 42}
                                v44 := GenS2{F0: -128, F1: 42}
                                _ = v42
                                _ = v43
                                _ = v44
                            }
                        }
                        _ = v41
                    } else {
                        fmt.Println(fi39)
                    }
                    _ = fi39
                    _ = v40
                }
            } else {
                switch "\n\t" {
                case "\n\t":
                    for fi45 := 0; fi45 < 8; fi45++ {
                        v46 := GenS3{F0: 65535, F1: (v2 * -32768), F2: false}
                        _ = fi45
                        _ = v46
                    }
                case "a":
                    if (0.0 >= (1e10 + (0.1 - 1.5))) {
                        var v47 int32 = 1
                        fmt.Println(v4)
                        v4 += v4
                        _ = v47
                    } else {
                        var v48 bool = true
                        switch "foo bar" {
                        case "foo bar":
                            fmt.Println(v2)
                        case "a":
                            fmt.Println(v2, v2)
                        case "\n\t":
                            fmt.Println(v48, v48, v4)
                        }
                        _ = v48
                    }
                }
                var v49 uint32 = 744
                if false {
                    v4 = (v4 + (^(v4)))
                    switch v35 {
                    case 32767:
                        v50 := GenS1{F0: genFn0(), F1: ("" + ("hello" + ("" + (("foo bar" + "foo bar") + "hello")))), F2: v49}
                        _ = v50
                    case -210:
                        fmt.Println(v35)
                    default:
                        v49 = (v49 >> ((1) & 63))
                    }
                    for fi51 := 0; fi51 < 9; fi51++ {
                        var v52 uint = 1
                        v52 -= (v52 % ((uint(v49)) | 1))
                        v53 := GenS4{F0: 844, F1: v49}
                        _ = fi51
                        _ = v52
                        _ = v53
                    }
                } else {
                    v54 := GenS5{F0: "a"}
                    _ = v54
                }
                _ = v49
            }
            v35 += (v35 * 190)
        } else {
            var v55 int8 = -120
            v56 := GenS5{F0: "foo bar"}
            fmt.Println(v55, v2)
            _ = v55
            _ = v56
        }
    case -2147483648:
        var v57 int = v35
        _ = v57
    }
    v58 := GenS1{F0: genFn0(), F1: "a", F2: 2}
    if (false || (true || false)) {
        if false {
            if (2147483648 == 2147483647) {
                v1[("" + "foo bar")] = false
                v35 += 32767
            } else {
                v59 := GenS3{F0: 1, F1: v2, F2: ((0 < 1) && true)}
                v60 := GenS6{F0: 661}
                v61 := GenS7{F0: 1}
                _ = v59
                _ = v60
                _ = v61
            }
        }
    } else {
        var v62 int64 = 9223372036854775807
        _ = v62
    }
    for fi63 := 0; fi63 < 10; fi63++ {
        v4 += genFn0()
        _ = fi63
    }
    if (v35 != 2147483647) {
        var v64 uint64 = v4
        _ = v64
    } else {
        for fi65 := 0; fi65 < 8; fi65++ {
            if true {
                v66 := []int{fi65}
                for fi67 := 0; fi67 < 19; fi67++ {
                    v68 := GenS8{F0: (!(false && ("foo bar" != (("a" + "hello") + "hello")))), F1: (!true)}
                    v35 = fi65
                    _ = fi67
                    _ = v68
                }
                _ = v66
            } else {
                switch ("foo bar" + "") {
                case "a":
                    var v69 uint16 = 0
                    _ = v69
                default:
                    fmt.Println(v2)
                }
            }
            fmt.Println(fi65)
            _ = fi65
        }
        for fi70 := 0; fi70 < 10; fi70++ {
            v71 := []int{v35, 0}
            v4 = 42
            v72 := v3[2147483647]
            _ = fi70
            _ = v71
            _ = v72
        }
        v73 := GenS1{F0: v4, F1: "foo bar", F2: 2147483648}
        _ = v73
    }
    var v74 float64 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v58
    _ = v74
}
