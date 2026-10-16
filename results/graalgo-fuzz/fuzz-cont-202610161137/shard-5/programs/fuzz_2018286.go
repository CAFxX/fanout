package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int16; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 uint64 }
func genFn0(p0 int64) int8 {
    fmt.Println(p0)
    switch 883 {
    case 4294967295:
        fmt.Println(p0)
    case -2147483648:
        v6 := []int{128, 9223372036854775807, -1, 699}
        _ = v6
    case 211:
        v7 := []int{}
        _ = v7
    default:
        v8 := GenS1{F0: -9223372036854775808, F1: 0}
        _ = v8
    }
    switch 625 {
    case 921:
        v9 := []int{2147483647, 42}
        _ = v9
    }
    switch 341 {
    case 4294967295:
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := GenS1{F0: (fi10 ^ 0), F1: 42}
            var v12 int16 = 528
            v13 := []int{}
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
    case 0:
        var v14 float64 = 1e10
        _ = v14
    case -402:
        if true {
            v15 := GenS1{F0: -9223372036854775808, F1: 901}
            _ = v15
        }
    default:
        v16 := []int{-1, 9223372036854775807}
        _ = v16
    }
    _ = p0
    return 68
}

func main() {
    var v1 int = 32767
    var v2 bool = false
    var v3 int8 = 1
    v4 := make(map[int]int)
    v5 := make(map[uint32]string)
    v17 := []int{}
    for fi18 := 0; fi18 < 1; fi18++ {
        v2 = v2
        for fi19 := 0; fi19 < 20; fi19++ {
            if false {
                if true {
                    var v20 string = "hello"
                    if v2 {
                        var v21 int32 = 42
                        _ = v21
                    } else {
                        if v2 {
                            v20 = ""
                            if (-2.25 > (1.5 / 1e10)) {
                                v22 := v5[2147483648]
                                _ = v22
                            } else {
                                v17 = append(v17, 128)
                                var v23 string = v20
                                _ = len(v5)
                                v3 = 0
                                _ = v23
                            }
                        } else {
                            if true {
                                v24 := GenS1{F0: v1, F1: 744}
                                _ = v24
                            } else {
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                                v25 := GenS2{F0: 1000, F1: 42, F2: 1}
                                _ = v25
                            }
                        }
                    }
                    v26 := GenS3{F0: 42, F1: (fi19 * 4294967295)}
                    _ = v20
                    _ = v26
                } else {
                    var v27 float32 = (1.5 / 3.14159)
                    if v2 {
                        fmt.Println(v2)
                        var v28 int64 = -773
                        var v29 uint32 = 2147483648
                        _ = v28
                        _ = v29
                    }
                    _ = v27
                }
                fmt.Println(v2)
                var v30 float32 = 0.0
                _ = v30
            }
            v4[(fi19 % ((0) | 1))] = (fi19 << ((2) & 63))
            var v31 int = fi18
            if true {
                v32 := GenS4{F0: 857, F1: 42}
                _ = v32
            }
            var v33 float32 = (((float32(3.14159) * 1.5) - (0.1 + 1e10)) * 0.0)
            _ = fi19
            _ = v31
            _ = v33
        }
        _ = fi18
    }
    delete(v5, 2147483648)
    if v2 {
        fmt.Println(v3)
        for fi34 := 0; fi34 < 16; fi34++ {
            for fi35 := 0; fi35 < 1; fi35++ {
                fmt.Println(v2, v1, v1)
                fmt.Println(v2, v2, v2)
                v36 := GenS5{F0: 781}
                _ = fi35
                _ = v36
            }
            _ = fi34
        }
        var v37 uint = 1
        _ = v37
    }
    v1 -= v1
    fmt.Println(v3, v2)
    fmt.Println(v3, v2)
    switch v1 {
    case -1000:
        v38 := GenS3{F0: 312, F1: v1}
        _ = v38
    default:
        if true {
            v39 := GenS6{F0: (-(v3)), F1: 631}
            var v40 uint16 = 32767
            _ = v39
            _ = v40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
