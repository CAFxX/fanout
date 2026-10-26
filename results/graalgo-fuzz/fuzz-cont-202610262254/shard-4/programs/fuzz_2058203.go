package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint32 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint32; F1 bool; F2 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 uint32; F1 int64; F2 uint32 }
func genFn0(p0 uint64) int16 {
    p0 += (^(p0))
    v7 := GenS1{F0: false, F1: 57, F2: 480}
    _ = p0
    _ = v7
    return 229
}
func genFn1(p0 int32) uint8 {
    var v8 int64 = 1
    v9 := GenS2{F0: (-(p0)), F1: p0}
    _ = p0
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 float64 = ((1e10 * float64(0.0)) * -2.25)
    var v2 int8 = 1
    var v3 uint16 = 32767
    var v4 int64 = 992
    var v5 int8 = v2
    v6 := make(map[string]bool)
    var v10 int = 2147483647
    if false {
        switch v10 {
        case -9223372036854775808:
            v11 := GenS3{F0: 531, F1: true, F2: "\n\t"}
            _ = v11
        case 128:
            v12 := GenS3{F0: 0, F1: ((v2 - (v5 >> ((640) & 63))) > -1), F2: ((("" + "a") + ("hello" + "")) + "foo bar")}
            _ = v12
        case -1:
            v13 := GenS3{F0: 2, F1: (v3 > 1), F2: (("foo bar" + "foo bar") + ("" + (("foo bar" + (("\n\t" + "\n\t") + "")) + "hello")))}
            _ = v13
        default:
            var v14 string = ""
            _ = v14
        }
    } else {
        v15 := []int{}
        _ = v15
    }
    fmt.Println(v4)
    var v16 int = (^(v10))
    v17 := GenS4{F0: uint8(v3)}
    v2 += (-(v5))
    switch "\n\t" {
    case "a":
        if (((v1 * v1) * (float64(0.1) + v1)) == float64(float32(-2.25))) {
            v18 := GenS5{F0: -1, F1: 4294967295}
            _ = v18
        } else {
            if (!true) {
                switch "foo bar" {
                case "hello":
                    v19 := []int{-9223372036854775808, (v16 >> ((4294967295) & 63))}
                    _ = v19
                case "a":
                    for fi20 := 0; fi20 < 14; fi20++ {
                        var v21 uint8 = genFn1(559)
                        v22 := []int{-757, (v10 ^ (fi20 | (v16 >> ((4294967295) & 63)))), 32767, v10}
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                default:
                    v23 := []int{v10, v10, v10, int(v3)}
                    _ = v23
                }
                v10 = v16
                v4 = v4
            } else {
                for fi24 := 0; fi24 < 18; fi24++ {
                    var v25 float32 = 0.0
                    _ = fi24
                    _ = v25
                }
                fmt.Println(v3, v1)
                for fi26 := 0; fi26 < 16; fi26++ {
                    v27 := GenS6{F0: 2, F1: v4, F2: 343}
                    v1 = (v1 + v1)
                    _ = fi26
                    _ = v27
                }
            }
            var v28 int = -35
            if (false && false) {
                v2 = (^(v2))
                switch "\n\t" {
                case "foo bar":
                    _ = len(v6)
                }
                var v29 int32 = 42
                _ = v29
            }
            _ = v28
        }
    case "hello":
        switch "a" {
        case "":
            var v30 uint32 = 4294967295
            _ = v30
        }
    default:
        v31 := GenS5{F0: 42, F1: 4294967295}
        _ = v31
    }
    if true {
        fmt.Println(v16)
        if (1 > -2147483648) {
            var v32 int64 = v4
            if true {
                var v33 uint64 = 782
                _ = v33
            }
            var v34 int = (v16 / (((^(v10))) | 1))
            _ = v32
            _ = v34
        } else {
            var v35 bool = true
            switch 32768 {
            case 4294967295:
                if (!true) {
                    switch v16 {
                    case 32767:
                        var v36 uint32 = 61
                        _ = v36
                    case -2147483648:
                        var v37 uint = 255
                        _ = v37
                    case 196:
                        var v38 uint = 2
                        _ = v38
                    default:
                        if ((v3 < 0) && v35) {
                            if (int32(v4) == -70) {
                                v2 += (v2 % ((v5) | 1))
                            }
                            if true {
                                v10 -= -9223372036854775808
                            } else {
                                v2 = v5
                                for fi39 := 0; fi39 < 14; fi39++ {
                                    if true {
                                        v40 := GenS1{F0: false, F1: 0, F2: 1}
                                        var v41 int64 = (v4 | v4)
                                        var v42 int = 42
                                        _ = v40
                                        _ = v41
                                        _ = v42
                                    } else {
                                        fmt.Println(v10, v1, v2)
                                        for fi43 := 0; fi43 < 14; fi43++ {
                                            switch (fi39 >> ((241) & 63)) {
                                            case -755:
                                                v44 := GenS3{F0: 65535, F1: ("foo bar" > "foo bar"), F2: ("foo bar" + "\n\t")}
                                                _ = v44
                                            case 4294967295:
                                                v16 = fi43
                                            default:
                                                var v45 uint16 = v3
                                                _ = v45
                                            }
                                            v10 = (v16 % ((-264) | 1))
                                            v46 := []int{(v10 >> ((18446744073709551615) & 63))}
                                            _ = fi43
                                            _ = v46
                                        }
                                    }
                                    v6[("" + (("" + "a") + ("foo bar" + (("\n\t" + ("" + "a")) + ""))))] = (float32(v1) > 0.0)
                                    _ = fi39
                                }
                                var v47 uint8 = 2
                                _ = v47
                            }
                            var v48 uint8 = 2
                            _ = v48
                        }
                    }
                    for fi49 := 0; fi49 < 15; fi49++ {
                        if true {
                            var v50 float32 = 1e10
                            v51 := GenS1{F0: (821 == 2), F1: 268, F2: 65535}
                            fmt.Println(v5)
                            _ = v50
                            _ = v51
                        } else {
                            if v35 {
                                var v52 int = 9223372036854775807
                                if v35 {
                                    for fi53 := 0; fi53 < 10; fi53++ {
                                        v2 += v2
                                        var v54 uint = 51
                                        _ = fi53
                                        _ = v54
                                    }
                                    v10 = fi49
                                } else {
                                    var v55 int8 = v2
                                    var v56 int32 = 1
                                    _ = v55
                                    _ = v56
                                }
                                v3 += 1
                                _ = v52
                            }
                        }
                        var v57 float32 = -2.25
                        switch fi49 {
                        case 32767:
                            var v58 float64 = v1
                            _ = v58
                        case 4294967295:
                            v6[("a" + "")] = (((-205 <= -32768) && (2 > 741)) || (18446744073709551615 >= 1))
                        case 1:
                            var v59 int64 = 9223372036854775807
                            _ = v59
                        }
                        _ = fi49
                        _ = v57
                    }
                } else {
                    var v60 string = "a"
                    _ = v60
                }
            default:
                switch "hello" {
                case "\n\t":
                    if v35 {
                        v1 -= v1
                    } else {
                        v61 := []int{v16, (^(v10)), v10, (v10 >> ((1) & 63))}
                        _ = v61
                    }
                case "":
                    if (v35 && false) {
                        fmt.Println(v3)
                        var v62 int32 = 2147483647
                        v62 += (v62 >> ((992) & 63))
                        _ = v62
                    }
                }
            }
            var v63 float32 = (((1e10 - (0.1 - float32((v1 - 3.14159)))) - 1e10) - 0.0)
            _ = v35
            _ = v63
        }
    }
    fmt.Println(v3, v2, v1)
    switch "foo bar" {
    case "a":
        v64 := GenS4{F0: 222}
        _ = v64
    case "":
        v6[("\n\t" + "")] = true
    }
    v65 := GenS5{F0: 42, F1: 0}
    var v66 uint8 = genFn1(-785)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v66:", v66)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
}
