package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 string; F1 int16; F2 int32 }
type GenS3 struct { F0 uint16; F1 string }

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]uint32)
    var v3 bool = true
    var v4 uint64 = 2
    v5 := make(map[uint32]uint32)
    switch "foo bar" {
    case "hello":
        var v6 uint64 = v1
        _ = v6
    case "foo bar":
        v1 = (v1 << ((65535) & 63))
    case "":
        v3 = (!false)
    }
    fmt.Println(v1, v1, v3)
    fmt.Println(v3, v1, v4)
    v7 := GenS1{F0: 58, F1: 0}
    v8 := []int{9223372036854775807, -2147483648, -454, 4294967295}
    var v9 uint8 = 157
    if v3 {
        for fi10 := 0; fi10 < 9; fi10++ {
            if ("a" == "") {
                v9 = v9
            } else {
                fmt.Println(v4)
            }
            v11 := GenS2{F0: "a", F1: -1, F2: int32(fi10)}
            _ = fi10
            _ = v11
        }
        v12 := GenS2{F0: "", F1: 623, F2: -1}
        _ = v12
    }
    if v3 {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        fmt.Println(v9, v9, v3)
        v13 := GenS1{F0: 32768, F1: 32767}
        _ = v13
    }
    v4 -= (v4 * v4)
    fmt.Println(v3, v9)
    fmt.Println(v3)
    fmt.Println(v1, v1, v4)
    v14 := GenS3{F0: 1, F1: "\n\t"}
    switch 512 {
    case 9223372036854775807:
        fmt.Println(v9)
    case 4294967295:
        delete(v5, 489)
    case -698:
        fmt.Println(v4, v4, v3)
    }
    if (v3 || ((32767 <= 42) && v3)) {
        v4 -= v4
    }
    v15 := GenS1{F0: 845, F1: 65535}
    v1 -= 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
