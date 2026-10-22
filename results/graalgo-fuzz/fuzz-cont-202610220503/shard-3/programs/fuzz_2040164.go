package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint; F1 int32 }

func main() {
    v1 := []string{}
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 4294967295
    var v4 int16 = 42
    var v5 int = -1
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 uint8 = 128
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v8 bool = false
        _ = fi6
        _ = v7
        _ = v8
    }
    v4 -= 675
    if false {
        v4 = (v4 * v4)
        switch "hello" {
        case "a":
            if false {
                var v9 uint8 = 127
                _ = v9
            }
        case "":
            var v10 float64 = (1.5 + 1e10)
            _ = v10
        case "foo bar":
            fmt.Println(v5)
        }
        v3 += (v3 ^ 0)
    } else {
        if true {
            v5 += (-(v5))
            v11 := GenS1{F0: 42, F1: 65535}
            var v12 float64 = 1.5
            _ = v11
            _ = v12
        }
        fmt.Println(v4, v3, v4)
    }
    var v13 float32 = 0.0
    if (uint(v3) < 997) {
        v14 := GenS1{F0: -1, F1: 42}
        v13 += -2.25
        v15 := GenS2{F0: 42, F1: -1}
        _ = v14
        _ = v15
    } else {
        if (v13 < v13) {
            v13 = v13
            v1 = append(v1, "")
        }
        for fi16 := 0; fi16 < 3; fi16++ {
            v3 = v3
            _ = fi16
        }
    }
    var v17 float32 = (3.14159 / v13)
    fmt.Println(v17)
    v18 := GenS2{F0: 255, F1: 0}
    if false {
        _ = len(v1)
        switch ("\n\t" + "a") {
        case "hello":
            var v19 int64 = 1
            _ = v19
        case "\n\t":
            if false {
                switch "foo bar" {
                case "":
                    v20 := GenS1{F0: -979, F1: 2}
                    _ = v20
                case "foo bar":
                    v4 = v4
                case "\n\t":
                    v21 := GenS1{F0: (v4 / (((-(v4))) | 1)), F1: 255}
                    _ = v21
                default:
                    if (uint32(v3) == 4294967295) {
                        var v22 uint32 = 296
                        switch 4294967295 {
                        case 9223372036854775807:
                            fmt.Println(v4, v22)
                        case -2147483648:
                            v13 = 1.5
                        case 4294967295:
                            v5 += (v5 >> ((18446744073709551615) & 63))
                        default:
                            for fi23 := 0; fi23 < 18; fi23++ {
                                switch ("" + "hello") {
                                case "foo bar":
                                    for fi24 := 0; fi24 < 11; fi24++ {
                                        var v25 float64 = (-2.25 / 0.1)
                                        _ = fi24
                                        _ = v25
                                    }
                                case "":
                                    v2[(^(v22))] = (^(v22))
                                default:
                                    v26 := GenS1{F0: (v4 >> ((2) & 63)), F1: uint(v3)}
                                    _ = v26
                                }
                                v5 += v5
                                v27 := GenS1{F0: v4, F1: 255}
                                _ = fi23
                                _ = v27
                            }
                        }
                        _ = v22
                    }
                }
                v28 := GenS1{F0: v4, F1: 0}
                var v29 int8 = -10
                _ = v28
                _ = v29
            }
        }
    }
    var v30 int16 = -854
    v3 -= (^(v3))
    if false {
        v3 += v3
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
