package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 string; F1 uint8; F2 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int64; F1 int16 }
func genFn0(p0 uint32, p1 uint64) uint8 {
    fmt.Println(p0)
    p0 = p0
    if false {
        fmt.Println(p0, p1)
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := []int{(fi6 & -743), (fi6 & (fi6 << ((198) & 63)))}
            _ = fi6
            _ = v7
        }
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 uint64) uint16 {
    var v8 int8 = 100
    v8 = (v8 ^ 42)
    for fi9 := 0; fi9 < 15; fi9++ {
        v8 += (v8 + v8)
        p0 = p0
        v10 := []int{}
        _ = fi9
        _ = v10
    }
    p0 -= p0
    _ = p0
    _ = v8
    return 0
}

func main() {
    var v1 int32 = -369
    var v2 bool = true
    v3 := []string{"hello", "hello", "a", "\n\t"}
    var v4 bool = (v2 || (!v2))
    v5 := []string{("hello" + ("a" + "\n\t")), "hello", "", "foo bar"}
    var v11 int32 = (v1 ^ -1)
    fmt.Println(v2, v1, v4)
    v12 := GenS1{F0: (float32(1e10) == 0.1)}
    if true {
        v2 = v2
        if len(v3) > 0 {
            v3[0] = ("" + "")
        }
    } else {
        v13 := GenS2{F0: v11, F1: 3, F2: 2147483647}
        v4 = v2
        v2 = v4
        _ = v13
    }
    v14 := GenS3{F0: genFn1(42), F1: 42}
    v3 = append(v3, ("hello" + (("a" + "foo bar") + "")))
    v11 += (v11 & v11)
    v1 += v1
    for fi15 := 0; fi15 < 13; fi15++ {
        for fi16 := 0; fi16 < 12; fi16++ {
            if (101 <= 65535) {
                switch 963 {
                case -851:
                    v17 := GenS4{F0: ("" + "hello"), F1: 127, F2: 578}
                    _ = v17
                case 9223372036854775807:
                    _ = len(v3)
                case -2147483648:
                    if (v4 && v2) {
                        switch (fi16 / ((2147483647) | 1)) {
                        case -1:
                            v3 = append(v3, ("a" + (((("" + "foo bar") + (("\n\t" + "a") + (("foo bar" + "a") + ("" + "\n\t")))) + ("a" + "hello")) + ("a" + (((("" + ("\n\t" + "hello")) + "") + "foo bar") + (("" + ("\n\t" + "")) + "a"))))))
                        case -547:
                            if (v4 && false) {
                                v18 := GenS4{F0: "\n\t", F1: 24, F2: -477}
                                fmt.Println(v2, v1, v11)
                                switch fi15 {
                                case 978:
                                    v19 := GenS5{F0: 42}
                                    _ = v19
                                case 155:
                                    var v20 int16 = -32768
                                    _ = v20
                                case -322:
                                    var v21 uint8 = 127
                                    v2 = (((v21 >> ((65535) & 63)) <= v21) || (-1 != -1))
                                    if v4 {
                                        fmt.Println(v4)
                                        v4 = true
                                    } else {
                                        if v4 {
                                            switch fi16 {
                                            case -679:
                                                v4 = v4
                                            case 42:
                                                var v22 uint = 255
                                                _ = v22
                                            }
                                            switch "hello" {
                                            case "":
                                                if (!v4) {
                                                    v3 = append(v3, ("foo bar" + "foo bar"))
                                                } else {
                                                    if v2 {
                                                        switch ("\n\t" + "foo bar") {
                                                        case "\n\t":
                                                            var v23 uint64 = 18446744073709551615
                                                            fmt.Println(v11)
                                                            fmt.Println(v23, v4)
                                                            _ = v23
                                                        }
                                                        var v24 float64 = (1.5 + 0.0)
                                                        var v25 int32 = (v11 - -1)
                                                        _ = v24
                                                        _ = v25
                                                    } else {
                                                        v26 := GenS6{F0: 0, F1: 18446744073709551615}
                                                        if v2 {
                                                            fmt.Println(fi16)
                                                        } else {
                                                            if ("a" >= ("" + ("a" + "foo bar"))) {
                                                                var v27 uint64 = 281
                                                                var v28 int16 = -1
                                                                if v2 {
                                                                    _ = len(v3)
                                                                } else {
                                                                    v5 = append(v5, "")
                                                                    v29 := GenS3{F0: 2, F1: 572}
                                                                    var v30 string = "\n\t"
                                                                    _ = v29
                                                                    _ = v30
                                                                }
                                                                v21 -= (v21 * v21)
                                                                if len(v5) > 0 {
                                                                    v5[0] = ("" + ("hello" + ("\n\t" + "")))
                                                                }
                                                                _ = v27
                                                                _ = v28
                                                            } else {
                                                                var v31 int32 = (v11 * 155)
                                                                v11 = (v11 & (v11 << ((790) & 63)))
                                                                v5 = append(v5, "a")
                                                                _ = v31
                                                            }
                                                            if (3.14159 <= 3.14159) {
                                                                v4 = v2
                                                                fmt.Println(v4)
                                                                switch (("foo bar" + "") + ("" + (("\n\t" + (("foo bar" + "foo bar") + "a")) + ("hello" + ("hello" + ""))))) {
                                                                case "a":
                                                                    if (false && v4) {
                                                                        var v32 uint = 61
                                                                        switch 32767 {
                                                                        case -616:
                                                                            v4 = (1 > 628)
                                                                        case -824:
                                                                            var v33 float32 = 0.0
                                                                            _ = v33
                                                                        case 9223372036854775807:
                                                                            v5 = append(v5, "foo bar")
                                                                        default:
                                                                            var v34 bool = (true && (888 > 65535))
                                                                            _ = v34
                                                                        }
                                                                        _ = v32
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        switch -9223372036854775808 {
                                                        case -241:
                                                            v35 := GenS7{F0: "foo bar"}
                                                            _ = v35
                                                        case -2147483648:
                                                            v2 = (v11 < v1)
                                                        default:
                                                            fmt.Println(v2, fi16, fi16)
                                                        }
                                                        _ = v26
                                                    }
                                                    fmt.Println(v4)
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                }
                                            case "a":
                                                var v36 uint16 = 2
                                                _ = v36
                                            case "hello":
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            default:
                                                var v37 float64 = (((1e10 * (3.14159 - 0.1)) * (0.0 / 1e10)) - (0.1 / 0.1))
                                                _ = v37
                                            }
                                        }
                                        fmt.Println(fi16)
                                    }
                                    _ = v21
                                default:
                                    v3 = append(v3, "a")
                                }
                                _ = v18
                            } else {
                                v4 = false
                                if (32767 > 1) {
                                    var v38 int64 = 745
                                    _ = v38
                                } else {
                                    v4 = false
                                }
                            }
                        case 128:
                            var v39 int8 = -128
                            _ = v39
                        default:
                            var v40 bool = (910 > 728)
                            _ = v40
                        }
                        var v41 uint = 2
                        _ = len(v5)
                        _ = v41
                    } else {
                        v42 := GenS4{F0: ("a" + ("" + "\n\t")), F1: genFn0(624, 18446744073709551615), F2: -1}
                        switch "hello" {
                        case "a":
                            v11 += -2147483648
                        case "\n\t":
                            v43 := GenS8{F0: 2147483647}
                            _ = v43
                        case "foo bar":
                            if len(v3) > 0 {
                                v3[0] = (("" + "\n\t") + ("\n\t" + "a"))
                            }
                            v1 = v1
                            v3 = append(v3, ("hello" + ("" + "")))
                        default:
                            v11 = v1
                        }
                        _ = v42
                    }
                default:
                    if (((fi16 >> ((18446744073709551615) & 63)) <= fi16) && v2) {
                        var v44 int16 = 214
                        var v45 uint = 255
                        var v46 float32 = float32(float64((float32(-2.25) - float32(0.1))))
                        _ = v44
                        _ = v45
                        _ = v46
                    }
                }
                v4 = v2
            }
            _ = fi16
        }
        v2 = v4
        _ = fi15
    }
    var v47 int64 = 66
    v1 = (-(v11))
    v48 := GenS9{F0: (v47 | v47), F1: 648}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
}
