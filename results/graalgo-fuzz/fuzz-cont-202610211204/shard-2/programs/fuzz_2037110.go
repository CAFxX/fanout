package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int8; F1 int64 }

func main() {
    v1 := make(map[int]bool)
    v2 := []string{}
    v3 := make(map[int]string)
    v4 := make(map[uint32]int)
    var v5 uint16 = 1
    v5 = (v5 & v5)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v6 uint64 = 4294967295
    if (3.14159 < 1.5) {
        fmt.Println(v6, v6)
        v6 = v6
    }
    switch 32767 {
    case 2147483647:
        v6 -= uint64(v5)
    case 955:
        v7 := GenS1{F0: -9223372036854775808}
        _ = v7
    default:
        v5 += (v5 | v5)
    }
    switch "hello" {
    case "\n\t":
        if (!(1 <= 128)) {
            v8 := GenS2{F0: false, F1: -357}
            _ = v8
        }
    case "foo bar":
        fmt.Println(v6, v6)
    case "a":
        fmt.Println(v5, v5)
    default:
        var v9 uint32 = 0
        _ = v9
    }
    fmt.Println(v5, v5)
    var v10 uint16 = v5
    for fi11 := 0; fi11 < 18; fi11++ {
        for fi12 := 0; fi12 < 20; fi12++ {
            v6 += (v6 % (((v6 << ((4294967295) & 63))) | 1))
            _ = fi12
        }
        v5 -= v10
        _ = fi11
    }
    v13 := GenS1{F0: -9223372036854775808}
    v14 := GenS3{F0: 0, F1: 0}
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 uint64 = (v6 << ((114) & 63))
        v10 = v10
        _ = fi15
        _ = v16
    }
    var v17 float32 = float32((0.0 + 3.14159))
    var v18 float32 = v17
    delete(v1, 181)
    for fi19 := 0; fi19 < 11; fi19++ {
        fmt.Println(v17, v6, fi19)
        _ = fi19
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
