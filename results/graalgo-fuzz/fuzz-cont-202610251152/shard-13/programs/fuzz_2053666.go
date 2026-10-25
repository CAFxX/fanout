package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }

func main() {
    var v1 int32 = 42
    var v2 int32 = (v1 | (v1 & v1))
    v3 := []uint32{118}
    var v4 float32 = ((0.0 * 1.5) / -2.25)
    switch ("foo bar" + (("hello" + "a") + "a")) {
    case "hello":
        for fi5 := 0; fi5 < 10; fi5++ {
            for fi6 := 0; fi6 < 16; fi6++ {
                v7 := GenS1{F0: 127}
                _ = fi6
                _ = v7
            }
            v8 := GenS1{F0: 127}
            v2 -= (v1 % (((v1 + (v2 * (v1 ^ v2)))) | 1))
            _ = fi5
            _ = v8
        }
    case "":
        var v9 float64 = float64(v4)
        _ = v9
    default:
        for fi10 := 0; fi10 < 15; fi10++ {
            _ = len(v3)
            _ = fi10
        }
    }
    if len(v3) > 0 {
        v3[0] = 42
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 uint64 = 1
        v13 := GenS1{F0: 88}
        _ = fi11
        _ = v12
        _ = v13
    }
    if len(v3) > 0 {
        v3[0] = 0
    }
    v14 := GenS1{F0: 1}
    var v15 int64 = 1
    var v16 uint = 65535
    v17 := GenS2{F0: false}
    v4 += v4
    var v18 int = 2147483647
    v18 = 2147483647
    v3 = append(v3, 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
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
