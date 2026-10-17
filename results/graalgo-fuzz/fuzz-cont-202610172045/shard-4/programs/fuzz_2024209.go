package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int16 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 int64; F1 int32; F2 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64; F1 uint8 }
type GenS8 struct { F0 uint16; F1 uint16; F2 int16 }
func genFn0(p0 int32) int8 {
    var v6 int64 = 42
    v7 := GenS1{F0: (v6 >> ((255) & 63)), F1: 2147483647, F2: 220}
    if (0 < -163) {
        if false {
            if (((p0 ^ -2147483648) != (-(p0))) || ((true || (1 != 42)) || true)) {
                var v8 uint8 = 208
                _ = v8
            } else {
                v9 := []int{9223372036854775807, 42}
                for fi10 := 0; fi10 < 15; fi10++ {
                    p0 -= 42
                    fmt.Println(fi10, p0)
                    _ = fi10
                }
                v11 := GenS2{F0: 199, F1: 100}
                _ = v11
                _ = v9
            }
            v12 := GenS1{F0: v6, F1: -299, F2: 1}
            var v13 float32 = (1e10 - ((float32(1.5) / 1.5) * 1e10))
            _ = v12
            _ = v13
        }
        var v14 float64 = -2.25
        _ = v14
    } else {
        var v15 string = "foo bar"
        v6 += v6
        _ = v15
    }
    v16 := []int{1, -1}
    if len(v16) > 0 {
        v16[0] = -9223372036854775808
    }
    _ = p0
    _ = v16
    _ = v6
    _ = v7
    return 1
}
func genFn1() uint16 {
    var v17 int64 = -9223372036854775808
    v18 := []int{42}
    if ((false || true) && true) {
        for fi19 := 0; fi19 < 12; fi19++ {
            for fi20 := 0; fi20 < 8; fi20++ {
                if ((!((!(float64(0.1) <= -2.25)) || false)) || (!((0.0 * float32(1e10)) <= (0.1 * 0.0)))) {
                    var v21 uint16 = 1
                    _ = v21
                }
                _ = fi20
            }
            v17 = v17
            v18 = append(v18, (-(fi19)))
            _ = fi19
        }
        if true {
            var v22 int16 = 32767
            _ = len(v18)
            if true {
                var v23 string = ("foo bar" + "foo bar")
                if (!((!(127 != -57)) && false)) {
                    v17 = v17
                    v22 -= (v22 | (v22 / ((v22) | 1)))
                    switch 32767 {
                    case -9223372036854775808:
                        if len(v18) > 0 {
                            v18[0] = 4294967295
                        }
                    }
                }
                _ = v23
            }
            _ = v22
        }
    }
    _ = v17
    _ = v18
    return 255
}

func main() {
    var v1 uint16 = 255
    var v2 int16 = 32767
    var v3 string = ("a" + ("\n\t" + ("" + ("" + "\n\t"))))
    var v4 string = v3
    var v5 float64 = 0.1
    if true {
        v5 += v5
    }
    switch "" {
    case "hello":
        fmt.Println(v2, v4, v1)
    default:
        switch "" {
        case "hello":
            var v24 string = "foo bar"
            _ = v24
        case "":
            var v25 uint8 = 128
            _ = v25
        default:
            var v26 uint8 = 1
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        switch 32767 {
        case 0:
            for fi28 := 0; fi28 < 13; fi28++ {
                fmt.Println(v3, v3, v2)
                fmt.Println(fi27, fi27)
                _ = fi28
            }
        case 32767:
            switch fi27 {
            case -999:
                v3 = "hello"
            default:
                v29 := []int{(fi27 << ((255) & 63)), (fi27 >> ((255) & 63)), fi27, fi27}
                _ = v29
            }
        case 42:
            v30 := GenS3{F0: 228}
            _ = v30
        }
        _ = fi27
    }
    switch ((("\n\t" + v4) + "\n\t") + v3) {
    case "\n\t":
        if (4294967295 != 42) {
            v31 := []int{883, 32767, 128, 202}
            if len(v31) >= 2 {
                v31 = v31[1:len(v31)]
            }
            var v32 int16 = 1000
            _ = v31
            _ = v32
        } else {
            if (1 >= 2147483648) {
                switch int(v1) {
                case 9223372036854775807:
                    var v33 int32 = int32(v1)
                    _ = v33
                case -9223372036854775808:
                    var v34 float32 = (0.1 * (0.0 / 1e10))
                    _ = v34
                }
            }
            var v35 float64 = -2.25
            v35 = float64(1.5)
            _ = v35
        }
    default:
        v2 = 329
    }
    v1 = (v1 - (v1 >> ((18446744073709551615) & 63)))
    for fi36 := 0; fi36 < 10; fi36++ {
        v4 = "hello"
        _ = fi36
    }
    var v37 int8 = 43
    if ("hello" <= v4) {
        for fi38 := 0; fi38 < 12; fi38++ {
            if (!(((v2 >> ((674) & 63)) >= (v2 | int16(v1))) || (1 != 217))) {
                fmt.Println(v37)
                v39 := []int{-794, (^(fi38)), 952, (^(fi38))}
                _ = v39
            } else {
                for fi40 := 0; fi40 < 13; fi40++ {
                    var v41 bool = false
                    var v42 int16 = -1
                    _ = fi40
                    _ = v41
                    _ = v42
                }
            }
            for fi43 := 0; fi43 < 16; fi43++ {
                var v44 float32 = (3.14159 * 3.14159)
                v44 = v44
                _ = fi43
                _ = v44
            }
            v45 := []int{fi38, (^(fi38)), fi38, (-(fi38))}
            _ = fi38
            _ = v45
        }
    } else {
        for fi46 := 0; fi46 < 4; fi46++ {
            var v47 string = (("hello" + (("\n\t" + ("" + "\n\t")) + "hello")) + "a")
            for fi48 := 0; fi48 < 12; fi48++ {
                v1 = genFn1()
                if true {
                    var v49 int16 = (v2 & (v2 ^ v2))
                    if true {
                        v50 := GenS2{F0: 127, F1: (v37 - 1)}
                        if false {
                            if false {
                                fmt.Println(v49)
                            } else {
                                fmt.Println(v37, v49, v3)
                            }
                        } else {
                            var v51 uint16 = v1
                            var v52 float32 = 1e10
                            v53 := GenS3{F0: -31}
                            _ = v51
                            _ = v52
                            _ = v53
                        }
                        fmt.Println(v5, v5, fi48)
                        _ = v50
                    } else {
                        var v54 string = v47
                        _ = v54
                    }
                    _ = v49
                }
                _ = fi48
            }
            v55 := GenS3{F0: -944}
            _ = fi46
            _ = v47
            _ = v55
        }
        v2 += (v2 << ((270) & 63))
    }
    v56 := []int{42, 9223372036854775807}
    var v57 bool = (!true)
    for fi58 := 0; fi58 < 7; fi58++ {
        fmt.Println(v2, fi58, v37)
        _ = fi58
    }
    if (!true) {
        if v57 {
            if len(v56) > 0 {
                v56[0] = -9223372036854775808
            }
            for fi59 := 0; fi59 < 8; fi59++ {
                var v60 bool = (3.14159 < 0.1)
                v3 = ("hello" + v4)
                _ = fi59
                _ = v60
            }
        }
    } else {
        v1 = 613
        fmt.Println(v1, v57)
        var v61 int = -9223372036854775808
        _ = v61
    }
    var v62 int = 567
    var v63 uint8 = 42
    v64 := GenS1{F0: 9223372036854775807, F1: v62, F2: (v2 + -250)}
    _ = len(v56)
    for fi65 := 0; fi65 < 18; fi65++ {
        if (!v57) {
            fmt.Println(v4, v4, v1)
            var v66 int64 = int64(v63)
            _ = v66
        } else {
            switch v4 {
            case "hello":
                switch (^(v62)) {
                case 32768:
                    for fi67 := 0; fi67 < 14; fi67++ {
                        v68 := GenS4{F0: 0, F1: 2147483647}
                        _ = fi67
                        _ = v68
                    }
                default:
                    if v57 {
                        var v69 float64 = ((v5 / (3.14159 / v5)) + float64(1.5))
                        _ = v69
                    }
                }
            case "\n\t":
                switch v4 {
                case "foo bar":
                    if len(v56) >= 2 {
                        v56 = v56[1:len(v56)]
                    }
                case "a":
                    switch v3 {
                    case "foo bar":
                        if (0 < 1) {
                            var v70 string = ((v4 + v4) + v4)
                            v71 := GenS2{F0: 2, F1: 1}
                            var v72 int = int(v2)
                            _ = v70
                            _ = v71
                            _ = v72
                        }
                    case "hello":
                        v56 = append(v56, (v62 / ((fi65) | 1)))
                    case "\n\t":
                        switch v3 {
                        case "":
                            fmt.Println(v4)
                        case "\n\t":
                            if v57 {
                                if (!v57) {
                                    v37 = (v37 & int8(v62))
                                }
                            } else {
                                v73 := GenS5{F0: -450, F1: -1, F2: fi65}
                                _ = v73
                            }
                        case "foo bar":
                            switch (fi65 * (fi65 - (fi65 & (fi65 >> ((849) & 63))))) {
                            case -418:
                                var v74 uint16 = (^(v1))
                                _ = v74
                            case 32767:
                                v4 = "\n\t"
                            case 53:
                                for fi75 := 0; fi75 < 9; fi75++ {
                                    var v76 uint = 0
                                    _ = fi75
                                    _ = v76
                                }
                            }
                        }
                    }
                }
            }
            var v77 float64 = (0.0 - v5)
            if false {
                var v78 int8 = genFn0(-2147483648)
                var v79 string = "a"
                _ = v78
                _ = v79
            } else {
                for fi80 := 0; fi80 < 4; fi80++ {
                    fmt.Println(v1)
                    if ((!(!v57)) || ((!v57) || (0 != 4294967295))) {
                        v81 := GenS2{F0: (v63 - 128), F1: v37}
                        v82 := GenS6{F0: (!(true && v57))}
                        if len(v56) >= 2 {
                            v56 = v56[1:len(v56)]
                        }
                        _ = v81
                        _ = v82
                    } else {
                        v83 := GenS7{F0: 57, F1: (v63 ^ 2)}
                        v4 = "hello"
                        v84 := GenS8{F0: (v1 | v1), F1: 854, F2: 0}
                        _ = v83
                        _ = v84
                    }
                    var v85 int = v62
                    _ = fi80
                    _ = v85
                }
                var v86 int64 = 383
                _ = v86
            }
            _ = v77
        }
        _ = fi65
    }
    var v87 float32 = (3.14159 * 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v56:", len(v56), v56)
    fmt.Println("v57:", v57)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v62
    _ = v63
    _ = v64
    _ = v87
}
