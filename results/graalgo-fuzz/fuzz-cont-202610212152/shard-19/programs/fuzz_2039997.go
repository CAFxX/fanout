package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int }
type GenS2 struct { F0 uint16; F1 int16; F2 uint }
func genFn0(p0 int) int8 {
    if ((582 < 0) && ((!(1 <= 429)) && true)) {
        p0 = p0
        switch "" {
        case "hello":
            v5 := GenS1{F0: 0, F1: 2, F2: 2147483647}
            _ = v5
        case "":
            fmt.Println(p0)
        case "foo bar":
            for fi6 := 0; fi6 < 12; fi6++ {
                var v7 uint32 = 1
                fmt.Println(v7)
                switch (-(fi6)) {
                case -9223372036854775808:
                    p0 = (^(fi6))
                default:
                    var v8 uint16 = 42
                    _ = v8
                }
                _ = fi6
                _ = v7
            }
        }
        if false {
            fmt.Println(p0)
            p0 = p0
        }
    }
    v9 := GenS1{F0: -339, F1: 2, F2: p0}
    var v10 uint8 = 255
    _ = p0
    _ = v10
    _ = v9
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{}
    var v3 float64 = -2.25
    var v4 float32 = float32((((float64(0.0) / v3) * (1e10 / float64((3.14159 - (float32(float64((3.14159 / 1.5))) - (1e10 + (1e10 * (float32(v3) - 3.14159)))))))) / (3.14159 * float64(float32(1e10)))))
    switch -1 {
    case -1:
        var v11 int32 = -526
        _ = v11
    case -675:
        var v12 int16 = -345
        _ = v12
    case -9223372036854775808:
        var v13 int16 = 32767
        _ = v13
    }
    var v14 int = 9223372036854775807
    v4 -= ((0.1 - (v4 * v4)) + v4)
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 float64 = (v3 - v3)
        v4 += float32(0.0)
        v4 += v4
        _ = fi15
        _ = v16
    }
    var v17 uint64 = 42
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 uint64 = (^(v17))
        fmt.Println(v4, v14)
        var v20 int8 = genFn0((fi18 % ((127) | 1)))
        _ = fi18
        _ = v19
        _ = v20
    }
    if (v4 == (v4 + (v4 / v4))) {
        if true {
            if (42 <= 42) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v3 = 0.1
                v17 = v17
            }
            v21 := v1[914]
            v3 -= float64(float32(-2.25))
            _ = v21
        }
        var v22 int64 = 0
        _ = v22
    } else {
        if ((-2.25 != ((v3 / -2.25) / v3)) && (42 < 0)) {
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
        } else {
            v3 -= -2.25
            v3 += (float64((3.14159 + ((v4 - v4) / ((3.14159 - v4) / v4)))) + (3.14159 / (v3 - float64(float32(0.1)))))
        }
        v14 = v14
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        if true {
            v24 := GenS2{F0: 65535, F1: 42, F2: 748}
            var v25 float64 = (v3 + (0.1 + 3.14159))
            var v26 uint8 = 1
            _ = v24
            _ = v25
            _ = v26
        }
        if len(v2) > 0 {
            v2[0] = ""
        }
        _ = fi23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
