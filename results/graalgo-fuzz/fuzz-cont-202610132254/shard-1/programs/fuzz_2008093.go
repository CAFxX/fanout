package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int64; F1 int32 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 4294967295
    v3 := make(map[string]int)
    v4 := []uint32{439, 2147483648, 65535}
    if (0 > 4294967295) {
        fmt.Println(v2)
    }
    switch 1 {
    case 0:
        for fi5 := 0; fi5 < 20; fi5++ {
            v4 = append(v4, 2)
            switch ("hello" + "") {
            case "foo bar":
                var v6 float32 = (1e10 + -2.25)
                _ = v6
            }
            var v7 bool = (1 < 127)
            _ = fi5
            _ = v7
        }
    case -9223372036854775808:
        if true {
            var v8 float32 = -2.25
            _ = v8
        } else {
            fmt.Println(v2)
            v9 := GenS1{F0: 42, F1: true}
            v2 += (v2 % ((1) | 1))
            _ = v9
        }
    default:
        v10 := GenS1{F0: 9223372036854775807, F1: true}
        _ = v10
    }
    v2 += v2
    switch ("foo bar" + "a") {
    case "a":
        _ = len(v4)
    case "foo bar":
        v2 = (v2 * v2)
    case "":
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 uint = 4294967295
            v2 -= uint64(v12)
            _ = fi11
            _ = v12
        }
    }
    v2 = 50
    v2 = v2
    v13 := GenS1{F0: -1, F1: true}
    v2 += (v2 - (v2 - (v2 | (v2 ^ v2))))
    v14 := GenS2{F0: -9223372036854775808, F1: 315}
    v15 := GenS1{F0: 0, F1: false}
    v2 = 64
    fmt.Println(v2)
    v4 = append(v4, 2147483648)
    v2 = v2
    var v16 string = "a"
    v17 := GenS2{F0: -1, F1: -2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
