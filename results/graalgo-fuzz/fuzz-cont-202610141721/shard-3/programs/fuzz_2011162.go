package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 int16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 2147483648
    v3 := make(map[uint32]int)
    var v4 int32 = -972
    v5 := make(map[uint32]int)
    v2 += (v2 & v2)
    v5[(^(v2))] = -458
    v4 = v4
    switch "\n\t" {
    case "":
        var v6 int = 229
        _ = v6
    case "\n\t":
        var v7 float64 = 0.0
        _ = v7
    default:
        fmt.Println(v4, v4)
    }
    v8 := GenS1{F0: -247}
    if false {
        if true {
            v2 = v2
            delete(v3, (v2 + (v2 - v2)))
        } else {
            v4 += (v4 % ((v4) | 1))
        }
        v9 := GenS2{F0: 65535}
        _ = v9
    }
    if true {
        v4 = v4
    } else {
        var v10 string = "a"
        _ = v10
    }
    var v11 uint64 = 0
    v12 := GenS2{F0: 18446744073709551615}
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 uint = 18446744073709551615
        v15 := GenS3{F0: (fi13 ^ (-(fi13))), F1: 42}
        _ = len(v3)
        _ = fi13
        _ = v14
        _ = v15
    }
    v11 = (^(v11))
    for fi16 := 0; fi16 < 18; fi16++ {
        v2 = (v2 + 635)
        if (((3.14159 == 1.5) && ((^(v11)) > v11)) && false) {
            var v17 uint32 = 2
            v18 := GenS3{F0: fi16, F1: 42}
            _ = v17
            _ = v18
        }
        _ = fi16
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        for fi20 := 0; fi20 < 8; fi20++ {
            v4 = v4
            fmt.Println(v11, fi19)
            var v21 bool = true
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    v3[330] = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
