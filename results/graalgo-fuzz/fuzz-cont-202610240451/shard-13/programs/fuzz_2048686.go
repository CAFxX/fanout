package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 string; F2 uint16 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint16; F1 bool; F2 int16 }
type GenS9 struct { F0 int64; F1 uint64; F2 int32 }
type GenS10 struct { F0 int; F1 uint8; F2 uint32 }

func main() {
    var v1 uint64 = 4294967295
    var v2 uint16 = 1
    v3 := make(map[string]uint32)
    var v4 uint8 = 128
    fmt.Println(v4)
    for fi5 := 0; fi5 < 16; fi5++ {
        for fi6 := 0; fi6 < 4; fi6++ {
            fmt.Println(v2, fi5, fi5)
            _ = fi6
        }
        _ = fi5
    }
    fmt.Println(v1)
    fmt.Println(v2, v4, v4)
    for fi7 := 0; fi7 < 15; fi7++ {
        delete(v3, ("" + (("hello" + "foo bar") + "hello")))
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := GenS1{F0: "foo bar", F1: 42}
            v10 := GenS1{F0: (("foo bar" + ((((("\n\t" + ("\n\t" + ("foo bar" + ("hello" + "foo bar")))) + ("foo bar" + "")) + "foo bar") + "hello") + ("a" + "hello"))) + "a"), F1: (^(v1))}
            if ((0.0 + 0.1) != 0.0) {
                v11 := GenS1{F0: "\n\t", F1: (v1 + v1)}
                var v12 uint64 = 4294967295
                v12 += v12
                _ = v11
                _ = v12
            } else {
                v4 = (v4 & (v4 << ((4294967295) & 63)))
                v13 := GenS2{F0: 18446744073709551615}
                _ = v13
            }
            switch "\n\t" {
            case "":
                var v14 int64 = -962
                _ = v14
            case "foo bar":
                v15 := GenS3{F0: 100, F1: ("\n\t" + "hello"), F2: v2}
                _ = v15
            case "a":
                if (false && false) {
                    switch ((("foo bar" + ("" + ("a" + "a"))) + "") + "foo bar") {
                    case "":
                        var v16 string = "\n\t"
                        _ = v16
                    case "a":
                        var v17 float32 = ((1.5 / 1.0) + 1.5)
                        _ = v17
                    case "hello":
                        if ((0.0 == (1.5 / (1e10 * 0.1))) || true) {
                            v18 := []int{fi7}
                            _ = v18
                        }
                    }
                    v19 := GenS4{F0: v1, F1: (v1 % ((v1) | 1))}
                    if (575 <= -892) {
                        if (42 == 42) {
                            v20 := GenS4{F0: (v1 + uint64(v4)), F1: (v1 >> ((255) & 63))}
                            _ = v20
                        } else {
                            switch ((("foo bar" + ("" + "\n\t")) + "foo bar") + (("" + "\n\t") + "\n\t")) {
                            case "a":
                                if ((0.1 - -2.25) > float64(((1.5 - 3.14159) * (3.14159 - ((3.14159 + (-2.25 / 3.14159)) - -2.25))))) {
                                    var v21 int8 = 100
                                    v1 += (^(v1))
                                    var v22 int64 = 9223372036854775807
                                    v23 := GenS1{F0: (("a" + "") + ""), F1: (v1 * (v1 - (v1 | (v1 << ((1) & 63)))))}
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                }
                            case "foo bar":
                                var v24 uint8 = v4
                                _ = v24
                            case "":
                                v25 := []int{}
                                _ = v25
                            default:
                                if false {
                                    if (-1 < 2147483647) {
                                        var v26 uint16 = uint16(v1)
                                        _ = v26
                                    }
                                }
                            }
                            v27 := GenS5{F0: 1, F1: 0}
                            v28 := GenS1{F0: "a", F1: v1}
                            var v29 float32 = ((1e10 / 0.1) - 1.5)
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                    }
                    _ = v19
                }
            default:
                var v30 int16 = 1
                v30 += (v30 & v30)
                v31 := []int{}
                _ = v30
                _ = v31
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        fmt.Println(v4, v4)
        _ = fi7
    }
    v32 := []int{2147483647, 9223372036854775807, int(v1)}
    if (!(true && true)) {
        v4 += 127
        if (2147483647 > 65535) {
            v33 := GenS6{F0: uint16(v1)}
            v2 = (v2 / (((v2 / (((v2 % ((v2) | 1))) | 1))) | 1))
            v1 += v1
            _ = v33
        } else {
            var v34 int64 = -9223372036854775808
            if (!true) {
                v4 = (v4 ^ (v4 | 42))
                fmt.Println(v2, v1)
            } else {
                var v35 int64 = v34
                if (-16 <= 79) {
                    v4 += 64
                    if true {
                        var v36 uint16 = v2
                        _ = v36
                    } else {
                        var v37 int32 = 2147483647
                        _ = v37
                    }
                    v35 = (v34 << ((0) & 63))
                } else {
                    v38 := GenS5{F0: (v34 & (v34 & (v35 >> ((579) & 63)))), F1: 266}
                    fmt.Println(v34, v1, v35)
                    v1 = (^(v1))
                    _ = v38
                }
                var v39 uint = 2
                _ = v35
                _ = v39
            }
            if false {
                if len(v32) > 0 {
                    v32[0] = 32767
                }
                if (!false) {
                    delete(v3, ((("foo bar" + "hello") + ("hello" + ("\n\t" + "a"))) + ("hello" + ("" + ""))))
                    v40 := GenS7{F0: "a"}
                    _ = v40
                }
                v34 = (^(v34))
            }
            _ = v34
        }
        v41 := GenS8{F0: v2, F1: ((!(2 < 2)) || true), F2: 32767}
        _ = v41
    }
    v2 = v2
    v4 = 127
    if true {
        var v42 float32 = (float32(0.1) + (((((1.5 / 0.1) - 0.1) - (-2.25 / 1.5)) * 1e10) - 1.5))
        _ = v42
    } else {
        _ = len(v32)
        var v43 bool = (128 > -754)
        _ = v43
    }
    switch "hello" {
    case "a":
        for fi44 := 0; fi44 < 9; fi44++ {
            switch fi44 {
            case 93:
                v2 -= (v2 + v2)
            case 9223372036854775807:
                fmt.Println(v2, v2, v1)
            default:
                var v45 float64 = 3.14159
                _ = v45
            }
            _ = fi44
        }
    case "hello":
        var v46 int32 = 0
        _ = v46
    default:
        var v47 bool = false
        _ = v47
    }
    for fi48 := 0; fi48 < 8; fi48++ {
        for fi49 := 0; fi49 < 18; fi49++ {
            v50 := GenS9{F0: 111, F1: (v1 << ((830) & 63)), F2: 2147483647}
            _ = fi49
            _ = v50
        }
        _ = fi48
    }
    var v51 uint = 1
    if ((false || false) || false) {
        v2 = v2
        v52 := GenS10{F0: 1, F1: (v4 % (((^(v4))) | 1)), F2: 927}
        _ = v52
    } else {
        for fi53 := 0; fi53 < 6; fi53++ {
            for fi54 := 0; fi54 < 12; fi54++ {
                v1 -= (v1 ^ 229)
                _ = len(v32)
                v55 := GenS7{F0: ("foo bar" + "\n\t")}
                v2 += (v2 * v2)
                _ = fi54
                _ = v55
            }
            for fi56 := 0; fi56 < 15; fi56++ {
                if false {
                    delete(v3, "foo bar")
                } else {
                    var v57 int8 = 127
                    _ = v57
                }
                v32 = append(v32, fi53)
                _ = fi56
            }
            _ = fi53
        }
        v4 = (v4 >> ((v51) & 63))
        fmt.Println(v1, v2)
    }
    v4 = (v4 - v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v51
}
