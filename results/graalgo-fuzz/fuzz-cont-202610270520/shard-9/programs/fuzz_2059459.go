package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint16 }

func main() {
    var v1 uint32 = 1
    var v2 uint32 = (^(v1))
    var v3 uint64 = 42
    v4 := make(map[uint32]int)
    var v5 int8 = 0
    fmt.Println(v2)
    switch -9223372036854775808 {
    case 225:
        v6 := GenS1{F0: v2, F1: 0, F2: 928}
        _ = v6
    case 980:
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 uint = uint(v3)
            if ((false && (((!(2147483647 > v2)) && false) || false)) && true) {
                v1 -= (v2 / ((v2) | 1))
            } else {
                v9 := v4[(^(v2))]
                if (v8 <= v8) {
                    var v10 bool = false
                    _ = v10
                }
                _ = v9
            }
            fmt.Println(v5, v8, v2)
            _ = fi7
            _ = v8
        }
    case 128:
        v11 := GenS1{F0: 2147483647, F1: 587, F2: 974}
        _ = v11
    default:
        if true {
            fmt.Println(v2, v2)
        }
    }
    v1 = 472
    v1 -= (^(v1))
    v5 = v5
    v5 = -128
    _ = len(v4)
    var v12 uint8 = 208
    v5 -= 0
    var v13 bool = false
    var v14 bool = false
    for fi15 := 0; fi15 < 20; fi15++ {
        for fi16 := 0; fi16 < 20; fi16++ {
            switch 1 {
            case -927:
                _ = len(v4)
            }
            _ = fi16
        }
        v4[(^(v1))] = fi15
        _ = fi15
    }
    var v17 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
