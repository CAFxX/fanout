package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 bool; F1 int32 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    var v3 uint16 = 255
    var v4 uint16 = 893
    var v5 uint8 = 172
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            v8 := v1["hello"]
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    v9 := GenS1{F0: 469}
    for fi10 := 0; fi10 < 16; fi10++ {
        if true {
            v4 += v4
        }
        fmt.Println(v4, v4, v4)
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := v1[""]
            var v13 uint64 = 4294967295
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    fmt.Println(v4, v4)
    if true {
        v14 := GenS1{F0: 42}
        v4 = v3
        var v15 bool = (!true)
        _ = v14
        _ = v15
    }
    fmt.Println(v3)
    v16 := GenS2{F0: -9223372036854775808, F1: true}
    switch "hello" {
    case "":
        for fi17 := 0; fi17 < 2; fi17++ {
            for fi18 := 0; fi18 < 13; fi18++ {
                v19 := GenS1{F0: 18446744073709551615}
                v3 = v4
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
    case "hello":
        v5 = (v5 % ((v5) | 1))
    case "foo bar":
        _ = len(v1)
    default:
        v20 := GenS3{F0: (!true), F1: 2147483647}
        _ = v20
    }
    var v21 uint64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
