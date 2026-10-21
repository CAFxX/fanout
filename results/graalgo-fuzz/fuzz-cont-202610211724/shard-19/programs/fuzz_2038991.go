package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int32 }
type GenS2 struct { F0 string; F1 int }

func main() {
    var v1 int = -1
    var v2 float64 = 1e10
    v3 := []uint32{127, 0}
    var v4 int64 = -377
    var v5 int64 = (v4 * -9223372036854775808)
    var v6 int8 = 127
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS1{F0: -32768, F1: 1, F2: 42}
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: "hello", F1: 4294967295}
    if (-32768 != int16(v1)) {
        var v10 string = "hello"
        _ = v10
    }
    var v11 int16 = -185
    var v12 uint16 = 0
    var v13 int16 = -383
    v5 += (v4 >> ((255) & 63))
    var v14 int = (v1 | (v1 + (^(v1))))
    _ = len(v3)
    if true {
        switch 42 {
        case 1:
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 bool = false
                _ = fi15
                _ = v16
            }
        }
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
