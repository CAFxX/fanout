package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint32 }

func main() {
    var v1 float64 = (3.14159 * 3.14159)
    v2 := []uint32{2147483647}
    v3 := make(map[string]int)
    v4 := []uint32{2147483647}
    v5 := make(map[string]uint32)
    var v6 uint16 = 32767
    if true {
        switch 63 {
        case 127:
            var v7 int = 4294967295
            _ = v7
        default:
            v8 := v5[""]
            _ = v8
        }
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 uint = 2
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: 671, F1: 2, F2: 2147483648}
        _ = v11
    }
    v1 -= v1
    _ = len(v2)
    var v12 uint16 = 137
    for fi13 := 0; fi13 < 5; fi13++ {
        fmt.Println(v6)
        _ = fi13
    }
    if (-2147483648 > 2147483647) {
        v12 -= v12
        v5[(("hello" + "a") + "hello")] = 802
        fmt.Println(v12, v12)
    } else {
        fmt.Println(v6)
    }
    if (349 >= 0) {
        var v14 uint8 = 1
        _ = v14
    }
    var v15 bool = (128 >= 128)
    v16 := v3[""]
    if (!true) {
        switch v16 {
        case 116:
            v12 -= v12
        case 32768:
            v17 := GenS1{F0: 4294967295, F1: 863, F2: 2147483647}
            _ = v17
        case -311:
            v18 := GenS1{F0: 65535, F1: 18446744073709551615, F2: 981}
            _ = v18
        default:
            fmt.Println(v1, v15)
        }
        v6 = 955
        fmt.Println(v15)
    } else {
        v19 := GenS1{F0: 255, F1: 42, F2: 42}
        var v20 int64 = 1
        _ = v19
        _ = v20
    }
    var v21 int = v16
    switch 759 {
    case -9223372036854775808:
        v15 = true
    case 32767:
        var v22 uint32 = 4294967295
        _ = v22
    }
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
