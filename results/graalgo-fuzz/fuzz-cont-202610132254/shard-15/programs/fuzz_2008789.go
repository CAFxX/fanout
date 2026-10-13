package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int16; F1 int64; F2 int8 }

func main() {
    var v1 int16 = -32768
    var v2 int64 = 427
    var v3 uint16 = 42
    v4 := []int{522}
    var v5 uint = 42
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 int32 = -1
        _ = fi6
        _ = v7
    }
    switch -2147483648 {
    case 42:
        fmt.Println(v3, v5)
    case 32767:
        var v8 int32 = 152
        _ = v8
    }
    fmt.Println(v2)
    var v9 uint8 = 255
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v10 := GenS1{F0: -2147483648, F1: ("a" + (((("foo bar" + "foo bar") + (("a" + ("" + "a")) + "")) + "\n\t") + ("hello" + "a")))}
    var v11 int8 = 42
    v12 := GenS1{F0: -907, F1: "a"}
    for fi13 := 0; fi13 < 14; fi13++ {
        for fi14 := 0; fi14 < 6; fi14++ {
            if true {
                if true {
                    if ("a" > "") {
                        if len(v4) > 0 {
                            v4[0] = fi13
                        }
                        v1 -= int16(v3)
                        if false {
                            switch "a" {
                            case "a":
                                fmt.Println(v5, fi14, v2)
                            case "foo bar":
                                v3 = 628
                                var v15 bool = false
                                v4 = append(v4, -257)
                                _ = v15
                            case "\n\t":
                                v16 := GenS1{F0: -2147483648, F1: "\n\t"}
                                _ = v16
                            }
                            fmt.Println(v11, v11)
                        }
                    } else {
                        switch "hello" {
                        case "foo bar":
                            if len(v4) > 0 {
                                v4[0] = int(v1)
                            }
                        case "":
                            v17 := GenS1{F0: -426, F1: ""}
                            _ = v17
                        case "hello":
                            v4 = append(v4, 2147483647)
                        default:
                            var v18 int64 = v2
                            _ = v18
                        }
                    }
                    var v19 int = fi14
                    fmt.Println(fi13)
                    _ = v19
                }
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(v9)
                case "a":
                    if (!(true && (false || true))) {
                        v9 = v9
                        var v20 uint8 = v9
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v20
                    } else {
                        var v21 bool = false
                        if len(v4) > 0 {
                            v4[0] = 128
                        }
                        v4 = append(v4, fi14)
                        v22 := GenS1{F0: -1, F1: ("foo bar" + "a")}
                        _ = v21
                        _ = v22
                    }
                default:
                    switch 127 {
                    case 32768:
                        v5 += (v5 ^ v5)
                    default:
                        _ = len(v4)
                    }
                }
                v23 := GenS1{F0: -1, F1: "\n\t"}
                _ = v23
            } else {
                v11 -= 108
                var v24 uint = v5
                _ = v24
            }
            fmt.Println(v2)
            v5 = v5
            _ = fi14
        }
        _ = fi13
    }
    if ((!(4294967295 != -9223372036854775808)) && true) {
        var v25 bool = (0.1 <= -2.25)
        for fi26 := 0; fi26 < 8; fi26++ {
            v27 := GenS2{F0: -1, F1: v2, F2: int8(v2)}
            _ = fi26
            _ = v27
        }
        _ = v25
    } else {
        v3 += v3
        v1 -= -811
        var v28 int64 = (v2 & -702)
        _ = v28
    }
    var v29 int16 = (v1 >> (((v5 >> (((v5 + v5)) & 63))) & 63))
    if false {
        v9 += (^(v9))
        v30 := GenS2{F0: 1000, F1: -411, F2: (v11 % ((v11) | 1))}
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
