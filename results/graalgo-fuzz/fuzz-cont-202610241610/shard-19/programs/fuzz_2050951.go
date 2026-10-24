package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint; F2 int16 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS4 struct { F0 int32; F1 int; F2 uint }

func main() {
    v1 := []uint32{1, 0, 42, 2147483648}
    var v2 float32 = 1e10
    var v3 float64 = 0.0
    var v4 int64 = 1
    var v5 int32 = 42
    v3 -= v3
    v4 = 9223372036854775807
    if (844 < 1) {
        if true {
            if len(v1) > 0 {
                v1[0] = 42
            }
            v3 = float64(0.0)
            if false {
                for fi6 := 0; fi6 < 16; fi6++ {
                    v7 := GenS1{F0: v5, F1: 0}
                    v2 = v2
                    _ = fi6
                    _ = v7
                }
            } else {
                if (70 >= -77) {
                    v1 = append(v1, 42)
                    fmt.Println(v2, v3, v3)
                }
                fmt.Println(v3, v4)
                if false {
                    v8 := GenS2{F0: 1, F1: 18446744073709551615, F2: 1}
                    if len(v1) > 0 {
                        v1[0] = 4294967295
                    }
                    v4 -= v4
                    _ = v8
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
        }
        v9 := GenS2{F0: -2147483648, F1: 18446744073709551615, F2: -32768}
        var v10 int16 = 32767
        _ = v10
        _ = v9
    }
    var v11 int8 = -40
    _ = len(v1)
    var v12 float64 = v3
    var v13 int64 = v4
    if ((1 > 244) && true) {
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := GenS1{F0: v5, F1: 128}
            v12 -= float64(-2.25)
            if false {
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 bool = false
                    v1 = append(v1, 2147483648)
                    if v17 {
                        v11 -= (v11 >> ((4294967295) & 63))
                        _ = len(v1)
                        switch fi14 {
                        case 0:
                            fmt.Println(v12)
                        }
                    }
                    _ = fi16
                    _ = v17
                }
                fmt.Println(fi14)
                var v18 int = 127
                _ = v18
            }
            _ = fi14
            _ = v15
        }
    } else {
        v19 := GenS3{F0: 255, F1: 877, F2: v5}
        v5 = 0
        _ = v19
    }
    v20 := GenS4{F0: (v5 | v5), F1: -9223372036854775808, F2: 65535}
    var v21 float64 = v3
    var v22 float32 = float32((float64(v2) * v21))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
