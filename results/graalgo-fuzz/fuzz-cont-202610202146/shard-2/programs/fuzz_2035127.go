package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 uint = 42
    v2 := make(map[int]bool)
    var v3 int16 = -1
    var v4 int16 = (v3 - -10)
    var v5 int64 = 9223372036854775807
    _ = len(v2)
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := []int{-1, fi6, (fi6 ^ (-(fi6)))}
        _ = fi6
        _ = v7
    }
    v1 = (v1 & (v1 ^ v1))
    fmt.Println(v3)
    v4 -= (v4 >> (((v1 | (v1 + 65535))) & 63))
    v5 -= v5
    v8 := GenS1{F0: 255}
    v1 += 65535
    var v9 int8 = -15
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 uint64 = 18446744073709551615
        var v12 int8 = -128
        switch "foo bar" {
        case "hello":
            switch "hello" {
            case "\n\t":
                fmt.Println(v12, v1, v11)
            case "a":
                for fi13 := 0; fi13 < 12; fi13++ {
                    var v14 uint16 = 42
                    var v15 float32 = 0.1
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            case "foo bar":
                v1 -= 1
            default:
                v11 += v11
            }
        case "a":
            var v16 uint = v1
            _ = v16
        case "":
            v4 = (v4 >> ((42) & 63))
        default:
            v5 -= v5
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    if true {
        var v17 uint32 = 4294967295
        v3 = 1000
        delete(v2, 127)
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
