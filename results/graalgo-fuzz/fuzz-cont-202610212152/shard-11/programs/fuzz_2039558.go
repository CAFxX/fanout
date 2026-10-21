package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int8 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := []int{-1, -880, 2147483647, 362}
    v2 := make(map[int]bool)
    v3 := make(map[string]string)
    var v4 uint8 = 45
    var v5 int8 = -128
    var v6 uint64 = 315
    if (4294967295 > 945) {
        v6 = v6
        var v7 uint = 42
        v8 := GenS1{F0: -9223372036854775808, F1: v7, F2: (v5 >> ((628) & 63))}
        _ = v7
        _ = v8
    } else {
        v2[-9223372036854775808] = (0.0 >= 1e10)
        v9 := GenS1{F0: -901, F1: 18446744073709551615, F2: v5}
        for fi10 := 0; fi10 < 12; fi10++ {
            fmt.Println(v5)
            var v11 int32 = 2147483647
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    if ((v6 == 930) || ((v6 ^ (^(v6))) <= v6)) {
        _ = len(v2)
        var v12 uint = 493
        var v13 float32 = float32(float64((1e10 / 1e10)))
        _ = v12
        _ = v13
    } else {
        var v14 int32 = 42
        v5 -= (-(v5))
        _ = v14
    }
    _ = len(v1)
    var v15 int = 32767
    fmt.Println(v6, v6, v5)
    if true {
        fmt.Println(v6, v15, v5)
        if (136 <= -1) {
            if (false && (((^(v15)) <= (v15 - (v15 + v15))) && false)) {
                var v16 int8 = 100
                switch "\n\t" {
                case "a":
                    var v17 int32 = 1
                    _ = v17
                case "\n\t":
                    delete(v2, -649)
                default:
                    v4 += v4
                }
                v16 += (v5 - (^(v16)))
                _ = v16
            } else {
                if true {
                    v4 -= (^(v4))
                    fmt.Println(v15, v6)
                    var v18 float64 = float64(-2.25)
                    _ = v18
                } else {
                    v19 := GenS1{F0: -300, F1: 930, F2: v5}
                    if true {
                        for fi20 := 0; fi20 < 9; fi20++ {
                            for fi21 := 0; fi21 < 14; fi21++ {
                                var v22 int64 = 1
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = fi21
                                _ = v22
                            }
                            _ = fi20
                        }
                        v23 := GenS2{F0: 100}
                        _ = v23
                    } else {
                        if ((("" + "\n\t") + "") <= "hello") {
                            v6 = v6
                        } else {
                            for fi24 := 0; fi24 < 15; fi24++ {
                                v5 += (v5 + v5)
                                _ = fi24
                            }
                            fmt.Println(v6, v15, v15)
                        }
                        v15 -= v15
                    }
                    v5 += -32
                    _ = v19
                }
            }
        }
        var v25 uint = 65535
        _ = v25
    } else {
        v6 -= v6
    }
    if (2147483648 >= 2147483647) {
        var v26 float32 = 0.0
        _ = v26
    } else {
        var v27 uint64 = (v6 + (v6 & 4294967295))
        _ = v27
    }
    _ = len(v1)
    var v28 string = ("" + "a")
    _ = len(v1)
    var v29 uint16 = 1
    if false {
        v6 -= v6
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
