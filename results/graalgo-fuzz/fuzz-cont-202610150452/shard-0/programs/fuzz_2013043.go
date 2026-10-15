package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 int32; F1 int32; F2 uint16 }

func main() {
    var v1 int16 = -790
    var v2 uint8 = 255
    var v3 int64 = -9223372036854775808
    var v4 bool = true
    v5 := []int{4294967295}
    var v6 uint8 = (^(v2))
    v6 = (v6 + (v2 / ((v6) | 1)))
    fmt.Println(v6, v1, v3)
    v4 = v4
    switch "foo bar" {
    case "":
        switch 42 {
        case 2147483647:
            switch "hello" {
            case "a":
                v1 += -32768
            default:
                switch 0 {
                case 1:
                    v1 = (v1 % (((v1 & (v1 + -32768))) | 1))
                case -2147483648:
                    v2 = (v6 - (^(v6)))
                }
            }
        case -419:
            var v7 float64 = 1e10
            _ = v7
        }
    case "a":
        v2 += (v6 * 243)
    case "hello":
        v8 := GenS1{F0: 2}
        _ = v8
    }
    if false {
        switch 0 {
        case 32767:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case 32768:
            v6 += (v6 / ((v6) | 1))
        case -2147483648:
            var v9 int16 = (v1 * v1)
            _ = v9
        default:
            switch "hello" {
            case "":
                v3 += (v3 - (v3 | 0))
            case "hello":
                v2 = (v6 + v2)
            case "a":
                if (!v4) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v6 -= v6
                    v10 := GenS1{F0: 18446744073709551615}
                    _ = v10
                } else {
                    if (int32(v3) == 1) {
                        var v11 int64 = (^(v3))
                        _ = v11
                    }
                }
            default:
                if len(v5) > 0 {
                    v5[0] = 2147483647
                }
            }
        }
        v6 = (v6 + v2)
        v6 = (v2 + (v6 ^ v2))
    } else {
        for fi12 := 0; fi12 < 4; fi12++ {
            if len(v5) > 0 {
                v5[0] = 2147483647
            }
            v4 = (((!(!(100 == -128))) || (v2 < v2)) && true)
            _ = fi12
        }
        if (v1 < v1) {
            v4 = v4
            for fi13 := 0; fi13 < 7; fi13++ {
                switch fi13 {
                case 128:
                    v14 := GenS2{F0: 42, F1: 65535}
                    _ = v14
                case 127:
                    switch int(v3) {
                    case 32768:
                        if v4 {
                            var v15 uint16 = 2
                            switch (fi13 & fi13) {
                            case -2147483648:
                                _ = len(v5)
                            case -1:
                                v15 = v15
                            }
                            _ = v15
                        } else {
                            switch ("a" + "") {
                            case "a":
                                if (v4 && v4) {
                                    v16 := GenS1{F0: 2}
                                    _ = v16
                                }
                            case "hello":
                                switch -9223372036854775808 {
                                case 32768:
                                    if len(v5) > 0 {
                                        v5[0] = (fi13 << ((255) & 63))
                                    }
                                case -507:
                                    fmt.Println(v3, v4)
                                case 0:
                                    fmt.Println(v3, v3)
                                default:
                                    fmt.Println(v6, v3, v1)
                                }
                            case "foo bar":
                                v5 = append(v5, (fi13 / ((32767) | 1)))
                            default:
                                if v4 {
                                    var v17 int16 = -32768
                                    for fi18 := 0; fi18 < 6; fi18++ {
                                        var v19 string = ("hello" + ((("\n\t" + "\n\t") + "foo bar") + "a"))
                                        v20 := GenS2{F0: 65535, F1: 255}
                                        _ = fi18
                                        _ = v19
                                        _ = v20
                                    }
                                    _ = v17
                                }
                            }
                        }
                    case 0:
                        for fi21 := 0; fi21 < 18; fi21++ {
                            v3 = 974
                            var v22 int8 = 1
                            v23 := GenS3{F0: v6, F1: 42}
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                    }
                case 32768:
                    fmt.Println(v4)
                default:
                    v4 = (v4 || v4)
                }
                _ = fi13
            }
            for fi24 := 0; fi24 < 2; fi24++ {
                v5 = append(v5, (fi24 + fi24))
                _ = fi24
            }
        } else {
            v25 := GenS4{F0: int32(v6), F1: 1, F2: 744}
            _ = v25
        }
    }
    v2 -= 240
    var v26 uint = 1
    v27 := GenS4{F0: -923, F1: -1, F2: 255}
    if false {
        var v28 bool = (3.14159 > -2.25)
        for fi29 := 0; fi29 < 6; fi29++ {
            v30 := GenS2{F0: 267, F1: 32767}
            _ = fi29
            _ = v30
        }
        _ = v28
    } else {
        fmt.Println(v26, v2)
        v6 -= (v6 >> ((v26) & 63))
    }
    v3 = -909
    for fi31 := 0; fi31 < 12; fi31++ {
        var v32 string = ""
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        if v4 {
            switch fi33 {
            case 32768:
                var v34 uint8 = v2
                _ = v34
            }
        } else {
            var v35 bool = v4
            _ = v35
        }
        var v36 bool = (!false)
        _ = fi33
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
