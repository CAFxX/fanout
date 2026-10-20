package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int32; F2 uint32 }
type GenS3 struct { F0 int }

func main() {
    var v1 uint32 = 1
    var v2 string = "\n\t"
    var v3 uint64 = 42
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := GenS1{F0: 32767}
        var v6 uint64 = v3
        v6 = 42
        _ = fi4
        _ = v5
        _ = v6
    }
    switch -516 {
    case 98:
        if true {
            v7 := []int{}
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = v7
        }
    case 512:
        var v8 int16 = 0
        _ = v8
    default:
        v9 := GenS2{F0: 1, F1: -612, F2: v1}
        _ = v9
    }
    v3 = v3
    switch v2 {
    case "foo bar":
        fmt.Println(v2)
    case "hello":
        v3 += (v3 * v3)
    }
    fmt.Println(v1, v3)
    var v10 int = 837
    var v11 uint32 = (v1 | v1)
    v12 := []int{}
    v2 = v2
    for fi13 := 0; fi13 < 11; fi13++ {
        fmt.Println(v3)
        v1 = (v11 << ((1) & 63))
        v1 = (v1 << ((18446744073709551615) & 63))
        _ = fi13
    }
    v14 := GenS2{F0: -105, F1: -2147483648, F2: v1}
    fmt.Println(v10, v10, v1)
    fmt.Println(v1, v11)
    switch v10 {
    case 0:
        v2 = "foo bar"
    case -9223372036854775808:
        var v15 int = (v10 | -903)
        _ = v15
    case 128:
        v16 := GenS3{F0: 32768}
        _ = v16
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        v10 += v10
        if len(v12) > 0 {
            v12[0] = 1
        }
        _ = fi17
    }
    var v18 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v18
    _ = v2
    _ = v3
}
