package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 string = "a"
    v2 := []int{-766, 9223372036854775807}
    var v3 uint32 = 2147483647
    var v4 uint64 = 68
    v1 = v1
    if len(v2) > 0 {
        v2[0] = 253
    }
    for fi5 := 0; fi5 < 10; fi5++ {
        v1 = v1
        v4 = (^(v4))
        v6 := GenS1{F0: v1, F1: v4}
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        v2 = append(v2, fi7)
        switch (fi7 / ((fi7) | 1)) {
        case 1:
            v1 = v1
        default:
            var v8 uint8 = 127
            _ = v8
        }
        for fi9 := 0; fi9 < 18; fi9++ {
            v4 = (^(v4))
            var v10 int8 = -128
            var v11 uint16 = 65535
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi7
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        if false {
            var v13 int8 = 42
            if (v4 == (v4 ^ (v4 ^ (v4 | (v4 ^ 18446744073709551615))))) {
                v14 := GenS2{F0: 65535}
                v4 = v4
                _ = v14
            }
            _ = v13
        } else {
            fmt.Println(v1)
            v1 = "a"
        }
        for fi15 := 0; fi15 < 17; fi15++ {
            v16 := GenS2{F0: 65535}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi15
            _ = v16
        }
        v17 := GenS3{F0: (v4 - (v4 % (((^(v4))) | 1)))}
        _ = fi12
        _ = v17
    }
    v3 = (v3 ^ 0)
    if ((!true) && (1e10 > float32(((float64(3.14159) * 1.5) + 0.0)))) {
        var v18 int16 = 42
        _ = v18
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        fmt.Println(v3)
        _ = fi19
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        for fi21 := 0; fi21 < 9; fi21++ {
            fmt.Println(fi20, fi20)
            var v22 int16 = 0
            var v23 float32 = 1.5
            var v24 int16 = v22
            var v25 uint64 = v4
            if (!((323 >= -9223372036854775808) && (!(0 <= 9223372036854775807)))) {
                var v26 int16 = 0
                _ = v26
            } else {
                var v27 int8 = 100
                var v28 int16 = int16(v4)
                _ = v27
                _ = v28
            }
            _ = len(v2)
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
            _ = v25
        }
        _ = fi20
    }
    v4 -= 2
    var v29 uint32 = v3
    fmt.Println(v3)
    v30 := GenS3{F0: (v4 / (((v4 >> ((uint(v4)) & 63))) | 1))}
    fmt.Println(v4)
    var v31 int8 = 100
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
