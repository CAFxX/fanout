package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 int64; F1 int64 }
func genFn0(p0 int64) int8 {
    v5 := GenS1{F0: -90, F1: 2}
    v6 := GenS2{F0: 160, F1: 324}
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 float32 = (0.0 + float32(((float64(3.14159) - float64(0.0)) * 3.14159)))
        _ = fi7
        _ = v8
    }
    var v9 int8 = 1
    for fi10 := 0; fi10 < 13; fi10++ {
        switch (fi10 >> ((18446744073709551615) & 63)) {
        case 42:
            p0 -= (p0 ^ p0)
        case -1:
            fmt.Println(p0)
        case 0:
            v9 += (v9 ^ (-(v9)))
        default:
            v11 := []int{127, 32767, fi10}
            _ = v11
        }
        v12 := GenS2{F0: 2, F1: 1}
        _ = fi10
        _ = v12
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v9
    return (v9 | (v9 >> ((42) & 63)))
}
func genFn1(p0 int, p1 int64, p2 int8) string {
    fmt.Println(p0, p2, p0)
    p1 -= p1
    v13 := GenS2{F0: 128, F1: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return ("hello" + ("\n\t" + "\n\t"))
}

func main() {
    var v1 uint64 = 4294967295
    v2 := []int{127, 42, 9223372036854775807}
    var v3 uint32 = 2147483647
    v4 := make(map[string]bool)
    if true {
        if (!false) {
            _ = len(v2)
        }
        v1 = v1
    } else {
        var v14 uint = 5
        _ = v14
    }
    fmt.Println(v1, v3)
    var v15 int64 = 9223372036854775807
    fmt.Println(v3, v1, v3)
    var v16 int8 = 100
    for fi17 := 0; fi17 < 10; fi17++ {
        v3 -= 42
        fmt.Println(v15)
        _ = fi17
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 uint8 = 237
        v19 += v19
        _ = fi18
        _ = v19
    }
    if false {
        switch ("" + "\n\t") {
        case "foo bar":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            var v20 uint16 = 367
            _ = v20
        }
        v1 = (v1 + (v1 | (v1 + v1)))
    }
    v16 = (-(v16))
    if true {
        switch 2147483647 {
        case -2147483648:
            switch 32768 {
            case 4294967295:
                var v21 float32 = (float32((1e10 - 1.5)) / 0.1)
                _ = v21
            default:
                fmt.Println(v16, v15)
            }
        default:
            fmt.Println(v16, v1)
        }
        v2 = append(v2, 128)
        if true {
            if false {
                for fi22 := 0; fi22 < 11; fi22++ {
                    if (false || false) {
                        v15 = v15
                        fmt.Println(v15)
                    } else {
                        _ = len(v2)
                        var v23 uint8 = 20
                        _ = v23
                    }
                    switch -1 {
                    case 42:
                        switch "foo bar" {
                        case "hello":
                            _ = len(v4)
                        case "a":
                            if false {
                                v2 = append(v2, (fi22 - fi22))
                                for fi24 := 0; fi24 < 9; fi24++ {
                                    if true {
                                        if true {
                                            var v25 int16 = -1
                                            v1 -= v1
                                            _ = v25
                                        }
                                        v2 = append(v2, -9223372036854775808)
                                        v3 += v3
                                        fmt.Println(v16)
                                    } else {
                                        v16 = v16
                                        v15 = (v15 - -9223372036854775808)
                                        var v26 float32 = (float32(0.1) * (-2.25 - 1e10))
                                        var v27 uint8 = 1
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v28 string = "\n\t"
                                            switch v28 {
                                            case "a":
                                                switch fi24 {
                                                case 468:
                                                    var v29 string = "foo bar"
                                                    _ = v29
                                                case -9223372036854775808:
                                                    if (!true) {
                                                        v16 += (v16 * (v16 + v16))
                                                    } else {
                                                        if true {
                                                            var v30 int32 = 60
                                                            _ = v30
                                                        }
                                                        var v31 int64 = (v15 % ((v15) | 1))
                                                        switch 32767 {
                                                        case -883:
                                                            if true {
                                                                switch (fi22 * int(v1)) {
                                                                case -1:
                                                                    _ = len(v4)
                                                                case 982:
                                                                    v32 := GenS3{F0: false}
                                                                    _ = v32
                                                                case -692:
                                                                    if (42 == 444) {
                                                                        var v33 float32 = 0.0
                                                                        if len(v2) > 0 {
                                                                            v2[0] = (^(fi24))
                                                                        }
                                                                        _ = v33
                                                                    }
                                                                }
                                                            }
                                                        case -137:
                                                            fmt.Println(v3, v27)
                                                        case 32767:
                                                            var v34 string = v28
                                                            _ = v34
                                                        }
                                                        _ = v31
                                                    }
                                                default:
                                                    v4[v28] = (false || (v15 <= (v15 ^ 0)))
                                                }
                                            }
                                            _ = v28
                                        case "\n\t":
                                            var v35 float32 = (-2.25 - (((float32(0.1) + (v26 - (v26 / 3.14159))) + 0.0) / ((v26 * v26) + 3.14159)))
                                            _ = v35
                                        default:
                                            v16 = 100
                                        }
                                        _ = v26
                                        _ = v27
                                    }
                                    _ = fi24
                                }
                            }
                        case "\n\t":
                            switch 127 {
                            case 933:
                                for fi36 := 0; fi36 < 6; fi36++ {
                                    if ("a" == "hello") {
                                        var v37 uint64 = v1
                                        if true {
                                            v38 := GenS1{F0: (v16 >> ((65535) & 63)), F1: 0}
                                            v39 := GenS1{F0: v16, F1: 1}
                                            fmt.Println(v16, v3, v37)
                                            _ = v38
                                            _ = v39
                                        }
                                        v3 -= v3
                                        _ = v37
                                    } else {
                                        var v40 string = ""
                                        _ = v40
                                    }
                                    _ = fi36
                                }
                            case 42:
                                var v41 float32 = (float32(((((0.0 - 0.0) / -2.25) * (0.1 - -2.25)) + (((1e10 / 1.5) + 1.5) - (1e10 - ((1e10 * (-2.25 / (-2.25 / ((float64((0.1 + 3.14159)) - float64(((float32((0.0 / float64((1.5 + 0.0)))) - 3.14159) + -2.25))) + float64((float32(-2.25) - 1.5)))))) / float64(1e10)))))) + 3.14159)
                                _ = v41
                            default:
                                v42 := GenS2{F0: 128, F1: v3}
                                _ = v42
                            }
                        }
                    case 352:
                        var v43 uint64 = (^(v1))
                        _ = v43
                    case -233:
                        var v44 int16 = -54
                        _ = v44
                    default:
                        var v45 string = "foo bar"
                        _ = v45
                    }
                    _ = fi22
                }
                var v46 int64 = v15
                var v47 int = 42
                _ = v46
                _ = v47
            } else {
                for fi48 := 0; fi48 < 3; fi48++ {
                    if true {
                        v2 = append(v2, -9223372036854775808)
                        fmt.Println(v16, v1, v15)
                    }
                    if true {
                        fmt.Println(v15)
                        v49 := GenS3{F0: false}
                        _ = v49
                    }
                    _ = len(v2)
                    _ = fi48
                }
                var v50 int8 = -1
                _ = v50
            }
        }
    }
    fmt.Println(v3, v1)
    var v51 bool = false
    v52 := GenS4{F0: 255, F1: (!v51)}
    v51 = true
    var v53 float64 = 1e10
    switch 128 {
    case 0:
        _ = len(v2)
    case 324:
        var v54 uint = 663
        _ = v54
    case -2147483648:
        if v51 {
            var v55 int16 = 32767
            _ = v55
        } else {
            for fi56 := 0; fi56 < 1; fi56++ {
                v16 -= (v16 ^ v16)
                v57 := GenS2{F0: 1, F1: v3}
                if (18446744073709551615 != (v1 / ((4294967295) | 1))) {
                    var v58 uint = 65535
                    switch "a" {
                    case "":
                        for fi59 := 0; fi59 < 5; fi59++ {
                            var v60 float64 = (1.5 - -2.25)
                            var v61 bool = (v51 && ((v16 << ((v58) & 63)) <= (^(v16))))
                            v53 -= v53
                            v51 = false
                            if (!v51) {
                                v62 := GenS5{F0: v15, F1: v15}
                                switch 32768 {
                                case 2147483647:
                                    var v63 float32 = 0.0
                                    _ = v63
                                case 42:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                default:
                                    var v64 uint64 = v1
                                    _ = v64
                                }
                                fmt.Println(v60, v53, v61)
                                _ = v62
                            }
                            _ = fi59
                            _ = v60
                            _ = v61
                        }
                    default:
                        v65 := GenS2{F0: 2, F1: 128}
                        _ = v65
                    }
                    _ = v58
                }
                _ = fi56
                _ = v57
            }
            v51 = (18446744073709551615 < 119)
            delete(v4, "\n\t")
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v52
    _ = v53
}
