package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint; F2 uint64 }

func main() {
    var v1 uint8 = 1
    var v2 int = -9223372036854775808
    v3 := []string{("hello" + ""), ("a" + "a")}
    var v4 int64 = 42
    for fi5 := 0; fi5 < 8; fi5++ {
        switch "foo bar" {
        case "a":
            fmt.Println(fi5)
        case "foo bar":
            v1 = v1
        case "":
            v4 -= (^(v4))
        default:
            v1 = (v1 << ((408) & 63))
        }
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 uint8 = (v1 & v1)
            _ = fi6
            _ = v7
        }
        var v8 uint32 = 1
        _ = fi5
        _ = v8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        v2 = fi9
        v2 = 213
        _ = fi9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        fmt.Println(v2)
        _ = fi10
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        fmt.Println(v1, v1, v2)
        for fi12 := 0; fi12 < 10; fi12++ {
            v13 := GenS1{F0: 0}
            v2 += 128
            v14 := GenS2{F0: (v1 + (^(v1)))}
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 int = int(v4)
        _ = fi11
        _ = v15
    }
    v2 += (-(v2))
    var v16 uint32 = 65535
    switch v2 {
    case -1:
        v17 := GenS3{F0: v4, F1: 775, F2: 0}
        _ = v17
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        v3 = append(v3, ("foo bar" + ("hello" + ("foo bar" + ("" + "foo bar")))))
        var v19 uint8 = (v1 & v1)
        v3 = append(v3, "hello")
        _ = fi18
        _ = v19
    }
    var v20 int8 = 0
    var v21 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
