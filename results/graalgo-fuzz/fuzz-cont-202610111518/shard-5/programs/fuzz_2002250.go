package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint }
type GenS2 struct { F0 int32; F1 uint; F2 int32 }

func main() {
    var v1 int64 = 1
    v2 := make(map[int]string)
    var v3 uint16 = 65535
    var v4 int32 = 0
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := v2[-2147483648]
        _ = fi5
        _ = v6
    }
    var v7 uint16 = 65535
    v3 += v7
    var v8 bool = true
    if (!v8) {
        var v9 int = 127
        for fi10 := 0; fi10 < 18; fi10++ {
            v11 := GenS1{F0: int16(v1), F1: 169, F2: 18446744073709551615}
            _ = fi10
            _ = v11
        }
        var v12 uint64 = 0
        _ = v12
        _ = v9
    }
    v1 += (v1 & (v1 | v1))
    if (127 < -28) {
        v13 := []int{int(v1), 1}
        v14 := GenS2{F0: 551, F1: 255, F2: -2147483648}
        _ = v13
        _ = v14
    }
    v15 := GenS2{F0: (v4 - -2147483648), F1: 18446744073709551615, F2: v4}
    fmt.Println(v8, v8, v4)
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 uint8 = 170
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
