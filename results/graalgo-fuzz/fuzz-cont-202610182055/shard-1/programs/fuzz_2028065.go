package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int64; F1 int64; F2 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 string; F1 uint16 }
type GenS6 struct { F0 uint; F1 uint; F2 string }
func genFn0() uint64 {
    for fi4 := 0; fi4 < 17; fi4++ {
        v5 := GenS1{F0: 2147483647, F1: 228}
        _ = fi4
        _ = v5
    }
    if (100 == 42) {
        var v6 bool = true
        _ = v6
    }
    return 42
}
func genFn1(p0 bool) uint {
    if false {
        switch 32768 {
        case 4294967295:
            fmt.Println(p0)
        case -389:
            var v7 uint16 = 255
            _ = v7
        case 862:
            p0 = ((p0 || ((705 < 77) || (65535 < 348))) && p0)
        default:
            if (-681 < 42) {
                var v8 int64 = 1
                _ = v8
            }
        }
    }
    p0 = false
    p0 = false
    _ = p0
    return 255
}

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = -128
    var v3 int8 = (v2 % ((100) | 1))
    switch "\n\t" {
    case "foo bar":
        v2 += v3
    case "":
        if true {
            v2 -= (v2 | v2)
            if (true && (v3 >= 0)) {
                v9 := []int{2147483647}
                switch 0 {
                case 1:
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                case 42:
                    switch (((("hello" + "hello") + "hello") + ("foo bar" + "")) + "a") {
                    case "a":
                        v10 := GenS1{F0: 42, F1: 127}
                        _ = v10
                    case "hello":
                        var v11 uint16 = 1
                        _ = v11
                    default:
                        switch "\n\t" {
                        case "foo bar":
                            var v12 uint8 = 206
                            _ = v12
                        case "":
                            v2 = v2
                        default:
                            for fi13 := 0; fi13 < 2; fi13++ {
                                fmt.Println(v3, v2)
                                _ = fi13
                            }
                        }
                    }
                case -2147483648:
                    _ = len(v1)
                }
                _ = v9
            } else {
                for fi14 := 0; fi14 < 6; fi14++ {
                    for fi15 := 0; fi15 < 14; fi15++ {
                        fmt.Println(fi14, v2, v2)
                        v2 = (v3 << ((0) & 63))
                        _ = fi15
                    }
                    switch ("" + "hello") {
                    case "hello":
                        v1[("\n\t" + ("" + ""))] = (-1 <= -1)
                    case "":
                        v2 = v3
                    }
                    v16 := GenS1{F0: 10, F1: 255}
                    _ = fi14
                    _ = v16
                }
            }
        }
    case "a":
        v1["foo bar"] = (false && true)
    default:
        v17 := v1["foo bar"]
        _ = v17
    }
    var v18 int8 = (v3 - (^(v3)))
    fmt.Println(v2, v3, v2)
    fmt.Println(v3)
    v19 := GenS2{F0: 42, F1: 978, F2: 1}
    if (0 != -1) {
        v20 := GenS1{F0: 65535, F1: 198}
        _ = v20
    }
    v21 := []int{-1, -2147483648, 0, -2147483648}
    v22 := GenS3{F0: -424}
    var v23 bool = false
    v18 = (-(v18))
    v23 = (2147483647 == 2147483647)
    var v24 int8 = (v18 % ((v2) | 1))
    if true {
        v25 := GenS4{F0: (v18 & 42), F1: -32768}
        _ = v25
    } else {
        var v26 int16 = 0
        _ = v26
    }
    if (1 == 0) {
        v24 -= (v24 >> ((65535) & 63))
        if (42 >= 255) {
            for fi27 := 0; fi27 < 12; fi27++ {
                for fi28 := 0; fi28 < 20; fi28++ {
                    fmt.Println(v2)
                    _ = len(v1)
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                    _ = fi28
                }
                for fi29 := 0; fi29 < 3; fi29++ {
                    v30 := GenS5{F0: "", F1: 1}
                    _ = fi29
                    _ = v30
                }
                _ = fi27
            }
            for fi31 := 0; fi31 < 16; fi31++ {
                for fi32 := 0; fi32 < 10; fi32++ {
                    if ("\n\t" > (("hello" + "\n\t") + ("a" + ("foo bar" + "\n\t")))) {
                        fmt.Println(v18)
                        fmt.Println(fi32)
                        v3 += (v24 % ((v2) | 1))
                    } else {
                        delete(v1, ("hello" + ("hello" + "hello")))
                        var v33 int64 = 712
                        var v34 int16 = 0
                        _ = v33
                        _ = v34
                    }
                    fmt.Println(v3, v18, v24)
                    fmt.Println(v18, fi32)
                    if (((fi31 == (^(fi32))) || false) || (!((!(4294967295 <= genFn0())) && v23))) {
                        switch (fi32 << ((1) & 63)) {
                        case -2147483648:
                            v3 += (-(v3))
                        case 4294967295:
                            v35 := GenS2{F0: -194, F1: 0, F2: genFn0()}
                            _ = v35
                        }
                        switch "foo bar" {
                        case "":
                            v36 := GenS3{F0: -9223372036854775808}
                            _ = v36
                        case "a":
                            v2 = (v18 & v3)
                        case "foo bar":
                            var v37 int32 = 1
                            _ = v37
                        default:
                            var v38 int64 = -1
                            _ = v38
                        }
                    } else {
                        var v39 string = "foo bar"
                        _ = v39
                    }
                    _ = fi32
                }
                var v40 string = ("\n\t" + ((("" + "\n\t") + "\n\t") + ("a" + "foo bar")))
                fmt.Println(fi31)
                _ = fi31
                _ = v40
            }
            v41 := GenS2{F0: 0, F1: -155, F2: 2}
            _ = v41
        } else {
            var v42 float32 = float32(((((3.14159 * -2.25) * float64(((float32(3.14159) / 1.5) * (1e10 + float32(0.1))))) / ((1.5 - 0.1) - (float64(((0.1 - 0.0) + 1.5)) - -2.25))) * -2.25))
            v42 = v42
            _ = v42
        }
    } else {
        if (float32((1e10 * 1.5)) <= (1.5 * 1.5)) {
            var v43 int32 = 1
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
            _ = v43
        } else {
            for fi44 := 0; fi44 < 20; fi44++ {
                for fi45 := 0; fi45 < 19; fi45++ {
                    v46 := GenS1{F0: 1, F1: 255}
                    v24 = (-(v3))
                    _ = fi45
                    _ = v46
                }
                var v47 uint16 = 2
                if (1.5 <= (0.1 - -2.25)) {
                    if (124 > 244) {
                        v18 = (v24 & 0)
                    }
                    fmt.Println(v18)
                    v2 -= v3
                } else {
                    if true {
                        if true {
                            v3 += v2
                        } else {
                            switch "" {
                            case "":
                                if false {
                                    switch "hello" {
                                    case "foo bar":
                                        v48 := GenS2{F0: 9223372036854775807, F1: -1, F2: 42}
                                        _ = v48
                                    default:
                                        var v49 float64 = float64(0.0)
                                        _ = v49
                                    }
                                }
                            case "hello":
                                v3 += v2
                            default:
                                v18 = (v2 ^ (-(v18)))
                            }
                        }
                    } else {
                        v3 += -1
                        var v50 uint32 = 2147483648
                        v50 = v50
                        _ = v50
                    }
                    var v51 float64 = ((1e10 - 0.1) / (1.5 / (-2.25 / 3.14159)))
                    var v52 int16 = 1
                    _ = v51
                    _ = v52
                }
                _ = fi44
                _ = v47
            }
            if (!true) {
                switch "hello" {
                case "foo bar":
                    _ = len(v21)
                case "hello":
                    if (!v23) {
                        for fi53 := 0; fi53 < 1; fi53++ {
                            var v54 uint16 = 42
                            _ = fi53
                            _ = v54
                        }
                        v2 -= v3
                    }
                case "":
                    v55 := v1[("a" + "\n\t")]
                    _ = v55
                default:
                    v24 += v24
                }
                if true {
                    var v56 int16 = 42
                    _ = v56
                } else {
                    v57 := GenS6{F0: 18446744073709551615, F1: 706, F2: "hello"}
                    _ = v57
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
