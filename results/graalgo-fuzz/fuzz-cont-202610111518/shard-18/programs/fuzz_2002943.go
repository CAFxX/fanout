package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 int32; F1 uint; F2 bool }
type GenS5 struct { F0 uint }
func genFn0() int {
    switch -786 {
    case 32767:
        if true {
            v6 := []int{127, -2147483648, 128}
            _ = v6
        }
    default:
        var v7 uint16 = 0
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 int16 = -1
        for fi10 := 0; fi10 < 1; fi10++ {
            v11 := GenS1{F0: 141, F1: 127}
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    return -9223372036854775808
}
func genFn1() int32 {
    v12 := []int{-986, 9223372036854775807, -9223372036854775808}
    var v13 uint32 = 451
    var v14 uint16 = 0
    var v15 int64 = -81
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    return 2147483647
}

func main() {
    var v1 int = 806
    v2 := []int{v1, v1, 32768}
    v3 := make(map[int]int)
    var v4 float64 = 0.1
    v5 := []uint32{4294967295, 1}
    fmt.Println(v1, v4)
    switch "\n\t" {
    case "foo bar":
        v4 = 0.1
    case "\n\t":
        fmt.Println(v1)
    case "":
        var v16 uint8 = 176
        _ = v16
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        fmt.Println(v4, v1, v1)
        _ = fi17
    }
    var v18 float32 = 1.5
    var v19 int = v1
    v5 = append(v5, 65535)
    switch v1 {
    case 1:
        if (65535 != 2147483647) {
            var v20 uint64 = 2
            v1 -= (-(v19))
            v21 := GenS1{F0: 608, F1: 1}
            _ = v20
            _ = v21
        }
    case 4294967295:
        v22 := GenS1{F0: 18446744073709551615, F1: 128}
        _ = v22
    case 624:
        v4 = float64(v18)
    default:
        var v23 int32 = -2147483648
        _ = v23
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        var v25 float32 = ((-2.25 * float32(3.14159)) / v18)
        _ = fi24
        _ = v25
    }
    var v26 uint = 2
    fmt.Println(v19, v4, v1)
    v18 = (((float32(3.14159) + (v18 + (v18 + (((v18 * float32(v4)) / ((-2.25 / 3.14159) - float32((v4 / v4)))) + v18)))) * (v18 + (((((3.14159 * ((v18 - v18) + v18)) - ((v18 - float32(float64(((0.1 / 3.14159) + -2.25)))) + float32(1.5))) - 3.14159) * (v18 / float32(float64((1e10 * 0.1))))) / float32((v4 * (v4 * (((((v4 - float64(-2.25)) - float64(0.1)) + 1e10) - v4) * v4))))))) + v18)
    v27 := GenS2{F0: 663, F1: 2147483647, F2: 18446744073709551615}
    for fi28 := 0; fi28 < 9; fi28++ {
        if false {
            for fi29 := 0; fi29 < 20; fi29++ {
                if (false && (4294967295 >= 379)) {
                    if (-9223372036854775808 != -9223372036854775808) {
                        v26 = v26
                        v19 += (fi28 + genFn0())
                        v1 += genFn0()
                        if true {
                            v19 -= v1
                            var v30 string = ""
                            _ = v30
                        }
                        v31 := v3[v19]
                        v19 -= (-(fi28))
                        _ = v31
                    } else {
                        v2 = append(v2, (v19 | (v1 >> ((v26) & 63))))
                    }
                    v1 = v19
                } else {
                    switch ("hello" + "") {
                    case "":
                        var v32 uint = v26
                        _ = v32
                    }
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v33 int64 = 448
                _ = fi29
                _ = v33
            }
            for fi34 := 0; fi34 < 4; fi34++ {
                switch "foo bar" {
                case "a":
                    v35 := GenS1{F0: 4294967295, F1: 1}
                    _ = v35
                case "hello":
                    v36 := GenS3{F0: 128, F1: 349}
                    _ = v36
                case "\n\t":
                    v37 := GenS3{F0: 1, F1: 0}
                    _ = v37
                default:
                    v38 := GenS4{F0: -948, F1: 577, F2: (v18 != 0.1)}
                    _ = v38
                }
                switch ("foo bar" + "hello") {
                case "\n\t":
                    v39 := GenS2{F0: 0, F1: -2147483648, F2: 4294967295}
                    _ = v39
                }
                _ = fi34
            }
            switch 42 {
            case 913:
                v40 := GenS1{F0: 4294967295, F1: 128}
                _ = v40
            case 203:
                var v41 int8 = -111
                _ = v41
            case 32767:
                for fi42 := 0; fi42 < 20; fi42++ {
                    var v43 int32 = 847
                    v18 += v18
                    fmt.Println(v43, v18)
                    _ = fi42
                    _ = v43
                }
            default:
                var v44 uint = (v26 - 520)
                _ = v44
            }
        } else {
            if (!("foo bar" >= "a")) {
                var v45 float64 = v4
                fmt.Println(v18, v45, v18)
                v19 += 0
                _ = v45
            } else {
                for fi46 := 0; fi46 < 9; fi46++ {
                    if (true && ((!true) && false)) {
                        v4 -= 3.14159
                    } else {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        switch ("a" + ((("hello" + (("hello" + "hello") + ("" + (("hello" + "foo bar") + ("\n\t" + "hello"))))) + ("a" + "foo bar")) + "foo bar")) {
                        case "":
                            switch "hello" {
                            case "hello":
                                v1 = (fi46 * v1)
                            default:
                                v26 = v26
                            }
                        }
                        if len(v2) > 0 {
                            v2[0] = v19
                        }
                        if (596 != -1) {
                            var v47 int8 = 0
                            _ = len(v5)
                            if (!false) {
                                v47 = v47
                                var v48 uint8 = 42
                                var v49 uint16 = 1
                                _ = v48
                                _ = v49
                            } else {
                                v3[v1] = 0
                            }
                            _ = v47
                        } else {
                            switch v19 {
                            case -2147483648:
                                var v50 int64 = 1
                                _ = v50
                            default:
                                switch ("foo bar" + "a") {
                                case "":
                                    if false {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        if true {
                                            v51 := v3[fi46]
                                            v1 += 4294967295
                                            _ = v51
                                        }
                                    } else {
                                        var v52 uint8 = 178
                                        v26 += v26
                                        fmt.Println(v26)
                                        var v53 int = fi28
                                        v18 = (float32(v4) + v18)
                                        _ = v52
                                        _ = v53
                                    }
                                default:
                                    if true {
                                        fmt.Println(v26, fi46, v1)
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v54 int = -1
                                            _ = v54
                                        case "hello":
                                            v55 := GenS4{F0: 42, F1: (v26 ^ (v26 * v26)), F2: ((2147483647 > 0) && (("a" + ("" + "a")) > "foo bar"))}
                                            _ = v55
                                        case "a":
                                            v1 = fi28
                                        }
                                        var v56 uint32 = 0
                                        _ = v56
                                    } else {
                                        v3[(fi28 & (fi28 >> (((v26 << (((v26 + 42)) & 63))) & 63)))] = 42
                                    }
                                }
                            }
                        }
                    }
                    v18 = 3.14159
                    _ = fi46
                }
                v57 := GenS5{F0: (v26 / (((v26 ^ v26)) | 1))}
                switch (v19 | v19) {
                case 377:
                    switch "" {
                    case "hello":
                        for fi58 := 0; fi58 < 18; fi58++ {
                            var v59 int = 4294967295
                            _ = fi58
                            _ = v59
                        }
                    }
                case 9223372036854775807:
                    fmt.Println(v26, fi28)
                case -744:
                    fmt.Println(v19)
                default:
                    for fi60 := 0; fi60 < 10; fi60++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi60
                    }
                }
                _ = v57
            }
        }
        v19 -= (v1 & genFn0())
        v61 := GenS4{F0: -1, F1: (^(v26)), F2: (!true)}
        _ = fi28
        _ = v61
    }
    switch "foo bar" {
    case "a":
        for fi62 := 0; fi62 < 18; fi62++ {
            var v63 string = ""
            var v64 uint16 = 65535
            _ = fi62
            _ = v63
            _ = v64
        }
    default:
        switch v19 {
        case 4294967295:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case -2147483648:
            delete(v3, 2147483647)
        default:
            var v65 int8 = 26
            _ = v65
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
