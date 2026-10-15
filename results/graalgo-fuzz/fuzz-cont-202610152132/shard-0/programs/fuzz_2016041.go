package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 uint16 }
func genFn0(p0 int) uint64 {
    if false {
        for fi6 := 0; fi6 < 20; fi6++ {
            fmt.Println(p0)
            _ = fi6
        }
    } else {
        v7 := GenS1{F0: 354}
        p0 = p0
        p0 += -2147483648
        _ = v7
    }
    p0 = p0
    var v8 float32 = (0.0 + (1e10 / -2.25))
    if (1 >= 100) {
        var v9 int8 = 127
        v10 := []int{p0, p0, (-(p0))}
        v11 := GenS1{F0: p0}
        _ = v10
        _ = v11
        _ = v9
    }
    v12 := GenS1{F0: p0}
    _ = p0
    _ = v12
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 2147483647
    v3 := make(map[string]int)
    v4 := make(map[uint32]uint32)
    v5 := []uint32{}
    v2 = (v2 / ((42) | 1))
    _ = len(v1)
    fmt.Println(v2)
    v13 := v1[-2147483648]
    v4[(v13 >> ((2) & 63))] = 2147483648
    switch 42 {
    case 42:
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 uint = 18446744073709551615
            _ = len(v1)
            v15 = v15
            _ = fi14
            _ = v15
        }
    default:
        if (false || ("\n\t" != "\n\t")) {
            var v16 uint16 = 32767
            _ = v16
        } else {
            if (!false) {
                v2 = (v2 >> ((4294967295) & 63))
                if (int64(v13) >= -141) {
                    v17 := GenS1{F0: -9223372036854775808}
                    _ = v17
                }
                fmt.Println(v13, v13)
            }
        }
    }
    v2 += (v13 ^ v2)
    if false {
        v2 = (v2 << ((0) & 63))
        v13 += (^(v13))
    } else {
        var v18 int8 = 43
        _ = v18
    }
    var v19 float64 = (((0.1 - float64(0.0)) + 0.1) * (0.0 / float64(1.5)))
    v5 = append(v5, (v2 >> ((1) & 63)))
    var v20 uint = 829
    if false {
        var v21 int32 = -2147483648
        switch -300 {
        case -2147483648:
            v22 := GenS2{F0: (!true)}
            _ = v22
        case 292:
            for fi23 := 0; fi23 < 9; fi23++ {
                var v24 float32 = (((-2.25 + ((0.1 * ((3.14159 - (1e10 - float32(0.1))) * 1.5)) / (0.1 - 3.14159))) / 1e10) + -2.25)
                _ = fi23
                _ = v24
            }
        case 875:
            if true {
                for fi25 := 0; fi25 < 7; fi25++ {
                    switch "" {
                    case "a":
                        for fi26 := 0; fi26 < 13; fi26++ {
                            if true {
                                v27 := GenS2{F0: false}
                                delete(v4, v2)
                                var v28 float64 = (v19 + v19)
                                _ = v27
                                _ = v28
                            } else {
                                var v29 int16 = -1
                                switch (fi25 ^ 32768) {
                                case 9223372036854775807:
                                    v19 = (3.14159 / 1e10)
                                default:
                                    _ = len(v1)
                                }
                                _ = v29
                            }
                            _ = fi26
                        }
                    case "\n\t":
                        v30 := GenS3{F0: v20, F1: 754}
                        _ = v30
                    }
                    if len(v5) > 0 {
                        v5[0] = v13
                    }
                    if true {
                        v21 = v21
                        switch "" {
                        case "":
                            fmt.Println(v2, v21, v20)
                        }
                    }
                    _ = fi25
                }
                for fi31 := 0; fi31 < 10; fi31++ {
                    v19 = -2.25
                    _ = fi31
                }
                var v32 uint64 = 42
                _ = v32
            } else {
                fmt.Println(v13)
                var v33 uint8 = 1
                v19 -= v19
                _ = v33
            }
        }
        var v34 float64 = float64(1.5)
        _ = v21
        _ = v34
    }
    delete(v4, v13)
    _ = len(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
