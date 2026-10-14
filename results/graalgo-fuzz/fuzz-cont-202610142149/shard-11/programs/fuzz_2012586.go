package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint32; F2 string }

func main() {
    var v1 uint8 = 255
    var v2 uint16 = 255
    var v3 uint32 = 0
    var v4 int8 = -128
    v2 -= 32767
    fmt.Println(v1, v3, v4)
    for fi5 := 0; fi5 < 17; fi5++ {
        for fi6 := 0; fi6 < 19; fi6++ {
            v4 += (v4 >> ((2) & 63))
            _ = fi6
        }
        _ = fi5
    }
    v7 := GenS1{F0: 32767, F1: 547}
    fmt.Println(v2)
    var v8 int32 = 2147483647
    v9 := GenS1{F0: 1, F1: 68}
    for fi10 := 0; fi10 < 1; fi10++ {
        switch int(v1) {
        case 128:
            fmt.Println(v4, v4)
        case 127:
            v8 = (^(v8))
        case -9223372036854775808:
            v11 := GenS1{F0: -32768, F1: 369}
            _ = v11
        }
        v12 := GenS2{F0: 605, F1: v3, F2: ("" + "")}
        _ = fi10
        _ = v12
    }
    var v13 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
