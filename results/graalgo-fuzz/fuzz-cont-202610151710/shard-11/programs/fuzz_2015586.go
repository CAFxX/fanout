package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 bool }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = -19
    var v3 int8 = -1
    var v4 bool = true
    var v5 float32 = (1e10 / -2.25)
    var v6 int32 = 246
    if (128 <= 255) {
        fmt.Println(v6)
        _ = len(v1)
    } else {
        var v7 bool = (127 >= 43)
        switch ("a" + "hello") {
        case "foo bar":
            v8 := GenS1{F0: v3}
            _ = v8
        case "\n\t":
            var v9 uint16 = 371
            _ = v9
        default:
            fmt.Println(v2, v6, v6)
        }
        v6 = (v6 / (((v6 ^ (v6 >> ((341) & 63)))) | 1))
        _ = v7
    }
    v10 := GenS2{F0: 9223372036854775807, F1: 65535}
    v11 := GenS3{F0: 4294967295, F1: v6}
    v12 := GenS1{F0: (v3 - (v3 << ((0) & 63)))}
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := GenS4{F0: (32767 > -469)}
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 18; fi16++ {
            fmt.Println(fi16)
            v3 += v3
            var v17 uint = 120
            _ = fi16
            _ = v17
        }
        fmt.Println(v3, v6, v5)
        _ = fi15
    }
    v2 -= -627
    var v18 string = "foo bar"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
