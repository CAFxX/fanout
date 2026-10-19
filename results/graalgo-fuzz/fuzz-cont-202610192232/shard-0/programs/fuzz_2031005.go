package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{684, 818, 4294967295, 1}
    var v3 int64 = -9223372036854775808
    v4 := make(map[uint32]string)
    var v5 uint32 = 433
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 string = "hello"
        v8 := GenS1{F0: v3}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint16 = 994
    fmt.Println(v9, v9)
    for fi10 := 0; fi10 < 2; fi10++ {
        v3 = (v3 ^ (v3 - (v3 + (v3 % ((349) | 1)))))
        _ = fi10
    }
    switch 9223372036854775807 {
    case 9223372036854775807:
        switch "\n\t" {
        case "a":
            v3 = (v3 & (^(v3)))
        case "hello":
            var v11 float32 = (1.5 + 1.5)
            _ = v11
        default:
            for fi12 := 0; fi12 < 12; fi12++ {
                if len(v2) > 0 {
                    v2[0] = (v5 ^ uint32(v9))
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v13 int = fi12
                _ = fi12
                _ = v13
            }
        }
    case -9223372036854775808:
        fmt.Println(v9)
    default:
        if len(v2) > 0 {
            v2[0] = 742
        }
    }
    v9 = v9
    v9 += (v9 << ((1) & 63))
    fmt.Println(v3)
    v3 -= v3
    if false {
        v9 -= v9
        v5 = (v5 ^ v5)
        v5 = (v5 - 898)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
