package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }

func main() {
    v1 := make(map[int]int)
    var v2 int8 = 0
    var v3 string = (("a" + "") + "foo bar")
    v4 := make(map[uint32]int)
    v5 := GenS1{F0: 42, F1: -213}
    v3 = v3
    var v6 uint16 = 708
    v2 = v2
    v7 := GenS1{F0: 42, F1: 526}
    fmt.Println(v3, v2)
    var v8 string = "foo bar"
    switch v8 {
    case "\n\t":
        var v9 float32 = (float32(3.14159) * 3.14159)
        _ = v9
    default:
        if (624 < -965) {
            var v10 float32 = 1.5
            _ = v10
        }
    }
    v6 = v6
    v8 = ""
    v6 -= (v6 % ((v6) | 1))
    for fi11 := 0; fi11 < 13; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            delete(v4, 612)
            _ = fi12
        }
        var v13 float32 = (1.5 + 0.1)
        _ = fi11
        _ = v13
    }
    v14 := GenS1{F0: 138, F1: 934}
    v15 := GenS1{F0: 42, F1: 9223372036854775807}
    for fi16 := 0; fi16 < 4; fi16++ {
        _ = len(v1)
        if true {
            v2 = v2
        }
        _ = fi16
    }
    var v17 uint64 = 0
    v17 = v17
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
