package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int; F1 bool; F2 int32 }

func main() {
    var v1 int64 = 1
    v2 := []uint32{2147483648, 2147483648, 2, 0}
    v3 := make(map[uint32]uint32)
    if (934 > 1) {
        var v4 float64 = 1.5
        var v5 uint16 = 946
        _ = v4
        _ = v5
    }
    v1 += (v1 << ((4294967295) & 63))
    var v6 uint16 = 0
    fmt.Println(v6, v1)
    v6 = v6
    if (true || (2147483648 > 2147483648)) {
        fmt.Println(v1)
        var v7 uint32 = 65535
        _ = v7
    } else {
        delete(v3, 446)
        for fi8 := 0; fi8 < 4; fi8++ {
            if true {
                v6 += (v6 / (((v6 % (((v6 << ((4294967295) & 63))) | 1))) | 1))
            }
            v9 := GenS1{F0: v1, F1: 4294967295}
            v6 = 2
            _ = fi8
            _ = v9
        }
    }
    var v10 string = "hello"
    switch (("\n\t" + "foo bar") + v10) {
    case "foo bar":
        var v11 int8 = 0
        _ = v11
    case "a":
        v12 := GenS1{F0: 9223372036854775807, F1: 197}
        _ = v12
    case "\n\t":
        var v13 uint = 18446744073709551615
        _ = v13
    }
    switch "hello" {
    case "\n\t":
        var v14 uint32 = 42
        _ = v14
    case "foo bar":
        for fi15 := 0; fi15 < 7; fi15++ {
            v16 := GenS2{F0: 1, F1: 4294967295, F2: 0}
            fmt.Println(v6, fi15)
            if true {
                fmt.Println(v1)
                var v17 uint = 58
                for fi18 := 0; fi18 < 14; fi18++ {
                    var v19 uint8 = 255
                    fmt.Println(v19)
                    v10 = v10
                    v2 = append(v2, 42)
                    _ = fi18
                    _ = v19
                }
                _ = v17
            }
            _ = fi15
            _ = v16
        }
    default:
        v20 := GenS3{F0: 0, F1: true, F2: 42}
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
}
