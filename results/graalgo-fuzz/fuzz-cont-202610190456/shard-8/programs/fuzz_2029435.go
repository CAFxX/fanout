package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 uint; F1 bool; F2 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16; F1 int8; F2 int8 }
type GenS6 struct { F0 int16 }
func genFn0() uint8 {
    var v5 bool = false
    switch 286 {
    case 42:
        var v6 int8 = 0
        _ = v6
    case 0:
        var v7 uint64 = 4294967295
        _ = v7
    default:
        fmt.Println(v5)
    }
    fmt.Println(v5)
    var v8 uint8 = 123
    _ = v5
    _ = v8
    return v8
}
func genFn1(p0 uint, p1 bool) int {
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 float32 = (((1.5 + (0.0 / -2.25)) - (0.1 * 0.0)) / 1.0)
        p0 -= p0
        var v11 uint = 764
        _ = fi9
        _ = v10
        _ = v11
    }
    switch "a" {
    case "":
        v12 := []int{9223372036854775807, 2147483647}
        _ = v12
    }
    if (("hello" + ("foo bar" + "")) >= (("foo bar" + "\n\t") + (("hello" + ("\n\t" + "a")) + "a"))) {
        switch -1 {
        case 9223372036854775807:
            p0 += (p0 + (p0 ^ 4294967295))
        case 2147483647:
            for fi13 := 0; fi13 < 13; fi13++ {
                switch 923 {
                case 9223372036854775807:
                    for fi14 := 0; fi14 < 20; fi14++ {
                        if (-1 <= 127) {
                            var v15 uint64 = 18446744073709551615
                            p0 += p0
                            fmt.Println(p0, fi13)
                            fmt.Println(p1, p1, p0)
                            _ = v15
                        } else {
                            v16 := GenS1{F0: 127}
                            v17 := []int{}
                            _ = v16
                            _ = v17
                        }
                        if (660 < (p0 / ((p0) | 1))) {
                            v18 := GenS2{F0: 128, F1: 0}
                            v19 := GenS3{F0: p0, F1: p1, F2: p1}
                            p1 = false
                            _ = v18
                            _ = v19
                        }
                        p0 = 2
                        fmt.Println(p0, p0, fi14)
                        p0 = (p0 - 933)
                        _ = fi14
                    }
                default:
                    var v20 string = "hello"
                    _ = v20
                }
                switch fi13 {
                case 2147483647:
                    p0 -= (p0 % ((0) | 1))
                case 9223372036854775807:
                    v21 := GenS4{F0: ""}
                    _ = v21
                case 127:
                    var v22 uint8 = 1
                    _ = v22
                }
                _ = fi13
            }
        case 32767:
            p0 += p0
        default:
            for fi23 := 0; fi23 < 12; fi23++ {
                for fi24 := 0; fi24 < 3; fi24++ {
                    v25 := GenS5{F0: 619, F1: 100, F2: 80}
                    p1 = p1
                    _ = fi24
                    _ = v25
                }
                _ = fi23
            }
        }
        fmt.Println(p0, p0)
        if ((p0 << (((^(p0))) & 63)) < (p0 + (p0 ^ p0))) {
            if p1 {
                v26 := []int{}
                p1 = (!((false && (!false)) && (906 > 42)))
                p1 = p1
                _ = v26
            }
            switch "foo bar" {
            case "a":
                for fi27 := 0; fi27 < 16; fi27++ {
                    var v28 uint64 = 2
                    var v29 uint64 = v28
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            default:
                p1 = (42 < 9223372036854775807)
            }
            for fi30 := 0; fi30 < 4; fi30++ {
                for fi31 := 0; fi31 < 7; fi31++ {
                    switch ((("hello" + "foo bar") + (("hello" + "foo bar") + ((("a" + ("foo bar" + "\n\t")) + "hello") + "\n\t"))) + "foo bar") {
                    case "a":
                        p1 = false
                    default:
                        p1 = (!(!(p1 || true)))
                    }
                    _ = fi31
                }
                p0 = 1
                v32 := GenS5{F0: -1, F1: 127, F2: 127}
                _ = fi30
                _ = v32
            }
        } else {
            var v33 bool = p1
            for fi34 := 0; fi34 < 16; fi34++ {
                for fi35 := 0; fi35 < 4; fi35++ {
                    p0 += p0
                    var v36 int32 = 2147483647
                    _ = fi35
                    _ = v36
                }
                _ = fi34
            }
            _ = v33
        }
    }
    _ = p0
    _ = p1
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]bool)
    var v3 float64 = (0.0 + -2.25)
    v4 := []int{42}
    v3 -= v3
    v3 -= (float64(1.5) + float64(0.1))
    v37 := GenS3{F0: 18446744073709551615, F1: (false || (!false)), F2: true}
    if (0 > 1) {
        var v38 uint64 = 158
        var v39 uint8 = 127
        switch 32767 {
        case 128:
            var v40 uint16 = 302
            _ = v40
        default:
            v41 := GenS5{F0: 1, F1: 1, F2: 93}
            _ = v41
        }
        _ = v38
        _ = v39
    } else {
        fmt.Println(v3)
        v42 := GenS3{F0: 422, F1: true, F2: (2 >= 327)}
        var v43 bool = (((("\n\t" + "a") + "foo bar") >= (("hello" + "a") + "\n\t")) || (2 <= 319))
        _ = v42
        _ = v43
    }
    var v44 int64 = 42
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 = 3.14159
    delete(v2, "a")
    for fi45 := 0; fi45 < 11; fi45++ {
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
        var v46 int64 = (^(v44))
        v46 = v44
        _ = fi45
        _ = v46
    }
    switch -1 {
    case 9223372036854775807:
        switch "" {
        case "\n\t":
            v1["a"] = (-128 != -32)
        case "hello":
            if (169 < 1) {
                for fi47 := 0; fi47 < 17; fi47++ {
                    v48 := GenS3{F0: 0, F1: false, F2: true}
                    v3 = v3
                    _ = fi47
                    _ = v48
                }
            } else {
                v3 -= 0.1
                switch ("foo bar" + "hello") {
                case "hello":
                    var v49 float64 = ((0.1 * (v3 * v3)) * (v3 + (v3 + v3)))
                    _ = v49
                case "a":
                    fmt.Println(v44, v3)
                }
            }
        }
    case -430:
        switch "foo bar" {
        case "foo bar":
            var v50 uint16 = 2
            _ = v50
        }
    default:
        if true {
            v44 = (v44 | -133)
            if true {
                fmt.Println(v3, v3)
                for fi51 := 0; fi51 < 1; fi51++ {
                    if false {
                        if true {
                            var v52 uint32 = 331
                            _ = v52
                        } else {
                            fmt.Println(fi51)
                        }
                    } else {
                        v53 := GenS5{F0: 0, F1: 100, F2: -123}
                        v54 := GenS1{F0: genFn0()}
                        v55 := GenS6{F0: 42}
                        _ = v53
                        _ = v54
                        _ = v55
                    }
                    var v56 float64 = 3.14159
                    _ = fi51
                    _ = v56
                }
            } else {
                fmt.Println(v44, v3)
                v57 := GenS4{F0: "\n\t"}
                switch ("hello" + "hello") {
                case "foo bar":
                    if (2147483647 < 65535) {
                        for fi58 := 0; fi58 < 3; fi58++ {
                            var v59 bool = (-1 >= 0)
                            fmt.Println(fi58, fi58)
                            _ = fi58
                            _ = v59
                        }
                        if true {
                            var v60 int16 = 130
                            fmt.Println(v3, v44, v3)
                            var v61 int32 = 2147483647
                            _ = v60
                            _ = v61
                        }
                    }
                case "a":
                    switch ("hello" + "") {
                    case "hello":
                        var v62 int8 = 9
                        _ = v62
                    default:
                        v3 = v3
                    }
                case "hello":
                    switch "foo bar" {
                    case "a":
                        var v63 string = ("\n\t" + "")
                        _ = v63
                    case "\n\t":
                        var v64 string = "\n\t"
                        _ = v64
                    case "foo bar":
                        switch "a" {
                        case "hello":
                            v3 -= float64(1e10)
                        case "\n\t":
                            v3 = v3
                        default:
                            switch 1 {
                            case -9223372036854775808:
                                v65 := GenS1{F0: 1}
                                _ = v65
                            case -420:
                                if len(v4) > 0 {
                                    v4[0] = 1
                                }
                            case 2147483647:
                                for fi66 := 0; fi66 < 9; fi66++ {
                                    v44 = 566
                                    _ = fi66
                                }
                            }
                        }
                    default:
                        v67 := GenS6{F0: 1000}
                        _ = v67
                    }
                default:
                    var v68 uint16 = 2
                    _ = v68
                }
                _ = v57
            }
        }
    }
    v2["hello"] = (255 != 18446744073709551615)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v44
}
