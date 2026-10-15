package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 bool; F1 int8; F2 int32 }

func main() {
    var v1 uint64 = 0
    var v2 float64 = ((-2.25 + 1.5) - 0.0)
    v3 := make(map[int]string)
    delete(v3, 9223372036854775807)
    v3[128] = ""
    switch 249 {
    case 0:
        var v4 uint8 = 128
        _ = v4
    case 32767:
        var v5 uint16 = 255
        _ = v5
    default:
        fmt.Println(v1, v1)
    }
    var v6 float32 = ((0.1 / 1.5) - (3.14159 / 1.5))
    v7 := []int{-1000, -1, 42, 0}
    _ = len(v3)
    v6 -= 3.14159
    v2 = v2
    v8 := GenS1{F0: 9223372036854775807}
    for fi9 := 0; fi9 < 17; fi9++ {
        switch "" {
        case "hello":
            switch -556 {
            case -2147483648:
                var v10 int8 = -128
                _ = v10
            case 9223372036854775807:
                var v11 string = ""
                _ = v11
            case 0:
                var v12 int16 = 42
                _ = v12
            }
        case "":
            fmt.Println(v2, fi9, v2)
        default:
            switch (fi9 ^ (fi9 % (((fi9 | (fi9 - (fi9 % ((fi9) | 1))))) | 1))) {
            case 2147483647:
                var v13 uint64 = 2
                _ = v13
            default:
                var v14 string = (("\n\t" + "foo bar") + "")
                _ = v14
            }
        }
        fmt.Println(v1)
        _ = fi9
    }
    v15 := GenS2{F0: (2147483647 == 666), F1: 127}
    var v16 int = 1
    delete(v3, v16)
    v2 -= ((v2 / (v2 * v2)) * 0.1)
    if true {
        _ = len(v3)
    }
    delete(v3, v16)
    switch (v16 ^ v16) {
    case 128:
        v3[v16] = (("a" + "a") + "")
    case 2147483647:
        var v17 int = v16
        _ = v17
    default:
        v18 := GenS3{F0: false, F1: 100, F2: -2147483648}
        _ = v18
    }
    v16 = (v16 % ((926) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
