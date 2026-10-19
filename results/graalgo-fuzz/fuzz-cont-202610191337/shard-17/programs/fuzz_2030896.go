package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int16; F1 int64; F2 int }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 uint16; F1 int; F2 uint16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int16; F1 int8; F2 uint8 }
func genFn0(p0 uint64, p1 int16, p2 uint8) uint32 {
    fmt.Println(p1)
    v5 := GenS1{F0: p0, F1: 397}
    p2 -= 191
    switch -339 {
    case 32767:
        v6 := GenS2{F0: -509, F1: 1, F2: 0}
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 4294967295
}
func genFn1() string {
    var v7 float64 = 3.14159
    var v8 int32 = 460
    switch -158 {
    case 830:
        switch 9223372036854775807 {
        case -1:
            if ((!(true && false)) && false) {
                if (-2147483648 != (v8 - 1)) {
                    v8 += v8
                    fmt.Println(v7)
                    v8 = (v8 | v8)
                }
                fmt.Println(v7, v7)
                var v9 string = "hello"
                _ = v9
            }
        case 42:
            var v10 float32 = 1e10
            _ = v10
        case -32:
            v11 := GenS3{F0: 0, F1: 0, F2: -9223372036854775808}
            _ = v11
        default:
            v8 -= -1
        }
    case 127:
        v7 = 0.1
    case -9223372036854775808:
        v7 -= v7
    }
    var v12 int16 = 0
    _ = v12
    _ = v7
    _ = v8
    return "foo bar"
}

func main() {
    var v1 int32 = -1
    v2 := make(map[string]string)
    v3 := []int{-2147483648, -9223372036854775808, -507, 127}
    var v4 string = (("" + "") + ("a" + ("\n\t" + "")))
    delete(v2, v4)
    var v13 int32 = v1
    var v14 uint16 = 0
    if (!(false && false)) {
        if (!false) {
            if true {
                if (100 > 42) {
                    if false {
                        for fi15 := 0; fi15 < 16; fi15++ {
                            v13 = v1
                            _ = fi15
                        }
                        var v16 uint = 18446744073709551615
                        _ = v16
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            v17 := GenS3{F0: 0, F1: 42, F2: 916}
                            _ = v17
                        default:
                            v13 -= (v1 % (((^(v13))) | 1))
                        }
                    }
                    v1 -= (v13 - (v13 & (v1 << ((1) & 63))))
                } else {
                    switch ((v4 + v4) + "\n\t") {
                    case "\n\t":
                        if (497 <= 4294967295) {
                            if true {
                                v18 := GenS4{F0: (3.14159 >= -2.25), F1: -9223372036854775808}
                                var v19 int32 = 2147483647
                                _ = v18
                                _ = v19
                            }
                            delete(v2, (("\n\t" + ((("hello" + v4) + v4) + v4)) + v4))
                        }
                    case "a":
                        var v20 float32 = float32(0.0)
                        _ = v20
                    }
                    switch -411 {
                    case 42:
                        _ = len(v2)
                    case 1:
                        var v21 int64 = -9223372036854775808
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 16; fi22++ {
                        fmt.Println(fi22, v1)
                        _ = fi22
                    }
                }
                if true {
                    var v23 int8 = -128
                    fmt.Println(v23, v14)
                    _ = v23
                } else {
                    v2[(v4 + v4)] = "a"
                    fmt.Println(v4, v14)
                    for fi24 := 0; fi24 < 6; fi24++ {
                        for fi25 := 0; fi25 < 15; fi25++ {
                            if len(v3) > 0 {
                                v3[0] = (fi24 << ((124) & 63))
                            }
                            _ = fi25
                        }
                        _ = fi24
                    }
                }
                switch "" {
                case "\n\t":
                    if len(v3) > 0 {
                        v3[0] = -9223372036854775808
                    }
                case "":
                    for fi26 := 0; fi26 < 3; fi26++ {
                        v27 := GenS5{F0: 1, F1: -1}
                        v13 = 215
                        v4 = v4
                        _ = fi26
                        _ = v27
                    }
                case "hello":
                    v4 = v4
                default:
                    switch 42 {
                    case 606:
                        v1 += -975
                    }
                }
            } else {
                var v28 uint16 = v14
                switch v4 {
                case "":
                    v13 = -961
                }
                _ = v28
            }
            v14 = v14
        } else {
            v1 += 42
        }
    }
    if (713 <= 42) {
        var v29 float64 = ((1.5 - (((float64(((((0.1 - -2.25) + 3.14159) + 1.5) / -2.25)) - (0.1 - 0.0)) * (((0.1 * 1.5) / float64((0.1 + 1e10))) * 1e10)) * 1e10)) - 0.0)
        _ = v29
    }
    if len(v3) > 0 {
        v3[0] = -421
    }
    var v30 int16 = 42
    for fi31 := 0; fi31 < 1; fi31++ {
        fmt.Println(v13, v30)
        _ = len(v3)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi31
    }
    fmt.Println(v14, v13)
    var v32 float64 = 1e10
    v30 = v30
    v33 := GenS6{F0: v14, F1: 556, F2: v14}
    v32 = v32
    if ((!false) && true) {
        fmt.Println(v30)
    } else {
        delete(v2, v4)
        var v34 string = ("hello" + ("hello" + v4))
        if (689 >= 1) {
            v35 := GenS5{F0: 147, F1: 95}
            for fi36 := 0; fi36 < 2; fi36++ {
                v37 := GenS3{F0: v30, F1: -332, F2: fi36}
                _ = fi36
                _ = v37
            }
            v38 := v2[v34]
            _ = v35
            _ = v38
        } else {
            v39 := GenS4{F0: (-304 == -32768), F1: 0}
            switch 32768 {
            case -91:
                if (true || ((^(v14)) < v14)) {
                    var v40 uint32 = 2147483647
                    switch -2147483648 {
                    case -426:
                        var v41 uint16 = v14
                        _ = v41
                    case -654:
                        if true {
                            fmt.Println(v34, v4, v14)
                            fmt.Println(v4, v40)
                            v4 = (v4 + ("\n\t" + ""))
                        } else {
                            v32 -= ((1.5 - (v32 * v32)) / float64(0.1))
                            switch 4294967295 {
                            case 1:
                                switch 1 {
                                case -2147483648:
                                    v13 = (^(v13))
                                case 0:
                                    var v42 uint16 = v14
                                    _ = v42
                                }
                            }
                            switch 1 {
                            case -2147483648:
                                fmt.Println(v14, v1)
                            case -607:
                                v43 := GenS6{F0: (v14 ^ (v14 << ((490) & 63))), F1: -2147483648, F2: v14}
                                _ = v43
                            default:
                                switch "hello" {
                                case "\n\t":
                                    fmt.Println(v34)
                                case "":
                                    v34 = v4
                                case "a":
                                    if (v30 != int16(v40)) {
                                        var v44 float32 = float32((((float64((((-2.25 / 1.5) - 0.1) + float32((v32 + float64(0.0))))) / (v32 - (float64(1e10) * ((v32 - float64(0.0)) + 1.5)))) + (v32 / float64(float32(0.1)))) / v32))
                                        var v45 int16 = 32767
                                        _ = len(v2)
                                        _ = v44
                                        _ = v45
                                    } else {
                                        v46 := GenS1{F0: 42, F1: 4294967295}
                                        var v47 string = v4
                                        switch (v47 + v4) {
                                        case "foo bar":
                                            var v48 bool = (-9223372036854775808 > -9223372036854775808)
                                            _ = v48
                                        case "":
                                            v14 -= (v14 / (((v14 / (((v14 >> ((39) & 63))) | 1))) | 1))
                                        default:
                                            fmt.Println(v34)
                                        }
                                        _ = v46
                                        _ = v47
                                    }
                                }
                            }
                        }
                    }
                    for fi49 := 0; fi49 < 9; fi49++ {
                        var v50 int = (fi49 / (((fi49 ^ fi49)) | 1))
                        for fi51 := 0; fi51 < 6; fi51++ {
                            if (!(float32((v32 / (3.14159 * float64(((-2.25 - 0.0) * ((-2.25 + 3.14159) + 0.1)))))) == float32((v32 - float64(3.14159))))) {
                                if (false || (!((v32 + (v32 + -2.25)) < v32))) {
                                    var v52 float64 = 0.1
                                    var v53 uint64 = 2
                                    v54 := GenS7{F0: 203}
                                    switch fi49 {
                                    case 1:
                                        v55 := GenS3{F0: v30, F1: -850, F2: fi49}
                                        _ = v55
                                    }
                                    v50 -= -2147483648
                                    v52 = 1.5
                                    v32 += 0.1
                                    _ = v52
                                    _ = v53
                                    _ = v54
                                }
                            } else {
                                if (36 <= int8(v13)) {
                                    var v56 int16 = v30
                                    fmt.Println(v1, v13)
                                    _ = v56
                                }
                            }
                            v14 = 255
                            v57 := GenS1{F0: 2, F1: 255}
                            v14 -= 2
                            _ = fi51
                            _ = v57
                        }
                        var v58 uint64 = 765
                        _ = fi49
                        _ = v50
                        _ = v58
                    }
                    _ = v40
                } else {
                    v13 = -1
                }
            }
            _ = v39
        }
        _ = v34
    }
    if true {
        var v59 int = 32768
        for fi60 := 0; fi60 < 1; fi60++ {
            v59 += v59
            fmt.Println(v13)
            _ = fi60
        }
        _ = v59
    } else {
        v61 := GenS3{F0: v30, F1: -9223372036854775808, F2: 0}
        var v62 uint64 = 110
        v63 := GenS8{F0: v30, F1: 42, F2: 0}
        _ = v61
        _ = v62
        _ = v63
    }
    switch -2147483648 {
    case 42:
        switch "hello" {
        case "\n\t":
            v64 := GenS3{F0: -336, F1: 0, F2: -2147483648}
            _ = v64
        }
    case 32767:
        v65 := GenS7{F0: 2}
        _ = v65
    }
    v66 := GenS6{F0: v14, F1: -2147483648, F2: v14}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v4
    _ = v66
}
