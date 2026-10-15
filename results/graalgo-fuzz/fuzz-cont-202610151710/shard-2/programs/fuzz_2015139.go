package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint16 }

func main() {
    var v1 uint = 255
    v2 := make(map[uint32]string)
    var v3 int16 = 42
    v4 := []uint32{2}
    var v5 uint8 = 2
    var v6 uint16 = 802
    v5 += uint8(v6)
    switch "hello" {
    case "":
        v4 = append(v4, 0)
    case "\n\t":
        var v7 bool = true
        _ = v7
    case "a":
        var v8 int8 = 0
        _ = v8
    default:
        v6 -= (v6 << (((v1 + 42)) & 63))
    }
    fmt.Println(v1)
    fmt.Println(v3, v6)
    v2[42] = ""
    switch 9223372036854775807 {
    case 127:
        _ = len(v4)
    case -30:
        switch 9223372036854775807 {
        case 32767:
            var v9 string = "hello"
            _ = v9
        case 4294967295:
            for fi10 := 0; fi10 < 20; fi10++ {
                switch fi10 {
                case 2147483647:
                    delete(v2, 2147483648)
                case 128:
                    var v11 int8 = 100
                    _ = v11
                case -9223372036854775808:
                    v1 -= 936
                default:
                    if (true || true) {
                        v5 -= 127
                    }
                }
                v4 = append(v4, 44)
                v6 += (v6 << ((42) & 63))
                _ = fi10
            }
        case 309:
            v12 := GenS1{F0: -311, F1: 100, F2: (v6 - (^(v6)))}
            _ = v12
        }
    case -334:
        fmt.Println(v1, v6)
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        v1 = (v1 % ((v1) | 1))
        _ = fi13
    }
    var v14 float64 = 1.5
    switch -931 {
    case 4294967295:
        v15 := v2[uint32(v5)]
        _ = v15
    case 9223372036854775807:
        v1 = (v1 & v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
