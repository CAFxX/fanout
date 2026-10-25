package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 uint32; F2 int }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int8; F1 uint }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int64; F1 int32; F2 bool }
type GenS10 struct { F0 uint8; F1 bool }
func genFn0(p0 uint16, p1 int32) uint8 {
    p0 -= 65535
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: 4294967295, F1: -9223372036854775808, F2: 902}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 593, F1: -1, F2: 1}
    _ = p0
    _ = p1
    _ = v9
    return 255
}

func main() {
    var v1 bool = (false && ((true || true) || false))
    var v2 bool = true
    var v3 uint = 42
    var v4 uint32 = 979
    var v5 float32 = -2.25
    v6 := []string{}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println(v2)
    v10 := GenS1{F0: (v4 + (v4 + (v4 & v4))), F1: int64(v3), F2: 65535}
    v11 := GenS2{F0: 42}
    if len(v6) > 0 {
        v6[0] = ("hello" + "hello")
    }
    var v12 string = (("hello" + ("a" + "a")) + ("hello" + "foo bar"))
    var v13 int64 = -11
    switch 42 {
    case 127:
        for fi14 := 0; fi14 < 4; fi14++ {
            switch v12 {
            case "hello":
                if v2 {
                    if v2 {
                        if v1 {
                            for fi15 := 0; fi15 < 12; fi15++ {
                                v16 := GenS1{F0: v4, F1: 0, F2: 4294967295}
                                var v17 bool = true
                                var v18 uint = v3
                                _ = fi15
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                            if (v1 && false) {
                                var v19 string = "foo bar"
                                v12 = (((v19 + v12) + ((v12 + v19) + ("a" + "a"))) + ((("hello" + "foo bar") + "hello") + (v19 + (v19 + (v19 + v19)))))
                                _ = v19
                            } else {
                                v2 = (32767 > -142)
                                for fi20 := 0; fi20 < 2; fi20++ {
                                    v21 := GenS3{F0: 699}
                                    v6 = append(v6, "a")
                                    var v22 uint32 = v4
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                            }
                        }
                        v12 = ("a" + "a")
                    } else {
                        v1 = ((((!(!v1)) || (v12 == "\n\t")) || false) || (157 == 0))
                        if v2 {
                            v2 = v1
                            v23 := GenS4{F0: 836, F1: 120, F2: 4294967295}
                            _ = v23
                        }
                        for fi24 := 0; fi24 < 14; fi24++ {
                            if (((v5 + -2.25) * (v5 - v5)) != (v5 - ((((v5 + (((((((0.0 - (float32((1e10 / (3.14159 + (0.0 / 0.1)))) * (0.0 + 0.1))) / (float32(3.14159) - (0.1 + (((v5 + ((float32(0.1) / ((3.14159 * (float32(((float64((v5 + (float32(0.1) / 1.5))) * ((0.0 * (1e10 + 0.0)) + float64(-2.25))) * -2.25)) + -2.25)) - v5)) * v5)) + float32(((-2.25 + (1e10 / ((-2.25 / 0.1) / 1.0))) * 0.1))) - v5)))) / v5) + v5) / v5) - (v5 + v5)) - (v5 * 0.1))) * 1.5) * (0.1 * v5)) + v5))) {
                                v25 := GenS5{F0: 2147483647, F1: (v4 << (((v3 >> ((v3) & 63))) & 63))}
                                v26 := GenS6{F0: v4}
                                _ = v25
                                _ = v26
                            } else {
                                var v27 int = 42
                                switch fi24 {
                                case 1:
                                    v12 = "a"
                                case 2147483647:
                                    v28 := GenS7{F0: -128, F1: v3}
                                    _ = v28
                                case -1:
                                    fmt.Println(v3, v3)
                                default:
                                    var v29 float64 = 0.1
                                    _ = v29
                                }
                                _ = v27
                            }
                            _ = fi24
                        }
                    }
                    for fi30 := 0; fi30 < 8; fi30++ {
                        v12 = v12
                        v5 = (v5 - v5)
                        fmt.Println(v12, v1, v12)
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = fi30
                    }
                } else {
                    if true {
                        if len(v6) > 0 {
                            v6[0] = ""
                        }
                        for fi31 := 0; fi31 < 16; fi31++ {
                            v13 = (v13 >> ((v3) & 63))
                            v5 -= (0.0 + v5)
                            v1 = false
                            _ = fi31
                        }
                    } else {
                        v6 = append(v6, "\n\t")
                    }
                }
            default:
                switch (fi14 % (((^(fi14))) | 1)) {
                case 32768:
                    var v32 uint16 = 718
                    _ = v32
                case 0:
                    if (1 == 42) {
                        if len(v6) > 0 {
                            v6[0] = v12
                        }
                        if (!v1) {
                            v5 += (v5 + v5)
                            fmt.Println(v4, v1, v4)
                            if true {
                                v4 = (v4 + (v4 ^ 2147483648))
                                var v33 float64 = 0.0
                                _ = v33
                            }
                        } else {
                            for fi34 := 0; fi34 < 6; fi34++ {
                                var v35 uint8 = 104
                                v36 := GenS8{F0: 430}
                                _ = fi34
                                _ = v35
                                _ = v36
                            }
                        }
                        for fi37 := 0; fi37 < 11; fi37++ {
                            var v38 float32 = v5
                            _ = fi37
                            _ = v38
                        }
                    }
                }
            }
            _ = fi14
        }
    case -9223372036854775808:
        v6 = append(v6, v12)
    case 42:
        v13 += (v13 - (v13 | v13))
    default:
        var v39 uint32 = 602
        _ = v39
    }
    for fi40 := 0; fi40 < 11; fi40++ {
        if false {
            if (false && v1) {
                v41 := GenS9{F0: 782, F1: 376, F2: true}
                _ = len(v6)
                var v42 float64 = (-2.25 - ((0.1 / -2.25) - -2.25))
                _ = v41
                _ = v42
            }
            var v43 int32 = 42
            _ = v43
        }
        v4 = v4
        for fi44 := 0; fi44 < 10; fi44++ {
            v1 = (-32768 > -32768)
            _ = len(v6)
            if false {
                v45 := GenS7{F0: 100, F1: 18446744073709551615}
                var v46 float32 = v5
                _ = v45
                _ = v46
            }
            _ = fi44
        }
        _ = fi40
    }
    var v47 bool = v1
    var v48 float32 = 3.14159
    _ = len(v6)
    var v49 int8 = 1
    var v50 uint32 = (v4 - 0)
    v48 = 0.1
    var v51 int8 = v49
    fmt.Println(v12, v2, v3)
    v52 := GenS10{F0: 43, F1: (-2147483648 < 0)}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
}
