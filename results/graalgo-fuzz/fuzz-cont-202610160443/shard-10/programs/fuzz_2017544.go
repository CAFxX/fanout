package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int64 }
type GenS2 struct { F0 int64; F1 uint8; F2 int64 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 int8; F1 int16 }

func main() {
    var v1 uint16 = 255
    var v2 bool = (1e10 <= 0.0)
    var v3 int16 = 42
    var v4 uint = 0
    var v5 string = ""
    var v6 int32 = 456
    v7 := []int{-9223372036854775808, -975, -2147483648, 2147483647}
    switch -827 {
    case 128:
        var v8 uint8 = 1
        _ = v8
    }
    v5 = v5
    v9 := GenS1{F0: 127, F1: -1, F2: 0}
    v7 = append(v7, -2147483648)
    _ = len(v7)
    v10 := GenS2{F0: -9223372036854775808, F1: 42, F2: -270}
    v11 := GenS3{F0: (v6 | (v6 << ((v4) & 63))), F1: 42}
    var v12 int64 = 42
    for fi13 := 0; fi13 < 9; fi13++ {
        if true {
            if true {
                _ = len(v7)
                if v2 {
                    var v14 float32 = 3.14159
                    _ = v14
                } else {
                    var v15 bool = (-1 >= 90)
                    fmt.Println(fi13)
                    _ = v15
                }
            } else {
                switch (fi13 | 127) {
                case -411:
                    var v16 int16 = (-(v3))
                    _ = v16
                case -9223372036854775808:
                    v17 := GenS3{F0: -863, F1: 32767}
                    _ = v17
                case 2147483647:
                    v18 := GenS3{F0: v6, F1: (v1 << ((647) & 63))}
                    _ = v18
                }
            }
            switch "a" {
            case "a":
                if len(v7) > 0 {
                    v7[0] = -2147483648
                }
            default:
                v3 = -32768
            }
            var v19 string = "a"
            _ = v19
        } else {
            var v20 int64 = v12
            _ = v20
        }
        v21 := GenS4{F0: 100, F1: 0}
        _ = fi13
        _ = v21
    }
    switch "\n\t" {
    case "a":
        var v22 bool = v2
        _ = v22
    case "\n\t":
        fmt.Println(v12)
    case "":
        var v23 int8 = 100
        _ = v23
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        var v25 bool = v2
        fmt.Println(v1, v12, v12)
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
