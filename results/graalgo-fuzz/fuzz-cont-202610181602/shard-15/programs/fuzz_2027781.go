package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int16 }
func genFn0(p0 uint8) string {
    if (-2.25 <= ((0.0 / (float32(0.1) * 3.14159)) * 1e10)) {
        v7 := GenS1{F0: 1}
        var v8 float32 = 3.14159
        _ = v7
        _ = v8
    } else {
        switch -339 {
        case -9223372036854775808:
            p0 = (p0 - (p0 | 255))
        default:
            for fi9 := 0; fi9 < 15; fi9++ {
                v10 := GenS1{F0: 2147483647}
                for fi11 := 0; fi11 < 4; fi11++ {
                    if false {
                        var v12 int = fi9
                        v12 = -310
                        v13 := []int{128, fi9, fi11, 359}
                        _ = v12
                        _ = v13
                    } else {
                        p0 -= p0
                        var v14 float64 = 1.5
                        _ = v14
                    }
                    _ = fi11
                }
                v15 := []int{}
                _ = fi9
                _ = v10
                _ = v15
            }
        }
        p0 += (p0 + 43)
    }
    v16 := GenS1{F0: 2147483647}
    var v17 int16 = 1
    _ = p0
    _ = v16
    _ = v17
    return "hello"
}
func genFn1(p0 uint16, p1 bool) int32 {
    fmt.Println(p0, p1)
    fmt.Println(p1, p1)
    fmt.Println(p1)
    fmt.Println(p0)
    if (-343 >= 496) {
        for fi18 := 0; fi18 < 1; fi18++ {
            fmt.Println(p0)
            p1 = p1
            p1 = (176 <= 1)
            _ = fi18
        }
    } else {
        var v19 float32 = 1.5
        var v20 float64 = (((1.5 / 3.14159) / 1e10) + (0.0 - (3.14159 + (3.14159 + float64(v19)))))
        p0 += 32767
        _ = v19
        _ = v20
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint32 = 2
    v2 := []int{128}
    v3 := make(map[int]string)
    var v4 bool = true
    var v5 uint8 = 0
    var v6 uint64 = 252
    v21 := v3[9223372036854775807]
    delete(v3, 541)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) > 0 {
        v2[0] = 711
    }
    var v22 int8 = 127
    var v23 string = "hello"
    v1 += v1
    v22 -= int8(v1)
    v22 -= v22
    var v24 uint8 = v5
    if (!v4) {
        var v25 uint = 1
        var v26 int16 = -294
        var v27 string = v23
        _ = v25
        _ = v26
        _ = v27
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v28 := GenS2{F0: 879, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
