package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 bool; F1 uint64; F2 int }

func main() {
    v1 := []string{}
    var v2 int16 = -32768
    var v3 uint64 = 0
    var v4 float32 = (0.1 + (3.14159 / 3.14159))
    switch -2147483648 {
    case -9223372036854775808:
        for fi5 := 0; fi5 < 2; fi5++ {
            var v6 float32 = v4
            fmt.Println(v3, fi5)
            _ = fi5
            _ = v6
        }
    case 1:
        if ((!((v2 | 42) <= v2)) && (("a" + "hello") <= "\n\t")) {
            v7 := GenS1{F0: 2, F1: 0}
            v8 := GenS1{F0: 0, F1: 302}
            var v9 float64 = ((1.5 / 0.1) / 3.14159)
            _ = v7
            _ = v8
            _ = v9
        } else {
            for fi10 := 0; fi10 < 7; fi10++ {
                for fi11 := 0; fi11 < 4; fi11++ {
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                    _ = fi11
                }
                if (-1 != -2147483648) {
                    switch "\n\t" {
                    case "hello":
                        switch "hello" {
                        case "\n\t":
                            fmt.Println(v2, v4, v2)
                        default:
                            var v12 int16 = (v2 & (v2 + -1))
                            _ = v12
                        }
                    case "\n\t":
                        if (v4 <= 0.0) {
                            for fi13 := 0; fi13 < 7; fi13++ {
                                fmt.Println(fi13)
                                switch ("" + "foo bar") {
                                case "hello":
                                    v3 = v3
                                case "\n\t":
                                    if true {
                                        v2 -= v2
                                    }
                                case "":
                                    switch "a" {
                                    case "":
                                        v4 = (v4 - (0.1 * 1.5))
                                    case "\n\t":
                                        v2 = v2
                                    default:
                                        var v14 int = (fi13 & fi10)
                                        _ = v14
                                    }
                                default:
                                    var v15 uint = 430
                                    _ = v15
                                }
                                if ((2147483648 > uint32(v3)) && ("\n\t" <= (("a" + "hello") + "a"))) {
                                    var v16 int64 = 384
                                    if len(v1) > 0 {
                                        v1[0] = "a"
                                    }
                                    _ = v16
                                } else {
                                    var v17 int = -24
                                    v3 -= v3
                                    v18 := GenS1{F0: 127, F1: uint(v3)}
                                    _ = v17
                                    _ = v18
                                }
                                _ = fi13
                            }
                            var v19 string = "hello"
                            _ = v19
                        }
                    case "":
                        var v20 int = fi10
                        _ = v20
                    default:
                        v1 = append(v1, "")
                    }
                    if len(v1) > 0 {
                        v1[0] = (("foo bar" + "a") + "foo bar")
                    }
                }
                _ = fi10
            }
        }
    case 854:
        if false {
            _ = len(v1)
            var v21 uint8 = 118
            var v22 bool = true
            _ = v21
            _ = v22
        } else {
            var v23 int = 0
            _ = v23
        }
    }
    var v24 float64 = 0.0
    if (-75 <= 127) {
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
    } else {
        for fi25 := 0; fi25 < 18; fi25++ {
            v2 -= 0
            _ = fi25
        }
        v4 = (v4 + v4)
        _ = len(v1)
    }
    v26 := GenS2{F0: false, F1: (v3 | v3), F2: -658}
    v4 = 0.1
    var v27 bool = (255 <= 127)
    var v28 int64 = -1
    for fi29 := 0; fi29 < 6; fi29++ {
        v30 := GenS1{F0: 2, F1: 255}
        _ = fi29
        _ = v30
    }
    var v31 string = "a"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
}
