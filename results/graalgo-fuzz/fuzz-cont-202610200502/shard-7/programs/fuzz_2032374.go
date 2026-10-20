package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint8; F2 int32 }

func main() {
    var v1 int16 = 1
    var v2 int32 = 662
    v3 := []string{"", ("hello" + "")}
    v4 := make(map[string]uint32)
    v5 := []uint32{2, 4294967295, 360, 740}
    var v6 int8 = -1
    var v7 float64 = 1e10
    fmt.Println(v7, v7)
    fmt.Println(v6, v1)
    if (!false) {
        fmt.Println(v7, v2)
        v6 += v6
        var v8 float32 = 3.14159
        _ = v8
    }
    var v9 string = "hello"
    switch (v9 + "\n\t") {
    case "foo bar":
        v10 := GenS1{F0: -9223372036854775808, F1: (v1 << ((0) & 63))}
        _ = v10
    case "":
        v6 += v6
    case "\n\t":
        switch 191 {
        case 0:
            fmt.Println(v7)
        }
    }
    var v11 int64 = 844
    switch "hello" {
    case "hello":
        v11 -= v11
    case "\n\t":
        for fi12 := 0; fi12 < 5; fi12++ {
            switch (v9 + "\n\t") {
            case "":
                v13 := GenS2{F0: fi12}
                _ = v13
            case "hello":
                for fi14 := 0; fi14 < 2; fi14++ {
                    var v15 uint32 = 4294967295
                    v9 = ("foo bar" + (v9 + v9))
                    v16 := GenS3{F0: 1, F1: 255, F2: v2}
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case "\n\t":
                fmt.Println(v9, v7, v11)
            }
            _ = fi12
        }
    default:
        var v17 int = -1
        _ = v17
    }
    v1 = (v1 ^ (v1 & v1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
