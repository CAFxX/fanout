package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint64 }

func main() {
    var v1 int8 = 0
    var v2 uint = 667
    v3 := []uint32{2147483648, 2}
    var v4 int8 = v1
    v5 := GenS1{F0: 65535, F1: 127, F2: 0}
    var v6 int32 = 2147483647
    if (false && ((true || false) || ("" > ("hello" + "")))) {
        var v7 uint8 = 255
        if false {
            v3 = append(v3, 2)
            switch 32767 {
            case 128:
                v8 := GenS1{F0: 0, F1: 464, F2: 1}
                _ = v8
            }
            var v9 float32 = -2.25
            _ = v9
        }
        if true {
            v3 = append(v3, 2147483648)
            for fi10 := 0; fi10 < 16; fi10++ {
                v1 -= (v4 << (((v2 >> (((^(v2))) & 63))) & 63))
                _ = fi10
            }
        }
        _ = v7
    }
    v1 -= v1
    switch -9223372036854775808 {
    case 32767:
        if len(v3) > 0 {
            v3[0] = 811
        }
    case 2147483647:
        v1 += (v1 | (v1 ^ v4))
    case 9223372036854775807:
        fmt.Println(v1, v1)
    }
    var v11 int = 128
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v12 int16 = -1
    var v13 float32 = 0.0
    if (!(true && ((v2 + 42) <= v2))) {
        v2 = v2
        v1 = -6
    }
    fmt.Println(v11, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
