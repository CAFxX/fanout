package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int }
func genFn0(p0 uint8, p1 uint) int8 {
    var v7 int = 1
    v8 := []int{v7, v7, 42}
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := GenS1{F0: 18446744073709551615, F1: 0}
        v11 := GenS2{F0: 382, F1: (!(uint16(p0) != 998))}
        _ = fi9
        _ = v10
        _ = v11
    }
    p0 = p0
    v12 := GenS3{F0: 0, F1: 42, F2: 806}
    _ = p0
    _ = p1
    _ = v12
    _ = v7
    _ = v8
    return 127
}
func genFn1(p0 uint, p1 uint, p2 uint16) int {
    p2 = (^(p2))
    v13 := []int{-2147483648, 800, -663}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 4294967295
}

func main() {
    var v1 int16 = 67
    var v2 uint = 2
    var v3 int = 904
    var v4 bool = false
    var v5 float64 = -2.25
    var v6 bool = v4
    for fi14 := 0; fi14 < 4; fi14++ {
        for fi15 := 0; fi15 < 20; fi15++ {
            switch (v3 ^ 0) {
            case 0:
                v16 := []int{fi14, (fi15 / ((fi15) | 1)), fi15}
                _ = v16
            }
            v17 := GenS2{F0: v2, F1: false}
            _ = fi15
            _ = v17
        }
        v18 := GenS2{F0: v2, F1: v4}
        var v19 uint32 = 287
        _ = fi14
        _ = v18
        _ = v19
    }
    if (!(v4 && v6)) {
        fmt.Println(v4)
        v2 = v2
        fmt.Println(v2)
    }
    v3 += -781
    if (42 == 556) {
        v1 -= (^(v1))
        var v20 uint = (v2 * v2)
        var v21 bool = (42 != 42)
        _ = v20
        _ = v21
    }
    v2 += 163
    var v22 uint16 = 255
    if ((v3 ^ 2147483647) <= (v3 | genFn1(1, 255, v22))) {
        for fi23 := 0; fi23 < 14; fi23++ {
            v6 = ((v1 | v1) > v1)
            _ = fi23
        }
        v24 := GenS3{F0: 147, F1: 2, F2: v22}
        v2 = 1
        _ = v24
    } else {
        v1 = (v1 & v1)
        for fi25 := 0; fi25 < 18; fi25++ {
            var v26 uint32 = 970
            var v27 int16 = -389
            _ = fi25
            _ = v26
            _ = v27
        }
        switch (v3 - 0) {
        case 785:
            switch v3 {
            case 32768:
                fmt.Println(v3, v4)
            default:
                if (100 != 1) {
                    if v6 {
                        fmt.Println(v2)
                        fmt.Println(v6)
                    } else {
                        var v28 uint64 = 42
                        v29 := []int{-627}
                        v2 += (v2 * v2)
                        _ = v28
                        _ = v29
                    }
                    fmt.Println(v2, v6)
                    v22 -= (v22 % (((^(v22))) | 1))
                }
            }
        case 305:
            v22 = v22
        default:
            fmt.Println(v3, v3, v5)
        }
    }
    v5 += 3.14159
    for fi30 := 0; fi30 < 3; fi30++ {
        v31 := GenS2{F0: 255, F1: (127 <= -9)}
        _ = fi30
        _ = v31
    }
    if ((!(true || v4)) || v6) {
        switch 128 {
        case -9223372036854775808:
            var v32 int16 = (v1 % (((v1 >> (((v2 & 42)) & 63))) | 1))
            _ = v32
        case 128:
            for fi33 := 0; fi33 < 11; fi33++ {
                v34 := GenS4{F0: (^(v3))}
                _ = fi33
                _ = v34
            }
        case -254:
            v5 -= ((float64(1.5) * v5) / float64(1.5))
        default:
            v35 := GenS3{F0: 4294967295, F1: (v22 << ((v2) & 63)), F2: v22}
            _ = v35
        }
        v3 = v3
        var v36 int8 = -71
        _ = v36
    } else {
        var v37 float32 = 0.0
        fmt.Println(v6)
        fmt.Println(v5, v1)
        _ = v37
    }
    switch -83 {
    case 2147483647:
        v4 = false
    case 32767:
        for fi38 := 0; fi38 < 1; fi38++ {
            var v39 uint16 = v22
            var v40 uint16 = v22
            v41 := []int{0, fi38, fi38, fi38}
            _ = fi38
            _ = v39
            _ = v40
            _ = v41
        }
    case 326:
        var v42 bool = (((-(v1)) != (v1 & (v1 + (v1 ^ v1)))) && false)
        _ = v42
    default:
        v4 = (v3 <= (v3 - v3))
    }
    v3 = (v3 | (^(v3)))
    fmt.Println(v6)
    v43 := GenS1{F0: 502, F1: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
