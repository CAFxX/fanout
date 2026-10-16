package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[int]int)
    var v2 int16 = -1
    var v3 int16 = (v2 ^ v2)
    var v4 float32 = 1e10
    var v5 uint = 396
    fmt.Println(v3, v2, v2)
    for fi6 := 0; fi6 < 12; fi6++ {
        if true {
            v5 = 2
            v5 += v5
        }
        _ = fi6
    }
    var v7 string = "hello"
    v8 := []int{}
    if (0 == (v2 - (v2 >> ((v5) & 63)))) {
        var v9 int64 = 593
        fmt.Println(v7)
        switch -1 {
        case 127:
            for fi10 := 0; fi10 < 17; fi10++ {
                v5 -= v5
                _ = fi10
            }
        case -9223372036854775808:
            var v11 uint16 = 464
            _ = v11
        default:
            v12 := GenS1{F0: 100, F1: 42}
            _ = v12
        }
        _ = v9
    }
    v5 -= (v5 % ((726) | 1))
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 uint8 = 127
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        for fi16 := 0; fi16 < 15; fi16++ {
            v17 := GenS2{F0: 255}
            v1[(-(fi16))] = (fi16 * -9223372036854775808)
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v18 := GenS1{F0: 85, F1: 0}
    for fi19 := 0; fi19 < 10; fi19++ {
        for fi20 := 0; fi20 < 12; fi20++ {
            v3 += v2
            var v21 int64 = 842
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    v5 += (v5 >> ((v5) & 63))
    v7 = (((v7 + (v7 + v7)) + (v7 + v7)) + ((v7 + "a") + ("a" + v7)))
    var v22 float32 = (v4 - 1.5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
