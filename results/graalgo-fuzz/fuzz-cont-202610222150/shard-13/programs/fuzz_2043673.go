package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }

func main() {
    var v1 uint32 = 0
    var v2 int16 = int16(v1)
    var v3 float32 = 0.0
    for fi4 := 0; fi4 < 20; fi4++ {
        v2 += (v2 - v2)
        v3 = v3
        _ = fi4
    }
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 int8 = 1
        var v7 int64 = -977
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint = 42
    for fi9 := 0; fi9 < 15; fi9++ {
        v2 = v2
        v10 := GenS1{F0: fi9, F1: 127}
        _ = fi9
        _ = v10
    }
    if ((^(v2)) < (v2 << (((v8 | 42)) & 63))) {
        var v11 int64 = 0
        _ = v11
    }
    v3 = (float32((0.0 + 1e10)) / 3.14159)
    v8 = v8
    if false {
        v8 = (v8 % ((uint(v1)) | 1))
    }
    if (v2 != (v2 & (v2 - (v2 - (v2 & -897))))) {
        v12 := []int{9223372036854775807}
        if false {
            var v13 int16 = 32767
            _ = v13
        } else {
            v12 = append(v12, -2147483648)
            for fi14 := 0; fi14 < 5; fi14++ {
                v1 -= (v1 - uint32(v8))
                _ = len(v12)
                for fi15 := 0; fi15 < 4; fi15++ {
                    if (127 != 0) {
                        if false {
                            v16 := GenS1{F0: fi15, F1: 163}
                            var v17 int16 = 0
                            var v18 uint8 = 128
                            switch ("\n\t" + "") {
                            case "hello":
                                fmt.Println(v1)
                            }
                            v17 -= (v17 | v2)
                            var v19 uint16 = 0
                            fmt.Println(v18, fi15, v18)
                            _ = v16
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                        v20 := GenS1{F0: fi15, F1: 116}
                        var v21 int64 = 0
                        _ = v20
                        _ = v21
                    } else {
                        v8 -= (v8 + (v8 * v8))
                        v22 := GenS1{F0: (fi14 * fi15), F1: 0}
                        var v23 uint32 = (^(v1))
                        _ = v22
                        _ = v23
                    }
                    fmt.Println(v2, v2)
                    var v24 int64 = 1
                    _ = fi15
                    _ = v24
                }
                _ = fi14
            }
            var v25 float32 = 1e10
            _ = v25
        }
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
}
