package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS3 struct { F0 int64; F1 int16; F2 uint16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 1
    v3 := make(map[int]bool)
    v4 := []string{(("hello" + ("" + ("" + "\n\t"))) + ("foo bar" + "")), ""}
    v5 := make(map[int]string)
    v2 -= v2
    v2 += v2
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS1{F0: 0}
        if false {
            var v8 float32 = 0.1
            fmt.Println(v2, v8)
            _ = v8
        } else {
            fmt.Println(v2, v2)
            if (v2 >= (-(v2))) {
                v2 = v2
            }
            var v9 int16 = v2
            _ = v9
        }
        v1[""] = 4294967295
        _ = fi6
        _ = v7
    }
    var v10 uint64 = 1
    v11 := GenS2{F0: 0, F1: 42, F2: 65535}
    var v12 float32 = ((-2.25 + 1e10) / (float32(float64(1.5)) - float32((1.5 * ((float64((1e10 / 1e10)) * float64(1.5)) + 0.1)))))
    v13 := GenS2{F0: 2, F1: 9223372036854775807, F2: 1}
    var v14 uint16 = 543
    for fi15 := 0; fi15 < 2; fi15++ {
        for fi16 := 0; fi16 < 4; fi16++ {
            v12 += v12
            _ = fi16
        }
        var v17 uint32 = 42
        switch (^(fi15)) {
        case -9223372036854775808:
            fmt.Println(v12, v12)
        case 640:
            for fi18 := 0; fi18 < 13; fi18++ {
                if true {
                    v19 := GenS2{F0: 430, F1: 9223372036854775807, F2: uint32(v10)}
                    if true {
                        v12 -= (float32((3.14159 * 1.5)) + (((-2.25 * v12) + 3.14159) * 0.0))
                    } else {
                        var v20 int = (^(fi15))
                        v21 := GenS3{F0: -9223372036854775808, F1: v2, F2: (^(v14))}
                        if ((-(fi18)) <= (v20 >> ((636) & 63))) {
                            v17 -= 284
                            v14 -= 1
                            var v22 uint = 65535
                            _ = v22
                        }
                        _ = v20
                        _ = v21
                    }
                    _ = v19
                }
                _ = fi18
            }
        case -1:
            v2 += v2
        default:
            var v23 uint64 = (v10 / ((1) | 1))
            _ = v23
        }
        _ = fi15
        _ = v17
    }
    var v24 bool = (false || false)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
