package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := []uint32{}
    var v2 int = 17
    var v3 int8 = -1
    var v4 int = v2
    v5 := GenS1{F0: 2147483647}
    if len(v1) > 0 {
        v1[0] = 65535
    }
    v4 += (v4 % ((v2) | 1))
    fmt.Println(v4)
    switch "a" {
    case "hello":
        switch (v2 << ((2) & 63)) {
        case 692:
            switch (v4 << ((18446744073709551615) & 63)) {
            case 127:
                v6 := GenS1{F0: 241}
                _ = v6
            case 1:
                var v7 uint64 = 361
                _ = v7
            case 2147483647:
                v3 = (v3 >> ((0) & 63))
            default:
                v3 += (v3 & 100)
            }
        }
    case "":
        if true {
            var v8 string = (((("hello" + "hello") + "\n\t") + "\n\t") + "")
            _ = v8
        }
    default:
        v4 -= v2
    }
    v9 := GenS1{F0: 1}
    if true {
        switch "\n\t" {
        case "a":
            v2 = (v4 ^ (v4 & (v2 + v4)))
        case "\n\t":
            v10 := GenS2{F0: 32767}
            _ = v10
        }
    }
    if len(v1) > 0 {
        v1[0] = 0
    }
    v3 -= (^(v3))
    switch v2 {
    case -9223372036854775808:
        var v11 uint8 = 42
        _ = v11
    default:
        var v12 int64 = 0
        _ = v12
    }
    switch v2 {
    case 4294967295:
        fmt.Println(v4, v4, v4)
    }
    v2 = v2
    for fi13 := 0; fi13 < 20; fi13++ {
        v1 = append(v1, 2147483647)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 int64 = 9223372036854775807
            var v16 uint32 = 175
            v17 := GenS2{F0: 32767}
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        _ = fi13
    }
    v4 += (v4 / (((v4 * -801)) | 1))
    fmt.Println(v4)
    v2 = 32768
    var v18 int = (v2 >> ((4294967295) & 63))
    v18 -= v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
