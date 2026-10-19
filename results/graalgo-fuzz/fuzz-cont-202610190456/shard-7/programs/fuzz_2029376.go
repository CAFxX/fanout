package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint }
type GenS2 struct { F0 uint; F1 int16; F2 int16 }
func genFn0(p0 int8, p1 uint16) uint32 {
    if true {
        var v7 int8 = p0
        _ = v7
    }
    var v8 string = "a"
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 int16 = -1
        var v11 uint8 = 0
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v8
    return 5
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]int)
    var v4 int16 = -1
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]uint32)
    v12 := []int{185, 32768, -334}
    if true {
        v13 := GenS1{F0: 927, F1: 119, F2: 4294967295}
        v4 += (v4 ^ (v4 >> ((884) & 63)))
        switch 32768 {
        case 127:
            for fi14 := 0; fi14 < 4; fi14++ {
                var v15 int64 = 42
                _ = fi14
                _ = v15
            }
        default:
            fmt.Println(v4)
        }
        _ = v13
    } else {
        fmt.Println(v4)
        if false {
            delete(v2, genFn0(1, 96))
        } else {
            switch 4294967295 {
            case -9223372036854775808:
                fmt.Println(v4)
            case 9223372036854775807:
                if (119 >= -865) {
                    for fi16 := 0; fi16 < 19; fi16++ {
                        switch fi16 {
                        case 323:
                            v4 = v4
                        default:
                            for fi17 := 0; fi17 < 9; fi17++ {
                                if true {
                                    v4 = v4
                                    v18 := GenS1{F0: 42, F1: 107, F2: 491}
                                    _ = v18
                                }
                                _ = len(v12)
                                v4 = v4
                                _ = fi17
                            }
                        }
                        if (-2.25 < (((1.5 - 1e10) - 1.5) - 0.0)) {
                            v4 -= 42
                            var v19 uint8 = 72
                            _ = v19
                        }
                        _ = fi16
                    }
                    v12 = append(v12, -1)
                } else {
                    var v20 uint64 = 4294967295
                    v20 = (v20 / ((18446744073709551615) | 1))
                    for fi21 := 0; fi21 < 18; fi21++ {
                        for fi22 := 0; fi22 < 9; fi22++ {
                            v20 = (v20 >> ((2) & 63))
                            if len(v12) > 0 {
                                v12[0] = (fi22 + fi21)
                            }
                            _ = fi22
                        }
                        _ = fi21
                    }
                    _ = v20
                }
            case 361:
                for fi23 := 0; fi23 < 20; fi23++ {
                    v4 = (v4 % (((v4 + (v4 | (v4 << ((504) & 63))))) | 1))
                    v4 = (v4 >> ((627) & 63))
                    _ = fi23
                }
            }
        }
        var v24 uint8 = 2
        _ = v24
    }
    v4 = v4
    for fi25 := 0; fi25 < 4; fi25++ {
        v4 += 42
        v26 := GenS2{F0: 941, F1: 1000, F2: (v4 + -1)}
        _ = fi25
        _ = v26
    }
    _ = len(v6)
    var v27 float64 = (1.5 / 1.0)
    if len(v12) > 0 {
        v12[0] = 2147483647
    }
    fmt.Println(v4, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
