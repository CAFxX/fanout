package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int64 }

func main() {
    v1 := []uint32{65535}
    v2 := make(map[uint32]string)
    v3 := make(map[string]string)
    var v4 int32 = 0
    v4 = (-(v4))
    v1 = append(v1, 2147483647)
    v4 = -742
    delete(v3, "")
    v4 = (v4 >> ((18446744073709551615) & 63))
    switch "foo bar" {
    case "foo bar":
        var v5 int64 = 386
        _ = v5
    case "":
        switch -9223372036854775808 {
        case 201:
            if (-1 == 905) {
                if ((-1 < 0) || (false && (59 < 159))) {
                    v4 += v4
                }
                var v6 uint32 = 0
                _ = v6
            } else {
                fmt.Println(v4)
                var v7 float32 = 1e10
                _ = v7
            }
        case -2147483648:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    default:
        v8 := GenS1{F0: 0, F1: 862, F2: -618}
        _ = v8
    }
    switch -908 {
    case 4294967295:
        _ = len(v2)
    case 127:
        for fi9 := 0; fi9 < 16; fi9++ {
            fmt.Println(fi9, fi9)
            var v10 bool = ("\n\t" >= "a")
            _ = fi9
            _ = v10
        }
    }
    v4 = 42
    var v11 int32 = (v4 ^ -695)
    var v12 int8 = -2
    fmt.Println(v11, v12)
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 int8 = v12
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 bool = true
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    fmt.Println(v12)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
