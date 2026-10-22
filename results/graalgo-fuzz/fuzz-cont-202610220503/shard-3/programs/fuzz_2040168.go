package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint64; F1 uint32 }
func genFn0(p0 uint) int {
    var v5 uint8 = 128
    v6 := GenS1{F0: 0, F1: (0 >= 1)}
    _ = p0
    _ = v5
    _ = v6
    return 32767
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 1
    var v3 float64 = 3.14159
    var v4 uint8 = 255
    if true {
        var v7 int = -1
        var v8 string = ""
        if (v4 != (v4 << ((4294967295) & 63))) {
            for fi9 := 0; fi9 < 12; fi9++ {
                v10 := GenS1{F0: 59, F1: (3.14159 >= 0.0)}
                v7 = 9223372036854775807
                v1[v2] = genFn0(0)
                _ = fi9
                _ = v10
            }
        }
        _ = v7
        _ = v8
    }
    v11 := GenS1{F0: 765, F1: false}
    v4 = (v4 + v4)
    var v12 uint16 = 783
    v4 = v4
    v12 = 32767
    if false {
        var v13 float32 = ((1e10 + 1e10) + float32((float64(1.5) * (v3 - (((0.0 * float64(0.0)) - v3) - v3)))))
        v14 := GenS1{F0: int32(v4), F1: false}
        if false {
            var v15 int64 = 441
            if ((false && (false && (true && ((v13 == 3.14159) && (genFn0(65535) != -143))))) && false) {
                fmt.Println(v4)
            }
            _ = v15
        } else {
            _ = len(v1)
            switch 4294967295 {
            case 418:
                var v16 int32 = 42
                _ = v16
            default:
                if (-1 == 1) {
                    var v17 int8 = 91
                    _ = v17
                } else {
                    if (255 < 0) {
                        for fi18 := 0; fi18 < 6; fi18++ {
                            var v19 int8 = -18
                            var v20 uint32 = v2
                            v3 = -2.25
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        if false {
                            v1[(v2 ^ (v2 - (v2 ^ (v2 | (v2 + 0)))))] = 4294967295
                            var v21 float64 = (1.5 - 1.5)
                            var v22 int = -2147483648
                            _ = v21
                            _ = v22
                        } else {
                            switch 134 {
                            case 2147483647:
                                var v23 string = "hello"
                                _ = v23
                            default:
                                var v24 uint = 0
                                _ = v24
                            }
                            v25 := GenS1{F0: -982, F1: false}
                            if (32767 > 32767) {
                                v3 += (((v3 * 0.0) - ((1.5 + -2.25) - 0.1)) * v3)
                                switch int(v12) {
                                case 9223372036854775807:
                                    v2 += 103
                                }
                            } else {
                                for fi26 := 0; fi26 < 15; fi26++ {
                                    v27 := GenS1{F0: 1, F1: false}
                                    if (v2 == v2) {
                                        var v28 bool = (!((-35 >= 456) || (!true)))
                                        v12 -= (v12 << ((53) & 63))
                                        _ = v28
                                    }
                                    _ = fi26
                                    _ = v27
                                }
                            }
                            _ = v25
                        }
                    } else {
                        v3 += v3
                        v29 := GenS2{F0: 708, F1: 2}
                        _ = v29
                    }
                    if false {
                        var v30 int8 = 1
                        v31 := []int{}
                        fmt.Println(v12, v4)
                        _ = v30
                        _ = v31
                    } else {
                        for fi32 := 0; fi32 < 7; fi32++ {
                            var v33 int = (fi32 * fi32)
                            _ = fi32
                            _ = v33
                        }
                        var v34 int = -865
                        _ = v34
                    }
                }
            }
        }
        _ = v13
        _ = v14
    } else {
        v2 += (v2 / ((uint32(v12)) | 1))
    }
    var v35 int = 0
    if true {
        for fi36 := 0; fi36 < 9; fi36++ {
            fmt.Println(v4, fi36, v3)
            _ = fi36
        }
        _ = len(v1)
    } else {
        if false {
            v3 += (v3 - v3)
            v12 = (^(v12))
            var v37 int16 = -494
            _ = v37
        }
        if ((v4 == (v4 / (((v4 * 255)) | 1))) || true) {
            var v38 uint8 = v4
            switch genFn0(65535) {
            case 42:
                var v39 int = (v35 ^ v35)
                _ = v39
            }
            _ = v38
        }
    }
    v4 = 255
    for fi40 := 0; fi40 < 14; fi40++ {
        v35 -= (v35 * v35)
        v41 := []int{32768, (v35 & (v35 % ((fi40) | 1)))}
        for fi42 := 0; fi42 < 4; fi42++ {
            var v43 float32 = (-2.25 / float32((-2.25 - v3)))
            _ = fi42
            _ = v43
        }
        _ = fi40
        _ = v41
    }
    v44 := GenS3{F0: 1, F1: (v2 ^ (^(v2)))}
    var v45 uint32 = 1
    var v46 int32 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v44
    _ = v45
    _ = v46
}
