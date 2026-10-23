package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint64 = 0
    var v2 int32 = -2147483648
    v3 := []int{9223372036854775807}
    v4 := GenS1{F0: 0, F1: 32767, F2: 65535}
    v1 = 699
    v1 = v1
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 int32 = (v2 << ((775) & 63))
        _ = fi5
        _ = v6
    }
    _ = len(v3)
    var v7 uint64 = (v1 | (v1 % (((^(v1))) | 1)))
    v8 := GenS1{F0: 559, F1: 2, F2: 0}
    var v9 int16 = 914
    fmt.Println(v9)
    for fi10 := 0; fi10 < 11; fi10++ {
        v1 = (v7 * 1)
        _ = fi10
    }
    if len(v3) > 0 {
        v3[0] = 538
    }
    if false {
        fmt.Println(v2)
    } else {
        switch "\n\t" {
        case "hello":
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 int = (fi11 ^ (fi11 % ((-2147483648) | 1)))
                v13 := GenS2{F0: -32768}
                v14 := GenS1{F0: 1, F1: 0, F2: 639}
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
        }
        var v15 uint8 = 0
        _ = v15
    }
    v16 := GenS2{F0: v9}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
