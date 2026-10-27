package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint }

func main() {
    var v1 float32 = float32(3.14159)
    var v2 int8 = 100
    var v3 int64 = 0
    var v4 uint16 = 2
    fmt.Println(v2, v3)
    for fi5 := 0; fi5 < 18; fi5++ {
        for fi6 := 0; fi6 < 13; fi6++ {
            v3 += 1
            _ = fi6
        }
        v2 += v2
        _ = fi5
    }
    v7 := GenS1{F0: v4, F1: v3, F2: 153}
    v8 := []int{127, 127, 0, -2147483648}
    var v9 uint16 = (v4 + (v4 - 938))
    if ((true && false) && true) {
        v1 -= float32(0.1)
        _ = len(v8)
    } else {
        for fi10 := 0; fi10 < 6; fi10++ {
            switch "hello" {
            case "hello":
                for fi11 := 0; fi11 < 11; fi11++ {
                    var v12 int32 = 2147483647
                    _ = fi11
                    _ = v12
                }
            case "foo bar":
                v13 := GenS1{F0: 65535, F1: (v3 % ((42) | 1)), F2: 1}
                _ = v13
            default:
                if false {
                    var v14 uint64 = 525
                    _ = v14
                }
            }
            var v15 string = "a"
            switch v15 {
            case "foo bar":
                if len(v8) > 0 {
                    v8[0] = -445
                }
            case "":
                if true {
                    v16 := GenS1{F0: (^(v4)), F1: v3, F2: 42}
                    v2 -= (^(v2))
                    _ = v16
                } else {
                    v9 = (v9 << ((763) & 63))
                }
            default:
                var v17 uint64 = 42
                _ = v17
            }
            _ = fi10
            _ = v15
        }
    }
    v18 := GenS1{F0: v9, F1: v3, F2: 931}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
