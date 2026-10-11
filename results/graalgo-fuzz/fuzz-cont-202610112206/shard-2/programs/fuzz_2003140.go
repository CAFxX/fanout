package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 string; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint }

func main() {
    var v1 uint32 = 65535
    v2 := make(map[int]bool)
    var v3 uint32 = 4294967295
    var v4 bool = (3 < 84)
    var v5 uint32 = (v1 >> ((18446744073709551615) & 63))
    for fi6 := 0; fi6 < 17; fi6++ {
        fmt.Println(v3, v1, v5)
        _ = fi6
    }
    var v7 int64 = 0
    for fi8 := 0; fi8 < 14; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            v4 = true
            _ = fi9
        }
        v1 -= (v5 >> ((252) & 63))
        v10 := GenS1{F0: fi8, F1: -1}
        _ = fi8
        _ = v10
    }
    var v11 uint8 = 181
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := GenS2{F0: ("\n\t" + "\n\t"), F1: -922, F2: (^(v11))}
        _ = fi12
        _ = v13
    }
    v11 += (v11 / (((v11 % (((v11 + v11)) | 1))) | 1))
    v5 += (v1 % ((650) | 1))
    switch "\n\t" {
    case "hello":
        v4 = (985 <= 202)
    default:
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 int16 = 950
            _ = fi14
            _ = v15
        }
    }
    var v16 uint8 = v11
    v4 = (!(1e10 <= (3.14159 * (1e10 / 1.0))))
    v17 := GenS3{F0: v16, F1: 97, F2: 0}
    var v18 uint = 798
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
