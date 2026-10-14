package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int16; F1 int32; F2 uint64 }
func genFn0(p0 int32) uint32 {
    v7 := []int{128, -241}
    fmt.Println(p0)
    if true {
        var v8 uint64 = 966
        if false {
            var v9 uint8 = 128
            var v10 uint = 289
            _ = v10
            _ = v9
        } else {
            if ((3.14159 >= 3.14159) && true) {
                for fi11 := 0; fi11 < 7; fi11++ {
                    fmt.Println(p0)
                    for fi12 := 0; fi12 < 10; fi12++ {
                        v8 += (v8 & v8)
                        v8 = v8
                        p0 += (p0 ^ 2147483647)
                        _ = fi12
                    }
                    _ = fi11
                }
            } else {
                var v13 uint = 4294967295
                _ = len(v7)
                _ = v13
            }
            if false {
                if len(v7) > 0 {
                    v7[0] = -1
                }
            } else {
                if true {
                    var v14 float64 = float64(float32((1.5 - -2.25)))
                    _ = len(v7)
                    for fi15 := 0; fi15 < 1; fi15++ {
                        if len(v7) > 0 {
                            v7[0] = (fi15 / (((fi15 & 32767)) | 1))
                        }
                        var v16 uint16 = 0
                        _ = fi15
                        _ = v16
                    }
                    _ = v14
                }
                p0 += (p0 << ((18446744073709551615) & 63))
            }
        }
        v17 := GenS1{F0: 434}
        _ = v17
        _ = v8
    } else {
        var v18 int64 = -1
        _ = v18
    }
    p0 = p0
    fmt.Println(p0)
    _ = p0
    _ = v7
    return 410
}

func main() {
    var v1 int64 = -245
    var v2 float64 = 3.14159
    var v3 int64 = v1
    var v4 string = (("" + "a") + "a")
    var v5 int64 = v3
    var v6 int16 = -453
    v2 = ((-2.25 / -2.25) - (1e10 - 0.0))
    v2 += float64(-2.25)
    for fi19 := 0; fi19 < 4; fi19++ {
        for fi20 := 0; fi20 < 18; fi20++ {
            fmt.Println(v6, v1, v1)
            _ = fi20
        }
        v21 := []int{(fi19 & fi19), 9223372036854775807}
        v3 = v3
        _ = fi19
        _ = v21
    }
    v22 := GenS2{F0: (^(v6)), F1: 2147483647}
    var v23 uint8 = 127
    for fi24 := 0; fi24 < 18; fi24++ {
        v25 := []int{fi24, (fi24 / (((fi24 << ((0) & 63))) | 1)), (fi24 + (fi24 << ((474) & 63)))}
        var v26 int = fi24
        v27 := GenS3{F0: v6, F1: -911, F2: 2}
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        v6 -= v6
        _ = fi28
    }
    v29 := GenS2{F0: 32767, F1: -144}
    for fi30 := 0; fi30 < 19; fi30++ {
        v23 += v23
        _ = fi30
    }
    var v31 uint = 18446744073709551615
    fmt.Println(v2, v3)
    var v32 uint32 = 509
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
