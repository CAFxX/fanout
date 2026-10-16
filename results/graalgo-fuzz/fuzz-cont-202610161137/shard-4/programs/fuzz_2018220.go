package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS4 struct { F0 int32; F1 int16; F2 int }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 int8; F1 uint16 }
type GenS7 struct { F0 bool; F1 int16; F2 uint }
type GenS8 struct { F0 bool; F1 int16 }
type GenS9 struct { F0 string; F1 uint; F2 uint32 }
type GenS10 struct { F0 uint8; F1 int32 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint64 }
func genFn0(p0 string, p1 string) uint32 {
    v4 := GenS1{F0: (p1 + p1), F1: 18446744073709551615}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v4
    return 1
}

func main() {
    var v1 string = ("hello" + ((("" + "\n\t") + "") + "foo bar"))
    v2 := []int{}
    v3 := make(map[string]string)
    var v5 int8 = -1
    var v6 int32 = -752
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint8 = 2
        _ = fi7
        _ = v8
    }
    switch "a" {
    case "hello":
        if false {
            v9 := GenS2{F0: 9223372036854775807}
            v2 = append(v2, 127)
            _ = v9
        }
    case "foo bar":
        var v10 uint16 = 65535
        _ = v10
    }
    if false {
        switch "foo bar" {
        case "":
            for fi11 := 0; fi11 < 2; fi11++ {
                v12 := GenS3{F0: 826, F1: 1, F2: 290}
                var v13 int = 2147483647
                v13 += -9223372036854775808
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        fmt.Println(v6)
    }
    if (42 < 127) {
        v6 = v6
        var v14 uint = 18446744073709551615
        _ = v14
    } else {
        delete(v3, v1)
        fmt.Println(v6, v5)
    }
    v6 -= (v6 | (v6 ^ v6))
    v15 := GenS4{F0: v6, F1: -32768, F2: 42}
    v6 -= 538
    v3[v1] = (v1 + ((v1 + (v1 + ("a" + "a"))) + "a"))
    if false {
        switch v1 {
        case "\n\t":
            for fi16 := 0; fi16 < 14; fi16++ {
                fmt.Println(v5, fi16, fi16)
                _ = fi16
            }
        case "a":
            v17 := GenS3{F0: 2, F1: 0, F2: 0}
            _ = v17
        case "hello":
            v6 = -991
        }
        var v18 int64 = int64(v5)
        if false {
            fmt.Println(v6, v6, v1)
            for fi19 := 0; fi19 < 16; fi19++ {
                var v20 int16 = -392
                if true {
                    v18 -= 1
                    if (!((false && (v5 >= v5)) && false)) {
                        for fi21 := 0; fi21 < 1; fi21++ {
                            fmt.Println(v20, v20, v1)
                            switch v1 {
                            case "a":
                                fmt.Println(v1, fi21)
                            }
                            if true {
                                var v22 int64 = (v18 % (((v18 >> ((275) & 63))) | 1))
                                _ = v22
                            } else {
                                fmt.Println(v20, v20)
                                switch (fi19 ^ 0) {
                                case -2147483648:
                                    var v23 float32 = ((float32(1.5) * 1.5) + 0.0)
                                    _ = v23
                                case 127:
                                    v5 = -24
                                case 128:
                                    if true {
                                        v20 += (^(v20))
                                    }
                                default:
                                    v24 := GenS1{F0: "a", F1: 4294967295}
                                    _ = v24
                                }
                                v25 := GenS5{F0: v6, F1: 2}
                                v6 += (v6 + v6)
                                v6 -= v6
                                _ = v25
                            }
                            _ = fi21
                        }
                        var v26 float64 = -2.25
                        if true {
                            for fi27 := 0; fi27 < 17; fi27++ {
                                var v28 uint = 42
                                _ = fi27
                                _ = v28
                            }
                            delete(v3, "foo bar")
                            var v29 int32 = (v6 % ((768) | 1))
                            _ = v29
                        }
                        _ = v26
                    }
                } else {
                    for fi30 := 0; fi30 < 16; fi30++ {
                        if (143 <= 0) {
                            v31 := v3[((v1 + v1) + (v1 + v1))]
                            _ = v31
                        }
                        switch (fi30 << ((2) & 63)) {
                        case 4294967295:
                            fmt.Println(fi30)
                            v32 := GenS6{F0: (v5 + v5), F1: 2}
                            _ = v32
                        case 337:
                            var v33 uint64 = 0
                            v34 := GenS4{F0: v6, F1: 32767, F2: -316}
                            v6 = 469
                            if ((0.0 < 0.1) || (v5 > v5)) {
                                var v35 bool = (((!false) && ((v20 + (-(v20))) != (v20 >> ((42) & 63)))) || true)
                                v36 := GenS5{F0: -492, F1: 406}
                                _ = v35
                                _ = v36
                            }
                            _ = v33
                            _ = v34
                        case 201:
                            switch (^(fi19)) {
                            case 32767:
                                if (v20 < (v20 << ((18446744073709551615) & 63))) {
                                    switch v1 {
                                    case "":
                                        v37 := GenS5{F0: 1, F1: genFn0("a", "foo bar")}
                                        v5 -= (v5 / ((v5) | 1))
                                        _ = len(v2)
                                        _ = v37
                                    case "hello":
                                        if (true && (128 > 1)) {
                                            if (797 != 1) {
                                                switch "\n\t" {
                                                case "\n\t":
                                                    _ = len(v2)
                                                default:
                                                    v38 := GenS7{F0: false, F1: -386, F2: 0}
                                                    _ = v38
                                                }
                                                _ = len(v3)
                                                if true {
                                                    v39 := GenS7{F0: ((v6 * (v6 ^ 1)) < v6), F1: 1, F2: 65535}
                                                    var v40 uint16 = 2
                                                    _ = v39
                                                    _ = v40
                                                }
                                            }
                                        } else {
                                            if false {
                                                v41 := GenS2{F0: 551}
                                                _ = v41
                                            } else {
                                                fmt.Println(v1, v5, v20)
                                                switch v1 {
                                                case "foo bar":
                                                    switch (fi19 ^ (fi30 * fi30)) {
                                                    case 4294967295:
                                                        var v42 uint = 214
                                                        _ = v42
                                                    case -441:
                                                        if len(v2) > 0 {
                                                            v2[0] = (fi30 & (^(fi30)))
                                                        }
                                                    case -595:
                                                        v6 = -1
                                                    }
                                                case "hello":
                                                    v43 := GenS8{F0: (true && false), F1: v20}
                                                    _ = v43
                                                }
                                                var v44 string = "\n\t"
                                                _ = v44
                                            }
                                        }
                                    case "a":
                                        fmt.Println(v20)
                                    }
                                }
                            case 128:
                                var v45 float32 = 3.14159
                                _ = v45
                            }
                        default:
                            fmt.Println(v20, v20)
                        }
                        if len(v2) > 0 {
                            v2[0] = (fi30 / (((fi30 * (-(fi30)))) | 1))
                        }
                        switch (v1 + v1) {
                        case "a":
                            v20 += (v20 - (v20 & v20))
                        case "foo bar":
                            delete(v3, v1)
                        default:
                            _ = len(v3)
                        }
                        if true {
                            if false {
                                v6 = (v6 % ((v6) | 1))
                                if false {
                                    var v46 int8 = v5
                                    v47 := GenS9{F0: v1, F1: 2, F2: 1}
                                    _ = v46
                                    _ = v47
                                } else {
                                    v6 -= (v6 % ((v6) | 1))
                                    var v48 int32 = v6
                                    _ = v48
                                }
                                fmt.Println(fi19, v20, fi19)
                            }
                        }
                        _ = fi30
                    }
                    v6 += (v6 * v6)
                }
                if (!false) {
                    switch (fi19 / (((fi19 % ((fi19) | 1))) | 1)) {
                    case 32768:
                        var v49 int = fi19
                        _ = v49
                    default:
                        v50 := GenS10{F0: 127, F1: -102}
                        _ = v50
                    }
                    switch fi19 {
                    case 4294967295:
                        v5 = (v5 * v5)
                    case -9223372036854775808:
                        v3["\n\t"] = v1
                    case 9223372036854775807:
                        if false {
                            fmt.Println(v5, v5, v20)
                        } else {
                            fmt.Println(v18, fi19)
                        }
                    default:
                        fmt.Println(fi19, fi19)
                    }
                } else {
                    for fi51 := 0; fi51 < 15; fi51++ {
                        v18 = 9223372036854775807
                        var v52 uint8 = 24
                        v53 := GenS6{F0: (v5 >> ((0) & 63)), F1: 1}
                        var v54 uint32 = 0
                        var v55 bool = ((v54 ^ v54) == 1)
                        _ = fi51
                        _ = v52
                        _ = v53
                        _ = v54
                        _ = v55
                    }
                }
                _ = fi19
                _ = v20
            }
        } else {
            for fi56 := 0; fi56 < 1; fi56++ {
                v57 := GenS11{F0: 146}
                v58 := GenS5{F0: (^(v6)), F1: 2}
                _ = fi56
                _ = v57
                _ = v58
            }
        }
        _ = v18
    } else {
        var v59 uint16 = 32767
        _ = v59
    }
    var v60 int64 = -9223372036854775808
    var v61 int16 = 0
    if false {
        if (65535 != 484) {
            for fi62 := 0; fi62 < 16; fi62++ {
                var v63 uint = 2
                _ = fi62
                _ = v63
            }
        }
    }
    if true {
        if (!true) {
            v64 := GenS12{F0: 18446744073709551615}
            _ = v64
        }
    }
    v65 := GenS6{F0: 100, F1: 117}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v60
    _ = v61
    _ = v65
}
