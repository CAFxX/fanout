package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int32 = 1
    v3 := make(map[int]string)
    var v4 uint32 = 65535
    fmt.Println(v1, v2)
    v3[1] = "foo bar"
    v5 := []int{}
    if false {
        if false {
            v3[42] = "foo bar"
        }
    } else {
        for fi6 := 0; fi6 < 16; fi6++ {
            v7 := GenS1{F0: -1, F1: true}
            if false {
                fmt.Println(v4, v2)
                v8 := GenS1{F0: 127, F1: (0.1 == ((((0.1 * -2.25) / 0.1) / (1e10 - (1e10 * 3.14159))) - 3.14159))}
                _ = v8
            } else {
                v4 += (v4 >> ((4294967295) & 63))
            }
            if true {
                _ = len(v3)
                var v9 int16 = 42
                var v10 int8 = int8(v2)
                _ = v10
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    }
    fmt.Println(v4, v4, v4)
    if false {
        for fi11 := 0; fi11 < 18; fi11++ {
            v2 = v2
            var v12 int64 = -1
            fmt.Println(v2, v1)
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 6; fi13++ {
            v1 = v1
            var v14 float64 = 3.14159
            _ = fi13
            _ = v14
        }
        fmt.Println(v2, v2)
    } else {
        if len(v5) > 0 {
            v5[0] = -745
        }
        for fi15 := 0; fi15 < 5; fi15++ {
            for fi16 := 0; fi16 < 18; fi16++ {
                if (false || (v2 >= (v2 / ((v2) | 1)))) {
                    fmt.Println(fi16, v1, v1)
                    v4 -= v4
                    fmt.Println(v1, v1, v4)
                }
                _ = fi16
            }
            _ = fi15
        }
    }
    v17 := v3[32768]
    fmt.Println(v2, v1, v1)
    _ = len(v3)
    var v18 float32 = ((-2.25 + 3.14159) + 1e10)
    var v19 int64 = 42
    v20 := GenS1{F0: -128, F1: true}
    v19 -= (v19 & v19)
    if false {
        if (-531 <= (v19 % ((int64(v2)) | 1))) {
            var v21 uint8 = 255
            for fi22 := 0; fi22 < 2; fi22++ {
                for fi23 := 0; fi23 < 9; fi23++ {
                    if (v17 <= v17) {
                        var v24 int32 = (v2 * 42)
                        _ = v24
                    } else {
                        v21 -= v21
                        if false {
                            var v25 int32 = (v2 + (v2 << ((4294967295) & 63)))
                            v2 += v25
                            var v26 int = (fi22 - (fi22 ^ -363))
                            _ = v25
                            _ = v26
                        } else {
                            var v27 float64 = -2.25
                            if (528 == (v1 / ((v1) | 1))) {
                                fmt.Println(v27, v4)
                            }
                            _ = v27
                        }
                        v18 = 3.14159
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi23
                }
                _ = fi22
            }
            var v28 float32 = (v18 + 0.1)
            _ = v21
            _ = v28
        } else {
            var v29 float64 = float64(3.14159)
            _ = v29
        }
        var v30 uint64 = v1
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
