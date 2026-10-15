package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0() int {
    var v7 uint8 = 42
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 int16 = 498
        fmt.Println(v7, fi8)
        if false {
            for fi10 := 0; fi10 < 18; fi10++ {
                v11 := GenS1{F0: -47, F1: 0}
                switch (fi10 % ((fi10) | 1)) {
                case 42:
                    v12 := []int{(fi8 | fi10), fi10, (fi8 << ((4294967295) & 63)), fi8}
                    _ = v12
                case -2147483648:
                    var v13 uint16 = 0
                    _ = v13
                case 1:
                    if (((((-9223372036854775808 < -1) && true) || false) || (-73 >= -128)) || true) {
                        v14 := GenS1{F0: -9223372036854775808, F1: 100}
                        var v15 int8 = 80
                        v9 += (v9 / (((^(v9))) | 1))
                        _ = v14
                        _ = v15
                    } else {
                        v16 := []int{}
                        fmt.Println(v9, fi10, fi10)
                        _ = v16
                    }
                default:
                    if ((0.1 - 1e10) != 1e10) {
                        v7 -= 2
                    } else {
                        var v17 uint = 2
                        var v18 int64 = 9223372036854775807
                        _ = v17
                        _ = v18
                    }
                }
                v7 += (^(v7))
                _ = fi10
                _ = v11
            }
            v19 := GenS1{F0: 9223372036854775807, F1: 1}
            v20 := []int{}
            _ = v19
            _ = v20
        }
        _ = fi8
        _ = v9
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        v7 += v7
        v7 += (v7 ^ (v7 >> ((18446744073709551615) & 63)))
        v7 -= 2
        _ = fi21
    }
    fmt.Println(v7)
    for fi22 := 0; fi22 < 17; fi22++ {
        for fi23 := 0; fi23 < 5; fi23++ {
            var v24 bool = (!true)
            var v25 uint64 = 229
            _ = fi23
            _ = v24
            _ = v25
        }
        for fi26 := 0; fi26 < 17; fi26++ {
            var v27 uint64 = 330
            fmt.Println(v27, fi22, fi22)
            _ = fi26
            _ = v27
        }
        _ = fi22
    }
    _ = v7
    return 2147483647
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 104
    v3 := []uint32{88, 2, 2147483647, 65535}
    v4 := make(map[int]int)
    v5 := []int{}
    var v6 uint32 = 2
    switch 1 {
    case 1:
        v6 += v6
    }
    v28 := GenS1{F0: -953, F1: 100}
    v6 = v6
    var v29 float64 = (1e10 + float64((float32((0.1 / 1.5)) * float32(1.5))))
    if false {
        var v30 uint64 = 4294967295
        v6 -= 2
        _ = v30
    }
    for fi31 := 0; fi31 < 8; fi31++ {
        fmt.Println(v2, v29)
        var v32 uint32 = 273
        for fi33 := 0; fi33 < 6; fi33++ {
            var v34 int64 = 59
            var v35 string = "a"
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = fi31
        _ = v32
    }
    var v36 int = 128
    if false {
        for fi37 := 0; fi37 < 15; fi37++ {
            v38 := GenS1{F0: 9223372036854775807, F1: 100}
            for fi39 := 0; fi39 < 1; fi39++ {
                v6 -= (v6 | v6)
                v2 = v2
                switch (v36 >> ((42) & 63)) {
                case -725:
                    v40 := GenS2{F0: 2, F1: (v6 >> ((18446744073709551615) & 63))}
                    _ = v40
                default:
                    v6 -= (v6 % (((v6 / (((^(v6))) | 1))) | 1))
                }
                _ = fi39
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi37
            _ = v38
        }
        v3 = append(v3, v6)
    } else {
        v2 = (v2 ^ (-(v2)))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v41 int = v36
        _ = v41
    }
    v6 -= (v6 & (v6 / ((872) | 1)))
    var v42 uint64 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
