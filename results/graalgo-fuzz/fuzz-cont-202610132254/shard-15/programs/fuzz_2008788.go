package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 bool }
func genFn0(p0 uint64, p1 string, p2 bool) int8 {
    var v7 string = p1
    var v8 int64 = 42
    v9 := GenS1{F0: 0}
    var v10 bool = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1() int {
    var v11 uint8 = 0
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 int16 = 1
        v14 := []int{-537, 9223372036854775807}
        for fi15 := 0; fi15 < 13; fi15++ {
            if (610 <= 0) {
                v13 = 1
                v11 = 87
                v16 := GenS1{F0: 2}
                _ = v16
            } else {
                v13 += (v13 % ((v13) | 1))
            }
            var v17 int16 = v13
            v14 = append(v14, (-(fi12)))
            v14 = append(v14, 4294967295)
            var v18 uint32 = 2147483648
            _ = fi15
            _ = v17
            _ = v18
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    v11 -= v11
    v11 -= (v11 >> ((502) & 63))
    _ = v11
    return 32767
}

func main() {
    var v1 float32 = 3.14159
    var v2 uint32 = 42
    var v3 int8 = -1
    var v4 string = "\n\t"
    v5 := make(map[string]uint32)
    v6 := []int{128, 998}
    if true {
        var v19 int32 = 1
        fmt.Println(v1, v4)
        switch "\n\t" {
        case "a":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        case "foo bar":
            v20 := v5[(v4 + v4)]
            _ = v20
        }
        _ = v19
    } else {
        v3 -= v3
        var v21 int32 = 1
        var v22 int32 = int32(v3)
        _ = v21
        _ = v22
    }
    v3 = -1
    v1 += (((float32(-2.25) / (v1 / float32(3.14159))) / v1) - v1)
    if (4294967295 > 1) {
        fmt.Println(v1)
        v23 := GenS1{F0: 65535}
        _ = v23
    }
    fmt.Println(v1)
    v3 = (v3 ^ (v3 | v3))
    for fi24 := 0; fi24 < 13; fi24++ {
        for fi25 := 0; fi25 < 2; fi25++ {
            v1 = v1
            _ = fi25
        }
        for fi26 := 0; fi26 < 8; fi26++ {
            var v27 int16 = int16(fi26)
            v28 := GenS2{F0: 0, F1: (128 < 0)}
            v29 := GenS2{F0: 1, F1: (true || (0 > 29))}
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
        }
        _ = fi24
    }
    for fi30 := 0; fi30 < 8; fi30++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi30
    }
    switch v4 {
    case "hello":
        var v31 string = v4
        _ = v31
    case "a":
        _ = len(v6)
    case "\n\t":
        v3 = (v3 + (v3 % (((v3 + genFn0(0, "a", true))) | 1)))
    }
    v3 -= v3
    fmt.Println(v3, v1)
    v32 := GenS2{F0: 42, F1: true}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
