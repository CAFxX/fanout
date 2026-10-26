package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
func genFn0(p0 int16, p1 uint) uint {
    if (!true) {
        v4 := []int{32768}
        _ = v4
    } else {
        if (255 >= p1) {
            var v5 int64 = -580
            _ = v5
        } else {
            p1 += p1
        }
    }
    var v6 uint16 = 32767
    var v7 uint8 = 255
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (p1 + uint(v6))
}
func genFn1(p0 bool, p1 uint8, p2 bool) int16 {
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 float32 = (-2.25 - 0.0)
        v10 := GenS1{F0: 0, F1: -1}
        _ = fi8
        _ = v10
        _ = v9
    }
    p0 = false
    var v11 uint8 = (p1 / ((127) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 42
}

func main() {
    var v1 uint = 4294967295
    var v2 uint = 65535
    v3 := make(map[uint32]bool)
    v1 -= (^(v2))
    v3[239] = true
    switch -2147483648 {
    case -9223372036854775808:
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 uint64 = 2
            v13 = v13
            _ = fi12
            _ = v13
        }
    }
    var v14 uint64 = 514
    for fi15 := 0; fi15 < 4; fi15++ {
        v1 = (v2 * (v1 << ((2) & 63)))
        _ = fi15
    }
    v14 -= (v14 << (((v2 / ((genFn0(-1, genFn0(32767, genFn0(32767, (v1 & 65535))))) | 1))) & 63))
    for fi16 := 0; fi16 < 8; fi16++ {
        fmt.Println(fi16, v14)
        _ = fi16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 float32 = 0.0
        if false {
            switch (fi17 - fi17) {
            case -951:
                v19 := []int{}
                _ = v19
            case 32767:
                var v20 uint32 = 0
                _ = v20
            case 159:
                v14 += v14
            default:
                fmt.Println(v14)
            }
            var v21 string = "\n\t"
            _ = v21
        }
        var v22 int16 = genFn1(true, 255, false)
        _ = fi17
        _ = v18
        _ = v22
    }
    fmt.Println(v1, v14)
    var v23 int8 = 1
    v3[1] = true
    v24 := []int{}
    if ((-2147483648 >= 4294967295) && (0.1 >= 1e10)) {
        v1 = 42
    } else {
        var v25 int = 0
        v24 = append(v24, -455)
        _ = v25
    }
    v14 = (v14 % ((v14) | 1))
    delete(v3, 2147483648)
    if ((float32(3.14159) / 1e10) != (3.14159 / 0.1)) {
        if len(v24) >= 2 {
            v24 = v24[1:len(v24)]
        }
    }
    v26 := GenS1{F0: 1, F1: -774}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v3
}
