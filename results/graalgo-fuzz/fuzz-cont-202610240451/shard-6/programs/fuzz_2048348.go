package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = -936
    var v3 int64 = 0
    var v4 int32 = -792
    var v5 bool = false
    var v6 int32 = v4
    var v7 int8 = 100
    fmt.Println(v5, v4)
    var v8 uint8 = 40
    if v5 {
        var v9 int16 = -1
        _ = v9
    } else {
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 int = fi10
            switch "" {
            case "":
                fmt.Println(v3, v7)
            case "\n\t":
                fmt.Println(v8, v11)
            case "hello":
                v2 = (-(v11))
            default:
                for fi12 := 0; fi12 < 17; fi12++ {
                    var v13 uint32 = 0
                    if true {
                        var v14 int16 = 0
                        v15 := GenS1{F0: v8, F1: 702}
                        if (v13 != uint32(v8)) {
                            var v16 bool = (!v5)
                            v17 := []int{}
                            _ = v16
                            _ = v17
                        }
                        _ = v14
                        _ = v15
                    } else {
                        var v18 uint8 = (v8 | (^(v8)))
                        _ = v18
                    }
                    _ = fi12
                    _ = v13
                }
            }
            _ = fi10
            _ = v11
        }
        if (0 < 4294967295) {
            var v19 int16 = -1
            _ = v19
        } else {
            v7 = (v7 / ((v7) | 1))
            for fi20 := 0; fi20 < 8; fi20++ {
                var v21 float32 = 3.14159
                _ = fi20
                _ = v21
            }
        }
    }
    var v22 float32 = (3.14159 / (0.1 + 0.1))
    v23 := GenS1{F0: v8, F1: (v2 % (((v2 + (v2 - (v2 & (-(v2)))))) | 1))}
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 int8 = (-(v7))
        _ = fi24
        _ = v25
    }
    v8 += 181
    v1[42] = 2
    fmt.Println(v3, v5, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
