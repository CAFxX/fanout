package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int32; F2 int16 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]int)
    v3 := make(map[uint32]bool)
    var v4 int64 = 0
    switch 127 {
    case -2147483648:
        v4 += -9223372036854775808
    case -9223372036854775808:
        v4 = (v4 / (((^(v4))) | 1))
    case 32767:
        var v5 int = int(v4)
        _ = v5
    default:
        fmt.Println(v4)
    }
    v6 := []int{}
    var v7 int8 = 58
    if (false && ((!(1.5 >= -2.25)) || (!(2 != 18446744073709551615)))) {
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := GenS1{F0: false}
            _ = fi8
            _ = v9
        }
    }
    var v10 float64 = (1.5 - ((1.5 * 1e10) / 1e10))
    fmt.Println(v4, v10)
    var v11 int = 32767
    if (v4 <= v4) {
        if true {
            v7 += (v7 | (v7 | (v7 * v7)))
            switch v11 {
            case 0:
                for fi12 := 0; fi12 < 13; fi12++ {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = fi12
                }
            case 9223372036854775807:
                var v13 int64 = 629
                _ = v13
            }
            for fi14 := 0; fi14 < 3; fi14++ {
                var v15 int8 = (v7 ^ v7)
                if true {
                    if false {
                        if true {
                            v15 -= -128
                        } else {
                            v16 := v1["foo bar"]
                            _ = v16
                        }
                    } else {
                        var v17 int64 = -9223372036854775808
                        fmt.Println(v7)
                        _ = v17
                    }
                    v7 += v7
                    v18 := GenS1{F0: false}
                    _ = v18
                } else {
                    var v19 int64 = -639
                    for fi20 := 0; fi20 < 5; fi20++ {
                        v15 = v15
                        _ = fi20
                    }
                    _ = v19
                }
                v7 -= (v15 ^ 0)
                _ = fi14
                _ = v15
            }
        }
    }
    var v21 int8 = v7
    v4 = (v4 & -9223372036854775808)
    var v22 int8 = v7
    for fi23 := 0; fi23 < 1; fi23++ {
        if false {
            for fi24 := 0; fi24 < 17; fi24++ {
                if ("foo bar" > "a") {
                    if (v21 != (v21 / (((-(v7))) | 1))) {
                        var v25 bool = true
                        _ = v25
                    } else {
                        if (false || false) {
                            v22 = (v22 * (v21 * 0))
                            delete(v1, "a")
                        } else {
                            _ = len(v6)
                        }
                        v7 = (v22 - -1)
                        fmt.Println(v21, v22, fi24)
                    }
                } else {
                    fmt.Println(v4, fi23, v10)
                }
                fmt.Println(fi24, v22)
                if true {
                    v21 -= v22
                    v21 = v7
                    delete(v3, 42)
                    v7 = (v22 << ((819) & 63))
                    v26 := v1["a"]
                    _ = v26
                }
                var v27 uint8 = 0
                _ = fi24
                _ = v27
            }
            v10 += v10
        }
        _ = fi23
    }
    v28 := GenS2{F0: false, F1: -2147483648, F2: -500}
    for fi29 := 0; fi29 < 8; fi29++ {
        v21 += v22
        _ = fi29
    }
    v4 += (^(v4))
    switch ((("" + "hello") + "foo bar") + (("a" + "\n\t") + "hello")) {
    case "foo bar":
        if false {
            _ = len(v6)
            v21 = (v21 << ((417) & 63))
            v30 := v3[1]
            _ = v30
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
