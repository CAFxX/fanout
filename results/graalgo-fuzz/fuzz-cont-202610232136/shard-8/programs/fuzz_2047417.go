package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = -1
    var v3 int32 = 42
    var v4 string = "\n\t"
    for fi5 := 0; fi5 < 1; fi5++ {
        if true {
            v6 := GenS1{F0: 674, F1: 1000}
            _ = v6
        }
        _ = fi5
    }
    switch 128 {
    case 1:
        var v7 float32 = 3.14159
        _ = v7
    case 9223372036854775807:
        v3 += (v3 | (v2 % ((v2) | 1)))
    case -1:
        switch v4 {
        case "hello":
            var v8 uint = 1
            _ = v8
        default:
            for fi9 := 0; fi9 < 9; fi9++ {
                for fi10 := 0; fi10 < 14; fi10++ {
                    v3 += 1000
                    if (true && true) {
                        var v11 string = v4
                        fmt.Println(fi10, v11)
                        _ = v11
                    }
                    v2 += 2147483647
                    v2 = (v3 & (v2 >> ((42) & 63)))
                    fmt.Println(v3, v4)
                    v3 += v2
                    _ = fi10
                }
                for fi12 := 0; fi12 < 12; fi12++ {
                    v3 = (v3 ^ (v2 & (-(v2))))
                    var v13 uint16 = 42
                    _ = fi12
                    _ = v13
                }
                _ = fi9
            }
        }
    default:
        v3 += 1
    }
    v14 := GenS2{F0: 0}
    v3 = (v2 - (v2 >> ((2) & 63)))
    var v15 int64 = -896
    var v16 int8 = -23
    v15 = v15
    v16 += 127
    v17 := GenS1{F0: 65535, F1: -1}
    var v18 int8 = int8(v15)
    fmt.Println(v18)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
