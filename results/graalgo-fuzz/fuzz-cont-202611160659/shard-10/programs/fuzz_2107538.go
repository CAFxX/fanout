package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint; F1 int }
func genFn0(p0 uint32) int16 {
    if (false || (0 <= 0)) {
        switch 763 {
        case -9223372036854775808:
            v4 := []int{2147483647, 127, -526}
            _ = v4
        }
        switch "" {
        case "":
            var v5 bool = true
            _ = v5
        case "a":
            var v6 uint16 = 255
            _ = v6
        default:
            for fi7 := 0; fi7 < 14; fi7++ {
                v8 := []int{(fi7 >> ((549) & 63)), (fi7 << ((0) & 63)), 32767, fi7}
                var v9 uint16 = 422
                _ = fi7
                _ = v8
                _ = v9
            }
        }
    } else {
        v10 := GenS1{F0: (p0 * p0), F1: 1}
        var v11 float32 = 0.0
        v12 := GenS2{F0: 42, F1: 42, F2: 1}
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := GenS3{F0: 847, F1: 876}
    _ = p0
    _ = v13
    return 761
}
func genFn1(p0 uint16, p1 int64, p2 uint) bool {
    v14 := []int{4294967295, 2147483647, 128, 32768}
    for fi15 := 0; fi15 < 1; fi15++ {
        _ = len(v14)
        _ = fi15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return ((p0 << ((4294967295) & 63)) < uint16(p2))
}

func main() {
    var v1 int = -1
    var v2 float32 = -2.25
    var v3 int16 = 1000
    var v16 string = ((("hello" + ("a" + "hello")) + "\n\t") + "\n\t")
    switch v1 {
    case -1:
        fmt.Println(v2, v1)
    case -43:
        var v17 int16 = (v3 + v3)
        _ = v17
    case 42:
        switch v1 {
        case -987:
            v18 := GenS3{F0: 1, F1: (v1 % ((2147483647) | 1))}
            _ = v18
        case -124:
            var v19 bool = true
            _ = v19
        default:
            for fi20 := 0; fi20 < 9; fi20++ {
                v21 := []int{(fi20 * fi20)}
                var v22 float64 = float64(v2)
                v16 = (v16 + (v16 + ("foo bar" + "\n\t")))
                _ = fi20
                _ = v21
                _ = v22
            }
        }
    }
    v23 := []int{v1, -181, (v1 * -869), (v1 << ((529) & 63))}
    v16 = "\n\t"
    v1 -= v1
    var v24 uint8 = 0
    fmt.Println(v2)
    var v25 int = 9223372036854775807
    var v26 int64 = 42
    var v27 int8 = -32
    if (809 > 165) {
        var v28 uint8 = v24
        _ = v28
    }
    fmt.Println(v2, v2)
    v23 = append(v23, (^(v25)))
    switch (v16 + ("hello" + (v16 + "a"))) {
    case "a":
        fmt.Println(v16, v3, v26)
    }
    var v29 int16 = v3
    for fi30 := 0; fi30 < 2; fi30++ {
        var v31 uint16 = 32767
        switch fi30 {
        case -2147483648:
            v2 -= (float32(1.5) * float32(0.1))
        case 4294967295:
            var v32 uint16 = 65535
            _ = v32
        case -127:
            var v33 int32 = -1
            _ = v33
        }
        _ = fi30
        _ = v31
    }
    v16 = v16
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
}
