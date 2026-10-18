package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 uint16 }
func genFn0(p0 int32, p1 uint8) int64 {
    if (-653 < -798) {
        p1 += p1
    }
    var v4 float32 = 3.14159
    _ = p0
    _ = p1
    _ = v4
    return 465
}
func genFn1(p0 int32, p1 uint8) int64 {
    switch -902 {
    case 1:
        p1 -= (^(p1))
    case 127:
        v5 := []int{-9223372036854775808}
        _ = v5
    default:
        for fi6 := 0; fi6 < 4; fi6++ {
            var v7 int64 = int64(fi6)
            if false {
                v8 := GenS1{F0: 1, F1: 938, F2: -217}
                fmt.Println(p1, p1, p0)
                v9 := GenS2{F0: v7}
                _ = v8
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    }
    fmt.Println(p0, p0)
    v10 := []int{127, 32767, -1}
    p1 += (^(p1))
    _ = p0
    _ = p1
    _ = v10
    return -9223372036854775808
}

func main() {
    var v1 int = -2147483648
    var v2 uint16 = 42
    var v3 uint64 = 0
    switch v1 {
    case 32768:
        switch 0 {
        case 32768:
            var v11 uint8 = 127
            _ = v11
        }
    }
    if true {
        if false {
            v12 := []int{42, (v1 * 127)}
            v3 += 4294967295
            _ = v12
        }
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 float32 = (1.5 + float32((1e10 - (3.14159 / ((1.5 + 0.1) - 3.14159)))))
            _ = fi13
            _ = v14
        }
    }
    var v15 uint = 65535
    if true {
        for fi16 := 0; fi16 < 19; fi16++ {
            v3 = (^(v3))
            var v17 float64 = -2.25
            _ = fi16
            _ = v17
        }
        v2 += (^(v2))
        var v18 float32 = -2.25
        _ = v18
    } else {
        var v19 uint = 65535
        var v20 int64 = 9223372036854775807
        if (!(2147483647 <= 2147483647)) {
            v3 -= 18
        }
        _ = v19
        _ = v20
    }
    if false {
        for fi21 := 0; fi21 < 17; fi21++ {
            v22 := []int{}
            v1 -= (v1 % ((-715) | 1))
            if true {
                v2 -= 680
            } else {
                if false {
                    v2 = uint16(v3)
                    v23 := GenS1{F0: 35, F1: 168, F2: 1}
                    _ = v23
                } else {
                    var v24 int32 = 2147483647
                    v25 := GenS3{F0: 1000, F1: v2}
                    _ = v24
                    _ = v25
                }
            }
            _ = fi21
            _ = v22
        }
        fmt.Println(v3, v2, v3)
    }
    var v26 int32 = -1
    v27 := []int{v1}
    var v28 uint16 = (v2 << (((v15 - v15)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
