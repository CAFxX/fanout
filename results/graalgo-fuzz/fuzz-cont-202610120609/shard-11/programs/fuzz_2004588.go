package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 int; F2 uint16 }
type GenS3 struct { F0 int; F1 uint8; F2 int32 }
type GenS4 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS5 struct { F0 int16; F1 uint16 }
func genFn0(p0 int, p1 uint) int8 {
    p0 = -2147483648
    p1 += (p1 / ((p1) | 1))
    var v4 uint64 = 677
    _ = p0
    _ = p1
    _ = v4
    return 1
}
func genFn1(p0 uint) uint {
    if true {
        v5 := GenS1{F0: "foo bar"}
        if true {
            var v6 int32 = -237
            _ = v6
        } else {
            p0 = (p0 ^ p0)
            p0 += (p0 * 167)
            switch 0 {
            case -9223372036854775808:
                var v7 int64 = 826
                _ = v7
            }
        }
        _ = v5
    } else {
        var v8 int32 = -1
        _ = v8
    }
    var v9 int32 = 48
    _ = p0
    _ = v9
    return p0
}

func main() {
    var v1 int = 32768
    v2 := make(map[int]string)
    v3 := make(map[uint32]uint32)
    v1 += 32767
    if true {
        v10 := v3[0]
        switch "\n\t" {
        case "\n\t":
            if ((2 <= 0) || (0 <= -598)) {
                if true {
                    v1 += v1
                    var v11 float32 = 0.1
                    var v12 int = (v1 + (^(v1)))
                    _ = v11
                    _ = v12
                } else {
                    var v13 int8 = 127
                    switch v1 {
                    case -2147483648:
                        for fi14 := 0; fi14 < 13; fi14++ {
                            if (!false) {
                                for fi15 := 0; fi15 < 20; fi15++ {
                                    var v16 int = fi14
                                    var v17 int = (-(fi15))
                                    var v18 uint64 = 674
                                    _ = fi15
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                }
                                v19 := GenS2{F0: 1, F1: v1, F2: 2}
                                _ = v19
                            }
                            v20 := []int{}
                            switch v1 {
                            case 42:
                                v21 := GenS3{F0: v1, F1: 0, F2: -2147483648}
                                _ = v21
                            case 4294967295:
                                if ((255 == genFn1(42)) || (!((65535 == 42) && false))) {
                                    var v22 uint16 = 32767
                                    _ = v22
                                } else {
                                    var v23 int8 = v13
                                    _ = v23
                                }
                            case -150:
                                switch 2147483647 {
                                case 113:
                                    var v24 int64 = -9223372036854775808
                                    _ = v24
                                case -1:
                                    _ = len(v20)
                                }
                            default:
                                v13 = (v13 ^ (v13 / ((v13) | 1)))
                            }
                            _ = fi14
                            _ = v20
                        }
                    case -9223372036854775808:
                        var v25 int16 = 32767
                        _ = v25
                    case 127:
                        var v26 int32 = 811
                        _ = v26
                    }
                    switch (v1 - 2147483647) {
                    case 0:
                        v13 = v13
                    case 9223372036854775807:
                        for fi27 := 0; fi27 < 14; fi27++ {
                            v10 -= (^(v10))
                            _ = fi27
                        }
                    case -9223372036854775808:
                        v13 = v13
                    }
                    _ = v13
                }
                fmt.Println(v10, v10)
                var v28 int32 = 0
                _ = v28
            }
        default:
            fmt.Println(v10)
        }
        _ = v10
    }
    var v29 int = (v1 % ((-1) | 1))
    v30 := GenS4{F0: 65535, F1: 788, F2: genFn0((v1 >> ((2) & 63)), 4294967295)}
    v31 := GenS5{F0: -1, F1: 0}
    switch 4294967295 {
    case 127:
        var v32 uint = 498
        _ = v32
    }
    v33 := []int{(v29 % ((128) | 1)), v29, 0, (v1 - v29)}
    v33 = append(v33, 425)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v33
}
