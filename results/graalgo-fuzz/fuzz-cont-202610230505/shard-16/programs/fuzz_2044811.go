package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int16 }
type GenS2 struct { F0 string; F1 string; F2 bool }
type GenS3 struct { F0 int64 }

func main() {
    var v1 bool = true
    v2 := make(map[uint32]uint32)
    var v3 int8 = 42
    var v4 int32 = -2147483648
    v5 := []int{-9223372036854775808}
    var v6 string = "foo bar"
    v7 := GenS1{F0: 151, F1: -893, F2: 42}
    v1 = v1
    fmt.Println(v3, v4, v4)
    var v8 uint8 = 127
    v1 = false
    for fi9 := 0; fi9 < 6; fi9++ {
        v10 := GenS2{F0: v6, F1: "\n\t", F2: (!(-9223372036854775808 > -1))}
        if true {
            var v11 uint16 = 65535
            if (fi9 <= (fi9 & int(v8))) {
                var v12 float64 = -2.25
                _ = v12
            } else {
                for fi13 := 0; fi13 < 1; fi13++ {
                    fmt.Println(fi9, v1, v3)
                    if (-548 >= -32768) {
                        switch fi9 {
                        case 32768:
                            switch fi9 {
                            case 2147483647:
                                v4 = (v4 - v4)
                            case 9223372036854775807:
                                var v14 uint16 = v11
                                _ = v14
                            default:
                                v15 := GenS1{F0: 2147483648, F1: -248, F2: 0}
                                _ = v15
                            }
                        case 32767:
                            v8 = v8
                        }
                    }
                    if (!((v3 | int8(v8)) == v3)) {
                        v8 += (v8 % ((2) | 1))
                        delete(v2, 255)
                        v4 = v4
                    } else {
                        v8 = 37
                        v16 := GenS2{F0: v6, F1: "", F2: ((v8 * v8) <= (v8 & v8))}
                        _ = v16
                    }
                    _ = fi13
                }
            }
            _ = v11
        }
        v8 = v8
        _ = fi9
        _ = v10
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v3 += 82
    for fi17 := 0; fi17 < 12; fi17++ {
        for fi18 := 0; fi18 < 1; fi18++ {
            if ((!(0.0 <= (0.0 + 1e10))) || (v3 >= (v3 << ((18446744073709551615) & 63)))) {
                delete(v2, 0)
                if (!(-2.25 != float64(float32(0.0)))) {
                    v6 = v6
                }
                var v19 int64 = -754
                _ = len(v5)
                v4 += v4
                _ = v19
            }
            fmt.Println(v6)
            if (128 < (fi17 << ((382) & 63))) {
                if v1 {
                    if (v1 || (((true && v1) || (-907 == int16(v3))) && (v1 && v1))) {
                        var v20 bool = false
                        v21 := GenS2{F0: v6, F1: v6, F2: v20}
                        _ = v20
                        _ = v21
                    }
                    var v22 uint8 = 127
                    _ = v22
                }
            }
            _ = fi18
        }
        if (42 <= 2) {
            var v23 uint8 = v8
            _ = v23
        } else {
            if len(v5) > 0 {
                v5[0] = (fi17 | -1)
            }
            if v1 {
                v24 := GenS3{F0: 820}
                _ = v24
            } else {
                fmt.Println(v8)
                if len(v5) > 0 {
                    v5[0] = fi17
                }
            }
        }
        _ = fi17
    }
    var v25 int32 = -2147483648
    v3 += (v3 * 100)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
