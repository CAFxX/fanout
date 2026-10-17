package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int32; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint16; F1 int16 }
func genFn0(p0 int64, p1 int, p2 string) int {
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 int = 1
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    p1 += p1
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS1{F0: 779, F1: p0}
        v11 := GenS1{F0: 18446744073709551615, F1: p0}
        _ = fi9
        _ = v10
        _ = v11
    }
    if false {
        v12 := GenS1{F0: 588, F1: -9223372036854775808}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return 9223372036854775807
}
func genFn1(p0 int8) int8 {
    v13 := GenS2{F0: 2147483647, F1: 577, F2: 313}
    p0 -= (p0 & (p0 % ((p0) | 1)))
    p0 = p0
    var v14 int = 32767
    _ = p0
    _ = v13
    _ = v14
    return 127
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = (((3.14159 - 0.0) * 3.14159) * ((3.14159 + 0.0) * (3.14159 * 3.14159)))
    var v3 uint = 255
    var v4 int32 = 987
    v5 := make(map[string]uint32)
    if (0 > uint8(v3)) {
        v3 = v3
    } else {
        v15 := GenS2{F0: int32(v3), F1: v3, F2: 2147483647}
        v16 := GenS1{F0: v3, F1: -9223372036854775808}
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: 1, F1: 1}
    v18 := GenS3{F0: 255, F1: 42}
    fmt.Println(v4, v3, v4)
    v2 += v2
    switch 217 {
    case 32767:
        v3 = 4294967295
    case 0:
        if ((-2.25 + -2.25) < (-2.25 * 1e10)) {
            v19 := []int{9223372036854775807, 32768}
            v3 -= 65535
            v3 = v3
            _ = v19
        } else {
            v2 = ((((v2 / float64(((1e10 * ((0.0 / 1e10) + 1e10)) - 0.1))) / 1e10) / (3.14159 / ((v2 - 1e10) * v2))) * 0.1)
        }
    }
    v2 = v2
    for fi20 := 0; fi20 < 12; fi20++ {
        v3 = v3
        v3 = v3
        _ = fi20
    }
    fmt.Println(v2)
    for fi21 := 0; fi21 < 10; fi21++ {
        var v22 float64 = (float64(3.14159) / 1.0)
        _ = fi21
        _ = v22
    }
    var v23 uint32 = 921
    v24 := GenS3{F0: 255, F1: -32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
