package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint }
type GenS2 struct { F0 int8; F1 uint32; F2 int16 }

func main() {
    var v1 float64 = (-2.25 * (3.14159 + 0.1))
    var v2 uint64 = 67
    var v3 int16 = 1
    var v4 uint8 = 42
    switch 9223372036854775807 {
    case 42:
        v1 = (v1 * ((0.0 * (1e10 - v1)) * v1))
    case 32767:
        var v5 float64 = v1
        _ = v5
    case 32768:
        v4 = v4
    default:
        if false {
            fmt.Println(v1, v2, v4)
            fmt.Println(v3)
        }
    }
    v6 := []int{32767, 32767, 32768}
    fmt.Println(v2, v2)
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(v4, v4)
        switch "\n\t" {
        case "hello":
            fmt.Println(v1, v4, v1)
        case "a":
            v4 -= 22
        case "\n\t":
            if (!(-396 != -1)) {
                var v8 int64 = 780
                v8 += v8
                _ = v8
            } else {
                if false {
                    var v9 uint8 = (v4 % (((v4 / ((v4) | 1))) | 1))
                    for fi10 := 0; fi10 < 9; fi10++ {
                        var v11 uint32 = 2147483648
                        if ((v1 * (v1 + v1)) >= (float64(0.1) + 0.0)) {
                            v3 -= 32767
                            v4 = uint8(v11)
                            var v12 string = ("hello" + "foo bar")
                            _ = v12
                        }
                        v4 += v9
                        if len(v6) > 0 {
                            v6[0] = (fi7 + 2147483647)
                        }
                        _ = fi10
                        _ = v11
                    }
                    if (0.1 < 0.1) {
                        for fi13 := 0; fi13 < 7; fi13++ {
                            v3 -= (v3 - v3)
                            v3 -= v3
                            var v14 bool = true
                            v14 = (float64(0.0) == -2.25)
                            var v15 int32 = 2147483647
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    } else {
                        var v16 int64 = 0
                        v6 = append(v6, fi7)
                        if (100 > 42) {
                            switch ((("a" + "hello") + "") + ("\n\t" + ("\n\t" + ""))) {
                            case "\n\t":
                                v1 = 0.0
                            case "hello":
                                v17 := GenS1{F0: "hello", F1: -2147483648, F2: 232}
                                _ = v17
                            case "a":
                                var v18 bool = (876 >= 127)
                                _ = v18
                            }
                        } else {
                            v1 = 0.0
                        }
                        _ = v16
                    }
                    _ = v9
                } else {
                    var v19 bool = (!true)
                    v20 := GenS1{F0: (("\n\t" + "a") + "hello"), F1: -514, F2: 18446744073709551615}
                    var v21 uint16 = 2
                    _ = v19
                    _ = v20
                    _ = v21
                }
            }
        }
        _ = fi7
    }
    var v22 uint32 = 4294967295
    v4 = (v4 ^ v4)
    if ((1 >= 42) || ((!(42 != -128)) || true)) {
        var v23 uint = 2
        for fi24 := 0; fi24 < 4; fi24++ {
            fmt.Println(v4, v2, v22)
            _ = fi24
        }
        _ = v23
    }
    var v25 float32 = (1.5 / -2.25)
    v26 := GenS2{F0: 127, F1: (^(v22)), F2: (v3 + (v3 * (v3 & v3)))}
    if len(v6) > 0 {
        v6[0] = 625
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v6
}
