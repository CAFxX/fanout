package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint64 }

func main() {
    var v1 float32 = 1e10
    v2 := make(map[string]uint32)
    v3 := []string{}
    var v4 int = -481
    if ((!((3.14159 <= 1e10) || ((false || false) && (!false)))) && true) {
        var v5 uint8 = 164
        _ = v5
    } else {
        switch "" {
        case "a":
            v1 += v1
        case "\n\t":
            v1 += v1
        case "hello":
            var v6 int = -2147483648
            _ = v6
        default:
            v7 := v2["a"]
            _ = v7
        }
        fmt.Println(v4, v1)
        switch v4 {
        case 9223372036854775807:
            v1 = v1
        }
    }
    var v8 uint = 58
    fmt.Println(v8)
    var v9 uint64 = 18446744073709551615
    v10 := GenS1{F0: 2147483647}
    v11 := GenS2{F0: v9, F1: 4294967295, F2: 18446744073709551615}
    v9 -= v9
    fmt.Println(v8)
    v12 := v2[(("foo bar" + ("hello" + ("foo bar" + ("foo bar" + ("foo bar" + ""))))) + "foo bar")]
    switch (v4 << (((v8 ^ (v8 / (((^(v8))) | 1)))) & 63)) {
    case 128:
        v9 = (^(v9))
    case 4294967295:
        v3 = append(v3, (("" + "a") + ""))
    case 2147483647:
        v9 = v9
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 int8 = -45
        _ = fi13
        _ = v14
    }
    if (!(2 == v12)) {
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 uint = 0
            v17 := GenS1{F0: 651}
            var v18 int16 = -386
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        if true {
            var v19 int32 = -1
            _ = v19
        }
    } else {
        fmt.Println(v9, v9, v9)
        _ = len(v2)
    }
    v9 += v9
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
