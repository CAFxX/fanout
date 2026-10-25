package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 bool }
func genFn0(p0 int32, p1 int64, p2 string) int {
    if (1 <= 1) {
        p1 -= p1
        p1 = (p1 + p1)
    } else {
        p0 += (p0 + 0)
        p1 += 9223372036854775807
        v7 := []int{-9223372036854775808}
        _ = v7
    }
    p0 = 2147483647
    var v8 uint16 = 65535
    v9 := GenS1{F0: 2147483648, F1: 893, F2: true}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return -987
}

func main() {
    v1 := []uint32{858, 2, 782}
    var v2 uint = 42
    v3 := []string{("\n\t" + ("foo bar" + "foo bar")), ("a" + "a"), "a", (((("hello" + "a") + "") + "\n\t") + "a")}
    var v4 bool = (!(!false))
    v5 := []int{-424, -9223372036854775808, -994}
    var v6 uint16 = 2
    var v10 float64 = (-2.25 * 1e10)
    fmt.Println(v2, v2, v2)
    var v11 uint32 = 717
    fmt.Println(v4, v4)
    for fi12 := 0; fi12 < 10; fi12++ {
        if v4 {
            switch ("a" + (("a" + "") + "")) {
            case "foo bar":
                switch (fi12 / ((genFn0(592, 9223372036854775807, ("\n\t" + "hello"))) | 1)) {
                case 42:
                    if (v6 <= 2) {
                        fmt.Println(v11, v10, v4)
                        if v4 {
                            switch "" {
                            case "foo bar":
                                var v13 uint32 = (v11 >> ((v2) & 63))
                                _ = v13
                            default:
                                _ = len(v1)
                            }
                            fmt.Println(v4, v11, v4)
                            fmt.Println(v2, v6)
                        } else {
                            if len(v1) > 0 {
                                v1[0] = (^(v11))
                            }
                        }
                        for fi14 := 0; fi14 < 10; fi14++ {
                            switch "foo bar" {
                            case "\n\t":
                                v2 -= v2
                            case "hello":
                                v15 := GenS1{F0: (v11 + v11), F1: 0, F2: (1 > 100)}
                                _ = v15
                            case "foo bar":
                                var v16 uint = 2
                                _ = v16
                            default:
                                switch (-(fi12)) {
                                case -9223372036854775808:
                                    v17 := GenS1{F0: v11, F1: 1, F2: v4}
                                    _ = v17
                                case 128:
                                    var v18 uint = 838
                                    _ = v18
                                case 42:
                                    v2 += (v2 * 65535)
                                    var v19 float64 = 3.14159
                                    _ = v19
                                default:
                                    if (127 != -113) {
                                        _ = len(v5)
                                    } else {
                                        fmt.Println(v10)
                                    }
                                }
                            }
                            fmt.Println(v11)
                            _ = fi14
                        }
                    } else {
                        var v20 int16 = 220
                        var v21 int8 = -1
                        v20 += v20
                        _ = v20
                        _ = v21
                    }
                case 0:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            default:
                v4 = (((fi12 != fi12) || (!false)) && (462 <= 9223372036854775807))
            }
            v11 += uint32(v6)
            var v22 int32 = -2147483648
            _ = v22
        }
        _ = fi12
    }
    v1 = append(v1, v11)
    var v23 float64 = float64(((1e10 / (0.0 * (1e10 / (float32((v10 * v10)) * (3.14159 * (3.14159 + (0.1 + 0.1))))))) + 1.5))
    fmt.Println(v4)
    switch ("foo bar" + ("hello" + "a")) {
    case "hello":
        var v24 uint32 = v11
        _ = v24
    case "a":
        var v25 int16 = 810
        _ = v25
    case "foo bar":
        if true {
            _ = len(v3)
            var v26 int = 32768
            _ = v26
        }
    default:
        v27 := GenS1{F0: v11, F1: -548, F2: v4}
        _ = v27
    }
    switch "hello" {
    case "":
        for fi28 := 0; fi28 < 13; fi28++ {
            var v29 int16 = -1
            for fi30 := 0; fi30 < 9; fi30++ {
                if len(v1) > 0 {
                    v1[0] = 42
                }
                var v31 float32 = float32(v10)
                fmt.Println(v31, fi28)
                v1 = append(v1, (v11 | (v11 - 448)))
                _ = fi30
                _ = v31
            }
            var v32 int16 = (v29 / ((0) | 1))
            _ = fi28
            _ = v29
            _ = v32
        }
    case "\n\t":
        var v33 string = ("a" + "")
        _ = v33
    }
    var v34 int64 = 42
    switch ((("a" + "") + "foo bar") + "a") {
    case "\n\t":
        v10 = (0.0 - v10)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
