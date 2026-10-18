package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint16 }

func main() {
    var v1 uint64 = 1
    var v2 uint = 0
    var v3 int16 = -32768
    v4 := make(map[string]string)
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := GenS1{F0: 236}
        fmt.Println(v1, fi5, v3)
        _ = fi5
        _ = v6
    }
    if true {
        v7 := GenS1{F0: 1}
        _ = v7
    }
    v8 := GenS2{F0: 42}
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS3{F0: 0, F1: 65535}
        v3 = v3
        _ = fi9
        _ = v10
    }
    switch "hello" {
    case "a":
        var v11 float32 = -2.25
        _ = v11
    case "hello":
        v3 -= 42
    case "\n\t":
        v12 := GenS3{F0: 481, F1: 2}
        _ = v12
    default:
        if false {
            for fi13 := 0; fi13 < 1; fi13++ {
                v2 -= v2
                _ = fi13
            }
        }
    }
    v14 := GenS2{F0: 2147483647}
    var v15 int64 = 0
    v16 := []int{-357, -876, -9223372036854775808, 1}
    v1 += (v1 ^ v1)
    delete(v4, "foo bar")
    v1 = (v1 | v1)
    switch 32767 {
    case 9223372036854775807:
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 float32 = -2.25
            _ = fi17
            _ = v18
        }
    case 0:
        for fi19 := 0; fi19 < 5; fi19++ {
            v20 := v4["a"]
            var v21 uint8 = 42
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    _ = len(v16)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
