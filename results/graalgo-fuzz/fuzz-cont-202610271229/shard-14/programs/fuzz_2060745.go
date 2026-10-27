package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 2147483648
    v3 := []int{32767, 128, 4294967295}
    var v4 uint16 = 0
    var v5 float32 = (1.5 * 0.0)
    var v6 int16 = 32767
    var v7 int16 = v6
    if (!(!false)) {
        if false {
            if ((!(42 <= 18446744073709551615)) && ((^(v6)) <= v6)) {
                v2 = 2147483648
                switch "foo bar" {
                case "foo bar":
                    v3 = append(v3, 127)
                }
                v8 := GenS1{F0: 0, F1: 100}
                _ = v8
            }
            var v9 int8 = -7
            _ = len(v1)
            _ = v9
        } else {
            var v10 uint = 222
            v10 = v10
            var v11 int64 = 0
            _ = v10
            _ = v11
        }
    } else {
        var v12 uint64 = 939
        v6 = (^(v7))
        switch "" {
        case "a":
            for fi13 := 0; fi13 < 19; fi13++ {
                v2 = v2
                v14 := GenS1{F0: 128, F1: 100}
                _ = fi13
                _ = v14
            }
        case "hello":
            fmt.Println(v5)
        case "\n\t":
            v2 = (v2 + (v2 - v2))
        default:
            v7 = 466
        }
        _ = v12
    }
    v6 = (v7 / (((v6 * 32767)) | 1))
    v2 = (v2 >> ((4294967295) & 63))
    var v15 float32 = v5
    v2 += (v2 * (v2 - 42))
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 uint32 = 1
        _ = fi16
        _ = v17
    }
    var v18 string = "a"
    v4 = (v4 ^ (v4 - v4))
    for fi19 := 0; fi19 < 5; fi19++ {
        v15 = 1e10
        switch ("" + v18) {
        case "":
            v20 := GenS1{F0: 203, F1: 127}
            _ = v20
        }
        _ = fi19
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        fmt.Println(v7, v4, v2)
        switch -9223372036854775808 {
        case 400:
            var v22 uint8 = 128
            _ = v22
        case 32767:
            fmt.Println(v4, v2, v6)
        case 390:
            v23 := GenS1{F0: 109, F1: -128}
            _ = v23
        }
        _ = fi21
    }
    _ = len(v1)
    v1[-2147483648] = 2147483647
    switch v18 {
    case "":
        switch (v18 + (v18 + (("a" + (v18 + "foo bar")) + (v18 + v18)))) {
        case "":
            if true {
                var v24 int64 = 0
                switch (("" + ("a" + "hello")) + v18) {
                case "":
                    if ((v24 / ((v24) | 1)) >= (v24 | v24)) {
                        v15 = (v15 * 3.14159)
                        switch v18 {
                        case "a":
                            switch v18 {
                            case "\n\t":
                                switch (v18 + v18) {
                                case "foo bar":
                                    fmt.Println(v24)
                                }
                            case "foo bar":
                                v4 = (v4 - (v4 ^ (v4 + v4)))
                            case "a":
                                v1[665] = v2
                            default:
                                fmt.Println(v18, v18)
                            }
                        case "":
                            v24 = (v24 + v24)
                        default:
                            var v25 uint8 = 2
                            _ = v25
                        }
                        for fi26 := 0; fi26 < 1; fi26++ {
                            _ = len(v3)
                            for fi27 := 0; fi27 < 2; fi27++ {
                                var v28 int = fi27
                                _ = fi27
                                _ = v28
                            }
                            v18 = (v18 + (v18 + (v18 + v18)))
                            _ = fi26
                        }
                    }
                case "hello":
                    if (!true) {
                        _ = len(v1)
                    }
                default:
                    v1[1] = v2
                }
                if (3.14159 > 1e10) {
                    fmt.Println(v18)
                    for fi29 := 0; fi29 < 13; fi29++ {
                        for fi30 := 0; fi30 < 20; fi30++ {
                            var v31 float32 = (((-2.25 - ((((v5 + ((v15 * ((v5 - (v15 / v5)) - 0.0)) - (((1e10 / v5) - float32(0.1)) * v5))) - v5) + (0.0 * 0.0)) - v15)) / v15) + v5)
                            _ = fi30
                            _ = v31
                        }
                        v7 = v7
                        var v32 int64 = -764
                        _ = fi29
                        _ = v32
                    }
                    v33 := GenS2{F0: 122, F1: 18446744073709551615}
                    _ = v33
                } else {
                    v6 += 506
                    var v34 uint32 = v2
                    fmt.Println(v7)
                    _ = v34
                }
                _ = v24
            } else {
                v35 := GenS1{F0: 127, F1: 100}
                v18 = "\n\t"
                _ = v35
            }
        case "hello":
            v15 = ((3.14159 + 1e10) - v5)
        }
    }
    v5 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
