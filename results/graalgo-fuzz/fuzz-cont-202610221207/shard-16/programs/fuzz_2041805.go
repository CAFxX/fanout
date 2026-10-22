package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int64; F1 uint16 }
func genFn0() string {
    var v4 uint64 = 0
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := GenS1{F0: (-(fi5)), F1: 42}
        var v7 uint = 979
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = v4
    return ((("a" + "foo bar") + "a") + "\n\t")
}

func main() {
    v1 := []string{"hello"}
    var v2 int8 = -128
    var v3 float64 = 3.14159
    fmt.Println(v3, v2)
    v8 := GenS1{F0: -15, F1: 42}
    v1 = append(v1, "")
    var v9 uint32 = 2
    v3 -= 3.14159
    switch 128 {
    case 0:
        if false {
            for fi10 := 0; fi10 < 3; fi10++ {
                switch fi10 {
                case -250:
                    v2 += v2
                default:
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(v9)
                    case "a":
                        if (("\n\t" + "foo bar") != ("\n\t" + ("foo bar" + ""))) {
                            if (v9 > v9) {
                                var v11 int32 = -1
                                _ = v11
                            } else {
                                switch fi10 {
                                case -9223372036854775808:
                                    var v12 int32 = 42
                                    _ = v12
                                case 4294967295:
                                    if true {
                                        fmt.Println(fi10, v9)
                                        if true {
                                            var v13 uint32 = v9
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = v13
                                        } else {
                                            var v14 int8 = (v2 << ((265) & 63))
                                            switch ("a" + "\n\t") {
                                            case "\n\t":
                                                if len(v1) > 0 {
                                                    v1[0] = "a"
                                                }
                                            case "a":
                                                for fi15 := 0; fi15 < 18; fi15++ {
                                                    switch fi15 {
                                                    case 2147483647:
                                                        v3 += (float64((((1.5 + 1e10) - (((0.0 * (1.5 / 1.0)) * (1.5 / 0.1)) - (3.14159 * -2.25))) - 1e10)) + (v3 / ((((1.5 + 1e10) - 0.1) + (v3 * (v3 * 0.0))) * (v3 + 1e10))))
                                                    case 127:
                                                        var v16 int64 = 442
                                                        _ = v16
                                                    case -2147483648:
                                                        v9 += 2
                                                    }
                                                    _ = fi15
                                                }
                                            default:
                                                var v17 uint16 = 2
                                                _ = v17
                                            }
                                            _ = v14
                                        }
                                        v9 = v9
                                    } else {
                                        var v18 uint = 1
                                        v19 := GenS2{F0: 42, F1: 704}
                                        _ = v18
                                        _ = v19
                                    }
                                case -1:
                                    v9 -= (v9 ^ (v9 | v9))
                                default:
                                    _ = len(v1)
                                }
                            }
                            var v20 uint = 2
                            _ = v20
                        }
                    }
                }
                for fi21 := 0; fi21 < 1; fi21++ {
                    v1 = append(v1, ("\n\t" + "a"))
                    _ = fi21
                }
                switch (fi10 / (((-(fi10))) | 1)) {
                case -2147483648:
                    var v22 int = (fi10 / ((-882) | 1))
                    _ = v22
                default:
                    v3 -= (v3 - (((v3 * -2.25) / v3) + v3))
                }
                _ = fi10
            }
            fmt.Println(v3, v9, v3)
            var v23 uint32 = v9
            _ = v23
        }
    default:
        if false {
            v3 = v3
            var v24 int16 = -32768
            if (!((((1.5 + 1e10) / ((1e10 * 1e10) * float32((-2.25 + ((0.0 / v3) * 1e10))))) + 3.14159) > 0.0)) {
                v2 -= (v2 ^ (v2 * (v2 / ((v2) | 1))))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = v24
        }
    }
    var v25 int16 = 42
    fmt.Println(v3, v3, v25)
    var v26 string = "a"
    switch ((v26 + (v26 + "a")) + v26) {
    case "a":
        var v27 int8 = 127
        _ = v27
    case "foo bar":
        var v28 bool = (float32((0.0 + float64((3.14159 * -2.25)))) > 0.1)
        _ = v28
    case "":
        v25 -= v25
    default:
        var v29 int = 2147483647
        _ = v29
    }
    fmt.Println(v2)
    var v30 int64 = -9223372036854775808
    v30 -= (v30 / ((v30) | 1))
    _ = len(v1)
    fmt.Println(v9, v25, v25)
    v2 = (v2 % ((v2) | 1))
    v9 = v9
    v25 += v25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v8
    _ = v9
}
