package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint32; F1 uint }
func genFn0() uint64 {
    v5 := []int{1}
    if false {
        switch 0 {
        case -9223372036854775808:
            for fi6 := 0; fi6 < 13; fi6++ {
                if (!(fi6 == 127)) {
                    v7 := GenS1{F0: 255}
                    if true {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        var v8 int8 = int8(fi6)
                        _ = v8
                    }
                    _ = v7
                } else {
                    var v9 int64 = -9223372036854775808
                    var v10 uint = 255
                    _ = v10
                    _ = v9
                }
                if (fi6 >= 147) {
                    var v11 int32 = 2147483647
                    if (fi6 != 198) {
                        v12 := GenS2{F0: ((0 <= 134) && true), F1: (v11 & (v11 * (v11 << ((891) & 63))))}
                        _ = v12
                    } else {
                        _ = len(v5)
                        v11 -= v11
                    }
                    _ = v11
                }
                _ = fi6
            }
        case 32768:
            if (-2.25 <= 0.1) {
                v5 = append(v5, 931)
            }
        default:
            _ = len(v5)
        }
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 int8 = 42
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        v16 := GenS2{F0: (32767 != -704), F1: -1}
        v5 = append(v5, -2147483648)
        _ = fi15
        _ = v16
    }
    _ = v5
    return 2
}

func main() {
    var v1 string = "hello"
    v2 := make(map[int]int)
    v3 := []int{42, 32767, 127}
    var v4 uint32 = 671
    if (!(122 >= 42)) {
        if false {
            for fi17 := 0; fi17 < 3; fi17++ {
                for fi18 := 0; fi18 < 8; fi18++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v19 uint8 = 83
                    var v20 int8 = 0
                    _ = len(v3)
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                v1 = (v1 + v1)
                if false {
                    var v21 uint16 = 2
                    _ = v21
                }
                _ = fi17
            }
            for fi22 := 0; fi22 < 2; fi22++ {
                v4 = (v4 / ((v4) | 1))
                _ = fi22
            }
        }
        fmt.Println(v1, v4)
    }
    v23 := v2[127]
    switch (v23 + 459) {
    case 0:
        v24 := GenS1{F0: 1}
        _ = v24
    default:
        var v25 uint16 = 65535
        _ = v25
    }
    var v26 float64 = 0.1
    v27 := GenS3{F0: 2147483648, F1: 255}
    switch v23 {
    case 2147483647:
        fmt.Println(v4, v4, v1)
    case 32768:
        v28 := GenS3{F0: (v4 >> ((2) & 63)), F1: 671}
        _ = v28
    case -162:
        var v29 int64 = 9223372036854775807
        _ = v29
    }
    var v30 int = (v23 / ((v23) | 1))
    v26 = 1.5
    fmt.Println(v23)
    var v31 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
