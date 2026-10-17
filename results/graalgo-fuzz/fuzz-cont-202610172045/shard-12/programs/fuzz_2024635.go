package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint32; F1 int8; F2 string }

func main() {
    var v1 float64 = (((0.0 / 3.14159) + 0.1) + 1e10)
    var v2 uint32 = 0
    var v3 uint16 = 783
    v4 := make(map[int]string)
    v5 := []int{32768, 9223372036854775807, -466, 602}
    fmt.Println(v1)
    for fi6 := 0; fi6 < 9; fi6++ {
        _ = len(v4)
        v5 = append(v5, (fi6 - fi6))
        v2 = v2
        _ = fi6
    }
    fmt.Println(v1, v1)
    var v7 string = "\n\t"
    v8 := GenS1{F0: v3}
    v4[115] = (v7 + ("" + ((((v7 + v7) + ("hello" + "\n\t")) + ("\n\t" + "a")) + v7)))
    v9 := GenS1{F0: v3}
    for fi10 := 0; fi10 < 1; fi10++ {
        v5 = append(v5, (fi10 << ((65535) & 63)))
        v11 := v4[(fi10 % (((fi10 * (^(fi10)))) | 1))]
        v12 := GenS2{F0: 42}
        _ = fi10
        _ = v11
        _ = v12
    }
    v3 -= v3
    fmt.Println(v7, v1)
    v13 := GenS3{F0: 581, F1: -128, F2: 0}
    v2 = v2
    if (-9223372036854775808 == -1) {
        v2 -= v2
        var v14 float32 = 0.0
        _ = v14
    } else {
        fmt.Println(v7)
        v1 += v1
        v15 := GenS2{F0: 100}
        _ = v15
    }
    switch "a" {
    case "foo bar":
        switch "\n\t" {
        case "foo bar":
            switch 32768 {
            case -1:
                v16 := GenS1{F0: 65535}
                _ = v16
            case 32768:
                switch 9223372036854775807 {
                case -1:
                    switch v7 {
                    case "":
                        var v17 int16 = 0
                        _ = v17
                    case "hello":
                        if ((18446744073709551615 != 18446744073709551615) && true) {
                            for fi18 := 0; fi18 < 20; fi18++ {
                                if ("foo bar" != (v7 + v7)) {
                                    v7 = v7
                                    var v19 int64 = -1
                                    if ((0.1 / float64((float32((v1 / v1)) - float32(0.0)))) > (v1 / (v1 / v1))) {
                                        var v20 uint16 = v3
                                        switch "hello" {
                                        case "a":
                                            v21 := GenS2{F0: -69}
                                            _ = v21
                                        case "foo bar":
                                            for fi22 := 0; fi22 < 5; fi22++ {
                                                var v23 string = v7
                                                v24 := GenS1{F0: 0}
                                                v20 = (v20 ^ v20)
                                                _ = fi22
                                                _ = v23
                                                _ = v24
                                            }
                                        case "\n\t":
                                            fmt.Println(fi18)
                                        }
                                        _ = v20
                                    }
                                    _ = v19
                                } else {
                                    for fi25 := 0; fi25 < 2; fi25++ {
                                        var v26 uint = 65535
                                        switch "\n\t" {
                                        case "hello":
                                            v26 += v26
                                        case "foo bar":
                                            _ = len(v5)
                                            v7 = (v7 + v7)
                                        case "\n\t":
                                            v7 = (v7 + "")
                                        default:
                                            var v27 int64 = -421
                                            _ = v27
                                        }
                                        if false {
                                            _ = len(v5)
                                        }
                                        _ = fi25
                                        _ = v26
                                    }
                                }
                                if true {
                                    v28 := v4[fi18]
                                    _ = v28
                                } else {
                                    v29 := GenS4{F0: 625, F1: -128}
                                    v2 -= (v2 / ((2) | 1))
                                    fmt.Println(v1, v2)
                                    _ = v29
                                }
                                if true {
                                    v1 = ((v1 / float64(3.14159)) * -2.25)
                                } else {
                                    switch fi18 {
                                    case 46:
                                        var v30 float32 = (3.14159 * 0.1)
                                        _ = v30
                                    case -94:
                                        switch "hello" {
                                        case "a":
                                            v2 -= 734
                                        case "hello":
                                            for fi31 := 0; fi31 < 5; fi31++ {
                                                var v32 float64 = (1e10 / (float64((0.0 - 1.5)) - v1))
                                                _ = fi31
                                                _ = v32
                                            }
                                        default:
                                            fmt.Println(v1, v7, v7)
                                        }
                                    default:
                                        var v33 uint64 = 2
                                        _ = v33
                                    }
                                    var v34 int64 = 42
                                    fmt.Println(v7)
                                    _ = v34
                                }
                                _ = fi18
                            }
                            v35 := GenS4{F0: 42, F1: -63}
                            v36 := GenS4{F0: (v3 >> ((585) & 63)), F1: -128}
                            _ = v35
                            _ = v36
                        } else {
                            for fi37 := 0; fi37 < 2; fi37++ {
                                v7 = "hello"
                                v38 := GenS4{F0: v3, F1: -1}
                                _ = fi37
                                _ = v38
                            }
                            v39 := GenS2{F0: 1}
                            _ = v39
                        }
                    }
                default:
                    v1 = -2.25
                }
            default:
                for fi40 := 0; fi40 < 2; fi40++ {
                    v3 = v3
                    v1 = ((v1 + float64(1e10)) + float64(1.5))
                    _ = fi40
                }
            }
        }
    }
    v7 = v7
    if (false && (-760 == 32767)) {
        v41 := GenS2{F0: -97}
        _ = v41
    } else {
        for fi42 := 0; fi42 < 9; fi42++ {
            v1 = 0.1
            var v43 int8 = 1
            _ = fi42
            _ = v43
        }
        v7 = ("a" + v7)
        switch int(v2) {
        case 1:
            if ((((true || false) && (v3 <= (v3 + 2))) || ((^(v3)) > (v3 - 32767))) || (55 == 0)) {
                var v44 uint16 = v3
                _ = v44
            } else {
                var v45 uint8 = 128
                for fi46 := 0; fi46 < 11; fi46++ {
                    for fi47 := 0; fi47 < 9; fi47++ {
                        v48 := GenS5{F0: 2147483647}
                        var v49 uint32 = (v2 * (v2 & (v2 - (v2 & (v2 | v2)))))
                        _ = fi47
                        _ = v48
                        _ = v49
                    }
                    fmt.Println(fi46)
                    _ = fi46
                }
                fmt.Println(v1)
                _ = v45
            }
        }
    }
    v50 := GenS6{F0: (v2 % ((84) | 1)), F1: -1, F2: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v7
    _ = v8
    _ = v9
}
