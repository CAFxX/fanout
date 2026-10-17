package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 uint64 }
func genFn0() uint64 {
    v7 := GenS1{F0: -1}
    switch 456 {
    case 42:
        var v8 int32 = 0
        _ = v8
    case 32767:
        var v9 float32 = (((1.5 * 1.5) * float32(0.1)) * (float32(1.5) / 1.5))
        _ = v9
    default:
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 float64 = 1e10
            _ = fi10
            _ = v11
        }
    }
    _ = v7
    return 4294967295
}
func genFn1(p0 uint16, p1 string) uint {
    p0 += 32767
    var v12 float64 = (0.1 * 1.5)
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = v12
    return 370
}

func main() {
    v1 := []string{"hello", "foo bar"}
    v2 := make(map[string]int)
    var v3 int64 = -715
    var v4 int = 9223372036854775807
    var v5 int16 = 42
    var v6 float32 = float32(1e10)
    v13 := GenS1{F0: 2147483647}
    v14 := GenS1{F0: 42}
    if false {
        v4 = (v4 * v4)
    }
    fmt.Println(v5, v5)
    fmt.Println(v5, v6, v6)
    var v15 uint16 = 32767
    fmt.Println(v6, v15)
    v6 = v6
    switch (v4 / ((v4) | 1)) {
    case 42:
        for fi16 := 0; fi16 < 2; fi16++ {
            if (2 >= 255) {
                var v17 int32 = 1
                for fi18 := 0; fi18 < 8; fi18++ {
                    if (!(false && false)) {
                        fmt.Println(v15, v3)
                        var v19 int64 = int64(v17)
                        v17 = (v17 - v17)
                        _ = v19
                    }
                    fmt.Println(v5, fi18)
                    _ = fi18
                }
                _ = v17
            } else {
                v20 := v2["foo bar"]
                v21 := GenS1{F0: 532}
                if (!true) {
                    var v22 int = v4
                    var v23 bool = true
                    if true {
                        var v24 uint8 = 128
                        v25 := GenS1{F0: -2147483648}
                        _ = v24
                        _ = v25
                    } else {
                        v26 := GenS2{F0: 18446744073709551615, F1: 42}
                        _ = v26
                    }
                    _ = v22
                    _ = v23
                } else {
                    var v27 int = v20
                    _ = v27
                }
                _ = v20
                _ = v21
            }
            _ = fi16
        }
    case 127:
        var v28 uint32 = 4294967295
        _ = v28
    case -1:
        var v29 uint = 333
        _ = v29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
