package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
func genFn0(p0 int64) int16 {
    fmt.Println(p0)
    var v6 int16 = -32768
    _ = p0
    _ = v6
    return (^(v6))
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]uint32)
    var v3 int32 = 0
    v4 := make(map[uint32]uint32)
    v5 := []int{}
    v3 += v3
    if false {
        if true {
            v3 -= (v3 - (v3 / (((^(v3))) | 1)))
        }
    }
    var v7 float64 = -2.25
    v3 = 0
    if (!(false || true)) {
        v7 = v7
        var v8 uint64 = 877
        _ = v8
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        for fi10 := 0; fi10 < 2; fi10++ {
            v7 += (v7 * 0.0)
            v3 = -538
            var v11 string = "\n\t"
            _ = fi10
            _ = v11
        }
        fmt.Println(fi9, v3)
        v7 += 3.14159
        _ = fi9
    }
    var v12 int16 = 1
    for fi13 := 0; fi13 < 14; fi13++ {
        fmt.Println(v12)
        for fi14 := 0; fi14 < 14; fi14++ {
            v15 := GenS1{F0: (0 != 118), F1: 9223372036854775807}
            v4[42] = 983
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    v7 = float64(float32(1.5))
    delete(v4, 42)
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 int8 = 42
        _ = fi16
        _ = v17
    }
    v12 = v12
    v12 = (v12 ^ (v12 >> ((314) & 63)))
    v3 -= (v3 % ((-101) | 1))
    if (32767 >= 255) {
        if (!(65535 != 395)) {
            v3 += int32(v12)
        } else {
            for fi18 := 0; fi18 < 8; fi18++ {
                v3 = (v3 | v3)
                v12 += genFn0(9223372036854775807)
                _ = fi18
            }
            if false {
                v7 = (v7 * 1.5)
            } else {
                if (("\n\t" > "foo bar") || false) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                if (-2.25 < 0.0) {
                    v19 := v4[42]
                    _ = v19
                } else {
                    var v20 int64 = -9223372036854775808
                    v4[1] = 4294967295
                    _ = v20
                }
                var v21 uint8 = 22
                _ = v21
            }
            for fi22 := 0; fi22 < 17; fi22++ {
                v12 -= (v12 | genFn0(-569))
                v3 = v3
                _ = fi22
            }
        }
    } else {
        switch (("" + "a") + "") {
        case "\n\t":
            var v23 int = 4294967295
            _ = v23
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
