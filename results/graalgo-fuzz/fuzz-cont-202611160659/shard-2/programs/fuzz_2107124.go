package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint16 }

func main() {
    var v1 uint = 588
    var v2 int = 9223372036854775807
    v3 := make(map[string]uint32)
    v3[("hello" + "a")] = 0
    fmt.Println(v1, v1)
    var v4 int64 = 483
    v1 += (v1 - 4294967295)
    if false {
        for fi5 := 0; fi5 < 12; fi5++ {
            for fi6 := 0; fi6 < 18; fi6++ {
                fmt.Println(fi6)
                _ = fi6
            }
            v7 := []int{}
            _ = fi5
            _ = v7
        }
        v2 = 0
    } else {
        v1 = (v1 & 65535)
        var v8 float32 = 1e10
        _ = v8
    }
    var v9 float64 = 3.14159
    v10 := GenS1{F0: v1, F1: false, F2: 32767}
    var v11 uint16 = 65535
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
