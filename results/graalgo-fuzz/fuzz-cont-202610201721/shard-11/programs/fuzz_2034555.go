package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 int }

func main() {
    var v1 bool = (false || false)
    v2 := make(map[string]bool)
    var v3 uint = 25
    var v4 int64 = -9223372036854775808
    var v5 uint32 = 2
    var v6 float32 = 1.5
    var v7 int8 = -128
    v4 = v4
    for fi8 := 0; fi8 < 15; fi8++ {
        v2[("\n\t" + ("\n\t" + ("a" + "a")))] = v1
        _ = fi8
    }
    if (!false) {
        v9 := []int{}
        _ = v9
    } else {
        v4 = (^(v4))
        delete(v2, "\n\t")
        for fi10 := 0; fi10 < 8; fi10++ {
            fmt.Println(v1, v5)
            v11 := GenS1{F0: fi10, F1: (v3 ^ 366)}
            var v12 int64 = (v4 + 9223372036854775807)
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    v6 = ((v6 - 3.14159) + v6)
    for fi13 := 0; fi13 < 10; fi13++ {
        v4 = v4
        _ = fi13
    }
    v4 += -272
    for fi14 := 0; fi14 < 10; fi14++ {
        v7 = v7
        var v15 int64 = v4
        _ = fi14
        _ = v15
    }
    v4 = v4
    switch (("a" + "foo bar") + ("a" + (("hello" + ("" + "foo bar")) + ("foo bar" + ("a" + "\n\t"))))) {
    case "a":
        v16 := GenS2{F0: ""}
        _ = v16
    case "\n\t":
        switch 126 {
        case 1:
            if (18446744073709551615 == uint64(v3)) {
                v5 = v5
                v17 := GenS2{F0: ("hello" + ("\n\t" + ("foo bar" + "a")))}
                _ = v17
            } else {
                for fi18 := 0; fi18 < 8; fi18++ {
                    if (96 <= 0) {
                        fmt.Println(v1)
                        if (fi18 != (fi18 ^ (fi18 / ((4294967295) | 1)))) {
                            var v19 int64 = -257
                            delete(v2, ("\n\t" + ("\n\t" + "hello")))
                            _ = v19
                        } else {
                            if (v1 || (true && (!(!v1)))) {
                                for fi20 := 0; fi20 < 20; fi20++ {
                                    v21 := GenS2{F0: "\n\t"}
                                    fmt.Println(v3, v4, v5)
                                    v22 := GenS1{F0: fi20, F1: 0}
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                                v3 -= (v3 >> ((2) & 63))
                            } else {
                                var v23 uint = v3
                                _ = v23
                            }
                        }
                        v24 := GenS3{F0: 892, F1: -128}
                        _ = v24
                    } else {
                        fmt.Println(v1)
                    }
                    _ = fi18
                }
                v3 = (v3 ^ v3)
            }
        case -9223372036854775808:
            fmt.Println(v1, v1, v6)
        default:
            if (v1 || v1) {
                for fi25 := 0; fi25 < 2; fi25++ {
                    v26 := GenS3{F0: 255, F1: v7}
                    v27 := GenS4{F0: (fi25 / ((fi25) | 1))}
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            }
        }
    case "foo bar":
        var v28 int8 = int8(v4)
        _ = v28
    }
    v1 = (((v4 + v4) > v4) || v1)
    fmt.Println(v7)
    fmt.Println(v3, v3)
    var v29 float64 = (0.1 * (0.1 / (1.5 * (((((0.1 * 3.14159) * 1.5) / (0.1 + (((0.0 - 3.14159) - 1.5) - float64((v6 - -2.25))))) / float64(3.14159)) + (1e10 / float64(1.5))))))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
