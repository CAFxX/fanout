package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = -1
    var v3 int8 = -128
    var v4 uint16 = 42
    v2 += (^(v2))
    v4 -= 65535
    fmt.Println(v4)
    v2 += (v2 % ((-564) | 1))
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 uint8 = 128
        if false {
            var v7 uint16 = v4
            fmt.Println(v3)
            _ = v7
        } else {
            v8 := []int{(fi5 ^ fi5)}
            _ = len(v8)
            v9 := GenS1{F0: 32767}
            _ = v8
            _ = v9
        }
        fmt.Println(v3)
        _ = fi5
        _ = v6
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        v11 := GenS1{F0: 343}
        v2 -= (v2 ^ -2147483648)
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: 1}
    var v13 bool = ((v3 % (((-(v3))) | 1)) == (v3 & (-(v3))))
    v4 -= v4
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
