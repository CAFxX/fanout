package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int16 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint16; F1 int; F2 int8 }
func genFn0(p0 uint64, p1 int32, p2 string) int {
    if (4294967295 >= 32768) {
        fmt.Println(p2)
        v6 := GenS1{F0: 236, F1: -1, F2: -102}
        _ = v6
    } else {
        v7 := []int{}
        _ = len(v7)
        p2 = (p2 + ((p2 + ("foo bar" + p2)) + p2))
        _ = v7
    }
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 32767
}
func genFn1(p0 uint64, p1 uint16) int8 {
    if false {
        p1 = (p1 >> ((595) & 63))
    } else {
        var v9 uint = 4294967295
        fmt.Println(p0, p0, v9)
        _ = v9
    }
    p1 -= 2
    var v10 int16 = 60
    var v11 float32 = (float32(((3.14159 - 1.5) / -2.25)) + ((0.0 + 3.14159) + ((1.5 / ((1.5 / 1.0) + 3.14159)) / 1.0)))
    var v12 string = "\n\t"
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return 1
}

func main() {
    v1 := []int{-9223372036854775808, 32767, -9223372036854775808, 995}
    var v2 float32 = ((0.0 + ((-2.25 * 3.14159) - -2.25)) + 1.5)
    var v3 float32 = 0.1
    var v4 bool = (-2.25 < 0.1)
    v5 := make(map[string]string)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3, v3)
    var v13 string = "hello"
    var v14 uint = 745
    fmt.Println(v4, v14)
    for fi15 := 0; fi15 < 16; fi15++ {
        for fi16 := 0; fi16 < 12; fi16++ {
            v3 = (float32(0.0) + (float32(-2.25) / v2))
            var v17 float64 = 1.5
            v1 = append(v1, (-(fi15)))
            var v18 int64 = -432
            v19 := GenS1{F0: v14, F1: -2147483648, F2: 656}
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
        var v20 uint16 = 0
        _ = fi15
        _ = v20
    }
    v2 -= ((-2.25 * 0.0) + v2)
    v13 = v13
    fmt.Println(v2, v13, v3)
    if v4 {
        var v21 string = v13
        if (v14 < 2) {
            v22 := v5[v13]
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v22
        } else {
            v23 := GenS1{F0: (v14 / ((v14) | 1)), F1: 1, F2: -338}
            v13 = v21
            _ = v23
        }
        _ = v21
    } else {
        v14 = (v14 - v14)
        fmt.Println(v13, v14, v2)
        var v24 uint8 = 0
        _ = v24
    }
    v2 -= float32(1.5)
    v25 := GenS2{F0: 1, F1: v14}
    for fi26 := 0; fi26 < 10; fi26++ {
        for fi27 := 0; fi27 < 13; fi27++ {
            v28 := GenS3{F0: 2, F1: 216, F2: 1}
            _ = fi27
            _ = v28
        }
        for fi29 := 0; fi29 < 16; fi29++ {
            var v30 int16 = -32768
            _ = fi29
            _ = v30
        }
        fmt.Println(v2, v4)
        _ = fi26
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        var v32 int64 = 9223372036854775807
        v1 = append(v1, 42)
        for fi33 := 0; fi33 < 8; fi33++ {
            if false {
                _ = len(v5)
                if (42 <= 0) {
                    var v34 int = fi31
                    _ = v34
                } else {
                    fmt.Println(v14, fi33)
                    v5[v13] = v13
                    v32 = (v32 + v32)
                }
            }
            v2 += v2
            v35 := GenS1{F0: (v14 | (v14 & v14)), F1: -1, F2: int16(v14)}
            _ = fi33
            _ = v35
        }
        _ = fi31
        _ = v32
    }
    for fi36 := 0; fi36 < 14; fi36++ {
        switch (fi36 ^ (fi36 >> ((v14) & 63))) {
        case -840:
            var v37 float64 = 0.0
            _ = v37
        case -1:
            var v38 int = fi36
            _ = v38
        case 801:
            var v39 int = (fi36 ^ genFn0(18446744073709551615, 2147483647, v13))
            _ = v39
        }
        fmt.Println(v14)
        _ = fi36
    }
    var v40 int8 = -86
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
