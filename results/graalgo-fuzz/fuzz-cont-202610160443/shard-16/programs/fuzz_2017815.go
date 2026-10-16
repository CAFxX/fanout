package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint64 }

func main() {
    var v1 int32 = 2147483647
    var v2 int = 1
    v3 := make(map[int]uint32)
    var v4 int = -596
    var v5 int64 = 411
    v4 = v4
    v1 += -2147483648
    var v6 uint16 = 32767
    var v7 int8 = -39
    if true {
        switch "foo bar" {
        case "\n\t":
            var v8 uint16 = 476
            _ = v8
        default:
            switch v2 {
            case 9223372036854775807:
                var v9 uint = 2
                _ = v9
            case 2147483647:
                for fi10 := 0; fi10 < 5; fi10++ {
                    var v11 int32 = (v1 >> ((65535) & 63))
                    _ = fi10
                    _ = v11
                }
            case -55:
                v12 := v3[(v4 << ((1) & 63))]
                _ = v12
            }
        }
        fmt.Println(v4, v6)
        v6 = (v6 << ((499) & 63))
    } else {
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 uint32 = 0
            _ = fi13
            _ = v14
        }
        v7 = v7
    }
    switch v4 {
    case 0:
        var v15 int64 = -1
        _ = v15
    case 42:
        var v16 int = v4
        _ = v16
    }
    v17 := GenS1{F0: (^(v6)), F1: 204, F2: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
