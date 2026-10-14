package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint16 }
func genFn0() int64 {
    var v7 int = 32767
    fmt.Println(v7)
    if true {
        v8 := GenS1{F0: v7, F1: 944, F2: 65535}
        v9 := GenS1{F0: v7, F1: 65535, F2: 324}
        _ = v8
        _ = v9
    }
    _ = v7
    return -9223372036854775808
}

func main() {
    var v1 uint16 = 532
    v2 := []string{"hello"}
    var v3 uint = 4294967295
    var v4 int8 = 127
    v5 := make(map[string]int)
    var v6 uint8 = 48
    var v10 float32 = float32((((((3.14159 - 1.5) / float64(-2.25)) - 1.5) * (3.14159 * 0.1)) * (float64((0.1 + 1.5)) * (0.1 - 0.1))))
    if true {
        for fi11 := 0; fi11 < 19; fi11++ {
            if (18446744073709551615 < 42) {
                switch 9223372036854775807 {
                case 128:
                    v2 = append(v2, "\n\t")
                case -599:
                    v6 = uint8(v1)
                case 871:
                    v12 := GenS1{F0: (fi11 % ((fi11) | 1)), F1: 0, F2: 32767}
                    _ = v12
                }
            }
            _ = fi11
        }
        v6 = (v6 * 1)
        var v13 int32 = -267
        _ = v13
    }
    if len(v2) > 0 {
        v2[0] = ""
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        if false {
            v15 := GenS1{F0: -18, F1: 2, F2: 983}
            if false {
                var v16 int32 = 410
                v17 := GenS1{F0: (-(fi14)), F1: 0, F2: (^(v1))}
                v6 += 42
                _ = v16
                _ = v17
            } else {
                delete(v5, "hello")
                fmt.Println(v3, v10, v1)
                v4 -= (^(v4))
            }
            _ = v15
        } else {
            fmt.Println(v3, v1)
            for fi18 := 0; fi18 < 20; fi18++ {
                var v19 uint8 = (v6 & (^(v6)))
                var v20 float32 = (v10 + v10)
                _ = fi18
                _ = v19
                _ = v20
            }
        }
        for fi21 := 0; fi21 < 13; fi21++ {
            v10 = 1.5
            var v22 uint8 = (v6 - v6)
            _ = fi21
            _ = v22
        }
        _ = fi14
    }
    if len(v2) > 0 {
        v2[0] = ""
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        var v24 float64 = (1e10 / 1.5)
        delete(v5, "")
        var v25 string = "\n\t"
        _ = fi23
        _ = v24
        _ = v25
    }
    v3 -= uint(v6)
    v26 := v5["foo bar"]
    var v27 uint16 = 119
    _ = len(v2)
    v6 = 127
    v28 := GenS1{F0: v26, F1: 0, F2: 743}
    v29 := GenS1{F0: 32767, F1: 664, F2: v27}
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
