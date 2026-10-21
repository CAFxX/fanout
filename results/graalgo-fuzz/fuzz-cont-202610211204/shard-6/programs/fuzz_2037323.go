package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint }
func genFn0(p0 string) int16 {
    var v7 string = ("a" + p0)
    p0 = (("\n\t" + "foo bar") + ((p0 + "foo bar") + v7))
    var v8 string = p0
    _ = p0
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = 3.14159
    var v3 int = 0
    var v4 uint32 = 1
    var v5 string = "a"
    v6 := []uint32{v4}
    v5 = "hello"
    fmt.Println(v2)
    v4 = 332
    var v9 uint64 = 2
    v10 := v1[v5]
    if false {
        var v11 string = v5
        var v12 uint16 = 32767
        switch (v3 % ((0) | 1)) {
        case -9223372036854775808:
            var v13 string = v11
            _ = v13
        case 32767:
            v14 := GenS1{F0: 2, F1: 65535, F2: 556}
            _ = v14
        case 127:
            if len(v6) > 0 {
                v6[0] = 2147483647
            }
        }
        _ = v11
        _ = v12
    } else {
        v15 := GenS1{F0: v9, F1: 820, F2: uint(v9)}
        _ = v15
    }
    var v16 float32 = 0.1
    switch 469 {
    case 32768:
        fmt.Println(v5, v10)
    default:
        var v17 int32 = -2147483648
        _ = v17
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 int = fi18
        v2 = v2
        v20 := GenS1{F0: v9, F1: 980, F2: 647}
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 int64 = 278
    v22 := GenS1{F0: (v9 ^ (v9 & 4294967295)), F1: 42, F2: 65535}
    var v23 float32 = v16
    var v24 uint8 = 128
    v23 = float32(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
