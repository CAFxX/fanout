package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 string; F1 int32; F2 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 string }

func main() {
    var v1 uint32 = 2
    v2 := []int{42}
    v3 := []int{-2147483648, 95}
    var v4 float32 = 3.14159
    v5 := GenS1{F0: -1, F1: 2}
    v6 := GenS1{F0: 0, F1: 570}
    v7 := GenS1{F0: 1, F1: 18446744073709551615}
    if true {
        for fi8 := 0; fi8 < 1; fi8++ {
            if true {
                v1 -= v1
                var v9 uint = 1
                _ = v9
            }
            v10 := GenS1{F0: fi8, F1: 255}
            _ = fi8
            _ = v10
        }
        if false {
            if (-1 == -1) {
                fmt.Println(v4, v1)
            }
        } else {
            v4 += (v4 + float32(0.0))
        }
    } else {
        for fi11 := 0; fi11 < 7; fi11++ {
            v1 = (v1 >> ((uint(v1)) & 63))
            _ = fi11
        }
        switch "\n\t" {
        case "a":
            if (42 < 42) {
                switch "a" {
                case "foo bar":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "hello":
                    var v12 int32 = 451
                    _ = v12
                case "\n\t":
                    v3 = append(v3, 127)
                default:
                    switch 127 {
                    case 127:
                        if (!false) {
                            switch ("" + "foo bar") {
                            case "foo bar":
                                var v13 uint = 345
                                _ = v13
                            case "hello":
                                if (-2147483648 == 858) {
                                    v1 -= (v1 / (((v1 ^ (v1 / ((v1) | 1)))) | 1))
                                    if true {
                                        for fi14 := 0; fi14 < 15; fi14++ {
                                            for fi15 := 0; fi15 < 6; fi15++ {
                                                var v16 int32 = 1
                                                v2 = append(v2, 171)
                                                if true {
                                                    if (!(!true)) {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        v17 := GenS2{F0: (("" + ("a" + ("hello" + ((("a" + (((("\n\t" + (("" + "") + "\n\t")) + "\n\t") + "hello") + (("foo bar" + "") + ""))) + ("\n\t" + "")) + "hello")))) + ("\n\t" + (("\n\t" + (("\n\t" + "\n\t") + "a")) + ((((("hello" + "") + "\n\t") + "a") + "a") + "")))), F1: 0, F2: 42}
                                                        v1 -= (v1 / (((v1 & (v1 / (((v1 * v1)) | 1)))) | 1))
                                                        v18 := GenS1{F0: (fi15 << ((183) & 63)), F1: 374}
                                                        switch (fi15 - (fi15 - (fi15 ^ (fi14 + (fi14 & -9223372036854775808))))) {
                                                        case 32767:
                                                            var v19 int64 = -854
                                                            _ = v19
                                                        }
                                                        _ = v17
                                                        _ = v18
                                                    }
                                                    v16 = v16
                                                    if false {
                                                        v16 = v16
                                                        switch "" {
                                                        case "\n\t":
                                                            var v20 uint64 = 21
                                                            _ = v20
                                                        }
                                                        if true {
                                                            v4 = 3.14159
                                                        }
                                                    } else {
                                                        v21 := GenS3{F0: "foo bar"}
                                                        _ = v21
                                                    }
                                                } else {
                                                    fmt.Println(fi14, v1, v16)
                                                }
                                                _ = fi15
                                                _ = v16
                                            }
                                            _ = fi14
                                        }
                                    } else {
                                        fmt.Println(v1, v4)
                                    }
                                } else {
                                    v22 := GenS4{F0: (42 >= 368), F1: 42}
                                    var v23 int32 = -850
                                    _ = v22
                                    _ = v23
                                }
                            }
                            fmt.Println(v1, v4)
                        }
                    }
                }
                v4 += float32(1.5)
                v1 += 42
            }
        case "hello":
            if (false && (!(-9223372036854775808 == 32767))) {
                v2 = append(v2, 42)
                for fi24 := 0; fi24 < 18; fi24++ {
                    if false {
                        if len(v2) > 0 {
                            v2[0] = (fi24 * (fi24 % ((-9223372036854775808) | 1)))
                        }
                        var v25 int = (fi24 / ((0) | 1))
                        _ = v25
                    } else {
                        var v26 int32 = 2147483647
                        v1 -= v1
                        if false {
                            v27 := GenS5{F0: uint(v1), F1: v26}
                            _ = v27
                        }
                        _ = v26
                    }
                    v28 := GenS3{F0: (("a" + "\n\t") + "foo bar")}
                    v29 := GenS6{F0: "hello"}
                    _ = fi24
                    _ = v28
                    _ = v29
                }
                var v30 int8 = -66
                _ = v30
            } else {
                var v31 int16 = 1
                fmt.Println(v31, v4)
                for fi32 := 0; fi32 < 7; fi32++ {
                    for fi33 := 0; fi33 < 1; fi33++ {
                        if false {
                            fmt.Println(v4)
                        }
                        _ = fi33
                    }
                    v2 = append(v2, 859)
                    v4 = (((1.5 + (((v4 / (((((v4 - 1.5) / v4) / 0.1) + float32(-2.25)) * v4)) * ((v4 / v4) + (float32(0.0) / v4))) + (v4 / v4))) - ((0.0 + v4) - ((v4 - v4) * -2.25))) * float32(float64(v4)))
                    _ = fi32
                }
                _ = v31
            }
        }
        var v34 float32 = ((-2.25 * 1.5) + (-2.25 + v4))
        _ = v34
    }
    fmt.Println(v4, v4)
    var v35 int16 = 996
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
