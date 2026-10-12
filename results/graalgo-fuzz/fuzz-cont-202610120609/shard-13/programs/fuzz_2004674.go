package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 uint8; F2 bool }
type GenS4 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS5 struct { F0 int; F1 int64; F2 uint32 }
type GenS6 struct { F0 uint; F1 bool; F2 uint32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint16; F1 uint32; F2 bool }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint8; F1 uint8 }
func genFn0(p0 int32, p1 bool, p2 int32) string {
    v5 := GenS1{F0: 42, F1: 4294967295, F2: 4294967295}
    for fi6 := 0; fi6 < 9; fi6++ {
        if (((("hello" + "") + "hello") + (("\n\t" + "hello") + ("" + "hello"))) == ("\n\t" + "a")) {
            for fi7 := 0; fi7 < 12; fi7++ {
                p2 = (p0 - 2147483647)
                _ = fi7
            }
            var v8 uint16 = 255
            if (42 >= 1) {
                for fi9 := 0; fi9 < 7; fi9++ {
                    if p1 {
                        switch "\n\t" {
                        case "a":
                            v8 -= v8
                        case "":
                            if true {
                                switch fi9 {
                                case 42:
                                    v8 += (v8 % ((2) | 1))
                                case 4294967295:
                                    fmt.Println(p2)
                                case 1:
                                    p2 -= -2147483648
                                default:
                                    v10 := []int{(fi9 / ((fi9) | 1))}
                                    _ = v10
                                }
                            } else {
                                var v11 bool = (4294967295 == 42)
                                _ = v11
                            }
                        case "\n\t":
                            var v12 uint = 255
                            _ = v12
                        }
                    }
                    v13 := []int{-1, fi6}
                    var v14 bool = false
                    _ = fi9
                    _ = v13
                    _ = v14
                }
                v15 := []int{-9223372036854775808, -2147483648, 127, (^(fi6))}
                switch "\n\t" {
                case "foo bar":
                    _ = len(v15)
                case "hello":
                    if true {
                        v16 := GenS2{F0: ("a" + "a")}
                        _ = v16
                    }
                default:
                    switch (fi6 ^ 0) {
                    case -373:
                        var v17 uint32 = 2
                        _ = v17
                    case 127:
                        if (42 == 1) {
                            for fi18 := 0; fi18 < 1; fi18++ {
                                v19 := GenS3{F0: 1, F1: 127, F2: p1}
                                v8 = v8
                                _ = fi18
                                _ = v19
                            }
                            _ = len(v15)
                        }
                    case 168:
                        p1 = true
                    default:
                        if (!(159 >= 127)) {
                            v8 = (v8 & v8)
                            switch 798 {
                            case -125:
                                p2 += p0
                            case -9223372036854775808:
                                for fi20 := 0; fi20 < 16; fi20++ {
                                    v15 = append(v15, (fi20 | fi20))
                                    if p1 {
                                        if true {
                                            p2 -= p2
                                            p0 = (p0 | (-(p2)))
                                        } else {
                                            fmt.Println(p0, p1, fi6)
                                            v8 += (v8 / ((v8) | 1))
                                            v21 := GenS4{F0: -32768, F1: 1, F2: v8}
                                            _ = v21
                                        }
                                    }
                                    _ = fi20
                                }
                            case 32767:
                                v22 := GenS5{F0: (fi6 * fi6), F1: 9223372036854775807, F2: 65535}
                                _ = v22
                            default:
                                if false {
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                }
                            }
                            fmt.Println(v8, p0, fi6)
                        } else {
                            var v23 uint32 = uint32(v8)
                            _ = v23
                        }
                    }
                }
                _ = v15
            } else {
                v24 := GenS2{F0: ("\n\t" + (("a" + "foo bar") + "\n\t"))}
                var v25 uint32 = 2147483647
                _ = v24
                _ = v25
            }
            _ = v8
        } else {
            p2 += 2147483647
        }
        _ = fi6
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        var v27 uint32 = 1
        _ = fi26
        _ = v27
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return "\n\t"
}

func main() {
    var v1 int8 = 0
    var v2 bool = true
    var v3 int32 = 1
    v4 := []uint32{}
    v28 := GenS6{F0: 0, F1: (32768 > -1), F2: 198}
    if v2 {
        fmt.Println(v2, v1)
    }
    v2 = v2
    fmt.Println(v3)
    v29 := GenS7{F0: 4294967295}
    if len(v4) > 0 {
        v4[0] = 686
    }
    v2 = (-2.25 == -2.25)
    v30 := GenS8{F0: 0, F1: 216, F2: (!v2)}
    fmt.Println(v3)
    v2 = (v2 && v2)
    switch "\n\t" {
    case "foo bar":
        v31 := GenS8{F0: 255, F1: 4294967295, F2: (false || (!(false && (v2 || true))))}
        _ = v31
    case "\n\t":
        v1 += 100
    }
    v3 = (v3 % ((v3) | 1))
    var v32 bool = (v2 && false)
    v32 = ((!(-9223372036854775808 < -696)) && (203 == 128))
    var v33 string = ("a" + "")
    if (4294967295 != 2) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v34 uint = 0
        _ = v34
    } else {
        switch "" {
        case "\n\t":
            fmt.Println(v32, v33, v3)
        default:
            fmt.Println(v1, v3)
        }
    }
    switch (((v33 + v33) + "\n\t") + v33) {
    case "\n\t":
        var v35 int16 = -134
        _ = v35
    case "foo bar":
        v1 += -119
    case "hello":
        if false {
            if false {
                if ((0.0 + float32(float64(((0.0 + 1.5) + (1e10 - 0.1))))) > (1e10 + (0.1 - -2.25))) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    if (65535 >= 32767) {
                        if v2 {
                            for fi36 := 0; fi36 < 8; fi36++ {
                                switch ("a" + v33) {
                                case "":
                                    var v37 float64 = 1e10
                                    _ = v37
                                default:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                v32 = (v3 < 0)
                                var v38 uint = 0
                                _ = fi36
                                _ = v38
                            }
                            v33 = v33
                        }
                    } else {
                        for fi39 := 0; fi39 < 20; fi39++ {
                            v1 = v1
                            if true {
                                var v40 int = fi39
                                v1 += (v1 << ((137) & 63))
                                switch v33 {
                                case "\n\t":
                                    var v41 int16 = 0
                                    _ = v41
                                case "hello":
                                    v40 = (fi39 & fi39)
                                default:
                                    v42 := GenS9{F0: -1}
                                    _ = v42
                                }
                                _ = v40
                            } else {
                                if false {
                                    var v43 int = -9223372036854775808
                                    v33 = "foo bar"
                                    _ = v43
                                }
                                v2 = true
                            }
                            _ = fi39
                        }
                        var v44 int16 = 1
                        v1 = v1
                        _ = v44
                    }
                    v32 = v2
                    v45 := GenS10{F0: 42, F1: 88}
                    _ = v45
                }
                for fi46 := 0; fi46 < 15; fi46++ {
                    for fi47 := 0; fi47 < 19; fi47++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v48 uint = 2
                        _ = fi47
                        _ = v48
                    }
                    switch fi46 {
                    case 1:
                        v4 = append(v4, 1)
                    case 105:
                        var v49 uint = 2
                        _ = v49
                    default:
                        var v50 int64 = -387
                        _ = v50
                    }
                    _ = fi46
                }
            } else {
                if (false || false) {
                    var v51 float64 = 3.14159
                    _ = v51
                } else {
                    if v32 {
                        _ = len(v4)
                    }
                }
            }
            var v52 int64 = 9223372036854775807
            _ = v52
        } else {
            var v53 int = -9223372036854775808
            _ = v53
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v4
}
