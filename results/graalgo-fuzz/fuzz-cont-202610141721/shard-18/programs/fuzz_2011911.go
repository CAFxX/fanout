package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 string; F1 int8; F2 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS6 struct { F0 uint8; F1 int64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8 }
func genFn0() uint32 {
    v6 := GenS1{F0: 335, F1: 657}
    if (42 < 100) {
        v7 := GenS1{F0: 1, F1: -772}
        if true {
            if (false && ((4294967295 > 263) || (!(false && true)))) {
                if true {
                    for fi8 := 0; fi8 < 14; fi8++ {
                        if (!(759 < -1)) {
                            v9 := []int{(fi8 + 42), (fi8 % ((32767) | 1)), fi8}
                            _ = v9
                        }
                        if (!false) {
                            switch "hello" {
                            case "hello":
                                var v10 uint = 255
                                _ = v10
                            }
                            for fi11 := 0; fi11 < 17; fi11++ {
                                var v12 uint16 = 2
                                _ = fi11
                                _ = v12
                            }
                            switch 4294967295 {
                            case 127:
                                var v13 int = 1
                                _ = v13
                            default:
                                v14 := GenS1{F0: 79, F1: fi8}
                                _ = v14
                            }
                        } else {
                            switch "\n\t" {
                            case "hello":
                                v15 := []int{(fi8 - 32767), (fi8 * 4294967295), 4294967295, (fi8 + fi8)}
                                _ = v15
                            case "a":
                                var v16 int64 = 42
                                _ = v16
                            case "":
                                for fi17 := 0; fi17 < 9; fi17++ {
                                    var v18 bool = (!(-2147483648 >= 42))
                                    if (78 <= 1) {
                                        v19 := []int{fi17, fi8, fi17}
                                        v20 := GenS2{F0: 0, F1: false}
                                        _ = v19
                                        _ = v20
                                    } else {
                                        fmt.Println(fi17)
                                        var v21 int16 = -1
                                        v21 += (v21 * v21)
                                        var v22 uint8 = 128
                                        v23 := []int{fi17, (-(fi8)), (fi8 | 128)}
                                        _ = v21
                                        _ = v22
                                        _ = v23
                                    }
                                    _ = fi17
                                    _ = v18
                                }
                            default:
                                if true {
                                    v24 := GenS1{F0: 65535, F1: fi8}
                                    _ = v24
                                }
                            }
                            switch "hello" {
                            case "a":
                                var v25 uint32 = 2
                                _ = v25
                            case "foo bar":
                                v26 := GenS3{F0: "foo bar", F1: -1, F2: ("hello" + (("" + "hello") + "\n\t"))}
                                _ = v26
                            case "hello":
                                var v27 uint64 = 4294967295
                                _ = v27
                            default:
                                var v28 int32 = 1
                                _ = v28
                            }
                        }
                        _ = fi8
                    }
                    var v29 uint = 0
                    _ = v29
                }
            }
            var v30 uint8 = 127
            _ = v30
        }
        var v31 uint32 = 689
        _ = v31
        _ = v7
    } else {
        var v32 uint8 = 39
        v33 := []int{-9223372036854775808, 32767, 32767}
        var v34 uint16 = 32767
        _ = v32
        _ = v33
        _ = v34
    }
    _ = v6
    return 2147483647
}
func genFn1(p0 uint32, p1 string, p2 uint64) uint {
    if (!(!(!false))) {
        switch 2147483647 {
        case 503:
            switch "foo bar" {
            case "foo bar":
                fmt.Println(p0, p1, p0)
            case "\n\t":
                var v35 uint = 122
                _ = v35
            }
        case 42:
            var v36 uint64 = p2
            _ = v36
        case 1:
            for fi37 := 0; fi37 < 12; fi37++ {
                var v38 int8 = -128
                _ = fi37
                _ = v38
            }
        default:
            var v39 uint64 = 769
            _ = v39
        }
        if true {
            var v40 int16 = -718
            p0 += 737
            _ = v40
        }
    } else {
        var v41 bool = (true && true)
        _ = v41
    }
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return 791
}

func main() {
    var v1 bool = true
    var v2 uint8 = 1
    var v3 uint8 = v2
    var v4 string = "\n\t"
    var v5 float64 = (((((3.14159 / 3.14159) + -2.25) + 1e10) - 0.1) + float64(3.14159))
    v2 += v3
    v42 := GenS4{F0: 0}
    for fi43 := 0; fi43 < 17; fi43++ {
        v44 := GenS5{F0: 1, F1: 61, F2: 0}
        _ = fi43
        _ = v44
    }
    v45 := GenS2{F0: 1, F1: true}
    fmt.Println(v4)
    if v1 {
        v3 = (^(v2))
        for fi46 := 0; fi46 < 7; fi46++ {
            v47 := GenS1{F0: 32767, F1: fi46}
            v48 := []int{fi46, fi46}
            v49 := GenS6{F0: (v2 | v2), F1: int64(fi46)}
            _ = fi46
            _ = v47
            _ = v48
            _ = v49
        }
        v50 := GenS7{F0: 0}
        _ = v50
    } else {
        for fi51 := 0; fi51 < 2; fi51++ {
            v2 = v3
            _ = fi51
        }
    }
    if false {
        v52 := []int{0}
        _ = v52
    }
    fmt.Println(v2)
    for fi53 := 0; fi53 < 10; fi53++ {
        for fi54 := 0; fi54 < 16; fi54++ {
            v2 = v3
            if v1 {
                v2 += (v2 * (v2 - (v3 % ((v2) | 1))))
                var v55 uint64 = 621
                _ = v55
            }
            v56 := GenS6{F0: v3, F1: -1}
            var v57 uint = 18446744073709551615
            _ = fi54
            _ = v56
            _ = v57
        }
        v58 := []int{}
        switch (fi53 ^ (-(fi53))) {
        case 42:
            for fi59 := 0; fi59 < 14; fi59++ {
                switch (fi53 | (fi59 | fi53)) {
                case -9223372036854775808:
                    var v60 float64 = ((((v5 - v5) / v5) + v5) - float64(1e10))
                    _ = v60
                case 127:
                    _ = len(v58)
                case -910:
                    if (!((1 < 1000) && v1)) {
                        if false {
                            var v61 int8 = -46
                            _ = v61
                        }
                    } else {
                        if len(v58) >= 2 {
                            v58 = v58[1:len(v58)]
                        }
                        var v62 float32 = 0.1
                        if v1 {
                            switch (v4 + (v4 + v4)) {
                            case "a":
                                if len(v58) > 0 {
                                    v58[0] = fi59
                                }
                            case "\n\t":
                                if (255 == 65535) {
                                    v5 -= 0.0
                                    if false {
                                        var v63 string = "foo bar"
                                        _ = v63
                                    }
                                    v1 = (32767 <= 42)
                                }
                            default:
                                fmt.Println(v3, v2)
                            }
                        }
                        _ = v62
                    }
                default:
                    if true {
                        v5 -= v5
                        v1 = true
                        _ = len(v58)
                        var v64 int16 = 570
                        v4 = (v4 + ((v4 + v4) + v4))
                        v5 = (3.14159 - v5)
                        v2 -= v2
                        switch v4 {
                        case "a":
                            if len(v58) >= 2 {
                                v58 = v58[1:len(v58)]
                            }
                        default:
                            fmt.Println(fi59, v3, v64)
                        }
                        v5 = 3.14159
                        _ = v64
                    }
                }
                var v65 int = fi59
                v65 = (v65 + fi59)
                _ = fi59
                _ = v65
            }
        case 128:
            v66 := GenS2{F0: 1, F1: v1}
            _ = v66
        case 493:
            for fi67 := 0; fi67 < 5; fi67++ {
                var v68 bool = (42 >= 719)
                if (4294967295 != 1) {
                    v2 -= v3
                } else {
                    var v69 int8 = 127
                    var v70 string = "hello"
                    _ = v69
                    _ = v70
                }
                if v1 {
                    var v71 uint16 = 1
                    switch v4 {
                    case "\n\t":
                        v72 := GenS2{F0: -1, F1: v1}
                        _ = v72
                    case "a":
                        v73 := GenS4{F0: 33}
                        _ = v73
                    case "foo bar":
                        v1 = false
                    }
                    _ = v71
                } else {
                    if len(v58) > 0 {
                        v58[0] = fi67
                    }
                }
                _ = fi67
                _ = v68
            }
        }
        _ = fi53
        _ = v58
    }
    v74 := []int{1, 1, 48, 932}
    var v75 int8 = 100
    v76 := GenS8{F0: (v3 - 9)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v74:", len(v74), v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v45
    _ = v5
    _ = v74
    _ = v75
    _ = v76
}
