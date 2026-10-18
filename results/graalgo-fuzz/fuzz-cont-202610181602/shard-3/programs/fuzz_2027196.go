package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 uint32 = 10
    var v2 uint64 = 18446744073709551615
    var v3 uint8 = 69
    var v4 uint = 1
    if true {
        var v5 int32 = 1
        _ = v5
    } else {
        for fi6 := 0; fi6 < 19; fi6++ {
            var v7 string = "hello"
            _ = fi6
            _ = v7
        }
    }
    v3 -= v3
    for fi8 := 0; fi8 < 19; fi8++ {
        switch "hello" {
        case "":
            v4 = (v4 ^ v4)
        case "a":
            v9 := GenS1{F0: -1}
            _ = v9
        default:
            v3 -= (v3 ^ v3)
        }
        switch "a" {
        case "a":
            var v10 float32 = float32(3.14159)
            _ = v10
        default:
            var v11 uint32 = (v1 & v1)
            _ = v11
        }
        _ = fi8
    }
    v12 := GenS1{F0: 810}
    fmt.Println(v4, v1)
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 int32 = -1
        _ = fi13
        _ = v14
    }
    v15 := []int{9223372036854775807, 127, 32767}
    v3 -= uint8(v4)
    switch -1 {
    case 9223372036854775807:
        if true {
            v4 += v4
            v2 -= v2
        }
    case 386:
        v2 += 1
    case 0:
        fmt.Println(v4, v3)
    default:
        var v16 int16 = 51
        _ = v16
    }
    fmt.Println(v1, v4)
    v4 += (^(v4))
    var v17 int = 128
    var v18 uint8 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
