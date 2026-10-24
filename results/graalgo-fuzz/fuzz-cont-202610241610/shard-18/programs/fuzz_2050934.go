package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
func genFn0(p0 uint16, p1 int16) int32 {
    v5 := []int{-107, 349, int(p0)}
    p1 -= (p1 % (((p1 >> ((65535) & 63))) | 1))
    for fi6 := 0; fi6 < 7; fi6++ {
        if true {
            var v7 uint = 18446744073709551615
            for fi8 := 0; fi8 < 19; fi8++ {
                if len(v5) > 0 {
                    v5[0] = (fi6 & (fi8 << ((v7) & 63)))
                }
                switch "hello" {
                case "\n\t":
                    p0 = p0
                default:
                    var v9 float64 = (0.0 + 0.1)
                    _ = v9
                }
                if true {
                    p1 = 808
                }
                _ = fi8
            }
            _ = v7
        }
        p1 += (p1 * p1)
        _ = fi6
    }
    p0 = (^(p0))
    if true {
        v5 = append(v5, 379)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return 42
}

func main() {
    var v1 int16 = 42
    var v2 int = 2147483647
    v3 := []string{"\n\t", "\n\t", ("" + "a"), "foo bar"}
    v4 := make(map[int]uint32)
    var v10 uint16 = 703
    var v11 string = "\n\t"
    switch v2 {
    case 32767:
        switch (v11 + "hello") {
        case "\n\t":
            v12 := GenS1{F0: (4294967295 != 42), F1: 42}
            _ = v12
        default:
            v11 = (v11 + (("foo bar" + "a") + "a"))
        }
    case 512:
        v10 -= (^(v10))
    case 42:
        v2 = 1
    }
    delete(v4, 232)
    switch "a" {
    case "hello":
        if false {
            if true {
                var v13 int = (v2 & (v2 | -933))
                for fi14 := 0; fi14 < 9; fi14++ {
                    if false {
                        var v15 int32 = 42
                        v3 = append(v3, "a")
                        _ = v15
                    }
                    v13 = int(v10)
                    _ = fi14
                }
                var v16 int64 = -9223372036854775808
                _ = v13
                _ = v16
            } else {
                v11 = v11
                if (!false) {
                    if false {
                        var v17 int32 = 505
                        for fi18 := 0; fi18 < 13; fi18++ {
                            if false {
                                for fi19 := 0; fi19 < 18; fi19++ {
                                    if false {
                                        var v20 int = fi18
                                        if ((2 != 128) && true) {
                                            v21 := GenS1{F0: false, F1: (-(v2))}
                                            var v22 uint16 = v10
                                            v22 += v10
                                            _ = v21
                                            _ = v22
                                        } else {
                                            var v23 int = fi19
                                            v4[(fi19 ^ (-(v2)))] = 496
                                            v10 -= 255
                                            _ = v23
                                        }
                                        _ = v20
                                    } else {
                                        if (true && false) {
                                            var v24 int8 = 0
                                            var v25 bool = (false || false)
                                            if v25 {
                                                var v26 uint16 = (v10 & (v10 / ((376) | 1)))
                                                v10 -= 1
                                                _ = v26
                                            } else {
                                                var v27 float64 = (0.1 * 3.14159)
                                                var v28 bool = v25
                                                _ = len(v3)
                                                _ = v27
                                                _ = v28
                                            }
                                            var v29 string = (v11 + v11)
                                            _ = v24
                                            _ = v25
                                            _ = v29
                                        }
                                    }
                                    v17 += (v17 ^ (v17 << ((uint(v10)) & 63)))
                                    _ = fi19
                                }
                                v2 -= (fi18 ^ int(v1))
                            }
                            _ = fi18
                        }
                        _ = v17
                    }
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(v10)
                    }
                }
                v30 := GenS1{F0: false, F1: v2}
                _ = v30
            }
        }
    case "":
        fmt.Println(v1)
    case "\n\t":
        if false {
            v11 = "hello"
            var v31 int16 = v1
            v32 := GenS1{F0: false, F1: (v2 / ((-172) | 1))}
            _ = v31
            _ = v32
        }
    default:
        var v33 string = ""
        _ = v33
    }
    var v34 bool = true
    var v35 int8 = 52
    fmt.Println(v10)
    fmt.Println(v11)
    _ = len(v3)
    v36 := GenS1{F0: (2 <= 4294967295), F1: -9223372036854775808}
    if (v34 || v34) {
        v34 = (65535 <= 42)
        v35 = (v35 | (v35 / ((v35) | 1)))
        v10 += v10
    }
    switch 32768 {
    case 1:
        if true {
            var v37 int64 = 0
            fmt.Println(v1, v2)
            _ = v37
        }
    case 127:
        for fi38 := 0; fi38 < 12; fi38++ {
            var v39 uint8 = 1
            for fi40 := 0; fi40 < 11; fi40++ {
                fmt.Println(v39, v2)
                _ = fi40
            }
            _ = fi38
            _ = v39
        }
    case -599:
        v41 := GenS1{F0: false, F1: (v2 & v2)}
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
