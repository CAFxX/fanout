package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }

func main() {
    v1 := []string{"", "", ("a" + "hello"), ""}
    var v2 uint32 = 280
    var v3 int8 = 1
    var v4 uint = 18446744073709551615
    var v5 string = "foo bar"
    var v6 bool = (v3 != 1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (!(v5 != "foo bar")) {
        var v7 int = 210
        _ = len(v1)
        _ = v7
    } else {
        var v8 bool = (!v6)
        for fi9 := 0; fi9 < 11; fi9++ {
            v5 = v5
            _ = fi9
        }
        _ = v8
    }
    switch v5 {
    case "hello":
        var v10 uint64 = 18446744073709551615
        _ = v10
    case "":
        _ = len(v1)
    default:
        fmt.Println(v5)
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        v4 += (v4 - 572)
        switch (("" + "hello") + v5) {
        case "foo bar":
            if v6 {
                v1 = append(v1, v5)
            } else {
                if (((0.0 / (((float64(((0.0 * 0.1) - (-2.25 + -2.25))) - float64(0.0)) * (float64(1e10) * 1e10)) / 1.5)) > 1.5) && (65535 < 2)) {
                    var v12 uint8 = 95
                    var v13 int64 = -9223372036854775808
                    _ = v12
                    _ = v13
                } else {
                    fmt.Println(v2, v5, v5)
                    if v6 {
                        for fi14 := 0; fi14 < 8; fi14++ {
                            v5 = (v5 + "a")
                            var v15 float32 = 3.14159
                            _ = fi14
                            _ = v15
                        }
                        var v16 uint16 = 0
                        if (606 == 1) {
                            fmt.Println(v3)
                            if ((v3 << ((v4) & 63)) <= (^(v3))) {
                                fmt.Println(v3, v16)
                                v4 = v4
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            if v6 {
                                v17 := GenS1{F0: (!((v3 & (v3 ^ (v3 | 31))) != (v3 % ((0) | 1)))), F1: v5}
                                _ = v17
                            } else {
                                fmt.Println(fi11, v4, v6)
                                fmt.Println(fi11, v2)
                            }
                        } else {
                            var v18 float32 = 1e10
                            _ = v18
                        }
                        _ = v16
                    } else {
                        var v19 string = ""
                        _ = v19
                    }
                }
            }
        case "\n\t":
            v3 = -128
        case "":
            var v20 float64 = -2.25
            _ = v20
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi11
    }
    _ = len(v1)
    var v21 uint16 = 255
    var v22 int = 0
    var v23 int = (v22 & v22)
    v23 -= (v23 % (((v22 ^ int(v4))) | 1))
    switch ((v5 + v5) + "") {
    case "a":
        var v24 uint = (^(v4))
        _ = v24
    case "hello":
        var v25 float64 = 1.5
        _ = v25
    case "\n\t":
        v26 := GenS1{F0: v6, F1: v5}
        _ = v26
    }
    v2 += (v2 + (v2 / ((v2) | 1)))
    var v27 string = v5
    v5 = ((v27 + v27) + v27)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        fmt.Println(v23)
        v1 = append(v1, v5)
    }
    switch (v22 | (v23 ^ (v22 << ((v4) & 63)))) {
    case 614:
        v6 = true
    case 2147483647:
        fmt.Println(v27)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
