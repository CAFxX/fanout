package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int32; F1 uint64; F2 int }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int }

func main() {
    var v1 int = 831
    var v2 uint16 = 255
    v3 := make(map[int]int)
    var v4 uint16 = v2
    v1 = (v1 * 127)
    switch "\n\t" {
    case "\n\t":
        v5 := GenS1{F0: 2147483647, F1: "\n\t"}
        _ = v5
    case "hello":
        if (((float32((3.14159 - 3.14159)) / 0.1) / 1.0) == (0.0 + -2.25)) {
            v4 += v2
            v6 := GenS2{F0: 0, F1: 2, F2: v1}
            switch v1 {
            case -1:
                v4 = v4
            case -14:
                var v7 bool = (false || false)
                _ = v7
            case 128:
                fmt.Println(v1)
            default:
                var v8 float64 = 3.14159
                _ = v8
            }
            _ = v6
        } else {
            fmt.Println(v4)
        }
    default:
        fmt.Println(v4, v4)
    }
    var v9 uint8 = 2
    var v10 uint64 = 2
    if true {
        var v11 uint8 = (v9 << ((27) & 63))
        _ = v11
    }
    var v12 int16 = 1000
    v13 := GenS1{F0: 217, F1: "a"}
    switch (v1 & (v1 + v1)) {
    case 9223372036854775807:
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 int32 = -1
            var v16 float32 = 0.0
            _ = fi14
            _ = v15
            _ = v16
        }
    case 0:
        var v17 int16 = (v12 % (((v12 ^ int16(v2))) | 1))
        _ = v17
    }
    v1 += (v1 >> ((98) & 63))
    v1 += 32767
    v12 = v12
    switch "hello" {
    case "hello":
        switch v1 {
        case 9223372036854775807:
            v1 = v1
        }
    default:
        v18 := GenS3{F0: int(v4), F1: (!true)}
        _ = v18
    }
    var v19 bool = false
    v20 := v3[(-(v1))]
    var v21 float32 = 1e10
    v4 += (^(v2))
    v22 := GenS4{F0: -327}
    if ("a" != ("\n\t" + "\n\t")) {
        v4 = (v2 * (v2 | v4))
        var v23 int8 = -1
        _ = v23
    } else {
        v9 = v9
        var v24 string = "hello"
        for fi25 := 0; fi25 < 10; fi25++ {
            v26 := GenS3{F0: v20, F1: false}
            _ = fi25
            _ = v26
        }
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v9
}
