package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 1
    v3 := make(map[string]string)
    v4 := make(map[string]int)
    var v5 int8 = 100
    var v6 uint = 233
    v7 := GenS1{F0: 18446744073709551615}
    var v8 int16 = -1
    v9 := v1[v2]
    v6 -= (^(v6))
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := GenS1{F0: 18446744073709551615}
        _ = fi10
        _ = v11
    }
    v6 = uint(v9)
    for fi12 := 0; fi12 < 11; fi12++ {
        v8 = (v8 & (v8 & -522))
        _ = fi12
    }
    if true {
        var v13 int64 = 42
        for fi14 := 0; fi14 < 1; fi14++ {
            v6 += (v6 & (^(v6)))
            v5 = v5
            _ = fi14
        }
        if ((v6 / ((255) | 1)) != (v6 % ((4294967295) | 1))) {
            var v15 bool = false
            v16 := GenS1{F0: uint64(v6)}
            _ = v15
            _ = v16
        }
        _ = v13
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
