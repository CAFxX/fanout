package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int32; F1 int32 }

func main() {
    var v1 uint16 = 255
    v2 := []string{"a", "hello"}
    var v3 int64 = 233
    var v4 string = "\n\t"
    var v5 int = 852
    v1 = (v1 % (((v1 - v1)) | 1))
    var v6 uint = 65535
    v7 := GenS1{F0: (v5 & v5), F1: false}
    v3 -= v3
    for fi8 := 0; fi8 < 15; fi8++ {
        switch v5 {
        case 32767:
            v9 := GenS1{F0: 578, F1: false}
            _ = v9
        default:
            var v10 uint16 = 255
            _ = v10
        }
        _ = fi8
    }
    switch -1 {
    case -2147483648:
        switch (v5 | 32768) {
        case -2147483648:
            fmt.Println(v6, v4, v1)
        case -470:
            fmt.Println(v1)
        case 436:
            fmt.Println(v1, v1, v3)
        default:
            var v11 float32 = 1.5
            _ = v11
        }
    default:
        fmt.Println(v1, v1, v4)
    }
    v4 = v4
    fmt.Println(v4)
    switch (v5 / (((v5 | 1)) | 1)) {
    case 1:
        var v12 int32 = -2147483648
        _ = v12
    case 0:
        v5 = v5
    case 32767:
        v13 := GenS2{F0: 42, F1: 152}
        _ = v13
    default:
        var v14 uint8 = 1
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
