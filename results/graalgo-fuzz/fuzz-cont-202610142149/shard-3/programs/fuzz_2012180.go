package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint64 }

func main() {
    v1 := []uint32{0}
    var v2 int32 = -2147483648
    var v3 int = 1
    var v4 int64 = 897
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 uint32 = 65535
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 32767, F1: v3, F2: 0}
    v3 = (v3 << ((369) & 63))
    v4 = (v4 & v4)
    var v8 uint16 = 65535
    v9 := GenS1{F0: -222, F1: v3, F2: 18446744073709551615}
    v10 := GenS1{F0: -32768, F1: -657, F2: 18446744073709551615}
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    if (true && true) {
        v1 = append(v1, 4294967295)
        fmt.Println(v8, v3, v4)
    } else {
        v2 = (v2 | (v2 >> ((255) & 63)))
        var v11 bool = (!(42 == 187))
        for fi12 := 0; fi12 < 2; fi12++ {
            if (v11 && (2 < 42)) {
                v13 := GenS1{F0: -888, F1: fi12, F2: 650}
                _ = v13
            }
            _ = fi12
        }
        _ = v11
    }
    var v14 bool = ((127 <= (v3 % ((197) | 1))) || (0 != 2))
    var v15 int = (v3 << ((936) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
