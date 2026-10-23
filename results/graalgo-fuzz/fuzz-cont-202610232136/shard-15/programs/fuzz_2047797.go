package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]string)
    var v3 uint32 = 1
    var v4 int64 = 1
    v5 := make(map[int]string)
    var v6 bool = (true && false)
    var v7 int32 = 184
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 float64 = 0.1
        _ = fi8
        _ = v9
    }
    switch "" {
    case "foo bar":
        var v10 uint64 = 223
        _ = v10
    case "hello":
        var v11 uint8 = 127
        _ = v11
    default:
        fmt.Println(v3, v4, v3)
    }
    fmt.Println(v7, v6)
    var v12 int64 = (v4 - v4)
    v6 = v6
    if (v3 <= (v3 & 4294967295)) {
        if true {
            var v13 int8 = 0
            switch 784 {
            case 4294967295:
                var v14 uint32 = v3
                _ = v14
            case -9223372036854775808:
                v15 := GenS1{F0: (v13 + v13), F1: ((!v6) && true)}
                _ = v15
            case 127:
                var v16 int32 = (v7 ^ v7)
                _ = v16
            }
            v17 := GenS1{F0: 88, F1: v6}
            _ = v13
            _ = v17
        }
    } else {
        v18 := GenS1{F0: 42, F1: (!v6)}
        v6 = true
        _ = v18
    }
    v6 = true
    v4 = -1
    for fi19 := 0; fi19 < 10; fi19++ {
        v4 -= 0
        v20 := GenS2{F0: 128}
        _ = fi19
        _ = v20
    }
    v3 = v3
    v12 = -962
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
