package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 bool }
type GenS2 struct { F0 uint16; F1 bool; F2 uint8 }

func main() {
    var v1 int32 = -154
    v2 := []uint32{4294967295, 2147483648}
    var v3 int16 = 996
    var v4 int32 = v1
    if (-889 > 9223372036854775807) {
        var v5 bool = (false && (("hello" + "\n\t") >= (("foo bar" + "foo bar") + "a")))
        v6 := GenS1{F0: 255, F1: 69, F2: v5}
        v7 := GenS1{F0: 128, F1: 56, F2: false}
        _ = v5
        _ = v6
        _ = v7
    }
    v8 := GenS2{F0: 427, F1: ("a" >= "a"), F2: 63}
    switch "\n\t" {
    case "":
        v2 = append(v2, 1)
    default:
        var v9 int8 = 1
        _ = v9
    }
    var v10 uint16 = 2
    var v11 uint64 = 1
    v10 = (v10 << ((1) & 63))
    for fi12 := 0; fi12 < 20; fi12++ {
        _ = len(v2)
        var v13 int32 = v1
        _ = fi12
        _ = v13
    }
    var v14 int = -2147483648
    v10 = (^(v10))
    v14 -= v14
    fmt.Println(v14)
    if (4294967295 < 1) {
        for fi15 := 0; fi15 < 13; fi15++ {
            v1 += (v4 | (v1 / (((v1 / (((v1 - 1)) | 1))) | 1)))
            _ = fi15
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
