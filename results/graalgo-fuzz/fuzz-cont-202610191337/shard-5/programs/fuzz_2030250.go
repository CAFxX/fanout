package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int64; F2 bool }

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 2
    v3 := []uint32{2, 4294967295, 1, 0}
    var v4 string = ((((("" + "\n\t") + "foo bar") + "a") + ("" + ("a" + ""))) + "\n\t")
    var v5 uint16 = 135
    var v6 float32 = ((float32(0.1) - 3.14159) * 1e10)
    var v7 uint32 = 42
    if (!true) {
        for fi8 := 0; fi8 < 15; fi8++ {
            for fi9 := 0; fi9 < 16; fi9++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(v2)
                _ = fi9
            }
            v5 -= v5
            if (0 == 128) {
                var v10 uint16 = v5
                var v11 string = "a"
                var v12 uint16 = (v10 * (v2 * v5))
                _ = v10
                _ = v11
                _ = v12
            }
            _ = fi8
        }
    }
    fmt.Println(v4)
    v6 -= ((float32(1.5) * (float32(-2.25) - float32(float64(v6)))) - v6)
    switch 0 {
    case 128:
        var v13 uint32 = v7
        _ = v13
    case 127:
        v14 := GenS1{F0: 65535, F1: 42}
        _ = v14
    case -1:
        v6 -= float32(((-2.25 * 3.14159) + 0.1))
    }
    v7 -= 982
    fmt.Println(v5, v4)
    var v15 int16 = -1
    v16 := GenS1{F0: (v7 * 2147483648), F1: 4294967295}
    fmt.Println(v7, v6)
    v4 = v4
    v15 += -1
    v5 = v2
    for fi17 := 0; fi17 < 14; fi17++ {
        v5 = v2
        var v18 int16 = v15
        if len(v3) > 0 {
            v3[0] = 1
        }
        _ = fi17
        _ = v18
    }
    v19 := GenS1{F0: v7, F1: 18446744073709551615}
    v20 := GenS2{F0: v7, F1: 0, F2: true}
    if (true && false) {
        for fi21 := 0; fi21 < 2; fi21++ {
            v6 = (1.5 / v6)
            v6 = v6
            _ = fi21
        }
        v22 := GenS1{F0: v7, F1: 0}
        _ = len(v1)
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
