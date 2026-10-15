package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 int32 = 634
    var v2 uint8 = 127
    v3 := make(map[int]uint32)
    delete(v3, -37)
    var v4 uint16 = 42
    switch (("foo bar" + ("hello" + "foo bar")) + "\n\t") {
    case "a":
        fmt.Println(v2, v1)
    case "":
        for fi5 := 0; fi5 < 15; fi5++ {
            v1 += 166
            fmt.Println(v2, v1)
            _ = fi5
        }
    }
    v6 := GenS1{F0: 4294967295, F1: false, F2: 255}
    v7 := []int{724}
    v8 := GenS2{F0: 218}
    var v9 bool = false
    v9 = false
    switch -418 {
    case 127:
        var v10 uint64 = 18446744073709551615
        _ = v10
    default:
        v4 -= (v4 + (v4 >> ((65535) & 63)))
    }
    v2 = v2
    var v11 uint = 1
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 int16 = 0
        v14 := v3[(fi12 << ((65535) & 63))]
        var v15 float64 = (0.1 + ((-2.25 - 1.5) - 0.0))
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 int32 = v1
    var v17 int64 = 9223372036854775807
    v18 := GenS3{F0: 440}
    v16 = (v1 >> (((v11 % (((v11 ^ 107)) | 1))) & 63))
    v19 := GenS3{F0: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
