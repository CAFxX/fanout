package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int64 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 3
    v3 := make(map[int]bool)
    v4 := []uint32{876, 784, 2147483648}
    var v5 int8 = (v2 << ((920) & 63))
    var v6 bool = true
    var v7 uint32 = 2
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 uint = 42
        _ = fi8
        _ = v9
    }
    if false {
        if (v6 || ((938 < 128) || false)) {
            var v10 uint8 = 67
            if ((v10 | 94) != (^(v10))) {
                var v11 uint16 = 206
                if (!v6) {
                    v6 = (42 < 42)
                }
                var v12 uint32 = 2147483648
                _ = v11
                _ = v12
            }
            _ = v10
        }
        if false {
            v2 = (v2 & (v5 / ((v5) | 1)))
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v13 := GenS1{F0: 958, F1: -128, F2: 0}
    for fi14 := 0; fi14 < 4; fi14++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v15 int64 = 958
        _ = fi14
        _ = v15
    }
    var v16 int = -331
    v2 -= (v2 | v2)
    v17 := GenS1{F0: uint(v7), F1: (v2 >> ((403) & 63)), F2: -1}
    var v18 uint16 = 255
    if true {
        for fi19 := 0; fi19 < 19; fi19++ {
            v6 = v6
            _ = fi19
        }
        var v20 uint16 = 131
        fmt.Println(v18)
        _ = v20
    } else {
        var v21 int16 = -260
        v2 += 100
        v18 = 698
        _ = v21
    }
    v6 = ((float32(float64(1e10)) < (-2.25 * (3.14159 / 1e10))) || (100 < v5))
    for fi22 := 0; fi22 < 11; fi22++ {
        for fi23 := 0; fi23 < 14; fi23++ {
            v4 = append(v4, (v7 % ((v7) | 1)))
            var v24 uint = 255
            var v25 string = ""
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi22
    }
    var v26 float32 = (3.14159 / 1e10)
    var v27 uint64 = 18446744073709551615
    if false {
        v28 := GenS1{F0: 42, F1: v2, F2: -295}
        fmt.Println(v16)
        v29 := GenS2{F0: 14}
        _ = v28
        _ = v29
    } else {
        for fi30 := 0; fi30 < 11; fi30++ {
            if false {
                for fi31 := 0; fi31 < 19; fi31++ {
                    var v32 uint = 255
                    switch v16 {
                    case 32767:
                        delete(v3, fi30)
                    case 2147483647:
                        v33 := GenS2{F0: 127}
                        _ = v33
                    default:
                        v34 := GenS1{F0: (v32 + (^(v32))), F1: (v5 / ((v2) | 1)), F2: 9223372036854775807}
                        _ = v34
                    }
                    switch v16 {
                    case 32767:
                        fmt.Println(v26, v27)
                    default:
                        if v6 {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            delete(v3, fi30)
                        } else {
                            if len(v4) > 0 {
                                v4[0] = v7
                            }
                        }
                    }
                    _ = fi31
                    _ = v32
                }
            }
            var v35 float32 = (v26 - v26)
            _ = fi30
            _ = v35
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
