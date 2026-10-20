package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int64; F2 bool }

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[string]uint32)
    var v3 int64 = -1
    var v4 int = -1
    if ((0.0 - float32((0.0 * 3.14159))) != (float32(-2.25) / 1.0)) {
        var v5 uint16 = 32767
        _ = v5
    } else {
        v6 := GenS1{F0: 1, F1: 0}
        v3 += v3
        var v7 uint = 42
        _ = v6
        _ = v7
    }
    if (!(!true)) {
        if true {
            for fi8 := 0; fi8 < 10; fi8++ {
                var v9 uint8 = 128
                _ = fi8
                _ = v9
            }
            switch v4 {
            case 4294967295:
                if true {
                    v10 := GenS2{F0: 234, F1: v3, F2: false}
                    _ = v10
                }
            }
        } else {
            var v11 int64 = v3
            var v12 int = 0
            var v13 int64 = (v3 * 9223372036854775807)
            _ = v11
            _ = v12
            _ = v13
        }
        fmt.Println(v1, v4, v1)
        if true {
            var v14 int8 = -1
            fmt.Println(v3)
            _ = v14
        }
    } else {
        v1 = v1
        v4 = v4
    }
    var v15 uint16 = 99
    var v16 bool = false
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 int32 = v1
        switch ("foo bar" + ("foo bar" + "hello")) {
        case "hello":
            var v19 int32 = (v1 / ((42) | 1))
            _ = v19
        }
        var v20 bool = (4294967295 >= 639)
        _ = fi17
        _ = v18
        _ = v20
    }
    v1 = v1
    v21 := v2["hello"]
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
