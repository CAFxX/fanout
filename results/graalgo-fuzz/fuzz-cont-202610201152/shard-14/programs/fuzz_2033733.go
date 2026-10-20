package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int; F1 uint32; F2 uint }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 string }

func main() {
    var v1 uint = 293
    var v2 int8 = 1
    var v3 int16 = 32767
    var v4 float64 = 3.14159
    var v5 uint32 = 75
    v6 := []int{-2147483648}
    if false {
        switch (("hello" + ("hello" + ("a" + ""))) + "\n\t") {
        case "foo bar":
            var v7 bool = (float32(0.0) <= float32(((-2.25 + -2.25) / 1.0)))
            _ = v7
        }
    }
    var v8 int8 = v2
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 float32 = (((((0.0 / ((0.1 / (1e10 + (0.0 + -2.25))) * -2.25)) + 3.14159) / 1.5) + 0.0) + -2.25)
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        for fi12 := 0; fi12 < 14; fi12++ {
            v5 = 585
            _ = fi12
        }
        v3 = (v3 << ((v1) & 63))
        _ = fi11
    }
    var v13 float32 = float32(v4)
    v13 = v13
    v4 = v4
    v8 += (v2 >> ((v1) & 63))
    if (!(!(1 != 42))) {
        v6 = append(v6, 32767)
        v3 += (^(v3))
    } else {
        switch 32768 {
        case 127:
            v3 = v3
        case 42:
            switch "a" {
            case "hello":
                for fi14 := 0; fi14 < 16; fi14++ {
                    var v15 float32 = v13
                    v16 := GenS1{F0: 42, F1: 352}
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            default:
                fmt.Println(v3, v4, v8)
            }
        }
    }
    v17 := GenS1{F0: v1, F1: 477}
    if (18446744073709551615 == 4294967295) {
        switch "" {
        case "\n\t":
            for fi18 := 0; fi18 < 7; fi18++ {
                var v19 float32 = v13
                _ = fi18
                _ = v19
            }
        default:
            v1 = v1
        }
        switch "hello" {
        case "a":
            if false {
                var v20 string = "a"
                switch -65 {
                case -9223372036854775808:
                    v20 = (v20 + v20)
                }
                _ = v20
            } else {
                for fi21 := 0; fi21 < 8; fi21++ {
                    fmt.Println(v2, v5, v13)
                    _ = fi21
                }
            }
        }
        fmt.Println(v3)
    }
    if true {
        v22 := GenS1{F0: v1, F1: 18446744073709551615}
        _ = v22
    } else {
        for fi23 := 0; fi23 < 13; fi23++ {
            if false {
                fmt.Println(v1)
                for fi24 := 0; fi24 < 19; fi24++ {
                    switch (fi24 % (((fi23 + fi23)) | 1)) {
                    case -9223372036854775808:
                        switch "hello" {
                        case "":
                            fmt.Println(v1, v3)
                        case "\n\t":
                            var v25 uint64 = 454
                            _ = v25
                        case "a":
                            if len(v6) > 0 {
                                v6[0] = (fi23 + fi24)
                            }
                        }
                    }
                    _ = fi24
                }
            } else {
                v26 := GenS2{F0: fi23, F1: v5, F2: v1}
                for fi27 := 0; fi27 < 13; fi27++ {
                    switch (^(fi23)) {
                    case -2147483648:
                        v8 -= (v2 << ((840) & 63))
                    case 127:
                        switch -1 {
                        case -9223372036854775808:
                            v2 = int8(fi27)
                        case -1:
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            if false {
                                if (42 < 1) {
                                    if len(v6) > 0 {
                                        v6[0] = fi23
                                    }
                                    v28 := GenS3{F0: v2, F1: (true || ((v3 - (v3 % ((0) | 1))) <= (v3 & (v3 | (v3 * (^(v3)))))))}
                                    v1 -= v1
                                    _ = v28
                                } else {
                                    v2 = v2
                                }
                                v13 -= v13
                                v29 := GenS3{F0: (^(v8)), F1: ((!true) && true)}
                                if (false && (v1 <= (v1 | 759))) {
                                    var v30 uint8 = 1
                                    _ = v30
                                }
                                v4 += (v4 + 1.5)
                                _ = v29
                            } else {
                                var v31 int = -1
                                var v32 string = ("foo bar" + "a")
                                switch fi27 {
                                case 1:
                                    var v33 int32 = 0
                                    _ = v33
                                }
                                _ = v31
                                _ = v32
                            }
                            v5 -= v5
                        }
                    case 884:
                        var v34 uint8 = 0
                        switch fi23 {
                        case -483:
                            fmt.Println(v5)
                        default:
                            var v35 uint8 = uint8(v1)
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            var v36 int32 = 0
                            _ = v35
                            _ = v36
                        }
                        _ = v34
                    default:
                        v1 += uint(v5)
                    }
                    _ = fi27
                }
                if true {
                    v37 := GenS4{F0: (v8 >> (((v1 / (((v1 + v1)) | 1))) & 63)), F1: "foo bar"}
                    v4 += ((((v4 / 1e10) + float64(1.5)) / v4) + 0.0)
                    _ = v37
                } else {
                    v1 = v1
                    fmt.Println(v5, v13, v2)
                }
                _ = v26
            }
            v6 = append(v6, fi23)
            _ = fi23
        }
    }
    if (true && (((v4 / 0.1) == 0.1) || true)) {
        var v38 uint8 = 1
        _ = v38
    }
    switch 777 {
    case 4294967295:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    default:
        for fi39 := 0; fi39 < 15; fi39++ {
            fmt.Println(v5, v5)
            v40 := GenS5{F0: (("a" + ("foo bar" + (("foo bar" + "hello") + ("hello" + "a")))) + "hello")}
            v5 = (v5 - v5)
            _ = fi39
            _ = v40
        }
    }
    v5 += v5
    if (false || (!false)) {
        v3 = v3
        var v41 float64 = float64((-2.25 + v13))
        _ = v41
    }
    var v42 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v8
}
