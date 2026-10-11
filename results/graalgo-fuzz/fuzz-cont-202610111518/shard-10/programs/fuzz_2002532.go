package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 int64 = -339
    v2 := []uint32{2}
    var v3 int = 4294967295
    v1 = (v1 % ((42) | 1))
    var v4 int16 = 1
    v1 -= v1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v5 := GenS1{F0: "hello", F1: 2147483647, F2: (v4 - -32768)}
    if len(v2) > 0 {
        v2[0] = 1
    }
    if true {
        fmt.Println(v1, v3)
    }
    var v6 float32 = (3.14159 / 1.5)
    if (!false) {
        var v7 float32 = v6
        for fi8 := 0; fi8 < 7; fi8++ {
            if false {
                v9 := GenS1{F0: "foo bar", F1: 4294967295, F2: v4}
                _ = v9
            } else {
                fmt.Println(v6, v1)
            }
            v10 := GenS2{F0: 99}
            _ = fi8
            _ = v10
        }
        _ = v7
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
