package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 int; F1 uint64; F2 uint64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8; F1 int64; F2 uint8 }
func genFn0() uint8 {
    for fi4 := 0; fi4 < 16; fi4++ {
        if (42 <= 174) {
            v5 := GenS1{F0: 65535, F1: -627, F2: 1}
            var v6 float32 = 1.5
            _ = v5
            _ = v6
        } else {
            var v7 uint8 = 2
            _ = v7
        }
        var v8 int32 = -2147483648
        var v9 uint32 = 0
        _ = fi4
        _ = v8
        _ = v9
    }
    var v10 uint64 = 2
    switch 4294967295 {
    case 9223372036854775807:
        var v11 float64 = 1e10
        _ = v11
    case 0:
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 int32 = -538
            if (-307 != 5) {
                if false {
                    if false {
                        fmt.Println(v10)
                        for fi14 := 0; fi14 < 11; fi14++ {
                            v15 := GenS2{F0: 100}
                            v13 = -263
                            v13 = (^(v13))
                            _ = fi14
                            _ = v15
                        }
                    }
                    for fi16 := 0; fi16 < 12; fi16++ {
                        v17 := GenS2{F0: -1}
                        _ = fi16
                        _ = v17
                    }
                    v18 := GenS3{F0: 128, F1: 4294967295}
                    _ = v18
                } else {
                    var v19 int8 = 0
                    v19 = v19
                    _ = v19
                }
            } else {
                for fi20 := 0; fi20 < 8; fi20++ {
                    v21 := GenS2{F0: 0}
                    v22 := GenS4{F0: fi12, F1: v10, F2: v10}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                var v23 float64 = -2.25
                v13 = (v13 >> ((4294967295) & 63))
                _ = v23
            }
            _ = fi12
            _ = v13
        }
    case 32768:
        v24 := GenS5{F0: 1}
        _ = v24
    }
    fmt.Println(v10)
    _ = v10
    return 128
}
func genFn1() int16 {
    v25 := []int{-2147483648, -9223372036854775808, -9223372036854775808, 0}
    for fi26 := 0; fi26 < 7; fi26++ {
        if false {
            var v27 uint8 = 2
            _ = v27
        }
        _ = fi26
    }
    var v28 uint32 = 4294967295
    _ = v25
    _ = v28
    return -77
}

func main() {
    var v1 float32 = (-2.25 - (1.5 + 3.14159))
    var v2 float32 = 0.0
    var v3 uint16 = 32767
    v2 = float32(1e10)
    for fi29 := 0; fi29 < 11; fi29++ {
        v30 := []int{fi29, (fi29 % ((0) | 1)), fi29, -861}
        _ = fi29
        _ = v30
    }
    switch 128 {
    case 4294967295:
        for fi31 := 0; fi31 < 20; fi31++ {
            v32 := GenS6{F0: 1, F1: -9223372036854775808, F2: 255}
            _ = fi31
            _ = v32
        }
    default:
        v33 := GenS1{F0: 255, F1: 356, F2: -1}
        _ = v33
    }
    var v34 float64 = 1.5
    for fi35 := 0; fi35 < 10; fi35++ {
        if false {
            var v36 float64 = float64(v1)
            _ = v36
        }
        v37 := GenS3{F0: genFn0(), F1: 65535}
        v1 = 0.1
        _ = fi35
        _ = v37
    }
    fmt.Println(v2)
    v3 -= v3
    for fi38 := 0; fi38 < 7; fi38++ {
        v39 := GenS4{F0: fi38, F1: 18446744073709551615, F2: 232}
        _ = fi38
        _ = v39
    }
    for fi40 := 0; fi40 < 4; fi40++ {
        var v41 uint = 605
        _ = fi40
        _ = v41
    }
    fmt.Println(v1, v1)
    var v42 int8 = -124
    if false {
        v42 += (v42 % ((100) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v42
}
