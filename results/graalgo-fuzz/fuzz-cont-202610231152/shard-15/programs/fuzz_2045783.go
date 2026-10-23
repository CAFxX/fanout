package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int8 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint }

func main() {
    v1 := []string{}
    var v2 int32 = -1
    var v3 bool = (((float32(0.0) * 0.0) * 3.14159) == 3.14159)
    var v4 int = -2147483648
    v5 := []int{127, (v4 - v4), (v4 | 227), (v4 >> ((255) & 63))}
    var v6 uint = 665
    fmt.Println(v4, v4)
    v5 = append(v5, v4)
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 float64 = float64(float32(-2.25))
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 460, F1: (v4 + (v4 - v4)), F2: -1}
    fmt.Println(v3, v3)
    fmt.Println(v6)
    v2 = -479
    v10 := GenS1{F0: 18446744073709551615, F1: (v4 << (((v6 % ((v6) | 1))) & 63)), F2: 42}
    var v11 float64 = 0.0
    v6 -= (v6 & v6)
    for fi12 := 0; fi12 < 13; fi12++ {
        fmt.Println(v6)
        v11 += v11
        v13 := GenS1{F0: 4294967295, F1: fi12, F2: 122}
        _ = fi12
        _ = v13
    }
    switch (-(v4)) {
    case 127:
        for fi14 := 0; fi14 < 12; fi14++ {
            v15 := GenS1{F0: 295, F1: -2147483648, F2: 1}
            for fi16 := 0; fi16 < 12; fi16++ {
                _ = len(v5)
                var v17 float64 = float64(0.1)
                _ = fi16
                _ = v17
            }
            _ = fi14
            _ = v15
        }
    }
    if true {
        if (255 > 2) {
            fmt.Println(v4, v6, v3)
            var v18 int8 = -128
            if (831 < 681) {
                var v19 string = "hello"
                _ = v19
            } else {
                fmt.Println(v11, v3)
                v18 = (-(v18))
                v20 := GenS2{F0: (9223372036854775807 <= 0), F1: -1}
                _ = v20
            }
            _ = v18
        }
    } else {
        var v21 uint16 = 32767
        v22 := GenS3{F0: v21, F1: 477, F2: v6}
        v6 += 0
        _ = v21
        _ = v22
    }
    _ = len(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
