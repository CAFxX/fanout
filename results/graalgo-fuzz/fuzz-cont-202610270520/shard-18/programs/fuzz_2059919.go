package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint8 }

func main() {
    v1 := []uint32{2, 170, 2, 134}
    var v2 int = -2147483648
    var v3 uint8 = 1
    var v4 uint8 = (v3 << ((65535) & 63))
    var v5 int32 = -185
    fmt.Println(v3, v3, v3)
    var v6 bool = (-2147483648 <= -2147483648)
    v3 += (v3 + 255)
    var v7 int32 = (v5 + (v5 % ((2147483647) | 1)))
    if (true || (!false)) {
        v5 = v7
        v5 += (v7 << ((2) & 63))
    }
    switch (-(v2)) {
    case -677:
        if (-919 == 0) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    case 392:
        if v6 {
            var v8 uint16 = 409
            _ = v8
        } else {
            var v9 int64 = 42
            v10 := GenS1{F0: 100, F1: 42, F2: v3}
            v2 = (v2 * (-(v2)))
            _ = v10
            _ = v9
        }
    }
    if v6 {
        v7 = v7
        var v11 uint64 = 4294967295
        var v12 uint32 = 65535
        _ = v11
        _ = v12
    } else {
        switch v2 {
        case -1:
            v3 -= v3
        case -465:
            v2 = 128
        case -77:
            v7 = 2147483647
        }
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 uint = 1
            _ = fi13
            _ = v14
        }
        fmt.Println(v5, v5, v3)
    }
    var v15 int64 = 0
    var v16 float64 = ((1.5 * (1.5 / 1.5)) / 1.0)
    switch "\n\t" {
    case "foo bar":
        v5 = v7
    case "a":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "hello":
        v16 = (v16 * 1e10)
    }
    switch "hello" {
    case "foo bar":
        fmt.Println(v2, v5, v6)
    default:
        v5 += 2147483647
    }
    v7 += v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
