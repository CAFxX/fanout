package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 2147483648
    var v3 int64 = 42
    var v4 float32 = (1e10 * ((3.14159 / (1e10 * float32((((0.1 * 3.14159) * -2.25) + 1.5)))) / (-2.25 * 0.1)))
    var v5 uint64 = 222
    var v6 uint8 = 0
    if false {
        var v7 int16 = int16(v5)
        if false {
            if (v2 < (v2 >> ((889) & 63))) {
                switch 1 {
                case 2147483647:
                    v3 += v3
                case 768:
                    v8 := GenS1{F0: 705, F1: v3}
                    _ = v8
                default:
                    for fi9 := 0; fi9 < 15; fi9++ {
                        v10 := []int{fi9}
                        v2 = (v2 + v2)
                        var v11 uint = 782
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                }
                var v12 uint = uint(v2)
                _ = v12
            }
            if true {
                fmt.Println(v2, v6)
                var v13 uint16 = 549
                _ = v13
            } else {
                switch -9223372036854775808 {
                case 42:
                    v14 := GenS2{F0: 482}
                    _ = v14
                }
                if (-2147483648 > -969) {
                    v15 := GenS1{F0: v3, F1: 0}
                    for fi16 := 0; fi16 < 6; fi16++ {
                        v5 = v5
                        v5 -= (^(v5))
                        var v17 int32 = 1
                        _ = fi16
                        _ = v17
                    }
                    _ = v15
                } else {
                    var v18 uint16 = 42
                    var v19 bool = true
                    _ = v18
                    _ = v19
                }
            }
        } else {
            for fi20 := 0; fi20 < 15; fi20++ {
                v2 = v2
                v21 := GenS2{F0: 443}
                _ = fi20
                _ = v21
            }
            var v22 uint32 = (v2 + 42)
            var v23 int64 = v3
            _ = v22
            _ = v23
        }
        _ = v7
    }
    var v24 int8 = 0
    var v25 uint16 = 65535
    v1[v2] = 472
    var v26 uint16 = (v25 & (^(v25)))
    v27 := []int{}
    var v28 string = "\n\t"
    v24 = (v24 | v24)
    if (true || (!true)) {
        fmt.Println(v25, v26)
        v29 := GenS1{F0: (v3 & (v3 | (^(v3)))), F1: int64(v2)}
        if (v28 <= ("hello" + v28)) {
            var v30 float64 = 3.14159
            v31 := GenS1{F0: (^(v3)), F1: v3}
            _ = v30
            _ = v31
        }
        _ = v29
    }
    if ((v2 - 224) < (v2 | 224)) {
        v32 := GenS3{F0: v25}
        for fi33 := 0; fi33 < 3; fi33++ {
            v26 += v25
            if len(v27) > 0 {
                v27[0] = (fi33 + fi33)
            }
            _ = fi33
        }
        _ = v32
    }
    if true {
        var v34 int32 = -2147483648
        v35 := GenS1{F0: -9223372036854775808, F1: (v3 & (v3 + int64(v5)))}
        _ = v34
        _ = v35
    }
    v3 += 9223372036854775807
    v36 := GenS4{F0: (v24 << ((210) & 63))}
    v37 := GenS5{F0: (^(v25))}
    fmt.Println(v24)
    v38 := GenS1{F0: (v3 % ((v3) | 1)), F1: v3}
    fmt.Println(v5, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
