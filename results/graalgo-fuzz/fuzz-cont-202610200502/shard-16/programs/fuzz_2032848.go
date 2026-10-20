package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint }
type GenS2 struct { F0 int32; F1 int; F2 bool }
func genFn0(p0 int64, p1 string, p2 bool) uint {
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 int16 = int16(p0)
        _ = fi5
        _ = v6
    }
    var v7 uint32 = 0
    p1 = p1
    for fi8 := 0; fi8 < 14; fi8++ {
        fmt.Println(p0, v7)
        var v9 uint16 = 0
        _ = fi8
        _ = v9
    }
    v10 := []int{0, 2147483647, 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return 2
}

func main() {
    var v1 int8 = 0
    v2 := make(map[int]string)
    var v3 int64 = -1
    v4 := []uint32{744}
    v1 = v1
    fmt.Println(v3)
    v1 += (^(v1))
    var v11 uint32 = 732
    var v12 string = "foo bar"
    var v13 int8 = v1
    for fi14 := 0; fi14 < 7; fi14++ {
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 int64 = (v3 & (-(v3)))
            if (true || (0.0 > -2.25)) {
                v4 = append(v4, (v11 | v11))
                if (-439 == 1) {
                    if ((v11 & (^(v11))) != (v11 << ((4294967295) & 63))) {
                        if (1e10 < (float64(-2.25) - 1e10)) {
                            if len(v4) > 0 {
                                v4[0] = v11
                            }
                            var v17 bool = true
                            _ = v17
                        }
                    }
                } else {
                    v18 := GenS1{F0: v12, F1: -1, F2: 856}
                    if false {
                        switch fi14 {
                        case 0:
                            v3 += v16
                        default:
                            var v19 int16 = 0
                            _ = v19
                        }
                        var v20 uint32 = v11
                        _ = v20
                    } else {
                        fmt.Println(fi15, fi15)
                        var v21 uint = 65535
                        v1 = (v13 << (((v21 | genFn0(v16, (v12 + (v12 + ("" + v12))), false))) & 63))
                        var v22 int64 = v16
                        _ = v21
                        _ = v22
                    }
                    _ = v18
                }
                if (true || (!true)) {
                    if false {
                        switch (v12 + v12) {
                        case "":
                            if (1 > (^(v1))) {
                                delete(v2, fi14)
                            } else {
                                fmt.Println(fi15)
                                var v23 uint32 = (v11 / (((^(v11))) | 1))
                                v24 := v2[(fi15 % (((fi14 / (((fi14 ^ 823)) | 1))) | 1))]
                                _ = v23
                                _ = v24
                            }
                            var v25 uint = genFn0(v3, (("a" + "\n\t") + ((v12 + v12) + v12)), false)
                            _ = v25
                        default:
                            fmt.Println(v1)
                            v26 := GenS2{F0: 42, F1: fi14, F2: (152 >= 10)}
                            v13 -= (v1 / ((v1) | 1))
                            v12 = (v12 + v12)
                            var v27 int64 = 665
                            var v28 string = v12
                            delete(v2, fi15)
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                        if true {
                            fmt.Println(v13, fi15, fi15)
                            var v29 int16 = 1000
                            fmt.Println(v29, v1)
                            _ = v29
                        } else {
                            v30 := GenS2{F0: -663, F1: int(v3), F2: true}
                            var v31 float32 = -2.25
                            var v32 int8 = (v1 - (v1 ^ (^(v13))))
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                    }
                    v11 = (v11 - (v11 ^ (v11 << ((42) & 63))))
                }
            } else {
                switch v12 {
                case "a":
                    v13 = v1
                case "\n\t":
                    v1 -= v13
                }
                var v33 uint32 = (v11 * (v11 / ((2147483648) | 1)))
                v12 = v12
                _ = v33
            }
            var v34 int16 = -985
            if (637 < 45) {
                _ = len(v2)
            }
            fmt.Println(fi15, v1, v16)
            _ = fi15
            _ = v16
            _ = v34
        }
        _ = fi14
    }
    switch "foo bar" {
    case "hello":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
