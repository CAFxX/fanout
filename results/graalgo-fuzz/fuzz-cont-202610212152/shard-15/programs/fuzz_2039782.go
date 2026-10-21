package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int16 }
type GenS2 struct { F0 int32; F1 int8; F2 uint16 }
type GenS3 struct { F0 int8; F1 uint8 }
func genFn0(p0 string, p1 uint64, p2 int64) int {
    v6 := GenS1{F0: 890, F1: 772, F2: 1}
    v7 := []int{}
    v8 := GenS1{F0: 65535, F1: 32767, F2: 32767}
    var v9 uint32 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 761
}
func genFn1(p0 int) uint64 {
    var v10 uint8 = 128
    v10 -= 128
    v10 -= v10
    switch -772 {
    case 32767:
        v11 := GenS2{F0: 42, F1: -1, F2: 1}
        _ = v11
    default:
        v12 := GenS2{F0: 0, F1: 100, F2: 65535}
        _ = v12
    }
    v10 -= v10
    _ = p0
    _ = v10
    return 4294967295
}

func main() {
    v1 := []uint32{2147483648, 65535, 715}
    var v2 int64 = 0
    var v3 int32 = -214
    var v4 int16 = -363
    var v5 int = 670
    v4 += -32768
    var v13 bool = true
    v13 = true
    if v13 {
        v4 += -32768
        var v14 float32 = ((0.0 / 1.5) / 1e10)
        for fi15 := 0; fi15 < 13; fi15++ {
            _ = len(v1)
            _ = fi15
        }
        _ = v14
    } else {
        if v13 {
            if (701 < v3) {
                v16 := GenS2{F0: (^(v3)), F1: 100, F2: 32767}
                if false {
                    v17 := GenS2{F0: 886, F1: 100, F2: 476}
                    var v18 uint = 0
                    _ = v17
                    _ = v18
                } else {
                    fmt.Println(v4, v2, v3)
                    fmt.Println(v13, v2, v3)
                    if (!v13) {
                        var v19 bool = v13
                        var v20 int32 = v3
                        _ = v19
                        _ = v20
                    }
                }
                v21 := GenS1{F0: 2147483647, F1: (v5 % (((v5 >> ((1) & 63))) | 1)), F2: v4}
                _ = v16
                _ = v21
            } else {
                v22 := GenS2{F0: (^(v3)), F1: 0, F2: 0}
                fmt.Println(v3, v2, v2)
                v13 = v13
                _ = v22
            }
        } else {
            v23 := GenS2{F0: v3, F1: -1, F2: 688}
            _ = v23
        }
        var v24 bool = v13
        v1 = append(v1, 2)
        _ = v24
    }
    v5 = 1
    v25 := GenS2{F0: 42, F1: 127, F2: 139}
    fmt.Println(v4, v3, v2)
    var v26 float32 = 3.14159
    fmt.Println(v26)
    var v27 int64 = 51
    var v28 uint16 = 65535
    var v29 uint64 = 42
    var v30 uint8 = 121
    v5 -= (^(v5))
    for fi31 := 0; fi31 < 8; fi31++ {
        if ((!v13) || v13) {
            for fi32 := 0; fi32 < 17; fi32++ {
                v33 := GenS3{F0: int8(v27), F1: v30}
                fmt.Println(v2, fi31)
                _ = fi32
                _ = v33
            }
            v29 = 0
            v1 = append(v1, 65535)
        } else {
            if (!v13) {
                var v34 float64 = 1.5
                _ = v34
            } else {
                var v35 int16 = (v4 * v4)
                fmt.Println(v35, v27)
                v4 += v4
                _ = v35
            }
            if v13 {
                var v36 string = (("\n\t" + ("hello" + "\n\t")) + (("" + "\n\t") + ("hello" + "a")))
                _ = v36
            } else {
                var v37 int32 = v3
                _ = v37
            }
            v27 -= (v27 << ((4294967295) & 63))
        }
        var v38 uint64 = v29
        _ = fi31
        _ = v38
    }
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
