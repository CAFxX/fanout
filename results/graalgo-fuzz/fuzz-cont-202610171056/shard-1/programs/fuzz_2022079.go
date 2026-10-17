package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint32; F2 int32 }

func main() {
    v1 := make(map[string]string)
    var v2 int8 = 0
    var v3 int = -2147483648
    var v4 int32 = -628
    var v5 int = 2147483647
    var v6 string = ""
    v5 += (v5 / ((42) | 1))
    v7 := GenS1{F0: (v5 & (v5 >> ((4294967295) & 63)))}
    v4 += 0
    v8 := GenS1{F0: -831}
    for fi9 := 0; fi9 < 10; fi9++ {
        fmt.Println(v2, v4)
        _ = fi9
    }
    if (true && (2 <= 78)) {
        v2 += -4
        var v10 uint16 = 2
        fmt.Println(v2)
        _ = v10
    }
    switch (v3 | v5) {
    case 1:
        var v11 uint64 = 95
        _ = v11
    case 429:
        if false {
            v12 := []int{}
            if len(v12) > 0 {
                v12[0] = (v3 | -2147483648)
            }
            var v13 int32 = -499
            _ = v12
            _ = v13
        } else {
            v4 += v4
            var v14 int32 = (v4 - (v4 / ((0) | 1)))
            _ = v14
        }
    case 42:
        var v15 uint16 = 32767
        _ = v15
    }
    v16 := []int{}
    if ((v6 + v6) != v6) {
        var v17 uint64 = 42
        _ = v17
    }
    v18 := GenS1{F0: 9223372036854775807}
    for fi19 := 0; fi19 < 13; fi19++ {
        for fi20 := 0; fi20 < 16; fi20++ {
            var v21 int64 = 1
            v2 -= v2
            if false {
                var v22 uint16 = 255
                v23 := v1[((v6 + v6) + (((v6 + "hello") + (v6 + (v6 + (v6 + v6)))) + "hello"))]
                v1[("a" + (v6 + ((v23 + "foo bar") + (((v6 + v6) + v6) + v6))))] = "\n\t"
                _ = v22
                _ = v23
            } else {
                switch "" {
                case "":
                    var v24 uint32 = 278
                    _ = v24
                }
                if (false || false) {
                    var v25 int8 = 1
                    fmt.Println(v25, v21, v6)
                    v25 = v25
                    _ = v25
                } else {
                    _ = len(v16)
                }
            }
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    var v26 bool = true
    if v26 {
        v3 = 2147483647
        v26 = (834 != 32767)
        v3 -= (v5 * v3)
    }
    _ = len(v1)
    if v26 {
        var v27 uint16 = 65535
        fmt.Println(v27)
        v28 := GenS2{F0: 165, F1: 2147483648, F2: -560}
        _ = v27
        _ = v28
    }
    switch 32767 {
    case 594:
        v2 = v2
    case 4294967295:
        var v29 float32 = (float32(1.5) + 3.14159)
        _ = v29
    case 0:
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
    default:
        fmt.Println(v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
