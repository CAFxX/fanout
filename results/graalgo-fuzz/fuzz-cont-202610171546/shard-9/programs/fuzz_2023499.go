package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 string }
type GenS2 struct { F0 uint32 }
func genFn0() string {
    v7 := []int{-177}
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 float32 = (-2.25 + 0.1)
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi8
        _ = v9
    }
    _ = v7
    return "a"
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = -848
    var v3 uint = 42
    var v4 int16 = 1000
    v5 := make(map[string]bool)
    v6 := []int{}
    v4 = v4
    v2 += v2
    if true {
        for fi10 := 0; fi10 < 3; fi10++ {
            v3 -= 1
            switch "" {
            case "a":
                var v11 uint32 = 411
                _ = v11
            default:
                v4 = v4
            }
            _ = fi10
        }
        var v12 int8 = 42
        if (v2 == v2) {
            _ = len(v5)
        } else {
            if true {
                var v13 int8 = v12
                v3 = 42
                if false {
                    for fi14 := 0; fi14 < 1; fi14++ {
                        switch "hello" {
                        case "foo bar":
                            if false {
                                if (!false) {
                                    fmt.Println(v4)
                                    v12 += (v13 ^ v13)
                                } else {
                                    var v15 uint8 = 128
                                    _ = v15
                                }
                                var v16 float64 = 1e10
                                fmt.Println(fi14, v2)
                                _ = v16
                            }
                        case "a":
                            var v17 int8 = int8(fi14)
                            _ = v17
                        default:
                            v3 += 2
                        }
                        for fi18 := 0; fi18 < 17; fi18++ {
                            var v19 int8 = (v13 << ((v3) & 63))
                            if true {
                                v19 = (v19 >> ((v3) & 63))
                            }
                            v4 = v4
                            v4 = (v4 + v4)
                            var v20 string = "\n\t"
                            v20 = "foo bar"
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        _ = fi14
                    }
                } else {
                    v21 := GenS1{F0: 166, F1: -1, F2: ("a" + "\n\t")}
                    v22 := GenS2{F0: 4294967295}
                    if false {
                        fmt.Println(v4, v12, v2)
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        fmt.Println(v12, v2)
                    }
                    _ = v21
                    _ = v22
                }
                _ = v13
            }
            v4 = (^(v4))
        }
        _ = v12
    }
    switch 4294967295 {
    case -9223372036854775808:
        v23 := GenS2{F0: 2}
        _ = v23
    case -1:
        var v24 uint16 = 65535
        _ = v24
    default:
        v3 = v3
    }
    if false {
        var v25 int64 = 672
        v26 := GenS1{F0: 138, F1: v25, F2: "\n\t"}
        _ = v25
        _ = v26
    } else {
        v27 := GenS1{F0: 2, F1: 42, F2: ((("a" + "\n\t") + "hello") + "a")}
        v3 = 18446744073709551615
        var v28 uint = (v3 | (^(v3)))
        _ = v27
        _ = v28
    }
    v3 = (v3 * (v3 | 1))
    var v29 int32 = 0
    switch (v2 / (((^(v2))) | 1)) {
    case 292:
        _ = len(v1)
    }
    switch (-(v2)) {
    case -813:
        for fi30 := 0; fi30 < 15; fi30++ {
            if false {
                var v31 int8 = -1
                _ = v31
            }
            var v32 int64 = -1
            v1[fi30] = (0.1 != 1.5)
            _ = fi30
            _ = v32
        }
    case -2147483648:
        var v33 int = v2
        _ = v33
    case 1:
        v29 += (v29 ^ (^(v29)))
    default:
        for fi34 := 0; fi34 < 20; fi34++ {
            fmt.Println(v2)
            if len(v6) > 0 {
                v6[0] = (fi34 * fi34)
            }
            var v35 int32 = v29
            _ = fi34
            _ = v35
        }
    }
    v2 -= v2
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
