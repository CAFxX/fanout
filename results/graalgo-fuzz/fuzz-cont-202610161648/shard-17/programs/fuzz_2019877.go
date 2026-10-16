package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 bool }
type GenS2 struct { F0 uint; F1 uint; F2 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8 }
func genFn0(p0 int8) int32 {
    var v6 float32 = -2.25
    var v7 bool = true
    p0 = -1
    fmt.Println(v6, p0, v7)
    fmt.Println(v7)
    _ = p0
    _ = v6
    _ = v7
    return 2147483647
}
func genFn1(p0 int64, p1 uint16) uint {
    p0 = p0
    var v8 uint32 = 65535
    v9 := GenS1{F0: 4294967295, F1: 985, F2: false}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    var v1 float32 = 0.1
    var v2 uint32 = 292
    var v3 uint32 = v2
    var v4 uint32 = 2147483648
    var v5 int32 = 0
    var v10 float32 = float32(0.0)
    var v11 int16 = 42
    for fi12 := 0; fi12 < 10; fi12++ {
        if false {
            v3 += 0
            v2 -= 2147483648
        } else {
            v13 := GenS1{F0: v3, F1: v2, F2: true}
            _ = v13
        }
        v14 := []int{fi12, fi12, (fi12 % ((-9223372036854775808) | 1)), fi12}
        v10 = (v10 + (((float32(1e10) * 3.14159) - (float32(0.1) + v1)) / 1.0))
        _ = fi12
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        if false {
            if false {
                v16 := GenS2{F0: 370, F1: 255, F2: 18446744073709551615}
                v17 := GenS1{F0: (v4 + v2), F1: (v2 - 65535), F2: true}
                var v18 int8 = 127
                _ = v16
                _ = v17
                _ = v18
            } else {
                for fi19 := 0; fi19 < 8; fi19++ {
                    v20 := []int{fi15, (fi15 ^ (fi15 >> ((1) & 63)))}
                    var v21 uint16 = 0
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
        }
        _ = fi15
    }
    v22 := []int{}
    for fi23 := 0; fi23 < 4; fi23++ {
        v1 = float32(((3.14159 / 0.1) / (float64((v10 - v1)) + (1.5 / (0.1 - -2.25)))))
        for fi24 := 0; fi24 < 7; fi24++ {
            if false {
                var v25 bool = false
                v26 := GenS3{F0: 65535}
                _ = v25
                _ = v26
            } else {
                var v27 bool = false
                _ = v27
            }
            v28 := GenS4{F0: -118}
            _ = fi24
            _ = v28
        }
        if len(v22) > 0 {
            v22[0] = 1
        }
        _ = fi23
    }
    if true {
        v29 := GenS2{F0: 2, F1: genFn1(42, 0), F2: 0}
        _ = v29
    } else {
        v22 = append(v22, int(v5))
        v30 := GenS5{F0: -25}
        _ = v30
    }
    v2 = v4
    v11 -= v11
    var v31 int8 = 0
    for fi32 := 0; fi32 < 10; fi32++ {
        var v33 bool = ((-(fi32)) != fi32)
        _ = fi32
        _ = v33
    }
    switch 127 {
    case -9223372036854775808:
        v31 += v31
    default:
        v34 := GenS1{F0: 4294967295, F1: v3, F2: true}
        _ = v34
    }
    _ = len(v22)
    var v35 int32 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v35
    _ = v4
    _ = v5
}
