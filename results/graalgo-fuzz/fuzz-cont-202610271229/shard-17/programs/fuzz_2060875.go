package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int64; F1 string; F2 int32 }
type GenS4 struct { F0 uint32; F1 uint32; F2 bool }
type GenS5 struct { F0 bool; F1 uint; F2 string }

func main() {
    var v1 uint32 = 42
    v2 := []string{}
    v3 := make(map[string]int)
    var v4 uint64 = 571
    var v5 uint32 = (v1 << ((18446744073709551615) & 63))
    v6 := make(map[uint32]uint32)
    var v7 int64 = 9223372036854775807
    fmt.Println(v4, v1)
    switch 4294967295 {
    case -224:
        v7 += v7
    case 56:
        v8 := GenS1{F0: 2, F1: ((0.0 * float32(0.1)) >= float32(0.0)), F2: uint(v5)}
        _ = v8
    case 938:
        for fi9 := 0; fi9 < 6; fi9++ {
            v7 = v7
            v10 := GenS2{F0: v4, F1: 1, F2: 42}
            _ = fi9
            _ = v10
        }
    }
    fmt.Println(v5, v4)
    v5 -= v1
    var v11 int8 = 127
    if true {
        v12 := GenS1{F0: 0, F1: false, F2: 18446744073709551615}
        _ = v12
    } else {
        if ((true && true) || (false && false)) {
            var v13 uint8 = 147
            if ((^(v5)) <= (v1 - v1)) {
                v11 -= (v11 | 104)
            } else {
                var v14 float64 = ((float64(float32(3.14159)) * 1.5) * 3.14159)
                if true {
                    if false {
                        v15 := GenS2{F0: v4, F1: v13, F2: (v13 / ((v13) | 1))}
                        v16 := GenS2{F0: (v4 ^ (v4 - 2)), F1: uint8(v1), F2: (^(v13))}
                        v14 = v14
                        _ = v15
                        _ = v16
                    } else {
                        if true {
                            var v17 int = -591
                            v7 += (v7 - v7)
                            _ = len(v6)
                            _ = v17
                        } else {
                            v18 := GenS1{F0: v13, F1: (826 >= 1), F2: uint(v4)}
                            _ = v18
                        }
                        v7 = v7
                    }
                } else {
                    switch ("foo bar" + "foo bar") {
                    case "a":
                        v1 -= 498
                    default:
                        v7 += (v7 ^ v7)
                    }
                    if ("hello" == (("\n\t" + "a") + "")) {
                        v14 = (v14 - v14)
                        v19 := GenS2{F0: (v4 | v4), F1: (^(v13)), F2: (v13 % ((v13) | 1))}
                        _ = v19
                    }
                    v20 := GenS3{F0: (v7 * (v7 / (((v7 >> ((42) & 63))) | 1))), F1: "a", F2: 2147483647}
                    _ = v20
                }
                for fi21 := 0; fi21 < 10; fi21++ {
                    switch (^(fi21)) {
                    case 2147483647:
                        v14 += v14
                    case 128:
                        v7 = (v7 * v7)
                    case -9223372036854775808:
                        v13 = v13
                    }
                    _ = len(v6)
                    switch fi21 {
                    case -2147483648:
                        v4 += (v4 << ((803) & 63))
                    }
                    _ = fi21
                }
                _ = v14
            }
            switch ("" + ("foo bar" + "a")) {
            case "foo bar":
                switch -9223372036854775808 {
                case -2147483648:
                    if (true || (32767 != 42)) {
                        if false {
                            for fi22 := 0; fi22 < 6; fi22++ {
                                switch fi22 {
                                case 32768:
                                    v6[v1] = 2147483647
                                default:
                                    var v23 int64 = v7
                                    _ = v23
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi22
                            }
                            v24 := GenS1{F0: v13, F1: (true && false), F2: 4294967295}
                            _ = v24
                        }
                    }
                case -9223372036854775808:
                    _ = len(v6)
                case -1:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
            _ = v13
        } else {
            v11 = (v11 + (v11 & (v11 >> ((1) & 63))))
        }
        if true {
            v25 := GenS3{F0: v7, F1: ("" + ""), F2: 42}
            switch "\n\t" {
            case "\n\t":
                for fi26 := 0; fi26 < 7; fi26++ {
                    v6[(v1 - v5)] = v5
                    _ = fi26
                }
            case "":
                v4 += (v4 / ((18446744073709551615) | 1))
            }
            v1 = 522
            _ = v25
        } else {
            v6[(v1 >> ((42) & 63))] = 636
        }
    }
    v5 = 717
    switch "foo bar" {
    case "hello":
        v7 = (v7 * v7)
    default:
        _ = len(v2)
    }
    if false {
        v11 = -36
        if (true || (461 >= 1)) {
            v27 := GenS3{F0: v7, F1: "a", F2: 1}
            for fi28 := 0; fi28 < 15; fi28++ {
                switch ("" + ("hello" + ("foo bar" + (("a" + "a") + "hello")))) {
                case "foo bar":
                    v7 += (v7 * (v7 << ((4294967295) & 63)))
                case "":
                    for fi29 := 0; fi29 < 5; fi29++ {
                        v30 := GenS4{F0: (v1 / ((v5) | 1)), F1: (v5 | v1), F2: (true && (2 == 380))}
                        v1 -= (v5 & 470)
                        fmt.Println(fi29)
                        _ = fi29
                        _ = v30
                    }
                case "hello":
                    for fi31 := 0; fi31 < 7; fi31++ {
                        var v32 int32 = 0
                        switch "" {
                        case "foo bar":
                            var v33 int64 = (v7 % (((v7 / (((v7 >> ((4294967295) & 63))) | 1))) | 1))
                            _ = v33
                        case "a":
                            fmt.Println(fi31, fi28, v1)
                        }
                        v34 := GenS2{F0: 785, F1: 0, F2: 181}
                        _ = fi31
                        _ = v32
                        _ = v34
                    }
                }
                switch int(v11) {
                case 128:
                    v5 = v5
                case -187:
                    v35 := GenS2{F0: (v4 ^ 224), F1: 2, F2: uint8(v4)}
                    _ = v35
                default:
                    v36 := GenS2{F0: (^(v4)), F1: 127, F2: 255}
                    _ = v36
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi28
            }
            _ = v27
        }
        _ = len(v2)
    } else {
        if false {
            v37 := GenS3{F0: (v7 & (v7 * v7)), F1: "foo bar", F2: 1}
            v1 += v1
            _ = v37
        } else {
            for fi38 := 0; fi38 < 2; fi38++ {
                for fi39 := 0; fi39 < 2; fi39++ {
                    _ = len(v6)
                    v2 = append(v2, ((("\n\t" + ("" + "\n\t")) + "foo bar") + (("a" + "hello") + ("hello" + (("a" + "foo bar") + "a")))))
                    _ = fi39
                }
                var v40 uint16 = 65535
                _ = fi38
                _ = v40
            }
            v5 += (v1 ^ v1)
            fmt.Println(v1)
        }
        delete(v6, 712)
        if (((v11 <= v11) && false) && true) {
            v41 := GenS1{F0: 67, F1: true, F2: 65535}
            v42 := GenS5{F0: ((1000 >= 42) || (92 != 1)), F1: 18446744073709551615, F2: (("foo bar" + "\n\t") + "")}
            v7 = (-(v7))
            _ = v41
            _ = v42
        } else {
            var v43 uint32 = 551
            v11 -= (v11 & v11)
            _ = len(v2)
            _ = v43
        }
    }
    switch "\n\t" {
    case "foo bar":
        var v44 string = ""
        _ = v44
    case "\n\t":
        v11 -= -128
    default:
        var v45 string = ("foo bar" + ("\n\t" + "hello"))
        _ = v45
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
