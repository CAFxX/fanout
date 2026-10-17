package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int32; F2 int64 }
type GenS3 struct { F0 int32 }

func main() {
    v1 := make(map[string]string)
    v2 := []string{}
    var v3 int32 = 42
    var v4 uint16 = 570
    switch "a" {
    case "a":
        fmt.Println(v4, v4)
    case "hello":
        var v5 int32 = (v3 ^ v3)
        _ = v5
    }
    v4 += (^(v4))
    var v6 bool = (1.5 > 1e10)
    if (-317 != 1000) {
        var v7 uint32 = 0
        _ = len(v1)
        _ = len(v2)
        _ = v7
    } else {
        fmt.Println(v3, v6, v6)
        if (255 > 255) {
            v2 = append(v2, (("foo bar" + "foo bar") + ""))
            var v8 uint32 = 258
            _ = v8
        } else {
            switch 2147483647 {
            case -105:
                v4 = v4
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    switch 0 {
    case 32768:
        v9 := GenS1{F0: 989}
        _ = v9
    default:
        var v10 uint32 = 563
        _ = v10
    }
    var v11 uint32 = 66
    if v6 {
        v12 := GenS2{F0: 42, F1: v3, F2: -397}
        v3 -= -2147483648
        _ = v12
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := GenS3{F0: (v3 >> ((42) & 63))}
        _ = fi13
        _ = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
