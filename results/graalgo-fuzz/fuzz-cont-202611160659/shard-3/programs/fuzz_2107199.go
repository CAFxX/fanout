package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 string }
type GenS2 struct { F0 int64; F1 bool; F2 bool }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64; F1 int8; F2 int }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int }

func main() {
    var v1 int = 42
    v2 := make(map[string]bool)
    v3 := make(map[int]bool)
    var v4 int = 583
    var v5 int8 = 0
    v6 := []uint32{0, 42, 0}
    v4 = v1
    switch ("a" + "hello") {
    case "":
        fmt.Println(v5)
    default:
        switch "" {
        case "a":
            var v7 uint8 = 1
            _ = v7
        default:
            var v8 string = "foo bar"
            _ = v8
        }
    }
    if true {
        v9 := GenS1{F0: 384, F1: 753, F2: "foo bar"}
        v1 = (v4 * (-(v4)))
        switch (("" + "foo bar") + "foo bar") {
        case "":
            v4 = (^(v1))
        case "\n\t":
            v10 := GenS2{F0: -9223372036854775808, F1: ((!(-2.25 >= 0.1)) && (9223372036854775807 >= 42)), F2: false}
            _ = v10
        case "hello":
            switch "a" {
            case "foo bar":
                v1 = (v1 + v4)
            default:
                if (-2147483648 != -2147483648) {
                    var v11 bool = false
                    if (((0.0 + (1.5 - ((0.1 / 1.5) * (1e10 * 1e10)))) > (0.0 - float32(1.5))) && v11) {
                        if true {
                            v12 := v3[(v4 % ((v1) | 1))]
                            _ = v12
                        } else {
                            switch (v1 << ((2) & 63)) {
                            case 9223372036854775807:
                                v2["a"] = (0 == 18446744073709551615)
                            case -1:
                                fmt.Println(v5)
                            case -462:
                                var v13 uint16 = 304
                                _ = v13
                            }
                        }
                        if (450 > 163) {
                            switch ("foo bar" + "\n\t") {
                            case "a":
                                if v11 {
                                    switch ("a" + "") {
                                    case "":
                                        v14 := GenS3{F0: ((-1 <= 1) || true), F1: 255}
                                        _ = v14
                                    }
                                    var v15 int = v1
                                    _ = v15
                                } else {
                                    v1 = v1
                                    for fi16 := 0; fi16 < 17; fi16++ {
                                        var v17 float32 = float32((1e10 - ((3.14159 / 1.5) * float64(0.0))))
                                        v17 += v17
                                        var v18 int16 = 42
                                        _ = fi16
                                        _ = v17
                                        _ = v18
                                    }
                                    var v19 int16 = -227
                                    _ = v19
                                }
                            default:
                                for fi20 := 0; fi20 < 18; fi20++ {
                                    for fi21 := 0; fi21 < 4; fi21++ {
                                        v5 = -2
                                        fmt.Println(fi20, fi21)
                                        if (!true) {
                                            v22 := GenS4{F0: (v5 % (((^(v5))) | 1))}
                                            v6 = append(v6, 2)
                                            _ = v22
                                        }
                                        _ = fi21
                                    }
                                    v23 := GenS5{F0: -467, F1: v5, F2: fi20}
                                    _ = fi20
                                    _ = v23
                                }
                            }
                            v24 := GenS5{F0: 1, F1: v5, F2: v4}
                            var v25 float64 = ((float64(0.0) * 0.1) + 1e10)
                            _ = v24
                            _ = v25
                        }
                    } else {
                        if len(v6) > 0 {
                            v6[0] = 2147483648
                        }
                        v26 := GenS6{F0: 42}
                        _ = v26
                    }
                    v6 = append(v6, 4294967295)
                    _ = v11
                }
            }
        default:
            fmt.Println(v4, v1)
        }
        _ = v9
    }
    var v27 int8 = 100
    var v28 int32 = -1
    v29 := GenS3{F0: (!(42 <= -1)), F1: 128}
    for fi30 := 0; fi30 < 13; fi30++ {
        v31 := GenS7{F0: v1}
        _ = fi30
        _ = v31
    }
    var v32 uint64 = 828
    v2[("hello" + ("foo bar" + (("\n\t" + "a") + "hello")))] = false
    var v33 int64 = 42
    v34 := v3[-9223372036854775808]
    if false {
        switch "a" {
        case "":
            var v35 bool = ((!("\n\t" > "hello")) && v34)
            _ = v35
        default:
            if (0 < 255) {
                var v36 float64 = 3.14159
                var v37 int32 = 1
                _ = v36
                _ = v37
            } else {
                fmt.Println(v33, v5)
            }
        }
        switch "a" {
        case "\n\t":
            if v34 {
                v1 -= 1
            } else {
                for fi38 := 0; fi38 < 9; fi38++ {
                    var v39 uint8 = 245
                    _ = fi38
                    _ = v39
                }
                fmt.Println(v1, v4, v34)
            }
        case "hello":
            if true {
                fmt.Println(v33, v33, v1)
                v5 -= (v27 << ((600) & 63))
            }
        case "":
            v27 += (v5 + (v5 - -49))
        }
    } else {
        switch (v1 ^ (v4 | -804)) {
        case 317:
            var v40 uint8 = 221
            _ = v40
        case -949:
            var v41 int = 887
            _ = v41
        }
    }
    fmt.Println(v33, v27, v34)
    for fi42 := 0; fi42 < 12; fi42++ {
        if false {
            fmt.Println(v27, v27)
            var v43 uint8 = 2
            _ = v43
        }
        v2["hello"] = v34
        _ = fi42
    }
    var v44 uint8 = 128
    var v45 int = (v4 ^ (v1 + 0))
    if (!(v34 || v34)) {
        fmt.Println(v45, v44, v28)
        v5 = int8(v44)
        v4 = v1
    } else {
        v5 = v27
        v44 += (v44 >> ((173) & 63))
    }
    v46 := GenS1{F0: 2, F1: 1, F2: "foo bar"}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
