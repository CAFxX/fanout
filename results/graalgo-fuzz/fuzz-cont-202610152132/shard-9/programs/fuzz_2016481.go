package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint }
func genFn0(p0 int8) uint32 {
    p0 = (p0 + (p0 | p0))
    p0 -= p0
    v5 := []int{4294967295, 481, 2147483647, 42}
    _ = p0
    _ = v5
    return 65535
}

func main() {
    v1 := []int{42, 1, 42, 32767}
    var v2 int8 = -115
    v3 := []int{}
    v4 := make(map[uint32]uint32)
    for fi6 := 0; fi6 < 16; fi6++ {
        for fi7 := 0; fi7 < 12; fi7++ {
            fmt.Println(fi6, fi6)
            v8 := GenS1{F0: 512, F1: 42, F2: -668}
            var v9 int64 = 1
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    v10 := GenS2{F0: 18446744073709551615}
    switch 127 {
    case 9223372036854775807:
        v2 = 0
    case 0:
        fmt.Println(v2)
    case -1:
        if (0 <= -1) {
            var v11 int16 = 0
            _ = v11
        }
    default:
        for fi12 := 0; fi12 < 12; fi12++ {
            v2 -= v2
            _ = fi12
        }
    }
    v2 = v2
    v2 += v2
    var v13 bool = (822 <= 820)
    fmt.Println(v2, v13)
    var v14 string = ""
    for fi15 := 0; fi15 < 3; fi15++ {
        v2 += -64
        if true {
            var v16 uint = 281
            if (0 > 1000) {
                v17 := GenS2{F0: (^(v16))}
                if v13 {
                    var v18 bool = v13
                    var v19 uint32 = 4294967295
                    v16 += (v16 << (((v16 + (v16 % (((v16 ^ (v16 - (v16 | 2)))) | 1)))) & 63))
                    _ = v18
                    _ = v19
                } else {
                    v13 = (!v13)
                }
                _ = v17
            }
            if (42 == 42) {
                for fi20 := 0; fi20 < 1; fi20++ {
                    v14 = (v14 + ("a" + "a"))
                    _ = fi20
                }
            } else {
                for fi21 := 0; fi21 < 4; fi21++ {
                    v2 = v2
                    var v22 uint = 4294967295
                    _ = fi21
                    _ = v22
                }
            }
            _ = v16
        } else {
            if v13 {
                if v13 {
                    if (2147483647 <= 1) {
                        fmt.Println(v13, fi15)
                        v14 = v14
                    }
                    for fi23 := 0; fi23 < 6; fi23++ {
                        v14 = v14
                        _ = fi23
                    }
                } else {
                    if ((0 == 42) && false) {
                        switch (v14 + v14) {
                        case "\n\t":
                            fmt.Println(v14)
                        case "a":
                            for fi24 := 0; fi24 < 14; fi24++ {
                                fmt.Println(fi15, v14)
                                switch v14 {
                                case "a":
                                    v2 = v2
                                case "foo bar":
                                    fmt.Println(fi24, v2, v13)
                                case "":
                                    v2 = v2
                                default:
                                    v25 := GenS1{F0: -1, F1: 4294967295, F2: 42}
                                    _ = v25
                                }
                                switch "a" {
                                case "foo bar":
                                    v2 = 39
                                }
                                _ = fi24
                            }
                        }
                        v14 = "foo bar"
                    }
                    v2 += 100
                    var v26 uint16 = 2
                    _ = v26
                }
                for fi27 := 0; fi27 < 1; fi27++ {
                    fmt.Println(v13, v2)
                    _ = fi27
                }
                fmt.Println(fi15)
            }
            var v28 float32 = (((1.5 * 0.1) * 1e10) - (0.0 / 1.5))
            _ = v28
        }
        _ = fi15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
