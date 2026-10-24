package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int; F1 int64; F2 uint16 }
type GenS4 struct { F0 uint16; F1 uint32; F2 string }
type GenS5 struct { F0 uint32; F1 int16 }
type GenS6 struct { F0 int8; F1 uint16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 string }
func genFn0(p0 uint, p1 uint16) uint8 {
    var v7 int64 = 474
    v8 := []int{4294967295, -2147483648}
    switch "" {
    case "\n\t":
        var v9 uint64 = 2
        _ = v9
    default:
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := GenS1{F0: (p1 * 382)}
            v12 := GenS2{F0: (p0 - (p0 ^ p0)), F1: 12}
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    var v13 uint16 = (p1 << (((p0 % ((0) | 1))) & 63))
    if ((-1 >= -225) && true) {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        v14 := GenS1{F0: (v13 << (((p0 & (p0 * p0))) & 63))}
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int32) uint64 {
    switch "hello" {
    case "a":
        v15 := []int{-1}
        _ = v15
    case "\n\t":
        v16 := GenS1{F0: 424}
        _ = v16
    }
    v17 := []int{9223372036854775807, 117, 4294967295}
    _ = p0
    _ = v17
    return 0
}

func main() {
    var v1 uint16 = 65535
    var v2 float32 = 1.5
    v3 := []int{-98}
    var v4 int8 = 42
    var v5 int64 = 42
    var v6 int32 = -368
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v18 string = "foo bar"
    if ((!((!(true && false)) || (true && (1 < (v1 ^ (v1 / ((32767) | 1))))))) || (42 < 4294967295)) {
        fmt.Println(v18)
        v3 = append(v3, -848)
    }
    v5 = v5
    if (!true) {
        if true {
            for fi19 := 0; fi19 < 1; fi19++ {
                fmt.Println(v1)
                _ = fi19
            }
            if true {
                var v20 uint8 = 214
                for fi21 := 0; fi21 < 19; fi21++ {
                    var v22 uint32 = 2147483648
                    _ = fi21
                    _ = v22
                }
                if true {
                    switch v18 {
                    case "\n\t":
                        fmt.Println(v2, v18, v18)
                    case "a":
                        v23 := GenS3{F0: 940, F1: (v5 & v5), F2: v1}
                        _ = v23
                    default:
                        var v24 int16 = 32767
                        _ = v24
                    }
                    var v25 float64 = -2.25
                    _ = v25
                }
                _ = v20
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            v6 = v6
        } else {
            v3 = append(v3, -1)
        }
    }
    for fi26 := 0; fi26 < 2; fi26++ {
        fmt.Println(v5, v4, v5)
        v4 += (v4 ^ v4)
        _ = fi26
    }
    var v27 uint = 65535
    if ((3.14159 * 3.14159) >= float64(1.5)) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if len(v3) > 0 {
            v3[0] = 128
        }
    } else {
        var v28 uint8 = 175
        for fi29 := 0; fi29 < 15; fi29++ {
            for fi30 := 0; fi30 < 15; fi30++ {
                switch v18 {
                case "a":
                    v31 := GenS4{F0: (v1 >> ((v27) & 63)), F1: 42, F2: v18}
                    _ = v31
                }
                _ = fi30
            }
            fmt.Println(fi29, v1)
            if (!(1 >= -1)) {
                var v32 uint16 = v1
                _ = v32
            }
            _ = fi29
        }
        for fi33 := 0; fi33 < 13; fi33++ {
            v5 += -132
            for fi34 := 0; fi34 < 6; fi34++ {
                v2 -= 0.0
                _ = fi34
            }
            _ = fi33
        }
        _ = v28
    }
    switch 0 {
    case 42:
        fmt.Println(v5)
    default:
        var v35 float64 = 1e10
        _ = v35
    }
    if (false && true) {
        for fi36 := 0; fi36 < 10; fi36++ {
            switch v18 {
            case "":
                v1 += (v1 * 1)
            case "\n\t":
                switch ((v18 + v18) + "foo bar") {
                case "hello":
                    v37 := GenS4{F0: (^(v1)), F1: 818, F2: v18}
                    _ = v37
                case "foo bar":
                    if true {
                        fmt.Println(v5, v6)
                    }
                case "":
                    v38 := GenS5{F0: 2, F1: -1}
                    _ = v38
                default:
                    if len(v3) > 0 {
                        v3[0] = -1
                    }
                }
            case "hello":
                v39 := GenS4{F0: (^(v1)), F1: 678, F2: "a"}
                _ = v39
            }
            v40 := GenS6{F0: v4, F1: (^(v1))}
            _ = fi36
            _ = v40
        }
        if (!true) {
            if len(v3) > 0 {
                v3[0] = 9223372036854775807
            }
            if false {
                for fi41 := 0; fi41 < 4; fi41++ {
                    if len(v3) > 0 {
                        v3[0] = 127
                    }
                    if true {
                        v27 -= (v27 / (((v27 ^ (v27 % ((v27) | 1)))) | 1))
                        for fi42 := 0; fi42 < 17; fi42++ {
                            if false {
                                v43 := GenS6{F0: (v4 + v4), F1: (v1 + 1)}
                                var v44 bool = ((-(v5)) < v5)
                                _ = v43
                                _ = v44
                            } else {
                                v45 := GenS7{F0: "a"}
                                switch (v18 + v18) {
                                case "\n\t":
                                    v6 -= 72
                                case "foo bar":
                                    var v46 int = (fi42 + fi42)
                                    _ = v46
                                case "a":
                                    var v47 int8 = (v4 | 100)
                                    _ = v47
                                }
                                v4 += v4
                                _ = v45
                            }
                            if false {
                                if len(v3) > 0 {
                                    v3[0] = (fi42 << ((v27) & 63))
                                }
                                fmt.Println(fi41, v18)
                                if false {
                                    _ = len(v3)
                                }
                            } else {
                                if true {
                                    var v48 uint64 = 18446744073709551615
                                    v27 += v27
                                    _ = v48
                                } else {
                                    v1 -= v1
                                    var v49 float64 = -2.25
                                    switch "\n\t" {
                                    case "\n\t":
                                        v18 = ("hello" + v18)
                                    }
                                    _ = len(v3)
                                    _ = v49
                                }
                            }
                            fmt.Println(v5)
                            _ = fi42
                        }
                        fmt.Println(v5, v18, v4)
                    } else {
                        if len(v3) > 0 {
                            v3[0] = (fi41 * fi41)
                        }
                    }
                    var v50 string = v18
                    _ = fi41
                    _ = v50
                }
            }
            switch v18 {
            case "a":
                var v51 uint16 = (v1 - 0)
                _ = v51
            default:
                fmt.Println(v27)
            }
        } else {
            var v52 int8 = (^(v4))
            v1 -= (v1 * 65535)
            _ = v52
        }
    }
    var v53 int32 = v6
    switch -245 {
    case 32767:
        switch v18 {
        case "\n\t":
            if false {
                v54 := GenS8{F0: v18}
                for fi55 := 0; fi55 < 10; fi55++ {
                    var v56 bool = (42 > 0)
                    var v57 int8 = (^(v4))
                    _ = fi55
                    _ = v56
                    _ = v57
                }
                var v58 float32 = 0.1
                _ = v54
                _ = v58
            }
        }
    case 32768:
        fmt.Println(v6, v27, v53)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
}
