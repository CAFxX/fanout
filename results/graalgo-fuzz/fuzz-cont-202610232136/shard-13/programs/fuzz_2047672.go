package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = -1
    var v3 uint64 = 42
    var v4 string = "hello"
    var v5 uint32 = 4294967295
    var v6 int8 = 42
    if false {
        if (v5 > v5) {
            var v7 uint8 = 42
            _ = v7
        } else {
            v6 = -128
            v8 := GenS1{F0: 32767, F1: 42}
            _ = v8
        }
        var v9 bool = false
        _ = v9
    }
    _ = len(v1)
    if (32767 == 1) {
        v3 = (v3 % ((v3) | 1))
        for fi10 := 0; fi10 < 1; fi10++ {
            fmt.Println(fi10, v5)
            _ = fi10
        }
    }
    delete(v1, (v5 ^ 4294967295))
    v11 := v1[(^(v5))]
    for fi12 := 0; fi12 < 20; fi12++ {
        fmt.Println(v3, v11)
        _ = fi12
    }
    v2 = (v2 | int64(v5))
    if false {
        if (1.5 > (0.0 + 0.0)) {
            var v13 float64 = (3.14159 + 0.0)
            v14 := []int{int(v5), (v11 << ((564) & 63))}
            for fi15 := 0; fi15 < 5; fi15++ {
                v2 -= (v2 % (((v2 - (-(v2)))) | 1))
                _ = fi15
            }
            _ = v13
            _ = v14
        }
        var v16 float64 = 1.5
        v16 = (v16 + 0.0)
        _ = v16
    } else {
        v17 := GenS1{F0: 65535, F1: 32767}
        fmt.Println(v3, v2)
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
