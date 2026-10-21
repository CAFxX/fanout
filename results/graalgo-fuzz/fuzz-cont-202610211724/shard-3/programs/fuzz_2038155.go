package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint8 }
func genFn0() int32 {
    switch 4294967295 {
    case -9223372036854775808:
        var v5 bool = true
        _ = v5
    case 2147483647:
        var v6 uint64 = 42
        _ = v6
    }
    if true {
        if (-1 >= -20) {
            var v7 uint = 65535
            for fi8 := 0; fi8 < 7; fi8++ {
                v9 := []int{}
                fmt.Println(fi8, fi8)
                _ = fi8
                _ = v9
            }
            var v10 uint = v7
            _ = v10
            _ = v7
        } else {
            var v11 bool = (true || false)
            if v11 {
                v12 := []int{801, 127}
                _ = v12
            }
            for fi13 := 0; fi13 < 6; fi13++ {
                v11 = v11
                v11 = false
                for fi14 := 0; fi14 < 14; fi14++ {
                    var v15 uint16 = 255
                    if (float64((((float32(3.14159) - 0.1) / 1.0) + 0.0)) >= -2.25) {
                        fmt.Println(fi14, v11, v15)
                        var v16 uint64 = 18446744073709551615
                        if (!true) {
                            switch "hello" {
                            case "hello":
                                switch fi13 {
                                case 127:
                                    v17 := GenS1{F0: 42, F1: v15}
                                    _ = v17
                                case 0:
                                    var v18 uint8 = 127
                                    _ = v18
                                }
                            }
                            var v19 int32 = -2147483648
                            var v20 float32 = 0.1
                            _ = v19
                            _ = v20
                        }
                        _ = v16
                    } else {
                        var v21 uint16 = v15
                        fmt.Println(fi14)
                        _ = v21
                    }
                    _ = fi14
                    _ = v15
                }
                _ = fi13
            }
            _ = v11
        }
    } else {
        var v22 int16 = 739
        _ = v22
    }
    return -1
}

func main() {
    var v1 int8 = 127
    v2 := []int{2147483647, 810}
    var v3 float32 = 1.5
    var v4 uint32 = 2
    var v23 int8 = -1
    if (((0.1 * 3.14159) - 1.5) != float64(0.0)) {
        var v24 int32 = 42
        _ = v24
    }
    if false {
        var v25 uint = 4294967295
        v2 = append(v2, 32767)
        v26 := GenS2{F0: 50, F1: v4}
        _ = v25
        _ = v26
    } else {
        fmt.Println(v1, v4)
    }
    var v27 uint16 = 215
    if false {
        for fi28 := 0; fi28 < 9; fi28++ {
            fmt.Println(fi28, v4)
            _ = fi28
        }
        v29 := GenS3{F0: 109}
        _ = v29
    }
    v30 := GenS1{F0: 384, F1: (v27 | (v27 * 65535))}
    for fi31 := 0; fi31 < 4; fi31++ {
        var v32 int8 = -128
        fmt.Println(v1, fi31, v3)
        _ = fi31
        _ = v32
    }
    v3 += (v3 + v3)
    _ = len(v2)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if len(v2) > 0 {
            v2[0] = 32768
        }
    } else {
        v1 = v1
    }
    v1 = 127
    if (v3 <= ((v3 / ((v3 * 3.14159) + ((float32((3.14159 / 1.0)) * (0.0 / v3)) + v3))) - (((1.5 - ((float32(1e10) / (float32(((1e10 - (1e10 / (1e10 / ((1e10 * 1.5) + 3.14159)))) + 1e10)) * (v3 / v3))) + v3)) + (0.1 * v3)) - (1.5 / v3)))) {
        v1 += 44
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v30
    _ = v4
}
