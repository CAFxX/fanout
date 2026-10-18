package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 int64 }

func main() {
    var v1 int8 = -128
    var v2 uint16 = 557
    var v3 int64 = 577
    v4 := []int{746, -95, -461}
    fmt.Println(v1, v2, v3)
    v1 += (v1 ^ 1)
    fmt.Println(v3)
    for fi5 := 0; fi5 < 18; fi5++ {
        fmt.Println(v1, v3, v1)
        v2 = v2
        if true {
            var v6 uint = 799
            _ = v6
        } else {
            _ = len(v4)
            if ((((0.1 + -2.25) * 1.5) / 1.0) == float32(-2.25)) {
                _ = len(v4)
                v3 = (v3 ^ 0)
                v7 := GenS1{F0: 0, F1: (v3 ^ (^(v3))), F2: "foo bar"}
                _ = v7
            }
        }
        _ = fi5
    }
    if false {
        var v8 uint64 = 1
        fmt.Println(v3, v1, v2)
        switch "\n\t" {
        case "\n\t":
            v9 := GenS1{F0: 1, F1: v3, F2: ("hello" + "hello")}
            _ = v9
        case "a":
            if true {
                if false {
                    if true {
                        switch 2147483647 {
                        case 1:
                            switch -9223372036854775808 {
                            case -2147483648:
                                if ((32768 != int(v3)) && (810 > 1)) {
                                    v1 -= (v1 & v1)
                                    v4 = append(v4, 42)
                                }
                            default:
                                v4 = append(v4, -9223372036854775808)
                            }
                        case 42:
                            switch 128 {
                            case -2147483648:
                                var v10 int8 = (v1 % ((1) | 1))
                                _ = v10
                            case 32767:
                                fmt.Println(v8)
                            case -1:
                                v3 -= 9223372036854775807
                            default:
                                for fi11 := 0; fi11 < 5; fi11++ {
                                    v4 = append(v4, (fi11 + fi11))
                                    _ = fi11
                                }
                            }
                        }
                    }
                    v1 = v1
                    var v12 int8 = (v1 >> ((1) & 63))
                    _ = v12
                } else {
                    switch 127 {
                    case -9223372036854775808:
                        fmt.Println(v3, v3, v2)
                    default:
                        if (!(false && (!(("" <= "\n\t") && (!true))))) {
                            fmt.Println(v1, v8, v1)
                            v4 = append(v4, -642)
                        }
                    }
                    if false {
                        for fi13 := 0; fi13 < 17; fi13++ {
                            var v14 int8 = v1
                            v1 = (-(v14))
                            _ = len(v4)
                            _ = fi13
                            _ = v14
                        }
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(v2, v3, v8)
                        case "a":
                            v15 := GenS1{F0: 2147483647, F1: v3, F2: ("hello" + "hello")}
                            _ = v15
                        }
                    } else {
                        for fi16 := 0; fi16 < 10; fi16++ {
                            var v17 uint64 = (v8 - (^(v8)))
                            v8 -= 1
                            _ = fi16
                            _ = v17
                        }
                        v18 := GenS2{F0: -32768}
                        switch "a" {
                        case "a":
                            fmt.Println(v2)
                        case "":
                            var v19 int = 5
                            _ = v19
                        case "hello":
                            if (v8 <= (v8 * v8)) {
                                switch "foo bar" {
                                case "":
                                    v2 = (^(v2))
                                }
                                switch 0 {
                                case 238:
                                    v2 += v2
                                case 0:
                                    _ = len(v4)
                                case 1:
                                    var v20 int16 = 1000
                                    _ = v20
                                }
                                v2 = v2
                            }
                        }
                        _ = v18
                    }
                    for fi21 := 0; fi21 < 15; fi21++ {
                        if len(v4) > 0 {
                            v4[0] = (fi21 / ((146) | 1))
                        }
                        v22 := GenS2{F0: 42}
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi21
                        _ = v22
                    }
                }
                v23 := GenS2{F0: 1000}
                v2 -= (v2 & (v2 >> ((uint(v2)) & 63)))
                _ = v23
            } else {
                for fi24 := 0; fi24 < 18; fi24++ {
                    if len(v4) > 0 {
                        v4[0] = (-(fi24))
                    }
                    v4 = append(v4, (^(fi24)))
                    var v25 int16 = 42
                    _ = fi24
                    _ = v25
                }
            }
        case "foo bar":
            v2 -= (^(v2))
        }
        _ = v8
    }
    switch ("foo bar" + (("" + "") + (("\n\t" + "hello") + "a"))) {
    case "":
        for fi26 := 0; fi26 < 3; fi26++ {
            for fi27 := 0; fi27 < 17; fi27++ {
                v28 := GenS3{F0: 2, F1: 0}
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
    case "a":
        fmt.Println(v2)
    default:
        for fi29 := 0; fi29 < 3; fi29++ {
            switch "hello" {
            case "\n\t":
                for fi30 := 0; fi30 < 5; fi30++ {
                    v1 += v1
                    if false {
                        fmt.Println(fi30, fi30)
                        v4 = append(v4, fi29)
                    }
                    var v31 float32 = ((float32((1.5 * 0.0)) + 0.0) * (1.5 * (0.1 * 0.0)))
                    _ = fi30
                    _ = v31
                }
            default:
                var v32 string = "hello"
                _ = v32
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v33 := GenS2{F0: 42}
            _ = fi29
            _ = v33
        }
    }
    v34 := GenS2{F0: 1000}
    v4 = append(v4, 1)
    v35 := GenS2{F0: -32768}
    v36 := GenS3{F0: 4294967295, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
