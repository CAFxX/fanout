package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 float32 = -2.25
    var v2 int64 = 9223372036854775807
    var v3 uint16 = 1
    v4 := []uint32{2, 2147483647, 2147483648, 573}
    v5 := make(map[int]bool)
    v1 -= 0.1
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: uint64(v3), F1: 2, F2: 1000}
        var v8 int8 = -128
        var v9 uint16 = (v3 >> ((255) & 63))
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v10 string = "a"
    var v11 int = -9223372036854775808
    v12 := GenS1{F0: 4294967295, F1: 0, F2: 1000}
    if true {
        v13 := GenS2{F0: -1}
        var v14 string = v10
        if false {
            var v15 int16 = 1000
            if true {
                if (false || false) {
                    fmt.Println(v14)
                    for fi16 := 0; fi16 < 6; fi16++ {
                        v17 := GenS3{F0: 0}
                        if ((2147483647 == 0) && (-128 != -1)) {
                            var v18 uint = 18446744073709551615
                            var v19 uint = (v18 - v18)
                            if ((1 != 0) || false) {
                                var v20 float32 = v1
                                v21 := GenS3{F0: uint32(v18)}
                                v22 := GenS2{F0: (v15 << ((v19) & 63))}
                                _ = v20
                                _ = v21
                                _ = v22
                            } else {
                                if (262 == 1) {
                                    v23 := v5[fi16]
                                    _ = v23
                                } else {
                                    v18 -= (v19 ^ (v19 ^ v18))
                                }
                                var v24 string = (v14 + v10)
                                if len(v4) > 0 {
                                    v4[0] = 101
                                }
                                _ = v24
                            }
                            _ = v18
                            _ = v19
                        } else {
                            if true {
                                var v25 int32 = 288
                                _ = v25
                            }
                        }
                        _ = len(v5)
                        _ = fi16
                        _ = v17
                    }
                    delete(v5, 0)
                }
            }
            _ = v15
        }
        _ = v13
        _ = v14
    } else {
        v10 = v10
        var v26 uint32 = 2147483648
        _ = v26
    }
    fmt.Println(v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
