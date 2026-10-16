package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint8 }

func main() {
    v1 := []int{}
    v2 := []int{-2147483648, 32767, -319}
    v3 := []uint32{}
    var v4 int32 = 1
    fmt.Println(v4)
    v5 := GenS1{F0: 65535, F1: 0, F2: 0}
    v4 = v4
    if true {
        for fi6 := 0; fi6 < 1; fi6++ {
            if false {
                fmt.Println(v4)
                var v7 uint32 = 1
                switch (^(fi6)) {
                case 32767:
                    for fi8 := 0; fi8 < 15; fi8++ {
                        var v9 uint32 = v7
                        v4 = -2147483648
                        _ = fi8
                        _ = v9
                    }
                default:
                    if (65535 == 2) {
                        var v10 float64 = (float64((3.14159 * ((1.5 - -2.25) * float32(((0.0 * -2.25) * ((-2.25 * 0.0) - ((0.1 / 1.0) * ((1e10 / (0.1 + (float64((1.5 / -2.25)) + (-2.25 * 1.5)))) - 0.1)))))))) * -2.25)
                        for fi11 := 0; fi11 < 12; fi11++ {
                            if (32767 > 857) {
                                var v12 uint = 65535
                                var v13 uint = v12
                                v14 := GenS1{F0: v7, F1: (v7 | (^(v7))), F2: 255}
                                _ = v12
                                _ = v13
                                _ = v14
                            } else {
                                v7 += 0
                                switch (^(fi11)) {
                                case 4294967295:
                                    v7 -= 772
                                case -193:
                                    switch "" {
                                    case "foo bar":
                                        _ = len(v3)
                                    }
                                case -2147483648:
                                    _ = len(v3)
                                }
                                var v15 float32 = 3.14159
                                _ = v15
                            }
                            var v16 int = (fi6 >> ((375) & 63))
                            var v17 float32 = (0.0 / -2.25)
                            v17 += (v17 / v17)
                            v7 += v7
                            _ = fi11
                            _ = v16
                            _ = v17
                        }
                        _ = v10
                    }
                }
                _ = v7
            }
            v4 = (v4 + (v4 % ((int32(fi6)) | 1)))
            _ = fi6
        }
        var v18 float32 = (-2.25 * (0.0 * float32(float64(0.1))))
        _ = v18
    } else {
        switch "hello" {
        case "\n\t":
            var v19 int = 127
            _ = v19
        case "a":
            v4 = v4
        case "hello":
            v4 -= v4
        }
    }
    var v20 int = -9223372036854775808
    var v21 string = ("foo bar" + ("" + "hello"))
    v20 = v20
    fmt.Println(v4, v20)
    for fi22 := 0; fi22 < 18; fi22++ {
        var v23 string = "hello"
        if len(v3) > 0 {
            v3[0] = 720
        }
        var v24 float32 = 0.0
        _ = fi22
        _ = v23
        _ = v24
    }
    v4 -= (v4 - (v4 * v4))
    v25 := GenS1{F0: 745, F1: 1, F2: 42}
    fmt.Println(v4)
    var v26 float64 = 0.0
    if (1 == v4) {
        v27 := GenS1{F0: 4294967295, F1: 4294967295, F2: 255}
        v4 += -958
        v20 -= v20
        _ = v27
    }
    switch v20 {
    case -79:
        for fi28 := 0; fi28 < 6; fi28++ {
            if len(v1) > 0 {
                v1[0] = v20
            }
            var v29 uint64 = 42
            _ = fi28
            _ = v29
        }
    default:
        v4 = v4
    }
    v4 += v4
    v21 = v21
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
