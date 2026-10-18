package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint; F2 uint }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint32) bool {
    p0 = 2147483648
    v6 := []int{-429, 0, 4294967295, 2147483647}
    _ = len(v6)
    if false {
        for fi7 := 0; fi7 < 10; fi7++ {
            switch fi7 {
            case 2147483647:
                _ = len(v6)
            case -9223372036854775808:
                fmt.Println(p0)
            }
            _ = fi7
        }
        fmt.Println(p0)
        p0 = 880
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        fmt.Println(p0)
    }
    var v8 uint = 4294967295
    _ = p0
    _ = v6
    _ = v8
    return ((940 >= 4294967295) && (!(0 != 9223372036854775807)))
}
func genFn1(p0 int32, p1 uint64, p2 uint64) uint16 {
    var v9 int64 = -9223372036854775808
    p0 -= (p0 & p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 0
}

func main() {
    v1 := []int{42, -624}
    v2 := make(map[string]bool)
    v3 := make(map[uint32]int)
    var v4 uint32 = 2147483648
    var v5 float32 = 0.1
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 int32 = 2147483647
        v12 := v2["a"]
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 float32 = v5
    var v14 string = "hello"
    fmt.Println(v4, v4)
    v14 = v14
    if (v4 >= 2147483647) {
        for fi15 := 0; fi15 < 16; fi15++ {
            for fi16 := 0; fi16 < 15; fi16++ {
                fmt.Println(v14, v14, v14)
                fmt.Println(fi16, fi16)
                _ = fi16
            }
            _ = fi15
        }
        if false {
            delete(v3, 2147483648)
        }
        v1 = append(v1, 127)
    } else {
        for fi17 := 0; fi17 < 9; fi17++ {
            if (42 < 42) {
                v18 := GenS1{F0: (42 == 0), F1: -32768}
                v13 = ((float32(float64((float32(((1e10 * 1.5) / 1.5)) / float32(3.14159)))) + 3.14159) + ((-2.25 * (((v13 * v5) / v13) + float32(3.14159))) - v13))
                var v19 float32 = float32(0.1)
                _ = v18
                _ = v19
            }
            var v20 uint = 42
            var v21 int32 = 372
            _ = fi17
            _ = v20
            _ = v21
        }
        if ((!true) && false) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    var v22 int8 = -128
    if false {
        v23 := GenS2{F0: v4, F1: 18446744073709551615, F2: 135}
        if true {
            fmt.Println(v5, v14, v4)
            for fi24 := 0; fi24 < 15; fi24++ {
                var v25 int64 = -9223372036854775808
                v13 = (1e10 * (v13 / v13))
                _ = fi24
                _ = v25
            }
            v13 = (v5 + float32((1e10 * (((0.1 / ((1e10 / (0.1 + 1.5)) - 1e10)) * 1e10) / ((float64(((((3.14159 / 0.1) / v13) - (v5 + -2.25)) / (v13 + v5))) * 0.0) * (float64(((((v5 - v13) * v13) - ((-2.25 / -2.25) - -2.25)) / v5)) * (0.0 - -2.25)))))))
        }
        var v26 int = 32768
        _ = v23
        _ = v26
    }
    var v27 int8 = v22
    switch 42 {
    case 127:
        for fi28 := 0; fi28 < 11; fi28++ {
            if (!false) {
                var v29 bool = true
                _ = v29
            }
            var v30 int32 = 365
            v31 := GenS3{F0: (v22 % (((v27 + (^(v27)))) | 1))}
            _ = fi28
            _ = v30
            _ = v31
        }
    case 4294967295:
        v14 = ""
    case -9223372036854775808:
        for fi32 := 0; fi32 < 8; fi32++ {
            var v33 int8 = -128
            _ = fi32
            _ = v33
        }
    }
    _ = len(v2)
    v34 := GenS1{F0: false, F1: 1000}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
