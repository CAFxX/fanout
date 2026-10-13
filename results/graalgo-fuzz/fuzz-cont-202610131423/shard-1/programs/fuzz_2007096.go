package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 int16 }
func genFn0() uint16 {
    v4 := GenS1{F0: 0, F1: ("" + ("" + ("foo bar" + ("\n\t" + "hello"))))}
    var v5 int8 = -128
    v6 := []int{-2147483648, 823, -579}
    _ = v4
    _ = v5
    _ = v6
    return 255
}

func main() {
    var v1 uint64 = 4294967295
    var v2 bool = ((!(((0.0 * (-2.25 - ((0.0 / (0.0 + 1e10)) / 1.0))) > -2.25) && (false || (true && true)))) || (6 != 255))
    var v3 uint16 = 1
    var v7 uint32 = 0
    fmt.Println(v3, v3, v1)
    if (v2 && (v1 == (^(v1)))) {
        var v8 float64 = 0.0
        _ = v8
    }
    var v9 int16 = -1
    fmt.Println(v9, v3, v9)
    v7 += 613
    if v2 {
        var v10 uint32 = v7
        _ = v10
    }
    v11 := []int{-2147483648, -555, -187, 42}
    switch ("a" + ("hello" + "hello")) {
    case "hello":
        v3 = v3
    case "":
        for fi12 := 0; fi12 < 3; fi12++ {
            if (v2 || v2) {
                if v2 {
                    if len(v11) > 0 {
                        v11[0] = 32767
                    }
                    v11 = append(v11, -9223372036854775808)
                } else {
                    v3 += (v3 + 65535)
                    v11 = append(v11, (fi12 & fi12))
                }
            }
            var v13 float64 = ((-2.25 / 1.0) / 1e10)
            _ = fi12
            _ = v13
        }
    case "\n\t":
        if false {
            if ((true && (!false)) && v2) {
                if ("" != ("foo bar" + "hello")) {
                    var v14 uint8 = 127
                    _ = v14
                } else {
                    if len(v11) > 0 {
                        v11[0] = 1
                    }
                    var v15 int8 = -125
                    var v16 int64 = 0
                    _ = v15
                    _ = v16
                }
                if (v2 && (!v2)) {
                    fmt.Println(v3, v9, v1)
                    switch "foo bar" {
                    case "a":
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                    case "foo bar":
                        if len(v11) > 0 {
                            v11[0] = -350
                        }
                    case "":
                        for fi17 := 0; fi17 < 13; fi17++ {
                            switch fi17 {
                            case -656:
                                var v18 float32 = float32(3.14159)
                                _ = v18
                            }
                            for fi19 := 0; fi19 < 8; fi19++ {
                                switch fi17 {
                                case 2147483647:
                                    var v20 int64 = 42
                                    _ = v20
                                case -2147483648:
                                    var v21 string = ((((((("hello" + ("" + ((("a" + "a") + "foo bar") + "\n\t"))) + "a") + "a") + (("hello" + "foo bar") + ((("\n\t" + "\n\t") + ("hello" + "\n\t")) + "\n\t"))) + ((("hello" + "foo bar") + "hello") + ("\n\t" + ""))) + ("a" + "foo bar")) + ((("foo bar" + "") + "hello") + (("hello" + ("foo bar" + ("hello" + "foo bar"))) + "foo bar")))
                                    _ = v21
                                case 4294967295:
                                    var v22 uint16 = (v3 >> ((4294967295) & 63))
                                    _ = v22
                                }
                                v23 := GenS1{F0: v1, F1: (("foo bar" + ("hello" + "")) + "\n\t")}
                                _ = fi19
                                _ = v23
                            }
                            _ = fi17
                        }
                    }
                } else {
                    if (100 >= 124) {
                        v24 := GenS2{F0: (v3 - (^(v3)))}
                        _ = v24
                    } else {
                        if (v2 || (((-498 > -1) || (1.5 == (1.5 - ((1e10 * 0.1) + 0.0)))) && (true && false))) {
                            v25 := GenS3{F0: 23, F1: (v9 ^ -1)}
                            var v26 float64 = float64(((1.5 - 1.5) / -2.25))
                            _ = v25
                            _ = v26
                        }
                    }
                    var v27 int16 = (v9 - v9)
                    _ = v27
                }
                switch -736 {
                case -1:
                    switch "a" {
                    case "\n\t":
                        fmt.Println(v3, v3, v2)
                    default:
                        if true {
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            v2 = ((((-2.25 / 1.5) >= 0.0) || (!true)) && false)
                        } else {
                            var v28 float64 = 3.14159
                            for fi29 := 0; fi29 < 12; fi29++ {
                                switch int(v1) {
                                case 9223372036854775807:
                                    fmt.Println(v3, v2, v3)
                                default:
                                    switch "hello" {
                                    case "foo bar":
                                        var v30 uint32 = v7
                                        _ = v30
                                    case "hello":
                                        for fi31 := 0; fi31 < 1; fi31++ {
                                            v28 += (0.0 - (0.0 / 3.14159))
                                            fmt.Println(v28, v28, v9)
                                            _ = fi31
                                        }
                                    default:
                                        _ = len(v11)
                                    }
                                }
                                fmt.Println(fi29, v28, v1)
                                _ = fi29
                            }
                            _ = v28
                        }
                    }
                case 4294967295:
                    _ = len(v11)
                }
            }
        } else {
            var v32 uint16 = uint16(v7)
            _ = v32
        }
    default:
        fmt.Println(v1, v2)
    }
    fmt.Println(v9, v7)
    var v33 int32 = 548
    var v34 uint8 = 255
    var v35 float64 = (1e10 + (0.0 + 0.1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v7
    _ = v9
}
